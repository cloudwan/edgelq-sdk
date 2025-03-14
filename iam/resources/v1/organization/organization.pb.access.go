// Code generated by protoc-gen-goten-resource
// Resource: Organization
// DO NOT EDIT!!!

package organization

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
	_ = &meta_service.Service{}
	_ = &meta.Meta{}
	_ = &multi_region_policy.MultiRegionPolicy{}
)

type OrganizationAccess interface {
	GetOrganization(context.Context, *GetQuery, ...gotenresource.GetOption) (*Organization, error)
	BatchGetOrganizations(context.Context, []*Reference, ...gotenresource.BatchGetOption) error
	QueryOrganizations(context.Context, *ListQuery, ...gotenresource.QueryOption) (*QueryResultSnapshot, error)
	SearchOrganizations(context.Context, *SearchQuery, ...gotenresource.QueryOption) (*QueryResultSnapshot, error)
	WatchOrganization(context.Context, *GetQuery, func(*OrganizationChange) error) error
	WatchOrganizations(context.Context, *WatchQuery, func(*QueryResultChange) error) error
	SaveOrganization(context.Context, *Organization, ...gotenresource.SaveOption) error
	DeleteOrganization(context.Context, *Reference, ...gotenresource.DeleteOption) error
}

type anyCastAccess struct {
	OrganizationAccess
}

func AsAnyCastAccess(access OrganizationAccess) gotenresource.Access {
	return &anyCastAccess{OrganizationAccess: access}
}

func (a *anyCastAccess) Get(ctx context.Context, q gotenresource.GetQuery, opts ...gotenresource.GetOption) (gotenresource.Resource, error) {
	if asOrganizationQuery, ok := q.(*GetQuery); ok {
		return a.GetOrganization(ctx, asOrganizationQuery, opts...)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Organization, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Query(ctx context.Context, q gotenresource.ListQuery, opts ...gotenresource.QueryOption) (gotenresource.QueryResultSnapshot, error) {
	if asOrganizationQuery, ok := q.(*ListQuery); ok {
		return a.QueryOrganizations(ctx, asOrganizationQuery, opts...)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Organization, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Search(ctx context.Context, q gotenresource.SearchQuery, opts ...gotenresource.QueryOption) (gotenresource.QueryResultSnapshot, error) {
	if asOrganizationQuery, ok := q.(*SearchQuery); ok {
		return a.SearchOrganizations(ctx, asOrganizationQuery, opts...)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Organization, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Watch(ctx context.Context, q gotenresource.GetQuery, cb func(ch gotenresource.ResourceChange) error) error {
	if asOrganizationQuery, ok := q.(*GetQuery); ok {
		return a.WatchOrganization(ctx, asOrganizationQuery, func(change *OrganizationChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Organization, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) WatchQuery(ctx context.Context, q gotenresource.WatchQuery, cb func(ch gotenresource.QueryResultChange) error) error {
	if asOrganizationQuery, ok := q.(*WatchQuery); ok {
		return a.WatchOrganizations(ctx, asOrganizationQuery, func(change *QueryResultChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Organization, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Save(ctx context.Context, res gotenresource.Resource, opts ...gotenresource.SaveOption) error {
	if asOrganizationRes, ok := res.(*Organization); ok {
		return a.SaveOrganization(ctx, asOrganizationRes, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Organization, got: %s",
		res.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Delete(ctx context.Context, ref gotenresource.Reference, opts ...gotenresource.DeleteOption) error {
	if asOrganizationRef, ok := ref.(*Reference); ok {
		return a.DeleteOrganization(ctx, asOrganizationRef, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Organization, got: %s",
		ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) BatchGet(ctx context.Context, toGet []gotenresource.Reference, opts ...gotenresource.BatchGetOption) error {
	organizationRefs := make([]*Reference, 0, len(toGet))
	for _, ref := range toGet {
		if asOrganizationRef, ok := ref.(*Reference); !ok {
			return status.Errorf(codes.Internal,
				"Unrecognized descriptor, expected Organization, got: %s",
				ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
		} else {
			organizationRefs = append(organizationRefs, asOrganizationRef)
		}
	}
	return a.BatchGetOrganizations(ctx, organizationRefs, opts...)
}

func (a *anyCastAccess) GetResourceDescriptors() []gotenresource.Descriptor {
	return []gotenresource.Descriptor{
		GetDescriptor(),
	}
}
