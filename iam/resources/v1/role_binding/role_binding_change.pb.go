// Code generated by protoc-gen-goten-go
// File: edgelq/iam/proto/v1/role_binding_change.proto
// DO NOT EDIT!!!

package role_binding

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
	organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1/organization"
	project "github.com/cloudwan/edgelq-sdk/iam/resources/v1/project"
	meta_service "github.com/cloudwan/goten-sdk/meta-service/resources/v1/service"
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
	_ = &project.Project{}
	_ = &fieldmaskpb.FieldMask{}
	_ = &meta_service.Service{}
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// RoleBindingChange is used by Watch notifications Responses to describe change
// of single RoleBinding One of Added, Modified, Removed
type RoleBindingChange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// RoleBinding change
	//
	// Types that are valid to be assigned to ChangeType:
	//	*RoleBindingChange_Added_
	//	*RoleBindingChange_Modified_
	//	*RoleBindingChange_Current_
	//	*RoleBindingChange_Removed_
	ChangeType isRoleBindingChange_ChangeType `protobuf_oneof:"change_type"`
}

func (m *RoleBindingChange) Reset() {
	*m = RoleBindingChange{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_iam_proto_v1_role_binding_change_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *RoleBindingChange) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*RoleBindingChange) ProtoMessage() {}

func (m *RoleBindingChange) ProtoReflect() preflect.Message {
	mi := &edgelq_iam_proto_v1_role_binding_change_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*RoleBindingChange) GotenMessage() {}

// Deprecated, Use RoleBindingChange.ProtoReflect.Descriptor instead.
func (*RoleBindingChange) Descriptor() ([]byte, []int) {
	return edgelq_iam_proto_v1_role_binding_change_proto_rawDescGZIP(), []int{0}
}

func (m *RoleBindingChange) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *RoleBindingChange) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *RoleBindingChange) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *RoleBindingChange) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

type isRoleBindingChange_ChangeType interface {
	isRoleBindingChange_ChangeType()
}

type RoleBindingChange_Added_ struct {
	// Added is returned when watched document is added, either created or
	// enters Query view
	Added *RoleBindingChange_Added `protobuf:"bytes,1,opt,name=added,proto3,oneof" firestore:"added"`
}
type RoleBindingChange_Modified_ struct {
	// Modified is returned when watched document is modified
	Modified *RoleBindingChange_Modified `protobuf:"bytes,2,opt,name=modified,proto3,oneof" firestore:"modified"`
}
type RoleBindingChange_Current_ struct {
	// Current is returned in stateless watch when document enters query view or
	// is modified within.
	Current *RoleBindingChange_Current `protobuf:"bytes,4,opt,name=current,proto3,oneof" firestore:"current"`
}
type RoleBindingChange_Removed_ struct {
	// Removed is returned when RoleBinding is deleted or leaves Query view
	Removed *RoleBindingChange_Removed `protobuf:"bytes,3,opt,name=removed,proto3,oneof" firestore:"removed"`
}

func (*RoleBindingChange_Added_) isRoleBindingChange_ChangeType()    {}
func (*RoleBindingChange_Modified_) isRoleBindingChange_ChangeType() {}
func (*RoleBindingChange_Current_) isRoleBindingChange_ChangeType()  {}
func (*RoleBindingChange_Removed_) isRoleBindingChange_ChangeType()  {}
func (m *RoleBindingChange) GetChangeType() isRoleBindingChange_ChangeType {
	if m != nil {
		return m.ChangeType
	}
	return nil
}
func (m *RoleBindingChange) GetAdded() *RoleBindingChange_Added {
	if x, ok := m.GetChangeType().(*RoleBindingChange_Added_); ok {
		return x.Added
	}
	return nil
}
func (m *RoleBindingChange) GetModified() *RoleBindingChange_Modified {
	if x, ok := m.GetChangeType().(*RoleBindingChange_Modified_); ok {
		return x.Modified
	}
	return nil
}
func (m *RoleBindingChange) GetCurrent() *RoleBindingChange_Current {
	if x, ok := m.GetChangeType().(*RoleBindingChange_Current_); ok {
		return x.Current
	}
	return nil
}
func (m *RoleBindingChange) GetRemoved() *RoleBindingChange_Removed {
	if x, ok := m.GetChangeType().(*RoleBindingChange_Removed_); ok {
		return x.Removed
	}
	return nil
}
func (m *RoleBindingChange) SetChangeType(ofv isRoleBindingChange_ChangeType) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "isRoleBindingChange_ChangeType", "RoleBindingChange"))
	}
	m.ChangeType = ofv
}
func (m *RoleBindingChange) SetAdded(fv *RoleBindingChange_Added) {
	m.SetChangeType(&RoleBindingChange_Added_{Added: fv})
}
func (m *RoleBindingChange) SetModified(fv *RoleBindingChange_Modified) {
	m.SetChangeType(&RoleBindingChange_Modified_{Modified: fv})
}
func (m *RoleBindingChange) SetCurrent(fv *RoleBindingChange_Current) {
	m.SetChangeType(&RoleBindingChange_Current_{Current: fv})
}
func (m *RoleBindingChange) SetRemoved(fv *RoleBindingChange_Removed) {
	m.SetChangeType(&RoleBindingChange_Removed_{Removed: fv})
}

