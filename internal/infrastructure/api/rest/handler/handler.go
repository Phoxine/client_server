package handler

import (
	"client_server/pkg/logger"

	users "client_server/internal/domain/users"
)

type Handler struct {
	log         logger.Logger
	userService *users.UserService
}

func New(log logger.Logger, userService *users.UserService) *Handler {
	return &Handler{
		log:         log,
		userService: userService,
	}
}
