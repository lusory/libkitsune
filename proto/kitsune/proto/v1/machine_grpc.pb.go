// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.4
// source: kitsune/proto/v1/machine.proto

package v1

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

// VirtualMachineRegistryServiceClient is the client API for VirtualMachineRegistryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VirtualMachineRegistryServiceClient interface {
	GetVirtualMachines(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (VirtualMachineRegistryService_GetVirtualMachinesClient, error)
	FindVirtualMachine(ctx context.Context, in *FindVirtualMachineRequest, opts ...grpc.CallOption) (*FindVirtualMachineResponse, error)
	CreateVirtualMachine(ctx context.Context, in *CreateVirtualMachineRequest, opts ...grpc.CallOption) (*CreateVirtualMachineResponse, error)
	DeleteVirtualMachine(ctx context.Context, in *DeleteVirtualMachineRequest, opts ...grpc.CallOption) (*DeleteVirtualMachineResponse, error)
	IsAlive(ctx context.Context, in *IsAliveRequest, opts ...grpc.CallOption) (*IsAliveResponse, error)
	GetAttachedImages(ctx context.Context, in *GetAttachedImagesRequest, opts ...grpc.CallOption) (*GetAttachedImagesResponse, error)
	AttachImage(ctx context.Context, in *AttachImageRequest, opts ...grpc.CallOption) (*AttachImageResponse, error)
	DetachImage(ctx context.Context, in *DetachImageRequest, opts ...grpc.CallOption) (*DetachImageResponse, error)
	SendACPIAction(ctx context.Context, in *SendACPIActionRequest, opts ...grpc.CallOption) (*SendACPIActionResponse, error)
	GetVNCServer(ctx context.Context, in *GetVNCServerRequest, opts ...grpc.CallOption) (*GetVNCServerResponse, error)
	GetMetadata(ctx context.Context, in *GetMetadataRequest, opts ...grpc.CallOption) (*GetMetadataResponse, error)
	SetMetadata(ctx context.Context, in *SetMetadataRequest, opts ...grpc.CallOption) (*SetMetadataResponse, error)
}

type virtualMachineRegistryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVirtualMachineRegistryServiceClient(cc grpc.ClientConnInterface) VirtualMachineRegistryServiceClient {
	return &virtualMachineRegistryServiceClient{cc}
}

