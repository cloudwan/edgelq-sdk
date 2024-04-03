// Code generated by protoc-gen-goten-object
// File: edgelq/ztp/proto/v1/tpm_attestation_cert.proto
// DO NOT EDIT!!!

package tpm_attestation_cert

// proto imports
import (
	project "github.com/cloudwan/edgelq-sdk/ztp/resources/v1/project"
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

type TpmAttestationCertFieldPathBuilder struct{}

func NewTpmAttestationCertFieldPathBuilder() TpmAttestationCertFieldPathBuilder {
	return TpmAttestationCertFieldPathBuilder{}
}
func (TpmAttestationCertFieldPathBuilder) Name() TpmAttestationCertPathSelectorName {
	return TpmAttestationCertPathSelectorName{}
}
func (TpmAttestationCertFieldPathBuilder) Metadata() TpmAttestationCertPathSelectorMetadata {
	return TpmAttestationCertPathSelectorMetadata{}
}
func (TpmAttestationCertFieldPathBuilder) DisplayName() TpmAttestationCertPathSelectorDisplayName {
	return TpmAttestationCertPathSelectorDisplayName{}
}
func (TpmAttestationCertFieldPathBuilder) Manufacturer() TpmAttestationCertPathSelectorManufacturer {
	return TpmAttestationCertPathSelectorManufacturer{}
}
func (TpmAttestationCertFieldPathBuilder) ProductName() TpmAttestationCertPathSelectorProductName {
	return TpmAttestationCertPathSelectorProductName{}
}
func (TpmAttestationCertFieldPathBuilder) TpmManufacturerCaCert() TpmAttestationCertPathSelectorTpmManufacturerCaCert {
	return TpmAttestationCertPathSelectorTpmManufacturerCaCert{}
}
func (TpmAttestationCertFieldPathBuilder) IdevidIssuerCaCert() TpmAttestationCertPathSelectorIdevidIssuerCaCert {
	return TpmAttestationCertPathSelectorIdevidIssuerCaCert{}
}
func (TpmAttestationCertFieldPathBuilder) LdevidIssuerCaCert() TpmAttestationCertPathSelectorLdevidIssuerCaCert {
	return TpmAttestationCertPathSelectorLdevidIssuerCaCert{}
}

type TpmAttestationCertPathSelectorName struct{}

func (TpmAttestationCertPathSelectorName) FieldPath() *TpmAttestationCert_FieldTerminalPath {
	return &TpmAttestationCert_FieldTerminalPath{selector: TpmAttestationCert_FieldPathSelectorName}
}

func (s TpmAttestationCertPathSelectorName) WithValue(value *Name) *TpmAttestationCert_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*TpmAttestationCert_FieldTerminalPathValue)
}

func (s TpmAttestationCertPathSelectorName) WithArrayOfValues(values []*Name) *TpmAttestationCert_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*TpmAttestationCert_FieldTerminalPathArrayOfValues)
}

type TpmAttestationCertPathSelectorMetadata struct{}

func (TpmAttestationCertPathSelectorMetadata) FieldPath() *TpmAttestationCert_FieldTerminalPath {
	return &TpmAttestationCert_FieldTerminalPath{selector: TpmAttestationCert_FieldPathSelectorMetadata}
}

func (s TpmAttestationCertPathSelectorMetadata) WithValue(value *meta.Meta) *TpmAttestationCert_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*TpmAttestationCert_FieldTerminalPathValue)
}

func (s TpmAttestationCertPathSelectorMetadata) WithArrayOfValues(values []*meta.Meta) *TpmAttestationCert_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*TpmAttestationCert_FieldTerminalPathArrayOfValues)
}

func (TpmAttestationCertPathSelectorMetadata) WithSubPath(subPath meta.Meta_FieldPath) *TpmAttestationCert_FieldSubPath {
	return &TpmAttestationCert_FieldSubPath{selector: TpmAttestationCert_FieldPathSelectorMetadata, subPath: subPath}
}

func (s TpmAttestationCertPathSelectorMetadata) WithSubValue(subPathValue meta.Meta_FieldPathValue) *TpmAttestationCert_FieldSubPathValue {
	return &TpmAttestationCert_FieldSubPathValue{TpmAttestationCert_FieldPath: s.WithSubPath(subPathValue), subPathValue: subPathValue}
}

func (s TpmAttestationCertPathSelectorMetadata) WithSubArrayOfValues(subPathArrayOfValues meta.Meta_FieldPathArrayOfValues) *TpmAttestationCert_FieldSubPathArrayOfValues {
	return &TpmAttestationCert_FieldSubPathArrayOfValues{TpmAttestationCert_FieldPath: s.WithSubPath(subPathArrayOfValues), subPathArrayOfValues: subPathArrayOfValues}
}

func (s TpmAttestationCertPathSelectorMetadata) WithSubArrayItemValue(subPathArrayItemValue meta.Meta_FieldPathArrayItemValue) *TpmAttestationCert_FieldSubPathArrayItemValue {
	return &TpmAttestationCert_FieldSubPathArrayItemValue{TpmAttestationCert_FieldPath: s.WithSubPath(subPathArrayItemValue), subPathItemValue: subPathArrayItemValue}
}

