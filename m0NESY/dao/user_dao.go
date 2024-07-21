package dao

import (
	"github.com/HsiaoCz/beast-clone/m0NESY/types"
	"gorm.io/gorm"
)

type UserDataModer interface {
	CreateUser(*types.User) (*types.User, error)
}

type UserDataMod struct {
	db *gorm.DB
}

func NewUserDataMod(db *gorm.DB) *UserDataMod {
	return &UserDataMod{
		db: db,
	}
}

func (u *UserDataMod) CreateUser(user *types.User) (*types.User, error) {
	tx := u.db.Model(&types.User{}).Create(user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return user, nil
}
