// Code generated by protoc-gen-goten-object
// File: edgelq/limits/proto/v1alpha2/limit_pool.proto
// DO NOT EDIT!!!

package limit_pool

// proto imports
import (
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
	multi_region_policy "github.com/cloudwan/edgelq-sdk/common/types/multi_region_policy"
	iam_organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/organization"
	meta_resource "github.com/cloudwan/edgelq-sdk/meta/resources/v1alpha2/resource"
	meta_service "github.com/cloudwan/edgelq-sdk/meta/resources/v1alpha2/service"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
)

// make sure we're using proto imports
var (
	_ = &ntt_meta.Meta{}
	_ = &multi_region_policy.MultiRegionPolicy{}
	_ = &iam_organization.Organization{}
	_ = &meta_resource.Resource{}
	_ = &meta_service.Service{}
	_ = &timestamp.Timestamp{}
)

type LimitPoolFieldPathBuilder struct{}

func NewLimitPoolFieldPathBuilder() LimitPoolFieldPathBuilder {
	return LimitPoolFieldPathBuilder{}
}
func (LimitPoolFieldPathBuilder) Name() LimitPoolPathSelectorName {
	return LimitPoolPathSelectorName{}
}
func (LimitPoolFieldPathBuilder) Service() LimitPoolPathSelectorService {
	return LimitPoolPathSelectorService{}
}
func (LimitPoolFieldPathBuilder) Resource() LimitPoolPathSelectorResource {
	return LimitPoolPathSelectorResource{}
}
func (LimitPoolFieldPathBuilder) ConfiguredSize() LimitPoolPathSelectorConfiguredSize {
	return LimitPoolPathSelectorConfiguredSize{}
}
func (LimitPoolFieldPathBuilder) ActiveSize() LimitPoolPathSelectorActiveSize {
	return LimitPoolPathSelectorActiveSize{}
}
func (LimitPoolFieldPathBuilder) Reserved() LimitPoolPathSelectorReserved {
	return LimitPoolPathSelectorReserved{}
}
func (LimitPoolFieldPathBuilder) Source() LimitPoolPathSelectorSource {
	return LimitPoolPathSelectorSource{}
}
func (LimitPoolFieldPathBuilder) Metadata() LimitPoolPathSelectorMetadata {
	return LimitPoolPathSelectorMetadata{}
}

type LimitPoolPathSelectorName struct{}

func (LimitPoolPathSelectorName) FieldPath() *LimitPool_FieldTerminalPath {
	return &LimitPool_FieldTerminalPath{selector: LimitPool_FieldPathSelectorName}
}

func (s LimitPoolPathSelectorName) WithValue(value *Name) *LimitPool_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*LimitPool_FieldTerminalPathValue)
}

func (s LimitPoolPathSelectorName) WithArrayOfValues(values []*Name) *LimitPool_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*LimitPool_FieldTerminalPathArrayOfValues)
}

type LimitPoolPathSelectorService struct{}

func (LimitPoolPathSelectorService) FieldPath() *LimitPool_FieldTerminalPath {
	return &LimitPool_FieldTerminalPath{selector: LimitPool_FieldPathSelectorService}
}

func (s LimitPoolPathSelectorService) WithValue(value *meta_service.Reference) *LimitPool_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*LimitPool_FieldTerminalPathValue)
}

func (s LimitPoolPathSelectorService) WithArrayOfValues(values []*meta_service.Reference) *LimitPool_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*LimitPool_FieldTerminalPathArrayOfValues)
}

type LimitPoolPathSelectorResource struct{}

func (LimitPoolPathSelectorResource) FieldPath() *LimitPool_FieldTerminalPath {
	return &LimitPool_FieldTerminalPath{selector: LimitPool_FieldPathSelectorResource}
}

func (s LimitPoolPathSelectorResource) WithValue(value *meta_resource.Reference) *LimitPool_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*LimitPool_FieldTerminalPathValue)
}

func (s LimitPoolPathSelectorResource) WithArrayOfValues(values []*meta_resource.Reference) *LimitPool_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*LimitPool_FieldTerminalPathArrayOfValues)
}

