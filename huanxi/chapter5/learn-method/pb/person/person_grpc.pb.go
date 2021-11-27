// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package person

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

// SearchServiceClient is the client API for SearchService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SearchServiceClient interface {
	//  四种方法
	Search(ctx context.Context, in *PersonReq, opts ...grpc.CallOption) (*PersonRes, error)
	StreamSearch(ctx context.Context, opts ...grpc.CallOption) (SearchService_StreamSearchClient, error)
	SearchStream(ctx context.Context, in *PersonReq, opts ...grpc.CallOption) (SearchService_SearchStreamClient, error)
	StreamSearchStream(ctx context.Context, opts ...grpc.CallOption) (SearchService_StreamSearchStreamClient, error)
}

type searchServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSearchServiceClient(cc grpc.ClientConnInterface) SearchServiceClient {
	return &searchServiceClient{cc}
}

func (c *searchServiceClient) Search(ctx context.Context, in *PersonReq, opts ...grpc.CallOption) (*PersonRes, error) {
	out := new(PersonRes)
	err := c.cc.Invoke(ctx, "/person.SearchService/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchServiceClient) StreamSearch(ctx context.Context, opts ...grpc.CallOption) (SearchService_StreamSearchClient, error) {
	stream, err := c.cc.NewStream(ctx, &SearchService_ServiceDesc.Streams[0], "/person.SearchService/StreamSearch", opts...)
	if err != nil {
		return nil, err
	}
	x := &searchServiceStreamSearchClient{stream}
	return x, nil
}

type SearchService_StreamSearchClient interface {
	Send(*PersonReq) error
	CloseAndRecv() (*PersonRes, error)
	grpc.ClientStream
}

type searchServiceStreamSearchClient struct {
	grpc.ClientStream
}

func (x *searchServiceStreamSearchClient) Send(m *PersonReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *searchServiceStreamSearchClient) CloseAndRecv() (*PersonRes, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(PersonRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *searchServiceClient) SearchStream(ctx context.Context, in *PersonReq, opts ...grpc.CallOption) (SearchService_SearchStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &SearchService_ServiceDesc.Streams[1], "/person.SearchService/SearchStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &searchServiceSearchStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SearchService_SearchStreamClient interface {
	Recv() (*PersonRes, error)
	grpc.ClientStream
}

type searchServiceSearchStreamClient struct {
	grpc.ClientStream
}

func (x *searchServiceSearchStreamClient) Recv() (*PersonRes, error) {
	m := new(PersonRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *searchServiceClient) StreamSearchStream(ctx context.Context, opts ...grpc.CallOption) (SearchService_StreamSearchStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &SearchService_ServiceDesc.Streams[2], "/person.SearchService/StreamSearchStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &searchServiceStreamSearchStreamClient{stream}
	return x, nil
}

type SearchService_StreamSearchStreamClient interface {
	Send(*PersonReq) error
	Recv() (*PersonRes, error)
	grpc.ClientStream
}

type searchServiceStreamSearchStreamClient struct {
	grpc.ClientStream
}

func (x *searchServiceStreamSearchStreamClient) Send(m *PersonReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *searchServiceStreamSearchStreamClient) Recv() (*PersonRes, error) {
	m := new(PersonRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SearchServiceServer is the server API for SearchService service.
// All implementations must embed UnimplementedSearchServiceServer
// for forward compatibility
type SearchServiceServer interface {
	//  四种方法
	Search(context.Context, *PersonReq) (*PersonRes, error)
	StreamSearch(SearchService_StreamSearchServer) error
	SearchStream(*PersonReq, SearchService_SearchStreamServer) error
	StreamSearchStream(SearchService_StreamSearchStreamServer) error
	mustEmbedUnimplementedSearchServiceServer()
}

// UnimplementedSearchServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSearchServiceServer struct {
}

func (UnimplementedSearchServiceServer) Search(context.Context, *PersonReq) (*PersonRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedSearchServiceServer) StreamSearch(SearchService_StreamSearchServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamSearch not implemented")
}
func (UnimplementedSearchServiceServer) SearchStream(*PersonReq, SearchService_SearchStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method SearchStream not implemented")
}
func (UnimplementedSearchServiceServer) StreamSearchStream(SearchService_StreamSearchStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamSearchStream not implemented")
}
func (UnimplementedSearchServiceServer) mustEmbedUnimplementedSearchServiceServer() {}

// UnsafeSearchServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SearchServiceServer will
// result in compilation errors.
type UnsafeSearchServiceServer interface {
	mustEmbedUnimplementedSearchServiceServer()
}

func RegisterSearchServiceServer(s grpc.ServiceRegistrar, srv SearchServiceServer) {
	s.RegisterService(&SearchService_ServiceDesc, srv)
}

func _SearchService_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PersonReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/person.SearchService/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).Search(ctx, req.(*PersonReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchService_StreamSearch_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SearchServiceServer).StreamSearch(&searchServiceStreamSearchServer{stream})
}

type SearchService_StreamSearchServer interface {
	SendAndClose(*PersonRes) error
	Recv() (*PersonReq, error)
	grpc.ServerStream
}

type searchServiceStreamSearchServer struct {
	grpc.ServerStream
}

func (x *searchServiceStreamSearchServer) SendAndClose(m *PersonRes) error {
	return x.ServerStream.SendMsg(m)
}

func (x *searchServiceStreamSearchServer) Recv() (*PersonReq, error) {
	m := new(PersonReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SearchService_SearchStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PersonReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SearchServiceServer).SearchStream(m, &searchServiceSearchStreamServer{stream})
}

type SearchService_SearchStreamServer interface {
	Send(*PersonRes) error
	grpc.ServerStream
}

type searchServiceSearchStreamServer struct {
	grpc.ServerStream
}

func (x *searchServiceSearchStreamServer) Send(m *PersonRes) error {
	return x.ServerStream.SendMsg(m)
}

func _SearchService_StreamSearchStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SearchServiceServer).StreamSearchStream(&searchServiceStreamSearchStreamServer{stream})
}

type SearchService_StreamSearchStreamServer interface {
	Send(*PersonRes) error
	Recv() (*PersonReq, error)
	grpc.ServerStream
}

type searchServiceStreamSearchStreamServer struct {
	grpc.ServerStream
}

func (x *searchServiceStreamSearchStreamServer) Send(m *PersonRes) error {
	return x.ServerStream.SendMsg(m)
}

func (x *searchServiceStreamSearchStreamServer) Recv() (*PersonReq, error) {
	m := new(PersonReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SearchService_ServiceDesc is the grpc.ServiceDesc for SearchService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SearchService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "person.SearchService",
	HandlerType: (*SearchServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Search",
			Handler:    _SearchService_Search_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamSearch",
			Handler:       _SearchService_StreamSearch_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "SearchStream",
			Handler:       _SearchService_SearchStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StreamSearchStream",
			Handler:       _SearchService_StreamSearchStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "person/person.proto",
}
