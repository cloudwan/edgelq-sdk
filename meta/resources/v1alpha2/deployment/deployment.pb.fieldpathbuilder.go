// Code generated by protoc-gen-goten-object
// File: edgelq/meta/proto/v1alpha2/deployment.proto
// DO NOT EDIT!!!

package deployment

// proto imports
import (
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
	region "github.com/cloudwan/edgelq-sdk/meta/resources/v1alpha2/region"
	service "github.com/cloudwan/edgelq-sdk/meta/resources/v1alpha2/service"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
)

// make sure we're using proto imports
var (
	_ = &ntt_meta.Meta{}
	_ = &region.Region{}
	_ = &service.Service{}
	_ = &timestamp.Timestamp{}
)

type DeploymentFieldPathBuilder struct{}

func NewDeploymentFieldPathBuilder() DeploymentFieldPathBuilder {
	return DeploymentFieldPathBuilder{}
}
func (DeploymentFieldPathBuilder) Name() DeploymentPathSelectorName {
	return DeploymentPathSelectorName{}
}
func (DeploymentFieldPathBuilder) Service() DeploymentPathSelectorService {
	return DeploymentPathSelectorService{}
}
func (DeploymentFieldPathBuilder) Metadata() DeploymentPathSelectorMetadata {
	return DeploymentPathSelectorMetadata{}
}

type DeploymentPathSelectorName struct{}

func (DeploymentPathSelectorName) FieldPath() *Deployment_FieldTerminalPath {
	return &Deployment_FieldTerminalPath{selector: Deployment_FieldPathSelectorName}
}

func (s DeploymentPathSelectorName) WithValue(value *Name) *Deployment_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*Deployment_FieldTerminalPathValue)
}

func (s DeploymentPathSelectorName) WithArrayOfValues(values []*Name) *Deployment_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Deployment_FieldTerminalPathArrayOfValues)
}

type DeploymentPathSelectorService struct{}

func (DeploymentPathSelectorService) FieldPath() *Deployment_FieldTerminalPath {
	return &Deployment_FieldTerminalPath{selector: Deployment_FieldPathSelectorService}
}

func (s DeploymentPathSelectorService) WithValue(value *service.Reference) *Deployment_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*Deployment_FieldTerminalPathValue)
}

func (s DeploymentPathSelectorService) WithArrayOfValues(values []*service.Reference) *Deployment_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Deployment_FieldTerminalPathArrayOfValues)
}

type DeploymentPathSelectorMetadata struct{}

func (DeploymentPathSelectorMetadata) FieldPath() *Deployment_FieldTerminalPath {
	return &Deployment_FieldTerminalPath{selector: Deployment_FieldPathSelectorMetadata}
}

func (s DeploymentPathSelectorMetadata) WithValue(value *ntt_meta.Meta) *Deployment_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*Deployment_FieldTerminalPathValue)
}

func (s DeploymentPathSelectorMetadata) WithArrayOfValues(values []*ntt_meta.Meta) *Deployment_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Deployment_FieldTerminalPathArrayOfValues)
}

func (DeploymentPathSelectorMetadata) WithSubPath(subPath ntt_meta.Meta_FieldPath) *Deployment_FieldSubPath {
	return &Deployment_FieldSubPath{selector: Deployment_FieldPathSelectorMetadata, subPath: subPath}
}

func (s DeploymentPathSelectorMetadata) WithSubValue(subPathValue ntt_meta.Meta_FieldPathValue) *Deployment_FieldSubPathValue {
	return &Deployment_FieldSubPathValue{Deployment_FieldPath: s.WithSubPath(subPathValue), subPathValue: subPathValue}
}

func (s DeploymentPathSelectorMetadata) WithSubArrayOfValues(subPathArrayOfValues ntt_meta.Meta_FieldPathArrayOfValues) *Deployment_FieldSubPathArrayOfValues {
	return &Deployment_FieldSubPathArrayOfValues{Deployment_FieldPath: s.WithSubPath(subPathArrayOfValues), subPathArrayOfValues: subPathArrayOfValues}
}

