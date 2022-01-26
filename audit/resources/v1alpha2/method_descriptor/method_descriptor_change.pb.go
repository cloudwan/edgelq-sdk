// Code generated by protoc-gen-goten-go
// File: edgelq/audit/proto/v1alpha2/method_descriptor_change.proto
// DO NOT EDIT!!!

package method_descriptor

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
	_ = &field_mask.FieldMask{}
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// MethodDescriptorChange is used by Watch notifications Responses to describe
// change of single MethodDescriptor One of Added, Modified, Removed
type MethodDescriptorChange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// MethodDescriptor change
	//
	// Types that are valid to be assigned to ChangeType:
	//	*MethodDescriptorChange_Added_
	//	*MethodDescriptorChange_Modified_
	//	*MethodDescriptorChange_Current_
	//	*MethodDescriptorChange_Removed_
	ChangeType isMethodDescriptorChange_ChangeType `protobuf_oneof:"change_type"`
}

func (m *MethodDescriptorChange) Reset() {
	*m = MethodDescriptorChange{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_audit_proto_v1alpha2_method_descriptor_change_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *MethodDescriptorChange) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*MethodDescriptorChange) ProtoMessage() {}

func (m *MethodDescriptorChange) ProtoReflect() preflect.Message {
	mi := &edgelq_audit_proto_v1alpha2_method_descriptor_change_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*MethodDescriptorChange) GotenMessage() {}

// Deprecated, Use MethodDescriptorChange.ProtoReflect.Descriptor instead.
func (*MethodDescriptorChange) Descriptor() ([]byte, []int) {
	return edgelq_audit_proto_v1alpha2_method_descriptor_change_proto_rawDescGZIP(), []int{0}
}

func (m *MethodDescriptorChange) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *MethodDescriptorChange) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *MethodDescriptorChange) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *MethodDescriptorChange) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

type isMethodDescriptorChange_ChangeType interface {
	isMethodDescriptorChange_ChangeType()
}

type MethodDescriptorChange_Added_ struct {
	// Added is returned when watched document is added, either created or
	// enters Query view
	Added *MethodDescriptorChange_Added `protobuf:"bytes,1,opt,name=added,proto3,oneof" firestore:"added"`
}
type MethodDescriptorChange_Modified_ struct {
	// Modified is returned when watched document is modified
	Modified *MethodDescriptorChange_Modified `protobuf:"bytes,2,opt,name=modified,proto3,oneof" firestore:"modified"`
}
type MethodDescriptorChange_Current_ struct {
	// Current is returned in stateless watch when document enters query view or
	// is modified within.
	Current *MethodDescriptorChange_Current `protobuf:"bytes,4,opt,name=current,proto3,oneof" firestore:"current"`
}
type MethodDescriptorChange_Removed_ struct {
	// Removed is returned when MethodDescriptor is deleted or leaves Query view
	Removed *MethodDescriptorChange_Removed `protobuf:"bytes,3,opt,name=removed,proto3,oneof" firestore:"removed"`
}

