// Code generated by protoc-gen-goten-go
// File: edgelq/limits/proto/v1alpha2/plan.proto
// DO NOT EDIT!!!

package plan

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
	common "github.com/cloudwan/edgelq-sdk/limits/resources/v1alpha2/common"
	meta_service "github.com/cloudwan/edgelq-sdk/meta/resources/v1alpha2/service"
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
	_ = &common.Allowance{}
	_ = &meta_service.Service{}
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Plan is a bundle with set of limits for various resources.
// Those resources should be mostly from service plan references,
// but it can also include resources from imported services too.
type Plan struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Name of Plan
	Name        *Name  `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty" firestore:"displayName"`
	// Primary service to which this plan applies
	Service *meta_service.Reference `protobuf:"bytes,3,opt,customtype=Reference,name=service,proto3" json:"service,omitempty" firestore:"service"`
	// List of limits per each resource. Resources may belong to different
	// services (dependant services)
	ResourceLimits []*common.Allowance `protobuf:"bytes,4,rep,name=resource_limits,json=resourceLimits,proto3" json:"resource_limits,omitempty" firestore:"resourceLimits"`
	// Metadata
	Metadata *ntt_meta.Meta `protobuf:"bytes,5,opt,name=metadata,proto3" json:"metadata,omitempty" firestore:"metadata"`
}

func (m *Plan) Reset() {
	*m = Plan{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_limits_proto_v1alpha2_plan_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *Plan) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*Plan) ProtoMessage() {}

func (m *Plan) ProtoReflect() preflect.Message {
	mi := &edgelq_limits_proto_v1alpha2_plan_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*Plan) GotenMessage() {}

// Deprecated, Use Plan.ProtoReflect.Descriptor instead.
func (*Plan) Descriptor() ([]byte, []int) {
	return edgelq_limits_proto_v1alpha2_plan_proto_rawDescGZIP(), []int{0}
}

func (m *Plan) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *Plan) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *Plan) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *Plan) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *Plan) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *Plan) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *Plan) GetService() *meta_service.Reference {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *Plan) GetResourceLimits() []*common.Allowance {
	if m != nil {
		return m.ResourceLimits
	}
	return nil
}

func (m *Plan) GetMetadata() *ntt_meta.Meta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Plan) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "Plan"))
	}
	m.Name = fv
}

func (m *Plan) SetDisplayName(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "DisplayName", "Plan"))
	}
	m.DisplayName = fv
}

func (m *Plan) SetService(fv *meta_service.Reference) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Service", "Plan"))
	}
	m.Service = fv
}

func (m *Plan) SetResourceLimits(fv []*common.Allowance) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ResourceLimits", "Plan"))
	}
	m.ResourceLimits = fv
}

func (m *Plan) SetMetadata(fv *ntt_meta.Meta) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Metadata", "Plan"))
	}
	m.Metadata = fv
}

var edgelq_limits_proto_v1alpha2_plan_proto preflect.FileDescriptor

