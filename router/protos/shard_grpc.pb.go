// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// ShardServiceClient is the client API for ShardService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShardServiceClient interface {
	ListShards(ctx context.Context, in *ShardRequest, opts ...grpc.CallOption) (*ShardReply, error)
	AddDataShard(ctx context.Context, in *AddShardRequest, opts ...grpc.CallOption) (*AddShardReply, error)
	AddWorldShard(ctx context.Context, in *AddWorldShardRequest, opts ...grpc.CallOption) (*AddShardReply, error)
}

type shardServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewShardServiceClient(cc grpc.ClientConnInterface) ShardServiceClient {
	return &shardServiceClient{cc}
}

func (c *shardServiceClient) ListShards(ctx context.Context, in *ShardRequest, opts ...grpc.CallOption) (*ShardReply, error) {
	out := new(ShardReply)
	err := c.cc.Invoke(ctx, "/yandex.spqr.ShardService/ListShards", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shardServiceClient) AddDataShard(ctx context.Context, in *AddShardRequest, opts ...grpc.CallOption) (*AddShardReply, error) {
	out := new(AddShardReply)
	err := c.cc.Invoke(ctx, "/yandex.spqr.ShardService/AddDataShard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shardServiceClient) AddWorldShard(ctx context.Context, in *AddWorldShardRequest, opts ...grpc.CallOption) (*AddShardReply, error) {
	out := new(AddShardReply)
	err := c.cc.Invoke(ctx, "/yandex.spqr.ShardService/AddWorldShard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShardServiceServer is the server API for ShardService service.
// All implementations must embed UnimplementedShardServiceServer
// for forward compatibility
type ShardServiceServer interface {
	ListShards(context.Context, *ShardRequest) (*ShardReply, error)
	AddDataShard(context.Context, *AddShardRequest) (*AddShardReply, error)
	AddWorldShard(context.Context, *AddWorldShardRequest) (*AddShardReply, error)
	mustEmbedUnimplementedShardServiceServer()
}

// UnimplementedShardServiceServer must be embedded to have forward compatible implementations.
type UnimplementedShardServiceServer struct {
}

func (UnimplementedShardServiceServer) ListShards(context.Context, *ShardRequest) (*ShardReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListShards not implemented")
}
func (UnimplementedShardServiceServer) AddDataShard(context.Context, *AddShardRequest) (*AddShardReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddDataShard not implemented")
}
func (UnimplementedShardServiceServer) AddWorldShard(context.Context, *AddWorldShardRequest) (*AddShardReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddWorldShard not implemented")
}
func (UnimplementedShardServiceServer) mustEmbedUnimplementedShardServiceServer() {}

// UnsafeShardServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShardServiceServer will
// result in compilation errors.
type UnsafeShardServiceServer interface {
	mustEmbedUnimplementedShardServiceServer()
}

func RegisterShardServiceServer(s grpc.ServiceRegistrar, srv ShardServiceServer) {
	s.RegisterService(&ShardService_ServiceDesc, srv)
}

func _ShardService_ListShards_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShardServiceServer).ListShards(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.spqr.ShardService/ListShards",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShardServiceServer).ListShards(ctx, req.(*ShardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShardService_AddDataShard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddShardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShardServiceServer).AddDataShard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.spqr.ShardService/AddDataShard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShardServiceServer).AddDataShard(ctx, req.(*AddShardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShardService_AddWorldShard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddWorldShardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShardServiceServer).AddWorldShard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.spqr.ShardService/AddWorldShard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShardServiceServer).AddWorldShard(ctx, req.(*AddWorldShardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ShardService_ServiceDesc is the grpc.ServiceDesc for ShardService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ShardService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.spqr.ShardService",
	HandlerType: (*ShardServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListShards",
			Handler:    _ShardService_ListShards_Handler,
		},
		{
			MethodName: "AddDataShard",
			Handler:    _ShardService_AddDataShard_Handler,
		},
		{
			MethodName: "AddWorldShard",
			Handler:    _ShardService_AddWorldShard_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/shard.proto",
}
