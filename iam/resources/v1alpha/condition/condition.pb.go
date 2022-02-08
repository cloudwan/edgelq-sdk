// Code generated by protoc-gen-goten-go
// File: edgelq/iam/proto/v1alpha/condition.proto
// DO NOT EDIT!!!

package condition

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
	organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/organization"
	project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/project"
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
	_ = &organization.Organization{}
	_ = &project.Project{}
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Parameter Type. Note: currently only STRING type is supported.
type Condition_ParameterType int32

const (
	Condition_TYPE_UNSPECIFIED Condition_ParameterType = 0
	Condition_STRING           Condition_ParameterType = 1
	Condition_INT64            Condition_ParameterType = 2
	Condition_DOUBLE           Condition_ParameterType = 3
	Condition_BOOL             Condition_ParameterType = 4
)

var (
	Condition_ParameterType_name = map[int32]string{
		0: "TYPE_UNSPECIFIED",
		1: "STRING",
		2: "INT64",
		3: "DOUBLE",
		4: "BOOL",
	}

	Condition_ParameterType_value = map[string]int32{
		"TYPE_UNSPECIFIED": 0,
		"STRING":           1,
		"INT64":            2,
		"DOUBLE":           3,
		"BOOL":             4,
	}
)

func (x Condition_ParameterType) Enum() *Condition_ParameterType {
	p := new(Condition_ParameterType)
	*p = x
	return p
}

func (x Condition_ParameterType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), preflect.EnumNumber(x))
}

func (Condition_ParameterType) Descriptor() preflect.EnumDescriptor {
	return edgelq_iam_proto_v1alpha_condition_proto_enumTypes[0].Descriptor()
}

func (Condition_ParameterType) Type() preflect.EnumType {
	return &edgelq_iam_proto_v1alpha_condition_proto_enumTypes[0]
}

func (x Condition_ParameterType) Number() preflect.EnumNumber {
	return preflect.EnumNumber(x)
}

// Deprecated, Use Condition_ParameterType.ProtoReflect.Descriptor instead.
func (Condition_ParameterType) EnumDescriptor() ([]byte, []int) {
	return edgelq_iam_proto_v1alpha_condition_proto_rawDescGZIP(), []int{0, 0}
}

// Condition Resource provides an extension to primary RBAC model, which allows
// customizable (scriptable) access. Conditions can only be used to further
// limit (or narrow) access compared to standard (or unconditional) RBAC.
//
// Usage:
//
// 1. Condition is created with expression and parameter declarations.
//    Consider this a definition, which will be later "bound".
// 2. When creating a [RoleBinding] pass [ConditionBinding] with reference to
// [Condition] and bound parameters.
//
// When checking for principal access (anonymous, user, service account), for
// RoleBinding to grant permissions included in bound [Role], both RBAC and
// Condition expression evaluation must grant access.
type Condition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Name of Condition
	Name *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// Display Name
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty" firestore:"displayName"`
	// Description
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty" firestore:"description"`
	// Condition expression in [Google
	// CEL](https://github.com/google/cel-spec/blob/v0.4.0/doc/intro.md), syntax,
	// e.g. `resource.name == "projects/xyz/instances/abc"`
	//
	// Accessible parameters:
	// | variable | type | description | examples |
	// |-|-|-|-|
	// | `resource.body` | `string` | Resource type | `"iam.edgelq.com/Condition"`
	// | | `request.action` | `string` | Action verb | `"create"`, `"batchGet"` |
	// | `request.body` | `dyn` | Request body (in native format). |
	// request.body.page_size | | `parameters` | `map(string, dyn)` | Bound
	// parameters. | `parameters.minSeverity` |
	//
	// Extension:
	//
	// Some
	//
	// Filter.`satisfies(other)`
	//
	// Access is determined by the return value. Return `true` to grant access or
	// `false` to deny. Any execution error results in access denied.
	Expression string `protobuf:"bytes,4,opt,name=expression,proto3" json:"expression,omitempty" firestore:"expression"`
	// Typed parameters declarations. When binding a Condition passed parameters
	// must correspond to declarations.
	ParameterDeclarations []*Condition_ParameterDeclaration `protobuf:"bytes,5,rep,name=parameter_declarations,json=parameterDeclarations,proto3" json:"parameter_declarations,omitempty" firestore:"parameterDeclarations"`
}