// RoleBinding has been added to query view
type RoleBindingChange_Added struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	RoleBinding   *RoleBinding `protobuf:"bytes,1,opt,name=role_binding,json=roleBinding,proto3" json:"role_binding,omitempty" firestore:"roleBinding"`
	// Integer describing index of added RoleBinding in resulting query view.
	ViewIndex int32 `protobuf:"varint,2,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty" firestore:"viewIndex"`
}

func (m *RoleBindingChange_Added) Reset() {
	*m = RoleBindingChange_Added{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_iam_proto_v1_role_binding_change_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *RoleBindingChange_Added) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*RoleBindingChange_Added) ProtoMessage() {}

func (m *RoleBindingChange_Added) ProtoReflect() preflect.Message {
	mi := &edgelq_iam_proto_v1_role_binding_change_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*RoleBindingChange_Added) GotenMessage() {}

// Deprecated, Use RoleBindingChange_Added.ProtoReflect.Descriptor instead.
func (*RoleBindingChange_Added) Descriptor() ([]byte, []int) {
	return edgelq_iam_proto_v1_role_binding_change_proto_rawDescGZIP(), []int{0, 0}
}

func (m *RoleBindingChange_Added) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *RoleBindingChange_Added) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *RoleBindingChange_Added) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *RoleBindingChange_Added) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *RoleBindingChange_Added) GetRoleBinding() *RoleBinding {
	if m != nil {
		return m.RoleBinding
	}
	return nil
}

func (m *RoleBindingChange_Added) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *RoleBindingChange_Added) SetRoleBinding(fv *RoleBinding) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "RoleBinding", "RoleBindingChange_Added"))
	}
	m.RoleBinding = fv
}

func (m *RoleBindingChange_Added) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "RoleBindingChange_Added"))
	}
	m.ViewIndex = fv
}

// RoleBinding changed some of it's fields - contains either full document or
// masked change
type RoleBindingChange_Modified struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Name of modified RoleBinding
	Name *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// New version of RoleBinding or masked difference, depending on
	// mask_changes instrumentation of issued [WatchRoleBindingRequest] or
	// [WatchRoleBindingsRequest]
	RoleBinding *RoleBinding `protobuf:"bytes,2,opt,name=role_binding,json=roleBinding,proto3" json:"role_binding,omitempty" firestore:"roleBinding"`
	// Used when mask_changes is set, contains field paths of modified
	// properties.
	FieldMask *RoleBinding_FieldMask `protobuf:"bytes,3,opt,customtype=RoleBinding_FieldMask,name=field_mask,json=fieldMask,proto3" json:"field_mask,omitempty" firestore:"fieldMask"`
	// Previous view index specifies previous position of modified RoleBinding.
	// When modification doesn't affect sorted order, value will remain
	// identical to [view_index].
	PreviousViewIndex int32 `protobuf:"varint,4,opt,name=previous_view_index,json=previousViewIndex,proto3" json:"previous_view_index,omitempty" firestore:"previousViewIndex"`
	// Integer specifying RoleBinding new index in resulting query view.
	ViewIndex int32 `protobuf:"varint,5,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty" firestore:"viewIndex"`
}

func (m *RoleBindingChange_Modified) Reset() {
	*m = RoleBindingChange_Modified{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_iam_proto_v1_role_binding_change_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *RoleBindingChange_Modified) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*RoleBindingChange_Modified) ProtoMessage() {}

