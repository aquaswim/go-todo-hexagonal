package restApi

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func loggerMiddleware() echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogMethod:    true,
		LogURI:       true,
		LogStatus:    true,
		LogRemoteIP:  true,
		LogUserAgent: true,
		LogError:     true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			level := zerolog.InfoLevel
			if v.Error != nil {
				level = zerolog.ErrorLevel
			}
			log.WithLevel(level).
				Str("method", v.Method).
				Str("URI", v.URI).
				Int("status", v.Status).
				Dur("duration", v.Latency).
				Str("ip", v.RemoteIP).
				Str("user_agent", v.UserAgent).
				Err(v.Error).
				Msg("request")

			return nil
		},
	})
}
