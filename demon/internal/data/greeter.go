package data

import (
	"context"

	"demon/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
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

func (g *greeterRepo) CreateUser(ctx context.Context, user *biz.User) (*biz.User, error) {
	return &biz.User{}, nil
}

func (g *greeterRepo) FindByID(ctx context.Context, uid string) (*biz.User, error) {
	return &biz.User{}, nil
}
