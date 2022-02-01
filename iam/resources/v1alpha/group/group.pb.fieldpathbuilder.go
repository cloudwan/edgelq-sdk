// Code generated by protoc-gen-goten-object
// File: edgelq/iam/proto/v1alpha/group.proto
// DO NOT EDIT!!!

package group

// proto imports
import (
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
	organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/organization"
	project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/project"
	syncing_meta "github.com/cloudwan/edgelq-sdk/meta/multi_region/proto/syncing_meta"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
)

// make sure we're using proto imports
var (
	_ = &ntt_meta.Meta{}
	_ = &organization.Organization{}
	_ = &project.Project{}
	_ = &syncing_meta.SyncingMeta{}
	_ = &timestamp.Timestamp{}
)

type GroupFieldPathBuilder struct{}

func NewGroupFieldPathBuilder() GroupFieldPathBuilder {
	return GroupFieldPathBuilder{}
}
func (GroupFieldPathBuilder) Name() GroupPathSelectorName {
	return GroupPathSelectorName{}
}
func (GroupFieldPathBuilder) DisplayName() GroupPathSelectorDisplayName {
	return GroupPathSelectorDisplayName{}
}
func (GroupFieldPathBuilder) Email() GroupPathSelectorEmail {
	return GroupPathSelectorEmail{}
}
func (GroupFieldPathBuilder) Metadata() GroupPathSelectorMetadata {
	return GroupPathSelectorMetadata{}
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

type GroupPathSelectorMetadata struct{}

func (GroupPathSelectorMetadata) FieldPath() *Group_FieldTerminalPath {
	return &Group_FieldTerminalPath{selector: Group_FieldPathSelectorMetadata}
}

func (s GroupPathSelectorMetadata) WithValue(value *ntt_meta.Meta) *Group_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*Group_FieldTerminalPathValue)
}

func (s GroupPathSelectorMetadata) WithArrayOfValues(values []*ntt_meta.Meta) *Group_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Group_FieldTerminalPathArrayOfValues)
}

func (GroupPathSelectorMetadata) WithSubPath(subPath ntt_meta.Meta_FieldPath) *Group_FieldSubPath {
	return &Group_FieldSubPath{selector: Group_FieldPathSelectorMetadata, subPath: subPath}
}

func (s GroupPathSelectorMetadata) WithSubValue(subPathValue ntt_meta.Meta_FieldPathValue) *Group_FieldSubPathValue {
	return &Group_FieldSubPathValue{Group_FieldPath: s.WithSubPath(subPathValue), subPathValue: subPathValue}
}

func (s GroupPathSelectorMetadata) WithSubArrayOfValues(subPathArrayOfValues ntt_meta.Meta_FieldPathArrayOfValues) *Group_FieldSubPathArrayOfValues {
	return &Group_FieldSubPathArrayOfValues{Group_FieldPath: s.WithSubPath(subPathArrayOfValues), subPathArrayOfValues: subPathArrayOfValues}
}

func (s GroupPathSelectorMetadata) WithSubArrayItemValue(subPathArrayItemValue ntt_meta.Meta_FieldPathArrayItemValue) *Group_FieldSubPathArrayItemValue {
	return &Group_FieldSubPathArrayItemValue{Group_FieldPath: s.WithSubPath(subPathArrayItemValue), subPathItemValue: subPathArrayItemValue}
}

func (GroupPathSelectorMetadata) CreateTime() GroupPathSelectorMetadataCreateTime {
	return GroupPathSelectorMetadataCreateTime{}
}

func (GroupPathSelectorMetadata) UpdateTime() GroupPathSelectorMetadataUpdateTime {
	return GroupPathSelectorMetadataUpdateTime{}
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

type GroupPathSelectorMetadataCreateTime struct{}

func (GroupPathSelectorMetadataCreateTime) FieldPath() *Group_FieldSubPath {
	return &Group_FieldSubPath{
		selector: Group_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().CreateTime().FieldPath(),
	}
}

func (s GroupPathSelectorMetadataCreateTime) WithValue(value *timestamp.Timestamp) *Group_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Group_FieldSubPathValue)
}

