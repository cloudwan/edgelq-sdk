// Code generated by protoc-gen-goten-client
// API: DeviceTypeService
// DO NOT EDIT!!!

package device_type_client

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/runtime/protoimpl"
)

// proto imports
import (
	device_type "github.com/cloudwan/edgelq-sdk/devices/resources/v1/device_type"
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
	_ = &device_type.DeviceType{}
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

// DeviceTypeServiceClient is the client API for DeviceTypeService.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DeviceTypeServiceClient interface {
	GetDeviceType(ctx context.Context, in *GetDeviceTypeRequest, opts ...grpc.CallOption) (*device_type.DeviceType, error)
	BatchGetDeviceTypes(ctx context.Context, in *BatchGetDeviceTypesRequest, opts ...grpc.CallOption) (*BatchGetDeviceTypesResponse, error)
	ListDeviceTypes(ctx context.Context, in *ListDeviceTypesRequest, opts ...grpc.CallOption) (*ListDeviceTypesResponse, error)
	WatchDeviceType(ctx context.Context, in *WatchDeviceTypeRequest, opts ...grpc.CallOption) (WatchDeviceTypeClientStream, error)
	WatchDeviceTypes(ctx context.Context, in *WatchDeviceTypesRequest, opts ...grpc.CallOption) (WatchDeviceTypesClientStream, error)
	CreateDeviceType(ctx context.Context, in *CreateDeviceTypeRequest, opts ...grpc.CallOption) (*device_type.DeviceType, error)
	UpdateDeviceType(ctx context.Context, in *UpdateDeviceTypeRequest, opts ...grpc.CallOption) (*device_type.DeviceType, error)
	DeleteDeviceType(ctx context.Context, in *DeleteDeviceTypeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type client struct {
	cc grpc.ClientConnInterface
}

func NewDeviceTypeServiceClient(cc grpc.ClientConnInterface) DeviceTypeServiceClient {
	return &client{cc}
}

func (c *client) GetDeviceType(ctx context.Context, in *GetDeviceTypeRequest, opts ...grpc.CallOption) (*device_type.DeviceType, error) {
	out := new(device_type.DeviceType)
	err := c.cc.Invoke(ctx, "/ntt.devices.v1.DeviceTypeService/GetDeviceType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) BatchGetDeviceTypes(ctx context.Context, in *BatchGetDeviceTypesRequest, opts ...grpc.CallOption) (*BatchGetDeviceTypesResponse, error) {
	out := new(BatchGetDeviceTypesResponse)
	err := c.cc.Invoke(ctx, "/ntt.devices.v1.DeviceTypeService/BatchGetDeviceTypes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) ListDeviceTypes(ctx context.Context, in *ListDeviceTypesRequest, opts ...grpc.CallOption) (*ListDeviceTypesResponse, error) {
	out := new(ListDeviceTypesResponse)
	err := c.cc.Invoke(ctx, "/ntt.devices.v1.DeviceTypeService/ListDeviceTypes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) WatchDeviceType(ctx context.Context, in *WatchDeviceTypeRequest, opts ...grpc.CallOption) (WatchDeviceTypeClientStream, error) {
	stream, err := c.cc.NewStream(ctx,
		&grpc.StreamDesc{
			StreamName:    "WatchDeviceType",
			ServerStreams: true,
		},
		"/ntt.devices.v1.DeviceTypeService/WatchDeviceType", opts...)
	if err != nil {
		return nil, err
	}
	x := &watchDeviceTypeWatchDeviceTypeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WatchDeviceTypeClientStream interface {
	Recv() (*WatchDeviceTypeResponse, error)
	grpc.ClientStream
}

type watchDeviceTypeWatchDeviceTypeClient struct {
	grpc.ClientStream
}

func (x *watchDeviceTypeWatchDeviceTypeClient) Recv() (*WatchDeviceTypeResponse, error) {
	m := new(WatchDeviceTypeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *client) WatchDeviceTypes(ctx context.Context, in *WatchDeviceTypesRequest, opts ...grpc.CallOption) (WatchDeviceTypesClientStream, error) {
	stream, err := c.cc.NewStream(ctx,
		&grpc.StreamDesc{
			StreamName:    "WatchDeviceTypes",
			ServerStreams: true,
		},
		"/ntt.devices.v1.DeviceTypeService/WatchDeviceTypes", opts...)
	if err != nil {
		return nil, err
	}
	x := &watchDeviceTypesWatchDeviceTypesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WatchDeviceTypesClientStream interface {
	Recv() (*WatchDeviceTypesResponse, error)
	grpc.ClientStream
}

type watchDeviceTypesWatchDeviceTypesClient struct {
	grpc.ClientStream
}

func (x *watchDeviceTypesWatchDeviceTypesClient) Recv() (*WatchDeviceTypesResponse, error) {
	m := new(WatchDeviceTypesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *client) CreateDeviceType(ctx context.Context, in *CreateDeviceTypeRequest, opts ...grpc.CallOption) (*device_type.DeviceType, error) {
	out := new(device_type.DeviceType)
	err := c.cc.Invoke(ctx, "/ntt.devices.v1.DeviceTypeService/CreateDeviceType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) UpdateDeviceType(ctx context.Context, in *UpdateDeviceTypeRequest, opts ...grpc.CallOption) (*device_type.DeviceType, error) {
	out := new(device_type.DeviceType)
	err := c.cc.Invoke(ctx, "/ntt.devices.v1.DeviceTypeService/UpdateDeviceType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) DeleteDeviceType(ctx context.Context, in *DeleteDeviceTypeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ntt.devices.v1.DeviceTypeService/DeleteDeviceType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
