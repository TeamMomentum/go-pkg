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
// source: google/ads/googleads/v19/errors/campaign_lifecycle_goal_error.proto

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

// Enum describing possible campaign lifecycle goal errors.
type CampaignLifecycleGoalErrorEnum_CampaignLifecycleGoalError int32

const (
	// Enum unspecified.
	CampaignLifecycleGoalErrorEnum_UNSPECIFIED CampaignLifecycleGoalErrorEnum_CampaignLifecycleGoalError = 0
	// The received error code is not known in this version.
	CampaignLifecycleGoalErrorEnum_UNKNOWN CampaignLifecycleGoalErrorEnum_CampaignLifecycleGoalError = 1
	// Campaign is not specified.
	CampaignLifecycleGoalErrorEnum_CAMPAIGN_MISSING CampaignLifecycleGoalErrorEnum_CampaignLifecycleGoalError = 2
	// Cannot find the specified campaign.
	CampaignLifecycleGoalErrorEnum_INVALID_CAMPAIGN CampaignLifecycleGoalErrorEnum_CampaignLifecycleGoalError = 3
	// Optimization mode is unspecified or invalid.
	CampaignLifecycleGoalErrorEnum_CUSTOMER_ACQUISITION_INVALID_OPTIMIZATION_MODE CampaignLifecycleGoalErrorEnum_CampaignLifecycleGoalError = 4
	// The configured lifecycle goal setting is not compatible with the bidding
	// strategy the campaign is using. Specifically, BID_HIGHER_FOR_NEW_CUSTOMER
	// requires conversion-value based bidding strategy type such as
	// MAXIMIZE_CONVERSION_VALUE.
	CampaignLifecycleGoalErrorEnum_INCOMPATIBLE_BIDDING_STRATEGY CampaignLifecycleGoalErrorEnum_CampaignLifecycleGoalError = 5
	// Lifecycle goals require the campaign to optimize towards purchase
	// conversion goal.
	CampaignLifecycleGoalErrorEnum_MISSING_PURCHASE_GOAL CampaignLifecycleGoalErrorEnum_CampaignLifecycleGoalError = 6
	// CampaignLifecycleGoal.customer_acquisition_goal_settings.value_settings.high_lifetime_value
	// is invalid or not allowed, such as when the specified value is smaller
	// than 0.01, when the optimization mode is not BID_HIGHER_FOR_NEW_CUSTOMER,
	// or when
	// CampaignLifecycleGoal.customer_acquisition_goal_settings.value_settings.high_lifetime_value
	// is specified smaller than/without
	// CampaignLifecycleGoal.customer_acquisition_goal_settings.value_settings.value.
	CampaignLifecycleGoalErrorEnum_CUSTOMER_ACQUISITION_INVALID_HIGH_LIFETIME_VALUE CampaignLifecycleGoalErrorEnum_CampaignLifecycleGoalError = 7
	// Customer acquisition goal is not supported on this campaign type.
	CampaignLifecycleGoalErrorEnum_CUSTOMER_ACQUISITION_UNSUPPORTED_CAMPAIGN_TYPE CampaignLifecycleGoalErrorEnum_CampaignLifecycleGoalError = 8
	// CampaignLifecycleGoal.customer_acquisition_goal_settings.value_settings.value
	// is invalid or not allowed, such as when the specified value is smaller
	// than 0.01, or when the optimization mode is not
	// BID_HIGHER_FOR_NEW_CUSTOMER.
	CampaignLifecycleGoalErrorEnum_CUSTOMER_ACQUISITION_INVALID_VALUE CampaignLifecycleGoalErrorEnum_CampaignLifecycleGoalError = 9
	// To use BID_HIGHER_FOR_NEW_CUSTOMER mode, either
	// CampaignLifecycleGoal.customer_acquisition_goal_settings.value_settings.value
	// or CustomerLifecycleGoal.customer_acquisition_goal_value_settings.value
	// must have been specified. If a manager account is managing your account's
	// conversion tracking, then only the CustomerLifecycleGoal of that manager
	// account is used.
	CampaignLifecycleGoalErrorEnum_CUSTOMER_ACQUISITION_VALUE_MISSING CampaignLifecycleGoalErrorEnum_CampaignLifecycleGoalError = 10
	// In order for a campaign to adopt the customer acquisition goal,
	// CustomerLifecycleGoal.lifecycle_goal_customer_definition_settings.existing_user_lists
	// must include active and accessible userlist with more than 1000 members
	// in the Search/Youtube network. If a manager account is managing your
	// account's conversion tracking, then only the CustomerLifecycleGoal of
	// that manager account is used. Also make sure that the manager account
	// shares audience segments with sub-accounts with continuous audience
	// sharing.
	CampaignLifecycleGoalErrorEnum_CUSTOMER_ACQUISITION_MISSING_EXISTING_CUSTOMER_DEFINITION CampaignLifecycleGoalErrorEnum_CampaignLifecycleGoalError = 11
	// In order for a campaign to adopt the customer acquisition goal with high
	// lifetime value optimization,
	// CustomerLifecycleGoal.lifecycle_goal_customer_definition_settings.high_lifetime_value_user_lists
	// must include active and accessible userlist with more than 1000 members
	// in the Search/Youtube network. If a manager account is managing your
	// account's conversion tracking, then only the CustomerLifecycleGoal of
	// that manager account is used. Also make sure that the manager account
	// shares audience segments with sub-accounts using continuous audience
	// sharing.
	CampaignLifecycleGoalErrorEnum_CUSTOMER_ACQUISITION_MISSING_HIGH_VALUE_CUSTOMER_DEFINITION CampaignLifecycleGoalErrorEnum_CampaignLifecycleGoalError = 12
)

