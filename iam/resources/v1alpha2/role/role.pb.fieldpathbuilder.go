// Code generated by protoc-gen-goten-object
// File: edgelq/iam/proto/v1alpha2/role.proto
// DO NOT EDIT!!!

package role

// proto imports
import (
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
	condition "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/condition"
	organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/organization"
	permission "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/permission"
	project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/project"
	policy "github.com/cloudwan/edgelq-sdk/meta/multi_region/proto/policy"
	syncing_meta "github.com/cloudwan/edgelq-sdk/meta/multi_region/proto/syncing_meta"
	meta_service "github.com/cloudwan/edgelq-sdk/meta/resources/v1alpha2/service"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
)

// make sure we're using proto imports
var (
	_ = &ntt_meta.Meta{}
	_ = &condition.Condition{}
	_ = &organization.Organization{}
	_ = &permission.Permission{}
	_ = &project.Project{}
	_ = &policy.Policy{}
	_ = &syncing_meta.SyncingMeta{}
	_ = &meta_service.Service{}
	_ = &timestamp.Timestamp{}
)

type RoleFieldPathBuilder struct{}

func NewRoleFieldPathBuilder() RoleFieldPathBuilder {
	return RoleFieldPathBuilder{}
}
func (RoleFieldPathBuilder) Name() RolePathSelectorName {
	return RolePathSelectorName{}
}
func (RoleFieldPathBuilder) DisplayName() RolePathSelectorDisplayName {
	return RolePathSelectorDisplayName{}
}
func (RoleFieldPathBuilder) IncludedPermissions() RolePathSelectorIncludedPermissions {
	return RolePathSelectorIncludedPermissions{}
}
func (RoleFieldPathBuilder) DefaultConditionBinding() RolePathSelectorDefaultConditionBinding {
	return RolePathSelectorDefaultConditionBinding{}
}
func (RoleFieldPathBuilder) Metadata() RolePathSelectorMetadata {
	return RolePathSelectorMetadata{}
}

type RolePathSelectorName struct{}

func (RolePathSelectorName) FieldPath() *Role_FieldTerminalPath {
	return &Role_FieldTerminalPath{selector: Role_FieldPathSelectorName}
}

func (s RolePathSelectorName) WithValue(value *Name) *Role_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*Role_FieldTerminalPathValue)
}

func (s RolePathSelectorName) WithArrayOfValues(values []*Name) *Role_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Role_FieldTerminalPathArrayOfValues)
}

type RolePathSelectorDisplayName struct{}

func (RolePathSelectorDisplayName) FieldPath() *Role_FieldTerminalPath {
	return &Role_FieldTerminalPath{selector: Role_FieldPathSelectorDisplayName}
}

func (s RolePathSelectorDisplayName) WithValue(value string) *Role_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*Role_FieldTerminalPathValue)
}

func (s RolePathSelectorDisplayName) WithArrayOfValues(values []string) *Role_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Role_FieldTerminalPathArrayOfValues)
}

type RolePathSelectorIncludedPermissions struct{}

func (RolePathSelectorIncludedPermissions) FieldPath() *Role_FieldTerminalPath {
	return &Role_FieldTerminalPath{selector: Role_FieldPathSelectorIncludedPermissions}
}

func (s RolePathSelectorIncludedPermissions) WithValue(value []*permission.Reference) *Role_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*Role_FieldTerminalPathValue)
}

func (s RolePathSelectorIncludedPermissions) WithArrayOfValues(values [][]*permission.Reference) *Role_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Role_FieldTerminalPathArrayOfValues)
}

func (s RolePathSelectorIncludedPermissions) WithItemValue(value *permission.Reference) *Role_FieldTerminalPathArrayItemValue {
	return s.FieldPath().WithIArrayItemValue(value).(*Role_FieldTerminalPathArrayItemValue)
}

type RolePathSelectorDefaultConditionBinding struct{}

func (RolePathSelectorDefaultConditionBinding) FieldPath() *Role_FieldTerminalPath {
	return &Role_FieldTerminalPath{selector: Role_FieldPathSelectorDefaultConditionBinding}
}

