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
	// DeleteCommentByID need two ids
	// userID and commentID
	DeleteCommentByID(context.Context, primitive.ObjectID, primitive.ObjectID) error
	GetCommentsByPostID(context.Context, primitive.ObjectID) ([]*types.Comment, error)
	// GetCommentsByPostIDAndParentID need two ids
	// postID and parentID
	GetCommentsByPostIDAndParentID(context.Context, primitive.ObjectID, primitive.ObjectID) ([]*types.Comment, error)
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

func (m *MongoCommentStore) GetCommentsByPostID(ctx context.Context, postID primitive.ObjectID) ([]*types.Comment, error) {
	filter := bson.D{
		{Key: "postID", Value: postID},
	}
	cusor, err := m.coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var comments []*types.Comment
	if err := cusor.All(ctx, comments); err != nil {
		return nil, err
	}
	return comments, nil
}

func (m *MongoCommentStore) GetCommentsByPostIDAndParentID(ctx context.Context, postID primitive.ObjectID, parentID primitive.ObjectID) ([]*types.Comment, error) {
	filter := bson.M{
		"postID":   postID,
		"parentID": parentID,
	}
	cur, err := m.coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var comments []*types.Comment
	if err := cur.All(ctx, comments); err != nil {
		return nil, err
	}
	return comments, nil
}