// Enum value maps for CampaignLifecycleGoalErrorEnum_CampaignLifecycleGoalError.
var (
	CampaignLifecycleGoalErrorEnum_CampaignLifecycleGoalError_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "CAMPAIGN_MISSING",
		3:  "INVALID_CAMPAIGN",
		4:  "CUSTOMER_ACQUISITION_INVALID_OPTIMIZATION_MODE",
		5:  "INCOMPATIBLE_BIDDING_STRATEGY",
		6:  "MISSING_PURCHASE_GOAL",
		7:  "CUSTOMER_ACQUISITION_INVALID_HIGH_LIFETIME_VALUE",
		8:  "CUSTOMER_ACQUISITION_UNSUPPORTED_CAMPAIGN_TYPE",
		9:  "CUSTOMER_ACQUISITION_INVALID_VALUE",
		10: "CUSTOMER_ACQUISITION_VALUE_MISSING",
		11: "CUSTOMER_ACQUISITION_MISSING_EXISTING_CUSTOMER_DEFINITION",
		12: "CUSTOMER_ACQUISITION_MISSING_HIGH_VALUE_CUSTOMER_DEFINITION",
	}
	CampaignLifecycleGoalErrorEnum_CampaignLifecycleGoalError_value = map[string]int32{
		"UNSPECIFIED":      0,
		"UNKNOWN":          1,
		"CAMPAIGN_MISSING": 2,
		"INVALID_CAMPAIGN": 3,
		"CUSTOMER_ACQUISITION_INVALID_OPTIMIZATION_MODE":              4,
		"INCOMPATIBLE_BIDDING_STRATEGY":                               5,
		"MISSING_PURCHASE_GOAL":                                       6,
		"CUSTOMER_ACQUISITION_INVALID_HIGH_LIFETIME_VALUE":            7,
		"CUSTOMER_ACQUISITION_UNSUPPORTED_CAMPAIGN_TYPE":              8,
		"CUSTOMER_ACQUISITION_INVALID_VALUE":                          9,
		"CUSTOMER_ACQUISITION_VALUE_MISSING":                          10,
		"CUSTOMER_ACQUISITION_MISSING_EXISTING_CUSTOMER_DEFINITION":   11,
		"CUSTOMER_ACQUISITION_MISSING_HIGH_VALUE_CUSTOMER_DEFINITION": 12,
	}
)

func (x CampaignLifecycleGoalErrorEnum_CampaignLifecycleGoalError) Enum() *CampaignLifecycleGoalErrorEnum_CampaignLifecycleGoalError {
	p := new(CampaignLifecycleGoalErrorEnum_CampaignLifecycleGoalError)
	*p = x
	return p
}

