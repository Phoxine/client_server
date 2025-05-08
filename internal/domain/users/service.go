package users

import (
	"context"
)

type UserService struct {
	userRepo UserRepository
}

func NewUserService(userRepo UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) GetUser(ctx context.Context, id int) (Users, error) {
	return s.userRepo.GetUser(ctx, id)
}

func (s *UserService) ListUser(ctx context.Context) ([]Users, error) {
	return s.userRepo.ListUser(ctx)
}

func (s *UserService) CreateUser(ctx context.Context, user Users) (int, error) {
	return s.userRepo.CreateUser(ctx, user)
}

func (s *UserService) UpdateUser(ctx context.Context, user Users) (int, error) {
	return s.userRepo.UpdateUser(ctx, user)
}

func (s *UserService) DeleteUser(ctx context.Context, ids []int) (int, error) {
	return s.userRepo.DeleteUser(ctx, ids)
}