var edgelq_limits_proto_v1alpha2_plan_proto_rawDesc = []byte{
	0x0a, 0x27, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x70,
	0x6c, 0x61, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x6e, 0x74, 0x74, 0x2e, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x1a, 0x19,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74,
	0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67,
	0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65,
	0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f,
	0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x74,
	0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x65, 0x64, 0x67,
	0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x65,
	0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x34, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6d, 0x65, 0x74, 0x61,
	0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x72, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x65, 0x64, 0x67, 0x65, 0x6c,
	0x71, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6d,
	0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x65, 0x64, 0x67, 0x65, 0x6c,
	0x71, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x32, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xeb,
	0x03, 0x0a, 0x04, 0x50, 0x6c, 0x61, 0x6e, 0x12, 0x20, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xb2, 0xda, 0x21, 0x08, 0x0a, 0x06, 0x0a, 0x04, 0x50,
	0x6c, 0x61, 0x6e, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73,
	0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x07,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x21, 0xb2,
	0xda, 0x21, 0x1d, 0x12, 0x1b, 0x0a, 0x19, 0x6e, 0x74, 0x74, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x0f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x61, 0x6e,
	0x63, 0x65, 0x52, 0x0e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x73, 0x12, 0x2b, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x3a,
	0xea, 0x01, 0xea, 0x41, 0x26, 0x0a, 0x16, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x2e, 0x65, 0x64,
	0x67, 0x65, 0x6c, 0x71, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x6c, 0x61, 0x6e, 0x12, 0x0c, 0x70,
	0x6c, 0x61, 0x6e, 0x73, 0x2f, 0x7b, 0x70, 0x6c, 0x61, 0x6e, 0x7d, 0x92, 0xd9, 0x21, 0x75, 0x0a,
	0x05, 0x70, 0x6c, 0x61, 0x6e, 0x73, 0x12, 0x05, 0x70, 0x6c, 0x61, 0x6e, 0x73, 0x22, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x4a, 0x25, 0x08, 0x02, 0x12, 0x06, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x0e, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x09, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4a, 0x38, 0x08, 0x03, 0x12, 0x06,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61,
	0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x09, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x11, 0x0a, 0x0f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x73, 0xda, 0x94, 0x23, 0x08, 0x12, 0x06, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0xc2, 0x85, 0x2c, 0x38, 0x22, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x0c, 0x64, 0x69, 0x73, 0x70,
	0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x22, 0x0f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x73, 0x42, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0x9d, 0x03, 0xe8,
	0xde, 0x21, 0x01, 0xd2, 0xff, 0xd0, 0x02, 0x43, 0x0a, 0x0a, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x12, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71,
	0x2f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x0a, 0x1a, 0x63, 0x6f, 0x6d,
	0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x2e, 0x70, 0x62, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x42, 0x09, 0x50, 0x6c, 0x61, 0x6e, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71,
	0x2f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x3b,
	0x70, 0x6c, 0x61, 0x6e, 0xd2, 0x84, 0xd1, 0x02, 0x49, 0x0a, 0x0d, 0x61, 0x75, 0x64, 0x69, 0x74,
	0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x12, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64,
	0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x2f, 0x61, 0x75, 0x64, 0x69,
	0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x32, 0xf2, 0x85, 0xd1, 0x02, 0x4b, 0x0a, 0x0e, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x64, 0x62,
	0x5f, 0x73, 0x79, 0x6e, 0x63, 0x65, 0x72, 0x12, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67,
	0x65, 0x6c, 0x71, 0x2f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x2f, 0x64, 0x62, 0x5f, 0x73, 0x79,
	0x6e, 0x63, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x70, 0x6c,
	0x61, 0x6e, 0xa2, 0x80, 0xd1, 0x02, 0x45, 0x0a, 0x0b, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x12, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71,
	0x2f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	edgelq_limits_proto_v1alpha2_plan_proto_rawDescOnce sync.Once
	edgelq_limits_proto_v1alpha2_plan_proto_rawDescData = edgelq_limits_proto_v1alpha2_plan_proto_rawDesc
)

func edgelq_limits_proto_v1alpha2_plan_proto_rawDescGZIP() []byte {
	edgelq_limits_proto_v1alpha2_plan_proto_rawDescOnce.Do(func() {
		edgelq_limits_proto_v1alpha2_plan_proto_rawDescData = protoimpl.X.CompressGZIP(edgelq_limits_proto_v1alpha2_plan_proto_rawDescData)
	})
	return edgelq_limits_proto_v1alpha2_plan_proto_rawDescData
}

var edgelq_limits_proto_v1alpha2_plan_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var edgelq_limits_proto_v1alpha2_plan_proto_goTypes = []interface{}{
	(*Plan)(nil),             // 0: ntt.limits.v1alpha2.Plan
	(*common.Allowance)(nil), // 1: ntt.limits.v1alpha2.Allowance
	(*ntt_meta.Meta)(nil),    // 2: ntt.types.Meta
}
var edgelq_limits_proto_v1alpha2_plan_proto_depIdxs = []int32{
	1, // 0: ntt.limits.v1alpha2.Plan.resource_limits:type_name -> ntt.limits.v1alpha2.Allowance
	2, // 1: ntt.limits.v1alpha2.Plan.metadata:type_name -> ntt.types.Meta
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { edgelq_limits_proto_v1alpha2_plan_proto_init() }
func edgelq_limits_proto_v1alpha2_plan_proto_init() {
	if edgelq_limits_proto_v1alpha2_plan_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		edgelq_limits_proto_v1alpha2_plan_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Plan); i {
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
			RawDescriptor: edgelq_limits_proto_v1alpha2_plan_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           edgelq_limits_proto_v1alpha2_plan_proto_goTypes,
		DependencyIndexes: edgelq_limits_proto_v1alpha2_plan_proto_depIdxs,
		MessageInfos:      edgelq_limits_proto_v1alpha2_plan_proto_msgTypes,
	}.Build()
	edgelq_limits_proto_v1alpha2_plan_proto = out.File
	edgelq_limits_proto_v1alpha2_plan_proto_rawDesc = nil
	edgelq_limits_proto_v1alpha2_plan_proto_goTypes = nil
	edgelq_limits_proto_v1alpha2_plan_proto_depIdxs = nil
}