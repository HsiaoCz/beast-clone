package types

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"

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
	Email       string `json:"email,omitempty"`
	PhoneNumber string `json:"phoneNumber,omitempty"`
	Password    string `json:"password"`
}

const (
	minFirstName = 4
	minLastName  = 4
	minPassword  = 8
	minNickName  = 4
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
	return errors
}

func EncryptedPassword(oldPassword string) string {
	h := md5.New()
	h.Write([]byte(os.Getenv("SECRET")))
	return hex.EncodeToString(h.Sum([]byte(oldPassword)))
}
