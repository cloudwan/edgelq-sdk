// Code generated by protoc-gen-goten-object
// File: edgelq/devices/proto/v1alpha2/os_image_profile.proto
// DO NOT EDIT!!!

package os_image_profile

// proto imports
import (
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
	multi_region_policy "github.com/cloudwan/edgelq-sdk/common/types/multi_region_policy"
	device_type "github.com/cloudwan/edgelq-sdk/devices/resources/v1alpha2/device_type"
	project "github.com/cloudwan/edgelq-sdk/devices/resources/v1alpha2/project"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
)

// make sure we're using proto imports
var (
	_ = &ntt_meta.Meta{}
	_ = &multi_region_policy.MultiRegionPolicy{}
	_ = &device_type.DeviceType{}
	_ = &project.Project{}
	_ = &timestamp.Timestamp{}
)

type OsImageProfileFieldPathBuilder struct{}

func NewOsImageProfileFieldPathBuilder() OsImageProfileFieldPathBuilder {
	return OsImageProfileFieldPathBuilder{}
}
func (OsImageProfileFieldPathBuilder) Name() OsImageProfilePathSelectorName {
	return OsImageProfilePathSelectorName{}
}
func (OsImageProfileFieldPathBuilder) Metadata() OsImageProfilePathSelectorMetadata {
	return OsImageProfilePathSelectorMetadata{}
}
func (OsImageProfileFieldPathBuilder) DisplayName() OsImageProfilePathSelectorDisplayName {
	return OsImageProfilePathSelectorDisplayName{}
}
func (OsImageProfileFieldPathBuilder) DeviceType() OsImageProfilePathSelectorDeviceType {
	return OsImageProfilePathSelectorDeviceType{}
}
func (OsImageProfileFieldPathBuilder) InstallAiAccelerator() OsImageProfilePathSelectorInstallAiAccelerator {
	return OsImageProfilePathSelectorInstallAiAccelerator{}
}
func (OsImageProfileFieldPathBuilder) Encryption() OsImageProfilePathSelectorEncryption {
	return OsImageProfilePathSelectorEncryption{}
}
func (OsImageProfileFieldPathBuilder) DiskMapping() OsImageProfilePathSelectorDiskMapping {
	return OsImageProfilePathSelectorDiskMapping{}
}
func (OsImageProfileFieldPathBuilder) NetworkAgent() OsImageProfilePathSelectorNetworkAgent {
	return OsImageProfilePathSelectorNetworkAgent{}
}
func (OsImageProfileFieldPathBuilder) Ntp() OsImageProfilePathSelectorNtp {
	return OsImageProfilePathSelectorNtp{}
}
func (OsImageProfileFieldPathBuilder) HttpProxy() OsImageProfilePathSelectorHttpProxy {
	return OsImageProfilePathSelectorHttpProxy{}
}
func (OsImageProfileFieldPathBuilder) HttpsProxy() OsImageProfilePathSelectorHttpsProxy {
	return OsImageProfilePathSelectorHttpsProxy{}
}
func (OsImageProfileFieldPathBuilder) NoProxy() OsImageProfilePathSelectorNoProxy {
	return OsImageProfilePathSelectorNoProxy{}
}

type OsImageProfilePathSelectorName struct{}

func (OsImageProfilePathSelectorName) FieldPath() *OsImageProfile_FieldTerminalPath {
	return &OsImageProfile_FieldTerminalPath{selector: OsImageProfile_FieldPathSelectorName}
}

func (s OsImageProfilePathSelectorName) WithValue(value *Name) *OsImageProfile_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*OsImageProfile_FieldTerminalPathValue)
}

func (s OsImageProfilePathSelectorName) WithArrayOfValues(values []*Name) *OsImageProfile_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*OsImageProfile_FieldTerminalPathArrayOfValues)
}

type OsImageProfilePathSelectorMetadata struct{}

func (OsImageProfilePathSelectorMetadata) FieldPath() *OsImageProfile_FieldTerminalPath {
	return &OsImageProfile_FieldTerminalPath{selector: OsImageProfile_FieldPathSelectorMetadata}
}

func (s OsImageProfilePathSelectorMetadata) WithValue(value *ntt_meta.Meta) *OsImageProfile_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*OsImageProfile_FieldTerminalPathValue)
}

func (s OsImageProfilePathSelectorMetadata) WithArrayOfValues(values []*ntt_meta.Meta) *OsImageProfile_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*OsImageProfile_FieldTerminalPathArrayOfValues)
}

func (OsImageProfilePathSelectorMetadata) WithSubPath(subPath ntt_meta.Meta_FieldPath) *OsImageProfile_FieldSubPath {
	return &OsImageProfile_FieldSubPath{selector: OsImageProfile_FieldPathSelectorMetadata, subPath: subPath}
}

