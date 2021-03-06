// Code generated by protoc-gen-goten-object
// File: edgelq/iam/proto/v1alpha/service_account.proto
// DO NOT EDIT!!!

package service_account

// proto imports
import (
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
	organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/organization"
	project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/project"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
)

// make sure we're using proto imports
var (
	_ = &ntt_meta.Meta{}
	_ = &organization.Organization{}
	_ = &project.Project{}
	_ = &timestamp.Timestamp{}
)

type ServiceAccountFieldPathBuilder struct{}

func NewServiceAccountFieldPathBuilder() ServiceAccountFieldPathBuilder {
	return ServiceAccountFieldPathBuilder{}
}
func (ServiceAccountFieldPathBuilder) Name() ServiceAccountPathSelectorName {
	return ServiceAccountPathSelectorName{}
}
func (ServiceAccountFieldPathBuilder) DisplayName() ServiceAccountPathSelectorDisplayName {
	return ServiceAccountPathSelectorDisplayName{}
}
func (ServiceAccountFieldPathBuilder) Email() ServiceAccountPathSelectorEmail {
	return ServiceAccountPathSelectorEmail{}
}
func (ServiceAccountFieldPathBuilder) Metadata() ServiceAccountPathSelectorMetadata {
	return ServiceAccountPathSelectorMetadata{}
}

type ServiceAccountPathSelectorName struct{}

func (ServiceAccountPathSelectorName) FieldPath() *ServiceAccount_FieldTerminalPath {
	return &ServiceAccount_FieldTerminalPath{selector: ServiceAccount_FieldPathSelectorName}
}

func (s ServiceAccountPathSelectorName) WithValue(value *Name) *ServiceAccount_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*ServiceAccount_FieldTerminalPathValue)
}

func (s ServiceAccountPathSelectorName) WithArrayOfValues(values []*Name) *ServiceAccount_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ServiceAccount_FieldTerminalPathArrayOfValues)
}

type ServiceAccountPathSelectorDisplayName struct{}

func (ServiceAccountPathSelectorDisplayName) FieldPath() *ServiceAccount_FieldTerminalPath {
	return &ServiceAccount_FieldTerminalPath{selector: ServiceAccount_FieldPathSelectorDisplayName}
}

func (s ServiceAccountPathSelectorDisplayName) WithValue(value string) *ServiceAccount_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*ServiceAccount_FieldTerminalPathValue)
}

func (s ServiceAccountPathSelectorDisplayName) WithArrayOfValues(values []string) *ServiceAccount_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ServiceAccount_FieldTerminalPathArrayOfValues)
}

type ServiceAccountPathSelectorEmail struct{}

func (ServiceAccountPathSelectorEmail) FieldPath() *ServiceAccount_FieldTerminalPath {
	return &ServiceAccount_FieldTerminalPath{selector: ServiceAccount_FieldPathSelectorEmail}
}

func (s ServiceAccountPathSelectorEmail) WithValue(value string) *ServiceAccount_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*ServiceAccount_FieldTerminalPathValue)
}

func (s ServiceAccountPathSelectorEmail) WithArrayOfValues(values []string) *ServiceAccount_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ServiceAccount_FieldTerminalPathArrayOfValues)
}

type ServiceAccountPathSelectorMetadata struct{}

func (ServiceAccountPathSelectorMetadata) FieldPath() *ServiceAccount_FieldTerminalPath {
	return &ServiceAccount_FieldTerminalPath{selector: ServiceAccount_FieldPathSelectorMetadata}
}

func (s ServiceAccountPathSelectorMetadata) WithValue(value *ntt_meta.Meta) *ServiceAccount_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*ServiceAccount_FieldTerminalPathValue)
}

func (s ServiceAccountPathSelectorMetadata) WithArrayOfValues(values []*ntt_meta.Meta) *ServiceAccount_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ServiceAccount_FieldTerminalPathArrayOfValues)
}

func (ServiceAccountPathSelectorMetadata) WithSubPath(subPath ntt_meta.Meta_FieldPath) *ServiceAccount_FieldSubPath {
	return &ServiceAccount_FieldSubPath{selector: ServiceAccount_FieldPathSelectorMetadata, subPath: subPath}
}

func (s ServiceAccountPathSelectorMetadata) WithSubValue(subPathValue ntt_meta.Meta_FieldPathValue) *ServiceAccount_FieldSubPathValue {
	return &ServiceAccount_FieldSubPathValue{ServiceAccount_FieldPath: s.WithSubPath(subPathValue), subPathValue: subPathValue}
}

