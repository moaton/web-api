package revenue

import "golang.org/x/net/context"

type Storage interface {
	GetRevenues(ctx context.Context, limit, offset int64) ([]Revenue, error)
	GetRevenueById(ctx context.Context, id int64) (Revenue, error)
	InsertRevenue(ctx context.Context, dto CreateRevenueDTO) error
	UpdateRevenue(ctx context.Context, dto UpdateRevenueDTO) error
	DeleteRevenue(ctx context.Context, id int64) error
}
