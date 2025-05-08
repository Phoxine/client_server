package users

import (
	echo "github.com/labstack/echo/v4"
)

type UserService struct {
	userRepo UserRepository
}

func NewUserService(userRepo UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) GetUser(id int, ctx echo.Context) (Users, error) {
	return s.userRepo.GetUser(id, ctx)
}

func (s *UserService) ListUser(ctx echo.Context) ([]Users, error) {
	return s.userRepo.ListUser(ctx)
}

func (s *UserService) CreateUser(user Users, ctx echo.Context) (int, error) {
	return s.userRepo.CreateUser(user, ctx)
}

func (s *UserService) UpdateUser(user Users, ctx echo.Context) (int, error) {
	return s.userRepo.UpdateUser(user, ctx)
}

func (s *UserService) DeleteUser(ids []int, ctx echo.Context) (int, error) {
	return s.userRepo.DeleteUser(ids, ctx)
}
