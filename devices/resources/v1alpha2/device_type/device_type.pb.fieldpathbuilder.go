// Code generated by protoc-gen-goten-object
// File: edgelq/devices/proto/v1alpha2/device_type.proto
// DO NOT EDIT!!!

package device_type

// proto imports
import (
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
)

// make sure we're using proto imports
var (
	_ = &ntt_meta.Meta{}
	_ = &timestamp.Timestamp{}
)

type DeviceTypeFieldPathBuilder struct{}

func NewDeviceTypeFieldPathBuilder() DeviceTypeFieldPathBuilder {
	return DeviceTypeFieldPathBuilder{}
}
func (DeviceTypeFieldPathBuilder) Name() DeviceTypePathSelectorName {
	return DeviceTypePathSelectorName{}
}
func (DeviceTypeFieldPathBuilder) Metadata() DeviceTypePathSelectorMetadata {
	return DeviceTypePathSelectorMetadata{}
}
func (DeviceTypeFieldPathBuilder) DisplayName() DeviceTypePathSelectorDisplayName {
	return DeviceTypePathSelectorDisplayName{}
}
func (DeviceTypeFieldPathBuilder) Hardware() DeviceTypePathSelectorHardware {
	return DeviceTypePathSelectorHardware{}
}
func (DeviceTypeFieldPathBuilder) Architecture() DeviceTypePathSelectorArchitecture {
	return DeviceTypePathSelectorArchitecture{}
}
func (DeviceTypeFieldPathBuilder) Description() DeviceTypePathSelectorDescription {
	return DeviceTypePathSelectorDescription{}
}

type DeviceTypePathSelectorName struct{}

func (DeviceTypePathSelectorName) FieldPath() *DeviceType_FieldTerminalPath {
	return &DeviceType_FieldTerminalPath{selector: DeviceType_FieldPathSelectorName}
}

func (s DeviceTypePathSelectorName) WithValue(value *Name) *DeviceType_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*DeviceType_FieldTerminalPathValue)
}

func (s DeviceTypePathSelectorName) WithArrayOfValues(values []*Name) *DeviceType_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*DeviceType_FieldTerminalPathArrayOfValues)
}

type DeviceTypePathSelectorMetadata struct{}

func (DeviceTypePathSelectorMetadata) FieldPath() *DeviceType_FieldTerminalPath {
	return &DeviceType_FieldTerminalPath{selector: DeviceType_FieldPathSelectorMetadata}
}

func (s DeviceTypePathSelectorMetadata) WithValue(value *ntt_meta.Meta) *DeviceType_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*DeviceType_FieldTerminalPathValue)
}

func (s DeviceTypePathSelectorMetadata) WithArrayOfValues(values []*ntt_meta.Meta) *DeviceType_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*DeviceType_FieldTerminalPathArrayOfValues)
}

func (DeviceTypePathSelectorMetadata) WithSubPath(subPath ntt_meta.Meta_FieldPath) *DeviceType_FieldSubPath {
	return &DeviceType_FieldSubPath{selector: DeviceType_FieldPathSelectorMetadata, subPath: subPath}
}

func (s DeviceTypePathSelectorMetadata) WithSubValue(subPathValue ntt_meta.Meta_FieldPathValue) *DeviceType_FieldSubPathValue {
	return &DeviceType_FieldSubPathValue{DeviceType_FieldPath: s.WithSubPath(subPathValue), subPathValue: subPathValue}
}

func (s DeviceTypePathSelectorMetadata) WithSubArrayOfValues(subPathArrayOfValues ntt_meta.Meta_FieldPathArrayOfValues) *DeviceType_FieldSubPathArrayOfValues {
	return &DeviceType_FieldSubPathArrayOfValues{DeviceType_FieldPath: s.WithSubPath(subPathArrayOfValues), subPathArrayOfValues: subPathArrayOfValues}
}

