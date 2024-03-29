// Code generated by protoc-gen-goten-object
// File: edgelq/applications/proto/v1alpha2/config_map.proto
// DO NOT EDIT!!!

package config_map

// proto imports
import (
	project "github.com/cloudwan/edgelq-sdk/applications/resources/v1alpha2/project"
	meta "github.com/cloudwan/goten-sdk/types/meta"
	multi_region_policy "github.com/cloudwan/goten-sdk/types/multi_region_policy"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

// make sure we're using proto imports
var (
	_ = &project.Project{}
	_ = &timestamppb.Timestamp{}
	_ = &meta.Meta{}
	_ = &multi_region_policy.MultiRegionPolicy{}
)

type ConfigMapFieldPathBuilder struct{}

func NewConfigMapFieldPathBuilder() ConfigMapFieldPathBuilder {
	return ConfigMapFieldPathBuilder{}
}
func (ConfigMapFieldPathBuilder) Name() ConfigMapPathSelectorName {
	return ConfigMapPathSelectorName{}
}
func (ConfigMapFieldPathBuilder) DisplayName() ConfigMapPathSelectorDisplayName {
	return ConfigMapPathSelectorDisplayName{}
}
func (ConfigMapFieldPathBuilder) Metadata() ConfigMapPathSelectorMetadata {
	return ConfigMapPathSelectorMetadata{}
}
func (ConfigMapFieldPathBuilder) Data() ConfigMapPathSelectorData {
	return ConfigMapPathSelectorData{}
}
func (ConfigMapFieldPathBuilder) BinaryData() ConfigMapPathSelectorBinaryData {
	return ConfigMapPathSelectorBinaryData{}
}

type ConfigMapPathSelectorName struct{}

func (ConfigMapPathSelectorName) FieldPath() *ConfigMap_FieldTerminalPath {
	return &ConfigMap_FieldTerminalPath{selector: ConfigMap_FieldPathSelectorName}
}

func (s ConfigMapPathSelectorName) WithValue(value *Name) *ConfigMap_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*ConfigMap_FieldTerminalPathValue)
}

func (s ConfigMapPathSelectorName) WithArrayOfValues(values []*Name) *ConfigMap_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ConfigMap_FieldTerminalPathArrayOfValues)
}

type ConfigMapPathSelectorDisplayName struct{}

func (ConfigMapPathSelectorDisplayName) FieldPath() *ConfigMap_FieldTerminalPath {
	return &ConfigMap_FieldTerminalPath{selector: ConfigMap_FieldPathSelectorDisplayName}
}

func (s ConfigMapPathSelectorDisplayName) WithValue(value string) *ConfigMap_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*ConfigMap_FieldTerminalPathValue)
}

func (s ConfigMapPathSelectorDisplayName) WithArrayOfValues(values []string) *ConfigMap_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ConfigMap_FieldTerminalPathArrayOfValues)
}

type ConfigMapPathSelectorMetadata struct{}

func (ConfigMapPathSelectorMetadata) FieldPath() *ConfigMap_FieldTerminalPath {
	return &ConfigMap_FieldTerminalPath{selector: ConfigMap_FieldPathSelectorMetadata}
}

func (s ConfigMapPathSelectorMetadata) WithValue(value *meta.Meta) *ConfigMap_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*ConfigMap_FieldTerminalPathValue)
}

func (s ConfigMapPathSelectorMetadata) WithArrayOfValues(values []*meta.Meta) *ConfigMap_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ConfigMap_FieldTerminalPathArrayOfValues)
}

func (ConfigMapPathSelectorMetadata) WithSubPath(subPath meta.Meta_FieldPath) *ConfigMap_FieldSubPath {
	return &ConfigMap_FieldSubPath{selector: ConfigMap_FieldPathSelectorMetadata, subPath: subPath}
}

func (s ConfigMapPathSelectorMetadata) WithSubValue(subPathValue meta.Meta_FieldPathValue) *ConfigMap_FieldSubPathValue {
	return &ConfigMap_FieldSubPathValue{ConfigMap_FieldPath: s.WithSubPath(subPathValue), subPathValue: subPathValue}
}

func (s ConfigMapPathSelectorMetadata) WithSubArrayOfValues(subPathArrayOfValues meta.Meta_FieldPathArrayOfValues) *ConfigMap_FieldSubPathArrayOfValues {
	return &ConfigMap_FieldSubPathArrayOfValues{ConfigMap_FieldPath: s.WithSubPath(subPathArrayOfValues), subPathArrayOfValues: subPathArrayOfValues}
}

