package types

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	CommentID string `json:"comment_id"`
	UserID    string `json:"user_id"`
	PostID    string `json:"post_id"`
	ParentID  string `json:"parent_id"`
	Content   string `json:"content"`
}
type CreateCommentParams struct {
	CommentID string `json:"comment_id"`
	UserID    string `json:"user_id"`
	PostID    string `json:"post_id"`
	ParentID  string `json:"parent_id"`
	Content   string `json:"content"`
}

func NewCommentFromParams(params CreateCommentParams) *Comment {
	return &Comment{
		CommentID: params.CommentID,
		UserID:    params.UserID,
		PostID:    params.PostID,
		ParentID:  params.ParentID,
		Content:   params.CommentID,
	}
}
