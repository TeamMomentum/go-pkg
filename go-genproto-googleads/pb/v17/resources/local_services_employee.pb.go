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
// source: google/ads/googleads/v17/resources/local_services_employee.proto

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

// A local services employee resource.
type LocalServicesEmployee struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Immutable. The resource name of the Local Services Verification.
	// Local Services Verification resource names have the form:
	//
	// `customers/{customer_id}/localServicesEmployees/{gls_employee_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The ID of the employee.
	Id *int64 `protobuf:"varint,2,opt,name=id,proto3,oneof" json:"id,omitempty"`
	// Output only. Timestamp of employee creation.
	// The format is "YYYY-MM-DD HH:MM:SS" in the Google Ads account's timezone.
	// Examples: "2018-03-05 09:15:00" or "2018-02-01 14:34:30"
	CreationDateTime string `protobuf:"bytes,3,opt,name=creation_date_time,json=creationDateTime,proto3" json:"creation_date_time,omitempty"`
	// Output only. Employee status, such as DELETED or ENABLED.
	Status enums.LocalServicesEmployeeStatusEnum_LocalServicesEmployeeStatus `protobuf:"varint,4,opt,name=status,proto3,enum=google.ads.googleads.v17.enums.LocalServicesEmployeeStatusEnum_LocalServicesEmployeeStatus" json:"status,omitempty"`
	// Output only. Employee type.
	Type enums.LocalServicesEmployeeTypeEnum_LocalServicesEmployeeType `protobuf:"varint,5,opt,name=type,proto3,enum=google.ads.googleads.v17.enums.LocalServicesEmployeeTypeEnum_LocalServicesEmployeeType" json:"type,omitempty"`
	// Output only. A list of degrees this employee has obtained, and wants to
	// feature.
	UniversityDegrees []*UniversityDegree `protobuf:"bytes,6,rep,name=university_degrees,json=universityDegrees,proto3" json:"university_degrees,omitempty"`
	// Output only. The institutions where the employee has completed their
	// residency.
	Residencies []*Residency `protobuf:"bytes,7,rep,name=residencies,proto3" json:"residencies,omitempty"`
	// Output only. The institutions where the employee has completed their
	// fellowship.
	Fellowships []*Fellowship `protobuf:"bytes,8,rep,name=fellowships,proto3" json:"fellowships,omitempty"`
	// Output only. Job title for this employee, such as "Senior partner" in legal
	// verticals.
	JobTitle *string `protobuf:"bytes,9,opt,name=job_title,json=jobTitle,proto3,oneof" json:"job_title,omitempty"`
	// Output only. The year that this employee started practicing in this field.
	YearStartedPracticing *int32 `protobuf:"varint,10,opt,name=year_started_practicing,json=yearStartedPracticing,proto3,oneof" json:"year_started_practicing,omitempty"`
	// Output only. Languages that the employee speaks, represented as language
	// tags from https://developers.google.com/admin-sdk/directory/v1/languages
	LanguagesSpoken []string `protobuf:"bytes,11,rep,name=languages_spoken,json=languagesSpoken,proto3" json:"languages_spoken,omitempty"`
	// Output only. Category of the employee. A list of Local Services category
	// IDs can be found at
	// https://developers.google.com/google-ads/api/data/codes-formats#local_services_ids.
	CategoryIds []string `protobuf:"bytes,12,rep,name=category_ids,json=categoryIds,proto3" json:"category_ids,omitempty"`
	// Output only. NPI id associated with the employee.
	NationalProviderIdNumber *string `protobuf:"bytes,13,opt,name=national_provider_id_number,json=nationalProviderIdNumber,proto3,oneof" json:"national_provider_id_number,omitempty"`
	// Output only. Email address of the employee.
	EmailAddress *string `protobuf:"bytes,14,opt,name=email_address,json=emailAddress,proto3,oneof" json:"email_address,omitempty"`
	// Output only. First name of the employee.
	FirstName *string `protobuf:"bytes,15,opt,name=first_name,json=firstName,proto3,oneof" json:"first_name,omitempty"`
	// Output only. Middle name of the employee.
	MiddleName *string `protobuf:"bytes,16,opt,name=middle_name,json=middleName,proto3,oneof" json:"middle_name,omitempty"`
	// Output only. Last name of the employee.
	LastName *string `protobuf:"bytes,17,opt,name=last_name,json=lastName,proto3,oneof" json:"last_name,omitempty"`
}

