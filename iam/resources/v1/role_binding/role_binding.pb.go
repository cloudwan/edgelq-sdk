// Code generated by protoc-gen-goten-go
// File: edgelq/iam/proto/v1/role_binding.proto
// DO NOT EDIT!!!

package role_binding

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
	condition "github.com/cloudwan/edgelq-sdk/iam/resources/v1/condition"
	organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1/organization"
	project "github.com/cloudwan/edgelq-sdk/iam/resources/v1/project"
	role "github.com/cloudwan/edgelq-sdk/iam/resources/v1/role"
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
	_ = &condition.Condition{}
	_ = &organization.Organization{}
	_ = &project.Project{}
	_ = &role.Role{}
	_ = &meta_service.Service{}
	_ = &meta.Meta{}
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// RoleBinding Resource
type RoleBinding struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Name of RoleBinding
	Name *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// Metadata is an object with information like create, update and delete time
	// (for async deleted resources), has user labels/annotations, sharding
	// information, multi-region syncing information and may have non-schema
	// owners (useful for taking ownership of resources belonging to lower level
	// services by higher ones).
	Metadata *meta.Meta `protobuf:"bytes,8,opt,name=metadata,proto3" json:"metadata,omitempty" firestore:"metadata"`
	// Applied role.
	Role *role.Reference `protobuf:"bytes,2,opt,customtype=Reference,name=role,proto3" json:"role,omitempty" firestore:"role"`
	// Binding members
	// Format of the string is one of:
	// - "allUsers" (anyone)
	// - "allAuthenticatedUsers" (anyone logged in from handling service point of
	// view)
	// - "user:admin.super@example.com"
	// - "serviceAccount:device_agent@watchdog.serviceaccounts.iam.edgelq.com"
	// - "group:nice.group@example.com"
	// - "domain:example.com" (anyone with exact email domain)
	Member string `protobuf:"bytes,4,opt,name=member,proto3" json:"member,omitempty" firestore:"member"`
	// All scope params defined as required by a role
	ScopeParams []*role.ScopeParam `protobuf:"bytes,5,rep,name=scope_params,json=scopeParams,proto3" json:"scope_params,omitempty" firestore:"scopeParams"`
	// optional executable conditions to be added to the role binding.
	// They are matched with the executable conditions in a role grants
	// by condition reference. If there is condition defined in RoleBinding,
	// but not in any role grant, then executable condition is applied to
	// all role grants.
	ExecutableConditions []*condition.ExecutableCondition `protobuf:"bytes,6,rep,name=executable_conditions,json=executableConditions,proto3" json:"executable_conditions,omitempty" firestore:"executableConditions"`
	// List of owned objects WITHIN role binding scope - for example
	// if role binding is a child of project and owner_objects contain some
	// device, member has all possible permissions for this device in project,
	// regardless of method. In case ownership is for whole role binding scope, it
	// will contain "-" string. This is computed based on a role with
	// wildcard grants.
	OwnedObjects []string `protobuf:"bytes,3,rep,name=owned_objects,json=ownedObjects,proto3" json:"owned_objects,omitempty" firestore:"ownedObjects"`
	// Member type, its prefix from member before ':' rune, fo example "user".
	// If member has no ':' (like allUsers), then it will contain same value.
	MemberType string `protobuf:"bytes,9,opt,name=member_type,json=memberType,proto3" json:"member_type,omitempty" firestore:"memberType"`
	// Role category, taken from role itself, allows for additional filtering.
	Category role.Role_Category `protobuf:"varint,14,opt,name=category,proto3,enum=ntt.iam.v1.Role_Category" json:"category,omitempty" firestore:"category"`
	// Internal field used by IAM controller to note role binding ancestry path
	// for Group type (RoleBindings inherited from Group)
	AncestryPath []*RoleBinding_Parent `protobuf:"bytes,7,rep,name=ancestry_path,json=ancestryPath,proto3" json:"ancestry_path,omitempty" firestore:"ancestryPath"`
	// Internal field used by IAM controller to note parent role binding
	// from Parent Organization type. It is much different compared to
	// "ancestry_path", we just need direct parent. This is because it is simpler
	// - controller for this inheritance does not need full ancestry path for loop
	// detection. Groups are complicated and it is legal for two groups containing
	// each other as members. But organization ancestry path does not allow loops,
	// so we can afford simple field!
	ParentByOrg *Reference `protobuf:"bytes,12,opt,customtype=Reference,name=parent_by_org,json=parentByOrg,proto3" json:"parent_by_org,omitempty" firestore:"parentByOrg"`
	// Internal field used to synchronize role binding with role.
	// This value increases when we spec in a role changes in a way requiring role
	// binding resynchronization.
	SpecGeneration int64 `protobuf:"varint,10,opt,name=spec_generation,json=specGeneration,proto3" json:"spec_generation,omitempty" firestore:"specGeneration"`
	// Internal field indicating if role binding has owned objects.
	// It is used for filtering, as its not possible to filter by
	// len(owned_objects) > 0
	HasOwnedObjects bool `protobuf:"varint,11,opt,name=has_owned_objects,json=hasOwnedObjects,proto3" json:"has_owned_objects,omitempty" firestore:"hasOwnedObjects"`
	// If true, then this RoleBinding DOES NOT apply to child entities.
	// TODO: For now IAM Server decides if this RB is not assignable, consider
	// opening this.
	DisableForChildScopes bool `protobuf:"varint,13,opt,name=disable_for_child_scopes,json=disableForChildScopes,proto3" json:"disable_for_child_scopes,omitempty" firestore:"disableForChildScopes"`
}

