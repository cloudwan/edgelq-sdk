// Code generated by protoc-gen-goten-go
// File: edgelq/iam/proto/v1/member_assignment.proto
// DO NOT EDIT!!!

package member_assignment

import (
	"fmt"
	"reflect"
	"sync"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	preflect "google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
)

// proto imports
import (
	iam_common "github.com/cloudwan/edgelq-sdk/iam/resources/v1/common"
	organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1/organization"
	role_binding "github.com/cloudwan/edgelq-sdk/iam/resources/v1/role_binding"
	meta_service "github.com/cloudwan/goten-sdk/meta-service/resources/v1/service"
	meta "github.com/cloudwan/goten-sdk/types/meta"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = fmt.Errorf
	_ = reflect.Method{}
	_ = sync.Once{}

	_ = protojson.MarshalOptions{}
	_ = proto.MarshalOptions{}
	_ = preflect.Value{}
	_ = protoimpl.DescBuilder{}
)

// make sure we're using proto imports
var (
	_ = &iam_common.PCR{}
	_ = &organization.Organization{}
	_ = &role_binding.RoleBinding{}
	_ = &meta_service.Service{}
	_ = &meta.Meta{}
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// MemberAssignment - it is an internal resource, not for end users. It tracks
// REGIONAL RoleBindings per combination of scope/member for organization and
// project RoleBindings.
//
// If multiple RoleBindings point to same project/org and member, they will all
// share single MemberAssignment.
//
// MemberAssignment are managed by IAM Server, they are created/updated/deleted
// in TX when RoleBinding is created/updated/deleted. MemberAssignment shares
// same region as RoleBinding, otherwise tx-level synchronization would not be
// possible.
//
// For example: We create RoleBinding { project = "X", member = "M", role =
// "r1"}. It will create MemberAssignment { scope = "projects/x", member = "M"
// }. If another RoleBinding with same project and member would be created, then
// no additional MemberAssignment is created. When last RoleBinding per
// scope/member is deleted, MemberAssignment is deleted.
//
// Not all RoleBindings however have MemberAssignment instances. We track only
// organization and project RoleBindings! Therefore, system and service
// RoleBindings don't get their MemberAssignment. However, there is some caveat
// about this...
//
// When lets say project enables Service "S", and we create RoleBinding WHERE:
// {project = "X", member = "M", role = "r1", metadata.services.allowedServices
// CONTAINS "S"}, then special MemberAssignment is created with params: { scope
// = "services/S", member = "M" }, apart of { scope = "projects/x", member = "M"
// } mentioned previously. We create those service MemberAssignment only for 3rd
// party services (non core EdgeLQ), so we know if User/ServiceAccount is
// eligible user of some service.
//
// Main task of MemberAssignment is to track participations of all users/service
// accounts in projects/organizations. We use it for things like ListMyProjects,
// ListMyOrganizations. We also use those special service MemberAssignment
// instances to track who is using Service by proxy of Project/Organization!
// With this we can forbid/allow specific users using particular service.
//
// System RoleBindings are managed only by EdgeLQ admins for internal cases, and
// we dont need this tracking.
type MemberAssignment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Name of MemberAssignment
	// First letter is source indicator (p, o, s for project, org, service).
	// Then we have proper scope identifier (projectId etc), then member
	// identifier (like user:$EMAIL).
	Name *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// Metadata is an object with information like create, update and delete time
	// (for async deleted resources), has user labels/annotations, sharding
	// information, multi-region syncing information and may have non-schema
	// owners (useful for taking ownership of resources belonging to lower level
	// services by higher ones).
	Metadata *meta.Meta `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty" firestore:"metadata"`
	// Points to Organization/Project of RoleBindings, OR service for those
	// special MemberAssignments. It is already part of name, but we track
	// in field for filter purpose.
	Scope *role_binding.ParentName `protobuf:"bytes,3,opt,customtype=ParentName,name=scope,proto3" json:"scope,omitempty" firestore:"scope"`
	// Populated for organization/project scopes. Skipped for service ones.
	ScopeTitle string `protobuf:"bytes,4,opt,name=scope_title,json=scopeTitle,proto3" json:"scope_title,omitempty" firestore:"scopeTitle"`
	// Populated for organization/project scopes. Skipped for service ones.
	ParentOrganization *organization.Name `protobuf:"bytes,5,opt,customtype=Name,name=parent_organization,json=parentOrganization,proto3" json:"parent_organization,omitempty" firestore:"parentOrganization"`
	// It has PARTIAL metadata inherited from scope (labels, annotations, tags).
	// Populated for organization/project scopes. Skipped for service ones.
	ScopeMetadata *meta.Meta `protobuf:"bytes,6,opt,name=scope_metadata,json=scopeMetadata,proto3" json:"scope_metadata,omitempty" firestore:"scopeMetadata"`
	// Populated for organization/project scopes. Skipped for service ones.
	// Contains multi_region_policy.default_control_region
	MultiRegionControlRegion string `protobuf:"bytes,7,opt,name=multi_region_control_region,json=multiRegionControlRegion,proto3" json:"multi_region_control_region,omitempty" firestore:"multiRegionControlRegion"`
	// Populated for organization/project scopes. Skipped for service ones.
	// Contains multi_region_policy.enabled_regions
	MultiRegionEnabledRegions []string `protobuf:"bytes,8,rep,name=multi_region_enabled_regions,json=multiRegionEnabledRegions,proto3" json:"multi_region_enabled_regions,omitempty" firestore:"multiRegionEnabledRegions"`
	// Populated for organization/project scopes. Skipped for service ones.
	// Contains allowed or enabled services.
	ScopeServices []*meta_service.Name `protobuf:"bytes,9,rep,customtype=Name,name=scope_services,json=scopeServices,proto3" json:"scope_services,omitempty" firestore:"scopeServices"`
	// Populated for organization/project scopes. Skipped for service ones.
	BusinessTier iam_common.BusinessTier `protobuf:"varint,10,opt,name=business_tier,json=businessTier,proto3,enum=ntt.iam.v1.BusinessTier" json:"business_tier,omitempty" firestore:"businessTier"`
	// Member pointed by RoleBinding. Part of name, but we also need for filtering
	// purposes.
	Member string `protobuf:"bytes,11,opt,name=member,proto3" json:"member,omitempty" firestore:"member"`
	// Region ID holding member resource (User, ServiceAccount...)
	MemberRegion string                       `protobuf:"bytes,12,opt,name=member_region,json=memberRegion,proto3" json:"member_region,omitempty" firestore:"memberRegion"`
	CtrlStatus   *MemberAssignment_WorkStatus `protobuf:"bytes,13,opt,name=ctrl_status,json=ctrlStatus,proto3" json:"ctrl_status,omitempty" firestore:"ctrlStatus"`
}

