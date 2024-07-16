package scripts

import (
	"github.com/HsiaoCz/beast-clone/beast/types"
	"gorm.io/gorm"
)

type PostFeed struct {
	db *gorm.DB
}

func NewPostFeed(db *gorm.DB) *PostFeed {
	return &PostFeed{
		db: db,
	}
}

func (p *PostFeed) CreatePost(post *types.Post) (*types.Post, error) {
	tx := p.db.Model(&types.Post{}).Create(post)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return post, nil
}
