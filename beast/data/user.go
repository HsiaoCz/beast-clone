package data

import (
	"github.com/HsiaoCz/beast-clone/beast/types"
	"gorm.io/gorm"
)

type UserStorer interface {
	CreateUser(*types.User) (*types.User, error)
	GetUserByID(string) (*types.User, error)
	GetUserByEmailAndPassword(string, string) error
	UpdateUser(string, *types.UserUpdate) (*types.User, error)
}

type UserStore struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) *UserStore {
	return &UserStore{
		db: db,
	}
}

func (u *UserStore) CreateUser(user *types.User) (*types.User, error) {
	tx := u.db.Model(&types.User{}).Create(user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return user, nil
}

func (u *UserStore) GetUserByID(user_id string) (*types.User, error) {
	var user types.User
	tx := u.db.Model(&types.User{}).Where("user_id = ?", user_id).First(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &user, nil
}

func (u *UserStore) GetUserByEmailAndPassword(email string, passwrod string) error {
	var user types.User
	tx := u.db.Model(&types.User{}).Find(&user, "email = ? AND user_password = ?", email, passwrod)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (u *UserStore) UpdateUser(user_id string, user_update *types.UserUpdate) (*types.User, error) {
	var user types.User
	tx := u.db.Model(&types.User{}).Where("user_id = ?", user_id).Updates(map[string]any{"username": user_update.Username, "avatar": user_update.Avatar, "synopsis": user_update.Synopsis, "background_image": user_update.Background_Image})
	if tx.Error != nil {
		return nil, tx.Error
	}
	tx1 := u.db.Model(&types.User{}).Find(&user, "user_id = ?", user_id)
	if tx1.Error != nil {
		return nil, tx.Error
	}
	return &user, nil
}
