// Code generated by protoc-gen-goten-object
// File: edgelq/iam/proto/v1/group.proto
// DO NOT EDIT!!!

package group

// proto imports
import (
	iam_common "github.com/cloudwan/edgelq-sdk/iam/resources/v1/common"
	organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1/organization"
	project "github.com/cloudwan/edgelq-sdk/iam/resources/v1/project"
	meta_common "github.com/cloudwan/goten-sdk/meta-service/resources/v1/common"
	meta_service "github.com/cloudwan/goten-sdk/meta-service/resources/v1/service"
	meta "github.com/cloudwan/goten-sdk/types/meta"
	multi_region_policy "github.com/cloudwan/goten-sdk/types/multi_region_policy"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

// make sure we're using proto imports
var (
	_ = &iam_common.PCR{}
	_ = &organization.Organization{}
	_ = &project.Project{}
	_ = &timestamppb.Timestamp{}
	_ = &meta_common.LabelledDomain{}
	_ = &meta_service.Service{}
	_ = &meta.Meta{}
	_ = &multi_region_policy.MultiRegionPolicy{}
)

type GroupFieldPathBuilder struct{}

func NewGroupFieldPathBuilder() GroupFieldPathBuilder {
	return GroupFieldPathBuilder{}
}
func (GroupFieldPathBuilder) Name() GroupPathSelectorName {
	return GroupPathSelectorName{}
}
func (GroupFieldPathBuilder) Metadata() GroupPathSelectorMetadata {
	return GroupPathSelectorMetadata{}
}
func (GroupFieldPathBuilder) DisplayName() GroupPathSelectorDisplayName {
	return GroupPathSelectorDisplayName{}
}
func (GroupFieldPathBuilder) Description() GroupPathSelectorDescription {
	return GroupPathSelectorDescription{}
}
func (GroupFieldPathBuilder) Email() GroupPathSelectorEmail {
	return GroupPathSelectorEmail{}
}

type GroupPathSelectorName struct{}

func (GroupPathSelectorName) FieldPath() *Group_FieldTerminalPath {
	return &Group_FieldTerminalPath{selector: Group_FieldPathSelectorName}
}

func (s GroupPathSelectorName) WithValue(value *Name) *Group_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*Group_FieldTerminalPathValue)
}

func (s GroupPathSelectorName) WithArrayOfValues(values []*Name) *Group_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Group_FieldTerminalPathArrayOfValues)
}

type GroupPathSelectorMetadata struct{}

func (GroupPathSelectorMetadata) FieldPath() *Group_FieldTerminalPath {
	return &Group_FieldTerminalPath{selector: Group_FieldPathSelectorMetadata}
}

func (s GroupPathSelectorMetadata) WithValue(value *meta.Meta) *Group_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*Group_FieldTerminalPathValue)
}

func (s GroupPathSelectorMetadata) WithArrayOfValues(values []*meta.Meta) *Group_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Group_FieldTerminalPathArrayOfValues)
}

func (GroupPathSelectorMetadata) WithSubPath(subPath meta.Meta_FieldPath) *Group_FieldSubPath {
	return &Group_FieldSubPath{selector: Group_FieldPathSelectorMetadata, subPath: subPath}
}

func (s GroupPathSelectorMetadata) WithSubValue(subPathValue meta.Meta_FieldPathValue) *Group_FieldSubPathValue {
	return &Group_FieldSubPathValue{Group_FieldPath: s.WithSubPath(subPathValue), subPathValue: subPathValue}
}

func (s GroupPathSelectorMetadata) WithSubArrayOfValues(subPathArrayOfValues meta.Meta_FieldPathArrayOfValues) *Group_FieldSubPathArrayOfValues {
	return &Group_FieldSubPathArrayOfValues{Group_FieldPath: s.WithSubPath(subPathArrayOfValues), subPathArrayOfValues: subPathArrayOfValues}
}

func (s GroupPathSelectorMetadata) WithSubArrayItemValue(subPathArrayItemValue meta.Meta_FieldPathArrayItemValue) *Group_FieldSubPathArrayItemValue {
	return &Group_FieldSubPathArrayItemValue{Group_FieldPath: s.WithSubPath(subPathArrayItemValue), subPathItemValue: subPathArrayItemValue}
}

