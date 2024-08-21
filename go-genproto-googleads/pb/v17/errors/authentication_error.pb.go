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
// source: google/ads/googleads/v17/errors/authentication_error.proto

// copy from https://github.com/dictav/go-genproto-googleads
// and changed by TeamMomentum

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

// Enum describing possible authentication errors.
type AuthenticationErrorEnum_AuthenticationError int32

const (
	// Enum unspecified.
	AuthenticationErrorEnum_UNSPECIFIED AuthenticationErrorEnum_AuthenticationError = 0
	// The received error code is not known in this version.
	AuthenticationErrorEnum_UNKNOWN AuthenticationErrorEnum_AuthenticationError = 1
	// Authentication of the request failed.
	AuthenticationErrorEnum_AUTHENTICATION_ERROR AuthenticationErrorEnum_AuthenticationError = 2
	// Client customer ID is not a number.
	AuthenticationErrorEnum_CLIENT_CUSTOMER_ID_INVALID AuthenticationErrorEnum_AuthenticationError = 5
	// No customer found for the provided customer ID.
	AuthenticationErrorEnum_CUSTOMER_NOT_FOUND AuthenticationErrorEnum_AuthenticationError = 8
	// Client's Google account is deleted.
	AuthenticationErrorEnum_GOOGLE_ACCOUNT_DELETED AuthenticationErrorEnum_AuthenticationError = 9
	// Google account login token in the cookie is invalid.
	AuthenticationErrorEnum_GOOGLE_ACCOUNT_COOKIE_INVALID AuthenticationErrorEnum_AuthenticationError = 10
	// A problem occurred during Google account authentication.
	AuthenticationErrorEnum_GOOGLE_ACCOUNT_AUTHENTICATION_FAILED AuthenticationErrorEnum_AuthenticationError = 25
	// The user in the Google account login token does not match the user ID in
	// the cookie.
	AuthenticationErrorEnum_GOOGLE_ACCOUNT_USER_AND_ADS_USER_MISMATCH AuthenticationErrorEnum_AuthenticationError = 12
	// Login cookie is required for authentication.
	AuthenticationErrorEnum_LOGIN_COOKIE_REQUIRED AuthenticationErrorEnum_AuthenticationError = 13
	// The Google account that generated the OAuth access
	// token is not associated with a Google Ads account. Create a new
	// account, or add the Google account to an existing Google Ads account.
	AuthenticationErrorEnum_NOT_ADS_USER AuthenticationErrorEnum_AuthenticationError = 14
	// OAuth token in the header is not valid.
	AuthenticationErrorEnum_OAUTH_TOKEN_INVALID AuthenticationErrorEnum_AuthenticationError = 15
	// OAuth token in the header has expired.
	AuthenticationErrorEnum_OAUTH_TOKEN_EXPIRED AuthenticationErrorEnum_AuthenticationError = 16
	// OAuth token in the header has been disabled.
	AuthenticationErrorEnum_OAUTH_TOKEN_DISABLED AuthenticationErrorEnum_AuthenticationError = 17
	// OAuth token in the header has been revoked.
	AuthenticationErrorEnum_OAUTH_TOKEN_REVOKED AuthenticationErrorEnum_AuthenticationError = 18
	// OAuth token HTTP header is malformed.
	AuthenticationErrorEnum_OAUTH_TOKEN_HEADER_INVALID AuthenticationErrorEnum_AuthenticationError = 19
	// Login cookie is not valid.
	AuthenticationErrorEnum_LOGIN_COOKIE_INVALID AuthenticationErrorEnum_AuthenticationError = 20
	// User ID in the header is not a valid ID.
	AuthenticationErrorEnum_USER_ID_INVALID AuthenticationErrorEnum_AuthenticationError = 22
	// An account administrator changed this account's authentication settings.
	// To access this Google Ads account, enable 2-Step Verification in your
	// Google account at https://www.google.com/landing/2step.
	AuthenticationErrorEnum_TWO_STEP_VERIFICATION_NOT_ENROLLED AuthenticationErrorEnum_AuthenticationError = 23
	// An account administrator changed this account's authentication settings.
	// To access this Google Ads account, enable Advanced Protection in your
	// Google account at https://landing.google.com/advancedprotection.
	AuthenticationErrorEnum_ADVANCED_PROTECTION_NOT_ENROLLED AuthenticationErrorEnum_AuthenticationError = 24
	// The Cloud organization associated with the project is not recognized.
	AuthenticationErrorEnum_ORGANIZATION_NOT_RECOGNIZED AuthenticationErrorEnum_AuthenticationError = 26
	// The Cloud organization associated with the project is not approved for
	// prod access.
	AuthenticationErrorEnum_ORGANIZATION_NOT_APPROVED AuthenticationErrorEnum_AuthenticationError = 27
	// The Cloud organization associated with the project is not associated with
	// the developer token.
	AuthenticationErrorEnum_ORGANIZATION_NOT_ASSOCIATED_WITH_DEVELOPER_TOKEN AuthenticationErrorEnum_AuthenticationError = 28
)

