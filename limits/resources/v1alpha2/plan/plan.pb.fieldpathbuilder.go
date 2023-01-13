// Code generated by protoc-gen-goten-object
// File: edgelq/limits/proto/v1alpha2/plan.proto
// DO NOT EDIT!!!

package plan

// proto imports
import (
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
	iam_iam_common "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/common"
	common "github.com/cloudwan/edgelq-sdk/limits/resources/v1alpha2/common"
	meta_resource "github.com/cloudwan/edgelq-sdk/meta/resources/v1alpha2/resource"
	meta_service "github.com/cloudwan/edgelq-sdk/meta/resources/v1alpha2/service"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
)

// make sure we're using proto imports
var (
	_ = &ntt_meta.Meta{}
	_ = &iam_iam_common.PCR{}
	_ = &common.Allowance{}
	_ = &meta_resource.Resource{}
	_ = &meta_service.Service{}
	_ = &timestamp.Timestamp{}
)

type PlanFieldPathBuilder struct{}

func NewPlanFieldPathBuilder() PlanFieldPathBuilder {
	return PlanFieldPathBuilder{}
}
func (PlanFieldPathBuilder) Name() PlanPathSelectorName {
	return PlanPathSelectorName{}
}
func (PlanFieldPathBuilder) DisplayName() PlanPathSelectorDisplayName {
	return PlanPathSelectorDisplayName{}
}
func (PlanFieldPathBuilder) Service() PlanPathSelectorService {
	return PlanPathSelectorService{}
}
func (PlanFieldPathBuilder) ResourceLimits() PlanPathSelectorResourceLimits {
	return PlanPathSelectorResourceLimits{}
}
func (PlanFieldPathBuilder) PlanLevel() PlanPathSelectorPlanLevel {
	return PlanPathSelectorPlanLevel{}
}
func (PlanFieldPathBuilder) BusinessTier() PlanPathSelectorBusinessTier {
	return PlanPathSelectorBusinessTier{}
}
func (PlanFieldPathBuilder) Metadata() PlanPathSelectorMetadata {
	return PlanPathSelectorMetadata{}
}

type PlanPathSelectorName struct{}

func (PlanPathSelectorName) FieldPath() *Plan_FieldTerminalPath {
	return &Plan_FieldTerminalPath{selector: Plan_FieldPathSelectorName}
}

func (s PlanPathSelectorName) WithValue(value *Name) *Plan_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*Plan_FieldTerminalPathValue)
}

func (s PlanPathSelectorName) WithArrayOfValues(values []*Name) *Plan_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Plan_FieldTerminalPathArrayOfValues)
}

type PlanPathSelectorDisplayName struct{}

func (PlanPathSelectorDisplayName) FieldPath() *Plan_FieldTerminalPath {
	return &Plan_FieldTerminalPath{selector: Plan_FieldPathSelectorDisplayName}
}

func (s PlanPathSelectorDisplayName) WithValue(value string) *Plan_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*Plan_FieldTerminalPathValue)
}

func (s PlanPathSelectorDisplayName) WithArrayOfValues(values []string) *Plan_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Plan_FieldTerminalPathArrayOfValues)
}

type PlanPathSelectorService struct{}

func (PlanPathSelectorService) FieldPath() *Plan_FieldTerminalPath {
	return &Plan_FieldTerminalPath{selector: Plan_FieldPathSelectorService}
}

func (s PlanPathSelectorService) WithValue(value *meta_service.Reference) *Plan_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*Plan_FieldTerminalPathValue)
}

func (s PlanPathSelectorService) WithArrayOfValues(values []*meta_service.Reference) *Plan_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Plan_FieldTerminalPathArrayOfValues)
}

type PlanPathSelectorResourceLimits struct{}

func (PlanPathSelectorResourceLimits) FieldPath() *Plan_FieldTerminalPath {
	return &Plan_FieldTerminalPath{selector: Plan_FieldPathSelectorResourceLimits}
}

func (s PlanPathSelectorResourceLimits) WithValue(value []*common.Allowance) *Plan_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*Plan_FieldTerminalPathValue)
}

