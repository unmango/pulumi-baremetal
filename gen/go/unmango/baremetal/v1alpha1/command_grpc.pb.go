// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             (unknown)
// source: unmango/baremetal/v1alpha1/command.proto

package baremetalv1alpha1

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
	CommandService_Exec_FullMethodName   = "/unmango.baremetal.v1alpha1.CommandService/Exec"
	CommandService_Create_FullMethodName = "/unmango.baremetal.v1alpha1.CommandService/Create"
	CommandService_Update_FullMethodName = "/unmango.baremetal.v1alpha1.CommandService/Update"
	CommandService_Delete_FullMethodName = "/unmango.baremetal.v1alpha1.CommandService/Delete"
)

// CommandServiceClient is the client API for CommandService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommandServiceClient interface {
	Exec(ctx context.Context, in *ExecRequest, opts ...grpc.CallOption) (*ExecResponse, error)
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
}

type commandServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCommandServiceClient(cc grpc.ClientConnInterface) CommandServiceClient {
	return &commandServiceClient{cc}
}

func (c *commandServiceClient) Exec(ctx context.Context, in *ExecRequest, opts ...grpc.CallOption) (*ExecResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ExecResponse)
	err := c.cc.Invoke(ctx, CommandService_Exec_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commandServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, CommandService_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commandServiceClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, CommandService_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commandServiceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, CommandService_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommandServiceServer is the server API for CommandService service.
// All implementations must embed UnimplementedCommandServiceServer
// for forward compatibility
type CommandServiceServer interface {
	Exec(context.Context, *ExecRequest) (*ExecResponse, error)
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	mustEmbedUnimplementedCommandServiceServer()
}

// UnimplementedCommandServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCommandServiceServer struct {
}

func (UnimplementedCommandServiceServer) Exec(context.Context, *ExecRequest) (*ExecResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Exec not implemented")
}
func (UnimplementedCommandServiceServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedCommandServiceServer) Update(context.Context, *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedCommandServiceServer) Delete(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedCommandServiceServer) mustEmbedUnimplementedCommandServiceServer() {}

// UnsafeCommandServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommandServiceServer will
// result in compilation errors.
type UnsafeCommandServiceServer interface {
	mustEmbedUnimplementedCommandServiceServer()
}

func RegisterCommandServiceServer(s grpc.ServiceRegistrar, srv CommandServiceServer) {
	s.RegisterService(&CommandService_ServiceDesc, srv)
}

func _CommandService_Exec_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommandServiceServer).Exec(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommandService_Exec_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommandServiceServer).Exec(ctx, req.(*ExecRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommandService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommandServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommandService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommandServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommandService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommandServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommandService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommandServiceServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommandService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommandServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommandService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommandServiceServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CommandService_ServiceDesc is the grpc.ServiceDesc for CommandService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CommandService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "unmango.baremetal.v1alpha1.CommandService",
	HandlerType: (*CommandServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Exec",
			Handler:    _CommandService_Exec_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _CommandService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _CommandService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _CommandService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "unmango/baremetal/v1alpha1/command.proto",
}
