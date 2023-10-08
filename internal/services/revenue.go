package services

import (
	"log"

	"github.com/moaton/web-api/internal/adapters/db"
	"github.com/moaton/web-api/internal/dto"
	"github.com/moaton/web-api/internal/models"
	"golang.org/x/net/context"
)

type RevenueService interface {
	GetRevenues(ctx context.Context, limit, offset int64) []models.Revenue
	GetRevenueById(ctx context.Context, id int64) models.Revenue
	CreateRevenue(ctx context.Context, dto dto.CreateRevenueDTO) error
	UpdateRevenue(ctx context.Context, dto dto.UpdateRevenueDTO) error
	DeleteRevenue(ctx context.Context, id int64) error
}

type revenueService struct {
	db *db.Repository
}

func NewRevenueService(db *db.Repository) RevenueService {
	return &revenueService{
		db: db,
	}
}

func (s *revenueService) GetRevenues(ctx context.Context, limit, offset int64) []models.Revenue {
	revenues, err := s.db.Revenue.GetRevenues(ctx, limit, offset)
	if err != nil {
		log.Println("GetRevenues err ", err)
		return []models.Revenue{}
	}
	return revenues
}

func (s *revenueService) GetRevenueById(ctx context.Context, id int64) models.Revenue {
	revenue, err := s.db.Revenue.GetRevenueById(ctx, id)
	if err != nil {
		log.Println("GetRevenueById err ", err)
		return models.Revenue{}
	}
	return revenue
}

func (s *revenueService) CreateRevenue(ctx context.Context, dto dto.CreateRevenueDTO) error {
	if err := s.db.Revenue.InsertRevenue(ctx, dto); err != nil {
		log.Println("InsertRevenue err ", err)
		return err
	}
	return nil
}

func (s *revenueService) UpdateRevenue(ctx context.Context, dto dto.UpdateRevenueDTO) error {
	if err := s.db.Revenue.UpdateRevenue(ctx, dto); err != nil {
		log.Println("UpdateRevenue err ", err)
		return err
	}
	return nil
}

func (s *revenueService) DeleteRevenue(ctx context.Context, id int64) error {
	if err := s.db.Revenue.DeleteRevenue(ctx, id); err != nil {
		log.Println("DeleteRevenue err ", err)
		return err
	}
	return nil
}
