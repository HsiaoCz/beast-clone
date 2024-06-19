package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/HsiaoCz/beast-clone/twitter/db"
	"github.com/HsiaoCz/beast-clone/twitter/types"
	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (m *UserApp) HandleGetUserByID(w http.ResponseWriter, r *http.Request) error {
	id := chi.URLParam(r, "uid")
	uid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return NewErrorResp(http.StatusBadRequest, "invalid uid")
	}
	user, err := m.db.Uc.GetUserByID(r.Context(), uid)
	if err != nil {
		return NewErrorResp(http.StatusInternalServerError, err.Error())
	}
	return WriteJson(w, http.StatusOK, map[string]any{
		"status": http.StatusOK,
		"user":   user,
	})
}

func (m *UserApp) HandleDeleteUserByID(w http.ResponseWriter, r *http.Request) error {
	id := chi.URLParam(r, "uid")
	uid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return NewErrorResp(http.StatusBadRequest, "invalid uid")
	}
	if err := m.db.Uc.DeleteUserByID(r.Context(), uid); err != nil {
		return NewErrorResp(http.StatusInternalServerError, err.Error())
	}
	return WriteJson(w, http.StatusOK, map[string]any{
		"status":  http.StatusOK,
		"message": fmt.Sprintf("delete user (uid=%v) success", uid),
	})
}

func (m *UserApp) HandleUpdateUserByID(w http.ResponseWriter, r *http.Request) error {
	id := chi.URLParam(r, "uid")
	uid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return NewErrorResp(http.StatusBadRequest, "invalid uid")
	}
	updateUserParams := types.UpdateUserParams{}
	if err := json.NewDecoder(r.Body).Decode(&updateUserParams); err != nil {
		return NewErrorResp(http.StatusBadRequest, err.Error())
	}
	msg := updateUserParams.Validate()
	if len(msg) != 0 {
		return WriteJson(w, http.StatusBadRequest, msg)
	}
	user, err := m.db.Uc.UpdateUserByID(r.Context(), uid, &updateUserParams)
	if err != nil {
		return NewErrorResp(http.StatusInternalServerError, err.Error())
	}
	return WriteJson(w, http.StatusOK, user)
}
