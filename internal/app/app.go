package app

import (
	"fmt"
	"log"

	"github.com/moaton/web-api/config"
	"github.com/moaton/web-api/internal/middleware"
	db "github.com/moaton/web-api/internal/repository"
	"github.com/moaton/web-api/internal/service"
	"github.com/moaton/web-api/internal/token"
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
	cache := cache.New()

	service := service.New(repo, cache)
	token := token.New(cfg.Secret)
	middleware := middleware.New(cfg.Secret, token)
	handler := rest.New(service, cache, middleware)

	go handler.ListenAndServe()
	log.Println("Running...")
}