func (m *RoleBinding) Reset() {
	*m = RoleBinding{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_iam_proto_v1_role_binding_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *RoleBinding) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*RoleBinding) ProtoMessage() {}

func (m *RoleBinding) ProtoReflect() preflect.Message {
	mi := &edgelq_iam_proto_v1_role_binding_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*RoleBinding) GotenMessage() {}

// Deprecated, Use RoleBinding.ProtoReflect.Descriptor instead.
func (*RoleBinding) Descriptor() ([]byte, []int) {
	return edgelq_iam_proto_v1_role_binding_proto_rawDescGZIP(), []int{0}
}

func (m *RoleBinding) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *RoleBinding) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *RoleBinding) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *RoleBinding) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *RoleBinding) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *RoleBinding) GetMetadata() *meta.Meta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *RoleBinding) GetRole() *role.Reference {
	if m != nil {
		return m.Role
	}
	return nil
}

func (m *RoleBinding) GetMember() string {
	if m != nil {
		return m.Member
	}
	return ""
}

func (m *RoleBinding) GetScopeParams() []*role.ScopeParam {
	if m != nil {
		return m.ScopeParams
	}
	return nil
}

func (m *RoleBinding) GetExecutableConditions() []*condition.ExecutableCondition {
	if m != nil {
		return m.ExecutableConditions
	}
	return nil
}

func (m *RoleBinding) GetOwnedObjects() []string {
	if m != nil {
		return m.OwnedObjects
	}
	return nil
}

func (m *RoleBinding) GetMemberType() string {
	if m != nil {
		return m.MemberType
	}
	return ""
}

func (m *RoleBinding) GetCategory() role.Role_Category {
	if m != nil {
		return m.Category
	}
	return role.Role_UNDEFINED
}

func (m *RoleBinding) GetAncestryPath() []*RoleBinding_Parent {
	if m != nil {
		return m.AncestryPath
	}
	return nil
}

func (m *RoleBinding) GetParentByOrg() *Reference {
	if m != nil {
		return m.ParentByOrg
	}
	return nil
}

func (m *RoleBinding) GetSpecGeneration() int64 {
	if m != nil {
		return m.SpecGeneration
	}
	return int64(0)
}

func (m *RoleBinding) GetHasOwnedObjects() bool {
	if m != nil {
		return m.HasOwnedObjects
	}
	return false
}

func (m *RoleBinding) GetDisableForChildScopes() bool {
	if m != nil {
		return m.DisableForChildScopes
	}
	return false
}

func (m *RoleBinding) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "RoleBinding"))
	}
	m.Name = fv
}

func (m *RoleBinding) SetMetadata(fv *meta.Meta) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Metadata", "RoleBinding"))
	}
	m.Metadata = fv
}

func (m *RoleBinding) SetRole(fv *role.Reference) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Role", "RoleBinding"))
	}
	m.Role = fv
}