func (s OsImageProfilePathSelectorMetadata) WithSubValue(subPathValue ntt_meta.Meta_FieldPathValue) *OsImageProfile_FieldSubPathValue {
	return &OsImageProfile_FieldSubPathValue{OsImageProfile_FieldPath: s.WithSubPath(subPathValue), subPathValue: subPathValue}
}

func (s OsImageProfilePathSelectorMetadata) WithSubArrayOfValues(subPathArrayOfValues ntt_meta.Meta_FieldPathArrayOfValues) *OsImageProfile_FieldSubPathArrayOfValues {
	return &OsImageProfile_FieldSubPathArrayOfValues{OsImageProfile_FieldPath: s.WithSubPath(subPathArrayOfValues), subPathArrayOfValues: subPathArrayOfValues}
}

func (s OsImageProfilePathSelectorMetadata) WithSubArrayItemValue(subPathArrayItemValue ntt_meta.Meta_FieldPathArrayItemValue) *OsImageProfile_FieldSubPathArrayItemValue {
	return &OsImageProfile_FieldSubPathArrayItemValue{OsImageProfile_FieldPath: s.WithSubPath(subPathArrayItemValue), subPathItemValue: subPathArrayItemValue}
}

func (OsImageProfilePathSelectorMetadata) CreateTime() OsImageProfilePathSelectorMetadataCreateTime {
	return OsImageProfilePathSelectorMetadataCreateTime{}
}

func (OsImageProfilePathSelectorMetadata) UpdateTime() OsImageProfilePathSelectorMetadataUpdateTime {
	return OsImageProfilePathSelectorMetadataUpdateTime{}
}

func (OsImageProfilePathSelectorMetadata) DeleteTime() OsImageProfilePathSelectorMetadataDeleteTime {
	return OsImageProfilePathSelectorMetadataDeleteTime{}
}

func (OsImageProfilePathSelectorMetadata) Uuid() OsImageProfilePathSelectorMetadataUuid {
	return OsImageProfilePathSelectorMetadataUuid{}
}

func (OsImageProfilePathSelectorMetadata) Tags() OsImageProfilePathSelectorMetadataTags {
	return OsImageProfilePathSelectorMetadataTags{}
}

func (OsImageProfilePathSelectorMetadata) Labels() OsImageProfilePathSelectorMetadataLabels {
	return OsImageProfilePathSelectorMetadataLabels{}
}

func (OsImageProfilePathSelectorMetadata) Annotations() OsImageProfilePathSelectorMetadataAnnotations {
	return OsImageProfilePathSelectorMetadataAnnotations{}
}

func (OsImageProfilePathSelectorMetadata) Generation() OsImageProfilePathSelectorMetadataGeneration {
	return OsImageProfilePathSelectorMetadataGeneration{}
}

func (OsImageProfilePathSelectorMetadata) ResourceVersion() OsImageProfilePathSelectorMetadataResourceVersion {
	return OsImageProfilePathSelectorMetadataResourceVersion{}
}

func (OsImageProfilePathSelectorMetadata) OwnerReferences() OsImageProfilePathSelectorMetadataOwnerReferences {
	return OsImageProfilePathSelectorMetadataOwnerReferences{}
}

func (OsImageProfilePathSelectorMetadata) Shards() OsImageProfilePathSelectorMetadataShards {
	return OsImageProfilePathSelectorMetadataShards{}
}

func (OsImageProfilePathSelectorMetadata) Syncing() OsImageProfilePathSelectorMetadataSyncing {
	return OsImageProfilePathSelectorMetadataSyncing{}
}

func (OsImageProfilePathSelectorMetadata) Lifecycle() OsImageProfilePathSelectorMetadataLifecycle {
	return OsImageProfilePathSelectorMetadataLifecycle{}
}

type OsImageProfilePathSelectorMetadataCreateTime struct{}

func (OsImageProfilePathSelectorMetadataCreateTime) FieldPath() *OsImageProfile_FieldSubPath {
	return &OsImageProfile_FieldSubPath{
		selector: OsImageProfile_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().CreateTime().FieldPath(),
	}
}

func (s OsImageProfilePathSelectorMetadataCreateTime) WithValue(value *timestamp.Timestamp) *OsImageProfile_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*OsImageProfile_FieldSubPathValue)
}

func (s OsImageProfilePathSelectorMetadataCreateTime) WithArrayOfValues(values []*timestamp.Timestamp) *OsImageProfile_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*OsImageProfile_FieldSubPathArrayOfValues)
}

type OsImageProfilePathSelectorMetadataUpdateTime struct{}

func (OsImageProfilePathSelectorMetadataUpdateTime) FieldPath() *OsImageProfile_FieldSubPath {
	return &OsImageProfile_FieldSubPath{
		selector: OsImageProfile_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().UpdateTime().FieldPath(),
	}
}

func (s OsImageProfilePathSelectorMetadataUpdateTime) WithValue(value *timestamp.Timestamp) *OsImageProfile_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*OsImageProfile_FieldSubPathValue)
}

