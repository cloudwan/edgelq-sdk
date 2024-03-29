// Code generated by protoc-gen-goten-go
// File: edgelq/monitoring/proto/v3/alerting_condition_change.proto
// DO NOT EDIT!!!

package alerting_condition

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
	alerting_policy "github.com/cloudwan/edgelq-sdk/monitoring/resources/v3/alerting_policy"
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
	_ = &alerting_policy.AlertingPolicy{}
	_ = &fieldmaskpb.FieldMask{}
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// AlertingConditionChange is used by Watch notifications Responses to describe
// change of single AlertingCondition One of Added, Modified, Removed
type AlertingConditionChange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// AlertingCondition change
	//
	// Types that are valid to be assigned to ChangeType:
	//	*AlertingConditionChange_Added_
	//	*AlertingConditionChange_Modified_
	//	*AlertingConditionChange_Current_
	//	*AlertingConditionChange_Removed_
	ChangeType isAlertingConditionChange_ChangeType `protobuf_oneof:"change_type"`
}

func (m *AlertingConditionChange) Reset() {
	*m = AlertingConditionChange{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_monitoring_proto_v3_alerting_condition_change_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *AlertingConditionChange) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*AlertingConditionChange) ProtoMessage() {}

func (m *AlertingConditionChange) ProtoReflect() preflect.Message {
	mi := &edgelq_monitoring_proto_v3_alerting_condition_change_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*AlertingConditionChange) GotenMessage() {}

// Deprecated, Use AlertingConditionChange.ProtoReflect.Descriptor instead.
func (*AlertingConditionChange) Descriptor() ([]byte, []int) {
	return edgelq_monitoring_proto_v3_alerting_condition_change_proto_rawDescGZIP(), []int{0}
}

func (m *AlertingConditionChange) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *AlertingConditionChange) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *AlertingConditionChange) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *AlertingConditionChange) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

type isAlertingConditionChange_ChangeType interface {
	isAlertingConditionChange_ChangeType()
}

type AlertingConditionChange_Added_ struct {
	// Added is returned when watched document is added, either created or
	// enters Query view
	Added *AlertingConditionChange_Added `protobuf:"bytes,1,opt,name=added,proto3,oneof" firestore:"added"`
}
type AlertingConditionChange_Modified_ struct {
	// Modified is returned when watched document is modified
	Modified *AlertingConditionChange_Modified `protobuf:"bytes,2,opt,name=modified,proto3,oneof" firestore:"modified"`
}
type AlertingConditionChange_Current_ struct {
	// Current is returned in stateless watch when document enters query view or
	// is modified within.
	Current *AlertingConditionChange_Current `protobuf:"bytes,4,opt,name=current,proto3,oneof" firestore:"current"`
}
type AlertingConditionChange_Removed_ struct {
	// Removed is returned when AlertingCondition is deleted or leaves Query
	// view
	Removed *AlertingConditionChange_Removed `protobuf:"bytes,3,opt,name=removed,proto3,oneof" firestore:"removed"`
}

func (*AlertingConditionChange_Added_) isAlertingConditionChange_ChangeType()    {}
func (*AlertingConditionChange_Modified_) isAlertingConditionChange_ChangeType() {}
func (*AlertingConditionChange_Current_) isAlertingConditionChange_ChangeType()  {}
func (*AlertingConditionChange_Removed_) isAlertingConditionChange_ChangeType()  {}
func (m *AlertingConditionChange) GetChangeType() isAlertingConditionChange_ChangeType {
	if m != nil {
		return m.ChangeType
	}
	return nil
}
func (m *AlertingConditionChange) GetAdded() *AlertingConditionChange_Added {
	if x, ok := m.GetChangeType().(*AlertingConditionChange_Added_); ok {
		return x.Added
	}
	return nil
}
func (m *AlertingConditionChange) GetModified() *AlertingConditionChange_Modified {
	if x, ok := m.GetChangeType().(*AlertingConditionChange_Modified_); ok {
		return x.Modified
	}
	return nil
}
func (m *AlertingConditionChange) GetCurrent() *AlertingConditionChange_Current {
	if x, ok := m.GetChangeType().(*AlertingConditionChange_Current_); ok {
		return x.Current
	}
	return nil
}
func (m *AlertingConditionChange) GetRemoved() *AlertingConditionChange_Removed {
	if x, ok := m.GetChangeType().(*AlertingConditionChange_Removed_); ok {
		return x.Removed
	}
	return nil
}
func (m *AlertingConditionChange) SetChangeType(ofv isAlertingConditionChange_ChangeType) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "isAlertingConditionChange_ChangeType", "AlertingConditionChange"))
	}
	m.ChangeType = ofv
}
func (m *AlertingConditionChange) SetAdded(fv *AlertingConditionChange_Added) {
	m.SetChangeType(&AlertingConditionChange_Added_{Added: fv})
}
func (m *AlertingConditionChange) SetModified(fv *AlertingConditionChange_Modified) {
	m.SetChangeType(&AlertingConditionChange_Modified_{Modified: fv})
}
func (m *AlertingConditionChange) SetCurrent(fv *AlertingConditionChange_Current) {
	m.SetChangeType(&AlertingConditionChange_Current_{Current: fv})
}
func (m *AlertingConditionChange) SetRemoved(fv *AlertingConditionChange_Removed) {
	m.SetChangeType(&AlertingConditionChange_Removed_{Removed: fv})
}