func (s ConfigMapPathSelectorMetadata) WithSubArrayItemValue(subPathArrayItemValue meta.Meta_FieldPathArrayItemValue) *ConfigMap_FieldSubPathArrayItemValue {
	return &ConfigMap_FieldSubPathArrayItemValue{ConfigMap_FieldPath: s.WithSubPath(subPathArrayItemValue), subPathItemValue: subPathArrayItemValue}
}

func (ConfigMapPathSelectorMetadata) CreateTime() ConfigMapPathSelectorMetadataCreateTime {
	return ConfigMapPathSelectorMetadataCreateTime{}
}

func (ConfigMapPathSelectorMetadata) UpdateTime() ConfigMapPathSelectorMetadataUpdateTime {
	return ConfigMapPathSelectorMetadataUpdateTime{}
}

func (ConfigMapPathSelectorMetadata) DeleteTime() ConfigMapPathSelectorMetadataDeleteTime {
	return ConfigMapPathSelectorMetadataDeleteTime{}
}

func (ConfigMapPathSelectorMetadata) Uuid() ConfigMapPathSelectorMetadataUuid {
	return ConfigMapPathSelectorMetadataUuid{}
}

func (ConfigMapPathSelectorMetadata) Tags() ConfigMapPathSelectorMetadataTags {
	return ConfigMapPathSelectorMetadataTags{}
}

func (ConfigMapPathSelectorMetadata) Labels() ConfigMapPathSelectorMetadataLabels {
	return ConfigMapPathSelectorMetadataLabels{}
}

func (ConfigMapPathSelectorMetadata) Annotations() ConfigMapPathSelectorMetadataAnnotations {
	return ConfigMapPathSelectorMetadataAnnotations{}
}

func (ConfigMapPathSelectorMetadata) Generation() ConfigMapPathSelectorMetadataGeneration {
	return ConfigMapPathSelectorMetadataGeneration{}
}

func (ConfigMapPathSelectorMetadata) ResourceVersion() ConfigMapPathSelectorMetadataResourceVersion {
	return ConfigMapPathSelectorMetadataResourceVersion{}
}

func (ConfigMapPathSelectorMetadata) OwnerReferences() ConfigMapPathSelectorMetadataOwnerReferences {
	return ConfigMapPathSelectorMetadataOwnerReferences{}
}

func (ConfigMapPathSelectorMetadata) Shards() ConfigMapPathSelectorMetadataShards {
	return ConfigMapPathSelectorMetadataShards{}
}

func (ConfigMapPathSelectorMetadata) Syncing() ConfigMapPathSelectorMetadataSyncing {
	return ConfigMapPathSelectorMetadataSyncing{}
}

func (ConfigMapPathSelectorMetadata) Lifecycle() ConfigMapPathSelectorMetadataLifecycle {
	return ConfigMapPathSelectorMetadataLifecycle{}
}

func (ConfigMapPathSelectorMetadata) Services() ConfigMapPathSelectorMetadataServices {
	return ConfigMapPathSelectorMetadataServices{}
}

type ConfigMapPathSelectorMetadataCreateTime struct{}

func (ConfigMapPathSelectorMetadataCreateTime) FieldPath() *ConfigMap_FieldSubPath {
	return &ConfigMap_FieldSubPath{
		selector: ConfigMap_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().CreateTime().FieldPath(),
	}
}

func (s ConfigMapPathSelectorMetadataCreateTime) WithValue(value *timestamppb.Timestamp) *ConfigMap_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ConfigMap_FieldSubPathValue)
}

func (s ConfigMapPathSelectorMetadataCreateTime) WithArrayOfValues(values []*timestamppb.Timestamp) *ConfigMap_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ConfigMap_FieldSubPathArrayOfValues)
}

type ConfigMapPathSelectorMetadataUpdateTime struct{}

func (ConfigMapPathSelectorMetadataUpdateTime) FieldPath() *ConfigMap_FieldSubPath {
	return &ConfigMap_FieldSubPath{
		selector: ConfigMap_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().UpdateTime().FieldPath(),
	}
}

func (s ConfigMapPathSelectorMetadataUpdateTime) WithValue(value *timestamppb.Timestamp) *ConfigMap_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ConfigMap_FieldSubPathValue)
}

