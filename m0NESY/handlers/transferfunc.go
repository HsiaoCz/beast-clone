package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/HsiaoCz/beast-clone/m0NESY/logger"
)

var StatusCode = &Status{Code: http.StatusOK}

type Status struct {
	Code int
}

type Handlerfunc func(w http.ResponseWriter, r *http.Request) error

type Map map[string]any

func TransferHandlerfunc(h Handlerfunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			defer logger.Logger.Error("the http server error", "method", r.Method, "path", r.URL.Path, "remote address", r.RemoteAddr, "error message", err)
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
		logger.Logger.Info("new request coming", "method", r.Method, "code", StatusCode.Code, "path", r.URL.Path, "remote address", r.RemoteAddr)
	}
}

func WriteJSON(w http.ResponseWriter, code int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	StatusCode.Code = code
	return json.NewEncoder(w).Encode(v)
}
