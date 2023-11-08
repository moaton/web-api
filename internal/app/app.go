package app

import (
	"fmt"
	"log"

	"github.com/moaton/web-api/config"
	db "github.com/moaton/web-api/internal/repository"
	"github.com/moaton/web-api/internal/services"
	"github.com/moaton/web-api/internal/transport/rest"
	"github.com/moaton/web-api/pkg/cache"
	"github.com/moaton/web-api/pkg/client/postgres"
)

func Run(cfg *config.Config) {
	url := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresHost, cfg.PostgresDB)
	client, err := postgres.NewClient(url)
	if err != nil {
		log.Fatalf("postgres.NewClient err %v", err)
	}
	repo := db.NewRepository(client)
	cache := cache.NewCache()

	service := services.NewService(repo)
	middleware := services.NewMiddleware()
	handler := rest.NewHandler(service, cache, middleware)

	go handler.ListenAndServe()
	log.Println("Running...")
}
