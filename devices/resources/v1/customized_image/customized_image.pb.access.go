// Code generated by protoc-gen-goten-resource
// Resource: CustomizedImage
// DO NOT EDIT!!!

package customized_image

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
	os_version "github.com/cloudwan/edgelq-sdk/devices/resources/v1/os_version"
	project "github.com/cloudwan/edgelq-sdk/devices/resources/v1/project"
	iam_service_account "github.com/cloudwan/edgelq-sdk/iam/resources/v1/service_account"
	iam_service_account_key "github.com/cloudwan/edgelq-sdk/iam/resources/v1/service_account_key"
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
	_ = &os_version.OsVersion{}
	_ = &project.Project{}
	_ = &iam_service_account.ServiceAccount{}
	_ = &iam_service_account_key.ServiceAccountKey{}
	_ = &meta.Meta{}
)

type CustomizedImageAccess interface {
	GetCustomizedImage(context.Context, *GetQuery, ...gotenresource.GetOption) (*CustomizedImage, error)
	BatchGetCustomizedImages(context.Context, []*Reference, ...gotenresource.BatchGetOption) error
	QueryCustomizedImages(context.Context, *ListQuery, ...gotenresource.QueryOption) (*QueryResultSnapshot, error)
	WatchCustomizedImage(context.Context, *GetQuery, func(*CustomizedImageChange) error) error
	WatchCustomizedImages(context.Context, *WatchQuery, func(*QueryResultChange) error) error
	SaveCustomizedImage(context.Context, *CustomizedImage, ...gotenresource.SaveOption) error
	DeleteCustomizedImage(context.Context, *Reference, ...gotenresource.DeleteOption) error
}

type anyCastAccess struct {
	CustomizedImageAccess
}

func AsAnyCastAccess(access CustomizedImageAccess) gotenresource.Access {
	return &anyCastAccess{CustomizedImageAccess: access}
}

func (a *anyCastAccess) Get(ctx context.Context, q gotenresource.GetQuery, opts ...gotenresource.GetOption) (gotenresource.Resource, error) {
	if asCustomizedImageQuery, ok := q.(*GetQuery); ok {
		return a.GetCustomizedImage(ctx, asCustomizedImageQuery, opts...)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected CustomizedImage, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Query(ctx context.Context, q gotenresource.ListQuery, opts ...gotenresource.QueryOption) (gotenresource.QueryResultSnapshot, error) {
	if asCustomizedImageQuery, ok := q.(*ListQuery); ok {
		return a.QueryCustomizedImages(ctx, asCustomizedImageQuery, opts...)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected CustomizedImage, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Search(ctx context.Context, q gotenresource.SearchQuery, opts ...gotenresource.QueryOption) (gotenresource.QueryResultSnapshot, error) {
	return nil, status.Errorf(codes.Internal, "Search is not available for CustomizedImage")
}

func (a *anyCastAccess) Watch(ctx context.Context, q gotenresource.GetQuery, cb func(ch gotenresource.ResourceChange) error) error {
	if asCustomizedImageQuery, ok := q.(*GetQuery); ok {
		return a.WatchCustomizedImage(ctx, asCustomizedImageQuery, func(change *CustomizedImageChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected CustomizedImage, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) WatchQuery(ctx context.Context, q gotenresource.WatchQuery, cb func(ch gotenresource.QueryResultChange) error) error {
	if asCustomizedImageQuery, ok := q.(*WatchQuery); ok {
		return a.WatchCustomizedImages(ctx, asCustomizedImageQuery, func(change *QueryResultChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected CustomizedImage, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Save(ctx context.Context, res gotenresource.Resource, opts ...gotenresource.SaveOption) error {
	if asCustomizedImageRes, ok := res.(*CustomizedImage); ok {
		return a.SaveCustomizedImage(ctx, asCustomizedImageRes, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected CustomizedImage, got: %s",
		res.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Delete(ctx context.Context, ref gotenresource.Reference, opts ...gotenresource.DeleteOption) error {
	if asCustomizedImageRef, ok := ref.(*Reference); ok {
		return a.DeleteCustomizedImage(ctx, asCustomizedImageRef, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected CustomizedImage, got: %s",
		ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) BatchGet(ctx context.Context, toGet []gotenresource.Reference, opts ...gotenresource.BatchGetOption) error {
	customizedImageRefs := make([]*Reference, 0, len(toGet))
	for _, ref := range toGet {
		if asCustomizedImageRef, ok := ref.(*Reference); !ok {
			return status.Errorf(codes.Internal,
				"Unrecognized descriptor, expected CustomizedImage, got: %s",
				ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
		} else {
			customizedImageRefs = append(customizedImageRefs, asCustomizedImageRef)
		}
	}
	return a.BatchGetCustomizedImages(ctx, customizedImageRefs, opts...)
}

func (a *anyCastAccess) GetResourceDescriptors() []gotenresource.Descriptor {
	return []gotenresource.Descriptor{
		GetDescriptor(),
	}
}
