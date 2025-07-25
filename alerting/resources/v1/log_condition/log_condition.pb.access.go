// Code generated by protoc-gen-goten-resource
// Resource: LogCondition
// DO NOT EDIT!!!

package log_condition

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
	rcommon "github.com/cloudwan/edgelq-sdk/alerting/resources/v1/common"
	document "github.com/cloudwan/edgelq-sdk/alerting/resources/v1/document"
	log_condition_template "github.com/cloudwan/edgelq-sdk/alerting/resources/v1/log_condition_template"
	policy "github.com/cloudwan/edgelq-sdk/alerting/resources/v1/policy"
	meta "github.com/cloudwan/goten-sdk/types/meta"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
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
	_ = &document.Document{}
	_ = &log_condition_template.LogConditionTemplate{}
	_ = &policy.Policy{}
	_ = &rcommon.LogCndSpec{}
	_ = &fieldmaskpb.FieldMask{}
	_ = &meta.Meta{}
)

type LogConditionAccess interface {
	GetLogCondition(context.Context, *GetQuery, ...gotenresource.GetOption) (*LogCondition, error)
	BatchGetLogConditions(context.Context, []*Reference, ...gotenresource.BatchGetOption) error
	QueryLogConditions(context.Context, *ListQuery, ...gotenresource.QueryOption) (*QueryResultSnapshot, error)
	SearchLogConditions(context.Context, *SearchQuery, ...gotenresource.QueryOption) (*QueryResultSnapshot, error)
	WatchLogCondition(context.Context, *GetQuery, func(*LogConditionChange) error) error
	WatchLogConditions(context.Context, *WatchQuery, func(*QueryResultChange) error) error
	SaveLogCondition(context.Context, *LogCondition, ...gotenresource.SaveOption) error
	DeleteLogCondition(context.Context, *Reference, ...gotenresource.DeleteOption) error
}

type anyCastAccess struct {
	LogConditionAccess
}

func AsAnyCastAccess(access LogConditionAccess) gotenresource.Access {
	return &anyCastAccess{LogConditionAccess: access}
}

func (a *anyCastAccess) Get(ctx context.Context, q gotenresource.GetQuery, opts ...gotenresource.GetOption) (gotenresource.Resource, error) {
	if asLogConditionQuery, ok := q.(*GetQuery); ok {
		return a.GetLogCondition(ctx, asLogConditionQuery, opts...)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected LogCondition, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Query(ctx context.Context, q gotenresource.ListQuery, opts ...gotenresource.QueryOption) (gotenresource.QueryResultSnapshot, error) {
	if asLogConditionQuery, ok := q.(*ListQuery); ok {
		return a.QueryLogConditions(ctx, asLogConditionQuery, opts...)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected LogCondition, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Search(ctx context.Context, q gotenresource.SearchQuery, opts ...gotenresource.QueryOption) (gotenresource.QueryResultSnapshot, error) {
	if asLogConditionQuery, ok := q.(*SearchQuery); ok {
		return a.SearchLogConditions(ctx, asLogConditionQuery, opts...)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected LogCondition, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Watch(ctx context.Context, q gotenresource.GetQuery, cb func(ch gotenresource.ResourceChange) error) error {
	if asLogConditionQuery, ok := q.(*GetQuery); ok {
		return a.WatchLogCondition(ctx, asLogConditionQuery, func(change *LogConditionChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected LogCondition, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) WatchQuery(ctx context.Context, q gotenresource.WatchQuery, cb func(ch gotenresource.QueryResultChange) error) error {
	if asLogConditionQuery, ok := q.(*WatchQuery); ok {
		return a.WatchLogConditions(ctx, asLogConditionQuery, func(change *QueryResultChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected LogCondition, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Save(ctx context.Context, res gotenresource.Resource, opts ...gotenresource.SaveOption) error {
	if asLogConditionRes, ok := res.(*LogCondition); ok {
		return a.SaveLogCondition(ctx, asLogConditionRes, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected LogCondition, got: %s",
		res.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Delete(ctx context.Context, ref gotenresource.Reference, opts ...gotenresource.DeleteOption) error {
	if asLogConditionRef, ok := ref.(*Reference); ok {
		return a.DeleteLogCondition(ctx, asLogConditionRef, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected LogCondition, got: %s",
		ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) BatchGet(ctx context.Context, toGet []gotenresource.Reference, opts ...gotenresource.BatchGetOption) error {
	logConditionRefs := make([]*Reference, 0, len(toGet))
	for _, ref := range toGet {
		if asLogConditionRef, ok := ref.(*Reference); !ok {
			return status.Errorf(codes.Internal,
				"Unrecognized descriptor, expected LogCondition, got: %s",
				ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
		} else {
			logConditionRefs = append(logConditionRefs, asLogConditionRef)
		}
	}
	return a.BatchGetLogConditions(ctx, logConditionRefs, opts...)
}

func (a *anyCastAccess) GetResourceDescriptors() []gotenresource.Descriptor {
	return []gotenresource.Descriptor{
		GetDescriptor(),
	}
}
