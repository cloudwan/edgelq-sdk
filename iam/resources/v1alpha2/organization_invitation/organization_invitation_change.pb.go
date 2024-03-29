// Code generated by protoc-gen-goten-go
// File: edgelq/iam/proto/v1alpha2/organization_invitation_change.proto
// DO NOT EDIT!!!

package organization_invitation

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
	_ = &organization.Organization{}
	_ = &fieldmaskpb.FieldMask{}
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// OrganizationInvitationChange is used by Watch notifications Responses to
// describe change of single OrganizationInvitation One of Added, Modified,
// Removed
type OrganizationInvitationChange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// OrganizationInvitation change
	//
	// Types that are valid to be assigned to ChangeType:
	//	*OrganizationInvitationChange_Added_
	//	*OrganizationInvitationChange_Modified_
	//	*OrganizationInvitationChange_Current_
	//	*OrganizationInvitationChange_Removed_
	ChangeType isOrganizationInvitationChange_ChangeType `protobuf_oneof:"change_type"`
}

func (m *OrganizationInvitationChange) Reset() {
	*m = OrganizationInvitationChange{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_iam_proto_v1alpha2_organization_invitation_change_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *OrganizationInvitationChange) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*OrganizationInvitationChange) ProtoMessage() {}

func (m *OrganizationInvitationChange) ProtoReflect() preflect.Message {
	mi := &edgelq_iam_proto_v1alpha2_organization_invitation_change_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*OrganizationInvitationChange) GotenMessage() {}

// Deprecated, Use OrganizationInvitationChange.ProtoReflect.Descriptor instead.
func (*OrganizationInvitationChange) Descriptor() ([]byte, []int) {
	return edgelq_iam_proto_v1alpha2_organization_invitation_change_proto_rawDescGZIP(), []int{0}
}

func (m *OrganizationInvitationChange) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *OrganizationInvitationChange) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *OrganizationInvitationChange) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *OrganizationInvitationChange) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

type isOrganizationInvitationChange_ChangeType interface {
	isOrganizationInvitationChange_ChangeType()
}

type OrganizationInvitationChange_Added_ struct {
	// Added is returned when watched document is added, either created or
	// enters Query view
	Added *OrganizationInvitationChange_Added `protobuf:"bytes,1,opt,name=added,proto3,oneof" firestore:"added"`
}
type OrganizationInvitationChange_Modified_ struct {
	// Modified is returned when watched document is modified
	Modified *OrganizationInvitationChange_Modified `protobuf:"bytes,2,opt,name=modified,proto3,oneof" firestore:"modified"`
}
type OrganizationInvitationChange_Current_ struct {
	// Current is returned in stateless watch when document enters query view or
	// is modified within.
	Current *OrganizationInvitationChange_Current `protobuf:"bytes,4,opt,name=current,proto3,oneof" firestore:"current"`
}
type OrganizationInvitationChange_Removed_ struct {
	// Removed is returned when OrganizationInvitation is deleted or leaves
	// Query view
	Removed *OrganizationInvitationChange_Removed `protobuf:"bytes,3,opt,name=removed,proto3,oneof" firestore:"removed"`
}

