// Code generated by protoc-gen-goten-go
// File: edgelq/ztp/proto/v1/edgelq_instance_change.proto
// DO NOT EDIT!!!

package edgelq_instance

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
	project "github.com/cloudwan/edgelq-sdk/ztp/resources/v1/project"
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

// EdgelqInstanceChange is used by Watch notifications Responses to describe
// change of single EdgelqInstance One of Added, Modified, Removed
type EdgelqInstanceChange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// EdgelqInstance change
	//
	// Types that are valid to be assigned to ChangeType:
	//	*EdgelqInstanceChange_Added_
	//	*EdgelqInstanceChange_Modified_
	//	*EdgelqInstanceChange_Current_
	//	*EdgelqInstanceChange_Removed_
	ChangeType isEdgelqInstanceChange_ChangeType `protobuf_oneof:"change_type"`
}

func (m *EdgelqInstanceChange) Reset() {
	*m = EdgelqInstanceChange{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_ztp_proto_v1_edgelq_instance_change_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *EdgelqInstanceChange) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*EdgelqInstanceChange) ProtoMessage() {}

func (m *EdgelqInstanceChange) ProtoReflect() preflect.Message {
	mi := &edgelq_ztp_proto_v1_edgelq_instance_change_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*EdgelqInstanceChange) GotenMessage() {}

// Deprecated, Use EdgelqInstanceChange.ProtoReflect.Descriptor instead.
func (*EdgelqInstanceChange) Descriptor() ([]byte, []int) {
	return edgelq_ztp_proto_v1_edgelq_instance_change_proto_rawDescGZIP(), []int{0}
}

func (m *EdgelqInstanceChange) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *EdgelqInstanceChange) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *EdgelqInstanceChange) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *EdgelqInstanceChange) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

type isEdgelqInstanceChange_ChangeType interface {
	isEdgelqInstanceChange_ChangeType()
}

type EdgelqInstanceChange_Added_ struct {
	// Added is returned when watched document is added, either created or
	// enters Query view
	Added *EdgelqInstanceChange_Added `protobuf:"bytes,1,opt,name=added,proto3,oneof" firestore:"added"`
}
type EdgelqInstanceChange_Modified_ struct {
	// Modified is returned when watched document is modified
	Modified *EdgelqInstanceChange_Modified `protobuf:"bytes,2,opt,name=modified,proto3,oneof" firestore:"modified"`
}
type EdgelqInstanceChange_Current_ struct {
	// Current is returned in stateless watch when document enters query view or
	// is modified within.
	Current *EdgelqInstanceChange_Current `protobuf:"bytes,4,opt,name=current,proto3,oneof" firestore:"current"`
}
type EdgelqInstanceChange_Removed_ struct {
	// Removed is returned when EdgelqInstance is deleted or leaves Query view
	Removed *EdgelqInstanceChange_Removed `protobuf:"bytes,3,opt,name=removed,proto3,oneof" firestore:"removed"`
}

