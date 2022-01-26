// Code generated by protoc-gen-goten-go
// File: edgelq/iam/proto/v1alpha2/condition_change.proto
// DO NOT EDIT!!!

package condition

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
	project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/project"
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

// ConditionChange is used by Watch notifications Responses to describe change
// of single Condition One of Added, Modified, Removed
type ConditionChange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Condition change
	//
	// Types that are valid to be assigned to ChangeType:
	//	*ConditionChange_Added_
	//	*ConditionChange_Modified_
	//	*ConditionChange_Current_
	//	*ConditionChange_Removed_
	ChangeType isConditionChange_ChangeType `protobuf_oneof:"change_type"`
}

func (m *ConditionChange) Reset() {
	*m = ConditionChange{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_iam_proto_v1alpha2_condition_change_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *ConditionChange) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*ConditionChange) ProtoMessage() {}

func (m *ConditionChange) ProtoReflect() preflect.Message {
	mi := &edgelq_iam_proto_v1alpha2_condition_change_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*ConditionChange) GotenMessage() {}

// Deprecated, Use ConditionChange.ProtoReflect.Descriptor instead.
func (*ConditionChange) Descriptor() ([]byte, []int) {
	return edgelq_iam_proto_v1alpha2_condition_change_proto_rawDescGZIP(), []int{0}
}

func (m *ConditionChange) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *ConditionChange) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *ConditionChange) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *ConditionChange) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

type isConditionChange_ChangeType interface {
	isConditionChange_ChangeType()
}

type ConditionChange_Added_ struct {
	// Added is returned when watched document is added, either created or
	// enters Query view
	Added *ConditionChange_Added `protobuf:"bytes,1,opt,name=added,proto3,oneof" firestore:"added"`
}
type ConditionChange_Modified_ struct {
	// Modified is returned when watched document is modified
	Modified *ConditionChange_Modified `protobuf:"bytes,2,opt,name=modified,proto3,oneof" firestore:"modified"`
}
type ConditionChange_Current_ struct {
	// Current is returned in stateless watch when document enters query view or
	// is modified within.
	Current *ConditionChange_Current `protobuf:"bytes,4,opt,name=current,proto3,oneof" firestore:"current"`
}
type ConditionChange_Removed_ struct {
	// Removed is returned when Condition is deleted or leaves Query view
	Removed *ConditionChange_Removed `protobuf:"bytes,3,opt,name=removed,proto3,oneof" firestore:"removed"`
}

func (*ConditionChange_Added_) isConditionChange_ChangeType()    {}
func (*ConditionChange_Modified_) isConditionChange_ChangeType() {}
func (*ConditionChange_Current_) isConditionChange_ChangeType()  {}
func (*ConditionChange_Removed_) isConditionChange_ChangeType()  {}
func (m *ConditionChange) GetChangeType() isConditionChange_ChangeType {
	if m != nil {
		return m.ChangeType
	}
	return nil
}
func (m *ConditionChange) GetAdded() *ConditionChange_Added {
	if x, ok := m.GetChangeType().(*ConditionChange_Added_); ok {
		return x.Added
	}
	return nil
}
func (m *ConditionChange) GetModified() *ConditionChange_Modified {
	if x, ok := m.GetChangeType().(*ConditionChange_Modified_); ok {
		return x.Modified
	}
	return nil
}
func (m *ConditionChange) GetCurrent() *ConditionChange_Current {
	if x, ok := m.GetChangeType().(*ConditionChange_Current_); ok {
		return x.Current
	}
	return nil
}
func (m *ConditionChange) GetRemoved() *ConditionChange_Removed {
	if x, ok := m.GetChangeType().(*ConditionChange_Removed_); ok {
		return x.Removed
	}
	return nil
}
func (m *ConditionChange) SetChangeType(ofv isConditionChange_ChangeType) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "isConditionChange_ChangeType", "ConditionChange"))
	}
	m.ChangeType = ofv
}
func (m *ConditionChange) SetAdded(fv *ConditionChange_Added) {
	m.SetChangeType(&ConditionChange_Added_{Added: fv})
}
func (m *ConditionChange) SetModified(fv *ConditionChange_Modified) {
	m.SetChangeType(&ConditionChange_Modified_{Modified: fv})
}
func (m *ConditionChange) SetCurrent(fv *ConditionChange_Current) {
	m.SetChangeType(&ConditionChange_Current_{Current: fv})
}
func (m *ConditionChange) SetRemoved(fv *ConditionChange_Removed) {
	m.SetChangeType(&ConditionChange_Removed_{Removed: fv})
}

// Condition has been added to query view
type ConditionChange_Added struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Condition     *Condition `protobuf:"bytes,1,opt,name=condition,proto3" json:"condition,omitempty" firestore:"condition"`
	// Integer describing index of added Condition in resulting query view.
	ViewIndex int32 `protobuf:"varint,2,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty" firestore:"viewIndex"`
}

