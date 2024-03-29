// Code generated by protoc-gen-goten-go
// File: edgelq/iam/proto/v1/project_invitation_change.proto
// DO NOT EDIT!!!

package project_invitation

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
	project "github.com/cloudwan/edgelq-sdk/iam/resources/v1/project"
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
	_ = &project.Project{}
	_ = &fieldmaskpb.FieldMask{}
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ProjectInvitationChange is used by Watch notifications Responses to describe
// change of single ProjectInvitation One of Added, Modified, Removed
type ProjectInvitationChange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// ProjectInvitation change
	//
	// Types that are valid to be assigned to ChangeType:
	//	*ProjectInvitationChange_Added_
	//	*ProjectInvitationChange_Modified_
	//	*ProjectInvitationChange_Current_
	//	*ProjectInvitationChange_Removed_
	ChangeType isProjectInvitationChange_ChangeType `protobuf_oneof:"change_type"`
}

func (m *ProjectInvitationChange) Reset() {
	*m = ProjectInvitationChange{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_iam_proto_v1_project_invitation_change_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *ProjectInvitationChange) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*ProjectInvitationChange) ProtoMessage() {}

func (m *ProjectInvitationChange) ProtoReflect() preflect.Message {
	mi := &edgelq_iam_proto_v1_project_invitation_change_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*ProjectInvitationChange) GotenMessage() {}

// Deprecated, Use ProjectInvitationChange.ProtoReflect.Descriptor instead.
func (*ProjectInvitationChange) Descriptor() ([]byte, []int) {
	return edgelq_iam_proto_v1_project_invitation_change_proto_rawDescGZIP(), []int{0}
}

func (m *ProjectInvitationChange) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *ProjectInvitationChange) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *ProjectInvitationChange) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *ProjectInvitationChange) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

type isProjectInvitationChange_ChangeType interface {
	isProjectInvitationChange_ChangeType()
}

type ProjectInvitationChange_Added_ struct {
	// Added is returned when watched document is added, either created or
	// enters Query view
	Added *ProjectInvitationChange_Added `protobuf:"bytes,1,opt,name=added,proto3,oneof" firestore:"added"`
}
type ProjectInvitationChange_Modified_ struct {
	// Modified is returned when watched document is modified
	Modified *ProjectInvitationChange_Modified `protobuf:"bytes,2,opt,name=modified,proto3,oneof" firestore:"modified"`
}
type ProjectInvitationChange_Current_ struct {
	// Current is returned in stateless watch when document enters query view or
	// is modified within.
	Current *ProjectInvitationChange_Current `protobuf:"bytes,4,opt,name=current,proto3,oneof" firestore:"current"`
}
type ProjectInvitationChange_Removed_ struct {
	// Removed is returned when ProjectInvitation is deleted or leaves Query
	// view
	Removed *ProjectInvitationChange_Removed `protobuf:"bytes,3,opt,name=removed,proto3,oneof" firestore:"removed"`
}

func (*ProjectInvitationChange_Added_) isProjectInvitationChange_ChangeType()    {}
func (*ProjectInvitationChange_Modified_) isProjectInvitationChange_ChangeType() {}
func (*ProjectInvitationChange_Current_) isProjectInvitationChange_ChangeType()  {}
func (*ProjectInvitationChange_Removed_) isProjectInvitationChange_ChangeType()  {}
func (m *ProjectInvitationChange) GetChangeType() isProjectInvitationChange_ChangeType {
	if m != nil {
		return m.ChangeType
	}
	return nil
}
func (m *ProjectInvitationChange) GetAdded() *ProjectInvitationChange_Added {
	if x, ok := m.GetChangeType().(*ProjectInvitationChange_Added_); ok {
		return x.Added
	}
	return nil
}
func (m *ProjectInvitationChange) GetModified() *ProjectInvitationChange_Modified {
	if x, ok := m.GetChangeType().(*ProjectInvitationChange_Modified_); ok {
		return x.Modified
	}
	return nil
}
func (m *ProjectInvitationChange) GetCurrent() *ProjectInvitationChange_Current {
	if x, ok := m.GetChangeType().(*ProjectInvitationChange_Current_); ok {
		return x.Current
	}
	return nil
}
func (m *ProjectInvitationChange) GetRemoved() *ProjectInvitationChange_Removed {
	if x, ok := m.GetChangeType().(*ProjectInvitationChange_Removed_); ok {
		return x.Removed
	}
	return nil
}
func (m *ProjectInvitationChange) SetChangeType(ofv isProjectInvitationChange_ChangeType) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "isProjectInvitationChange_ChangeType", "ProjectInvitationChange"))
	}
	m.ChangeType = ofv
}
func (m *ProjectInvitationChange) SetAdded(fv *ProjectInvitationChange_Added) {
	m.SetChangeType(&ProjectInvitationChange_Added_{Added: fv})
}
func (m *ProjectInvitationChange) SetModified(fv *ProjectInvitationChange_Modified) {
	m.SetChangeType(&ProjectInvitationChange_Modified_{Modified: fv})
}
func (m *ProjectInvitationChange) SetCurrent(fv *ProjectInvitationChange_Current) {
	m.SetChangeType(&ProjectInvitationChange_Current_{Current: fv})
}
func (m *ProjectInvitationChange) SetRemoved(fv *ProjectInvitationChange_Removed) {
	m.SetChangeType(&ProjectInvitationChange_Removed_{Removed: fv})
}

