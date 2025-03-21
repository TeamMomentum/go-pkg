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
// source: google/ads/googleads/v18/enums/placeholder_type.proto

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

// Possible placeholder types for a feed mapping.
type PlaceholderTypeEnum_PlaceholderType int32

const (
	// Not specified.
	PlaceholderTypeEnum_UNSPECIFIED PlaceholderTypeEnum_PlaceholderType = 0
	// Used for return value only. Represents value unknown in this version.
	PlaceholderTypeEnum_UNKNOWN PlaceholderTypeEnum_PlaceholderType = 1
	// Lets you show links in your ad to pages from your website, including the
	// main landing page.
	PlaceholderTypeEnum_SITELINK PlaceholderTypeEnum_PlaceholderType = 2
	// Lets you attach a phone number to an ad, allowing customers to call
	// directly from the ad.
	PlaceholderTypeEnum_CALL PlaceholderTypeEnum_PlaceholderType = 3
	// Lets you provide users with a link that points to a mobile app in
	// addition to a website.
	PlaceholderTypeEnum_APP PlaceholderTypeEnum_PlaceholderType = 4
	// Lets you show locations of businesses from your Business Profile
	// in your ad. This helps people find your locations by showing your
	// ads with your address, a map to your location, or the distance to your
	// business. This extension type is useful to draw customers to your
	// brick-and-mortar location.
	PlaceholderTypeEnum_LOCATION PlaceholderTypeEnum_PlaceholderType = 5
	// If you sell your product through retail chains, affiliate location
	// extensions let you show nearby stores that carry your products.
	PlaceholderTypeEnum_AFFILIATE_LOCATION PlaceholderTypeEnum_PlaceholderType = 6
	// Lets you include additional text with your search ads that provide
	// detailed information about your business, including products and services
	// you offer. Callouts appear in ads at the top and bottom of Google search
	// results.
	PlaceholderTypeEnum_CALLOUT PlaceholderTypeEnum_PlaceholderType = 7
	// Lets you add more info to your ad, specific to some predefined categories
	// such as types, brands, styles, etc. A minimum of 3 text (SNIPPETS) values
	// are required.
	PlaceholderTypeEnum_STRUCTURED_SNIPPET PlaceholderTypeEnum_PlaceholderType = 8
	// Allows users to see your ad, click an icon, and contact you directly by
	// text message. With one tap on your ad, people can contact you to book an
	// appointment, get a quote, ask for information, or request a service.
	PlaceholderTypeEnum_MESSAGE PlaceholderTypeEnum_PlaceholderType = 9
	// Lets you display prices for a list of items along with your ads. A price
	// feed is composed of three to eight price table rows.
	PlaceholderTypeEnum_PRICE PlaceholderTypeEnum_PlaceholderType = 10
	// Lets you highlight sales and other promotions that let users see how
	// they can save by buying now.
	PlaceholderTypeEnum_PROMOTION PlaceholderTypeEnum_PlaceholderType = 11
	// Lets you dynamically inject custom data into the title and description
	// of your ads.
	PlaceholderTypeEnum_AD_CUSTOMIZER PlaceholderTypeEnum_PlaceholderType = 12
	// Indicates that this feed is for education dynamic remarketing.
	PlaceholderTypeEnum_DYNAMIC_EDUCATION PlaceholderTypeEnum_PlaceholderType = 13
	// Indicates that this feed is for flight dynamic remarketing.
	PlaceholderTypeEnum_DYNAMIC_FLIGHT PlaceholderTypeEnum_PlaceholderType = 14
	// Indicates that this feed is for a custom dynamic remarketing type. Use
	// this only if the other business types don't apply to your products or
	// services.
	PlaceholderTypeEnum_DYNAMIC_CUSTOM PlaceholderTypeEnum_PlaceholderType = 15
	// Indicates that this feed is for hotels and rentals dynamic remarketing.
	PlaceholderTypeEnum_DYNAMIC_HOTEL PlaceholderTypeEnum_PlaceholderType = 16
	// Indicates that this feed is for real estate dynamic remarketing.
	PlaceholderTypeEnum_DYNAMIC_REAL_ESTATE PlaceholderTypeEnum_PlaceholderType = 17
	// Indicates that this feed is for travel dynamic remarketing.
	PlaceholderTypeEnum_DYNAMIC_TRAVEL PlaceholderTypeEnum_PlaceholderType = 18
	// Indicates that this feed is for local deals dynamic remarketing.
	PlaceholderTypeEnum_DYNAMIC_LOCAL PlaceholderTypeEnum_PlaceholderType = 19
	// Indicates that this feed is for job dynamic remarketing.
	PlaceholderTypeEnum_DYNAMIC_JOB PlaceholderTypeEnum_PlaceholderType = 20
	// Lets you attach an image to an ad.
	PlaceholderTypeEnum_IMAGE PlaceholderTypeEnum_PlaceholderType = 21
)