func (*OrganizationInvitationChange_Added_) isOrganizationInvitationChange_ChangeType()    {}
func (*OrganizationInvitationChange_Modified_) isOrganizationInvitationChange_ChangeType() {}
func (*OrganizationInvitationChange_Current_) isOrganizationInvitationChange_ChangeType()  {}
func (*OrganizationInvitationChange_Removed_) isOrganizationInvitationChange_ChangeType()  {}
func (m *OrganizationInvitationChange) GetChangeType() isOrganizationInvitationChange_ChangeType {
	if m != nil {
		return m.ChangeType
	}
	return nil
}
func (m *OrganizationInvitationChange) GetAdded() *OrganizationInvitationChange_Added {
	if x, ok := m.GetChangeType().(*OrganizationInvitationChange_Added_); ok {
		return x.Added
	}
	return nil
}
func (m *OrganizationInvitationChange) GetModified() *OrganizationInvitationChange_Modified {
	if x, ok := m.GetChangeType().(*OrganizationInvitationChange_Modified_); ok {
		return x.Modified
	}
	return nil
}
func (m *OrganizationInvitationChange) GetCurrent() *OrganizationInvitationChange_Current {
	if x, ok := m.GetChangeType().(*OrganizationInvitationChange_Current_); ok {
		return x.Current
	}
	return nil
}
func (m *OrganizationInvitationChange) GetRemoved() *OrganizationInvitationChange_Removed {
	if x, ok := m.GetChangeType().(*OrganizationInvitationChange_Removed_); ok {
		return x.Removed
	}
	return nil
}
func (m *OrganizationInvitationChange) SetChangeType(ofv isOrganizationInvitationChange_ChangeType) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "isOrganizationInvitationChange_ChangeType", "OrganizationInvitationChange"))
	}
	m.ChangeType = ofv
}
func (m *OrganizationInvitationChange) SetAdded(fv *OrganizationInvitationChange_Added) {
	m.SetChangeType(&OrganizationInvitationChange_Added_{Added: fv})
}
func (m *OrganizationInvitationChange) SetModified(fv *OrganizationInvitationChange_Modified) {
	m.SetChangeType(&OrganizationInvitationChange_Modified_{Modified: fv})
}
func (m *OrganizationInvitationChange) SetCurrent(fv *OrganizationInvitationChange_Current) {
	m.SetChangeType(&OrganizationInvitationChange_Current_{Current: fv})
}
func (m *OrganizationInvitationChange) SetRemoved(fv *OrganizationInvitationChange_Removed) {
	m.SetChangeType(&OrganizationInvitationChange_Removed_{Removed: fv})
}

// OrganizationInvitation has been added to query view
type OrganizationInvitationChange_Added struct {
	state                  protoimpl.MessageState
	sizeCache              protoimpl.SizeCache
	unknownFields          protoimpl.UnknownFields
	OrganizationInvitation *OrganizationInvitation `protobuf:"bytes,1,opt,name=organization_invitation,json=organizationInvitation,proto3" json:"organization_invitation,omitempty" firestore:"organizationInvitation"`
	// Integer describing index of added OrganizationInvitation in resulting
	// query view.
	ViewIndex int32 `protobuf:"varint,2,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty" firestore:"viewIndex"`
}

func (m *OrganizationInvitationChange_Added) Reset() {
	*m = OrganizationInvitationChange_Added{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_iam_proto_v1alpha2_organization_invitation_change_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *OrganizationInvitationChange_Added) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*OrganizationInvitationChange_Added) ProtoMessage() {}

func (m *OrganizationInvitationChange_Added) ProtoReflect() preflect.Message {
	mi := &edgelq_iam_proto_v1alpha2_organization_invitation_change_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*OrganizationInvitationChange_Added) GotenMessage() {}

// Deprecated, Use OrganizationInvitationChange_Added.ProtoReflect.Descriptor instead.
func (*OrganizationInvitationChange_Added) Descriptor() ([]byte, []int) {
	return edgelq_iam_proto_v1alpha2_organization_invitation_change_proto_rawDescGZIP(), []int{0, 0}
}

func (m *OrganizationInvitationChange_Added) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *OrganizationInvitationChange_Added) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *OrganizationInvitationChange_Added) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *OrganizationInvitationChange_Added) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *OrganizationInvitationChange_Added) GetOrganizationInvitation() *OrganizationInvitation {
	if m != nil {
		return m.OrganizationInvitation
	}
	return nil
}

func (m *OrganizationInvitationChange_Added) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *OrganizationInvitationChange_Added) SetOrganizationInvitation(fv *OrganizationInvitation) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "OrganizationInvitation", "OrganizationInvitationChange_Added"))
	}
	m.OrganizationInvitation = fv
}

func (m *OrganizationInvitationChange_Added) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "OrganizationInvitationChange_Added"))
	}
	m.ViewIndex = fv
}

