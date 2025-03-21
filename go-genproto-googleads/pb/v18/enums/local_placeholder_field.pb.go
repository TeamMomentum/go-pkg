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
// source: google/ads/googleads/v18/enums/local_placeholder_field.proto

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

// Possible values for Local placeholder fields.
type LocalPlaceholderFieldEnum_LocalPlaceholderField int32

const (
	// Not specified.
	LocalPlaceholderFieldEnum_UNSPECIFIED LocalPlaceholderFieldEnum_LocalPlaceholderField = 0
	// Used for return value only. Represents value unknown in this version.
	LocalPlaceholderFieldEnum_UNKNOWN LocalPlaceholderFieldEnum_LocalPlaceholderField = 1
	// Data Type: STRING. Required. Unique ID.
	LocalPlaceholderFieldEnum_DEAL_ID LocalPlaceholderFieldEnum_LocalPlaceholderField = 2
	// Data Type: STRING. Required. Main headline with local deal title to be
	// shown in dynamic ad.
	LocalPlaceholderFieldEnum_DEAL_NAME LocalPlaceholderFieldEnum_LocalPlaceholderField = 3
	// Data Type: STRING. Local deal subtitle to be shown in dynamic ad.
	LocalPlaceholderFieldEnum_SUBTITLE LocalPlaceholderFieldEnum_LocalPlaceholderField = 4
	// Data Type: STRING. Description of local deal to be shown in dynamic ad.
	LocalPlaceholderFieldEnum_DESCRIPTION LocalPlaceholderFieldEnum_LocalPlaceholderField = 5
	// Data Type: STRING. Price to be shown in the ad. Highly recommended for
	// dynamic ads. Example: "100.00 USD"
	LocalPlaceholderFieldEnum_PRICE LocalPlaceholderFieldEnum_LocalPlaceholderField = 6
	// Data Type: STRING. Formatted price to be shown in the ad.
	// Example: "Starting at $100.00 USD", "$80 - $100"
	LocalPlaceholderFieldEnum_FORMATTED_PRICE LocalPlaceholderFieldEnum_LocalPlaceholderField = 7
	// Data Type: STRING. Sale price to be shown in the ad.
	// Example: "80.00 USD"
	LocalPlaceholderFieldEnum_SALE_PRICE LocalPlaceholderFieldEnum_LocalPlaceholderField = 8
	// Data Type: STRING. Formatted sale price to be shown in the ad.
	// Example: "On sale for $80.00", "$60 - $80"
	LocalPlaceholderFieldEnum_FORMATTED_SALE_PRICE LocalPlaceholderFieldEnum_LocalPlaceholderField = 9
	// Data Type: URL. Image to be displayed in the ad.
	LocalPlaceholderFieldEnum_IMAGE_URL LocalPlaceholderFieldEnum_LocalPlaceholderField = 10
	// Data Type: STRING. Complete property address, including postal code.
	LocalPlaceholderFieldEnum_ADDRESS LocalPlaceholderFieldEnum_LocalPlaceholderField = 11
	// Data Type: STRING. Category of local deal used to group like items
	// together for recommendation engine.
	LocalPlaceholderFieldEnum_CATEGORY LocalPlaceholderFieldEnum_LocalPlaceholderField = 12
	// Data Type: STRING_LIST. Keywords used for product retrieval.
	LocalPlaceholderFieldEnum_CONTEXTUAL_KEYWORDS LocalPlaceholderFieldEnum_LocalPlaceholderField = 13
	// Data Type: URL_LIST. Required. Final URLs to be used in ad when using
	// Upgraded URLs; the more specific the better (for example, the individual
	// URL of a specific local deal and its location).
	LocalPlaceholderFieldEnum_FINAL_URLS LocalPlaceholderFieldEnum_LocalPlaceholderField = 14
	// Data Type: URL_LIST. Final mobile URLs for the ad when using Upgraded
	// URLs.
	LocalPlaceholderFieldEnum_FINAL_MOBILE_URLS LocalPlaceholderFieldEnum_LocalPlaceholderField = 15
	// Data Type: URL. Tracking template for the ad when using Upgraded URLs.
	LocalPlaceholderFieldEnum_TRACKING_URL LocalPlaceholderFieldEnum_LocalPlaceholderField = 16
	// Data Type: STRING. Android app link. Must be formatted as:
	// android-app://{package_id}/{scheme}/{host_path}.
	// The components are defined as follows:
	// package_id: app ID as specified in Google Play.
	// scheme: the scheme to pass to the application. Can be HTTP, or a custom
	//
	//	scheme.
	//
	// host_path: identifies the specific content within your application.
	LocalPlaceholderFieldEnum_ANDROID_APP_LINK LocalPlaceholderFieldEnum_LocalPlaceholderField = 17
	// Data Type: STRING_LIST. List of recommended local deal IDs to show
	// together with this item.
	LocalPlaceholderFieldEnum_SIMILAR_DEAL_IDS LocalPlaceholderFieldEnum_LocalPlaceholderField = 18
	// Data Type: STRING. iOS app link.
	LocalPlaceholderFieldEnum_IOS_APP_LINK LocalPlaceholderFieldEnum_LocalPlaceholderField = 19
	// Data Type: INT64. iOS app store ID.
	LocalPlaceholderFieldEnum_IOS_APP_STORE_ID LocalPlaceholderFieldEnum_LocalPlaceholderField = 20
)

