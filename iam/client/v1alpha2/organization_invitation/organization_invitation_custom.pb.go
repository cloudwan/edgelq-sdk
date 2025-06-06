// Code generated by protoc-gen-goten-go
// File: edgelq/iam/proto/v1alpha2/organization_invitation_custom.proto
// DO NOT EDIT!!!

package organization_invitation_client

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
	organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/organization"
	organization_invitation "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/organization_invitation"
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
	_ = &organization.Organization{}
	_ = &organization_invitation.OrganizationInvitation{}
	_ = &project_invitation.ProjectInvitation{}
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Request message for method
// [AcceptOrganizationInvitation][ntt.iam.v1alpha2.AcceptOrganizationInvitation]
type AcceptOrganizationInvitationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// name of ntt.iam.v1alpha2.OrganizationInvitation
	Name *organization_invitation.Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty"`
}

func (m *AcceptOrganizationInvitationRequest) Reset() {
	*m = AcceptOrganizationInvitationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_iam_proto_v1alpha2_organization_invitation_custom_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *AcceptOrganizationInvitationRequest) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*AcceptOrganizationInvitationRequest) ProtoMessage() {}

func (m *AcceptOrganizationInvitationRequest) ProtoReflect() preflect.Message {
	mi := &edgelq_iam_proto_v1alpha2_organization_invitation_custom_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*AcceptOrganizationInvitationRequest) GotenMessage() {}

// Deprecated, Use AcceptOrganizationInvitationRequest.ProtoReflect.Descriptor instead.
func (*AcceptOrganizationInvitationRequest) Descriptor() ([]byte, []int) {
	return edgelq_iam_proto_v1alpha2_organization_invitation_custom_proto_rawDescGZIP(), []int{0}
}

func (m *AcceptOrganizationInvitationRequest) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *AcceptOrganizationInvitationRequest) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *AcceptOrganizationInvitationRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *AcceptOrganizationInvitationRequest) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *AcceptOrganizationInvitationRequest) GetName() *organization_invitation.Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *AcceptOrganizationInvitationRequest) SetName(fv *organization_invitation.Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "AcceptOrganizationInvitationRequest"))
	}
	m.Name = fv
}

// Response message for method
// [AcceptOrganizationInvitation][ntt.iam.v1alpha2.AcceptOrganizationInvitation]
type AcceptOrganizationInvitationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (m *AcceptOrganizationInvitationResponse) Reset() {
	*m = AcceptOrganizationInvitationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_iam_proto_v1alpha2_organization_invitation_custom_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *AcceptOrganizationInvitationResponse) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*AcceptOrganizationInvitationResponse) ProtoMessage() {}

func (m *AcceptOrganizationInvitationResponse) ProtoReflect() preflect.Message {
	mi := &edgelq_iam_proto_v1alpha2_organization_invitation_custom_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*AcceptOrganizationInvitationResponse) GotenMessage() {}

// Deprecated, Use AcceptOrganizationInvitationResponse.ProtoReflect.Descriptor instead.
func (*AcceptOrganizationInvitationResponse) Descriptor() ([]byte, []int) {
	return edgelq_iam_proto_v1alpha2_organization_invitation_custom_proto_rawDescGZIP(), []int{1}
}

func (m *AcceptOrganizationInvitationResponse) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *AcceptOrganizationInvitationResponse) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *AcceptOrganizationInvitationResponse) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *AcceptOrganizationInvitationResponse) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

// Request message for method
// [DeclineOrganizationInvitation][ntt.iam.v1alpha2.DeclineOrganizationInvitation]
type DeclineOrganizationInvitationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	//  name of ntt.iam.v1alpha2.OrganizationInvitation
	Name *organization_invitation.Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty"`
	// Additional filter for invitations, e.g. state = "PENDING"
	Filter *project_invitation.Filter `protobuf:"bytes,5,opt,customtype=Filter,name=filter,proto3" json:"filter,omitempty"`
}

func (m *DeclineOrganizationInvitationRequest) Reset() {
	*m = DeclineOrganizationInvitationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_iam_proto_v1alpha2_organization_invitation_custom_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *DeclineOrganizationInvitationRequest) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*DeclineOrganizationInvitationRequest) ProtoMessage() {}

func (m *DeclineOrganizationInvitationRequest) ProtoReflect() preflect.Message {
	mi := &edgelq_iam_proto_v1alpha2_organization_invitation_custom_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*DeclineOrganizationInvitationRequest) GotenMessage() {}