func (s RolePathSelectorDefaultConditionBinding) WithValue(value *condition.ConditionBinding) *Role_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*Role_FieldTerminalPathValue)
}

func (s RolePathSelectorDefaultConditionBinding) WithArrayOfValues(values []*condition.ConditionBinding) *Role_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Role_FieldTerminalPathArrayOfValues)
}

func (RolePathSelectorDefaultConditionBinding) WithSubPath(subPath condition.ConditionBinding_FieldPath) *Role_FieldSubPath {
	return &Role_FieldSubPath{selector: Role_FieldPathSelectorDefaultConditionBinding, subPath: subPath}
}

func (s RolePathSelectorDefaultConditionBinding) WithSubValue(subPathValue condition.ConditionBinding_FieldPathValue) *Role_FieldSubPathValue {
	return &Role_FieldSubPathValue{Role_FieldPath: s.WithSubPath(subPathValue), subPathValue: subPathValue}
}

func (s RolePathSelectorDefaultConditionBinding) WithSubArrayOfValues(subPathArrayOfValues condition.ConditionBinding_FieldPathArrayOfValues) *Role_FieldSubPathArrayOfValues {
	return &Role_FieldSubPathArrayOfValues{Role_FieldPath: s.WithSubPath(subPathArrayOfValues), subPathArrayOfValues: subPathArrayOfValues}
}

func (s RolePathSelectorDefaultConditionBinding) WithSubArrayItemValue(subPathArrayItemValue condition.ConditionBinding_FieldPathArrayItemValue) *Role_FieldSubPathArrayItemValue {
	return &Role_FieldSubPathArrayItemValue{Role_FieldPath: s.WithSubPath(subPathArrayItemValue), subPathItemValue: subPathArrayItemValue}
}

func (RolePathSelectorDefaultConditionBinding) Condition() RolePathSelectorDefaultConditionBindingCondition {
	return RolePathSelectorDefaultConditionBindingCondition{}
}

func (RolePathSelectorDefaultConditionBinding) Parameters() RolePathSelectorDefaultConditionBindingParameters {
	return RolePathSelectorDefaultConditionBindingParameters{}
}

type RolePathSelectorDefaultConditionBindingCondition struct{}

func (RolePathSelectorDefaultConditionBindingCondition) FieldPath() *Role_FieldSubPath {
	return &Role_FieldSubPath{
		selector: Role_FieldPathSelectorDefaultConditionBinding,
		subPath:  condition.NewConditionBindingFieldPathBuilder().Condition().FieldPath(),
	}
}

func (s RolePathSelectorDefaultConditionBindingCondition) WithValue(value *condition.Reference) *Role_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Role_FieldSubPathValue)
}

func (s RolePathSelectorDefaultConditionBindingCondition) WithArrayOfValues(values []*condition.Reference) *Role_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Role_FieldSubPathArrayOfValues)
}

type RolePathSelectorDefaultConditionBindingParameters struct{}

func (RolePathSelectorDefaultConditionBindingParameters) FieldPath() *Role_FieldSubPath {
	return &Role_FieldSubPath{
		selector: Role_FieldPathSelectorDefaultConditionBinding,
		subPath:  condition.NewConditionBindingFieldPathBuilder().Parameters().FieldPath(),
	}
}

func (s RolePathSelectorDefaultConditionBindingParameters) WithValue(value map[string]string) *Role_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Role_FieldSubPathValue)
}

func (s RolePathSelectorDefaultConditionBindingParameters) WithArrayOfValues(values []map[string]string) *Role_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Role_FieldSubPathArrayOfValues)
}

func (RolePathSelectorDefaultConditionBindingParameters) WithKey(key string) RoleMapPathSelectorDefaultConditionBindingParameters {
	return RoleMapPathSelectorDefaultConditionBindingParameters{key: key}
}

type RoleMapPathSelectorDefaultConditionBindingParameters struct {
	key string
}

