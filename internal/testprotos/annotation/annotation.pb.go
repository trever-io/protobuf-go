// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/testprotos/annotation/annotation.proto

package annotation

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	unsafe "unsafe"
)

var file_internal_testprotos_annotation_annotation_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         37383685,
		Name:          "go_annotation.track_field_use",
		Tag:           "varint,37383685,opt,name=track_field_use",
		Filename:      "internal/testprotos/annotation/annotation.proto",
	},
}

// Extension fields to descriptorpb.MessageOptions.
var (
	// Setting this on a message enables tracking of which fields in the message
	// a specific binary might access. As a consequence, it also disables the use
	// of the message accessor methods to satisfy interfaces: they can only be
	// called directly.
	//
	// optional bool track_field_use = 37383685;
	E_TrackFieldUse = &file_internal_testprotos_annotation_annotation_proto_extTypes[0]
)

var File_internal_testprotos_annotation_annotation_proto protoreflect.FileDescriptor

var file_internal_testprotos_annotation_annotation_proto_rawDesc = string([]byte{
	0x0a, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0d, 0x67, 0x6f, 0x5f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x3a, 0x4a, 0x0a, 0x0f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x5f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x75, 0x73, 0x65, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x85, 0xdc, 0xe9, 0x11, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0d, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x55, 0x73, 0x65, 0x42, 0x3b,
	0x5a, 0x39, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e,
	0x6f, 0x72, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
})

var file_internal_testprotos_annotation_annotation_proto_goTypes = []any{
	(*descriptorpb.MessageOptions)(nil), // 0: google.protobuf.MessageOptions
}
var file_internal_testprotos_annotation_annotation_proto_depIdxs = []int32{
	0, // 0: go_annotation.track_field_use:extendee -> google.protobuf.MessageOptions
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	0, // [0:1] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_testprotos_annotation_annotation_proto_init() }
func file_internal_testprotos_annotation_annotation_proto_init() {
	if File_internal_testprotos_annotation_annotation_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_internal_testprotos_annotation_annotation_proto_rawDesc), len(file_internal_testprotos_annotation_annotation_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_internal_testprotos_annotation_annotation_proto_goTypes,
		DependencyIndexes: file_internal_testprotos_annotation_annotation_proto_depIdxs,
		ExtensionInfos:    file_internal_testprotos_annotation_annotation_proto_extTypes,
	}.Build()
	File_internal_testprotos_annotation_annotation_proto = out.File
	file_internal_testprotos_annotation_annotation_proto_goTypes = nil
	file_internal_testprotos_annotation_annotation_proto_depIdxs = nil
}
