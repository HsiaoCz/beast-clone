package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/HsiaoCz/beast-clone/m0NESY/dao"
	"github.com/HsiaoCz/beast-clone/m0NESY/types"
)

type UserHandlers struct {
	ud dao.UserDataModer
}

func NewUserHandlers(ud dao.UserDataModer) *UserHandlers {
	return &UserHandlers{
		ud: ud,
	}
}

func (u *UserHandlers) HandleCreateUser(w http.ResponseWriter, r *http.Request) error {
	var create_user_params types.CreateUserParams
	if err := json.NewDecoder(r.Body).Decode(&create_user_params); err != nil {
		return ErrorMessage(http.StatusBadRequest, err.Error())
	}
	user, err := u.ud.CreateUser(types.NewUserFromParams(create_user_params))
	if err != nil {
		return ErrorMessage(http.StatusInternalServerError, err.Error())
	}
	return WriteJSON(w, http.StatusOK, Map{
		"status":  http.StatusOK,
		"message": "create user success",
		"user":    user,
	})
}

func (u *UserHandlers) HandleGetUserByID(w http.ResponseWriter, r *http.Request) error {
	user_id := r.URL.Query().Get("uid")
	user, err := u.ud.GetUserByID(user_id)
	if err != nil {
		return ErrorMessage(http.StatusInternalServerError, err.Error())
	}
	return WriteJSON(w, http.StatusOK, Map{
		"status":  http.StatusOK,
		"message": "get user success",
		"user":    user,
	})
}

func (u *UserHandlers) HandleDeleteUserByID(w http.ResponseWriter, r *http.Request) error {
	user_id := r.URL.Query().Get("uid")
	if err := u.ud.DeleteUserByID(user_id); err != nil {
		return ErrorMessage(http.StatusInternalServerError, err.Error())
	}
	return WriteJSON(w, http.StatusOK, Map{
		"status":  http.StatusOK,
		"message": fmt.Sprintf("this user (%s) has been deleted", user_id),
	})
}

func (u *UserHandlers) HandleUpdateUser(w http.ResponseWriter, r *http.Request) error {
	var user_update types.UserUpdate
	if err := json.NewDecoder(r.Body).Decode(&user_update); err != nil {
		return ErrorMessage(http.StatusBadRequest, err.Error())
	}
	user, err := u.ud.UpdateUser(r.PathValue("uid"), &user_update)
	if err != nil {
		return ErrorMessage(http.StatusInternalServerError, err.Error())
	}
	return WriteJSON(w, http.StatusOK, Map{
		"status":  http.StatusOK,
		"message": "update user success",
		"user":    user,
	})
}
