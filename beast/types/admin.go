package types

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	UserID       string `json:"user_id"`
	Username     string `json:"username"`
	UserPassword string `json:"-"`
	Avatar       string `json:"avatar"`
	Email        string `json:"email"`
}
