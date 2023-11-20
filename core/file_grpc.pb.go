// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.2
// source: file.proto

package core

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

// FileServiceClient is the client API for FileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FileServiceClient interface {
	SendPartialFile(ctx context.Context, in *PartialFileRequest, opts ...grpc.CallOption) (*PartialFileResponse, error)
	DownloadPartialFile(ctx context.Context, in *DownloadRequest, opts ...grpc.CallOption) (FileService_DownloadPartialFileClient, error)
}

type fileServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFileServiceClient(cc grpc.ClientConnInterface) FileServiceClient {
	return &fileServiceClient{cc}
}

func (c *fileServiceClient) SendPartialFile(ctx context.Context, in *PartialFileRequest, opts ...grpc.CallOption) (*PartialFileResponse, error) {
	out := new(PartialFileResponse)
	err := c.cc.Invoke(ctx, "/core.FileService/SendPartialFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) DownloadPartialFile(ctx context.Context, in *DownloadRequest, opts ...grpc.CallOption) (FileService_DownloadPartialFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &FileService_ServiceDesc.Streams[0], "/core.FileService/DownloadPartialFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileServiceDownloadPartialFileClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FileService_DownloadPartialFileClient interface {
	Recv() (*DownloadResponse, error)
	grpc.ClientStream
}

type fileServiceDownloadPartialFileClient struct {
	grpc.ClientStream
}

func (x *fileServiceDownloadPartialFileClient) Recv() (*DownloadResponse, error) {
	m := new(DownloadResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FileServiceServer is the server API for FileService service.
// All implementations must embed UnimplementedFileServiceServer
// for forward compatibility
type FileServiceServer interface {
	SendPartialFile(context.Context, *PartialFileRequest) (*PartialFileResponse, error)
	DownloadPartialFile(*DownloadRequest, FileService_DownloadPartialFileServer) error
	mustEmbedUnimplementedFileServiceServer()
}

// UnimplementedFileServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFileServiceServer struct {
}

func (UnimplementedFileServiceServer) SendPartialFile(context.Context, *PartialFileRequest) (*PartialFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendPartialFile not implemented")
}
func (UnimplementedFileServiceServer) DownloadPartialFile(*DownloadRequest, FileService_DownloadPartialFileServer) error {
	return status.Errorf(codes.Unimplemented, "method DownloadPartialFile not implemented")
}
func (UnimplementedFileServiceServer) mustEmbedUnimplementedFileServiceServer() {}

// UnsafeFileServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FileServiceServer will
// result in compilation errors.
type UnsafeFileServiceServer interface {
	mustEmbedUnimplementedFileServiceServer()
}

func RegisterFileServiceServer(s grpc.ServiceRegistrar, srv FileServiceServer) {
	s.RegisterService(&FileService_ServiceDesc, srv)
}

func _FileService_SendPartialFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PartialFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).SendPartialFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.FileService/SendPartialFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).SendPartialFile(ctx, req.(*PartialFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_DownloadPartialFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DownloadRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FileServiceServer).DownloadPartialFile(m, &fileServiceDownloadPartialFileServer{stream})
}

type FileService_DownloadPartialFileServer interface {
	Send(*DownloadResponse) error
	grpc.ServerStream
}

type fileServiceDownloadPartialFileServer struct {
	grpc.ServerStream
}

func (x *fileServiceDownloadPartialFileServer) Send(m *DownloadResponse) error {
	return x.ServerStream.SendMsg(m)
}

// FileService_ServiceDesc is the grpc.ServiceDesc for FileService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FileService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "core.FileService",
	HandlerType: (*FileServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendPartialFile",
			Handler:    _FileService_SendPartialFile_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DownloadPartialFile",
			Handler:       _FileService_DownloadPartialFile_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "file.proto",
}