func (m *Condition) Reset() {
	*m = Condition{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_iam_proto_v1alpha_condition_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *Condition) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*Condition) ProtoMessage() {}

func (m *Condition) ProtoReflect() preflect.Message {
	mi := &edgelq_iam_proto_v1alpha_condition_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*Condition) GotenMessage() {}

// Deprecated, Use Condition.ProtoReflect.Descriptor instead.
func (*Condition) Descriptor() ([]byte, []int) {
	return edgelq_iam_proto_v1alpha_condition_proto_rawDescGZIP(), []int{0}
}

func (m *Condition) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *Condition) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *Condition) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *Condition) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *Condition) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *Condition) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *Condition) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Condition) GetExpression() string {
	if m != nil {
		return m.Expression
	}
	return ""
}

func (m *Condition) GetParameterDeclarations() []*Condition_ParameterDeclaration {
	if m != nil {
		return m.ParameterDeclarations
	}
	return nil
}

func (m *Condition) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "Condition"))
	}
	m.Name = fv
}

func (m *Condition) SetDisplayName(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "DisplayName", "Condition"))
	}
	m.DisplayName = fv
}

func (m *Condition) SetDescription(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Description", "Condition"))
	}
	m.Description = fv
}

func (m *Condition) SetExpression(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Expression", "Condition"))
	}
	m.Expression = fv
}

func (m *Condition) SetParameterDeclarations(fv []*Condition_ParameterDeclaration) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ParameterDeclarations", "Condition"))
	}
	m.ParameterDeclarations = fv
}

// ConditionBinding represents instantiantion of condi
type ConditionBinding struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Reference to Condition which may also be parameterized
	Condition *Reference `protobuf:"bytes,1,opt,customtype=Reference,name=condition,proto3" json:"condition,omitempty" firestore:"condition"`
	// Parameters in string form. Parameters must match
	// [declarations][ntt.iam.v1alpha.Condition.parameter_declarations]
	Parameters map[string]string `protobuf:"bytes,2,rep,name=parameters,proto3" json:"parameters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" firestore:"parameters"`
}

func (m *ConditionBinding) Reset() {
	*m = ConditionBinding{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_iam_proto_v1alpha_condition_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *ConditionBinding) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*ConditionBinding) ProtoMessage() {}

func (m *ConditionBinding) ProtoReflect() preflect.Message {
	mi := &edgelq_iam_proto_v1alpha_condition_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*ConditionBinding) GotenMessage() {}

// Deprecated, Use ConditionBinding.ProtoReflect.Descriptor instead.
func (*ConditionBinding) Descriptor() ([]byte, []int) {
	return edgelq_iam_proto_v1alpha_condition_proto_rawDescGZIP(), []int{1}
}

func (m *ConditionBinding) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *ConditionBinding) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *ConditionBinding) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *ConditionBinding) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *ConditionBinding) GetCondition() *Reference {
	if m != nil {
		return m.Condition
	}
	return nil
}

func (m *ConditionBinding) GetParameters() map[string]string {
	if m != nil {
		return m.Parameters
	}
	return nil
}

func (m *ConditionBinding) SetCondition(fv *Reference) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Condition", "ConditionBinding"))
	}
	m.Condition = fv
}

func (m *ConditionBinding) SetParameters(fv map[string]string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Parameters", "ConditionBinding"))
	}
	m.Parameters = fv
}

// Parameter Declarations used
type Condition_ParameterDeclaration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Parameter Key - must be unique within condition. Defined parameter
	// variables are accessible in condition expression via `parameters.<key>`,
	// e.g.`parameters.projectId`
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty" firestore:"key"`
	// Parameter value type
	Type Condition_ParameterType `protobuf:"varint,2,opt,name=type,proto3,enum=ntt.iam.v1alpha.Condition_ParameterType" json:"type,omitempty" firestore:"type"`
}

func (m *Condition_ParameterDeclaration) Reset() {
	*m = Condition_ParameterDeclaration{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_iam_proto_v1alpha_condition_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *Condition_ParameterDeclaration) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*Condition_ParameterDeclaration) ProtoMessage() {}

