package rest

import (
	"client_server/internal/infrastructure/api/rest/handler"
	"client_server/internal/infrastructure/api/rest/middleware"

	echoSwagger "github.com/swaggo/echo-swagger"
)

func (s *Server) routes(h *handler.Handler, m *middleware.Middleware) {
	prefix := s.echo.Group("/api/v1")
	h.RegisterUserRoutes(prefix)
	h.RegisterAuthRoutes(prefix)
	s.echo.Use(m.TransactionMiddleware())
	s.echo.Use(m.AuthenticationMiddleware())
	s.echo.GET("/swagger/*", echoSwagger.WrapHandler)
}
