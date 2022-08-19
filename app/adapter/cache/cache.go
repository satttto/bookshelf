package cache

import (
	"context"
)

type Cache interface {
	Get(ctx context.Context, key string) (string, error)
}
