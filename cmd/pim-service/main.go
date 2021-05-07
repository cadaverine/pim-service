package main

import (
	"context"
	"fmt"

	"net"
	"net/http"

	"gitlab.com/cadaverine/pim-service/config"
	gw "gitlab.com/cadaverine/pim-service/gen/pim-service"
	"gitlab.com/cadaverine/pim-service/helpers/db"
	"gitlab.com/cadaverine/pim-service/service"
	"go.uber.org/zap"

	"log"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/spf13/viper"

	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

var logger *zap.SugaredLogger

func init() {
	viper.AutomaticEnv()

	viper.SetDefault(config.Host, "localhost")
	viper.SetDefault(config.GrpcPort, 9090)
	viper.SetDefault(config.HttpPort, 7070)
	viper.SetDefault(config.DbHost, "localhost")
	viper.SetDefault(config.DbPort, "5432")
	viper.SetDefault(config.DbUser, "postgres")
	viper.SetDefault(config.DbName, "pim_db")
	viper.SetDefault(config.DbPass, "postgres")
	viper.SetDefault(config.DbMock, false)

	zapLogger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}

	logger = zapLogger.Sugar()
}

func main() {
	if err := run(); err != nil {
		logger.Fatalf("error: %v", err)
	}
}

func run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	dbAdp, err := db.New(ctx, viper.GetBool(config.DbMock), db.Conf{
		Host: viper.GetString(config.DbHost),
		Port: viper.GetString(config.DbPort),
		User: viper.GetString(config.DbUser),
		Pass: viper.GetString(config.DbPass),
		Name: viper.GetString(config.DbName),
	})
	if err != nil {
		return err
	}

	svc := service.NewPimService(dbAdp)

	grpcPort := viper.GetInt(config.GrpcPort)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", grpcPort))
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
	httpMux.Handle("/swagger-ui/",
		http.StripPrefix("/swagger-ui/",
			http.FileServer(http.Dir("swagger-ui/dist")),
		),
	)

	httpPort := viper.GetInt(config.HttpPort)

	group.Go(func() error {
		return http.ListenAndServe(fmt.Sprintf(":%v", httpPort), httpMux)
	})

	logger.Info(fmt.Sprintf("server listening on ':%v'", httpPort))

	return group.Wait()
}

func registerRoutes(mux *runtime.ServeMux, svc *service.PimService) {
	mux.HandlePath(http.MethodPost, "/upload-xml", svc.UploadXML)
}