// Enum value maps for AuthenticationErrorEnum_AuthenticationError.
var (
	AuthenticationErrorEnum_AuthenticationError_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "AUTHENTICATION_ERROR",
		5:  "CLIENT_CUSTOMER_ID_INVALID",
		8:  "CUSTOMER_NOT_FOUND",
		9:  "GOOGLE_ACCOUNT_DELETED",
		10: "GOOGLE_ACCOUNT_COOKIE_INVALID",
		25: "GOOGLE_ACCOUNT_AUTHENTICATION_FAILED",
		12: "GOOGLE_ACCOUNT_USER_AND_ADS_USER_MISMATCH",
		13: "LOGIN_COOKIE_REQUIRED",
		14: "NOT_ADS_USER",
		15: "OAUTH_TOKEN_INVALID",
		16: "OAUTH_TOKEN_EXPIRED",
		17: "OAUTH_TOKEN_DISABLED",
		18: "OAUTH_TOKEN_REVOKED",
		19: "OAUTH_TOKEN_HEADER_INVALID",
		20: "LOGIN_COOKIE_INVALID",
		22: "USER_ID_INVALID",
		23: "TWO_STEP_VERIFICATION_NOT_ENROLLED",
		24: "ADVANCED_PROTECTION_NOT_ENROLLED",
		26: "ORGANIZATION_NOT_RECOGNIZED",
		27: "ORGANIZATION_NOT_APPROVED",
		28: "ORGANIZATION_NOT_ASSOCIATED_WITH_DEVELOPER_TOKEN",
	}
	AuthenticationErrorEnum_AuthenticationError_value = map[string]int32{
		"UNSPECIFIED":                                      0,
		"UNKNOWN":                                          1,
		"AUTHENTICATION_ERROR":                             2,
		"CLIENT_CUSTOMER_ID_INVALID":                       5,
		"CUSTOMER_NOT_FOUND":                               8,
		"GOOGLE_ACCOUNT_DELETED":                           9,
		"GOOGLE_ACCOUNT_COOKIE_INVALID":                    10,
		"GOOGLE_ACCOUNT_AUTHENTICATION_FAILED":             25,
		"GOOGLE_ACCOUNT_USER_AND_ADS_USER_MISMATCH":        12,
		"LOGIN_COOKIE_REQUIRED":                            13,
		"NOT_ADS_USER":                                     14,
		"OAUTH_TOKEN_INVALID":                              15,
		"OAUTH_TOKEN_EXPIRED":                              16,
		"OAUTH_TOKEN_DISABLED":                             17,
		"OAUTH_TOKEN_REVOKED":                              18,
		"OAUTH_TOKEN_HEADER_INVALID":                       19,
		"LOGIN_COOKIE_INVALID":                             20,
		"USER_ID_INVALID":                                  22,
		"TWO_STEP_VERIFICATION_NOT_ENROLLED":               23,
		"ADVANCED_PROTECTION_NOT_ENROLLED":                 24,
		"ORGANIZATION_NOT_RECOGNIZED":                      26,
		"ORGANIZATION_NOT_APPROVED":                        27,
		"ORGANIZATION_NOT_ASSOCIATED_WITH_DEVELOPER_TOKEN": 28,
	}
)

func (x AuthenticationErrorEnum_AuthenticationError) Enum() *AuthenticationErrorEnum_AuthenticationError {
	p := new(AuthenticationErrorEnum_AuthenticationError)
	*p = x
	return p
}

func (x AuthenticationErrorEnum_AuthenticationError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AuthenticationErrorEnum_AuthenticationError) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v17_errors_authentication_error_proto_enumTypes[0].Descriptor()
}

