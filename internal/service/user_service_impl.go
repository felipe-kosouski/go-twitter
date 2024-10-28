package service

import (
	"context"
	"go_twitter/internal/models"
	"go_twitter/internal/repository"
)

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (u userService) GetUsers(ctx context.Context) ([]*models.User, error) {
	return u.userRepo.GetUsers(ctx)
}

func (u userService) GetUserById(ctx context.Context, id int) (*models.User, error) {
	return u.userRepo.GetUserById(ctx, id)
}

func (u userService) CreateUser(ctx context.Context, user *models.User) error {
	return u.userRepo.CreateUser(ctx, user)
}

func (u userService) UpdateUser(ctx context.Context, user *models.User) error {
	return u.userRepo.UpdateUser(ctx, user)
}

func (u userService) DeleteUser(ctx context.Context, id int) error {
	return u.userRepo.DeleteUser(ctx, id)
}
