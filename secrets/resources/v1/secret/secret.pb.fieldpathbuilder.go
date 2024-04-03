// Code generated by protoc-gen-goten-object
// File: edgelq/secrets/proto/v1/secret.proto
// DO NOT EDIT!!!

package secret

// proto imports
import (
	project "github.com/cloudwan/edgelq-sdk/secrets/resources/v1/project"
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

type SecretFieldPathBuilder struct{}

func NewSecretFieldPathBuilder() SecretFieldPathBuilder {
	return SecretFieldPathBuilder{}
}
func (SecretFieldPathBuilder) Name() SecretPathSelectorName {
	return SecretPathSelectorName{}
}
func (SecretFieldPathBuilder) Metadata() SecretPathSelectorMetadata {
	return SecretPathSelectorMetadata{}
}
func (SecretFieldPathBuilder) EncData() SecretPathSelectorEncData {
	return SecretPathSelectorEncData{}
}
func (SecretFieldPathBuilder) Data() SecretPathSelectorData {
	return SecretPathSelectorData{}
}

type SecretPathSelectorName struct{}

func (SecretPathSelectorName) FieldPath() *Secret_FieldTerminalPath {
	return &Secret_FieldTerminalPath{selector: Secret_FieldPathSelectorName}
}

func (s SecretPathSelectorName) WithValue(value *Name) *Secret_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*Secret_FieldTerminalPathValue)
}

func (s SecretPathSelectorName) WithArrayOfValues(values []*Name) *Secret_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Secret_FieldTerminalPathArrayOfValues)
}

type SecretPathSelectorMetadata struct{}

func (SecretPathSelectorMetadata) FieldPath() *Secret_FieldTerminalPath {
	return &Secret_FieldTerminalPath{selector: Secret_FieldPathSelectorMetadata}
}

func (s SecretPathSelectorMetadata) WithValue(value *meta.Meta) *Secret_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*Secret_FieldTerminalPathValue)
}

func (s SecretPathSelectorMetadata) WithArrayOfValues(values []*meta.Meta) *Secret_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Secret_FieldTerminalPathArrayOfValues)
}

func (SecretPathSelectorMetadata) WithSubPath(subPath meta.Meta_FieldPath) *Secret_FieldSubPath {
	return &Secret_FieldSubPath{selector: Secret_FieldPathSelectorMetadata, subPath: subPath}
}

func (s SecretPathSelectorMetadata) WithSubValue(subPathValue meta.Meta_FieldPathValue) *Secret_FieldSubPathValue {
	return &Secret_FieldSubPathValue{Secret_FieldPath: s.WithSubPath(subPathValue), subPathValue: subPathValue}
}

func (s SecretPathSelectorMetadata) WithSubArrayOfValues(subPathArrayOfValues meta.Meta_FieldPathArrayOfValues) *Secret_FieldSubPathArrayOfValues {
	return &Secret_FieldSubPathArrayOfValues{Secret_FieldPath: s.WithSubPath(subPathArrayOfValues), subPathArrayOfValues: subPathArrayOfValues}
}

func (s SecretPathSelectorMetadata) WithSubArrayItemValue(subPathArrayItemValue meta.Meta_FieldPathArrayItemValue) *Secret_FieldSubPathArrayItemValue {
	return &Secret_FieldSubPathArrayItemValue{Secret_FieldPath: s.WithSubPath(subPathArrayItemValue), subPathItemValue: subPathArrayItemValue}
}

func (SecretPathSelectorMetadata) CreateTime() SecretPathSelectorMetadataCreateTime {
	return SecretPathSelectorMetadataCreateTime{}
}

func (SecretPathSelectorMetadata) UpdateTime() SecretPathSelectorMetadataUpdateTime {
	return SecretPathSelectorMetadataUpdateTime{}
}

func (SecretPathSelectorMetadata) DeleteTime() SecretPathSelectorMetadataDeleteTime {
	return SecretPathSelectorMetadataDeleteTime{}
}

func (SecretPathSelectorMetadata) Uuid() SecretPathSelectorMetadataUuid {
	return SecretPathSelectorMetadataUuid{}
}

func (SecretPathSelectorMetadata) Tags() SecretPathSelectorMetadataTags {
	return SecretPathSelectorMetadataTags{}
}

