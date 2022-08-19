package server

import (
	"context"

	pb "github.com/satttto/bookshelf/proto-client/api"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *BookshelfServer) ListBooks(ctx context.Context, req *emptypb.Empty) (*pb.ListBooksResponse, error) {
	return &pb.ListBooksResponse{}, nil
}
