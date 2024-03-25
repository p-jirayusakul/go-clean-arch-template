package middleware

import (
	"context"
	"log/slog"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/p-jirayusakul/go-clean-arch-template/pkg/common"
)

func LogHandler(logger *slog.Logger) echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogRequestID:     true,
		LogRemoteIP:      true,
		LogURI:           true,
		LogHost:          true,
		LogMethod:        true,
		LogUserAgent:     true,
		LogStatus:        true,
		LogError:         true,
		LogLatency:       true,
		LogContentLength: true,
		Skipper: func(c echo.Context) bool {
			// Skip middleware if path is equal health check
			if c.Request().URL.Path == "/" || c.Request().URL.Path == "" {
				return true
			} else if strings.Contains(c.Request().URL.Path, common.DOCS_URL) {
				return true
			}
			return false
		},
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			var errMsg string
			var logLevel slog.Level

			if v.Error != nil {
				errMsg = v.Error.Error()
				logLevel = slog.LevelError
			} else {
				logLevel = slog.LevelInfo
			}

			logger.LogAttrs(context.Background(), logLevel, "REQUEST",
				slog.String("id", v.RequestID),
				slog.String("remote_ip", v.RemoteIP),
				slog.String("host", v.Host),
				slog.String("method", v.Method),
				slog.String("uri", v.URI),
				slog.String("user_agent", v.UserAgent),
				slog.Int("status", v.Status),
				slog.String("error", errMsg),
				slog.String("latency", v.Latency.String()),
			)
			return nil
		},
	})
}
