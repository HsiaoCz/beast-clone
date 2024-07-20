package main

import (
	"log"
	"net/http"
	"os"

	"github.com/HsiaoCz/beast-clone/anne/handlers"
	"github.com/HsiaoCz/beast-clone/anne/service"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	var (
		port                = os.Getenv("PORT")
		defaultPriceFetcher = &service.DefaultPriceFetcher{}
		priceHandler        = handlers.NewPriceHandler(defaultPriceFetcher)
		router              = http.NewServeMux()
	)

	router.HandleFunc("GET /price", handlers.TransferHandlerfunc(priceHandler.HandleGetPrice))
	logrus.WithFields(logrus.Fields{
		"listen address": port,
	}).Info("http server is running")
	http.ListenAndServe(port, router)
}
