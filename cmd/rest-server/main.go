package main

import (
	"context"
	"github.com/golobby/container/v3"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"hexagonal-todo/internal"
	"hexagonal-todo/internal/adapter/config"
	"net"
	"os"
	"os/signal"
	"time"
)

func init() {
	//zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Logger.Level(zerolog.DebugLevel).Output(zerolog.NewConsoleWriter(func(w *zerolog.ConsoleWriter) {
		w.TimeFormat = time.RFC3339
		//w.NoColor = true
	}))

	// boot the container
	internal.InitContainer()
}

func main() {
	var server *echo.Echo

	container.MustResolve(container.Global, &server)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		log.Info().Msg("Gracefully shutting down...")
		err := server.Shutdown(context.Background())
		if err != nil {
			log.Error().Err(err).Msg("error shutdown server %s")
		}
	}()

	// And we serve HTTP until the world ends.
	var cfg *config.RestConfig
	container.MustResolve(container.Global, &cfg)

	server.HideBanner = true
	server.HidePort = true

	log.Info().
		Str("port", cfg.Port).
		Msg("server started")
	_ = server.Start(net.JoinHostPort("0.0.0.0", cfg.Port))
	// do other cleanup here
	container.MustCall(container.Global, func(pgPool *pgxpool.Pool) {
		pgPool.Close()
	})
}
