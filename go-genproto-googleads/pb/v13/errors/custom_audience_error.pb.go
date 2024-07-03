// Copyright 2023 Google LLC
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
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.2
// source: google/ads/googleads/v13/errors/custom_audience_error.proto

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

// Enum describing possible custom audience errors.
type CustomAudienceErrorEnum_CustomAudienceError int32

const (
	// Enum unspecified.
	CustomAudienceErrorEnum_UNSPECIFIED CustomAudienceErrorEnum_CustomAudienceError = 0
	// The received error code is not known in this version.
	CustomAudienceErrorEnum_UNKNOWN CustomAudienceErrorEnum_CustomAudienceError = 1
	// New name in the custom audience is duplicated ignoring cases.
	CustomAudienceErrorEnum_NAME_ALREADY_USED CustomAudienceErrorEnum_CustomAudienceError = 2
	// Cannot remove a custom audience while it's still being used as targeting.
	CustomAudienceErrorEnum_CANNOT_REMOVE_WHILE_IN_USE CustomAudienceErrorEnum_CustomAudienceError = 3
	// Cannot update or remove a custom audience that is already removed.
	CustomAudienceErrorEnum_RESOURCE_ALREADY_REMOVED CustomAudienceErrorEnum_CustomAudienceError = 4
	// The pair of [type, value] already exists in members.
	CustomAudienceErrorEnum_MEMBER_TYPE_AND_PARAMETER_ALREADY_EXISTED CustomAudienceErrorEnum_CustomAudienceError = 5
	// Member type is invalid.
	CustomAudienceErrorEnum_INVALID_MEMBER_TYPE CustomAudienceErrorEnum_CustomAudienceError = 6
	// Member type does not have associated value.
	CustomAudienceErrorEnum_MEMBER_TYPE_AND_VALUE_DOES_NOT_MATCH CustomAudienceErrorEnum_CustomAudienceError = 7
	// Custom audience contains a member that violates policy.
	CustomAudienceErrorEnum_POLICY_VIOLATION CustomAudienceErrorEnum_CustomAudienceError = 8
	// Change in custom audience type is not allowed.
	CustomAudienceErrorEnum_INVALID_TYPE_CHANGE CustomAudienceErrorEnum_CustomAudienceError = 9
)

// Enum value maps for CustomAudienceErrorEnum_CustomAudienceError.
var (
	CustomAudienceErrorEnum_CustomAudienceError_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "NAME_ALREADY_USED",
		3: "CANNOT_REMOVE_WHILE_IN_USE",
		4: "RESOURCE_ALREADY_REMOVED",
		5: "MEMBER_TYPE_AND_PARAMETER_ALREADY_EXISTED",
		6: "INVALID_MEMBER_TYPE",
		7: "MEMBER_TYPE_AND_VALUE_DOES_NOT_MATCH",
		8: "POLICY_VIOLATION",
		9: "INVALID_TYPE_CHANGE",
	}
	CustomAudienceErrorEnum_CustomAudienceError_value = map[string]int32{
		"UNSPECIFIED":                               0,
		"UNKNOWN":                                   1,
		"NAME_ALREADY_USED":                         2,
		"CANNOT_REMOVE_WHILE_IN_USE":                3,
		"RESOURCE_ALREADY_REMOVED":                  4,
		"MEMBER_TYPE_AND_PARAMETER_ALREADY_EXISTED": 5,
		"INVALID_MEMBER_TYPE":                       6,
		"MEMBER_TYPE_AND_VALUE_DOES_NOT_MATCH":      7,
		"POLICY_VIOLATION":                          8,
		"INVALID_TYPE_CHANGE":                       9,
	}
)

func (x CustomAudienceErrorEnum_CustomAudienceError) Enum() *CustomAudienceErrorEnum_CustomAudienceError {
	p := new(CustomAudienceErrorEnum_CustomAudienceError)
	*p = x
	return p
}

func (x CustomAudienceErrorEnum_CustomAudienceError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CustomAudienceErrorEnum_CustomAudienceError) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v13_errors_custom_audience_error_proto_enumTypes[0].Descriptor()
}

func (CustomAudienceErrorEnum_CustomAudienceError) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v13_errors_custom_audience_error_proto_enumTypes[0]
}

func (x CustomAudienceErrorEnum_CustomAudienceError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CustomAudienceErrorEnum_CustomAudienceError.Descriptor instead.
func (CustomAudienceErrorEnum_CustomAudienceError) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v13_errors_custom_audience_error_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible custom audience errors.
type CustomAudienceErrorEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CustomAudienceErrorEnum) Reset() {
	*x = CustomAudienceErrorEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v13_errors_custom_audience_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomAudienceErrorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomAudienceErrorEnum) ProtoMessage() {}

