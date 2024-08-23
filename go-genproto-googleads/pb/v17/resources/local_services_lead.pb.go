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
// source: google/ads/googleads/v17/resources/local_services_lead.proto

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

// Data from Local Services Lead.
// Contains details of Lead which is generated when user calls, messages or
// books service from advertiser.
// More info: https://ads.google.com/local-services-ads
type LocalServicesLead struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The resource name of the local services lead data.
	// Local Services Lead resource name have the form
	//
	// `customers/{customer_id}/localServicesLead/{local_services_lead_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. ID of this Lead.
	Id int64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	// Output only. Service category of the lead. For example:
	// `xcat:service_area_business_hvac`,
	// `xcat:service_area_business_real_estate_agent`, etc.
	// For more details see:
	// https://developers.google.com/google-ads/api/data/codes-formats#local_services_ids
	CategoryId string `protobuf:"bytes,3,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	// Output only. Service for the  category. For example: `buyer_agent`,
	// `seller_agent` for the category of
	// `xcat:service_area_business_real_estate_agent`.
	ServiceId string `protobuf:"bytes,4,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	// Output only. Lead's contact details.
	ContactDetails *ContactDetails `protobuf:"bytes,5,opt,name=contact_details,json=contactDetails,proto3" json:"contact_details,omitempty"`
	// Output only. Type of Local Services lead: phone, message, booking, etc.
	LeadType enums.LocalServicesLeadTypeEnum_LeadType `protobuf:"varint,6,opt,name=lead_type,json=leadType,proto3,enum=google.ads.googleads.v17.enums.LocalServicesLeadTypeEnum_LeadType" json:"lead_type,omitempty"`
	// Output only. Current status of lead.
	LeadStatus enums.LocalServicesLeadStatusEnum_LeadStatus `protobuf:"varint,7,opt,name=lead_status,json=leadStatus,proto3,enum=google.ads.googleads.v17.enums.LocalServicesLeadStatusEnum_LeadStatus" json:"lead_status,omitempty"`
	// Output only. The date time at which lead was created by Local Services Ads.
	// The format is "YYYY-MM-DD HH:MM:SS" in the Google Ads account's timezone.
	// Examples: "2018-03-05 09:15:00" or "2018-02-01 14:34:30"
	CreationDateTime string `protobuf:"bytes,8,opt,name=creation_date_time,json=creationDateTime,proto3" json:"creation_date_time,omitempty"`
	// Output only. Language used by the Local Services provider linked to lead.
	// See https://developers.google.com/google-ads/api/data/codes-formats#locales
	Locale string `protobuf:"bytes,9,opt,name=locale,proto3" json:"locale,omitempty"`
	// Output only. Note added by advertiser for the lead.
	Note *Note `protobuf:"bytes,10,opt,name=note,proto3,oneof" json:"note,omitempty"`
	// Output only. True if the advertiser was charged for the lead.
	LeadCharged bool `protobuf:"varint,11,opt,name=lead_charged,json=leadCharged,proto3" json:"lead_charged,omitempty"`
	// Output only. Credit details of the lead.
	CreditDetails *CreditDetails `protobuf:"bytes,12,opt,name=credit_details,json=creditDetails,proto3,oneof" json:"credit_details,omitempty"`
}

func (x *LocalServicesLead) Reset() {
	*x = LocalServicesLead{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v17_resources_local_services_lead_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocalServicesLead) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocalServicesLead) ProtoMessage() {}

func (x *LocalServicesLead) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v17_resources_local_services_lead_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocalServicesLead.ProtoReflect.Descriptor instead.
func (*LocalServicesLead) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v17_resources_local_services_lead_proto_rawDescGZIP(), []int{0}
}

func (x *LocalServicesLead) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *LocalServicesLead) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *LocalServicesLead) GetCategoryId() string {
	if x != nil {
		return x.CategoryId
	}
	return ""
}

func (x *LocalServicesLead) GetServiceId() string {
	if x != nil {
		return x.ServiceId
	}
	return ""
}

func (x *LocalServicesLead) GetContactDetails() *ContactDetails {
	if x != nil {
		return x.ContactDetails
	}
	return nil
}

func (x *LocalServicesLead) GetLeadType() enums.LocalServicesLeadTypeEnum_LeadType {
	if x != nil {
		return x.LeadType
	}
	return enums.LocalServicesLeadTypeEnum_LeadType(0)
}