func (SecretPathSelectorMetadata) Labels() SecretPathSelectorMetadataLabels {
	return SecretPathSelectorMetadataLabels{}
}

func (SecretPathSelectorMetadata) Annotations() SecretPathSelectorMetadataAnnotations {
	return SecretPathSelectorMetadataAnnotations{}
}

func (SecretPathSelectorMetadata) Generation() SecretPathSelectorMetadataGeneration {
	return SecretPathSelectorMetadataGeneration{}
}

func (SecretPathSelectorMetadata) ResourceVersion() SecretPathSelectorMetadataResourceVersion {
	return SecretPathSelectorMetadataResourceVersion{}
}

func (SecretPathSelectorMetadata) OwnerReferences() SecretPathSelectorMetadataOwnerReferences {
	return SecretPathSelectorMetadataOwnerReferences{}
}

func (SecretPathSelectorMetadata) Shards() SecretPathSelectorMetadataShards {
	return SecretPathSelectorMetadataShards{}
}

func (SecretPathSelectorMetadata) Syncing() SecretPathSelectorMetadataSyncing {
	return SecretPathSelectorMetadataSyncing{}
}

func (SecretPathSelectorMetadata) Lifecycle() SecretPathSelectorMetadataLifecycle {
	return SecretPathSelectorMetadataLifecycle{}
}

func (SecretPathSelectorMetadata) Services() SecretPathSelectorMetadataServices {
	return SecretPathSelectorMetadataServices{}
}

type SecretPathSelectorMetadataCreateTime struct{}

func (SecretPathSelectorMetadataCreateTime) FieldPath() *Secret_FieldSubPath {
	return &Secret_FieldSubPath{
		selector: Secret_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().CreateTime().FieldPath(),
	}
}

func (s SecretPathSelectorMetadataCreateTime) WithValue(value *timestamppb.Timestamp) *Secret_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Secret_FieldSubPathValue)
}

func (s SecretPathSelectorMetadataCreateTime) WithArrayOfValues(values []*timestamppb.Timestamp) *Secret_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Secret_FieldSubPathArrayOfValues)
}

type SecretPathSelectorMetadataUpdateTime struct{}

func (SecretPathSelectorMetadataUpdateTime) FieldPath() *Secret_FieldSubPath {
	return &Secret_FieldSubPath{
		selector: Secret_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().UpdateTime().FieldPath(),
	}
}

func (s SecretPathSelectorMetadataUpdateTime) WithValue(value *timestamppb.Timestamp) *Secret_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Secret_FieldSubPathValue)
}

func (s SecretPathSelectorMetadataUpdateTime) WithArrayOfValues(values []*timestamppb.Timestamp) *Secret_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Secret_FieldSubPathArrayOfValues)
}

type SecretPathSelectorMetadataDeleteTime struct{}

func (SecretPathSelectorMetadataDeleteTime) FieldPath() *Secret_FieldSubPath {
	return &Secret_FieldSubPath{
		selector: Secret_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().DeleteTime().FieldPath(),
	}
}

func (s SecretPathSelectorMetadataDeleteTime) WithValue(value *timestamppb.Timestamp) *Secret_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Secret_FieldSubPathValue)
}

func (s SecretPathSelectorMetadataDeleteTime) WithArrayOfValues(values []*timestamppb.Timestamp) *Secret_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Secret_FieldSubPathArrayOfValues)
}

type SecretPathSelectorMetadataUuid struct{}

func (SecretPathSelectorMetadataUuid) FieldPath() *Secret_FieldSubPath {
	return &Secret_FieldSubPath{
		selector: Secret_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Uuid().FieldPath(),
	}
}

func (s SecretPathSelectorMetadataUuid) WithValue(value string) *Secret_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Secret_FieldSubPathValue)
}

func (s SecretPathSelectorMetadataUuid) WithArrayOfValues(values []string) *Secret_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Secret_FieldSubPathArrayOfValues)
}

type SecretPathSelectorMetadataTags struct{}

func (SecretPathSelectorMetadataTags) FieldPath() *Secret_FieldSubPath {
	return &Secret_FieldSubPath{
		selector: Secret_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Tags().FieldPath(),
	}
}

