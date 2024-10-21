package main

import (
	"context"
	"encoding/json"
	"net/http"

	"math/rand"

	"github.com/sirupsen/logrus"
)

type CtxType string

const CtxKey CtxType = "RequestID"

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
	return json.NewEncoder(w).Encode(v)
}

func TransferHandlerfunc(h Handlerfunc) http.HandlerFunc {
	ctx := context.Background()
	ctx = context.WithValue(ctx, CtxKey, rand.Intn(10000000))
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(ctx, w, r); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		}
	}
}
