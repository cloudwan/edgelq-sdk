// Code generated by protoc-gen-goten-go
// File: edgelq/limits/proto/v1alpha2/plan_assignment.proto
// DO NOT EDIT!!!

package plan_assignment

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
	iam_organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/organization"
	iam_project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/project"
	accepted_plan "github.com/cloudwan/edgelq-sdk/limits/resources/v1alpha2/accepted_plan"
	common "github.com/cloudwan/edgelq-sdk/limits/resources/v1alpha2/common"
	plan "github.com/cloudwan/edgelq-sdk/limits/resources/v1alpha2/plan"
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
	_ = &iam_organization.Organization{}
	_ = &iam_project.Project{}
	_ = &accepted_plan.AcceptedPlan{}
	_ = &common.Allowance{}
	_ = &plan.Plan{}
	_ = &meta_service.Service{}
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// PlanAssignment binds plan (with optional customizations) with specific
// entity (assignee - either system, organization or project).
// Plan assignments (sum of them) are used to compute configured size of
// limit pools (system and organizations) AND limits (projects).
// They are not modifiable - limits are passed from system to projects,
// so each PlanAssignment must be agreed by higher instance (parent organization
// or system). System recognizes acceptance by existence of AcceptedPlan
// instances - it synchronizes PlanAssignment instances with them.
type PlanAssignment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Name of PlanAssignment
	Name *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// Plan with basic set of limits
	Plan *plan.Reference `protobuf:"bytes,2,opt,customtype=Reference,name=plan,proto3" json:"plan,omitempty" firestore:"plan"`
	// Service associated with plan
	Service *meta_service.Reference `protobuf:"bytes,3,opt,customtype=Reference,name=service,proto3" json:"service,omitempty" firestore:"service"`
	// Additional extensions (but may be negative) on top of regular plan.
	Extensions []*common.Allowance `protobuf:"bytes,4,rep,name=extensions,proto3" json:"extensions,omitempty" firestore:"extensions"`
	// Regional distributions across regions. This field makes sense only
	// for projects, because limit pools are regionless.
	RegionalDistributions []*common.RegionalDistribution `protobuf:"bytes,5,rep,name=regional_distributions,json=regionalDistributions,proto3" json:"regional_distributions,omitempty" firestore:"regionalDistributions"`
	// Source with which this PlanAssignment is synchronized.
	Source *accepted_plan.Reference `protobuf:"bytes,6,opt,customtype=Reference,name=source,proto3" json:"source,omitempty" firestore:"source"`
	// Metadata
	Metadata *ntt_meta.Meta `protobuf:"bytes,7,opt,name=metadata,proto3" json:"metadata,omitempty" firestore:"metadata"`
}

func (m *PlanAssignment) Reset() {
	*m = PlanAssignment{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_limits_proto_v1alpha2_plan_assignment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *PlanAssignment) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*PlanAssignment) ProtoMessage() {}

func (m *PlanAssignment) ProtoReflect() preflect.Message {
	mi := &edgelq_limits_proto_v1alpha2_plan_assignment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*PlanAssignment) GotenMessage() {}

// Deprecated, Use PlanAssignment.ProtoReflect.Descriptor instead.
func (*PlanAssignment) Descriptor() ([]byte, []int) {
	return edgelq_limits_proto_v1alpha2_plan_assignment_proto_rawDescGZIP(), []int{0}
}

func (m *PlanAssignment) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *PlanAssignment) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *PlanAssignment) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *PlanAssignment) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *PlanAssignment) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *PlanAssignment) GetPlan() *plan.Reference {
	if m != nil {
		return m.Plan
	}
	return nil
}

func (m *PlanAssignment) GetService() *meta_service.Reference {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *PlanAssignment) GetExtensions() []*common.Allowance {
	if m != nil {
		return m.Extensions
	}
	return nil
}

func (m *PlanAssignment) GetRegionalDistributions() []*common.RegionalDistribution {
	if m != nil {
		return m.RegionalDistributions
	}
	return nil
}

func (m *PlanAssignment) GetSource() *accepted_plan.Reference {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *PlanAssignment) GetMetadata() *ntt_meta.Meta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *PlanAssignment) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "PlanAssignment"))
	}
	m.Name = fv
}

func (m *PlanAssignment) SetPlan(fv *plan.Reference) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Plan", "PlanAssignment"))
	}
	m.Plan = fv
}

func (m *PlanAssignment) SetService(fv *meta_service.Reference) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Service", "PlanAssignment"))
	}
	m.Service = fv
}

func (m *PlanAssignment) SetExtensions(fv []*common.Allowance) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Extensions", "PlanAssignment"))
	}
	m.Extensions = fv
}

