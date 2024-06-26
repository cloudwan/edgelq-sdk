// Code generated by protoc-gen-goten-client
// API: PublicService
// DO NOT EDIT!!!

package public_client

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/runtime/protoimpl"
)

// proto imports
import (
	device "github.com/cloudwan/edgelq-sdk/devices/resources/v1alpha2/device"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = protoimpl.DescBuilder{}
	_ = new(context.Context)
	_ = grpc.ClientConn{}
)

// make sure we're using proto imports
var (
	_ = &device.Device{}
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

// PublicServiceClient is the client API for PublicService.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PublicServiceClient interface {
	ListPublicDevices(ctx context.Context, in *ListPublicDevicesRequest, opts ...grpc.CallOption) (*ListPublicDevicesResponse, error)
}

type client struct {
	cc grpc.ClientConnInterface
}

func NewPublicServiceClient(cc grpc.ClientConnInterface) PublicServiceClient {
	return &client{cc}
}

func (c *client) ListPublicDevices(ctx context.Context, in *ListPublicDevicesRequest, opts ...grpc.CallOption) (*ListPublicDevicesResponse, error) {
	out := new(ListPublicDevicesResponse)
	err := c.cc.Invoke(ctx, "/ntt.devices.v1alpha2.PublicService/ListPublicDevices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
