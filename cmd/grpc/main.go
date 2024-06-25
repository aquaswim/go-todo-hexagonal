package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"hexagonal-todo/internal"
	"net"
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

	if err := server.Serve(lis); err != nil {
		log.Fatal().Err(err).Msg("failed to serve")
	}
}