func (AuthenticationErrorEnum_AuthenticationError) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v17_errors_authentication_error_proto_enumTypes[0]
}

func (x AuthenticationErrorEnum_AuthenticationError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AuthenticationErrorEnum_AuthenticationError.Descriptor instead.
func (AuthenticationErrorEnum_AuthenticationError) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v17_errors_authentication_error_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible authentication errors.
type AuthenticationErrorEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AuthenticationErrorEnum) Reset() {
	*x = AuthenticationErrorEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v17_errors_authentication_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticationErrorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticationErrorEnum) ProtoMessage() {}

func (x *AuthenticationErrorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v17_errors_authentication_error_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticationErrorEnum.ProtoReflect.Descriptor instead.
func (*AuthenticationErrorEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v17_errors_authentication_error_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v17_errors_authentication_error_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v17_errors_authentication_error_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x37, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x22, 0xde, 0x05,
	0x0a, 0x17, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0xc2, 0x05, 0x0a, 0x13, 0x41, 0x75,
	0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12,
	0x18, 0x0a, 0x14, 0x41, 0x55, 0x54, 0x48, 0x45, 0x4e, 0x54, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x12, 0x1e, 0x0a, 0x1a, 0x43, 0x4c, 0x49,
	0x45, 0x4e, 0x54, 0x5f, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x45, 0x52, 0x5f, 0x49, 0x44, 0x5f,
	0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x05, 0x12, 0x16, 0x0a, 0x12, 0x43, 0x55, 0x53,
	0x54, 0x4f, 0x4d, 0x45, 0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10,
	0x08, 0x12, 0x1a, 0x0a, 0x16, 0x47, 0x4f, 0x4f, 0x47, 0x4c, 0x45, 0x5f, 0x41, 0x43, 0x43, 0x4f,
	0x55, 0x4e, 0x54, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x09, 0x12, 0x21, 0x0a,
	0x1d, 0x47, 0x4f, 0x4f, 0x47, 0x4c, 0x45, 0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f,
	0x43, 0x4f, 0x4f, 0x4b, 0x49, 0x45, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x0a,
	0x12, 0x28, 0x0a, 0x24, 0x47, 0x4f, 0x4f, 0x47, 0x4c, 0x45, 0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55,
	0x4e, 0x54, 0x5f, 0x41, 0x55, 0x54, 0x48, 0x45, 0x4e, 0x54, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x19, 0x12, 0x2d, 0x0a, 0x29, 0x47, 0x4f,
	0x4f, 0x47, 0x4c, 0x45, 0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x55, 0x53, 0x45,
	0x52, 0x5f, 0x41, 0x4e, 0x44, 0x5f, 0x41, 0x44, 0x53, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x4d,
	0x49, 0x53, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x10, 0x0c, 0x12, 0x19, 0x0a, 0x15, 0x4c, 0x4f, 0x47,
	0x49, 0x4e, 0x5f, 0x43, 0x4f, 0x4f, 0x4b, 0x49, 0x45, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x49, 0x52,
	0x45, 0x44, 0x10, 0x0d, 0x12, 0x10, 0x0a, 0x0c, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x44, 0x53, 0x5f,
	0x55, 0x53, 0x45, 0x52, 0x10, 0x0e, 0x12, 0x17, 0x0a, 0x13, 0x4f, 0x41, 0x55, 0x54, 0x48, 0x5f,
	0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x0f, 0x12,
	0x17, 0x0a, 0x13, 0x4f, 0x41, 0x55, 0x54, 0x48, 0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x5f, 0x45,
	0x58, 0x50, 0x49, 0x52, 0x45, 0x44, 0x10, 0x10, 0x12, 0x18, 0x0a, 0x14, 0x4f, 0x41, 0x55, 0x54,
	0x48, 0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x5f, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x44,
	0x10, 0x11, 0x12, 0x17, 0x0a, 0x13, 0x4f, 0x41, 0x55, 0x54, 0x48, 0x5f, 0x54, 0x4f, 0x4b, 0x45,
	0x4e, 0x5f, 0x52, 0x45, 0x56, 0x4f, 0x4b, 0x45, 0x44, 0x10, 0x12, 0x12, 0x1e, 0x0a, 0x1a, 0x4f,
	0x41, 0x55, 0x54, 0x48, 0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x5f, 0x48, 0x45, 0x41, 0x44, 0x45,
	0x52, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x13, 0x12, 0x18, 0x0a, 0x14, 0x4c,
	0x4f, 0x47, 0x49, 0x4e, 0x5f, 0x43, 0x4f, 0x4f, 0x4b, 0x49, 0x45, 0x5f, 0x49, 0x4e, 0x56, 0x41,
	0x4c, 0x49, 0x44, 0x10, 0x14, 0x12, 0x13, 0x0a, 0x0f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x49, 0x44,
	0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x16, 0x12, 0x26, 0x0a, 0x22, 0x54, 0x57,
	0x4f, 0x5f, 0x53, 0x54, 0x45, 0x50, 0x5f, 0x56, 0x45, 0x52, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x4e, 0x52, 0x4f, 0x4c, 0x4c, 0x45, 0x44,
	0x10, 0x17, 0x12, 0x24, 0x0a, 0x20, 0x41, 0x44, 0x56, 0x41, 0x4e, 0x43, 0x45, 0x44, 0x5f, 0x50,
	0x52, 0x4f, 0x54, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x4e,
	0x52, 0x4f, 0x4c, 0x4c, 0x45, 0x44, 0x10, 0x18, 0x12, 0x1f, 0x0a, 0x1b, 0x4f, 0x52, 0x47, 0x41,
	0x4e, 0x49, 0x5a, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x52, 0x45, 0x43,
	0x4f, 0x47, 0x4e, 0x49, 0x5a, 0x45, 0x44, 0x10, 0x1a, 0x12, 0x1d, 0x0a, 0x19, 0x4f, 0x52, 0x47,
	0x41, 0x4e, 0x49, 0x5a, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x50,
	0x50, 0x52, 0x4f, 0x56, 0x45, 0x44, 0x10, 0x1b, 0x12, 0x34, 0x0a, 0x30, 0x4f, 0x52, 0x47, 0x41,
	0x4e, 0x49, 0x5a, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x53, 0x53,
	0x4f, 0x43, 0x49, 0x41, 0x54, 0x45, 0x44, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x5f, 0x44, 0x45, 0x56,
	0x45, 0x4c, 0x4f, 0x50, 0x45, 0x52, 0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x10, 0x1c, 0x42, 0xf8,
	0x01, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x42, 0x18, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x45, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e,
	0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x37, 0x2f, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0x3b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa,
	0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x37, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0xca, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x37, 0x5c, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x73, 0xea, 0x02, 0x23, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64,
	0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31,
	0x37, 0x3a, 0x3a, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_google_ads_googleads_v17_errors_authentication_error_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v17_errors_authentication_error_proto_rawDescData = file_google_ads_googleads_v17_errors_authentication_error_proto_rawDesc
)

func file_google_ads_googleads_v17_errors_authentication_error_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v17_errors_authentication_error_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v17_errors_authentication_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v17_errors_authentication_error_proto_rawDescData)
	})
	return file_google_ads_googleads_v17_errors_authentication_error_proto_rawDescData
}

var file_google_ads_googleads_v17_errors_authentication_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v17_errors_authentication_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v17_errors_authentication_error_proto_goTypes = []any{
	(AuthenticationErrorEnum_AuthenticationError)(0), // 0: google.ads.googleads.v17.errors.AuthenticationErrorEnum.AuthenticationError
	(*AuthenticationErrorEnum)(nil),                  // 1: google.ads.googleads.v17.errors.AuthenticationErrorEnum
}
var file_google_ads_googleads_v17_errors_authentication_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v17_errors_authentication_error_proto_init() }
func file_google_ads_googleads_v17_errors_authentication_error_proto_init() {
	if File_google_ads_googleads_v17_errors_authentication_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v17_errors_authentication_error_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*AuthenticationErrorEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v17_errors_authentication_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v17_errors_authentication_error_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v17_errors_authentication_error_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v17_errors_authentication_error_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v17_errors_authentication_error_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v17_errors_authentication_error_proto = out.File
	file_google_ads_googleads_v17_errors_authentication_error_proto_rawDesc = nil
	file_google_ads_googleads_v17_errors_authentication_error_proto_goTypes = nil
	file_google_ads_googleads_v17_errors_authentication_error_proto_depIdxs = nil
}
