package scripts

import (
	"context"
	"errors"

	"github.com/HsiaoCz/beast-clone/reader/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoUri = "mongodb://127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=2000&appName=mongosh+2.2.5"
	dbname   = "reader"
	userColl = "users"
)

type Feed struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func Newfeed() (*Feed, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoUri))
	if err != nil {
		return nil, err
	}
	return &Feed{client: client, coll: client.Database(dbname).Collection(userColl)}, nil
}

func (f *Feed) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	var check models.User
	filter := bson.D{{Key: "email", Value: user.Email}}
	if err := f.coll.FindOne(ctx, filter).Decode(&check); err != mongo.ErrNoDocuments {
		return nil, errors.New("create user failed because this record exists")
	}
	result, err := f.coll.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	user.ID = result.InsertedID.(primitive.ObjectID)
	return user, nil
}

func (f *Feed) DeleteUser(ctx context.Context, uid primitive.ObjectID) error {
	result, err := f.coll.DeleteOne(ctx, bson.D{{Key: "_id", Value: uid}})
	if err != nil {
		return errors.New("delete this record failed")
	}
	if result.DeletedCount == 0 {
		return errors.New("database dosen't have this record")
	}
	return nil
}