func (s GroupPathSelectorMetadataCreateTime) WithArrayOfValues(values []*timestamp.Timestamp) *Group_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Group_FieldSubPathArrayOfValues)
}

type GroupPathSelectorMetadataUpdateTime struct{}

func (GroupPathSelectorMetadataUpdateTime) FieldPath() *Group_FieldSubPath {
	return &Group_FieldSubPath{
		selector: Group_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().UpdateTime().FieldPath(),
	}
}

func (s GroupPathSelectorMetadataUpdateTime) WithValue(value *timestamp.Timestamp) *Group_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Group_FieldSubPathValue)
}

func (s GroupPathSelectorMetadataUpdateTime) WithArrayOfValues(values []*timestamp.Timestamp) *Group_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Group_FieldSubPathArrayOfValues)
}

type GroupPathSelectorMetadataUuid struct{}

func (GroupPathSelectorMetadataUuid) FieldPath() *Group_FieldSubPath {
	return &Group_FieldSubPath{
		selector: Group_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Uuid().FieldPath(),
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
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Tags().FieldPath(),
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
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Labels().FieldPath(),
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
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Labels().WithKey(s.key).FieldPath(),
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
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Annotations().FieldPath(),
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
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Annotations().WithKey(s.key).FieldPath(),
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
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Generation().FieldPath(),
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
		subPath:  ntt_meta.NewMetaFieldPathBuilder().ResourceVersion().FieldPath(),
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
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().FieldPath(),
	}
}

func (s GroupPathSelectorMetadataOwnerReferences) WithValue(value []*ntt_meta.OwnerReference) *Group_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Group_FieldSubPathValue)
}

func (s GroupPathSelectorMetadataOwnerReferences) WithArrayOfValues(values [][]*ntt_meta.OwnerReference) *Group_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Group_FieldSubPathArrayOfValues)
}

func (s GroupPathSelectorMetadataOwnerReferences) WithItemValue(value *ntt_meta.OwnerReference) *Group_FieldSubPathArrayItemValue {
	return s.FieldPath().WithIArrayItemValue(value).(*Group_FieldSubPathArrayItemValue)
}

func (GroupPathSelectorMetadataOwnerReferences) ApiVersion() GroupPathSelectorMetadataOwnerReferencesApiVersion {
	return GroupPathSelectorMetadataOwnerReferencesApiVersion{}
}

func (GroupPathSelectorMetadataOwnerReferences) Kind() GroupPathSelectorMetadataOwnerReferencesKind {
	return GroupPathSelectorMetadataOwnerReferencesKind{}
}

func (GroupPathSelectorMetadataOwnerReferences) Name() GroupPathSelectorMetadataOwnerReferencesName {
	return GroupPathSelectorMetadataOwnerReferencesName{}
}

func (GroupPathSelectorMetadataOwnerReferences) Uid() GroupPathSelectorMetadataOwnerReferencesUid {
	return GroupPathSelectorMetadataOwnerReferencesUid{}
}

func (GroupPathSelectorMetadataOwnerReferences) Controller() GroupPathSelectorMetadataOwnerReferencesController {
	return GroupPathSelectorMetadataOwnerReferencesController{}
}

func (GroupPathSelectorMetadataOwnerReferences) BlockOwnerDeletion() GroupPathSelectorMetadataOwnerReferencesBlockOwnerDeletion {
	return GroupPathSelectorMetadataOwnerReferencesBlockOwnerDeletion{}
}

type GroupPathSelectorMetadataOwnerReferencesApiVersion struct{}

func (GroupPathSelectorMetadataOwnerReferencesApiVersion) FieldPath() *Group_FieldSubPath {
	return &Group_FieldSubPath{
		selector: Group_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().ApiVersion().FieldPath(),
	}
}

func (s GroupPathSelectorMetadataOwnerReferencesApiVersion) WithValue(value string) *Group_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Group_FieldSubPathValue)
}