// Enum value maps for PlaceholderTypeEnum_PlaceholderType.
var (
	PlaceholderTypeEnum_PlaceholderType_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "SITELINK",
		3:  "CALL",
		4:  "APP",
		5:  "LOCATION",
		6:  "AFFILIATE_LOCATION",
		7:  "CALLOUT",
		8:  "STRUCTURED_SNIPPET",
		9:  "MESSAGE",
		10: "PRICE",
		11: "PROMOTION",
		12: "AD_CUSTOMIZER",
		13: "DYNAMIC_EDUCATION",
		14: "DYNAMIC_FLIGHT",
		15: "DYNAMIC_CUSTOM",
		16: "DYNAMIC_HOTEL",
		17: "DYNAMIC_REAL_ESTATE",
		18: "DYNAMIC_TRAVEL",
		19: "DYNAMIC_LOCAL",
		20: "DYNAMIC_JOB",
		21: "IMAGE",
	}
	PlaceholderTypeEnum_PlaceholderType_value = map[string]int32{
		"UNSPECIFIED":         0,
		"UNKNOWN":             1,
		"SITELINK":            2,
		"CALL":                3,
		"APP":                 4,
		"LOCATION":            5,
		"AFFILIATE_LOCATION":  6,
		"CALLOUT":             7,
		"STRUCTURED_SNIPPET":  8,
		"MESSAGE":             9,
		"PRICE":               10,
		"PROMOTION":           11,
		"AD_CUSTOMIZER":       12,
		"DYNAMIC_EDUCATION":   13,
		"DYNAMIC_FLIGHT":      14,
		"DYNAMIC_CUSTOM":      15,
		"DYNAMIC_HOTEL":       16,
		"DYNAMIC_REAL_ESTATE": 17,
		"DYNAMIC_TRAVEL":      18,
		"DYNAMIC_LOCAL":       19,
		"DYNAMIC_JOB":         20,
		"IMAGE":               21,
	}
)

func (x PlaceholderTypeEnum_PlaceholderType) Enum() *PlaceholderTypeEnum_PlaceholderType {
	p := new(PlaceholderTypeEnum_PlaceholderType)
	*p = x
	return p
}

func (x PlaceholderTypeEnum_PlaceholderType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PlaceholderTypeEnum_PlaceholderType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v18_enums_placeholder_type_proto_enumTypes[0].Descriptor()
}

func (PlaceholderTypeEnum_PlaceholderType) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v18_enums_placeholder_type_proto_enumTypes[0]
}

func (x PlaceholderTypeEnum_PlaceholderType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PlaceholderTypeEnum_PlaceholderType.Descriptor instead.
func (PlaceholderTypeEnum_PlaceholderType) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v18_enums_placeholder_type_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible placeholder types for a feed mapping.
type PlaceholderTypeEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PlaceholderTypeEnum) Reset() {
	*x = PlaceholderTypeEnum{}
	mi := &file_google_ads_googleads_v18_enums_placeholder_type_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PlaceholderTypeEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlaceholderTypeEnum) ProtoMessage() {}

