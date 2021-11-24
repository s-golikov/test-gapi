// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: garwinapis/seo/v1/page_type.proto

package seo

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

// Represents the possible page types of a website.
type PageType int32

const (
	// The unspecified page.
	PageType_PAGE_TYPE_UNSPECIFIED PageType = 0
	// The home page.
	PageType_PAGE_TYPE_HOME PageType = 1
	// The product category page.
	PageType_PAGE_TYPE_CATEGORY PageType = 2
	// The product details page.
	PageType_PAGE_TYPE_PRODUCT PageType = 3
	// The promo details page.
	PageType_PAGE_TYPE_PROMO PageType = 4
	// The search results page.
	PageType_PAGE_TYPE_SEARCH_RESULTS PageType = 5
	// The cart page.
	PageType_PAGE_TYPE_CART PageType = 6
	// The checkout page.
	PageType_PAGE_TYPE_CHECKOUT PageType = 7
	// The personal account related pages.
	PageType_PAGE_TYPE_ACCOUNT PageType = 8
)

// Enum value maps for PageType.
var (
	PageType_name = map[int32]string{
		0: "PAGE_TYPE_UNSPECIFIED",
		1: "PAGE_TYPE_HOME",
		2: "PAGE_TYPE_CATEGORY",
		3: "PAGE_TYPE_PRODUCT",
		4: "PAGE_TYPE_PROMO",
		5: "PAGE_TYPE_SEARCH_RESULTS",
		6: "PAGE_TYPE_CART",
		7: "PAGE_TYPE_CHECKOUT",
		8: "PAGE_TYPE_ACCOUNT",
	}
	PageType_value = map[string]int32{
		"PAGE_TYPE_UNSPECIFIED":    0,
		"PAGE_TYPE_HOME":           1,
		"PAGE_TYPE_CATEGORY":       2,
		"PAGE_TYPE_PRODUCT":        3,
		"PAGE_TYPE_PROMO":          4,
		"PAGE_TYPE_SEARCH_RESULTS": 5,
		"PAGE_TYPE_CART":           6,
		"PAGE_TYPE_CHECKOUT":       7,
		"PAGE_TYPE_ACCOUNT":        8,
	}
)

func (x PageType) Enum() *PageType {
	p := new(PageType)
	*p = x
	return p
}

func (x PageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PageType) Descriptor() protoreflect.EnumDescriptor {
	return file_garwinapis_seo_v1_page_type_proto_enumTypes[0].Descriptor()
}

func (PageType) Type() protoreflect.EnumType {
	return &file_garwinapis_seo_v1_page_type_proto_enumTypes[0]
}

func (x PageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PageType.Descriptor instead.
func (PageType) EnumDescriptor() ([]byte, []int) {
	return file_garwinapis_seo_v1_page_type_proto_rawDescGZIP(), []int{0}
}

var File_garwinapis_seo_v1_page_type_proto protoreflect.FileDescriptor

var file_garwinapis_seo_v1_page_type_proto_rawDesc = []byte{
	0x0a, 0x21, 0x67, 0x61, 0x72, 0x77, 0x69, 0x6e, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x73, 0x65, 0x6f,
	0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x11, 0x67, 0x61, 0x72, 0x77, 0x69, 0x6e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x73, 0x65, 0x6f, 0x2e, 0x76, 0x31, 0x2a, 0xde, 0x01, 0x0a, 0x08, 0x50, 0x61, 0x67, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x15, 0x50, 0x41, 0x47, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x12,
	0x0a, 0x0e, 0x50, 0x41, 0x47, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x48, 0x4f, 0x4d, 0x45,
	0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x50, 0x41, 0x47, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x10, 0x02, 0x12, 0x15, 0x0a, 0x11, 0x50, 0x41,
	0x47, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43, 0x54, 0x10,
	0x03, 0x12, 0x13, 0x0a, 0x0f, 0x50, 0x41, 0x47, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50,
	0x52, 0x4f, 0x4d, 0x4f, 0x10, 0x04, 0x12, 0x1c, 0x0a, 0x18, 0x50, 0x41, 0x47, 0x45, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x53, 0x45, 0x41, 0x52, 0x43, 0x48, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c,
	0x54, 0x53, 0x10, 0x05, 0x12, 0x12, 0x0a, 0x0e, 0x50, 0x41, 0x47, 0x45, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x43, 0x41, 0x52, 0x54, 0x10, 0x06, 0x12, 0x16, 0x0a, 0x12, 0x50, 0x41, 0x47, 0x45,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x48, 0x45, 0x43, 0x4b, 0x4f, 0x55, 0x54, 0x10, 0x07,
	0x12, 0x15, 0x0a, 0x11, 0x50, 0x41, 0x47, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x43,
	0x43, 0x4f, 0x55, 0x4e, 0x54, 0x10, 0x08, 0x42, 0x0c, 0x5a, 0x0a, 0x73, 0x65, 0x6f, 0x2f, 0x76,
	0x31, 0x3b, 0x73, 0x65, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_garwinapis_seo_v1_page_type_proto_rawDescOnce sync.Once
	file_garwinapis_seo_v1_page_type_proto_rawDescData = file_garwinapis_seo_v1_page_type_proto_rawDesc
)

func file_garwinapis_seo_v1_page_type_proto_rawDescGZIP() []byte {
	file_garwinapis_seo_v1_page_type_proto_rawDescOnce.Do(func() {
		file_garwinapis_seo_v1_page_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_garwinapis_seo_v1_page_type_proto_rawDescData)
	})
	return file_garwinapis_seo_v1_page_type_proto_rawDescData
}

var file_garwinapis_seo_v1_page_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_garwinapis_seo_v1_page_type_proto_goTypes = []interface{}{
	(PageType)(0), // 0: garwinapis.seo.v1.PageType
}
var file_garwinapis_seo_v1_page_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_garwinapis_seo_v1_page_type_proto_init() }
func file_garwinapis_seo_v1_page_type_proto_init() {
	if File_garwinapis_seo_v1_page_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_garwinapis_seo_v1_page_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_garwinapis_seo_v1_page_type_proto_goTypes,
		DependencyIndexes: file_garwinapis_seo_v1_page_type_proto_depIdxs,
		EnumInfos:         file_garwinapis_seo_v1_page_type_proto_enumTypes,
	}.Build()
	File_garwinapis_seo_v1_page_type_proto = out.File
	file_garwinapis_seo_v1_page_type_proto_rawDesc = nil
	file_garwinapis_seo_v1_page_type_proto_goTypes = nil
	file_garwinapis_seo_v1_page_type_proto_depIdxs = nil
}
