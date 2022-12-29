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

// MasterDocGRPCClient is the client API for MasterDocGRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MasterDocGRPCClient interface {
	Exist(ctx context.Context, in *ExistRequest, opts ...grpc.CallOption) (*ExistResponse, error)
	Create(ctx context.Context, in *CreateDocRequest, opts ...grpc.CallOption) (*CreateDocResponse, error)
	GetAll(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetAllDocsResponse, error)
	Delete(ctx context.Context, in *DeleteDocRequest, opts ...grpc.CallOption) (*Empty, error)
	UpdateLinks(ctx context.Context, in *UpdateLinksRequest, opts ...grpc.CallOption) (*UpdateLinksResponse, error)
	GetAbsents(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetAbsentsResponse, error)
}

type masterDocGRPCClient struct {
	cc grpc.ClientConnInterface
}

func NewMasterDocGRPCClient(cc grpc.ClientConnInterface) MasterDocGRPCClient {
	return &masterDocGRPCClient{cc}
}

func (c *masterDocGRPCClient) Exist(ctx context.Context, in *ExistRequest, opts ...grpc.CallOption) (*ExistResponse, error) {
	out := new(ExistResponse)
	err := c.cc.Invoke(ctx, "/master.v1.MasterDocGRPC/Exist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterDocGRPCClient) Create(ctx context.Context, in *CreateDocRequest, opts ...grpc.CallOption) (*CreateDocResponse, error) {
	out := new(CreateDocResponse)
	err := c.cc.Invoke(ctx, "/master.v1.MasterDocGRPC/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterDocGRPCClient) GetAll(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetAllDocsResponse, error) {
	out := new(GetAllDocsResponse)
	err := c.cc.Invoke(ctx, "/master.v1.MasterDocGRPC/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterDocGRPCClient) Delete(ctx context.Context, in *DeleteDocRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/master.v1.MasterDocGRPC/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterDocGRPCClient) UpdateLinks(ctx context.Context, in *UpdateLinksRequest, opts ...grpc.CallOption) (*UpdateLinksResponse, error) {
	out := new(UpdateLinksResponse)
	err := c.cc.Invoke(ctx, "/master.v1.MasterDocGRPC/UpdateLinks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterDocGRPCClient) GetAbsents(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetAbsentsResponse, error) {
	out := new(GetAbsentsResponse)
	err := c.cc.Invoke(ctx, "/master.v1.MasterDocGRPC/GetAbsents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MasterDocGRPCServer is the server API for MasterDocGRPC service.
// All implementations must embed UnimplementedMasterDocGRPCServer
// for forward compatibility
type MasterDocGRPCServer interface {
	Exist(context.Context, *ExistRequest) (*ExistResponse, error)
	Create(context.Context, *CreateDocRequest) (*CreateDocResponse, error)
	GetAll(context.Context, *Empty) (*GetAllDocsResponse, error)
	Delete(context.Context, *DeleteDocRequest) (*Empty, error)
	UpdateLinks(context.Context, *UpdateLinksRequest) (*UpdateLinksResponse, error)
	GetAbsents(context.Context, *Empty) (*GetAbsentsResponse, error)
	mustEmbedUnimplementedMasterDocGRPCServer()
}

// UnimplementedMasterDocGRPCServer must be embedded to have forward compatible implementations.
type UnimplementedMasterDocGRPCServer struct {
}

func (UnimplementedMasterDocGRPCServer) Exist(context.Context, *ExistRequest) (*ExistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Exist not implemented")
}
func (UnimplementedMasterDocGRPCServer) Create(context.Context, *CreateDocRequest) (*CreateDocResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedMasterDocGRPCServer) GetAll(context.Context, *Empty) (*GetAllDocsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedMasterDocGRPCServer) Delete(context.Context, *DeleteDocRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedMasterDocGRPCServer) UpdateLinks(context.Context, *UpdateLinksRequest) (*UpdateLinksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateLinks not implemented")
}
func (UnimplementedMasterDocGRPCServer) GetAbsents(context.Context, *Empty) (*GetAbsentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAbsents not implemented")
}
func (UnimplementedMasterDocGRPCServer) mustEmbedUnimplementedMasterDocGRPCServer() {}

// UnsafeMasterDocGRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MasterDocGRPCServer will
// result in compilation errors.
type UnsafeMasterDocGRPCServer interface {
	mustEmbedUnimplementedMasterDocGRPCServer()
}

func RegisterMasterDocGRPCServer(s grpc.ServiceRegistrar, srv MasterDocGRPCServer) {
	s.RegisterService(&MasterDocGRPC_ServiceDesc, srv)
}

func _MasterDocGRPC_Exist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterDocGRPCServer).Exist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/master.v1.MasterDocGRPC/Exist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterDocGRPCServer).Exist(ctx, req.(*ExistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterDocGRPC_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDocRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterDocGRPCServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/master.v1.MasterDocGRPC/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterDocGRPCServer).Create(ctx, req.(*CreateDocRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterDocGRPC_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterDocGRPCServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/master.v1.MasterDocGRPC/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterDocGRPCServer).GetAll(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterDocGRPC_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDocRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterDocGRPCServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/master.v1.MasterDocGRPC/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterDocGRPCServer).Delete(ctx, req.(*DeleteDocRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterDocGRPC_UpdateLinks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateLinksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterDocGRPCServer).UpdateLinks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/master.v1.MasterDocGRPC/UpdateLinks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterDocGRPCServer).UpdateLinks(ctx, req.(*UpdateLinksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterDocGRPC_GetAbsents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterDocGRPCServer).GetAbsents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/master.v1.MasterDocGRPC/GetAbsents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterDocGRPCServer).GetAbsents(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// MasterDocGRPC_ServiceDesc is the grpc.ServiceDesc for MasterDocGRPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MasterDocGRPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "master.v1.MasterDocGRPC",
	HandlerType: (*MasterDocGRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Exist",
			Handler:    _MasterDocGRPC_Exist_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _MasterDocGRPC_Create_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _MasterDocGRPC_GetAll_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _MasterDocGRPC_Delete_Handler,
		},
		{
			MethodName: "UpdateLinks",
			Handler:    _MasterDocGRPC_UpdateLinks_Handler,
		},
		{
			MethodName: "GetAbsents",
			Handler:    _MasterDocGRPC_GetAbsents_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "master_service.proto",
}

// MasterChapterGRPCClient is the client API for MasterChapterGRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MasterChapterGRPCClient interface {
	Create(ctx context.Context, in *CreateChapterRequest, opts ...grpc.CallOption) (*CreateChapterResponse, error)
	GetAll(ctx context.Context, in *GetAllChaptersRequest, opts ...grpc.CallOption) (*GetAllChaptersResponse, error)
}

type masterChapterGRPCClient struct {
	cc grpc.ClientConnInterface
}

func NewMasterChapterGRPCClient(cc grpc.ClientConnInterface) MasterChapterGRPCClient {
	return &masterChapterGRPCClient{cc}
}

func (c *masterChapterGRPCClient) Create(ctx context.Context, in *CreateChapterRequest, opts ...grpc.CallOption) (*CreateChapterResponse, error) {
	out := new(CreateChapterResponse)
	err := c.cc.Invoke(ctx, "/master.v1.MasterChapterGRPC/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterChapterGRPCClient) GetAll(ctx context.Context, in *GetAllChaptersRequest, opts ...grpc.CallOption) (*GetAllChaptersResponse, error) {
	out := new(GetAllChaptersResponse)
	err := c.cc.Invoke(ctx, "/master.v1.MasterChapterGRPC/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MasterChapterGRPCServer is the server API for MasterChapterGRPC service.
// All implementations must embed UnimplementedMasterChapterGRPCServer
// for forward compatibility
type MasterChapterGRPCServer interface {
	Create(context.Context, *CreateChapterRequest) (*CreateChapterResponse, error)
	GetAll(context.Context, *GetAllChaptersRequest) (*GetAllChaptersResponse, error)
	mustEmbedUnimplementedMasterChapterGRPCServer()
}

// UnimplementedMasterChapterGRPCServer must be embedded to have forward compatible implementations.
type UnimplementedMasterChapterGRPCServer struct {
}

func (UnimplementedMasterChapterGRPCServer) Create(context.Context, *CreateChapterRequest) (*CreateChapterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedMasterChapterGRPCServer) GetAll(context.Context, *GetAllChaptersRequest) (*GetAllChaptersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedMasterChapterGRPCServer) mustEmbedUnimplementedMasterChapterGRPCServer() {}

// UnsafeMasterChapterGRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MasterChapterGRPCServer will
// result in compilation errors.
type UnsafeMasterChapterGRPCServer interface {
	mustEmbedUnimplementedMasterChapterGRPCServer()
}

func RegisterMasterChapterGRPCServer(s grpc.ServiceRegistrar, srv MasterChapterGRPCServer) {
	s.RegisterService(&MasterChapterGRPC_ServiceDesc, srv)
}

func _MasterChapterGRPC_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateChapterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterChapterGRPCServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/master.v1.MasterChapterGRPC/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterChapterGRPCServer).Create(ctx, req.(*CreateChapterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterChapterGRPC_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllChaptersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterChapterGRPCServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/master.v1.MasterChapterGRPC/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterChapterGRPCServer).GetAll(ctx, req.(*GetAllChaptersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MasterChapterGRPC_ServiceDesc is the grpc.ServiceDesc for MasterChapterGRPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MasterChapterGRPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "master.v1.MasterChapterGRPC",
	HandlerType: (*MasterChapterGRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _MasterChapterGRPC_Create_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _MasterChapterGRPC_GetAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "master_service.proto",
}

// MasterParagraphGRPCClient is the client API for MasterParagraphGRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MasterParagraphGRPCClient interface {
	Create(ctx context.Context, in *CreateParagraphsRequest, opts ...grpc.CallOption) (*Empty, error)
	GetOne(ctx context.Context, in *GetOneParagraphRequest, opts ...grpc.CallOption) (*GetOneParagraphResponse, error)
	Update(ctx context.Context, in *UpdateParagraphRequest, opts ...grpc.CallOption) (*Empty, error)
}

type masterParagraphGRPCClient struct {
	cc grpc.ClientConnInterface
}

func NewMasterParagraphGRPCClient(cc grpc.ClientConnInterface) MasterParagraphGRPCClient {
	return &masterParagraphGRPCClient{cc}
}

func (c *masterParagraphGRPCClient) Create(ctx context.Context, in *CreateParagraphsRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/master.v1.MasterParagraphGRPC/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterParagraphGRPCClient) GetOne(ctx context.Context, in *GetOneParagraphRequest, opts ...grpc.CallOption) (*GetOneParagraphResponse, error) {
	out := new(GetOneParagraphResponse)
	err := c.cc.Invoke(ctx, "/master.v1.MasterParagraphGRPC/GetOne", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterParagraphGRPCClient) Update(ctx context.Context, in *UpdateParagraphRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/master.v1.MasterParagraphGRPC/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MasterParagraphGRPCServer is the server API for MasterParagraphGRPC service.
// All implementations must embed UnimplementedMasterParagraphGRPCServer
// for forward compatibility
type MasterParagraphGRPCServer interface {
	Create(context.Context, *CreateParagraphsRequest) (*Empty, error)
	GetOne(context.Context, *GetOneParagraphRequest) (*GetOneParagraphResponse, error)
	Update(context.Context, *UpdateParagraphRequest) (*Empty, error)
	mustEmbedUnimplementedMasterParagraphGRPCServer()
}

// UnimplementedMasterParagraphGRPCServer must be embedded to have forward compatible implementations.
type UnimplementedMasterParagraphGRPCServer struct {
}

func (UnimplementedMasterParagraphGRPCServer) Create(context.Context, *CreateParagraphsRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedMasterParagraphGRPCServer) GetOne(context.Context, *GetOneParagraphRequest) (*GetOneParagraphResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOne not implemented")
}
func (UnimplementedMasterParagraphGRPCServer) Update(context.Context, *UpdateParagraphRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedMasterParagraphGRPCServer) mustEmbedUnimplementedMasterParagraphGRPCServer() {}

// UnsafeMasterParagraphGRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MasterParagraphGRPCServer will
// result in compilation errors.
type UnsafeMasterParagraphGRPCServer interface {
	mustEmbedUnimplementedMasterParagraphGRPCServer()
}

func RegisterMasterParagraphGRPCServer(s grpc.ServiceRegistrar, srv MasterParagraphGRPCServer) {
	s.RegisterService(&MasterParagraphGRPC_ServiceDesc, srv)
}

func _MasterParagraphGRPC_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateParagraphsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterParagraphGRPCServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/master.v1.MasterParagraphGRPC/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterParagraphGRPCServer).Create(ctx, req.(*CreateParagraphsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterParagraphGRPC_GetOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOneParagraphRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterParagraphGRPCServer).GetOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/master.v1.MasterParagraphGRPC/GetOne",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterParagraphGRPCServer).GetOne(ctx, req.(*GetOneParagraphRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterParagraphGRPC_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateParagraphRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterParagraphGRPCServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/master.v1.MasterParagraphGRPC/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterParagraphGRPCServer).Update(ctx, req.(*UpdateParagraphRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MasterParagraphGRPC_ServiceDesc is the grpc.ServiceDesc for MasterParagraphGRPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MasterParagraphGRPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "master.v1.MasterParagraphGRPC",
	HandlerType: (*MasterParagraphGRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _MasterParagraphGRPC_Create_Handler,
		},
		{
			MethodName: "GetOne",
			Handler:    _MasterParagraphGRPC_GetOne_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _MasterParagraphGRPC_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "master_service.proto",
}
