// Code generated by protoc-gen-goten-client
// API: ProjectService
// DO NOT EDIT!!!

package project_client

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/runtime/protoimpl"
)

// proto imports
import (
	project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/project"
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
	_ = &project.Project{}
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

// ProjectServiceClient is the client API for ProjectService.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProjectServiceClient interface {
	GetProject(ctx context.Context, in *GetProjectRequest, opts ...grpc.CallOption) (*project.Project, error)
	BatchGetProjects(ctx context.Context, in *BatchGetProjectsRequest, opts ...grpc.CallOption) (*BatchGetProjectsResponse, error)
	ListProjects(ctx context.Context, in *ListProjectsRequest, opts ...grpc.CallOption) (*ListProjectsResponse, error)
	WatchProject(ctx context.Context, in *WatchProjectRequest, opts ...grpc.CallOption) (WatchProjectClientStream, error)
	WatchProjects(ctx context.Context, in *WatchProjectsRequest, opts ...grpc.CallOption) (WatchProjectsClientStream, error)
	CreateProject(ctx context.Context, in *CreateProjectRequest, opts ...grpc.CallOption) (*project.Project, error)
	UpdateProject(ctx context.Context, in *UpdateProjectRequest, opts ...grpc.CallOption) (*project.Project, error)
	DeleteProject(ctx context.Context, in *DeleteProjectRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	ListMyProjects(ctx context.Context, in *ListMyProjectsRequest, opts ...grpc.CallOption) (*ListMyProjectsResponse, error)
}

type client struct {
	cc grpc.ClientConnInterface
}

func NewProjectServiceClient(cc grpc.ClientConnInterface) ProjectServiceClient {
	return &client{cc}
}

func (c *client) GetProject(ctx context.Context, in *GetProjectRequest, opts ...grpc.CallOption) (*project.Project, error) {
	out := new(project.Project)
	err := c.cc.Invoke(ctx, "/ntt.iam.v1alpha.ProjectService/GetProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) BatchGetProjects(ctx context.Context, in *BatchGetProjectsRequest, opts ...grpc.CallOption) (*BatchGetProjectsResponse, error) {
	out := new(BatchGetProjectsResponse)
	err := c.cc.Invoke(ctx, "/ntt.iam.v1alpha.ProjectService/BatchGetProjects", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) ListProjects(ctx context.Context, in *ListProjectsRequest, opts ...grpc.CallOption) (*ListProjectsResponse, error) {
	out := new(ListProjectsResponse)
	err := c.cc.Invoke(ctx, "/ntt.iam.v1alpha.ProjectService/ListProjects", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) WatchProject(ctx context.Context, in *WatchProjectRequest, opts ...grpc.CallOption) (WatchProjectClientStream, error) {
	stream, err := c.cc.NewStream(ctx,
		&grpc.StreamDesc{
			StreamName:    "WatchProject",
			ServerStreams: true,
		},
		"/ntt.iam.v1alpha.ProjectService/WatchProject", opts...)
	if err != nil {
		return nil, err
	}
	x := &watchProjectWatchProjectClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WatchProjectClientStream interface {
	Recv() (*WatchProjectResponse, error)
	grpc.ClientStream
}

type watchProjectWatchProjectClient struct {
	grpc.ClientStream
}

func (x *watchProjectWatchProjectClient) Recv() (*WatchProjectResponse, error) {
	m := new(WatchProjectResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *client) WatchProjects(ctx context.Context, in *WatchProjectsRequest, opts ...grpc.CallOption) (WatchProjectsClientStream, error) {
	stream, err := c.cc.NewStream(ctx,
		&grpc.StreamDesc{
			StreamName:    "WatchProjects",
			ServerStreams: true,
		},
		"/ntt.iam.v1alpha.ProjectService/WatchProjects", opts...)
	if err != nil {
		return nil, err
	}
	x := &watchProjectsWatchProjectsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WatchProjectsClientStream interface {
	Recv() (*WatchProjectsResponse, error)
	grpc.ClientStream
}

type watchProjectsWatchProjectsClient struct {
	grpc.ClientStream
}

func (x *watchProjectsWatchProjectsClient) Recv() (*WatchProjectsResponse, error) {
	m := new(WatchProjectsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *client) CreateProject(ctx context.Context, in *CreateProjectRequest, opts ...grpc.CallOption) (*project.Project, error) {
	out := new(project.Project)
	err := c.cc.Invoke(ctx, "/ntt.iam.v1alpha.ProjectService/CreateProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) UpdateProject(ctx context.Context, in *UpdateProjectRequest, opts ...grpc.CallOption) (*project.Project, error) {
	out := new(project.Project)
	err := c.cc.Invoke(ctx, "/ntt.iam.v1alpha.ProjectService/UpdateProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) DeleteProject(ctx context.Context, in *DeleteProjectRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/ntt.iam.v1alpha.ProjectService/DeleteProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) ListMyProjects(ctx context.Context, in *ListMyProjectsRequest, opts ...grpc.CallOption) (*ListMyProjectsResponse, error) {
	out := new(ListMyProjectsResponse)
	err := c.cc.Invoke(ctx, "/ntt.iam.v1alpha.ProjectService/ListMyProjects", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
