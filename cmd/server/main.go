package main

import (
	"flag"
	"fmt"
	"github.com/go-kit/kit/log"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"github.com/peterbourgon/ff"
	"github.com/tchaudhry91/bottle/crate"
	"github.com/tchaudhry91/bottle/crate/pb"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fs := flag.NewFlagSet("crate", flag.ExitOnError)
	var (
		grpcAddr = fs.String("grpc-addr", "127.0.0.1:14999", "address to bind on")
		dbFile   = fs.String("db-file", "crate.db", "Local file to store bottles in")
	)

	err := ff.Parse(fs, os.Args[1:], ff.WithEnvVarNoPrefix())
	if err != nil {
		panic(err)
	}

	var logger log.Logger
	{
		logger = log.NewJSONLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	// Database
	var store crate.BottleStore
	if *dbFile != "" {
		store, err = crate.NewBoltBottleStore(*dbFile)
		if err != nil {
			panic(err)
		}
	}

	// Base Service
	var svc crate.CrateService
	svc = crate.NewCrateService(store)

	// Logging Middleware
	svc = crate.NewLoggingMiddleware(logger, svc)

	// gRPC Transport
	grpcServer := crate.NewGRPCServer(crate.MakeServerEndpoints(svc), logger)

	// gRPC Listener
	listener, err := net.Listen("tcp", *grpcAddr)
	if err != nil {
		logger.Log("transport", "grpc", "during", "listen", "err", err)
		panic(err)
	}

	// Base Server
	baseServer := grpc.NewServer(grpc.UnaryInterceptor(kitgrpc.Interceptor))
	pb.RegisterCrateServer(baseServer, grpcServer)

	// Make Shutdown Channel
	shutdown := make(chan error, 1)
	// Make Interrupt Channel
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	go func() {
		logger.Log("transport", "gRPC", "addr", *grpcAddr)
		err = baseServer.Serve(listener)
		shutdown <- err
	}()

	select {
	case signalKill := <-interrupt:
		logger.Log("msg", fmt.Sprintf("Stopping Server: %s", signalKill.String()))
	case err := <-shutdown:
		logger.Log("error", err)
	}

	err = listener.Close()
	if err != nil {
		logger.Log("error", err)
	}

}
