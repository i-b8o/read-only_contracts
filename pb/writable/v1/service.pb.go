// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: service.proto

package pb_writable

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
	0x0b, 0x77, 0x72, 0x69, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x0d, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xe6, 0x08, 0x0a, 0x16,
	0x57, 0x72, 0x69, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x47, 0x52, 0x50, 0x43, 0x12, 0x61, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x2e, 0x77, 0x72, 0x69,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x67, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x25, 0x2e, 0x77, 0x72, 0x69, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x61, 0x0a, 0x10, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x2e,
	0x77, 0x72, 0x69, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x77, 0x72, 0x69, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x58, 0x0a, 0x0d,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x12, 0x21, 0x2e,
	0x77, 0x72, 0x69, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x22, 0x2e, 0x77, 0x72, 0x69, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x82, 0x01, 0x0a, 0x1b, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x73, 0x46, 0x6f, 0x72, 0x52, 0x65, 0x67, 0x75,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x2e, 0x77, 0x72, 0x69, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x68, 0x61, 0x70, 0x74,
	0x65, 0x72, 0x73, 0x46, 0x6f, 0x72, 0x52, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x77, 0x72, 0x69, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x68, 0x61, 0x70,
	0x74, 0x65, 0x72, 0x73, 0x46, 0x6f, 0x72, 0x52, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6a, 0x0a, 0x13, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x6c, 0x50, 0x61, 0x72, 0x61, 0x67, 0x72, 0x61, 0x70,
	0x68, 0x73, 0x12, 0x27, 0x2e, 0x77, 0x72, 0x69, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x6c, 0x50, 0x61, 0x72, 0x61, 0x67, 0x72,
	0x61, 0x70, 0x68, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x77, 0x72,
	0x69, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x41, 0x6c, 0x6c, 0x50, 0x61, 0x72, 0x61, 0x67, 0x72, 0x61, 0x70, 0x68, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x67, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4f, 0x6e, 0x65, 0x50, 0x61, 0x72, 0x61, 0x67, 0x72, 0x61, 0x70, 0x68, 0x12, 0x26, 0x2e,
	0x77, 0x72, 0x69, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4f, 0x6e, 0x65, 0x50, 0x61, 0x72, 0x61, 0x67, 0x72, 0x61, 0x70, 0x68, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x77, 0x72, 0x69, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x6e, 0x65, 0x50, 0x61, 0x72,
	0x61, 0x67, 0x72, 0x61, 0x70, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x7f, 0x0a, 0x1a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x72, 0x61, 0x67, 0x72,
	0x61, 0x70, 0x68, 0x73, 0x46, 0x6f, 0x72, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x12, 0x2e,
	0x2e, 0x77, 0x72, 0x69, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x50, 0x61, 0x72, 0x61, 0x67, 0x72, 0x61, 0x70, 0x68, 0x73, 0x46, 0x6f, 0x72,
	0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f,
	0x2e, 0x77, 0x72, 0x69, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x50, 0x61, 0x72, 0x61, 0x67, 0x72, 0x61, 0x70, 0x68, 0x73, 0x46, 0x6f, 0x72,
	0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x73, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72, 0x61, 0x67, 0x72, 0x61, 0x70,
	0x68, 0x73, 0x57, 0x69, 0x74, 0x68, 0x48, 0x72, 0x65, 0x66, 0x73, 0x12, 0x2a, 0x2e, 0x77, 0x72,
	0x69, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72,
	0x61, 0x67, 0x72, 0x61, 0x70, 0x68, 0x73, 0x57, 0x69, 0x74, 0x68, 0x48, 0x72, 0x65, 0x66, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x77, 0x72, 0x69, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72, 0x61, 0x67, 0x72, 0x61,
	0x70, 0x68, 0x73, 0x57, 0x69, 0x74, 0x68, 0x48, 0x72, 0x65, 0x66, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5b, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x73, 0x12, 0x22, 0x2e, 0x77, 0x72, 0x69, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x68, 0x61,
	0x70, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x77,
	0x72, 0x69, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x7f, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x67, 0x75, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x42, 0x79, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x2e, 0x2e, 0x77, 0x72, 0x69, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x42,
	0x79, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2f, 0x2e, 0x77, 0x72, 0x69, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x42,
	0x79, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x19, 0x5a, 0x17, 0x77, 0x72, 0x69, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x2f, 0x76, 0x31, 0x3b, 0x70, 0x62, 0x5f, 0x77, 0x72, 0x69, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_service_proto_goTypes = []interface{}{
	(*CreateRegulationRequest)(nil),             // 0: writable.v1.CreateRegulationRequest
	(*DeleteRegulationRequest)(nil),             // 1: writable.v1.DeleteRegulationRequest
	(*CreateChapterRequest)(nil),                // 2: writable.v1.CreateChapterRequest
	(*DeleteChaptersForRegulationRequest)(nil),  // 3: writable.v1.DeleteChaptersForRegulationRequest
	(*CreateAllParagraphsRequest)(nil),          // 4: writable.v1.CreateAllParagraphsRequest
	(*UpdateOneParagraphRequest)(nil),           // 5: writable.v1.UpdateOneParagraphRequest
	(*DeleteParagraphsForChapterRequest)(nil),   // 6: writable.v1.DeleteParagraphsForChapterRequest
	(*GetParagraphsWithHrefsRequest)(nil),       // 7: writable.v1.GetParagraphsWithHrefsRequest
	(*GetAllChaptersRequest)(nil),               // 8: writable.v1.GetAllChaptersRequest
	(*GetRegulationIdByChapterIdRequest)(nil),   // 9: writable.v1.GetRegulationIdByChapterIdRequest
	(*CreateRegulationResponse)(nil),            // 10: writable.v1.CreateRegulationResponse
	(*DeleteRegulationResponse)(nil),            // 11: writable.v1.DeleteRegulationResponse
	(*CreateChapterResponse)(nil),               // 12: writable.v1.CreateChapterResponse
	(*DeleteChaptersForRegulationResponse)(nil), // 13: writable.v1.DeleteChaptersForRegulationResponse
	(*CreateAllParagraphsResponse)(nil),         // 14: writable.v1.CreateAllParagraphsResponse
	(*UpdateOneParagraphResponse)(nil),          // 15: writable.v1.UpdateOneParagraphResponse
	(*DeleteParagraphsForChapterResponse)(nil),  // 16: writable.v1.DeleteParagraphsForChapterResponse
	(*GetParagraphsWithHrefsResponse)(nil),      // 17: writable.v1.GetParagraphsWithHrefsResponse
	(*GetAllChaptersResponse)(nil),              // 18: writable.v1.GetAllChaptersResponse
	(*GetRegulationIdByChapterIdResponse)(nil),  // 19: writable.v1.GetRegulationIdByChapterIdResponse
}
var file_service_proto_depIdxs = []int32{
	0,  // 0: writable.v1.WritableRegulationGRPC.CreateRegulation:input_type -> writable.v1.CreateRegulationRequest
	1,  // 1: writable.v1.WritableRegulationGRPC.DeleteRegulation:input_type -> writable.v1.DeleteRegulationRequest
	2,  // 2: writable.v1.WritableRegulationGRPC.CreateChapter:input_type -> writable.v1.CreateChapterRequest
	3,  // 3: writable.v1.WritableRegulationGRPC.DeleteChaptersForRegulation:input_type -> writable.v1.DeleteChaptersForRegulationRequest
	4,  // 4: writable.v1.WritableRegulationGRPC.CreateAllParagraphs:input_type -> writable.v1.CreateAllParagraphsRequest
	5,  // 5: writable.v1.WritableRegulationGRPC.UpdateOneParagraph:input_type -> writable.v1.UpdateOneParagraphRequest
	6,  // 6: writable.v1.WritableRegulationGRPC.DeleteParagraphsForChapter:input_type -> writable.v1.DeleteParagraphsForChapterRequest
	7,  // 7: writable.v1.WritableRegulationGRPC.GetParagraphsWithHrefs:input_type -> writable.v1.GetParagraphsWithHrefsRequest
	8,  // 8: writable.v1.WritableRegulationGRPC.GetAllChapters:input_type -> writable.v1.GetAllChaptersRequest
	9,  // 9: writable.v1.WritableRegulationGRPC.GetRegulationIdByChapterId:input_type -> writable.v1.GetRegulationIdByChapterIdRequest
	10, // 10: writable.v1.WritableRegulationGRPC.CreateRegulation:output_type -> writable.v1.CreateRegulationResponse
	11, // 11: writable.v1.WritableRegulationGRPC.DeleteRegulation:output_type -> writable.v1.DeleteRegulationResponse
	12, // 12: writable.v1.WritableRegulationGRPC.CreateChapter:output_type -> writable.v1.CreateChapterResponse
	13, // 13: writable.v1.WritableRegulationGRPC.DeleteChaptersForRegulation:output_type -> writable.v1.DeleteChaptersForRegulationResponse
	14, // 14: writable.v1.WritableRegulationGRPC.CreateAllParagraphs:output_type -> writable.v1.CreateAllParagraphsResponse
	15, // 15: writable.v1.WritableRegulationGRPC.UpdateOneParagraph:output_type -> writable.v1.UpdateOneParagraphResponse
	16, // 16: writable.v1.WritableRegulationGRPC.DeleteParagraphsForChapter:output_type -> writable.v1.DeleteParagraphsForChapterResponse
	17, // 17: writable.v1.WritableRegulationGRPC.GetParagraphsWithHrefs:output_type -> writable.v1.GetParagraphsWithHrefsResponse
	18, // 18: writable.v1.WritableRegulationGRPC.GetAllChapters:output_type -> writable.v1.GetAllChaptersResponse
	19, // 19: writable.v1.WritableRegulationGRPC.GetRegulationIdByChapterId:output_type -> writable.v1.GetRegulationIdByChapterIdResponse
	10, // [10:20] is the sub-list for method output_type
	0,  // [0:10] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
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