func (m *PlanAssignment) SetRegionalDistributions(fv []*common.RegionalDistribution) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "RegionalDistributions", "PlanAssignment"))
	}
	m.RegionalDistributions = fv
}

func (m *PlanAssignment) SetSource(fv *accepted_plan.Reference) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Source", "PlanAssignment"))
	}
	m.Source = fv
}

func (m *PlanAssignment) SetMetadata(fv *ntt_meta.Meta) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Metadata", "PlanAssignment"))
	}
	m.Metadata = fv
}

var edgelq_limits_proto_v1alpha2_plan_assignment_proto preflect.FileDescriptor

var edgelq_limits_proto_v1alpha2_plan_assignment_proto_rawDesc = []byte{
	0x0a, 0x32, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x70,
	0x6c, 0x61, 0x6e, 0x5f, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x6e, 0x74, 0x74, 0x2e, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69,
	0x61, 0x6d, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x69,
	0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71,
	0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x34,
	0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x70,
	0x69, 0x6c, 0x65, 0x72, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x32, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x2c, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x6f, 0x72, 0x67, 0x61, 0x6e,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x65,
	0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6d,
	0x65, 0x74, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x32, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x27, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x70,
	0x6c, 0x61, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xab, 0x07, 0x0a, 0x0e, 0x50, 0x6c,
	0x61, 0x6e, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x16, 0xb2, 0xda, 0x21, 0x12,
	0x0a, 0x10, 0x0a, 0x0e, 0x50, 0x6c, 0x61, 0x6e, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x70, 0x6c, 0x61, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x10, 0xf0, 0xd9, 0x21, 0x01, 0xb2, 0xda, 0x21, 0x08,
	0x12, 0x06, 0x0a, 0x04, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x04, 0x70, 0x6c, 0x61, 0x6e, 0x12, 0x3f,
	0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x25, 0xf0, 0xd9, 0x21, 0x01, 0xb2, 0xda, 0x21, 0x1d, 0x12, 0x1b, 0x0a, 0x19, 0x6e, 0x74, 0x74,
	0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x44, 0x0a, 0x0a, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x61,
	0x6e, 0x63, 0x65, 0x42, 0x04, 0xf0, 0xd9, 0x21, 0x01, 0x52, 0x0a, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x66, 0x0a, 0x16, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x61,
	0x6c, 0x5f, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x52, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x61, 0x6c, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x04, 0xf0, 0xd9, 0x21, 0x01, 0x52, 0x15, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x61, 0x6c,
	0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2e, 0x0a,
	0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x16, 0xb2,
	0xda, 0x21, 0x12, 0x12, 0x10, 0x0a, 0x0c, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x50,
	0x6c, 0x61, 0x6e, 0x10, 0x03, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x2b, 0x0a,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x61,
	0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x3a, 0xfa, 0x03, 0xea, 0x41, 0xbb,
	0x01, 0x0a, 0x20, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x6c, 0x61, 0x6e, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x21, 0x70, 0x6c, 0x61, 0x6e, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x61, 0x73, 0x73, 0x69, 0x67,
	0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x7d, 0x12, 0x34, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x41,
	0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x6c, 0x61, 0x6e,
	0x5f, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x7d, 0x12, 0x3e, 0x6f, 0x72,
	0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6f, 0x72, 0x67,
	0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x41,
	0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x6c, 0x61, 0x6e,
	0x5f, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x7d, 0x92, 0xd9, 0x21, 0xdb,
	0x01, 0x0a, 0x0f, 0x70, 0x6c, 0x61, 0x6e, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x12, 0x0f, 0x70, 0x6c, 0x61, 0x6e, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x1a, 0x12, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x4e, 0x6f, 0x6e, 0x65, 0x1a, 0x18, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x1a, 0x1d, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x32, 0x2f, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x4a, 0x1d, 0x08, 0x02, 0x12, 0x06, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x06, 0x0a, 0x04, 0x70, 0x6c, 0x61, 0x6e, 0x12, 0x09, 0x0a, 0x07, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x4a, 0x45, 0x08, 0x03, 0x12, 0x06, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x06, 0x0a, 0x04, 0x70, 0x6c, 0x61, 0x6e, 0x12, 0x09, 0x0a, 0x07, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x0c, 0x0a, 0x0a, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x18, 0x0a, 0x16, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x64,
	0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0xda, 0x94, 0x23, 0x08,
	0x12, 0x06, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0xc2, 0x85, 0x2c, 0x4b, 0x22, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x04, 0x70, 0x6c, 0x61, 0x6e, 0x22, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x22, 0x0a, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x16, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0xff, 0x03, 0xe8, 0xde, 0x21, 0x01, 0xd2, 0xff,
	0xd0, 0x02, 0x59, 0x0a, 0x15, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x40, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f,
	0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x2f, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x70, 0x6c, 0x61,
	0x6e, 0x5f, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x0a, 0x1a, 0x63, 0x6f,
	0x6d, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x2e, 0x70, 0x62, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x42, 0x13, 0x50, 0x6c, 0x61, 0x6e, 0x41, 0x73,
	0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x54, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x61, 0x73, 0x73, 0x69, 0x67,
	0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x3b, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x61, 0x73, 0x73, 0x69, 0x67,
	0x6e, 0x6d, 0x65, 0x6e, 0x74, 0xd2, 0x84, 0xd1, 0x02, 0x49, 0x0a, 0x0d, 0x61, 0x75, 0x64, 0x69,
	0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x12, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65,
	0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x2f, 0x61, 0x75, 0x64,
	0x69, 0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x32, 0xf2, 0x85, 0xd1, 0x02, 0x61, 0x0a, 0x19, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x61,
	0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x62, 0x5f, 0x73, 0x79, 0x6e,
	0x63, 0x65, 0x72, 0x12, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x2f, 0x64, 0x62, 0x5f, 0x73, 0x79, 0x6e, 0x63, 0x65, 0x72,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x61,
	0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0xa2, 0x80, 0xd1, 0x02, 0x5b, 0x0a, 0x16,
	0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x5f,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65,
	0x6c, 0x71, 0x2f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x61,
	0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	edgelq_limits_proto_v1alpha2_plan_assignment_proto_rawDescOnce sync.Once
	edgelq_limits_proto_v1alpha2_plan_assignment_proto_rawDescData = edgelq_limits_proto_v1alpha2_plan_assignment_proto_rawDesc
)

