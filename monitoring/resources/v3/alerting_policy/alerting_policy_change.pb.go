// Code generated by protoc-gen-goten-go
// File: edgelq/monitoring/proto/v3/alerting_policy_change.proto
// DO NOT EDIT!!!

package alerting_policy

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
	project "github.com/cloudwan/edgelq-sdk/monitoring/resources/v3/project"
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
	_ = &project.Project{}
	_ = &field_mask.FieldMask{}
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// AlertingPolicyChange is used by Watch notifications Responses to describe
// change of single AlertingPolicy One of Added, Modified, Removed
type AlertingPolicyChange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// AlertingPolicy change
	//
	// Types that are valid to be assigned to ChangeType:
	//	*AlertingPolicyChange_Added_
	//	*AlertingPolicyChange_Modified_
	//	*AlertingPolicyChange_Current_
	//	*AlertingPolicyChange_Removed_
	ChangeType isAlertingPolicyChange_ChangeType `protobuf_oneof:"change_type"`
}

func (m *AlertingPolicyChange) Reset() {
	*m = AlertingPolicyChange{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_monitoring_proto_v3_alerting_policy_change_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *AlertingPolicyChange) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*AlertingPolicyChange) ProtoMessage() {}

func (m *AlertingPolicyChange) ProtoReflect() preflect.Message {
	mi := &edgelq_monitoring_proto_v3_alerting_policy_change_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*AlertingPolicyChange) GotenMessage() {}

// Deprecated, Use AlertingPolicyChange.ProtoReflect.Descriptor instead.
func (*AlertingPolicyChange) Descriptor() ([]byte, []int) {
	return edgelq_monitoring_proto_v3_alerting_policy_change_proto_rawDescGZIP(), []int{0}
}

func (m *AlertingPolicyChange) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *AlertingPolicyChange) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *AlertingPolicyChange) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *AlertingPolicyChange) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

type isAlertingPolicyChange_ChangeType interface {
	isAlertingPolicyChange_ChangeType()
}

type AlertingPolicyChange_Added_ struct {
	// Added is returned when watched document is added, either created or
	// enters Query view
	Added *AlertingPolicyChange_Added `protobuf:"bytes,1,opt,name=added,proto3,oneof" firestore:"added"`
}
type AlertingPolicyChange_Modified_ struct {
	// Modified is returned when watched document is modified
	Modified *AlertingPolicyChange_Modified `protobuf:"bytes,2,opt,name=modified,proto3,oneof" firestore:"modified"`
}
type AlertingPolicyChange_Current_ struct {
	// Current is returned in stateless watch when document enters query view or
	// is modified within.
	Current *AlertingPolicyChange_Current `protobuf:"bytes,4,opt,name=current,proto3,oneof" firestore:"current"`
}
type AlertingPolicyChange_Removed_ struct {
	// Removed is returned when AlertingPolicy is deleted or leaves Query view
	Removed *AlertingPolicyChange_Removed `protobuf:"bytes,3,opt,name=removed,proto3,oneof" firestore:"removed"`
}

