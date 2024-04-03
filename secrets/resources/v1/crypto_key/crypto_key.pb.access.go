// Code generated by protoc-gen-goten-resource
// Resource: CryptoKey
// DO NOT EDIT!!!

package crypto_key

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
	project "github.com/cloudwan/edgelq-sdk/secrets/resources/v1/project"
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

type CryptoKeyAccess interface {
	GetCryptoKey(context.Context, *GetQuery) (*CryptoKey, error)
	BatchGetCryptoKeys(context.Context, []*Reference, ...gotenresource.BatchGetOption) error
	QueryCryptoKeys(context.Context, *ListQuery) (*QueryResultSnapshot, error)
	WatchCryptoKey(context.Context, *GetQuery, func(*CryptoKeyChange) error) error
	WatchCryptoKeys(context.Context, *WatchQuery, func(*QueryResultChange) error) error
	SaveCryptoKey(context.Context, *CryptoKey, ...gotenresource.SaveOption) error
	DeleteCryptoKey(context.Context, *Reference, ...gotenresource.DeleteOption) error
}

type anyCastAccess struct {
	CryptoKeyAccess
}

func AsAnyCastAccess(access CryptoKeyAccess) gotenresource.Access {
	return &anyCastAccess{CryptoKeyAccess: access}
}

func (a *anyCastAccess) Get(ctx context.Context, q gotenresource.GetQuery) (gotenresource.Resource, error) {
	if asCryptoKeyQuery, ok := q.(*GetQuery); ok {
		return a.GetCryptoKey(ctx, asCryptoKeyQuery)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected CryptoKey, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Query(ctx context.Context, q gotenresource.ListQuery) (gotenresource.QueryResultSnapshot, error) {
	if asCryptoKeyQuery, ok := q.(*ListQuery); ok {
		return a.QueryCryptoKeys(ctx, asCryptoKeyQuery)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected CryptoKey, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Search(ctx context.Context, q gotenresource.SearchQuery) (gotenresource.QueryResultSnapshot, error) {
	return nil, status.Errorf(codes.Internal, "Search is not available for CryptoKey")
}

func (a *anyCastAccess) Watch(ctx context.Context, q gotenresource.GetQuery, cb func(ch gotenresource.ResourceChange) error) error {
	if asCryptoKeyQuery, ok := q.(*GetQuery); ok {
		return a.WatchCryptoKey(ctx, asCryptoKeyQuery, func(change *CryptoKeyChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected CryptoKey, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) WatchQuery(ctx context.Context, q gotenresource.WatchQuery, cb func(ch gotenresource.QueryResultChange) error) error {
	if asCryptoKeyQuery, ok := q.(*WatchQuery); ok {
		return a.WatchCryptoKeys(ctx, asCryptoKeyQuery, func(change *QueryResultChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected CryptoKey, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Save(ctx context.Context, res gotenresource.Resource, opts ...gotenresource.SaveOption) error {
	if asCryptoKeyRes, ok := res.(*CryptoKey); ok {
		return a.SaveCryptoKey(ctx, asCryptoKeyRes, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected CryptoKey, got: %s",
		res.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Delete(ctx context.Context, ref gotenresource.Reference, opts ...gotenresource.DeleteOption) error {
	if asCryptoKeyRef, ok := ref.(*Reference); ok {
		return a.DeleteCryptoKey(ctx, asCryptoKeyRef, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected CryptoKey, got: %s",
		ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) BatchGet(ctx context.Context, toGet []gotenresource.Reference, opts ...gotenresource.BatchGetOption) error {
	cryptoKeyRefs := make([]*Reference, 0, len(toGet))
	for _, ref := range toGet {
		if asCryptoKeyRef, ok := ref.(*Reference); !ok {
			return status.Errorf(codes.Internal,
				"Unrecognized descriptor, expected CryptoKey, got: %s",
				ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
		} else {
			cryptoKeyRefs = append(cryptoKeyRefs, asCryptoKeyRef)
		}
	}
	return a.BatchGetCryptoKeys(ctx, cryptoKeyRefs, opts...)
}

func (a *anyCastAccess) GetResourceDescriptors() []gotenresource.Descriptor {
	return []gotenresource.Descriptor{
		GetDescriptor(),
	}
}