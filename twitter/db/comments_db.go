package db

import (
	"context"

	"github.com/HsiaoCz/beast-clone/twitter/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CommentCaser interface {
	CreateComment(context.Context, *types.Comment) (*types.Comment, error)
	// DeleteCommentByID need to ids
	// userID and commentID
	DeleteCommentByID(context.Context, primitive.ObjectID, primitive.ObjectID) error
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

func (m *MongoCommentStore) CreateComment(ctx context.Context, comment *types.Comment) (*types.Comment, error) {
	resp, err := m.coll.InsertOne(ctx, comment)
	if err != nil {
		return nil, err
	}
	comment.ID = resp.InsertedID.(primitive.ObjectID)
	return comment, nil
}

func (m *MongoCommentStore) DeleteCommentByID(ctx context.Context, userID primitive.ObjectID, commentID primitive.ObjectID) error {
	filter := bson.M{
		"_id":    commentID,
		"userID": userID,
	}
	resp, err := m.coll.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	if resp.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil
}
