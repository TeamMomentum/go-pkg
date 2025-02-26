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
// source: google/ads/googleads/v18/errors/request_error.proto

package errors

import (
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

// Enum describing possible request errors.
type RequestErrorEnum_RequestError int32

const (
	// Enum unspecified.
	RequestErrorEnum_UNSPECIFIED RequestErrorEnum_RequestError = 0
	// The received error code is not known in this version.
	RequestErrorEnum_UNKNOWN RequestErrorEnum_RequestError = 1
	// Resource name is required for this request.
	RequestErrorEnum_RESOURCE_NAME_MISSING RequestErrorEnum_RequestError = 3
	// Resource name provided is malformed.
	RequestErrorEnum_RESOURCE_NAME_MALFORMED RequestErrorEnum_RequestError = 4
	// Resource name provided is malformed.
	RequestErrorEnum_BAD_RESOURCE_ID RequestErrorEnum_RequestError = 17
	// Customer ID is invalid.
	RequestErrorEnum_INVALID_CUSTOMER_ID RequestErrorEnum_RequestError = 16
	// Mutate operation should have either create, update, or remove specified.
	RequestErrorEnum_OPERATION_REQUIRED RequestErrorEnum_RequestError = 5
	// Requested resource not found.
	RequestErrorEnum_RESOURCE_NOT_FOUND RequestErrorEnum_RequestError = 6
	// Next page token specified in user request is invalid.
	RequestErrorEnum_INVALID_PAGE_TOKEN RequestErrorEnum_RequestError = 7
	// Next page token specified in user request has expired.
	RequestErrorEnum_EXPIRED_PAGE_TOKEN RequestErrorEnum_RequestError = 8
	// Page size specified in user request is invalid.
	RequestErrorEnum_INVALID_PAGE_SIZE RequestErrorEnum_RequestError = 22
	// Setting the page size is not supported, and will be unavailable in a
	// future version.
	RequestErrorEnum_PAGE_SIZE_NOT_SUPPORTED RequestErrorEnum_RequestError = 40
	// Required field is missing.
	RequestErrorEnum_REQUIRED_FIELD_MISSING RequestErrorEnum_RequestError = 9
	// The field cannot be modified because it's immutable. It's also possible
	// that the field can be modified using 'create' operation but not 'update'.
	RequestErrorEnum_IMMUTABLE_FIELD RequestErrorEnum_RequestError = 11
	// Received too many entries in request.
	RequestErrorEnum_TOO_MANY_MUTATE_OPERATIONS RequestErrorEnum_RequestError = 13
	// Request cannot be executed by a manager account.
	RequestErrorEnum_CANNOT_BE_EXECUTED_BY_MANAGER_ACCOUNT RequestErrorEnum_RequestError = 14
	// Mutate request was attempting to modify a readonly field.
	// For instance, Budget fields can be requested for Ad Group,
	// but are read-only for adGroups:mutate.
	RequestErrorEnum_CANNOT_MODIFY_FOREIGN_FIELD RequestErrorEnum_RequestError = 15
	// Enum value is not permitted.
	RequestErrorEnum_INVALID_ENUM_VALUE RequestErrorEnum_RequestError = 18
	// The developer-token parameter is required for all requests.
	RequestErrorEnum_DEVELOPER_TOKEN_PARAMETER_MISSING RequestErrorEnum_RequestError = 19
	// The login-customer-id parameter is required for this request.
	RequestErrorEnum_LOGIN_CUSTOMER_ID_PARAMETER_MISSING RequestErrorEnum_RequestError = 20
	// page_token is set in the validate only request
	RequestErrorEnum_VALIDATE_ONLY_REQUEST_HAS_PAGE_TOKEN RequestErrorEnum_RequestError = 21
	// return_summary_row cannot be enabled if request did not select any
	// metrics field.
	RequestErrorEnum_CANNOT_RETURN_SUMMARY_ROW_FOR_REQUEST_WITHOUT_METRICS RequestErrorEnum_RequestError = 29
	// return_summary_row should not be enabled for validate only requests.
	RequestErrorEnum_CANNOT_RETURN_SUMMARY_ROW_FOR_VALIDATE_ONLY_REQUESTS RequestErrorEnum_RequestError = 30
	// return_summary_row parameter value should be the same between requests
	// with page_token field set and their original request.
	RequestErrorEnum_INCONSISTENT_RETURN_SUMMARY_ROW_VALUE RequestErrorEnum_RequestError = 31
	// The total results count cannot be returned if it was not requested in the
	// original request.
	RequestErrorEnum_TOTAL_RESULTS_COUNT_NOT_ORIGINALLY_REQUESTED RequestErrorEnum_RequestError = 32
	// Deadline specified by the client was too short.
	RequestErrorEnum_RPC_DEADLINE_TOO_SHORT RequestErrorEnum_RequestError = 33
	// This API version has been sunset and is no longer supported.
	RequestErrorEnum_UNSUPPORTED_VERSION RequestErrorEnum_RequestError = 38
	// The Google Cloud project in the request was not found.
	RequestErrorEnum_CLOUD_PROJECT_NOT_FOUND RequestErrorEnum_RequestError = 39
)

// Enum value maps for RequestErrorEnum_RequestError.
var (
	RequestErrorEnum_RequestError_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		3:  "RESOURCE_NAME_MISSING",
		4:  "RESOURCE_NAME_MALFORMED",
		17: "BAD_RESOURCE_ID",
		16: "INVALID_CUSTOMER_ID",
		5:  "OPERATION_REQUIRED",
		6:  "RESOURCE_NOT_FOUND",
		7:  "INVALID_PAGE_TOKEN",
		8:  "EXPIRED_PAGE_TOKEN",
		22: "INVALID_PAGE_SIZE",
		40: "PAGE_SIZE_NOT_SUPPORTED",
		9:  "REQUIRED_FIELD_MISSING",
		11: "IMMUTABLE_FIELD",
		13: "TOO_MANY_MUTATE_OPERATIONS",
		14: "CANNOT_BE_EXECUTED_BY_MANAGER_ACCOUNT",
		15: "CANNOT_MODIFY_FOREIGN_FIELD",
		18: "INVALID_ENUM_VALUE",
		19: "DEVELOPER_TOKEN_PARAMETER_MISSING",
		20: "LOGIN_CUSTOMER_ID_PARAMETER_MISSING",
		21: "VALIDATE_ONLY_REQUEST_HAS_PAGE_TOKEN",
		29: "CANNOT_RETURN_SUMMARY_ROW_FOR_REQUEST_WITHOUT_METRICS",
		30: "CANNOT_RETURN_SUMMARY_ROW_FOR_VALIDATE_ONLY_REQUESTS",
		31: "INCONSISTENT_RETURN_SUMMARY_ROW_VALUE",
		32: "TOTAL_RESULTS_COUNT_NOT_ORIGINALLY_REQUESTED",
		33: "RPC_DEADLINE_TOO_SHORT",
		38: "UNSUPPORTED_VERSION",
		39: "CLOUD_PROJECT_NOT_FOUND",
	}
	RequestErrorEnum_RequestError_value = map[string]int32{
		"UNSPECIFIED":                                           0,
		"UNKNOWN":                                               1,
		"RESOURCE_NAME_MISSING":                                 3,
		"RESOURCE_NAME_MALFORMED":                               4,
		"BAD_RESOURCE_ID":                                       17,
		"INVALID_CUSTOMER_ID":                                   16,
		"OPERATION_REQUIRED":                                    5,
		"RESOURCE_NOT_FOUND":                                    6,
		"INVALID_PAGE_TOKEN":                                    7,
		"EXPIRED_PAGE_TOKEN":                                    8,
		"INVALID_PAGE_SIZE":                                     22,
		"PAGE_SIZE_NOT_SUPPORTED":                               40,
		"REQUIRED_FIELD_MISSING":                                9,
		"IMMUTABLE_FIELD":                                       11,
		"TOO_MANY_MUTATE_OPERATIONS":                            13,
		"CANNOT_BE_EXECUTED_BY_MANAGER_ACCOUNT":                 14,
		"CANNOT_MODIFY_FOREIGN_FIELD":                           15,
		"INVALID_ENUM_VALUE":                                    18,
		"DEVELOPER_TOKEN_PARAMETER_MISSING":                     19,
		"LOGIN_CUSTOMER_ID_PARAMETER_MISSING":                   20,
		"VALIDATE_ONLY_REQUEST_HAS_PAGE_TOKEN":                  21,
		"CANNOT_RETURN_SUMMARY_ROW_FOR_REQUEST_WITHOUT_METRICS": 29,
		"CANNOT_RETURN_SUMMARY_ROW_FOR_VALIDATE_ONLY_REQUESTS":  30,
		"INCONSISTENT_RETURN_SUMMARY_ROW_VALUE":                 31,
		"TOTAL_RESULTS_COUNT_NOT_ORIGINALLY_REQUESTED":          32,
		"RPC_DEADLINE_TOO_SHORT":                                33,
		"UNSUPPORTED_VERSION":                                   38,
		"CLOUD_PROJECT_NOT_FOUND":                               39,
	}
)

