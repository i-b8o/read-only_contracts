// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: service.proto

package pb_reader

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_service_proto protoreflect.FileDescriptor

var file_service_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x09, 0x72, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x0d, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x88, 0x02, 0x0a, 0x0a, 0x52, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x47, 0x52, 0x50, 0x43, 0x12, 0x54, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x67, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x2e, 0x72, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x72, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x67, 0x75, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4b,
	0x0a, 0x0a, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x12, 0x1c, 0x2e, 0x72,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x70,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x72, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x57, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x73, 0x12, 0x20, 0x2e,
	0x72, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x21, 0x2e, 0x72, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x15, 0x5a, 0x13, 0x72, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x3b, 0x70, 0x62, 0x5f, 0x72, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var file_service_proto_goTypes = []interface{}{
	(*GetRegulationRequest)(nil),   // 0: reader.v1.GetRegulationRequest
	(*GetChapterRequest)(nil),      // 1: reader.v1.GetChapterRequest
	(*GetAllChaptersRequest)(nil),  // 2: reader.v1.GetAllChaptersRequest
	(*GetRegulationResponse)(nil),  // 3: reader.v1.GetRegulationResponse
	(*GetChapterResponse)(nil),     // 4: reader.v1.GetChapterResponse
	(*GetAllChaptersResponse)(nil), // 5: reader.v1.GetAllChaptersResponse
}
var file_service_proto_depIdxs = []int32{
	0, // 0: reader.v1.ReaderGRPC.GetRegulation:input_type -> reader.v1.GetRegulationRequest
	1, // 1: reader.v1.ReaderGRPC.GetChapter:input_type -> reader.v1.GetChapterRequest
	2, // 2: reader.v1.ReaderGRPC.GetAllChapters:input_type -> reader.v1.GetAllChaptersRequest
	3, // 3: reader.v1.ReaderGRPC.GetRegulation:output_type -> reader.v1.GetRegulationResponse
	4, // 4: reader.v1.ReaderGRPC.GetChapter:output_type -> reader.v1.GetChapterResponse
	5, // 5: reader.v1.ReaderGRPC.GetAllChapters:output_type -> reader.v1.GetAllChaptersResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_service_proto_init() }
func file_service_proto_init() {
	if File_service_proto != nil {
		return
	}
	file_message_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_proto_goTypes,
		DependencyIndexes: file_service_proto_depIdxs,
	}.Build()
	File_service_proto = out.File
	file_service_proto_rawDesc = nil
	file_service_proto_goTypes = nil
	file_service_proto_depIdxs = nil
}