func (GroupPathSelectorMetadata) CreateTime() GroupPathSelectorMetadataCreateTime {
	return GroupPathSelectorMetadataCreateTime{}
}

func (GroupPathSelectorMetadata) UpdateTime() GroupPathSelectorMetadataUpdateTime {
	return GroupPathSelectorMetadataUpdateTime{}
}

func (GroupPathSelectorMetadata) DeleteTime() GroupPathSelectorMetadataDeleteTime {
	return GroupPathSelectorMetadataDeleteTime{}
}

func (GroupPathSelectorMetadata) Uuid() GroupPathSelectorMetadataUuid {
	return GroupPathSelectorMetadataUuid{}
}

func (GroupPathSelectorMetadata) Tags() GroupPathSelectorMetadataTags {
	return GroupPathSelectorMetadataTags{}
}

func (GroupPathSelectorMetadata) Labels() GroupPathSelectorMetadataLabels {
	return GroupPathSelectorMetadataLabels{}
}

func (GroupPathSelectorMetadata) Annotations() GroupPathSelectorMetadataAnnotations {
	return GroupPathSelectorMetadataAnnotations{}
}

func (GroupPathSelectorMetadata) Generation() GroupPathSelectorMetadataGeneration {
	return GroupPathSelectorMetadataGeneration{}
}

func (GroupPathSelectorMetadata) ResourceVersion() GroupPathSelectorMetadataResourceVersion {
	return GroupPathSelectorMetadataResourceVersion{}
}

func (GroupPathSelectorMetadata) OwnerReferences() GroupPathSelectorMetadataOwnerReferences {
	return GroupPathSelectorMetadataOwnerReferences{}
}

func (GroupPathSelectorMetadata) Shards() GroupPathSelectorMetadataShards {
	return GroupPathSelectorMetadataShards{}
}

func (GroupPathSelectorMetadata) Syncing() GroupPathSelectorMetadataSyncing {
	return GroupPathSelectorMetadataSyncing{}
}

func (GroupPathSelectorMetadata) Lifecycle() GroupPathSelectorMetadataLifecycle {
	return GroupPathSelectorMetadataLifecycle{}
}

func (GroupPathSelectorMetadata) Services() GroupPathSelectorMetadataServices {
	return GroupPathSelectorMetadataServices{}
}

type GroupPathSelectorMetadataCreateTime struct{}

func (GroupPathSelectorMetadataCreateTime) FieldPath() *Group_FieldSubPath {
	return &Group_FieldSubPath{
		selector: Group_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().CreateTime().FieldPath(),
	}
}

func (s GroupPathSelectorMetadataCreateTime) WithValue(value *timestamppb.Timestamp) *Group_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Group_FieldSubPathValue)
}

func (s GroupPathSelectorMetadataCreateTime) WithArrayOfValues(values []*timestamppb.Timestamp) *Group_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Group_FieldSubPathArrayOfValues)
}

type GroupPathSelectorMetadataUpdateTime struct{}

func (GroupPathSelectorMetadataUpdateTime) FieldPath() *Group_FieldSubPath {
	return &Group_FieldSubPath{
		selector: Group_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().UpdateTime().FieldPath(),
	}
}

func (s GroupPathSelectorMetadataUpdateTime) WithValue(value *timestamppb.Timestamp) *Group_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Group_FieldSubPathValue)
}

func (s GroupPathSelectorMetadataUpdateTime) WithArrayOfValues(values []*timestamppb.Timestamp) *Group_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Group_FieldSubPathArrayOfValues)
}

type GroupPathSelectorMetadataDeleteTime struct{}

func (GroupPathSelectorMetadataDeleteTime) FieldPath() *Group_FieldSubPath {
	return &Group_FieldSubPath{
		selector: Group_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().DeleteTime().FieldPath(),
	}
}

func (s GroupPathSelectorMetadataDeleteTime) WithValue(value *timestamppb.Timestamp) *Group_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Group_FieldSubPathValue)
}

func (s GroupPathSelectorMetadataDeleteTime) WithArrayOfValues(values []*timestamppb.Timestamp) *Group_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Group_FieldSubPathArrayOfValues)
}

type GroupPathSelectorMetadataUuid struct{}

func (GroupPathSelectorMetadataUuid) FieldPath() *Group_FieldSubPath {
	return &Group_FieldSubPath{
		selector: Group_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Uuid().FieldPath(),
	}
}

