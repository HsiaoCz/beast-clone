package types

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"regexp"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	bcryptCost      = 12
	minFirstNameLen = 2
	minLastNameLen  = 2
	minPasswordLen  = 7
)

type User struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName         string             `bson:"firstName" json:"firstName"`
	LastName          string             `bson:"lastName" json:"lastName"`
	Email             string             `bson:"email" json:"email"`
	EncryptedPassword string             `bson:"EncryptedPassword" json:"-"`
	IsAdmin           bool               `bson:"isAdmin" json:"isAdmin"`
}

type CreateUserParam struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	IsAdmin   bool   `json:"isAdmin"`
}

func (params CreateUserParam) ValidateCreateUserParam() map[string]string {
	errors := map[string]string{}

	if len(params.FirstName) < minFirstNameLen {
		errors["firstName"] = fmt.Sprintf("firstName length should be at least %d characters", minFirstNameLen)
	}
	if len(params.LastName) < minLastNameLen {
		errors["lastName"] = fmt.Sprintf("lastName length should be at least %d characters", minLastNameLen)
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
		FirstName:         params.FirstName,
		LastName:          params.LastName,
		Email:             params.Email,
		EncryptedPassword: encpw,
		IsAdmin:           params.IsAdmin,
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
