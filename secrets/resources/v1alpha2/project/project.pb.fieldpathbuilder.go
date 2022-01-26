// Code generated by protoc-gen-goten-object
// File: edgelq/secrets/proto/v1alpha2/project.proto
// DO NOT EDIT!!!

package project

// proto imports
import (
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
	policy "github.com/cloudwan/edgelq-sdk/meta/multi_region/proto/policy"
	syncing_meta "github.com/cloudwan/edgelq-sdk/meta/multi_region/proto/syncing_meta"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
)

// make sure we're using proto imports
var (
	_ = &ntt_meta.Meta{}
	_ = &policy.Policy{}
	_ = &syncing_meta.SyncingMeta{}
	_ = &timestamp.Timestamp{}
)

type ProjectFieldPathBuilder struct{}

func NewProjectFieldPathBuilder() ProjectFieldPathBuilder {
	return ProjectFieldPathBuilder{}
}
func (ProjectFieldPathBuilder) Name() ProjectPathSelectorName {
	return ProjectPathSelectorName{}
}
func (ProjectFieldPathBuilder) Metadata() ProjectPathSelectorMetadata {
	return ProjectPathSelectorMetadata{}
}
func (ProjectFieldPathBuilder) MultiRegionPolicy() ProjectPathSelectorMultiRegionPolicy {
	return ProjectPathSelectorMultiRegionPolicy{}
}

type ProjectPathSelectorName struct{}

func (ProjectPathSelectorName) FieldPath() *Project_FieldTerminalPath {
	return &Project_FieldTerminalPath{selector: Project_FieldPathSelectorName}
}

func (s ProjectPathSelectorName) WithValue(value *Name) *Project_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*Project_FieldTerminalPathValue)
}

func (s ProjectPathSelectorName) WithArrayOfValues(values []*Name) *Project_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Project_FieldTerminalPathArrayOfValues)
}

type ProjectPathSelectorMetadata struct{}

func (ProjectPathSelectorMetadata) FieldPath() *Project_FieldTerminalPath {
	return &Project_FieldTerminalPath{selector: Project_FieldPathSelectorMetadata}
}

func (s ProjectPathSelectorMetadata) WithValue(value *ntt_meta.Meta) *Project_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*Project_FieldTerminalPathValue)
}

func (s ProjectPathSelectorMetadata) WithArrayOfValues(values []*ntt_meta.Meta) *Project_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Project_FieldTerminalPathArrayOfValues)
}

func (ProjectPathSelectorMetadata) WithSubPath(subPath ntt_meta.Meta_FieldPath) *Project_FieldSubPath {
	return &Project_FieldSubPath{selector: Project_FieldPathSelectorMetadata, subPath: subPath}
}

func (s ProjectPathSelectorMetadata) WithSubValue(subPathValue ntt_meta.Meta_FieldPathValue) *Project_FieldSubPathValue {
	return &Project_FieldSubPathValue{Project_FieldPath: s.WithSubPath(subPathValue), subPathValue: subPathValue}
}

func (s ProjectPathSelectorMetadata) WithSubArrayOfValues(subPathArrayOfValues ntt_meta.Meta_FieldPathArrayOfValues) *Project_FieldSubPathArrayOfValues {
	return &Project_FieldSubPathArrayOfValues{Project_FieldPath: s.WithSubPath(subPathArrayOfValues), subPathArrayOfValues: subPathArrayOfValues}
}

func (s ProjectPathSelectorMetadata) WithSubArrayItemValue(subPathArrayItemValue ntt_meta.Meta_FieldPathArrayItemValue) *Project_FieldSubPathArrayItemValue {
	return &Project_FieldSubPathArrayItemValue{Project_FieldPath: s.WithSubPath(subPathArrayItemValue), subPathItemValue: subPathArrayItemValue}
}

func (ProjectPathSelectorMetadata) CreateTime() ProjectPathSelectorMetadataCreateTime {
	return ProjectPathSelectorMetadataCreateTime{}
}

