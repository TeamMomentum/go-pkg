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
// source: google/ads/googleads/v18/common/matching_function.proto

package common

import (
	enums "github.com/TeamMomentum/go-pkg/go-genproto-googleads/pb/v18/enums"
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

// Matching function associated with a
// CustomerFeed, CampaignFeed, or AdGroupFeed. The matching function is used
// to filter the set of feed items selected.
type MatchingFunction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// String representation of the Function.
	//
	// Examples:
	//
	// 1. IDENTITY(true) or IDENTITY(false). All or no feed items served.
	// 2. EQUALS(CONTEXT.DEVICE,"Mobile")
	// 3. IN(FEED_ITEM_ID,{1000001,1000002,1000003})
	// 4. CONTAINS_ANY(FeedAttribute[12345678,0],{"Mars cruise","Venus cruise"})
	// 5. AND(IN(FEED_ITEM_ID,{10001,10002}),EQUALS(CONTEXT.DEVICE,"Mobile"))
	//
	// For more details, visit
	// https://developers.google.com/google-ads/api/docs/extensions/feeds/matching-functions
	//
	// Note that because multiple strings may represent the same underlying
	// function (whitespace and single versus double quotation marks, for
	// example), the value returned may not be identical to the string sent in a
	// mutate request.
	FunctionString *string `protobuf:"bytes,5,opt,name=function_string,json=functionString,proto3,oneof" json:"function_string,omitempty"`
	// Operator for a function.
	Operator enums.MatchingFunctionOperatorEnum_MatchingFunctionOperator `protobuf:"varint,4,opt,name=operator,proto3,enum=google.ads.googleads.v18.enums.MatchingFunctionOperatorEnum_MatchingFunctionOperator" json:"operator,omitempty"`
	// The operands on the left hand side of the equation. This is also the
	// operand to be used for single operand expressions such as NOT.
	LeftOperands []*Operand `protobuf:"bytes,2,rep,name=left_operands,json=leftOperands,proto3" json:"left_operands,omitempty"`
	// The operands on the right hand side of the equation.
	RightOperands []*Operand `protobuf:"bytes,3,rep,name=right_operands,json=rightOperands,proto3" json:"right_operands,omitempty"`
}

func (x *MatchingFunction) Reset() {
	*x = MatchingFunction{}
	mi := &file_google_ads_googleads_v18_common_matching_function_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MatchingFunction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchingFunction) ProtoMessage() {}

func (x *MatchingFunction) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v18_common_matching_function_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchingFunction.ProtoReflect.Descriptor instead.
func (*MatchingFunction) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v18_common_matching_function_proto_rawDescGZIP(), []int{0}
}

func (x *MatchingFunction) GetFunctionString() string {
	if x != nil && x.FunctionString != nil {
		return *x.FunctionString
	}
	return ""
}

func (x *MatchingFunction) GetOperator() enums.MatchingFunctionOperatorEnum_MatchingFunctionOperator {
	if x != nil {
		return x.Operator
	}
	return enums.MatchingFunctionOperatorEnum_MatchingFunctionOperator(0)
}

func (x *MatchingFunction) GetLeftOperands() []*Operand {
	if x != nil {
		return x.LeftOperands
	}
	return nil
}

func (x *MatchingFunction) GetRightOperands() []*Operand {
	if x != nil {
		return x.RightOperands
	}
	return nil
}

// An operand in a matching function.
type Operand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Different operands that can be used in a matching function. Required.
	//
	// Types that are assignable to FunctionArgumentOperand:
	//
	//	*Operand_ConstantOperand_
	//	*Operand_FeedAttributeOperand_
	//	*Operand_FunctionOperand_
	//	*Operand_RequestContextOperand_
	FunctionArgumentOperand isOperand_FunctionArgumentOperand `protobuf_oneof:"function_argument_operand"`
}

func (x *Operand) Reset() {
	*x = Operand{}
	mi := &file_google_ads_googleads_v18_common_matching_function_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Operand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Operand) ProtoMessage() {}

