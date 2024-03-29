// Code generated by protoc-gen-goten-object
// File: edgelq/limits/proto/v1alpha2/limit.proto
// DO NOT EDIT!!!

package limit

// proto imports
import (
	iam_iam_common "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/common"
	iam_organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/organization"
	iam_project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/project"
	limit_pool "github.com/cloudwan/edgelq-sdk/limits/resources/v1alpha2/limit_pool"
	meta_resource "github.com/cloudwan/edgelq-sdk/meta/resources/v1alpha2/resource"
	meta_service "github.com/cloudwan/edgelq-sdk/meta/resources/v1alpha2/service"
	meta "github.com/cloudwan/goten-sdk/types/meta"
	multi_region_policy "github.com/cloudwan/goten-sdk/types/multi_region_policy"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

// make sure we're using proto imports
var (
	_ = &iam_iam_common.PCR{}
	_ = &iam_organization.Organization{}
	_ = &iam_project.Project{}
	_ = &limit_pool.LimitPool{}
	_ = &meta_resource.Resource{}
	_ = &meta_service.Service{}
	_ = &timestamppb.Timestamp{}
	_ = &meta.Meta{}
	_ = &multi_region_policy.MultiRegionPolicy{}
)

type LimitFieldPathBuilder struct{}

func NewLimitFieldPathBuilder() LimitFieldPathBuilder {
	return LimitFieldPathBuilder{}
}
func (LimitFieldPathBuilder) Name() LimitPathSelectorName {
	return LimitPathSelectorName{}
}
func (LimitFieldPathBuilder) Service() LimitPathSelectorService {
	return LimitPathSelectorService{}
}
func (LimitFieldPathBuilder) Resource() LimitPathSelectorResource {
	return LimitPathSelectorResource{}
}
func (LimitFieldPathBuilder) Region() LimitPathSelectorRegion {
	return LimitPathSelectorRegion{}
}
func (LimitFieldPathBuilder) ConfiguredLimit() LimitPathSelectorConfiguredLimit {
	return LimitPathSelectorConfiguredLimit{}
}
func (LimitFieldPathBuilder) ActiveLimit() LimitPathSelectorActiveLimit {
	return LimitPathSelectorActiveLimit{}
}
func (LimitFieldPathBuilder) Usage() LimitPathSelectorUsage {
	return LimitPathSelectorUsage{}
}
func (LimitFieldPathBuilder) Source() LimitPathSelectorSource {
	return LimitPathSelectorSource{}
}
func (LimitFieldPathBuilder) Metadata() LimitPathSelectorMetadata {
	return LimitPathSelectorMetadata{}
}

type LimitPathSelectorName struct{}

func (LimitPathSelectorName) FieldPath() *Limit_FieldTerminalPath {
	return &Limit_FieldTerminalPath{selector: Limit_FieldPathSelectorName}
}

func (s LimitPathSelectorName) WithValue(value *Name) *Limit_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*Limit_FieldTerminalPathValue)
}

func (s LimitPathSelectorName) WithArrayOfValues(values []*Name) *Limit_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Limit_FieldTerminalPathArrayOfValues)
}

type LimitPathSelectorService struct{}

func (LimitPathSelectorService) FieldPath() *Limit_FieldTerminalPath {
	return &Limit_FieldTerminalPath{selector: Limit_FieldPathSelectorService}
}

func (s LimitPathSelectorService) WithValue(value *meta_service.Reference) *Limit_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*Limit_FieldTerminalPathValue)
}

func (s LimitPathSelectorService) WithArrayOfValues(values []*meta_service.Reference) *Limit_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Limit_FieldTerminalPathArrayOfValues)
}

type LimitPathSelectorResource struct{}

func (LimitPathSelectorResource) FieldPath() *Limit_FieldTerminalPath {
	return &Limit_FieldTerminalPath{selector: Limit_FieldPathSelectorResource}
}

func (s LimitPathSelectorResource) WithValue(value *meta_resource.Reference) *Limit_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*Limit_FieldTerminalPathValue)
}

func (s LimitPathSelectorResource) WithArrayOfValues(values []*meta_resource.Reference) *Limit_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Limit_FieldTerminalPathArrayOfValues)
}