func (x CampaignLifecycleGoalErrorEnum_CampaignLifecycleGoalError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CampaignLifecycleGoalErrorEnum_CampaignLifecycleGoalError) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v19_errors_campaign_lifecycle_goal_error_proto_enumTypes[0].Descriptor()
}

func (CampaignLifecycleGoalErrorEnum_CampaignLifecycleGoalError) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v19_errors_campaign_lifecycle_goal_error_proto_enumTypes[0]
}

func (x CampaignLifecycleGoalErrorEnum_CampaignLifecycleGoalError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CampaignLifecycleGoalErrorEnum_CampaignLifecycleGoalError.Descriptor instead.
func (CampaignLifecycleGoalErrorEnum_CampaignLifecycleGoalError) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v19_errors_campaign_lifecycle_goal_error_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible campaign lifecycle goal errors.
type CampaignLifecycleGoalErrorEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CampaignLifecycleGoalErrorEnum) Reset() {
	*x = CampaignLifecycleGoalErrorEnum{}
	mi := &file_google_ads_googleads_v19_errors_campaign_lifecycle_goal_error_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CampaignLifecycleGoalErrorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CampaignLifecycleGoalErrorEnum) ProtoMessage() {}

func (x *CampaignLifecycleGoalErrorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v19_errors_campaign_lifecycle_goal_error_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CampaignLifecycleGoalErrorEnum.ProtoReflect.Descriptor instead.
func (*CampaignLifecycleGoalErrorEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v19_errors_campaign_lifecycle_goal_error_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v19_errors_campaign_lifecycle_goal_error_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v19_errors_campaign_lifecycle_goal_error_proto_rawDesc = []byte{
	0x0a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x39, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x2f, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x5f, 0x6c, 0x69, 0x66, 0x65, 0x63,
	0x79, 0x63, 0x6c, 0x65, 0x5f, 0x67, 0x6f, 0x61, 0x6c, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x22, 0xb5, 0x04, 0x0a, 0x1e, 0x43, 0x61, 0x6d, 0x70, 0x61,
	0x69, 0x67, 0x6e, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x47, 0x6f, 0x61, 0x6c,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0x92, 0x04, 0x0a, 0x1a, 0x43, 0x61,
	0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x47,
	0x6f, 0x61, 0x6c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x41, 0x4d, 0x50, 0x41, 0x49,
	0x47, 0x4e, 0x5f, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x14, 0x0a, 0x10,
	0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x43, 0x41, 0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e,
	0x10, 0x03, 0x12, 0x32, 0x0a, 0x2e, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x45, 0x52, 0x5f, 0x41,
	0x43, 0x51, 0x55, 0x49, 0x53, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c,
	0x49, 0x44, 0x5f, 0x4f, 0x50, 0x54, 0x49, 0x4d, 0x49, 0x5a, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x4d, 0x4f, 0x44, 0x45, 0x10, 0x04, 0x12, 0x21, 0x0a, 0x1d, 0x49, 0x4e, 0x43, 0x4f, 0x4d, 0x50,
	0x41, 0x54, 0x49, 0x42, 0x4c, 0x45, 0x5f, 0x42, 0x49, 0x44, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x53,
	0x54, 0x52, 0x41, 0x54, 0x45, 0x47, 0x59, 0x10, 0x05, 0x12, 0x19, 0x0a, 0x15, 0x4d, 0x49, 0x53,
	0x53, 0x49, 0x4e, 0x47, 0x5f, 0x50, 0x55, 0x52, 0x43, 0x48, 0x41, 0x53, 0x45, 0x5f, 0x47, 0x4f,
	0x41, 0x4c, 0x10, 0x06, 0x12, 0x34, 0x0a, 0x30, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x45, 0x52,
	0x5f, 0x41, 0x43, 0x51, 0x55, 0x49, 0x53, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x49, 0x4e, 0x56,
	0x41, 0x4c, 0x49, 0x44, 0x5f, 0x48, 0x49, 0x47, 0x48, 0x5f, 0x4c, 0x49, 0x46, 0x45, 0x54, 0x49,
	0x4d, 0x45, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x10, 0x07, 0x12, 0x32, 0x0a, 0x2e, 0x43, 0x55,
	0x53, 0x54, 0x4f, 0x4d, 0x45, 0x52, 0x5f, 0x41, 0x43, 0x51, 0x55, 0x49, 0x53, 0x49, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x5f, 0x43,
	0x41, 0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x08, 0x12, 0x26,
	0x0a, 0x22, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x45, 0x52, 0x5f, 0x41, 0x43, 0x51, 0x55, 0x49,
	0x53, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x56,
	0x41, 0x4c, 0x55, 0x45, 0x10, 0x09, 0x12, 0x26, 0x0a, 0x22, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d,
	0x45, 0x52, 0x5f, 0x41, 0x43, 0x51, 0x55, 0x49, 0x53, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x56,
	0x41, 0x4c, 0x55, 0x45, 0x5f, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x10, 0x0a, 0x12, 0x3d,
	0x0a, 0x39, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x45, 0x52, 0x5f, 0x41, 0x43, 0x51, 0x55, 0x49,
	0x53, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x5f, 0x45,
	0x58, 0x49, 0x53, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x45, 0x52,
	0x5f, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x0b, 0x12, 0x3f, 0x0a,
	0x3b, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x45, 0x52, 0x5f, 0x41, 0x43, 0x51, 0x55, 0x49, 0x53,
	0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x5f, 0x48, 0x49,
	0x47, 0x48, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x45,
	0x52, 0x5f, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x0c, 0x42, 0xff,
	0x01, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x42, 0x1f, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e,
	0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x47, 0x6f, 0x61, 0x6c, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76,
	0x31, 0x39, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x3b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31,
	0x39, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xca, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c,
	0x56, 0x31, 0x39, 0x5c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xea, 0x02, 0x23, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x39, 0x3a, 0x3a, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v19_errors_campaign_lifecycle_goal_error_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v19_errors_campaign_lifecycle_goal_error_proto_rawDescData = file_google_ads_googleads_v19_errors_campaign_lifecycle_goal_error_proto_rawDesc
)

func file_google_ads_googleads_v19_errors_campaign_lifecycle_goal_error_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v19_errors_campaign_lifecycle_goal_error_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v19_errors_campaign_lifecycle_goal_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v19_errors_campaign_lifecycle_goal_error_proto_rawDescData)
	})
	return file_google_ads_googleads_v19_errors_campaign_lifecycle_goal_error_proto_rawDescData
}

