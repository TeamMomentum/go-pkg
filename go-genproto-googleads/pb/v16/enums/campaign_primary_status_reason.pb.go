// Copyright 2023 Google LLC
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
// source: google/ads/googleads/v16/enums/campaign_primary_status_reason.proto

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

// Enum describing the possible campaign primary status reasons.  Provides
// insight into why a campaign is not serving or not serving optimally. These
// reasons are aggregated to determine an overall campaign primary status.
type CampaignPrimaryStatusReasonEnum_CampaignPrimaryStatusReason int32

const (
	// Not specified.
	CampaignPrimaryStatusReasonEnum_UNSPECIFIED CampaignPrimaryStatusReasonEnum_CampaignPrimaryStatusReason = 0
	// Used for return value only. Represents value unknown in this version.
	CampaignPrimaryStatusReasonEnum_UNKNOWN CampaignPrimaryStatusReasonEnum_CampaignPrimaryStatusReason = 1
	// The user-specified campaign status is removed.
	CampaignPrimaryStatusReasonEnum_CAMPAIGN_REMOVED CampaignPrimaryStatusReasonEnum_CampaignPrimaryStatusReason = 2
	// The user-specified campaign status is paused.
	CampaignPrimaryStatusReasonEnum_CAMPAIGN_PAUSED CampaignPrimaryStatusReasonEnum_CampaignPrimaryStatusReason = 3
	// The user-specified time for this campaign to start is in the future.
	CampaignPrimaryStatusReasonEnum_CAMPAIGN_PENDING CampaignPrimaryStatusReasonEnum_CampaignPrimaryStatusReason = 4
	// The user-specified time for this campaign to end has passed.
	CampaignPrimaryStatusReasonEnum_CAMPAIGN_ENDED CampaignPrimaryStatusReasonEnum_CampaignPrimaryStatusReason = 5
	// The campaign is a draft.
	CampaignPrimaryStatusReasonEnum_CAMPAIGN_DRAFT CampaignPrimaryStatusReasonEnum_CampaignPrimaryStatusReason = 6
	// The bidding strategy has incorrect user-specified settings.
	CampaignPrimaryStatusReasonEnum_BIDDING_STRATEGY_MISCONFIGURED CampaignPrimaryStatusReasonEnum_CampaignPrimaryStatusReason = 7
	// The bidding strategy is limited by user-specified settings such as lack
	// of data or similar.
	CampaignPrimaryStatusReasonEnum_BIDDING_STRATEGY_LIMITED CampaignPrimaryStatusReasonEnum_CampaignPrimaryStatusReason = 8
	// The automated bidding system is adjusting to user-specified changes to
	// the bidding strategy.
	CampaignPrimaryStatusReasonEnum_BIDDING_STRATEGY_LEARNING CampaignPrimaryStatusReasonEnum_CampaignPrimaryStatusReason = 9
	// Campaign could capture more conversion value by adjusting CPA/ROAS
	// targets.
	CampaignPrimaryStatusReasonEnum_BIDDING_STRATEGY_CONSTRAINED CampaignPrimaryStatusReasonEnum_CampaignPrimaryStatusReason = 10
	// The budget is limiting the campaign's ability to serve.
	CampaignPrimaryStatusReasonEnum_BUDGET_CONSTRAINED CampaignPrimaryStatusReasonEnum_CampaignPrimaryStatusReason = 11
	// The budget has incorrect user-specified settings.
	CampaignPrimaryStatusReasonEnum_BUDGET_MISCONFIGURED CampaignPrimaryStatusReasonEnum_CampaignPrimaryStatusReason = 12
	// Campaign is not targeting all relevant queries.
	CampaignPrimaryStatusReasonEnum_SEARCH_VOLUME_LIMITED CampaignPrimaryStatusReasonEnum_CampaignPrimaryStatusReason = 13
	// The user-specified ad group statuses are all paused.
	CampaignPrimaryStatusReasonEnum_AD_GROUPS_PAUSED CampaignPrimaryStatusReasonEnum_CampaignPrimaryStatusReason = 14
	// No eligible ad groups exist in this campaign.
	CampaignPrimaryStatusReasonEnum_NO_AD_GROUPS CampaignPrimaryStatusReasonEnum_CampaignPrimaryStatusReason = 15
	// The user-specified keyword statuses are all paused.
	CampaignPrimaryStatusReasonEnum_KEYWORDS_PAUSED CampaignPrimaryStatusReasonEnum_CampaignPrimaryStatusReason = 16
	// No eligible keywords exist in this campaign.
	CampaignPrimaryStatusReasonEnum_NO_KEYWORDS CampaignPrimaryStatusReasonEnum_CampaignPrimaryStatusReason = 17
	// The user-specified ad group ad statuses are all paused.
	CampaignPrimaryStatusReasonEnum_AD_GROUP_ADS_PAUSED CampaignPrimaryStatusReasonEnum_CampaignPrimaryStatusReason = 18
	// No eligible ad group ads exist in this campaign.
	CampaignPrimaryStatusReasonEnum_NO_AD_GROUP_ADS CampaignPrimaryStatusReasonEnum_CampaignPrimaryStatusReason = 19
	// At least one ad in this campaign is limited by policy.
	CampaignPrimaryStatusReasonEnum_HAS_ADS_LIMITED_BY_POLICY CampaignPrimaryStatusReasonEnum_CampaignPrimaryStatusReason = 20
	// At least one ad in this campaign is disapproved.
	CampaignPrimaryStatusReasonEnum_HAS_ADS_DISAPPROVED CampaignPrimaryStatusReasonEnum_CampaignPrimaryStatusReason = 21
	// Most ads in this campaign are pending review.
	CampaignPrimaryStatusReasonEnum_MOST_ADS_UNDER_REVIEW CampaignPrimaryStatusReasonEnum_CampaignPrimaryStatusReason = 22
	// The campaign has a lead form goal, and the lead form extension is
	// missing.
	CampaignPrimaryStatusReasonEnum_MISSING_LEAD_FORM_EXTENSION CampaignPrimaryStatusReasonEnum_CampaignPrimaryStatusReason = 23
	// The campaign has a call goal, and the call extension is missing.
	CampaignPrimaryStatusReasonEnum_MISSING_CALL_EXTENSION CampaignPrimaryStatusReasonEnum_CampaignPrimaryStatusReason = 24
	// The lead form extension is under review.
	CampaignPrimaryStatusReasonEnum_LEAD_FORM_EXTENSION_UNDER_REVIEW CampaignPrimaryStatusReasonEnum_CampaignPrimaryStatusReason = 25
	// The lead extension is disapproved.
	CampaignPrimaryStatusReasonEnum_LEAD_FORM_EXTENSION_DISAPPROVED CampaignPrimaryStatusReasonEnum_CampaignPrimaryStatusReason = 26
	// The call extension is under review.
	CampaignPrimaryStatusReasonEnum_CALL_EXTENSION_UNDER_REVIEW CampaignPrimaryStatusReasonEnum_CampaignPrimaryStatusReason = 27
	// The call extension is disapproved.
	CampaignPrimaryStatusReasonEnum_CALL_EXTENSION_DISAPPROVED CampaignPrimaryStatusReasonEnum_CampaignPrimaryStatusReason = 28
	// No eligible mobile application ad group criteria exist in this campaign.
	CampaignPrimaryStatusReasonEnum_NO_MOBILE_APPLICATION_AD_GROUP_CRITERIA CampaignPrimaryStatusReasonEnum_CampaignPrimaryStatusReason = 29
	// The user-specified campaign group status is paused.
	CampaignPrimaryStatusReasonEnum_CAMPAIGN_GROUP_PAUSED CampaignPrimaryStatusReasonEnum_CampaignPrimaryStatusReason = 30
	// The user-specified times of all group budgets associated with the parent
	// campaign group has passed.
	CampaignPrimaryStatusReasonEnum_CAMPAIGN_GROUP_ALL_GROUP_BUDGETS_ENDED CampaignPrimaryStatusReasonEnum_CampaignPrimaryStatusReason = 31
	// The app associated with this ACi campaign is not released in the target
	// countries of the campaign.
	CampaignPrimaryStatusReasonEnum_APP_NOT_RELEASED CampaignPrimaryStatusReasonEnum_CampaignPrimaryStatusReason = 32
	// The app associated with this ACi campaign is partially released in the
	// target countries of the campaign.
	CampaignPrimaryStatusReasonEnum_APP_PARTIALLY_RELEASED CampaignPrimaryStatusReasonEnum_CampaignPrimaryStatusReason = 33
	// At least one asset group in this campaign is disapproved.
	CampaignPrimaryStatusReasonEnum_HAS_ASSET_GROUPS_DISAPPROVED CampaignPrimaryStatusReasonEnum_CampaignPrimaryStatusReason = 34
	// At least one asset group in this campaign is limited by policy.
	CampaignPrimaryStatusReasonEnum_HAS_ASSET_GROUPS_LIMITED_BY_POLICY CampaignPrimaryStatusReasonEnum_CampaignPrimaryStatusReason = 35
	// Most asset groups in this campaign are pending review.
	CampaignPrimaryStatusReasonEnum_MOST_ASSET_GROUPS_UNDER_REVIEW CampaignPrimaryStatusReasonEnum_CampaignPrimaryStatusReason = 36
	// No eligible asset groups exist in this campaign.
	CampaignPrimaryStatusReasonEnum_NO_ASSET_GROUPS CampaignPrimaryStatusReasonEnum_CampaignPrimaryStatusReason = 37
	// All asset groups in this campaign are paused.
	CampaignPrimaryStatusReasonEnum_ASSET_GROUPS_PAUSED CampaignPrimaryStatusReasonEnum_CampaignPrimaryStatusReason = 38
)

