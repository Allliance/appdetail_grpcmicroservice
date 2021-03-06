// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package appdetail

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

// AppDetailClient is the client API for AppDetail service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AppDetailClient interface {
	GetAppDetail(ctx context.Context, in *GetAppDetailRequest, opts ...grpc.CallOption) (*GetAppDetailReply, error)
}

type appDetailClient struct {
	cc grpc.ClientConnInterface
}

func NewAppDetailClient(cc grpc.ClientConnInterface) AppDetailClient {
	return &appDetailClient{cc}
}

func (c *appDetailClient) GetAppDetail(ctx context.Context, in *GetAppDetailRequest, opts ...grpc.CallOption) (*GetAppDetailReply, error) {
	out := new(GetAppDetailReply)
	err := c.cc.Invoke(ctx, "/appdetail.AppDetail/GetAppDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppDetailServer is the server API for AppDetail service.
// All implementations must embed UnimplementedAppDetailServer
// for forward compatibility
type AppDetailServer interface {
	GetAppDetail(context.Context, *GetAppDetailRequest) (*GetAppDetailReply, error)
	mustEmbedUnimplementedAppDetailServer()
}

// UnimplementedAppDetailServer must be embedded to have forward compatible implementations.
type UnimplementedAppDetailServer struct {
}

func (UnimplementedAppDetailServer) GetAppDetail(context.Context, *GetAppDetailRequest) (*GetAppDetailReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppDetail not implemented")
}
func (UnimplementedAppDetailServer) mustEmbedUnimplementedAppDetailServer() {}

// UnsafeAppDetailServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AppDetailServer will
// result in compilation errors.
type UnsafeAppDetailServer interface {
	mustEmbedUnimplementedAppDetailServer()
}

func RegisterAppDetailServer(s grpc.ServiceRegistrar, srv AppDetailServer) {
	s.RegisterService(&AppDetail_ServiceDesc, srv)
}

func _AppDetail_GetAppDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppDetailServer).GetAppDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appdetail.AppDetail/GetAppDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppDetailServer).GetAppDetail(ctx, req.(*GetAppDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AppDetail_ServiceDesc is the grpc.ServiceDesc for AppDetail service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AppDetail_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "appdetail.AppDetail",
	HandlerType: (*AppDetailServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAppDetail",
			Handler:    _AppDetail_GetAppDetail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "appdetail.proto",
}
