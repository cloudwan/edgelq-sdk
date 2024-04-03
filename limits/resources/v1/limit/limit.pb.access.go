// Code generated by protoc-gen-goten-resource
// Resource: Limit
// DO NOT EDIT!!!

package limit

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
	iam_project "github.com/cloudwan/edgelq-sdk/iam/resources/v1/project"
	limit_pool "github.com/cloudwan/edgelq-sdk/limits/resources/v1/limit_pool"
	meta_resource "github.com/cloudwan/goten-sdk/meta-service/resources/v1/resource"
	meta_service "github.com/cloudwan/goten-sdk/meta-service/resources/v1/service"
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
	_ = &iam_project.Project{}
	_ = &limit_pool.LimitPool{}
	_ = &meta_resource.Resource{}
	_ = &meta_service.Service{}
	_ = &meta.Meta{}
)

type LimitAccess interface {
	GetLimit(context.Context, *GetQuery) (*Limit, error)
	BatchGetLimits(context.Context, []*Reference, ...gotenresource.BatchGetOption) error
	QueryLimits(context.Context, *ListQuery) (*QueryResultSnapshot, error)
	WatchLimit(context.Context, *GetQuery, func(*LimitChange) error) error
	WatchLimits(context.Context, *WatchQuery, func(*QueryResultChange) error) error
	SaveLimit(context.Context, *Limit, ...gotenresource.SaveOption) error
	DeleteLimit(context.Context, *Reference, ...gotenresource.DeleteOption) error
}

type anyCastAccess struct {
	LimitAccess
}

func AsAnyCastAccess(access LimitAccess) gotenresource.Access {
	return &anyCastAccess{LimitAccess: access}
}

func (a *anyCastAccess) Get(ctx context.Context, q gotenresource.GetQuery) (gotenresource.Resource, error) {
	if asLimitQuery, ok := q.(*GetQuery); ok {
		return a.GetLimit(ctx, asLimitQuery)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Limit, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Query(ctx context.Context, q gotenresource.ListQuery) (gotenresource.QueryResultSnapshot, error) {
	if asLimitQuery, ok := q.(*ListQuery); ok {
		return a.QueryLimits(ctx, asLimitQuery)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Limit, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Search(ctx context.Context, q gotenresource.SearchQuery) (gotenresource.QueryResultSnapshot, error) {
	return nil, status.Errorf(codes.Internal, "Search is not available for Limit")
}

func (a *anyCastAccess) Watch(ctx context.Context, q gotenresource.GetQuery, cb func(ch gotenresource.ResourceChange) error) error {
	if asLimitQuery, ok := q.(*GetQuery); ok {
		return a.WatchLimit(ctx, asLimitQuery, func(change *LimitChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Limit, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) WatchQuery(ctx context.Context, q gotenresource.WatchQuery, cb func(ch gotenresource.QueryResultChange) error) error {
	if asLimitQuery, ok := q.(*WatchQuery); ok {
		return a.WatchLimits(ctx, asLimitQuery, func(change *QueryResultChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Limit, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Save(ctx context.Context, res gotenresource.Resource, opts ...gotenresource.SaveOption) error {
	if asLimitRes, ok := res.(*Limit); ok {
		return a.SaveLimit(ctx, asLimitRes, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Limit, got: %s",
		res.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Delete(ctx context.Context, ref gotenresource.Reference, opts ...gotenresource.DeleteOption) error {
	if asLimitRef, ok := ref.(*Reference); ok {
		return a.DeleteLimit(ctx, asLimitRef, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Limit, got: %s",
		ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) BatchGet(ctx context.Context, toGet []gotenresource.Reference, opts ...gotenresource.BatchGetOption) error {
	limitRefs := make([]*Reference, 0, len(toGet))
	for _, ref := range toGet {
		if asLimitRef, ok := ref.(*Reference); !ok {
			return status.Errorf(codes.Internal,
				"Unrecognized descriptor, expected Limit, got: %s",
				ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
		} else {
			limitRefs = append(limitRefs, asLimitRef)
		}
	}
	return a.BatchGetLimits(ctx, limitRefs, opts...)
}

func (a *anyCastAccess) GetResourceDescriptors() []gotenresource.Descriptor {
	return []gotenresource.Descriptor{
		GetDescriptor(),
	}
}