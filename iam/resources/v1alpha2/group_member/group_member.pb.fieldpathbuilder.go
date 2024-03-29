// Code generated by protoc-gen-goten-object
// File: edgelq/iam/proto/v1alpha2/group_member.proto
// DO NOT EDIT!!!

package group_member

// proto imports
import (
	iam_common "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/common"
	group "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/group"
	organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/organization"
	project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/project"
	meta_service "github.com/cloudwan/edgelq-sdk/meta/resources/v1alpha2/service"
	meta "github.com/cloudwan/goten-sdk/types/meta"
	multi_region_policy "github.com/cloudwan/goten-sdk/types/multi_region_policy"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

// make sure we're using proto imports
var (
	_ = &iam_common.PCR{}
	_ = &group.Group{}
	_ = &organization.Organization{}
	_ = &project.Project{}
	_ = &meta_service.Service{}
	_ = &timestamppb.Timestamp{}
	_ = &meta.Meta{}
	_ = &multi_region_policy.MultiRegionPolicy{}
)

type GroupMemberFieldPathBuilder struct{}

func NewGroupMemberFieldPathBuilder() GroupMemberFieldPathBuilder {
	return GroupMemberFieldPathBuilder{}
}
func (GroupMemberFieldPathBuilder) Name() GroupMemberPathSelectorName {
	return GroupMemberPathSelectorName{}
}
func (GroupMemberFieldPathBuilder) Member() GroupMemberPathSelectorMember {
	return GroupMemberPathSelectorMember{}
}
func (GroupMemberFieldPathBuilder) ParentMember() GroupMemberPathSelectorParentMember {
	return GroupMemberPathSelectorParentMember{}
}
func (GroupMemberFieldPathBuilder) MinAncestryMembers() GroupMemberPathSelectorMinAncestryMembers {
	return GroupMemberPathSelectorMinAncestryMembers{}
}
func (GroupMemberFieldPathBuilder) Metadata() GroupMemberPathSelectorMetadata {
	return GroupMemberPathSelectorMetadata{}
}

type GroupMemberPathSelectorName struct{}

func (GroupMemberPathSelectorName) FieldPath() *GroupMember_FieldTerminalPath {
	return &GroupMember_FieldTerminalPath{selector: GroupMember_FieldPathSelectorName}
}

func (s GroupMemberPathSelectorName) WithValue(value *Name) *GroupMember_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*GroupMember_FieldTerminalPathValue)
}

func (s GroupMemberPathSelectorName) WithArrayOfValues(values []*Name) *GroupMember_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*GroupMember_FieldTerminalPathArrayOfValues)
}

type GroupMemberPathSelectorMember struct{}

func (GroupMemberPathSelectorMember) FieldPath() *GroupMember_FieldTerminalPath {
	return &GroupMember_FieldTerminalPath{selector: GroupMember_FieldPathSelectorMember}
}

func (s GroupMemberPathSelectorMember) WithValue(value string) *GroupMember_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*GroupMember_FieldTerminalPathValue)
}

func (s GroupMemberPathSelectorMember) WithArrayOfValues(values []string) *GroupMember_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*GroupMember_FieldTerminalPathArrayOfValues)
}

type GroupMemberPathSelectorParentMember struct{}

func (GroupMemberPathSelectorParentMember) FieldPath() *GroupMember_FieldTerminalPath {
	return &GroupMember_FieldTerminalPath{selector: GroupMember_FieldPathSelectorParentMember}
}

func (s GroupMemberPathSelectorParentMember) WithValue(value string) *GroupMember_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*GroupMember_FieldTerminalPathValue)
}

func (s GroupMemberPathSelectorParentMember) WithArrayOfValues(values []string) *GroupMember_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*GroupMember_FieldTerminalPathArrayOfValues)
}

type GroupMemberPathSelectorMinAncestryMembers struct{}

func (GroupMemberPathSelectorMinAncestryMembers) FieldPath() *GroupMember_FieldTerminalPath {
	return &GroupMember_FieldTerminalPath{selector: GroupMember_FieldPathSelectorMinAncestryMembers}
}

func (s GroupMemberPathSelectorMinAncestryMembers) WithValue(value []string) *GroupMember_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*GroupMember_FieldTerminalPathValue)
}