func (s DeviceTypePathSelectorMetadata) WithSubArrayItemValue(subPathArrayItemValue ntt_meta.Meta_FieldPathArrayItemValue) *DeviceType_FieldSubPathArrayItemValue {
	return &DeviceType_FieldSubPathArrayItemValue{DeviceType_FieldPath: s.WithSubPath(subPathArrayItemValue), subPathItemValue: subPathArrayItemValue}
}

func (DeviceTypePathSelectorMetadata) CreateTime() DeviceTypePathSelectorMetadataCreateTime {
	return DeviceTypePathSelectorMetadataCreateTime{}
}

func (DeviceTypePathSelectorMetadata) UpdateTime() DeviceTypePathSelectorMetadataUpdateTime {
	return DeviceTypePathSelectorMetadataUpdateTime{}
}

func (DeviceTypePathSelectorMetadata) DeleteTime() DeviceTypePathSelectorMetadataDeleteTime {
	return DeviceTypePathSelectorMetadataDeleteTime{}
}

func (DeviceTypePathSelectorMetadata) Uuid() DeviceTypePathSelectorMetadataUuid {
	return DeviceTypePathSelectorMetadataUuid{}
}

func (DeviceTypePathSelectorMetadata) Tags() DeviceTypePathSelectorMetadataTags {
	return DeviceTypePathSelectorMetadataTags{}
}

func (DeviceTypePathSelectorMetadata) Labels() DeviceTypePathSelectorMetadataLabels {
	return DeviceTypePathSelectorMetadataLabels{}
}

func (DeviceTypePathSelectorMetadata) Annotations() DeviceTypePathSelectorMetadataAnnotations {
	return DeviceTypePathSelectorMetadataAnnotations{}
}

func (DeviceTypePathSelectorMetadata) Generation() DeviceTypePathSelectorMetadataGeneration {
	return DeviceTypePathSelectorMetadataGeneration{}
}

func (DeviceTypePathSelectorMetadata) ResourceVersion() DeviceTypePathSelectorMetadataResourceVersion {
	return DeviceTypePathSelectorMetadataResourceVersion{}
}

func (DeviceTypePathSelectorMetadata) OwnerReferences() DeviceTypePathSelectorMetadataOwnerReferences {
	return DeviceTypePathSelectorMetadataOwnerReferences{}
}

func (DeviceTypePathSelectorMetadata) Shards() DeviceTypePathSelectorMetadataShards {
	return DeviceTypePathSelectorMetadataShards{}
}

func (DeviceTypePathSelectorMetadata) Syncing() DeviceTypePathSelectorMetadataSyncing {
	return DeviceTypePathSelectorMetadataSyncing{}
}

func (DeviceTypePathSelectorMetadata) Lifecycle() DeviceTypePathSelectorMetadataLifecycle {
	return DeviceTypePathSelectorMetadataLifecycle{}
}

type DeviceTypePathSelectorMetadataCreateTime struct{}

func (DeviceTypePathSelectorMetadataCreateTime) FieldPath() *DeviceType_FieldSubPath {
	return &DeviceType_FieldSubPath{
		selector: DeviceType_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().CreateTime().FieldPath(),
	}
}

func (s DeviceTypePathSelectorMetadataCreateTime) WithValue(value *timestamp.Timestamp) *DeviceType_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*DeviceType_FieldSubPathValue)
}

func (s DeviceTypePathSelectorMetadataCreateTime) WithArrayOfValues(values []*timestamp.Timestamp) *DeviceType_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*DeviceType_FieldSubPathArrayOfValues)
}

type DeviceTypePathSelectorMetadataUpdateTime struct{}

func (DeviceTypePathSelectorMetadataUpdateTime) FieldPath() *DeviceType_FieldSubPath {
	return &DeviceType_FieldSubPath{
		selector: DeviceType_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().UpdateTime().FieldPath(),
	}
}

func (s DeviceTypePathSelectorMetadataUpdateTime) WithValue(value *timestamp.Timestamp) *DeviceType_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*DeviceType_FieldSubPathValue)
}

