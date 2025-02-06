// Code generated by protoc-gen-goten-resource
// Resource: Project
// DO NOT EDIT!!!

package project

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
	iam_common "github.com/cloudwan/edgelq-sdk/iam/resources/v1/common"
	organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1/organization"
	meta_service "github.com/cloudwan/goten-sdk/meta-service/resources/v1/service"
	meta "github.com/cloudwan/goten-sdk/types/meta"
	multi_region_policy "github.com/cloudwan/goten-sdk/types/multi_region_policy"
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
	_ = &iam_common.PCR{}
	_ = &organization.Organization{}
	_ = &meta_service.Service{}
	_ = &meta.Meta{}
	_ = &multi_region_policy.MultiRegionPolicy{}
)

type ProjectAccess interface {
	GetProject(context.Context, *GetQuery, ...gotenresource.GetOption) (*Project, error)
	BatchGetProjects(context.Context, []*Reference, ...gotenresource.BatchGetOption) error
	QueryProjects(context.Context, *ListQuery, ...gotenresource.QueryOption) (*QueryResultSnapshot, error)
	SearchProjects(context.Context, *SearchQuery, ...gotenresource.QueryOption) (*QueryResultSnapshot, error)
	WatchProject(context.Context, *GetQuery, func(*ProjectChange) error) error
	WatchProjects(context.Context, *WatchQuery, func(*QueryResultChange) error) error
	SaveProject(context.Context, *Project, ...gotenresource.SaveOption) error
	DeleteProject(context.Context, *Reference, ...gotenresource.DeleteOption) error
}

type anyCastAccess struct {
	ProjectAccess
}

func AsAnyCastAccess(access ProjectAccess) gotenresource.Access {
	return &anyCastAccess{ProjectAccess: access}
}

func (a *anyCastAccess) Get(ctx context.Context, q gotenresource.GetQuery, opts ...gotenresource.GetOption) (gotenresource.Resource, error) {
	if asProjectQuery, ok := q.(*GetQuery); ok {
		return a.GetProject(ctx, asProjectQuery, opts...)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Project, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Query(ctx context.Context, q gotenresource.ListQuery, opts ...gotenresource.QueryOption) (gotenresource.QueryResultSnapshot, error) {
	if asProjectQuery, ok := q.(*ListQuery); ok {
		return a.QueryProjects(ctx, asProjectQuery, opts...)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Project, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Search(ctx context.Context, q gotenresource.SearchQuery, opts ...gotenresource.QueryOption) (gotenresource.QueryResultSnapshot, error) {
	if asProjectQuery, ok := q.(*SearchQuery); ok {
		return a.SearchProjects(ctx, asProjectQuery, opts...)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Project, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Watch(ctx context.Context, q gotenresource.GetQuery, cb func(ch gotenresource.ResourceChange) error) error {
	if asProjectQuery, ok := q.(*GetQuery); ok {
		return a.WatchProject(ctx, asProjectQuery, func(change *ProjectChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Project, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) WatchQuery(ctx context.Context, q gotenresource.WatchQuery, cb func(ch gotenresource.QueryResultChange) error) error {
	if asProjectQuery, ok := q.(*WatchQuery); ok {
		return a.WatchProjects(ctx, asProjectQuery, func(change *QueryResultChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Project, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Save(ctx context.Context, res gotenresource.Resource, opts ...gotenresource.SaveOption) error {
	if asProjectRes, ok := res.(*Project); ok {
		return a.SaveProject(ctx, asProjectRes, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Project, got: %s",
		res.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Delete(ctx context.Context, ref gotenresource.Reference, opts ...gotenresource.DeleteOption) error {
	if asProjectRef, ok := ref.(*Reference); ok {
		return a.DeleteProject(ctx, asProjectRef, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Project, got: %s",
		ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) BatchGet(ctx context.Context, toGet []gotenresource.Reference, opts ...gotenresource.BatchGetOption) error {
	projectRefs := make([]*Reference, 0, len(toGet))
	for _, ref := range toGet {
		if asProjectRef, ok := ref.(*Reference); !ok {
			return status.Errorf(codes.Internal,
				"Unrecognized descriptor, expected Project, got: %s",
				ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
		} else {
			projectRefs = append(projectRefs, asProjectRef)
		}
	}
	return a.BatchGetProjects(ctx, projectRefs, opts...)
}

func (a *anyCastAccess) GetResourceDescriptors() []gotenresource.Descriptor {
	return []gotenresource.Descriptor{
		GetDescriptor(),
	}
}
