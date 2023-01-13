// Code generated by protoc-gen-goten-resource
// Resource: AttestationDomain
// DO NOT EDIT!!!

package attestation_domain

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
	iam_common "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/common"
	project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/project"
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
	_ = &iam_common.PCR{}
	_ = &project.Project{}
)

type AttestationDomainAccess interface {
	GetAttestationDomain(context.Context, *GetQuery) (*AttestationDomain, error)
	BatchGetAttestationDomains(context.Context, []*Reference, ...gotenresource.BatchGetOption) error
	QueryAttestationDomains(context.Context, *ListQuery) (*QueryResultSnapshot, error)
	WatchAttestationDomain(context.Context, *GetQuery, func(*AttestationDomainChange) error) error
	WatchAttestationDomains(context.Context, *WatchQuery, func(*QueryResultChange) error) error
	SaveAttestationDomain(context.Context, *AttestationDomain, ...gotenresource.SaveOption) error
	DeleteAttestationDomain(context.Context, *Reference, ...gotenresource.DeleteOption) error
}

type anyCastAccess struct {
	AttestationDomainAccess
}

func AsAnyCastAccess(access AttestationDomainAccess) gotenresource.Access {
	return &anyCastAccess{AttestationDomainAccess: access}
}

func (a *anyCastAccess) Get(ctx context.Context, q gotenresource.GetQuery) (gotenresource.Resource, error) {
	if asAttestationDomainQuery, ok := q.(*GetQuery); ok {
		return a.GetAttestationDomain(ctx, asAttestationDomainQuery)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected AttestationDomain, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Query(ctx context.Context, q gotenresource.ListQuery) (gotenresource.QueryResultSnapshot, error) {
	if asAttestationDomainQuery, ok := q.(*ListQuery); ok {
		return a.QueryAttestationDomains(ctx, asAttestationDomainQuery)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected AttestationDomain, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Search(ctx context.Context, q gotenresource.SearchQuery) (gotenresource.SearchQueryResultSnapshot, error) {
	return nil, status.Errorf(codes.Internal, "Search is not available for AttestationDomain")
}

func (a *anyCastAccess) Watch(ctx context.Context, q gotenresource.GetQuery, cb func(ch gotenresource.ResourceChange) error) error {
	if asAttestationDomainQuery, ok := q.(*GetQuery); ok {
		return a.WatchAttestationDomain(ctx, asAttestationDomainQuery, func(change *AttestationDomainChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected AttestationDomain, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) WatchQuery(ctx context.Context, q gotenresource.WatchQuery, cb func(ch gotenresource.QueryResultChange) error) error {
	if asAttestationDomainQuery, ok := q.(*WatchQuery); ok {
		return a.WatchAttestationDomains(ctx, asAttestationDomainQuery, func(change *QueryResultChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected AttestationDomain, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Save(ctx context.Context, res gotenresource.Resource, opts ...gotenresource.SaveOption) error {
	if asAttestationDomainRes, ok := res.(*AttestationDomain); ok {
		return a.SaveAttestationDomain(ctx, asAttestationDomainRes, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected AttestationDomain, got: %s",
		res.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Delete(ctx context.Context, ref gotenresource.Reference, opts ...gotenresource.DeleteOption) error {
	if asAttestationDomainRef, ok := ref.(*Reference); ok {
		return a.DeleteAttestationDomain(ctx, asAttestationDomainRef, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected AttestationDomain, got: %s",
		ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) BatchGet(ctx context.Context, toGet []gotenresource.Reference, opts ...gotenresource.BatchGetOption) error {
	attestationDomainRefs := make([]*Reference, 0, len(toGet))
	for _, ref := range toGet {
		if asAttestationDomainRef, ok := ref.(*Reference); !ok {
			return status.Errorf(codes.Internal,
				"Unrecognized descriptor, expected AttestationDomain, got: %s",
				ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
		} else {
			attestationDomainRefs = append(attestationDomainRefs, asAttestationDomainRef)
		}
	}
	return a.BatchGetAttestationDomains(ctx, attestationDomainRefs, opts...)
}

func (a *anyCastAccess) GetResourceDescriptors() []gotenresource.Descriptor {
	return []gotenresource.Descriptor{
		GetDescriptor(),
	}
}
