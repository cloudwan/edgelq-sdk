// Code generated by protoc-gen-goten-go
// File: edgelq/applications/proto/v1alpha2/distribution.proto
// DO NOT EDIT!!!

package distribution

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
	pod "github.com/cloudwan/edgelq-sdk/applications/resources/v1alpha2/pod"
	project "github.com/cloudwan/edgelq-sdk/applications/resources/v1alpha2/project"
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
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
	_ = &pod.Pod{}
	_ = &project.Project{}
	_ = &ntt_meta.Meta{}
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Distribution Resource
type Distribution struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Name of Distribution
	Name *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// Display name of Distribution
	DisplayName string               `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty" firestore:"displayName"`
	Metadata    *ntt_meta.Meta       `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty" firestore:"metadata"`
	Spec        *Distribution_Spec   `protobuf:"bytes,4,opt,name=spec,proto3" json:"spec,omitempty" firestore:"spec"`
	Status      *Distribution_Status `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty" firestore:"status"`
}

func (m *Distribution) Reset() {
	*m = Distribution{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_applications_proto_v1alpha2_distribution_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *Distribution) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*Distribution) ProtoMessage() {}

func (m *Distribution) ProtoReflect() preflect.Message {
	mi := &edgelq_applications_proto_v1alpha2_distribution_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*Distribution) GotenMessage() {}

// Deprecated, Use Distribution.ProtoReflect.Descriptor instead.
func (*Distribution) Descriptor() ([]byte, []int) {
	return edgelq_applications_proto_v1alpha2_distribution_proto_rawDescGZIP(), []int{0}
}

func (m *Distribution) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *Distribution) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *Distribution) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *Distribution) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *Distribution) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *Distribution) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *Distribution) GetMetadata() *ntt_meta.Meta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Distribution) GetSpec() *Distribution_Spec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *Distribution) GetStatus() *Distribution_Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *Distribution) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "Distribution"))
	}
	m.Name = fv
}

func (m *Distribution) SetDisplayName(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "DisplayName", "Distribution"))
	}
	m.DisplayName = fv
}

func (m *Distribution) SetMetadata(fv *ntt_meta.Meta) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Metadata", "Distribution"))
	}
	m.Metadata = fv
}

func (m *Distribution) SetSpec(fv *Distribution_Spec) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Spec", "Distribution"))
	}
	m.Spec = fv
}

func (m *Distribution) SetStatus(fv *Distribution_Status) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Status", "Distribution"))
	}
	m.Status = fv
}

// A label selector is a label query over a set of resources. The result of
// matchLabels and matchExpressions are ANDed. An empty label selector matches
// all objects. A null label selector matches no objects.
type LabelSelector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// matchLabels is a map of {key,value} pairs. A single {key,value} in the
	// matchLabels map is equivalent to an element of matchExpressions, whose key
	// field is "key", the operator is "In", and the values array contains only
	// "value". The requirements are ANDed. +optional
	MatchLabels map[string]string `protobuf:"bytes,1,rep,name=match_labels,json=matchLabels,proto3" json:"match_labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" firestore:"matchLabels"`
	// matchExpressions is a list of label selector requirements. The requirements
	// are ANDed. +optional
	MatchExpressions []*LabelSelectorRequirement `protobuf:"bytes,2,rep,name=match_expressions,json=matchExpressions,proto3" json:"match_expressions,omitempty" firestore:"matchExpressions"`
}

func (m *LabelSelector) Reset() {
	*m = LabelSelector{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_applications_proto_v1alpha2_distribution_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *LabelSelector) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*LabelSelector) ProtoMessage() {}

func (m *LabelSelector) ProtoReflect() preflect.Message {
	mi := &edgelq_applications_proto_v1alpha2_distribution_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*LabelSelector) GotenMessage() {}

// Deprecated, Use LabelSelector.ProtoReflect.Descriptor instead.
func (*LabelSelector) Descriptor() ([]byte, []int) {
	return edgelq_applications_proto_v1alpha2_distribution_proto_rawDescGZIP(), []int{1}
}

