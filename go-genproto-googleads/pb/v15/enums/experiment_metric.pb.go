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
// source: google/ads/googleads/v15/enums/experiment_metric.proto

package enums

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

// The type of experiment metric.
type ExperimentMetricEnum_ExperimentMetric int32

const (
	// Not specified.
	ExperimentMetricEnum_UNSPECIFIED ExperimentMetricEnum_ExperimentMetric = 0
	// The value is unknown in this version.
	ExperimentMetricEnum_UNKNOWN ExperimentMetricEnum_ExperimentMetric = 1
	// The goal of the experiment is clicks.
	ExperimentMetricEnum_CLICKS ExperimentMetricEnum_ExperimentMetric = 2
	// The goal of the experiment is impressions.
	ExperimentMetricEnum_IMPRESSIONS ExperimentMetricEnum_ExperimentMetric = 3
	// The goal of the experiment is cost.
	ExperimentMetricEnum_COST ExperimentMetricEnum_ExperimentMetric = 4
	// The goal of the experiment is conversion rate.
	ExperimentMetricEnum_CONVERSIONS_PER_INTERACTION_RATE ExperimentMetricEnum_ExperimentMetric = 5
	// The goal of the experiment is cost per conversion.
	ExperimentMetricEnum_COST_PER_CONVERSION ExperimentMetricEnum_ExperimentMetric = 6
	// The goal of the experiment is conversion value per cost.
	ExperimentMetricEnum_CONVERSIONS_VALUE_PER_COST ExperimentMetricEnum_ExperimentMetric = 7
	// The goal of the experiment is avg cpc.
	ExperimentMetricEnum_AVERAGE_CPC ExperimentMetricEnum_ExperimentMetric = 8
	// The goal of the experiment is ctr.
	ExperimentMetricEnum_CTR ExperimentMetricEnum_ExperimentMetric = 9
	// The goal of the experiment is incremental conversions.
	ExperimentMetricEnum_INCREMENTAL_CONVERSIONS ExperimentMetricEnum_ExperimentMetric = 10
	// The goal of the experiment is completed video views.
	ExperimentMetricEnum_COMPLETED_VIDEO_VIEWS ExperimentMetricEnum_ExperimentMetric = 11
	// The goal of the experiment is custom algorithms.
	ExperimentMetricEnum_CUSTOM_ALGORITHMS ExperimentMetricEnum_ExperimentMetric = 12
	// The goal of the experiment is conversions.
	ExperimentMetricEnum_CONVERSIONS ExperimentMetricEnum_ExperimentMetric = 13
	// The goal of the experiment is conversion value.
	ExperimentMetricEnum_CONVERSION_VALUE ExperimentMetricEnum_ExperimentMetric = 14
)

// Enum value maps for ExperimentMetricEnum_ExperimentMetric.
var (
	ExperimentMetricEnum_ExperimentMetric_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "CLICKS",
		3:  "IMPRESSIONS",
		4:  "COST",
		5:  "CONVERSIONS_PER_INTERACTION_RATE",
		6:  "COST_PER_CONVERSION",
		7:  "CONVERSIONS_VALUE_PER_COST",
		8:  "AVERAGE_CPC",
		9:  "CTR",
		10: "INCREMENTAL_CONVERSIONS",
		11: "COMPLETED_VIDEO_VIEWS",
		12: "CUSTOM_ALGORITHMS",
		13: "CONVERSIONS",
		14: "CONVERSION_VALUE",
	}
	ExperimentMetricEnum_ExperimentMetric_value = map[string]int32{
		"UNSPECIFIED":                      0,
		"UNKNOWN":                          1,
		"CLICKS":                           2,
		"IMPRESSIONS":                      3,
		"COST":                             4,
		"CONVERSIONS_PER_INTERACTION_RATE": 5,
		"COST_PER_CONVERSION":              6,
		"CONVERSIONS_VALUE_PER_COST":       7,
		"AVERAGE_CPC":                      8,
		"CTR":                              9,
		"INCREMENTAL_CONVERSIONS":          10,
		"COMPLETED_VIDEO_VIEWS":            11,
		"CUSTOM_ALGORITHMS":                12,
		"CONVERSIONS":                      13,
		"CONVERSION_VALUE":                 14,
	}
)

func (x ExperimentMetricEnum_ExperimentMetric) Enum() *ExperimentMetricEnum_ExperimentMetric {
	p := new(ExperimentMetricEnum_ExperimentMetric)
	*p = x
	return p
}

func (x ExperimentMetricEnum_ExperimentMetric) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ExperimentMetricEnum_ExperimentMetric) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v15_enums_experiment_metric_proto_enumTypes[0].Descriptor()
}

func (ExperimentMetricEnum_ExperimentMetric) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v15_enums_experiment_metric_proto_enumTypes[0]
}

func (x ExperimentMetricEnum_ExperimentMetric) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ExperimentMetricEnum_ExperimentMetric.Descriptor instead.
func (ExperimentMetricEnum_ExperimentMetric) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v15_enums_experiment_metric_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing the type of experiment metric.
type ExperimentMetricEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ExperimentMetricEnum) Reset() {
	*x = ExperimentMetricEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v15_enums_experiment_metric_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExperimentMetricEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExperimentMetricEnum) ProtoMessage() {}

