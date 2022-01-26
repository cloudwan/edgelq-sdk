// Code generated by protoc-gen-goten-go
// File: edgelq/iam/proto/v1alpha2/role_change.proto
// DO NOT EDIT!!!

package role

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

// RoleChange is used by Watch notifications Responses to describe change of
// single Role One of Added, Modified, Removed
type RoleChange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Role change
	//
	// Types that are valid to be assigned to ChangeType:
	//	*RoleChange_Added_
	//	*RoleChange_Modified_
	//	*RoleChange_Current_
	//	*RoleChange_Removed_
	ChangeType isRoleChange_ChangeType `protobuf_oneof:"change_type"`
}

func (m *RoleChange) Reset() {
	*m = RoleChange{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_iam_proto_v1alpha2_role_change_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *RoleChange) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*RoleChange) ProtoMessage() {}

func (m *RoleChange) ProtoReflect() preflect.Message {
	mi := &edgelq_iam_proto_v1alpha2_role_change_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*RoleChange) GotenMessage() {}

// Deprecated, Use RoleChange.ProtoReflect.Descriptor instead.
func (*RoleChange) Descriptor() ([]byte, []int) {
	return edgelq_iam_proto_v1alpha2_role_change_proto_rawDescGZIP(), []int{0}
}

func (m *RoleChange) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *RoleChange) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *RoleChange) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *RoleChange) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

type isRoleChange_ChangeType interface {
	isRoleChange_ChangeType()
}

type RoleChange_Added_ struct {
	// Added is returned when watched document is added, either created or
	// enters Query view
	Added *RoleChange_Added `protobuf:"bytes,1,opt,name=added,proto3,oneof" firestore:"added"`
}
type RoleChange_Modified_ struct {
	// Modified is returned when watched document is modified
	Modified *RoleChange_Modified `protobuf:"bytes,2,opt,name=modified,proto3,oneof" firestore:"modified"`
}
type RoleChange_Current_ struct {
	// Current is returned in stateless watch when document enters query view or
	// is modified within.
	Current *RoleChange_Current `protobuf:"bytes,4,opt,name=current,proto3,oneof" firestore:"current"`
}
type RoleChange_Removed_ struct {
	// Removed is returned when Role is deleted or leaves Query view
	Removed *RoleChange_Removed `protobuf:"bytes,3,opt,name=removed,proto3,oneof" firestore:"removed"`
}

func (*RoleChange_Added_) isRoleChange_ChangeType()    {}
func (*RoleChange_Modified_) isRoleChange_ChangeType() {}
func (*RoleChange_Current_) isRoleChange_ChangeType()  {}
func (*RoleChange_Removed_) isRoleChange_ChangeType()  {}
func (m *RoleChange) GetChangeType() isRoleChange_ChangeType {
	if m != nil {
		return m.ChangeType
	}
	return nil
}
func (m *RoleChange) GetAdded() *RoleChange_Added {
	if x, ok := m.GetChangeType().(*RoleChange_Added_); ok {
		return x.Added
	}
	return nil
}
func (m *RoleChange) GetModified() *RoleChange_Modified {
	if x, ok := m.GetChangeType().(*RoleChange_Modified_); ok {
		return x.Modified
	}
	return nil
}
func (m *RoleChange) GetCurrent() *RoleChange_Current {
	if x, ok := m.GetChangeType().(*RoleChange_Current_); ok {
		return x.Current
	}
	return nil
}
func (m *RoleChange) GetRemoved() *RoleChange_Removed {
	if x, ok := m.GetChangeType().(*RoleChange_Removed_); ok {
		return x.Removed
	}
	return nil
}
func (m *RoleChange) SetChangeType(ofv isRoleChange_ChangeType) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "isRoleChange_ChangeType", "RoleChange"))
	}
	m.ChangeType = ofv
}
func (m *RoleChange) SetAdded(fv *RoleChange_Added) {
	m.SetChangeType(&RoleChange_Added_{Added: fv})
}
func (m *RoleChange) SetModified(fv *RoleChange_Modified) {
	m.SetChangeType(&RoleChange_Modified_{Modified: fv})
}
func (m *RoleChange) SetCurrent(fv *RoleChange_Current) {
	m.SetChangeType(&RoleChange_Current_{Current: fv})
}
func (m *RoleChange) SetRemoved(fv *RoleChange_Removed) {
	m.SetChangeType(&RoleChange_Removed_{Removed: fv})
}

// Role has been added to query view
type RoleChange_Added struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Role          *Role `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty" firestore:"role"`
	// Integer describing index of added Role in resulting query view.
	ViewIndex int32 `protobuf:"varint,2,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty" firestore:"viewIndex"`
}