func (s GroupPathSelectorMetadataUuid) WithValue(value string) *Group_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Group_FieldSubPathValue)
}

func (s GroupPathSelectorMetadataUuid) WithArrayOfValues(values []string) *Group_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Group_FieldSubPathArrayOfValues)
}

type GroupPathSelectorMetadataTags struct{}

func (GroupPathSelectorMetadataTags) FieldPath() *Group_FieldSubPath {
	return &Group_FieldSubPath{
		selector: Group_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Tags().FieldPath(),
	}
}

func (s GroupPathSelectorMetadataTags) WithValue(value []string) *Group_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Group_FieldSubPathValue)
}

func (s GroupPathSelectorMetadataTags) WithArrayOfValues(values [][]string) *Group_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Group_FieldSubPathArrayOfValues)
}

func (s GroupPathSelectorMetadataTags) WithItemValue(value string) *Group_FieldSubPathArrayItemValue {
	return s.FieldPath().WithIArrayItemValue(value).(*Group_FieldSubPathArrayItemValue)
}

type GroupPathSelectorMetadataLabels struct{}

func (GroupPathSelectorMetadataLabels) FieldPath() *Group_FieldSubPath {
	return &Group_FieldSubPath{
		selector: Group_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Labels().FieldPath(),
	}
}

func (s GroupPathSelectorMetadataLabels) WithValue(value map[string]string) *Group_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Group_FieldSubPathValue)
}

func (s GroupPathSelectorMetadataLabels) WithArrayOfValues(values []map[string]string) *Group_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Group_FieldSubPathArrayOfValues)
}

func (GroupPathSelectorMetadataLabels) WithKey(key string) GroupMapPathSelectorMetadataLabels {
	return GroupMapPathSelectorMetadataLabels{key: key}
}

type GroupMapPathSelectorMetadataLabels struct {
	key string
}

func (s GroupMapPathSelectorMetadataLabels) FieldPath() *Group_FieldSubPath {
	return &Group_FieldSubPath{
		selector: Group_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Labels().WithKey(s.key).FieldPath(),
	}
}

func (s GroupMapPathSelectorMetadataLabels) WithValue(value string) *Group_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Group_FieldSubPathValue)
}

func (s GroupMapPathSelectorMetadataLabels) WithArrayOfValues(values []string) *Group_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Group_FieldSubPathArrayOfValues)
}

type GroupPathSelectorMetadataAnnotations struct{}

func (GroupPathSelectorMetadataAnnotations) FieldPath() *Group_FieldSubPath {
	return &Group_FieldSubPath{
		selector: Group_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Annotations().FieldPath(),
	}
}

func (s GroupPathSelectorMetadataAnnotations) WithValue(value map[string]string) *Group_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Group_FieldSubPathValue)
}

func (s GroupPathSelectorMetadataAnnotations) WithArrayOfValues(values []map[string]string) *Group_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Group_FieldSubPathArrayOfValues)
}

func (GroupPathSelectorMetadataAnnotations) WithKey(key string) GroupMapPathSelectorMetadataAnnotations {
	return GroupMapPathSelectorMetadataAnnotations{key: key}
}

type GroupMapPathSelectorMetadataAnnotations struct {
	key string
}

func (s GroupMapPathSelectorMetadataAnnotations) FieldPath() *Group_FieldSubPath {
	return &Group_FieldSubPath{
		selector: Group_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Annotations().WithKey(s.key).FieldPath(),
	}
}

func (s GroupMapPathSelectorMetadataAnnotations) WithValue(value string) *Group_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Group_FieldSubPathValue)
}

func (s GroupMapPathSelectorMetadataAnnotations) WithArrayOfValues(values []string) *Group_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Group_FieldSubPathArrayOfValues)
}

type GroupPathSelectorMetadataGeneration struct{}

func (GroupPathSelectorMetadataGeneration) FieldPath() *Group_FieldSubPath {
	return &Group_FieldSubPath{
		selector: Group_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Generation().FieldPath(),
	}
}

func (s GroupPathSelectorMetadataGeneration) WithValue(value int64) *Group_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Group_FieldSubPathValue)
}

func (s GroupPathSelectorMetadataGeneration) WithArrayOfValues(values []int64) *Group_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Group_FieldSubPathArrayOfValues)
}

type GroupPathSelectorMetadataResourceVersion struct{}