var file_google_ads_googleads_v19_errors_campaign_lifecycle_goal_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v19_errors_campaign_lifecycle_goal_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v19_errors_campaign_lifecycle_goal_error_proto_goTypes = []any{
	(CampaignLifecycleGoalErrorEnum_CampaignLifecycleGoalError)(0), // 0: google.ads.googleads.v19.errors.CampaignLifecycleGoalErrorEnum.CampaignLifecycleGoalError
	(*CampaignLifecycleGoalErrorEnum)(nil),                         // 1: google.ads.googleads.v19.errors.CampaignLifecycleGoalErrorEnum
}
var file_google_ads_googleads_v19_errors_campaign_lifecycle_goal_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v19_errors_campaign_lifecycle_goal_error_proto_init() }
func file_google_ads_googleads_v19_errors_campaign_lifecycle_goal_error_proto_init() {
	if File_google_ads_googleads_v19_errors_campaign_lifecycle_goal_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v19_errors_campaign_lifecycle_goal_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v19_errors_campaign_lifecycle_goal_error_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v19_errors_campaign_lifecycle_goal_error_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v19_errors_campaign_lifecycle_goal_error_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v19_errors_campaign_lifecycle_goal_error_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v19_errors_campaign_lifecycle_goal_error_proto = out.File
	file_google_ads_googleads_v19_errors_campaign_lifecycle_goal_error_proto_rawDesc = nil
	file_google_ads_googleads_v19_errors_campaign_lifecycle_goal_error_proto_goTypes = nil
	file_google_ads_googleads_v19_errors_campaign_lifecycle_goal_error_proto_depIdxs = nil
}