func (m *RoleChange_Added) Reset() {
	*m = RoleChange_Added{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_iam_proto_v1alpha2_role_change_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *RoleChange_Added) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*RoleChange_Added) ProtoMessage() {}

func (m *RoleChange_Added) ProtoReflect() preflect.Message {
	mi := &edgelq_iam_proto_v1alpha2_role_change_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*RoleChange_Added) GotenMessage() {}

// Deprecated, Use RoleChange_Added.ProtoReflect.Descriptor instead.
func (*RoleChange_Added) Descriptor() ([]byte, []int) {
	return edgelq_iam_proto_v1alpha2_role_change_proto_rawDescGZIP(), []int{0, 0}
}

func (m *RoleChange_Added) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *RoleChange_Added) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *RoleChange_Added) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *RoleChange_Added) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *RoleChange_Added) GetRole() *Role {
	if m != nil {
		return m.Role
	}
	return nil
}

func (m *RoleChange_Added) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *RoleChange_Added) SetRole(fv *Role) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Role", "RoleChange_Added"))
	}
	m.Role = fv
}

func (m *RoleChange_Added) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "RoleChange_Added"))
	}
	m.ViewIndex = fv
}

// Role changed some of it's fields - contains either full document or masked
// change
type RoleChange_Modified struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Name of modified Role
	Name *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// New version of Role or masked difference, depending on mask_changes
	// instrumentation of issued [WatchRoleRequest] or [WatchRolesRequest]
	Role *Role `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty" firestore:"role"`
	// Used when mask_changes is set, contains field paths of modified
	// properties.
	FieldMask *Role_FieldMask `protobuf:"bytes,3,opt,customtype=Role_FieldMask,name=field_mask,json=fieldMask,proto3" json:"field_mask,omitempty" firestore:"fieldMask"`
	// Previous view index specifies previous position of modified Role.
	// When modification doesn't affect sorted order, value will remain
	// identical to [view_index].
	PreviousViewIndex int32 `protobuf:"varint,4,opt,name=previous_view_index,json=previousViewIndex,proto3" json:"previous_view_index,omitempty" firestore:"previousViewIndex"`
	// Integer specifying Role new index in resulting query view.
	ViewIndex int32 `protobuf:"varint,5,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty" firestore:"viewIndex"`
}

func (m *RoleChange_Modified) Reset() {
	*m = RoleChange_Modified{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_iam_proto_v1alpha2_role_change_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *RoleChange_Modified) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*RoleChange_Modified) ProtoMessage() {}

func (m *RoleChange_Modified) ProtoReflect() preflect.Message {
	mi := &edgelq_iam_proto_v1alpha2_role_change_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*RoleChange_Modified) GotenMessage() {}

// Deprecated, Use RoleChange_Modified.ProtoReflect.Descriptor instead.
func (*RoleChange_Modified) Descriptor() ([]byte, []int) {
	return edgelq_iam_proto_v1alpha2_role_change_proto_rawDescGZIP(), []int{0, 1}
}

func (m *RoleChange_Modified) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *RoleChange_Modified) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *RoleChange_Modified) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *RoleChange_Modified) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *RoleChange_Modified) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *RoleChange_Modified) GetRole() *Role {
	if m != nil {
		return m.Role
	}
	return nil
}

func (m *RoleChange_Modified) GetFieldMask() *Role_FieldMask {
	if m != nil {
		return m.FieldMask
	}
	return nil
}

func (m *RoleChange_Modified) GetPreviousViewIndex() int32 {
	if m != nil {
		return m.PreviousViewIndex
	}
	return int32(0)
}

func (m *RoleChange_Modified) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *RoleChange_Modified) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "RoleChange_Modified"))
	}
	m.Name = fv
}

func (m *RoleChange_Modified) SetRole(fv *Role) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Role", "RoleChange_Modified"))
	}
	m.Role = fv
}

func (m *RoleChange_Modified) SetFieldMask(fv *Role_FieldMask) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "FieldMask", "RoleChange_Modified"))
	}
	m.FieldMask = fv
}

func (m *RoleChange_Modified) SetPreviousViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "PreviousViewIndex", "RoleChange_Modified"))
	}
	m.PreviousViewIndex = fv
}

func (m *RoleChange_Modified) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "RoleChange_Modified"))
	}
	m.ViewIndex = fv
}

// Role has been added or modified in a query view. Version used for stateless
// watching
type RoleChange_Current struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Role          *Role `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty" firestore:"role"`
}

func (m *RoleChange_Current) Reset() {
	*m = RoleChange_Current{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_iam_proto_v1alpha2_role_change_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *RoleChange_Current) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*RoleChange_Current) ProtoMessage() {}

func (m *RoleChange_Current) ProtoReflect() preflect.Message {
	mi := &edgelq_iam_proto_v1alpha2_role_change_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*RoleChange_Current) GotenMessage() {}

