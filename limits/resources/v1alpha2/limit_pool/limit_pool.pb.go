// Code generated by protoc-gen-goten-go
// File: edgelq/limits/proto/v1alpha2/limit_pool.proto
// DO NOT EDIT!!!

package limit_pool

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
	iam_organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/organization"
	meta_resource "github.com/cloudwan/edgelq-sdk/meta/resources/v1alpha2/resource"
	meta_service "github.com/cloudwan/edgelq-sdk/meta/resources/v1alpha2/service"
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
	_ = &iam_organization.Organization{}
	_ = &meta_resource.Resource{}
	_ = &meta_service.Service{}
	_ = &meta.Meta{}
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// LimitPool represents a pool from which next LimitPool
// or final Limit instances are allocated from. LimitPool
// can either belong to a system (first pool, without parent) or
// organization (which got pool from system or parent organization).
// Size of child limit or limit pools must not exceed parent
// limit pool.
// LimitPool is in a way similar to Limit - it has scope and
// associated resource but:
// * LimitPool is non-final object. It cannot be used to
// allocate resources directly. It need child Limit first. So
// field with values in LimitPool are configured size, active
// size and reserved. Fields in Limit are configured limit,
// active limit and usage - they need different interpretation.
// * LimitPool is not regional, because it may feed limits in
// other regions. Limit has extra tracker in their final region
// for state synchronization purpose.
// Because LimitPool instances within one scope (system, organization)
// are managed by primary region of that scope, changes in them
// are done with only single transaction.
// LimitPool instances cannot be created or deleted. They cannot be
// also be modified (there are only some state fields controlled
// by system). LimitPools are created/modified along with
// PlanAssignment instances.
type LimitPool struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Name of LimitPool
	Name *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// Referenced service
	Service *meta_service.Reference `protobuf:"bytes,2,opt,customtype=Reference,name=service,proto3" json:"service,omitempty" firestore:"service"`
	// Referenced resource type
	Resource *meta_resource.Reference `protobuf:"bytes,3,opt,customtype=Reference,name=resource,proto3" json:"resource,omitempty" firestore:"resource"`
	// Configured size of pool according to PlanAssignment instances
	// belonging to same scope (system or organization).
	ConfiguredSize int64 `protobuf:"varint,4,opt,name=configured_size,json=configuredSize,proto3" json:"configured_size,omitempty" firestore:"configuredSize"`
	// Active pool size. It will be normally equal to configured size.
	// However, if configured size goes down below reserved value,
	// then active size will be equal to that "reserved".
	// It will also prevent parent LimitPool (if any) from getting
	// limit value back.
	ActiveSize int64 `protobuf:"varint,5,opt,name=active_size,json=activeSize,proto3" json:"active_size,omitempty" firestore:"activeSize"`
	// Configured OR Active Size of all child Limit and LimitPool instances -
	// whichever is bigger.
	Reserved int64 `protobuf:"varint,6,opt,name=reserved,proto3" json:"reserved,omitempty" firestore:"reserved"`
	// LimitPool source that gave birth to this LimitPool. Any change
	// in configured/active value in current LimitPool increases/decreases
	// reserved field in source LimitPool.
	Source *Reference `protobuf:"bytes,7,opt,customtype=Reference,name=source,proto3" json:"source,omitempty" firestore:"source"`
	// Metadata
	Metadata *meta.Meta `protobuf:"bytes,8,opt,name=metadata,proto3" json:"metadata,omitempty" firestore:"metadata"`
}

func (m *LimitPool) Reset() {
	*m = LimitPool{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_limits_proto_v1alpha2_limit_pool_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *LimitPool) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*LimitPool) ProtoMessage() {}

func (m *LimitPool) ProtoReflect() preflect.Message {
	mi := &edgelq_limits_proto_v1alpha2_limit_pool_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*LimitPool) GotenMessage() {}

// Deprecated, Use LimitPool.ProtoReflect.Descriptor instead.
func (*LimitPool) Descriptor() ([]byte, []int) {
	return edgelq_limits_proto_v1alpha2_limit_pool_proto_rawDescGZIP(), []int{0}
}

func (m *LimitPool) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *LimitPool) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *LimitPool) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *LimitPool) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *LimitPool) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *LimitPool) GetService() *meta_service.Reference {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *LimitPool) GetResource() *meta_resource.Reference {
	if m != nil {
		return m.Resource
	}
	return nil
}

func (m *LimitPool) GetConfiguredSize() int64 {
	if m != nil {
		return m.ConfiguredSize
	}
	return int64(0)
}

