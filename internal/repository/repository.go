package db

import (
	"database/sql"
	"sync"

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

func (r *Repository) Close() {
	var wg sync.WaitGroup
	wg.Add(2)
	r.Revenue.Close(&wg)
	r.User.Close(&wg)
	wg.Wait()
}
