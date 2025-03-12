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
// source: google/ads/googleads/v18/enums/recommendation_type.proto

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

// Types of recommendations.
type RecommendationTypeEnum_RecommendationType int32

const (
	// Not specified.
	RecommendationTypeEnum_UNSPECIFIED RecommendationTypeEnum_RecommendationType = 0
	// Used for return value only. Represents value unknown in this version.
	RecommendationTypeEnum_UNKNOWN RecommendationTypeEnum_RecommendationType = 1
	// Budget recommendation for campaigns that are currently budget-constrained
	// (as opposed to the FORECASTING_CAMPAIGN_BUDGET recommendation, which
	// applies to campaigns that are expected to become budget-constrained in
	// the future).
	RecommendationTypeEnum_CAMPAIGN_BUDGET RecommendationTypeEnum_RecommendationType = 2
	// Keyword recommendation.
	RecommendationTypeEnum_KEYWORD RecommendationTypeEnum_RecommendationType = 3
	// Recommendation to add a new text ad.
	RecommendationTypeEnum_TEXT_AD RecommendationTypeEnum_RecommendationType = 4
	// Recommendation to update a campaign to use a Target CPA bidding strategy.
	RecommendationTypeEnum_TARGET_CPA_OPT_IN RecommendationTypeEnum_RecommendationType = 5
	// Recommendation to update a campaign to use the Maximize Conversions
	// bidding strategy.
	RecommendationTypeEnum_MAXIMIZE_CONVERSIONS_OPT_IN RecommendationTypeEnum_RecommendationType = 6
	// Recommendation to enable Enhanced Cost Per Click for a campaign.
	RecommendationTypeEnum_ENHANCED_CPC_OPT_IN RecommendationTypeEnum_RecommendationType = 7
	// Recommendation to start showing your campaign's ads on Google Search
	// Partners Websites.
	RecommendationTypeEnum_SEARCH_PARTNERS_OPT_IN RecommendationTypeEnum_RecommendationType = 8
	// Recommendation to update a campaign to use a Maximize Clicks bidding
	// strategy.
	RecommendationTypeEnum_MAXIMIZE_CLICKS_OPT_IN RecommendationTypeEnum_RecommendationType = 9
	// Recommendation to start using the "Optimize" ad rotation setting for the
	// given ad group.
	RecommendationTypeEnum_OPTIMIZE_AD_ROTATION RecommendationTypeEnum_RecommendationType = 10
	// Recommendation to change an existing keyword from one match type to a
	// broader match type.
	RecommendationTypeEnum_KEYWORD_MATCH_TYPE RecommendationTypeEnum_RecommendationType = 14
	// Recommendation to move unused budget from one budget to a constrained
	// budget.
	RecommendationTypeEnum_MOVE_UNUSED_BUDGET RecommendationTypeEnum_RecommendationType = 15
	// Budget recommendation for campaigns that are expected to become
	// budget-constrained in the future (as opposed to the CAMPAIGN_BUDGET
	// recommendation, which applies to campaigns that are currently
	// budget-constrained).
	RecommendationTypeEnum_FORECASTING_CAMPAIGN_BUDGET RecommendationTypeEnum_RecommendationType = 16
	// Recommendation to update a campaign to use a Target ROAS bidding
	// strategy.
	RecommendationTypeEnum_TARGET_ROAS_OPT_IN RecommendationTypeEnum_RecommendationType = 17
	// Recommendation to add a new responsive search ad.
	RecommendationTypeEnum_RESPONSIVE_SEARCH_AD RecommendationTypeEnum_RecommendationType = 18
	// Budget recommendation for campaigns whose ROI is predicted to increase
	// with a budget adjustment.
	RecommendationTypeEnum_MARGINAL_ROI_CAMPAIGN_BUDGET RecommendationTypeEnum_RecommendationType = 19
	// Recommendation to add broad match versions of keywords for fully
	// automated conversion-based bidding campaigns.
	RecommendationTypeEnum_USE_BROAD_MATCH_KEYWORD RecommendationTypeEnum_RecommendationType = 20
	// Recommendation to add new responsive search ad assets.
	RecommendationTypeEnum_RESPONSIVE_SEARCH_AD_ASSET RecommendationTypeEnum_RecommendationType = 21
	// Recommendation to upgrade a Smart Shopping campaign to a Performance Max
	// campaign.
	RecommendationTypeEnum_UPGRADE_SMART_SHOPPING_CAMPAIGN_TO_PERFORMANCE_MAX RecommendationTypeEnum_RecommendationType = 22
	// Recommendation to improve strength of responsive search ad.
	RecommendationTypeEnum_RESPONSIVE_SEARCH_AD_IMPROVE_AD_STRENGTH RecommendationTypeEnum_RecommendationType = 23
	// Recommendation to update a campaign to use Display Expansion.
	RecommendationTypeEnum_DISPLAY_EXPANSION_OPT_IN RecommendationTypeEnum_RecommendationType = 24
	// Recommendation to upgrade a Local campaign to a Performance Max
	// campaign.
	RecommendationTypeEnum_UPGRADE_LOCAL_CAMPAIGN_TO_PERFORMANCE_MAX RecommendationTypeEnum_RecommendationType = 25
	// Recommendation to raise target CPA when it is too low and there are very
	// few or no conversions.
	// It is applied asynchronously and can take minutes
	// depending on the number of ad groups there are in the related campaign.
	RecommendationTypeEnum_RAISE_TARGET_CPA_BID_TOO_LOW RecommendationTypeEnum_RecommendationType = 26
	// Recommendation to raise the budget in advance of a seasonal event that is
	// forecasted to increase traffic, and change bidding strategy from maximize
	// conversion value to target ROAS.
	RecommendationTypeEnum_FORECASTING_SET_TARGET_ROAS RecommendationTypeEnum_RecommendationType = 27
	// Recommendation to add callout assets to campaign or customer level.
	RecommendationTypeEnum_CALLOUT_ASSET RecommendationTypeEnum_RecommendationType = 28
	// Recommendation to add sitelink assets to campaign or customer level.
	RecommendationTypeEnum_SITELINK_ASSET RecommendationTypeEnum_RecommendationType = 29
	// Recommendation to add call assets to campaign or customer level.
	RecommendationTypeEnum_CALL_ASSET RecommendationTypeEnum_RecommendationType = 30
	// Recommendation to add the age group attribute to offers that are
	// demoted because of a missing age group.
	RecommendationTypeEnum_SHOPPING_ADD_AGE_GROUP RecommendationTypeEnum_RecommendationType = 31
	// Recommendation to add a color to offers that are demoted
	// because of a missing color.
	RecommendationTypeEnum_SHOPPING_ADD_COLOR RecommendationTypeEnum_RecommendationType = 32
	// Recommendation to add a gender to offers that are demoted
	// because of a missing gender.
	RecommendationTypeEnum_SHOPPING_ADD_GENDER RecommendationTypeEnum_RecommendationType = 33
	// Recommendation to add a GTIN (Global Trade Item Number) to offers
	// that are demoted because of a missing GTIN.
	RecommendationTypeEnum_SHOPPING_ADD_GTIN RecommendationTypeEnum_RecommendationType = 34
	// Recommendation to add more identifiers to offers that are demoted because
	// of missing identifiers.
	RecommendationTypeEnum_SHOPPING_ADD_MORE_IDENTIFIERS RecommendationTypeEnum_RecommendationType = 35
	// Recommendation to add the size to offers that are demoted
	// because of a missing size.
	RecommendationTypeEnum_SHOPPING_ADD_SIZE RecommendationTypeEnum_RecommendationType = 36
	// Recommendation informing a customer about a campaign that cannot serve
	// because no products are being targeted.
	RecommendationTypeEnum_SHOPPING_ADD_PRODUCTS_TO_CAMPAIGN RecommendationTypeEnum_RecommendationType = 37
	// The shopping recommendation informing a customer about campaign with a
	// high percentage of disapproved products.
	RecommendationTypeEnum_SHOPPING_FIX_DISAPPROVED_PRODUCTS RecommendationTypeEnum_RecommendationType = 38
	// Recommendation to create a catch-all campaign that targets all offers.
	RecommendationTypeEnum_SHOPPING_TARGET_ALL_OFFERS RecommendationTypeEnum_RecommendationType = 39
	// Recommendation to fix Merchant Center account suspension issues.
	RecommendationTypeEnum_SHOPPING_FIX_SUSPENDED_MERCHANT_CENTER_ACCOUNT RecommendationTypeEnum_RecommendationType = 40
	// Recommendation to fix Merchant Center account suspension warning issues.
	RecommendationTypeEnum_SHOPPING_FIX_MERCHANT_CENTER_ACCOUNT_SUSPENSION_WARNING RecommendationTypeEnum_RecommendationType = 41
	// Recommendation to migrate offers targeted by Regular Shopping Campaigns
	// to existing Performance Max campaigns.
	RecommendationTypeEnum_SHOPPING_MIGRATE_REGULAR_SHOPPING_CAMPAIGN_OFFERS_TO_PERFORMANCE_MAX RecommendationTypeEnum_RecommendationType = 42
	// Recommendation to enable dynamic image extensions on the account,
	// allowing Google to find the best images from ad landing pages and
	// complement text ads.
	RecommendationTypeEnum_DYNAMIC_IMAGE_EXTENSION_OPT_IN RecommendationTypeEnum_RecommendationType = 43
	// Recommendation to raise Target CPA based on Google predictions modeled
	// from past conversions. It is applied asynchronously and can take minutes
	// depending on the number of ad groups there are in the related campaign.
	RecommendationTypeEnum_RAISE_TARGET_CPA RecommendationTypeEnum_RecommendationType = 44
	// Recommendation to lower Target ROAS.
	RecommendationTypeEnum_LOWER_TARGET_ROAS RecommendationTypeEnum_RecommendationType = 45
	// Recommendation to opt into Performance Max campaigns.
	RecommendationTypeEnum_PERFORMANCE_MAX_OPT_IN RecommendationTypeEnum_RecommendationType = 46
	// Recommendation to improve the asset group strength of a Performance Max
	// campaign to an "Excellent" rating.
	RecommendationTypeEnum_IMPROVE_PERFORMANCE_MAX_AD_STRENGTH RecommendationTypeEnum_RecommendationType = 47
	// Recommendation to migrate Dynamic Search Ads to Performance Max
	// campaigns.
	RecommendationTypeEnum_MIGRATE_DYNAMIC_SEARCH_ADS_CAMPAIGN_TO_PERFORMANCE_MAX RecommendationTypeEnum_RecommendationType = 48
	// Recommendation to set a target CPA for campaigns that do not have one
	// specified, in advance of a seasonal event that is forecasted to increase
	// traffic.
	RecommendationTypeEnum_FORECASTING_SET_TARGET_CPA RecommendationTypeEnum_RecommendationType = 49
	// Recommendation to set a target CPA for campaigns that do not have one
	// specified.
	RecommendationTypeEnum_SET_TARGET_CPA RecommendationTypeEnum_RecommendationType = 50
	// Recommendation to set a target ROAS for campaigns that do not have one
	// specified.
	RecommendationTypeEnum_SET_TARGET_ROAS RecommendationTypeEnum_RecommendationType = 51
	// Recommendation to update a campaign to use the Maximize Conversion Value
	// bidding strategy.
	RecommendationTypeEnum_MAXIMIZE_CONVERSION_VALUE_OPT_IN RecommendationTypeEnum_RecommendationType = 52
	// Recommendation to deploy Google Tag on more pages.
	RecommendationTypeEnum_IMPROVE_GOOGLE_TAG_COVERAGE RecommendationTypeEnum_RecommendationType = 53
	// Recommendation to turn on Final URL expansion for your Performance Max
	// campaigns.
	RecommendationTypeEnum_PERFORMANCE_MAX_FINAL_URL_OPT_IN RecommendationTypeEnum_RecommendationType = 54
	// Recommendation to update a customer list that hasn't been updated
	// in the last 90 days.
	RecommendationTypeEnum_REFRESH_CUSTOMER_MATCH_LIST RecommendationTypeEnum_RecommendationType = 55
	// Recommendation to create a custom audience.
	RecommendationTypeEnum_CUSTOM_AUDIENCE_OPT_IN RecommendationTypeEnum_RecommendationType = 56
	// Recommendation to add lead form assets to campaign or customer level.
	RecommendationTypeEnum_LEAD_FORM_ASSET RecommendationTypeEnum_RecommendationType = 57
	// Recommendation to improve the strength of ads in
	// Demand Gen campaigns.
	RecommendationTypeEnum_IMPROVE_DEMAND_GEN_AD_STRENGTH RecommendationTypeEnum_RecommendationType = 58
)