func (s DeploymentPathSelectorMetadata) WithSubArrayItemValue(subPathArrayItemValue ntt_meta.Meta_FieldPathArrayItemValue) *Deployment_FieldSubPathArrayItemValue {
	return &Deployment_FieldSubPathArrayItemValue{Deployment_FieldPath: s.WithSubPath(subPathArrayItemValue), subPathItemValue: subPathArrayItemValue}
}

func (DeploymentPathSelectorMetadata) CreateTime() DeploymentPathSelectorMetadataCreateTime {
	return DeploymentPathSelectorMetadataCreateTime{}
}

func (DeploymentPathSelectorMetadata) UpdateTime() DeploymentPathSelectorMetadataUpdateTime {
	return DeploymentPathSelectorMetadataUpdateTime{}
}

func (DeploymentPathSelectorMetadata) Uuid() DeploymentPathSelectorMetadataUuid {
	return DeploymentPathSelectorMetadataUuid{}
}

func (DeploymentPathSelectorMetadata) Tags() DeploymentPathSelectorMetadataTags {
	return DeploymentPathSelectorMetadataTags{}
}

func (DeploymentPathSelectorMetadata) Labels() DeploymentPathSelectorMetadataLabels {
	return DeploymentPathSelectorMetadataLabels{}
}

func (DeploymentPathSelectorMetadata) Annotations() DeploymentPathSelectorMetadataAnnotations {
	return DeploymentPathSelectorMetadataAnnotations{}
}

func (DeploymentPathSelectorMetadata) Generation() DeploymentPathSelectorMetadataGeneration {
	return DeploymentPathSelectorMetadataGeneration{}
}

func (DeploymentPathSelectorMetadata) ResourceVersion() DeploymentPathSelectorMetadataResourceVersion {
	return DeploymentPathSelectorMetadataResourceVersion{}
}

func (DeploymentPathSelectorMetadata) OwnerReferences() DeploymentPathSelectorMetadataOwnerReferences {
	return DeploymentPathSelectorMetadataOwnerReferences{}
}

func (DeploymentPathSelectorMetadata) Shards() DeploymentPathSelectorMetadataShards {
	return DeploymentPathSelectorMetadataShards{}
}

func (DeploymentPathSelectorMetadata) Syncing() DeploymentPathSelectorMetadataSyncing {
	return DeploymentPathSelectorMetadataSyncing{}
}

type DeploymentPathSelectorMetadataCreateTime struct{}

func (DeploymentPathSelectorMetadataCreateTime) FieldPath() *Deployment_FieldSubPath {
	return &Deployment_FieldSubPath{
		selector: Deployment_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().CreateTime().FieldPath(),
	}
}

func (s DeploymentPathSelectorMetadataCreateTime) WithValue(value *timestamp.Timestamp) *Deployment_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Deployment_FieldSubPathValue)
}

func (s DeploymentPathSelectorMetadataCreateTime) WithArrayOfValues(values []*timestamp.Timestamp) *Deployment_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Deployment_FieldSubPathArrayOfValues)
}

type DeploymentPathSelectorMetadataUpdateTime struct{}

func (DeploymentPathSelectorMetadataUpdateTime) FieldPath() *Deployment_FieldSubPath {
	return &Deployment_FieldSubPath{
		selector: Deployment_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().UpdateTime().FieldPath(),
	}
}

func (s DeploymentPathSelectorMetadataUpdateTime) WithValue(value *timestamp.Timestamp) *Deployment_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Deployment_FieldSubPathValue)
}

func (s DeploymentPathSelectorMetadataUpdateTime) WithArrayOfValues(values []*timestamp.Timestamp) *Deployment_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Deployment_FieldSubPathArrayOfValues)
}

type DeploymentPathSelectorMetadataUuid struct{}

func (DeploymentPathSelectorMetadataUuid) FieldPath() *Deployment_FieldSubPath {
	return &Deployment_FieldSubPath{
		selector: Deployment_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Uuid().FieldPath(),
	}
}

