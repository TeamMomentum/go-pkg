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
// source: google/ads/googleads/v13/services/travel_asset_suggestion_service.proto

package services

import (
	context "context"
	enums "github.com/dictav/go-genproto-googleads/pb/v13/enums"
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
// [TravelSuggestAssetsService.SuggestTravelAssets][].
type SuggestTravelAssetsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The ID of the customer.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The language specifications in BCP 47 format (for example, en-US,
	// zh-CN, etc.) for the asset suggestions. Text will be in this language.
	// Usually matches one of the campaign target languages.
	LanguageOption string `protobuf:"bytes,2,opt,name=language_option,json=languageOption,proto3" json:"language_option,omitempty"`
	// The Google Maps Place IDs of hotels for which assets are requested. See
	// https://developers.google.com/places/web-service/place-id for more
	// information.
	PlaceId []string `protobuf:"bytes,3,rep,name=place_id,json=placeId,proto3" json:"place_id,omitempty"`
}

func (x *SuggestTravelAssetsRequest) Reset() {
	*x = SuggestTravelAssetsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v13_services_travel_asset_suggestion_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuggestTravelAssetsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuggestTravelAssetsRequest) ProtoMessage() {}

func (x *SuggestTravelAssetsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v13_services_travel_asset_suggestion_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuggestTravelAssetsRequest.ProtoReflect.Descriptor instead.
func (*SuggestTravelAssetsRequest) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v13_services_travel_asset_suggestion_service_proto_rawDescGZIP(), []int{0}
}

func (x *SuggestTravelAssetsRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *SuggestTravelAssetsRequest) GetLanguageOption() string {
	if x != nil {
		return x.LanguageOption
	}
	return ""
}

func (x *SuggestTravelAssetsRequest) GetPlaceId() []string {
	if x != nil {
		return x.PlaceId
	}
	return nil
}

// Response message for
// [TravelSuggestAssetsService.SuggestTravelAssets][].
type SuggestTravelAssetsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Asset suggestions for each place ID submitted in the request.
	HotelAssetSuggestions []*HotelAssetSuggestion `protobuf:"bytes,1,rep,name=hotel_asset_suggestions,json=hotelAssetSuggestions,proto3" json:"hotel_asset_suggestions,omitempty"`
}

func (x *SuggestTravelAssetsResponse) Reset() {
	*x = SuggestTravelAssetsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v13_services_travel_asset_suggestion_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuggestTravelAssetsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuggestTravelAssetsResponse) ProtoMessage() {}

func (x *SuggestTravelAssetsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v13_services_travel_asset_suggestion_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuggestTravelAssetsResponse.ProtoReflect.Descriptor instead.
func (*SuggestTravelAssetsResponse) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v13_services_travel_asset_suggestion_service_proto_rawDescGZIP(), []int{1}
}

func (x *SuggestTravelAssetsResponse) GetHotelAssetSuggestions() []*HotelAssetSuggestion {
	if x != nil {
		return x.HotelAssetSuggestions
	}
	return nil
}

// Message containing the asset suggestions for a hotel.
type HotelAssetSuggestion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Google Places ID of the hotel.
	PlaceId string `protobuf:"bytes,1,opt,name=place_id,json=placeId,proto3" json:"place_id,omitempty"`
	// Suggested final URL for an AssetGroup.
	FinalUrl string `protobuf:"bytes,2,opt,name=final_url,json=finalUrl,proto3" json:"final_url,omitempty"`
	// Hotel name in requested language.
	HotelName string `protobuf:"bytes,3,opt,name=hotel_name,json=hotelName,proto3" json:"hotel_name,omitempty"`
	// Call to action type.
	CallToAction enums.CallToActionTypeEnum_CallToActionType `protobuf:"varint,4,opt,name=call_to_action,json=callToAction,proto3,enum=google.ads.googleads.v13.enums.CallToActionTypeEnum_CallToActionType" json:"call_to_action,omitempty"`
	// Text assets such as headline, description, etc.
	TextAssets []*HotelTextAsset `protobuf:"bytes,5,rep,name=text_assets,json=textAssets,proto3" json:"text_assets,omitempty"`
	// Image assets such as landscape/portrait/square, etc.
	ImageAssets []*HotelImageAsset `protobuf:"bytes,6,rep,name=image_assets,json=imageAssets,proto3" json:"image_assets,omitempty"`
	// The status of the hotel asset suggestion.
	Status enums.HotelAssetSuggestionStatusEnum_HotelAssetSuggestionStatus `protobuf:"varint,7,opt,name=status,proto3,enum=google.ads.googleads.v13.enums.HotelAssetSuggestionStatusEnum_HotelAssetSuggestionStatus" json:"status,omitempty"`
}