type LimitPoolPathSelectorConfiguredSize struct{}

func (LimitPoolPathSelectorConfiguredSize) FieldPath() *LimitPool_FieldTerminalPath {
	return &LimitPool_FieldTerminalPath{selector: LimitPool_FieldPathSelectorConfiguredSize}
}

func (s LimitPoolPathSelectorConfiguredSize) WithValue(value int64) *LimitPool_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*LimitPool_FieldTerminalPathValue)
}

func (s LimitPoolPathSelectorConfiguredSize) WithArrayOfValues(values []int64) *LimitPool_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*LimitPool_FieldTerminalPathArrayOfValues)
}

type LimitPoolPathSelectorActiveSize struct{}

func (LimitPoolPathSelectorActiveSize) FieldPath() *LimitPool_FieldTerminalPath {
	return &LimitPool_FieldTerminalPath{selector: LimitPool_FieldPathSelectorActiveSize}
}

func (s LimitPoolPathSelectorActiveSize) WithValue(value int64) *LimitPool_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*LimitPool_FieldTerminalPathValue)
}

func (s LimitPoolPathSelectorActiveSize) WithArrayOfValues(values []int64) *LimitPool_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*LimitPool_FieldTerminalPathArrayOfValues)
}

type LimitPoolPathSelectorReserved struct{}

func (LimitPoolPathSelectorReserved) FieldPath() *LimitPool_FieldTerminalPath {
	return &LimitPool_FieldTerminalPath{selector: LimitPool_FieldPathSelectorReserved}
}

func (s LimitPoolPathSelectorReserved) WithValue(value int64) *LimitPool_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*LimitPool_FieldTerminalPathValue)
}

func (s LimitPoolPathSelectorReserved) WithArrayOfValues(values []int64) *LimitPool_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*LimitPool_FieldTerminalPathArrayOfValues)
}

type LimitPoolPathSelectorSource struct{}

func (LimitPoolPathSelectorSource) FieldPath() *LimitPool_FieldTerminalPath {
	return &LimitPool_FieldTerminalPath{selector: LimitPool_FieldPathSelectorSource}
}

func (s LimitPoolPathSelectorSource) WithValue(value *Reference) *LimitPool_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*LimitPool_FieldTerminalPathValue)
}

func (s LimitPoolPathSelectorSource) WithArrayOfValues(values []*Reference) *LimitPool_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*LimitPool_FieldTerminalPathArrayOfValues)
}

type LimitPoolPathSelectorMetadata struct{}

func (LimitPoolPathSelectorMetadata) FieldPath() *LimitPool_FieldTerminalPath {
	return &LimitPool_FieldTerminalPath{selector: LimitPool_FieldPathSelectorMetadata}
}

func (s LimitPoolPathSelectorMetadata) WithValue(value *ntt_meta.Meta) *LimitPool_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*LimitPool_FieldTerminalPathValue)
}

func (s LimitPoolPathSelectorMetadata) WithArrayOfValues(values []*ntt_meta.Meta) *LimitPool_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*LimitPool_FieldTerminalPathArrayOfValues)
}

func (LimitPoolPathSelectorMetadata) WithSubPath(subPath ntt_meta.Meta_FieldPath) *LimitPool_FieldSubPath {
	return &LimitPool_FieldSubPath{selector: LimitPool_FieldPathSelectorMetadata, subPath: subPath}
}

func (s LimitPoolPathSelectorMetadata) WithSubValue(subPathValue ntt_meta.Meta_FieldPathValue) *LimitPool_FieldSubPathValue {
	return &LimitPool_FieldSubPathValue{LimitPool_FieldPath: s.WithSubPath(subPathValue), subPathValue: subPathValue}
}

func (s LimitPoolPathSelectorMetadata) WithSubArrayOfValues(subPathArrayOfValues ntt_meta.Meta_FieldPathArrayOfValues) *LimitPool_FieldSubPathArrayOfValues {
	return &LimitPool_FieldSubPathArrayOfValues{LimitPool_FieldPath: s.WithSubPath(subPathArrayOfValues), subPathArrayOfValues: subPathArrayOfValues}
}

