// Code generated by protoc-gen-goten-object
// File: edgelq/iam/proto/v1/permission.proto
// DO NOT EDIT!!!

package permission

// proto imports
import (
	meta_service "github.com/cloudwan/goten-sdk/meta-service/resources/v1/service"
	meta "github.com/cloudwan/goten-sdk/types/meta"
	multi_region_policy "github.com/cloudwan/goten-sdk/types/multi_region_policy"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

// make sure we're using proto imports
var (
	_ = &timestamppb.Timestamp{}
	_ = &meta_service.Service{}
	_ = &meta.Meta{}
	_ = &multi_region_policy.MultiRegionPolicy{}
)

type PermissionFieldPathBuilder struct{}

func NewPermissionFieldPathBuilder() PermissionFieldPathBuilder {
	return PermissionFieldPathBuilder{}
}
func (PermissionFieldPathBuilder) Name() PermissionPathSelectorName {
	return PermissionPathSelectorName{}
}
func (PermissionFieldPathBuilder) Metadata() PermissionPathSelectorMetadata {
	return PermissionPathSelectorMetadata{}
}
func (PermissionFieldPathBuilder) Title() PermissionPathSelectorTitle {
	return PermissionPathSelectorTitle{}
}
func (PermissionFieldPathBuilder) Description() PermissionPathSelectorDescription {
	return PermissionPathSelectorDescription{}
}

type PermissionPathSelectorName struct{}

func (PermissionPathSelectorName) FieldPath() *Permission_FieldTerminalPath {
	return &Permission_FieldTerminalPath{selector: Permission_FieldPathSelectorName}
}

func (s PermissionPathSelectorName) WithValue(value *Name) *Permission_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*Permission_FieldTerminalPathValue)
}

func (s PermissionPathSelectorName) WithArrayOfValues(values []*Name) *Permission_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Permission_FieldTerminalPathArrayOfValues)
}

type PermissionPathSelectorMetadata struct{}

func (PermissionPathSelectorMetadata) FieldPath() *Permission_FieldTerminalPath {
	return &Permission_FieldTerminalPath{selector: Permission_FieldPathSelectorMetadata}
}

func (s PermissionPathSelectorMetadata) WithValue(value *meta.Meta) *Permission_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*Permission_FieldTerminalPathValue)
}

func (s PermissionPathSelectorMetadata) WithArrayOfValues(values []*meta.Meta) *Permission_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Permission_FieldTerminalPathArrayOfValues)
}

func (PermissionPathSelectorMetadata) WithSubPath(subPath meta.Meta_FieldPath) *Permission_FieldSubPath {
	return &Permission_FieldSubPath{selector: Permission_FieldPathSelectorMetadata, subPath: subPath}
}

func (s PermissionPathSelectorMetadata) WithSubValue(subPathValue meta.Meta_FieldPathValue) *Permission_FieldSubPathValue {
	return &Permission_FieldSubPathValue{Permission_FieldPath: s.WithSubPath(subPathValue), subPathValue: subPathValue}
}

func (s PermissionPathSelectorMetadata) WithSubArrayOfValues(subPathArrayOfValues meta.Meta_FieldPathArrayOfValues) *Permission_FieldSubPathArrayOfValues {
	return &Permission_FieldSubPathArrayOfValues{Permission_FieldPath: s.WithSubPath(subPathArrayOfValues), subPathArrayOfValues: subPathArrayOfValues}
}

func (s PermissionPathSelectorMetadata) WithSubArrayItemValue(subPathArrayItemValue meta.Meta_FieldPathArrayItemValue) *Permission_FieldSubPathArrayItemValue {
	return &Permission_FieldSubPathArrayItemValue{Permission_FieldPath: s.WithSubPath(subPathArrayItemValue), subPathItemValue: subPathArrayItemValue}
}

func (PermissionPathSelectorMetadata) CreateTime() PermissionPathSelectorMetadataCreateTime {
	return PermissionPathSelectorMetadataCreateTime{}
}

func (PermissionPathSelectorMetadata) UpdateTime() PermissionPathSelectorMetadataUpdateTime {
	return PermissionPathSelectorMetadataUpdateTime{}
}

