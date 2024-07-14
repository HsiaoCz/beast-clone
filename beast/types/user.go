package types

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"regexp"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserID           string `json:"userID"`
	Username         string `json:"username"`
	Email            string `json:"email"`
	UserPassword     string `json:"-"`
	Synopsis         string `json:"synopsis"`
	Avatar           string `json:"avatar"`
	Background_Image string `json:"background_image"`
}

type CreateUserParams struct {
	Username         string `json:"username"`
	Email            string `json:"email"`
	Password         string `json:"password"`
	Synopsis         string `json:"synopsis"`
	Avatar           string `json:"avatar"`
	Background_Image string `json:"background_image"`
}

var (
	minUsernameLen = 4
	maxUsernameLen = 40
)

func (p CreateUserParams) Validate() map[string]string {
	errors := map[string]string{}
	if len(p.Username) < minUsernameLen || len(p.Username) > maxUsernameLen {
		errors["username"] = fmt.Sprintf("username should more than %d and less than %d", minUsernameLen, maxUsernameLen)
	}
	if !isValidEmail(p.Email) {
		errors["email"] = fmt.Sprintf("your email %s  is invalid", p.Email)
	}
	return errors
}

func isValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+[a-z]{2,4}$`)
	return emailRegex.MatchString(email)
}

func NewUserFromParams(params CreateUserParams) *User {
	return &User{
		UserID:           uuid.New().String(),
		Username:         params.Username,
		UserPassword:     encryptPassword(params.Password),
		Background_Image: params.Background_Image,
		Avatar:           params.Avatar,
		Email:            params.Email,
		Synopsis:         params.Synopsis,
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
