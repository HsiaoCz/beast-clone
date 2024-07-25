package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/HsiaoCz/beast-clone/m0NESY/dao"
	"github.com/HsiaoCz/beast-clone/m0NESY/types"
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
	var create_comment_params types.CreateCommentParams
	if err := json.NewDecoder(r.Body).Decode(&create_comment_params); err != nil {
		return ErrorMessage(http.StatusBadRequest, err.Error())
	}
	comment, err := c.comment.CreateComment(types.NewCommentFromParams(create_comment_params))
	if err != nil {
		return ErrorMessage(http.StatusInternalServerError, err.Error())
	}
	return WriteJSON(w, http.StatusOK, comment)
}
