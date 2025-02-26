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
// source: google/ads/googleads/v18/enums/external_conversion_source.proto

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

// The external conversion source that is associated with a ConversionAction.
type ExternalConversionSourceEnum_ExternalConversionSource int32

const (
	// Not specified.
	ExternalConversionSourceEnum_UNSPECIFIED ExternalConversionSourceEnum_ExternalConversionSource = 0
	// Represents value unknown in this version.
	ExternalConversionSourceEnum_UNKNOWN ExternalConversionSourceEnum_ExternalConversionSource = 1
	// Conversion that occurs when a user navigates to a particular webpage
	// after viewing an ad; Displayed in Google Ads UI as 'Website'.
	ExternalConversionSourceEnum_WEBPAGE ExternalConversionSourceEnum_ExternalConversionSource = 2
	// Conversion that comes from linked Google Analytics goal or transaction;
	// Displayed in Google Ads UI as 'Analytics'.
	ExternalConversionSourceEnum_ANALYTICS ExternalConversionSourceEnum_ExternalConversionSource = 3
	// Website conversion that is uploaded through ConversionUploadService;
	// Displayed in Google Ads UI as 'Import from clicks'.
	ExternalConversionSourceEnum_UPLOAD ExternalConversionSourceEnum_ExternalConversionSource = 4
	// Conversion that occurs when a user clicks on a call extension directly on
	// an ad; Displayed in Google Ads UI as 'Calls from ads'.
	ExternalConversionSourceEnum_AD_CALL_METRICS ExternalConversionSourceEnum_ExternalConversionSource = 5
	// Conversion that occurs when a user calls a dynamically-generated phone
	// number (by installed javascript) from an advertiser's website after
	// clicking on an ad; Displayed in Google Ads UI as 'Calls from website'.
	ExternalConversionSourceEnum_WEBSITE_CALL_METRICS ExternalConversionSourceEnum_ExternalConversionSource = 6
	// Conversion that occurs when a user visits an advertiser's retail store
	// after clicking on a Google ad;
	// Displayed in Google Ads UI as 'Store visits'.
	ExternalConversionSourceEnum_STORE_VISITS ExternalConversionSourceEnum_ExternalConversionSource = 7
	// Conversion that occurs when a user takes an in-app action such as a
	// purchase in an Android app;
	// Displayed in Google Ads UI as 'Android in-app action'.
	ExternalConversionSourceEnum_ANDROID_IN_APP ExternalConversionSourceEnum_ExternalConversionSource = 8
	// Conversion that occurs when a user takes an in-app action such as a
	// purchase in an iOS app;
	// Displayed in Google Ads UI as 'iOS in-app action'.
	ExternalConversionSourceEnum_IOS_IN_APP ExternalConversionSourceEnum_ExternalConversionSource = 9
	// Conversion that occurs when a user opens an iOS app for the first time;
	// Displayed in Google Ads UI as 'iOS app install (first open)'.
	ExternalConversionSourceEnum_IOS_FIRST_OPEN ExternalConversionSourceEnum_ExternalConversionSource = 10
	// Legacy app conversions that do not have an AppPlatform provided;
	// Displayed in Google Ads UI as 'Mobile app'.
	ExternalConversionSourceEnum_APP_UNSPECIFIED ExternalConversionSourceEnum_ExternalConversionSource = 11
	// Conversion that occurs when a user opens an Android app for the first
	// time; Displayed in Google Ads UI as 'Android app install (first open)'.
	ExternalConversionSourceEnum_ANDROID_FIRST_OPEN ExternalConversionSourceEnum_ExternalConversionSource = 12
	// Call conversion that is uploaded through ConversionUploadService;
	// Displayed in Google Ads UI as 'Import from calls'.
	ExternalConversionSourceEnum_UPLOAD_CALLS ExternalConversionSourceEnum_ExternalConversionSource = 13
	// Conversion that comes from a linked Firebase event;
	// Displayed in Google Ads UI as 'Firebase'.
	ExternalConversionSourceEnum_FIREBASE ExternalConversionSourceEnum_ExternalConversionSource = 14
	// Conversion that occurs when a user clicks on a mobile phone number;
	// Displayed in Google Ads UI as 'Phone number clicks'.
	ExternalConversionSourceEnum_CLICK_TO_CALL ExternalConversionSourceEnum_ExternalConversionSource = 15
	// Conversion that comes from Salesforce;
	// Displayed in Google Ads UI as 'Salesforce.com'.
	ExternalConversionSourceEnum_SALESFORCE ExternalConversionSourceEnum_ExternalConversionSource = 16
	// Conversion that comes from in-store purchases recorded by CRM;
	// Displayed in Google Ads UI as 'Store sales (data partner)'.
	ExternalConversionSourceEnum_STORE_SALES_CRM ExternalConversionSourceEnum_ExternalConversionSource = 17
	// Conversion that comes from in-store purchases from payment network;
	// Displayed in Google Ads UI as 'Store sales (payment network)'.
	ExternalConversionSourceEnum_STORE_SALES_PAYMENT_NETWORK ExternalConversionSourceEnum_ExternalConversionSource = 18
	// Codeless Google Play conversion;
	// Displayed in Google Ads UI as 'Google Play'.
	ExternalConversionSourceEnum_GOOGLE_PLAY ExternalConversionSourceEnum_ExternalConversionSource = 19
	// Conversion that comes from a linked third-party app analytics event;
	// Displayed in Google Ads UI as 'Third-party app analytics'.
	ExternalConversionSourceEnum_THIRD_PARTY_APP_ANALYTICS ExternalConversionSourceEnum_ExternalConversionSource = 20
	// Conversion that is controlled by Google Attribution.
	ExternalConversionSourceEnum_GOOGLE_ATTRIBUTION ExternalConversionSourceEnum_ExternalConversionSource = 21
	// Store Sales conversion based on first-party or third-party merchant data
	// uploads. Displayed in Google Ads UI as 'Store sales (direct upload)'.
	ExternalConversionSourceEnum_STORE_SALES_DIRECT_UPLOAD ExternalConversionSourceEnum_ExternalConversionSource = 23
	// Store Sales conversion based on first-party or third-party merchant
	// data uploads and/or from in-store purchases using cards from payment
	// networks. Displayed in Google Ads UI as 'Store sales'.
	ExternalConversionSourceEnum_STORE_SALES ExternalConversionSourceEnum_ExternalConversionSource = 24
	// Conversions imported from Search Ads 360 Floodlight data.
	ExternalConversionSourceEnum_SEARCH_ADS_360 ExternalConversionSourceEnum_ExternalConversionSource = 25
	// Conversions that track local actions from Google's products and services
	// after interacting with an ad.
	ExternalConversionSourceEnum_GOOGLE_HOSTED ExternalConversionSourceEnum_ExternalConversionSource = 27
	// Conversions reported by Floodlight tags.
	ExternalConversionSourceEnum_FLOODLIGHT ExternalConversionSourceEnum_ExternalConversionSource = 29
	// Conversions that come from Google Analytics specifically for Search Ads
	// 360. Displayed in Google Ads UI as Analytics (SA360).
	ExternalConversionSourceEnum_ANALYTICS_SEARCH_ADS_360 ExternalConversionSourceEnum_ExternalConversionSource = 31
	// Conversion that comes from a linked Firebase event for Search Ads 360.
	ExternalConversionSourceEnum_FIREBASE_SEARCH_ADS_360 ExternalConversionSourceEnum_ExternalConversionSource = 33
	// Conversion that is reported by Floodlight for DV360.
	ExternalConversionSourceEnum_DISPLAY_AND_VIDEO_360_FLOODLIGHT ExternalConversionSourceEnum_ExternalConversionSource = 34
)

