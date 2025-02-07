// Code generated by protoc-gen-goten-resource
// Resource: Device
// DO NOT EDIT!!!

package device

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
	api "github.com/cloudwan/edgelq-sdk/common/api"
	project "github.com/cloudwan/edgelq-sdk/devices/resources/v1/project"
	iam_attestation_domain "github.com/cloudwan/edgelq-sdk/iam/resources/v1/attestation_domain"
	iam_iam_common "github.com/cloudwan/edgelq-sdk/iam/resources/v1/common"
	iam_service_account "github.com/cloudwan/edgelq-sdk/iam/resources/v1/service_account"
	logging_bucket "github.com/cloudwan/edgelq-sdk/logging/resources/v1/bucket"
	monitoring_bucket "github.com/cloudwan/edgelq-sdk/monitoring/resources/v4/bucket"
	meta "github.com/cloudwan/goten-sdk/types/meta"
	latlng "google.golang.org/genproto/googleapis/type/latlng"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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
	_ = &api.HealthCheckSpec{}
	_ = &project.Project{}
	_ = &iam_attestation_domain.AttestationDomain{}
	_ = &iam_iam_common.PCR{}
	_ = &iam_service_account.ServiceAccount{}
	_ = &logging_bucket.Bucket{}
	_ = &monitoring_bucket.Bucket{}
	_ = &durationpb.Duration{}
	_ = &fieldmaskpb.FieldMask{}
	_ = &timestamppb.Timestamp{}
	_ = &latlng.LatLng{}
	_ = &meta.Meta{}
)

type DeviceAccess interface {
	GetDevice(context.Context, *GetQuery, ...gotenresource.GetOption) (*Device, error)
	BatchGetDevices(context.Context, []*Reference, ...gotenresource.BatchGetOption) error
	QueryDevices(context.Context, *ListQuery, ...gotenresource.QueryOption) (*QueryResultSnapshot, error)
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

func (a *anyCastAccess) Get(ctx context.Context, q gotenresource.GetQuery, opts ...gotenresource.GetOption) (gotenresource.Resource, error) {
	if asDeviceQuery, ok := q.(*GetQuery); ok {
		return a.GetDevice(ctx, asDeviceQuery, opts...)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Device, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Query(ctx context.Context, q gotenresource.ListQuery, opts ...gotenresource.QueryOption) (gotenresource.QueryResultSnapshot, error) {
	if asDeviceQuery, ok := q.(*ListQuery); ok {
		return a.QueryDevices(ctx, asDeviceQuery, opts...)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Device, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Search(ctx context.Context, q gotenresource.SearchQuery, opts ...gotenresource.QueryOption) (gotenresource.QueryResultSnapshot, error) {
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