func (*MethodDescriptorChange_Added_) isMethodDescriptorChange_ChangeType()    {}
func (*MethodDescriptorChange_Modified_) isMethodDescriptorChange_ChangeType() {}
func (*MethodDescriptorChange_Current_) isMethodDescriptorChange_ChangeType()  {}
func (*MethodDescriptorChange_Removed_) isMethodDescriptorChange_ChangeType()  {}
func (m *MethodDescriptorChange) GetChangeType() isMethodDescriptorChange_ChangeType {
	if m != nil {
		return m.ChangeType
	}
	return nil
}
func (m *MethodDescriptorChange) GetAdded() *MethodDescriptorChange_Added {
	if x, ok := m.GetChangeType().(*MethodDescriptorChange_Added_); ok {
		return x.Added
	}
	return nil
}
func (m *MethodDescriptorChange) GetModified() *MethodDescriptorChange_Modified {
	if x, ok := m.GetChangeType().(*MethodDescriptorChange_Modified_); ok {
		return x.Modified
	}
	return nil
}
func (m *MethodDescriptorChange) GetCurrent() *MethodDescriptorChange_Current {
	if x, ok := m.GetChangeType().(*MethodDescriptorChange_Current_); ok {
		return x.Current
	}
	return nil
}
func (m *MethodDescriptorChange) GetRemoved() *MethodDescriptorChange_Removed {
	if x, ok := m.GetChangeType().(*MethodDescriptorChange_Removed_); ok {
		return x.Removed
	}
	return nil
}
func (m *MethodDescriptorChange) SetChangeType(ofv isMethodDescriptorChange_ChangeType) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "isMethodDescriptorChange_ChangeType", "MethodDescriptorChange"))
	}
	m.ChangeType = ofv
}
func (m *MethodDescriptorChange) SetAdded(fv *MethodDescriptorChange_Added) {
	m.SetChangeType(&MethodDescriptorChange_Added_{Added: fv})
}
func (m *MethodDescriptorChange) SetModified(fv *MethodDescriptorChange_Modified) {
	m.SetChangeType(&MethodDescriptorChange_Modified_{Modified: fv})
}
func (m *MethodDescriptorChange) SetCurrent(fv *MethodDescriptorChange_Current) {
	m.SetChangeType(&MethodDescriptorChange_Current_{Current: fv})
}
func (m *MethodDescriptorChange) SetRemoved(fv *MethodDescriptorChange_Removed) {
	m.SetChangeType(&MethodDescriptorChange_Removed_{Removed: fv})
}

// MethodDescriptor has been added to query view
type MethodDescriptorChange_Added struct {
	state            protoimpl.MessageState
	sizeCache        protoimpl.SizeCache
	unknownFields    protoimpl.UnknownFields
	MethodDescriptor *MethodDescriptor `protobuf:"bytes,1,opt,name=method_descriptor,json=methodDescriptor,proto3" json:"method_descriptor,omitempty" firestore:"methodDescriptor"`
	// Integer describing index of added MethodDescriptor in resulting query
	// view.
	ViewIndex int32 `protobuf:"varint,2,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty" firestore:"viewIndex"`
}

func (m *MethodDescriptorChange_Added) Reset() {
	*m = MethodDescriptorChange_Added{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_audit_proto_v1alpha2_method_descriptor_change_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *MethodDescriptorChange_Added) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*MethodDescriptorChange_Added) ProtoMessage() {}

func (m *MethodDescriptorChange_Added) ProtoReflect() preflect.Message {
	mi := &edgelq_audit_proto_v1alpha2_method_descriptor_change_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*MethodDescriptorChange_Added) GotenMessage() {}

// Deprecated, Use MethodDescriptorChange_Added.ProtoReflect.Descriptor instead.
func (*MethodDescriptorChange_Added) Descriptor() ([]byte, []int) {
	return edgelq_audit_proto_v1alpha2_method_descriptor_change_proto_rawDescGZIP(), []int{0, 0}
}

func (m *MethodDescriptorChange_Added) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *MethodDescriptorChange_Added) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *MethodDescriptorChange_Added) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *MethodDescriptorChange_Added) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *MethodDescriptorChange_Added) GetMethodDescriptor() *MethodDescriptor {
	if m != nil {
		return m.MethodDescriptor
	}
	return nil
}

func (m *MethodDescriptorChange_Added) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *MethodDescriptorChange_Added) SetMethodDescriptor(fv *MethodDescriptor) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "MethodDescriptor", "MethodDescriptorChange_Added"))
	}
	m.MethodDescriptor = fv
}

func (m *MethodDescriptorChange_Added) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "MethodDescriptorChange_Added"))
	}
	m.ViewIndex = fv
}