func (ProjectPathSelectorMetadata) UpdateTime() ProjectPathSelectorMetadataUpdateTime {
	return ProjectPathSelectorMetadataUpdateTime{}
}

func (ProjectPathSelectorMetadata) Uuid() ProjectPathSelectorMetadataUuid {
	return ProjectPathSelectorMetadataUuid{}
}

func (ProjectPathSelectorMetadata) Tags() ProjectPathSelectorMetadataTags {
	return ProjectPathSelectorMetadataTags{}
}

func (ProjectPathSelectorMetadata) Labels() ProjectPathSelectorMetadataLabels {
	return ProjectPathSelectorMetadataLabels{}
}

func (ProjectPathSelectorMetadata) Annotations() ProjectPathSelectorMetadataAnnotations {
	return ProjectPathSelectorMetadataAnnotations{}
}

func (ProjectPathSelectorMetadata) Generation() ProjectPathSelectorMetadataGeneration {
	return ProjectPathSelectorMetadataGeneration{}
}

func (ProjectPathSelectorMetadata) ResourceVersion() ProjectPathSelectorMetadataResourceVersion {
	return ProjectPathSelectorMetadataResourceVersion{}
}

func (ProjectPathSelectorMetadata) OwnerReferences() ProjectPathSelectorMetadataOwnerReferences {
	return ProjectPathSelectorMetadataOwnerReferences{}
}

func (ProjectPathSelectorMetadata) Shards() ProjectPathSelectorMetadataShards {
	return ProjectPathSelectorMetadataShards{}
}

func (ProjectPathSelectorMetadata) Syncing() ProjectPathSelectorMetadataSyncing {
	return ProjectPathSelectorMetadataSyncing{}
}

type ProjectPathSelectorMetadataCreateTime struct{}

func (ProjectPathSelectorMetadataCreateTime) FieldPath() *Project_FieldSubPath {
	return &Project_FieldSubPath{
		selector: Project_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().CreateTime().FieldPath(),
	}
}

func (s ProjectPathSelectorMetadataCreateTime) WithValue(value *timestamp.Timestamp) *Project_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Project_FieldSubPathValue)
}

func (s ProjectPathSelectorMetadataCreateTime) WithArrayOfValues(values []*timestamp.Timestamp) *Project_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Project_FieldSubPathArrayOfValues)
}

type ProjectPathSelectorMetadataUpdateTime struct{}

func (ProjectPathSelectorMetadataUpdateTime) FieldPath() *Project_FieldSubPath {
	return &Project_FieldSubPath{
		selector: Project_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().UpdateTime().FieldPath(),
	}
}

func (s ProjectPathSelectorMetadataUpdateTime) WithValue(value *timestamp.Timestamp) *Project_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Project_FieldSubPathValue)
}

func (s ProjectPathSelectorMetadataUpdateTime) WithArrayOfValues(values []*timestamp.Timestamp) *Project_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Project_FieldSubPathArrayOfValues)
}

type ProjectPathSelectorMetadataUuid struct{}

func (ProjectPathSelectorMetadataUuid) FieldPath() *Project_FieldSubPath {
	return &Project_FieldSubPath{
		selector: Project_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Uuid().FieldPath(),
	}
}

func (s ProjectPathSelectorMetadataUuid) WithValue(value string) *Project_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Project_FieldSubPathValue)
}

func (s ProjectPathSelectorMetadataUuid) WithArrayOfValues(values []string) *Project_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Project_FieldSubPathArrayOfValues)
}

type ProjectPathSelectorMetadataTags struct{}

func (ProjectPathSelectorMetadataTags) FieldPath() *Project_FieldSubPath {
	return &Project_FieldSubPath{
		selector: Project_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Tags().FieldPath(),
	}
}

func (s ProjectPathSelectorMetadataTags) WithValue(value []string) *Project_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Project_FieldSubPathValue)
}

