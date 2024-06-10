package scripts

import (
	"context"
	"errors"

	"github.com/HsiaoCz/beast-clone/twitter/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoUrl = "mongodb://127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=2000&appName=mongosh+2.2.5"
	dbname   = "twtest"
	userColl = "users"
)

type TestUserStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewTestStore() (*TestUserStore, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoUrl))
	if err != nil {
		return nil, err
	}
	return &TestUserStore{
		client: client,
		coll:   client.Database(dbname).Collection(userColl),
	}, nil
}

func (t *TestUserStore) CreateUser(ctx context.Context, user *types.User) (*types.User, error) {
	filter := bson.M{
		"email":       user.Email,
		"phoneNumber": user.PhoneNumber,
	}
	cusor := t.coll.FindOne(ctx, filter)
	if cusor.Err() != mongo.ErrNoDocuments {
		// log.Fatal(cusor.Err())
		return nil, errors.New("error record exists")
	}
	result, err := t.coll.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	user.ID = result.InsertedID.(primitive.ObjectID)
	return user, nil
}