func (m *Condition_ParameterDeclaration) ProtoReflect() preflect.Message {
	mi := &edgelq_iam_proto_v1alpha_condition_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*Condition_ParameterDeclaration) GotenMessage() {}

// Deprecated, Use Condition_ParameterDeclaration.ProtoReflect.Descriptor instead.
func (*Condition_ParameterDeclaration) Descriptor() ([]byte, []int) {
	return edgelq_iam_proto_v1alpha_condition_proto_rawDescGZIP(), []int{0, 0}
}

func (m *Condition_ParameterDeclaration) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *Condition_ParameterDeclaration) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *Condition_ParameterDeclaration) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *Condition_ParameterDeclaration) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *Condition_ParameterDeclaration) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Condition_ParameterDeclaration) GetType() Condition_ParameterType {
	if m != nil {
		return m.Type
	}
	return Condition_TYPE_UNSPECIFIED
}

func (m *Condition_ParameterDeclaration) SetKey(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Key", "Condition_ParameterDeclaration"))
	}
	m.Key = fv
}

func (m *Condition_ParameterDeclaration) SetType(fv Condition_ParameterType) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Type", "Condition_ParameterDeclaration"))
	}
	m.Type = fv
}

var edgelq_iam_proto_v1alpha_condition_proto preflect.FileDescriptor

var edgelq_iam_proto_v1alpha_condition_proto_rawDesc = []byte{
	0x0a, 0x28, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x63, 0x6f, 0x6e, 0x64, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x6e, 0x74, 0x74, 0x2e,
	0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x19, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x74, 0x65,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x65, 0x64,
	0x67, 0x65, 0x6c, 0x71, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x2b, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x6f, 0x72, 0x67,
	0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x26, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x85, 0x06, 0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x11, 0xb2, 0xda, 0x21, 0x0d, 0x0a, 0x0b, 0x0a, 0x09, 0x43, 0x6f, 0x6e,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x66, 0x0a, 0x16, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x5f, 0x64,
	0x65, 0x63, 0x6c, 0x61, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2f, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x44, 0x65, 0x63, 0x6c, 0x61, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x15, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x44, 0x65, 0x63,
	0x6c, 0x61, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x66, 0x0a, 0x14, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x44, 0x65, 0x63, 0x6c, 0x61, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x3c, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x28, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x22, 0x52, 0x0a, 0x0d, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x54, 0x52, 0x49,
	0x4e, 0x47, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x49, 0x4e, 0x54, 0x36, 0x34, 0x10, 0x02, 0x12,
	0x0a, 0x0a, 0x06, 0x44, 0x4f, 0x55, 0x42, 0x4c, 0x45, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x42,
	0x4f, 0x4f, 0x4c, 0x10, 0x04, 0x3a, 0xc7, 0x02, 0xea, 0x41, 0x92, 0x01, 0x0a, 0x18, 0x69, 0x61,
	0x6d, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x6f, 0x6e,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x7b, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x12, 0x29,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x7d, 0x2f, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x63,
	0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x12, 0x33, 0x6f, 0x72, 0x67, 0x61, 0x6e,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x7b, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x92, 0xd9,
	0x21, 0x49, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x0a,
	0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x12, 0x67, 0x6f, 0x74, 0x65,
	0x6e, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x6f, 0x6e, 0x65, 0x1a, 0x07,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x1a, 0x0c, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0xfa, 0xde, 0x21, 0x0b, 0x0a,
	0x09, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0xda, 0x94, 0x23, 0x08, 0x12, 0x06,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0xc2, 0x85, 0x2c, 0x45, 0x22, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x0a, 0x65, 0x78, 0x70,
	0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x16, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74,
	0x65, 0x72, 0x5f, 0x64, 0x65, 0x63, 0x6c, 0x61, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22,
	0xd5, 0x01, 0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x69, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x12, 0x2f, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x11, 0xb2, 0xda, 0x21, 0x0d, 0x12, 0x0b, 0x0a,
	0x09, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x51, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74,
	0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x6e, 0x74, 0x74, 0x2e,
	0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x43, 0x6f, 0x6e, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x70, 0x61,
	0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x1a, 0x3d, 0x0a, 0x0f, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0xb4, 0x03, 0xe8, 0xde, 0x21, 0x01, 0xd2, 0xff,
	0xd0, 0x02, 0x49, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x12, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71,
	0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x2f, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x92, 0x8c, 0xd1, 0x02,
	0x53, 0x0a, 0x14, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67,
	0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69,
	0x6e, 0x67, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x63, 0x6f, 0x6e, 0x64, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61,
	0x6d, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x42, 0x0e, 0x43, 0x6f,
	0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x44,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x2f, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x3b, 0x63, 0x6f, 0x6e, 0x64, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0xd2, 0x84, 0xd1, 0x02, 0x45, 0x0a, 0x0d, 0x61, 0x75, 0x64, 0x69, 0x74,
	0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x12, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64,
	0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x68, 0x61,
	0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0xa2, 0x80,
	0xd1, 0x02, 0x4b, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c,
	0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x2f, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	edgelq_iam_proto_v1alpha_condition_proto_rawDescOnce sync.Once
	edgelq_iam_proto_v1alpha_condition_proto_rawDescData = edgelq_iam_proto_v1alpha_condition_proto_rawDesc
)