func (*AlertingPolicyChange_Added_) isAlertingPolicyChange_ChangeType()    {}
func (*AlertingPolicyChange_Modified_) isAlertingPolicyChange_ChangeType() {}
func (*AlertingPolicyChange_Current_) isAlertingPolicyChange_ChangeType()  {}
func (*AlertingPolicyChange_Removed_) isAlertingPolicyChange_ChangeType()  {}
func (m *AlertingPolicyChange) GetChangeType() isAlertingPolicyChange_ChangeType {
	if m != nil {
		return m.ChangeType
	}
	return nil
}
func (m *AlertingPolicyChange) GetAdded() *AlertingPolicyChange_Added {
	if x, ok := m.GetChangeType().(*AlertingPolicyChange_Added_); ok {
		return x.Added
	}
	return nil
}
func (m *AlertingPolicyChange) GetModified() *AlertingPolicyChange_Modified {
	if x, ok := m.GetChangeType().(*AlertingPolicyChange_Modified_); ok {
		return x.Modified
	}
	return nil
}
func (m *AlertingPolicyChange) GetCurrent() *AlertingPolicyChange_Current {
	if x, ok := m.GetChangeType().(*AlertingPolicyChange_Current_); ok {
		return x.Current
	}
	return nil
}
func (m *AlertingPolicyChange) GetRemoved() *AlertingPolicyChange_Removed {
	if x, ok := m.GetChangeType().(*AlertingPolicyChange_Removed_); ok {
		return x.Removed
	}
	return nil
}
func (m *AlertingPolicyChange) SetChangeType(ofv isAlertingPolicyChange_ChangeType) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "isAlertingPolicyChange_ChangeType", "AlertingPolicyChange"))
	}
	m.ChangeType = ofv
}
func (m *AlertingPolicyChange) SetAdded(fv *AlertingPolicyChange_Added) {
	m.SetChangeType(&AlertingPolicyChange_Added_{Added: fv})
}
func (m *AlertingPolicyChange) SetModified(fv *AlertingPolicyChange_Modified) {
	m.SetChangeType(&AlertingPolicyChange_Modified_{Modified: fv})
}
func (m *AlertingPolicyChange) SetCurrent(fv *AlertingPolicyChange_Current) {
	m.SetChangeType(&AlertingPolicyChange_Current_{Current: fv})
}
func (m *AlertingPolicyChange) SetRemoved(fv *AlertingPolicyChange_Removed) {
	m.SetChangeType(&AlertingPolicyChange_Removed_{Removed: fv})
}

// AlertingPolicy has been added to query view
type AlertingPolicyChange_Added struct {
	state          protoimpl.MessageState
	sizeCache      protoimpl.SizeCache
	unknownFields  protoimpl.UnknownFields
	AlertingPolicy *AlertingPolicy `protobuf:"bytes,1,opt,name=alerting_policy,json=alertingPolicy,proto3" json:"alerting_policy,omitempty" firestore:"alertingPolicy"`
	// Integer describing index of added AlertingPolicy in resulting query view.
	ViewIndex int32 `protobuf:"varint,2,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty" firestore:"viewIndex"`
}

func (m *AlertingPolicyChange_Added) Reset() {
	*m = AlertingPolicyChange_Added{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_monitoring_proto_v3_alerting_policy_change_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *AlertingPolicyChange_Added) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*AlertingPolicyChange_Added) ProtoMessage() {}

func (m *AlertingPolicyChange_Added) ProtoReflect() preflect.Message {
	mi := &edgelq_monitoring_proto_v3_alerting_policy_change_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*AlertingPolicyChange_Added) GotenMessage() {}

// Deprecated, Use AlertingPolicyChange_Added.ProtoReflect.Descriptor instead.
func (*AlertingPolicyChange_Added) Descriptor() ([]byte, []int) {
	return edgelq_monitoring_proto_v3_alerting_policy_change_proto_rawDescGZIP(), []int{0, 0}
}

func (m *AlertingPolicyChange_Added) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *AlertingPolicyChange_Added) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *AlertingPolicyChange_Added) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *AlertingPolicyChange_Added) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *AlertingPolicyChange_Added) GetAlertingPolicy() *AlertingPolicy {
	if m != nil {
		return m.AlertingPolicy
	}
	return nil
}

func (m *AlertingPolicyChange_Added) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *AlertingPolicyChange_Added) SetAlertingPolicy(fv *AlertingPolicy) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "AlertingPolicy", "AlertingPolicyChange_Added"))
	}
	m.AlertingPolicy = fv
}

func (m *AlertingPolicyChange_Added) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "AlertingPolicyChange_Added"))
	}
	m.ViewIndex = fv
}