// MethodDescriptor changed some of it's fields - contains either full
// document or masked change
type MethodDescriptorChange_Modified struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Name of modified MethodDescriptor
	Name *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// New version of MethodDescriptor or masked difference, depending on
	// mask_changes instrumentation of issued [WatchMethodDescriptorRequest] or
	// [WatchMethodDescriptorsRequest]
	MethodDescriptor *MethodDescriptor `protobuf:"bytes,2,opt,name=method_descriptor,json=methodDescriptor,proto3" json:"method_descriptor,omitempty" firestore:"methodDescriptor"`
	// Used when mask_changes is set, contains field paths of modified
	// properties.
	FieldMask *MethodDescriptor_FieldMask `protobuf:"bytes,3,opt,customtype=MethodDescriptor_FieldMask,name=field_mask,json=fieldMask,proto3" json:"field_mask,omitempty" firestore:"fieldMask"`
	// Previous view index specifies previous position of modified
	// MethodDescriptor. When modification doesn't affect sorted order, value
	// will remain identical to [view_index].
	PreviousViewIndex int32 `protobuf:"varint,4,opt,name=previous_view_index,json=previousViewIndex,proto3" json:"previous_view_index,omitempty" firestore:"previousViewIndex"`
	// Integer specifying MethodDescriptor new index in resulting query view.
	ViewIndex int32 `protobuf:"varint,5,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty" firestore:"viewIndex"`
}

func (m *MethodDescriptorChange_Modified) Reset() {
	*m = MethodDescriptorChange_Modified{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_audit_proto_v1alpha2_method_descriptor_change_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *MethodDescriptorChange_Modified) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*MethodDescriptorChange_Modified) ProtoMessage() {}

func (m *MethodDescriptorChange_Modified) ProtoReflect() preflect.Message {
	mi := &edgelq_audit_proto_v1alpha2_method_descriptor_change_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*MethodDescriptorChange_Modified) GotenMessage() {}

// Deprecated, Use MethodDescriptorChange_Modified.ProtoReflect.Descriptor instead.
func (*MethodDescriptorChange_Modified) Descriptor() ([]byte, []int) {
	return edgelq_audit_proto_v1alpha2_method_descriptor_change_proto_rawDescGZIP(), []int{0, 1}
}

func (m *MethodDescriptorChange_Modified) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *MethodDescriptorChange_Modified) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *MethodDescriptorChange_Modified) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *MethodDescriptorChange_Modified) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *MethodDescriptorChange_Modified) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *MethodDescriptorChange_Modified) GetMethodDescriptor() *MethodDescriptor {
	if m != nil {
		return m.MethodDescriptor
	}
	return nil
}

func (m *MethodDescriptorChange_Modified) GetFieldMask() *MethodDescriptor_FieldMask {
	if m != nil {
		return m.FieldMask
	}
	return nil
}

func (m *MethodDescriptorChange_Modified) GetPreviousViewIndex() int32 {
	if m != nil {
		return m.PreviousViewIndex
	}
	return int32(0)
}

func (m *MethodDescriptorChange_Modified) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *MethodDescriptorChange_Modified) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "MethodDescriptorChange_Modified"))
	}
	m.Name = fv
}

func (m *MethodDescriptorChange_Modified) SetMethodDescriptor(fv *MethodDescriptor) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "MethodDescriptor", "MethodDescriptorChange_Modified"))
	}
	m.MethodDescriptor = fv
}

func (m *MethodDescriptorChange_Modified) SetFieldMask(fv *MethodDescriptor_FieldMask) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "FieldMask", "MethodDescriptorChange_Modified"))
	}
	m.FieldMask = fv
}

func (m *MethodDescriptorChange_Modified) SetPreviousViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "PreviousViewIndex", "MethodDescriptorChange_Modified"))
	}
	m.PreviousViewIndex = fv
}

func (m *MethodDescriptorChange_Modified) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "MethodDescriptorChange_Modified"))
	}
	m.ViewIndex = fv
}

// MethodDescriptor has been added or modified in a query view. Version used
// for stateless watching
type MethodDescriptorChange_Current struct {
	state            protoimpl.MessageState
	sizeCache        protoimpl.SizeCache
	unknownFields    protoimpl.UnknownFields
	MethodDescriptor *MethodDescriptor `protobuf:"bytes,1,opt,name=method_descriptor,json=methodDescriptor,proto3" json:"method_descriptor,omitempty" firestore:"methodDescriptor"`
}