// Deprecated, Use RoleChange_Current.ProtoReflect.Descriptor instead.
func (*RoleChange_Current) Descriptor() ([]byte, []int) {
	return edgelq_iam_proto_v1alpha2_role_change_proto_rawDescGZIP(), []int{0, 2}
}

func (m *RoleChange_Current) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *RoleChange_Current) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *RoleChange_Current) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *RoleChange_Current) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *RoleChange_Current) GetRole() *Role {
	if m != nil {
		return m.Role
	}
	return nil
}

func (m *RoleChange_Current) SetRole(fv *Role) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Role", "RoleChange_Current"))
	}
	m.Role = fv
}

// Removed is returned when Role is deleted or leaves Query view
type RoleChange_Removed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Name          *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// Integer specifying removed Role index. Not populated in stateless watch
	// type.
	ViewIndex int32 `protobuf:"varint,2,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty" firestore:"viewIndex"`
}

func (m *RoleChange_Removed) Reset() {
	*m = RoleChange_Removed{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_iam_proto_v1alpha2_role_change_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *RoleChange_Removed) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*RoleChange_Removed) ProtoMessage() {}

func (m *RoleChange_Removed) ProtoReflect() preflect.Message {
	mi := &edgelq_iam_proto_v1alpha2_role_change_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*RoleChange_Removed) GotenMessage() {}

// Deprecated, Use RoleChange_Removed.ProtoReflect.Descriptor instead.
func (*RoleChange_Removed) Descriptor() ([]byte, []int) {
	return edgelq_iam_proto_v1alpha2_role_change_proto_rawDescGZIP(), []int{0, 3}
}

func (m *RoleChange_Removed) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *RoleChange_Removed) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *RoleChange_Removed) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *RoleChange_Removed) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *RoleChange_Removed) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *RoleChange_Removed) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *RoleChange_Removed) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "RoleChange_Removed"))
	}
	m.Name = fv
}

func (m *RoleChange_Removed) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "RoleChange_Removed"))
	}
	m.ViewIndex = fv
}

var edgelq_iam_proto_v1alpha2_role_change_proto preflect.FileDescriptor

var edgelq_iam_proto_v1alpha2_role_change_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x72, 0x6f, 0x6c, 0x65,
	0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x6e,
	0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x1a,
	0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f,
	0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x74,
	0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x65,
	0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xf6, 0x05, 0x0a, 0x0a, 0x52, 0x6f, 0x6c, 0x65, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x12, 0x3a, 0x0a, 0x05, 0x61, 0x64, 0x64, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x22, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x32, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e,
	0x41, 0x64, 0x64, 0x65, 0x64, 0x48, 0x00, 0x52, 0x05, 0x61, 0x64, 0x64, 0x65, 0x64, 0x12, 0x43,
	0x0a, 0x08, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x25, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x32, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x4d,
	0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x48, 0x00, 0x52, 0x08, 0x6d, 0x6f, 0x64, 0x69, 0x66,
	0x69, 0x65, 0x64, 0x12, 0x40, 0x0a, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x07, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x40, 0x0a, 0x07, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x48, 0x00, 0x52, 0x07,
	0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x1a, 0x52, 0x0a, 0x05, 0x41, 0x64, 0x64, 0x65, 0x64,
	0x12, 0x2a, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x32, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x76, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x1a, 0xf0, 0x01, 0x0a, 0x08,
	0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x20, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xb2, 0xda, 0x21, 0x08, 0x0a, 0x06, 0x0a, 0x04,
	0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x04, 0x72, 0x6f,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69,
	0x61, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x52, 0x6f, 0x6c, 0x65,
	0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x47, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x6d, 0x61, 0x73, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x42, 0x0c, 0xb2, 0xda, 0x21, 0x08, 0x32, 0x06, 0x0a, 0x04,
	0x52, 0x6f, 0x6c, 0x65, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x12,
	0x2e, 0x0a, 0x13, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x5f, 0x76, 0x69, 0x65, 0x77,
	0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x70, 0x72,
	0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x56, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12,
	0x1d, 0x0a, 0x0a, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x76, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x1a, 0x35,
	0x0a, 0x07, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x04, 0x72, 0x6f, 0x6c,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61,
	0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52,
	0x04, 0x72, 0x6f, 0x6c, 0x65, 0x1a, 0x4a, 0x0a, 0x07, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64,
	0x12, 0x20, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c,
	0xb2, 0xda, 0x21, 0x08, 0x0a, 0x06, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x76, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x3a, 0x0a, 0x9a, 0xd9, 0x21, 0x06, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x42, 0x0d, 0x0a,
	0x0b, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x42, 0x6d, 0xe8, 0xde,
	0x21, 0x00, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e,
	0x70, 0x62, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x42, 0x0f, 0x52, 0x6f, 0x6c,
	0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x00, 0x5a, 0x3b,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x32, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x3b, 0x72, 0x6f, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	edgelq_iam_proto_v1alpha2_role_change_proto_rawDescOnce sync.Once
	edgelq_iam_proto_v1alpha2_role_change_proto_rawDescData = edgelq_iam_proto_v1alpha2_role_change_proto_rawDesc
)

