package data

import (
	"context"
	"errors"

	"github.com/HsiaoCz/beast-clone/clean/sweetwater/internal/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserDataInter interface {
	CreateUser(context.Context, *types.User) (*types.User, error)
	DeleteUser(context.Context, primitive.ObjectID) error
	GetUserByID(context.Context, primitive.ObjectID) (*types.User, error)
	UpdateUser(context.Context, *types.UserUpdateParams) (*types.User, error)
}

type UserData struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func UserDataInit(client *mongo.Client, coll *mongo.Collection) *UserData {
	return &UserData{
		client: client,
		coll:   coll,
	}
}

func (u *UserData) CreateUser(ctx context.Context, user *types.User) (*types.User, error) {
	// 第一步：判断数据库里面是否已经存在当前用户
	// 基于什么来判断呢？ 基于字段 email
	filter := bson.D{
		{Key: "email", Value: user.Email},
	}
	if err := u.coll.FindOne(ctx, filter).Err(); err != mongo.ErrNoDocuments {
		return nil, errors.New("database has this record")
	}
	// 第二步，将提交的数据插入数据库
	result, err := u.coll.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	// 第三步，将结果的id赋值给user.id
	user.ID = result.InsertedID.(primitive.ObjectID)
	// 第四步 返回结果，这里就是原来的user,不过是赋值过id的user
	return user, nil
}

func (u *UserData) DeleteUser(ctx context.Context, id primitive.ObjectID) error {
	// 第一步: 组装filter 字段为id
	filter := bson.D{
		{Key: "_id", Value: id},
	}
	// 删除记录
	result, err := u.coll.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("there is no this record")
	}
	return nil
}

func (u *UserData) GetUserByID(ctx context.Context, id primitive.ObjectID) (*types.User, error) {
	var user types.User
	filter := bson.D{
		{Key: "_id", Value: id},
	}
	if err := u.coll.FindOne(ctx, filter).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserData) UpdateUser(ctx context.Context, updateParmas *types.UserUpdateParams) (*types.User, error) {
	return nil, nil
}
