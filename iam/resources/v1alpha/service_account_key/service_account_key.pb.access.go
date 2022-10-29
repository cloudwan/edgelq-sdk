// Code generated by protoc-gen-goten-resource
// Resource: ServiceAccountKey
// DO NOT EDIT!!!

package service_account_key

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
	service_account "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/service_account"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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
	_ = &service_account.ServiceAccount{}
	_ = &timestamp.Timestamp{}
)

type ServiceAccountKeyAccess interface {
	GetServiceAccountKey(context.Context, *GetQuery) (*ServiceAccountKey, error)
	BatchGetServiceAccountKeys(context.Context, []*Reference, ...gotenresource.BatchGetOption) error
	QueryServiceAccountKeys(context.Context, *ListQuery) (*QueryResultSnapshot, error)
	WatchServiceAccountKey(context.Context, *GetQuery, func(*ServiceAccountKeyChange) error) error
	WatchServiceAccountKeys(context.Context, *WatchQuery, func(*QueryResultChange) error) error
	SaveServiceAccountKey(context.Context, *ServiceAccountKey, ...gotenresource.SaveOption) error
	DeleteServiceAccountKey(context.Context, *Reference, ...gotenresource.DeleteOption) error
}

type anyCastAccess struct {
	ServiceAccountKeyAccess
}

func AsAnyCastAccess(access ServiceAccountKeyAccess) gotenresource.Access {
	return &anyCastAccess{ServiceAccountKeyAccess: access}
}

func (a *anyCastAccess) Get(ctx context.Context, q gotenresource.GetQuery) (gotenresource.Resource, error) {
	if asServiceAccountKeyQuery, ok := q.(*GetQuery); ok {
		return a.GetServiceAccountKey(ctx, asServiceAccountKeyQuery)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected ServiceAccountKey, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Query(ctx context.Context, q gotenresource.ListQuery) (gotenresource.QueryResultSnapshot, error) {
	if asServiceAccountKeyQuery, ok := q.(*ListQuery); ok {
		return a.QueryServiceAccountKeys(ctx, asServiceAccountKeyQuery)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected ServiceAccountKey, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Search(ctx context.Context, q gotenresource.SearchQuery) (gotenresource.SearchQueryResultSnapshot, error) {
	return nil, status.Errorf(codes.Internal, "Search is not available for ServiceAccountKey")
}

func (a *anyCastAccess) Watch(ctx context.Context, q gotenresource.GetQuery, cb func(ch gotenresource.ResourceChange) error) error {
	if asServiceAccountKeyQuery, ok := q.(*GetQuery); ok {
		return a.WatchServiceAccountKey(ctx, asServiceAccountKeyQuery, func(change *ServiceAccountKeyChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected ServiceAccountKey, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) WatchQuery(ctx context.Context, q gotenresource.WatchQuery, cb func(ch gotenresource.QueryResultChange) error) error {
	if asServiceAccountKeyQuery, ok := q.(*WatchQuery); ok {
		return a.WatchServiceAccountKeys(ctx, asServiceAccountKeyQuery, func(change *QueryResultChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected ServiceAccountKey, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Save(ctx context.Context, res gotenresource.Resource, opts ...gotenresource.SaveOption) error {
	if asServiceAccountKeyRes, ok := res.(*ServiceAccountKey); ok {
		return a.SaveServiceAccountKey(ctx, asServiceAccountKeyRes, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected ServiceAccountKey, got: %s",
		res.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Delete(ctx context.Context, ref gotenresource.Reference, opts ...gotenresource.DeleteOption) error {
	if asServiceAccountKeyRef, ok := ref.(*Reference); ok {
		return a.DeleteServiceAccountKey(ctx, asServiceAccountKeyRef, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected ServiceAccountKey, got: %s",
		ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) BatchGet(ctx context.Context, toGet []gotenresource.Reference, opts ...gotenresource.BatchGetOption) error {
	serviceAccountKeyRefs := make([]*Reference, 0, len(toGet))
	for _, ref := range toGet {
		if asServiceAccountKeyRef, ok := ref.(*Reference); !ok {
			return status.Errorf(codes.Internal,
				"Unrecognized descriptor, expected ServiceAccountKey, got: %s",
				ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
		} else {
			serviceAccountKeyRefs = append(serviceAccountKeyRefs, asServiceAccountKeyRef)
		}
	}
	return a.BatchGetServiceAccountKeys(ctx, serviceAccountKeyRefs, opts...)
}

func (a *anyCastAccess) GetResourceDescriptors() []gotenresource.Descriptor {
	return []gotenresource.Descriptor{
		GetDescriptor(),
	}
}
