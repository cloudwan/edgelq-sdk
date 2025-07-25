// Code generated by protoc-gen-goten-resource
// Resource: LogConditionTemplate
// DO NOT EDIT!!!

package log_condition_template

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
	policy_template "github.com/cloudwan/edgelq-sdk/alerting/resources/v1/policy_template"
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
	_ = &document.Document{}
	_ = &policy_template.PolicyTemplate{}
	_ = &rcommon.LogCndSpec{}
	_ = &meta.Meta{}
)

type LogConditionTemplateAccess interface {
	GetLogConditionTemplate(context.Context, *GetQuery, ...gotenresource.GetOption) (*LogConditionTemplate, error)
	BatchGetLogConditionTemplates(context.Context, []*Reference, ...gotenresource.BatchGetOption) error
	QueryLogConditionTemplates(context.Context, *ListQuery, ...gotenresource.QueryOption) (*QueryResultSnapshot, error)
	SearchLogConditionTemplates(context.Context, *SearchQuery, ...gotenresource.QueryOption) (*QueryResultSnapshot, error)
	WatchLogConditionTemplate(context.Context, *GetQuery, func(*LogConditionTemplateChange) error) error
	WatchLogConditionTemplates(context.Context, *WatchQuery, func(*QueryResultChange) error) error
	SaveLogConditionTemplate(context.Context, *LogConditionTemplate, ...gotenresource.SaveOption) error
	DeleteLogConditionTemplate(context.Context, *Reference, ...gotenresource.DeleteOption) error
}

type anyCastAccess struct {
	LogConditionTemplateAccess
}

func AsAnyCastAccess(access LogConditionTemplateAccess) gotenresource.Access {
	return &anyCastAccess{LogConditionTemplateAccess: access}
}

func (a *anyCastAccess) Get(ctx context.Context, q gotenresource.GetQuery, opts ...gotenresource.GetOption) (gotenresource.Resource, error) {
	if asLogConditionTemplateQuery, ok := q.(*GetQuery); ok {
		return a.GetLogConditionTemplate(ctx, asLogConditionTemplateQuery, opts...)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected LogConditionTemplate, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Query(ctx context.Context, q gotenresource.ListQuery, opts ...gotenresource.QueryOption) (gotenresource.QueryResultSnapshot, error) {
	if asLogConditionTemplateQuery, ok := q.(*ListQuery); ok {
		return a.QueryLogConditionTemplates(ctx, asLogConditionTemplateQuery, opts...)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected LogConditionTemplate, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Search(ctx context.Context, q gotenresource.SearchQuery, opts ...gotenresource.QueryOption) (gotenresource.QueryResultSnapshot, error) {
	if asLogConditionTemplateQuery, ok := q.(*SearchQuery); ok {
		return a.SearchLogConditionTemplates(ctx, asLogConditionTemplateQuery, opts...)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected LogConditionTemplate, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Watch(ctx context.Context, q gotenresource.GetQuery, cb func(ch gotenresource.ResourceChange) error) error {
	if asLogConditionTemplateQuery, ok := q.(*GetQuery); ok {
		return a.WatchLogConditionTemplate(ctx, asLogConditionTemplateQuery, func(change *LogConditionTemplateChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected LogConditionTemplate, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) WatchQuery(ctx context.Context, q gotenresource.WatchQuery, cb func(ch gotenresource.QueryResultChange) error) error {
	if asLogConditionTemplateQuery, ok := q.(*WatchQuery); ok {
		return a.WatchLogConditionTemplates(ctx, asLogConditionTemplateQuery, func(change *QueryResultChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected LogConditionTemplate, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Save(ctx context.Context, res gotenresource.Resource, opts ...gotenresource.SaveOption) error {
	if asLogConditionTemplateRes, ok := res.(*LogConditionTemplate); ok {
		return a.SaveLogConditionTemplate(ctx, asLogConditionTemplateRes, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected LogConditionTemplate, got: %s",
		res.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Delete(ctx context.Context, ref gotenresource.Reference, opts ...gotenresource.DeleteOption) error {
	if asLogConditionTemplateRef, ok := ref.(*Reference); ok {
		return a.DeleteLogConditionTemplate(ctx, asLogConditionTemplateRef, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected LogConditionTemplate, got: %s",
		ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) BatchGet(ctx context.Context, toGet []gotenresource.Reference, opts ...gotenresource.BatchGetOption) error {
	logConditionTemplateRefs := make([]*Reference, 0, len(toGet))
	for _, ref := range toGet {
		if asLogConditionTemplateRef, ok := ref.(*Reference); !ok {
			return status.Errorf(codes.Internal,
				"Unrecognized descriptor, expected LogConditionTemplate, got: %s",
				ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
		} else {
			logConditionTemplateRefs = append(logConditionTemplateRefs, asLogConditionTemplateRef)
		}
	}
	return a.BatchGetLogConditionTemplates(ctx, logConditionTemplateRefs, opts...)
}

func (a *anyCastAccess) GetResourceDescriptors() []gotenresource.Descriptor {
	return []gotenresource.Descriptor{
		GetDescriptor(),
	}
}
