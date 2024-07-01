package main

import (
	"log"
	"net/http"
	"os"

	"github.com/HsiaoCz/beast-clone/m0NESY/logger"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	if err := logger.InitLogger("./info.log", "./debug.log", "./error.log", "./warn.log"); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello My Man"))
	})
	logger.Logger.Info("http server is running", "listen addresss", os.Getenv("PORT"))
	http.ListenAndServe(os.Getenv("PORT"), nil)
}