func (s LimitPoolPathSelectorMetadata) WithSubArrayItemValue(subPathArrayItemValue ntt_meta.Meta_FieldPathArrayItemValue) *LimitPool_FieldSubPathArrayItemValue {
	return &LimitPool_FieldSubPathArrayItemValue{LimitPool_FieldPath: s.WithSubPath(subPathArrayItemValue), subPathItemValue: subPathArrayItemValue}
}

func (LimitPoolPathSelectorMetadata) CreateTime() LimitPoolPathSelectorMetadataCreateTime {
	return LimitPoolPathSelectorMetadataCreateTime{}
}

func (LimitPoolPathSelectorMetadata) UpdateTime() LimitPoolPathSelectorMetadataUpdateTime {
	return LimitPoolPathSelectorMetadataUpdateTime{}
}

func (LimitPoolPathSelectorMetadata) Uuid() LimitPoolPathSelectorMetadataUuid {
	return LimitPoolPathSelectorMetadataUuid{}
}

func (LimitPoolPathSelectorMetadata) Tags() LimitPoolPathSelectorMetadataTags {
	return LimitPoolPathSelectorMetadataTags{}
}

func (LimitPoolPathSelectorMetadata) Labels() LimitPoolPathSelectorMetadataLabels {
	return LimitPoolPathSelectorMetadataLabels{}
}

func (LimitPoolPathSelectorMetadata) Annotations() LimitPoolPathSelectorMetadataAnnotations {
	return LimitPoolPathSelectorMetadataAnnotations{}
}

func (LimitPoolPathSelectorMetadata) Generation() LimitPoolPathSelectorMetadataGeneration {
	return LimitPoolPathSelectorMetadataGeneration{}
}

func (LimitPoolPathSelectorMetadata) ResourceVersion() LimitPoolPathSelectorMetadataResourceVersion {
	return LimitPoolPathSelectorMetadataResourceVersion{}
}

func (LimitPoolPathSelectorMetadata) OwnerReferences() LimitPoolPathSelectorMetadataOwnerReferences {
	return LimitPoolPathSelectorMetadataOwnerReferences{}
}

func (LimitPoolPathSelectorMetadata) Shards() LimitPoolPathSelectorMetadataShards {
	return LimitPoolPathSelectorMetadataShards{}
}

func (LimitPoolPathSelectorMetadata) Syncing() LimitPoolPathSelectorMetadataSyncing {
	return LimitPoolPathSelectorMetadataSyncing{}
}

type LimitPoolPathSelectorMetadataCreateTime struct{}

func (LimitPoolPathSelectorMetadataCreateTime) FieldPath() *LimitPool_FieldSubPath {
	return &LimitPool_FieldSubPath{
		selector: LimitPool_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().CreateTime().FieldPath(),
	}
}

func (s LimitPoolPathSelectorMetadataCreateTime) WithValue(value *timestamp.Timestamp) *LimitPool_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*LimitPool_FieldSubPathValue)
}

func (s LimitPoolPathSelectorMetadataCreateTime) WithArrayOfValues(values []*timestamp.Timestamp) *LimitPool_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*LimitPool_FieldSubPathArrayOfValues)
}

type LimitPoolPathSelectorMetadataUpdateTime struct{}

func (LimitPoolPathSelectorMetadataUpdateTime) FieldPath() *LimitPool_FieldSubPath {
	return &LimitPool_FieldSubPath{
		selector: LimitPool_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().UpdateTime().FieldPath(),
	}
}

func (s LimitPoolPathSelectorMetadataUpdateTime) WithValue(value *timestamp.Timestamp) *LimitPool_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*LimitPool_FieldSubPathValue)
}

func (s LimitPoolPathSelectorMetadataUpdateTime) WithArrayOfValues(values []*timestamp.Timestamp) *LimitPool_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*LimitPool_FieldSubPathArrayOfValues)
}

type LimitPoolPathSelectorMetadataUuid struct{}

func (LimitPoolPathSelectorMetadataUuid) FieldPath() *LimitPool_FieldSubPath {
	return &LimitPool_FieldSubPath{
		selector: LimitPool_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Uuid().FieldPath(),
	}
}