func (TpmAttestationCertPathSelectorMetadata) CreateTime() TpmAttestationCertPathSelectorMetadataCreateTime {
	return TpmAttestationCertPathSelectorMetadataCreateTime{}
}

func (TpmAttestationCertPathSelectorMetadata) UpdateTime() TpmAttestationCertPathSelectorMetadataUpdateTime {
	return TpmAttestationCertPathSelectorMetadataUpdateTime{}
}

func (TpmAttestationCertPathSelectorMetadata) DeleteTime() TpmAttestationCertPathSelectorMetadataDeleteTime {
	return TpmAttestationCertPathSelectorMetadataDeleteTime{}
}

func (TpmAttestationCertPathSelectorMetadata) Uuid() TpmAttestationCertPathSelectorMetadataUuid {
	return TpmAttestationCertPathSelectorMetadataUuid{}
}

func (TpmAttestationCertPathSelectorMetadata) Tags() TpmAttestationCertPathSelectorMetadataTags {
	return TpmAttestationCertPathSelectorMetadataTags{}
}

func (TpmAttestationCertPathSelectorMetadata) Labels() TpmAttestationCertPathSelectorMetadataLabels {
	return TpmAttestationCertPathSelectorMetadataLabels{}
}

func (TpmAttestationCertPathSelectorMetadata) Annotations() TpmAttestationCertPathSelectorMetadataAnnotations {
	return TpmAttestationCertPathSelectorMetadataAnnotations{}
}

func (TpmAttestationCertPathSelectorMetadata) Generation() TpmAttestationCertPathSelectorMetadataGeneration {
	return TpmAttestationCertPathSelectorMetadataGeneration{}
}

func (TpmAttestationCertPathSelectorMetadata) ResourceVersion() TpmAttestationCertPathSelectorMetadataResourceVersion {
	return TpmAttestationCertPathSelectorMetadataResourceVersion{}
}

func (TpmAttestationCertPathSelectorMetadata) OwnerReferences() TpmAttestationCertPathSelectorMetadataOwnerReferences {
	return TpmAttestationCertPathSelectorMetadataOwnerReferences{}
}

func (TpmAttestationCertPathSelectorMetadata) Shards() TpmAttestationCertPathSelectorMetadataShards {
	return TpmAttestationCertPathSelectorMetadataShards{}
}

func (TpmAttestationCertPathSelectorMetadata) Syncing() TpmAttestationCertPathSelectorMetadataSyncing {
	return TpmAttestationCertPathSelectorMetadataSyncing{}
}

func (TpmAttestationCertPathSelectorMetadata) Lifecycle() TpmAttestationCertPathSelectorMetadataLifecycle {
	return TpmAttestationCertPathSelectorMetadataLifecycle{}
}

func (TpmAttestationCertPathSelectorMetadata) Services() TpmAttestationCertPathSelectorMetadataServices {
	return TpmAttestationCertPathSelectorMetadataServices{}
}

type TpmAttestationCertPathSelectorMetadataCreateTime struct{}

func (TpmAttestationCertPathSelectorMetadataCreateTime) FieldPath() *TpmAttestationCert_FieldSubPath {
	return &TpmAttestationCert_FieldSubPath{
		selector: TpmAttestationCert_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().CreateTime().FieldPath(),
	}
}

func (s TpmAttestationCertPathSelectorMetadataCreateTime) WithValue(value *timestamppb.Timestamp) *TpmAttestationCert_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*TpmAttestationCert_FieldSubPathValue)
}

func (s TpmAttestationCertPathSelectorMetadataCreateTime) WithArrayOfValues(values []*timestamppb.Timestamp) *TpmAttestationCert_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*TpmAttestationCert_FieldSubPathArrayOfValues)
}

type TpmAttestationCertPathSelectorMetadataUpdateTime struct{}

func (TpmAttestationCertPathSelectorMetadataUpdateTime) FieldPath() *TpmAttestationCert_FieldSubPath {
	return &TpmAttestationCert_FieldSubPath{
		selector: TpmAttestationCert_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().UpdateTime().FieldPath(),
	}
}

func (s TpmAttestationCertPathSelectorMetadataUpdateTime) WithValue(value *timestamppb.Timestamp) *TpmAttestationCert_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*TpmAttestationCert_FieldSubPathValue)
}

func (s TpmAttestationCertPathSelectorMetadataUpdateTime) WithArrayOfValues(values []*timestamppb.Timestamp) *TpmAttestationCert_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*TpmAttestationCert_FieldSubPathArrayOfValues)
}

type TpmAttestationCertPathSelectorMetadataDeleteTime struct{}

func (TpmAttestationCertPathSelectorMetadataDeleteTime) FieldPath() *TpmAttestationCert_FieldSubPath {
	return &TpmAttestationCert_FieldSubPath{
		selector: TpmAttestationCert_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().DeleteTime().FieldPath(),
	}
}

