package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/moaton/web-api/config"
	"github.com/moaton/web-api/internal/app"
	"github.com/moaton/web-api/pkg/logger"
)

func init() {

}

func main() {
	cfg := config.GetConfig()

	logger.SetLogger(cfg.IsDebug)

	// revenueService := usecase.NewUseCase()
	go app.Run(cfg)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-sig
	log.Println("Gracfull shutdown")
}