func (s GroupPathSelectorMetadataOwnerReferencesApiVersion) WithArrayOfValues(values []string) *Group_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Group_FieldSubPathArrayOfValues)
}

type GroupPathSelectorMetadataOwnerReferencesKind struct{}

func (GroupPathSelectorMetadataOwnerReferencesKind) FieldPath() *Group_FieldSubPath {
	return &Group_FieldSubPath{
		selector: Group_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().Kind().FieldPath(),
	}
}

func (s GroupPathSelectorMetadataOwnerReferencesKind) WithValue(value string) *Group_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Group_FieldSubPathValue)
}

func (s GroupPathSelectorMetadataOwnerReferencesKind) WithArrayOfValues(values []string) *Group_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Group_FieldSubPathArrayOfValues)
}

type GroupPathSelectorMetadataOwnerReferencesName struct{}

func (GroupPathSelectorMetadataOwnerReferencesName) FieldPath() *Group_FieldSubPath {
	return &Group_FieldSubPath{
		selector: Group_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().Name().FieldPath(),
	}
}

func (s GroupPathSelectorMetadataOwnerReferencesName) WithValue(value string) *Group_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Group_FieldSubPathValue)
}

func (s GroupPathSelectorMetadataOwnerReferencesName) WithArrayOfValues(values []string) *Group_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Group_FieldSubPathArrayOfValues)
}

type GroupPathSelectorMetadataOwnerReferencesUid struct{}

func (GroupPathSelectorMetadataOwnerReferencesUid) FieldPath() *Group_FieldSubPath {
	return &Group_FieldSubPath{
		selector: Group_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().Uid().FieldPath(),
	}
}

func (s GroupPathSelectorMetadataOwnerReferencesUid) WithValue(value string) *Group_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Group_FieldSubPathValue)
}

func (s GroupPathSelectorMetadataOwnerReferencesUid) WithArrayOfValues(values []string) *Group_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Group_FieldSubPathArrayOfValues)
}

type GroupPathSelectorMetadataOwnerReferencesController struct{}

func (GroupPathSelectorMetadataOwnerReferencesController) FieldPath() *Group_FieldSubPath {
	return &Group_FieldSubPath{
		selector: Group_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().Controller().FieldPath(),
	}
}

func (s GroupPathSelectorMetadataOwnerReferencesController) WithValue(value bool) *Group_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Group_FieldSubPathValue)
}

func (s GroupPathSelectorMetadataOwnerReferencesController) WithArrayOfValues(values []bool) *Group_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Group_FieldSubPathArrayOfValues)
}

type GroupPathSelectorMetadataOwnerReferencesBlockOwnerDeletion struct{}

func (GroupPathSelectorMetadataOwnerReferencesBlockOwnerDeletion) FieldPath() *Group_FieldSubPath {
	return &Group_FieldSubPath{
		selector: Group_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().BlockOwnerDeletion().FieldPath(),
	}
}

func (s GroupPathSelectorMetadataOwnerReferencesBlockOwnerDeletion) WithValue(value bool) *Group_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Group_FieldSubPathValue)
}

func (s GroupPathSelectorMetadataOwnerReferencesBlockOwnerDeletion) WithArrayOfValues(values []bool) *Group_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Group_FieldSubPathArrayOfValues)
}

type GroupPathSelectorMetadataShards struct{}

func (GroupPathSelectorMetadataShards) FieldPath() *Group_FieldSubPath {
	return &Group_FieldSubPath{
		selector: Group_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Shards().FieldPath(),
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
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Shards().WithKey(s.key).FieldPath(),
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
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Syncing().FieldPath(),
	}
}

func (s GroupPathSelectorMetadataSyncing) WithValue(value *syncing_meta.SyncingMeta) *Group_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Group_FieldSubPathValue)
}

func (s GroupPathSelectorMetadataSyncing) WithArrayOfValues(values []*syncing_meta.SyncingMeta) *Group_FieldSubPathArrayOfValues {
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
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Syncing().OwningRegion().FieldPath(),
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
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Syncing().Regions().FieldPath(),
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
