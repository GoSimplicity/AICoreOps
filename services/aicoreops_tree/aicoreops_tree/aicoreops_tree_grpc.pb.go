// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: aicoreops_tree.proto

package aicoreops_tree

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	AicoreopsTree_Ping_FullMethodName = "/aicoreops_tree.Aicoreops_tree/Ping"
)

// AicoreopsTreeClient is the client API for AicoreopsTree service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AicoreopsTreeClient interface {
	Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type aicoreopsTreeClient struct {
	cc grpc.ClientConnInterface
}

func NewAicoreopsTreeClient(cc grpc.ClientConnInterface) AicoreopsTreeClient {
	return &aicoreopsTreeClient{cc}
}

func (c *aicoreopsTreeClient) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, AicoreopsTree_Ping_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AicoreopsTreeServer is the server API for AicoreopsTree service.
// All implementations must embed UnimplementedAicoreopsTreeServer
// for forward compatibility
type AicoreopsTreeServer interface {
	Ping(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedAicoreopsTreeServer()
}

// UnimplementedAicoreopsTreeServer must be embedded to have forward compatible implementations.
type UnimplementedAicoreopsTreeServer struct {
}

func (UnimplementedAicoreopsTreeServer) Ping(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedAicoreopsTreeServer) mustEmbedUnimplementedAicoreopsTreeServer() {}

// UnsafeAicoreopsTreeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AicoreopsTreeServer will
// result in compilation errors.
type UnsafeAicoreopsTreeServer interface {
	mustEmbedUnimplementedAicoreopsTreeServer()
}

func RegisterAicoreopsTreeServer(s grpc.ServiceRegistrar, srv AicoreopsTreeServer) {
	s.RegisterService(&AicoreopsTree_ServiceDesc, srv)
}

func _AicoreopsTree_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AicoreopsTreeServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AicoreopsTree_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AicoreopsTreeServer).Ping(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// AicoreopsTree_ServiceDesc is the grpc.ServiceDesc for AicoreopsTree service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AicoreopsTree_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "aicoreops_tree.Aicoreops_tree",
	HandlerType: (*AicoreopsTreeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _AicoreopsTree_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "aicoreops_tree.proto",
}
