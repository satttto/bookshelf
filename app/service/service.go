package service

import (
	"context"

	"github.com/satttto/bookshelf/app/adapter/cache"
	"github.com/satttto/bookshelf/app/adapter/rdb"
	"github.com/satttto/bookshelf/app/model"
	"github.com/satttto/bookshelf/app/pkg/logger"
)

type BookshelfService interface {
	AddBook(ctx context.Context, in *AddBookServiceInput) (*model.Book, error)
	ListBooks(ctx context.Context) ([]model.Book, error)
	GetBook(ctx context.Context, in *GetBookServiceInput) (*model.Book, error)
}

type BookshelfServiceImp struct {
	logger *logger.Logger
	db     rdb.DB
	cache  cache.Cache
}

func New(logger *logger.Logger, db rdb.DB, cache cache.Cache) BookshelfService {
	return &BookshelfServiceImp{
		logger: logger,
		db:     db,
		cache:  cache,
	}
}