func (c *virtualMachineRegistryServiceClient) GetVirtualMachines(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (VirtualMachineRegistryService_GetVirtualMachinesClient, error) {
	stream, err := c.cc.NewStream(ctx, &VirtualMachineRegistryService_ServiceDesc.Streams[0], "/kitsune.proto.v1.VirtualMachineRegistryService/GetVirtualMachines", opts...)
	if err != nil {
		return nil, err
	}
	x := &virtualMachineRegistryServiceGetVirtualMachinesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type VirtualMachineRegistryService_GetVirtualMachinesClient interface {
	Recv() (*VirtualMachine, error)
	grpc.ClientStream
}

type virtualMachineRegistryServiceGetVirtualMachinesClient struct {
	grpc.ClientStream
}

func (x *virtualMachineRegistryServiceGetVirtualMachinesClient) Recv() (*VirtualMachine, error) {
	m := new(VirtualMachine)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *virtualMachineRegistryServiceClient) FindVirtualMachine(ctx context.Context, in *FindVirtualMachineRequest, opts ...grpc.CallOption) (*FindVirtualMachineResponse, error) {
	out := new(FindVirtualMachineResponse)
	err := c.cc.Invoke(ctx, "/kitsune.proto.v1.VirtualMachineRegistryService/FindVirtualMachine", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *virtualMachineRegistryServiceClient) CreateVirtualMachine(ctx context.Context, in *CreateVirtualMachineRequest, opts ...grpc.CallOption) (*CreateVirtualMachineResponse, error) {
	out := new(CreateVirtualMachineResponse)
	err := c.cc.Invoke(ctx, "/kitsune.proto.v1.VirtualMachineRegistryService/CreateVirtualMachine", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *virtualMachineRegistryServiceClient) DeleteVirtualMachine(ctx context.Context, in *DeleteVirtualMachineRequest, opts ...grpc.CallOption) (*DeleteVirtualMachineResponse, error) {
	out := new(DeleteVirtualMachineResponse)
	err := c.cc.Invoke(ctx, "/kitsune.proto.v1.VirtualMachineRegistryService/DeleteVirtualMachine", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *virtualMachineRegistryServiceClient) IsAlive(ctx context.Context, in *IsAliveRequest, opts ...grpc.CallOption) (*IsAliveResponse, error) {
	out := new(IsAliveResponse)
	err := c.cc.Invoke(ctx, "/kitsune.proto.v1.VirtualMachineRegistryService/IsAlive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *virtualMachineRegistryServiceClient) GetAttachedImages(ctx context.Context, in *GetAttachedImagesRequest, opts ...grpc.CallOption) (*GetAttachedImagesResponse, error) {
	out := new(GetAttachedImagesResponse)
	err := c.cc.Invoke(ctx, "/kitsune.proto.v1.VirtualMachineRegistryService/GetAttachedImages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *virtualMachineRegistryServiceClient) AttachImage(ctx context.Context, in *AttachImageRequest, opts ...grpc.CallOption) (*AttachImageResponse, error) {
	out := new(AttachImageResponse)
	err := c.cc.Invoke(ctx, "/kitsune.proto.v1.VirtualMachineRegistryService/AttachImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *virtualMachineRegistryServiceClient) DetachImage(ctx context.Context, in *DetachImageRequest, opts ...grpc.CallOption) (*DetachImageResponse, error) {
	out := new(DetachImageResponse)
	err := c.cc.Invoke(ctx, "/kitsune.proto.v1.VirtualMachineRegistryService/DetachImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *virtualMachineRegistryServiceClient) SendACPIAction(ctx context.Context, in *SendACPIActionRequest, opts ...grpc.CallOption) (*SendACPIActionResponse, error) {
	out := new(SendACPIActionResponse)
	err := c.cc.Invoke(ctx, "/kitsune.proto.v1.VirtualMachineRegistryService/SendACPIAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *virtualMachineRegistryServiceClient) GetVNCServer(ctx context.Context, in *GetVNCServerRequest, opts ...grpc.CallOption) (*GetVNCServerResponse, error) {
	out := new(GetVNCServerResponse)
	err := c.cc.Invoke(ctx, "/kitsune.proto.v1.VirtualMachineRegistryService/GetVNCServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *virtualMachineRegistryServiceClient) GetMetadata(ctx context.Context, in *GetMetadataRequest, opts ...grpc.CallOption) (*GetMetadataResponse, error) {
	out := new(GetMetadataResponse)
	err := c.cc.Invoke(ctx, "/kitsune.proto.v1.VirtualMachineRegistryService/GetMetadata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *virtualMachineRegistryServiceClient) SetMetadata(ctx context.Context, in *SetMetadataRequest, opts ...grpc.CallOption) (*SetMetadataResponse, error) {
	out := new(SetMetadataResponse)
	err := c.cc.Invoke(ctx, "/kitsune.proto.v1.VirtualMachineRegistryService/SetMetadata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VirtualMachineRegistryServiceServer is the server API for VirtualMachineRegistryService service.
// All implementations must embed UnimplementedVirtualMachineRegistryServiceServer
// for forward compatibility
type VirtualMachineRegistryServiceServer interface {
	GetVirtualMachines(*emptypb.Empty, VirtualMachineRegistryService_GetVirtualMachinesServer) error
	FindVirtualMachine(context.Context, *FindVirtualMachineRequest) (*FindVirtualMachineResponse, error)
	CreateVirtualMachine(context.Context, *CreateVirtualMachineRequest) (*CreateVirtualMachineResponse, error)
	DeleteVirtualMachine(context.Context, *DeleteVirtualMachineRequest) (*DeleteVirtualMachineResponse, error)
	IsAlive(context.Context, *IsAliveRequest) (*IsAliveResponse, error)
	GetAttachedImages(context.Context, *GetAttachedImagesRequest) (*GetAttachedImagesResponse, error)
	AttachImage(context.Context, *AttachImageRequest) (*AttachImageResponse, error)
	DetachImage(context.Context, *DetachImageRequest) (*DetachImageResponse, error)
	SendACPIAction(context.Context, *SendACPIActionRequest) (*SendACPIActionResponse, error)
	GetVNCServer(context.Context, *GetVNCServerRequest) (*GetVNCServerResponse, error)
	GetMetadata(context.Context, *GetMetadataRequest) (*GetMetadataResponse, error)
	SetMetadata(context.Context, *SetMetadataRequest) (*SetMetadataResponse, error)
	mustEmbedUnimplementedVirtualMachineRegistryServiceServer()
}

// UnimplementedVirtualMachineRegistryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedVirtualMachineRegistryServiceServer struct {
}

func (UnimplementedVirtualMachineRegistryServiceServer) GetVirtualMachines(*emptypb.Empty, VirtualMachineRegistryService_GetVirtualMachinesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetVirtualMachines not implemented")
}
func (UnimplementedVirtualMachineRegistryServiceServer) FindVirtualMachine(context.Context, *FindVirtualMachineRequest) (*FindVirtualMachineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindVirtualMachine not implemented")
}
func (UnimplementedVirtualMachineRegistryServiceServer) CreateVirtualMachine(context.Context, *CreateVirtualMachineRequest) (*CreateVirtualMachineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateVirtualMachine not implemented")
}
func (UnimplementedVirtualMachineRegistryServiceServer) DeleteVirtualMachine(context.Context, *DeleteVirtualMachineRequest) (*DeleteVirtualMachineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteVirtualMachine not implemented")
}
func (UnimplementedVirtualMachineRegistryServiceServer) IsAlive(context.Context, *IsAliveRequest) (*IsAliveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsAlive not implemented")
}
func (UnimplementedVirtualMachineRegistryServiceServer) GetAttachedImages(context.Context, *GetAttachedImagesRequest) (*GetAttachedImagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAttachedImages not implemented")
}
func (UnimplementedVirtualMachineRegistryServiceServer) AttachImage(context.Context, *AttachImageRequest) (*AttachImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AttachImage not implemented")
}
func (UnimplementedVirtualMachineRegistryServiceServer) DetachImage(context.Context, *DetachImageRequest) (*DetachImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DetachImage not implemented")
}
func (UnimplementedVirtualMachineRegistryServiceServer) SendACPIAction(context.Context, *SendACPIActionRequest) (*SendACPIActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendACPIAction not implemented")
}
func (UnimplementedVirtualMachineRegistryServiceServer) GetVNCServer(context.Context, *GetVNCServerRequest) (*GetVNCServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVNCServer not implemented")
}
func (UnimplementedVirtualMachineRegistryServiceServer) GetMetadata(context.Context, *GetMetadataRequest) (*GetMetadataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMetadata not implemented")
}
func (UnimplementedVirtualMachineRegistryServiceServer) SetMetadata(context.Context, *SetMetadataRequest) (*SetMetadataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetMetadata not implemented")
}
func (UnimplementedVirtualMachineRegistryServiceServer) mustEmbedUnimplementedVirtualMachineRegistryServiceServer() {
}

// UnsafeVirtualMachineRegistryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VirtualMachineRegistryServiceServer will
// result in compilation errors.
type UnsafeVirtualMachineRegistryServiceServer interface {
	mustEmbedUnimplementedVirtualMachineRegistryServiceServer()
}

func RegisterVirtualMachineRegistryServiceServer(s grpc.ServiceRegistrar, srv VirtualMachineRegistryServiceServer) {
	s.RegisterService(&VirtualMachineRegistryService_ServiceDesc, srv)
}

func _VirtualMachineRegistryService_GetVirtualMachines_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(VirtualMachineRegistryServiceServer).GetVirtualMachines(m, &virtualMachineRegistryServiceGetVirtualMachinesServer{stream})
}

type VirtualMachineRegistryService_GetVirtualMachinesServer interface {
	Send(*VirtualMachine) error
	grpc.ServerStream
}

type virtualMachineRegistryServiceGetVirtualMachinesServer struct {
	grpc.ServerStream
}

func (x *virtualMachineRegistryServiceGetVirtualMachinesServer) Send(m *VirtualMachine) error {
	return x.ServerStream.SendMsg(m)
}

func _VirtualMachineRegistryService_FindVirtualMachine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindVirtualMachineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VirtualMachineRegistryServiceServer).FindVirtualMachine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kitsune.proto.v1.VirtualMachineRegistryService/FindVirtualMachine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VirtualMachineRegistryServiceServer).FindVirtualMachine(ctx, req.(*FindVirtualMachineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VirtualMachineRegistryService_CreateVirtualMachine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateVirtualMachineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VirtualMachineRegistryServiceServer).CreateVirtualMachine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kitsune.proto.v1.VirtualMachineRegistryService/CreateVirtualMachine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VirtualMachineRegistryServiceServer).CreateVirtualMachine(ctx, req.(*CreateVirtualMachineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VirtualMachineRegistryService_DeleteVirtualMachine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteVirtualMachineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VirtualMachineRegistryServiceServer).DeleteVirtualMachine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kitsune.proto.v1.VirtualMachineRegistryService/DeleteVirtualMachine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VirtualMachineRegistryServiceServer).DeleteVirtualMachine(ctx, req.(*DeleteVirtualMachineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VirtualMachineRegistryService_IsAlive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsAliveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VirtualMachineRegistryServiceServer).IsAlive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kitsune.proto.v1.VirtualMachineRegistryService/IsAlive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VirtualMachineRegistryServiceServer).IsAlive(ctx, req.(*IsAliveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VirtualMachineRegistryService_GetAttachedImages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAttachedImagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VirtualMachineRegistryServiceServer).GetAttachedImages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kitsune.proto.v1.VirtualMachineRegistryService/GetAttachedImages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VirtualMachineRegistryServiceServer).GetAttachedImages(ctx, req.(*GetAttachedImagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VirtualMachineRegistryService_AttachImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AttachImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VirtualMachineRegistryServiceServer).AttachImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kitsune.proto.v1.VirtualMachineRegistryService/AttachImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VirtualMachineRegistryServiceServer).AttachImage(ctx, req.(*AttachImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VirtualMachineRegistryService_DetachImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DetachImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VirtualMachineRegistryServiceServer).DetachImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kitsune.proto.v1.VirtualMachineRegistryService/DetachImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VirtualMachineRegistryServiceServer).DetachImage(ctx, req.(*DetachImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VirtualMachineRegistryService_SendACPIAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendACPIActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VirtualMachineRegistryServiceServer).SendACPIAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kitsune.proto.v1.VirtualMachineRegistryService/SendACPIAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VirtualMachineRegistryServiceServer).SendACPIAction(ctx, req.(*SendACPIActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VirtualMachineRegistryService_GetVNCServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVNCServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VirtualMachineRegistryServiceServer).GetVNCServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kitsune.proto.v1.VirtualMachineRegistryService/GetVNCServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VirtualMachineRegistryServiceServer).GetVNCServer(ctx, req.(*GetVNCServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VirtualMachineRegistryService_GetMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VirtualMachineRegistryServiceServer).GetMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kitsune.proto.v1.VirtualMachineRegistryService/GetMetadata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VirtualMachineRegistryServiceServer).GetMetadata(ctx, req.(*GetMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VirtualMachineRegistryService_SetMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VirtualMachineRegistryServiceServer).SetMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kitsune.proto.v1.VirtualMachineRegistryService/SetMetadata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VirtualMachineRegistryServiceServer).SetMetadata(ctx, req.(*SetMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// VirtualMachineRegistryService_ServiceDesc is the grpc.ServiceDesc for VirtualMachineRegistryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VirtualMachineRegistryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kitsune.proto.v1.VirtualMachineRegistryService",
	HandlerType: (*VirtualMachineRegistryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindVirtualMachine",
			Handler:    _VirtualMachineRegistryService_FindVirtualMachine_Handler,
		},
		{
			MethodName: "CreateVirtualMachine",
			Handler:    _VirtualMachineRegistryService_CreateVirtualMachine_Handler,
		},
		{
			MethodName: "DeleteVirtualMachine",
			Handler:    _VirtualMachineRegistryService_DeleteVirtualMachine_Handler,
		},
		{
			MethodName: "IsAlive",
			Handler:    _VirtualMachineRegistryService_IsAlive_Handler,
		},
		{
			MethodName: "GetAttachedImages",
			Handler:    _VirtualMachineRegistryService_GetAttachedImages_Handler,
		},
		{
			MethodName: "AttachImage",
			Handler:    _VirtualMachineRegistryService_AttachImage_Handler,
		},
		{
			MethodName: "DetachImage",
			Handler:    _VirtualMachineRegistryService_DetachImage_Handler,
		},
		{
			MethodName: "SendACPIAction",
			Handler:    _VirtualMachineRegistryService_SendACPIAction_Handler,
		},
		{
			MethodName: "GetVNCServer",
			Handler:    _VirtualMachineRegistryService_GetVNCServer_Handler,
		},
		{
			MethodName: "GetMetadata",
			Handler:    _VirtualMachineRegistryService_GetMetadata_Handler,
		},
		{
			MethodName: "SetMetadata",
			Handler:    _VirtualMachineRegistryService_SetMetadata_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetVirtualMachines",
			Handler:       _VirtualMachineRegistryService_GetVirtualMachines_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "kitsune/proto/v1/machine.proto",
}
