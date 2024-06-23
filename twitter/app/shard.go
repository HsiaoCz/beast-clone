package app

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type Status struct {
	Code int
}

type Handlerfunc func(w http.ResponseWriter, r *http.Request) error

func TransferHandlerfunc(h Handlerfunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		status := &Status{Code: http.StatusOK}
		if err := h(w, r); err != nil {
			defer slog.Error("the http server error", "method", r.Method, "path", r.URL.Path, "remote address", r.RemoteAddr, "error message", err)
			if e, ok := err.(ErrorMsg); ok {
				status.Code = e.Status
				WriteJson(w, e.Status, &e)
			} else {
				errMsg := ErrorMsg{
					Status:  http.StatusInternalServerError,
					Message: err.Error(),
				}
				status.Code = e.Status
				WriteJson(w, errMsg.Status, &errMsg)
			}
		}
		slog.Info("new request coming", "method", r.Method, "code", status.Code, "path", r.URL.Path, "remote address", r.RemoteAddr)
	}
}

func WriteJson(w http.ResponseWriter, code int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(v)
}
