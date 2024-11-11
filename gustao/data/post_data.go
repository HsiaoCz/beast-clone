package data

import (
	"context"

	"github.com/HsiaoCz/beast-clone/gustao/types"
	"go.mongodb.org/mongo-driver/mongo"
)

type PostDataInter interface {
	CreatePost(context.Context, *types.Posts) (*types.Posts, error)
	GetPostByID(context.Context, *types.Posts) (*types.Posts, error)
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
