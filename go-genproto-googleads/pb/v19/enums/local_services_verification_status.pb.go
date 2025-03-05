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
// source: google/ads/googleads/v19/enums/local_services_verification_status.proto

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

// Enum describing status of a particular Local Services Ads verification
// category.
type LocalServicesVerificationStatusEnum_LocalServicesVerificationStatus int32

const (
	// Not specified.
	LocalServicesVerificationStatusEnum_UNSPECIFIED LocalServicesVerificationStatusEnum_LocalServicesVerificationStatus = 0
	// Unknown verification status.
	LocalServicesVerificationStatusEnum_UNKNOWN LocalServicesVerificationStatusEnum_LocalServicesVerificationStatus = 1
	// Verification has started, but has not finished.
	LocalServicesVerificationStatusEnum_NEEDS_REVIEW LocalServicesVerificationStatusEnum_LocalServicesVerificationStatus = 2
	// Verification has failed.
	LocalServicesVerificationStatusEnum_FAILED LocalServicesVerificationStatusEnum_LocalServicesVerificationStatus = 3
	// Verification has passed.
	LocalServicesVerificationStatusEnum_PASSED LocalServicesVerificationStatusEnum_LocalServicesVerificationStatus = 4
	// Verification is not applicable.
	LocalServicesVerificationStatusEnum_NOT_APPLICABLE LocalServicesVerificationStatusEnum_LocalServicesVerificationStatus = 5
	// Verification is required but pending submission.
	LocalServicesVerificationStatusEnum_NO_SUBMISSION LocalServicesVerificationStatusEnum_LocalServicesVerificationStatus = 6
	// Not all required verification has been submitted.
	LocalServicesVerificationStatusEnum_PARTIAL_SUBMISSION LocalServicesVerificationStatusEnum_LocalServicesVerificationStatus = 7
	// Verification needs review by Local Services Ads Ops Specialist.
	LocalServicesVerificationStatusEnum_PENDING_ESCALATION LocalServicesVerificationStatusEnum_LocalServicesVerificationStatus = 8
)

// Enum value maps for LocalServicesVerificationStatusEnum_LocalServicesVerificationStatus.
var (
	LocalServicesVerificationStatusEnum_LocalServicesVerificationStatus_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "NEEDS_REVIEW",
		3: "FAILED",
		4: "PASSED",
		5: "NOT_APPLICABLE",
		6: "NO_SUBMISSION",
		7: "PARTIAL_SUBMISSION",
		8: "PENDING_ESCALATION",
	}
	LocalServicesVerificationStatusEnum_LocalServicesVerificationStatus_value = map[string]int32{
		"UNSPECIFIED":        0,
		"UNKNOWN":            1,
		"NEEDS_REVIEW":       2,
		"FAILED":             3,
		"PASSED":             4,
		"NOT_APPLICABLE":     5,
		"NO_SUBMISSION":      6,
		"PARTIAL_SUBMISSION": 7,
		"PENDING_ESCALATION": 8,
	}
)

func (x LocalServicesVerificationStatusEnum_LocalServicesVerificationStatus) Enum() *LocalServicesVerificationStatusEnum_LocalServicesVerificationStatus {
	p := new(LocalServicesVerificationStatusEnum_LocalServicesVerificationStatus)
	*p = x
	return p
}

func (x LocalServicesVerificationStatusEnum_LocalServicesVerificationStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LocalServicesVerificationStatusEnum_LocalServicesVerificationStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v19_enums_local_services_verification_status_proto_enumTypes[0].Descriptor()
}

func (LocalServicesVerificationStatusEnum_LocalServicesVerificationStatus) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v19_enums_local_services_verification_status_proto_enumTypes[0]
}

func (x LocalServicesVerificationStatusEnum_LocalServicesVerificationStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LocalServicesVerificationStatusEnum_LocalServicesVerificationStatus.Descriptor instead.
func (LocalServicesVerificationStatusEnum_LocalServicesVerificationStatus) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v19_enums_local_services_verification_status_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing status of a particular Local Services Ads
// verification category.
type LocalServicesVerificationStatusEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LocalServicesVerificationStatusEnum) Reset() {
	*x = LocalServicesVerificationStatusEnum{}
	mi := &file_google_ads_googleads_v19_enums_local_services_verification_status_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LocalServicesVerificationStatusEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocalServicesVerificationStatusEnum) ProtoMessage() {}

func (x *LocalServicesVerificationStatusEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v19_enums_local_services_verification_status_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocalServicesVerificationStatusEnum.ProtoReflect.Descriptor instead.
func (*LocalServicesVerificationStatusEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v19_enums_local_services_verification_status_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v19_enums_local_services_verification_status_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v19_enums_local_services_verification_status_proto_rawDesc = []byte{
	0x0a, 0x47, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x39, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x5f,
	0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x39, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x22, 0xe8, 0x01, 0x0a, 0x23, 0x4c, 0x6f,
	0x63, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x56, 0x65, 0x72, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x75,
	0x6d, 0x22, 0xc0, 0x01, 0x0a, 0x1f, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57,
	0x4e, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x4e, 0x45, 0x45, 0x44, 0x53, 0x5f, 0x52, 0x45, 0x56,
	0x49, 0x45, 0x57, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10,
	0x03, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x41, 0x53, 0x53, 0x45, 0x44, 0x10, 0x04, 0x12, 0x12, 0x0a,
	0x0e, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x50, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x42, 0x4c, 0x45, 0x10,
	0x05, 0x12, 0x11, 0x0a, 0x0d, 0x4e, 0x4f, 0x5f, 0x53, 0x55, 0x42, 0x4d, 0x49, 0x53, 0x53, 0x49,
	0x4f, 0x4e, 0x10, 0x06, 0x12, 0x16, 0x0a, 0x12, 0x50, 0x41, 0x52, 0x54, 0x49, 0x41, 0x4c, 0x5f,
	0x53, 0x55, 0x42, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x10, 0x07, 0x12, 0x16, 0x0a, 0x12,
	0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x45, 0x53, 0x43, 0x41, 0x4c, 0x41, 0x54, 0x49,
	0x4f, 0x4e, 0x10, 0x08, 0x42, 0xfe, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x42, 0x24, 0x4c, 0x6f, 0x63,
	0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x39, 0x2f, 0x65, 0x6e, 0x75,
	0x6d, 0x73, 0x3b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02,
	0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x39, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xca,
	0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x39, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73,
	0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x39, 0x3a, 0x3a,
	0x45, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v19_enums_local_services_verification_status_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v19_enums_local_services_verification_status_proto_rawDescData = file_google_ads_googleads_v19_enums_local_services_verification_status_proto_rawDesc
)

func file_google_ads_googleads_v19_enums_local_services_verification_status_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v19_enums_local_services_verification_status_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v19_enums_local_services_verification_status_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v19_enums_local_services_verification_status_proto_rawDescData)
	})
	return file_google_ads_googleads_v19_enums_local_services_verification_status_proto_rawDescData
}

var file_google_ads_googleads_v19_enums_local_services_verification_status_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v19_enums_local_services_verification_status_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v19_enums_local_services_verification_status_proto_goTypes = []any{
	(LocalServicesVerificationStatusEnum_LocalServicesVerificationStatus)(0), // 0: google.ads.googleads.v19.enums.LocalServicesVerificationStatusEnum.LocalServicesVerificationStatus
	(*LocalServicesVerificationStatusEnum)(nil),                              // 1: google.ads.googleads.v19.enums.LocalServicesVerificationStatusEnum
}
var file_google_ads_googleads_v19_enums_local_services_verification_status_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v19_enums_local_services_verification_status_proto_init() }
func file_google_ads_googleads_v19_enums_local_services_verification_status_proto_init() {
	if File_google_ads_googleads_v19_enums_local_services_verification_status_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v19_enums_local_services_verification_status_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v19_enums_local_services_verification_status_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v19_enums_local_services_verification_status_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v19_enums_local_services_verification_status_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v19_enums_local_services_verification_status_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v19_enums_local_services_verification_status_proto = out.File
	file_google_ads_googleads_v19_enums_local_services_verification_status_proto_rawDesc = nil
	file_google_ads_googleads_v19_enums_local_services_verification_status_proto_goTypes = nil
	file_google_ads_googleads_v19_enums_local_services_verification_status_proto_depIdxs = nil
}
