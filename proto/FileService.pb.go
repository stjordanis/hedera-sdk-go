// Code generated by protoc-gen-go. DO NOT EDIT.
// source: FileService.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("FileService.proto", fileDescriptor_5003a56c7e235889) }

var fileDescriptor_5003a56c7e235889 = []byte{
	// 244 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xcd, 0x4a, 0x03, 0x31,
	0x14, 0x85, 0x17, 0xa2, 0x8b, 0x3b, 0x75, 0xa4, 0xb3, 0x33, 0x0f, 0x20, 0x2e, 0x24, 0x82, 0x6e,
	0x45, 0x45, 0x45, 0x70, 0xe9, 0xdf, 0x03, 0x5c, 0x93, 0x6b, 0x33, 0xd2, 0xde, 0x84, 0x24, 0x2d,
	0xf4, 0x91, 0x7d, 0x0b, 0x99, 0x24, 0x95, 0xa0, 0x0b, 0xe9, 0xb8, 0x9a, 0xc9, 0xb9, 0xe7, 0x3b,
	0x27, 0x09, 0x81, 0xe9, 0x7d, 0x3f, 0xa7, 0x67, 0xf2, 0xab, 0x5e, 0x91, 0x74, 0xde, 0x46, 0xdb,
	0xed, 0xa6, 0x8f, 0x68, 0x1e, 0x97, 0xe4, 0xd7, 0x59, 0x13, 0xed, 0x13, 0x05, 0x67, 0x39, 0x14,
	0x8f, 0x38, 0x7c, 0xf1, 0xc8, 0x01, 0x55, 0xec, 0x2d, 0xff, 0x18, 0x4d, 0xab, 0x51, 0x96, 0xce,
	0x3e, 0x77, 0xa0, 0xa9, 0x7a, 0xba, 0x0b, 0x00, 0xe5, 0x09, 0x23, 0x0d, 0x62, 0xd7, 0x65, 0x97,
	0xac, 0x38, 0x21, 0x7e, 0x6b, 0x9b, 0x9a, 0x81, 0x5e, 0x3a, 0xfd, 0x0f, 0x5a, 0xd3, 0x9c, 0x46,
	0xd2, 0x57, 0xb0, 0x8f, 0xce, 0x11, 0xeb, 0x5b, 0xcb, 0x91, 0x38, 0x6e, 0x1d, 0x70, 0x0a, 0xed,
	0x8c, 0xe2, 0xd0, 0xbd, 0x49, 0x98, 0x14, 0x77, 0xba, 0x6e, 0x71, 0x50, 0x56, 0xdf, 0xc0, 0x09,
	0x34, 0x05, 0x78, 0xe0, 0x77, 0xfb, 0x97, 0xfb, 0x12, 0x26, 0x61, 0x1d, 0x22, 0x2d, 0xee, 0xd2,
	0x19, 0xb7, 0xde, 0xde, 0x35, 0xb4, 0x99, 0x7f, 0x65, 0x3d, 0x2a, 0xe1, 0xe6, 0x18, 0x8e, 0x94,
	0x5d, 0x48, 0x43, 0x9a, 0x3c, 0x1a, 0x0c, 0x66, 0xe6, 0xd1, 0x19, 0x19, 0xea, 0x27, 0x26, 0x3f,
	0x70, 0x85, 0x6f, 0x7b, 0xe9, 0xff, 0xfc, 0x2b, 0x00, 0x00, 0xff, 0xff, 0xee, 0xb2, 0x9f, 0x39,
	0x83, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FileServiceClient is the client API for FileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FileServiceClient interface {
	CreateFile(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
	UpdateFile(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
	DeleteFile(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
	AppendContent(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
	GetFileContent(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error)
	GetFileInfo(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error)
	SystemDelete(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
	SystemUndelete(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
}

type fileServiceClient struct {
	cc *grpc.ClientConn
}

func NewFileServiceClient(cc *grpc.ClientConn) FileServiceClient {
	return &fileServiceClient{cc}
}

func (c *fileServiceClient) CreateFile(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.FileService/createFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) UpdateFile(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.FileService/updateFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) DeleteFile(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.FileService/deleteFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) AppendContent(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.FileService/appendContent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) GetFileContent(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.FileService/getFileContent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) GetFileInfo(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.FileService/getFileInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) SystemDelete(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.FileService/systemDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) SystemUndelete(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.FileService/systemUndelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FileServiceServer is the server API for FileService service.
type FileServiceServer interface {
	CreateFile(context.Context, *Transaction) (*TransactionResponse, error)
	UpdateFile(context.Context, *Transaction) (*TransactionResponse, error)
	DeleteFile(context.Context, *Transaction) (*TransactionResponse, error)
	AppendContent(context.Context, *Transaction) (*TransactionResponse, error)
	GetFileContent(context.Context, *Query) (*Response, error)
	GetFileInfo(context.Context, *Query) (*Response, error)
	SystemDelete(context.Context, *Transaction) (*TransactionResponse, error)
	SystemUndelete(context.Context, *Transaction) (*TransactionResponse, error)
}

// UnimplementedFileServiceServer can be embedded to have forward compatible implementations.
type UnimplementedFileServiceServer struct {
}

func (*UnimplementedFileServiceServer) CreateFile(ctx context.Context, req *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFile not implemented")
}
func (*UnimplementedFileServiceServer) UpdateFile(ctx context.Context, req *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFile not implemented")
}
func (*UnimplementedFileServiceServer) DeleteFile(ctx context.Context, req *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFile not implemented")
}
func (*UnimplementedFileServiceServer) AppendContent(ctx context.Context, req *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppendContent not implemented")
}
func (*UnimplementedFileServiceServer) GetFileContent(ctx context.Context, req *Query) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFileContent not implemented")
}
func (*UnimplementedFileServiceServer) GetFileInfo(ctx context.Context, req *Query) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFileInfo not implemented")
}
func (*UnimplementedFileServiceServer) SystemDelete(ctx context.Context, req *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemDelete not implemented")
}
func (*UnimplementedFileServiceServer) SystemUndelete(ctx context.Context, req *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemUndelete not implemented")
}

func RegisterFileServiceServer(s *grpc.Server, srv FileServiceServer) {
	s.RegisterService(&_FileService_serviceDesc, srv)
}

func _FileService_CreateFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).CreateFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.FileService/CreateFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).CreateFile(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_UpdateFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).UpdateFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.FileService/UpdateFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).UpdateFile(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_DeleteFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).DeleteFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.FileService/DeleteFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).DeleteFile(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_AppendContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).AppendContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.FileService/AppendContent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).AppendContent(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_GetFileContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).GetFileContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.FileService/GetFileContent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).GetFileContent(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_GetFileInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).GetFileInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.FileService/GetFileInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).GetFileInfo(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_SystemDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).SystemDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.FileService/SystemDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).SystemDelete(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_SystemUndelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).SystemUndelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.FileService/SystemUndelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).SystemUndelete(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

var _FileService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.FileService",
	HandlerType: (*FileServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createFile",
			Handler:    _FileService_CreateFile_Handler,
		},
		{
			MethodName: "updateFile",
			Handler:    _FileService_UpdateFile_Handler,
		},
		{
			MethodName: "deleteFile",
			Handler:    _FileService_DeleteFile_Handler,
		},
		{
			MethodName: "appendContent",
			Handler:    _FileService_AppendContent_Handler,
		},
		{
			MethodName: "getFileContent",
			Handler:    _FileService_GetFileContent_Handler,
		},
		{
			MethodName: "getFileInfo",
			Handler:    _FileService_GetFileInfo_Handler,
		},
		{
			MethodName: "systemDelete",
			Handler:    _FileService_SystemDelete_Handler,
		},
		{
			MethodName: "systemUndelete",
			Handler:    _FileService_SystemUndelete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "FileService.proto",
}
