// Code generated by protoc-gen-goten-client
// API: GroupMemberService
// DO NOT EDIT!!!

package group_member_client

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/runtime/protoimpl"
)

// proto imports
import (
	group_member "github.com/cloudwan/edgelq-sdk/iam/resources/v1/group_member"
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
	_ = &group_member.GroupMember{}
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

// GroupMemberServiceClient is the client API for GroupMemberService.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GroupMemberServiceClient interface {
	GetGroupMember(ctx context.Context, in *GetGroupMemberRequest, opts ...grpc.CallOption) (*group_member.GroupMember, error)
	BatchGetGroupMembers(ctx context.Context, in *BatchGetGroupMembersRequest, opts ...grpc.CallOption) (*BatchGetGroupMembersResponse, error)
	ListGroupMembers(ctx context.Context, in *ListGroupMembersRequest, opts ...grpc.CallOption) (*ListGroupMembersResponse, error)
	WatchGroupMember(ctx context.Context, in *WatchGroupMemberRequest, opts ...grpc.CallOption) (WatchGroupMemberClientStream, error)
	WatchGroupMembers(ctx context.Context, in *WatchGroupMembersRequest, opts ...grpc.CallOption) (WatchGroupMembersClientStream, error)
	CreateGroupMember(ctx context.Context, in *CreateGroupMemberRequest, opts ...grpc.CallOption) (*group_member.GroupMember, error)
	UpdateGroupMember(ctx context.Context, in *UpdateGroupMemberRequest, opts ...grpc.CallOption) (*group_member.GroupMember, error)
	DeleteGroupMember(ctx context.Context, in *DeleteGroupMemberRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ListGroupMembersWithMembers(ctx context.Context, in *ListGroupMembersWithMembersRequest, opts ...grpc.CallOption) (*ListGroupMembersWithMembersResponse, error)
}

type client struct {
	cc grpc.ClientConnInterface
}

func NewGroupMemberServiceClient(cc grpc.ClientConnInterface) GroupMemberServiceClient {
	return &client{cc}
}

func (c *client) GetGroupMember(ctx context.Context, in *GetGroupMemberRequest, opts ...grpc.CallOption) (*group_member.GroupMember, error) {
	out := new(group_member.GroupMember)
	err := c.cc.Invoke(ctx, "/ntt.iam.v1.GroupMemberService/GetGroupMember", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) BatchGetGroupMembers(ctx context.Context, in *BatchGetGroupMembersRequest, opts ...grpc.CallOption) (*BatchGetGroupMembersResponse, error) {
	out := new(BatchGetGroupMembersResponse)
	err := c.cc.Invoke(ctx, "/ntt.iam.v1.GroupMemberService/BatchGetGroupMembers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) ListGroupMembers(ctx context.Context, in *ListGroupMembersRequest, opts ...grpc.CallOption) (*ListGroupMembersResponse, error) {
	out := new(ListGroupMembersResponse)
	err := c.cc.Invoke(ctx, "/ntt.iam.v1.GroupMemberService/ListGroupMembers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) WatchGroupMember(ctx context.Context, in *WatchGroupMemberRequest, opts ...grpc.CallOption) (WatchGroupMemberClientStream, error) {
	stream, err := c.cc.NewStream(ctx,
		&grpc.StreamDesc{
			StreamName:    "WatchGroupMember",
			ServerStreams: true,
		},
		"/ntt.iam.v1.GroupMemberService/WatchGroupMember", opts...)
	if err != nil {
		return nil, err
	}
	x := &watchGroupMemberWatchGroupMemberClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WatchGroupMemberClientStream interface {
	Recv() (*WatchGroupMemberResponse, error)
	grpc.ClientStream
}

type watchGroupMemberWatchGroupMemberClient struct {
	grpc.ClientStream
}

func (x *watchGroupMemberWatchGroupMemberClient) Recv() (*WatchGroupMemberResponse, error) {
	m := new(WatchGroupMemberResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *client) WatchGroupMembers(ctx context.Context, in *WatchGroupMembersRequest, opts ...grpc.CallOption) (WatchGroupMembersClientStream, error) {
	stream, err := c.cc.NewStream(ctx,
		&grpc.StreamDesc{
			StreamName:    "WatchGroupMembers",
			ServerStreams: true,
		},
		"/ntt.iam.v1.GroupMemberService/WatchGroupMembers", opts...)
	if err != nil {
		return nil, err
	}
	x := &watchGroupMembersWatchGroupMembersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WatchGroupMembersClientStream interface {
	Recv() (*WatchGroupMembersResponse, error)
	grpc.ClientStream
}

type watchGroupMembersWatchGroupMembersClient struct {
	grpc.ClientStream
}

func (x *watchGroupMembersWatchGroupMembersClient) Recv() (*WatchGroupMembersResponse, error) {
	m := new(WatchGroupMembersResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *client) CreateGroupMember(ctx context.Context, in *CreateGroupMemberRequest, opts ...grpc.CallOption) (*group_member.GroupMember, error) {
	out := new(group_member.GroupMember)
	err := c.cc.Invoke(ctx, "/ntt.iam.v1.GroupMemberService/CreateGroupMember", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) UpdateGroupMember(ctx context.Context, in *UpdateGroupMemberRequest, opts ...grpc.CallOption) (*group_member.GroupMember, error) {
	out := new(group_member.GroupMember)
	err := c.cc.Invoke(ctx, "/ntt.iam.v1.GroupMemberService/UpdateGroupMember", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) DeleteGroupMember(ctx context.Context, in *DeleteGroupMemberRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ntt.iam.v1.GroupMemberService/DeleteGroupMember", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) ListGroupMembersWithMembers(ctx context.Context, in *ListGroupMembersWithMembersRequest, opts ...grpc.CallOption) (*ListGroupMembersWithMembersResponse, error) {
	out := new(ListGroupMembersWithMembersResponse)
	err := c.cc.Invoke(ctx, "/ntt.iam.v1.GroupMemberService/ListGroupMembersWithMembers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
