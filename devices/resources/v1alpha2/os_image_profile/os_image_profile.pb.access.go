// Code generated by protoc-gen-goten-resource
// Resource: OsImageProfile
// DO NOT EDIT!!!

package os_image_profile

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
	device_type "github.com/cloudwan/edgelq-sdk/devices/resources/v1alpha2/device_type"
	project "github.com/cloudwan/edgelq-sdk/devices/resources/v1alpha2/project"
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
	_ = &device_type.DeviceType{}
	_ = &project.Project{}
	_ = &meta.Meta{}
)

type OsImageProfileAccess interface {
	GetOsImageProfile(context.Context, *GetQuery) (*OsImageProfile, error)
	BatchGetOsImageProfiles(context.Context, []*Reference, ...gotenresource.BatchGetOption) error
	QueryOsImageProfiles(context.Context, *ListQuery) (*QueryResultSnapshot, error)
	WatchOsImageProfile(context.Context, *GetQuery, func(*OsImageProfileChange) error) error
	WatchOsImageProfiles(context.Context, *WatchQuery, func(*QueryResultChange) error) error
	SaveOsImageProfile(context.Context, *OsImageProfile, ...gotenresource.SaveOption) error
	DeleteOsImageProfile(context.Context, *Reference, ...gotenresource.DeleteOption) error
}

type anyCastAccess struct {
	OsImageProfileAccess
}

func AsAnyCastAccess(access OsImageProfileAccess) gotenresource.Access {
	return &anyCastAccess{OsImageProfileAccess: access}
}

func (a *anyCastAccess) Get(ctx context.Context, q gotenresource.GetQuery) (gotenresource.Resource, error) {
	if asOsImageProfileQuery, ok := q.(*GetQuery); ok {
		return a.GetOsImageProfile(ctx, asOsImageProfileQuery)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected OsImageProfile, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Query(ctx context.Context, q gotenresource.ListQuery) (gotenresource.QueryResultSnapshot, error) {
	if asOsImageProfileQuery, ok := q.(*ListQuery); ok {
		return a.QueryOsImageProfiles(ctx, asOsImageProfileQuery)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected OsImageProfile, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Search(ctx context.Context, q gotenresource.SearchQuery) (gotenresource.QueryResultSnapshot, error) {
	return nil, status.Errorf(codes.Internal, "Search is not available for OsImageProfile")
}

func (a *anyCastAccess) Watch(ctx context.Context, q gotenresource.GetQuery, cb func(ch gotenresource.ResourceChange) error) error {
	if asOsImageProfileQuery, ok := q.(*GetQuery); ok {
		return a.WatchOsImageProfile(ctx, asOsImageProfileQuery, func(change *OsImageProfileChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected OsImageProfile, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) WatchQuery(ctx context.Context, q gotenresource.WatchQuery, cb func(ch gotenresource.QueryResultChange) error) error {
	if asOsImageProfileQuery, ok := q.(*WatchQuery); ok {
		return a.WatchOsImageProfiles(ctx, asOsImageProfileQuery, func(change *QueryResultChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected OsImageProfile, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Save(ctx context.Context, res gotenresource.Resource, opts ...gotenresource.SaveOption) error {
	if asOsImageProfileRes, ok := res.(*OsImageProfile); ok {
		return a.SaveOsImageProfile(ctx, asOsImageProfileRes, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected OsImageProfile, got: %s",
		res.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Delete(ctx context.Context, ref gotenresource.Reference, opts ...gotenresource.DeleteOption) error {
	if asOsImageProfileRef, ok := ref.(*Reference); ok {
		return a.DeleteOsImageProfile(ctx, asOsImageProfileRef, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected OsImageProfile, got: %s",
		ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) BatchGet(ctx context.Context, toGet []gotenresource.Reference, opts ...gotenresource.BatchGetOption) error {
	osImageProfileRefs := make([]*Reference, 0, len(toGet))
	for _, ref := range toGet {
		if asOsImageProfileRef, ok := ref.(*Reference); !ok {
			return status.Errorf(codes.Internal,
				"Unrecognized descriptor, expected OsImageProfile, got: %s",
				ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
		} else {
			osImageProfileRefs = append(osImageProfileRefs, asOsImageProfileRef)
		}
	}
	return a.BatchGetOsImageProfiles(ctx, osImageProfileRefs, opts...)
}

func (a *anyCastAccess) GetResourceDescriptors() []gotenresource.Descriptor {
	return []gotenresource.Descriptor{
		GetDescriptor(),
	}
}
