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
// source: google/ads/googleads/v18/services/campaign_label_service.proto

package services

import (
	context "context"
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
// [CampaignLabelService.MutateCampaignLabels][google.ads.googleads.v18.services.CampaignLabelService.MutateCampaignLabels].
type MutateCampaignLabelsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. ID of the customer whose campaign-label relationships are being
	// modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The list of operations to perform on campaign-label
	// relationships.
	Operations []*CampaignLabelOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	// If true, successful operations will be carried out and invalid
	// operations will return errors. If false, all operations will be carried
	// out in one transaction if and only if they are all valid.
	// Default is false.
	PartialFailure bool `protobuf:"varint,3,opt,name=partial_failure,json=partialFailure,proto3" json:"partial_failure,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly bool `protobuf:"varint,4,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
}

func (x *MutateCampaignLabelsRequest) Reset() {
	*x = MutateCampaignLabelsRequest{}
	mi := &file_google_ads_googleads_v18_services_campaign_label_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MutateCampaignLabelsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateCampaignLabelsRequest) ProtoMessage() {}

func (x *MutateCampaignLabelsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v18_services_campaign_label_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateCampaignLabelsRequest.ProtoReflect.Descriptor instead.
func (*MutateCampaignLabelsRequest) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v18_services_campaign_label_service_proto_rawDescGZIP(), []int{0}
}

func (x *MutateCampaignLabelsRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *MutateCampaignLabelsRequest) GetOperations() []*CampaignLabelOperation {
	if x != nil {
		return x.Operations
	}
	return nil
}

func (x *MutateCampaignLabelsRequest) GetPartialFailure() bool {
	if x != nil {
		return x.PartialFailure
	}
	return false
}

func (x *MutateCampaignLabelsRequest) GetValidateOnly() bool {
	if x != nil {
		return x.ValidateOnly
	}
	return false
}

// A single operation (create, remove) on a campaign-label relationship.
type CampaignLabelOperation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The mutate operation.
	//
	// Types that are assignable to Operation:
	//
	//	*CampaignLabelOperation_Create
	//	*CampaignLabelOperation_Remove
	Operation isCampaignLabelOperation_Operation `protobuf_oneof:"operation"`
}

func (x *CampaignLabelOperation) Reset() {
	*x = CampaignLabelOperation{}
	mi := &file_google_ads_googleads_v18_services_campaign_label_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CampaignLabelOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CampaignLabelOperation) ProtoMessage() {}

func (x *CampaignLabelOperation) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v18_services_campaign_label_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CampaignLabelOperation.ProtoReflect.Descriptor instead.
func (*CampaignLabelOperation) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v18_services_campaign_label_service_proto_rawDescGZIP(), []int{1}
}