func (GroupPathSelectorMetadataResourceVersion) FieldPath() *Group_FieldSubPath {
	return &Group_FieldSubPath{
		selector: Group_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().ResourceVersion().FieldPath(),
	}
}

func (s GroupPathSelectorMetadataResourceVersion) WithValue(value string) *Group_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Group_FieldSubPathValue)
}

func (s GroupPathSelectorMetadataResourceVersion) WithArrayOfValues(values []string) *Group_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Group_FieldSubPathArrayOfValues)
}

type GroupPathSelectorMetadataOwnerReferences struct{}

func (GroupPathSelectorMetadataOwnerReferences) FieldPath() *Group_FieldSubPath {
	return &Group_FieldSubPath{
		selector: Group_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().OwnerReferences().FieldPath(),
	}
}

func (s GroupPathSelectorMetadataOwnerReferences) WithValue(value []*meta.OwnerReference) *Group_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Group_FieldSubPathValue)
}

func (s GroupPathSelectorMetadataOwnerReferences) WithArrayOfValues(values [][]*meta.OwnerReference) *Group_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Group_FieldSubPathArrayOfValues)
}

func (s GroupPathSelectorMetadataOwnerReferences) WithItemValue(value *meta.OwnerReference) *Group_FieldSubPathArrayItemValue {
	return s.FieldPath().WithIArrayItemValue(value).(*Group_FieldSubPathArrayItemValue)
}

func (GroupPathSelectorMetadataOwnerReferences) Kind() GroupPathSelectorMetadataOwnerReferencesKind {
	return GroupPathSelectorMetadataOwnerReferencesKind{}
}

func (GroupPathSelectorMetadataOwnerReferences) Version() GroupPathSelectorMetadataOwnerReferencesVersion {
	return GroupPathSelectorMetadataOwnerReferencesVersion{}
}

func (GroupPathSelectorMetadataOwnerReferences) Name() GroupPathSelectorMetadataOwnerReferencesName {
	return GroupPathSelectorMetadataOwnerReferencesName{}
}

func (GroupPathSelectorMetadataOwnerReferences) Region() GroupPathSelectorMetadataOwnerReferencesRegion {
	return GroupPathSelectorMetadataOwnerReferencesRegion{}
}

func (GroupPathSelectorMetadataOwnerReferences) Controller() GroupPathSelectorMetadataOwnerReferencesController {
	return GroupPathSelectorMetadataOwnerReferencesController{}
}

func (GroupPathSelectorMetadataOwnerReferences) RequiresOwnerReference() GroupPathSelectorMetadataOwnerReferencesRequiresOwnerReference {
	return GroupPathSelectorMetadataOwnerReferencesRequiresOwnerReference{}
}

func (GroupPathSelectorMetadataOwnerReferences) UnsetOnDelete() GroupPathSelectorMetadataOwnerReferencesUnsetOnDelete {
	return GroupPathSelectorMetadataOwnerReferencesUnsetOnDelete{}
}

type GroupPathSelectorMetadataOwnerReferencesKind struct{}

func (GroupPathSelectorMetadataOwnerReferencesKind) FieldPath() *Group_FieldSubPath {
	return &Group_FieldSubPath{
		selector: Group_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().OwnerReferences().Kind().FieldPath(),
	}
}

func (s GroupPathSelectorMetadataOwnerReferencesKind) WithValue(value string) *Group_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Group_FieldSubPathValue)
}

func (s GroupPathSelectorMetadataOwnerReferencesKind) WithArrayOfValues(values []string) *Group_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Group_FieldSubPathArrayOfValues)
}

type GroupPathSelectorMetadataOwnerReferencesVersion struct{}

func (GroupPathSelectorMetadataOwnerReferencesVersion) FieldPath() *Group_FieldSubPath {
	return &Group_FieldSubPath{
		selector: Group_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().OwnerReferences().Version().FieldPath(),
	}
}

func (s GroupPathSelectorMetadataOwnerReferencesVersion) WithValue(value string) *Group_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Group_FieldSubPathValue)
}

func (s GroupPathSelectorMetadataOwnerReferencesVersion) WithArrayOfValues(values []string) *Group_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Group_FieldSubPathArrayOfValues)
}

type GroupPathSelectorMetadataOwnerReferencesName struct{}

