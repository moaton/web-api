package postgres

import (
	"database/sql"
	"time"

	"github.com/moaton/web-api/pkg/logger"
	"github.com/tinrab/retry"
)

func NewClient(url string) (*sql.DB, error) {
	var db *sql.DB
	var err error
	retry.ForeverSleep(2*time.Second, func(i int) error {
		db, err = sql.Open("postgres", url)
		if err != nil {
			logger.Errorf("sql.Open err %v", err)
			return err
		}
		err = db.Ping()
		if err != nil {
			logger.Errorf("db.Ping err %v", err)
			return err
		}
		return nil
	})

	return db, nil
}