func (m *MethodDescriptorChange_Current) Reset() {
	*m = MethodDescriptorChange_Current{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_audit_proto_v1alpha2_method_descriptor_change_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *MethodDescriptorChange_Current) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*MethodDescriptorChange_Current) ProtoMessage() {}

func (m *MethodDescriptorChange_Current) ProtoReflect() preflect.Message {
	mi := &edgelq_audit_proto_v1alpha2_method_descriptor_change_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*MethodDescriptorChange_Current) GotenMessage() {}

// Deprecated, Use MethodDescriptorChange_Current.ProtoReflect.Descriptor instead.
func (*MethodDescriptorChange_Current) Descriptor() ([]byte, []int) {
	return edgelq_audit_proto_v1alpha2_method_descriptor_change_proto_rawDescGZIP(), []int{0, 2}
}

func (m *MethodDescriptorChange_Current) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *MethodDescriptorChange_Current) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *MethodDescriptorChange_Current) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *MethodDescriptorChange_Current) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *MethodDescriptorChange_Current) GetMethodDescriptor() *MethodDescriptor {
	if m != nil {
		return m.MethodDescriptor
	}
	return nil
}

func (m *MethodDescriptorChange_Current) SetMethodDescriptor(fv *MethodDescriptor) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "MethodDescriptor", "MethodDescriptorChange_Current"))
	}
	m.MethodDescriptor = fv
}

// Removed is returned when MethodDescriptor is deleted or leaves Query view
type MethodDescriptorChange_Removed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Name          *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// Integer specifying removed MethodDescriptor index. Not populated in
	// stateless watch type.
	ViewIndex int32 `protobuf:"varint,2,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty" firestore:"viewIndex"`
}

func (m *MethodDescriptorChange_Removed) Reset() {
	*m = MethodDescriptorChange_Removed{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_audit_proto_v1alpha2_method_descriptor_change_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *MethodDescriptorChange_Removed) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*MethodDescriptorChange_Removed) ProtoMessage() {}

func (m *MethodDescriptorChange_Removed) ProtoReflect() preflect.Message {
	mi := &edgelq_audit_proto_v1alpha2_method_descriptor_change_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*MethodDescriptorChange_Removed) GotenMessage() {}

// Deprecated, Use MethodDescriptorChange_Removed.ProtoReflect.Descriptor instead.
func (*MethodDescriptorChange_Removed) Descriptor() ([]byte, []int) {
	return edgelq_audit_proto_v1alpha2_method_descriptor_change_proto_rawDescGZIP(), []int{0, 3}
}

func (m *MethodDescriptorChange_Removed) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *MethodDescriptorChange_Removed) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *MethodDescriptorChange_Removed) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *MethodDescriptorChange_Removed) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *MethodDescriptorChange_Removed) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *MethodDescriptorChange_Removed) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *MethodDescriptorChange_Removed) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "MethodDescriptorChange_Removed"))
	}
	m.Name = fv
}

func (m *MethodDescriptorChange_Removed) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "MethodDescriptorChange_Removed"))
	}
	m.ViewIndex = fv
}

var edgelq_audit_proto_v1alpha2_method_descriptor_change_proto preflect.FileDescriptor