func (s ProjectPathSelectorMetadataTags) WithArrayOfValues(values [][]string) *Project_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Project_FieldSubPathArrayOfValues)
}

func (s ProjectPathSelectorMetadataTags) WithItemValue(value string) *Project_FieldSubPathArrayItemValue {
	return s.FieldPath().WithIArrayItemValue(value).(*Project_FieldSubPathArrayItemValue)
}

type ProjectPathSelectorMetadataLabels struct{}

func (ProjectPathSelectorMetadataLabels) FieldPath() *Project_FieldSubPath {
	return &Project_FieldSubPath{
		selector: Project_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Labels().FieldPath(),
	}
}

func (s ProjectPathSelectorMetadataLabels) WithValue(value map[string]string) *Project_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Project_FieldSubPathValue)
}

func (s ProjectPathSelectorMetadataLabels) WithArrayOfValues(values []map[string]string) *Project_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Project_FieldSubPathArrayOfValues)
}

func (ProjectPathSelectorMetadataLabels) WithKey(key string) ProjectMapPathSelectorMetadataLabels {
	return ProjectMapPathSelectorMetadataLabels{key: key}
}

type ProjectMapPathSelectorMetadataLabels struct {
	key string
}

func (s ProjectMapPathSelectorMetadataLabels) FieldPath() *Project_FieldSubPath {
	return &Project_FieldSubPath{
		selector: Project_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Labels().WithKey(s.key).FieldPath(),
	}
}

func (s ProjectMapPathSelectorMetadataLabels) WithValue(value string) *Project_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Project_FieldSubPathValue)
}

func (s ProjectMapPathSelectorMetadataLabels) WithArrayOfValues(values []string) *Project_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Project_FieldSubPathArrayOfValues)
}

type ProjectPathSelectorMetadataAnnotations struct{}

func (ProjectPathSelectorMetadataAnnotations) FieldPath() *Project_FieldSubPath {
	return &Project_FieldSubPath{
		selector: Project_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Annotations().FieldPath(),
	}
}

func (s ProjectPathSelectorMetadataAnnotations) WithValue(value map[string]string) *Project_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Project_FieldSubPathValue)
}

func (s ProjectPathSelectorMetadataAnnotations) WithArrayOfValues(values []map[string]string) *Project_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Project_FieldSubPathArrayOfValues)
}

func (ProjectPathSelectorMetadataAnnotations) WithKey(key string) ProjectMapPathSelectorMetadataAnnotations {
	return ProjectMapPathSelectorMetadataAnnotations{key: key}
}

type ProjectMapPathSelectorMetadataAnnotations struct {
	key string
}

func (s ProjectMapPathSelectorMetadataAnnotations) FieldPath() *Project_FieldSubPath {
	return &Project_FieldSubPath{
		selector: Project_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Annotations().WithKey(s.key).FieldPath(),
	}
}

func (s ProjectMapPathSelectorMetadataAnnotations) WithValue(value string) *Project_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Project_FieldSubPathValue)
}

func (s ProjectMapPathSelectorMetadataAnnotations) WithArrayOfValues(values []string) *Project_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Project_FieldSubPathArrayOfValues)
}

type ProjectPathSelectorMetadataGeneration struct{}

func (ProjectPathSelectorMetadataGeneration) FieldPath() *Project_FieldSubPath {
	return &Project_FieldSubPath{
		selector: Project_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Generation().FieldPath(),
	}
}

func (s ProjectPathSelectorMetadataGeneration) WithValue(value int64) *Project_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Project_FieldSubPathValue)
}

func (s ProjectPathSelectorMetadataGeneration) WithArrayOfValues(values []int64) *Project_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Project_FieldSubPathArrayOfValues)
}

type ProjectPathSelectorMetadataResourceVersion struct{}

func (ProjectPathSelectorMetadataResourceVersion) FieldPath() *Project_FieldSubPath {
	return &Project_FieldSubPath{
		selector: Project_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().ResourceVersion().FieldPath(),
	}
}

