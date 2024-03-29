// Code generated by protoc-gen-goten-go
// File: edgelq/limits/proto/v1/limit.proto
// DO NOT EDIT!!!

package limit

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
	iam_project "github.com/cloudwan/edgelq-sdk/iam/resources/v1/project"
	limit_pool "github.com/cloudwan/edgelq-sdk/limits/resources/v1/limit_pool"
	meta_resource "github.com/cloudwan/goten-sdk/meta-service/resources/v1/resource"
	meta_service "github.com/cloudwan/goten-sdk/meta-service/resources/v1/service"
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
	_ = &iam_project.Project{}
	_ = &limit_pool.LimitPool{}
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

// Limit resource represents usage/limit of specific resource
// in specific project and region - however instances of those
// resources are managed by primary region of project. This
// ensures consistency and allows modification of project limits
// in one database transaction.
// Limit instances cannot be created or deleted. They cannot be
// also be modified (there are only some state fields controlled
// by system). Limits are created/modified along with
// PlanAssignment instances.
type Limit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Name of Limit
	// When creating a new instance, this field is optional and if not provided,
	// it will be generated automatically. Last ID segment must conform to the
	// following regex:
	// [a-zA-Z0-9-]{1,128}\\/[a-zA-Z0-9-.]{1,128}\\/[a-zA-Z]{1,128}
	Name *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// Metadata is an object with information like create, update and delete time
	// (for async deleted resources), has user labels/annotations, sharding
	// information, multi-region syncing information and may have non-schema
	// owners (useful for taking ownership of resources belonging to lower level
	// services by higher ones).
	Metadata *meta.Meta `protobuf:"bytes,9,opt,name=metadata,proto3" json:"metadata,omitempty" firestore:"metadata"`
	// Referenced service
	Service *meta_service.Reference `protobuf:"bytes,2,opt,customtype=Reference,name=service,proto3" json:"service,omitempty" firestore:"service"`
	// Referenced resource type
	Resource *meta_resource.Reference `protobuf:"bytes,3,opt,customtype=Reference,name=resource,proto3" json:"resource,omitempty" firestore:"resource"`
	// Region ID to which this limit applies. This may be different
	// compared to default project region, because projects themselves
	// can be multi-regional.
	Region string `protobuf:"bytes,4,opt,name=region,proto3" json:"region,omitempty" firestore:"region"`
	// Configured limit - it is always in sync with limits
	// predicted by PlanAssignment instances.
	ConfiguredLimit int64 `protobuf:"varint,5,opt,name=configured_limit,json=configuredLimit,proto3" json:"configured_limit,omitempty" firestore:"configuredLimit"`
	// Active limit - it is always equal to configured limit with exceptions:
	// * If configured limit has just changed, then active limit will for
	// a moment contain old value before synchronization.
	// * If it turns out that usage increased in the very same moment
	// when configured limit decreased, then active limit will hold
	// old value indicating problem.
	// It is important to note that, because projects may be multi-regional,
	// limits are tracked by their FINAL region (See region field).
	// State fields (active limit and usage) are synced asynchronously.
	// From user perspective this field can be hidden if in line with
	// configured limit.
	ActiveLimit int64 `protobuf:"varint,6,opt,name=active_limit,json=activeLimit,proto3" json:"active_limit,omitempty" firestore:"activeLimit"`
	// Number of resources currently in existence.
	Usage int64 `protobuf:"varint,7,opt,name=usage,proto3" json:"usage,omitempty" firestore:"usage"`
	// LimitPool sources that supplied reservation to this limit. Any change
	// in configured/active value in current Limit increases/decreases
	// reserved field in source LimitPool.
	// Number of sources depends whether resource is regional or not - if
	// resource is regional, it will be only one source pointing to the
	// specific region. Non-regional resources will point to all regions
	// where project is present.
	Sources []*limit_pool.Reference `protobuf:"bytes,8,rep,customtype=Reference,name=sources,proto3" json:"sources,omitempty" firestore:"sources"`
}

func (m *Limit) Reset() {
	*m = Limit{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_limits_proto_v1_limit_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *Limit) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*Limit) ProtoMessage() {}

func (m *Limit) ProtoReflect() preflect.Message {
	mi := &edgelq_limits_proto_v1_limit_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*Limit) GotenMessage() {}

// Deprecated, Use Limit.ProtoReflect.Descriptor instead.
func (*Limit) Descriptor() ([]byte, []int) {
	return edgelq_limits_proto_v1_limit_proto_rawDescGZIP(), []int{0}
}

