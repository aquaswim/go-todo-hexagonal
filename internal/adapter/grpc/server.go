package grpc

import (
	"errors"
	"github.com/golobby/container/v3"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"hexagonal-todo/internal/adapter/grpc/interceptors"
	"hexagonal-todo/internal/adapter/grpc/pb"
	"hexagonal-todo/internal/core/port"
	"net"
)

type grpcServer struct {
	// todo listen port
	lis    net.Listener
	server *grpc.Server
}

func (g *grpcServer) Start() error {
	lis, err := net.Listen("tcp", net.JoinHostPort("0.0.0.0", "5001"))
	if err != nil {
		return err
	}
	log.Info().Msgf("server started on: %s", lis.Addr())
	g.lis = lis

	if err := g.server.Serve(lis); err != nil && errors.Is(err, grpc.ErrServerStopped) {
		return err
	}
	return nil
}

func (g *grpcServer) Stop() error {
	return g.lis.Close()
}

func New() port.Server {
	log.Debug().Msg("initializing grpc server")
	logger := interceptors.Logging{}
	recov := interceptors.Recover{}
	auth := interceptors.Auth{}
	errorConverter := interceptors.ErrorConverter{}

	container.MustFill(container.Global, &auth)

	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(logger.UnaryInterceptor, recov.UnaryInterceptor, auth.UnaryInterceptor, errorConverter.UnaryInterceptor),
		grpc.ChainStreamInterceptor(logger.StreamInterceptor, recov.StreamInterceptor, auth.StreamInterceptor, errorConverter.StreamInterceptor),
	)

	var pub handlerPublic
	var withAuth handlerWithAuth

	container.MustFill(container.Global, &pub)
	container.MustFill(container.Global, &withAuth)

	pb.RegisterTodoHexagonalServiceServer(server, &pub)
	pb.RegisterTodoHexagonalServiceWithAuthServer(server, &withAuth)

	reflection.Register(server)

	return &grpcServer{
		server: server,
	}
}
