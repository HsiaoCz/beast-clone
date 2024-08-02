package app

import "github.com/HsiaoCz/beast-clone/simon/store"

type BookingApp struct {
	store *store.Store
}

func BookingAppInit(store *store.Store) *BookingApp {
	return &BookingApp{
		store: store,
	}
}
