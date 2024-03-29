// Code generated by protoc-gen-goten-object
// File: edgelq/monitoring/proto/v3/project.proto
// DO NOT EDIT!!!

package project

// proto imports
import (
	meta "github.com/cloudwan/goten-sdk/types/meta"
	multi_region_policy "github.com/cloudwan/goten-sdk/types/multi_region_policy"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

// make sure we're using proto imports
var (
	_ = &timestamppb.Timestamp{}
	_ = &meta.Meta{}
	_ = &multi_region_policy.MultiRegionPolicy{}
)

type ProjectFieldPathBuilder struct{}

func NewProjectFieldPathBuilder() ProjectFieldPathBuilder {
	return ProjectFieldPathBuilder{}
}
func (ProjectFieldPathBuilder) Name() ProjectPathSelectorName {
	return ProjectPathSelectorName{}
}
func (ProjectFieldPathBuilder) Title() ProjectPathSelectorTitle {
	return ProjectPathSelectorTitle{}
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

type ProjectPathSelectorTitle struct{}

func (ProjectPathSelectorTitle) FieldPath() *Project_FieldTerminalPath {
	return &Project_FieldTerminalPath{selector: Project_FieldPathSelectorTitle}
}

func (s ProjectPathSelectorTitle) WithValue(value string) *Project_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*Project_FieldTerminalPathValue)
}

func (s ProjectPathSelectorTitle) WithArrayOfValues(values []string) *Project_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Project_FieldTerminalPathArrayOfValues)
}

type ProjectPathSelectorMetadata struct{}

func (ProjectPathSelectorMetadata) FieldPath() *Project_FieldTerminalPath {
	return &Project_FieldTerminalPath{selector: Project_FieldPathSelectorMetadata}
}

func (s ProjectPathSelectorMetadata) WithValue(value *meta.Meta) *Project_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*Project_FieldTerminalPathValue)
}

func (s ProjectPathSelectorMetadata) WithArrayOfValues(values []*meta.Meta) *Project_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Project_FieldTerminalPathArrayOfValues)
}

func (ProjectPathSelectorMetadata) WithSubPath(subPath meta.Meta_FieldPath) *Project_FieldSubPath {
	return &Project_FieldSubPath{selector: Project_FieldPathSelectorMetadata, subPath: subPath}
}

func (s ProjectPathSelectorMetadata) WithSubValue(subPathValue meta.Meta_FieldPathValue) *Project_FieldSubPathValue {
	return &Project_FieldSubPathValue{Project_FieldPath: s.WithSubPath(subPathValue), subPathValue: subPathValue}
}

func (s ProjectPathSelectorMetadata) WithSubArrayOfValues(subPathArrayOfValues meta.Meta_FieldPathArrayOfValues) *Project_FieldSubPathArrayOfValues {
	return &Project_FieldSubPathArrayOfValues{Project_FieldPath: s.WithSubPath(subPathArrayOfValues), subPathArrayOfValues: subPathArrayOfValues}
}

func (s ProjectPathSelectorMetadata) WithSubArrayItemValue(subPathArrayItemValue meta.Meta_FieldPathArrayItemValue) *Project_FieldSubPathArrayItemValue {
	return &Project_FieldSubPathArrayItemValue{Project_FieldPath: s.WithSubPath(subPathArrayItemValue), subPathItemValue: subPathArrayItemValue}
}

func (ProjectPathSelectorMetadata) CreateTime() ProjectPathSelectorMetadataCreateTime {
	return ProjectPathSelectorMetadataCreateTime{}
}

func (ProjectPathSelectorMetadata) UpdateTime() ProjectPathSelectorMetadataUpdateTime {
	return ProjectPathSelectorMetadataUpdateTime{}
}

func (ProjectPathSelectorMetadata) DeleteTime() ProjectPathSelectorMetadataDeleteTime {
	return ProjectPathSelectorMetadataDeleteTime{}
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

func (ProjectPathSelectorMetadata) Lifecycle() ProjectPathSelectorMetadataLifecycle {
	return ProjectPathSelectorMetadataLifecycle{}
}

func (ProjectPathSelectorMetadata) Services() ProjectPathSelectorMetadataServices {
	return ProjectPathSelectorMetadataServices{}
}

type ProjectPathSelectorMetadataCreateTime struct{}

func (ProjectPathSelectorMetadataCreateTime) FieldPath() *Project_FieldSubPath {
	return &Project_FieldSubPath{
		selector: Project_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().CreateTime().FieldPath(),
	}
}

func (s ProjectPathSelectorMetadataCreateTime) WithValue(value *timestamppb.Timestamp) *Project_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Project_FieldSubPathValue)
}

