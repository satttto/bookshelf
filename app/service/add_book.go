package service

import (
	"context"

	"github.com/satttto/bookshelf/app/model"
)

type AddBookServiceInput struct {
}

func (s *BookshelfServiceImp) AddBook(ctx context.Context, in AddBookServiceInput) (model.Book, error) {
	return model.Book{
		ID: 1,
	}, nil
}