// ProjectInvitation has been added to query view
type ProjectInvitationChange_Added struct {
	state             protoimpl.MessageState
	sizeCache         protoimpl.SizeCache
	unknownFields     protoimpl.UnknownFields
	ProjectInvitation *ProjectInvitation `protobuf:"bytes,1,opt,name=project_invitation,json=projectInvitation,proto3" json:"project_invitation,omitempty" firestore:"projectInvitation"`
	// Integer describing index of added ProjectInvitation in resulting query
	// view.
	ViewIndex int32 `protobuf:"varint,2,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty" firestore:"viewIndex"`
}

func (m *ProjectInvitationChange_Added) Reset() {
	*m = ProjectInvitationChange_Added{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_iam_proto_v1_project_invitation_change_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *ProjectInvitationChange_Added) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*ProjectInvitationChange_Added) ProtoMessage() {}

func (m *ProjectInvitationChange_Added) ProtoReflect() preflect.Message {
	mi := &edgelq_iam_proto_v1_project_invitation_change_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*ProjectInvitationChange_Added) GotenMessage() {}

// Deprecated, Use ProjectInvitationChange_Added.ProtoReflect.Descriptor instead.
func (*ProjectInvitationChange_Added) Descriptor() ([]byte, []int) {
	return edgelq_iam_proto_v1_project_invitation_change_proto_rawDescGZIP(), []int{0, 0}
}

func (m *ProjectInvitationChange_Added) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *ProjectInvitationChange_Added) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *ProjectInvitationChange_Added) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *ProjectInvitationChange_Added) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *ProjectInvitationChange_Added) GetProjectInvitation() *ProjectInvitation {
	if m != nil {
		return m.ProjectInvitation
	}
	return nil
}

func (m *ProjectInvitationChange_Added) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *ProjectInvitationChange_Added) SetProjectInvitation(fv *ProjectInvitation) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ProjectInvitation", "ProjectInvitationChange_Added"))
	}
	m.ProjectInvitation = fv
}

func (m *ProjectInvitationChange_Added) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "ProjectInvitationChange_Added"))
	}
	m.ViewIndex = fv
}

// ProjectInvitation changed some of it's fields - contains either full
// document or masked change
type ProjectInvitationChange_Modified struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Name of modified ProjectInvitation
	Name *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// New version of ProjectInvitation or masked difference, depending on
	// mask_changes instrumentation of issued [WatchProjectInvitationRequest] or
	// [WatchProjectInvitationsRequest]
	ProjectInvitation *ProjectInvitation `protobuf:"bytes,2,opt,name=project_invitation,json=projectInvitation,proto3" json:"project_invitation,omitempty" firestore:"projectInvitation"`
	// Used when mask_changes is set, contains field paths of modified
	// properties.
	FieldMask *ProjectInvitation_FieldMask `protobuf:"bytes,3,opt,customtype=ProjectInvitation_FieldMask,name=field_mask,json=fieldMask,proto3" json:"field_mask,omitempty" firestore:"fieldMask"`
	// Previous view index specifies previous position of modified
	// ProjectInvitation. When modification doesn't affect sorted order, value
	// will remain identical to [view_index].
	PreviousViewIndex int32 `protobuf:"varint,4,opt,name=previous_view_index,json=previousViewIndex,proto3" json:"previous_view_index,omitempty" firestore:"previousViewIndex"`
	// Integer specifying ProjectInvitation new index in resulting query view.
	ViewIndex int32 `protobuf:"varint,5,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty" firestore:"viewIndex"`
}

