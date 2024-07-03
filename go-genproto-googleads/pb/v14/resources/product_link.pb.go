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
// source: google/ads/googleads/v14/resources/product_link.proto

package resources

import (
	enums "github.com/dictav/go-genproto-googleads/pb/v14/enums"
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

// Represents the data sharing connection between  a Google
// Ads customer and another product.
type ProductLink struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Immutable. Resource name of the product link.
	// ProductLink resource names have the form:
	//
	// `customers/{customer_id}/productLinks/{product_link_id} `
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The ID of the link.
	// This field is read only.
	ProductLinkId *int64 `protobuf:"varint,2,opt,name=product_link_id,json=productLinkId,proto3,oneof" json:"product_link_id,omitempty"`
	// Output only. The type of the linked product.
	Type enums.LinkedProductTypeEnum_LinkedProductType `protobuf:"varint,3,opt,name=type,proto3,enum=google.ads.googleads.v14.enums.LinkedProductTypeEnum_LinkedProductType" json:"type,omitempty"`
	// A product linked to this account.
	//
	// Types that are assignable to LinkedProduct:
	//	*ProductLink_DataPartner
	//	*ProductLink_GoogleAds
	LinkedProduct isProductLink_LinkedProduct `protobuf_oneof:"linked_product"`
}

func (x *ProductLink) Reset() {
	*x = ProductLink{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v14_resources_product_link_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductLink) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductLink) ProtoMessage() {}

func (x *ProductLink) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v14_resources_product_link_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductLink.ProtoReflect.Descriptor instead.
func (*ProductLink) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v14_resources_product_link_proto_rawDescGZIP(), []int{0}
}

func (x *ProductLink) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *ProductLink) GetProductLinkId() int64 {
	if x != nil && x.ProductLinkId != nil {
		return *x.ProductLinkId
	}
	return 0
}

func (x *ProductLink) GetType() enums.LinkedProductTypeEnum_LinkedProductType {
	if x != nil {
		return x.Type
	}
	return enums.LinkedProductTypeEnum_LinkedProductType(0)
}

func (m *ProductLink) GetLinkedProduct() isProductLink_LinkedProduct {
	if m != nil {
		return m.LinkedProduct
	}
	return nil
}

func (x *ProductLink) GetDataPartner() *DataPartnerIdentifier {
	if x, ok := x.GetLinkedProduct().(*ProductLink_DataPartner); ok {
		return x.DataPartner
	}
	return nil
}

func (x *ProductLink) GetGoogleAds() *GoogleAdsIdentifier {
	if x, ok := x.GetLinkedProduct().(*ProductLink_GoogleAds); ok {
		return x.GoogleAds
	}
	return nil
}

type isProductLink_LinkedProduct interface {
	isProductLink_LinkedProduct()
}

type ProductLink_DataPartner struct {
	// Immutable. Data partner link.
	DataPartner *DataPartnerIdentifier `protobuf:"bytes,4,opt,name=data_partner,json=dataPartner,proto3,oneof"`
}

type ProductLink_GoogleAds struct {
	// Immutable. Google Ads link.
	GoogleAds *GoogleAdsIdentifier `protobuf:"bytes,5,opt,name=google_ads,json=googleAds,proto3,oneof"`
}

func (*ProductLink_DataPartner) isProductLink_LinkedProduct() {}

func (*ProductLink_GoogleAds) isProductLink_LinkedProduct() {}

// The identifier for Data Partner account.
type DataPartnerIdentifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Immutable. The customer ID of the Data partner account.
	// This field is required and should not be empty when creating a new
	// data partner link. It is unable to be modified after the creation of
	// the link.
	DataPartnerId *int64 `protobuf:"varint,1,opt,name=data_partner_id,json=dataPartnerId,proto3,oneof" json:"data_partner_id,omitempty"`
}

func (x *DataPartnerIdentifier) Reset() {
	*x = DataPartnerIdentifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v14_resources_product_link_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataPartnerIdentifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataPartnerIdentifier) ProtoMessage() {}

func (x *DataPartnerIdentifier) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v14_resources_product_link_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataPartnerIdentifier.ProtoReflect.Descriptor instead.
func (*DataPartnerIdentifier) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v14_resources_product_link_proto_rawDescGZIP(), []int{1}
}

func (x *DataPartnerIdentifier) GetDataPartnerId() int64 {
	if x != nil && x.DataPartnerId != nil {
		return *x.DataPartnerId
	}
	return 0
}

// The identifier for Google Ads account.
type GoogleAdsIdentifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Immutable. The resource name of the Google Ads account.
	// This field is required and should not be empty when creating a new
	// Google Ads link. It is unable to be modified after the creation of
	// the link.
	Customer *string `protobuf:"bytes,1,opt,name=customer,proto3,oneof" json:"customer,omitempty"`
}

func (x *GoogleAdsIdentifier) Reset() {
	*x = GoogleAdsIdentifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v14_resources_product_link_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoogleAdsIdentifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoogleAdsIdentifier) ProtoMessage() {}

func (x *GoogleAdsIdentifier) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v14_resources_product_link_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoogleAdsIdentifier.ProtoReflect.Descriptor instead.
func (*GoogleAdsIdentifier) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v14_resources_product_link_proto_rawDescGZIP(), []int{2}
}

func (x *GoogleAdsIdentifier) GetCustomer() string {
	if x != nil && x.Customer != nil {
		return *x.Customer
	}
	return ""
}

