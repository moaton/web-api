package app

import (
	"fmt"
	"log"

	"github.com/moaton/web-api/internal/adapters/db"
	"github.com/moaton/web-api/internal/models"
	"github.com/moaton/web-api/pkg/client/postgres"
)

func Run(cfg *models.Config) {
	url := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresHost, cfg.PostgresDB)
	client, err := postgres.NewClient(url)
	if err != nil {
		log.Fatalf("postgres.NewClient err %v", err)
	}
	db.NewRepository(client)

	// err := http.ListenAndServe(":3000", handler)
	// if err != nil {
	// 	log.Println("ListenAndServe err ", err)
	// }
}