func edgelq_limits_proto_v1alpha2_plan_assignment_proto_rawDescGZIP() []byte {
	edgelq_limits_proto_v1alpha2_plan_assignment_proto_rawDescOnce.Do(func() {
		edgelq_limits_proto_v1alpha2_plan_assignment_proto_rawDescData = protoimpl.X.CompressGZIP(edgelq_limits_proto_v1alpha2_plan_assignment_proto_rawDescData)
	})
	return edgelq_limits_proto_v1alpha2_plan_assignment_proto_rawDescData
}

var edgelq_limits_proto_v1alpha2_plan_assignment_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var edgelq_limits_proto_v1alpha2_plan_assignment_proto_goTypes = []interface{}{
	(*PlanAssignment)(nil),              // 0: ntt.limits.v1alpha2.PlanAssignment
	(*common.Allowance)(nil),            // 1: ntt.limits.v1alpha2.Allowance
	(*common.RegionalDistribution)(nil), // 2: ntt.limits.v1alpha2.RegionalDistribution
	(*ntt_meta.Meta)(nil),               // 3: ntt.types.Meta
}
var edgelq_limits_proto_v1alpha2_plan_assignment_proto_depIdxs = []int32{
	1, // 0: ntt.limits.v1alpha2.PlanAssignment.extensions:type_name -> ntt.limits.v1alpha2.Allowance
	2, // 1: ntt.limits.v1alpha2.PlanAssignment.regional_distributions:type_name -> ntt.limits.v1alpha2.RegionalDistribution
	3, // 2: ntt.limits.v1alpha2.PlanAssignment.metadata:type_name -> ntt.types.Meta
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { edgelq_limits_proto_v1alpha2_plan_assignment_proto_init() }
func edgelq_limits_proto_v1alpha2_plan_assignment_proto_init() {
	if edgelq_limits_proto_v1alpha2_plan_assignment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		edgelq_limits_proto_v1alpha2_plan_assignment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlanAssignment); i {
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
			RawDescriptor: edgelq_limits_proto_v1alpha2_plan_assignment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           edgelq_limits_proto_v1alpha2_plan_assignment_proto_goTypes,
		DependencyIndexes: edgelq_limits_proto_v1alpha2_plan_assignment_proto_depIdxs,
		MessageInfos:      edgelq_limits_proto_v1alpha2_plan_assignment_proto_msgTypes,
	}.Build()
	edgelq_limits_proto_v1alpha2_plan_assignment_proto = out.File
	edgelq_limits_proto_v1alpha2_plan_assignment_proto_rawDesc = nil
	edgelq_limits_proto_v1alpha2_plan_assignment_proto_goTypes = nil
	edgelq_limits_proto_v1alpha2_plan_assignment_proto_depIdxs = nil
}