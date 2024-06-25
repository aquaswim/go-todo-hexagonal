package grpc

import (
	"github.com/golobby/container/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"hexagonal-todo/internal/adapter/grpc/pb"
)

func NewServer() (*grpc.Server, error) {
	server := grpc.NewServer()

	var handler h

	container.MustFill(container.Global, &handler)

	pb.RegisterTodoHexagonalServiceServer(server, &handler)

	reflection.Register(server)
	// todo add error handling, logging

	return server, nil
}
