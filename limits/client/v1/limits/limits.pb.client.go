// Code generated by protoc-gen-goten-client
// Service: Limits
// DO NOT EDIT!!!

package limits_client

import (
	"google.golang.org/grpc"
)

// proto imports
import (
	accepted_plan_client "github.com/cloudwan/edgelq-sdk/limits/client/v1/accepted_plan"
	limit_client "github.com/cloudwan/edgelq-sdk/limits/client/v1/limit"
	limit_pool_client "github.com/cloudwan/edgelq-sdk/limits/client/v1/limit_pool"
	plan_client "github.com/cloudwan/edgelq-sdk/limits/client/v1/plan"
	plan_assignment_client "github.com/cloudwan/edgelq-sdk/limits/client/v1/plan_assignment"
	plan_assignment_request_client "github.com/cloudwan/edgelq-sdk/limits/client/v1/plan_assignment_request"
	accepted_plan "github.com/cloudwan/edgelq-sdk/limits/resources/v1/accepted_plan"
	limit "github.com/cloudwan/edgelq-sdk/limits/resources/v1/limit"
	limit_pool "github.com/cloudwan/edgelq-sdk/limits/resources/v1/limit_pool"
	plan "github.com/cloudwan/edgelq-sdk/limits/resources/v1/plan"
	plan_assignment "github.com/cloudwan/edgelq-sdk/limits/resources/v1/plan_assignment"
	plan_assignment_request "github.com/cloudwan/edgelq-sdk/limits/resources/v1/plan_assignment_request"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = grpc.ClientConn{}
)

// make sure we're using proto imports
var (
	_ = &accepted_plan.AcceptedPlan{}
	_ = &accepted_plan_client.GetAcceptedPlanRequest{}
	_ = &limit.Limit{}
	_ = &limit_pool.LimitPool{}
	_ = &limit_pool_client.GetLimitPoolRequest{}
	_ = &limit_client.GetLimitRequest{}
	_ = &plan.Plan{}
	_ = &plan_assignment.PlanAssignment{}
	_ = &plan_assignment_request.PlanAssignmentRequest{}
	_ = &plan_assignment_request_client.GetPlanAssignmentRequestRequest{}
	_ = &plan_assignment_client.GetPlanAssignmentRequest{}
	_ = &plan_client.GetPlanRequest{}
)

type LimitsClient interface {
	accepted_plan_client.AcceptedPlanServiceClient
	limit_pool_client.LimitPoolServiceClient
	limit_client.LimitServiceClient
	plan_assignment_request_client.PlanAssignmentRequestServiceClient
	plan_assignment_client.PlanAssignmentServiceClient
	plan_client.PlanServiceClient
}

type limitsClient struct {
	accepted_plan_client.AcceptedPlanServiceClient
	limit_pool_client.LimitPoolServiceClient
	limit_client.LimitServiceClient
	plan_assignment_request_client.PlanAssignmentRequestServiceClient
	plan_assignment_client.PlanAssignmentServiceClient
	plan_client.PlanServiceClient
}

func NewLimitsClient(cc grpc.ClientConnInterface) LimitsClient {
	return &limitsClient{
		AcceptedPlanServiceClient:          accepted_plan_client.NewAcceptedPlanServiceClient(cc),
		LimitPoolServiceClient:             limit_pool_client.NewLimitPoolServiceClient(cc),
		LimitServiceClient:                 limit_client.NewLimitServiceClient(cc),
		PlanAssignmentRequestServiceClient: plan_assignment_request_client.NewPlanAssignmentRequestServiceClient(cc),
		PlanAssignmentServiceClient:        plan_assignment_client.NewPlanAssignmentServiceClient(cc),
		PlanServiceClient:                  plan_client.NewPlanServiceClient(cc),
	}
}
