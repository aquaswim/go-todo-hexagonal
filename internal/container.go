package internal

import (
	"github.com/golobby/container/v3"
	"github.com/jackc/pgx/v5/pgxpool"
	"hexagonal-todo/internal/adapter/config"
	"hexagonal-todo/internal/adapter/grpc"
	restApi "hexagonal-todo/internal/adapter/rest-api"
	"hexagonal-todo/internal/adapter/storage/pgsql"
	"hexagonal-todo/internal/adapter/storage/pgsql/repositories"
	tokenManager "hexagonal-todo/internal/adapter/token_manager"
	"hexagonal-todo/internal/core/service"
)

func InitContainer() {
	container.MustSingletonLazy(container.Global, config.RestConfigFromENV)
	container.MustSingletonLazy(container.Global, config.DBConfigFromENV)
	container.MustSingletonLazy(container.Global, config.JwtConfigFromENV)

	container.MustSingletonLazy(container.Global, func(cfg *config.DBConfig) *pgxpool.Pool {
		pool, err := pgsql.Connect(cfg)
		if err != nil {
			panic(err)
		}
		return pool
	})
	container.MustNamedSingletonLazy(container.Global, "db", pgsql.NewCloser)

	container.MustSingletonLazy(container.Global, repositories.NewTodoRepo)
	container.MustSingletonLazy(container.Global, repositories.NewUserRepo)

	container.MustSingletonLazy(container.Global, tokenManager.NewJwtTokenManager)

	container.MustSingletonLazy(container.Global, service.NewTodoService)
	container.MustSingletonLazy(container.Global, service.NewAuthService)

	container.MustNamedSingletonLazy(container.Global, "rest", restApi.New)

	container.MustNamedSingletonLazy(container.Global, "grpc", grpc.New)
}

func ContainerResolve[T any]() T {
	var t T
	container.MustResolve(container.Global, &t)
	return t
}
func ContainerNamedResolve[T any](name string) T {
	var t T
	container.MustNamedResolve(container.Global, &t, name)
	return t
}
