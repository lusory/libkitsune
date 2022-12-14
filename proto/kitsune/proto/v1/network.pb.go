//
// This file is part of kitsune, licensed under the Apache License, Version 2.0 (the "License").
//
// Copyright (c) 2022-present lusory contributors
//
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.4
// source: kitsune/proto/v1/network.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type NetworkAddressFamily int32

const (
	NetworkAddressFamily_IPV4    NetworkAddressFamily = 0
	NetworkAddressFamily_IPV6    NetworkAddressFamily = 1
	NetworkAddressFamily_UNIX    NetworkAddressFamily = 2
	NetworkAddressFamily_VSOCK   NetworkAddressFamily = 3
	NetworkAddressFamily_UNKNOWN NetworkAddressFamily = 4
)

// Enum value maps for NetworkAddressFamily.
var (
	NetworkAddressFamily_name = map[int32]string{
		0: "IPV4",
		1: "IPV6",
		2: "UNIX",
		3: "VSOCK",
		4: "UNKNOWN",
	}
	NetworkAddressFamily_value = map[string]int32{
		"IPV4":    0,
		"IPV6":    1,
		"UNIX":    2,
		"VSOCK":   3,
		"UNKNOWN": 4,
	}
)

func (x NetworkAddressFamily) Enum() *NetworkAddressFamily {
	p := new(NetworkAddressFamily)
	*p = x
	return p
}

func (x NetworkAddressFamily) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NetworkAddressFamily) Descriptor() protoreflect.EnumDescriptor {
	return file_kitsune_proto_v1_network_proto_enumTypes[0].Descriptor()
}

func (NetworkAddressFamily) Type() protoreflect.EnumType {
	return &file_kitsune_proto_v1_network_proto_enumTypes[0]
}

func (x NetworkAddressFamily) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NetworkAddressFamily.Descriptor instead.
func (NetworkAddressFamily) EnumDescriptor() ([]byte, []int) {
	return file_kitsune_proto_v1_network_proto_rawDescGZIP(), []int{0}
}

var File_kitsune_proto_v1_network_proto protoreflect.FileDescriptor

var file_kitsune_proto_v1_network_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x6b, 0x69, 0x74, 0x73, 0x75, 0x6e, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x76, 0x31, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x10, 0x6b, 0x69, 0x74, 0x73, 0x75, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1d, 0x6b, 0x69, 0x74, 0x73, 0x75, 0x6e, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x4c,
	0x0a, 0x14, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x50, 0x56, 0x34, 0x10, 0x00,
	0x12, 0x08, 0x0a, 0x04, 0x49, 0x50, 0x56, 0x36, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x55, 0x4e,
	0x49, 0x58, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x56, 0x53, 0x4f, 0x43, 0x4b, 0x10, 0x03, 0x12,
	0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x04, 0x42, 0x38, 0x0a, 0x10,
	0x6b, 0x69, 0x74, 0x73, 0x75, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31,
	0x50, 0x01, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c,
	0x75, 0x73, 0x6f, 0x72, 0x79, 0x2f, 0x6b, 0x69, 0x74, 0x73, 0x75, 0x6e, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kitsune_proto_v1_network_proto_rawDescOnce sync.Once
	file_kitsune_proto_v1_network_proto_rawDescData = file_kitsune_proto_v1_network_proto_rawDesc
)

func file_kitsune_proto_v1_network_proto_rawDescGZIP() []byte {
	file_kitsune_proto_v1_network_proto_rawDescOnce.Do(func() {
		file_kitsune_proto_v1_network_proto_rawDescData = protoimpl.X.CompressGZIP(file_kitsune_proto_v1_network_proto_rawDescData)
	})
	return file_kitsune_proto_v1_network_proto_rawDescData
}

var file_kitsune_proto_v1_network_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_kitsune_proto_v1_network_proto_goTypes = []interface{}{
	(NetworkAddressFamily)(0), // 0: kitsune.proto.v1.NetworkAddressFamily
}
var file_kitsune_proto_v1_network_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_kitsune_proto_v1_network_proto_init() }
func file_kitsune_proto_v1_network_proto_init() {
	if File_kitsune_proto_v1_network_proto != nil {
		return
	}
	file_kitsune_proto_v1_common_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_kitsune_proto_v1_network_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kitsune_proto_v1_network_proto_goTypes,
		DependencyIndexes: file_kitsune_proto_v1_network_proto_depIdxs,
		EnumInfos:         file_kitsune_proto_v1_network_proto_enumTypes,
	}.Build()
	File_kitsune_proto_v1_network_proto = out.File
	file_kitsune_proto_v1_network_proto_rawDesc = nil
	file_kitsune_proto_v1_network_proto_goTypes = nil
	file_kitsune_proto_v1_network_proto_depIdxs = nil
}