func (s RoleMapPathSelectorDefaultConditionBindingParameters) FieldPath() *Role_FieldSubPath {
	return &Role_FieldSubPath{
		selector: Role_FieldPathSelectorDefaultConditionBinding,
		subPath:  condition.NewConditionBindingFieldPathBuilder().Parameters().WithKey(s.key).FieldPath(),
	}
}

func (s RoleMapPathSelectorDefaultConditionBindingParameters) WithValue(value string) *Role_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Role_FieldSubPathValue)
}

func (s RoleMapPathSelectorDefaultConditionBindingParameters) WithArrayOfValues(values []string) *Role_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Role_FieldSubPathArrayOfValues)
}

type RolePathSelectorMetadata struct{}

func (RolePathSelectorMetadata) FieldPath() *Role_FieldTerminalPath {
	return &Role_FieldTerminalPath{selector: Role_FieldPathSelectorMetadata}
}

func (s RolePathSelectorMetadata) WithValue(value *ntt_meta.Meta) *Role_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*Role_FieldTerminalPathValue)
}

func (s RolePathSelectorMetadata) WithArrayOfValues(values []*ntt_meta.Meta) *Role_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Role_FieldTerminalPathArrayOfValues)
}

func (RolePathSelectorMetadata) WithSubPath(subPath ntt_meta.Meta_FieldPath) *Role_FieldSubPath {
	return &Role_FieldSubPath{selector: Role_FieldPathSelectorMetadata, subPath: subPath}
}

func (s RolePathSelectorMetadata) WithSubValue(subPathValue ntt_meta.Meta_FieldPathValue) *Role_FieldSubPathValue {
	return &Role_FieldSubPathValue{Role_FieldPath: s.WithSubPath(subPathValue), subPathValue: subPathValue}
}

func (s RolePathSelectorMetadata) WithSubArrayOfValues(subPathArrayOfValues ntt_meta.Meta_FieldPathArrayOfValues) *Role_FieldSubPathArrayOfValues {
	return &Role_FieldSubPathArrayOfValues{Role_FieldPath: s.WithSubPath(subPathArrayOfValues), subPathArrayOfValues: subPathArrayOfValues}
}

func (s RolePathSelectorMetadata) WithSubArrayItemValue(subPathArrayItemValue ntt_meta.Meta_FieldPathArrayItemValue) *Role_FieldSubPathArrayItemValue {
	return &Role_FieldSubPathArrayItemValue{Role_FieldPath: s.WithSubPath(subPathArrayItemValue), subPathItemValue: subPathArrayItemValue}
}

func (RolePathSelectorMetadata) CreateTime() RolePathSelectorMetadataCreateTime {
	return RolePathSelectorMetadataCreateTime{}
}

func (RolePathSelectorMetadata) UpdateTime() RolePathSelectorMetadataUpdateTime {
	return RolePathSelectorMetadataUpdateTime{}
}

func (RolePathSelectorMetadata) Uuid() RolePathSelectorMetadataUuid {
	return RolePathSelectorMetadataUuid{}
}

func (RolePathSelectorMetadata) Tags() RolePathSelectorMetadataTags {
	return RolePathSelectorMetadataTags{}
}

func (RolePathSelectorMetadata) Labels() RolePathSelectorMetadataLabels {
	return RolePathSelectorMetadataLabels{}
}

func (RolePathSelectorMetadata) Annotations() RolePathSelectorMetadataAnnotations {
	return RolePathSelectorMetadataAnnotations{}
}

func (RolePathSelectorMetadata) Generation() RolePathSelectorMetadataGeneration {
	return RolePathSelectorMetadataGeneration{}
}

func (RolePathSelectorMetadata) ResourceVersion() RolePathSelectorMetadataResourceVersion {
	return RolePathSelectorMetadataResourceVersion{}
}

func (RolePathSelectorMetadata) OwnerReferences() RolePathSelectorMetadataOwnerReferences {
	return RolePathSelectorMetadataOwnerReferences{}
}

