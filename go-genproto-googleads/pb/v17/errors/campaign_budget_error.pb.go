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
// source: google/ads/googleads/v17/errors/campaign_budget_error.proto

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

// Enum describing possible campaign budget errors.
type CampaignBudgetErrorEnum_CampaignBudgetError int32

const (
	// Enum unspecified.
	CampaignBudgetErrorEnum_UNSPECIFIED CampaignBudgetErrorEnum_CampaignBudgetError = 0
	// The received error code is not known in this version.
	CampaignBudgetErrorEnum_UNKNOWN CampaignBudgetErrorEnum_CampaignBudgetError = 1
	// The campaign budget cannot be shared.
	CampaignBudgetErrorEnum_CAMPAIGN_BUDGET_CANNOT_BE_SHARED CampaignBudgetErrorEnum_CampaignBudgetError = 17
	// The requested campaign budget no longer exists.
	CampaignBudgetErrorEnum_CAMPAIGN_BUDGET_REMOVED CampaignBudgetErrorEnum_CampaignBudgetError = 2
	// The campaign budget is associated with at least one campaign, and so the
	// campaign budget cannot be removed.
	CampaignBudgetErrorEnum_CAMPAIGN_BUDGET_IN_USE CampaignBudgetErrorEnum_CampaignBudgetError = 3
	// Customer is not on the allow-list for this campaign budget period.
	CampaignBudgetErrorEnum_CAMPAIGN_BUDGET_PERIOD_NOT_AVAILABLE CampaignBudgetErrorEnum_CampaignBudgetError = 4
	// This field is not mutable on implicitly shared campaign budgets
	CampaignBudgetErrorEnum_CANNOT_MODIFY_FIELD_OF_IMPLICITLY_SHARED_CAMPAIGN_BUDGET CampaignBudgetErrorEnum_CampaignBudgetError = 6
	// Cannot change explicitly shared campaign budgets back to implicitly
	// shared ones.
	CampaignBudgetErrorEnum_CANNOT_UPDATE_CAMPAIGN_BUDGET_TO_IMPLICITLY_SHARED CampaignBudgetErrorEnum_CampaignBudgetError = 7
	// An implicit campaign budget without a name cannot be changed to
	// explicitly shared campaign budget.
	CampaignBudgetErrorEnum_CANNOT_UPDATE_CAMPAIGN_BUDGET_TO_EXPLICITLY_SHARED_WITHOUT_NAME CampaignBudgetErrorEnum_CampaignBudgetError = 8
	// Cannot change an implicitly shared campaign budget to an explicitly
	// shared one.
	CampaignBudgetErrorEnum_CANNOT_UPDATE_CAMPAIGN_BUDGET_TO_EXPLICITLY_SHARED CampaignBudgetErrorEnum_CampaignBudgetError = 9
	// Only explicitly shared campaign budgets can be used with multiple
	// campaigns.
	CampaignBudgetErrorEnum_CANNOT_USE_IMPLICITLY_SHARED_CAMPAIGN_BUDGET_WITH_MULTIPLE_CAMPAIGNS CampaignBudgetErrorEnum_CampaignBudgetError = 10
	// A campaign budget with this name already exists.
	CampaignBudgetErrorEnum_DUPLICATE_NAME CampaignBudgetErrorEnum_CampaignBudgetError = 11
	// A money amount was not in the expected currency.
	CampaignBudgetErrorEnum_MONEY_AMOUNT_IN_WRONG_CURRENCY CampaignBudgetErrorEnum_CampaignBudgetError = 12
	// A money amount was less than the minimum CPC for currency.
	CampaignBudgetErrorEnum_MONEY_AMOUNT_LESS_THAN_CURRENCY_MINIMUM_CPC CampaignBudgetErrorEnum_CampaignBudgetError = 13
	// A money amount was greater than the maximum allowed.
	CampaignBudgetErrorEnum_MONEY_AMOUNT_TOO_LARGE CampaignBudgetErrorEnum_CampaignBudgetError = 14
	// A money amount was negative.
	CampaignBudgetErrorEnum_NEGATIVE_MONEY_AMOUNT CampaignBudgetErrorEnum_CampaignBudgetError = 15
	// A money amount was not a multiple of a minimum unit.
	CampaignBudgetErrorEnum_NON_MULTIPLE_OF_MINIMUM_CURRENCY_UNIT CampaignBudgetErrorEnum_CampaignBudgetError = 16
	// Total budget amount must be unset when BudgetPeriod is DAILY.
	CampaignBudgetErrorEnum_TOTAL_BUDGET_AMOUNT_MUST_BE_UNSET_FOR_BUDGET_PERIOD_DAILY CampaignBudgetErrorEnum_CampaignBudgetError = 18
	// The period of the budget is not allowed.
	CampaignBudgetErrorEnum_INVALID_PERIOD CampaignBudgetErrorEnum_CampaignBudgetError = 19
	// Cannot use accelerated delivery method on this budget.
	CampaignBudgetErrorEnum_CANNOT_USE_ACCELERATED_DELIVERY_MODE CampaignBudgetErrorEnum_CampaignBudgetError = 20
	// Budget amount must be unset when BudgetPeriod is CUSTOM.
	CampaignBudgetErrorEnum_BUDGET_AMOUNT_MUST_BE_UNSET_FOR_CUSTOM_BUDGET_PERIOD CampaignBudgetErrorEnum_CampaignBudgetError = 21
)

