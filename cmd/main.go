package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/kelseyhightower/envconfig"
	"github.com/moaton/web-api/internal/app"
	"github.com/moaton/web-api/internal/models"
)

func init() {

}

func main() {
	var cfg models.Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatalf("envconfig.Process err %v", err)
	}
	// revenueService := usecase.NewUseCase()
	go app.Run(&cfg)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-sig
	log.Println("Gracfull shutdown")
}
