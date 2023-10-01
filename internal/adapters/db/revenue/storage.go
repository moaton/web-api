package revenue

import (
	"context"

	"github.com/moaton/web-api/internal/entities/revenue"
)

type RevenueStorage struct {
}

func (r *RevenueStorage) GetRevenues(ctx context.Context, limit, offset int64) ([]revenue.Revenue, error) {
	return []revenue.Revenue{}, nil
}

func (r *RevenueStorage) GetRevenueById(ctx context.Context, id int64) (revenue.Revenue, error) {
	return revenue.Revenue{}, nil
}

func (r *RevenueStorage) InsertRevenue(ctx context.Context, dto revenue.CreateRevenueDTO) error {
	return nil
}

func (r *RevenueStorage) UpdateRevenue(ctx context.Context, dto revenue.UpdateRevenueDTO) error {
	return nil
}

func (r *RevenueStorage) DeleteRevenue(ctx context.Context, id int64) error {
	return nil
}
