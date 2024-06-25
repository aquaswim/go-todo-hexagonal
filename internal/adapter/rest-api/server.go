package restApi

import (
	"context"
	"errors"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/golobby/container/v3"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	oapiMiddleware "github.com/oapi-codegen/echo-middleware"
	"github.com/rs/zerolog/log"
	"hexagonal-todo/internal/adapter/config"
	"hexagonal-todo/internal/core/port"
	"net"
	"net/http"
)

type restServer struct {
	cfg *config.RestConfig
	e   *echo.Echo
}

func (r *restServer) Start() error {
	log.Info().
		Str("port", r.cfg.Port).
		Msg("server started")
	err := r.e.Start(net.JoinHostPort("0.0.0.0", r.cfg.Port))
	if errors.Is(err, http.ErrServerClosed) {
		return nil
	}
	return err
}

func (r *restServer) Stop() error {
	err := r.e.Shutdown(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func New(cfg *config.RestConfig) port.Server {
	log.Debug().Msg("initializing rest-api server")
	swagger, err := GetSwagger()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to load swagger spec")
	}
	// Clear out the servers array in the swagger spec, that skips validating
	// that h names match. We don't know how this thing will be run.
	swagger.Servers = nil
	handler := &h{}

	container.MustFill(container.Global, handler)

	e := echo.New()
	e.Use(loggerMiddleware())
	e.Use(echoMiddleware.Recover())
	e.Use(oapiMiddleware.OapiRequestValidatorWithOptions(swagger, &oapiMiddleware.Options{
		Options: openapi3filter.Options{
			AuthenticationFunc: newAuthenticator(),
		},
	}))
	e.HTTPErrorHandler = errorHandler

	RegisterHandlers(e, NewStrictHandler(handler, nil))

	e.HideBanner = true
	e.HidePort = true

	return &restServer{
		cfg: cfg,
		e:   e,
	}
}
