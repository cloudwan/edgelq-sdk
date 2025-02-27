// Code generated by protoc-gen-goten-resource
// Resource: Hardware
// DO NOT EDIT!!!

package hardware

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
	edgelq_instance "github.com/cloudwan/edgelq-sdk/ztp/resources/v1/edgelq_instance"
	project "github.com/cloudwan/edgelq-sdk/ztp/resources/v1/project"
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
	_ = &edgelq_instance.EdgelqInstance{}
	_ = &project.Project{}
	_ = &meta.Meta{}
)

type HardwareAccess interface {
	GetHardware(context.Context, *GetQuery, ...gotenresource.GetOption) (*Hardware, error)
	BatchGetHardwares(context.Context, []*Reference, ...gotenresource.BatchGetOption) error
	QueryHardwares(context.Context, *ListQuery, ...gotenresource.QueryOption) (*QueryResultSnapshot, error)
	WatchHardware(context.Context, *GetQuery, func(*HardwareChange) error) error
	WatchHardwares(context.Context, *WatchQuery, func(*QueryResultChange) error) error
	SaveHardware(context.Context, *Hardware, ...gotenresource.SaveOption) error
	DeleteHardware(context.Context, *Reference, ...gotenresource.DeleteOption) error
}

type anyCastAccess struct {
	HardwareAccess
}

func AsAnyCastAccess(access HardwareAccess) gotenresource.Access {
	return &anyCastAccess{HardwareAccess: access}
}

func (a *anyCastAccess) Get(ctx context.Context, q gotenresource.GetQuery, opts ...gotenresource.GetOption) (gotenresource.Resource, error) {
	if asHardwareQuery, ok := q.(*GetQuery); ok {
		return a.GetHardware(ctx, asHardwareQuery, opts...)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Hardware, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Query(ctx context.Context, q gotenresource.ListQuery, opts ...gotenresource.QueryOption) (gotenresource.QueryResultSnapshot, error) {
	if asHardwareQuery, ok := q.(*ListQuery); ok {
		return a.QueryHardwares(ctx, asHardwareQuery, opts...)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Hardware, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Search(ctx context.Context, q gotenresource.SearchQuery, opts ...gotenresource.QueryOption) (gotenresource.QueryResultSnapshot, error) {
	return nil, status.Errorf(codes.Internal, "Search is not available for Hardware")
}

func (a *anyCastAccess) Watch(ctx context.Context, q gotenresource.GetQuery, cb func(ch gotenresource.ResourceChange) error) error {
	if asHardwareQuery, ok := q.(*GetQuery); ok {
		return a.WatchHardware(ctx, asHardwareQuery, func(change *HardwareChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Hardware, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) WatchQuery(ctx context.Context, q gotenresource.WatchQuery, cb func(ch gotenresource.QueryResultChange) error) error {
	if asHardwareQuery, ok := q.(*WatchQuery); ok {
		return a.WatchHardwares(ctx, asHardwareQuery, func(change *QueryResultChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Hardware, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Save(ctx context.Context, res gotenresource.Resource, opts ...gotenresource.SaveOption) error {
	if asHardwareRes, ok := res.(*Hardware); ok {
		return a.SaveHardware(ctx, asHardwareRes, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Hardware, got: %s",
		res.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Delete(ctx context.Context, ref gotenresource.Reference, opts ...gotenresource.DeleteOption) error {
	if asHardwareRef, ok := ref.(*Reference); ok {
		return a.DeleteHardware(ctx, asHardwareRef, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Hardware, got: %s",
		ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) BatchGet(ctx context.Context, toGet []gotenresource.Reference, opts ...gotenresource.BatchGetOption) error {
	hardwareRefs := make([]*Reference, 0, len(toGet))
	for _, ref := range toGet {
		if asHardwareRef, ok := ref.(*Reference); !ok {
			return status.Errorf(codes.Internal,
				"Unrecognized descriptor, expected Hardware, got: %s",
				ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
		} else {
			hardwareRefs = append(hardwareRefs, asHardwareRef)
		}
	}
	return a.BatchGetHardwares(ctx, hardwareRefs, opts...)
}

func (a *anyCastAccess) GetResourceDescriptors() []gotenresource.Descriptor {
	return []gotenresource.Descriptor{
		GetDescriptor(),
	}
}
