// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: service.proto

package pb_writer

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// WriterGRPCClient is the client API for WriterGRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WriterGRPCClient interface {
	CreateRegulation(ctx context.Context, in *CreateRegulationRequest, opts ...grpc.CallOption) (*CreateRegulationResponse, error)
	DeleteRegulation(ctx context.Context, in *DeleteRegulationRequest, opts ...grpc.CallOption) (*Empty, error)
	CreateChapter(ctx context.Context, in *CreateChapterRequest, opts ...grpc.CallOption) (*CreateChapterResponse, error)
	GetAllChaptersIds(ctx context.Context, in *GetAllChaptersIdsRequest, opts ...grpc.CallOption) (*GetAllChaptersIdsResponse, error)
	GetRegulationIdByChapterId(ctx context.Context, in *GetRegulationIdByChapterIdRequest, opts ...grpc.CallOption) (*GetRegulationIdByChapterIdResponse, error)
	CreateAllParagraphs(ctx context.Context, in *CreateAllParagraphsRequest, opts ...grpc.CallOption) (*Empty, error)
	UpdateOneParagraph(ctx context.Context, in *UpdateOneParagraphRequest, opts ...grpc.CallOption) (*Empty, error)
	DeleteParagraphsForChapter(ctx context.Context, in *DeleteParagraphsForChapterRequest, opts ...grpc.CallOption) (*Empty, error)
	GetParagraphsWithHrefs(ctx context.Context, in *GetParagraphsWithHrefsRequest, opts ...grpc.CallOption) (*GetParagraphsWithHrefsResponse, error)
}

type writerGRPCClient struct {
	cc grpc.ClientConnInterface
}

func NewWriterGRPCClient(cc grpc.ClientConnInterface) WriterGRPCClient {
	return &writerGRPCClient{cc}
}

