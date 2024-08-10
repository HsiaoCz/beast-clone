package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

var config = fiber.Config{
	AppName: "music",
	// ErrorHandler: func(c *fiber.Ctx, err error) error {
	// 	if
	// },
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	var (
		port = os.Getenv("PORT")

		r = fiber.New(config)
	)

	go func() {
		if err := r.Listen(port); err != nil {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	if err := r.Shutdown(); err != nil {
		log.Fatal(err)
	}
}