func edgelq_iam_proto_v1alpha_condition_proto_rawDescGZIP() []byte {
	edgelq_iam_proto_v1alpha_condition_proto_rawDescOnce.Do(func() {
		edgelq_iam_proto_v1alpha_condition_proto_rawDescData = protoimpl.X.CompressGZIP(edgelq_iam_proto_v1alpha_condition_proto_rawDescData)
	})
	return edgelq_iam_proto_v1alpha_condition_proto_rawDescData
}

var edgelq_iam_proto_v1alpha_condition_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var edgelq_iam_proto_v1alpha_condition_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var edgelq_iam_proto_v1alpha_condition_proto_goTypes = []interface{}{
	(Condition_ParameterType)(0),           // 0: ntt.iam.v1alpha.Condition_ParameterType
	(*Condition)(nil),                      // 1: ntt.iam.v1alpha.Condition
	(*ConditionBinding)(nil),               // 2: ntt.iam.v1alpha.ConditionBinding
	(*Condition_ParameterDeclaration)(nil), // 3: ntt.iam.v1alpha.Condition.ParameterDeclaration
	nil,                                    // 4: ntt.iam.v1alpha.ConditionBinding.ParametersEntry
}
var edgelq_iam_proto_v1alpha_condition_proto_depIdxs = []int32{
	3, // 0: ntt.iam.v1alpha.Condition.parameter_declarations:type_name -> ntt.iam.v1alpha.Condition.ParameterDeclaration
	4, // 1: ntt.iam.v1alpha.ConditionBinding.parameters:type_name -> ntt.iam.v1alpha.ConditionBinding.ParametersEntry
	0, // 2: ntt.iam.v1alpha.Condition.ParameterDeclaration.type:type_name -> ntt.iam.v1alpha.Condition_ParameterType
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { edgelq_iam_proto_v1alpha_condition_proto_init() }
func edgelq_iam_proto_v1alpha_condition_proto_init() {
	if edgelq_iam_proto_v1alpha_condition_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		edgelq_iam_proto_v1alpha_condition_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Condition); i {
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
		edgelq_iam_proto_v1alpha_condition_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConditionBinding); i {
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
		edgelq_iam_proto_v1alpha_condition_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Condition_ParameterDeclaration); i {
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
			RawDescriptor: edgelq_iam_proto_v1alpha_condition_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           edgelq_iam_proto_v1alpha_condition_proto_goTypes,
		DependencyIndexes: edgelq_iam_proto_v1alpha_condition_proto_depIdxs,
		EnumInfos:         edgelq_iam_proto_v1alpha_condition_proto_enumTypes,
		MessageInfos:      edgelq_iam_proto_v1alpha_condition_proto_msgTypes,
	}.Build()
	edgelq_iam_proto_v1alpha_condition_proto = out.File
	edgelq_iam_proto_v1alpha_condition_proto_rawDesc = nil
	edgelq_iam_proto_v1alpha_condition_proto_goTypes = nil
	edgelq_iam_proto_v1alpha_condition_proto_depIdxs = nil
}