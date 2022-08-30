package server

import (
	"context"

	"github.com/satttto/bookshelf/app/service"
	pb "github.com/satttto/bookshelf/proto-client/api"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *BookshelfServer) GetBook(ctx context.Context, req *pb.GetBookRequest) (*pb.GetBookResponse, error) {
	in := convertGetBookReqToServiceInput(req)

	out, err := s.service.GetBook(ctx, in)
	if err != nil {
		return nil, err
	}

	return &pb.GetBookResponse{
		Book: &pb.Book{
			Id:        out.ID,
			Title:     out.Title,
			Category:  convertCategoryModelToPb(out.Category),
			Author:    out.Auther,
			CreatedAt: timestamppb.New(out.CreatedAt),
			UpdatedAt: timestamppb.New(out.UpdatedAt),
			DeletedAt: timestamppb.New(out.DeletedAt),
		},
	}, nil
}

func convertGetBookReqToServiceInput(req *pb.GetBookRequest) *service.GetBookServiceInput {
	return &service.GetBookServiceInput{
		Id: req.Id,
	}
}