func (PermissionPathSelectorMetadata) DeleteTime() PermissionPathSelectorMetadataDeleteTime {
	return PermissionPathSelectorMetadataDeleteTime{}
}

func (PermissionPathSelectorMetadata) Uuid() PermissionPathSelectorMetadataUuid {
	return PermissionPathSelectorMetadataUuid{}
}

func (PermissionPathSelectorMetadata) Tags() PermissionPathSelectorMetadataTags {
	return PermissionPathSelectorMetadataTags{}
}

func (PermissionPathSelectorMetadata) Labels() PermissionPathSelectorMetadataLabels {
	return PermissionPathSelectorMetadataLabels{}
}

func (PermissionPathSelectorMetadata) Annotations() PermissionPathSelectorMetadataAnnotations {
	return PermissionPathSelectorMetadataAnnotations{}
}

func (PermissionPathSelectorMetadata) Generation() PermissionPathSelectorMetadataGeneration {
	return PermissionPathSelectorMetadataGeneration{}
}

func (PermissionPathSelectorMetadata) ResourceVersion() PermissionPathSelectorMetadataResourceVersion {
	return PermissionPathSelectorMetadataResourceVersion{}
}

func (PermissionPathSelectorMetadata) OwnerReferences() PermissionPathSelectorMetadataOwnerReferences {
	return PermissionPathSelectorMetadataOwnerReferences{}
}

func (PermissionPathSelectorMetadata) Shards() PermissionPathSelectorMetadataShards {
	return PermissionPathSelectorMetadataShards{}
}

func (PermissionPathSelectorMetadata) Syncing() PermissionPathSelectorMetadataSyncing {
	return PermissionPathSelectorMetadataSyncing{}
}

func (PermissionPathSelectorMetadata) Lifecycle() PermissionPathSelectorMetadataLifecycle {
	return PermissionPathSelectorMetadataLifecycle{}
}

func (PermissionPathSelectorMetadata) Services() PermissionPathSelectorMetadataServices {
	return PermissionPathSelectorMetadataServices{}
}

type PermissionPathSelectorMetadataCreateTime struct{}

func (PermissionPathSelectorMetadataCreateTime) FieldPath() *Permission_FieldSubPath {
	return &Permission_FieldSubPath{
		selector: Permission_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().CreateTime().FieldPath(),
	}
}

func (s PermissionPathSelectorMetadataCreateTime) WithValue(value *timestamppb.Timestamp) *Permission_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Permission_FieldSubPathValue)
}

func (s PermissionPathSelectorMetadataCreateTime) WithArrayOfValues(values []*timestamppb.Timestamp) *Permission_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Permission_FieldSubPathArrayOfValues)
}

type PermissionPathSelectorMetadataUpdateTime struct{}

func (PermissionPathSelectorMetadataUpdateTime) FieldPath() *Permission_FieldSubPath {
	return &Permission_FieldSubPath{
		selector: Permission_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().UpdateTime().FieldPath(),
	}
}

func (s PermissionPathSelectorMetadataUpdateTime) WithValue(value *timestamppb.Timestamp) *Permission_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Permission_FieldSubPathValue)
}

func (s PermissionPathSelectorMetadataUpdateTime) WithArrayOfValues(values []*timestamppb.Timestamp) *Permission_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Permission_FieldSubPathArrayOfValues)
}

type PermissionPathSelectorMetadataDeleteTime struct{}

func (PermissionPathSelectorMetadataDeleteTime) FieldPath() *Permission_FieldSubPath {
	return &Permission_FieldSubPath{
		selector: Permission_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().DeleteTime().FieldPath(),
	}
}

func (s PermissionPathSelectorMetadataDeleteTime) WithValue(value *timestamppb.Timestamp) *Permission_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Permission_FieldSubPathValue)
}

func (s PermissionPathSelectorMetadataDeleteTime) WithArrayOfValues(values []*timestamppb.Timestamp) *Permission_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Permission_FieldSubPathArrayOfValues)
}

type PermissionPathSelectorMetadataUuid struct{}

func (PermissionPathSelectorMetadataUuid) FieldPath() *Permission_FieldSubPath {
	return &Permission_FieldSubPath{
		selector: Permission_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Uuid().FieldPath(),
	}
}