func (m *RoleBindingChange_Modified) ProtoReflect() preflect.Message {
	mi := &edgelq_iam_proto_v1_role_binding_change_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*RoleBindingChange_Modified) GotenMessage() {}

// Deprecated, Use RoleBindingChange_Modified.ProtoReflect.Descriptor instead.
func (*RoleBindingChange_Modified) Descriptor() ([]byte, []int) {
	return edgelq_iam_proto_v1_role_binding_change_proto_rawDescGZIP(), []int{0, 1}
}

func (m *RoleBindingChange_Modified) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *RoleBindingChange_Modified) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *RoleBindingChange_Modified) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *RoleBindingChange_Modified) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *RoleBindingChange_Modified) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *RoleBindingChange_Modified) GetRoleBinding() *RoleBinding {
	if m != nil {
		return m.RoleBinding
	}
	return nil
}

func (m *RoleBindingChange_Modified) GetFieldMask() *RoleBinding_FieldMask {
	if m != nil {
		return m.FieldMask
	}
	return nil
}

func (m *RoleBindingChange_Modified) GetPreviousViewIndex() int32 {
	if m != nil {
		return m.PreviousViewIndex
	}
	return int32(0)
}

func (m *RoleBindingChange_Modified) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *RoleBindingChange_Modified) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "RoleBindingChange_Modified"))
	}
	m.Name = fv
}

func (m *RoleBindingChange_Modified) SetRoleBinding(fv *RoleBinding) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "RoleBinding", "RoleBindingChange_Modified"))
	}
	m.RoleBinding = fv
}

func (m *RoleBindingChange_Modified) SetFieldMask(fv *RoleBinding_FieldMask) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "FieldMask", "RoleBindingChange_Modified"))
	}
	m.FieldMask = fv
}

func (m *RoleBindingChange_Modified) SetPreviousViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "PreviousViewIndex", "RoleBindingChange_Modified"))
	}
	m.PreviousViewIndex = fv
}

func (m *RoleBindingChange_Modified) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "RoleBindingChange_Modified"))
	}
	m.ViewIndex = fv
}

// RoleBinding has been added or modified in a query view. Version used for
// stateless watching
type RoleBindingChange_Current struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	RoleBinding   *RoleBinding `protobuf:"bytes,1,opt,name=role_binding,json=roleBinding,proto3" json:"role_binding,omitempty" firestore:"roleBinding"`
}

func (m *RoleBindingChange_Current) Reset() {
	*m = RoleBindingChange_Current{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_iam_proto_v1_role_binding_change_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *RoleBindingChange_Current) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*RoleBindingChange_Current) ProtoMessage() {}

func (m *RoleBindingChange_Current) ProtoReflect() preflect.Message {
	mi := &edgelq_iam_proto_v1_role_binding_change_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*RoleBindingChange_Current) GotenMessage() {}

// Deprecated, Use RoleBindingChange_Current.ProtoReflect.Descriptor instead.
func (*RoleBindingChange_Current) Descriptor() ([]byte, []int) {
	return edgelq_iam_proto_v1_role_binding_change_proto_rawDescGZIP(), []int{0, 2}
}

func (m *RoleBindingChange_Current) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *RoleBindingChange_Current) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *RoleBindingChange_Current) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *RoleBindingChange_Current) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *RoleBindingChange_Current) GetRoleBinding() *RoleBinding {
	if m != nil {
		return m.RoleBinding
	}
	return nil
}

func (m *RoleBindingChange_Current) SetRoleBinding(fv *RoleBinding) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "RoleBinding", "RoleBindingChange_Current"))
	}
	m.RoleBinding = fv
}

// Removed is returned when RoleBinding is deleted or leaves Query view
type RoleBindingChange_Removed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Name          *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// Integer specifying removed RoleBinding index. Not populated in stateless
	// watch type.
	ViewIndex int32 `protobuf:"varint,2,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty" firestore:"viewIndex"`
}

func (m *RoleBindingChange_Removed) Reset() {
	*m = RoleBindingChange_Removed{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_iam_proto_v1_role_binding_change_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *RoleBindingChange_Removed) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*RoleBindingChange_Removed) ProtoMessage() {}

