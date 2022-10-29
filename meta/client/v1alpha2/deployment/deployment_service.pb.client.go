// Code generated by protoc-gen-goten-client
// API: DeploymentService
// DO NOT EDIT!!!

package deployment_client

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/runtime/protoimpl"
)

// proto imports
import (
	deployment "github.com/cloudwan/edgelq-sdk/meta/resources/v1alpha2/deployment"
	empty "github.com/golang/protobuf/ptypes/empty"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = protoimpl.DescBuilder{}
	_ = new(context.Context)
	_ = grpc.ClientConn{}
)

// make sure we're using proto imports
var (
	_ = &deployment.Deployment{}
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

// DeploymentServiceClient is the client API for DeploymentService.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DeploymentServiceClient interface {
	GetDeployment(ctx context.Context, in *GetDeploymentRequest, opts ...grpc.CallOption) (*deployment.Deployment, error)
	BatchGetDeployments(ctx context.Context, in *BatchGetDeploymentsRequest, opts ...grpc.CallOption) (*BatchGetDeploymentsResponse, error)
	ListDeployments(ctx context.Context, in *ListDeploymentsRequest, opts ...grpc.CallOption) (*ListDeploymentsResponse, error)
	WatchDeployment(ctx context.Context, in *WatchDeploymentRequest, opts ...grpc.CallOption) (WatchDeploymentClientStream, error)
	WatchDeployments(ctx context.Context, in *WatchDeploymentsRequest, opts ...grpc.CallOption) (WatchDeploymentsClientStream, error)
	CreateDeployment(ctx context.Context, in *CreateDeploymentRequest, opts ...grpc.CallOption) (*deployment.Deployment, error)
	UpdateDeployment(ctx context.Context, in *UpdateDeploymentRequest, opts ...grpc.CallOption) (*deployment.Deployment, error)
	DeleteDeployment(ctx context.Context, in *DeleteDeploymentRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type client struct {
	cc grpc.ClientConnInterface
}

func NewDeploymentServiceClient(cc grpc.ClientConnInterface) DeploymentServiceClient {
	return &client{cc}
}

func (c *client) GetDeployment(ctx context.Context, in *GetDeploymentRequest, opts ...grpc.CallOption) (*deployment.Deployment, error) {
	out := new(deployment.Deployment)
	err := c.cc.Invoke(ctx, "/ntt.meta.v1alpha2.DeploymentService/GetDeployment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) BatchGetDeployments(ctx context.Context, in *BatchGetDeploymentsRequest, opts ...grpc.CallOption) (*BatchGetDeploymentsResponse, error) {
	out := new(BatchGetDeploymentsResponse)
	err := c.cc.Invoke(ctx, "/ntt.meta.v1alpha2.DeploymentService/BatchGetDeployments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) ListDeployments(ctx context.Context, in *ListDeploymentsRequest, opts ...grpc.CallOption) (*ListDeploymentsResponse, error) {
	out := new(ListDeploymentsResponse)
	err := c.cc.Invoke(ctx, "/ntt.meta.v1alpha2.DeploymentService/ListDeployments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) WatchDeployment(ctx context.Context, in *WatchDeploymentRequest, opts ...grpc.CallOption) (WatchDeploymentClientStream, error) {
	stream, err := c.cc.NewStream(ctx,
		&grpc.StreamDesc{
			StreamName:    "WatchDeployment",
			ServerStreams: true,
		},
		"/ntt.meta.v1alpha2.DeploymentService/WatchDeployment", opts...)
	if err != nil {
		return nil, err
	}
	x := &watchDeploymentWatchDeploymentClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WatchDeploymentClientStream interface {
	Recv() (*WatchDeploymentResponse, error)
	grpc.ClientStream
}

type watchDeploymentWatchDeploymentClient struct {
	grpc.ClientStream
}

func (x *watchDeploymentWatchDeploymentClient) Recv() (*WatchDeploymentResponse, error) {
	m := new(WatchDeploymentResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *client) WatchDeployments(ctx context.Context, in *WatchDeploymentsRequest, opts ...grpc.CallOption) (WatchDeploymentsClientStream, error) {
	stream, err := c.cc.NewStream(ctx,
		&grpc.StreamDesc{
			StreamName:    "WatchDeployments",
			ServerStreams: true,
		},
		"/ntt.meta.v1alpha2.DeploymentService/WatchDeployments", opts...)
	if err != nil {
		return nil, err
	}
	x := &watchDeploymentsWatchDeploymentsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WatchDeploymentsClientStream interface {
	Recv() (*WatchDeploymentsResponse, error)
	grpc.ClientStream
}

type watchDeploymentsWatchDeploymentsClient struct {
	grpc.ClientStream
}

func (x *watchDeploymentsWatchDeploymentsClient) Recv() (*WatchDeploymentsResponse, error) {
	m := new(WatchDeploymentsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *client) CreateDeployment(ctx context.Context, in *CreateDeploymentRequest, opts ...grpc.CallOption) (*deployment.Deployment, error) {
	out := new(deployment.Deployment)
	err := c.cc.Invoke(ctx, "/ntt.meta.v1alpha2.DeploymentService/CreateDeployment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) UpdateDeployment(ctx context.Context, in *UpdateDeploymentRequest, opts ...grpc.CallOption) (*deployment.Deployment, error) {
	out := new(deployment.Deployment)
	err := c.cc.Invoke(ctx, "/ntt.meta.v1alpha2.DeploymentService/UpdateDeployment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) DeleteDeployment(ctx context.Context, in *DeleteDeploymentRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/ntt.meta.v1alpha2.DeploymentService/DeleteDeployment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
