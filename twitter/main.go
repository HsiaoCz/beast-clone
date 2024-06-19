package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/HsiaoCz/beast-clone/twitter/app"
	"github.com/HsiaoCz/beast-clone/twitter/db"
	"github.com/HsiaoCz/beast-clone/twitter/etc"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	if err := etc.ParseConfig(); err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(etc.Conf.App.MongoUri))
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		if err := client.Ping(ctx, &readpref.ReadPref{}); err != nil {
			log.Fatal(err)
		}
	}()

	var (
		logger        = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))
		port          = etc.Conf.App.Port
		userColl      = client.Database(etc.Conf.App.DBName).Collection(etc.Conf.App.UserColl)
		mongoUserCase = db.NewMongoUserCase(client, userColl)
		dbs           = &db.DBS{Uc: mongoUserCase}
		userApp       = app.NewUserApp(dbs)
		router        = chi.NewMux()
	)

	slog.SetDefault(logger)

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.RequestID)

	router.Post("/user", app.TransferHandler(userApp.HandleCreateUser))
	router.Get("/user/{uid}", app.TransferHandler(userApp.HandleGetUserByID))
	router.Delete("/user/{uid}", app.TransferHandler(userApp.HandleDeleteUserByID))
	router.Delete("/user/{uid}", app.TransferHandler(userApp.HandleUpdateUserByID))

	slog.Info("the server is running", "listen address", port)
	http.ListenAndServe(port, router)
}
