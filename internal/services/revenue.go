package services

import (
	"github.com/moaton/web-api/internal/models"
	db "github.com/moaton/web-api/internal/repository"
	"github.com/moaton/web-api/pkg/cache"
	"golang.org/x/net/context"
)

type RevenueService interface {
	GetRevenues(ctx context.Context, limit, offset int64) ([]models.Revenue, int64, error)
	GetRevenueById(ctx context.Context, id int64) (models.Revenue, error)
	CreateRevenue(ctx context.Context, revenue models.Revenue) (int64, error)
	UpdateRevenue(ctx context.Context, revenue models.Revenue) error
	DeleteRevenue(ctx context.Context, id int64) error
}

type revenueService struct {
	db    *db.Repository
	cache *cache.Cache
}

func newRevenueService(db *db.Repository, cache *cache.Cache) RevenueService {
	return &revenueService{
		db:    db,
		cache: cache,
	}
}

func (s *revenueService) GetRevenues(ctx context.Context, limit, offset int64) ([]models.Revenue, int64, error) {
	revenues, total, err := s.db.Revenue.GetRevenues(ctx, limit, offset)
	if err != nil {
		return []models.Revenue{}, 0, err
	}
	return revenues, total, nil
}

func (s *revenueService) GetRevenueById(ctx context.Context, id int64) (models.Revenue, error) {
	revenue, err := s.db.Revenue.GetRevenueById(ctx, id)
	return revenue, err
}

func (s *revenueService) CreateRevenue(ctx context.Context, revenue models.Revenue) (int64, error) {
	id, err := s.db.Revenue.InsertRevenue(ctx, revenue)
	return id, err
}

func (s *revenueService) UpdateRevenue(ctx context.Context, revenue models.Revenue) error {
	err := s.db.Revenue.UpdateRevenue(ctx, revenue)
	return err
}

func (s *revenueService) DeleteRevenue(ctx context.Context, id int64) error {
	err := s.db.Revenue.DeleteRevenue(ctx, id)
	return err
}