// Enum value maps for CampaignBudgetErrorEnum_CampaignBudgetError.
var (
	CampaignBudgetErrorEnum_CampaignBudgetError_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		17: "CAMPAIGN_BUDGET_CANNOT_BE_SHARED",
		2:  "CAMPAIGN_BUDGET_REMOVED",
		3:  "CAMPAIGN_BUDGET_IN_USE",
		4:  "CAMPAIGN_BUDGET_PERIOD_NOT_AVAILABLE",
		6:  "CANNOT_MODIFY_FIELD_OF_IMPLICITLY_SHARED_CAMPAIGN_BUDGET",
		7:  "CANNOT_UPDATE_CAMPAIGN_BUDGET_TO_IMPLICITLY_SHARED",
		8:  "CANNOT_UPDATE_CAMPAIGN_BUDGET_TO_EXPLICITLY_SHARED_WITHOUT_NAME",
		9:  "CANNOT_UPDATE_CAMPAIGN_BUDGET_TO_EXPLICITLY_SHARED",
		10: "CANNOT_USE_IMPLICITLY_SHARED_CAMPAIGN_BUDGET_WITH_MULTIPLE_CAMPAIGNS",
		11: "DUPLICATE_NAME",
		12: "MONEY_AMOUNT_IN_WRONG_CURRENCY",
		13: "MONEY_AMOUNT_LESS_THAN_CURRENCY_MINIMUM_CPC",
		14: "MONEY_AMOUNT_TOO_LARGE",
		15: "NEGATIVE_MONEY_AMOUNT",
		16: "NON_MULTIPLE_OF_MINIMUM_CURRENCY_UNIT",
		18: "TOTAL_BUDGET_AMOUNT_MUST_BE_UNSET_FOR_BUDGET_PERIOD_DAILY",
		19: "INVALID_PERIOD",
		20: "CANNOT_USE_ACCELERATED_DELIVERY_MODE",
		21: "BUDGET_AMOUNT_MUST_BE_UNSET_FOR_CUSTOM_BUDGET_PERIOD",
	}
	CampaignBudgetErrorEnum_CampaignBudgetError_value = map[string]int32{
		"UNSPECIFIED":                          0,
		"UNKNOWN":                              1,
		"CAMPAIGN_BUDGET_CANNOT_BE_SHARED":     17,
		"CAMPAIGN_BUDGET_REMOVED":              2,
		"CAMPAIGN_BUDGET_IN_USE":               3,
		"CAMPAIGN_BUDGET_PERIOD_NOT_AVAILABLE": 4,
		"CANNOT_MODIFY_FIELD_OF_IMPLICITLY_SHARED_CAMPAIGN_BUDGET":             6,
		"CANNOT_UPDATE_CAMPAIGN_BUDGET_TO_IMPLICITLY_SHARED":                   7,
		"CANNOT_UPDATE_CAMPAIGN_BUDGET_TO_EXPLICITLY_SHARED_WITHOUT_NAME":      8,
		"CANNOT_UPDATE_CAMPAIGN_BUDGET_TO_EXPLICITLY_SHARED":                   9,
		"CANNOT_USE_IMPLICITLY_SHARED_CAMPAIGN_BUDGET_WITH_MULTIPLE_CAMPAIGNS": 10,
		"DUPLICATE_NAME":                                            11,
		"MONEY_AMOUNT_IN_WRONG_CURRENCY":                            12,
		"MONEY_AMOUNT_LESS_THAN_CURRENCY_MINIMUM_CPC":               13,
		"MONEY_AMOUNT_TOO_LARGE":                                    14,
		"NEGATIVE_MONEY_AMOUNT":                                     15,
		"NON_MULTIPLE_OF_MINIMUM_CURRENCY_UNIT":                     16,
		"TOTAL_BUDGET_AMOUNT_MUST_BE_UNSET_FOR_BUDGET_PERIOD_DAILY": 18,
		"INVALID_PERIOD":                                            19,
		"CANNOT_USE_ACCELERATED_DELIVERY_MODE":                      20,
		"BUDGET_AMOUNT_MUST_BE_UNSET_FOR_CUSTOM_BUDGET_PERIOD":      21,
	}
)

