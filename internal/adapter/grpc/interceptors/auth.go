package interceptors

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"hexagonal-todo/internal/core/domain"
	coreHelpers "hexagonal-todo/internal/core/helpers"
	"hexagonal-todo/internal/core/port"
	"strings"
)

type Auth struct {
	authService port.AuthService `container:"type"`
}

func (a Auth) UnaryInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	if a.isNeedAuth(info.FullMethod) {
		tokenData, err := a.handlerAuth(ctx)
		if err != nil {
			return nil, err
		}
		ctx = coreHelpers.SetAuthCtx(ctx, tokenData)
	}
	return handler(ctx, req)
}

func (a Auth) StreamInterceptor(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	if a.isNeedAuth(info.FullMethod) {
		_, err := a.handlerAuth(ss.Context())
		if err != nil {
			return err
		}
		// todo search how to replace context in server stream and add the tokenData
	}
	return handler(srv, ss)
}

func (a Auth) isNeedAuth(method string) bool {
	return strings.HasPrefix(method, "/grpc.TodoHexagonalServiceWithAuth/")
}

func (a Auth) handlerAuth(ctx context.Context) (*domain.TokenData, error) {
	authorizationHeader := metadata.ValueFromIncomingContext(ctx, "authorization")
	// get first value of header
	if len(authorizationHeader) == 0 {
		return nil, status.Error(codes.Unauthenticated, "authorization header is empty")
	}

	token := strings.TrimPrefix(authorizationHeader[0], "Bearer ")

	if token == "" {
		return nil, status.Error(codes.Unauthenticated, "authorization header is empty")
	}

	return a.authService.ValidateToken(ctx, token)
}