func (s SecretPathSelectorMetadataTags) WithValue(value []string) *Secret_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Secret_FieldSubPathValue)
}

func (s SecretPathSelectorMetadataTags) WithArrayOfValues(values [][]string) *Secret_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Secret_FieldSubPathArrayOfValues)
}

func (s SecretPathSelectorMetadataTags) WithItemValue(value string) *Secret_FieldSubPathArrayItemValue {
	return s.FieldPath().WithIArrayItemValue(value).(*Secret_FieldSubPathArrayItemValue)
}

type SecretPathSelectorMetadataLabels struct{}

func (SecretPathSelectorMetadataLabels) FieldPath() *Secret_FieldSubPath {
	return &Secret_FieldSubPath{
		selector: Secret_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Labels().FieldPath(),
	}
}

func (s SecretPathSelectorMetadataLabels) WithValue(value map[string]string) *Secret_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Secret_FieldSubPathValue)
}

func (s SecretPathSelectorMetadataLabels) WithArrayOfValues(values []map[string]string) *Secret_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Secret_FieldSubPathArrayOfValues)
}

func (SecretPathSelectorMetadataLabels) WithKey(key string) SecretMapPathSelectorMetadataLabels {
	return SecretMapPathSelectorMetadataLabels{key: key}
}

type SecretMapPathSelectorMetadataLabels struct {
	key string
}

func (s SecretMapPathSelectorMetadataLabels) FieldPath() *Secret_FieldSubPath {
	return &Secret_FieldSubPath{
		selector: Secret_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Labels().WithKey(s.key).FieldPath(),
	}
}

func (s SecretMapPathSelectorMetadataLabels) WithValue(value string) *Secret_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Secret_FieldSubPathValue)
}

func (s SecretMapPathSelectorMetadataLabels) WithArrayOfValues(values []string) *Secret_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Secret_FieldSubPathArrayOfValues)
}

type SecretPathSelectorMetadataAnnotations struct{}

func (SecretPathSelectorMetadataAnnotations) FieldPath() *Secret_FieldSubPath {
	return &Secret_FieldSubPath{
		selector: Secret_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Annotations().FieldPath(),
	}
}

func (s SecretPathSelectorMetadataAnnotations) WithValue(value map[string]string) *Secret_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Secret_FieldSubPathValue)
}

func (s SecretPathSelectorMetadataAnnotations) WithArrayOfValues(values []map[string]string) *Secret_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Secret_FieldSubPathArrayOfValues)
}

func (SecretPathSelectorMetadataAnnotations) WithKey(key string) SecretMapPathSelectorMetadataAnnotations {
	return SecretMapPathSelectorMetadataAnnotations{key: key}
}

type SecretMapPathSelectorMetadataAnnotations struct {
	key string
}

func (s SecretMapPathSelectorMetadataAnnotations) FieldPath() *Secret_FieldSubPath {
	return &Secret_FieldSubPath{
		selector: Secret_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Annotations().WithKey(s.key).FieldPath(),
	}
}

func (s SecretMapPathSelectorMetadataAnnotations) WithValue(value string) *Secret_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Secret_FieldSubPathValue)
}

func (s SecretMapPathSelectorMetadataAnnotations) WithArrayOfValues(values []string) *Secret_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Secret_FieldSubPathArrayOfValues)
}

type SecretPathSelectorMetadataGeneration struct{}

func (SecretPathSelectorMetadataGeneration) FieldPath() *Secret_FieldSubPath {
	return &Secret_FieldSubPath{
		selector: Secret_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Generation().FieldPath(),
	}
}

func (s SecretPathSelectorMetadataGeneration) WithValue(value int64) *Secret_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Secret_FieldSubPathValue)
}

func (s SecretPathSelectorMetadataGeneration) WithArrayOfValues(values []int64) *Secret_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Secret_FieldSubPathArrayOfValues)
}

type SecretPathSelectorMetadataResourceVersion struct{}

func (SecretPathSelectorMetadataResourceVersion) FieldPath() *Secret_FieldSubPath {
	return &Secret_FieldSubPath{
		selector: Secret_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().ResourceVersion().FieldPath(),
	}
}