// Deprecated, Use DeclineOrganizationInvitationRequest.ProtoReflect.Descriptor instead.
func (*DeclineOrganizationInvitationRequest) Descriptor() ([]byte, []int) {
	return edgelq_iam_proto_v1alpha2_organization_invitation_custom_proto_rawDescGZIP(), []int{2}
}

func (m *DeclineOrganizationInvitationRequest) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *DeclineOrganizationInvitationRequest) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *DeclineOrganizationInvitationRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *DeclineOrganizationInvitationRequest) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *DeclineOrganizationInvitationRequest) GetName() *organization_invitation.Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *DeclineOrganizationInvitationRequest) GetFilter() *project_invitation.Filter {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *DeclineOrganizationInvitationRequest) SetName(fv *organization_invitation.Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "DeclineOrganizationInvitationRequest"))
	}
	m.Name = fv
}

func (m *DeclineOrganizationInvitationRequest) SetFilter(fv *project_invitation.Filter) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Filter", "DeclineOrganizationInvitationRequest"))
	}
	m.Filter = fv
}

// Response message for method
// [DeclineOrganizationInvitation][ntt.iam.v1alpha2.DeclineOrganizationInvitation]
type DeclineOrganizationInvitationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (m *DeclineOrganizationInvitationResponse) Reset() {
	*m = DeclineOrganizationInvitationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_iam_proto_v1alpha2_organization_invitation_custom_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *DeclineOrganizationInvitationResponse) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*DeclineOrganizationInvitationResponse) ProtoMessage() {}

func (m *DeclineOrganizationInvitationResponse) ProtoReflect() preflect.Message {
	mi := &edgelq_iam_proto_v1alpha2_organization_invitation_custom_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*DeclineOrganizationInvitationResponse) GotenMessage() {}

// Deprecated, Use DeclineOrganizationInvitationResponse.ProtoReflect.Descriptor instead.
func (*DeclineOrganizationInvitationResponse) Descriptor() ([]byte, []int) {
	return edgelq_iam_proto_v1alpha2_organization_invitation_custom_proto_rawDescGZIP(), []int{3}
}

func (m *DeclineOrganizationInvitationResponse) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *DeclineOrganizationInvitationResponse) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *DeclineOrganizationInvitationResponse) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *DeclineOrganizationInvitationResponse) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

// Request message for method
// [ListMyOrganizationInvitations][ntt.iam.v1alpha2.ListMyOrganizationInvitations]
type ListMyOrganizationInvitationsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Parent name of ntt.iam.v1alpha2.OrganizationInvitation
	Parent *organization_invitation.ParentName `protobuf:"bytes,1,opt,customtype=ParentName,name=parent,proto3" json:"parent,omitempty"`
	// Additional filter for invitations, e.g. state = "PENDING"
	Filter *organization_invitation.Filter `protobuf:"bytes,5,opt,customtype=Filter,name=filter,proto3" json:"filter,omitempty"`
}

func (m *ListMyOrganizationInvitationsRequest) Reset() {
	*m = ListMyOrganizationInvitationsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_iam_proto_v1alpha2_organization_invitation_custom_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *ListMyOrganizationInvitationsRequest) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*ListMyOrganizationInvitationsRequest) ProtoMessage() {}

func (m *ListMyOrganizationInvitationsRequest) ProtoReflect() preflect.Message {
	mi := &edgelq_iam_proto_v1alpha2_organization_invitation_custom_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*ListMyOrganizationInvitationsRequest) GotenMessage() {}

// Deprecated, Use ListMyOrganizationInvitationsRequest.ProtoReflect.Descriptor instead.
func (*ListMyOrganizationInvitationsRequest) Descriptor() ([]byte, []int) {
	return edgelq_iam_proto_v1alpha2_organization_invitation_custom_proto_rawDescGZIP(), []int{4}
}

func (m *ListMyOrganizationInvitationsRequest) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *ListMyOrganizationInvitationsRequest) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *ListMyOrganizationInvitationsRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *ListMyOrganizationInvitationsRequest) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *ListMyOrganizationInvitationsRequest) GetParent() *organization_invitation.ParentName {
	if m != nil {
		return m.Parent
	}
	return nil
}

func (m *ListMyOrganizationInvitationsRequest) GetFilter() *organization_invitation.Filter {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *ListMyOrganizationInvitationsRequest) SetParent(fv *organization_invitation.ParentName) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Parent", "ListMyOrganizationInvitationsRequest"))
	}
	m.Parent = fv
}