func (s DeploymentPathSelectorMetadataUuid) WithValue(value string) *Deployment_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Deployment_FieldSubPathValue)
}

func (s DeploymentPathSelectorMetadataUuid) WithArrayOfValues(values []string) *Deployment_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Deployment_FieldSubPathArrayOfValues)
}

type DeploymentPathSelectorMetadataTags struct{}

func (DeploymentPathSelectorMetadataTags) FieldPath() *Deployment_FieldSubPath {
	return &Deployment_FieldSubPath{
		selector: Deployment_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Tags().FieldPath(),
	}
}

func (s DeploymentPathSelectorMetadataTags) WithValue(value []string) *Deployment_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Deployment_FieldSubPathValue)
}

func (s DeploymentPathSelectorMetadataTags) WithArrayOfValues(values [][]string) *Deployment_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Deployment_FieldSubPathArrayOfValues)
}

func (s DeploymentPathSelectorMetadataTags) WithItemValue(value string) *Deployment_FieldSubPathArrayItemValue {
	return s.FieldPath().WithIArrayItemValue(value).(*Deployment_FieldSubPathArrayItemValue)
}

type DeploymentPathSelectorMetadataLabels struct{}

func (DeploymentPathSelectorMetadataLabels) FieldPath() *Deployment_FieldSubPath {
	return &Deployment_FieldSubPath{
		selector: Deployment_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Labels().FieldPath(),
	}
}

func (s DeploymentPathSelectorMetadataLabels) WithValue(value map[string]string) *Deployment_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Deployment_FieldSubPathValue)
}

func (s DeploymentPathSelectorMetadataLabels) WithArrayOfValues(values []map[string]string) *Deployment_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Deployment_FieldSubPathArrayOfValues)
}

func (DeploymentPathSelectorMetadataLabels) WithKey(key string) DeploymentMapPathSelectorMetadataLabels {
	return DeploymentMapPathSelectorMetadataLabels{key: key}
}

type DeploymentMapPathSelectorMetadataLabels struct {
	key string
}

func (s DeploymentMapPathSelectorMetadataLabels) FieldPath() *Deployment_FieldSubPath {
	return &Deployment_FieldSubPath{
		selector: Deployment_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Labels().WithKey(s.key).FieldPath(),
	}
}

func (s DeploymentMapPathSelectorMetadataLabels) WithValue(value string) *Deployment_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Deployment_FieldSubPathValue)
}

func (s DeploymentMapPathSelectorMetadataLabels) WithArrayOfValues(values []string) *Deployment_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Deployment_FieldSubPathArrayOfValues)
}

type DeploymentPathSelectorMetadataAnnotations struct{}

func (DeploymentPathSelectorMetadataAnnotations) FieldPath() *Deployment_FieldSubPath {
	return &Deployment_FieldSubPath{
		selector: Deployment_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Annotations().FieldPath(),
	}
}

func (s DeploymentPathSelectorMetadataAnnotations) WithValue(value map[string]string) *Deployment_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Deployment_FieldSubPathValue)
}

func (s DeploymentPathSelectorMetadataAnnotations) WithArrayOfValues(values []map[string]string) *Deployment_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Deployment_FieldSubPathArrayOfValues)
}

func (DeploymentPathSelectorMetadataAnnotations) WithKey(key string) DeploymentMapPathSelectorMetadataAnnotations {
	return DeploymentMapPathSelectorMetadataAnnotations{key: key}
}

type DeploymentMapPathSelectorMetadataAnnotations struct {
	key string
}

func (s DeploymentMapPathSelectorMetadataAnnotations) FieldPath() *Deployment_FieldSubPath {
	return &Deployment_FieldSubPath{
		selector: Deployment_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Annotations().WithKey(s.key).FieldPath(),
	}
}

func (s DeploymentMapPathSelectorMetadataAnnotations) WithValue(value string) *Deployment_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Deployment_FieldSubPathValue)
}

