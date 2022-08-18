package server

import (
	"context"

	pb "github.com/satttto/bookshelf/proto-client/api"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

func New() *grpc.Server {
	grpcServer := grpc.NewServer()
	bookshelf := newBookshelfServer()
	pb.RegisterBookshelfServiceServer(grpcServer, bookshelf)
	return grpcServer
}

type BookshelfServer struct {
	pb.UnimplementedBookshelfServiceServer
}

func newBookshelfServer() *BookshelfServer {
	return &BookshelfServer{}
}

func (s *BookshelfServer) AddBook(ctx context.Context, req *pb.AddBookRequest) (*pb.AddBookResponse, error) {
	return &pb.AddBookResponse{}, nil
}

func (s *BookshelfServer) ListBooks(ctx context.Context, req *emptypb.Empty) (*pb.ListBooksResponse, error) {
	return &pb.ListBooksResponse{}, nil
}