func (s DeviceTypePathSelectorMetadataUpdateTime) WithArrayOfValues(values []*timestamp.Timestamp) *DeviceType_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*DeviceType_FieldSubPathArrayOfValues)
}

type DeviceTypePathSelectorMetadataDeleteTime struct{}

func (DeviceTypePathSelectorMetadataDeleteTime) FieldPath() *DeviceType_FieldSubPath {
	return &DeviceType_FieldSubPath{
		selector: DeviceType_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().DeleteTime().FieldPath(),
	}
}

func (s DeviceTypePathSelectorMetadataDeleteTime) WithValue(value *timestamp.Timestamp) *DeviceType_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*DeviceType_FieldSubPathValue)
}

func (s DeviceTypePathSelectorMetadataDeleteTime) WithArrayOfValues(values []*timestamp.Timestamp) *DeviceType_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*DeviceType_FieldSubPathArrayOfValues)
}

type DeviceTypePathSelectorMetadataUuid struct{}

func (DeviceTypePathSelectorMetadataUuid) FieldPath() *DeviceType_FieldSubPath {
	return &DeviceType_FieldSubPath{
		selector: DeviceType_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Uuid().FieldPath(),
	}
}

func (s DeviceTypePathSelectorMetadataUuid) WithValue(value string) *DeviceType_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*DeviceType_FieldSubPathValue)
}

func (s DeviceTypePathSelectorMetadataUuid) WithArrayOfValues(values []string) *DeviceType_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*DeviceType_FieldSubPathArrayOfValues)
}

type DeviceTypePathSelectorMetadataTags struct{}

func (DeviceTypePathSelectorMetadataTags) FieldPath() *DeviceType_FieldSubPath {
	return &DeviceType_FieldSubPath{
		selector: DeviceType_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Tags().FieldPath(),
	}
}

func (s DeviceTypePathSelectorMetadataTags) WithValue(value []string) *DeviceType_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*DeviceType_FieldSubPathValue)
}

func (s DeviceTypePathSelectorMetadataTags) WithArrayOfValues(values [][]string) *DeviceType_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*DeviceType_FieldSubPathArrayOfValues)
}

func (s DeviceTypePathSelectorMetadataTags) WithItemValue(value string) *DeviceType_FieldSubPathArrayItemValue {
	return s.FieldPath().WithIArrayItemValue(value).(*DeviceType_FieldSubPathArrayItemValue)
}

type DeviceTypePathSelectorMetadataLabels struct{}

func (DeviceTypePathSelectorMetadataLabels) FieldPath() *DeviceType_FieldSubPath {
	return &DeviceType_FieldSubPath{
		selector: DeviceType_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Labels().FieldPath(),
	}
}

func (s DeviceTypePathSelectorMetadataLabels) WithValue(value map[string]string) *DeviceType_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*DeviceType_FieldSubPathValue)
}

func (s DeviceTypePathSelectorMetadataLabels) WithArrayOfValues(values []map[string]string) *DeviceType_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*DeviceType_FieldSubPathArrayOfValues)
}

func (DeviceTypePathSelectorMetadataLabels) WithKey(key string) DeviceTypeMapPathSelectorMetadataLabels {
	return DeviceTypeMapPathSelectorMetadataLabels{key: key}
}

type DeviceTypeMapPathSelectorMetadataLabels struct {
	key string
}

func (s DeviceTypeMapPathSelectorMetadataLabels) FieldPath() *DeviceType_FieldSubPath {
	return &DeviceType_FieldSubPath{
		selector: DeviceType_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Labels().WithKey(s.key).FieldPath(),
	}
}

func (s DeviceTypeMapPathSelectorMetadataLabels) WithValue(value string) *DeviceType_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*DeviceType_FieldSubPathValue)
}

func (s DeviceTypeMapPathSelectorMetadataLabels) WithArrayOfValues(values []string) *DeviceType_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*DeviceType_FieldSubPathArrayOfValues)
}

type DeviceTypePathSelectorMetadataAnnotations struct{}