func (x *HotelAssetSuggestion) Reset() {
	*x = HotelAssetSuggestion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v13_services_travel_asset_suggestion_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HotelAssetSuggestion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HotelAssetSuggestion) ProtoMessage() {}

func (x *HotelAssetSuggestion) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v13_services_travel_asset_suggestion_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HotelAssetSuggestion.ProtoReflect.Descriptor instead.
func (*HotelAssetSuggestion) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v13_services_travel_asset_suggestion_service_proto_rawDescGZIP(), []int{2}
}

func (x *HotelAssetSuggestion) GetPlaceId() string {
	if x != nil {
		return x.PlaceId
	}
	return ""
}

func (x *HotelAssetSuggestion) GetFinalUrl() string {
	if x != nil {
		return x.FinalUrl
	}
	return ""
}

func (x *HotelAssetSuggestion) GetHotelName() string {
	if x != nil {
		return x.HotelName
	}
	return ""
}

func (x *HotelAssetSuggestion) GetCallToAction() enums.CallToActionTypeEnum_CallToActionType {
	if x != nil {
		return x.CallToAction
	}
	return enums.CallToActionTypeEnum_CallToActionType(0)
}

func (x *HotelAssetSuggestion) GetTextAssets() []*HotelTextAsset {
	if x != nil {
		return x.TextAssets
	}
	return nil
}

func (x *HotelAssetSuggestion) GetImageAssets() []*HotelImageAsset {
	if x != nil {
		return x.ImageAssets
	}
	return nil
}

func (x *HotelAssetSuggestion) GetStatus() enums.HotelAssetSuggestionStatusEnum_HotelAssetSuggestionStatus {
	if x != nil {
		return x.Status
	}
	return enums.HotelAssetSuggestionStatusEnum_HotelAssetSuggestionStatus(0)
}

// A single text asset suggestion for a hotel.
type HotelTextAsset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Asset text in requested language.
	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	// The text asset type. For example, HEADLINE, DESCRIPTION, etc.
	AssetFieldType enums.AssetFieldTypeEnum_AssetFieldType `protobuf:"varint,2,opt,name=asset_field_type,json=assetFieldType,proto3,enum=google.ads.googleads.v13.enums.AssetFieldTypeEnum_AssetFieldType" json:"asset_field_type,omitempty"`
}

func (x *HotelTextAsset) Reset() {
	*x = HotelTextAsset{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v13_services_travel_asset_suggestion_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HotelTextAsset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HotelTextAsset) ProtoMessage() {}

func (x *HotelTextAsset) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v13_services_travel_asset_suggestion_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HotelTextAsset.ProtoReflect.Descriptor instead.
func (*HotelTextAsset) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v13_services_travel_asset_suggestion_service_proto_rawDescGZIP(), []int{3}
}

func (x *HotelTextAsset) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *HotelTextAsset) GetAssetFieldType() enums.AssetFieldTypeEnum_AssetFieldType {
	if x != nil {
		return x.AssetFieldType
	}
	return enums.AssetFieldTypeEnum_AssetFieldType(0)
}