func (m *RoleBinding) SetMember(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Member", "RoleBinding"))
	}
	m.Member = fv
}

func (m *RoleBinding) SetScopeParams(fv []*role.ScopeParam) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ScopeParams", "RoleBinding"))
	}
	m.ScopeParams = fv
}

func (m *RoleBinding) SetExecutableConditions(fv []*condition.ExecutableCondition) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ExecutableConditions", "RoleBinding"))
	}
	m.ExecutableConditions = fv
}

func (m *RoleBinding) SetOwnedObjects(fv []string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "OwnedObjects", "RoleBinding"))
	}
	m.OwnedObjects = fv
}

func (m *RoleBinding) SetMemberType(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "MemberType", "RoleBinding"))
	}
	m.MemberType = fv
}

func (m *RoleBinding) SetCategory(fv role.Role_Category) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Category", "RoleBinding"))
	}
	m.Category = fv
}

func (m *RoleBinding) SetAncestryPath(fv []*RoleBinding_Parent) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "AncestryPath", "RoleBinding"))
	}
	m.AncestryPath = fv
}

func (m *RoleBinding) SetParentByOrg(fv *Reference) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ParentByOrg", "RoleBinding"))
	}
	m.ParentByOrg = fv
}

func (m *RoleBinding) SetSpecGeneration(fv int64) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "SpecGeneration", "RoleBinding"))
	}
	m.SpecGeneration = fv
}

func (m *RoleBinding) SetHasOwnedObjects(fv bool) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "HasOwnedObjects", "RoleBinding"))
	}
	m.HasOwnedObjects = fv
}

func (m *RoleBinding) SetDisableForChildScopes(fv bool) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "DisableForChildScopes", "RoleBinding"))
	}
	m.DisableForChildScopes = fv
}

// Provides information about inheritance of this role binding - if it was
// created from another role binding.
type RoleBinding_Parent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Role bindings which have group as member have child for each service
	// account and user in that group. All those role bindings have one common
	// parent pointing at group role binding
	Parent *Reference `protobuf:"bytes,1,opt,customtype=Reference,name=parent,proto3" json:"parent,omitempty" firestore:"parent"`
	// Member of the parent role binding
	Member string `protobuf:"bytes,2,opt,name=member,proto3" json:"member,omitempty" firestore:"member"`
}

func (m *RoleBinding_Parent) Reset() {
	*m = RoleBinding_Parent{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_iam_proto_v1_role_binding_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *RoleBinding_Parent) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*RoleBinding_Parent) ProtoMessage() {}

func (m *RoleBinding_Parent) ProtoReflect() preflect.Message {
	mi := &edgelq_iam_proto_v1_role_binding_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*RoleBinding_Parent) GotenMessage() {}

// Deprecated, Use RoleBinding_Parent.ProtoReflect.Descriptor instead.
func (*RoleBinding_Parent) Descriptor() ([]byte, []int) {
	return edgelq_iam_proto_v1_role_binding_proto_rawDescGZIP(), []int{0, 0}
}

func (m *RoleBinding_Parent) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *RoleBinding_Parent) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *RoleBinding_Parent) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *RoleBinding_Parent) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *RoleBinding_Parent) GetParent() *Reference {
	if m != nil {
		return m.Parent
	}
	return nil
}

func (m *RoleBinding_Parent) GetMember() string {
	if m != nil {
		return m.Member
	}
	return ""
}

func (m *RoleBinding_Parent) SetParent(fv *Reference) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Parent", "RoleBinding_Parent"))
	}
	m.Parent = fv
}

func (m *RoleBinding_Parent) SetMember(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Member", "RoleBinding_Parent"))
	}
	m.Member = fv
}

var edgelq_iam_proto_v1_role_binding_proto preflect.FileDescriptor