func (x *CustomAudienceErrorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v13_errors_custom_audience_error_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomAudienceErrorEnum.ProtoReflect.Descriptor instead.
func (*CustomAudienceErrorEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v13_errors_custom_audience_error_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v13_errors_custom_audience_error_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v13_errors_custom_audience_error_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x33, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x61, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63,
	0x65, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x33, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x22, 0xc5,
	0x02, 0x0a, 0x17, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x41, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63,
	0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0xa9, 0x02, 0x0a, 0x13, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x41, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01,
	0x12, 0x15, 0x0a, 0x11, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x41, 0x4c, 0x52, 0x45, 0x41, 0x44, 0x59,
	0x5f, 0x55, 0x53, 0x45, 0x44, 0x10, 0x02, 0x12, 0x1e, 0x0a, 0x1a, 0x43, 0x41, 0x4e, 0x4e, 0x4f,
	0x54, 0x5f, 0x52, 0x45, 0x4d, 0x4f, 0x56, 0x45, 0x5f, 0x57, 0x48, 0x49, 0x4c, 0x45, 0x5f, 0x49,
	0x4e, 0x5f, 0x55, 0x53, 0x45, 0x10, 0x03, 0x12, 0x1c, 0x0a, 0x18, 0x52, 0x45, 0x53, 0x4f, 0x55,
	0x52, 0x43, 0x45, 0x5f, 0x41, 0x4c, 0x52, 0x45, 0x41, 0x44, 0x59, 0x5f, 0x52, 0x45, 0x4d, 0x4f,
	0x56, 0x45, 0x44, 0x10, 0x04, 0x12, 0x2d, 0x0a, 0x29, 0x4d, 0x45, 0x4d, 0x42, 0x45, 0x52, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x4e, 0x44, 0x5f, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x45, 0x54,
	0x45, 0x52, 0x5f, 0x41, 0x4c, 0x52, 0x45, 0x41, 0x44, 0x59, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54,
	0x45, 0x44, 0x10, 0x05, 0x12, 0x17, 0x0a, 0x13, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f,
	0x4d, 0x45, 0x4d, 0x42, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x06, 0x12, 0x28, 0x0a,
	0x24, 0x4d, 0x45, 0x4d, 0x42, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x4e, 0x44,
	0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x44, 0x4f, 0x45, 0x53, 0x5f, 0x4e, 0x4f, 0x54, 0x5f,
	0x4d, 0x41, 0x54, 0x43, 0x48, 0x10, 0x07, 0x12, 0x14, 0x0a, 0x10, 0x50, 0x4f, 0x4c, 0x49, 0x43,
	0x59, 0x5f, 0x56, 0x49, 0x4f, 0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x08, 0x12, 0x17, 0x0a,
	0x13, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x48,
	0x41, 0x4e, 0x47, 0x45, 0x10, 0x09, 0x42, 0xf8, 0x01, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x33, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x42, 0x18,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x41, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f,
	0x76, 0x31, 0x33, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x3b, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56,
	0x31, 0x33, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xca, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73,
	0x5c, 0x56, 0x31, 0x33, 0x5c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xea, 0x02, 0x23, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x33, 0x3a, 0x3a, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v13_errors_custom_audience_error_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v13_errors_custom_audience_error_proto_rawDescData = file_google_ads_googleads_v13_errors_custom_audience_error_proto_rawDesc
)

func file_google_ads_googleads_v13_errors_custom_audience_error_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v13_errors_custom_audience_error_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v13_errors_custom_audience_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v13_errors_custom_audience_error_proto_rawDescData)
	})
	return file_google_ads_googleads_v13_errors_custom_audience_error_proto_rawDescData
}

var file_google_ads_googleads_v13_errors_custom_audience_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v13_errors_custom_audience_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v13_errors_custom_audience_error_proto_goTypes = []interface{}{
	(CustomAudienceErrorEnum_CustomAudienceError)(0), // 0: google.ads.googleads.v13.errors.CustomAudienceErrorEnum.CustomAudienceError
	(*CustomAudienceErrorEnum)(nil),                  // 1: google.ads.googleads.v13.errors.CustomAudienceErrorEnum
}
var file_google_ads_googleads_v13_errors_custom_audience_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v13_errors_custom_audience_error_proto_init() }
func file_google_ads_googleads_v13_errors_custom_audience_error_proto_init() {
	if File_google_ads_googleads_v13_errors_custom_audience_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v13_errors_custom_audience_error_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomAudienceErrorEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v13_errors_custom_audience_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v13_errors_custom_audience_error_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v13_errors_custom_audience_error_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v13_errors_custom_audience_error_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v13_errors_custom_audience_error_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v13_errors_custom_audience_error_proto = out.File
	file_google_ads_googleads_v13_errors_custom_audience_error_proto_rawDesc = nil
	file_google_ads_googleads_v13_errors_custom_audience_error_proto_goTypes = nil
	file_google_ads_googleads_v13_errors_custom_audience_error_proto_depIdxs = nil
}
