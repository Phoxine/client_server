package users

import (
	"context"
)

type UserRepository interface {
	GetUser(context.Context, int) (Users, error)
	ListUser(context.Context) ([]Users, error)
	CreateUser(context.Context, Users) (int, error)
	UpdateUser(context.Context, Users) (int, error)
	DeleteUser(context.Context, []int) (int, error)
}