func (s PlanPathSelectorResourceLimits) WithArrayOfValues(values [][]*common.Allowance) *Plan_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Plan_FieldTerminalPathArrayOfValues)
}

func (s PlanPathSelectorResourceLimits) WithItemValue(value *common.Allowance) *Plan_FieldTerminalPathArrayItemValue {
	return s.FieldPath().WithIArrayItemValue(value).(*Plan_FieldTerminalPathArrayItemValue)
}
func (PlanPathSelectorResourceLimits) WithSubPath(subPath common.Allowance_FieldPath) *Plan_FieldSubPath {
	return &Plan_FieldSubPath{selector: Plan_FieldPathSelectorResourceLimits, subPath: subPath}
}

func (s PlanPathSelectorResourceLimits) WithSubValue(subPathValue common.Allowance_FieldPathValue) *Plan_FieldSubPathValue {
	return &Plan_FieldSubPathValue{Plan_FieldPath: s.WithSubPath(subPathValue), subPathValue: subPathValue}
}

func (s PlanPathSelectorResourceLimits) WithSubArrayOfValues(subPathArrayOfValues common.Allowance_FieldPathArrayOfValues) *Plan_FieldSubPathArrayOfValues {
	return &Plan_FieldSubPathArrayOfValues{Plan_FieldPath: s.WithSubPath(subPathArrayOfValues), subPathArrayOfValues: subPathArrayOfValues}
}

func (s PlanPathSelectorResourceLimits) WithSubArrayItemValue(subPathArrayItemValue common.Allowance_FieldPathArrayItemValue) *Plan_FieldSubPathArrayItemValue {
	return &Plan_FieldSubPathArrayItemValue{Plan_FieldPath: s.WithSubPath(subPathArrayItemValue), subPathItemValue: subPathArrayItemValue}
}

func (PlanPathSelectorResourceLimits) Resource() PlanPathSelectorResourceLimitsResource {
	return PlanPathSelectorResourceLimitsResource{}
}

func (PlanPathSelectorResourceLimits) Value() PlanPathSelectorResourceLimitsValue {
	return PlanPathSelectorResourceLimitsValue{}
}

type PlanPathSelectorResourceLimitsResource struct{}

func (PlanPathSelectorResourceLimitsResource) FieldPath() *Plan_FieldSubPath {
	return &Plan_FieldSubPath{
		selector: Plan_FieldPathSelectorResourceLimits,
		subPath:  common.NewAllowanceFieldPathBuilder().Resource().FieldPath(),
	}
}

func (s PlanPathSelectorResourceLimitsResource) WithValue(value *meta_resource.Reference) *Plan_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Plan_FieldSubPathValue)
}

func (s PlanPathSelectorResourceLimitsResource) WithArrayOfValues(values []*meta_resource.Reference) *Plan_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Plan_FieldSubPathArrayOfValues)
}

type PlanPathSelectorResourceLimitsValue struct{}

func (PlanPathSelectorResourceLimitsValue) FieldPath() *Plan_FieldSubPath {
	return &Plan_FieldSubPath{
		selector: Plan_FieldPathSelectorResourceLimits,
		subPath:  common.NewAllowanceFieldPathBuilder().Value().FieldPath(),
	}
}

func (s PlanPathSelectorResourceLimitsValue) WithValue(value int64) *Plan_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Plan_FieldSubPathValue)
}

func (s PlanPathSelectorResourceLimitsValue) WithArrayOfValues(values []int64) *Plan_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Plan_FieldSubPathArrayOfValues)
}

type PlanPathSelectorPlanLevel struct{}

func (PlanPathSelectorPlanLevel) FieldPath() *Plan_FieldTerminalPath {
	return &Plan_FieldTerminalPath{selector: Plan_FieldPathSelectorPlanLevel}
}

func (s PlanPathSelectorPlanLevel) WithValue(value Plan_PlanLevel) *Plan_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*Plan_FieldTerminalPathValue)
}

func (s PlanPathSelectorPlanLevel) WithArrayOfValues(values []Plan_PlanLevel) *Plan_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Plan_FieldTerminalPathArrayOfValues)
}

type PlanPathSelectorBusinessTier struct{}