func (m *LabelSelector) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *LabelSelector) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *LabelSelector) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *LabelSelector) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *LabelSelector) GetMatchLabels() map[string]string {
	if m != nil {
		return m.MatchLabels
	}
	return nil
}

func (m *LabelSelector) GetMatchExpressions() []*LabelSelectorRequirement {
	if m != nil {
		return m.MatchExpressions
	}
	return nil
}

func (m *LabelSelector) SetMatchLabels(fv map[string]string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "MatchLabels", "LabelSelector"))
	}
	m.MatchLabels = fv
}

func (m *LabelSelector) SetMatchExpressions(fv []*LabelSelectorRequirement) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "MatchExpressions", "LabelSelector"))
	}
	m.MatchExpressions = fv
}

// A label selector requirement is a selector that contains values, a key, and
// an operator that relates the key and values.
type LabelSelectorRequirement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// key is the label key that the selector applies to.
	// +patchMergeKey=key
	// +patchStrategy=merge
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty" firestore:"key"`
	// operator represents a key's relationship to a set of values.
	// Valid operators are In, NotIn, Exists and DoesNotExist.
	Operator string `protobuf:"bytes,2,opt,name=operator,proto3" json:"operator,omitempty" firestore:"operator"`
	// values is an array of string values. If the operator is In or NotIn,
	// the values array must be non-empty. If the operator is Exists or
	// DoesNotExist, the values array must be empty. This array is replaced during
	// a strategic merge patch. +optional
	Values []string `protobuf:"bytes,3,rep,name=values,proto3" json:"values,omitempty" firestore:"values"`
}

func (m *LabelSelectorRequirement) Reset() {
	*m = LabelSelectorRequirement{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_applications_proto_v1alpha2_distribution_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *LabelSelectorRequirement) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*LabelSelectorRequirement) ProtoMessage() {}

func (m *LabelSelectorRequirement) ProtoReflect() preflect.Message {
	mi := &edgelq_applications_proto_v1alpha2_distribution_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*LabelSelectorRequirement) GotenMessage() {}

// Deprecated, Use LabelSelectorRequirement.ProtoReflect.Descriptor instead.
func (*LabelSelectorRequirement) Descriptor() ([]byte, []int) {
	return edgelq_applications_proto_v1alpha2_distribution_proto_rawDescGZIP(), []int{2}
}

func (m *LabelSelectorRequirement) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *LabelSelectorRequirement) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *LabelSelectorRequirement) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *LabelSelectorRequirement) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *LabelSelectorRequirement) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *LabelSelectorRequirement) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *LabelSelectorRequirement) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *LabelSelectorRequirement) SetKey(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Key", "LabelSelectorRequirement"))
	}
	m.Key = fv
}

func (m *LabelSelectorRequirement) SetOperator(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Operator", "LabelSelectorRequirement"))
	}
	m.Operator = fv
}

func (m *LabelSelectorRequirement) SetValues(fv []string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Values", "LabelSelectorRequirement"))
	}
	m.Values = fv
}

// Spec defines the configuration of a Distribution
type Distribution_Spec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Selector      *LabelSelector `protobuf:"bytes,1,opt,name=selector,proto3" json:"selector,omitempty" firestore:"selector"`
	// template defines the probing config to be distributed.
	Template *Distribution_Spec_Template `protobuf:"bytes,2,opt,name=template,proto3" json:"template,omitempty" firestore:"template"`
}

func (m *Distribution_Spec) Reset() {
	*m = Distribution_Spec{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_applications_proto_v1alpha2_distribution_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *Distribution_Spec) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*Distribution_Spec) ProtoMessage() {}

func (m *Distribution_Spec) ProtoReflect() preflect.Message {
	mi := &edgelq_applications_proto_v1alpha2_distribution_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*Distribution_Spec) GotenMessage() {}

// Deprecated, Use Distribution_Spec.ProtoReflect.Descriptor instead.
func (*Distribution_Spec) Descriptor() ([]byte, []int) {
	return edgelq_applications_proto_v1alpha2_distribution_proto_rawDescGZIP(), []int{0, 0}
}

