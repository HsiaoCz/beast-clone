package st

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"regexp"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	bcryptCost     = 12
	minUsername    = 4
	minPasswordLen = 7
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username string             `bson:"username" json:"username"`
	Email    string             `bson:"email" json:"email"`
	Password string             `bson:"password" json:"-"`
	IsAdmin  bool               `bson:"isAdmin" json:"isAdmin"`
}

type CreateUserParam struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"isAdmin"`
}

func (params CreateUserParam) ValidateCreateUserParam() map[string]string {
	errors := map[string]string{}
	if len(params.Username) < minUsername {
		errors["username"] = fmt.Sprintf("username length should be at least %d characters", minUsername)
	}
	if len(params.Password) < minPasswordLen {
		errors["password"] = fmt.Sprintf("password length should be at least %d characters", minPasswordLen)
	}
	if !isEmailValidata(params.Email) {
		errors["email"] = fmt.Sprintf("email %s is invalid", params.Email)
	}
	return errors
}

func isEmailValidata(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}

func NewUserFromReq(params CreateUserParam) *User {
	encpw := encryptPassword(params.Password)
	return &User{
		Username: params.Username,
		Email:    params.Email,
		Password: encpw,
		IsAdmin:  params.IsAdmin,
	}
}

type UserLoginParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (p UserLoginParams) EncryptedPassword() *UserLoginParams {
	encpw := encryptPassword(p.Password)
	p.Password = encpw
	return &p
}

func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(os.Getenv("SECRET")))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

type UpdateUserParams struct {
	Username string `json:"username"`
}

func (p UpdateUserParams) ValidateUpdateUserParams() map[string]string {
	errors := map[string]string{}

	if len(p.Username) < minUsername {
		errors["username"] = fmt.Sprintf("username length should be at least %d characters", minUsername)
	}
	return errors
}

type VerifyUserPasswordParmas struct {
	Password       string `json:"password"`
	VerifyPassword string `json:"verify_password"`
}

func (v VerifyUserPasswordParmas) Validate() bool {
	return v.Password == v.VerifyPassword
}

func (v VerifyUserPasswordParmas) EncryptedUserPassword() string {
	return encryptPassword(v.Password)
}

// UserInfo context
type UserInfo struct {
	UserID  primitive.ObjectID
	Email   string
	IsAdmin bool
}
