// Code generated by protoc-gen-goten-go
// File: edgelq/monitoring/proto/v3/monitored_resource_descriptor.proto
// DO NOT EDIT!!!

package monitored_resource_descriptor

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
	common "github.com/cloudwan/edgelq-sdk/monitoring/resources/v3/common"
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
	_ = &common.LabelDescriptor{}
	_ = &meta.Meta{}
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// An object that describes the schema of a
// [MonitoredResource][google.api.MonitoredResource] object using a type name
// and a set of labels.  For example, the monitored resource descriptor for
// Google Compute Engine VM instances has a type of
// `"gce_instance"` and specifies the use of the labels `"instance_id"` and
// `"zone"` to identify particular VM instances.
//
// Different APIs can support different monitored resource types. APIs generally
// provide a `list` method that returns the monitored resource descriptors used
// by the API.
type MonitoredResourceDescriptor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Optional. The resource name of the monitored resource descriptor:
	// `"projects/{project_id}/monitoredResourceDescriptors/{type}"` where
	// {type} is the value of the `type` field in this object and
	// {project_id} is a project ID that provides API-specific context for
	// accessing the type.  APIs that do not use project information can use the
	// resource name format `"monitoredResourceDescriptors/{type}"`.
	// NOTE: currently only `"monitoredResourceDescriptors/{type}"` form is
	// supported.
	Name *Name `protobuf:"bytes,5,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// Required. The monitored resource type. For example, the type
	// `"cloudsql_database"` represents databases in Google Cloud SQL.
	// The maximum length of this value is 256 characters.
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty" firestore:"type"`
	// Optional. A concise name for the monitored resource type that might be
	// displayed in user interfaces. It should be a Title Cased Noun Phrase,
	// without any article or other determiners. For example,
	// `"Google Cloud SQL Database"`.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty" firestore:"displayName"`
	// Optional. A detailed description of the monitored resource type that might
	// be used in documentation.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty" firestore:"description"`
	// Required. A set of labels used to describe instances of this monitored
	// resource type. For example, an individual Google Cloud SQL database is
	// identified by values for the labels `"database_id"` and `"zone"`.
	Labels []*common.LabelDescriptor `protobuf:"bytes,4,rep,name=labels,proto3" json:"labels,omitempty" firestore:"labels"`
	// Promoted Label Key Sets allow defining multiple indexing rules for
	// underlying backend enabling query optimizations. Metric promoted label sets
	// are combined with MonitoredResource promoted label sets and result in
	// PromotedKeySet.
	PromotedLabelKeySets []*common.LabelKeySet `protobuf:"bytes,32,rep,name=promoted_label_key_sets,json=promotedLabelKeySets,proto3" json:"promoted_label_key_sets,omitempty" firestore:"promotedLabelKeySets"`
	// Metadata
	Metadata *meta.Meta `protobuf:"bytes,33,opt,name=metadata,proto3" json:"metadata,omitempty" firestore:"metadata"`
}

func (m *MonitoredResourceDescriptor) Reset() {
	*m = MonitoredResourceDescriptor{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_monitoring_proto_v3_monitored_resource_descriptor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *MonitoredResourceDescriptor) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*MonitoredResourceDescriptor) ProtoMessage() {}

func (m *MonitoredResourceDescriptor) ProtoReflect() preflect.Message {
	mi := &edgelq_monitoring_proto_v3_monitored_resource_descriptor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*MonitoredResourceDescriptor) GotenMessage() {}

// Deprecated, Use MonitoredResourceDescriptor.ProtoReflect.Descriptor instead.
func (*MonitoredResourceDescriptor) Descriptor() ([]byte, []int) {
	return edgelq_monitoring_proto_v3_monitored_resource_descriptor_proto_rawDescGZIP(), []int{0}
}

func (m *MonitoredResourceDescriptor) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *MonitoredResourceDescriptor) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *MonitoredResourceDescriptor) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *MonitoredResourceDescriptor) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *MonitoredResourceDescriptor) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *MonitoredResourceDescriptor) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *MonitoredResourceDescriptor) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *MonitoredResourceDescriptor) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *MonitoredResourceDescriptor) GetLabels() []*common.LabelDescriptor {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *MonitoredResourceDescriptor) GetPromotedLabelKeySets() []*common.LabelKeySet {
	if m != nil {
		return m.PromotedLabelKeySets
	}
	return nil
}