func (s ConfigMapPathSelectorMetadataUpdateTime) WithArrayOfValues(values []*timestamppb.Timestamp) *ConfigMap_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ConfigMap_FieldSubPathArrayOfValues)
}

type ConfigMapPathSelectorMetadataDeleteTime struct{}

func (ConfigMapPathSelectorMetadataDeleteTime) FieldPath() *ConfigMap_FieldSubPath {
	return &ConfigMap_FieldSubPath{
		selector: ConfigMap_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().DeleteTime().FieldPath(),
	}
}

func (s ConfigMapPathSelectorMetadataDeleteTime) WithValue(value *timestamppb.Timestamp) *ConfigMap_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ConfigMap_FieldSubPathValue)
}

func (s ConfigMapPathSelectorMetadataDeleteTime) WithArrayOfValues(values []*timestamppb.Timestamp) *ConfigMap_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ConfigMap_FieldSubPathArrayOfValues)
}

type ConfigMapPathSelectorMetadataUuid struct{}

func (ConfigMapPathSelectorMetadataUuid) FieldPath() *ConfigMap_FieldSubPath {
	return &ConfigMap_FieldSubPath{
		selector: ConfigMap_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Uuid().FieldPath(),
	}
}

func (s ConfigMapPathSelectorMetadataUuid) WithValue(value string) *ConfigMap_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ConfigMap_FieldSubPathValue)
}

func (s ConfigMapPathSelectorMetadataUuid) WithArrayOfValues(values []string) *ConfigMap_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ConfigMap_FieldSubPathArrayOfValues)
}

type ConfigMapPathSelectorMetadataTags struct{}

func (ConfigMapPathSelectorMetadataTags) FieldPath() *ConfigMap_FieldSubPath {
	return &ConfigMap_FieldSubPath{
		selector: ConfigMap_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Tags().FieldPath(),
	}
}

func (s ConfigMapPathSelectorMetadataTags) WithValue(value []string) *ConfigMap_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ConfigMap_FieldSubPathValue)
}

func (s ConfigMapPathSelectorMetadataTags) WithArrayOfValues(values [][]string) *ConfigMap_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ConfigMap_FieldSubPathArrayOfValues)
}

func (s ConfigMapPathSelectorMetadataTags) WithItemValue(value string) *ConfigMap_FieldSubPathArrayItemValue {
	return s.FieldPath().WithIArrayItemValue(value).(*ConfigMap_FieldSubPathArrayItemValue)
}

type ConfigMapPathSelectorMetadataLabels struct{}

func (ConfigMapPathSelectorMetadataLabels) FieldPath() *ConfigMap_FieldSubPath {
	return &ConfigMap_FieldSubPath{
		selector: ConfigMap_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Labels().FieldPath(),
	}
}

func (s ConfigMapPathSelectorMetadataLabels) WithValue(value map[string]string) *ConfigMap_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ConfigMap_FieldSubPathValue)
}

func (s ConfigMapPathSelectorMetadataLabels) WithArrayOfValues(values []map[string]string) *ConfigMap_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ConfigMap_FieldSubPathArrayOfValues)
}

func (ConfigMapPathSelectorMetadataLabels) WithKey(key string) ConfigMapMapPathSelectorMetadataLabels {
	return ConfigMapMapPathSelectorMetadataLabels{key: key}
}

type ConfigMapMapPathSelectorMetadataLabels struct {
	key string
}

func (s ConfigMapMapPathSelectorMetadataLabels) FieldPath() *ConfigMap_FieldSubPath {
	return &ConfigMap_FieldSubPath{
		selector: ConfigMap_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Labels().WithKey(s.key).FieldPath(),
	}
}

func (s ConfigMapMapPathSelectorMetadataLabels) WithValue(value string) *ConfigMap_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ConfigMap_FieldSubPathValue)
}

func (s ConfigMapMapPathSelectorMetadataLabels) WithArrayOfValues(values []string) *ConfigMap_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ConfigMap_FieldSubPathArrayOfValues)
}

type ConfigMapPathSelectorMetadataAnnotations struct{}

func (ConfigMapPathSelectorMetadataAnnotations) FieldPath() *ConfigMap_FieldSubPath {
	return &ConfigMap_FieldSubPath{
		selector: ConfigMap_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Annotations().FieldPath(),
	}
}

func (s ConfigMapPathSelectorMetadataAnnotations) WithValue(value map[string]string) *ConfigMap_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ConfigMap_FieldSubPathValue)
}

