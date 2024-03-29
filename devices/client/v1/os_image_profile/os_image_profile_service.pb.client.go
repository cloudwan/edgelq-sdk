// Code generated by protoc-gen-goten-client
// API: OsImageProfileService
// DO NOT EDIT!!!

package os_image_profile_client

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/runtime/protoimpl"
)

// proto imports
import (
	os_image_profile "github.com/cloudwan/edgelq-sdk/devices/resources/v1/os_image_profile"
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
	_ = &os_image_profile.OsImageProfile{}
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

// OsImageProfileServiceClient is the client API for OsImageProfileService.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OsImageProfileServiceClient interface {
	GetOsImageProfile(ctx context.Context, in *GetOsImageProfileRequest, opts ...grpc.CallOption) (*os_image_profile.OsImageProfile, error)
	BatchGetOsImageProfiles(ctx context.Context, in *BatchGetOsImageProfilesRequest, opts ...grpc.CallOption) (*BatchGetOsImageProfilesResponse, error)
	ListOsImageProfiles(ctx context.Context, in *ListOsImageProfilesRequest, opts ...grpc.CallOption) (*ListOsImageProfilesResponse, error)
	WatchOsImageProfile(ctx context.Context, in *WatchOsImageProfileRequest, opts ...grpc.CallOption) (WatchOsImageProfileClientStream, error)
	WatchOsImageProfiles(ctx context.Context, in *WatchOsImageProfilesRequest, opts ...grpc.CallOption) (WatchOsImageProfilesClientStream, error)
	CreateOsImageProfile(ctx context.Context, in *CreateOsImageProfileRequest, opts ...grpc.CallOption) (*os_image_profile.OsImageProfile, error)
	UpdateOsImageProfile(ctx context.Context, in *UpdateOsImageProfileRequest, opts ...grpc.CallOption) (*os_image_profile.OsImageProfile, error)
	DeleteOsImageProfile(ctx context.Context, in *DeleteOsImageProfileRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type client struct {
	cc grpc.ClientConnInterface
}

func NewOsImageProfileServiceClient(cc grpc.ClientConnInterface) OsImageProfileServiceClient {
	return &client{cc}
}

func (c *client) GetOsImageProfile(ctx context.Context, in *GetOsImageProfileRequest, opts ...grpc.CallOption) (*os_image_profile.OsImageProfile, error) {
	out := new(os_image_profile.OsImageProfile)
	err := c.cc.Invoke(ctx, "/ntt.devices.v1.OsImageProfileService/GetOsImageProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) BatchGetOsImageProfiles(ctx context.Context, in *BatchGetOsImageProfilesRequest, opts ...grpc.CallOption) (*BatchGetOsImageProfilesResponse, error) {
	out := new(BatchGetOsImageProfilesResponse)
	err := c.cc.Invoke(ctx, "/ntt.devices.v1.OsImageProfileService/BatchGetOsImageProfiles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) ListOsImageProfiles(ctx context.Context, in *ListOsImageProfilesRequest, opts ...grpc.CallOption) (*ListOsImageProfilesResponse, error) {
	out := new(ListOsImageProfilesResponse)
	err := c.cc.Invoke(ctx, "/ntt.devices.v1.OsImageProfileService/ListOsImageProfiles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) WatchOsImageProfile(ctx context.Context, in *WatchOsImageProfileRequest, opts ...grpc.CallOption) (WatchOsImageProfileClientStream, error) {
	stream, err := c.cc.NewStream(ctx,
		&grpc.StreamDesc{
			StreamName:    "WatchOsImageProfile",
			ServerStreams: true,
		},
		"/ntt.devices.v1.OsImageProfileService/WatchOsImageProfile", opts...)
	if err != nil {
		return nil, err
	}
	x := &watchOsImageProfileWatchOsImageProfileClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WatchOsImageProfileClientStream interface {
	Recv() (*WatchOsImageProfileResponse, error)
	grpc.ClientStream
}

type watchOsImageProfileWatchOsImageProfileClient struct {
	grpc.ClientStream
}

func (x *watchOsImageProfileWatchOsImageProfileClient) Recv() (*WatchOsImageProfileResponse, error) {
	m := new(WatchOsImageProfileResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *client) WatchOsImageProfiles(ctx context.Context, in *WatchOsImageProfilesRequest, opts ...grpc.CallOption) (WatchOsImageProfilesClientStream, error) {
	stream, err := c.cc.NewStream(ctx,
		&grpc.StreamDesc{
			StreamName:    "WatchOsImageProfiles",
			ServerStreams: true,
		},
		"/ntt.devices.v1.OsImageProfileService/WatchOsImageProfiles", opts...)
	if err != nil {
		return nil, err
	}
	x := &watchOsImageProfilesWatchOsImageProfilesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WatchOsImageProfilesClientStream interface {
	Recv() (*WatchOsImageProfilesResponse, error)
	grpc.ClientStream
}

type watchOsImageProfilesWatchOsImageProfilesClient struct {
	grpc.ClientStream
}

func (x *watchOsImageProfilesWatchOsImageProfilesClient) Recv() (*WatchOsImageProfilesResponse, error) {
	m := new(WatchOsImageProfilesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *client) CreateOsImageProfile(ctx context.Context, in *CreateOsImageProfileRequest, opts ...grpc.CallOption) (*os_image_profile.OsImageProfile, error) {
	out := new(os_image_profile.OsImageProfile)
	err := c.cc.Invoke(ctx, "/ntt.devices.v1.OsImageProfileService/CreateOsImageProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) UpdateOsImageProfile(ctx context.Context, in *UpdateOsImageProfileRequest, opts ...grpc.CallOption) (*os_image_profile.OsImageProfile, error) {
	out := new(os_image_profile.OsImageProfile)
	err := c.cc.Invoke(ctx, "/ntt.devices.v1.OsImageProfileService/UpdateOsImageProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) DeleteOsImageProfile(ctx context.Context, in *DeleteOsImageProfileRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ntt.devices.v1.OsImageProfileService/DeleteOsImageProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
