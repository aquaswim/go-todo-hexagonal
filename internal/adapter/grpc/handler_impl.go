package grpc

import (
	"context"
	"hexagonal-todo/internal/adapter/grpc/pb"
)

type h struct {
	pb.UnimplementedTodoHexagonalServiceServer
}

func (h h) GetHealth(_ context.Context, _ *pb.HealthCheck_Payload) (*pb.HealthCheck_Result, error) {
	return &pb.HealthCheck_Result{Healthy: true}, nil
}
