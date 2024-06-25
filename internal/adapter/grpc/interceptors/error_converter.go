package interceptors

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"hexagonal-todo/internal/core/domain"
)

type ErrorConverter struct {
}

var appErrorToGrpcMap = map[domain.AppErrorCode]codes.Code{
	domain.ErrCodeInternal:     codes.Internal,
	domain.ErrCodeNotFound:     codes.NotFound,
	domain.ErrCodeBadRequest:   codes.InvalidArgument,
	domain.ErrCodeForbidden:    codes.PermissionDenied,
	domain.ErrCodeUnauthorized: codes.Unauthenticated,
}

func (e ErrorConverter) UnaryInterceptor(ctx context.Context, req any, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	resp, err := handler(ctx, req)
	if err != nil {
		err = e.convertError(err)
	}
	return resp, err
}

func (e ErrorConverter) StreamInterceptor(srv any, ss grpc.ServerStream, _ *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	err := handler(srv, ss)
	if err != nil {
		err = e.convertError(err)
	}
	return err
}

func (ErrorConverter) convertError(err error) error {
	var appError domain.AppError
	if errors.As(err, &appError) {
		grpcErrorCode, ok := appErrorToGrpcMap[appError.Code]
		if !ok {
			return status.Errorf(codes.Unknown, appError.Error())
		}
		return status.Errorf(grpcErrorCode, "%s", appError.Err)
	} else {
		return err
	}
}