func (*EdgelqInstanceChange_Added_) isEdgelqInstanceChange_ChangeType()    {}
func (*EdgelqInstanceChange_Modified_) isEdgelqInstanceChange_ChangeType() {}
func (*EdgelqInstanceChange_Current_) isEdgelqInstanceChange_ChangeType()  {}
func (*EdgelqInstanceChange_Removed_) isEdgelqInstanceChange_ChangeType()  {}
func (m *EdgelqInstanceChange) GetChangeType() isEdgelqInstanceChange_ChangeType {
	if m != nil {
		return m.ChangeType
	}
	return nil
}
func (m *EdgelqInstanceChange) GetAdded() *EdgelqInstanceChange_Added {
	if x, ok := m.GetChangeType().(*EdgelqInstanceChange_Added_); ok {
		return x.Added
	}
	return nil
}
func (m *EdgelqInstanceChange) GetModified() *EdgelqInstanceChange_Modified {
	if x, ok := m.GetChangeType().(*EdgelqInstanceChange_Modified_); ok {
		return x.Modified
	}
	return nil
}
func (m *EdgelqInstanceChange) GetCurrent() *EdgelqInstanceChange_Current {
	if x, ok := m.GetChangeType().(*EdgelqInstanceChange_Current_); ok {
		return x.Current
	}
	return nil
}
func (m *EdgelqInstanceChange) GetRemoved() *EdgelqInstanceChange_Removed {
	if x, ok := m.GetChangeType().(*EdgelqInstanceChange_Removed_); ok {
		return x.Removed
	}
	return nil
}
func (m *EdgelqInstanceChange) SetChangeType(ofv isEdgelqInstanceChange_ChangeType) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "isEdgelqInstanceChange_ChangeType", "EdgelqInstanceChange"))
	}
	m.ChangeType = ofv
}
func (m *EdgelqInstanceChange) SetAdded(fv *EdgelqInstanceChange_Added) {
	m.SetChangeType(&EdgelqInstanceChange_Added_{Added: fv})
}
func (m *EdgelqInstanceChange) SetModified(fv *EdgelqInstanceChange_Modified) {
	m.SetChangeType(&EdgelqInstanceChange_Modified_{Modified: fv})
}
func (m *EdgelqInstanceChange) SetCurrent(fv *EdgelqInstanceChange_Current) {
	m.SetChangeType(&EdgelqInstanceChange_Current_{Current: fv})
}
func (m *EdgelqInstanceChange) SetRemoved(fv *EdgelqInstanceChange_Removed) {
	m.SetChangeType(&EdgelqInstanceChange_Removed_{Removed: fv})
}

// EdgelqInstance has been added to query view
type EdgelqInstanceChange_Added struct {
	state          protoimpl.MessageState
	sizeCache      protoimpl.SizeCache
	unknownFields  protoimpl.UnknownFields
	EdgelqInstance *EdgelqInstance `protobuf:"bytes,1,opt,name=edgelq_instance,json=edgelqInstance,proto3" json:"edgelq_instance,omitempty" firestore:"edgelqInstance"`
	// Integer describing index of added EdgelqInstance in resulting query view.
	ViewIndex int32 `protobuf:"varint,2,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty" firestore:"viewIndex"`
}

func (m *EdgelqInstanceChange_Added) Reset() {
	*m = EdgelqInstanceChange_Added{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_ztp_proto_v1_edgelq_instance_change_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *EdgelqInstanceChange_Added) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*EdgelqInstanceChange_Added) ProtoMessage() {}

func (m *EdgelqInstanceChange_Added) ProtoReflect() preflect.Message {
	mi := &edgelq_ztp_proto_v1_edgelq_instance_change_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*EdgelqInstanceChange_Added) GotenMessage() {}

// Deprecated, Use EdgelqInstanceChange_Added.ProtoReflect.Descriptor instead.
func (*EdgelqInstanceChange_Added) Descriptor() ([]byte, []int) {
	return edgelq_ztp_proto_v1_edgelq_instance_change_proto_rawDescGZIP(), []int{0, 0}
}

func (m *EdgelqInstanceChange_Added) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *EdgelqInstanceChange_Added) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *EdgelqInstanceChange_Added) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *EdgelqInstanceChange_Added) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *EdgelqInstanceChange_Added) GetEdgelqInstance() *EdgelqInstance {
	if m != nil {
		return m.EdgelqInstance
	}
	return nil
}

func (m *EdgelqInstanceChange_Added) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *EdgelqInstanceChange_Added) SetEdgelqInstance(fv *EdgelqInstance) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "EdgelqInstance", "EdgelqInstanceChange_Added"))
	}
	m.EdgelqInstance = fv
}

func (m *EdgelqInstanceChange_Added) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "EdgelqInstanceChange_Added"))
	}
	m.ViewIndex = fv
}

