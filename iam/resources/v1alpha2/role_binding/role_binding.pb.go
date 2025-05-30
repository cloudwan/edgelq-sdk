// Code generated by protoc-gen-goten-go
// File: edgelq/iam/proto/v1alpha2/role_binding.proto
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
	condition "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/condition"
	organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/organization"
	project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/project"
	role "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/role"
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
	_ = &meta.Meta{}
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// RoleBinding in iam RBAC model is the way of granting access to a party (user,
// service account, etc) to edgelq resources. Creating RoleBinding requires user
// to answer 3 questions:
//
// Who: with member field, e.g. "user:wile.e.coyote@customers.acme.com"
// What: specify scope (or parent), e.g.: `projects/acme/roleBindings/<uuid>`
// How: bind role and optional condition to grant access to resources within
// above scope
//
// RoleBindings are additive, meaning that creating a new RoleBinding may only
// extend ability of given member to perform actions. In other words RoleBinding
// doesn't affect other RoleBindings and
// [PermissionCheck][ntt.iam.v1alpha2.PermissionCheck] method needs to find
// *any* RoleBinding granting permission
type RoleBinding struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Name of RoleBinding
	// When creating a new instance, this field is optional and if not provided,
	// it will be generated automatically. Last ID segment must conform to the
	// following regex: [\\w.|-]{1,128}
	Name *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty"`
	// Role
	Role *role.Reference `protobuf:"bytes,2,opt,customtype=Reference,name=role,proto3" json:"role,omitempty"`
	// Binding members
	// Format of the string is one of:
	// - "allUsers" (anyone)
	// - "allAuthenticatedUsers" (anyone logged in)
	// - "user:admin.super@example.com"
	// - "serviceAccount:device_agent@watchdog.serviceaccounts.iam.edgelq.com"
	// - "group:nice.group@example.com"
	// - "domain:example.com" (anyone with exact email domain)
	Member string `protobuf:"bytes,4,opt,name=member,proto3" json:"member,omitempty"`
	// optional ConditionBinding
	// TODO: Make it repeated and make sure backend will check all before
	// verifying
	ConditionBinding *condition.ConditionBinding `protobuf:"bytes,6,opt,name=condition_binding,json=conditionBinding,proto3" json:"condition_binding,omitempty"`
	// Internal field used by IAM controller to note role binding ancestry path
	AncestryPath []*RoleBinding_Parent `protobuf:"bytes,7,rep,name=ancestry_path,json=ancestryPath,proto3" json:"ancestry_path,omitempty"`
	// Metadata
	Metadata *meta.Meta `protobuf:"bytes,8,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (m *RoleBinding) Reset() {
	*m = RoleBinding{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_iam_proto_v1alpha2_role_binding_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *RoleBinding) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*RoleBinding) ProtoMessage() {}

func (m *RoleBinding) ProtoReflect() preflect.Message {
	mi := &edgelq_iam_proto_v1alpha2_role_binding_proto_msgTypes[0]
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
	return edgelq_iam_proto_v1alpha2_role_binding_proto_rawDescGZIP(), []int{0}
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

func (m *RoleBinding) GetConditionBinding() *condition.ConditionBinding {
	if m != nil {
		return m.ConditionBinding
	}
	return nil
}

func (m *RoleBinding) GetAncestryPath() []*RoleBinding_Parent {
	if m != nil {
		return m.AncestryPath
	}
	return nil
}

func (m *RoleBinding) GetMetadata() *meta.Meta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *RoleBinding) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "RoleBinding"))
	}
	m.Name = fv
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

func (m *RoleBinding) SetConditionBinding(fv *condition.ConditionBinding) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ConditionBinding", "RoleBinding"))
	}
	m.ConditionBinding = fv
}

func (m *RoleBinding) SetAncestryPath(fv []*RoleBinding_Parent) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "AncestryPath", "RoleBinding"))
	}
	m.AncestryPath = fv
}

func (m *RoleBinding) SetMetadata(fv *meta.Meta) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Metadata", "RoleBinding"))
	}
	m.Metadata = fv
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
	Parent *Reference `protobuf:"bytes,1,opt,customtype=Reference,name=parent,proto3" json:"parent,omitempty"`
	// Member of the parent role binding
	Member string `protobuf:"bytes,2,opt,name=member,proto3" json:"member,omitempty"`
}

func (m *RoleBinding_Parent) Reset() {
	*m = RoleBinding_Parent{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_iam_proto_v1alpha2_role_binding_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *RoleBinding_Parent) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*RoleBinding_Parent) ProtoMessage() {}

func (m *RoleBinding_Parent) ProtoReflect() preflect.Message {
	mi := &edgelq_iam_proto_v1alpha2_role_binding_proto_msgTypes[1]
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
	return edgelq_iam_proto_v1alpha2_role_binding_proto_rawDescGZIP(), []int{0, 0}
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

var edgelq_iam_proto_v1alpha2_role_binding_proto preflect.FileDescriptor

var edgelq_iam_proto_v1alpha2_role_binding_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x72, 0x6f, 0x6c, 0x65,
	0x5f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10,
	0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32,
	0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x74,
	0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67,
	0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f,
	0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d,
	0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24,
	0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x61, 0x75, 0x64,
	0x69, 0x74, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61,
	0x75, 0x64, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x67, 0x6f, 0x74, 0x65,
	0x6e, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x29, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x63, 0x6f,
	0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbb, 0x08,
	0x0a, 0x0b, 0x52, 0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x27, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x13, 0xb2, 0xda, 0x21,
	0x0f, 0x0a, 0x0d, 0x0a, 0x0b, 0x52, 0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2b, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x17, 0xfa, 0x41, 0x06, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0xb2,
	0xda, 0x21, 0x0a, 0x12, 0x08, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x10, 0x01, 0x52, 0x04, 0x72,
	0x6f, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x4f, 0x0a, 0x11, 0x63,
	0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x10, 0x63, 0x6f, 0x6e, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x49, 0x0a, 0x0d,
	0x61, 0x6e, 0x63, 0x65, 0x73, 0x74, 0x72, 0x79, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x07, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x2e, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x52, 0x0c, 0x61, 0x6e, 0x63, 0x65, 0x73,
	0x74, 0x72, 0x79, 0x50, 0x61, 0x74, 0x68, 0x12, 0x2d, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67, 0x6f, 0x74, 0x65,
	0x6e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x5f, 0x0a, 0x06, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x12, 0x3d, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x25, 0xfa, 0x41, 0x0d, 0x0a, 0x0b, 0x52, 0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0xb2, 0xda, 0x21, 0x11, 0x12, 0x0f, 0x0a, 0x0b, 0x52, 0x6f, 0x6c, 0x65, 0x42, 0x69,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x10, 0x06, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x3a, 0x8b, 0x05, 0xea, 0x41, 0xa3, 0x01, 0x0a, 0x1a,
	0x69, 0x61, 0x6d, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x52,
	0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x1b, 0x72, 0x6f, 0x6c, 0x65,
	0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x7b, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x62,
	0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x7d, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x72, 0x6f, 0x6c, 0x65,
	0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x7b, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x62,
	0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x7d, 0x12, 0x38, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x73, 0x2f, 0x7b, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x7d, 0x92, 0xd9, 0x21, 0x77, 0x0a, 0x0c, 0x72, 0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x73, 0x12, 0x0c, 0x72, 0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x73, 0x1a, 0x12, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x4e, 0x6f, 0x6e, 0x65, 0x1a, 0x07, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x1a, 0x0c,
	0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2a, 0x0e, 0x5b, 0x5c,
	0x77, 0x2e, 0x7c, 0x2d, 0x5d, 0x7b, 0x31, 0x2c, 0x31, 0x32, 0x38, 0x7d, 0x38, 0x05, 0x42, 0x1c,
	0x08, 0x02, 0x12, 0x06, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x06, 0x0a, 0x04, 0x72, 0x6f,
	0x6c, 0x65, 0x12, 0x08, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0xb2, 0xdf, 0x21, 0xfe,
	0x01, 0x0a, 0xf1, 0x01, 0x0a, 0x42, 0x0a, 0x0a, 0x62, 0x79, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x46, 0x73, 0x12, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x1a, 0x21, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x72,
	0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x2d, 0x2a, 0x09, 0x66,
	0x69, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x0a, 0x7e, 0x0a, 0x0d, 0x62, 0x79, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x4d, 0x6f, 0x6e, 0x67, 0x6f, 0x12, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x1a, 0x21, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x73, 0x2f, 0x2d, 0x1a, 0x2b, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x7d, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x2f,
	0x2d, 0x1a, 0x0e, 0x72, 0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x2f,
	0x2d, 0x2a, 0x05, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x12, 0x14, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65,
	0x10, 0x01, 0x1a, 0x0a, 0x62, 0x79, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x46, 0x73, 0x12, 0x15,
	0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x1a, 0x0d, 0x62, 0x79, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x4d, 0x6f, 0x6e, 0x67, 0x6f, 0x12, 0x08, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0xda,
	0x94, 0x23, 0x18, 0x12, 0x06, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0xe2, 0xde, 0x21, 0x02, 0x08,
	0x01, 0xc2, 0x85, 0x2c, 0x40, 0x22, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x04, 0x72, 0x6f, 0x6c,
	0x65, 0x22, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x11, 0x63, 0x6f, 0x6e, 0x64, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x2a, 0x0d, 0x61, 0x6e,
	0x63, 0x65, 0x73, 0x74, 0x72, 0x79, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x42, 0x08, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x4a, 0x04, 0x08, 0x05, 0x10, 0x06, 0x42, 0xaa, 0x02, 0xe8, 0xde,
	0x21, 0x01, 0xd2, 0xff, 0xd0, 0x02, 0x50, 0x0a, 0x12, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x62, 0x69,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x3a, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e,
	0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x5f,
	0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0xa2, 0x80, 0xd1, 0x02, 0x52, 0x0a, 0x13, 0x72, 0x6f,
	0x6c, 0x65, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x12, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61,
	0x6d, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x32, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x0a, 0x17,
	0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x70, 0x62, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x42, 0x10, 0x52, 0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4b, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e,
	0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x72,
	0x6f, 0x6c, 0x65, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x3b, 0x72, 0x6f, 0x6c, 0x65,
	0x5f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	edgelq_iam_proto_v1alpha2_role_binding_proto_rawDescOnce sync.Once
	edgelq_iam_proto_v1alpha2_role_binding_proto_rawDescData = edgelq_iam_proto_v1alpha2_role_binding_proto_rawDesc
)

func edgelq_iam_proto_v1alpha2_role_binding_proto_rawDescGZIP() []byte {
	edgelq_iam_proto_v1alpha2_role_binding_proto_rawDescOnce.Do(func() {
		edgelq_iam_proto_v1alpha2_role_binding_proto_rawDescData = protoimpl.X.CompressGZIP(edgelq_iam_proto_v1alpha2_role_binding_proto_rawDescData)
	})
	return edgelq_iam_proto_v1alpha2_role_binding_proto_rawDescData
}

var edgelq_iam_proto_v1alpha2_role_binding_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var edgelq_iam_proto_v1alpha2_role_binding_proto_goTypes = []interface{}{
	(*RoleBinding)(nil),                // 0: ntt.iam.v1alpha2.RoleBinding
	(*RoleBinding_Parent)(nil),         // 1: ntt.iam.v1alpha2.RoleBinding.Parent
	(*condition.ConditionBinding)(nil), // 2: ntt.iam.v1alpha2.ConditionBinding
	(*meta.Meta)(nil),                  // 3: goten.types.Meta
}
var edgelq_iam_proto_v1alpha2_role_binding_proto_depIdxs = []int32{
	2, // 0: ntt.iam.v1alpha2.RoleBinding.condition_binding:type_name -> ntt.iam.v1alpha2.ConditionBinding
	1, // 1: ntt.iam.v1alpha2.RoleBinding.ancestry_path:type_name -> ntt.iam.v1alpha2.RoleBinding.Parent
	3, // 2: ntt.iam.v1alpha2.RoleBinding.metadata:type_name -> goten.types.Meta
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { edgelq_iam_proto_v1alpha2_role_binding_proto_init() }
func edgelq_iam_proto_v1alpha2_role_binding_proto_init() {
	if edgelq_iam_proto_v1alpha2_role_binding_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		edgelq_iam_proto_v1alpha2_role_binding_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		edgelq_iam_proto_v1alpha2_role_binding_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: edgelq_iam_proto_v1alpha2_role_binding_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           edgelq_iam_proto_v1alpha2_role_binding_proto_goTypes,
		DependencyIndexes: edgelq_iam_proto_v1alpha2_role_binding_proto_depIdxs,
		MessageInfos:      edgelq_iam_proto_v1alpha2_role_binding_proto_msgTypes,
	}.Build()
	edgelq_iam_proto_v1alpha2_role_binding_proto = out.File
	edgelq_iam_proto_v1alpha2_role_binding_proto_rawDesc = nil
	edgelq_iam_proto_v1alpha2_role_binding_proto_goTypes = nil
	edgelq_iam_proto_v1alpha2_role_binding_proto_depIdxs = nil
}
