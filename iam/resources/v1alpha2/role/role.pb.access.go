// Code generated by protoc-gen-goten-resource
// Resource: Role
// DO NOT EDIT!!!

package role

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
	condition "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/condition"
	permission "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/permission"
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
	_ = &ntt_meta.Meta{}
	_ = &condition.Condition{}
	_ = &permission.Permission{}
)

type RoleAccess interface {
	GetRole(context.Context, *GetQuery) (*Role, error)
	BatchGetRoles(context.Context, []*Reference, ...gotenresource.BatchGetOption) error
	QueryRoles(context.Context, *ListQuery) (*QueryResultSnapshot, error)
	WatchRole(context.Context, *GetQuery, func(*RoleChange) error) error
	WatchRoles(context.Context, *WatchQuery, func(*QueryResultChange) error) error
	SaveRole(context.Context, *Role, ...gotenresource.SaveOption) error
	DeleteRole(context.Context, *Reference, ...gotenresource.DeleteOption) error
}

type anyCastAccess struct {
	RoleAccess
}

func AsAnyCastAccess(access RoleAccess) gotenresource.Access {
	return &anyCastAccess{RoleAccess: access}
}

func (a *anyCastAccess) Get(ctx context.Context, q gotenresource.GetQuery) (gotenresource.Resource, error) {
	if asRoleQuery, ok := q.(*GetQuery); ok {
		return a.GetRole(ctx, asRoleQuery)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Role, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Query(ctx context.Context, q gotenresource.ListQuery) (gotenresource.QueryResultSnapshot, error) {
	if asRoleQuery, ok := q.(*ListQuery); ok {
		return a.QueryRoles(ctx, asRoleQuery)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Role, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Search(ctx context.Context, q gotenresource.SearchQuery) (gotenresource.SearchQueryResultSnapshot, error) {
	return nil, status.Errorf(codes.Internal, "Search is not available for Role")
}

func (a *anyCastAccess) Watch(ctx context.Context, q gotenresource.GetQuery, cb func(ch gotenresource.ResourceChange) error) error {
	if asRoleQuery, ok := q.(*GetQuery); ok {
		return a.WatchRole(ctx, asRoleQuery, func(change *RoleChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Role, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) WatchQuery(ctx context.Context, q gotenresource.WatchQuery, cb func(ch gotenresource.QueryResultChange) error) error {
	if asRoleQuery, ok := q.(*WatchQuery); ok {
		return a.WatchRoles(ctx, asRoleQuery, func(change *QueryResultChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Role, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Save(ctx context.Context, res gotenresource.Resource, opts ...gotenresource.SaveOption) error {
	if asRoleRes, ok := res.(*Role); ok {
		return a.SaveRole(ctx, asRoleRes, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Role, got: %s",
		res.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Delete(ctx context.Context, ref gotenresource.Reference, opts ...gotenresource.DeleteOption) error {
	if asRoleRef, ok := ref.(*Reference); ok {
		return a.DeleteRole(ctx, asRoleRef, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Role, got: %s",
		ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) BatchGet(ctx context.Context, toGet []gotenresource.Reference, opts ...gotenresource.BatchGetOption) error {
	roleRefs := make([]*Reference, 0, len(toGet))
	for _, ref := range toGet {
		if asRoleRef, ok := ref.(*Reference); !ok {
			return status.Errorf(codes.Internal,
				"Unrecognized descriptor, expected Role, got: %s",
				ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
		} else {
			roleRefs = append(roleRefs, asRoleRef)
		}
	}
	return a.BatchGetRoles(ctx, roleRefs, opts...)
}

func (a *anyCastAccess) GetResourceDescriptors() []gotenresource.Descriptor {
	return []gotenresource.Descriptor{
		GetDescriptor(),
	}
}
