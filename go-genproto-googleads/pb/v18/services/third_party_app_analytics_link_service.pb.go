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
// source: google/ads/googleads/v18/services/third_party_app_analytics_link_service.proto

package services

import (
	context "context"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Request message for
// [ThirdPartyAppAnalyticsLinkService.RegenerateShareableLinkId][google.ads.googleads.v18.services.ThirdPartyAppAnalyticsLinkService.RegenerateShareableLinkId].
type RegenerateShareableLinkIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Resource name of the third party app analytics link.
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
}

func (x *RegenerateShareableLinkIdRequest) Reset() {
	*x = RegenerateShareableLinkIdRequest{}
	mi := &file_google_ads_googleads_v18_services_third_party_app_analytics_link_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegenerateShareableLinkIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegenerateShareableLinkIdRequest) ProtoMessage() {}

func (x *RegenerateShareableLinkIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v18_services_third_party_app_analytics_link_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegenerateShareableLinkIdRequest.ProtoReflect.Descriptor instead.
func (*RegenerateShareableLinkIdRequest) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v18_services_third_party_app_analytics_link_service_proto_rawDescGZIP(), []int{0}
}

func (x *RegenerateShareableLinkIdRequest) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

// Response message for
// [ThirdPartyAppAnalyticsLinkService.RegenerateShareableLinkId][google.ads.googleads.v18.services.ThirdPartyAppAnalyticsLinkService.RegenerateShareableLinkId].
type RegenerateShareableLinkIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RegenerateShareableLinkIdResponse) Reset() {
	*x = RegenerateShareableLinkIdResponse{}
	mi := &file_google_ads_googleads_v18_services_third_party_app_analytics_link_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegenerateShareableLinkIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegenerateShareableLinkIdResponse) ProtoMessage() {}

func (x *RegenerateShareableLinkIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v18_services_third_party_app_analytics_link_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegenerateShareableLinkIdResponse.ProtoReflect.Descriptor instead.
func (*RegenerateShareableLinkIdResponse) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v18_services_third_party_app_analytics_link_service_proto_rawDescGZIP(), []int{1}
}

var File_google_ads_googleads_v18_services_third_party_app_analytics_link_service_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v18_services_third_party_app_analytics_link_service_proto_rawDesc = []byte{
	0x0a, 0x4e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x38, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2f, 0x74, 0x68, 0x69, 0x72, 0x64, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x5f,
	0x61, 0x70, 0x70, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x5f, 0x6c, 0x69,
	0x6e, 0x6b, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x38, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x81, 0x01, 0x0a, 0x20, 0x52, 0x65, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x53, 0x68, 0x61, 0x72, 0x65, 0x61, 0x62, 0x6c, 0x65, 0x4c, 0x69, 0x6e,
	0x6b, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x5d, 0x0a, 0x0d, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x38, 0xfa, 0x41, 0x35, 0x0a, 0x33, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x54, 0x68, 0x69, 0x72, 0x64, 0x50, 0x61, 0x72, 0x74, 0x79, 0x41, 0x70, 0x70, 0x41, 0x6e,
	0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x0c, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x23, 0x0a, 0x21, 0x52, 0x65, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x53, 0x68, 0x61, 0x72, 0x65, 0x61, 0x62, 0x6c, 0x65,
	0x4c, 0x69, 0x6e, 0x6b, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xf8,
	0x02, 0x0a, 0x21, 0x54, 0x68, 0x69, 0x72, 0x64, 0x50, 0x61, 0x72, 0x74, 0x79, 0x41, 0x70, 0x70,
	0x41, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x4c, 0x69, 0x6e, 0x6b, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x8b, 0x02, 0x0a, 0x19, 0x52, 0x65, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x53, 0x68, 0x61, 0x72, 0x65, 0x61, 0x62, 0x6c, 0x65, 0x4c, 0x69, 0x6e, 0x6b,
	0x49, 0x64, 0x12, 0x43, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x38, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x53, 0x68, 0x61, 0x72, 0x65, 0x61, 0x62, 0x6c, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x49, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x44, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x38, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x53, 0x68, 0x61, 0x72, 0x65, 0x61, 0x62, 0x6c, 0x65, 0x4c,
	0x69, 0x6e, 0x6b, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x63, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x5d, 0x3a, 0x01, 0x2a, 0x22, 0x58, 0x2f, 0x76, 0x31, 0x38, 0x2f, 0x7b,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x2a, 0x2f, 0x74, 0x68, 0x69, 0x72, 0x64, 0x50,
	0x61, 0x72, 0x74, 0x79, 0x41, 0x70, 0x70, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73,
	0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x2f, 0x2a, 0x7d, 0x3a, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x53, 0x68, 0x61, 0x72, 0x65, 0x61, 0x62, 0x6c, 0x65, 0x4c, 0x69, 0x6e, 0x6b,
	0x49, 0x64, 0x1a, 0x45, 0xca, 0x41, 0x18, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0xd2,
	0x41, 0x27, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74,
	0x68, 0x2f, 0x61, 0x64, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x42, 0x92, 0x02, 0x0a, 0x25, 0x63, 0x6f,
	0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x38, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x42, 0x26, 0x54, 0x68, 0x69, 0x72, 0x64, 0x50, 0x61, 0x72, 0x74, 0x79, 0x41,
	0x70, 0x70, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x4c, 0x69, 0x6e, 0x6b, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x49, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67,
	0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2f, 0x76, 0x31, 0x38, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x3b,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02,
	0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x38, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0xca, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x38, 0x5c, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xea, 0x02, 0x25, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a,
	0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a,
	0x3a, 0x56, 0x31, 0x38, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v18_services_third_party_app_analytics_link_service_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v18_services_third_party_app_analytics_link_service_proto_rawDescData = file_google_ads_googleads_v18_services_third_party_app_analytics_link_service_proto_rawDesc
)

