// Code generated by protoc-gen-goten-client
// API: AuthorizationService
// DO NOT EDIT!!!

package authorization_client

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/runtime/protoimpl"
)

// proto imports
import (
	role_binding "github.com/cloudwan/edgelq-sdk/iam/resources/v1/role_binding"
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

// AuthorizationServiceClient is the client API for AuthorizationService.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthorizationServiceClient interface {
	GetPrincipal(ctx context.Context, in *GetPrincipalRequest, opts ...grpc.CallOption) (*GetPrincipalResponse, error)
	WatchPrincipalUpdates(ctx context.Context, in *WatchPrincipalUpdatesRequest, opts ...grpc.CallOption) (WatchPrincipalUpdatesClientStream, error)
	CheckMyRoleBindings(ctx context.Context, in *CheckMyRoleBindingsRequest, opts ...grpc.CallOption) (*CheckMyRoleBindingsResponse, error)
}

type client struct {
	cc grpc.ClientConnInterface
}

func NewAuthorizationServiceClient(cc grpc.ClientConnInterface) AuthorizationServiceClient {
	return &client{cc}
}

func (c *client) GetPrincipal(ctx context.Context, in *GetPrincipalRequest, opts ...grpc.CallOption) (*GetPrincipalResponse, error) {
	out := new(GetPrincipalResponse)
	err := c.cc.Invoke(ctx, "/ntt.iam.v1.AuthorizationService/GetPrincipal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) WatchPrincipalUpdates(ctx context.Context, in *WatchPrincipalUpdatesRequest, opts ...grpc.CallOption) (WatchPrincipalUpdatesClientStream, error) {
	stream, err := c.cc.NewStream(ctx,
		&grpc.StreamDesc{
			StreamName:    "WatchPrincipalUpdates",
			ServerStreams: true,
		},
		"/ntt.iam.v1.AuthorizationService/WatchPrincipalUpdates", opts...)
	if err != nil {
		return nil, err
	}
	x := &watchPrincipalUpdatesWatchPrincipalUpdatesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WatchPrincipalUpdatesClientStream interface {
	Recv() (*WatchPrincipalUpdatesResponse, error)
	grpc.ClientStream
}

type watchPrincipalUpdatesWatchPrincipalUpdatesClient struct {
	grpc.ClientStream
}

func (x *watchPrincipalUpdatesWatchPrincipalUpdatesClient) Recv() (*WatchPrincipalUpdatesResponse, error) {
	m := new(WatchPrincipalUpdatesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *client) CheckMyRoleBindings(ctx context.Context, in *CheckMyRoleBindingsRequest, opts ...grpc.CallOption) (*CheckMyRoleBindingsResponse, error) {
	out := new(CheckMyRoleBindingsResponse)
	err := c.cc.Invoke(ctx, "/ntt.iam.v1.AuthorizationService/CheckMyRoleBindings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}