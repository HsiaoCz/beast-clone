package data

import (
	"context"

	"github.com/HsiaoCz/beast-clone/latency/internal/biz"
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

func (r *greeterRepo) Save(ctx context.Context, g *biz.User) (*biz.User, error) {
	return g, nil
}

func (r *greeterRepo) Update(ctx context.Context, g *biz.User) (*biz.User, error) {
	return g, nil
}

func (r *greeterRepo) FindByID(context.Context, int64) (*biz.User, error) {
	return nil, nil
}

func (r *greeterRepo) ListAll(context.Context) ([]*biz.User, error) {
	return nil, nil
}
