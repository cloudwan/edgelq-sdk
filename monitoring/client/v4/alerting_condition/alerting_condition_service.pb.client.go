// Code generated by protoc-gen-goten-client
// API: AlertingConditionService
// DO NOT EDIT!!!

package alerting_condition_client

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/runtime/protoimpl"
)

// proto imports
import (
	alerting_condition "github.com/cloudwan/edgelq-sdk/monitoring/resources/v4/alerting_condition"
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
	_ = &alerting_condition.AlertingCondition{}
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

// AlertingConditionServiceClient is the client API for AlertingConditionService.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AlertingConditionServiceClient interface {
	GetAlertingCondition(ctx context.Context, in *GetAlertingConditionRequest, opts ...grpc.CallOption) (*alerting_condition.AlertingCondition, error)
	BatchGetAlertingConditions(ctx context.Context, in *BatchGetAlertingConditionsRequest, opts ...grpc.CallOption) (*BatchGetAlertingConditionsResponse, error)
	ListAlertingConditions(ctx context.Context, in *ListAlertingConditionsRequest, opts ...grpc.CallOption) (*ListAlertingConditionsResponse, error)
	WatchAlertingCondition(ctx context.Context, in *WatchAlertingConditionRequest, opts ...grpc.CallOption) (WatchAlertingConditionClientStream, error)
	WatchAlertingConditions(ctx context.Context, in *WatchAlertingConditionsRequest, opts ...grpc.CallOption) (WatchAlertingConditionsClientStream, error)
	CreateAlertingCondition(ctx context.Context, in *CreateAlertingConditionRequest, opts ...grpc.CallOption) (*alerting_condition.AlertingCondition, error)
	UpdateAlertingCondition(ctx context.Context, in *UpdateAlertingConditionRequest, opts ...grpc.CallOption) (*alerting_condition.AlertingCondition, error)
	DeleteAlertingCondition(ctx context.Context, in *DeleteAlertingConditionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	SearchAlertingConditions(ctx context.Context, in *SearchAlertingConditionsRequest, opts ...grpc.CallOption) (*SearchAlertingConditionsResponse, error)
}

type client struct {
	cc grpc.ClientConnInterface
}

func NewAlertingConditionServiceClient(cc grpc.ClientConnInterface) AlertingConditionServiceClient {
	return &client{cc}
}

func (c *client) GetAlertingCondition(ctx context.Context, in *GetAlertingConditionRequest, opts ...grpc.CallOption) (*alerting_condition.AlertingCondition, error) {
	out := new(alerting_condition.AlertingCondition)
	err := c.cc.Invoke(ctx, "/ntt.monitoring.v4.AlertingConditionService/GetAlertingCondition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) BatchGetAlertingConditions(ctx context.Context, in *BatchGetAlertingConditionsRequest, opts ...grpc.CallOption) (*BatchGetAlertingConditionsResponse, error) {
	out := new(BatchGetAlertingConditionsResponse)
	err := c.cc.Invoke(ctx, "/ntt.monitoring.v4.AlertingConditionService/BatchGetAlertingConditions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) ListAlertingConditions(ctx context.Context, in *ListAlertingConditionsRequest, opts ...grpc.CallOption) (*ListAlertingConditionsResponse, error) {
	out := new(ListAlertingConditionsResponse)
	err := c.cc.Invoke(ctx, "/ntt.monitoring.v4.AlertingConditionService/ListAlertingConditions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) WatchAlertingCondition(ctx context.Context, in *WatchAlertingConditionRequest, opts ...grpc.CallOption) (WatchAlertingConditionClientStream, error) {
	stream, err := c.cc.NewStream(ctx,
		&grpc.StreamDesc{
			StreamName:    "WatchAlertingCondition",
			ServerStreams: true,
		},
		"/ntt.monitoring.v4.AlertingConditionService/WatchAlertingCondition", opts...)
	if err != nil {
		return nil, err
	}
	x := &watchAlertingConditionWatchAlertingConditionClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WatchAlertingConditionClientStream interface {
	Recv() (*WatchAlertingConditionResponse, error)
	grpc.ClientStream
}

type watchAlertingConditionWatchAlertingConditionClient struct {
	grpc.ClientStream
}

func (x *watchAlertingConditionWatchAlertingConditionClient) Recv() (*WatchAlertingConditionResponse, error) {
	m := new(WatchAlertingConditionResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *client) WatchAlertingConditions(ctx context.Context, in *WatchAlertingConditionsRequest, opts ...grpc.CallOption) (WatchAlertingConditionsClientStream, error) {
	stream, err := c.cc.NewStream(ctx,
		&grpc.StreamDesc{
			StreamName:    "WatchAlertingConditions",
			ServerStreams: true,
		},
		"/ntt.monitoring.v4.AlertingConditionService/WatchAlertingConditions", opts...)
	if err != nil {
		return nil, err
	}
	x := &watchAlertingConditionsWatchAlertingConditionsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WatchAlertingConditionsClientStream interface {
	Recv() (*WatchAlertingConditionsResponse, error)
	grpc.ClientStream
}

type watchAlertingConditionsWatchAlertingConditionsClient struct {
	grpc.ClientStream
}

func (x *watchAlertingConditionsWatchAlertingConditionsClient) Recv() (*WatchAlertingConditionsResponse, error) {
	m := new(WatchAlertingConditionsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *client) CreateAlertingCondition(ctx context.Context, in *CreateAlertingConditionRequest, opts ...grpc.CallOption) (*alerting_condition.AlertingCondition, error) {
	out := new(alerting_condition.AlertingCondition)
	err := c.cc.Invoke(ctx, "/ntt.monitoring.v4.AlertingConditionService/CreateAlertingCondition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) UpdateAlertingCondition(ctx context.Context, in *UpdateAlertingConditionRequest, opts ...grpc.CallOption) (*alerting_condition.AlertingCondition, error) {
	out := new(alerting_condition.AlertingCondition)
	err := c.cc.Invoke(ctx, "/ntt.monitoring.v4.AlertingConditionService/UpdateAlertingCondition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) DeleteAlertingCondition(ctx context.Context, in *DeleteAlertingConditionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ntt.monitoring.v4.AlertingConditionService/DeleteAlertingCondition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) SearchAlertingConditions(ctx context.Context, in *SearchAlertingConditionsRequest, opts ...grpc.CallOption) (*SearchAlertingConditionsResponse, error) {
	out := new(SearchAlertingConditionsResponse)
	err := c.cc.Invoke(ctx, "/ntt.monitoring.v4.AlertingConditionService/SearchAlertingConditions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}