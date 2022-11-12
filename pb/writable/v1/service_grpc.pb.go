// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: service.proto

package pb_writable

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// WritableRegulationGRPCClient is the client API for WritableRegulationGRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WritableRegulationGRPCClient interface {
	CreateRegulation(ctx context.Context, in *CreateRegulationRequest, opts ...grpc.CallOption) (*CreateRegulationResponse, error)
	DeleteRegulation(ctx context.Context, in *DeleteRegulationRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	CreateChapter(ctx context.Context, in *CreateChapterRequest, opts ...grpc.CallOption) (*CreateChapterResponse, error)
	DeleteChaptersForRegulation(ctx context.Context, in *DeleteChaptersForRegulationRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	CreateAllParagraphs(ctx context.Context, in *CreateAllParagraphsRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	UpdateOneParagraph(ctx context.Context, in *UpdateOneParagraphRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	DeleteParagraphsForChapter(ctx context.Context, in *DeleteParagraphsForChapterRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	GetParagraphsWithHrefs(ctx context.Context, in *GetParagraphsWithHrefsRequest, opts ...grpc.CallOption) (*GetParagraphsWithHrefsResponse, error)
	GetAllChapters(ctx context.Context, in *GetAllChaptersRequest, opts ...grpc.CallOption) (*GetAllChaptersResponse, error)
	GetRegulationIdByChapterId(ctx context.Context, in *GetRegulationIdByChapterIdRequest, opts ...grpc.CallOption) (*GetRegulationIdByChapterIdResponse, error)
}

type writableRegulationGRPCClient struct {
	cc grpc.ClientConnInterface
}

func NewWritableRegulationGRPCClient(cc grpc.ClientConnInterface) WritableRegulationGRPCClient {
	return &writableRegulationGRPCClient{cc}
}

func (c *writableRegulationGRPCClient) CreateRegulation(ctx context.Context, in *CreateRegulationRequest, opts ...grpc.CallOption) (*CreateRegulationResponse, error) {
	out := new(CreateRegulationResponse)
	err := c.cc.Invoke(ctx, "/writable.v1.WritableRegulationGRPC/CreateRegulation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *writableRegulationGRPCClient) DeleteRegulation(ctx context.Context, in *DeleteRegulationRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/writable.v1.WritableRegulationGRPC/DeleteRegulation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *writableRegulationGRPCClient) CreateChapter(ctx context.Context, in *CreateChapterRequest, opts ...grpc.CallOption) (*CreateChapterResponse, error) {
	out := new(CreateChapterResponse)
	err := c.cc.Invoke(ctx, "/writable.v1.WritableRegulationGRPC/CreateChapter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *writableRegulationGRPCClient) DeleteChaptersForRegulation(ctx context.Context, in *DeleteChaptersForRegulationRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/writable.v1.WritableRegulationGRPC/DeleteChaptersForRegulation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *writableRegulationGRPCClient) CreateAllParagraphs(ctx context.Context, in *CreateAllParagraphsRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/writable.v1.WritableRegulationGRPC/CreateAllParagraphs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *writableRegulationGRPCClient) UpdateOneParagraph(ctx context.Context, in *UpdateOneParagraphRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/writable.v1.WritableRegulationGRPC/UpdateOneParagraph", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *writableRegulationGRPCClient) DeleteParagraphsForChapter(ctx context.Context, in *DeleteParagraphsForChapterRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/writable.v1.WritableRegulationGRPC/DeleteParagraphsForChapter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *writableRegulationGRPCClient) GetParagraphsWithHrefs(ctx context.Context, in *GetParagraphsWithHrefsRequest, opts ...grpc.CallOption) (*GetParagraphsWithHrefsResponse, error) {
	out := new(GetParagraphsWithHrefsResponse)
	err := c.cc.Invoke(ctx, "/writable.v1.WritableRegulationGRPC/GetParagraphsWithHrefs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *writableRegulationGRPCClient) GetAllChapters(ctx context.Context, in *GetAllChaptersRequest, opts ...grpc.CallOption) (*GetAllChaptersResponse, error) {
	out := new(GetAllChaptersResponse)
	err := c.cc.Invoke(ctx, "/writable.v1.WritableRegulationGRPC/GetAllChapters", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *writableRegulationGRPCClient) GetRegulationIdByChapterId(ctx context.Context, in *GetRegulationIdByChapterIdRequest, opts ...grpc.CallOption) (*GetRegulationIdByChapterIdResponse, error) {
	out := new(GetRegulationIdByChapterIdResponse)
	err := c.cc.Invoke(ctx, "/writable.v1.WritableRegulationGRPC/GetRegulationIdByChapterId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WritableRegulationGRPCServer is the server API for WritableRegulationGRPC service.
// All implementations must embed UnimplementedWritableRegulationGRPCServer
// for forward compatibility
type WritableRegulationGRPCServer interface {
	CreateRegulation(context.Context, *CreateRegulationRequest) (*CreateRegulationResponse, error)
	DeleteRegulation(context.Context, *DeleteRegulationRequest) (*empty.Empty, error)
	CreateChapter(context.Context, *CreateChapterRequest) (*CreateChapterResponse, error)
	DeleteChaptersForRegulation(context.Context, *DeleteChaptersForRegulationRequest) (*empty.Empty, error)
	CreateAllParagraphs(context.Context, *CreateAllParagraphsRequest) (*empty.Empty, error)
	UpdateOneParagraph(context.Context, *UpdateOneParagraphRequest) (*empty.Empty, error)
	DeleteParagraphsForChapter(context.Context, *DeleteParagraphsForChapterRequest) (*empty.Empty, error)
	GetParagraphsWithHrefs(context.Context, *GetParagraphsWithHrefsRequest) (*GetParagraphsWithHrefsResponse, error)
	GetAllChapters(context.Context, *GetAllChaptersRequest) (*GetAllChaptersResponse, error)
	GetRegulationIdByChapterId(context.Context, *GetRegulationIdByChapterIdRequest) (*GetRegulationIdByChapterIdResponse, error)
	mustEmbedUnimplementedWritableRegulationGRPCServer()
}

// UnimplementedWritableRegulationGRPCServer must be embedded to have forward compatible implementations.
type UnimplementedWritableRegulationGRPCServer struct {
}

func (UnimplementedWritableRegulationGRPCServer) CreateRegulation(context.Context, *CreateRegulationRequest) (*CreateRegulationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRegulation not implemented")
}
func (UnimplementedWritableRegulationGRPCServer) DeleteRegulation(context.Context, *DeleteRegulationRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRegulation not implemented")
}
func (UnimplementedWritableRegulationGRPCServer) CreateChapter(context.Context, *CreateChapterRequest) (*CreateChapterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateChapter not implemented")
}
func (UnimplementedWritableRegulationGRPCServer) DeleteChaptersForRegulation(context.Context, *DeleteChaptersForRegulationRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteChaptersForRegulation not implemented")
}
func (UnimplementedWritableRegulationGRPCServer) CreateAllParagraphs(context.Context, *CreateAllParagraphsRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAllParagraphs not implemented")
}
func (UnimplementedWritableRegulationGRPCServer) UpdateOneParagraph(context.Context, *UpdateOneParagraphRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOneParagraph not implemented")
}
func (UnimplementedWritableRegulationGRPCServer) DeleteParagraphsForChapter(context.Context, *DeleteParagraphsForChapterRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteParagraphsForChapter not implemented")
}
func (UnimplementedWritableRegulationGRPCServer) GetParagraphsWithHrefs(context.Context, *GetParagraphsWithHrefsRequest) (*GetParagraphsWithHrefsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetParagraphsWithHrefs not implemented")
}
func (UnimplementedWritableRegulationGRPCServer) GetAllChapters(context.Context, *GetAllChaptersRequest) (*GetAllChaptersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllChapters not implemented")
}
func (UnimplementedWritableRegulationGRPCServer) GetRegulationIdByChapterId(context.Context, *GetRegulationIdByChapterIdRequest) (*GetRegulationIdByChapterIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRegulationIdByChapterId not implemented")
}
func (UnimplementedWritableRegulationGRPCServer) mustEmbedUnimplementedWritableRegulationGRPCServer() {
}

// UnsafeWritableRegulationGRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WritableRegulationGRPCServer will
// result in compilation errors.
type UnsafeWritableRegulationGRPCServer interface {
	mustEmbedUnimplementedWritableRegulationGRPCServer()
}

func RegisterWritableRegulationGRPCServer(s grpc.ServiceRegistrar, srv WritableRegulationGRPCServer) {
	s.RegisterService(&WritableRegulationGRPC_ServiceDesc, srv)
}

func _WritableRegulationGRPC_CreateRegulation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRegulationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WritableRegulationGRPCServer).CreateRegulation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/writable.v1.WritableRegulationGRPC/CreateRegulation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WritableRegulationGRPCServer).CreateRegulation(ctx, req.(*CreateRegulationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WritableRegulationGRPC_DeleteRegulation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRegulationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WritableRegulationGRPCServer).DeleteRegulation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/writable.v1.WritableRegulationGRPC/DeleteRegulation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WritableRegulationGRPCServer).DeleteRegulation(ctx, req.(*DeleteRegulationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WritableRegulationGRPC_CreateChapter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateChapterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WritableRegulationGRPCServer).CreateChapter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/writable.v1.WritableRegulationGRPC/CreateChapter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WritableRegulationGRPCServer).CreateChapter(ctx, req.(*CreateChapterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WritableRegulationGRPC_DeleteChaptersForRegulation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteChaptersForRegulationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WritableRegulationGRPCServer).DeleteChaptersForRegulation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/writable.v1.WritableRegulationGRPC/DeleteChaptersForRegulation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WritableRegulationGRPCServer).DeleteChaptersForRegulation(ctx, req.(*DeleteChaptersForRegulationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WritableRegulationGRPC_CreateAllParagraphs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAllParagraphsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WritableRegulationGRPCServer).CreateAllParagraphs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/writable.v1.WritableRegulationGRPC/CreateAllParagraphs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WritableRegulationGRPCServer).CreateAllParagraphs(ctx, req.(*CreateAllParagraphsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WritableRegulationGRPC_UpdateOneParagraph_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOneParagraphRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WritableRegulationGRPCServer).UpdateOneParagraph(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/writable.v1.WritableRegulationGRPC/UpdateOneParagraph",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WritableRegulationGRPCServer).UpdateOneParagraph(ctx, req.(*UpdateOneParagraphRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WritableRegulationGRPC_DeleteParagraphsForChapter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteParagraphsForChapterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WritableRegulationGRPCServer).DeleteParagraphsForChapter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/writable.v1.WritableRegulationGRPC/DeleteParagraphsForChapter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WritableRegulationGRPCServer).DeleteParagraphsForChapter(ctx, req.(*DeleteParagraphsForChapterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WritableRegulationGRPC_GetParagraphsWithHrefs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetParagraphsWithHrefsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WritableRegulationGRPCServer).GetParagraphsWithHrefs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/writable.v1.WritableRegulationGRPC/GetParagraphsWithHrefs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WritableRegulationGRPCServer).GetParagraphsWithHrefs(ctx, req.(*GetParagraphsWithHrefsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WritableRegulationGRPC_GetAllChapters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllChaptersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WritableRegulationGRPCServer).GetAllChapters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/writable.v1.WritableRegulationGRPC/GetAllChapters",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WritableRegulationGRPCServer).GetAllChapters(ctx, req.(*GetAllChaptersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WritableRegulationGRPC_GetRegulationIdByChapterId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRegulationIdByChapterIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WritableRegulationGRPCServer).GetRegulationIdByChapterId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/writable.v1.WritableRegulationGRPC/GetRegulationIdByChapterId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WritableRegulationGRPCServer).GetRegulationIdByChapterId(ctx, req.(*GetRegulationIdByChapterIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WritableRegulationGRPC_ServiceDesc is the grpc.ServiceDesc for WritableRegulationGRPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WritableRegulationGRPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "writable.v1.WritableRegulationGRPC",
	HandlerType: (*WritableRegulationGRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRegulation",
			Handler:    _WritableRegulationGRPC_CreateRegulation_Handler,
		},
		{
			MethodName: "DeleteRegulation",
			Handler:    _WritableRegulationGRPC_DeleteRegulation_Handler,
		},
		{
			MethodName: "CreateChapter",
			Handler:    _WritableRegulationGRPC_CreateChapter_Handler,
		},
		{
			MethodName: "DeleteChaptersForRegulation",
			Handler:    _WritableRegulationGRPC_DeleteChaptersForRegulation_Handler,
		},
		{
			MethodName: "CreateAllParagraphs",
			Handler:    _WritableRegulationGRPC_CreateAllParagraphs_Handler,
		},
		{
			MethodName: "UpdateOneParagraph",
			Handler:    _WritableRegulationGRPC_UpdateOneParagraph_Handler,
		},
		{
			MethodName: "DeleteParagraphsForChapter",
			Handler:    _WritableRegulationGRPC_DeleteParagraphsForChapter_Handler,
		},
		{
			MethodName: "GetParagraphsWithHrefs",
			Handler:    _WritableRegulationGRPC_GetParagraphsWithHrefs_Handler,
		},
		{
			MethodName: "GetAllChapters",
			Handler:    _WritableRegulationGRPC_GetAllChapters_Handler,
		},
		{
			MethodName: "GetRegulationIdByChapterId",
			Handler:    _WritableRegulationGRPC_GetRegulationIdByChapterId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
