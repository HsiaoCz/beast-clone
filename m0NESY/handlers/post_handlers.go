package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/HsiaoCz/beast-clone/m0NESY/dao"
	"github.com/HsiaoCz/beast-clone/m0NESY/types"
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
	var create_post_params types.CreatePostParams
	if err := json.NewDecoder(r.Body).Decode(&create_post_params); err != nil {
		return ErrorMessage(http.StatusBadRequest, err.Error())
	}
	post, err := p.post.CreatePost(types.NewPostFromParams(create_post_params))
	if err != nil {
		return ErrorMessage(http.StatusInternalServerError, err.Error())
	}
	return WriteJSON(w, http.StatusOK, Map{
		"status":  http.StatusOK,
		"message": "create post success",
		"post":    post,
	})
}

func (p *PostHandler) HandleGetPostByID(w http.ResponseWriter, r *http.Request) error {
	post_id := r.URL.Query().Get("pid")
	post, err := p.post.GetPostByID(post_id)
	if err != nil {
		return ErrorMessage(http.StatusInternalServerError, err.Error())
	}
	return WriteJSON(w, http.StatusOK, Map{
		"status": http.StatusOK,
		"post":   post,
	})
}

func (p *PostHandler) HandleGetPostsByUserID(w http.ResponseWriter, r *http.Request) error {
	posts, err := p.post.GetPostByUserID(r.URL.Query().Get("uid"))
	if err != nil {
		return ErrorMessage(http.StatusInternalServerError, err.Error())
	}
	return WriteJSON(w, http.StatusOK, Map{
		"status": http.StatusOK,
		"posts":  posts,
	})
}
