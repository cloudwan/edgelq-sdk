// Code generated by protoc-gen-goten-go
// File: edgelq/proxies/proto/v1alpha2/project.proto
// DO NOT EDIT!!!

package project

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
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
	multi_region_policy "github.com/cloudwan/edgelq-sdk/common/types/multi_region_policy"
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
	_ = &ntt_meta.Meta{}
	_ = &multi_region_policy.MultiRegionPolicy{}
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Project Resource
type Project struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Name of Project
	Name *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// Metadata
	Metadata *ntt_meta.Meta `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty" firestore:"metadata"`
	// Multi region policy
	MultiRegionPolicy *multi_region_policy.MultiRegionPolicy `protobuf:"bytes,3,opt,name=multi_region_policy,json=multiRegionPolicy,proto3" json:"multi_region_policy,omitempty" firestore:"multiRegionPolicy"`
}

func (m *Project) Reset() {
	*m = Project{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_proxies_proto_v1alpha2_project_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *Project) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*Project) ProtoMessage() {}

func (m *Project) ProtoReflect() preflect.Message {
	mi := &edgelq_proxies_proto_v1alpha2_project_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*Project) GotenMessage() {}

// Deprecated, Use Project.ProtoReflect.Descriptor instead.
func (*Project) Descriptor() ([]byte, []int) {
	return edgelq_proxies_proto_v1alpha2_project_proto_rawDescGZIP(), []int{0}
}

func (m *Project) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *Project) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *Project) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *Project) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *Project) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *Project) GetMetadata() *ntt_meta.Meta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Project) GetMultiRegionPolicy() *multi_region_policy.MultiRegionPolicy {
	if m != nil {
		return m.MultiRegionPolicy
	}
	return nil
}

func (m *Project) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "Project"))
	}
	m.Name = fv
}

func (m *Project) SetMetadata(fv *ntt_meta.Meta) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Metadata", "Project"))
	}
	m.Metadata = fv
}

func (m *Project) SetMultiRegionPolicy(fv *multi_region_policy.MultiRegionPolicy) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "MultiRegionPolicy", "Project"))
	}
	m.MultiRegionPolicy = fv
}

var edgelq_proxies_proto_v1alpha2_project_proto preflect.FileDescriptor

var edgelq_proxies_proto_v1alpha2_project_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x70, 0x72, 0x6f, 0x78, 0x69, 0x65, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x6e,
	0x74, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x32, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20,
	0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1e, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x24, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x34, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6d,
	0x65, 0x74, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x72, 0x73, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f,
	0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x65, 0x64,
	0x67, 0x65, 0x6c, 0x71, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2d, 0x65, 0x64,
	0x67, 0x65, 0x6c, 0x71, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbc, 0x02, 0x0a, 0x07,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x23, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0f, 0xb2, 0xda, 0x21, 0x0b, 0x0a, 0x09, 0x0a, 0x07, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2b, 0x0a, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x56, 0x0a, 0x13, 0x6d, 0x75, 0x6c,
	0x74, 0x69, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x42, 0x08, 0xca, 0xc6, 0x27, 0x04, 0x62, 0x02, 0x08, 0x01, 0x52, 0x11,
	0x6d, 0x75, 0x6c, 0x74, 0x69, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x3a, 0x86, 0x01, 0xea, 0x41, 0x30, 0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x78, 0x69, 0x65, 0x73,
	0x2e, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x12, 0x12, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x92, 0xd9, 0x21, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x22,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0xda, 0x94, 0x23, 0x08, 0x12, 0x06, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0xc2, 0x85, 0x2c, 0x25, 0x22, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x13, 0x6d, 0x75, 0x6c,
	0x74, 0x69, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x42, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0xbe, 0x03, 0xe8, 0xde, 0x21,
	0x01, 0xd2, 0xff, 0xd0, 0x02, 0x4a, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c,
	0x71, 0x2f, 0x70, 0x72, 0x6f, 0x78, 0x69, 0x65, 0x73, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x69, 0x65,
	0x73, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x42, 0x0c, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77,
	0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x70, 0x72, 0x6f, 0x78, 0x69, 0x65,
	0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x32, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x3b, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0xd2, 0x84, 0xd1, 0x02, 0x4a, 0x0a, 0x0d, 0x61, 0x75, 0x64, 0x69, 0x74,
	0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x12, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64,
	0x67, 0x65, 0x6c, 0x71, 0x2f, 0x70, 0x72, 0x6f, 0x78, 0x69, 0x65, 0x73, 0x2f, 0x61, 0x75, 0x64,
	0x69, 0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x32, 0xf2, 0x85, 0xd1, 0x02, 0x52, 0x0a, 0x11, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x5f, 0x64, 0x62, 0x5f, 0x73, 0x79, 0x6e, 0x63, 0x65, 0x72, 0x12, 0x3d, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e,
	0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x70, 0x72, 0x6f, 0x78, 0x69, 0x65, 0x73, 0x2f,
	0x64, 0x62, 0x5f, 0x73, 0x79, 0x6e, 0x63, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x32, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0xa2, 0x80, 0xd1, 0x02, 0x4c, 0x0a,
	0x0e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12,
	0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x70, 0x72, 0x6f, 0x78,
	0x69, 0x65, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x32, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	edgelq_proxies_proto_v1alpha2_project_proto_rawDescOnce sync.Once
	edgelq_proxies_proto_v1alpha2_project_proto_rawDescData = edgelq_proxies_proto_v1alpha2_project_proto_rawDesc
)