func (x *LocalServicesEmployee) Reset() {
	*x = LocalServicesEmployee{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v17_resources_local_services_employee_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocalServicesEmployee) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocalServicesEmployee) ProtoMessage() {}

func (x *LocalServicesEmployee) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v17_resources_local_services_employee_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocalServicesEmployee.ProtoReflect.Descriptor instead.
func (*LocalServicesEmployee) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v17_resources_local_services_employee_proto_rawDescGZIP(), []int{0}
}

func (x *LocalServicesEmployee) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *LocalServicesEmployee) GetId() int64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *LocalServicesEmployee) GetCreationDateTime() string {
	if x != nil {
		return x.CreationDateTime
	}
	return ""
}

func (x *LocalServicesEmployee) GetStatus() enums.LocalServicesEmployeeStatusEnum_LocalServicesEmployeeStatus {
	if x != nil {
		return x.Status
	}
	return enums.LocalServicesEmployeeStatusEnum_LocalServicesEmployeeStatus(0)
}

func (x *LocalServicesEmployee) GetType() enums.LocalServicesEmployeeTypeEnum_LocalServicesEmployeeType {
	if x != nil {
		return x.Type
	}
	return enums.LocalServicesEmployeeTypeEnum_LocalServicesEmployeeType(0)
}

func (x *LocalServicesEmployee) GetUniversityDegrees() []*UniversityDegree {
	if x != nil {
		return x.UniversityDegrees
	}
	return nil
}

func (x *LocalServicesEmployee) GetResidencies() []*Residency {
	if x != nil {
		return x.Residencies
	}
	return nil
}

func (x *LocalServicesEmployee) GetFellowships() []*Fellowship {
	if x != nil {
		return x.Fellowships
	}
	return nil
}

func (x *LocalServicesEmployee) GetJobTitle() string {
	if x != nil && x.JobTitle != nil {
		return *x.JobTitle
	}
	return ""
}

func (x *LocalServicesEmployee) GetYearStartedPracticing() int32 {
	if x != nil && x.YearStartedPracticing != nil {
		return *x.YearStartedPracticing
	}
	return 0
}

func (x *LocalServicesEmployee) GetLanguagesSpoken() []string {
	if x != nil {
		return x.LanguagesSpoken
	}
	return nil
}

func (x *LocalServicesEmployee) GetCategoryIds() []string {
	if x != nil {
		return x.CategoryIds
	}
	return nil
}

func (x *LocalServicesEmployee) GetNationalProviderIdNumber() string {
	if x != nil && x.NationalProviderIdNumber != nil {
		return *x.NationalProviderIdNumber
	}
	return ""
}

func (x *LocalServicesEmployee) GetEmailAddress() string {
	if x != nil && x.EmailAddress != nil {
		return *x.EmailAddress
	}
	return ""
}

func (x *LocalServicesEmployee) GetFirstName() string {
	if x != nil && x.FirstName != nil {
		return *x.FirstName
	}
	return ""
}

func (x *LocalServicesEmployee) GetMiddleName() string {
	if x != nil && x.MiddleName != nil {
		return *x.MiddleName
	}
	return ""
}

func (x *LocalServicesEmployee) GetLastName() string {
	if x != nil && x.LastName != nil {
		return *x.LastName
	}
	return ""
}

