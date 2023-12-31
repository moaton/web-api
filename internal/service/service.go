package service

import (
	db "github.com/moaton/web-api/internal/repository"
	"github.com/moaton/web-api/pkg/cache"
)

type Service struct {
	UserService    UserService
	RevenueService RevenueService
}

func New(db *db.Repository, cache *cache.Cache) *Service {
	return &Service{
		UserService:    newUserService(db, cache),
		RevenueService: newRevenueService(db, cache),
	}
}
