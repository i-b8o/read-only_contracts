// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: service.proto

package pb_searcher

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

// SearcherGRPCClient is the client API for SearcherGRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SearcherGRPCClient interface {
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponseMessage, error)
	SearchDocs(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponseMessage, error)
	SearchChapters(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponseMessage, error)
	SearchPargaraphs(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponseMessage, error)
}

type searcherGRPCClient struct {
	cc grpc.ClientConnInterface
}

func NewSearcherGRPCClient(cc grpc.ClientConnInterface) SearcherGRPCClient {
	return &searcherGRPCClient{cc}
}

func (c *searcherGRPCClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponseMessage, error) {
	out := new(SearchResponseMessage)
	err := c.cc.Invoke(ctx, "/searcher.v1.SearcherGRPC/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searcherGRPCClient) SearchDocs(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponseMessage, error) {
	out := new(SearchResponseMessage)
	err := c.cc.Invoke(ctx, "/searcher.v1.SearcherGRPC/SearchDocs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searcherGRPCClient) SearchChapters(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponseMessage, error) {
	out := new(SearchResponseMessage)
	err := c.cc.Invoke(ctx, "/searcher.v1.SearcherGRPC/SearchChapters", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searcherGRPCClient) SearchPargaraphs(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponseMessage, error) {
	out := new(SearchResponseMessage)
	err := c.cc.Invoke(ctx, "/searcher.v1.SearcherGRPC/SearchPargaraphs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SearcherGRPCServer is the server API for SearcherGRPC service.
// All implementations must embed UnimplementedSearcherGRPCServer
// for forward compatibility
type SearcherGRPCServer interface {
	Search(context.Context, *SearchRequest) (*SearchResponseMessage, error)
	SearchDocs(context.Context, *SearchRequest) (*SearchResponseMessage, error)
	SearchChapters(context.Context, *SearchRequest) (*SearchResponseMessage, error)
	SearchPargaraphs(context.Context, *SearchRequest) (*SearchResponseMessage, error)
	mustEmbedUnimplementedSearcherGRPCServer()
}

// UnimplementedSearcherGRPCServer must be embedded to have forward compatible implementations.
type UnimplementedSearcherGRPCServer struct {
}

func (UnimplementedSearcherGRPCServer) Search(context.Context, *SearchRequest) (*SearchResponseMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedSearcherGRPCServer) SearchDocs(context.Context, *SearchRequest) (*SearchResponseMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchDocs not implemented")
}
func (UnimplementedSearcherGRPCServer) SearchChapters(context.Context, *SearchRequest) (*SearchResponseMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchChapters not implemented")
}
func (UnimplementedSearcherGRPCServer) SearchPargaraphs(context.Context, *SearchRequest) (*SearchResponseMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchPargaraphs not implemented")
}
func (UnimplementedSearcherGRPCServer) mustEmbedUnimplementedSearcherGRPCServer() {}

// UnsafeSearcherGRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SearcherGRPCServer will
// result in compilation errors.
type UnsafeSearcherGRPCServer interface {
	mustEmbedUnimplementedSearcherGRPCServer()
}

func RegisterSearcherGRPCServer(s grpc.ServiceRegistrar, srv SearcherGRPCServer) {
	s.RegisterService(&SearcherGRPC_ServiceDesc, srv)
}

func _SearcherGRPC_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearcherGRPCServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/searcher.v1.SearcherGRPC/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearcherGRPCServer).Search(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearcherGRPC_SearchDocs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearcherGRPCServer).SearchDocs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/searcher.v1.SearcherGRPC/SearchDocs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearcherGRPCServer).SearchDocs(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearcherGRPC_SearchChapters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearcherGRPCServer).SearchChapters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/searcher.v1.SearcherGRPC/SearchChapters",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearcherGRPCServer).SearchChapters(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearcherGRPC_SearchPargaraphs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearcherGRPCServer).SearchPargaraphs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/searcher.v1.SearcherGRPC/SearchPargaraphs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearcherGRPCServer).SearchPargaraphs(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SearcherGRPC_ServiceDesc is the grpc.ServiceDesc for SearcherGRPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SearcherGRPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "searcher.v1.SearcherGRPC",
	HandlerType: (*SearcherGRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Search",
			Handler:    _SearcherGRPC_Search_Handler,
		},
		{
			MethodName: "SearchDocs",
			Handler:    _SearcherGRPC_SearchDocs_Handler,
		},
		{
			MethodName: "SearchChapters",
			Handler:    _SearcherGRPC_SearchChapters_Handler,
		},
		{
			MethodName: "SearchPargaraphs",
			Handler:    _SearcherGRPC_SearchPargaraphs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
