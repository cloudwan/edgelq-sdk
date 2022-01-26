// Code generated by protoc-gen-goten-client
// API: ServiceAccountKeyService
// DO NOT EDIT!!!

package service_account_key_client

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/runtime/protoimpl"
)

// proto imports
import (
	service_account_key "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/service_account_key"
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
	_ = &service_account_key.ServiceAccountKey{}
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

// ServiceAccountKeyServiceClient is the client API for ServiceAccountKeyService.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServiceAccountKeyServiceClient interface {
	GetServiceAccountKey(ctx context.Context, in *GetServiceAccountKeyRequest, opts ...grpc.CallOption) (*service_account_key.ServiceAccountKey, error)
	BatchGetServiceAccountKeys(ctx context.Context, in *BatchGetServiceAccountKeysRequest, opts ...grpc.CallOption) (*BatchGetServiceAccountKeysResponse, error)
	ListServiceAccountKeys(ctx context.Context, in *ListServiceAccountKeysRequest, opts ...grpc.CallOption) (*ListServiceAccountKeysResponse, error)
	WatchServiceAccountKey(ctx context.Context, in *WatchServiceAccountKeyRequest, opts ...grpc.CallOption) (WatchServiceAccountKeyClientStream, error)
	WatchServiceAccountKeys(ctx context.Context, in *WatchServiceAccountKeysRequest, opts ...grpc.CallOption) (WatchServiceAccountKeysClientStream, error)
	CreateServiceAccountKey(ctx context.Context, in *CreateServiceAccountKeyRequest, opts ...grpc.CallOption) (*service_account_key.ServiceAccountKey, error)
	UpdateServiceAccountKey(ctx context.Context, in *UpdateServiceAccountKeyRequest, opts ...grpc.CallOption) (*service_account_key.ServiceAccountKey, error)
	DeleteServiceAccountKey(ctx context.Context, in *DeleteServiceAccountKeyRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type client struct {
	cc grpc.ClientConnInterface
}

func NewServiceAccountKeyServiceClient(cc grpc.ClientConnInterface) ServiceAccountKeyServiceClient {
	return &client{cc}
}

func (c *client) GetServiceAccountKey(ctx context.Context, in *GetServiceAccountKeyRequest, opts ...grpc.CallOption) (*service_account_key.ServiceAccountKey, error) {
	out := new(service_account_key.ServiceAccountKey)
	err := c.cc.Invoke(ctx, "/ntt.iam.v1alpha2.ServiceAccountKeyService/GetServiceAccountKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) BatchGetServiceAccountKeys(ctx context.Context, in *BatchGetServiceAccountKeysRequest, opts ...grpc.CallOption) (*BatchGetServiceAccountKeysResponse, error) {
	out := new(BatchGetServiceAccountKeysResponse)
	err := c.cc.Invoke(ctx, "/ntt.iam.v1alpha2.ServiceAccountKeyService/BatchGetServiceAccountKeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) ListServiceAccountKeys(ctx context.Context, in *ListServiceAccountKeysRequest, opts ...grpc.CallOption) (*ListServiceAccountKeysResponse, error) {
	out := new(ListServiceAccountKeysResponse)
	err := c.cc.Invoke(ctx, "/ntt.iam.v1alpha2.ServiceAccountKeyService/ListServiceAccountKeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) WatchServiceAccountKey(ctx context.Context, in *WatchServiceAccountKeyRequest, opts ...grpc.CallOption) (WatchServiceAccountKeyClientStream, error) {
	stream, err := c.cc.NewStream(ctx,
		&grpc.StreamDesc{
			StreamName:    "WatchServiceAccountKey",
			ServerStreams: true,
		},
		"/ntt.iam.v1alpha2.ServiceAccountKeyService/WatchServiceAccountKey", opts...)
	if err != nil {
		return nil, err
	}
	x := &watchServiceAccountKeyWatchServiceAccountKeyClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WatchServiceAccountKeyClientStream interface {
	Recv() (*WatchServiceAccountKeyResponse, error)
	grpc.ClientStream
}

type watchServiceAccountKeyWatchServiceAccountKeyClient struct {
	grpc.ClientStream
}

func (x *watchServiceAccountKeyWatchServiceAccountKeyClient) Recv() (*WatchServiceAccountKeyResponse, error) {
	m := new(WatchServiceAccountKeyResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *client) WatchServiceAccountKeys(ctx context.Context, in *WatchServiceAccountKeysRequest, opts ...grpc.CallOption) (WatchServiceAccountKeysClientStream, error) {
	stream, err := c.cc.NewStream(ctx,
		&grpc.StreamDesc{
			StreamName:    "WatchServiceAccountKeys",
			ServerStreams: true,
		},
		"/ntt.iam.v1alpha2.ServiceAccountKeyService/WatchServiceAccountKeys", opts...)
	if err != nil {
		return nil, err
	}
	x := &watchServiceAccountKeysWatchServiceAccountKeysClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WatchServiceAccountKeysClientStream interface {
	Recv() (*WatchServiceAccountKeysResponse, error)
	grpc.ClientStream
}

type watchServiceAccountKeysWatchServiceAccountKeysClient struct {
	grpc.ClientStream
}

func (x *watchServiceAccountKeysWatchServiceAccountKeysClient) Recv() (*WatchServiceAccountKeysResponse, error) {
	m := new(WatchServiceAccountKeysResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *client) CreateServiceAccountKey(ctx context.Context, in *CreateServiceAccountKeyRequest, opts ...grpc.CallOption) (*service_account_key.ServiceAccountKey, error) {
	out := new(service_account_key.ServiceAccountKey)
	err := c.cc.Invoke(ctx, "/ntt.iam.v1alpha2.ServiceAccountKeyService/CreateServiceAccountKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) UpdateServiceAccountKey(ctx context.Context, in *UpdateServiceAccountKeyRequest, opts ...grpc.CallOption) (*service_account_key.ServiceAccountKey, error) {
	out := new(service_account_key.ServiceAccountKey)
	err := c.cc.Invoke(ctx, "/ntt.iam.v1alpha2.ServiceAccountKeyService/UpdateServiceAccountKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) DeleteServiceAccountKey(ctx context.Context, in *DeleteServiceAccountKeyRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/ntt.iam.v1alpha2.ServiceAccountKeyService/DeleteServiceAccountKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