func (x CampaignBudgetErrorEnum_CampaignBudgetError) Enum() *CampaignBudgetErrorEnum_CampaignBudgetError {
	p := new(CampaignBudgetErrorEnum_CampaignBudgetError)
	*p = x
	return p
}

func (x CampaignBudgetErrorEnum_CampaignBudgetError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CampaignBudgetErrorEnum_CampaignBudgetError) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v17_errors_campaign_budget_error_proto_enumTypes[0].Descriptor()
}

func (CampaignBudgetErrorEnum_CampaignBudgetError) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v17_errors_campaign_budget_error_proto_enumTypes[0]
}

func (x CampaignBudgetErrorEnum_CampaignBudgetError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CampaignBudgetErrorEnum_CampaignBudgetError.Descriptor instead.
func (CampaignBudgetErrorEnum_CampaignBudgetError) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v17_errors_campaign_budget_error_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible campaign budget errors.
type CampaignBudgetErrorEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CampaignBudgetErrorEnum) Reset() {
	*x = CampaignBudgetErrorEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v17_errors_campaign_budget_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CampaignBudgetErrorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CampaignBudgetErrorEnum) ProtoMessage() {}

func (x *CampaignBudgetErrorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v17_errors_campaign_budget_error_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CampaignBudgetErrorEnum.ProtoReflect.Descriptor instead.
func (*CampaignBudgetErrorEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v17_errors_campaign_budget_error_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v17_errors_campaign_budget_error_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v17_errors_campaign_budget_error_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x37, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x2f, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x5f, 0x62, 0x75, 0x64, 0x67, 0x65,
	0x74, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x22, 0x97,
	0x07, 0x0a, 0x17, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x42, 0x75, 0x64, 0x67, 0x65,
	0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0xfb, 0x06, 0x0a, 0x13, 0x43,
	0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01,
	0x12, 0x24, 0x0a, 0x20, 0x43, 0x41, 0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e, 0x5f, 0x42, 0x55, 0x44,
	0x47, 0x45, 0x54, 0x5f, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x42, 0x45, 0x5f, 0x53, 0x48,
	0x41, 0x52, 0x45, 0x44, 0x10, 0x11, 0x12, 0x1b, 0x0a, 0x17, 0x43, 0x41, 0x4d, 0x50, 0x41, 0x49,
	0x47, 0x4e, 0x5f, 0x42, 0x55, 0x44, 0x47, 0x45, 0x54, 0x5f, 0x52, 0x45, 0x4d, 0x4f, 0x56, 0x45,
	0x44, 0x10, 0x02, 0x12, 0x1a, 0x0a, 0x16, 0x43, 0x41, 0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e, 0x5f,
	0x42, 0x55, 0x44, 0x47, 0x45, 0x54, 0x5f, 0x49, 0x4e, 0x5f, 0x55, 0x53, 0x45, 0x10, 0x03, 0x12,
	0x28, 0x0a, 0x24, 0x43, 0x41, 0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e, 0x5f, 0x42, 0x55, 0x44, 0x47,
	0x45, 0x54, 0x5f, 0x50, 0x45, 0x52, 0x49, 0x4f, 0x44, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x56,
	0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x04, 0x12, 0x3c, 0x0a, 0x38, 0x43, 0x41, 0x4e,
	0x4e, 0x4f, 0x54, 0x5f, 0x4d, 0x4f, 0x44, 0x49, 0x46, 0x59, 0x5f, 0x46, 0x49, 0x45, 0x4c, 0x44,
	0x5f, 0x4f, 0x46, 0x5f, 0x49, 0x4d, 0x50, 0x4c, 0x49, 0x43, 0x49, 0x54, 0x4c, 0x59, 0x5f, 0x53,
	0x48, 0x41, 0x52, 0x45, 0x44, 0x5f, 0x43, 0x41, 0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e, 0x5f, 0x42,
	0x55, 0x44, 0x47, 0x45, 0x54, 0x10, 0x06, 0x12, 0x36, 0x0a, 0x32, 0x43, 0x41, 0x4e, 0x4e, 0x4f,
	0x54, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x43, 0x41, 0x4d, 0x50, 0x41, 0x49, 0x47,
	0x4e, 0x5f, 0x42, 0x55, 0x44, 0x47, 0x45, 0x54, 0x5f, 0x54, 0x4f, 0x5f, 0x49, 0x4d, 0x50, 0x4c,
	0x49, 0x43, 0x49, 0x54, 0x4c, 0x59, 0x5f, 0x53, 0x48, 0x41, 0x52, 0x45, 0x44, 0x10, 0x07, 0x12,
	0x43, 0x0a, 0x3f, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45,
	0x5f, 0x43, 0x41, 0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e, 0x5f, 0x42, 0x55, 0x44, 0x47, 0x45, 0x54,
	0x5f, 0x54, 0x4f, 0x5f, 0x45, 0x58, 0x50, 0x4c, 0x49, 0x43, 0x49, 0x54, 0x4c, 0x59, 0x5f, 0x53,
	0x48, 0x41, 0x52, 0x45, 0x44, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x4f, 0x55, 0x54, 0x5f, 0x4e, 0x41,
	0x4d, 0x45, 0x10, 0x08, 0x12, 0x36, 0x0a, 0x32, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x55,
	0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x43, 0x41, 0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e, 0x5f, 0x42,
	0x55, 0x44, 0x47, 0x45, 0x54, 0x5f, 0x54, 0x4f, 0x5f, 0x45, 0x58, 0x50, 0x4c, 0x49, 0x43, 0x49,
	0x54, 0x4c, 0x59, 0x5f, 0x53, 0x48, 0x41, 0x52, 0x45, 0x44, 0x10, 0x09, 0x12, 0x48, 0x0a, 0x44,
	0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x55, 0x53, 0x45, 0x5f, 0x49, 0x4d, 0x50, 0x4c, 0x49,
	0x43, 0x49, 0x54, 0x4c, 0x59, 0x5f, 0x53, 0x48, 0x41, 0x52, 0x45, 0x44, 0x5f, 0x43, 0x41, 0x4d,
	0x50, 0x41, 0x49, 0x47, 0x4e, 0x5f, 0x42, 0x55, 0x44, 0x47, 0x45, 0x54, 0x5f, 0x57, 0x49, 0x54,
	0x48, 0x5f, 0x4d, 0x55, 0x4c, 0x54, 0x49, 0x50, 0x4c, 0x45, 0x5f, 0x43, 0x41, 0x4d, 0x50, 0x41,
	0x49, 0x47, 0x4e, 0x53, 0x10, 0x0a, 0x12, 0x12, 0x0a, 0x0e, 0x44, 0x55, 0x50, 0x4c, 0x49, 0x43,
	0x41, 0x54, 0x45, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x0b, 0x12, 0x22, 0x0a, 0x1e, 0x4d, 0x4f,
	0x4e, 0x45, 0x59, 0x5f, 0x41, 0x4d, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x49, 0x4e, 0x5f, 0x57, 0x52,
	0x4f, 0x4e, 0x47, 0x5f, 0x43, 0x55, 0x52, 0x52, 0x45, 0x4e, 0x43, 0x59, 0x10, 0x0c, 0x12, 0x2f,
	0x0a, 0x2b, 0x4d, 0x4f, 0x4e, 0x45, 0x59, 0x5f, 0x41, 0x4d, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x4c,
	0x45, 0x53, 0x53, 0x5f, 0x54, 0x48, 0x41, 0x4e, 0x5f, 0x43, 0x55, 0x52, 0x52, 0x45, 0x4e, 0x43,
	0x59, 0x5f, 0x4d, 0x49, 0x4e, 0x49, 0x4d, 0x55, 0x4d, 0x5f, 0x43, 0x50, 0x43, 0x10, 0x0d, 0x12,
	0x1a, 0x0a, 0x16, 0x4d, 0x4f, 0x4e, 0x45, 0x59, 0x5f, 0x41, 0x4d, 0x4f, 0x55, 0x4e, 0x54, 0x5f,
	0x54, 0x4f, 0x4f, 0x5f, 0x4c, 0x41, 0x52, 0x47, 0x45, 0x10, 0x0e, 0x12, 0x19, 0x0a, 0x15, 0x4e,
	0x45, 0x47, 0x41, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x4d, 0x4f, 0x4e, 0x45, 0x59, 0x5f, 0x41, 0x4d,
	0x4f, 0x55, 0x4e, 0x54, 0x10, 0x0f, 0x12, 0x29, 0x0a, 0x25, 0x4e, 0x4f, 0x4e, 0x5f, 0x4d, 0x55,
	0x4c, 0x54, 0x49, 0x50, 0x4c, 0x45, 0x5f, 0x4f, 0x46, 0x5f, 0x4d, 0x49, 0x4e, 0x49, 0x4d, 0x55,
	0x4d, 0x5f, 0x43, 0x55, 0x52, 0x52, 0x45, 0x4e, 0x43, 0x59, 0x5f, 0x55, 0x4e, 0x49, 0x54, 0x10,
	0x10, 0x12, 0x3d, 0x0a, 0x39, 0x54, 0x4f, 0x54, 0x41, 0x4c, 0x5f, 0x42, 0x55, 0x44, 0x47, 0x45,
	0x54, 0x5f, 0x41, 0x4d, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x4d, 0x55, 0x53, 0x54, 0x5f, 0x42, 0x45,
	0x5f, 0x55, 0x4e, 0x53, 0x45, 0x54, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x42, 0x55, 0x44, 0x47, 0x45,
	0x54, 0x5f, 0x50, 0x45, 0x52, 0x49, 0x4f, 0x44, 0x5f, 0x44, 0x41, 0x49, 0x4c, 0x59, 0x10, 0x12,
	0x12, 0x12, 0x0a, 0x0e, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x50, 0x45, 0x52, 0x49,
	0x4f, 0x44, 0x10, 0x13, 0x12, 0x28, 0x0a, 0x24, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x55,
	0x53, 0x45, 0x5f, 0x41, 0x43, 0x43, 0x45, 0x4c, 0x45, 0x52, 0x41, 0x54, 0x45, 0x44, 0x5f, 0x44,
	0x45, 0x4c, 0x49, 0x56, 0x45, 0x52, 0x59, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x10, 0x14, 0x12, 0x38,
	0x0a, 0x34, 0x42, 0x55, 0x44, 0x47, 0x45, 0x54, 0x5f, 0x41, 0x4d, 0x4f, 0x55, 0x4e, 0x54, 0x5f,
	0x4d, 0x55, 0x53, 0x54, 0x5f, 0x42, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x45, 0x54, 0x5f, 0x46, 0x4f,
	0x52, 0x5f, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x5f, 0x42, 0x55, 0x44, 0x47, 0x45, 0x54, 0x5f,
	0x50, 0x45, 0x52, 0x49, 0x4f, 0x44, 0x10, 0x15, 0x42, 0xf8, 0x01, 0x0a, 0x23, 0x63, 0x6f, 0x6d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0x42, 0x18, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2f, 0x76, 0x31, 0x37, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x3b, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73,
	0x2e, 0x56, 0x31, 0x37, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xca, 0x02, 0x1f, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41,
	0x64, 0x73, 0x5c, 0x56, 0x31, 0x37, 0x5c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xea, 0x02, 0x23,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x37, 0x3a, 0x3a, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v17_errors_campaign_budget_error_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v17_errors_campaign_budget_error_proto_rawDescData = file_google_ads_googleads_v17_errors_campaign_budget_error_proto_rawDesc
)

func file_google_ads_googleads_v17_errors_campaign_budget_error_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v17_errors_campaign_budget_error_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v17_errors_campaign_budget_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v17_errors_campaign_budget_error_proto_rawDescData)
	})
	return file_google_ads_googleads_v17_errors_campaign_budget_error_proto_rawDescData
}