// Enum value maps for LocalPlaceholderFieldEnum_LocalPlaceholderField.
var (
	LocalPlaceholderFieldEnum_LocalPlaceholderField_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "DEAL_ID",
		3:  "DEAL_NAME",
		4:  "SUBTITLE",
		5:  "DESCRIPTION",
		6:  "PRICE",
		7:  "FORMATTED_PRICE",
		8:  "SALE_PRICE",
		9:  "FORMATTED_SALE_PRICE",
		10: "IMAGE_URL",
		11: "ADDRESS",
		12: "CATEGORY",
		13: "CONTEXTUAL_KEYWORDS",
		14: "FINAL_URLS",
		15: "FINAL_MOBILE_URLS",
		16: "TRACKING_URL",
		17: "ANDROID_APP_LINK",
		18: "SIMILAR_DEAL_IDS",
		19: "IOS_APP_LINK",
		20: "IOS_APP_STORE_ID",
	}
	LocalPlaceholderFieldEnum_LocalPlaceholderField_value = map[string]int32{
		"UNSPECIFIED":          0,
		"UNKNOWN":              1,
		"DEAL_ID":              2,
		"DEAL_NAME":            3,
		"SUBTITLE":             4,
		"DESCRIPTION":          5,
		"PRICE":                6,
		"FORMATTED_PRICE":      7,
		"SALE_PRICE":           8,
		"FORMATTED_SALE_PRICE": 9,
		"IMAGE_URL":            10,
		"ADDRESS":              11,
		"CATEGORY":             12,
		"CONTEXTUAL_KEYWORDS":  13,
		"FINAL_URLS":           14,
		"FINAL_MOBILE_URLS":    15,
		"TRACKING_URL":         16,
		"ANDROID_APP_LINK":     17,
		"SIMILAR_DEAL_IDS":     18,
		"IOS_APP_LINK":         19,
		"IOS_APP_STORE_ID":     20,
	}
)

func (x LocalPlaceholderFieldEnum_LocalPlaceholderField) Enum() *LocalPlaceholderFieldEnum_LocalPlaceholderField {
	p := new(LocalPlaceholderFieldEnum_LocalPlaceholderField)
	*p = x
	return p
}

func (x LocalPlaceholderFieldEnum_LocalPlaceholderField) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LocalPlaceholderFieldEnum_LocalPlaceholderField) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v18_enums_local_placeholder_field_proto_enumTypes[0].Descriptor()
}

func (LocalPlaceholderFieldEnum_LocalPlaceholderField) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v18_enums_local_placeholder_field_proto_enumTypes[0]
}

