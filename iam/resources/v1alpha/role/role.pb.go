// Code generated by protoc-gen-goten-go
// File: edgelq/iam/proto/v1alpha/role.proto
// DO NOT EDIT!!!

package role

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
	condition "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/condition"
	permission "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/permission"
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
	_ = &permission.Permission{}
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Role is a named bag of [Permissions][ntt.iam.v1alpha.Permission] that can be
// granted on certain scope (system, organization or project) to member (user,
// service account, group, anyone) with [RoleBinding].
//
// Note: currently roles are defined as system-level resources, but
// per-organization or per-project support is expected.
type Role struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Name of Role
	Name *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// Display Name
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty" firestore:"displayName"`
	// Included Permissions in this Role. Binding this role grants all following
	// Permissions.
	IncludedPermissions []*permission.Reference `protobuf:"bytes,3,rep,customtype=Reference,name=included_permissions,json=includedPermissions,proto3" json:"included_permissions,omitempty" firestore:"includedPermissions"`
	// Default conditionBinding (optional)
	DefaultConditionBinding *condition.ConditionBinding `protobuf:"bytes,5,opt,name=default_condition_binding,json=defaultConditionBinding,proto3" json:"default_condition_binding,omitempty" firestore:"defaultConditionBinding"`
}

func (m *Role) Reset() {
	*m = Role{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_iam_proto_v1alpha_role_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *Role) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*Role) ProtoMessage() {}

func (m *Role) ProtoReflect() preflect.Message {
	mi := &edgelq_iam_proto_v1alpha_role_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*Role) GotenMessage() {}

// Deprecated, Use Role.ProtoReflect.Descriptor instead.
func (*Role) Descriptor() ([]byte, []int) {
	return edgelq_iam_proto_v1alpha_role_proto_rawDescGZIP(), []int{0}
}

func (m *Role) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *Role) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *Role) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *Role) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *Role) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *Role) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *Role) GetIncludedPermissions() []*permission.Reference {
	if m != nil {
		return m.IncludedPermissions
	}
	return nil
}

func (m *Role) GetDefaultConditionBinding() *condition.ConditionBinding {
	if m != nil {
		return m.DefaultConditionBinding
	}
	return nil
}

func (m *Role) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "Role"))
	}
	m.Name = fv
}

func (m *Role) SetDisplayName(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "DisplayName", "Role"))
	}
	m.DisplayName = fv
}

func (m *Role) SetIncludedPermissions(fv []*permission.Reference) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "IncludedPermissions", "Role"))
	}
	m.IncludedPermissions = fv
}

func (m *Role) SetDefaultConditionBinding(fv *condition.ConditionBinding) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "DefaultConditionBinding", "Role"))
	}
	m.DefaultConditionBinding = fv
}

var edgelq_iam_proto_v1alpha_role_proto preflect.FileDescriptor