// OrganizationInvitation changed some of it's fields - contains either full
// document or masked change
type OrganizationInvitationChange_Modified struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Name of modified OrganizationInvitation
	Name *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// New version of OrganizationInvitation or masked difference, depending on
	// mask_changes instrumentation of issued
	// [WatchOrganizationInvitationRequest] or
	// [WatchOrganizationInvitationsRequest]
	OrganizationInvitation *OrganizationInvitation `protobuf:"bytes,2,opt,name=organization_invitation,json=organizationInvitation,proto3" json:"organization_invitation,omitempty" firestore:"organizationInvitation"`
	// Used when mask_changes is set, contains field paths of modified
	// properties.
	FieldMask *OrganizationInvitation_FieldMask `protobuf:"bytes,3,opt,customtype=OrganizationInvitation_FieldMask,name=field_mask,json=fieldMask,proto3" json:"field_mask,omitempty" firestore:"fieldMask"`
	// Previous view index specifies previous position of modified
	// OrganizationInvitation. When modification doesn't affect sorted order,
	// value will remain identical to [view_index].
	PreviousViewIndex int32 `protobuf:"varint,4,opt,name=previous_view_index,json=previousViewIndex,proto3" json:"previous_view_index,omitempty" firestore:"previousViewIndex"`
	// Integer specifying OrganizationInvitation new index in resulting query
	// view.
	ViewIndex int32 `protobuf:"varint,5,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty" firestore:"viewIndex"`
}

func (m *OrganizationInvitationChange_Modified) Reset() {
	*m = OrganizationInvitationChange_Modified{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_iam_proto_v1alpha2_organization_invitation_change_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *OrganizationInvitationChange_Modified) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*OrganizationInvitationChange_Modified) ProtoMessage() {}

func (m *OrganizationInvitationChange_Modified) ProtoReflect() preflect.Message {
	mi := &edgelq_iam_proto_v1alpha2_organization_invitation_change_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*OrganizationInvitationChange_Modified) GotenMessage() {}

// Deprecated, Use OrganizationInvitationChange_Modified.ProtoReflect.Descriptor instead.
func (*OrganizationInvitationChange_Modified) Descriptor() ([]byte, []int) {
	return edgelq_iam_proto_v1alpha2_organization_invitation_change_proto_rawDescGZIP(), []int{0, 1}
}

func (m *OrganizationInvitationChange_Modified) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *OrganizationInvitationChange_Modified) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *OrganizationInvitationChange_Modified) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *OrganizationInvitationChange_Modified) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *OrganizationInvitationChange_Modified) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *OrganizationInvitationChange_Modified) GetOrganizationInvitation() *OrganizationInvitation {
	if m != nil {
		return m.OrganizationInvitation
	}
	return nil
}

func (m *OrganizationInvitationChange_Modified) GetFieldMask() *OrganizationInvitation_FieldMask {
	if m != nil {
		return m.FieldMask
	}
	return nil
}

func (m *OrganizationInvitationChange_Modified) GetPreviousViewIndex() int32 {
	if m != nil {
		return m.PreviousViewIndex
	}
	return int32(0)
}

func (m *OrganizationInvitationChange_Modified) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *OrganizationInvitationChange_Modified) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "OrganizationInvitationChange_Modified"))
	}
	m.Name = fv
}

func (m *OrganizationInvitationChange_Modified) SetOrganizationInvitation(fv *OrganizationInvitation) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "OrganizationInvitation", "OrganizationInvitationChange_Modified"))
	}
	m.OrganizationInvitation = fv
}

func (m *OrganizationInvitationChange_Modified) SetFieldMask(fv *OrganizationInvitation_FieldMask) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "FieldMask", "OrganizationInvitationChange_Modified"))
	}
	m.FieldMask = fv
}

func (m *OrganizationInvitationChange_Modified) SetPreviousViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "PreviousViewIndex", "OrganizationInvitationChange_Modified"))
	}
	m.PreviousViewIndex = fv
}

func (m *OrganizationInvitationChange_Modified) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "OrganizationInvitationChange_Modified"))
	}
	m.ViewIndex = fv
}