func (s ConfigMapPathSelectorMetadataAnnotations) WithArrayOfValues(values []map[string]string) *ConfigMap_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ConfigMap_FieldSubPathArrayOfValues)
}

func (ConfigMapPathSelectorMetadataAnnotations) WithKey(key string) ConfigMapMapPathSelectorMetadataAnnotations {
	return ConfigMapMapPathSelectorMetadataAnnotations{key: key}
}

type ConfigMapMapPathSelectorMetadataAnnotations struct {
	key string
}

func (s ConfigMapMapPathSelectorMetadataAnnotations) FieldPath() *ConfigMap_FieldSubPath {
	return &ConfigMap_FieldSubPath{
		selector: ConfigMap_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Annotations().WithKey(s.key).FieldPath(),
	}
}

func (s ConfigMapMapPathSelectorMetadataAnnotations) WithValue(value string) *ConfigMap_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ConfigMap_FieldSubPathValue)
}

func (s ConfigMapMapPathSelectorMetadataAnnotations) WithArrayOfValues(values []string) *ConfigMap_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ConfigMap_FieldSubPathArrayOfValues)
}

type ConfigMapPathSelectorMetadataGeneration struct{}

func (ConfigMapPathSelectorMetadataGeneration) FieldPath() *ConfigMap_FieldSubPath {
	return &ConfigMap_FieldSubPath{
		selector: ConfigMap_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Generation().FieldPath(),
	}
}

func (s ConfigMapPathSelectorMetadataGeneration) WithValue(value int64) *ConfigMap_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ConfigMap_FieldSubPathValue)
}

func (s ConfigMapPathSelectorMetadataGeneration) WithArrayOfValues(values []int64) *ConfigMap_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ConfigMap_FieldSubPathArrayOfValues)
}

type ConfigMapPathSelectorMetadataResourceVersion struct{}

func (ConfigMapPathSelectorMetadataResourceVersion) FieldPath() *ConfigMap_FieldSubPath {
	return &ConfigMap_FieldSubPath{
		selector: ConfigMap_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().ResourceVersion().FieldPath(),
	}
}

func (s ConfigMapPathSelectorMetadataResourceVersion) WithValue(value string) *ConfigMap_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ConfigMap_FieldSubPathValue)
}

func (s ConfigMapPathSelectorMetadataResourceVersion) WithArrayOfValues(values []string) *ConfigMap_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ConfigMap_FieldSubPathArrayOfValues)
}

type ConfigMapPathSelectorMetadataOwnerReferences struct{}

func (ConfigMapPathSelectorMetadataOwnerReferences) FieldPath() *ConfigMap_FieldSubPath {
	return &ConfigMap_FieldSubPath{
		selector: ConfigMap_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().OwnerReferences().FieldPath(),
	}
}

func (s ConfigMapPathSelectorMetadataOwnerReferences) WithValue(value []*meta.OwnerReference) *ConfigMap_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ConfigMap_FieldSubPathValue)
}

func (s ConfigMapPathSelectorMetadataOwnerReferences) WithArrayOfValues(values [][]*meta.OwnerReference) *ConfigMap_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ConfigMap_FieldSubPathArrayOfValues)
}

func (s ConfigMapPathSelectorMetadataOwnerReferences) WithItemValue(value *meta.OwnerReference) *ConfigMap_FieldSubPathArrayItemValue {
	return s.FieldPath().WithIArrayItemValue(value).(*ConfigMap_FieldSubPathArrayItemValue)
}

func (ConfigMapPathSelectorMetadataOwnerReferences) Kind() ConfigMapPathSelectorMetadataOwnerReferencesKind {
	return ConfigMapPathSelectorMetadataOwnerReferencesKind{}
}

func (ConfigMapPathSelectorMetadataOwnerReferences) Version() ConfigMapPathSelectorMetadataOwnerReferencesVersion {
	return ConfigMapPathSelectorMetadataOwnerReferencesVersion{}
}

func (ConfigMapPathSelectorMetadataOwnerReferences) Name() ConfigMapPathSelectorMetadataOwnerReferencesName {
	return ConfigMapPathSelectorMetadataOwnerReferencesName{}
}

func (ConfigMapPathSelectorMetadataOwnerReferences) Region() ConfigMapPathSelectorMetadataOwnerReferencesRegion {
	return ConfigMapPathSelectorMetadataOwnerReferencesRegion{}
}

