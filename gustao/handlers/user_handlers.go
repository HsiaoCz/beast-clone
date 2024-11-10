package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/HsiaoCz/beast-clone/gustao/data"
	"github.com/HsiaoCz/beast-clone/gustao/types"
)

type UserHandlers struct {
	data *data.UserData
}

func UserHandlersInit(data *data.UserData) *UserHandlers {
	return &UserHandlers{
		data: data,
	}
}

func (u *UserHandlers) HandleCreateUser(w http.ResponseWriter, r *http.Request) error {
	// 1. get create user params
	var create_user_params types.CreateUserReq
	if err := json.NewDecoder(r.Body).Decode(&create_user_params); err != nil {
		return ErrorMessage(http.StatusBadRequest, "please check the params")
	}
	// get user
	user := types.CreateUserFromParams(create_user_params)

	// in database
	userReturn, err := u.data.CreateUser(r.Context(), user)
	if err != nil {
		return ErrorMessage(http.StatusInternalServerError, err.Error())
	}
	return WriteJson(w, http.StatusOK, userReturn)
}
