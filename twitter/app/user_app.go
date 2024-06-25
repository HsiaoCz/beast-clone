package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/HsiaoCz/beast-clone/twitter/app/middleware"
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
		return ErrorMessage(http.StatusBadRequest, "please chech the request params")
	}
	msg := userCreateParams.Validate()
	if len(msg) != 0 {
		return WriteJson(w, http.StatusBadRequest, msg)
	}
	user := types.NewUserFormParams(userCreateParams)
	userresp, err := u.db.Uc.CreateUser(r.Context(), user)
	if err != nil {
		return ErrorMessage(http.StatusInternalServerError, err.Error())
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
		return ErrorMessage(http.StatusBadRequest, "invalid uid")
	}
	user, err := m.db.Uc.GetUserByID(r.Context(), uid)
	if err != nil {
		return ErrorMessage(http.StatusInternalServerError, err.Error())
	}
	return WriteJson(w, http.StatusOK, map[string]any{
		"status": http.StatusOK,
		"user":   user,
	})
}

func (m *UserApp) HandleDeleteUserByID(w http.ResponseWriter, r *http.Request) error {
	// maybe delete user need user login
	// so the userID can get from the context
	// id := chi.URLParam(r, "uid")
	// uid, err := primitive.ObjectIDFromHex(id)
	// if err != nil {
	// 	return NewErrorResp(http.StatusBadRequest, "invalid uid")
	// }
	userInfo, ok := r.Context().Value(middleware.CtxUserInfoKey).(*types.UserInfo)
	if !ok {
		return ErrorMessage(http.StatusNonAuthoritativeInfo, "user need login")
	}
	if err := m.db.Uc.DeleteUserByID(r.Context(), userInfo.UserID); err != nil {
		return ErrorMessage(http.StatusInternalServerError, err.Error())
	}
	return WriteJson(w, http.StatusOK, map[string]any{
		"status":  http.StatusOK,
		"message": fmt.Sprintf("delete user (uid=%v) success", userInfo.UserID),
	})
}

func (m *UserApp) HandleUpdateUserByID(w http.ResponseWriter, r *http.Request) error {
	// update user need user login
	// so we can get the uid from the user context
	// id := chi.URLParam(r, "uid")
	// uid, err := primitive.ObjectIDFromHex(id)
	// if err != nil {
	// 	return NewErrorResp(http.StatusBadRequest, "invalid uid")
	// }
	userInfo, ok := r.Context().Value(middleware.CtxUserInfoKey).(*types.UserInfo)
	if !ok {
		return ErrorMessage(http.StatusNonAuthoritativeInfo, "user need login")
	}
	updateUserParams := types.UpdateUserParams{}
	if err := json.NewDecoder(r.Body).Decode(&updateUserParams); err != nil {
		return ErrorMessage(http.StatusBadRequest, err.Error())
	}
	msg := updateUserParams.Validate()
	if len(msg) != 0 {
		return WriteJson(w, http.StatusBadRequest, msg)
	}
	user, err := m.db.Uc.UpdateUserByID(r.Context(), userInfo.UserID, &updateUserParams)
	if err != nil {
		return ErrorMessage(http.StatusInternalServerError, err.Error())
	}
	return WriteJson(w, http.StatusOK, user)
}

func (u *UserApp) HandleFollowingUser(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (u *UserApp) HandleUnfollowingUser(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (u *UserApp) HandleBlackPerson(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (u *UserApp) HandleUnBlackPerson(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (u *UserApp) HandleGetFollowings(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (u *UserApp) HandleGetFollowers(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (u *UserApp) HandleGetBlacks(w http.ResponseWriter, r *http.Request) error {
	return nil
}
