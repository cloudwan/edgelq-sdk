// Code generated by protoc-gen-goten-go
// File: edgelq/applications/proto/v1/config_map.proto
// DO NOT EDIT!!!

package config_map

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
	project "github.com/cloudwan/edgelq-sdk/applications/resources/v1/project"
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
	_ = &project.Project{}
	_ = &meta.Meta{}
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ConfigMap Resource
type ConfigMap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Name of ConfigMap
	// When creating a new instance, this field is optional and if not provided,
	// it will be generated automatically. Last ID segment must conform to the
	// following regex: [a-z][a-z0-9\\-]{0,28}[a-z0-9]
	Name *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty"`
	// Metadata is an object with information like create, update and delete time
	// (for async deleted resources), has user labels/annotations, sharding
	// information, multi-region syncing information and may have non-schema
	// owners (useful for taking ownership of resources belonging to lower level
	// services by higher ones).
	Metadata *meta.Meta `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Display name of ConfigMap
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Optional. Description of the ConfigMap.
	Description string `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	// Data contains the configuration data.
	// Each key must consist of alphanumeric characters, '-', '_' or '.'.
	// Values with non-UTF-8 byte sequences must use the BinaryData field.
	// The keys stored in Data must not overlap with the keys in
	// the BinaryData field, this is enforced during validation process.
	// +optional
	Data map[string]string `protobuf:"bytes,4,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// BinaryData contains the binary data.
	// Each key must consist of alphanumeric characters, '-', '_' or '.'.
	// BinaryData can contain byte sequences that are not in the UTF-8 range.
	// The keys stored in BinaryData must not overlap with the ones in
	// the Data field, this is enforced during validation process.
	// Using this field will require 1.10+ apiserver and
	// kubelet.
	// +optional
	BinaryData map[string][]byte `protobuf:"bytes,5,rep,name=binary_data,json=binaryData,proto3" json:"binary_data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *ConfigMap) Reset() {
	*m = ConfigMap{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_applications_proto_v1_config_map_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *ConfigMap) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*ConfigMap) ProtoMessage() {}

func (m *ConfigMap) ProtoReflect() preflect.Message {
	mi := &edgelq_applications_proto_v1_config_map_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*ConfigMap) GotenMessage() {}

// Deprecated, Use ConfigMap.ProtoReflect.Descriptor instead.
func (*ConfigMap) Descriptor() ([]byte, []int) {
	return edgelq_applications_proto_v1_config_map_proto_rawDescGZIP(), []int{0}
}

func (m *ConfigMap) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *ConfigMap) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *ConfigMap) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *ConfigMap) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *ConfigMap) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *ConfigMap) GetMetadata() *meta.Meta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *ConfigMap) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *ConfigMap) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ConfigMap) GetData() map[string]string {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *ConfigMap) GetBinaryData() map[string][]byte {
	if m != nil {
		return m.BinaryData
	}
	return nil
}

func (m *ConfigMap) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "ConfigMap"))
	}
	m.Name = fv
}

func (m *ConfigMap) SetMetadata(fv *meta.Meta) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Metadata", "ConfigMap"))
	}
	m.Metadata = fv
}

func (m *ConfigMap) SetDisplayName(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "DisplayName", "ConfigMap"))
	}
	m.DisplayName = fv
}

func (m *ConfigMap) SetDescription(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Description", "ConfigMap"))
	}
	m.Description = fv
}

func (m *ConfigMap) SetData(fv map[string]string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Data", "ConfigMap"))
	}
	m.Data = fv
}

func (m *ConfigMap) SetBinaryData(fv map[string][]byte) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "BinaryData", "ConfigMap"))
	}
	m.BinaryData = fv
}

var edgelq_applications_proto_v1_config_map_proto preflect.FileDescriptor