func (s LimitPoolPathSelectorMetadataUuid) WithValue(value string) *LimitPool_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*LimitPool_FieldSubPathValue)
}

func (s LimitPoolPathSelectorMetadataUuid) WithArrayOfValues(values []string) *LimitPool_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*LimitPool_FieldSubPathArrayOfValues)
}

type LimitPoolPathSelectorMetadataTags struct{}

func (LimitPoolPathSelectorMetadataTags) FieldPath() *LimitPool_FieldSubPath {
	return &LimitPool_FieldSubPath{
		selector: LimitPool_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Tags().FieldPath(),
	}
}

func (s LimitPoolPathSelectorMetadataTags) WithValue(value []string) *LimitPool_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*LimitPool_FieldSubPathValue)
}

func (s LimitPoolPathSelectorMetadataTags) WithArrayOfValues(values [][]string) *LimitPool_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*LimitPool_FieldSubPathArrayOfValues)
}

func (s LimitPoolPathSelectorMetadataTags) WithItemValue(value string) *LimitPool_FieldSubPathArrayItemValue {
	return s.FieldPath().WithIArrayItemValue(value).(*LimitPool_FieldSubPathArrayItemValue)
}

type LimitPoolPathSelectorMetadataLabels struct{}

func (LimitPoolPathSelectorMetadataLabels) FieldPath() *LimitPool_FieldSubPath {
	return &LimitPool_FieldSubPath{
		selector: LimitPool_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Labels().FieldPath(),
	}
}

func (s LimitPoolPathSelectorMetadataLabels) WithValue(value map[string]string) *LimitPool_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*LimitPool_FieldSubPathValue)
}

func (s LimitPoolPathSelectorMetadataLabels) WithArrayOfValues(values []map[string]string) *LimitPool_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*LimitPool_FieldSubPathArrayOfValues)
}

func (LimitPoolPathSelectorMetadataLabels) WithKey(key string) LimitPoolMapPathSelectorMetadataLabels {
	return LimitPoolMapPathSelectorMetadataLabels{key: key}
}

type LimitPoolMapPathSelectorMetadataLabels struct {
	key string
}

func (s LimitPoolMapPathSelectorMetadataLabels) FieldPath() *LimitPool_FieldSubPath {
	return &LimitPool_FieldSubPath{
		selector: LimitPool_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Labels().WithKey(s.key).FieldPath(),
	}
}

func (s LimitPoolMapPathSelectorMetadataLabels) WithValue(value string) *LimitPool_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*LimitPool_FieldSubPathValue)
}

func (s LimitPoolMapPathSelectorMetadataLabels) WithArrayOfValues(values []string) *LimitPool_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*LimitPool_FieldSubPathArrayOfValues)
}

type LimitPoolPathSelectorMetadataAnnotations struct{}

func (LimitPoolPathSelectorMetadataAnnotations) FieldPath() *LimitPool_FieldSubPath {
	return &LimitPool_FieldSubPath{
		selector: LimitPool_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Annotations().FieldPath(),
	}
}

func (s LimitPoolPathSelectorMetadataAnnotations) WithValue(value map[string]string) *LimitPool_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*LimitPool_FieldSubPathValue)
}

func (s LimitPoolPathSelectorMetadataAnnotations) WithArrayOfValues(values []map[string]string) *LimitPool_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*LimitPool_FieldSubPathArrayOfValues)
}

func (LimitPoolPathSelectorMetadataAnnotations) WithKey(key string) LimitPoolMapPathSelectorMetadataAnnotations {
	return LimitPoolMapPathSelectorMetadataAnnotations{key: key}
}

type LimitPoolMapPathSelectorMetadataAnnotations struct {
	key string
}

func (s LimitPoolMapPathSelectorMetadataAnnotations) FieldPath() *LimitPool_FieldSubPath {
	return &LimitPool_FieldSubPath{
		selector: LimitPool_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Annotations().WithKey(s.key).FieldPath(),
	}
}

func (s LimitPoolMapPathSelectorMetadataAnnotations) WithValue(value string) *LimitPool_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*LimitPool_FieldSubPathValue)
}

