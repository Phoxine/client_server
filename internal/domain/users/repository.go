package users

import (
	echo "github.com/labstack/echo/v4"
)

type UserRepository interface {
	GetUser(int, echo.Context) (Users, error)
	ListUser(echo.Context) ([]Users, error)
	CreateUser(Users, echo.Context) (int, error)
	UpdateUser(Users, echo.Context) (int, error)
	DeleteUser([]int, echo.Context) (int, error)
}
