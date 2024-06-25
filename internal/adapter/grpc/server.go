package grpc

import (
	"github.com/golobby/container/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"hexagonal-todo/internal/adapter/grpc/interceptors"
	"hexagonal-todo/internal/adapter/grpc/pb"
)

func NewServer() (*grpc.Server, error) {
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

	return server, nil
}