// AlertingPolicy changed some of it's fields - contains either full document
// or masked change
type AlertingPolicyChange_Modified struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Name of modified AlertingPolicy
	Name *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// New version of AlertingPolicy or masked difference, depending on
	// mask_changes instrumentation of issued [WatchAlertingPolicyRequest] or
	// [WatchAlertingPoliciesRequest]
	AlertingPolicy *AlertingPolicy `protobuf:"bytes,2,opt,name=alerting_policy,json=alertingPolicy,proto3" json:"alerting_policy,omitempty" firestore:"alertingPolicy"`
	// Used when mask_changes is set, contains field paths of modified
	// properties.
	FieldMask *AlertingPolicy_FieldMask `protobuf:"bytes,3,opt,customtype=AlertingPolicy_FieldMask,name=field_mask,json=fieldMask,proto3" json:"field_mask,omitempty" firestore:"fieldMask"`
	// Previous view index specifies previous position of modified
	// AlertingPolicy. When modification doesn't affect sorted order, value will
	// remain identical to [view_index].
	PreviousViewIndex int32 `protobuf:"varint,4,opt,name=previous_view_index,json=previousViewIndex,proto3" json:"previous_view_index,omitempty" firestore:"previousViewIndex"`
	// Integer specifying AlertingPolicy new index in resulting query view.
	ViewIndex int32 `protobuf:"varint,5,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty" firestore:"viewIndex"`
}

func (m *AlertingPolicyChange_Modified) Reset() {
	*m = AlertingPolicyChange_Modified{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_monitoring_proto_v3_alerting_policy_change_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *AlertingPolicyChange_Modified) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*AlertingPolicyChange_Modified) ProtoMessage() {}

func (m *AlertingPolicyChange_Modified) ProtoReflect() preflect.Message {
	mi := &edgelq_monitoring_proto_v3_alerting_policy_change_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*AlertingPolicyChange_Modified) GotenMessage() {}

// Deprecated, Use AlertingPolicyChange_Modified.ProtoReflect.Descriptor instead.
func (*AlertingPolicyChange_Modified) Descriptor() ([]byte, []int) {
	return edgelq_monitoring_proto_v3_alerting_policy_change_proto_rawDescGZIP(), []int{0, 1}
}

func (m *AlertingPolicyChange_Modified) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *AlertingPolicyChange_Modified) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *AlertingPolicyChange_Modified) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *AlertingPolicyChange_Modified) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *AlertingPolicyChange_Modified) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *AlertingPolicyChange_Modified) GetAlertingPolicy() *AlertingPolicy {
	if m != nil {
		return m.AlertingPolicy
	}
	return nil
}

func (m *AlertingPolicyChange_Modified) GetFieldMask() *AlertingPolicy_FieldMask {
	if m != nil {
		return m.FieldMask
	}
	return nil
}

func (m *AlertingPolicyChange_Modified) GetPreviousViewIndex() int32 {
	if m != nil {
		return m.PreviousViewIndex
	}
	return int32(0)
}

func (m *AlertingPolicyChange_Modified) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *AlertingPolicyChange_Modified) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "AlertingPolicyChange_Modified"))
	}
	m.Name = fv
}

func (m *AlertingPolicyChange_Modified) SetAlertingPolicy(fv *AlertingPolicy) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "AlertingPolicy", "AlertingPolicyChange_Modified"))
	}
	m.AlertingPolicy = fv
}

func (m *AlertingPolicyChange_Modified) SetFieldMask(fv *AlertingPolicy_FieldMask) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "FieldMask", "AlertingPolicyChange_Modified"))
	}
	m.FieldMask = fv
}

func (m *AlertingPolicyChange_Modified) SetPreviousViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "PreviousViewIndex", "AlertingPolicyChange_Modified"))
	}
	m.PreviousViewIndex = fv
}

func (m *AlertingPolicyChange_Modified) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "AlertingPolicyChange_Modified"))
	}
	m.ViewIndex = fv
}

