// Code generated by protoc-gen-goten-client
// API: ActivityLogService
// DO NOT EDIT!!!

package activity_log_client

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/runtime/protoimpl"
)

// proto imports
import (
	activity_log "github.com/cloudwan/edgelq-sdk/audit/resources/v1alpha/activity_log"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = protoimpl.DescBuilder{}
	_ = new(context.Context)
	_ = grpc.ClientConn{}
)

// make sure we're using proto imports
var (
	_ = &activity_log.ActivityLog{}
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

// ActivityLogServiceClient is the client API for ActivityLogService.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ActivityLogServiceClient interface {
	ListActivityLogs(ctx context.Context, in *ListActivityLogsRequest, opts ...grpc.CallOption) (*ListActivityLogsResponse, error)
	CreateActivityLogs(ctx context.Context, in *CreateActivityLogsRequest, opts ...grpc.CallOption) (*CreateActivityLogsResponse, error)
}

type client struct {
	cc grpc.ClientConnInterface
}

func NewActivityLogServiceClient(cc grpc.ClientConnInterface) ActivityLogServiceClient {
	return &client{cc}
}

func (c *client) ListActivityLogs(ctx context.Context, in *ListActivityLogsRequest, opts ...grpc.CallOption) (*ListActivityLogsResponse, error) {
	out := new(ListActivityLogsResponse)
	err := c.cc.Invoke(ctx, "/ntt.audit.v1alpha.ActivityLogService/ListActivityLogs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) CreateActivityLogs(ctx context.Context, in *CreateActivityLogsRequest, opts ...grpc.CallOption) (*CreateActivityLogsResponse, error) {
	out := new(CreateActivityLogsResponse)
	err := c.cc.Invoke(ctx, "/ntt.audit.v1alpha.ActivityLogService/CreateActivityLogs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
