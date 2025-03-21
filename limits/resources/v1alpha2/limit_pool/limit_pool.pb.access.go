// Code generated by protoc-gen-goten-resource
// Resource: LimitPool
// DO NOT EDIT!!!

package limit_pool

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
	iam_organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/organization"
	meta_resource "github.com/cloudwan/edgelq-sdk/meta/resources/v1alpha2/resource"
	meta_service "github.com/cloudwan/edgelq-sdk/meta/resources/v1alpha2/service"
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
	_ = &iam_organization.Organization{}
	_ = &meta_resource.Resource{}
	_ = &meta_service.Service{}
	_ = &meta.Meta{}
)

type LimitPoolAccess interface {
	GetLimitPool(context.Context, *GetQuery, ...gotenresource.GetOption) (*LimitPool, error)
	BatchGetLimitPools(context.Context, []*Reference, ...gotenresource.BatchGetOption) error
	QueryLimitPools(context.Context, *ListQuery, ...gotenresource.QueryOption) (*QueryResultSnapshot, error)
	WatchLimitPool(context.Context, *GetQuery, func(*LimitPoolChange) error) error
	WatchLimitPools(context.Context, *WatchQuery, func(*QueryResultChange) error) error
	SaveLimitPool(context.Context, *LimitPool, ...gotenresource.SaveOption) error
	DeleteLimitPool(context.Context, *Reference, ...gotenresource.DeleteOption) error
}

type anyCastAccess struct {
	LimitPoolAccess
}

func AsAnyCastAccess(access LimitPoolAccess) gotenresource.Access {
	return &anyCastAccess{LimitPoolAccess: access}
}

func (a *anyCastAccess) Get(ctx context.Context, q gotenresource.GetQuery, opts ...gotenresource.GetOption) (gotenresource.Resource, error) {
	if asLimitPoolQuery, ok := q.(*GetQuery); ok {
		return a.GetLimitPool(ctx, asLimitPoolQuery, opts...)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected LimitPool, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Query(ctx context.Context, q gotenresource.ListQuery, opts ...gotenresource.QueryOption) (gotenresource.QueryResultSnapshot, error) {
	if asLimitPoolQuery, ok := q.(*ListQuery); ok {
		return a.QueryLimitPools(ctx, asLimitPoolQuery, opts...)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected LimitPool, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Search(ctx context.Context, q gotenresource.SearchQuery, opts ...gotenresource.QueryOption) (gotenresource.QueryResultSnapshot, error) {
	return nil, status.Errorf(codes.Internal, "Search is not available for LimitPool")
}

func (a *anyCastAccess) Watch(ctx context.Context, q gotenresource.GetQuery, cb func(ch gotenresource.ResourceChange) error) error {
	if asLimitPoolQuery, ok := q.(*GetQuery); ok {
		return a.WatchLimitPool(ctx, asLimitPoolQuery, func(change *LimitPoolChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected LimitPool, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) WatchQuery(ctx context.Context, q gotenresource.WatchQuery, cb func(ch gotenresource.QueryResultChange) error) error {
	if asLimitPoolQuery, ok := q.(*WatchQuery); ok {
		return a.WatchLimitPools(ctx, asLimitPoolQuery, func(change *QueryResultChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected LimitPool, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Save(ctx context.Context, res gotenresource.Resource, opts ...gotenresource.SaveOption) error {
	if asLimitPoolRes, ok := res.(*LimitPool); ok {
		return a.SaveLimitPool(ctx, asLimitPoolRes, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected LimitPool, got: %s",
		res.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Delete(ctx context.Context, ref gotenresource.Reference, opts ...gotenresource.DeleteOption) error {
	if asLimitPoolRef, ok := ref.(*Reference); ok {
		return a.DeleteLimitPool(ctx, asLimitPoolRef, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected LimitPool, got: %s",
		ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) BatchGet(ctx context.Context, toGet []gotenresource.Reference, opts ...gotenresource.BatchGetOption) error {
	limitPoolRefs := make([]*Reference, 0, len(toGet))
	for _, ref := range toGet {
		if asLimitPoolRef, ok := ref.(*Reference); !ok {
			return status.Errorf(codes.Internal,
				"Unrecognized descriptor, expected LimitPool, got: %s",
				ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
		} else {
			limitPoolRefs = append(limitPoolRefs, asLimitPoolRef)
		}
	}
	return a.BatchGetLimitPools(ctx, limitPoolRefs, opts...)
}

func (a *anyCastAccess) GetResourceDescriptors() []gotenresource.Descriptor {
	return []gotenresource.Descriptor{
		GetDescriptor(),
	}
}