func (PlanPathSelectorBusinessTier) FieldPath() *Plan_FieldTerminalPath {
	return &Plan_FieldTerminalPath{selector: Plan_FieldPathSelectorBusinessTier}
}

func (s PlanPathSelectorBusinessTier) WithValue(value iam_iam_common.BusinessTier) *Plan_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*Plan_FieldTerminalPathValue)
}

func (s PlanPathSelectorBusinessTier) WithArrayOfValues(values []iam_iam_common.BusinessTier) *Plan_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Plan_FieldTerminalPathArrayOfValues)
}

type PlanPathSelectorMetadata struct{}

func (PlanPathSelectorMetadata) FieldPath() *Plan_FieldTerminalPath {
	return &Plan_FieldTerminalPath{selector: Plan_FieldPathSelectorMetadata}
}

func (s PlanPathSelectorMetadata) WithValue(value *ntt_meta.Meta) *Plan_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*Plan_FieldTerminalPathValue)
}

func (s PlanPathSelectorMetadata) WithArrayOfValues(values []*ntt_meta.Meta) *Plan_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Plan_FieldTerminalPathArrayOfValues)
}

func (PlanPathSelectorMetadata) WithSubPath(subPath ntt_meta.Meta_FieldPath) *Plan_FieldSubPath {
	return &Plan_FieldSubPath{selector: Plan_FieldPathSelectorMetadata, subPath: subPath}
}

func (s PlanPathSelectorMetadata) WithSubValue(subPathValue ntt_meta.Meta_FieldPathValue) *Plan_FieldSubPathValue {
	return &Plan_FieldSubPathValue{Plan_FieldPath: s.WithSubPath(subPathValue), subPathValue: subPathValue}
}

func (s PlanPathSelectorMetadata) WithSubArrayOfValues(subPathArrayOfValues ntt_meta.Meta_FieldPathArrayOfValues) *Plan_FieldSubPathArrayOfValues {
	return &Plan_FieldSubPathArrayOfValues{Plan_FieldPath: s.WithSubPath(subPathArrayOfValues), subPathArrayOfValues: subPathArrayOfValues}
}

func (s PlanPathSelectorMetadata) WithSubArrayItemValue(subPathArrayItemValue ntt_meta.Meta_FieldPathArrayItemValue) *Plan_FieldSubPathArrayItemValue {
	return &Plan_FieldSubPathArrayItemValue{Plan_FieldPath: s.WithSubPath(subPathArrayItemValue), subPathItemValue: subPathArrayItemValue}
}

func (PlanPathSelectorMetadata) CreateTime() PlanPathSelectorMetadataCreateTime {
	return PlanPathSelectorMetadataCreateTime{}
}

func (PlanPathSelectorMetadata) UpdateTime() PlanPathSelectorMetadataUpdateTime {
	return PlanPathSelectorMetadataUpdateTime{}
}

func (PlanPathSelectorMetadata) DeleteTime() PlanPathSelectorMetadataDeleteTime {
	return PlanPathSelectorMetadataDeleteTime{}
}

func (PlanPathSelectorMetadata) Uuid() PlanPathSelectorMetadataUuid {
	return PlanPathSelectorMetadataUuid{}
}

func (PlanPathSelectorMetadata) Tags() PlanPathSelectorMetadataTags {
	return PlanPathSelectorMetadataTags{}
}

func (PlanPathSelectorMetadata) Labels() PlanPathSelectorMetadataLabels {
	return PlanPathSelectorMetadataLabels{}
}

func (PlanPathSelectorMetadata) Annotations() PlanPathSelectorMetadataAnnotations {
	return PlanPathSelectorMetadataAnnotations{}
}

func (PlanPathSelectorMetadata) Generation() PlanPathSelectorMetadataGeneration {
	return PlanPathSelectorMetadataGeneration{}
}

func (PlanPathSelectorMetadata) ResourceVersion() PlanPathSelectorMetadataResourceVersion {
	return PlanPathSelectorMetadataResourceVersion{}
}

func (PlanPathSelectorMetadata) OwnerReferences() PlanPathSelectorMetadataOwnerReferences {
	return PlanPathSelectorMetadataOwnerReferences{}
}

