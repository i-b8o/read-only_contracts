// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: master_service.proto

package pb_master

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_master_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_master_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_master_service_proto_rawDescGZIP(), []int{0}
}

var File_master_service_proto protoreflect.FileDescriptor

var file_master_service_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x1a, 0x14, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x32, 0xaf, 0x03, 0x0a, 0x0a, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x47, 0x52, 0x50, 0x43, 0x12,
	0x54, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72,
	0x12, 0x1f, 0x2e, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x20, 0x2e, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x61, 0x72, 0x61, 0x67, 0x72, 0x61, 0x70, 0x68, 0x73, 0x12, 0x22, 0x2e, 0x6d, 0x61, 0x73, 0x74,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x61,
	0x67, 0x72, 0x61, 0x70, 0x68, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e,
	0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x00, 0x12, 0x5d, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x67, 0x75, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x2e, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x6d, 0x61, 0x73, 0x74,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x67, 0x75,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x54, 0x0a, 0x0d, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6e, 0x6b,
	0x73, 0x12, 0x1f, 0x2e, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x20, 0x2e, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x2e, 0x6d, 0x61, 0x73,
	0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x67,
	0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10,
	0x2e, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x00, 0x42, 0x15, 0x5a, 0x13, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b,
	0x70, 0x62, 0x5f, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_master_service_proto_rawDescOnce sync.Once
	file_master_service_proto_rawDescData = file_master_service_proto_rawDesc
)

func file_master_service_proto_rawDescGZIP() []byte {
	file_master_service_proto_rawDescOnce.Do(func() {
		file_master_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_master_service_proto_rawDescData)
	})
	return file_master_service_proto_rawDescData
}

var file_master_service_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_master_service_proto_goTypes = []interface{}{
	(*Empty)(nil),                    // 0: master.v1.Empty
	(*CreateChapterRequest)(nil),     // 1: master.v1.CreateChapterRequest
	(*CreateParagraphsRequest)(nil),  // 2: master.v1.CreateParagraphsRequest
	(*CreateRegulationRequest)(nil),  // 3: master.v1.CreateRegulationRequest
	(*GenerateLinksRequest)(nil),     // 4: master.v1.GenerateLinksRequest
	(*DeleteRegulationRequest)(nil),  // 5: master.v1.DeleteRegulationRequest
	(*CreateChapterResponse)(nil),    // 6: master.v1.CreateChapterResponse
	(*CreateRegulationResponse)(nil), // 7: master.v1.CreateRegulationResponse
	(*GenerateLinksResponse)(nil),    // 8: master.v1.GenerateLinksResponse
}
var file_master_service_proto_depIdxs = []int32{
	1, // 0: master.v1.MasterGRPC.CreateChapter:input_type -> master.v1.CreateChapterRequest
	2, // 1: master.v1.MasterGRPC.CreateParagraphs:input_type -> master.v1.CreateParagraphsRequest
	3, // 2: master.v1.MasterGRPC.CreateRegulation:input_type -> master.v1.CreateRegulationRequest
	4, // 3: master.v1.MasterGRPC.GenerateLinks:input_type -> master.v1.GenerateLinksRequest
	5, // 4: master.v1.MasterGRPC.DeleteRegulation:input_type -> master.v1.DeleteRegulationRequest
	6, // 5: master.v1.MasterGRPC.CreateChapter:output_type -> master.v1.CreateChapterResponse
	0, // 6: master.v1.MasterGRPC.CreateParagraphs:output_type -> master.v1.Empty
	7, // 7: master.v1.MasterGRPC.CreateRegulation:output_type -> master.v1.CreateRegulationResponse
	8, // 8: master.v1.MasterGRPC.GenerateLinks:output_type -> master.v1.GenerateLinksResponse
	0, // 9: master.v1.MasterGRPC.DeleteRegulation:output_type -> master.v1.Empty
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_master_service_proto_init() }
func file_master_service_proto_init() {
	if File_master_service_proto != nil {
		return
	}
	file_master_message_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_master_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_master_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_master_service_proto_goTypes,
		DependencyIndexes: file_master_service_proto_depIdxs,
		MessageInfos:      file_master_service_proto_msgTypes,
	}.Build()
	File_master_service_proto = out.File
	file_master_service_proto_rawDesc = nil
	file_master_service_proto_goTypes = nil
	file_master_service_proto_depIdxs = nil
}