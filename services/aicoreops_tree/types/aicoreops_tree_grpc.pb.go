// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: aicoreops_tree.proto

package types

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
	ResourceTreeService_ListTreeNode_FullMethodName        = "/tree_rpc.ResourceTreeService/ListTreeNode"
	ResourceTreeService_SelectTreeNode_FullMethodName      = "/tree_rpc.ResourceTreeService/SelectTreeNode"
	ResourceTreeService_GetTopTreeNode_FullMethodName      = "/tree_rpc.ResourceTreeService/GetTopTreeNode"
	ResourceTreeService_ListLeafTreeNode_FullMethodName    = "/tree_rpc.ResourceTreeService/ListLeafTreeNode"
	ResourceTreeService_CreateTreeNode_FullMethodName      = "/tree_rpc.ResourceTreeService/CreateTreeNode"
	ResourceTreeService_DeleteTreeNode_FullMethodName      = "/tree_rpc.ResourceTreeService/DeleteTreeNode"
	ResourceTreeService_GetChildrenTreeNode_FullMethodName = "/tree_rpc.ResourceTreeService/GetChildrenTreeNode"
	ResourceTreeService_UpdateTreeNode_FullMethodName      = "/tree_rpc.ResourceTreeService/UpdateTreeNode"
	ResourceTreeService_SyncCMDB_FullMethodName            = "/tree_rpc.ResourceTreeService/SyncCMDB"
)

// ResourceTreeServiceClient is the client API for ResourceTreeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// 资源树服务
type ResourceTreeServiceClient interface {
	// 获取树节点列表
	ListTreeNode(ctx context.Context, in *ListTreeNodeRequest, opts ...grpc.CallOption) (*ListTreeNodeResponse, error)
	// 选择树节点
	SelectTreeNode(ctx context.Context, in *SelectTreeNodeRequest, opts ...grpc.CallOption) (*SelectTreeNodeResponse, error)
	// 获取顶级树节点
	GetTopTreeNode(ctx context.Context, in *GetTopTreeNodeRequest, opts ...grpc.CallOption) (*GetTopTreeNodeResponse, error)
	// 获取叶子节点列表
	ListLeafTreeNode(ctx context.Context, in *ListLeafTreeNodeRequest, opts ...grpc.CallOption) (*ListLeafTreeNodeResponse, error)
	// 创建树节点
	CreateTreeNode(ctx context.Context, in *CreateTreeNodeRequest, opts ...grpc.CallOption) (*CreateTreeNodeResponse, error)
	// 删除树节点
	DeleteTreeNode(ctx context.Context, in *DeleteTreeNodeRequest, opts ...grpc.CallOption) (*DeleteTreeNodeResponse, error)
	// 获取子节点
	GetChildrenTreeNode(ctx context.Context, in *GetChildrenTreeNodeRequest, opts ...grpc.CallOption) (*GetChildrenTreeNodeResponse, error)
	// 更新树节点
	UpdateTreeNode(ctx context.Context, in *UpdateTreeNodeRequest, opts ...grpc.CallOption) (*UpdateTreeNodeResponse, error)
	// 同步CMDB资源
	SyncCMDB(ctx context.Context, in *SyncCMDBRequest, opts ...grpc.CallOption) (*SyncCMDBResponse, error)
}

type resourceTreeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewResourceTreeServiceClient(cc grpc.ClientConnInterface) ResourceTreeServiceClient {
	return &resourceTreeServiceClient{cc}
}

