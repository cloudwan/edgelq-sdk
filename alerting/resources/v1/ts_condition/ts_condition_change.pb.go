// Code generated by protoc-gen-goten-go
// File: edgelq/alerting/proto/v1/ts_condition_change.proto
// DO NOT EDIT!!!

package ts_condition

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
	policy "github.com/cloudwan/edgelq-sdk/alerting/resources/v1/policy"
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
	_ = &policy.Policy{}
	_ = &fieldmaskpb.FieldMask{}
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// TsConditionChange is used by Watch notifications Responses to describe change
// of single TsCondition One of Added, Modified, Removed
type TsConditionChange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// TsCondition change
	//
	// Types that are valid to be assigned to ChangeType:
	//	*TsConditionChange_Added_
	//	*TsConditionChange_Modified_
	//	*TsConditionChange_Current_
	//	*TsConditionChange_Removed_
	ChangeType isTsConditionChange_ChangeType `protobuf_oneof:"change_type"`
}

func (m *TsConditionChange) Reset() {
	*m = TsConditionChange{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_alerting_proto_v1_ts_condition_change_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *TsConditionChange) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*TsConditionChange) ProtoMessage() {}

func (m *TsConditionChange) ProtoReflect() preflect.Message {
	mi := &edgelq_alerting_proto_v1_ts_condition_change_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*TsConditionChange) GotenMessage() {}

// Deprecated, Use TsConditionChange.ProtoReflect.Descriptor instead.
func (*TsConditionChange) Descriptor() ([]byte, []int) {
	return edgelq_alerting_proto_v1_ts_condition_change_proto_rawDescGZIP(), []int{0}
}

func (m *TsConditionChange) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *TsConditionChange) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *TsConditionChange) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *TsConditionChange) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

type isTsConditionChange_ChangeType interface {
	isTsConditionChange_ChangeType()
}

type TsConditionChange_Added_ struct {
	// Added is returned when watched document is added, either created or
	// enters Query view
	Added *TsConditionChange_Added `protobuf:"bytes,1,opt,name=added,proto3,oneof"`
}
type TsConditionChange_Modified_ struct {
	// Modified is returned when watched document is modified
	Modified *TsConditionChange_Modified `protobuf:"bytes,2,opt,name=modified,proto3,oneof"`
}
type TsConditionChange_Current_ struct {
	// Current is returned in stateless watch when document enters query view or
	// is modified within.
	Current *TsConditionChange_Current `protobuf:"bytes,4,opt,name=current,proto3,oneof"`
}
type TsConditionChange_Removed_ struct {
	// Removed is returned when TsCondition is deleted or leaves Query view
	Removed *TsConditionChange_Removed `protobuf:"bytes,3,opt,name=removed,proto3,oneof"`
}

func (*TsConditionChange_Added_) isTsConditionChange_ChangeType()    {}
func (*TsConditionChange_Modified_) isTsConditionChange_ChangeType() {}
func (*TsConditionChange_Current_) isTsConditionChange_ChangeType()  {}
func (*TsConditionChange_Removed_) isTsConditionChange_ChangeType()  {}
func (m *TsConditionChange) GetChangeType() isTsConditionChange_ChangeType {
	if m != nil {
		return m.ChangeType
	}
	return nil
}
func (m *TsConditionChange) GetAdded() *TsConditionChange_Added {
	if x, ok := m.GetChangeType().(*TsConditionChange_Added_); ok {
		return x.Added
	}
	return nil
}
func (m *TsConditionChange) GetModified() *TsConditionChange_Modified {
	if x, ok := m.GetChangeType().(*TsConditionChange_Modified_); ok {
		return x.Modified
	}
	return nil
}
func (m *TsConditionChange) GetCurrent() *TsConditionChange_Current {
	if x, ok := m.GetChangeType().(*TsConditionChange_Current_); ok {
		return x.Current
	}
	return nil
}
func (m *TsConditionChange) GetRemoved() *TsConditionChange_Removed {
	if x, ok := m.GetChangeType().(*TsConditionChange_Removed_); ok {
		return x.Removed
	}
	return nil
}
func (m *TsConditionChange) SetChangeType(ofv isTsConditionChange_ChangeType) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "isTsConditionChange_ChangeType", "TsConditionChange"))
	}
	m.ChangeType = ofv
}
func (m *TsConditionChange) SetAdded(fv *TsConditionChange_Added) {
	m.SetChangeType(&TsConditionChange_Added_{Added: fv})
}
func (m *TsConditionChange) SetModified(fv *TsConditionChange_Modified) {
	m.SetChangeType(&TsConditionChange_Modified_{Modified: fv})
}
func (m *TsConditionChange) SetCurrent(fv *TsConditionChange_Current) {
	m.SetChangeType(&TsConditionChange_Current_{Current: fv})
}
func (m *TsConditionChange) SetRemoved(fv *TsConditionChange_Removed) {
	m.SetChangeType(&TsConditionChange_Removed_{Removed: fv})
}