func (s SecretPathSelectorMetadataResourceVersion) WithValue(value string) *Secret_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Secret_FieldSubPathValue)
}

func (s SecretPathSelectorMetadataResourceVersion) WithArrayOfValues(values []string) *Secret_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Secret_FieldSubPathArrayOfValues)
}

type SecretPathSelectorMetadataOwnerReferences struct{}

func (SecretPathSelectorMetadataOwnerReferences) FieldPath() *Secret_FieldSubPath {
	return &Secret_FieldSubPath{
		selector: Secret_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().OwnerReferences().FieldPath(),
	}
}

func (s SecretPathSelectorMetadataOwnerReferences) WithValue(value []*meta.OwnerReference) *Secret_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Secret_FieldSubPathValue)
}

func (s SecretPathSelectorMetadataOwnerReferences) WithArrayOfValues(values [][]*meta.OwnerReference) *Secret_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Secret_FieldSubPathArrayOfValues)
}

func (s SecretPathSelectorMetadataOwnerReferences) WithItemValue(value *meta.OwnerReference) *Secret_FieldSubPathArrayItemValue {
	return s.FieldPath().WithIArrayItemValue(value).(*Secret_FieldSubPathArrayItemValue)
}

func (SecretPathSelectorMetadataOwnerReferences) Kind() SecretPathSelectorMetadataOwnerReferencesKind {
	return SecretPathSelectorMetadataOwnerReferencesKind{}
}

func (SecretPathSelectorMetadataOwnerReferences) Version() SecretPathSelectorMetadataOwnerReferencesVersion {
	return SecretPathSelectorMetadataOwnerReferencesVersion{}
}

func (SecretPathSelectorMetadataOwnerReferences) Name() SecretPathSelectorMetadataOwnerReferencesName {
	return SecretPathSelectorMetadataOwnerReferencesName{}
}

func (SecretPathSelectorMetadataOwnerReferences) Region() SecretPathSelectorMetadataOwnerReferencesRegion {
	return SecretPathSelectorMetadataOwnerReferencesRegion{}
}

func (SecretPathSelectorMetadataOwnerReferences) Controller() SecretPathSelectorMetadataOwnerReferencesController {
	return SecretPathSelectorMetadataOwnerReferencesController{}
}

func (SecretPathSelectorMetadataOwnerReferences) RequiresOwnerReference() SecretPathSelectorMetadataOwnerReferencesRequiresOwnerReference {
	return SecretPathSelectorMetadataOwnerReferencesRequiresOwnerReference{}
}

type SecretPathSelectorMetadataOwnerReferencesKind struct{}

func (SecretPathSelectorMetadataOwnerReferencesKind) FieldPath() *Secret_FieldSubPath {
	return &Secret_FieldSubPath{
		selector: Secret_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().OwnerReferences().Kind().FieldPath(),
	}
}

func (s SecretPathSelectorMetadataOwnerReferencesKind) WithValue(value string) *Secret_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Secret_FieldSubPathValue)
}

func (s SecretPathSelectorMetadataOwnerReferencesKind) WithArrayOfValues(values []string) *Secret_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Secret_FieldSubPathArrayOfValues)
}

type SecretPathSelectorMetadataOwnerReferencesVersion struct{}

func (SecretPathSelectorMetadataOwnerReferencesVersion) FieldPath() *Secret_FieldSubPath {
	return &Secret_FieldSubPath{
		selector: Secret_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().OwnerReferences().Version().FieldPath(),
	}
}

func (s SecretPathSelectorMetadataOwnerReferencesVersion) WithValue(value string) *Secret_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Secret_FieldSubPathValue)
}

func (s SecretPathSelectorMetadataOwnerReferencesVersion) WithArrayOfValues(values []string) *Secret_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Secret_FieldSubPathArrayOfValues)
}

type SecretPathSelectorMetadataOwnerReferencesName struct{}

func (SecretPathSelectorMetadataOwnerReferencesName) FieldPath() *Secret_FieldSubPath {
	return &Secret_FieldSubPath{
		selector: Secret_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().OwnerReferences().Name().FieldPath(),
	}
}

func (s SecretPathSelectorMetadataOwnerReferencesName) WithValue(value string) *Secret_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Secret_FieldSubPathValue)
}