func (s TpmAttestationCertPathSelectorMetadataDeleteTime) WithValue(value *timestamppb.Timestamp) *TpmAttestationCert_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*TpmAttestationCert_FieldSubPathValue)
}

func (s TpmAttestationCertPathSelectorMetadataDeleteTime) WithArrayOfValues(values []*timestamppb.Timestamp) *TpmAttestationCert_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*TpmAttestationCert_FieldSubPathArrayOfValues)
}

type TpmAttestationCertPathSelectorMetadataUuid struct{}

func (TpmAttestationCertPathSelectorMetadataUuid) FieldPath() *TpmAttestationCert_FieldSubPath {
	return &TpmAttestationCert_FieldSubPath{
		selector: TpmAttestationCert_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Uuid().FieldPath(),
	}
}

func (s TpmAttestationCertPathSelectorMetadataUuid) WithValue(value string) *TpmAttestationCert_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*TpmAttestationCert_FieldSubPathValue)
}

func (s TpmAttestationCertPathSelectorMetadataUuid) WithArrayOfValues(values []string) *TpmAttestationCert_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*TpmAttestationCert_FieldSubPathArrayOfValues)
}

type TpmAttestationCertPathSelectorMetadataTags struct{}

func (TpmAttestationCertPathSelectorMetadataTags) FieldPath() *TpmAttestationCert_FieldSubPath {
	return &TpmAttestationCert_FieldSubPath{
		selector: TpmAttestationCert_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Tags().FieldPath(),
	}
}

func (s TpmAttestationCertPathSelectorMetadataTags) WithValue(value []string) *TpmAttestationCert_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*TpmAttestationCert_FieldSubPathValue)
}

func (s TpmAttestationCertPathSelectorMetadataTags) WithArrayOfValues(values [][]string) *TpmAttestationCert_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*TpmAttestationCert_FieldSubPathArrayOfValues)
}

func (s TpmAttestationCertPathSelectorMetadataTags) WithItemValue(value string) *TpmAttestationCert_FieldSubPathArrayItemValue {
	return s.FieldPath().WithIArrayItemValue(value).(*TpmAttestationCert_FieldSubPathArrayItemValue)
}

type TpmAttestationCertPathSelectorMetadataLabels struct{}

func (TpmAttestationCertPathSelectorMetadataLabels) FieldPath() *TpmAttestationCert_FieldSubPath {
	return &TpmAttestationCert_FieldSubPath{
		selector: TpmAttestationCert_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Labels().FieldPath(),
	}
}

func (s TpmAttestationCertPathSelectorMetadataLabels) WithValue(value map[string]string) *TpmAttestationCert_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*TpmAttestationCert_FieldSubPathValue)
}

func (s TpmAttestationCertPathSelectorMetadataLabels) WithArrayOfValues(values []map[string]string) *TpmAttestationCert_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*TpmAttestationCert_FieldSubPathArrayOfValues)
}

func (TpmAttestationCertPathSelectorMetadataLabels) WithKey(key string) TpmAttestationCertMapPathSelectorMetadataLabels {
	return TpmAttestationCertMapPathSelectorMetadataLabels{key: key}
}

type TpmAttestationCertMapPathSelectorMetadataLabels struct {
	key string
}

func (s TpmAttestationCertMapPathSelectorMetadataLabels) FieldPath() *TpmAttestationCert_FieldSubPath {
	return &TpmAttestationCert_FieldSubPath{
		selector: TpmAttestationCert_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Labels().WithKey(s.key).FieldPath(),
	}
}

func (s TpmAttestationCertMapPathSelectorMetadataLabels) WithValue(value string) *TpmAttestationCert_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*TpmAttestationCert_FieldSubPathValue)
}

func (s TpmAttestationCertMapPathSelectorMetadataLabels) WithArrayOfValues(values []string) *TpmAttestationCert_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*TpmAttestationCert_FieldSubPathArrayOfValues)
}

type TpmAttestationCertPathSelectorMetadataAnnotations struct{}

func (TpmAttestationCertPathSelectorMetadataAnnotations) FieldPath() *TpmAttestationCert_FieldSubPath {
	return &TpmAttestationCert_FieldSubPath{
		selector: TpmAttestationCert_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Annotations().FieldPath(),
	}
}

func (s TpmAttestationCertPathSelectorMetadataAnnotations) WithValue(value map[string]string) *TpmAttestationCert_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*TpmAttestationCert_FieldSubPathValue)
}

func (s TpmAttestationCertPathSelectorMetadataAnnotations) WithArrayOfValues(values []map[string]string) *TpmAttestationCert_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*TpmAttestationCert_FieldSubPathArrayOfValues)
}

func (TpmAttestationCertPathSelectorMetadataAnnotations) WithKey(key string) TpmAttestationCertMapPathSelectorMetadataAnnotations {
	return TpmAttestationCertMapPathSelectorMetadataAnnotations{key: key}
}

type TpmAttestationCertMapPathSelectorMetadataAnnotations struct {
	key string
}

