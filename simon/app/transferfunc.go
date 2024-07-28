package app

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

var StatusCode = &Status{Code: http.StatusOK}

type Status struct {
	Code int
}

type Handlerfunc func(w http.ResponseWriter, r *http.Request) error

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
				WriteJson(w, e.Status, &e)
			} else {
				errMsg := ErrorMsg{
					Status:  http.StatusInternalServerError,
					Message: err.Error(),
				}
				StatusCode.Code = e.Status
				WriteJson(w, errMsg.Status, &errMsg)
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

func WriteJson(w http.ResponseWriter, code int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	StatusCode.Code = code
	return json.NewEncoder(w).Encode(v)
}
