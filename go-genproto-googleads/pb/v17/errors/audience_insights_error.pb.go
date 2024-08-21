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
// source: google/ads/googleads/v17/errors/audience_insights_error.proto

// copy from https://github.com/dictav/go-genproto-googleads
// and changed by TeamMomentum

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

// Enum describing possible errors from AudienceInsightsService.
type AudienceInsightsErrorEnum_AudienceInsightsError int32

const (
	// Enum unspecified.
	AudienceInsightsErrorEnum_UNSPECIFIED AudienceInsightsErrorEnum_AudienceInsightsError = 0
	// The received error code is not known in this version.
	AudienceInsightsErrorEnum_UNKNOWN AudienceInsightsErrorEnum_AudienceInsightsError = 1
	// The dimensions cannot be used with topic audience combinations.
	AudienceInsightsErrorEnum_DIMENSION_INCOMPATIBLE_WITH_TOPIC_AUDIENCE_COMBINATIONS AudienceInsightsErrorEnum_AudienceInsightsError = 2
)

// Enum value maps for AudienceInsightsErrorEnum_AudienceInsightsError.
var (
	AudienceInsightsErrorEnum_AudienceInsightsError_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "DIMENSION_INCOMPATIBLE_WITH_TOPIC_AUDIENCE_COMBINATIONS",
	}
	AudienceInsightsErrorEnum_AudienceInsightsError_value = map[string]int32{
		"UNSPECIFIED": 0,
		"UNKNOWN":     1,
		"DIMENSION_INCOMPATIBLE_WITH_TOPIC_AUDIENCE_COMBINATIONS": 2,
	}
)

func (x AudienceInsightsErrorEnum_AudienceInsightsError) Enum() *AudienceInsightsErrorEnum_AudienceInsightsError {
	p := new(AudienceInsightsErrorEnum_AudienceInsightsError)
	*p = x
	return p
}

func (x AudienceInsightsErrorEnum_AudienceInsightsError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AudienceInsightsErrorEnum_AudienceInsightsError) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v17_errors_audience_insights_error_proto_enumTypes[0].Descriptor()
}

func (AudienceInsightsErrorEnum_AudienceInsightsError) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v17_errors_audience_insights_error_proto_enumTypes[0]
}

func (x AudienceInsightsErrorEnum_AudienceInsightsError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AudienceInsightsErrorEnum_AudienceInsightsError.Descriptor instead.
func (AudienceInsightsErrorEnum_AudienceInsightsError) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v17_errors_audience_insights_error_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible errors returned from
// the AudienceInsightsService.
type AudienceInsightsErrorEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AudienceInsightsErrorEnum) Reset() {
	*x = AudienceInsightsErrorEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v17_errors_audience_insights_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AudienceInsightsErrorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AudienceInsightsErrorEnum) ProtoMessage() {}

func (x *AudienceInsightsErrorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v17_errors_audience_insights_error_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AudienceInsightsErrorEnum.ProtoReflect.Descriptor instead.
func (*AudienceInsightsErrorEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v17_errors_audience_insights_error_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v17_errors_audience_insights_error_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v17_errors_audience_insights_error_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x37, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x6e, 0x73, 0x69, 0x67,
	0x68, 0x74, 0x73, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0x22, 0x8f, 0x01, 0x0a, 0x19, 0x41, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x6e, 0x73,
	0x69, 0x67, 0x68, 0x74, 0x73, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0x72,
	0x0a, 0x15, 0x41, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x6e, 0x73, 0x69, 0x67, 0x68,
	0x74, 0x73, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x3b, 0x0a, 0x37, 0x44, 0x49, 0x4d, 0x45, 0x4e, 0x53, 0x49,
	0x4f, 0x4e, 0x5f, 0x49, 0x4e, 0x43, 0x4f, 0x4d, 0x50, 0x41, 0x54, 0x49, 0x42, 0x4c, 0x45, 0x5f,
	0x57, 0x49, 0x54, 0x48, 0x5f, 0x54, 0x4f, 0x50, 0x49, 0x43, 0x5f, 0x41, 0x55, 0x44, 0x49, 0x45,
	0x4e, 0x43, 0x45, 0x5f, 0x43, 0x4f, 0x4d, 0x42, 0x49, 0x4e, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x53,
	0x10, 0x02, 0x42, 0xfa, 0x01, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x37, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x42, 0x1a, 0x41, 0x75, 0x64, 0x69,
	0x65, 0x6e, 0x63, 0x65, 0x49, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x73, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31,
	0x37, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x3b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xa2,
	0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41,
	0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x37,
	0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xca, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56,
	0x31, 0x37, 0x5c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xea, 0x02, 0x23, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41,
	0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x37, 0x3a, 0x3a, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v17_errors_audience_insights_error_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v17_errors_audience_insights_error_proto_rawDescData = file_google_ads_googleads_v17_errors_audience_insights_error_proto_rawDesc
)

func file_google_ads_googleads_v17_errors_audience_insights_error_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v17_errors_audience_insights_error_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v17_errors_audience_insights_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v17_errors_audience_insights_error_proto_rawDescData)
	})
	return file_google_ads_googleads_v17_errors_audience_insights_error_proto_rawDescData
}

var file_google_ads_googleads_v17_errors_audience_insights_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v17_errors_audience_insights_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v17_errors_audience_insights_error_proto_goTypes = []any{
	(AudienceInsightsErrorEnum_AudienceInsightsError)(0), // 0: google.ads.googleads.v17.errors.AudienceInsightsErrorEnum.AudienceInsightsError
	(*AudienceInsightsErrorEnum)(nil),                    // 1: google.ads.googleads.v17.errors.AudienceInsightsErrorEnum
}
var file_google_ads_googleads_v17_errors_audience_insights_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v17_errors_audience_insights_error_proto_init() }
func file_google_ads_googleads_v17_errors_audience_insights_error_proto_init() {
	if File_google_ads_googleads_v17_errors_audience_insights_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v17_errors_audience_insights_error_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*AudienceInsightsErrorEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v17_errors_audience_insights_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v17_errors_audience_insights_error_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v17_errors_audience_insights_error_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v17_errors_audience_insights_error_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v17_errors_audience_insights_error_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v17_errors_audience_insights_error_proto = out.File
	file_google_ads_googleads_v17_errors_audience_insights_error_proto_rawDesc = nil
	file_google_ads_googleads_v17_errors_audience_insights_error_proto_goTypes = nil
	file_google_ads_googleads_v17_errors_audience_insights_error_proto_depIdxs = nil
}
