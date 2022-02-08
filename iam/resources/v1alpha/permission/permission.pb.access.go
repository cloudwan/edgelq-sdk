// Code generated by protoc-gen-goten-resource
// Resource: Permission
// DO NOT EDIT!!!

package permission

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cloudwan/goten-sdk/runtime/api/watch_type"
	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import ()

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
var ()

type PermissionAccess interface {
	GetPermission(context.Context, *GetQuery) (*Permission, error)
	BatchGetPermissions(context.Context, []*Reference, ...gotenresource.BatchGetOption) error
	QueryPermissions(context.Context, *ListQuery) (*QueryResultSnapshot, error)
	WatchPermission(context.Context, *GetQuery, func(*PermissionChange) error) error
	WatchPermissions(context.Context, *WatchQuery, func(*QueryResultChange) error) error
	SavePermission(context.Context, *Permission, ...gotenresource.SaveOption) error
	DeletePermission(context.Context, *Reference, ...gotenresource.DeleteOption) error
}

type anyCastAccess struct {
	PermissionAccess
}

func AsAnyCastAccess(access PermissionAccess) gotenresource.Access {
	return &anyCastAccess{PermissionAccess: access}
}

func (a *anyCastAccess) Get(ctx context.Context, q gotenresource.GetQuery) (gotenresource.Resource, error) {
	if asPermissionQuery, ok := q.(*GetQuery); ok {
		return a.GetPermission(ctx, asPermissionQuery)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Permission, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Query(ctx context.Context, q gotenresource.ListQuery) (gotenresource.QueryResultSnapshot, error) {
	if asPermissionQuery, ok := q.(*ListQuery); ok {
		return a.QueryPermissions(ctx, asPermissionQuery)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Permission, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Search(ctx context.Context, q gotenresource.SearchQuery) (gotenresource.SearchQueryResultSnapshot, error) {
	return nil, status.Errorf(codes.Internal, "Search is not available for Permission")
}

func (a *anyCastAccess) Watch(ctx context.Context, q gotenresource.GetQuery, cb func(ch gotenresource.ResourceChange) error) error {
	if asPermissionQuery, ok := q.(*GetQuery); ok {
		return a.WatchPermission(ctx, asPermissionQuery, func(change *PermissionChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Permission, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) WatchQuery(ctx context.Context, q gotenresource.WatchQuery, cb func(ch gotenresource.QueryResultChange) error) error {
	if asPermissionQuery, ok := q.(*WatchQuery); ok {
		return a.WatchPermissions(ctx, asPermissionQuery, func(change *QueryResultChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Permission, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Save(ctx context.Context, res gotenresource.Resource, opts ...gotenresource.SaveOption) error {
	if asPermissionRes, ok := res.(*Permission); ok {
		return a.SavePermission(ctx, asPermissionRes, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Permission, got: %s",
		res.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Delete(ctx context.Context, ref gotenresource.Reference, opts ...gotenresource.DeleteOption) error {
	if asPermissionRef, ok := ref.(*Reference); ok {
		return a.DeletePermission(ctx, asPermissionRef, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Permission, got: %s",
		ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) BatchGet(ctx context.Context, toGet []gotenresource.Reference, opts ...gotenresource.BatchGetOption) error {
	permissionRefs := make([]*Reference, 0, len(toGet))
	for _, ref := range toGet {
		if asPermissionRef, ok := ref.(*Reference); !ok {
			return status.Errorf(codes.Internal,
				"Unrecognized descriptor, expected Permission, got: %s",
				ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
		} else {
			permissionRefs = append(permissionRefs, asPermissionRef)
		}
	}
	return a.BatchGetPermissions(ctx, permissionRefs, opts...)
}

func (a *anyCastAccess) GetResourceDescriptors() []gotenresource.Descriptor {
	return []gotenresource.Descriptor{
		GetDescriptor(),
	}
}