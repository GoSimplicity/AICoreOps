// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: prometheus_rpc.proto

package prometheus_rpc

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
	PrometheusRpc_Ping_FullMethodName = "/prometheus_rpc.Prometheus_rpc/Ping"
)

// PrometheusRpcClient is the client API for PrometheusRpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PrometheusRpcClient interface {
	Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type prometheusRpcClient struct {
	cc grpc.ClientConnInterface
}

func NewPrometheusRpcClient(cc grpc.ClientConnInterface) PrometheusRpcClient {
	return &prometheusRpcClient{cc}
}

func (c *prometheusRpcClient) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, PrometheusRpc_Ping_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PrometheusRpcServer is the server API for PrometheusRpc service.
// All implementations must embed UnimplementedPrometheusRpcServer
// for forward compatibility
type PrometheusRpcServer interface {
	Ping(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedPrometheusRpcServer()
}

// UnimplementedPrometheusRpcServer must be embedded to have forward compatible implementations.
type UnimplementedPrometheusRpcServer struct {
}

func (UnimplementedPrometheusRpcServer) Ping(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedPrometheusRpcServer) mustEmbedUnimplementedPrometheusRpcServer() {}

// UnsafePrometheusRpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PrometheusRpcServer will
// result in compilation errors.
type UnsafePrometheusRpcServer interface {
	mustEmbedUnimplementedPrometheusRpcServer()
}

func RegisterPrometheusRpcServer(s grpc.ServiceRegistrar, srv PrometheusRpcServer) {
	s.RegisterService(&PrometheusRpc_ServiceDesc, srv)
}

func _PrometheusRpc_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrometheusRpcServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PrometheusRpc_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrometheusRpcServer).Ping(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// PrometheusRpc_ServiceDesc is the grpc.ServiceDesc for PrometheusRpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PrometheusRpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "prometheus_rpc.Prometheus_rpc",
	HandlerType: (*PrometheusRpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _PrometheusRpc_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "prometheus_rpc.proto",
}
