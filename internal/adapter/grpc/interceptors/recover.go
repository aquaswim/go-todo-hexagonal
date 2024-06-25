package interceptors

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Recover struct {
}

func (Recover) UnaryInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	defer func() {
		if errr := recover(); errr != nil {
			err = status.Errorf(codes.Internal, "panic: %v", errr)
		}
	}()
	return handler(ctx, req)
}

func (Recover) StreamInterceptor(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
	defer func() {
		if errr := recover(); errr != nil {
			err = status.Errorf(codes.Internal, "panic: %v", errr)
		}
	}()
	return handler(srv, ss)
}
