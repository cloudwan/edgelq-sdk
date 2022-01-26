// Code generated by protoc-gen-goten-client
// API: OrganizationInvitationService
// DO NOT EDIT!!!

package organization_invitation_client

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/runtime/protoimpl"
)

// proto imports
import (
	organization_invitation "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/organization_invitation"
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
	_ = &organization_invitation.OrganizationInvitation{}
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

// OrganizationInvitationServiceClient is the client API for OrganizationInvitationService.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OrganizationInvitationServiceClient interface {
	GetOrganizationInvitation(ctx context.Context, in *GetOrganizationInvitationRequest, opts ...grpc.CallOption) (*organization_invitation.OrganizationInvitation, error)
	BatchGetOrganizationInvitations(ctx context.Context, in *BatchGetOrganizationInvitationsRequest, opts ...grpc.CallOption) (*BatchGetOrganizationInvitationsResponse, error)
	ListOrganizationInvitations(ctx context.Context, in *ListOrganizationInvitationsRequest, opts ...grpc.CallOption) (*ListOrganizationInvitationsResponse, error)
	WatchOrganizationInvitation(ctx context.Context, in *WatchOrganizationInvitationRequest, opts ...grpc.CallOption) (WatchOrganizationInvitationClientStream, error)
	WatchOrganizationInvitations(ctx context.Context, in *WatchOrganizationInvitationsRequest, opts ...grpc.CallOption) (WatchOrganizationInvitationsClientStream, error)
	CreateOrganizationInvitation(ctx context.Context, in *CreateOrganizationInvitationRequest, opts ...grpc.CallOption) (*organization_invitation.OrganizationInvitation, error)
	UpdateOrganizationInvitation(ctx context.Context, in *UpdateOrganizationInvitationRequest, opts ...grpc.CallOption) (*organization_invitation.OrganizationInvitation, error)
	DeleteOrganizationInvitation(ctx context.Context, in *DeleteOrganizationInvitationRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	AcceptOrganizationInvitation(ctx context.Context, in *AcceptOrganizationInvitationRequest, opts ...grpc.CallOption) (*AcceptOrganizationInvitationResponse, error)
	DeclineOrganizationInvitation(ctx context.Context, in *DeclineOrganizationInvitationRequest, opts ...grpc.CallOption) (*DeclineOrganizationInvitationResponse, error)
	ListMyOrganizationInvitations(ctx context.Context, in *ListMyOrganizationInvitationsRequest, opts ...grpc.CallOption) (*ListMyOrganizationInvitationsResponse, error)
}

type client struct {
	cc grpc.ClientConnInterface
}

func NewOrganizationInvitationServiceClient(cc grpc.ClientConnInterface) OrganizationInvitationServiceClient {
	return &client{cc}
}

func (c *client) GetOrganizationInvitation(ctx context.Context, in *GetOrganizationInvitationRequest, opts ...grpc.CallOption) (*organization_invitation.OrganizationInvitation, error) {
	out := new(organization_invitation.OrganizationInvitation)
	err := c.cc.Invoke(ctx, "/ntt.iam.v1alpha2.OrganizationInvitationService/GetOrganizationInvitation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) BatchGetOrganizationInvitations(ctx context.Context, in *BatchGetOrganizationInvitationsRequest, opts ...grpc.CallOption) (*BatchGetOrganizationInvitationsResponse, error) {
	out := new(BatchGetOrganizationInvitationsResponse)
	err := c.cc.Invoke(ctx, "/ntt.iam.v1alpha2.OrganizationInvitationService/BatchGetOrganizationInvitations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) ListOrganizationInvitations(ctx context.Context, in *ListOrganizationInvitationsRequest, opts ...grpc.CallOption) (*ListOrganizationInvitationsResponse, error) {
	out := new(ListOrganizationInvitationsResponse)
	err := c.cc.Invoke(ctx, "/ntt.iam.v1alpha2.OrganizationInvitationService/ListOrganizationInvitations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) WatchOrganizationInvitation(ctx context.Context, in *WatchOrganizationInvitationRequest, opts ...grpc.CallOption) (WatchOrganizationInvitationClientStream, error) {
	stream, err := c.cc.NewStream(ctx,
		&grpc.StreamDesc{
			StreamName:    "WatchOrganizationInvitation",
			ServerStreams: true,
		},
		"/ntt.iam.v1alpha2.OrganizationInvitationService/WatchOrganizationInvitation", opts...)
	if err != nil {
		return nil, err
	}
	x := &watchOrganizationInvitationWatchOrganizationInvitationClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WatchOrganizationInvitationClientStream interface {
	Recv() (*WatchOrganizationInvitationResponse, error)
	grpc.ClientStream
}

type watchOrganizationInvitationWatchOrganizationInvitationClient struct {
	grpc.ClientStream
}

func (x *watchOrganizationInvitationWatchOrganizationInvitationClient) Recv() (*WatchOrganizationInvitationResponse, error) {
	m := new(WatchOrganizationInvitationResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *client) WatchOrganizationInvitations(ctx context.Context, in *WatchOrganizationInvitationsRequest, opts ...grpc.CallOption) (WatchOrganizationInvitationsClientStream, error) {
	stream, err := c.cc.NewStream(ctx,
		&grpc.StreamDesc{
			StreamName:    "WatchOrganizationInvitations",
			ServerStreams: true,
		},
		"/ntt.iam.v1alpha2.OrganizationInvitationService/WatchOrganizationInvitations", opts...)
	if err != nil {
		return nil, err
	}
	x := &watchOrganizationInvitationsWatchOrganizationInvitationsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WatchOrganizationInvitationsClientStream interface {
	Recv() (*WatchOrganizationInvitationsResponse, error)
	grpc.ClientStream
}

type watchOrganizationInvitationsWatchOrganizationInvitationsClient struct {
	grpc.ClientStream
}

func (x *watchOrganizationInvitationsWatchOrganizationInvitationsClient) Recv() (*WatchOrganizationInvitationsResponse, error) {
	m := new(WatchOrganizationInvitationsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *client) CreateOrganizationInvitation(ctx context.Context, in *CreateOrganizationInvitationRequest, opts ...grpc.CallOption) (*organization_invitation.OrganizationInvitation, error) {
	out := new(organization_invitation.OrganizationInvitation)
	err := c.cc.Invoke(ctx, "/ntt.iam.v1alpha2.OrganizationInvitationService/CreateOrganizationInvitation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) UpdateOrganizationInvitation(ctx context.Context, in *UpdateOrganizationInvitationRequest, opts ...grpc.CallOption) (*organization_invitation.OrganizationInvitation, error) {
	out := new(organization_invitation.OrganizationInvitation)
	err := c.cc.Invoke(ctx, "/ntt.iam.v1alpha2.OrganizationInvitationService/UpdateOrganizationInvitation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) DeleteOrganizationInvitation(ctx context.Context, in *DeleteOrganizationInvitationRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/ntt.iam.v1alpha2.OrganizationInvitationService/DeleteOrganizationInvitation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) AcceptOrganizationInvitation(ctx context.Context, in *AcceptOrganizationInvitationRequest, opts ...grpc.CallOption) (*AcceptOrganizationInvitationResponse, error) {
	out := new(AcceptOrganizationInvitationResponse)
	err := c.cc.Invoke(ctx, "/ntt.iam.v1alpha2.OrganizationInvitationService/AcceptOrganizationInvitation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) DeclineOrganizationInvitation(ctx context.Context, in *DeclineOrganizationInvitationRequest, opts ...grpc.CallOption) (*DeclineOrganizationInvitationResponse, error) {
	out := new(DeclineOrganizationInvitationResponse)
	err := c.cc.Invoke(ctx, "/ntt.iam.v1alpha2.OrganizationInvitationService/DeclineOrganizationInvitation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) ListMyOrganizationInvitations(ctx context.Context, in *ListMyOrganizationInvitationsRequest, opts ...grpc.CallOption) (*ListMyOrganizationInvitationsResponse, error) {
	out := new(ListMyOrganizationInvitationsResponse)
	err := c.cc.Invoke(ctx, "/ntt.iam.v1alpha2.OrganizationInvitationService/ListMyOrganizationInvitations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
