package service

import (
	"context"

	"github.com/HsiaoCz/beast-clone/anne/types"
)

type FetchPricer interface {
	FetchPrice(context.Context, string) (*types.Price, error)
}
