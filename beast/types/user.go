package types

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserID           string
	Username         string
	Email            string
	UserPassword     string
	Synopsis         string
	Avatar           string
	Background_Image string
}
