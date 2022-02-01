// Code generated by protoc-gen-goten-go
// File: edgelq/iam/proto/v1alpha/group_change.proto
// DO NOT EDIT!!!

package group

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
	organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/organization"
	project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/project"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
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
	_ = &project.Project{}
	_ = &field_mask.FieldMask{}
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// GroupChange is used by Watch notifications Responses to describe change of
// single Group One of Added, Modified, Removed
type GroupChange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Group change
	//
	// Types that are valid to be assigned to ChangeType:
	//	*GroupChange_Added_
	//	*GroupChange_Modified_
	//	*GroupChange_Current_
	//	*GroupChange_Removed_
	ChangeType isGroupChange_ChangeType `protobuf_oneof:"change_type"`
}

func (m *GroupChange) Reset() {
	*m = GroupChange{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_iam_proto_v1alpha_group_change_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *GroupChange) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*GroupChange) ProtoMessage() {}

func (m *GroupChange) ProtoReflect() preflect.Message {
	mi := &edgelq_iam_proto_v1alpha_group_change_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*GroupChange) GotenMessage() {}

// Deprecated, Use GroupChange.ProtoReflect.Descriptor instead.
func (*GroupChange) Descriptor() ([]byte, []int) {
	return edgelq_iam_proto_v1alpha_group_change_proto_rawDescGZIP(), []int{0}
}

func (m *GroupChange) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *GroupChange) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *GroupChange) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *GroupChange) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

type isGroupChange_ChangeType interface {
	isGroupChange_ChangeType()
}

type GroupChange_Added_ struct {
	// Added is returned when watched document is added, either created or
	// enters Query view
	Added *GroupChange_Added `protobuf:"bytes,1,opt,name=added,proto3,oneof" firestore:"added"`
}
type GroupChange_Modified_ struct {
	// Modified is returned when watched document is modified
	Modified *GroupChange_Modified `protobuf:"bytes,2,opt,name=modified,proto3,oneof" firestore:"modified"`
}
type GroupChange_Current_ struct {
	// Current is returned in stateless watch when document enters query view or
	// is modified within.
	Current *GroupChange_Current `protobuf:"bytes,4,opt,name=current,proto3,oneof" firestore:"current"`
}
type GroupChange_Removed_ struct {
	// Removed is returned when Group is deleted or leaves Query view
	Removed *GroupChange_Removed `protobuf:"bytes,3,opt,name=removed,proto3,oneof" firestore:"removed"`
}

func (*GroupChange_Added_) isGroupChange_ChangeType()    {}
func (*GroupChange_Modified_) isGroupChange_ChangeType() {}
func (*GroupChange_Current_) isGroupChange_ChangeType()  {}
func (*GroupChange_Removed_) isGroupChange_ChangeType()  {}
func (m *GroupChange) GetChangeType() isGroupChange_ChangeType {
	if m != nil {
		return m.ChangeType
	}
	return nil
}
func (m *GroupChange) GetAdded() *GroupChange_Added {
	if x, ok := m.GetChangeType().(*GroupChange_Added_); ok {
		return x.Added
	}
	return nil
}
func (m *GroupChange) GetModified() *GroupChange_Modified {
	if x, ok := m.GetChangeType().(*GroupChange_Modified_); ok {
		return x.Modified
	}
	return nil
}
func (m *GroupChange) GetCurrent() *GroupChange_Current {
	if x, ok := m.GetChangeType().(*GroupChange_Current_); ok {
		return x.Current
	}
	return nil
}
func (m *GroupChange) GetRemoved() *GroupChange_Removed {
	if x, ok := m.GetChangeType().(*GroupChange_Removed_); ok {
		return x.Removed
	}
	return nil
}
func (m *GroupChange) SetChangeType(ofv isGroupChange_ChangeType) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "isGroupChange_ChangeType", "GroupChange"))
	}
	m.ChangeType = ofv
}
func (m *GroupChange) SetAdded(fv *GroupChange_Added) {
	m.SetChangeType(&GroupChange_Added_{Added: fv})
}
func (m *GroupChange) SetModified(fv *GroupChange_Modified) {
	m.SetChangeType(&GroupChange_Modified_{Modified: fv})
}
func (m *GroupChange) SetCurrent(fv *GroupChange_Current) {
	m.SetChangeType(&GroupChange_Current_{Current: fv})
}
func (m *GroupChange) SetRemoved(fv *GroupChange_Removed) {
	m.SetChangeType(&GroupChange_Removed_{Removed: fv})
}

// Group has been added to query view
type GroupChange_Added struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Group         *Group `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty" firestore:"group"`
	// Integer describing index of added Group in resulting query view.
	ViewIndex int32 `protobuf:"varint,2,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty" firestore:"viewIndex"`
}