func (c *writerGRPCClient) CreateRegulation(ctx context.Context, in *CreateRegulationRequest, opts ...grpc.CallOption) (*CreateRegulationResponse, error) {
	out := new(CreateRegulationResponse)
	err := c.cc.Invoke(ctx, "/writer.v1.WriterGRPC/CreateRegulation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *writerGRPCClient) DeleteRegulation(ctx context.Context, in *DeleteRegulationRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/writer.v1.WriterGRPC/DeleteRegulation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *writerGRPCClient) CreateChapter(ctx context.Context, in *CreateChapterRequest, opts ...grpc.CallOption) (*CreateChapterResponse, error) {
	out := new(CreateChapterResponse)
	err := c.cc.Invoke(ctx, "/writer.v1.WriterGRPC/CreateChapter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *writerGRPCClient) GetAllChaptersIds(ctx context.Context, in *GetAllChaptersIdsRequest, opts ...grpc.CallOption) (*GetAllChaptersIdsResponse, error) {
	out := new(GetAllChaptersIdsResponse)
	err := c.cc.Invoke(ctx, "/writer.v1.WriterGRPC/GetAllChaptersIds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *writerGRPCClient) GetRegulationIdByChapterId(ctx context.Context, in *GetRegulationIdByChapterIdRequest, opts ...grpc.CallOption) (*GetRegulationIdByChapterIdResponse, error) {
	out := new(GetRegulationIdByChapterIdResponse)
	err := c.cc.Invoke(ctx, "/writer.v1.WriterGRPC/GetRegulationIdByChapterId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *writerGRPCClient) CreateAllParagraphs(ctx context.Context, in *CreateAllParagraphsRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/writer.v1.WriterGRPC/CreateAllParagraphs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *writerGRPCClient) UpdateOneParagraph(ctx context.Context, in *UpdateOneParagraphRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/writer.v1.WriterGRPC/UpdateOneParagraph", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *writerGRPCClient) DeleteParagraphsForChapter(ctx context.Context, in *DeleteParagraphsForChapterRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/writer.v1.WriterGRPC/DeleteParagraphsForChapter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *writerGRPCClient) GetParagraphsWithHrefs(ctx context.Context, in *GetParagraphsWithHrefsRequest, opts ...grpc.CallOption) (*GetParagraphsWithHrefsResponse, error) {
	out := new(GetParagraphsWithHrefsResponse)
	err := c.cc.Invoke(ctx, "/writer.v1.WriterGRPC/GetParagraphsWithHrefs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WriterGRPCServer is the server API for WriterGRPC service.
// All implementations must embed UnimplementedWriterGRPCServer
// for forward compatibility
type WriterGRPCServer interface {
	CreateRegulation(context.Context, *CreateRegulationRequest) (*CreateRegulationResponse, error)
	DeleteRegulation(context.Context, *DeleteRegulationRequest) (*Empty, error)
	CreateChapter(context.Context, *CreateChapterRequest) (*CreateChapterResponse, error)
	GetAllChaptersIds(context.Context, *GetAllChaptersIdsRequest) (*GetAllChaptersIdsResponse, error)
	GetRegulationIdByChapterId(context.Context, *GetRegulationIdByChapterIdRequest) (*GetRegulationIdByChapterIdResponse, error)
	CreateAllParagraphs(context.Context, *CreateAllParagraphsRequest) (*Empty, error)
	UpdateOneParagraph(context.Context, *UpdateOneParagraphRequest) (*Empty, error)
	DeleteParagraphsForChapter(context.Context, *DeleteParagraphsForChapterRequest) (*Empty, error)
	GetParagraphsWithHrefs(context.Context, *GetParagraphsWithHrefsRequest) (*GetParagraphsWithHrefsResponse, error)
	mustEmbedUnimplementedWriterGRPCServer()
}

// UnimplementedWriterGRPCServer must be embedded to have forward compatible implementations.
type UnimplementedWriterGRPCServer struct {
}

func (UnimplementedWriterGRPCServer) CreateRegulation(context.Context, *CreateRegulationRequest) (*CreateRegulationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRegulation not implemented")
}
func (UnimplementedWriterGRPCServer) DeleteRegulation(context.Context, *DeleteRegulationRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRegulation not implemented")
}
func (UnimplementedWriterGRPCServer) CreateChapter(context.Context, *CreateChapterRequest) (*CreateChapterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateChapter not implemented")
}
func (UnimplementedWriterGRPCServer) GetAllChaptersIds(context.Context, *GetAllChaptersIdsRequest) (*GetAllChaptersIdsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllChaptersIds not implemented")
}
func (UnimplementedWriterGRPCServer) GetRegulationIdByChapterId(context.Context, *GetRegulationIdByChapterIdRequest) (*GetRegulationIdByChapterIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRegulationIdByChapterId not implemented")
}
func (UnimplementedWriterGRPCServer) CreateAllParagraphs(context.Context, *CreateAllParagraphsRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAllParagraphs not implemented")
}
func (UnimplementedWriterGRPCServer) UpdateOneParagraph(context.Context, *UpdateOneParagraphRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOneParagraph not implemented")
}
func (UnimplementedWriterGRPCServer) DeleteParagraphsForChapter(context.Context, *DeleteParagraphsForChapterRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteParagraphsForChapter not implemented")
}
func (UnimplementedWriterGRPCServer) GetParagraphsWithHrefs(context.Context, *GetParagraphsWithHrefsRequest) (*GetParagraphsWithHrefsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetParagraphsWithHrefs not implemented")
}
func (UnimplementedWriterGRPCServer) mustEmbedUnimplementedWriterGRPCServer() {}

// UnsafeWriterGRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WriterGRPCServer will
// result in compilation errors.
type UnsafeWriterGRPCServer interface {
	mustEmbedUnimplementedWriterGRPCServer()
}

func RegisterWriterGRPCServer(s grpc.ServiceRegistrar, srv WriterGRPCServer) {
	s.RegisterService(&WriterGRPC_ServiceDesc, srv)
}

func _WriterGRPC_CreateRegulation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRegulationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WriterGRPCServer).CreateRegulation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/writer.v1.WriterGRPC/CreateRegulation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WriterGRPCServer).CreateRegulation(ctx, req.(*CreateRegulationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WriterGRPC_DeleteRegulation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRegulationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WriterGRPCServer).DeleteRegulation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/writer.v1.WriterGRPC/DeleteRegulation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WriterGRPCServer).DeleteRegulation(ctx, req.(*DeleteRegulationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WriterGRPC_CreateChapter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateChapterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WriterGRPCServer).CreateChapter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/writer.v1.WriterGRPC/CreateChapter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WriterGRPCServer).CreateChapter(ctx, req.(*CreateChapterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WriterGRPC_GetAllChaptersIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllChaptersIdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WriterGRPCServer).GetAllChaptersIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/writer.v1.WriterGRPC/GetAllChaptersIds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WriterGRPCServer).GetAllChaptersIds(ctx, req.(*GetAllChaptersIdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WriterGRPC_GetRegulationIdByChapterId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRegulationIdByChapterIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WriterGRPCServer).GetRegulationIdByChapterId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/writer.v1.WriterGRPC/GetRegulationIdByChapterId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WriterGRPCServer).GetRegulationIdByChapterId(ctx, req.(*GetRegulationIdByChapterIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WriterGRPC_CreateAllParagraphs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAllParagraphsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WriterGRPCServer).CreateAllParagraphs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/writer.v1.WriterGRPC/CreateAllParagraphs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WriterGRPCServer).CreateAllParagraphs(ctx, req.(*CreateAllParagraphsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WriterGRPC_UpdateOneParagraph_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOneParagraphRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WriterGRPCServer).UpdateOneParagraph(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/writer.v1.WriterGRPC/UpdateOneParagraph",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WriterGRPCServer).UpdateOneParagraph(ctx, req.(*UpdateOneParagraphRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WriterGRPC_DeleteParagraphsForChapter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteParagraphsForChapterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WriterGRPCServer).DeleteParagraphsForChapter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/writer.v1.WriterGRPC/DeleteParagraphsForChapter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WriterGRPCServer).DeleteParagraphsForChapter(ctx, req.(*DeleteParagraphsForChapterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WriterGRPC_GetParagraphsWithHrefs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetParagraphsWithHrefsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WriterGRPCServer).GetParagraphsWithHrefs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/writer.v1.WriterGRPC/GetParagraphsWithHrefs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WriterGRPCServer).GetParagraphsWithHrefs(ctx, req.(*GetParagraphsWithHrefsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WriterGRPC_ServiceDesc is the grpc.ServiceDesc for WriterGRPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WriterGRPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "writer.v1.WriterGRPC",
	HandlerType: (*WriterGRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRegulation",
			Handler:    _WriterGRPC_CreateRegulation_Handler,
		},
		{
			MethodName: "DeleteRegulation",
			Handler:    _WriterGRPC_DeleteRegulation_Handler,
		},
		{
			MethodName: "CreateChapter",
			Handler:    _WriterGRPC_CreateChapter_Handler,
		},
		{
			MethodName: "GetAllChaptersIds",
			Handler:    _WriterGRPC_GetAllChaptersIds_Handler,
		},
		{
			MethodName: "GetRegulationIdByChapterId",
			Handler:    _WriterGRPC_GetRegulationIdByChapterId_Handler,
		},
		{
			MethodName: "CreateAllParagraphs",
			Handler:    _WriterGRPC_CreateAllParagraphs_Handler,
		},
		{
			MethodName: "UpdateOneParagraph",
			Handler:    _WriterGRPC_UpdateOneParagraph_Handler,
		},
		{
			MethodName: "DeleteParagraphsForChapter",
			Handler:    _WriterGRPC_DeleteParagraphsForChapter_Handler,
		},
		{
			MethodName: "GetParagraphsWithHrefs",
			Handler:    _WriterGRPC_GetParagraphsWithHrefs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