func (s OsImageProfilePathSelectorMetadataUpdateTime) WithArrayOfValues(values []*timestamp.Timestamp) *OsImageProfile_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*OsImageProfile_FieldSubPathArrayOfValues)
}

type OsImageProfilePathSelectorMetadataDeleteTime struct{}

func (OsImageProfilePathSelectorMetadataDeleteTime) FieldPath() *OsImageProfile_FieldSubPath {
	return &OsImageProfile_FieldSubPath{
		selector: OsImageProfile_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().DeleteTime().FieldPath(),
	}
}

func (s OsImageProfilePathSelectorMetadataDeleteTime) WithValue(value *timestamp.Timestamp) *OsImageProfile_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*OsImageProfile_FieldSubPathValue)
}

func (s OsImageProfilePathSelectorMetadataDeleteTime) WithArrayOfValues(values []*timestamp.Timestamp) *OsImageProfile_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*OsImageProfile_FieldSubPathArrayOfValues)
}

type OsImageProfilePathSelectorMetadataUuid struct{}

func (OsImageProfilePathSelectorMetadataUuid) FieldPath() *OsImageProfile_FieldSubPath {
	return &OsImageProfile_FieldSubPath{
		selector: OsImageProfile_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Uuid().FieldPath(),
	}
}

func (s OsImageProfilePathSelectorMetadataUuid) WithValue(value string) *OsImageProfile_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*OsImageProfile_FieldSubPathValue)
}

func (s OsImageProfilePathSelectorMetadataUuid) WithArrayOfValues(values []string) *OsImageProfile_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*OsImageProfile_FieldSubPathArrayOfValues)
}

type OsImageProfilePathSelectorMetadataTags struct{}

func (OsImageProfilePathSelectorMetadataTags) FieldPath() *OsImageProfile_FieldSubPath {
	return &OsImageProfile_FieldSubPath{
		selector: OsImageProfile_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Tags().FieldPath(),
	}
}

func (s OsImageProfilePathSelectorMetadataTags) WithValue(value []string) *OsImageProfile_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*OsImageProfile_FieldSubPathValue)
}

func (s OsImageProfilePathSelectorMetadataTags) WithArrayOfValues(values [][]string) *OsImageProfile_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*OsImageProfile_FieldSubPathArrayOfValues)
}

func (s OsImageProfilePathSelectorMetadataTags) WithItemValue(value string) *OsImageProfile_FieldSubPathArrayItemValue {
	return s.FieldPath().WithIArrayItemValue(value).(*OsImageProfile_FieldSubPathArrayItemValue)
}

type OsImageProfilePathSelectorMetadataLabels struct{}

func (OsImageProfilePathSelectorMetadataLabels) FieldPath() *OsImageProfile_FieldSubPath {
	return &OsImageProfile_FieldSubPath{
		selector: OsImageProfile_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Labels().FieldPath(),
	}
}

func (s OsImageProfilePathSelectorMetadataLabels) WithValue(value map[string]string) *OsImageProfile_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*OsImageProfile_FieldSubPathValue)
}

func (s OsImageProfilePathSelectorMetadataLabels) WithArrayOfValues(values []map[string]string) *OsImageProfile_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*OsImageProfile_FieldSubPathArrayOfValues)
}

func (OsImageProfilePathSelectorMetadataLabels) WithKey(key string) OsImageProfileMapPathSelectorMetadataLabels {
	return OsImageProfileMapPathSelectorMetadataLabels{key: key}
}

type OsImageProfileMapPathSelectorMetadataLabels struct {
	key string
}

func (s OsImageProfileMapPathSelectorMetadataLabels) FieldPath() *OsImageProfile_FieldSubPath {
	return &OsImageProfile_FieldSubPath{
		selector: OsImageProfile_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Labels().WithKey(s.key).FieldPath(),
	}
}

func (s OsImageProfileMapPathSelectorMetadataLabels) WithValue(value string) *OsImageProfile_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*OsImageProfile_FieldSubPathValue)
}

func (s OsImageProfileMapPathSelectorMetadataLabels) WithArrayOfValues(values []string) *OsImageProfile_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*OsImageProfile_FieldSubPathArrayOfValues)
}

type OsImageProfilePathSelectorMetadataAnnotations struct{}

func (OsImageProfilePathSelectorMetadataAnnotations) FieldPath() *OsImageProfile_FieldSubPath {
	return &OsImageProfile_FieldSubPath{
		selector: OsImageProfile_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Annotations().FieldPath(),
	}
}

func (s OsImageProfilePathSelectorMetadataAnnotations) WithValue(value map[string]string) *OsImageProfile_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*OsImageProfile_FieldSubPathValue)
}

func (s OsImageProfilePathSelectorMetadataAnnotations) WithArrayOfValues(values []map[string]string) *OsImageProfile_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*OsImageProfile_FieldSubPathArrayOfValues)
}

