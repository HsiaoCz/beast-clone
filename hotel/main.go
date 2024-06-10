package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/HsiaoCz/beast-clone/hotel/conf"
	"github.com/HsiaoCz/beast-clone/hotel/handlers"
	"github.com/HsiaoCz/beast-clone/hotel/handlers/middlewares"
	"github.com/HsiaoCz/beast-clone/hotel/storage"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var config = fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		if apiError, ok := err.(handlers.APIError); ok {
			return c.Status(apiError.Status).JSON(&apiError)
		}
		aError := handlers.NewAPIError(http.StatusInternalServerError, err.Error())
		return c.Status(aError.Status).JSON(&aError)
	},
}

func main() {
	if err := conf.ParseConfig(); err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(conf.Conf.App.MongoUri))
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		if err := client.Ping(ctx, &readpref.ReadPref{}); err != nil {
			log.Fatal(err)
		}
	}()
	var (
		userColl       = client.Database(conf.Conf.App.DBName).Collection(conf.Conf.App.UserColl)
		mongoUserStore = storage.NewMongoUserStore(client, userColl)
		store          = &storage.Store{User: mongoUserStore}
		userHandlers   = handlers.NewUserHandlers(store)
		app            = fiber.New(config)
		v1             = app.Group("/api/v1")
	)
	{
		// router
		v1.Post("/user", userHandlers.HandleCreateUser)
		v1.Post("/user/login", userHandlers.HandleUserLogin)
		v1.Get("/user", middlewares.JWTAuthMiddleware(), userHandlers.HandleGetUserByID)
		v1.Delete("/user/:id", middlewares.JWTAuthMiddleware(), userHandlers.HandleDeleteUser)
		v1.Post("/user/:id", middlewares.JWTAuthMiddleware(), userHandlers.HandleUpdateUser)

		// router for booking
		// v1.Post("/")
	}

	go func() {
		if err := app.Listen(conf.Conf.App.Port); err != nil {
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
