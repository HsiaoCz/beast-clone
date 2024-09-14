package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/HsiaoCz/beast-clone/something/handlers"
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
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGOURL")))
	if err != nil {
		log.Fatal(err)
	}
	go func(ctx context.Context) {
		if err := client.Ping(ctx, &readpref.ReadPref{}); err != nil {
			log.Fatal(err)
		}
	}(ctx)

	var (
		port        = os.Getenv("PORT")
		userHandler = &handlers.UserHandlers{}
		router      = http.NewServeMux()
	)

	{
		router.HandleFunc("POST /user", handlers.TransferHandlerfunc(userHandler.HandleCreateUser))
	}

	http.ListenAndServe(port, router)
}
