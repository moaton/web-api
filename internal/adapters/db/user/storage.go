package user

import (
	"context"

	"github.com/moaton/web-api/internal/entities/user"
)

type UserStorage struct {
}

func (u *UserStorage) GetUserByEmail(ctx context.Context, email string) (user.User, error) {
	return user.User{}, nil
}

func (u *UserStorage) InsertUser(ctx context.Context, user user.User) error {
	return nil
}

func (u *UserStorage) UpdateUser(ctx context.Context, user user.User) error {
	return nil
}

func (u *UserStorage) DeleteUser(ctx context.Context, email string) error {
	return nil
}