func (s TpmAttestationCertMapPathSelectorMetadataAnnotations) FieldPath() *TpmAttestationCert_FieldSubPath {
	return &TpmAttestationCert_FieldSubPath{
		selector: TpmAttestationCert_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Annotations().WithKey(s.key).FieldPath(),
	}
}

func (s TpmAttestationCertMapPathSelectorMetadataAnnotations) WithValue(value string) *TpmAttestationCert_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*TpmAttestationCert_FieldSubPathValue)
}

func (s TpmAttestationCertMapPathSelectorMetadataAnnotations) WithArrayOfValues(values []string) *TpmAttestationCert_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*TpmAttestationCert_FieldSubPathArrayOfValues)
}

type TpmAttestationCertPathSelectorMetadataGeneration struct{}

func (TpmAttestationCertPathSelectorMetadataGeneration) FieldPath() *TpmAttestationCert_FieldSubPath {
	return &TpmAttestationCert_FieldSubPath{
		selector: TpmAttestationCert_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Generation().FieldPath(),
	}
}

func (s TpmAttestationCertPathSelectorMetadataGeneration) WithValue(value int64) *TpmAttestationCert_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*TpmAttestationCert_FieldSubPathValue)
}

func (s TpmAttestationCertPathSelectorMetadataGeneration) WithArrayOfValues(values []int64) *TpmAttestationCert_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*TpmAttestationCert_FieldSubPathArrayOfValues)
}

type TpmAttestationCertPathSelectorMetadataResourceVersion struct{}

func (TpmAttestationCertPathSelectorMetadataResourceVersion) FieldPath() *TpmAttestationCert_FieldSubPath {
	return &TpmAttestationCert_FieldSubPath{
		selector: TpmAttestationCert_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().ResourceVersion().FieldPath(),
	}
}

func (s TpmAttestationCertPathSelectorMetadataResourceVersion) WithValue(value string) *TpmAttestationCert_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*TpmAttestationCert_FieldSubPathValue)
}

func (s TpmAttestationCertPathSelectorMetadataResourceVersion) WithArrayOfValues(values []string) *TpmAttestationCert_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*TpmAttestationCert_FieldSubPathArrayOfValues)
}

type TpmAttestationCertPathSelectorMetadataOwnerReferences struct{}

func (TpmAttestationCertPathSelectorMetadataOwnerReferences) FieldPath() *TpmAttestationCert_FieldSubPath {
	return &TpmAttestationCert_FieldSubPath{
		selector: TpmAttestationCert_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().OwnerReferences().FieldPath(),
	}
}

func (s TpmAttestationCertPathSelectorMetadataOwnerReferences) WithValue(value []*meta.OwnerReference) *TpmAttestationCert_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*TpmAttestationCert_FieldSubPathValue)
}

func (s TpmAttestationCertPathSelectorMetadataOwnerReferences) WithArrayOfValues(values [][]*meta.OwnerReference) *TpmAttestationCert_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*TpmAttestationCert_FieldSubPathArrayOfValues)
}

func (s TpmAttestationCertPathSelectorMetadataOwnerReferences) WithItemValue(value *meta.OwnerReference) *TpmAttestationCert_FieldSubPathArrayItemValue {
	return s.FieldPath().WithIArrayItemValue(value).(*TpmAttestationCert_FieldSubPathArrayItemValue)
}

func (TpmAttestationCertPathSelectorMetadataOwnerReferences) Kind() TpmAttestationCertPathSelectorMetadataOwnerReferencesKind {
	return TpmAttestationCertPathSelectorMetadataOwnerReferencesKind{}
}

func (TpmAttestationCertPathSelectorMetadataOwnerReferences) Version() TpmAttestationCertPathSelectorMetadataOwnerReferencesVersion {
	return TpmAttestationCertPathSelectorMetadataOwnerReferencesVersion{}
}

func (TpmAttestationCertPathSelectorMetadataOwnerReferences) Name() TpmAttestationCertPathSelectorMetadataOwnerReferencesName {
	return TpmAttestationCertPathSelectorMetadataOwnerReferencesName{}
}

func (TpmAttestationCertPathSelectorMetadataOwnerReferences) Region() TpmAttestationCertPathSelectorMetadataOwnerReferencesRegion {
	return TpmAttestationCertPathSelectorMetadataOwnerReferencesRegion{}
}

func (TpmAttestationCertPathSelectorMetadataOwnerReferences) Controller() TpmAttestationCertPathSelectorMetadataOwnerReferencesController {
	return TpmAttestationCertPathSelectorMetadataOwnerReferencesController{}
}

func (TpmAttestationCertPathSelectorMetadataOwnerReferences) RequiresOwnerReference() TpmAttestationCertPathSelectorMetadataOwnerReferencesRequiresOwnerReference {
	return TpmAttestationCertPathSelectorMetadataOwnerReferencesRequiresOwnerReference{}
}

type TpmAttestationCertPathSelectorMetadataOwnerReferencesKind struct{}

