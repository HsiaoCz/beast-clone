package data

import (
	"context"

	"github.com/HsiaoCz/beast-clone/gustao/types"
	"go.mongodb.org/mongo-driver/mongo"
)

type PostDataInter interface {
	CreatePost(context.Context, *types.Posts) (*types.Posts, error)
	GetPostByID(context.Context, string) (*types.Posts, error)
	GetPostsByUserID(context.Context, string) ([]*types.Posts, error)
}

type PostData struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func PostDataInit(client *mongo.Client, coll *mongo.Collection) *PostData {
	return &PostData{
		client: client,
		coll:   coll,
	}
}

func (p *PostData) CreatePost(ctx context.Context, posts *types.Posts) (*types.Posts, error) {
	return nil, nil
}

func (p *PostData) GetPostByID(ctx context.Context, post_id string) (*types.Posts, error) {
	return nil, nil
}

func (p *PostData) GetPostsByUserID(ctx context.Context, user_id string) ([]*types.Posts, error) {
	return nil, nil
}
