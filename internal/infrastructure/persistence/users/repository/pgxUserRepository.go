package persistence

import (
	users "client_server/internal/domain/users"
	"errors"
	"fmt"

	shared "client_server/internal/domain/shared"
	"context"

	pgx "github.com/jackc/pgx/v5"
)

type PgxUserRepository struct {
}

func NewPgxUserRepository() users.UserRepository {
	return &PgxUserRepository{}
}

func getTx(ctx context.Context) (pgx.Tx, error) {
	txIface := ctx.Value(shared.TxKey)
	tx, ok := txIface.(pgx.Tx)
	if !ok {
		return nil, errors.New("no transaction found")
	}
	return tx, nil
}

func (u *PgxUserRepository) CreateUser(ctx context.Context, user users.Users) (int, error) {
	tx, err := getTx(ctx)
	if err != nil {
		return 0, err
	}

	var id int
	if err := tx.QueryRow(
		ctx,
		"INSERT INTO users (name, email, created_at) VALUES ($1, $2, $3) RETURNING id",
		user.Name, user.Email, user.CreatedAt).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (u *PgxUserRepository) UpdateUser(ctx context.Context, user users.Users) (int, error) {
	tx, err := getTx(ctx)
	if err != nil {
		return 0, err
	}

	result, err := tx.Exec(
		ctx,
		"UPDATE users SET name = $1, email = $2 WHERE id = $3",
		user.Name, user.Email, user.ID)
	if err != nil {
		return 0, err
	}
	return int(result.RowsAffected()), nil
}

func (u *PgxUserRepository) DeleteUser(ctx context.Context, ids []int) (int, error) {
	if len(ids) == 0 {
		return 0, nil
	}
	tx, err := getTx(ctx)
	if err != nil {
		return 0, err
	}

	result, err := tx.Exec(ctx, "DELETE FROM users WHERE id = ANY($1)", ids)
	if err != nil {
		return 0, err
	}
	return int(result.RowsAffected()), nil
}

func (u *PgxUserRepository) GetUser(ctx context.Context, id int) (users.Users, error) {
	tx, err := getTx(ctx)
	if err != nil {
		return users.Users{}, err
	}

	var user users.Users
	if err := tx.QueryRow(
		ctx,
		"SELECT id, name, email, created_at FROM users WHERE id = $1",
		id).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt); err != nil {
		return users.Users{}, err
	}
	return user, nil
}

func (u *PgxUserRepository) ListUser(ctx context.Context) ([]users.Users, error) {
	tx, err := getTx(ctx)
	if err != nil {
		return nil, err
	}

	var userList []users.Users
	rows, err := tx.Query(ctx, "SELECT id, name, email, created_at FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user users.Users
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt); err != nil {
			return nil, fmt.Errorf("scan failed: %w", err)
		}
		userList = append(userList, user)
	}
	return userList, nil
}
