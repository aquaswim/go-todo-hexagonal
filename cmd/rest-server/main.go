package main

import (
	"context"
	"github.com/golobby/container/v3"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"hexagonal-todo/internal"
	"hexagonal-todo/internal/adapter/config"
	"net"
	"os"
	"os/signal"
)

func main() {
	internal.InitContainer()

	var server *echo.Echo

	container.MustResolve(container.Global, &server)

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
	var cfg *config.RestConfig
	container.MustResolve(container.Global, &cfg)

	_ = server.Start(net.JoinHostPort("0.0.0.0", cfg.Port))
	// do other cleanup here
	container.MustCall(container.Global, func(pgPool *pgxpool.Pool) {
		pgPool.Close()
	})
}
