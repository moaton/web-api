package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/moaton/web-api/config"
	"github.com/moaton/web-api/internal/app"
	"github.com/moaton/web-api/pkg/logger"
)

func init() {

}

func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	cfg := config.GetConfig()

	logger.SetLogger(cfg.IsDebug)

	wg.Add(1)
	go app.Run(ctx, &wg, cfg)

	sig := make(chan os.Signal, 1)
	done := make(chan int, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		<-sig
		cancel()
		wg.Wait()
		done <- 1
	}()

	<-done
	logger.Info("Graceful shutdown")
}