var edgelq_iam_proto_v1_role_binding_proto_rawDesc = []byte{
	0x0a, 0x26, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61,
	0x6d, 0x2e, 0x76, 0x31, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x24, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f, 0x72, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71,
	0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16,
	0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6d, 0x65, 0x74, 0x61,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69,
	0x61, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x65, 0x64, 0x67,
	0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31,
	0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcd, 0x10, 0x0a, 0x0b,
	0x52, 0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x27, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x13, 0xb2, 0xda, 0x21, 0x0f, 0x0a,
	0x0d, 0x0a, 0x0b, 0x52, 0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x22, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x0e, 0xb2, 0xda, 0x21, 0x0a, 0x12, 0x08, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x10,
	0x01, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xca, 0xc6, 0x27, 0x04, 0x2a, 0x02, 0x68,
	0x01, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x39, 0x0a, 0x0c, 0x73, 0x63, 0x6f,
	0x70, 0x65, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x6f,
	0x70, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x0b, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x12, 0x54, 0x0a, 0x15, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31,
	0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x64, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x14, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x29, 0x0a, 0x0d, 0x6f, 0x77,
	0x6e, 0x65, 0x64, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x09, 0x42, 0x04, 0xf0, 0xd9, 0x21, 0x01, 0x52, 0x0c, 0x6f, 0x77, 0x6e, 0x65, 0x64, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x25, 0x0a, 0x0b, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xf0, 0xd9, 0x21, 0x01,
	0x52, 0x0a, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x3b, 0x0a, 0x08,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19,
	0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6c, 0x65,
	0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x42, 0x04, 0xf0, 0xd9, 0x21, 0x01, 0x52,
	0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x43, 0x0a, 0x0d, 0x61, 0x6e, 0x63,
	0x65, 0x73, 0x74, 0x72, 0x79, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f,
	0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x52, 0x0c, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x74, 0x72, 0x79, 0x50, 0x61, 0x74, 0x68, 0x12, 0x39,
	0x0a, 0x0d, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x62, 0x79, 0x5f, 0x6f, 0x72, 0x67, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x09, 0x42, 0x15, 0xb2, 0xda, 0x21, 0x11, 0x12, 0x0f, 0x0a, 0x0b, 0x52,
	0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x10, 0x05, 0x52, 0x0b, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x4f, 0x72, 0x67, 0x12, 0x2d, 0x0a, 0x0f, 0x73, 0x70, 0x65,
	0x63, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x03, 0x42, 0x04, 0xf0, 0xd9, 0x21, 0x01, 0x52, 0x0e, 0x73, 0x70, 0x65, 0x63, 0x47, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x11, 0x68, 0x61, 0x73, 0x5f,
	0x6f, 0x77, 0x6e, 0x65, 0x64, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x08, 0x42, 0x04, 0xf0, 0xd9, 0x21, 0x01, 0x52, 0x0f, 0x68, 0x61, 0x73, 0x4f, 0x77,
	0x6e, 0x65, 0x64, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x3d, 0x0a, 0x18, 0x64, 0x69,
	0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x66, 0x6f, 0x72, 0x5f, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x5f,
	0x73, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x42, 0x04, 0xf0, 0xd9,
	0x21, 0x01, 0x52, 0x15, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x46, 0x6f, 0x72, 0x43, 0x68,
	0x69, 0x6c, 0x64, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x1a, 0x5e, 0x0a, 0x06, 0x50, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x12, 0x2d, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x15, 0xb2, 0xda, 0x21, 0x11, 0x12, 0x0f, 0x0a, 0x0b, 0x52, 0x6f, 0x6c,
	0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x10, 0x06, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x12, 0x25, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x0d, 0xca, 0xc6, 0x27, 0x09, 0x2a, 0x07, 0x3a, 0x03, 0x08, 0x80, 0x02, 0x68,
	0x01, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x3a, 0xff, 0x09, 0xea, 0x41, 0xd3, 0x01,
	0x0a, 0x1a, 0x69, 0x61, 0x6d, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x52, 0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x1b, 0x72, 0x6f,
	0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x7b, 0x72, 0x6f, 0x6c, 0x65,
	0x5f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x7d, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x72, 0x6f,
	0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x7b, 0x72, 0x6f, 0x6c, 0x65,
	0x5f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x7d, 0x12, 0x38, 0x6f, 0x72, 0x67, 0x61, 0x6e,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64,
	0x69, 0x6e, 0x67, 0x73, 0x2f, 0x7b, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x7d, 0x12, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x7b, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x7d, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64,
	0x69, 0x6e, 0x67, 0x73, 0x2f, 0x7b, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x7d, 0x92, 0xd9, 0x21, 0x8f, 0x01, 0x0a, 0x0c, 0x72, 0x6f, 0x6c, 0x65, 0x42, 0x69,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x0c, 0x72, 0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64,
	0x69, 0x6e, 0x67, 0x73, 0x1a, 0x12, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x6e, 0x74, 0x74, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x6f, 0x6e, 0x65, 0x1a, 0x07, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x1a, 0x0c, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a,
	0x16, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2a, 0x0e, 0x5b, 0x5c, 0x77, 0x2e, 0x7c, 0x2d, 0x5d,
	0x7b, 0x31, 0x2c, 0x31, 0x32, 0x38, 0x7d, 0x38, 0x05, 0x42, 0x1c, 0x08, 0x02, 0x12, 0x06, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x06, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x08, 0x0a,
	0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0xb2, 0xdf, 0x21, 0xa3, 0x05, 0x0a, 0x85, 0x05, 0x0a,
	0x42, 0x0a, 0x0a, 0x62, 0x79, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x46, 0x73, 0x12, 0x06, 0x6d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x1a, 0x21, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f,
	0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x42, 0x69,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x2d, 0x2a, 0x09, 0x66, 0x69, 0x72, 0x65, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x0a, 0xa1, 0x01, 0x0a, 0x0d, 0x62, 0x79, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x4d, 0x6f, 0x6e, 0x67, 0x6f, 0x12, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x1a, 0x21, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x7b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x7d, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x2d,
	0x1a, 0x21, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x7d, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x73, 0x2f, 0x2d, 0x1a, 0x2b, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x7b, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x7d, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x2d,
	0x1a, 0x0e, 0x72, 0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x2d,
	0x2a, 0x05, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x0a, 0x7f, 0x0a, 0x20, 0x62, 0x79, 0x4e, 0x61, 0x6d,
	0x65, 0x4d, 0x6f, 0x6e, 0x67, 0x6f, 0x46, 0x6f, 0x72, 0x49, 0x61, 0x6d, 0x53, 0x63, 0x6f, 0x70,
	0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x1a, 0x21, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x73, 0x2f, 0x2d, 0x1a, 0x2b, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x7d, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x2f,
	0x2d, 0x2a, 0x05, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x0a, 0x99, 0x01, 0x0a, 0x07, 0x62, 0x79, 0x53,
	0x63, 0x6f, 0x70, 0x65, 0x12, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x1a, 0x21, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x72,
	0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x2d, 0x1a, 0x2b, 0x6f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6f, 0x72,
	0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x72, 0x6f, 0x6c, 0x65,
	0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x2d, 0x1a, 0x21, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2f, 0x7b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x7d, 0x2f, 0x72,
	0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x2d, 0x1a, 0x0e, 0x72,
	0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x2d, 0x2a, 0x05, 0x6d,
	0x6f, 0x6e, 0x67, 0x6f, 0x12, 0x14, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x10, 0x01, 0x1a, 0x0a,
	0x62, 0x79, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x46, 0x73, 0x12, 0x53, 0x0a, 0x0d, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x5f, 0x62, 0x79, 0x5f, 0x6f, 0x72, 0x67, 0x10, 0x01, 0x1a, 0x20, 0x62,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x4d, 0x6f, 0x6e, 0x67, 0x6f, 0x46, 0x6f, 0x72, 0x49, 0x61, 0x6d,
	0x53, 0x63, 0x6f, 0x70, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x22,
	0x1e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2d,
	0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x2d, 0x12,
	0x13, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x1a, 0x07, 0x62, 0x79, 0x53,
	0x63, 0x6f, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x0f,
	0x0a, 0x0d, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x62, 0x79, 0x5f, 0x6f, 0x72, 0x67, 0xda,
	0x94, 0x23, 0x18, 0x12, 0x06, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0xe2, 0xde, 0x21, 0x02, 0x08,
	0x01, 0xc2, 0x85, 0x2c, 0xc5, 0x01, 0x22, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x04, 0x72, 0x6f,
	0x6c, 0x65, 0x22, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x0b, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x22, 0x15, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x63, 0x6f,
	0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x0c, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x5f,
	0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x0d, 0x6f, 0x77, 0x6e, 0x65, 0x64, 0x5f, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x22, 0x18, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x66,
	0x6f, 0x72, 0x5f, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x5f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x2a,
	0x0d, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x74, 0x72, 0x79, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x2a, 0x0d,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x62, 0x79, 0x5f, 0x6f, 0x72, 0x67, 0x42, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0x0f, 0x73, 0x70, 0x65, 0x63, 0x5f, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x11, 0x68, 0x61, 0x73, 0x5f, 0x6f, 0x77,
	0x6e, 0x65, 0x64, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x42, 0x92, 0x02, 0xe8, 0xde,
	0x21, 0x01, 0xd2, 0xff, 0xd0, 0x02, 0x4a, 0x0a, 0x12, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x62, 0x69,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x34, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e,
	0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0xa2, 0x80, 0xd1, 0x02, 0x4c, 0x0a, 0x13, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x62, 0x69, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x35, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e,
	0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e,
	0x70, 0x62, 0x2e, 0x76, 0x31, 0x42, 0x10, 0x52, 0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65,
	0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x62, 0x69, 0x6e, 0x64,
	0x69, 0x6e, 0x67, 0x3b, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	edgelq_iam_proto_v1_role_binding_proto_rawDescOnce sync.Once
	edgelq_iam_proto_v1_role_binding_proto_rawDescData = edgelq_iam_proto_v1_role_binding_proto_rawDesc
)

