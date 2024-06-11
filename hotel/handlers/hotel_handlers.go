package handlers

import "github.com/HsiaoCz/beast-clone/hotel/storage"

type HotelHandlers struct {
	store *storage.Store
}

func NewHotelHandlers(store *storage.Store) *HotelHandlers {
	return &HotelHandlers{
		store: store,
	}
}

