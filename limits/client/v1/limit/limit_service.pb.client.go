// Code generated by protoc-gen-goten-client
// API: LimitService
// DO NOT EDIT!!!

package limit_client

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/runtime/protoimpl"
)

// proto imports
import (
	limit "github.com/cloudwan/edgelq-sdk/limits/resources/v1/limit"
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
	_ = &limit.Limit{}
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

// LimitServiceClient is the client API for LimitService.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LimitServiceClient interface {
	GetLimit(ctx context.Context, in *GetLimitRequest, opts ...grpc.CallOption) (*limit.Limit, error)
	BatchGetLimits(ctx context.Context, in *BatchGetLimitsRequest, opts ...grpc.CallOption) (*BatchGetLimitsResponse, error)
	ListLimits(ctx context.Context, in *ListLimitsRequest, opts ...grpc.CallOption) (*ListLimitsResponse, error)
	WatchLimit(ctx context.Context, in *WatchLimitRequest, opts ...grpc.CallOption) (WatchLimitClientStream, error)
	WatchLimits(ctx context.Context, in *WatchLimitsRequest, opts ...grpc.CallOption) (WatchLimitsClientStream, error)
	UpdateLimit(ctx context.Context, in *UpdateLimitRequest, opts ...grpc.CallOption) (*limit.Limit, error)
	DeleteLimit(ctx context.Context, in *DeleteLimitRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	MigrateLimitSource(ctx context.Context, in *MigrateLimitSourceRequest, opts ...grpc.CallOption) (*limit.Limit, error)
}

type client struct {
	cc grpc.ClientConnInterface
}

func NewLimitServiceClient(cc grpc.ClientConnInterface) LimitServiceClient {
	return &client{cc}
}

func (c *client) GetLimit(ctx context.Context, in *GetLimitRequest, opts ...grpc.CallOption) (*limit.Limit, error) {
	out := new(limit.Limit)
	err := c.cc.Invoke(ctx, "/ntt.limits.v1.LimitService/GetLimit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) BatchGetLimits(ctx context.Context, in *BatchGetLimitsRequest, opts ...grpc.CallOption) (*BatchGetLimitsResponse, error) {
	out := new(BatchGetLimitsResponse)
	err := c.cc.Invoke(ctx, "/ntt.limits.v1.LimitService/BatchGetLimits", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) ListLimits(ctx context.Context, in *ListLimitsRequest, opts ...grpc.CallOption) (*ListLimitsResponse, error) {
	out := new(ListLimitsResponse)
	err := c.cc.Invoke(ctx, "/ntt.limits.v1.LimitService/ListLimits", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) WatchLimit(ctx context.Context, in *WatchLimitRequest, opts ...grpc.CallOption) (WatchLimitClientStream, error) {
	stream, err := c.cc.NewStream(ctx,
		&grpc.StreamDesc{
			StreamName:    "WatchLimit",
			ServerStreams: true,
		},
		"/ntt.limits.v1.LimitService/WatchLimit", opts...)
	if err != nil {
		return nil, err
	}
	x := &watchLimitWatchLimitClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WatchLimitClientStream interface {
	Recv() (*WatchLimitResponse, error)
	grpc.ClientStream
}

type watchLimitWatchLimitClient struct {
	grpc.ClientStream
}

func (x *watchLimitWatchLimitClient) Recv() (*WatchLimitResponse, error) {
	m := new(WatchLimitResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *client) WatchLimits(ctx context.Context, in *WatchLimitsRequest, opts ...grpc.CallOption) (WatchLimitsClientStream, error) {
	stream, err := c.cc.NewStream(ctx,
		&grpc.StreamDesc{
			StreamName:    "WatchLimits",
			ServerStreams: true,
		},
		"/ntt.limits.v1.LimitService/WatchLimits", opts...)
	if err != nil {
		return nil, err
	}
	x := &watchLimitsWatchLimitsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WatchLimitsClientStream interface {
	Recv() (*WatchLimitsResponse, error)
	grpc.ClientStream
}

type watchLimitsWatchLimitsClient struct {
	grpc.ClientStream
}

func (x *watchLimitsWatchLimitsClient) Recv() (*WatchLimitsResponse, error) {
	m := new(WatchLimitsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *client) UpdateLimit(ctx context.Context, in *UpdateLimitRequest, opts ...grpc.CallOption) (*limit.Limit, error) {
	out := new(limit.Limit)
	err := c.cc.Invoke(ctx, "/ntt.limits.v1.LimitService/UpdateLimit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) DeleteLimit(ctx context.Context, in *DeleteLimitRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ntt.limits.v1.LimitService/DeleteLimit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) MigrateLimitSource(ctx context.Context, in *MigrateLimitSourceRequest, opts ...grpc.CallOption) (*limit.Limit, error) {
	out := new(limit.Limit)
	err := c.cc.Invoke(ctx, "/ntt.limits.v1.LimitService/MigrateLimitSource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