func (s SecretPathSelectorMetadataOwnerReferencesName) WithArrayOfValues(values []string) *Secret_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Secret_FieldSubPathArrayOfValues)
}

type SecretPathSelectorMetadataOwnerReferencesRegion struct{}

func (SecretPathSelectorMetadataOwnerReferencesRegion) FieldPath() *Secret_FieldSubPath {
	return &Secret_FieldSubPath{
		selector: Secret_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().OwnerReferences().Region().FieldPath(),
	}
}

func (s SecretPathSelectorMetadataOwnerReferencesRegion) WithValue(value string) *Secret_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Secret_FieldSubPathValue)
}

func (s SecretPathSelectorMetadataOwnerReferencesRegion) WithArrayOfValues(values []string) *Secret_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Secret_FieldSubPathArrayOfValues)
}

type SecretPathSelectorMetadataOwnerReferencesController struct{}

func (SecretPathSelectorMetadataOwnerReferencesController) FieldPath() *Secret_FieldSubPath {
	return &Secret_FieldSubPath{
		selector: Secret_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().OwnerReferences().Controller().FieldPath(),
	}
}

func (s SecretPathSelectorMetadataOwnerReferencesController) WithValue(value bool) *Secret_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Secret_FieldSubPathValue)
}

func (s SecretPathSelectorMetadataOwnerReferencesController) WithArrayOfValues(values []bool) *Secret_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Secret_FieldSubPathArrayOfValues)
}

type SecretPathSelectorMetadataOwnerReferencesRequiresOwnerReference struct{}

func (SecretPathSelectorMetadataOwnerReferencesRequiresOwnerReference) FieldPath() *Secret_FieldSubPath {
	return &Secret_FieldSubPath{
		selector: Secret_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().OwnerReferences().RequiresOwnerReference().FieldPath(),
	}
}

func (s SecretPathSelectorMetadataOwnerReferencesRequiresOwnerReference) WithValue(value bool) *Secret_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Secret_FieldSubPathValue)
}

func (s SecretPathSelectorMetadataOwnerReferencesRequiresOwnerReference) WithArrayOfValues(values []bool) *Secret_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Secret_FieldSubPathArrayOfValues)
}

type SecretPathSelectorMetadataShards struct{}

func (SecretPathSelectorMetadataShards) FieldPath() *Secret_FieldSubPath {
	return &Secret_FieldSubPath{
		selector: Secret_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Shards().FieldPath(),
	}
}

func (s SecretPathSelectorMetadataShards) WithValue(value map[string]int64) *Secret_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Secret_FieldSubPathValue)
}

func (s SecretPathSelectorMetadataShards) WithArrayOfValues(values []map[string]int64) *Secret_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Secret_FieldSubPathArrayOfValues)
}

func (SecretPathSelectorMetadataShards) WithKey(key string) SecretMapPathSelectorMetadataShards {
	return SecretMapPathSelectorMetadataShards{key: key}
}

type SecretMapPathSelectorMetadataShards struct {
	key string
}

func (s SecretMapPathSelectorMetadataShards) FieldPath() *Secret_FieldSubPath {
	return &Secret_FieldSubPath{
		selector: Secret_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Shards().WithKey(s.key).FieldPath(),
	}
}

func (s SecretMapPathSelectorMetadataShards) WithValue(value int64) *Secret_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Secret_FieldSubPathValue)
}

func (s SecretMapPathSelectorMetadataShards) WithArrayOfValues(values []int64) *Secret_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Secret_FieldSubPathArrayOfValues)
}

type SecretPathSelectorMetadataSyncing struct{}

func (SecretPathSelectorMetadataSyncing) FieldPath() *Secret_FieldSubPath {
	return &Secret_FieldSubPath{
		selector: Secret_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Syncing().FieldPath(),
	}
}

func (s SecretPathSelectorMetadataSyncing) WithValue(value *meta.SyncingMeta) *Secret_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Secret_FieldSubPathValue)
}

func (s SecretPathSelectorMetadataSyncing) WithArrayOfValues(values []*meta.SyncingMeta) *Secret_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Secret_FieldSubPathArrayOfValues)
}

