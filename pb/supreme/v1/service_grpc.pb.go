// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: service.proto

package pb_supreme

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

// SupremeRegulationGRPCClient is the client API for SupremeRegulationGRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SupremeRegulationGRPCClient interface {
	CreateChapter(ctx context.Context, in *CreateChapterRequest, opts ...grpc.CallOption) (*CreateChapterResponse, error)
	CreateParagraphs(ctx context.Context, in *CreateParagraphsRequest, opts ...grpc.CallOption) (*Empty, error)
	CreateRegulation(ctx context.Context, in *CreateRegulationRequest, opts ...grpc.CallOption) (*CreateRegulationResponse, error)
	GenerateLinks(ctx context.Context, in *GenerateLinksRequest, opts ...grpc.CallOption) (*GenerateLinksResponse, error)
	DeleteRegulation(ctx context.Context, in *DeleteRegulationRequest, opts ...grpc.CallOption) (*Empty, error)
}

type supremeRegulationGRPCClient struct {
	cc grpc.ClientConnInterface
}

func NewSupremeRegulationGRPCClient(cc grpc.ClientConnInterface) SupremeRegulationGRPCClient {
	return &supremeRegulationGRPCClient{cc}
}

func (c *supremeRegulationGRPCClient) CreateChapter(ctx context.Context, in *CreateChapterRequest, opts ...grpc.CallOption) (*CreateChapterResponse, error) {
	out := new(CreateChapterResponse)
	err := c.cc.Invoke(ctx, "/supreme.v1.SupremeRegulationGRPC/CreateChapter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supremeRegulationGRPCClient) CreateParagraphs(ctx context.Context, in *CreateParagraphsRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/supreme.v1.SupremeRegulationGRPC/CreateParagraphs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supremeRegulationGRPCClient) CreateRegulation(ctx context.Context, in *CreateRegulationRequest, opts ...grpc.CallOption) (*CreateRegulationResponse, error) {
	out := new(CreateRegulationResponse)
	err := c.cc.Invoke(ctx, "/supreme.v1.SupremeRegulationGRPC/CreateRegulation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supremeRegulationGRPCClient) GenerateLinks(ctx context.Context, in *GenerateLinksRequest, opts ...grpc.CallOption) (*GenerateLinksResponse, error) {
	out := new(GenerateLinksResponse)
	err := c.cc.Invoke(ctx, "/supreme.v1.SupremeRegulationGRPC/GenerateLinks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supremeRegulationGRPCClient) DeleteRegulation(ctx context.Context, in *DeleteRegulationRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/supreme.v1.SupremeRegulationGRPC/DeleteRegulation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SupremeRegulationGRPCServer is the server API for SupremeRegulationGRPC service.
// All implementations must embed UnimplementedSupremeRegulationGRPCServer
// for forward compatibility
type SupremeRegulationGRPCServer interface {
	CreateChapter(context.Context, *CreateChapterRequest) (*CreateChapterResponse, error)
	CreateParagraphs(context.Context, *CreateParagraphsRequest) (*Empty, error)
	CreateRegulation(context.Context, *CreateRegulationRequest) (*CreateRegulationResponse, error)
	GenerateLinks(context.Context, *GenerateLinksRequest) (*GenerateLinksResponse, error)
	DeleteRegulation(context.Context, *DeleteRegulationRequest) (*Empty, error)
	mustEmbedUnimplementedSupremeRegulationGRPCServer()
}

// UnimplementedSupremeRegulationGRPCServer must be embedded to have forward compatible implementations.
type UnimplementedSupremeRegulationGRPCServer struct {
}

func (UnimplementedSupremeRegulationGRPCServer) CreateChapter(context.Context, *CreateChapterRequest) (*CreateChapterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateChapter not implemented")
}
func (UnimplementedSupremeRegulationGRPCServer) CreateParagraphs(context.Context, *CreateParagraphsRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateParagraphs not implemented")
}
func (UnimplementedSupremeRegulationGRPCServer) CreateRegulation(context.Context, *CreateRegulationRequest) (*CreateRegulationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRegulation not implemented")
}
func (UnimplementedSupremeRegulationGRPCServer) GenerateLinks(context.Context, *GenerateLinksRequest) (*GenerateLinksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateLinks not implemented")
}
func (UnimplementedSupremeRegulationGRPCServer) DeleteRegulation(context.Context, *DeleteRegulationRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRegulation not implemented")
}
func (UnimplementedSupremeRegulationGRPCServer) mustEmbedUnimplementedSupremeRegulationGRPCServer() {}

// UnsafeSupremeRegulationGRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SupremeRegulationGRPCServer will
// result in compilation errors.
type UnsafeSupremeRegulationGRPCServer interface {
	mustEmbedUnimplementedSupremeRegulationGRPCServer()
}

func RegisterSupremeRegulationGRPCServer(s grpc.ServiceRegistrar, srv SupremeRegulationGRPCServer) {
	s.RegisterService(&SupremeRegulationGRPC_ServiceDesc, srv)
}

func _SupremeRegulationGRPC_CreateChapter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateChapterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupremeRegulationGRPCServer).CreateChapter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/supreme.v1.SupremeRegulationGRPC/CreateChapter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupremeRegulationGRPCServer).CreateChapter(ctx, req.(*CreateChapterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SupremeRegulationGRPC_CreateParagraphs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateParagraphsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupremeRegulationGRPCServer).CreateParagraphs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/supreme.v1.SupremeRegulationGRPC/CreateParagraphs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupremeRegulationGRPCServer).CreateParagraphs(ctx, req.(*CreateParagraphsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SupremeRegulationGRPC_CreateRegulation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRegulationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupremeRegulationGRPCServer).CreateRegulation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/supreme.v1.SupremeRegulationGRPC/CreateRegulation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupremeRegulationGRPCServer).CreateRegulation(ctx, req.(*CreateRegulationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SupremeRegulationGRPC_GenerateLinks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateLinksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupremeRegulationGRPCServer).GenerateLinks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/supreme.v1.SupremeRegulationGRPC/GenerateLinks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupremeRegulationGRPCServer).GenerateLinks(ctx, req.(*GenerateLinksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SupremeRegulationGRPC_DeleteRegulation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRegulationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupremeRegulationGRPCServer).DeleteRegulation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/supreme.v1.SupremeRegulationGRPC/DeleteRegulation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupremeRegulationGRPCServer).DeleteRegulation(ctx, req.(*DeleteRegulationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SupremeRegulationGRPC_ServiceDesc is the grpc.ServiceDesc for SupremeRegulationGRPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SupremeRegulationGRPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "supreme.v1.SupremeRegulationGRPC",
	HandlerType: (*SupremeRegulationGRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateChapter",
			Handler:    _SupremeRegulationGRPC_CreateChapter_Handler,
		},
		{
			MethodName: "CreateParagraphs",
			Handler:    _SupremeRegulationGRPC_CreateParagraphs_Handler,
		},
		{
			MethodName: "CreateRegulation",
			Handler:    _SupremeRegulationGRPC_CreateRegulation_Handler,
		},
		{
			MethodName: "GenerateLinks",
			Handler:    _SupremeRegulationGRPC_GenerateLinks_Handler,
		},
		{
			MethodName: "DeleteRegulation",
			Handler:    _SupremeRegulationGRPC_DeleteRegulation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