func (s PermissionPathSelectorMetadataUuid) WithValue(value string) *Permission_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Permission_FieldSubPathValue)
}

func (s PermissionPathSelectorMetadataUuid) WithArrayOfValues(values []string) *Permission_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Permission_FieldSubPathArrayOfValues)
}

type PermissionPathSelectorMetadataTags struct{}

func (PermissionPathSelectorMetadataTags) FieldPath() *Permission_FieldSubPath {
	return &Permission_FieldSubPath{
		selector: Permission_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Tags().FieldPath(),
	}
}

func (s PermissionPathSelectorMetadataTags) WithValue(value []string) *Permission_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Permission_FieldSubPathValue)
}

func (s PermissionPathSelectorMetadataTags) WithArrayOfValues(values [][]string) *Permission_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Permission_FieldSubPathArrayOfValues)
}

func (s PermissionPathSelectorMetadataTags) WithItemValue(value string) *Permission_FieldSubPathArrayItemValue {
	return s.FieldPath().WithIArrayItemValue(value).(*Permission_FieldSubPathArrayItemValue)
}

type PermissionPathSelectorMetadataLabels struct{}

func (PermissionPathSelectorMetadataLabels) FieldPath() *Permission_FieldSubPath {
	return &Permission_FieldSubPath{
		selector: Permission_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Labels().FieldPath(),
	}
}

func (s PermissionPathSelectorMetadataLabels) WithValue(value map[string]string) *Permission_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Permission_FieldSubPathValue)
}

func (s PermissionPathSelectorMetadataLabels) WithArrayOfValues(values []map[string]string) *Permission_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Permission_FieldSubPathArrayOfValues)
}

func (PermissionPathSelectorMetadataLabels) WithKey(key string) PermissionMapPathSelectorMetadataLabels {
	return PermissionMapPathSelectorMetadataLabels{key: key}
}

type PermissionMapPathSelectorMetadataLabels struct {
	key string
}

func (s PermissionMapPathSelectorMetadataLabels) FieldPath() *Permission_FieldSubPath {
	return &Permission_FieldSubPath{
		selector: Permission_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Labels().WithKey(s.key).FieldPath(),
	}
}

func (s PermissionMapPathSelectorMetadataLabels) WithValue(value string) *Permission_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Permission_FieldSubPathValue)
}

func (s PermissionMapPathSelectorMetadataLabels) WithArrayOfValues(values []string) *Permission_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Permission_FieldSubPathArrayOfValues)
}

type PermissionPathSelectorMetadataAnnotations struct{}

func (PermissionPathSelectorMetadataAnnotations) FieldPath() *Permission_FieldSubPath {
	return &Permission_FieldSubPath{
		selector: Permission_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Annotations().FieldPath(),
	}
}

func (s PermissionPathSelectorMetadataAnnotations) WithValue(value map[string]string) *Permission_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Permission_FieldSubPathValue)
}

func (s PermissionPathSelectorMetadataAnnotations) WithArrayOfValues(values []map[string]string) *Permission_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Permission_FieldSubPathArrayOfValues)
}

func (PermissionPathSelectorMetadataAnnotations) WithKey(key string) PermissionMapPathSelectorMetadataAnnotations {
	return PermissionMapPathSelectorMetadataAnnotations{key: key}
}

type PermissionMapPathSelectorMetadataAnnotations struct {
	key string
}

func (s PermissionMapPathSelectorMetadataAnnotations) FieldPath() *Permission_FieldSubPath {
	return &Permission_FieldSubPath{
		selector: Permission_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Annotations().WithKey(s.key).FieldPath(),
	}
}

func (s PermissionMapPathSelectorMetadataAnnotations) WithValue(value string) *Permission_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Permission_FieldSubPathValue)
}

func (s PermissionMapPathSelectorMetadataAnnotations) WithArrayOfValues(values []string) *Permission_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Permission_FieldSubPathArrayOfValues)
}

type PermissionPathSelectorMetadataGeneration struct{}

func (PermissionPathSelectorMetadataGeneration) FieldPath() *Permission_FieldSubPath {
	return &Permission_FieldSubPath{
		selector: Permission_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Generation().FieldPath(),
	}
}

func (s PermissionPathSelectorMetadataGeneration) WithValue(value int64) *Permission_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Permission_FieldSubPathValue)
}