func (x *Operand) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v18_common_matching_function_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Operand.ProtoReflect.Descriptor instead.
func (*Operand) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v18_common_matching_function_proto_rawDescGZIP(), []int{1}
}

func (m *Operand) GetFunctionArgumentOperand() isOperand_FunctionArgumentOperand {
	if m != nil {
		return m.FunctionArgumentOperand
	}
	return nil
}

func (x *Operand) GetConstantOperand() *Operand_ConstantOperand {
	if x, ok := x.GetFunctionArgumentOperand().(*Operand_ConstantOperand_); ok {
		return x.ConstantOperand
	}
	return nil
}

func (x *Operand) GetFeedAttributeOperand() *Operand_FeedAttributeOperand {
	if x, ok := x.GetFunctionArgumentOperand().(*Operand_FeedAttributeOperand_); ok {
		return x.FeedAttributeOperand
	}
	return nil
}

func (x *Operand) GetFunctionOperand() *Operand_FunctionOperand {
	if x, ok := x.GetFunctionArgumentOperand().(*Operand_FunctionOperand_); ok {
		return x.FunctionOperand
	}
	return nil
}

func (x *Operand) GetRequestContextOperand() *Operand_RequestContextOperand {
	if x, ok := x.GetFunctionArgumentOperand().(*Operand_RequestContextOperand_); ok {
		return x.RequestContextOperand
	}
	return nil
}

type isOperand_FunctionArgumentOperand interface {
	isOperand_FunctionArgumentOperand()
}

type Operand_ConstantOperand_ struct {
	// A constant operand in a matching function.
	ConstantOperand *Operand_ConstantOperand `protobuf:"bytes,1,opt,name=constant_operand,json=constantOperand,proto3,oneof"`
}

type Operand_FeedAttributeOperand_ struct {
	// This operand specifies a feed attribute in feed.
	FeedAttributeOperand *Operand_FeedAttributeOperand `protobuf:"bytes,2,opt,name=feed_attribute_operand,json=feedAttributeOperand,proto3,oneof"`
}

type Operand_FunctionOperand_ struct {
	// A function operand in a matching function.
	// Used to represent nested functions.
	FunctionOperand *Operand_FunctionOperand `protobuf:"bytes,3,opt,name=function_operand,json=functionOperand,proto3,oneof"`
}

type Operand_RequestContextOperand_ struct {
	// An operand in a function referring to a value in the request context.
	RequestContextOperand *Operand_RequestContextOperand `protobuf:"bytes,4,opt,name=request_context_operand,json=requestContextOperand,proto3,oneof"`
}

func (*Operand_ConstantOperand_) isOperand_FunctionArgumentOperand() {}

func (*Operand_FeedAttributeOperand_) isOperand_FunctionArgumentOperand() {}

func (*Operand_FunctionOperand_) isOperand_FunctionArgumentOperand() {}

func (*Operand_RequestContextOperand_) isOperand_FunctionArgumentOperand() {}

// A constant operand in a matching function.
type Operand_ConstantOperand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Constant operand values. Required.
	//
	// Types that are assignable to ConstantOperandValue:
	//
	//	*Operand_ConstantOperand_StringValue
	//	*Operand_ConstantOperand_LongValue
	//	*Operand_ConstantOperand_BooleanValue
	//	*Operand_ConstantOperand_DoubleValue
	ConstantOperandValue isOperand_ConstantOperand_ConstantOperandValue `protobuf_oneof:"constant_operand_value"`
}

func (x *Operand_ConstantOperand) Reset() {
	*x = Operand_ConstantOperand{}
	mi := &file_google_ads_googleads_v18_common_matching_function_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Operand_ConstantOperand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Operand_ConstantOperand) ProtoMessage() {}

func (x *Operand_ConstantOperand) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v18_common_matching_function_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Operand_ConstantOperand.ProtoReflect.Descriptor instead.
func (*Operand_ConstantOperand) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v18_common_matching_function_proto_rawDescGZIP(), []int{1, 0}
}

