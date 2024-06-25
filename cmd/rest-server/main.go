package main

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"hexagonal-todo/internal"
	"hexagonal-todo/internal/adapter/config"
	"hexagonal-todo/internal/core/port"
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
	server := internal.ContainerResolve[*echo.Echo]()

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
	restCfg := internal.ContainerResolve[*config.RestConfig]()

	server.HideBanner = true
	server.HidePort = true

	log.Info().
		Str("port", restCfg.Port).
		Msg("server started")
	_ = server.Start(net.JoinHostPort("0.0.0.0", restCfg.Port))
	// do other cleanup here
	_ = internal.ContainerNamedResolve[port.Closable]("db").Close()
}
