package revenue

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/moaton/web-api/internal/dto"
	"github.com/moaton/web-api/internal/models"
)

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

func (s *storage) InsertRevenue(ctx context.Context, dto dto.CreateRevenueDTO) error {
	return nil
}

func (s *storage) UpdateRevenue(ctx context.Context, dto dto.UpdateRevenueDTO) error {
	return nil
}

func (s *storage) DeleteRevenue(ctx context.Context, id int64) error {
	return nil
}
