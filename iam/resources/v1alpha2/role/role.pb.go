// Code generated by protoc-gen-goten-go
// File: edgelq/iam/proto/v1alpha2/role.proto
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
	condition "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/condition"
	permission "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/permission"
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
	_ = &permission.Permission{}
	_ = &meta.Meta{}
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Role Resource
type Role struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Name of Role
	Name *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty"`
	// Display Name
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Included Permissions in this Role. Binding this role grants all following
	// Permissions.
	IncludedPermissions []*permission.Reference `protobuf:"bytes,3,rep,customtype=Reference,name=included_permissions,json=includedPermissions,proto3" json:"included_permissions,omitempty"`
	// Default conditionBinding (optional), cannot be used with required
	// TODO: Deprecated...
	DefaultConditionBinding *condition.ConditionBinding `protobuf:"bytes,5,opt,name=default_condition_binding,json=defaultConditionBinding,proto3" json:"default_condition_binding,omitempty"`
	// Condition bindings that will be copied into created role bindings
	// TODO: Works only if only one element is specified
	IncludedConditionBindings []*condition.ConditionBinding `protobuf:"bytes,6,rep,name=included_condition_bindings,json=includedConditionBindings,proto3" json:"included_condition_bindings,omitempty"`
	// List of conditions that must be used for this role. Parameters must be
	// defined by client.
	// TODO: Works only if only one element is specified
	RequiredConditions []*condition.Reference `protobuf:"bytes,7,rep,customtype=Reference,name=required_conditions,json=requiredConditions,proto3" json:"required_conditions,omitempty"`
	// Metadata
	Metadata *meta.Meta `protobuf:"bytes,4,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (m *Role) Reset() {
	*m = Role{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_iam_proto_v1alpha2_role_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *Role) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*Role) ProtoMessage() {}

func (m *Role) ProtoReflect() preflect.Message {
	mi := &edgelq_iam_proto_v1alpha2_role_proto_msgTypes[0]
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
	return edgelq_iam_proto_v1alpha2_role_proto_rawDescGZIP(), []int{0}
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

func (m *Role) GetIncludedConditionBindings() []*condition.ConditionBinding {
	if m != nil {
		return m.IncludedConditionBindings
	}
	return nil
}

func (m *Role) GetRequiredConditions() []*condition.Reference {
	if m != nil {
		return m.RequiredConditions
	}
	return nil
}

func (m *Role) GetMetadata() *meta.Meta {
	if m != nil {
		return m.Metadata
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

func (m *Role) SetIncludedConditionBindings(fv []*condition.ConditionBinding) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "IncludedConditionBindings", "Role"))
	}
	m.IncludedConditionBindings = fv
}

func (m *Role) SetRequiredConditions(fv []*condition.Reference) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "RequiredConditions", "Role"))
	}
	m.RequiredConditions = fv
}

func (m *Role) SetMetadata(fv *meta.Meta) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Metadata", "Role"))
	}
	m.Metadata = fv
}

var edgelq_iam_proto_v1alpha2_role_proto preflect.FileDescriptor

var edgelq_iam_proto_v1alpha2_role_proto_rawDesc = []byte{
	0x0a, 0x24, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x72, 0x6f, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x61,
	0x75, 0x64, 0x69, 0x74, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x67, 0x6f,
	0x74, 0x65, 0x6e, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f,
	0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x84, 0x06, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xb2, 0xda, 0x21, 0x08, 0x0a, 0x06, 0x0a, 0x04,
	0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69,
	0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x56, 0x0a,
	0x14, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x5f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x42, 0x23, 0xfa, 0x41, 0x0c,
	0x0a, 0x0a, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0xb2, 0xda, 0x21, 0x10,
	0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x10, 0x06,
	0x52, 0x13, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x5e, 0x0a, 0x19, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74,
	0x5f, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69,
	0x61, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x43, 0x6f, 0x6e, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x17, 0x64, 0x65,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x69,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x62, 0x0a, 0x1b, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65,
	0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x62, 0x69, 0x6e, 0x64,
	0x69, 0x6e, 0x67, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6e, 0x74, 0x74,
	0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x43, 0x6f,
	0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x19,
	0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x52, 0x0a, 0x13, 0x72, 0x65, 0x71,
	0x75, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x42, 0x21, 0xfa, 0x41, 0x0b, 0x0a, 0x09, 0x43, 0x6f, 0x6e,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0xb2, 0xda, 0x21, 0x0f, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x6f,
	0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x01, 0x52, 0x12, 0x72, 0x65, 0x71, 0x75, 0x69,
	0x72, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2d, 0x0a,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4d, 0x65,
	0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x3a, 0x97, 0x02, 0xea,
	0x41, 0x23, 0x0a, 0x13, 0x69, 0x61, 0x6d, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x0c, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x2f, 0x7b,
	0x72, 0x6f, 0x6c, 0x65, 0x7d, 0x92, 0xd9, 0x21, 0x3f, 0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73,
	0x12, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x2a, 0x13, 0x5b, 0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a,
	0x30, 0x2d, 0x39, 0x2d, 0x5d, 0x7b, 0x31, 0x2c, 0x31, 0x32, 0x38, 0x7d, 0x42, 0x1a, 0x08, 0x02,
	0x12, 0x06, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70,
	0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0xb2, 0xdf, 0x21, 0x18, 0x12, 0x16, 0x0a, 0x14,
	0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x5f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0xda, 0x94, 0x23, 0x08, 0x12, 0x06, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0xc2, 0x85, 0x2c, 0x81, 0x01, 0x22, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x0c, 0x64, 0x69, 0x73,
	0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x14, 0x69, 0x6e, 0x63, 0x6c, 0x75,
	0x64, 0x65, 0x64, 0x5f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22,
	0x13, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x22, 0x1b, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x5f, 0x63,
	0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x73, 0x22, 0x19, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x64, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x42, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0xf3, 0x01, 0xe8, 0xde, 0x21, 0x01, 0xd2, 0xff, 0xd0,
	0x02, 0x40, 0x0a, 0x0a, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x32,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x72, 0x6f,
	0x6c, 0x65, 0xa2, 0x80, 0xd1, 0x02, 0x42, 0x0a, 0x0b, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x12, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71,
	0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x32, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x6e,
	0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x32, 0x42, 0x09, 0x52, 0x6f, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x32, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x3b, 0x72, 0x6f, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	edgelq_iam_proto_v1alpha2_role_proto_rawDescOnce sync.Once
	edgelq_iam_proto_v1alpha2_role_proto_rawDescData = edgelq_iam_proto_v1alpha2_role_proto_rawDesc
)

func edgelq_iam_proto_v1alpha2_role_proto_rawDescGZIP() []byte {
	edgelq_iam_proto_v1alpha2_role_proto_rawDescOnce.Do(func() {
		edgelq_iam_proto_v1alpha2_role_proto_rawDescData = protoimpl.X.CompressGZIP(edgelq_iam_proto_v1alpha2_role_proto_rawDescData)
	})
	return edgelq_iam_proto_v1alpha2_role_proto_rawDescData
}

var edgelq_iam_proto_v1alpha2_role_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var edgelq_iam_proto_v1alpha2_role_proto_goTypes = []interface{}{
	(*Role)(nil),                       // 0: ntt.iam.v1alpha2.Role
	(*condition.ConditionBinding)(nil), // 1: ntt.iam.v1alpha2.ConditionBinding
	(*meta.Meta)(nil),                  // 2: goten.types.Meta
}
var edgelq_iam_proto_v1alpha2_role_proto_depIdxs = []int32{
	1, // 0: ntt.iam.v1alpha2.Role.default_condition_binding:type_name -> ntt.iam.v1alpha2.ConditionBinding
	1, // 1: ntt.iam.v1alpha2.Role.included_condition_bindings:type_name -> ntt.iam.v1alpha2.ConditionBinding
	2, // 2: ntt.iam.v1alpha2.Role.metadata:type_name -> goten.types.Meta
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { edgelq_iam_proto_v1alpha2_role_proto_init() }
func edgelq_iam_proto_v1alpha2_role_proto_init() {
	if edgelq_iam_proto_v1alpha2_role_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		edgelq_iam_proto_v1alpha2_role_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: edgelq_iam_proto_v1alpha2_role_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           edgelq_iam_proto_v1alpha2_role_proto_goTypes,
		DependencyIndexes: edgelq_iam_proto_v1alpha2_role_proto_depIdxs,
		MessageInfos:      edgelq_iam_proto_v1alpha2_role_proto_msgTypes,
	}.Build()
	edgelq_iam_proto_v1alpha2_role_proto = out.File
	edgelq_iam_proto_v1alpha2_role_proto_rawDesc = nil
	edgelq_iam_proto_v1alpha2_role_proto_goTypes = nil
	edgelq_iam_proto_v1alpha2_role_proto_depIdxs = nil
}
