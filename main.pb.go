// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protocol/main.proto

package protocol

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

func init() { proto.RegisterFile("protocol/main.proto", fileDescriptor_9e8e7d5cf664e777) }

var fileDescriptor_9e8e7d5cf664e777 = []byte{
	// 349 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x51, 0x4b, 0xfb, 0x30,
	0x14, 0xc5, 0xfb, 0xb2, 0xfd, 0xff, 0x84, 0x31, 0x5d, 0xe6, 0x26, 0x74, 0x32, 0xb0, 0xbe, 0x77,
	0xa0, 0x2f, 0xa2, 0x0f, 0x43, 0x37, 0x1f, 0x84, 0xc1, 0x64, 0x45, 0xd8, 0x6b, 0x57, 0xae, 0x5d,
	0x58, 0xd7, 0x8c, 0xf4, 0x76, 0xa0, 0xdf, 0xd7, 0xef, 0x21, 0x4d, 0xd2, 0x25, 0xa5, 0x54, 0x1f,
	0x7b, 0xce, 0xb9, 0xbf, 0x7b, 0x9a, 0x84, 0xf4, 0x0f, 0x82, 0x23, 0x8f, 0x78, 0x32, 0xd9, 0x87,
	0x2c, 0xf5, 0xe5, 0x17, 0xfd, 0x5f, 0x8a, 0xee, 0x28, 0xe6, 0x3c, 0x4e, 0x60, 0x22, 0x85, 0x4d,
	0xfe, 0x31, 0x81, 0xfd, 0x01, 0x3f, 0x55, 0xcc, 0xbd, 0x38, 0xcd, 0x66, 0x80, 0xf9, 0x41, 0xab,
	0x83, 0x93, 0xba, 0xc9, 0xa3, 0x1d, 0xa0, 0x96, 0xcd, 0xa2, 0x30, 0xc7, 0xad, 0x12, 0x6f, 0xbf,
	0x5b, 0xa4, 0x13, 0x20, 0x17, 0x10, 0x80, 0x38, 0xb2, 0x08, 0xe8, 0x94, 0x90, 0xd9, 0x16, 0xa2,
	0x5d, 0x50, 0x00, 0xe9, 0xd0, 0x57, 0xeb, 0xfd, 0x72, 0xbd, 0xff, 0x52, 0xac, 0x77, 0x2f, 0xfd,
	0x12, 0xe6, 0x9b, 0xf4, 0x0a, 0x32, 0xcf, 0xa1, 0xf7, 0xa4, 0xa5, 0x66, 0x07, 0x26, 0x23, 0x85,
	0xb7, 0x50, 0x84, 0xfb, 0xcc, 0x6d, 0x40, 0x7a, 0x0e, 0x7d, 0x20, 0xed, 0x80, 0xc5, 0xe9, 0x6b,
	0x4a, 0x87, 0xd6, 0xa8, 0x54, 0xfe, 0x9c, 0x7d, 0x24, 0xff, 0x8a, 0xe4, 0x32, 0xc7, 0xc6, 0xce,
	0xcd, 0xc3, 0x4b, 0xd2, 0x5d, 0xb0, 0x0c, 0x9f, 0x92, 0xe4, 0x59, 0x1e, 0x58, 0x46, 0xc7, 0xa6,
	0x40, 0xd5, 0xd1, 0x45, 0x46, 0x4d, 0xbe, 0x3a, 0x83, 0x39, 0xe9, 0xae, 0x00, 0x05, 0x83, 0x23,
	0x28, 0xdd, 0x06, 0x56, 0x1d, 0x0d, 0x3c, 0x37, 0xbe, 0xd2, 0x3d, 0x87, 0xae, 0x49, 0xaf, 0x9a,
	0x5d, 0xf0, 0x98, 0x5e, 0x57, 0x41, 0xc0, 0x2c, 0x53, 0xb3, 0xc6, 0xbf, 0x44, 0x54, 0xbf, 0x29,
	0xe9, 0xcc, 0x21, 0x01, 0x2c, 0xdb, 0x5d, 0x99, 0x09, 0x5b, 0xd7, 0xbc, 0x33, 0xe3, 0x2e, 0xf5,
	0x25, 0xaf, 0x49, 0x4f, 0xfe, 0x77, 0x8e, 0x5b, 0x2e, 0xd8, 0x57, 0x88, 0x8c, 0xa7, 0x76, 0xb5,
	0x9a, 0x59, 0xaf, 0x56, 0x8b, 0x28, 0xf2, 0x3b, 0xe9, 0xcf, 0x04, 0x84, 0x08, 0x55, 0xf6, 0x8d,
	0xf5, 0xe0, 0xea, 0xb6, 0xa6, 0x5b, 0xaf, 0xb2, 0x62, 0x7b, 0xce, 0xa6, 0x2d, 0x9d, 0xbb, 0x9f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x39, 0x93, 0xb1, 0x60, 0x6e, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StoreServiceClient is the client API for StoreService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StoreServiceClient interface {
	// Setup
	CheckSetup(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*CheckSetupRes, error)
	Setup(ctx context.Context, in *SetupParams, opts ...grpc.CallOption) (*empty.Empty, error)
	SignIn(ctx context.Context, in *SignInParams, opts ...grpc.CallOption) (*empty.Empty, error)
	SignOut(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
	// Bucket
	ListAllBuckets(ctx context.Context, in *ListAllBucketsParams, opts ...grpc.CallOption) (*ListAllBucketsRes, error)
	RetrieveBucket(ctx context.Context, in *RetrieveBucketParams, opts ...grpc.CallOption) (*Bucket, error)
	RetrieveBucketLog(ctx context.Context, in *RetreiveBucketLogParams, opts ...grpc.CallOption) (*RetreiveBucketLogRes, error)
	DeleteBucket(ctx context.Context, in *DeleteBucketParams, opts ...grpc.CallOption) (*OpRes, error)
	// Auth
	ListAuthorization(ctx context.Context, in *ListAuthorizationParams, opts ...grpc.CallOption) (*ListAuthorizationRes, error)
	CreateAuthorization(ctx context.Context, in *CreateAuthorizationParams, opts ...grpc.CallOption) (*Authorization, error)
}

type storeServiceClient struct {
	cc *grpc.ClientConn
}

func NewStoreServiceClient(cc *grpc.ClientConn) StoreServiceClient {
	return &storeServiceClient{cc}
}

func (c *storeServiceClient) CheckSetup(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*CheckSetupRes, error) {
	out := new(CheckSetupRes)
	err := c.cc.Invoke(ctx, "/protocol.StoreService/CheckSetup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeServiceClient) Setup(ctx context.Context, in *SetupParams, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/protocol.StoreService/Setup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeServiceClient) SignIn(ctx context.Context, in *SignInParams, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/protocol.StoreService/SignIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeServiceClient) SignOut(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/protocol.StoreService/SignOut", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeServiceClient) ListAllBuckets(ctx context.Context, in *ListAllBucketsParams, opts ...grpc.CallOption) (*ListAllBucketsRes, error) {
	out := new(ListAllBucketsRes)
	err := c.cc.Invoke(ctx, "/protocol.StoreService/ListAllBuckets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeServiceClient) RetrieveBucket(ctx context.Context, in *RetrieveBucketParams, opts ...grpc.CallOption) (*Bucket, error) {
	out := new(Bucket)
	err := c.cc.Invoke(ctx, "/protocol.StoreService/RetrieveBucket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeServiceClient) RetrieveBucketLog(ctx context.Context, in *RetreiveBucketLogParams, opts ...grpc.CallOption) (*RetreiveBucketLogRes, error) {
	out := new(RetreiveBucketLogRes)
	err := c.cc.Invoke(ctx, "/protocol.StoreService/RetrieveBucketLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeServiceClient) DeleteBucket(ctx context.Context, in *DeleteBucketParams, opts ...grpc.CallOption) (*OpRes, error) {
	out := new(OpRes)
	err := c.cc.Invoke(ctx, "/protocol.StoreService/DeleteBucket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeServiceClient) ListAuthorization(ctx context.Context, in *ListAuthorizationParams, opts ...grpc.CallOption) (*ListAuthorizationRes, error) {
	out := new(ListAuthorizationRes)
	err := c.cc.Invoke(ctx, "/protocol.StoreService/ListAuthorization", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeServiceClient) CreateAuthorization(ctx context.Context, in *CreateAuthorizationParams, opts ...grpc.CallOption) (*Authorization, error) {
	out := new(Authorization)
	err := c.cc.Invoke(ctx, "/protocol.StoreService/CreateAuthorization", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StoreServiceServer is the server API for StoreService service.
type StoreServiceServer interface {
	// Setup
	CheckSetup(context.Context, *empty.Empty) (*CheckSetupRes, error)
	Setup(context.Context, *SetupParams) (*empty.Empty, error)
	SignIn(context.Context, *SignInParams) (*empty.Empty, error)
	SignOut(context.Context, *empty.Empty) (*empty.Empty, error)
	// Bucket
	ListAllBuckets(context.Context, *ListAllBucketsParams) (*ListAllBucketsRes, error)
	RetrieveBucket(context.Context, *RetrieveBucketParams) (*Bucket, error)
	RetrieveBucketLog(context.Context, *RetreiveBucketLogParams) (*RetreiveBucketLogRes, error)
	DeleteBucket(context.Context, *DeleteBucketParams) (*OpRes, error)
	// Auth
	ListAuthorization(context.Context, *ListAuthorizationParams) (*ListAuthorizationRes, error)
	CreateAuthorization(context.Context, *CreateAuthorizationParams) (*Authorization, error)
}

// UnimplementedStoreServiceServer can be embedded to have forward compatible implementations.
type UnimplementedStoreServiceServer struct {
}

func (*UnimplementedStoreServiceServer) CheckSetup(ctx context.Context, req *empty.Empty) (*CheckSetupRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckSetup not implemented")
}
func (*UnimplementedStoreServiceServer) Setup(ctx context.Context, req *SetupParams) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Setup not implemented")
}
func (*UnimplementedStoreServiceServer) SignIn(ctx context.Context, req *SignInParams) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignIn not implemented")
}
func (*UnimplementedStoreServiceServer) SignOut(ctx context.Context, req *empty.Empty) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignOut not implemented")
}
func (*UnimplementedStoreServiceServer) ListAllBuckets(ctx context.Context, req *ListAllBucketsParams) (*ListAllBucketsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAllBuckets not implemented")
}
func (*UnimplementedStoreServiceServer) RetrieveBucket(ctx context.Context, req *RetrieveBucketParams) (*Bucket, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RetrieveBucket not implemented")
}
func (*UnimplementedStoreServiceServer) RetrieveBucketLog(ctx context.Context, req *RetreiveBucketLogParams) (*RetreiveBucketLogRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RetrieveBucketLog not implemented")
}
func (*UnimplementedStoreServiceServer) DeleteBucket(ctx context.Context, req *DeleteBucketParams) (*OpRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBucket not implemented")
}
func (*UnimplementedStoreServiceServer) ListAuthorization(ctx context.Context, req *ListAuthorizationParams) (*ListAuthorizationRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAuthorization not implemented")
}
func (*UnimplementedStoreServiceServer) CreateAuthorization(ctx context.Context, req *CreateAuthorizationParams) (*Authorization, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAuthorization not implemented")
}

func RegisterStoreServiceServer(s *grpc.Server, srv StoreServiceServer) {
	s.RegisterService(&_StoreService_serviceDesc, srv)
}

func _StoreService_CheckSetup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServiceServer).CheckSetup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.StoreService/CheckSetup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServiceServer).CheckSetup(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoreService_Setup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetupParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServiceServer).Setup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.StoreService/Setup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServiceServer).Setup(ctx, req.(*SetupParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoreService_SignIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignInParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServiceServer).SignIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.StoreService/SignIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServiceServer).SignIn(ctx, req.(*SignInParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoreService_SignOut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServiceServer).SignOut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.StoreService/SignOut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServiceServer).SignOut(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoreService_ListAllBuckets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAllBucketsParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServiceServer).ListAllBuckets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.StoreService/ListAllBuckets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServiceServer).ListAllBuckets(ctx, req.(*ListAllBucketsParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoreService_RetrieveBucket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RetrieveBucketParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServiceServer).RetrieveBucket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.StoreService/RetrieveBucket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServiceServer).RetrieveBucket(ctx, req.(*RetrieveBucketParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoreService_RetrieveBucketLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RetreiveBucketLogParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServiceServer).RetrieveBucketLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.StoreService/RetrieveBucketLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServiceServer).RetrieveBucketLog(ctx, req.(*RetreiveBucketLogParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoreService_DeleteBucket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBucketParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServiceServer).DeleteBucket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.StoreService/DeleteBucket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServiceServer).DeleteBucket(ctx, req.(*DeleteBucketParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoreService_ListAuthorization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAuthorizationParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServiceServer).ListAuthorization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.StoreService/ListAuthorization",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServiceServer).ListAuthorization(ctx, req.(*ListAuthorizationParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoreService_CreateAuthorization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAuthorizationParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServiceServer).CreateAuthorization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.StoreService/CreateAuthorization",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServiceServer).CreateAuthorization(ctx, req.(*CreateAuthorizationParams))
	}
	return interceptor(ctx, in, info, handler)
}

var _StoreService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protocol.StoreService",
	HandlerType: (*StoreServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckSetup",
			Handler:    _StoreService_CheckSetup_Handler,
		},
		{
			MethodName: "Setup",
			Handler:    _StoreService_Setup_Handler,
		},
		{
			MethodName: "SignIn",
			Handler:    _StoreService_SignIn_Handler,
		},
		{
			MethodName: "SignOut",
			Handler:    _StoreService_SignOut_Handler,
		},
		{
			MethodName: "ListAllBuckets",
			Handler:    _StoreService_ListAllBuckets_Handler,
		},
		{
			MethodName: "RetrieveBucket",
			Handler:    _StoreService_RetrieveBucket_Handler,
		},
		{
			MethodName: "RetrieveBucketLog",
			Handler:    _StoreService_RetrieveBucketLog_Handler,
		},
		{
			MethodName: "DeleteBucket",
			Handler:    _StoreService_DeleteBucket_Handler,
		},
		{
			MethodName: "ListAuthorization",
			Handler:    _StoreService_ListAuthorization_Handler,
		},
		{
			MethodName: "CreateAuthorization",
			Handler:    _StoreService_CreateAuthorization_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protocol/main.proto",
}
