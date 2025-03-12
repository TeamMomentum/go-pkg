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
// source: google/ads/googleads/v18/enums/offline_user_data_job_failure_reason.proto

package enums

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

// The failure reason of an offline user data job.
type OfflineUserDataJobFailureReasonEnum_OfflineUserDataJobFailureReason int32

const (
	// Not specified.
	OfflineUserDataJobFailureReasonEnum_UNSPECIFIED OfflineUserDataJobFailureReasonEnum_OfflineUserDataJobFailureReason = 0
	// Used for return value only. Represents value unknown in this version.
	OfflineUserDataJobFailureReasonEnum_UNKNOWN OfflineUserDataJobFailureReasonEnum_OfflineUserDataJobFailureReason = 1
	// The matched transactions are insufficient.
	OfflineUserDataJobFailureReasonEnum_INSUFFICIENT_MATCHED_TRANSACTIONS OfflineUserDataJobFailureReasonEnum_OfflineUserDataJobFailureReason = 2
	// The uploaded transactions are insufficient.
	OfflineUserDataJobFailureReasonEnum_INSUFFICIENT_TRANSACTIONS OfflineUserDataJobFailureReasonEnum_OfflineUserDataJobFailureReason = 3
	// The average transaction value is unusually high for your account. If this
	//
	//	is intended, contact support to request an exception. Learn more at
	//	https://support.google.com/google-ads/answer/10018944#transaction_value
	OfflineUserDataJobFailureReasonEnum_HIGH_AVERAGE_TRANSACTION_VALUE OfflineUserDataJobFailureReasonEnum_OfflineUserDataJobFailureReason = 4
	// The average transaction value is unusually low for your account. If this
	// is intended, contact support to request an exception. Learn more at
	// https://support.google.com/google-ads/answer/10018944#transaction_value
	OfflineUserDataJobFailureReasonEnum_LOW_AVERAGE_TRANSACTION_VALUE OfflineUserDataJobFailureReasonEnum_OfflineUserDataJobFailureReason = 5
	// There's a currency code that you haven't used before in your uploads. If
	// this is intended, contact support to request an exception. Learn more at
	// https://support.google.com/google-ads/answer/10018944#Unrecognized_currency
	OfflineUserDataJobFailureReasonEnum_NEWLY_OBSERVED_CURRENCY_CODE OfflineUserDataJobFailureReasonEnum_OfflineUserDataJobFailureReason = 6
)

// Enum value maps for OfflineUserDataJobFailureReasonEnum_OfflineUserDataJobFailureReason.
var (
	OfflineUserDataJobFailureReasonEnum_OfflineUserDataJobFailureReason_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "INSUFFICIENT_MATCHED_TRANSACTIONS",
		3: "INSUFFICIENT_TRANSACTIONS",
		4: "HIGH_AVERAGE_TRANSACTION_VALUE",
		5: "LOW_AVERAGE_TRANSACTION_VALUE",
		6: "NEWLY_OBSERVED_CURRENCY_CODE",
	}
	OfflineUserDataJobFailureReasonEnum_OfflineUserDataJobFailureReason_value = map[string]int32{
		"UNSPECIFIED":                       0,
		"UNKNOWN":                           1,
		"INSUFFICIENT_MATCHED_TRANSACTIONS": 2,
		"INSUFFICIENT_TRANSACTIONS":         3,
		"HIGH_AVERAGE_TRANSACTION_VALUE":    4,
		"LOW_AVERAGE_TRANSACTION_VALUE":     5,
		"NEWLY_OBSERVED_CURRENCY_CODE":      6,
	}
)

func (x OfflineUserDataJobFailureReasonEnum_OfflineUserDataJobFailureReason) Enum() *OfflineUserDataJobFailureReasonEnum_OfflineUserDataJobFailureReason {
	p := new(OfflineUserDataJobFailureReasonEnum_OfflineUserDataJobFailureReason)
	*p = x
	return p
}

func (x OfflineUserDataJobFailureReasonEnum_OfflineUserDataJobFailureReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OfflineUserDataJobFailureReasonEnum_OfflineUserDataJobFailureReason) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v18_enums_offline_user_data_job_failure_reason_proto_enumTypes[0].Descriptor()
}

func (OfflineUserDataJobFailureReasonEnum_OfflineUserDataJobFailureReason) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v18_enums_offline_user_data_job_failure_reason_proto_enumTypes[0]
}

func (x OfflineUserDataJobFailureReasonEnum_OfflineUserDataJobFailureReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OfflineUserDataJobFailureReasonEnum_OfflineUserDataJobFailureReason.Descriptor instead.
func (OfflineUserDataJobFailureReasonEnum_OfflineUserDataJobFailureReason) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v18_enums_offline_user_data_job_failure_reason_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing reasons why an offline user data job
// failed to be processed.
type OfflineUserDataJobFailureReasonEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OfflineUserDataJobFailureReasonEnum) Reset() {
	*x = OfflineUserDataJobFailureReasonEnum{}
	mi := &file_google_ads_googleads_v18_enums_offline_user_data_job_failure_reason_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OfflineUserDataJobFailureReasonEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OfflineUserDataJobFailureReasonEnum) ProtoMessage() {}

func (x *OfflineUserDataJobFailureReasonEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v18_enums_offline_user_data_job_failure_reason_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OfflineUserDataJobFailureReasonEnum.ProtoReflect.Descriptor instead.
func (*OfflineUserDataJobFailureReasonEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v18_enums_offline_user_data_job_failure_reason_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v18_enums_offline_user_data_job_failure_reason_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v18_enums_offline_user_data_job_failure_reason_proto_rawDesc = []byte{
	0x0a, 0x49, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x38, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x6f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x64, 0x61,
	0x74, 0x61, 0x5f, 0x6a, 0x6f, 0x62, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x5f, 0x72,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x76, 0x31, 0x38, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x22, 0x96, 0x02, 0x0a, 0x23,
	0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x4a,
	0x6f, 0x62, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x45,
	0x6e, 0x75, 0x6d, 0x22, 0xee, 0x01, 0x0a, 0x1f, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x4a, 0x6f, 0x62, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72,
	0x65, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x25, 0x0a, 0x21, 0x49, 0x4e, 0x53, 0x55, 0x46, 0x46, 0x49,
	0x43, 0x49, 0x45, 0x4e, 0x54, 0x5f, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x45, 0x44, 0x5f, 0x54, 0x52,
	0x41, 0x4e, 0x53, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x53, 0x10, 0x02, 0x12, 0x1d, 0x0a, 0x19,
	0x49, 0x4e, 0x53, 0x55, 0x46, 0x46, 0x49, 0x43, 0x49, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x52, 0x41,
	0x4e, 0x53, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x53, 0x10, 0x03, 0x12, 0x22, 0x0a, 0x1e, 0x48,
	0x49, 0x47, 0x48, 0x5f, 0x41, 0x56, 0x45, 0x52, 0x41, 0x47, 0x45, 0x5f, 0x54, 0x52, 0x41, 0x4e,
	0x53, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x10, 0x04, 0x12,
	0x21, 0x0a, 0x1d, 0x4c, 0x4f, 0x57, 0x5f, 0x41, 0x56, 0x45, 0x52, 0x41, 0x47, 0x45, 0x5f, 0x54,
	0x52, 0x41, 0x4e, 0x53, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45,
	0x10, 0x05, 0x12, 0x20, 0x0a, 0x1c, 0x4e, 0x45, 0x57, 0x4c, 0x59, 0x5f, 0x4f, 0x42, 0x53, 0x45,
	0x52, 0x56, 0x45, 0x44, 0x5f, 0x43, 0x55, 0x52, 0x52, 0x45, 0x4e, 0x43, 0x59, 0x5f, 0x43, 0x4f,
	0x44, 0x45, 0x10, 0x06, 0x42, 0xfe, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x76, 0x31, 0x38, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x42, 0x24, 0x4f, 0x66, 0x66,
	0x6c, 0x69, 0x6e, 0x65, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x4a, 0x6f, 0x62, 0x46,
	0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x38, 0x2f, 0x65, 0x6e, 0x75,
	0x6d, 0x73, 0x3b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02,
	0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x38, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xca,
	0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x38, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73,
	0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x38, 0x3a, 0x3a,
	0x45, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v18_enums_offline_user_data_job_failure_reason_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v18_enums_offline_user_data_job_failure_reason_proto_rawDescData = file_google_ads_googleads_v18_enums_offline_user_data_job_failure_reason_proto_rawDesc
)

func file_google_ads_googleads_v18_enums_offline_user_data_job_failure_reason_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v18_enums_offline_user_data_job_failure_reason_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v18_enums_offline_user_data_job_failure_reason_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v18_enums_offline_user_data_job_failure_reason_proto_rawDescData)
	})
	return file_google_ads_googleads_v18_enums_offline_user_data_job_failure_reason_proto_rawDescData
}

var file_google_ads_googleads_v18_enums_offline_user_data_job_failure_reason_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v18_enums_offline_user_data_job_failure_reason_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v18_enums_offline_user_data_job_failure_reason_proto_goTypes = []any{
	(OfflineUserDataJobFailureReasonEnum_OfflineUserDataJobFailureReason)(0), // 0: google.ads.googleads.v18.enums.OfflineUserDataJobFailureReasonEnum.OfflineUserDataJobFailureReason
	(*OfflineUserDataJobFailureReasonEnum)(nil),                              // 1: google.ads.googleads.v18.enums.OfflineUserDataJobFailureReasonEnum
}
var file_google_ads_googleads_v18_enums_offline_user_data_job_failure_reason_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v18_enums_offline_user_data_job_failure_reason_proto_init() }
func file_google_ads_googleads_v18_enums_offline_user_data_job_failure_reason_proto_init() {
	if File_google_ads_googleads_v18_enums_offline_user_data_job_failure_reason_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v18_enums_offline_user_data_job_failure_reason_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v18_enums_offline_user_data_job_failure_reason_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v18_enums_offline_user_data_job_failure_reason_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v18_enums_offline_user_data_job_failure_reason_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v18_enums_offline_user_data_job_failure_reason_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v18_enums_offline_user_data_job_failure_reason_proto = out.File
	file_google_ads_googleads_v18_enums_offline_user_data_job_failure_reason_proto_rawDesc = nil
	file_google_ads_googleads_v18_enums_offline_user_data_job_failure_reason_proto_goTypes = nil
	file_google_ads_googleads_v18_enums_offline_user_data_job_failure_reason_proto_depIdxs = nil
}