func (x *LocalServicesLead) GetLeadStatus() enums.LocalServicesLeadStatusEnum_LeadStatus {
	if x != nil {
		return x.LeadStatus
	}
	return enums.LocalServicesLeadStatusEnum_LeadStatus(0)
}

func (x *LocalServicesLead) GetCreationDateTime() string {
	if x != nil {
		return x.CreationDateTime
	}
	return ""
}

func (x *LocalServicesLead) GetLocale() string {
	if x != nil {
		return x.Locale
	}
	return ""
}

func (x *LocalServicesLead) GetNote() *Note {
	if x != nil {
		return x.Note
	}
	return nil
}

func (x *LocalServicesLead) GetLeadCharged() bool {
	if x != nil {
		return x.LeadCharged
	}
	return false
}

func (x *LocalServicesLead) GetCreditDetails() *CreditDetails {
	if x != nil {
		return x.CreditDetails
	}
	return nil
}

// Fields containing consumer contact details.
type ContactDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. Consumer phone number in E164 format.
	PhoneNumber string `protobuf:"bytes,1,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	// Output only. Consumer email address.
	Email string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	// Output only. Consumer name if consumer provided name from Message or
	// Booking form on google.com
	ConsumerName string `protobuf:"bytes,3,opt,name=consumer_name,json=consumerName,proto3" json:"consumer_name,omitempty"`
}

func (x *ContactDetails) Reset() {
	*x = ContactDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v17_resources_local_services_lead_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContactDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContactDetails) ProtoMessage() {}

func (x *ContactDetails) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v17_resources_local_services_lead_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContactDetails.ProtoReflect.Descriptor instead.
func (*ContactDetails) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v17_resources_local_services_lead_proto_rawDescGZIP(), []int{1}
}

func (x *ContactDetails) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *ContactDetails) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *ContactDetails) GetConsumerName() string {
	if x != nil {
		return x.ConsumerName
	}
	return ""
}

// Represents a note added to a lead by the advertiser. Advertisers can edit
// notes, which will reset edit time and change description.
type Note struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The date time when lead note was edited. The format is
	// "YYYY-MM-DD HH:MM:SS" in the Google Ads account's timezone. Examples:
	// "2018-03-05 09:15:00" or "2018-02-01 14:34:30"
	EditDateTime string `protobuf:"bytes,1,opt,name=edit_date_time,json=editDateTime,proto3" json:"edit_date_time,omitempty"`
	// Output only. Content of lead note.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *Note) Reset() {
	*x = Note{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v17_resources_local_services_lead_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Note) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Note) ProtoMessage() {}

func (x *Note) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v17_resources_local_services_lead_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Note.ProtoReflect.Descriptor instead.
func (*Note) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v17_resources_local_services_lead_proto_rawDescGZIP(), []int{2}
}

func (x *Note) GetEditDateTime() string {
	if x != nil {
		return x.EditDateTime
	}
	return ""
}

func (x *Note) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

// Represents the credit details of a lead.
type CreditDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. Credit state of the lead.
	CreditState enums.LocalServicesCreditStateEnum_CreditState `protobuf:"varint,1,opt,name=credit_state,json=creditState,proto3,enum=google.ads.googleads.v17.enums.LocalServicesCreditStateEnum_CreditState" json:"credit_state,omitempty"`
	// Output only. The date time when the credit state of the lead was last
	// updated. The format is "YYYY-MM-DD HH:MM:SS" in the Google Ads account's
	// timezone. Examples: "2018-03-05 09:15:00" or "2018-02-01 14:34:30"
	CreditStateLastUpdateDateTime string `protobuf:"bytes,2,opt,name=credit_state_last_update_date_time,json=creditStateLastUpdateDateTime,proto3" json:"credit_state_last_update_date_time,omitempty"`
}

func (x *CreditDetails) Reset() {
	*x = CreditDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v17_resources_local_services_lead_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreditDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreditDetails) ProtoMessage() {}

func (x *CreditDetails) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v17_resources_local_services_lead_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreditDetails.ProtoReflect.Descriptor instead.
func (*CreditDetails) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v17_resources_local_services_lead_proto_rawDescGZIP(), []int{3}
}

func (x *CreditDetails) GetCreditState() enums.LocalServicesCreditStateEnum_CreditState {
	if x != nil {
		return x.CreditState
	}
	return enums.LocalServicesCreditStateEnum_CreditState(0)
}

func (x *CreditDetails) GetCreditStateLastUpdateDateTime() string {
	if x != nil {
		return x.CreditStateLastUpdateDateTime
	}
	return ""
}

var File_google_ads_googleads_v17_resources_local_services_lead_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v17_resources_local_services_lead_proto_rawDesc = []byte{
	0x0a, 0x3c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x37, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x5f, 0x6c, 0x65, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x1a, 0x45, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x37, 0x2f, 0x65, 0x6e, 0x75,
	0x6d, 0x73, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x5f, 0x6c, 0x65, 0x61, 0x64, 0x5f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f,
	0x76, 0x31, 0x37, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x5f, 0x6c, 0x65, 0x61, 0x64, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3d, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2f, 0x76, 0x31, 0x37, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x6c,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x5f, 0x6c, 0x65, 0x61, 0x64, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61,
	0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb7, 0x07, 0x0a, 0x11, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x4c, 0x65, 0x61, 0x64, 0x12, 0x57, 0x0a, 0x0d, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x32, 0xe0, 0x41, 0x03, 0xfa, 0x41, 0x2c, 0x0a, 0x2a, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x4c, 0x65, 0x61, 0x64, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x24, 0x0a, 0x0b, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03,
	0xe0, 0x41, 0x03, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12,
	0x22, 0x0a, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x09, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x49, 0x64, 0x12, 0x60, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x5f, 0x64,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x64, 0x0a, 0x09, 0x6c, 0x65, 0x61, 0x64, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x42, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x37, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x4c, 0x65, 0x61, 0x64, 0x54, 0x79, 0x70, 0x65, 0x45,
	0x6e, 0x75, 0x6d, 0x2e, 0x4c, 0x65, 0x61, 0x64, 0x54, 0x79, 0x70, 0x65, 0x42, 0x03, 0xe0, 0x41,
	0x03, 0x52, 0x08, 0x6c, 0x65, 0x61, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x6c, 0x0a, 0x0b, 0x6c,
	0x65, 0x61, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x46, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x4c,
	0x65, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x4c, 0x65,
	0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x6c,
	0x65, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x31, 0x0a, 0x12, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x10, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x06,
	0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41,
	0x03, 0x52, 0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x12, 0x46, 0x0a, 0x04, 0x6e, 0x6f, 0x74,
	0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x37, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x4e, 0x6f, 0x74,
	0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x26, 0x0a, 0x0c, 0x6c, 0x65, 0x61, 0x64, 0x5f, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65,
	0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0b, 0x6c, 0x65,
	0x61, 0x64, 0x43, 0x68, 0x61, 0x72, 0x67, 0x65, 0x64, 0x12, 0x62, 0x0a, 0x0e, 0x63, 0x72, 0x65,
	0x64, 0x69, 0x74, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x31, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x01, 0x52, 0x0d, 0x63, 0x72, 0x65,
	0x64, 0x69, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x88, 0x01, 0x01, 0x3a, 0x74, 0xea,
	0x41, 0x71, 0x0a, 0x2a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4c, 0x6f, 0x63,
	0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x4c, 0x65, 0x61, 0x64, 0x12, 0x43,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x4c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61,
	0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x5f, 0x6c, 0x65, 0x61, 0x64, 0x5f,
	0x69, 0x64, 0x7d, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x6f, 0x74, 0x65, 0x42, 0x11, 0x0a, 0x0f,
	0x5f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22,
	0x7d, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x12, 0x26, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0b, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x28, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03,
	0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x58,
	0x0a, 0x04, 0x4e, 0x6f, 0x74, 0x65, 0x12, 0x29, 0x0a, 0x0e, 0x65, 0x64, 0x69, 0x74, 0x5f, 0x64,
	0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03,
	0xe0, 0x41, 0x03, 0x52, 0x0c, 0x65, 0x64, 0x69, 0x74, 0x44, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x25, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xd1, 0x01, 0x0a, 0x0d, 0x43, 0x72, 0x65,
	0x64, 0x69, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x70, 0x0a, 0x0c, 0x63, 0x72,
	0x65, 0x64, 0x69, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x48, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x43,
	0x72, 0x65, 0x64, 0x69, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x43,
	0x72, 0x65, 0x64, 0x69, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52,
	0x0b, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x4e, 0x0a, 0x22,
	0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x6c, 0x61, 0x73,
	0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x1d, 0x63,
	0x72, 0x65, 0x64, 0x69, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4c, 0x61, 0x73, 0x74, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x88, 0x02, 0x0a,
	0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0x16, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x4c, 0x65, 0x61, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x4b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67,
	0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x37, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x3b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xa2, 0x02,
	0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64,
	0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x37, 0x2e,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xca, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73,
	0x5c, 0x56, 0x31, 0x37, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xea, 0x02,
	0x26, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x37, 0x3a, 0x3a, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v17_resources_local_services_lead_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v17_resources_local_services_lead_proto_rawDescData = file_google_ads_googleads_v17_resources_local_services_lead_proto_rawDesc
)

func file_google_ads_googleads_v17_resources_local_services_lead_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v17_resources_local_services_lead_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v17_resources_local_services_lead_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v17_resources_local_services_lead_proto_rawDescData)
	})
	return file_google_ads_googleads_v17_resources_local_services_lead_proto_rawDescData
}

var file_google_ads_googleads_v17_resources_local_services_lead_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_ads_googleads_v17_resources_local_services_lead_proto_goTypes = []any{
	(*LocalServicesLead)(nil),                           // 0: google.ads.googleads.v17.resources.LocalServicesLead
	(*ContactDetails)(nil),                              // 1: google.ads.googleads.v17.resources.ContactDetails
	(*Note)(nil),                                        // 2: google.ads.googleads.v17.resources.Note
	(*CreditDetails)(nil),                               // 3: google.ads.googleads.v17.resources.CreditDetails
	(enums.LocalServicesLeadTypeEnum_LeadType)(0),       // 4: google.ads.googleads.v17.enums.LocalServicesLeadTypeEnum.LeadType
	(enums.LocalServicesLeadStatusEnum_LeadStatus)(0),   // 5: google.ads.googleads.v17.enums.LocalServicesLeadStatusEnum.LeadStatus
	(enums.LocalServicesCreditStateEnum_CreditState)(0), // 6: google.ads.googleads.v17.enums.LocalServicesCreditStateEnum.CreditState
}
var file_google_ads_googleads_v17_resources_local_services_lead_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v17.resources.LocalServicesLead.contact_details:type_name -> google.ads.googleads.v17.resources.ContactDetails
	4, // 1: google.ads.googleads.v17.resources.LocalServicesLead.lead_type:type_name -> google.ads.googleads.v17.enums.LocalServicesLeadTypeEnum.LeadType
	5, // 2: google.ads.googleads.v17.resources.LocalServicesLead.lead_status:type_name -> google.ads.googleads.v17.enums.LocalServicesLeadStatusEnum.LeadStatus
	2, // 3: google.ads.googleads.v17.resources.LocalServicesLead.note:type_name -> google.ads.googleads.v17.resources.Note
	3, // 4: google.ads.googleads.v17.resources.LocalServicesLead.credit_details:type_name -> google.ads.googleads.v17.resources.CreditDetails
	6, // 5: google.ads.googleads.v17.resources.CreditDetails.credit_state:type_name -> google.ads.googleads.v17.enums.LocalServicesCreditStateEnum.CreditState
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v17_resources_local_services_lead_proto_init() }
func file_google_ads_googleads_v17_resources_local_services_lead_proto_init() {
	if File_google_ads_googleads_v17_resources_local_services_lead_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v17_resources_local_services_lead_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*LocalServicesLead); i {
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
		file_google_ads_googleads_v17_resources_local_services_lead_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ContactDetails); i {
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
		file_google_ads_googleads_v17_resources_local_services_lead_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Note); i {
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
		file_google_ads_googleads_v17_resources_local_services_lead_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*CreditDetails); i {
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
	file_google_ads_googleads_v17_resources_local_services_lead_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v17_resources_local_services_lead_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v17_resources_local_services_lead_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v17_resources_local_services_lead_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v17_resources_local_services_lead_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v17_resources_local_services_lead_proto = out.File
	file_google_ads_googleads_v17_resources_local_services_lead_proto_rawDesc = nil
	file_google_ads_googleads_v17_resources_local_services_lead_proto_goTypes = nil
	file_google_ads_googleads_v17_resources_local_services_lead_proto_depIdxs = nil
}