func (x RequestErrorEnum_RequestError) Enum() *RequestErrorEnum_RequestError {
	p := new(RequestErrorEnum_RequestError)
	*p = x
	return p
}

func (x RequestErrorEnum_RequestError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RequestErrorEnum_RequestError) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v18_errors_request_error_proto_enumTypes[0].Descriptor()
}

func (RequestErrorEnum_RequestError) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v18_errors_request_error_proto_enumTypes[0]
}

func (x RequestErrorEnum_RequestError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RequestErrorEnum_RequestError.Descriptor instead.
func (RequestErrorEnum_RequestError) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v18_errors_request_error_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible request errors.
type RequestErrorEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RequestErrorEnum) Reset() {
	*x = RequestErrorEnum{}
	mi := &file_google_ads_googleads_v18_errors_request_error_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RequestErrorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestErrorEnum) ProtoMessage() {}

func (x *RequestErrorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v18_errors_request_error_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestErrorEnum.ProtoReflect.Descriptor instead.
func (*RequestErrorEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v18_errors_request_error_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v18_errors_request_error_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v18_errors_request_error_proto_rawDesc = []byte{
	0x0a, 0x33, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x38, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x38, 0x2e,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x22, 0x8e, 0x07, 0x0a, 0x10, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0xf9, 0x06, 0x0a, 0x0c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x0f, 0x0a, 0x0b,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a,
	0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x52, 0x45,
	0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x4d, 0x49, 0x53, 0x53,
	0x49, 0x4e, 0x47, 0x10, 0x03, 0x12, 0x1b, 0x0a, 0x17, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43,
	0x45, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x4d, 0x41, 0x4c, 0x46, 0x4f, 0x52, 0x4d, 0x45, 0x44,
	0x10, 0x04, 0x12, 0x13, 0x0a, 0x0f, 0x42, 0x41, 0x44, 0x5f, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52,
	0x43, 0x45, 0x5f, 0x49, 0x44, 0x10, 0x11, 0x12, 0x17, 0x0a, 0x13, 0x49, 0x4e, 0x56, 0x41, 0x4c,
	0x49, 0x44, 0x5f, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x45, 0x52, 0x5f, 0x49, 0x44, 0x10, 0x10,
	0x12, 0x16, 0x0a, 0x12, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45,
	0x51, 0x55, 0x49, 0x52, 0x45, 0x44, 0x10, 0x05, 0x12, 0x16, 0x0a, 0x12, 0x52, 0x45, 0x53, 0x4f,
	0x55, 0x52, 0x43, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x06,
	0x12, 0x16, 0x0a, 0x12, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x50, 0x41, 0x47, 0x45,
	0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x10, 0x07, 0x12, 0x16, 0x0a, 0x12, 0x45, 0x58, 0x50, 0x49,
	0x52, 0x45, 0x44, 0x5f, 0x50, 0x41, 0x47, 0x45, 0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x10, 0x08,
	0x12, 0x15, 0x0a, 0x11, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x50, 0x41, 0x47, 0x45,
	0x5f, 0x53, 0x49, 0x5a, 0x45, 0x10, 0x16, 0x12, 0x1b, 0x0a, 0x17, 0x50, 0x41, 0x47, 0x45, 0x5f,
	0x53, 0x49, 0x5a, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54,
	0x45, 0x44, 0x10, 0x28, 0x12, 0x1a, 0x0a, 0x16, 0x52, 0x45, 0x51, 0x55, 0x49, 0x52, 0x45, 0x44,
	0x5f, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x10, 0x09,
	0x12, 0x13, 0x0a, 0x0f, 0x49, 0x4d, 0x4d, 0x55, 0x54, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x46, 0x49,
	0x45, 0x4c, 0x44, 0x10, 0x0b, 0x12, 0x1e, 0x0a, 0x1a, 0x54, 0x4f, 0x4f, 0x5f, 0x4d, 0x41, 0x4e,
	0x59, 0x5f, 0x4d, 0x55, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49,
	0x4f, 0x4e, 0x53, 0x10, 0x0d, 0x12, 0x29, 0x0a, 0x25, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f,
	0x42, 0x45, 0x5f, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x45, 0x44, 0x5f, 0x42, 0x59, 0x5f, 0x4d,
	0x41, 0x4e, 0x41, 0x47, 0x45, 0x52, 0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x10, 0x0e,
	0x12, 0x1f, 0x0a, 0x1b, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x4d, 0x4f, 0x44, 0x49, 0x46,
	0x59, 0x5f, 0x46, 0x4f, 0x52, 0x45, 0x49, 0x47, 0x4e, 0x5f, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x10,
	0x0f, 0x12, 0x16, 0x0a, 0x12, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x45, 0x4e, 0x55,
	0x4d, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x10, 0x12, 0x12, 0x25, 0x0a, 0x21, 0x44, 0x45, 0x56,
	0x45, 0x4c, 0x4f, 0x50, 0x45, 0x52, 0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x5f, 0x50, 0x41, 0x52,
	0x41, 0x4d, 0x45, 0x54, 0x45, 0x52, 0x5f, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x10, 0x13,
	0x12, 0x27, 0x0a, 0x23, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x5f, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d,
	0x45, 0x52, 0x5f, 0x49, 0x44, 0x5f, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x45, 0x54, 0x45, 0x52, 0x5f,
	0x4d, 0x49, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x10, 0x14, 0x12, 0x28, 0x0a, 0x24, 0x56, 0x41, 0x4c,
	0x49, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x4f, 0x4e, 0x4c, 0x59, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45,
	0x53, 0x54, 0x5f, 0x48, 0x41, 0x53, 0x5f, 0x50, 0x41, 0x47, 0x45, 0x5f, 0x54, 0x4f, 0x4b, 0x45,
	0x4e, 0x10, 0x15, 0x12, 0x39, 0x0a, 0x35, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x52, 0x45,
	0x54, 0x55, 0x52, 0x4e, 0x5f, 0x53, 0x55, 0x4d, 0x4d, 0x41, 0x52, 0x59, 0x5f, 0x52, 0x4f, 0x57,
	0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x57, 0x49, 0x54,
	0x48, 0x4f, 0x55, 0x54, 0x5f, 0x4d, 0x45, 0x54, 0x52, 0x49, 0x43, 0x53, 0x10, 0x1d, 0x12, 0x38,
	0x0a, 0x34, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x52, 0x45, 0x54, 0x55, 0x52, 0x4e, 0x5f,
	0x53, 0x55, 0x4d, 0x4d, 0x41, 0x52, 0x59, 0x5f, 0x52, 0x4f, 0x57, 0x5f, 0x46, 0x4f, 0x52, 0x5f,
	0x56, 0x41, 0x4c, 0x49, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x4f, 0x4e, 0x4c, 0x59, 0x5f, 0x52, 0x45,
	0x51, 0x55, 0x45, 0x53, 0x54, 0x53, 0x10, 0x1e, 0x12, 0x29, 0x0a, 0x25, 0x49, 0x4e, 0x43, 0x4f,
	0x4e, 0x53, 0x49, 0x53, 0x54, 0x45, 0x4e, 0x54, 0x5f, 0x52, 0x45, 0x54, 0x55, 0x52, 0x4e, 0x5f,
	0x53, 0x55, 0x4d, 0x4d, 0x41, 0x52, 0x59, 0x5f, 0x52, 0x4f, 0x57, 0x5f, 0x56, 0x41, 0x4c, 0x55,
	0x45, 0x10, 0x1f, 0x12, 0x30, 0x0a, 0x2c, 0x54, 0x4f, 0x54, 0x41, 0x4c, 0x5f, 0x52, 0x45, 0x53,
	0x55, 0x4c, 0x54, 0x53, 0x5f, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x4f,
	0x52, 0x49, 0x47, 0x49, 0x4e, 0x41, 0x4c, 0x4c, 0x59, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53,
	0x54, 0x45, 0x44, 0x10, 0x20, 0x12, 0x1a, 0x0a, 0x16, 0x52, 0x50, 0x43, 0x5f, 0x44, 0x45, 0x41,
	0x44, 0x4c, 0x49, 0x4e, 0x45, 0x5f, 0x54, 0x4f, 0x4f, 0x5f, 0x53, 0x48, 0x4f, 0x52, 0x54, 0x10,
	0x21, 0x12, 0x17, 0x0a, 0x13, 0x55, 0x4e, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x45, 0x44,
	0x5f, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x10, 0x26, 0x12, 0x1b, 0x0a, 0x17, 0x43, 0x4c,
	0x4f, 0x55, 0x44, 0x5f, 0x50, 0x52, 0x4f, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x4e, 0x4f, 0x54, 0x5f,
	0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x27, 0x42, 0xf1, 0x01, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x38, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x42,
	0x11, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c,
	0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x38, 0x2f, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0x3b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41,
	0x41, 0xaa, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x38, 0x2e, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x73, 0xca, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73,
	0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x38, 0x5c, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0xea, 0x02, 0x23, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a,
	0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a,
	0x56, 0x31, 0x38, 0x3a, 0x3a, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v18_errors_request_error_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v18_errors_request_error_proto_rawDescData = file_google_ads_googleads_v18_errors_request_error_proto_rawDesc
)

func file_google_ads_googleads_v18_errors_request_error_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v18_errors_request_error_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v18_errors_request_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v18_errors_request_error_proto_rawDescData)
	})
	return file_google_ads_googleads_v18_errors_request_error_proto_rawDescData
}

var file_google_ads_googleads_v18_errors_request_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v18_errors_request_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v18_errors_request_error_proto_goTypes = []any{
	(RequestErrorEnum_RequestError)(0), // 0: google.ads.googleads.v18.errors.RequestErrorEnum.RequestError
	(*RequestErrorEnum)(nil),           // 1: google.ads.googleads.v18.errors.RequestErrorEnum
}
var file_google_ads_googleads_v18_errors_request_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v18_errors_request_error_proto_init() }
func file_google_ads_googleads_v18_errors_request_error_proto_init() {
	if File_google_ads_googleads_v18_errors_request_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v18_errors_request_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v18_errors_request_error_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v18_errors_request_error_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v18_errors_request_error_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v18_errors_request_error_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v18_errors_request_error_proto = out.File
	file_google_ads_googleads_v18_errors_request_error_proto_rawDesc = nil
	file_google_ads_googleads_v18_errors_request_error_proto_goTypes = nil
	file_google_ads_googleads_v18_errors_request_error_proto_depIdxs = nil
}
