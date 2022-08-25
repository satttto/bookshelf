package redis

import (
	"context"

	"github.com/go-redis/redis/v8"

	"github.com/satttto/bookshelf/app/adapter/cache"
)

type Redis struct {
	client *redis.Client
}

func New(addr string, username string, password string) (cache.Cache, error) {
	cache := redis.NewClient(&redis.Options{
		Addr:     addr,
		Username: username,
		Password: password,
	})

	return &Redis{
		client: cache,
	}, nil
}

func (r *Redis) Get(ctx context.Context, key string) (string, error) {
	return "", nil
}

func (r *Redis) Put(ctx context.Context, key string, value string) error {
	return nil
}
