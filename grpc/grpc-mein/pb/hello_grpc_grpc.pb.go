// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: hello_grpc.proto

package hello_grpc

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

// HelloServiceClient is the client API for HelloService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HelloServiceClient interface {
	SayHi(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Res, error)
}

type helloServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHelloServiceClient(cc grpc.ClientConnInterface) HelloServiceClient {
	return &helloServiceClient{cc}
}

func (c *helloServiceClient) SayHi(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Res, error) {
	out := new(Res)
	err := c.cc.Invoke(ctx, "/hello_grpc.HelloService/SayHi", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloServiceServer is the server API for HelloService service.
// All implementations must embed UnimplementedHelloServiceServer
// for forward compatibility
type HelloServiceServer interface {
	SayHi(context.Context, *Req) (*Res, error)
	mustEmbedUnimplementedHelloServiceServer()
}

// UnimplementedHelloServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHelloServiceServer struct {
}

func (UnimplementedHelloServiceServer) SayHi(context.Context, *Req) (*Res, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHi not implemented")
}
func (UnimplementedHelloServiceServer) mustEmbedUnimplementedHelloServiceServer() {}

// UnsafeHelloServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HelloServiceServer will
// result in compilation errors.
type UnsafeHelloServiceServer interface {
	mustEmbedUnimplementedHelloServiceServer()
}

func RegisterHelloServiceServer(s grpc.ServiceRegistrar, srv HelloServiceServer) {
	s.RegisterService(&HelloService_ServiceDesc, srv)
}

func _HelloService_SayHi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServiceServer).SayHi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hello_grpc.HelloService/SayHi",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServiceServer).SayHi(ctx, req.(*Req))
	}
	return interceptor(ctx, in, info, handler)
}

// HelloService_ServiceDesc is the grpc.ServiceDesc for HelloService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HelloService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hello_grpc.HelloService",
	HandlerType: (*HelloServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHi",
			Handler:    _HelloService_SayHi_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hello_grpc.proto",
}