func (m *Distribution_Spec) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *Distribution_Spec) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *Distribution_Spec) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *Distribution_Spec) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *Distribution_Spec) GetSelector() *LabelSelector {
	if m != nil {
		return m.Selector
	}
	return nil
}

func (m *Distribution_Spec) GetTemplate() *Distribution_Spec_Template {
	if m != nil {
		return m.Template
	}
	return nil
}

func (m *Distribution_Spec) SetSelector(fv *LabelSelector) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Selector", "Distribution_Spec"))
	}
	m.Selector = fv
}

func (m *Distribution_Spec) SetTemplate(fv *Distribution_Spec_Template) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Template", "Distribution_Spec"))
	}
	m.Template = fv
}

type Distribution_Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (m *Distribution_Status) Reset() {
	*m = Distribution_Status{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_applications_proto_v1alpha2_distribution_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *Distribution_Status) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*Distribution_Status) ProtoMessage() {}

func (m *Distribution_Status) ProtoReflect() preflect.Message {
	mi := &edgelq_applications_proto_v1alpha2_distribution_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*Distribution_Status) GotenMessage() {}

// Deprecated, Use Distribution_Status.ProtoReflect.Descriptor instead.
func (*Distribution_Status) Descriptor() ([]byte, []int) {
	return edgelq_applications_proto_v1alpha2_distribution_proto_rawDescGZIP(), []int{0, 1}
}

func (m *Distribution_Status) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *Distribution_Status) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *Distribution_Status) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *Distribution_Status) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

type Distribution_Spec_Template struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Metadata      *ntt_meta.Meta `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty" firestore:"metadata"`
	Spec          *pod.Pod_Spec  `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty" firestore:"spec"`
}

func (m *Distribution_Spec_Template) Reset() {
	*m = Distribution_Spec_Template{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_applications_proto_v1alpha2_distribution_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *Distribution_Spec_Template) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*Distribution_Spec_Template) ProtoMessage() {}

func (m *Distribution_Spec_Template) ProtoReflect() preflect.Message {
	mi := &edgelq_applications_proto_v1alpha2_distribution_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*Distribution_Spec_Template) GotenMessage() {}

// Deprecated, Use Distribution_Spec_Template.ProtoReflect.Descriptor instead.
func (*Distribution_Spec_Template) Descriptor() ([]byte, []int) {
	return edgelq_applications_proto_v1alpha2_distribution_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (m *Distribution_Spec_Template) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *Distribution_Spec_Template) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *Distribution_Spec_Template) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *Distribution_Spec_Template) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *Distribution_Spec_Template) GetMetadata() *ntt_meta.Meta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Distribution_Spec_Template) GetSpec() *pod.Pod_Spec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *Distribution_Spec_Template) SetMetadata(fv *ntt_meta.Meta) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Metadata", "Distribution_Spec_Template"))
	}
	m.Metadata = fv
}

func (m *Distribution_Spec_Template) SetSpec(fv *pod.Pod_Spec) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Spec", "Distribution_Spec_Template"))
	}
	m.Spec = fv
}

var edgelq_applications_proto_v1alpha2_distribution_proto preflect.FileDescriptor

