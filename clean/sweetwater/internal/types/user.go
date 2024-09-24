package types

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"regexp"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username string             `bson:"username" json:"username"`
	Password string             `bson:"password" json:"-"`
	Email    string             `bson:"email" json:"email"`
	Content  string             `bson:"content" json:"content"`
	Avatar   string             `bson:"avatar" json:"avatar"`
	ReadTime string             `bson:"readTime" json:"readTime"`
	IsAdmin  bool               `bson:"isAdmin" json:"isAdmin"`
}

// update user
type UserUpdateParams struct {
	Username string `json:"username"`
	Content  string `json:"content"`
	Avatar   string `json:"avatar"`
}

// user create params
type UserCreateParams struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"isAdmin"`
}

const (
	minUsername = 4
	minPassword = 8
)

func (param UserCreateParams) ValidateParams() map[string]string {
	errors := map[string]string{}
	if len(param.Username) < minUsername {
		errors["username"] = fmt.Sprintf("the username shouldent short the %d", minUsername)
	}
	if len(param.Password) < minPassword {
		errors["password"] = fmt.Sprintf("the password shouldent short the %d", minPassword)
	}
	if !isEmailvalid(param.Email) {
		errors["email"] = fmt.Sprintf("email %s is invaild", param.Email)
	}
	return errors
}

func isEmailvalid(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+.[a-z]{2,4}$`)
	return emailRegex.MatchString(email)
}

func NewUserFromParams(params UserCreateParams) *User {
	return &User{
		Username: params.Username,
		Password: encrptyedPassword(params.Password),
		Email:    params.Email,
		IsAdmin:  params.IsAdmin,
		ReadTime: "0",
		Avatar:   "./users/data/avatar/1111.jpg",
		Content:  "",
	}
}

func encrptyedPassword(oldPassword string) string {
	h := md5.New()
	h.Write([]byte(os.Getenv("SECERT")))
	return hex.EncodeToString(h.Sum([]byte(oldPassword)))
}

type UserLoginParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewUserLoginParams(userLoginParams UserLoginParams) UserLoginParams {
	return UserLoginParams{
		Email:    userLoginParams.Email,
		Password: encrptyedPassword(userLoginParams.Password),
	}
}

// UserInfo context
type UserInfo struct {
	UserID  primitive.ObjectID
	Email   string
	IsAdmin bool
}