func (x LocalPlaceholderFieldEnum_LocalPlaceholderField) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LocalPlaceholderFieldEnum_LocalPlaceholderField.Descriptor instead.
func (LocalPlaceholderFieldEnum_LocalPlaceholderField) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v18_enums_local_placeholder_field_proto_rawDescGZIP(), []int{0, 0}
}

// Values for Local placeholder fields.
// For more information about dynamic remarketing feeds, see
// https://support.google.com/google-ads/answer/6053288.
type LocalPlaceholderFieldEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LocalPlaceholderFieldEnum) Reset() {
	*x = LocalPlaceholderFieldEnum{}
	mi := &file_google_ads_googleads_v18_enums_local_placeholder_field_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LocalPlaceholderFieldEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocalPlaceholderFieldEnum) ProtoMessage() {}

func (x *LocalPlaceholderFieldEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v18_enums_local_placeholder_field_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocalPlaceholderFieldEnum.ProtoReflect.Descriptor instead.
func (*LocalPlaceholderFieldEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v18_enums_local_placeholder_field_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v18_enums_local_placeholder_field_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v18_enums_local_placeholder_field_proto_rawDesc = []byte{
	0x0a, 0x3c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x38, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x68, 0x6f, 0x6c, 0x64,
	0x65, 0x72, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x38, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x22, 0xa8,
	0x03, 0x0a, 0x19, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x68, 0x6f, 0x6c,
	0x64, 0x65, 0x72, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0x8a, 0x03, 0x0a,
	0x15, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x68, 0x6f, 0x6c, 0x64, 0x65,
	0x72, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x45, 0x41, 0x4c, 0x5f, 0x49, 0x44, 0x10,
	0x02, 0x12, 0x0d, 0x0a, 0x09, 0x44, 0x45, 0x41, 0x4c, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x03,
	0x12, 0x0c, 0x0a, 0x08, 0x53, 0x55, 0x42, 0x54, 0x49, 0x54, 0x4c, 0x45, 0x10, 0x04, 0x12, 0x0f,
	0x0a, 0x0b, 0x44, 0x45, 0x53, 0x43, 0x52, 0x49, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x05, 0x12,
	0x09, 0x0a, 0x05, 0x50, 0x52, 0x49, 0x43, 0x45, 0x10, 0x06, 0x12, 0x13, 0x0a, 0x0f, 0x46, 0x4f,
	0x52, 0x4d, 0x41, 0x54, 0x54, 0x45, 0x44, 0x5f, 0x50, 0x52, 0x49, 0x43, 0x45, 0x10, 0x07, 0x12,
	0x0e, 0x0a, 0x0a, 0x53, 0x41, 0x4c, 0x45, 0x5f, 0x50, 0x52, 0x49, 0x43, 0x45, 0x10, 0x08, 0x12,
	0x18, 0x0a, 0x14, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x54, 0x54, 0x45, 0x44, 0x5f, 0x53, 0x41, 0x4c,
	0x45, 0x5f, 0x50, 0x52, 0x49, 0x43, 0x45, 0x10, 0x09, 0x12, 0x0d, 0x0a, 0x09, 0x49, 0x4d, 0x41,
	0x47, 0x45, 0x5f, 0x55, 0x52, 0x4c, 0x10, 0x0a, 0x12, 0x0b, 0x0a, 0x07, 0x41, 0x44, 0x44, 0x52,
	0x45, 0x53, 0x53, 0x10, 0x0b, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52,
	0x59, 0x10, 0x0c, 0x12, 0x17, 0x0a, 0x13, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x58, 0x54, 0x55, 0x41,
	0x4c, 0x5f, 0x4b, 0x45, 0x59, 0x57, 0x4f, 0x52, 0x44, 0x53, 0x10, 0x0d, 0x12, 0x0e, 0x0a, 0x0a,
	0x46, 0x49, 0x4e, 0x41, 0x4c, 0x5f, 0x55, 0x52, 0x4c, 0x53, 0x10, 0x0e, 0x12, 0x15, 0x0a, 0x11,
	0x46, 0x49, 0x4e, 0x41, 0x4c, 0x5f, 0x4d, 0x4f, 0x42, 0x49, 0x4c, 0x45, 0x5f, 0x55, 0x52, 0x4c,
	0x53, 0x10, 0x0f, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x52, 0x41, 0x43, 0x4b, 0x49, 0x4e, 0x47, 0x5f,
	0x55, 0x52, 0x4c, 0x10, 0x10, 0x12, 0x14, 0x0a, 0x10, 0x41, 0x4e, 0x44, 0x52, 0x4f, 0x49, 0x44,
	0x5f, 0x41, 0x50, 0x50, 0x5f, 0x4c, 0x49, 0x4e, 0x4b, 0x10, 0x11, 0x12, 0x14, 0x0a, 0x10, 0x53,
	0x49, 0x4d, 0x49, 0x4c, 0x41, 0x52, 0x5f, 0x44, 0x45, 0x41, 0x4c, 0x5f, 0x49, 0x44, 0x53, 0x10,
	0x12, 0x12, 0x10, 0x0a, 0x0c, 0x49, 0x4f, 0x53, 0x5f, 0x41, 0x50, 0x50, 0x5f, 0x4c, 0x49, 0x4e,
	0x4b, 0x10, 0x13, 0x12, 0x14, 0x0a, 0x10, 0x49, 0x4f, 0x53, 0x5f, 0x41, 0x50, 0x50, 0x5f, 0x53,
	0x54, 0x4f, 0x52, 0x45, 0x5f, 0x49, 0x44, 0x10, 0x14, 0x42, 0xf4, 0x01, 0x0a, 0x22, 0x63, 0x6f,
	0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x38, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x42, 0x1a, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x68, 0x6f, 0x6c, 0x64,
	0x65, 0x72, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x38, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x3b, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73,
	0x2e, 0x56, 0x31, 0x38, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64,
	0x73, 0x5c, 0x56, 0x31, 0x38, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xea, 0x02, 0x22, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x38, 0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v18_enums_local_placeholder_field_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v18_enums_local_placeholder_field_proto_rawDescData = file_google_ads_googleads_v18_enums_local_placeholder_field_proto_rawDesc
)

func file_google_ads_googleads_v18_enums_local_placeholder_field_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v18_enums_local_placeholder_field_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v18_enums_local_placeholder_field_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v18_enums_local_placeholder_field_proto_rawDescData)
	})
	return file_google_ads_googleads_v18_enums_local_placeholder_field_proto_rawDescData
}