func (TpmAttestationCertPathSelectorMetadataOwnerReferencesKind) FieldPath() *TpmAttestationCert_FieldSubPath {
	return &TpmAttestationCert_FieldSubPath{
		selector: TpmAttestationCert_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().OwnerReferences().Kind().FieldPath(),
	}
}

func (s TpmAttestationCertPathSelectorMetadataOwnerReferencesKind) WithValue(value string) *TpmAttestationCert_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*TpmAttestationCert_FieldSubPathValue)
}

func (s TpmAttestationCertPathSelectorMetadataOwnerReferencesKind) WithArrayOfValues(values []string) *TpmAttestationCert_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*TpmAttestationCert_FieldSubPathArrayOfValues)
}

type TpmAttestationCertPathSelectorMetadataOwnerReferencesVersion struct{}

func (TpmAttestationCertPathSelectorMetadataOwnerReferencesVersion) FieldPath() *TpmAttestationCert_FieldSubPath {
	return &TpmAttestationCert_FieldSubPath{
		selector: TpmAttestationCert_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().OwnerReferences().Version().FieldPath(),
	}
}

func (s TpmAttestationCertPathSelectorMetadataOwnerReferencesVersion) WithValue(value string) *TpmAttestationCert_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*TpmAttestationCert_FieldSubPathValue)
}

func (s TpmAttestationCertPathSelectorMetadataOwnerReferencesVersion) WithArrayOfValues(values []string) *TpmAttestationCert_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*TpmAttestationCert_FieldSubPathArrayOfValues)
}

type TpmAttestationCertPathSelectorMetadataOwnerReferencesName struct{}

func (TpmAttestationCertPathSelectorMetadataOwnerReferencesName) FieldPath() *TpmAttestationCert_FieldSubPath {
	return &TpmAttestationCert_FieldSubPath{
		selector: TpmAttestationCert_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().OwnerReferences().Name().FieldPath(),
	}
}

func (s TpmAttestationCertPathSelectorMetadataOwnerReferencesName) WithValue(value string) *TpmAttestationCert_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*TpmAttestationCert_FieldSubPathValue)
}

func (s TpmAttestationCertPathSelectorMetadataOwnerReferencesName) WithArrayOfValues(values []string) *TpmAttestationCert_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*TpmAttestationCert_FieldSubPathArrayOfValues)
}

type TpmAttestationCertPathSelectorMetadataOwnerReferencesRegion struct{}

func (TpmAttestationCertPathSelectorMetadataOwnerReferencesRegion) FieldPath() *TpmAttestationCert_FieldSubPath {
	return &TpmAttestationCert_FieldSubPath{
		selector: TpmAttestationCert_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().OwnerReferences().Region().FieldPath(),
	}
}

func (s TpmAttestationCertPathSelectorMetadataOwnerReferencesRegion) WithValue(value string) *TpmAttestationCert_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*TpmAttestationCert_FieldSubPathValue)
}

func (s TpmAttestationCertPathSelectorMetadataOwnerReferencesRegion) WithArrayOfValues(values []string) *TpmAttestationCert_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*TpmAttestationCert_FieldSubPathArrayOfValues)
}

type TpmAttestationCertPathSelectorMetadataOwnerReferencesController struct{}

func (TpmAttestationCertPathSelectorMetadataOwnerReferencesController) FieldPath() *TpmAttestationCert_FieldSubPath {
	return &TpmAttestationCert_FieldSubPath{
		selector: TpmAttestationCert_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().OwnerReferences().Controller().FieldPath(),
	}
}

func (s TpmAttestationCertPathSelectorMetadataOwnerReferencesController) WithValue(value bool) *TpmAttestationCert_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*TpmAttestationCert_FieldSubPathValue)
}

func (s TpmAttestationCertPathSelectorMetadataOwnerReferencesController) WithArrayOfValues(values []bool) *TpmAttestationCert_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*TpmAttestationCert_FieldSubPathArrayOfValues)
}

type TpmAttestationCertPathSelectorMetadataOwnerReferencesRequiresOwnerReference struct{}

func (TpmAttestationCertPathSelectorMetadataOwnerReferencesRequiresOwnerReference) FieldPath() *TpmAttestationCert_FieldSubPath {
	return &TpmAttestationCert_FieldSubPath{
		selector: TpmAttestationCert_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().OwnerReferences().RequiresOwnerReference().FieldPath(),
	}
}

func (s TpmAttestationCertPathSelectorMetadataOwnerReferencesRequiresOwnerReference) WithValue(value bool) *TpmAttestationCert_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*TpmAttestationCert_FieldSubPathValue)
}

func (s TpmAttestationCertPathSelectorMetadataOwnerReferencesRequiresOwnerReference) WithArrayOfValues(values []bool) *TpmAttestationCert_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*TpmAttestationCert_FieldSubPathArrayOfValues)
}

type TpmAttestationCertPathSelectorMetadataShards struct{}

func (TpmAttestationCertPathSelectorMetadataShards) FieldPath() *TpmAttestationCert_FieldSubPath {
	return &TpmAttestationCert_FieldSubPath{
		selector: TpmAttestationCert_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Shards().FieldPath(),
	}
}

