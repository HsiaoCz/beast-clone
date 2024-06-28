package data

import (
	"context"
	"errors"

	"github.com/HsiaoCz/beast-clone/latency/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type greeterRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewGreeterRepo(data *Data, logger log.Logger) biz.GreeterRepo {
	return &greeterRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *greeterRepo) Save(ctx context.Context, u *biz.User) (*biz.User, error) {
	filter := bson.D{
		{Key: "email", Value: u.Email},
	}
	result := r.data.coll.FindOne(ctx, filter)
	if result.Err() != mongo.ErrNoDocuments {
		return nil, errors.New("database has this record")
	}
	cursor, err := r.data.coll.InsertOne(ctx, u)
	if err != nil {
		return nil, err
	}
	u.ID = cursor.InsertedID.(primitive.ObjectID).String()
	return u, nil
}

func (r *greeterRepo) Update(ctx context.Context, g *biz.User) (*biz.User, error) {
	return g, nil
}

func (r *greeterRepo) FindByID(ctx context.Context, id primitive.ObjectID) (*biz.User, error) {
	filter := bson.D{
		{Key: "_id", Value: id},
	}
	result := r.data.coll.FindOne(ctx, filter)
	if result.Err() != nil {
		return nil, result.Err()
	}
	user := &biz.User{}
	if err := result.Decode(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (r *greeterRepo) ListAll(context.Context) ([]*biz.User, error) {
	return nil, nil
}
