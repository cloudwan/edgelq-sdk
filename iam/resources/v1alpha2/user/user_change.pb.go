// Code generated by protoc-gen-goten-go
// File: edgelq/iam/proto/v1alpha2/user_change.proto
// DO NOT EDIT!!!

package user

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

// UserChange is used by Watch notifications Responses to describe change of
// single User One of Added, Modified, Removed
type UserChange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// User change
	//
	// Types that are valid to be assigned to ChangeType:
	//	*UserChange_Added_
	//	*UserChange_Modified_
	//	*UserChange_Current_
	//	*UserChange_Removed_
	ChangeType isUserChange_ChangeType `protobuf_oneof:"change_type"`
}

func (m *UserChange) Reset() {
	*m = UserChange{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_iam_proto_v1alpha2_user_change_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *UserChange) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*UserChange) ProtoMessage() {}

func (m *UserChange) ProtoReflect() preflect.Message {
	mi := &edgelq_iam_proto_v1alpha2_user_change_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*UserChange) GotenMessage() {}

// Deprecated, Use UserChange.ProtoReflect.Descriptor instead.
func (*UserChange) Descriptor() ([]byte, []int) {
	return edgelq_iam_proto_v1alpha2_user_change_proto_rawDescGZIP(), []int{0}
}

func (m *UserChange) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *UserChange) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *UserChange) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *UserChange) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

type isUserChange_ChangeType interface {
	isUserChange_ChangeType()
}

type UserChange_Added_ struct {
	// Added is returned when watched document is added, either created or
	// enters Query view
	Added *UserChange_Added `protobuf:"bytes,1,opt,name=added,proto3,oneof" firestore:"added"`
}
type UserChange_Modified_ struct {
	// Modified is returned when watched document is modified
	Modified *UserChange_Modified `protobuf:"bytes,2,opt,name=modified,proto3,oneof" firestore:"modified"`
}
type UserChange_Current_ struct {
	// Current is returned in stateless watch when document enters query view or
	// is modified within.
	Current *UserChange_Current `protobuf:"bytes,4,opt,name=current,proto3,oneof" firestore:"current"`
}
type UserChange_Removed_ struct {
	// Removed is returned when User is deleted or leaves Query view
	Removed *UserChange_Removed `protobuf:"bytes,3,opt,name=removed,proto3,oneof" firestore:"removed"`
}

func (*UserChange_Added_) isUserChange_ChangeType()    {}
func (*UserChange_Modified_) isUserChange_ChangeType() {}
func (*UserChange_Current_) isUserChange_ChangeType()  {}
func (*UserChange_Removed_) isUserChange_ChangeType()  {}
func (m *UserChange) GetChangeType() isUserChange_ChangeType {
	if m != nil {
		return m.ChangeType
	}
	return nil
}
func (m *UserChange) GetAdded() *UserChange_Added {
	if x, ok := m.GetChangeType().(*UserChange_Added_); ok {
		return x.Added
	}
	return nil
}
func (m *UserChange) GetModified() *UserChange_Modified {
	if x, ok := m.GetChangeType().(*UserChange_Modified_); ok {
		return x.Modified
	}
	return nil
}
func (m *UserChange) GetCurrent() *UserChange_Current {
	if x, ok := m.GetChangeType().(*UserChange_Current_); ok {
		return x.Current
	}
	return nil
}
func (m *UserChange) GetRemoved() *UserChange_Removed {
	if x, ok := m.GetChangeType().(*UserChange_Removed_); ok {
		return x.Removed
	}
	return nil
}
func (m *UserChange) SetChangeType(ofv isUserChange_ChangeType) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "isUserChange_ChangeType", "UserChange"))
	}
	m.ChangeType = ofv
}
func (m *UserChange) SetAdded(fv *UserChange_Added) {
	m.SetChangeType(&UserChange_Added_{Added: fv})
}
func (m *UserChange) SetModified(fv *UserChange_Modified) {
	m.SetChangeType(&UserChange_Modified_{Modified: fv})
}
func (m *UserChange) SetCurrent(fv *UserChange_Current) {
	m.SetChangeType(&UserChange_Current_{Current: fv})
}
func (m *UserChange) SetRemoved(fv *UserChange_Removed) {
	m.SetChangeType(&UserChange_Removed_{Removed: fv})
}

// User has been added to query view
type UserChange_Added struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	User          *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty" firestore:"user"`
	// Integer describing index of added User in resulting query view.
	ViewIndex int32 `protobuf:"varint,2,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty" firestore:"viewIndex"`
}

