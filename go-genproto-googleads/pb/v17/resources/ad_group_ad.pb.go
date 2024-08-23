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
// source: google/ads/googleads/v17/resources/ad_group_ad.proto

package resources

import (
	common "github.com/TeamMomentum/go-pkg/go-genproto-googleads/pb/v17/common"
	enums "github.com/TeamMomentum/go-pkg/go-genproto-googleads/pb/v17/enums"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// An ad group ad.
type AdGroupAd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Immutable. The resource name of the ad.
	// Ad group ad resource names have the form:
	//
	// `customers/{customer_id}/adGroupAds/{ad_group_id}~{ad_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The status of the ad.
	Status enums.AdGroupAdStatusEnum_AdGroupAdStatus `protobuf:"varint,3,opt,name=status,proto3,enum=google.ads.googleads.v17.enums.AdGroupAdStatusEnum_AdGroupAdStatus" json:"status,omitempty"`
	// Immutable. The ad group to which the ad belongs.
	AdGroup *string `protobuf:"bytes,9,opt,name=ad_group,json=adGroup,proto3,oneof" json:"ad_group,omitempty"`
	// Immutable. The ad.
	Ad *Ad `protobuf:"bytes,5,opt,name=ad,proto3" json:"ad,omitempty"`
	// Output only. Policy information for the ad.
	PolicySummary *AdGroupAdPolicySummary `protobuf:"bytes,6,opt,name=policy_summary,json=policySummary,proto3" json:"policy_summary,omitempty"`
	// Output only. Overall ad strength for this ad group ad.
	AdStrength enums.AdStrengthEnum_AdStrength `protobuf:"varint,7,opt,name=ad_strength,json=adStrength,proto3,enum=google.ads.googleads.v17.enums.AdStrengthEnum_AdStrength" json:"ad_strength,omitempty"`
	// Output only. A list of recommendations to improve the ad strength. For
	// example, a recommendation could be "Try adding a few more unique headlines
	// or unpinning some assets.".
	ActionItems []string `protobuf:"bytes,13,rep,name=action_items,json=actionItems,proto3" json:"action_items,omitempty"`
	// Output only. The resource names of labels attached to this ad group ad.
	Labels []string `protobuf:"bytes,10,rep,name=labels,proto3" json:"labels,omitempty"`
	// Output only. Provides aggregated view into why an ad group ad is not
	// serving or not serving optimally.
	PrimaryStatus enums.AdGroupAdPrimaryStatusEnum_AdGroupAdPrimaryStatus `protobuf:"varint,16,opt,name=primary_status,json=primaryStatus,proto3,enum=google.ads.googleads.v17.enums.AdGroupAdPrimaryStatusEnum_AdGroupAdPrimaryStatus" json:"primary_status,omitempty"`
	// Output only. Provides reasons for why an ad group ad is not serving or not
	// serving optimally.
	PrimaryStatusReasons []enums.AdGroupAdPrimaryStatusReasonEnum_AdGroupAdPrimaryStatusReason `protobuf:"varint,17,rep,packed,name=primary_status_reasons,json=primaryStatusReasons,proto3,enum=google.ads.googleads.v17.enums.AdGroupAdPrimaryStatusReasonEnum_AdGroupAdPrimaryStatusReason" json:"primary_status_reasons,omitempty"`
}

func (x *AdGroupAd) Reset() {
	*x = AdGroupAd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v17_resources_ad_group_ad_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdGroupAd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdGroupAd) ProtoMessage() {}

func (x *AdGroupAd) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v17_resources_ad_group_ad_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdGroupAd.ProtoReflect.Descriptor instead.
func (*AdGroupAd) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v17_resources_ad_group_ad_proto_rawDescGZIP(), []int{0}
}

func (x *AdGroupAd) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *AdGroupAd) GetStatus() enums.AdGroupAdStatusEnum_AdGroupAdStatus {
	if x != nil {
		return x.Status
	}
	return enums.AdGroupAdStatusEnum_AdGroupAdStatus(0)
}

func (x *AdGroupAd) GetAdGroup() string {
	if x != nil && x.AdGroup != nil {
		return *x.AdGroup
	}
	return ""
}

func (x *AdGroupAd) GetAd() *Ad {
	if x != nil {
		return x.Ad
	}
	return nil
}

func (x *AdGroupAd) GetPolicySummary() *AdGroupAdPolicySummary {
	if x != nil {
		return x.PolicySummary
	}
	return nil
}

