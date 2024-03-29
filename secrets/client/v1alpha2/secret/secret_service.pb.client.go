// Code generated by protoc-gen-goten-client
// API: SecretService
// DO NOT EDIT!!!

package secret_client

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/runtime/protoimpl"
)

// proto imports
import (
	secret "github.com/cloudwan/edgelq-sdk/secrets/resources/v1alpha2/secret"
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
	_ = &secret.Secret{}
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

// SecretServiceClient is the client API for SecretService.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SecretServiceClient interface {
	GetSecret(ctx context.Context, in *GetSecretRequest, opts ...grpc.CallOption) (*secret.Secret, error)
	BatchGetSecrets(ctx context.Context, in *BatchGetSecretsRequest, opts ...grpc.CallOption) (*BatchGetSecretsResponse, error)
	ListSecrets(ctx context.Context, in *ListSecretsRequest, opts ...grpc.CallOption) (*ListSecretsResponse, error)
	WatchSecret(ctx context.Context, in *WatchSecretRequest, opts ...grpc.CallOption) (WatchSecretClientStream, error)
	WatchSecrets(ctx context.Context, in *WatchSecretsRequest, opts ...grpc.CallOption) (WatchSecretsClientStream, error)
	CreateSecret(ctx context.Context, in *CreateSecretRequest, opts ...grpc.CallOption) (*secret.Secret, error)
	UpdateSecret(ctx context.Context, in *UpdateSecretRequest, opts ...grpc.CallOption) (*secret.Secret, error)
	DeleteSecret(ctx context.Context, in *DeleteSecretRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type client struct {
	cc grpc.ClientConnInterface
}

func NewSecretServiceClient(cc grpc.ClientConnInterface) SecretServiceClient {
	return &client{cc}
}

func (c *client) GetSecret(ctx context.Context, in *GetSecretRequest, opts ...grpc.CallOption) (*secret.Secret, error) {
	out := new(secret.Secret)
	err := c.cc.Invoke(ctx, "/ntt.secrets.v1alpha2.SecretService/GetSecret", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) BatchGetSecrets(ctx context.Context, in *BatchGetSecretsRequest, opts ...grpc.CallOption) (*BatchGetSecretsResponse, error) {
	out := new(BatchGetSecretsResponse)
	err := c.cc.Invoke(ctx, "/ntt.secrets.v1alpha2.SecretService/BatchGetSecrets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) ListSecrets(ctx context.Context, in *ListSecretsRequest, opts ...grpc.CallOption) (*ListSecretsResponse, error) {
	out := new(ListSecretsResponse)
	err := c.cc.Invoke(ctx, "/ntt.secrets.v1alpha2.SecretService/ListSecrets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) WatchSecret(ctx context.Context, in *WatchSecretRequest, opts ...grpc.CallOption) (WatchSecretClientStream, error) {
	stream, err := c.cc.NewStream(ctx,
		&grpc.StreamDesc{
			StreamName:    "WatchSecret",
			ServerStreams: true,
		},
		"/ntt.secrets.v1alpha2.SecretService/WatchSecret", opts...)
	if err != nil {
		return nil, err
	}
	x := &watchSecretWatchSecretClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WatchSecretClientStream interface {
	Recv() (*WatchSecretResponse, error)
	grpc.ClientStream
}

type watchSecretWatchSecretClient struct {
	grpc.ClientStream
}

func (x *watchSecretWatchSecretClient) Recv() (*WatchSecretResponse, error) {
	m := new(WatchSecretResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *client) WatchSecrets(ctx context.Context, in *WatchSecretsRequest, opts ...grpc.CallOption) (WatchSecretsClientStream, error) {
	stream, err := c.cc.NewStream(ctx,
		&grpc.StreamDesc{
			StreamName:    "WatchSecrets",
			ServerStreams: true,
		},
		"/ntt.secrets.v1alpha2.SecretService/WatchSecrets", opts...)
	if err != nil {
		return nil, err
	}
	x := &watchSecretsWatchSecretsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WatchSecretsClientStream interface {
	Recv() (*WatchSecretsResponse, error)
	grpc.ClientStream
}

type watchSecretsWatchSecretsClient struct {
	grpc.ClientStream
}

func (x *watchSecretsWatchSecretsClient) Recv() (*WatchSecretsResponse, error) {
	m := new(WatchSecretsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *client) CreateSecret(ctx context.Context, in *CreateSecretRequest, opts ...grpc.CallOption) (*secret.Secret, error) {
	out := new(secret.Secret)
	err := c.cc.Invoke(ctx, "/ntt.secrets.v1alpha2.SecretService/CreateSecret", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) UpdateSecret(ctx context.Context, in *UpdateSecretRequest, opts ...grpc.CallOption) (*secret.Secret, error) {
	out := new(secret.Secret)
	err := c.cc.Invoke(ctx, "/ntt.secrets.v1alpha2.SecretService/UpdateSecret", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) DeleteSecret(ctx context.Context, in *DeleteSecretRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ntt.secrets.v1alpha2.SecretService/DeleteSecret", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