func (s ProjectPathSelectorMetadataResourceVersion) WithValue(value string) *Project_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Project_FieldSubPathValue)
}

func (s ProjectPathSelectorMetadataResourceVersion) WithArrayOfValues(values []string) *Project_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Project_FieldSubPathArrayOfValues)
}

type ProjectPathSelectorMetadataOwnerReferences struct{}

func (ProjectPathSelectorMetadataOwnerReferences) FieldPath() *Project_FieldSubPath {
	return &Project_FieldSubPath{
		selector: Project_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().FieldPath(),
	}
}

func (s ProjectPathSelectorMetadataOwnerReferences) WithValue(value []*ntt_meta.OwnerReference) *Project_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Project_FieldSubPathValue)
}

func (s ProjectPathSelectorMetadataOwnerReferences) WithArrayOfValues(values [][]*ntt_meta.OwnerReference) *Project_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Project_FieldSubPathArrayOfValues)
}

func (s ProjectPathSelectorMetadataOwnerReferences) WithItemValue(value *ntt_meta.OwnerReference) *Project_FieldSubPathArrayItemValue {
	return s.FieldPath().WithIArrayItemValue(value).(*Project_FieldSubPathArrayItemValue)
}

func (ProjectPathSelectorMetadataOwnerReferences) ApiVersion() ProjectPathSelectorMetadataOwnerReferencesApiVersion {
	return ProjectPathSelectorMetadataOwnerReferencesApiVersion{}
}

func (ProjectPathSelectorMetadataOwnerReferences) Kind() ProjectPathSelectorMetadataOwnerReferencesKind {
	return ProjectPathSelectorMetadataOwnerReferencesKind{}
}

func (ProjectPathSelectorMetadataOwnerReferences) Name() ProjectPathSelectorMetadataOwnerReferencesName {
	return ProjectPathSelectorMetadataOwnerReferencesName{}
}

func (ProjectPathSelectorMetadataOwnerReferences) Uid() ProjectPathSelectorMetadataOwnerReferencesUid {
	return ProjectPathSelectorMetadataOwnerReferencesUid{}
}

func (ProjectPathSelectorMetadataOwnerReferences) Controller() ProjectPathSelectorMetadataOwnerReferencesController {
	return ProjectPathSelectorMetadataOwnerReferencesController{}
}

func (ProjectPathSelectorMetadataOwnerReferences) BlockOwnerDeletion() ProjectPathSelectorMetadataOwnerReferencesBlockOwnerDeletion {
	return ProjectPathSelectorMetadataOwnerReferencesBlockOwnerDeletion{}
}

type ProjectPathSelectorMetadataOwnerReferencesApiVersion struct{}

func (ProjectPathSelectorMetadataOwnerReferencesApiVersion) FieldPath() *Project_FieldSubPath {
	return &Project_FieldSubPath{
		selector: Project_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().ApiVersion().FieldPath(),
	}
}

func (s ProjectPathSelectorMetadataOwnerReferencesApiVersion) WithValue(value string) *Project_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Project_FieldSubPathValue)
}

func (s ProjectPathSelectorMetadataOwnerReferencesApiVersion) WithArrayOfValues(values []string) *Project_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Project_FieldSubPathArrayOfValues)
}

type ProjectPathSelectorMetadataOwnerReferencesKind struct{}

func (ProjectPathSelectorMetadataOwnerReferencesKind) FieldPath() *Project_FieldSubPath {
	return &Project_FieldSubPath{
		selector: Project_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().Kind().FieldPath(),
	}
}

func (s ProjectPathSelectorMetadataOwnerReferencesKind) WithValue(value string) *Project_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Project_FieldSubPathValue)
}

func (s ProjectPathSelectorMetadataOwnerReferencesKind) WithArrayOfValues(values []string) *Project_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Project_FieldSubPathArrayOfValues)
}

type ProjectPathSelectorMetadataOwnerReferencesName struct{}

