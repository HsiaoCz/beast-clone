package db

import "gorm.io/gorm"

type UserStorer interface {
	CreateUser()
}

type MysqlUserStore struct {
	db *gorm.DB
}

func NewMysqlUserStore(db *gorm.DB) *MysqlUserStore {
	return &MysqlUserStore{
		db: db,
	}
}

func (m *MysqlUserStore) CreateUser() {}
