package storage

import (
	"context"
	"errors"

	"github.com/HsiaoCz/beast-clone/reader/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserStorer interface {
	CreateUser(context.Context, *models.User) (*models.User, error)
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

func (m *MongoUserStore) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	res := models.User{}
	filter := bson.D{{Key: "email", Value: user.Email}}
	if err := m.coll.FindOne(ctx, filter).Decode(&res); err != mongo.ErrNoDocuments {
		return nil, errors.New("this record exists")
	}
	result, err := m.coll.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	user.ID = result.InsertedID.(primitive.ObjectID)
	return user, nil
}
