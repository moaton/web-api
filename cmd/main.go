package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/moaton/web-api/internal/app"
)

func init() {

}

func main() {
	fmt.Println("Hello world")
	// revenueService := usecase.NewUseCase()
	go app.Run()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-sig
	log.Println("Gracfull shutdown")
}
