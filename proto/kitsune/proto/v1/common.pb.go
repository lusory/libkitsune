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
// source: kitsune/proto/v1/common.proto

package v1

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

type UUID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *UUID) Reset() {
	*x = UUID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kitsune_proto_v1_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UUID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UUID) ProtoMessage() {}

func (x *UUID) ProtoReflect() protoreflect.Message {
	mi := &file_kitsune_proto_v1_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UUID.ProtoReflect.Descriptor instead.
func (*UUID) Descriptor() ([]byte, []int) {
	return file_kitsune_proto_v1_common_proto_rawDescGZIP(), []int{0}
}

func (x *UUID) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type string  `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Msg  *string `protobuf:"bytes,2,opt,name=msg,proto3,oneof" json:"msg,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kitsune_proto_v1_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_kitsune_proto_v1_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_kitsune_proto_v1_common_proto_rawDescGZIP(), []int{1}
}

func (x *Error) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Error) GetMsg() string {
	if x != nil && x.Msg != nil {
		return *x.Msg
	}
	return ""
}

type MetadataMap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data map[string]string `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *MetadataMap) Reset() {
	*x = MetadataMap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kitsune_proto_v1_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetadataMap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetadataMap) ProtoMessage() {}

func (x *MetadataMap) ProtoReflect() protoreflect.Message {
	mi := &file_kitsune_proto_v1_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetadataMap.ProtoReflect.Descriptor instead.
func (*MetadataMap) Descriptor() ([]byte, []int) {
	return file_kitsune_proto_v1_common_proto_rawDescGZIP(), []int{2}
}

func (x *MetadataMap) GetData() map[string]string {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetMetadataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id *UUID `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetMetadataRequest) Reset() {
	*x = GetMetadataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kitsune_proto_v1_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMetadataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMetadataRequest) ProtoMessage() {}

func (x *GetMetadataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kitsune_proto_v1_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMetadataRequest.ProtoReflect.Descriptor instead.
func (*GetMetadataRequest) Descriptor() ([]byte, []int) {
	return file_kitsune_proto_v1_common_proto_rawDescGZIP(), []int{3}
}

func (x *GetMetadataRequest) GetId() *UUID {
	if x != nil {
		return x.Id
	}
	return nil
}

type GetMetadataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to MetadataOrError:
	//	*GetMetadataResponse_Meta
	//	*GetMetadataResponse_Error
	MetadataOrError isGetMetadataResponse_MetadataOrError `protobuf_oneof:"metadata_or_error"`
}

func (x *GetMetadataResponse) Reset() {
	*x = GetMetadataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kitsune_proto_v1_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMetadataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMetadataResponse) ProtoMessage() {}

func (x *GetMetadataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kitsune_proto_v1_common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMetadataResponse.ProtoReflect.Descriptor instead.
func (*GetMetadataResponse) Descriptor() ([]byte, []int) {
	return file_kitsune_proto_v1_common_proto_rawDescGZIP(), []int{4}
}

func (m *GetMetadataResponse) GetMetadataOrError() isGetMetadataResponse_MetadataOrError {
	if m != nil {
		return m.MetadataOrError
	}
	return nil
}

func (x *GetMetadataResponse) GetMeta() *MetadataMap {
	if x, ok := x.GetMetadataOrError().(*GetMetadataResponse_Meta); ok {
		return x.Meta
	}
	return nil
}

func (x *GetMetadataResponse) GetError() *Error {
	if x, ok := x.GetMetadataOrError().(*GetMetadataResponse_Error); ok {
		return x.Error
	}
	return nil
}

type isGetMetadataResponse_MetadataOrError interface {
	isGetMetadataResponse_MetadataOrError()
}

type GetMetadataResponse_Meta struct {
	Meta *MetadataMap `protobuf:"bytes,1,opt,name=meta,proto3,oneof"`
}

type GetMetadataResponse_Error struct {
	Error *Error `protobuf:"bytes,2,opt,name=error,proto3,oneof"`
}

func (*GetMetadataResponse_Meta) isGetMetadataResponse_MetadataOrError() {}

func (*GetMetadataResponse_Error) isGetMetadataResponse_MetadataOrError() {}

type SetMetadataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   *UUID        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Meta *MetadataMap `protobuf:"bytes,2,opt,name=meta,proto3" json:"meta,omitempty"`
}

func (x *SetMetadataRequest) Reset() {
	*x = SetMetadataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kitsune_proto_v1_common_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetMetadataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetMetadataRequest) ProtoMessage() {}

func (x *SetMetadataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kitsune_proto_v1_common_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetMetadataRequest.ProtoReflect.Descriptor instead.
func (*SetMetadataRequest) Descriptor() ([]byte, []int) {
	return file_kitsune_proto_v1_common_proto_rawDescGZIP(), []int{5}
}

func (x *SetMetadataRequest) GetId() *UUID {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *SetMetadataRequest) GetMeta() *MetadataMap {
	if x != nil {
		return x.Meta
	}
	return nil
}

type SetMetadataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error *Error `protobuf:"bytes,1,opt,name=error,proto3,oneof" json:"error,omitempty"`
}

func (x *SetMetadataResponse) Reset() {
	*x = SetMetadataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kitsune_proto_v1_common_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetMetadataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetMetadataResponse) ProtoMessage() {}