func (m *CampaignLabelOperation) GetOperation() isCampaignLabelOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (x *CampaignLabelOperation) GetCreate() *resources.CampaignLabel {
	if x, ok := x.GetOperation().(*CampaignLabelOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (x *CampaignLabelOperation) GetRemove() string {
	if x, ok := x.GetOperation().(*CampaignLabelOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

type isCampaignLabelOperation_Operation interface {
	isCampaignLabelOperation_Operation()
}

type CampaignLabelOperation_Create struct {
	// Create operation: No resource name is expected for the new campaign-label
	// relationship.
	Create *resources.CampaignLabel `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type CampaignLabelOperation_Remove struct {
	// Remove operation: A resource name for the campaign-label relationship
	// being removed, in this format:
	//
	// `customers/{customer_id}/campaignLabels/{campaign_id}~{label_id}`
	Remove string `protobuf:"bytes,2,opt,name=remove,proto3,oneof"`
}

func (*CampaignLabelOperation_Create) isCampaignLabelOperation_Operation() {}

func (*CampaignLabelOperation_Remove) isCampaignLabelOperation_Operation() {}

// Response message for a campaign labels mutate.
type MutateCampaignLabelsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (for example, auth
	// errors), we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results []*MutateCampaignLabelResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *MutateCampaignLabelsResponse) Reset() {
	*x = MutateCampaignLabelsResponse{}
	mi := &file_google_ads_googleads_v18_services_campaign_label_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MutateCampaignLabelsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateCampaignLabelsResponse) ProtoMessage() {}

func (x *MutateCampaignLabelsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v18_services_campaign_label_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateCampaignLabelsResponse.ProtoReflect.Descriptor instead.
func (*MutateCampaignLabelsResponse) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v18_services_campaign_label_service_proto_rawDescGZIP(), []int{2}
}

func (x *MutateCampaignLabelsResponse) GetPartialFailureError() *status.Status {
	if x != nil {
		return x.PartialFailureError
	}
	return nil
}

func (x *MutateCampaignLabelsResponse) GetResults() []*MutateCampaignLabelResult {
	if x != nil {
		return x.Results
	}
	return nil
}

// The result for a campaign label mutate.
type MutateCampaignLabelResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Returned for successful operations.
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
}

func (x *MutateCampaignLabelResult) Reset() {
	*x = MutateCampaignLabelResult{}
	mi := &file_google_ads_googleads_v18_services_campaign_label_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MutateCampaignLabelResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateCampaignLabelResult) ProtoMessage() {}

func (x *MutateCampaignLabelResult) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v18_services_campaign_label_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateCampaignLabelResult.ProtoReflect.Descriptor instead.
func (*MutateCampaignLabelResult) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v18_services_campaign_label_service_proto_rawDescGZIP(), []int{3}
}

func (x *MutateCampaignLabelResult) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

var File_google_ads_googleads_v18_services_campaign_label_service_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v18_services_campaign_label_service_proto_rawDesc = []byte{
	0x0a, 0x3e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x38, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2f, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x5f, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x38, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x1a, 0x37, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x38, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e,
	0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf1, 0x01, 0x0a, 0x1b, 0x4d, 0x75, 0x74,
	0x61, 0x74, 0x65, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0,
	0x41, 0x02, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x5e,
	0x0a, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x39, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x38, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x03, 0xe0,
	0x41, 0x02, 0x52, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x27,
	0x0a, 0x0f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c,
	0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x6e, 0x6c, 0x79, 0x22, 0xb9, 0x01, 0x0a,
	0x16, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4b, 0x0a, 0x06, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x38, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x6d,
	0x70, 0x61, 0x69, 0x67, 0x6e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x48, 0x00, 0x52, 0x06, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x12, 0x45, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x2b, 0xfa, 0x41, 0x28, 0x0a, 0x26, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x48, 0x00, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xbe, 0x01, 0x0a, 0x1c, 0x4d, 0x75, 0x74,
	0x61, 0x74, 0x65, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x15, 0x70, 0x61, 0x72,
	0x74, 0x69, 0x61, 0x6c, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x5f, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x13, 0x70, 0x61,
	0x72, 0x74, 0x69, 0x61, 0x6c, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x12, 0x56, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x38, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x43, 0x61, 0x6d,
	0x70, 0x61, 0x69, 0x67, 0x6e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x22, 0x6d, 0x0a, 0x19, 0x4d, 0x75, 0x74,
	0x61, 0x74, 0x65, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x50, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2b, 0xfa,
	0x41, 0x28, 0x0a, 0x26, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x61, 0x6d,
	0x70, 0x61, 0x69, 0x67, 0x6e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x32, 0xd1, 0x02, 0x0a, 0x14, 0x43, 0x61, 0x6d,
	0x70, 0x61, 0x69, 0x67, 0x6e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0xf1, 0x01, 0x0a, 0x14, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x43, 0x61, 0x6d, 0x70,
	0x61, 0x69, 0x67, 0x6e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x3e, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x76, 0x31, 0x38, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d,
	0x75, 0x74, 0x61, 0x74, 0x65, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x4c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3f, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x76, 0x31, 0x38, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d,
	0x75, 0x74, 0x61, 0x74, 0x65, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x4c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x58, 0xda, 0x41, 0x16,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x2c, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x39, 0x3a, 0x01, 0x2a, 0x22,
	0x34, 0x2f, 0x76, 0x31, 0x38, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f,
	0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x3d, 0x2a, 0x7d, 0x2f,
	0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x3a, 0x6d,
	0x75, 0x74, 0x61, 0x74, 0x65, 0x1a, 0x45, 0xca, 0x41, 0x18, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0xd2, 0x41, 0x27, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x64, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x42, 0x85, 0x02, 0x0a,
	0x25, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x38, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x42, 0x19, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74,
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
	file_google_ads_googleads_v18_services_campaign_label_service_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v18_services_campaign_label_service_proto_rawDescData = file_google_ads_googleads_v18_services_campaign_label_service_proto_rawDesc
)

func file_google_ads_googleads_v18_services_campaign_label_service_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v18_services_campaign_label_service_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v18_services_campaign_label_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v18_services_campaign_label_service_proto_rawDescData)
	})
	return file_google_ads_googleads_v18_services_campaign_label_service_proto_rawDescData
}

var file_google_ads_googleads_v18_services_campaign_label_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_ads_googleads_v18_services_campaign_label_service_proto_goTypes = []any{
	(*MutateCampaignLabelsRequest)(nil),  // 0: google.ads.googleads.v18.services.MutateCampaignLabelsRequest
	(*CampaignLabelOperation)(nil),       // 1: google.ads.googleads.v18.services.CampaignLabelOperation
	(*MutateCampaignLabelsResponse)(nil), // 2: google.ads.googleads.v18.services.MutateCampaignLabelsResponse
	(*MutateCampaignLabelResult)(nil),    // 3: google.ads.googleads.v18.services.MutateCampaignLabelResult
	(*resources.CampaignLabel)(nil),      // 4: google.ads.googleads.v18.resources.CampaignLabel
	(*status.Status)(nil),                // 5: google.rpc.Status
}
var file_google_ads_googleads_v18_services_campaign_label_service_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v18.services.MutateCampaignLabelsRequest.operations:type_name -> google.ads.googleads.v18.services.CampaignLabelOperation
	4, // 1: google.ads.googleads.v18.services.CampaignLabelOperation.create:type_name -> google.ads.googleads.v18.resources.CampaignLabel
	5, // 2: google.ads.googleads.v18.services.MutateCampaignLabelsResponse.partial_failure_error:type_name -> google.rpc.Status
	3, // 3: google.ads.googleads.v18.services.MutateCampaignLabelsResponse.results:type_name -> google.ads.googleads.v18.services.MutateCampaignLabelResult
	0, // 4: google.ads.googleads.v18.services.CampaignLabelService.MutateCampaignLabels:input_type -> google.ads.googleads.v18.services.MutateCampaignLabelsRequest
	2, // 5: google.ads.googleads.v18.services.CampaignLabelService.MutateCampaignLabels:output_type -> google.ads.googleads.v18.services.MutateCampaignLabelsResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v18_services_campaign_label_service_proto_init() }
func file_google_ads_googleads_v18_services_campaign_label_service_proto_init() {
	if File_google_ads_googleads_v18_services_campaign_label_service_proto != nil {
		return
	}
	file_google_ads_googleads_v18_services_campaign_label_service_proto_msgTypes[1].OneofWrappers = []any{
		(*CampaignLabelOperation_Create)(nil),
		(*CampaignLabelOperation_Remove)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v18_services_campaign_label_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_ads_googleads_v18_services_campaign_label_service_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v18_services_campaign_label_service_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v18_services_campaign_label_service_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v18_services_campaign_label_service_proto = out.File
	file_google_ads_googleads_v18_services_campaign_label_service_proto_rawDesc = nil
	file_google_ads_googleads_v18_services_campaign_label_service_proto_goTypes = nil
	file_google_ads_googleads_v18_services_campaign_label_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CampaignLabelServiceClient is the client API for CampaignLabelService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CampaignLabelServiceClient interface {
	// Creates and removes campaign-label relationships.
	// Operation statuses are returned.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[DatabaseError]()
	//	[FieldError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[LabelError]()
	//	[MutateError]()
	//	[NewResourceCreationError]()
	//	[QuotaError]()
	//	[RequestError]()
	MutateCampaignLabels(ctx context.Context, in *MutateCampaignLabelsRequest, opts ...grpc.CallOption) (*MutateCampaignLabelsResponse, error)
}

type campaignLabelServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCampaignLabelServiceClient(cc grpc.ClientConnInterface) CampaignLabelServiceClient {
	return &campaignLabelServiceClient{cc}
}

func (c *campaignLabelServiceClient) MutateCampaignLabels(ctx context.Context, in *MutateCampaignLabelsRequest, opts ...grpc.CallOption) (*MutateCampaignLabelsResponse, error) {
	out := new(MutateCampaignLabelsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v18.services.CampaignLabelService/MutateCampaignLabels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CampaignLabelServiceServer is the server API for CampaignLabelService service.
type CampaignLabelServiceServer interface {
	// Creates and removes campaign-label relationships.
	// Operation statuses are returned.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[DatabaseError]()
	//	[FieldError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[LabelError]()
	//	[MutateError]()
	//	[NewResourceCreationError]()
	//	[QuotaError]()
	//	[RequestError]()
	MutateCampaignLabels(context.Context, *MutateCampaignLabelsRequest) (*MutateCampaignLabelsResponse, error)
}

// UnimplementedCampaignLabelServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCampaignLabelServiceServer struct {
}

func (*UnimplementedCampaignLabelServiceServer) MutateCampaignLabels(context.Context, *MutateCampaignLabelsRequest) (*MutateCampaignLabelsResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method MutateCampaignLabels not implemented")
}

func RegisterCampaignLabelServiceServer(s *grpc.Server, srv CampaignLabelServiceServer) {
	s.RegisterService(&_CampaignLabelService_serviceDesc, srv)
}

func _CampaignLabelService_MutateCampaignLabels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateCampaignLabelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignLabelServiceServer).MutateCampaignLabels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v18.services.CampaignLabelService/MutateCampaignLabels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignLabelServiceServer).MutateCampaignLabels(ctx, req.(*MutateCampaignLabelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CampaignLabelService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v18.services.CampaignLabelService",
	HandlerType: (*CampaignLabelServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateCampaignLabels",
			Handler:    _CampaignLabelService_MutateCampaignLabels_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v18/services/campaign_label_service.proto",
}
