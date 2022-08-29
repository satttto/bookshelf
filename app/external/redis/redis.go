package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"

	"github.com/satttto/bookshelf/app/adapter/cache"
	"github.com/satttto/bookshelf/app/model"
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

func (r *Redis) GetBook(ctx context.Context, id int64) (*model.Book, bool, error) {
	strId := fmt.Sprintf("book-%d", id)

	strBook, err := r.client.Get(ctx, strId).Result()
	if err == redis.Nil {
		return nil, false, nil
	} else if err != nil {
		return nil, false, err
	}

	var book model.Book
	if err := json.Unmarshal([]byte(strBook), &book); err != nil {
		return nil, false, errors.Wrap(err, "failed to unmarshal book data")
	}

	return &book, true, nil
}

func (r *Redis) PutBook(ctx context.Context, book *model.Book) error {
	strBook, err := json.Marshal(*book)
	if err != nil {
		return errors.Wrap(err, "failed to marshal book data")
	}

	strId := fmt.Sprintf("book-%d", book.ID)
	if err := r.client.Set(ctx, strId, strBook, 10*time.Hour).Err(); err != nil {
		return errors.Wrap(err, "failed to set a book to Redis cache")
	}

	return nil
}
