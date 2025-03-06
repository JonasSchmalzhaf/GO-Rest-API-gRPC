// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: service.proto

package gRPC

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
	DatabaseService_GetMultipleDBs_FullMethodName = "/DatabaseService/getMultipleDBs"
	DatabaseService_GetSingleDB_FullMethodName    = "/DatabaseService/getSingleDB"
	DatabaseService_CreateSingleDB_FullMethodName = "/DatabaseService/createSingleDB"
	DatabaseService_UpdateSingleDB_FullMethodName = "/DatabaseService/updateSingleDB"
)

// DatabaseServiceClient is the client API for DatabaseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DatabaseServiceClient interface {
	GetMultipleDBs(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	GetSingleDB(ctx context.Context, in *GetSingleRequest, opts ...grpc.CallOption) (*GetSingleResponse, error)
	CreateSingleDB(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	UpdateSingleDB(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
}

type databaseServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDatabaseServiceClient(cc grpc.ClientConnInterface) DatabaseServiceClient {
	return &databaseServiceClient{cc}
}

func (c *databaseServiceClient) GetMultipleDBs(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, DatabaseService_GetMultipleDBs_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseServiceClient) GetSingleDB(ctx context.Context, in *GetSingleRequest, opts ...grpc.CallOption) (*GetSingleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetSingleResponse)
	err := c.cc.Invoke(ctx, DatabaseService_GetSingleDB_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseServiceClient) CreateSingleDB(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, DatabaseService_CreateSingleDB_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseServiceClient) UpdateSingleDB(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, DatabaseService_UpdateSingleDB_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DatabaseServiceServer is the server API for DatabaseService service.
// All implementations must embed UnimplementedDatabaseServiceServer
// for forward compatibility.
type DatabaseServiceServer interface {
	GetMultipleDBs(context.Context, *GetRequest) (*GetResponse, error)
	GetSingleDB(context.Context, *GetSingleRequest) (*GetSingleResponse, error)
	CreateSingleDB(context.Context, *CreateRequest) (*CreateResponse, error)
	UpdateSingleDB(context.Context, *UpdateRequest) (*UpdateResponse, error)
	mustEmbedUnimplementedDatabaseServiceServer()
}

// UnimplementedDatabaseServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDatabaseServiceServer struct{}

func (UnimplementedDatabaseServiceServer) GetMultipleDBs(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMultipleDBs not implemented")
}
func (UnimplementedDatabaseServiceServer) GetSingleDB(context.Context, *GetSingleRequest) (*GetSingleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSingleDB not implemented")
}
func (UnimplementedDatabaseServiceServer) CreateSingleDB(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSingleDB not implemented")
}
func (UnimplementedDatabaseServiceServer) UpdateSingleDB(context.Context, *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSingleDB not implemented")
}
func (UnimplementedDatabaseServiceServer) mustEmbedUnimplementedDatabaseServiceServer() {}
func (UnimplementedDatabaseServiceServer) testEmbeddedByValue()                         {}

// UnsafeDatabaseServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DatabaseServiceServer will
// result in compilation errors.
type UnsafeDatabaseServiceServer interface {
	mustEmbedUnimplementedDatabaseServiceServer()
}

func RegisterDatabaseServiceServer(s grpc.ServiceRegistrar, srv DatabaseServiceServer) {
	// If the following call pancis, it indicates UnimplementedDatabaseServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&DatabaseService_ServiceDesc, srv)
}

func _DatabaseService_GetMultipleDBs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServiceServer).GetMultipleDBs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DatabaseService_GetMultipleDBs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServiceServer).GetMultipleDBs(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatabaseService_GetSingleDB_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSingleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServiceServer).GetSingleDB(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DatabaseService_GetSingleDB_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServiceServer).GetSingleDB(ctx, req.(*GetSingleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatabaseService_CreateSingleDB_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServiceServer).CreateSingleDB(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DatabaseService_CreateSingleDB_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServiceServer).CreateSingleDB(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatabaseService_UpdateSingleDB_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServiceServer).UpdateSingleDB(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DatabaseService_UpdateSingleDB_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServiceServer).UpdateSingleDB(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DatabaseService_ServiceDesc is the grpc.ServiceDesc for DatabaseService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DatabaseService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "DatabaseService",
	HandlerType: (*DatabaseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getMultipleDBs",
			Handler:    _DatabaseService_GetMultipleDBs_Handler,
		},
		{
			MethodName: "getSingleDB",
			Handler:    _DatabaseService_GetSingleDB_Handler,
		},
		{
			MethodName: "createSingleDB",
			Handler:    _DatabaseService_CreateSingleDB_Handler,
		},
		{
			MethodName: "updateSingleDB",
			Handler:    _DatabaseService_UpdateSingleDB_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
