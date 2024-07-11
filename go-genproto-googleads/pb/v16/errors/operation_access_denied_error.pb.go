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
// 	protoc-gen-go v1.34.2
// 	protoc        v4.25.3
// source: google/ads/googleads/v16/errors/operation_access_denied_error.proto

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

// Enum describing possible operation access denied errors.
type OperationAccessDeniedErrorEnum_OperationAccessDeniedError int32

const (
	// Enum unspecified.
	OperationAccessDeniedErrorEnum_UNSPECIFIED OperationAccessDeniedErrorEnum_OperationAccessDeniedError = 0
	// The received error code is not known in this version.
	OperationAccessDeniedErrorEnum_UNKNOWN OperationAccessDeniedErrorEnum_OperationAccessDeniedError = 1
	// Unauthorized invocation of a service's method (get, mutate, etc.)
	OperationAccessDeniedErrorEnum_ACTION_NOT_PERMITTED OperationAccessDeniedErrorEnum_OperationAccessDeniedError = 2
	// Unauthorized CREATE operation in invoking a service's mutate method.
	OperationAccessDeniedErrorEnum_CREATE_OPERATION_NOT_PERMITTED OperationAccessDeniedErrorEnum_OperationAccessDeniedError = 3
	// Unauthorized REMOVE operation in invoking a service's mutate method.
	OperationAccessDeniedErrorEnum_REMOVE_OPERATION_NOT_PERMITTED OperationAccessDeniedErrorEnum_OperationAccessDeniedError = 4
	// Unauthorized UPDATE operation in invoking a service's mutate method.
	OperationAccessDeniedErrorEnum_UPDATE_OPERATION_NOT_PERMITTED OperationAccessDeniedErrorEnum_OperationAccessDeniedError = 5
	// A mutate action is not allowed on this resource, from this client.
	OperationAccessDeniedErrorEnum_MUTATE_ACTION_NOT_PERMITTED_FOR_CLIENT OperationAccessDeniedErrorEnum_OperationAccessDeniedError = 6
	// This operation is not permitted on this campaign type
	OperationAccessDeniedErrorEnum_OPERATION_NOT_PERMITTED_FOR_CAMPAIGN_TYPE OperationAccessDeniedErrorEnum_OperationAccessDeniedError = 7
	// A CREATE operation may not set status to REMOVED.
	OperationAccessDeniedErrorEnum_CREATE_AS_REMOVED_NOT_PERMITTED OperationAccessDeniedErrorEnum_OperationAccessDeniedError = 8
	// This operation is not allowed because the resource is removed.
	OperationAccessDeniedErrorEnum_OPERATION_NOT_PERMITTED_FOR_REMOVED_RESOURCE OperationAccessDeniedErrorEnum_OperationAccessDeniedError = 9
	// This operation is not permitted on this ad group type.
	OperationAccessDeniedErrorEnum_OPERATION_NOT_PERMITTED_FOR_AD_GROUP_TYPE OperationAccessDeniedErrorEnum_OperationAccessDeniedError = 10
	// The mutate is not allowed for this customer.
	OperationAccessDeniedErrorEnum_MUTATE_NOT_PERMITTED_FOR_CUSTOMER OperationAccessDeniedErrorEnum_OperationAccessDeniedError = 11
)

