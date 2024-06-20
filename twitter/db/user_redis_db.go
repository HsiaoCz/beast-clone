package db

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type UserRedisStorer interface {
	SubscribeUser(context.Context, string, string) (string, error)
	UnSubscribeUser(context.Context, string, string) (string, error)
	BlackUser(context.Context, string, string) (string, error)
	UnBlackUser(context.Context, string, string) (string, error)
}

type RedisUserStore struct {
	db *redis.Client
}

func NewRedisUserStore(db *redis.Client) *RedisUserStore {
	return &RedisUserStore{
		db: db,
	}
}

func (r *RedisUserStore) SubscribeUser(ctx context.Context, userID string, beUserID string) (string, error) {
	return "", nil
}

func (r *RedisUserStore) UnSubscribeUser(ctx context.Context, userID string, subUserID string) (string, error) {
	return "", nil
}

func (r *RedisUserStore) BlackUser(ctx context.Context, userID string, blackUserID string) (string, error) {
	return "", nil
}
func (r *RedisUserStore) UnBlackUser(ctx context.Context, userID string, blackUserID string) (string, error) {
	return "", nil
}
