// Code generated by protoc-gen-goten-go
// File: edgelq/audit/proto/v1alpha/audited_resource_descriptor_change.proto
// DO NOT EDIT!!!

package audited_resource_descriptor

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

// AuditedResourceDescriptorChange is used by Watch notifications Responses to
// describe change of single AuditedResourceDescriptor One of Added, Modified,
// Removed
type AuditedResourceDescriptorChange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// AuditedResourceDescriptor change
	//
	// Types that are valid to be assigned to ChangeType:
	//	*AuditedResourceDescriptorChange_Added_
	//	*AuditedResourceDescriptorChange_Modified_
	//	*AuditedResourceDescriptorChange_Current_
	//	*AuditedResourceDescriptorChange_Removed_
	ChangeType isAuditedResourceDescriptorChange_ChangeType `protobuf_oneof:"change_type"`
}

func (m *AuditedResourceDescriptorChange) Reset() {
	*m = AuditedResourceDescriptorChange{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_audit_proto_v1alpha_audited_resource_descriptor_change_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *AuditedResourceDescriptorChange) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*AuditedResourceDescriptorChange) ProtoMessage() {}

func (m *AuditedResourceDescriptorChange) ProtoReflect() preflect.Message {
	mi := &edgelq_audit_proto_v1alpha_audited_resource_descriptor_change_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*AuditedResourceDescriptorChange) GotenMessage() {}

// Deprecated, Use AuditedResourceDescriptorChange.ProtoReflect.Descriptor instead.
func (*AuditedResourceDescriptorChange) Descriptor() ([]byte, []int) {
	return edgelq_audit_proto_v1alpha_audited_resource_descriptor_change_proto_rawDescGZIP(), []int{0}
}

func (m *AuditedResourceDescriptorChange) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *AuditedResourceDescriptorChange) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *AuditedResourceDescriptorChange) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *AuditedResourceDescriptorChange) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

type isAuditedResourceDescriptorChange_ChangeType interface {
	isAuditedResourceDescriptorChange_ChangeType()
}

type AuditedResourceDescriptorChange_Added_ struct {
	// Added is returned when watched document is added, either created or
	// enters Query view
	Added *AuditedResourceDescriptorChange_Added `protobuf:"bytes,1,opt,name=added,proto3,oneof" firestore:"added"`
}
type AuditedResourceDescriptorChange_Modified_ struct {
	// Modified is returned when watched document is modified
	Modified *AuditedResourceDescriptorChange_Modified `protobuf:"bytes,2,opt,name=modified,proto3,oneof" firestore:"modified"`
}
type AuditedResourceDescriptorChange_Current_ struct {
	// Current is returned in stateless watch when document enters query view or
	// is modified within.
	Current *AuditedResourceDescriptorChange_Current `protobuf:"bytes,4,opt,name=current,proto3,oneof" firestore:"current"`
}
type AuditedResourceDescriptorChange_Removed_ struct {
	// Removed is returned when AuditedResourceDescriptor is deleted or leaves
	// Query view
	Removed *AuditedResourceDescriptorChange_Removed `protobuf:"bytes,3,opt,name=removed,proto3,oneof" firestore:"removed"`
}