type LimitPathSelectorRegion struct{}

func (LimitPathSelectorRegion) FieldPath() *Limit_FieldTerminalPath {
	return &Limit_FieldTerminalPath{selector: Limit_FieldPathSelectorRegion}
}

func (s LimitPathSelectorRegion) WithValue(value string) *Limit_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*Limit_FieldTerminalPathValue)
}

func (s LimitPathSelectorRegion) WithArrayOfValues(values []string) *Limit_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Limit_FieldTerminalPathArrayOfValues)
}

type LimitPathSelectorConfiguredLimit struct{}

func (LimitPathSelectorConfiguredLimit) FieldPath() *Limit_FieldTerminalPath {
	return &Limit_FieldTerminalPath{selector: Limit_FieldPathSelectorConfiguredLimit}
}

func (s LimitPathSelectorConfiguredLimit) WithValue(value int64) *Limit_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*Limit_FieldTerminalPathValue)
}

func (s LimitPathSelectorConfiguredLimit) WithArrayOfValues(values []int64) *Limit_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Limit_FieldTerminalPathArrayOfValues)
}

type LimitPathSelectorActiveLimit struct{}

func (LimitPathSelectorActiveLimit) FieldPath() *Limit_FieldTerminalPath {
	return &Limit_FieldTerminalPath{selector: Limit_FieldPathSelectorActiveLimit}
}

func (s LimitPathSelectorActiveLimit) WithValue(value int64) *Limit_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*Limit_FieldTerminalPathValue)
}

func (s LimitPathSelectorActiveLimit) WithArrayOfValues(values []int64) *Limit_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Limit_FieldTerminalPathArrayOfValues)
}

type LimitPathSelectorUsage struct{}

func (LimitPathSelectorUsage) FieldPath() *Limit_FieldTerminalPath {
	return &Limit_FieldTerminalPath{selector: Limit_FieldPathSelectorUsage}
}

func (s LimitPathSelectorUsage) WithValue(value int64) *Limit_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*Limit_FieldTerminalPathValue)
}

func (s LimitPathSelectorUsage) WithArrayOfValues(values []int64) *Limit_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Limit_FieldTerminalPathArrayOfValues)
}

type LimitPathSelectorSource struct{}

func (LimitPathSelectorSource) FieldPath() *Limit_FieldTerminalPath {
	return &Limit_FieldTerminalPath{selector: Limit_FieldPathSelectorSource}
}

func (s LimitPathSelectorSource) WithValue(value *limit_pool.Reference) *Limit_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*Limit_FieldTerminalPathValue)
}

func (s LimitPathSelectorSource) WithArrayOfValues(values []*limit_pool.Reference) *Limit_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Limit_FieldTerminalPathArrayOfValues)
}

type LimitPathSelectorMetadata struct{}

func (LimitPathSelectorMetadata) FieldPath() *Limit_FieldTerminalPath {
	return &Limit_FieldTerminalPath{selector: Limit_FieldPathSelectorMetadata}
}

func (s LimitPathSelectorMetadata) WithValue(value *meta.Meta) *Limit_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*Limit_FieldTerminalPathValue)
}

func (s LimitPathSelectorMetadata) WithArrayOfValues(values []*meta.Meta) *Limit_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Limit_FieldTerminalPathArrayOfValues)
}

func (LimitPathSelectorMetadata) WithSubPath(subPath meta.Meta_FieldPath) *Limit_FieldSubPath {
	return &Limit_FieldSubPath{selector: Limit_FieldPathSelectorMetadata, subPath: subPath}
}

func (s LimitPathSelectorMetadata) WithSubValue(subPathValue meta.Meta_FieldPathValue) *Limit_FieldSubPathValue {
	return &Limit_FieldSubPathValue{Limit_FieldPath: s.WithSubPath(subPathValue), subPathValue: subPathValue}
}

func (s LimitPathSelectorMetadata) WithSubArrayOfValues(subPathArrayOfValues meta.Meta_FieldPathArrayOfValues) *Limit_FieldSubPathArrayOfValues {
	return &Limit_FieldSubPathArrayOfValues{Limit_FieldPath: s.WithSubPath(subPathArrayOfValues), subPathArrayOfValues: subPathArrayOfValues}
}