func (m *ProjectInvitationChange_Modified) Reset() {
	*m = ProjectInvitationChange_Modified{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_iam_proto_v1_project_invitation_change_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *ProjectInvitationChange_Modified) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*ProjectInvitationChange_Modified) ProtoMessage() {}

func (m *ProjectInvitationChange_Modified) ProtoReflect() preflect.Message {
	mi := &edgelq_iam_proto_v1_project_invitation_change_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*ProjectInvitationChange_Modified) GotenMessage() {}

// Deprecated, Use ProjectInvitationChange_Modified.ProtoReflect.Descriptor instead.
func (*ProjectInvitationChange_Modified) Descriptor() ([]byte, []int) {
	return edgelq_iam_proto_v1_project_invitation_change_proto_rawDescGZIP(), []int{0, 1}
}

func (m *ProjectInvitationChange_Modified) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *ProjectInvitationChange_Modified) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *ProjectInvitationChange_Modified) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *ProjectInvitationChange_Modified) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *ProjectInvitationChange_Modified) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *ProjectInvitationChange_Modified) GetProjectInvitation() *ProjectInvitation {
	if m != nil {
		return m.ProjectInvitation
	}
	return nil
}

func (m *ProjectInvitationChange_Modified) GetFieldMask() *ProjectInvitation_FieldMask {
	if m != nil {
		return m.FieldMask
	}
	return nil
}

func (m *ProjectInvitationChange_Modified) GetPreviousViewIndex() int32 {
	if m != nil {
		return m.PreviousViewIndex
	}
	return int32(0)
}

func (m *ProjectInvitationChange_Modified) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *ProjectInvitationChange_Modified) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "ProjectInvitationChange_Modified"))
	}
	m.Name = fv
}

func (m *ProjectInvitationChange_Modified) SetProjectInvitation(fv *ProjectInvitation) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ProjectInvitation", "ProjectInvitationChange_Modified"))
	}
	m.ProjectInvitation = fv
}

func (m *ProjectInvitationChange_Modified) SetFieldMask(fv *ProjectInvitation_FieldMask) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "FieldMask", "ProjectInvitationChange_Modified"))
	}
	m.FieldMask = fv
}

func (m *ProjectInvitationChange_Modified) SetPreviousViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "PreviousViewIndex", "ProjectInvitationChange_Modified"))
	}
	m.PreviousViewIndex = fv
}

func (m *ProjectInvitationChange_Modified) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "ProjectInvitationChange_Modified"))
	}
	m.ViewIndex = fv
}

// ProjectInvitation has been added or modified in a query view. Version used
// for stateless watching
type ProjectInvitationChange_Current struct {
	state             protoimpl.MessageState
	sizeCache         protoimpl.SizeCache
	unknownFields     protoimpl.UnknownFields
	ProjectInvitation *ProjectInvitation `protobuf:"bytes,1,opt,name=project_invitation,json=projectInvitation,proto3" json:"project_invitation,omitempty" firestore:"projectInvitation"`
}

func (m *ProjectInvitationChange_Current) Reset() {
	*m = ProjectInvitationChange_Current{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_iam_proto_v1_project_invitation_change_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *ProjectInvitationChange_Current) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*ProjectInvitationChange_Current) ProtoMessage() {}

func (m *ProjectInvitationChange_Current) ProtoReflect() preflect.Message {
	mi := &edgelq_iam_proto_v1_project_invitation_change_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*ProjectInvitationChange_Current) GotenMessage() {}

// Deprecated, Use ProjectInvitationChange_Current.ProtoReflect.Descriptor instead.
func (*ProjectInvitationChange_Current) Descriptor() ([]byte, []int) {
	return edgelq_iam_proto_v1_project_invitation_change_proto_rawDescGZIP(), []int{0, 2}
}

func (m *ProjectInvitationChange_Current) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *ProjectInvitationChange_Current) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *ProjectInvitationChange_Current) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *ProjectInvitationChange_Current) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *ProjectInvitationChange_Current) GetProjectInvitation() *ProjectInvitation {
	if m != nil {
		return m.ProjectInvitation
	}
	return nil
}

func (m *ProjectInvitationChange_Current) SetProjectInvitation(fv *ProjectInvitation) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ProjectInvitation", "ProjectInvitationChange_Current"))
	}
	m.ProjectInvitation = fv
}

// Removed is returned when ProjectInvitation is deleted or leaves Query view
type ProjectInvitationChange_Removed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Name          *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// Integer specifying removed ProjectInvitation index. Not populated in
	// stateless watch type.
	ViewIndex int32 `protobuf:"varint,2,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty" firestore:"viewIndex"`
}

