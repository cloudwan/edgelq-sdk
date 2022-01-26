// Code generated by protoc-gen-goten-go
// File: edgelq/audit/proto/v1alpha/method_descriptor.proto
// DO NOT EDIT!!!

package method_descriptor

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
	audit_common "github.com/cloudwan/edgelq-sdk/audit/common/v1alpha"
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
	_ = &audit_common.Authentication{}
	_ = &ntt_meta.Meta{}
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// MethodDescriptor Resource
type MethodDescriptor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Name of MethodDescriptor - contains service name and
	// method type name, separated by '/' sign. Example name:
	// "iam.edgelq.com/CreateRoleBinding"
	Name *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// Optional. A concise name for the audited object type that might be
	// displayed in user interfaces. It should be a Title Cased Noun Phrase,
	// without any article or other determiners.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty" firestore:"displayName"`
	// Optional. A detailed description of the audited method type that might
	// be used in documentation.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty" firestore:"description"`
	// A set of labels used to describe instances of this audited
	// method type. For example, for "UpdateRoleBinding" we can define member
	// label. This can allow us to make query like "who tried to give user X
	// permissions to those things?"
	Labels []*audit_common.LabelDescriptor `protobuf:"bytes,4,rep,name=labels,proto3" json:"labels,omitempty" firestore:"labels"`
	// Promoted Label Key Sets allow defining multiple indexing rules for
	// underlying backend enabling query optimizations.
	PromotedLabelKeySets []*audit_common.LabelKeySet `protobuf:"bytes,5,rep,name=promoted_label_key_sets,json=promotedLabelKeySets,proto3" json:"promoted_label_key_sets,omitempty" firestore:"promotedLabelKeySets"`
	// List of API versions which define this method
	Versions []string `protobuf:"bytes,6,rep,name=versions,proto3" json:"versions,omitempty" firestore:"versions"`
	// Metadata
	Metadata *ntt_meta.Meta `protobuf:"bytes,7,opt,name=metadata,proto3" json:"metadata,omitempty" firestore:"metadata"`
}

func (m *MethodDescriptor) Reset() {
	*m = MethodDescriptor{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_audit_proto_v1alpha_method_descriptor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *MethodDescriptor) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*MethodDescriptor) ProtoMessage() {}

func (m *MethodDescriptor) ProtoReflect() preflect.Message {
	mi := &edgelq_audit_proto_v1alpha_method_descriptor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*MethodDescriptor) GotenMessage() {}

// Deprecated, Use MethodDescriptor.ProtoReflect.Descriptor instead.
func (*MethodDescriptor) Descriptor() ([]byte, []int) {
	return edgelq_audit_proto_v1alpha_method_descriptor_proto_rawDescGZIP(), []int{0}
}

func (m *MethodDescriptor) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *MethodDescriptor) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *MethodDescriptor) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *MethodDescriptor) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *MethodDescriptor) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *MethodDescriptor) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *MethodDescriptor) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *MethodDescriptor) GetLabels() []*audit_common.LabelDescriptor {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *MethodDescriptor) GetPromotedLabelKeySets() []*audit_common.LabelKeySet {
	if m != nil {
		return m.PromotedLabelKeySets
	}
	return nil
}

func (m *MethodDescriptor) GetVersions() []string {
	if m != nil {
		return m.Versions
	}
	return nil
}

func (m *MethodDescriptor) GetMetadata() *ntt_meta.Meta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *MethodDescriptor) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "MethodDescriptor"))
	}
	m.Name = fv
}

func (m *MethodDescriptor) SetDisplayName(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "DisplayName", "MethodDescriptor"))
	}
	m.DisplayName = fv
}

func (m *MethodDescriptor) SetDescription(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Description", "MethodDescriptor"))
	}
	m.Description = fv
}

func (m *MethodDescriptor) SetLabels(fv []*audit_common.LabelDescriptor) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Labels", "MethodDescriptor"))
	}
	m.Labels = fv
}

func (m *MethodDescriptor) SetPromotedLabelKeySets(fv []*audit_common.LabelKeySet) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "PromotedLabelKeySets", "MethodDescriptor"))
	}
	m.PromotedLabelKeySets = fv
}

func (m *MethodDescriptor) SetVersions(fv []string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Versions", "MethodDescriptor"))
	}
	m.Versions = fv
}

func (m *MethodDescriptor) SetMetadata(fv *ntt_meta.Meta) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Metadata", "MethodDescriptor"))
	}
	m.Metadata = fv
}

var edgelq_audit_proto_v1alpha_method_descriptor_proto preflect.FileDescriptor

