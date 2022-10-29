// Code generated by protoc-gen-goten-client
// API: AuditedResourceDescriptorService
// DO NOT EDIT!!!

package audited_resource_descriptor_client

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/runtime/protoimpl"
)

// proto imports
import (
	audited_resource_descriptor "github.com/cloudwan/edgelq-sdk/audit/resources/v1alpha2/audited_resource_descriptor"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = protoimpl.DescBuilder{}
	_ = new(context.Context)
	_ = grpc.ClientConn{}
)

// make sure we're using proto imports
var (
	_ = &audited_resource_descriptor.AuditedResourceDescriptor{}
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

// AuditedResourceDescriptorServiceClient is the client API for AuditedResourceDescriptorService.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuditedResourceDescriptorServiceClient interface {
	GetAuditedResourceDescriptor(ctx context.Context, in *GetAuditedResourceDescriptorRequest, opts ...grpc.CallOption) (*audited_resource_descriptor.AuditedResourceDescriptor, error)
	BatchGetAuditedResourceDescriptors(ctx context.Context, in *BatchGetAuditedResourceDescriptorsRequest, opts ...grpc.CallOption) (*BatchGetAuditedResourceDescriptorsResponse, error)
	ListAuditedResourceDescriptors(ctx context.Context, in *ListAuditedResourceDescriptorsRequest, opts ...grpc.CallOption) (*ListAuditedResourceDescriptorsResponse, error)
	WatchAuditedResourceDescriptor(ctx context.Context, in *WatchAuditedResourceDescriptorRequest, opts ...grpc.CallOption) (WatchAuditedResourceDescriptorClientStream, error)
	WatchAuditedResourceDescriptors(ctx context.Context, in *WatchAuditedResourceDescriptorsRequest, opts ...grpc.CallOption) (WatchAuditedResourceDescriptorsClientStream, error)
	CreateAuditedResourceDescriptor(ctx context.Context, in *CreateAuditedResourceDescriptorRequest, opts ...grpc.CallOption) (*audited_resource_descriptor.AuditedResourceDescriptor, error)
	UpdateAuditedResourceDescriptor(ctx context.Context, in *UpdateAuditedResourceDescriptorRequest, opts ...grpc.CallOption) (*audited_resource_descriptor.AuditedResourceDescriptor, error)
}

type client struct {
	cc grpc.ClientConnInterface
}

func NewAuditedResourceDescriptorServiceClient(cc grpc.ClientConnInterface) AuditedResourceDescriptorServiceClient {
	return &client{cc}
}

func (c *client) GetAuditedResourceDescriptor(ctx context.Context, in *GetAuditedResourceDescriptorRequest, opts ...grpc.CallOption) (*audited_resource_descriptor.AuditedResourceDescriptor, error) {
	out := new(audited_resource_descriptor.AuditedResourceDescriptor)
	err := c.cc.Invoke(ctx, "/ntt.audit.v1alpha2.AuditedResourceDescriptorService/GetAuditedResourceDescriptor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) BatchGetAuditedResourceDescriptors(ctx context.Context, in *BatchGetAuditedResourceDescriptorsRequest, opts ...grpc.CallOption) (*BatchGetAuditedResourceDescriptorsResponse, error) {
	out := new(BatchGetAuditedResourceDescriptorsResponse)
	err := c.cc.Invoke(ctx, "/ntt.audit.v1alpha2.AuditedResourceDescriptorService/BatchGetAuditedResourceDescriptors", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) ListAuditedResourceDescriptors(ctx context.Context, in *ListAuditedResourceDescriptorsRequest, opts ...grpc.CallOption) (*ListAuditedResourceDescriptorsResponse, error) {
	out := new(ListAuditedResourceDescriptorsResponse)
	err := c.cc.Invoke(ctx, "/ntt.audit.v1alpha2.AuditedResourceDescriptorService/ListAuditedResourceDescriptors", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) WatchAuditedResourceDescriptor(ctx context.Context, in *WatchAuditedResourceDescriptorRequest, opts ...grpc.CallOption) (WatchAuditedResourceDescriptorClientStream, error) {
	stream, err := c.cc.NewStream(ctx,
		&grpc.StreamDesc{
			StreamName:    "WatchAuditedResourceDescriptor",
			ServerStreams: true,
		},
		"/ntt.audit.v1alpha2.AuditedResourceDescriptorService/WatchAuditedResourceDescriptor", opts...)
	if err != nil {
		return nil, err
	}
	x := &watchAuditedResourceDescriptorWatchAuditedResourceDescriptorClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WatchAuditedResourceDescriptorClientStream interface {
	Recv() (*WatchAuditedResourceDescriptorResponse, error)
	grpc.ClientStream
}

type watchAuditedResourceDescriptorWatchAuditedResourceDescriptorClient struct {
	grpc.ClientStream
}

func (x *watchAuditedResourceDescriptorWatchAuditedResourceDescriptorClient) Recv() (*WatchAuditedResourceDescriptorResponse, error) {
	m := new(WatchAuditedResourceDescriptorResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *client) WatchAuditedResourceDescriptors(ctx context.Context, in *WatchAuditedResourceDescriptorsRequest, opts ...grpc.CallOption) (WatchAuditedResourceDescriptorsClientStream, error) {
	stream, err := c.cc.NewStream(ctx,
		&grpc.StreamDesc{
			StreamName:    "WatchAuditedResourceDescriptors",
			ServerStreams: true,
		},
		"/ntt.audit.v1alpha2.AuditedResourceDescriptorService/WatchAuditedResourceDescriptors", opts...)
	if err != nil {
		return nil, err
	}
	x := &watchAuditedResourceDescriptorsWatchAuditedResourceDescriptorsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WatchAuditedResourceDescriptorsClientStream interface {
	Recv() (*WatchAuditedResourceDescriptorsResponse, error)
	grpc.ClientStream
}

type watchAuditedResourceDescriptorsWatchAuditedResourceDescriptorsClient struct {
	grpc.ClientStream
}

func (x *watchAuditedResourceDescriptorsWatchAuditedResourceDescriptorsClient) Recv() (*WatchAuditedResourceDescriptorsResponse, error) {
	m := new(WatchAuditedResourceDescriptorsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *client) CreateAuditedResourceDescriptor(ctx context.Context, in *CreateAuditedResourceDescriptorRequest, opts ...grpc.CallOption) (*audited_resource_descriptor.AuditedResourceDescriptor, error) {
	out := new(audited_resource_descriptor.AuditedResourceDescriptor)
	err := c.cc.Invoke(ctx, "/ntt.audit.v1alpha2.AuditedResourceDescriptorService/CreateAuditedResourceDescriptor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) UpdateAuditedResourceDescriptor(ctx context.Context, in *UpdateAuditedResourceDescriptorRequest, opts ...grpc.CallOption) (*audited_resource_descriptor.AuditedResourceDescriptor, error) {
	out := new(audited_resource_descriptor.AuditedResourceDescriptor)
	err := c.cc.Invoke(ctx, "/ntt.audit.v1alpha2.AuditedResourceDescriptorService/UpdateAuditedResourceDescriptor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
