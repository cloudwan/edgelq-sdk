// Code generated by protoc-gen-goten-go
// File: edgelq/iam/proto/v1/common_client.proto
// DO NOT EDIT!!!

package iam_common

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
	group "github.com/cloudwan/edgelq-sdk/iam/resources/v1/group"
	service_account "github.com/cloudwan/edgelq-sdk/iam/resources/v1/service_account"
	user "github.com/cloudwan/edgelq-sdk/iam/resources/v1/user"
	view "github.com/cloudwan/goten-sdk/types/view"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
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
	_ = &service_account.ServiceAccount{}
	_ = &user.User{}
	_ = &fieldmaskpb.FieldMask{}
	_ = view.View(0)
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// MembersMasks determines what fields should be obtained for matching
// members to GroupMembers.
type MembersMasks struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// View defines list of standard response fields present in member items.
	// Additional fields can be amended by specific field masks.
	View view.View `protobuf:"varint,1,opt,name=view,proto3,enum=goten.types.View" json:"view,omitempty" firestore:"view"`
	// A list of extra fields to be obtained for each member User on top of
	// fields defined by request field view
	UserMask *user.User_FieldMask `protobuf:"bytes,2,opt,customtype=User_FieldMask,name=user_mask,json=userMask,proto3" json:"user_mask,omitempty" firestore:"userMask"`
	// A list of extra fields to be obtained for each member ServiceAccount on
	// top of fields defined by request field view
	ServiceAccountMask *service_account.ServiceAccount_FieldMask `protobuf:"bytes,3,opt,customtype=ServiceAccount_FieldMask,name=service_account_mask,json=serviceAccountMask,proto3" json:"service_account_mask,omitempty" firestore:"serviceAccountMask"`
	// A list of extra fields to be obtained for each member Group on top of
	// fields defined by request field view
	GroupMask *group.Group_FieldMask `protobuf:"bytes,4,opt,customtype=Group_FieldMask,name=group_mask,json=groupMask,proto3" json:"group_mask,omitempty" firestore:"groupMask"`
}

func (m *MembersMasks) Reset() {
	*m = MembersMasks{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_iam_proto_v1_common_client_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *MembersMasks) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*MembersMasks) ProtoMessage() {}

func (m *MembersMasks) ProtoReflect() preflect.Message {
	mi := &edgelq_iam_proto_v1_common_client_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*MembersMasks) GotenMessage() {}

// Deprecated, Use MembersMasks.ProtoReflect.Descriptor instead.
func (*MembersMasks) Descriptor() ([]byte, []int) {
	return edgelq_iam_proto_v1_common_client_proto_rawDescGZIP(), []int{0}
}

func (m *MembersMasks) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *MembersMasks) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *MembersMasks) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *MembersMasks) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *MembersMasks) GetView() view.View {
	if m != nil {
		return m.View
	}
	return view.View_UNSPECIFIED
}

func (m *MembersMasks) GetUserMask() *user.User_FieldMask {
	if m != nil {
		return m.UserMask
	}
	return nil
}

func (m *MembersMasks) GetServiceAccountMask() *service_account.ServiceAccount_FieldMask {
	if m != nil {
		return m.ServiceAccountMask
	}
	return nil
}

func (m *MembersMasks) GetGroupMask() *group.Group_FieldMask {
	if m != nil {
		return m.GroupMask
	}
	return nil
}

func (m *MembersMasks) SetView(fv view.View) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "View", "MembersMasks"))
	}
	m.View = fv
}

func (m *MembersMasks) SetUserMask(fv *user.User_FieldMask) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "UserMask", "MembersMasks"))
	}
	m.UserMask = fv
}

func (m *MembersMasks) SetServiceAccountMask(fv *service_account.ServiceAccount_FieldMask) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ServiceAccountMask", "MembersMasks"))
	}
	m.ServiceAccountMask = fv
}

func (m *MembersMasks) SetGroupMask(fv *group.Group_FieldMask) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "GroupMask", "MembersMasks"))
	}
	m.GroupMask = fv
}

