// Code generated by protoc-gen-goten-go
// File: edgelq/alerting/proto/v1/ts_entry_change.proto
// DO NOT EDIT!!!

package ts_entry

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
	ts_condition "github.com/cloudwan/edgelq-sdk/alerting/resources/v1/ts_condition"
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
	_ = &ts_condition.TsCondition{}
	_ = &fieldmaskpb.FieldMask{}
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// TsEntryChange is used by Watch notifications Responses to describe change of
// single TsEntry One of Added, Modified, Removed
type TsEntryChange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// TsEntry change
	//
	// Types that are valid to be assigned to ChangeType:
	//	*TsEntryChange_Added_
	//	*TsEntryChange_Modified_
	//	*TsEntryChange_Current_
	//	*TsEntryChange_Removed_
	ChangeType isTsEntryChange_ChangeType `protobuf_oneof:"change_type"`
}

func (m *TsEntryChange) Reset() {
	*m = TsEntryChange{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_alerting_proto_v1_ts_entry_change_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *TsEntryChange) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*TsEntryChange) ProtoMessage() {}

func (m *TsEntryChange) ProtoReflect() preflect.Message {
	mi := &edgelq_alerting_proto_v1_ts_entry_change_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*TsEntryChange) GotenMessage() {}

// Deprecated, Use TsEntryChange.ProtoReflect.Descriptor instead.
func (*TsEntryChange) Descriptor() ([]byte, []int) {
	return edgelq_alerting_proto_v1_ts_entry_change_proto_rawDescGZIP(), []int{0}
}

func (m *TsEntryChange) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *TsEntryChange) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *TsEntryChange) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *TsEntryChange) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

type isTsEntryChange_ChangeType interface {
	isTsEntryChange_ChangeType()
}

type TsEntryChange_Added_ struct {
	// Added is returned when watched document is added, either created or
	// enters Query view
	Added *TsEntryChange_Added `protobuf:"bytes,1,opt,name=added,proto3,oneof"`
}
type TsEntryChange_Modified_ struct {
	// Modified is returned when watched document is modified
	Modified *TsEntryChange_Modified `protobuf:"bytes,2,opt,name=modified,proto3,oneof"`
}
type TsEntryChange_Current_ struct {
	// Current is returned in stateless watch when document enters query view or
	// is modified within.
	Current *TsEntryChange_Current `protobuf:"bytes,4,opt,name=current,proto3,oneof"`
}
type TsEntryChange_Removed_ struct {
	// Removed is returned when TsEntry is deleted or leaves Query view
	Removed *TsEntryChange_Removed `protobuf:"bytes,3,opt,name=removed,proto3,oneof"`
}

func (*TsEntryChange_Added_) isTsEntryChange_ChangeType()    {}
func (*TsEntryChange_Modified_) isTsEntryChange_ChangeType() {}
func (*TsEntryChange_Current_) isTsEntryChange_ChangeType()  {}
func (*TsEntryChange_Removed_) isTsEntryChange_ChangeType()  {}
func (m *TsEntryChange) GetChangeType() isTsEntryChange_ChangeType {
	if m != nil {
		return m.ChangeType
	}
	return nil
}
func (m *TsEntryChange) GetAdded() *TsEntryChange_Added {
	if x, ok := m.GetChangeType().(*TsEntryChange_Added_); ok {
		return x.Added
	}
	return nil
}
func (m *TsEntryChange) GetModified() *TsEntryChange_Modified {
	if x, ok := m.GetChangeType().(*TsEntryChange_Modified_); ok {
		return x.Modified
	}
	return nil
}
func (m *TsEntryChange) GetCurrent() *TsEntryChange_Current {
	if x, ok := m.GetChangeType().(*TsEntryChange_Current_); ok {
		return x.Current
	}
	return nil
}
func (m *TsEntryChange) GetRemoved() *TsEntryChange_Removed {
	if x, ok := m.GetChangeType().(*TsEntryChange_Removed_); ok {
		return x.Removed
	}
	return nil
}
func (m *TsEntryChange) SetChangeType(ofv isTsEntryChange_ChangeType) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "isTsEntryChange_ChangeType", "TsEntryChange"))
	}
	m.ChangeType = ofv
}
func (m *TsEntryChange) SetAdded(fv *TsEntryChange_Added) {
	m.SetChangeType(&TsEntryChange_Added_{Added: fv})
}
func (m *TsEntryChange) SetModified(fv *TsEntryChange_Modified) {
	m.SetChangeType(&TsEntryChange_Modified_{Modified: fv})
}
func (m *TsEntryChange) SetCurrent(fv *TsEntryChange_Current) {
	m.SetChangeType(&TsEntryChange_Current_{Current: fv})
}
func (m *TsEntryChange) SetRemoved(fv *TsEntryChange_Removed) {
	m.SetChangeType(&TsEntryChange_Removed_{Removed: fv})
}

