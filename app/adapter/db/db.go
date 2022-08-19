package db

import (
	"context"

	"github.com/satttto/bookshelf/app/model"
)

type DB interface {
	Migrate() error
	AddBook(ctx context.Context, book *model.Book) error
}
