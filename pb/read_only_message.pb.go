// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: read_only_message.proto

package pb

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

type ReadOnlyParagraph struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID        uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Num       uint64 `protobuf:"varint,2,opt,name=Num,proto3" json:"Num,omitempty"`
	HasLinks  bool   `protobuf:"varint,3,opt,name=HasLinks,proto3" json:"HasLinks,omitempty"`
	IsTable   bool   `protobuf:"varint,4,opt,name=IsTable,proto3" json:"IsTable,omitempty"`
	IsNFT     bool   `protobuf:"varint,5,opt,name=IsNFT,proto3" json:"IsNFT,omitempty"`
	Class     string `protobuf:"bytes,6,opt,name=Class,proto3" json:"Class,omitempty"`
	Content   string `protobuf:"bytes,7,opt,name=Content,proto3" json:"Content,omitempty"`
	ChapterID uint64 `protobuf:"varint,8,opt,name=ChapterID,proto3" json:"ChapterID,omitempty"`
}

func (x *ReadOnlyParagraph) Reset() {
	*x = ReadOnlyParagraph{}
	if protoimpl.UnsafeEnabled {
		mi := &file_read_only_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadOnlyParagraph) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadOnlyParagraph) ProtoMessage() {}

func (x *ReadOnlyParagraph) ProtoReflect() protoreflect.Message {
	mi := &file_read_only_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadOnlyParagraph.ProtoReflect.Descriptor instead.
func (*ReadOnlyParagraph) Descriptor() ([]byte, []int) {
	return file_read_only_message_proto_rawDescGZIP(), []int{0}
}

func (x *ReadOnlyParagraph) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *ReadOnlyParagraph) GetNum() uint64 {
	if x != nil {
		return x.Num
	}
	return 0
}

func (x *ReadOnlyParagraph) GetHasLinks() bool {
	if x != nil {
		return x.HasLinks
	}
	return false
}

func (x *ReadOnlyParagraph) GetIsTable() bool {
	if x != nil {
		return x.IsTable
	}
	return false
}

func (x *ReadOnlyParagraph) GetIsNFT() bool {
	if x != nil {
		return x.IsNFT
	}
	return false
}

func (x *ReadOnlyParagraph) GetClass() string {
	if x != nil {
		return x.Class
	}
	return ""
}

func (x *ReadOnlyParagraph) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *ReadOnlyParagraph) GetChapterID() uint64 {
	if x != nil {
		return x.ChapterID
	}
	return 0
}

