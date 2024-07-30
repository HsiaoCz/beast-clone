package store

import (
	"context"
	"errors"

	"github.com/HsiaoCz/beast-clone/simon/st"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserStorer interface {
	CreateUser(context.Context, *st.User) (*st.User, error)
	GetUserByEmail(context.Context, string) (*st.User, error)
}

type UserStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func UserStoreInit(client *mongo.Client, coll *mongo.Collection) *UserStore {
	return &UserStore{
		client: client,
		coll:   coll,
	}
}

func (u *UserStore) CreateUser(ctx context.Context, user *st.User) (*st.User, error) {
	var check st.User
	filter := bson.D{{Key: "email", Value: user.Email}}
	if err := u.coll.FindOne(ctx, filter).Decode(&check); err != mongo.ErrNoDocuments {
		return nil, errors.New("create user failed because this record exists")
	}
	result, err := u.coll.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	user.ID = result.InsertedID.(primitive.ObjectID)
	return user, nil
}

func (u *UserStore) GetUserByEmail(ctx context.Context, email string) (*st.User, error) {
	var check st.User
	filter := bson.D{
		{Key: "email", Value: email},
	}
	if err := u.coll.FindOne(ctx, filter).Decode(&check); err != nil {
		return nil, errors.New("database doesnt hava this record")
	}
	return &check, nil
}
