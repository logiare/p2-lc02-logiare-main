package helper

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func NewErrorResponse(message string) ErrorResponse {
	return ErrorResponse{Message: message}
}

func CustomLogger(ctx context.Context) *log.Entry {
	log.SetFormatter(&log.JSONFormatter{})
	if ctx == nil {
		return log.WithFields(log.Fields{
			"at": time.Now().Format("2006-01-02 15:04:05"),
		})
	}

	return log.WithFields(log.Fields{
		"at":         time.Now().Format("2006-01-02 15:04:05"),
		"method":     ctx.Value("method"),
		"path":       ctx.Value("path"),
		"remote":     ctx.Value("ip"),
		"user-agent": ctx.Value("user-agent"),
	})
}

func MiddlewareLogger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		ctx = context.WithValue(ctx, "ip", c.RealIP())
		ctx = context.WithValue(ctx, "user-agent", c.Request().UserAgent())
		ctx = context.WithValue(ctx, "method", c.Request().Method)
		ctx = context.WithValue(ctx, "path", c.Request().RequestURI)
		CustomLogger(ctx).Info("request")
		return next(c)
	}
}

func RespondJSON(c echo.Context, status int, data interface{}) error {
	return c.JSON(status, data)
}

func RespondError(c echo.Context, status int, message string) error {
	return c.JSON(status, ErrorResponse{Message: message})
}

// WriteJSON is kept for compatibility
func WriteJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
}