func (m *MemberAssignment) Reset() {
	*m = MemberAssignment{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_iam_proto_v1_member_assignment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *MemberAssignment) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*MemberAssignment) ProtoMessage() {}

func (m *MemberAssignment) ProtoReflect() preflect.Message {
	mi := &edgelq_iam_proto_v1_member_assignment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*MemberAssignment) GotenMessage() {}

// Deprecated, Use MemberAssignment.ProtoReflect.Descriptor instead.
func (*MemberAssignment) Descriptor() ([]byte, []int) {
	return edgelq_iam_proto_v1_member_assignment_proto_rawDescGZIP(), []int{0}
}

func (m *MemberAssignment) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *MemberAssignment) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *MemberAssignment) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *MemberAssignment) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *MemberAssignment) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *MemberAssignment) GetMetadata() *meta.Meta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *MemberAssignment) GetScope() *role_binding.ParentName {
	if m != nil {
		return m.Scope
	}
	return nil
}

func (m *MemberAssignment) GetScopeTitle() string {
	if m != nil {
		return m.ScopeTitle
	}
	return ""
}

func (m *MemberAssignment) GetParentOrganization() *organization.Name {
	if m != nil {
		return m.ParentOrganization
	}
	return nil
}

func (m *MemberAssignment) GetScopeMetadata() *meta.Meta {
	if m != nil {
		return m.ScopeMetadata
	}
	return nil
}

func (m *MemberAssignment) GetMultiRegionControlRegion() string {
	if m != nil {
		return m.MultiRegionControlRegion
	}
	return ""
}

func (m *MemberAssignment) GetMultiRegionEnabledRegions() []string {
	if m != nil {
		return m.MultiRegionEnabledRegions
	}
	return nil
}

func (m *MemberAssignment) GetScopeServices() []*meta_service.Name {
	if m != nil {
		return m.ScopeServices
	}
	return nil
}

