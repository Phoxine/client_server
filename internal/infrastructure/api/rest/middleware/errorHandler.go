package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (m *Middleware) ErrorHandlerMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := next(c)

			if err != nil {
				m.log.Error(err.Error())
				if he, ok := err.(*echo.HTTPError); ok {
					return c.JSON(he.Code, map[string]interface{}{
						"error": he.Message,
					})
				}

				return c.JSON(http.StatusInternalServerError, map[string]interface{}{
					"error": err.Error(),
				})
			}
			return nil
		}
	}
}