func (s LimitPathSelectorMetadata) WithSubArrayItemValue(subPathArrayItemValue meta.Meta_FieldPathArrayItemValue) *Limit_FieldSubPathArrayItemValue {
	return &Limit_FieldSubPathArrayItemValue{Limit_FieldPath: s.WithSubPath(subPathArrayItemValue), subPathItemValue: subPathArrayItemValue}
}

func (LimitPathSelectorMetadata) CreateTime() LimitPathSelectorMetadataCreateTime {
	return LimitPathSelectorMetadataCreateTime{}
}

func (LimitPathSelectorMetadata) UpdateTime() LimitPathSelectorMetadataUpdateTime {
	return LimitPathSelectorMetadataUpdateTime{}
}

func (LimitPathSelectorMetadata) DeleteTime() LimitPathSelectorMetadataDeleteTime {
	return LimitPathSelectorMetadataDeleteTime{}
}

func (LimitPathSelectorMetadata) Uuid() LimitPathSelectorMetadataUuid {
	return LimitPathSelectorMetadataUuid{}
}

func (LimitPathSelectorMetadata) Tags() LimitPathSelectorMetadataTags {
	return LimitPathSelectorMetadataTags{}
}

func (LimitPathSelectorMetadata) Labels() LimitPathSelectorMetadataLabels {
	return LimitPathSelectorMetadataLabels{}
}

func (LimitPathSelectorMetadata) Annotations() LimitPathSelectorMetadataAnnotations {
	return LimitPathSelectorMetadataAnnotations{}
}

func (LimitPathSelectorMetadata) Generation() LimitPathSelectorMetadataGeneration {
	return LimitPathSelectorMetadataGeneration{}
}

func (LimitPathSelectorMetadata) ResourceVersion() LimitPathSelectorMetadataResourceVersion {
	return LimitPathSelectorMetadataResourceVersion{}
}

func (LimitPathSelectorMetadata) OwnerReferences() LimitPathSelectorMetadataOwnerReferences {
	return LimitPathSelectorMetadataOwnerReferences{}
}

func (LimitPathSelectorMetadata) Shards() LimitPathSelectorMetadataShards {
	return LimitPathSelectorMetadataShards{}
}

func (LimitPathSelectorMetadata) Syncing() LimitPathSelectorMetadataSyncing {
	return LimitPathSelectorMetadataSyncing{}
}

func (LimitPathSelectorMetadata) Lifecycle() LimitPathSelectorMetadataLifecycle {
	return LimitPathSelectorMetadataLifecycle{}
}

func (LimitPathSelectorMetadata) Services() LimitPathSelectorMetadataServices {
	return LimitPathSelectorMetadataServices{}
}

type LimitPathSelectorMetadataCreateTime struct{}

func (LimitPathSelectorMetadataCreateTime) FieldPath() *Limit_FieldSubPath {
	return &Limit_FieldSubPath{
		selector: Limit_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().CreateTime().FieldPath(),
	}
}

func (s LimitPathSelectorMetadataCreateTime) WithValue(value *timestamppb.Timestamp) *Limit_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Limit_FieldSubPathValue)
}

func (s LimitPathSelectorMetadataCreateTime) WithArrayOfValues(values []*timestamppb.Timestamp) *Limit_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Limit_FieldSubPathArrayOfValues)
}

type LimitPathSelectorMetadataUpdateTime struct{}

func (LimitPathSelectorMetadataUpdateTime) FieldPath() *Limit_FieldSubPath {
	return &Limit_FieldSubPath{
		selector: Limit_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().UpdateTime().FieldPath(),
	}
}

func (s LimitPathSelectorMetadataUpdateTime) WithValue(value *timestamppb.Timestamp) *Limit_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Limit_FieldSubPathValue)
}

func (s LimitPathSelectorMetadataUpdateTime) WithArrayOfValues(values []*timestamppb.Timestamp) *Limit_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Limit_FieldSubPathArrayOfValues)
}

