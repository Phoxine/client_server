package middleware

import (
	"context"
	"net/http"

	shared "client_server/internal/domain/shared"

	"github.com/labstack/echo/v4"
)

func (m *Middleware) TransactionMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := c.Request().Context()
			tx, err := m.txManager.Begin(ctx)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, "could not begin transaction")
			}

			defer func() {
				if p := recover(); p != nil {
					_ = tx.Rollback(ctx)
					panic(p)
				}
			}()
			ctx = context.WithValue(ctx, shared.TxKey, tx)
			c.SetRequest(c.Request().WithContext(ctx))

			if err := next(c); err != nil {
				_ = tx.Rollback(ctx)
				return err
			}

			if err := tx.Commit(ctx); err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, "could not commit transaction")
			}
			return nil
		}
	}
}
