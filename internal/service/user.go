package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/moaton/web-api/internal/models"
	db "github.com/moaton/web-api/internal/repository"
	"github.com/moaton/web-api/pkg/cache"
	"github.com/moaton/web-api/pkg/logger"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Refresh(ctx context.Context, refreshToken string) (int64, map[string]string, error)
	Auth(ctx context.Context, email, password string) (int64, map[string]string, error)
	CreateUser(ctx context.Context, User models.User) (map[string]string, error)
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

func (s *userService) Refresh(ctx context.Context, refreshToken string) (int64, map[string]string, error) {
	email, err := Decode(refreshToken)
	if err != nil {
		return 0, map[string]string{}, fmt.Errorf("refresh token decode err %v", err)
	}
	user, err := s.db.User.GetUserByEmail(ctx, email)
	if err != nil {
		return 0, map[string]string{}, errors.New("user not found")
	}
	if _, err := s.cache.Get(user.ID); err != nil {
		return 0, map[string]string{}, errors.New("refresh token has been expired")
	}
	tokens, err := Encode(user.ID, user.Email)
	if err != nil {
		return 0, map[string]string{}, errors.New("tokens haven't been created")
	}
	return user.ID, tokens, nil
}

func (s *userService) Auth(ctx context.Context, email, password string) (int64, map[string]string, error) {
	user, err := s.db.User.GetUserByEmail(ctx, email)
	if err != nil {
		return 0, map[string]string{}, errors.New("user not found")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return 0, map[string]string{}, errors.New("invalid email or password")
	}

	tokens, err := Encode(user.ID, user.Email)
	if err != nil {
		return 0, map[string]string{}, errors.New("tokens haven't been created")
	}
	return user.ID, tokens, nil
}

func (s *userService) CreateUser(ctx context.Context, user models.User) (map[string]string, error) {

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.Errorf("CreateUser GenerateFromPassword err %v", err)
	}
	user.Password = string(password)

	id, err := s.db.User.InsertUser(ctx, user)
	if err != nil {
		return map[string]string{}, err
	}

	tokens, err := Encode(id, user.Email)
	if err != nil {
		logger.Errorf("CreateUser Encode err %v", err)
	}

	if err := s.cache.Set(id, tokens["refresh_token"], time.Now().Add(time.Hour*24).Unix()); err != nil {
		logger.Errorf("Refresh cache.Set err %v", err)
	}

	return tokens, nil
}

func (s *userService) UpdateUser(ctx context.Context, user models.User) error {
	err := s.db.User.UpdateUser(ctx, user)
	return err
}

func (s *userService) DeleteUser(ctx context.Context, email string) error {
	err := s.db.User.DeleteUser(ctx, email)
	return err
}
