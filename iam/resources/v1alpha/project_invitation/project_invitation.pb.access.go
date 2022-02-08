// Code generated by protoc-gen-goten-resource
// Resource: ProjectInvitation
// DO NOT EDIT!!!

package project_invitation

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cloudwan/goten-sdk/runtime/api/watch_type"
	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	iam_common "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/common"
	project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/project"
)

// ensure the imports are used
var (
	_ = context.Context(nil)

	_ = codes.Internal
	_ = status.Status{}

	_ = watch_type.WatchType_STATEFUL
	_ = gotenobject.FieldPath(nil)
	_ = gotenresource.ListQuery(nil)
)

// make sure we're using proto imports
var (
	_ = &iam_common.Actor{}
	_ = &project.Project{}
)

type ProjectInvitationAccess interface {
	GetProjectInvitation(context.Context, *GetQuery) (*ProjectInvitation, error)
	BatchGetProjectInvitations(context.Context, []*Reference, ...gotenresource.BatchGetOption) error
	QueryProjectInvitations(context.Context, *ListQuery) (*QueryResultSnapshot, error)
	WatchProjectInvitation(context.Context, *GetQuery, func(*ProjectInvitationChange) error) error
	WatchProjectInvitations(context.Context, *WatchQuery, func(*QueryResultChange) error) error
	SaveProjectInvitation(context.Context, *ProjectInvitation, ...gotenresource.SaveOption) error
	DeleteProjectInvitation(context.Context, *Reference, ...gotenresource.DeleteOption) error
}

type anyCastAccess struct {
	ProjectInvitationAccess
}

func AsAnyCastAccess(access ProjectInvitationAccess) gotenresource.Access {
	return &anyCastAccess{ProjectInvitationAccess: access}
}

func (a *anyCastAccess) Get(ctx context.Context, q gotenresource.GetQuery) (gotenresource.Resource, error) {
	if asProjectInvitationQuery, ok := q.(*GetQuery); ok {
		return a.GetProjectInvitation(ctx, asProjectInvitationQuery)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected ProjectInvitation, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Query(ctx context.Context, q gotenresource.ListQuery) (gotenresource.QueryResultSnapshot, error) {
	if asProjectInvitationQuery, ok := q.(*ListQuery); ok {
		return a.QueryProjectInvitations(ctx, asProjectInvitationQuery)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected ProjectInvitation, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Search(ctx context.Context, q gotenresource.SearchQuery) (gotenresource.SearchQueryResultSnapshot, error) {
	return nil, status.Errorf(codes.Internal, "Search is not available for ProjectInvitation")
}

func (a *anyCastAccess) Watch(ctx context.Context, q gotenresource.GetQuery, cb func(ch gotenresource.ResourceChange) error) error {
	if asProjectInvitationQuery, ok := q.(*GetQuery); ok {
		return a.WatchProjectInvitation(ctx, asProjectInvitationQuery, func(change *ProjectInvitationChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected ProjectInvitation, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) WatchQuery(ctx context.Context, q gotenresource.WatchQuery, cb func(ch gotenresource.QueryResultChange) error) error {
	if asProjectInvitationQuery, ok := q.(*WatchQuery); ok {
		return a.WatchProjectInvitations(ctx, asProjectInvitationQuery, func(change *QueryResultChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected ProjectInvitation, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Save(ctx context.Context, res gotenresource.Resource, opts ...gotenresource.SaveOption) error {
	if asProjectInvitationRes, ok := res.(*ProjectInvitation); ok {
		return a.SaveProjectInvitation(ctx, asProjectInvitationRes, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected ProjectInvitation, got: %s",
		res.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Delete(ctx context.Context, ref gotenresource.Reference, opts ...gotenresource.DeleteOption) error {
	if asProjectInvitationRef, ok := ref.(*Reference); ok {
		return a.DeleteProjectInvitation(ctx, asProjectInvitationRef, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected ProjectInvitation, got: %s",
		ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) BatchGet(ctx context.Context, toGet []gotenresource.Reference, opts ...gotenresource.BatchGetOption) error {
	projectInvitationRefs := make([]*Reference, 0, len(toGet))
	for _, ref := range toGet {
		if asProjectInvitationRef, ok := ref.(*Reference); !ok {
			return status.Errorf(codes.Internal,
				"Unrecognized descriptor, expected ProjectInvitation, got: %s",
				ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
		} else {
			projectInvitationRefs = append(projectInvitationRefs, asProjectInvitationRef)
		}
	}
	return a.BatchGetProjectInvitations(ctx, projectInvitationRefs, opts...)
}

func (a *anyCastAccess) GetResourceDescriptors() []gotenresource.Descriptor {
	return []gotenresource.Descriptor{
		GetDescriptor(),
	}
}