// MembersInfo maps Users, ServiceAccounts or Groups to specific
// member data.
type MembersInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Member identifier - as in format used in RoleBindings/GroupMembers.
	Member string `protobuf:"bytes,1,opt,name=member,proto3" json:"member,omitempty" firestore:"member"`
	// List of users. May contain more than 1 item, if there are
	// multiple users sharing same email.
	// Array may be empty if member is for service account or
	// group type, or if no user exists with specified email.
	Users []*user.User `protobuf:"bytes,2,rep,name=users,proto3" json:"users,omitempty" firestore:"users"`
	// List of service accounts. It will be empty if member
	// matches user or group type, or if no matching ServiceAccount
	// was found. Length of this array is from 0 to 1.
	ServiceAccounts []*service_account.ServiceAccount `protobuf:"bytes,3,rep,name=service_accounts,json=serviceAccounts,proto3" json:"service_accounts,omitempty" firestore:"serviceAccounts"`
	// List of groups. It will be empty if member matches
	// user or service account type, or if no matching Group
	// was found. Length of this array is from 0 to 1.
	Groups []*group.Group `protobuf:"bytes,4,rep,name=groups,proto3" json:"groups,omitempty" firestore:"groups"`
}

func (m *MembersInfo) Reset() {
	*m = MembersInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_iam_proto_v1_common_client_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *MembersInfo) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*MembersInfo) ProtoMessage() {}

func (m *MembersInfo) ProtoReflect() preflect.Message {
	mi := &edgelq_iam_proto_v1_common_client_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*MembersInfo) GotenMessage() {}

// Deprecated, Use MembersInfo.ProtoReflect.Descriptor instead.
func (*MembersInfo) Descriptor() ([]byte, []int) {
	return edgelq_iam_proto_v1_common_client_proto_rawDescGZIP(), []int{1}
}

func (m *MembersInfo) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *MembersInfo) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *MembersInfo) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *MembersInfo) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *MembersInfo) GetMember() string {
	if m != nil {
		return m.Member
	}
	return ""
}

func (m *MembersInfo) GetUsers() []*user.User {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *MembersInfo) GetServiceAccounts() []*service_account.ServiceAccount {
	if m != nil {
		return m.ServiceAccounts
	}
	return nil
}

func (m *MembersInfo) GetGroups() []*group.Group {
	if m != nil {
		return m.Groups
	}
	return nil
}

func (m *MembersInfo) SetMember(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Member", "MembersInfo"))
	}
	m.Member = fv
}

func (m *MembersInfo) SetUsers(fv []*user.User) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Users", "MembersInfo"))
	}
	m.Users = fv
}

func (m *MembersInfo) SetServiceAccounts(fv []*service_account.ServiceAccount) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ServiceAccounts", "MembersInfo"))
	}
	m.ServiceAccounts = fv
}

func (m *MembersInfo) SetGroups(fv []*group.Group) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Groups", "MembersInfo"))
	}
	m.Groups = fv
}

var edgelq_iam_proto_v1_common_client_proto preflect.FileDescriptor

var edgelq_iam_proto_v1_common_client_proto_rawDesc = []byte{
	0x0a, 0x27, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6e, 0x74, 0x74, 0x2e, 0x69,
	0x61, 0x6d, 0x2e, 0x76, 0x31, 0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x74, 0x65, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61,
	0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x29, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2d, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x65, 0x64, 0x67,
	0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x65, 0x64, 0x67,
	0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69,
	0x61, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xac, 0x02, 0x0a, 0x0c, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x73, 0x4d, 0x61, 0x73, 0x6b, 0x73, 0x12, 0x25, 0x0a, 0x04, 0x76, 0x69, 0x65, 0x77,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x52, 0x04, 0x76, 0x69, 0x65, 0x77, 0x12,
	0x45, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x42, 0x0c,
	0xb2, 0xda, 0x21, 0x08, 0x32, 0x06, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x64, 0x0a, 0x14, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b,
	0x42, 0x16, 0xb2, 0xda, 0x21, 0x12, 0x32, 0x10, 0x0a, 0x0e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x12, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x48, 0x0a, 0x0a,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x42, 0x0d, 0xb2, 0xda,
	0x21, 0x09, 0x32, 0x07, 0x0a, 0x05, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x09, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x4d, 0x61, 0x73, 0x6b, 0x22, 0xbf, 0x01, 0x0a, 0x0b, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x26,
	0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0x45, 0x0a, 0x10, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x0f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x29, 0x0a,
	0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x52, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x42, 0x55, 0xe8, 0xde, 0x21, 0x00, 0x0a, 0x11,
	0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x70, 0x62, 0x2e, 0x76,
	0x31, 0x50, 0x01, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f,
	0x69, 0x61, 0x6d, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x3b, 0x69, 0x61, 0x6d, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	edgelq_iam_proto_v1_common_client_proto_rawDescOnce sync.Once
	edgelq_iam_proto_v1_common_client_proto_rawDescData = edgelq_iam_proto_v1_common_client_proto_rawDesc
)

