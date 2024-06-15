package scripts

import (
	"context"

	"github.com/HsiaoCz/beast-clone/twitter/etc"
	"github.com/HsiaoCz/beast-clone/twitter/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type postFeedStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func newPostFeedStore(ctx context.Context) (*postFeedStore, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(etc.Conf.App.MongoUri))
	if err != nil {
		return nil, err
	}
	return &postFeedStore{client: client, coll: client.Database(etc.Conf.App.DBName).Collection(etc.Conf.App.UserColl)}, nil
}

func (p *postFeedStore) createPost(ctx context.Context, post *types.Post) (*types.Post, error) {
	result, err := p.coll.InsertOne(ctx, post)
	if err != nil {
		return nil, err
	}
	post.ID = result.InsertedID.(primitive.ObjectID)
	return post, nil
}