func (m *MemberAssignment) GetBusinessTier() iam_common.BusinessTier {
	if m != nil {
		return m.BusinessTier
	}
	return iam_common.BusinessTier_UNDEFINED
}

func (m *MemberAssignment) GetMember() string {
	if m != nil {
		return m.Member
	}
	return ""
}

func (m *MemberAssignment) GetMemberRegion() string {
	if m != nil {
		return m.MemberRegion
	}
	return ""
}

func (m *MemberAssignment) GetCtrlStatus() *MemberAssignment_WorkStatus {
	if m != nil {
		return m.CtrlStatus
	}
	return nil
}

func (m *MemberAssignment) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "MemberAssignment"))
	}
	m.Name = fv
}

func (m *MemberAssignment) SetMetadata(fv *meta.Meta) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Metadata", "MemberAssignment"))
	}
	m.Metadata = fv
}

func (m *MemberAssignment) SetScope(fv *role_binding.ParentName) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Scope", "MemberAssignment"))
	}
	m.Scope = fv
}

func (m *MemberAssignment) SetScopeTitle(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ScopeTitle", "MemberAssignment"))
	}
	m.ScopeTitle = fv
}

func (m *MemberAssignment) SetParentOrganization(fv *organization.Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ParentOrganization", "MemberAssignment"))
	}
	m.ParentOrganization = fv
}

func (m *MemberAssignment) SetScopeMetadata(fv *meta.Meta) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ScopeMetadata", "MemberAssignment"))
	}
	m.ScopeMetadata = fv
}

func (m *MemberAssignment) SetMultiRegionControlRegion(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "MultiRegionControlRegion", "MemberAssignment"))
	}
	m.MultiRegionControlRegion = fv
}

func (m *MemberAssignment) SetMultiRegionEnabledRegions(fv []string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "MultiRegionEnabledRegions", "MemberAssignment"))
	}
	m.MultiRegionEnabledRegions = fv
}

func (m *MemberAssignment) SetScopeServices(fv []*meta_service.Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ScopeServices", "MemberAssignment"))
	}
	m.ScopeServices = fv
}

func (m *MemberAssignment) SetBusinessTier(fv iam_common.BusinessTier) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "BusinessTier", "MemberAssignment"))
	}
	m.BusinessTier = fv
}

func (m *MemberAssignment) SetMember(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Member", "MemberAssignment"))
	}
	m.Member = fv
}

func (m *MemberAssignment) SetMemberRegion(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "MemberRegion", "MemberAssignment"))
	}
	m.MemberRegion = fv
}

func (m *MemberAssignment) SetCtrlStatus(fv *MemberAssignment_WorkStatus) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "CtrlStatus", "MemberAssignment"))
	}
	m.CtrlStatus = fv
}

type MemberAssignment_WorkStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// if controller has some work on this resource
	Pending bool `protobuf:"varint,1,opt,name=pending,proto3" json:"pending,omitempty" firestore:"pending"`
	// If this resource should be deleted.
	PendingDeletion bool `protobuf:"varint,2,opt,name=pending_deletion,json=pendingDeletion,proto3" json:"pending_deletion,omitempty" firestore:"pendingDeletion"`
}

func (m *MemberAssignment_WorkStatus) Reset() {
	*m = MemberAssignment_WorkStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_iam_proto_v1_member_assignment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *MemberAssignment_WorkStatus) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*MemberAssignment_WorkStatus) ProtoMessage() {}

func (m *MemberAssignment_WorkStatus) ProtoReflect() preflect.Message {
	mi := &edgelq_iam_proto_v1_member_assignment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*MemberAssignment_WorkStatus) GotenMessage() {}

// Deprecated, Use MemberAssignment_WorkStatus.ProtoReflect.Descriptor instead.
func (*MemberAssignment_WorkStatus) Descriptor() ([]byte, []int) {
	return edgelq_iam_proto_v1_member_assignment_proto_rawDescGZIP(), []int{0, 0}
}

func (m *MemberAssignment_WorkStatus) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *MemberAssignment_WorkStatus) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *MemberAssignment_WorkStatus) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *MemberAssignment_WorkStatus) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *MemberAssignment_WorkStatus) GetPending() bool {
	if m != nil {
		return m.Pending
	}
	return false
}

func (m *MemberAssignment_WorkStatus) GetPendingDeletion() bool {
	if m != nil {
		return m.PendingDeletion
	}
	return false
}

func (m *MemberAssignment_WorkStatus) SetPending(fv bool) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Pending", "MemberAssignment_WorkStatus"))
	}
	m.Pending = fv
}

