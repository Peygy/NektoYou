// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.0--rc3
// source: protos/graph_auth/graph_auth.proto

package graph_auth

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
	SignInService_GeneratePairOfTokens_FullMethodName = "/graph_auth.SignInService/GeneratePairOfTokens"
)

// SignInServiceClient is the client API for SignInService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SignInServiceClient interface {
	GeneratePairOfTokens(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*SignInResponce, error)
}

type signInServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSignInServiceClient(cc grpc.ClientConnInterface) SignInServiceClient {
	return &signInServiceClient{cc}
}

func (c *signInServiceClient) GeneratePairOfTokens(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*SignInResponce, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SignInResponce)
	err := c.cc.Invoke(ctx, SignInService_GeneratePairOfTokens_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SignInServiceServer is the server API for SignInService service.
// All implementations must embed UnimplementedSignInServiceServer
// for forward compatibility.
type SignInServiceServer interface {
	GeneratePairOfTokens(context.Context, *SignInRequest) (*SignInResponce, error)
	mustEmbedUnimplementedSignInServiceServer()
}

// UnimplementedSignInServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSignInServiceServer struct{}

func (UnimplementedSignInServiceServer) GeneratePairOfTokens(context.Context, *SignInRequest) (*SignInResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GeneratePairOfTokens not implemented")
}
func (UnimplementedSignInServiceServer) mustEmbedUnimplementedSignInServiceServer() {}
func (UnimplementedSignInServiceServer) testEmbeddedByValue()                       {}

// UnsafeSignInServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SignInServiceServer will
// result in compilation errors.
type UnsafeSignInServiceServer interface {
	mustEmbedUnimplementedSignInServiceServer()
}

func RegisterSignInServiceServer(s grpc.ServiceRegistrar, srv SignInServiceServer) {
	// If the following call pancis, it indicates UnimplementedSignInServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SignInService_ServiceDesc, srv)
}

func _SignInService_GeneratePairOfTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignInServiceServer).GeneratePairOfTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SignInService_GeneratePairOfTokens_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignInServiceServer).GeneratePairOfTokens(ctx, req.(*SignInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SignInService_ServiceDesc is the grpc.ServiceDesc for SignInService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SignInService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "graph_auth.SignInService",
	HandlerType: (*SignInServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GeneratePairOfTokens",
			Handler:    _SignInService_GeneratePairOfTokens_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/graph_auth/graph_auth.proto",
}
