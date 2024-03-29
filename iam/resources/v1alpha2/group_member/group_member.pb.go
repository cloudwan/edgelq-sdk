// Code generated by protoc-gen-goten-go
// File: edgelq/iam/proto/v1alpha2/group_member.proto
// DO NOT EDIT!!!

package group_member

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
	group "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/group"
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
	_ = &group.Group{}
	_ = &meta.Meta{}
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// GroupMember Resource
type GroupMember struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Name of GroupMember
	// When creating a new instance, this field is optional and if not provided,
	// it will be generated automatically. Last ID segment must conform to the
	// following regex: [a-z][a-z0-9\-]{0,28}[a-z0-9]
	Name *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// Format of the string is one of:
	// - "allUsers" (anyone)
	// - "allAuthenticatedUsers" (anyone logged in)
	// - "user:admin.super@example.com"
	// - "serviceAccount:device_agent@watchdog.serviceaccounts.iam.edgelq.com"
	// - "group:nice.group@example.com"
	Member string `protobuf:"bytes,2,opt,name=member,proto3" json:"member,omitempty" firestore:"member"`
	// Internal field used by the IAM controller to know member ID (email) of a
	// group this member belongs to.
	ParentMember string `protobuf:"bytes,5,opt,name=parent_member,json=parentMember,proto3" json:"parent_member,omitempty" firestore:"parentMember"`
	// Internal field used the IAM controller to know common ancestors if this
	// group member is actually a copy from another group, which can happen if
	// group is member in another group. List contains member IDs of other groups.
	// Last element is always pointing to the group which made an actual copy and
	// is responsible for deletion when needed. If particular group member is
	// copied multiple times (which also can happen if group is indirect member
	// via multiple memberships) then this field contains list only of common
	// ancestors. This field has purpose of indicating owner (last item in this
	// list) and to break cycles (previous elements). List is always empty if this
	// group member is DIRECT member of this group and controller cannot modify
	// it.
	MinAncestryMembers []string `protobuf:"bytes,6,rep,name=min_ancestry_members,json=minAncestryMembers,proto3" json:"min_ancestry_members,omitempty" firestore:"minAncestryMembers"`
	// Metadata
	Metadata *meta.Meta `protobuf:"bytes,7,opt,name=metadata,proto3" json:"metadata,omitempty" firestore:"metadata"`
}

func (m *GroupMember) Reset() {
	*m = GroupMember{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_iam_proto_v1alpha2_group_member_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *GroupMember) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*GroupMember) ProtoMessage() {}

func (m *GroupMember) ProtoReflect() preflect.Message {
	mi := &edgelq_iam_proto_v1alpha2_group_member_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*GroupMember) GotenMessage() {}

// Deprecated, Use GroupMember.ProtoReflect.Descriptor instead.
func (*GroupMember) Descriptor() ([]byte, []int) {
	return edgelq_iam_proto_v1alpha2_group_member_proto_rawDescGZIP(), []int{0}
}

func (m *GroupMember) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *GroupMember) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *GroupMember) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *GroupMember) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *GroupMember) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *GroupMember) GetMember() string {
	if m != nil {
		return m.Member
	}
	return ""
}

func (m *GroupMember) GetParentMember() string {
	if m != nil {
		return m.ParentMember
	}
	return ""
}

func (m *GroupMember) GetMinAncestryMembers() []string {
	if m != nil {
		return m.MinAncestryMembers
	}
	return nil
}

func (m *GroupMember) GetMetadata() *meta.Meta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *GroupMember) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "GroupMember"))
	}
	m.Name = fv
}

func (m *GroupMember) SetMember(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Member", "GroupMember"))
	}
	m.Member = fv
}

func (m *GroupMember) SetParentMember(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ParentMember", "GroupMember"))
	}
	m.ParentMember = fv
}

func (m *GroupMember) SetMinAncestryMembers(fv []string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "MinAncestryMembers", "GroupMember"))
	}
	m.MinAncestryMembers = fv
}

func (m *GroupMember) SetMetadata(fv *meta.Meta) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Metadata", "GroupMember"))
	}
	m.Metadata = fv
}

var edgelq_iam_proto_v1alpha2_group_member_proto preflect.FileDescriptor