func edgelq_iam_proto_v1alpha2_role_change_proto_rawDescGZIP() []byte {
	edgelq_iam_proto_v1alpha2_role_change_proto_rawDescOnce.Do(func() {
		edgelq_iam_proto_v1alpha2_role_change_proto_rawDescData = protoimpl.X.CompressGZIP(edgelq_iam_proto_v1alpha2_role_change_proto_rawDescData)
	})
	return edgelq_iam_proto_v1alpha2_role_change_proto_rawDescData
}

var edgelq_iam_proto_v1alpha2_role_change_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var edgelq_iam_proto_v1alpha2_role_change_proto_goTypes = []interface{}{
	(*RoleChange)(nil),          // 0: ntt.iam.v1alpha2.RoleChange
	(*RoleChange_Added)(nil),    // 1: ntt.iam.v1alpha2.RoleChange.Added
	(*RoleChange_Modified)(nil), // 2: ntt.iam.v1alpha2.RoleChange.Modified
	(*RoleChange_Current)(nil),  // 3: ntt.iam.v1alpha2.RoleChange.Current
	(*RoleChange_Removed)(nil),  // 4: ntt.iam.v1alpha2.RoleChange.Removed
	(*Role)(nil),                // 5: ntt.iam.v1alpha2.Role
	(*Role_FieldMask)(nil),      // 6: ntt.iam.v1alpha2.Role_FieldMask
}
var edgelq_iam_proto_v1alpha2_role_change_proto_depIdxs = []int32{
	1, // 0: ntt.iam.v1alpha2.RoleChange.added:type_name -> ntt.iam.v1alpha2.RoleChange.Added
	2, // 1: ntt.iam.v1alpha2.RoleChange.modified:type_name -> ntt.iam.v1alpha2.RoleChange.Modified
	3, // 2: ntt.iam.v1alpha2.RoleChange.current:type_name -> ntt.iam.v1alpha2.RoleChange.Current
	4, // 3: ntt.iam.v1alpha2.RoleChange.removed:type_name -> ntt.iam.v1alpha2.RoleChange.Removed
	5, // 4: ntt.iam.v1alpha2.RoleChange.Added.role:type_name -> ntt.iam.v1alpha2.Role
	5, // 5: ntt.iam.v1alpha2.RoleChange.Modified.role:type_name -> ntt.iam.v1alpha2.Role
	6, // 6: ntt.iam.v1alpha2.RoleChange.Modified.field_mask:type_name -> ntt.iam.v1alpha2.Role_FieldMask
	5, // 7: ntt.iam.v1alpha2.RoleChange.Current.role:type_name -> ntt.iam.v1alpha2.Role
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { edgelq_iam_proto_v1alpha2_role_change_proto_init() }
func edgelq_iam_proto_v1alpha2_role_change_proto_init() {
	if edgelq_iam_proto_v1alpha2_role_change_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		edgelq_iam_proto_v1alpha2_role_change_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleChange); i {
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
		edgelq_iam_proto_v1alpha2_role_change_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleChange_Added); i {
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
		edgelq_iam_proto_v1alpha2_role_change_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleChange_Modified); i {
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
		edgelq_iam_proto_v1alpha2_role_change_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleChange_Current); i {
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
		edgelq_iam_proto_v1alpha2_role_change_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleChange_Removed); i {
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

	edgelq_iam_proto_v1alpha2_role_change_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*RoleChange_Added_)(nil),
		(*RoleChange_Modified_)(nil),
		(*RoleChange_Current_)(nil),
		(*RoleChange_Removed_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: edgelq_iam_proto_v1alpha2_role_change_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           edgelq_iam_proto_v1alpha2_role_change_proto_goTypes,
		DependencyIndexes: edgelq_iam_proto_v1alpha2_role_change_proto_depIdxs,
		MessageInfos:      edgelq_iam_proto_v1alpha2_role_change_proto_msgTypes,
	}.Build()
	edgelq_iam_proto_v1alpha2_role_change_proto = out.File
	edgelq_iam_proto_v1alpha2_role_change_proto_rawDesc = nil
	edgelq_iam_proto_v1alpha2_role_change_proto_goTypes = nil
	edgelq_iam_proto_v1alpha2_role_change_proto_depIdxs = nil
}