func (ConfigMapPathSelectorMetadataOwnerReferences) Controller() ConfigMapPathSelectorMetadataOwnerReferencesController {
	return ConfigMapPathSelectorMetadataOwnerReferencesController{}
}

func (ConfigMapPathSelectorMetadataOwnerReferences) RequiresOwnerReference() ConfigMapPathSelectorMetadataOwnerReferencesRequiresOwnerReference {
	return ConfigMapPathSelectorMetadataOwnerReferencesRequiresOwnerReference{}
}

type ConfigMapPathSelectorMetadataOwnerReferencesKind struct{}

func (ConfigMapPathSelectorMetadataOwnerReferencesKind) FieldPath() *ConfigMap_FieldSubPath {
	return &ConfigMap_FieldSubPath{
		selector: ConfigMap_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().OwnerReferences().Kind().FieldPath(),
	}
}

func (s ConfigMapPathSelectorMetadataOwnerReferencesKind) WithValue(value string) *ConfigMap_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ConfigMap_FieldSubPathValue)
}

func (s ConfigMapPathSelectorMetadataOwnerReferencesKind) WithArrayOfValues(values []string) *ConfigMap_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ConfigMap_FieldSubPathArrayOfValues)
}

type ConfigMapPathSelectorMetadataOwnerReferencesVersion struct{}

func (ConfigMapPathSelectorMetadataOwnerReferencesVersion) FieldPath() *ConfigMap_FieldSubPath {
	return &ConfigMap_FieldSubPath{
		selector: ConfigMap_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().OwnerReferences().Version().FieldPath(),
	}
}

func (s ConfigMapPathSelectorMetadataOwnerReferencesVersion) WithValue(value string) *ConfigMap_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ConfigMap_FieldSubPathValue)
}

func (s ConfigMapPathSelectorMetadataOwnerReferencesVersion) WithArrayOfValues(values []string) *ConfigMap_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ConfigMap_FieldSubPathArrayOfValues)
}

type ConfigMapPathSelectorMetadataOwnerReferencesName struct{}

func (ConfigMapPathSelectorMetadataOwnerReferencesName) FieldPath() *ConfigMap_FieldSubPath {
	return &ConfigMap_FieldSubPath{
		selector: ConfigMap_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().OwnerReferences().Name().FieldPath(),
	}
}

func (s ConfigMapPathSelectorMetadataOwnerReferencesName) WithValue(value string) *ConfigMap_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ConfigMap_FieldSubPathValue)
}

func (s ConfigMapPathSelectorMetadataOwnerReferencesName) WithArrayOfValues(values []string) *ConfigMap_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ConfigMap_FieldSubPathArrayOfValues)
}

type ConfigMapPathSelectorMetadataOwnerReferencesRegion struct{}

func (ConfigMapPathSelectorMetadataOwnerReferencesRegion) FieldPath() *ConfigMap_FieldSubPath {
	return &ConfigMap_FieldSubPath{
		selector: ConfigMap_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().OwnerReferences().Region().FieldPath(),
	}
}

func (s ConfigMapPathSelectorMetadataOwnerReferencesRegion) WithValue(value string) *ConfigMap_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ConfigMap_FieldSubPathValue)
}

func (s ConfigMapPathSelectorMetadataOwnerReferencesRegion) WithArrayOfValues(values []string) *ConfigMap_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ConfigMap_FieldSubPathArrayOfValues)
}

type ConfigMapPathSelectorMetadataOwnerReferencesController struct{}

func (ConfigMapPathSelectorMetadataOwnerReferencesController) FieldPath() *ConfigMap_FieldSubPath {
	return &ConfigMap_FieldSubPath{
		selector: ConfigMap_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().OwnerReferences().Controller().FieldPath(),
	}
}

func (s ConfigMapPathSelectorMetadataOwnerReferencesController) WithValue(value bool) *ConfigMap_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ConfigMap_FieldSubPathValue)
}

func (s ConfigMapPathSelectorMetadataOwnerReferencesController) WithArrayOfValues(values []bool) *ConfigMap_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ConfigMap_FieldSubPathArrayOfValues)
}

type ConfigMapPathSelectorMetadataOwnerReferencesRequiresOwnerReference struct{}

func (ConfigMapPathSelectorMetadataOwnerReferencesRequiresOwnerReference) FieldPath() *ConfigMap_FieldSubPath {
	return &ConfigMap_FieldSubPath{
		selector: ConfigMap_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().OwnerReferences().RequiresOwnerReference().FieldPath(),
	}
}

