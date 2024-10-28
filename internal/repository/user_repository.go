package repository

import (
	"context"
	"github.com/jackc/pgx/v4"
	"go_twitter/internal/models"
)

type UserRepository interface {
	GetUsers(ctx context.Context) ([]*models.User, error)
	GetUserById(ctx context.Context, id int) (*models.User, error)
	CreateUser(ctx context.Context, user *models.User) error
	UpdateUser(ctx context.Context, user *models.User) error
	DeleteUser(ctx context.Context, id int) error
}

type userRepository struct {
	db *pgx.Conn
}

func NewUserRepository(db *pgx.Conn) UserRepository {
	return &userRepository{db: db}
}

func (u userRepository) GetUsers(ctx context.Context) ([]*models.User, error) {
	rows, err := u.db.Query(ctx, "SELECT id, username, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		user := &models.User{}
		err := rows.Scan(&user.ID, &user.Username, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func (u userRepository) CreateUser(ctx context.Context, user *models.User) error {
	_, err := u.db.Exec(ctx, "INSERT INTO users (username, email) VALUES ($1, $2)", user.Username, user.Email)
	return err
}

func (u userRepository) GetUserById(ctx context.Context, id int) (*models.User, error) {
	user := &models.User{}
	err := u.db.QueryRow(ctx, "SELECT id, username, email FROM users WHERE id=$1", id).Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u userRepository) UpdateUser(ctx context.Context, user *models.User) error {
	_, err := u.db.Exec(ctx, "UPDATE users SET username = $1, email = $2 WHERE id = $3", user.Username, user.Email, user.ID)
	return err
}

func (u userRepository) DeleteUser(ctx context.Context, id int) error {
	_, err := u.db.Exec(ctx, "DELETE FROM users WHERE id = $1", id)
	return err
}
