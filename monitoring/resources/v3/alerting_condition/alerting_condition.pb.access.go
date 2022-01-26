// Code generated by protoc-gen-goten-resource
// Resource: AlertingCondition
// DO NOT EDIT!!!

package alerting_condition

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
	alerting_policy "github.com/cloudwan/edgelq-sdk/monitoring/resources/v3/alerting_policy"
	duration "github.com/golang/protobuf/ptypes/duration"
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
	_ = &alerting_policy.AlertingPolicy{}
	_ = &monitoring_common.LabelDescriptor{}
	_ = &duration.Duration{}
)

type AlertingConditionAccess interface {
	GetAlertingCondition(context.Context, *GetQuery) (*AlertingCondition, error)
	BatchGetAlertingConditions(context.Context, []*Reference, ...gotenresource.BatchGetOption) error
	QueryAlertingConditions(context.Context, *ListQuery) (*QueryResultSnapshot, error)
	WatchAlertingCondition(context.Context, *GetQuery, func(*AlertingConditionChange) error) error
	WatchAlertingConditions(context.Context, *WatchQuery, func(*QueryResultChange) error) error
	SaveAlertingCondition(context.Context, *AlertingCondition, ...gotenresource.SaveOption) error
	DeleteAlertingCondition(context.Context, *Reference, ...gotenresource.DeleteOption) error
}

type anyCastAccess struct {
	AlertingConditionAccess
}

func AsAnyCastAccess(access AlertingConditionAccess) gotenresource.Access {
	return &anyCastAccess{AlertingConditionAccess: access}
}

func (a *anyCastAccess) Get(ctx context.Context, q gotenresource.GetQuery) (gotenresource.Resource, error) {
	if asAlertingConditionQuery, ok := q.(*GetQuery); ok {
		return a.GetAlertingCondition(ctx, asAlertingConditionQuery)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected AlertingCondition, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Query(ctx context.Context, q gotenresource.ListQuery) (gotenresource.QueryResultSnapshot, error) {
	if asAlertingConditionQuery, ok := q.(*ListQuery); ok {
		return a.QueryAlertingConditions(ctx, asAlertingConditionQuery)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected AlertingCondition, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Search(ctx context.Context, q gotenresource.SearchQuery) (gotenresource.SearchQueryResultSnapshot, error) {
	return nil, status.Errorf(codes.Internal, "Search is not available for AlertingCondition")
}

func (a *anyCastAccess) Watch(ctx context.Context, q gotenresource.GetQuery, cb func(ch gotenresource.ResourceChange) error) error {
	if asAlertingConditionQuery, ok := q.(*GetQuery); ok {
		return a.WatchAlertingCondition(ctx, asAlertingConditionQuery, func(change *AlertingConditionChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected AlertingCondition, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) WatchQuery(ctx context.Context, q gotenresource.WatchQuery, cb func(ch gotenresource.QueryResultChange) error) error {
	if asAlertingConditionQuery, ok := q.(*WatchQuery); ok {
		return a.WatchAlertingConditions(ctx, asAlertingConditionQuery, func(change *QueryResultChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected AlertingCondition, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Save(ctx context.Context, res gotenresource.Resource, opts ...gotenresource.SaveOption) error {
	if asAlertingConditionRes, ok := res.(*AlertingCondition); ok {
		return a.SaveAlertingCondition(ctx, asAlertingConditionRes, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected AlertingCondition, got: %s",
		res.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Delete(ctx context.Context, ref gotenresource.Reference, opts ...gotenresource.DeleteOption) error {
	if asAlertingConditionRef, ok := ref.(*Reference); ok {
		return a.DeleteAlertingCondition(ctx, asAlertingConditionRef, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected AlertingCondition, got: %s",
		ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) BatchGet(ctx context.Context, toGet []gotenresource.Reference, opts ...gotenresource.BatchGetOption) error {
	alertingConditionRefs := make([]*Reference, 0, len(toGet))
	for _, ref := range toGet {
		if asAlertingConditionRef, ok := ref.(*Reference); !ok {
			return status.Errorf(codes.Internal,
				"Unrecognized descriptor, expected AlertingCondition, got: %s",
				ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
		} else {
			alertingConditionRefs = append(alertingConditionRefs, asAlertingConditionRef)
		}
	}
	return a.BatchGetAlertingConditions(ctx, alertingConditionRefs, opts...)
}

func (a *anyCastAccess) GetResourceDescriptors() []gotenresource.Descriptor {
	return []gotenresource.Descriptor{
		GetDescriptor(),
	}
}
