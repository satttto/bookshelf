package service

import (
	"context"

	"github.com/satttto/bookshelf/app/model"
)

type AddBookServiceInput struct {
	Title    string
	Category model.Category
	Author   string
}

func (s *BookshelfServiceImp) AddBook(ctx context.Context, in *AddBookServiceInput) (*model.Book, error) {
	book := &model.Book{
		Title:    in.Title,
		Category: in.Category,
		Auther:   in.Author,
	}

	book, err := s.db.AddBook(ctx, book)
	if err != nil {
		return nil, err
	}

	if err := s.cache.PutBook(ctx, book); err != nil {
		return nil, err
	}

	return book, nil
}