func (x *AdGroupAd) GetAdStrength() enums.AdStrengthEnum_AdStrength {
	if x != nil {
		return x.AdStrength
	}
	return enums.AdStrengthEnum_AdStrength(0)
}

func (x *AdGroupAd) GetActionItems() []string {
	if x != nil {
		return x.ActionItems
	}
	return nil
}

func (x *AdGroupAd) GetLabels() []string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *AdGroupAd) GetPrimaryStatus() enums.AdGroupAdPrimaryStatusEnum_AdGroupAdPrimaryStatus {
	if x != nil {
		return x.PrimaryStatus
	}
	return enums.AdGroupAdPrimaryStatusEnum_AdGroupAdPrimaryStatus(0)
}

func (x *AdGroupAd) GetPrimaryStatusReasons() []enums.AdGroupAdPrimaryStatusReasonEnum_AdGroupAdPrimaryStatusReason {
	if x != nil {
		return x.PrimaryStatusReasons
	}
	return nil
}

// Contains policy information for an ad.
type AdGroupAdPolicySummary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The list of policy findings for this ad.
	PolicyTopicEntries []*common.PolicyTopicEntry `protobuf:"bytes,1,rep,name=policy_topic_entries,json=policyTopicEntries,proto3" json:"policy_topic_entries,omitempty"`
	// Output only. Where in the review process this ad is.
	ReviewStatus enums.PolicyReviewStatusEnum_PolicyReviewStatus `protobuf:"varint,2,opt,name=review_status,json=reviewStatus,proto3,enum=google.ads.googleads.v17.enums.PolicyReviewStatusEnum_PolicyReviewStatus" json:"review_status,omitempty"`
	// Output only. The overall approval status of this ad, calculated based on
	// the status of its individual policy topic entries.
	ApprovalStatus enums.PolicyApprovalStatusEnum_PolicyApprovalStatus `protobuf:"varint,3,opt,name=approval_status,json=approvalStatus,proto3,enum=google.ads.googleads.v17.enums.PolicyApprovalStatusEnum_PolicyApprovalStatus" json:"approval_status,omitempty"`
}

func (x *AdGroupAdPolicySummary) Reset() {
	*x = AdGroupAdPolicySummary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v17_resources_ad_group_ad_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdGroupAdPolicySummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdGroupAdPolicySummary) ProtoMessage() {}

func (x *AdGroupAdPolicySummary) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v17_resources_ad_group_ad_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdGroupAdPolicySummary.ProtoReflect.Descriptor instead.
func (*AdGroupAdPolicySummary) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v17_resources_ad_group_ad_proto_rawDescGZIP(), []int{1}
}

func (x *AdGroupAdPolicySummary) GetPolicyTopicEntries() []*common.PolicyTopicEntry {
	if x != nil {
		return x.PolicyTopicEntries
	}
	return nil
}

func (x *AdGroupAdPolicySummary) GetReviewStatus() enums.PolicyReviewStatusEnum_PolicyReviewStatus {
	if x != nil {
		return x.ReviewStatus
	}
	return enums.PolicyReviewStatusEnum_PolicyReviewStatus(0)
}

func (x *AdGroupAdPolicySummary) GetApprovalStatus() enums.PolicyApprovalStatusEnum_PolicyApprovalStatus {
	if x != nil {
		return x.ApprovalStatus
	}
	return enums.PolicyApprovalStatusEnum_PolicyApprovalStatus(0)
}

