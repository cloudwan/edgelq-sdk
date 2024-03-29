// Code generated by protoc-gen-goten-go
// File: edgelq/iam/proto/v1alpha2/project_invitation_custom.proto
// DO NOT EDIT!!!

package project_invitation_client

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
	project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/project"
	project_invitation "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/project_invitation"
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
	_ = &project_invitation.ProjectInvitation{}
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Request message for method
// [AcceptProjectInvitation][ntt.iam.v1alpha2.AcceptProjectInvitation]
type AcceptProjectInvitationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	//  name of ntt.iam.v1alpha2.ProjectInvitation
	Name *project_invitation.Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
}

func (m *AcceptProjectInvitationRequest) Reset() {
	*m = AcceptProjectInvitationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_iam_proto_v1alpha2_project_invitation_custom_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *AcceptProjectInvitationRequest) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*AcceptProjectInvitationRequest) ProtoMessage() {}

func (m *AcceptProjectInvitationRequest) ProtoReflect() preflect.Message {
	mi := &edgelq_iam_proto_v1alpha2_project_invitation_custom_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*AcceptProjectInvitationRequest) GotenMessage() {}

// Deprecated, Use AcceptProjectInvitationRequest.ProtoReflect.Descriptor instead.
func (*AcceptProjectInvitationRequest) Descriptor() ([]byte, []int) {
	return edgelq_iam_proto_v1alpha2_project_invitation_custom_proto_rawDescGZIP(), []int{0}
}

func (m *AcceptProjectInvitationRequest) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *AcceptProjectInvitationRequest) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *AcceptProjectInvitationRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *AcceptProjectInvitationRequest) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *AcceptProjectInvitationRequest) GetName() *project_invitation.Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *AcceptProjectInvitationRequest) SetName(fv *project_invitation.Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "AcceptProjectInvitationRequest"))
	}
	m.Name = fv
}

// Response message for method
// [AcceptProjectInvitation][ntt.iam.v1alpha2.AcceptProjectInvitation]
type AcceptProjectInvitationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (m *AcceptProjectInvitationResponse) Reset() {
	*m = AcceptProjectInvitationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_iam_proto_v1alpha2_project_invitation_custom_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *AcceptProjectInvitationResponse) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*AcceptProjectInvitationResponse) ProtoMessage() {}

func (m *AcceptProjectInvitationResponse) ProtoReflect() preflect.Message {
	mi := &edgelq_iam_proto_v1alpha2_project_invitation_custom_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*AcceptProjectInvitationResponse) GotenMessage() {}

// Deprecated, Use AcceptProjectInvitationResponse.ProtoReflect.Descriptor instead.
func (*AcceptProjectInvitationResponse) Descriptor() ([]byte, []int) {
	return edgelq_iam_proto_v1alpha2_project_invitation_custom_proto_rawDescGZIP(), []int{1}
}

func (m *AcceptProjectInvitationResponse) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *AcceptProjectInvitationResponse) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *AcceptProjectInvitationResponse) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *AcceptProjectInvitationResponse) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

// Request message for method
// [DeclineProjectInvitation][ntt.iam.v1alpha2.DeclineProjectInvitation]
type DeclineProjectInvitationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	//  name of ntt.iam.v1alpha2.ProjectInvitation
	Name *project_invitation.Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
}

func (m *DeclineProjectInvitationRequest) Reset() {
	*m = DeclineProjectInvitationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_iam_proto_v1alpha2_project_invitation_custom_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *DeclineProjectInvitationRequest) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*DeclineProjectInvitationRequest) ProtoMessage() {}

func (m *DeclineProjectInvitationRequest) ProtoReflect() preflect.Message {
	mi := &edgelq_iam_proto_v1alpha2_project_invitation_custom_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*DeclineProjectInvitationRequest) GotenMessage() {}

// Deprecated, Use DeclineProjectInvitationRequest.ProtoReflect.Descriptor instead.
func (*DeclineProjectInvitationRequest) Descriptor() ([]byte, []int) {
	return edgelq_iam_proto_v1alpha2_project_invitation_custom_proto_rawDescGZIP(), []int{2}
}

