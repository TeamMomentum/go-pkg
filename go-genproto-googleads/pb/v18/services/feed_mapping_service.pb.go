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
// source: google/ads/googleads/v18/services/feed_mapping_service.proto

package services

import (
	context "context"
	enums "github.com/TeamMomentum/go-pkg/go-genproto-googleads/pb/v18/enums"
	resources "github.com/TeamMomentum/go-pkg/go-genproto-googleads/pb/v18/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status1 "google.golang.org/grpc/status"
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
// [FeedMappingService.MutateFeedMappings][google.ads.googleads.v18.services.FeedMappingService.MutateFeedMappings].
type MutateFeedMappingsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The ID of the customer whose feed mappings are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The list of operations to perform on individual feed mappings.
	Operations []*FeedMappingOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
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
	ResponseContentType enums.ResponseContentTypeEnum_ResponseContentType `protobuf:"varint,5,opt,name=response_content_type,json=responseContentType,proto3,enum=google.ads.googleads.v18.enums.ResponseContentTypeEnum_ResponseContentType" json:"response_content_type,omitempty"`
}

func (x *MutateFeedMappingsRequest) Reset() {
	*x = MutateFeedMappingsRequest{}
	mi := &file_google_ads_googleads_v18_services_feed_mapping_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MutateFeedMappingsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateFeedMappingsRequest) ProtoMessage() {}

func (x *MutateFeedMappingsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v18_services_feed_mapping_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateFeedMappingsRequest.ProtoReflect.Descriptor instead.
func (*MutateFeedMappingsRequest) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v18_services_feed_mapping_service_proto_rawDescGZIP(), []int{0}
}

func (x *MutateFeedMappingsRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *MutateFeedMappingsRequest) GetOperations() []*FeedMappingOperation {
	if x != nil {
		return x.Operations
	}
	return nil
}

func (x *MutateFeedMappingsRequest) GetPartialFailure() bool {
	if x != nil {
		return x.PartialFailure
	}
	return false
}

func (x *MutateFeedMappingsRequest) GetValidateOnly() bool {
	if x != nil {
		return x.ValidateOnly
	}
	return false
}

func (x *MutateFeedMappingsRequest) GetResponseContentType() enums.ResponseContentTypeEnum_ResponseContentType {
	if x != nil {
		return x.ResponseContentType
	}
	return enums.ResponseContentTypeEnum_ResponseContentType(0)
}

// A single operation (create, remove) on a feed mapping.
type FeedMappingOperation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The mutate operation.
	//
	// Types that are assignable to Operation:
	//
	//	*FeedMappingOperation_Create
	//	*FeedMappingOperation_Remove
	Operation isFeedMappingOperation_Operation `protobuf_oneof:"operation"`
}

func (x *FeedMappingOperation) Reset() {
	*x = FeedMappingOperation{}
	mi := &file_google_ads_googleads_v18_services_feed_mapping_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FeedMappingOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeedMappingOperation) ProtoMessage() {}

func (x *FeedMappingOperation) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v18_services_feed_mapping_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeedMappingOperation.ProtoReflect.Descriptor instead.
func (*FeedMappingOperation) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v18_services_feed_mapping_service_proto_rawDescGZIP(), []int{1}
}

