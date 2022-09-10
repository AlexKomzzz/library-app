// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: library.proto

package api

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

// LibraryClient is the client API for Library service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LibraryClient interface {
	SeachAuthor(ctx context.Context, in *Book, opts ...grpc.CallOption) (*Authors, error)
	SeachBook(ctx context.Context, in *Author, opts ...grpc.CallOption) (*Books, error)
}

type libraryClient struct {
	cc grpc.ClientConnInterface
}

func NewLibraryClient(cc grpc.ClientConnInterface) LibraryClient {
	return &libraryClient{cc}
}

func (c *libraryClient) SeachAuthor(ctx context.Context, in *Book, opts ...grpc.CallOption) (*Authors, error) {
	out := new(Authors)
	err := c.cc.Invoke(ctx, "/api.Library/SeachAuthor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryClient) SeachBook(ctx context.Context, in *Author, opts ...grpc.CallOption) (*Books, error) {
	out := new(Books)
	err := c.cc.Invoke(ctx, "/api.Library/SeachBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LibraryServer is the server API for Library service.
// All implementations must embed UnimplementedLibraryServer
// for forward compatibility
type LibraryServer interface {
	SeachAuthor(context.Context, *Book) (*Authors, error)
	SeachBook(context.Context, *Author) (*Books, error)
	mustEmbedUnimplementedLibraryServer()
}

// UnimplementedLibraryServer must be embedded to have forward compatible implementations.
type UnimplementedLibraryServer struct {
}

func (UnimplementedLibraryServer) SeachAuthor(context.Context, *Book) (*Authors, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SeachAuthor not implemented")
}
func (UnimplementedLibraryServer) SeachBook(context.Context, *Author) (*Books, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SeachBook not implemented")
}
func (UnimplementedLibraryServer) mustEmbedUnimplementedLibraryServer() {}

// UnsafeLibraryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LibraryServer will
// result in compilation errors.
type UnsafeLibraryServer interface {
	mustEmbedUnimplementedLibraryServer()
}

func RegisterLibraryServer(s grpc.ServiceRegistrar, srv LibraryServer) {
	s.RegisterService(&Library_ServiceDesc, srv)
}

func _Library_SeachAuthor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Book)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServer).SeachAuthor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Library/SeachAuthor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServer).SeachAuthor(ctx, req.(*Book))
	}
	return interceptor(ctx, in, info, handler)
}

func _Library_SeachBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Author)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServer).SeachBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Library/SeachBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServer).SeachBook(ctx, req.(*Author))
	}
	return interceptor(ctx, in, info, handler)
}

// Library_ServiceDesc is the grpc.ServiceDesc for Library service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Library_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.Library",
	HandlerType: (*LibraryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SeachAuthor",
			Handler:    _Library_SeachAuthor_Handler,
		},
		{
			MethodName: "SeachBook",
			Handler:    _Library_SeachBook_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "library.proto",
}