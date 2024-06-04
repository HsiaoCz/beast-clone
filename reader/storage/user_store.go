package storage

import (
	"context"
	"errors"

	"github.com/HsiaoCz/beast-clone/reader/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserStorer interface {
	CreateUser(context.Context, *models.User) (*models.User, error)
	UpdateUser(context.Context, primitive.ObjectID, *models.UserUpdateParams) (*models.User, error)
	GetUserByID(context.Context, primitive.ObjectID) (*models.User, error)
	// GetUserByEmail(context.Context, string) (*models.User, error)
	DeleteUserByID(context.Context, primitive.ObjectID) error
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

func (m *MongoUserStore) UpdateUser(ctx context.Context, uid primitive.ObjectID, userupdate *models.UserUpdateParams) (*models.User, error) {
	filter := bson.D{{Key: "_id", Value: uid}}

	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "username", Value: userupdate.Username},
			{Key: "content", Value: userupdate.Content},
			{Key: "avatar", Value: userupdate.Avatar},
		}},
	}
	updateOptions := options.Update().SetUpsert(true)
	_, err := m.coll.UpdateOne(ctx, filter, update, updateOptions)
	if err != nil {
		return nil, err
	}
	res := models.User{}
	if err := m.coll.FindOne(ctx, filter).Decode(&res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (m *MongoUserStore) GetUserByID(ctx context.Context, uid primitive.ObjectID) (*models.User, error) {
	user := models.User{}
	filter := bson.D{{Key: "_id", Value: uid}}
	if err := m.coll.FindOne(ctx, filter).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

// seems we dont need this shit
// func (m *MongoUserStore) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
// 	user := models.User{}
// 	filter := bson.D{{Key: "email", Value: email}}
// 	if err := m.coll.FindOne(ctx, filter).Decode(&user); err != nil {
// 		return nil, err
// 	}
// 	return &user, nil
// }

func (m *MongoUserStore) DeleteUserByID(ctx context.Context, uid primitive.ObjectID) error {
	result, err := m.coll.DeleteOne(ctx, bson.D{{Key: "_id", Value: uid}})
	if err != nil {
		return errors.New("delete this record failed")
	}
	if result.DeletedCount == 0 {
		return errors.New("database dosen't have this record")
	}
	return nil
}