func (m *UserChange_Added) Reset() {
	*m = UserChange_Added{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_iam_proto_v1alpha2_user_change_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *UserChange_Added) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*UserChange_Added) ProtoMessage() {}

func (m *UserChange_Added) ProtoReflect() preflect.Message {
	mi := &edgelq_iam_proto_v1alpha2_user_change_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*UserChange_Added) GotenMessage() {}

// Deprecated, Use UserChange_Added.ProtoReflect.Descriptor instead.
func (*UserChange_Added) Descriptor() ([]byte, []int) {
	return edgelq_iam_proto_v1alpha2_user_change_proto_rawDescGZIP(), []int{0, 0}
}

func (m *UserChange_Added) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *UserChange_Added) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *UserChange_Added) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *UserChange_Added) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *UserChange_Added) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *UserChange_Added) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *UserChange_Added) SetUser(fv *User) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "User", "UserChange_Added"))
	}
	m.User = fv
}

func (m *UserChange_Added) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "UserChange_Added"))
	}
	m.ViewIndex = fv
}

// User changed some of it's fields - contains either full document or masked
// change
type UserChange_Modified struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Name of modified User
	Name *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// New version of User or masked difference, depending on mask_changes
	// instrumentation of issued [WatchUserRequest] or [WatchUsersRequest]
	User *User `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty" firestore:"user"`
	// Used when mask_changes is set, contains field paths of modified
	// properties.
	FieldMask *User_FieldMask `protobuf:"bytes,3,opt,customtype=User_FieldMask,name=field_mask,json=fieldMask,proto3" json:"field_mask,omitempty" firestore:"fieldMask"`
	// Previous view index specifies previous position of modified User.
	// When modification doesn't affect sorted order, value will remain
	// identical to [view_index].
	PreviousViewIndex int32 `protobuf:"varint,4,opt,name=previous_view_index,json=previousViewIndex,proto3" json:"previous_view_index,omitempty" firestore:"previousViewIndex"`
	// Integer specifying User new index in resulting query view.
	ViewIndex int32 `protobuf:"varint,5,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty" firestore:"viewIndex"`
}

func (m *UserChange_Modified) Reset() {
	*m = UserChange_Modified{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_iam_proto_v1alpha2_user_change_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *UserChange_Modified) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*UserChange_Modified) ProtoMessage() {}

func (m *UserChange_Modified) ProtoReflect() preflect.Message {
	mi := &edgelq_iam_proto_v1alpha2_user_change_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*UserChange_Modified) GotenMessage() {}

// Deprecated, Use UserChange_Modified.ProtoReflect.Descriptor instead.
func (*UserChange_Modified) Descriptor() ([]byte, []int) {
	return edgelq_iam_proto_v1alpha2_user_change_proto_rawDescGZIP(), []int{0, 1}
}

func (m *UserChange_Modified) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *UserChange_Modified) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *UserChange_Modified) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *UserChange_Modified) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *UserChange_Modified) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *UserChange_Modified) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *UserChange_Modified) GetFieldMask() *User_FieldMask {
	if m != nil {
		return m.FieldMask
	}
	return nil
}

func (m *UserChange_Modified) GetPreviousViewIndex() int32 {
	if m != nil {
		return m.PreviousViewIndex
	}
	return int32(0)
}

func (m *UserChange_Modified) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *UserChange_Modified) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "UserChange_Modified"))
	}
	m.Name = fv
}

func (m *UserChange_Modified) SetUser(fv *User) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "User", "UserChange_Modified"))
	}
	m.User = fv
}

func (m *UserChange_Modified) SetFieldMask(fv *User_FieldMask) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "FieldMask", "UserChange_Modified"))
	}
	m.FieldMask = fv
}

func (m *UserChange_Modified) SetPreviousViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "PreviousViewIndex", "UserChange_Modified"))
	}
	m.PreviousViewIndex = fv
}

func (m *UserChange_Modified) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "UserChange_Modified"))
	}
	m.ViewIndex = fv
}

// User has been added or modified in a query view. Version used for stateless
// watching
type UserChange_Current struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	User          *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty" firestore:"user"`
}

func (m *UserChange_Current) Reset() {
	*m = UserChange_Current{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_iam_proto_v1alpha2_user_change_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *UserChange_Current) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*UserChange_Current) ProtoMessage() {}

