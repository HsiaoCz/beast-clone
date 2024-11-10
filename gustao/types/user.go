package types

import (
	"crypto/md5"
	"encoding/hex"
	"os"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

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
	Password string `json:"password"`
	Age      int    `json:"age"`
}

func CreateUserFromParams(parmas CreateUserReq) *User {
	return &User{
		User_ID:          uuid.NewString(),
		Username:         parmas.Username,
		Email:            parmas.Email,
		Birthday:         parmas.Birthday,
		Age:              parmas.Age,
		Gender:           parmas.Gender,
		User_Password:    encryptPassword(parmas.Password),
		Synopsis:         "",
		Avatar:           "./picture/avatar/12334.jpg",
		Background_image: "./picture/brackground/12334.jpg",
	}
}

func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(os.Getenv("SECRET")))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
