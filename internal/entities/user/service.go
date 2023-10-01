package user

import (
	"context"
)

type Service interface {
	GetUserByEmail(ctx context.Context, email string) (User, error)
	CreateUser(ctx context.Context, User User) error
	UpdateUser(ctx context.Context, User User) error
	DeleteUser(ctx context.Context, email string) error
}

type service struct {
	storage Storage
}

func NewService(storage Storage) Service {
	return &service{
		storage: storage,
	}
}

func (s *service) GetUserByEmail(ctx context.Context, email string) (User, error) {
	user, err := s.storage.GetUserByEmail(ctx, email)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (s *service) CreateUser(ctx context.Context, User User) error {
	if err := s.storage.InsertUser(ctx, User); err != nil {
		return err
	}
	return nil
}

func (s *service) UpdateUser(ctx context.Context, User User) error {
	if err := s.storage.UpdateUser(ctx, User); err != nil {
		return err
	}
	return nil
}

func (s *service) DeleteUser(ctx context.Context, email string) error {
	if err := s.storage.DeleteUser(ctx, email); err != nil {
		return err
	}
	return nil
}