func edgelq_proxies_proto_v1alpha2_project_proto_rawDescGZIP() []byte {
	edgelq_proxies_proto_v1alpha2_project_proto_rawDescOnce.Do(func() {
		edgelq_proxies_proto_v1alpha2_project_proto_rawDescData = protoimpl.X.CompressGZIP(edgelq_proxies_proto_v1alpha2_project_proto_rawDescData)
	})
	return edgelq_proxies_proto_v1alpha2_project_proto_rawDescData
}

var edgelq_proxies_proto_v1alpha2_project_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var edgelq_proxies_proto_v1alpha2_project_proto_goTypes = []interface{}{
	(*Project)(nil),       // 0: ntt.proxies.v1alpha2.Project
	(*ntt_meta.Meta)(nil), // 1: ntt.types.Meta
	(*multi_region_policy.MultiRegionPolicy)(nil), // 2: ntt.types.MultiRegionPolicy
}
var edgelq_proxies_proto_v1alpha2_project_proto_depIdxs = []int32{
	1, // 0: ntt.proxies.v1alpha2.Project.metadata:type_name -> ntt.types.Meta
	2, // 1: ntt.proxies.v1alpha2.Project.multi_region_policy:type_name -> ntt.types.MultiRegionPolicy
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { edgelq_proxies_proto_v1alpha2_project_proto_init() }
func edgelq_proxies_proto_v1alpha2_project_proto_init() {
	if edgelq_proxies_proto_v1alpha2_project_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		edgelq_proxies_proto_v1alpha2_project_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Project); i {
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
			RawDescriptor: edgelq_proxies_proto_v1alpha2_project_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           edgelq_proxies_proto_v1alpha2_project_proto_goTypes,
		DependencyIndexes: edgelq_proxies_proto_v1alpha2_project_proto_depIdxs,
		MessageInfos:      edgelq_proxies_proto_v1alpha2_project_proto_msgTypes,
	}.Build()
	edgelq_proxies_proto_v1alpha2_project_proto = out.File
	edgelq_proxies_proto_v1alpha2_project_proto_rawDesc = nil
	edgelq_proxies_proto_v1alpha2_project_proto_goTypes = nil
	edgelq_proxies_proto_v1alpha2_project_proto_depIdxs = nil
}