func (DeviceTypePathSelectorMetadataAnnotations) FieldPath() *DeviceType_FieldSubPath {
	return &DeviceType_FieldSubPath{
		selector: DeviceType_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Annotations().FieldPath(),
	}
}

func (s DeviceTypePathSelectorMetadataAnnotations) WithValue(value map[string]string) *DeviceType_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*DeviceType_FieldSubPathValue)
}

func (s DeviceTypePathSelectorMetadataAnnotations) WithArrayOfValues(values []map[string]string) *DeviceType_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*DeviceType_FieldSubPathArrayOfValues)
}

func (DeviceTypePathSelectorMetadataAnnotations) WithKey(key string) DeviceTypeMapPathSelectorMetadataAnnotations {
	return DeviceTypeMapPathSelectorMetadataAnnotations{key: key}
}

type DeviceTypeMapPathSelectorMetadataAnnotations struct {
	key string
}

func (s DeviceTypeMapPathSelectorMetadataAnnotations) FieldPath() *DeviceType_FieldSubPath {
	return &DeviceType_FieldSubPath{
		selector: DeviceType_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Annotations().WithKey(s.key).FieldPath(),
	}
}

func (s DeviceTypeMapPathSelectorMetadataAnnotations) WithValue(value string) *DeviceType_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*DeviceType_FieldSubPathValue)
}

func (s DeviceTypeMapPathSelectorMetadataAnnotations) WithArrayOfValues(values []string) *DeviceType_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*DeviceType_FieldSubPathArrayOfValues)
}

type DeviceTypePathSelectorMetadataGeneration struct{}

func (DeviceTypePathSelectorMetadataGeneration) FieldPath() *DeviceType_FieldSubPath {
	return &DeviceType_FieldSubPath{
		selector: DeviceType_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Generation().FieldPath(),
	}
}

func (s DeviceTypePathSelectorMetadataGeneration) WithValue(value int64) *DeviceType_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*DeviceType_FieldSubPathValue)
}

func (s DeviceTypePathSelectorMetadataGeneration) WithArrayOfValues(values []int64) *DeviceType_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*DeviceType_FieldSubPathArrayOfValues)
}

type DeviceTypePathSelectorMetadataResourceVersion struct{}

func (DeviceTypePathSelectorMetadataResourceVersion) FieldPath() *DeviceType_FieldSubPath {
	return &DeviceType_FieldSubPath{
		selector: DeviceType_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().ResourceVersion().FieldPath(),
	}
}

func (s DeviceTypePathSelectorMetadataResourceVersion) WithValue(value string) *DeviceType_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*DeviceType_FieldSubPathValue)
}

func (s DeviceTypePathSelectorMetadataResourceVersion) WithArrayOfValues(values []string) *DeviceType_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*DeviceType_FieldSubPathArrayOfValues)
}

type DeviceTypePathSelectorMetadataOwnerReferences struct{}

func (DeviceTypePathSelectorMetadataOwnerReferences) FieldPath() *DeviceType_FieldSubPath {
	return &DeviceType_FieldSubPath{
		selector: DeviceType_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().FieldPath(),
	}
}

func (s DeviceTypePathSelectorMetadataOwnerReferences) WithValue(value []*ntt_meta.OwnerReference) *DeviceType_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*DeviceType_FieldSubPathValue)
}

func (s DeviceTypePathSelectorMetadataOwnerReferences) WithArrayOfValues(values [][]*ntt_meta.OwnerReference) *DeviceType_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*DeviceType_FieldSubPathArrayOfValues)
}

func (s DeviceTypePathSelectorMetadataOwnerReferences) WithItemValue(value *ntt_meta.OwnerReference) *DeviceType_FieldSubPathArrayItemValue {
	return s.FieldPath().WithIArrayItemValue(value).(*DeviceType_FieldSubPathArrayItemValue)
}

func (DeviceTypePathSelectorMetadataOwnerReferences) Kind() DeviceTypePathSelectorMetadataOwnerReferencesKind {
	return DeviceTypePathSelectorMetadataOwnerReferencesKind{}
}