func (m *MonitoredResourceDescriptor) GetMetadata() *meta.Meta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *MonitoredResourceDescriptor) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "MonitoredResourceDescriptor"))
	}
	m.Name = fv
}

func (m *MonitoredResourceDescriptor) SetType(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Type", "MonitoredResourceDescriptor"))
	}
	m.Type = fv
}

func (m *MonitoredResourceDescriptor) SetDisplayName(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "DisplayName", "MonitoredResourceDescriptor"))
	}
	m.DisplayName = fv
}

func (m *MonitoredResourceDescriptor) SetDescription(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Description", "MonitoredResourceDescriptor"))
	}
	m.Description = fv
}

func (m *MonitoredResourceDescriptor) SetLabels(fv []*common.LabelDescriptor) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Labels", "MonitoredResourceDescriptor"))
	}
	m.Labels = fv
}

func (m *MonitoredResourceDescriptor) SetPromotedLabelKeySets(fv []*common.LabelKeySet) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "PromotedLabelKeySets", "MonitoredResourceDescriptor"))
	}
	m.PromotedLabelKeySets = fv
}

func (m *MonitoredResourceDescriptor) SetMetadata(fv *meta.Meta) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Metadata", "MonitoredResourceDescriptor"))
	}
	m.Metadata = fv
}

var edgelq_monitoring_proto_v3_monitored_resource_descriptor_proto preflect.FileDescriptor

var edgelq_monitoring_proto_v3_monitored_resource_descriptor_proto_rawDesc = []byte{
	0x0a, 0x3e, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x69, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x33, 0x2f, 0x6d, 0x6f, 0x6e,
	0x69, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x11, 0x6e, 0x74, 0x74, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x33, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20,
	0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x24, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x75, 0x64, 0x69,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x27, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69,
	0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x33, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf8, 0x05, 0x0a, 0x1b, 0x4d, 0x6f, 0x6e,
	0x69, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x12, 0x37, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x23, 0xb2, 0xda, 0x21, 0x1f, 0x0a, 0x1d, 0x0a, 0x1b,
	0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x40, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x2c, 0xca, 0xc6, 0x27, 0x28, 0x2a, 0x26, 0x52, 0x24, 0x42, 0x22, 0x5e, 0x5b, 0x41, 0x2d, 0x5a,
	0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5f, 0x5c, 0x2d, 0x2e, 0x2c, 0x2b, 0x21, 0x2a, 0x28, 0x29,
	0x25, 0x5c, 0x5c, 0x2f, 0x5d, 0x7b, 0x31, 0x2c, 0x32, 0x35, 0x36, 0x7d, 0x24, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x46, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x6d,
	0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x4c, 0x61, 0x62,
	0x65, 0x6c, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x42, 0x0a, 0xca, 0xc6,
	0x27, 0x06, 0x42, 0x04, 0x12, 0x02, 0x08, 0x19, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x12, 0x55, 0x0a, 0x17, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x65, 0x64, 0x5f, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x73, 0x65, 0x74, 0x73, 0x18, 0x20, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1e, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69,
	0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x4b, 0x65, 0x79, 0x53, 0x65,
	0x74, 0x52, 0x14, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x65, 0x64, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x4b, 0x65, 0x79, 0x53, 0x65, 0x74, 0x73, 0x12, 0x2d, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x21, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67, 0x6f, 0x74, 0x65,
	0x6e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x3a, 0xca, 0x02, 0xea, 0x41, 0x71, 0x0a, 0x31, 0x6d, 0x6f,
	0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x12,
	0x3c, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x7b, 0x6d,
	0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x7d, 0x92, 0xd9, 0x21,
	0x70, 0x0a, 0x1c, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x73, 0x12,
	0x1c, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x73, 0x2a, 0x0e, 0x5b,
	0x5c, 0x77, 0x2e, 0x2f, 0x2d, 0x5d, 0x7b, 0x34, 0x2c, 0x31, 0x32, 0x38, 0x7d, 0x42, 0x22, 0x08,
	0x02, 0x12, 0x06, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x0c, 0x64, 0x69, 0x73,
	0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x06, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0xda, 0x94, 0x23, 0x08, 0x12, 0x06, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0xc2, 0x85, 0x2c,
	0x52, 0x22, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x06, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x22, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x17, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x65, 0x64, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x73, 0x65, 0x74, 0x73, 0x42, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x42, 0xa4, 0x03, 0xe8, 0xde, 0x21, 0x01, 0xd2, 0xff, 0xd0, 0x02, 0x73, 0x0a,
	0x23, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x5f, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x12, 0x4c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71,
	0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2f, 0x76, 0x33, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x5f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x6f, 0x72, 0xa2, 0x80, 0xd1, 0x02, 0x75, 0x0a, 0x24, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x65, 0x64, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x4d, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77,
	0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f,
	0x72, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x6d,
	0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x0a, 0x18, 0x63, 0x6f,
	0x6d, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67,
	0x2e, 0x70, 0x62, 0x2e, 0x76, 0x33, 0x42, 0x20, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x65,
	0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x6e, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f,
	0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e,
	0x67, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x6d,
	0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x3b, 0x6d, 0x6f, 0x6e,
	0x69, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	edgelq_monitoring_proto_v3_monitored_resource_descriptor_proto_rawDescOnce sync.Once
	edgelq_monitoring_proto_v3_monitored_resource_descriptor_proto_rawDescData = edgelq_monitoring_proto_v3_monitored_resource_descriptor_proto_rawDesc
)

