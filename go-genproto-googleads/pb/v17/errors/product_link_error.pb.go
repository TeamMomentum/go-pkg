// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.25.3
// source: google/ads/googleads/v17/errors/product_link_error.proto

package errors

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

// Enum describing possible ProductLink errors.
type ProductLinkErrorEnum_ProductLinkError int32

const (
	// Enum unspecified.
	ProductLinkErrorEnum_UNSPECIFIED ProductLinkErrorEnum_ProductLinkError = 0
	// The received error code is not known in this version.
	ProductLinkErrorEnum_UNKNOWN ProductLinkErrorEnum_ProductLinkError = 1
	// The requested operation is invalid. For example, you are not allowed to
	// remove a link from a partner account.
	ProductLinkErrorEnum_INVALID_OPERATION ProductLinkErrorEnum_ProductLinkError = 2
	// The creation request is not permitted.
	ProductLinkErrorEnum_CREATION_NOT_PERMITTED ProductLinkErrorEnum_ProductLinkError = 3
	// A link cannot be created because a pending link already exists.
	ProductLinkErrorEnum_INVITATION_EXISTS ProductLinkErrorEnum_ProductLinkError = 4
	// A link cannot be created because an active link already exists.
	ProductLinkErrorEnum_LINK_EXISTS ProductLinkErrorEnum_ProductLinkError = 5
)

// Enum value maps for ProductLinkErrorEnum_ProductLinkError.
var (
	ProductLinkErrorEnum_ProductLinkError_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "INVALID_OPERATION",
		3: "CREATION_NOT_PERMITTED",
		4: "INVITATION_EXISTS",
		5: "LINK_EXISTS",
	}
	ProductLinkErrorEnum_ProductLinkError_value = map[string]int32{
		"UNSPECIFIED":            0,
		"UNKNOWN":                1,
		"INVALID_OPERATION":      2,
		"CREATION_NOT_PERMITTED": 3,
		"INVITATION_EXISTS":      4,
		"LINK_EXISTS":            5,
	}
)

func (x ProductLinkErrorEnum_ProductLinkError) Enum() *ProductLinkErrorEnum_ProductLinkError {
	p := new(ProductLinkErrorEnum_ProductLinkError)
	*p = x
	return p
}

func (x ProductLinkErrorEnum_ProductLinkError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProductLinkErrorEnum_ProductLinkError) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v17_errors_product_link_error_proto_enumTypes[0].Descriptor()
}

func (ProductLinkErrorEnum_ProductLinkError) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v17_errors_product_link_error_proto_enumTypes[0]
}

func (x ProductLinkErrorEnum_ProductLinkError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProductLinkErrorEnum_ProductLinkError.Descriptor instead.
func (ProductLinkErrorEnum_ProductLinkError) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v17_errors_product_link_error_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible ProductLink errors.
type ProductLinkErrorEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ProductLinkErrorEnum) Reset() {
	*x = ProductLinkErrorEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v17_errors_product_link_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductLinkErrorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductLinkErrorEnum) ProtoMessage() {}

func (x *ProductLinkErrorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v17_errors_product_link_error_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductLinkErrorEnum.ProtoReflect.Descriptor instead.
func (*ProductLinkErrorEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v17_errors_product_link_error_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v17_errors_product_link_error_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v17_errors_product_link_error_proto_rawDesc = []byte{
	0x0a, 0x38, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x37, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x31, 0x37, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x22, 0xa4, 0x01, 0x0a, 0x14,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x45, 0x6e, 0x75, 0x6d, 0x22, 0x8b, 0x01, 0x0a, 0x10, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x4c, 0x69, 0x6e, 0x6b, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e,
	0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x49, 0x4e, 0x56, 0x41, 0x4c,
	0x49, 0x44, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x02, 0x12, 0x1a,
	0x0a, 0x16, 0x43, 0x52, 0x45, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x50,
	0x45, 0x52, 0x4d, 0x49, 0x54, 0x54, 0x45, 0x44, 0x10, 0x03, 0x12, 0x15, 0x0a, 0x11, 0x49, 0x4e,
	0x56, 0x49, 0x54, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x53, 0x10,
	0x04, 0x12, 0x0f, 0x0a, 0x0b, 0x4c, 0x49, 0x4e, 0x4b, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x53,
	0x10, 0x05, 0x42, 0xf5, 0x01, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x37, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x42, 0x15, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x45, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x37, 0x2f, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x73, 0x3b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41,
	0xaa, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x37, 0x2e, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0xca, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x37, 0x5c, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0xea, 0x02, 0x23, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41,
	0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56,
	0x31, 0x37, 0x3a, 0x3a, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_google_ads_googleads_v17_errors_product_link_error_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v17_errors_product_link_error_proto_rawDescData = file_google_ads_googleads_v17_errors_product_link_error_proto_rawDesc
)

func file_google_ads_googleads_v17_errors_product_link_error_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v17_errors_product_link_error_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v17_errors_product_link_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v17_errors_product_link_error_proto_rawDescData)
	})
	return file_google_ads_googleads_v17_errors_product_link_error_proto_rawDescData
}

var file_google_ads_googleads_v17_errors_product_link_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v17_errors_product_link_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v17_errors_product_link_error_proto_goTypes = []any{
	(ProductLinkErrorEnum_ProductLinkError)(0), // 0: google.ads.googleads.v17.errors.ProductLinkErrorEnum.ProductLinkError
	(*ProductLinkErrorEnum)(nil),               // 1: google.ads.googleads.v17.errors.ProductLinkErrorEnum
}
var file_google_ads_googleads_v17_errors_product_link_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v17_errors_product_link_error_proto_init() }
func file_google_ads_googleads_v17_errors_product_link_error_proto_init() {
	if File_google_ads_googleads_v17_errors_product_link_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v17_errors_product_link_error_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ProductLinkErrorEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v17_errors_product_link_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v17_errors_product_link_error_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v17_errors_product_link_error_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v17_errors_product_link_error_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v17_errors_product_link_error_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v17_errors_product_link_error_proto = out.File
	file_google_ads_googleads_v17_errors_product_link_error_proto_rawDesc = nil
	file_google_ads_googleads_v17_errors_product_link_error_proto_goTypes = nil
	file_google_ads_googleads_v17_errors_product_link_error_proto_depIdxs = nil
}
