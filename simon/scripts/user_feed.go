package scripts

import (
	"context"
	"errors"
	"os"

	"github.com/HsiaoCz/beast-clone/simon/st"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserFeed struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewUserfeed() (*UserFeed, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("MONGORUL")))
	if err != nil {
		return nil, err
	}
	return &UserFeed{client: client, coll: client.Database(os.Getenv("DBNAME")).Collection(os.Getenv("USERCOLL"))}, nil
}

func (f *UserFeed) CreateUser(ctx context.Context, user *st.User) (*st.User, error) {
	var check st.User
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
