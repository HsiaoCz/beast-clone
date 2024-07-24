package scripts

import (
	"github.com/HsiaoCz/beast-clone/hotel/types"
	"gorm.io/gorm"
)

type UserFeed struct {
	db *gorm.DB
}

func NewUserFeed(db *gorm.DB) *UserFeed {
	return &UserFeed{
		db: db,
	}
}

func (u *UserFeed)CreateUser(user *types.User)(*types.User,error){
	return nil,nil
}