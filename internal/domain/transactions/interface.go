package transactions

import "context"

type Transaction interface {
	Commit(context.Context) error
	Rollback(context.Context) error
}

type TransactionService interface {
	Begin(context.Context) (Transaction, error)
}
