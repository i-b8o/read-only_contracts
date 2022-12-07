// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
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

// WriterRegulationGRPCClient is the client API for WriterRegulationGRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WriterRegulationGRPCClient interface {
	Create(ctx context.Context, in *CreateRegulationRequest, opts ...grpc.CallOption) (*CreateRegulationResponse, error)
	Delete(ctx context.Context, in *DeleteRegulationRequest, opts ...grpc.CallOption) (*Empty, error)
	GetAll(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetRegulationsResponse, error)
}

type writerRegulationGRPCClient struct {
	cc grpc.ClientConnInterface
}

func NewWriterRegulationGRPCClient(cc grpc.ClientConnInterface) WriterRegulationGRPCClient {
	return &writerRegulationGRPCClient{cc}
}

func (c *writerRegulationGRPCClient) Create(ctx context.Context, in *CreateRegulationRequest, opts ...grpc.CallOption) (*CreateRegulationResponse, error) {
	out := new(CreateRegulationResponse)
	err := c.cc.Invoke(ctx, "/writer.v1.WriterRegulationGRPC/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *writerRegulationGRPCClient) Delete(ctx context.Context, in *DeleteRegulationRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/writer.v1.WriterRegulationGRPC/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *writerRegulationGRPCClient) GetAll(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetRegulationsResponse, error) {
	out := new(GetRegulationsResponse)
	err := c.cc.Invoke(ctx, "/writer.v1.WriterRegulationGRPC/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WriterRegulationGRPCServer is the server API for WriterRegulationGRPC service.
// All implementations must embed UnimplementedWriterRegulationGRPCServer
// for forward compatibility
type WriterRegulationGRPCServer interface {
	Create(context.Context, *CreateRegulationRequest) (*CreateRegulationResponse, error)
	Delete(context.Context, *DeleteRegulationRequest) (*Empty, error)
	GetAll(context.Context, *Empty) (*GetRegulationsResponse, error)
	mustEmbedUnimplementedWriterRegulationGRPCServer()
}

// UnimplementedWriterRegulationGRPCServer must be embedded to have forward compatible implementations.
type UnimplementedWriterRegulationGRPCServer struct {
}

func (UnimplementedWriterRegulationGRPCServer) Create(context.Context, *CreateRegulationRequest) (*CreateRegulationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedWriterRegulationGRPCServer) Delete(context.Context, *DeleteRegulationRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedWriterRegulationGRPCServer) GetAll(context.Context, *Empty) (*GetRegulationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedWriterRegulationGRPCServer) mustEmbedUnimplementedWriterRegulationGRPCServer() {}

// UnsafeWriterRegulationGRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WriterRegulationGRPCServer will
// result in compilation errors.
type UnsafeWriterRegulationGRPCServer interface {
	mustEmbedUnimplementedWriterRegulationGRPCServer()
}

func RegisterWriterRegulationGRPCServer(s grpc.ServiceRegistrar, srv WriterRegulationGRPCServer) {
	s.RegisterService(&WriterRegulationGRPC_ServiceDesc, srv)
}

func _WriterRegulationGRPC_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRegulationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WriterRegulationGRPCServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/writer.v1.WriterRegulationGRPC/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WriterRegulationGRPCServer).Create(ctx, req.(*CreateRegulationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WriterRegulationGRPC_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRegulationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WriterRegulationGRPCServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/writer.v1.WriterRegulationGRPC/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WriterRegulationGRPCServer).Delete(ctx, req.(*DeleteRegulationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WriterRegulationGRPC_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WriterRegulationGRPCServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/writer.v1.WriterRegulationGRPC/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WriterRegulationGRPCServer).GetAll(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// WriterRegulationGRPC_ServiceDesc is the grpc.ServiceDesc for WriterRegulationGRPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WriterRegulationGRPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "writer.v1.WriterRegulationGRPC",
	HandlerType: (*WriterRegulationGRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _WriterRegulationGRPC_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _WriterRegulationGRPC_Delete_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _WriterRegulationGRPC_GetAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}

// WriterChapterGRPCClient is the client API for WriterChapterGRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WriterChapterGRPCClient interface {
	Create(ctx context.Context, in *CreateChapterRequest, opts ...grpc.CallOption) (*CreateChapterResponse, error)
	GetAll(ctx context.Context, in *GetAllChaptersIdsRequest, opts ...grpc.CallOption) (*GetAllChaptersIdsResponse, error)
	GetRegulationId(ctx context.Context, in *GetRegulationIdByChapterIdRequest, opts ...grpc.CallOption) (*GetRegulationIdByChapterIdResponse, error)
}

type writerChapterGRPCClient struct {
	cc grpc.ClientConnInterface
}

func NewWriterChapterGRPCClient(cc grpc.ClientConnInterface) WriterChapterGRPCClient {
	return &writerChapterGRPCClient{cc}
}

func (c *writerChapterGRPCClient) Create(ctx context.Context, in *CreateChapterRequest, opts ...grpc.CallOption) (*CreateChapterResponse, error) {
	out := new(CreateChapterResponse)
	err := c.cc.Invoke(ctx, "/writer.v1.WriterChapterGRPC/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *writerChapterGRPCClient) GetAll(ctx context.Context, in *GetAllChaptersIdsRequest, opts ...grpc.CallOption) (*GetAllChaptersIdsResponse, error) {
	out := new(GetAllChaptersIdsResponse)
	err := c.cc.Invoke(ctx, "/writer.v1.WriterChapterGRPC/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *writerChapterGRPCClient) GetRegulationId(ctx context.Context, in *GetRegulationIdByChapterIdRequest, opts ...grpc.CallOption) (*GetRegulationIdByChapterIdResponse, error) {
	out := new(GetRegulationIdByChapterIdResponse)
	err := c.cc.Invoke(ctx, "/writer.v1.WriterChapterGRPC/GetRegulationId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WriterChapterGRPCServer is the server API for WriterChapterGRPC service.
// All implementations must embed UnimplementedWriterChapterGRPCServer
// for forward compatibility
type WriterChapterGRPCServer interface {
	Create(context.Context, *CreateChapterRequest) (*CreateChapterResponse, error)
	GetAll(context.Context, *GetAllChaptersIdsRequest) (*GetAllChaptersIdsResponse, error)
	GetRegulationId(context.Context, *GetRegulationIdByChapterIdRequest) (*GetRegulationIdByChapterIdResponse, error)
	mustEmbedUnimplementedWriterChapterGRPCServer()
}

// UnimplementedWriterChapterGRPCServer must be embedded to have forward compatible implementations.
type UnimplementedWriterChapterGRPCServer struct {
}

func (UnimplementedWriterChapterGRPCServer) Create(context.Context, *CreateChapterRequest) (*CreateChapterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedWriterChapterGRPCServer) GetAll(context.Context, *GetAllChaptersIdsRequest) (*GetAllChaptersIdsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedWriterChapterGRPCServer) GetRegulationId(context.Context, *GetRegulationIdByChapterIdRequest) (*GetRegulationIdByChapterIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRegulationId not implemented")
}
func (UnimplementedWriterChapterGRPCServer) mustEmbedUnimplementedWriterChapterGRPCServer() {}

// UnsafeWriterChapterGRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WriterChapterGRPCServer will
// result in compilation errors.
type UnsafeWriterChapterGRPCServer interface {
	mustEmbedUnimplementedWriterChapterGRPCServer()
}

func RegisterWriterChapterGRPCServer(s grpc.ServiceRegistrar, srv WriterChapterGRPCServer) {
	s.RegisterService(&WriterChapterGRPC_ServiceDesc, srv)
}

func _WriterChapterGRPC_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateChapterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WriterChapterGRPCServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/writer.v1.WriterChapterGRPC/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WriterChapterGRPCServer).Create(ctx, req.(*CreateChapterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WriterChapterGRPC_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllChaptersIdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WriterChapterGRPCServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/writer.v1.WriterChapterGRPC/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WriterChapterGRPCServer).GetAll(ctx, req.(*GetAllChaptersIdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WriterChapterGRPC_GetRegulationId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRegulationIdByChapterIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WriterChapterGRPCServer).GetRegulationId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/writer.v1.WriterChapterGRPC/GetRegulationId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WriterChapterGRPCServer).GetRegulationId(ctx, req.(*GetRegulationIdByChapterIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WriterChapterGRPC_ServiceDesc is the grpc.ServiceDesc for WriterChapterGRPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WriterChapterGRPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "writer.v1.WriterChapterGRPC",
	HandlerType: (*WriterChapterGRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _WriterChapterGRPC_Create_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _WriterChapterGRPC_GetAll_Handler,
		},
		{
			MethodName: "GetRegulationId",
			Handler:    _WriterChapterGRPC_GetRegulationId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}

// WriterParagraphGRPCClient is the client API for WriterParagraphGRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WriterParagraphGRPCClient interface {
	GetOne(ctx context.Context, in *GetOneParagraphRequest, opts ...grpc.CallOption) (*GetOneParagraphResponse, error)
	CreateAll(ctx context.Context, in *CreateAllParagraphsRequest, opts ...grpc.CallOption) (*Empty, error)
	Update(ctx context.Context, in *UpdateOneParagraphRequest, opts ...grpc.CallOption) (*Empty, error)
	GetWithHrefs(ctx context.Context, in *GetParagraphsWithHrefsRequest, opts ...grpc.CallOption) (*GetParagraphsWithHrefsResponse, error)
}

type writerParagraphGRPCClient struct {
	cc grpc.ClientConnInterface
}

func NewWriterParagraphGRPCClient(cc grpc.ClientConnInterface) WriterParagraphGRPCClient {
	return &writerParagraphGRPCClient{cc}
}

func (c *writerParagraphGRPCClient) GetOne(ctx context.Context, in *GetOneParagraphRequest, opts ...grpc.CallOption) (*GetOneParagraphResponse, error) {
	out := new(GetOneParagraphResponse)
	err := c.cc.Invoke(ctx, "/writer.v1.WriterParagraphGRPC/GetOne", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *writerParagraphGRPCClient) CreateAll(ctx context.Context, in *CreateAllParagraphsRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/writer.v1.WriterParagraphGRPC/CreateAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *writerParagraphGRPCClient) Update(ctx context.Context, in *UpdateOneParagraphRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/writer.v1.WriterParagraphGRPC/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *writerParagraphGRPCClient) GetWithHrefs(ctx context.Context, in *GetParagraphsWithHrefsRequest, opts ...grpc.CallOption) (*GetParagraphsWithHrefsResponse, error) {
	out := new(GetParagraphsWithHrefsResponse)
	err := c.cc.Invoke(ctx, "/writer.v1.WriterParagraphGRPC/GetWithHrefs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WriterParagraphGRPCServer is the server API for WriterParagraphGRPC service.
// All implementations must embed UnimplementedWriterParagraphGRPCServer
// for forward compatibility
type WriterParagraphGRPCServer interface {
	GetOne(context.Context, *GetOneParagraphRequest) (*GetOneParagraphResponse, error)
	CreateAll(context.Context, *CreateAllParagraphsRequest) (*Empty, error)
	Update(context.Context, *UpdateOneParagraphRequest) (*Empty, error)
	GetWithHrefs(context.Context, *GetParagraphsWithHrefsRequest) (*GetParagraphsWithHrefsResponse, error)
	mustEmbedUnimplementedWriterParagraphGRPCServer()
}

// UnimplementedWriterParagraphGRPCServer must be embedded to have forward compatible implementations.
type UnimplementedWriterParagraphGRPCServer struct {
}

func (UnimplementedWriterParagraphGRPCServer) GetOne(context.Context, *GetOneParagraphRequest) (*GetOneParagraphResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOne not implemented")
}
func (UnimplementedWriterParagraphGRPCServer) CreateAll(context.Context, *CreateAllParagraphsRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAll not implemented")
}
func (UnimplementedWriterParagraphGRPCServer) Update(context.Context, *UpdateOneParagraphRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedWriterParagraphGRPCServer) GetWithHrefs(context.Context, *GetParagraphsWithHrefsRequest) (*GetParagraphsWithHrefsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWithHrefs not implemented")
}
func (UnimplementedWriterParagraphGRPCServer) mustEmbedUnimplementedWriterParagraphGRPCServer() {}

// UnsafeWriterParagraphGRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WriterParagraphGRPCServer will
// result in compilation errors.
type UnsafeWriterParagraphGRPCServer interface {
	mustEmbedUnimplementedWriterParagraphGRPCServer()
}

func RegisterWriterParagraphGRPCServer(s grpc.ServiceRegistrar, srv WriterParagraphGRPCServer) {
	s.RegisterService(&WriterParagraphGRPC_ServiceDesc, srv)
}

func _WriterParagraphGRPC_GetOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOneParagraphRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WriterParagraphGRPCServer).GetOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/writer.v1.WriterParagraphGRPC/GetOne",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WriterParagraphGRPCServer).GetOne(ctx, req.(*GetOneParagraphRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WriterParagraphGRPC_CreateAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAllParagraphsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WriterParagraphGRPCServer).CreateAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/writer.v1.WriterParagraphGRPC/CreateAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WriterParagraphGRPCServer).CreateAll(ctx, req.(*CreateAllParagraphsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WriterParagraphGRPC_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOneParagraphRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WriterParagraphGRPCServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/writer.v1.WriterParagraphGRPC/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WriterParagraphGRPCServer).Update(ctx, req.(*UpdateOneParagraphRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WriterParagraphGRPC_GetWithHrefs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetParagraphsWithHrefsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WriterParagraphGRPCServer).GetWithHrefs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/writer.v1.WriterParagraphGRPC/GetWithHrefs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WriterParagraphGRPCServer).GetWithHrefs(ctx, req.(*GetParagraphsWithHrefsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WriterParagraphGRPC_ServiceDesc is the grpc.ServiceDesc for WriterParagraphGRPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WriterParagraphGRPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "writer.v1.WriterParagraphGRPC",
	HandlerType: (*WriterParagraphGRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOne",
			Handler:    _WriterParagraphGRPC_GetOne_Handler,
		},
		{
			MethodName: "CreateAll",
			Handler:    _WriterParagraphGRPC_CreateAll_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _WriterParagraphGRPC_Update_Handler,
		},
		{
			MethodName: "GetWithHrefs",
			Handler:    _WriterParagraphGRPC_GetWithHrefs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