func (*AuditedResourceDescriptorChange_Added_) isAuditedResourceDescriptorChange_ChangeType()    {}
func (*AuditedResourceDescriptorChange_Modified_) isAuditedResourceDescriptorChange_ChangeType() {}
func (*AuditedResourceDescriptorChange_Current_) isAuditedResourceDescriptorChange_ChangeType()  {}
func (*AuditedResourceDescriptorChange_Removed_) isAuditedResourceDescriptorChange_ChangeType()  {}
func (m *AuditedResourceDescriptorChange) GetChangeType() isAuditedResourceDescriptorChange_ChangeType {
	if m != nil {
		return m.ChangeType
	}
	return nil
}
func (m *AuditedResourceDescriptorChange) GetAdded() *AuditedResourceDescriptorChange_Added {
	if x, ok := m.GetChangeType().(*AuditedResourceDescriptorChange_Added_); ok {
		return x.Added
	}
	return nil
}
func (m *AuditedResourceDescriptorChange) GetModified() *AuditedResourceDescriptorChange_Modified {
	if x, ok := m.GetChangeType().(*AuditedResourceDescriptorChange_Modified_); ok {
		return x.Modified
	}
	return nil
}
func (m *AuditedResourceDescriptorChange) GetCurrent() *AuditedResourceDescriptorChange_Current {
	if x, ok := m.GetChangeType().(*AuditedResourceDescriptorChange_Current_); ok {
		return x.Current
	}
	return nil
}
func (m *AuditedResourceDescriptorChange) GetRemoved() *AuditedResourceDescriptorChange_Removed {
	if x, ok := m.GetChangeType().(*AuditedResourceDescriptorChange_Removed_); ok {
		return x.Removed
	}
	return nil
}
func (m *AuditedResourceDescriptorChange) SetChangeType(ofv isAuditedResourceDescriptorChange_ChangeType) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "isAuditedResourceDescriptorChange_ChangeType", "AuditedResourceDescriptorChange"))
	}
	m.ChangeType = ofv
}
func (m *AuditedResourceDescriptorChange) SetAdded(fv *AuditedResourceDescriptorChange_Added) {
	m.SetChangeType(&AuditedResourceDescriptorChange_Added_{Added: fv})
}
func (m *AuditedResourceDescriptorChange) SetModified(fv *AuditedResourceDescriptorChange_Modified) {
	m.SetChangeType(&AuditedResourceDescriptorChange_Modified_{Modified: fv})
}
func (m *AuditedResourceDescriptorChange) SetCurrent(fv *AuditedResourceDescriptorChange_Current) {
	m.SetChangeType(&AuditedResourceDescriptorChange_Current_{Current: fv})
}
func (m *AuditedResourceDescriptorChange) SetRemoved(fv *AuditedResourceDescriptorChange_Removed) {
	m.SetChangeType(&AuditedResourceDescriptorChange_Removed_{Removed: fv})
}

// AuditedResourceDescriptor has been added to query view
type AuditedResourceDescriptorChange_Added struct {
	state                     protoimpl.MessageState
	sizeCache                 protoimpl.SizeCache
	unknownFields             protoimpl.UnknownFields
	AuditedResourceDescriptor *AuditedResourceDescriptor `protobuf:"bytes,1,opt,name=audited_resource_descriptor,json=auditedResourceDescriptor,proto3" json:"audited_resource_descriptor,omitempty" firestore:"auditedResourceDescriptor"`
	// Integer describing index of added AuditedResourceDescriptor in resulting
	// query view.
	ViewIndex int32 `protobuf:"varint,2,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty" firestore:"viewIndex"`
}

func (m *AuditedResourceDescriptorChange_Added) Reset() {
	*m = AuditedResourceDescriptorChange_Added{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_audit_proto_v1alpha_audited_resource_descriptor_change_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *AuditedResourceDescriptorChange_Added) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*AuditedResourceDescriptorChange_Added) ProtoMessage() {}

func (m *AuditedResourceDescriptorChange_Added) ProtoReflect() preflect.Message {
	mi := &edgelq_audit_proto_v1alpha_audited_resource_descriptor_change_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*AuditedResourceDescriptorChange_Added) GotenMessage() {}

// Deprecated, Use AuditedResourceDescriptorChange_Added.ProtoReflect.Descriptor instead.
func (*AuditedResourceDescriptorChange_Added) Descriptor() ([]byte, []int) {
	return edgelq_audit_proto_v1alpha_audited_resource_descriptor_change_proto_rawDescGZIP(), []int{0, 0}
}

func (m *AuditedResourceDescriptorChange_Added) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *AuditedResourceDescriptorChange_Added) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *AuditedResourceDescriptorChange_Added) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *AuditedResourceDescriptorChange_Added) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *AuditedResourceDescriptorChange_Added) GetAuditedResourceDescriptor() *AuditedResourceDescriptor {
	if m != nil {
		return m.AuditedResourceDescriptor
	}
	return nil
}