func (m *GroupChange_Added) Reset() {
	*m = GroupChange_Added{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_iam_proto_v1alpha_group_change_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *GroupChange_Added) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*GroupChange_Added) ProtoMessage() {}

func (m *GroupChange_Added) ProtoReflect() preflect.Message {
	mi := &edgelq_iam_proto_v1alpha_group_change_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*GroupChange_Added) GotenMessage() {}

// Deprecated, Use GroupChange_Added.ProtoReflect.Descriptor instead.
func (*GroupChange_Added) Descriptor() ([]byte, []int) {
	return edgelq_iam_proto_v1alpha_group_change_proto_rawDescGZIP(), []int{0, 0}
}

func (m *GroupChange_Added) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *GroupChange_Added) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *GroupChange_Added) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *GroupChange_Added) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *GroupChange_Added) GetGroup() *Group {
	if m != nil {
		return m.Group
	}
	return nil
}

func (m *GroupChange_Added) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *GroupChange_Added) SetGroup(fv *Group) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Group", "GroupChange_Added"))
	}
	m.Group = fv
}

func (m *GroupChange_Added) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "GroupChange_Added"))
	}
	m.ViewIndex = fv
}

// Group changed some of it's fields - contains either full document or masked
// change
type GroupChange_Modified struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Name of modified Group
	Name *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// New version of Group or masked difference, depending on mask_changes
	// instrumentation of issued [WatchGroupRequest] or [WatchGroupsRequest]
	Group *Group `protobuf:"bytes,2,opt,name=group,proto3" json:"group,omitempty" firestore:"group"`
	// Used when mask_changes is set, contains field paths of modified
	// properties.
	FieldMask *Group_FieldMask `protobuf:"bytes,3,opt,customtype=Group_FieldMask,name=field_mask,json=fieldMask,proto3" json:"field_mask,omitempty" firestore:"fieldMask"`
	// Previous view index specifies previous position of modified Group.
	// When modification doesn't affect sorted order, value will remain
	// identical to [view_index].
	PreviousViewIndex int32 `protobuf:"varint,4,opt,name=previous_view_index,json=previousViewIndex,proto3" json:"previous_view_index,omitempty" firestore:"previousViewIndex"`
	// Integer specifying Group new index in resulting query view.
	ViewIndex int32 `protobuf:"varint,5,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty" firestore:"viewIndex"`
}

func (m *GroupChange_Modified) Reset() {
	*m = GroupChange_Modified{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_iam_proto_v1alpha_group_change_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *GroupChange_Modified) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*GroupChange_Modified) ProtoMessage() {}

func (m *GroupChange_Modified) ProtoReflect() preflect.Message {
	mi := &edgelq_iam_proto_v1alpha_group_change_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*GroupChange_Modified) GotenMessage() {}

// Deprecated, Use GroupChange_Modified.ProtoReflect.Descriptor instead.
func (*GroupChange_Modified) Descriptor() ([]byte, []int) {
	return edgelq_iam_proto_v1alpha_group_change_proto_rawDescGZIP(), []int{0, 1}
}

func (m *GroupChange_Modified) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *GroupChange_Modified) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *GroupChange_Modified) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *GroupChange_Modified) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *GroupChange_Modified) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *GroupChange_Modified) GetGroup() *Group {
	if m != nil {
		return m.Group
	}
	return nil
}

func (m *GroupChange_Modified) GetFieldMask() *Group_FieldMask {
	if m != nil {
		return m.FieldMask
	}
	return nil
}

func (m *GroupChange_Modified) GetPreviousViewIndex() int32 {
	if m != nil {
		return m.PreviousViewIndex
	}
	return int32(0)
}

func (m *GroupChange_Modified) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *GroupChange_Modified) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "GroupChange_Modified"))
	}
	m.Name = fv
}

func (m *GroupChange_Modified) SetGroup(fv *Group) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Group", "GroupChange_Modified"))
	}
	m.Group = fv
}

func (m *GroupChange_Modified) SetFieldMask(fv *Group_FieldMask) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "FieldMask", "GroupChange_Modified"))
	}
	m.FieldMask = fv
}

func (m *GroupChange_Modified) SetPreviousViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "PreviousViewIndex", "GroupChange_Modified"))
	}
	m.PreviousViewIndex = fv
}

func (m *GroupChange_Modified) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "GroupChange_Modified"))
	}
	m.ViewIndex = fv
}

// Group has been added or modified in a query view. Version used for
// stateless watching
type GroupChange_Current struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Group         *Group `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty" firestore:"group"`
}

