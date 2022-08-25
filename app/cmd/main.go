package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"

	"go.uber.org/zap"
	"google.golang.org/grpc/reflection"

	"github.com/satttto/bookshelf/app/config"
	"github.com/satttto/bookshelf/app/server"
	"github.com/satttto/bookshelf/app/service"
)

func main() {
	os.Exit(run())
}

func run() int {
	// Config
	config, err := config.ReadConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read config : %s\n", err)
		return 1
	}

	logger, err := setupLogger(config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to setup logger: %s\n", err)
		return 1
	}

	db, err := setupDB(config)
	if err != nil {
		logger.Error("Failed to set up db")
	}

	cache, err := setupCache(config)
	if err != nil {
		logger.Error("Failed to set up cache")
	}

	bookshelfService := service.New(logger, db, cache)

	server := server.New(bookshelfService)
	if !config.IsProduction() {
		// Allowing grpcurl reqests to test/debug.
		reflection.Register(server)
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", config.GRPCPort))
	if err != nil {
		logger.Error("Failed to set up a listener")
		return 1
	}

	go func() {
		logger.Info("Started listening", zap.Int("port", config.GRPCPort))
		server.Serve(listener)
	}()

	// Shutdown with Ctrl + C
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	logger.Info("Stopped listening")
	server.GracefulStop()

	return 0
}
