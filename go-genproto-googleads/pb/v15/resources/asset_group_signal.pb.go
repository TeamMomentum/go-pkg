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
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.2
// source: google/ads/googleads/v15/resources/asset_group_signal.proto

package resources

import (
	common "github.com/dictav/go-genproto-googleads/pb/v15/common"
	enums "github.com/dictav/go-genproto-googleads/pb/v15/enums"
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

// AssetGroupSignal represents a signal in an asset group. The existence of a
// signal tells the performance max campaign who's most likely to convert.
// Performance Max uses the signal to look for new people with similar or
// stronger intent to find conversions across Search, Display, Video, and more.
type AssetGroupSignal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Immutable. The resource name of the asset group signal.
	// Asset group signal resource name have the form:
	//
	// `customers/{customer_id}/assetGroupSignals/{asset_group_id}~{signal_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Immutable. The asset group which this asset group signal belongs to.
	AssetGroup string `protobuf:"bytes,2,opt,name=asset_group,json=assetGroup,proto3" json:"asset_group,omitempty"`
	// Output only. Approval status is the output value for search theme signal
	// after Google ads policy review. When using Audience signal, this field is
	// not used and will be absent.
	ApprovalStatus enums.AssetGroupSignalApprovalStatusEnum_AssetGroupSignalApprovalStatus `protobuf:"varint,6,opt,name=approval_status,json=approvalStatus,proto3,enum=google.ads.googleads.v15.enums.AssetGroupSignalApprovalStatusEnum_AssetGroupSignalApprovalStatus" json:"approval_status,omitempty"`
	// Output only. Computed for SearchTheme signals.
	// When using Audience signal, this field is not used and will be absent.
	DisapprovalReasons []string `protobuf:"bytes,7,rep,name=disapproval_reasons,json=disapprovalReasons,proto3" json:"disapproval_reasons,omitempty"`
	// The signal of the asset group.
	//
	// Types that are assignable to Signal:
	//	*AssetGroupSignal_Audience
	//	*AssetGroupSignal_SearchTheme
	Signal isAssetGroupSignal_Signal `protobuf_oneof:"signal"`
}

func (x *AssetGroupSignal) Reset() {
	*x = AssetGroupSignal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v15_resources_asset_group_signal_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssetGroupSignal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssetGroupSignal) ProtoMessage() {}

func (x *AssetGroupSignal) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v15_resources_asset_group_signal_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssetGroupSignal.ProtoReflect.Descriptor instead.
func (*AssetGroupSignal) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v15_resources_asset_group_signal_proto_rawDescGZIP(), []int{0}
}

func (x *AssetGroupSignal) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *AssetGroupSignal) GetAssetGroup() string {
	if x != nil {
		return x.AssetGroup
	}
	return ""
}

func (x *AssetGroupSignal) GetApprovalStatus() enums.AssetGroupSignalApprovalStatusEnum_AssetGroupSignalApprovalStatus {
	if x != nil {
		return x.ApprovalStatus
	}
	return enums.AssetGroupSignalApprovalStatusEnum_AssetGroupSignalApprovalStatus(0)
}

func (x *AssetGroupSignal) GetDisapprovalReasons() []string {
	if x != nil {
		return x.DisapprovalReasons
	}
	return nil
}

func (m *AssetGroupSignal) GetSignal() isAssetGroupSignal_Signal {
	if m != nil {
		return m.Signal
	}
	return nil
}

func (x *AssetGroupSignal) GetAudience() *common.AudienceInfo {
	if x, ok := x.GetSignal().(*AssetGroupSignal_Audience); ok {
		return x.Audience
	}
	return nil
}

func (x *AssetGroupSignal) GetSearchTheme() *common.SearchThemeInfo {
	if x, ok := x.GetSignal().(*AssetGroupSignal_SearchTheme); ok {
		return x.SearchTheme
	}
	return nil
}

type isAssetGroupSignal_Signal interface {
	isAssetGroupSignal_Signal()
}

type AssetGroupSignal_Audience struct {
	// Immutable. The audience signal to be used by the performance max
	// campaign.
	Audience *common.AudienceInfo `protobuf:"bytes,4,opt,name=audience,proto3,oneof"`
}

type AssetGroupSignal_SearchTheme struct {
	// Immutable. The search_theme signal to be used by the performance max
	// campaign.
	// Mutate errors of search_theme criterion includes
	// AssetGroupSignalError.UNSPECIFIED
	// AssetGroupSignalError.UNKNOWN
	// AssetGroupSignalError.TOO_MANY_WORDS
	// AssetGroupSignalError.SEARCH_THEME_POLICY_VIOLATION
	// FieldError.REQUIRED
	// StringFormatError.ILLEGAL_CHARS
	// StringLengthError.TOO_LONG
	// ResourceCountLimitExceededError.RESOURCE_LIMIT
	SearchTheme *common.SearchThemeInfo `protobuf:"bytes,5,opt,name=search_theme,json=searchTheme,proto3,oneof"`
}

func (*AssetGroupSignal_Audience) isAssetGroupSignal_Signal() {}

func (*AssetGroupSignal_SearchTheme) isAssetGroupSignal_Signal() {}