func (OsImageProfilePathSelectorMetadataAnnotations) WithKey(key string) OsImageProfileMapPathSelectorMetadataAnnotations {
	return OsImageProfileMapPathSelectorMetadataAnnotations{key: key}
}

type OsImageProfileMapPathSelectorMetadataAnnotations struct {
	key string
}

func (s OsImageProfileMapPathSelectorMetadataAnnotations) FieldPath() *OsImageProfile_FieldSubPath {
	return &OsImageProfile_FieldSubPath{
		selector: OsImageProfile_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Annotations().WithKey(s.key).FieldPath(),
	}
}

func (s OsImageProfileMapPathSelectorMetadataAnnotations) WithValue(value string) *OsImageProfile_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*OsImageProfile_FieldSubPathValue)
}

func (s OsImageProfileMapPathSelectorMetadataAnnotations) WithArrayOfValues(values []string) *OsImageProfile_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*OsImageProfile_FieldSubPathArrayOfValues)
}

type OsImageProfilePathSelectorMetadataGeneration struct{}

func (OsImageProfilePathSelectorMetadataGeneration) FieldPath() *OsImageProfile_FieldSubPath {
	return &OsImageProfile_FieldSubPath{
		selector: OsImageProfile_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Generation().FieldPath(),
	}
}

func (s OsImageProfilePathSelectorMetadataGeneration) WithValue(value int64) *OsImageProfile_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*OsImageProfile_FieldSubPathValue)
}

func (s OsImageProfilePathSelectorMetadataGeneration) WithArrayOfValues(values []int64) *OsImageProfile_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*OsImageProfile_FieldSubPathArrayOfValues)
}

type OsImageProfilePathSelectorMetadataResourceVersion struct{}

func (OsImageProfilePathSelectorMetadataResourceVersion) FieldPath() *OsImageProfile_FieldSubPath {
	return &OsImageProfile_FieldSubPath{
		selector: OsImageProfile_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().ResourceVersion().FieldPath(),
	}
}

func (s OsImageProfilePathSelectorMetadataResourceVersion) WithValue(value string) *OsImageProfile_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*OsImageProfile_FieldSubPathValue)
}

func (s OsImageProfilePathSelectorMetadataResourceVersion) WithArrayOfValues(values []string) *OsImageProfile_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*OsImageProfile_FieldSubPathArrayOfValues)
}

type OsImageProfilePathSelectorMetadataOwnerReferences struct{}

func (OsImageProfilePathSelectorMetadataOwnerReferences) FieldPath() *OsImageProfile_FieldSubPath {
	return &OsImageProfile_FieldSubPath{
		selector: OsImageProfile_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().FieldPath(),
	}
}

func (s OsImageProfilePathSelectorMetadataOwnerReferences) WithValue(value []*ntt_meta.OwnerReference) *OsImageProfile_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*OsImageProfile_FieldSubPathValue)
}

func (s OsImageProfilePathSelectorMetadataOwnerReferences) WithArrayOfValues(values [][]*ntt_meta.OwnerReference) *OsImageProfile_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*OsImageProfile_FieldSubPathArrayOfValues)
}

func (s OsImageProfilePathSelectorMetadataOwnerReferences) WithItemValue(value *ntt_meta.OwnerReference) *OsImageProfile_FieldSubPathArrayItemValue {
	return s.FieldPath().WithIArrayItemValue(value).(*OsImageProfile_FieldSubPathArrayItemValue)
}

func (OsImageProfilePathSelectorMetadataOwnerReferences) Kind() OsImageProfilePathSelectorMetadataOwnerReferencesKind {
	return OsImageProfilePathSelectorMetadataOwnerReferencesKind{}
}

func (OsImageProfilePathSelectorMetadataOwnerReferences) Version() OsImageProfilePathSelectorMetadataOwnerReferencesVersion {
	return OsImageProfilePathSelectorMetadataOwnerReferencesVersion{}
}

func (OsImageProfilePathSelectorMetadataOwnerReferences) Name() OsImageProfilePathSelectorMetadataOwnerReferencesName {
	return OsImageProfilePathSelectorMetadataOwnerReferencesName{}
}

func (OsImageProfilePathSelectorMetadataOwnerReferences) Region() OsImageProfilePathSelectorMetadataOwnerReferencesRegion {
	return OsImageProfilePathSelectorMetadataOwnerReferencesRegion{}
}

func (OsImageProfilePathSelectorMetadataOwnerReferences) Controller() OsImageProfilePathSelectorMetadataOwnerReferencesController {
	return OsImageProfilePathSelectorMetadataOwnerReferencesController{}
}

func (OsImageProfilePathSelectorMetadataOwnerReferences) BlockOwnerDeletion() OsImageProfilePathSelectorMetadataOwnerReferencesBlockOwnerDeletion {
	return OsImageProfilePathSelectorMetadataOwnerReferencesBlockOwnerDeletion{}
}