func (PlanPathSelectorMetadata) Shards() PlanPathSelectorMetadataShards {
	return PlanPathSelectorMetadataShards{}
}

func (PlanPathSelectorMetadata) Syncing() PlanPathSelectorMetadataSyncing {
	return PlanPathSelectorMetadataSyncing{}
}

func (PlanPathSelectorMetadata) Lifecycle() PlanPathSelectorMetadataLifecycle {
	return PlanPathSelectorMetadataLifecycle{}
}

type PlanPathSelectorMetadataCreateTime struct{}

func (PlanPathSelectorMetadataCreateTime) FieldPath() *Plan_FieldSubPath {
	return &Plan_FieldSubPath{
		selector: Plan_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().CreateTime().FieldPath(),
	}
}

func (s PlanPathSelectorMetadataCreateTime) WithValue(value *timestamp.Timestamp) *Plan_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Plan_FieldSubPathValue)
}

func (s PlanPathSelectorMetadataCreateTime) WithArrayOfValues(values []*timestamp.Timestamp) *Plan_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Plan_FieldSubPathArrayOfValues)
}

type PlanPathSelectorMetadataUpdateTime struct{}

func (PlanPathSelectorMetadataUpdateTime) FieldPath() *Plan_FieldSubPath {
	return &Plan_FieldSubPath{
		selector: Plan_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().UpdateTime().FieldPath(),
	}
}

func (s PlanPathSelectorMetadataUpdateTime) WithValue(value *timestamp.Timestamp) *Plan_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Plan_FieldSubPathValue)
}

func (s PlanPathSelectorMetadataUpdateTime) WithArrayOfValues(values []*timestamp.Timestamp) *Plan_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Plan_FieldSubPathArrayOfValues)
}

type PlanPathSelectorMetadataDeleteTime struct{}

func (PlanPathSelectorMetadataDeleteTime) FieldPath() *Plan_FieldSubPath {
	return &Plan_FieldSubPath{
		selector: Plan_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().DeleteTime().FieldPath(),
	}
}

func (s PlanPathSelectorMetadataDeleteTime) WithValue(value *timestamp.Timestamp) *Plan_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Plan_FieldSubPathValue)
}

func (s PlanPathSelectorMetadataDeleteTime) WithArrayOfValues(values []*timestamp.Timestamp) *Plan_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Plan_FieldSubPathArrayOfValues)
}

type PlanPathSelectorMetadataUuid struct{}

func (PlanPathSelectorMetadataUuid) FieldPath() *Plan_FieldSubPath {
	return &Plan_FieldSubPath{
		selector: Plan_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Uuid().FieldPath(),
	}
}

func (s PlanPathSelectorMetadataUuid) WithValue(value string) *Plan_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Plan_FieldSubPathValue)
}

func (s PlanPathSelectorMetadataUuid) WithArrayOfValues(values []string) *Plan_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Plan_FieldSubPathArrayOfValues)
}

type PlanPathSelectorMetadataTags struct{}

func (PlanPathSelectorMetadataTags) FieldPath() *Plan_FieldSubPath {
	return &Plan_FieldSubPath{
		selector: Plan_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Tags().FieldPath(),
	}
}

func (s PlanPathSelectorMetadataTags) WithValue(value []string) *Plan_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Plan_FieldSubPathValue)
}

func (s PlanPathSelectorMetadataTags) WithArrayOfValues(values [][]string) *Plan_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Plan_FieldSubPathArrayOfValues)
}

func (s PlanPathSelectorMetadataTags) WithItemValue(value string) *Plan_FieldSubPathArrayItemValue {
	return s.FieldPath().WithIArrayItemValue(value).(*Plan_FieldSubPathArrayItemValue)
}

type PlanPathSelectorMetadataLabels struct{}

func (PlanPathSelectorMetadataLabels) FieldPath() *Plan_FieldSubPath {
	return &Plan_FieldSubPath{
		selector: Plan_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Labels().FieldPath(),
	}
}

func (s PlanPathSelectorMetadataLabels) WithValue(value map[string]string) *Plan_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Plan_FieldSubPathValue)
}

func (s PlanPathSelectorMetadataLabels) WithArrayOfValues(values []map[string]string) *Plan_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Plan_FieldSubPathArrayOfValues)
}

