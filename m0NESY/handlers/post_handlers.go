package handlers

import (
	"net/http"

	"github.com/HsiaoCz/beast-clone/m0NESY/dao"
)

type PostHandler struct {
	post dao.PostDataModer
}

func NewPostHandler(post dao.PostDataModer) *PostHandler {
	return &PostHandler{
		post: post,
	}
}

func (p *PostHandler) HandleCreatePost(w http.ResponseWriter, r *http.Request) error {
	return nil
}
