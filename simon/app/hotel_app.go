package app

import "github.com/HsiaoCz/beast-clone/simon/store"

type HotelApp struct {
	store *store.Store
}

func HotelAppInit(store *store.Store) *HotelApp {
	return &HotelApp{
		store: store,
	}
}