func (m *ProjectInvitationChange_Removed) Reset() {
	*m = ProjectInvitationChange_Removed{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_iam_proto_v1_project_invitation_change_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *ProjectInvitationChange_Removed) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*ProjectInvitationChange_Removed) ProtoMessage() {}

func (m *ProjectInvitationChange_Removed) ProtoReflect() preflect.Message {
	mi := &edgelq_iam_proto_v1_project_invitation_change_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*ProjectInvitationChange_Removed) GotenMessage() {}

// Deprecated, Use ProjectInvitationChange_Removed.ProtoReflect.Descriptor instead.
func (*ProjectInvitationChange_Removed) Descriptor() ([]byte, []int) {
	return edgelq_iam_proto_v1_project_invitation_change_proto_rawDescGZIP(), []int{0, 3}
}

func (m *ProjectInvitationChange_Removed) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *ProjectInvitationChange_Removed) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *ProjectInvitationChange_Removed) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *ProjectInvitationChange_Removed) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *ProjectInvitationChange_Removed) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *ProjectInvitationChange_Removed) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *ProjectInvitationChange_Removed) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "ProjectInvitationChange_Removed"))
	}
	m.Name = fv
}

func (m *ProjectInvitationChange_Removed) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "ProjectInvitationChange_Removed"))
	}
	m.ViewIndex = fv
}

var edgelq_iam_proto_v1_project_invitation_change_proto preflect.FileDescriptor

var edgelq_iam_proto_v1_project_invitation_change_proto_rawDesc = []byte{
	0x0a, 0x33, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x6e,
	0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76,
	0x31, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d,
	0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67,
	0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x2c, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x6e, 0x76,
	0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb9, 0x07,
	0x0a, 0x17, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x41, 0x0a, 0x05, 0x61, 0x64, 0x64,
	0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69,
	0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x76,
	0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x41, 0x64,
	0x64, 0x65, 0x64, 0x48, 0x00, 0x52, 0x05, 0x61, 0x64, 0x64, 0x65, 0x64, 0x12, 0x4a, 0x0a, 0x08,
	0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c,
	0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x48, 0x00, 0x52, 0x08,
	0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x47, 0x0a, 0x07, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x6e, 0x74, 0x74, 0x2e,
	0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e,
	0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x12, 0x47, 0x0a, 0x07, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x48,
	0x00, 0x52, 0x07, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x1a, 0x74, 0x0a, 0x05, 0x41, 0x64,
	0x64, 0x65, 0x64, 0x12, 0x4c, 0x0a, 0x12, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69,
	0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x11,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x76, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x1a, 0xac, 0x02, 0x0a, 0x08, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x2d, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x19, 0xb2, 0xda, 0x21,
	0x15, 0x0a, 0x13, 0x0a, 0x11, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x76, 0x69,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x4c, 0x0a, 0x12,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69,
	0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x76,
	0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x11, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x54, 0x0a, 0x0a, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x42, 0x19, 0xb2, 0xda, 0x21, 0x15,
	0x32, 0x13, 0x0a, 0x11, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b,
	0x12, 0x2e, 0x0a, 0x13, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x5f, 0x76, 0x69, 0x65,
	0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x70,
	0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x56, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x76, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x1a,
	0x57, 0x0a, 0x07, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x4c, 0x0a, 0x12, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x11, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e,
	0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x57, 0x0a, 0x07, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x64, 0x12, 0x2d, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x19, 0xb2, 0xda, 0x21, 0x15, 0x0a, 0x13, 0x0a, 0x11, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x76, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x3a, 0x17, 0x9a, 0xd9, 0x21, 0x13, 0x0a, 0x11, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0d, 0x0a, 0x0b, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x42, 0x8a, 0x01, 0xe8, 0xde, 0x21, 0x00,
	0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x70, 0x62,
	0x2e, 0x76, 0x31, 0x42, 0x1c, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x76, 0x69,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x00, 0x5a, 0x51, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f,
	0x69, 0x61, 0x6d, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x3b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x6e, 0x76, 0x69,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	edgelq_iam_proto_v1_project_invitation_change_proto_rawDescOnce sync.Once
	edgelq_iam_proto_v1_project_invitation_change_proto_rawDescData = edgelq_iam_proto_v1_project_invitation_change_proto_rawDesc
)

