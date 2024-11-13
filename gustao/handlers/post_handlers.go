package handlers

import (
	"net/http"

	"github.com/HsiaoCz/beast-clone/gustao/data"
)

type PostHandlers struct {
	post data.PostDataInter
}

func PostHandlersInit(post data.PostDataInter) *PostHandlers {
	return &PostHandlers{
		post: post,
	}
}

func (p *PostHandlers) HandleCreatePost(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (p *PostHandlers) HandleGetPostByID(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (p *PostHandlers) HandleGetPostsByUserID(w http.ResponseWriter, r *http.Request) error {
	return nil
}
