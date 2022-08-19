package server

import (
	"github.com/satttto/bookshelf/app/model"
	pb "github.com/satttto/bookshelf/proto-client/api"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertCategoryModelToPb(category model.Category) pb.Category {
	switch category {
	case model.CATEGORY_TECH:
		return pb.Category_CATEGORY_TECH
	default:
		return pb.Category_CATEGORY_UNKNOWN
	}
}

func convertCategoryPbToModel(category pb.Category) model.Category {
	switch category {
	case pb.Category_CATEGORY_TECH:
		return model.CATEGORY_TECH
	default:
		return model.CATEGORY_UNKNOWN
	}
}

func convertBookModelToPb(model model.Book) pb.Book {
	return pb.Book{
		Id:        model.ID,
		Title:     model.Title,
		Category:  convertCategoryModelToPb(model.Category),
		Author:    model.Auther,
		CreatedAt: timestamppb.New(model.CreatedAt),
		UpdatedAt: timestamppb.New(model.UpdatedAt),
		DeletedAt: timestamppb.New(model.DeletedAt),
	}
}
