// Code generated by protoc-gen-goten-client
// API: MonitoredResourceDescriptorService
// DO NOT EDIT!!!

package monitored_resource_descriptor_client

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/runtime/protoimpl"
)

// proto imports
import (
	monitored_resource_descriptor "github.com/cloudwan/edgelq-sdk/monitoring/resources/v3/monitored_resource_descriptor"
	empty "github.com/golang/protobuf/ptypes/empty"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = protoimpl.DescBuilder{}
	_ = new(context.Context)
	_ = grpc.ClientConn{}
)

// make sure we're using proto imports
var (
	_ = &monitored_resource_descriptor.MonitoredResourceDescriptor{}
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

// MonitoredResourceDescriptorServiceClient is the client API for MonitoredResourceDescriptorService.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MonitoredResourceDescriptorServiceClient interface {
	BatchGetMonitoredResourceDescriptors(ctx context.Context, in *BatchGetMonitoredResourceDescriptorsRequest, opts ...grpc.CallOption) (*BatchGetMonitoredResourceDescriptorsResponse, error)
	WatchMonitoredResourceDescriptor(ctx context.Context, in *WatchMonitoredResourceDescriptorRequest, opts ...grpc.CallOption) (WatchMonitoredResourceDescriptorClientStream, error)
	WatchMonitoredResourceDescriptors(ctx context.Context, in *WatchMonitoredResourceDescriptorsRequest, opts ...grpc.CallOption) (WatchMonitoredResourceDescriptorsClientStream, error)
	CreateMonitoredResourceDescriptor(ctx context.Context, in *CreateMonitoredResourceDescriptorRequest, opts ...grpc.CallOption) (*monitored_resource_descriptor.MonitoredResourceDescriptor, error)
	UpdateMonitoredResourceDescriptor(ctx context.Context, in *UpdateMonitoredResourceDescriptorRequest, opts ...grpc.CallOption) (*monitored_resource_descriptor.MonitoredResourceDescriptor, error)
	DeleteMonitoredResourceDescriptor(ctx context.Context, in *DeleteMonitoredResourceDescriptorRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	GetMonitoredResourceDescriptor(ctx context.Context, in *GetMonitoredResourceDescriptorRequest, opts ...grpc.CallOption) (*monitored_resource_descriptor.MonitoredResourceDescriptor, error)
	ListMonitoredResourceDescriptors(ctx context.Context, in *ListMonitoredResourceDescriptorsRequest, opts ...grpc.CallOption) (*ListMonitoredResourceDescriptorsResponse, error)
}

type client struct {
	cc grpc.ClientConnInterface
}

func NewMonitoredResourceDescriptorServiceClient(cc grpc.ClientConnInterface) MonitoredResourceDescriptorServiceClient {
	return &client{cc}
}

func (c *client) BatchGetMonitoredResourceDescriptors(ctx context.Context, in *BatchGetMonitoredResourceDescriptorsRequest, opts ...grpc.CallOption) (*BatchGetMonitoredResourceDescriptorsResponse, error) {
	out := new(BatchGetMonitoredResourceDescriptorsResponse)
	err := c.cc.Invoke(ctx, "/ntt.monitoring.v3.MonitoredResourceDescriptorService/BatchGetMonitoredResourceDescriptors", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) WatchMonitoredResourceDescriptor(ctx context.Context, in *WatchMonitoredResourceDescriptorRequest, opts ...grpc.CallOption) (WatchMonitoredResourceDescriptorClientStream, error) {
	stream, err := c.cc.NewStream(ctx,
		&grpc.StreamDesc{
			StreamName:    "WatchMonitoredResourceDescriptor",
			ServerStreams: true,
		},
		"/ntt.monitoring.v3.MonitoredResourceDescriptorService/WatchMonitoredResourceDescriptor", opts...)
	if err != nil {
		return nil, err
	}
	x := &watchMonitoredResourceDescriptorWatchMonitoredResourceDescriptorClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WatchMonitoredResourceDescriptorClientStream interface {
	Recv() (*WatchMonitoredResourceDescriptorResponse, error)
	grpc.ClientStream
}

type watchMonitoredResourceDescriptorWatchMonitoredResourceDescriptorClient struct {
	grpc.ClientStream
}

func (x *watchMonitoredResourceDescriptorWatchMonitoredResourceDescriptorClient) Recv() (*WatchMonitoredResourceDescriptorResponse, error) {
	m := new(WatchMonitoredResourceDescriptorResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *client) WatchMonitoredResourceDescriptors(ctx context.Context, in *WatchMonitoredResourceDescriptorsRequest, opts ...grpc.CallOption) (WatchMonitoredResourceDescriptorsClientStream, error) {
	stream, err := c.cc.NewStream(ctx,
		&grpc.StreamDesc{
			StreamName:    "WatchMonitoredResourceDescriptors",
			ServerStreams: true,
		},
		"/ntt.monitoring.v3.MonitoredResourceDescriptorService/WatchMonitoredResourceDescriptors", opts...)
	if err != nil {
		return nil, err
	}
	x := &watchMonitoredResourceDescriptorsWatchMonitoredResourceDescriptorsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WatchMonitoredResourceDescriptorsClientStream interface {
	Recv() (*WatchMonitoredResourceDescriptorsResponse, error)
	grpc.ClientStream
}

type watchMonitoredResourceDescriptorsWatchMonitoredResourceDescriptorsClient struct {
	grpc.ClientStream
}

func (x *watchMonitoredResourceDescriptorsWatchMonitoredResourceDescriptorsClient) Recv() (*WatchMonitoredResourceDescriptorsResponse, error) {
	m := new(WatchMonitoredResourceDescriptorsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *client) CreateMonitoredResourceDescriptor(ctx context.Context, in *CreateMonitoredResourceDescriptorRequest, opts ...grpc.CallOption) (*monitored_resource_descriptor.MonitoredResourceDescriptor, error) {
	out := new(monitored_resource_descriptor.MonitoredResourceDescriptor)
	err := c.cc.Invoke(ctx, "/ntt.monitoring.v3.MonitoredResourceDescriptorService/CreateMonitoredResourceDescriptor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) UpdateMonitoredResourceDescriptor(ctx context.Context, in *UpdateMonitoredResourceDescriptorRequest, opts ...grpc.CallOption) (*monitored_resource_descriptor.MonitoredResourceDescriptor, error) {
	out := new(monitored_resource_descriptor.MonitoredResourceDescriptor)
	err := c.cc.Invoke(ctx, "/ntt.monitoring.v3.MonitoredResourceDescriptorService/UpdateMonitoredResourceDescriptor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) DeleteMonitoredResourceDescriptor(ctx context.Context, in *DeleteMonitoredResourceDescriptorRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/ntt.monitoring.v3.MonitoredResourceDescriptorService/DeleteMonitoredResourceDescriptor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) GetMonitoredResourceDescriptor(ctx context.Context, in *GetMonitoredResourceDescriptorRequest, opts ...grpc.CallOption) (*monitored_resource_descriptor.MonitoredResourceDescriptor, error) {
	out := new(monitored_resource_descriptor.MonitoredResourceDescriptor)
	err := c.cc.Invoke(ctx, "/ntt.monitoring.v3.MonitoredResourceDescriptorService/GetMonitoredResourceDescriptor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) ListMonitoredResourceDescriptors(ctx context.Context, in *ListMonitoredResourceDescriptorsRequest, opts ...grpc.CallOption) (*ListMonitoredResourceDescriptorsResponse, error) {
	out := new(ListMonitoredResourceDescriptorsResponse)
	err := c.cc.Invoke(ctx, "/ntt.monitoring.v3.MonitoredResourceDescriptorService/ListMonitoredResourceDescriptors", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