func (c *resourceTreeServiceClient) ListTreeNode(ctx context.Context, in *ListTreeNodeRequest, opts ...grpc.CallOption) (*ListTreeNodeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListTreeNodeResponse)
	err := c.cc.Invoke(ctx, ResourceTreeService_ListTreeNode_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceTreeServiceClient) SelectTreeNode(ctx context.Context, in *SelectTreeNodeRequest, opts ...grpc.CallOption) (*SelectTreeNodeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SelectTreeNodeResponse)
	err := c.cc.Invoke(ctx, ResourceTreeService_SelectTreeNode_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceTreeServiceClient) GetTopTreeNode(ctx context.Context, in *GetTopTreeNodeRequest, opts ...grpc.CallOption) (*GetTopTreeNodeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTopTreeNodeResponse)
	err := c.cc.Invoke(ctx, ResourceTreeService_GetTopTreeNode_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceTreeServiceClient) ListLeafTreeNode(ctx context.Context, in *ListLeafTreeNodeRequest, opts ...grpc.CallOption) (*ListLeafTreeNodeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListLeafTreeNodeResponse)
	err := c.cc.Invoke(ctx, ResourceTreeService_ListLeafTreeNode_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceTreeServiceClient) CreateTreeNode(ctx context.Context, in *CreateTreeNodeRequest, opts ...grpc.CallOption) (*CreateTreeNodeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateTreeNodeResponse)
	err := c.cc.Invoke(ctx, ResourceTreeService_CreateTreeNode_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceTreeServiceClient) DeleteTreeNode(ctx context.Context, in *DeleteTreeNodeRequest, opts ...grpc.CallOption) (*DeleteTreeNodeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteTreeNodeResponse)
	err := c.cc.Invoke(ctx, ResourceTreeService_DeleteTreeNode_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceTreeServiceClient) GetChildrenTreeNode(ctx context.Context, in *GetChildrenTreeNodeRequest, opts ...grpc.CallOption) (*GetChildrenTreeNodeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetChildrenTreeNodeResponse)
	err := c.cc.Invoke(ctx, ResourceTreeService_GetChildrenTreeNode_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceTreeServiceClient) UpdateTreeNode(ctx context.Context, in *UpdateTreeNodeRequest, opts ...grpc.CallOption) (*UpdateTreeNodeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateTreeNodeResponse)
	err := c.cc.Invoke(ctx, ResourceTreeService_UpdateTreeNode_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceTreeServiceClient) SyncCMDB(ctx context.Context, in *SyncCMDBRequest, opts ...grpc.CallOption) (*SyncCMDBResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SyncCMDBResponse)
	err := c.cc.Invoke(ctx, ResourceTreeService_SyncCMDB_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResourceTreeServiceServer is the server API for ResourceTreeService service.
// All implementations must embed UnimplementedResourceTreeServiceServer
// for forward compatibility
//
// 资源树服务
type ResourceTreeServiceServer interface {
	// 获取树节点列表
	ListTreeNode(context.Context, *ListTreeNodeRequest) (*ListTreeNodeResponse, error)
	// 选择树节点
	SelectTreeNode(context.Context, *SelectTreeNodeRequest) (*SelectTreeNodeResponse, error)
	// 获取顶级树节点
	GetTopTreeNode(context.Context, *GetTopTreeNodeRequest) (*GetTopTreeNodeResponse, error)
	// 获取叶子节点列表
	ListLeafTreeNode(context.Context, *ListLeafTreeNodeRequest) (*ListLeafTreeNodeResponse, error)
	// 创建树节点
	CreateTreeNode(context.Context, *CreateTreeNodeRequest) (*CreateTreeNodeResponse, error)
	// 删除树节点
	DeleteTreeNode(context.Context, *DeleteTreeNodeRequest) (*DeleteTreeNodeResponse, error)
	// 获取子节点
	GetChildrenTreeNode(context.Context, *GetChildrenTreeNodeRequest) (*GetChildrenTreeNodeResponse, error)
	// 更新树节点
	UpdateTreeNode(context.Context, *UpdateTreeNodeRequest) (*UpdateTreeNodeResponse, error)
	// 同步CMDB资源
	SyncCMDB(context.Context, *SyncCMDBRequest) (*SyncCMDBResponse, error)
	mustEmbedUnimplementedResourceTreeServiceServer()
}

// UnimplementedResourceTreeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedResourceTreeServiceServer struct {
}

func (UnimplementedResourceTreeServiceServer) ListTreeNode(context.Context, *ListTreeNodeRequest) (*ListTreeNodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTreeNode not implemented")
}
func (UnimplementedResourceTreeServiceServer) SelectTreeNode(context.Context, *SelectTreeNodeRequest) (*SelectTreeNodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SelectTreeNode not implemented")
}
func (UnimplementedResourceTreeServiceServer) GetTopTreeNode(context.Context, *GetTopTreeNodeRequest) (*GetTopTreeNodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTopTreeNode not implemented")
}
func (UnimplementedResourceTreeServiceServer) ListLeafTreeNode(context.Context, *ListLeafTreeNodeRequest) (*ListLeafTreeNodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListLeafTreeNode not implemented")
}
func (UnimplementedResourceTreeServiceServer) CreateTreeNode(context.Context, *CreateTreeNodeRequest) (*CreateTreeNodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTreeNode not implemented")
}
func (UnimplementedResourceTreeServiceServer) DeleteTreeNode(context.Context, *DeleteTreeNodeRequest) (*DeleteTreeNodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTreeNode not implemented")
}
func (UnimplementedResourceTreeServiceServer) GetChildrenTreeNode(context.Context, *GetChildrenTreeNodeRequest) (*GetChildrenTreeNodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChildrenTreeNode not implemented")
}
func (UnimplementedResourceTreeServiceServer) UpdateTreeNode(context.Context, *UpdateTreeNodeRequest) (*UpdateTreeNodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTreeNode not implemented")
}
func (UnimplementedResourceTreeServiceServer) SyncCMDB(context.Context, *SyncCMDBRequest) (*SyncCMDBResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyncCMDB not implemented")
}
func (UnimplementedResourceTreeServiceServer) mustEmbedUnimplementedResourceTreeServiceServer() {}

// UnsafeResourceTreeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ResourceTreeServiceServer will
// result in compilation errors.
type UnsafeResourceTreeServiceServer interface {
	mustEmbedUnimplementedResourceTreeServiceServer()
}

func RegisterResourceTreeServiceServer(s grpc.ServiceRegistrar, srv ResourceTreeServiceServer) {
	s.RegisterService(&ResourceTreeService_ServiceDesc, srv)
}

func _ResourceTreeService_ListTreeNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTreeNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceTreeServiceServer).ListTreeNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResourceTreeService_ListTreeNode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceTreeServiceServer).ListTreeNode(ctx, req.(*ListTreeNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceTreeService_SelectTreeNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SelectTreeNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceTreeServiceServer).SelectTreeNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResourceTreeService_SelectTreeNode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceTreeServiceServer).SelectTreeNode(ctx, req.(*SelectTreeNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceTreeService_GetTopTreeNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTopTreeNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceTreeServiceServer).GetTopTreeNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResourceTreeService_GetTopTreeNode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceTreeServiceServer).GetTopTreeNode(ctx, req.(*GetTopTreeNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceTreeService_ListLeafTreeNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListLeafTreeNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceTreeServiceServer).ListLeafTreeNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResourceTreeService_ListLeafTreeNode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceTreeServiceServer).ListLeafTreeNode(ctx, req.(*ListLeafTreeNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceTreeService_CreateTreeNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTreeNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceTreeServiceServer).CreateTreeNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResourceTreeService_CreateTreeNode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceTreeServiceServer).CreateTreeNode(ctx, req.(*CreateTreeNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceTreeService_DeleteTreeNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTreeNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceTreeServiceServer).DeleteTreeNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResourceTreeService_DeleteTreeNode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceTreeServiceServer).DeleteTreeNode(ctx, req.(*DeleteTreeNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceTreeService_GetChildrenTreeNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChildrenTreeNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceTreeServiceServer).GetChildrenTreeNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResourceTreeService_GetChildrenTreeNode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceTreeServiceServer).GetChildrenTreeNode(ctx, req.(*GetChildrenTreeNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceTreeService_UpdateTreeNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTreeNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceTreeServiceServer).UpdateTreeNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResourceTreeService_UpdateTreeNode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceTreeServiceServer).UpdateTreeNode(ctx, req.(*UpdateTreeNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceTreeService_SyncCMDB_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncCMDBRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceTreeServiceServer).SyncCMDB(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResourceTreeService_SyncCMDB_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceTreeServiceServer).SyncCMDB(ctx, req.(*SyncCMDBRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ResourceTreeService_ServiceDesc is the grpc.ServiceDesc for ResourceTreeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ResourceTreeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tree_rpc.ResourceTreeService",
	HandlerType: (*ResourceTreeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListTreeNode",
			Handler:    _ResourceTreeService_ListTreeNode_Handler,
		},
		{
			MethodName: "SelectTreeNode",
			Handler:    _ResourceTreeService_SelectTreeNode_Handler,
		},
		{
			MethodName: "GetTopTreeNode",
			Handler:    _ResourceTreeService_GetTopTreeNode_Handler,
		},
		{
			MethodName: "ListLeafTreeNode",
			Handler:    _ResourceTreeService_ListLeafTreeNode_Handler,
		},
		{
			MethodName: "CreateTreeNode",
			Handler:    _ResourceTreeService_CreateTreeNode_Handler,
		},
		{
			MethodName: "DeleteTreeNode",
			Handler:    _ResourceTreeService_DeleteTreeNode_Handler,
		},
		{
			MethodName: "GetChildrenTreeNode",
			Handler:    _ResourceTreeService_GetChildrenTreeNode_Handler,
		},
		{
			MethodName: "UpdateTreeNode",
			Handler:    _ResourceTreeService_UpdateTreeNode_Handler,
		},
		{
			MethodName: "SyncCMDB",
			Handler:    _ResourceTreeService_SyncCMDB_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "aicoreops_tree.proto",
}
