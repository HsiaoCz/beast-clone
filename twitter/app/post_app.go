package app

import (
	"encoding/json"
	"net/http"

	"github.com/HsiaoCz/beast-clone/twitter/app/middleware"
	"github.com/HsiaoCz/beast-clone/twitter/db"
	"github.com/HsiaoCz/beast-clone/twitter/types"
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
		return NewErrorResp(http.StatusNonAuthoritativeInfo, "user need login")
	}
	// TODO if user not login
	// need redirect to userlogin page
	if err := json.NewDecoder(r.Body).Decode(&postCreateParam); err != nil {
		return NewErrorResp(http.StatusBadRequest, err.Error())
	}

	post := types.NewPostFromParams(postCreateParam)
	post.UserID = userInfo.UserID

	postResp, err := p.db.Pc.CreatePost(r.Context(), post)
	if err != nil {
		return NewErrorResp(http.StatusInternalServerError, err.Error())
	}

	return WriteJson(w, http.StatusOK, postResp)
}