var edgelq_iam_proto_v1alpha_role_proto_rawDesc = []byte{
	0x0a, 0x23, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f,
	0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x65,
	0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb7, 0x03, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65,
	0x12, 0x20, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c,
	0xb2, 0xda, 0x21, 0x08, 0x0a, 0x06, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x29, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xde, 0x21, 0x02, 0x08, 0x04,
	0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x54, 0x0a,
	0x14, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x5f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x42, 0x21, 0xfa, 0x41, 0x0c,
	0x0a, 0x0a, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0xb2, 0xda, 0x21, 0x0e,
	0x12, 0x0c, 0x0a, 0x0a, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x13,
	0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x5d, 0x0a, 0x19, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x63,
	0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x17, 0x64, 0x65, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x69, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x3a, 0xac, 0x01, 0xea, 0x41, 0x23, 0x0a, 0x13, 0x69, 0x61, 0x6d, 0x2e, 0x65, 0x64,
	0x67, 0x65, 0x6c, 0x71, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x0c, 0x72,
	0x6f, 0x6c, 0x65, 0x73, 0x2f, 0x7b, 0x72, 0x6f, 0x6c, 0x65, 0x7d, 0x92, 0xd9, 0x21, 0x30, 0x0a,
	0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x22, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x4a, 0x1a, 0x08, 0x02, 0x12, 0x06, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x0e, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0xfa,
	0xde, 0x21, 0x06, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0xb2, 0xdf, 0x21, 0x0a, 0x0a, 0x08, 0x6c,
	0x69, 0x73, 0x74, 0x56, 0x69, 0x65, 0x77, 0xda, 0x94, 0x23, 0x08, 0x12, 0x06, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0xc2, 0x85, 0x2c, 0x2a, 0x22, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x0c, 0x64,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x14, 0x69, 0x6e, 0x63,
	0x6c, 0x75, 0x64, 0x65, 0x64, 0x5f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x42, 0x87, 0x03, 0xe8, 0xde, 0x21, 0x01, 0xd2, 0xff, 0xd0, 0x02, 0x3f, 0x0a, 0x0a, 0x72,
	0x6f, 0x6c, 0x65, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65,
	0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x92, 0x8c, 0xd1, 0x02,
	0x49, 0x0a, 0x0f, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69,
	0x6e, 0x67, 0x12, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69,
	0x61, 0x6d, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e,
	0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x42, 0x09, 0x52, 0x6f, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x3b, 0x72, 0x6f, 0x6c, 0x65, 0xd2, 0x84, 0xd1, 0x02, 0x45,
	0x0a, 0x0d, 0x61, 0x75, 0x64, 0x69, 0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x12,
	0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f,
	0x61, 0x75, 0x64, 0x69, 0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0xa2, 0x80, 0xd1, 0x02, 0x41, 0x0a, 0x0b, 0x72, 0x6f, 0x6c, 0x65,
	0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67,
	0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	edgelq_iam_proto_v1alpha_role_proto_rawDescOnce sync.Once
	edgelq_iam_proto_v1alpha_role_proto_rawDescData = edgelq_iam_proto_v1alpha_role_proto_rawDesc
)

func edgelq_iam_proto_v1alpha_role_proto_rawDescGZIP() []byte {
	edgelq_iam_proto_v1alpha_role_proto_rawDescOnce.Do(func() {
		edgelq_iam_proto_v1alpha_role_proto_rawDescData = protoimpl.X.CompressGZIP(edgelq_iam_proto_v1alpha_role_proto_rawDescData)
	})
	return edgelq_iam_proto_v1alpha_role_proto_rawDescData
}

var edgelq_iam_proto_v1alpha_role_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var edgelq_iam_proto_v1alpha_role_proto_goTypes = []interface{}{
	(*Role)(nil),                       // 0: ntt.iam.v1alpha.Role
	(*condition.ConditionBinding)(nil), // 1: ntt.iam.v1alpha.ConditionBinding
}
var edgelq_iam_proto_v1alpha_role_proto_depIdxs = []int32{
	1, // 0: ntt.iam.v1alpha.Role.default_condition_binding:type_name -> ntt.iam.v1alpha.ConditionBinding
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { edgelq_iam_proto_v1alpha_role_proto_init() }
func edgelq_iam_proto_v1alpha_role_proto_init() {
	if edgelq_iam_proto_v1alpha_role_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		edgelq_iam_proto_v1alpha_role_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Role); i {
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
			RawDescriptor: edgelq_iam_proto_v1alpha_role_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           edgelq_iam_proto_v1alpha_role_proto_goTypes,
		DependencyIndexes: edgelq_iam_proto_v1alpha_role_proto_depIdxs,
		MessageInfos:      edgelq_iam_proto_v1alpha_role_proto_msgTypes,
	}.Build()
	edgelq_iam_proto_v1alpha_role_proto = out.File
	edgelq_iam_proto_v1alpha_role_proto_rawDesc = nil
	edgelq_iam_proto_v1alpha_role_proto_goTypes = nil
	edgelq_iam_proto_v1alpha_role_proto_depIdxs = nil
}
