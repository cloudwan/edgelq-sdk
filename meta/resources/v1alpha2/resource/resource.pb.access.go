// Code generated by protoc-gen-goten-resource
// Resource: Resource
// DO NOT EDIT!!!

package resource

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
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
	service "github.com/cloudwan/edgelq-sdk/meta/resources/v1alpha2/service"
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
	_ = &ntt_meta.Meta{}
	_ = &service.Service{}
)

type ResourceAccess interface {
	GetResource(context.Context, *GetQuery) (*Resource, error)
	BatchGetResources(context.Context, []*Reference, ...gotenresource.BatchGetOption) error
	QueryResources(context.Context, *ListQuery) (*QueryResultSnapshot, error)
	WatchResource(context.Context, *GetQuery, func(*ResourceChange) error) error
	WatchResources(context.Context, *WatchQuery, func(*QueryResultChange) error) error
	SaveResource(context.Context, *Resource, ...gotenresource.SaveOption) error
	DeleteResource(context.Context, *Reference, ...gotenresource.DeleteOption) error
}

type anyCastAccess struct {
	ResourceAccess
}

func AsAnyCastAccess(access ResourceAccess) gotenresource.Access {
	return &anyCastAccess{ResourceAccess: access}
}

func (a *anyCastAccess) Get(ctx context.Context, q gotenresource.GetQuery) (gotenresource.Resource, error) {
	if asResourceQuery, ok := q.(*GetQuery); ok {
		return a.GetResource(ctx, asResourceQuery)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Resource, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Query(ctx context.Context, q gotenresource.ListQuery) (gotenresource.QueryResultSnapshot, error) {
	if asResourceQuery, ok := q.(*ListQuery); ok {
		return a.QueryResources(ctx, asResourceQuery)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Resource, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Search(ctx context.Context, q gotenresource.SearchQuery) (gotenresource.SearchQueryResultSnapshot, error) {
	return nil, status.Errorf(codes.Internal, "Search is not available for Resource")
}

func (a *anyCastAccess) Watch(ctx context.Context, q gotenresource.GetQuery, cb func(ch gotenresource.ResourceChange) error) error {
	if asResourceQuery, ok := q.(*GetQuery); ok {
		return a.WatchResource(ctx, asResourceQuery, func(change *ResourceChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Resource, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) WatchQuery(ctx context.Context, q gotenresource.WatchQuery, cb func(ch gotenresource.QueryResultChange) error) error {
	if asResourceQuery, ok := q.(*WatchQuery); ok {
		return a.WatchResources(ctx, asResourceQuery, func(change *QueryResultChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Resource, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Save(ctx context.Context, res gotenresource.Resource, opts ...gotenresource.SaveOption) error {
	if asResourceRes, ok := res.(*Resource); ok {
		return a.SaveResource(ctx, asResourceRes, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Resource, got: %s",
		res.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Delete(ctx context.Context, ref gotenresource.Reference, opts ...gotenresource.DeleteOption) error {
	if asResourceRef, ok := ref.(*Reference); ok {
		return a.DeleteResource(ctx, asResourceRef, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Resource, got: %s",
		ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) BatchGet(ctx context.Context, toGet []gotenresource.Reference, opts ...gotenresource.BatchGetOption) error {
	resourceRefs := make([]*Reference, 0, len(toGet))
	for _, ref := range toGet {
		if asResourceRef, ok := ref.(*Reference); !ok {
			return status.Errorf(codes.Internal,
				"Unrecognized descriptor, expected Resource, got: %s",
				ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
		} else {
			resourceRefs = append(resourceRefs, asResourceRef)
		}
	}
	return a.BatchGetResources(ctx, resourceRefs, opts...)
}

func (a *anyCastAccess) GetResourceDescriptors() []gotenresource.Descriptor {
	return []gotenresource.Descriptor{
		GetDescriptor(),
	}
}