func (s LimitPoolMapPathSelectorMetadataAnnotations) WithArrayOfValues(values []string) *LimitPool_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*LimitPool_FieldSubPathArrayOfValues)
}

type LimitPoolPathSelectorMetadataGeneration struct{}

func (LimitPoolPathSelectorMetadataGeneration) FieldPath() *LimitPool_FieldSubPath {
	return &LimitPool_FieldSubPath{
		selector: LimitPool_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Generation().FieldPath(),
	}
}

func (s LimitPoolPathSelectorMetadataGeneration) WithValue(value int64) *LimitPool_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*LimitPool_FieldSubPathValue)
}

func (s LimitPoolPathSelectorMetadataGeneration) WithArrayOfValues(values []int64) *LimitPool_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*LimitPool_FieldSubPathArrayOfValues)
}

type LimitPoolPathSelectorMetadataResourceVersion struct{}

func (LimitPoolPathSelectorMetadataResourceVersion) FieldPath() *LimitPool_FieldSubPath {
	return &LimitPool_FieldSubPath{
		selector: LimitPool_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().ResourceVersion().FieldPath(),
	}
}

func (s LimitPoolPathSelectorMetadataResourceVersion) WithValue(value string) *LimitPool_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*LimitPool_FieldSubPathValue)
}

func (s LimitPoolPathSelectorMetadataResourceVersion) WithArrayOfValues(values []string) *LimitPool_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*LimitPool_FieldSubPathArrayOfValues)
}

type LimitPoolPathSelectorMetadataOwnerReferences struct{}

func (LimitPoolPathSelectorMetadataOwnerReferences) FieldPath() *LimitPool_FieldSubPath {
	return &LimitPool_FieldSubPath{
		selector: LimitPool_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().FieldPath(),
	}
}

func (s LimitPoolPathSelectorMetadataOwnerReferences) WithValue(value []*ntt_meta.OwnerReference) *LimitPool_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*LimitPool_FieldSubPathValue)
}

func (s LimitPoolPathSelectorMetadataOwnerReferences) WithArrayOfValues(values [][]*ntt_meta.OwnerReference) *LimitPool_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*LimitPool_FieldSubPathArrayOfValues)
}

func (s LimitPoolPathSelectorMetadataOwnerReferences) WithItemValue(value *ntt_meta.OwnerReference) *LimitPool_FieldSubPathArrayItemValue {
	return s.FieldPath().WithIArrayItemValue(value).(*LimitPool_FieldSubPathArrayItemValue)
}

func (LimitPoolPathSelectorMetadataOwnerReferences) ApiVersion() LimitPoolPathSelectorMetadataOwnerReferencesApiVersion {
	return LimitPoolPathSelectorMetadataOwnerReferencesApiVersion{}
}

func (LimitPoolPathSelectorMetadataOwnerReferences) Kind() LimitPoolPathSelectorMetadataOwnerReferencesKind {
	return LimitPoolPathSelectorMetadataOwnerReferencesKind{}
}

func (LimitPoolPathSelectorMetadataOwnerReferences) Name() LimitPoolPathSelectorMetadataOwnerReferencesName {
	return LimitPoolPathSelectorMetadataOwnerReferencesName{}
}

func (LimitPoolPathSelectorMetadataOwnerReferences) Uid() LimitPoolPathSelectorMetadataOwnerReferencesUid {
	return LimitPoolPathSelectorMetadataOwnerReferencesUid{}
}

func (LimitPoolPathSelectorMetadataOwnerReferences) Controller() LimitPoolPathSelectorMetadataOwnerReferencesController {
	return LimitPoolPathSelectorMetadataOwnerReferencesController{}
}

func (LimitPoolPathSelectorMetadataOwnerReferences) BlockOwnerDeletion() LimitPoolPathSelectorMetadataOwnerReferencesBlockOwnerDeletion {
	return LimitPoolPathSelectorMetadataOwnerReferencesBlockOwnerDeletion{}
}

type LimitPoolPathSelectorMetadataOwnerReferencesApiVersion struct{}

func (LimitPoolPathSelectorMetadataOwnerReferencesApiVersion) FieldPath() *LimitPool_FieldSubPath {
	return &LimitPool_FieldSubPath{
		selector: LimitPool_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().ApiVersion().FieldPath(),
	}
}