func (m *UserChange_Current) ProtoReflect() preflect.Message {
	mi := &edgelq_iam_proto_v1alpha2_user_change_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*UserChange_Current) GotenMessage() {}

// Deprecated, Use UserChange_Current.ProtoReflect.Descriptor instead.
func (*UserChange_Current) Descriptor() ([]byte, []int) {
	return edgelq_iam_proto_v1alpha2_user_change_proto_rawDescGZIP(), []int{0, 2}
}

func (m *UserChange_Current) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *UserChange_Current) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *UserChange_Current) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *UserChange_Current) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *UserChange_Current) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *UserChange_Current) SetUser(fv *User) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "User", "UserChange_Current"))
	}
	m.User = fv
}

// Removed is returned when User is deleted or leaves Query view
type UserChange_Removed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Name          *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// Integer specifying removed User index. Not populated in stateless watch
	// type.
	ViewIndex int32 `protobuf:"varint,2,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty" firestore:"viewIndex"`
}

func (m *UserChange_Removed) Reset() {
	*m = UserChange_Removed{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_iam_proto_v1alpha2_user_change_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *UserChange_Removed) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*UserChange_Removed) ProtoMessage() {}

func (m *UserChange_Removed) ProtoReflect() preflect.Message {
	mi := &edgelq_iam_proto_v1alpha2_user_change_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*UserChange_Removed) GotenMessage() {}

// Deprecated, Use UserChange_Removed.ProtoReflect.Descriptor instead.
func (*UserChange_Removed) Descriptor() ([]byte, []int) {
	return edgelq_iam_proto_v1alpha2_user_change_proto_rawDescGZIP(), []int{0, 3}
}

func (m *UserChange_Removed) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *UserChange_Removed) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *UserChange_Removed) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *UserChange_Removed) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *UserChange_Removed) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *UserChange_Removed) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *UserChange_Removed) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "UserChange_Removed"))
	}
	m.Name = fv
}

func (m *UserChange_Removed) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "UserChange_Removed"))
	}
	m.ViewIndex = fv
}

var edgelq_iam_proto_v1alpha2_user_change_proto preflect.FileDescriptor

var edgelq_iam_proto_v1alpha2_user_change_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x75, 0x73, 0x65, 0x72,
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
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xf6, 0x05, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x12, 0x3a, 0x0a, 0x05, 0x61, 0x64, 0x64, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x22, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x32, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e,
	0x41, 0x64, 0x64, 0x65, 0x64, 0x48, 0x00, 0x52, 0x05, 0x61, 0x64, 0x64, 0x65, 0x64, 0x12, 0x43,
	0x0a, 0x08, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x25, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x32, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x4d,
	0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x48, 0x00, 0x52, 0x08, 0x6d, 0x6f, 0x64, 0x69, 0x66,
	0x69, 0x65, 0x64, 0x12, 0x40, 0x0a, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x07, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x40, 0x0a, 0x07, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x48, 0x00, 0x52, 0x07,
	0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x1a, 0x52, 0x0a, 0x05, 0x41, 0x64, 0x64, 0x65, 0x64,
	0x12, 0x2a, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x32, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a,
	0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x76, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x1a, 0xf0, 0x01, 0x0a, 0x08,
	0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x20, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xb2, 0xda, 0x21, 0x08, 0x0a, 0x06, 0x0a, 0x04,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69,
	0x61, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x47, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x6d, 0x61, 0x73, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x42, 0x0c, 0xb2, 0xda, 0x21, 0x08, 0x32, 0x06, 0x0a, 0x04,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x12,
	0x2e, 0x0a, 0x13, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x5f, 0x76, 0x69, 0x65, 0x77,
	0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x70, 0x72,
	0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x56, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12,
	0x1d, 0x0a, 0x0a, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x76, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x1a, 0x35,
	0x0a, 0x07, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61,
	0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x4a, 0x0a, 0x07, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64,
	0x12, 0x20, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c,
	0xb2, 0xda, 0x21, 0x08, 0x0a, 0x06, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x76, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x3a, 0x0a, 0x9a, 0xd9, 0x21, 0x06, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x42, 0x0d, 0x0a,
	0x0b, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x42, 0x6d, 0xe8, 0xde,
	0x21, 0x00, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e,
	0x70, 0x62, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x42, 0x0f, 0x55, 0x73, 0x65,
	0x72, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x00, 0x5a, 0x3b,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x32, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x3b, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	edgelq_iam_proto_v1alpha2_user_change_proto_rawDescOnce sync.Once
	edgelq_iam_proto_v1alpha2_user_change_proto_rawDescData = edgelq_iam_proto_v1alpha2_user_change_proto_rawDesc
)

