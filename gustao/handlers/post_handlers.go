package handlers

import "net/http"

type PostHandlers struct{}

func (p *PostHandlers) HandleCreatePost(w http.ResponseWriter, r *http.Request) error {
	return nil
}
