// Code generated by protoc-gen-goten-client
// API: RoleBindingService
// DO NOT EDIT!!!

package role_binding_client

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/runtime/protoimpl"
)

// proto imports
import (
	role_binding "github.com/cloudwan/edgelq-sdk/iam/resources/v1/role_binding"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = protoimpl.DescBuilder{}
	_ = new(context.Context)
	_ = grpc.ClientConn{}
)

// make sure we're using proto imports
var (
	_ = &role_binding.RoleBinding{}
	_ = &emptypb.Empty{}
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

// RoleBindingServiceClient is the client API for RoleBindingService.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RoleBindingServiceClient interface {
	GetRoleBinding(ctx context.Context, in *GetRoleBindingRequest, opts ...grpc.CallOption) (*role_binding.RoleBinding, error)
	BatchGetRoleBindings(ctx context.Context, in *BatchGetRoleBindingsRequest, opts ...grpc.CallOption) (*BatchGetRoleBindingsResponse, error)
	ListRoleBindings(ctx context.Context, in *ListRoleBindingsRequest, opts ...grpc.CallOption) (*ListRoleBindingsResponse, error)
	WatchRoleBinding(ctx context.Context, in *WatchRoleBindingRequest, opts ...grpc.CallOption) (WatchRoleBindingClientStream, error)
	WatchRoleBindings(ctx context.Context, in *WatchRoleBindingsRequest, opts ...grpc.CallOption) (WatchRoleBindingsClientStream, error)
	CreateRoleBinding(ctx context.Context, in *CreateRoleBindingRequest, opts ...grpc.CallOption) (*role_binding.RoleBinding, error)
	UpdateRoleBinding(ctx context.Context, in *UpdateRoleBindingRequest, opts ...grpc.CallOption) (*role_binding.RoleBinding, error)
	DeleteRoleBinding(ctx context.Context, in *DeleteRoleBindingRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ListRoleBindingsWithMembers(ctx context.Context, in *ListRoleBindingsWithMembersRequest, opts ...grpc.CallOption) (*ListRoleBindingsWithMembersResponse, error)
}

type client struct {
	cc grpc.ClientConnInterface
}

func NewRoleBindingServiceClient(cc grpc.ClientConnInterface) RoleBindingServiceClient {
	return &client{cc}
}

func (c *client) GetRoleBinding(ctx context.Context, in *GetRoleBindingRequest, opts ...grpc.CallOption) (*role_binding.RoleBinding, error) {
	out := new(role_binding.RoleBinding)
	err := c.cc.Invoke(ctx, "/ntt.iam.v1.RoleBindingService/GetRoleBinding", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) BatchGetRoleBindings(ctx context.Context, in *BatchGetRoleBindingsRequest, opts ...grpc.CallOption) (*BatchGetRoleBindingsResponse, error) {
	out := new(BatchGetRoleBindingsResponse)
	err := c.cc.Invoke(ctx, "/ntt.iam.v1.RoleBindingService/BatchGetRoleBindings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) ListRoleBindings(ctx context.Context, in *ListRoleBindingsRequest, opts ...grpc.CallOption) (*ListRoleBindingsResponse, error) {
	out := new(ListRoleBindingsResponse)
	err := c.cc.Invoke(ctx, "/ntt.iam.v1.RoleBindingService/ListRoleBindings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) WatchRoleBinding(ctx context.Context, in *WatchRoleBindingRequest, opts ...grpc.CallOption) (WatchRoleBindingClientStream, error) {
	stream, err := c.cc.NewStream(ctx,
		&grpc.StreamDesc{
			StreamName:    "WatchRoleBinding",
			ServerStreams: true,
		},
		"/ntt.iam.v1.RoleBindingService/WatchRoleBinding", opts...)
	if err != nil {
		return nil, err
	}
	x := &watchRoleBindingWatchRoleBindingClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WatchRoleBindingClientStream interface {
	Recv() (*WatchRoleBindingResponse, error)
	grpc.ClientStream
}

type watchRoleBindingWatchRoleBindingClient struct {
	grpc.ClientStream
}

func (x *watchRoleBindingWatchRoleBindingClient) Recv() (*WatchRoleBindingResponse, error) {
	m := new(WatchRoleBindingResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *client) WatchRoleBindings(ctx context.Context, in *WatchRoleBindingsRequest, opts ...grpc.CallOption) (WatchRoleBindingsClientStream, error) {
	stream, err := c.cc.NewStream(ctx,
		&grpc.StreamDesc{
			StreamName:    "WatchRoleBindings",
			ServerStreams: true,
		},
		"/ntt.iam.v1.RoleBindingService/WatchRoleBindings", opts...)
	if err != nil {
		return nil, err
	}
	x := &watchRoleBindingsWatchRoleBindingsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WatchRoleBindingsClientStream interface {
	Recv() (*WatchRoleBindingsResponse, error)
	grpc.ClientStream
}

type watchRoleBindingsWatchRoleBindingsClient struct {
	grpc.ClientStream
}

func (x *watchRoleBindingsWatchRoleBindingsClient) Recv() (*WatchRoleBindingsResponse, error) {
	m := new(WatchRoleBindingsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *client) CreateRoleBinding(ctx context.Context, in *CreateRoleBindingRequest, opts ...grpc.CallOption) (*role_binding.RoleBinding, error) {
	out := new(role_binding.RoleBinding)
	err := c.cc.Invoke(ctx, "/ntt.iam.v1.RoleBindingService/CreateRoleBinding", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) UpdateRoleBinding(ctx context.Context, in *UpdateRoleBindingRequest, opts ...grpc.CallOption) (*role_binding.RoleBinding, error) {
	out := new(role_binding.RoleBinding)
	err := c.cc.Invoke(ctx, "/ntt.iam.v1.RoleBindingService/UpdateRoleBinding", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) DeleteRoleBinding(ctx context.Context, in *DeleteRoleBindingRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ntt.iam.v1.RoleBindingService/DeleteRoleBinding", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) ListRoleBindingsWithMembers(ctx context.Context, in *ListRoleBindingsWithMembersRequest, opts ...grpc.CallOption) (*ListRoleBindingsWithMembersResponse, error) {
	out := new(ListRoleBindingsWithMembersResponse)
	err := c.cc.Invoke(ctx, "/ntt.iam.v1.RoleBindingService/ListRoleBindingsWithMembers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
