package service

import (
	"context"

	"github.com/satttto/bookshelf/app/adapter/rdb"
	"github.com/satttto/bookshelf/app/model"
	"github.com/satttto/bookshelf/app/pkg/logger"
)

type BookshelfService interface {
	AddBook(ctx context.Context, in AddBookServiceInput) (model.Book, error)
	ListBooks(ctx context.Context) ([]model.Book, error)
}

type BookshelfServiceImp struct {
	db     rdb.DB
	logger *logger.Logger
}

func New(logger *logger.Logger, db rdb.DB) BookshelfService {
	return &BookshelfServiceImp{
		logger: logger,
		db:     db,
	}
}
