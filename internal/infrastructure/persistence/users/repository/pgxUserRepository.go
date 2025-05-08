package persistence

import (
	users "client_server/internal/domain/users"
	"fmt"
	http "net/http"

	database "client_server/internal/infrastructure/shared/database"

	pgx "github.com/jackc/pgx/v5"
	echo "github.com/labstack/echo/v4"
)

type PgxUserRepository struct {
	db *database.PgxPool
}

func NewPgxUserRepository(db *database.PgxPool) users.UserRepository {
	return &PgxUserRepository{db: db}
}

func getTx(ctx echo.Context) (pgx.Tx, error) {
	txIface := ctx.Get("tx")
	tx, ok := txIface.(pgx.Tx)
	if !ok {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "no transaction found")
	}
	return tx, nil
}

func (u *PgxUserRepository) CreateUser(user users.Users, ctx echo.Context) (int, error) {
	tx, err := getTx(ctx)
	if err != nil {
		return 0, err
	}

	var id int
	if err := tx.QueryRow(
		ctx.Request().Context(),
		"INSERT INTO users (name, email, created_at) VALUES ($1, $2, $3) RETURNING id",
		user.Name, user.Email, user.CreatedAt).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (u *PgxUserRepository) UpdateUser(user users.Users, ctx echo.Context) (int, error) {
	tx, err := getTx(ctx)
	if err != nil {
		return 0, err
	}

	result, err := tx.Exec(
		ctx.Request().Context(),
		"UPDATE users SET name = $1, email = $2 WHERE id = $3",
		user.Name, user.Email, user.ID)
	if err != nil {
		return 0, err
	}
	return int(result.RowsAffected()), nil
}

func (u *PgxUserRepository) DeleteUser(ids []int, ctx echo.Context) (int, error) {
	if len(ids) == 0 {
		return 0, nil
	}
	tx, err := getTx(ctx)
	if err != nil {
		return 0, err
	}

	result, err := tx.Exec(ctx.Request().Context(), "DELETE FROM users WHERE id = ANY($1)", ids)
	if err != nil {
		return 0, err
	}
	return int(result.RowsAffected()), nil
}

func (u *PgxUserRepository) GetUser(id int, ctx echo.Context) (users.Users, error) {
	tx, err := getTx(ctx)
	if err != nil {
		return users.Users{}, err
	}

	var user users.Users
	if err := tx.QueryRow(
		ctx.Request().Context(),
		"SELECT id, name, email, created_at FROM users WHERE id = $1",
		id).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt); err != nil {
		return users.Users{}, err
	}
	return user, nil
}

func (u *PgxUserRepository) ListUser(ctx echo.Context) ([]users.Users, error) {
	tx, err := getTx(ctx)
	if err != nil {
		return nil, err
	}

	var userList []users.Users
	rows, err := tx.Query(ctx.Request().Context(), "SELECT id, name, email, created_at FROM users")
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