func (x *PlaceholderTypeEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v18_enums_placeholder_type_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlaceholderTypeEnum.ProtoReflect.Descriptor instead.
func (*PlaceholderTypeEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v18_enums_placeholder_type_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v18_enums_placeholder_type_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v18_enums_placeholder_type_proto_rawDesc = []byte{
	0x0a, 0x35, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x38, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31,
	0x38, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x22, 0x9b, 0x03, 0x0a, 0x13, 0x50, 0x6c, 0x61, 0x63,
	0x65, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x22,
	0x83, 0x03, 0x0a, 0x0f, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10,
	0x01, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x49, 0x54, 0x45, 0x4c, 0x49, 0x4e, 0x4b, 0x10, 0x02, 0x12,
	0x08, 0x0a, 0x04, 0x43, 0x41, 0x4c, 0x4c, 0x10, 0x03, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x50, 0x50,
	0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08, 0x4c, 0x4f, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x05,
	0x12, 0x16, 0x0a, 0x12, 0x41, 0x46, 0x46, 0x49, 0x4c, 0x49, 0x41, 0x54, 0x45, 0x5f, 0x4c, 0x4f,
	0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x06, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x41, 0x4c, 0x4c,
	0x4f, 0x55, 0x54, 0x10, 0x07, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x52, 0x55, 0x43, 0x54, 0x55,
	0x52, 0x45, 0x44, 0x5f, 0x53, 0x4e, 0x49, 0x50, 0x50, 0x45, 0x54, 0x10, 0x08, 0x12, 0x0b, 0x0a,
	0x07, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x10, 0x09, 0x12, 0x09, 0x0a, 0x05, 0x50, 0x52,
	0x49, 0x43, 0x45, 0x10, 0x0a, 0x12, 0x0d, 0x0a, 0x09, 0x50, 0x52, 0x4f, 0x4d, 0x4f, 0x54, 0x49,
	0x4f, 0x4e, 0x10, 0x0b, 0x12, 0x11, 0x0a, 0x0d, 0x41, 0x44, 0x5f, 0x43, 0x55, 0x53, 0x54, 0x4f,
	0x4d, 0x49, 0x5a, 0x45, 0x52, 0x10, 0x0c, 0x12, 0x15, 0x0a, 0x11, 0x44, 0x59, 0x4e, 0x41, 0x4d,
	0x49, 0x43, 0x5f, 0x45, 0x44, 0x55, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x0d, 0x12, 0x12,
	0x0a, 0x0e, 0x44, 0x59, 0x4e, 0x41, 0x4d, 0x49, 0x43, 0x5f, 0x46, 0x4c, 0x49, 0x47, 0x48, 0x54,
	0x10, 0x0e, 0x12, 0x12, 0x0a, 0x0e, 0x44, 0x59, 0x4e, 0x41, 0x4d, 0x49, 0x43, 0x5f, 0x43, 0x55,
	0x53, 0x54, 0x4f, 0x4d, 0x10, 0x0f, 0x12, 0x11, 0x0a, 0x0d, 0x44, 0x59, 0x4e, 0x41, 0x4d, 0x49,
	0x43, 0x5f, 0x48, 0x4f, 0x54, 0x45, 0x4c, 0x10, 0x10, 0x12, 0x17, 0x0a, 0x13, 0x44, 0x59, 0x4e,
	0x41, 0x4d, 0x49, 0x43, 0x5f, 0x52, 0x45, 0x41, 0x4c, 0x5f, 0x45, 0x53, 0x54, 0x41, 0x54, 0x45,
	0x10, 0x11, 0x12, 0x12, 0x0a, 0x0e, 0x44, 0x59, 0x4e, 0x41, 0x4d, 0x49, 0x43, 0x5f, 0x54, 0x52,
	0x41, 0x56, 0x45, 0x4c, 0x10, 0x12, 0x12, 0x11, 0x0a, 0x0d, 0x44, 0x59, 0x4e, 0x41, 0x4d, 0x49,
	0x43, 0x5f, 0x4c, 0x4f, 0x43, 0x41, 0x4c, 0x10, 0x13, 0x12, 0x0f, 0x0a, 0x0b, 0x44, 0x59, 0x4e,
	0x41, 0x4d, 0x49, 0x43, 0x5f, 0x4a, 0x4f, 0x42, 0x10, 0x14, 0x12, 0x09, 0x0a, 0x05, 0x49, 0x4d,
	0x41, 0x47, 0x45, 0x10, 0x15, 0x42, 0xee, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x76, 0x31, 0x38, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x42, 0x14, 0x50, 0x6c,
	0x61, 0x63, 0x65, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c,
	0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x38, 0x2f, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x3b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa,
	0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x38, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73,
	0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x38, 0x5c, 0x45, 0x6e, 0x75, 0x6d,
	0x73, 0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a,
	0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x38, 0x3a,
	0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v18_enums_placeholder_type_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v18_enums_placeholder_type_proto_rawDescData = file_google_ads_googleads_v18_enums_placeholder_type_proto_rawDesc
)

func file_google_ads_googleads_v18_enums_placeholder_type_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v18_enums_placeholder_type_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v18_enums_placeholder_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v18_enums_placeholder_type_proto_rawDescData)
	})
	return file_google_ads_googleads_v18_enums_placeholder_type_proto_rawDescData
}

var file_google_ads_googleads_v18_enums_placeholder_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v18_enums_placeholder_type_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v18_enums_placeholder_type_proto_goTypes = []any{
	(PlaceholderTypeEnum_PlaceholderType)(0), // 0: google.ads.googleads.v18.enums.PlaceholderTypeEnum.PlaceholderType
	(*PlaceholderTypeEnum)(nil),              // 1: google.ads.googleads.v18.enums.PlaceholderTypeEnum
}
var file_google_ads_googleads_v18_enums_placeholder_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v18_enums_placeholder_type_proto_init() }
func file_google_ads_googleads_v18_enums_placeholder_type_proto_init() {
	if File_google_ads_googleads_v18_enums_placeholder_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v18_enums_placeholder_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v18_enums_placeholder_type_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v18_enums_placeholder_type_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v18_enums_placeholder_type_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v18_enums_placeholder_type_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v18_enums_placeholder_type_proto = out.File
	file_google_ads_googleads_v18_enums_placeholder_type_proto_rawDesc = nil
	file_google_ads_googleads_v18_enums_placeholder_type_proto_goTypes = nil
	file_google_ads_googleads_v18_enums_placeholder_type_proto_depIdxs = nil
}
