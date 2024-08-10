package handlers

import (
	"context"
	"encoding/json"
	"math"
	"math/rand"
	"net/http"
	"time"

	"github.com/HsiaoCz/beast-clone/m0NESY/types"
	"go.uber.org/zap"
)

var StatusCode = &Status{Code: http.StatusOK}

type Status struct {
	Code int
}

type Handlerfunc func(w http.ResponseWriter, r *http.Request) error

type Map map[string]any

func TransferHandlerfunc(h Handlerfunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestID := rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(math.MaxInt64)
		ctx := context.WithValue(r.Context(), types.CtxRequestIDKey, requestID)
		r = r.WithContext(ctx)
		if err := h(w, r); err != nil {
			defer zap.L().Error("the http server error", zap.Int64("requestID", requestID), zap.String("method", r.Method), zap.String("path", r.URL.Path), zap.String("remote address", r.RemoteAddr), zap.String("error message", err.Error()))
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
		zap.L().Info("new request coming", zap.Int64("requestID", requestID), zap.String("method", r.Method), zap.Int("code", StatusCode.Code), zap.String("path", r.URL.Path), zap.String("remote address", r.RemoteAddr))
	}
}

func WriteJSON(w http.ResponseWriter, code int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	StatusCode.Code = code
	return json.NewEncoder(w).Encode(v)
}
