// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: master_service.proto

package pb_master

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

// MasterGRPCClient is the client API for MasterGRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MasterGRPCClient interface {
	CreateChapter(ctx context.Context, in *CreateChapterRequest, opts ...grpc.CallOption) (*CreateChapterResponse, error)
	CreateParagraphs(ctx context.Context, in *CreateParagraphsRequest, opts ...grpc.CallOption) (*Empty, error)
	CreateRegulation(ctx context.Context, in *CreateRegulationRequest, opts ...grpc.CallOption) (*CreateRegulationResponse, error)
	GenerateLinks(ctx context.Context, in *GenerateLinksRequest, opts ...grpc.CallOption) (*GenerateLinksResponse, error)
	DeleteRegulation(ctx context.Context, in *DeleteRegulationRequest, opts ...grpc.CallOption) (*Empty, error)
}

type masterGRPCClient struct {
	cc grpc.ClientConnInterface
}

func NewMasterGRPCClient(cc grpc.ClientConnInterface) MasterGRPCClient {
	return &masterGRPCClient{cc}
}

func (c *masterGRPCClient) CreateChapter(ctx context.Context, in *CreateChapterRequest, opts ...grpc.CallOption) (*CreateChapterResponse, error) {
	out := new(CreateChapterResponse)
	err := c.cc.Invoke(ctx, "/master.v1.MasterGRPC/CreateChapter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterGRPCClient) CreateParagraphs(ctx context.Context, in *CreateParagraphsRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/master.v1.MasterGRPC/CreateParagraphs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterGRPCClient) CreateRegulation(ctx context.Context, in *CreateRegulationRequest, opts ...grpc.CallOption) (*CreateRegulationResponse, error) {
	out := new(CreateRegulationResponse)
	err := c.cc.Invoke(ctx, "/master.v1.MasterGRPC/CreateRegulation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterGRPCClient) GenerateLinks(ctx context.Context, in *GenerateLinksRequest, opts ...grpc.CallOption) (*GenerateLinksResponse, error) {
	out := new(GenerateLinksResponse)
	err := c.cc.Invoke(ctx, "/master.v1.MasterGRPC/GenerateLinks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterGRPCClient) DeleteRegulation(ctx context.Context, in *DeleteRegulationRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/master.v1.MasterGRPC/DeleteRegulation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MasterGRPCServer is the server API for MasterGRPC service.
// All implementations must embed UnimplementedMasterGRPCServer
// for forward compatibility
type MasterGRPCServer interface {
	CreateChapter(context.Context, *CreateChapterRequest) (*CreateChapterResponse, error)
	CreateParagraphs(context.Context, *CreateParagraphsRequest) (*Empty, error)
	CreateRegulation(context.Context, *CreateRegulationRequest) (*CreateRegulationResponse, error)
	GenerateLinks(context.Context, *GenerateLinksRequest) (*GenerateLinksResponse, error)
	DeleteRegulation(context.Context, *DeleteRegulationRequest) (*Empty, error)
	mustEmbedUnimplementedMasterGRPCServer()
}

// UnimplementedMasterGRPCServer must be embedded to have forward compatible implementations.
type UnimplementedMasterGRPCServer struct {
}

func (UnimplementedMasterGRPCServer) CreateChapter(context.Context, *CreateChapterRequest) (*CreateChapterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateChapter not implemented")
}
func (UnimplementedMasterGRPCServer) CreateParagraphs(context.Context, *CreateParagraphsRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateParagraphs not implemented")
}
func (UnimplementedMasterGRPCServer) CreateRegulation(context.Context, *CreateRegulationRequest) (*CreateRegulationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRegulation not implemented")
}
func (UnimplementedMasterGRPCServer) GenerateLinks(context.Context, *GenerateLinksRequest) (*GenerateLinksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateLinks not implemented")
}
func (UnimplementedMasterGRPCServer) DeleteRegulation(context.Context, *DeleteRegulationRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRegulation not implemented")
}
func (UnimplementedMasterGRPCServer) mustEmbedUnimplementedMasterGRPCServer() {}

// UnsafeMasterGRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MasterGRPCServer will
// result in compilation errors.
type UnsafeMasterGRPCServer interface {
	mustEmbedUnimplementedMasterGRPCServer()
}

func RegisterMasterGRPCServer(s grpc.ServiceRegistrar, srv MasterGRPCServer) {
	s.RegisterService(&MasterGRPC_ServiceDesc, srv)
}

func _MasterGRPC_CreateChapter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateChapterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterGRPCServer).CreateChapter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/master.v1.MasterGRPC/CreateChapter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterGRPCServer).CreateChapter(ctx, req.(*CreateChapterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterGRPC_CreateParagraphs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateParagraphsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterGRPCServer).CreateParagraphs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/master.v1.MasterGRPC/CreateParagraphs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterGRPCServer).CreateParagraphs(ctx, req.(*CreateParagraphsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterGRPC_CreateRegulation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRegulationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterGRPCServer).CreateRegulation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/master.v1.MasterGRPC/CreateRegulation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterGRPCServer).CreateRegulation(ctx, req.(*CreateRegulationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterGRPC_GenerateLinks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateLinksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterGRPCServer).GenerateLinks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/master.v1.MasterGRPC/GenerateLinks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterGRPCServer).GenerateLinks(ctx, req.(*GenerateLinksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterGRPC_DeleteRegulation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRegulationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterGRPCServer).DeleteRegulation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/master.v1.MasterGRPC/DeleteRegulation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterGRPCServer).DeleteRegulation(ctx, req.(*DeleteRegulationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MasterGRPC_ServiceDesc is the grpc.ServiceDesc for MasterGRPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MasterGRPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "master.v1.MasterGRPC",
	HandlerType: (*MasterGRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateChapter",
			Handler:    _MasterGRPC_CreateChapter_Handler,
		},
		{
			MethodName: "CreateParagraphs",
			Handler:    _MasterGRPC_CreateParagraphs_Handler,
		},
		{
			MethodName: "CreateRegulation",
			Handler:    _MasterGRPC_CreateRegulation_Handler,
		},
		{
			MethodName: "GenerateLinks",
			Handler:    _MasterGRPC_GenerateLinks_Handler,
		},
		{
			MethodName: "DeleteRegulation",
			Handler:    _MasterGRPC_DeleteRegulation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "master_service.proto",
}