func (m *ConditionChange_Added) Reset() {
	*m = ConditionChange_Added{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_iam_proto_v1alpha2_condition_change_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *ConditionChange_Added) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*ConditionChange_Added) ProtoMessage() {}

func (m *ConditionChange_Added) ProtoReflect() preflect.Message {
	mi := &edgelq_iam_proto_v1alpha2_condition_change_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*ConditionChange_Added) GotenMessage() {}

// Deprecated, Use ConditionChange_Added.ProtoReflect.Descriptor instead.
func (*ConditionChange_Added) Descriptor() ([]byte, []int) {
	return edgelq_iam_proto_v1alpha2_condition_change_proto_rawDescGZIP(), []int{0, 0}
}

func (m *ConditionChange_Added) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *ConditionChange_Added) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *ConditionChange_Added) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *ConditionChange_Added) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *ConditionChange_Added) GetCondition() *Condition {
	if m != nil {
		return m.Condition
	}
	return nil
}

func (m *ConditionChange_Added) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *ConditionChange_Added) SetCondition(fv *Condition) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Condition", "ConditionChange_Added"))
	}
	m.Condition = fv
}

func (m *ConditionChange_Added) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "ConditionChange_Added"))
	}
	m.ViewIndex = fv
}

// Condition changed some of it's fields - contains either full document or
// masked change
type ConditionChange_Modified struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Name of modified Condition
	Name *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// New version of Condition or masked difference, depending on mask_changes
	// instrumentation of issued [WatchConditionRequest] or
	// [WatchConditionsRequest]
	Condition *Condition `protobuf:"bytes,2,opt,name=condition,proto3" json:"condition,omitempty" firestore:"condition"`
	// Used when mask_changes is set, contains field paths of modified
	// properties.
	FieldMask *Condition_FieldMask `protobuf:"bytes,3,opt,customtype=Condition_FieldMask,name=field_mask,json=fieldMask,proto3" json:"field_mask,omitempty" firestore:"fieldMask"`
	// Previous view index specifies previous position of modified Condition.
	// When modification doesn't affect sorted order, value will remain
	// identical to [view_index].
	PreviousViewIndex int32 `protobuf:"varint,4,opt,name=previous_view_index,json=previousViewIndex,proto3" json:"previous_view_index,omitempty" firestore:"previousViewIndex"`
	// Integer specifying Condition new index in resulting query view.
	ViewIndex int32 `protobuf:"varint,5,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty" firestore:"viewIndex"`
}

func (m *ConditionChange_Modified) Reset() {
	*m = ConditionChange_Modified{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_iam_proto_v1alpha2_condition_change_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *ConditionChange_Modified) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*ConditionChange_Modified) ProtoMessage() {}

func (m *ConditionChange_Modified) ProtoReflect() preflect.Message {
	mi := &edgelq_iam_proto_v1alpha2_condition_change_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*ConditionChange_Modified) GotenMessage() {}

// Deprecated, Use ConditionChange_Modified.ProtoReflect.Descriptor instead.
func (*ConditionChange_Modified) Descriptor() ([]byte, []int) {
	return edgelq_iam_proto_v1alpha2_condition_change_proto_rawDescGZIP(), []int{0, 1}
}

func (m *ConditionChange_Modified) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *ConditionChange_Modified) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *ConditionChange_Modified) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *ConditionChange_Modified) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *ConditionChange_Modified) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *ConditionChange_Modified) GetCondition() *Condition {
	if m != nil {
		return m.Condition
	}
	return nil
}

func (m *ConditionChange_Modified) GetFieldMask() *Condition_FieldMask {
	if m != nil {
		return m.FieldMask
	}
	return nil
}

func (m *ConditionChange_Modified) GetPreviousViewIndex() int32 {
	if m != nil {
		return m.PreviousViewIndex
	}
	return int32(0)
}

func (m *ConditionChange_Modified) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *ConditionChange_Modified) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "ConditionChange_Modified"))
	}
	m.Name = fv
}

func (m *ConditionChange_Modified) SetCondition(fv *Condition) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Condition", "ConditionChange_Modified"))
	}
	m.Condition = fv
}

func (m *ConditionChange_Modified) SetFieldMask(fv *Condition_FieldMask) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "FieldMask", "ConditionChange_Modified"))
	}
	m.FieldMask = fv
}

func (m *ConditionChange_Modified) SetPreviousViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "PreviousViewIndex", "ConditionChange_Modified"))
	}
	m.PreviousViewIndex = fv
}

func (m *ConditionChange_Modified) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "ConditionChange_Modified"))
	}
	m.ViewIndex = fv
}

// Condition has been added or modified in a query view. Version used for
// stateless watching
type ConditionChange_Current struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Condition     *Condition `protobuf:"bytes,1,opt,name=condition,proto3" json:"condition,omitempty" firestore:"condition"`
}