func (RolePathSelectorMetadata) Shards() RolePathSelectorMetadataShards {
	return RolePathSelectorMetadataShards{}
}

func (RolePathSelectorMetadata) Syncing() RolePathSelectorMetadataSyncing {
	return RolePathSelectorMetadataSyncing{}
}

type RolePathSelectorMetadataCreateTime struct{}

func (RolePathSelectorMetadataCreateTime) FieldPath() *Role_FieldSubPath {
	return &Role_FieldSubPath{
		selector: Role_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().CreateTime().FieldPath(),
	}
}

func (s RolePathSelectorMetadataCreateTime) WithValue(value *timestamp.Timestamp) *Role_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Role_FieldSubPathValue)
}

func (s RolePathSelectorMetadataCreateTime) WithArrayOfValues(values []*timestamp.Timestamp) *Role_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Role_FieldSubPathArrayOfValues)
}

type RolePathSelectorMetadataUpdateTime struct{}

func (RolePathSelectorMetadataUpdateTime) FieldPath() *Role_FieldSubPath {
	return &Role_FieldSubPath{
		selector: Role_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().UpdateTime().FieldPath(),
	}
}

func (s RolePathSelectorMetadataUpdateTime) WithValue(value *timestamp.Timestamp) *Role_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Role_FieldSubPathValue)
}

func (s RolePathSelectorMetadataUpdateTime) WithArrayOfValues(values []*timestamp.Timestamp) *Role_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Role_FieldSubPathArrayOfValues)
}

type RolePathSelectorMetadataUuid struct{}

func (RolePathSelectorMetadataUuid) FieldPath() *Role_FieldSubPath {
	return &Role_FieldSubPath{
		selector: Role_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Uuid().FieldPath(),
	}
}

func (s RolePathSelectorMetadataUuid) WithValue(value string) *Role_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Role_FieldSubPathValue)
}

func (s RolePathSelectorMetadataUuid) WithArrayOfValues(values []string) *Role_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Role_FieldSubPathArrayOfValues)
}

type RolePathSelectorMetadataTags struct{}

func (RolePathSelectorMetadataTags) FieldPath() *Role_FieldSubPath {
	return &Role_FieldSubPath{
		selector: Role_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Tags().FieldPath(),
	}
}

func (s RolePathSelectorMetadataTags) WithValue(value []string) *Role_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Role_FieldSubPathValue)
}

func (s RolePathSelectorMetadataTags) WithArrayOfValues(values [][]string) *Role_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Role_FieldSubPathArrayOfValues)
}

func (s RolePathSelectorMetadataTags) WithItemValue(value string) *Role_FieldSubPathArrayItemValue {
	return s.FieldPath().WithIArrayItemValue(value).(*Role_FieldSubPathArrayItemValue)
}

type RolePathSelectorMetadataLabels struct{}

func (RolePathSelectorMetadataLabels) FieldPath() *Role_FieldSubPath {
	return &Role_FieldSubPath{
		selector: Role_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Labels().FieldPath(),
	}
}

func (s RolePathSelectorMetadataLabels) WithValue(value map[string]string) *Role_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Role_FieldSubPathValue)
}

func (s RolePathSelectorMetadataLabels) WithArrayOfValues(values []map[string]string) *Role_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Role_FieldSubPathArrayOfValues)
}

func (RolePathSelectorMetadataLabels) WithKey(key string) RoleMapPathSelectorMetadataLabels {
	return RoleMapPathSelectorMetadataLabels{key: key}
}

type RoleMapPathSelectorMetadataLabels struct {
	key string
}

func (s RoleMapPathSelectorMetadataLabels) FieldPath() *Role_FieldSubPath {
	return &Role_FieldSubPath{
		selector: Role_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Labels().WithKey(s.key).FieldPath(),
	}
}

func (s RoleMapPathSelectorMetadataLabels) WithValue(value string) *Role_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Role_FieldSubPathValue)
}

func (s RoleMapPathSelectorMetadataLabels) WithArrayOfValues(values []string) *Role_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Role_FieldSubPathArrayOfValues)
}

