package db

import (
	"context"
	"errors"

	"github.com/HsiaoCz/beast-clone/twitter/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	// well this is totally shit
	// filter := bson.D{
	// 	{Key: "$get", Value: bson.D{
	// 		{Key: "email", Value: user.Email},
	// 		{Key: "phoneNumber", Value: user.PhoneNumber},
	// 	}},
	// }

	filter := bson.M{
		"email":       user.Email,
		"phoneNumber": user.PhoneNumber,
	}
	cusor := m.Coll.FindOne(ctx, filter)
	if cusor.Err() != mongo.ErrNoDocuments {
		// log.Fatal(cusor.Err())
		return nil, errors.New("error record exists")
	}
	result, err := m.Coll.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	user.ID = result.InsertedID.(primitive.ObjectID)
	return user, nil
}