func (m *DeclineProjectInvitationRequest) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *DeclineProjectInvitationRequest) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *DeclineProjectInvitationRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *DeclineProjectInvitationRequest) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *DeclineProjectInvitationRequest) GetName() *project_invitation.Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *DeclineProjectInvitationRequest) SetName(fv *project_invitation.Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "DeclineProjectInvitationRequest"))
	}
	m.Name = fv
}

// Response message for method
// [DeclineProjectInvitation][ntt.iam.v1alpha2.DeclineProjectInvitation]
type DeclineProjectInvitationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (m *DeclineProjectInvitationResponse) Reset() {
	*m = DeclineProjectInvitationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_iam_proto_v1alpha2_project_invitation_custom_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *DeclineProjectInvitationResponse) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*DeclineProjectInvitationResponse) ProtoMessage() {}

func (m *DeclineProjectInvitationResponse) ProtoReflect() preflect.Message {
	mi := &edgelq_iam_proto_v1alpha2_project_invitation_custom_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*DeclineProjectInvitationResponse) GotenMessage() {}

// Deprecated, Use DeclineProjectInvitationResponse.ProtoReflect.Descriptor instead.
func (*DeclineProjectInvitationResponse) Descriptor() ([]byte, []int) {
	return edgelq_iam_proto_v1alpha2_project_invitation_custom_proto_rawDescGZIP(), []int{3}
}

func (m *DeclineProjectInvitationResponse) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *DeclineProjectInvitationResponse) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *DeclineProjectInvitationResponse) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *DeclineProjectInvitationResponse) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

// Request message for method
// [ListMyProjectInvitations][ntt.iam.v1alpha2.ListMyProjectInvitations]
type ListMyProjectInvitationsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Parent name of ntt.iam.v1alpha2.ProjectInvitation
	Parent *project_invitation.ParentName `protobuf:"bytes,1,opt,customtype=ParentName,name=parent,proto3" json:"parent,omitempty" firestore:"parent"`
	// Additional filter for invitations, e.g. state = "PENDING"
	Filter *project_invitation.Filter `protobuf:"bytes,5,opt,customtype=Filter,name=filter,proto3" json:"filter,omitempty" firestore:"filter"`
}

func (m *ListMyProjectInvitationsRequest) Reset() {
	*m = ListMyProjectInvitationsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_iam_proto_v1alpha2_project_invitation_custom_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *ListMyProjectInvitationsRequest) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*ListMyProjectInvitationsRequest) ProtoMessage() {}

func (m *ListMyProjectInvitationsRequest) ProtoReflect() preflect.Message {
	mi := &edgelq_iam_proto_v1alpha2_project_invitation_custom_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*ListMyProjectInvitationsRequest) GotenMessage() {}

// Deprecated, Use ListMyProjectInvitationsRequest.ProtoReflect.Descriptor instead.
func (*ListMyProjectInvitationsRequest) Descriptor() ([]byte, []int) {
	return edgelq_iam_proto_v1alpha2_project_invitation_custom_proto_rawDescGZIP(), []int{4}
}

func (m *ListMyProjectInvitationsRequest) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *ListMyProjectInvitationsRequest) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *ListMyProjectInvitationsRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *ListMyProjectInvitationsRequest) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *ListMyProjectInvitationsRequest) GetParent() *project_invitation.ParentName {
	if m != nil {
		return m.Parent
	}
	return nil
}

func (m *ListMyProjectInvitationsRequest) GetFilter() *project_invitation.Filter {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *ListMyProjectInvitationsRequest) SetParent(fv *project_invitation.ParentName) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Parent", "ListMyProjectInvitationsRequest"))
	}
	m.Parent = fv
}

func (m *ListMyProjectInvitationsRequest) SetFilter(fv *project_invitation.Filter) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Filter", "ListMyProjectInvitationsRequest"))
	}
	m.Filter = fv
}

