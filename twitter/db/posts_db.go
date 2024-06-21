package db

import (
	"context"
	"errors"

	"github.com/HsiaoCz/beast-clone/twitter/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PostCaser interface {
	CreatePost(context.Context, *types.Post) (*types.Post, error)
	DeletePostByID(context.Context, primitive.ObjectID) error
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

func (m *MongoPostStore) DeletePostByID(ctx context.Context, pid primitive.ObjectID) error {
	filter := bson.D{
		{Key: "_id", Value: pid},
	}

	res, err := m.coll.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return errors.New("no record")
	}
	return nil
}
