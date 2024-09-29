package types

import "gorm.io/gorm"

type User struct {
	gorm.Model
	User_ID       string `gorm:"column:user_id;" json:"user_id"`
	Username      string `gorm:"column:username;" json:"username"`
	Email         string `gorm:"column:email" json:"email"`
	User_Password string `gorm:"column:user_password" json:"user_password"`
	
}