// EdgelqInstance changed some of it's fields - contains either full document
// or masked change
type EdgelqInstanceChange_Modified struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Name of modified EdgelqInstance
	Name *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// New version of EdgelqInstance or masked difference, depending on
	// mask_changes instrumentation of issued [WatchEdgelqInstanceRequest] or
	// [WatchEdgelqInstancesRequest]
	EdgelqInstance *EdgelqInstance `protobuf:"bytes,2,opt,name=edgelq_instance,json=edgelqInstance,proto3" json:"edgelq_instance,omitempty" firestore:"edgelqInstance"`
	// Used when mask_changes is set, contains field paths of modified
	// properties.
	FieldMask *EdgelqInstance_FieldMask `protobuf:"bytes,3,opt,customtype=EdgelqInstance_FieldMask,name=field_mask,json=fieldMask,proto3" json:"field_mask,omitempty" firestore:"fieldMask"`
	// Previous view index specifies previous position of modified
	// EdgelqInstance. When modification doesn't affect sorted order, value will
	// remain identical to [view_index].
	PreviousViewIndex int32 `protobuf:"varint,4,opt,name=previous_view_index,json=previousViewIndex,proto3" json:"previous_view_index,omitempty" firestore:"previousViewIndex"`
	// Integer specifying EdgelqInstance new index in resulting query view.
	ViewIndex int32 `protobuf:"varint,5,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty" firestore:"viewIndex"`
}

func (m *EdgelqInstanceChange_Modified) Reset() {
	*m = EdgelqInstanceChange_Modified{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_ztp_proto_v1_edgelq_instance_change_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *EdgelqInstanceChange_Modified) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*EdgelqInstanceChange_Modified) ProtoMessage() {}

func (m *EdgelqInstanceChange_Modified) ProtoReflect() preflect.Message {
	mi := &edgelq_ztp_proto_v1_edgelq_instance_change_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*EdgelqInstanceChange_Modified) GotenMessage() {}

// Deprecated, Use EdgelqInstanceChange_Modified.ProtoReflect.Descriptor instead.
func (*EdgelqInstanceChange_Modified) Descriptor() ([]byte, []int) {
	return edgelq_ztp_proto_v1_edgelq_instance_change_proto_rawDescGZIP(), []int{0, 1}
}

func (m *EdgelqInstanceChange_Modified) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *EdgelqInstanceChange_Modified) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *EdgelqInstanceChange_Modified) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *EdgelqInstanceChange_Modified) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *EdgelqInstanceChange_Modified) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *EdgelqInstanceChange_Modified) GetEdgelqInstance() *EdgelqInstance {
	if m != nil {
		return m.EdgelqInstance
	}
	return nil
}

func (m *EdgelqInstanceChange_Modified) GetFieldMask() *EdgelqInstance_FieldMask {
	if m != nil {
		return m.FieldMask
	}
	return nil
}

func (m *EdgelqInstanceChange_Modified) GetPreviousViewIndex() int32 {
	if m != nil {
		return m.PreviousViewIndex
	}
	return int32(0)
}

func (m *EdgelqInstanceChange_Modified) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *EdgelqInstanceChange_Modified) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "EdgelqInstanceChange_Modified"))
	}
	m.Name = fv
}

func (m *EdgelqInstanceChange_Modified) SetEdgelqInstance(fv *EdgelqInstance) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "EdgelqInstance", "EdgelqInstanceChange_Modified"))
	}
	m.EdgelqInstance = fv
}

func (m *EdgelqInstanceChange_Modified) SetFieldMask(fv *EdgelqInstance_FieldMask) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "FieldMask", "EdgelqInstanceChange_Modified"))
	}
	m.FieldMask = fv
}

func (m *EdgelqInstanceChange_Modified) SetPreviousViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "PreviousViewIndex", "EdgelqInstanceChange_Modified"))
	}
	m.PreviousViewIndex = fv
}

func (m *EdgelqInstanceChange_Modified) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "EdgelqInstanceChange_Modified"))
	}
	m.ViewIndex = fv
}

// EdgelqInstance has been added or modified in a query view. Version used for
// stateless watching
type EdgelqInstanceChange_Current struct {
	state          protoimpl.MessageState
	sizeCache      protoimpl.SizeCache
	unknownFields  protoimpl.UnknownFields
	EdgelqInstance *EdgelqInstance `protobuf:"bytes,1,opt,name=edgelq_instance,json=edgelqInstance,proto3" json:"edgelq_instance,omitempty" firestore:"edgelqInstance"`
}