func (m *RoleBindingChange_Removed) ProtoReflect() preflect.Message {
	mi := &edgelq_iam_proto_v1_role_binding_change_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*RoleBindingChange_Removed) GotenMessage() {}

// Deprecated, Use RoleBindingChange_Removed.ProtoReflect.Descriptor instead.
func (*RoleBindingChange_Removed) Descriptor() ([]byte, []int) {
	return edgelq_iam_proto_v1_role_binding_change_proto_rawDescGZIP(), []int{0, 3}
}

func (m *RoleBindingChange_Removed) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *RoleBindingChange_Removed) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *RoleBindingChange_Removed) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *RoleBindingChange_Removed) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *RoleBindingChange_Removed) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *RoleBindingChange_Removed) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *RoleBindingChange_Removed) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "RoleBindingChange_Removed"))
	}
	m.Name = fv
}

func (m *RoleBindingChange_Removed) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "RoleBindingChange_Removed"))
	}
	m.ViewIndex = fv
}

var edgelq_iam_proto_v1_role_binding_change_proto preflect.FileDescriptor

var edgelq_iam_proto_v1_role_binding_change_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0a, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x1a, 0x19, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61,
	0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x74, 0x65,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x65, 0x64, 0x67, 0x65, 0x6c,
	0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x72,
	0x6f, 0x6c, 0x65, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xcd, 0x06, 0x0a, 0x11, 0x52, 0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x3b, 0x0a, 0x05, 0x61, 0x64, 0x64, 0x65, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x65, 0x64, 0x48, 0x00, 0x52, 0x05, 0x61,
	0x64, 0x64, 0x65, 0x64, 0x12, 0x44, 0x0a, 0x08, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x48, 0x00,
	0x52, 0x08, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x41, 0x0a, 0x07, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6e, 0x74,
	0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x48, 0x00, 0x52, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x41, 0x0a,
	0x07, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25,
	0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6c, 0x65,
	0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x64, 0x48, 0x00, 0x52, 0x07, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64,
	0x1a, 0x62, 0x0a, 0x05, 0x41, 0x64, 0x64, 0x65, 0x64, 0x12, 0x3a, 0x0a, 0x0c, 0x72, 0x6f, 0x6c,
	0x65, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6c,
	0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x0b, 0x72, 0x6f, 0x6c, 0x65, 0x42, 0x69,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x76, 0x69, 0x65, 0x77, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x1a, 0x8e, 0x02, 0x0a, 0x08, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65,
	0x64, 0x12, 0x27, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x13, 0xb2, 0xda, 0x21, 0x0f, 0x0a, 0x0d, 0x0a, 0x0b, 0x52, 0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3a, 0x0a, 0x0c, 0x72, 0x6f,
	0x6c, 0x65, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f,
	0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x0b, 0x72, 0x6f, 0x6c, 0x65, 0x42,
	0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x4e, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x6d, 0x61, 0x73, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x42, 0x13, 0xb2, 0xda, 0x21, 0x0f, 0x32, 0x0d, 0x0a, 0x0b,
	0x52, 0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x09, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x2e, 0x0a, 0x13, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f,
	0x75, 0x73, 0x5f, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x11, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x56, 0x69, 0x65,
	0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x76, 0x69, 0x65, 0x77,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x1a, 0x45, 0x0a, 0x07, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x12, 0x3a, 0x0a, 0x0c, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52,
	0x0b, 0x72, 0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x1a, 0x51, 0x0a, 0x07,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x12, 0x27, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x13, 0xb2, 0xda, 0x21, 0x0f, 0x0a, 0x0d, 0x0a, 0x0b, 0x52,
	0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x76, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x3a,
	0x11, 0x9a, 0xd9, 0x21, 0x0d, 0x0a, 0x0b, 0x52, 0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x42, 0x0d, 0x0a, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x42, 0x78, 0xe8, 0xde, 0x21, 0x00, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x74, 0x74,
	0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x42, 0x16, 0x52, 0x6f, 0x6c, 0x65,
	0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x00, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71,
	0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x3b, 0x72,
	0x6f, 0x6c, 0x65, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	edgelq_iam_proto_v1_role_binding_change_proto_rawDescOnce sync.Once
	edgelq_iam_proto_v1_role_binding_change_proto_rawDescData = edgelq_iam_proto_v1_role_binding_change_proto_rawDesc
)

