// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// UserProfilesClient is the client API for UserProfiles service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserProfilesClient interface {
	CreateUserProfile(ctx context.Context, in *CreateUserProfileRequest, opts ...grpc.CallOption) (*UserProfile, error)
	DeleteUserProfile(ctx context.Context, in *DeleteUserProfileRequest, opts ...grpc.CallOption) (*EmptyReq, error)
	ListUserProfiles(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*ListUserProfilesResponse, error)
}

type userProfilesClient struct {
	cc grpc.ClientConnInterface
}

func NewUserProfilesClient(cc grpc.ClientConnInterface) UserProfilesClient {
	return &userProfilesClient{cc}
}

func (c *userProfilesClient) CreateUserProfile(ctx context.Context, in *CreateUserProfileRequest, opts ...grpc.CallOption) (*UserProfile, error) {
	out := new(UserProfile)
	err := c.cc.Invoke(ctx, "/api.UserProfiles/CreateUserProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userProfilesClient) DeleteUserProfile(ctx context.Context, in *DeleteUserProfileRequest, opts ...grpc.CallOption) (*EmptyReq, error) {
	out := new(EmptyReq)
	err := c.cc.Invoke(ctx, "/api.UserProfiles/DeleteUserProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userProfilesClient) ListUserProfiles(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*ListUserProfilesResponse, error) {
	out := new(ListUserProfilesResponse)
	err := c.cc.Invoke(ctx, "/api.UserProfiles/ListUserProfiles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserProfilesServer is the server API for UserProfiles service.
// All implementations should embed UnimplementedUserProfilesServer
// for forward compatibility
type UserProfilesServer interface {
	CreateUserProfile(context.Context, *CreateUserProfileRequest) (*UserProfile, error)
	DeleteUserProfile(context.Context, *DeleteUserProfileRequest) (*EmptyReq, error)
	ListUserProfiles(context.Context, *EmptyReq) (*ListUserProfilesResponse, error)
}

// UnimplementedUserProfilesServer should be embedded to have forward compatible implementations.
type UnimplementedUserProfilesServer struct {
}

func (UnimplementedUserProfilesServer) CreateUserProfile(context.Context, *CreateUserProfileRequest) (*UserProfile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUserProfile not implemented")
}
func (UnimplementedUserProfilesServer) DeleteUserProfile(context.Context, *DeleteUserProfileRequest) (*EmptyReq, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserProfile not implemented")
}
func (UnimplementedUserProfilesServer) ListUserProfiles(context.Context, *EmptyReq) (*ListUserProfilesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUserProfiles not implemented")
}

// UnsafeUserProfilesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserProfilesServer will
// result in compilation errors.
type UnsafeUserProfilesServer interface {
	mustEmbedUnimplementedUserProfilesServer()
}

func RegisterUserProfilesServer(s grpc.ServiceRegistrar, srv UserProfilesServer) {
	s.RegisterService(&UserProfiles_ServiceDesc, srv)
}

func _UserProfiles_CreateUserProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserProfilesServer).CreateUserProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.UserProfiles/CreateUserProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserProfilesServer).CreateUserProfile(ctx, req.(*CreateUserProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserProfiles_DeleteUserProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserProfilesServer).DeleteUserProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.UserProfiles/DeleteUserProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserProfilesServer).DeleteUserProfile(ctx, req.(*DeleteUserProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserProfiles_ListUserProfiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserProfilesServer).ListUserProfiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.UserProfiles/ListUserProfiles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserProfilesServer).ListUserProfiles(ctx, req.(*EmptyReq))
	}
	return interceptor(ctx, in, info, handler)
}

// UserProfiles_ServiceDesc is the grpc.ServiceDesc for UserProfiles service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserProfiles_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.UserProfiles",
	HandlerType: (*UserProfilesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUserProfile",
			Handler:    _UserProfiles_CreateUserProfile_Handler,
		},
		{
			MethodName: "DeleteUserProfile",
			Handler:    _UserProfiles_DeleteUserProfile_Handler,
		},
		{
			MethodName: "ListUserProfiles",
			Handler:    _UserProfiles_ListUserProfiles_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/user.proto",
}