type ReadOnlyChapter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID           uint64               `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name         string               `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Num          string               `protobuf:"bytes,3,opt,name=Num,proto3" json:"Num,omitempty"`
	RegulationID string               `protobuf:"bytes,4,opt,name=RegulationID,proto3" json:"RegulationID,omitempty"`
	OrderNum     string               `protobuf:"bytes,5,opt,name=OrderNum,proto3" json:"OrderNum,omitempty"`
	Paragraphs   []*ReadOnlyParagraph `protobuf:"bytes,6,rep,name=Paragraphs,proto3" json:"Paragraphs,omitempty"`
	UpdatedAt    string               `protobuf:"bytes,7,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
}

func (x *ReadOnlyChapter) Reset() {
	*x = ReadOnlyChapter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_read_only_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadOnlyChapter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadOnlyChapter) ProtoMessage() {}

func (x *ReadOnlyChapter) ProtoReflect() protoreflect.Message {
	mi := &file_read_only_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadOnlyChapter.ProtoReflect.Descriptor instead.
func (*ReadOnlyChapter) Descriptor() ([]byte, []int) {
	return file_read_only_message_proto_rawDescGZIP(), []int{1}
}

func (x *ReadOnlyChapter) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *ReadOnlyChapter) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ReadOnlyChapter) GetNum() string {
	if x != nil {
		return x.Num
	}
	return ""
}

func (x *ReadOnlyChapter) GetRegulationID() string {
	if x != nil {
		return x.RegulationID
	}
	return ""
}

func (x *ReadOnlyChapter) GetOrderNum() string {
	if x != nil {
		return x.OrderNum
	}
	return ""
}

func (x *ReadOnlyChapter) GetParagraphs() []*ReadOnlyParagraph {
	if x != nil {
		return x.Paragraphs
	}
	return nil
}

func (x *ReadOnlyChapter) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type GetRegulationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *GetRegulationRequest) Reset() {
	*x = GetRegulationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_read_only_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRegulationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRegulationRequest) ProtoMessage() {}

func (x *GetRegulationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_read_only_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRegulationRequest.ProtoReflect.Descriptor instead.
func (*GetRegulationRequest) Descriptor() ([]byte, []int) {
	return file_read_only_message_proto_rawDescGZIP(), []int{2}
}

func (x *GetRegulationRequest) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type GetRegulationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Abbreviation string `protobuf:"bytes,2,opt,name=Abbreviation,proto3" json:"Abbreviation,omitempty"`
	Title        string `protobuf:"bytes,3,opt,name=Title,proto3" json:"Title,omitempty"`
}

func (x *GetRegulationResponse) Reset() {
	*x = GetRegulationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_read_only_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRegulationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRegulationResponse) ProtoMessage() {}

func (x *GetRegulationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_read_only_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRegulationResponse.ProtoReflect.Descriptor instead.
func (*GetRegulationResponse) Descriptor() ([]byte, []int) {
	return file_read_only_message_proto_rawDescGZIP(), []int{3}
}

func (x *GetRegulationResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetRegulationResponse) GetAbbreviation() string {
	if x != nil {
		return x.Abbreviation
	}
	return ""
}

func (x *GetRegulationResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type GetChapterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *GetChapterRequest) Reset() {
	*x = GetChapterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_read_only_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChapterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChapterRequest) ProtoMessage() {}

func (x *GetChapterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_read_only_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChapterRequest.ProtoReflect.Descriptor instead.
func (*GetChapterRequest) Descriptor() ([]byte, []int) {
	return file_read_only_message_proto_rawDescGZIP(), []int{4}
}

func (x *GetChapterRequest) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type GetChapterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID           uint64               `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name         string               `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Num          string               `protobuf:"bytes,3,opt,name=Num,proto3" json:"Num,omitempty"`
	RegulationID string               `protobuf:"bytes,4,opt,name=RegulationID,proto3" json:"RegulationID,omitempty"`
	OrderNum     string               `protobuf:"bytes,5,opt,name=OrderNum,proto3" json:"OrderNum,omitempty"`
	Paragraphs   []*ReadOnlyParagraph `protobuf:"bytes,6,rep,name=Paragraphs,proto3" json:"Paragraphs,omitempty"`
	UpdatedAt    string               `protobuf:"bytes,7,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
}

func (x *GetChapterResponse) Reset() {
	*x = GetChapterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_read_only_message_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChapterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChapterResponse) ProtoMessage() {}

func (x *GetChapterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_read_only_message_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChapterResponse.ProtoReflect.Descriptor instead.
func (*GetChapterResponse) Descriptor() ([]byte, []int) {
	return file_read_only_message_proto_rawDescGZIP(), []int{5}
}

func (x *GetChapterResponse) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *GetChapterResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetChapterResponse) GetNum() string {
	if x != nil {
		return x.Num
	}
	return ""
}

func (x *GetChapterResponse) GetRegulationID() string {
	if x != nil {
		return x.RegulationID
	}
	return ""
}

func (x *GetChapterResponse) GetOrderNum() string {
	if x != nil {
		return x.OrderNum
	}
	return ""
}

func (x *GetChapterResponse) GetParagraphs() []*ReadOnlyParagraph {
	if x != nil {
		return x.Paragraphs
	}
	return nil
}

func (x *GetChapterResponse) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type GetAllChaptersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *GetAllChaptersRequest) Reset() {
	*x = GetAllChaptersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_read_only_message_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllChaptersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllChaptersRequest) ProtoMessage() {}

func (x *GetAllChaptersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_read_only_message_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllChaptersRequest.ProtoReflect.Descriptor instead.
func (*GetAllChaptersRequest) Descriptor() ([]byte, []int) {
	return file_read_only_message_proto_rawDescGZIP(), []int{6}
}

func (x *GetAllChaptersRequest) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type GetAllChaptersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chapters []*ReadOnlyChapter `protobuf:"bytes,1,rep,name=Chapters,proto3" json:"Chapters,omitempty"`
}

func (x *GetAllChaptersResponse) Reset() {
	*x = GetAllChaptersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_read_only_message_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllChaptersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllChaptersResponse) ProtoMessage() {}