func (s GroupMemberPathSelectorMinAncestryMembers) WithArrayOfValues(values [][]string) *GroupMember_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*GroupMember_FieldTerminalPathArrayOfValues)
}

func (s GroupMemberPathSelectorMinAncestryMembers) WithItemValue(value string) *GroupMember_FieldTerminalPathArrayItemValue {
	return s.FieldPath().WithIArrayItemValue(value).(*GroupMember_FieldTerminalPathArrayItemValue)
}

type GroupMemberPathSelectorMetadata struct{}

func (GroupMemberPathSelectorMetadata) FieldPath() *GroupMember_FieldTerminalPath {
	return &GroupMember_FieldTerminalPath{selector: GroupMember_FieldPathSelectorMetadata}
}

func (s GroupMemberPathSelectorMetadata) WithValue(value *meta.Meta) *GroupMember_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*GroupMember_FieldTerminalPathValue)
}

func (s GroupMemberPathSelectorMetadata) WithArrayOfValues(values []*meta.Meta) *GroupMember_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*GroupMember_FieldTerminalPathArrayOfValues)
}

func (GroupMemberPathSelectorMetadata) WithSubPath(subPath meta.Meta_FieldPath) *GroupMember_FieldSubPath {
	return &GroupMember_FieldSubPath{selector: GroupMember_FieldPathSelectorMetadata, subPath: subPath}
}

func (s GroupMemberPathSelectorMetadata) WithSubValue(subPathValue meta.Meta_FieldPathValue) *GroupMember_FieldSubPathValue {
	return &GroupMember_FieldSubPathValue{GroupMember_FieldPath: s.WithSubPath(subPathValue), subPathValue: subPathValue}
}

func (s GroupMemberPathSelectorMetadata) WithSubArrayOfValues(subPathArrayOfValues meta.Meta_FieldPathArrayOfValues) *GroupMember_FieldSubPathArrayOfValues {
	return &GroupMember_FieldSubPathArrayOfValues{GroupMember_FieldPath: s.WithSubPath(subPathArrayOfValues), subPathArrayOfValues: subPathArrayOfValues}
}

func (s GroupMemberPathSelectorMetadata) WithSubArrayItemValue(subPathArrayItemValue meta.Meta_FieldPathArrayItemValue) *GroupMember_FieldSubPathArrayItemValue {
	return &GroupMember_FieldSubPathArrayItemValue{GroupMember_FieldPath: s.WithSubPath(subPathArrayItemValue), subPathItemValue: subPathArrayItemValue}
}

func (GroupMemberPathSelectorMetadata) CreateTime() GroupMemberPathSelectorMetadataCreateTime {
	return GroupMemberPathSelectorMetadataCreateTime{}
}

func (GroupMemberPathSelectorMetadata) UpdateTime() GroupMemberPathSelectorMetadataUpdateTime {
	return GroupMemberPathSelectorMetadataUpdateTime{}
}

func (GroupMemberPathSelectorMetadata) DeleteTime() GroupMemberPathSelectorMetadataDeleteTime {
	return GroupMemberPathSelectorMetadataDeleteTime{}
}

func (GroupMemberPathSelectorMetadata) Uuid() GroupMemberPathSelectorMetadataUuid {
	return GroupMemberPathSelectorMetadataUuid{}
}

func (GroupMemberPathSelectorMetadata) Tags() GroupMemberPathSelectorMetadataTags {
	return GroupMemberPathSelectorMetadataTags{}
}

func (GroupMemberPathSelectorMetadata) Labels() GroupMemberPathSelectorMetadataLabels {
	return GroupMemberPathSelectorMetadataLabels{}
}

func (GroupMemberPathSelectorMetadata) Annotations() GroupMemberPathSelectorMetadataAnnotations {
	return GroupMemberPathSelectorMetadataAnnotations{}
}

func (GroupMemberPathSelectorMetadata) Generation() GroupMemberPathSelectorMetadataGeneration {
	return GroupMemberPathSelectorMetadataGeneration{}
}

func (GroupMemberPathSelectorMetadata) ResourceVersion() GroupMemberPathSelectorMetadataResourceVersion {
	return GroupMemberPathSelectorMetadataResourceVersion{}
}