func (m *Operand_ConstantOperand) GetConstantOperandValue() isOperand_ConstantOperand_ConstantOperandValue {
	if m != nil {
		return m.ConstantOperandValue
	}
	return nil
}

func (x *Operand_ConstantOperand) GetStringValue() string {
	if x, ok := x.GetConstantOperandValue().(*Operand_ConstantOperand_StringValue); ok {
		return x.StringValue
	}
	return ""
}

func (x *Operand_ConstantOperand) GetLongValue() int64 {
	if x, ok := x.GetConstantOperandValue().(*Operand_ConstantOperand_LongValue); ok {
		return x.LongValue
	}
	return 0
}

func (x *Operand_ConstantOperand) GetBooleanValue() bool {
	if x, ok := x.GetConstantOperandValue().(*Operand_ConstantOperand_BooleanValue); ok {
		return x.BooleanValue
	}
	return false
}

func (x *Operand_ConstantOperand) GetDoubleValue() float64 {
	if x, ok := x.GetConstantOperandValue().(*Operand_ConstantOperand_DoubleValue); ok {
		return x.DoubleValue
	}
	return 0
}

type isOperand_ConstantOperand_ConstantOperandValue interface {
	isOperand_ConstantOperand_ConstantOperandValue()
}

type Operand_ConstantOperand_StringValue struct {
	// String value of the operand if it is a string type.
	StringValue string `protobuf:"bytes,5,opt,name=string_value,json=stringValue,proto3,oneof"`
}

type Operand_ConstantOperand_LongValue struct {
	// Int64 value of the operand if it is a int64 type.
	LongValue int64 `protobuf:"varint,6,opt,name=long_value,json=longValue,proto3,oneof"`
}

type Operand_ConstantOperand_BooleanValue struct {
	// Boolean value of the operand if it is a boolean type.
	BooleanValue bool `protobuf:"varint,7,opt,name=boolean_value,json=booleanValue,proto3,oneof"`
}

type Operand_ConstantOperand_DoubleValue struct {
	// Double value of the operand if it is a double type.
	DoubleValue float64 `protobuf:"fixed64,8,opt,name=double_value,json=doubleValue,proto3,oneof"`
}

func (*Operand_ConstantOperand_StringValue) isOperand_ConstantOperand_ConstantOperandValue() {}

func (*Operand_ConstantOperand_LongValue) isOperand_ConstantOperand_ConstantOperandValue() {}

func (*Operand_ConstantOperand_BooleanValue) isOperand_ConstantOperand_ConstantOperandValue() {}

func (*Operand_ConstantOperand_DoubleValue) isOperand_ConstantOperand_ConstantOperandValue() {}

// A feed attribute operand in a matching function.
// Used to represent a feed attribute in feed.
type Operand_FeedAttributeOperand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The associated feed. Required.
	FeedId *int64 `protobuf:"varint,3,opt,name=feed_id,json=feedId,proto3,oneof" json:"feed_id,omitempty"`
	// Id of the referenced feed attribute. Required.
	FeedAttributeId *int64 `protobuf:"varint,4,opt,name=feed_attribute_id,json=feedAttributeId,proto3,oneof" json:"feed_attribute_id,omitempty"`
}

func (x *Operand_FeedAttributeOperand) Reset() {
	*x = Operand_FeedAttributeOperand{}
	mi := &file_google_ads_googleads_v18_common_matching_function_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Operand_FeedAttributeOperand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Operand_FeedAttributeOperand) ProtoMessage() {}

func (x *Operand_FeedAttributeOperand) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v18_common_matching_function_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Operand_FeedAttributeOperand.ProtoReflect.Descriptor instead.
func (*Operand_FeedAttributeOperand) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v18_common_matching_function_proto_rawDescGZIP(), []int{1, 1}
}

func (x *Operand_FeedAttributeOperand) GetFeedId() int64 {
	if x != nil && x.FeedId != nil {
		return *x.FeedId
	}
	return 0
}

