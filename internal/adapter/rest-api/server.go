package restApi

import (
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/golobby/container/v3"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	oapiMiddleware "github.com/oapi-codegen/echo-middleware"
)

func New() (*echo.Echo, error) {
	swagger, err := GetSwagger()
	if err != nil {
		return nil, err
	}
	// Clear out the servers array in the swagger spec, that skips validating
	// that h names match. We don't know how this thing will be run.
	swagger.Servers = nil
	handler := &h{}

	container.MustFill(container.Global, handler)

	e := echo.New()
	//todo: proper global logger
	e.Use(echoMiddleware.Logger())
	e.Use(echoMiddleware.Recover())
	e.Use(oapiMiddleware.OapiRequestValidatorWithOptions(swagger, &oapiMiddleware.Options{
		Options: openapi3filter.Options{
			AuthenticationFunc: newAuthenticator(),
		},
	}))
	e.HTTPErrorHandler = errorHandler

	RegisterHandlers(e, NewStrictHandler(handler, nil))

	return e, err
}
