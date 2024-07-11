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
// source: google/ads/googleads/v16/services/ad_group_bid_modifier_service.proto

package services

import (
	context "context"
	enums "github.com/dictav/go-genproto-googleads/pb/v16/enums"
	resources "github.com/dictav/go-genproto-googleads/pb/v16/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status1 "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
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
// [AdGroupBidModifierService.MutateAdGroupBidModifiers][google.ads.googleads.v16.services.AdGroupBidModifierService.MutateAdGroupBidModifiers].
type MutateAdGroupBidModifiersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. ID of the customer whose ad group bid modifiers are being
	// modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The list of operations to perform on individual ad group bid
	// modifiers.
	Operations []*AdGroupBidModifierOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	// If true, successful operations will be carried out and invalid
	// operations will return errors. If false, all operations will be carried
	// out in one transaction if and only if they are all valid.
	// Default is false.
	PartialFailure bool `protobuf:"varint,3,opt,name=partial_failure,json=partialFailure,proto3" json:"partial_failure,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly bool `protobuf:"varint,4,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
	// The response content type setting. Determines whether the mutable resource
	// or just the resource name should be returned post mutation.
	ResponseContentType enums.ResponseContentTypeEnum_ResponseContentType `protobuf:"varint,5,opt,name=response_content_type,json=responseContentType,proto3,enum=google.ads.googleads.v16.enums.ResponseContentTypeEnum_ResponseContentType" json:"response_content_type,omitempty"`
}

func (x *MutateAdGroupBidModifiersRequest) Reset() {
	*x = MutateAdGroupBidModifiersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v16_services_ad_group_bid_modifier_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateAdGroupBidModifiersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateAdGroupBidModifiersRequest) ProtoMessage() {}

func (x *MutateAdGroupBidModifiersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v16_services_ad_group_bid_modifier_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateAdGroupBidModifiersRequest.ProtoReflect.Descriptor instead.
func (*MutateAdGroupBidModifiersRequest) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v16_services_ad_group_bid_modifier_service_proto_rawDescGZIP(), []int{0}
}

func (x *MutateAdGroupBidModifiersRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *MutateAdGroupBidModifiersRequest) GetOperations() []*AdGroupBidModifierOperation {
	if x != nil {
		return x.Operations
	}
	return nil
}

func (x *MutateAdGroupBidModifiersRequest) GetPartialFailure() bool {
	if x != nil {
		return x.PartialFailure
	}
	return false
}

func (x *MutateAdGroupBidModifiersRequest) GetValidateOnly() bool {
	if x != nil {
		return x.ValidateOnly
	}
	return false
}

func (x *MutateAdGroupBidModifiersRequest) GetResponseContentType() enums.ResponseContentTypeEnum_ResponseContentType {
	if x != nil {
		return x.ResponseContentType
	}
	return enums.ResponseContentTypeEnum_ResponseContentType(0)
}

// A single operation (create, remove, update) on an ad group bid modifier.
type AdGroupBidModifierOperation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// FieldMask that determines which resource fields are modified in an update.
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,4,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The mutate operation.
	//
	// Types that are assignable to Operation:
	//
	//	*AdGroupBidModifierOperation_Create
	//	*AdGroupBidModifierOperation_Update
	//	*AdGroupBidModifierOperation_Remove
	Operation isAdGroupBidModifierOperation_Operation `protobuf_oneof:"operation"`
}

func (x *AdGroupBidModifierOperation) Reset() {
	*x = AdGroupBidModifierOperation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v16_services_ad_group_bid_modifier_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdGroupBidModifierOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdGroupBidModifierOperation) ProtoMessage() {}

func (x *AdGroupBidModifierOperation) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v16_services_ad_group_bid_modifier_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdGroupBidModifierOperation.ProtoReflect.Descriptor instead.
func (*AdGroupBidModifierOperation) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v16_services_ad_group_bid_modifier_service_proto_rawDescGZIP(), []int{1}
}

func (x *AdGroupBidModifierOperation) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

