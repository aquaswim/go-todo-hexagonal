package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"hexagonal-todo/internal"
	"hexagonal-todo/internal/core/port"
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
	server := internal.ContainerNamedResolve[port.Server]("grpc")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		log.Info().Msg("Gracefully shutting down...")
		err := server.Stop()
		if err != nil {
			log.Error().Err(err).Msg("error shutdown server %s")
		}
	}()

	if err := server.Start(); err != nil {
		log.Fatal().Err(err).Msg("failed to serve")
	}

	_ = internal.ContainerNamedResolve[port.Closable]("db").Close()
}