// AlertingCondition has been added to query view
type AlertingConditionChange_Added struct {
	state             protoimpl.MessageState
	sizeCache         protoimpl.SizeCache
	unknownFields     protoimpl.UnknownFields
	AlertingCondition *AlertingCondition `protobuf:"bytes,1,opt,name=alerting_condition,json=alertingCondition,proto3" json:"alerting_condition,omitempty" firestore:"alertingCondition"`
	// Integer describing index of added AlertingCondition in resulting query
	// view.
	ViewIndex int32 `protobuf:"varint,2,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty" firestore:"viewIndex"`
}

func (m *AlertingConditionChange_Added) Reset() {
	*m = AlertingConditionChange_Added{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_monitoring_proto_v3_alerting_condition_change_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *AlertingConditionChange_Added) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*AlertingConditionChange_Added) ProtoMessage() {}

func (m *AlertingConditionChange_Added) ProtoReflect() preflect.Message {
	mi := &edgelq_monitoring_proto_v3_alerting_condition_change_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*AlertingConditionChange_Added) GotenMessage() {}

// Deprecated, Use AlertingConditionChange_Added.ProtoReflect.Descriptor instead.
func (*AlertingConditionChange_Added) Descriptor() ([]byte, []int) {
	return edgelq_monitoring_proto_v3_alerting_condition_change_proto_rawDescGZIP(), []int{0, 0}
}

func (m *AlertingConditionChange_Added) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *AlertingConditionChange_Added) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *AlertingConditionChange_Added) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *AlertingConditionChange_Added) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *AlertingConditionChange_Added) GetAlertingCondition() *AlertingCondition {
	if m != nil {
		return m.AlertingCondition
	}
	return nil
}

func (m *AlertingConditionChange_Added) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *AlertingConditionChange_Added) SetAlertingCondition(fv *AlertingCondition) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "AlertingCondition", "AlertingConditionChange_Added"))
	}
	m.AlertingCondition = fv
}

func (m *AlertingConditionChange_Added) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "AlertingConditionChange_Added"))
	}
	m.ViewIndex = fv
}

// AlertingCondition changed some of it's fields - contains either full
// document or masked change
type AlertingConditionChange_Modified struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Name of modified AlertingCondition
	Name *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// New version of AlertingCondition or masked difference, depending on
	// mask_changes instrumentation of issued [WatchAlertingConditionRequest] or
	// [WatchAlertingConditionsRequest]
	AlertingCondition *AlertingCondition `protobuf:"bytes,2,opt,name=alerting_condition,json=alertingCondition,proto3" json:"alerting_condition,omitempty" firestore:"alertingCondition"`
	// Used when mask_changes is set, contains field paths of modified
	// properties.
	FieldMask *AlertingCondition_FieldMask `protobuf:"bytes,3,opt,customtype=AlertingCondition_FieldMask,name=field_mask,json=fieldMask,proto3" json:"field_mask,omitempty" firestore:"fieldMask"`
	// Previous view index specifies previous position of modified
	// AlertingCondition. When modification doesn't affect sorted order, value
	// will remain identical to [view_index].
	PreviousViewIndex int32 `protobuf:"varint,4,opt,name=previous_view_index,json=previousViewIndex,proto3" json:"previous_view_index,omitempty" firestore:"previousViewIndex"`
	// Integer specifying AlertingCondition new index in resulting query view.
	ViewIndex int32 `protobuf:"varint,5,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty" firestore:"viewIndex"`
}

func (m *AlertingConditionChange_Modified) Reset() {
	*m = AlertingConditionChange_Modified{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_monitoring_proto_v3_alerting_condition_change_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *AlertingConditionChange_Modified) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*AlertingConditionChange_Modified) ProtoMessage() {}

