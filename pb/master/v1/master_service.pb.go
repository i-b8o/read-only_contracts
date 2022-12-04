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

type Regulation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID             uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	RegulationName string `protobuf:"bytes,2,opt,name=RegulationName,proto3" json:"RegulationName,omitempty"`
	Abbreviation   string `protobuf:"bytes,3,opt,name=Abbreviation,proto3" json:"Abbreviation,omitempty"`
	Title          string `protobuf:"bytes,4,opt,name=Title,proto3" json:"Title,omitempty"`
}

func (x *Regulation) Reset() {
	*x = Regulation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_master_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Regulation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Regulation) ProtoMessage() {}

func (x *Regulation) ProtoReflect() protoreflect.Message {
	mi := &file_master_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Regulation.ProtoReflect.Descriptor instead.
func (*Regulation) Descriptor() ([]byte, []int) {
	return file_master_service_proto_rawDescGZIP(), []int{1}
}

func (x *Regulation) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Regulation) GetRegulationName() string {
	if x != nil {
		return x.RegulationName
	}
	return ""
}

func (x *Regulation) GetAbbreviation() string {
	if x != nil {
		return x.Abbreviation
	}
	return ""
}

func (x *Regulation) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type GetAllRegulationsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Regulations []*Regulation `protobuf:"bytes,1,rep,name=Regulations,proto3" json:"Regulations,omitempty"`
}

func (x *GetAllRegulationsResponse) Reset() {
	*x = GetAllRegulationsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_master_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllRegulationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllRegulationsResponse) ProtoMessage() {}

func (x *GetAllRegulationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_master_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllRegulationsResponse.ProtoReflect.Descriptor instead.
func (*GetAllRegulationsResponse) Descriptor() ([]byte, []int) {
	return file_master_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetAllRegulationsResponse) GetRegulations() []*Regulation {
	if x != nil {
		return x.Regulations
	}
	return nil
}

type CreateChapterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PseudoId     string `protobuf:"bytes,1,opt,name=PseudoId,proto3" json:"PseudoId,omitempty"`
	RegulationId uint64 `protobuf:"varint,2,opt,name=RegulationId,proto3" json:"RegulationId,omitempty"`
	ChapterName  string `protobuf:"bytes,3,opt,name=ChapterName,proto3" json:"ChapterName,omitempty"`
	ChapterNum   string `protobuf:"bytes,4,opt,name=ChapterNum,proto3" json:"ChapterNum,omitempty"`
	OrderNum     uint32 `protobuf:"varint,5,opt,name=OrderNum,proto3" json:"OrderNum,omitempty"`
}

func (x *CreateChapterRequest) Reset() {
	*x = CreateChapterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_master_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateChapterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateChapterRequest) ProtoMessage() {}

func (x *CreateChapterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_master_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateChapterRequest.ProtoReflect.Descriptor instead.
func (*CreateChapterRequest) Descriptor() ([]byte, []int) {
	return file_master_service_proto_rawDescGZIP(), []int{3}
}

func (x *CreateChapterRequest) GetPseudoId() string {
	if x != nil {
		return x.PseudoId
	}
	return ""
}

func (x *CreateChapterRequest) GetRegulationId() uint64 {
	if x != nil {
		return x.RegulationId
	}
	return 0
}

func (x *CreateChapterRequest) GetChapterName() string {
	if x != nil {
		return x.ChapterName
	}
	return ""
}

func (x *CreateChapterRequest) GetChapterNum() string {
	if x != nil {
		return x.ChapterNum
	}
	return ""
}

func (x *CreateChapterRequest) GetOrderNum() uint32 {
	if x != nil {
		return x.OrderNum
	}
	return 0
}

type CreateChapterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *CreateChapterResponse) Reset() {
	*x = CreateChapterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_master_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateChapterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateChapterResponse) ProtoMessage() {}