// TsEntry has been added to query view
type TsEntryChange_Added struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	TsEntry       *TsEntry `protobuf:"bytes,1,opt,name=ts_entry,json=tsEntry,proto3" json:"ts_entry,omitempty"`
	// Integer describing index of added TsEntry in resulting query view.
	ViewIndex int32 `protobuf:"varint,2,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty"`
}

func (m *TsEntryChange_Added) Reset() {
	*m = TsEntryChange_Added{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_alerting_proto_v1_ts_entry_change_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *TsEntryChange_Added) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*TsEntryChange_Added) ProtoMessage() {}

func (m *TsEntryChange_Added) ProtoReflect() preflect.Message {
	mi := &edgelq_alerting_proto_v1_ts_entry_change_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*TsEntryChange_Added) GotenMessage() {}

// Deprecated, Use TsEntryChange_Added.ProtoReflect.Descriptor instead.
func (*TsEntryChange_Added) Descriptor() ([]byte, []int) {
	return edgelq_alerting_proto_v1_ts_entry_change_proto_rawDescGZIP(), []int{0, 0}
}

func (m *TsEntryChange_Added) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *TsEntryChange_Added) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *TsEntryChange_Added) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *TsEntryChange_Added) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *TsEntryChange_Added) GetTsEntry() *TsEntry {
	if m != nil {
		return m.TsEntry
	}
	return nil
}

func (m *TsEntryChange_Added) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *TsEntryChange_Added) SetTsEntry(fv *TsEntry) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "TsEntry", "TsEntryChange_Added"))
	}
	m.TsEntry = fv
}

func (m *TsEntryChange_Added) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "TsEntryChange_Added"))
	}
	m.ViewIndex = fv
}

// TsEntry changed some of it's fields - contains either full document or
// masked change
type TsEntryChange_Modified struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Name of modified TsEntry
	Name *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty"`
	// New version of TsEntry or masked difference, depending on mask_changes
	// instrumentation of issued [WatchTsEntryRequest] or
	// [WatchTsEntriesRequest]
	TsEntry *TsEntry `protobuf:"bytes,2,opt,name=ts_entry,json=tsEntry,proto3" json:"ts_entry,omitempty"`
	// Used when mask_changes is set, contains field paths of modified
	// properties.
	FieldMask *TsEntry_FieldMask `protobuf:"bytes,3,opt,customtype=TsEntry_FieldMask,name=field_mask,json=fieldMask,proto3" json:"field_mask,omitempty"`
	// Previous view index specifies previous position of modified TsEntry.
	// When modification doesn't affect sorted order, value will remain
	// identical to [view_index].
	PreviousViewIndex int32 `protobuf:"varint,4,opt,name=previous_view_index,json=previousViewIndex,proto3" json:"previous_view_index,omitempty"`
	// Integer specifying TsEntry new index in resulting query view.
	ViewIndex int32 `protobuf:"varint,5,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty"`
}

func (m *TsEntryChange_Modified) Reset() {
	*m = TsEntryChange_Modified{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_alerting_proto_v1_ts_entry_change_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *TsEntryChange_Modified) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*TsEntryChange_Modified) ProtoMessage() {}

func (m *TsEntryChange_Modified) ProtoReflect() preflect.Message {
	mi := &edgelq_alerting_proto_v1_ts_entry_change_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*TsEntryChange_Modified) GotenMessage() {}

// Deprecated, Use TsEntryChange_Modified.ProtoReflect.Descriptor instead.
func (*TsEntryChange_Modified) Descriptor() ([]byte, []int) {
	return edgelq_alerting_proto_v1_ts_entry_change_proto_rawDescGZIP(), []int{0, 1}
}