type RolePathSelectorMetadataAnnotations struct{}

func (RolePathSelectorMetadataAnnotations) FieldPath() *Role_FieldSubPath {
	return &Role_FieldSubPath{
		selector: Role_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Annotations().FieldPath(),
	}
}

func (s RolePathSelectorMetadataAnnotations) WithValue(value map[string]string) *Role_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Role_FieldSubPathValue)
}

func (s RolePathSelectorMetadataAnnotations) WithArrayOfValues(values []map[string]string) *Role_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Role_FieldSubPathArrayOfValues)
}

func (RolePathSelectorMetadataAnnotations) WithKey(key string) RoleMapPathSelectorMetadataAnnotations {
	return RoleMapPathSelectorMetadataAnnotations{key: key}
}

type RoleMapPathSelectorMetadataAnnotations struct {
	key string
}

func (s RoleMapPathSelectorMetadataAnnotations) FieldPath() *Role_FieldSubPath {
	return &Role_FieldSubPath{
		selector: Role_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Annotations().WithKey(s.key).FieldPath(),
	}
}

func (s RoleMapPathSelectorMetadataAnnotations) WithValue(value string) *Role_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Role_FieldSubPathValue)
}

func (s RoleMapPathSelectorMetadataAnnotations) WithArrayOfValues(values []string) *Role_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Role_FieldSubPathArrayOfValues)
}

type RolePathSelectorMetadataGeneration struct{}

func (RolePathSelectorMetadataGeneration) FieldPath() *Role_FieldSubPath {
	return &Role_FieldSubPath{
		selector: Role_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Generation().FieldPath(),
	}
}

func (s RolePathSelectorMetadataGeneration) WithValue(value int64) *Role_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Role_FieldSubPathValue)
}

func (s RolePathSelectorMetadataGeneration) WithArrayOfValues(values []int64) *Role_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Role_FieldSubPathArrayOfValues)
}

type RolePathSelectorMetadataResourceVersion struct{}

func (RolePathSelectorMetadataResourceVersion) FieldPath() *Role_FieldSubPath {
	return &Role_FieldSubPath{
		selector: Role_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().ResourceVersion().FieldPath(),
	}
}

func (s RolePathSelectorMetadataResourceVersion) WithValue(value string) *Role_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Role_FieldSubPathValue)
}

func (s RolePathSelectorMetadataResourceVersion) WithArrayOfValues(values []string) *Role_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Role_FieldSubPathArrayOfValues)
}

type RolePathSelectorMetadataOwnerReferences struct{}

func (RolePathSelectorMetadataOwnerReferences) FieldPath() *Role_FieldSubPath {
	return &Role_FieldSubPath{
		selector: Role_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().FieldPath(),
	}
}

func (s RolePathSelectorMetadataOwnerReferences) WithValue(value []*ntt_meta.OwnerReference) *Role_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Role_FieldSubPathValue)
}

func (s RolePathSelectorMetadataOwnerReferences) WithArrayOfValues(values [][]*ntt_meta.OwnerReference) *Role_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Role_FieldSubPathArrayOfValues)
}

func (s RolePathSelectorMetadataOwnerReferences) WithItemValue(value *ntt_meta.OwnerReference) *Role_FieldSubPathArrayItemValue {
	return s.FieldPath().WithIArrayItemValue(value).(*Role_FieldSubPathArrayItemValue)
}

func (RolePathSelectorMetadataOwnerReferences) ApiVersion() RolePathSelectorMetadataOwnerReferencesApiVersion {
	return RolePathSelectorMetadataOwnerReferencesApiVersion{}
}

func (RolePathSelectorMetadataOwnerReferences) Kind() RolePathSelectorMetadataOwnerReferencesKind {
	return RolePathSelectorMetadataOwnerReferencesKind{}
}

func (RolePathSelectorMetadataOwnerReferences) Name() RolePathSelectorMetadataOwnerReferencesName {
	return RolePathSelectorMetadataOwnerReferencesName{}
}