func (m *ConditionChange_Current) Reset() {
	*m = ConditionChange_Current{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_iam_proto_v1alpha2_condition_change_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *ConditionChange_Current) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*ConditionChange_Current) ProtoMessage() {}

func (m *ConditionChange_Current) ProtoReflect() preflect.Message {
	mi := &edgelq_iam_proto_v1alpha2_condition_change_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*ConditionChange_Current) GotenMessage() {}

// Deprecated, Use ConditionChange_Current.ProtoReflect.Descriptor instead.
func (*ConditionChange_Current) Descriptor() ([]byte, []int) {
	return edgelq_iam_proto_v1alpha2_condition_change_proto_rawDescGZIP(), []int{0, 2}
}

func (m *ConditionChange_Current) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *ConditionChange_Current) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *ConditionChange_Current) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *ConditionChange_Current) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *ConditionChange_Current) GetCondition() *Condition {
	if m != nil {
		return m.Condition
	}
	return nil
}

func (m *ConditionChange_Current) SetCondition(fv *Condition) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Condition", "ConditionChange_Current"))
	}
	m.Condition = fv
}

// Removed is returned when Condition is deleted or leaves Query view
type ConditionChange_Removed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Name          *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// Integer specifying removed Condition index. Not populated in stateless
	// watch type.
	ViewIndex int32 `protobuf:"varint,2,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty" firestore:"viewIndex"`
}

func (m *ConditionChange_Removed) Reset() {
	*m = ConditionChange_Removed{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_iam_proto_v1alpha2_condition_change_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *ConditionChange_Removed) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*ConditionChange_Removed) ProtoMessage() {}

func (m *ConditionChange_Removed) ProtoReflect() preflect.Message {
	mi := &edgelq_iam_proto_v1alpha2_condition_change_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*ConditionChange_Removed) GotenMessage() {}

// Deprecated, Use ConditionChange_Removed.ProtoReflect.Descriptor instead.
func (*ConditionChange_Removed) Descriptor() ([]byte, []int) {
	return edgelq_iam_proto_v1alpha2_condition_change_proto_rawDescGZIP(), []int{0, 3}
}

func (m *ConditionChange_Removed) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *ConditionChange_Removed) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *ConditionChange_Removed) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *ConditionChange_Removed) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *ConditionChange_Removed) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *ConditionChange_Removed) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *ConditionChange_Removed) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "ConditionChange_Removed"))
	}
	m.Name = fv
}

func (m *ConditionChange_Removed) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "ConditionChange_Removed"))
	}
	m.ViewIndex = fv
}

var edgelq_iam_proto_v1alpha2_condition_change_proto preflect.FileDescriptor

var edgelq_iam_proto_v1alpha2_condition_change_proto_rawDesc = []byte{
	0x0a, 0x30, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x63, 0x6f, 0x6e, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x10, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x32, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x29, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x63, 0x6f,
	0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd0, 0x06,
	0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x12, 0x3f, 0x0a, 0x05, 0x61, 0x64, 0x64, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x27, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x32, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x65, 0x64, 0x48, 0x00, 0x52, 0x05, 0x61, 0x64, 0x64,
	0x65, 0x64, 0x12, 0x48, 0x0a, 0x08, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64,
	0x48, 0x00, 0x52, 0x08, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x45, 0x0a, 0x07,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e,
	0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32,
	0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x07, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x12, 0x45, 0x0a, 0x07, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x48,
	0x00, 0x52, 0x07, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x1a, 0x61, 0x0a, 0x05, 0x41, 0x64,
	0x64, 0x65, 0x64, 0x12, 0x39, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d,
	0x0a, 0x0a, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x76, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x1a, 0x89, 0x02,
	0x0a, 0x08, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x25, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x11, 0xb2, 0xda, 0x21, 0x0d, 0x0a, 0x0b,
	0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x39, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4c, 0x0a, 0x0a,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x42, 0x11, 0xb2, 0xda,
	0x21, 0x0d, 0x32, 0x0b, 0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x2e, 0x0a, 0x13, 0x70, 0x72,
	0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x5f, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75,
	0x73, 0x56, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x69,
	0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x76, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x1a, 0x44, 0x0a, 0x07, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x12, 0x39, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61,
	0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x1a,
	0x4f, 0x0a, 0x07, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x12, 0x25, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x11, 0xb2, 0xda, 0x21, 0x0d, 0x0a, 0x0b,
	0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x76, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x3a, 0x0f, 0x9a, 0xd9, 0x21, 0x0b, 0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x42, 0x0d, 0x0a, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x42, 0x7c, 0xe8, 0xde, 0x21, 0x00, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x74, 0x74, 0x2e,
	0x69, 0x61, 0x6d, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x42,
	0x14, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x00, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67,
	0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x63, 0x6f, 0x6e, 0x64, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x3b, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	edgelq_iam_proto_v1alpha2_condition_change_proto_rawDescOnce sync.Once
	edgelq_iam_proto_v1alpha2_condition_change_proto_rawDescData = edgelq_iam_proto_v1alpha2_condition_change_proto_rawDesc
)

