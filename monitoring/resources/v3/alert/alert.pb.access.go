// Code generated by protoc-gen-goten-resource
// Resource: Alert
// DO NOT EDIT!!!

package alert

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
	monitoring_common "github.com/cloudwan/edgelq-sdk/monitoring/common/v3"
	alerting_condition "github.com/cloudwan/edgelq-sdk/monitoring/resources/v3/alerting_condition"
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
	_ = &alerting_condition.AlertingCondition{}
	_ = &monitoring_common.LabelDescriptor{}
)

type AlertAccess interface {
	GetAlert(context.Context, *GetQuery) (*Alert, error)
	BatchGetAlerts(context.Context, []*Reference, ...gotenresource.BatchGetOption) error
	QueryAlerts(context.Context, *ListQuery) (*QueryResultSnapshot, error)
	WatchAlert(context.Context, *GetQuery, func(*AlertChange) error) error
	WatchAlerts(context.Context, *WatchQuery, func(*QueryResultChange) error) error
	SaveAlert(context.Context, *Alert, ...gotenresource.SaveOption) error
	DeleteAlert(context.Context, *Reference, ...gotenresource.DeleteOption) error
}

type anyCastAccess struct {
	AlertAccess
}

func AsAnyCastAccess(access AlertAccess) gotenresource.Access {
	return &anyCastAccess{AlertAccess: access}
}

func (a *anyCastAccess) Get(ctx context.Context, q gotenresource.GetQuery) (gotenresource.Resource, error) {
	if asAlertQuery, ok := q.(*GetQuery); ok {
		return a.GetAlert(ctx, asAlertQuery)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Alert, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Query(ctx context.Context, q gotenresource.ListQuery) (gotenresource.QueryResultSnapshot, error) {
	if asAlertQuery, ok := q.(*ListQuery); ok {
		return a.QueryAlerts(ctx, asAlertQuery)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Alert, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Search(ctx context.Context, q gotenresource.SearchQuery) (gotenresource.SearchQueryResultSnapshot, error) {
	return nil, status.Errorf(codes.Internal, "Search is not available for Alert")
}

func (a *anyCastAccess) Watch(ctx context.Context, q gotenresource.GetQuery, cb func(ch gotenresource.ResourceChange) error) error {
	if asAlertQuery, ok := q.(*GetQuery); ok {
		return a.WatchAlert(ctx, asAlertQuery, func(change *AlertChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Alert, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) WatchQuery(ctx context.Context, q gotenresource.WatchQuery, cb func(ch gotenresource.QueryResultChange) error) error {
	if asAlertQuery, ok := q.(*WatchQuery); ok {
		return a.WatchAlerts(ctx, asAlertQuery, func(change *QueryResultChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Alert, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Save(ctx context.Context, res gotenresource.Resource, opts ...gotenresource.SaveOption) error {
	if asAlertRes, ok := res.(*Alert); ok {
		return a.SaveAlert(ctx, asAlertRes, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Alert, got: %s",
		res.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Delete(ctx context.Context, ref gotenresource.Reference, opts ...gotenresource.DeleteOption) error {
	if asAlertRef, ok := ref.(*Reference); ok {
		return a.DeleteAlert(ctx, asAlertRef, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected Alert, got: %s",
		ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) BatchGet(ctx context.Context, toGet []gotenresource.Reference, opts ...gotenresource.BatchGetOption) error {
	alertRefs := make([]*Reference, 0, len(toGet))
	for _, ref := range toGet {
		if asAlertRef, ok := ref.(*Reference); !ok {
			return status.Errorf(codes.Internal,
				"Unrecognized descriptor, expected Alert, got: %s",
				ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
		} else {
			alertRefs = append(alertRefs, asAlertRef)
		}
	}
	return a.BatchGetAlerts(ctx, alertRefs, opts...)
}

func (a *anyCastAccess) GetResourceDescriptors() []gotenresource.Descriptor {
	return []gotenresource.Descriptor{
		GetDescriptor(),
	}
}