func (s PermissionPathSelectorMetadataGeneration) WithArrayOfValues(values []int64) *Permission_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Permission_FieldSubPathArrayOfValues)
}

type PermissionPathSelectorMetadataResourceVersion struct{}

func (PermissionPathSelectorMetadataResourceVersion) FieldPath() *Permission_FieldSubPath {
	return &Permission_FieldSubPath{
		selector: Permission_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().ResourceVersion().FieldPath(),
	}
}

func (s PermissionPathSelectorMetadataResourceVersion) WithValue(value string) *Permission_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Permission_FieldSubPathValue)
}

func (s PermissionPathSelectorMetadataResourceVersion) WithArrayOfValues(values []string) *Permission_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Permission_FieldSubPathArrayOfValues)
}

type PermissionPathSelectorMetadataOwnerReferences struct{}

func (PermissionPathSelectorMetadataOwnerReferences) FieldPath() *Permission_FieldSubPath {
	return &Permission_FieldSubPath{
		selector: Permission_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().OwnerReferences().FieldPath(),
	}
}

func (s PermissionPathSelectorMetadataOwnerReferences) WithValue(value []*meta.OwnerReference) *Permission_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Permission_FieldSubPathValue)
}

func (s PermissionPathSelectorMetadataOwnerReferences) WithArrayOfValues(values [][]*meta.OwnerReference) *Permission_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Permission_FieldSubPathArrayOfValues)
}

func (s PermissionPathSelectorMetadataOwnerReferences) WithItemValue(value *meta.OwnerReference) *Permission_FieldSubPathArrayItemValue {
	return s.FieldPath().WithIArrayItemValue(value).(*Permission_FieldSubPathArrayItemValue)
}

func (PermissionPathSelectorMetadataOwnerReferences) Kind() PermissionPathSelectorMetadataOwnerReferencesKind {
	return PermissionPathSelectorMetadataOwnerReferencesKind{}
}

func (PermissionPathSelectorMetadataOwnerReferences) Version() PermissionPathSelectorMetadataOwnerReferencesVersion {
	return PermissionPathSelectorMetadataOwnerReferencesVersion{}
}

func (PermissionPathSelectorMetadataOwnerReferences) Name() PermissionPathSelectorMetadataOwnerReferencesName {
	return PermissionPathSelectorMetadataOwnerReferencesName{}
}

func (PermissionPathSelectorMetadataOwnerReferences) Region() PermissionPathSelectorMetadataOwnerReferencesRegion {
	return PermissionPathSelectorMetadataOwnerReferencesRegion{}
}

func (PermissionPathSelectorMetadataOwnerReferences) Controller() PermissionPathSelectorMetadataOwnerReferencesController {
	return PermissionPathSelectorMetadataOwnerReferencesController{}
}

func (PermissionPathSelectorMetadataOwnerReferences) RequiresOwnerReference() PermissionPathSelectorMetadataOwnerReferencesRequiresOwnerReference {
	return PermissionPathSelectorMetadataOwnerReferencesRequiresOwnerReference{}
}

type PermissionPathSelectorMetadataOwnerReferencesKind struct{}

func (PermissionPathSelectorMetadataOwnerReferencesKind) FieldPath() *Permission_FieldSubPath {
	return &Permission_FieldSubPath{
		selector: Permission_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().OwnerReferences().Kind().FieldPath(),
	}
}

func (s PermissionPathSelectorMetadataOwnerReferencesKind) WithValue(value string) *Permission_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Permission_FieldSubPathValue)
}

func (s PermissionPathSelectorMetadataOwnerReferencesKind) WithArrayOfValues(values []string) *Permission_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Permission_FieldSubPathArrayOfValues)
}

type PermissionPathSelectorMetadataOwnerReferencesVersion struct{}

func (PermissionPathSelectorMetadataOwnerReferencesVersion) FieldPath() *Permission_FieldSubPath {
	return &Permission_FieldSubPath{
		selector: Permission_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().OwnerReferences().Version().FieldPath(),
	}
}

func (s PermissionPathSelectorMetadataOwnerReferencesVersion) WithValue(value string) *Permission_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Permission_FieldSubPathValue)
}