var File_google_ads_googleads_v17_resources_ad_group_ad_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v17_resources_ad_group_ad_proto_rawDesc = []byte{
	0x0a, 0x34, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x37, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x61, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x61, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x37,
	0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0x2c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2f, 0x76, 0x31, 0x37, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76,
	0x31, 0x37, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x61, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x5f, 0x61, 0x64, 0x5f, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x46, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f,
	0x76, 0x31, 0x37, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x61, 0x64, 0x5f, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x5f, 0x61, 0x64, 0x5f, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x37, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x37, 0x2f, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x2f, 0x61, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x61, 0x64, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x30, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2f, 0x76, 0x31, 0x37, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x61, 0x64, 0x5f, 0x73, 0x74,
	0x72, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2f, 0x76, 0x31, 0x37, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x5f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x39, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f,
	0x76, 0x31, 0x37, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x5f, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x37, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x84, 0x08,
	0x0a, 0x09, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x64, 0x12, 0x4f, 0x0a, 0x0d, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x2a, 0xe0, 0x41, 0x05, 0xfa, 0x41, 0x24, 0x0a, 0x22, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x64, 0x52, 0x0c,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x5b, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x43, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x41, 0x64,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x75,
	0x6d, 0x2e, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x48, 0x0a, 0x08, 0x61, 0x64, 0x5f,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x42, 0x28, 0xe0, 0x41, 0x05,
	0xfa, 0x41, 0x22, 0x0a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x64,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x48, 0x00, 0x52, 0x07, 0x61, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x88, 0x01, 0x01, 0x12, 0x3b, 0x0a, 0x02, 0x61, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x26, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x64, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x52, 0x02, 0x61, 0x64,
	0x12, 0x66, 0x0a, 0x0e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x73, 0x75, 0x6d, 0x6d, 0x61,
	0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x37, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x64,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x64, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x75, 0x6d,
	0x6d, 0x61, 0x72, 0x79, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0d, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x5f, 0x0a, 0x0b, 0x61, 0x64, 0x5f, 0x73,
	0x74, 0x72, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x39, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x41,
	0x64, 0x53, 0x74, 0x72, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x41, 0x64,
	0x53, 0x74, 0x72, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x61,
	0x64, 0x53, 0x74, 0x72, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x26, 0x0a, 0x0c, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x09, 0x42,
	0x03, 0xe0, 0x41, 0x03, 0x52, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x74, 0x65, 0x6d,
	0x73, 0x12, 0x47, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28,
	0x09, 0x42, 0x2f, 0xe0, 0x41, 0x03, 0xfa, 0x41, 0x29, 0x0a, 0x27, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x64, 0x4c, 0x61, 0x62,
	0x65, 0x6c, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x7d, 0x0a, 0x0e, 0x70, 0x72,
	0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x10, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x51, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x2e, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x64, 0x50, 0x72, 0x69,
	0x6d, 0x61, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x41,
	0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x64, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0d, 0x70, 0x72, 0x69, 0x6d,
	0x61, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x98, 0x01, 0x0a, 0x16, 0x70, 0x72,
	0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x72, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x73, 0x18, 0x11, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x5d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x41, 0x64, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x41, 0x64, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x41, 0x64, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x41, 0x64, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x14,
	0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x73, 0x3a, 0x61, 0xea, 0x41, 0x5e, 0x0a, 0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x64, 0x12, 0x38, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x61, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x64,
	0x73, 0x2f, 0x7b, 0x61, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x7d, 0x7e,
	0x7b, 0x61, 0x64, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x61, 0x64, 0x5f, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x22, 0xf4, 0x02, 0x0a, 0x16, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x41, 0x64, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12,
	0x68, 0x0a, 0x14, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x5f,
	0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x12, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x54, 0x6f, 0x70,
	0x69, 0x63, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x73, 0x0a, 0x0d, 0x72, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x49, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x03, 0xe0, 0x41, 0x03,
	0x52, 0x0c, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x7b,
	0x0a, 0x0f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x4d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x37, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x41,
	0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x75,
	0x6d, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0e, 0x61, 0x70, 0x70,
	0x72, 0x6f, 0x76, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x80, 0x02, 0x0a, 0x26,
	0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0x0e, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41,
	0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31,
	0x37, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x3b, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x22, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41,
	0x64, 0x73, 0x2e, 0x56, 0x31, 0x37, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0xca, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x37, 0x5c, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0xea, 0x02, 0x26, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a,
	0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a,
	0x56, 0x31, 0x37, 0x3a, 0x3a, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v17_resources_ad_group_ad_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v17_resources_ad_group_ad_proto_rawDescData = file_google_ads_googleads_v17_resources_ad_group_ad_proto_rawDesc
)

func file_google_ads_googleads_v17_resources_ad_group_ad_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v17_resources_ad_group_ad_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v17_resources_ad_group_ad_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v17_resources_ad_group_ad_proto_rawDescData)
	})
	return file_google_ads_googleads_v17_resources_ad_group_ad_proto_rawDescData
}