var edgelq_audit_proto_v1alpha_method_descriptor_proto_rawDesc = []byte{
	0x0a, 0x32, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x6d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x6e, 0x74, 0x74, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69,
	0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71,
	0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x69, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x65, 0x64, 0x67, 0x65,
	0x6c, 0x71, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1e, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x27, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfc, 0x04, 0x0a, 0x10, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x12, 0x2c,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xb2, 0xda,
	0x21, 0x14, 0x0a, 0x12, 0x0a, 0x10, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x3a, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x22, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x55, 0x0a,
	0x17, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x65, 0x64, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f,
	0x6b, 0x65, 0x79, 0x5f, 0x73, 0x65, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x4b, 0x65, 0x79, 0x53, 0x65, 0x74, 0x52, 0x14,
	0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x65, 0x64, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x4b, 0x65, 0x79,
	0x53, 0x65, 0x74, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x2b, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4d,
	0x65, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x3a, 0x98, 0x02,
	0xea, 0x41, 0x4a, 0x0a, 0x21, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x6c,
	0x71, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x12, 0x25, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x7b, 0x6d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x7d, 0x92, 0xd9, 0x21,
	0x4b, 0x0a, 0x11, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x6f, 0x72, 0x73, 0x12, 0x11, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x73, 0x22, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x3a, 0x1d, 0x5b,
	0x5c, 0x77, 0x2e, 0x2f, 0x2d, 0x5d, 0x7b, 0x31, 0x2c, 0x31, 0x32, 0x38, 0x7d, 0x2f, 0x5b, 0x5c,
	0x77, 0x2e, 0x2f, 0x2d, 0x5d, 0x7b, 0x31, 0x2c, 0x31, 0x32, 0x38, 0x7d, 0xfa, 0xde, 0x21, 0x12,
	0x0a, 0x10, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x6f, 0x72, 0xda, 0x94, 0x23, 0x08, 0x12, 0x06, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0xc2, 0x85,
	0x2c, 0x56, 0x22, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61,
	0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x22, 0x17, 0x70, 0x72, 0x6f,
	0x6d, 0x6f, 0x74, 0x65, 0x64, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f, 0x6b, 0x65, 0x79, 0x5f,
	0x73, 0x65, 0x74, 0x73, 0x22, 0x08, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0x87, 0x04, 0xe8, 0xde, 0x21, 0x01, 0xd2,
	0xff, 0xd0, 0x02, 0x5b, 0x0a, 0x17, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x40, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77,
	0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x6d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x92,
	0x8c, 0xd1, 0x02, 0x65, 0x0a, 0x1c, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69,
	0x6e, 0x67, 0x12, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x61,
	0x75, 0x64, 0x69, 0x74, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x0a, 0x18, 0x63, 0x6f, 0x6d, 0x2e, 0x6e,
	0x74, 0x74, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x42, 0x15, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x56, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61,
	0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x2f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x6f, 0x72, 0x3b, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x6f, 0x72, 0xd2, 0x84, 0xd1, 0x02, 0x47, 0x0a, 0x0d, 0x61, 0x75, 0x64, 0x69, 0x74,
	0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x12, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64,
	0x67, 0x65, 0x6c, 0x71, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74,
	0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0xa2, 0x80, 0xd1, 0x02, 0x5d, 0x0a, 0x18, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12,
	0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x61, 0x75, 0x64, 0x69,
	0x74, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x2f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	edgelq_audit_proto_v1alpha_method_descriptor_proto_rawDescOnce sync.Once
	edgelq_audit_proto_v1alpha_method_descriptor_proto_rawDescData = edgelq_audit_proto_v1alpha_method_descriptor_proto_rawDesc
)

func edgelq_audit_proto_v1alpha_method_descriptor_proto_rawDescGZIP() []byte {
	edgelq_audit_proto_v1alpha_method_descriptor_proto_rawDescOnce.Do(func() {
		edgelq_audit_proto_v1alpha_method_descriptor_proto_rawDescData = protoimpl.X.CompressGZIP(edgelq_audit_proto_v1alpha_method_descriptor_proto_rawDescData)
	})
	return edgelq_audit_proto_v1alpha_method_descriptor_proto_rawDescData
}

var edgelq_audit_proto_v1alpha_method_descriptor_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var edgelq_audit_proto_v1alpha_method_descriptor_proto_goTypes = []interface{}{
	(*MethodDescriptor)(nil),             // 0: ntt.audit.v1alpha.MethodDescriptor
	(*audit_common.LabelDescriptor)(nil), // 1: ntt.audit.v1alpha.LabelDescriptor
	(*audit_common.LabelKeySet)(nil),     // 2: ntt.audit.v1alpha.LabelKeySet
	(*ntt_meta.Meta)(nil),                // 3: ntt.types.Meta
}
var edgelq_audit_proto_v1alpha_method_descriptor_proto_depIdxs = []int32{
	1, // 0: ntt.audit.v1alpha.MethodDescriptor.labels:type_name -> ntt.audit.v1alpha.LabelDescriptor
	2, // 1: ntt.audit.v1alpha.MethodDescriptor.promoted_label_key_sets:type_name -> ntt.audit.v1alpha.LabelKeySet
	3, // 2: ntt.audit.v1alpha.MethodDescriptor.metadata:type_name -> ntt.types.Meta
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { edgelq_audit_proto_v1alpha_method_descriptor_proto_init() }
func edgelq_audit_proto_v1alpha_method_descriptor_proto_init() {
	if edgelq_audit_proto_v1alpha_method_descriptor_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		edgelq_audit_proto_v1alpha_method_descriptor_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MethodDescriptor); i {
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
			RawDescriptor: edgelq_audit_proto_v1alpha_method_descriptor_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           edgelq_audit_proto_v1alpha_method_descriptor_proto_goTypes,
		DependencyIndexes: edgelq_audit_proto_v1alpha_method_descriptor_proto_depIdxs,
		MessageInfos:      edgelq_audit_proto_v1alpha_method_descriptor_proto_msgTypes,
	}.Build()
	edgelq_audit_proto_v1alpha_method_descriptor_proto = out.File
	edgelq_audit_proto_v1alpha_method_descriptor_proto_rawDesc = nil
	edgelq_audit_proto_v1alpha_method_descriptor_proto_goTypes = nil
	edgelq_audit_proto_v1alpha_method_descriptor_proto_depIdxs = nil
}