func (s PermissionPathSelectorMetadataOwnerReferencesVersion) WithArrayOfValues(values []string) *Permission_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Permission_FieldSubPathArrayOfValues)
}

type PermissionPathSelectorMetadataOwnerReferencesName struct{}

func (PermissionPathSelectorMetadataOwnerReferencesName) FieldPath() *Permission_FieldSubPath {
	return &Permission_FieldSubPath{
		selector: Permission_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().OwnerReferences().Name().FieldPath(),
	}
}

func (s PermissionPathSelectorMetadataOwnerReferencesName) WithValue(value string) *Permission_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Permission_FieldSubPathValue)
}

func (s PermissionPathSelectorMetadataOwnerReferencesName) WithArrayOfValues(values []string) *Permission_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Permission_FieldSubPathArrayOfValues)
}

type PermissionPathSelectorMetadataOwnerReferencesRegion struct{}

func (PermissionPathSelectorMetadataOwnerReferencesRegion) FieldPath() *Permission_FieldSubPath {
	return &Permission_FieldSubPath{
		selector: Permission_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().OwnerReferences().Region().FieldPath(),
	}
}

func (s PermissionPathSelectorMetadataOwnerReferencesRegion) WithValue(value string) *Permission_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Permission_FieldSubPathValue)
}

func (s PermissionPathSelectorMetadataOwnerReferencesRegion) WithArrayOfValues(values []string) *Permission_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Permission_FieldSubPathArrayOfValues)
}

type PermissionPathSelectorMetadataOwnerReferencesController struct{}

func (PermissionPathSelectorMetadataOwnerReferencesController) FieldPath() *Permission_FieldSubPath {
	return &Permission_FieldSubPath{
		selector: Permission_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().OwnerReferences().Controller().FieldPath(),
	}
}

func (s PermissionPathSelectorMetadataOwnerReferencesController) WithValue(value bool) *Permission_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Permission_FieldSubPathValue)
}

func (s PermissionPathSelectorMetadataOwnerReferencesController) WithArrayOfValues(values []bool) *Permission_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Permission_FieldSubPathArrayOfValues)
}

type PermissionPathSelectorMetadataOwnerReferencesRequiresOwnerReference struct{}

func (PermissionPathSelectorMetadataOwnerReferencesRequiresOwnerReference) FieldPath() *Permission_FieldSubPath {
	return &Permission_FieldSubPath{
		selector: Permission_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().OwnerReferences().RequiresOwnerReference().FieldPath(),
	}
}

func (s PermissionPathSelectorMetadataOwnerReferencesRequiresOwnerReference) WithValue(value bool) *Permission_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Permission_FieldSubPathValue)
}

func (s PermissionPathSelectorMetadataOwnerReferencesRequiresOwnerReference) WithArrayOfValues(values []bool) *Permission_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Permission_FieldSubPathArrayOfValues)
}

type PermissionPathSelectorMetadataShards struct{}

func (PermissionPathSelectorMetadataShards) FieldPath() *Permission_FieldSubPath {
	return &Permission_FieldSubPath{
		selector: Permission_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Shards().FieldPath(),
	}
}

func (s PermissionPathSelectorMetadataShards) WithValue(value map[string]int64) *Permission_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Permission_FieldSubPathValue)
}

func (s PermissionPathSelectorMetadataShards) WithArrayOfValues(values []map[string]int64) *Permission_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Permission_FieldSubPathArrayOfValues)
}

func (PermissionPathSelectorMetadataShards) WithKey(key string) PermissionMapPathSelectorMetadataShards {
	return PermissionMapPathSelectorMetadataShards{key: key}
}

type PermissionMapPathSelectorMetadataShards struct {
	key string
}

func (s PermissionMapPathSelectorMetadataShards) FieldPath() *Permission_FieldSubPath {
	return &Permission_FieldSubPath{
		selector: Permission_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Shards().WithKey(s.key).FieldPath(),
	}
}

func (s PermissionMapPathSelectorMetadataShards) WithValue(value int64) *Permission_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Permission_FieldSubPathValue)
}

func (s PermissionMapPathSelectorMetadataShards) WithArrayOfValues(values []int64) *Permission_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Permission_FieldSubPathArrayOfValues)
}