// Enum value maps for RecommendationTypeEnum_RecommendationType.
var (
	RecommendationTypeEnum_RecommendationType_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "CAMPAIGN_BUDGET",
		3:  "KEYWORD",
		4:  "TEXT_AD",
		5:  "TARGET_CPA_OPT_IN",
		6:  "MAXIMIZE_CONVERSIONS_OPT_IN",
		7:  "ENHANCED_CPC_OPT_IN",
		8:  "SEARCH_PARTNERS_OPT_IN",
		9:  "MAXIMIZE_CLICKS_OPT_IN",
		10: "OPTIMIZE_AD_ROTATION",
		14: "KEYWORD_MATCH_TYPE",
		15: "MOVE_UNUSED_BUDGET",
		16: "FORECASTING_CAMPAIGN_BUDGET",
		17: "TARGET_ROAS_OPT_IN",
		18: "RESPONSIVE_SEARCH_AD",
		19: "MARGINAL_ROI_CAMPAIGN_BUDGET",
		20: "USE_BROAD_MATCH_KEYWORD",
		21: "RESPONSIVE_SEARCH_AD_ASSET",
		22: "UPGRADE_SMART_SHOPPING_CAMPAIGN_TO_PERFORMANCE_MAX",
		23: "RESPONSIVE_SEARCH_AD_IMPROVE_AD_STRENGTH",
		24: "DISPLAY_EXPANSION_OPT_IN",
		25: "UPGRADE_LOCAL_CAMPAIGN_TO_PERFORMANCE_MAX",
		26: "RAISE_TARGET_CPA_BID_TOO_LOW",
		27: "FORECASTING_SET_TARGET_ROAS",
		28: "CALLOUT_ASSET",
		29: "SITELINK_ASSET",
		30: "CALL_ASSET",
		31: "SHOPPING_ADD_AGE_GROUP",
		32: "SHOPPING_ADD_COLOR",
		33: "SHOPPING_ADD_GENDER",
		34: "SHOPPING_ADD_GTIN",
		35: "SHOPPING_ADD_MORE_IDENTIFIERS",
		36: "SHOPPING_ADD_SIZE",
		37: "SHOPPING_ADD_PRODUCTS_TO_CAMPAIGN",
		38: "SHOPPING_FIX_DISAPPROVED_PRODUCTS",
		39: "SHOPPING_TARGET_ALL_OFFERS",
		40: "SHOPPING_FIX_SUSPENDED_MERCHANT_CENTER_ACCOUNT",
		41: "SHOPPING_FIX_MERCHANT_CENTER_ACCOUNT_SUSPENSION_WARNING",
		42: "SHOPPING_MIGRATE_REGULAR_SHOPPING_CAMPAIGN_OFFERS_TO_PERFORMANCE_MAX",
		43: "DYNAMIC_IMAGE_EXTENSION_OPT_IN",
		44: "RAISE_TARGET_CPA",
		45: "LOWER_TARGET_ROAS",
		46: "PERFORMANCE_MAX_OPT_IN",
		47: "IMPROVE_PERFORMANCE_MAX_AD_STRENGTH",
		48: "MIGRATE_DYNAMIC_SEARCH_ADS_CAMPAIGN_TO_PERFORMANCE_MAX",
		49: "FORECASTING_SET_TARGET_CPA",
		50: "SET_TARGET_CPA",
		51: "SET_TARGET_ROAS",
		52: "MAXIMIZE_CONVERSION_VALUE_OPT_IN",
		53: "IMPROVE_GOOGLE_TAG_COVERAGE",
		54: "PERFORMANCE_MAX_FINAL_URL_OPT_IN",
		55: "REFRESH_CUSTOMER_MATCH_LIST",
		56: "CUSTOM_AUDIENCE_OPT_IN",
		57: "LEAD_FORM_ASSET",
		58: "IMPROVE_DEMAND_GEN_AD_STRENGTH",
	}
	RecommendationTypeEnum_RecommendationType_value = map[string]int32{
		"UNSPECIFIED":                  0,
		"UNKNOWN":                      1,
		"CAMPAIGN_BUDGET":              2,
		"KEYWORD":                      3,
		"TEXT_AD":                      4,
		"TARGET_CPA_OPT_IN":            5,
		"MAXIMIZE_CONVERSIONS_OPT_IN":  6,
		"ENHANCED_CPC_OPT_IN":          7,
		"SEARCH_PARTNERS_OPT_IN":       8,
		"MAXIMIZE_CLICKS_OPT_IN":       9,
		"OPTIMIZE_AD_ROTATION":         10,
		"KEYWORD_MATCH_TYPE":           14,
		"MOVE_UNUSED_BUDGET":           15,
		"FORECASTING_CAMPAIGN_BUDGET":  16,
		"TARGET_ROAS_OPT_IN":           17,
		"RESPONSIVE_SEARCH_AD":         18,
		"MARGINAL_ROI_CAMPAIGN_BUDGET": 19,
		"USE_BROAD_MATCH_KEYWORD":      20,
		"RESPONSIVE_SEARCH_AD_ASSET":   21,
		"UPGRADE_SMART_SHOPPING_CAMPAIGN_TO_PERFORMANCE_MAX":      22,
		"RESPONSIVE_SEARCH_AD_IMPROVE_AD_STRENGTH":                23,
		"DISPLAY_EXPANSION_OPT_IN":                                24,
		"UPGRADE_LOCAL_CAMPAIGN_TO_PERFORMANCE_MAX":               25,
		"RAISE_TARGET_CPA_BID_TOO_LOW":                            26,
		"FORECASTING_SET_TARGET_ROAS":                             27,
		"CALLOUT_ASSET":                                           28,
		"SITELINK_ASSET":                                          29,
		"CALL_ASSET":                                              30,
		"SHOPPING_ADD_AGE_GROUP":                                  31,
		"SHOPPING_ADD_COLOR":                                      32,
		"SHOPPING_ADD_GENDER":                                     33,
		"SHOPPING_ADD_GTIN":                                       34,
		"SHOPPING_ADD_MORE_IDENTIFIERS":                           35,
		"SHOPPING_ADD_SIZE":                                       36,
		"SHOPPING_ADD_PRODUCTS_TO_CAMPAIGN":                       37,
		"SHOPPING_FIX_DISAPPROVED_PRODUCTS":                       38,
		"SHOPPING_TARGET_ALL_OFFERS":                              39,
		"SHOPPING_FIX_SUSPENDED_MERCHANT_CENTER_ACCOUNT":          40,
		"SHOPPING_FIX_MERCHANT_CENTER_ACCOUNT_SUSPENSION_WARNING": 41,
		"SHOPPING_MIGRATE_REGULAR_SHOPPING_CAMPAIGN_OFFERS_TO_PERFORMANCE_MAX": 42,
		"DYNAMIC_IMAGE_EXTENSION_OPT_IN":                                       43,
		"RAISE_TARGET_CPA":                                                     44,
		"LOWER_TARGET_ROAS":                                                    45,
		"PERFORMANCE_MAX_OPT_IN":                                               46,
		"IMPROVE_PERFORMANCE_MAX_AD_STRENGTH":                                  47,
		"MIGRATE_DYNAMIC_SEARCH_ADS_CAMPAIGN_TO_PERFORMANCE_MAX":               48,
		"FORECASTING_SET_TARGET_CPA":                                           49,
		"SET_TARGET_CPA":                                                       50,
		"SET_TARGET_ROAS":                                                      51,
		"MAXIMIZE_CONVERSION_VALUE_OPT_IN":                                     52,
		"IMPROVE_GOOGLE_TAG_COVERAGE":                                          53,
		"PERFORMANCE_MAX_FINAL_URL_OPT_IN":                                     54,
		"REFRESH_CUSTOMER_MATCH_LIST":                                          55,
		"CUSTOM_AUDIENCE_OPT_IN":                                               56,
		"LEAD_FORM_ASSET":                                                      57,
		"IMPROVE_DEMAND_GEN_AD_STRENGTH":                                       58,
	}
)

