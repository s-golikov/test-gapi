// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: garwinapis/search/v1/common.proto

package search

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

// Represents configuration parameters for sorting search results.
type SortingConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the parameter that is used for sorting.
	Param string `protobuf:"bytes,1,opt,name=param,proto3" json:"param,omitempty"`
	// The sorting direction.
	Order string `protobuf:"bytes,2,opt,name=order,proto3" json:"order,omitempty"`
}

func (x *SortingConfig) Reset() {
	*x = SortingConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_garwinapis_search_v1_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SortingConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SortingConfig) ProtoMessage() {}

func (x *SortingConfig) ProtoReflect() protoreflect.Message {
	mi := &file_garwinapis_search_v1_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SortingConfig.ProtoReflect.Descriptor instead.
func (*SortingConfig) Descriptor() ([]byte, []int) {
	return file_garwinapis_search_v1_common_proto_rawDescGZIP(), []int{0}
}

func (x *SortingConfig) GetParam() string {
	if x != nil {
		return x.Param
	}
	return ""
}

func (x *SortingConfig) GetOrder() string {
	if x != nil {
		return x.Order
	}
	return ""
}

// Represents configuration parameters for paginating search results.
type PaginationConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Requested page number.
	PageNumber int32 `protobuf:"varint,1,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"`
	// Requested page size.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *PaginationConfig) Reset() {
	*x = PaginationConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_garwinapis_search_v1_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaginationConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaginationConfig) ProtoMessage() {}

func (x *PaginationConfig) ProtoReflect() protoreflect.Message {
	mi := &file_garwinapis_search_v1_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaginationConfig.ProtoReflect.Descriptor instead.
func (*PaginationConfig) Descriptor() ([]byte, []int) {
	return file_garwinapis_search_v1_common_proto_rawDescGZIP(), []int{1}
}

func (x *PaginationConfig) GetPageNumber() int32 {
	if x != nil {
		return x.PageNumber
	}
	return 0
}

func (x *PaginationConfig) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

var File_garwinapis_search_v1_common_proto protoreflect.FileDescriptor

var file_garwinapis_search_v1_common_proto_rawDesc = []byte{
	0x0a, 0x21, 0x67, 0x61, 0x72, 0x77, 0x69, 0x6e, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x73, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x14, 0x67, 0x61, 0x72, 0x77, 0x69, 0x6e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x22, 0x3b, 0x0a, 0x0d, 0x53, 0x6f, 0x72,
	0x74, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61,
	0x72, 0x61, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x12, 0x14, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x50, 0x0a, 0x10, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x42, 0x12, 0x5a, 0x10, 0x73, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_garwinapis_search_v1_common_proto_rawDescOnce sync.Once
	file_garwinapis_search_v1_common_proto_rawDescData = file_garwinapis_search_v1_common_proto_rawDesc
)

func file_garwinapis_search_v1_common_proto_rawDescGZIP() []byte {
	file_garwinapis_search_v1_common_proto_rawDescOnce.Do(func() {
		file_garwinapis_search_v1_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_garwinapis_search_v1_common_proto_rawDescData)
	})
	return file_garwinapis_search_v1_common_proto_rawDescData
}

var file_garwinapis_search_v1_common_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_garwinapis_search_v1_common_proto_goTypes = []interface{}{
	(*SortingConfig)(nil),    // 0: garwinapis.search.v1.SortingConfig
	(*PaginationConfig)(nil), // 1: garwinapis.search.v1.PaginationConfig
}
var file_garwinapis_search_v1_common_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_garwinapis_search_v1_common_proto_init() }
func file_garwinapis_search_v1_common_proto_init() {
	if File_garwinapis_search_v1_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_garwinapis_search_v1_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SortingConfig); i {
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
		file_garwinapis_search_v1_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaginationConfig); i {
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
			RawDescriptor: file_garwinapis_search_v1_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_garwinapis_search_v1_common_proto_goTypes,
		DependencyIndexes: file_garwinapis_search_v1_common_proto_depIdxs,
		MessageInfos:      file_garwinapis_search_v1_common_proto_msgTypes,
	}.Build()
	File_garwinapis_search_v1_common_proto = out.File
	file_garwinapis_search_v1_common_proto_rawDesc = nil
	file_garwinapis_search_v1_common_proto_goTypes = nil
	file_garwinapis_search_v1_common_proto_depIdxs = nil
}
