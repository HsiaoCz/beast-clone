package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/HsiaoCz/beast-clone/simon/app"
	"github.com/HsiaoCz/beast-clone/simon/store"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGOURL")))
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		if err := client.Ping(ctx, &readpref.ReadPref{}); err != nil {
			log.Fatal(err)
		}
	}()

	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})

	var (
		userColl  = client.Database(os.Getenv("DBNAME")).Collection(os.Getenv("USERCOLL"))
		userStore = store.UserStoreInit(client, userColl)
		store     = &store.Store{Us: userStore}
		userApp   = app.UserAppInit(store)
		router    = http.NewServeMux()
	)

	{
		// router
		router.HandleFunc("POST /user/signup", app.TransferHandlerfunc(userApp.HandleCreateUser))
		router.HandleFunc("POST /user/login", app.TransferHandlerfunc(userApp.HandleUserLogin))
		router.HandleFunc("GET /user", app.TransferHandlerfunc(userApp.HandleGetUserByID))
	}

	logrus.WithFields(logrus.Fields{
		"listen address": os.Getenv("PORT"),
	}).Info("the http server is running")
	http.ListenAndServe(os.Getenv("PORT"), router)
}