// TsCondition has been added to query view
type TsConditionChange_Added struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	TsCondition   *TsCondition `protobuf:"bytes,1,opt,name=ts_condition,json=tsCondition,proto3" json:"ts_condition,omitempty"`
	// Integer describing index of added TsCondition in resulting query view.
	ViewIndex int32 `protobuf:"varint,2,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty"`
}

func (m *TsConditionChange_Added) Reset() {
	*m = TsConditionChange_Added{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_alerting_proto_v1_ts_condition_change_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *TsConditionChange_Added) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*TsConditionChange_Added) ProtoMessage() {}

func (m *TsConditionChange_Added) ProtoReflect() preflect.Message {
	mi := &edgelq_alerting_proto_v1_ts_condition_change_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*TsConditionChange_Added) GotenMessage() {}

// Deprecated, Use TsConditionChange_Added.ProtoReflect.Descriptor instead.
func (*TsConditionChange_Added) Descriptor() ([]byte, []int) {
	return edgelq_alerting_proto_v1_ts_condition_change_proto_rawDescGZIP(), []int{0, 0}
}

func (m *TsConditionChange_Added) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *TsConditionChange_Added) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *TsConditionChange_Added) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *TsConditionChange_Added) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *TsConditionChange_Added) GetTsCondition() *TsCondition {
	if m != nil {
		return m.TsCondition
	}
	return nil
}

func (m *TsConditionChange_Added) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *TsConditionChange_Added) SetTsCondition(fv *TsCondition) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "TsCondition", "TsConditionChange_Added"))
	}
	m.TsCondition = fv
}

func (m *TsConditionChange_Added) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "TsConditionChange_Added"))
	}
	m.ViewIndex = fv
}

// TsCondition changed some of it's fields - contains either full document or
// masked change
type TsConditionChange_Modified struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Name of modified TsCondition
	Name *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty"`
	// New version of TsCondition or masked difference, depending on
	// mask_changes instrumentation of issued [WatchTsConditionRequest] or
	// [WatchTsConditionsRequest]
	TsCondition *TsCondition `protobuf:"bytes,2,opt,name=ts_condition,json=tsCondition,proto3" json:"ts_condition,omitempty"`
	// Used when mask_changes is set, contains field paths of modified
	// properties.
	FieldMask *TsCondition_FieldMask `protobuf:"bytes,3,opt,customtype=TsCondition_FieldMask,name=field_mask,json=fieldMask,proto3" json:"field_mask,omitempty"`
	// Previous view index specifies previous position of modified TsCondition.
	// When modification doesn't affect sorted order, value will remain
	// identical to [view_index].
	PreviousViewIndex int32 `protobuf:"varint,4,opt,name=previous_view_index,json=previousViewIndex,proto3" json:"previous_view_index,omitempty"`
	// Integer specifying TsCondition new index in resulting query view.
	ViewIndex int32 `protobuf:"varint,5,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty"`
}

func (m *TsConditionChange_Modified) Reset() {
	*m = TsConditionChange_Modified{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_alerting_proto_v1_ts_condition_change_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *TsConditionChange_Modified) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*TsConditionChange_Modified) ProtoMessage() {}

func (m *TsConditionChange_Modified) ProtoReflect() preflect.Message {
	mi := &edgelq_alerting_proto_v1_ts_condition_change_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*TsConditionChange_Modified) GotenMessage() {}

// Deprecated, Use TsConditionChange_Modified.ProtoReflect.Descriptor instead.
func (*TsConditionChange_Modified) Descriptor() ([]byte, []int) {
	return edgelq_alerting_proto_v1_ts_condition_change_proto_rawDescGZIP(), []int{0, 1}
}