func (m *EdgelqInstanceChange_Current) Reset() {
	*m = EdgelqInstanceChange_Current{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_ztp_proto_v1_edgelq_instance_change_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *EdgelqInstanceChange_Current) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*EdgelqInstanceChange_Current) ProtoMessage() {}

func (m *EdgelqInstanceChange_Current) ProtoReflect() preflect.Message {
	mi := &edgelq_ztp_proto_v1_edgelq_instance_change_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*EdgelqInstanceChange_Current) GotenMessage() {}

// Deprecated, Use EdgelqInstanceChange_Current.ProtoReflect.Descriptor instead.
func (*EdgelqInstanceChange_Current) Descriptor() ([]byte, []int) {
	return edgelq_ztp_proto_v1_edgelq_instance_change_proto_rawDescGZIP(), []int{0, 2}
}

func (m *EdgelqInstanceChange_Current) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *EdgelqInstanceChange_Current) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *EdgelqInstanceChange_Current) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *EdgelqInstanceChange_Current) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *EdgelqInstanceChange_Current) GetEdgelqInstance() *EdgelqInstance {
	if m != nil {
		return m.EdgelqInstance
	}
	return nil
}

func (m *EdgelqInstanceChange_Current) SetEdgelqInstance(fv *EdgelqInstance) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "EdgelqInstance", "EdgelqInstanceChange_Current"))
	}
	m.EdgelqInstance = fv
}

// Removed is returned when EdgelqInstance is deleted or leaves Query view
type EdgelqInstanceChange_Removed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Name          *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// Integer specifying removed EdgelqInstance index. Not populated in
	// stateless watch type.
	ViewIndex int32 `protobuf:"varint,2,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty" firestore:"viewIndex"`
}

func (m *EdgelqInstanceChange_Removed) Reset() {
	*m = EdgelqInstanceChange_Removed{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_ztp_proto_v1_edgelq_instance_change_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *EdgelqInstanceChange_Removed) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*EdgelqInstanceChange_Removed) ProtoMessage() {}

func (m *EdgelqInstanceChange_Removed) ProtoReflect() preflect.Message {
	mi := &edgelq_ztp_proto_v1_edgelq_instance_change_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*EdgelqInstanceChange_Removed) GotenMessage() {}

// Deprecated, Use EdgelqInstanceChange_Removed.ProtoReflect.Descriptor instead.
func (*EdgelqInstanceChange_Removed) Descriptor() ([]byte, []int) {
	return edgelq_ztp_proto_v1_edgelq_instance_change_proto_rawDescGZIP(), []int{0, 3}
}

func (m *EdgelqInstanceChange_Removed) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *EdgelqInstanceChange_Removed) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *EdgelqInstanceChange_Removed) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *EdgelqInstanceChange_Removed) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *EdgelqInstanceChange_Removed) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *EdgelqInstanceChange_Removed) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *EdgelqInstanceChange_Removed) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "EdgelqInstanceChange_Removed"))
	}
	m.Name = fv
}

func (m *EdgelqInstanceChange_Removed) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "EdgelqInstanceChange_Removed"))
	}
	m.ViewIndex = fv
}

var edgelq_ztp_proto_v1_edgelq_instance_change_proto preflect.FileDescriptor