func (PlanPathSelectorMetadataLabels) WithKey(key string) PlanMapPathSelectorMetadataLabels {
	return PlanMapPathSelectorMetadataLabels{key: key}
}

type PlanMapPathSelectorMetadataLabels struct {
	key string
}

func (s PlanMapPathSelectorMetadataLabels) FieldPath() *Plan_FieldSubPath {
	return &Plan_FieldSubPath{
		selector: Plan_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Labels().WithKey(s.key).FieldPath(),
	}
}

func (s PlanMapPathSelectorMetadataLabels) WithValue(value string) *Plan_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Plan_FieldSubPathValue)
}

func (s PlanMapPathSelectorMetadataLabels) WithArrayOfValues(values []string) *Plan_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Plan_FieldSubPathArrayOfValues)
}

type PlanPathSelectorMetadataAnnotations struct{}

func (PlanPathSelectorMetadataAnnotations) FieldPath() *Plan_FieldSubPath {
	return &Plan_FieldSubPath{
		selector: Plan_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Annotations().FieldPath(),
	}
}

func (s PlanPathSelectorMetadataAnnotations) WithValue(value map[string]string) *Plan_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Plan_FieldSubPathValue)
}

func (s PlanPathSelectorMetadataAnnotations) WithArrayOfValues(values []map[string]string) *Plan_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Plan_FieldSubPathArrayOfValues)
}

func (PlanPathSelectorMetadataAnnotations) WithKey(key string) PlanMapPathSelectorMetadataAnnotations {
	return PlanMapPathSelectorMetadataAnnotations{key: key}
}

type PlanMapPathSelectorMetadataAnnotations struct {
	key string
}

func (s PlanMapPathSelectorMetadataAnnotations) FieldPath() *Plan_FieldSubPath {
	return &Plan_FieldSubPath{
		selector: Plan_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Annotations().WithKey(s.key).FieldPath(),
	}
}

func (s PlanMapPathSelectorMetadataAnnotations) WithValue(value string) *Plan_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Plan_FieldSubPathValue)
}

func (s PlanMapPathSelectorMetadataAnnotations) WithArrayOfValues(values []string) *Plan_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Plan_FieldSubPathArrayOfValues)
}

type PlanPathSelectorMetadataGeneration struct{}

func (PlanPathSelectorMetadataGeneration) FieldPath() *Plan_FieldSubPath {
	return &Plan_FieldSubPath{
		selector: Plan_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Generation().FieldPath(),
	}
}

func (s PlanPathSelectorMetadataGeneration) WithValue(value int64) *Plan_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Plan_FieldSubPathValue)
}

func (s PlanPathSelectorMetadataGeneration) WithArrayOfValues(values []int64) *Plan_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Plan_FieldSubPathArrayOfValues)
}

type PlanPathSelectorMetadataResourceVersion struct{}

func (PlanPathSelectorMetadataResourceVersion) FieldPath() *Plan_FieldSubPath {
	return &Plan_FieldSubPath{
		selector: Plan_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().ResourceVersion().FieldPath(),
	}
}

func (s PlanPathSelectorMetadataResourceVersion) WithValue(value string) *Plan_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Plan_FieldSubPathValue)
}

func (s PlanPathSelectorMetadataResourceVersion) WithArrayOfValues(values []string) *Plan_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Plan_FieldSubPathArrayOfValues)
}

type PlanPathSelectorMetadataOwnerReferences struct{}

func (PlanPathSelectorMetadataOwnerReferences) FieldPath() *Plan_FieldSubPath {
	return &Plan_FieldSubPath{
		selector: Plan_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().FieldPath(),
	}
}

func (s PlanPathSelectorMetadataOwnerReferences) WithValue(value []*ntt_meta.OwnerReference) *Plan_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Plan_FieldSubPathValue)
}

func (s PlanPathSelectorMetadataOwnerReferences) WithArrayOfValues(values [][]*ntt_meta.OwnerReference) *Plan_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Plan_FieldSubPathArrayOfValues)
}