// Response message for method
// [ListMyProjectInvitations][ntt.iam.v1alpha2.ListMyProjectInvitations]
type ListMyProjectInvitationsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// The list of ProjectInvitations
	ProjectInvitations []*project_invitation.ProjectInvitation `protobuf:"bytes,1,rep,name=project_invitations,json=projectInvitations,proto3" json:"project_invitations,omitempty" firestore:"projectInvitations"`
}

func (m *ListMyProjectInvitationsResponse) Reset() {
	*m = ListMyProjectInvitationsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_iam_proto_v1alpha2_project_invitation_custom_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *ListMyProjectInvitationsResponse) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*ListMyProjectInvitationsResponse) ProtoMessage() {}

func (m *ListMyProjectInvitationsResponse) ProtoReflect() preflect.Message {
	mi := &edgelq_iam_proto_v1alpha2_project_invitation_custom_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*ListMyProjectInvitationsResponse) GotenMessage() {}

// Deprecated, Use ListMyProjectInvitationsResponse.ProtoReflect.Descriptor instead.
func (*ListMyProjectInvitationsResponse) Descriptor() ([]byte, []int) {
	return edgelq_iam_proto_v1alpha2_project_invitation_custom_proto_rawDescGZIP(), []int{5}
}

func (m *ListMyProjectInvitationsResponse) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *ListMyProjectInvitationsResponse) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *ListMyProjectInvitationsResponse) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *ListMyProjectInvitationsResponse) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *ListMyProjectInvitationsResponse) GetProjectInvitations() []*project_invitation.ProjectInvitation {
	if m != nil {
		return m.ProjectInvitations
	}
	return nil
}

func (m *ListMyProjectInvitationsResponse) SetProjectInvitations(fv []*project_invitation.ProjectInvitation) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ProjectInvitations", "ListMyProjectInvitationsResponse"))
	}
	m.ProjectInvitations = fv
}

// ResendInvitationRequest
type ResendProjectInvitationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	//  name of ntt.iam.v1alpha2.ProjectInvitation
	Name *project_invitation.Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
}

func (m *ResendProjectInvitationRequest) Reset() {
	*m = ResendProjectInvitationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_iam_proto_v1alpha2_project_invitation_custom_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *ResendProjectInvitationRequest) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*ResendProjectInvitationRequest) ProtoMessage() {}

func (m *ResendProjectInvitationRequest) ProtoReflect() preflect.Message {
	mi := &edgelq_iam_proto_v1alpha2_project_invitation_custom_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*ResendProjectInvitationRequest) GotenMessage() {}

// Deprecated, Use ResendProjectInvitationRequest.ProtoReflect.Descriptor instead.
func (*ResendProjectInvitationRequest) Descriptor() ([]byte, []int) {
	return edgelq_iam_proto_v1alpha2_project_invitation_custom_proto_rawDescGZIP(), []int{6}
}

func (m *ResendProjectInvitationRequest) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *ResendProjectInvitationRequest) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *ResendProjectInvitationRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *ResendProjectInvitationRequest) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *ResendProjectInvitationRequest) GetName() *project_invitation.Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *ResendProjectInvitationRequest) SetName(fv *project_invitation.Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "ResendProjectInvitationRequest"))
	}
	m.Name = fv
}

// ResendInvitationResponse
type ResendProjectInvitationResponse struct {
	state             protoimpl.MessageState
	sizeCache         protoimpl.SizeCache
	unknownFields     protoimpl.UnknownFields
	ProjectInvitation *project_invitation.ProjectInvitation `protobuf:"bytes,1,opt,name=project_invitation,json=projectInvitation,proto3" json:"project_invitation,omitempty" firestore:"projectInvitation"`
}

func (m *ResendProjectInvitationResponse) Reset() {
	*m = ResendProjectInvitationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_iam_proto_v1alpha2_project_invitation_custom_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *ResendProjectInvitationResponse) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*ResendProjectInvitationResponse) ProtoMessage() {}

func (m *ResendProjectInvitationResponse) ProtoReflect() preflect.Message {
	mi := &edgelq_iam_proto_v1alpha2_project_invitation_custom_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*ResendProjectInvitationResponse) GotenMessage() {}

