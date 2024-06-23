package db

import (
	"context"

	"github.com/HsiaoCz/beast-clone/twitter/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CommentCaser interface {
	CreateComment(context.Context, primitive.ObjectID, *types.Comment) (*types.Comment, error)
}

type MongoCommentStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoCommentStore(client *mongo.Client, coll *mongo.Collection) *MongoCommentStore {
	return &MongoCommentStore{
		client: client,
		coll:   coll,
	}
}

func (m *MongoCommentStore) CreateComment(ctx context.Context, userID primitive.ObjectID, comment *types.Comment) (*types.Comment, error) {
	return nil, nil
}