func (m *AlertingConditionChange_Modified) ProtoReflect() preflect.Message {
	mi := &edgelq_monitoring_proto_v3_alerting_condition_change_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*AlertingConditionChange_Modified) GotenMessage() {}

// Deprecated, Use AlertingConditionChange_Modified.ProtoReflect.Descriptor instead.
func (*AlertingConditionChange_Modified) Descriptor() ([]byte, []int) {
	return edgelq_monitoring_proto_v3_alerting_condition_change_proto_rawDescGZIP(), []int{0, 1}
}

func (m *AlertingConditionChange_Modified) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *AlertingConditionChange_Modified) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *AlertingConditionChange_Modified) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *AlertingConditionChange_Modified) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *AlertingConditionChange_Modified) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *AlertingConditionChange_Modified) GetAlertingCondition() *AlertingCondition {
	if m != nil {
		return m.AlertingCondition
	}
	return nil
}

func (m *AlertingConditionChange_Modified) GetFieldMask() *AlertingCondition_FieldMask {
	if m != nil {
		return m.FieldMask
	}
	return nil
}

func (m *AlertingConditionChange_Modified) GetPreviousViewIndex() int32 {
	if m != nil {
		return m.PreviousViewIndex
	}
	return int32(0)
}

func (m *AlertingConditionChange_Modified) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *AlertingConditionChange_Modified) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "AlertingConditionChange_Modified"))
	}
	m.Name = fv
}

func (m *AlertingConditionChange_Modified) SetAlertingCondition(fv *AlertingCondition) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "AlertingCondition", "AlertingConditionChange_Modified"))
	}
	m.AlertingCondition = fv
}

func (m *AlertingConditionChange_Modified) SetFieldMask(fv *AlertingCondition_FieldMask) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "FieldMask", "AlertingConditionChange_Modified"))
	}
	m.FieldMask = fv
}

func (m *AlertingConditionChange_Modified) SetPreviousViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "PreviousViewIndex", "AlertingConditionChange_Modified"))
	}
	m.PreviousViewIndex = fv
}

func (m *AlertingConditionChange_Modified) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "AlertingConditionChange_Modified"))
	}
	m.ViewIndex = fv
}

// AlertingCondition has been added or modified in a query view. Version used
// for stateless watching
type AlertingConditionChange_Current struct {
	state             protoimpl.MessageState
	sizeCache         protoimpl.SizeCache
	unknownFields     protoimpl.UnknownFields
	AlertingCondition *AlertingCondition `protobuf:"bytes,1,opt,name=alerting_condition,json=alertingCondition,proto3" json:"alerting_condition,omitempty" firestore:"alertingCondition"`
}

func (m *AlertingConditionChange_Current) Reset() {
	*m = AlertingConditionChange_Current{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_monitoring_proto_v3_alerting_condition_change_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *AlertingConditionChange_Current) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*AlertingConditionChange_Current) ProtoMessage() {}

func (m *AlertingConditionChange_Current) ProtoReflect() preflect.Message {
	mi := &edgelq_monitoring_proto_v3_alerting_condition_change_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*AlertingConditionChange_Current) GotenMessage() {}

// Deprecated, Use AlertingConditionChange_Current.ProtoReflect.Descriptor instead.
func (*AlertingConditionChange_Current) Descriptor() ([]byte, []int) {
	return edgelq_monitoring_proto_v3_alerting_condition_change_proto_rawDescGZIP(), []int{0, 2}
}

func (m *AlertingConditionChange_Current) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *AlertingConditionChange_Current) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *AlertingConditionChange_Current) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *AlertingConditionChange_Current) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *AlertingConditionChange_Current) GetAlertingCondition() *AlertingCondition {
	if m != nil {
		return m.AlertingCondition
	}
	return nil
}

func (m *AlertingConditionChange_Current) SetAlertingCondition(fv *AlertingCondition) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "AlertingCondition", "AlertingConditionChange_Current"))
	}
	m.AlertingCondition = fv
}

// Removed is returned when AlertingCondition is deleted or leaves Query view
type AlertingConditionChange_Removed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Name          *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// Integer specifying removed AlertingCondition index. Not populated in
	// stateless watch type.
	ViewIndex int32 `protobuf:"varint,2,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty" firestore:"viewIndex"`
}

func (m *AlertingConditionChange_Removed) Reset() {
	*m = AlertingConditionChange_Removed{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_monitoring_proto_v3_alerting_condition_change_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *AlertingConditionChange_Removed) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*AlertingConditionChange_Removed) ProtoMessage() {}

func (m *AlertingConditionChange_Removed) ProtoReflect() preflect.Message {
	mi := &edgelq_monitoring_proto_v3_alerting_condition_change_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*AlertingConditionChange_Removed) GotenMessage() {}

