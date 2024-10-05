package main

import (
	"log"
	"net/http"
	"os"

	"github.com/HsiaoCz/beast-clone/underface/handlers"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
	var (
		port         = os.Getenv("PORT")
		userHandlers = &handlers.UserHandlers{}
		router       = http.NewServeMux()
	)

	{
		router.HandleFunc("POST /user", handlers.TransferHandlerfunc(userHandlers.HandleCreateUser))
	}

	logrus.WithFields(logrus.Fields{
		"listen address": port,
	}).Info("the http server is running")

	// listen and serve
	http.ListenAndServe(port, router)
}
