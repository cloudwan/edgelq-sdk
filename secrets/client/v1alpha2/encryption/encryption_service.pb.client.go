// Code generated by protoc-gen-goten-client
// API: EncryptionService
// DO NOT EDIT!!!

package encryption_client

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/runtime/protoimpl"
)

// proto imports
import (
	project "github.com/cloudwan/edgelq-sdk/secrets/resources/v1alpha2/project"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = protoimpl.DescBuilder{}
	_ = context.Context(nil)
	_ = grpc.ClientConn{}
)

// make sure we're using proto imports
var (
	_ = &project.Project{}
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

// EncryptionServiceClient is the client API for EncryptionService.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EncryptionServiceClient interface {
	EncryptData(ctx context.Context, in *EncryptDataRequest, opts ...grpc.CallOption) (*EncryptDataResponse, error)
	DecryptData(ctx context.Context, in *DecryptDataRequest, opts ...grpc.CallOption) (*DecryptDataResponse, error)
}

type client struct {
	cc grpc.ClientConnInterface
}

func NewEncryptionServiceClient(cc grpc.ClientConnInterface) EncryptionServiceClient {
	return &client{cc}
}

func (c *client) EncryptData(ctx context.Context, in *EncryptDataRequest, opts ...grpc.CallOption) (*EncryptDataResponse, error) {
	out := new(EncryptDataResponse)
	err := c.cc.Invoke(ctx, "/ntt.secrets.v1alpha2.EncryptionService/EncryptData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) DecryptData(ctx context.Context, in *DecryptDataRequest, opts ...grpc.CallOption) (*DecryptDataResponse, error) {
	out := new(DecryptDataResponse)
	err := c.cc.Invoke(ctx, "/ntt.secrets.v1alpha2.EncryptionService/DecryptData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
