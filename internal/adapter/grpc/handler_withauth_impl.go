package grpc

import (
	"context"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	"hexagonal-todo/internal/adapter/grpc/pb"
	"hexagonal-todo/internal/core/port"
)

type handlerWithAuth struct {
	pb.UnimplementedTodoHexagonalServiceWithAuthServer
	authService port.AuthService `container:"type"`
}

func (h handlerWithAuth) AuthProfile(ctx context.Context, _ *emptypb.Empty) (*pb.Auth_UserData, error) {
	profile, err := h.authService.MyProfile(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.Auth_UserData{
		Id:       profile.Id,
		Email:    profile.Email,
		Password: "-redacted-",
		FullName: profile.FullName,
	}, nil
}
