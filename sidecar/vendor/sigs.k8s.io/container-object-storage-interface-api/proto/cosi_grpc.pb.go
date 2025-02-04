// Code generated by make; DO NOT EDIT.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.27.2
// source: cosi.proto

package cosi

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

const (
	Identity_DriverGetInfo_FullMethodName = "/cosi.v1alpha1.Identity/DriverGetInfo"
)

// IdentityClient is the client API for Identity service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IdentityClient interface {
	// This call is meant to retrieve the unique provisioner Identity.
	// This identity will have to be set in BucketClaim.DriverName field in order to invoke this specific provisioner.
	DriverGetInfo(ctx context.Context, in *DriverGetInfoRequest, opts ...grpc.CallOption) (*DriverGetInfoResponse, error)
}

type identityClient struct {
	cc grpc.ClientConnInterface
}

func NewIdentityClient(cc grpc.ClientConnInterface) IdentityClient {
	return &identityClient{cc}
}

func (c *identityClient) DriverGetInfo(ctx context.Context, in *DriverGetInfoRequest, opts ...grpc.CallOption) (*DriverGetInfoResponse, error) {
	out := new(DriverGetInfoResponse)
	err := c.cc.Invoke(ctx, Identity_DriverGetInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IdentityServer is the server API for Identity service.
// All implementations must embed UnimplementedIdentityServer
// for forward compatibility
type IdentityServer interface {
	// This call is meant to retrieve the unique provisioner Identity.
	// This identity will have to be set in BucketClaim.DriverName field in order to invoke this specific provisioner.
	DriverGetInfo(context.Context, *DriverGetInfoRequest) (*DriverGetInfoResponse, error)
	mustEmbedUnimplementedIdentityServer()
}

// UnimplementedIdentityServer must be embedded to have forward compatible implementations.
type UnimplementedIdentityServer struct {
}

func (UnimplementedIdentityServer) DriverGetInfo(context.Context, *DriverGetInfoRequest) (*DriverGetInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DriverGetInfo not implemented")
}
func (UnimplementedIdentityServer) mustEmbedUnimplementedIdentityServer() {}

// UnsafeIdentityServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IdentityServer will
// result in compilation errors.
type UnsafeIdentityServer interface {
	mustEmbedUnimplementedIdentityServer()
}

func RegisterIdentityServer(s grpc.ServiceRegistrar, srv IdentityServer) {
	s.RegisterService(&Identity_ServiceDesc, srv)
}

func _Identity_DriverGetInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DriverGetInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServer).DriverGetInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Identity_DriverGetInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServer).DriverGetInfo(ctx, req.(*DriverGetInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Identity_ServiceDesc is the grpc.ServiceDesc for Identity service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Identity_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cosi.v1alpha1.Identity",
	HandlerType: (*IdentityServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DriverGetInfo",
			Handler:    _Identity_DriverGetInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosi.proto",
}

const (
	Provisioner_DriverCreateBucket_FullMethodName       = "/cosi.v1alpha1.Provisioner/DriverCreateBucket"
	Provisioner_DriverDeleteBucket_FullMethodName       = "/cosi.v1alpha1.Provisioner/DriverDeleteBucket"
	Provisioner_DriverGrantBucketAccess_FullMethodName  = "/cosi.v1alpha1.Provisioner/DriverGrantBucketAccess"
	Provisioner_DriverRevokeBucketAccess_FullMethodName = "/cosi.v1alpha1.Provisioner/DriverRevokeBucketAccess"
)

// ProvisionerClient is the client API for Provisioner service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProvisionerClient interface {
	// This call is made to create the bucket in the backend.
	// This call is idempotent
	//  1. If a bucket that matches both name and parameters already exists, then OK (success) must be returned.
	//  2. If a bucket by same name, but different parameters is provided, then the appropriate error code ALREADY_EXISTS must be returned.
	DriverCreateBucket(ctx context.Context, in *DriverCreateBucketRequest, opts ...grpc.CallOption) (*DriverCreateBucketResponse, error)
	// This call is made to delete the bucket in the backend.
	// If the bucket has already been deleted, then no error should be returned.
	DriverDeleteBucket(ctx context.Context, in *DriverDeleteBucketRequest, opts ...grpc.CallOption) (*DriverDeleteBucketResponse, error)
	// This call grants access to an account. The account_name in the request shall be used as a unique identifier to create credentials.
	// The account_id returned in the response will be used as the unique identifier for deleting this access when calling DriverRevokeBucketAccess.
	DriverGrantBucketAccess(ctx context.Context, in *DriverGrantBucketAccessRequest, opts ...grpc.CallOption) (*DriverGrantBucketAccessResponse, error)
	// This call revokes all access to a particular bucket from a principal.
	DriverRevokeBucketAccess(ctx context.Context, in *DriverRevokeBucketAccessRequest, opts ...grpc.CallOption) (*DriverRevokeBucketAccessResponse, error)
}

type provisionerClient struct {
	cc grpc.ClientConnInterface
}

func NewProvisionerClient(cc grpc.ClientConnInterface) ProvisionerClient {
	return &provisionerClient{cc}
}

func (c *provisionerClient) DriverCreateBucket(ctx context.Context, in *DriverCreateBucketRequest, opts ...grpc.CallOption) (*DriverCreateBucketResponse, error) {
	out := new(DriverCreateBucketResponse)
	err := c.cc.Invoke(ctx, Provisioner_DriverCreateBucket_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *provisionerClient) DriverDeleteBucket(ctx context.Context, in *DriverDeleteBucketRequest, opts ...grpc.CallOption) (*DriverDeleteBucketResponse, error) {
	out := new(DriverDeleteBucketResponse)
	err := c.cc.Invoke(ctx, Provisioner_DriverDeleteBucket_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *provisionerClient) DriverGrantBucketAccess(ctx context.Context, in *DriverGrantBucketAccessRequest, opts ...grpc.CallOption) (*DriverGrantBucketAccessResponse, error) {
	out := new(DriverGrantBucketAccessResponse)
	err := c.cc.Invoke(ctx, Provisioner_DriverGrantBucketAccess_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *provisionerClient) DriverRevokeBucketAccess(ctx context.Context, in *DriverRevokeBucketAccessRequest, opts ...grpc.CallOption) (*DriverRevokeBucketAccessResponse, error) {
	out := new(DriverRevokeBucketAccessResponse)
	err := c.cc.Invoke(ctx, Provisioner_DriverRevokeBucketAccess_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProvisionerServer is the server API for Provisioner service.
// All implementations must embed UnimplementedProvisionerServer
// for forward compatibility
type ProvisionerServer interface {
	// This call is made to create the bucket in the backend.
	// This call is idempotent
	//  1. If a bucket that matches both name and parameters already exists, then OK (success) must be returned.
	//  2. If a bucket by same name, but different parameters is provided, then the appropriate error code ALREADY_EXISTS must be returned.
	DriverCreateBucket(context.Context, *DriverCreateBucketRequest) (*DriverCreateBucketResponse, error)
	// This call is made to delete the bucket in the backend.
	// If the bucket has already been deleted, then no error should be returned.
	DriverDeleteBucket(context.Context, *DriverDeleteBucketRequest) (*DriverDeleteBucketResponse, error)
	// This call grants access to an account. The account_name in the request shall be used as a unique identifier to create credentials.
	// The account_id returned in the response will be used as the unique identifier for deleting this access when calling DriverRevokeBucketAccess.
	DriverGrantBucketAccess(context.Context, *DriverGrantBucketAccessRequest) (*DriverGrantBucketAccessResponse, error)
	// This call revokes all access to a particular bucket from a principal.
	DriverRevokeBucketAccess(context.Context, *DriverRevokeBucketAccessRequest) (*DriverRevokeBucketAccessResponse, error)
	mustEmbedUnimplementedProvisionerServer()
}

// UnimplementedProvisionerServer must be embedded to have forward compatible implementations.
type UnimplementedProvisionerServer struct {
}

func (UnimplementedProvisionerServer) DriverCreateBucket(context.Context, *DriverCreateBucketRequest) (*DriverCreateBucketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DriverCreateBucket not implemented")
}
func (UnimplementedProvisionerServer) DriverDeleteBucket(context.Context, *DriverDeleteBucketRequest) (*DriverDeleteBucketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DriverDeleteBucket not implemented")
}
func (UnimplementedProvisionerServer) DriverGrantBucketAccess(context.Context, *DriverGrantBucketAccessRequest) (*DriverGrantBucketAccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DriverGrantBucketAccess not implemented")
}
func (UnimplementedProvisionerServer) DriverRevokeBucketAccess(context.Context, *DriverRevokeBucketAccessRequest) (*DriverRevokeBucketAccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DriverRevokeBucketAccess not implemented")
}
func (UnimplementedProvisionerServer) mustEmbedUnimplementedProvisionerServer() {}

// UnsafeProvisionerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProvisionerServer will
// result in compilation errors.
type UnsafeProvisionerServer interface {
	mustEmbedUnimplementedProvisionerServer()
}

func RegisterProvisionerServer(s grpc.ServiceRegistrar, srv ProvisionerServer) {
	s.RegisterService(&Provisioner_ServiceDesc, srv)
}

func _Provisioner_DriverCreateBucket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DriverCreateBucketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProvisionerServer).DriverCreateBucket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Provisioner_DriverCreateBucket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProvisionerServer).DriverCreateBucket(ctx, req.(*DriverCreateBucketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Provisioner_DriverDeleteBucket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DriverDeleteBucketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProvisionerServer).DriverDeleteBucket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Provisioner_DriverDeleteBucket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProvisionerServer).DriverDeleteBucket(ctx, req.(*DriverDeleteBucketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Provisioner_DriverGrantBucketAccess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DriverGrantBucketAccessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProvisionerServer).DriverGrantBucketAccess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Provisioner_DriverGrantBucketAccess_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProvisionerServer).DriverGrantBucketAccess(ctx, req.(*DriverGrantBucketAccessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Provisioner_DriverRevokeBucketAccess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DriverRevokeBucketAccessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProvisionerServer).DriverRevokeBucketAccess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Provisioner_DriverRevokeBucketAccess_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProvisionerServer).DriverRevokeBucketAccess(ctx, req.(*DriverRevokeBucketAccessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Provisioner_ServiceDesc is the grpc.ServiceDesc for Provisioner service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Provisioner_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cosi.v1alpha1.Provisioner",
	HandlerType: (*ProvisionerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DriverCreateBucket",
			Handler:    _Provisioner_DriverCreateBucket_Handler,
		},
		{
			MethodName: "DriverDeleteBucket",
			Handler:    _Provisioner_DriverDeleteBucket_Handler,
		},
		{
			MethodName: "DriverGrantBucketAccess",
			Handler:    _Provisioner_DriverGrantBucketAccess_Handler,
		},
		{
			MethodName: "DriverRevokeBucketAccess",
			Handler:    _Provisioner_DriverRevokeBucketAccess_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosi.proto",
}