func (s ProjectPathSelectorMetadataCreateTime) WithArrayOfValues(values []*timestamppb.Timestamp) *Project_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Project_FieldSubPathArrayOfValues)
}

type ProjectPathSelectorMetadataUpdateTime struct{}

func (ProjectPathSelectorMetadataUpdateTime) FieldPath() *Project_FieldSubPath {
	return &Project_FieldSubPath{
		selector: Project_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().UpdateTime().FieldPath(),
	}
}

func (s ProjectPathSelectorMetadataUpdateTime) WithValue(value *timestamppb.Timestamp) *Project_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Project_FieldSubPathValue)
}

func (s ProjectPathSelectorMetadataUpdateTime) WithArrayOfValues(values []*timestamppb.Timestamp) *Project_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Project_FieldSubPathArrayOfValues)
}

type ProjectPathSelectorMetadataDeleteTime struct{}

func (ProjectPathSelectorMetadataDeleteTime) FieldPath() *Project_FieldSubPath {
	return &Project_FieldSubPath{
		selector: Project_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().DeleteTime().FieldPath(),
	}
}

func (s ProjectPathSelectorMetadataDeleteTime) WithValue(value *timestamppb.Timestamp) *Project_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Project_FieldSubPathValue)
}

func (s ProjectPathSelectorMetadataDeleteTime) WithArrayOfValues(values []*timestamppb.Timestamp) *Project_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Project_FieldSubPathArrayOfValues)
}

type ProjectPathSelectorMetadataUuid struct{}

func (ProjectPathSelectorMetadataUuid) FieldPath() *Project_FieldSubPath {
	return &Project_FieldSubPath{
		selector: Project_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Uuid().FieldPath(),
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
		subPath:  meta.NewMetaFieldPathBuilder().Tags().FieldPath(),
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
		subPath:  meta.NewMetaFieldPathBuilder().Labels().FieldPath(),
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
		subPath:  meta.NewMetaFieldPathBuilder().Labels().WithKey(s.key).FieldPath(),
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
		subPath:  meta.NewMetaFieldPathBuilder().Annotations().FieldPath(),
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
		subPath:  meta.NewMetaFieldPathBuilder().Annotations().WithKey(s.key).FieldPath(),
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
		subPath:  meta.NewMetaFieldPathBuilder().Generation().FieldPath(),
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
		subPath:  meta.NewMetaFieldPathBuilder().ResourceVersion().FieldPath(),
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
		subPath:  meta.NewMetaFieldPathBuilder().OwnerReferences().FieldPath(),
	}
}

func (s ProjectPathSelectorMetadataOwnerReferences) WithValue(value []*meta.OwnerReference) *Project_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Project_FieldSubPathValue)
}

func (s ProjectPathSelectorMetadataOwnerReferences) WithArrayOfValues(values [][]*meta.OwnerReference) *Project_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Project_FieldSubPathArrayOfValues)
}

func (s ProjectPathSelectorMetadataOwnerReferences) WithItemValue(value *meta.OwnerReference) *Project_FieldSubPathArrayItemValue {
	return s.FieldPath().WithIArrayItemValue(value).(*Project_FieldSubPathArrayItemValue)
}

func (ProjectPathSelectorMetadataOwnerReferences) Kind() ProjectPathSelectorMetadataOwnerReferencesKind {
	return ProjectPathSelectorMetadataOwnerReferencesKind{}
}

func (ProjectPathSelectorMetadataOwnerReferences) Version() ProjectPathSelectorMetadataOwnerReferencesVersion {
	return ProjectPathSelectorMetadataOwnerReferencesVersion{}
}

func (ProjectPathSelectorMetadataOwnerReferences) Name() ProjectPathSelectorMetadataOwnerReferencesName {
	return ProjectPathSelectorMetadataOwnerReferencesName{}
}

func (ProjectPathSelectorMetadataOwnerReferences) Region() ProjectPathSelectorMetadataOwnerReferencesRegion {
	return ProjectPathSelectorMetadataOwnerReferencesRegion{}
}

func (ProjectPathSelectorMetadataOwnerReferences) Controller() ProjectPathSelectorMetadataOwnerReferencesController {
	return ProjectPathSelectorMetadataOwnerReferencesController{}
}

func (ProjectPathSelectorMetadataOwnerReferences) RequiresOwnerReference() ProjectPathSelectorMetadataOwnerReferencesRequiresOwnerReference {
	return ProjectPathSelectorMetadataOwnerReferencesRequiresOwnerReference{}
}

