package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/HsiaoCz/beast-clone/gustao/db"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	// get env
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	// set logrus
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})

	// init sqlite
	if err := db.Init(); err != nil {
		log.Fatal(err)
	}

	// connect mongo db
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

	var (
		port   = os.Getenv("PORT")
		router = http.NewServeMux()
	)

	logrus.WithFields(logrus.Fields{
		"listen address": port,
	}).Info("the http server is running....")

	http.ListenAndServe(port, router)
}
