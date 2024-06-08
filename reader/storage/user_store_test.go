package storage

import (
	"context"
	"testing"

	"github.com/HsiaoCz/beast-clone/reader/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ()

type UserTest struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewUserTest(mongoUrl, dbName, userColl string) (*UserTest, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoUrl))
	if err != nil {
		return nil, err
	}
	return &UserTest{
		client: client,
		coll:   client.Database(dbName).Collection(userColl),
	}, nil
}

func (u *UserTest) UpdateUser(ctx context.Context, uid primitive.ObjectID, userupdate *models.UserUpdateParams) (*models.User, error) {
	filter := bson.D{{Key: "_id", Value: uid}}

	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "username", Value: userupdate.Username},
			{Key: "content", Value: userupdate.Content},
			{Key: "avatar", Value: userupdate.Avatar},
		}},
	}
	updateOptions := options.Update().SetUpsert(true)
	_, err := u.coll.UpdateOne(ctx, filter, update, updateOptions)
	if err != nil {
		return nil, err
	}
	res := models.User{}
	if err := u.coll.FindOne(ctx, filter).Decode(&res); err != nil {
		return nil, err
	}
	return &res, nil
}

func TestUpdateUser(t *testing.T) {
	var (
		mongoUri = "mongodb://127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=2000&appName=mongosh+2.2.5"
		dbname   = "reader"
		userColl = "users"
	)

	ut, err := NewUserTest(mongoUri, dbname, userColl)
	if err != nil {
		t.Fatal(err)
	}
	userUpdateParams := models.UserUpdateParams{
		Username: "santinal",
		Content:  "慢煮生活",
		Avatar:   "./data/user/avatar/11212.jpg",
	}
	uid, err := primitive.ObjectIDFromHex("665d0a80ad93bc8e7712c26d")
	if err != nil {
		t.Fatal(err)
	}
	result, err := ut.UpdateUser(context.Background(), uid, &userUpdateParams)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}
