package dao

import (
	"github.com/HsiaoCz/beast-clone/m0NESY/types"
	"gorm.io/gorm"
)

type CommentDataModer interface {
	CreateComment(*types.Comment) (*types.Comment, error)
}

type CommentDataMod struct {
	db *gorm.DB
}

func NewCommentDataMod(db *gorm.DB) *CommentDataMod {
	return &CommentDataMod{db: db}
}

func (c *CommentDataMod) CreateComment(comment *types.Comment) (*types.Comment, error) {
	tx := c.db.Model(&types.Comment{}).Create(comment)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return comment, nil
}