func edgelq_iam_proto_v1alpha2_user_change_proto_rawDescGZIP() []byte {
	edgelq_iam_proto_v1alpha2_user_change_proto_rawDescOnce.Do(func() {
		edgelq_iam_proto_v1alpha2_user_change_proto_rawDescData = protoimpl.X.CompressGZIP(edgelq_iam_proto_v1alpha2_user_change_proto_rawDescData)
	})
	return edgelq_iam_proto_v1alpha2_user_change_proto_rawDescData
}

var edgelq_iam_proto_v1alpha2_user_change_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var edgelq_iam_proto_v1alpha2_user_change_proto_goTypes = []interface{}{
	(*UserChange)(nil),          // 0: ntt.iam.v1alpha2.UserChange
	(*UserChange_Added)(nil),    // 1: ntt.iam.v1alpha2.UserChange.Added
	(*UserChange_Modified)(nil), // 2: ntt.iam.v1alpha2.UserChange.Modified
	(*UserChange_Current)(nil),  // 3: ntt.iam.v1alpha2.UserChange.Current
	(*UserChange_Removed)(nil),  // 4: ntt.iam.v1alpha2.UserChange.Removed
	(*User)(nil),                // 5: ntt.iam.v1alpha2.User
	(*User_FieldMask)(nil),      // 6: ntt.iam.v1alpha2.User_FieldMask
}
var edgelq_iam_proto_v1alpha2_user_change_proto_depIdxs = []int32{
	1, // 0: ntt.iam.v1alpha2.UserChange.added:type_name -> ntt.iam.v1alpha2.UserChange.Added
	2, // 1: ntt.iam.v1alpha2.UserChange.modified:type_name -> ntt.iam.v1alpha2.UserChange.Modified
	3, // 2: ntt.iam.v1alpha2.UserChange.current:type_name -> ntt.iam.v1alpha2.UserChange.Current
	4, // 3: ntt.iam.v1alpha2.UserChange.removed:type_name -> ntt.iam.v1alpha2.UserChange.Removed
	5, // 4: ntt.iam.v1alpha2.UserChange.Added.user:type_name -> ntt.iam.v1alpha2.User
	5, // 5: ntt.iam.v1alpha2.UserChange.Modified.user:type_name -> ntt.iam.v1alpha2.User
	6, // 6: ntt.iam.v1alpha2.UserChange.Modified.field_mask:type_name -> ntt.iam.v1alpha2.User_FieldMask
	5, // 7: ntt.iam.v1alpha2.UserChange.Current.user:type_name -> ntt.iam.v1alpha2.User
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { edgelq_iam_proto_v1alpha2_user_change_proto_init() }
func edgelq_iam_proto_v1alpha2_user_change_proto_init() {
	if edgelq_iam_proto_v1alpha2_user_change_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		edgelq_iam_proto_v1alpha2_user_change_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserChange); i {
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
		edgelq_iam_proto_v1alpha2_user_change_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserChange_Added); i {
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
		edgelq_iam_proto_v1alpha2_user_change_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserChange_Modified); i {
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
		edgelq_iam_proto_v1alpha2_user_change_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserChange_Current); i {
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
		edgelq_iam_proto_v1alpha2_user_change_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserChange_Removed); i {
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

	edgelq_iam_proto_v1alpha2_user_change_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*UserChange_Added_)(nil),
		(*UserChange_Modified_)(nil),
		(*UserChange_Current_)(nil),
		(*UserChange_Removed_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: edgelq_iam_proto_v1alpha2_user_change_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           edgelq_iam_proto_v1alpha2_user_change_proto_goTypes,
		DependencyIndexes: edgelq_iam_proto_v1alpha2_user_change_proto_depIdxs,
		MessageInfos:      edgelq_iam_proto_v1alpha2_user_change_proto_msgTypes,
	}.Build()
	edgelq_iam_proto_v1alpha2_user_change_proto = out.File
	edgelq_iam_proto_v1alpha2_user_change_proto_rawDesc = nil
	edgelq_iam_proto_v1alpha2_user_change_proto_goTypes = nil
	edgelq_iam_proto_v1alpha2_user_change_proto_depIdxs = nil
}
