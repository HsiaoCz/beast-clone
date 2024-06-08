package app

import (
	"encoding/json"
	"net/http"

	"github.com/HsiaoCz/beast-clone/twitter/db"
	"github.com/HsiaoCz/beast-clone/twitter/types"
)

type UserApp struct {
	db *db.DBS
}

func NewUserApp(db *db.DBS) *UserApp {
	return &UserApp{
		db: db,
	}
}

func (u *UserApp) HandleCreateUser(w http.ResponseWriter, r *http.Request) error {
	userCreateParams := types.CreateUserParams{}
	if err := json.NewDecoder(r.Body).Decode(&userCreateParams); err != nil {
		return NewErrorResp(http.StatusBadRequest, "please chech the request params")
	}
	msg := userCreateParams.Validate()
	if len(msg) != 0 {
		return WriteJson(w, http.StatusBadRequest, msg)
	}
	user := types.NewUserFormParams(userCreateParams)
	userresp, err := u.db.Uc.CreateUser(r.Context(), user)
	if err != nil {
		return NewErrorResp(http.StatusInternalServerError, err.Error())
	}
	return WriteJson(w, http.StatusOK, map[string]any{
		"status": http.StatusOK,
		"user":   userresp,
	})
}