func edgelq_iam_proto_v1alpha2_condition_change_proto_rawDescGZIP() []byte {
	edgelq_iam_proto_v1alpha2_condition_change_proto_rawDescOnce.Do(func() {
		edgelq_iam_proto_v1alpha2_condition_change_proto_rawDescData = protoimpl.X.CompressGZIP(edgelq_iam_proto_v1alpha2_condition_change_proto_rawDescData)
	})
	return edgelq_iam_proto_v1alpha2_condition_change_proto_rawDescData
}

var edgelq_iam_proto_v1alpha2_condition_change_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var edgelq_iam_proto_v1alpha2_condition_change_proto_goTypes = []interface{}{
	(*ConditionChange)(nil),          // 0: ntt.iam.v1alpha2.ConditionChange
	(*ConditionChange_Added)(nil),    // 1: ntt.iam.v1alpha2.ConditionChange.Added
	(*ConditionChange_Modified)(nil), // 2: ntt.iam.v1alpha2.ConditionChange.Modified
	(*ConditionChange_Current)(nil),  // 3: ntt.iam.v1alpha2.ConditionChange.Current
	(*ConditionChange_Removed)(nil),  // 4: ntt.iam.v1alpha2.ConditionChange.Removed
	(*Condition)(nil),                // 5: ntt.iam.v1alpha2.Condition
	(*Condition_FieldMask)(nil),      // 6: ntt.iam.v1alpha2.Condition_FieldMask
}
var edgelq_iam_proto_v1alpha2_condition_change_proto_depIdxs = []int32{
	1, // 0: ntt.iam.v1alpha2.ConditionChange.added:type_name -> ntt.iam.v1alpha2.ConditionChange.Added
	2, // 1: ntt.iam.v1alpha2.ConditionChange.modified:type_name -> ntt.iam.v1alpha2.ConditionChange.Modified
	3, // 2: ntt.iam.v1alpha2.ConditionChange.current:type_name -> ntt.iam.v1alpha2.ConditionChange.Current
	4, // 3: ntt.iam.v1alpha2.ConditionChange.removed:type_name -> ntt.iam.v1alpha2.ConditionChange.Removed
	5, // 4: ntt.iam.v1alpha2.ConditionChange.Added.condition:type_name -> ntt.iam.v1alpha2.Condition
	5, // 5: ntt.iam.v1alpha2.ConditionChange.Modified.condition:type_name -> ntt.iam.v1alpha2.Condition
	6, // 6: ntt.iam.v1alpha2.ConditionChange.Modified.field_mask:type_name -> ntt.iam.v1alpha2.Condition_FieldMask
	5, // 7: ntt.iam.v1alpha2.ConditionChange.Current.condition:type_name -> ntt.iam.v1alpha2.Condition
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { edgelq_iam_proto_v1alpha2_condition_change_proto_init() }
func edgelq_iam_proto_v1alpha2_condition_change_proto_init() {
	if edgelq_iam_proto_v1alpha2_condition_change_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		edgelq_iam_proto_v1alpha2_condition_change_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConditionChange); i {
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
		edgelq_iam_proto_v1alpha2_condition_change_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConditionChange_Added); i {
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
		edgelq_iam_proto_v1alpha2_condition_change_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConditionChange_Modified); i {
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
		edgelq_iam_proto_v1alpha2_condition_change_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConditionChange_Current); i {
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
		edgelq_iam_proto_v1alpha2_condition_change_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConditionChange_Removed); i {
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

	edgelq_iam_proto_v1alpha2_condition_change_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ConditionChange_Added_)(nil),
		(*ConditionChange_Modified_)(nil),
		(*ConditionChange_Current_)(nil),
		(*ConditionChange_Removed_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: edgelq_iam_proto_v1alpha2_condition_change_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           edgelq_iam_proto_v1alpha2_condition_change_proto_goTypes,
		DependencyIndexes: edgelq_iam_proto_v1alpha2_condition_change_proto_depIdxs,
		MessageInfos:      edgelq_iam_proto_v1alpha2_condition_change_proto_msgTypes,
	}.Build()
	edgelq_iam_proto_v1alpha2_condition_change_proto = out.File
	edgelq_iam_proto_v1alpha2_condition_change_proto_rawDesc = nil
	edgelq_iam_proto_v1alpha2_condition_change_proto_goTypes = nil
	edgelq_iam_proto_v1alpha2_condition_change_proto_depIdxs = nil
}
