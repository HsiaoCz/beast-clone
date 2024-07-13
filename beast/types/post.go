package types

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	UserID   string
	PostID   string
	Content  string
	PostPath string
}