// Deprecated, Use ResendProjectInvitationResponse.ProtoReflect.Descriptor instead.
func (*ResendProjectInvitationResponse) Descriptor() ([]byte, []int) {
	return edgelq_iam_proto_v1alpha2_project_invitation_custom_proto_rawDescGZIP(), []int{7}
}

func (m *ResendProjectInvitationResponse) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *ResendProjectInvitationResponse) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *ResendProjectInvitationResponse) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *ResendProjectInvitationResponse) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *ResendProjectInvitationResponse) GetProjectInvitation() *project_invitation.ProjectInvitation {
	if m != nil {
		return m.ProjectInvitation
	}
	return nil
}

func (m *ResendProjectInvitationResponse) SetProjectInvitation(fv *project_invitation.ProjectInvitation) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ProjectInvitation", "ResendProjectInvitationResponse"))
	}
	m.ProjectInvitation = fv
}

var edgelq_iam_proto_v1alpha2_project_invitation_custom_proto preflect.FileDescriptor

var edgelq_iam_proto_v1alpha2_project_invitation_custom_proto_rawDesc = []byte{
	0x0a, 0x39, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x5f, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x6e, 0x74, 0x74,
	0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x1a, 0x20, 0x67,
	0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20,
	0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x24, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x32, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69,
	0x61, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x32, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x39, 0x65, 0x64, 0x67, 0x65,
	0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x6e,
	0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4f, 0x0a, 0x1e, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x19, 0xb2, 0xda, 0x21, 0x15, 0x0a, 0x13, 0x0a, 0x11, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x21, 0x0a, 0x1f, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x50, 0x0a, 0x1f, 0x44, 0x65, 0x63,
	0x6c, 0x69, 0x6e, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x19, 0xb2, 0xda, 0x21, 0x15,
	0x0a, 0x13, 0x0a, 0x11, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x22, 0x0a, 0x20, 0x44,
	0x65, 0x63, 0x6c, 0x69, 0x6e, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x76,
	0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x87, 0x01, 0x0a, 0x1f, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x79, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x19, 0xb2, 0xda, 0x21, 0x15, 0x3a, 0x13, 0x0a, 0x11, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x31, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x19, 0xb2, 0xda, 0x21, 0x15, 0x1a, 0x13, 0x0a, 0x11,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x78, 0x0a, 0x20, 0x4c, 0x69, 0x73,
	0x74, 0x4d, 0x79, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x54, 0x0a,
	0x13, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6e, 0x74, 0x74,
	0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x12, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x22, 0x4f, 0x0a, 0x1e, 0x52, 0x65, 0x73, 0x65, 0x6e, 0x64, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x19, 0xb2, 0xda, 0x21, 0x15, 0x0a, 0x13, 0x0a, 0x11, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x75, 0x0a, 0x1f, 0x52, 0x65, 0x73, 0x65, 0x6e, 0x64, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x52, 0x0a, 0x12, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x5f, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e,
	0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x11, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x96, 0x01, 0x0a, 0x17,
	0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x70, 0x62, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x42, 0x1c, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x00, 0x5a, 0x5b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67,
	0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f,
	0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3b, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x5f, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	edgelq_iam_proto_v1alpha2_project_invitation_custom_proto_rawDescOnce sync.Once
	edgelq_iam_proto_v1alpha2_project_invitation_custom_proto_rawDescData = edgelq_iam_proto_v1alpha2_project_invitation_custom_proto_rawDesc
)

func edgelq_iam_proto_v1alpha2_project_invitation_custom_proto_rawDescGZIP() []byte {
	edgelq_iam_proto_v1alpha2_project_invitation_custom_proto_rawDescOnce.Do(func() {
		edgelq_iam_proto_v1alpha2_project_invitation_custom_proto_rawDescData = protoimpl.X.CompressGZIP(edgelq_iam_proto_v1alpha2_project_invitation_custom_proto_rawDescData)
	})
	return edgelq_iam_proto_v1alpha2_project_invitation_custom_proto_rawDescData
}

