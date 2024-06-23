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
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))
	slog.SetDefault(logger)
	userHandlers := &UserHandlers{}
	router := http.NewServeMux()

	router.HandleFunc("POST /user", TransferHandlefunc(userHandlers.HandleCreateUser))

	slog.Info("the server is running", "listen address", os.Getenv("PORT"))
	http.ListenAndServe(os.Getenv("PORT"), router)
}
