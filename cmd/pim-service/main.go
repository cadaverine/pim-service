package main

import (
	"context"
	"flag"
	"net"
	"net/http"

	gw "gitlab.com/cadaverine/pim-service/gen"
	"gitlab.com/cadaverine/pim-service/service"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

var (
	host     = flag.String("host", "localhost", "host of the service")
	grpcPort = flag.String("grpc_port", ":9090", "grpc port")
	httpPort = flag.String("http_port", ":7070", "http port")
	network  = flag.String("network", "tcp", `one of "tcp" or "unix". Must be consistent to -endpoint`)
)

func init() {
	flag.Parse()
	flag.PrintDefaults()
}

func main() {
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}

func run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	lis, err := net.Listen(*network, *grpcPort)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()

	svc := service.NewPimService()
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

	glog.V(2).Infof("server listening on '%s%s'", *host, *httpPort)

	return group.Wait()
}