type ProjectPathSelectorMetadataOwnerReferencesKind struct{}

func (ProjectPathSelectorMetadataOwnerReferencesKind) FieldPath() *Project_FieldSubPath {
	return &Project_FieldSubPath{
		selector: Project_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().OwnerReferences().Kind().FieldPath(),
	}
}

func (s ProjectPathSelectorMetadataOwnerReferencesKind) WithValue(value string) *Project_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Project_FieldSubPathValue)
}

func (s ProjectPathSelectorMetadataOwnerReferencesKind) WithArrayOfValues(values []string) *Project_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Project_FieldSubPathArrayOfValues)
}

type ProjectPathSelectorMetadataOwnerReferencesVersion struct{}

func (ProjectPathSelectorMetadataOwnerReferencesVersion) FieldPath() *Project_FieldSubPath {
	return &Project_FieldSubPath{
		selector: Project_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().OwnerReferences().Version().FieldPath(),
	}
}

func (s ProjectPathSelectorMetadataOwnerReferencesVersion) WithValue(value string) *Project_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Project_FieldSubPathValue)
}

func (s ProjectPathSelectorMetadataOwnerReferencesVersion) WithArrayOfValues(values []string) *Project_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Project_FieldSubPathArrayOfValues)
}

type ProjectPathSelectorMetadataOwnerReferencesName struct{}

func (ProjectPathSelectorMetadataOwnerReferencesName) FieldPath() *Project_FieldSubPath {
	return &Project_FieldSubPath{
		selector: Project_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().OwnerReferences().Name().FieldPath(),
	}
}

func (s ProjectPathSelectorMetadataOwnerReferencesName) WithValue(value string) *Project_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Project_FieldSubPathValue)
}

func (s ProjectPathSelectorMetadataOwnerReferencesName) WithArrayOfValues(values []string) *Project_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Project_FieldSubPathArrayOfValues)
}

type ProjectPathSelectorMetadataOwnerReferencesRegion struct{}

func (ProjectPathSelectorMetadataOwnerReferencesRegion) FieldPath() *Project_FieldSubPath {
	return &Project_FieldSubPath{
		selector: Project_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().OwnerReferences().Region().FieldPath(),
	}
}

func (s ProjectPathSelectorMetadataOwnerReferencesRegion) WithValue(value string) *Project_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Project_FieldSubPathValue)
}

func (s ProjectPathSelectorMetadataOwnerReferencesRegion) WithArrayOfValues(values []string) *Project_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Project_FieldSubPathArrayOfValues)
}

type ProjectPathSelectorMetadataOwnerReferencesController struct{}

func (ProjectPathSelectorMetadataOwnerReferencesController) FieldPath() *Project_FieldSubPath {
	return &Project_FieldSubPath{
		selector: Project_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().OwnerReferences().Controller().FieldPath(),
	}
}

func (s ProjectPathSelectorMetadataOwnerReferencesController) WithValue(value bool) *Project_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Project_FieldSubPathValue)
}

func (s ProjectPathSelectorMetadataOwnerReferencesController) WithArrayOfValues(values []bool) *Project_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Project_FieldSubPathArrayOfValues)
}

type ProjectPathSelectorMetadataOwnerReferencesRequiresOwnerReference struct{}

func (ProjectPathSelectorMetadataOwnerReferencesRequiresOwnerReference) FieldPath() *Project_FieldSubPath {
	return &Project_FieldSubPath{
		selector: Project_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().OwnerReferences().RequiresOwnerReference().FieldPath(),
	}
}

func (s ProjectPathSelectorMetadataOwnerReferencesRequiresOwnerReference) WithValue(value bool) *Project_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Project_FieldSubPathValue)
}

func (s ProjectPathSelectorMetadataOwnerReferencesRequiresOwnerReference) WithArrayOfValues(values []bool) *Project_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Project_FieldSubPathArrayOfValues)
}

type ProjectPathSelectorMetadataShards struct{}

func (ProjectPathSelectorMetadataShards) FieldPath() *Project_FieldSubPath {
	return &Project_FieldSubPath{
		selector: Project_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Shards().FieldPath(),
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
		subPath:  meta.NewMetaFieldPathBuilder().Shards().WithKey(s.key).FieldPath(),
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
		subPath:  meta.NewMetaFieldPathBuilder().Syncing().FieldPath(),
	}
}

func (s ProjectPathSelectorMetadataSyncing) WithValue(value *meta.SyncingMeta) *Project_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Project_FieldSubPathValue)
}

