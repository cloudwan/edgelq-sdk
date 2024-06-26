// Code generated by protoc-gen-goten-object
// File: edgelq/limits/proto/v1alpha2/common.proto
// DO NOT EDIT!!!

package common

// proto imports
import (
	meta_resource "github.com/cloudwan/edgelq-sdk/meta/resources/v1alpha2/resource"
	meta_service "github.com/cloudwan/edgelq-sdk/meta/resources/v1alpha2/service"
	meta "github.com/cloudwan/goten-sdk/types/meta"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

// make sure we're using proto imports
var (
	_ = &meta_resource.Resource{}
	_ = &meta_service.Service{}
	_ = &timestamppb.Timestamp{}
	_ = &meta.Meta{}
)

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

type RegionalDistributionFieldPathBuilder struct{}

func NewRegionalDistributionFieldPathBuilder() RegionalDistributionFieldPathBuilder {
	return RegionalDistributionFieldPathBuilder{}
}
func (RegionalDistributionFieldPathBuilder) Resource() RegionalDistributionPathSelectorResource {
	return RegionalDistributionPathSelectorResource{}
}
func (RegionalDistributionFieldPathBuilder) LimitsByRegion() RegionalDistributionPathSelectorLimitsByRegion {
	return RegionalDistributionPathSelectorLimitsByRegion{}
}

type RegionalDistributionPathSelectorResource struct{}

func (RegionalDistributionPathSelectorResource) FieldPath() *RegionalDistribution_FieldTerminalPath {
	return &RegionalDistribution_FieldTerminalPath{selector: RegionalDistribution_FieldPathSelectorResource}
}

func (s RegionalDistributionPathSelectorResource) WithValue(value *meta_resource.Reference) *RegionalDistribution_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*RegionalDistribution_FieldTerminalPathValue)
}

func (s RegionalDistributionPathSelectorResource) WithArrayOfValues(values []*meta_resource.Reference) *RegionalDistribution_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*RegionalDistribution_FieldTerminalPathArrayOfValues)
}

type RegionalDistributionPathSelectorLimitsByRegion struct{}

func (RegionalDistributionPathSelectorLimitsByRegion) FieldPath() *RegionalDistribution_FieldTerminalPath {
	return &RegionalDistribution_FieldTerminalPath{selector: RegionalDistribution_FieldPathSelectorLimitsByRegion}
}

func (s RegionalDistributionPathSelectorLimitsByRegion) WithValue(value map[string]int64) *RegionalDistribution_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*RegionalDistribution_FieldTerminalPathValue)
}

func (s RegionalDistributionPathSelectorLimitsByRegion) WithArrayOfValues(values []map[string]int64) *RegionalDistribution_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*RegionalDistribution_FieldTerminalPathArrayOfValues)
}

func (RegionalDistributionPathSelectorLimitsByRegion) WithKey(key string) RegionalDistributionMapPathSelectorLimitsByRegion {
	return RegionalDistributionMapPathSelectorLimitsByRegion{key: key}
}

type RegionalDistributionMapPathSelectorLimitsByRegion struct {
	key string
}

func (s RegionalDistributionMapPathSelectorLimitsByRegion) FieldPath() *RegionalDistribution_FieldPathMap {
	return &RegionalDistribution_FieldPathMap{selector: RegionalDistribution_FieldPathSelectorLimitsByRegion, key: s.key}
}

func (s RegionalDistributionMapPathSelectorLimitsByRegion) WithValue(value int64) *RegionalDistribution_FieldPathMapValue {
	return s.FieldPath().WithIValue(value).(*RegionalDistribution_FieldPathMapValue)
}

func (s RegionalDistributionMapPathSelectorLimitsByRegion) WithArrayOfValues(values []int64) *RegionalDistribution_FieldPathMapArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*RegionalDistribution_FieldPathMapArrayOfValues)
}