func (x *Operand_FeedAttributeOperand) GetFeedAttributeId() int64 {
	if x != nil && x.FeedAttributeId != nil {
		return *x.FeedAttributeId
	}
	return 0
}

// A function operand in a matching function.
// Used to represent nested functions.
type Operand_FunctionOperand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The matching function held in this operand.
	MatchingFunction *MatchingFunction `protobuf:"bytes,1,opt,name=matching_function,json=matchingFunction,proto3" json:"matching_function,omitempty"`
}

func (x *Operand_FunctionOperand) Reset() {
	*x = Operand_FunctionOperand{}
	mi := &file_google_ads_googleads_v18_common_matching_function_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Operand_FunctionOperand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Operand_FunctionOperand) ProtoMessage() {}

func (x *Operand_FunctionOperand) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v18_common_matching_function_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Operand_FunctionOperand.ProtoReflect.Descriptor instead.
func (*Operand_FunctionOperand) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v18_common_matching_function_proto_rawDescGZIP(), []int{1, 2}
}

func (x *Operand_FunctionOperand) GetMatchingFunction() *MatchingFunction {
	if x != nil {
		return x.MatchingFunction
	}
	return nil
}

// An operand in a function referring to a value in the request context.
type Operand_RequestContextOperand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Type of value to be referred in the request context.
	ContextType enums.MatchingFunctionContextTypeEnum_MatchingFunctionContextType `protobuf:"varint,1,opt,name=context_type,json=contextType,proto3,enum=google.ads.googleads.v18.enums.MatchingFunctionContextTypeEnum_MatchingFunctionContextType" json:"context_type,omitempty"`
}

func (x *Operand_RequestContextOperand) Reset() {
	*x = Operand_RequestContextOperand{}
	mi := &file_google_ads_googleads_v18_common_matching_function_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Operand_RequestContextOperand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Operand_RequestContextOperand) ProtoMessage() {}

func (x *Operand_RequestContextOperand) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v18_common_matching_function_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Operand_RequestContextOperand.ProtoReflect.Descriptor instead.
func (*Operand_RequestContextOperand) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v18_common_matching_function_proto_rawDescGZIP(), []int{1, 3}
}

func (x *Operand_RequestContextOperand) GetContextType() enums.MatchingFunctionContextTypeEnum_MatchingFunctionContextType {
	if x != nil {
		return x.ContextType
	}
	return enums.MatchingFunctionContextTypeEnum_MatchingFunctionContextType(0)
}