func (s ProjectPathSelectorMetadataSyncing) WithArrayOfValues(values []*meta.SyncingMeta) *Project_FieldSubPathArrayOfValues {
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
		subPath:  meta.NewMetaFieldPathBuilder().Syncing().OwningRegion().FieldPath(),
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
		subPath:  meta.NewMetaFieldPathBuilder().Syncing().Regions().FieldPath(),
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

type ProjectPathSelectorMetadataLifecycle struct{}

func (ProjectPathSelectorMetadataLifecycle) FieldPath() *Project_FieldSubPath {
	return &Project_FieldSubPath{
		selector: Project_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Lifecycle().FieldPath(),
	}
}

func (s ProjectPathSelectorMetadataLifecycle) WithValue(value *meta.Lifecycle) *Project_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Project_FieldSubPathValue)
}

func (s ProjectPathSelectorMetadataLifecycle) WithArrayOfValues(values []*meta.Lifecycle) *Project_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Project_FieldSubPathArrayOfValues)
}

func (ProjectPathSelectorMetadataLifecycle) State() ProjectPathSelectorMetadataLifecycleState {
	return ProjectPathSelectorMetadataLifecycleState{}
}

func (ProjectPathSelectorMetadataLifecycle) BlockDeletion() ProjectPathSelectorMetadataLifecycleBlockDeletion {
	return ProjectPathSelectorMetadataLifecycleBlockDeletion{}
}

type ProjectPathSelectorMetadataLifecycleState struct{}

func (ProjectPathSelectorMetadataLifecycleState) FieldPath() *Project_FieldSubPath {
	return &Project_FieldSubPath{
		selector: Project_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Lifecycle().State().FieldPath(),
	}
}

func (s ProjectPathSelectorMetadataLifecycleState) WithValue(value meta.Lifecycle_State) *Project_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Project_FieldSubPathValue)
}

func (s ProjectPathSelectorMetadataLifecycleState) WithArrayOfValues(values []meta.Lifecycle_State) *Project_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Project_FieldSubPathArrayOfValues)
}

type ProjectPathSelectorMetadataLifecycleBlockDeletion struct{}

func (ProjectPathSelectorMetadataLifecycleBlockDeletion) FieldPath() *Project_FieldSubPath {
	return &Project_FieldSubPath{
		selector: Project_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Lifecycle().BlockDeletion().FieldPath(),
	}
}

func (s ProjectPathSelectorMetadataLifecycleBlockDeletion) WithValue(value bool) *Project_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Project_FieldSubPathValue)
}

func (s ProjectPathSelectorMetadataLifecycleBlockDeletion) WithArrayOfValues(values []bool) *Project_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Project_FieldSubPathArrayOfValues)
}

type ProjectPathSelectorMetadataServices struct{}

func (ProjectPathSelectorMetadataServices) FieldPath() *Project_FieldSubPath {
	return &Project_FieldSubPath{
		selector: Project_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Services().FieldPath(),
	}
}

func (s ProjectPathSelectorMetadataServices) WithValue(value *meta.ServicesInfo) *Project_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Project_FieldSubPathValue)
}

func (s ProjectPathSelectorMetadataServices) WithArrayOfValues(values []*meta.ServicesInfo) *Project_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Project_FieldSubPathArrayOfValues)
}

func (ProjectPathSelectorMetadataServices) OwningService() ProjectPathSelectorMetadataServicesOwningService {
	return ProjectPathSelectorMetadataServicesOwningService{}
}

func (ProjectPathSelectorMetadataServices) AllowedServices() ProjectPathSelectorMetadataServicesAllowedServices {
	return ProjectPathSelectorMetadataServicesAllowedServices{}
}

type ProjectPathSelectorMetadataServicesOwningService struct{}

func (ProjectPathSelectorMetadataServicesOwningService) FieldPath() *Project_FieldSubPath {
	return &Project_FieldSubPath{
		selector: Project_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Services().OwningService().FieldPath(),
	}
}

func (s ProjectPathSelectorMetadataServicesOwningService) WithValue(value string) *Project_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Project_FieldSubPathValue)
}

func (s ProjectPathSelectorMetadataServicesOwningService) WithArrayOfValues(values []string) *Project_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Project_FieldSubPathArrayOfValues)
}

type ProjectPathSelectorMetadataServicesAllowedServices struct{}

func (ProjectPathSelectorMetadataServicesAllowedServices) FieldPath() *Project_FieldSubPath {
	return &Project_FieldSubPath{
		selector: Project_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Services().AllowedServices().FieldPath(),
	}
}

func (s ProjectPathSelectorMetadataServicesAllowedServices) WithValue(value []string) *Project_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Project_FieldSubPathValue)
}

