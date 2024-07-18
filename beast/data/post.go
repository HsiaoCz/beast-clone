package data

import (
	"github.com/HsiaoCz/beast-clone/beast/types"
	"gorm.io/gorm"
)

type PostStorer interface {
	CreatePost(*types.Post) (*types.Post, error)
	GetPostByID(string) (*types.Post, error)
	DeletePostByID(string) error
}

type PostStore struct {
	db *gorm.DB
}

func NewPostStore(db *gorm.DB) *PostStore {
	return &PostStore{
		db: db,
	}
}

func (p *PostStore) CreatePost(post *types.Post) (*types.Post, error) {
	tx := p.db.Model(&types.Post{}).Create(post)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return post, nil
}

func (p *PostStore) GetPostByID(post_id string) (*types.Post, error) {
	var post types.Post
	tx := p.db.Model(&types.Post{}).Find(&post, "post_id = ?", post_id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &post, nil
}

func (p *PostStore) DeletePostByID(post_id string) error {
	return p.db.Where("post_id = ?", post_id).Delete(&types.Post{}).Error
}