// AlertingPolicy has been added or modified in a query view. Version used for
// stateless watching
type AlertingPolicyChange_Current struct {
	state          protoimpl.MessageState
	sizeCache      protoimpl.SizeCache
	unknownFields  protoimpl.UnknownFields
	AlertingPolicy *AlertingPolicy `protobuf:"bytes,1,opt,name=alerting_policy,json=alertingPolicy,proto3" json:"alerting_policy,omitempty" firestore:"alertingPolicy"`
}

func (m *AlertingPolicyChange_Current) Reset() {
	*m = AlertingPolicyChange_Current{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_monitoring_proto_v3_alerting_policy_change_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *AlertingPolicyChange_Current) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*AlertingPolicyChange_Current) ProtoMessage() {}

func (m *AlertingPolicyChange_Current) ProtoReflect() preflect.Message {
	mi := &edgelq_monitoring_proto_v3_alerting_policy_change_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*AlertingPolicyChange_Current) GotenMessage() {}

// Deprecated, Use AlertingPolicyChange_Current.ProtoReflect.Descriptor instead.
func (*AlertingPolicyChange_Current) Descriptor() ([]byte, []int) {
	return edgelq_monitoring_proto_v3_alerting_policy_change_proto_rawDescGZIP(), []int{0, 2}
}

func (m *AlertingPolicyChange_Current) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *AlertingPolicyChange_Current) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *AlertingPolicyChange_Current) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *AlertingPolicyChange_Current) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *AlertingPolicyChange_Current) GetAlertingPolicy() *AlertingPolicy {
	if m != nil {
		return m.AlertingPolicy
	}
	return nil
}

func (m *AlertingPolicyChange_Current) SetAlertingPolicy(fv *AlertingPolicy) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "AlertingPolicy", "AlertingPolicyChange_Current"))
	}
	m.AlertingPolicy = fv
}

// Removed is returned when AlertingPolicy is deleted or leaves Query view
type AlertingPolicyChange_Removed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Name          *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// Integer specifying removed AlertingPolicy index. Not populated in
	// stateless watch type.
	ViewIndex int32 `protobuf:"varint,2,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty" firestore:"viewIndex"`
}

func (m *AlertingPolicyChange_Removed) Reset() {
	*m = AlertingPolicyChange_Removed{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_monitoring_proto_v3_alerting_policy_change_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *AlertingPolicyChange_Removed) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*AlertingPolicyChange_Removed) ProtoMessage() {}

func (m *AlertingPolicyChange_Removed) ProtoReflect() preflect.Message {
	mi := &edgelq_monitoring_proto_v3_alerting_policy_change_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*AlertingPolicyChange_Removed) GotenMessage() {}

// Deprecated, Use AlertingPolicyChange_Removed.ProtoReflect.Descriptor instead.
func (*AlertingPolicyChange_Removed) Descriptor() ([]byte, []int) {
	return edgelq_monitoring_proto_v3_alerting_policy_change_proto_rawDescGZIP(), []int{0, 3}
}

func (m *AlertingPolicyChange_Removed) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *AlertingPolicyChange_Removed) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *AlertingPolicyChange_Removed) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *AlertingPolicyChange_Removed) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *AlertingPolicyChange_Removed) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *AlertingPolicyChange_Removed) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *AlertingPolicyChange_Removed) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "AlertingPolicyChange_Removed"))
	}
	m.Name = fv
}

func (m *AlertingPolicyChange_Removed) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "AlertingPolicyChange_Removed"))
	}
	m.ViewIndex = fv
}

var edgelq_monitoring_proto_v3_alerting_policy_change_proto preflect.FileDescriptor