func edgelq_iam_proto_v1_role_binding_proto_rawDescGZIP() []byte {
	edgelq_iam_proto_v1_role_binding_proto_rawDescOnce.Do(func() {
		edgelq_iam_proto_v1_role_binding_proto_rawDescData = protoimpl.X.CompressGZIP(edgelq_iam_proto_v1_role_binding_proto_rawDescData)
	})
	return edgelq_iam_proto_v1_role_binding_proto_rawDescData
}

var edgelq_iam_proto_v1_role_binding_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var edgelq_iam_proto_v1_role_binding_proto_goTypes = []interface{}{
	(*RoleBinding)(nil),                   // 0: ntt.iam.v1.RoleBinding
	(*RoleBinding_Parent)(nil),            // 1: ntt.iam.v1.RoleBinding.Parent
	(*meta.Meta)(nil),                     // 2: goten.types.Meta
	(*role.ScopeParam)(nil),               // 3: ntt.iam.v1.ScopeParam
	(*condition.ExecutableCondition)(nil), // 4: ntt.iam.v1.ExecutableCondition
	(role.Role_Category)(0),               // 5: ntt.iam.v1.Role_Category
}
var edgelq_iam_proto_v1_role_binding_proto_depIdxs = []int32{
	2, // 0: ntt.iam.v1.RoleBinding.metadata:type_name -> goten.types.Meta
	3, // 1: ntt.iam.v1.RoleBinding.scope_params:type_name -> ntt.iam.v1.ScopeParam
	4, // 2: ntt.iam.v1.RoleBinding.executable_conditions:type_name -> ntt.iam.v1.ExecutableCondition
	5, // 3: ntt.iam.v1.RoleBinding.category:type_name -> ntt.iam.v1.Role_Category
	1, // 4: ntt.iam.v1.RoleBinding.ancestry_path:type_name -> ntt.iam.v1.RoleBinding.Parent
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { edgelq_iam_proto_v1_role_binding_proto_init() }
func edgelq_iam_proto_v1_role_binding_proto_init() {
	if edgelq_iam_proto_v1_role_binding_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		edgelq_iam_proto_v1_role_binding_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleBinding); i {
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
		edgelq_iam_proto_v1_role_binding_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleBinding_Parent); i {
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
			RawDescriptor: edgelq_iam_proto_v1_role_binding_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           edgelq_iam_proto_v1_role_binding_proto_goTypes,
		DependencyIndexes: edgelq_iam_proto_v1_role_binding_proto_depIdxs,
		MessageInfos:      edgelq_iam_proto_v1_role_binding_proto_msgTypes,
	}.Build()
	edgelq_iam_proto_v1_role_binding_proto = out.File
	edgelq_iam_proto_v1_role_binding_proto_rawDesc = nil
	edgelq_iam_proto_v1_role_binding_proto_goTypes = nil
	edgelq_iam_proto_v1_role_binding_proto_depIdxs = nil
}