func (ProjectPathSelectorMetadataOwnerReferencesName) FieldPath() *Project_FieldSubPath {
	return &Project_FieldSubPath{
		selector: Project_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().Name().FieldPath(),
	}
}

func (s ProjectPathSelectorMetadataOwnerReferencesName) WithValue(value string) *Project_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Project_FieldSubPathValue)
}

func (s ProjectPathSelectorMetadataOwnerReferencesName) WithArrayOfValues(values []string) *Project_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Project_FieldSubPathArrayOfValues)
}

type ProjectPathSelectorMetadataOwnerReferencesUid struct{}

func (ProjectPathSelectorMetadataOwnerReferencesUid) FieldPath() *Project_FieldSubPath {
	return &Project_FieldSubPath{
		selector: Project_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().Uid().FieldPath(),
	}
}

func (s ProjectPathSelectorMetadataOwnerReferencesUid) WithValue(value string) *Project_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Project_FieldSubPathValue)
}

func (s ProjectPathSelectorMetadataOwnerReferencesUid) WithArrayOfValues(values []string) *Project_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Project_FieldSubPathArrayOfValues)
}

type ProjectPathSelectorMetadataOwnerReferencesController struct{}

func (ProjectPathSelectorMetadataOwnerReferencesController) FieldPath() *Project_FieldSubPath {
	return &Project_FieldSubPath{
		selector: Project_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().Controller().FieldPath(),
	}
}

func (s ProjectPathSelectorMetadataOwnerReferencesController) WithValue(value bool) *Project_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Project_FieldSubPathValue)
}

func (s ProjectPathSelectorMetadataOwnerReferencesController) WithArrayOfValues(values []bool) *Project_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Project_FieldSubPathArrayOfValues)
}

type ProjectPathSelectorMetadataOwnerReferencesBlockOwnerDeletion struct{}

func (ProjectPathSelectorMetadataOwnerReferencesBlockOwnerDeletion) FieldPath() *Project_FieldSubPath {
	return &Project_FieldSubPath{
		selector: Project_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().BlockOwnerDeletion().FieldPath(),
	}
}

func (s ProjectPathSelectorMetadataOwnerReferencesBlockOwnerDeletion) WithValue(value bool) *Project_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Project_FieldSubPathValue)
}

func (s ProjectPathSelectorMetadataOwnerReferencesBlockOwnerDeletion) WithArrayOfValues(values []bool) *Project_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Project_FieldSubPathArrayOfValues)
}

type ProjectPathSelectorMetadataShards struct{}

func (ProjectPathSelectorMetadataShards) FieldPath() *Project_FieldSubPath {
	return &Project_FieldSubPath{
		selector: Project_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Shards().FieldPath(),
	}
}

func (s ProjectPathSelectorMetadataShards) WithValue(value map[string]int64) *Project_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Project_FieldSubPathValue)
}

func (s ProjectPathSelectorMetadataShards) WithArrayOfValues(values []map[string]int64) *Project_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Project_FieldSubPathArrayOfValues)
}

func (ProjectPathSelectorMetadataShards) WithKey(key string) ProjectMapPathSelectorMetadataShards {
	return ProjectMapPathSelectorMetadataShards{key: key}
}

type ProjectMapPathSelectorMetadataShards struct {
	key string
}

func (s ProjectMapPathSelectorMetadataShards) FieldPath() *Project_FieldSubPath {
	return &Project_FieldSubPath{
		selector: Project_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Shards().WithKey(s.key).FieldPath(),
	}
}

func (s ProjectMapPathSelectorMetadataShards) WithValue(value int64) *Project_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Project_FieldSubPathValue)
}

func (s ProjectMapPathSelectorMetadataShards) WithArrayOfValues(values []int64) *Project_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Project_FieldSubPathArrayOfValues)
}

type ProjectPathSelectorMetadataSyncing struct{}

func (ProjectPathSelectorMetadataSyncing) FieldPath() *Project_FieldSubPath {
	return &Project_FieldSubPath{
		selector: Project_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Syncing().FieldPath(),
	}
}