var edgelq_monitoring_proto_v3_alerting_policy_change_proto_rawDesc = []byte{
	0x0a, 0x37, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x69, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x6e, 0x74, 0x74, 0x2e, 0x6d,
	0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x1a, 0x19, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d,
	0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x74,
	0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x30, 0x65, 0x64, 0x67, 0x65,
	0x6c, 0x71, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x5f,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb4, 0x07, 0x0a,
	0x14, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x45, 0x0a, 0x05, 0x61, 0x64, 0x64, 0x65, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74,
	0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e,
	0x67, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x41, 0x64,
	0x64, 0x65, 0x64, 0x48, 0x00, 0x52, 0x05, 0x61, 0x64, 0x64, 0x65, 0x64, 0x12, 0x4e, 0x0a, 0x08,
	0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30,
	0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x33, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64,
	0x48, 0x00, 0x52, 0x08, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x4b, 0x0a, 0x07,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e,
	0x6e, 0x74, 0x74, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76,
	0x33, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x48, 0x00,
	0x52, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x4b, 0x0a, 0x07, 0x72, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x6e, 0x74, 0x74,
	0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x41,
	0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x48, 0x00, 0x52, 0x07, 0x72,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x1a, 0x72, 0x0a, 0x05, 0x41, 0x64, 0x64, 0x65, 0x64, 0x12,
	0x4a, 0x0a, 0x0f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x6d,
	0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x41, 0x6c, 0x65,
	0x72, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x0e, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x76,
	0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x76, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x1a, 0xa4, 0x02, 0x0a, 0x08, 0x4d,
	0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x2a, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x16, 0xb2, 0xda, 0x21, 0x12, 0x0a, 0x10, 0x0a, 0x0e, 0x41,
	0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x4a, 0x0a, 0x0f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x5f,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6e,
	0x74, 0x74, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33,
	0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52,
	0x0e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12,
	0x51, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x42,
	0x16, 0xb2, 0xda, 0x21, 0x12, 0x32, 0x10, 0x0a, 0x0e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e,
	0x67, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61,
	0x73, 0x6b, 0x12, 0x2e, 0x0a, 0x13, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x5f, 0x76,
	0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x11, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x56, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x76, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x1a, 0x55, 0x0a, 0x07, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x4a, 0x0a, 0x0f,
	0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x6d, 0x6f, 0x6e, 0x69,
	0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x69,
	0x6e, 0x67, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x0e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69,
	0x6e, 0x67, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x1a, 0x54, 0x0a, 0x07, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x64, 0x12, 0x2a, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x16, 0xb2, 0xda, 0x21, 0x12, 0x0a, 0x10, 0x0a, 0x0e, 0x41, 0x6c, 0x65, 0x72, 0x74,
	0x69, 0x6e, 0x67, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x76, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x3a, 0x14,
	0x9a, 0xd9, 0x21, 0x10, 0x0a, 0x0e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x42, 0x0d, 0x0a, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x42, 0x8f, 0x01, 0xe8, 0xde, 0x21, 0x00, 0x0a, 0x18, 0x63, 0x6f, 0x6d, 0x2e,
	0x6e, 0x74, 0x74, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x70,
	0x62, 0x2e, 0x76, 0x33, 0x42, 0x19, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x00, 0x5a, 0x52, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6d, 0x6f,
	0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x3b, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	edgelq_monitoring_proto_v3_alerting_policy_change_proto_rawDescOnce sync.Once
	edgelq_monitoring_proto_v3_alerting_policy_change_proto_rawDescData = edgelq_monitoring_proto_v3_alerting_policy_change_proto_rawDesc
)

func edgelq_monitoring_proto_v3_alerting_policy_change_proto_rawDescGZIP() []byte {
	edgelq_monitoring_proto_v3_alerting_policy_change_proto_rawDescOnce.Do(func() {
		edgelq_monitoring_proto_v3_alerting_policy_change_proto_rawDescData = protoimpl.X.CompressGZIP(edgelq_monitoring_proto_v3_alerting_policy_change_proto_rawDescData)
	})
	return edgelq_monitoring_proto_v3_alerting_policy_change_proto_rawDescData
}