func edgelq_iam_proto_v1_project_invitation_change_proto_rawDescGZIP() []byte {
	edgelq_iam_proto_v1_project_invitation_change_proto_rawDescOnce.Do(func() {
		edgelq_iam_proto_v1_project_invitation_change_proto_rawDescData = protoimpl.X.CompressGZIP(edgelq_iam_proto_v1_project_invitation_change_proto_rawDescData)
	})
	return edgelq_iam_proto_v1_project_invitation_change_proto_rawDescData
}

var edgelq_iam_proto_v1_project_invitation_change_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var edgelq_iam_proto_v1_project_invitation_change_proto_goTypes = []interface{}{
	(*ProjectInvitationChange)(nil),          // 0: ntt.iam.v1.ProjectInvitationChange
	(*ProjectInvitationChange_Added)(nil),    // 1: ntt.iam.v1.ProjectInvitationChange.Added
	(*ProjectInvitationChange_Modified)(nil), // 2: ntt.iam.v1.ProjectInvitationChange.Modified
	(*ProjectInvitationChange_Current)(nil),  // 3: ntt.iam.v1.ProjectInvitationChange.Current
	(*ProjectInvitationChange_Removed)(nil),  // 4: ntt.iam.v1.ProjectInvitationChange.Removed
	(*ProjectInvitation)(nil),                // 5: ntt.iam.v1.ProjectInvitation
	(*ProjectInvitation_FieldMask)(nil),      // 6: ntt.iam.v1.ProjectInvitation_FieldMask
}
var edgelq_iam_proto_v1_project_invitation_change_proto_depIdxs = []int32{
	1, // 0: ntt.iam.v1.ProjectInvitationChange.added:type_name -> ntt.iam.v1.ProjectInvitationChange.Added
	2, // 1: ntt.iam.v1.ProjectInvitationChange.modified:type_name -> ntt.iam.v1.ProjectInvitationChange.Modified
	3, // 2: ntt.iam.v1.ProjectInvitationChange.current:type_name -> ntt.iam.v1.ProjectInvitationChange.Current
	4, // 3: ntt.iam.v1.ProjectInvitationChange.removed:type_name -> ntt.iam.v1.ProjectInvitationChange.Removed
	5, // 4: ntt.iam.v1.ProjectInvitationChange.Added.project_invitation:type_name -> ntt.iam.v1.ProjectInvitation
	5, // 5: ntt.iam.v1.ProjectInvitationChange.Modified.project_invitation:type_name -> ntt.iam.v1.ProjectInvitation
	6, // 6: ntt.iam.v1.ProjectInvitationChange.Modified.field_mask:type_name -> ntt.iam.v1.ProjectInvitation_FieldMask
	5, // 7: ntt.iam.v1.ProjectInvitationChange.Current.project_invitation:type_name -> ntt.iam.v1.ProjectInvitation
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { edgelq_iam_proto_v1_project_invitation_change_proto_init() }
func edgelq_iam_proto_v1_project_invitation_change_proto_init() {
	if edgelq_iam_proto_v1_project_invitation_change_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		edgelq_iam_proto_v1_project_invitation_change_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectInvitationChange); i {
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
		edgelq_iam_proto_v1_project_invitation_change_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectInvitationChange_Added); i {
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
		edgelq_iam_proto_v1_project_invitation_change_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectInvitationChange_Modified); i {
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
		edgelq_iam_proto_v1_project_invitation_change_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectInvitationChange_Current); i {
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
		edgelq_iam_proto_v1_project_invitation_change_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectInvitationChange_Removed); i {
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

	edgelq_iam_proto_v1_project_invitation_change_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ProjectInvitationChange_Added_)(nil),
		(*ProjectInvitationChange_Modified_)(nil),
		(*ProjectInvitationChange_Current_)(nil),
		(*ProjectInvitationChange_Removed_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: edgelq_iam_proto_v1_project_invitation_change_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           edgelq_iam_proto_v1_project_invitation_change_proto_goTypes,
		DependencyIndexes: edgelq_iam_proto_v1_project_invitation_change_proto_depIdxs,
		MessageInfos:      edgelq_iam_proto_v1_project_invitation_change_proto_msgTypes,
	}.Build()
	edgelq_iam_proto_v1_project_invitation_change_proto = out.File
	edgelq_iam_proto_v1_project_invitation_change_proto_rawDesc = nil
	edgelq_iam_proto_v1_project_invitation_change_proto_goTypes = nil
	edgelq_iam_proto_v1_project_invitation_change_proto_depIdxs = nil
}
