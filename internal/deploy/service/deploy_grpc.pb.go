// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package service

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

// DeployClient is the client API for Deploy service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DeployClient interface {
	DeployProject(ctx context.Context, in *DeployRequest, opts ...grpc.CallOption) (*DeployReply, error)
	CheckStatus(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusReply, error)
	DeleteProjectDeployment(ctx context.Context, in *DeleteProjectDeploymentRequest, opts ...grpc.CallOption) (*DeleteProjectDeploymentReply, error)
}

type deployClient struct {
	cc grpc.ClientConnInterface
}

func NewDeployClient(cc grpc.ClientConnInterface) DeployClient {
	return &deployClient{cc}
}

func (c *deployClient) DeployProject(ctx context.Context, in *DeployRequest, opts ...grpc.CallOption) (*DeployReply, error) {
	out := new(DeployReply)
	err := c.cc.Invoke(ctx, "/service.Deploy/DeployProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deployClient) CheckStatus(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusReply, error) {
	out := new(StatusReply)
	err := c.cc.Invoke(ctx, "/service.Deploy/CheckStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deployClient) DeleteProjectDeployment(ctx context.Context, in *DeleteProjectDeploymentRequest, opts ...grpc.CallOption) (*DeleteProjectDeploymentReply, error) {
	out := new(DeleteProjectDeploymentReply)
	err := c.cc.Invoke(ctx, "/service.Deploy/DeleteProjectDeployment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeployServer is the server API for Deploy service.
// All implementations must embed UnimplementedDeployServer
// for forward compatibility
type DeployServer interface {
	DeployProject(context.Context, *DeployRequest) (*DeployReply, error)
	CheckStatus(context.Context, *StatusRequest) (*StatusReply, error)
	DeleteProjectDeployment(context.Context, *DeleteProjectDeploymentRequest) (*DeleteProjectDeploymentReply, error)
	mustEmbedUnimplementedDeployServer()
}

// UnimplementedDeployServer must be embedded to have forward compatible implementations.
type UnimplementedDeployServer struct {
}

func (UnimplementedDeployServer) DeployProject(context.Context, *DeployRequest) (*DeployReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeployProject not implemented")
}
func (UnimplementedDeployServer) CheckStatus(context.Context, *StatusRequest) (*StatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckStatus not implemented")
}
func (UnimplementedDeployServer) DeleteProjectDeployment(context.Context, *DeleteProjectDeploymentRequest) (*DeleteProjectDeploymentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProjectDeployment not implemented")
}
func (UnimplementedDeployServer) mustEmbedUnimplementedDeployServer() {}

// UnsafeDeployServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DeployServer will
// result in compilation errors.
type UnsafeDeployServer interface {
	mustEmbedUnimplementedDeployServer()
}

func RegisterDeployServer(s grpc.ServiceRegistrar, srv DeployServer) {
	s.RegisterService(&Deploy_ServiceDesc, srv)
}

func _Deploy_DeployProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeployRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeployServer).DeployProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Deploy/DeployProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeployServer).DeployProject(ctx, req.(*DeployRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Deploy_CheckStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeployServer).CheckStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Deploy/CheckStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeployServer).CheckStatus(ctx, req.(*StatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Deploy_DeleteProjectDeployment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProjectDeploymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeployServer).DeleteProjectDeployment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Deploy/DeleteProjectDeployment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeployServer).DeleteProjectDeployment(ctx, req.(*DeleteProjectDeploymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Deploy_ServiceDesc is the grpc.ServiceDesc for Deploy service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Deploy_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.Deploy",
	HandlerType: (*DeployServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeployProject",
			Handler:    _Deploy_DeployProject_Handler,
		},
		{
			MethodName: "CheckStatus",
			Handler:    _Deploy_CheckStatus_Handler,
		},
		{
			MethodName: "DeleteProjectDeployment",
			Handler:    _Deploy_DeleteProjectDeployment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "deploy.proto",
}
