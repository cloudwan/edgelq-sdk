// Code generated by protoc-gen-goten-resource
// Resource: Service
// DO NOT EDIT!!!

package service

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
	_ = &meta.Meta{}
)

type ServiceAccess interface {
	GetService(context.Context, *GetQuery, ...gotenresource.GetOption) (*Service, error)
	BatchGetServices(context.Context, []*Reference, ...gotenresource.BatchGetOption) error
	QueryServices(context.Context, *ListQuery, ...gotenresource.QueryOption) (*QueryResultSnapshot, error)
	WatchService(context.Context, *GetQuery, func(*ServiceChange) error) error
	WatchServices(context.Context, *WatchQuery, func(*QueryResultChange) error) error
	SaveService(context.Context, *Service, ...gotenresource.SaveOption) error
	DeleteService(context.Context, *Reference, ...gotenresource.DeleteOption) error
}

type anyCastAccess struct {
	ServiceAccess
}

func AsAnyCastAccess(access ServiceAccess) gotenresource.Access {
	return &anyCastAccess{ServiceAccess: access}
}

func (a *anyCastAccess) Get(ctx context.Context, q gotenresource.GetQuery, opts ...gotenresource.GetOption) (gotenresource.Resource, error) {
	if asServiceQuery, ok := q.(*GetQuery); ok {
		return a.GetService(ctx, asServiceQuery, opts...)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Service, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Query(ctx context.Context, q gotenresource.ListQuery, opts ...gotenresource.QueryOption) (gotenresource.QueryResultSnapshot, error) {
	if asServiceQuery, ok := q.(*ListQuery); ok {
		return a.QueryServices(ctx, asServiceQuery, opts...)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Service, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Search(ctx context.Context, q gotenresource.SearchQuery, opts ...gotenresource.QueryOption) (gotenresource.QueryResultSnapshot, error) {
	return nil, status.Errorf(codes.Internal, "Search is not available for Service")
}

func (a *anyCastAccess) Watch(ctx context.Context, q gotenresource.GetQuery, cb func(ch gotenresource.ResourceChange) error) error {
	if asServiceQuery, ok := q.(*GetQuery); ok {
		return a.WatchService(ctx, asServiceQuery, func(change *ServiceChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Service, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) WatchQuery(ctx context.Context, q gotenresource.WatchQuery, cb func(ch gotenresource.QueryResultChange) error) error {
	if asServiceQuery, ok := q.(*WatchQuery); ok {
		return a.WatchServices(ctx, asServiceQuery, func(change *QueryResultChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Service, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Save(ctx context.Context, res gotenresource.Resource, opts ...gotenresource.SaveOption) error {
	if asServiceRes, ok := res.(*Service); ok {
		return a.SaveService(ctx, asServiceRes, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Service, got: %s",
		res.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Delete(ctx context.Context, ref gotenresource.Reference, opts ...gotenresource.DeleteOption) error {
	if asServiceRef, ok := ref.(*Reference); ok {
		return a.DeleteService(ctx, asServiceRef, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Service, got: %s",
		ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) BatchGet(ctx context.Context, toGet []gotenresource.Reference, opts ...gotenresource.BatchGetOption) error {
	serviceRefs := make([]*Reference, 0, len(toGet))
	for _, ref := range toGet {
		if asServiceRef, ok := ref.(*Reference); !ok {
			return status.Errorf(codes.Internal,
				"Unrecognized descriptor, expected Service, got: %s",
				ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
		} else {
			serviceRefs = append(serviceRefs, asServiceRef)
		}
	}
	return a.BatchGetServices(ctx, serviceRefs, opts...)
}

func (a *anyCastAccess) GetResourceDescriptors() []gotenresource.Descriptor {
	return []gotenresource.Descriptor{
		GetDescriptor(),
	}
}