func (s ProjectPathSelectorMetadataSyncing) WithValue(value *syncing_meta.SyncingMeta) *Project_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Project_FieldSubPathValue)
}

func (s ProjectPathSelectorMetadataSyncing) WithArrayOfValues(values []*syncing_meta.SyncingMeta) *Project_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Project_FieldSubPathArrayOfValues)
}

func (ProjectPathSelectorMetadataSyncing) OwningRegion() ProjectPathSelectorMetadataSyncingOwningRegion {
	return ProjectPathSelectorMetadataSyncingOwningRegion{}
}

func (ProjectPathSelectorMetadataSyncing) Regions() ProjectPathSelectorMetadataSyncingRegions {
	return ProjectPathSelectorMetadataSyncingRegions{}
}

type ProjectPathSelectorMetadataSyncingOwningRegion struct{}

func (ProjectPathSelectorMetadataSyncingOwningRegion) FieldPath() *Project_FieldSubPath {
	return &Project_FieldSubPath{
		selector: Project_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Syncing().OwningRegion().FieldPath(),
	}
}

func (s ProjectPathSelectorMetadataSyncingOwningRegion) WithValue(value string) *Project_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Project_FieldSubPathValue)
}

func (s ProjectPathSelectorMetadataSyncingOwningRegion) WithArrayOfValues(values []string) *Project_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Project_FieldSubPathArrayOfValues)
}

type ProjectPathSelectorMetadataSyncingRegions struct{}

func (ProjectPathSelectorMetadataSyncingRegions) FieldPath() *Project_FieldSubPath {
	return &Project_FieldSubPath{
		selector: Project_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Syncing().Regions().FieldPath(),
	}
}

func (s ProjectPathSelectorMetadataSyncingRegions) WithValue(value []string) *Project_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Project_FieldSubPathValue)
}

func (s ProjectPathSelectorMetadataSyncingRegions) WithArrayOfValues(values [][]string) *Project_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Project_FieldSubPathArrayOfValues)
}

func (s ProjectPathSelectorMetadataSyncingRegions) WithItemValue(value string) *Project_FieldSubPathArrayItemValue {
	return s.FieldPath().WithIArrayItemValue(value).(*Project_FieldSubPathArrayItemValue)
}

type ProjectPathSelectorMultiRegionPolicy struct{}

func (ProjectPathSelectorMultiRegionPolicy) FieldPath() *Project_FieldTerminalPath {
	return &Project_FieldTerminalPath{selector: Project_FieldPathSelectorMultiRegionPolicy}
}

func (s ProjectPathSelectorMultiRegionPolicy) WithValue(value *policy.Policy) *Project_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*Project_FieldTerminalPathValue)
}

func (s ProjectPathSelectorMultiRegionPolicy) WithArrayOfValues(values []*policy.Policy) *Project_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Project_FieldTerminalPathArrayOfValues)
}

func (ProjectPathSelectorMultiRegionPolicy) WithSubPath(subPath policy.Policy_FieldPath) *Project_FieldSubPath {
	return &Project_FieldSubPath{selector: Project_FieldPathSelectorMultiRegionPolicy, subPath: subPath}
}

func (s ProjectPathSelectorMultiRegionPolicy) WithSubValue(subPathValue policy.Policy_FieldPathValue) *Project_FieldSubPathValue {
	return &Project_FieldSubPathValue{Project_FieldPath: s.WithSubPath(subPathValue), subPathValue: subPathValue}
}

func (s ProjectPathSelectorMultiRegionPolicy) WithSubArrayOfValues(subPathArrayOfValues policy.Policy_FieldPathArrayOfValues) *Project_FieldSubPathArrayOfValues {
	return &Project_FieldSubPathArrayOfValues{Project_FieldPath: s.WithSubPath(subPathArrayOfValues), subPathArrayOfValues: subPathArrayOfValues}
}