func (DeviceTypePathSelectorMetadataOwnerReferences) Version() DeviceTypePathSelectorMetadataOwnerReferencesVersion {
	return DeviceTypePathSelectorMetadataOwnerReferencesVersion{}
}

func (DeviceTypePathSelectorMetadataOwnerReferences) Name() DeviceTypePathSelectorMetadataOwnerReferencesName {
	return DeviceTypePathSelectorMetadataOwnerReferencesName{}
}

func (DeviceTypePathSelectorMetadataOwnerReferences) Region() DeviceTypePathSelectorMetadataOwnerReferencesRegion {
	return DeviceTypePathSelectorMetadataOwnerReferencesRegion{}
}

func (DeviceTypePathSelectorMetadataOwnerReferences) Controller() DeviceTypePathSelectorMetadataOwnerReferencesController {
	return DeviceTypePathSelectorMetadataOwnerReferencesController{}
}

func (DeviceTypePathSelectorMetadataOwnerReferences) BlockOwnerDeletion() DeviceTypePathSelectorMetadataOwnerReferencesBlockOwnerDeletion {
	return DeviceTypePathSelectorMetadataOwnerReferencesBlockOwnerDeletion{}
}

func (DeviceTypePathSelectorMetadataOwnerReferences) RequiresOwnerReference() DeviceTypePathSelectorMetadataOwnerReferencesRequiresOwnerReference {
	return DeviceTypePathSelectorMetadataOwnerReferencesRequiresOwnerReference{}
}

type DeviceTypePathSelectorMetadataOwnerReferencesKind struct{}

func (DeviceTypePathSelectorMetadataOwnerReferencesKind) FieldPath() *DeviceType_FieldSubPath {
	return &DeviceType_FieldSubPath{
		selector: DeviceType_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().Kind().FieldPath(),
	}
}

func (s DeviceTypePathSelectorMetadataOwnerReferencesKind) WithValue(value string) *DeviceType_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*DeviceType_FieldSubPathValue)
}

func (s DeviceTypePathSelectorMetadataOwnerReferencesKind) WithArrayOfValues(values []string) *DeviceType_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*DeviceType_FieldSubPathArrayOfValues)
}

type DeviceTypePathSelectorMetadataOwnerReferencesVersion struct{}

func (DeviceTypePathSelectorMetadataOwnerReferencesVersion) FieldPath() *DeviceType_FieldSubPath {
	return &DeviceType_FieldSubPath{
		selector: DeviceType_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().Version().FieldPath(),
	}
}

func (s DeviceTypePathSelectorMetadataOwnerReferencesVersion) WithValue(value string) *DeviceType_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*DeviceType_FieldSubPathValue)
}

func (s DeviceTypePathSelectorMetadataOwnerReferencesVersion) WithArrayOfValues(values []string) *DeviceType_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*DeviceType_FieldSubPathArrayOfValues)
}

type DeviceTypePathSelectorMetadataOwnerReferencesName struct{}

func (DeviceTypePathSelectorMetadataOwnerReferencesName) FieldPath() *DeviceType_FieldSubPath {
	return &DeviceType_FieldSubPath{
		selector: DeviceType_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().Name().FieldPath(),
	}
}

func (s DeviceTypePathSelectorMetadataOwnerReferencesName) WithValue(value string) *DeviceType_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*DeviceType_FieldSubPathValue)
}

func (s DeviceTypePathSelectorMetadataOwnerReferencesName) WithArrayOfValues(values []string) *DeviceType_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*DeviceType_FieldSubPathArrayOfValues)
}

type DeviceTypePathSelectorMetadataOwnerReferencesRegion struct{}

func (DeviceTypePathSelectorMetadataOwnerReferencesRegion) FieldPath() *DeviceType_FieldSubPath {
	return &DeviceType_FieldSubPath{
		selector: DeviceType_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().Region().FieldPath(),
	}
}

