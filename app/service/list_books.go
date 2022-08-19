package service

import (
	"context"

	"github.com/satttto/bookshelf/app/model"
)

func (s *BookshelfServiceImp) ListBooks(ctx context.Context) ([]model.Book, error) {
	books, err := s.db.ListBooks(ctx)
	if err != nil {
		return nil, err
	}

	return books, nil
}