func (x *SetMetadataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kitsune_proto_v1_common_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetMetadataResponse.ProtoReflect.Descriptor instead.
func (*SetMetadataResponse) Descriptor() ([]byte, []int) {
	return file_kitsune_proto_v1_common_proto_rawDescGZIP(), []int{6}
}

func (x *SetMetadataResponse) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_kitsune_proto_v1_common_proto protoreflect.FileDescriptor

var file_kitsune_proto_v1_common_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x6b, 0x69, 0x74, 0x73, 0x75, 0x6e, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x10, 0x6b, 0x69, 0x74, 0x73, 0x75, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76,
	0x31, 0x22, 0x1c, 0x0a, 0x04, 0x55, 0x55, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0x3a, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x15, 0x0a, 0x03,
	0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x03, 0x6d, 0x73, 0x67,
	0x88, 0x01, 0x01, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x6d, 0x73, 0x67, 0x22, 0x83, 0x01, 0x0a, 0x0b,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x4d, 0x61, 0x70, 0x12, 0x3b, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x6b, 0x69, 0x74, 0x73,
	0x75, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x4d, 0x61, 0x70, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x37, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x61,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0x3c, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6b, 0x69, 0x74, 0x73, 0x75, 0x6e, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x55, 0x49, 0x44, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x90, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6b, 0x69, 0x74, 0x73, 0x75, 0x6e, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x4d, 0x61, 0x70, 0x48, 0x00, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x2f, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6b, 0x69,
	0x74, 0x73, 0x75, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x13, 0x0a,
	0x11, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6f, 0x72, 0x5f, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x22, 0x6f, 0x0a, 0x12, 0x53, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6b, 0x69, 0x74, 0x73, 0x75, 0x6e, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x55, 0x49, 0x44, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x31, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x6b, 0x69, 0x74, 0x73, 0x75, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76,
	0x31, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x4d, 0x61, 0x70, 0x52, 0x04, 0x6d,
	0x65, 0x74, 0x61, 0x22, 0x53, 0x0a, 0x13, 0x53, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6b, 0x69, 0x74, 0x73,
	0x75, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x42, 0x08,
	0x0a, 0x06, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x38, 0x0a, 0x10, 0x6b, 0x69, 0x74, 0x73,
	0x75, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x22,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x75, 0x73, 0x6f, 0x72,
	0x79, 0x2f, 0x6b, 0x69, 0x74, 0x73, 0x75, 0x6e, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kitsune_proto_v1_common_proto_rawDescOnce sync.Once
	file_kitsune_proto_v1_common_proto_rawDescData = file_kitsune_proto_v1_common_proto_rawDesc
)

func file_kitsune_proto_v1_common_proto_rawDescGZIP() []byte {
	file_kitsune_proto_v1_common_proto_rawDescOnce.Do(func() {
		file_kitsune_proto_v1_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_kitsune_proto_v1_common_proto_rawDescData)
	})
	return file_kitsune_proto_v1_common_proto_rawDescData
}

var file_kitsune_proto_v1_common_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_kitsune_proto_v1_common_proto_goTypes = []interface{}{
	(*UUID)(nil),                // 0: kitsune.proto.v1.UUID
	(*Error)(nil),               // 1: kitsune.proto.v1.Error
	(*MetadataMap)(nil),         // 2: kitsune.proto.v1.MetadataMap
	(*GetMetadataRequest)(nil),  // 3: kitsune.proto.v1.GetMetadataRequest
	(*GetMetadataResponse)(nil), // 4: kitsune.proto.v1.GetMetadataResponse
	(*SetMetadataRequest)(nil),  // 5: kitsune.proto.v1.SetMetadataRequest
	(*SetMetadataResponse)(nil), // 6: kitsune.proto.v1.SetMetadataResponse
	nil,                         // 7: kitsune.proto.v1.MetadataMap.DataEntry
}
var file_kitsune_proto_v1_common_proto_depIdxs = []int32{
	7, // 0: kitsune.proto.v1.MetadataMap.data:type_name -> kitsune.proto.v1.MetadataMap.DataEntry
	0, // 1: kitsune.proto.v1.GetMetadataRequest.id:type_name -> kitsune.proto.v1.UUID
	2, // 2: kitsune.proto.v1.GetMetadataResponse.meta:type_name -> kitsune.proto.v1.MetadataMap
	1, // 3: kitsune.proto.v1.GetMetadataResponse.error:type_name -> kitsune.proto.v1.Error
	0, // 4: kitsune.proto.v1.SetMetadataRequest.id:type_name -> kitsune.proto.v1.UUID
	2, // 5: kitsune.proto.v1.SetMetadataRequest.meta:type_name -> kitsune.proto.v1.MetadataMap
	1, // 6: kitsune.proto.v1.SetMetadataResponse.error:type_name -> kitsune.proto.v1.Error
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_kitsune_proto_v1_common_proto_init() }
func file_kitsune_proto_v1_common_proto_init() {
	if File_kitsune_proto_v1_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kitsune_proto_v1_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UUID); i {
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
		file_kitsune_proto_v1_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
		file_kitsune_proto_v1_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetadataMap); i {
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
		file_kitsune_proto_v1_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMetadataRequest); i {
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
		file_kitsune_proto_v1_common_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMetadataResponse); i {
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
		file_kitsune_proto_v1_common_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetMetadataRequest); i {
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
		file_kitsune_proto_v1_common_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetMetadataResponse); i {
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
	file_kitsune_proto_v1_common_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_kitsune_proto_v1_common_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*GetMetadataResponse_Meta)(nil),
		(*GetMetadataResponse_Error)(nil),
	}
	file_kitsune_proto_v1_common_proto_msgTypes[6].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_kitsune_proto_v1_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kitsune_proto_v1_common_proto_goTypes,
		DependencyIndexes: file_kitsune_proto_v1_common_proto_depIdxs,
		MessageInfos:      file_kitsune_proto_v1_common_proto_msgTypes,
	}.Build()
	File_kitsune_proto_v1_common_proto = out.File
	file_kitsune_proto_v1_common_proto_rawDesc = nil
	file_kitsune_proto_v1_common_proto_goTypes = nil
	file_kitsune_proto_v1_common_proto_depIdxs = nil
}