// Enum value maps for ExternalConversionSourceEnum_ExternalConversionSource.
var (
	ExternalConversionSourceEnum_ExternalConversionSource_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "WEBPAGE",
		3:  "ANALYTICS",
		4:  "UPLOAD",
		5:  "AD_CALL_METRICS",
		6:  "WEBSITE_CALL_METRICS",
		7:  "STORE_VISITS",
		8:  "ANDROID_IN_APP",
		9:  "IOS_IN_APP",
		10: "IOS_FIRST_OPEN",
		11: "APP_UNSPECIFIED",
		12: "ANDROID_FIRST_OPEN",
		13: "UPLOAD_CALLS",
		14: "FIREBASE",
		15: "CLICK_TO_CALL",
		16: "SALESFORCE",
		17: "STORE_SALES_CRM",
		18: "STORE_SALES_PAYMENT_NETWORK",
		19: "GOOGLE_PLAY",
		20: "THIRD_PARTY_APP_ANALYTICS",
		21: "GOOGLE_ATTRIBUTION",
		23: "STORE_SALES_DIRECT_UPLOAD",
		24: "STORE_SALES",
		25: "SEARCH_ADS_360",
		27: "GOOGLE_HOSTED",
		29: "FLOODLIGHT",
		31: "ANALYTICS_SEARCH_ADS_360",
		33: "FIREBASE_SEARCH_ADS_360",
		34: "DISPLAY_AND_VIDEO_360_FLOODLIGHT",
	}
	ExternalConversionSourceEnum_ExternalConversionSource_value = map[string]int32{
		"UNSPECIFIED":                      0,
		"UNKNOWN":                          1,
		"WEBPAGE":                          2,
		"ANALYTICS":                        3,
		"UPLOAD":                           4,
		"AD_CALL_METRICS":                  5,
		"WEBSITE_CALL_METRICS":             6,
		"STORE_VISITS":                     7,
		"ANDROID_IN_APP":                   8,
		"IOS_IN_APP":                       9,
		"IOS_FIRST_OPEN":                   10,
		"APP_UNSPECIFIED":                  11,
		"ANDROID_FIRST_OPEN":               12,
		"UPLOAD_CALLS":                     13,
		"FIREBASE":                         14,
		"CLICK_TO_CALL":                    15,
		"SALESFORCE":                       16,
		"STORE_SALES_CRM":                  17,
		"STORE_SALES_PAYMENT_NETWORK":      18,
		"GOOGLE_PLAY":                      19,
		"THIRD_PARTY_APP_ANALYTICS":        20,
		"GOOGLE_ATTRIBUTION":               21,
		"STORE_SALES_DIRECT_UPLOAD":        23,
		"STORE_SALES":                      24,
		"SEARCH_ADS_360":                   25,
		"GOOGLE_HOSTED":                    27,
		"FLOODLIGHT":                       29,
		"ANALYTICS_SEARCH_ADS_360":         31,
		"FIREBASE_SEARCH_ADS_360":          33,
		"DISPLAY_AND_VIDEO_360_FLOODLIGHT": 34,
	}
)

