package types

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	CommentID string
	UserID    string
	PostID    string
	ParentID  string
	Content   string
}