func (OsImageProfilePathSelectorMetadataOwnerReferences) RequiresOwnerReference() OsImageProfilePathSelectorMetadataOwnerReferencesRequiresOwnerReference {
	return OsImageProfilePathSelectorMetadataOwnerReferencesRequiresOwnerReference{}
}

type OsImageProfilePathSelectorMetadataOwnerReferencesKind struct{}

func (OsImageProfilePathSelectorMetadataOwnerReferencesKind) FieldPath() *OsImageProfile_FieldSubPath {
	return &OsImageProfile_FieldSubPath{
		selector: OsImageProfile_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().Kind().FieldPath(),
	}
}

func (s OsImageProfilePathSelectorMetadataOwnerReferencesKind) WithValue(value string) *OsImageProfile_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*OsImageProfile_FieldSubPathValue)
}

func (s OsImageProfilePathSelectorMetadataOwnerReferencesKind) WithArrayOfValues(values []string) *OsImageProfile_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*OsImageProfile_FieldSubPathArrayOfValues)
}

type OsImageProfilePathSelectorMetadataOwnerReferencesVersion struct{}

func (OsImageProfilePathSelectorMetadataOwnerReferencesVersion) FieldPath() *OsImageProfile_FieldSubPath {
	return &OsImageProfile_FieldSubPath{
		selector: OsImageProfile_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().Version().FieldPath(),
	}
}

func (s OsImageProfilePathSelectorMetadataOwnerReferencesVersion) WithValue(value string) *OsImageProfile_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*OsImageProfile_FieldSubPathValue)
}

func (s OsImageProfilePathSelectorMetadataOwnerReferencesVersion) WithArrayOfValues(values []string) *OsImageProfile_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*OsImageProfile_FieldSubPathArrayOfValues)
}

type OsImageProfilePathSelectorMetadataOwnerReferencesName struct{}

func (OsImageProfilePathSelectorMetadataOwnerReferencesName) FieldPath() *OsImageProfile_FieldSubPath {
	return &OsImageProfile_FieldSubPath{
		selector: OsImageProfile_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().Name().FieldPath(),
	}
}

func (s OsImageProfilePathSelectorMetadataOwnerReferencesName) WithValue(value string) *OsImageProfile_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*OsImageProfile_FieldSubPathValue)
}

func (s OsImageProfilePathSelectorMetadataOwnerReferencesName) WithArrayOfValues(values []string) *OsImageProfile_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*OsImageProfile_FieldSubPathArrayOfValues)
}

type OsImageProfilePathSelectorMetadataOwnerReferencesRegion struct{}

func (OsImageProfilePathSelectorMetadataOwnerReferencesRegion) FieldPath() *OsImageProfile_FieldSubPath {
	return &OsImageProfile_FieldSubPath{
		selector: OsImageProfile_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().Region().FieldPath(),
	}
}

func (s OsImageProfilePathSelectorMetadataOwnerReferencesRegion) WithValue(value string) *OsImageProfile_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*OsImageProfile_FieldSubPathValue)
}

func (s OsImageProfilePathSelectorMetadataOwnerReferencesRegion) WithArrayOfValues(values []string) *OsImageProfile_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*OsImageProfile_FieldSubPathArrayOfValues)
}

type OsImageProfilePathSelectorMetadataOwnerReferencesController struct{}

func (OsImageProfilePathSelectorMetadataOwnerReferencesController) FieldPath() *OsImageProfile_FieldSubPath {
	return &OsImageProfile_FieldSubPath{
		selector: OsImageProfile_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().Controller().FieldPath(),
	}
}

func (s OsImageProfilePathSelectorMetadataOwnerReferencesController) WithValue(value bool) *OsImageProfile_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*OsImageProfile_FieldSubPathValue)
}

func (s OsImageProfilePathSelectorMetadataOwnerReferencesController) WithArrayOfValues(values []bool) *OsImageProfile_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*OsImageProfile_FieldSubPathArrayOfValues)
}

type OsImageProfilePathSelectorMetadataOwnerReferencesBlockOwnerDeletion struct{}

func (OsImageProfilePathSelectorMetadataOwnerReferencesBlockOwnerDeletion) FieldPath() *OsImageProfile_FieldSubPath {
	return &OsImageProfile_FieldSubPath{
		selector: OsImageProfile_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().BlockOwnerDeletion().FieldPath(),
	}
}

func (s OsImageProfilePathSelectorMetadataOwnerReferencesBlockOwnerDeletion) WithValue(value bool) *OsImageProfile_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*OsImageProfile_FieldSubPathValue)
}

func (s OsImageProfilePathSelectorMetadataOwnerReferencesBlockOwnerDeletion) WithArrayOfValues(values []bool) *OsImageProfile_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*OsImageProfile_FieldSubPathArrayOfValues)
}

type OsImageProfilePathSelectorMetadataOwnerReferencesRequiresOwnerReference struct{}