func (m *TsConditionChange_Modified) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *TsConditionChange_Modified) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *TsConditionChange_Modified) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *TsConditionChange_Modified) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *TsConditionChange_Modified) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *TsConditionChange_Modified) GetTsCondition() *TsCondition {
	if m != nil {
		return m.TsCondition
	}
	return nil
}

func (m *TsConditionChange_Modified) GetFieldMask() *TsCondition_FieldMask {
	if m != nil {
		return m.FieldMask
	}
	return nil
}

func (m *TsConditionChange_Modified) GetPreviousViewIndex() int32 {
	if m != nil {
		return m.PreviousViewIndex
	}
	return int32(0)
}

func (m *TsConditionChange_Modified) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *TsConditionChange_Modified) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "TsConditionChange_Modified"))
	}
	m.Name = fv
}

func (m *TsConditionChange_Modified) SetTsCondition(fv *TsCondition) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "TsCondition", "TsConditionChange_Modified"))
	}
	m.TsCondition = fv
}

func (m *TsConditionChange_Modified) SetFieldMask(fv *TsCondition_FieldMask) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "FieldMask", "TsConditionChange_Modified"))
	}
	m.FieldMask = fv
}

func (m *TsConditionChange_Modified) SetPreviousViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "PreviousViewIndex", "TsConditionChange_Modified"))
	}
	m.PreviousViewIndex = fv
}

func (m *TsConditionChange_Modified) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "TsConditionChange_Modified"))
	}
	m.ViewIndex = fv
}

// TsCondition has been added or modified in a query view. Version used for
// stateless watching
type TsConditionChange_Current struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	TsCondition   *TsCondition `protobuf:"bytes,1,opt,name=ts_condition,json=tsCondition,proto3" json:"ts_condition,omitempty"`
}

func (m *TsConditionChange_Current) Reset() {
	*m = TsConditionChange_Current{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_alerting_proto_v1_ts_condition_change_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *TsConditionChange_Current) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*TsConditionChange_Current) ProtoMessage() {}

func (m *TsConditionChange_Current) ProtoReflect() preflect.Message {
	mi := &edgelq_alerting_proto_v1_ts_condition_change_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*TsConditionChange_Current) GotenMessage() {}

// Deprecated, Use TsConditionChange_Current.ProtoReflect.Descriptor instead.
func (*TsConditionChange_Current) Descriptor() ([]byte, []int) {
	return edgelq_alerting_proto_v1_ts_condition_change_proto_rawDescGZIP(), []int{0, 2}
}

func (m *TsConditionChange_Current) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *TsConditionChange_Current) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *TsConditionChange_Current) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *TsConditionChange_Current) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *TsConditionChange_Current) GetTsCondition() *TsCondition {
	if m != nil {
		return m.TsCondition
	}
	return nil
}

func (m *TsConditionChange_Current) SetTsCondition(fv *TsCondition) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "TsCondition", "TsConditionChange_Current"))
	}
	m.TsCondition = fv
}

// Removed is returned when TsCondition is deleted or leaves Query view
type TsConditionChange_Removed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Name          *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty"`
	// Integer specifying removed TsCondition index. Not populated in stateless
	// watch type.
	ViewIndex int32 `protobuf:"varint,2,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty"`
}

func (m *TsConditionChange_Removed) Reset() {
	*m = TsConditionChange_Removed{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_alerting_proto_v1_ts_condition_change_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *TsConditionChange_Removed) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*TsConditionChange_Removed) ProtoMessage() {}

func (m *TsConditionChange_Removed) ProtoReflect() preflect.Message {
	mi := &edgelq_alerting_proto_v1_ts_condition_change_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*TsConditionChange_Removed) GotenMessage() {}

// Deprecated, Use TsConditionChange_Removed.ProtoReflect.Descriptor instead.
func (*TsConditionChange_Removed) Descriptor() ([]byte, []int) {
	return edgelq_alerting_proto_v1_ts_condition_change_proto_rawDescGZIP(), []int{0, 3}
}

func (m *TsConditionChange_Removed) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *TsConditionChange_Removed) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *TsConditionChange_Removed) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *TsConditionChange_Removed) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *TsConditionChange_Removed) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *TsConditionChange_Removed) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *TsConditionChange_Removed) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "TsConditionChange_Removed"))
	}
	m.Name = fv
}