func edgelq_monitoring_proto_v3_monitored_resource_descriptor_proto_rawDescGZIP() []byte {
	edgelq_monitoring_proto_v3_monitored_resource_descriptor_proto_rawDescOnce.Do(func() {
		edgelq_monitoring_proto_v3_monitored_resource_descriptor_proto_rawDescData = protoimpl.X.CompressGZIP(edgelq_monitoring_proto_v3_monitored_resource_descriptor_proto_rawDescData)
	})
	return edgelq_monitoring_proto_v3_monitored_resource_descriptor_proto_rawDescData
}

var edgelq_monitoring_proto_v3_monitored_resource_descriptor_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var edgelq_monitoring_proto_v3_monitored_resource_descriptor_proto_goTypes = []interface{}{
	(*MonitoredResourceDescriptor)(nil), // 0: ntt.monitoring.v3.MonitoredResourceDescriptor
	(*common.LabelDescriptor)(nil),      // 1: ntt.monitoring.v3.LabelDescriptor
	(*common.LabelKeySet)(nil),          // 2: ntt.monitoring.v3.LabelKeySet
	(*meta.Meta)(nil),                   // 3: goten.types.Meta
}
var edgelq_monitoring_proto_v3_monitored_resource_descriptor_proto_depIdxs = []int32{
	1, // 0: ntt.monitoring.v3.MonitoredResourceDescriptor.labels:type_name -> ntt.monitoring.v3.LabelDescriptor
	2, // 1: ntt.monitoring.v3.MonitoredResourceDescriptor.promoted_label_key_sets:type_name -> ntt.monitoring.v3.LabelKeySet
	3, // 2: ntt.monitoring.v3.MonitoredResourceDescriptor.metadata:type_name -> goten.types.Meta
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { edgelq_monitoring_proto_v3_monitored_resource_descriptor_proto_init() }
func edgelq_monitoring_proto_v3_monitored_resource_descriptor_proto_init() {
	if edgelq_monitoring_proto_v3_monitored_resource_descriptor_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		edgelq_monitoring_proto_v3_monitored_resource_descriptor_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MonitoredResourceDescriptor); i {
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
			RawDescriptor: edgelq_monitoring_proto_v3_monitored_resource_descriptor_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           edgelq_monitoring_proto_v3_monitored_resource_descriptor_proto_goTypes,
		DependencyIndexes: edgelq_monitoring_proto_v3_monitored_resource_descriptor_proto_depIdxs,
		MessageInfos:      edgelq_monitoring_proto_v3_monitored_resource_descriptor_proto_msgTypes,
	}.Build()
	edgelq_monitoring_proto_v3_monitored_resource_descriptor_proto = out.File
	edgelq_monitoring_proto_v3_monitored_resource_descriptor_proto_rawDesc = nil
	edgelq_monitoring_proto_v3_monitored_resource_descriptor_proto_goTypes = nil
	edgelq_monitoring_proto_v3_monitored_resource_descriptor_proto_depIdxs = nil
}