func (OsImageProfilePathSelectorMetadataOwnerReferencesRequiresOwnerReference) FieldPath() *OsImageProfile_FieldSubPath {
	return &OsImageProfile_FieldSubPath{
		selector: OsImageProfile_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().RequiresOwnerReference().FieldPath(),
	}
}

func (s OsImageProfilePathSelectorMetadataOwnerReferencesRequiresOwnerReference) WithValue(value bool) *OsImageProfile_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*OsImageProfile_FieldSubPathValue)
}

func (s OsImageProfilePathSelectorMetadataOwnerReferencesRequiresOwnerReference) WithArrayOfValues(values []bool) *OsImageProfile_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*OsImageProfile_FieldSubPathArrayOfValues)
}

type OsImageProfilePathSelectorMetadataShards struct{}

func (OsImageProfilePathSelectorMetadataShards) FieldPath() *OsImageProfile_FieldSubPath {
	return &OsImageProfile_FieldSubPath{
		selector: OsImageProfile_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Shards().FieldPath(),
	}
}

func (s OsImageProfilePathSelectorMetadataShards) WithValue(value map[string]int64) *OsImageProfile_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*OsImageProfile_FieldSubPathValue)
}

func (s OsImageProfilePathSelectorMetadataShards) WithArrayOfValues(values []map[string]int64) *OsImageProfile_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*OsImageProfile_FieldSubPathArrayOfValues)
}

func (OsImageProfilePathSelectorMetadataShards) WithKey(key string) OsImageProfileMapPathSelectorMetadataShards {
	return OsImageProfileMapPathSelectorMetadataShards{key: key}
}

type OsImageProfileMapPathSelectorMetadataShards struct {
	key string
}

func (s OsImageProfileMapPathSelectorMetadataShards) FieldPath() *OsImageProfile_FieldSubPath {
	return &OsImageProfile_FieldSubPath{
		selector: OsImageProfile_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Shards().WithKey(s.key).FieldPath(),
	}
}

func (s OsImageProfileMapPathSelectorMetadataShards) WithValue(value int64) *OsImageProfile_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*OsImageProfile_FieldSubPathValue)
}

func (s OsImageProfileMapPathSelectorMetadataShards) WithArrayOfValues(values []int64) *OsImageProfile_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*OsImageProfile_FieldSubPathArrayOfValues)
}

type OsImageProfilePathSelectorMetadataSyncing struct{}

func (OsImageProfilePathSelectorMetadataSyncing) FieldPath() *OsImageProfile_FieldSubPath {
	return &OsImageProfile_FieldSubPath{
		selector: OsImageProfile_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Syncing().FieldPath(),
	}
}

func (s OsImageProfilePathSelectorMetadataSyncing) WithValue(value *ntt_meta.SyncingMeta) *OsImageProfile_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*OsImageProfile_FieldSubPathValue)
}

func (s OsImageProfilePathSelectorMetadataSyncing) WithArrayOfValues(values []*ntt_meta.SyncingMeta) *OsImageProfile_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*OsImageProfile_FieldSubPathArrayOfValues)
}

func (OsImageProfilePathSelectorMetadataSyncing) OwningRegion() OsImageProfilePathSelectorMetadataSyncingOwningRegion {
	return OsImageProfilePathSelectorMetadataSyncingOwningRegion{}
}

func (OsImageProfilePathSelectorMetadataSyncing) Regions() OsImageProfilePathSelectorMetadataSyncingRegions {
	return OsImageProfilePathSelectorMetadataSyncingRegions{}
}

type OsImageProfilePathSelectorMetadataSyncingOwningRegion struct{}

func (OsImageProfilePathSelectorMetadataSyncingOwningRegion) FieldPath() *OsImageProfile_FieldSubPath {
	return &OsImageProfile_FieldSubPath{
		selector: OsImageProfile_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Syncing().OwningRegion().FieldPath(),
	}
}

func (s OsImageProfilePathSelectorMetadataSyncingOwningRegion) WithValue(value string) *OsImageProfile_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*OsImageProfile_FieldSubPathValue)
}

func (s OsImageProfilePathSelectorMetadataSyncingOwningRegion) WithArrayOfValues(values []string) *OsImageProfile_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*OsImageProfile_FieldSubPathArrayOfValues)
}

type OsImageProfilePathSelectorMetadataSyncingRegions struct{}

func (OsImageProfilePathSelectorMetadataSyncingRegions) FieldPath() *OsImageProfile_FieldSubPath {
	return &OsImageProfile_FieldSubPath{
		selector: OsImageProfile_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Syncing().Regions().FieldPath(),
	}
}

func (s OsImageProfilePathSelectorMetadataSyncingRegions) WithValue(value []string) *OsImageProfile_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*OsImageProfile_FieldSubPathValue)
}

func (s OsImageProfilePathSelectorMetadataSyncingRegions) WithArrayOfValues(values [][]string) *OsImageProfile_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*OsImageProfile_FieldSubPathArrayOfValues)
}

