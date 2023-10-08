package db

import (
	"context"
	"database/sql"

	"github.com/moaton/web-api/internal/adapters/db/revenue"
	"github.com/moaton/web-api/internal/adapters/db/user"
	"github.com/moaton/web-api/internal/dto"
	"github.com/moaton/web-api/internal/models"
)

type RevenueStorage interface {
	GetRevenues(ctx context.Context, limit, offset int64) ([]models.Revenue, error)
	GetRevenueById(ctx context.Context, id int64) (models.Revenue, error)
	InsertRevenue(ctx context.Context, dto dto.CreateRevenueDTO) error
	UpdateRevenue(ctx context.Context, dto dto.UpdateRevenueDTO) error
	DeleteRevenue(ctx context.Context, id int64) error
}

type UserStorage interface {
	GetUserByEmail(ctx context.Context, email string) (models.User, error)
	InsertUser(ctx context.Context, user models.User) error
	UpdateUser(ctx context.Context, user models.User) error
	DeleteUser(ctx context.Context, email string) error
}

type Repository struct {
	Revenue RevenueStorage
	User    UserStorage
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Revenue: revenue.NewRevenueStorage(db),
		User:    user.NewUserStorage(db),
	}
}