func (s LimitPoolPathSelectorMetadataOwnerReferencesApiVersion) WithValue(value string) *LimitPool_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*LimitPool_FieldSubPathValue)
}

func (s LimitPoolPathSelectorMetadataOwnerReferencesApiVersion) WithArrayOfValues(values []string) *LimitPool_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*LimitPool_FieldSubPathArrayOfValues)
}

type LimitPoolPathSelectorMetadataOwnerReferencesKind struct{}

func (LimitPoolPathSelectorMetadataOwnerReferencesKind) FieldPath() *LimitPool_FieldSubPath {
	return &LimitPool_FieldSubPath{
		selector: LimitPool_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().Kind().FieldPath(),
	}
}

func (s LimitPoolPathSelectorMetadataOwnerReferencesKind) WithValue(value string) *LimitPool_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*LimitPool_FieldSubPathValue)
}

func (s LimitPoolPathSelectorMetadataOwnerReferencesKind) WithArrayOfValues(values []string) *LimitPool_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*LimitPool_FieldSubPathArrayOfValues)
}

type LimitPoolPathSelectorMetadataOwnerReferencesName struct{}

func (LimitPoolPathSelectorMetadataOwnerReferencesName) FieldPath() *LimitPool_FieldSubPath {
	return &LimitPool_FieldSubPath{
		selector: LimitPool_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().Name().FieldPath(),
	}
}

func (s LimitPoolPathSelectorMetadataOwnerReferencesName) WithValue(value string) *LimitPool_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*LimitPool_FieldSubPathValue)
}

func (s LimitPoolPathSelectorMetadataOwnerReferencesName) WithArrayOfValues(values []string) *LimitPool_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*LimitPool_FieldSubPathArrayOfValues)
}

type LimitPoolPathSelectorMetadataOwnerReferencesUid struct{}

func (LimitPoolPathSelectorMetadataOwnerReferencesUid) FieldPath() *LimitPool_FieldSubPath {
	return &LimitPool_FieldSubPath{
		selector: LimitPool_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().Uid().FieldPath(),
	}
}

func (s LimitPoolPathSelectorMetadataOwnerReferencesUid) WithValue(value string) *LimitPool_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*LimitPool_FieldSubPathValue)
}

func (s LimitPoolPathSelectorMetadataOwnerReferencesUid) WithArrayOfValues(values []string) *LimitPool_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*LimitPool_FieldSubPathArrayOfValues)
}

type LimitPoolPathSelectorMetadataOwnerReferencesController struct{}

func (LimitPoolPathSelectorMetadataOwnerReferencesController) FieldPath() *LimitPool_FieldSubPath {
	return &LimitPool_FieldSubPath{
		selector: LimitPool_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().Controller().FieldPath(),
	}
}

func (s LimitPoolPathSelectorMetadataOwnerReferencesController) WithValue(value bool) *LimitPool_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*LimitPool_FieldSubPathValue)
}

func (s LimitPoolPathSelectorMetadataOwnerReferencesController) WithArrayOfValues(values []bool) *LimitPool_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*LimitPool_FieldSubPathArrayOfValues)
}

type LimitPoolPathSelectorMetadataOwnerReferencesBlockOwnerDeletion struct{}

func (LimitPoolPathSelectorMetadataOwnerReferencesBlockOwnerDeletion) FieldPath() *LimitPool_FieldSubPath {
	return &LimitPool_FieldSubPath{
		selector: LimitPool_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().BlockOwnerDeletion().FieldPath(),
	}
}

func (s LimitPoolPathSelectorMetadataOwnerReferencesBlockOwnerDeletion) WithValue(value bool) *LimitPool_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*LimitPool_FieldSubPathValue)
}

func (s LimitPoolPathSelectorMetadataOwnerReferencesBlockOwnerDeletion) WithArrayOfValues(values []bool) *LimitPool_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*LimitPool_FieldSubPathArrayOfValues)
}

type LimitPoolPathSelectorMetadataShards struct{}