func (GroupMemberPathSelectorMetadata) OwnerReferences() GroupMemberPathSelectorMetadataOwnerReferences {
	return GroupMemberPathSelectorMetadataOwnerReferences{}
}

func (GroupMemberPathSelectorMetadata) Shards() GroupMemberPathSelectorMetadataShards {
	return GroupMemberPathSelectorMetadataShards{}
}

func (GroupMemberPathSelectorMetadata) Syncing() GroupMemberPathSelectorMetadataSyncing {
	return GroupMemberPathSelectorMetadataSyncing{}
}

func (GroupMemberPathSelectorMetadata) Lifecycle() GroupMemberPathSelectorMetadataLifecycle {
	return GroupMemberPathSelectorMetadataLifecycle{}
}

func (GroupMemberPathSelectorMetadata) Services() GroupMemberPathSelectorMetadataServices {
	return GroupMemberPathSelectorMetadataServices{}
}

type GroupMemberPathSelectorMetadataCreateTime struct{}

func (GroupMemberPathSelectorMetadataCreateTime) FieldPath() *GroupMember_FieldSubPath {
	return &GroupMember_FieldSubPath{
		selector: GroupMember_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().CreateTime().FieldPath(),
	}
}

func (s GroupMemberPathSelectorMetadataCreateTime) WithValue(value *timestamppb.Timestamp) *GroupMember_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*GroupMember_FieldSubPathValue)
}

func (s GroupMemberPathSelectorMetadataCreateTime) WithArrayOfValues(values []*timestamppb.Timestamp) *GroupMember_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*GroupMember_FieldSubPathArrayOfValues)
}

type GroupMemberPathSelectorMetadataUpdateTime struct{}

func (GroupMemberPathSelectorMetadataUpdateTime) FieldPath() *GroupMember_FieldSubPath {
	return &GroupMember_FieldSubPath{
		selector: GroupMember_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().UpdateTime().FieldPath(),
	}
}

func (s GroupMemberPathSelectorMetadataUpdateTime) WithValue(value *timestamppb.Timestamp) *GroupMember_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*GroupMember_FieldSubPathValue)
}

func (s GroupMemberPathSelectorMetadataUpdateTime) WithArrayOfValues(values []*timestamppb.Timestamp) *GroupMember_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*GroupMember_FieldSubPathArrayOfValues)
}

type GroupMemberPathSelectorMetadataDeleteTime struct{}

func (GroupMemberPathSelectorMetadataDeleteTime) FieldPath() *GroupMember_FieldSubPath {
	return &GroupMember_FieldSubPath{
		selector: GroupMember_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().DeleteTime().FieldPath(),
	}
}

func (s GroupMemberPathSelectorMetadataDeleteTime) WithValue(value *timestamppb.Timestamp) *GroupMember_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*GroupMember_FieldSubPathValue)
}

func (s GroupMemberPathSelectorMetadataDeleteTime) WithArrayOfValues(values []*timestamppb.Timestamp) *GroupMember_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*GroupMember_FieldSubPathArrayOfValues)
}

type GroupMemberPathSelectorMetadataUuid struct{}

func (GroupMemberPathSelectorMetadataUuid) FieldPath() *GroupMember_FieldSubPath {
	return &GroupMember_FieldSubPath{
		selector: GroupMember_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Uuid().FieldPath(),
	}
}

func (s GroupMemberPathSelectorMetadataUuid) WithValue(value string) *GroupMember_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*GroupMember_FieldSubPathValue)
}

func (s GroupMemberPathSelectorMetadataUuid) WithArrayOfValues(values []string) *GroupMember_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*GroupMember_FieldSubPathArrayOfValues)
}

type GroupMemberPathSelectorMetadataTags struct{}

func (GroupMemberPathSelectorMetadataTags) FieldPath() *GroupMember_FieldSubPath {
	return &GroupMember_FieldSubPath{
		selector: GroupMember_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Tags().FieldPath(),
	}
}

func (s GroupMemberPathSelectorMetadataTags) WithValue(value []string) *GroupMember_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*GroupMember_FieldSubPathValue)
}

func (s GroupMemberPathSelectorMetadataTags) WithArrayOfValues(values [][]string) *GroupMember_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*GroupMember_FieldSubPathArrayOfValues)
}