type LimitPathSelectorMetadataDeleteTime struct{}

func (LimitPathSelectorMetadataDeleteTime) FieldPath() *Limit_FieldSubPath {
	return &Limit_FieldSubPath{
		selector: Limit_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().DeleteTime().FieldPath(),
	}
}

func (s LimitPathSelectorMetadataDeleteTime) WithValue(value *timestamppb.Timestamp) *Limit_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Limit_FieldSubPathValue)
}

func (s LimitPathSelectorMetadataDeleteTime) WithArrayOfValues(values []*timestamppb.Timestamp) *Limit_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Limit_FieldSubPathArrayOfValues)
}

type LimitPathSelectorMetadataUuid struct{}

func (LimitPathSelectorMetadataUuid) FieldPath() *Limit_FieldSubPath {
	return &Limit_FieldSubPath{
		selector: Limit_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Uuid().FieldPath(),
	}
}

func (s LimitPathSelectorMetadataUuid) WithValue(value string) *Limit_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Limit_FieldSubPathValue)
}

func (s LimitPathSelectorMetadataUuid) WithArrayOfValues(values []string) *Limit_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Limit_FieldSubPathArrayOfValues)
}

type LimitPathSelectorMetadataTags struct{}

func (LimitPathSelectorMetadataTags) FieldPath() *Limit_FieldSubPath {
	return &Limit_FieldSubPath{
		selector: Limit_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Tags().FieldPath(),
	}
}

func (s LimitPathSelectorMetadataTags) WithValue(value []string) *Limit_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Limit_FieldSubPathValue)
}

func (s LimitPathSelectorMetadataTags) WithArrayOfValues(values [][]string) *Limit_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Limit_FieldSubPathArrayOfValues)
}

func (s LimitPathSelectorMetadataTags) WithItemValue(value string) *Limit_FieldSubPathArrayItemValue {
	return s.FieldPath().WithIArrayItemValue(value).(*Limit_FieldSubPathArrayItemValue)
}

type LimitPathSelectorMetadataLabels struct{}

func (LimitPathSelectorMetadataLabels) FieldPath() *Limit_FieldSubPath {
	return &Limit_FieldSubPath{
		selector: Limit_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Labels().FieldPath(),
	}
}

func (s LimitPathSelectorMetadataLabels) WithValue(value map[string]string) *Limit_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Limit_FieldSubPathValue)
}

func (s LimitPathSelectorMetadataLabels) WithArrayOfValues(values []map[string]string) *Limit_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Limit_FieldSubPathArrayOfValues)
}

func (LimitPathSelectorMetadataLabels) WithKey(key string) LimitMapPathSelectorMetadataLabels {
	return LimitMapPathSelectorMetadataLabels{key: key}
}

type LimitMapPathSelectorMetadataLabels struct {
	key string
}

func (s LimitMapPathSelectorMetadataLabels) FieldPath() *Limit_FieldSubPath {
	return &Limit_FieldSubPath{
		selector: Limit_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Labels().WithKey(s.key).FieldPath(),
	}
}

func (s LimitMapPathSelectorMetadataLabels) WithValue(value string) *Limit_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Limit_FieldSubPathValue)
}

func (s LimitMapPathSelectorMetadataLabels) WithArrayOfValues(values []string) *Limit_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Limit_FieldSubPathArrayOfValues)
}

type LimitPathSelectorMetadataAnnotations struct{}

func (LimitPathSelectorMetadataAnnotations) FieldPath() *Limit_FieldSubPath {
	return &Limit_FieldSubPath{
		selector: Limit_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Annotations().FieldPath(),
	}
}

func (s LimitPathSelectorMetadataAnnotations) WithValue(value map[string]string) *Limit_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Limit_FieldSubPathValue)
}

func (s LimitPathSelectorMetadataAnnotations) WithArrayOfValues(values []map[string]string) *Limit_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Limit_FieldSubPathArrayOfValues)
}

func (LimitPathSelectorMetadataAnnotations) WithKey(key string) LimitMapPathSelectorMetadataAnnotations {
	return LimitMapPathSelectorMetadataAnnotations{key: key}
}