var edgelq_ztp_proto_v1_edgelq_instance_change_proto_rawDesc = []byte{
	0x0a, 0x30, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x7a, 0x74, 0x70, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x5f, 0x69, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x6e, 0x74, 0x74, 0x2e, 0x7a, 0x74, 0x70, 0x2e, 0x76, 0x31, 0x1a, 0x19,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74,
	0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67,
	0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65,
	0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x65, 0x64,
	0x67, 0x65, 0x6c, 0x71, 0x2f, 0x7a, 0x74, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76,
	0x31, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x83, 0x07, 0x0a, 0x14, 0x45, 0x64, 0x67, 0x65,
	0x6c, 0x71, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x12, 0x3e, 0x0a, 0x05, 0x61, 0x64, 0x64, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x26, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x7a, 0x74, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x64, 0x67,
	0x65, 0x6c, 0x71, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x2e, 0x41, 0x64, 0x64, 0x65, 0x64, 0x48, 0x00, 0x52, 0x05, 0x61, 0x64, 0x64, 0x65, 0x64,
	0x12, 0x47, 0x0a, 0x08, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x29, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x7a, 0x74, 0x70, 0x2e, 0x76, 0x31, 0x2e,
	0x45, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x48, 0x00, 0x52,
	0x08, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x44, 0x0a, 0x07, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x6e, 0x74, 0x74,
	0x2e, 0x7a, 0x74, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x49, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12,
	0x44, 0x0a, 0x07, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x28, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x7a, 0x74, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x64,
	0x67, 0x65, 0x6c, 0x71, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x48, 0x00, 0x52, 0x07, 0x72, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x64, 0x1a, 0x6b, 0x0a, 0x05, 0x41, 0x64, 0x64, 0x65, 0x64, 0x12, 0x43,
	0x0a, 0x0f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x7a, 0x74,
	0x70, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x49, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x52, 0x0e, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x49, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x76, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x1a, 0x9d, 0x02, 0x0a, 0x08, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12,
	0x2a, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x16, 0xb2,
	0xda, 0x21, 0x12, 0x0a, 0x10, 0x0a, 0x0e, 0x45, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x43, 0x0a, 0x0f, 0x65,
	0x64, 0x67, 0x65, 0x6c, 0x71, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x7a, 0x74, 0x70, 0x2e, 0x76,
	0x31, 0x2e, 0x45, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x52, 0x0e, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x12, 0x51, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b,
	0x42, 0x16, 0xb2, 0xda, 0x21, 0x12, 0x32, 0x10, 0x0a, 0x0e, 0x45, 0x64, 0x67, 0x65, 0x6c, 0x71,
	0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4d,
	0x61, 0x73, 0x6b, 0x12, 0x2e, 0x0a, 0x13, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x5f,
	0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x11, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x56, 0x69, 0x65, 0x77, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x76, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x1a, 0x4e, 0x0a, 0x07, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x43, 0x0a,
	0x0f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x7a, 0x74, 0x70,
	0x2e, 0x76, 0x31, 0x2e, 0x45, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x52, 0x0e, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x1a, 0x54, 0x0a, 0x07, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x12, 0x2a, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x16, 0xb2, 0xda, 0x21,
	0x12, 0x0a, 0x10, 0x0a, 0x0e, 0x45, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x49, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x69, 0x65,
	0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x76,
	0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x3a, 0x14, 0x9a, 0xd9, 0x21, 0x10, 0x0a, 0x0e,
	0x45, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x42, 0x0d,
	0x0a, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x42, 0x81, 0x01,
	0xe8, 0xde, 0x21, 0x00, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x7a, 0x74,
	0x70, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x42, 0x19, 0x45, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x49,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x00, 0x5a, 0x4b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71,
	0x2f, 0x7a, 0x74, 0x70, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x3b, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	edgelq_ztp_proto_v1_edgelq_instance_change_proto_rawDescOnce sync.Once
	edgelq_ztp_proto_v1_edgelq_instance_change_proto_rawDescData = edgelq_ztp_proto_v1_edgelq_instance_change_proto_rawDesc
)

func edgelq_ztp_proto_v1_edgelq_instance_change_proto_rawDescGZIP() []byte {
	edgelq_ztp_proto_v1_edgelq_instance_change_proto_rawDescOnce.Do(func() {
		edgelq_ztp_proto_v1_edgelq_instance_change_proto_rawDescData = protoimpl.X.CompressGZIP(edgelq_ztp_proto_v1_edgelq_instance_change_proto_rawDescData)
	})
	return edgelq_ztp_proto_v1_edgelq_instance_change_proto_rawDescData
}