func (s GroupMemberPathSelectorMetadataTags) WithItemValue(value string) *GroupMember_FieldSubPathArrayItemValue {
	return s.FieldPath().WithIArrayItemValue(value).(*GroupMember_FieldSubPathArrayItemValue)
}

type GroupMemberPathSelectorMetadataLabels struct{}

func (GroupMemberPathSelectorMetadataLabels) FieldPath() *GroupMember_FieldSubPath {
	return &GroupMember_FieldSubPath{
		selector: GroupMember_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Labels().FieldPath(),
	}
}

func (s GroupMemberPathSelectorMetadataLabels) WithValue(value map[string]string) *GroupMember_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*GroupMember_FieldSubPathValue)
}

func (s GroupMemberPathSelectorMetadataLabels) WithArrayOfValues(values []map[string]string) *GroupMember_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*GroupMember_FieldSubPathArrayOfValues)
}

func (GroupMemberPathSelectorMetadataLabels) WithKey(key string) GroupMemberMapPathSelectorMetadataLabels {
	return GroupMemberMapPathSelectorMetadataLabels{key: key}
}

type GroupMemberMapPathSelectorMetadataLabels struct {
	key string
}

func (s GroupMemberMapPathSelectorMetadataLabels) FieldPath() *GroupMember_FieldSubPath {
	return &GroupMember_FieldSubPath{
		selector: GroupMember_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Labels().WithKey(s.key).FieldPath(),
	}
}

func (s GroupMemberMapPathSelectorMetadataLabels) WithValue(value string) *GroupMember_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*GroupMember_FieldSubPathValue)
}

func (s GroupMemberMapPathSelectorMetadataLabels) WithArrayOfValues(values []string) *GroupMember_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*GroupMember_FieldSubPathArrayOfValues)
}

type GroupMemberPathSelectorMetadataAnnotations struct{}

func (GroupMemberPathSelectorMetadataAnnotations) FieldPath() *GroupMember_FieldSubPath {
	return &GroupMember_FieldSubPath{
		selector: GroupMember_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Annotations().FieldPath(),
	}
}

func (s GroupMemberPathSelectorMetadataAnnotations) WithValue(value map[string]string) *GroupMember_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*GroupMember_FieldSubPathValue)
}

func (s GroupMemberPathSelectorMetadataAnnotations) WithArrayOfValues(values []map[string]string) *GroupMember_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*GroupMember_FieldSubPathArrayOfValues)
}

func (GroupMemberPathSelectorMetadataAnnotations) WithKey(key string) GroupMemberMapPathSelectorMetadataAnnotations {
	return GroupMemberMapPathSelectorMetadataAnnotations{key: key}
}

type GroupMemberMapPathSelectorMetadataAnnotations struct {
	key string
}

func (s GroupMemberMapPathSelectorMetadataAnnotations) FieldPath() *GroupMember_FieldSubPath {
	return &GroupMember_FieldSubPath{
		selector: GroupMember_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Annotations().WithKey(s.key).FieldPath(),
	}
}

func (s GroupMemberMapPathSelectorMetadataAnnotations) WithValue(value string) *GroupMember_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*GroupMember_FieldSubPathValue)
}

func (s GroupMemberMapPathSelectorMetadataAnnotations) WithArrayOfValues(values []string) *GroupMember_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*GroupMember_FieldSubPathArrayOfValues)
}

type GroupMemberPathSelectorMetadataGeneration struct{}

func (GroupMemberPathSelectorMetadataGeneration) FieldPath() *GroupMember_FieldSubPath {
	return &GroupMember_FieldSubPath{
		selector: GroupMember_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Generation().FieldPath(),
	}
}

func (s GroupMemberPathSelectorMetadataGeneration) WithValue(value int64) *GroupMember_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*GroupMember_FieldSubPathValue)
}

func (s GroupMemberPathSelectorMetadataGeneration) WithArrayOfValues(values []int64) *GroupMember_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*GroupMember_FieldSubPathArrayOfValues)
}

type GroupMemberPathSelectorMetadataResourceVersion struct{}

