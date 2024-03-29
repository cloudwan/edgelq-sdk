// Code generated by protoc-gen-goten-client
// API: LimitPoolService
// DO NOT EDIT!!!

package limit_pool_client

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/runtime/protoimpl"
)

// proto imports
import (
	limit_pool "github.com/cloudwan/edgelq-sdk/limits/resources/v1alpha2/limit_pool"
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
	_ = &limit_pool.LimitPool{}
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

// LimitPoolServiceClient is the client API for LimitPoolService.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LimitPoolServiceClient interface {
	GetLimitPool(ctx context.Context, in *GetLimitPoolRequest, opts ...grpc.CallOption) (*limit_pool.LimitPool, error)
	BatchGetLimitPools(ctx context.Context, in *BatchGetLimitPoolsRequest, opts ...grpc.CallOption) (*BatchGetLimitPoolsResponse, error)
	ListLimitPools(ctx context.Context, in *ListLimitPoolsRequest, opts ...grpc.CallOption) (*ListLimitPoolsResponse, error)
	WatchLimitPool(ctx context.Context, in *WatchLimitPoolRequest, opts ...grpc.CallOption) (WatchLimitPoolClientStream, error)
	WatchLimitPools(ctx context.Context, in *WatchLimitPoolsRequest, opts ...grpc.CallOption) (WatchLimitPoolsClientStream, error)
	UpdateLimitPool(ctx context.Context, in *UpdateLimitPoolRequest, opts ...grpc.CallOption) (*limit_pool.LimitPool, error)
	DeleteLimitPool(ctx context.Context, in *DeleteLimitPoolRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	MigrateLimitPoolSource(ctx context.Context, in *MigrateLimitPoolSourceRequest, opts ...grpc.CallOption) (*limit_pool.LimitPool, error)
}

type client struct {
	cc grpc.ClientConnInterface
}

func NewLimitPoolServiceClient(cc grpc.ClientConnInterface) LimitPoolServiceClient {
	return &client{cc}
}

func (c *client) GetLimitPool(ctx context.Context, in *GetLimitPoolRequest, opts ...grpc.CallOption) (*limit_pool.LimitPool, error) {
	out := new(limit_pool.LimitPool)
	err := c.cc.Invoke(ctx, "/ntt.limits.v1alpha2.LimitPoolService/GetLimitPool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) BatchGetLimitPools(ctx context.Context, in *BatchGetLimitPoolsRequest, opts ...grpc.CallOption) (*BatchGetLimitPoolsResponse, error) {
	out := new(BatchGetLimitPoolsResponse)
	err := c.cc.Invoke(ctx, "/ntt.limits.v1alpha2.LimitPoolService/BatchGetLimitPools", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) ListLimitPools(ctx context.Context, in *ListLimitPoolsRequest, opts ...grpc.CallOption) (*ListLimitPoolsResponse, error) {
	out := new(ListLimitPoolsResponse)
	err := c.cc.Invoke(ctx, "/ntt.limits.v1alpha2.LimitPoolService/ListLimitPools", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) WatchLimitPool(ctx context.Context, in *WatchLimitPoolRequest, opts ...grpc.CallOption) (WatchLimitPoolClientStream, error) {
	stream, err := c.cc.NewStream(ctx,
		&grpc.StreamDesc{
			StreamName:    "WatchLimitPool",
			ServerStreams: true,
		},
		"/ntt.limits.v1alpha2.LimitPoolService/WatchLimitPool", opts...)
	if err != nil {
		return nil, err
	}
	x := &watchLimitPoolWatchLimitPoolClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WatchLimitPoolClientStream interface {
	Recv() (*WatchLimitPoolResponse, error)
	grpc.ClientStream
}

type watchLimitPoolWatchLimitPoolClient struct {
	grpc.ClientStream
}

func (x *watchLimitPoolWatchLimitPoolClient) Recv() (*WatchLimitPoolResponse, error) {
	m := new(WatchLimitPoolResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *client) WatchLimitPools(ctx context.Context, in *WatchLimitPoolsRequest, opts ...grpc.CallOption) (WatchLimitPoolsClientStream, error) {
	stream, err := c.cc.NewStream(ctx,
		&grpc.StreamDesc{
			StreamName:    "WatchLimitPools",
			ServerStreams: true,
		},
		"/ntt.limits.v1alpha2.LimitPoolService/WatchLimitPools", opts...)
	if err != nil {
		return nil, err
	}
	x := &watchLimitPoolsWatchLimitPoolsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WatchLimitPoolsClientStream interface {
	Recv() (*WatchLimitPoolsResponse, error)
	grpc.ClientStream
}

type watchLimitPoolsWatchLimitPoolsClient struct {
	grpc.ClientStream
}

func (x *watchLimitPoolsWatchLimitPoolsClient) Recv() (*WatchLimitPoolsResponse, error) {
	m := new(WatchLimitPoolsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *client) UpdateLimitPool(ctx context.Context, in *UpdateLimitPoolRequest, opts ...grpc.CallOption) (*limit_pool.LimitPool, error) {
	out := new(limit_pool.LimitPool)
	err := c.cc.Invoke(ctx, "/ntt.limits.v1alpha2.LimitPoolService/UpdateLimitPool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) DeleteLimitPool(ctx context.Context, in *DeleteLimitPoolRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ntt.limits.v1alpha2.LimitPoolService/DeleteLimitPool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) MigrateLimitPoolSource(ctx context.Context, in *MigrateLimitPoolSourceRequest, opts ...grpc.CallOption) (*limit_pool.LimitPool, error) {
	out := new(limit_pool.LimitPool)
	err := c.cc.Invoke(ctx, "/ntt.limits.v1alpha2.LimitPoolService/MigrateLimitPoolSource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