func (x *CreateChapterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_master_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateChapterResponse.ProtoReflect.Descriptor instead.
func (*CreateChapterResponse) Descriptor() ([]byte, []int) {
	return file_master_service_proto_rawDescGZIP(), []int{4}
}

func (x *CreateChapterResponse) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type Paragraph struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ParagraphId       uint64 `protobuf:"varint,1,opt,name=ParagraphId,proto3" json:"ParagraphId,omitempty"`
	ParagraphOrderNum uint32 `protobuf:"varint,2,opt,name=ParagraphOrderNum,proto3" json:"ParagraphOrderNum,omitempty"`
	IsTable           bool   `protobuf:"varint,3,opt,name=IsTable,proto3" json:"IsTable,omitempty"`
	IsNFT             bool   `protobuf:"varint,4,opt,name=IsNFT,proto3" json:"IsNFT,omitempty"`
	HasLinks          bool   `protobuf:"varint,5,opt,name=HasLinks,proto3" json:"HasLinks,omitempty"`
	ParagraphClass    string `protobuf:"bytes,6,opt,name=ParagraphClass,proto3" json:"ParagraphClass,omitempty"`
	ParagraphText     string `protobuf:"bytes,7,opt,name=ParagraphText,proto3" json:"ParagraphText,omitempty"`
	ChapterId         uint64 `protobuf:"varint,8,opt,name=ChapterId,proto3" json:"ChapterId,omitempty"`
}

func (x *Paragraph) Reset() {
	*x = Paragraph{}
	if protoimpl.UnsafeEnabled {
		mi := &file_master_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Paragraph) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Paragraph) ProtoMessage() {}

func (x *Paragraph) ProtoReflect() protoreflect.Message {
	mi := &file_master_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Paragraph.ProtoReflect.Descriptor instead.
func (*Paragraph) Descriptor() ([]byte, []int) {
	return file_master_service_proto_rawDescGZIP(), []int{5}
}

func (x *Paragraph) GetParagraphId() uint64 {
	if x != nil {
		return x.ParagraphId
	}
	return 0
}

func (x *Paragraph) GetParagraphOrderNum() uint32 {
	if x != nil {
		return x.ParagraphOrderNum
	}
	return 0
}

func (x *Paragraph) GetIsTable() bool {
	if x != nil {
		return x.IsTable
	}
	return false
}

func (x *Paragraph) GetIsNFT() bool {
	if x != nil {
		return x.IsNFT
	}
	return false
}

func (x *Paragraph) GetHasLinks() bool {
	if x != nil {
		return x.HasLinks
	}
	return false
}

func (x *Paragraph) GetParagraphClass() string {
	if x != nil {
		return x.ParagraphClass
	}
	return ""
}

func (x *Paragraph) GetParagraphText() string {
	if x != nil {
		return x.ParagraphText
	}
	return ""
}

func (x *Paragraph) GetChapterId() uint64 {
	if x != nil {
		return x.ChapterId
	}
	return 0
}

type CreateParagraphsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paragraphs []*Paragraph `protobuf:"bytes,1,rep,name=paragraphs,proto3" json:"paragraphs,omitempty"`
}

func (x *CreateParagraphsRequest) Reset() {
	*x = CreateParagraphsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_master_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateParagraphsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateParagraphsRequest) ProtoMessage() {}

func (x *CreateParagraphsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_master_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateParagraphsRequest.ProtoReflect.Descriptor instead.
func (*CreateParagraphsRequest) Descriptor() ([]byte, []int) {
	return file_master_service_proto_rawDescGZIP(), []int{6}
}

func (x *CreateParagraphsRequest) GetParagraphs() []*Paragraph {
	if x != nil {
		return x.Paragraphs
	}
	return nil
}

type CreateRegulationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PseudoId       string `protobuf:"bytes,1,opt,name=PseudoId,proto3" json:"PseudoId,omitempty"`
	RegulationName string `protobuf:"bytes,2,opt,name=RegulationName,proto3" json:"RegulationName,omitempty"`
	Abbreviation   string `protobuf:"bytes,3,opt,name=Abbreviation,proto3" json:"Abbreviation,omitempty"`
	Title          string `protobuf:"bytes,4,opt,name=Title,proto3" json:"Title,omitempty"`
}

