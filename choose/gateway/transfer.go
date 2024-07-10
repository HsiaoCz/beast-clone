package gateway

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func TransferHTTPHandler(fn fiber.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(&fiber.Ctx{}); err != nil {
			slog.Error("error", "err", err)
			WriteJSON(w, http.StatusInternalServerError, &RespError{
				Status: http.StatusInternalServerError,
				Error:  err.Error(),
			})
		}
	}
}

type RespError struct {
	Status int    `json:"status"`
	Error  string `json:"error"`
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}