// Enum value maps for OperationAccessDeniedErrorEnum_OperationAccessDeniedError.
var (
	OperationAccessDeniedErrorEnum_OperationAccessDeniedError_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "ACTION_NOT_PERMITTED",
		3:  "CREATE_OPERATION_NOT_PERMITTED",
		4:  "REMOVE_OPERATION_NOT_PERMITTED",
		5:  "UPDATE_OPERATION_NOT_PERMITTED",
		6:  "MUTATE_ACTION_NOT_PERMITTED_FOR_CLIENT",
		7:  "OPERATION_NOT_PERMITTED_FOR_CAMPAIGN_TYPE",
		8:  "CREATE_AS_REMOVED_NOT_PERMITTED",
		9:  "OPERATION_NOT_PERMITTED_FOR_REMOVED_RESOURCE",
		10: "OPERATION_NOT_PERMITTED_FOR_AD_GROUP_TYPE",
		11: "MUTATE_NOT_PERMITTED_FOR_CUSTOMER",
	}
	OperationAccessDeniedErrorEnum_OperationAccessDeniedError_value = map[string]int32{
		"UNSPECIFIED":                                  0,
		"UNKNOWN":                                      1,
		"ACTION_NOT_PERMITTED":                         2,
		"CREATE_OPERATION_NOT_PERMITTED":               3,
		"REMOVE_OPERATION_NOT_PERMITTED":               4,
		"UPDATE_OPERATION_NOT_PERMITTED":               5,
		"MUTATE_ACTION_NOT_PERMITTED_FOR_CLIENT":       6,
		"OPERATION_NOT_PERMITTED_FOR_CAMPAIGN_TYPE":    7,
		"CREATE_AS_REMOVED_NOT_PERMITTED":              8,
		"OPERATION_NOT_PERMITTED_FOR_REMOVED_RESOURCE": 9,
		"OPERATION_NOT_PERMITTED_FOR_AD_GROUP_TYPE":    10,
		"MUTATE_NOT_PERMITTED_FOR_CUSTOMER":            11,
	}
)

func (x OperationAccessDeniedErrorEnum_OperationAccessDeniedError) Enum() *OperationAccessDeniedErrorEnum_OperationAccessDeniedError {
	p := new(OperationAccessDeniedErrorEnum_OperationAccessDeniedError)
	*p = x
	return p
}

func (x OperationAccessDeniedErrorEnum_OperationAccessDeniedError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OperationAccessDeniedErrorEnum_OperationAccessDeniedError) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v16_errors_operation_access_denied_error_proto_enumTypes[0].Descriptor()
}

func (OperationAccessDeniedErrorEnum_OperationAccessDeniedError) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v16_errors_operation_access_denied_error_proto_enumTypes[0]
}

func (x OperationAccessDeniedErrorEnum_OperationAccessDeniedError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OperationAccessDeniedErrorEnum_OperationAccessDeniedError.Descriptor instead.
func (OperationAccessDeniedErrorEnum_OperationAccessDeniedError) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v16_errors_operation_access_denied_error_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible operation access denied errors.
type OperationAccessDeniedErrorEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OperationAccessDeniedErrorEnum) Reset() {
	*x = OperationAccessDeniedErrorEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v16_errors_operation_access_denied_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationAccessDeniedErrorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationAccessDeniedErrorEnum) ProtoMessage() {}

func (x *OperationAccessDeniedErrorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v16_errors_operation_access_denied_error_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationAccessDeniedErrorEnum.ProtoReflect.Descriptor instead.
func (*OperationAccessDeniedErrorEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v16_errors_operation_access_denied_error_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v16_errors_operation_access_denied_error_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v16_errors_operation_access_denied_error_proto_rawDesc = []byte{
	0x0a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x36, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x5f, 0x64, 0x65, 0x6e, 0x69, 0x65, 0x64, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x36, 0x2e,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x22, 0xeb, 0x03, 0x0a, 0x1e, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x44, 0x65, 0x6e, 0x69, 0x65, 0x64,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0xc8, 0x03, 0x0a, 0x1a, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x44, 0x65, 0x6e,
	0x69, 0x65, 0x64, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x54, 0x54, 0x45, 0x44, 0x10, 0x02,
	0x12, 0x22, 0x0a, 0x1e, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x54, 0x54,
	0x45, 0x44, 0x10, 0x03, 0x12, 0x22, 0x0a, 0x1e, 0x52, 0x45, 0x4d, 0x4f, 0x56, 0x45, 0x5f, 0x4f,
	0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x50, 0x45, 0x52,
	0x4d, 0x49, 0x54, 0x54, 0x45, 0x44, 0x10, 0x04, 0x12, 0x22, 0x0a, 0x1e, 0x55, 0x50, 0x44, 0x41,
	0x54, 0x45, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4e, 0x4f, 0x54,
	0x5f, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x54, 0x54, 0x45, 0x44, 0x10, 0x05, 0x12, 0x2a, 0x0a, 0x26,
	0x4d, 0x55, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4e, 0x4f,
	0x54, 0x5f, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x54, 0x54, 0x45, 0x44, 0x5f, 0x46, 0x4f, 0x52, 0x5f,
	0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x10, 0x06, 0x12, 0x2d, 0x0a, 0x29, 0x4f, 0x50, 0x45, 0x52,
	0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x54,
	0x54, 0x45, 0x44, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x43, 0x41, 0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x07, 0x12, 0x23, 0x0a, 0x1f, 0x43, 0x52, 0x45, 0x41, 0x54,
	0x45, 0x5f, 0x41, 0x53, 0x5f, 0x52, 0x45, 0x4d, 0x4f, 0x56, 0x45, 0x44, 0x5f, 0x4e, 0x4f, 0x54,
	0x5f, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x54, 0x54, 0x45, 0x44, 0x10, 0x08, 0x12, 0x30, 0x0a, 0x2c,
	0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x50, 0x45,
	0x52, 0x4d, 0x49, 0x54, 0x54, 0x45, 0x44, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x52, 0x45, 0x4d, 0x4f,
	0x56, 0x45, 0x44, 0x5f, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x10, 0x09, 0x12, 0x2d,
	0x0a, 0x29, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4e, 0x4f, 0x54, 0x5f,
	0x50, 0x45, 0x52, 0x4d, 0x49, 0x54, 0x54, 0x45, 0x44, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x41, 0x44,
	0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x0a, 0x12, 0x25, 0x0a,
	0x21, 0x4d, 0x55, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x50, 0x45, 0x52, 0x4d,
	0x49, 0x54, 0x54, 0x45, 0x44, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d,
	0x45, 0x52, 0x10, 0x0b, 0x42, 0xff, 0x01, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x76, 0x31, 0x36, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x42, 0x1f, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x44, 0x65, 0x6e,
	0x69, 0x65, 0x64, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x45, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f,
	0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x36, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x3b,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1f, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x36, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xca, 0x02,
	0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x36, 0x5c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0xea, 0x02, 0x23, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x36, 0x3a, 0x3a,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v16_errors_operation_access_denied_error_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v16_errors_operation_access_denied_error_proto_rawDescData = file_google_ads_googleads_v16_errors_operation_access_denied_error_proto_rawDesc
)

func file_google_ads_googleads_v16_errors_operation_access_denied_error_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v16_errors_operation_access_denied_error_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v16_errors_operation_access_denied_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v16_errors_operation_access_denied_error_proto_rawDescData)
	})
	return file_google_ads_googleads_v16_errors_operation_access_denied_error_proto_rawDescData
}

var file_google_ads_googleads_v16_errors_operation_access_denied_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v16_errors_operation_access_denied_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v16_errors_operation_access_denied_error_proto_goTypes = []any{
	(OperationAccessDeniedErrorEnum_OperationAccessDeniedError)(0), // 0: google.ads.googleads.v16.errors.OperationAccessDeniedErrorEnum.OperationAccessDeniedError
	(*OperationAccessDeniedErrorEnum)(nil),                         // 1: google.ads.googleads.v16.errors.OperationAccessDeniedErrorEnum
}
var file_google_ads_googleads_v16_errors_operation_access_denied_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v16_errors_operation_access_denied_error_proto_init() }
func file_google_ads_googleads_v16_errors_operation_access_denied_error_proto_init() {
	if File_google_ads_googleads_v16_errors_operation_access_denied_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v16_errors_operation_access_denied_error_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*OperationAccessDeniedErrorEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v16_errors_operation_access_denied_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v16_errors_operation_access_denied_error_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v16_errors_operation_access_denied_error_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v16_errors_operation_access_denied_error_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v16_errors_operation_access_denied_error_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v16_errors_operation_access_denied_error_proto = out.File
	file_google_ads_googleads_v16_errors_operation_access_denied_error_proto_rawDesc = nil
	file_google_ads_googleads_v16_errors_operation_access_denied_error_proto_goTypes = nil
	file_google_ads_googleads_v16_errors_operation_access_denied_error_proto_depIdxs = nil
}
