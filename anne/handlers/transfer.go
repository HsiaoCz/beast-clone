package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

var StatusCode = &Status{Code: http.StatusOK}

type Handlerfunc func(w http.ResponseWriter, r *http.Request) error

type Status struct {
	Code int
}

func WriteJSON(w http.ResponseWriter, code int, v any) error {
	w.Header().Set("Content-Type", "application")
	w.WriteHeader(code)
	StatusCode.Code = code
	return json.NewEncoder(w).Encode(v)
}

func TransferHandlerfunc(h Handlerfunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		if err := h(w, r); err != nil {
			defer func() {
				logrus.WithFields(logrus.Fields{
					"method":         r.Method,
					"path":           r.URL.Path,
					"remote address": r.RemoteAddr,
					"error message":  err,
				}).Error("the http server error")
			}()
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
		logrus.WithFields(logrus.Fields{
			"method":         r.Method,
			"code":           StatusCode.Code,
			"path":           r.URL.Path,
			"remote address": r.RemoteAddr,
			"cost":           time.Since(start),
		}).Info("new request coming")
	}
}

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