func (s ProjectPathSelectorMetadataServicesAllowedServices) WithArrayOfValues(values [][]string) *Project_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Project_FieldSubPathArrayOfValues)
}

func (s ProjectPathSelectorMetadataServicesAllowedServices) WithItemValue(value string) *Project_FieldSubPathArrayItemValue {
	return s.FieldPath().WithIArrayItemValue(value).(*Project_FieldSubPathArrayItemValue)
}

type ProjectPathSelectorMultiRegionPolicy struct{}

func (ProjectPathSelectorMultiRegionPolicy) FieldPath() *Project_FieldTerminalPath {
	return &Project_FieldTerminalPath{selector: Project_FieldPathSelectorMultiRegionPolicy}
}

func (s ProjectPathSelectorMultiRegionPolicy) WithValue(value *multi_region_policy.MultiRegionPolicy) *Project_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*Project_FieldTerminalPathValue)
}

func (s ProjectPathSelectorMultiRegionPolicy) WithArrayOfValues(values []*multi_region_policy.MultiRegionPolicy) *Project_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Project_FieldTerminalPathArrayOfValues)
}

func (ProjectPathSelectorMultiRegionPolicy) WithSubPath(subPath multi_region_policy.MultiRegionPolicy_FieldPath) *Project_FieldSubPath {
	return &Project_FieldSubPath{selector: Project_FieldPathSelectorMultiRegionPolicy, subPath: subPath}
}

func (s ProjectPathSelectorMultiRegionPolicy) WithSubValue(subPathValue multi_region_policy.MultiRegionPolicy_FieldPathValue) *Project_FieldSubPathValue {
	return &Project_FieldSubPathValue{Project_FieldPath: s.WithSubPath(subPathValue), subPathValue: subPathValue}
}

func (s ProjectPathSelectorMultiRegionPolicy) WithSubArrayOfValues(subPathArrayOfValues multi_region_policy.MultiRegionPolicy_FieldPathArrayOfValues) *Project_FieldSubPathArrayOfValues {
	return &Project_FieldSubPathArrayOfValues{Project_FieldPath: s.WithSubPath(subPathArrayOfValues), subPathArrayOfValues: subPathArrayOfValues}
}

func (s ProjectPathSelectorMultiRegionPolicy) WithSubArrayItemValue(subPathArrayItemValue multi_region_policy.MultiRegionPolicy_FieldPathArrayItemValue) *Project_FieldSubPathArrayItemValue {
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
		subPath:  multi_region_policy.NewMultiRegionPolicyFieldPathBuilder().EnabledRegions().FieldPath(),
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
		subPath:  multi_region_policy.NewMultiRegionPolicyFieldPathBuilder().DefaultControlRegion().FieldPath(),
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
		subPath:  multi_region_policy.NewMultiRegionPolicyFieldPathBuilder().CriteriaForDisabledSync().FieldPath(),
	}
}

func (s ProjectPathSelectorMultiRegionPolicyCriteriaForDisabledSync) WithValue(value []*multi_region_policy.MultiRegionPolicy_CriteriaForDisabledSync) *Project_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Project_FieldSubPathValue)
}

func (s ProjectPathSelectorMultiRegionPolicyCriteriaForDisabledSync) WithArrayOfValues(values [][]*multi_region_policy.MultiRegionPolicy_CriteriaForDisabledSync) *Project_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Project_FieldSubPathArrayOfValues)
}

func (s ProjectPathSelectorMultiRegionPolicyCriteriaForDisabledSync) WithItemValue(value *multi_region_policy.MultiRegionPolicy_CriteriaForDisabledSync) *Project_FieldSubPathArrayItemValue {
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
		subPath:  multi_region_policy.NewMultiRegionPolicyFieldPathBuilder().CriteriaForDisabledSync().ResourceTypeName().FieldPath(),
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
		subPath:  multi_region_policy.NewMultiRegionPolicyFieldPathBuilder().CriteriaForDisabledSync().SourceRegion().FieldPath(),
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
		subPath:  multi_region_policy.NewMultiRegionPolicyFieldPathBuilder().CriteriaForDisabledSync().DestRegion().FieldPath(),
	}
}

func (s ProjectPathSelectorMultiRegionPolicyCriteriaForDisabledSyncDestRegion) WithValue(value string) *Project_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Project_FieldSubPathValue)
}

func (s ProjectPathSelectorMultiRegionPolicyCriteriaForDisabledSyncDestRegion) WithArrayOfValues(values []string) *Project_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Project_FieldSubPathArrayOfValues)
}