func (s ConfigMapPathSelectorMetadataOwnerReferencesRequiresOwnerReference) WithValue(value bool) *ConfigMap_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ConfigMap_FieldSubPathValue)
}

func (s ConfigMapPathSelectorMetadataOwnerReferencesRequiresOwnerReference) WithArrayOfValues(values []bool) *ConfigMap_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ConfigMap_FieldSubPathArrayOfValues)
}

type ConfigMapPathSelectorMetadataShards struct{}

func (ConfigMapPathSelectorMetadataShards) FieldPath() *ConfigMap_FieldSubPath {
	return &ConfigMap_FieldSubPath{
		selector: ConfigMap_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Shards().FieldPath(),
	}
}

func (s ConfigMapPathSelectorMetadataShards) WithValue(value map[string]int64) *ConfigMap_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ConfigMap_FieldSubPathValue)
}

func (s ConfigMapPathSelectorMetadataShards) WithArrayOfValues(values []map[string]int64) *ConfigMap_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ConfigMap_FieldSubPathArrayOfValues)
}

func (ConfigMapPathSelectorMetadataShards) WithKey(key string) ConfigMapMapPathSelectorMetadataShards {
	return ConfigMapMapPathSelectorMetadataShards{key: key}
}

type ConfigMapMapPathSelectorMetadataShards struct {
	key string
}

func (s ConfigMapMapPathSelectorMetadataShards) FieldPath() *ConfigMap_FieldSubPath {
	return &ConfigMap_FieldSubPath{
		selector: ConfigMap_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Shards().WithKey(s.key).FieldPath(),
	}
}

func (s ConfigMapMapPathSelectorMetadataShards) WithValue(value int64) *ConfigMap_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ConfigMap_FieldSubPathValue)
}

func (s ConfigMapMapPathSelectorMetadataShards) WithArrayOfValues(values []int64) *ConfigMap_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ConfigMap_FieldSubPathArrayOfValues)
}

type ConfigMapPathSelectorMetadataSyncing struct{}

func (ConfigMapPathSelectorMetadataSyncing) FieldPath() *ConfigMap_FieldSubPath {
	return &ConfigMap_FieldSubPath{
		selector: ConfigMap_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Syncing().FieldPath(),
	}
}

func (s ConfigMapPathSelectorMetadataSyncing) WithValue(value *meta.SyncingMeta) *ConfigMap_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ConfigMap_FieldSubPathValue)
}

func (s ConfigMapPathSelectorMetadataSyncing) WithArrayOfValues(values []*meta.SyncingMeta) *ConfigMap_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ConfigMap_FieldSubPathArrayOfValues)
}

func (ConfigMapPathSelectorMetadataSyncing) OwningRegion() ConfigMapPathSelectorMetadataSyncingOwningRegion {
	return ConfigMapPathSelectorMetadataSyncingOwningRegion{}
}

func (ConfigMapPathSelectorMetadataSyncing) Regions() ConfigMapPathSelectorMetadataSyncingRegions {
	return ConfigMapPathSelectorMetadataSyncingRegions{}
}

type ConfigMapPathSelectorMetadataSyncingOwningRegion struct{}

func (ConfigMapPathSelectorMetadataSyncingOwningRegion) FieldPath() *ConfigMap_FieldSubPath {
	return &ConfigMap_FieldSubPath{
		selector: ConfigMap_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Syncing().OwningRegion().FieldPath(),
	}
}

func (s ConfigMapPathSelectorMetadataSyncingOwningRegion) WithValue(value string) *ConfigMap_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ConfigMap_FieldSubPathValue)
}

func (s ConfigMapPathSelectorMetadataSyncingOwningRegion) WithArrayOfValues(values []string) *ConfigMap_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ConfigMap_FieldSubPathArrayOfValues)
}

type ConfigMapPathSelectorMetadataSyncingRegions struct{}

func (ConfigMapPathSelectorMetadataSyncingRegions) FieldPath() *ConfigMap_FieldSubPath {
	return &ConfigMap_FieldSubPath{
		selector: ConfigMap_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Syncing().Regions().FieldPath(),
	}
}

func (s ConfigMapPathSelectorMetadataSyncingRegions) WithValue(value []string) *ConfigMap_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ConfigMap_FieldSubPathValue)
}