func (s OsImageProfilePathSelectorMetadataSyncingRegions) WithItemValue(value string) *OsImageProfile_FieldSubPathArrayItemValue {
	return s.FieldPath().WithIArrayItemValue(value).(*OsImageProfile_FieldSubPathArrayItemValue)
}

type OsImageProfilePathSelectorMetadataLifecycle struct{}

func (OsImageProfilePathSelectorMetadataLifecycle) FieldPath() *OsImageProfile_FieldSubPath {
	return &OsImageProfile_FieldSubPath{
		selector: OsImageProfile_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Lifecycle().FieldPath(),
	}
}

func (s OsImageProfilePathSelectorMetadataLifecycle) WithValue(value *ntt_meta.Lifecycle) *OsImageProfile_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*OsImageProfile_FieldSubPathValue)
}

func (s OsImageProfilePathSelectorMetadataLifecycle) WithArrayOfValues(values []*ntt_meta.Lifecycle) *OsImageProfile_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*OsImageProfile_FieldSubPathArrayOfValues)
}

func (OsImageProfilePathSelectorMetadataLifecycle) State() OsImageProfilePathSelectorMetadataLifecycleState {
	return OsImageProfilePathSelectorMetadataLifecycleState{}
}

func (OsImageProfilePathSelectorMetadataLifecycle) BlockDeletion() OsImageProfilePathSelectorMetadataLifecycleBlockDeletion {
	return OsImageProfilePathSelectorMetadataLifecycleBlockDeletion{}
}

type OsImageProfilePathSelectorMetadataLifecycleState struct{}

func (OsImageProfilePathSelectorMetadataLifecycleState) FieldPath() *OsImageProfile_FieldSubPath {
	return &OsImageProfile_FieldSubPath{
		selector: OsImageProfile_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Lifecycle().State().FieldPath(),
	}
}

func (s OsImageProfilePathSelectorMetadataLifecycleState) WithValue(value ntt_meta.Lifecycle_State) *OsImageProfile_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*OsImageProfile_FieldSubPathValue)
}

func (s OsImageProfilePathSelectorMetadataLifecycleState) WithArrayOfValues(values []ntt_meta.Lifecycle_State) *OsImageProfile_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*OsImageProfile_FieldSubPathArrayOfValues)
}

type OsImageProfilePathSelectorMetadataLifecycleBlockDeletion struct{}

func (OsImageProfilePathSelectorMetadataLifecycleBlockDeletion) FieldPath() *OsImageProfile_FieldSubPath {
	return &OsImageProfile_FieldSubPath{
		selector: OsImageProfile_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Lifecycle().BlockDeletion().FieldPath(),
	}
}

func (s OsImageProfilePathSelectorMetadataLifecycleBlockDeletion) WithValue(value bool) *OsImageProfile_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*OsImageProfile_FieldSubPathValue)
}

func (s OsImageProfilePathSelectorMetadataLifecycleBlockDeletion) WithArrayOfValues(values []bool) *OsImageProfile_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*OsImageProfile_FieldSubPathArrayOfValues)
}

type OsImageProfilePathSelectorDisplayName struct{}

func (OsImageProfilePathSelectorDisplayName) FieldPath() *OsImageProfile_FieldTerminalPath {
	return &OsImageProfile_FieldTerminalPath{selector: OsImageProfile_FieldPathSelectorDisplayName}
}

func (s OsImageProfilePathSelectorDisplayName) WithValue(value string) *OsImageProfile_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*OsImageProfile_FieldTerminalPathValue)
}

func (s OsImageProfilePathSelectorDisplayName) WithArrayOfValues(values []string) *OsImageProfile_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*OsImageProfile_FieldTerminalPathArrayOfValues)
}

type OsImageProfilePathSelectorDeviceType struct{}

func (OsImageProfilePathSelectorDeviceType) FieldPath() *OsImageProfile_FieldTerminalPath {
	return &OsImageProfile_FieldTerminalPath{selector: OsImageProfile_FieldPathSelectorDeviceType}
}

func (s OsImageProfilePathSelectorDeviceType) WithValue(value *device_type.Reference) *OsImageProfile_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*OsImageProfile_FieldTerminalPathValue)
}

func (s OsImageProfilePathSelectorDeviceType) WithArrayOfValues(values []*device_type.Reference) *OsImageProfile_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*OsImageProfile_FieldTerminalPathArrayOfValues)
}

type OsImageProfilePathSelectorInstallAiAccelerator struct{}

func (OsImageProfilePathSelectorInstallAiAccelerator) FieldPath() *OsImageProfile_FieldTerminalPath {
	return &OsImageProfile_FieldTerminalPath{selector: OsImageProfile_FieldPathSelectorInstallAiAccelerator}
}

func (s OsImageProfilePathSelectorInstallAiAccelerator) WithValue(value bool) *OsImageProfile_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*OsImageProfile_FieldTerminalPathValue)
}