func (GroupMemberPathSelectorMetadataResourceVersion) FieldPath() *GroupMember_FieldSubPath {
	return &GroupMember_FieldSubPath{
		selector: GroupMember_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().ResourceVersion().FieldPath(),
	}
}

func (s GroupMemberPathSelectorMetadataResourceVersion) WithValue(value string) *GroupMember_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*GroupMember_FieldSubPathValue)
}

func (s GroupMemberPathSelectorMetadataResourceVersion) WithArrayOfValues(values []string) *GroupMember_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*GroupMember_FieldSubPathArrayOfValues)
}

type GroupMemberPathSelectorMetadataOwnerReferences struct{}

func (GroupMemberPathSelectorMetadataOwnerReferences) FieldPath() *GroupMember_FieldSubPath {
	return &GroupMember_FieldSubPath{
		selector: GroupMember_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().OwnerReferences().FieldPath(),
	}
}

func (s GroupMemberPathSelectorMetadataOwnerReferences) WithValue(value []*meta.OwnerReference) *GroupMember_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*GroupMember_FieldSubPathValue)
}

func (s GroupMemberPathSelectorMetadataOwnerReferences) WithArrayOfValues(values [][]*meta.OwnerReference) *GroupMember_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*GroupMember_FieldSubPathArrayOfValues)
}

func (s GroupMemberPathSelectorMetadataOwnerReferences) WithItemValue(value *meta.OwnerReference) *GroupMember_FieldSubPathArrayItemValue {
	return s.FieldPath().WithIArrayItemValue(value).(*GroupMember_FieldSubPathArrayItemValue)
}

func (GroupMemberPathSelectorMetadataOwnerReferences) Kind() GroupMemberPathSelectorMetadataOwnerReferencesKind {
	return GroupMemberPathSelectorMetadataOwnerReferencesKind{}
}

func (GroupMemberPathSelectorMetadataOwnerReferences) Version() GroupMemberPathSelectorMetadataOwnerReferencesVersion {
	return GroupMemberPathSelectorMetadataOwnerReferencesVersion{}
}

func (GroupMemberPathSelectorMetadataOwnerReferences) Name() GroupMemberPathSelectorMetadataOwnerReferencesName {
	return GroupMemberPathSelectorMetadataOwnerReferencesName{}
}

func (GroupMemberPathSelectorMetadataOwnerReferences) Region() GroupMemberPathSelectorMetadataOwnerReferencesRegion {
	return GroupMemberPathSelectorMetadataOwnerReferencesRegion{}
}

func (GroupMemberPathSelectorMetadataOwnerReferences) Controller() GroupMemberPathSelectorMetadataOwnerReferencesController {
	return GroupMemberPathSelectorMetadataOwnerReferencesController{}
}

func (GroupMemberPathSelectorMetadataOwnerReferences) RequiresOwnerReference() GroupMemberPathSelectorMetadataOwnerReferencesRequiresOwnerReference {
	return GroupMemberPathSelectorMetadataOwnerReferencesRequiresOwnerReference{}
}

type GroupMemberPathSelectorMetadataOwnerReferencesKind struct{}

func (GroupMemberPathSelectorMetadataOwnerReferencesKind) FieldPath() *GroupMember_FieldSubPath {
	return &GroupMember_FieldSubPath{
		selector: GroupMember_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().OwnerReferences().Kind().FieldPath(),
	}
}

func (s GroupMemberPathSelectorMetadataOwnerReferencesKind) WithValue(value string) *GroupMember_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*GroupMember_FieldSubPathValue)
}

func (s GroupMemberPathSelectorMetadataOwnerReferencesKind) WithArrayOfValues(values []string) *GroupMember_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*GroupMember_FieldSubPathArrayOfValues)
}

type GroupMemberPathSelectorMetadataOwnerReferencesVersion struct{}

func (GroupMemberPathSelectorMetadataOwnerReferencesVersion) FieldPath() *GroupMember_FieldSubPath {
	return &GroupMember_FieldSubPath{
		selector: GroupMember_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().OwnerReferences().Version().FieldPath(),
	}
}

func (s GroupMemberPathSelectorMetadataOwnerReferencesVersion) WithValue(value string) *GroupMember_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*GroupMember_FieldSubPathValue)
}

