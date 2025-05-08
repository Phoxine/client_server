// wire.go
//go:build wireinject
// +build wireinject

package main

import (
	users "client_server/internal/domain/users"
	server "client_server/internal/infrastructure/api/rest"
	pgxTransactionService "client_server/internal/infrastructure/persistence/transactions"
	pgxUserRepo "client_server/internal/infrastructure/persistence/users/repository"
	database "client_server/internal/infrastructure/shared/database"
	client_config "client_server/pkg/config"
	logger "client_server/pkg/logger"
	utils "client_server/pkg/utils"

	"github.com/google/wire"
)

func configPath() string {
	return utils.GetEnv("CONFIG_PATH", "config.yaml")
}

// InitializeServer is a Wire injector function.
func InitializeServer() (*server.Server, error) {
	wire.Build(
		server.New,
		pgxTransactionService.NewPgxTransactionService,
		pgxUserRepo.NewPgxUserRepository,
		users.NewUserService,
		database.NewPgxPool,
		logger.NewLogrusLogger,
		client_config.NewClientConfig,
		configPath,
	)
	return nil, nil
}