type LimitMapPathSelectorMetadataAnnotations struct {
	key string
}

func (s LimitMapPathSelectorMetadataAnnotations) FieldPath() *Limit_FieldSubPath {
	return &Limit_FieldSubPath{
		selector: Limit_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Annotations().WithKey(s.key).FieldPath(),
	}
}

func (s LimitMapPathSelectorMetadataAnnotations) WithValue(value string) *Limit_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Limit_FieldSubPathValue)
}

func (s LimitMapPathSelectorMetadataAnnotations) WithArrayOfValues(values []string) *Limit_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Limit_FieldSubPathArrayOfValues)
}

type LimitPathSelectorMetadataGeneration struct{}

func (LimitPathSelectorMetadataGeneration) FieldPath() *Limit_FieldSubPath {
	return &Limit_FieldSubPath{
		selector: Limit_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Generation().FieldPath(),
	}
}

func (s LimitPathSelectorMetadataGeneration) WithValue(value int64) *Limit_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Limit_FieldSubPathValue)
}

func (s LimitPathSelectorMetadataGeneration) WithArrayOfValues(values []int64) *Limit_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Limit_FieldSubPathArrayOfValues)
}

type LimitPathSelectorMetadataResourceVersion struct{}

func (LimitPathSelectorMetadataResourceVersion) FieldPath() *Limit_FieldSubPath {
	return &Limit_FieldSubPath{
		selector: Limit_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().ResourceVersion().FieldPath(),
	}
}

func (s LimitPathSelectorMetadataResourceVersion) WithValue(value string) *Limit_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Limit_FieldSubPathValue)
}

func (s LimitPathSelectorMetadataResourceVersion) WithArrayOfValues(values []string) *Limit_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Limit_FieldSubPathArrayOfValues)
}

type LimitPathSelectorMetadataOwnerReferences struct{}

func (LimitPathSelectorMetadataOwnerReferences) FieldPath() *Limit_FieldSubPath {
	return &Limit_FieldSubPath{
		selector: Limit_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().OwnerReferences().FieldPath(),
	}
}

func (s LimitPathSelectorMetadataOwnerReferences) WithValue(value []*meta.OwnerReference) *Limit_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Limit_FieldSubPathValue)
}

func (s LimitPathSelectorMetadataOwnerReferences) WithArrayOfValues(values [][]*meta.OwnerReference) *Limit_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Limit_FieldSubPathArrayOfValues)
}

func (s LimitPathSelectorMetadataOwnerReferences) WithItemValue(value *meta.OwnerReference) *Limit_FieldSubPathArrayItemValue {
	return s.FieldPath().WithIArrayItemValue(value).(*Limit_FieldSubPathArrayItemValue)
}

func (LimitPathSelectorMetadataOwnerReferences) Kind() LimitPathSelectorMetadataOwnerReferencesKind {
	return LimitPathSelectorMetadataOwnerReferencesKind{}
}

func (LimitPathSelectorMetadataOwnerReferences) Version() LimitPathSelectorMetadataOwnerReferencesVersion {
	return LimitPathSelectorMetadataOwnerReferencesVersion{}
}

func (LimitPathSelectorMetadataOwnerReferences) Name() LimitPathSelectorMetadataOwnerReferencesName {
	return LimitPathSelectorMetadataOwnerReferencesName{}
}

func (LimitPathSelectorMetadataOwnerReferences) Region() LimitPathSelectorMetadataOwnerReferencesRegion {
	return LimitPathSelectorMetadataOwnerReferencesRegion{}
}

func (LimitPathSelectorMetadataOwnerReferences) Controller() LimitPathSelectorMetadataOwnerReferencesController {
	return LimitPathSelectorMetadataOwnerReferencesController{}
}

func (LimitPathSelectorMetadataOwnerReferences) RequiresOwnerReference() LimitPathSelectorMetadataOwnerReferencesRequiresOwnerReference {
	return LimitPathSelectorMetadataOwnerReferencesRequiresOwnerReference{}
}

type LimitPathSelectorMetadataOwnerReferencesKind struct{}