func (m *AuditedResourceDescriptorChange_Added) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *AuditedResourceDescriptorChange_Added) SetAuditedResourceDescriptor(fv *AuditedResourceDescriptor) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "AuditedResourceDescriptor", "AuditedResourceDescriptorChange_Added"))
	}
	m.AuditedResourceDescriptor = fv
}

func (m *AuditedResourceDescriptorChange_Added) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "AuditedResourceDescriptorChange_Added"))
	}
	m.ViewIndex = fv
}

// AuditedResourceDescriptor changed some of it's fields - contains either
// full document or masked change
type AuditedResourceDescriptorChange_Modified struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Name of modified AuditedResourceDescriptor
	Name *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// New version of AuditedResourceDescriptor or masked difference, depending
	// on mask_changes instrumentation of issued
	// [WatchAuditedResourceDescriptorRequest] or
	// [WatchAuditedResourceDescriptorsRequest]
	AuditedResourceDescriptor *AuditedResourceDescriptor `protobuf:"bytes,2,opt,name=audited_resource_descriptor,json=auditedResourceDescriptor,proto3" json:"audited_resource_descriptor,omitempty" firestore:"auditedResourceDescriptor"`
	// Used when mask_changes is set, contains field paths of modified
	// properties.
	FieldMask *AuditedResourceDescriptor_FieldMask `protobuf:"bytes,3,opt,customtype=AuditedResourceDescriptor_FieldMask,name=field_mask,json=fieldMask,proto3" json:"field_mask,omitempty" firestore:"fieldMask"`
	// Previous view index specifies previous position of modified
	// AuditedResourceDescriptor. When modification doesn't affect sorted order,
	// value will remain identical to [view_index].
	PreviousViewIndex int32 `protobuf:"varint,4,opt,name=previous_view_index,json=previousViewIndex,proto3" json:"previous_view_index,omitempty" firestore:"previousViewIndex"`
	// Integer specifying AuditedResourceDescriptor new index in resulting query
	// view.
	ViewIndex int32 `protobuf:"varint,5,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty" firestore:"viewIndex"`
}

func (m *AuditedResourceDescriptorChange_Modified) Reset() {
	*m = AuditedResourceDescriptorChange_Modified{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_audit_proto_v1alpha_audited_resource_descriptor_change_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *AuditedResourceDescriptorChange_Modified) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*AuditedResourceDescriptorChange_Modified) ProtoMessage() {}

func (m *AuditedResourceDescriptorChange_Modified) ProtoReflect() preflect.Message {
	mi := &edgelq_audit_proto_v1alpha_audited_resource_descriptor_change_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*AuditedResourceDescriptorChange_Modified) GotenMessage() {}

// Deprecated, Use AuditedResourceDescriptorChange_Modified.ProtoReflect.Descriptor instead.
func (*AuditedResourceDescriptorChange_Modified) Descriptor() ([]byte, []int) {
	return edgelq_audit_proto_v1alpha_audited_resource_descriptor_change_proto_rawDescGZIP(), []int{0, 1}
}

func (m *AuditedResourceDescriptorChange_Modified) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *AuditedResourceDescriptorChange_Modified) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *AuditedResourceDescriptorChange_Modified) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *AuditedResourceDescriptorChange_Modified) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *AuditedResourceDescriptorChange_Modified) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *AuditedResourceDescriptorChange_Modified) GetAuditedResourceDescriptor() *AuditedResourceDescriptor {
	if m != nil {
		return m.AuditedResourceDescriptor
	}
	return nil
}

func (m *AuditedResourceDescriptorChange_Modified) GetFieldMask() *AuditedResourceDescriptor_FieldMask {
	if m != nil {
		return m.FieldMask
	}
	return nil
}

func (m *AuditedResourceDescriptorChange_Modified) GetPreviousViewIndex() int32 {
	if m != nil {
		return m.PreviousViewIndex
	}
	return int32(0)
}

