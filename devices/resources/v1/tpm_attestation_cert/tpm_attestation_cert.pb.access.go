// Code generated by protoc-gen-goten-resource
// Resource: TpmAttestationCert
// DO NOT EDIT!!!

package tpm_attestation_cert

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
	project "github.com/cloudwan/edgelq-sdk/devices/resources/v1/project"
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

type TpmAttestationCertAccess interface {
	GetTpmAttestationCert(context.Context, *GetQuery, ...gotenresource.GetOption) (*TpmAttestationCert, error)
	BatchGetTpmAttestationCerts(context.Context, []*Reference, ...gotenresource.BatchGetOption) error
	QueryTpmAttestationCerts(context.Context, *ListQuery, ...gotenresource.QueryOption) (*QueryResultSnapshot, error)
	WatchTpmAttestationCert(context.Context, *GetQuery, func(*TpmAttestationCertChange) error) error
	WatchTpmAttestationCerts(context.Context, *WatchQuery, func(*QueryResultChange) error) error
	SaveTpmAttestationCert(context.Context, *TpmAttestationCert, ...gotenresource.SaveOption) error
	DeleteTpmAttestationCert(context.Context, *Reference, ...gotenresource.DeleteOption) error
}

type anyCastAccess struct {
	TpmAttestationCertAccess
}

func AsAnyCastAccess(access TpmAttestationCertAccess) gotenresource.Access {
	return &anyCastAccess{TpmAttestationCertAccess: access}
}

func (a *anyCastAccess) Get(ctx context.Context, q gotenresource.GetQuery, opts ...gotenresource.GetOption) (gotenresource.Resource, error) {
	if asTpmAttestationCertQuery, ok := q.(*GetQuery); ok {
		return a.GetTpmAttestationCert(ctx, asTpmAttestationCertQuery, opts...)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected TpmAttestationCert, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Query(ctx context.Context, q gotenresource.ListQuery, opts ...gotenresource.QueryOption) (gotenresource.QueryResultSnapshot, error) {
	if asTpmAttestationCertQuery, ok := q.(*ListQuery); ok {
		return a.QueryTpmAttestationCerts(ctx, asTpmAttestationCertQuery, opts...)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected TpmAttestationCert, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Search(ctx context.Context, q gotenresource.SearchQuery, opts ...gotenresource.QueryOption) (gotenresource.QueryResultSnapshot, error) {
	return nil, status.Errorf(codes.Internal, "Search is not available for TpmAttestationCert")
}

func (a *anyCastAccess) Watch(ctx context.Context, q gotenresource.GetQuery, cb func(ch gotenresource.ResourceChange) error) error {
	if asTpmAttestationCertQuery, ok := q.(*GetQuery); ok {
		return a.WatchTpmAttestationCert(ctx, asTpmAttestationCertQuery, func(change *TpmAttestationCertChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected TpmAttestationCert, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) WatchQuery(ctx context.Context, q gotenresource.WatchQuery, cb func(ch gotenresource.QueryResultChange) error) error {
	if asTpmAttestationCertQuery, ok := q.(*WatchQuery); ok {
		return a.WatchTpmAttestationCerts(ctx, asTpmAttestationCertQuery, func(change *QueryResultChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected TpmAttestationCert, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Save(ctx context.Context, res gotenresource.Resource, opts ...gotenresource.SaveOption) error {
	if asTpmAttestationCertRes, ok := res.(*TpmAttestationCert); ok {
		return a.SaveTpmAttestationCert(ctx, asTpmAttestationCertRes, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected TpmAttestationCert, got: %s",
		res.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Delete(ctx context.Context, ref gotenresource.Reference, opts ...gotenresource.DeleteOption) error {
	if asTpmAttestationCertRef, ok := ref.(*Reference); ok {
		return a.DeleteTpmAttestationCert(ctx, asTpmAttestationCertRef, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected TpmAttestationCert, got: %s",
		ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) BatchGet(ctx context.Context, toGet []gotenresource.Reference, opts ...gotenresource.BatchGetOption) error {
	tpmAttestationCertRefs := make([]*Reference, 0, len(toGet))
	for _, ref := range toGet {
		if asTpmAttestationCertRef, ok := ref.(*Reference); !ok {
			return status.Errorf(codes.Internal,
				"Unrecognized descriptor, expected TpmAttestationCert, got: %s",
				ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
		} else {
			tpmAttestationCertRefs = append(tpmAttestationCertRefs, asTpmAttestationCertRef)
		}
	}
	return a.BatchGetTpmAttestationCerts(ctx, tpmAttestationCertRefs, opts...)
}

func (a *anyCastAccess) GetResourceDescriptors() []gotenresource.Descriptor {
	return []gotenresource.Descriptor{
		GetDescriptor(),
	}
}
