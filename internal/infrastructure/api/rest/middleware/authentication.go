package middleware

import (
	"net/http"
	"strings"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func (m *Middleware) AuthenticationMiddleware() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey:   []byte("secret"),
		Skipper:      skipper,
		ErrorHandler: errorHandler,
	})
}

func skipper(c echo.Context) bool {
	return strings.HasPrefix(c.Request().URL.Path, "/swagger") || c.Request().URL.Path == "/api/v1/auth/login"
}

func errorHandler(c echo.Context, err error) error {
	if c.Request().Header.Get("Authorization") == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "no authorization header").SetInternal(err)
	}
	return echo.NewHTTPError(http.StatusUnauthorized, "invalid or expired jwt").SetInternal(err)
}
