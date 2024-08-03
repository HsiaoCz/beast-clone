package app

import (
	"net/http"

	"github.com/HsiaoCz/beast-clone/simon/store"
)

type HotelApp struct {
	store *store.Store
}

func HotelAppInit(store *store.Store) *HotelApp {
	return &HotelApp{
		store: store,
	}
}

func (H *HotelApp) HandleGetHotels(w http.ResponseWriter, r *http.Request) error {
	return nil
}
