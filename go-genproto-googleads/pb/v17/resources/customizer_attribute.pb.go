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
// source: google/ads/googleads/v17/resources/customizer_attribute.proto

// copy from https://github.com/dictav/go-genproto-googleads
// and changed by TeamMomentum

package resources

import (
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

// A customizer attribute.
// Use CustomerCustomizer, CampaignCustomizer, AdGroupCustomizer, or
// AdGroupCriterionCustomizer to associate a customizer attribute and
// set its value at the customer, campaign, ad group, or ad group criterion
// level, respectively.
type CustomizerAttribute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Immutable. The resource name of the customizer attribute.
	// Customizer Attribute resource names have the form:
	//
	// `customers/{customer_id}/customizerAttributes/{customizer_attribute_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The ID of the customizer attribute.
	Id int64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	// Required. Immutable. Name of the customizer attribute. Required. It must
	// have a minimum length of 1 and maximum length of 40. Name of an enabled
	// customizer attribute must be unique (case insensitive).
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Immutable. The type of the customizer attribute.
	Type enums.CustomizerAttributeTypeEnum_CustomizerAttributeType `protobuf:"varint,4,opt,name=type,proto3,enum=google.ads.googleads.v17.enums.CustomizerAttributeTypeEnum_CustomizerAttributeType" json:"type,omitempty"`
	// Output only. The status of the customizer attribute.
	Status enums.CustomizerAttributeStatusEnum_CustomizerAttributeStatus `protobuf:"varint,5,opt,name=status,proto3,enum=google.ads.googleads.v17.enums.CustomizerAttributeStatusEnum_CustomizerAttributeStatus" json:"status,omitempty"`
}

func (x *CustomizerAttribute) Reset() {
	*x = CustomizerAttribute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v17_resources_customizer_attribute_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomizerAttribute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomizerAttribute) ProtoMessage() {}

func (x *CustomizerAttribute) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v17_resources_customizer_attribute_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomizerAttribute.ProtoReflect.Descriptor instead.
func (*CustomizerAttribute) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v17_resources_customizer_attribute_proto_rawDescGZIP(), []int{0}
}

func (x *CustomizerAttribute) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *CustomizerAttribute) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CustomizerAttribute) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CustomizerAttribute) GetType() enums.CustomizerAttributeTypeEnum_CustomizerAttributeType {
	if x != nil {
		return x.Type
	}
	return enums.CustomizerAttributeTypeEnum_CustomizerAttributeType(0)
}

func (x *CustomizerAttribute) GetStatus() enums.CustomizerAttributeStatusEnum_CustomizerAttributeStatus {
	if x != nil {
		return x.Status
	}
	return enums.CustomizerAttributeStatusEnum_CustomizerAttributeStatus(0)
}

var File_google_ads_googleads_v17_resources_customizer_attribute_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v17_resources_customizer_attribute_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x37, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x65, 0x72, 0x5f,
	0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x1a, 0x40, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x37, 0x2f, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x65, 0x72, 0x5f, 0x61,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64,
	0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x37, 0x2f,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x65, 0x72,
	0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x80, 0x04, 0x0a, 0x13, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x65, 0x72,
	0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x12, 0x59, 0x0a, 0x0d, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x34, 0xe0, 0x41, 0x05, 0xfa, 0x41, 0x2e, 0x0a, 0x2c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x65, 0x72, 0x41, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe0, 0x41, 0x02, 0xe0, 0x41, 0x05, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x6c, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x53, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x65, 0x72, 0x41,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d,
	0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x65, 0x72, 0x41, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x74, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x57, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x65, 0x72, 0x41,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e,
	0x75, 0x6d, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x65, 0x72, 0x41, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x03, 0xe0, 0x41,
	0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x3a, 0x79, 0xea, 0x41, 0x76, 0x0a, 0x2c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69,
	0x7a, 0x65, 0x72, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x12, 0x46, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x65, 0x72,
	0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x69, 0x7a, 0x65, 0x72, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x5f, 0x69, 0x64, 0x7d, 0x42, 0x8a, 0x02, 0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42,
	0x18, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x65, 0x72, 0x41, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2f, 0x76, 0x31, 0x37, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x3b, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02,
	0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x37, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0xca, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73,
	0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x37, 0x5c, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xea, 0x02, 0x26, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64,
	0x73, 0x3a, 0x3a, 0x56, 0x31, 0x37, 0x3a, 0x3a, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v17_resources_customizer_attribute_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v17_resources_customizer_attribute_proto_rawDescData = file_google_ads_googleads_v17_resources_customizer_attribute_proto_rawDesc
)

func file_google_ads_googleads_v17_resources_customizer_attribute_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v17_resources_customizer_attribute_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v17_resources_customizer_attribute_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v17_resources_customizer_attribute_proto_rawDescData)
	})
	return file_google_ads_googleads_v17_resources_customizer_attribute_proto_rawDescData
}

var file_google_ads_googleads_v17_resources_customizer_attribute_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v17_resources_customizer_attribute_proto_goTypes = []any{
	(*CustomizerAttribute)(nil),                                        // 0: google.ads.googleads.v17.resources.CustomizerAttribute
	(enums.CustomizerAttributeTypeEnum_CustomizerAttributeType)(0),     // 1: google.ads.googleads.v17.enums.CustomizerAttributeTypeEnum.CustomizerAttributeType
	(enums.CustomizerAttributeStatusEnum_CustomizerAttributeStatus)(0), // 2: google.ads.googleads.v17.enums.CustomizerAttributeStatusEnum.CustomizerAttributeStatus
}
var file_google_ads_googleads_v17_resources_customizer_attribute_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v17.resources.CustomizerAttribute.type:type_name -> google.ads.googleads.v17.enums.CustomizerAttributeTypeEnum.CustomizerAttributeType
	2, // 1: google.ads.googleads.v17.resources.CustomizerAttribute.status:type_name -> google.ads.googleads.v17.enums.CustomizerAttributeStatusEnum.CustomizerAttributeStatus
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v17_resources_customizer_attribute_proto_init() }
func file_google_ads_googleads_v17_resources_customizer_attribute_proto_init() {
	if File_google_ads_googleads_v17_resources_customizer_attribute_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v17_resources_customizer_attribute_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CustomizerAttribute); i {
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
			RawDescriptor: file_google_ads_googleads_v17_resources_customizer_attribute_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v17_resources_customizer_attribute_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v17_resources_customizer_attribute_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v17_resources_customizer_attribute_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v17_resources_customizer_attribute_proto = out.File
	file_google_ads_googleads_v17_resources_customizer_attribute_proto_rawDesc = nil
	file_google_ads_googleads_v17_resources_customizer_attribute_proto_goTypes = nil
	file_google_ads_googleads_v17_resources_customizer_attribute_proto_depIdxs = nil
}
