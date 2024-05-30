package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/HsiaoCz/beast-clone/hotel/handlers"
	"github.com/HsiaoCz/beast-clone/hotel/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGOURL")))
	if err != nil {
		log.Fatal(err)
	}

	var (
		userColl       = client.Database(os.Getenv("DBNAME")).Collection(os.Getenv("USERCOLL"))
		mongoUserStore = storage.NewMongoUserStore(client, userColl)
		store          = &storage.Store{User: mongoUserStore}
		userHandlers   = handlers.NewUserHandlers(store)
		app            = fiber.New()
		v1             = app.Group("/api/v1")
	)
	{
		// router
		v1.Post("/user", userHandlers.HandleCreateUser)
	}

	go func() {
		if err := app.Listen(os.Getenv("PORT")); err != nil {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	if err := app.Shutdown(); err != nil {
		log.Fatal(err)
	}
}
