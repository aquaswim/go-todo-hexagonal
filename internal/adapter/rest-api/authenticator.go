package restApi

import (
	"context"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/golobby/container/v3"
	oapiMiddleware "github.com/oapi-codegen/echo-middleware"
	"hexagonal-todo/internal/core/domain"
	coreHelpers "hexagonal-todo/internal/core/helpers"
	"hexagonal-todo/internal/core/port"
	"strings"
)

const bearerPrefix = "Bearer "

func newAuthenticator() openapi3filter.AuthenticationFunc {
	var tokenManager port.TokenManager
	container.MustResolve(container.Global, &tokenManager)

	return func(ctx context.Context, input *openapi3filter.AuthenticationInput) error {
		if input.SecuritySchemeName != "BearerAuth" {
			return domain.NewAppErrorString(domain.ErrCodeInternal, "Security scheme is not bearer auth")
		}

		// get token from request
		authHeader := input.RequestValidationInput.Request.Header.Get("Authorization")
		if authHeader == "" {
			return domain.AppErrTokenNotProvided
		}

		// token should have prefix "Bearer "
		if !strings.HasPrefix(authHeader, bearerPrefix) {
			return domain.AppErrTokenInvalidProvided
		}
		token := strings.TrimPrefix(authHeader, bearerPrefix)

		tokenData, err := tokenManager.DecodeToken(ctx, token)
		if err != nil {
			return err
		}

		// set token data to context that can be retrieved in request handler
		eCtx := oapiMiddleware.GetEchoContext(ctx)
		eCtx.SetRequest(
			eCtx.Request().WithContext(
				coreHelpers.SetAuthCtx(eCtx.Request().Context(), tokenData),
			),
		)

		return nil
	}
}