var edgelq_ztp_proto_v1_edgelq_instance_change_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var edgelq_ztp_proto_v1_edgelq_instance_change_proto_goTypes = []interface{}{
	(*EdgelqInstanceChange)(nil),          // 0: ntt.ztp.v1.EdgelqInstanceChange
	(*EdgelqInstanceChange_Added)(nil),    // 1: ntt.ztp.v1.EdgelqInstanceChange.Added
	(*EdgelqInstanceChange_Modified)(nil), // 2: ntt.ztp.v1.EdgelqInstanceChange.Modified
	(*EdgelqInstanceChange_Current)(nil),  // 3: ntt.ztp.v1.EdgelqInstanceChange.Current
	(*EdgelqInstanceChange_Removed)(nil),  // 4: ntt.ztp.v1.EdgelqInstanceChange.Removed
	(*EdgelqInstance)(nil),                // 5: ntt.ztp.v1.EdgelqInstance
	(*EdgelqInstance_FieldMask)(nil),      // 6: ntt.ztp.v1.EdgelqInstance_FieldMask
}
var edgelq_ztp_proto_v1_edgelq_instance_change_proto_depIdxs = []int32{
	1, // 0: ntt.ztp.v1.EdgelqInstanceChange.added:type_name -> ntt.ztp.v1.EdgelqInstanceChange.Added
	2, // 1: ntt.ztp.v1.EdgelqInstanceChange.modified:type_name -> ntt.ztp.v1.EdgelqInstanceChange.Modified
	3, // 2: ntt.ztp.v1.EdgelqInstanceChange.current:type_name -> ntt.ztp.v1.EdgelqInstanceChange.Current
	4, // 3: ntt.ztp.v1.EdgelqInstanceChange.removed:type_name -> ntt.ztp.v1.EdgelqInstanceChange.Removed
	5, // 4: ntt.ztp.v1.EdgelqInstanceChange.Added.edgelq_instance:type_name -> ntt.ztp.v1.EdgelqInstance
	5, // 5: ntt.ztp.v1.EdgelqInstanceChange.Modified.edgelq_instance:type_name -> ntt.ztp.v1.EdgelqInstance
	6, // 6: ntt.ztp.v1.EdgelqInstanceChange.Modified.field_mask:type_name -> ntt.ztp.v1.EdgelqInstance_FieldMask
	5, // 7: ntt.ztp.v1.EdgelqInstanceChange.Current.edgelq_instance:type_name -> ntt.ztp.v1.EdgelqInstance
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { edgelq_ztp_proto_v1_edgelq_instance_change_proto_init() }
func edgelq_ztp_proto_v1_edgelq_instance_change_proto_init() {
	if edgelq_ztp_proto_v1_edgelq_instance_change_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		edgelq_ztp_proto_v1_edgelq_instance_change_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EdgelqInstanceChange); i {
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
		edgelq_ztp_proto_v1_edgelq_instance_change_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EdgelqInstanceChange_Added); i {
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
		edgelq_ztp_proto_v1_edgelq_instance_change_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EdgelqInstanceChange_Modified); i {
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
		edgelq_ztp_proto_v1_edgelq_instance_change_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EdgelqInstanceChange_Current); i {
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
		edgelq_ztp_proto_v1_edgelq_instance_change_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EdgelqInstanceChange_Removed); i {
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

	edgelq_ztp_proto_v1_edgelq_instance_change_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*EdgelqInstanceChange_Added_)(nil),
		(*EdgelqInstanceChange_Modified_)(nil),
		(*EdgelqInstanceChange_Current_)(nil),
		(*EdgelqInstanceChange_Removed_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: edgelq_ztp_proto_v1_edgelq_instance_change_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           edgelq_ztp_proto_v1_edgelq_instance_change_proto_goTypes,
		DependencyIndexes: edgelq_ztp_proto_v1_edgelq_instance_change_proto_depIdxs,
		MessageInfos:      edgelq_ztp_proto_v1_edgelq_instance_change_proto_msgTypes,
	}.Build()
	edgelq_ztp_proto_v1_edgelq_instance_change_proto = out.File
	edgelq_ztp_proto_v1_edgelq_instance_change_proto_rawDesc = nil
	edgelq_ztp_proto_v1_edgelq_instance_change_proto_goTypes = nil
	edgelq_ztp_proto_v1_edgelq_instance_change_proto_depIdxs = nil
}