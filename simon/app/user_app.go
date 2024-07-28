package app

import "net/http"

type UserApp struct{}

func UserAppInit() *UserApp {
	return &UserApp{}
}

func (u *UserApp) HandleCreateUser(w http.ResponseWriter, r *http.Request) error {
	return nil
}