func (s PlanPathSelectorMetadataOwnerReferences) WithItemValue(value *ntt_meta.OwnerReference) *Plan_FieldSubPathArrayItemValue {
	return s.FieldPath().WithIArrayItemValue(value).(*Plan_FieldSubPathArrayItemValue)
}

func (PlanPathSelectorMetadataOwnerReferences) Kind() PlanPathSelectorMetadataOwnerReferencesKind {
	return PlanPathSelectorMetadataOwnerReferencesKind{}
}

func (PlanPathSelectorMetadataOwnerReferences) Version() PlanPathSelectorMetadataOwnerReferencesVersion {
	return PlanPathSelectorMetadataOwnerReferencesVersion{}
}

func (PlanPathSelectorMetadataOwnerReferences) Name() PlanPathSelectorMetadataOwnerReferencesName {
	return PlanPathSelectorMetadataOwnerReferencesName{}
}

func (PlanPathSelectorMetadataOwnerReferences) Region() PlanPathSelectorMetadataOwnerReferencesRegion {
	return PlanPathSelectorMetadataOwnerReferencesRegion{}
}

func (PlanPathSelectorMetadataOwnerReferences) Controller() PlanPathSelectorMetadataOwnerReferencesController {
	return PlanPathSelectorMetadataOwnerReferencesController{}
}

func (PlanPathSelectorMetadataOwnerReferences) BlockOwnerDeletion() PlanPathSelectorMetadataOwnerReferencesBlockOwnerDeletion {
	return PlanPathSelectorMetadataOwnerReferencesBlockOwnerDeletion{}
}

func (PlanPathSelectorMetadataOwnerReferences) RequiresOwnerReference() PlanPathSelectorMetadataOwnerReferencesRequiresOwnerReference {
	return PlanPathSelectorMetadataOwnerReferencesRequiresOwnerReference{}
}

type PlanPathSelectorMetadataOwnerReferencesKind struct{}

func (PlanPathSelectorMetadataOwnerReferencesKind) FieldPath() *Plan_FieldSubPath {
	return &Plan_FieldSubPath{
		selector: Plan_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().Kind().FieldPath(),
	}
}

func (s PlanPathSelectorMetadataOwnerReferencesKind) WithValue(value string) *Plan_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Plan_FieldSubPathValue)
}

func (s PlanPathSelectorMetadataOwnerReferencesKind) WithArrayOfValues(values []string) *Plan_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Plan_FieldSubPathArrayOfValues)
}

type PlanPathSelectorMetadataOwnerReferencesVersion struct{}

func (PlanPathSelectorMetadataOwnerReferencesVersion) FieldPath() *Plan_FieldSubPath {
	return &Plan_FieldSubPath{
		selector: Plan_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().Version().FieldPath(),
	}
}

func (s PlanPathSelectorMetadataOwnerReferencesVersion) WithValue(value string) *Plan_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Plan_FieldSubPathValue)
}

func (s PlanPathSelectorMetadataOwnerReferencesVersion) WithArrayOfValues(values []string) *Plan_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Plan_FieldSubPathArrayOfValues)
}

type PlanPathSelectorMetadataOwnerReferencesName struct{}

func (PlanPathSelectorMetadataOwnerReferencesName) FieldPath() *Plan_FieldSubPath {
	return &Plan_FieldSubPath{
		selector: Plan_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().Name().FieldPath(),
	}
}

func (s PlanPathSelectorMetadataOwnerReferencesName) WithValue(value string) *Plan_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Plan_FieldSubPathValue)
}

func (s PlanPathSelectorMetadataOwnerReferencesName) WithArrayOfValues(values []string) *Plan_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Plan_FieldSubPathArrayOfValues)
}

type PlanPathSelectorMetadataOwnerReferencesRegion struct{}

func (PlanPathSelectorMetadataOwnerReferencesRegion) FieldPath() *Plan_FieldSubPath {
	return &Plan_FieldSubPath{
		selector: Plan_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().Region().FieldPath(),
	}
}

func (s PlanPathSelectorMetadataOwnerReferencesRegion) WithValue(value string) *Plan_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Plan_FieldSubPathValue)
}

func (s PlanPathSelectorMetadataOwnerReferencesRegion) WithArrayOfValues(values []string) *Plan_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Plan_FieldSubPathArrayOfValues)
}