func (m *LimitPool) GetActiveSize() int64 {
	if m != nil {
		return m.ActiveSize
	}
	return int64(0)
}

func (m *LimitPool) GetReserved() int64 {
	if m != nil {
		return m.Reserved
	}
	return int64(0)
}

func (m *LimitPool) GetSource() *Reference {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *LimitPool) GetMetadata() *meta.Meta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *LimitPool) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "LimitPool"))
	}
	m.Name = fv
}

func (m *LimitPool) SetService(fv *meta_service.Reference) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Service", "LimitPool"))
	}
	m.Service = fv
}

func (m *LimitPool) SetResource(fv *meta_resource.Reference) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Resource", "LimitPool"))
	}
	m.Resource = fv
}

func (m *LimitPool) SetConfiguredSize(fv int64) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ConfiguredSize", "LimitPool"))
	}
	m.ConfiguredSize = fv
}

func (m *LimitPool) SetActiveSize(fv int64) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ActiveSize", "LimitPool"))
	}
	m.ActiveSize = fv
}

func (m *LimitPool) SetReserved(fv int64) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Reserved", "LimitPool"))
	}
	m.Reserved = fv
}

func (m *LimitPool) SetSource(fv *Reference) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Source", "LimitPool"))
	}
	m.Source = fv
}

func (m *LimitPool) SetMetadata(fv *meta.Meta) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Metadata", "LimitPool"))
	}
	m.Metadata = fv
}

var edgelq_limits_proto_v1alpha2_limit_pool_proto preflect.FileDescriptor

var edgelq_limits_proto_v1alpha2_limit_pool_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x5f, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x13, 0x6e, 0x74, 0x74, 0x2e, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x32, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
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
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71,
	0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x69, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x65, 0x64, 0x67, 0x65,
	0x6c, 0x71, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x16, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6d, 0x65,
	0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71,
	0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x32, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x2c, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x6f, 0x72,
	0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x27, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb3, 0x07, 0x0a, 0x09, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x50, 0x6f, 0x6f, 0x6c, 0x12, 0x25, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x11, 0xb2, 0xda, 0x21, 0x0d, 0x0a, 0x0b, 0x0a, 0x09, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x50, 0x6f, 0x6f, 0x6c, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3f,
	0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x25, 0xf0, 0xd9, 0x21, 0x01, 0xb2, 0xda, 0x21, 0x1d, 0x12, 0x1b, 0x0a, 0x17, 0x6d, 0x65, 0x74,
	0x61, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x10, 0x01, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x42, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x26, 0xf0, 0xd9, 0x21, 0x01, 0xb2, 0xda, 0x21, 0x1e, 0x12, 0x1c, 0x0a, 0x18, 0x6d,
	0x65, 0x74, 0x61, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x10, 0x01, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x12, 0x2d, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65,
	0x64, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x42, 0x04, 0xf0, 0xd9,
	0x21, 0x01, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x64, 0x53, 0x69,
	0x7a, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x12,
	0x2f, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x17, 0xf0, 0xd9, 0x21, 0x01, 0xb2, 0xda, 0x21, 0x0f, 0x12, 0x0d, 0x0a, 0x09, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x50, 0x6f, 0x6f, 0x6c, 0x10, 0x01, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x12, 0x2d, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x3a,
	0xad, 0x04, 0xea, 0x41, 0x6c, 0x0a, 0x1b, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x2e, 0x65, 0x64,
	0x67, 0x65, 0x6c, 0x71, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x50, 0x6f,
	0x6f, 0x6c, 0x12, 0x17, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x50, 0x6f, 0x6f, 0x6c, 0x73, 0x2f, 0x7b,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x70, 0x6f, 0x6f, 0x6c, 0x7d, 0x12, 0x34, 0x6f, 0x72, 0x67,
	0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6f, 0x72, 0x67, 0x61,
	0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x50,
	0x6f, 0x6f, 0x6c, 0x73, 0x2f, 0x7b, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x70, 0x6f, 0x6f, 0x6c,
	0x7d, 0x92, 0xd9, 0x21, 0xfa, 0x01, 0x0a, 0x0a, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x50, 0x6f, 0x6f,
	0x6c, 0x73, 0x12, 0x0a, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x50, 0x6f, 0x6f, 0x6c, 0x73, 0x1a, 0x12,
	0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x6f,
	0x6e, 0x65, 0x1a, 0x1b, 0x69, 0x61, 0x6d, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2a,
	0x25, 0x5b, 0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a, 0x30, 0x2d, 0x39, 0x2d, 0x2e, 0x5d, 0x7b, 0x31,
	0x2c, 0x31, 0x32, 0x38, 0x7d, 0x5c, 0x2f, 0x5b, 0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a, 0x5d, 0x7b,
	0x31, 0x2c, 0x31, 0x32, 0x38, 0x7d, 0x38, 0x05, 0x42, 0x35, 0x08, 0x02, 0x12, 0x06, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x12, 0x11, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x64, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x12, 0x0a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x42,
	0x4f, 0x08, 0x03, 0x12, 0x06, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x09, 0x0a, 0x07, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x0a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x12, 0x11, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x64,
	0x5f, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x0d, 0x0a, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x12, 0x0a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64,
	0xb2, 0xdf, 0x21, 0x52, 0x0a, 0x50, 0x0a, 0x4e, 0x0a, 0x06, 0x62, 0x79, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x1a, 0x29, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x50, 0x6f, 0x6f, 0x6c, 0x73, 0x2f,
	0x2d, 0x1a, 0x0c, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x50, 0x6f, 0x6f, 0x6c, 0x73, 0x2f, 0x2d, 0x2a,
	0x05, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0xda, 0x94, 0x23, 0x08, 0x12, 0x06, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0xe2, 0xde, 0x21, 0x02, 0x08, 0x02, 0xc2, 0x85, 0x2c, 0x53, 0x22, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x08, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x0f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65,
	0x64, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2a, 0x0b,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x2a, 0x08, 0x72, 0x65, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x64, 0x42, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42,
	0xa8, 0x02, 0xe8, 0xde, 0x21, 0x01, 0xd2, 0xff, 0xd0, 0x02, 0x4f, 0x0a, 0x10, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x5f, 0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x3b, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77,
	0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73,
	0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x70, 0x6f, 0x6f, 0x6c, 0xa2, 0x80, 0xd1, 0x02, 0x51, 0x0a,
	0x11, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x12, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x70, 0x6f, 0x6f, 0x6c,
	0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73,
	0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x42, 0x0e, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x50, 0x6f, 0x6f, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4a,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x32, 0x2f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x70, 0x6f, 0x6f, 0x6c, 0x3b,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x70, 0x6f, 0x6f, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	edgelq_limits_proto_v1alpha2_limit_pool_proto_rawDescOnce sync.Once
	edgelq_limits_proto_v1alpha2_limit_pool_proto_rawDescData = edgelq_limits_proto_v1alpha2_limit_pool_proto_rawDesc
)

