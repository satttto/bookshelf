package service

import (
	"context"

	"github.com/satttto/bookshelf/app/model"
)

type GetBookServiceInput struct {
	Id int64
}

func (s BookshelfServiceImp) GetBook(ctx context.Context, in *GetBookServiceInput) (*model.Book, error) {
	book, found, err := s.cache.GetBook(ctx, in.Id)
	if err != nil {
		return nil, err
	}

	if !found {
		book, err = s.db.GetBook(ctx, in.Id)
		if err != nil {
			return nil, err
		}
	}

	return book, nil

}
