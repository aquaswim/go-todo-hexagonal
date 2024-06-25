package interceptors

import (
	"context"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"time"
)

type Logging struct{}

func (Logging) UnaryInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	start := time.Now()
	res, err := handler(ctx, req)
	log.Info().
		Str("method", info.FullMethod).
		Dur("duration", time.Since(start)).
		Any("UA", md.Get("user-agent")).
		Err(err).
		Msg("Unary Request")
	return res, err
}

func (Logging) StreamInterceptor(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	md, _ := metadata.FromIncomingContext(ss.Context())
	start := time.Now()
	err := handler(srv, ss)
	log.Info().
		Str("method", info.FullMethod).
		Bool("is client stream", info.IsClientStream).
		Bool("is server stream", info.IsServerStream).
		Dur("duration", time.Since(start)).
		Any("UA", md.Get("user-agent")).
		Err(err).
		Msg("Stream Request")
	return err
}