func (LimitPathSelectorMetadataOwnerReferencesKind) FieldPath() *Limit_FieldSubPath {
	return &Limit_FieldSubPath{
		selector: Limit_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().OwnerReferences().Kind().FieldPath(),
	}
}

func (s LimitPathSelectorMetadataOwnerReferencesKind) WithValue(value string) *Limit_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Limit_FieldSubPathValue)
}

func (s LimitPathSelectorMetadataOwnerReferencesKind) WithArrayOfValues(values []string) *Limit_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Limit_FieldSubPathArrayOfValues)
}

type LimitPathSelectorMetadataOwnerReferencesVersion struct{}

func (LimitPathSelectorMetadataOwnerReferencesVersion) FieldPath() *Limit_FieldSubPath {
	return &Limit_FieldSubPath{
		selector: Limit_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().OwnerReferences().Version().FieldPath(),
	}
}

func (s LimitPathSelectorMetadataOwnerReferencesVersion) WithValue(value string) *Limit_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Limit_FieldSubPathValue)
}

func (s LimitPathSelectorMetadataOwnerReferencesVersion) WithArrayOfValues(values []string) *Limit_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Limit_FieldSubPathArrayOfValues)
}

type LimitPathSelectorMetadataOwnerReferencesName struct{}

func (LimitPathSelectorMetadataOwnerReferencesName) FieldPath() *Limit_FieldSubPath {
	return &Limit_FieldSubPath{
		selector: Limit_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().OwnerReferences().Name().FieldPath(),
	}
}

func (s LimitPathSelectorMetadataOwnerReferencesName) WithValue(value string) *Limit_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Limit_FieldSubPathValue)
}

func (s LimitPathSelectorMetadataOwnerReferencesName) WithArrayOfValues(values []string) *Limit_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Limit_FieldSubPathArrayOfValues)
}

type LimitPathSelectorMetadataOwnerReferencesRegion struct{}

func (LimitPathSelectorMetadataOwnerReferencesRegion) FieldPath() *Limit_FieldSubPath {
	return &Limit_FieldSubPath{
		selector: Limit_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().OwnerReferences().Region().FieldPath(),
	}
}

func (s LimitPathSelectorMetadataOwnerReferencesRegion) WithValue(value string) *Limit_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Limit_FieldSubPathValue)
}

func (s LimitPathSelectorMetadataOwnerReferencesRegion) WithArrayOfValues(values []string) *Limit_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Limit_FieldSubPathArrayOfValues)
}

type LimitPathSelectorMetadataOwnerReferencesController struct{}

func (LimitPathSelectorMetadataOwnerReferencesController) FieldPath() *Limit_FieldSubPath {
	return &Limit_FieldSubPath{
		selector: Limit_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().OwnerReferences().Controller().FieldPath(),
	}
}

func (s LimitPathSelectorMetadataOwnerReferencesController) WithValue(value bool) *Limit_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Limit_FieldSubPathValue)
}

func (s LimitPathSelectorMetadataOwnerReferencesController) WithArrayOfValues(values []bool) *Limit_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Limit_FieldSubPathArrayOfValues)
}

type LimitPathSelectorMetadataOwnerReferencesRequiresOwnerReference struct{}

func (LimitPathSelectorMetadataOwnerReferencesRequiresOwnerReference) FieldPath() *Limit_FieldSubPath {
	return &Limit_FieldSubPath{
		selector: Limit_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().OwnerReferences().RequiresOwnerReference().FieldPath(),
	}
}

func (s LimitPathSelectorMetadataOwnerReferencesRequiresOwnerReference) WithValue(value bool) *Limit_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Limit_FieldSubPathValue)
}

func (s LimitPathSelectorMetadataOwnerReferencesRequiresOwnerReference) WithArrayOfValues(values []bool) *Limit_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Limit_FieldSubPathArrayOfValues)
}

type LimitPathSelectorMetadataShards struct{}

func (LimitPathSelectorMetadataShards) FieldPath() *Limit_FieldSubPath {
	return &Limit_FieldSubPath{
		selector: Limit_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Shards().FieldPath(),
	}
}

func (s LimitPathSelectorMetadataShards) WithValue(value map[string]int64) *Limit_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Limit_FieldSubPathValue)
}