var edgelq_applications_proto_v1alpha2_distribution_proto_rawDesc = []byte{
	0x0a, 0x35, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x32, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x6e, 0x74, 0x74, 0x2e, 0x61, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x32, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67,
	0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20,
	0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x34, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x63, 0x6f, 0x6d,
	0x70, 0x69, 0x6c, 0x65, 0x72, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x61, 0x75,
	0x64, 0x69, 0x74, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x65, 0x64, 0x67,
	0x65, 0x6c, 0x71, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x65, 0x64, 0x67,
	0x65, 0x6c, 0x71, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f,
	0x70, 0x6f, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa8, 0x07, 0x0a, 0x0c, 0x44, 0x69,
	0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x14, 0xb2, 0xda, 0x21, 0x10, 0x0a, 0x0e,
	0x0a, 0x0c, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70,
	0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2b, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6e, 0x74, 0x74, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x40, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x44,
	0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x70, 0x65, 0x63,
	0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12, 0x46, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x61, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x32, 0x2e, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x1a, 0x91,
	0x02, 0x0a, 0x04, 0x53, 0x70, 0x65, 0x63, 0x12, 0x44, 0x0a, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x6e, 0x74, 0x74, 0x2e,
	0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x53, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x52, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x51, 0x0a,
	0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x35, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x44, 0x69, 0x73, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x54, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x1a, 0x70, 0x0a, 0x08, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x2b, 0x0a, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x37, 0x0a, 0x04, 0x73, 0x70, 0x65,
	0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x61, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x32, 0x2e, 0x50, 0x6f, 0x64, 0x2e, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70,
	0x65, 0x63, 0x1a, 0x08, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x3a, 0xf5, 0x02, 0xea,
	0x41, 0x57, 0x0a, 0x24, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x69, 0x73, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x64, 0x69, 0x73,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x64, 0x69, 0x73, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x92, 0xd9, 0x21, 0x8a, 0x01, 0x0a, 0x0d,
	0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x0d, 0x64,
	0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x07, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x4a, 0x59, 0x08, 0x02, 0x12,
	0x06, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0f, 0x0a, 0x0d, 0x73, 0x70, 0x65, 0x63, 0x2e,
	0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x08, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x0f, 0x0a, 0x0d, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x74,
	0x61, 0x67, 0x73, 0x12, 0x11, 0x0a, 0x0f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x50, 0x05, 0xb2, 0xdf, 0x21, 0x4c, 0x0a, 0x4a, 0x0a, 0x48,
	0x0a, 0x0d, 0x62, 0x79, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x1a, 0x22, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x7d, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x2d, 0x2a, 0x05, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0xda, 0x94, 0x23, 0x08, 0x12, 0x06, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0xc2, 0x85, 0x2c, 0x2c, 0x22, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x0c,
	0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x04, 0x73, 0x70,
	0x65, 0x63, 0x2a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x8f, 0x02, 0x0a, 0x0d, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x53, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x5c, 0x0a, 0x0c, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x6e,
	0x74, 0x74, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x53, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x12, 0x60, 0x0a, 0x11, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x65, 0x78,
	0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x33, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x10, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x45, 0x78, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x3e, 0x0a, 0x10, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x60, 0x0a, 0x18, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x53,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x12, 0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x42, 0xa3, 0x03, 0xe8, 0xde, 0x21, 0x01, 0xd2,
	0xff, 0xd0, 0x02, 0x59, 0x0a, 0x12, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64,
	0x67, 0x65, 0x6c, 0x71, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32,
	0x2f, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x0a, 0x20, 0x63,
	0x6f, 0x6d, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x42,
	0x11, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x54, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71,
	0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32,
	0x2f, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x3b, 0x64, 0x69,
	0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0xd2, 0x84, 0xd1, 0x02, 0x4f, 0x0a,
	0x0d, 0x61, 0x75, 0x64, 0x69, 0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x12, 0x3e,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x68, 0x61, 0x6e,
	0x64, 0x6c, 0x65, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0xa2, 0x80,
	0xd1, 0x02, 0x5b, 0x0a, 0x13, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64,
	0x67, 0x65, 0x6c, 0x71, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x32, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	edgelq_applications_proto_v1alpha2_distribution_proto_rawDescOnce sync.Once
	edgelq_applications_proto_v1alpha2_distribution_proto_rawDescData = edgelq_applications_proto_v1alpha2_distribution_proto_rawDesc
)

func edgelq_applications_proto_v1alpha2_distribution_proto_rawDescGZIP() []byte {
	edgelq_applications_proto_v1alpha2_distribution_proto_rawDescOnce.Do(func() {
		edgelq_applications_proto_v1alpha2_distribution_proto_rawDescData = protoimpl.X.CompressGZIP(edgelq_applications_proto_v1alpha2_distribution_proto_rawDescData)
	})
	return edgelq_applications_proto_v1alpha2_distribution_proto_rawDescData
}