func (s ServiceAccountPathSelectorMetadata) WithSubArrayOfValues(subPathArrayOfValues ntt_meta.Meta_FieldPathArrayOfValues) *ServiceAccount_FieldSubPathArrayOfValues {
	return &ServiceAccount_FieldSubPathArrayOfValues{ServiceAccount_FieldPath: s.WithSubPath(subPathArrayOfValues), subPathArrayOfValues: subPathArrayOfValues}
}

func (s ServiceAccountPathSelectorMetadata) WithSubArrayItemValue(subPathArrayItemValue ntt_meta.Meta_FieldPathArrayItemValue) *ServiceAccount_FieldSubPathArrayItemValue {
	return &ServiceAccount_FieldSubPathArrayItemValue{ServiceAccount_FieldPath: s.WithSubPath(subPathArrayItemValue), subPathItemValue: subPathArrayItemValue}
}

func (ServiceAccountPathSelectorMetadata) CreateTime() ServiceAccountPathSelectorMetadataCreateTime {
	return ServiceAccountPathSelectorMetadataCreateTime{}
}

func (ServiceAccountPathSelectorMetadata) UpdateTime() ServiceAccountPathSelectorMetadataUpdateTime {
	return ServiceAccountPathSelectorMetadataUpdateTime{}
}

func (ServiceAccountPathSelectorMetadata) Uuid() ServiceAccountPathSelectorMetadataUuid {
	return ServiceAccountPathSelectorMetadataUuid{}
}

func (ServiceAccountPathSelectorMetadata) Tags() ServiceAccountPathSelectorMetadataTags {
	return ServiceAccountPathSelectorMetadataTags{}
}

func (ServiceAccountPathSelectorMetadata) Labels() ServiceAccountPathSelectorMetadataLabels {
	return ServiceAccountPathSelectorMetadataLabels{}
}

func (ServiceAccountPathSelectorMetadata) Annotations() ServiceAccountPathSelectorMetadataAnnotations {
	return ServiceAccountPathSelectorMetadataAnnotations{}
}

func (ServiceAccountPathSelectorMetadata) Generation() ServiceAccountPathSelectorMetadataGeneration {
	return ServiceAccountPathSelectorMetadataGeneration{}
}

func (ServiceAccountPathSelectorMetadata) ResourceVersion() ServiceAccountPathSelectorMetadataResourceVersion {
	return ServiceAccountPathSelectorMetadataResourceVersion{}
}

func (ServiceAccountPathSelectorMetadata) OwnerReferences() ServiceAccountPathSelectorMetadataOwnerReferences {
	return ServiceAccountPathSelectorMetadataOwnerReferences{}
}

func (ServiceAccountPathSelectorMetadata) Shards() ServiceAccountPathSelectorMetadataShards {
	return ServiceAccountPathSelectorMetadataShards{}
}

func (ServiceAccountPathSelectorMetadata) Syncing() ServiceAccountPathSelectorMetadataSyncing {
	return ServiceAccountPathSelectorMetadataSyncing{}
}

type ServiceAccountPathSelectorMetadataCreateTime struct{}

func (ServiceAccountPathSelectorMetadataCreateTime) FieldPath() *ServiceAccount_FieldSubPath {
	return &ServiceAccount_FieldSubPath{
		selector: ServiceAccount_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().CreateTime().FieldPath(),
	}
}

func (s ServiceAccountPathSelectorMetadataCreateTime) WithValue(value *timestamp.Timestamp) *ServiceAccount_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ServiceAccount_FieldSubPathValue)
}

func (s ServiceAccountPathSelectorMetadataCreateTime) WithArrayOfValues(values []*timestamp.Timestamp) *ServiceAccount_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ServiceAccount_FieldSubPathArrayOfValues)
}

type ServiceAccountPathSelectorMetadataUpdateTime struct{}

func (ServiceAccountPathSelectorMetadataUpdateTime) FieldPath() *ServiceAccount_FieldSubPath {
	return &ServiceAccount_FieldSubPath{
		selector: ServiceAccount_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().UpdateTime().FieldPath(),
	}
}

func (s ServiceAccountPathSelectorMetadataUpdateTime) WithValue(value *timestamp.Timestamp) *ServiceAccount_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ServiceAccount_FieldSubPathValue)
}

func (s ServiceAccountPathSelectorMetadataUpdateTime) WithArrayOfValues(values []*timestamp.Timestamp) *ServiceAccount_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ServiceAccount_FieldSubPathArrayOfValues)
}

