package service

import (
	"context"
	"go_twitter/internal/models"
)

type UserService interface {
	GetUsers(ctx context.Context) ([]*models.User, error)
	GetUserById(ctx context.Context, id int) (*models.User, error)
	CreateUser(ctx context.Context, user *models.User) error
	UpdateUser(ctx context.Context, user *models.User) error
	DeleteUser(ctx context.Context, id int) error
}
