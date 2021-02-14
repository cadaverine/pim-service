package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"net"
	"net/http"

	gw "gitlab.com/cadaverine/pim-service/gen"
	"gitlab.com/cadaverine/pim-service/service"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"

	"github.com/jackc/pgx/v4/log/log15adapter"
	"github.com/jackc/pgx/v4/pgxpool"
	log "gopkg.in/inconshreveable/log15.v2"
)

var (
	host     = pflag.String("host", "localhost", "host of the service")
	grpcPort = pflag.String("grpc_port", ":9090", "grpc port")
	httpPort = pflag.String("http_port", ":7070", "http port")
	network  = pflag.String("network", "tcp", `one of "tcp" or "unix". Must be consistent to -endpoint`)
	dbHost   = pflag.String("db_host", "database", "")
	dbPort   = pflag.String("db_port", "5432", "")
	dbUser   = pflag.String("db_user", "postgres", "")
	dbName   = pflag.String("db_name", "pim_db", "")
	dbPass   = pflag.String("db_pass", "postgres", "")
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

	dbConf, err := getDBConfig(*dbHost, *dbPort, *dbUser, *dbPass, *dbName)
	if err != nil {
		log.Crit("Unable to create db config", "error", err)
		os.Exit(1)
	}

	log.Info("db pool creation...")

	db, err := pgxpool.ConnectConfig(ctx, dbConf)
	if err != nil {
		log.Crit("Unable to create connection pool", "error", err)
		os.Exit(1)
	}

	log.Info("db pool created")

	lis, err := net.Listen(*network, *grpcPort)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()

	svc := service.NewPimService(db)
	gw.RegisterPimServiceServer(grpcServer, svc)

	var group errgroup.Group

	group.Go(func() error {
		return grpcServer.Serve(lis)
	})

	mux := runtime.NewServeMux()

	group.Go(func() error {
		return gw.RegisterPimServiceHandlerServer(ctx, mux, svc)
	})

	group.Go(func() error {
		return http.ListenAndServe(*httpPort, mux)
	})

	log.Info(fmt.Sprintf("server listening on '%s%s'", *host, *httpPort))

	return group.Wait()
}

func getDBConfig(host, port, user, password, dbname string) (*pgxpool.Config, error) {
	dsnTemplate := "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable"

	dsn := fmt.Sprintf(dsnTemplate, host, port, user, password, dbname)

	logger := log15adapter.NewLogger(log.New("module", "pgx"))

	poolConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	poolConfig.ConnConfig.Logger = logger

	return poolConfig, nil
}