func (s TpmAttestationCertPathSelectorMetadataShards) WithValue(value map[string]int64) *TpmAttestationCert_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*TpmAttestationCert_FieldSubPathValue)
}

func (s TpmAttestationCertPathSelectorMetadataShards) WithArrayOfValues(values []map[string]int64) *TpmAttestationCert_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*TpmAttestationCert_FieldSubPathArrayOfValues)
}

func (TpmAttestationCertPathSelectorMetadataShards) WithKey(key string) TpmAttestationCertMapPathSelectorMetadataShards {
	return TpmAttestationCertMapPathSelectorMetadataShards{key: key}
}

type TpmAttestationCertMapPathSelectorMetadataShards struct {
	key string
}

func (s TpmAttestationCertMapPathSelectorMetadataShards) FieldPath() *TpmAttestationCert_FieldSubPath {
	return &TpmAttestationCert_FieldSubPath{
		selector: TpmAttestationCert_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Shards().WithKey(s.key).FieldPath(),
	}
}

func (s TpmAttestationCertMapPathSelectorMetadataShards) WithValue(value int64) *TpmAttestationCert_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*TpmAttestationCert_FieldSubPathValue)
}

func (s TpmAttestationCertMapPathSelectorMetadataShards) WithArrayOfValues(values []int64) *TpmAttestationCert_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*TpmAttestationCert_FieldSubPathArrayOfValues)
}

type TpmAttestationCertPathSelectorMetadataSyncing struct{}

func (TpmAttestationCertPathSelectorMetadataSyncing) FieldPath() *TpmAttestationCert_FieldSubPath {
	return &TpmAttestationCert_FieldSubPath{
		selector: TpmAttestationCert_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Syncing().FieldPath(),
	}
}

func (s TpmAttestationCertPathSelectorMetadataSyncing) WithValue(value *meta.SyncingMeta) *TpmAttestationCert_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*TpmAttestationCert_FieldSubPathValue)
}

func (s TpmAttestationCertPathSelectorMetadataSyncing) WithArrayOfValues(values []*meta.SyncingMeta) *TpmAttestationCert_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*TpmAttestationCert_FieldSubPathArrayOfValues)
}

func (TpmAttestationCertPathSelectorMetadataSyncing) OwningRegion() TpmAttestationCertPathSelectorMetadataSyncingOwningRegion {
	return TpmAttestationCertPathSelectorMetadataSyncingOwningRegion{}
}

func (TpmAttestationCertPathSelectorMetadataSyncing) Regions() TpmAttestationCertPathSelectorMetadataSyncingRegions {
	return TpmAttestationCertPathSelectorMetadataSyncingRegions{}
}

type TpmAttestationCertPathSelectorMetadataSyncingOwningRegion struct{}

func (TpmAttestationCertPathSelectorMetadataSyncingOwningRegion) FieldPath() *TpmAttestationCert_FieldSubPath {
	return &TpmAttestationCert_FieldSubPath{
		selector: TpmAttestationCert_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Syncing().OwningRegion().FieldPath(),
	}
}

func (s TpmAttestationCertPathSelectorMetadataSyncingOwningRegion) WithValue(value string) *TpmAttestationCert_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*TpmAttestationCert_FieldSubPathValue)
}

func (s TpmAttestationCertPathSelectorMetadataSyncingOwningRegion) WithArrayOfValues(values []string) *TpmAttestationCert_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*TpmAttestationCert_FieldSubPathArrayOfValues)
}

type TpmAttestationCertPathSelectorMetadataSyncingRegions struct{}

func (TpmAttestationCertPathSelectorMetadataSyncingRegions) FieldPath() *TpmAttestationCert_FieldSubPath {
	return &TpmAttestationCert_FieldSubPath{
		selector: TpmAttestationCert_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Syncing().Regions().FieldPath(),
	}
}

func (s TpmAttestationCertPathSelectorMetadataSyncingRegions) WithValue(value []string) *TpmAttestationCert_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*TpmAttestationCert_FieldSubPathValue)
}

func (s TpmAttestationCertPathSelectorMetadataSyncingRegions) WithArrayOfValues(values [][]string) *TpmAttestationCert_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*TpmAttestationCert_FieldSubPathArrayOfValues)
}

func (s TpmAttestationCertPathSelectorMetadataSyncingRegions) WithItemValue(value string) *TpmAttestationCert_FieldSubPathArrayItemValue {
	return s.FieldPath().WithIArrayItemValue(value).(*TpmAttestationCert_FieldSubPathArrayItemValue)
}

type TpmAttestationCertPathSelectorMetadataLifecycle struct{}

func (TpmAttestationCertPathSelectorMetadataLifecycle) FieldPath() *TpmAttestationCert_FieldSubPath {
	return &TpmAttestationCert_FieldSubPath{
		selector: TpmAttestationCert_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Lifecycle().FieldPath(),
	}
}