func (m *ListMyOrganizationInvitationsRequest) SetFilter(fv *organization_invitation.Filter) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Filter", "ListMyOrganizationInvitationsRequest"))
	}
	m.Filter = fv
}

// Response message for method
// [ListMyOrganizationInvitations][ntt.iam.v1alpha2.ListMyOrganizationInvitations]
type ListMyOrganizationInvitationsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// The list of OrganizationInvitations
	OrganizationInvitations []*organization_invitation.OrganizationInvitation `protobuf:"bytes,1,rep,name=organization_invitations,json=organizationInvitations,proto3" json:"organization_invitations,omitempty"`
}

func (m *ListMyOrganizationInvitationsResponse) Reset() {
	*m = ListMyOrganizationInvitationsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_iam_proto_v1alpha2_organization_invitation_custom_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *ListMyOrganizationInvitationsResponse) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*ListMyOrganizationInvitationsResponse) ProtoMessage() {}

func (m *ListMyOrganizationInvitationsResponse) ProtoReflect() preflect.Message {
	mi := &edgelq_iam_proto_v1alpha2_organization_invitation_custom_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*ListMyOrganizationInvitationsResponse) GotenMessage() {}

// Deprecated, Use ListMyOrganizationInvitationsResponse.ProtoReflect.Descriptor instead.
func (*ListMyOrganizationInvitationsResponse) Descriptor() ([]byte, []int) {
	return edgelq_iam_proto_v1alpha2_organization_invitation_custom_proto_rawDescGZIP(), []int{5}
}

func (m *ListMyOrganizationInvitationsResponse) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *ListMyOrganizationInvitationsResponse) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *ListMyOrganizationInvitationsResponse) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *ListMyOrganizationInvitationsResponse) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *ListMyOrganizationInvitationsResponse) GetOrganizationInvitations() []*organization_invitation.OrganizationInvitation {
	if m != nil {
		return m.OrganizationInvitations
	}
	return nil
}

func (m *ListMyOrganizationInvitationsResponse) SetOrganizationInvitations(fv []*organization_invitation.OrganizationInvitation) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "OrganizationInvitations", "ListMyOrganizationInvitationsResponse"))
	}
	m.OrganizationInvitations = fv
}

var edgelq_iam_proto_v1alpha2_organization_invitation_custom_proto preflect.FileDescriptor

var edgelq_iam_proto_v1alpha2_organization_invitation_custom_proto_rawDesc = []byte{
	0x0a, 0x3e, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x6f, 0x72, 0x67, 0x61,
	0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x10, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x32, 0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x61, 0x75,
	0x64, 0x69, 0x74, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x37, 0x65, 0x64, 0x67,
	0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3e, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f,
	0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x76,
	0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x59, 0x0a, 0x23, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x4f, 0x72,
	0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1e, 0xb2, 0xda, 0x21, 0x1a, 0x0a,
	0x18, 0x0a, 0x16, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x26, 0x0a, 0x24, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x8d, 0x01, 0x0a, 0x24, 0x44, 0x65, 0x63, 0x6c,
	0x69, 0x6e, 0x65, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x32, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1e,
	0xb2, 0xda, 0x21, 0x1a, 0x0a, 0x18, 0x0a, 0x16, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x31, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x19, 0xb2, 0xda, 0x21, 0x15, 0x1a, 0x13, 0x0a, 0x11, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x27, 0x0a, 0x25, 0x44, 0x65, 0x63, 0x6c, 0x69,
	0x6e, 0x65, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e,
	0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x96, 0x01, 0x0a, 0x24, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x79, 0x4f, 0x72, 0x67, 0x61, 0x6e,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x06, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1e, 0xb2, 0xda, 0x21, 0x1a, 0x3a,
	0x18, 0x0a, 0x16, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x12, 0x36, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x1e, 0xb2, 0xda, 0x21, 0x1a, 0x1a, 0x18, 0x0a, 0x16, 0x4f, 0x72, 0x67, 0x61, 0x6e,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x8c, 0x01, 0x0a, 0x25, 0x4c, 0x69,
	0x73, 0x74, 0x4d, 0x79, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x63, 0x0a, 0x18, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x17, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x76,
	0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0xa5, 0x01, 0x0a, 0x17, 0x63, 0x6f, 0x6d,
	0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x32, 0x42, 0x21, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x00, 0x5a, 0x65, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65,
	0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x3b, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	edgelq_iam_proto_v1alpha2_organization_invitation_custom_proto_rawDescOnce sync.Once
	edgelq_iam_proto_v1alpha2_organization_invitation_custom_proto_rawDescData = edgelq_iam_proto_v1alpha2_organization_invitation_custom_proto_rawDesc
)

