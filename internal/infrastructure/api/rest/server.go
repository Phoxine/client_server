package rest

import (
	transaction "client_server/internal/domain/transactions"
	users "client_server/internal/domain/users"
	"client_server/internal/infrastructure/api/rest/handler"
	"client_server/internal/infrastructure/api/rest/middleware"
	client_config "client_server/pkg/config"
	"client_server/pkg/logger"
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Server struct {
	echo *echo.Echo
	log  logger.Logger
	cfg  *client_config.ClientConfig
}

func New(
	userService *users.UserService,
	txManager transaction.TransactionService,
	logger logger.Logger,
	cfg *client_config.ClientConfig,
) (*Server, error) {

	server := &Server{
		echo: echo.New(),
		log:  logger,
		cfg:  cfg,
	}

	server.routes(
		handler.New(server.log, userService),
		middleware.New(server.log, txManager),
	)

	return server, nil
}

func (s *Server) Start(ctx context.Context) error {
	errorChan := make(chan error, 1)
	defer s.log.Flush()
	go s.start(errorChan)
	s.log.Info("Starting HTTP server on port: ", s.cfg.ServerConfig.Port)
	s.log.Info("Is env production: ", s.cfg.ServerConfig.IsProduction)
	select {
	case <-ctx.Done():
		s.log.Info("Shutting down the server")
		if shutdownErr := s.echo.Shutdown(ctx); shutdownErr != nil {
			s.log.Error(fmt.Sprintf("Error shutting down the server: %v", shutdownErr))
			return shutdownErr
		}
	case err := <-errorChan:
		s.log.Fatal(fmt.Sprintf("Failed to start HTTP server: %v", err))
		return err
	}

	return nil
}

func (s *Server) start(errorChan chan<- error) {
	defer close(errorChan)
	if err := s.echo.Start(
		fmt.Sprintf(":%d", s.cfg.ServerConfig.Port),
	); err != nil && !errors.Is(err, http.ErrServerClosed) {
		errorChan <- err
	}
}
