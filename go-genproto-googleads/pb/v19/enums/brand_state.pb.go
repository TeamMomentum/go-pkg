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
// source: google/ads/googleads/v19/enums/brand_state.proto

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

// Enumeration of different brand states.
type BrandStateEnum_BrandState int32

const (
	// No value has been specified.
	BrandStateEnum_UNSPECIFIED BrandStateEnum_BrandState = 0
	// Used for return value only. Represents value unknown in this version.
	BrandStateEnum_UNKNOWN BrandStateEnum_BrandState = 1
	// Brand is verified and globally available for selection
	BrandStateEnum_ENABLED BrandStateEnum_BrandState = 2
	// Brand was globally available in past but is no longer a valid brand
	// (based on business criteria)
	BrandStateEnum_DEPRECATED BrandStateEnum_BrandState = 3
	// Brand is unverified and customer scoped, but can be selected by customer
	// (only who requested for same) for targeting
	BrandStateEnum_UNVERIFIED BrandStateEnum_BrandState = 4
	// Was a customer-scoped (unverified) brand, which got approved by business
	// and added to the global list. Its assigned CKG MID should be used instead
	// of this
	BrandStateEnum_APPROVED BrandStateEnum_BrandState = 5
	// Was a customer-scoped (unverified) brand, but the request was canceled by
	// customer and this brand id is no longer valid
	BrandStateEnum_CANCELLED BrandStateEnum_BrandState = 6
	// Was a customer-scoped (unverified) brand, but the request was rejected by
	// internal business team and this brand id is no longer valid
	BrandStateEnum_REJECTED BrandStateEnum_BrandState = 7
)

// Enum value maps for BrandStateEnum_BrandState.
var (
	BrandStateEnum_BrandState_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "ENABLED",
		3: "DEPRECATED",
		4: "UNVERIFIED",
		5: "APPROVED",
		6: "CANCELLED",
		7: "REJECTED",
	}
	BrandStateEnum_BrandState_value = map[string]int32{
		"UNSPECIFIED": 0,
		"UNKNOWN":     1,
		"ENABLED":     2,
		"DEPRECATED":  3,
		"UNVERIFIED":  4,
		"APPROVED":    5,
		"CANCELLED":   6,
		"REJECTED":    7,
	}
)

func (x BrandStateEnum_BrandState) Enum() *BrandStateEnum_BrandState {
	p := new(BrandStateEnum_BrandState)
	*p = x
	return p
}

func (x BrandStateEnum_BrandState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BrandStateEnum_BrandState) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v19_enums_brand_state_proto_enumTypes[0].Descriptor()
}

func (BrandStateEnum_BrandState) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v19_enums_brand_state_proto_enumTypes[0]
}

func (x BrandStateEnum_BrandState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BrandStateEnum_BrandState.Descriptor instead.
func (BrandStateEnum_BrandState) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v19_enums_brand_state_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible brand states.
type BrandStateEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BrandStateEnum) Reset() {
	*x = BrandStateEnum{}
	mi := &file_google_ads_googleads_v19_enums_brand_state_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BrandStateEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BrandStateEnum) ProtoMessage() {}

func (x *BrandStateEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v19_enums_brand_state_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BrandStateEnum.ProtoReflect.Descriptor instead.
func (*BrandStateEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v19_enums_brand_state_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v19_enums_brand_state_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v19_enums_brand_state_proto_rawDesc = []byte{
	0x0a, 0x30, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x39, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e, 0x65, 0x6e, 0x75,
	0x6d, 0x73, 0x22, 0x95, 0x01, 0x0a, 0x0e, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0x82, 0x01, 0x0a, 0x0a, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x45, 0x4e, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x12,
	0x0e, 0x0a, 0x0a, 0x44, 0x45, 0x50, 0x52, 0x45, 0x43, 0x41, 0x54, 0x45, 0x44, 0x10, 0x03, 0x12,
	0x0e, 0x0a, 0x0a, 0x55, 0x4e, 0x56, 0x45, 0x52, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x04, 0x12,
	0x0c, 0x0a, 0x08, 0x41, 0x50, 0x50, 0x52, 0x4f, 0x56, 0x45, 0x44, 0x10, 0x05, 0x12, 0x0d, 0x0a,
	0x09, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x4c, 0x45, 0x44, 0x10, 0x06, 0x12, 0x0c, 0x0a, 0x08,
	0x52, 0x45, 0x4a, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x07, 0x42, 0xe9, 0x01, 0x0a, 0x22, 0x63,
	0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x42, 0x0f, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c,
	0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x39, 0x2f, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x3b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa,
	0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x39, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73,
	0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x39, 0x5c, 0x45, 0x6e, 0x75, 0x6d,
	0x73, 0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a,
	0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x39, 0x3a,
	0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v19_enums_brand_state_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v19_enums_brand_state_proto_rawDescData = file_google_ads_googleads_v19_enums_brand_state_proto_rawDesc
)

func file_google_ads_googleads_v19_enums_brand_state_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v19_enums_brand_state_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v19_enums_brand_state_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v19_enums_brand_state_proto_rawDescData)
	})
	return file_google_ads_googleads_v19_enums_brand_state_proto_rawDescData
}

var file_google_ads_googleads_v19_enums_brand_state_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v19_enums_brand_state_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v19_enums_brand_state_proto_goTypes = []any{
	(BrandStateEnum_BrandState)(0), // 0: google.ads.googleads.v19.enums.BrandStateEnum.BrandState
	(*BrandStateEnum)(nil),         // 1: google.ads.googleads.v19.enums.BrandStateEnum
}
var file_google_ads_googleads_v19_enums_brand_state_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v19_enums_brand_state_proto_init() }
func file_google_ads_googleads_v19_enums_brand_state_proto_init() {
	if File_google_ads_googleads_v19_enums_brand_state_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v19_enums_brand_state_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v19_enums_brand_state_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v19_enums_brand_state_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v19_enums_brand_state_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v19_enums_brand_state_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v19_enums_brand_state_proto = out.File
	file_google_ads_googleads_v19_enums_brand_state_proto_rawDesc = nil
	file_google_ads_googleads_v19_enums_brand_state_proto_goTypes = nil
	file_google_ads_googleads_v19_enums_brand_state_proto_depIdxs = nil
}
