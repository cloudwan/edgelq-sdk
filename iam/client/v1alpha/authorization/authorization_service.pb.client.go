// Code generated by protoc-gen-goten-client
// API: AuthorizationService
// DO NOT EDIT!!!

package authorization_client

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

// AuthorizationServiceClient is the client API for AuthorizationService.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthorizationServiceClient interface {
	CheckPermissions(ctx context.Context, in *CheckPermissionsRequest, opts ...grpc.CallOption) (*CheckPermissionsResponse, error)
	CheckMyPermissions(ctx context.Context, in *CheckMyPermissionsRequest, opts ...grpc.CallOption) (*CheckMyPermissionsResponse, error)
}

type client struct {
	cc grpc.ClientConnInterface
}

func NewAuthorizationServiceClient(cc grpc.ClientConnInterface) AuthorizationServiceClient {
	return &client{cc}
}

func (c *client) CheckPermissions(ctx context.Context, in *CheckPermissionsRequest, opts ...grpc.CallOption) (*CheckPermissionsResponse, error) {
	out := new(CheckPermissionsResponse)
	err := c.cc.Invoke(ctx, "/ntt.iam.v1alpha.AuthorizationService/CheckPermissions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) CheckMyPermissions(ctx context.Context, in *CheckMyPermissionsRequest, opts ...grpc.CallOption) (*CheckMyPermissionsResponse, error) {
	out := new(CheckMyPermissionsResponse)
	err := c.cc.Invoke(ctx, "/ntt.iam.v1alpha.AuthorizationService/CheckMyPermissions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}