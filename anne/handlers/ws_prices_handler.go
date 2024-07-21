package handlers

import "net/http"

type WsPriceHandler struct {
}

func NewWsPriceHandler() *WsPriceHandler {
	return &WsPriceHandler{}
}

func (wsh *WsPriceHandler) HandlePriceFetch(w http.ResponseWriter, r *http.Request) error {
	return nil
}
