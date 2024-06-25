package grpc

import (
	"context"
	"hexagonal-todo/internal/adapter/grpc/pb"
	"hexagonal-todo/internal/core/domain"
	"hexagonal-todo/internal/core/port"
)

type h struct {
	pb.UnimplementedTodoHexagonalServiceServer
	authService port.AuthService `container:"type"`
}

func (h h) GetHealth(_ context.Context, _ *pb.HealthCheck_Payload) (*pb.HealthCheck_Result, error) {
	return &pb.HealthCheck_Result{Healthy: true}, nil
}

func (h h) AuthLogin(ctx context.Context, credential *pb.Auth_LoginCredential) (*pb.Auth_LoginResult, error) {
	res, err := h.authService.Login(ctx, &domain.LoginCredential{
		Email:    credential.GetEmail(),
		Password: credential.GetPassword(),
	})
	if err != nil {
		return nil, err
	}

	return &pb.Auth_LoginResult{Token: res.Token}, nil
}

func (h h) AuthRegister(ctx context.Context, data *pb.Auth_RegisterData) (*pb.Auth_UserData, error) {
	res, err := h.authService.Register(ctx, &domain.UserData{
		LoginCredential: domain.LoginCredential{
			Email:    data.Email,
			Password: data.Password,
		},
		FullName: data.FullName,
	})
	if err != nil {
		return nil, err
	}
	return &pb.Auth_UserData{
		Id:       res.Id,
		Email:    res.Email,
		Password: res.Password,
		FullName: res.FullName,
	}, nil
}