var File_google_ads_googleads_v15_resources_asset_group_signal_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v15_resources_asset_group_signal_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x35, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x5f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x35, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x63, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x47, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x35, 0x2f, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x6c, 0x5f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68,
	0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb3, 0x05, 0x0a, 0x10, 0x41, 0x73, 0x73, 0x65, 0x74,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x12, 0x56, 0x0a, 0x0d, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x31, 0xe0, 0x41, 0x05, 0xfa, 0x41, 0x2b, 0x0a, 0x29, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x73, 0x73, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53,
	0x69, 0x67, 0x6e, 0x61, 0x6c, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x4c, 0x0a, 0x0b, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2b, 0xe0, 0x41, 0x05, 0xfa, 0x41, 0x25,
	0x0a, 0x23, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x73, 0x73, 0x65, 0x74,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x0a, 0x61, 0x73, 0x73, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x12, 0x8f, 0x01, 0x0a, 0x0f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x61, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x41, 0x73, 0x73,
	0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x41, 0x70, 0x70,
	0x72, 0x6f, 0x76, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x2e,
	0x41, 0x73, 0x73, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c,
	0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x03,
	0xe0, 0x41, 0x03, 0x52, 0x0e, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x34, 0x0a, 0x13, 0x64, 0x69, 0x73, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76,
	0x61, 0x6c, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09,
	0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x12, 0x64, 0x69, 0x73, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76,
	0x61, 0x6c, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x73, 0x12, 0x50, 0x0a, 0x08, 0x61, 0x75, 0x64,
	0x69, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x75,
	0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x48,
	0x00, 0x52, 0x08, 0x61, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x5a, 0x0a, 0x0c, 0x73,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x5f, 0x74, 0x68, 0x65, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x30, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x68, 0x65, 0x6d, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x48, 0x00, 0x52, 0x0b, 0x73, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x54, 0x68, 0x65, 0x6d, 0x65, 0x3a, 0x79, 0xea, 0x41, 0x76, 0x0a, 0x29, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x73, 0x73, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x12, 0x49, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d,
	0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x69, 0x67, 0x6e, 0x61,
	0x6c, 0x73, 0x2f, 0x7b, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f,
	0x69, 0x64, 0x7d, 0x7e, 0x7b, 0x63, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x7d, 0x42, 0x08, 0x0a, 0x06, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x42, 0x87, 0x02, 0x0a,
	0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0x15, 0x41, 0x73, 0x73, 0x65, 0x74, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x4b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e,
	0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x35, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x3b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xa2, 0x02, 0x03,
	0x47, 0x41, 0x41, 0xaa, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73,
	0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x35, 0x2e, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xca, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c,
	0x56, 0x31, 0x35, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xea, 0x02, 0x26,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x35, 0x3a, 0x3a, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v15_resources_asset_group_signal_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v15_resources_asset_group_signal_proto_rawDescData = file_google_ads_googleads_v15_resources_asset_group_signal_proto_rawDesc
)

func file_google_ads_googleads_v15_resources_asset_group_signal_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v15_resources_asset_group_signal_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v15_resources_asset_group_signal_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v15_resources_asset_group_signal_proto_rawDescData)
	})
	return file_google_ads_googleads_v15_resources_asset_group_signal_proto_rawDescData
}

var file_google_ads_googleads_v15_resources_asset_group_signal_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v15_resources_asset_group_signal_proto_goTypes = []interface{}{
	(*AssetGroupSignal)(nil), // 0: google.ads.googleads.v15.resources.AssetGroupSignal
	(enums.AssetGroupSignalApprovalStatusEnum_AssetGroupSignalApprovalStatus)(0), // 1: google.ads.googleads.v15.enums.AssetGroupSignalApprovalStatusEnum.AssetGroupSignalApprovalStatus
	(*common.AudienceInfo)(nil),    // 2: google.ads.googleads.v15.common.AudienceInfo
	(*common.SearchThemeInfo)(nil), // 3: google.ads.googleads.v15.common.SearchThemeInfo
}
var file_google_ads_googleads_v15_resources_asset_group_signal_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v15.resources.AssetGroupSignal.approval_status:type_name -> google.ads.googleads.v15.enums.AssetGroupSignalApprovalStatusEnum.AssetGroupSignalApprovalStatus
	2, // 1: google.ads.googleads.v15.resources.AssetGroupSignal.audience:type_name -> google.ads.googleads.v15.common.AudienceInfo
	3, // 2: google.ads.googleads.v15.resources.AssetGroupSignal.search_theme:type_name -> google.ads.googleads.v15.common.SearchThemeInfo
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v15_resources_asset_group_signal_proto_init() }
func file_google_ads_googleads_v15_resources_asset_group_signal_proto_init() {
	if File_google_ads_googleads_v15_resources_asset_group_signal_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v15_resources_asset_group_signal_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssetGroupSignal); i {
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
	file_google_ads_googleads_v15_resources_asset_group_signal_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*AssetGroupSignal_Audience)(nil),
		(*AssetGroupSignal_SearchTheme)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v15_resources_asset_group_signal_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v15_resources_asset_group_signal_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v15_resources_asset_group_signal_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v15_resources_asset_group_signal_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v15_resources_asset_group_signal_proto = out.File
	file_google_ads_googleads_v15_resources_asset_group_signal_proto_rawDesc = nil
	file_google_ads_googleads_v15_resources_asset_group_signal_proto_goTypes = nil
	file_google_ads_googleads_v15_resources_asset_group_signal_proto_depIdxs = nil
}