func (s DeploymentMapPathSelectorMetadataAnnotations) WithArrayOfValues(values []string) *Deployment_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Deployment_FieldSubPathArrayOfValues)
}

type DeploymentPathSelectorMetadataGeneration struct{}

func (DeploymentPathSelectorMetadataGeneration) FieldPath() *Deployment_FieldSubPath {
	return &Deployment_FieldSubPath{
		selector: Deployment_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Generation().FieldPath(),
	}
}

func (s DeploymentPathSelectorMetadataGeneration) WithValue(value int64) *Deployment_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Deployment_FieldSubPathValue)
}

func (s DeploymentPathSelectorMetadataGeneration) WithArrayOfValues(values []int64) *Deployment_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Deployment_FieldSubPathArrayOfValues)
}

type DeploymentPathSelectorMetadataResourceVersion struct{}

func (DeploymentPathSelectorMetadataResourceVersion) FieldPath() *Deployment_FieldSubPath {
	return &Deployment_FieldSubPath{
		selector: Deployment_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().ResourceVersion().FieldPath(),
	}
}

func (s DeploymentPathSelectorMetadataResourceVersion) WithValue(value string) *Deployment_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Deployment_FieldSubPathValue)
}

func (s DeploymentPathSelectorMetadataResourceVersion) WithArrayOfValues(values []string) *Deployment_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Deployment_FieldSubPathArrayOfValues)
}

type DeploymentPathSelectorMetadataOwnerReferences struct{}

func (DeploymentPathSelectorMetadataOwnerReferences) FieldPath() *Deployment_FieldSubPath {
	return &Deployment_FieldSubPath{
		selector: Deployment_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().FieldPath(),
	}
}

func (s DeploymentPathSelectorMetadataOwnerReferences) WithValue(value []*ntt_meta.OwnerReference) *Deployment_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Deployment_FieldSubPathValue)
}

func (s DeploymentPathSelectorMetadataOwnerReferences) WithArrayOfValues(values [][]*ntt_meta.OwnerReference) *Deployment_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Deployment_FieldSubPathArrayOfValues)
}

func (s DeploymentPathSelectorMetadataOwnerReferences) WithItemValue(value *ntt_meta.OwnerReference) *Deployment_FieldSubPathArrayItemValue {
	return s.FieldPath().WithIArrayItemValue(value).(*Deployment_FieldSubPathArrayItemValue)
}

func (DeploymentPathSelectorMetadataOwnerReferences) ApiVersion() DeploymentPathSelectorMetadataOwnerReferencesApiVersion {
	return DeploymentPathSelectorMetadataOwnerReferencesApiVersion{}
}

func (DeploymentPathSelectorMetadataOwnerReferences) Kind() DeploymentPathSelectorMetadataOwnerReferencesKind {
	return DeploymentPathSelectorMetadataOwnerReferencesKind{}
}

func (DeploymentPathSelectorMetadataOwnerReferences) Name() DeploymentPathSelectorMetadataOwnerReferencesName {
	return DeploymentPathSelectorMetadataOwnerReferencesName{}
}

func (DeploymentPathSelectorMetadataOwnerReferences) Uid() DeploymentPathSelectorMetadataOwnerReferencesUid {
	return DeploymentPathSelectorMetadataOwnerReferencesUid{}
}

func (DeploymentPathSelectorMetadataOwnerReferences) Controller() DeploymentPathSelectorMetadataOwnerReferencesController {
	return DeploymentPathSelectorMetadataOwnerReferencesController{}
}

func (DeploymentPathSelectorMetadataOwnerReferences) BlockOwnerDeletion() DeploymentPathSelectorMetadataOwnerReferencesBlockOwnerDeletion {
	return DeploymentPathSelectorMetadataOwnerReferencesBlockOwnerDeletion{}
}

type DeploymentPathSelectorMetadataOwnerReferencesApiVersion struct{}

