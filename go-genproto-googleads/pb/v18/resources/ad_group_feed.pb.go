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
// source: google/ads/googleads/v18/resources/ad_group_feed.proto

package resources

import (
	common "github.com/TeamMomentum/go-pkg/go-genproto-googleads/pb/v18/common"
	enums "github.com/TeamMomentum/go-pkg/go-genproto-googleads/pb/v18/enums"
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

// An ad group feed.
type AdGroupFeed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Immutable. The resource name of the ad group feed.
	// Ad group feed resource names have the form:
	//
	// `customers/{customer_id}/adGroupFeeds/{ad_group_id}~{feed_id}
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Immutable. The feed being linked to the ad group.
	Feed *string `protobuf:"bytes,7,opt,name=feed,proto3,oneof" json:"feed,omitempty"`
	// Immutable. The ad group being linked to the feed.
	AdGroup *string `protobuf:"bytes,8,opt,name=ad_group,json=adGroup,proto3,oneof" json:"ad_group,omitempty"`
	// Indicates which placeholder types the feed may populate under the connected
	// ad group. Required.
	PlaceholderTypes []enums.PlaceholderTypeEnum_PlaceholderType `protobuf:"varint,4,rep,packed,name=placeholder_types,json=placeholderTypes,proto3,enum=google.ads.googleads.v18.enums.PlaceholderTypeEnum_PlaceholderType" json:"placeholder_types,omitempty"`
	// Matching function associated with the AdGroupFeed.
	// The matching function is used to filter the set of feed items selected.
	// Required.
	MatchingFunction *common.MatchingFunction `protobuf:"bytes,5,opt,name=matching_function,json=matchingFunction,proto3" json:"matching_function,omitempty"`
	// Output only. Status of the ad group feed.
	// This field is read-only.
	Status enums.FeedLinkStatusEnum_FeedLinkStatus `protobuf:"varint,6,opt,name=status,proto3,enum=google.ads.googleads.v18.enums.FeedLinkStatusEnum_FeedLinkStatus" json:"status,omitempty"`
}

func (x *AdGroupFeed) Reset() {
	*x = AdGroupFeed{}
	mi := &file_google_ads_googleads_v18_resources_ad_group_feed_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AdGroupFeed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdGroupFeed) ProtoMessage() {}

func (x *AdGroupFeed) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v18_resources_ad_group_feed_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdGroupFeed.ProtoReflect.Descriptor instead.
func (*AdGroupFeed) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v18_resources_ad_group_feed_proto_rawDescGZIP(), []int{0}
}

func (x *AdGroupFeed) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *AdGroupFeed) GetFeed() string {
	if x != nil && x.Feed != nil {
		return *x.Feed
	}
	return ""
}

func (x *AdGroupFeed) GetAdGroup() string {
	if x != nil && x.AdGroup != nil {
		return *x.AdGroup
	}
	return ""
}

func (x *AdGroupFeed) GetPlaceholderTypes() []enums.PlaceholderTypeEnum_PlaceholderType {
	if x != nil {
		return x.PlaceholderTypes
	}
	return nil
}

func (x *AdGroupFeed) GetMatchingFunction() *common.MatchingFunction {
	if x != nil {
		return x.MatchingFunction
	}
	return nil
}

func (x *AdGroupFeed) GetStatus() enums.FeedLinkStatusEnum_FeedLinkStatus {
	if x != nil {
		return x.Status
	}
	return enums.FeedLinkStatusEnum_FeedLinkStatus(0)
}

