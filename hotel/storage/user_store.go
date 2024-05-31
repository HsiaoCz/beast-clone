package storage

import (
	"context"
	"errors"

	"github.com/HsiaoCz/beast-clone/hotel/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserStorer interface {
	CreateUser(context.Context, *types.User) (*types.User, error)
	GetUserByEmail(context.Context, string) (*types.User, error)
	GetUserByID(context.Context, primitive.ObjectID) (*types.User, error)
}

type MongoUserStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoUserStore(client *mongo.Client, coll *mongo.Collection) *MongoUserStore {
	return &MongoUserStore{
		client: client,
		coll:   coll,
	}
}

func (m *MongoUserStore) CreateUser(ctx context.Context, user *types.User) (*types.User, error) {
	var check types.User
	filter := bson.D{{Key: "email", Value: user.Email}}
	if err := m.coll.FindOne(ctx, filter).Decode(&check); err != mongo.ErrNoDocuments {
		return nil, errors.New("create user failed because this record exists")
	}
	result, err := m.coll.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	user.ID = result.InsertedID.(primitive.ObjectID)
	return user, nil
}

func (m *MongoUserStore) GetUserByEmail(ctx context.Context, email string) (*types.User, error) {
	check := types.User{}
	filter := bson.D{{Key: "email", Value: email}}
	if err := m.coll.FindOne(ctx, filter).Decode(&check); err != nil {
		return nil, errors.New("database doesnt havh this record")
	}
	return &check, nil
}

func (m *MongoUserStore) GetUserByID(ctx context.Context, id primitive.ObjectID) (*types.User, error) {
	user := types.User{}
	if err := m.coll.FindOne(ctx, bson.D{{Key: "_id", Value: id}}).Decode(&user); err != nil {
		return nil, errors.New("query database failed please check the query params")
	}
	return &user, nil
}
