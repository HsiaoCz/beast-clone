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
	tx := c.db.Model(&types.Comment{}).Delete("comment_id = ?", comment_id)
	return tx.Error
}

func (c *CommentStore) GetCommentByPostID(post_id string) ([]*types.Comment, error) {
	return nil, nil
}