var file_google_ads_googleads_v17_resources_ad_group_ad_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_ads_googleads_v17_resources_ad_group_ad_proto_goTypes = []any{
	(*AdGroupAd)(nil),                              // 0: google.ads.googleads.v17.resources.AdGroupAd
	(*AdGroupAdPolicySummary)(nil),                 // 1: google.ads.googleads.v17.resources.AdGroupAdPolicySummary
	(enums.AdGroupAdStatusEnum_AdGroupAdStatus)(0), // 2: google.ads.googleads.v17.enums.AdGroupAdStatusEnum.AdGroupAdStatus
	(*Ad)(nil),                           // 3: google.ads.googleads.v17.resources.Ad
	(enums.AdStrengthEnum_AdStrength)(0), // 4: google.ads.googleads.v17.enums.AdStrengthEnum.AdStrength
	(enums.AdGroupAdPrimaryStatusEnum_AdGroupAdPrimaryStatus)(0),             // 5: google.ads.googleads.v17.enums.AdGroupAdPrimaryStatusEnum.AdGroupAdPrimaryStatus
	(enums.AdGroupAdPrimaryStatusReasonEnum_AdGroupAdPrimaryStatusReason)(0), // 6: google.ads.googleads.v17.enums.AdGroupAdPrimaryStatusReasonEnum.AdGroupAdPrimaryStatusReason
	(*common.PolicyTopicEntry)(nil),                                          // 7: google.ads.googleads.v17.common.PolicyTopicEntry
	(enums.PolicyReviewStatusEnum_PolicyReviewStatus)(0),                     // 8: google.ads.googleads.v17.enums.PolicyReviewStatusEnum.PolicyReviewStatus
	(enums.PolicyApprovalStatusEnum_PolicyApprovalStatus)(0),                 // 9: google.ads.googleads.v17.enums.PolicyApprovalStatusEnum.PolicyApprovalStatus
}
var file_google_ads_googleads_v17_resources_ad_group_ad_proto_depIdxs = []int32{
	2, // 0: google.ads.googleads.v17.resources.AdGroupAd.status:type_name -> google.ads.googleads.v17.enums.AdGroupAdStatusEnum.AdGroupAdStatus
	3, // 1: google.ads.googleads.v17.resources.AdGroupAd.ad:type_name -> google.ads.googleads.v17.resources.Ad
	1, // 2: google.ads.googleads.v17.resources.AdGroupAd.policy_summary:type_name -> google.ads.googleads.v17.resources.AdGroupAdPolicySummary
	4, // 3: google.ads.googleads.v17.resources.AdGroupAd.ad_strength:type_name -> google.ads.googleads.v17.enums.AdStrengthEnum.AdStrength
	5, // 4: google.ads.googleads.v17.resources.AdGroupAd.primary_status:type_name -> google.ads.googleads.v17.enums.AdGroupAdPrimaryStatusEnum.AdGroupAdPrimaryStatus
	6, // 5: google.ads.googleads.v17.resources.AdGroupAd.primary_status_reasons:type_name -> google.ads.googleads.v17.enums.AdGroupAdPrimaryStatusReasonEnum.AdGroupAdPrimaryStatusReason
	7, // 6: google.ads.googleads.v17.resources.AdGroupAdPolicySummary.policy_topic_entries:type_name -> google.ads.googleads.v17.common.PolicyTopicEntry
	8, // 7: google.ads.googleads.v17.resources.AdGroupAdPolicySummary.review_status:type_name -> google.ads.googleads.v17.enums.PolicyReviewStatusEnum.PolicyReviewStatus
	9, // 8: google.ads.googleads.v17.resources.AdGroupAdPolicySummary.approval_status:type_name -> google.ads.googleads.v17.enums.PolicyApprovalStatusEnum.PolicyApprovalStatus
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v17_resources_ad_group_ad_proto_init() }
func file_google_ads_googleads_v17_resources_ad_group_ad_proto_init() {
	if File_google_ads_googleads_v17_resources_ad_group_ad_proto != nil {
		return
	}
	file_google_ads_googleads_v17_resources_ad_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v17_resources_ad_group_ad_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*AdGroupAd); i {
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
		file_google_ads_googleads_v17_resources_ad_group_ad_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*AdGroupAdPolicySummary); i {
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
	file_google_ads_googleads_v17_resources_ad_group_ad_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v17_resources_ad_group_ad_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v17_resources_ad_group_ad_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v17_resources_ad_group_ad_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v17_resources_ad_group_ad_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v17_resources_ad_group_ad_proto = out.File
	file_google_ads_googleads_v17_resources_ad_group_ad_proto_rawDesc = nil
	file_google_ads_googleads_v17_resources_ad_group_ad_proto_goTypes = nil
	file_google_ads_googleads_v17_resources_ad_group_ad_proto_depIdxs = nil
}
