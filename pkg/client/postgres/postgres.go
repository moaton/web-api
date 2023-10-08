package postgres

import (
	"database/sql"
	"log"
)

func NewClient(url string) (*sql.DB, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Printf("sql.Open err %v", err)
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		log.Printf("db.Ping err %v", err)
		return nil, err
	}

	return db, nil
}