var File_google_ads_googleads_v18_resources_ad_group_feed_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v18_resources_ad_group_feed_proto_rawDesc = []byte{
	0x0a, 0x36, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x38, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x61, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x66, 0x65,
	0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x38, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0x37, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2f, 0x76, 0x31, 0x38, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x61,
	0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x35, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64,
	0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x38, 0x2f,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x66, 0x65, 0x65, 0x64, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x35, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2f, 0x76, 0x31, 0x38, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x70, 0x6c, 0x61,
	0x63, 0x65, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x9b, 0x05, 0x0a, 0x0b, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x46, 0x65, 0x65, 0x64, 0x12,
	0x51, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2c, 0xe0, 0x41, 0x05, 0xfa, 0x41, 0x26, 0x0a, 0x24,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x46, 0x65, 0x65, 0x64, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x3e, 0x0a, 0x04, 0x66, 0x65, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x25, 0xe0, 0x41, 0x05, 0xfa, 0x41, 0x1f, 0x0a, 0x1d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x46, 0x65, 0x65, 0x64, 0x48, 0x00, 0x52, 0x04, 0x66, 0x65, 0x65, 0x64, 0x88,
	0x01, 0x01, 0x12, 0x48, 0x0a, 0x08, 0x61, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x28, 0xe0, 0x41, 0x05, 0xfa, 0x41, 0x22, 0x0a, 0x20, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x48, 0x01,
	0x52, 0x07, 0x61, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x88, 0x01, 0x01, 0x12, 0x70, 0x0a, 0x11,
	0x70, 0x6c, 0x61, 0x63, 0x65, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x43, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x38, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x68, 0x6f,
	0x6c, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x50, 0x6c, 0x61,
	0x63, 0x65, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x10, 0x70, 0x6c,
	0x61, 0x63, 0x65, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x5e,
	0x0a, 0x11, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x75, 0x6e, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x31, 0x38, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x69, 0x6e, 0x67, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x10, 0x6d, 0x61,
	0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x5e,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x41,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x38, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e,
	0x46, 0x65, 0x65, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e,
	0x75, 0x6d, 0x2e, 0x46, 0x65, 0x65, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x3a, 0x67,
	0xea, 0x41, 0x64, 0x0a, 0x24, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x64,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x46, 0x65, 0x65, 0x64, 0x12, 0x3c, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x7d, 0x2f, 0x61, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x46, 0x65, 0x65, 0x64, 0x73, 0x2f,
	0x7b, 0x61, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x7d, 0x7e, 0x7b, 0x66,
	0x65, 0x65, 0x64, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x66, 0x65, 0x65, 0x64,
	0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x61, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x82, 0x02,
	0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x38, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0x10, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x46, 0x65, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2f, 0x76, 0x31, 0x38, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x3b,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa,
	0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x38, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0xca, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64,
	0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x38, 0x5c,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xea, 0x02, 0x26, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41,
	0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x38, 0x3a, 0x3a, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v18_resources_ad_group_feed_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v18_resources_ad_group_feed_proto_rawDescData = file_google_ads_googleads_v18_resources_ad_group_feed_proto_rawDesc
)

func file_google_ads_googleads_v18_resources_ad_group_feed_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v18_resources_ad_group_feed_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v18_resources_ad_group_feed_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v18_resources_ad_group_feed_proto_rawDescData)
	})
	return file_google_ads_googleads_v18_resources_ad_group_feed_proto_rawDescData
}

var file_google_ads_googleads_v18_resources_ad_group_feed_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v18_resources_ad_group_feed_proto_goTypes = []any{
	(*AdGroupFeed)(nil),                            // 0: google.ads.googleads.v18.resources.AdGroupFeed
	(enums.PlaceholderTypeEnum_PlaceholderType)(0), // 1: google.ads.googleads.v18.enums.PlaceholderTypeEnum.PlaceholderType
	(*common.MatchingFunction)(nil),                // 2: google.ads.googleads.v18.common.MatchingFunction
	(enums.FeedLinkStatusEnum_FeedLinkStatus)(0),   // 3: google.ads.googleads.v18.enums.FeedLinkStatusEnum.FeedLinkStatus
}
var file_google_ads_googleads_v18_resources_ad_group_feed_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v18.resources.AdGroupFeed.placeholder_types:type_name -> google.ads.googleads.v18.enums.PlaceholderTypeEnum.PlaceholderType
	2, // 1: google.ads.googleads.v18.resources.AdGroupFeed.matching_function:type_name -> google.ads.googleads.v18.common.MatchingFunction
	3, // 2: google.ads.googleads.v18.resources.AdGroupFeed.status:type_name -> google.ads.googleads.v18.enums.FeedLinkStatusEnum.FeedLinkStatus
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v18_resources_ad_group_feed_proto_init() }
func file_google_ads_googleads_v18_resources_ad_group_feed_proto_init() {
	if File_google_ads_googleads_v18_resources_ad_group_feed_proto != nil {
		return
	}
	file_google_ads_googleads_v18_resources_ad_group_feed_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v18_resources_ad_group_feed_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v18_resources_ad_group_feed_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v18_resources_ad_group_feed_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v18_resources_ad_group_feed_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v18_resources_ad_group_feed_proto = out.File
	file_google_ads_googleads_v18_resources_ad_group_feed_proto_rawDesc = nil
	file_google_ads_googleads_v18_resources_ad_group_feed_proto_goTypes = nil
	file_google_ads_googleads_v18_resources_ad_group_feed_proto_depIdxs = nil
}