func (SecretPathSelectorMetadataSyncing) OwningRegion() SecretPathSelectorMetadataSyncingOwningRegion {
	return SecretPathSelectorMetadataSyncingOwningRegion{}
}

func (SecretPathSelectorMetadataSyncing) Regions() SecretPathSelectorMetadataSyncingRegions {
	return SecretPathSelectorMetadataSyncingRegions{}
}

type SecretPathSelectorMetadataSyncingOwningRegion struct{}

func (SecretPathSelectorMetadataSyncingOwningRegion) FieldPath() *Secret_FieldSubPath {
	return &Secret_FieldSubPath{
		selector: Secret_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Syncing().OwningRegion().FieldPath(),
	}
}

func (s SecretPathSelectorMetadataSyncingOwningRegion) WithValue(value string) *Secret_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Secret_FieldSubPathValue)
}

func (s SecretPathSelectorMetadataSyncingOwningRegion) WithArrayOfValues(values []string) *Secret_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Secret_FieldSubPathArrayOfValues)
}

type SecretPathSelectorMetadataSyncingRegions struct{}

func (SecretPathSelectorMetadataSyncingRegions) FieldPath() *Secret_FieldSubPath {
	return &Secret_FieldSubPath{
		selector: Secret_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Syncing().Regions().FieldPath(),
	}
}

func (s SecretPathSelectorMetadataSyncingRegions) WithValue(value []string) *Secret_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Secret_FieldSubPathValue)
}

func (s SecretPathSelectorMetadataSyncingRegions) WithArrayOfValues(values [][]string) *Secret_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Secret_FieldSubPathArrayOfValues)
}

func (s SecretPathSelectorMetadataSyncingRegions) WithItemValue(value string) *Secret_FieldSubPathArrayItemValue {
	return s.FieldPath().WithIArrayItemValue(value).(*Secret_FieldSubPathArrayItemValue)
}

type SecretPathSelectorMetadataLifecycle struct{}

func (SecretPathSelectorMetadataLifecycle) FieldPath() *Secret_FieldSubPath {
	return &Secret_FieldSubPath{
		selector: Secret_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Lifecycle().FieldPath(),
	}
}

func (s SecretPathSelectorMetadataLifecycle) WithValue(value *meta.Lifecycle) *Secret_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Secret_FieldSubPathValue)
}

func (s SecretPathSelectorMetadataLifecycle) WithArrayOfValues(values []*meta.Lifecycle) *Secret_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Secret_FieldSubPathArrayOfValues)
}

func (SecretPathSelectorMetadataLifecycle) State() SecretPathSelectorMetadataLifecycleState {
	return SecretPathSelectorMetadataLifecycleState{}
}

func (SecretPathSelectorMetadataLifecycle) BlockDeletion() SecretPathSelectorMetadataLifecycleBlockDeletion {
	return SecretPathSelectorMetadataLifecycleBlockDeletion{}
}

type SecretPathSelectorMetadataLifecycleState struct{}

func (SecretPathSelectorMetadataLifecycleState) FieldPath() *Secret_FieldSubPath {
	return &Secret_FieldSubPath{
		selector: Secret_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Lifecycle().State().FieldPath(),
	}
}

func (s SecretPathSelectorMetadataLifecycleState) WithValue(value meta.Lifecycle_State) *Secret_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Secret_FieldSubPathValue)
}

func (s SecretPathSelectorMetadataLifecycleState) WithArrayOfValues(values []meta.Lifecycle_State) *Secret_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Secret_FieldSubPathArrayOfValues)
}

type SecretPathSelectorMetadataLifecycleBlockDeletion struct{}

func (SecretPathSelectorMetadataLifecycleBlockDeletion) FieldPath() *Secret_FieldSubPath {
	return &Secret_FieldSubPath{
		selector: Secret_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Lifecycle().BlockDeletion().FieldPath(),
	}
}

func (s SecretPathSelectorMetadataLifecycleBlockDeletion) WithValue(value bool) *Secret_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Secret_FieldSubPathValue)
}

func (s SecretPathSelectorMetadataLifecycleBlockDeletion) WithArrayOfValues(values []bool) *Secret_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Secret_FieldSubPathArrayOfValues)
}

type SecretPathSelectorMetadataServices struct{}

