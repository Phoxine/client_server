package database

import (
	"fmt"

	"context"

	client_config "client_server/pkg/config"

	transactions "client_server/internal/domain/transactions"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PgxPool struct {
	Pool *pgxpool.Pool
}

func NewPgxPool(config *client_config.ClientConfig) *PgxPool {
	return &PgxPool{
		Pool: generatePool(config),
	}
}

func generatePool(config *client_config.ClientConfig) *pgxpool.Pool {
	host := config.Postgres.Host
	user := config.Postgres.User
	password := config.Postgres.Password
	dbName := config.Postgres.DBName
	port := config.Postgres.Port

	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, dbName)
	pool, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		panic(err)
	}
	return pool
}

func (pool *PgxPool) GetTransaction(ctx context.Context) (transactions.Transaction, error) {
	return pool.Pool.Begin(ctx)
}
