// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: tagger.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var file_tagger_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         847939,
		Name:          "tagger.tags",
		Tag:           "bytes,847939,opt,name=tags",
		Filename:      "tagger.proto",
	},
	{
		ExtendedType:  (*descriptorpb.OneofOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         847939,
		Name:          "tagger.oneof_tags",
		Tag:           "bytes,847939,opt,name=oneof_tags",
		Filename:      "tagger.proto",
	},
}

// Extension fields to descriptorpb.FieldOptions.
var (
	// Multiple Tags can be spcified.
	//
	// optional string tags = 847939;
	E_Tags = &file_tagger_proto_extTypes[0]
)

// Extension fields to descriptorpb.OneofOptions.
var (
	// Multiple Tags can be spcified.
	//
	// optional string oneof_tags = 847939;
	E_OneofTags = &file_tagger_proto_extTypes[1]
)

var File_tagger_proto protoreflect.FileDescriptor

var file_tagger_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x74, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x74, 0x61, 0x67, 0x67, 0x65, 0x72, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3a, 0x33, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73,
	0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0xc3, 0xe0, 0x33, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x3a, 0x3e, 0x0a,
	0x0a, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x5f, 0x74, 0x61, 0x67, 0x73, 0x12, 0x1d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4f, 0x6e,
	0x65, 0x6f, 0x66, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xc3, 0xe0, 0x33, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x54, 0x61, 0x67, 0x73, 0x42, 0x2b, 0x5a,
	0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x76, 0x69,
	0x64, 0x61, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var file_tagger_proto_goTypes = []interface{}{
	(*descriptorpb.FieldOptions)(nil), // 0: google.protobuf.FieldOptions
	(*descriptorpb.OneofOptions)(nil), // 1: google.protobuf.OneofOptions
}
var file_tagger_proto_depIdxs = []int32{
	0, // 0: tagger.tags:extendee -> google.protobuf.FieldOptions
	1, // 1: tagger.oneof_tags:extendee -> google.protobuf.OneofOptions
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	0, // [0:2] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tagger_proto_init() }
func file_tagger_proto_init() {
	if File_tagger_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tagger_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 2,
			NumServices:   0,
		},
		GoTypes:           file_tagger_proto_goTypes,
		DependencyIndexes: file_tagger_proto_depIdxs,
		ExtensionInfos:    file_tagger_proto_extTypes,
	}.Build()
	File_tagger_proto = out.File
	file_tagger_proto_rawDesc = nil
	file_tagger_proto_goTypes = nil
	file_tagger_proto_depIdxs = nil
}
