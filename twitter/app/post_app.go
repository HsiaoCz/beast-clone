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

type PostApp struct {
	db *db.DBS
}

func NewPostApp(db *db.DBS) *PostApp {
	return &PostApp{
		db: db,
	}
}

func (p *PostApp) HandleCreatePost(w http.ResponseWriter, r *http.Request) error {
	var (
		postCreateParam = types.CreatePostParams{}
	)
	userInfo, ok := r.Context().Value(middleware.CtxUserInfoKey).(*types.UserInfo)
	if !ok {
		return ErrorMessage(http.StatusNonAuthoritativeInfo, "user need login")
	}
	// TODO if user not login
	// need redirect to userlogin page
	if err := json.NewDecoder(r.Body).Decode(&postCreateParam); err != nil {
		return ErrorMessage(http.StatusBadRequest, err.Error())
	}

	post := types.NewPostFromParams(postCreateParam)
	post.UserID = userInfo.UserID

	postResp, err := p.db.Pc.CreatePost(r.Context(), post)
	if err != nil {
		return ErrorMessage(http.StatusInternalServerError, err.Error())
	}

	return WriteJson(w, http.StatusOK, postResp)
}

func (p *PostApp) HandleDeletePost(w http.ResponseWriter, r *http.Request) error {
	paramid := chi.URLParam(r, "pid")
	postID, err := primitive.ObjectIDFromHex(paramid)
	if err != nil {
		return ErrorMessage(http.StatusBadRequest, "invalid post id")
	}
	if err := p.db.Pc.DeletePostByID(r.Context(), postID); err != nil {
		return ErrorMessage(http.StatusInternalServerError, err.Error())
	}
	return WriteJson(w, http.StatusOK, map[string]any{
		"status":  http.StatusOK,
		"message": "delete post success",
	})
}
