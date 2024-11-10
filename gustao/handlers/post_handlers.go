package handlers

import (
	"net/http"

	"github.com/HsiaoCz/beast-clone/gustao/data"
)

type PostHandlers struct{
	Post *data.PostData
}

func (p *PostHandlers) HandleCreatePost(w http.ResponseWriter, r *http.Request) error {
	return nil
}
