package handlers

import "net/http"

type UserHandler struct{}

func (u *UserHandler) HandleUserCreate(w http.ResponseWriter, r *http.Request) error {
	return nil
}