type ServiceAccountPathSelectorMetadataUuid struct{}

func (ServiceAccountPathSelectorMetadataUuid) FieldPath() *ServiceAccount_FieldSubPath {
	return &ServiceAccount_FieldSubPath{
		selector: ServiceAccount_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Uuid().FieldPath(),
	}
}

func (s ServiceAccountPathSelectorMetadataUuid) WithValue(value string) *ServiceAccount_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ServiceAccount_FieldSubPathValue)
}

func (s ServiceAccountPathSelectorMetadataUuid) WithArrayOfValues(values []string) *ServiceAccount_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ServiceAccount_FieldSubPathArrayOfValues)
}

type ServiceAccountPathSelectorMetadataTags struct{}

func (ServiceAccountPathSelectorMetadataTags) FieldPath() *ServiceAccount_FieldSubPath {
	return &ServiceAccount_FieldSubPath{
		selector: ServiceAccount_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Tags().FieldPath(),
	}
}

func (s ServiceAccountPathSelectorMetadataTags) WithValue(value []string) *ServiceAccount_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ServiceAccount_FieldSubPathValue)
}

func (s ServiceAccountPathSelectorMetadataTags) WithArrayOfValues(values [][]string) *ServiceAccount_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ServiceAccount_FieldSubPathArrayOfValues)
}

func (s ServiceAccountPathSelectorMetadataTags) WithItemValue(value string) *ServiceAccount_FieldSubPathArrayItemValue {
	return s.FieldPath().WithIArrayItemValue(value).(*ServiceAccount_FieldSubPathArrayItemValue)
}

type ServiceAccountPathSelectorMetadataLabels struct{}

func (ServiceAccountPathSelectorMetadataLabels) FieldPath() *ServiceAccount_FieldSubPath {
	return &ServiceAccount_FieldSubPath{
		selector: ServiceAccount_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Labels().FieldPath(),
	}
}

func (s ServiceAccountPathSelectorMetadataLabels) WithValue(value map[string]string) *ServiceAccount_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ServiceAccount_FieldSubPathValue)
}

func (s ServiceAccountPathSelectorMetadataLabels) WithArrayOfValues(values []map[string]string) *ServiceAccount_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ServiceAccount_FieldSubPathArrayOfValues)
}

func (ServiceAccountPathSelectorMetadataLabels) WithKey(key string) ServiceAccountMapPathSelectorMetadataLabels {
	return ServiceAccountMapPathSelectorMetadataLabels{key: key}
}

type ServiceAccountMapPathSelectorMetadataLabels struct {
	key string
}

func (s ServiceAccountMapPathSelectorMetadataLabels) FieldPath() *ServiceAccount_FieldSubPath {
	return &ServiceAccount_FieldSubPath{
		selector: ServiceAccount_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Labels().WithKey(s.key).FieldPath(),
	}
}

func (s ServiceAccountMapPathSelectorMetadataLabels) WithValue(value string) *ServiceAccount_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ServiceAccount_FieldSubPathValue)
}

func (s ServiceAccountMapPathSelectorMetadataLabels) WithArrayOfValues(values []string) *ServiceAccount_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ServiceAccount_FieldSubPathArrayOfValues)
}

type ServiceAccountPathSelectorMetadataAnnotations struct{}

func (ServiceAccountPathSelectorMetadataAnnotations) FieldPath() *ServiceAccount_FieldSubPath {
	return &ServiceAccount_FieldSubPath{
		selector: ServiceAccount_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Annotations().FieldPath(),
	}
}

func (s ServiceAccountPathSelectorMetadataAnnotations) WithValue(value map[string]string) *ServiceAccount_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ServiceAccount_FieldSubPathValue)
}

func (s ServiceAccountPathSelectorMetadataAnnotations) WithArrayOfValues(values []map[string]string) *ServiceAccount_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ServiceAccount_FieldSubPathArrayOfValues)
}

func (ServiceAccountPathSelectorMetadataAnnotations) WithKey(key string) ServiceAccountMapPathSelectorMetadataAnnotations {
	return ServiceAccountMapPathSelectorMetadataAnnotations{key: key}
}

type ServiceAccountMapPathSelectorMetadataAnnotations struct {
	key string
}

func (s ServiceAccountMapPathSelectorMetadataAnnotations) FieldPath() *ServiceAccount_FieldSubPath {
	return &ServiceAccount_FieldSubPath{
		selector: ServiceAccount_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Annotations().WithKey(s.key).FieldPath(),
	}
}

