// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.1
// source: prometheus_rpc.proto

package types

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
	PrometheusRpc_GetMonitorAlertmanagerPoolList_FullMethodName = "/prometheus_rpc.Prometheus_rpc/GetMonitorAlertmanagerPoolList"
	PrometheusRpc_CreateMonitorAlertManagerPool_FullMethodName  = "/prometheus_rpc.Prometheus_rpc/CreateMonitorAlertManagerPool"
	PrometheusRpc_UpdateMonitorAlertManagerPool_FullMethodName  = "/prometheus_rpc.Prometheus_rpc/UpdateMonitorAlertManagerPool"
	PrometheusRpc_DeleteMonitorAlertManagerPool_FullMethodName  = "/prometheus_rpc.Prometheus_rpc/DeleteMonitorAlertManagerPool"
)

// PrometheusRpcClient is the client API for PrometheusRpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PrometheusRpcClient interface {
	// Alertmanager 相关接口
	GetMonitorAlertmanagerPoolList(ctx context.Context, in *GetAlertmanagerPoolListRequest, opts ...grpc.CallOption) (*GetAlertmanagerPoolListResponse, error)
	CreateMonitorAlertManagerPool(ctx context.Context, in *CreateMonitorAlertManagerPoolRequest, opts ...grpc.CallOption) (*CreateMonitorAlertManagerPoolResponse, error)
	UpdateMonitorAlertManagerPool(ctx context.Context, in *UpdateMonitorAlertManagerPoolRequest, opts ...grpc.CallOption) (*UpdateMonitorAlertManagerPoolResponse, error)
	DeleteMonitorAlertManagerPool(ctx context.Context, in *DeleteMonitorAlertManagerPoolRequest, opts ...grpc.CallOption) (*DeleteMonitorAlertManagerPoolResponse, error)
}

type prometheusRpcClient struct {
	cc grpc.ClientConnInterface
}

func NewPrometheusRpcClient(cc grpc.ClientConnInterface) PrometheusRpcClient {
	return &prometheusRpcClient{cc}
}