type PlanPathSelectorMetadataOwnerReferencesController struct{}

func (PlanPathSelectorMetadataOwnerReferencesController) FieldPath() *Plan_FieldSubPath {
	return &Plan_FieldSubPath{
		selector: Plan_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().Controller().FieldPath(),
	}
}

func (s PlanPathSelectorMetadataOwnerReferencesController) WithValue(value bool) *Plan_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Plan_FieldSubPathValue)
}

func (s PlanPathSelectorMetadataOwnerReferencesController) WithArrayOfValues(values []bool) *Plan_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Plan_FieldSubPathArrayOfValues)
}

type PlanPathSelectorMetadataOwnerReferencesBlockOwnerDeletion struct{}

func (PlanPathSelectorMetadataOwnerReferencesBlockOwnerDeletion) FieldPath() *Plan_FieldSubPath {
	return &Plan_FieldSubPath{
		selector: Plan_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().BlockOwnerDeletion().FieldPath(),
	}
}

func (s PlanPathSelectorMetadataOwnerReferencesBlockOwnerDeletion) WithValue(value bool) *Plan_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Plan_FieldSubPathValue)
}

func (s PlanPathSelectorMetadataOwnerReferencesBlockOwnerDeletion) WithArrayOfValues(values []bool) *Plan_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Plan_FieldSubPathArrayOfValues)
}

type PlanPathSelectorMetadataOwnerReferencesRequiresOwnerReference struct{}

func (PlanPathSelectorMetadataOwnerReferencesRequiresOwnerReference) FieldPath() *Plan_FieldSubPath {
	return &Plan_FieldSubPath{
		selector: Plan_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().RequiresOwnerReference().FieldPath(),
	}
}

func (s PlanPathSelectorMetadataOwnerReferencesRequiresOwnerReference) WithValue(value bool) *Plan_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Plan_FieldSubPathValue)
}

func (s PlanPathSelectorMetadataOwnerReferencesRequiresOwnerReference) WithArrayOfValues(values []bool) *Plan_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Plan_FieldSubPathArrayOfValues)
}

type PlanPathSelectorMetadataShards struct{}

func (PlanPathSelectorMetadataShards) FieldPath() *Plan_FieldSubPath {
	return &Plan_FieldSubPath{
		selector: Plan_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Shards().FieldPath(),
	}
}

func (s PlanPathSelectorMetadataShards) WithValue(value map[string]int64) *Plan_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Plan_FieldSubPathValue)
}

func (s PlanPathSelectorMetadataShards) WithArrayOfValues(values []map[string]int64) *Plan_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Plan_FieldSubPathArrayOfValues)
}

func (PlanPathSelectorMetadataShards) WithKey(key string) PlanMapPathSelectorMetadataShards {
	return PlanMapPathSelectorMetadataShards{key: key}
}

type PlanMapPathSelectorMetadataShards struct {
	key string
}

func (s PlanMapPathSelectorMetadataShards) FieldPath() *Plan_FieldSubPath {
	return &Plan_FieldSubPath{
		selector: Plan_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Shards().WithKey(s.key).FieldPath(),
	}
}

func (s PlanMapPathSelectorMetadataShards) WithValue(value int64) *Plan_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Plan_FieldSubPathValue)
}

func (s PlanMapPathSelectorMetadataShards) WithArrayOfValues(values []int64) *Plan_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Plan_FieldSubPathArrayOfValues)
}

type PlanPathSelectorMetadataSyncing struct{}

func (PlanPathSelectorMetadataSyncing) FieldPath() *Plan_FieldSubPath {
	return &Plan_FieldSubPath{
		selector: Plan_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Syncing().FieldPath(),
	}
}

func (s PlanPathSelectorMetadataSyncing) WithValue(value *ntt_meta.SyncingMeta) *Plan_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Plan_FieldSubPathValue)
}

func (s PlanPathSelectorMetadataSyncing) WithArrayOfValues(values []*ntt_meta.SyncingMeta) *Plan_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Plan_FieldSubPathArrayOfValues)
}