var file_google_ads_googleads_v17_errors_campaign_budget_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v17_errors_campaign_budget_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v17_errors_campaign_budget_error_proto_goTypes = []any{
	(CampaignBudgetErrorEnum_CampaignBudgetError)(0), // 0: google.ads.googleads.v17.errors.CampaignBudgetErrorEnum.CampaignBudgetError
	(*CampaignBudgetErrorEnum)(nil),                  // 1: google.ads.googleads.v17.errors.CampaignBudgetErrorEnum
}
var file_google_ads_googleads_v17_errors_campaign_budget_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v17_errors_campaign_budget_error_proto_init() }
func file_google_ads_googleads_v17_errors_campaign_budget_error_proto_init() {
	if File_google_ads_googleads_v17_errors_campaign_budget_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v17_errors_campaign_budget_error_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CampaignBudgetErrorEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v17_errors_campaign_budget_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v17_errors_campaign_budget_error_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v17_errors_campaign_budget_error_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v17_errors_campaign_budget_error_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v17_errors_campaign_budget_error_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v17_errors_campaign_budget_error_proto = out.File
	file_google_ads_googleads_v17_errors_campaign_budget_error_proto_rawDesc = nil
	file_google_ads_googleads_v17_errors_campaign_budget_error_proto_goTypes = nil
	file_google_ads_googleads_v17_errors_campaign_budget_error_proto_depIdxs = nil
}