func (m *TsConditionChange_Removed) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "TsConditionChange_Removed"))
	}
	m.ViewIndex = fv
}

var edgelq_alerting_proto_v1_ts_condition_change_proto preflect.FileDescriptor

var edgelq_alerting_proto_v1_ts_condition_change_proto_rawDesc = []byte{
	0x0a, 0x32, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e,
	0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x73, 0x5f, 0x63, 0x6f,
	0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x6e, 0x74, 0x74, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69,
	0x6e, 0x67, 0x2e, 0x76, 0x31, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x61, 0x6c, 0x65, 0x72,
	0x74, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x73,
	0x5f, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xf0, 0x06, 0x0a, 0x11, 0x54, 0x73, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x40, 0x0a, 0x05, 0x61, 0x64, 0x64, 0x65, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x61, 0x6c, 0x65, 0x72,
	0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x73, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x65, 0x64, 0x48,
	0x00, 0x52, 0x05, 0x61, 0x64, 0x64, 0x65, 0x64, 0x12, 0x49, 0x0a, 0x08, 0x6d, 0x6f, 0x64, 0x69,
	0x66, 0x69, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x6e, 0x74, 0x74,
	0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x73, 0x43,
	0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x4d,
	0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x48, 0x00, 0x52, 0x08, 0x6d, 0x6f, 0x64, 0x69, 0x66,
	0x69, 0x65, 0x64, 0x12, 0x46, 0x0a, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74,
	0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x73, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x48, 0x00, 0x52, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x46, 0x0a, 0x07, 0x72,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x6e,
	0x74, 0x74, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x54,
	0x73, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x48, 0x00, 0x52, 0x07, 0x72, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x64, 0x1a, 0x67, 0x0a, 0x05, 0x41, 0x64, 0x64, 0x65, 0x64, 0x12, 0x3f, 0x0a, 0x0c,
	0x74, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x73, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0b, 0x74, 0x73, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a,
	0x0a, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x76, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x1a, 0x93, 0x02, 0x0a,
	0x08, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x27, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x13, 0xb2, 0xda, 0x21, 0x0f, 0x0a, 0x0d, 0x0a,
	0x0b, 0x54, 0x73, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x3f, 0x0a, 0x0c, 0x74, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x61,
	0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x73, 0x43, 0x6f, 0x6e,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x74, 0x73, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x4e, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73,
	0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d,
	0x61, 0x73, 0x6b, 0x42, 0x13, 0xb2, 0xda, 0x21, 0x0f, 0x32, 0x0d, 0x0a, 0x0b, 0x54, 0x73, 0x43,
	0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4d,
	0x61, 0x73, 0x6b, 0x12, 0x2e, 0x0a, 0x13, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x5f,
	0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x11, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x56, 0x69, 0x65, 0x77, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x76, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x1a, 0x4a, 0x0a, 0x07, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x3f, 0x0a,
	0x0c, 0x74, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69,
	0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x73, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x0b, 0x74, 0x73, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x51,
	0x0a, 0x07, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x12, 0x27, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x13, 0xb2, 0xda, 0x21, 0x0f, 0x0a, 0x0d, 0x0a,
	0x0b, 0x54, 0x73, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x76, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x3a, 0x11, 0x9a, 0xd9, 0x21, 0x0d, 0x0a, 0x0b, 0x54, 0x73, 0x43, 0x6f, 0x6e, 0x64, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0d, 0x0a, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x42, 0x82, 0x01, 0xe8, 0xde, 0x21, 0x00, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e,
	0x6e, 0x74, 0x74, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x62, 0x2e,
	0x76, 0x31, 0x42, 0x16, 0x54, 0x73, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x00, 0x5a, 0x4a, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61,
	0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e,
	0x67, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x74,
	0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x3b, 0x74, 0x73, 0x5f, 0x63,
	0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	edgelq_alerting_proto_v1_ts_condition_change_proto_rawDescOnce sync.Once
	edgelq_alerting_proto_v1_ts_condition_change_proto_rawDescData = edgelq_alerting_proto_v1_ts_condition_change_proto_rawDesc
)