func (x *GetAllChaptersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_read_only_message_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllChaptersResponse.ProtoReflect.Descriptor instead.
func (*GetAllChaptersResponse) Descriptor() ([]byte, []int) {
	return file_read_only_message_proto_rawDescGZIP(), []int{7}
}

func (x *GetAllChaptersResponse) GetChapters() []*ReadOnlyChapter {
	if x != nil {
		return x.Chapters
	}
	return nil
}

type GetParagraphsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *GetParagraphsRequest) Reset() {
	*x = GetParagraphsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_read_only_message_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetParagraphsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetParagraphsRequest) ProtoMessage() {}

func (x *GetParagraphsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_read_only_message_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetParagraphsRequest.ProtoReflect.Descriptor instead.
func (*GetParagraphsRequest) Descriptor() ([]byte, []int) {
	return file_read_only_message_proto_rawDescGZIP(), []int{8}
}

func (x *GetParagraphsRequest) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type GetParagraphsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paragraphs []*ReadOnlyParagraph `protobuf:"bytes,1,rep,name=Paragraphs,proto3" json:"Paragraphs,omitempty"`
}

func (x *GetParagraphsResponse) Reset() {
	*x = GetParagraphsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_read_only_message_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetParagraphsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetParagraphsResponse) ProtoMessage() {}

func (x *GetParagraphsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_read_only_message_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetParagraphsResponse.ProtoReflect.Descriptor instead.
func (*GetParagraphsResponse) Descriptor() ([]byte, []int) {
	return file_read_only_message_proto_rawDescGZIP(), []int{9}
}

func (x *GetParagraphsResponse) GetParagraphs() []*ReadOnlyParagraph {
	if x != nil {
		return x.Paragraphs
	}
	return nil
}

type SearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SearchQuery string `protobuf:"bytes,1,opt,name=SearchQuery,proto3" json:"SearchQuery,omitempty"`
	Limit       string `protobuf:"bytes,2,opt,name=Limit,proto3" json:"Limit,omitempty"`
	Offset      string `protobuf:"bytes,3,opt,name=Offset,proto3" json:"Offset,omitempty"`
}

func (x *SearchRequest) Reset() {
	*x = SearchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_read_only_message_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchRequest) ProtoMessage() {}

func (x *SearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_read_only_message_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchRequest.ProtoReflect.Descriptor instead.
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return file_read_only_message_proto_rawDescGZIP(), []int{10}
}

func (x *SearchRequest) GetSearchQuery() string {
	if x != nil {
		return x.SearchQuery
	}
	return ""
}

func (x *SearchRequest) GetLimit() string {
	if x != nil {
		return x.Limit
	}
	return ""
}

func (x *SearchRequest) GetOffset() string {
	if x != nil {
		return x.Offset
	}
	return ""
}

type SearchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RID       string `protobuf:"bytes,1,opt,name=RID,proto3" json:"RID,omitempty"`
	RName     string `protobuf:"bytes,2,opt,name=RName,proto3" json:"RName,omitempty"`
	CID       string `protobuf:"bytes,3,opt,name=CID,proto3" json:"CID,omitempty"`
	CName     string `protobuf:"bytes,4,opt,name=CName,proto3" json:"CName,omitempty"`
	UpdatedAt string `protobuf:"bytes,5,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
	PID       string `protobuf:"bytes,6,opt,name=PID,proto3" json:"PID,omitempty"`
	Text      string `protobuf:"bytes,7,opt,name=Text,proto3" json:"Text,omitempty"`
	Count     string `protobuf:"bytes,8,opt,name=Count,proto3" json:"Count,omitempty"`
}

func (x *SearchResponse) Reset() {
	*x = SearchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_read_only_message_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchResponse) ProtoMessage() {}

func (x *SearchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_read_only_message_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchResponse.ProtoReflect.Descriptor instead.
func (*SearchResponse) Descriptor() ([]byte, []int) {
	return file_read_only_message_proto_rawDescGZIP(), []int{11}
}

func (x *SearchResponse) GetRID() string {
	if x != nil {
		return x.RID
	}
	return ""
}

func (x *SearchResponse) GetRName() string {
	if x != nil {
		return x.RName
	}
	return ""
}

func (x *SearchResponse) GetCID() string {
	if x != nil {
		return x.CID
	}
	return ""
}

func (x *SearchResponse) GetCName() string {
	if x != nil {
		return x.CName
	}
	return ""
}

func (x *SearchResponse) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *SearchResponse) GetPID() string {
	if x != nil {
		return x.PID
	}
	return ""
}

func (x *SearchResponse) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *SearchResponse) GetCount() string {
	if x != nil {
		return x.Count
	}
	return ""
}

type SearchResponseMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response []*SearchResponse `protobuf:"bytes,1,rep,name=response,proto3" json:"response,omitempty"`
}

func (x *SearchResponseMessage) Reset() {
	*x = SearchResponseMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_read_only_message_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchResponseMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchResponseMessage) ProtoMessage() {}

func (x *SearchResponseMessage) ProtoReflect() protoreflect.Message {
	mi := &file_read_only_message_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchResponseMessage.ProtoReflect.Descriptor instead.
func (*SearchResponseMessage) Descriptor() ([]byte, []int) {
	return file_read_only_message_proto_rawDescGZIP(), []int{12}
}

func (x *SearchResponseMessage) GetResponse() []*SearchResponse {
	if x != nil {
		return x.Response
	}
	return nil
}

var File_read_only_message_proto protoreflect.FileDescriptor

var file_read_only_message_proto_rawDesc = []byte{
	0x0a, 0x17, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcf, 0x01, 0x0a, 0x11, 0x52, 0x65,
	0x61, 0x64, 0x4f, 0x6e, 0x6c, 0x79, 0x50, 0x61, 0x72, 0x61, 0x67, 0x72, 0x61, 0x70, 0x68, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x12,
	0x10, 0x0a, 0x03, 0x4e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x4e, 0x75,
	0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x48, 0x61, 0x73, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x08, 0x48, 0x61, 0x73, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x49, 0x73, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x49, 0x73, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x49, 0x73, 0x4e, 0x46, 0x54,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x49, 0x73, 0x4e, 0x46, 0x54, 0x12, 0x14, 0x0a,
	0x05, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x43, 0x6c,
	0x61, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x49, 0x44, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x09, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x49, 0x44, 0x22, 0xd9, 0x01, 0x0a, 0x0f,
	0x52, 0x65, 0x61, 0x64, 0x4f, 0x6e, 0x6c, 0x79, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x12,
	0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x4e, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x4e, 0x75, 0x6d, 0x12, 0x22, 0x0a, 0x0c, 0x52, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x52, 0x65, 0x67,
	0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x4e, 0x75, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x4e, 0x75, 0x6d, 0x12, 0x32, 0x0a, 0x0a, 0x50, 0x61, 0x72, 0x61, 0x67, 0x72, 0x61,
	0x70, 0x68, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x52, 0x65, 0x61, 0x64,
	0x4f, 0x6e, 0x6c, 0x79, 0x50, 0x61, 0x72, 0x61, 0x67, 0x72, 0x61, 0x70, 0x68, 0x52, 0x0a, 0x50,
	0x61, 0x72, 0x61, 0x67, 0x72, 0x61, 0x70, 0x68, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x26, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x67, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x22,
	0x65, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x52, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c,
	0x41, 0x62, 0x62, 0x72, 0x65, 0x76, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x41, 0x62, 0x62, 0x72, 0x65, 0x76, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x23, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61,
	0x70, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x22, 0xdc, 0x01, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02,
	0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x4e, 0x75, 0x6d, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x4e, 0x75, 0x6d, 0x12, 0x22, 0x0a, 0x0c, 0x52, 0x65, 0x67, 0x75,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x52, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x12, 0x32, 0x0a, 0x0a, 0x50, 0x61, 0x72, 0x61,
	0x67, 0x72, 0x61, 0x70, 0x68, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x52,
	0x65, 0x61, 0x64, 0x4f, 0x6e, 0x6c, 0x79, 0x50, 0x61, 0x72, 0x61, 0x67, 0x72, 0x61, 0x70, 0x68,
	0x52, 0x0a, 0x50, 0x61, 0x72, 0x61, 0x67, 0x72, 0x61, 0x70, 0x68, 0x73, 0x12, 0x1c, 0x0a, 0x09,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x27, 0x0a, 0x15, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x02, 0x49, 0x44, 0x22, 0x46, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x68, 0x61,
	0x70, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a,
	0x08, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x4f, 0x6e, 0x6c, 0x79, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65,
	0x72, 0x52, 0x08, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x73, 0x22, 0x26, 0x0a, 0x14, 0x47,
	0x65, 0x74, 0x50, 0x61, 0x72, 0x61, 0x67, 0x72, 0x61, 0x70, 0x68, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x02, 0x49, 0x44, 0x22, 0x4b, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72, 0x61, 0x67, 0x72,
	0x61, 0x70, 0x68, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x0a,
	0x50, 0x61, 0x72, 0x61, 0x67, 0x72, 0x61, 0x70, 0x68, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x4f, 0x6e, 0x6c, 0x79, 0x50, 0x61, 0x72, 0x61, 0x67,
	0x72, 0x61, 0x70, 0x68, 0x52, 0x0a, 0x50, 0x61, 0x72, 0x61, 0x67, 0x72, 0x61, 0x70, 0x68, 0x73,
	0x22, 0x5f, 0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x20, 0x0a, 0x0b, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x22, 0xba, 0x01, 0x0a, 0x0e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x52, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x52, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x52, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x52, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x43, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x43, 0x49, 0x44, 0x12, 0x14,
	0x0a, 0x05, 0x43, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x43,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x50, 0x49, 0x44, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x50, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x65, 0x78, 0x74, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x54, 0x65, 0x78, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x44,
	0x0a, 0x15, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2b, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_read_only_message_proto_rawDescOnce sync.Once
	file_read_only_message_proto_rawDescData = file_read_only_message_proto_rawDesc
)

func file_read_only_message_proto_rawDescGZIP() []byte {
	file_read_only_message_proto_rawDescOnce.Do(func() {
		file_read_only_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_read_only_message_proto_rawDescData)
	})
	return file_read_only_message_proto_rawDescData
}

var file_read_only_message_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_read_only_message_proto_goTypes = []interface{}{
	(*ReadOnlyParagraph)(nil),      // 0: ReadOnlyParagraph
	(*ReadOnlyChapter)(nil),        // 1: ReadOnlyChapter
	(*GetRegulationRequest)(nil),   // 2: GetRegulationRequest
	(*GetRegulationResponse)(nil),  // 3: GetRegulationResponse
	(*GetChapterRequest)(nil),      // 4: GetChapterRequest
	(*GetChapterResponse)(nil),     // 5: GetChapterResponse
	(*GetAllChaptersRequest)(nil),  // 6: GetAllChaptersRequest
	(*GetAllChaptersResponse)(nil), // 7: GetAllChaptersResponse
	(*GetParagraphsRequest)(nil),   // 8: GetParagraphsRequest
	(*GetParagraphsResponse)(nil),  // 9: GetParagraphsResponse
	(*SearchRequest)(nil),          // 10: SearchRequest
	(*SearchResponse)(nil),         // 11: SearchResponse
	(*SearchResponseMessage)(nil),  // 12: SearchResponseMessage
}
var file_read_only_message_proto_depIdxs = []int32{
	0,  // 0: ReadOnlyChapter.Paragraphs:type_name -> ReadOnlyParagraph
	0,  // 1: GetChapterResponse.Paragraphs:type_name -> ReadOnlyParagraph
	1,  // 2: GetAllChaptersResponse.Chapters:type_name -> ReadOnlyChapter
	0,  // 3: GetParagraphsResponse.Paragraphs:type_name -> ReadOnlyParagraph
	11, // 4: SearchResponseMessage.response:type_name -> SearchResponse
	5,  // [5:5] is the sub-list for method output_type
	5,  // [5:5] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_read_only_message_proto_init() }
func file_read_only_message_proto_init() {
	if File_read_only_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_read_only_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadOnlyParagraph); i {
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
		file_read_only_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadOnlyChapter); i {
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
		file_read_only_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRegulationRequest); i {
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
		file_read_only_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRegulationResponse); i {
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
		file_read_only_message_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetChapterRequest); i {
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
		file_read_only_message_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetChapterResponse); i {
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
		file_read_only_message_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllChaptersRequest); i {
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
		file_read_only_message_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllChaptersResponse); i {
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
		file_read_only_message_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetParagraphsRequest); i {
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
		file_read_only_message_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetParagraphsResponse); i {
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
		file_read_only_message_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchRequest); i {
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
		file_read_only_message_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchResponse); i {
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
		file_read_only_message_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchResponseMessage); i {
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
			RawDescriptor: file_read_only_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_read_only_message_proto_goTypes,
		DependencyIndexes: file_read_only_message_proto_depIdxs,
		MessageInfos:      file_read_only_message_proto_msgTypes,
	}.Build()
	File_read_only_message_proto = out.File
	file_read_only_message_proto_rawDesc = nil
	file_read_only_message_proto_goTypes = nil
	file_read_only_message_proto_depIdxs = nil
}