func (s TpmAttestationCertPathSelectorMetadataLifecycle) WithValue(value *meta.Lifecycle) *TpmAttestationCert_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*TpmAttestationCert_FieldSubPathValue)
}

func (s TpmAttestationCertPathSelectorMetadataLifecycle) WithArrayOfValues(values []*meta.Lifecycle) *TpmAttestationCert_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*TpmAttestationCert_FieldSubPathArrayOfValues)
}

func (TpmAttestationCertPathSelectorMetadataLifecycle) State() TpmAttestationCertPathSelectorMetadataLifecycleState {
	return TpmAttestationCertPathSelectorMetadataLifecycleState{}
}

func (TpmAttestationCertPathSelectorMetadataLifecycle) BlockDeletion() TpmAttestationCertPathSelectorMetadataLifecycleBlockDeletion {
	return TpmAttestationCertPathSelectorMetadataLifecycleBlockDeletion{}
}

type TpmAttestationCertPathSelectorMetadataLifecycleState struct{}

func (TpmAttestationCertPathSelectorMetadataLifecycleState) FieldPath() *TpmAttestationCert_FieldSubPath {
	return &TpmAttestationCert_FieldSubPath{
		selector: TpmAttestationCert_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Lifecycle().State().FieldPath(),
	}
}

func (s TpmAttestationCertPathSelectorMetadataLifecycleState) WithValue(value meta.Lifecycle_State) *TpmAttestationCert_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*TpmAttestationCert_FieldSubPathValue)
}

func (s TpmAttestationCertPathSelectorMetadataLifecycleState) WithArrayOfValues(values []meta.Lifecycle_State) *TpmAttestationCert_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*TpmAttestationCert_FieldSubPathArrayOfValues)
}

type TpmAttestationCertPathSelectorMetadataLifecycleBlockDeletion struct{}

func (TpmAttestationCertPathSelectorMetadataLifecycleBlockDeletion) FieldPath() *TpmAttestationCert_FieldSubPath {
	return &TpmAttestationCert_FieldSubPath{
		selector: TpmAttestationCert_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Lifecycle().BlockDeletion().FieldPath(),
	}
}

func (s TpmAttestationCertPathSelectorMetadataLifecycleBlockDeletion) WithValue(value bool) *TpmAttestationCert_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*TpmAttestationCert_FieldSubPathValue)
}

func (s TpmAttestationCertPathSelectorMetadataLifecycleBlockDeletion) WithArrayOfValues(values []bool) *TpmAttestationCert_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*TpmAttestationCert_FieldSubPathArrayOfValues)
}

type TpmAttestationCertPathSelectorMetadataServices struct{}

func (TpmAttestationCertPathSelectorMetadataServices) FieldPath() *TpmAttestationCert_FieldSubPath {
	return &TpmAttestationCert_FieldSubPath{
		selector: TpmAttestationCert_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Services().FieldPath(),
	}
}

func (s TpmAttestationCertPathSelectorMetadataServices) WithValue(value *meta.ServicesInfo) *TpmAttestationCert_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*TpmAttestationCert_FieldSubPathValue)
}

func (s TpmAttestationCertPathSelectorMetadataServices) WithArrayOfValues(values []*meta.ServicesInfo) *TpmAttestationCert_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*TpmAttestationCert_FieldSubPathArrayOfValues)
}

func (TpmAttestationCertPathSelectorMetadataServices) OwningService() TpmAttestationCertPathSelectorMetadataServicesOwningService {
	return TpmAttestationCertPathSelectorMetadataServicesOwningService{}
}

func (TpmAttestationCertPathSelectorMetadataServices) AllowedServices() TpmAttestationCertPathSelectorMetadataServicesAllowedServices {
	return TpmAttestationCertPathSelectorMetadataServicesAllowedServices{}
}

type TpmAttestationCertPathSelectorMetadataServicesOwningService struct{}

func (TpmAttestationCertPathSelectorMetadataServicesOwningService) FieldPath() *TpmAttestationCert_FieldSubPath {
	return &TpmAttestationCert_FieldSubPath{
		selector: TpmAttestationCert_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Services().OwningService().FieldPath(),
	}
}

func (s TpmAttestationCertPathSelectorMetadataServicesOwningService) WithValue(value string) *TpmAttestationCert_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*TpmAttestationCert_FieldSubPathValue)
}

func (s TpmAttestationCertPathSelectorMetadataServicesOwningService) WithArrayOfValues(values []string) *TpmAttestationCert_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*TpmAttestationCert_FieldSubPathArrayOfValues)
}

type TpmAttestationCertPathSelectorMetadataServicesAllowedServices struct{}

func (TpmAttestationCertPathSelectorMetadataServicesAllowedServices) FieldPath() *TpmAttestationCert_FieldSubPath {
	return &TpmAttestationCert_FieldSubPath{
		selector: TpmAttestationCert_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Services().AllowedServices().FieldPath(),
	}
}

func (s TpmAttestationCertPathSelectorMetadataServicesAllowedServices) WithValue(value []string) *TpmAttestationCert_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*TpmAttestationCert_FieldSubPathValue)
}