var edgelq_applications_proto_v1alpha2_distribution_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var edgelq_applications_proto_v1alpha2_distribution_proto_goTypes = []interface{}{
	(*Distribution)(nil),               // 0: ntt.applications.v1alpha2.Distribution
	(*LabelSelector)(nil),              // 1: ntt.applications.v1alpha2.LabelSelector
	(*LabelSelectorRequirement)(nil),   // 2: ntt.applications.v1alpha2.LabelSelectorRequirement
	(*Distribution_Spec)(nil),          // 3: ntt.applications.v1alpha2.Distribution.Spec
	(*Distribution_Status)(nil),        // 4: ntt.applications.v1alpha2.Distribution.Status
	(*Distribution_Spec_Template)(nil), // 5: ntt.applications.v1alpha2.Distribution.Spec.Template
	nil,                                // 6: ntt.applications.v1alpha2.LabelSelector.MatchLabelsEntry
	(*ntt_meta.Meta)(nil),              // 7: ntt.types.Meta
	(*pod.Pod_Spec)(nil),               // 8: ntt.applications.v1alpha2.Pod.Spec
}
var edgelq_applications_proto_v1alpha2_distribution_proto_depIdxs = []int32{
	7, // 0: ntt.applications.v1alpha2.Distribution.metadata:type_name -> ntt.types.Meta
	3, // 1: ntt.applications.v1alpha2.Distribution.spec:type_name -> ntt.applications.v1alpha2.Distribution.Spec
	4, // 2: ntt.applications.v1alpha2.Distribution.status:type_name -> ntt.applications.v1alpha2.Distribution.Status
	6, // 3: ntt.applications.v1alpha2.LabelSelector.match_labels:type_name -> ntt.applications.v1alpha2.LabelSelector.MatchLabelsEntry
	2, // 4: ntt.applications.v1alpha2.LabelSelector.match_expressions:type_name -> ntt.applications.v1alpha2.LabelSelectorRequirement
	1, // 5: ntt.applications.v1alpha2.Distribution.Spec.selector:type_name -> ntt.applications.v1alpha2.LabelSelector
	5, // 6: ntt.applications.v1alpha2.Distribution.Spec.template:type_name -> ntt.applications.v1alpha2.Distribution.Spec.Template
	7, // 7: ntt.applications.v1alpha2.Distribution.Spec.Template.metadata:type_name -> ntt.types.Meta
	8, // 8: ntt.applications.v1alpha2.Distribution.Spec.Template.spec:type_name -> ntt.applications.v1alpha2.Pod.Spec
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { edgelq_applications_proto_v1alpha2_distribution_proto_init() }
func edgelq_applications_proto_v1alpha2_distribution_proto_init() {
	if edgelq_applications_proto_v1alpha2_distribution_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		edgelq_applications_proto_v1alpha2_distribution_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Distribution); i {
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
		edgelq_applications_proto_v1alpha2_distribution_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LabelSelector); i {
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
		edgelq_applications_proto_v1alpha2_distribution_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LabelSelectorRequirement); i {
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
		edgelq_applications_proto_v1alpha2_distribution_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Distribution_Spec); i {
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
		edgelq_applications_proto_v1alpha2_distribution_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Distribution_Status); i {
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
		edgelq_applications_proto_v1alpha2_distribution_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Distribution_Spec_Template); i {
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
			RawDescriptor: edgelq_applications_proto_v1alpha2_distribution_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           edgelq_applications_proto_v1alpha2_distribution_proto_goTypes,
		DependencyIndexes: edgelq_applications_proto_v1alpha2_distribution_proto_depIdxs,
		MessageInfos:      edgelq_applications_proto_v1alpha2_distribution_proto_msgTypes,
	}.Build()
	edgelq_applications_proto_v1alpha2_distribution_proto = out.File
	edgelq_applications_proto_v1alpha2_distribution_proto_rawDesc = nil
	edgelq_applications_proto_v1alpha2_distribution_proto_goTypes = nil
	edgelq_applications_proto_v1alpha2_distribution_proto_depIdxs = nil
}