func (s ConfigMapPathSelectorMetadataSyncingRegions) WithArrayOfValues(values [][]string) *ConfigMap_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ConfigMap_FieldSubPathArrayOfValues)
}

func (s ConfigMapPathSelectorMetadataSyncingRegions) WithItemValue(value string) *ConfigMap_FieldSubPathArrayItemValue {
	return s.FieldPath().WithIArrayItemValue(value).(*ConfigMap_FieldSubPathArrayItemValue)
}

type ConfigMapPathSelectorMetadataLifecycle struct{}

func (ConfigMapPathSelectorMetadataLifecycle) FieldPath() *ConfigMap_FieldSubPath {
	return &ConfigMap_FieldSubPath{
		selector: ConfigMap_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Lifecycle().FieldPath(),
	}
}

func (s ConfigMapPathSelectorMetadataLifecycle) WithValue(value *meta.Lifecycle) *ConfigMap_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ConfigMap_FieldSubPathValue)
}

func (s ConfigMapPathSelectorMetadataLifecycle) WithArrayOfValues(values []*meta.Lifecycle) *ConfigMap_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ConfigMap_FieldSubPathArrayOfValues)
}

func (ConfigMapPathSelectorMetadataLifecycle) State() ConfigMapPathSelectorMetadataLifecycleState {
	return ConfigMapPathSelectorMetadataLifecycleState{}
}

func (ConfigMapPathSelectorMetadataLifecycle) BlockDeletion() ConfigMapPathSelectorMetadataLifecycleBlockDeletion {
	return ConfigMapPathSelectorMetadataLifecycleBlockDeletion{}
}

type ConfigMapPathSelectorMetadataLifecycleState struct{}

func (ConfigMapPathSelectorMetadataLifecycleState) FieldPath() *ConfigMap_FieldSubPath {
	return &ConfigMap_FieldSubPath{
		selector: ConfigMap_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Lifecycle().State().FieldPath(),
	}
}

func (s ConfigMapPathSelectorMetadataLifecycleState) WithValue(value meta.Lifecycle_State) *ConfigMap_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ConfigMap_FieldSubPathValue)
}

func (s ConfigMapPathSelectorMetadataLifecycleState) WithArrayOfValues(values []meta.Lifecycle_State) *ConfigMap_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ConfigMap_FieldSubPathArrayOfValues)
}

type ConfigMapPathSelectorMetadataLifecycleBlockDeletion struct{}

func (ConfigMapPathSelectorMetadataLifecycleBlockDeletion) FieldPath() *ConfigMap_FieldSubPath {
	return &ConfigMap_FieldSubPath{
		selector: ConfigMap_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Lifecycle().BlockDeletion().FieldPath(),
	}
}

func (s ConfigMapPathSelectorMetadataLifecycleBlockDeletion) WithValue(value bool) *ConfigMap_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ConfigMap_FieldSubPathValue)
}

func (s ConfigMapPathSelectorMetadataLifecycleBlockDeletion) WithArrayOfValues(values []bool) *ConfigMap_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ConfigMap_FieldSubPathArrayOfValues)
}

type ConfigMapPathSelectorMetadataServices struct{}

func (ConfigMapPathSelectorMetadataServices) FieldPath() *ConfigMap_FieldSubPath {
	return &ConfigMap_FieldSubPath{
		selector: ConfigMap_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Services().FieldPath(),
	}
}

func (s ConfigMapPathSelectorMetadataServices) WithValue(value *meta.ServicesInfo) *ConfigMap_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ConfigMap_FieldSubPathValue)
}

func (s ConfigMapPathSelectorMetadataServices) WithArrayOfValues(values []*meta.ServicesInfo) *ConfigMap_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ConfigMap_FieldSubPathArrayOfValues)
}

func (ConfigMapPathSelectorMetadataServices) OwningService() ConfigMapPathSelectorMetadataServicesOwningService {
	return ConfigMapPathSelectorMetadataServicesOwningService{}
}

func (ConfigMapPathSelectorMetadataServices) AllowedServices() ConfigMapPathSelectorMetadataServicesAllowedServices {
	return ConfigMapPathSelectorMetadataServicesAllowedServices{}
}

type ConfigMapPathSelectorMetadataServicesOwningService struct{}

func (ConfigMapPathSelectorMetadataServicesOwningService) FieldPath() *ConfigMap_FieldSubPath {
	return &ConfigMap_FieldSubPath{
		selector: ConfigMap_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Services().OwningService().FieldPath(),
	}
}

