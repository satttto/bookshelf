package service

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/satttto/bookshelf/app/model"
)

type AddBookServiceInput struct {
	Title    string
	Category model.Category
	Author   string
}

func (s *BookshelfServiceImp) AddBook(ctx context.Context, in AddBookServiceInput) (model.Book, error) {
	book := model.Book{
		Title:    in.Title,
		Category: in.Category,
		Auther:   in.Author,
	}

	book, err := s.db.AddBook(ctx, book)
	if err != nil {
		return model.Book{}, err
	}

	json, err := json.Marshal(book)
	if err != nil {
		return model.Book{}, err
	}
	if err := s.cache.Put(ctx, strconv.FormatInt(book.ID, 10), string(json)); err != nil {
		return model.Book{}, err
	}

	return book, nil
}
