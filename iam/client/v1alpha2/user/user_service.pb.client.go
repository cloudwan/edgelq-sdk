// Code generated by protoc-gen-goten-client
// API: UserService
// DO NOT EDIT!!!

package user_client

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/runtime/protoimpl"
)

// proto imports
import (
	user "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/user"
	empty "github.com/golang/protobuf/ptypes/empty"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = protoimpl.DescBuilder{}
	_ = context.Context(nil)
	_ = grpc.ClientConn{}
)

// make sure we're using proto imports
var (
	_ = &user.User{}
	_ = &empty.Empty{}
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UserServiceClient is the client API for UserService.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*user.User, error)
	BatchGetUsers(ctx context.Context, in *BatchGetUsersRequest, opts ...grpc.CallOption) (*BatchGetUsersResponse, error)
	ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error)
	WatchUser(ctx context.Context, in *WatchUserRequest, opts ...grpc.CallOption) (WatchUserClientStream, error)
	WatchUsers(ctx context.Context, in *WatchUsersRequest, opts ...grpc.CallOption) (WatchUsersClientStream, error)
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*user.User, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*user.User, error)
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	GetUserByEmail(ctx context.Context, in *GetUserByEmailRequest, opts ...grpc.CallOption) (*user.User, error)
	BatchGetUsersByEmail(ctx context.Context, in *BatchGetUsersByEmailRequest, opts ...grpc.CallOption) (*BatchGetUsersByEmailResponse, error)
	GetMySettings(ctx context.Context, in *GetMySettingsRequest, opts ...grpc.CallOption) (*GetMySettingsResponse, error)
	SetMySettings(ctx context.Context, in *SetMySettingsRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	RefreshUserFromIdToken(ctx context.Context, in *RefreshUserFromIdTokenRequest, opts ...grpc.CallOption) (*RefreshUserFromIdTokenResponse, error)
	ResendVerificationEmail(ctx context.Context, in *ResendVerificationEmailRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type client struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &client{cc}
}

func (c *client) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*user.User, error) {
	out := new(user.User)
	err := c.cc.Invoke(ctx, "/ntt.iam.v1alpha2.UserService/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) BatchGetUsers(ctx context.Context, in *BatchGetUsersRequest, opts ...grpc.CallOption) (*BatchGetUsersResponse, error) {
	out := new(BatchGetUsersResponse)
	err := c.cc.Invoke(ctx, "/ntt.iam.v1alpha2.UserService/BatchGetUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error) {
	out := new(ListUsersResponse)
	err := c.cc.Invoke(ctx, "/ntt.iam.v1alpha2.UserService/ListUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) WatchUser(ctx context.Context, in *WatchUserRequest, opts ...grpc.CallOption) (WatchUserClientStream, error) {
	stream, err := c.cc.NewStream(ctx,
		&grpc.StreamDesc{
			StreamName:    "WatchUser",
			ServerStreams: true,
		},
		"/ntt.iam.v1alpha2.UserService/WatchUser", opts...)
	if err != nil {
		return nil, err
	}
	x := &watchUserWatchUserClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WatchUserClientStream interface {
	Recv() (*WatchUserResponse, error)
	grpc.ClientStream
}

type watchUserWatchUserClient struct {
	grpc.ClientStream
}

func (x *watchUserWatchUserClient) Recv() (*WatchUserResponse, error) {
	m := new(WatchUserResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *client) WatchUsers(ctx context.Context, in *WatchUsersRequest, opts ...grpc.CallOption) (WatchUsersClientStream, error) {
	stream, err := c.cc.NewStream(ctx,
		&grpc.StreamDesc{
			StreamName:    "WatchUsers",
			ServerStreams: true,
		},
		"/ntt.iam.v1alpha2.UserService/WatchUsers", opts...)
	if err != nil {
		return nil, err
	}
	x := &watchUsersWatchUsersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WatchUsersClientStream interface {
	Recv() (*WatchUsersResponse, error)
	grpc.ClientStream
}

type watchUsersWatchUsersClient struct {
	grpc.ClientStream
}

func (x *watchUsersWatchUsersClient) Recv() (*WatchUsersResponse, error) {
	m := new(WatchUsersResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *client) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*user.User, error) {
	out := new(user.User)
	err := c.cc.Invoke(ctx, "/ntt.iam.v1alpha2.UserService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*user.User, error) {
	out := new(user.User)
	err := c.cc.Invoke(ctx, "/ntt.iam.v1alpha2.UserService/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/ntt.iam.v1alpha2.UserService/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) GetUserByEmail(ctx context.Context, in *GetUserByEmailRequest, opts ...grpc.CallOption) (*user.User, error) {
	out := new(user.User)
	err := c.cc.Invoke(ctx, "/ntt.iam.v1alpha2.UserService/GetUserByEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) BatchGetUsersByEmail(ctx context.Context, in *BatchGetUsersByEmailRequest, opts ...grpc.CallOption) (*BatchGetUsersByEmailResponse, error) {
	out := new(BatchGetUsersByEmailResponse)
	err := c.cc.Invoke(ctx, "/ntt.iam.v1alpha2.UserService/BatchGetUsersByEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) GetMySettings(ctx context.Context, in *GetMySettingsRequest, opts ...grpc.CallOption) (*GetMySettingsResponse, error) {
	out := new(GetMySettingsResponse)
	err := c.cc.Invoke(ctx, "/ntt.iam.v1alpha2.UserService/GetMySettings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) SetMySettings(ctx context.Context, in *SetMySettingsRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/ntt.iam.v1alpha2.UserService/SetMySettings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) RefreshUserFromIdToken(ctx context.Context, in *RefreshUserFromIdTokenRequest, opts ...grpc.CallOption) (*RefreshUserFromIdTokenResponse, error) {
	out := new(RefreshUserFromIdTokenResponse)
	err := c.cc.Invoke(ctx, "/ntt.iam.v1alpha2.UserService/RefreshUserFromIdToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) ResendVerificationEmail(ctx context.Context, in *ResendVerificationEmailRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/ntt.iam.v1alpha2.UserService/ResendVerificationEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
