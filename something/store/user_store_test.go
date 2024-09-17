package store

import (
	"context"
	"os"
	"testing"

	"github.com/HsiaoCz/beast-clone/hotel/types"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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

func (u *UserTest) UpdateUser(ctx context.Context, uid primitive.ObjectID, userupdate *types.UpdateUserParams) (*types.User, error) {
	filter := bson.D{{Key: "_id", Value: uid}}

	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "firstName", Value: userupdate.FirstName},
			{Key: "lastName", Value: userupdate.LastName},
		}},
	}
	updateOptions := options.Update().SetUpsert(true)
	_, err := u.coll.UpdateOne(ctx, filter, update, updateOptions)
	if err != nil {
		return nil, err
	}
	res := types.User{}
	if err := u.coll.FindOne(ctx, filter).Decode(&res); err != nil {
		return nil, err
	}
	return &res, nil
}

func TestUpdateUser(t *testing.T) {
	if err := godotenv.Load(); err != nil {
		t.Fatal(err)
	}
	var (
		mongoURL     = os.Getenv("MONGOURL")
		dbName       = os.Getenv("DBNAME")
		userCollName = os.Getenv("USERCOLL")
	)

	ut, err := NewUserTest(mongoURL, dbName, userCollName)
	if err != nil {
		t.Fatal(err)
	}
	userUpdateParams := types.UpdateUserParams{
		FirstName: "shangs",
		LastName:  "santinal",
	}
	uid, err := primitive.ObjectIDFromHex("66597e0918d1205254d56d75")
	if err != nil {
		t.Fatal(err)
	}
	result, err := ut.UpdateUser(context.Background(), uid, &userUpdateParams)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}