func (m *AuditedResourceDescriptorChange_Modified) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *AuditedResourceDescriptorChange_Modified) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "AuditedResourceDescriptorChange_Modified"))
	}
	m.Name = fv
}

func (m *AuditedResourceDescriptorChange_Modified) SetAuditedResourceDescriptor(fv *AuditedResourceDescriptor) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "AuditedResourceDescriptor", "AuditedResourceDescriptorChange_Modified"))
	}
	m.AuditedResourceDescriptor = fv
}

func (m *AuditedResourceDescriptorChange_Modified) SetFieldMask(fv *AuditedResourceDescriptor_FieldMask) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "FieldMask", "AuditedResourceDescriptorChange_Modified"))
	}
	m.FieldMask = fv
}

func (m *AuditedResourceDescriptorChange_Modified) SetPreviousViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "PreviousViewIndex", "AuditedResourceDescriptorChange_Modified"))
	}
	m.PreviousViewIndex = fv
}

func (m *AuditedResourceDescriptorChange_Modified) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "AuditedResourceDescriptorChange_Modified"))
	}
	m.ViewIndex = fv
}

// AuditedResourceDescriptor has been added or modified in a query view.
// Version used for stateless watching
type AuditedResourceDescriptorChange_Current struct {
	state                     protoimpl.MessageState
	sizeCache                 protoimpl.SizeCache
	unknownFields             protoimpl.UnknownFields
	AuditedResourceDescriptor *AuditedResourceDescriptor `protobuf:"bytes,1,opt,name=audited_resource_descriptor,json=auditedResourceDescriptor,proto3" json:"audited_resource_descriptor,omitempty" firestore:"auditedResourceDescriptor"`
}

func (m *AuditedResourceDescriptorChange_Current) Reset() {
	*m = AuditedResourceDescriptorChange_Current{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_audit_proto_v1alpha_audited_resource_descriptor_change_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *AuditedResourceDescriptorChange_Current) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*AuditedResourceDescriptorChange_Current) ProtoMessage() {}

func (m *AuditedResourceDescriptorChange_Current) ProtoReflect() preflect.Message {
	mi := &edgelq_audit_proto_v1alpha_audited_resource_descriptor_change_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*AuditedResourceDescriptorChange_Current) GotenMessage() {}

// Deprecated, Use AuditedResourceDescriptorChange_Current.ProtoReflect.Descriptor instead.
func (*AuditedResourceDescriptorChange_Current) Descriptor() ([]byte, []int) {
	return edgelq_audit_proto_v1alpha_audited_resource_descriptor_change_proto_rawDescGZIP(), []int{0, 2}
}

func (m *AuditedResourceDescriptorChange_Current) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *AuditedResourceDescriptorChange_Current) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *AuditedResourceDescriptorChange_Current) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *AuditedResourceDescriptorChange_Current) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *AuditedResourceDescriptorChange_Current) GetAuditedResourceDescriptor() *AuditedResourceDescriptor {
	if m != nil {
		return m.AuditedResourceDescriptor
	}
	return nil
}

func (m *AuditedResourceDescriptorChange_Current) SetAuditedResourceDescriptor(fv *AuditedResourceDescriptor) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "AuditedResourceDescriptor", "AuditedResourceDescriptorChange_Current"))
	}
	m.AuditedResourceDescriptor = fv
}

// Removed is returned when AuditedResourceDescriptor is deleted or leaves
// Query view
type AuditedResourceDescriptorChange_Removed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Name          *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// Integer specifying removed AuditedResourceDescriptor index. Not populated
	// in stateless watch type.
	ViewIndex int32 `protobuf:"varint,2,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty" firestore:"viewIndex"`
}

func (m *AuditedResourceDescriptorChange_Removed) Reset() {
	*m = AuditedResourceDescriptorChange_Removed{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_audit_proto_v1alpha_audited_resource_descriptor_change_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *AuditedResourceDescriptorChange_Removed) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*AuditedResourceDescriptorChange_Removed) ProtoMessage() {}

