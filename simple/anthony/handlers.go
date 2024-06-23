package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type UserHandlers struct{}

func (u *UserHandlers) HandleCreateUser(w http.ResponseWriter, r *http.Request) error {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		return ErrorMessage(http.StatusBadRequest, err.Error())
	}
	fmt.Println(user)

	return WriteJSON(w, http.StatusOK, &user)
}
