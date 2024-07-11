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
// source: google/ads/googleads/v16/resources/keyword_plan_ad_group.proto

package resources

import (
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

// A Keyword Planner ad group.
// Max number of keyword plan ad groups per plan: 200.
type KeywordPlanAdGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Immutable. The resource name of the Keyword Planner ad group.
	// KeywordPlanAdGroup resource names have the form:
	//
	// `customers/{customer_id}/keywordPlanAdGroups/{kp_ad_group_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The keyword plan campaign to which this ad group belongs.
	KeywordPlanCampaign *string `protobuf:"bytes,6,opt,name=keyword_plan_campaign,json=keywordPlanCampaign,proto3,oneof" json:"keyword_plan_campaign,omitempty"`
	// Output only. The ID of the keyword plan ad group.
	Id *int64 `protobuf:"varint,7,opt,name=id,proto3,oneof" json:"id,omitempty"`
	// The name of the keyword plan ad group.
	//
	// This field is required and should not be empty when creating keyword plan
	// ad group.
	Name *string `protobuf:"bytes,8,opt,name=name,proto3,oneof" json:"name,omitempty"`
	// A default ad group max cpc bid in micros in account currency for all
	// biddable keywords under the keyword plan ad group.
	// If not set, will inherit from parent campaign.
	CpcBidMicros *int64 `protobuf:"varint,9,opt,name=cpc_bid_micros,json=cpcBidMicros,proto3,oneof" json:"cpc_bid_micros,omitempty"`
}

func (x *KeywordPlanAdGroup) Reset() {
	*x = KeywordPlanAdGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v16_resources_keyword_plan_ad_group_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeywordPlanAdGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeywordPlanAdGroup) ProtoMessage() {}

func (x *KeywordPlanAdGroup) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v16_resources_keyword_plan_ad_group_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeywordPlanAdGroup.ProtoReflect.Descriptor instead.
func (*KeywordPlanAdGroup) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v16_resources_keyword_plan_ad_group_proto_rawDescGZIP(), []int{0}
}

func (x *KeywordPlanAdGroup) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *KeywordPlanAdGroup) GetKeywordPlanCampaign() string {
	if x != nil && x.KeywordPlanCampaign != nil {
		return *x.KeywordPlanCampaign
	}
	return ""
}

func (x *KeywordPlanAdGroup) GetId() int64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *KeywordPlanAdGroup) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *KeywordPlanAdGroup) GetCpcBidMicros() int64 {
	if x != nil && x.CpcBidMicros != nil {
		return *x.CpcBidMicros
	}
	return 0
}

var File_google_ads_googleads_v16_resources_keyword_plan_ad_group_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v16_resources_keyword_plan_ad_group_proto_rawDesc = []byte{
	0x0a, 0x3e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x36, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x70, 0x6c, 0x61,
	0x6e, 0x5f, 0x61, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x36, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xef, 0x03, 0x0a, 0x12, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x50, 0x6c, 0x61, 0x6e,
	0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x58, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x33,
	0xe0, 0x41, 0x05, 0xfa, 0x41, 0x2d, 0x0a, 0x2b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x41, 0x64, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x6a, 0x0a, 0x15, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x70, 0x6c, 0x61,
	0x6e, 0x5f, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x31, 0xfa, 0x41, 0x2e, 0x0a, 0x2c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x43, 0x61, 0x6d, 0x70, 0x61,
	0x69, 0x67, 0x6e, 0x48, 0x00, 0x52, 0x13, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x50, 0x6c,
	0x61, 0x6e, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x18, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x01,
	0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01,
	0x12, 0x29, 0x0a, 0x0e, 0x63, 0x70, 0x63, 0x5f, 0x62, 0x69, 0x64, 0x5f, 0x6d, 0x69, 0x63, 0x72,
	0x6f, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x48, 0x03, 0x52, 0x0c, 0x63, 0x70, 0x63, 0x42,
	0x69, 0x64, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x88, 0x01, 0x01, 0x3a, 0x78, 0xea, 0x41, 0x75,
	0x0a, 0x2b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4b, 0x65, 0x79, 0x77, 0x6f,
	0x72, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x46, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x50, 0x6c,
	0x61, 0x6e, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2f, 0x7b, 0x6b, 0x65, 0x79, 0x77,
	0x6f, 0x72, 0x64, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x61, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x18, 0x0a, 0x16, 0x5f, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72,
	0x64, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x42,
	0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42,
	0x11, 0x0a, 0x0f, 0x5f, 0x63, 0x70, 0x63, 0x5f, 0x62, 0x69, 0x64, 0x5f, 0x6d, 0x69, 0x63, 0x72,
	0x6f, 0x73, 0x42, 0x89, 0x02, 0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x36, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0x17, 0x4b,
	0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31,
	0x36, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x3b, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x22, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41,
	0x64, 0x73, 0x2e, 0x56, 0x31, 0x36, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0xca, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x36, 0x5c, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0xea, 0x02, 0x26, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a,
	0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a,
	0x56, 0x31, 0x36, 0x3a, 0x3a, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v16_resources_keyword_plan_ad_group_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v16_resources_keyword_plan_ad_group_proto_rawDescData = file_google_ads_googleads_v16_resources_keyword_plan_ad_group_proto_rawDesc
)

func file_google_ads_googleads_v16_resources_keyword_plan_ad_group_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v16_resources_keyword_plan_ad_group_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v16_resources_keyword_plan_ad_group_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v16_resources_keyword_plan_ad_group_proto_rawDescData)
	})
	return file_google_ads_googleads_v16_resources_keyword_plan_ad_group_proto_rawDescData
}

var file_google_ads_googleads_v16_resources_keyword_plan_ad_group_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v16_resources_keyword_plan_ad_group_proto_goTypes = []any{
	(*KeywordPlanAdGroup)(nil), // 0: google.ads.googleads.v16.resources.KeywordPlanAdGroup
}
var file_google_ads_googleads_v16_resources_keyword_plan_ad_group_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v16_resources_keyword_plan_ad_group_proto_init() }
func file_google_ads_googleads_v16_resources_keyword_plan_ad_group_proto_init() {
	if File_google_ads_googleads_v16_resources_keyword_plan_ad_group_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v16_resources_keyword_plan_ad_group_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*KeywordPlanAdGroup); i {
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
	file_google_ads_googleads_v16_resources_keyword_plan_ad_group_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v16_resources_keyword_plan_ad_group_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v16_resources_keyword_plan_ad_group_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v16_resources_keyword_plan_ad_group_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v16_resources_keyword_plan_ad_group_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v16_resources_keyword_plan_ad_group_proto = out.File
	file_google_ads_googleads_v16_resources_keyword_plan_ad_group_proto_rawDesc = nil
	file_google_ads_googleads_v16_resources_keyword_plan_ad_group_proto_goTypes = nil
	file_google_ads_googleads_v16_resources_keyword_plan_ad_group_proto_depIdxs = nil
}