func (s DeviceTypePathSelectorMetadataOwnerReferencesRegion) WithValue(value string) *DeviceType_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*DeviceType_FieldSubPathValue)
}

func (s DeviceTypePathSelectorMetadataOwnerReferencesRegion) WithArrayOfValues(values []string) *DeviceType_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*DeviceType_FieldSubPathArrayOfValues)
}

type DeviceTypePathSelectorMetadataOwnerReferencesController struct{}

func (DeviceTypePathSelectorMetadataOwnerReferencesController) FieldPath() *DeviceType_FieldSubPath {
	return &DeviceType_FieldSubPath{
		selector: DeviceType_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().Controller().FieldPath(),
	}
}

func (s DeviceTypePathSelectorMetadataOwnerReferencesController) WithValue(value bool) *DeviceType_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*DeviceType_FieldSubPathValue)
}

func (s DeviceTypePathSelectorMetadataOwnerReferencesController) WithArrayOfValues(values []bool) *DeviceType_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*DeviceType_FieldSubPathArrayOfValues)
}

type DeviceTypePathSelectorMetadataOwnerReferencesBlockOwnerDeletion struct{}

func (DeviceTypePathSelectorMetadataOwnerReferencesBlockOwnerDeletion) FieldPath() *DeviceType_FieldSubPath {
	return &DeviceType_FieldSubPath{
		selector: DeviceType_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().BlockOwnerDeletion().FieldPath(),
	}
}

func (s DeviceTypePathSelectorMetadataOwnerReferencesBlockOwnerDeletion) WithValue(value bool) *DeviceType_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*DeviceType_FieldSubPathValue)
}

func (s DeviceTypePathSelectorMetadataOwnerReferencesBlockOwnerDeletion) WithArrayOfValues(values []bool) *DeviceType_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*DeviceType_FieldSubPathArrayOfValues)
}

type DeviceTypePathSelectorMetadataOwnerReferencesRequiresOwnerReference struct{}

func (DeviceTypePathSelectorMetadataOwnerReferencesRequiresOwnerReference) FieldPath() *DeviceType_FieldSubPath {
	return &DeviceType_FieldSubPath{
		selector: DeviceType_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().RequiresOwnerReference().FieldPath(),
	}
}

func (s DeviceTypePathSelectorMetadataOwnerReferencesRequiresOwnerReference) WithValue(value bool) *DeviceType_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*DeviceType_FieldSubPathValue)
}

func (s DeviceTypePathSelectorMetadataOwnerReferencesRequiresOwnerReference) WithArrayOfValues(values []bool) *DeviceType_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*DeviceType_FieldSubPathArrayOfValues)
}

type DeviceTypePathSelectorMetadataShards struct{}

func (DeviceTypePathSelectorMetadataShards) FieldPath() *DeviceType_FieldSubPath {
	return &DeviceType_FieldSubPath{
		selector: DeviceType_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Shards().FieldPath(),
	}
}

func (s DeviceTypePathSelectorMetadataShards) WithValue(value map[string]int64) *DeviceType_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*DeviceType_FieldSubPathValue)
}

func (s DeviceTypePathSelectorMetadataShards) WithArrayOfValues(values []map[string]int64) *DeviceType_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*DeviceType_FieldSubPathArrayOfValues)
}

func (DeviceTypePathSelectorMetadataShards) WithKey(key string) DeviceTypeMapPathSelectorMetadataShards {
	return DeviceTypeMapPathSelectorMetadataShards{key: key}
}

type DeviceTypeMapPathSelectorMetadataShards struct {
	key string
}

func (s DeviceTypeMapPathSelectorMetadataShards) FieldPath() *DeviceType_FieldSubPath {
	return &DeviceType_FieldSubPath{
		selector: DeviceType_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Shards().WithKey(s.key).FieldPath(),
	}
}

func (s DeviceTypeMapPathSelectorMetadataShards) WithValue(value int64) *DeviceType_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*DeviceType_FieldSubPathValue)
}

