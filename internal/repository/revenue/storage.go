package revenue

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/lib/pq"
	"github.com/moaton/web-api/internal/models"
	"github.com/moaton/web-api/pkg/logger"
)

type RevenueStorage interface {
	GetRevenues(ctx context.Context, limit, offset int64) ([]models.Revenue, int64, error)
	GetRevenueById(ctx context.Context, id int64) (models.Revenue, error)
	InsertRevenue(ctx context.Context, revenue models.Revenue) (int64, error)
	UpdateRevenue(ctx context.Context, revenue models.Revenue) error
	DeleteRevenue(ctx context.Context, id int64) error
}
type storage struct {
	db *sql.DB
}

func NewRevenueStorage(db *sql.DB) RevenueStorage {
	return &storage{
		db: db,
	}
}

func (s *storage) GetRevenues(ctx context.Context, limit, offset int64) ([]models.Revenue, int64, error) {
	var total int64
	err := s.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM revenues LIMIT $1 OFFSET $2", limit, offset).Scan(&total)
	if err != nil {
		logger.Errorf("GetRevenues total err %v", err)
		return []models.Revenue{}, 0, err
	}

	rows, err := s.db.QueryContext(ctx, "SELECT id, title, description, amount, type, createdat, updatedat FROM revenues LIMIT $1 OFFSET $2", limit, offset)
	if err != nil {
		logger.Errorf("GetRevenues SELECT err %v", err)
		return []models.Revenue{}, 0, err
	}
	revenues := []models.Revenue{}
	for rows.Next() {
		revenue := models.Revenue{}
		err := rows.Scan(&revenue.ID, &revenue.Title, &revenue.Description, &revenue.Amount, &revenue.Type, &revenue.CreatedAt, &revenue.UpdatedAt)
		if err != nil {
			logger.Errorf("GetRevenues rows.Scan err %v", err)
			continue
		}
		revenues = append(revenues, revenue)
	}
	return revenues, total, nil
}

func (s *storage) GetRevenueById(ctx context.Context, id int64) (models.Revenue, error) {
	var revenue models.Revenue
	err := s.db.QueryRowContext(ctx, "SELECT id, title, description, amount, type, createdat, updatedat FROM revenues WHERE id = $1", id).Scan(&revenue.ID, &revenue.Title, &revenue.Description, &revenue.Amount, &revenue.Type, &revenue.CreatedAt, &revenue.UpdatedAt)
	if err != nil {
		logger.Errorf("GetRevenueById total err %v", err)
		return models.Revenue{}, err
	}
	return revenue, nil
}

func (s *storage) InsertRevenue(ctx context.Context, revenue models.Revenue) (int64, error) {
	var id int64
	err := s.db.QueryRowContext(ctx, "INSERT INTO revenues (title, description, amount, type, createdat, updatedat) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id", revenue.Title, revenue.Description, revenue.Amount, revenue.Type, time.Now().UTC(), time.Now().UTC()).Scan(&id)
	return id, err
}

func (s *storage) UpdateRevenue(ctx context.Context, revenue models.Revenue) error {
	_, err := s.db.ExecContext(ctx, "UPDATE revenues SET title = $1, description = $2, amount = $3, type = $4, updatedat = $5 WHERE id = $6", revenue.Title, revenue.Description, revenue.Amount, revenue.Type, time.Now().UTC(), revenue.ID)
	return err
}

func (s *storage) DeleteRevenue(ctx context.Context, id int64) error {
	_, err := s.db.ExecContext(ctx, "DELETE FROM revenues WHERE id = $1", id)
	return err
}
