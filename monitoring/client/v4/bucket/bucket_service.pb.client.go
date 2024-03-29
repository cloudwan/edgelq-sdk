// Code generated by protoc-gen-goten-client
// API: BucketService
// DO NOT EDIT!!!

package bucket_client

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/runtime/protoimpl"
)

// proto imports
import (
	bucket "github.com/cloudwan/edgelq-sdk/monitoring/resources/v4/bucket"
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
	_ = &bucket.Bucket{}
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

// BucketServiceClient is the client API for BucketService.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BucketServiceClient interface {
	GetBucket(ctx context.Context, in *GetBucketRequest, opts ...grpc.CallOption) (*bucket.Bucket, error)
	BatchGetBuckets(ctx context.Context, in *BatchGetBucketsRequest, opts ...grpc.CallOption) (*BatchGetBucketsResponse, error)
	ListBuckets(ctx context.Context, in *ListBucketsRequest, opts ...grpc.CallOption) (*ListBucketsResponse, error)
	WatchBucket(ctx context.Context, in *WatchBucketRequest, opts ...grpc.CallOption) (WatchBucketClientStream, error)
	WatchBuckets(ctx context.Context, in *WatchBucketsRequest, opts ...grpc.CallOption) (WatchBucketsClientStream, error)
	CreateBucket(ctx context.Context, in *CreateBucketRequest, opts ...grpc.CallOption) (*bucket.Bucket, error)
	UpdateBucket(ctx context.Context, in *UpdateBucketRequest, opts ...grpc.CallOption) (*bucket.Bucket, error)
	DeleteBucket(ctx context.Context, in *DeleteBucketRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type client struct {
	cc grpc.ClientConnInterface
}

func NewBucketServiceClient(cc grpc.ClientConnInterface) BucketServiceClient {
	return &client{cc}
}

func (c *client) GetBucket(ctx context.Context, in *GetBucketRequest, opts ...grpc.CallOption) (*bucket.Bucket, error) {
	out := new(bucket.Bucket)
	err := c.cc.Invoke(ctx, "/ntt.monitoring.v4.BucketService/GetBucket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) BatchGetBuckets(ctx context.Context, in *BatchGetBucketsRequest, opts ...grpc.CallOption) (*BatchGetBucketsResponse, error) {
	out := new(BatchGetBucketsResponse)
	err := c.cc.Invoke(ctx, "/ntt.monitoring.v4.BucketService/BatchGetBuckets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) ListBuckets(ctx context.Context, in *ListBucketsRequest, opts ...grpc.CallOption) (*ListBucketsResponse, error) {
	out := new(ListBucketsResponse)
	err := c.cc.Invoke(ctx, "/ntt.monitoring.v4.BucketService/ListBuckets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) WatchBucket(ctx context.Context, in *WatchBucketRequest, opts ...grpc.CallOption) (WatchBucketClientStream, error) {
	stream, err := c.cc.NewStream(ctx,
		&grpc.StreamDesc{
			StreamName:    "WatchBucket",
			ServerStreams: true,
		},
		"/ntt.monitoring.v4.BucketService/WatchBucket", opts...)
	if err != nil {
		return nil, err
	}
	x := &watchBucketWatchBucketClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WatchBucketClientStream interface {
	Recv() (*WatchBucketResponse, error)
	grpc.ClientStream
}

type watchBucketWatchBucketClient struct {
	grpc.ClientStream
}

func (x *watchBucketWatchBucketClient) Recv() (*WatchBucketResponse, error) {
	m := new(WatchBucketResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *client) WatchBuckets(ctx context.Context, in *WatchBucketsRequest, opts ...grpc.CallOption) (WatchBucketsClientStream, error) {
	stream, err := c.cc.NewStream(ctx,
		&grpc.StreamDesc{
			StreamName:    "WatchBuckets",
			ServerStreams: true,
		},
		"/ntt.monitoring.v4.BucketService/WatchBuckets", opts...)
	if err != nil {
		return nil, err
	}
	x := &watchBucketsWatchBucketsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WatchBucketsClientStream interface {
	Recv() (*WatchBucketsResponse, error)
	grpc.ClientStream
}

type watchBucketsWatchBucketsClient struct {
	grpc.ClientStream
}

func (x *watchBucketsWatchBucketsClient) Recv() (*WatchBucketsResponse, error) {
	m := new(WatchBucketsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *client) CreateBucket(ctx context.Context, in *CreateBucketRequest, opts ...grpc.CallOption) (*bucket.Bucket, error) {
	out := new(bucket.Bucket)
	err := c.cc.Invoke(ctx, "/ntt.monitoring.v4.BucketService/CreateBucket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) UpdateBucket(ctx context.Context, in *UpdateBucketRequest, opts ...grpc.CallOption) (*bucket.Bucket, error) {
	out := new(bucket.Bucket)
	err := c.cc.Invoke(ctx, "/ntt.monitoring.v4.BucketService/UpdateBucket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) DeleteBucket(ctx context.Context, in *DeleteBucketRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ntt.monitoring.v4.BucketService/DeleteBucket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
