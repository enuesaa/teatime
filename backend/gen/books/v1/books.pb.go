// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: books/v1/books.proto

package books

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

type BooksServiceListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BooksServiceListRequest) Reset() {
	*x = BooksServiceListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_books_v1_books_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BooksServiceListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BooksServiceListRequest) ProtoMessage() {}

func (x *BooksServiceListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_books_v1_books_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BooksServiceListRequest.ProtoReflect.Descriptor instead.
func (*BooksServiceListRequest) Descriptor() ([]byte, []int) {
	return file_books_v1_books_proto_rawDescGZIP(), []int{0}
}

type BooksServiceListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *BooksServiceListResponse) Reset() {
	*x = BooksServiceListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_books_v1_books_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BooksServiceListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BooksServiceListResponse) ProtoMessage() {}

func (x *BooksServiceListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_books_v1_books_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BooksServiceListResponse.ProtoReflect.Descriptor instead.
func (*BooksServiceListResponse) Descriptor() ([]byte, []int) {
	return file_books_v1_books_proto_rawDescGZIP(), []int{1}
}

func (x *BooksServiceListResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_books_v1_books_proto protoreflect.FileDescriptor

var file_books_v1_books_proto_rawDesc = []byte{
	0x0a, 0x14, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x76, 0x31,
	0x22, 0x19, 0x0a, 0x17, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x34, 0x0a, 0x18, 0x42,
	0x6f, 0x6f, 0x6b, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x32, 0x5f, 0x0a, 0x0c, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x4f, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x21, 0x2e, 0x62, 0x6f, 0x6f, 0x6b,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x62,
	0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x65, 0x6e, 0x75, 0x65, 0x73, 0x61, 0x61, 0x2f, 0x74, 0x65, 0x61, 0x74, 0x69, 0x6d, 0x65,
	0x2d, 0x61, 0x70, 0x70, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x67, 0x65, 0x6e,
	0x3b, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_books_v1_books_proto_rawDescOnce sync.Once
	file_books_v1_books_proto_rawDescData = file_books_v1_books_proto_rawDesc
)

func file_books_v1_books_proto_rawDescGZIP() []byte {
	file_books_v1_books_proto_rawDescOnce.Do(func() {
		file_books_v1_books_proto_rawDescData = protoimpl.X.CompressGZIP(file_books_v1_books_proto_rawDescData)
	})
	return file_books_v1_books_proto_rawDescData
}

var file_books_v1_books_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_books_v1_books_proto_goTypes = []interface{}{
	(*BooksServiceListRequest)(nil),  // 0: books.v1.BooksServiceListRequest
	(*BooksServiceListResponse)(nil), // 1: books.v1.BooksServiceListResponse
}
var file_books_v1_books_proto_depIdxs = []int32{
	0, // 0: books.v1.BooksService.List:input_type -> books.v1.BooksServiceListRequest
	1, // 1: books.v1.BooksService.List:output_type -> books.v1.BooksServiceListResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_books_v1_books_proto_init() }
func file_books_v1_books_proto_init() {
	if File_books_v1_books_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_books_v1_books_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BooksServiceListRequest); i {
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
		file_books_v1_books_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BooksServiceListResponse); i {
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
			RawDescriptor: file_books_v1_books_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_books_v1_books_proto_goTypes,
		DependencyIndexes: file_books_v1_books_proto_depIdxs,
		MessageInfos:      file_books_v1_books_proto_msgTypes,
	}.Build()
	File_books_v1_books_proto = out.File
	file_books_v1_books_proto_rawDesc = nil
	file_books_v1_books_proto_goTypes = nil
	file_books_v1_books_proto_depIdxs = nil
}
