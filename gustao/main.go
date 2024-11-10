package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/HsiaoCz/beast-clone/gustao/data"
	"github.com/HsiaoCz/beast-clone/gustao/db"
	"github.com/HsiaoCz/beast-clone/gustao/handlers"
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
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

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
		port         = os.Getenv("PORT")
		router       = http.NewServeMux()
		userData     = data.UserDataInit(db.Get())
		userHandlers = handlers.UserHandlersInit(userData)
	)

	{
		// user handlefunc
		router.HandleFunc("POST /api/v1/user", handlers.TransferHandlerfunc(userHandlers.HandleCreateUser))
	}

	server := http.Server{
		Addr:         port,
		Handler:      router,
		ReadTimeout:  time.Millisecond * 1500,
		WriteTimeout: time.Millisecond * 1500,
	}

	go func() {
		logrus.WithFields(logrus.Fields{
			"listen address": port,
		}).Info("the http server is running....")
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	<-interrupt

	log.Println("shutting down the server....")

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

	log.Println("server gracefully shut down....")
}