func (x ExternalConversionSourceEnum_ExternalConversionSource) Enum() *ExternalConversionSourceEnum_ExternalConversionSource {
	p := new(ExternalConversionSourceEnum_ExternalConversionSource)
	*p = x
	return p
}

func (x ExternalConversionSourceEnum_ExternalConversionSource) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ExternalConversionSourceEnum_ExternalConversionSource) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v18_enums_external_conversion_source_proto_enumTypes[0].Descriptor()
}

func (ExternalConversionSourceEnum_ExternalConversionSource) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v18_enums_external_conversion_source_proto_enumTypes[0]
}

func (x ExternalConversionSourceEnum_ExternalConversionSource) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ExternalConversionSourceEnum_ExternalConversionSource.Descriptor instead.
func (ExternalConversionSourceEnum_ExternalConversionSource) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v18_enums_external_conversion_source_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing the external conversion source that is
// associated with a ConversionAction.
type ExternalConversionSourceEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ExternalConversionSourceEnum) Reset() {
	*x = ExternalConversionSourceEnum{}
	mi := &file_google_ads_googleads_v18_enums_external_conversion_source_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExternalConversionSourceEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExternalConversionSourceEnum) ProtoMessage() {}

func (x *ExternalConversionSourceEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v18_enums_external_conversion_source_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExternalConversionSourceEnum.ProtoReflect.Descriptor instead.
func (*ExternalConversionSourceEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v18_enums_external_conversion_source_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v18_enums_external_conversion_source_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v18_enums_external_conversion_source_proto_rawDesc = []byte{
	0x0a, 0x3f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x38, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x38, 0x2e, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x22, 0xb0, 0x05, 0x0a, 0x1c, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x43, 0x6f,
	0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x45, 0x6e,
	0x75, 0x6d, 0x22, 0x8f, 0x05, 0x0a, 0x18, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x43,
	0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12,
	0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x0b, 0x0a,
	0x07, 0x57, 0x45, 0x42, 0x50, 0x41, 0x47, 0x45, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x41, 0x4e,
	0x41, 0x4c, 0x59, 0x54, 0x49, 0x43, 0x53, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x50, 0x4c,
	0x4f, 0x41, 0x44, 0x10, 0x04, 0x12, 0x13, 0x0a, 0x0f, 0x41, 0x44, 0x5f, 0x43, 0x41, 0x4c, 0x4c,
	0x5f, 0x4d, 0x45, 0x54, 0x52, 0x49, 0x43, 0x53, 0x10, 0x05, 0x12, 0x18, 0x0a, 0x14, 0x57, 0x45,
	0x42, 0x53, 0x49, 0x54, 0x45, 0x5f, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x4d, 0x45, 0x54, 0x52, 0x49,
	0x43, 0x53, 0x10, 0x06, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x54, 0x4f, 0x52, 0x45, 0x5f, 0x56, 0x49,
	0x53, 0x49, 0x54, 0x53, 0x10, 0x07, 0x12, 0x12, 0x0a, 0x0e, 0x41, 0x4e, 0x44, 0x52, 0x4f, 0x49,
	0x44, 0x5f, 0x49, 0x4e, 0x5f, 0x41, 0x50, 0x50, 0x10, 0x08, 0x12, 0x0e, 0x0a, 0x0a, 0x49, 0x4f,
	0x53, 0x5f, 0x49, 0x4e, 0x5f, 0x41, 0x50, 0x50, 0x10, 0x09, 0x12, 0x12, 0x0a, 0x0e, 0x49, 0x4f,
	0x53, 0x5f, 0x46, 0x49, 0x52, 0x53, 0x54, 0x5f, 0x4f, 0x50, 0x45, 0x4e, 0x10, 0x0a, 0x12, 0x13,
	0x0a, 0x0f, 0x41, 0x50, 0x50, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x0b, 0x12, 0x16, 0x0a, 0x12, 0x41, 0x4e, 0x44, 0x52, 0x4f, 0x49, 0x44, 0x5f, 0x46,
	0x49, 0x52, 0x53, 0x54, 0x5f, 0x4f, 0x50, 0x45, 0x4e, 0x10, 0x0c, 0x12, 0x10, 0x0a, 0x0c, 0x55,
	0x50, 0x4c, 0x4f, 0x41, 0x44, 0x5f, 0x43, 0x41, 0x4c, 0x4c, 0x53, 0x10, 0x0d, 0x12, 0x0c, 0x0a,
	0x08, 0x46, 0x49, 0x52, 0x45, 0x42, 0x41, 0x53, 0x45, 0x10, 0x0e, 0x12, 0x11, 0x0a, 0x0d, 0x43,
	0x4c, 0x49, 0x43, 0x4b, 0x5f, 0x54, 0x4f, 0x5f, 0x43, 0x41, 0x4c, 0x4c, 0x10, 0x0f, 0x12, 0x0e,
	0x0a, 0x0a, 0x53, 0x41, 0x4c, 0x45, 0x53, 0x46, 0x4f, 0x52, 0x43, 0x45, 0x10, 0x10, 0x12, 0x13,
	0x0a, 0x0f, 0x53, 0x54, 0x4f, 0x52, 0x45, 0x5f, 0x53, 0x41, 0x4c, 0x45, 0x53, 0x5f, 0x43, 0x52,
	0x4d, 0x10, 0x11, 0x12, 0x1f, 0x0a, 0x1b, 0x53, 0x54, 0x4f, 0x52, 0x45, 0x5f, 0x53, 0x41, 0x4c,
	0x45, 0x53, 0x5f, 0x50, 0x41, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x4e, 0x45, 0x54, 0x57, 0x4f,
	0x52, 0x4b, 0x10, 0x12, 0x12, 0x0f, 0x0a, 0x0b, 0x47, 0x4f, 0x4f, 0x47, 0x4c, 0x45, 0x5f, 0x50,
	0x4c, 0x41, 0x59, 0x10, 0x13, 0x12, 0x1d, 0x0a, 0x19, 0x54, 0x48, 0x49, 0x52, 0x44, 0x5f, 0x50,
	0x41, 0x52, 0x54, 0x59, 0x5f, 0x41, 0x50, 0x50, 0x5f, 0x41, 0x4e, 0x41, 0x4c, 0x59, 0x54, 0x49,
	0x43, 0x53, 0x10, 0x14, 0x12, 0x16, 0x0a, 0x12, 0x47, 0x4f, 0x4f, 0x47, 0x4c, 0x45, 0x5f, 0x41,
	0x54, 0x54, 0x52, 0x49, 0x42, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x15, 0x12, 0x1d, 0x0a, 0x19,
	0x53, 0x54, 0x4f, 0x52, 0x45, 0x5f, 0x53, 0x41, 0x4c, 0x45, 0x53, 0x5f, 0x44, 0x49, 0x52, 0x45,
	0x43, 0x54, 0x5f, 0x55, 0x50, 0x4c, 0x4f, 0x41, 0x44, 0x10, 0x17, 0x12, 0x0f, 0x0a, 0x0b, 0x53,
	0x54, 0x4f, 0x52, 0x45, 0x5f, 0x53, 0x41, 0x4c, 0x45, 0x53, 0x10, 0x18, 0x12, 0x12, 0x0a, 0x0e,
	0x53, 0x45, 0x41, 0x52, 0x43, 0x48, 0x5f, 0x41, 0x44, 0x53, 0x5f, 0x33, 0x36, 0x30, 0x10, 0x19,
	0x12, 0x11, 0x0a, 0x0d, 0x47, 0x4f, 0x4f, 0x47, 0x4c, 0x45, 0x5f, 0x48, 0x4f, 0x53, 0x54, 0x45,
	0x44, 0x10, 0x1b, 0x12, 0x0e, 0x0a, 0x0a, 0x46, 0x4c, 0x4f, 0x4f, 0x44, 0x4c, 0x49, 0x47, 0x48,
	0x54, 0x10, 0x1d, 0x12, 0x1c, 0x0a, 0x18, 0x41, 0x4e, 0x41, 0x4c, 0x59, 0x54, 0x49, 0x43, 0x53,
	0x5f, 0x53, 0x45, 0x41, 0x52, 0x43, 0x48, 0x5f, 0x41, 0x44, 0x53, 0x5f, 0x33, 0x36, 0x30, 0x10,
	0x1f, 0x12, 0x1b, 0x0a, 0x17, 0x46, 0x49, 0x52, 0x45, 0x42, 0x41, 0x53, 0x45, 0x5f, 0x53, 0x45,
	0x41, 0x52, 0x43, 0x48, 0x5f, 0x41, 0x44, 0x53, 0x5f, 0x33, 0x36, 0x30, 0x10, 0x21, 0x12, 0x24,
	0x0a, 0x20, 0x44, 0x49, 0x53, 0x50, 0x4c, 0x41, 0x59, 0x5f, 0x41, 0x4e, 0x44, 0x5f, 0x56, 0x49,
	0x44, 0x45, 0x4f, 0x5f, 0x33, 0x36, 0x30, 0x5f, 0x46, 0x4c, 0x4f, 0x4f, 0x44, 0x4c, 0x49, 0x47,
	0x48, 0x54, 0x10, 0x22, 0x42, 0xf7, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x76, 0x31, 0x38, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x42, 0x1d, 0x45, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x6f,
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
	file_google_ads_googleads_v18_enums_external_conversion_source_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v18_enums_external_conversion_source_proto_rawDescData = file_google_ads_googleads_v18_enums_external_conversion_source_proto_rawDesc
)

func file_google_ads_googleads_v18_enums_external_conversion_source_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v18_enums_external_conversion_source_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v18_enums_external_conversion_source_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v18_enums_external_conversion_source_proto_rawDescData)
	})
	return file_google_ads_googleads_v18_enums_external_conversion_source_proto_rawDescData
}