func (m *Limit) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *Limit) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *Limit) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *Limit) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *Limit) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *Limit) GetMetadata() *meta.Meta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Limit) GetService() *meta_service.Reference {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *Limit) GetResource() *meta_resource.Reference {
	if m != nil {
		return m.Resource
	}
	return nil
}

func (m *Limit) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *Limit) GetConfiguredLimit() int64 {
	if m != nil {
		return m.ConfiguredLimit
	}
	return int64(0)
}

func (m *Limit) GetActiveLimit() int64 {
	if m != nil {
		return m.ActiveLimit
	}
	return int64(0)
}

func (m *Limit) GetUsage() int64 {
	if m != nil {
		return m.Usage
	}
	return int64(0)
}

func (m *Limit) GetSources() []*limit_pool.Reference {
	if m != nil {
		return m.Sources
	}
	return nil
}

func (m *Limit) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "Limit"))
	}
	m.Name = fv
}

func (m *Limit) SetMetadata(fv *meta.Meta) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Metadata", "Limit"))
	}
	m.Metadata = fv
}

func (m *Limit) SetService(fv *meta_service.Reference) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Service", "Limit"))
	}
	m.Service = fv
}

func (m *Limit) SetResource(fv *meta_resource.Reference) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Resource", "Limit"))
	}
	m.Resource = fv
}

func (m *Limit) SetRegion(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Region", "Limit"))
	}
	m.Region = fv
}

func (m *Limit) SetConfiguredLimit(fv int64) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ConfiguredLimit", "Limit"))
	}
	m.ConfiguredLimit = fv
}

func (m *Limit) SetActiveLimit(fv int64) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ActiveLimit", "Limit"))
	}
	m.ActiveLimit = fv
}

func (m *Limit) SetUsage(fv int64) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Usage", "Limit"))
	}
	m.Usage = fv
}

func (m *Limit) SetSources(fv []*limit_pool.Reference) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Sources", "Limit"))
	}
	m.Sources = fv
}

var edgelq_limits_proto_v1_limit_proto preflect.FileDescriptor

