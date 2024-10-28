package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

var StatusCode = &Status{Code: http.StatusOK}

type ErrorMsg struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (e ErrorMsg) Error() string {
	return e.Message
}

func ErrorMessage(status int, message string) ErrorMsg {
	return ErrorMsg{
		Status:  status,
		Message: message,
	}
}

type Status struct {
	Code int
}

type Handlefunc func(w http.ResponseWriter, r *http.Request) error

func TransferHandlefunc(h Handlefunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			defer slog.Error("the http server error", "method", r.Method, "path", r.URL.Path, "remote address", r.RemoteAddr, "error message", err)
			if e, ok := err.(ErrorMsg); ok {
				StatusCode.Code = e.Status
				WriteJSON(w, e.Status, &e)
			} else {
				errMsg := ErrorMsg{
					Status:  http.StatusInternalServerError,
					Message: err.Error(),
				}
				StatusCode.Code = errMsg.Status
				WriteJSON(w, errMsg.Status, &errMsg)
			}
		}
		slog.Info("new request coming", "method", r.Method, "code", StatusCode.Code, "path", r.URL.Path, "remote address", r.RemoteAddr)
	}
}

func WriteJSON(w http.ResponseWriter, code int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	StatusCode.Code = code
	return json.NewEncoder(w).Encode(v)
}
