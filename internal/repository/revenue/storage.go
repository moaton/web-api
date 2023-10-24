package revenue

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/moaton/web-api/internal/models"
)

type RevenueStorage interface {
	GetRevenues(ctx context.Context, limit, offset int64) ([]models.Revenue, error)
	GetRevenueById(ctx context.Context, id int64) (models.Revenue, error)
	InsertRevenue(ctx context.Context, revenue models.Revenue) error
	UpdateRevenue(ctx context.Context, revenue models.Revenue) error
	DeleteRevenue(ctx context.Context, id int64) error
}
type storage struct {
	db *sql.DB
}

func NewRevenueStorage(db *sql.DB) *storage {
	return &storage{
		db: db,
	}
}

func (s *storage) GetRevenues(ctx context.Context, limit, offset int64) ([]models.Revenue, error) {
	return []models.Revenue{}, nil
}

func (s *storage) GetRevenueById(ctx context.Context, id int64) (models.Revenue, error) {
	return models.Revenue{}, nil
}

func (s *storage) InsertRevenue(ctx context.Context, revenue models.Revenue) error {
	return nil
}

func (s *storage) UpdateRevenue(ctx context.Context, revenue models.Revenue) error {
	return nil
}

func (s *storage) DeleteRevenue(ctx context.Context, id int64) error {
	return nil
}
