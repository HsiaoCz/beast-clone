package types

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserID          string `gorm:"column:user_id;" json:"user_id"`
	Username        string `gorm:"column:username;" json:"username"`
	UserPassword    string `gorm:"column:user_password;" json:"-"`
	Gender          string `gorm:"column:gender;" json:"gender"`
	Age             string `gorm:"column:age;" json:"age"`
	Birthday        string `gorm:"column:birthday;" json:"birthday"`
	Avatar          string `gorm:"column:avatar;" json:"avatar"`
	BackgroundImage string `gorm:"column:background_image;" json:"background_image"`
}

type CreateUserParam struct {
	Username string `json:"username"`
}
