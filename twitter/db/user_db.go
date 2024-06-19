package db

import (
	"context"
	"errors"

	"github.com/HsiaoCz/beast-clone/twitter/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserCaser interface {
	CreateUser(context.Context, *types.User) (*types.User, error)
	GetUserByID(context.Context, primitive.ObjectID) (*types.User, error)
	DeleteUserByID(context.Context, primitive.ObjectID) error
	UpdateUserByID(context.Context, primitive.ObjectID, *types.UpdateUserParams) (*types.User, error)
}

type MongoUserCase struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoUserCase(client *mongo.Client, coll *mongo.Collection) *MongoUserCase {
	return &MongoUserCase{
		client: client,
		coll:   coll,
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
	cusor := m.coll.FindOne(ctx, filter)
	if cusor.Err() != mongo.ErrNoDocuments {
		// log.Fatal(cusor.Err())
		return nil, errors.New("error record exists")
	}
	result, err := m.coll.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	user.ID = result.InsertedID.(primitive.ObjectID)
	return user, nil
}

func (m *MongoUserCase) GetUserByID(ctx context.Context, uid primitive.ObjectID) (*types.User, error) {
	filter := bson.D{
		{Key: "_id", Value: uid},
	}
	user := types.User{}
	if err := m.coll.FindOne(ctx, filter).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (m *MongoUserCase) DeleteUserByID(ctx context.Context, uid primitive.ObjectID) error {
	filter := bson.D{
		{Key: "_id", Value: uid},
	}
	_, err := m.coll.DeleteOne(ctx, filter)
	return err
}

func (m *MongoUserCase) UpdateUserByID(ctx context.Context, uid primitive.ObjectID, updateUserParams *types.UpdateUserParams) (*types.User, error) {
	filter := bson.D{
		{Key: "_id", Value: uid},
	}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "firstName", Value: updateUserParams.FirstName},
			{Key: "lastName", Value: updateUserParams.LastName},
			{Key: "nickName", Value: updateUserParams.NickName},
			{Key: "avatar", Value: updateUserParams.Avatar},
			{Key: "backgroudImage", Value: updateUserParams.BackgroundImage},
			{Key: "synopsis", Value: updateUserParams.Synopsis},
		}},
	}
	updateOptions := options.Update().SetUpsert(true)
	_, err := m.coll.UpdateOne(ctx, filter, update, updateOptions)
	if err != nil {
		return nil, err
	}
	res := &types.User{}
	if err := m.coll.FindOne(ctx, filter).Decode(res); err != nil {
		return nil, err
	}

	return res, nil
}
