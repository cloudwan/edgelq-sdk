// Code generated by protoc-gen-goten-object
// File: edgelq/limits/proto/v1/common.proto
// DO NOT EDIT!!!

package common

// proto imports
import (
	iam_iam_common "github.com/cloudwan/edgelq-sdk/iam/resources/v1/common"
	iam_organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1/organization"
	plan "github.com/cloudwan/edgelq-sdk/limits/resources/v1/plan"
	meta_resource "github.com/cloudwan/goten-sdk/meta-service/resources/v1/resource"
	meta_service "github.com/cloudwan/goten-sdk/meta-service/resources/v1/service"
	meta "github.com/cloudwan/goten-sdk/types/meta"
	multi_region_policy "github.com/cloudwan/goten-sdk/types/multi_region_policy"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

// make sure we're using proto imports
var (
	_ = &iam_iam_common.PCR{}
	_ = &iam_organization.Organization{}
	_ = &plan.Plan{}
	_ = &timestamppb.Timestamp{}
	_ = &meta_resource.Resource{}
	_ = &meta_service.Service{}
	_ = &meta.Meta{}
	_ = &multi_region_policy.MultiRegionPolicy{}
)

type RegionalPlanAssignmentFieldPathBuilder struct{}

func NewRegionalPlanAssignmentFieldPathBuilder() RegionalPlanAssignmentFieldPathBuilder {
	return RegionalPlanAssignmentFieldPathBuilder{}
}
func (RegionalPlanAssignmentFieldPathBuilder) Plan() RegionalPlanAssignmentPathSelectorPlan {
	return RegionalPlanAssignmentPathSelectorPlan{}
}
func (RegionalPlanAssignmentFieldPathBuilder) Region() RegionalPlanAssignmentPathSelectorRegion {
	return RegionalPlanAssignmentPathSelectorRegion{}
}
func (RegionalPlanAssignmentFieldPathBuilder) PlanGeneration() RegionalPlanAssignmentPathSelectorPlanGeneration {
	return RegionalPlanAssignmentPathSelectorPlanGeneration{}
}

type RegionalPlanAssignmentPathSelectorPlan struct{}

func (RegionalPlanAssignmentPathSelectorPlan) FieldPath() *RegionalPlanAssignment_FieldTerminalPath {
	return &RegionalPlanAssignment_FieldTerminalPath{selector: RegionalPlanAssignment_FieldPathSelectorPlan}
}

func (s RegionalPlanAssignmentPathSelectorPlan) WithValue(value *plan.Reference) *RegionalPlanAssignment_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*RegionalPlanAssignment_FieldTerminalPathValue)
}

func (s RegionalPlanAssignmentPathSelectorPlan) WithArrayOfValues(values []*plan.Reference) *RegionalPlanAssignment_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*RegionalPlanAssignment_FieldTerminalPathArrayOfValues)
}

type RegionalPlanAssignmentPathSelectorRegion struct{}

func (RegionalPlanAssignmentPathSelectorRegion) FieldPath() *RegionalPlanAssignment_FieldTerminalPath {
	return &RegionalPlanAssignment_FieldTerminalPath{selector: RegionalPlanAssignment_FieldPathSelectorRegion}
}

func (s RegionalPlanAssignmentPathSelectorRegion) WithValue(value string) *RegionalPlanAssignment_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*RegionalPlanAssignment_FieldTerminalPathValue)
}

func (s RegionalPlanAssignmentPathSelectorRegion) WithArrayOfValues(values []string) *RegionalPlanAssignment_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*RegionalPlanAssignment_FieldTerminalPathArrayOfValues)
}

type RegionalPlanAssignmentPathSelectorPlanGeneration struct{}

func (RegionalPlanAssignmentPathSelectorPlanGeneration) FieldPath() *RegionalPlanAssignment_FieldTerminalPath {
	return &RegionalPlanAssignment_FieldTerminalPath{selector: RegionalPlanAssignment_FieldPathSelectorPlanGeneration}
}

func (s RegionalPlanAssignmentPathSelectorPlanGeneration) WithValue(value int64) *RegionalPlanAssignment_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*RegionalPlanAssignment_FieldTerminalPathValue)
}

func (s RegionalPlanAssignmentPathSelectorPlanGeneration) WithArrayOfValues(values []int64) *RegionalPlanAssignment_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*RegionalPlanAssignment_FieldTerminalPathArrayOfValues)
}

type AllowanceFieldPathBuilder struct{}

func NewAllowanceFieldPathBuilder() AllowanceFieldPathBuilder {
	return AllowanceFieldPathBuilder{}
}
func (AllowanceFieldPathBuilder) Resource() AllowancePathSelectorResource {
	return AllowancePathSelectorResource{}
}
func (AllowanceFieldPathBuilder) Value() AllowancePathSelectorValue {
	return AllowancePathSelectorValue{}
}
func (AllowanceFieldPathBuilder) Region() AllowancePathSelectorRegion {
	return AllowancePathSelectorRegion{}
}

type AllowancePathSelectorResource struct{}

func (AllowancePathSelectorResource) FieldPath() *Allowance_FieldTerminalPath {
	return &Allowance_FieldTerminalPath{selector: Allowance_FieldPathSelectorResource}
}

func (s AllowancePathSelectorResource) WithValue(value *meta_resource.Reference) *Allowance_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*Allowance_FieldTerminalPathValue)
}

func (s AllowancePathSelectorResource) WithArrayOfValues(values []*meta_resource.Reference) *Allowance_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Allowance_FieldTerminalPathArrayOfValues)
}

type AllowancePathSelectorValue struct{}

func (AllowancePathSelectorValue) FieldPath() *Allowance_FieldTerminalPath {
	return &Allowance_FieldTerminalPath{selector: Allowance_FieldPathSelectorValue}
}

func (s AllowancePathSelectorValue) WithValue(value int64) *Allowance_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*Allowance_FieldTerminalPathValue)
}

func (s AllowancePathSelectorValue) WithArrayOfValues(values []int64) *Allowance_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Allowance_FieldTerminalPathArrayOfValues)
}

type AllowancePathSelectorRegion struct{}

func (AllowancePathSelectorRegion) FieldPath() *Allowance_FieldTerminalPath {
	return &Allowance_FieldTerminalPath{selector: Allowance_FieldPathSelectorRegion}
}

func (s AllowancePathSelectorRegion) WithValue(value string) *Allowance_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*Allowance_FieldTerminalPathValue)
}

func (s AllowancePathSelectorRegion) WithArrayOfValues(values []string) *Allowance_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Allowance_FieldTerminalPathArrayOfValues)
}