func (s ConfigMapPathSelectorMetadataServicesOwningService) WithValue(value string) *ConfigMap_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ConfigMap_FieldSubPathValue)
}

func (s ConfigMapPathSelectorMetadataServicesOwningService) WithArrayOfValues(values []string) *ConfigMap_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ConfigMap_FieldSubPathArrayOfValues)
}

type ConfigMapPathSelectorMetadataServicesAllowedServices struct{}

func (ConfigMapPathSelectorMetadataServicesAllowedServices) FieldPath() *ConfigMap_FieldSubPath {
	return &ConfigMap_FieldSubPath{
		selector: ConfigMap_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Services().AllowedServices().FieldPath(),
	}
}

func (s ConfigMapPathSelectorMetadataServicesAllowedServices) WithValue(value []string) *ConfigMap_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ConfigMap_FieldSubPathValue)
}

func (s ConfigMapPathSelectorMetadataServicesAllowedServices) WithArrayOfValues(values [][]string) *ConfigMap_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ConfigMap_FieldSubPathArrayOfValues)
}

func (s ConfigMapPathSelectorMetadataServicesAllowedServices) WithItemValue(value string) *ConfigMap_FieldSubPathArrayItemValue {
	return s.FieldPath().WithIArrayItemValue(value).(*ConfigMap_FieldSubPathArrayItemValue)
}

type ConfigMapPathSelectorData struct{}

func (ConfigMapPathSelectorData) FieldPath() *ConfigMap_FieldTerminalPath {
	return &ConfigMap_FieldTerminalPath{selector: ConfigMap_FieldPathSelectorData}
}

func (s ConfigMapPathSelectorData) WithValue(value map[string]string) *ConfigMap_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*ConfigMap_FieldTerminalPathValue)
}

func (s ConfigMapPathSelectorData) WithArrayOfValues(values []map[string]string) *ConfigMap_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ConfigMap_FieldTerminalPathArrayOfValues)
}

func (ConfigMapPathSelectorData) WithKey(key string) ConfigMapMapPathSelectorData {
	return ConfigMapMapPathSelectorData{key: key}
}

type ConfigMapMapPathSelectorData struct {
	key string
}

func (s ConfigMapMapPathSelectorData) FieldPath() *ConfigMap_FieldPathMap {
	return &ConfigMap_FieldPathMap{selector: ConfigMap_FieldPathSelectorData, key: s.key}
}

func (s ConfigMapMapPathSelectorData) WithValue(value string) *ConfigMap_FieldPathMapValue {
	return s.FieldPath().WithIValue(value).(*ConfigMap_FieldPathMapValue)
}

func (s ConfigMapMapPathSelectorData) WithArrayOfValues(values []string) *ConfigMap_FieldPathMapArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ConfigMap_FieldPathMapArrayOfValues)
}

type ConfigMapPathSelectorBinaryData struct{}

func (ConfigMapPathSelectorBinaryData) FieldPath() *ConfigMap_FieldTerminalPath {
	return &ConfigMap_FieldTerminalPath{selector: ConfigMap_FieldPathSelectorBinaryData}
}

func (s ConfigMapPathSelectorBinaryData) WithValue(value map[string][]byte) *ConfigMap_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*ConfigMap_FieldTerminalPathValue)
}

func (s ConfigMapPathSelectorBinaryData) WithArrayOfValues(values []map[string][]byte) *ConfigMap_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ConfigMap_FieldTerminalPathArrayOfValues)
}

func (ConfigMapPathSelectorBinaryData) WithKey(key string) ConfigMapMapPathSelectorBinaryData {
	return ConfigMapMapPathSelectorBinaryData{key: key}
}

type ConfigMapMapPathSelectorBinaryData struct {
	key string
}

func (s ConfigMapMapPathSelectorBinaryData) FieldPath() *ConfigMap_FieldPathMap {
	return &ConfigMap_FieldPathMap{selector: ConfigMap_FieldPathSelectorBinaryData, key: s.key}
}

func (s ConfigMapMapPathSelectorBinaryData) WithValue(value []byte) *ConfigMap_FieldPathMapValue {
	return s.FieldPath().WithIValue(value).(*ConfigMap_FieldPathMapValue)
}

func (s ConfigMapMapPathSelectorBinaryData) WithArrayOfValues(values [][]byte) *ConfigMap_FieldPathMapArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ConfigMap_FieldPathMapArrayOfValues)
}
