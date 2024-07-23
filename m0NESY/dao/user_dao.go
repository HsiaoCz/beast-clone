package dao

import (
	"github.com/HsiaoCz/beast-clone/m0NESY/types"
	"gorm.io/gorm"
)

type UserDataModer interface {
	CreateUser(*types.User) (*types.User, error)
	GetUserByID(string) (*types.User, error)
	DeleteUserByID(string) error
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

func (u *UserDataMod) GetUserByID(user_id string) (*types.User, error) {
	var user types.User
	tx := u.db.Model(&types.User{}).Find(&user, "user_id = ?", user_id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &user, nil
}

func (u *UserDataMod) DeleteUserByID(user_id string) error {
	return u.db.Where("user_id = ?", user_id).Delete(&types.User{}).Error
}
