// Code generated by protoc-gen-goten-go
// File: edgelq/limits/proto/v1alpha2/common.proto
// DO NOT EDIT!!!

package common

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
	meta_resource "github.com/cloudwan/edgelq-sdk/meta/resources/v1alpha2/resource"
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
	_ = &meta_resource.Resource{}
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Allowance informs how many instances of resource for given type are
// permitted.
type Allowance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Reference to resource
	Resource *meta_resource.Reference `protobuf:"bytes,1,opt,customtype=Reference,name=resource,proto3" json:"resource,omitempty"`
	// Limit or extension for given resource.
	Value int64 `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *Allowance) Reset() {
	*m = Allowance{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_limits_proto_v1alpha2_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *Allowance) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*Allowance) ProtoMessage() {}

func (m *Allowance) ProtoReflect() preflect.Message {
	mi := &edgelq_limits_proto_v1alpha2_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*Allowance) GotenMessage() {}

// Deprecated, Use Allowance.ProtoReflect.Descriptor instead.
func (*Allowance) Descriptor() ([]byte, []int) {
	return edgelq_limits_proto_v1alpha2_common_proto_rawDescGZIP(), []int{0}
}

func (m *Allowance) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *Allowance) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *Allowance) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *Allowance) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *Allowance) GetResource() *meta_resource.Reference {
	if m != nil {
		return m.Resource
	}
	return nil
}

func (m *Allowance) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return int64(0)
}

func (m *Allowance) SetResource(fv *meta_resource.Reference) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Resource", "Allowance"))
	}
	m.Resource = fv
}

func (m *Allowance) SetValue(fv int64) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Value", "Allowance"))
	}
	m.Value = fv
}

// RegionalDistribution shows distribution of certain resource allowances
// across regions.
type RegionalDistribution struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Reference to resource
	Resource *meta_resource.Reference `protobuf:"bytes,1,opt,customtype=Reference,name=resource,proto3" json:"resource,omitempty"`
	// Limit per region.
	LimitsByRegion map[string]int64 `protobuf:"bytes,2,rep,name=limits_by_region,json=limitsByRegion,proto3" json:"limits_by_region,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (m *RegionalDistribution) Reset() {
	*m = RegionalDistribution{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_limits_proto_v1alpha2_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *RegionalDistribution) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*RegionalDistribution) ProtoMessage() {}

func (m *RegionalDistribution) ProtoReflect() preflect.Message {
	mi := &edgelq_limits_proto_v1alpha2_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*RegionalDistribution) GotenMessage() {}

// Deprecated, Use RegionalDistribution.ProtoReflect.Descriptor instead.
func (*RegionalDistribution) Descriptor() ([]byte, []int) {
	return edgelq_limits_proto_v1alpha2_common_proto_rawDescGZIP(), []int{1}
}

func (m *RegionalDistribution) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *RegionalDistribution) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *RegionalDistribution) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *RegionalDistribution) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *RegionalDistribution) GetResource() *meta_resource.Reference {
	if m != nil {
		return m.Resource
	}
	return nil
}

func (m *RegionalDistribution) GetLimitsByRegion() map[string]int64 {
	if m != nil {
		return m.LimitsByRegion
	}
	return nil
}

func (m *RegionalDistribution) SetResource(fv *meta_resource.Reference) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Resource", "RegionalDistribution"))
	}
	m.Resource = fv
}

func (m *RegionalDistribution) SetLimitsByRegion(fv map[string]int64) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "LimitsByRegion", "RegionalDistribution"))
	}
	m.LimitsByRegion = fv
}

var edgelq_limits_proto_v1alpha2_common_proto preflect.FileDescriptor

var edgelq_limits_proto_v1alpha2_common_proto_rawDesc = []byte{
	0x0a, 0x29, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x6e, 0x74, 0x74,
	0x2e, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32,
	0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x29, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x61, 0x0a, 0x09,
	0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x08, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x22, 0xb2, 0xda, 0x21,
	0x1e, 0x12, 0x1c, 0x0a, 0x18, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x10, 0x01, 0x52,
	0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0x82, 0x02, 0x0a, 0x14, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x44, 0x69, 0x73, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3e, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x22, 0xb2, 0xda, 0x21, 0x1e,
	0x12, 0x1c, 0x0a, 0x18, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x10, 0x01, 0x52, 0x08,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x67, 0x0a, 0x10, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x73, 0x5f, 0x62, 0x79, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x61,
	0x6c, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x73, 0x42, 0x79, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x0e, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x42, 0x79, 0x52, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x1a, 0x41, 0x0a, 0x13, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x42, 0x79, 0x52, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x42, 0x73, 0xe8, 0xde, 0x21, 0x01, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e,
	0x6e, 0x74, 0x74, 0x2e, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x42, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c,
	0x71, 0x2f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	edgelq_limits_proto_v1alpha2_common_proto_rawDescOnce sync.Once
	edgelq_limits_proto_v1alpha2_common_proto_rawDescData = edgelq_limits_proto_v1alpha2_common_proto_rawDesc
)

func edgelq_limits_proto_v1alpha2_common_proto_rawDescGZIP() []byte {
	edgelq_limits_proto_v1alpha2_common_proto_rawDescOnce.Do(func() {
		edgelq_limits_proto_v1alpha2_common_proto_rawDescData = protoimpl.X.CompressGZIP(edgelq_limits_proto_v1alpha2_common_proto_rawDescData)
	})
	return edgelq_limits_proto_v1alpha2_common_proto_rawDescData
}

var edgelq_limits_proto_v1alpha2_common_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var edgelq_limits_proto_v1alpha2_common_proto_goTypes = []interface{}{
	(*Allowance)(nil),            // 0: ntt.limits.v1alpha2.Allowance
	(*RegionalDistribution)(nil), // 1: ntt.limits.v1alpha2.RegionalDistribution
	nil,                          // 2: ntt.limits.v1alpha2.RegionalDistribution.LimitsByRegionEntry
}
var edgelq_limits_proto_v1alpha2_common_proto_depIdxs = []int32{
	2, // 0: ntt.limits.v1alpha2.RegionalDistribution.limits_by_region:type_name -> ntt.limits.v1alpha2.RegionalDistribution.LimitsByRegionEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { edgelq_limits_proto_v1alpha2_common_proto_init() }
func edgelq_limits_proto_v1alpha2_common_proto_init() {
	if edgelq_limits_proto_v1alpha2_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		edgelq_limits_proto_v1alpha2_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Allowance); i {
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
		edgelq_limits_proto_v1alpha2_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegionalDistribution); i {
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
			RawDescriptor: edgelq_limits_proto_v1alpha2_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           edgelq_limits_proto_v1alpha2_common_proto_goTypes,
		DependencyIndexes: edgelq_limits_proto_v1alpha2_common_proto_depIdxs,
		MessageInfos:      edgelq_limits_proto_v1alpha2_common_proto_msgTypes,
	}.Build()
	edgelq_limits_proto_v1alpha2_common_proto = out.File
	edgelq_limits_proto_v1alpha2_common_proto_rawDesc = nil
	edgelq_limits_proto_v1alpha2_common_proto_goTypes = nil
	edgelq_limits_proto_v1alpha2_common_proto_depIdxs = nil
}