func (GroupPathSelectorMetadataOwnerReferencesName) FieldPath() *Group_FieldSubPath {
	return &Group_FieldSubPath{
		selector: Group_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().OwnerReferences().Name().FieldPath(),
	}
}

func (s GroupPathSelectorMetadataOwnerReferencesName) WithValue(value string) *Group_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Group_FieldSubPathValue)
}

func (s GroupPathSelectorMetadataOwnerReferencesName) WithArrayOfValues(values []string) *Group_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Group_FieldSubPathArrayOfValues)
}

type GroupPathSelectorMetadataOwnerReferencesRegion struct{}

func (GroupPathSelectorMetadataOwnerReferencesRegion) FieldPath() *Group_FieldSubPath {
	return &Group_FieldSubPath{
		selector: Group_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().OwnerReferences().Region().FieldPath(),
	}
}

func (s GroupPathSelectorMetadataOwnerReferencesRegion) WithValue(value string) *Group_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Group_FieldSubPathValue)
}

func (s GroupPathSelectorMetadataOwnerReferencesRegion) WithArrayOfValues(values []string) *Group_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Group_FieldSubPathArrayOfValues)
}

type GroupPathSelectorMetadataOwnerReferencesController struct{}

func (GroupPathSelectorMetadataOwnerReferencesController) FieldPath() *Group_FieldSubPath {
	return &Group_FieldSubPath{
		selector: Group_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().OwnerReferences().Controller().FieldPath(),
	}
}

func (s GroupPathSelectorMetadataOwnerReferencesController) WithValue(value bool) *Group_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Group_FieldSubPathValue)
}

func (s GroupPathSelectorMetadataOwnerReferencesController) WithArrayOfValues(values []bool) *Group_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Group_FieldSubPathArrayOfValues)
}

type GroupPathSelectorMetadataOwnerReferencesRequiresOwnerReference struct{}

func (GroupPathSelectorMetadataOwnerReferencesRequiresOwnerReference) FieldPath() *Group_FieldSubPath {
	return &Group_FieldSubPath{
		selector: Group_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().OwnerReferences().RequiresOwnerReference().FieldPath(),
	}
}

func (s GroupPathSelectorMetadataOwnerReferencesRequiresOwnerReference) WithValue(value bool) *Group_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Group_FieldSubPathValue)
}

func (s GroupPathSelectorMetadataOwnerReferencesRequiresOwnerReference) WithArrayOfValues(values []bool) *Group_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Group_FieldSubPathArrayOfValues)
}

type GroupPathSelectorMetadataOwnerReferencesUnsetOnDelete struct{}

func (GroupPathSelectorMetadataOwnerReferencesUnsetOnDelete) FieldPath() *Group_FieldSubPath {
	return &Group_FieldSubPath{
		selector: Group_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().OwnerReferences().UnsetOnDelete().FieldPath(),
	}
}

func (s GroupPathSelectorMetadataOwnerReferencesUnsetOnDelete) WithValue(value bool) *Group_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Group_FieldSubPathValue)
}

func (s GroupPathSelectorMetadataOwnerReferencesUnsetOnDelete) WithArrayOfValues(values []bool) *Group_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Group_FieldSubPathArrayOfValues)
}

type GroupPathSelectorMetadataShards struct{}

func (GroupPathSelectorMetadataShards) FieldPath() *Group_FieldSubPath {
	return &Group_FieldSubPath{
		selector: Group_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Shards().FieldPath(),
	}
}

func (s GroupPathSelectorMetadataShards) WithValue(value map[string]int64) *Group_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Group_FieldSubPathValue)
}

func (s GroupPathSelectorMetadataShards) WithArrayOfValues(values []map[string]int64) *Group_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Group_FieldSubPathArrayOfValues)
}

func (GroupPathSelectorMetadataShards) WithKey(key string) GroupMapPathSelectorMetadataShards {
	return GroupMapPathSelectorMetadataShards{key: key}
}

type GroupMapPathSelectorMetadataShards struct {
	key string
}

func (s GroupMapPathSelectorMetadataShards) FieldPath() *Group_FieldSubPath {
	return &Group_FieldSubPath{
		selector: Group_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Shards().WithKey(s.key).FieldPath(),
	}
}

func (s GroupMapPathSelectorMetadataShards) WithValue(value int64) *Group_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Group_FieldSubPathValue)
}

func (s GroupMapPathSelectorMetadataShards) WithArrayOfValues(values []int64) *Group_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Group_FieldSubPathArrayOfValues)
}