func (PlanPathSelectorMetadataSyncing) OwningRegion() PlanPathSelectorMetadataSyncingOwningRegion {
	return PlanPathSelectorMetadataSyncingOwningRegion{}
}

func (PlanPathSelectorMetadataSyncing) Regions() PlanPathSelectorMetadataSyncingRegions {
	return PlanPathSelectorMetadataSyncingRegions{}
}

type PlanPathSelectorMetadataSyncingOwningRegion struct{}

func (PlanPathSelectorMetadataSyncingOwningRegion) FieldPath() *Plan_FieldSubPath {
	return &Plan_FieldSubPath{
		selector: Plan_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Syncing().OwningRegion().FieldPath(),
	}
}

func (s PlanPathSelectorMetadataSyncingOwningRegion) WithValue(value string) *Plan_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Plan_FieldSubPathValue)
}

func (s PlanPathSelectorMetadataSyncingOwningRegion) WithArrayOfValues(values []string) *Plan_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Plan_FieldSubPathArrayOfValues)
}

type PlanPathSelectorMetadataSyncingRegions struct{}

func (PlanPathSelectorMetadataSyncingRegions) FieldPath() *Plan_FieldSubPath {
	return &Plan_FieldSubPath{
		selector: Plan_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Syncing().Regions().FieldPath(),
	}
}

func (s PlanPathSelectorMetadataSyncingRegions) WithValue(value []string) *Plan_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Plan_FieldSubPathValue)
}

func (s PlanPathSelectorMetadataSyncingRegions) WithArrayOfValues(values [][]string) *Plan_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Plan_FieldSubPathArrayOfValues)
}

func (s PlanPathSelectorMetadataSyncingRegions) WithItemValue(value string) *Plan_FieldSubPathArrayItemValue {
	return s.FieldPath().WithIArrayItemValue(value).(*Plan_FieldSubPathArrayItemValue)
}

type PlanPathSelectorMetadataLifecycle struct{}

func (PlanPathSelectorMetadataLifecycle) FieldPath() *Plan_FieldSubPath {
	return &Plan_FieldSubPath{
		selector: Plan_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Lifecycle().FieldPath(),
	}
}

func (s PlanPathSelectorMetadataLifecycle) WithValue(value *ntt_meta.Lifecycle) *Plan_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Plan_FieldSubPathValue)
}

func (s PlanPathSelectorMetadataLifecycle) WithArrayOfValues(values []*ntt_meta.Lifecycle) *Plan_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Plan_FieldSubPathArrayOfValues)
}

func (PlanPathSelectorMetadataLifecycle) State() PlanPathSelectorMetadataLifecycleState {
	return PlanPathSelectorMetadataLifecycleState{}
}

func (PlanPathSelectorMetadataLifecycle) BlockDeletion() PlanPathSelectorMetadataLifecycleBlockDeletion {
	return PlanPathSelectorMetadataLifecycleBlockDeletion{}
}

type PlanPathSelectorMetadataLifecycleState struct{}

func (PlanPathSelectorMetadataLifecycleState) FieldPath() *Plan_FieldSubPath {
	return &Plan_FieldSubPath{
		selector: Plan_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Lifecycle().State().FieldPath(),
	}
}

func (s PlanPathSelectorMetadataLifecycleState) WithValue(value ntt_meta.Lifecycle_State) *Plan_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Plan_FieldSubPathValue)
}

func (s PlanPathSelectorMetadataLifecycleState) WithArrayOfValues(values []ntt_meta.Lifecycle_State) *Plan_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Plan_FieldSubPathArrayOfValues)
}

type PlanPathSelectorMetadataLifecycleBlockDeletion struct{}

func (PlanPathSelectorMetadataLifecycleBlockDeletion) FieldPath() *Plan_FieldSubPath {
	return &Plan_FieldSubPath{
		selector: Plan_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Lifecycle().BlockDeletion().FieldPath(),
	}
}

func (s PlanPathSelectorMetadataLifecycleBlockDeletion) WithValue(value bool) *Plan_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Plan_FieldSubPathValue)
}

func (s PlanPathSelectorMetadataLifecycleBlockDeletion) WithArrayOfValues(values []bool) *Plan_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Plan_FieldSubPathArrayOfValues)
}
