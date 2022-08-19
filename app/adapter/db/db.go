package db

import (
	"context"

	"github.com/satttto/bookshelf/app/model"
)

type DB interface {
	Migrate() error
	AddBook(ctx context.Context, book model.Book) (model.Book, error)
	ListBooks(ctx context.Context) ([]model.Book, error)
}
