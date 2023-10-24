package user

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/moaton/web-api/internal/models"
)

type UserStorage interface {
	GetUserByEmail(ctx context.Context, email string) (models.User, error)
	InsertUser(ctx context.Context, user models.User) error
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

func (s *storage) InsertUser(ctx context.Context, user models.User) error {
	var id int64
	_ = s.db.QueryRowContext(ctx, "INSERT INTO user(name,) VALUES ($1) RETURNING id").Scan(&id)
	return nil
}

func (s *storage) UpdateUser(ctx context.Context, user models.User) error {
	return nil
}

func (s *storage) DeleteUser(ctx context.Context, email string) error {
	return nil
}