func (m *MemberAssignment_WorkStatus) SetPendingDeletion(fv bool) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "PendingDeletion", "MemberAssignment_WorkStatus"))
	}
	m.PendingDeletion = fv
}

var edgelq_iam_proto_v1_member_assignment_proto preflect.FileDescriptor

var edgelq_iam_proto_v1_member_assignment_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x61, 0x73, 0x73,
	0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6e,
	0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69,
	0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x67,
	0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61,
	0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc2, 0x0a, 0x0a, 0x10, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x2c, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xb2, 0xda, 0x21, 0x14,
	0x0a, 0x12, 0x0a, 0x10, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67,
	0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x2d, 0x0a, 0x05, 0x73, 0x63, 0x6f,
	0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x17, 0xf0, 0xd9, 0x21, 0x01, 0xb2, 0xda,
	0x21, 0x0f, 0x3a, 0x0d, 0x0a, 0x0b, 0x52, 0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x25, 0x0a, 0x0b, 0x73, 0x63, 0x6f, 0x70,
	0x65, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xf0,
	0xd9, 0x21, 0x01, 0x52, 0x0a, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x49, 0x0a, 0x13, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xf0, 0xd9,
	0x21, 0x01, 0xb2, 0xda, 0x21, 0x10, 0x0a, 0x0e, 0x0a, 0x0c, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x12, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x4f, 0x72,
	0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3e, 0x0a, 0x0e, 0x73, 0x63,
	0x6f, 0x70, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x4d, 0x65, 0x74, 0x61, 0x42, 0x04, 0xf0, 0xd9, 0x21, 0x01, 0x52, 0x0d, 0x73, 0x63, 0x6f,
	0x70, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x43, 0x0a, 0x1b, 0x6d, 0x75,
	0x6c, 0x74, 0x69, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x04, 0xf0, 0xd9, 0x21, 0x01, 0x52, 0x18, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x52, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12,
	0x45, 0x0a, 0x1c, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x08, 0x20, 0x03, 0x28, 0x09, 0x42, 0x04, 0xf0, 0xd9, 0x21, 0x01, 0x52, 0x19, 0x6d, 0x75, 0x6c,
	0x74, 0x69, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x52,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x49, 0x0a, 0x0e, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x42, 0x22,
	0xf0, 0xd9, 0x21, 0x01, 0xb2, 0xda, 0x21, 0x1a, 0x0a, 0x18, 0x0a, 0x16, 0x6d, 0x65, 0x74, 0x61,
	0x2e, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x52, 0x0d, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x12, 0x43, 0x0a, 0x0d, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x69,
	0x65, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69,
	0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x54, 0x69,
	0x65, 0x72, 0x42, 0x04, 0xf0, 0xd9, 0x21, 0x01, 0x52, 0x0c, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65,
	0x73, 0x73, 0x54, 0x69, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xf0, 0xd9, 0x21, 0x01, 0x52, 0x06, 0x6d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x0d, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xf0, 0xd9, 0x21,
	0x01, 0x52, 0x0c, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12,
	0x48, 0x0a, 0x0b, 0x63, 0x74, 0x72, 0x6c, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76,
	0x31, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0a, 0x63,
	0x74, 0x72, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x1a, 0x51, 0x0a, 0x0a, 0x57, 0x6f, 0x72,
	0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x65, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x12, 0x29, 0x0a, 0x10, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x70, 0x65, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0xed, 0x03, 0xea,
	0x41, 0x59, 0x0a, 0x1f, 0x69, 0x61, 0x6d, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x36, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x72, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x41, 0x73, 0x73, 0x69,
	0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f,
	0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x7d, 0x92, 0xd9, 0x21, 0x4a, 0x0a,
	0x11, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x12, 0x11, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x18, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2a,
	0x08, 0x2e, 0x7b, 0x31, 0x2c, 0x35, 0x31, 0x32, 0x7d, 0xb2, 0xdf, 0x21, 0xac, 0x02, 0x0a, 0x92,
	0x02, 0x0a, 0x31, 0x0a, 0x07, 0x62, 0x79, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x05, 0x73, 0x63,
	0x6f, 0x70, 0x65, 0x1a, 0x1d, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2d, 0x2f, 0x6d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x2f, 0x2d, 0x20, 0x02, 0x0a, 0x3c, 0x0a, 0x0c, 0x62, 0x79, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x54,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x0b, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x5f, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x1a, 0x1d, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2d, 0x2f, 0x6d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x2d,
	0x20, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x10, 0x01, 0x12, 0x32,
	0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x10, 0x01, 0x22, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x2f, 0x2d, 0x22, 0x0f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2d, 0x22, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2f, 0x2d, 0x12, 0x15, 0x0a, 0x13, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x6f, 0x72, 0x67,
	0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x0e, 0x73, 0x63, 0x6f,
	0x70, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x36, 0x0a, 0x34, 0x73,
	0x63, 0x6f, 0x70, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2d, 0x69, 0x73, 0x2d, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2d, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x15, 0x0a, 0x13, 0x63, 0x74, 0x72, 0x6c, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x2e, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0xda, 0x94, 0x23, 0x08, 0x12, 0x06,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0xe2, 0xde, 0x21, 0x02, 0x08, 0x03, 0x42, 0xb5, 0x02, 0xe8,
	0xde, 0x21, 0x01, 0xd2, 0xff, 0xd0, 0x02, 0x54, 0x0a, 0x17, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x5f, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x12, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61,
	0x6d, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x5f, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0xa2, 0x80, 0xd1, 0x02,
	0x56, 0x0a, 0x18, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x3a, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e,
	0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x61, 0x73, 0x73,
	0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x74, 0x74,
	0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x42, 0x15, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x4f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f,
	0x69, 0x61, 0x6d, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65,
	0x6e, 0x74, 0x3b, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	edgelq_iam_proto_v1_member_assignment_proto_rawDescOnce sync.Once
	edgelq_iam_proto_v1_member_assignment_proto_rawDescData = edgelq_iam_proto_v1_member_assignment_proto_rawDesc
)