// A list of degrees this employee has obtained, and wants to feature.
type UniversityDegree struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. Name of the university at which the degree was obtained.
	InstitutionName *string `protobuf:"bytes,1,opt,name=institution_name,json=institutionName,proto3,oneof" json:"institution_name,omitempty"`
	// Output only. Name of the degree obtained.
	Degree *string `protobuf:"bytes,2,opt,name=degree,proto3,oneof" json:"degree,omitempty"`
	// Output only. Year of graduation.
	GraduationYear *int32 `protobuf:"varint,3,opt,name=graduation_year,json=graduationYear,proto3,oneof" json:"graduation_year,omitempty"`
}

func (x *UniversityDegree) Reset() {
	*x = UniversityDegree{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v17_resources_local_services_employee_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UniversityDegree) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UniversityDegree) ProtoMessage() {}

func (x *UniversityDegree) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v17_resources_local_services_employee_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UniversityDegree.ProtoReflect.Descriptor instead.
func (*UniversityDegree) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v17_resources_local_services_employee_proto_rawDescGZIP(), []int{1}
}

func (x *UniversityDegree) GetInstitutionName() string {
	if x != nil && x.InstitutionName != nil {
		return *x.InstitutionName
	}
	return ""
}

func (x *UniversityDegree) GetDegree() string {
	if x != nil && x.Degree != nil {
		return *x.Degree
	}
	return ""
}

func (x *UniversityDegree) GetGraduationYear() int32 {
	if x != nil && x.GraduationYear != nil {
		return *x.GraduationYear
	}
	return 0
}

// Details about the employee's medical residency.
// Residency is a stage of graduate medical education in which a qualified
// medical professional practices under the supervision of a senior clinician.
type Residency struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. Name of the institution at which the residency was completed.
	InstitutionName *string `protobuf:"bytes,1,opt,name=institution_name,json=institutionName,proto3,oneof" json:"institution_name,omitempty"`
	// Output only. Year of completion.
	CompletionYear *int32 `protobuf:"varint,2,opt,name=completion_year,json=completionYear,proto3,oneof" json:"completion_year,omitempty"`
}

func (x *Residency) Reset() {
	*x = Residency{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v17_resources_local_services_employee_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Residency) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Residency) ProtoMessage() {}

func (x *Residency) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v17_resources_local_services_employee_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Residency.ProtoReflect.Descriptor instead.
func (*Residency) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v17_resources_local_services_employee_proto_rawDescGZIP(), []int{2}
}

func (x *Residency) GetInstitutionName() string {
	if x != nil && x.InstitutionName != nil {
		return *x.InstitutionName
	}
	return ""
}

func (x *Residency) GetCompletionYear() int32 {
	if x != nil && x.CompletionYear != nil {
		return *x.CompletionYear
	}
	return 0
}

// Details about the employee's medical Fellowship.
// Fellowship is a period of medical training that the professional undertakes
// after finishing their residency.
type Fellowship struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. Name of the instutition at which the fellowship was completed.
	InstitutionName *string `protobuf:"bytes,1,opt,name=institution_name,json=institutionName,proto3,oneof" json:"institution_name,omitempty"`
	// Output only. Year of completion.
	CompletionYear *int32 `protobuf:"varint,2,opt,name=completion_year,json=completionYear,proto3,oneof" json:"completion_year,omitempty"`
}

func (x *Fellowship) Reset() {
	*x = Fellowship{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v17_resources_local_services_employee_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Fellowship) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Fellowship) ProtoMessage() {}

func (x *Fellowship) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v17_resources_local_services_employee_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Fellowship.ProtoReflect.Descriptor instead.
func (*Fellowship) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v17_resources_local_services_employee_proto_rawDescGZIP(), []int{3}
}

func (x *Fellowship) GetInstitutionName() string {
	if x != nil && x.InstitutionName != nil {
		return *x.InstitutionName
	}
	return ""
}

func (x *Fellowship) GetCompletionYear() int32 {
	if x != nil && x.CompletionYear != nil {
		return *x.CompletionYear
	}
	return 0
}

