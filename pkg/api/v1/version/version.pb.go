// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: pkg/api/v1/version/version.proto

// Copyright (c) 2018 Bhojpur Consulting Private Limited, India. All rights reserved.

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package version

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Version struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Major      uint32 `protobuf:"varint,1,opt,name=major,proto3" json:"major,omitempty"`
	Minor      uint32 `protobuf:"varint,2,opt,name=minor,proto3" json:"minor,omitempty"`
	Micro      uint32 `protobuf:"varint,3,opt,name=micro,proto3" json:"micro,omitempty"`
	Additional string `protobuf:"bytes,4,opt,name=additional,proto3" json:"additional,omitempty"`
}

func (x *Version) Reset() {
	*x = Version{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_v1_version_version_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Version) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Version) ProtoMessage() {}

func (x *Version) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_v1_version_version_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Version.ProtoReflect.Descriptor instead.
func (*Version) Descriptor() ([]byte, []int) {
	return file_pkg_api_v1_version_version_proto_rawDescGZIP(), []int{0}
}

func (x *Version) GetMajor() uint32 {
	if x != nil {
		return x.Major
	}
	return 0
}

func (x *Version) GetMinor() uint32 {
	if x != nil {
		return x.Minor
	}
	return 0
}

func (x *Version) GetMicro() uint32 {
	if x != nil {
		return x.Micro
	}
	return 0
}

func (x *Version) GetAdditional() string {
	if x != nil {
		return x.Additional
	}
	return ""
}

var File_pkg_api_v1_version_version_proto protoreflect.FileDescriptor

var file_pkg_api_v1_version_version_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x76, 0x31, 0x2e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6b, 0x0a, 0x07, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x61, 0x6a, 0x6f, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6d, 0x61, 0x6a, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05,
	0x6d, 0x69, 0x6e, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6d, 0x69, 0x6e,
	0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x05, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x64, 0x64, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x64,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x32, 0x42, 0x0a, 0x03, 0x41, 0x50, 0x49, 0x12,
	0x3b, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x13, 0x2e, 0x76, 0x31, 0x2e, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x42, 0x34, 0x5a, 0x32,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x68, 0x6f, 0x6a, 0x70,
	0x75, 0x72, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x3b, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_api_v1_version_version_proto_rawDescOnce sync.Once
	file_pkg_api_v1_version_version_proto_rawDescData = file_pkg_api_v1_version_version_proto_rawDesc
)

func file_pkg_api_v1_version_version_proto_rawDescGZIP() []byte {
	file_pkg_api_v1_version_version_proto_rawDescOnce.Do(func() {
		file_pkg_api_v1_version_version_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_api_v1_version_version_proto_rawDescData)
	})
	return file_pkg_api_v1_version_version_proto_rawDescData
}

var file_pkg_api_v1_version_version_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pkg_api_v1_version_version_proto_goTypes = []interface{}{
	(*Version)(nil),       // 0: v1.version.Version
	(*emptypb.Empty)(nil), // 1: google.protobuf.Empty
}
var file_pkg_api_v1_version_version_proto_depIdxs = []int32{
	1, // 0: v1.version.API.GetVersion:input_type -> google.protobuf.Empty
	0, // 1: v1.version.API.GetVersion:output_type -> v1.version.Version
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_api_v1_version_version_proto_init() }
func file_pkg_api_v1_version_version_proto_init() {
	if File_pkg_api_v1_version_version_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_api_v1_version_version_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Version); i {
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
			RawDescriptor: file_pkg_api_v1_version_version_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_api_v1_version_version_proto_goTypes,
		DependencyIndexes: file_pkg_api_v1_version_version_proto_depIdxs,
		MessageInfos:      file_pkg_api_v1_version_version_proto_msgTypes,
	}.Build()
	File_pkg_api_v1_version_version_proto = out.File
	file_pkg_api_v1_version_version_proto_rawDesc = nil
	file_pkg_api_v1_version_version_proto_goTypes = nil
	file_pkg_api_v1_version_version_proto_depIdxs = nil
}