func (x RecommendationTypeEnum_RecommendationType) Enum() *RecommendationTypeEnum_RecommendationType {
	p := new(RecommendationTypeEnum_RecommendationType)
	*p = x
	return p
}

func (x RecommendationTypeEnum_RecommendationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RecommendationTypeEnum_RecommendationType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v18_enums_recommendation_type_proto_enumTypes[0].Descriptor()
}

func (RecommendationTypeEnum_RecommendationType) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v18_enums_recommendation_type_proto_enumTypes[0]
}

func (x RecommendationTypeEnum_RecommendationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RecommendationTypeEnum_RecommendationType.Descriptor instead.
func (RecommendationTypeEnum_RecommendationType) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v18_enums_recommendation_type_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing types of recommendations.
type RecommendationTypeEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RecommendationTypeEnum) Reset() {
	*x = RecommendationTypeEnum{}
	mi := &file_google_ads_googleads_v18_enums_recommendation_type_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RecommendationTypeEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecommendationTypeEnum) ProtoMessage() {}

func (x *RecommendationTypeEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v18_enums_recommendation_type_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecommendationTypeEnum.ProtoReflect.Descriptor instead.
func (*RecommendationTypeEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v18_enums_recommendation_type_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v18_enums_recommendation_type_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v18_enums_recommendation_type_proto_rawDesc = []byte{
	0x0a, 0x38, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x38, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x31, 0x38, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x22, 0xf3, 0x0d, 0x0a, 0x16, 0x52,
	0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0xd8, 0x0d, 0x0a, 0x12, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a,
	0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x43, 0x41,
	0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e, 0x5f, 0x42, 0x55, 0x44, 0x47, 0x45, 0x54, 0x10, 0x02, 0x12,
	0x0b, 0x0a, 0x07, 0x4b, 0x45, 0x59, 0x57, 0x4f, 0x52, 0x44, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07,
	0x54, 0x45, 0x58, 0x54, 0x5f, 0x41, 0x44, 0x10, 0x04, 0x12, 0x15, 0x0a, 0x11, 0x54, 0x41, 0x52,
	0x47, 0x45, 0x54, 0x5f, 0x43, 0x50, 0x41, 0x5f, 0x4f, 0x50, 0x54, 0x5f, 0x49, 0x4e, 0x10, 0x05,
	0x12, 0x1f, 0x0a, 0x1b, 0x4d, 0x41, 0x58, 0x49, 0x4d, 0x49, 0x5a, 0x45, 0x5f, 0x43, 0x4f, 0x4e,
	0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x53, 0x5f, 0x4f, 0x50, 0x54, 0x5f, 0x49, 0x4e, 0x10,
	0x06, 0x12, 0x17, 0x0a, 0x13, 0x45, 0x4e, 0x48, 0x41, 0x4e, 0x43, 0x45, 0x44, 0x5f, 0x43, 0x50,
	0x43, 0x5f, 0x4f, 0x50, 0x54, 0x5f, 0x49, 0x4e, 0x10, 0x07, 0x12, 0x1a, 0x0a, 0x16, 0x53, 0x45,
	0x41, 0x52, 0x43, 0x48, 0x5f, 0x50, 0x41, 0x52, 0x54, 0x4e, 0x45, 0x52, 0x53, 0x5f, 0x4f, 0x50,
	0x54, 0x5f, 0x49, 0x4e, 0x10, 0x08, 0x12, 0x1a, 0x0a, 0x16, 0x4d, 0x41, 0x58, 0x49, 0x4d, 0x49,
	0x5a, 0x45, 0x5f, 0x43, 0x4c, 0x49, 0x43, 0x4b, 0x53, 0x5f, 0x4f, 0x50, 0x54, 0x5f, 0x49, 0x4e,
	0x10, 0x09, 0x12, 0x18, 0x0a, 0x14, 0x4f, 0x50, 0x54, 0x49, 0x4d, 0x49, 0x5a, 0x45, 0x5f, 0x41,
	0x44, 0x5f, 0x52, 0x4f, 0x54, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x0a, 0x12, 0x16, 0x0a, 0x12,
	0x4b, 0x45, 0x59, 0x57, 0x4f, 0x52, 0x44, 0x5f, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x10, 0x0e, 0x12, 0x16, 0x0a, 0x12, 0x4d, 0x4f, 0x56, 0x45, 0x5f, 0x55, 0x4e, 0x55,
	0x53, 0x45, 0x44, 0x5f, 0x42, 0x55, 0x44, 0x47, 0x45, 0x54, 0x10, 0x0f, 0x12, 0x1f, 0x0a, 0x1b,
	0x46, 0x4f, 0x52, 0x45, 0x43, 0x41, 0x53, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x43, 0x41, 0x4d, 0x50,
	0x41, 0x49, 0x47, 0x4e, 0x5f, 0x42, 0x55, 0x44, 0x47, 0x45, 0x54, 0x10, 0x10, 0x12, 0x16, 0x0a,
	0x12, 0x54, 0x41, 0x52, 0x47, 0x45, 0x54, 0x5f, 0x52, 0x4f, 0x41, 0x53, 0x5f, 0x4f, 0x50, 0x54,
	0x5f, 0x49, 0x4e, 0x10, 0x11, 0x12, 0x18, 0x0a, 0x14, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53,
	0x49, 0x56, 0x45, 0x5f, 0x53, 0x45, 0x41, 0x52, 0x43, 0x48, 0x5f, 0x41, 0x44, 0x10, 0x12, 0x12,
	0x20, 0x0a, 0x1c, 0x4d, 0x41, 0x52, 0x47, 0x49, 0x4e, 0x41, 0x4c, 0x5f, 0x52, 0x4f, 0x49, 0x5f,
	0x43, 0x41, 0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e, 0x5f, 0x42, 0x55, 0x44, 0x47, 0x45, 0x54, 0x10,
	0x13, 0x12, 0x1b, 0x0a, 0x17, 0x55, 0x53, 0x45, 0x5f, 0x42, 0x52, 0x4f, 0x41, 0x44, 0x5f, 0x4d,
	0x41, 0x54, 0x43, 0x48, 0x5f, 0x4b, 0x45, 0x59, 0x57, 0x4f, 0x52, 0x44, 0x10, 0x14, 0x12, 0x1e,
	0x0a, 0x1a, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x49, 0x56, 0x45, 0x5f, 0x53, 0x45, 0x41,
	0x52, 0x43, 0x48, 0x5f, 0x41, 0x44, 0x5f, 0x41, 0x53, 0x53, 0x45, 0x54, 0x10, 0x15, 0x12, 0x36,
	0x0a, 0x32, 0x55, 0x50, 0x47, 0x52, 0x41, 0x44, 0x45, 0x5f, 0x53, 0x4d, 0x41, 0x52, 0x54, 0x5f,
	0x53, 0x48, 0x4f, 0x50, 0x50, 0x49, 0x4e, 0x47, 0x5f, 0x43, 0x41, 0x4d, 0x50, 0x41, 0x49, 0x47,
	0x4e, 0x5f, 0x54, 0x4f, 0x5f, 0x50, 0x45, 0x52, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x4e, 0x43, 0x45,
	0x5f, 0x4d, 0x41, 0x58, 0x10, 0x16, 0x12, 0x2c, 0x0a, 0x28, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e,
	0x53, 0x49, 0x56, 0x45, 0x5f, 0x53, 0x45, 0x41, 0x52, 0x43, 0x48, 0x5f, 0x41, 0x44, 0x5f, 0x49,
	0x4d, 0x50, 0x52, 0x4f, 0x56, 0x45, 0x5f, 0x41, 0x44, 0x5f, 0x53, 0x54, 0x52, 0x45, 0x4e, 0x47,
	0x54, 0x48, 0x10, 0x17, 0x12, 0x1c, 0x0a, 0x18, 0x44, 0x49, 0x53, 0x50, 0x4c, 0x41, 0x59, 0x5f,
	0x45, 0x58, 0x50, 0x41, 0x4e, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x4f, 0x50, 0x54, 0x5f, 0x49, 0x4e,
	0x10, 0x18, 0x12, 0x2d, 0x0a, 0x29, 0x55, 0x50, 0x47, 0x52, 0x41, 0x44, 0x45, 0x5f, 0x4c, 0x4f,
	0x43, 0x41, 0x4c, 0x5f, 0x43, 0x41, 0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e, 0x5f, 0x54, 0x4f, 0x5f,
	0x50, 0x45, 0x52, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x4e, 0x43, 0x45, 0x5f, 0x4d, 0x41, 0x58, 0x10,
	0x19, 0x12, 0x20, 0x0a, 0x1c, 0x52, 0x41, 0x49, 0x53, 0x45, 0x5f, 0x54, 0x41, 0x52, 0x47, 0x45,
	0x54, 0x5f, 0x43, 0x50, 0x41, 0x5f, 0x42, 0x49, 0x44, 0x5f, 0x54, 0x4f, 0x4f, 0x5f, 0x4c, 0x4f,
	0x57, 0x10, 0x1a, 0x12, 0x1f, 0x0a, 0x1b, 0x46, 0x4f, 0x52, 0x45, 0x43, 0x41, 0x53, 0x54, 0x49,
	0x4e, 0x47, 0x5f, 0x53, 0x45, 0x54, 0x5f, 0x54, 0x41, 0x52, 0x47, 0x45, 0x54, 0x5f, 0x52, 0x4f,
	0x41, 0x53, 0x10, 0x1b, 0x12, 0x11, 0x0a, 0x0d, 0x43, 0x41, 0x4c, 0x4c, 0x4f, 0x55, 0x54, 0x5f,
	0x41, 0x53, 0x53, 0x45, 0x54, 0x10, 0x1c, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x49, 0x54, 0x45, 0x4c,
	0x49, 0x4e, 0x4b, 0x5f, 0x41, 0x53, 0x53, 0x45, 0x54, 0x10, 0x1d, 0x12, 0x0e, 0x0a, 0x0a, 0x43,
	0x41, 0x4c, 0x4c, 0x5f, 0x41, 0x53, 0x53, 0x45, 0x54, 0x10, 0x1e, 0x12, 0x1a, 0x0a, 0x16, 0x53,
	0x48, 0x4f, 0x50, 0x50, 0x49, 0x4e, 0x47, 0x5f, 0x41, 0x44, 0x44, 0x5f, 0x41, 0x47, 0x45, 0x5f,
	0x47, 0x52, 0x4f, 0x55, 0x50, 0x10, 0x1f, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x48, 0x4f, 0x50, 0x50,
	0x49, 0x4e, 0x47, 0x5f, 0x41, 0x44, 0x44, 0x5f, 0x43, 0x4f, 0x4c, 0x4f, 0x52, 0x10, 0x20, 0x12,
	0x17, 0x0a, 0x13, 0x53, 0x48, 0x4f, 0x50, 0x50, 0x49, 0x4e, 0x47, 0x5f, 0x41, 0x44, 0x44, 0x5f,
	0x47, 0x45, 0x4e, 0x44, 0x45, 0x52, 0x10, 0x21, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x48, 0x4f, 0x50,
	0x50, 0x49, 0x4e, 0x47, 0x5f, 0x41, 0x44, 0x44, 0x5f, 0x47, 0x54, 0x49, 0x4e, 0x10, 0x22, 0x12,
	0x21, 0x0a, 0x1d, 0x53, 0x48, 0x4f, 0x50, 0x50, 0x49, 0x4e, 0x47, 0x5f, 0x41, 0x44, 0x44, 0x5f,
	0x4d, 0x4f, 0x52, 0x45, 0x5f, 0x49, 0x44, 0x45, 0x4e, 0x54, 0x49, 0x46, 0x49, 0x45, 0x52, 0x53,
	0x10, 0x23, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x48, 0x4f, 0x50, 0x50, 0x49, 0x4e, 0x47, 0x5f, 0x41,
	0x44, 0x44, 0x5f, 0x53, 0x49, 0x5a, 0x45, 0x10, 0x24, 0x12, 0x25, 0x0a, 0x21, 0x53, 0x48, 0x4f,
	0x50, 0x50, 0x49, 0x4e, 0x47, 0x5f, 0x41, 0x44, 0x44, 0x5f, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43,
	0x54, 0x53, 0x5f, 0x54, 0x4f, 0x5f, 0x43, 0x41, 0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e, 0x10, 0x25,
	0x12, 0x25, 0x0a, 0x21, 0x53, 0x48, 0x4f, 0x50, 0x50, 0x49, 0x4e, 0x47, 0x5f, 0x46, 0x49, 0x58,
	0x5f, 0x44, 0x49, 0x53, 0x41, 0x50, 0x50, 0x52, 0x4f, 0x56, 0x45, 0x44, 0x5f, 0x50, 0x52, 0x4f,
	0x44, 0x55, 0x43, 0x54, 0x53, 0x10, 0x26, 0x12, 0x1e, 0x0a, 0x1a, 0x53, 0x48, 0x4f, 0x50, 0x50,
	0x49, 0x4e, 0x47, 0x5f, 0x54, 0x41, 0x52, 0x47, 0x45, 0x54, 0x5f, 0x41, 0x4c, 0x4c, 0x5f, 0x4f,
	0x46, 0x46, 0x45, 0x52, 0x53, 0x10, 0x27, 0x12, 0x32, 0x0a, 0x2e, 0x53, 0x48, 0x4f, 0x50, 0x50,
	0x49, 0x4e, 0x47, 0x5f, 0x46, 0x49, 0x58, 0x5f, 0x53, 0x55, 0x53, 0x50, 0x45, 0x4e, 0x44, 0x45,
	0x44, 0x5f, 0x4d, 0x45, 0x52, 0x43, 0x48, 0x41, 0x4e, 0x54, 0x5f, 0x43, 0x45, 0x4e, 0x54, 0x45,
	0x52, 0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x10, 0x28, 0x12, 0x3b, 0x0a, 0x37, 0x53,
	0x48, 0x4f, 0x50, 0x50, 0x49, 0x4e, 0x47, 0x5f, 0x46, 0x49, 0x58, 0x5f, 0x4d, 0x45, 0x52, 0x43,
	0x48, 0x41, 0x4e, 0x54, 0x5f, 0x43, 0x45, 0x4e, 0x54, 0x45, 0x52, 0x5f, 0x41, 0x43, 0x43, 0x4f,
	0x55, 0x4e, 0x54, 0x5f, 0x53, 0x55, 0x53, 0x50, 0x45, 0x4e, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x57,
	0x41, 0x52, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x29, 0x12, 0x48, 0x0a, 0x44, 0x53, 0x48, 0x4f, 0x50,
	0x50, 0x49, 0x4e, 0x47, 0x5f, 0x4d, 0x49, 0x47, 0x52, 0x41, 0x54, 0x45, 0x5f, 0x52, 0x45, 0x47,
	0x55, 0x4c, 0x41, 0x52, 0x5f, 0x53, 0x48, 0x4f, 0x50, 0x50, 0x49, 0x4e, 0x47, 0x5f, 0x43, 0x41,
	0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e, 0x5f, 0x4f, 0x46, 0x46, 0x45, 0x52, 0x53, 0x5f, 0x54, 0x4f,
	0x5f, 0x50, 0x45, 0x52, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x4e, 0x43, 0x45, 0x5f, 0x4d, 0x41, 0x58,
	0x10, 0x2a, 0x12, 0x22, 0x0a, 0x1e, 0x44, 0x59, 0x4e, 0x41, 0x4d, 0x49, 0x43, 0x5f, 0x49, 0x4d,
	0x41, 0x47, 0x45, 0x5f, 0x45, 0x58, 0x54, 0x45, 0x4e, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x4f, 0x50,
	0x54, 0x5f, 0x49, 0x4e, 0x10, 0x2b, 0x12, 0x14, 0x0a, 0x10, 0x52, 0x41, 0x49, 0x53, 0x45, 0x5f,
	0x54, 0x41, 0x52, 0x47, 0x45, 0x54, 0x5f, 0x43, 0x50, 0x41, 0x10, 0x2c, 0x12, 0x15, 0x0a, 0x11,
	0x4c, 0x4f, 0x57, 0x45, 0x52, 0x5f, 0x54, 0x41, 0x52, 0x47, 0x45, 0x54, 0x5f, 0x52, 0x4f, 0x41,
	0x53, 0x10, 0x2d, 0x12, 0x1a, 0x0a, 0x16, 0x50, 0x45, 0x52, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x4e,
	0x43, 0x45, 0x5f, 0x4d, 0x41, 0x58, 0x5f, 0x4f, 0x50, 0x54, 0x5f, 0x49, 0x4e, 0x10, 0x2e, 0x12,
	0x27, 0x0a, 0x23, 0x49, 0x4d, 0x50, 0x52, 0x4f, 0x56, 0x45, 0x5f, 0x50, 0x45, 0x52, 0x46, 0x4f,
	0x52, 0x4d, 0x41, 0x4e, 0x43, 0x45, 0x5f, 0x4d, 0x41, 0x58, 0x5f, 0x41, 0x44, 0x5f, 0x53, 0x54,
	0x52, 0x45, 0x4e, 0x47, 0x54, 0x48, 0x10, 0x2f, 0x12, 0x3a, 0x0a, 0x36, 0x4d, 0x49, 0x47, 0x52,
	0x41, 0x54, 0x45, 0x5f, 0x44, 0x59, 0x4e, 0x41, 0x4d, 0x49, 0x43, 0x5f, 0x53, 0x45, 0x41, 0x52,
	0x43, 0x48, 0x5f, 0x41, 0x44, 0x53, 0x5f, 0x43, 0x41, 0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e, 0x5f,
	0x54, 0x4f, 0x5f, 0x50, 0x45, 0x52, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x4e, 0x43, 0x45, 0x5f, 0x4d,
	0x41, 0x58, 0x10, 0x30, 0x12, 0x1e, 0x0a, 0x1a, 0x46, 0x4f, 0x52, 0x45, 0x43, 0x41, 0x53, 0x54,
	0x49, 0x4e, 0x47, 0x5f, 0x53, 0x45, 0x54, 0x5f, 0x54, 0x41, 0x52, 0x47, 0x45, 0x54, 0x5f, 0x43,
	0x50, 0x41, 0x10, 0x31, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x45, 0x54, 0x5f, 0x54, 0x41, 0x52, 0x47,
	0x45, 0x54, 0x5f, 0x43, 0x50, 0x41, 0x10, 0x32, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x45, 0x54, 0x5f,
	0x54, 0x41, 0x52, 0x47, 0x45, 0x54, 0x5f, 0x52, 0x4f, 0x41, 0x53, 0x10, 0x33, 0x12, 0x24, 0x0a,
	0x20, 0x4d, 0x41, 0x58, 0x49, 0x4d, 0x49, 0x5a, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52,
	0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x4f, 0x50, 0x54, 0x5f, 0x49,
	0x4e, 0x10, 0x34, 0x12, 0x1f, 0x0a, 0x1b, 0x49, 0x4d, 0x50, 0x52, 0x4f, 0x56, 0x45, 0x5f, 0x47,
	0x4f, 0x4f, 0x47, 0x4c, 0x45, 0x5f, 0x54, 0x41, 0x47, 0x5f, 0x43, 0x4f, 0x56, 0x45, 0x52, 0x41,
	0x47, 0x45, 0x10, 0x35, 0x12, 0x24, 0x0a, 0x20, 0x50, 0x45, 0x52, 0x46, 0x4f, 0x52, 0x4d, 0x41,
	0x4e, 0x43, 0x45, 0x5f, 0x4d, 0x41, 0x58, 0x5f, 0x46, 0x49, 0x4e, 0x41, 0x4c, 0x5f, 0x55, 0x52,
	0x4c, 0x5f, 0x4f, 0x50, 0x54, 0x5f, 0x49, 0x4e, 0x10, 0x36, 0x12, 0x1f, 0x0a, 0x1b, 0x52, 0x45,
	0x46, 0x52, 0x45, 0x53, 0x48, 0x5f, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x45, 0x52, 0x5f, 0x4d,
	0x41, 0x54, 0x43, 0x48, 0x5f, 0x4c, 0x49, 0x53, 0x54, 0x10, 0x37, 0x12, 0x1a, 0x0a, 0x16, 0x43,
	0x55, 0x53, 0x54, 0x4f, 0x4d, 0x5f, 0x41, 0x55, 0x44, 0x49, 0x45, 0x4e, 0x43, 0x45, 0x5f, 0x4f,
	0x50, 0x54, 0x5f, 0x49, 0x4e, 0x10, 0x38, 0x12, 0x13, 0x0a, 0x0f, 0x4c, 0x45, 0x41, 0x44, 0x5f,
	0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x41, 0x53, 0x53, 0x45, 0x54, 0x10, 0x39, 0x12, 0x22, 0x0a, 0x1e,
	0x49, 0x4d, 0x50, 0x52, 0x4f, 0x56, 0x45, 0x5f, 0x44, 0x45, 0x4d, 0x41, 0x4e, 0x44, 0x5f, 0x47,
	0x45, 0x4e, 0x5f, 0x41, 0x44, 0x5f, 0x53, 0x54, 0x52, 0x45, 0x4e, 0x47, 0x54, 0x48, 0x10, 0x3a,
	0x42, 0xf1, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31,
	0x38, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x42, 0x17, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e,
	0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x38, 0x2f, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x3b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1e,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x38, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xca, 0x02,
	0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x38, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xea,
	0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x38, 0x3a, 0x3a, 0x45,
	0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v18_enums_recommendation_type_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v18_enums_recommendation_type_proto_rawDescData = file_google_ads_googleads_v18_enums_recommendation_type_proto_rawDesc
)

func file_google_ads_googleads_v18_enums_recommendation_type_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v18_enums_recommendation_type_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v18_enums_recommendation_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v18_enums_recommendation_type_proto_rawDescData)
	})
	return file_google_ads_googleads_v18_enums_recommendation_type_proto_rawDescData
}

var file_google_ads_googleads_v18_enums_recommendation_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v18_enums_recommendation_type_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v18_enums_recommendation_type_proto_goTypes = []any{
	(RecommendationTypeEnum_RecommendationType)(0), // 0: google.ads.googleads.v18.enums.RecommendationTypeEnum.RecommendationType
	(*RecommendationTypeEnum)(nil),                 // 1: google.ads.googleads.v18.enums.RecommendationTypeEnum
}
var file_google_ads_googleads_v18_enums_recommendation_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v18_enums_recommendation_type_proto_init() }
func file_google_ads_googleads_v18_enums_recommendation_type_proto_init() {
	if File_google_ads_googleads_v18_enums_recommendation_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v18_enums_recommendation_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v18_enums_recommendation_type_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v18_enums_recommendation_type_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v18_enums_recommendation_type_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v18_enums_recommendation_type_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v18_enums_recommendation_type_proto = out.File
	file_google_ads_googleads_v18_enums_recommendation_type_proto_rawDesc = nil
	file_google_ads_googleads_v18_enums_recommendation_type_proto_goTypes = nil
	file_google_ads_googleads_v18_enums_recommendation_type_proto_depIdxs = nil
}
