package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/HsiaoCz/beast-clone/wechat/handlers"
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
	userHandlers := &handlers.UserHandlers{}

	router := http.NewServeMux()
	router.HandleFunc("GET /hello", handlers.TransferHandlefunc(userHandlers.HandleUserSignup))
	router.HandleFunc("GET /user", handlers.TransferHandlefunc(userHandlers.HandleUserLogin))

	slog.Info("the http server is runing", "listen address", os.Getenv("PORT"))
	http.ListenAndServe(os.Getenv("PORT"), router)
}
