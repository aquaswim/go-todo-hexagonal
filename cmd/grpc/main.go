package main

import (
	"errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"hexagonal-todo/internal"
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
	lis, err := net.Listen("tcp", "0.0.0.0:5001")
	if err != nil {
		log.Fatal().Err(err).Msg("failed to listen")
	}
	log.Info().Msg("server started on 127.0.0.1:5001")
	server := internal.ContainerResolve[*grpc.Server]()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		log.Info().Msg("Gracefully shutting down...")
		err := lis.Close()
		if err != nil {
			log.Error().Err(err).Msg("error shutdown server %s")
		}
	}()

	if err := server.Serve(lis); err != nil && errors.Is(err, grpc.ErrServerStopped) {
		log.Fatal().Err(err).Msg("failed to serve")
	}

	_ = internal.ContainerNamedResolve[port.Closable]("db").Close()
}
