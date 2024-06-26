// Code generated by protoc-gen-goten-client
// API: ProvisioningApprovalRequestService
// DO NOT EDIT!!!

package provisioning_approval_request_client

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/runtime/protoimpl"
)

// proto imports
import (
	provisioning_approval_request "github.com/cloudwan/edgelq-sdk/devices/resources/v1/provisioning_approval_request"
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
	_ = &provisioning_approval_request.ProvisioningApprovalRequest{}
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

// ProvisioningApprovalRequestServiceClient is the client API for ProvisioningApprovalRequestService.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProvisioningApprovalRequestServiceClient interface {
	GetProvisioningApprovalRequest(ctx context.Context, in *GetProvisioningApprovalRequestRequest, opts ...grpc.CallOption) (*provisioning_approval_request.ProvisioningApprovalRequest, error)
	BatchGetProvisioningApprovalRequests(ctx context.Context, in *BatchGetProvisioningApprovalRequestsRequest, opts ...grpc.CallOption) (*BatchGetProvisioningApprovalRequestsResponse, error)
	ListProvisioningApprovalRequests(ctx context.Context, in *ListProvisioningApprovalRequestsRequest, opts ...grpc.CallOption) (*ListProvisioningApprovalRequestsResponse, error)
	WatchProvisioningApprovalRequest(ctx context.Context, in *WatchProvisioningApprovalRequestRequest, opts ...grpc.CallOption) (WatchProvisioningApprovalRequestClientStream, error)
	WatchProvisioningApprovalRequests(ctx context.Context, in *WatchProvisioningApprovalRequestsRequest, opts ...grpc.CallOption) (WatchProvisioningApprovalRequestsClientStream, error)
	CreateProvisioningApprovalRequest(ctx context.Context, in *CreateProvisioningApprovalRequestRequest, opts ...grpc.CallOption) (*provisioning_approval_request.ProvisioningApprovalRequest, error)
	UpdateProvisioningApprovalRequest(ctx context.Context, in *UpdateProvisioningApprovalRequestRequest, opts ...grpc.CallOption) (*provisioning_approval_request.ProvisioningApprovalRequest, error)
	DeleteProvisioningApprovalRequest(ctx context.Context, in *DeleteProvisioningApprovalRequestRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ProvisionDeviceForApprovedRequest(ctx context.Context, in *ProvisionDeviceForApprovedRequestRequest, opts ...grpc.CallOption) (*ProvisionDeviceForApprovedRequestResponse, error)
}

type client struct {
	cc grpc.ClientConnInterface
}

func NewProvisioningApprovalRequestServiceClient(cc grpc.ClientConnInterface) ProvisioningApprovalRequestServiceClient {
	return &client{cc}
}

func (c *client) GetProvisioningApprovalRequest(ctx context.Context, in *GetProvisioningApprovalRequestRequest, opts ...grpc.CallOption) (*provisioning_approval_request.ProvisioningApprovalRequest, error) {
	out := new(provisioning_approval_request.ProvisioningApprovalRequest)
	err := c.cc.Invoke(ctx, "/ntt.devices.v1.ProvisioningApprovalRequestService/GetProvisioningApprovalRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) BatchGetProvisioningApprovalRequests(ctx context.Context, in *BatchGetProvisioningApprovalRequestsRequest, opts ...grpc.CallOption) (*BatchGetProvisioningApprovalRequestsResponse, error) {
	out := new(BatchGetProvisioningApprovalRequestsResponse)
	err := c.cc.Invoke(ctx, "/ntt.devices.v1.ProvisioningApprovalRequestService/BatchGetProvisioningApprovalRequests", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) ListProvisioningApprovalRequests(ctx context.Context, in *ListProvisioningApprovalRequestsRequest, opts ...grpc.CallOption) (*ListProvisioningApprovalRequestsResponse, error) {
	out := new(ListProvisioningApprovalRequestsResponse)
	err := c.cc.Invoke(ctx, "/ntt.devices.v1.ProvisioningApprovalRequestService/ListProvisioningApprovalRequests", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) WatchProvisioningApprovalRequest(ctx context.Context, in *WatchProvisioningApprovalRequestRequest, opts ...grpc.CallOption) (WatchProvisioningApprovalRequestClientStream, error) {
	stream, err := c.cc.NewStream(ctx,
		&grpc.StreamDesc{
			StreamName:    "WatchProvisioningApprovalRequest",
			ServerStreams: true,
		},
		"/ntt.devices.v1.ProvisioningApprovalRequestService/WatchProvisioningApprovalRequest", opts...)
	if err != nil {
		return nil, err
	}
	x := &watchProvisioningApprovalRequestWatchProvisioningApprovalRequestClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WatchProvisioningApprovalRequestClientStream interface {
	Recv() (*WatchProvisioningApprovalRequestResponse, error)
	grpc.ClientStream
}

type watchProvisioningApprovalRequestWatchProvisioningApprovalRequestClient struct {
	grpc.ClientStream
}

func (x *watchProvisioningApprovalRequestWatchProvisioningApprovalRequestClient) Recv() (*WatchProvisioningApprovalRequestResponse, error) {
	m := new(WatchProvisioningApprovalRequestResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *client) WatchProvisioningApprovalRequests(ctx context.Context, in *WatchProvisioningApprovalRequestsRequest, opts ...grpc.CallOption) (WatchProvisioningApprovalRequestsClientStream, error) {
	stream, err := c.cc.NewStream(ctx,
		&grpc.StreamDesc{
			StreamName:    "WatchProvisioningApprovalRequests",
			ServerStreams: true,
		},
		"/ntt.devices.v1.ProvisioningApprovalRequestService/WatchProvisioningApprovalRequests", opts...)
	if err != nil {
		return nil, err
	}
	x := &watchProvisioningApprovalRequestsWatchProvisioningApprovalRequestsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WatchProvisioningApprovalRequestsClientStream interface {
	Recv() (*WatchProvisioningApprovalRequestsResponse, error)
	grpc.ClientStream
}

type watchProvisioningApprovalRequestsWatchProvisioningApprovalRequestsClient struct {
	grpc.ClientStream
}

func (x *watchProvisioningApprovalRequestsWatchProvisioningApprovalRequestsClient) Recv() (*WatchProvisioningApprovalRequestsResponse, error) {
	m := new(WatchProvisioningApprovalRequestsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *client) CreateProvisioningApprovalRequest(ctx context.Context, in *CreateProvisioningApprovalRequestRequest, opts ...grpc.CallOption) (*provisioning_approval_request.ProvisioningApprovalRequest, error) {
	out := new(provisioning_approval_request.ProvisioningApprovalRequest)
	err := c.cc.Invoke(ctx, "/ntt.devices.v1.ProvisioningApprovalRequestService/CreateProvisioningApprovalRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) UpdateProvisioningApprovalRequest(ctx context.Context, in *UpdateProvisioningApprovalRequestRequest, opts ...grpc.CallOption) (*provisioning_approval_request.ProvisioningApprovalRequest, error) {
	out := new(provisioning_approval_request.ProvisioningApprovalRequest)
	err := c.cc.Invoke(ctx, "/ntt.devices.v1.ProvisioningApprovalRequestService/UpdateProvisioningApprovalRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) DeleteProvisioningApprovalRequest(ctx context.Context, in *DeleteProvisioningApprovalRequestRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ntt.devices.v1.ProvisioningApprovalRequestService/DeleteProvisioningApprovalRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) ProvisionDeviceForApprovedRequest(ctx context.Context, in *ProvisionDeviceForApprovedRequestRequest, opts ...grpc.CallOption) (*ProvisionDeviceForApprovedRequestResponse, error) {
	out := new(ProvisionDeviceForApprovedRequestResponse)
	err := c.cc.Invoke(ctx, "/ntt.devices.v1.ProvisioningApprovalRequestService/ProvisionDeviceForApprovedRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
