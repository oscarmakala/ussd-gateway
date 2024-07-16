// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: proto/ussd/v1/ussd.proto

package ussd

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

// UssdServiceClient is the client API for UssdService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UssdServiceClient interface {
	HandleUssdRequest(ctx context.Context, in *UssdRequest, opts ...grpc.CallOption) (*UssdResponse, error)
}

type ussdServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUssdServiceClient(cc grpc.ClientConnInterface) UssdServiceClient {
	return &ussdServiceClient{cc}
}

func (c *ussdServiceClient) HandleUssdRequest(ctx context.Context, in *UssdRequest, opts ...grpc.CallOption) (*UssdResponse, error) {
	out := new(UssdResponse)
	err := c.cc.Invoke(ctx, "/ussd.UssdService/HandleUssdRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UssdServiceServer is the server API for UssdService service.
// All implementations must embed UnimplementedUssdServiceServer
// for forward compatibility
type UssdServiceServer interface {
	HandleUssdRequest(context.Context, *UssdRequest) (*UssdResponse, error)
	mustEmbedUnimplementedUssdServiceServer()
}

// UnimplementedUssdServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUssdServiceServer struct {
}

func (UnimplementedUssdServiceServer) HandleUssdRequest(context.Context, *UssdRequest) (*UssdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleUssdRequest not implemented")
}
func (UnimplementedUssdServiceServer) mustEmbedUnimplementedUssdServiceServer() {}

// UnsafeUssdServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UssdServiceServer will
// result in compilation errors.
type UnsafeUssdServiceServer interface {
	mustEmbedUnimplementedUssdServiceServer()
}

func RegisterUssdServiceServer(s grpc.ServiceRegistrar, srv UssdServiceServer) {
	s.RegisterService(&UssdService_ServiceDesc, srv)
}

func _UssdService_HandleUssdRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UssdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UssdServiceServer).HandleUssdRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ussd.UssdService/HandleUssdRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UssdServiceServer).HandleUssdRequest(ctx, req.(*UssdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UssdService_ServiceDesc is the grpc.ServiceDesc for UssdService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UssdService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ussd.UssdService",
	HandlerType: (*UssdServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HandleUssdRequest",
			Handler:    _UssdService_HandleUssdRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/ussd/v1/ussd.proto",
}