// Enum value maps for CampaignPrimaryStatusReasonEnum_CampaignPrimaryStatusReason.
var (
	CampaignPrimaryStatusReasonEnum_CampaignPrimaryStatusReason_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "CAMPAIGN_REMOVED",
		3:  "CAMPAIGN_PAUSED",
		4:  "CAMPAIGN_PENDING",
		5:  "CAMPAIGN_ENDED",
		6:  "CAMPAIGN_DRAFT",
		7:  "BIDDING_STRATEGY_MISCONFIGURED",
		8:  "BIDDING_STRATEGY_LIMITED",
		9:  "BIDDING_STRATEGY_LEARNING",
		10: "BIDDING_STRATEGY_CONSTRAINED",
		11: "BUDGET_CONSTRAINED",
		12: "BUDGET_MISCONFIGURED",
		13: "SEARCH_VOLUME_LIMITED",
		14: "AD_GROUPS_PAUSED",
		15: "NO_AD_GROUPS",
		16: "KEYWORDS_PAUSED",
		17: "NO_KEYWORDS",
		18: "AD_GROUP_ADS_PAUSED",
		19: "NO_AD_GROUP_ADS",
		20: "HAS_ADS_LIMITED_BY_POLICY",
		21: "HAS_ADS_DISAPPROVED",
		22: "MOST_ADS_UNDER_REVIEW",
		23: "MISSING_LEAD_FORM_EXTENSION",
		24: "MISSING_CALL_EXTENSION",
		25: "LEAD_FORM_EXTENSION_UNDER_REVIEW",
		26: "LEAD_FORM_EXTENSION_DISAPPROVED",
		27: "CALL_EXTENSION_UNDER_REVIEW",
		28: "CALL_EXTENSION_DISAPPROVED",
		29: "NO_MOBILE_APPLICATION_AD_GROUP_CRITERIA",
		30: "CAMPAIGN_GROUP_PAUSED",
		31: "CAMPAIGN_GROUP_ALL_GROUP_BUDGETS_ENDED",
		32: "APP_NOT_RELEASED",
		33: "APP_PARTIALLY_RELEASED",
		34: "HAS_ASSET_GROUPS_DISAPPROVED",
		35: "HAS_ASSET_GROUPS_LIMITED_BY_POLICY",
		36: "MOST_ASSET_GROUPS_UNDER_REVIEW",
		37: "NO_ASSET_GROUPS",
		38: "ASSET_GROUPS_PAUSED",
	}
	CampaignPrimaryStatusReasonEnum_CampaignPrimaryStatusReason_value = map[string]int32{
		"UNSPECIFIED":                             0,
		"UNKNOWN":                                 1,
		"CAMPAIGN_REMOVED":                        2,
		"CAMPAIGN_PAUSED":                         3,
		"CAMPAIGN_PENDING":                        4,
		"CAMPAIGN_ENDED":                          5,
		"CAMPAIGN_DRAFT":                          6,
		"BIDDING_STRATEGY_MISCONFIGURED":          7,
		"BIDDING_STRATEGY_LIMITED":                8,
		"BIDDING_STRATEGY_LEARNING":               9,
		"BIDDING_STRATEGY_CONSTRAINED":            10,
		"BUDGET_CONSTRAINED":                      11,
		"BUDGET_MISCONFIGURED":                    12,
		"SEARCH_VOLUME_LIMITED":                   13,
		"AD_GROUPS_PAUSED":                        14,
		"NO_AD_GROUPS":                            15,
		"KEYWORDS_PAUSED":                         16,
		"NO_KEYWORDS":                             17,
		"AD_GROUP_ADS_PAUSED":                     18,
		"NO_AD_GROUP_ADS":                         19,
		"HAS_ADS_LIMITED_BY_POLICY":               20,
		"HAS_ADS_DISAPPROVED":                     21,
		"MOST_ADS_UNDER_REVIEW":                   22,
		"MISSING_LEAD_FORM_EXTENSION":             23,
		"MISSING_CALL_EXTENSION":                  24,
		"LEAD_FORM_EXTENSION_UNDER_REVIEW":        25,
		"LEAD_FORM_EXTENSION_DISAPPROVED":         26,
		"CALL_EXTENSION_UNDER_REVIEW":             27,
		"CALL_EXTENSION_DISAPPROVED":              28,
		"NO_MOBILE_APPLICATION_AD_GROUP_CRITERIA": 29,
		"CAMPAIGN_GROUP_PAUSED":                   30,
		"CAMPAIGN_GROUP_ALL_GROUP_BUDGETS_ENDED":  31,
		"APP_NOT_RELEASED":                        32,
		"APP_PARTIALLY_RELEASED":                  33,
		"HAS_ASSET_GROUPS_DISAPPROVED":            34,
		"HAS_ASSET_GROUPS_LIMITED_BY_POLICY":      35,
		"MOST_ASSET_GROUPS_UNDER_REVIEW":          36,
		"NO_ASSET_GROUPS":                         37,
		"ASSET_GROUPS_PAUSED":                     38,
	}
)