func (m *TsEntryChange_Modified) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *TsEntryChange_Modified) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *TsEntryChange_Modified) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *TsEntryChange_Modified) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *TsEntryChange_Modified) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *TsEntryChange_Modified) GetTsEntry() *TsEntry {
	if m != nil {
		return m.TsEntry
	}
	return nil
}

func (m *TsEntryChange_Modified) GetFieldMask() *TsEntry_FieldMask {
	if m != nil {
		return m.FieldMask
	}
	return nil
}

func (m *TsEntryChange_Modified) GetPreviousViewIndex() int32 {
	if m != nil {
		return m.PreviousViewIndex
	}
	return int32(0)
}

func (m *TsEntryChange_Modified) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *TsEntryChange_Modified) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "TsEntryChange_Modified"))
	}
	m.Name = fv
}

func (m *TsEntryChange_Modified) SetTsEntry(fv *TsEntry) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "TsEntry", "TsEntryChange_Modified"))
	}
	m.TsEntry = fv
}

func (m *TsEntryChange_Modified) SetFieldMask(fv *TsEntry_FieldMask) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "FieldMask", "TsEntryChange_Modified"))
	}
	m.FieldMask = fv
}

func (m *TsEntryChange_Modified) SetPreviousViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "PreviousViewIndex", "TsEntryChange_Modified"))
	}
	m.PreviousViewIndex = fv
}

func (m *TsEntryChange_Modified) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "TsEntryChange_Modified"))
	}
	m.ViewIndex = fv
}

// TsEntry has been added or modified in a query view. Version used for
// stateless watching
type TsEntryChange_Current struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	TsEntry       *TsEntry `protobuf:"bytes,1,opt,name=ts_entry,json=tsEntry,proto3" json:"ts_entry,omitempty"`
}

func (m *TsEntryChange_Current) Reset() {
	*m = TsEntryChange_Current{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_alerting_proto_v1_ts_entry_change_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *TsEntryChange_Current) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*TsEntryChange_Current) ProtoMessage() {}

func (m *TsEntryChange_Current) ProtoReflect() preflect.Message {
	mi := &edgelq_alerting_proto_v1_ts_entry_change_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*TsEntryChange_Current) GotenMessage() {}

// Deprecated, Use TsEntryChange_Current.ProtoReflect.Descriptor instead.
func (*TsEntryChange_Current) Descriptor() ([]byte, []int) {
	return edgelq_alerting_proto_v1_ts_entry_change_proto_rawDescGZIP(), []int{0, 2}
}

func (m *TsEntryChange_Current) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *TsEntryChange_Current) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *TsEntryChange_Current) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *TsEntryChange_Current) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *TsEntryChange_Current) GetTsEntry() *TsEntry {
	if m != nil {
		return m.TsEntry
	}
	return nil
}

func (m *TsEntryChange_Current) SetTsEntry(fv *TsEntry) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "TsEntry", "TsEntryChange_Current"))
	}
	m.TsEntry = fv
}

// Removed is returned when TsEntry is deleted or leaves Query view
type TsEntryChange_Removed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Name          *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty"`
	// Integer specifying removed TsEntry index. Not populated in stateless
	// watch type.
	ViewIndex int32 `protobuf:"varint,2,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty"`
}

func (m *TsEntryChange_Removed) Reset() {
	*m = TsEntryChange_Removed{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_alerting_proto_v1_ts_entry_change_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *TsEntryChange_Removed) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*TsEntryChange_Removed) ProtoMessage() {}

func (m *TsEntryChange_Removed) ProtoReflect() preflect.Message {
	mi := &edgelq_alerting_proto_v1_ts_entry_change_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*TsEntryChange_Removed) GotenMessage() {}

// Deprecated, Use TsEntryChange_Removed.ProtoReflect.Descriptor instead.
func (*TsEntryChange_Removed) Descriptor() ([]byte, []int) {
	return edgelq_alerting_proto_v1_ts_entry_change_proto_rawDescGZIP(), []int{0, 3}
}