func (s ProjectPathSelectorMultiRegionPolicy) WithSubArrayItemValue(subPathArrayItemValue policy.Policy_FieldPathArrayItemValue) *Project_FieldSubPathArrayItemValue {
	return &Project_FieldSubPathArrayItemValue{Project_FieldPath: s.WithSubPath(subPathArrayItemValue), subPathItemValue: subPathArrayItemValue}
}

func (ProjectPathSelectorMultiRegionPolicy) EnabledRegions() ProjectPathSelectorMultiRegionPolicyEnabledRegions {
	return ProjectPathSelectorMultiRegionPolicyEnabledRegions{}
}

func (ProjectPathSelectorMultiRegionPolicy) DefaultControlRegion() ProjectPathSelectorMultiRegionPolicyDefaultControlRegion {
	return ProjectPathSelectorMultiRegionPolicyDefaultControlRegion{}
}

func (ProjectPathSelectorMultiRegionPolicy) CriteriaForDisabledSync() ProjectPathSelectorMultiRegionPolicyCriteriaForDisabledSync {
	return ProjectPathSelectorMultiRegionPolicyCriteriaForDisabledSync{}
}

type ProjectPathSelectorMultiRegionPolicyEnabledRegions struct{}

func (ProjectPathSelectorMultiRegionPolicyEnabledRegions) FieldPath() *Project_FieldSubPath {
	return &Project_FieldSubPath{
		selector: Project_FieldPathSelectorMultiRegionPolicy,
		subPath:  policy.NewPolicyFieldPathBuilder().EnabledRegions().FieldPath(),
	}
}

func (s ProjectPathSelectorMultiRegionPolicyEnabledRegions) WithValue(value []string) *Project_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Project_FieldSubPathValue)
}

func (s ProjectPathSelectorMultiRegionPolicyEnabledRegions) WithArrayOfValues(values [][]string) *Project_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Project_FieldSubPathArrayOfValues)
}

func (s ProjectPathSelectorMultiRegionPolicyEnabledRegions) WithItemValue(value string) *Project_FieldSubPathArrayItemValue {
	return s.FieldPath().WithIArrayItemValue(value).(*Project_FieldSubPathArrayItemValue)
}

type ProjectPathSelectorMultiRegionPolicyDefaultControlRegion struct{}

func (ProjectPathSelectorMultiRegionPolicyDefaultControlRegion) FieldPath() *Project_FieldSubPath {
	return &Project_FieldSubPath{
		selector: Project_FieldPathSelectorMultiRegionPolicy,
		subPath:  policy.NewPolicyFieldPathBuilder().DefaultControlRegion().FieldPath(),
	}
}

func (s ProjectPathSelectorMultiRegionPolicyDefaultControlRegion) WithValue(value string) *Project_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Project_FieldSubPathValue)
}

func (s ProjectPathSelectorMultiRegionPolicyDefaultControlRegion) WithArrayOfValues(values []string) *Project_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Project_FieldSubPathArrayOfValues)
}

type ProjectPathSelectorMultiRegionPolicyCriteriaForDisabledSync struct{}

func (ProjectPathSelectorMultiRegionPolicyCriteriaForDisabledSync) FieldPath() *Project_FieldSubPath {
	return &Project_FieldSubPath{
		selector: Project_FieldPathSelectorMultiRegionPolicy,
		subPath:  policy.NewPolicyFieldPathBuilder().CriteriaForDisabledSync().FieldPath(),
	}
}

func (s ProjectPathSelectorMultiRegionPolicyCriteriaForDisabledSync) WithValue(value []*policy.Policy_CriteriaForDisabledSync) *Project_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Project_FieldSubPathValue)
}

func (s ProjectPathSelectorMultiRegionPolicyCriteriaForDisabledSync) WithArrayOfValues(values [][]*policy.Policy_CriteriaForDisabledSync) *Project_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Project_FieldSubPathArrayOfValues)
}

