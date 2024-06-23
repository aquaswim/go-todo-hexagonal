package internal

import (
	"github.com/golobby/container/v3"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"hexagonal-todo/internal/adapter/config"
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
	container.MustSingletonLazy(container.Global, repositories.NewTodoRepo)
	container.MustSingletonLazy(container.Global, repositories.NewUserRepo)

	container.MustSingletonLazy(container.Global, tokenManager.NewJwtTokenManager)

	container.MustSingletonLazy(container.Global, service.NewTodoService)
	container.MustSingletonLazy(container.Global, service.NewAuthService)

	container.MustSingletonLazy(container.Global, func() *echo.Echo {
		server, err := restApi.New()
		if err != nil {
			panic(err)
		}
		return server
	})
}