func (s ServiceAccountMapPathSelectorMetadataAnnotations) WithValue(value string) *ServiceAccount_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ServiceAccount_FieldSubPathValue)
}

func (s ServiceAccountMapPathSelectorMetadataAnnotations) WithArrayOfValues(values []string) *ServiceAccount_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ServiceAccount_FieldSubPathArrayOfValues)
}

type ServiceAccountPathSelectorMetadataGeneration struct{}

func (ServiceAccountPathSelectorMetadataGeneration) FieldPath() *ServiceAccount_FieldSubPath {
	return &ServiceAccount_FieldSubPath{
		selector: ServiceAccount_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Generation().FieldPath(),
	}
}

func (s ServiceAccountPathSelectorMetadataGeneration) WithValue(value int64) *ServiceAccount_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ServiceAccount_FieldSubPathValue)
}

func (s ServiceAccountPathSelectorMetadataGeneration) WithArrayOfValues(values []int64) *ServiceAccount_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ServiceAccount_FieldSubPathArrayOfValues)
}

type ServiceAccountPathSelectorMetadataResourceVersion struct{}

func (ServiceAccountPathSelectorMetadataResourceVersion) FieldPath() *ServiceAccount_FieldSubPath {
	return &ServiceAccount_FieldSubPath{
		selector: ServiceAccount_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().ResourceVersion().FieldPath(),
	}
}

func (s ServiceAccountPathSelectorMetadataResourceVersion) WithValue(value string) *ServiceAccount_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ServiceAccount_FieldSubPathValue)
}

func (s ServiceAccountPathSelectorMetadataResourceVersion) WithArrayOfValues(values []string) *ServiceAccount_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ServiceAccount_FieldSubPathArrayOfValues)
}

type ServiceAccountPathSelectorMetadataOwnerReferences struct{}

func (ServiceAccountPathSelectorMetadataOwnerReferences) FieldPath() *ServiceAccount_FieldSubPath {
	return &ServiceAccount_FieldSubPath{
		selector: ServiceAccount_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().FieldPath(),
	}
}

func (s ServiceAccountPathSelectorMetadataOwnerReferences) WithValue(value []*ntt_meta.OwnerReference) *ServiceAccount_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ServiceAccount_FieldSubPathValue)
}

func (s ServiceAccountPathSelectorMetadataOwnerReferences) WithArrayOfValues(values [][]*ntt_meta.OwnerReference) *ServiceAccount_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ServiceAccount_FieldSubPathArrayOfValues)
}

func (s ServiceAccountPathSelectorMetadataOwnerReferences) WithItemValue(value *ntt_meta.OwnerReference) *ServiceAccount_FieldSubPathArrayItemValue {
	return s.FieldPath().WithIArrayItemValue(value).(*ServiceAccount_FieldSubPathArrayItemValue)
}

func (ServiceAccountPathSelectorMetadataOwnerReferences) ApiVersion() ServiceAccountPathSelectorMetadataOwnerReferencesApiVersion {
	return ServiceAccountPathSelectorMetadataOwnerReferencesApiVersion{}
}

func (ServiceAccountPathSelectorMetadataOwnerReferences) Kind() ServiceAccountPathSelectorMetadataOwnerReferencesKind {
	return ServiceAccountPathSelectorMetadataOwnerReferencesKind{}
}

func (ServiceAccountPathSelectorMetadataOwnerReferences) Name() ServiceAccountPathSelectorMetadataOwnerReferencesName {
	return ServiceAccountPathSelectorMetadataOwnerReferencesName{}
}

func (ServiceAccountPathSelectorMetadataOwnerReferences) Uid() ServiceAccountPathSelectorMetadataOwnerReferencesUid {
	return ServiceAccountPathSelectorMetadataOwnerReferencesUid{}
}

func (ServiceAccountPathSelectorMetadataOwnerReferences) Controller() ServiceAccountPathSelectorMetadataOwnerReferencesController {
	return ServiceAccountPathSelectorMetadataOwnerReferencesController{}
}

func (ServiceAccountPathSelectorMetadataOwnerReferences) BlockOwnerDeletion() ServiceAccountPathSelectorMetadataOwnerReferencesBlockOwnerDeletion {
	return ServiceAccountPathSelectorMetadataOwnerReferencesBlockOwnerDeletion{}
}

type ServiceAccountPathSelectorMetadataOwnerReferencesApiVersion struct{}

