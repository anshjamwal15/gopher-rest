package server

import (
	"log"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
)

func StartServerWithGracefulShutdown(a *fiber.App) {

	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		if err := a.Shutdown(); err != nil {
			log.Printf("Server is shutting down! Error: %v", err)
		}
		close(idleConnsClosed)
	}()

	err := a.Listen(os.Getenv("SERVER_URL"))

	if err != nil {
		log.Printf("Server is not running! Reason: %v", err)
	}
	<-idleConnsClosed
}

func StartServer(a *fiber.App) {
	if err := a.Listen(os.Getenv("SERVER_URL")); err != nil {
		log.Printf("Server is not running! Reason: %v", err)
	}
}
