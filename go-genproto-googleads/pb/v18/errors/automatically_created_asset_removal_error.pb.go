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
// 	protoc-gen-go v1.35.2
// 	protoc        v4.25.3
// source: google/ads/googleads/v18/errors/automatically_created_asset_removal_error.proto

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

// Enum describing possible automatically created asset removal errors.
type AutomaticallyCreatedAssetRemovalErrorEnum_AutomaticallyCreatedAssetRemovalError int32

const (
	// Enum unspecified.
	AutomaticallyCreatedAssetRemovalErrorEnum_UNSPECIFIED AutomaticallyCreatedAssetRemovalErrorEnum_AutomaticallyCreatedAssetRemovalError = 0
	// The received error code is not known in this version.
	AutomaticallyCreatedAssetRemovalErrorEnum_UNKNOWN AutomaticallyCreatedAssetRemovalErrorEnum_AutomaticallyCreatedAssetRemovalError = 1
	// The ad does not exist.
	AutomaticallyCreatedAssetRemovalErrorEnum_AD_DOES_NOT_EXIST AutomaticallyCreatedAssetRemovalErrorEnum_AutomaticallyCreatedAssetRemovalError = 2
	// Ad type is not supported. Only Responsive Search Ad type is supported.
	AutomaticallyCreatedAssetRemovalErrorEnum_INVALID_AD_TYPE AutomaticallyCreatedAssetRemovalErrorEnum_AutomaticallyCreatedAssetRemovalError = 3
	// The asset does not exist.
	AutomaticallyCreatedAssetRemovalErrorEnum_ASSET_DOES_NOT_EXIST AutomaticallyCreatedAssetRemovalErrorEnum_AutomaticallyCreatedAssetRemovalError = 4
	// The asset field type does not match.
	AutomaticallyCreatedAssetRemovalErrorEnum_ASSET_FIELD_TYPE_DOES_NOT_MATCH AutomaticallyCreatedAssetRemovalErrorEnum_AutomaticallyCreatedAssetRemovalError = 5
	// Not an automatically created asset.
	AutomaticallyCreatedAssetRemovalErrorEnum_NOT_AN_AUTOMATICALLY_CREATED_ASSET AutomaticallyCreatedAssetRemovalErrorEnum_AutomaticallyCreatedAssetRemovalError = 6
)

// Enum value maps for AutomaticallyCreatedAssetRemovalErrorEnum_AutomaticallyCreatedAssetRemovalError.
var (
	AutomaticallyCreatedAssetRemovalErrorEnum_AutomaticallyCreatedAssetRemovalError_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "AD_DOES_NOT_EXIST",
		3: "INVALID_AD_TYPE",
		4: "ASSET_DOES_NOT_EXIST",
		5: "ASSET_FIELD_TYPE_DOES_NOT_MATCH",
		6: "NOT_AN_AUTOMATICALLY_CREATED_ASSET",
	}
	AutomaticallyCreatedAssetRemovalErrorEnum_AutomaticallyCreatedAssetRemovalError_value = map[string]int32{
		"UNSPECIFIED":                        0,
		"UNKNOWN":                            1,
		"AD_DOES_NOT_EXIST":                  2,
		"INVALID_AD_TYPE":                    3,
		"ASSET_DOES_NOT_EXIST":               4,
		"ASSET_FIELD_TYPE_DOES_NOT_MATCH":    5,
		"NOT_AN_AUTOMATICALLY_CREATED_ASSET": 6,
	}
)

func (x AutomaticallyCreatedAssetRemovalErrorEnum_AutomaticallyCreatedAssetRemovalError) Enum() *AutomaticallyCreatedAssetRemovalErrorEnum_AutomaticallyCreatedAssetRemovalError {
	p := new(AutomaticallyCreatedAssetRemovalErrorEnum_AutomaticallyCreatedAssetRemovalError)
	*p = x
	return p
}

func (x AutomaticallyCreatedAssetRemovalErrorEnum_AutomaticallyCreatedAssetRemovalError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AutomaticallyCreatedAssetRemovalErrorEnum_AutomaticallyCreatedAssetRemovalError) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v18_errors_automatically_created_asset_removal_error_proto_enumTypes[0].Descriptor()
}

func (AutomaticallyCreatedAssetRemovalErrorEnum_AutomaticallyCreatedAssetRemovalError) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v18_errors_automatically_created_asset_removal_error_proto_enumTypes[0]
}

func (x AutomaticallyCreatedAssetRemovalErrorEnum_AutomaticallyCreatedAssetRemovalError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AutomaticallyCreatedAssetRemovalErrorEnum_AutomaticallyCreatedAssetRemovalError.Descriptor instead.
func (AutomaticallyCreatedAssetRemovalErrorEnum_AutomaticallyCreatedAssetRemovalError) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v18_errors_automatically_created_asset_removal_error_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible automatically created asset removal
// errors.
type AutomaticallyCreatedAssetRemovalErrorEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AutomaticallyCreatedAssetRemovalErrorEnum) Reset() {
	*x = AutomaticallyCreatedAssetRemovalErrorEnum{}
	mi := &file_google_ads_googleads_v18_errors_automatically_created_asset_removal_error_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AutomaticallyCreatedAssetRemovalErrorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AutomaticallyCreatedAssetRemovalErrorEnum) ProtoMessage() {}

