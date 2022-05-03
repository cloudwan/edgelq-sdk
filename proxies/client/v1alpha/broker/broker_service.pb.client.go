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
import (
	project "github.com/cloudwan/edgelq-sdk/proxies/resources/v1alpha/project"
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

// BrokerServiceClient is the client API for BrokerService.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BrokerServiceClient interface {
	Connect(ctx context.Context, opts ...grpc.CallOption) (ConnectClientStream, error)
	Listen(ctx context.Context, opts ...grpc.CallOption) (ListenClientStream, error)
	Accept(ctx context.Context, opts ...grpc.CallOption) (AcceptClientStream, error)
}

type client struct {
	cc grpc.ClientConnInterface
}

func NewBrokerServiceClient(cc grpc.ClientConnInterface) BrokerServiceClient {
	return &client{cc}
}

func (c *client) Connect(ctx context.Context, opts ...grpc.CallOption) (ConnectClientStream, error) {
	stream, err := c.cc.NewStream(ctx,
		&grpc.StreamDesc{
			StreamName:    "Connect",
			ServerStreams: true,
			ClientStreams: true,
		},
		"/ntt.proxies.v1alpha.BrokerService/Connect", opts...)
	if err != nil {
		return nil, err
	}
	x := &connectConnectClient{stream}
	return x, nil
}

type ConnectClientStream interface {
	Send(*ConnectRequest) error
	Recv() (*ConnectResponse, error)
	grpc.ClientStream
}

type connectConnectClient struct {
	grpc.ClientStream
}

func (x *connectConnectClient) Send(m *ConnectRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *connectConnectClient) Recv() (*ConnectResponse, error) {
	m := new(ConnectResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *client) Listen(ctx context.Context, opts ...grpc.CallOption) (ListenClientStream, error) {
	stream, err := c.cc.NewStream(ctx,
		&grpc.StreamDesc{
			StreamName:    "Listen",
			ServerStreams: true,
			ClientStreams: true,
		},
		"/ntt.proxies.v1alpha.BrokerService/Listen", opts...)
	if err != nil {
		return nil, err
	}
	x := &listenListenClient{stream}
	return x, nil
}

type ListenClientStream interface {
	Send(*ListenRequest) error
	Recv() (*ListenResponse, error)
	grpc.ClientStream
}

type listenListenClient struct {
	grpc.ClientStream
}

func (x *listenListenClient) Send(m *ListenRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *listenListenClient) Recv() (*ListenResponse, error) {
	m := new(ListenResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *client) Accept(ctx context.Context, opts ...grpc.CallOption) (AcceptClientStream, error) {
	stream, err := c.cc.NewStream(ctx,
		&grpc.StreamDesc{
			StreamName:    "Accept",
			ServerStreams: true,
			ClientStreams: true,
		},
		"/ntt.proxies.v1alpha.BrokerService/Accept", opts...)
	if err != nil {
		return nil, err
	}
	x := &acceptAcceptClient{stream}
	return x, nil
}

type AcceptClientStream interface {
	Send(*AcceptRequest) error
	Recv() (*AcceptResponse, error)
	grpc.ClientStream
}

type acceptAcceptClient struct {
	grpc.ClientStream
}

func (x *acceptAcceptClient) Send(m *AcceptRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *acceptAcceptClient) Recv() (*AcceptResponse, error) {
	m := new(AcceptResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}