type GroupPathSelectorMetadataSyncing struct{}

func (GroupPathSelectorMetadataSyncing) FieldPath() *Group_FieldSubPath {
	return &Group_FieldSubPath{
		selector: Group_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Syncing().FieldPath(),
	}
}

func (s GroupPathSelectorMetadataSyncing) WithValue(value *meta.SyncingMeta) *Group_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Group_FieldSubPathValue)
}

func (s GroupPathSelectorMetadataSyncing) WithArrayOfValues(values []*meta.SyncingMeta) *Group_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Group_FieldSubPathArrayOfValues)
}

func (GroupPathSelectorMetadataSyncing) OwningRegion() GroupPathSelectorMetadataSyncingOwningRegion {
	return GroupPathSelectorMetadataSyncingOwningRegion{}
}

func (GroupPathSelectorMetadataSyncing) Regions() GroupPathSelectorMetadataSyncingRegions {
	return GroupPathSelectorMetadataSyncingRegions{}
}

type GroupPathSelectorMetadataSyncingOwningRegion struct{}

func (GroupPathSelectorMetadataSyncingOwningRegion) FieldPath() *Group_FieldSubPath {
	return &Group_FieldSubPath{
		selector: Group_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Syncing().OwningRegion().FieldPath(),
	}
}

func (s GroupPathSelectorMetadataSyncingOwningRegion) WithValue(value string) *Group_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Group_FieldSubPathValue)
}

func (s GroupPathSelectorMetadataSyncingOwningRegion) WithArrayOfValues(values []string) *Group_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Group_FieldSubPathArrayOfValues)
}

type GroupPathSelectorMetadataSyncingRegions struct{}

func (GroupPathSelectorMetadataSyncingRegions) FieldPath() *Group_FieldSubPath {
	return &Group_FieldSubPath{
		selector: Group_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Syncing().Regions().FieldPath(),
	}
}

func (s GroupPathSelectorMetadataSyncingRegions) WithValue(value []string) *Group_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Group_FieldSubPathValue)
}

func (s GroupPathSelectorMetadataSyncingRegions) WithArrayOfValues(values [][]string) *Group_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Group_FieldSubPathArrayOfValues)
}

func (s GroupPathSelectorMetadataSyncingRegions) WithItemValue(value string) *Group_FieldSubPathArrayItemValue {
	return s.FieldPath().WithIArrayItemValue(value).(*Group_FieldSubPathArrayItemValue)
}

type GroupPathSelectorMetadataLifecycle struct{}

func (GroupPathSelectorMetadataLifecycle) FieldPath() *Group_FieldSubPath {
	return &Group_FieldSubPath{
		selector: Group_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Lifecycle().FieldPath(),
	}
}

func (s GroupPathSelectorMetadataLifecycle) WithValue(value *meta.Lifecycle) *Group_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Group_FieldSubPathValue)
}

func (s GroupPathSelectorMetadataLifecycle) WithArrayOfValues(values []*meta.Lifecycle) *Group_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Group_FieldSubPathArrayOfValues)
}

func (GroupPathSelectorMetadataLifecycle) State() GroupPathSelectorMetadataLifecycleState {
	return GroupPathSelectorMetadataLifecycleState{}
}

func (GroupPathSelectorMetadataLifecycle) BlockDeletion() GroupPathSelectorMetadataLifecycleBlockDeletion {
	return GroupPathSelectorMetadataLifecycleBlockDeletion{}
}

type GroupPathSelectorMetadataLifecycleState struct{}

func (GroupPathSelectorMetadataLifecycleState) FieldPath() *Group_FieldSubPath {
	return &Group_FieldSubPath{
		selector: Group_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Lifecycle().State().FieldPath(),
	}
}

func (s GroupPathSelectorMetadataLifecycleState) WithValue(value meta.Lifecycle_State) *Group_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Group_FieldSubPathValue)
}

func (s GroupPathSelectorMetadataLifecycleState) WithArrayOfValues(values []meta.Lifecycle_State) *Group_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Group_FieldSubPathArrayOfValues)
}

type GroupPathSelectorMetadataLifecycleBlockDeletion struct{}

func (GroupPathSelectorMetadataLifecycleBlockDeletion) FieldPath() *Group_FieldSubPath {
	return &Group_FieldSubPath{
		selector: Group_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Lifecycle().BlockDeletion().FieldPath(),
	}
}

