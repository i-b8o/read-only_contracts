// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: service.proto

package pb_supreme

import (
	empty "github.com/golang/protobuf/ptypes/empty"
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
	0x0a, 0x73, 0x75, 0x70, 0x72, 0x65, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x0d, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xce, 0x03, 0x0a, 0x15, 0x53, 0x75, 0x70, 0x72,
	0x65, 0x6d, 0x65, 0x52, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x52, 0x50,
	0x43, 0x12, 0x56, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x70, 0x74,
	0x65, 0x72, 0x12, 0x20, 0x2e, 0x73, 0x75, 0x70, 0x72, 0x65, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x73, 0x75, 0x70, 0x72, 0x65, 0x6d, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x10, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x61, 0x67, 0x72, 0x61, 0x70, 0x68, 0x73, 0x12, 0x23, 0x2e,
	0x73, 0x75, 0x70, 0x72, 0x65, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x61, 0x72, 0x61, 0x67, 0x72, 0x61, 0x70, 0x68, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x5f, 0x0a, 0x10,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x23, 0x2e, 0x73, 0x75, 0x70, 0x72, 0x65, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x73, 0x75, 0x70, 0x72, 0x65, 0x6d, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x56, 0x0a,
	0x0d, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x12, 0x20,
	0x2e, 0x73, 0x75, 0x70, 0x72, 0x65, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x21, 0x2e, 0x73, 0x75, 0x70, 0x72, 0x65, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x67, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x2e, 0x73, 0x75, 0x70, 0x72,
	0x65, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x67,
	0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x17, 0x5a, 0x15, 0x73, 0x75, 0x70, 0x72,
	0x65, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x62, 0x5f, 0x73, 0x75, 0x70, 0x72, 0x65, 0x6d,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_service_proto_goTypes = []interface{}{
	(*CreateChapterRequest)(nil),     // 0: supreme.v1.CreateChapterRequest
	(*CreateParagraphsRequest)(nil),  // 1: supreme.v1.CreateParagraphsRequest
	(*CreateRegulationRequest)(nil),  // 2: supreme.v1.CreateRegulationRequest
	(*GenerateLinksRequest)(nil),     // 3: supreme.v1.GenerateLinksRequest
	(*DeleteRegulationRequest)(nil),  // 4: supreme.v1.DeleteRegulationRequest
	(*CreateChapterResponse)(nil),    // 5: supreme.v1.CreateChapterResponse
	(*empty.Empty)(nil),              // 6: google.protobuf.Empty
	(*CreateRegulationResponse)(nil), // 7: supreme.v1.CreateRegulationResponse
	(*GenerateLinksResponse)(nil),    // 8: supreme.v1.GenerateLinksResponse
}
var file_service_proto_depIdxs = []int32{
	0, // 0: supreme.v1.SupremeRegulationGRPC.CreateChapter:input_type -> supreme.v1.CreateChapterRequest
	1, // 1: supreme.v1.SupremeRegulationGRPC.CreateParagraphs:input_type -> supreme.v1.CreateParagraphsRequest
	2, // 2: supreme.v1.SupremeRegulationGRPC.CreateRegulation:input_type -> supreme.v1.CreateRegulationRequest
	3, // 3: supreme.v1.SupremeRegulationGRPC.GenerateLinks:input_type -> supreme.v1.GenerateLinksRequest
	4, // 4: supreme.v1.SupremeRegulationGRPC.DeleteRegulation:input_type -> supreme.v1.DeleteRegulationRequest
	5, // 5: supreme.v1.SupremeRegulationGRPC.CreateChapter:output_type -> supreme.v1.CreateChapterResponse
	6, // 6: supreme.v1.SupremeRegulationGRPC.CreateParagraphs:output_type -> google.protobuf.Empty
	7, // 7: supreme.v1.SupremeRegulationGRPC.CreateRegulation:output_type -> supreme.v1.CreateRegulationResponse
	8, // 8: supreme.v1.SupremeRegulationGRPC.GenerateLinks:output_type -> supreme.v1.GenerateLinksResponse
	6, // 9: supreme.v1.SupremeRegulationGRPC.DeleteRegulation:output_type -> google.protobuf.Empty
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
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