// A single image asset suggestion for a hotel.
type HotelImageAsset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// URI for the image.
	Uri string `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
	// The Image asset type. For example, MARKETING_IMAGE,
	// PORTRAIT_MARKETING_IMAGE, etc.
	AssetFieldType enums.AssetFieldTypeEnum_AssetFieldType `protobuf:"varint,2,opt,name=asset_field_type,json=assetFieldType,proto3,enum=google.ads.googleads.v13.enums.AssetFieldTypeEnum_AssetFieldType" json:"asset_field_type,omitempty"`
}

func (x *HotelImageAsset) Reset() {
	*x = HotelImageAsset{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v13_services_travel_asset_suggestion_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HotelImageAsset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HotelImageAsset) ProtoMessage() {}

func (x *HotelImageAsset) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v13_services_travel_asset_suggestion_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HotelImageAsset.ProtoReflect.Descriptor instead.
func (*HotelImageAsset) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v13_services_travel_asset_suggestion_service_proto_rawDescGZIP(), []int{4}
}

func (x *HotelImageAsset) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *HotelImageAsset) GetAssetFieldType() enums.AssetFieldTypeEnum_AssetFieldType {
	if x != nil {
		return x.AssetFieldType
	}
	return enums.AssetFieldTypeEnum_AssetFieldType(0)
}

var File_google_ads_googleads_v13_services_travel_asset_suggestion_service_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v13_services_travel_asset_suggestion_service_proto_rawDesc = []byte{
	0x0a, 0x47, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x33, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2f, 0x74, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74,
	0x5f, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x33, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x35, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2f, 0x76, 0x31, 0x33, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x38, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x33, 0x2f, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x2f, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x74, 0x6f, 0x5f, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x42, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x33, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x68, 0x6f,
	0x74, 0x65, 0x6c, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76,
	0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8b, 0x01, 0x0a, 0x1a, 0x53, 0x75,
	0x67, 0x67, 0x65, 0x73, 0x74, 0x54, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x41, 0x73, 0x73, 0x65, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0,
	0x41, 0x02, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2c,
	0x0a, 0x0f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0e, 0x6c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08,
	0x70, 0x6c, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07,
	0x70, 0x6c, 0x61, 0x63, 0x65, 0x49, 0x64, 0x22, 0x8e, 0x01, 0x0a, 0x1b, 0x53, 0x75, 0x67, 0x67,
	0x65, 0x73, 0x74, 0x54, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6f, 0x0a, 0x17, 0x68, 0x6f, 0x74, 0x65, 0x6c,
	0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x33, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x48, 0x6f, 0x74,
	0x65, 0x6c, 0x41, 0x73, 0x73, 0x65, 0x74, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x15, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x41, 0x73, 0x73, 0x65, 0x74, 0x53, 0x75, 0x67,
	0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xf8, 0x03, 0x0a, 0x14, 0x48, 0x6f, 0x74,
	0x65, 0x6c, 0x41, 0x73, 0x73, 0x65, 0x74, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x66, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x72, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x68, 0x6f, 0x74,
	0x65, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x68,
	0x6f, 0x74, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x6b, 0x0a, 0x0e, 0x63, 0x61, 0x6c, 0x6c,
	0x5f, 0x74, 0x6f, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x45, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x33, 0x2e, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x54, 0x6f, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x54, 0x6f, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x63, 0x61, 0x6c, 0x6c, 0x54, 0x6f, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x52, 0x0a, 0x0b, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x61, 0x73,
	0x73, 0x65, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x76, 0x31, 0x33, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x48,
	0x6f, 0x74, 0x65, 0x6c, 0x54, 0x65, 0x78, 0x74, 0x41, 0x73, 0x73, 0x65, 0x74, 0x52, 0x0a, 0x74,
	0x65, 0x78, 0x74, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x12, 0x55, 0x0a, 0x0c, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x32, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x33, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x41, 0x73,
	0x73, 0x65, 0x74, 0x52, 0x0b, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73,
	0x12, 0x71, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x59, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x33, 0x2e, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x2e, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x41, 0x73, 0x73, 0x65, 0x74, 0x53, 0x75, 0x67, 0x67,
	0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x75, 0x6d,
	0x2e, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x41, 0x73, 0x73, 0x65, 0x74, 0x53, 0x75, 0x67, 0x67, 0x65,
	0x73, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0x91, 0x01, 0x0a, 0x0e, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x54, 0x65, 0x78,
	0x74, 0x41, 0x73, 0x73, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x6b, 0x0a, 0x10, 0x61, 0x73,
	0x73, 0x65, 0x74, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x41, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x33, 0x2e,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0e, 0x61, 0x73, 0x73, 0x65, 0x74, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x22, 0x90, 0x01, 0x0a, 0x0f, 0x48, 0x6f, 0x74, 0x65,
	0x6c, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x41, 0x73, 0x73, 0x65, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x72, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x69, 0x12, 0x6b, 0x0a,
	0x10, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x41, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x33, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x41, 0x73, 0x73, 0x65,
	0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0e, 0x61, 0x73, 0x73, 0x65,
	0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x32, 0xd9, 0x02, 0x0a, 0x1c, 0x54,
	0x72, 0x61, 0x76, 0x65, 0x6c, 0x41, 0x73, 0x73, 0x65, 0x74, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xf1, 0x01, 0x0a, 0x13,
	0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x54, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x41, 0x73, 0x73,
	0x65, 0x74, 0x73, 0x12, 0x3d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x33, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x54,
	0x72, 0x61, 0x76, 0x65, 0x6c, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x3e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x33, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x54, 0x72,
	0x61, 0x76, 0x65, 0x6c, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x5b, 0xda, 0x41, 0x1b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x2c, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x5f, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x37, 0x3a, 0x01, 0x2a, 0x22, 0x32, 0x2f, 0x76, 0x31,
	0x33, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x3d, 0x2a, 0x7d, 0x3a, 0x73, 0x75, 0x67, 0x67,
	0x65, 0x73, 0x74, 0x54, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x1a,
	0x45, 0xca, 0x41, 0x18, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x27, 0x68,
	0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61,
	0x64, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x42, 0x8d, 0x02, 0x0a, 0x25, 0x63, 0x6f, 0x6d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x33, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x42, 0x21, 0x54, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x41, 0x73, 0x73, 0x65, 0x74, 0x53, 0x75, 0x67,
	0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x49, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f,
	0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x33, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31,
	0x33, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xca, 0x02, 0x21, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64,
	0x73, 0x5c, 0x56, 0x31, 0x33, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xea, 0x02,
	0x25, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x33, 0x3a, 0x3a, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v13_services_travel_asset_suggestion_service_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v13_services_travel_asset_suggestion_service_proto_rawDescData = file_google_ads_googleads_v13_services_travel_asset_suggestion_service_proto_rawDesc
)

func file_google_ads_googleads_v13_services_travel_asset_suggestion_service_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v13_services_travel_asset_suggestion_service_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v13_services_travel_asset_suggestion_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v13_services_travel_asset_suggestion_service_proto_rawDescData)
	})
	return file_google_ads_googleads_v13_services_travel_asset_suggestion_service_proto_rawDescData
}

var file_google_ads_googleads_v13_services_travel_asset_suggestion_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_google_ads_googleads_v13_services_travel_asset_suggestion_service_proto_goTypes = []interface{}{
	(*SuggestTravelAssetsRequest)(nil),                                   // 0: google.ads.googleads.v13.services.SuggestTravelAssetsRequest
	(*SuggestTravelAssetsResponse)(nil),                                  // 1: google.ads.googleads.v13.services.SuggestTravelAssetsResponse
	(*HotelAssetSuggestion)(nil),                                         // 2: google.ads.googleads.v13.services.HotelAssetSuggestion
	(*HotelTextAsset)(nil),                                               // 3: google.ads.googleads.v13.services.HotelTextAsset
	(*HotelImageAsset)(nil),                                              // 4: google.ads.googleads.v13.services.HotelImageAsset
	(enums.CallToActionTypeEnum_CallToActionType)(0),                     // 5: google.ads.googleads.v13.enums.CallToActionTypeEnum.CallToActionType
	(enums.HotelAssetSuggestionStatusEnum_HotelAssetSuggestionStatus)(0), // 6: google.ads.googleads.v13.enums.HotelAssetSuggestionStatusEnum.HotelAssetSuggestionStatus
	(enums.AssetFieldTypeEnum_AssetFieldType)(0),                         // 7: google.ads.googleads.v13.enums.AssetFieldTypeEnum.AssetFieldType
}
var file_google_ads_googleads_v13_services_travel_asset_suggestion_service_proto_depIdxs = []int32{
	2, // 0: google.ads.googleads.v13.services.SuggestTravelAssetsResponse.hotel_asset_suggestions:type_name -> google.ads.googleads.v13.services.HotelAssetSuggestion
	5, // 1: google.ads.googleads.v13.services.HotelAssetSuggestion.call_to_action:type_name -> google.ads.googleads.v13.enums.CallToActionTypeEnum.CallToActionType
	3, // 2: google.ads.googleads.v13.services.HotelAssetSuggestion.text_assets:type_name -> google.ads.googleads.v13.services.HotelTextAsset
	4, // 3: google.ads.googleads.v13.services.HotelAssetSuggestion.image_assets:type_name -> google.ads.googleads.v13.services.HotelImageAsset
	6, // 4: google.ads.googleads.v13.services.HotelAssetSuggestion.status:type_name -> google.ads.googleads.v13.enums.HotelAssetSuggestionStatusEnum.HotelAssetSuggestionStatus
	7, // 5: google.ads.googleads.v13.services.HotelTextAsset.asset_field_type:type_name -> google.ads.googleads.v13.enums.AssetFieldTypeEnum.AssetFieldType
	7, // 6: google.ads.googleads.v13.services.HotelImageAsset.asset_field_type:type_name -> google.ads.googleads.v13.enums.AssetFieldTypeEnum.AssetFieldType
	0, // 7: google.ads.googleads.v13.services.TravelAssetSuggestionService.SuggestTravelAssets:input_type -> google.ads.googleads.v13.services.SuggestTravelAssetsRequest
	1, // 8: google.ads.googleads.v13.services.TravelAssetSuggestionService.SuggestTravelAssets:output_type -> google.ads.googleads.v13.services.SuggestTravelAssetsResponse
	8, // [8:9] is the sub-list for method output_type
	7, // [7:8] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v13_services_travel_asset_suggestion_service_proto_init() }
func file_google_ads_googleads_v13_services_travel_asset_suggestion_service_proto_init() {
	if File_google_ads_googleads_v13_services_travel_asset_suggestion_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v13_services_travel_asset_suggestion_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SuggestTravelAssetsRequest); i {
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
		file_google_ads_googleads_v13_services_travel_asset_suggestion_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SuggestTravelAssetsResponse); i {
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
		file_google_ads_googleads_v13_services_travel_asset_suggestion_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HotelAssetSuggestion); i {
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
		file_google_ads_googleads_v13_services_travel_asset_suggestion_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HotelTextAsset); i {
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
		file_google_ads_googleads_v13_services_travel_asset_suggestion_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HotelImageAsset); i {
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
			RawDescriptor: file_google_ads_googleads_v13_services_travel_asset_suggestion_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_ads_googleads_v13_services_travel_asset_suggestion_service_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v13_services_travel_asset_suggestion_service_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v13_services_travel_asset_suggestion_service_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v13_services_travel_asset_suggestion_service_proto = out.File
	file_google_ads_googleads_v13_services_travel_asset_suggestion_service_proto_rawDesc = nil
	file_google_ads_googleads_v13_services_travel_asset_suggestion_service_proto_goTypes = nil
	file_google_ads_googleads_v13_services_travel_asset_suggestion_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TravelAssetSuggestionServiceClient is the client API for TravelAssetSuggestionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TravelAssetSuggestionServiceClient interface {
	// Returns Travel Asset suggestions. Asset
	// suggestions are returned on a best-effort basis. There are no guarantees
	// that all possible asset types will be returned for any given hotel
	// property.
	SuggestTravelAssets(ctx context.Context, in *SuggestTravelAssetsRequest, opts ...grpc.CallOption) (*SuggestTravelAssetsResponse, error)
}

type travelAssetSuggestionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTravelAssetSuggestionServiceClient(cc grpc.ClientConnInterface) TravelAssetSuggestionServiceClient {
	return &travelAssetSuggestionServiceClient{cc}
}

func (c *travelAssetSuggestionServiceClient) SuggestTravelAssets(ctx context.Context, in *SuggestTravelAssetsRequest, opts ...grpc.CallOption) (*SuggestTravelAssetsResponse, error) {
	out := new(SuggestTravelAssetsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v13.services.TravelAssetSuggestionService/SuggestTravelAssets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TravelAssetSuggestionServiceServer is the server API for TravelAssetSuggestionService service.
type TravelAssetSuggestionServiceServer interface {
	// Returns Travel Asset suggestions. Asset
	// suggestions are returned on a best-effort basis. There are no guarantees
	// that all possible asset types will be returned for any given hotel
	// property.
	SuggestTravelAssets(context.Context, *SuggestTravelAssetsRequest) (*SuggestTravelAssetsResponse, error)
}

// UnimplementedTravelAssetSuggestionServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTravelAssetSuggestionServiceServer struct {
}

func (*UnimplementedTravelAssetSuggestionServiceServer) SuggestTravelAssets(context.Context, *SuggestTravelAssetsRequest) (*SuggestTravelAssetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SuggestTravelAssets not implemented")
}

func RegisterTravelAssetSuggestionServiceServer(s *grpc.Server, srv TravelAssetSuggestionServiceServer) {
	s.RegisterService(&_TravelAssetSuggestionService_serviceDesc, srv)
}

func _TravelAssetSuggestionService_SuggestTravelAssets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SuggestTravelAssetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TravelAssetSuggestionServiceServer).SuggestTravelAssets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v13.services.TravelAssetSuggestionService/SuggestTravelAssets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TravelAssetSuggestionServiceServer).SuggestTravelAssets(ctx, req.(*SuggestTravelAssetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TravelAssetSuggestionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v13.services.TravelAssetSuggestionService",
	HandlerType: (*TravelAssetSuggestionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SuggestTravelAssets",
			Handler:    _TravelAssetSuggestionService_SuggestTravelAssets_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v13/services/travel_asset_suggestion_service.proto",
}
