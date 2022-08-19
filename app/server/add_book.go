package server

import (
	"context"

	"github.com/satttto/bookshelf/app/service"
	pb "github.com/satttto/bookshelf/proto-client/api"
)

func (s *BookshelfServer) AddBook(ctx context.Context, req *pb.AddBookRequest) (*pb.AddBookResponse, error) {
	in := convertAddBookReqToServiceInput(req)
	out, err := s.service.AddBook(ctx, in)
	if err != nil {
		return nil, err
	}

	return &pb.AddBookResponse{
		Book: &pb.Book{
			Id:       out.ID,
			Title:    "",
			Category: 0,
			Author:   "",
		},
	}, nil
}

func convertAddBookReqToServiceInput(req *pb.AddBookRequest) service.AddBookServiceInput {
	return service.AddBookServiceInput{}
}
