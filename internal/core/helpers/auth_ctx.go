package helpers

import (
	"context"
	"hexagonal-todo/internal/core/domain"
)

const authCtxKey = "app_authkey"

func SetAuthCtx(ctx context.Context, data *domain.TokenData) context.Context {
	return context.WithValue(ctx, authCtxKey, data)
}

func GetAuthCtx(ctx context.Context) (*domain.TokenData, error) {
	v := ctx.Value(authCtxKey)
	if v == nil {
		return nil, domain.AppErrAuthCtxEmpty
	}
	return v.(*domain.TokenData), nil
}
