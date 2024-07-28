package store

import (
	"context"

	"github.com/HsiaoCz/beast-clone/simon/st"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserStorer interface {
	CreateUser(context.Context, *st.User) (*st.User, error)
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
