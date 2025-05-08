package persistence

import (
	transaction "client_server/internal/domain/transactions"
	"context"

	database "client_server/internal/infrastructure/shared/database"

	"github.com/jackc/pgx/v5"
)

type PgxTransactionService struct {
	db *database.PgxPool
}

type Tx struct {
	pgx.Tx
}

func NewPgxTransactionService(db *database.PgxPool) transaction.TransactionService {
	return &PgxTransactionService{db: db}
}

func (t *PgxTransactionService) Begin(ctx context.Context) (transaction.Transaction, error) {
	tx, err := t.db.Pool.Begin(ctx)
	if err != nil {
		return nil, err
	}
	return &Tx{Tx: tx}, nil
}

func (t *Tx) Commit(ctx context.Context) error {
	return t.Tx.Commit(ctx)
}

func (t *Tx) Rollback(ctx context.Context) error {
	return t.Tx.Rollback(ctx)
}
