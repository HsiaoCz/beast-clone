package handlers

import "net/http"

type PriceHandler struct{}

func (p *PriceHandler) HandleGetPrice(w http.ResponseWriter, r *http.Request) error {
	return ErrorMessage(http.StatusInternalServerError, "something wrong")
}