// OrganizationInvitation has been added or modified in a query view. Version
// used for stateless watching
type OrganizationInvitationChange_Current struct {
	state                  protoimpl.MessageState
	sizeCache              protoimpl.SizeCache
	unknownFields          protoimpl.UnknownFields
	OrganizationInvitation *OrganizationInvitation `protobuf:"bytes,1,opt,name=organization_invitation,json=organizationInvitation,proto3" json:"organization_invitation,omitempty" firestore:"organizationInvitation"`
}

func (m *OrganizationInvitationChange_Current) Reset() {
	*m = OrganizationInvitationChange_Current{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_iam_proto_v1alpha2_organization_invitation_change_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *OrganizationInvitationChange_Current) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*OrganizationInvitationChange_Current) ProtoMessage() {}

func (m *OrganizationInvitationChange_Current) ProtoReflect() preflect.Message {
	mi := &edgelq_iam_proto_v1alpha2_organization_invitation_change_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*OrganizationInvitationChange_Current) GotenMessage() {}

// Deprecated, Use OrganizationInvitationChange_Current.ProtoReflect.Descriptor instead.
func (*OrganizationInvitationChange_Current) Descriptor() ([]byte, []int) {
	return edgelq_iam_proto_v1alpha2_organization_invitation_change_proto_rawDescGZIP(), []int{0, 2}
}

func (m *OrganizationInvitationChange_Current) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *OrganizationInvitationChange_Current) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *OrganizationInvitationChange_Current) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *OrganizationInvitationChange_Current) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *OrganizationInvitationChange_Current) GetOrganizationInvitation() *OrganizationInvitation {
	if m != nil {
		return m.OrganizationInvitation
	}
	return nil
}

func (m *OrganizationInvitationChange_Current) SetOrganizationInvitation(fv *OrganizationInvitation) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "OrganizationInvitation", "OrganizationInvitationChange_Current"))
	}
	m.OrganizationInvitation = fv
}

// Removed is returned when OrganizationInvitation is deleted or leaves Query
// view
type OrganizationInvitationChange_Removed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Name          *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// Integer specifying removed OrganizationInvitation index. Not populated in
	// stateless watch type.
	ViewIndex int32 `protobuf:"varint,2,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty" firestore:"viewIndex"`
}

func (m *OrganizationInvitationChange_Removed) Reset() {
	*m = OrganizationInvitationChange_Removed{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_iam_proto_v1alpha2_organization_invitation_change_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *OrganizationInvitationChange_Removed) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*OrganizationInvitationChange_Removed) ProtoMessage() {}

func (m *OrganizationInvitationChange_Removed) ProtoReflect() preflect.Message {
	mi := &edgelq_iam_proto_v1alpha2_organization_invitation_change_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*OrganizationInvitationChange_Removed) GotenMessage() {}

// Deprecated, Use OrganizationInvitationChange_Removed.ProtoReflect.Descriptor instead.
func (*OrganizationInvitationChange_Removed) Descriptor() ([]byte, []int) {
	return edgelq_iam_proto_v1alpha2_organization_invitation_change_proto_rawDescGZIP(), []int{0, 3}
}

func (m *OrganizationInvitationChange_Removed) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *OrganizationInvitationChange_Removed) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *OrganizationInvitationChange_Removed) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *OrganizationInvitationChange_Removed) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *OrganizationInvitationChange_Removed) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *OrganizationInvitationChange_Removed) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *OrganizationInvitationChange_Removed) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "OrganizationInvitationChange_Removed"))
	}
	m.Name = fv
}

func (m *OrganizationInvitationChange_Removed) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "OrganizationInvitationChange_Removed"))
	}
	m.ViewIndex = fv
}

var edgelq_iam_proto_v1alpha2_organization_invitation_change_proto preflect.FileDescriptor