func edgelq_alerting_proto_v1_ts_condition_change_proto_rawDescGZIP() []byte {
	edgelq_alerting_proto_v1_ts_condition_change_proto_rawDescOnce.Do(func() {
		edgelq_alerting_proto_v1_ts_condition_change_proto_rawDescData = protoimpl.X.CompressGZIP(edgelq_alerting_proto_v1_ts_condition_change_proto_rawDescData)
	})
	return edgelq_alerting_proto_v1_ts_condition_change_proto_rawDescData
}

var edgelq_alerting_proto_v1_ts_condition_change_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var edgelq_alerting_proto_v1_ts_condition_change_proto_goTypes = []interface{}{
	(*TsConditionChange)(nil),          // 0: ntt.alerting.v1.TsConditionChange
	(*TsConditionChange_Added)(nil),    // 1: ntt.alerting.v1.TsConditionChange.Added
	(*TsConditionChange_Modified)(nil), // 2: ntt.alerting.v1.TsConditionChange.Modified
	(*TsConditionChange_Current)(nil),  // 3: ntt.alerting.v1.TsConditionChange.Current
	(*TsConditionChange_Removed)(nil),  // 4: ntt.alerting.v1.TsConditionChange.Removed
	(*TsCondition)(nil),                // 5: ntt.alerting.v1.TsCondition
	(*TsCondition_FieldMask)(nil),      // 6: ntt.alerting.v1.TsCondition_FieldMask
}
var edgelq_alerting_proto_v1_ts_condition_change_proto_depIdxs = []int32{
	1, // 0: ntt.alerting.v1.TsConditionChange.added:type_name -> ntt.alerting.v1.TsConditionChange.Added
	2, // 1: ntt.alerting.v1.TsConditionChange.modified:type_name -> ntt.alerting.v1.TsConditionChange.Modified
	3, // 2: ntt.alerting.v1.TsConditionChange.current:type_name -> ntt.alerting.v1.TsConditionChange.Current
	4, // 3: ntt.alerting.v1.TsConditionChange.removed:type_name -> ntt.alerting.v1.TsConditionChange.Removed
	5, // 4: ntt.alerting.v1.TsConditionChange.Added.ts_condition:type_name -> ntt.alerting.v1.TsCondition
	5, // 5: ntt.alerting.v1.TsConditionChange.Modified.ts_condition:type_name -> ntt.alerting.v1.TsCondition
	6, // 6: ntt.alerting.v1.TsConditionChange.Modified.field_mask:type_name -> ntt.alerting.v1.TsCondition_FieldMask
	5, // 7: ntt.alerting.v1.TsConditionChange.Current.ts_condition:type_name -> ntt.alerting.v1.TsCondition
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { edgelq_alerting_proto_v1_ts_condition_change_proto_init() }
func edgelq_alerting_proto_v1_ts_condition_change_proto_init() {
	if edgelq_alerting_proto_v1_ts_condition_change_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		edgelq_alerting_proto_v1_ts_condition_change_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TsConditionChange); i {
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
		edgelq_alerting_proto_v1_ts_condition_change_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TsConditionChange_Added); i {
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
		edgelq_alerting_proto_v1_ts_condition_change_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TsConditionChange_Modified); i {
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
		edgelq_alerting_proto_v1_ts_condition_change_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TsConditionChange_Current); i {
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
		edgelq_alerting_proto_v1_ts_condition_change_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TsConditionChange_Removed); i {
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

	edgelq_alerting_proto_v1_ts_condition_change_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*TsConditionChange_Added_)(nil),
		(*TsConditionChange_Modified_)(nil),
		(*TsConditionChange_Current_)(nil),
		(*TsConditionChange_Removed_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: edgelq_alerting_proto_v1_ts_condition_change_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           edgelq_alerting_proto_v1_ts_condition_change_proto_goTypes,
		DependencyIndexes: edgelq_alerting_proto_v1_ts_condition_change_proto_depIdxs,
		MessageInfos:      edgelq_alerting_proto_v1_ts_condition_change_proto_msgTypes,
	}.Build()
	edgelq_alerting_proto_v1_ts_condition_change_proto = out.File
	edgelq_alerting_proto_v1_ts_condition_change_proto_rawDesc = nil
	edgelq_alerting_proto_v1_ts_condition_change_proto_goTypes = nil
	edgelq_alerting_proto_v1_ts_condition_change_proto_depIdxs = nil
}