var edgelq_audit_proto_v1alpha2_method_descriptor_change_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x6d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x5f,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x6e, 0x74,
	0x74, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32,
	0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67,
	0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f,
	0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x33,
	0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x6d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xdf, 0x07, 0x0a, 0x16, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x48,
	0x0a, 0x05, 0x61, 0x64, 0x64, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e,
	0x6e, 0x74, 0x74, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x32, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x6f, 0x72, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x65, 0x64, 0x48,
	0x00, 0x52, 0x05, 0x61, 0x64, 0x64, 0x65, 0x64, 0x12, 0x51, 0x0a, 0x08, 0x6d, 0x6f, 0x64, 0x69,
	0x66, 0x69, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x6e, 0x74, 0x74,
	0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x48,
	0x00, 0x52, 0x08, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x4e, 0x0a, 0x07, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x6e,
	0x74, 0x74, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x32, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x6f, 0x72, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x48, 0x00, 0x52, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x4e, 0x0a, 0x07, 0x72,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x6e,
	0x74, 0x74, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x32, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x6f, 0x72, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64,
	0x48, 0x00, 0x52, 0x07, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x1a, 0x79, 0x0a, 0x05, 0x41,
	0x64, 0x64, 0x65, 0x64, 0x12, 0x51, 0x0a, 0x11, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x24, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x32, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x10, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x69, 0x65, 0x77, 0x5f,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x76, 0x69, 0x65,
	0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x1a, 0xaf, 0x02, 0x0a, 0x08, 0x4d, 0x6f, 0x64, 0x69, 0x66,
	0x69, 0x65, 0x64, 0x12, 0x2c, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x18, 0xb2, 0xda, 0x21, 0x14, 0x0a, 0x12, 0x0a, 0x10, 0x4d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x51, 0x0a, 0x11, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6e,
	0x74, 0x74, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x32, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x6f, 0x72, 0x52, 0x10, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x6f, 0x72, 0x12, 0x53, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61,
	0x73, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x4d, 0x61, 0x73, 0x6b, 0x42, 0x18, 0xb2, 0xda, 0x21, 0x14, 0x32, 0x12, 0x0a, 0x10, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x09,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x2e, 0x0a, 0x13, 0x70, 0x72, 0x65,
	0x76, 0x69, 0x6f, 0x75, 0x73, 0x5f, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73,
	0x56, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x69, 0x65,
	0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x76,
	0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x1a, 0x5c, 0x0a, 0x07, 0x43, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x12, 0x51, 0x0a, 0x11, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24,
	0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x32, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x6f, 0x72, 0x52, 0x10, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x1a, 0x56, 0x0a, 0x07, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x64, 0x12, 0x2c, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x18, 0xb2, 0xda, 0x21, 0x14, 0x0a, 0x12, 0x0a, 0x10, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x76, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x3a, 0x16,
	0x9a, 0xd9, 0x21, 0x12, 0x0a, 0x10, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x42, 0x0d, 0x0a, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x42, 0x97, 0x01, 0xe8, 0xde, 0x21, 0x00, 0x0a, 0x19, 0x63, 0x6f,
	0x6d, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x70, 0x62, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x42, 0x1b, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x00, 0x5a, 0x57, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65,
	0x6c, 0x71, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x6d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x3b, 0x6d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	edgelq_audit_proto_v1alpha2_method_descriptor_change_proto_rawDescOnce sync.Once
	edgelq_audit_proto_v1alpha2_method_descriptor_change_proto_rawDescData = edgelq_audit_proto_v1alpha2_method_descriptor_change_proto_rawDesc
)

func edgelq_audit_proto_v1alpha2_method_descriptor_change_proto_rawDescGZIP() []byte {
	edgelq_audit_proto_v1alpha2_method_descriptor_change_proto_rawDescOnce.Do(func() {
		edgelq_audit_proto_v1alpha2_method_descriptor_change_proto_rawDescData = protoimpl.X.CompressGZIP(edgelq_audit_proto_v1alpha2_method_descriptor_change_proto_rawDescData)
	})
	return edgelq_audit_proto_v1alpha2_method_descriptor_change_proto_rawDescData
}

