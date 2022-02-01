// Code generated by protoc-gen-goten-resource
// Resource: Device
// DO NOT EDIT!!!

package device

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
	project "github.com/cloudwan/edgelq-sdk/devices/resources/v1alpha/project"
	iam_service_account "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/service_account"
	duration "github.com/golang/protobuf/ptypes/duration"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
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
	_ = &project.Project{}
	_ = &iam_service_account.ServiceAccount{}
	_ = &duration.Duration{}
	_ = &field_mask.FieldMask{}
	_ = &timestamp.Timestamp{}
)

type DeviceAccess interface {
	GetDevice(context.Context, *GetQuery) (*Device, error)
	BatchGetDevices(context.Context, []*Reference, ...gotenresource.BatchGetOption) error
	QueryDevices(context.Context, *ListQuery) (*QueryResultSnapshot, error)
	WatchDevice(context.Context, *GetQuery, func(*DeviceChange) error) error
	WatchDevices(context.Context, *WatchQuery, func(*QueryResultChange) error) error
	SaveDevice(context.Context, *Device, ...gotenresource.SaveOption) error
	DeleteDevice(context.Context, *Reference, ...gotenresource.DeleteOption) error
}

type anyCastAccess struct {
	DeviceAccess
}

func AsAnyCastAccess(access DeviceAccess) gotenresource.Access {
	return &anyCastAccess{DeviceAccess: access}
}

func (a *anyCastAccess) Get(ctx context.Context, q gotenresource.GetQuery) (gotenresource.Resource, error) {
	if asDeviceQuery, ok := q.(*GetQuery); ok {
		return a.GetDevice(ctx, asDeviceQuery)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Device, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Query(ctx context.Context, q gotenresource.ListQuery) (gotenresource.QueryResultSnapshot, error) {
	if asDeviceQuery, ok := q.(*ListQuery); ok {
		return a.QueryDevices(ctx, asDeviceQuery)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Device, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Search(ctx context.Context, q gotenresource.SearchQuery) (gotenresource.SearchQueryResultSnapshot, error) {
	return nil, status.Errorf(codes.Internal, "Search is not available for Device")
}

func (a *anyCastAccess) Watch(ctx context.Context, q gotenresource.GetQuery, cb func(ch gotenresource.ResourceChange) error) error {
	if asDeviceQuery, ok := q.(*GetQuery); ok {
		return a.WatchDevice(ctx, asDeviceQuery, func(change *DeviceChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Device, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) WatchQuery(ctx context.Context, q gotenresource.WatchQuery, cb func(ch gotenresource.QueryResultChange) error) error {
	if asDeviceQuery, ok := q.(*WatchQuery); ok {
		return a.WatchDevices(ctx, asDeviceQuery, func(change *QueryResultChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Device, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Save(ctx context.Context, res gotenresource.Resource, opts ...gotenresource.SaveOption) error {
	if asDeviceRes, ok := res.(*Device); ok {
		return a.SaveDevice(ctx, asDeviceRes, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Device, got: %s",
		res.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Delete(ctx context.Context, ref gotenresource.Reference, opts ...gotenresource.DeleteOption) error {
	if asDeviceRef, ok := ref.(*Reference); ok {
		return a.DeleteDevice(ctx, asDeviceRef, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Device, got: %s",
		ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) BatchGet(ctx context.Context, toGet []gotenresource.Reference, opts ...gotenresource.BatchGetOption) error {
	deviceRefs := make([]*Reference, 0, len(toGet))
	for _, ref := range toGet {
		if asDeviceRef, ok := ref.(*Reference); !ok {
			return status.Errorf(codes.Internal,
				"Unrecognized descriptor, expected Device, got: %s",
				ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
		} else {
			deviceRefs = append(deviceRefs, asDeviceRef)
		}
	}
	return a.BatchGetDevices(ctx, deviceRefs, opts...)
}

func (a *anyCastAccess) GetResourceDescriptors() []gotenresource.Descriptor {
	return []gotenresource.Descriptor{
		GetDescriptor(),
	}
}