// Deprecated, Use AlertingConditionChange_Removed.ProtoReflect.Descriptor instead.
func (*AlertingConditionChange_Removed) Descriptor() ([]byte, []int) {
	return edgelq_monitoring_proto_v3_alerting_condition_change_proto_rawDescGZIP(), []int{0, 3}
}

func (m *AlertingConditionChange_Removed) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *AlertingConditionChange_Removed) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *AlertingConditionChange_Removed) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *AlertingConditionChange_Removed) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *AlertingConditionChange_Removed) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *AlertingConditionChange_Removed) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *AlertingConditionChange_Removed) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "AlertingConditionChange_Removed"))
	}
	m.Name = fv
}

func (m *AlertingConditionChange_Removed) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "AlertingConditionChange_Removed"))
	}
	m.ViewIndex = fv
}

var edgelq_monitoring_proto_v3_alerting_condition_change_proto preflect.FileDescriptor

var edgelq_monitoring_proto_v3_alerting_condition_change_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x69, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x6e, 0x74,
	0x74, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x1a,
	0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f,
	0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x74,
	0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x33, 0x65,
	0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69,
	0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xea, 0x07, 0x0a, 0x17, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x43,
	0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x48,
	0x0a, 0x05, 0x61, 0x64, 0x64, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e,
	0x6e, 0x74, 0x74, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76,
	0x33, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x65, 0x64, 0x48,
	0x00, 0x52, 0x05, 0x61, 0x64, 0x64, 0x65, 0x64, 0x12, 0x51, 0x0a, 0x08, 0x6d, 0x6f, 0x64, 0x69,
	0x66, 0x69, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x6e, 0x74, 0x74,
	0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x41,
	0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x48,
	0x00, 0x52, 0x08, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x4e, 0x0a, 0x07, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x6e,
	0x74, 0x74, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33,
	0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x48, 0x00, 0x52, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x4e, 0x0a, 0x07, 0x72,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x6e,
	0x74, 0x74, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33,
	0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64,
	0x48, 0x00, 0x52, 0x07, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x1a, 0x7b, 0x0a, 0x05, 0x41,
	0x64, 0x64, 0x65, 0x64, 0x12, 0x53, 0x0a, 0x12, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67,
	0x5f, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x24, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x33, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x11, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67,
	0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x69, 0x65,
	0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x76,
	0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x1a, 0xb3, 0x02, 0x0a, 0x08, 0x4d, 0x6f, 0x64,
	0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x2d, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x19, 0xb2, 0xda, 0x21, 0x15, 0x0a, 0x13, 0x0a, 0x11, 0x41, 0x6c, 0x65,
	0x72, 0x74, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x53, 0x0a, 0x12, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67,
	0x5f, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x24, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x33, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x11, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67,
	0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x54, 0x0a, 0x0a, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x42, 0x19, 0xb2, 0xda, 0x21, 0x15, 0x32,
	0x13, 0x0a, 0x11, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x64, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x12,
	0x2e, 0x0a, 0x13, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x5f, 0x76, 0x69, 0x65, 0x77,
	0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x70, 0x72,
	0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x56, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12,
	0x1d, 0x0a, 0x0a, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x76, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x1a, 0x5e,
	0x0a, 0x07, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x53, 0x0a, 0x12, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x6d, 0x6f, 0x6e, 0x69,
	0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x69,
	0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x11, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x57,
	0x0a, 0x07, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x12, 0x2d, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x19, 0xb2, 0xda, 0x21, 0x15, 0x0a, 0x13, 0x0a,
	0x11, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x69, 0x65, 0x77,
	0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x76, 0x69,
	0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x3a, 0x17, 0x9a, 0xd9, 0x21, 0x13, 0x0a, 0x11, 0x41,
	0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x0d, 0x0a, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x42,
	0x98, 0x01, 0xe8, 0xde, 0x21, 0x00, 0x0a, 0x18, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x74, 0x74, 0x2e,
	0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x33,
	0x42, 0x1c, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x00,
	0x5a, 0x58, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6d, 0x6f, 0x6e,
	0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f,
	0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x3b, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67,
	0x5f, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	edgelq_monitoring_proto_v3_alerting_condition_change_proto_rawDescOnce sync.Once
	edgelq_monitoring_proto_v3_alerting_condition_change_proto_rawDescData = edgelq_monitoring_proto_v3_alerting_condition_change_proto_rawDesc
)

