package server

import (
	"github.com/satttto/bookshelf/app/service"
	pb "github.com/satttto/bookshelf/proto-client/api"
	"google.golang.org/grpc"
)

func New(service service.BookshelfService) *grpc.Server {
	grpcServer := grpc.NewServer()
	bookshelf := newBookshelfServer(service)
	pb.RegisterBookshelfServiceServer(grpcServer, bookshelf)
	return grpcServer
}

type BookshelfServer struct {
	pb.UnimplementedBookshelfServiceServer
	service service.BookshelfService
}

func newBookshelfServer(service service.BookshelfService) pb.BookshelfServiceServer {
	return &BookshelfServer{
		service: service,
	}
}
