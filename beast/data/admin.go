package data

import (
	"github.com/HsiaoCz/beast-clone/beast/types"
	"gorm.io/gorm"
)

type AdminStorer interface {
	CreateAdmin(*types.Admin) (*types.Admin, error)
}

type AdminStore struct {
	db *gorm.DB
}

func NewAdminStore(db *gorm.DB) *AdminStore {
	return &AdminStore{
		db: db,
	}
}

func (a *AdminStore) CreateAdmin(admin *types.Admin) (*types.Admin, error) {
	return admin, nil
}

func (a *AdminStore) DeleteAdmin(user_id string) error {
	return a.db.Where("user_id = ?", user_id).Delete(&types.Admin{}).Error
}
