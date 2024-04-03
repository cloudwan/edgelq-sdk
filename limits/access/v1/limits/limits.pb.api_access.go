// Code generated by protoc-gen-goten-access
// Service: Limits
// DO NOT EDIT!!!

package limits_access

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"

	accepted_plan_access "github.com/cloudwan/edgelq-sdk/limits/access/v1/accepted_plan"
	limit_access "github.com/cloudwan/edgelq-sdk/limits/access/v1/limit"
	limit_pool_access "github.com/cloudwan/edgelq-sdk/limits/access/v1/limit_pool"
	plan_access "github.com/cloudwan/edgelq-sdk/limits/access/v1/plan"
	plan_assignment_access "github.com/cloudwan/edgelq-sdk/limits/access/v1/plan_assignment"
	plan_assignment_request_access "github.com/cloudwan/edgelq-sdk/limits/access/v1/plan_assignment_request"
	limits_client "github.com/cloudwan/edgelq-sdk/limits/client/v1/limits"
	accepted_plan "github.com/cloudwan/edgelq-sdk/limits/resources/v1/accepted_plan"
	limit "github.com/cloudwan/edgelq-sdk/limits/resources/v1/limit"
	limit_pool "github.com/cloudwan/edgelq-sdk/limits/resources/v1/limit_pool"
	plan "github.com/cloudwan/edgelq-sdk/limits/resources/v1/plan"
	plan_assignment "github.com/cloudwan/edgelq-sdk/limits/resources/v1/plan_assignment"
	plan_assignment_request "github.com/cloudwan/edgelq-sdk/limits/resources/v1/plan_assignment_request"
)

type LimitsApiAccess interface {
	gotenresource.Access

	accepted_plan.AcceptedPlanAccess
	limit.LimitAccess
	limit_pool.LimitPoolAccess
	plan.PlanAccess
	plan_assignment.PlanAssignmentAccess
	plan_assignment_request.PlanAssignmentRequestAccess
}

type apiLimitsAccess struct {
	gotenresource.Access

	accepted_plan.AcceptedPlanAccess
	limit.LimitAccess
	limit_pool.LimitPoolAccess
	plan.PlanAccess
	plan_assignment.PlanAssignmentAccess
	plan_assignment_request.PlanAssignmentRequestAccess
}

func NewApiAccess(client limits_client.LimitsClient) LimitsApiAccess {

	acceptedPlanAccess := accepted_plan_access.NewApiAcceptedPlanAccess(client)
	limitAccess := limit_access.NewApiLimitAccess(client)
	limitPoolAccess := limit_pool_access.NewApiLimitPoolAccess(client)
	planAccess := plan_access.NewApiPlanAccess(client)
	planAssignmentAccess := plan_assignment_access.NewApiPlanAssignmentAccess(client)
	planAssignmentRequestAccess := plan_assignment_request_access.NewApiPlanAssignmentRequestAccess(client)

	return &apiLimitsAccess{
		Access: gotenresource.NewCompositeAccess(

			accepted_plan.AsAnyCastAccess(acceptedPlanAccess),
			limit.AsAnyCastAccess(limitAccess),
			limit_pool.AsAnyCastAccess(limitPoolAccess),
			plan.AsAnyCastAccess(planAccess),
			plan_assignment.AsAnyCastAccess(planAssignmentAccess),
			plan_assignment_request.AsAnyCastAccess(planAssignmentRequestAccess),
		),

		AcceptedPlanAccess:          acceptedPlanAccess,
		LimitAccess:                 limitAccess,
		LimitPoolAccess:             limitPoolAccess,
		PlanAccess:                  planAccess,
		PlanAssignmentAccess:        planAssignmentAccess,
		PlanAssignmentRequestAccess: planAssignmentRequestAccess,
	}
}