// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.4
// source: kitsune/proto/v1/image.proto

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

// ImageRegistryServiceClient is the client API for ImageRegistryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ImageRegistryServiceClient interface {
	GetImages(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (ImageRegistryService_GetImagesClient, error)
	FindImage(ctx context.Context, in *FindImageRequest, opts ...grpc.CallOption) (*FindImageResponse, error)
	CreateImage(ctx context.Context, in *CreateImageRequest, opts ...grpc.CallOption) (*CreateImageResponse, error)
	DeleteImage(ctx context.Context, in *DeleteImageRequest, opts ...grpc.CallOption) (*DeleteImageResponse, error)
	GetMetadata(ctx context.Context, in *GetMetadataRequest, opts ...grpc.CallOption) (*GetMetadataResponse, error)
	SetMetadata(ctx context.Context, in *SetMetadataRequest, opts ...grpc.CallOption) (*SetMetadataResponse, error)
}

type imageRegistryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewImageRegistryServiceClient(cc grpc.ClientConnInterface) ImageRegistryServiceClient {
	return &imageRegistryServiceClient{cc}
}

func (c *imageRegistryServiceClient) GetImages(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (ImageRegistryService_GetImagesClient, error) {
	stream, err := c.cc.NewStream(ctx, &ImageRegistryService_ServiceDesc.Streams[0], "/kitsune.proto.v1.ImageRegistryService/GetImages", opts...)
	if err != nil {
		return nil, err
	}
	x := &imageRegistryServiceGetImagesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ImageRegistryService_GetImagesClient interface {
	Recv() (*Image, error)
	grpc.ClientStream
}

type imageRegistryServiceGetImagesClient struct {
	grpc.ClientStream
}

func (x *imageRegistryServiceGetImagesClient) Recv() (*Image, error) {
	m := new(Image)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *imageRegistryServiceClient) FindImage(ctx context.Context, in *FindImageRequest, opts ...grpc.CallOption) (*FindImageResponse, error) {
	out := new(FindImageResponse)
	err := c.cc.Invoke(ctx, "/kitsune.proto.v1.ImageRegistryService/FindImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageRegistryServiceClient) CreateImage(ctx context.Context, in *CreateImageRequest, opts ...grpc.CallOption) (*CreateImageResponse, error) {
	out := new(CreateImageResponse)
	err := c.cc.Invoke(ctx, "/kitsune.proto.v1.ImageRegistryService/CreateImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageRegistryServiceClient) DeleteImage(ctx context.Context, in *DeleteImageRequest, opts ...grpc.CallOption) (*DeleteImageResponse, error) {
	out := new(DeleteImageResponse)
	err := c.cc.Invoke(ctx, "/kitsune.proto.v1.ImageRegistryService/DeleteImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageRegistryServiceClient) GetMetadata(ctx context.Context, in *GetMetadataRequest, opts ...grpc.CallOption) (*GetMetadataResponse, error) {
	out := new(GetMetadataResponse)
	err := c.cc.Invoke(ctx, "/kitsune.proto.v1.ImageRegistryService/GetMetadata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageRegistryServiceClient) SetMetadata(ctx context.Context, in *SetMetadataRequest, opts ...grpc.CallOption) (*SetMetadataResponse, error) {
	out := new(SetMetadataResponse)
	err := c.cc.Invoke(ctx, "/kitsune.proto.v1.ImageRegistryService/SetMetadata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ImageRegistryServiceServer is the server API for ImageRegistryService service.
// All implementations must embed UnimplementedImageRegistryServiceServer
// for forward compatibility
type ImageRegistryServiceServer interface {
	GetImages(*emptypb.Empty, ImageRegistryService_GetImagesServer) error
	FindImage(context.Context, *FindImageRequest) (*FindImageResponse, error)
	CreateImage(context.Context, *CreateImageRequest) (*CreateImageResponse, error)
	DeleteImage(context.Context, *DeleteImageRequest) (*DeleteImageResponse, error)
	GetMetadata(context.Context, *GetMetadataRequest) (*GetMetadataResponse, error)
	SetMetadata(context.Context, *SetMetadataRequest) (*SetMetadataResponse, error)
	mustEmbedUnimplementedImageRegistryServiceServer()
}

// UnimplementedImageRegistryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedImageRegistryServiceServer struct {
}

func (UnimplementedImageRegistryServiceServer) GetImages(*emptypb.Empty, ImageRegistryService_GetImagesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetImages not implemented")
}
func (UnimplementedImageRegistryServiceServer) FindImage(context.Context, *FindImageRequest) (*FindImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindImage not implemented")
}
func (UnimplementedImageRegistryServiceServer) CreateImage(context.Context, *CreateImageRequest) (*CreateImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateImage not implemented")
}
func (UnimplementedImageRegistryServiceServer) DeleteImage(context.Context, *DeleteImageRequest) (*DeleteImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteImage not implemented")
}
func (UnimplementedImageRegistryServiceServer) GetMetadata(context.Context, *GetMetadataRequest) (*GetMetadataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMetadata not implemented")
}
func (UnimplementedImageRegistryServiceServer) SetMetadata(context.Context, *SetMetadataRequest) (*SetMetadataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetMetadata not implemented")
}
func (UnimplementedImageRegistryServiceServer) mustEmbedUnimplementedImageRegistryServiceServer() {}

// UnsafeImageRegistryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ImageRegistryServiceServer will
// result in compilation errors.
type UnsafeImageRegistryServiceServer interface {
	mustEmbedUnimplementedImageRegistryServiceServer()
}

func RegisterImageRegistryServiceServer(s grpc.ServiceRegistrar, srv ImageRegistryServiceServer) {
	s.RegisterService(&ImageRegistryService_ServiceDesc, srv)
}

func _ImageRegistryService_GetImages_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ImageRegistryServiceServer).GetImages(m, &imageRegistryServiceGetImagesServer{stream})
}

type ImageRegistryService_GetImagesServer interface {
	Send(*Image) error
	grpc.ServerStream
}

type imageRegistryServiceGetImagesServer struct {
	grpc.ServerStream
}

func (x *imageRegistryServiceGetImagesServer) Send(m *Image) error {
	return x.ServerStream.SendMsg(m)
}

func _ImageRegistryService_FindImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageRegistryServiceServer).FindImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kitsune.proto.v1.ImageRegistryService/FindImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageRegistryServiceServer).FindImage(ctx, req.(*FindImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImageRegistryService_CreateImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageRegistryServiceServer).CreateImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kitsune.proto.v1.ImageRegistryService/CreateImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageRegistryServiceServer).CreateImage(ctx, req.(*CreateImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImageRegistryService_DeleteImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageRegistryServiceServer).DeleteImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kitsune.proto.v1.ImageRegistryService/DeleteImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageRegistryServiceServer).DeleteImage(ctx, req.(*DeleteImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImageRegistryService_GetMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageRegistryServiceServer).GetMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kitsune.proto.v1.ImageRegistryService/GetMetadata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageRegistryServiceServer).GetMetadata(ctx, req.(*GetMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImageRegistryService_SetMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageRegistryServiceServer).SetMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kitsune.proto.v1.ImageRegistryService/SetMetadata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageRegistryServiceServer).SetMetadata(ctx, req.(*SetMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ImageRegistryService_ServiceDesc is the grpc.ServiceDesc for ImageRegistryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ImageRegistryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kitsune.proto.v1.ImageRegistryService",
	HandlerType: (*ImageRegistryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindImage",
			Handler:    _ImageRegistryService_FindImage_Handler,
		},
		{
			MethodName: "CreateImage",
			Handler:    _ImageRegistryService_CreateImage_Handler,
		},
		{
			MethodName: "DeleteImage",
			Handler:    _ImageRegistryService_DeleteImage_Handler,
		},
		{
			MethodName: "GetMetadata",
			Handler:    _ImageRegistryService_GetMetadata_Handler,
		},
		{
			MethodName: "SetMetadata",
			Handler:    _ImageRegistryService_SetMetadata_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetImages",
			Handler:       _ImageRegistryService_GetImages_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "kitsune/proto/v1/image.proto",
}