func edgelq_iam_proto_v1alpha2_organization_invitation_custom_proto_rawDescGZIP() []byte {
	edgelq_iam_proto_v1alpha2_organization_invitation_custom_proto_rawDescOnce.Do(func() {
		edgelq_iam_proto_v1alpha2_organization_invitation_custom_proto_rawDescData = protoimpl.X.CompressGZIP(edgelq_iam_proto_v1alpha2_organization_invitation_custom_proto_rawDescData)
	})
	return edgelq_iam_proto_v1alpha2_organization_invitation_custom_proto_rawDescData
}

var edgelq_iam_proto_v1alpha2_organization_invitation_custom_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var edgelq_iam_proto_v1alpha2_organization_invitation_custom_proto_goTypes = []interface{}{
	(*AcceptOrganizationInvitationRequest)(nil),            // 0: ntt.iam.v1alpha2.AcceptOrganizationInvitationRequest
	(*AcceptOrganizationInvitationResponse)(nil),           // 1: ntt.iam.v1alpha2.AcceptOrganizationInvitationResponse
	(*DeclineOrganizationInvitationRequest)(nil),           // 2: ntt.iam.v1alpha2.DeclineOrganizationInvitationRequest
	(*DeclineOrganizationInvitationResponse)(nil),          // 3: ntt.iam.v1alpha2.DeclineOrganizationInvitationResponse
	(*ListMyOrganizationInvitationsRequest)(nil),           // 4: ntt.iam.v1alpha2.ListMyOrganizationInvitationsRequest
	(*ListMyOrganizationInvitationsResponse)(nil),          // 5: ntt.iam.v1alpha2.ListMyOrganizationInvitationsResponse
	(*organization_invitation.OrganizationInvitation)(nil), // 6: ntt.iam.v1alpha2.OrganizationInvitation
}
var edgelq_iam_proto_v1alpha2_organization_invitation_custom_proto_depIdxs = []int32{
	6, // 0: ntt.iam.v1alpha2.ListMyOrganizationInvitationsResponse.organization_invitations:type_name -> ntt.iam.v1alpha2.OrganizationInvitation
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { edgelq_iam_proto_v1alpha2_organization_invitation_custom_proto_init() }
func edgelq_iam_proto_v1alpha2_organization_invitation_custom_proto_init() {
	if edgelq_iam_proto_v1alpha2_organization_invitation_custom_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		edgelq_iam_proto_v1alpha2_organization_invitation_custom_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AcceptOrganizationInvitationRequest); i {
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
		edgelq_iam_proto_v1alpha2_organization_invitation_custom_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AcceptOrganizationInvitationResponse); i {
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
		edgelq_iam_proto_v1alpha2_organization_invitation_custom_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeclineOrganizationInvitationRequest); i {
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
		edgelq_iam_proto_v1alpha2_organization_invitation_custom_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeclineOrganizationInvitationResponse); i {
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
		edgelq_iam_proto_v1alpha2_organization_invitation_custom_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMyOrganizationInvitationsRequest); i {
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
		edgelq_iam_proto_v1alpha2_organization_invitation_custom_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMyOrganizationInvitationsResponse); i {
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
			RawDescriptor: edgelq_iam_proto_v1alpha2_organization_invitation_custom_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           edgelq_iam_proto_v1alpha2_organization_invitation_custom_proto_goTypes,
		DependencyIndexes: edgelq_iam_proto_v1alpha2_organization_invitation_custom_proto_depIdxs,
		MessageInfos:      edgelq_iam_proto_v1alpha2_organization_invitation_custom_proto_msgTypes,
	}.Build()
	edgelq_iam_proto_v1alpha2_organization_invitation_custom_proto = out.File
	edgelq_iam_proto_v1alpha2_organization_invitation_custom_proto_rawDesc = nil
	edgelq_iam_proto_v1alpha2_organization_invitation_custom_proto_goTypes = nil
	edgelq_iam_proto_v1alpha2_organization_invitation_custom_proto_depIdxs = nil
}
