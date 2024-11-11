package data

import (
	"context"

	"github.com/HsiaoCz/beast-clone/gustao/types"
	"gorm.io/gorm"
)

type UserDataInter interface {
	CreateUser(context.Context, *types.User) (*types.User, error)
	UpdateUser(context.Context, *types.User) (*types.User, error)
}

type UserData struct {
	db *gorm.DB
}

func UserDataInit(db *gorm.DB) *UserData {
	return &UserData{
		db: db,
	}
}

func (u *UserData) CreateUser(ctx context.Context, user *types.User) (*types.User, error) {
	tx := u.db.Debug().WithContext(ctx).Model(&types.User{}).Create(user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return user, nil
}

func (u *UserData)UpdateUser(ctx context.Context,user *types.User)(*types.User,error){
	return nil,nil
}