func (DeploymentPathSelectorMetadataOwnerReferencesApiVersion) FieldPath() *Deployment_FieldSubPath {
	return &Deployment_FieldSubPath{
		selector: Deployment_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().ApiVersion().FieldPath(),
	}
}

func (s DeploymentPathSelectorMetadataOwnerReferencesApiVersion) WithValue(value string) *Deployment_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Deployment_FieldSubPathValue)
}

func (s DeploymentPathSelectorMetadataOwnerReferencesApiVersion) WithArrayOfValues(values []string) *Deployment_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Deployment_FieldSubPathArrayOfValues)
}

type DeploymentPathSelectorMetadataOwnerReferencesKind struct{}

func (DeploymentPathSelectorMetadataOwnerReferencesKind) FieldPath() *Deployment_FieldSubPath {
	return &Deployment_FieldSubPath{
		selector: Deployment_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().Kind().FieldPath(),
	}
}

func (s DeploymentPathSelectorMetadataOwnerReferencesKind) WithValue(value string) *Deployment_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Deployment_FieldSubPathValue)
}

func (s DeploymentPathSelectorMetadataOwnerReferencesKind) WithArrayOfValues(values []string) *Deployment_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Deployment_FieldSubPathArrayOfValues)
}

type DeploymentPathSelectorMetadataOwnerReferencesName struct{}

func (DeploymentPathSelectorMetadataOwnerReferencesName) FieldPath() *Deployment_FieldSubPath {
	return &Deployment_FieldSubPath{
		selector: Deployment_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().Name().FieldPath(),
	}
}

func (s DeploymentPathSelectorMetadataOwnerReferencesName) WithValue(value string) *Deployment_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Deployment_FieldSubPathValue)
}

func (s DeploymentPathSelectorMetadataOwnerReferencesName) WithArrayOfValues(values []string) *Deployment_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Deployment_FieldSubPathArrayOfValues)
}

type DeploymentPathSelectorMetadataOwnerReferencesUid struct{}

func (DeploymentPathSelectorMetadataOwnerReferencesUid) FieldPath() *Deployment_FieldSubPath {
	return &Deployment_FieldSubPath{
		selector: Deployment_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().Uid().FieldPath(),
	}
}

func (s DeploymentPathSelectorMetadataOwnerReferencesUid) WithValue(value string) *Deployment_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Deployment_FieldSubPathValue)
}

func (s DeploymentPathSelectorMetadataOwnerReferencesUid) WithArrayOfValues(values []string) *Deployment_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Deployment_FieldSubPathArrayOfValues)
}

type DeploymentPathSelectorMetadataOwnerReferencesController struct{}

func (DeploymentPathSelectorMetadataOwnerReferencesController) FieldPath() *Deployment_FieldSubPath {
	return &Deployment_FieldSubPath{
		selector: Deployment_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().Controller().FieldPath(),
	}
}

func (s DeploymentPathSelectorMetadataOwnerReferencesController) WithValue(value bool) *Deployment_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Deployment_FieldSubPathValue)
}

func (s DeploymentPathSelectorMetadataOwnerReferencesController) WithArrayOfValues(values []bool) *Deployment_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Deployment_FieldSubPathArrayOfValues)
}

type DeploymentPathSelectorMetadataOwnerReferencesBlockOwnerDeletion struct{}

func (DeploymentPathSelectorMetadataOwnerReferencesBlockOwnerDeletion) FieldPath() *Deployment_FieldSubPath {
	return &Deployment_FieldSubPath{
		selector: Deployment_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().BlockOwnerDeletion().FieldPath(),
	}
}

func (s DeploymentPathSelectorMetadataOwnerReferencesBlockOwnerDeletion) WithValue(value bool) *Deployment_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Deployment_FieldSubPathValue)
}

func (s DeploymentPathSelectorMetadataOwnerReferencesBlockOwnerDeletion) WithArrayOfValues(values []bool) *Deployment_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Deployment_FieldSubPathArrayOfValues)
}

type DeploymentPathSelectorMetadataShards struct{}