func (m *GroupChange_Current) Reset() {
	*m = GroupChange_Current{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_iam_proto_v1alpha_group_change_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *GroupChange_Current) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*GroupChange_Current) ProtoMessage() {}

func (m *GroupChange_Current) ProtoReflect() preflect.Message {
	mi := &edgelq_iam_proto_v1alpha_group_change_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*GroupChange_Current) GotenMessage() {}

// Deprecated, Use GroupChange_Current.ProtoReflect.Descriptor instead.
func (*GroupChange_Current) Descriptor() ([]byte, []int) {
	return edgelq_iam_proto_v1alpha_group_change_proto_rawDescGZIP(), []int{0, 2}
}

func (m *GroupChange_Current) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *GroupChange_Current) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *GroupChange_Current) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *GroupChange_Current) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *GroupChange_Current) GetGroup() *Group {
	if m != nil {
		return m.Group
	}
	return nil
}

func (m *GroupChange_Current) SetGroup(fv *Group) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Group", "GroupChange_Current"))
	}
	m.Group = fv
}

// Removed is returned when Group is deleted or leaves Query view
type GroupChange_Removed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Name          *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// Integer specifying removed Group index. Not populated in stateless watch
	// type.
	ViewIndex int32 `protobuf:"varint,2,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty" firestore:"viewIndex"`
}

func (m *GroupChange_Removed) Reset() {
	*m = GroupChange_Removed{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_iam_proto_v1alpha_group_change_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *GroupChange_Removed) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*GroupChange_Removed) ProtoMessage() {}

func (m *GroupChange_Removed) ProtoReflect() preflect.Message {
	mi := &edgelq_iam_proto_v1alpha_group_change_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*GroupChange_Removed) GotenMessage() {}

// Deprecated, Use GroupChange_Removed.ProtoReflect.Descriptor instead.
func (*GroupChange_Removed) Descriptor() ([]byte, []int) {
	return edgelq_iam_proto_v1alpha_group_change_proto_rawDescGZIP(), []int{0, 3}
}

func (m *GroupChange_Removed) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *GroupChange_Removed) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *GroupChange_Removed) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *GroupChange_Removed) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *GroupChange_Removed) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *GroupChange_Removed) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *GroupChange_Removed) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "GroupChange_Removed"))
	}
	m.Name = fv
}

func (m *GroupChange_Removed) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "GroupChange_Removed"))
	}
	m.ViewIndex = fv
}

var edgelq_iam_proto_v1alpha_group_change_proto preflect.FileDescriptor

var edgelq_iam_proto_v1alpha_group_change_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x6e,
	0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x19,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74,
	0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67,
	0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65,
	0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x67, 0x6f,
	0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x24, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x92, 0x06, 0x0a, 0x0b, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x3a, 0x0a, 0x05, 0x61, 0x64, 0x64, 0x65, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x65, 0x64, 0x48, 0x00, 0x52, 0x05, 0x61, 0x64, 0x64,
	0x65, 0x64, 0x12, 0x43, 0x0a, 0x08, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x48, 0x00, 0x52, 0x08, 0x6d,
	0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x40, 0x0a, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69,
	0x61, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x48, 0x00,
	0x52, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x40, 0x0a, 0x07, 0x72, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6e, 0x74, 0x74,
	0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64,
	0x48, 0x00, 0x52, 0x07, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x1a, 0x54, 0x0a, 0x05, 0x41,
	0x64, 0x64, 0x65, 0x64, 0x12, 0x2c, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x05, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x76, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x1a, 0xf4, 0x01, 0x0a, 0x08, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x21,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0d, 0xb2, 0xda,
	0x21, 0x09, 0x0a, 0x07, 0x0a, 0x05, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x2c, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12,
	0x48, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x42,
	0x0d, 0xb2, 0xda, 0x21, 0x09, 0x32, 0x07, 0x0a, 0x05, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x09,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x2e, 0x0a, 0x13, 0x70, 0x72, 0x65,
	0x76, 0x69, 0x6f, 0x75, 0x73, 0x5f, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73,
	0x56, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x69, 0x65,
	0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x76,
	0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x1a, 0x37, 0x0a, 0x07, 0x43, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x12, 0x2c, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x1a, 0x4b, 0x0a, 0x07, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x12, 0x21, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0d, 0xb2, 0xda, 0x21, 0x09,
	0x0a, 0x07, 0x0a, 0x05, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x76, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x3a, 0x1c,
	0x9a, 0xd9, 0x21, 0x07, 0x0a, 0x05, 0x47, 0x72, 0x6f, 0x75, 0x70, 0xfa, 0xde, 0x21, 0x0d, 0x0a,
	0x0b, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x42, 0x0d, 0x0a, 0x0b,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x42, 0xbe, 0x01, 0xe8, 0xde,
	0x21, 0x00, 0x92, 0x8c, 0xd1, 0x02, 0x4b, 0x0a, 0x10, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65,
	0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e,
	0x70, 0x62, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x42, 0x10, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x00, 0x5a, 0x3c,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x3b, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	edgelq_iam_proto_v1alpha_group_change_proto_rawDescOnce sync.Once
	edgelq_iam_proto_v1alpha_group_change_proto_rawDescData = edgelq_iam_proto_v1alpha_group_change_proto_rawDesc
)

