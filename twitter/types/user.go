package types

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"regexp"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName       string             `bson:"firstName" json:"firstName"`
	LastName        string             `bson:"lastName" json:"lastName"`
	NickName        string             `bson:"nickName" json:"nickName"`
	Email           string             `bson:"email" json:"email"`
	PhoneNumber     string             `bson:"phoneNumber" json:"phoneNumber"`
	Password        string             `bson:"password" json:"-"`
	Avatar          string             `bson:"avatar" json:"avatar"`
	BackgroundImage string             `bson:"backgroundImage" json:"backgroundImage"`
	Synopsis        string             `bson:"synopsis" json:"synopsis"`
	JoinedTime      string             `bson:"joinedTime" json:"joinedTime"`
	Following       string             `bson:"following" json:"following"` // use redis
	Followers       string             `bson:"follwoers" json:"followers"` // use redis
}

type CreateUserParams struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	NickName    string `json:"nickName"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"password"`
}

const (
	minFirstName   = 4
	minLastName    = 4
	minPassword    = 8
	minNickName    = 4
	lenPhoneNumber = 11
)

func (params CreateUserParams) Validate() map[string]string {
	errors := map[string]string{}
	if len(params.FirstName) < minFirstName {
		errors["firstName"] = fmt.Sprintf("the firstName shoulden't short then %d", minFirstName)
	}
	if len(params.LastName) < minLastName {
		errors["lastName"] = fmt.Sprintf("the lastName shoulden't short then %d", minLastName)
	}
	if len(params.Password) < minPassword {
		errors["password"] = fmt.Sprintf("the password shoulden't short then %d", minPassword)
	}
	if len(params.NickName) < minNickName {
		errors["nickName"] = fmt.Sprintf("the nickName shoulden't short the %d", minNickName)
	}
	if len(params.PhoneNumber) != lenPhoneNumber {
		errors["phoneNumber"] = fmt.Sprintf("the phoneNumber should equal %d", lenPhoneNumber)
	}
	if !isEmailValidata(params.Email) {
		errors["email"] = fmt.Sprintf("the email %s is invalid", params.Email)
	}
	return errors
}
func isEmailValidata(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}

func EncryptedPassword(oldPassword string) string {
	h := md5.New()
	h.Write([]byte(os.Getenv("SECRET")))
	return hex.EncodeToString(h.Sum([]byte(oldPassword)))
}

func NewUserFormParams(parmas CreateUserParams) *User {
	user := &User{
		FirstName:       parmas.FirstName,
		LastName:        parmas.LastName,
		NickName:        parmas.NickName,
		Password:        EncryptedPassword(parmas.Password),
		Email:           "",
		PhoneNumber:     "",
		Avatar:          "./static/user/avatar/1211.jpg",
		BackgroundImage: "./static/user/background/1234.jpg",
		Synopsis:        "",
		JoinedTime:      time.Now().Format("2006/01/02"),
		Following:       "0",
		Followers:       "0",
	}
	if parmas.Email != "" {
		user.Email = parmas.Email
	}
	if parmas.PhoneNumber != "" {
		user.PhoneNumber = parmas.PhoneNumber
	}
	return user
}
