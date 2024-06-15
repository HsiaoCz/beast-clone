package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type Status struct {
	Code int
}

type Handlefunc func(w http.ResponseWriter, r *http.Request) error

func TransferHandlefunc(h Handlefunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		status := &Status{}
		if err := h(w, r); err != nil {
			slog.Error("the http server error", "method", r.Method, "path", r.URL.Path, "remote address", r.RemoteAddr, "error message", err)
			if e, ok := err.(AppError); ok {
				status.Code = e.Status
				WriteJSON(w, e.Status, &e)
			} else {
				aErr := AppError{
					Status:  http.StatusInternalServerError,
					Message: err.Error(),
				}
				status.Code = e.Status
				WriteJSON(w, aErr.Status, &aErr)
			}
		}
		slog.Info("new request coming", "method", r.Method, "code", status.Code, "path", r.URL.Path, "remote address", r.RemoteAddr)
	}
}

func WriteJSON(w http.ResponseWriter, code int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(v)
}