func edgelq_iam_proto_v1_role_binding_change_proto_rawDescGZIP() []byte {
	edgelq_iam_proto_v1_role_binding_change_proto_rawDescOnce.Do(func() {
		edgelq_iam_proto_v1_role_binding_change_proto_rawDescData = protoimpl.X.CompressGZIP(edgelq_iam_proto_v1_role_binding_change_proto_rawDescData)
	})
	return edgelq_iam_proto_v1_role_binding_change_proto_rawDescData
}

var edgelq_iam_proto_v1_role_binding_change_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var edgelq_iam_proto_v1_role_binding_change_proto_goTypes = []interface{}{
	(*RoleBindingChange)(nil),          // 0: ntt.iam.v1.RoleBindingChange
	(*RoleBindingChange_Added)(nil),    // 1: ntt.iam.v1.RoleBindingChange.Added
	(*RoleBindingChange_Modified)(nil), // 2: ntt.iam.v1.RoleBindingChange.Modified
	(*RoleBindingChange_Current)(nil),  // 3: ntt.iam.v1.RoleBindingChange.Current
	(*RoleBindingChange_Removed)(nil),  // 4: ntt.iam.v1.RoleBindingChange.Removed
	(*RoleBinding)(nil),                // 5: ntt.iam.v1.RoleBinding
	(*RoleBinding_FieldMask)(nil),      // 6: ntt.iam.v1.RoleBinding_FieldMask
}
var edgelq_iam_proto_v1_role_binding_change_proto_depIdxs = []int32{
	1, // 0: ntt.iam.v1.RoleBindingChange.added:type_name -> ntt.iam.v1.RoleBindingChange.Added
	2, // 1: ntt.iam.v1.RoleBindingChange.modified:type_name -> ntt.iam.v1.RoleBindingChange.Modified
	3, // 2: ntt.iam.v1.RoleBindingChange.current:type_name -> ntt.iam.v1.RoleBindingChange.Current
	4, // 3: ntt.iam.v1.RoleBindingChange.removed:type_name -> ntt.iam.v1.RoleBindingChange.Removed
	5, // 4: ntt.iam.v1.RoleBindingChange.Added.role_binding:type_name -> ntt.iam.v1.RoleBinding
	5, // 5: ntt.iam.v1.RoleBindingChange.Modified.role_binding:type_name -> ntt.iam.v1.RoleBinding
	6, // 6: ntt.iam.v1.RoleBindingChange.Modified.field_mask:type_name -> ntt.iam.v1.RoleBinding_FieldMask
	5, // 7: ntt.iam.v1.RoleBindingChange.Current.role_binding:type_name -> ntt.iam.v1.RoleBinding
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { edgelq_iam_proto_v1_role_binding_change_proto_init() }
func edgelq_iam_proto_v1_role_binding_change_proto_init() {
	if edgelq_iam_proto_v1_role_binding_change_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		edgelq_iam_proto_v1_role_binding_change_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleBindingChange); i {
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
		edgelq_iam_proto_v1_role_binding_change_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleBindingChange_Added); i {
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
		edgelq_iam_proto_v1_role_binding_change_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleBindingChange_Modified); i {
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
		edgelq_iam_proto_v1_role_binding_change_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleBindingChange_Current); i {
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
		edgelq_iam_proto_v1_role_binding_change_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleBindingChange_Removed); i {
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

	edgelq_iam_proto_v1_role_binding_change_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*RoleBindingChange_Added_)(nil),
		(*RoleBindingChange_Modified_)(nil),
		(*RoleBindingChange_Current_)(nil),
		(*RoleBindingChange_Removed_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: edgelq_iam_proto_v1_role_binding_change_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           edgelq_iam_proto_v1_role_binding_change_proto_goTypes,
		DependencyIndexes: edgelq_iam_proto_v1_role_binding_change_proto_depIdxs,
		MessageInfos:      edgelq_iam_proto_v1_role_binding_change_proto_msgTypes,
	}.Build()
	edgelq_iam_proto_v1_role_binding_change_proto = out.File
	edgelq_iam_proto_v1_role_binding_change_proto_rawDesc = nil
	edgelq_iam_proto_v1_role_binding_change_proto_goTypes = nil
	edgelq_iam_proto_v1_role_binding_change_proto_depIdxs = nil
}