func file_google_ads_googleads_v18_services_third_party_app_analytics_link_service_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v18_services_third_party_app_analytics_link_service_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v18_services_third_party_app_analytics_link_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v18_services_third_party_app_analytics_link_service_proto_rawDescData)
	})
	return file_google_ads_googleads_v18_services_third_party_app_analytics_link_service_proto_rawDescData
}

var file_google_ads_googleads_v18_services_third_party_app_analytics_link_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_ads_googleads_v18_services_third_party_app_analytics_link_service_proto_goTypes = []any{
	(*RegenerateShareableLinkIdRequest)(nil),  // 0: google.ads.googleads.v18.services.RegenerateShareableLinkIdRequest
	(*RegenerateShareableLinkIdResponse)(nil), // 1: google.ads.googleads.v18.services.RegenerateShareableLinkIdResponse
}
var file_google_ads_googleads_v18_services_third_party_app_analytics_link_service_proto_depIdxs = []int32{
	0, // 0: google.ads.googleads.v18.services.ThirdPartyAppAnalyticsLinkService.RegenerateShareableLinkId:input_type -> google.ads.googleads.v18.services.RegenerateShareableLinkIdRequest
	1, // 1: google.ads.googleads.v18.services.ThirdPartyAppAnalyticsLinkService.RegenerateShareableLinkId:output_type -> google.ads.googleads.v18.services.RegenerateShareableLinkIdResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_google_ads_googleads_v18_services_third_party_app_analytics_link_service_proto_init()
}
func file_google_ads_googleads_v18_services_third_party_app_analytics_link_service_proto_init() {
	if File_google_ads_googleads_v18_services_third_party_app_analytics_link_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v18_services_third_party_app_analytics_link_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_ads_googleads_v18_services_third_party_app_analytics_link_service_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v18_services_third_party_app_analytics_link_service_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v18_services_third_party_app_analytics_link_service_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v18_services_third_party_app_analytics_link_service_proto = out.File
	file_google_ads_googleads_v18_services_third_party_app_analytics_link_service_proto_rawDesc = nil
	file_google_ads_googleads_v18_services_third_party_app_analytics_link_service_proto_goTypes = nil
	file_google_ads_googleads_v18_services_third_party_app_analytics_link_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ThirdPartyAppAnalyticsLinkServiceClient is the client API for ThirdPartyAppAnalyticsLinkService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ThirdPartyAppAnalyticsLinkServiceClient interface {
	// Regenerate ThirdPartyAppAnalyticsLink.shareable_link_id that should be
	// provided to the third party when setting up app analytics.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RequestError]()
	RegenerateShareableLinkId(ctx context.Context, in *RegenerateShareableLinkIdRequest, opts ...grpc.CallOption) (*RegenerateShareableLinkIdResponse, error)
}

type thirdPartyAppAnalyticsLinkServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewThirdPartyAppAnalyticsLinkServiceClient(cc grpc.ClientConnInterface) ThirdPartyAppAnalyticsLinkServiceClient {
	return &thirdPartyAppAnalyticsLinkServiceClient{cc}
}

func (c *thirdPartyAppAnalyticsLinkServiceClient) RegenerateShareableLinkId(ctx context.Context, in *RegenerateShareableLinkIdRequest, opts ...grpc.CallOption) (*RegenerateShareableLinkIdResponse, error) {
	out := new(RegenerateShareableLinkIdResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v18.services.ThirdPartyAppAnalyticsLinkService/RegenerateShareableLinkId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ThirdPartyAppAnalyticsLinkServiceServer is the server API for ThirdPartyAppAnalyticsLinkService service.
type ThirdPartyAppAnalyticsLinkServiceServer interface {
	// Regenerate ThirdPartyAppAnalyticsLink.shareable_link_id that should be
	// provided to the third party when setting up app analytics.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RequestError]()
	RegenerateShareableLinkId(context.Context, *RegenerateShareableLinkIdRequest) (*RegenerateShareableLinkIdResponse, error)
}

// UnimplementedThirdPartyAppAnalyticsLinkServiceServer can be embedded to have forward compatible implementations.
type UnimplementedThirdPartyAppAnalyticsLinkServiceServer struct {
}

func (*UnimplementedThirdPartyAppAnalyticsLinkServiceServer) RegenerateShareableLinkId(context.Context, *RegenerateShareableLinkIdRequest) (*RegenerateShareableLinkIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegenerateShareableLinkId not implemented")
}

func RegisterThirdPartyAppAnalyticsLinkServiceServer(s *grpc.Server, srv ThirdPartyAppAnalyticsLinkServiceServer) {
	s.RegisterService(&_ThirdPartyAppAnalyticsLinkService_serviceDesc, srv)
}

func _ThirdPartyAppAnalyticsLinkService_RegenerateShareableLinkId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegenerateShareableLinkIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThirdPartyAppAnalyticsLinkServiceServer).RegenerateShareableLinkId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v18.services.ThirdPartyAppAnalyticsLinkService/RegenerateShareableLinkId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThirdPartyAppAnalyticsLinkServiceServer).RegenerateShareableLinkId(ctx, req.(*RegenerateShareableLinkIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ThirdPartyAppAnalyticsLinkService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v18.services.ThirdPartyAppAnalyticsLinkService",
	HandlerType: (*ThirdPartyAppAnalyticsLinkServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegenerateShareableLinkId",
			Handler:    _ThirdPartyAppAnalyticsLinkService_RegenerateShareableLinkId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v18/services/third_party_app_analytics_link_service.proto",
}
