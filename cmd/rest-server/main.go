package main

import (
	"context"
	"hexagonal-todo/internal/adapter/config"
	restApi "hexagonal-todo/internal/adapter/rest-api"
	"hexagonal-todo/internal/adapter/storage/pgsql"
	"hexagonal-todo/internal/adapter/storage/pgsql/repositories"
	"hexagonal-todo/internal/core/service"
	"net"
	"os"
	"os/signal"
)

func main() {
	cfg := config.New()

	pgPool, err := pgsql.Connect(cfg)
	if err != nil {
		panic(err)
	}

	todoRepo := repositories.NewTodoRepo(pgPool)
	todoService := service.NewTodoService(todoRepo)

	server, err := restApi.New(
		todoService,
	)
	if err != nil {
		panic(err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		server.Logger.Infof("Gracefully shutting down...")
		err := server.Shutdown(context.Background())
		if err != nil {
			server.Logger.Errorf("error shutdown server %s", err)
		}
	}()

	// And we serve HTTP until the world ends.
	port := "8080"
	_ = server.Start(net.JoinHostPort("0.0.0.0", port))
	// do other cleanup here
}