func (ServiceAccountPathSelectorMetadataOwnerReferencesApiVersion) FieldPath() *ServiceAccount_FieldSubPath {
	return &ServiceAccount_FieldSubPath{
		selector: ServiceAccount_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().ApiVersion().FieldPath(),
	}
}

func (s ServiceAccountPathSelectorMetadataOwnerReferencesApiVersion) WithValue(value string) *ServiceAccount_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ServiceAccount_FieldSubPathValue)
}

func (s ServiceAccountPathSelectorMetadataOwnerReferencesApiVersion) WithArrayOfValues(values []string) *ServiceAccount_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ServiceAccount_FieldSubPathArrayOfValues)
}

type ServiceAccountPathSelectorMetadataOwnerReferencesKind struct{}

func (ServiceAccountPathSelectorMetadataOwnerReferencesKind) FieldPath() *ServiceAccount_FieldSubPath {
	return &ServiceAccount_FieldSubPath{
		selector: ServiceAccount_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().Kind().FieldPath(),
	}
}

func (s ServiceAccountPathSelectorMetadataOwnerReferencesKind) WithValue(value string) *ServiceAccount_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ServiceAccount_FieldSubPathValue)
}

func (s ServiceAccountPathSelectorMetadataOwnerReferencesKind) WithArrayOfValues(values []string) *ServiceAccount_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ServiceAccount_FieldSubPathArrayOfValues)
}

type ServiceAccountPathSelectorMetadataOwnerReferencesName struct{}

func (ServiceAccountPathSelectorMetadataOwnerReferencesName) FieldPath() *ServiceAccount_FieldSubPath {
	return &ServiceAccount_FieldSubPath{
		selector: ServiceAccount_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().Name().FieldPath(),
	}
}

func (s ServiceAccountPathSelectorMetadataOwnerReferencesName) WithValue(value string) *ServiceAccount_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ServiceAccount_FieldSubPathValue)
}

func (s ServiceAccountPathSelectorMetadataOwnerReferencesName) WithArrayOfValues(values []string) *ServiceAccount_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ServiceAccount_FieldSubPathArrayOfValues)
}

type ServiceAccountPathSelectorMetadataOwnerReferencesUid struct{}

func (ServiceAccountPathSelectorMetadataOwnerReferencesUid) FieldPath() *ServiceAccount_FieldSubPath {
	return &ServiceAccount_FieldSubPath{
		selector: ServiceAccount_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().Uid().FieldPath(),
	}
}

func (s ServiceAccountPathSelectorMetadataOwnerReferencesUid) WithValue(value string) *ServiceAccount_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ServiceAccount_FieldSubPathValue)
}

func (s ServiceAccountPathSelectorMetadataOwnerReferencesUid) WithArrayOfValues(values []string) *ServiceAccount_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ServiceAccount_FieldSubPathArrayOfValues)
}

type ServiceAccountPathSelectorMetadataOwnerReferencesController struct{}

func (ServiceAccountPathSelectorMetadataOwnerReferencesController) FieldPath() *ServiceAccount_FieldSubPath {
	return &ServiceAccount_FieldSubPath{
		selector: ServiceAccount_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().Controller().FieldPath(),
	}
}

func (s ServiceAccountPathSelectorMetadataOwnerReferencesController) WithValue(value bool) *ServiceAccount_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ServiceAccount_FieldSubPathValue)
}

func (s ServiceAccountPathSelectorMetadataOwnerReferencesController) WithArrayOfValues(values []bool) *ServiceAccount_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ServiceAccount_FieldSubPathArrayOfValues)
}

type ServiceAccountPathSelectorMetadataOwnerReferencesBlockOwnerDeletion struct{}

func (ServiceAccountPathSelectorMetadataOwnerReferencesBlockOwnerDeletion) FieldPath() *ServiceAccount_FieldSubPath {
	return &ServiceAccount_FieldSubPath{
		selector: ServiceAccount_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().BlockOwnerDeletion().FieldPath(),
	}
}

func (s ServiceAccountPathSelectorMetadataOwnerReferencesBlockOwnerDeletion) WithValue(value bool) *ServiceAccount_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ServiceAccount_FieldSubPathValue)
}

func (s ServiceAccountPathSelectorMetadataOwnerReferencesBlockOwnerDeletion) WithArrayOfValues(values []bool) *ServiceAccount_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ServiceAccount_FieldSubPathArrayOfValues)
}

type ServiceAccountPathSelectorMetadataShards struct{}

