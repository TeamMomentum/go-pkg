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
// source: google/ads/googleads/v19/resources/asset_set_asset.proto

package resources

import (
	enums "github.com/TeamMomentum/go-pkg/go-genproto-googleads/pb/v19/enums"
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

// AssetSetAsset is the link between an asset and an asset set.
// Adding an AssetSetAsset links an asset with an asset set.
type AssetSetAsset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Immutable. The resource name of the asset set asset.
	// Asset set asset resource names have the form:
	//
	// `customers/{customer_id}/assetSetAssets/{asset_set_id}~{asset_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Immutable. The asset set which this asset set asset is linking to.
	AssetSet string `protobuf:"bytes,2,opt,name=asset_set,json=assetSet,proto3" json:"asset_set,omitempty"`
	// Immutable. The asset which this asset set asset is linking to.
	Asset string `protobuf:"bytes,3,opt,name=asset,proto3" json:"asset,omitempty"`
	// Output only. The status of the asset set asset. Read-only.
	Status enums.AssetSetAssetStatusEnum_AssetSetAssetStatus `protobuf:"varint,4,opt,name=status,proto3,enum=google.ads.googleads.v19.enums.AssetSetAssetStatusEnum_AssetSetAssetStatus" json:"status,omitempty"`
}

func (x *AssetSetAsset) Reset() {
	*x = AssetSetAsset{}
	mi := &file_google_ads_googleads_v19_resources_asset_set_asset_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AssetSetAsset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssetSetAsset) ProtoMessage() {}

func (x *AssetSetAsset) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v19_resources_asset_set_asset_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssetSetAsset.ProtoReflect.Descriptor instead.
func (*AssetSetAsset) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v19_resources_asset_set_asset_proto_rawDescGZIP(), []int{0}
}

func (x *AssetSetAsset) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *AssetSetAsset) GetAssetSet() string {
	if x != nil {
		return x.AssetSet
	}
	return ""
}

func (x *AssetSetAsset) GetAsset() string {
	if x != nil {
		return x.Asset
	}
	return ""
}

func (x *AssetSetAsset) GetStatus() enums.AssetSetAssetStatusEnum_AssetSetAssetStatus {
	if x != nil {
		return x.Status
	}
	return enums.AssetSetAssetStatusEnum_AssetSetAssetStatus(0)
}

var File_google_ads_googleads_v19_resources_asset_set_asset_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v19_resources_asset_set_asset_proto_rawDesc = []byte{
	0x0a, 0x38, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x39, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x74, 0x5f, 0x61,
	0x73, 0x73, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x31, 0x39, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0x3b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x39, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x61,
	0x73, 0x73, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x74, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65,
	0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc3, 0x03, 0x0a, 0x0d, 0x41, 0x73, 0x73, 0x65,
	0x74, 0x53, 0x65, 0x74, 0x41, 0x73, 0x73, 0x65, 0x74, 0x12, 0x53, 0x0a, 0x0d, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x2e, 0xe0, 0x41, 0x05, 0xfa, 0x41, 0x28, 0x0a, 0x26, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x41, 0x73, 0x73, 0x65, 0x74, 0x53, 0x65, 0x74, 0x41, 0x73, 0x73, 0x65, 0x74,
	0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x46,
	0x0a, 0x09, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x29, 0xe0, 0x41, 0x05, 0xfa, 0x41, 0x23, 0x0a, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x73, 0x73, 0x65, 0x74, 0x53, 0x65, 0x74, 0x52, 0x08, 0x61, 0x73,
	0x73, 0x65, 0x74, 0x53, 0x65, 0x74, 0x12, 0x3c, 0x0a, 0x05, 0x61, 0x73, 0x73, 0x65, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x26, 0xe0, 0x41, 0x05, 0xfa, 0x41, 0x20, 0x0a, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x73, 0x73, 0x65, 0x74, 0x52, 0x05, 0x61,
	0x73, 0x73, 0x65, 0x74, 0x12, 0x68, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x4b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x53, 0x65, 0x74, 0x41, 0x73,
	0x73, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x41, 0x73,
	0x73, 0x65, 0x74, 0x53, 0x65, 0x74, 0x41, 0x73, 0x73, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x3a, 0x6d,
	0xea, 0x41, 0x6a, 0x0a, 0x26, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x73,
	0x73, 0x65, 0x74, 0x53, 0x65, 0x74, 0x41, 0x73, 0x73, 0x65, 0x74, 0x12, 0x40, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x53, 0x65, 0x74, 0x41, 0x73, 0x73,
	0x65, 0x74, 0x73, 0x2f, 0x7b, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x74, 0x5f, 0x69,
	0x64, 0x7d, 0x7e, 0x7b, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x84, 0x02,
	0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0x12, 0x41, 0x73, 0x73, 0x65, 0x74, 0x53,
	0x65, 0x74, 0x41, 0x73, 0x73, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x39, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x3b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41,
	0x41, 0xaa, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x39, 0x2e, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xca, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c,
	0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31,
	0x39, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xea, 0x02, 0x26, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x39, 0x3a, 0x3a, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v19_resources_asset_set_asset_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v19_resources_asset_set_asset_proto_rawDescData = file_google_ads_googleads_v19_resources_asset_set_asset_proto_rawDesc
)

func file_google_ads_googleads_v19_resources_asset_set_asset_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v19_resources_asset_set_asset_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v19_resources_asset_set_asset_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v19_resources_asset_set_asset_proto_rawDescData)
	})
	return file_google_ads_googleads_v19_resources_asset_set_asset_proto_rawDescData
}

var file_google_ads_googleads_v19_resources_asset_set_asset_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v19_resources_asset_set_asset_proto_goTypes = []any{
	(*AssetSetAsset)(nil), // 0: google.ads.googleads.v19.resources.AssetSetAsset
	(enums.AssetSetAssetStatusEnum_AssetSetAssetStatus)(0), // 1: google.ads.googleads.v19.enums.AssetSetAssetStatusEnum.AssetSetAssetStatus
}
var file_google_ads_googleads_v19_resources_asset_set_asset_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v19.resources.AssetSetAsset.status:type_name -> google.ads.googleads.v19.enums.AssetSetAssetStatusEnum.AssetSetAssetStatus
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v19_resources_asset_set_asset_proto_init() }
func file_google_ads_googleads_v19_resources_asset_set_asset_proto_init() {
	if File_google_ads_googleads_v19_resources_asset_set_asset_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v19_resources_asset_set_asset_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v19_resources_asset_set_asset_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v19_resources_asset_set_asset_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v19_resources_asset_set_asset_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v19_resources_asset_set_asset_proto = out.File
	file_google_ads_googleads_v19_resources_asset_set_asset_proto_rawDesc = nil
	file_google_ads_googleads_v19_resources_asset_set_asset_proto_goTypes = nil
	file_google_ads_googleads_v19_resources_asset_set_asset_proto_depIdxs = nil
}
