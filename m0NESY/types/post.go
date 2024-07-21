package types

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	UserID   string `json:"user_id"`
	PostID   string `json:"post_id"`
	Content  string `json:"content"`
	PostPath string `json:"post_path"`
}

type CreatePostParams struct {
	UserID   string `json:"user_id"`
	PostID   string `json:"post_id"`
	Content  string `json:"content"`
	PostPath string `json:"post_path"`
}

func NewPostFromParams(params CreatePostParams) *Post {
	return &Post{
		UserID:   params.UserID,
		PostID:   params.PostID,
		Content:  params.Content,
		PostPath: params.PostPath,
	}
}
