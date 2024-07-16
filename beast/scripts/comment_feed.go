package scripts

import (
	"github.com/HsiaoCz/beast-clone/twitter/types"
	"gorm.io/gorm"
)

type CommentFeed struct {
	db *gorm.DB
}

func NewCommentFeed(db *gorm.DB) *CommentFeed {
	return &CommentFeed{
		db: db,
	}
}

func (c *CommentFeed) CreateComment(comment *types.Comment) (*types.Comment, error) {
	tx := c.db.Model(&types.Comment{}).Create(comment)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return comment, nil
}