func (m *AuditedResourceDescriptorChange_Removed) ProtoReflect() preflect.Message {
	mi := &edgelq_audit_proto_v1alpha_audited_resource_descriptor_change_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*AuditedResourceDescriptorChange_Removed) GotenMessage() {}

// Deprecated, Use AuditedResourceDescriptorChange_Removed.ProtoReflect.Descriptor instead.
func (*AuditedResourceDescriptorChange_Removed) Descriptor() ([]byte, []int) {
	return edgelq_audit_proto_v1alpha_audited_resource_descriptor_change_proto_rawDescGZIP(), []int{0, 3}
}

func (m *AuditedResourceDescriptorChange_Removed) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *AuditedResourceDescriptorChange_Removed) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *AuditedResourceDescriptorChange_Removed) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *AuditedResourceDescriptorChange_Removed) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *AuditedResourceDescriptorChange_Removed) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *AuditedResourceDescriptorChange_Removed) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *AuditedResourceDescriptorChange_Removed) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "AuditedResourceDescriptorChange_Removed"))
	}
	m.Name = fv
}

func (m *AuditedResourceDescriptorChange_Removed) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "AuditedResourceDescriptorChange_Removed"))
	}
	m.ViewIndex = fv
}

var edgelq_audit_proto_v1alpha_audited_resource_descriptor_change_proto preflect.FileDescriptor

var edgelq_audit_proto_v1alpha_audited_resource_descriptor_change_proto_rawDesc = []byte{
	0x0a, 0x43, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x61, 0x75, 0x64,
	0x69, 0x74, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x6e, 0x74, 0x74, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3c, 0x65, 0x64, 0x67, 0x65,
	0x6c, 0x71, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x65, 0x64, 0x5f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa3, 0x09, 0x0a, 0x1f, 0x41, 0x75, 0x64,
	0x69, 0x74, 0x65, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x50, 0x0a, 0x05,
	0x61, 0x64, 0x64, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x6e, 0x74,
	0x74, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e,
	0x41, 0x75, 0x64, 0x69, 0x74, 0x65, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e,
	0x41, 0x64, 0x64, 0x65, 0x64, 0x48, 0x00, 0x52, 0x05, 0x61, 0x64, 0x64, 0x65, 0x64, 0x12, 0x59,
	0x0a, 0x08, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x3b, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x74, 0x65, 0x64, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x48, 0x00, 0x52,
	0x08, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x56, 0x0a, 0x07, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x6e, 0x74, 0x74,
	0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x41,
	0x75, 0x64, 0x69, 0x74, 0x65, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x12, 0x56, 0x0a, 0x07, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x74, 0x65, 0x64, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x48, 0x00,
	0x52, 0x07, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x1a, 0x94, 0x01, 0x0a, 0x05, 0x41, 0x64,
	0x64, 0x65, 0x64, 0x12, 0x6c, 0x0a, 0x1b, 0x61, 0x75, 0x64, 0x69, 0x74, 0x65, 0x64, 0x5f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x61,
	0x75, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x41, 0x75, 0x64,
	0x69, 0x74, 0x65, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x19, 0x61, 0x75, 0x64, 0x69, 0x74, 0x65, 0x64, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f,
	0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x76, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x1a, 0xdc, 0x02, 0x0a, 0x08, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x35, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x21, 0xb2, 0xda, 0x21,
	0x1d, 0x0a, 0x1b, 0x0a, 0x19, 0x41, 0x75, 0x64, 0x69, 0x74, 0x65, 0x64, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x6c, 0x0a, 0x1b, 0x61, 0x75, 0x64, 0x69, 0x74, 0x65, 0x64, 0x5f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x6e, 0x74, 0x74, 0x2e,
	0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x41, 0x75,
	0x64, 0x69, 0x74, 0x65, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x19, 0x61, 0x75, 0x64, 0x69, 0x74, 0x65, 0x64,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x6f, 0x72, 0x12, 0x5c, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61,
	0x73, 0x6b, 0x42, 0x21, 0xb2, 0xda, 0x21, 0x1d, 0x32, 0x1b, 0x0a, 0x19, 0x41, 0x75, 0x64, 0x69,
	0x74, 0x65, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b,
	0x12, 0x2e, 0x0a, 0x13, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x5f, 0x76, 0x69, 0x65,
	0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x70,
	0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x56, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x76, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x1a,
	0x77, 0x0a, 0x07, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x6c, 0x0a, 0x1b, 0x61, 0x75,
	0x64, 0x69, 0x74, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2c, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x74, 0x65, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x19, 0x61,
	0x75, 0x64, 0x69, 0x74, 0x65, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x1a, 0x5f, 0x0a, 0x07, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x64, 0x12, 0x35, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x21, 0xb2, 0xda, 0x21, 0x1d, 0x0a, 0x1b, 0x0a, 0x19, 0x41, 0x75, 0x64, 0x69, 0x74,
	0x65, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x6f, 0x72, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x69,
	0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x76, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x3a, 0x44, 0x9a, 0xd9, 0x21, 0x1b, 0x0a,
	0x19, 0x41, 0x75, 0x64, 0x69, 0x74, 0x65, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0xfa, 0xde, 0x21, 0x21, 0x0a, 0x1f,
	0x41, 0x75, 0x64, 0x69, 0x74, 0x65, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x42,
	0x0d, 0x0a, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x42, 0xb0,
	0x02, 0xe8, 0xde, 0x21, 0x00, 0x92, 0x8c, 0xd1, 0x02, 0x79, 0x0a, 0x26, 0x61, 0x75, 0x64, 0x69,
	0x74, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69,
	0x6e, 0x67, 0x12, 0x4f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x61,
	0x75, 0x64, 0x69, 0x74, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x65, 0x64, 0x5f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x6f, 0x72, 0x0a, 0x18, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x61, 0x75, 0x64,
	0x69, 0x74, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x42, 0x24, 0x41,
	0x75, 0x64, 0x69, 0x74, 0x65, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x00, 0x5a, 0x6a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c,
	0x71, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x65,
	0x64, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x6f, 0x72, 0x3b, 0x61, 0x75, 0x64, 0x69, 0x74, 0x65, 0x64, 0x5f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f,
	0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	edgelq_audit_proto_v1alpha_audited_resource_descriptor_change_proto_rawDescOnce sync.Once
	edgelq_audit_proto_v1alpha_audited_resource_descriptor_change_proto_rawDescData = edgelq_audit_proto_v1alpha_audited_resource_descriptor_change_proto_rawDesc
)