func edgelq_monitoring_proto_v3_alerting_condition_change_proto_rawDescGZIP() []byte {
	edgelq_monitoring_proto_v3_alerting_condition_change_proto_rawDescOnce.Do(func() {
		edgelq_monitoring_proto_v3_alerting_condition_change_proto_rawDescData = protoimpl.X.CompressGZIP(edgelq_monitoring_proto_v3_alerting_condition_change_proto_rawDescData)
	})
	return edgelq_monitoring_proto_v3_alerting_condition_change_proto_rawDescData
}

var edgelq_monitoring_proto_v3_alerting_condition_change_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var edgelq_monitoring_proto_v3_alerting_condition_change_proto_goTypes = []interface{}{
	(*AlertingConditionChange)(nil),          // 0: ntt.monitoring.v3.AlertingConditionChange
	(*AlertingConditionChange_Added)(nil),    // 1: ntt.monitoring.v3.AlertingConditionChange.Added
	(*AlertingConditionChange_Modified)(nil), // 2: ntt.monitoring.v3.AlertingConditionChange.Modified
	(*AlertingConditionChange_Current)(nil),  // 3: ntt.monitoring.v3.AlertingConditionChange.Current
	(*AlertingConditionChange_Removed)(nil),  // 4: ntt.monitoring.v3.AlertingConditionChange.Removed
	(*AlertingCondition)(nil),                // 5: ntt.monitoring.v3.AlertingCondition
	(*AlertingCondition_FieldMask)(nil),      // 6: ntt.monitoring.v3.AlertingCondition_FieldMask
}
var edgelq_monitoring_proto_v3_alerting_condition_change_proto_depIdxs = []int32{
	1, // 0: ntt.monitoring.v3.AlertingConditionChange.added:type_name -> ntt.monitoring.v3.AlertingConditionChange.Added
	2, // 1: ntt.monitoring.v3.AlertingConditionChange.modified:type_name -> ntt.monitoring.v3.AlertingConditionChange.Modified
	3, // 2: ntt.monitoring.v3.AlertingConditionChange.current:type_name -> ntt.monitoring.v3.AlertingConditionChange.Current
	4, // 3: ntt.monitoring.v3.AlertingConditionChange.removed:type_name -> ntt.monitoring.v3.AlertingConditionChange.Removed
	5, // 4: ntt.monitoring.v3.AlertingConditionChange.Added.alerting_condition:type_name -> ntt.monitoring.v3.AlertingCondition
	5, // 5: ntt.monitoring.v3.AlertingConditionChange.Modified.alerting_condition:type_name -> ntt.monitoring.v3.AlertingCondition
	6, // 6: ntt.monitoring.v3.AlertingConditionChange.Modified.field_mask:type_name -> ntt.monitoring.v3.AlertingCondition_FieldMask
	5, // 7: ntt.monitoring.v3.AlertingConditionChange.Current.alerting_condition:type_name -> ntt.monitoring.v3.AlertingCondition
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { edgelq_monitoring_proto_v3_alerting_condition_change_proto_init() }
func edgelq_monitoring_proto_v3_alerting_condition_change_proto_init() {
	if edgelq_monitoring_proto_v3_alerting_condition_change_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		edgelq_monitoring_proto_v3_alerting_condition_change_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlertingConditionChange); i {
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
		edgelq_monitoring_proto_v3_alerting_condition_change_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlertingConditionChange_Added); i {
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
		edgelq_monitoring_proto_v3_alerting_condition_change_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlertingConditionChange_Modified); i {
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
		edgelq_monitoring_proto_v3_alerting_condition_change_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlertingConditionChange_Current); i {
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
		edgelq_monitoring_proto_v3_alerting_condition_change_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlertingConditionChange_Removed); i {
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

	edgelq_monitoring_proto_v3_alerting_condition_change_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*AlertingConditionChange_Added_)(nil),
		(*AlertingConditionChange_Modified_)(nil),
		(*AlertingConditionChange_Current_)(nil),
		(*AlertingConditionChange_Removed_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: edgelq_monitoring_proto_v3_alerting_condition_change_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           edgelq_monitoring_proto_v3_alerting_condition_change_proto_goTypes,
		DependencyIndexes: edgelq_monitoring_proto_v3_alerting_condition_change_proto_depIdxs,
		MessageInfos:      edgelq_monitoring_proto_v3_alerting_condition_change_proto_msgTypes,
	}.Build()
	edgelq_monitoring_proto_v3_alerting_condition_change_proto = out.File
	edgelq_monitoring_proto_v3_alerting_condition_change_proto_rawDesc = nil
	edgelq_monitoring_proto_v3_alerting_condition_change_proto_goTypes = nil
	edgelq_monitoring_proto_v3_alerting_condition_change_proto_depIdxs = nil
}
