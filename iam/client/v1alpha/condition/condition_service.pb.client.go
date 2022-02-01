// Code generated by protoc-gen-goten-client
// API: ConditionService
// DO NOT EDIT!!!

package condition_client

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/runtime/protoimpl"
)

// proto imports
import (
	condition "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/condition"
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
	_ = &condition.Condition{}
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

// ConditionServiceClient is the client API for ConditionService.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConditionServiceClient interface {
	GetCondition(ctx context.Context, in *GetConditionRequest, opts ...grpc.CallOption) (*condition.Condition, error)
	BatchGetConditions(ctx context.Context, in *BatchGetConditionsRequest, opts ...grpc.CallOption) (*BatchGetConditionsResponse, error)
	ListConditions(ctx context.Context, in *ListConditionsRequest, opts ...grpc.CallOption) (*ListConditionsResponse, error)
	WatchCondition(ctx context.Context, in *WatchConditionRequest, opts ...grpc.CallOption) (WatchConditionClientStream, error)
	WatchConditions(ctx context.Context, in *WatchConditionsRequest, opts ...grpc.CallOption) (WatchConditionsClientStream, error)
	CreateCondition(ctx context.Context, in *CreateConditionRequest, opts ...grpc.CallOption) (*condition.Condition, error)
	UpdateCondition(ctx context.Context, in *UpdateConditionRequest, opts ...grpc.CallOption) (*condition.Condition, error)
	DeleteCondition(ctx context.Context, in *DeleteConditionRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type client struct {
	cc grpc.ClientConnInterface
}

func NewConditionServiceClient(cc grpc.ClientConnInterface) ConditionServiceClient {
	return &client{cc}
}

func (c *client) GetCondition(ctx context.Context, in *GetConditionRequest, opts ...grpc.CallOption) (*condition.Condition, error) {
	out := new(condition.Condition)
	err := c.cc.Invoke(ctx, "/ntt.iam.v1alpha.ConditionService/GetCondition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) BatchGetConditions(ctx context.Context, in *BatchGetConditionsRequest, opts ...grpc.CallOption) (*BatchGetConditionsResponse, error) {
	out := new(BatchGetConditionsResponse)
	err := c.cc.Invoke(ctx, "/ntt.iam.v1alpha.ConditionService/BatchGetConditions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) ListConditions(ctx context.Context, in *ListConditionsRequest, opts ...grpc.CallOption) (*ListConditionsResponse, error) {
	out := new(ListConditionsResponse)
	err := c.cc.Invoke(ctx, "/ntt.iam.v1alpha.ConditionService/ListConditions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) WatchCondition(ctx context.Context, in *WatchConditionRequest, opts ...grpc.CallOption) (WatchConditionClientStream, error) {
	stream, err := c.cc.NewStream(ctx,
		&grpc.StreamDesc{
			StreamName:    "WatchCondition",
			ServerStreams: true,
		},
		"/ntt.iam.v1alpha.ConditionService/WatchCondition", opts...)
	if err != nil {
		return nil, err
	}
	x := &watchConditionWatchConditionClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WatchConditionClientStream interface {
	Recv() (*WatchConditionResponse, error)
	grpc.ClientStream
}

type watchConditionWatchConditionClient struct {
	grpc.ClientStream
}

func (x *watchConditionWatchConditionClient) Recv() (*WatchConditionResponse, error) {
	m := new(WatchConditionResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *client) WatchConditions(ctx context.Context, in *WatchConditionsRequest, opts ...grpc.CallOption) (WatchConditionsClientStream, error) {
	stream, err := c.cc.NewStream(ctx,
		&grpc.StreamDesc{
			StreamName:    "WatchConditions",
			ServerStreams: true,
		},
		"/ntt.iam.v1alpha.ConditionService/WatchConditions", opts...)
	if err != nil {
		return nil, err
	}
	x := &watchConditionsWatchConditionsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WatchConditionsClientStream interface {
	Recv() (*WatchConditionsResponse, error)
	grpc.ClientStream
}

type watchConditionsWatchConditionsClient struct {
	grpc.ClientStream
}

func (x *watchConditionsWatchConditionsClient) Recv() (*WatchConditionsResponse, error) {
	m := new(WatchConditionsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *client) CreateCondition(ctx context.Context, in *CreateConditionRequest, opts ...grpc.CallOption) (*condition.Condition, error) {
	out := new(condition.Condition)
	err := c.cc.Invoke(ctx, "/ntt.iam.v1alpha.ConditionService/CreateCondition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) UpdateCondition(ctx context.Context, in *UpdateConditionRequest, opts ...grpc.CallOption) (*condition.Condition, error) {
	out := new(condition.Condition)
	err := c.cc.Invoke(ctx, "/ntt.iam.v1alpha.ConditionService/UpdateCondition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) DeleteCondition(ctx context.Context, in *DeleteConditionRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/ntt.iam.v1alpha.ConditionService/DeleteCondition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