func (m *TsEntryChange_Removed) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *TsEntryChange_Removed) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *TsEntryChange_Removed) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *TsEntryChange_Removed) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *TsEntryChange_Removed) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *TsEntryChange_Removed) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *TsEntryChange_Removed) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "TsEntryChange_Removed"))
	}
	m.Name = fv
}

func (m *TsEntryChange_Removed) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "TsEntryChange_Removed"))
	}
	m.ViewIndex = fv
}

var edgelq_alerting_proto_v1_ts_entry_change_proto preflect.FileDescriptor

var edgelq_alerting_proto_v1_ts_entry_change_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e,
	0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x73, 0x5f, 0x65, 0x6e,
	0x74, 0x72, 0x79, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0f, 0x6e, 0x74, 0x74, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76,
	0x31, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d,
	0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67,
	0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x27, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x73, 0x5f, 0x65, 0x6e, 0x74,
	0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa8, 0x06, 0x0a, 0x0d, 0x54, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x3c, 0x0a, 0x05, 0x61, 0x64,
	0x64, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6e, 0x74, 0x74, 0x2e,
	0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x65, 0x64, 0x48,
	0x00, 0x52, 0x05, 0x61, 0x64, 0x64, 0x65, 0x64, 0x12, 0x45, 0x0a, 0x08, 0x6d, 0x6f, 0x64, 0x69,
	0x66, 0x69, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x6e, 0x74, 0x74,
	0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66,
	0x69, 0x65, 0x64, 0x48, 0x00, 0x52, 0x08, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12,
	0x42, 0x0a, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x26, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x07, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x12, 0x42, 0x0a, 0x07, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74,
	0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x48, 0x00, 0x52, 0x07,
	0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x1a, 0x5b, 0x0a, 0x05, 0x41, 0x64, 0x64, 0x65, 0x64,
	0x12, 0x33, 0x0a, 0x08, 0x74, 0x73, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x74, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x76, 0x69, 0x65, 0x77, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x1a, 0xff, 0x01, 0x0a, 0x08, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65,
	0x64, 0x12, 0x23, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x0f, 0xb2, 0xda, 0x21, 0x0b, 0x0a, 0x09, 0x0a, 0x07, 0x54, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x08, 0x74, 0x73, 0x5f, 0x65, 0x6e, 0x74,
	0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x61,
	0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x07, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x4a, 0x0a, 0x0a, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x42, 0x0f, 0xb2, 0xda, 0x21,
	0x0b, 0x32, 0x09, 0x0a, 0x07, 0x54, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x2e, 0x0a, 0x13, 0x70, 0x72, 0x65, 0x76, 0x69,
	0x6f, 0x75, 0x73, 0x5f, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x56, 0x69,
	0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x69, 0x65, 0x77, 0x5f,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x76, 0x69, 0x65,
	0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x1a, 0x3e, 0x0a, 0x07, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x12, 0x33, 0x0a, 0x08, 0x74, 0x73, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69,
	0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x74,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x1a, 0x4d, 0x0a, 0x07, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x64, 0x12, 0x23, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x0f, 0xb2, 0xda, 0x21, 0x0b, 0x0a, 0x09, 0x0a, 0x07, 0x54, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x76, 0x69, 0x65, 0x77,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x3a, 0x0d, 0x9a, 0xd9, 0x21, 0x09, 0x0a, 0x07, 0x54, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x42, 0x0d, 0x0a, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x42, 0x76, 0xe8, 0xde, 0x21, 0x00, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x6e,
	0x74, 0x74, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x62, 0x2e, 0x76,
	0x31, 0x42, 0x12, 0x54, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x00, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67,
	0x65, 0x6c, 0x71, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x73, 0x5f, 0x65, 0x6e, 0x74,
	0x72, 0x79, 0x3b, 0x74, 0x73, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	edgelq_alerting_proto_v1_ts_entry_change_proto_rawDescOnce sync.Once
	edgelq_alerting_proto_v1_ts_entry_change_proto_rawDescData = edgelq_alerting_proto_v1_ts_entry_change_proto_rawDesc
)

