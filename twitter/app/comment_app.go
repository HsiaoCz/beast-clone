package app

import (
	"net/http"

	"github.com/HsiaoCz/beast-clone/twitter/db"
)

type CommentApp struct {
	db *db.DBS
}

func NewCommentApp(db *db.DBS) *CommentApp {
	return &CommentApp{
		db: db,
	}
}

func (m *CommentApp) HandleCreateComment(w http.ResponseWriter, r *http.Request) error {
	return nil
}
