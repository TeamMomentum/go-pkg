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
// source: google/ads/googleads/v18/enums/promotion_placeholder_field.proto

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

// Possible values for Promotion placeholder fields.
type PromotionPlaceholderFieldEnum_PromotionPlaceholderField int32

const (
	// Not specified.
	PromotionPlaceholderFieldEnum_UNSPECIFIED PromotionPlaceholderFieldEnum_PromotionPlaceholderField = 0
	// Used for return value only. Represents value unknown in this version.
	PromotionPlaceholderFieldEnum_UNKNOWN PromotionPlaceholderFieldEnum_PromotionPlaceholderField = 1
	// Data Type: STRING. The text that appears on the ad when the extension is
	// shown.
	PromotionPlaceholderFieldEnum_PROMOTION_TARGET PromotionPlaceholderFieldEnum_PromotionPlaceholderField = 2
	// Data Type: STRING. Lets you add "up to" phrase to the promotion,
	// in case you have variable promotion rates.
	PromotionPlaceholderFieldEnum_DISCOUNT_MODIFIER PromotionPlaceholderFieldEnum_PromotionPlaceholderField = 3
	// Data Type: INT64. Takes a value in micros, where 1 million micros
	// represents 1%, and is shown as a percentage when rendered.
	PromotionPlaceholderFieldEnum_PERCENT_OFF PromotionPlaceholderFieldEnum_PromotionPlaceholderField = 4
	// Data Type: MONEY. Requires a currency and an amount of money.
	PromotionPlaceholderFieldEnum_MONEY_AMOUNT_OFF PromotionPlaceholderFieldEnum_PromotionPlaceholderField = 5
	// Data Type: STRING. A string that the user enters to get the discount.
	PromotionPlaceholderFieldEnum_PROMOTION_CODE PromotionPlaceholderFieldEnum_PromotionPlaceholderField = 6
	// Data Type: MONEY. A minimum spend before the user qualifies for the
	// promotion.
	PromotionPlaceholderFieldEnum_ORDERS_OVER_AMOUNT PromotionPlaceholderFieldEnum_PromotionPlaceholderField = 7
	// Data Type: DATE. The start date of the promotion.
	PromotionPlaceholderFieldEnum_PROMOTION_START PromotionPlaceholderFieldEnum_PromotionPlaceholderField = 8
	// Data Type: DATE. The end date of the promotion.
	PromotionPlaceholderFieldEnum_PROMOTION_END PromotionPlaceholderFieldEnum_PromotionPlaceholderField = 9
	// Data Type: STRING. Describes the associated event for the promotion using
	// one of the PromotionExtensionOccasion enum values, for example NEW_YEARS.
	PromotionPlaceholderFieldEnum_OCCASION PromotionPlaceholderFieldEnum_PromotionPlaceholderField = 10
	// Data Type: URL_LIST. Final URLs to be used in the ad when using Upgraded
	// URLs.
	PromotionPlaceholderFieldEnum_FINAL_URLS PromotionPlaceholderFieldEnum_PromotionPlaceholderField = 11
	// Data Type: URL_LIST. Final mobile URLs for the ad when using Upgraded
	// URLs.
	PromotionPlaceholderFieldEnum_FINAL_MOBILE_URLS PromotionPlaceholderFieldEnum_PromotionPlaceholderField = 12
	// Data Type: URL. Tracking template for the ad when using Upgraded URLs.
	PromotionPlaceholderFieldEnum_TRACKING_URL PromotionPlaceholderFieldEnum_PromotionPlaceholderField = 13
	// Data Type: STRING. A string represented by a language code for the
	// promotion.
	PromotionPlaceholderFieldEnum_LANGUAGE PromotionPlaceholderFieldEnum_PromotionPlaceholderField = 14
	// Data Type: STRING. Final URL suffix for the ad when using parallel
	// tracking.
	PromotionPlaceholderFieldEnum_FINAL_URL_SUFFIX PromotionPlaceholderFieldEnum_PromotionPlaceholderField = 15
)

