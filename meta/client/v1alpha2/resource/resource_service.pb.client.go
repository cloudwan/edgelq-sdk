// Code generated by protoc-gen-goten-client
// API: ResourceService
// DO NOT EDIT!!!

package resource_client

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/runtime/protoimpl"
)

// proto imports
import (
	resource "github.com/cloudwan/edgelq-sdk/meta/resources/v1alpha2/resource"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = protoimpl.DescBuilder{}
	_ = new(context.Context)
	_ = grpc.ClientConn{}
)

// make sure we're using proto imports
var (
	_ = &resource.Resource{}
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

// ResourceServiceClient is the client API for ResourceService.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ResourceServiceClient interface {
	GetResource(ctx context.Context, in *GetResourceRequest, opts ...grpc.CallOption) (*resource.Resource, error)
	BatchGetResources(ctx context.Context, in *BatchGetResourcesRequest, opts ...grpc.CallOption) (*BatchGetResourcesResponse, error)
	ListResources(ctx context.Context, in *ListResourcesRequest, opts ...grpc.CallOption) (*ListResourcesResponse, error)
	WatchResource(ctx context.Context, in *WatchResourceRequest, opts ...grpc.CallOption) (WatchResourceClientStream, error)
	WatchResources(ctx context.Context, in *WatchResourcesRequest, opts ...grpc.CallOption) (WatchResourcesClientStream, error)
}

type client struct {
	cc grpc.ClientConnInterface
}

func NewResourceServiceClient(cc grpc.ClientConnInterface) ResourceServiceClient {
	return &client{cc}
}

func (c *client) GetResource(ctx context.Context, in *GetResourceRequest, opts ...grpc.CallOption) (*resource.Resource, error) {
	out := new(resource.Resource)
	err := c.cc.Invoke(ctx, "/ntt.meta.v1alpha2.ResourceService/GetResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) BatchGetResources(ctx context.Context, in *BatchGetResourcesRequest, opts ...grpc.CallOption) (*BatchGetResourcesResponse, error) {
	out := new(BatchGetResourcesResponse)
	err := c.cc.Invoke(ctx, "/ntt.meta.v1alpha2.ResourceService/BatchGetResources", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) ListResources(ctx context.Context, in *ListResourcesRequest, opts ...grpc.CallOption) (*ListResourcesResponse, error) {
	out := new(ListResourcesResponse)
	err := c.cc.Invoke(ctx, "/ntt.meta.v1alpha2.ResourceService/ListResources", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) WatchResource(ctx context.Context, in *WatchResourceRequest, opts ...grpc.CallOption) (WatchResourceClientStream, error) {
	stream, err := c.cc.NewStream(ctx,
		&grpc.StreamDesc{
			StreamName:    "WatchResource",
			ServerStreams: true,
		},
		"/ntt.meta.v1alpha2.ResourceService/WatchResource", opts...)
	if err != nil {
		return nil, err
	}
	x := &watchResourceWatchResourceClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WatchResourceClientStream interface {
	Recv() (*WatchResourceResponse, error)
	grpc.ClientStream
}

type watchResourceWatchResourceClient struct {
	grpc.ClientStream
}

func (x *watchResourceWatchResourceClient) Recv() (*WatchResourceResponse, error) {
	m := new(WatchResourceResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *client) WatchResources(ctx context.Context, in *WatchResourcesRequest, opts ...grpc.CallOption) (WatchResourcesClientStream, error) {
	stream, err := c.cc.NewStream(ctx,
		&grpc.StreamDesc{
			StreamName:    "WatchResources",
			ServerStreams: true,
		},
		"/ntt.meta.v1alpha2.ResourceService/WatchResources", opts...)
	if err != nil {
		return nil, err
	}
	x := &watchResourcesWatchResourcesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WatchResourcesClientStream interface {
	Recv() (*WatchResourcesResponse, error)
	grpc.ClientStream
}

type watchResourcesWatchResourcesClient struct {
	grpc.ClientStream
}

func (x *watchResourcesWatchResourcesClient) Recv() (*WatchResourcesResponse, error) {
	m := new(WatchResourcesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}