var edgelq_iam_proto_v1alpha2_project_invitation_custom_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var edgelq_iam_proto_v1alpha2_project_invitation_custom_proto_goTypes = []interface{}{
	(*AcceptProjectInvitationRequest)(nil),       // 0: ntt.iam.v1alpha2.AcceptProjectInvitationRequest
	(*AcceptProjectInvitationResponse)(nil),      // 1: ntt.iam.v1alpha2.AcceptProjectInvitationResponse
	(*DeclineProjectInvitationRequest)(nil),      // 2: ntt.iam.v1alpha2.DeclineProjectInvitationRequest
	(*DeclineProjectInvitationResponse)(nil),     // 3: ntt.iam.v1alpha2.DeclineProjectInvitationResponse
	(*ListMyProjectInvitationsRequest)(nil),      // 4: ntt.iam.v1alpha2.ListMyProjectInvitationsRequest
	(*ListMyProjectInvitationsResponse)(nil),     // 5: ntt.iam.v1alpha2.ListMyProjectInvitationsResponse
	(*ResendProjectInvitationRequest)(nil),       // 6: ntt.iam.v1alpha2.ResendProjectInvitationRequest
	(*ResendProjectInvitationResponse)(nil),      // 7: ntt.iam.v1alpha2.ResendProjectInvitationResponse
	(*project_invitation.ProjectInvitation)(nil), // 8: ntt.iam.v1alpha2.ProjectInvitation
}
var edgelq_iam_proto_v1alpha2_project_invitation_custom_proto_depIdxs = []int32{
	8, // 0: ntt.iam.v1alpha2.ListMyProjectInvitationsResponse.project_invitations:type_name -> ntt.iam.v1alpha2.ProjectInvitation
	8, // 1: ntt.iam.v1alpha2.ResendProjectInvitationResponse.project_invitation:type_name -> ntt.iam.v1alpha2.ProjectInvitation
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { edgelq_iam_proto_v1alpha2_project_invitation_custom_proto_init() }
func edgelq_iam_proto_v1alpha2_project_invitation_custom_proto_init() {
	if edgelq_iam_proto_v1alpha2_project_invitation_custom_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		edgelq_iam_proto_v1alpha2_project_invitation_custom_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AcceptProjectInvitationRequest); i {
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
		edgelq_iam_proto_v1alpha2_project_invitation_custom_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AcceptProjectInvitationResponse); i {
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
		edgelq_iam_proto_v1alpha2_project_invitation_custom_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeclineProjectInvitationRequest); i {
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
		edgelq_iam_proto_v1alpha2_project_invitation_custom_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeclineProjectInvitationResponse); i {
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
		edgelq_iam_proto_v1alpha2_project_invitation_custom_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMyProjectInvitationsRequest); i {
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
		edgelq_iam_proto_v1alpha2_project_invitation_custom_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMyProjectInvitationsResponse); i {
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
		edgelq_iam_proto_v1alpha2_project_invitation_custom_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResendProjectInvitationRequest); i {
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
		edgelq_iam_proto_v1alpha2_project_invitation_custom_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResendProjectInvitationResponse); i {
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
			RawDescriptor: edgelq_iam_proto_v1alpha2_project_invitation_custom_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           edgelq_iam_proto_v1alpha2_project_invitation_custom_proto_goTypes,
		DependencyIndexes: edgelq_iam_proto_v1alpha2_project_invitation_custom_proto_depIdxs,
		MessageInfos:      edgelq_iam_proto_v1alpha2_project_invitation_custom_proto_msgTypes,
	}.Build()
	edgelq_iam_proto_v1alpha2_project_invitation_custom_proto = out.File
	edgelq_iam_proto_v1alpha2_project_invitation_custom_proto_rawDesc = nil
	edgelq_iam_proto_v1alpha2_project_invitation_custom_proto_goTypes = nil
	edgelq_iam_proto_v1alpha2_project_invitation_custom_proto_depIdxs = nil
}