func (s GroupMemberPathSelectorMetadataOwnerReferencesVersion) WithArrayOfValues(values []string) *GroupMember_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*GroupMember_FieldSubPathArrayOfValues)
}

type GroupMemberPathSelectorMetadataOwnerReferencesName struct{}

func (GroupMemberPathSelectorMetadataOwnerReferencesName) FieldPath() *GroupMember_FieldSubPath {
	return &GroupMember_FieldSubPath{
		selector: GroupMember_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().OwnerReferences().Name().FieldPath(),
	}
}

func (s GroupMemberPathSelectorMetadataOwnerReferencesName) WithValue(value string) *GroupMember_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*GroupMember_FieldSubPathValue)
}

func (s GroupMemberPathSelectorMetadataOwnerReferencesName) WithArrayOfValues(values []string) *GroupMember_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*GroupMember_FieldSubPathArrayOfValues)
}

type GroupMemberPathSelectorMetadataOwnerReferencesRegion struct{}

func (GroupMemberPathSelectorMetadataOwnerReferencesRegion) FieldPath() *GroupMember_FieldSubPath {
	return &GroupMember_FieldSubPath{
		selector: GroupMember_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().OwnerReferences().Region().FieldPath(),
	}
}

func (s GroupMemberPathSelectorMetadataOwnerReferencesRegion) WithValue(value string) *GroupMember_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*GroupMember_FieldSubPathValue)
}

func (s GroupMemberPathSelectorMetadataOwnerReferencesRegion) WithArrayOfValues(values []string) *GroupMember_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*GroupMember_FieldSubPathArrayOfValues)
}

type GroupMemberPathSelectorMetadataOwnerReferencesController struct{}

func (GroupMemberPathSelectorMetadataOwnerReferencesController) FieldPath() *GroupMember_FieldSubPath {
	return &GroupMember_FieldSubPath{
		selector: GroupMember_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().OwnerReferences().Controller().FieldPath(),
	}
}

func (s GroupMemberPathSelectorMetadataOwnerReferencesController) WithValue(value bool) *GroupMember_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*GroupMember_FieldSubPathValue)
}

func (s GroupMemberPathSelectorMetadataOwnerReferencesController) WithArrayOfValues(values []bool) *GroupMember_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*GroupMember_FieldSubPathArrayOfValues)
}

type GroupMemberPathSelectorMetadataOwnerReferencesRequiresOwnerReference struct{}

func (GroupMemberPathSelectorMetadataOwnerReferencesRequiresOwnerReference) FieldPath() *GroupMember_FieldSubPath {
	return &GroupMember_FieldSubPath{
		selector: GroupMember_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().OwnerReferences().RequiresOwnerReference().FieldPath(),
	}
}

func (s GroupMemberPathSelectorMetadataOwnerReferencesRequiresOwnerReference) WithValue(value bool) *GroupMember_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*GroupMember_FieldSubPathValue)
}

func (s GroupMemberPathSelectorMetadataOwnerReferencesRequiresOwnerReference) WithArrayOfValues(values []bool) *GroupMember_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*GroupMember_FieldSubPathArrayOfValues)
}

type GroupMemberPathSelectorMetadataShards struct{}

func (GroupMemberPathSelectorMetadataShards) FieldPath() *GroupMember_FieldSubPath {
	return &GroupMember_FieldSubPath{
		selector: GroupMember_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Shards().FieldPath(),
	}
}

func (s GroupMemberPathSelectorMetadataShards) WithValue(value map[string]int64) *GroupMember_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*GroupMember_FieldSubPathValue)
}

func (s GroupMemberPathSelectorMetadataShards) WithArrayOfValues(values []map[string]int64) *GroupMember_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*GroupMember_FieldSubPathArrayOfValues)
}

func (GroupMemberPathSelectorMetadataShards) WithKey(key string) GroupMemberMapPathSelectorMetadataShards {
	return GroupMemberMapPathSelectorMetadataShards{key: key}
}

type GroupMemberMapPathSelectorMetadataShards struct {
	key string
}

func (s GroupMemberMapPathSelectorMetadataShards) FieldPath() *GroupMember_FieldSubPath {
	return &GroupMember_FieldSubPath{
		selector: GroupMember_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Shards().WithKey(s.key).FieldPath(),
	}
}

