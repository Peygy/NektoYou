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
	SignUpService_SignUp_FullMethodName = "/graph_auth.SignUpService/SignUp"
)

// SignUpServiceClient is the client API for SignUpService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SignUpServiceClient interface {
	SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*SignUpResponce, error)
}

type signUpServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSignUpServiceClient(cc grpc.ClientConnInterface) SignUpServiceClient {
	return &signUpServiceClient{cc}
}

func (c *signUpServiceClient) SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*SignUpResponce, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SignUpResponce)
	err := c.cc.Invoke(ctx, SignUpService_SignUp_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SignUpServiceServer is the server API for SignUpService service.
// All implementations must embed UnimplementedSignUpServiceServer
// for forward compatibility.
type SignUpServiceServer interface {
	SignUp(context.Context, *SignUpRequest) (*SignUpResponce, error)
	mustEmbedUnimplementedSignUpServiceServer()
}

// UnimplementedSignUpServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSignUpServiceServer struct{}

func (UnimplementedSignUpServiceServer) SignUp(context.Context, *SignUpRequest) (*SignUpResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignUp not implemented")
}
func (UnimplementedSignUpServiceServer) mustEmbedUnimplementedSignUpServiceServer() {}
func (UnimplementedSignUpServiceServer) testEmbeddedByValue()                       {}

// UnsafeSignUpServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SignUpServiceServer will
// result in compilation errors.
type UnsafeSignUpServiceServer interface {
	mustEmbedUnimplementedSignUpServiceServer()
}

func RegisterSignUpServiceServer(s grpc.ServiceRegistrar, srv SignUpServiceServer) {
	// If the following call pancis, it indicates UnimplementedSignUpServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SignUpService_ServiceDesc, srv)
}

func _SignUpService_SignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignUpServiceServer).SignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SignUpService_SignUp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignUpServiceServer).SignUp(ctx, req.(*SignUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SignUpService_ServiceDesc is the grpc.ServiceDesc for SignUpService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SignUpService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "graph_auth.SignUpService",
	HandlerType: (*SignUpServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignUp",
			Handler:    _SignUpService_SignUp_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/graph_auth/graph_auth.proto",
}
