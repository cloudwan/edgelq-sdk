// Code generated by protoc-gen-goten-client
// Service: Limits
// DO NOT EDIT!!!

package limits_client

import (
	gotenclient "github.com/cloudwan/goten-sdk/runtime/client"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	accepted_plan_client "github.com/cloudwan/edgelq-sdk/limits/client/v1alpha2/accepted_plan"
	limit_client "github.com/cloudwan/edgelq-sdk/limits/client/v1alpha2/limit"
	limit_pool_client "github.com/cloudwan/edgelq-sdk/limits/client/v1alpha2/limit_pool"
	plan_client "github.com/cloudwan/edgelq-sdk/limits/client/v1alpha2/plan"
	plan_assignment_client "github.com/cloudwan/edgelq-sdk/limits/client/v1alpha2/plan_assignment"
	plan_assignment_request_client "github.com/cloudwan/edgelq-sdk/limits/client/v1alpha2/plan_assignment_request"
	accepted_plan "github.com/cloudwan/edgelq-sdk/limits/resources/v1alpha2/accepted_plan"
	limit "github.com/cloudwan/edgelq-sdk/limits/resources/v1alpha2/limit"
	limit_pool "github.com/cloudwan/edgelq-sdk/limits/resources/v1alpha2/limit_pool"
	plan "github.com/cloudwan/edgelq-sdk/limits/resources/v1alpha2/plan"
	plan_assignment "github.com/cloudwan/edgelq-sdk/limits/resources/v1alpha2/plan_assignment"
	plan_assignment_request "github.com/cloudwan/edgelq-sdk/limits/resources/v1alpha2/plan_assignment_request"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = new(gotenclient.MethodDescriptor)
	_ = gotenresource.WildcardId
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

var (
	descriptorInitialized bool
	limitsDescriptor      *LimitsDescriptor
)

type LimitsDescriptor struct{}

func (d *LimitsDescriptor) GetServiceName() string {
	return "limits"
}

func (d *LimitsDescriptor) GetServiceDomain() string {
	return "limits.edgelq.com"
}

func (d *LimitsDescriptor) GetVersion() string {
	return "v1alpha2"
}

func (d *LimitsDescriptor) GetNextVersion() string {

	return "v1"
}

func (d *LimitsDescriptor) AllResourceDescriptors() []gotenresource.Descriptor {
	return []gotenresource.Descriptor{
		accepted_plan.GetDescriptor(),
		limit.GetDescriptor(),
		limit_pool.GetDescriptor(),
		plan.GetDescriptor(),
		plan_assignment.GetDescriptor(),
		plan_assignment_request.GetDescriptor(),
	}
}

func (d *LimitsDescriptor) AllApiDescriptors() []gotenclient.ApiDescriptor {
	return []gotenclient.ApiDescriptor{
		accepted_plan_client.GetAcceptedPlanServiceDescriptor(),
		limit_pool_client.GetLimitPoolServiceDescriptor(),
		limit_client.GetLimitServiceDescriptor(),
		plan_assignment_request_client.GetPlanAssignmentRequestServiceDescriptor(),
		plan_assignment_client.GetPlanAssignmentServiceDescriptor(),
		plan_client.GetPlanServiceDescriptor(),
	}
}

func (d *LimitsDescriptor) AllImportedServiceInfos() []gotenclient.ServiceImportInfo {
	return []gotenclient.ServiceImportInfo{
		{
			Domain:  "iam.edgelq.com",
			Version: "v1alpha2",
		},
		{
			Domain:  "meta.edgelq.com",
			Version: "v1alpha2",
		},
	}
}

func GetLimitsDescriptor() *LimitsDescriptor {
	return limitsDescriptor
}

func initDescriptor() {
	limitsDescriptor = &LimitsDescriptor{}
	gotenclient.GetRegistry().RegisterSvcDescriptor(limitsDescriptor)
}

func init() {
	if !descriptorInitialized {
		initDescriptor()
		descriptorInitialized = true
	}
}