func (m *FeedMappingOperation) GetOperation() isFeedMappingOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (x *FeedMappingOperation) GetCreate() *resources.FeedMapping {
	if x, ok := x.GetOperation().(*FeedMappingOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (x *FeedMappingOperation) GetRemove() string {
	if x, ok := x.GetOperation().(*FeedMappingOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

type isFeedMappingOperation_Operation interface {
	isFeedMappingOperation_Operation()
}

type FeedMappingOperation_Create struct {
	// Create operation: No resource name is expected for the new feed mapping.
	Create *resources.FeedMapping `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type FeedMappingOperation_Remove struct {
	// Remove operation: A resource name for the removed feed mapping is
	// expected, in this format:
	//
	// `customers/{customer_id}/feedMappings/{feed_id}~{feed_mapping_id}`
	Remove string `protobuf:"bytes,3,opt,name=remove,proto3,oneof"`
}

func (*FeedMappingOperation_Create) isFeedMappingOperation_Operation() {}

func (*FeedMappingOperation_Remove) isFeedMappingOperation_Operation() {}

// Response message for a feed mapping mutate.
type MutateFeedMappingsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (for example, auth
	// errors), we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results []*MutateFeedMappingResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *MutateFeedMappingsResponse) Reset() {
	*x = MutateFeedMappingsResponse{}
	mi := &file_google_ads_googleads_v18_services_feed_mapping_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MutateFeedMappingsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateFeedMappingsResponse) ProtoMessage() {}

func (x *MutateFeedMappingsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v18_services_feed_mapping_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateFeedMappingsResponse.ProtoReflect.Descriptor instead.
func (*MutateFeedMappingsResponse) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v18_services_feed_mapping_service_proto_rawDescGZIP(), []int{2}
}

func (x *MutateFeedMappingsResponse) GetPartialFailureError() *status.Status {
	if x != nil {
		return x.PartialFailureError
	}
	return nil
}

func (x *MutateFeedMappingsResponse) GetResults() []*MutateFeedMappingResult {
	if x != nil {
		return x.Results
	}
	return nil
}

// The result for the feed mapping mutate.
type MutateFeedMappingResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Returned for successful operations.
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The mutated feed mapping with only mutable fields after mutate. The field
	// will only be returned when response_content_type is set to
	// "MUTABLE_RESOURCE".
	FeedMapping *resources.FeedMapping `protobuf:"bytes,2,opt,name=feed_mapping,json=feedMapping,proto3" json:"feed_mapping,omitempty"`
}

func (x *MutateFeedMappingResult) Reset() {
	*x = MutateFeedMappingResult{}
	mi := &file_google_ads_googleads_v18_services_feed_mapping_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MutateFeedMappingResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateFeedMappingResult) ProtoMessage() {}

func (x *MutateFeedMappingResult) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v18_services_feed_mapping_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateFeedMappingResult.ProtoReflect.Descriptor instead.
func (*MutateFeedMappingResult) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v18_services_feed_mapping_service_proto_rawDescGZIP(), []int{3}
}

func (x *MutateFeedMappingResult) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *MutateFeedMappingResult) GetFeedMapping() *resources.FeedMapping {
	if x != nil {
		return x.FeedMapping
	}
	return nil
}

var File_google_ads_googleads_v18_services_feed_mapping_service_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v18_services_feed_mapping_service_proto_rawDesc = []byte{
	0x0a, 0x3c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x38, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2f, 0x66, 0x65, 0x65, 0x64, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x38, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x1a, 0x3a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x38, 0x2f, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x35, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x38, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x2f, 0x66, 0x65, 0x65, 0x64, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65,
	0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x72, 0x70, 0x63, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xee, 0x02, 0x0a, 0x19, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x46, 0x65, 0x65, 0x64, 0x4d,
	0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24,
	0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x5c, 0x0a, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x38, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x46, 0x65, 0x65,
	0x64, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x66, 0x61,
	0x69, 0x6c, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x70, 0x61, 0x72,
	0x74, 0x69, 0x61, 0x6c, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0c, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x6e, 0x6c, 0x79,
	0x12, 0x7f, 0x0a, 0x15, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x4b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x38, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x13, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x22, 0xb3, 0x01, 0x0a, 0x14, 0x46, 0x65, 0x65, 0x64, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x49, 0x0a, 0x06, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x76, 0x31, 0x38, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e,
	0x46, 0x65, 0x65, 0x64, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x48, 0x00, 0x52, 0x06, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x43, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x29, 0xfa, 0x41, 0x26, 0x0a, 0x24, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x46, 0x65, 0x65, 0x64, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67,
	0x48, 0x00, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xba, 0x01, 0x0a, 0x1a, 0x4d, 0x75, 0x74, 0x61,
	0x74, 0x65, 0x46, 0x65, 0x65, 0x64, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x15, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61,
	0x6c, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72,
	0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x13, 0x70, 0x61, 0x72, 0x74, 0x69,
	0x61, 0x6c, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x54,
	0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x3a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x38, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x46, 0x65, 0x65, 0x64, 0x4d, 0x61,
	0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x07, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x73, 0x22, 0xbd, 0x01, 0x0a, 0x17, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x46,
	0x65, 0x65, 0x64, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x4e, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x29, 0xfa, 0x41, 0x26, 0x0a, 0x24, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x46, 0x65, 0x65, 0x64, 0x4d, 0x61, 0x70, 0x70, 0x69,
	0x6e, 0x67, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x52, 0x0a, 0x0c, 0x66, 0x65, 0x65, 0x64, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31,
	0x38, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x46, 0x65, 0x65, 0x64,
	0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x0b, 0x66, 0x65, 0x65, 0x64, 0x4d, 0x61, 0x70,
	0x70, 0x69, 0x6e, 0x67, 0x32, 0xc7, 0x02, 0x0a, 0x12, 0x46, 0x65, 0x65, 0x64, 0x4d, 0x61, 0x70,
	0x70, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xe9, 0x01, 0x0a, 0x12,
	0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x46, 0x65, 0x65, 0x64, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x73, 0x12, 0x3c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x38, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x46, 0x65, 0x65,
	0x64, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x3d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x38, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x46, 0x65, 0x65, 0x64, 0x4d,
	0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x56, 0xda, 0x41, 0x16, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x2c,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x37,
	0x3a, 0x01, 0x2a, 0x22, 0x32, 0x2f, 0x76, 0x31, 0x38, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x3d, 0x2a, 0x7d, 0x2f, 0x66, 0x65, 0x65, 0x64, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x73,
	0x3a, 0x6d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x1a, 0x45, 0xca, 0x41, 0x18, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x27, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77,
	0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x64, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x42, 0x83,
	0x02, 0x0a, 0x25, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x38, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x42, 0x17, 0x46, 0x65, 0x65, 0x64, 0x4d, 0x61,
	0x70, 0x70, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x49, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x38, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xa2, 0x02,
	0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64,
	0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x38, 0x2e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xca, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c,
	0x56, 0x31, 0x38, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xea, 0x02, 0x25, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x38, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v18_services_feed_mapping_service_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v18_services_feed_mapping_service_proto_rawDescData = file_google_ads_googleads_v18_services_feed_mapping_service_proto_rawDesc
)

func file_google_ads_googleads_v18_services_feed_mapping_service_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v18_services_feed_mapping_service_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v18_services_feed_mapping_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v18_services_feed_mapping_service_proto_rawDescData)
	})
	return file_google_ads_googleads_v18_services_feed_mapping_service_proto_rawDescData
}

var file_google_ads_googleads_v18_services_feed_mapping_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_ads_googleads_v18_services_feed_mapping_service_proto_goTypes = []any{
	(*MutateFeedMappingsRequest)(nil),                      // 0: google.ads.googleads.v18.services.MutateFeedMappingsRequest
	(*FeedMappingOperation)(nil),                           // 1: google.ads.googleads.v18.services.FeedMappingOperation
	(*MutateFeedMappingsResponse)(nil),                     // 2: google.ads.googleads.v18.services.MutateFeedMappingsResponse
	(*MutateFeedMappingResult)(nil),                        // 3: google.ads.googleads.v18.services.MutateFeedMappingResult
	(enums.ResponseContentTypeEnum_ResponseContentType)(0), // 4: google.ads.googleads.v18.enums.ResponseContentTypeEnum.ResponseContentType
	(*resources.FeedMapping)(nil),                          // 5: google.ads.googleads.v18.resources.FeedMapping
	(*status.Status)(nil),                                  // 6: google.rpc.Status
}
var file_google_ads_googleads_v18_services_feed_mapping_service_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v18.services.MutateFeedMappingsRequest.operations:type_name -> google.ads.googleads.v18.services.FeedMappingOperation
	4, // 1: google.ads.googleads.v18.services.MutateFeedMappingsRequest.response_content_type:type_name -> google.ads.googleads.v18.enums.ResponseContentTypeEnum.ResponseContentType
	5, // 2: google.ads.googleads.v18.services.FeedMappingOperation.create:type_name -> google.ads.googleads.v18.resources.FeedMapping
	6, // 3: google.ads.googleads.v18.services.MutateFeedMappingsResponse.partial_failure_error:type_name -> google.rpc.Status
	3, // 4: google.ads.googleads.v18.services.MutateFeedMappingsResponse.results:type_name -> google.ads.googleads.v18.services.MutateFeedMappingResult
	5, // 5: google.ads.googleads.v18.services.MutateFeedMappingResult.feed_mapping:type_name -> google.ads.googleads.v18.resources.FeedMapping
	0, // 6: google.ads.googleads.v18.services.FeedMappingService.MutateFeedMappings:input_type -> google.ads.googleads.v18.services.MutateFeedMappingsRequest
	2, // 7: google.ads.googleads.v18.services.FeedMappingService.MutateFeedMappings:output_type -> google.ads.googleads.v18.services.MutateFeedMappingsResponse
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v18_services_feed_mapping_service_proto_init() }
func file_google_ads_googleads_v18_services_feed_mapping_service_proto_init() {
	if File_google_ads_googleads_v18_services_feed_mapping_service_proto != nil {
		return
	}
	file_google_ads_googleads_v18_services_feed_mapping_service_proto_msgTypes[1].OneofWrappers = []any{
		(*FeedMappingOperation_Create)(nil),
		(*FeedMappingOperation_Remove)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v18_services_feed_mapping_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_ads_googleads_v18_services_feed_mapping_service_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v18_services_feed_mapping_service_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v18_services_feed_mapping_service_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v18_services_feed_mapping_service_proto = out.File
	file_google_ads_googleads_v18_services_feed_mapping_service_proto_rawDesc = nil
	file_google_ads_googleads_v18_services_feed_mapping_service_proto_goTypes = nil
	file_google_ads_googleads_v18_services_feed_mapping_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FeedMappingServiceClient is the client API for FeedMappingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FeedMappingServiceClient interface {
	// Creates or removes feed mappings. Operation statuses are
	// returned.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[DatabaseError]()
	//	[DistinctError]()
	//	[FeedMappingError]()
	//	[FieldError]()
	//	[HeaderError]()
	//	[IdError]()
	//	[InternalError]()
	//	[MutateError]()
	//	[NotEmptyError]()
	//	[OperationAccessDeniedError]()
	//	[OperatorError]()
	//	[QuotaError]()
	//	[RangeError]()
	//	[RequestError]()
	//	[SizeLimitError]()
	//	[StringFormatError]()
	//	[StringLengthError]()
	MutateFeedMappings(ctx context.Context, in *MutateFeedMappingsRequest, opts ...grpc.CallOption) (*MutateFeedMappingsResponse, error)
}

type feedMappingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFeedMappingServiceClient(cc grpc.ClientConnInterface) FeedMappingServiceClient {
	return &feedMappingServiceClient{cc}
}

func (c *feedMappingServiceClient) MutateFeedMappings(ctx context.Context, in *MutateFeedMappingsRequest, opts ...grpc.CallOption) (*MutateFeedMappingsResponse, error) {
	out := new(MutateFeedMappingsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v18.services.FeedMappingService/MutateFeedMappings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FeedMappingServiceServer is the server API for FeedMappingService service.
type FeedMappingServiceServer interface {
	// Creates or removes feed mappings. Operation statuses are
	// returned.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[DatabaseError]()
	//	[DistinctError]()
	//	[FeedMappingError]()
	//	[FieldError]()
	//	[HeaderError]()
	//	[IdError]()
	//	[InternalError]()
	//	[MutateError]()
	//	[NotEmptyError]()
	//	[OperationAccessDeniedError]()
	//	[OperatorError]()
	//	[QuotaError]()
	//	[RangeError]()
	//	[RequestError]()
	//	[SizeLimitError]()
	//	[StringFormatError]()
	//	[StringLengthError]()
	MutateFeedMappings(context.Context, *MutateFeedMappingsRequest) (*MutateFeedMappingsResponse, error)
}

// UnimplementedFeedMappingServiceServer can be embedded to have forward compatible implementations.
type UnimplementedFeedMappingServiceServer struct {
}

func (*UnimplementedFeedMappingServiceServer) MutateFeedMappings(context.Context, *MutateFeedMappingsRequest) (*MutateFeedMappingsResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method MutateFeedMappings not implemented")
}

func RegisterFeedMappingServiceServer(s *grpc.Server, srv FeedMappingServiceServer) {
	s.RegisterService(&_FeedMappingService_serviceDesc, srv)
}

func _FeedMappingService_MutateFeedMappings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateFeedMappingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedMappingServiceServer).MutateFeedMappings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v18.services.FeedMappingService/MutateFeedMappings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedMappingServiceServer).MutateFeedMappings(ctx, req.(*MutateFeedMappingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FeedMappingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v18.services.FeedMappingService",
	HandlerType: (*FeedMappingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateFeedMappings",
			Handler:    _FeedMappingService_MutateFeedMappings_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v18/services/feed_mapping_service.proto",
}