func (DeploymentPathSelectorMetadataShards) FieldPath() *Deployment_FieldSubPath {
	return &Deployment_FieldSubPath{
		selector: Deployment_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Shards().FieldPath(),
	}
}

func (s DeploymentPathSelectorMetadataShards) WithValue(value map[string]int64) *Deployment_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Deployment_FieldSubPathValue)
}

func (s DeploymentPathSelectorMetadataShards) WithArrayOfValues(values []map[string]int64) *Deployment_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Deployment_FieldSubPathArrayOfValues)
}

func (DeploymentPathSelectorMetadataShards) WithKey(key string) DeploymentMapPathSelectorMetadataShards {
	return DeploymentMapPathSelectorMetadataShards{key: key}
}

type DeploymentMapPathSelectorMetadataShards struct {
	key string
}

func (s DeploymentMapPathSelectorMetadataShards) FieldPath() *Deployment_FieldSubPath {
	return &Deployment_FieldSubPath{
		selector: Deployment_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Shards().WithKey(s.key).FieldPath(),
	}
}

func (s DeploymentMapPathSelectorMetadataShards) WithValue(value int64) *Deployment_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Deployment_FieldSubPathValue)
}

func (s DeploymentMapPathSelectorMetadataShards) WithArrayOfValues(values []int64) *Deployment_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Deployment_FieldSubPathArrayOfValues)
}

type DeploymentPathSelectorMetadataSyncing struct{}

func (DeploymentPathSelectorMetadataSyncing) FieldPath() *Deployment_FieldSubPath {
	return &Deployment_FieldSubPath{
		selector: Deployment_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Syncing().FieldPath(),
	}
}

func (s DeploymentPathSelectorMetadataSyncing) WithValue(value *ntt_meta.SyncingMeta) *Deployment_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Deployment_FieldSubPathValue)
}

func (s DeploymentPathSelectorMetadataSyncing) WithArrayOfValues(values []*ntt_meta.SyncingMeta) *Deployment_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Deployment_FieldSubPathArrayOfValues)
}

func (DeploymentPathSelectorMetadataSyncing) OwningRegion() DeploymentPathSelectorMetadataSyncingOwningRegion {
	return DeploymentPathSelectorMetadataSyncingOwningRegion{}
}

func (DeploymentPathSelectorMetadataSyncing) Regions() DeploymentPathSelectorMetadataSyncingRegions {
	return DeploymentPathSelectorMetadataSyncingRegions{}
}

type DeploymentPathSelectorMetadataSyncingOwningRegion struct{}

func (DeploymentPathSelectorMetadataSyncingOwningRegion) FieldPath() *Deployment_FieldSubPath {
	return &Deployment_FieldSubPath{
		selector: Deployment_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Syncing().OwningRegion().FieldPath(),
	}
}

func (s DeploymentPathSelectorMetadataSyncingOwningRegion) WithValue(value string) *Deployment_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Deployment_FieldSubPathValue)
}

func (s DeploymentPathSelectorMetadataSyncingOwningRegion) WithArrayOfValues(values []string) *Deployment_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Deployment_FieldSubPathArrayOfValues)
}

type DeploymentPathSelectorMetadataSyncingRegions struct{}

func (DeploymentPathSelectorMetadataSyncingRegions) FieldPath() *Deployment_FieldSubPath {
	return &Deployment_FieldSubPath{
		selector: Deployment_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Syncing().Regions().FieldPath(),
	}
}

func (s DeploymentPathSelectorMetadataSyncingRegions) WithValue(value []string) *Deployment_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Deployment_FieldSubPathValue)
}

func (s DeploymentPathSelectorMetadataSyncingRegions) WithArrayOfValues(values [][]string) *Deployment_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Deployment_FieldSubPathArrayOfValues)
}

func (s DeploymentPathSelectorMetadataSyncingRegions) WithItemValue(value string) *Deployment_FieldSubPathArrayItemValue {
	return s.FieldPath().WithIArrayItemValue(value).(*Deployment_FieldSubPathArrayItemValue)
}