var edgelq_iam_proto_v1alpha2_organization_invitation_change_proto_rawDesc = []byte{
	0x0a, 0x3e, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x6f, 0x72, 0x67, 0x61,
	0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x10, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x32, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20,
	0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x37, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x6f, 0x72, 0x67, 0x61,
	0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbe, 0x08, 0x0a, 0x1c, 0x4f, 0x72,
	0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x4c, 0x0a, 0x05, 0x61, 0x64,
	0x64, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x6e, 0x74, 0x74, 0x2e,
	0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x4f, 0x72, 0x67,
	0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x65, 0x64, 0x48,
	0x00, 0x52, 0x05, 0x61, 0x64, 0x64, 0x65, 0x64, 0x12, 0x55, 0x0a, 0x08, 0x6d, 0x6f, 0x64, 0x69,
	0x66, 0x69, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x6e, 0x74, 0x74,
	0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x4f, 0x72,
	0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66,
	0x69, 0x65, 0x64, 0x48, 0x00, 0x52, 0x08, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12,
	0x52, 0x0a, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x36, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x32, 0x2e, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x07, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x12, 0x52, 0x0a, 0x07, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x48, 0x00, 0x52, 0x07,
	0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x1a, 0x89, 0x01, 0x0a, 0x05, 0x41, 0x64, 0x64, 0x65,
	0x64, 0x12, 0x61, 0x0a, 0x17, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x28, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x16, 0x6f, 0x72,
	0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x76, 0x69, 0x65, 0x77, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x1a, 0xcb, 0x02, 0x0a, 0x08, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64,
	0x12, 0x32, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1e,
	0xb2, 0xda, 0x21, 0x1a, 0x0a, 0x18, 0x0a, 0x16, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x61, 0x0a, 0x17, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x16, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x76,
	0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x59, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x42, 0x1e, 0xb2, 0xda, 0x21, 0x1a, 0x32, 0x18, 0x0a,
	0x16, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x76,
	0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61,
	0x73, 0x6b, 0x12, 0x2e, 0x0a, 0x13, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x5f, 0x76,
	0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x11, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x56, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x76, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x1a, 0x6c, 0x0a, 0x07, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x61, 0x0a, 0x17,
	0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x76,
	0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e,
	0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32,
	0x2e, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x76,
	0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x16, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a,
	0x5c, 0x0a, 0x07, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x12, 0x32, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1e, 0xb2, 0xda, 0x21, 0x1a, 0x0a, 0x18,
	0x0a, 0x16, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e,
	0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x76, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x3a, 0x1c, 0x9a,
	0xd9, 0x21, 0x18, 0x0a, 0x16, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0d, 0x0a, 0x0b, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x42, 0xa5, 0x01, 0xe8, 0xde, 0x21,
	0x00, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x70,
	0x62, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x42, 0x21, 0x4f, 0x72, 0x67, 0x61,
	0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x00, 0x5a,
	0x61, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x32, 0x2f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3b, 0x6f, 0x72, 0x67, 0x61, 0x6e,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	edgelq_iam_proto_v1alpha2_organization_invitation_change_proto_rawDescOnce sync.Once
	edgelq_iam_proto_v1alpha2_organization_invitation_change_proto_rawDescData = edgelq_iam_proto_v1alpha2_organization_invitation_change_proto_rawDesc
)

func edgelq_iam_proto_v1alpha2_organization_invitation_change_proto_rawDescGZIP() []byte {
	edgelq_iam_proto_v1alpha2_organization_invitation_change_proto_rawDescOnce.Do(func() {
		edgelq_iam_proto_v1alpha2_organization_invitation_change_proto_rawDescData = protoimpl.X.CompressGZIP(edgelq_iam_proto_v1alpha2_organization_invitation_change_proto_rawDescData)
	})
	return edgelq_iam_proto_v1alpha2_organization_invitation_change_proto_rawDescData
}