func edgelq_iam_proto_v1_member_assignment_proto_rawDescGZIP() []byte {
	edgelq_iam_proto_v1_member_assignment_proto_rawDescOnce.Do(func() {
		edgelq_iam_proto_v1_member_assignment_proto_rawDescData = protoimpl.X.CompressGZIP(edgelq_iam_proto_v1_member_assignment_proto_rawDescData)
	})
	return edgelq_iam_proto_v1_member_assignment_proto_rawDescData
}

var edgelq_iam_proto_v1_member_assignment_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var edgelq_iam_proto_v1_member_assignment_proto_goTypes = []interface{}{
	(*MemberAssignment)(nil),            // 0: ntt.iam.v1.MemberAssignment
	(*MemberAssignment_WorkStatus)(nil), // 1: ntt.iam.v1.MemberAssignment.WorkStatus
	(*meta.Meta)(nil),                   // 2: goten.types.Meta
	(iam_common.BusinessTier)(0),        // 3: ntt.iam.v1.BusinessTier
}
var edgelq_iam_proto_v1_member_assignment_proto_depIdxs = []int32{
	2, // 0: ntt.iam.v1.MemberAssignment.metadata:type_name -> goten.types.Meta
	2, // 1: ntt.iam.v1.MemberAssignment.scope_metadata:type_name -> goten.types.Meta
	3, // 2: ntt.iam.v1.MemberAssignment.business_tier:type_name -> ntt.iam.v1.BusinessTier
	1, // 3: ntt.iam.v1.MemberAssignment.ctrl_status:type_name -> ntt.iam.v1.MemberAssignment.WorkStatus
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { edgelq_iam_proto_v1_member_assignment_proto_init() }
func edgelq_iam_proto_v1_member_assignment_proto_init() {
	if edgelq_iam_proto_v1_member_assignment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		edgelq_iam_proto_v1_member_assignment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberAssignment); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		edgelq_iam_proto_v1_member_assignment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberAssignment_WorkStatus); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}

	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: edgelq_iam_proto_v1_member_assignment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           edgelq_iam_proto_v1_member_assignment_proto_goTypes,
		DependencyIndexes: edgelq_iam_proto_v1_member_assignment_proto_depIdxs,
		MessageInfos:      edgelq_iam_proto_v1_member_assignment_proto_msgTypes,
	}.Build()
	edgelq_iam_proto_v1_member_assignment_proto = out.File
	edgelq_iam_proto_v1_member_assignment_proto_rawDesc = nil
	edgelq_iam_proto_v1_member_assignment_proto_goTypes = nil
	edgelq_iam_proto_v1_member_assignment_proto_depIdxs = nil
}
