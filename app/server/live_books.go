package server

import (
	"context"

	pb "github.com/satttto/bookshelf/proto-client/api"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *BookshelfServer) ListBooks(ctx context.Context, req *emptypb.Empty) (*pb.ListBooksResponse, error) {
	books, err := s.service.ListBooks(ctx)
	if err != nil {
		return nil, err
	}

	var bookPb pb.Book
	bookPbs := make([]*pb.Book, len(books))
	for i, book := range books {
		bookPb = convertBookModelToPb(book)
		bookPbs[i] = &bookPb
	}

	return &pb.ListBooksResponse{
		Books: bookPbs,
	}, nil
}