func (x *CreateRegulationRequest) Reset() {
	*x = CreateRegulationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_master_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRegulationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRegulationRequest) ProtoMessage() {}

func (x *CreateRegulationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_master_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRegulationRequest.ProtoReflect.Descriptor instead.
func (*CreateRegulationRequest) Descriptor() ([]byte, []int) {
	return file_master_service_proto_rawDescGZIP(), []int{7}
}

func (x *CreateRegulationRequest) GetPseudoId() string {
	if x != nil {
		return x.PseudoId
	}
	return ""
}

func (x *CreateRegulationRequest) GetRegulationName() string {
	if x != nil {
		return x.RegulationName
	}
	return ""
}

func (x *CreateRegulationRequest) GetAbbreviation() string {
	if x != nil {
		return x.Abbreviation
	}
	return ""
}

func (x *CreateRegulationRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type CreateRegulationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *CreateRegulationResponse) Reset() {
	*x = CreateRegulationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_master_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRegulationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRegulationResponse) ProtoMessage() {}

func (x *CreateRegulationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_master_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRegulationResponse.ProtoReflect.Descriptor instead.
func (*CreateRegulationResponse) Descriptor() ([]byte, []int) {
	return file_master_service_proto_rawDescGZIP(), []int{8}
}

func (x *CreateRegulationResponse) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type GenerateLinksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *GenerateLinksRequest) Reset() {
	*x = GenerateLinksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_master_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateLinksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateLinksRequest) ProtoMessage() {}

func (x *GenerateLinksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_master_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateLinksRequest.ProtoReflect.Descriptor instead.
func (*GenerateLinksRequest) Descriptor() ([]byte, []int) {
	return file_master_service_proto_rawDescGZIP(), []int{9}
}

func (x *GenerateLinksRequest) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type GenerateLinksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *GenerateLinksResponse) Reset() {
	*x = GenerateLinksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_master_service_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateLinksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateLinksResponse) ProtoMessage() {}

func (x *GenerateLinksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_master_service_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateLinksResponse.ProtoReflect.Descriptor instead.
func (*GenerateLinksResponse) Descriptor() ([]byte, []int) {
	return file_master_service_proto_rawDescGZIP(), []int{10}
}

func (x *GenerateLinksResponse) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type DeleteRegulationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *DeleteRegulationRequest) Reset() {
	*x = DeleteRegulationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_master_service_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRegulationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRegulationRequest) ProtoMessage() {}

func (x *DeleteRegulationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_master_service_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRegulationRequest.ProtoReflect.Descriptor instead.
func (*DeleteRegulationRequest) Descriptor() ([]byte, []int) {
	return file_master_service_proto_rawDescGZIP(), []int{11}
}

func (x *DeleteRegulationRequest) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

var File_master_service_proto protoreflect.FileDescriptor

var file_master_service_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x7e, 0x0a, 0x0a, 0x52, 0x65,
	0x67, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x12, 0x26, 0x0a, 0x0e, 0x52, 0x65, 0x67, 0x75,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x52, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x22, 0x0a, 0x0c, 0x41, 0x62, 0x62, 0x72, 0x65, 0x76, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x41, 0x62, 0x62, 0x72, 0x65, 0x76, 0x69, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x54, 0x0a, 0x19, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x0b, 0x52, 0x65, 0x67, 0x75, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6d,
	0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x52, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x22, 0xb4, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x70, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x73, 0x65,
	0x75, 0x64, 0x6f, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x73, 0x65,
	0x75, 0x64, 0x6f, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x52, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x52, 0x65, 0x67,
	0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x68, 0x61,
	0x70, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x43,
	0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x22, 0x27, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44,
	0x22, 0x93, 0x02, 0x0a, 0x09, 0x50, 0x61, 0x72, 0x61, 0x67, 0x72, 0x61, 0x70, 0x68, 0x12, 0x20,
	0x0a, 0x0b, 0x50, 0x61, 0x72, 0x61, 0x67, 0x72, 0x61, 0x70, 0x68, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0b, 0x50, 0x61, 0x72, 0x61, 0x67, 0x72, 0x61, 0x70, 0x68, 0x49, 0x64,
	0x12, 0x2c, 0x0a, 0x11, 0x50, 0x61, 0x72, 0x61, 0x67, 0x72, 0x61, 0x70, 0x68, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x4e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x11, 0x50, 0x61, 0x72,
	0x61, 0x67, 0x72, 0x61, 0x70, 0x68, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x12, 0x18,
	0x0a, 0x07, 0x49, 0x73, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x49, 0x73, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x49, 0x73, 0x4e, 0x46,
	0x54, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x49, 0x73, 0x4e, 0x46, 0x54, 0x12, 0x1a,
	0x0a, 0x08, 0x48, 0x61, 0x73, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x48, 0x61, 0x73, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x50, 0x61,
	0x72, 0x61, 0x67, 0x72, 0x61, 0x70, 0x68, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x50, 0x61, 0x72, 0x61, 0x67, 0x72, 0x61, 0x70, 0x68, 0x43, 0x6c, 0x61,
	0x73, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x50, 0x61, 0x72, 0x61, 0x67, 0x72, 0x61, 0x70, 0x68, 0x54,
	0x65, 0x78, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x50, 0x61, 0x72, 0x61, 0x67,
	0x72, 0x61, 0x70, 0x68, 0x54, 0x65, 0x78, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x68, 0x61, 0x70,
	0x74, 0x65, 0x72, 0x49, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x43, 0x68, 0x61,
	0x70, 0x74, 0x65, 0x72, 0x49, 0x64, 0x22, 0x4f, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x50, 0x61, 0x72, 0x61, 0x67, 0x72, 0x61, 0x70, 0x68, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x34, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x67, 0x72, 0x61, 0x70, 0x68, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x67, 0x72, 0x61, 0x70, 0x68, 0x52, 0x0a, 0x70, 0x61, 0x72,
	0x61, 0x67, 0x72, 0x61, 0x70, 0x68, 0x73, 0x22, 0x97, 0x01, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x73, 0x65, 0x75, 0x64, 0x6f, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x73, 0x65, 0x75, 0x64, 0x6f, 0x49, 0x64, 0x12,
	0x26, 0x0a, 0x0e, 0x52, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x52, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x41, 0x62, 0x62, 0x72, 0x65,
	0x76, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x41,
	0x62, 0x62, 0x72, 0x65, 0x76, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x54,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c,
	0x65, 0x22, 0x2a, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x67, 0x75, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x22, 0x26, 0x0a,
	0x14, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x02, 0x49, 0x44, 0x22, 0x27, 0x0a, 0x15, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x22, 0x29,
	0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x32, 0xfe, 0x03, 0x0a, 0x0a, 0x4d, 0x61,
	0x73, 0x74, 0x65, 0x72, 0x47, 0x52, 0x50, 0x43, 0x12, 0x4d, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x52, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x10, 0x2e,
	0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x24, 0x2e, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x52, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x54, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x12, 0x1f, 0x2e, 0x6d, 0x61, 0x73, 0x74, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x70, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x6d, 0x61, 0x73, 0x74,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x70,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4a, 0x0a,
	0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x61, 0x67, 0x72, 0x61, 0x70, 0x68,
	0x73, 0x12, 0x22, 0x2e, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x61, 0x67, 0x72, 0x61, 0x70, 0x68, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x5d, 0x0a, 0x10, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x2e,
	0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x23, 0x2e, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x54, 0x0a, 0x0d, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x12, 0x1f, 0x2e, 0x6d, 0x61, 0x73, 0x74,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4c, 0x69,
	0x6e, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x6d, 0x61, 0x73,
	0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4c,
	0x69, 0x6e, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4a,
	0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x22, 0x2e, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x15, 0x5a, 0x13, 0x6d, 0x61,
	0x73, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x62, 0x5f, 0x6d, 0x61, 0x73, 0x74, 0x65,
	0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_master_service_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_master_service_proto_goTypes = []interface{}{
	(*Empty)(nil),                     // 0: master.v1.Empty
	(*Regulation)(nil),                // 1: master.v1.Regulation
	(*GetAllRegulationsResponse)(nil), // 2: master.v1.GetAllRegulationsResponse
	(*CreateChapterRequest)(nil),      // 3: master.v1.CreateChapterRequest
	(*CreateChapterResponse)(nil),     // 4: master.v1.CreateChapterResponse
	(*Paragraph)(nil),                 // 5: master.v1.Paragraph
	(*CreateParagraphsRequest)(nil),   // 6: master.v1.CreateParagraphsRequest
	(*CreateRegulationRequest)(nil),   // 7: master.v1.CreateRegulationRequest
	(*CreateRegulationResponse)(nil),  // 8: master.v1.CreateRegulationResponse
	(*GenerateLinksRequest)(nil),      // 9: master.v1.GenerateLinksRequest
	(*GenerateLinksResponse)(nil),     // 10: master.v1.GenerateLinksResponse
	(*DeleteRegulationRequest)(nil),   // 11: master.v1.DeleteRegulationRequest
}
var file_master_service_proto_depIdxs = []int32{
	1,  // 0: master.v1.GetAllRegulationsResponse.Regulations:type_name -> master.v1.Regulation
	5,  // 1: master.v1.CreateParagraphsRequest.paragraphs:type_name -> master.v1.Paragraph
	0,  // 2: master.v1.MasterGRPC.GetAllRegulations:input_type -> master.v1.Empty
	3,  // 3: master.v1.MasterGRPC.CreateChapter:input_type -> master.v1.CreateChapterRequest
	6,  // 4: master.v1.MasterGRPC.CreateParagraphs:input_type -> master.v1.CreateParagraphsRequest
	7,  // 5: master.v1.MasterGRPC.CreateRegulation:input_type -> master.v1.CreateRegulationRequest
	9,  // 6: master.v1.MasterGRPC.GenerateLinks:input_type -> master.v1.GenerateLinksRequest
	11, // 7: master.v1.MasterGRPC.DeleteRegulation:input_type -> master.v1.DeleteRegulationRequest
	2,  // 8: master.v1.MasterGRPC.GetAllRegulations:output_type -> master.v1.GetAllRegulationsResponse
	4,  // 9: master.v1.MasterGRPC.CreateChapter:output_type -> master.v1.CreateChapterResponse
	0,  // 10: master.v1.MasterGRPC.CreateParagraphs:output_type -> master.v1.Empty
	8,  // 11: master.v1.MasterGRPC.CreateRegulation:output_type -> master.v1.CreateRegulationResponse
	10, // 12: master.v1.MasterGRPC.GenerateLinks:output_type -> master.v1.GenerateLinksResponse
	0,  // 13: master.v1.MasterGRPC.DeleteRegulation:output_type -> master.v1.Empty
	8,  // [8:14] is the sub-list for method output_type
	2,  // [2:8] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_master_service_proto_init() }
func file_master_service_proto_init() {
	if File_master_service_proto != nil {
		return
	}
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
		file_master_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Regulation); i {
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
		file_master_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllRegulationsResponse); i {
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
		file_master_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateChapterRequest); i {
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
		file_master_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateChapterResponse); i {
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
		file_master_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Paragraph); i {
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
		file_master_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateParagraphsRequest); i {
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
		file_master_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRegulationRequest); i {
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
		file_master_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRegulationResponse); i {
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
		file_master_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateLinksRequest); i {
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
		file_master_service_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateLinksResponse); i {
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
		file_master_service_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRegulationRequest); i {
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
			NumMessages:   12,
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
