package main

import (
	"context"
	"log"
	"os"

	"github.com/HsiaoCz/beast-clone/clean/sweetwater/internal/server"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGOURL")))
	if err != nil {
		log.Fatal(err)
	}
	go func(context.Context) {
		if err := client.Ping(ctx, &readpref.ReadPref{}); err != nil {
			log.Fatal(err)
		}
	}(ctx)

	var (
		port   = os.Getenv("PORT")
		app    = fiber.New()
		server = server.HttpServerInit(app, client)
	)
	if err := server.Register(port); err != nil {
		log.Fatal(err)
	}
}
