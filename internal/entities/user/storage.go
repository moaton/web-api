package user

import "context"

type Storage interface {
	GetUserByEmail(ctx context.Context, email string) (User, error)
	InsertUser(ctx context.Context, user User) error
	UpdateUser(ctx context.Context, user User) error
	DeleteUser(ctx context.Context, email string) error
}