func edgelq_iam_proto_v1_common_client_proto_rawDescGZIP() []byte {
	edgelq_iam_proto_v1_common_client_proto_rawDescOnce.Do(func() {
		edgelq_iam_proto_v1_common_client_proto_rawDescData = protoimpl.X.CompressGZIP(edgelq_iam_proto_v1_common_client_proto_rawDescData)
	})
	return edgelq_iam_proto_v1_common_client_proto_rawDescData
}

var edgelq_iam_proto_v1_common_client_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var edgelq_iam_proto_v1_common_client_proto_goTypes = []interface{}{
	(*MembersMasks)(nil),                             // 0: ntt.iam.v1.MembersMasks
	(*MembersInfo)(nil),                              // 1: ntt.iam.v1.MembersInfo
	(view.View)(0),                                   // 2: goten.types.View
	(*user.User_FieldMask)(nil),                      // 3: ntt.iam.v1.User_FieldMask
	(*service_account.ServiceAccount_FieldMask)(nil), // 4: ntt.iam.v1.ServiceAccount_FieldMask
	(*group.Group_FieldMask)(nil),                    // 5: ntt.iam.v1.Group_FieldMask
	(*user.User)(nil),                                // 6: ntt.iam.v1.User
	(*service_account.ServiceAccount)(nil),           // 7: ntt.iam.v1.ServiceAccount
	(*group.Group)(nil),                              // 8: ntt.iam.v1.Group
}
var edgelq_iam_proto_v1_common_client_proto_depIdxs = []int32{
	2, // 0: ntt.iam.v1.MembersMasks.view:type_name -> goten.types.View
	3, // 1: ntt.iam.v1.MembersMasks.user_mask:type_name -> ntt.iam.v1.User_FieldMask
	4, // 2: ntt.iam.v1.MembersMasks.service_account_mask:type_name -> ntt.iam.v1.ServiceAccount_FieldMask
	5, // 3: ntt.iam.v1.MembersMasks.group_mask:type_name -> ntt.iam.v1.Group_FieldMask
	6, // 4: ntt.iam.v1.MembersInfo.users:type_name -> ntt.iam.v1.User
	7, // 5: ntt.iam.v1.MembersInfo.service_accounts:type_name -> ntt.iam.v1.ServiceAccount
	8, // 6: ntt.iam.v1.MembersInfo.groups:type_name -> ntt.iam.v1.Group
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { edgelq_iam_proto_v1_common_client_proto_init() }
func edgelq_iam_proto_v1_common_client_proto_init() {
	if edgelq_iam_proto_v1_common_client_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		edgelq_iam_proto_v1_common_client_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MembersMasks); i {
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
		edgelq_iam_proto_v1_common_client_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MembersInfo); i {
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
			RawDescriptor: edgelq_iam_proto_v1_common_client_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           edgelq_iam_proto_v1_common_client_proto_goTypes,
		DependencyIndexes: edgelq_iam_proto_v1_common_client_proto_depIdxs,
		MessageInfos:      edgelq_iam_proto_v1_common_client_proto_msgTypes,
	}.Build()
	edgelq_iam_proto_v1_common_client_proto = out.File
	edgelq_iam_proto_v1_common_client_proto_rawDesc = nil
	edgelq_iam_proto_v1_common_client_proto_goTypes = nil
	edgelq_iam_proto_v1_common_client_proto_depIdxs = nil
}