func (ServiceAccountPathSelectorMetadataShards) FieldPath() *ServiceAccount_FieldSubPath {
	return &ServiceAccount_FieldSubPath{
		selector: ServiceAccount_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Shards().FieldPath(),
	}
}

func (s ServiceAccountPathSelectorMetadataShards) WithValue(value map[string]int64) *ServiceAccount_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ServiceAccount_FieldSubPathValue)
}

func (s ServiceAccountPathSelectorMetadataShards) WithArrayOfValues(values []map[string]int64) *ServiceAccount_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ServiceAccount_FieldSubPathArrayOfValues)
}

func (ServiceAccountPathSelectorMetadataShards) WithKey(key string) ServiceAccountMapPathSelectorMetadataShards {
	return ServiceAccountMapPathSelectorMetadataShards{key: key}
}

type ServiceAccountMapPathSelectorMetadataShards struct {
	key string
}

func (s ServiceAccountMapPathSelectorMetadataShards) FieldPath() *ServiceAccount_FieldSubPath {
	return &ServiceAccount_FieldSubPath{
		selector: ServiceAccount_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Shards().WithKey(s.key).FieldPath(),
	}
}

func (s ServiceAccountMapPathSelectorMetadataShards) WithValue(value int64) *ServiceAccount_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ServiceAccount_FieldSubPathValue)
}

func (s ServiceAccountMapPathSelectorMetadataShards) WithArrayOfValues(values []int64) *ServiceAccount_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ServiceAccount_FieldSubPathArrayOfValues)
}

type ServiceAccountPathSelectorMetadataSyncing struct{}

func (ServiceAccountPathSelectorMetadataSyncing) FieldPath() *ServiceAccount_FieldSubPath {
	return &ServiceAccount_FieldSubPath{
		selector: ServiceAccount_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Syncing().FieldPath(),
	}
}

func (s ServiceAccountPathSelectorMetadataSyncing) WithValue(value *ntt_meta.SyncingMeta) *ServiceAccount_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ServiceAccount_FieldSubPathValue)
}

func (s ServiceAccountPathSelectorMetadataSyncing) WithArrayOfValues(values []*ntt_meta.SyncingMeta) *ServiceAccount_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ServiceAccount_FieldSubPathArrayOfValues)
}

func (ServiceAccountPathSelectorMetadataSyncing) OwningRegion() ServiceAccountPathSelectorMetadataSyncingOwningRegion {
	return ServiceAccountPathSelectorMetadataSyncingOwningRegion{}
}

func (ServiceAccountPathSelectorMetadataSyncing) Regions() ServiceAccountPathSelectorMetadataSyncingRegions {
	return ServiceAccountPathSelectorMetadataSyncingRegions{}
}

type ServiceAccountPathSelectorMetadataSyncingOwningRegion struct{}

func (ServiceAccountPathSelectorMetadataSyncingOwningRegion) FieldPath() *ServiceAccount_FieldSubPath {
	return &ServiceAccount_FieldSubPath{
		selector: ServiceAccount_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Syncing().OwningRegion().FieldPath(),
	}
}

func (s ServiceAccountPathSelectorMetadataSyncingOwningRegion) WithValue(value string) *ServiceAccount_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ServiceAccount_FieldSubPathValue)
}

func (s ServiceAccountPathSelectorMetadataSyncingOwningRegion) WithArrayOfValues(values []string) *ServiceAccount_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ServiceAccount_FieldSubPathArrayOfValues)
}

type ServiceAccountPathSelectorMetadataSyncingRegions struct{}

func (ServiceAccountPathSelectorMetadataSyncingRegions) FieldPath() *ServiceAccount_FieldSubPath {
	return &ServiceAccount_FieldSubPath{
		selector: ServiceAccount_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Syncing().Regions().FieldPath(),
	}
}

func (s ServiceAccountPathSelectorMetadataSyncingRegions) WithValue(value []string) *ServiceAccount_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ServiceAccount_FieldSubPathValue)
}

func (s ServiceAccountPathSelectorMetadataSyncingRegions) WithArrayOfValues(values [][]string) *ServiceAccount_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ServiceAccount_FieldSubPathArrayOfValues)
}

func (s ServiceAccountPathSelectorMetadataSyncingRegions) WithItemValue(value string) *ServiceAccount_FieldSubPathArrayItemValue {
	return s.FieldPath().WithIArrayItemValue(value).(*ServiceAccount_FieldSubPathArrayItemValue)
}