func (s GroupMemberMapPathSelectorMetadataShards) WithValue(value int64) *GroupMember_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*GroupMember_FieldSubPathValue)
}

func (s GroupMemberMapPathSelectorMetadataShards) WithArrayOfValues(values []int64) *GroupMember_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*GroupMember_FieldSubPathArrayOfValues)
}

type GroupMemberPathSelectorMetadataSyncing struct{}

func (GroupMemberPathSelectorMetadataSyncing) FieldPath() *GroupMember_FieldSubPath {
	return &GroupMember_FieldSubPath{
		selector: GroupMember_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Syncing().FieldPath(),
	}
}

func (s GroupMemberPathSelectorMetadataSyncing) WithValue(value *meta.SyncingMeta) *GroupMember_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*GroupMember_FieldSubPathValue)
}

func (s GroupMemberPathSelectorMetadataSyncing) WithArrayOfValues(values []*meta.SyncingMeta) *GroupMember_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*GroupMember_FieldSubPathArrayOfValues)
}

func (GroupMemberPathSelectorMetadataSyncing) OwningRegion() GroupMemberPathSelectorMetadataSyncingOwningRegion {
	return GroupMemberPathSelectorMetadataSyncingOwningRegion{}
}

func (GroupMemberPathSelectorMetadataSyncing) Regions() GroupMemberPathSelectorMetadataSyncingRegions {
	return GroupMemberPathSelectorMetadataSyncingRegions{}
}

type GroupMemberPathSelectorMetadataSyncingOwningRegion struct{}

func (GroupMemberPathSelectorMetadataSyncingOwningRegion) FieldPath() *GroupMember_FieldSubPath {
	return &GroupMember_FieldSubPath{
		selector: GroupMember_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Syncing().OwningRegion().FieldPath(),
	}
}

func (s GroupMemberPathSelectorMetadataSyncingOwningRegion) WithValue(value string) *GroupMember_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*GroupMember_FieldSubPathValue)
}

func (s GroupMemberPathSelectorMetadataSyncingOwningRegion) WithArrayOfValues(values []string) *GroupMember_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*GroupMember_FieldSubPathArrayOfValues)
}

type GroupMemberPathSelectorMetadataSyncingRegions struct{}

func (GroupMemberPathSelectorMetadataSyncingRegions) FieldPath() *GroupMember_FieldSubPath {
	return &GroupMember_FieldSubPath{
		selector: GroupMember_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Syncing().Regions().FieldPath(),
	}
}

func (s GroupMemberPathSelectorMetadataSyncingRegions) WithValue(value []string) *GroupMember_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*GroupMember_FieldSubPathValue)
}

func (s GroupMemberPathSelectorMetadataSyncingRegions) WithArrayOfValues(values [][]string) *GroupMember_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*GroupMember_FieldSubPathArrayOfValues)
}

func (s GroupMemberPathSelectorMetadataSyncingRegions) WithItemValue(value string) *GroupMember_FieldSubPathArrayItemValue {
	return s.FieldPath().WithIArrayItemValue(value).(*GroupMember_FieldSubPathArrayItemValue)
}

type GroupMemberPathSelectorMetadataLifecycle struct{}

func (GroupMemberPathSelectorMetadataLifecycle) FieldPath() *GroupMember_FieldSubPath {
	return &GroupMember_FieldSubPath{
		selector: GroupMember_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Lifecycle().FieldPath(),
	}
}

func (s GroupMemberPathSelectorMetadataLifecycle) WithValue(value *meta.Lifecycle) *GroupMember_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*GroupMember_FieldSubPathValue)
}

func (s GroupMemberPathSelectorMetadataLifecycle) WithArrayOfValues(values []*meta.Lifecycle) *GroupMember_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*GroupMember_FieldSubPathArrayOfValues)
}

func (GroupMemberPathSelectorMetadataLifecycle) State() GroupMemberPathSelectorMetadataLifecycleState {
	return GroupMemberPathSelectorMetadataLifecycleState{}
}

func (GroupMemberPathSelectorMetadataLifecycle) BlockDeletion() GroupMemberPathSelectorMetadataLifecycleBlockDeletion {
	return GroupMemberPathSelectorMetadataLifecycleBlockDeletion{}
}

