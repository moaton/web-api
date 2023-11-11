package service

import (
	"context"
	"errors"

	"github.com/moaton/web-api/internal/models"
	db "github.com/moaton/web-api/internal/repository"
	"github.com/moaton/web-api/pkg/cache"
	"github.com/moaton/web-api/pkg/logger"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	GetUserById(ctx context.Context, id int64) (models.User, error)
	Auth(ctx context.Context, email, password string) (models.User, error)
	CreateUser(ctx context.Context, User models.User) (models.User, error)
	UpdateUser(ctx context.Context, User models.User) error
	DeleteUser(ctx context.Context, email string) error
}

type userService struct {
	db    *db.Repository
	cache *cache.Cache
}

func newUserService(db *db.Repository, cache *cache.Cache) UserService {
	return &userService{
		db:    db,
		cache: cache,
	}
}

func (s *userService) GetUserById(ctx context.Context, id int64) (models.User, error) {
	return s.db.User.GetUserById(ctx, id)
}

func (s *userService) Auth(ctx context.Context, email, password string) (models.User, error) {
	user, err := s.db.User.GetUserByEmail(ctx, email)
	if err != nil {
		return models.User{}, errors.New("user not found")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return models.User{}, errors.New("invalid email or password")
	}

	return user, nil
}

func (s *userService) CreateUser(ctx context.Context, user models.User) (models.User, error) {

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.Errorf("CreateUser GenerateFromPassword err %v", err)
	}
	user.Password = string(password)

	id, err := s.db.User.InsertUser(ctx, user)
	if err != nil {
		return models.User{}, err
	}
	user.ID = id

	return user, nil
}

func (s *userService) UpdateUser(ctx context.Context, user models.User) error {
	err := s.db.User.UpdateUser(ctx, user)
	return err
}

func (s *userService) DeleteUser(ctx context.Context, email string) error {
	err := s.db.User.DeleteUser(ctx, email)
	return err
}