func (SecretPathSelectorMetadataServices) FieldPath() *Secret_FieldSubPath {
	return &Secret_FieldSubPath{
		selector: Secret_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Services().FieldPath(),
	}
}

func (s SecretPathSelectorMetadataServices) WithValue(value *meta.ServicesInfo) *Secret_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Secret_FieldSubPathValue)
}

func (s SecretPathSelectorMetadataServices) WithArrayOfValues(values []*meta.ServicesInfo) *Secret_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Secret_FieldSubPathArrayOfValues)
}

func (SecretPathSelectorMetadataServices) OwningService() SecretPathSelectorMetadataServicesOwningService {
	return SecretPathSelectorMetadataServicesOwningService{}
}

func (SecretPathSelectorMetadataServices) AllowedServices() SecretPathSelectorMetadataServicesAllowedServices {
	return SecretPathSelectorMetadataServicesAllowedServices{}
}

type SecretPathSelectorMetadataServicesOwningService struct{}

func (SecretPathSelectorMetadataServicesOwningService) FieldPath() *Secret_FieldSubPath {
	return &Secret_FieldSubPath{
		selector: Secret_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Services().OwningService().FieldPath(),
	}
}

func (s SecretPathSelectorMetadataServicesOwningService) WithValue(value string) *Secret_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Secret_FieldSubPathValue)
}

func (s SecretPathSelectorMetadataServicesOwningService) WithArrayOfValues(values []string) *Secret_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Secret_FieldSubPathArrayOfValues)
}

type SecretPathSelectorMetadataServicesAllowedServices struct{}

func (SecretPathSelectorMetadataServicesAllowedServices) FieldPath() *Secret_FieldSubPath {
	return &Secret_FieldSubPath{
		selector: Secret_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Services().AllowedServices().FieldPath(),
	}
}

func (s SecretPathSelectorMetadataServicesAllowedServices) WithValue(value []string) *Secret_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Secret_FieldSubPathValue)
}

func (s SecretPathSelectorMetadataServicesAllowedServices) WithArrayOfValues(values [][]string) *Secret_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Secret_FieldSubPathArrayOfValues)
}

func (s SecretPathSelectorMetadataServicesAllowedServices) WithItemValue(value string) *Secret_FieldSubPathArrayItemValue {
	return s.FieldPath().WithIArrayItemValue(value).(*Secret_FieldSubPathArrayItemValue)
}

type SecretPathSelectorEncData struct{}

func (SecretPathSelectorEncData) FieldPath() *Secret_FieldTerminalPath {
	return &Secret_FieldTerminalPath{selector: Secret_FieldPathSelectorEncData}
}

func (s SecretPathSelectorEncData) WithValue(value []byte) *Secret_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*Secret_FieldTerminalPathValue)
}

func (s SecretPathSelectorEncData) WithArrayOfValues(values [][]byte) *Secret_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Secret_FieldTerminalPathArrayOfValues)
}

type SecretPathSelectorData struct{}

func (SecretPathSelectorData) FieldPath() *Secret_FieldTerminalPath {
	return &Secret_FieldTerminalPath{selector: Secret_FieldPathSelectorData}
}

func (s SecretPathSelectorData) WithValue(value map[string]string) *Secret_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*Secret_FieldTerminalPathValue)
}

func (s SecretPathSelectorData) WithArrayOfValues(values []map[string]string) *Secret_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Secret_FieldTerminalPathArrayOfValues)
}

func (SecretPathSelectorData) WithKey(key string) SecretMapPathSelectorData {
	return SecretMapPathSelectorData{key: key}
}

type SecretMapPathSelectorData struct {
	key string
}

func (s SecretMapPathSelectorData) FieldPath() *Secret_FieldPathMap {
	return &Secret_FieldPathMap{selector: Secret_FieldPathSelectorData, key: s.key}
}

func (s SecretMapPathSelectorData) WithValue(value string) *Secret_FieldPathMapValue {
	return s.FieldPath().WithIValue(value).(*Secret_FieldPathMapValue)
}

func (s SecretMapPathSelectorData) WithArrayOfValues(values []string) *Secret_FieldPathMapArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Secret_FieldPathMapArrayOfValues)
}