func (LimitPoolPathSelectorMetadataShards) FieldPath() *LimitPool_FieldSubPath {
	return &LimitPool_FieldSubPath{
		selector: LimitPool_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Shards().FieldPath(),
	}
}

func (s LimitPoolPathSelectorMetadataShards) WithValue(value map[string]int64) *LimitPool_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*LimitPool_FieldSubPathValue)
}

func (s LimitPoolPathSelectorMetadataShards) WithArrayOfValues(values []map[string]int64) *LimitPool_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*LimitPool_FieldSubPathArrayOfValues)
}

func (LimitPoolPathSelectorMetadataShards) WithKey(key string) LimitPoolMapPathSelectorMetadataShards {
	return LimitPoolMapPathSelectorMetadataShards{key: key}
}

type LimitPoolMapPathSelectorMetadataShards struct {
	key string
}

func (s LimitPoolMapPathSelectorMetadataShards) FieldPath() *LimitPool_FieldSubPath {
	return &LimitPool_FieldSubPath{
		selector: LimitPool_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Shards().WithKey(s.key).FieldPath(),
	}
}

func (s LimitPoolMapPathSelectorMetadataShards) WithValue(value int64) *LimitPool_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*LimitPool_FieldSubPathValue)
}

func (s LimitPoolMapPathSelectorMetadataShards) WithArrayOfValues(values []int64) *LimitPool_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*LimitPool_FieldSubPathArrayOfValues)
}

type LimitPoolPathSelectorMetadataSyncing struct{}

func (LimitPoolPathSelectorMetadataSyncing) FieldPath() *LimitPool_FieldSubPath {
	return &LimitPool_FieldSubPath{
		selector: LimitPool_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Syncing().FieldPath(),
	}
}

func (s LimitPoolPathSelectorMetadataSyncing) WithValue(value *ntt_meta.SyncingMeta) *LimitPool_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*LimitPool_FieldSubPathValue)
}

func (s LimitPoolPathSelectorMetadataSyncing) WithArrayOfValues(values []*ntt_meta.SyncingMeta) *LimitPool_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*LimitPool_FieldSubPathArrayOfValues)
}

func (LimitPoolPathSelectorMetadataSyncing) OwningRegion() LimitPoolPathSelectorMetadataSyncingOwningRegion {
	return LimitPoolPathSelectorMetadataSyncingOwningRegion{}
}

func (LimitPoolPathSelectorMetadataSyncing) Regions() LimitPoolPathSelectorMetadataSyncingRegions {
	return LimitPoolPathSelectorMetadataSyncingRegions{}
}

type LimitPoolPathSelectorMetadataSyncingOwningRegion struct{}

func (LimitPoolPathSelectorMetadataSyncingOwningRegion) FieldPath() *LimitPool_FieldSubPath {
	return &LimitPool_FieldSubPath{
		selector: LimitPool_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Syncing().OwningRegion().FieldPath(),
	}
}

func (s LimitPoolPathSelectorMetadataSyncingOwningRegion) WithValue(value string) *LimitPool_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*LimitPool_FieldSubPathValue)
}

func (s LimitPoolPathSelectorMetadataSyncingOwningRegion) WithArrayOfValues(values []string) *LimitPool_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*LimitPool_FieldSubPathArrayOfValues)
}

type LimitPoolPathSelectorMetadataSyncingRegions struct{}

func (LimitPoolPathSelectorMetadataSyncingRegions) FieldPath() *LimitPool_FieldSubPath {
	return &LimitPool_FieldSubPath{
		selector: LimitPool_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Syncing().Regions().FieldPath(),
	}
}

func (s LimitPoolPathSelectorMetadataSyncingRegions) WithValue(value []string) *LimitPool_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*LimitPool_FieldSubPathValue)
}

func (s LimitPoolPathSelectorMetadataSyncingRegions) WithArrayOfValues(values [][]string) *LimitPool_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*LimitPool_FieldSubPathArrayOfValues)
}

func (s LimitPoolPathSelectorMetadataSyncingRegions) WithItemValue(value string) *LimitPool_FieldSubPathArrayItemValue {
	return s.FieldPath().WithIArrayItemValue(value).(*LimitPool_FieldSubPathArrayItemValue)
}
