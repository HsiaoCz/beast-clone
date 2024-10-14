package types

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Age              int    `gorm:"column:age" json:"age"`
	User_ID          string `gorm:"column:user_id;" json:"user_id"`
	Username         string `gorm:"column:username;" json:"username"`
	Email            string `gorm:"column:email" json:"email"`
	Gender           string `gorm:"column:gender" json:"gender"`
	Birthday         string `gorm:"column:birthday" json:"birthday"`
	User_Password    string `gorm:"column:user_password" json:"-"`
	Synopsis         string `gorm:"column:synopsis" json:"synopsis"`
	Avatar           string `gorm:"column:avatar" json:"avatar"`
	Background_image string `gorm:"column:background_image" json:"background_image"`
}

type CreateUserReq struct {
	Username string `json:"username"`
	Gender   string `json:"gender"`
	Email    string `json:"email"`
	Birthday string `json:"birthday"`
	
}
