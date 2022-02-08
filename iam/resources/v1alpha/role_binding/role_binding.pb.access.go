// Code generated by protoc-gen-goten-resource
// Resource: RoleBinding
// DO NOT EDIT!!!

package role_binding

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
	condition "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/condition"
	organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/organization"
	project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/project"
	role "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/role"
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
	_ = &condition.Condition{}
	_ = &organization.Organization{}
	_ = &project.Project{}
	_ = &role.Role{}
)

type RoleBindingAccess interface {
	GetRoleBinding(context.Context, *GetQuery) (*RoleBinding, error)
	BatchGetRoleBindings(context.Context, []*Reference, ...gotenresource.BatchGetOption) error
	QueryRoleBindings(context.Context, *ListQuery) (*QueryResultSnapshot, error)
	WatchRoleBinding(context.Context, *GetQuery, func(*RoleBindingChange) error) error
	WatchRoleBindings(context.Context, *WatchQuery, func(*QueryResultChange) error) error
	SaveRoleBinding(context.Context, *RoleBinding, ...gotenresource.SaveOption) error
	DeleteRoleBinding(context.Context, *Reference, ...gotenresource.DeleteOption) error
}

type anyCastAccess struct {
	RoleBindingAccess
}

func AsAnyCastAccess(access RoleBindingAccess) gotenresource.Access {
	return &anyCastAccess{RoleBindingAccess: access}
}

func (a *anyCastAccess) Get(ctx context.Context, q gotenresource.GetQuery) (gotenresource.Resource, error) {
	if asRoleBindingQuery, ok := q.(*GetQuery); ok {
		return a.GetRoleBinding(ctx, asRoleBindingQuery)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected RoleBinding, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Query(ctx context.Context, q gotenresource.ListQuery) (gotenresource.QueryResultSnapshot, error) {
	if asRoleBindingQuery, ok := q.(*ListQuery); ok {
		return a.QueryRoleBindings(ctx, asRoleBindingQuery)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected RoleBinding, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Search(ctx context.Context, q gotenresource.SearchQuery) (gotenresource.SearchQueryResultSnapshot, error) {
	return nil, status.Errorf(codes.Internal, "Search is not available for RoleBinding")
}

func (a *anyCastAccess) Watch(ctx context.Context, q gotenresource.GetQuery, cb func(ch gotenresource.ResourceChange) error) error {
	if asRoleBindingQuery, ok := q.(*GetQuery); ok {
		return a.WatchRoleBinding(ctx, asRoleBindingQuery, func(change *RoleBindingChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected RoleBinding, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) WatchQuery(ctx context.Context, q gotenresource.WatchQuery, cb func(ch gotenresource.QueryResultChange) error) error {
	if asRoleBindingQuery, ok := q.(*WatchQuery); ok {
		return a.WatchRoleBindings(ctx, asRoleBindingQuery, func(change *QueryResultChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected RoleBinding, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Save(ctx context.Context, res gotenresource.Resource, opts ...gotenresource.SaveOption) error {
	if asRoleBindingRes, ok := res.(*RoleBinding); ok {
		return a.SaveRoleBinding(ctx, asRoleBindingRes, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected RoleBinding, got: %s",
		res.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Delete(ctx context.Context, ref gotenresource.Reference, opts ...gotenresource.DeleteOption) error {
	if asRoleBindingRef, ok := ref.(*Reference); ok {
		return a.DeleteRoleBinding(ctx, asRoleBindingRef, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected RoleBinding, got: %s",
		ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) BatchGet(ctx context.Context, toGet []gotenresource.Reference, opts ...gotenresource.BatchGetOption) error {
	roleBindingRefs := make([]*Reference, 0, len(toGet))
	for _, ref := range toGet {
		if asRoleBindingRef, ok := ref.(*Reference); !ok {
			return status.Errorf(codes.Internal,
				"Unrecognized descriptor, expected RoleBinding, got: %s",
				ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
		} else {
			roleBindingRefs = append(roleBindingRefs, asRoleBindingRef)
		}
	}
	return a.BatchGetRoleBindings(ctx, roleBindingRefs, opts...)
}

func (a *anyCastAccess) GetResourceDescriptors() []gotenresource.Descriptor {
	return []gotenresource.Descriptor{
		GetDescriptor(),
	}
}