func edgelq_iam_proto_v1alpha_group_change_proto_rawDescGZIP() []byte {
	edgelq_iam_proto_v1alpha_group_change_proto_rawDescOnce.Do(func() {
		edgelq_iam_proto_v1alpha_group_change_proto_rawDescData = protoimpl.X.CompressGZIP(edgelq_iam_proto_v1alpha_group_change_proto_rawDescData)
	})
	return edgelq_iam_proto_v1alpha_group_change_proto_rawDescData
}

var edgelq_iam_proto_v1alpha_group_change_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var edgelq_iam_proto_v1alpha_group_change_proto_goTypes = []interface{}{
	(*GroupChange)(nil),          // 0: ntt.iam.v1alpha.GroupChange
	(*GroupChange_Added)(nil),    // 1: ntt.iam.v1alpha.GroupChange.Added
	(*GroupChange_Modified)(nil), // 2: ntt.iam.v1alpha.GroupChange.Modified
	(*GroupChange_Current)(nil),  // 3: ntt.iam.v1alpha.GroupChange.Current
	(*GroupChange_Removed)(nil),  // 4: ntt.iam.v1alpha.GroupChange.Removed
	(*Group)(nil),                // 5: ntt.iam.v1alpha.Group
	(*Group_FieldMask)(nil),      // 6: ntt.iam.v1alpha.Group_FieldMask
}
var edgelq_iam_proto_v1alpha_group_change_proto_depIdxs = []int32{
	1, // 0: ntt.iam.v1alpha.GroupChange.added:type_name -> ntt.iam.v1alpha.GroupChange.Added
	2, // 1: ntt.iam.v1alpha.GroupChange.modified:type_name -> ntt.iam.v1alpha.GroupChange.Modified
	3, // 2: ntt.iam.v1alpha.GroupChange.current:type_name -> ntt.iam.v1alpha.GroupChange.Current
	4, // 3: ntt.iam.v1alpha.GroupChange.removed:type_name -> ntt.iam.v1alpha.GroupChange.Removed
	5, // 4: ntt.iam.v1alpha.GroupChange.Added.group:type_name -> ntt.iam.v1alpha.Group
	5, // 5: ntt.iam.v1alpha.GroupChange.Modified.group:type_name -> ntt.iam.v1alpha.Group
	6, // 6: ntt.iam.v1alpha.GroupChange.Modified.field_mask:type_name -> ntt.iam.v1alpha.Group_FieldMask
	5, // 7: ntt.iam.v1alpha.GroupChange.Current.group:type_name -> ntt.iam.v1alpha.Group
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { edgelq_iam_proto_v1alpha_group_change_proto_init() }
func edgelq_iam_proto_v1alpha_group_change_proto_init() {
	if edgelq_iam_proto_v1alpha_group_change_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		edgelq_iam_proto_v1alpha_group_change_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupChange); i {
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
		edgelq_iam_proto_v1alpha_group_change_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupChange_Added); i {
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
		edgelq_iam_proto_v1alpha_group_change_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupChange_Modified); i {
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
		edgelq_iam_proto_v1alpha_group_change_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupChange_Current); i {
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
		edgelq_iam_proto_v1alpha_group_change_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupChange_Removed); i {
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

	edgelq_iam_proto_v1alpha_group_change_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*GroupChange_Added_)(nil),
		(*GroupChange_Modified_)(nil),
		(*GroupChange_Current_)(nil),
		(*GroupChange_Removed_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: edgelq_iam_proto_v1alpha_group_change_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           edgelq_iam_proto_v1alpha_group_change_proto_goTypes,
		DependencyIndexes: edgelq_iam_proto_v1alpha_group_change_proto_depIdxs,
		MessageInfos:      edgelq_iam_proto_v1alpha_group_change_proto_msgTypes,
	}.Build()
	edgelq_iam_proto_v1alpha_group_change_proto = out.File
	edgelq_iam_proto_v1alpha_group_change_proto_rawDesc = nil
	edgelq_iam_proto_v1alpha_group_change_proto_goTypes = nil
	edgelq_iam_proto_v1alpha_group_change_proto_depIdxs = nil
}