var File_google_ads_googleads_v17_resources_local_services_employee_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v17_resources_local_services_employee_proto_rawDesc = []byte{
	0x0a, 0x40, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x37, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x5f, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x37,
	0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x5f, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x41, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2f, 0x76, 0x31, 0x37, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x6c, 0x6f, 0x63, 0x61,
	0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x5f, 0x65, 0x6d, 0x70, 0x6c, 0x6f,
	0x79, 0x65, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xff, 0x0a, 0x0a, 0x15, 0x4c,
	0x6f, 0x63, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x45, 0x6d, 0x70, 0x6c,
	0x6f, 0x79, 0x65, 0x65, 0x12, 0x5b, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x36, 0xe0, 0x41, 0x05,
	0xfa, 0x41, 0x30, 0x0a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4c, 0x6f,
	0x63, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x45, 0x6d, 0x70, 0x6c, 0x6f,
	0x79, 0x65, 0x65, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x18, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0,
	0x41, 0x03, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x31, 0x0a, 0x12, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x10, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x78,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x5b,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e,
	0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x45, 0x6d, 0x70,
	0x6c, 0x6f, 0x79, 0x65, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x2e,
	0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x45, 0x6d, 0x70,
	0x6c, 0x6f, 0x79, 0x65, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x03, 0xe0, 0x41, 0x03,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x70, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x57, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31,
	0x37, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x54, 0x79, 0x70, 0x65, 0x42,
	0x03, 0xe0, 0x41, 0x03, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x68, 0x0a, 0x12, 0x75, 0x6e,
	0x69, 0x76, 0x65, 0x72, 0x73, 0x69, 0x74, 0x79, 0x5f, 0x64, 0x65, 0x67, 0x72, 0x65, 0x65, 0x73,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31,
	0x37, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x55, 0x6e, 0x69, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x74, 0x79, 0x44, 0x65, 0x67, 0x72, 0x65, 0x65, 0x42, 0x03, 0xe0, 0x41,
	0x03, 0x52, 0x11, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x69, 0x74, 0x79, 0x44, 0x65, 0x67,
	0x72, 0x65, 0x65, 0x73, 0x12, 0x54, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x63,
	0x69, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x31, 0x37, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x52,
	0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0b, 0x72,
	0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x12, 0x55, 0x0a, 0x0b, 0x66, 0x65,
	0x6c, 0x6c, 0x6f, 0x77, 0x73, 0x68, 0x69, 0x70, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2e, 0x46, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x73, 0x68, 0x69, 0x70, 0x42,
	0x03, 0xe0, 0x41, 0x03, 0x52, 0x0b, 0x66, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x73, 0x68, 0x69, 0x70,
	0x73, 0x12, 0x25, 0x0a, 0x09, 0x6a, 0x6f, 0x62, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x01, 0x52, 0x08, 0x6a, 0x6f, 0x62,
	0x54, 0x69, 0x74, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x12, 0x40, 0x0a, 0x17, 0x79, 0x65, 0x61, 0x72,
	0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x70, 0x72, 0x61, 0x63, 0x74, 0x69, 0x63,
	0x69, 0x6e, 0x67, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x02,
	0x52, 0x15, 0x79, 0x65, 0x61, 0x72, 0x53, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x50, 0x72, 0x61,
	0x63, 0x74, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x88, 0x01, 0x01, 0x12, 0x2e, 0x0a, 0x10, 0x6c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x73, 0x5f, 0x73, 0x70, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x0b,
	0x20, 0x03, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0f, 0x6c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x73, 0x53, 0x70, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x26, 0x0a, 0x0c, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x09,
	0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0b, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49,
	0x64, 0x73, 0x12, 0x47, 0x0a, 0x1b, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x03, 0x52, 0x18,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x49, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x2d, 0x0a, 0x0d, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x04, 0x52, 0x0c, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x88, 0x01, 0x01, 0x12, 0x27, 0x0a, 0x0a, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03,
	0xe0, 0x41, 0x03, 0x48, 0x05, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x88, 0x01, 0x01, 0x12, 0x29, 0x0a, 0x0b, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x06, 0x52,
	0x0a, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x25,
	0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x07, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x88, 0x01, 0x01, 0x3a, 0x75, 0xea, 0x41, 0x72, 0x0a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x12, 0x40, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x73, 0x2f, 0x7b, 0x67, 0x6c, 0x73, 0x5f,
	0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x05, 0x0a, 0x03,
	0x5f, 0x69, 0x64, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x6a, 0x6f, 0x62, 0x5f, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x42, 0x1a, 0x0a, 0x18, 0x5f, 0x79, 0x65, 0x61, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x65, 0x64, 0x5f, 0x70, 0x72, 0x61, 0x63, 0x74, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x42, 0x1e, 0x0a,
	0x1c, 0x5f, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x10, 0x0a,
	0x0e, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42,
	0x0d, 0x0a, 0x0b, 0x5f, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0e,
	0x0a, 0x0c, 0x5f, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0c,
	0x0a, 0x0a, 0x5f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xd0, 0x01, 0x0a,
	0x10, 0x55, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x69, 0x74, 0x79, 0x44, 0x65, 0x67, 0x72, 0x65,
	0x65, 0x12, 0x33, 0x0a, 0x10, 0x69, 0x6e, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03,
	0x48, 0x00, 0x52, 0x0f, 0x69, 0x6e, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x4e,
	0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x20, 0x0a, 0x06, 0x64, 0x65, 0x67, 0x72, 0x65, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x01, 0x52, 0x06, 0x64,
	0x65, 0x67, 0x72, 0x65, 0x65, 0x88, 0x01, 0x01, 0x12, 0x31, 0x0a, 0x0f, 0x67, 0x72, 0x61, 0x64,
	0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x79, 0x65, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x02, 0x52, 0x0e, 0x67, 0x72, 0x61, 0x64, 0x75, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x59, 0x65, 0x61, 0x72, 0x88, 0x01, 0x01, 0x42, 0x13, 0x0a, 0x11, 0x5f,
	0x69, 0x6e, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x42, 0x09, 0x0a, 0x07, 0x5f, 0x64, 0x65, 0x67, 0x72, 0x65, 0x65, 0x42, 0x12, 0x0a, 0x10, 0x5f,
	0x67, 0x72, 0x61, 0x64, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x79, 0x65, 0x61, 0x72, 0x22,
	0x9c, 0x01, 0x0a, 0x09, 0x52, 0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x33, 0x0a,
	0x10, 0x69, 0x6e, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x00, 0x52, 0x0f,
	0x69, 0x6e, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x88,
	0x01, 0x01, 0x12, 0x31, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x79, 0x65, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x03, 0xe0, 0x41, 0x03,
	0x48, 0x01, 0x52, 0x0e, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x59, 0x65,
	0x61, 0x72, 0x88, 0x01, 0x01, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x69, 0x74,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x63,
	0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x79, 0x65, 0x61, 0x72, 0x22, 0x9d,
	0x01, 0x0a, 0x0a, 0x46, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x73, 0x68, 0x69, 0x70, 0x12, 0x33, 0x0a,
	0x10, 0x69, 0x6e, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x00, 0x52, 0x0f,
	0x69, 0x6e, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x88,
	0x01, 0x01, 0x12, 0x31, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x79, 0x65, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x03, 0xe0, 0x41, 0x03,
	0x48, 0x01, 0x52, 0x0e, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x59, 0x65,
	0x61, 0x72, 0x88, 0x01, 0x01, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x69, 0x74,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x63,
	0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x79, 0x65, 0x61, 0x72, 0x42, 0x8c,
	0x02, 0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0x1a, 0x4c, 0x6f, 0x63, 0x61, 0x6c,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61,
	0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x37,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x3b, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x22, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64,
	0x73, 0x2e, 0x56, 0x31, 0x37, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xca,
	0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x37, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0xea, 0x02, 0x26, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41,
	0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56,
	0x31, 0x37, 0x3a, 0x3a, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v17_resources_local_services_employee_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v17_resources_local_services_employee_proto_rawDescData = file_google_ads_googleads_v17_resources_local_services_employee_proto_rawDesc
)

