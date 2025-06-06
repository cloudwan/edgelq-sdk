// Code generated by protoc-gen-goten-resource
// Resource: Group
// DO NOT EDIT!!!

package group

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
	"github.com/cloudwan/goten-sdk/types/watch_type"
)

// proto imports
import (
	organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/organization"
	project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/project"
	meta "github.com/cloudwan/goten-sdk/types/meta"
)

// ensure the imports are used
var (
	_ = new(context.Context)

	_ = codes.Internal
	_ = status.Status{}

	_ = watch_type.WatchType_STATEFUL
	_ = new(gotenobject.FieldPath)
	_ = new(gotenresource.ListQuery)
)

// make sure we're using proto imports
var (
	_ = &organization.Organization{}
	_ = &project.Project{}
	_ = &meta.Meta{}
)

type GroupAccess interface {
	GetGroup(context.Context, *GetQuery, ...gotenresource.GetOption) (*Group, error)
	BatchGetGroups(context.Context, []*Reference, ...gotenresource.BatchGetOption) error
	QueryGroups(context.Context, *ListQuery, ...gotenresource.QueryOption) (*QueryResultSnapshot, error)
	WatchGroup(context.Context, *GetQuery, func(*GroupChange) error) error
	WatchGroups(context.Context, *WatchQuery, func(*QueryResultChange) error) error
	SaveGroup(context.Context, *Group, ...gotenresource.SaveOption) error
	DeleteGroup(context.Context, *Reference, ...gotenresource.DeleteOption) error
}

type anyCastAccess struct {
	GroupAccess
}

func AsAnyCastAccess(access GroupAccess) gotenresource.Access {
	return &anyCastAccess{GroupAccess: access}
}

func (a *anyCastAccess) Get(ctx context.Context, q gotenresource.GetQuery, opts ...gotenresource.GetOption) (gotenresource.Resource, error) {
	if asGroupQuery, ok := q.(*GetQuery); ok {
		return a.GetGroup(ctx, asGroupQuery, opts...)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Group, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Query(ctx context.Context, q gotenresource.ListQuery, opts ...gotenresource.QueryOption) (gotenresource.QueryResultSnapshot, error) {
	if asGroupQuery, ok := q.(*ListQuery); ok {
		return a.QueryGroups(ctx, asGroupQuery, opts...)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Group, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Search(ctx context.Context, q gotenresource.SearchQuery, opts ...gotenresource.QueryOption) (gotenresource.QueryResultSnapshot, error) {
	return nil, status.Errorf(codes.Internal, "Search is not available for Group")
}

func (a *anyCastAccess) Watch(ctx context.Context, q gotenresource.GetQuery, cb func(ch gotenresource.ResourceChange) error) error {
	if asGroupQuery, ok := q.(*GetQuery); ok {
		return a.WatchGroup(ctx, asGroupQuery, func(change *GroupChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Group, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) WatchQuery(ctx context.Context, q gotenresource.WatchQuery, cb func(ch gotenresource.QueryResultChange) error) error {
	if asGroupQuery, ok := q.(*WatchQuery); ok {
		return a.WatchGroups(ctx, asGroupQuery, func(change *QueryResultChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Group, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Save(ctx context.Context, res gotenresource.Resource, opts ...gotenresource.SaveOption) error {
	if asGroupRes, ok := res.(*Group); ok {
		return a.SaveGroup(ctx, asGroupRes, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Group, got: %s",
		res.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Delete(ctx context.Context, ref gotenresource.Reference, opts ...gotenresource.DeleteOption) error {
	if asGroupRef, ok := ref.(*Reference); ok {
		return a.DeleteGroup(ctx, asGroupRef, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Group, got: %s",
		ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) BatchGet(ctx context.Context, toGet []gotenresource.Reference, opts ...gotenresource.BatchGetOption) error {
	groupRefs := make([]*Reference, 0, len(toGet))
	for _, ref := range toGet {
		if asGroupRef, ok := ref.(*Reference); !ok {
			return status.Errorf(codes.Internal,
				"Unrecognized descriptor, expected Group, got: %s",
				ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
		} else {
			groupRefs = append(groupRefs, asGroupRef)
		}
	}
	return a.BatchGetGroups(ctx, groupRefs, opts...)
}

func (a *anyCastAccess) GetResourceDescriptors() []gotenresource.Descriptor {
	return []gotenresource.Descriptor{
		GetDescriptor(),
	}
}
