package db

import (
	"context"

	"github.com/HsiaoCz/beast-clone/twitter/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PostCaser interface {
	CreatePost(context.Context, *types.Post) (*types.Post, error)
}

type MongoPostStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoPostStore(client *mongo.Client, coll *mongo.Collection) *MongoPostStore {
	return &MongoPostStore{
		client: client,
		coll:   coll,
	}
}

func (m *MongoPostStore) CreatePost(ctx context.Context, post *types.Post) (*types.Post, error) {
	resp, err := m.coll.InsertOne(ctx, post)
	if err != nil {
		return nil, err
	}
	post.ID = resp.InsertedID.(primitive.ObjectID)
	return post, nil
}