var File_google_ads_googleads_v18_common_matching_function_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v18_common_matching_function_proto_rawDesc = []byte{
	0x0a, 0x37, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x38, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x75, 0x6e, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x38, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x1a, 0x43, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2f, 0x76, 0x31, 0x38, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68,
	0x69, 0x6e, 0x67, 0x5f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x3f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x38, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xe7, 0x02, 0x0a, 0x10, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x46, 0x75, 0x6e,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x0f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x0e, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x88, 0x01, 0x01, 0x12, 0x71, 0x0a, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x55, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x38,
	0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x46,
	0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x45,
	0x6e, 0x75, 0x6d, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x46, 0x75, 0x6e, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x08, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x4d, 0x0a, 0x0d, 0x6c, 0x65, 0x66, 0x74, 0x5f, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x6e, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x38, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x6e, 0x64, 0x52, 0x0c, 0x6c, 0x65, 0x66, 0x74, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x6e, 0x64, 0x73, 0x12, 0x4f, 0x0a, 0x0e, 0x72, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x6e, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x38, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x6e, 0x64, 0x52, 0x0d, 0x72, 0x69, 0x67, 0x68, 0x74, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x6e, 0x64, 0x73, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x66, 0x75, 0x6e, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0xbc, 0x08, 0x0a, 0x07, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x6e, 0x64, 0x12, 0x65, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x74, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x38, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x38, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x6e, 0x64, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x6e, 0x64, 0x48, 0x00, 0x52, 0x0f, 0x63, 0x6f,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x6e, 0x64, 0x12, 0x75, 0x0a,
	0x16, 0x66, 0x65, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x5f,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3d, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x38, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x6e, 0x64, 0x2e, 0x46, 0x65, 0x65, 0x64, 0x41, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x6e, 0x64, 0x48, 0x00, 0x52, 0x14,
	0x66, 0x65, 0x65, 0x64, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x6e, 0x64, 0x12, 0x65, 0x0a, 0x10, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x38,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x38, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x6e, 0x64, 0x2e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x6e, 0x64, 0x48, 0x00, 0x52, 0x0f, 0x66, 0x75, 0x6e, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x6e, 0x64, 0x12, 0x78, 0x0a, 0x17, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3e, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x38, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x6e, 0x64, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x6e, 0x64, 0x48, 0x00, 0x52, 0x15,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x6e, 0x64, 0x1a, 0xbd, 0x01, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x6e, 0x64, 0x12, 0x23, 0x0a, 0x0c, 0x73, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1f,
	0x0a, 0x0a, 0x6c, 0x6f, 0x6e, 0x67, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x03, 0x48, 0x00, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x25, 0x0a, 0x0d, 0x62, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x0c, 0x62, 0x6f, 0x6f, 0x6c, 0x65, 0x61,
	0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x23, 0x0a, 0x0c, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65,
	0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x48, 0x00, 0x52, 0x0b,
	0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x18, 0x0a, 0x16, 0x63,
	0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x6e, 0x64, 0x5f,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x87, 0x01, 0x0a, 0x14, 0x46, 0x65, 0x65, 0x64, 0x41, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x6e, 0x64, 0x12, 0x1c,
	0x0a, 0x07, 0x66, 0x65, 0x65, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x48,
	0x00, 0x52, 0x06, 0x66, 0x65, 0x65, 0x64, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x2f, 0x0a, 0x11,
	0x66, 0x65, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x48, 0x01, 0x52, 0x0f, 0x66, 0x65, 0x65, 0x64, 0x41,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a,
	0x08, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x5f, 0x69, 0x64, 0x42, 0x14, 0x0a, 0x12, 0x5f, 0x66, 0x65,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x1a,
	0x71, 0x0a, 0x0f, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x6e, 0x64, 0x12, 0x5e, 0x0a, 0x11, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x5f, 0x66,
	0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x38, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x10, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x1a, 0x97, 0x01, 0x0a, 0x15, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x6e, 0x64, 0x12, 0x7e, 0x0a, 0x0c,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x5b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x38, 0x2e, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x46, 0x75, 0x6e, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x54, 0x79, 0x70, 0x65, 0x45,
	0x6e, 0x75, 0x6d, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x46, 0x75, 0x6e, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x54, 0x79, 0x70, 0x65, 0x42, 0x1b, 0x0a, 0x19,
	0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x6e, 0x64, 0x42, 0xf5, 0x01, 0x0a, 0x23, 0x63, 0x6f,
	0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x38, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x42, 0x15, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x46, 0x75, 0x6e, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f,
	0x76, 0x31, 0x38, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56,
	0x31, 0x38, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0xca, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73,
	0x5c, 0x56, 0x31, 0x38, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0xea, 0x02, 0x23, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x38, 0x3a, 0x3a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v18_common_matching_function_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v18_common_matching_function_proto_rawDescData = file_google_ads_googleads_v18_common_matching_function_proto_rawDesc
)

func file_google_ads_googleads_v18_common_matching_function_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v18_common_matching_function_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v18_common_matching_function_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v18_common_matching_function_proto_rawDescData)
	})
	return file_google_ads_googleads_v18_common_matching_function_proto_rawDescData
}

