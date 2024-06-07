package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	var (
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))
		port   = ":3003"
	)
	slog.SetDefault(logger)
	router := chi.NewMux()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.RequestID)

	router.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})
	slog.Info("the server is running", "listen address", port)
	http.ListenAndServe(port, router)
}
