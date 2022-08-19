package service

import (
	"context"

	"github.com/satttto/bookshelf/app/adapter/db"
	"github.com/satttto/bookshelf/app/model"
	"github.com/satttto/bookshelf/app/pkg/logger"
)

type BookshelfService interface {
	AddBook(ctx context.Context, in AddBookServiceInput) (model.Book, error)
}

type BookshelfServiceImp struct {
	db     db.DB
	logger *logger.Logger
}

func New(logger *logger.Logger, db db.DB) BookshelfService {
	return &BookshelfServiceImp{
		logger: logger,
		db:     db,
	}
}
