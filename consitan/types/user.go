package types

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Username string
	UserID   string
	Email    string
}
