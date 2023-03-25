// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: v1/bookmark.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AddBookmarkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url  string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *AddBookmarkRequest) Reset() {
	*x = AddBookmarkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_bookmark_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddBookmarkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddBookmarkRequest) ProtoMessage() {}

func (x *AddBookmarkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_bookmark_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddBookmarkRequest.ProtoReflect.Descriptor instead.
func (*AddBookmarkRequest) Descriptor() ([]byte, []int) {
	return file_v1_bookmark_proto_rawDescGZIP(), []int{0}
}

func (x *AddBookmarkRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *AddBookmarkRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type AddBookmarResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AddBookmarResponse) Reset() {
	*x = AddBookmarResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_bookmark_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddBookmarResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddBookmarResponse) ProtoMessage() {}

func (x *AddBookmarResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_bookmark_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddBookmarResponse.ProtoReflect.Descriptor instead.
func (*AddBookmarResponse) Descriptor() ([]byte, []int) {
	return file_v1_bookmark_proto_rawDescGZIP(), []int{1}
}

func (x *AddBookmarResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ListBookmarkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *ListBookmarkRequest) Reset() {
	*x = ListBookmarkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_bookmark_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBookmarkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBookmarkRequest) ProtoMessage() {}

func (x *ListBookmarkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_bookmark_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBookmarkRequest.ProtoReflect.Descriptor instead.
func (*ListBookmarkRequest) Descriptor() ([]byte, []int) {
	return file_v1_bookmark_proto_rawDescGZIP(), []int{2}
}

func (x *ListBookmarkRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

type ListBookmarkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page  int32                        `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Items []*ListBookmarkResponse_Item `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ListBookmarkResponse) Reset() {
	*x = ListBookmarkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_bookmark_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBookmarkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBookmarkResponse) ProtoMessage() {}

func (x *ListBookmarkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_bookmark_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBookmarkResponse.ProtoReflect.Descriptor instead.
func (*ListBookmarkResponse) Descriptor() ([]byte, []int) {
	return file_v1_bookmark_proto_rawDescGZIP(), []int{3}
}

func (x *ListBookmarkResponse) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListBookmarkResponse) GetItems() []*ListBookmarkResponse_Item {
	if x != nil {
		return x.Items
	}
	return nil
}

type GetBookmarkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetBookmarkRequest) Reset() {
	*x = GetBookmarkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_bookmark_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBookmarkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBookmarkRequest) ProtoMessage() {}

func (x *GetBookmarkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_bookmark_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBookmarkRequest.ProtoReflect.Descriptor instead.
func (*GetBookmarkRequest) Descriptor() ([]byte, []int) {
	return file_v1_bookmark_proto_rawDescGZIP(), []int{4}
}

func (x *GetBookmarkRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetBookmarkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Url  string `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *GetBookmarkResponse) Reset() {
	*x = GetBookmarkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_bookmark_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBookmarkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBookmarkResponse) ProtoMessage() {}

func (x *GetBookmarkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_bookmark_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBookmarkResponse.ProtoReflect.Descriptor instead.
func (*GetBookmarkResponse) Descriptor() ([]byte, []int) {
	return file_v1_bookmark_proto_rawDescGZIP(), []int{5}
}

func (x *GetBookmarkResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetBookmarkResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetBookmarkResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type UpdateBookmarkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UpdateBookmarkRequest) Reset() {
	*x = UpdateBookmarkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_bookmark_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBookmarkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBookmarkRequest) ProtoMessage() {}

func (x *UpdateBookmarkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_bookmark_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBookmarkRequest.ProtoReflect.Descriptor instead.
func (*UpdateBookmarkRequest) Descriptor() ([]byte, []int) {
	return file_v1_bookmark_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateBookmarkRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpdateBookmarkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateBookmarkResponse) Reset() {
	*x = UpdateBookmarkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_bookmark_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBookmarkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBookmarkResponse) ProtoMessage() {}

func (x *UpdateBookmarkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_bookmark_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBookmarkResponse.ProtoReflect.Descriptor instead.
func (*UpdateBookmarkResponse) Descriptor() ([]byte, []int) {
	return file_v1_bookmark_proto_rawDescGZIP(), []int{7}
}

type RemoveBookmarkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RemoveBookmarkRequest) Reset() {
	*x = RemoveBookmarkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_bookmark_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveBookmarkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveBookmarkRequest) ProtoMessage() {}

func (x *RemoveBookmarkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_bookmark_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveBookmarkRequest.ProtoReflect.Descriptor instead.
func (*RemoveBookmarkRequest) Descriptor() ([]byte, []int) {
	return file_v1_bookmark_proto_rawDescGZIP(), []int{8}
}

func (x *RemoveBookmarkRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type RemoveBookmarkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RemoveBookmarkResponse) Reset() {
	*x = RemoveBookmarkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_bookmark_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveBookmarkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveBookmarkResponse) ProtoMessage() {}

func (x *RemoveBookmarkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_bookmark_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveBookmarkResponse.ProtoReflect.Descriptor instead.
func (*RemoveBookmarkResponse) Descriptor() ([]byte, []int) {
	return file_v1_bookmark_proto_rawDescGZIP(), []int{9}
}

type ListBookmarkResponse_Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Url  string `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *ListBookmarkResponse_Item) Reset() {
	*x = ListBookmarkResponse_Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_bookmark_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBookmarkResponse_Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBookmarkResponse_Item) ProtoMessage() {}

func (x *ListBookmarkResponse_Item) ProtoReflect() protoreflect.Message {
	mi := &file_v1_bookmark_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBookmarkResponse_Item.ProtoReflect.Descriptor instead.
func (*ListBookmarkResponse_Item) Descriptor() ([]byte, []int) {
	return file_v1_bookmark_proto_rawDescGZIP(), []int{3, 0}
}

func (x *ListBookmarkResponse_Item) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ListBookmarkResponse_Item) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ListBookmarkResponse_Item) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

var File_v1_bookmark_proto protoreflect.FileDescriptor

var file_v1_bookmark_proto_rawDesc = []byte{
	0x0a, 0x11, 0x76, 0x31, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x22, 0x3a, 0x0a, 0x12, 0x41, 0x64, 0x64, 0x42, 0x6f,
	0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x24, 0x0a, 0x12, 0x41, 0x64, 0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x29, 0x0a, 0x13, 0x4c, 0x69, 0x73,
	0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x22, 0x9d, 0x01, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6f, 0x6f,
	0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x12, 0x33, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61,
	0x72, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x1a, 0x3c, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x75, 0x72, 0x6c, 0x22, 0x24, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x6d,
	0x61, 0x72, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4b, 0x0a, 0x13, 0x47, 0x65,
	0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x27, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x18, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61,
	0x72, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x27, 0x0a, 0x15, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x18, 0x0a, 0x16, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x42, 0x6f, 0x6f,
	0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xe8, 0x02,
	0x0a, 0x08, 0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x3f, 0x0a, 0x0b, 0x41, 0x64,
	0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x16, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x64, 0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0c, 0x4c,
	0x69, 0x73, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x17, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6f,
	0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x40, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x12,
	0x16, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x49, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b,
	0x6d, 0x61, 0x72, 0x6b, 0x12, 0x19, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1a, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x6d,
	0x61, 0x72, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x49, 0x0a,
	0x0e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x12,
	0x19, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x6d,
	0x61, 0x72, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x72, 0x0a, 0x06, 0x63, 0x6f, 0x6d, 0x2e,
	0x76, 0x31, 0x42, 0x0d, 0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x65, 0x6e, 0x75, 0x65, 0x73, 0x61, 0x61, 0x2f, 0x74, 0x65, 0x61, 0x74, 0x69, 0x6d, 0x65, 0x2d,
	0x61, 0x70, 0x70, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x62, 0x75, 0x66, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x56, 0x58, 0x58, 0xaa, 0x02, 0x02, 0x56,
	0x31, 0xca, 0x02, 0x02, 0x56, 0x31, 0xe2, 0x02, 0x0e, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x02, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_bookmark_proto_rawDescOnce sync.Once
	file_v1_bookmark_proto_rawDescData = file_v1_bookmark_proto_rawDesc
)

func file_v1_bookmark_proto_rawDescGZIP() []byte {
	file_v1_bookmark_proto_rawDescOnce.Do(func() {
		file_v1_bookmark_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_bookmark_proto_rawDescData)
	})
	return file_v1_bookmark_proto_rawDescData
}

var file_v1_bookmark_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_v1_bookmark_proto_goTypes = []interface{}{
	(*AddBookmarkRequest)(nil),        // 0: v1.AddBookmarkRequest
	(*AddBookmarResponse)(nil),        // 1: v1.AddBookmarResponse
	(*ListBookmarkRequest)(nil),       // 2: v1.ListBookmarkRequest
	(*ListBookmarkResponse)(nil),      // 3: v1.ListBookmarkResponse
	(*GetBookmarkRequest)(nil),        // 4: v1.GetBookmarkRequest
	(*GetBookmarkResponse)(nil),       // 5: v1.GetBookmarkResponse
	(*UpdateBookmarkRequest)(nil),     // 6: v1.UpdateBookmarkRequest
	(*UpdateBookmarkResponse)(nil),    // 7: v1.UpdateBookmarkResponse
	(*RemoveBookmarkRequest)(nil),     // 8: v1.RemoveBookmarkRequest
	(*RemoveBookmarkResponse)(nil),    // 9: v1.RemoveBookmarkResponse
	(*ListBookmarkResponse_Item)(nil), // 10: v1.ListBookmarkResponse.Item
}
var file_v1_bookmark_proto_depIdxs = []int32{
	10, // 0: v1.ListBookmarkResponse.items:type_name -> v1.ListBookmarkResponse.Item
	0,  // 1: v1.Bookmark.AddBookmark:input_type -> v1.AddBookmarkRequest
	2,  // 2: v1.Bookmark.ListBookmark:input_type -> v1.ListBookmarkRequest
	4,  // 3: v1.Bookmark.GetBookmark:input_type -> v1.GetBookmarkRequest
	6,  // 4: v1.Bookmark.UpdateBookmark:input_type -> v1.UpdateBookmarkRequest
	8,  // 5: v1.Bookmark.RemoveBookmark:input_type -> v1.RemoveBookmarkRequest
	1,  // 6: v1.Bookmark.AddBookmark:output_type -> v1.AddBookmarResponse
	3,  // 7: v1.Bookmark.ListBookmark:output_type -> v1.ListBookmarkResponse
	5,  // 8: v1.Bookmark.GetBookmark:output_type -> v1.GetBookmarkResponse
	7,  // 9: v1.Bookmark.UpdateBookmark:output_type -> v1.UpdateBookmarkResponse
	9,  // 10: v1.Bookmark.RemoveBookmark:output_type -> v1.RemoveBookmarkResponse
	6,  // [6:11] is the sub-list for method output_type
	1,  // [1:6] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_v1_bookmark_proto_init() }
func file_v1_bookmark_proto_init() {
	if File_v1_bookmark_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_bookmark_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddBookmarkRequest); i {
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
		file_v1_bookmark_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddBookmarResponse); i {
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
		file_v1_bookmark_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBookmarkRequest); i {
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
		file_v1_bookmark_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBookmarkResponse); i {
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
		file_v1_bookmark_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBookmarkRequest); i {
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
		file_v1_bookmark_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBookmarkResponse); i {
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
		file_v1_bookmark_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBookmarkRequest); i {
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
		file_v1_bookmark_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBookmarkResponse); i {
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
		file_v1_bookmark_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveBookmarkRequest); i {
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
		file_v1_bookmark_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveBookmarkResponse); i {
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
		file_v1_bookmark_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBookmarkResponse_Item); i {
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
			RawDescriptor: file_v1_bookmark_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_bookmark_proto_goTypes,
		DependencyIndexes: file_v1_bookmark_proto_depIdxs,
		MessageInfos:      file_v1_bookmark_proto_msgTypes,
	}.Build()
	File_v1_bookmark_proto = out.File
	file_v1_bookmark_proto_rawDesc = nil
	file_v1_bookmark_proto_goTypes = nil
	file_v1_bookmark_proto_depIdxs = nil
}
