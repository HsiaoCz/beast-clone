package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/HsiaoCz/beast-clone/underface/types"
	"github.com/sirupsen/logrus"
)

type UserHandlers struct{}

func (u *UserHandlers) HandleCreateUser(w http.ResponseWriter, r *http.Request) error {
	var user types.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		return ErrorMessage(http.StatusBadRequest, "check the input")
	}
	requestID, ok := r.Context().Value(types.CtxRequestIDKey).(int64)
	if !ok {
		logrus.Error("context fucked")
	}
	logrus.WithFields(logrus.Fields{
		"requestID": requestID,
	}).Info("the request fucking ID")
	return WriteJson(w, http.StatusOK, &user)
}

func (u *UserHandlers) HandleGetUserByID(w http.ResponseWriter, r *http.Request) error {
	return nil
}
