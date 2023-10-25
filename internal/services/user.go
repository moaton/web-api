package services

import (
	"context"

	"github.com/moaton/web-api/internal/models"
	db "github.com/moaton/web-api/internal/repository"
	"github.com/moaton/web-api/pkg/logger"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	GetUserByEmail(ctx context.Context, email, password string) (models.User, error)
	CreateUser(ctx context.Context, User models.User) (int64, error)
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

func (s *userService) GetUserByEmail(ctx context.Context, email, password string) (models.User, error) {
	user, err := s.db.User.GetUserByEmail(ctx, email)
	if err != nil {
		return models.User{}, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (s *userService) CreateUser(ctx context.Context, user models.User) (int64, error) {

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.Errorf("CreateUser GenerateFromPassword err %v", err)
	}
	user.Password = string(password)

	id, err := s.db.User.InsertUser(ctx, user)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *userService) UpdateUser(ctx context.Context, user models.User) error {
	err := s.db.User.UpdateUser(ctx, user)
	return err
}

func (s *userService) DeleteUser(ctx context.Context, email string) error {
	err := s.db.User.DeleteUser(ctx, email)
	return err
}