// Enum value maps for PromotionPlaceholderFieldEnum_PromotionPlaceholderField.
var (
	PromotionPlaceholderFieldEnum_PromotionPlaceholderField_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "PROMOTION_TARGET",
		3:  "DISCOUNT_MODIFIER",
		4:  "PERCENT_OFF",
		5:  "MONEY_AMOUNT_OFF",
		6:  "PROMOTION_CODE",
		7:  "ORDERS_OVER_AMOUNT",
		8:  "PROMOTION_START",
		9:  "PROMOTION_END",
		10: "OCCASION",
		11: "FINAL_URLS",
		12: "FINAL_MOBILE_URLS",
		13: "TRACKING_URL",
		14: "LANGUAGE",
		15: "FINAL_URL_SUFFIX",
	}
	PromotionPlaceholderFieldEnum_PromotionPlaceholderField_value = map[string]int32{
		"UNSPECIFIED":        0,
		"UNKNOWN":            1,
		"PROMOTION_TARGET":   2,
		"DISCOUNT_MODIFIER":  3,
		"PERCENT_OFF":        4,
		"MONEY_AMOUNT_OFF":   5,
		"PROMOTION_CODE":     6,
		"ORDERS_OVER_AMOUNT": 7,
		"PROMOTION_START":    8,
		"PROMOTION_END":      9,
		"OCCASION":           10,
		"FINAL_URLS":         11,
		"FINAL_MOBILE_URLS":  12,
		"TRACKING_URL":       13,
		"LANGUAGE":           14,
		"FINAL_URL_SUFFIX":   15,
	}
)

func (x PromotionPlaceholderFieldEnum_PromotionPlaceholderField) Enum() *PromotionPlaceholderFieldEnum_PromotionPlaceholderField {
	p := new(PromotionPlaceholderFieldEnum_PromotionPlaceholderField)
	*p = x
	return p
}

func (x PromotionPlaceholderFieldEnum_PromotionPlaceholderField) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PromotionPlaceholderFieldEnum_PromotionPlaceholderField) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v18_enums_promotion_placeholder_field_proto_enumTypes[0].Descriptor()
}

func (PromotionPlaceholderFieldEnum_PromotionPlaceholderField) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v18_enums_promotion_placeholder_field_proto_enumTypes[0]
}

func (x PromotionPlaceholderFieldEnum_PromotionPlaceholderField) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PromotionPlaceholderFieldEnum_PromotionPlaceholderField.Descriptor instead.
func (PromotionPlaceholderFieldEnum_PromotionPlaceholderField) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v18_enums_promotion_placeholder_field_proto_rawDescGZIP(), []int{0, 0}
}

// Values for Promotion placeholder fields.
type PromotionPlaceholderFieldEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PromotionPlaceholderFieldEnum) Reset() {
	*x = PromotionPlaceholderFieldEnum{}
	mi := &file_google_ads_googleads_v18_enums_promotion_placeholder_field_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PromotionPlaceholderFieldEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PromotionPlaceholderFieldEnum) ProtoMessage() {}

func (x *PromotionPlaceholderFieldEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v18_enums_promotion_placeholder_field_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PromotionPlaceholderFieldEnum.ProtoReflect.Descriptor instead.
func (*PromotionPlaceholderFieldEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v18_enums_promotion_placeholder_field_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v18_enums_promotion_placeholder_field_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v18_enums_promotion_placeholder_field_proto_rawDesc = []byte{
	0x0a, 0x40, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x38, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x6c, 0x61, 0x63, 0x65,
	0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x38, 0x2e, 0x65, 0x6e, 0x75,
	0x6d, 0x73, 0x22, 0xee, 0x02, 0x0a, 0x1d, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e,
	0x50, 0x6c, 0x61, 0x63, 0x65, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x45, 0x6e, 0x75, 0x6d, 0x22, 0xcc, 0x02, 0x0a, 0x19, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69,
	0x6f, 0x6e, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01,
	0x12, 0x14, 0x0a, 0x10, 0x50, 0x52, 0x4f, 0x4d, 0x4f, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x41,
	0x52, 0x47, 0x45, 0x54, 0x10, 0x02, 0x12, 0x15, 0x0a, 0x11, 0x44, 0x49, 0x53, 0x43, 0x4f, 0x55,
	0x4e, 0x54, 0x5f, 0x4d, 0x4f, 0x44, 0x49, 0x46, 0x49, 0x45, 0x52, 0x10, 0x03, 0x12, 0x0f, 0x0a,
	0x0b, 0x50, 0x45, 0x52, 0x43, 0x45, 0x4e, 0x54, 0x5f, 0x4f, 0x46, 0x46, 0x10, 0x04, 0x12, 0x14,
	0x0a, 0x10, 0x4d, 0x4f, 0x4e, 0x45, 0x59, 0x5f, 0x41, 0x4d, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x4f,
	0x46, 0x46, 0x10, 0x05, 0x12, 0x12, 0x0a, 0x0e, 0x50, 0x52, 0x4f, 0x4d, 0x4f, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x10, 0x06, 0x12, 0x16, 0x0a, 0x12, 0x4f, 0x52, 0x44, 0x45,
	0x52, 0x53, 0x5f, 0x4f, 0x56, 0x45, 0x52, 0x5f, 0x41, 0x4d, 0x4f, 0x55, 0x4e, 0x54, 0x10, 0x07,
	0x12, 0x13, 0x0a, 0x0f, 0x50, 0x52, 0x4f, 0x4d, 0x4f, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54,
	0x41, 0x52, 0x54, 0x10, 0x08, 0x12, 0x11, 0x0a, 0x0d, 0x50, 0x52, 0x4f, 0x4d, 0x4f, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x45, 0x4e, 0x44, 0x10, 0x09, 0x12, 0x0c, 0x0a, 0x08, 0x4f, 0x43, 0x43, 0x41,
	0x53, 0x49, 0x4f, 0x4e, 0x10, 0x0a, 0x12, 0x0e, 0x0a, 0x0a, 0x46, 0x49, 0x4e, 0x41, 0x4c, 0x5f,
	0x55, 0x52, 0x4c, 0x53, 0x10, 0x0b, 0x12, 0x15, 0x0a, 0x11, 0x46, 0x49, 0x4e, 0x41, 0x4c, 0x5f,
	0x4d, 0x4f, 0x42, 0x49, 0x4c, 0x45, 0x5f, 0x55, 0x52, 0x4c, 0x53, 0x10, 0x0c, 0x12, 0x10, 0x0a,
	0x0c, 0x54, 0x52, 0x41, 0x43, 0x4b, 0x49, 0x4e, 0x47, 0x5f, 0x55, 0x52, 0x4c, 0x10, 0x0d, 0x12,
	0x0c, 0x0a, 0x08, 0x4c, 0x41, 0x4e, 0x47, 0x55, 0x41, 0x47, 0x45, 0x10, 0x0e, 0x12, 0x14, 0x0a,
	0x10, 0x46, 0x49, 0x4e, 0x41, 0x4c, 0x5f, 0x55, 0x52, 0x4c, 0x5f, 0x53, 0x55, 0x46, 0x46, 0x49,
	0x58, 0x10, 0x0f, 0x42, 0xf8, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x31, 0x38, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x42, 0x1e, 0x50, 0x72, 0x6f, 0x6d,
	0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2f, 0x76, 0x31, 0x38, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x3b, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56,
	0x31, 0x38, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c,
	0x56, 0x31, 0x38, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41,
	0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x38, 0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v18_enums_promotion_placeholder_field_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v18_enums_promotion_placeholder_field_proto_rawDescData = file_google_ads_googleads_v18_enums_promotion_placeholder_field_proto_rawDesc
)

func file_google_ads_googleads_v18_enums_promotion_placeholder_field_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v18_enums_promotion_placeholder_field_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v18_enums_promotion_placeholder_field_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v18_enums_promotion_placeholder_field_proto_rawDescData)
	})
	return file_google_ads_googleads_v18_enums_promotion_placeholder_field_proto_rawDescData
}

var file_google_ads_googleads_v18_enums_promotion_placeholder_field_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v18_enums_promotion_placeholder_field_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v18_enums_promotion_placeholder_field_proto_goTypes = []any{
	(PromotionPlaceholderFieldEnum_PromotionPlaceholderField)(0), // 0: google.ads.googleads.v18.enums.PromotionPlaceholderFieldEnum.PromotionPlaceholderField
	(*PromotionPlaceholderFieldEnum)(nil),                        // 1: google.ads.googleads.v18.enums.PromotionPlaceholderFieldEnum
}
var file_google_ads_googleads_v18_enums_promotion_placeholder_field_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v18_enums_promotion_placeholder_field_proto_init() }
func file_google_ads_googleads_v18_enums_promotion_placeholder_field_proto_init() {
	if File_google_ads_googleads_v18_enums_promotion_placeholder_field_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v18_enums_promotion_placeholder_field_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v18_enums_promotion_placeholder_field_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v18_enums_promotion_placeholder_field_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v18_enums_promotion_placeholder_field_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v18_enums_promotion_placeholder_field_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v18_enums_promotion_placeholder_field_proto = out.File
	file_google_ads_googleads_v18_enums_promotion_placeholder_field_proto_rawDesc = nil
	file_google_ads_googleads_v18_enums_promotion_placeholder_field_proto_goTypes = nil
	file_google_ads_googleads_v18_enums_promotion_placeholder_field_proto_depIdxs = nil
}
