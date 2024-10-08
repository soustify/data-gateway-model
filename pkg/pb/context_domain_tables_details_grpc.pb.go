// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: context_domain_tables_details.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ContextDomainTableDetailService_Create_FullMethodName  = "/domain.ContextDomainTableDetailService/Create"
	ContextDomainTableDetailService_Update_FullMethodName  = "/domain.ContextDomainTableDetailService/Update"
	ContextDomainTableDetailService_FindAll_FullMethodName = "/domain.ContextDomainTableDetailService/FindAll"
)

// ContextDomainTableDetailServiceClient is the client API for ContextDomainTableDetailService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ContextDomainTableDetailServiceClient interface {
	Create(ctx context.Context, in *ContextDomainTableDetail, opts ...grpc.CallOption) (*ContextDomainTableDetailResponse, error)
	Update(ctx context.Context, in *ContextDomainTableDetail, opts ...grpc.CallOption) (*ContextDomainTableDetailResponse, error)
	FindAll(ctx context.Context, in *ContextDomainTableDetailListRequest, opts ...grpc.CallOption) (*ContextDomainTableDetailListResponse, error)
}

type contextDomainTableDetailServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewContextDomainTableDetailServiceClient(cc grpc.ClientConnInterface) ContextDomainTableDetailServiceClient {
	return &contextDomainTableDetailServiceClient{cc}
}

func (c *contextDomainTableDetailServiceClient) Create(ctx context.Context, in *ContextDomainTableDetail, opts ...grpc.CallOption) (*ContextDomainTableDetailResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ContextDomainTableDetailResponse)
	err := c.cc.Invoke(ctx, ContextDomainTableDetailService_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contextDomainTableDetailServiceClient) Update(ctx context.Context, in *ContextDomainTableDetail, opts ...grpc.CallOption) (*ContextDomainTableDetailResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ContextDomainTableDetailResponse)
	err := c.cc.Invoke(ctx, ContextDomainTableDetailService_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contextDomainTableDetailServiceClient) FindAll(ctx context.Context, in *ContextDomainTableDetailListRequest, opts ...grpc.CallOption) (*ContextDomainTableDetailListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ContextDomainTableDetailListResponse)
	err := c.cc.Invoke(ctx, ContextDomainTableDetailService_FindAll_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContextDomainTableDetailServiceServer is the server API for ContextDomainTableDetailService service.
// All implementations must embed UnimplementedContextDomainTableDetailServiceServer
// for forward compatibility.
type ContextDomainTableDetailServiceServer interface {
	Create(context.Context, *ContextDomainTableDetail) (*ContextDomainTableDetailResponse, error)
	Update(context.Context, *ContextDomainTableDetail) (*ContextDomainTableDetailResponse, error)
	FindAll(context.Context, *ContextDomainTableDetailListRequest) (*ContextDomainTableDetailListResponse, error)
	mustEmbedUnimplementedContextDomainTableDetailServiceServer()
}

// UnimplementedContextDomainTableDetailServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedContextDomainTableDetailServiceServer struct{}

func (UnimplementedContextDomainTableDetailServiceServer) Create(context.Context, *ContextDomainTableDetail) (*ContextDomainTableDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedContextDomainTableDetailServiceServer) Update(context.Context, *ContextDomainTableDetail) (*ContextDomainTableDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedContextDomainTableDetailServiceServer) FindAll(context.Context, *ContextDomainTableDetailListRequest) (*ContextDomainTableDetailListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAll not implemented")
}
func (UnimplementedContextDomainTableDetailServiceServer) mustEmbedUnimplementedContextDomainTableDetailServiceServer() {
}
func (UnimplementedContextDomainTableDetailServiceServer) testEmbeddedByValue() {}

// UnsafeContextDomainTableDetailServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ContextDomainTableDetailServiceServer will
// result in compilation errors.
type UnsafeContextDomainTableDetailServiceServer interface {
	mustEmbedUnimplementedContextDomainTableDetailServiceServer()
}

func RegisterContextDomainTableDetailServiceServer(s grpc.ServiceRegistrar, srv ContextDomainTableDetailServiceServer) {
	// If the following call pancis, it indicates UnimplementedContextDomainTableDetailServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ContextDomainTableDetailService_ServiceDesc, srv)
}

func _ContextDomainTableDetailService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContextDomainTableDetail)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContextDomainTableDetailServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContextDomainTableDetailService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContextDomainTableDetailServiceServer).Create(ctx, req.(*ContextDomainTableDetail))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContextDomainTableDetailService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContextDomainTableDetail)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContextDomainTableDetailServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContextDomainTableDetailService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContextDomainTableDetailServiceServer).Update(ctx, req.(*ContextDomainTableDetail))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContextDomainTableDetailService_FindAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContextDomainTableDetailListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContextDomainTableDetailServiceServer).FindAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContextDomainTableDetailService_FindAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContextDomainTableDetailServiceServer).FindAll(ctx, req.(*ContextDomainTableDetailListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ContextDomainTableDetailService_ServiceDesc is the grpc.ServiceDesc for ContextDomainTableDetailService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ContextDomainTableDetailService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "domain.ContextDomainTableDetailService",
	HandlerType: (*ContextDomainTableDetailServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ContextDomainTableDetailService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ContextDomainTableDetailService_Update_Handler,
		},
		{
			MethodName: "FindAll",
			Handler:    _ContextDomainTableDetailService_FindAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "context_domain_tables_details.proto",
}
