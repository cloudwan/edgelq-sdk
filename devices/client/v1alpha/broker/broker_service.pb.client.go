// Code generated by protoc-gen-goten-client
// API: BrokerService
// DO NOT EDIT!!!

package broker_client

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/runtime/protoimpl"
)

// proto imports
import ()

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = protoimpl.DescBuilder{}
	_ = context.Context(nil)
	_ = grpc.ClientConn{}
)

// make sure we're using proto imports
var ()

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

// BrokerServiceClient is the client API for BrokerService.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BrokerServiceClient interface {
	ListenForConnections(ctx context.Context, opts ...grpc.CallOption) (ListenForConnectionsClientStream, error)
	OpenConnectionChannelSocket(ctx context.Context, opts ...grpc.CallOption) (OpenConnectionChannelSocketClientStream, error)
	ConnectToDevice(ctx context.Context, opts ...grpc.CallOption) (ConnectToDeviceClientStream, error)
}

type client struct {
	cc grpc.ClientConnInterface
}

func NewBrokerServiceClient(cc grpc.ClientConnInterface) BrokerServiceClient {
	return &client{cc}
}

func (c *client) ListenForConnections(ctx context.Context, opts ...grpc.CallOption) (ListenForConnectionsClientStream, error) {
	stream, err := c.cc.NewStream(ctx,
		&grpc.StreamDesc{
			StreamName:    "ListenForConnections",
			ServerStreams: true,
			ClientStreams: true,
		},
		"/ntt.devices.v1alpha.BrokerService/ListenForConnections", opts...)
	if err != nil {
		return nil, err
	}
	x := &listenForConnectionsListenForConnectionsClient{stream}
	return x, nil
}

type ListenForConnectionsClientStream interface {
	Send(*ListenForConnectionsRequest) error
	Recv() (*ListenForConnectionsResponse, error)
	grpc.ClientStream
}

type listenForConnectionsListenForConnectionsClient struct {
	grpc.ClientStream
}

func (x *listenForConnectionsListenForConnectionsClient) Send(m *ListenForConnectionsRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *listenForConnectionsListenForConnectionsClient) Recv() (*ListenForConnectionsResponse, error) {
	m := new(ListenForConnectionsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *client) OpenConnectionChannelSocket(ctx context.Context, opts ...grpc.CallOption) (OpenConnectionChannelSocketClientStream, error) {
	stream, err := c.cc.NewStream(ctx,
		&grpc.StreamDesc{
			StreamName:    "OpenConnectionChannelSocket",
			ServerStreams: true,
			ClientStreams: true,
		},
		"/ntt.devices.v1alpha.BrokerService/OpenConnectionChannelSocket", opts...)
	if err != nil {
		return nil, err
	}
	x := &openConnectionChannelSocketOpenConnectionChannelSocketClient{stream}
	return x, nil
}

type OpenConnectionChannelSocketClientStream interface {
	Send(*OpenConnectionChannelSocketRequest) error
	Recv() (*OpenConnectionChannelSocketResponse, error)
	grpc.ClientStream
}

type openConnectionChannelSocketOpenConnectionChannelSocketClient struct {
	grpc.ClientStream
}

func (x *openConnectionChannelSocketOpenConnectionChannelSocketClient) Send(m *OpenConnectionChannelSocketRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *openConnectionChannelSocketOpenConnectionChannelSocketClient) Recv() (*OpenConnectionChannelSocketResponse, error) {
	m := new(OpenConnectionChannelSocketResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *client) ConnectToDevice(ctx context.Context, opts ...grpc.CallOption) (ConnectToDeviceClientStream, error) {
	stream, err := c.cc.NewStream(ctx,
		&grpc.StreamDesc{
			StreamName:    "ConnectToDevice",
			ServerStreams: true,
			ClientStreams: true,
		},
		"/ntt.devices.v1alpha.BrokerService/ConnectToDevice", opts...)
	if err != nil {
		return nil, err
	}
	x := &connectToDeviceConnectToDeviceClient{stream}
	return x, nil
}

type ConnectToDeviceClientStream interface {
	Send(*ConnectToDeviceRequest) error
	Recv() (*ConnectToDeviceResponse, error)
	grpc.ClientStream
}

type connectToDeviceConnectToDeviceClient struct {
	grpc.ClientStream
}

func (x *connectToDeviceConnectToDeviceClient) Send(m *ConnectToDeviceRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *connectToDeviceConnectToDeviceClient) Recv() (*ConnectToDeviceResponse, error) {
	m := new(ConnectToDeviceResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}