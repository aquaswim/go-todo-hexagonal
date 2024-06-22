package restApi

import (
	"errors"
	"fmt"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/labstack/echo/v4"
	"hexagonal-todo/internal/core/domain"
	"net/http"
)

func errorHandler(err error, c echo.Context) {
	httpCode := http.StatusInternalServerError
	responseBody := Error{
		Code:    "-",
		Message: "Internal Server Error",
	}

	var HTTPError *echo.HTTPError
	var appError domain.AppError
	var openapiReqError *openapi3filter.RequestError
	switch {
	case errors.As(err, &openapiReqError):
		httpCode = http.StatusBadRequest
		responseBody.Code = domain.ErrCodeBadRequest.String()
		responseBody.Message = openapiReqError.Reason
	case errors.As(err, &appError):
		httpCode = appErrorToHTTPCode(&appError)
		responseBody.Code = appError.Code.String()
		responseBody.Message = appError.Err.Error()
	case errors.As(err, &HTTPError):
		httpCode = HTTPError.Code
		responseBody.Code = fmt.Sprintf("H_%d", HTTPError.Code)
		responseBody.Message = ""
		responseBody.Detail = &HTTPError.Message
	}

	err = c.JSON(httpCode, responseBody)
	if err != nil {
		c.Logger().Errorf("failed to send error response: %s", err)
	}
}

func appErrorToHTTPCode(appError *domain.AppError) int {
	if appError == nil {
		return http.StatusInternalServerError
	}
	switch appError.Code {
	case domain.ErrCodeNotFound:
		return http.StatusNotFound
	case domain.ErrCodeBadRequest:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
