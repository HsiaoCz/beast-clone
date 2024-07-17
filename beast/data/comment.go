package data

import (
	"github.com/HsiaoCz/beast-clone/beast/types"
	"gorm.io/gorm"
)

type CommentStorer interface {
	CreateComment(*types.Comment) (*types.Comment, error)
}

type CommentStore struct {
	db *gorm.DB
}

func NewCommentStore(db *gorm.DB) *CommentStore {
	return &CommentStore{
		db: db,
	}
}

func (c *CommentStore) CreateComment(comment *types.Comment) (*types.Comment, error) {
	tx := c.db.Model(&types.Comment{}).Create(comment)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return comment, nil
}

func (c *CommentStore) DeleteComment(comment_id string) error {
	return c.db.Where("comment_id = ?", comment_id).Delete(&types.Comment{}).Error
}

func (c *CommentStore) GetCommentByPostID(post_id string) ([]*types.Comment, error) {
	return nil, nil
}