type PermissionPathSelectorMetadataSyncing struct{}

func (PermissionPathSelectorMetadataSyncing) FieldPath() *Permission_FieldSubPath {
	return &Permission_FieldSubPath{
		selector: Permission_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Syncing().FieldPath(),
	}
}

func (s PermissionPathSelectorMetadataSyncing) WithValue(value *meta.SyncingMeta) *Permission_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Permission_FieldSubPathValue)
}

func (s PermissionPathSelectorMetadataSyncing) WithArrayOfValues(values []*meta.SyncingMeta) *Permission_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Permission_FieldSubPathArrayOfValues)
}

func (PermissionPathSelectorMetadataSyncing) OwningRegion() PermissionPathSelectorMetadataSyncingOwningRegion {
	return PermissionPathSelectorMetadataSyncingOwningRegion{}
}

func (PermissionPathSelectorMetadataSyncing) Regions() PermissionPathSelectorMetadataSyncingRegions {
	return PermissionPathSelectorMetadataSyncingRegions{}
}

type PermissionPathSelectorMetadataSyncingOwningRegion struct{}

func (PermissionPathSelectorMetadataSyncingOwningRegion) FieldPath() *Permission_FieldSubPath {
	return &Permission_FieldSubPath{
		selector: Permission_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Syncing().OwningRegion().FieldPath(),
	}
}

func (s PermissionPathSelectorMetadataSyncingOwningRegion) WithValue(value string) *Permission_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Permission_FieldSubPathValue)
}

func (s PermissionPathSelectorMetadataSyncingOwningRegion) WithArrayOfValues(values []string) *Permission_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Permission_FieldSubPathArrayOfValues)
}

type PermissionPathSelectorMetadataSyncingRegions struct{}

func (PermissionPathSelectorMetadataSyncingRegions) FieldPath() *Permission_FieldSubPath {
	return &Permission_FieldSubPath{
		selector: Permission_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Syncing().Regions().FieldPath(),
	}
}

func (s PermissionPathSelectorMetadataSyncingRegions) WithValue(value []string) *Permission_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Permission_FieldSubPathValue)
}

func (s PermissionPathSelectorMetadataSyncingRegions) WithArrayOfValues(values [][]string) *Permission_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Permission_FieldSubPathArrayOfValues)
}

func (s PermissionPathSelectorMetadataSyncingRegions) WithItemValue(value string) *Permission_FieldSubPathArrayItemValue {
	return s.FieldPath().WithIArrayItemValue(value).(*Permission_FieldSubPathArrayItemValue)
}

type PermissionPathSelectorMetadataLifecycle struct{}

func (PermissionPathSelectorMetadataLifecycle) FieldPath() *Permission_FieldSubPath {
	return &Permission_FieldSubPath{
		selector: Permission_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Lifecycle().FieldPath(),
	}
}

func (s PermissionPathSelectorMetadataLifecycle) WithValue(value *meta.Lifecycle) *Permission_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Permission_FieldSubPathValue)
}

func (s PermissionPathSelectorMetadataLifecycle) WithArrayOfValues(values []*meta.Lifecycle) *Permission_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Permission_FieldSubPathArrayOfValues)
}

func (PermissionPathSelectorMetadataLifecycle) State() PermissionPathSelectorMetadataLifecycleState {
	return PermissionPathSelectorMetadataLifecycleState{}
}

func (PermissionPathSelectorMetadataLifecycle) BlockDeletion() PermissionPathSelectorMetadataLifecycleBlockDeletion {
	return PermissionPathSelectorMetadataLifecycleBlockDeletion{}
}

type PermissionPathSelectorMetadataLifecycleState struct{}

func (PermissionPathSelectorMetadataLifecycleState) FieldPath() *Permission_FieldSubPath {
	return &Permission_FieldSubPath{
		selector: Permission_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Lifecycle().State().FieldPath(),
	}
}

func (s PermissionPathSelectorMetadataLifecycleState) WithValue(value meta.Lifecycle_State) *Permission_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Permission_FieldSubPathValue)
}

func (s PermissionPathSelectorMetadataLifecycleState) WithArrayOfValues(values []meta.Lifecycle_State) *Permission_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Permission_FieldSubPathArrayOfValues)
}

