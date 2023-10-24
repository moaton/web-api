package services

import "github.com/moaton/web-api/internal/adapters/db"

type Service struct {
	UserService    UserService
	RevenueService RevenueService
}

func NewService(db *db.Repository) *Service {
	return &Service{
		UserService:    newUserService(db),
		RevenueService: newRevenueService(db),
	}
}
