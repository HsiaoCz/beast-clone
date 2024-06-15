package handlers

import (
	"net/http"
)

type UserHandlers struct {
}

func (u *UserHandlers) HandleUserSignup(w http.ResponseWriter, r *http.Request) error {
	return NewAppError(http.StatusBadRequest, "what's up ASL")
}

func (u *UserHandlers) HandleUserLogin(w http.ResponseWriter, r *http.Request) error {
	return WriteJSON(w, http.StatusOK, map[string]any{
		"status":  http.StatusOK,
		"message": "all is well",
	})
}
