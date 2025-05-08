package middleware

import (
	"client_server/pkg/logger"

	transaction "client_server/internal/domain/transactions"
)

type Middleware struct {
	log       logger.Logger
	txManager transaction.TransactionService
}

func New(log logger.Logger, txManager transaction.TransactionService) *Middleware {
	return &Middleware{
		log:       log,
		txManager: txManager,
	}
}