func (x *AutomaticallyCreatedAssetRemovalErrorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v18_errors_automatically_created_asset_removal_error_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AutomaticallyCreatedAssetRemovalErrorEnum.ProtoReflect.Descriptor instead.
func (*AutomaticallyCreatedAssetRemovalErrorEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v18_errors_automatically_created_asset_removal_error_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v18_errors_automatically_created_asset_removal_error_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v18_errors_automatically_created_asset_removal_error_proto_rawDesc = []byte{
	0x0a, 0x4f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x38, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x6c, 0x79, 0x5f,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x72, 0x65,
	0x6d, 0x6f, 0x76, 0x61, 0x6c, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x38, 0x2e, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0x22, 0x86, 0x02, 0x0a, 0x29, 0x41, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x69, 0x63,
	0x61, 0x6c, 0x6c, 0x79, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x73, 0x73, 0x65, 0x74,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x61, 0x6c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x45, 0x6e, 0x75, 0x6d,
	0x22, 0xd8, 0x01, 0x0a, 0x25, 0x41, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x61, 0x6c,
	0x6c, 0x79, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x73, 0x73, 0x65, 0x74, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x61, 0x6c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55,
	0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x41, 0x44, 0x5f, 0x44,
	0x4f, 0x45, 0x53, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x10, 0x02, 0x12,
	0x13, 0x0a, 0x0f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x41, 0x44, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x10, 0x03, 0x12, 0x18, 0x0a, 0x14, 0x41, 0x53, 0x53, 0x45, 0x54, 0x5f, 0x44, 0x4f,
	0x45, 0x53, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x10, 0x04, 0x12, 0x23,
	0x0a, 0x1f, 0x41, 0x53, 0x53, 0x45, 0x54, 0x5f, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x44, 0x4f, 0x45, 0x53, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x4d, 0x41, 0x54, 0x43,
	0x48, 0x10, 0x05, 0x12, 0x26, 0x0a, 0x22, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x4e, 0x5f, 0x41, 0x55,
	0x54, 0x4f, 0x4d, 0x41, 0x54, 0x49, 0x43, 0x41, 0x4c, 0x4c, 0x59, 0x5f, 0x43, 0x52, 0x45, 0x41,
	0x54, 0x45, 0x44, 0x5f, 0x41, 0x53, 0x53, 0x45, 0x54, 0x10, 0x06, 0x42, 0x8a, 0x02, 0x0a, 0x23,
	0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x38, 0x2e, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x73, 0x42, 0x2a, 0x41, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x61, 0x6c,
	0x6c, 0x79, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x73, 0x73, 0x65, 0x74, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x61, 0x6c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x45, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67,
	0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x38, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x3b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02,
	0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x38, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0xca, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x38, 0x5c, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0xea, 0x02, 0x23, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73,
	0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x38,
	0x3a, 0x3a, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v18_errors_automatically_created_asset_removal_error_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v18_errors_automatically_created_asset_removal_error_proto_rawDescData = file_google_ads_googleads_v18_errors_automatically_created_asset_removal_error_proto_rawDesc
)

func file_google_ads_googleads_v18_errors_automatically_created_asset_removal_error_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v18_errors_automatically_created_asset_removal_error_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v18_errors_automatically_created_asset_removal_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v18_errors_automatically_created_asset_removal_error_proto_rawDescData)
	})
	return file_google_ads_googleads_v18_errors_automatically_created_asset_removal_error_proto_rawDescData
}

var file_google_ads_googleads_v18_errors_automatically_created_asset_removal_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v18_errors_automatically_created_asset_removal_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v18_errors_automatically_created_asset_removal_error_proto_goTypes = []any{
	(AutomaticallyCreatedAssetRemovalErrorEnum_AutomaticallyCreatedAssetRemovalError)(0), // 0: google.ads.googleads.v18.errors.AutomaticallyCreatedAssetRemovalErrorEnum.AutomaticallyCreatedAssetRemovalError
	(*AutomaticallyCreatedAssetRemovalErrorEnum)(nil),                                    // 1: google.ads.googleads.v18.errors.AutomaticallyCreatedAssetRemovalErrorEnum
}
var file_google_ads_googleads_v18_errors_automatically_created_asset_removal_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_google_ads_googleads_v18_errors_automatically_created_asset_removal_error_proto_init()
}
func file_google_ads_googleads_v18_errors_automatically_created_asset_removal_error_proto_init() {
	if File_google_ads_googleads_v18_errors_automatically_created_asset_removal_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v18_errors_automatically_created_asset_removal_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v18_errors_automatically_created_asset_removal_error_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v18_errors_automatically_created_asset_removal_error_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v18_errors_automatically_created_asset_removal_error_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v18_errors_automatically_created_asset_removal_error_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v18_errors_automatically_created_asset_removal_error_proto = out.File
	file_google_ads_googleads_v18_errors_automatically_created_asset_removal_error_proto_rawDesc = nil
	file_google_ads_googleads_v18_errors_automatically_created_asset_removal_error_proto_goTypes = nil
	file_google_ads_googleads_v18_errors_automatically_created_asset_removal_error_proto_depIdxs = nil
}