func edgelq_audit_proto_v1alpha_audited_resource_descriptor_change_proto_rawDescGZIP() []byte {
	edgelq_audit_proto_v1alpha_audited_resource_descriptor_change_proto_rawDescOnce.Do(func() {
		edgelq_audit_proto_v1alpha_audited_resource_descriptor_change_proto_rawDescData = protoimpl.X.CompressGZIP(edgelq_audit_proto_v1alpha_audited_resource_descriptor_change_proto_rawDescData)
	})
	return edgelq_audit_proto_v1alpha_audited_resource_descriptor_change_proto_rawDescData
}

var edgelq_audit_proto_v1alpha_audited_resource_descriptor_change_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var edgelq_audit_proto_v1alpha_audited_resource_descriptor_change_proto_goTypes = []interface{}{
	(*AuditedResourceDescriptorChange)(nil),          // 0: ntt.audit.v1alpha.AuditedResourceDescriptorChange
	(*AuditedResourceDescriptorChange_Added)(nil),    // 1: ntt.audit.v1alpha.AuditedResourceDescriptorChange.Added
	(*AuditedResourceDescriptorChange_Modified)(nil), // 2: ntt.audit.v1alpha.AuditedResourceDescriptorChange.Modified
	(*AuditedResourceDescriptorChange_Current)(nil),  // 3: ntt.audit.v1alpha.AuditedResourceDescriptorChange.Current
	(*AuditedResourceDescriptorChange_Removed)(nil),  // 4: ntt.audit.v1alpha.AuditedResourceDescriptorChange.Removed
	(*AuditedResourceDescriptor)(nil),                // 5: ntt.audit.v1alpha.AuditedResourceDescriptor
	(*AuditedResourceDescriptor_FieldMask)(nil),      // 6: ntt.audit.v1alpha.AuditedResourceDescriptor_FieldMask
}
var edgelq_audit_proto_v1alpha_audited_resource_descriptor_change_proto_depIdxs = []int32{
	1, // 0: ntt.audit.v1alpha.AuditedResourceDescriptorChange.added:type_name -> ntt.audit.v1alpha.AuditedResourceDescriptorChange.Added
	2, // 1: ntt.audit.v1alpha.AuditedResourceDescriptorChange.modified:type_name -> ntt.audit.v1alpha.AuditedResourceDescriptorChange.Modified
	3, // 2: ntt.audit.v1alpha.AuditedResourceDescriptorChange.current:type_name -> ntt.audit.v1alpha.AuditedResourceDescriptorChange.Current
	4, // 3: ntt.audit.v1alpha.AuditedResourceDescriptorChange.removed:type_name -> ntt.audit.v1alpha.AuditedResourceDescriptorChange.Removed
	5, // 4: ntt.audit.v1alpha.AuditedResourceDescriptorChange.Added.audited_resource_descriptor:type_name -> ntt.audit.v1alpha.AuditedResourceDescriptor
	5, // 5: ntt.audit.v1alpha.AuditedResourceDescriptorChange.Modified.audited_resource_descriptor:type_name -> ntt.audit.v1alpha.AuditedResourceDescriptor
	6, // 6: ntt.audit.v1alpha.AuditedResourceDescriptorChange.Modified.field_mask:type_name -> ntt.audit.v1alpha.AuditedResourceDescriptor_FieldMask
	5, // 7: ntt.audit.v1alpha.AuditedResourceDescriptorChange.Current.audited_resource_descriptor:type_name -> ntt.audit.v1alpha.AuditedResourceDescriptor
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { edgelq_audit_proto_v1alpha_audited_resource_descriptor_change_proto_init() }
func edgelq_audit_proto_v1alpha_audited_resource_descriptor_change_proto_init() {
	if edgelq_audit_proto_v1alpha_audited_resource_descriptor_change_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		edgelq_audit_proto_v1alpha_audited_resource_descriptor_change_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuditedResourceDescriptorChange); i {
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
		edgelq_audit_proto_v1alpha_audited_resource_descriptor_change_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuditedResourceDescriptorChange_Added); i {
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
		edgelq_audit_proto_v1alpha_audited_resource_descriptor_change_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuditedResourceDescriptorChange_Modified); i {
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
		edgelq_audit_proto_v1alpha_audited_resource_descriptor_change_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuditedResourceDescriptorChange_Current); i {
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
		edgelq_audit_proto_v1alpha_audited_resource_descriptor_change_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuditedResourceDescriptorChange_Removed); i {
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

	edgelq_audit_proto_v1alpha_audited_resource_descriptor_change_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*AuditedResourceDescriptorChange_Added_)(nil),
		(*AuditedResourceDescriptorChange_Modified_)(nil),
		(*AuditedResourceDescriptorChange_Current_)(nil),
		(*AuditedResourceDescriptorChange_Removed_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: edgelq_audit_proto_v1alpha_audited_resource_descriptor_change_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           edgelq_audit_proto_v1alpha_audited_resource_descriptor_change_proto_goTypes,
		DependencyIndexes: edgelq_audit_proto_v1alpha_audited_resource_descriptor_change_proto_depIdxs,
		MessageInfos:      edgelq_audit_proto_v1alpha_audited_resource_descriptor_change_proto_msgTypes,
	}.Build()
	edgelq_audit_proto_v1alpha_audited_resource_descriptor_change_proto = out.File
	edgelq_audit_proto_v1alpha_audited_resource_descriptor_change_proto_rawDesc = nil
	edgelq_audit_proto_v1alpha_audited_resource_descriptor_change_proto_goTypes = nil
	edgelq_audit_proto_v1alpha_audited_resource_descriptor_change_proto_depIdxs = nil
}
