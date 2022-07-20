// Code generated by protoc-gen-goten-client
// API: AcceptedPlanService
// DO NOT EDIT!!!

package accepted_plan_client

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/runtime/protoimpl"
)

// proto imports
import (
	accepted_plan "github.com/cloudwan/edgelq-sdk/limits/resources/v1alpha2/accepted_plan"
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
	_ = &accepted_plan.AcceptedPlan{}
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

// AcceptedPlanServiceClient is the client API for AcceptedPlanService.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AcceptedPlanServiceClient interface {
	GetAcceptedPlan(ctx context.Context, in *GetAcceptedPlanRequest, opts ...grpc.CallOption) (*accepted_plan.AcceptedPlan, error)
	BatchGetAcceptedPlans(ctx context.Context, in *BatchGetAcceptedPlansRequest, opts ...grpc.CallOption) (*BatchGetAcceptedPlansResponse, error)
	ListAcceptedPlans(ctx context.Context, in *ListAcceptedPlansRequest, opts ...grpc.CallOption) (*ListAcceptedPlansResponse, error)
	WatchAcceptedPlan(ctx context.Context, in *WatchAcceptedPlanRequest, opts ...grpc.CallOption) (WatchAcceptedPlanClientStream, error)
	WatchAcceptedPlans(ctx context.Context, in *WatchAcceptedPlansRequest, opts ...grpc.CallOption) (WatchAcceptedPlansClientStream, error)
	CreateAcceptedPlan(ctx context.Context, in *CreateAcceptedPlanRequest, opts ...grpc.CallOption) (*accepted_plan.AcceptedPlan, error)
	UpdateAcceptedPlan(ctx context.Context, in *UpdateAcceptedPlanRequest, opts ...grpc.CallOption) (*accepted_plan.AcceptedPlan, error)
	DeleteAcceptedPlan(ctx context.Context, in *DeleteAcceptedPlanRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type client struct {
	cc grpc.ClientConnInterface
}

func NewAcceptedPlanServiceClient(cc grpc.ClientConnInterface) AcceptedPlanServiceClient {
	return &client{cc}
}

func (c *client) GetAcceptedPlan(ctx context.Context, in *GetAcceptedPlanRequest, opts ...grpc.CallOption) (*accepted_plan.AcceptedPlan, error) {
	out := new(accepted_plan.AcceptedPlan)
	err := c.cc.Invoke(ctx, "/ntt.limits.v1alpha2.AcceptedPlanService/GetAcceptedPlan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) BatchGetAcceptedPlans(ctx context.Context, in *BatchGetAcceptedPlansRequest, opts ...grpc.CallOption) (*BatchGetAcceptedPlansResponse, error) {
	out := new(BatchGetAcceptedPlansResponse)
	err := c.cc.Invoke(ctx, "/ntt.limits.v1alpha2.AcceptedPlanService/BatchGetAcceptedPlans", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) ListAcceptedPlans(ctx context.Context, in *ListAcceptedPlansRequest, opts ...grpc.CallOption) (*ListAcceptedPlansResponse, error) {
	out := new(ListAcceptedPlansResponse)
	err := c.cc.Invoke(ctx, "/ntt.limits.v1alpha2.AcceptedPlanService/ListAcceptedPlans", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) WatchAcceptedPlan(ctx context.Context, in *WatchAcceptedPlanRequest, opts ...grpc.CallOption) (WatchAcceptedPlanClientStream, error) {
	stream, err := c.cc.NewStream(ctx,
		&grpc.StreamDesc{
			StreamName:    "WatchAcceptedPlan",
			ServerStreams: true,
		},
		"/ntt.limits.v1alpha2.AcceptedPlanService/WatchAcceptedPlan", opts...)
	if err != nil {
		return nil, err
	}
	x := &watchAcceptedPlanWatchAcceptedPlanClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WatchAcceptedPlanClientStream interface {
	Recv() (*WatchAcceptedPlanResponse, error)
	grpc.ClientStream
}

type watchAcceptedPlanWatchAcceptedPlanClient struct {
	grpc.ClientStream
}

func (x *watchAcceptedPlanWatchAcceptedPlanClient) Recv() (*WatchAcceptedPlanResponse, error) {
	m := new(WatchAcceptedPlanResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *client) WatchAcceptedPlans(ctx context.Context, in *WatchAcceptedPlansRequest, opts ...grpc.CallOption) (WatchAcceptedPlansClientStream, error) {
	stream, err := c.cc.NewStream(ctx,
		&grpc.StreamDesc{
			StreamName:    "WatchAcceptedPlans",
			ServerStreams: true,
		},
		"/ntt.limits.v1alpha2.AcceptedPlanService/WatchAcceptedPlans", opts...)
	if err != nil {
		return nil, err
	}
	x := &watchAcceptedPlansWatchAcceptedPlansClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WatchAcceptedPlansClientStream interface {
	Recv() (*WatchAcceptedPlansResponse, error)
	grpc.ClientStream
}

type watchAcceptedPlansWatchAcceptedPlansClient struct {
	grpc.ClientStream
}

func (x *watchAcceptedPlansWatchAcceptedPlansClient) Recv() (*WatchAcceptedPlansResponse, error) {
	m := new(WatchAcceptedPlansResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *client) CreateAcceptedPlan(ctx context.Context, in *CreateAcceptedPlanRequest, opts ...grpc.CallOption) (*accepted_plan.AcceptedPlan, error) {
	out := new(accepted_plan.AcceptedPlan)
	err := c.cc.Invoke(ctx, "/ntt.limits.v1alpha2.AcceptedPlanService/CreateAcceptedPlan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) UpdateAcceptedPlan(ctx context.Context, in *UpdateAcceptedPlanRequest, opts ...grpc.CallOption) (*accepted_plan.AcceptedPlan, error) {
	out := new(accepted_plan.AcceptedPlan)
	err := c.cc.Invoke(ctx, "/ntt.limits.v1alpha2.AcceptedPlanService/UpdateAcceptedPlan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) DeleteAcceptedPlan(ctx context.Context, in *DeleteAcceptedPlanRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/ntt.limits.v1alpha2.AcceptedPlanService/DeleteAcceptedPlan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}