package scripts

import (
	"context"

	"github.com/HsiaoCz/beast-clone/twitter/etc"
	"github.com/HsiaoCz/beast-clone/twitter/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type commentFeedStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func newCommentFeedStore(ctx context.Context) (*commentFeedStore, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(etc.Conf.App.MongoUri))
	if err != nil {
		return nil, err
	}
	return &commentFeedStore{client: client, coll: client.Database(etc.Conf.App.DBName).Collection("comments")}, nil
}

func (c *commentFeedStore) CreateComment(ctx context.Context, comment *types.Comment) (*types.Comment, error) {
	result, err := c.coll.InsertOne(ctx, comment)
	if err != nil {
		return nil, err
	}
	comment.ID = result.InsertedID.(primitive.ObjectID)
	return comment, nil
}
