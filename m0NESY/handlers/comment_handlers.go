package handlers

import (
	"net/http"

	"github.com/HsiaoCz/beast-clone/m0NESY/dao"
)

type CommentHandlers struct {
	comment dao.CommentDataModer
}

func NewCommentHandlers(comment dao.CommentDataModer) *CommentHandlers {
	return &CommentHandlers{
		comment: comment,
	}
}

func (c *CommentHandlers) HandleCreateComment(w http.ResponseWriter, r *http.Request) error {
	return nil
}
