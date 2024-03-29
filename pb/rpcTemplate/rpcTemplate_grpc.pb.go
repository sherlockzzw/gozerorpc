// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: rpcTemplate.proto

package rpcTemplate

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

const (
	UserDetailService_UserDetailList_FullMethodName   = "/rpcTemplate.UserDetailService/UserDetailList"
	UserDetailService_UserDetailCount_FullMethodName  = "/rpcTemplate.UserDetailService/UserDetailCount"
	UserDetailService_UserDetailUpdate_FullMethodName = "/rpcTemplate.UserDetailService/UserDetailUpdate"
	UserDetailService_UserDetailInsert_FullMethodName = "/rpcTemplate.UserDetailService/UserDetailInsert"
	UserDetailService_UserDetailDelete_FullMethodName = "/rpcTemplate.UserDetailService/UserDetailDelete"
)

// UserDetailServiceClient is the client API for UserDetailService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserDetailServiceClient interface {
	UserDetailList(ctx context.Context, in *UserDetailListRequest, opts ...grpc.CallOption) (*UserDetailListResponse, error)
	UserDetailCount(ctx context.Context, in *UserDetailCountReq, opts ...grpc.CallOption) (*UserDetailCountResp, error)
	UserDetailUpdate(ctx context.Context, in *UserDetailUpdateRequest, opts ...grpc.CallOption) (*UserDetailUpdateResponse, error)
	UserDetailInsert(ctx context.Context, in *UserDetailInsertReq, opts ...grpc.CallOption) (*UserDetailInsertResp, error)
	UserDetailDelete(ctx context.Context, in *UserDetailDeleteReq, opts ...grpc.CallOption) (*UserDetailDeleteResp, error)
}

type userDetailServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserDetailServiceClient(cc grpc.ClientConnInterface) UserDetailServiceClient {
	return &userDetailServiceClient{cc}
}