var edgelq_monitoring_proto_v3_alerting_policy_change_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var edgelq_monitoring_proto_v3_alerting_policy_change_proto_goTypes = []interface{}{
	(*AlertingPolicyChange)(nil),          // 0: ntt.monitoring.v3.AlertingPolicyChange
	(*AlertingPolicyChange_Added)(nil),    // 1: ntt.monitoring.v3.AlertingPolicyChange.Added
	(*AlertingPolicyChange_Modified)(nil), // 2: ntt.monitoring.v3.AlertingPolicyChange.Modified
	(*AlertingPolicyChange_Current)(nil),  // 3: ntt.monitoring.v3.AlertingPolicyChange.Current
	(*AlertingPolicyChange_Removed)(nil),  // 4: ntt.monitoring.v3.AlertingPolicyChange.Removed
	(*AlertingPolicy)(nil),                // 5: ntt.monitoring.v3.AlertingPolicy
	(*AlertingPolicy_FieldMask)(nil),      // 6: ntt.monitoring.v3.AlertingPolicy_FieldMask
}
var edgelq_monitoring_proto_v3_alerting_policy_change_proto_depIdxs = []int32{
	1, // 0: ntt.monitoring.v3.AlertingPolicyChange.added:type_name -> ntt.monitoring.v3.AlertingPolicyChange.Added
	2, // 1: ntt.monitoring.v3.AlertingPolicyChange.modified:type_name -> ntt.monitoring.v3.AlertingPolicyChange.Modified
	3, // 2: ntt.monitoring.v3.AlertingPolicyChange.current:type_name -> ntt.monitoring.v3.AlertingPolicyChange.Current
	4, // 3: ntt.monitoring.v3.AlertingPolicyChange.removed:type_name -> ntt.monitoring.v3.AlertingPolicyChange.Removed
	5, // 4: ntt.monitoring.v3.AlertingPolicyChange.Added.alerting_policy:type_name -> ntt.monitoring.v3.AlertingPolicy
	5, // 5: ntt.monitoring.v3.AlertingPolicyChange.Modified.alerting_policy:type_name -> ntt.monitoring.v3.AlertingPolicy
	6, // 6: ntt.monitoring.v3.AlertingPolicyChange.Modified.field_mask:type_name -> ntt.monitoring.v3.AlertingPolicy_FieldMask
	5, // 7: ntt.monitoring.v3.AlertingPolicyChange.Current.alerting_policy:type_name -> ntt.monitoring.v3.AlertingPolicy
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { edgelq_monitoring_proto_v3_alerting_policy_change_proto_init() }
func edgelq_monitoring_proto_v3_alerting_policy_change_proto_init() {
	if edgelq_monitoring_proto_v3_alerting_policy_change_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		edgelq_monitoring_proto_v3_alerting_policy_change_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlertingPolicyChange); i {
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
		edgelq_monitoring_proto_v3_alerting_policy_change_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlertingPolicyChange_Added); i {
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
		edgelq_monitoring_proto_v3_alerting_policy_change_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlertingPolicyChange_Modified); i {
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
		edgelq_monitoring_proto_v3_alerting_policy_change_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlertingPolicyChange_Current); i {
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
		edgelq_monitoring_proto_v3_alerting_policy_change_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlertingPolicyChange_Removed); i {
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

	edgelq_monitoring_proto_v3_alerting_policy_change_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*AlertingPolicyChange_Added_)(nil),
		(*AlertingPolicyChange_Modified_)(nil),
		(*AlertingPolicyChange_Current_)(nil),
		(*AlertingPolicyChange_Removed_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: edgelq_monitoring_proto_v3_alerting_policy_change_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           edgelq_monitoring_proto_v3_alerting_policy_change_proto_goTypes,
		DependencyIndexes: edgelq_monitoring_proto_v3_alerting_policy_change_proto_depIdxs,
		MessageInfos:      edgelq_monitoring_proto_v3_alerting_policy_change_proto_msgTypes,
	}.Build()
	edgelq_monitoring_proto_v3_alerting_policy_change_proto = out.File
	edgelq_monitoring_proto_v3_alerting_policy_change_proto_rawDesc = nil
	edgelq_monitoring_proto_v3_alerting_policy_change_proto_goTypes = nil
	edgelq_monitoring_proto_v3_alerting_policy_change_proto_depIdxs = nil
}