func (s GroupPathSelectorMetadataLifecycleBlockDeletion) WithValue(value bool) *Group_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Group_FieldSubPathValue)
}

func (s GroupPathSelectorMetadataLifecycleBlockDeletion) WithArrayOfValues(values []bool) *Group_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Group_FieldSubPathArrayOfValues)
}

type GroupPathSelectorMetadataServices struct{}

func (GroupPathSelectorMetadataServices) FieldPath() *Group_FieldSubPath {
	return &Group_FieldSubPath{
		selector: Group_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Services().FieldPath(),
	}
}

func (s GroupPathSelectorMetadataServices) WithValue(value *meta.ServicesInfo) *Group_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Group_FieldSubPathValue)
}

func (s GroupPathSelectorMetadataServices) WithArrayOfValues(values []*meta.ServicesInfo) *Group_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Group_FieldSubPathArrayOfValues)
}

func (GroupPathSelectorMetadataServices) OwningService() GroupPathSelectorMetadataServicesOwningService {
	return GroupPathSelectorMetadataServicesOwningService{}
}

func (GroupPathSelectorMetadataServices) AllowedServices() GroupPathSelectorMetadataServicesAllowedServices {
	return GroupPathSelectorMetadataServicesAllowedServices{}
}

type GroupPathSelectorMetadataServicesOwningService struct{}

func (GroupPathSelectorMetadataServicesOwningService) FieldPath() *Group_FieldSubPath {
	return &Group_FieldSubPath{
		selector: Group_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Services().OwningService().FieldPath(),
	}
}

func (s GroupPathSelectorMetadataServicesOwningService) WithValue(value string) *Group_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Group_FieldSubPathValue)
}

func (s GroupPathSelectorMetadataServicesOwningService) WithArrayOfValues(values []string) *Group_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Group_FieldSubPathArrayOfValues)
}

type GroupPathSelectorMetadataServicesAllowedServices struct{}

func (GroupPathSelectorMetadataServicesAllowedServices) FieldPath() *Group_FieldSubPath {
	return &Group_FieldSubPath{
		selector: Group_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Services().AllowedServices().FieldPath(),
	}
}

func (s GroupPathSelectorMetadataServicesAllowedServices) WithValue(value []string) *Group_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Group_FieldSubPathValue)
}

func (s GroupPathSelectorMetadataServicesAllowedServices) WithArrayOfValues(values [][]string) *Group_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Group_FieldSubPathArrayOfValues)
}

func (s GroupPathSelectorMetadataServicesAllowedServices) WithItemValue(value string) *Group_FieldSubPathArrayItemValue {
	return s.FieldPath().WithIArrayItemValue(value).(*Group_FieldSubPathArrayItemValue)
}

type GroupPathSelectorDisplayName struct{}

func (GroupPathSelectorDisplayName) FieldPath() *Group_FieldTerminalPath {
	return &Group_FieldTerminalPath{selector: Group_FieldPathSelectorDisplayName}
}

func (s GroupPathSelectorDisplayName) WithValue(value string) *Group_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*Group_FieldTerminalPathValue)
}

func (s GroupPathSelectorDisplayName) WithArrayOfValues(values []string) *Group_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Group_FieldTerminalPathArrayOfValues)
}

type GroupPathSelectorDescription struct{}

func (GroupPathSelectorDescription) FieldPath() *Group_FieldTerminalPath {
	return &Group_FieldTerminalPath{selector: Group_FieldPathSelectorDescription}
}

func (s GroupPathSelectorDescription) WithValue(value string) *Group_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*Group_FieldTerminalPathValue)
}

func (s GroupPathSelectorDescription) WithArrayOfValues(values []string) *Group_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Group_FieldTerminalPathArrayOfValues)
}

type GroupPathSelectorEmail struct{}

func (GroupPathSelectorEmail) FieldPath() *Group_FieldTerminalPath {
	return &Group_FieldTerminalPath{selector: Group_FieldPathSelectorEmail}
}

func (s GroupPathSelectorEmail) WithValue(value string) *Group_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*Group_FieldTerminalPathValue)
}

func (s GroupPathSelectorEmail) WithArrayOfValues(values []string) *Group_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Group_FieldTerminalPathArrayOfValues)
}
