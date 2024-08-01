package store

import (
	"context"
	"errors"

	"github.com/HsiaoCz/beast-clone/simon/st"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserStorer interface {
	CreateUser(context.Context, *st.User) (*st.User, error)
	GetUserByEmail(context.Context, string) (*st.User, error)
	GetUserByID(context.Context, primitive.ObjectID) (*st.User, error)
	DeleteUserByID(context.Context, primitive.ObjectID) error
	UpdateUser(context.Context, primitive.ObjectID, *st.UpdateUserParams) (*st.User, error)
	VerifyUserPassword(context.Context, primitive.ObjectID, string) error
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

func (u *UserStore) GetUserByID(ctx context.Context, uid primitive.ObjectID) (*st.User, error) {
	var user st.User
	filter := bson.D{
		{Key: "_id", Value: uid},
	}
	if err := u.coll.FindOne(ctx, filter).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserStore) DeleteUserByID(ctx context.Context, uid primitive.ObjectID) error {
	result, err := u.coll.DeleteOne(ctx, bson.D{{Key: "_id", Value: uid}})
	if err != nil {
		return errors.New("delete this record failed")
	}
	if result.DeletedCount == 0 {
		return errors.New("database doesn't have this record")
	}
	return nil
}

func (u *UserStore) UpdateUser(ctx context.Context, uid primitive.ObjectID, params *st.UpdateUserParams) (*st.User, error) {
	filter := bson.D{{Key: "_id", Value: uid}}

	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "username", Value: params.Username},
		}},
	}

	updateOptions := options.Update().SetUpsert(true)
	_, err := u.coll.UpdateOne(ctx, filter, update, updateOptions)
	if err != nil {
		return nil, err
	}
	var user st.User
	if err := u.coll.FindOne(ctx, filter).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserStore) VerifyUserPassword(ctx context.Context, uid primitive.ObjectID, newPasswd string) error {
	filter := bson.D{
		{Key: "_id", Value: uid},
	}

	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "password", Value: newPasswd},
		}},
	}

	updateOption := options.Update().SetUpsert(true)

	res, err := u.coll.UpdateOne(ctx, filter, update, updateOption)
	if err != nil {
		return err
	}

	if res.ModifiedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}