var file_google_ads_googleads_v18_enums_external_conversion_source_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v18_enums_external_conversion_source_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v18_enums_external_conversion_source_proto_goTypes = []any{
	(ExternalConversionSourceEnum_ExternalConversionSource)(0), // 0: google.ads.googleads.v18.enums.ExternalConversionSourceEnum.ExternalConversionSource
	(*ExternalConversionSourceEnum)(nil),                       // 1: google.ads.googleads.v18.enums.ExternalConversionSourceEnum
}
var file_google_ads_googleads_v18_enums_external_conversion_source_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v18_enums_external_conversion_source_proto_init() }
func file_google_ads_googleads_v18_enums_external_conversion_source_proto_init() {
	if File_google_ads_googleads_v18_enums_external_conversion_source_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v18_enums_external_conversion_source_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v18_enums_external_conversion_source_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v18_enums_external_conversion_source_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v18_enums_external_conversion_source_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v18_enums_external_conversion_source_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v18_enums_external_conversion_source_proto = out.File
	file_google_ads_googleads_v18_enums_external_conversion_source_proto_rawDesc = nil
	file_google_ads_googleads_v18_enums_external_conversion_source_proto_goTypes = nil
	file_google_ads_googleads_v18_enums_external_conversion_source_proto_depIdxs = nil
}
