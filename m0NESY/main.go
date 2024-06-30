package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello My Man"))
	})
	slog.Info("http server is running", "listen addresss", os.Getenv("PORT"))
	http.ListenAndServe(os.Getenv("PORT"), nil)
}