var edgelq_iam_proto_v1alpha2_organization_invitation_change_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var edgelq_iam_proto_v1alpha2_organization_invitation_change_proto_goTypes = []interface{}{
	(*OrganizationInvitationChange)(nil),          // 0: ntt.iam.v1alpha2.OrganizationInvitationChange
	(*OrganizationInvitationChange_Added)(nil),    // 1: ntt.iam.v1alpha2.OrganizationInvitationChange.Added
	(*OrganizationInvitationChange_Modified)(nil), // 2: ntt.iam.v1alpha2.OrganizationInvitationChange.Modified
	(*OrganizationInvitationChange_Current)(nil),  // 3: ntt.iam.v1alpha2.OrganizationInvitationChange.Current
	(*OrganizationInvitationChange_Removed)(nil),  // 4: ntt.iam.v1alpha2.OrganizationInvitationChange.Removed
	(*OrganizationInvitation)(nil),                // 5: ntt.iam.v1alpha2.OrganizationInvitation
	(*OrganizationInvitation_FieldMask)(nil),      // 6: ntt.iam.v1alpha2.OrganizationInvitation_FieldMask
}
var edgelq_iam_proto_v1alpha2_organization_invitation_change_proto_depIdxs = []int32{
	1, // 0: ntt.iam.v1alpha2.OrganizationInvitationChange.added:type_name -> ntt.iam.v1alpha2.OrganizationInvitationChange.Added
	2, // 1: ntt.iam.v1alpha2.OrganizationInvitationChange.modified:type_name -> ntt.iam.v1alpha2.OrganizationInvitationChange.Modified
	3, // 2: ntt.iam.v1alpha2.OrganizationInvitationChange.current:type_name -> ntt.iam.v1alpha2.OrganizationInvitationChange.Current
	4, // 3: ntt.iam.v1alpha2.OrganizationInvitationChange.removed:type_name -> ntt.iam.v1alpha2.OrganizationInvitationChange.Removed
	5, // 4: ntt.iam.v1alpha2.OrganizationInvitationChange.Added.organization_invitation:type_name -> ntt.iam.v1alpha2.OrganizationInvitation
	5, // 5: ntt.iam.v1alpha2.OrganizationInvitationChange.Modified.organization_invitation:type_name -> ntt.iam.v1alpha2.OrganizationInvitation
	6, // 6: ntt.iam.v1alpha2.OrganizationInvitationChange.Modified.field_mask:type_name -> ntt.iam.v1alpha2.OrganizationInvitation_FieldMask
	5, // 7: ntt.iam.v1alpha2.OrganizationInvitationChange.Current.organization_invitation:type_name -> ntt.iam.v1alpha2.OrganizationInvitation
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { edgelq_iam_proto_v1alpha2_organization_invitation_change_proto_init() }
func edgelq_iam_proto_v1alpha2_organization_invitation_change_proto_init() {
	if edgelq_iam_proto_v1alpha2_organization_invitation_change_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		edgelq_iam_proto_v1alpha2_organization_invitation_change_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrganizationInvitationChange); i {
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
		edgelq_iam_proto_v1alpha2_organization_invitation_change_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrganizationInvitationChange_Added); i {
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
		edgelq_iam_proto_v1alpha2_organization_invitation_change_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrganizationInvitationChange_Modified); i {
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
		edgelq_iam_proto_v1alpha2_organization_invitation_change_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrganizationInvitationChange_Current); i {
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
		edgelq_iam_proto_v1alpha2_organization_invitation_change_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrganizationInvitationChange_Removed); i {
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

	edgelq_iam_proto_v1alpha2_organization_invitation_change_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*OrganizationInvitationChange_Added_)(nil),
		(*OrganizationInvitationChange_Modified_)(nil),
		(*OrganizationInvitationChange_Current_)(nil),
		(*OrganizationInvitationChange_Removed_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: edgelq_iam_proto_v1alpha2_organization_invitation_change_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           edgelq_iam_proto_v1alpha2_organization_invitation_change_proto_goTypes,
		DependencyIndexes: edgelq_iam_proto_v1alpha2_organization_invitation_change_proto_depIdxs,
		MessageInfos:      edgelq_iam_proto_v1alpha2_organization_invitation_change_proto_msgTypes,
	}.Build()
	edgelq_iam_proto_v1alpha2_organization_invitation_change_proto = out.File
	edgelq_iam_proto_v1alpha2_organization_invitation_change_proto_rawDesc = nil
	edgelq_iam_proto_v1alpha2_organization_invitation_change_proto_goTypes = nil
	edgelq_iam_proto_v1alpha2_organization_invitation_change_proto_depIdxs = nil
}
