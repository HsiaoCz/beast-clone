package app

import (
	"net/http"

	"github.com/HsiaoCz/beast-clone/simon/store"
)

type AdminApp struct {
	store *store.Store
}

func AdminAppInit(store *store.Store) *AdminApp {
	return &AdminApp{
		store: store,
	}
}

func (a *AdminApp) HandleCreateHotel(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (a *AdminApp) HandleCreateRoom(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (a *AdminApp) HandleDeleteHotel(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (a *AdminApp) HandleDeleteRoom(w http.ResponseWriter, r *http.Request) error {
	return nil
}