var edgelq_applications_proto_v1_config_map_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x6d, 0x61, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x13, 0x6e, 0x74, 0x74, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x76, 0x31, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
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
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x86, 0x06, 0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4d, 0x61, 0x70, 0x12, 0x25,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x11, 0xb2, 0xda,
	0x21, 0x0d, 0x0a, 0x0b, 0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4d, 0x61, 0x70, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x2e, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0xca, 0xc6, 0x27, 0x07,
	0x2a, 0x05, 0x22, 0x03, 0x08, 0x80, 0x02, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0xca, 0xc6, 0x27, 0x07, 0x2a,
	0x05, 0x22, 0x03, 0x08, 0x80, 0x02, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x3c, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x28, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4d, 0x61,
	0x70, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x4f, 0x0a, 0x0b, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x5f, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x61, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x4d, 0x61, 0x70, 0x2e, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x44, 0x61, 0x74,
	0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x44, 0x61,
	0x74, 0x61, 0x1a, 0x37, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3d, 0x0a, 0x0f, 0x42,
	0x69, 0x6e, 0x61, 0x72, 0x79, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x3a, 0xbc, 0x02, 0xea, 0x41, 0x60,
	0x0a, 0x21, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x65,
	0x64, 0x67, 0x65, 0x6c, 0x71, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x4d, 0x61, 0x70, 0x12, 0x3b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x7b, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4d,
	0x61, 0x70, 0x73, 0x2f, 0x7b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x6d, 0x61, 0x70, 0x7d,
	0x92, 0xd9, 0x21, 0x3d, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4d, 0x61, 0x70, 0x73,
	0x12, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4d, 0x61, 0x70, 0x73, 0x1a, 0x07, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x18, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x38,
	0x05, 0xb2, 0xdf, 0x21, 0x87, 0x01, 0x0a, 0x84, 0x01, 0x0a, 0x81, 0x01, 0x0a, 0x0d, 0x62, 0x79,
	0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x0c, 0x64, 0x69, 0x73,
	0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x1a, 0x30, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x72, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4d, 0x61, 0x70, 0x73, 0x2f, 0x2d, 0x1a, 0x29, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f,
	0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2d, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x4d, 0x61, 0x70, 0x73, 0x2f, 0x2d, 0x2a, 0x05, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0xda, 0x94, 0x23,
	0x08, 0x12, 0x06, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0xa8, 0x02, 0xe8, 0xde, 0x21, 0x01,
	0xd2, 0xff, 0xd0, 0x02, 0x4f, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x6d, 0x61,
	0x70, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67,
	0x65, 0x6c, 0x71, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x5f, 0x6d, 0x61, 0x70, 0xa2, 0x80, 0xd1, 0x02, 0x51, 0x0a, 0x11, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x5f, 0x6d, 0x61, 0x70, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x3c, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61,
	0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x6d, 0x61, 0x70, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e,
	0x6e, 0x74, 0x74, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x42, 0x0e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4d, 0x61,
	0x70, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64,
	0x67, 0x65, 0x6c, 0x71, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x6d, 0x61, 0x70, 0x3b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x5f, 0x6d, 0x61, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	edgelq_applications_proto_v1_config_map_proto_rawDescOnce sync.Once
	edgelq_applications_proto_v1_config_map_proto_rawDescData = edgelq_applications_proto_v1_config_map_proto_rawDesc
)

func edgelq_applications_proto_v1_config_map_proto_rawDescGZIP() []byte {
	edgelq_applications_proto_v1_config_map_proto_rawDescOnce.Do(func() {
		edgelq_applications_proto_v1_config_map_proto_rawDescData = protoimpl.X.CompressGZIP(edgelq_applications_proto_v1_config_map_proto_rawDescData)
	})
	return edgelq_applications_proto_v1_config_map_proto_rawDescData
}

var edgelq_applications_proto_v1_config_map_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var edgelq_applications_proto_v1_config_map_proto_goTypes = []interface{}{
	(*ConfigMap)(nil), // 0: ntt.applications.v1.ConfigMap
	nil,               // 1: ntt.applications.v1.ConfigMap.DataEntry
	nil,               // 2: ntt.applications.v1.ConfigMap.BinaryDataEntry
	(*meta.Meta)(nil), // 3: goten.types.Meta
}
var edgelq_applications_proto_v1_config_map_proto_depIdxs = []int32{
	3, // 0: ntt.applications.v1.ConfigMap.metadata:type_name -> goten.types.Meta
	1, // 1: ntt.applications.v1.ConfigMap.data:type_name -> ntt.applications.v1.ConfigMap.DataEntry
	2, // 2: ntt.applications.v1.ConfigMap.binary_data:type_name -> ntt.applications.v1.ConfigMap.BinaryDataEntry
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { edgelq_applications_proto_v1_config_map_proto_init() }
func edgelq_applications_proto_v1_config_map_proto_init() {
	if edgelq_applications_proto_v1_config_map_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		edgelq_applications_proto_v1_config_map_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigMap); i {
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
			RawDescriptor: edgelq_applications_proto_v1_config_map_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           edgelq_applications_proto_v1_config_map_proto_goTypes,
		DependencyIndexes: edgelq_applications_proto_v1_config_map_proto_depIdxs,
		MessageInfos:      edgelq_applications_proto_v1_config_map_proto_msgTypes,
	}.Build()
	edgelq_applications_proto_v1_config_map_proto = out.File
	edgelq_applications_proto_v1_config_map_proto_rawDesc = nil
	edgelq_applications_proto_v1_config_map_proto_goTypes = nil
	edgelq_applications_proto_v1_config_map_proto_depIdxs = nil
}
