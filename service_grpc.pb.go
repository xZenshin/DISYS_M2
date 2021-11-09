// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package DISYS_M2

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TokenRingClient is the client API for TokenRing service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TokenRingClient interface {
	GrantToken(ctx context.Context, in *Token, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type tokenRingClient struct {
	cc grpc.ClientConnInterface
}

func NewTokenRingClient(cc grpc.ClientConnInterface) TokenRingClient {
	return &tokenRingClient{cc}
}

func (c *tokenRingClient) GrantToken(ctx context.Context, in *Token, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/TokenRing/grantToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TokenRingServer is the server API for TokenRing service.
// All implementations must embed UnimplementedTokenRingServer
// for forward compatibility
type TokenRingServer interface {
	GrantToken(context.Context, *Token) (*emptypb.Empty, error)
	mustEmbedUnimplementedTokenRingServer()
}

// UnimplementedTokenRingServer must be embedded to have forward compatible implementations.
type UnimplementedTokenRingServer struct {
}

func (UnimplementedTokenRingServer) GrantToken(context.Context, *Token) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GrantToken not implemented")
}
func (UnimplementedTokenRingServer) mustEmbedUnimplementedTokenRingServer() {}

// UnsafeTokenRingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TokenRingServer will
// result in compilation errors.
type UnsafeTokenRingServer interface {
	mustEmbedUnimplementedTokenRingServer()
}

func RegisterTokenRingServer(s grpc.ServiceRegistrar, srv TokenRingServer) {
	s.RegisterService(&TokenRing_ServiceDesc, srv)
}

func _TokenRing_GrantToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Token)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenRingServer).GrantToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TokenRing/grantToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenRingServer).GrantToken(ctx, req.(*Token))
	}
	return interceptor(ctx, in, info, handler)
}

// TokenRing_ServiceDesc is the grpc.ServiceDesc for TokenRing service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TokenRing_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "TokenRing",
	HandlerType: (*TokenRingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "grantToken",
			Handler:    _TokenRing_GrantToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
