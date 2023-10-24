package services

import (
	"context"

	"github.com/moaton/web-api/internal/repository"
	"github.com/moaton/web-api/internal/models"
)

type UserService interface {
	GetUserByEmail(ctx context.Context, email string) (models.User, error)
	CreateUser(ctx context.Context, User models.User) error
	UpdateUser(ctx context.Context, User models.User) error
	DeleteUser(ctx context.Context, email string) error
}

type userService struct {
	db *db.Repository
}

func newUserService(db *db.Repository) UserService {
	return &userService{
		db: db,
	}
}

func (s *userService) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	user, err := s.db.User.GetUserByEmail(ctx, email)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (s *userService) CreateUser(ctx context.Context, User models.User) error {
	if err := s.db.User.InsertUser(ctx, User); err != nil {
		return err
	}
	return nil
}

func (s *userService) UpdateUser(ctx context.Context, User models.User) error {
	if err := s.db.User.UpdateUser(ctx, User); err != nil {
		return err
	}
	return nil
}

func (s *userService) DeleteUser(ctx context.Context, email string) error {
	if err := s.db.User.DeleteUser(ctx, email); err != nil {
		return err
	}
	return nil
}