func (s ProjectPathSelectorMultiRegionPolicyCriteriaForDisabledSync) WithItemValue(value *policy.Policy_CriteriaForDisabledSync) *Project_FieldSubPathArrayItemValue {
	return s.FieldPath().WithIArrayItemValue(value).(*Project_FieldSubPathArrayItemValue)
}

func (ProjectPathSelectorMultiRegionPolicyCriteriaForDisabledSync) ResourceTypeName() ProjectPathSelectorMultiRegionPolicyCriteriaForDisabledSyncResourceTypeName {
	return ProjectPathSelectorMultiRegionPolicyCriteriaForDisabledSyncResourceTypeName{}
}

func (ProjectPathSelectorMultiRegionPolicyCriteriaForDisabledSync) SourceRegion() ProjectPathSelectorMultiRegionPolicyCriteriaForDisabledSyncSourceRegion {
	return ProjectPathSelectorMultiRegionPolicyCriteriaForDisabledSyncSourceRegion{}
}

func (ProjectPathSelectorMultiRegionPolicyCriteriaForDisabledSync) DestRegion() ProjectPathSelectorMultiRegionPolicyCriteriaForDisabledSyncDestRegion {
	return ProjectPathSelectorMultiRegionPolicyCriteriaForDisabledSyncDestRegion{}
}

type ProjectPathSelectorMultiRegionPolicyCriteriaForDisabledSyncResourceTypeName struct{}

func (ProjectPathSelectorMultiRegionPolicyCriteriaForDisabledSyncResourceTypeName) FieldPath() *Project_FieldSubPath {
	return &Project_FieldSubPath{
		selector: Project_FieldPathSelectorMultiRegionPolicy,
		subPath:  policy.NewPolicyFieldPathBuilder().CriteriaForDisabledSync().ResourceTypeName().FieldPath(),
	}
}

func (s ProjectPathSelectorMultiRegionPolicyCriteriaForDisabledSyncResourceTypeName) WithValue(value string) *Project_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Project_FieldSubPathValue)
}

func (s ProjectPathSelectorMultiRegionPolicyCriteriaForDisabledSyncResourceTypeName) WithArrayOfValues(values []string) *Project_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Project_FieldSubPathArrayOfValues)
}

type ProjectPathSelectorMultiRegionPolicyCriteriaForDisabledSyncSourceRegion struct{}

func (ProjectPathSelectorMultiRegionPolicyCriteriaForDisabledSyncSourceRegion) FieldPath() *Project_FieldSubPath {
	return &Project_FieldSubPath{
		selector: Project_FieldPathSelectorMultiRegionPolicy,
		subPath:  policy.NewPolicyFieldPathBuilder().CriteriaForDisabledSync().SourceRegion().FieldPath(),
	}
}

func (s ProjectPathSelectorMultiRegionPolicyCriteriaForDisabledSyncSourceRegion) WithValue(value string) *Project_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Project_FieldSubPathValue)
}

func (s ProjectPathSelectorMultiRegionPolicyCriteriaForDisabledSyncSourceRegion) WithArrayOfValues(values []string) *Project_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Project_FieldSubPathArrayOfValues)
}

type ProjectPathSelectorMultiRegionPolicyCriteriaForDisabledSyncDestRegion struct{}

func (ProjectPathSelectorMultiRegionPolicyCriteriaForDisabledSyncDestRegion) FieldPath() *Project_FieldSubPath {
	return &Project_FieldSubPath{
		selector: Project_FieldPathSelectorMultiRegionPolicy,
		subPath:  policy.NewPolicyFieldPathBuilder().CriteriaForDisabledSync().DestRegion().FieldPath(),
	}
}

func (s ProjectPathSelectorMultiRegionPolicyCriteriaForDisabledSyncDestRegion) WithValue(value string) *Project_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Project_FieldSubPathValue)
}

func (s ProjectPathSelectorMultiRegionPolicyCriteriaForDisabledSyncDestRegion) WithArrayOfValues(values []string) *Project_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Project_FieldSubPathArrayOfValues)
}
