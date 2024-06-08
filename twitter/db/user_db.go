package db

import (
	"context"

	"github.com/HsiaoCz/beast-clone/twitter/types"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserCaser interface {
	CreateUser(context.Context, *types.User) (*types.User, error)
}

type MongoUserCase struct {
	Client *mongo.Client
	Coll   *mongo.Collection
}

func NewMongoUserCase(client *mongo.Client, coll *mongo.Collection) *MongoUserCase {
	return &MongoUserCase{
		Client: client,
		Coll:   coll,
	}
}

func (m *MongoUserCase) CreateUser(ctx context.Context, user *types.User) (*types.User, error) {
	return nil, nil
}