func (s LimitPathSelectorMetadataShards) WithArrayOfValues(values []map[string]int64) *Limit_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Limit_FieldSubPathArrayOfValues)
}

func (LimitPathSelectorMetadataShards) WithKey(key string) LimitMapPathSelectorMetadataShards {
	return LimitMapPathSelectorMetadataShards{key: key}
}

type LimitMapPathSelectorMetadataShards struct {
	key string
}

func (s LimitMapPathSelectorMetadataShards) FieldPath() *Limit_FieldSubPath {
	return &Limit_FieldSubPath{
		selector: Limit_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Shards().WithKey(s.key).FieldPath(),
	}
}

func (s LimitMapPathSelectorMetadataShards) WithValue(value int64) *Limit_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Limit_FieldSubPathValue)
}

func (s LimitMapPathSelectorMetadataShards) WithArrayOfValues(values []int64) *Limit_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Limit_FieldSubPathArrayOfValues)
}

type LimitPathSelectorMetadataSyncing struct{}

func (LimitPathSelectorMetadataSyncing) FieldPath() *Limit_FieldSubPath {
	return &Limit_FieldSubPath{
		selector: Limit_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Syncing().FieldPath(),
	}
}

func (s LimitPathSelectorMetadataSyncing) WithValue(value *meta.SyncingMeta) *Limit_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Limit_FieldSubPathValue)
}

func (s LimitPathSelectorMetadataSyncing) WithArrayOfValues(values []*meta.SyncingMeta) *Limit_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Limit_FieldSubPathArrayOfValues)
}

func (LimitPathSelectorMetadataSyncing) OwningRegion() LimitPathSelectorMetadataSyncingOwningRegion {
	return LimitPathSelectorMetadataSyncingOwningRegion{}
}

func (LimitPathSelectorMetadataSyncing) Regions() LimitPathSelectorMetadataSyncingRegions {
	return LimitPathSelectorMetadataSyncingRegions{}
}

type LimitPathSelectorMetadataSyncingOwningRegion struct{}

func (LimitPathSelectorMetadataSyncingOwningRegion) FieldPath() *Limit_FieldSubPath {
	return &Limit_FieldSubPath{
		selector: Limit_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Syncing().OwningRegion().FieldPath(),
	}
}

func (s LimitPathSelectorMetadataSyncingOwningRegion) WithValue(value string) *Limit_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Limit_FieldSubPathValue)
}

func (s LimitPathSelectorMetadataSyncingOwningRegion) WithArrayOfValues(values []string) *Limit_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Limit_FieldSubPathArrayOfValues)
}

type LimitPathSelectorMetadataSyncingRegions struct{}

func (LimitPathSelectorMetadataSyncingRegions) FieldPath() *Limit_FieldSubPath {
	return &Limit_FieldSubPath{
		selector: Limit_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Syncing().Regions().FieldPath(),
	}
}

func (s LimitPathSelectorMetadataSyncingRegions) WithValue(value []string) *Limit_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Limit_FieldSubPathValue)
}

func (s LimitPathSelectorMetadataSyncingRegions) WithArrayOfValues(values [][]string) *Limit_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Limit_FieldSubPathArrayOfValues)
}

func (s LimitPathSelectorMetadataSyncingRegions) WithItemValue(value string) *Limit_FieldSubPathArrayItemValue {
	return s.FieldPath().WithIArrayItemValue(value).(*Limit_FieldSubPathArrayItemValue)
}

type LimitPathSelectorMetadataLifecycle struct{}

func (LimitPathSelectorMetadataLifecycle) FieldPath() *Limit_FieldSubPath {
	return &Limit_FieldSubPath{
		selector: Limit_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Lifecycle().FieldPath(),
	}
}

func (s LimitPathSelectorMetadataLifecycle) WithValue(value *meta.Lifecycle) *Limit_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Limit_FieldSubPathValue)
}

func (s LimitPathSelectorMetadataLifecycle) WithArrayOfValues(values []*meta.Lifecycle) *Limit_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Limit_FieldSubPathArrayOfValues)
}