func (c *userDetailServiceClient) UserDetailList(ctx context.Context, in *UserDetailListRequest, opts ...grpc.CallOption) (*UserDetailListResponse, error) {
	out := new(UserDetailListResponse)
	err := c.cc.Invoke(ctx, UserDetailService_UserDetailList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userDetailServiceClient) UserDetailCount(ctx context.Context, in *UserDetailCountReq, opts ...grpc.CallOption) (*UserDetailCountResp, error) {
	out := new(UserDetailCountResp)
	err := c.cc.Invoke(ctx, UserDetailService_UserDetailCount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userDetailServiceClient) UserDetailUpdate(ctx context.Context, in *UserDetailUpdateRequest, opts ...grpc.CallOption) (*UserDetailUpdateResponse, error) {
	out := new(UserDetailUpdateResponse)
	err := c.cc.Invoke(ctx, UserDetailService_UserDetailUpdate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userDetailServiceClient) UserDetailInsert(ctx context.Context, in *UserDetailInsertReq, opts ...grpc.CallOption) (*UserDetailInsertResp, error) {
	out := new(UserDetailInsertResp)
	err := c.cc.Invoke(ctx, UserDetailService_UserDetailInsert_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userDetailServiceClient) UserDetailDelete(ctx context.Context, in *UserDetailDeleteReq, opts ...grpc.CallOption) (*UserDetailDeleteResp, error) {
	out := new(UserDetailDeleteResp)
	err := c.cc.Invoke(ctx, UserDetailService_UserDetailDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserDetailServiceServer is the server API for UserDetailService service.
// All implementations must embed UnimplementedUserDetailServiceServer
// for forward compatibility
type UserDetailServiceServer interface {
	UserDetailList(context.Context, *UserDetailListRequest) (*UserDetailListResponse, error)
	UserDetailCount(context.Context, *UserDetailCountReq) (*UserDetailCountResp, error)
	UserDetailUpdate(context.Context, *UserDetailUpdateRequest) (*UserDetailUpdateResponse, error)
	UserDetailInsert(context.Context, *UserDetailInsertReq) (*UserDetailInsertResp, error)
	UserDetailDelete(context.Context, *UserDetailDeleteReq) (*UserDetailDeleteResp, error)
	mustEmbedUnimplementedUserDetailServiceServer()
}

// UnimplementedUserDetailServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserDetailServiceServer struct {
}

func (UnimplementedUserDetailServiceServer) UserDetailList(context.Context, *UserDetailListRequest) (*UserDetailListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserDetailList not implemented")
}
func (UnimplementedUserDetailServiceServer) UserDetailCount(context.Context, *UserDetailCountReq) (*UserDetailCountResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserDetailCount not implemented")
}
func (UnimplementedUserDetailServiceServer) UserDetailUpdate(context.Context, *UserDetailUpdateRequest) (*UserDetailUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserDetailUpdate not implemented")
}
func (UnimplementedUserDetailServiceServer) UserDetailInsert(context.Context, *UserDetailInsertReq) (*UserDetailInsertResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserDetailInsert not implemented")
}
func (UnimplementedUserDetailServiceServer) UserDetailDelete(context.Context, *UserDetailDeleteReq) (*UserDetailDeleteResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserDetailDelete not implemented")
}
func (UnimplementedUserDetailServiceServer) mustEmbedUnimplementedUserDetailServiceServer() {}

// UnsafeUserDetailServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserDetailServiceServer will
// result in compilation errors.
type UnsafeUserDetailServiceServer interface {
	mustEmbedUnimplementedUserDetailServiceServer()
}

func RegisterUserDetailServiceServer(s grpc.ServiceRegistrar, srv UserDetailServiceServer) {
	s.RegisterService(&UserDetailService_ServiceDesc, srv)
}

func _UserDetailService_UserDetailList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserDetailListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserDetailServiceServer).UserDetailList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserDetailService_UserDetailList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserDetailServiceServer).UserDetailList(ctx, req.(*UserDetailListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserDetailService_UserDetailCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserDetailCountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserDetailServiceServer).UserDetailCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserDetailService_UserDetailCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserDetailServiceServer).UserDetailCount(ctx, req.(*UserDetailCountReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserDetailService_UserDetailUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserDetailUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserDetailServiceServer).UserDetailUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserDetailService_UserDetailUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserDetailServiceServer).UserDetailUpdate(ctx, req.(*UserDetailUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserDetailService_UserDetailInsert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserDetailInsertReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserDetailServiceServer).UserDetailInsert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserDetailService_UserDetailInsert_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserDetailServiceServer).UserDetailInsert(ctx, req.(*UserDetailInsertReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserDetailService_UserDetailDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserDetailDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserDetailServiceServer).UserDetailDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserDetailService_UserDetailDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserDetailServiceServer).UserDetailDelete(ctx, req.(*UserDetailDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

// UserDetailService_ServiceDesc is the grpc.ServiceDesc for UserDetailService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserDetailService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rpcTemplate.UserDetailService",
	HandlerType: (*UserDetailServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserDetailList",
			Handler:    _UserDetailService_UserDetailList_Handler,
		},
		{
			MethodName: "UserDetailCount",
			Handler:    _UserDetailService_UserDetailCount_Handler,
		},
		{
			MethodName: "UserDetailUpdate",
			Handler:    _UserDetailService_UserDetailUpdate_Handler,
		},
		{
			MethodName: "UserDetailInsert",
			Handler:    _UserDetailService_UserDetailInsert_Handler,
		},
		{
			MethodName: "UserDetailDelete",
			Handler:    _UserDetailService_UserDetailDelete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpcTemplate.proto",
}

const (
	TestUserService_GetUserInfo_FullMethodName = "/rpcTemplate.TestUserService/GetUserInfo"
	TestUserService_AddUser_FullMethodName     = "/rpcTemplate.TestUserService/AddUser"
	TestUserService_GetUserList_FullMethodName = "/rpcTemplate.TestUserService/GetUserList"
	TestUserService_UpdUser_FullMethodName     = "/rpcTemplate.TestUserService/UpdUser"
	TestUserService_DelUser_FullMethodName     = "/rpcTemplate.TestUserService/DelUser"
)

// TestUserServiceClient is the client API for TestUserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TestUserServiceClient interface {
	GetUserInfo(ctx context.Context, in *GetUserRes, opts ...grpc.CallOption) (*GetUserResp, error)
	AddUser(ctx context.Context, in *AddUserRes, opts ...grpc.CallOption) (*AddUserResp, error)
	GetUserList(ctx context.Context, in *GetUserListReq, opts ...grpc.CallOption) (*GetUserListResp, error)
	UpdUser(ctx context.Context, in *UpdUserReq, opts ...grpc.CallOption) (*UpdUserResp, error)
	DelUser(ctx context.Context, in *DelUserReq, opts ...grpc.CallOption) (*DelUserResp, error)
}

type testUserServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTestUserServiceClient(cc grpc.ClientConnInterface) TestUserServiceClient {
	return &testUserServiceClient{cc}
}

func (c *testUserServiceClient) GetUserInfo(ctx context.Context, in *GetUserRes, opts ...grpc.CallOption) (*GetUserResp, error) {
	out := new(GetUserResp)
	err := c.cc.Invoke(ctx, TestUserService_GetUserInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testUserServiceClient) AddUser(ctx context.Context, in *AddUserRes, opts ...grpc.CallOption) (*AddUserResp, error) {
	out := new(AddUserResp)
	err := c.cc.Invoke(ctx, TestUserService_AddUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testUserServiceClient) GetUserList(ctx context.Context, in *GetUserListReq, opts ...grpc.CallOption) (*GetUserListResp, error) {
	out := new(GetUserListResp)
	err := c.cc.Invoke(ctx, TestUserService_GetUserList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testUserServiceClient) UpdUser(ctx context.Context, in *UpdUserReq, opts ...grpc.CallOption) (*UpdUserResp, error) {
	out := new(UpdUserResp)
	err := c.cc.Invoke(ctx, TestUserService_UpdUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testUserServiceClient) DelUser(ctx context.Context, in *DelUserReq, opts ...grpc.CallOption) (*DelUserResp, error) {
	out := new(DelUserResp)
	err := c.cc.Invoke(ctx, TestUserService_DelUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestUserServiceServer is the server API for TestUserService service.
// All implementations must embed UnimplementedTestUserServiceServer
// for forward compatibility
type TestUserServiceServer interface {
	GetUserInfo(context.Context, *GetUserRes) (*GetUserResp, error)
	AddUser(context.Context, *AddUserRes) (*AddUserResp, error)
	GetUserList(context.Context, *GetUserListReq) (*GetUserListResp, error)
	UpdUser(context.Context, *UpdUserReq) (*UpdUserResp, error)
	DelUser(context.Context, *DelUserReq) (*DelUserResp, error)
	mustEmbedUnimplementedTestUserServiceServer()
}

// UnimplementedTestUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTestUserServiceServer struct {
}

func (UnimplementedTestUserServiceServer) GetUserInfo(context.Context, *GetUserRes) (*GetUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfo not implemented")
}
func (UnimplementedTestUserServiceServer) AddUser(context.Context, *AddUserRes) (*AddUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUser not implemented")
}
func (UnimplementedTestUserServiceServer) GetUserList(context.Context, *GetUserListReq) (*GetUserListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserList not implemented")
}
func (UnimplementedTestUserServiceServer) UpdUser(context.Context, *UpdUserReq) (*UpdUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdUser not implemented")
}
func (UnimplementedTestUserServiceServer) DelUser(context.Context, *DelUserReq) (*DelUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelUser not implemented")
}
func (UnimplementedTestUserServiceServer) mustEmbedUnimplementedTestUserServiceServer() {}

// UnsafeTestUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TestUserServiceServer will
// result in compilation errors.
type UnsafeTestUserServiceServer interface {
	mustEmbedUnimplementedTestUserServiceServer()
}

func RegisterTestUserServiceServer(s grpc.ServiceRegistrar, srv TestUserServiceServer) {
	s.RegisterService(&TestUserService_ServiceDesc, srv)
}

func _TestUserService_GetUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestUserServiceServer).GetUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TestUserService_GetUserInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestUserServiceServer).GetUserInfo(ctx, req.(*GetUserRes))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestUserService_AddUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserRes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestUserServiceServer).AddUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TestUserService_AddUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestUserServiceServer).AddUser(ctx, req.(*AddUserRes))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestUserService_GetUserList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestUserServiceServer).GetUserList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TestUserService_GetUserList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestUserServiceServer).GetUserList(ctx, req.(*GetUserListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestUserService_UpdUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestUserServiceServer).UpdUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TestUserService_UpdUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestUserServiceServer).UpdUser(ctx, req.(*UpdUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestUserService_DelUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestUserServiceServer).DelUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TestUserService_DelUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestUserServiceServer).DelUser(ctx, req.(*DelUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

// TestUserService_ServiceDesc is the grpc.ServiceDesc for TestUserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TestUserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rpcTemplate.TestUserService",
	HandlerType: (*TestUserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserInfo",
			Handler:    _TestUserService_GetUserInfo_Handler,
		},
		{
			MethodName: "AddUser",
			Handler:    _TestUserService_AddUser_Handler,
		},
		{
			MethodName: "GetUserList",
			Handler:    _TestUserService_GetUserList_Handler,
		},
		{
			MethodName: "UpdUser",
			Handler:    _TestUserService_UpdUser_Handler,
		},
		{
			MethodName: "DelUser",
			Handler:    _TestUserService_DelUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpcTemplate.proto",
}
