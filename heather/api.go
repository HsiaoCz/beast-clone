package main

import (
	"context"
	"encoding/json"
	"math/rand"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

var StatusCode = &Status{Code: http.StatusOK}

type Status struct {
	Code int
}
type CtxType string

const CtxKey CtxType = "RequsetID"

type Handlerfunc func(context.Context, http.ResponseWriter, *http.Request) error

type PriceResponse struct {
	Ticker string  `json:"ticker"`
	Price  float64 `json:"price"`
}

type JSONAPIServer struct {
	listenAddr string
	svc        PriceFetcher
}

func NewJSONAPIServer(listenAddr string, svc PriceFetcher) *JSONAPIServer {
	return &JSONAPIServer{
		listenAddr: listenAddr,
		svc:        svc,
	}
}

func (s *JSONAPIServer) Run() {
	router := http.NewServeMux()

	router.HandleFunc("GET /price", TransferHandlerfunc(s.handleFetchPrice))

	logrus.WithFields(logrus.Fields{
		"listen address": s.listenAddr,
	}).Info("the http server is running")
	http.ListenAndServe(s.listenAddr, router)
}

func (s *JSONAPIServer) handleFetchPrice(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	ticker := r.URL.Query().Get("ticker")

	price, err := s.svc.FetchPrice(ctx, ticker)
	if err != nil {
		return err
	}

	priceResp := PriceResponse{
		Ticker: ticker,
		Price:  price,
	}

	return writeJSON(w, http.StatusOK, &priceResp)
}

func writeJSON(w http.ResponseWriter, code int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	StatusCode.Code = code
	return json.NewEncoder(w).Encode(v)
}

func TransferHandlerfunc(h Handlerfunc) http.HandlerFunc {
	ctx := context.Background()
	ctx = context.WithValue(ctx, CtxKey, rand.Intn(10000000))
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		if err := h(ctx, w, r); err != nil {
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
				writeJSON(w, e.Status, &e)
			} else {
				errMsg := ErrorMsg{
					Status:  http.StatusInternalServerError,
					Message: err.Error(),
				}
				StatusCode.Code = e.Status
				writeJSON(w, errMsg.Status, &errMsg)
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
