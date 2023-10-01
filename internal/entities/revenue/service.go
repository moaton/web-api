package revenue

import (
	"log"

	"golang.org/x/net/context"
)

type Service interface {
	GetRevenues(ctx context.Context, limit, offset int64) []Revenue
	GetRevenueById(ctx context.Context, id int64) Revenue
	CreateRevenue(ctx context.Context, dto CreateRevenueDTO) error
	UpdateRevenue(ctx context.Context, dto UpdateRevenueDTO) error
	DeleteRevenue(ctx context.Context, id int64) error
}

type service struct {
	storage Storage
}

func NewService(storage Storage) Service {
	return &service{
		storage: storage,
	}
}

func (s *service) GetRevenues(ctx context.Context, limit, offset int64) []Revenue {
	revenues, err := s.storage.GetRevenues(ctx, limit, offset)
	if err != nil {
		log.Println("GetRevenues err ", err)
		return []Revenue{}
	}
	return revenues
}

func (s *service) GetRevenueById(ctx context.Context, id int64) Revenue {
	revenue, err := s.storage.GetRevenueById(ctx, id)
	if err != nil {
		log.Println("GetRevenueById err ", err)
		return Revenue{}
	}
	return revenue
}

func (s *service) CreateRevenue(ctx context.Context, dto CreateRevenueDTO) error {
	if err := s.storage.InsertRevenue(ctx, dto); err != nil {
		log.Println("InsertRevenue err ", err)
		return err
	}
	return nil
}

func (s *service) UpdateRevenue(ctx context.Context, dto UpdateRevenueDTO) error {
	if err := s.storage.UpdateRevenue(ctx, dto); err != nil {
		log.Println("UpdateRevenue err ", err)
		return err
	}
	return nil
}

func (s *service) DeleteRevenue(ctx context.Context, id int64) error {
	if err := s.storage.DeleteRevenue(ctx, id); err != nil {
		log.Println("DeleteRevenue err ", err)
		return err
	}
	return nil
}
