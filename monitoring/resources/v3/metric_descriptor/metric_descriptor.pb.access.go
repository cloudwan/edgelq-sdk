// Code generated by protoc-gen-goten-resource
// Resource: MetricDescriptor
// DO NOT EDIT!!!

package metric_descriptor

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
	common "github.com/cloudwan/edgelq-sdk/monitoring/resources/v3/common"
	monitored_resource_descriptor "github.com/cloudwan/edgelq-sdk/monitoring/resources/v3/monitored_resource_descriptor"
	project "github.com/cloudwan/edgelq-sdk/monitoring/resources/v3/project"
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
	_ = api.LaunchStage(0)
	_ = &common.LabelDescriptor{}
	_ = &monitored_resource_descriptor.MonitoredResourceDescriptor{}
	_ = &project.Project{}
	_ = &meta.Meta{}
)

type MetricDescriptorAccess interface {
	GetMetricDescriptor(context.Context, *GetQuery, ...gotenresource.GetOption) (*MetricDescriptor, error)
	BatchGetMetricDescriptors(context.Context, []*Reference, ...gotenresource.BatchGetOption) error
	QueryMetricDescriptors(context.Context, *ListQuery, ...gotenresource.QueryOption) (*QueryResultSnapshot, error)
	WatchMetricDescriptor(context.Context, *GetQuery, func(*MetricDescriptorChange) error) error
	WatchMetricDescriptors(context.Context, *WatchQuery, func(*QueryResultChange) error) error
	SaveMetricDescriptor(context.Context, *MetricDescriptor, ...gotenresource.SaveOption) error
	DeleteMetricDescriptor(context.Context, *Reference, ...gotenresource.DeleteOption) error
}

type anyCastAccess struct {
	MetricDescriptorAccess
}

func AsAnyCastAccess(access MetricDescriptorAccess) gotenresource.Access {
	return &anyCastAccess{MetricDescriptorAccess: access}
}

func (a *anyCastAccess) Get(ctx context.Context, q gotenresource.GetQuery, opts ...gotenresource.GetOption) (gotenresource.Resource, error) {
	if asMetricDescriptorQuery, ok := q.(*GetQuery); ok {
		return a.GetMetricDescriptor(ctx, asMetricDescriptorQuery, opts...)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected MetricDescriptor, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Query(ctx context.Context, q gotenresource.ListQuery, opts ...gotenresource.QueryOption) (gotenresource.QueryResultSnapshot, error) {
	if asMetricDescriptorQuery, ok := q.(*ListQuery); ok {
		return a.QueryMetricDescriptors(ctx, asMetricDescriptorQuery, opts...)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected MetricDescriptor, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Search(ctx context.Context, q gotenresource.SearchQuery, opts ...gotenresource.QueryOption) (gotenresource.QueryResultSnapshot, error) {
	return nil, status.Errorf(codes.Internal, "Search is not available for MetricDescriptor")
}

func (a *anyCastAccess) Watch(ctx context.Context, q gotenresource.GetQuery, cb func(ch gotenresource.ResourceChange) error) error {
	if asMetricDescriptorQuery, ok := q.(*GetQuery); ok {
		return a.WatchMetricDescriptor(ctx, asMetricDescriptorQuery, func(change *MetricDescriptorChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected MetricDescriptor, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) WatchQuery(ctx context.Context, q gotenresource.WatchQuery, cb func(ch gotenresource.QueryResultChange) error) error {
	if asMetricDescriptorQuery, ok := q.(*WatchQuery); ok {
		return a.WatchMetricDescriptors(ctx, asMetricDescriptorQuery, func(change *QueryResultChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected MetricDescriptor, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Save(ctx context.Context, res gotenresource.Resource, opts ...gotenresource.SaveOption) error {
	if asMetricDescriptorRes, ok := res.(*MetricDescriptor); ok {
		return a.SaveMetricDescriptor(ctx, asMetricDescriptorRes, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected MetricDescriptor, got: %s",
		res.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Delete(ctx context.Context, ref gotenresource.Reference, opts ...gotenresource.DeleteOption) error {
	if asMetricDescriptorRef, ok := ref.(*Reference); ok {
		return a.DeleteMetricDescriptor(ctx, asMetricDescriptorRef, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected MetricDescriptor, got: %s",
		ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) BatchGet(ctx context.Context, toGet []gotenresource.Reference, opts ...gotenresource.BatchGetOption) error {
	metricDescriptorRefs := make([]*Reference, 0, len(toGet))
	for _, ref := range toGet {
		if asMetricDescriptorRef, ok := ref.(*Reference); !ok {
			return status.Errorf(codes.Internal,
				"Unrecognized descriptor, expected MetricDescriptor, got: %s",
				ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
		} else {
			metricDescriptorRefs = append(metricDescriptorRefs, asMetricDescriptorRef)
		}
	}
	return a.BatchGetMetricDescriptors(ctx, metricDescriptorRefs, opts...)
}

func (a *anyCastAccess) GetResourceDescriptors() []gotenresource.Descriptor {
	return []gotenresource.Descriptor{
		GetDescriptor(),
	}
}
