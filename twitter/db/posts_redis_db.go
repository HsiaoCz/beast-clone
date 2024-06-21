package db

import "context"

type PostRedisCaser interface {
	LikePost(context.Context, string) (string, error)
	UnLikePost(context.Context, string) (string, error)
	WatchPost(context.Context, string) (string, error)
}
