package app

import (
	"net/http"

	"github.com/HsiaoCz/beast-clone/twitter/db"
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
	return nil
}