var File_google_ads_googleads_v14_resources_product_link_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v14_resources_product_link_proto_rawDesc = []byte{
	0x0a, 0x35, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x34, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6c, 0x69, 0x6e,
	0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31,
	0x34, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0x38, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2f, 0x76, 0x31, 0x34, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x6c, 0x69, 0x6e, 0x6b,
	0x65, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xc1, 0x04, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x6e,
	0x6b, 0x12, 0x51, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2c, 0xe0, 0x41, 0x05, 0xfa, 0x41, 0x26,
	0x0a, 0x24, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f,
	0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0,
	0x41, 0x03, 0x48, 0x01, 0x52, 0x0d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x6e,
	0x6b, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x60, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x47, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x34, 0x2e,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x4c, 0x69, 0x6e, 0x6b,
	0x65, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x42, 0x03, 0xe0,
	0x41, 0x03, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x63, 0x0a, 0x0c, 0x64, 0x61, 0x74, 0x61,
	0x5f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x39,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x34, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x48, 0x00,
	0x52, 0x0b, 0x64, 0x61, 0x74, 0x61, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x12, 0x5d, 0x0a,
	0x0a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5f, 0x61, 0x64, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x37, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x34, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x48,
	0x00, 0x52, 0x09, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x61, 0xea, 0x41,
	0x5e, 0x0a, 0x24, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x36, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d,
	0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x2f, 0x7b, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x69, 0x64, 0x7d, 0x42,
	0x10, 0x0a, 0x0e, 0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6c, 0x69,
	0x6e, 0x6b, 0x5f, 0x69, 0x64, 0x22, 0x5d, 0x0a, 0x15, 0x44, 0x61, 0x74, 0x61, 0x50, 0x61, 0x72,
	0x74, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x30,
	0x0a, 0x0f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x48, 0x00, 0x52, 0x0d,
	0x64, 0x61, 0x74, 0x61, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01,
	0x42, 0x12, 0x0a, 0x10, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x22, 0x6e, 0x0a, 0x13, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64,
	0x73, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x4a, 0x0a, 0x08, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x29, 0xe0,
	0x41, 0x05, 0xfa, 0x41, 0x23, 0x0a, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x48, 0x00, 0x52, 0x08, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x42, 0x82, 0x02, 0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x76, 0x31, 0x34, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42,
	0x10, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x4b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x34, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x3b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31,
	0x34, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xca, 0x02, 0x22, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41,
	0x64, 0x73, 0x5c, 0x56, 0x31, 0x34, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0xea, 0x02, 0x26, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x34, 0x3a, 0x3a,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_google_ads_googleads_v14_resources_product_link_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v14_resources_product_link_proto_rawDescData = file_google_ads_googleads_v14_resources_product_link_proto_rawDesc
)

func file_google_ads_googleads_v14_resources_product_link_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v14_resources_product_link_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v14_resources_product_link_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v14_resources_product_link_proto_rawDescData)
	})
	return file_google_ads_googleads_v14_resources_product_link_proto_rawDescData
}

var file_google_ads_googleads_v14_resources_product_link_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_ads_googleads_v14_resources_product_link_proto_goTypes = []interface{}{
	(*ProductLink)(nil),                                // 0: google.ads.googleads.v14.resources.ProductLink
	(*DataPartnerIdentifier)(nil),                      // 1: google.ads.googleads.v14.resources.DataPartnerIdentifier
	(*GoogleAdsIdentifier)(nil),                        // 2: google.ads.googleads.v14.resources.GoogleAdsIdentifier
	(enums.LinkedProductTypeEnum_LinkedProductType)(0), // 3: google.ads.googleads.v14.enums.LinkedProductTypeEnum.LinkedProductType
}
var file_google_ads_googleads_v14_resources_product_link_proto_depIdxs = []int32{
	3, // 0: google.ads.googleads.v14.resources.ProductLink.type:type_name -> google.ads.googleads.v14.enums.LinkedProductTypeEnum.LinkedProductType
	1, // 1: google.ads.googleads.v14.resources.ProductLink.data_partner:type_name -> google.ads.googleads.v14.resources.DataPartnerIdentifier
	2, // 2: google.ads.googleads.v14.resources.ProductLink.google_ads:type_name -> google.ads.googleads.v14.resources.GoogleAdsIdentifier
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v14_resources_product_link_proto_init() }
func file_google_ads_googleads_v14_resources_product_link_proto_init() {
	if File_google_ads_googleads_v14_resources_product_link_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v14_resources_product_link_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductLink); i {
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
		file_google_ads_googleads_v14_resources_product_link_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataPartnerIdentifier); i {
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
		file_google_ads_googleads_v14_resources_product_link_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoogleAdsIdentifier); i {
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
	file_google_ads_googleads_v14_resources_product_link_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ProductLink_DataPartner)(nil),
		(*ProductLink_GoogleAds)(nil),
	}
	file_google_ads_googleads_v14_resources_product_link_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_google_ads_googleads_v14_resources_product_link_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v14_resources_product_link_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v14_resources_product_link_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v14_resources_product_link_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v14_resources_product_link_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v14_resources_product_link_proto = out.File
	file_google_ads_googleads_v14_resources_product_link_proto_rawDesc = nil
	file_google_ads_googleads_v14_resources_product_link_proto_goTypes = nil
	file_google_ads_googleads_v14_resources_product_link_proto_depIdxs = nil
}