func (RolePathSelectorMetadataOwnerReferences) Uid() RolePathSelectorMetadataOwnerReferencesUid {
	return RolePathSelectorMetadataOwnerReferencesUid{}
}

func (RolePathSelectorMetadataOwnerReferences) Controller() RolePathSelectorMetadataOwnerReferencesController {
	return RolePathSelectorMetadataOwnerReferencesController{}
}

func (RolePathSelectorMetadataOwnerReferences) BlockOwnerDeletion() RolePathSelectorMetadataOwnerReferencesBlockOwnerDeletion {
	return RolePathSelectorMetadataOwnerReferencesBlockOwnerDeletion{}
}

type RolePathSelectorMetadataOwnerReferencesApiVersion struct{}

func (RolePathSelectorMetadataOwnerReferencesApiVersion) FieldPath() *Role_FieldSubPath {
	return &Role_FieldSubPath{
		selector: Role_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().ApiVersion().FieldPath(),
	}
}

func (s RolePathSelectorMetadataOwnerReferencesApiVersion) WithValue(value string) *Role_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Role_FieldSubPathValue)
}

func (s RolePathSelectorMetadataOwnerReferencesApiVersion) WithArrayOfValues(values []string) *Role_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Role_FieldSubPathArrayOfValues)
}

type RolePathSelectorMetadataOwnerReferencesKind struct{}

func (RolePathSelectorMetadataOwnerReferencesKind) FieldPath() *Role_FieldSubPath {
	return &Role_FieldSubPath{
		selector: Role_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().Kind().FieldPath(),
	}
}

func (s RolePathSelectorMetadataOwnerReferencesKind) WithValue(value string) *Role_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Role_FieldSubPathValue)
}

func (s RolePathSelectorMetadataOwnerReferencesKind) WithArrayOfValues(values []string) *Role_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Role_FieldSubPathArrayOfValues)
}

type RolePathSelectorMetadataOwnerReferencesName struct{}

func (RolePathSelectorMetadataOwnerReferencesName) FieldPath() *Role_FieldSubPath {
	return &Role_FieldSubPath{
		selector: Role_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().Name().FieldPath(),
	}
}

func (s RolePathSelectorMetadataOwnerReferencesName) WithValue(value string) *Role_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Role_FieldSubPathValue)
}

func (s RolePathSelectorMetadataOwnerReferencesName) WithArrayOfValues(values []string) *Role_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Role_FieldSubPathArrayOfValues)
}

type RolePathSelectorMetadataOwnerReferencesUid struct{}

func (RolePathSelectorMetadataOwnerReferencesUid) FieldPath() *Role_FieldSubPath {
	return &Role_FieldSubPath{
		selector: Role_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().Uid().FieldPath(),
	}
}

func (s RolePathSelectorMetadataOwnerReferencesUid) WithValue(value string) *Role_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Role_FieldSubPathValue)
}

func (s RolePathSelectorMetadataOwnerReferencesUid) WithArrayOfValues(values []string) *Role_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Role_FieldSubPathArrayOfValues)
}

type RolePathSelectorMetadataOwnerReferencesController struct{}

func (RolePathSelectorMetadataOwnerReferencesController) FieldPath() *Role_FieldSubPath {
	return &Role_FieldSubPath{
		selector: Role_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().Controller().FieldPath(),
	}
}

func (s RolePathSelectorMetadataOwnerReferencesController) WithValue(value bool) *Role_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Role_FieldSubPathValue)
}

func (s RolePathSelectorMetadataOwnerReferencesController) WithArrayOfValues(values []bool) *Role_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Role_FieldSubPathArrayOfValues)
}

type RolePathSelectorMetadataOwnerReferencesBlockOwnerDeletion struct{}

func (RolePathSelectorMetadataOwnerReferencesBlockOwnerDeletion) FieldPath() *Role_FieldSubPath {
	return &Role_FieldSubPath{
		selector: Role_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().BlockOwnerDeletion().FieldPath(),
	}
}

