// Code generated by protoc-gen-goten-resource
// Resource: DeviceHardware
// DO NOT EDIT!!!

package device_hardware

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
	device "github.com/cloudwan/edgelq-sdk/devices/resources/v1/device"
	project "github.com/cloudwan/edgelq-sdk/devices/resources/v1/project"
	provisioning_policy "github.com/cloudwan/edgelq-sdk/devices/resources/v1/provisioning_policy"
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
	_ = &device.Device{}
	_ = &project.Project{}
	_ = &provisioning_policy.ProvisioningPolicy{}
	_ = &meta.Meta{}
)

type DeviceHardwareAccess interface {
	GetDeviceHardware(context.Context, *GetQuery) (*DeviceHardware, error)
	BatchGetDeviceHardwares(context.Context, []*Reference, ...gotenresource.BatchGetOption) error
	QueryDeviceHardwares(context.Context, *ListQuery) (*QueryResultSnapshot, error)
	WatchDeviceHardware(context.Context, *GetQuery, func(*DeviceHardwareChange) error) error
	WatchDeviceHardwares(context.Context, *WatchQuery, func(*QueryResultChange) error) error
	SaveDeviceHardware(context.Context, *DeviceHardware, ...gotenresource.SaveOption) error
	DeleteDeviceHardware(context.Context, *Reference, ...gotenresource.DeleteOption) error
}

type anyCastAccess struct {
	DeviceHardwareAccess
}

func AsAnyCastAccess(access DeviceHardwareAccess) gotenresource.Access {
	return &anyCastAccess{DeviceHardwareAccess: access}
}

func (a *anyCastAccess) Get(ctx context.Context, q gotenresource.GetQuery) (gotenresource.Resource, error) {
	if asDeviceHardwareQuery, ok := q.(*GetQuery); ok {
		return a.GetDeviceHardware(ctx, asDeviceHardwareQuery)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected DeviceHardware, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Query(ctx context.Context, q gotenresource.ListQuery) (gotenresource.QueryResultSnapshot, error) {
	if asDeviceHardwareQuery, ok := q.(*ListQuery); ok {
		return a.QueryDeviceHardwares(ctx, asDeviceHardwareQuery)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected DeviceHardware, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Search(ctx context.Context, q gotenresource.SearchQuery) (gotenresource.QueryResultSnapshot, error) {
	return nil, status.Errorf(codes.Internal, "Search is not available for DeviceHardware")
}

func (a *anyCastAccess) Watch(ctx context.Context, q gotenresource.GetQuery, cb func(ch gotenresource.ResourceChange) error) error {
	if asDeviceHardwareQuery, ok := q.(*GetQuery); ok {
		return a.WatchDeviceHardware(ctx, asDeviceHardwareQuery, func(change *DeviceHardwareChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected DeviceHardware, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) WatchQuery(ctx context.Context, q gotenresource.WatchQuery, cb func(ch gotenresource.QueryResultChange) error) error {
	if asDeviceHardwareQuery, ok := q.(*WatchQuery); ok {
		return a.WatchDeviceHardwares(ctx, asDeviceHardwareQuery, func(change *QueryResultChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected DeviceHardware, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Save(ctx context.Context, res gotenresource.Resource, opts ...gotenresource.SaveOption) error {
	if asDeviceHardwareRes, ok := res.(*DeviceHardware); ok {
		return a.SaveDeviceHardware(ctx, asDeviceHardwareRes, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected DeviceHardware, got: %s",
		res.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Delete(ctx context.Context, ref gotenresource.Reference, opts ...gotenresource.DeleteOption) error {
	if asDeviceHardwareRef, ok := ref.(*Reference); ok {
		return a.DeleteDeviceHardware(ctx, asDeviceHardwareRef, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected DeviceHardware, got: %s",
		ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) BatchGet(ctx context.Context, toGet []gotenresource.Reference, opts ...gotenresource.BatchGetOption) error {
	deviceHardwareRefs := make([]*Reference, 0, len(toGet))
	for _, ref := range toGet {
		if asDeviceHardwareRef, ok := ref.(*Reference); !ok {
			return status.Errorf(codes.Internal,
				"Unrecognized descriptor, expected DeviceHardware, got: %s",
				ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
		} else {
			deviceHardwareRefs = append(deviceHardwareRefs, asDeviceHardwareRef)
		}
	}
	return a.BatchGetDeviceHardwares(ctx, deviceHardwareRefs, opts...)
}

func (a *anyCastAccess) GetResourceDescriptors() []gotenresource.Descriptor {
	return []gotenresource.Descriptor{
		GetDescriptor(),
	}
}
