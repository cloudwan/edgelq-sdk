// Code generated by protoc-gen-goten-resource
// Resource: ServiceAccount
// DO NOT EDIT!!!

package service_account

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
	_ = &project.Project{}
	_ = &meta.Meta{}
)

type ServiceAccountAccess interface {
	GetServiceAccount(context.Context, *GetQuery, ...gotenresource.GetOption) (*ServiceAccount, error)
	BatchGetServiceAccounts(context.Context, []*Reference, ...gotenresource.BatchGetOption) error
	QueryServiceAccounts(context.Context, *ListQuery, ...gotenresource.QueryOption) (*QueryResultSnapshot, error)
	WatchServiceAccount(context.Context, *GetQuery, func(*ServiceAccountChange) error) error
	WatchServiceAccounts(context.Context, *WatchQuery, func(*QueryResultChange) error) error
	SaveServiceAccount(context.Context, *ServiceAccount, ...gotenresource.SaveOption) error
	DeleteServiceAccount(context.Context, *Reference, ...gotenresource.DeleteOption) error
}

type anyCastAccess struct {
	ServiceAccountAccess
}

func AsAnyCastAccess(access ServiceAccountAccess) gotenresource.Access {
	return &anyCastAccess{ServiceAccountAccess: access}
}

func (a *anyCastAccess) Get(ctx context.Context, q gotenresource.GetQuery, opts ...gotenresource.GetOption) (gotenresource.Resource, error) {
	if asServiceAccountQuery, ok := q.(*GetQuery); ok {
		return a.GetServiceAccount(ctx, asServiceAccountQuery, opts...)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected ServiceAccount, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Query(ctx context.Context, q gotenresource.ListQuery, opts ...gotenresource.QueryOption) (gotenresource.QueryResultSnapshot, error) {
	if asServiceAccountQuery, ok := q.(*ListQuery); ok {
		return a.QueryServiceAccounts(ctx, asServiceAccountQuery, opts...)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected ServiceAccount, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Search(ctx context.Context, q gotenresource.SearchQuery, opts ...gotenresource.QueryOption) (gotenresource.QueryResultSnapshot, error) {
	return nil, status.Errorf(codes.Internal, "Search is not available for ServiceAccount")
}

func (a *anyCastAccess) Watch(ctx context.Context, q gotenresource.GetQuery, cb func(ch gotenresource.ResourceChange) error) error {
	if asServiceAccountQuery, ok := q.(*GetQuery); ok {
		return a.WatchServiceAccount(ctx, asServiceAccountQuery, func(change *ServiceAccountChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected ServiceAccount, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) WatchQuery(ctx context.Context, q gotenresource.WatchQuery, cb func(ch gotenresource.QueryResultChange) error) error {
	if asServiceAccountQuery, ok := q.(*WatchQuery); ok {
		return a.WatchServiceAccounts(ctx, asServiceAccountQuery, func(change *QueryResultChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected ServiceAccount, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Save(ctx context.Context, res gotenresource.Resource, opts ...gotenresource.SaveOption) error {
	if asServiceAccountRes, ok := res.(*ServiceAccount); ok {
		return a.SaveServiceAccount(ctx, asServiceAccountRes, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected ServiceAccount, got: %s",
		res.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Delete(ctx context.Context, ref gotenresource.Reference, opts ...gotenresource.DeleteOption) error {
	if asServiceAccountRef, ok := ref.(*Reference); ok {
		return a.DeleteServiceAccount(ctx, asServiceAccountRef, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected ServiceAccount, got: %s",
		ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) BatchGet(ctx context.Context, toGet []gotenresource.Reference, opts ...gotenresource.BatchGetOption) error {
	serviceAccountRefs := make([]*Reference, 0, len(toGet))
	for _, ref := range toGet {
		if asServiceAccountRef, ok := ref.(*Reference); !ok {
			return status.Errorf(codes.Internal,
				"Unrecognized descriptor, expected ServiceAccount, got: %s",
				ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
		} else {
			serviceAccountRefs = append(serviceAccountRefs, asServiceAccountRef)
		}
	}
	return a.BatchGetServiceAccounts(ctx, serviceAccountRefs, opts...)
}

func (a *anyCastAccess) GetResourceDescriptors() []gotenresource.Descriptor {
	return []gotenresource.Descriptor{
		GetDescriptor(),
	}
}
