package cache

import (
	"context"

	"github.com/satttto/bookshelf/app/model"
)

type Cache interface {
	GetBook(ctx context.Context, id int64) (*model.Book, bool, error)
	PutBook(ctx context.Context, book *model.Book) error
}