var edgelq_audit_proto_v1alpha2_method_descriptor_change_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var edgelq_audit_proto_v1alpha2_method_descriptor_change_proto_goTypes = []interface{}{
	(*MethodDescriptorChange)(nil),          // 0: ntt.audit.v1alpha2.MethodDescriptorChange
	(*MethodDescriptorChange_Added)(nil),    // 1: ntt.audit.v1alpha2.MethodDescriptorChange.Added
	(*MethodDescriptorChange_Modified)(nil), // 2: ntt.audit.v1alpha2.MethodDescriptorChange.Modified
	(*MethodDescriptorChange_Current)(nil),  // 3: ntt.audit.v1alpha2.MethodDescriptorChange.Current
	(*MethodDescriptorChange_Removed)(nil),  // 4: ntt.audit.v1alpha2.MethodDescriptorChange.Removed
	(*MethodDescriptor)(nil),                // 5: ntt.audit.v1alpha2.MethodDescriptor
	(*MethodDescriptor_FieldMask)(nil),      // 6: ntt.audit.v1alpha2.MethodDescriptor_FieldMask
}
var edgelq_audit_proto_v1alpha2_method_descriptor_change_proto_depIdxs = []int32{
	1, // 0: ntt.audit.v1alpha2.MethodDescriptorChange.added:type_name -> ntt.audit.v1alpha2.MethodDescriptorChange.Added
	2, // 1: ntt.audit.v1alpha2.MethodDescriptorChange.modified:type_name -> ntt.audit.v1alpha2.MethodDescriptorChange.Modified
	3, // 2: ntt.audit.v1alpha2.MethodDescriptorChange.current:type_name -> ntt.audit.v1alpha2.MethodDescriptorChange.Current
	4, // 3: ntt.audit.v1alpha2.MethodDescriptorChange.removed:type_name -> ntt.audit.v1alpha2.MethodDescriptorChange.Removed
	5, // 4: ntt.audit.v1alpha2.MethodDescriptorChange.Added.method_descriptor:type_name -> ntt.audit.v1alpha2.MethodDescriptor
	5, // 5: ntt.audit.v1alpha2.MethodDescriptorChange.Modified.method_descriptor:type_name -> ntt.audit.v1alpha2.MethodDescriptor
	6, // 6: ntt.audit.v1alpha2.MethodDescriptorChange.Modified.field_mask:type_name -> ntt.audit.v1alpha2.MethodDescriptor_FieldMask
	5, // 7: ntt.audit.v1alpha2.MethodDescriptorChange.Current.method_descriptor:type_name -> ntt.audit.v1alpha2.MethodDescriptor
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { edgelq_audit_proto_v1alpha2_method_descriptor_change_proto_init() }
func edgelq_audit_proto_v1alpha2_method_descriptor_change_proto_init() {
	if edgelq_audit_proto_v1alpha2_method_descriptor_change_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		edgelq_audit_proto_v1alpha2_method_descriptor_change_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MethodDescriptorChange); i {
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
		edgelq_audit_proto_v1alpha2_method_descriptor_change_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MethodDescriptorChange_Added); i {
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
		edgelq_audit_proto_v1alpha2_method_descriptor_change_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MethodDescriptorChange_Modified); i {
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
		edgelq_audit_proto_v1alpha2_method_descriptor_change_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MethodDescriptorChange_Current); i {
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
		edgelq_audit_proto_v1alpha2_method_descriptor_change_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MethodDescriptorChange_Removed); i {
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

	edgelq_audit_proto_v1alpha2_method_descriptor_change_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*MethodDescriptorChange_Added_)(nil),
		(*MethodDescriptorChange_Modified_)(nil),
		(*MethodDescriptorChange_Current_)(nil),
		(*MethodDescriptorChange_Removed_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: edgelq_audit_proto_v1alpha2_method_descriptor_change_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           edgelq_audit_proto_v1alpha2_method_descriptor_change_proto_goTypes,
		DependencyIndexes: edgelq_audit_proto_v1alpha2_method_descriptor_change_proto_depIdxs,
		MessageInfos:      edgelq_audit_proto_v1alpha2_method_descriptor_change_proto_msgTypes,
	}.Build()
	edgelq_audit_proto_v1alpha2_method_descriptor_change_proto = out.File
	edgelq_audit_proto_v1alpha2_method_descriptor_change_proto_rawDesc = nil
	edgelq_audit_proto_v1alpha2_method_descriptor_change_proto_goTypes = nil
	edgelq_audit_proto_v1alpha2_method_descriptor_change_proto_depIdxs = nil
}
