package scripts

import (
	"github.com/HsiaoCz/beast-clone/beast/types"
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

func (u *UserFeed) CreateUser(user *types.User) (*types.User, error) {
	tx := u.db.Model(&types.User{}).Create(user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return user, nil
}
