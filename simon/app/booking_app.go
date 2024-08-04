package app

import (
	"net/http"

	"github.com/HsiaoCz/beast-clone/simon/st"
	"github.com/HsiaoCz/beast-clone/simon/store"
)

type BookingApp struct {
	store *store.Store
}

func BookingAppInit(store *store.Store) *BookingApp {
	return &BookingApp{
		store: store,
	}
}

func (b *BookingApp) HandleGetBookings(w http.ResponseWriter, r *http.Request) error {
	// get booking need login
	userInfo, ok := r.Context().Value(st.CtxUserInfoKey).(*st.UserInfo)
	if !ok {
		return ErrorMessage(http.StatusUnauthorized, "user unlogin")
	}
	_ = userInfo
	return nil
}
