package user

import (
	"context"
	"database/sql"
	"sync"

	_ "github.com/lib/pq"
	"github.com/moaton/web-api/internal/models"
)

type UserStorage interface {
	GetUserById(ctx context.Context, id int64) (models.User, error)
	GetUserByEmail(ctx context.Context, email string) (models.User, error)
	InsertUser(ctx context.Context, user models.User) (int64, error)
	UpdateUser(ctx context.Context, user models.User) error
	DeleteUser(ctx context.Context, email string) error

	Close(wg *sync.WaitGroup)
}

type storage struct {
	db *sql.DB
}

func NewUserStorage(db *sql.DB) UserStorage {
	return &storage{
		db: db,
	}
}

func (s *storage) Close(wg *sync.WaitGroup) {
	s.db.Close()
	wg.Done()
}

func (s *storage) GetUserById(ctx context.Context, id int64) (models.User, error) {
	user := models.User{}
	err := s.db.QueryRowContext(ctx, "SELECT id, email, name, password FROM users WHERE id = $1", id).Scan(&user.ID, &user.Email, &user.Name, &user.Password)
	return user, err
}

func (s *storage) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	user := models.User{}
	err := s.db.QueryRowContext(ctx, "SELECT id, email, name, password FROM users WHERE email = $1", email).Scan(&user.ID, &user.Email, &user.Name, &user.Password)
	return user, err
}

func (s *storage) InsertUser(ctx context.Context, user models.User) (int64, error) {
	var id int64
	err := s.db.QueryRowContext(ctx, "INSERT INTO users (email, name, password) VALUES ($1, $2, $3) RETURNING id", user.Email, user.Name, user.Password).Scan(&id)
	return id, err
}

func (s *storage) UpdateUser(ctx context.Context, user models.User) error {
	_, err := s.db.ExecContext(ctx, "UPDATE users SET email = $1, name = $2, password = $3 WHERE id = $4", user.Email, user.Name, user.Password, user.ID)
	return err
}

func (s *storage) DeleteUser(ctx context.Context, email string) error {
	_, err := s.db.ExecContext(ctx, "DELETE FROM users WHERE email = $1", email)
	return err
}
