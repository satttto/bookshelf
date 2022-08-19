package server

import (
	"context"

	"github.com/satttto/bookshelf/app/service"
	pb "github.com/satttto/bookshelf/proto-client/api"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *BookshelfServer) AddBook(ctx context.Context, req *pb.AddBookRequest) (*pb.AddBookResponse, error) {
	in := convertAddBookReqToServiceInput(req)

	out, err := s.service.AddBook(ctx, in)
	if err != nil {
		return nil, err
	}

	return &pb.AddBookResponse{
		Book: &pb.Book{
			Id:        out.ID,
			Title:     out.Title,
			Category:  convertCategoryModelToPb(out.Category),
			Author:    out.Auther,
			CreatedAt: timestamppb.New(out.CreatedAt),
			UpdatedAt: timestamppb.New(out.UpdatedAt),
		},
	}, nil
}

func convertAddBookReqToServiceInput(req *pb.AddBookRequest) service.AddBookServiceInput {
	return service.AddBookServiceInput{
		Title:    req.Title,
		Category: convertCategoryPbToModel(req.Category),
		Author:   req.Author,
	}
}