func (s RolePathSelectorMetadataOwnerReferencesBlockOwnerDeletion) WithValue(value bool) *Role_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Role_FieldSubPathValue)
}

func (s RolePathSelectorMetadataOwnerReferencesBlockOwnerDeletion) WithArrayOfValues(values []bool) *Role_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Role_FieldSubPathArrayOfValues)
}

type RolePathSelectorMetadataShards struct{}

func (RolePathSelectorMetadataShards) FieldPath() *Role_FieldSubPath {
	return &Role_FieldSubPath{
		selector: Role_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Shards().FieldPath(),
	}
}

func (s RolePathSelectorMetadataShards) WithValue(value map[string]int64) *Role_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Role_FieldSubPathValue)
}

func (s RolePathSelectorMetadataShards) WithArrayOfValues(values []map[string]int64) *Role_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Role_FieldSubPathArrayOfValues)
}

func (RolePathSelectorMetadataShards) WithKey(key string) RoleMapPathSelectorMetadataShards {
	return RoleMapPathSelectorMetadataShards{key: key}
}

type RoleMapPathSelectorMetadataShards struct {
	key string
}

func (s RoleMapPathSelectorMetadataShards) FieldPath() *Role_FieldSubPath {
	return &Role_FieldSubPath{
		selector: Role_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Shards().WithKey(s.key).FieldPath(),
	}
}

func (s RoleMapPathSelectorMetadataShards) WithValue(value int64) *Role_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Role_FieldSubPathValue)
}

func (s RoleMapPathSelectorMetadataShards) WithArrayOfValues(values []int64) *Role_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Role_FieldSubPathArrayOfValues)
}

type RolePathSelectorMetadataSyncing struct{}

func (RolePathSelectorMetadataSyncing) FieldPath() *Role_FieldSubPath {
	return &Role_FieldSubPath{
		selector: Role_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Syncing().FieldPath(),
	}
}

func (s RolePathSelectorMetadataSyncing) WithValue(value *syncing_meta.SyncingMeta) *Role_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Role_FieldSubPathValue)
}

func (s RolePathSelectorMetadataSyncing) WithArrayOfValues(values []*syncing_meta.SyncingMeta) *Role_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Role_FieldSubPathArrayOfValues)
}

func (RolePathSelectorMetadataSyncing) OwningRegion() RolePathSelectorMetadataSyncingOwningRegion {
	return RolePathSelectorMetadataSyncingOwningRegion{}
}

func (RolePathSelectorMetadataSyncing) Regions() RolePathSelectorMetadataSyncingRegions {
	return RolePathSelectorMetadataSyncingRegions{}
}

type RolePathSelectorMetadataSyncingOwningRegion struct{}

func (RolePathSelectorMetadataSyncingOwningRegion) FieldPath() *Role_FieldSubPath {
	return &Role_FieldSubPath{
		selector: Role_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Syncing().OwningRegion().FieldPath(),
	}
}

func (s RolePathSelectorMetadataSyncingOwningRegion) WithValue(value string) *Role_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Role_FieldSubPathValue)
}

func (s RolePathSelectorMetadataSyncingOwningRegion) WithArrayOfValues(values []string) *Role_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Role_FieldSubPathArrayOfValues)
}

type RolePathSelectorMetadataSyncingRegions struct{}

func (RolePathSelectorMetadataSyncingRegions) FieldPath() *Role_FieldSubPath {
	return &Role_FieldSubPath{
		selector: Role_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Syncing().Regions().FieldPath(),
	}
}

func (s RolePathSelectorMetadataSyncingRegions) WithValue(value []string) *Role_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*Role_FieldSubPathValue)
}

func (s RolePathSelectorMetadataSyncingRegions) WithArrayOfValues(values [][]string) *Role_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Role_FieldSubPathArrayOfValues)
}

func (s RolePathSelectorMetadataSyncingRegions) WithItemValue(value string) *Role_FieldSubPathArrayItemValue {
	return s.FieldPath().WithIArrayItemValue(value).(*Role_FieldSubPathArrayItemValue)
}
