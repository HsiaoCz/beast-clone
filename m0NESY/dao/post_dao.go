package dao

import (
	"github.com/HsiaoCz/beast-clone/m0NESY/types"
	"gorm.io/gorm"
)

type PostDataModer interface {
	CreatePost(*types.Post) (*types.Post, error)
	GetPostByID(string) (*types.Post, error)
}

type PostDataMod struct {
	db *gorm.DB
}

func NewPostDataMod(db *gorm.DB) *PostDataMod {
	return &PostDataMod{
		db: db,
	}
}

func (p *PostDataMod) CreatePost(post *types.Post) (*types.Post, error) {
	tx := p.db.Model(&types.Post{}).Create(post)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return post, nil
}

func (p *PostDataMod) GetPostByID(post_id string) (*types.Post, error) {
	var post types.Post
	tx := p.db.Model(&types.Post{}).Find(&post, "post_id = ?", post_id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &post, nil
}