func (m *AdGroupBidModifierOperation) GetOperation() isAdGroupBidModifierOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (x *AdGroupBidModifierOperation) GetCreate() *resources.AdGroupBidModifier {
	if x, ok := x.GetOperation().(*AdGroupBidModifierOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (x *AdGroupBidModifierOperation) GetUpdate() *resources.AdGroupBidModifier {
	if x, ok := x.GetOperation().(*AdGroupBidModifierOperation_Update); ok {
		return x.Update
	}
	return nil
}

func (x *AdGroupBidModifierOperation) GetRemove() string {
	if x, ok := x.GetOperation().(*AdGroupBidModifierOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

type isAdGroupBidModifierOperation_Operation interface {
	isAdGroupBidModifierOperation_Operation()
}

type AdGroupBidModifierOperation_Create struct {
	// Create operation: No resource name is expected for the new ad group bid
	// modifier.
	Create *resources.AdGroupBidModifier `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type AdGroupBidModifierOperation_Update struct {
	// Update operation: The ad group bid modifier is expected to have a valid
	// resource name.
	Update *resources.AdGroupBidModifier `protobuf:"bytes,2,opt,name=update,proto3,oneof"`
}

type AdGroupBidModifierOperation_Remove struct {
	// Remove operation: A resource name for the removed ad group bid modifier
	// is expected, in this format:
	//
	// `customers/{customer_id}/adGroupBidModifiers/{ad_group_id}~{criterion_id}`
	Remove string `protobuf:"bytes,3,opt,name=remove,proto3,oneof"`
}

func (*AdGroupBidModifierOperation_Create) isAdGroupBidModifierOperation_Operation() {}

func (*AdGroupBidModifierOperation_Update) isAdGroupBidModifierOperation_Operation() {}

func (*AdGroupBidModifierOperation_Remove) isAdGroupBidModifierOperation_Operation() {}

// Response message for ad group bid modifiers mutate.
type MutateAdGroupBidModifiersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (for example, auth
	// errors), we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results []*MutateAdGroupBidModifierResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *MutateAdGroupBidModifiersResponse) Reset() {
	*x = MutateAdGroupBidModifiersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v16_services_ad_group_bid_modifier_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateAdGroupBidModifiersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateAdGroupBidModifiersResponse) ProtoMessage() {}

func (x *MutateAdGroupBidModifiersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v16_services_ad_group_bid_modifier_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateAdGroupBidModifiersResponse.ProtoReflect.Descriptor instead.
func (*MutateAdGroupBidModifiersResponse) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v16_services_ad_group_bid_modifier_service_proto_rawDescGZIP(), []int{2}
}

func (x *MutateAdGroupBidModifiersResponse) GetPartialFailureError() *status.Status {
	if x != nil {
		return x.PartialFailureError
	}
	return nil
}

func (x *MutateAdGroupBidModifiersResponse) GetResults() []*MutateAdGroupBidModifierResult {
	if x != nil {
		return x.Results
	}
	return nil
}

// The result for the criterion mutate.
type MutateAdGroupBidModifierResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Returned for successful operations.
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The mutated ad group bid modifier with only mutable fields after mutate.
	// The field will only be returned when response_content_type is set to
	// "MUTABLE_RESOURCE".
	AdGroupBidModifier *resources.AdGroupBidModifier `protobuf:"bytes,2,opt,name=ad_group_bid_modifier,json=adGroupBidModifier,proto3" json:"ad_group_bid_modifier,omitempty"`
}

func (x *MutateAdGroupBidModifierResult) Reset() {
	*x = MutateAdGroupBidModifierResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v16_services_ad_group_bid_modifier_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateAdGroupBidModifierResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateAdGroupBidModifierResult) ProtoMessage() {}

func (x *MutateAdGroupBidModifierResult) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v16_services_ad_group_bid_modifier_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateAdGroupBidModifierResult.ProtoReflect.Descriptor instead.
func (*MutateAdGroupBidModifierResult) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v16_services_ad_group_bid_modifier_service_proto_rawDescGZIP(), []int{3}
}

func (x *MutateAdGroupBidModifierResult) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *MutateAdGroupBidModifierResult) GetAdGroupBidModifier() *resources.AdGroupBidModifier {
	if x != nil {
		return x.AdGroupBidModifier
	}
	return nil
}

var File_google_ads_googleads_v16_services_ad_group_bid_modifier_service_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v16_services_ad_group_bid_modifier_service_proto_rawDesc = []byte{
	0x0a, 0x45, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x36, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2f, 0x61, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x62, 0x69, 0x64,
	0x5f, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31,
	0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x3a, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2f, 0x76, 0x31, 0x36, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x36,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x61, 0x64, 0x5f, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x5f, 0x62, 0x69, 0x64, 0x5f, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfc, 0x02, 0x0a, 0x20, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x41,
	0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x69, 0x64, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0b, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03,
	0xe0, 0x41, 0x02, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x63, 0x0a, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x3e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x36, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x42,
	0x69, 0x64, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x5f,
	0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x70,
	0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x12, 0x23, 0x0a,
	0x0d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x6e,
	0x6c, 0x79, 0x12, 0x7f, 0x0a, 0x15, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x4b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x36, 0x2e, 0x65, 0x6e, 0x75,
	0x6d, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x13,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x22, 0xd7, 0x02, 0x0a, 0x1b, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x42,
	0x69, 0x64, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61,
	0x73, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x4d, 0x61, 0x73, 0x6b, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b,
	0x12, 0x50, 0x0a, 0x06, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x36, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x36, 0x2e, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x69, 0x64,
	0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x48, 0x00, 0x52, 0x06, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x12, 0x50, 0x0a, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x36, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x36, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x42,
	0x69, 0x64, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x48, 0x00, 0x52, 0x06, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x4a, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x30, 0xfa, 0x41, 0x2d, 0x0a, 0x2b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x69, 0x64, 0x4d, 0x6f,
	0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x48, 0x00, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x42, 0x0b, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xc8, 0x01,
	0x0a, 0x21, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x42,
	0x69, 0x64, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x15, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x66,
	0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x13, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x46,
	0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x5b, 0x0a, 0x07, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x41, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x69,
	0x64, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52,
	0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x22, 0xe2, 0x01, 0x0a, 0x1e, 0x4d, 0x75, 0x74,
	0x61, 0x74, 0x65, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x69, 0x64, 0x4d, 0x6f, 0x64,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x55, 0x0a, 0x0d, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x30, 0xfa, 0x41, 0x2d, 0x0a, 0x2b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x69, 0x64, 0x4d, 0x6f, 0x64, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x69, 0x0a, 0x15, 0x61, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x62,
	0x69, 0x64, 0x5f, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x36, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x36, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x69,
	0x64, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x12, 0x61, 0x64, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x42, 0x69, 0x64, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x32, 0xea, 0x02,
	0x0a, 0x19, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x69, 0x64, 0x4d, 0x6f, 0x64, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x85, 0x02, 0x0a, 0x19,
	0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x69, 0x64,
	0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x12, 0x43, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x31, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x75,
	0x74, 0x61, 0x74, 0x65, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x69, 0x64, 0x4d, 0x6f,
	0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x44,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x42, 0x69, 0x64, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5d, 0xda, 0x41, 0x16, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x2c, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x3e, 0x3a, 0x01, 0x2a, 0x22, 0x39, 0x2f, 0x76, 0x31, 0x36, 0x2f, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x3d, 0x2a, 0x7d, 0x2f, 0x61, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x42, 0x69, 0x64, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x3a, 0x6d, 0x75, 0x74,
	0x61, 0x74, 0x65, 0x1a, 0x45, 0xca, 0x41, 0x18, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0xd2, 0x41, 0x27, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75,
	0x74, 0x68, 0x2f, 0x61, 0x64, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x42, 0x8a, 0x02, 0x0a, 0x25, 0x63,
	0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x42, 0x1e, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x69, 0x64,
	0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x49, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67,
	0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64,
	0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x36, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56,
	0x31, 0x36, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xca, 0x02, 0x21, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41,
	0x64, 0x73, 0x5c, 0x56, 0x31, 0x36, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xea,
	0x02, 0x25, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x36, 0x3a, 0x3a, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v16_services_ad_group_bid_modifier_service_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v16_services_ad_group_bid_modifier_service_proto_rawDescData = file_google_ads_googleads_v16_services_ad_group_bid_modifier_service_proto_rawDesc
)

func file_google_ads_googleads_v16_services_ad_group_bid_modifier_service_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v16_services_ad_group_bid_modifier_service_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v16_services_ad_group_bid_modifier_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v16_services_ad_group_bid_modifier_service_proto_rawDescData)
	})
	return file_google_ads_googleads_v16_services_ad_group_bid_modifier_service_proto_rawDescData
}

var file_google_ads_googleads_v16_services_ad_group_bid_modifier_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_ads_googleads_v16_services_ad_group_bid_modifier_service_proto_goTypes = []any{
	(*MutateAdGroupBidModifiersRequest)(nil),               // 0: google.ads.googleads.v16.services.MutateAdGroupBidModifiersRequest
	(*AdGroupBidModifierOperation)(nil),                    // 1: google.ads.googleads.v16.services.AdGroupBidModifierOperation
	(*MutateAdGroupBidModifiersResponse)(nil),              // 2: google.ads.googleads.v16.services.MutateAdGroupBidModifiersResponse
	(*MutateAdGroupBidModifierResult)(nil),                 // 3: google.ads.googleads.v16.services.MutateAdGroupBidModifierResult
	(enums.ResponseContentTypeEnum_ResponseContentType)(0), // 4: google.ads.googleads.v16.enums.ResponseContentTypeEnum.ResponseContentType
	(*fieldmaskpb.FieldMask)(nil),                          // 5: google.protobuf.FieldMask
	(*resources.AdGroupBidModifier)(nil),                   // 6: google.ads.googleads.v16.resources.AdGroupBidModifier
	(*status.Status)(nil),                                  // 7: google.rpc.Status
}
var file_google_ads_googleads_v16_services_ad_group_bid_modifier_service_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v16.services.MutateAdGroupBidModifiersRequest.operations:type_name -> google.ads.googleads.v16.services.AdGroupBidModifierOperation
	4, // 1: google.ads.googleads.v16.services.MutateAdGroupBidModifiersRequest.response_content_type:type_name -> google.ads.googleads.v16.enums.ResponseContentTypeEnum.ResponseContentType
	5, // 2: google.ads.googleads.v16.services.AdGroupBidModifierOperation.update_mask:type_name -> google.protobuf.FieldMask
	6, // 3: google.ads.googleads.v16.services.AdGroupBidModifierOperation.create:type_name -> google.ads.googleads.v16.resources.AdGroupBidModifier
	6, // 4: google.ads.googleads.v16.services.AdGroupBidModifierOperation.update:type_name -> google.ads.googleads.v16.resources.AdGroupBidModifier
	7, // 5: google.ads.googleads.v16.services.MutateAdGroupBidModifiersResponse.partial_failure_error:type_name -> google.rpc.Status
	3, // 6: google.ads.googleads.v16.services.MutateAdGroupBidModifiersResponse.results:type_name -> google.ads.googleads.v16.services.MutateAdGroupBidModifierResult
	6, // 7: google.ads.googleads.v16.services.MutateAdGroupBidModifierResult.ad_group_bid_modifier:type_name -> google.ads.googleads.v16.resources.AdGroupBidModifier
	0, // 8: google.ads.googleads.v16.services.AdGroupBidModifierService.MutateAdGroupBidModifiers:input_type -> google.ads.googleads.v16.services.MutateAdGroupBidModifiersRequest
	2, // 9: google.ads.googleads.v16.services.AdGroupBidModifierService.MutateAdGroupBidModifiers:output_type -> google.ads.googleads.v16.services.MutateAdGroupBidModifiersResponse
	9, // [9:10] is the sub-list for method output_type
	8, // [8:9] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v16_services_ad_group_bid_modifier_service_proto_init() }
func file_google_ads_googleads_v16_services_ad_group_bid_modifier_service_proto_init() {
	if File_google_ads_googleads_v16_services_ad_group_bid_modifier_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v16_services_ad_group_bid_modifier_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*MutateAdGroupBidModifiersRequest); i {
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
		file_google_ads_googleads_v16_services_ad_group_bid_modifier_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*AdGroupBidModifierOperation); i {
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
		file_google_ads_googleads_v16_services_ad_group_bid_modifier_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*MutateAdGroupBidModifiersResponse); i {
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
		file_google_ads_googleads_v16_services_ad_group_bid_modifier_service_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*MutateAdGroupBidModifierResult); i {
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
	file_google_ads_googleads_v16_services_ad_group_bid_modifier_service_proto_msgTypes[1].OneofWrappers = []any{
		(*AdGroupBidModifierOperation_Create)(nil),
		(*AdGroupBidModifierOperation_Update)(nil),
		(*AdGroupBidModifierOperation_Remove)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v16_services_ad_group_bid_modifier_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_ads_googleads_v16_services_ad_group_bid_modifier_service_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v16_services_ad_group_bid_modifier_service_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v16_services_ad_group_bid_modifier_service_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v16_services_ad_group_bid_modifier_service_proto = out.File
	file_google_ads_googleads_v16_services_ad_group_bid_modifier_service_proto_rawDesc = nil
	file_google_ads_googleads_v16_services_ad_group_bid_modifier_service_proto_goTypes = nil
	file_google_ads_googleads_v16_services_ad_group_bid_modifier_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AdGroupBidModifierServiceClient is the client API for AdGroupBidModifierService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdGroupBidModifierServiceClient interface {
	// Creates, updates, or removes ad group bid modifiers.
	// Operation statuses are returned.
	//
	// List of thrown errors:
	//
	//	[AdGroupBidModifierError]()
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[ContextError]()
	//	[CriterionError]()
	//	[DatabaseError]()
	//	[DistinctError]()
	//	[FieldError]()
	//	[FieldMaskError]()
	//	[HeaderError]()
	//	[IdError]()
	//	[InternalError]()
	//	[MutateError]()
	//	[NewResourceCreationError]()
	//	[NotEmptyError]()
	//	[OperatorError]()
	//	[QuotaError]()
	//	[RangeError]()
	//	[RequestError]()
	//	[ResourceCountLimitExceededError]()
	//	[SizeLimitError]()
	//	[StringFormatError]()
	//	[StringLengthError]()
	MutateAdGroupBidModifiers(ctx context.Context, in *MutateAdGroupBidModifiersRequest, opts ...grpc.CallOption) (*MutateAdGroupBidModifiersResponse, error)
}

type adGroupBidModifierServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdGroupBidModifierServiceClient(cc grpc.ClientConnInterface) AdGroupBidModifierServiceClient {
	return &adGroupBidModifierServiceClient{cc}
}

func (c *adGroupBidModifierServiceClient) MutateAdGroupBidModifiers(ctx context.Context, in *MutateAdGroupBidModifiersRequest, opts ...grpc.CallOption) (*MutateAdGroupBidModifiersResponse, error) {
	out := new(MutateAdGroupBidModifiersResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v16.services.AdGroupBidModifierService/MutateAdGroupBidModifiers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdGroupBidModifierServiceServer is the server API for AdGroupBidModifierService service.
type AdGroupBidModifierServiceServer interface {
	// Creates, updates, or removes ad group bid modifiers.
	// Operation statuses are returned.
	//
	// List of thrown errors:
	//
	//	[AdGroupBidModifierError]()
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[ContextError]()
	//	[CriterionError]()
	//	[DatabaseError]()
	//	[DistinctError]()
	//	[FieldError]()
	//	[FieldMaskError]()
	//	[HeaderError]()
	//	[IdError]()
	//	[InternalError]()
	//	[MutateError]()
	//	[NewResourceCreationError]()
	//	[NotEmptyError]()
	//	[OperatorError]()
	//	[QuotaError]()
	//	[RangeError]()
	//	[RequestError]()
	//	[ResourceCountLimitExceededError]()
	//	[SizeLimitError]()
	//	[StringFormatError]()
	//	[StringLengthError]()
	MutateAdGroupBidModifiers(context.Context, *MutateAdGroupBidModifiersRequest) (*MutateAdGroupBidModifiersResponse, error)
}

// UnimplementedAdGroupBidModifierServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAdGroupBidModifierServiceServer struct {
}

func (*UnimplementedAdGroupBidModifierServiceServer) MutateAdGroupBidModifiers(context.Context, *MutateAdGroupBidModifiersRequest) (*MutateAdGroupBidModifiersResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method MutateAdGroupBidModifiers not implemented")
}

func RegisterAdGroupBidModifierServiceServer(s *grpc.Server, srv AdGroupBidModifierServiceServer) {
	s.RegisterService(&_AdGroupBidModifierService_serviceDesc, srv)
}

func _AdGroupBidModifierService_MutateAdGroupBidModifiers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateAdGroupBidModifiersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdGroupBidModifierServiceServer).MutateAdGroupBidModifiers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v16.services.AdGroupBidModifierService/MutateAdGroupBidModifiers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdGroupBidModifierServiceServer).MutateAdGroupBidModifiers(ctx, req.(*MutateAdGroupBidModifiersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AdGroupBidModifierService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v16.services.AdGroupBidModifierService",
	HandlerType: (*AdGroupBidModifierServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateAdGroupBidModifiers",
			Handler:    _AdGroupBidModifierService_MutateAdGroupBidModifiers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v16/services/ad_group_bid_modifier_service.proto",
}