func (LimitPathSelectorMetadataLifecycle) State() LimitPathSelectorMetadataLifecycleState {
	return LimitPathSelectorMetadataLifecycleState{}
}

func (LimitPathSelectorMetadataLifecycle) BlockDeletion() LimitPathSelectorMetadataLifecycleBlockDeletion {
	return LimitPathSelectorMetadataLifecycleBlockDeletion{}
}

type LimitPathSelectorMetadataLifecycleState struct{}

func (LimitPathSelectorMetadataLifecycleState) FieldPath() *Limit_FieldSubPath {
	return &Limit_FieldSubPath{
		selector: Limit_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Lifecycle().State().FieldPath(),
	}
}

func (s LimitPathSelectorMetadataLifecycleState) WithValue(value meta.Lifecycle_State) *Limit_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Limit_FieldSubPathValue)
}

func (s LimitPathSelectorMetadataLifecycleState) WithArrayOfValues(values []meta.Lifecycle_State) *Limit_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Limit_FieldSubPathArrayOfValues)
}

type LimitPathSelectorMetadataLifecycleBlockDeletion struct{}

func (LimitPathSelectorMetadataLifecycleBlockDeletion) FieldPath() *Limit_FieldSubPath {
	return &Limit_FieldSubPath{
		selector: Limit_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Lifecycle().BlockDeletion().FieldPath(),
	}
}

func (s LimitPathSelectorMetadataLifecycleBlockDeletion) WithValue(value bool) *Limit_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Limit_FieldSubPathValue)
}

func (s LimitPathSelectorMetadataLifecycleBlockDeletion) WithArrayOfValues(values []bool) *Limit_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Limit_FieldSubPathArrayOfValues)
}

type LimitPathSelectorMetadataServices struct{}

func (LimitPathSelectorMetadataServices) FieldPath() *Limit_FieldSubPath {
	return &Limit_FieldSubPath{
		selector: Limit_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Services().FieldPath(),
	}
}

func (s LimitPathSelectorMetadataServices) WithValue(value *meta.ServicesInfo) *Limit_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Limit_FieldSubPathValue)
}

func (s LimitPathSelectorMetadataServices) WithArrayOfValues(values []*meta.ServicesInfo) *Limit_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Limit_FieldSubPathArrayOfValues)
}

func (LimitPathSelectorMetadataServices) OwningService() LimitPathSelectorMetadataServicesOwningService {
	return LimitPathSelectorMetadataServicesOwningService{}
}

func (LimitPathSelectorMetadataServices) AllowedServices() LimitPathSelectorMetadataServicesAllowedServices {
	return LimitPathSelectorMetadataServicesAllowedServices{}
}

type LimitPathSelectorMetadataServicesOwningService struct{}

func (LimitPathSelectorMetadataServicesOwningService) FieldPath() *Limit_FieldSubPath {
	return &Limit_FieldSubPath{
		selector: Limit_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Services().OwningService().FieldPath(),
	}
}

func (s LimitPathSelectorMetadataServicesOwningService) WithValue(value string) *Limit_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Limit_FieldSubPathValue)
}

func (s LimitPathSelectorMetadataServicesOwningService) WithArrayOfValues(values []string) *Limit_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Limit_FieldSubPathArrayOfValues)
}

type LimitPathSelectorMetadataServicesAllowedServices struct{}

func (LimitPathSelectorMetadataServicesAllowedServices) FieldPath() *Limit_FieldSubPath {
	return &Limit_FieldSubPath{
		selector: Limit_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Services().AllowedServices().FieldPath(),
	}
}

func (s LimitPathSelectorMetadataServicesAllowedServices) WithValue(value []string) *Limit_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Limit_FieldSubPathValue)
}

func (s LimitPathSelectorMetadataServicesAllowedServices) WithArrayOfValues(values [][]string) *Limit_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Limit_FieldSubPathArrayOfValues)
}

func (s LimitPathSelectorMetadataServicesAllowedServices) WithItemValue(value string) *Limit_FieldSubPathArrayItemValue {
	return s.FieldPath().WithIArrayItemValue(value).(*Limit_FieldSubPathArrayItemValue)
}
