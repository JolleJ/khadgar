package user

import (
	"context"
	"jollej/db-scout/internal/domain/user"
)

type UserService struct {
	repo user.UserRepository
}

func NewUserService(repo user.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (u *UserService) List(ctx context.Context) []user.User {
	return u.repo.List(&ctx)
}
