// Code generated by protoc-gen-goten-client
// API: HardwareService
// DO NOT EDIT!!!

package hardware_client

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/runtime/protoimpl"
)

// proto imports
import (
	hardware "github.com/cloudwan/edgelq-sdk/ztp/resources/v1/hardware"
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
	_ = &hardware.Hardware{}
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

// HardwareServiceClient is the client API for HardwareService.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HardwareServiceClient interface {
	GetHardware(ctx context.Context, in *GetHardwareRequest, opts ...grpc.CallOption) (*hardware.Hardware, error)
	BatchGetHardwares(ctx context.Context, in *BatchGetHardwaresRequest, opts ...grpc.CallOption) (*BatchGetHardwaresResponse, error)
	ListHardwares(ctx context.Context, in *ListHardwaresRequest, opts ...grpc.CallOption) (*ListHardwaresResponse, error)
	WatchHardware(ctx context.Context, in *WatchHardwareRequest, opts ...grpc.CallOption) (WatchHardwareClientStream, error)
	WatchHardwares(ctx context.Context, in *WatchHardwaresRequest, opts ...grpc.CallOption) (WatchHardwaresClientStream, error)
	CreateHardware(ctx context.Context, in *CreateHardwareRequest, opts ...grpc.CallOption) (*hardware.Hardware, error)
	UpdateHardware(ctx context.Context, in *UpdateHardwareRequest, opts ...grpc.CallOption) (*hardware.Hardware, error)
	DeleteHardware(ctx context.Context, in *DeleteHardwareRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type client struct {
	cc grpc.ClientConnInterface
}

func NewHardwareServiceClient(cc grpc.ClientConnInterface) HardwareServiceClient {
	return &client{cc}
}

func (c *client) GetHardware(ctx context.Context, in *GetHardwareRequest, opts ...grpc.CallOption) (*hardware.Hardware, error) {
	out := new(hardware.Hardware)
	err := c.cc.Invoke(ctx, "/ntt.ztp.v1.HardwareService/GetHardware", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) BatchGetHardwares(ctx context.Context, in *BatchGetHardwaresRequest, opts ...grpc.CallOption) (*BatchGetHardwaresResponse, error) {
	out := new(BatchGetHardwaresResponse)
	err := c.cc.Invoke(ctx, "/ntt.ztp.v1.HardwareService/BatchGetHardwares", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) ListHardwares(ctx context.Context, in *ListHardwaresRequest, opts ...grpc.CallOption) (*ListHardwaresResponse, error) {
	out := new(ListHardwaresResponse)
	err := c.cc.Invoke(ctx, "/ntt.ztp.v1.HardwareService/ListHardwares", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) WatchHardware(ctx context.Context, in *WatchHardwareRequest, opts ...grpc.CallOption) (WatchHardwareClientStream, error) {
	stream, err := c.cc.NewStream(ctx,
		&grpc.StreamDesc{
			StreamName:    "WatchHardware",
			ServerStreams: true,
		},
		"/ntt.ztp.v1.HardwareService/WatchHardware", opts...)
	if err != nil {
		return nil, err
	}
	x := &watchHardwareWatchHardwareClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WatchHardwareClientStream interface {
	Recv() (*WatchHardwareResponse, error)
	grpc.ClientStream
}

type watchHardwareWatchHardwareClient struct {
	grpc.ClientStream
}

func (x *watchHardwareWatchHardwareClient) Recv() (*WatchHardwareResponse, error) {
	m := new(WatchHardwareResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *client) WatchHardwares(ctx context.Context, in *WatchHardwaresRequest, opts ...grpc.CallOption) (WatchHardwaresClientStream, error) {
	stream, err := c.cc.NewStream(ctx,
		&grpc.StreamDesc{
			StreamName:    "WatchHardwares",
			ServerStreams: true,
		},
		"/ntt.ztp.v1.HardwareService/WatchHardwares", opts...)
	if err != nil {
		return nil, err
	}
	x := &watchHardwaresWatchHardwaresClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WatchHardwaresClientStream interface {
	Recv() (*WatchHardwaresResponse, error)
	grpc.ClientStream
}

type watchHardwaresWatchHardwaresClient struct {
	grpc.ClientStream
}

func (x *watchHardwaresWatchHardwaresClient) Recv() (*WatchHardwaresResponse, error) {
	m := new(WatchHardwaresResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *client) CreateHardware(ctx context.Context, in *CreateHardwareRequest, opts ...grpc.CallOption) (*hardware.Hardware, error) {
	out := new(hardware.Hardware)
	err := c.cc.Invoke(ctx, "/ntt.ztp.v1.HardwareService/CreateHardware", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) UpdateHardware(ctx context.Context, in *UpdateHardwareRequest, opts ...grpc.CallOption) (*hardware.Hardware, error) {
	out := new(hardware.Hardware)
	err := c.cc.Invoke(ctx, "/ntt.ztp.v1.HardwareService/UpdateHardware", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) DeleteHardware(ctx context.Context, in *DeleteHardwareRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ntt.ztp.v1.HardwareService/DeleteHardware", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