var file_google_ads_googleads_v18_enums_local_placeholder_field_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v18_enums_local_placeholder_field_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v18_enums_local_placeholder_field_proto_goTypes = []any{
	(LocalPlaceholderFieldEnum_LocalPlaceholderField)(0), // 0: google.ads.googleads.v18.enums.LocalPlaceholderFieldEnum.LocalPlaceholderField
	(*LocalPlaceholderFieldEnum)(nil),                    // 1: google.ads.googleads.v18.enums.LocalPlaceholderFieldEnum
}
var file_google_ads_googleads_v18_enums_local_placeholder_field_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v18_enums_local_placeholder_field_proto_init() }
func file_google_ads_googleads_v18_enums_local_placeholder_field_proto_init() {
	if File_google_ads_googleads_v18_enums_local_placeholder_field_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v18_enums_local_placeholder_field_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v18_enums_local_placeholder_field_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v18_enums_local_placeholder_field_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v18_enums_local_placeholder_field_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v18_enums_local_placeholder_field_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v18_enums_local_placeholder_field_proto = out.File
	file_google_ads_googleads_v18_enums_local_placeholder_field_proto_rawDesc = nil
	file_google_ads_googleads_v18_enums_local_placeholder_field_proto_goTypes = nil
	file_google_ads_googleads_v18_enums_local_placeholder_field_proto_depIdxs = nil
}
