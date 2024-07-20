package handlers

import (
	"context"
	"math"
	"math/rand"
	"net/http"

	"github.com/HsiaoCz/beast-clone/anne/service"
	"github.com/HsiaoCz/beast-clone/anne/types"
)

type PriceHandler struct {
	fetcher service.FetchPricer
}

func NewPriceHandler(fetcher service.FetchPricer) *PriceHandler {
	return &PriceHandler{
		fetcher: fetcher,
	}
}

func (p *PriceHandler) HandleGetPrice(w http.ResponseWriter, r *http.Request) error {
	ticker := r.URL.Query().Get("ticker")
	ctx := context.WithValue(r.Context(), types.CtxRequestID, rand.Intn(math.MaxInt))
	price, err := p.fetcher.FetchPrice(ctx, ticker)
	if err != nil {
		return ErrorMessage(http.StatusBadRequest, err.Error())
	}
	return WriteJSON(w, http.StatusOK, map[string]any{
		"status":  http.StatusOK,
		"message": "get price success",
		"price":   price,
	})
}