func (s DeviceTypeMapPathSelectorMetadataShards) WithArrayOfValues(values []int64) *DeviceType_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*DeviceType_FieldSubPathArrayOfValues)
}

type DeviceTypePathSelectorMetadataSyncing struct{}

func (DeviceTypePathSelectorMetadataSyncing) FieldPath() *DeviceType_FieldSubPath {
	return &DeviceType_FieldSubPath{
		selector: DeviceType_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Syncing().FieldPath(),
	}
}

func (s DeviceTypePathSelectorMetadataSyncing) WithValue(value *ntt_meta.SyncingMeta) *DeviceType_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*DeviceType_FieldSubPathValue)
}

func (s DeviceTypePathSelectorMetadataSyncing) WithArrayOfValues(values []*ntt_meta.SyncingMeta) *DeviceType_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*DeviceType_FieldSubPathArrayOfValues)
}

func (DeviceTypePathSelectorMetadataSyncing) OwningRegion() DeviceTypePathSelectorMetadataSyncingOwningRegion {
	return DeviceTypePathSelectorMetadataSyncingOwningRegion{}
}

func (DeviceTypePathSelectorMetadataSyncing) Regions() DeviceTypePathSelectorMetadataSyncingRegions {
	return DeviceTypePathSelectorMetadataSyncingRegions{}
}

type DeviceTypePathSelectorMetadataSyncingOwningRegion struct{}

func (DeviceTypePathSelectorMetadataSyncingOwningRegion) FieldPath() *DeviceType_FieldSubPath {
	return &DeviceType_FieldSubPath{
		selector: DeviceType_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Syncing().OwningRegion().FieldPath(),
	}
}

func (s DeviceTypePathSelectorMetadataSyncingOwningRegion) WithValue(value string) *DeviceType_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*DeviceType_FieldSubPathValue)
}

func (s DeviceTypePathSelectorMetadataSyncingOwningRegion) WithArrayOfValues(values []string) *DeviceType_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*DeviceType_FieldSubPathArrayOfValues)
}

type DeviceTypePathSelectorMetadataSyncingRegions struct{}

func (DeviceTypePathSelectorMetadataSyncingRegions) FieldPath() *DeviceType_FieldSubPath {
	return &DeviceType_FieldSubPath{
		selector: DeviceType_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Syncing().Regions().FieldPath(),
	}
}

func (s DeviceTypePathSelectorMetadataSyncingRegions) WithValue(value []string) *DeviceType_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*DeviceType_FieldSubPathValue)
}

func (s DeviceTypePathSelectorMetadataSyncingRegions) WithArrayOfValues(values [][]string) *DeviceType_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*DeviceType_FieldSubPathArrayOfValues)
}

func (s DeviceTypePathSelectorMetadataSyncingRegions) WithItemValue(value string) *DeviceType_FieldSubPathArrayItemValue {
	return s.FieldPath().WithIArrayItemValue(value).(*DeviceType_FieldSubPathArrayItemValue)
}

type DeviceTypePathSelectorMetadataLifecycle struct{}

func (DeviceTypePathSelectorMetadataLifecycle) FieldPath() *DeviceType_FieldSubPath {
	return &DeviceType_FieldSubPath{
		selector: DeviceType_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Lifecycle().FieldPath(),
	}
}

func (s DeviceTypePathSelectorMetadataLifecycle) WithValue(value *ntt_meta.Lifecycle) *DeviceType_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*DeviceType_FieldSubPathValue)
}

func (s DeviceTypePathSelectorMetadataLifecycle) WithArrayOfValues(values []*ntt_meta.Lifecycle) *DeviceType_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*DeviceType_FieldSubPathArrayOfValues)
}

func (DeviceTypePathSelectorMetadataLifecycle) State() DeviceTypePathSelectorMetadataLifecycleState {
	return DeviceTypePathSelectorMetadataLifecycleState{}
}

func (DeviceTypePathSelectorMetadataLifecycle) BlockDeletion() DeviceTypePathSelectorMetadataLifecycleBlockDeletion {
	return DeviceTypePathSelectorMetadataLifecycleBlockDeletion{}
}