func edgelq_limits_proto_v1alpha2_limit_pool_proto_rawDescGZIP() []byte {
	edgelq_limits_proto_v1alpha2_limit_pool_proto_rawDescOnce.Do(func() {
		edgelq_limits_proto_v1alpha2_limit_pool_proto_rawDescData = protoimpl.X.CompressGZIP(edgelq_limits_proto_v1alpha2_limit_pool_proto_rawDescData)
	})
	return edgelq_limits_proto_v1alpha2_limit_pool_proto_rawDescData
}

var edgelq_limits_proto_v1alpha2_limit_pool_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var edgelq_limits_proto_v1alpha2_limit_pool_proto_goTypes = []interface{}{
	(*LimitPool)(nil), // 0: ntt.limits.v1alpha2.LimitPool
	(*meta.Meta)(nil), // 1: goten.types.Meta
}
var edgelq_limits_proto_v1alpha2_limit_pool_proto_depIdxs = []int32{
	1, // 0: ntt.limits.v1alpha2.LimitPool.metadata:type_name -> goten.types.Meta
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { edgelq_limits_proto_v1alpha2_limit_pool_proto_init() }
func edgelq_limits_proto_v1alpha2_limit_pool_proto_init() {
	if edgelq_limits_proto_v1alpha2_limit_pool_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		edgelq_limits_proto_v1alpha2_limit_pool_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LimitPool); i {
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
			RawDescriptor: edgelq_limits_proto_v1alpha2_limit_pool_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           edgelq_limits_proto_v1alpha2_limit_pool_proto_goTypes,
		DependencyIndexes: edgelq_limits_proto_v1alpha2_limit_pool_proto_depIdxs,
		MessageInfos:      edgelq_limits_proto_v1alpha2_limit_pool_proto_msgTypes,
	}.Build()
	edgelq_limits_proto_v1alpha2_limit_pool_proto = out.File
	edgelq_limits_proto_v1alpha2_limit_pool_proto_rawDesc = nil
	edgelq_limits_proto_v1alpha2_limit_pool_proto_goTypes = nil
	edgelq_limits_proto_v1alpha2_limit_pool_proto_depIdxs = nil
}
