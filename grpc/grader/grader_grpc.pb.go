// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package grader

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

// GraderServiceClient is the client API for GraderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GraderServiceClient interface {
	SayHello(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
}

type graderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGraderServiceClient(cc grpc.ClientConnInterface) GraderServiceClient {
	return &graderServiceClient{cc}
}

func (c *graderServiceClient) SayHello(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/grader.GraderService/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GraderServiceServer is the server API for GraderService service.
// All implementations must embed UnimplementedGraderServiceServer
// for forward compatibility
type GraderServiceServer interface {
	SayHello(context.Context, *Message) (*Message, error)
	mustEmbedUnimplementedGraderServiceServer()
}

// UnimplementedGraderServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGraderServiceServer struct {
}

func (UnimplementedGraderServiceServer) SayHello(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedGraderServiceServer) mustEmbedUnimplementedGraderServiceServer() {}

// UnsafeGraderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GraderServiceServer will
// result in compilation errors.
type UnsafeGraderServiceServer interface {
	mustEmbedUnimplementedGraderServiceServer()
}

func RegisterGraderServiceServer(s grpc.ServiceRegistrar, srv GraderServiceServer) {
	s.RegisterService(&GraderService_ServiceDesc, srv)
}

func _GraderService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GraderServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grader.GraderService/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GraderServiceServer).SayHello(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

// GraderService_ServiceDesc is the grpc.ServiceDesc for GraderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GraderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grader.GraderService",
	HandlerType: (*GraderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _GraderService_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grader.proto",
}