type GroupMemberPathSelectorMetadataLifecycleState struct{}

func (GroupMemberPathSelectorMetadataLifecycleState) FieldPath() *GroupMember_FieldSubPath {
	return &GroupMember_FieldSubPath{
		selector: GroupMember_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Lifecycle().State().FieldPath(),
	}
}

func (s GroupMemberPathSelectorMetadataLifecycleState) WithValue(value meta.Lifecycle_State) *GroupMember_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*GroupMember_FieldSubPathValue)
}

func (s GroupMemberPathSelectorMetadataLifecycleState) WithArrayOfValues(values []meta.Lifecycle_State) *GroupMember_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*GroupMember_FieldSubPathArrayOfValues)
}

type GroupMemberPathSelectorMetadataLifecycleBlockDeletion struct{}

func (GroupMemberPathSelectorMetadataLifecycleBlockDeletion) FieldPath() *GroupMember_FieldSubPath {
	return &GroupMember_FieldSubPath{
		selector: GroupMember_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Lifecycle().BlockDeletion().FieldPath(),
	}
}

func (s GroupMemberPathSelectorMetadataLifecycleBlockDeletion) WithValue(value bool) *GroupMember_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*GroupMember_FieldSubPathValue)
}

func (s GroupMemberPathSelectorMetadataLifecycleBlockDeletion) WithArrayOfValues(values []bool) *GroupMember_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*GroupMember_FieldSubPathArrayOfValues)
}

type GroupMemberPathSelectorMetadataServices struct{}

func (GroupMemberPathSelectorMetadataServices) FieldPath() *GroupMember_FieldSubPath {
	return &GroupMember_FieldSubPath{
		selector: GroupMember_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Services().FieldPath(),
	}
}

func (s GroupMemberPathSelectorMetadataServices) WithValue(value *meta.ServicesInfo) *GroupMember_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*GroupMember_FieldSubPathValue)
}

func (s GroupMemberPathSelectorMetadataServices) WithArrayOfValues(values []*meta.ServicesInfo) *GroupMember_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*GroupMember_FieldSubPathArrayOfValues)
}

func (GroupMemberPathSelectorMetadataServices) OwningService() GroupMemberPathSelectorMetadataServicesOwningService {
	return GroupMemberPathSelectorMetadataServicesOwningService{}
}

func (GroupMemberPathSelectorMetadataServices) AllowedServices() GroupMemberPathSelectorMetadataServicesAllowedServices {
	return GroupMemberPathSelectorMetadataServicesAllowedServices{}
}

type GroupMemberPathSelectorMetadataServicesOwningService struct{}

func (GroupMemberPathSelectorMetadataServicesOwningService) FieldPath() *GroupMember_FieldSubPath {
	return &GroupMember_FieldSubPath{
		selector: GroupMember_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Services().OwningService().FieldPath(),
	}
}

func (s GroupMemberPathSelectorMetadataServicesOwningService) WithValue(value string) *GroupMember_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*GroupMember_FieldSubPathValue)
}

func (s GroupMemberPathSelectorMetadataServicesOwningService) WithArrayOfValues(values []string) *GroupMember_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*GroupMember_FieldSubPathArrayOfValues)
}

type GroupMemberPathSelectorMetadataServicesAllowedServices struct{}

func (GroupMemberPathSelectorMetadataServicesAllowedServices) FieldPath() *GroupMember_FieldSubPath {
	return &GroupMember_FieldSubPath{
		selector: GroupMember_FieldPathSelectorMetadata,
		subPath:  meta.NewMetaFieldPathBuilder().Services().AllowedServices().FieldPath(),
	}
}

func (s GroupMemberPathSelectorMetadataServicesAllowedServices) WithValue(value []string) *GroupMember_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*GroupMember_FieldSubPathValue)
}

func (s GroupMemberPathSelectorMetadataServicesAllowedServices) WithArrayOfValues(values [][]string) *GroupMember_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*GroupMember_FieldSubPathArrayOfValues)
}

func (s GroupMemberPathSelectorMetadataServicesAllowedServices) WithItemValue(value string) *GroupMember_FieldSubPathArrayItemValue {
	return s.FieldPath().WithIArrayItemValue(value).(*GroupMember_FieldSubPathArrayItemValue)
}