var file_google_ads_googleads_v18_common_matching_function_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_google_ads_googleads_v18_common_matching_function_proto_goTypes = []any{
	(*MatchingFunction)(nil),                                               // 0: google.ads.googleads.v18.common.MatchingFunction
	(*Operand)(nil),                                                        // 1: google.ads.googleads.v18.common.Operand
	(*Operand_ConstantOperand)(nil),                                        // 2: google.ads.googleads.v18.common.Operand.ConstantOperand
	(*Operand_FeedAttributeOperand)(nil),                                   // 3: google.ads.googleads.v18.common.Operand.FeedAttributeOperand
	(*Operand_FunctionOperand)(nil),                                        // 4: google.ads.googleads.v18.common.Operand.FunctionOperand
	(*Operand_RequestContextOperand)(nil),                                  // 5: google.ads.googleads.v18.common.Operand.RequestContextOperand
	(enums.MatchingFunctionOperatorEnum_MatchingFunctionOperator)(0),       // 6: google.ads.googleads.v18.enums.MatchingFunctionOperatorEnum.MatchingFunctionOperator
	(enums.MatchingFunctionContextTypeEnum_MatchingFunctionContextType)(0), // 7: google.ads.googleads.v18.enums.MatchingFunctionContextTypeEnum.MatchingFunctionContextType
}
var file_google_ads_googleads_v18_common_matching_function_proto_depIdxs = []int32{
	6, // 0: google.ads.googleads.v18.common.MatchingFunction.operator:type_name -> google.ads.googleads.v18.enums.MatchingFunctionOperatorEnum.MatchingFunctionOperator
	1, // 1: google.ads.googleads.v18.common.MatchingFunction.left_operands:type_name -> google.ads.googleads.v18.common.Operand
	1, // 2: google.ads.googleads.v18.common.MatchingFunction.right_operands:type_name -> google.ads.googleads.v18.common.Operand
	2, // 3: google.ads.googleads.v18.common.Operand.constant_operand:type_name -> google.ads.googleads.v18.common.Operand.ConstantOperand
	3, // 4: google.ads.googleads.v18.common.Operand.feed_attribute_operand:type_name -> google.ads.googleads.v18.common.Operand.FeedAttributeOperand
	4, // 5: google.ads.googleads.v18.common.Operand.function_operand:type_name -> google.ads.googleads.v18.common.Operand.FunctionOperand
	5, // 6: google.ads.googleads.v18.common.Operand.request_context_operand:type_name -> google.ads.googleads.v18.common.Operand.RequestContextOperand
	0, // 7: google.ads.googleads.v18.common.Operand.FunctionOperand.matching_function:type_name -> google.ads.googleads.v18.common.MatchingFunction
	7, // 8: google.ads.googleads.v18.common.Operand.RequestContextOperand.context_type:type_name -> google.ads.googleads.v18.enums.MatchingFunctionContextTypeEnum.MatchingFunctionContextType
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v18_common_matching_function_proto_init() }
func file_google_ads_googleads_v18_common_matching_function_proto_init() {
	if File_google_ads_googleads_v18_common_matching_function_proto != nil {
		return
	}
	file_google_ads_googleads_v18_common_matching_function_proto_msgTypes[0].OneofWrappers = []any{}
	file_google_ads_googleads_v18_common_matching_function_proto_msgTypes[1].OneofWrappers = []any{
		(*Operand_ConstantOperand_)(nil),
		(*Operand_FeedAttributeOperand_)(nil),
		(*Operand_FunctionOperand_)(nil),
		(*Operand_RequestContextOperand_)(nil),
	}
	file_google_ads_googleads_v18_common_matching_function_proto_msgTypes[2].OneofWrappers = []any{
		(*Operand_ConstantOperand_StringValue)(nil),
		(*Operand_ConstantOperand_LongValue)(nil),
		(*Operand_ConstantOperand_BooleanValue)(nil),
		(*Operand_ConstantOperand_DoubleValue)(nil),
	}
	file_google_ads_googleads_v18_common_matching_function_proto_msgTypes[3].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v18_common_matching_function_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v18_common_matching_function_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v18_common_matching_function_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v18_common_matching_function_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v18_common_matching_function_proto = out.File
	file_google_ads_googleads_v18_common_matching_function_proto_rawDesc = nil
	file_google_ads_googleads_v18_common_matching_function_proto_goTypes = nil
	file_google_ads_googleads_v18_common_matching_function_proto_depIdxs = nil
}
