// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: api/bookshelf.proto

package api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Category int32

const (
	Category_CATEGORY_UNKNOWN Category = 0
	Category_CATEGORY_TECH    Category = 1
)

// Enum value maps for Category.
var (
	Category_name = map[int32]string{
		0: "CATEGORY_UNKNOWN",
		1: "CATEGORY_TECH",
	}
	Category_value = map[string]int32{
		"CATEGORY_UNKNOWN": 0,
		"CATEGORY_TECH":    1,
	}
)

func (x Category) Enum() *Category {
	p := new(Category)
	*p = x
	return p
}

func (x Category) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Category) Descriptor() protoreflect.EnumDescriptor {
	return file_api_bookshelf_proto_enumTypes[0].Descriptor()
}

func (Category) Type() protoreflect.EnumType {
	return &file_api_bookshelf_proto_enumTypes[0]
}

func (x Category) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Category.Descriptor instead.
func (Category) EnumDescriptor() ([]byte, []int) {
	return file_api_bookshelf_proto_rawDescGZIP(), []int{0}
}

type Book struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title     string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Category  Category               `protobuf:"varint,3,opt,name=category,proto3,enum=Category" json:"category,omitempty"`
	Author    string                 `protobuf:"bytes,4,opt,name=author,proto3" json:"author,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,50,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,51,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt *timestamppb.Timestamp `protobuf:"bytes,52,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *Book) Reset() {
	*x = Book{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_bookshelf_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Book) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Book) ProtoMessage() {}

func (x *Book) ProtoReflect() protoreflect.Message {
	mi := &file_api_bookshelf_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Book.ProtoReflect.Descriptor instead.
func (*Book) Descriptor() ([]byte, []int) {
	return file_api_bookshelf_proto_rawDescGZIP(), []int{0}
}

func (x *Book) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Book) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Book) GetCategory() Category {
	if x != nil {
		return x.Category
	}
	return Category_CATEGORY_UNKNOWN
}

func (x *Book) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *Book) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Book) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Book) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

type AddBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title    string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Category Category `protobuf:"varint,2,opt,name=category,proto3,enum=Category" json:"category,omitempty"`
	Author   string   `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
}

func (x *AddBookRequest) Reset() {
	*x = AddBookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_bookshelf_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddBookRequest) ProtoMessage() {}

func (x *AddBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_bookshelf_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddBookRequest.ProtoReflect.Descriptor instead.
func (*AddBookRequest) Descriptor() ([]byte, []int) {
	return file_api_bookshelf_proto_rawDescGZIP(), []int{1}
}

func (x *AddBookRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *AddBookRequest) GetCategory() Category {
	if x != nil {
		return x.Category
	}
	return Category_CATEGORY_UNKNOWN
}

func (x *AddBookRequest) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

type AddBookResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Book *Book `protobuf:"bytes,1,opt,name=book,proto3" json:"book,omitempty"`
}

func (x *AddBookResponse) Reset() {
	*x = AddBookResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_bookshelf_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddBookResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddBookResponse) ProtoMessage() {}

func (x *AddBookResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_bookshelf_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddBookResponse.ProtoReflect.Descriptor instead.
func (*AddBookResponse) Descriptor() ([]byte, []int) {
	return file_api_bookshelf_proto_rawDescGZIP(), []int{2}
}

func (x *AddBookResponse) GetBook() *Book {
	if x != nil {
		return x.Book
	}
	return nil
}

type ListBooksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Books []*Book `protobuf:"bytes,1,rep,name=books,proto3" json:"books,omitempty"`
}

func (x *ListBooksResponse) Reset() {
	*x = ListBooksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_bookshelf_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBooksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBooksResponse) ProtoMessage() {}

func (x *ListBooksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_bookshelf_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBooksResponse.ProtoReflect.Descriptor instead.
func (*ListBooksResponse) Descriptor() ([]byte, []int) {
	return file_api_bookshelf_proto_rawDescGZIP(), []int{3}
}

func (x *ListBooksResponse) GetBooks() []*Book {
	if x != nil {
		return x.Books
	}
	return nil
}

type GetBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetBookRequest) Reset() {
	*x = GetBookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_bookshelf_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBookRequest) ProtoMessage() {}

func (x *GetBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_bookshelf_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBookRequest.ProtoReflect.Descriptor instead.
func (*GetBookRequest) Descriptor() ([]byte, []int) {
	return file_api_bookshelf_proto_rawDescGZIP(), []int{4}
}

func (x *GetBookRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetBookResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Book *Book `protobuf:"bytes,1,opt,name=book,proto3" json:"book,omitempty"`
}

func (x *GetBookResponse) Reset() {
	*x = GetBookResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_bookshelf_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBookResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBookResponse) ProtoMessage() {}

func (x *GetBookResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_bookshelf_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBookResponse.ProtoReflect.Descriptor instead.
func (*GetBookResponse) Descriptor() ([]byte, []int) {
	return file_api_bookshelf_proto_rawDescGZIP(), []int{5}
}

func (x *GetBookResponse) GetBook() *Book {
	if x != nil {
		return x.Book
	}
	return nil
}

var File_api_bookshelf_proto protoreflect.FileDescriptor

var file_api_bookshelf_proto_rawDesc = []byte{
	0x0a, 0x13, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x9c, 0x02, 0x0a, 0x04, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x25, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x09, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52,
	0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x32, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x33, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x34, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x22, 0x65, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x25, 0x0a, 0x08, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x09, 0x2e, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x22, 0x2c, 0x0a, 0x0f, 0x41, 0x64, 0x64,
	0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x04,
	0x62, 0x6f, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x42, 0x6f, 0x6f,
	0x6b, 0x52, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x22, 0x30, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x42,
	0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x05,
	0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x42, 0x6f,
	0x6f, 0x6b, 0x52, 0x05, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2c, 0x0a, 0x0f, 0x47,
	0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19,
	0x0a, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x42,
	0x6f, 0x6f, 0x6b, 0x52, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x2a, 0x33, 0x0a, 0x08, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52,
	0x59, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x43,
	0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x54, 0x45, 0x43, 0x48, 0x10, 0x01, 0x32, 0xa7,
	0x01, 0x0a, 0x10, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x2c, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x0f,
	0x2e, 0x41, 0x64, 0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x10, 0x2e, 0x41, 0x64, 0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x37, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x12, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x12, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6f, 0x6f,
	0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x07, 0x47, 0x65,
	0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x0f, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x61, 0x74, 0x74, 0x74, 0x74, 0x6f, 0x2f, 0x62,
	0x6f, 0x6f, 0x6b, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61,
	0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_bookshelf_proto_rawDescOnce sync.Once
	file_api_bookshelf_proto_rawDescData = file_api_bookshelf_proto_rawDesc
)

func file_api_bookshelf_proto_rawDescGZIP() []byte {
	file_api_bookshelf_proto_rawDescOnce.Do(func() {
		file_api_bookshelf_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_bookshelf_proto_rawDescData)
	})
	return file_api_bookshelf_proto_rawDescData
}

var file_api_bookshelf_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_bookshelf_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_bookshelf_proto_goTypes = []interface{}{
	(Category)(0),                 // 0: Category
	(*Book)(nil),                  // 1: Book
	(*AddBookRequest)(nil),        // 2: AddBookRequest
	(*AddBookResponse)(nil),       // 3: AddBookResponse
	(*ListBooksResponse)(nil),     // 4: ListBooksResponse
	(*GetBookRequest)(nil),        // 5: GetBookRequest
	(*GetBookResponse)(nil),       // 6: GetBookResponse
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 8: google.protobuf.Empty
}
var file_api_bookshelf_proto_depIdxs = []int32{
	0,  // 0: Book.category:type_name -> Category
	7,  // 1: Book.created_at:type_name -> google.protobuf.Timestamp
	7,  // 2: Book.updated_at:type_name -> google.protobuf.Timestamp
	7,  // 3: Book.deleted_at:type_name -> google.protobuf.Timestamp
	0,  // 4: AddBookRequest.category:type_name -> Category
	1,  // 5: AddBookResponse.book:type_name -> Book
	1,  // 6: ListBooksResponse.books:type_name -> Book
	1,  // 7: GetBookResponse.book:type_name -> Book
	2,  // 8: BookshelfService.AddBook:input_type -> AddBookRequest
	8,  // 9: BookshelfService.ListBooks:input_type -> google.protobuf.Empty
	5,  // 10: BookshelfService.GetBook:input_type -> GetBookRequest
	3,  // 11: BookshelfService.AddBook:output_type -> AddBookResponse
	4,  // 12: BookshelfService.ListBooks:output_type -> ListBooksResponse
	6,  // 13: BookshelfService.GetBook:output_type -> GetBookResponse
	11, // [11:14] is the sub-list for method output_type
	8,  // [8:11] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_api_bookshelf_proto_init() }
func file_api_bookshelf_proto_init() {
	if File_api_bookshelf_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_bookshelf_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Book); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_bookshelf_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddBookRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_bookshelf_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddBookResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_bookshelf_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBooksResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_bookshelf_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBookRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_bookshelf_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBookResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_bookshelf_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_bookshelf_proto_goTypes,
		DependencyIndexes: file_api_bookshelf_proto_depIdxs,
		EnumInfos:         file_api_bookshelf_proto_enumTypes,
		MessageInfos:      file_api_bookshelf_proto_msgTypes,
	}.Build()
	File_api_bookshelf_proto = out.File
	file_api_bookshelf_proto_rawDesc = nil
	file_api_bookshelf_proto_goTypes = nil
	file_api_bookshelf_proto_depIdxs = nil
}
