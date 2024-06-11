package handlers

import "github.com/HsiaoCz/beast-clone/hotel/storage"

type BookingHandlers struct {
	store *storage.Store
}

func NewBookingHandlers(store *storage.Store) *BookingHandlers {
	return &BookingHandlers{
		store: store,
	}
}