func edgelq_alerting_proto_v1_ts_entry_change_proto_rawDescGZIP() []byte {
	edgelq_alerting_proto_v1_ts_entry_change_proto_rawDescOnce.Do(func() {
		edgelq_alerting_proto_v1_ts_entry_change_proto_rawDescData = protoimpl.X.CompressGZIP(edgelq_alerting_proto_v1_ts_entry_change_proto_rawDescData)
	})
	return edgelq_alerting_proto_v1_ts_entry_change_proto_rawDescData
}

var edgelq_alerting_proto_v1_ts_entry_change_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var edgelq_alerting_proto_v1_ts_entry_change_proto_goTypes = []interface{}{
	(*TsEntryChange)(nil),          // 0: ntt.alerting.v1.TsEntryChange
	(*TsEntryChange_Added)(nil),    // 1: ntt.alerting.v1.TsEntryChange.Added
	(*TsEntryChange_Modified)(nil), // 2: ntt.alerting.v1.TsEntryChange.Modified
	(*TsEntryChange_Current)(nil),  // 3: ntt.alerting.v1.TsEntryChange.Current
	(*TsEntryChange_Removed)(nil),  // 4: ntt.alerting.v1.TsEntryChange.Removed
	(*TsEntry)(nil),                // 5: ntt.alerting.v1.TsEntry
	(*TsEntry_FieldMask)(nil),      // 6: ntt.alerting.v1.TsEntry_FieldMask
}
var edgelq_alerting_proto_v1_ts_entry_change_proto_depIdxs = []int32{
	1, // 0: ntt.alerting.v1.TsEntryChange.added:type_name -> ntt.alerting.v1.TsEntryChange.Added
	2, // 1: ntt.alerting.v1.TsEntryChange.modified:type_name -> ntt.alerting.v1.TsEntryChange.Modified
	3, // 2: ntt.alerting.v1.TsEntryChange.current:type_name -> ntt.alerting.v1.TsEntryChange.Current
	4, // 3: ntt.alerting.v1.TsEntryChange.removed:type_name -> ntt.alerting.v1.TsEntryChange.Removed
	5, // 4: ntt.alerting.v1.TsEntryChange.Added.ts_entry:type_name -> ntt.alerting.v1.TsEntry
	5, // 5: ntt.alerting.v1.TsEntryChange.Modified.ts_entry:type_name -> ntt.alerting.v1.TsEntry
	6, // 6: ntt.alerting.v1.TsEntryChange.Modified.field_mask:type_name -> ntt.alerting.v1.TsEntry_FieldMask
	5, // 7: ntt.alerting.v1.TsEntryChange.Current.ts_entry:type_name -> ntt.alerting.v1.TsEntry
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { edgelq_alerting_proto_v1_ts_entry_change_proto_init() }
func edgelq_alerting_proto_v1_ts_entry_change_proto_init() {
	if edgelq_alerting_proto_v1_ts_entry_change_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		edgelq_alerting_proto_v1_ts_entry_change_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TsEntryChange); i {
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
		edgelq_alerting_proto_v1_ts_entry_change_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TsEntryChange_Added); i {
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
		edgelq_alerting_proto_v1_ts_entry_change_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TsEntryChange_Modified); i {
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
		edgelq_alerting_proto_v1_ts_entry_change_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TsEntryChange_Current); i {
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
		edgelq_alerting_proto_v1_ts_entry_change_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TsEntryChange_Removed); i {
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

	edgelq_alerting_proto_v1_ts_entry_change_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*TsEntryChange_Added_)(nil),
		(*TsEntryChange_Modified_)(nil),
		(*TsEntryChange_Current_)(nil),
		(*TsEntryChange_Removed_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: edgelq_alerting_proto_v1_ts_entry_change_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           edgelq_alerting_proto_v1_ts_entry_change_proto_goTypes,
		DependencyIndexes: edgelq_alerting_proto_v1_ts_entry_change_proto_depIdxs,
		MessageInfos:      edgelq_alerting_proto_v1_ts_entry_change_proto_msgTypes,
	}.Build()
	edgelq_alerting_proto_v1_ts_entry_change_proto = out.File
	edgelq_alerting_proto_v1_ts_entry_change_proto_rawDesc = nil
	edgelq_alerting_proto_v1_ts_entry_change_proto_goTypes = nil
	edgelq_alerting_proto_v1_ts_entry_change_proto_depIdxs = nil
}
