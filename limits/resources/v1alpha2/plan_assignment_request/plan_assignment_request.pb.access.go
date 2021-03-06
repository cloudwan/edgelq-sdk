// Code generated by protoc-gen-goten-resource
// Resource: PlanAssignmentRequest
// DO NOT EDIT!!!

package plan_assignment_request

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
	iam_organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/organization"
	iam_project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/project"
	common "github.com/cloudwan/edgelq-sdk/limits/resources/v1alpha2/common"
	plan "github.com/cloudwan/edgelq-sdk/limits/resources/v1alpha2/plan"
	plan_assignment "github.com/cloudwan/edgelq-sdk/limits/resources/v1alpha2/plan_assignment"
	meta_service "github.com/cloudwan/edgelq-sdk/meta/resources/v1alpha2/service"
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
	_ = &iam_organization.Organization{}
	_ = &iam_project.Project{}
	_ = &common.Allowance{}
	_ = &plan.Plan{}
	_ = &plan_assignment.PlanAssignment{}
	_ = &meta_service.Service{}
)

type PlanAssignmentRequestAccess interface {
	GetPlanAssignmentRequest(context.Context, *GetQuery) (*PlanAssignmentRequest, error)
	BatchGetPlanAssignmentRequests(context.Context, []*Reference, ...gotenresource.BatchGetOption) error
	QueryPlanAssignmentRequests(context.Context, *ListQuery) (*QueryResultSnapshot, error)
	WatchPlanAssignmentRequest(context.Context, *GetQuery, func(*PlanAssignmentRequestChange) error) error
	WatchPlanAssignmentRequests(context.Context, *WatchQuery, func(*QueryResultChange) error) error
	SavePlanAssignmentRequest(context.Context, *PlanAssignmentRequest, ...gotenresource.SaveOption) error
	DeletePlanAssignmentRequest(context.Context, *Reference, ...gotenresource.DeleteOption) error
}

type anyCastAccess struct {
	PlanAssignmentRequestAccess
}

func AsAnyCastAccess(access PlanAssignmentRequestAccess) gotenresource.Access {
	return &anyCastAccess{PlanAssignmentRequestAccess: access}
}

func (a *anyCastAccess) Get(ctx context.Context, q gotenresource.GetQuery) (gotenresource.Resource, error) {
	if asPlanAssignmentRequestQuery, ok := q.(*GetQuery); ok {
		return a.GetPlanAssignmentRequest(ctx, asPlanAssignmentRequestQuery)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected PlanAssignmentRequest, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Query(ctx context.Context, q gotenresource.ListQuery) (gotenresource.QueryResultSnapshot, error) {
	if asPlanAssignmentRequestQuery, ok := q.(*ListQuery); ok {
		return a.QueryPlanAssignmentRequests(ctx, asPlanAssignmentRequestQuery)
	}
	return nil, status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected PlanAssignmentRequest, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Search(ctx context.Context, q gotenresource.SearchQuery) (gotenresource.SearchQueryResultSnapshot, error) {
	return nil, status.Errorf(codes.Internal, "Search is not available for PlanAssignmentRequest")
}

func (a *anyCastAccess) Watch(ctx context.Context, q gotenresource.GetQuery, cb func(ch gotenresource.ResourceChange) error) error {
	if asPlanAssignmentRequestQuery, ok := q.(*GetQuery); ok {
		return a.WatchPlanAssignmentRequest(ctx, asPlanAssignmentRequestQuery, func(change *PlanAssignmentRequestChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected PlanAssignmentRequest, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) WatchQuery(ctx context.Context, q gotenresource.WatchQuery, cb func(ch gotenresource.QueryResultChange) error) error {
	if asPlanAssignmentRequestQuery, ok := q.(*WatchQuery); ok {
		return a.WatchPlanAssignmentRequests(ctx, asPlanAssignmentRequestQuery, func(change *QueryResultChange) error {
			return cb(change)
		})
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected PlanAssignmentRequest, got: %s",
		q.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Save(ctx context.Context, res gotenresource.Resource, opts ...gotenresource.SaveOption) error {
	if asPlanAssignmentRequestRes, ok := res.(*PlanAssignmentRequest); ok {
		return a.SavePlanAssignmentRequest(ctx, asPlanAssignmentRequestRes, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected PlanAssignmentRequest, got: %s",
		res.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) Delete(ctx context.Context, ref gotenresource.Reference, opts ...gotenresource.DeleteOption) error {
	if asPlanAssignmentRequestRef, ok := ref.(*Reference); ok {
		return a.DeletePlanAssignmentRequest(ctx, asPlanAssignmentRequestRef, opts...)
	}
	return status.Errorf(codes.Internal,
		"Unrecognized descriptor, expected PlanAssignmentRequest, got: %s",
		ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
}

func (a *anyCastAccess) BatchGet(ctx context.Context, toGet []gotenresource.Reference, opts ...gotenresource.BatchGetOption) error {
	planAssignmentRequestRefs := make([]*Reference, 0, len(toGet))
	for _, ref := range toGet {
		if asPlanAssignmentRequestRef, ok := ref.(*Reference); !ok {
			return status.Errorf(codes.Internal,
				"Unrecognized descriptor, expected PlanAssignmentRequest, got: %s",
				ref.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName())
		} else {
			planAssignmentRequestRefs = append(planAssignmentRequestRefs, asPlanAssignmentRequestRef)
		}
	}
	return a.BatchGetPlanAssignmentRequests(ctx, planAssignmentRequestRefs, opts...)
}

func (a *anyCastAccess) GetResourceDescriptors() []gotenresource.Descriptor {
	return []gotenresource.Descriptor{
		GetDescriptor(),
	}
}
