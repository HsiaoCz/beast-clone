package db

import (
	"context"
)

type UserRedisStorer interface {
	SubscribeUser(context.Context, string, string) (string, error)
	UnSubscribeUser(context.Context, string, string) (string, error)
	BlackUser(context.Context, string, string) (string, error)
	UnBlackUser(context.Context, string, string) (string, error)
}