var edgelq_iam_proto_v1alpha2_group_member_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x5f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10,
	0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32,
	0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x74,
	0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67,
	0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f,
	0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d,
	0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24,
	0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb3, 0x06, 0x0a,
	0x0b, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x13, 0xb2, 0xda, 0x21, 0x0f,
	0x0a, 0x0d, 0x0a, 0x0b, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x29, 0x0a,
	0x0d, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xf0, 0xd9, 0x21, 0x01, 0x52, 0x0c, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x3a, 0x0a, 0x14, 0x6d, 0x69, 0x6e, 0x5f,
	0x61, 0x6e, 0x63, 0x65, 0x73, 0x74, 0x72, 0x79, 0x5f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x42, 0x08, 0xea, 0xde, 0x21, 0x04, 0x0a, 0x02, 0x10, 0x04,
	0x52, 0x12, 0x6d, 0x69, 0x6e, 0x41, 0x6e, 0x63, 0x65, 0x73, 0x74, 0x72, 0x79, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x73, 0x12, 0x2d, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x3a, 0xc0, 0x04, 0xea, 0x41, 0xd0, 0x01, 0x0a, 0x1a, 0x69, 0x61, 0x6d, 0x2e,
	0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2a, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2f, 0x7b,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x7d, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x73, 0x2f, 0x7b, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x7d, 0x12, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2f, 0x7b, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x7d, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x73, 0x2f, 0x7b, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x7d, 0x12, 0x47, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x7b, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2f, 0x7b, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x7d, 0x2f, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x5f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x7d, 0x92, 0xd9, 0x21, 0x5d, 0x0a, 0x0c,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x0c, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x1a, 0x05, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x38, 0x05, 0x42, 0x14, 0x08, 0x02, 0x12, 0x06, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x08, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x20, 0x08, 0x03, 0x12, 0x06, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x08, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x0a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xb2, 0xdf, 0x21, 0xb9, 0x01,
	0x0a, 0xb6, 0x01, 0x0a, 0xa2, 0x01, 0x0a, 0x06, 0x62, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x1a, 0x30, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2f,
	0x7b, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x7d, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x73, 0x2f, 0x2d, 0x1a, 0x3a, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2f, 0x7b, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x7d, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73,
	0x2f, 0x2d, 0x1a, 0x1d, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2f, 0x7b, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x7d, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x2f,
	0x2d, 0x2a, 0x05, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x12, 0x0f, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x2a, 0x05, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0xda, 0x94, 0x23, 0x08, 0x12, 0x06, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0xc2, 0x85, 0x2c, 0x3d, 0x22, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2a, 0x0d, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2a, 0x14, 0x6d, 0x69, 0x6e, 0x5f, 0x61, 0x6e, 0x63, 0x65,
	0x73, 0x74, 0x72, 0x79, 0x5f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x42, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x4a, 0x04, 0x08, 0x03, 0x10, 0x04, 0x4a, 0x04, 0x08, 0x04,
	0x10, 0x05, 0x42, 0xaa, 0x02, 0xe8, 0xde, 0x21, 0x01, 0xd2, 0xff, 0xd0, 0x02, 0x50, 0x0a, 0x12,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x12, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69,
	0x61, 0x6d, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x32, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0xa2, 0x80,
	0xd1, 0x02, 0x52, 0x0a, 0x13, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64,
	0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69,
	0x61, 0x6d, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x42, 0x10,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x4b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69,
	0x61, 0x6d, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x3b, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	edgelq_iam_proto_v1alpha2_group_member_proto_rawDescOnce sync.Once
	edgelq_iam_proto_v1alpha2_group_member_proto_rawDescData = edgelq_iam_proto_v1alpha2_group_member_proto_rawDesc
)

func edgelq_iam_proto_v1alpha2_group_member_proto_rawDescGZIP() []byte {
	edgelq_iam_proto_v1alpha2_group_member_proto_rawDescOnce.Do(func() {
		edgelq_iam_proto_v1alpha2_group_member_proto_rawDescData = protoimpl.X.CompressGZIP(edgelq_iam_proto_v1alpha2_group_member_proto_rawDescData)
	})
	return edgelq_iam_proto_v1alpha2_group_member_proto_rawDescData
}

var edgelq_iam_proto_v1alpha2_group_member_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var edgelq_iam_proto_v1alpha2_group_member_proto_goTypes = []interface{}{
	(*GroupMember)(nil), // 0: ntt.iam.v1alpha2.GroupMember
	(*meta.Meta)(nil),   // 1: goten.types.Meta
}
var edgelq_iam_proto_v1alpha2_group_member_proto_depIdxs = []int32{
	1, // 0: ntt.iam.v1alpha2.GroupMember.metadata:type_name -> goten.types.Meta
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { edgelq_iam_proto_v1alpha2_group_member_proto_init() }
func edgelq_iam_proto_v1alpha2_group_member_proto_init() {
	if edgelq_iam_proto_v1alpha2_group_member_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		edgelq_iam_proto_v1alpha2_group_member_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupMember); i {
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
			RawDescriptor: edgelq_iam_proto_v1alpha2_group_member_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           edgelq_iam_proto_v1alpha2_group_member_proto_goTypes,
		DependencyIndexes: edgelq_iam_proto_v1alpha2_group_member_proto_depIdxs,
		MessageInfos:      edgelq_iam_proto_v1alpha2_group_member_proto_msgTypes,
	}.Build()
	edgelq_iam_proto_v1alpha2_group_member_proto = out.File
	edgelq_iam_proto_v1alpha2_group_member_proto_rawDesc = nil
	edgelq_iam_proto_v1alpha2_group_member_proto_goTypes = nil
	edgelq_iam_proto_v1alpha2_group_member_proto_depIdxs = nil
}