func (s TpmAttestationCertPathSelectorMetadataServicesAllowedServices) WithArrayOfValues(values [][]string) *TpmAttestationCert_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*TpmAttestationCert_FieldSubPathArrayOfValues)
}

func (s TpmAttestationCertPathSelectorMetadataServicesAllowedServices) WithItemValue(value string) *TpmAttestationCert_FieldSubPathArrayItemValue {
	return s.FieldPath().WithIArrayItemValue(value).(*TpmAttestationCert_FieldSubPathArrayItemValue)
}

type TpmAttestationCertPathSelectorDisplayName struct{}

func (TpmAttestationCertPathSelectorDisplayName) FieldPath() *TpmAttestationCert_FieldTerminalPath {
	return &TpmAttestationCert_FieldTerminalPath{selector: TpmAttestationCert_FieldPathSelectorDisplayName}
}

func (s TpmAttestationCertPathSelectorDisplayName) WithValue(value string) *TpmAttestationCert_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*TpmAttestationCert_FieldTerminalPathValue)
}

func (s TpmAttestationCertPathSelectorDisplayName) WithArrayOfValues(values []string) *TpmAttestationCert_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*TpmAttestationCert_FieldTerminalPathArrayOfValues)
}

type TpmAttestationCertPathSelectorManufacturer struct{}

func (TpmAttestationCertPathSelectorManufacturer) FieldPath() *TpmAttestationCert_FieldTerminalPath {
	return &TpmAttestationCert_FieldTerminalPath{selector: TpmAttestationCert_FieldPathSelectorManufacturer}
}

func (s TpmAttestationCertPathSelectorManufacturer) WithValue(value string) *TpmAttestationCert_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*TpmAttestationCert_FieldTerminalPathValue)
}

func (s TpmAttestationCertPathSelectorManufacturer) WithArrayOfValues(values []string) *TpmAttestationCert_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*TpmAttestationCert_FieldTerminalPathArrayOfValues)
}

type TpmAttestationCertPathSelectorProductName struct{}

func (TpmAttestationCertPathSelectorProductName) FieldPath() *TpmAttestationCert_FieldTerminalPath {
	return &TpmAttestationCert_FieldTerminalPath{selector: TpmAttestationCert_FieldPathSelectorProductName}
}

func (s TpmAttestationCertPathSelectorProductName) WithValue(value string) *TpmAttestationCert_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*TpmAttestationCert_FieldTerminalPathValue)
}

func (s TpmAttestationCertPathSelectorProductName) WithArrayOfValues(values []string) *TpmAttestationCert_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*TpmAttestationCert_FieldTerminalPathArrayOfValues)
}

type TpmAttestationCertPathSelectorTpmManufacturerCaCert struct{}

func (TpmAttestationCertPathSelectorTpmManufacturerCaCert) FieldPath() *TpmAttestationCert_FieldTerminalPath {
	return &TpmAttestationCert_FieldTerminalPath{selector: TpmAttestationCert_FieldPathSelectorTpmManufacturerCaCert}
}

func (s TpmAttestationCertPathSelectorTpmManufacturerCaCert) WithValue(value string) *TpmAttestationCert_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*TpmAttestationCert_FieldTerminalPathValue)
}

func (s TpmAttestationCertPathSelectorTpmManufacturerCaCert) WithArrayOfValues(values []string) *TpmAttestationCert_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*TpmAttestationCert_FieldTerminalPathArrayOfValues)
}

type TpmAttestationCertPathSelectorIdevidIssuerCaCert struct{}

func (TpmAttestationCertPathSelectorIdevidIssuerCaCert) FieldPath() *TpmAttestationCert_FieldTerminalPath {
	return &TpmAttestationCert_FieldTerminalPath{selector: TpmAttestationCert_FieldPathSelectorIdevidIssuerCaCert}
}

func (s TpmAttestationCertPathSelectorIdevidIssuerCaCert) WithValue(value string) *TpmAttestationCert_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*TpmAttestationCert_FieldTerminalPathValue)
}

func (s TpmAttestationCertPathSelectorIdevidIssuerCaCert) WithArrayOfValues(values []string) *TpmAttestationCert_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*TpmAttestationCert_FieldTerminalPathArrayOfValues)
}

type TpmAttestationCertPathSelectorLdevidIssuerCaCert struct{}

func (TpmAttestationCertPathSelectorLdevidIssuerCaCert) FieldPath() *TpmAttestationCert_FieldTerminalPath {
	return &TpmAttestationCert_FieldTerminalPath{selector: TpmAttestationCert_FieldPathSelectorLdevidIssuerCaCert}
}

func (s TpmAttestationCertPathSelectorLdevidIssuerCaCert) WithValue(value string) *TpmAttestationCert_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*TpmAttestationCert_FieldTerminalPathValue)
}

func (s TpmAttestationCertPathSelectorLdevidIssuerCaCert) WithArrayOfValues(values []string) *TpmAttestationCert_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*TpmAttestationCert_FieldTerminalPathArrayOfValues)
}