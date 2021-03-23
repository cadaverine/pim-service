package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"net"
	"net/http"

	gw "gitlab.com/cadaverine/pim-service/gen/pim-service"
	"gitlab.com/cadaverine/pim-service/helpers/db"
	"gitlab.com/cadaverine/pim-service/service"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"

	log "gopkg.in/inconshreveable/log15.v2"
)

var (
	host     = pflag.String("host", "localhost", "host of the service")
	grpcPort = pflag.String("grpc_port", ":9090", "grpc port")
	httpPort = pflag.String("http_port", ":7070", "http port")
	network  = pflag.String("network", "tcp", `one of "tcp" or "unix". Must be consistent to -endpoint`)
	dbHost   = pflag.String("db_host", "localhost", "")
	dbPort   = pflag.String("db_port", "5432", "")
	dbUser   = pflag.String("db_user", "postgres", "")
	dbName   = pflag.String("db_name", "pim_db", "")
	dbPass   = pflag.String("db_pass", "postgres", "")
	dbMock   = pflag.Bool("db_mock", false, "use db mock")
)

func init() {
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.AutomaticEnv()

	pflag.PrintDefaults()
}

func main() {
	if err := run(); err != nil {
		log.Crit("error", err)
		os.Exit(1)
	}
}

func run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	dbAdp, err := db.New(ctx, *dbMock, db.Conf{
		Host: *dbHost,
		Port: *dbPort,
		User: *dbUser,
		Pass: *dbPass,
		Name: *dbName,
	})
	if err != nil {
		return err
	}

	svc := service.NewPimService(dbAdp)

	lis, err := net.Listen(*network, *grpcPort)
	if err != nil {
		return err
	}
	defer lis.Close()

	grpcServer := grpc.NewServer()

	gw.RegisterPimServiceServer(grpc.NewServer(), svc)

	var group errgroup.Group

	group.Go(func() error {
		return grpcServer.Serve(lis)
	})

	gwMux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			EmitUnpopulated: false,
		},
	}))

	registerRoutes(gwMux, svc)

	err = gw.RegisterPimServiceHandlerServer(ctx, gwMux, svc)
	if err != nil {
		return err
	}

	httpMux := http.NewServeMux()
	httpMux.Handle("/", gwMux)
	httpMux.Handle("/swagger",
		http.StripPrefix("/swagger",
			http.FileServer(http.Dir("swagger-ui/dist")),
		),
	)

	group.Go(func() error {
		return http.ListenAndServe(*httpPort, httpMux)
	})

	log.Info(fmt.Sprintf("server listening on '%s%s'", *host, *httpPort))

	return group.Wait()
}

func registerRoutes(mux *runtime.ServeMux, svc *service.PimService) {
	mux.HandlePath(http.MethodPost, "/upload-xml", svc.UploadXML)
}
