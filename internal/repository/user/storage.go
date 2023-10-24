package user

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/moaton/web-api/internal/models"
)

type UserStorage interface {
	GetUserByEmail(ctx context.Context, email string) (models.User, error)
	InsertUser(ctx context.Context, user models.User) (int64, error)
	UpdateUser(ctx context.Context, user models.User) error
	DeleteUser(ctx context.Context, email string) error
}

type storage struct {
	db *sql.DB
}

func NewUserStorage(db *sql.DB) *storage {
	return &storage{
		db: db,
	}
}

func (s *storage) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	return models.User{}, nil
}

func (s *storage) InsertUser(ctx context.Context, user models.User) (int64, error) {
	var id int64
	err := s.db.QueryRowContext(ctx, "INSERT INTO user(email, name, password) VALUES ($1, %2, %3) RETURNING id", user.Email, user.Name, user.Password).Scan(&id)
	return id, err
}

func (s *storage) UpdateUser(ctx context.Context, user models.User) error {
	return nil
}

func (s *storage) DeleteUser(ctx context.Context, email string) error {
	return nil
}
