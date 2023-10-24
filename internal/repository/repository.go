package db

import (
	"database/sql"

	"github.com/moaton/web-api/internal/repository/revenue"
	"github.com/moaton/web-api/internal/repository/user"
)

type Repository struct {
	Revenue revenue.RevenueStorage
	User    user.UserStorage
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Revenue: revenue.NewRevenueStorage(db),
		User:    user.NewUserStorage(db),
	}
}