func (x CampaignPrimaryStatusReasonEnum_CampaignPrimaryStatusReason) Enum() *CampaignPrimaryStatusReasonEnum_CampaignPrimaryStatusReason {
	p := new(CampaignPrimaryStatusReasonEnum_CampaignPrimaryStatusReason)
	*p = x
	return p
}

func (x CampaignPrimaryStatusReasonEnum_CampaignPrimaryStatusReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CampaignPrimaryStatusReasonEnum_CampaignPrimaryStatusReason) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v16_enums_campaign_primary_status_reason_proto_enumTypes[0].Descriptor()
}

func (CampaignPrimaryStatusReasonEnum_CampaignPrimaryStatusReason) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v16_enums_campaign_primary_status_reason_proto_enumTypes[0]
}

func (x CampaignPrimaryStatusReasonEnum_CampaignPrimaryStatusReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CampaignPrimaryStatusReasonEnum_CampaignPrimaryStatusReason.Descriptor instead.
func (CampaignPrimaryStatusReasonEnum_CampaignPrimaryStatusReason) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v16_enums_campaign_primary_status_reason_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible campaign primary status reasons.
type CampaignPrimaryStatusReasonEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CampaignPrimaryStatusReasonEnum) Reset() {
	*x = CampaignPrimaryStatusReasonEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v16_enums_campaign_primary_status_reason_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CampaignPrimaryStatusReasonEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CampaignPrimaryStatusReasonEnum) ProtoMessage() {}

func (x *CampaignPrimaryStatusReasonEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v16_enums_campaign_primary_status_reason_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CampaignPrimaryStatusReasonEnum.ProtoReflect.Descriptor instead.
func (*CampaignPrimaryStatusReasonEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v16_enums_campaign_primary_status_reason_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v16_enums_campaign_primary_status_reason_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v16_enums_campaign_primary_status_reason_proto_rawDesc = []byte{
	0x0a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x36, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x5f, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72,
	0x79, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x36, 0x2e,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x22, 0xf2, 0x08, 0x0a, 0x1f, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69,
	0x67, 0x6e, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0xce, 0x08, 0x0a, 0x1b, 0x43, 0x61,
	0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e,
	0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x41, 0x4d, 0x50, 0x41,
	0x49, 0x47, 0x4e, 0x5f, 0x52, 0x45, 0x4d, 0x4f, 0x56, 0x45, 0x44, 0x10, 0x02, 0x12, 0x13, 0x0a,
	0x0f, 0x43, 0x41, 0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e, 0x5f, 0x50, 0x41, 0x55, 0x53, 0x45, 0x44,
	0x10, 0x03, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x41, 0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e, 0x5f, 0x50,
	0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x04, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x41, 0x4d, 0x50,
	0x41, 0x49, 0x47, 0x4e, 0x5f, 0x45, 0x4e, 0x44, 0x45, 0x44, 0x10, 0x05, 0x12, 0x12, 0x0a, 0x0e,
	0x43, 0x41, 0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e, 0x5f, 0x44, 0x52, 0x41, 0x46, 0x54, 0x10, 0x06,
	0x12, 0x22, 0x0a, 0x1e, 0x42, 0x49, 0x44, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x54, 0x52, 0x41,
	0x54, 0x45, 0x47, 0x59, 0x5f, 0x4d, 0x49, 0x53, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x55, 0x52,
	0x45, 0x44, 0x10, 0x07, 0x12, 0x1c, 0x0a, 0x18, 0x42, 0x49, 0x44, 0x44, 0x49, 0x4e, 0x47, 0x5f,
	0x53, 0x54, 0x52, 0x41, 0x54, 0x45, 0x47, 0x59, 0x5f, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x45, 0x44,
	0x10, 0x08, 0x12, 0x1d, 0x0a, 0x19, 0x42, 0x49, 0x44, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x54,
	0x52, 0x41, 0x54, 0x45, 0x47, 0x59, 0x5f, 0x4c, 0x45, 0x41, 0x52, 0x4e, 0x49, 0x4e, 0x47, 0x10,
	0x09, 0x12, 0x20, 0x0a, 0x1c, 0x42, 0x49, 0x44, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x54, 0x52,
	0x41, 0x54, 0x45, 0x47, 0x59, 0x5f, 0x43, 0x4f, 0x4e, 0x53, 0x54, 0x52, 0x41, 0x49, 0x4e, 0x45,
	0x44, 0x10, 0x0a, 0x12, 0x16, 0x0a, 0x12, 0x42, 0x55, 0x44, 0x47, 0x45, 0x54, 0x5f, 0x43, 0x4f,
	0x4e, 0x53, 0x54, 0x52, 0x41, 0x49, 0x4e, 0x45, 0x44, 0x10, 0x0b, 0x12, 0x18, 0x0a, 0x14, 0x42,
	0x55, 0x44, 0x47, 0x45, 0x54, 0x5f, 0x4d, 0x49, 0x53, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x55,
	0x52, 0x45, 0x44, 0x10, 0x0c, 0x12, 0x19, 0x0a, 0x15, 0x53, 0x45, 0x41, 0x52, 0x43, 0x48, 0x5f,
	0x56, 0x4f, 0x4c, 0x55, 0x4d, 0x45, 0x5f, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x45, 0x44, 0x10, 0x0d,
	0x12, 0x14, 0x0a, 0x10, 0x41, 0x44, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x53, 0x5f, 0x50, 0x41,
	0x55, 0x53, 0x45, 0x44, 0x10, 0x0e, 0x12, 0x10, 0x0a, 0x0c, 0x4e, 0x4f, 0x5f, 0x41, 0x44, 0x5f,
	0x47, 0x52, 0x4f, 0x55, 0x50, 0x53, 0x10, 0x0f, 0x12, 0x13, 0x0a, 0x0f, 0x4b, 0x45, 0x59, 0x57,
	0x4f, 0x52, 0x44, 0x53, 0x5f, 0x50, 0x41, 0x55, 0x53, 0x45, 0x44, 0x10, 0x10, 0x12, 0x0f, 0x0a,
	0x0b, 0x4e, 0x4f, 0x5f, 0x4b, 0x45, 0x59, 0x57, 0x4f, 0x52, 0x44, 0x53, 0x10, 0x11, 0x12, 0x17,
	0x0a, 0x13, 0x41, 0x44, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x41, 0x44, 0x53, 0x5f, 0x50,
	0x41, 0x55, 0x53, 0x45, 0x44, 0x10, 0x12, 0x12, 0x13, 0x0a, 0x0f, 0x4e, 0x4f, 0x5f, 0x41, 0x44,
	0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x41, 0x44, 0x53, 0x10, 0x13, 0x12, 0x1d, 0x0a, 0x19,
	0x48, 0x41, 0x53, 0x5f, 0x41, 0x44, 0x53, 0x5f, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x45, 0x44, 0x5f,
	0x42, 0x59, 0x5f, 0x50, 0x4f, 0x4c, 0x49, 0x43, 0x59, 0x10, 0x14, 0x12, 0x17, 0x0a, 0x13, 0x48,
	0x41, 0x53, 0x5f, 0x41, 0x44, 0x53, 0x5f, 0x44, 0x49, 0x53, 0x41, 0x50, 0x50, 0x52, 0x4f, 0x56,
	0x45, 0x44, 0x10, 0x15, 0x12, 0x19, 0x0a, 0x15, 0x4d, 0x4f, 0x53, 0x54, 0x5f, 0x41, 0x44, 0x53,
	0x5f, 0x55, 0x4e, 0x44, 0x45, 0x52, 0x5f, 0x52, 0x45, 0x56, 0x49, 0x45, 0x57, 0x10, 0x16, 0x12,
	0x1f, 0x0a, 0x1b, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x5f, 0x4c, 0x45, 0x41, 0x44, 0x5f,
	0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x45, 0x58, 0x54, 0x45, 0x4e, 0x53, 0x49, 0x4f, 0x4e, 0x10, 0x17,
	0x12, 0x1a, 0x0a, 0x16, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x5f, 0x43, 0x41, 0x4c, 0x4c,
	0x5f, 0x45, 0x58, 0x54, 0x45, 0x4e, 0x53, 0x49, 0x4f, 0x4e, 0x10, 0x18, 0x12, 0x24, 0x0a, 0x20,
	0x4c, 0x45, 0x41, 0x44, 0x5f, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x45, 0x58, 0x54, 0x45, 0x4e, 0x53,
	0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x44, 0x45, 0x52, 0x5f, 0x52, 0x45, 0x56, 0x49, 0x45, 0x57,
	0x10, 0x19, 0x12, 0x23, 0x0a, 0x1f, 0x4c, 0x45, 0x41, 0x44, 0x5f, 0x46, 0x4f, 0x52, 0x4d, 0x5f,
	0x45, 0x58, 0x54, 0x45, 0x4e, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x44, 0x49, 0x53, 0x41, 0x50, 0x50,
	0x52, 0x4f, 0x56, 0x45, 0x44, 0x10, 0x1a, 0x12, 0x1f, 0x0a, 0x1b, 0x43, 0x41, 0x4c, 0x4c, 0x5f,
	0x45, 0x58, 0x54, 0x45, 0x4e, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x44, 0x45, 0x52, 0x5f,
	0x52, 0x45, 0x56, 0x49, 0x45, 0x57, 0x10, 0x1b, 0x12, 0x1e, 0x0a, 0x1a, 0x43, 0x41, 0x4c, 0x4c,
	0x5f, 0x45, 0x58, 0x54, 0x45, 0x4e, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x44, 0x49, 0x53, 0x41, 0x50,
	0x50, 0x52, 0x4f, 0x56, 0x45, 0x44, 0x10, 0x1c, 0x12, 0x2b, 0x0a, 0x27, 0x4e, 0x4f, 0x5f, 0x4d,
	0x4f, 0x42, 0x49, 0x4c, 0x45, 0x5f, 0x41, 0x50, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x41, 0x44, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x43, 0x52, 0x49, 0x54, 0x45,
	0x52, 0x49, 0x41, 0x10, 0x1d, 0x12, 0x19, 0x0a, 0x15, 0x43, 0x41, 0x4d, 0x50, 0x41, 0x49, 0x47,
	0x4e, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x50, 0x41, 0x55, 0x53, 0x45, 0x44, 0x10, 0x1e,
	0x12, 0x2a, 0x0a, 0x26, 0x43, 0x41, 0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e, 0x5f, 0x47, 0x52, 0x4f,
	0x55, 0x50, 0x5f, 0x41, 0x4c, 0x4c, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x42, 0x55, 0x44,
	0x47, 0x45, 0x54, 0x53, 0x5f, 0x45, 0x4e, 0x44, 0x45, 0x44, 0x10, 0x1f, 0x12, 0x14, 0x0a, 0x10,
	0x41, 0x50, 0x50, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x52, 0x45, 0x4c, 0x45, 0x41, 0x53, 0x45, 0x44,
	0x10, 0x20, 0x12, 0x1a, 0x0a, 0x16, 0x41, 0x50, 0x50, 0x5f, 0x50, 0x41, 0x52, 0x54, 0x49, 0x41,
	0x4c, 0x4c, 0x59, 0x5f, 0x52, 0x45, 0x4c, 0x45, 0x41, 0x53, 0x45, 0x44, 0x10, 0x21, 0x12, 0x20,
	0x0a, 0x1c, 0x48, 0x41, 0x53, 0x5f, 0x41, 0x53, 0x53, 0x45, 0x54, 0x5f, 0x47, 0x52, 0x4f, 0x55,
	0x50, 0x53, 0x5f, 0x44, 0x49, 0x53, 0x41, 0x50, 0x50, 0x52, 0x4f, 0x56, 0x45, 0x44, 0x10, 0x22,
	0x12, 0x26, 0x0a, 0x22, 0x48, 0x41, 0x53, 0x5f, 0x41, 0x53, 0x53, 0x45, 0x54, 0x5f, 0x47, 0x52,
	0x4f, 0x55, 0x50, 0x53, 0x5f, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x45, 0x44, 0x5f, 0x42, 0x59, 0x5f,
	0x50, 0x4f, 0x4c, 0x49, 0x43, 0x59, 0x10, 0x23, 0x12, 0x22, 0x0a, 0x1e, 0x4d, 0x4f, 0x53, 0x54,
	0x5f, 0x41, 0x53, 0x53, 0x45, 0x54, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x53, 0x5f, 0x55, 0x4e,
	0x44, 0x45, 0x52, 0x5f, 0x52, 0x45, 0x56, 0x49, 0x45, 0x57, 0x10, 0x24, 0x12, 0x13, 0x0a, 0x0f,
	0x4e, 0x4f, 0x5f, 0x41, 0x53, 0x53, 0x45, 0x54, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x53, 0x10,
	0x25, 0x12, 0x17, 0x0a, 0x13, 0x41, 0x53, 0x53, 0x45, 0x54, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50,
	0x53, 0x5f, 0x50, 0x41, 0x55, 0x53, 0x45, 0x44, 0x10, 0x26, 0x42, 0xfa, 0x01, 0x0a, 0x22, 0x63,
	0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x36, 0x2e, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x42, 0x20, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x50, 0x72, 0x69, 0x6d, 0x61,
	0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f,
	0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x36, 0x2f, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x3b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41,
	0xaa, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x36, 0x2e, 0x45, 0x6e, 0x75, 0x6d,
	0x73, 0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x36, 0x5c, 0x45, 0x6e, 0x75,
	0x6d, 0x73, 0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73,
	0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x36,
	0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v16_enums_campaign_primary_status_reason_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v16_enums_campaign_primary_status_reason_proto_rawDescData = file_google_ads_googleads_v16_enums_campaign_primary_status_reason_proto_rawDesc
)

func file_google_ads_googleads_v16_enums_campaign_primary_status_reason_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v16_enums_campaign_primary_status_reason_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v16_enums_campaign_primary_status_reason_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v16_enums_campaign_primary_status_reason_proto_rawDescData)
	})
	return file_google_ads_googleads_v16_enums_campaign_primary_status_reason_proto_rawDescData
}

