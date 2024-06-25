package app

import (
	"encoding/json"
	"net/http"

	"github.com/HsiaoCz/beast-clone/twitter/app/middleware"
	"github.com/HsiaoCz/beast-clone/twitter/db"
	"github.com/HsiaoCz/beast-clone/twitter/types"
	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	userInfo, ok := r.Context().Value(middleware.CtxUserInfoKey).(types.UserInfo)
	if !ok {
		return ErrorMessage(http.StatusNonAuthoritativeInfo, "user need login")
	}
	var createCommentParam types.CreateCommentParams
	if err := json.NewDecoder(r.Body).Decode(&createCommentParam); err != nil {
		return ErrorMessage(http.StatusBadRequest, err.Error())
	}
	msg := createCommentParam.Validate()
	if len(msg) != 0 {
		return WriteJson(w, http.StatusBadRequest, msg)
	}
	comment := types.NewCommentsFromParams(createCommentParam, userInfo.UserID)

	commentResp, err := m.db.CS.CreateComment(r.Context(), comment)
	if err != nil {
		return ErrorMessage(http.StatusInternalServerError, err.Error())
	}
	return WriteJson(w, http.StatusOK, commentResp)
}

func (m *CommentApp) HandleDeleteCommentByID(w http.ResponseWriter, r *http.Request) error {
	userInfo, ok := r.Context().Value(middleware.CtxUserInfoKey).(types.UserInfo)
	if !ok {
		return ErrorMessage(http.StatusNonAuthoritativeInfo, "user need login")
	}

	id := chi.URLParam(r, "cid")
	cid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ErrorMessage(http.StatusBadRequest, err.Error())
	}
	if err := m.db.CS.DeleteCommentByID(r.Context(), userInfo.UserID, cid); err != nil {
		return ErrorMessage(http.StatusInternalServerError, err.Error())
	}
	return WriteJson(w, http.StatusOK, map[string]any{
		"status":  http.StatusOK,
		"message": "delete comment success",
	})
}

func (m *CommentApp) HandleGetCommentByPostID(w http.ResponseWriter, r *http.Request) error {
	id := r.PathValue("pid")
	pid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ErrorMessage(http.StatusBadRequest, err.Error())
	}
	comments, err := m.db.CS.GetCommentsByPostID(r.Context(), pid)
	if err != nil {
		return ErrorMessage(http.StatusInternalServerError, err.Error())
	}
	return WriteJson(w, http.StatusOK, comments)
}