var edgelq_limits_proto_v1_limit_proto_rawDesc = []byte{
	0x0a, 0x22, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6e, 0x74, 0x74, 0x2e, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73,
	0x2e, 0x76, 0x31, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
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
	0x6f, 0x1a, 0x24, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f,
	0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x67,
	0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x6d, 0x65, 0x74,
	0x61, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x21, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x5f, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x88, 0x07, 0x0a,
	0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x21, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x0d, 0xb2, 0xda, 0x21, 0x09, 0x0a, 0x07, 0x0a, 0x05, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x08, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67, 0x6f,
	0x74, 0x65, 0x6e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x3e, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x24, 0xf0, 0xd9, 0x21, 0x01, 0xb2,
	0xda, 0x21, 0x1c, 0x12, 0x1a, 0x0a, 0x16, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x67, 0x6f, 0x74, 0x65,
	0x6e, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x10, 0x01, 0x52,
	0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x41, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x25, 0xf0, 0xd9, 0x21, 0x01,
	0xb2, 0xda, 0x21, 0x1d, 0x12, 0x1b, 0x0a, 0x17, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x67, 0x6f, 0x74,
	0x65, 0x6e, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x10,
	0x01, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x06, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xf0, 0xd9, 0x21,
	0x01, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x0a, 0x10, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x64, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x42, 0x04, 0xf0, 0xd9, 0x21, 0x01, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x75, 0x72, 0x65, 0x64, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x75, 0x73, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x75, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x31, 0x0a, 0x07, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x08,
	0x20, 0x03, 0x28, 0x09, 0x42, 0x17, 0xf0, 0xd9, 0x21, 0x01, 0xb2, 0xda, 0x21, 0x0f, 0x12, 0x0d,
	0x0a, 0x09, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x50, 0x6f, 0x6f, 0x6c, 0x10, 0x01, 0x52, 0x07, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x3a, 0xee, 0x03, 0xea, 0x41, 0x3c, 0x0a, 0x17, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x73, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x21, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f,
	0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73,
	0x2f, 0x7b, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x7d, 0x92, 0xd9, 0x21, 0xff, 0x01, 0x0a, 0x06, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x73, 0x12, 0x06, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x1a, 0x16, 0x69,
	0x61, 0x6d, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2a, 0x3a, 0x5b, 0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a, 0x30, 0x2d,
	0x39, 0x2d, 0x5d, 0x7b, 0x31, 0x2c, 0x31, 0x32, 0x38, 0x7d, 0x5c, 0x2f, 0x5b, 0x61, 0x2d, 0x7a,
	0x41, 0x2d, 0x5a, 0x30, 0x2d, 0x39, 0x2d, 0x2e, 0x5d, 0x7b, 0x31, 0x2c, 0x31, 0x32, 0x38, 0x7d,
	0x5c, 0x2f, 0x5b, 0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a, 0x5d, 0x7b, 0x31, 0x2c, 0x31, 0x32, 0x38,
	0x7d, 0x38, 0x05, 0x42, 0x3d, 0x08, 0x02, 0x12, 0x06, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x0a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x08, 0x0a, 0x06, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75,
	0x72, 0x65, 0x64, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x07, 0x0a, 0x05, 0x75, 0x73, 0x61,
	0x67, 0x65, 0x42, 0x58, 0x08, 0x03, 0x12, 0x06, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x09,
	0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x0a, 0x0a, 0x08, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x08, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12,
	0x12, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x64, 0x5f, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x12, 0x0e, 0x0a, 0x0c, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x12, 0x07, 0x0a, 0x05, 0x75, 0x73, 0x61, 0x67, 0x65, 0xb2, 0xdf, 0x21, 0x36,
	0x0a, 0x34, 0x0a, 0x32, 0x0a, 0x06, 0x62, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x1a, 0x1b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x2f, 0x2d, 0x2a,
	0x05, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0xda, 0x94, 0x23, 0x08, 0x12, 0x06, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0xe2, 0xde, 0x21, 0x02, 0x08, 0x02, 0xc2, 0x85, 0x2c, 0x5b, 0x22, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x08, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x22, 0x10, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x64, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22,
	0x07, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2a, 0x0c, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2a, 0x05, 0x75, 0x73, 0x61, 0x67, 0x65, 0x42, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0xee, 0x01, 0xe8, 0xde, 0x21, 0x01, 0xd2, 0xff,
	0xd0, 0x02, 0x3f, 0x0a, 0x0b, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x12, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x73, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0xa2, 0x80, 0xd1, 0x02, 0x41, 0x0a, 0x0c, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c,
	0x71, 0x2f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x74,
	0x74, 0x2e, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x42, 0x0a,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3a, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61,
	0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x3b, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	edgelq_limits_proto_v1_limit_proto_rawDescOnce sync.Once
	edgelq_limits_proto_v1_limit_proto_rawDescData = edgelq_limits_proto_v1_limit_proto_rawDesc
)

func edgelq_limits_proto_v1_limit_proto_rawDescGZIP() []byte {
	edgelq_limits_proto_v1_limit_proto_rawDescOnce.Do(func() {
		edgelq_limits_proto_v1_limit_proto_rawDescData = protoimpl.X.CompressGZIP(edgelq_limits_proto_v1_limit_proto_rawDescData)
	})
	return edgelq_limits_proto_v1_limit_proto_rawDescData
}

var edgelq_limits_proto_v1_limit_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var edgelq_limits_proto_v1_limit_proto_goTypes = []interface{}{
	(*Limit)(nil),     // 0: ntt.limits.v1.Limit
	(*meta.Meta)(nil), // 1: goten.types.Meta
}
var edgelq_limits_proto_v1_limit_proto_depIdxs = []int32{
	1, // 0: ntt.limits.v1.Limit.metadata:type_name -> goten.types.Meta
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { edgelq_limits_proto_v1_limit_proto_init() }
func edgelq_limits_proto_v1_limit_proto_init() {
	if edgelq_limits_proto_v1_limit_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		edgelq_limits_proto_v1_limit_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Limit); i {
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
			RawDescriptor: edgelq_limits_proto_v1_limit_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           edgelq_limits_proto_v1_limit_proto_goTypes,
		DependencyIndexes: edgelq_limits_proto_v1_limit_proto_depIdxs,
		MessageInfos:      edgelq_limits_proto_v1_limit_proto_msgTypes,
	}.Build()
	edgelq_limits_proto_v1_limit_proto = out.File
	edgelq_limits_proto_v1_limit_proto_rawDesc = nil
	edgelq_limits_proto_v1_limit_proto_goTypes = nil
	edgelq_limits_proto_v1_limit_proto_depIdxs = nil
}