type DeviceTypePathSelectorMetadataLifecycleState struct{}

func (DeviceTypePathSelectorMetadataLifecycleState) FieldPath() *DeviceType_FieldSubPath {
	return &DeviceType_FieldSubPath{
		selector: DeviceType_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Lifecycle().State().FieldPath(),
	}
}

func (s DeviceTypePathSelectorMetadataLifecycleState) WithValue(value ntt_meta.Lifecycle_State) *DeviceType_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*DeviceType_FieldSubPathValue)
}

func (s DeviceTypePathSelectorMetadataLifecycleState) WithArrayOfValues(values []ntt_meta.Lifecycle_State) *DeviceType_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*DeviceType_FieldSubPathArrayOfValues)
}

type DeviceTypePathSelectorMetadataLifecycleBlockDeletion struct{}

func (DeviceTypePathSelectorMetadataLifecycleBlockDeletion) FieldPath() *DeviceType_FieldSubPath {
	return &DeviceType_FieldSubPath{
		selector: DeviceType_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Lifecycle().BlockDeletion().FieldPath(),
	}
}

func (s DeviceTypePathSelectorMetadataLifecycleBlockDeletion) WithValue(value bool) *DeviceType_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*DeviceType_FieldSubPathValue)
}

func (s DeviceTypePathSelectorMetadataLifecycleBlockDeletion) WithArrayOfValues(values []bool) *DeviceType_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*DeviceType_FieldSubPathArrayOfValues)
}

type DeviceTypePathSelectorDisplayName struct{}

func (DeviceTypePathSelectorDisplayName) FieldPath() *DeviceType_FieldTerminalPath {
	return &DeviceType_FieldTerminalPath{selector: DeviceType_FieldPathSelectorDisplayName}
}

func (s DeviceTypePathSelectorDisplayName) WithValue(value string) *DeviceType_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*DeviceType_FieldTerminalPathValue)
}

func (s DeviceTypePathSelectorDisplayName) WithArrayOfValues(values []string) *DeviceType_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*DeviceType_FieldTerminalPathArrayOfValues)
}

type DeviceTypePathSelectorHardware struct{}

func (DeviceTypePathSelectorHardware) FieldPath() *DeviceType_FieldTerminalPath {
	return &DeviceType_FieldTerminalPath{selector: DeviceType_FieldPathSelectorHardware}
}

func (s DeviceTypePathSelectorHardware) WithValue(value DeviceType_Platform) *DeviceType_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*DeviceType_FieldTerminalPathValue)
}

func (s DeviceTypePathSelectorHardware) WithArrayOfValues(values []DeviceType_Platform) *DeviceType_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*DeviceType_FieldTerminalPathArrayOfValues)
}

type DeviceTypePathSelectorArchitecture struct{}

func (DeviceTypePathSelectorArchitecture) FieldPath() *DeviceType_FieldTerminalPath {
	return &DeviceType_FieldTerminalPath{selector: DeviceType_FieldPathSelectorArchitecture}
}

func (s DeviceTypePathSelectorArchitecture) WithValue(value DeviceType_Architecture) *DeviceType_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*DeviceType_FieldTerminalPathValue)
}

func (s DeviceTypePathSelectorArchitecture) WithArrayOfValues(values []DeviceType_Architecture) *DeviceType_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*DeviceType_FieldTerminalPathArrayOfValues)
}

type DeviceTypePathSelectorDescription struct{}

func (DeviceTypePathSelectorDescription) FieldPath() *DeviceType_FieldTerminalPath {
	return &DeviceType_FieldTerminalPath{selector: DeviceType_FieldPathSelectorDescription}
}

func (s DeviceTypePathSelectorDescription) WithValue(value string) *DeviceType_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*DeviceType_FieldTerminalPathValue)
}

func (s DeviceTypePathSelectorDescription) WithArrayOfValues(values []string) *DeviceType_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*DeviceType_FieldTerminalPathArrayOfValues)
}