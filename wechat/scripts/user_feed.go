package scripts

import (
	"context"
	"errors"
	"os"

	"github.com/HsiaoCz/beast-clone/wechat/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserFeed struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewUserFeed(ctx context.Context) (*UserFeed, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGOURL")))
	if err != nil {
		return nil, err
	}
	return &UserFeed{
		client: client,
		coll:   client.Database(os.Getenv("DBNAME")).Collection(os.Getenv("USERCOLL")),
	}, nil
}

func (u *UserFeed) CreateUser(ctx context.Context, user *types.User) (*types.User, error) {
	filter := bson.D{
		{Key: "email", Value: user.Email},
	}
	result := u.coll.FindOne(ctx, filter)
	if result.Err() != mongo.ErrNoDocuments {
		return nil, errors.New("this record exists")
	}
	cursor, err := u.coll.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	user.ID = cursor.InsertedID.(primitive.ObjectID)
	return user, nil
}
