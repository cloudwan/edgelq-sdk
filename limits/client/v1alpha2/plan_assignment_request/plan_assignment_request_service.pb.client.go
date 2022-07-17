// Code generated by protoc-gen-goten-client
// API: PlanAssignmentRequestService
// DO NOT EDIT!!!

package plan_assignment_request_client

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/runtime/protoimpl"
)

// proto imports
import (
	plan_assignment_request "github.com/cloudwan/edgelq-sdk/limits/resources/v1alpha2/plan_assignment_request"
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
	_ = &plan_assignment_request.PlanAssignmentRequest{}
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

// PlanAssignmentRequestServiceClient is the client API for PlanAssignmentRequestService.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PlanAssignmentRequestServiceClient interface {
	GetPlanAssignmentRequest(ctx context.Context, in *GetPlanAssignmentRequestRequest, opts ...grpc.CallOption) (*plan_assignment_request.PlanAssignmentRequest, error)
	BatchGetPlanAssignmentRequests(ctx context.Context, in *BatchGetPlanAssignmentRequestsRequest, opts ...grpc.CallOption) (*BatchGetPlanAssignmentRequestsResponse, error)
	ListPlanAssignmentRequests(ctx context.Context, in *ListPlanAssignmentRequestsRequest, opts ...grpc.CallOption) (*ListPlanAssignmentRequestsResponse, error)
	WatchPlanAssignmentRequest(ctx context.Context, in *WatchPlanAssignmentRequestRequest, opts ...grpc.CallOption) (WatchPlanAssignmentRequestClientStream, error)
	WatchPlanAssignmentRequests(ctx context.Context, in *WatchPlanAssignmentRequestsRequest, opts ...grpc.CallOption) (WatchPlanAssignmentRequestsClientStream, error)
	CreatePlanAssignmentRequest(ctx context.Context, in *CreatePlanAssignmentRequestRequest, opts ...grpc.CallOption) (*plan_assignment_request.PlanAssignmentRequest, error)
	UpdatePlanAssignmentRequest(ctx context.Context, in *UpdatePlanAssignmentRequestRequest, opts ...grpc.CallOption) (*plan_assignment_request.PlanAssignmentRequest, error)
	DeletePlanAssignmentRequest(ctx context.Context, in *DeletePlanAssignmentRequestRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	AcceptPlanAssignment(ctx context.Context, in *AcceptPlanAssignmentRequest, opts ...grpc.CallOption) (*AcceptPlanAssignmentResponse, error)
	DeclinePlanAssignment(ctx context.Context, in *DeclinePlanAssignmentRequest, opts ...grpc.CallOption) (*DeclinePlanAssignmentResponse, error)
}

type client struct {
	cc grpc.ClientConnInterface
}

func NewPlanAssignmentRequestServiceClient(cc grpc.ClientConnInterface) PlanAssignmentRequestServiceClient {
	return &client{cc}
}

func (c *client) GetPlanAssignmentRequest(ctx context.Context, in *GetPlanAssignmentRequestRequest, opts ...grpc.CallOption) (*plan_assignment_request.PlanAssignmentRequest, error) {
	out := new(plan_assignment_request.PlanAssignmentRequest)
	err := c.cc.Invoke(ctx, "/ntt.limits.v1alpha2.PlanAssignmentRequestService/GetPlanAssignmentRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) BatchGetPlanAssignmentRequests(ctx context.Context, in *BatchGetPlanAssignmentRequestsRequest, opts ...grpc.CallOption) (*BatchGetPlanAssignmentRequestsResponse, error) {
	out := new(BatchGetPlanAssignmentRequestsResponse)
	err := c.cc.Invoke(ctx, "/ntt.limits.v1alpha2.PlanAssignmentRequestService/BatchGetPlanAssignmentRequests", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) ListPlanAssignmentRequests(ctx context.Context, in *ListPlanAssignmentRequestsRequest, opts ...grpc.CallOption) (*ListPlanAssignmentRequestsResponse, error) {
	out := new(ListPlanAssignmentRequestsResponse)
	err := c.cc.Invoke(ctx, "/ntt.limits.v1alpha2.PlanAssignmentRequestService/ListPlanAssignmentRequests", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) WatchPlanAssignmentRequest(ctx context.Context, in *WatchPlanAssignmentRequestRequest, opts ...grpc.CallOption) (WatchPlanAssignmentRequestClientStream, error) {
	stream, err := c.cc.NewStream(ctx,
		&grpc.StreamDesc{
			StreamName:    "WatchPlanAssignmentRequest",
			ServerStreams: true,
		},
		"/ntt.limits.v1alpha2.PlanAssignmentRequestService/WatchPlanAssignmentRequest", opts...)
	if err != nil {
		return nil, err
	}
	x := &watchPlanAssignmentRequestWatchPlanAssignmentRequestClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WatchPlanAssignmentRequestClientStream interface {
	Recv() (*WatchPlanAssignmentRequestResponse, error)
	grpc.ClientStream
}

type watchPlanAssignmentRequestWatchPlanAssignmentRequestClient struct {
	grpc.ClientStream
}

func (x *watchPlanAssignmentRequestWatchPlanAssignmentRequestClient) Recv() (*WatchPlanAssignmentRequestResponse, error) {
	m := new(WatchPlanAssignmentRequestResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *client) WatchPlanAssignmentRequests(ctx context.Context, in *WatchPlanAssignmentRequestsRequest, opts ...grpc.CallOption) (WatchPlanAssignmentRequestsClientStream, error) {
	stream, err := c.cc.NewStream(ctx,
		&grpc.StreamDesc{
			StreamName:    "WatchPlanAssignmentRequests",
			ServerStreams: true,
		},
		"/ntt.limits.v1alpha2.PlanAssignmentRequestService/WatchPlanAssignmentRequests", opts...)
	if err != nil {
		return nil, err
	}
	x := &watchPlanAssignmentRequestsWatchPlanAssignmentRequestsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WatchPlanAssignmentRequestsClientStream interface {
	Recv() (*WatchPlanAssignmentRequestsResponse, error)
	grpc.ClientStream
}

type watchPlanAssignmentRequestsWatchPlanAssignmentRequestsClient struct {
	grpc.ClientStream
}

func (x *watchPlanAssignmentRequestsWatchPlanAssignmentRequestsClient) Recv() (*WatchPlanAssignmentRequestsResponse, error) {
	m := new(WatchPlanAssignmentRequestsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *client) CreatePlanAssignmentRequest(ctx context.Context, in *CreatePlanAssignmentRequestRequest, opts ...grpc.CallOption) (*plan_assignment_request.PlanAssignmentRequest, error) {
	out := new(plan_assignment_request.PlanAssignmentRequest)
	err := c.cc.Invoke(ctx, "/ntt.limits.v1alpha2.PlanAssignmentRequestService/CreatePlanAssignmentRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) UpdatePlanAssignmentRequest(ctx context.Context, in *UpdatePlanAssignmentRequestRequest, opts ...grpc.CallOption) (*plan_assignment_request.PlanAssignmentRequest, error) {
	out := new(plan_assignment_request.PlanAssignmentRequest)
	err := c.cc.Invoke(ctx, "/ntt.limits.v1alpha2.PlanAssignmentRequestService/UpdatePlanAssignmentRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) DeletePlanAssignmentRequest(ctx context.Context, in *DeletePlanAssignmentRequestRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/ntt.limits.v1alpha2.PlanAssignmentRequestService/DeletePlanAssignmentRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) AcceptPlanAssignment(ctx context.Context, in *AcceptPlanAssignmentRequest, opts ...grpc.CallOption) (*AcceptPlanAssignmentResponse, error) {
	out := new(AcceptPlanAssignmentResponse)
	err := c.cc.Invoke(ctx, "/ntt.limits.v1alpha2.PlanAssignmentRequestService/AcceptPlanAssignment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) DeclinePlanAssignment(ctx context.Context, in *DeclinePlanAssignmentRequest, opts ...grpc.CallOption) (*DeclinePlanAssignmentResponse, error) {
	out := new(DeclinePlanAssignmentResponse)
	err := c.cc.Invoke(ctx, "/ntt.limits.v1alpha2.PlanAssignmentRequestService/DeclinePlanAssignment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