func file_google_ads_googleads_v17_resources_local_services_employee_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v17_resources_local_services_employee_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v17_resources_local_services_employee_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v17_resources_local_services_employee_proto_rawDescData)
	})
	return file_google_ads_googleads_v17_resources_local_services_employee_proto_rawDescData
}

var file_google_ads_googleads_v17_resources_local_services_employee_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_ads_googleads_v17_resources_local_services_employee_proto_goTypes = []any{
	(*LocalServicesEmployee)(nil), // 0: google.ads.googleads.v17.resources.LocalServicesEmployee
	(*UniversityDegree)(nil),      // 1: google.ads.googleads.v17.resources.UniversityDegree
	(*Residency)(nil),             // 2: google.ads.googleads.v17.resources.Residency
	(*Fellowship)(nil),            // 3: google.ads.googleads.v17.resources.Fellowship
	(enums.LocalServicesEmployeeStatusEnum_LocalServicesEmployeeStatus)(0), // 4: google.ads.googleads.v17.enums.LocalServicesEmployeeStatusEnum.LocalServicesEmployeeStatus
	(enums.LocalServicesEmployeeTypeEnum_LocalServicesEmployeeType)(0),     // 5: google.ads.googleads.v17.enums.LocalServicesEmployeeTypeEnum.LocalServicesEmployeeType
}
var file_google_ads_googleads_v17_resources_local_services_employee_proto_depIdxs = []int32{
	4, // 0: google.ads.googleads.v17.resources.LocalServicesEmployee.status:type_name -> google.ads.googleads.v17.enums.LocalServicesEmployeeStatusEnum.LocalServicesEmployeeStatus
	5, // 1: google.ads.googleads.v17.resources.LocalServicesEmployee.type:type_name -> google.ads.googleads.v17.enums.LocalServicesEmployeeTypeEnum.LocalServicesEmployeeType
	1, // 2: google.ads.googleads.v17.resources.LocalServicesEmployee.university_degrees:type_name -> google.ads.googleads.v17.resources.UniversityDegree
	2, // 3: google.ads.googleads.v17.resources.LocalServicesEmployee.residencies:type_name -> google.ads.googleads.v17.resources.Residency
	3, // 4: google.ads.googleads.v17.resources.LocalServicesEmployee.fellowships:type_name -> google.ads.googleads.v17.resources.Fellowship
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v17_resources_local_services_employee_proto_init() }
func file_google_ads_googleads_v17_resources_local_services_employee_proto_init() {
	if File_google_ads_googleads_v17_resources_local_services_employee_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v17_resources_local_services_employee_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*LocalServicesEmployee); i {
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
		file_google_ads_googleads_v17_resources_local_services_employee_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*UniversityDegree); i {
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
		file_google_ads_googleads_v17_resources_local_services_employee_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Residency); i {
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
		file_google_ads_googleads_v17_resources_local_services_employee_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Fellowship); i {
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
	file_google_ads_googleads_v17_resources_local_services_employee_proto_msgTypes[0].OneofWrappers = []any{}
	file_google_ads_googleads_v17_resources_local_services_employee_proto_msgTypes[1].OneofWrappers = []any{}
	file_google_ads_googleads_v17_resources_local_services_employee_proto_msgTypes[2].OneofWrappers = []any{}
	file_google_ads_googleads_v17_resources_local_services_employee_proto_msgTypes[3].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v17_resources_local_services_employee_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v17_resources_local_services_employee_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v17_resources_local_services_employee_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v17_resources_local_services_employee_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v17_resources_local_services_employee_proto = out.File
	file_google_ads_googleads_v17_resources_local_services_employee_proto_rawDesc = nil
	file_google_ads_googleads_v17_resources_local_services_employee_proto_goTypes = nil
	file_google_ads_googleads_v17_resources_local_services_employee_proto_depIdxs = nil
}
