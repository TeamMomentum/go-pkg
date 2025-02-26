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
// source: google/ads/googleads/v18/enums/video_thumbnail.proto

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

// Enum listing the possible types of a video thumbnail.
type VideoThumbnailEnum_VideoThumbnail int32

const (
	// The type has not been specified.
	VideoThumbnailEnum_UNSPECIFIED VideoThumbnailEnum_VideoThumbnail = 0
	// The received value is not known in this version.
	// This is a response-only value.
	VideoThumbnailEnum_UNKNOWN VideoThumbnailEnum_VideoThumbnail = 1
	// The default thumbnail. Can be auto-generated or user-uploaded.
	VideoThumbnailEnum_DEFAULT_THUMBNAIL VideoThumbnailEnum_VideoThumbnail = 2
	// Thumbnail 1, generated from the video.
	VideoThumbnailEnum_THUMBNAIL_1 VideoThumbnailEnum_VideoThumbnail = 3
	// Thumbnail 2, generated from the video.
	VideoThumbnailEnum_THUMBNAIL_2 VideoThumbnailEnum_VideoThumbnail = 4
	// Thumbnail 3, generated from the video.
	VideoThumbnailEnum_THUMBNAIL_3 VideoThumbnailEnum_VideoThumbnail = 5
)

// Enum value maps for VideoThumbnailEnum_VideoThumbnail.
var (
	VideoThumbnailEnum_VideoThumbnail_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "DEFAULT_THUMBNAIL",
		3: "THUMBNAIL_1",
		4: "THUMBNAIL_2",
		5: "THUMBNAIL_3",
	}
	VideoThumbnailEnum_VideoThumbnail_value = map[string]int32{
		"UNSPECIFIED":       0,
		"UNKNOWN":           1,
		"DEFAULT_THUMBNAIL": 2,
		"THUMBNAIL_1":       3,
		"THUMBNAIL_2":       4,
		"THUMBNAIL_3":       5,
	}
)

func (x VideoThumbnailEnum_VideoThumbnail) Enum() *VideoThumbnailEnum_VideoThumbnail {
	p := new(VideoThumbnailEnum_VideoThumbnail)
	*p = x
	return p
}

func (x VideoThumbnailEnum_VideoThumbnail) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VideoThumbnailEnum_VideoThumbnail) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v18_enums_video_thumbnail_proto_enumTypes[0].Descriptor()
}

func (VideoThumbnailEnum_VideoThumbnail) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v18_enums_video_thumbnail_proto_enumTypes[0]
}

func (x VideoThumbnailEnum_VideoThumbnail) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VideoThumbnailEnum_VideoThumbnail.Descriptor instead.
func (VideoThumbnailEnum_VideoThumbnail) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v18_enums_video_thumbnail_proto_rawDescGZIP(), []int{0, 0}
}

// Defines the thumbnail to use for In-Display video ads. Note that
// DEFAULT_THUMBNAIL may have been uploaded by the user while thumbnails 1-3 are
// auto-generated from the video.
type VideoThumbnailEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *VideoThumbnailEnum) Reset() {
	*x = VideoThumbnailEnum{}
	mi := &file_google_ads_googleads_v18_enums_video_thumbnail_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VideoThumbnailEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoThumbnailEnum) ProtoMessage() {}

func (x *VideoThumbnailEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v18_enums_video_thumbnail_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoThumbnailEnum.ProtoReflect.Descriptor instead.
func (*VideoThumbnailEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v18_enums_video_thumbnail_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v18_enums_video_thumbnail_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v18_enums_video_thumbnail_proto_rawDesc = []byte{
	0x0a, 0x34, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x38, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x38,
	0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x22, 0x8e, 0x01, 0x0a, 0x12, 0x56, 0x69, 0x64, 0x65, 0x6f,
	0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0x78, 0x0a,
	0x0e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x12,
	0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x15, 0x0a,
	0x11, 0x44, 0x45, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x5f, 0x54, 0x48, 0x55, 0x4d, 0x42, 0x4e, 0x41,
	0x49, 0x4c, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x54, 0x48, 0x55, 0x4d, 0x42, 0x4e, 0x41, 0x49,
	0x4c, 0x5f, 0x31, 0x10, 0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x54, 0x48, 0x55, 0x4d, 0x42, 0x4e, 0x41,
	0x49, 0x4c, 0x5f, 0x32, 0x10, 0x04, 0x12, 0x0f, 0x0a, 0x0b, 0x54, 0x48, 0x55, 0x4d, 0x42, 0x4e,
	0x41, 0x49, 0x4c, 0x5f, 0x33, 0x10, 0x05, 0x42, 0xed, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x38, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x42, 0x13,
	0x56, 0x69, 0x64, 0x65, 0x6f, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f,
	0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x38, 0x2f, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x3b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41,
	0xaa, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x38, 0x2e, 0x45, 0x6e, 0x75, 0x6d,
	0x73, 0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x38, 0x5c, 0x45, 0x6e, 0x75,
	0x6d, 0x73, 0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73,
	0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x38,
	0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v18_enums_video_thumbnail_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v18_enums_video_thumbnail_proto_rawDescData = file_google_ads_googleads_v18_enums_video_thumbnail_proto_rawDesc
)

func file_google_ads_googleads_v18_enums_video_thumbnail_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v18_enums_video_thumbnail_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v18_enums_video_thumbnail_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v18_enums_video_thumbnail_proto_rawDescData)
	})
	return file_google_ads_googleads_v18_enums_video_thumbnail_proto_rawDescData
}

var file_google_ads_googleads_v18_enums_video_thumbnail_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v18_enums_video_thumbnail_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v18_enums_video_thumbnail_proto_goTypes = []any{
	(VideoThumbnailEnum_VideoThumbnail)(0), // 0: google.ads.googleads.v18.enums.VideoThumbnailEnum.VideoThumbnail
	(*VideoThumbnailEnum)(nil),             // 1: google.ads.googleads.v18.enums.VideoThumbnailEnum
}
var file_google_ads_googleads_v18_enums_video_thumbnail_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v18_enums_video_thumbnail_proto_init() }
func file_google_ads_googleads_v18_enums_video_thumbnail_proto_init() {
	if File_google_ads_googleads_v18_enums_video_thumbnail_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v18_enums_video_thumbnail_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v18_enums_video_thumbnail_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v18_enums_video_thumbnail_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v18_enums_video_thumbnail_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v18_enums_video_thumbnail_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v18_enums_video_thumbnail_proto = out.File
	file_google_ads_googleads_v18_enums_video_thumbnail_proto_rawDesc = nil
	file_google_ads_googleads_v18_enums_video_thumbnail_proto_goTypes = nil
	file_google_ads_googleads_v18_enums_video_thumbnail_proto_depIdxs = nil
}