func (s OsImageProfilePathSelectorInstallAiAccelerator) WithArrayOfValues(values []bool) *OsImageProfile_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*OsImageProfile_FieldTerminalPathArrayOfValues)
}

type OsImageProfilePathSelectorEncryption struct{}

func (OsImageProfilePathSelectorEncryption) FieldPath() *OsImageProfile_FieldTerminalPath {
	return &OsImageProfile_FieldTerminalPath{selector: OsImageProfile_FieldPathSelectorEncryption}
}

func (s OsImageProfilePathSelectorEncryption) WithValue(value bool) *OsImageProfile_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*OsImageProfile_FieldTerminalPathValue)
}

func (s OsImageProfilePathSelectorEncryption) WithArrayOfValues(values []bool) *OsImageProfile_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*OsImageProfile_FieldTerminalPathArrayOfValues)
}

type OsImageProfilePathSelectorDiskMapping struct{}

func (OsImageProfilePathSelectorDiskMapping) FieldPath() *OsImageProfile_FieldTerminalPath {
	return &OsImageProfile_FieldTerminalPath{selector: OsImageProfile_FieldPathSelectorDiskMapping}
}

func (s OsImageProfilePathSelectorDiskMapping) WithValue(value string) *OsImageProfile_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*OsImageProfile_FieldTerminalPathValue)
}

func (s OsImageProfilePathSelectorDiskMapping) WithArrayOfValues(values []string) *OsImageProfile_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*OsImageProfile_FieldTerminalPathArrayOfValues)
}

type OsImageProfilePathSelectorNetworkAgent struct{}

func (OsImageProfilePathSelectorNetworkAgent) FieldPath() *OsImageProfile_FieldTerminalPath {
	return &OsImageProfile_FieldTerminalPath{selector: OsImageProfile_FieldPathSelectorNetworkAgent}
}

func (s OsImageProfilePathSelectorNetworkAgent) WithValue(value string) *OsImageProfile_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*OsImageProfile_FieldTerminalPathValue)
}

func (s OsImageProfilePathSelectorNetworkAgent) WithArrayOfValues(values []string) *OsImageProfile_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*OsImageProfile_FieldTerminalPathArrayOfValues)
}

type OsImageProfilePathSelectorNtp struct{}

func (OsImageProfilePathSelectorNtp) FieldPath() *OsImageProfile_FieldTerminalPath {
	return &OsImageProfile_FieldTerminalPath{selector: OsImageProfile_FieldPathSelectorNtp}
}

func (s OsImageProfilePathSelectorNtp) WithValue(value string) *OsImageProfile_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*OsImageProfile_FieldTerminalPathValue)
}

func (s OsImageProfilePathSelectorNtp) WithArrayOfValues(values []string) *OsImageProfile_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*OsImageProfile_FieldTerminalPathArrayOfValues)
}

type OsImageProfilePathSelectorHttpProxy struct{}

func (OsImageProfilePathSelectorHttpProxy) FieldPath() *OsImageProfile_FieldTerminalPath {
	return &OsImageProfile_FieldTerminalPath{selector: OsImageProfile_FieldPathSelectorHttpProxy}
}

func (s OsImageProfilePathSelectorHttpProxy) WithValue(value string) *OsImageProfile_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*OsImageProfile_FieldTerminalPathValue)
}

func (s OsImageProfilePathSelectorHttpProxy) WithArrayOfValues(values []string) *OsImageProfile_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*OsImageProfile_FieldTerminalPathArrayOfValues)
}

type OsImageProfilePathSelectorHttpsProxy struct{}

func (OsImageProfilePathSelectorHttpsProxy) FieldPath() *OsImageProfile_FieldTerminalPath {
	return &OsImageProfile_FieldTerminalPath{selector: OsImageProfile_FieldPathSelectorHttpsProxy}
}

func (s OsImageProfilePathSelectorHttpsProxy) WithValue(value string) *OsImageProfile_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*OsImageProfile_FieldTerminalPathValue)
}

func (s OsImageProfilePathSelectorHttpsProxy) WithArrayOfValues(values []string) *OsImageProfile_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*OsImageProfile_FieldTerminalPathArrayOfValues)
}

type OsImageProfilePathSelectorNoProxy struct{}

func (OsImageProfilePathSelectorNoProxy) FieldPath() *OsImageProfile_FieldTerminalPath {
	return &OsImageProfile_FieldTerminalPath{selector: OsImageProfile_FieldPathSelectorNoProxy}
}

func (s OsImageProfilePathSelectorNoProxy) WithValue(value string) *OsImageProfile_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*OsImageProfile_FieldTerminalPathValue)
}

func (s OsImageProfilePathSelectorNoProxy) WithArrayOfValues(values []string) *OsImageProfile_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*OsImageProfile_FieldTerminalPathArrayOfValues)
}