func (x *ExperimentMetricEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v15_enums_experiment_metric_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExperimentMetricEnum.ProtoReflect.Descriptor instead.
func (*ExperimentMetricEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v15_enums_experiment_metric_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v15_enums_experiment_metric_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v15_enums_experiment_metric_proto_rawDesc = []byte{
	0x0a, 0x36, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x35, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x35, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x22, 0xdf, 0x02, 0x0a, 0x14, 0x45, 0x78, 0x70,
	0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x45, 0x6e, 0x75,
	0x6d, 0x22, 0xc6, 0x02, 0x0a, 0x10, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x4c, 0x49, 0x43, 0x4b, 0x53, 0x10, 0x02,
	0x12, 0x0f, 0x0a, 0x0b, 0x49, 0x4d, 0x50, 0x52, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x53, 0x10,
	0x03, 0x12, 0x08, 0x0a, 0x04, 0x43, 0x4f, 0x53, 0x54, 0x10, 0x04, 0x12, 0x24, 0x0a, 0x20, 0x43,
	0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x53, 0x5f, 0x50, 0x45, 0x52, 0x5f, 0x49,
	0x4e, 0x54, 0x45, 0x52, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x41, 0x54, 0x45, 0x10,
	0x05, 0x12, 0x17, 0x0a, 0x13, 0x43, 0x4f, 0x53, 0x54, 0x5f, 0x50, 0x45, 0x52, 0x5f, 0x43, 0x4f,
	0x4e, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x10, 0x06, 0x12, 0x1e, 0x0a, 0x1a, 0x43, 0x4f,
	0x4e, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x53, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f,
	0x50, 0x45, 0x52, 0x5f, 0x43, 0x4f, 0x53, 0x54, 0x10, 0x07, 0x12, 0x0f, 0x0a, 0x0b, 0x41, 0x56,
	0x45, 0x52, 0x41, 0x47, 0x45, 0x5f, 0x43, 0x50, 0x43, 0x10, 0x08, 0x12, 0x07, 0x0a, 0x03, 0x43,
	0x54, 0x52, 0x10, 0x09, 0x12, 0x1b, 0x0a, 0x17, 0x49, 0x4e, 0x43, 0x52, 0x45, 0x4d, 0x45, 0x4e,
	0x54, 0x41, 0x4c, 0x5f, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x53, 0x10,
	0x0a, 0x12, 0x19, 0x0a, 0x15, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x5f, 0x56,
	0x49, 0x44, 0x45, 0x4f, 0x5f, 0x56, 0x49, 0x45, 0x57, 0x53, 0x10, 0x0b, 0x12, 0x15, 0x0a, 0x11,
	0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x5f, 0x41, 0x4c, 0x47, 0x4f, 0x52, 0x49, 0x54, 0x48, 0x4d,
	0x53, 0x10, 0x0c, 0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f,
	0x4e, 0x53, 0x10, 0x0d, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x49,
	0x4f, 0x4e, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x10, 0x0e, 0x42, 0xef, 0x01, 0x0a, 0x22, 0x63,
	0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x42, 0x15, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f,
	0x76, 0x31, 0x35, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x3b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0xa2,
	0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41,
	0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x35,
	0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c,
	0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31,
	0x35, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73,
	0x3a, 0x3a, 0x56, 0x31, 0x35, 0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v15_enums_experiment_metric_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v15_enums_experiment_metric_proto_rawDescData = file_google_ads_googleads_v15_enums_experiment_metric_proto_rawDesc
)

func file_google_ads_googleads_v15_enums_experiment_metric_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v15_enums_experiment_metric_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v15_enums_experiment_metric_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v15_enums_experiment_metric_proto_rawDescData)
	})
	return file_google_ads_googleads_v15_enums_experiment_metric_proto_rawDescData
}

var file_google_ads_googleads_v15_enums_experiment_metric_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v15_enums_experiment_metric_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v15_enums_experiment_metric_proto_goTypes = []interface{}{
	(ExperimentMetricEnum_ExperimentMetric)(0), // 0: google.ads.googleads.v15.enums.ExperimentMetricEnum.ExperimentMetric
	(*ExperimentMetricEnum)(nil),               // 1: google.ads.googleads.v15.enums.ExperimentMetricEnum
}
var file_google_ads_googleads_v15_enums_experiment_metric_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v15_enums_experiment_metric_proto_init() }
func file_google_ads_googleads_v15_enums_experiment_metric_proto_init() {
	if File_google_ads_googleads_v15_enums_experiment_metric_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v15_enums_experiment_metric_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExperimentMetricEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v15_enums_experiment_metric_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v15_enums_experiment_metric_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v15_enums_experiment_metric_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v15_enums_experiment_metric_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v15_enums_experiment_metric_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v15_enums_experiment_metric_proto = out.File
	file_google_ads_googleads_v15_enums_experiment_metric_proto_rawDesc = nil
	file_google_ads_googleads_v15_enums_experiment_metric_proto_goTypes = nil
	file_google_ads_googleads_v15_enums_experiment_metric_proto_depIdxs = nil
}