var file_google_ads_googleads_v16_enums_campaign_primary_status_reason_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v16_enums_campaign_primary_status_reason_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v16_enums_campaign_primary_status_reason_proto_goTypes = []any{
	(CampaignPrimaryStatusReasonEnum_CampaignPrimaryStatusReason)(0), // 0: google.ads.googleads.v16.enums.CampaignPrimaryStatusReasonEnum.CampaignPrimaryStatusReason
	(*CampaignPrimaryStatusReasonEnum)(nil),                          // 1: google.ads.googleads.v16.enums.CampaignPrimaryStatusReasonEnum
}
var file_google_ads_googleads_v16_enums_campaign_primary_status_reason_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v16_enums_campaign_primary_status_reason_proto_init() }
func file_google_ads_googleads_v16_enums_campaign_primary_status_reason_proto_init() {
	if File_google_ads_googleads_v16_enums_campaign_primary_status_reason_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v16_enums_campaign_primary_status_reason_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CampaignPrimaryStatusReasonEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v16_enums_campaign_primary_status_reason_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v16_enums_campaign_primary_status_reason_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v16_enums_campaign_primary_status_reason_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v16_enums_campaign_primary_status_reason_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v16_enums_campaign_primary_status_reason_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v16_enums_campaign_primary_status_reason_proto = out.File
	file_google_ads_googleads_v16_enums_campaign_primary_status_reason_proto_rawDesc = nil
	file_google_ads_googleads_v16_enums_campaign_primary_status_reason_proto_goTypes = nil
	file_google_ads_googleads_v16_enums_campaign_primary_status_reason_proto_depIdxs = nil
}
