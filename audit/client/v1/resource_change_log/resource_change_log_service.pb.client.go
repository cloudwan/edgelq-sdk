// Code generated by protoc-gen-goten-client
// API: ResourceChangeLogService
// DO NOT EDIT!!!

package resource_change_log_client

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/runtime/protoimpl"
)

// proto imports
import (
	resource_change_log "github.com/cloudwan/edgelq-sdk/audit/resources/v1/resource_change_log"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = protoimpl.DescBuilder{}
	_ = new(context.Context)
	_ = grpc.ClientConn{}
)

// make sure we're using proto imports
var (
	_ = &resource_change_log.ResourceChangeLog{}
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

// ResourceChangeLogServiceClient is the client API for ResourceChangeLogService.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ResourceChangeLogServiceClient interface {
	ListResourceChangeLogs(ctx context.Context, in *ListResourceChangeLogsRequest, opts ...grpc.CallOption) (*ListResourceChangeLogsResponse, error)
	CreatePreCommittedResourceChangeLogs(ctx context.Context, in *CreatePreCommittedResourceChangeLogsRequest, opts ...grpc.CallOption) (*CreatePreCommittedResourceChangeLogsResponse, error)
	SetResourceChangeLogsCommitState(ctx context.Context, in *SetResourceChangeLogsCommitStateRequest, opts ...grpc.CallOption) (*SetResourceChangeLogsCommitStateResponse, error)
}

type client struct {
	cc grpc.ClientConnInterface
}

func NewResourceChangeLogServiceClient(cc grpc.ClientConnInterface) ResourceChangeLogServiceClient {
	return &client{cc}
}

func (c *client) ListResourceChangeLogs(ctx context.Context, in *ListResourceChangeLogsRequest, opts ...grpc.CallOption) (*ListResourceChangeLogsResponse, error) {
	out := new(ListResourceChangeLogsResponse)
	err := c.cc.Invoke(ctx, "/ntt.audit.v1.ResourceChangeLogService/ListResourceChangeLogs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) CreatePreCommittedResourceChangeLogs(ctx context.Context, in *CreatePreCommittedResourceChangeLogsRequest, opts ...grpc.CallOption) (*CreatePreCommittedResourceChangeLogsResponse, error) {
	out := new(CreatePreCommittedResourceChangeLogsResponse)
	err := c.cc.Invoke(ctx, "/ntt.audit.v1.ResourceChangeLogService/CreatePreCommittedResourceChangeLogs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) SetResourceChangeLogsCommitState(ctx context.Context, in *SetResourceChangeLogsCommitStateRequest, opts ...grpc.CallOption) (*SetResourceChangeLogsCommitStateResponse, error) {
	out := new(SetResourceChangeLogsCommitStateResponse)
	err := c.cc.Invoke(ctx, "/ntt.audit.v1.ResourceChangeLogService/SetResourceChangeLogsCommitState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