type PermissionPathSelectorMetadataLifecycleBlockDeletion struct{}

func (PermissionPathSelectorMetadataLifecycleBlockDeletion) FieldPath() *Permission_FieldSubPath {
	return &Permission_FieldSubPath{
		selector: Permission_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Lifecycle().BlockDeletion().FieldPath(),
	}
}

func (s PermissionPathSelectorMetadataLifecycleBlockDeletion) WithValue(value bool) *Permission_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Permission_FieldSubPathValue)
}

func (s PermissionPathSelectorMetadataLifecycleBlockDeletion) WithArrayOfValues(values []bool) *Permission_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Permission_FieldSubPathArrayOfValues)
}

type PermissionPathSelectorMetadataServices struct{}

func (PermissionPathSelectorMetadataServices) FieldPath() *Permission_FieldSubPath {
	return &Permission_FieldSubPath{
		selector: Permission_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Services().FieldPath(),
	}
}

func (s PermissionPathSelectorMetadataServices) WithValue(value *meta.ServicesInfo) *Permission_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Permission_FieldSubPathValue)
}

func (s PermissionPathSelectorMetadataServices) WithArrayOfValues(values []*meta.ServicesInfo) *Permission_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Permission_FieldSubPathArrayOfValues)
}

func (PermissionPathSelectorMetadataServices) OwningService() PermissionPathSelectorMetadataServicesOwningService {
	return PermissionPathSelectorMetadataServicesOwningService{}
}

func (PermissionPathSelectorMetadataServices) AllowedServices() PermissionPathSelectorMetadataServicesAllowedServices {
	return PermissionPathSelectorMetadataServicesAllowedServices{}
}

type PermissionPathSelectorMetadataServicesOwningService struct{}

func (PermissionPathSelectorMetadataServicesOwningService) FieldPath() *Permission_FieldSubPath {
	return &Permission_FieldSubPath{
		selector: Permission_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Services().OwningService().FieldPath(),
	}
}

func (s PermissionPathSelectorMetadataServicesOwningService) WithValue(value string) *Permission_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Permission_FieldSubPathValue)
}

func (s PermissionPathSelectorMetadataServicesOwningService) WithArrayOfValues(values []string) *Permission_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Permission_FieldSubPathArrayOfValues)
}

type PermissionPathSelectorMetadataServicesAllowedServices struct{}

func (PermissionPathSelectorMetadataServicesAllowedServices) FieldPath() *Permission_FieldSubPath {
	return &Permission_FieldSubPath{
		selector: Permission_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Services().AllowedServices().FieldPath(),
	}
}

func (s PermissionPathSelectorMetadataServicesAllowedServices) WithValue(value []string) *Permission_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Permission_FieldSubPathValue)
}

func (s PermissionPathSelectorMetadataServicesAllowedServices) WithArrayOfValues(values [][]string) *Permission_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Permission_FieldSubPathArrayOfValues)
}

func (s PermissionPathSelectorMetadataServicesAllowedServices) WithItemValue(value string) *Permission_FieldSubPathArrayItemValue {
	return s.FieldPath().WithIArrayItemValue(value).(*Permission_FieldSubPathArrayItemValue)
}

type PermissionPathSelectorTitle struct{}

func (PermissionPathSelectorTitle) FieldPath() *Permission_FieldTerminalPath {
	return &Permission_FieldTerminalPath{selector: Permission_FieldPathSelectorTitle}
}

func (s PermissionPathSelectorTitle) WithValue(value string) *Permission_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*Permission_FieldTerminalPathValue)
}

func (s PermissionPathSelectorTitle) WithArrayOfValues(values []string) *Permission_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Permission_FieldTerminalPathArrayOfValues)
}

type PermissionPathSelectorDescription struct{}

func (PermissionPathSelectorDescription) FieldPath() *Permission_FieldTerminalPath {
	return &Permission_FieldTerminalPath{selector: Permission_FieldPathSelectorDescription}
}

func (s PermissionPathSelectorDescription) WithValue(value string) *Permission_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*Permission_FieldTerminalPathValue)
}

func (s PermissionPathSelectorDescription) WithArrayOfValues(values []string) *Permission_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Permission_FieldTerminalPathArrayOfValues)
}