func (c *prometheusRpcClient) GetMonitorAlertmanagerPoolList(ctx context.Context, in *GetAlertmanagerPoolListRequest, opts ...grpc.CallOption) (*GetAlertmanagerPoolListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAlertmanagerPoolListResponse)
	err := c.cc.Invoke(ctx, PrometheusRpc_GetMonitorAlertmanagerPoolList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *prometheusRpcClient) CreateMonitorAlertManagerPool(ctx context.Context, in *CreateMonitorAlertManagerPoolRequest, opts ...grpc.CallOption) (*CreateMonitorAlertManagerPoolResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateMonitorAlertManagerPoolResponse)
	err := c.cc.Invoke(ctx, PrometheusRpc_CreateMonitorAlertManagerPool_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *prometheusRpcClient) UpdateMonitorAlertManagerPool(ctx context.Context, in *UpdateMonitorAlertManagerPoolRequest, opts ...grpc.CallOption) (*UpdateMonitorAlertManagerPoolResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateMonitorAlertManagerPoolResponse)
	err := c.cc.Invoke(ctx, PrometheusRpc_UpdateMonitorAlertManagerPool_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *prometheusRpcClient) DeleteMonitorAlertManagerPool(ctx context.Context, in *DeleteMonitorAlertManagerPoolRequest, opts ...grpc.CallOption) (*DeleteMonitorAlertManagerPoolResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteMonitorAlertManagerPoolResponse)
	err := c.cc.Invoke(ctx, PrometheusRpc_DeleteMonitorAlertManagerPool_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PrometheusRpcServer is the server API for PrometheusRpc service.
// All implementations must embed UnimplementedPrometheusRpcServer
// for forward compatibility.
type PrometheusRpcServer interface {
	// Alertmanager 相关接口
	GetMonitorAlertmanagerPoolList(context.Context, *GetAlertmanagerPoolListRequest) (*GetAlertmanagerPoolListResponse, error)
	CreateMonitorAlertManagerPool(context.Context, *CreateMonitorAlertManagerPoolRequest) (*CreateMonitorAlertManagerPoolResponse, error)
	UpdateMonitorAlertManagerPool(context.Context, *UpdateMonitorAlertManagerPoolRequest) (*UpdateMonitorAlertManagerPoolResponse, error)
	DeleteMonitorAlertManagerPool(context.Context, *DeleteMonitorAlertManagerPoolRequest) (*DeleteMonitorAlertManagerPoolResponse, error)
	mustEmbedUnimplementedPrometheusRpcServer()
}

// UnimplementedPrometheusRpcServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPrometheusRpcServer struct{}

func (UnimplementedPrometheusRpcServer) GetMonitorAlertmanagerPoolList(context.Context, *GetAlertmanagerPoolListRequest) (*GetAlertmanagerPoolListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMonitorAlertmanagerPoolList not implemented")
}
func (UnimplementedPrometheusRpcServer) CreateMonitorAlertManagerPool(context.Context, *CreateMonitorAlertManagerPoolRequest) (*CreateMonitorAlertManagerPoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMonitorAlertManagerPool not implemented")
}
func (UnimplementedPrometheusRpcServer) UpdateMonitorAlertManagerPool(context.Context, *UpdateMonitorAlertManagerPoolRequest) (*UpdateMonitorAlertManagerPoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMonitorAlertManagerPool not implemented")
}
func (UnimplementedPrometheusRpcServer) DeleteMonitorAlertManagerPool(context.Context, *DeleteMonitorAlertManagerPoolRequest) (*DeleteMonitorAlertManagerPoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMonitorAlertManagerPool not implemented")
}
func (UnimplementedPrometheusRpcServer) mustEmbedUnimplementedPrometheusRpcServer() {}
func (UnimplementedPrometheusRpcServer) testEmbeddedByValue()                       {}

// UnsafePrometheusRpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PrometheusRpcServer will
// result in compilation errors.
type UnsafePrometheusRpcServer interface {
	mustEmbedUnimplementedPrometheusRpcServer()
}

func RegisterPrometheusRpcServer(s grpc.ServiceRegistrar, srv PrometheusRpcServer) {
	// If the following call pancis, it indicates UnimplementedPrometheusRpcServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PrometheusRpc_ServiceDesc, srv)
}

func _PrometheusRpc_GetMonitorAlertmanagerPoolList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAlertmanagerPoolListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrometheusRpcServer).GetMonitorAlertmanagerPoolList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PrometheusRpc_GetMonitorAlertmanagerPoolList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrometheusRpcServer).GetMonitorAlertmanagerPoolList(ctx, req.(*GetAlertmanagerPoolListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrometheusRpc_CreateMonitorAlertManagerPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMonitorAlertManagerPoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrometheusRpcServer).CreateMonitorAlertManagerPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PrometheusRpc_CreateMonitorAlertManagerPool_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrometheusRpcServer).CreateMonitorAlertManagerPool(ctx, req.(*CreateMonitorAlertManagerPoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrometheusRpc_UpdateMonitorAlertManagerPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMonitorAlertManagerPoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrometheusRpcServer).UpdateMonitorAlertManagerPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PrometheusRpc_UpdateMonitorAlertManagerPool_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrometheusRpcServer).UpdateMonitorAlertManagerPool(ctx, req.(*UpdateMonitorAlertManagerPoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrometheusRpc_DeleteMonitorAlertManagerPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMonitorAlertManagerPoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrometheusRpcServer).DeleteMonitorAlertManagerPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PrometheusRpc_DeleteMonitorAlertManagerPool_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrometheusRpcServer).DeleteMonitorAlertManagerPool(ctx, req.(*DeleteMonitorAlertManagerPoolRequest))
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
			MethodName: "GetMonitorAlertmanagerPoolList",
			Handler:    _PrometheusRpc_GetMonitorAlertmanagerPoolList_Handler,
		},
		{
			MethodName: "CreateMonitorAlertManagerPool",
			Handler:    _PrometheusRpc_CreateMonitorAlertManagerPool_Handler,
		},
		{
			MethodName: "UpdateMonitorAlertManagerPool",
			Handler:    _PrometheusRpc_UpdateMonitorAlertManagerPool_Handler,
		},
		{
			MethodName: "DeleteMonitorAlertManagerPool",
			Handler:    _PrometheusRpc_DeleteMonitorAlertManagerPool_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "prometheus_rpc.proto",
}
