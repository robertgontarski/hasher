// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: proto/hasher.proto

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

const (
	HasherService_HashEmail_FullMethodName = "/hasher.HasherService/HashEmail"
	HasherService_HashPhone_FullMethodName = "/hasher.HasherService/HashPhone"
	HasherService_HashName_FullMethodName  = "/hasher.HasherService/HashName"
)

// HasherServiceClient is the client API for HasherService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HasherServiceClient interface {
	HashEmail(ctx context.Context, in *HashEmailRequest, opts ...grpc.CallOption) (*HashResponse, error)
	HashPhone(ctx context.Context, in *HashPhoneRequest, opts ...grpc.CallOption) (*HashResponse, error)
	HashName(ctx context.Context, in *HashNameRequest, opts ...grpc.CallOption) (*HashResponse, error)
}

type hasherServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHasherServiceClient(cc grpc.ClientConnInterface) HasherServiceClient {
	return &hasherServiceClient{cc}
}

func (c *hasherServiceClient) HashEmail(ctx context.Context, in *HashEmailRequest, opts ...grpc.CallOption) (*HashResponse, error) {
	out := new(HashResponse)
	err := c.cc.Invoke(ctx, HasherService_HashEmail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hasherServiceClient) HashPhone(ctx context.Context, in *HashPhoneRequest, opts ...grpc.CallOption) (*HashResponse, error) {
	out := new(HashResponse)
	err := c.cc.Invoke(ctx, HasherService_HashPhone_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hasherServiceClient) HashName(ctx context.Context, in *HashNameRequest, opts ...grpc.CallOption) (*HashResponse, error) {
	out := new(HashResponse)
	err := c.cc.Invoke(ctx, HasherService_HashName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HasherServiceServer is the server API for HasherService service.
// All implementations must embed UnimplementedHasherServiceServer
// for forward compatibility
type HasherServiceServer interface {
	HashEmail(context.Context, *HashEmailRequest) (*HashResponse, error)
	HashPhone(context.Context, *HashPhoneRequest) (*HashResponse, error)
	HashName(context.Context, *HashNameRequest) (*HashResponse, error)
	mustEmbedUnimplementedHasherServiceServer()
}

// UnimplementedHasherServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHasherServiceServer struct {
}

func (UnimplementedHasherServiceServer) HashEmail(context.Context, *HashEmailRequest) (*HashResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HashEmail not implemented")
}
func (UnimplementedHasherServiceServer) HashPhone(context.Context, *HashPhoneRequest) (*HashResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HashPhone not implemented")
}
func (UnimplementedHasherServiceServer) HashName(context.Context, *HashNameRequest) (*HashResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HashName not implemented")
}
func (UnimplementedHasherServiceServer) mustEmbedUnimplementedHasherServiceServer() {}

// UnsafeHasherServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HasherServiceServer will
// result in compilation errors.
type UnsafeHasherServiceServer interface {
	mustEmbedUnimplementedHasherServiceServer()
}

func RegisterHasherServiceServer(s grpc.ServiceRegistrar, srv HasherServiceServer) {
	s.RegisterService(&HasherService_ServiceDesc, srv)
}

func _HasherService_HashEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HashEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HasherServiceServer).HashEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HasherService_HashEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HasherServiceServer).HashEmail(ctx, req.(*HashEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HasherService_HashPhone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HashPhoneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HasherServiceServer).HashPhone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HasherService_HashPhone_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HasherServiceServer).HashPhone(ctx, req.(*HashPhoneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HasherService_HashName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HashNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HasherServiceServer).HashName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HasherService_HashName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HasherServiceServer).HashName(ctx, req.(*HashNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HasherService_ServiceDesc is the grpc.ServiceDesc for HasherService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HasherService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hasher.HasherService",
	HandlerType: (*HasherServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HashEmail",
			Handler:    _HasherService_HashEmail_Handler,
		},
		{
			MethodName: "HashPhone",
			Handler:    _HasherService_HashPhone_Handler,
		},
		{
			MethodName: "HashName",
			Handler:    _HasherService_HashName_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/hasher.proto",
}
