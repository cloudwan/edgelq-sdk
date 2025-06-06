// Code generated by protoc-gen-goten-go
// File: edgelq/monitoring/proto/v4/alert_change.proto
// DO NOT EDIT!!!

package alert

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
	alerting_condition "github.com/cloudwan/edgelq-sdk/monitoring/resources/v4/alerting_condition"
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
	_ = &alerting_condition.AlertingCondition{}
	_ = &fieldmaskpb.FieldMask{}
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// AlertChange is used by Watch notifications Responses to describe change of
// single Alert One of Added, Modified, Removed
type AlertChange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Alert change
	//
	// Types that are valid to be assigned to ChangeType:
	//	*AlertChange_Added_
	//	*AlertChange_Modified_
	//	*AlertChange_Current_
	//	*AlertChange_Removed_
	ChangeType isAlertChange_ChangeType `protobuf_oneof:"change_type"`
}

func (m *AlertChange) Reset() {
	*m = AlertChange{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_monitoring_proto_v4_alert_change_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *AlertChange) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*AlertChange) ProtoMessage() {}

func (m *AlertChange) ProtoReflect() preflect.Message {
	mi := &edgelq_monitoring_proto_v4_alert_change_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*AlertChange) GotenMessage() {}

// Deprecated, Use AlertChange.ProtoReflect.Descriptor instead.
func (*AlertChange) Descriptor() ([]byte, []int) {
	return edgelq_monitoring_proto_v4_alert_change_proto_rawDescGZIP(), []int{0}
}

func (m *AlertChange) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *AlertChange) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *AlertChange) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *AlertChange) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

type isAlertChange_ChangeType interface {
	isAlertChange_ChangeType()
}

type AlertChange_Added_ struct {
	// Added is returned when watched document is added, either created or
	// enters Query view
	Added *AlertChange_Added `protobuf:"bytes,1,opt,name=added,proto3,oneof"`
}
type AlertChange_Modified_ struct {
	// Modified is returned when watched document is modified
	Modified *AlertChange_Modified `protobuf:"bytes,2,opt,name=modified,proto3,oneof"`
}
type AlertChange_Current_ struct {
	// Current is returned in stateless watch when document enters query view or
	// is modified within.
	Current *AlertChange_Current `protobuf:"bytes,4,opt,name=current,proto3,oneof"`
}
type AlertChange_Removed_ struct {
	// Removed is returned when Alert is deleted or leaves Query view
	Removed *AlertChange_Removed `protobuf:"bytes,3,opt,name=removed,proto3,oneof"`
}

func (*AlertChange_Added_) isAlertChange_ChangeType()    {}
func (*AlertChange_Modified_) isAlertChange_ChangeType() {}
func (*AlertChange_Current_) isAlertChange_ChangeType()  {}
func (*AlertChange_Removed_) isAlertChange_ChangeType()  {}
func (m *AlertChange) GetChangeType() isAlertChange_ChangeType {
	if m != nil {
		return m.ChangeType
	}
	return nil
}
func (m *AlertChange) GetAdded() *AlertChange_Added {
	if x, ok := m.GetChangeType().(*AlertChange_Added_); ok {
		return x.Added
	}
	return nil
}
func (m *AlertChange) GetModified() *AlertChange_Modified {
	if x, ok := m.GetChangeType().(*AlertChange_Modified_); ok {
		return x.Modified
	}
	return nil
}
func (m *AlertChange) GetCurrent() *AlertChange_Current {
	if x, ok := m.GetChangeType().(*AlertChange_Current_); ok {
		return x.Current
	}
	return nil
}
func (m *AlertChange) GetRemoved() *AlertChange_Removed {
	if x, ok := m.GetChangeType().(*AlertChange_Removed_); ok {
		return x.Removed
	}
	return nil
}
func (m *AlertChange) SetChangeType(ofv isAlertChange_ChangeType) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "isAlertChange_ChangeType", "AlertChange"))
	}
	m.ChangeType = ofv
}
func (m *AlertChange) SetAdded(fv *AlertChange_Added) {
	m.SetChangeType(&AlertChange_Added_{Added: fv})
}
func (m *AlertChange) SetModified(fv *AlertChange_Modified) {
	m.SetChangeType(&AlertChange_Modified_{Modified: fv})
}
func (m *AlertChange) SetCurrent(fv *AlertChange_Current) {
	m.SetChangeType(&AlertChange_Current_{Current: fv})
}
func (m *AlertChange) SetRemoved(fv *AlertChange_Removed) {
	m.SetChangeType(&AlertChange_Removed_{Removed: fv})
}

// Alert has been added to query view
type AlertChange_Added struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Alert         *Alert `protobuf:"bytes,1,opt,name=alert,proto3" json:"alert,omitempty"`
	// Integer describing index of added Alert in resulting query view.
	ViewIndex int32 `protobuf:"varint,2,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty"`
}

func (m *AlertChange_Added) Reset() {
	*m = AlertChange_Added{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_monitoring_proto_v4_alert_change_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *AlertChange_Added) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*AlertChange_Added) ProtoMessage() {}

func (m *AlertChange_Added) ProtoReflect() preflect.Message {
	mi := &edgelq_monitoring_proto_v4_alert_change_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*AlertChange_Added) GotenMessage() {}

// Deprecated, Use AlertChange_Added.ProtoReflect.Descriptor instead.
func (*AlertChange_Added) Descriptor() ([]byte, []int) {
	return edgelq_monitoring_proto_v4_alert_change_proto_rawDescGZIP(), []int{0, 0}
}

func (m *AlertChange_Added) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *AlertChange_Added) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *AlertChange_Added) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *AlertChange_Added) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *AlertChange_Added) GetAlert() *Alert {
	if m != nil {
		return m.Alert
	}
	return nil
}

func (m *AlertChange_Added) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *AlertChange_Added) SetAlert(fv *Alert) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Alert", "AlertChange_Added"))
	}
	m.Alert = fv
}

func (m *AlertChange_Added) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "AlertChange_Added"))
	}
	m.ViewIndex = fv
}

// Alert changed some of it's fields - contains either full document or masked
// change
type AlertChange_Modified struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Name of modified Alert
	Name *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty"`
	// New version of Alert or masked difference, depending on mask_changes
	// instrumentation of issued [WatchAlertRequest] or [WatchAlertsRequest]
	Alert *Alert `protobuf:"bytes,2,opt,name=alert,proto3" json:"alert,omitempty"`
	// Used when mask_changes is set, contains field paths of modified
	// properties.
	FieldMask *Alert_FieldMask `protobuf:"bytes,3,opt,customtype=Alert_FieldMask,name=field_mask,json=fieldMask,proto3" json:"field_mask,omitempty"`
	// Previous view index specifies previous position of modified Alert.
	// When modification doesn't affect sorted order, value will remain
	// identical to [view_index].
	PreviousViewIndex int32 `protobuf:"varint,4,opt,name=previous_view_index,json=previousViewIndex,proto3" json:"previous_view_index,omitempty"`
	// Integer specifying Alert new index in resulting query view.
	ViewIndex int32 `protobuf:"varint,5,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty"`
}

func (m *AlertChange_Modified) Reset() {
	*m = AlertChange_Modified{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_monitoring_proto_v4_alert_change_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *AlertChange_Modified) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*AlertChange_Modified) ProtoMessage() {}

func (m *AlertChange_Modified) ProtoReflect() preflect.Message {
	mi := &edgelq_monitoring_proto_v4_alert_change_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*AlertChange_Modified) GotenMessage() {}

// Deprecated, Use AlertChange_Modified.ProtoReflect.Descriptor instead.
func (*AlertChange_Modified) Descriptor() ([]byte, []int) {
	return edgelq_monitoring_proto_v4_alert_change_proto_rawDescGZIP(), []int{0, 1}
}

func (m *AlertChange_Modified) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *AlertChange_Modified) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *AlertChange_Modified) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *AlertChange_Modified) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *AlertChange_Modified) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *AlertChange_Modified) GetAlert() *Alert {
	if m != nil {
		return m.Alert
	}
	return nil
}

func (m *AlertChange_Modified) GetFieldMask() *Alert_FieldMask {
	if m != nil {
		return m.FieldMask
	}
	return nil
}

func (m *AlertChange_Modified) GetPreviousViewIndex() int32 {
	if m != nil {
		return m.PreviousViewIndex
	}
	return int32(0)
}

func (m *AlertChange_Modified) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *AlertChange_Modified) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "AlertChange_Modified"))
	}
	m.Name = fv
}

func (m *AlertChange_Modified) SetAlert(fv *Alert) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Alert", "AlertChange_Modified"))
	}
	m.Alert = fv
}

func (m *AlertChange_Modified) SetFieldMask(fv *Alert_FieldMask) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "FieldMask", "AlertChange_Modified"))
	}
	m.FieldMask = fv
}

func (m *AlertChange_Modified) SetPreviousViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "PreviousViewIndex", "AlertChange_Modified"))
	}
	m.PreviousViewIndex = fv
}

func (m *AlertChange_Modified) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "AlertChange_Modified"))
	}
	m.ViewIndex = fv
}

// Alert has been added or modified in a query view. Version used for
// stateless watching
type AlertChange_Current struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Alert         *Alert `protobuf:"bytes,1,opt,name=alert,proto3" json:"alert,omitempty"`
}

func (m *AlertChange_Current) Reset() {
	*m = AlertChange_Current{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_monitoring_proto_v4_alert_change_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *AlertChange_Current) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*AlertChange_Current) ProtoMessage() {}

func (m *AlertChange_Current) ProtoReflect() preflect.Message {
	mi := &edgelq_monitoring_proto_v4_alert_change_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*AlertChange_Current) GotenMessage() {}

// Deprecated, Use AlertChange_Current.ProtoReflect.Descriptor instead.
func (*AlertChange_Current) Descriptor() ([]byte, []int) {
	return edgelq_monitoring_proto_v4_alert_change_proto_rawDescGZIP(), []int{0, 2}
}

func (m *AlertChange_Current) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *AlertChange_Current) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *AlertChange_Current) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *AlertChange_Current) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *AlertChange_Current) GetAlert() *Alert {
	if m != nil {
		return m.Alert
	}
	return nil
}

func (m *AlertChange_Current) SetAlert(fv *Alert) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Alert", "AlertChange_Current"))
	}
	m.Alert = fv
}

// Removed is returned when Alert is deleted or leaves Query view
type AlertChange_Removed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Name          *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty"`
	// Integer specifying removed Alert index. Not populated in stateless watch
	// type.
	ViewIndex int32 `protobuf:"varint,2,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty"`
}

func (m *AlertChange_Removed) Reset() {
	*m = AlertChange_Removed{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_monitoring_proto_v4_alert_change_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *AlertChange_Removed) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*AlertChange_Removed) ProtoMessage() {}

func (m *AlertChange_Removed) ProtoReflect() preflect.Message {
	mi := &edgelq_monitoring_proto_v4_alert_change_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*AlertChange_Removed) GotenMessage() {}

// Deprecated, Use AlertChange_Removed.ProtoReflect.Descriptor instead.
func (*AlertChange_Removed) Descriptor() ([]byte, []int) {
	return edgelq_monitoring_proto_v4_alert_change_proto_rawDescGZIP(), []int{0, 3}
}

func (m *AlertChange_Removed) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *AlertChange_Removed) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *AlertChange_Removed) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *AlertChange_Removed) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *AlertChange_Removed) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *AlertChange_Removed) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *AlertChange_Removed) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "AlertChange_Removed"))
	}
	m.Name = fv
}

func (m *AlertChange_Removed) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "AlertChange_Removed"))
	}
	m.ViewIndex = fv
}

var edgelq_monitoring_proto_v4_alert_change_proto preflect.FileDescriptor

var edgelq_monitoring_proto_v4_alert_change_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x69, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x34, 0x2f, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x11, 0x6e, 0x74, 0x74, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x34, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20,
	0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x26, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x69, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x34, 0x2f, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8f, 0x06, 0x0a, 0x0b, 0x41, 0x6c, 0x65,
	0x72, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x3c, 0x0a, 0x05, 0x61, 0x64, 0x64, 0x65,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x6d, 0x6f,
	0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x34, 0x2e, 0x41, 0x6c, 0x65, 0x72,
	0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x65, 0x64, 0x48, 0x00, 0x52,
	0x05, 0x61, 0x64, 0x64, 0x65, 0x64, 0x12, 0x45, 0x0a, 0x08, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69,
	0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x6d,
	0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x34, 0x2e, 0x41, 0x6c, 0x65,
	0x72, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65,
	0x64, 0x48, 0x00, 0x52, 0x08, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x42, 0x0a,
	0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26,
	0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x34, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x12, 0x42, 0x0a, 0x07, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x69, 0x6e, 0x67, 0x2e, 0x76, 0x34, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x48, 0x00, 0x52, 0x07, 0x72, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x64, 0x1a, 0x56, 0x0a, 0x05, 0x41, 0x64, 0x64, 0x65, 0x64, 0x12, 0x2e,
	0x0a, 0x05, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x6e, 0x74, 0x74, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76,
	0x34, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x05, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x76, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x1a, 0xf6, 0x01,
	0x0a, 0x08, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x21, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0d, 0xb2, 0xda, 0x21, 0x09, 0x0a, 0x07,
	0x0a, 0x05, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a,
	0x05, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6e,
	0x74, 0x74, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x34,
	0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x05, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x12, 0x48, 0x0a,
	0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x42, 0x0d, 0xb2,
	0xda, 0x21, 0x09, 0x32, 0x07, 0x0a, 0x05, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x09, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x2e, 0x0a, 0x13, 0x70, 0x72, 0x65, 0x76, 0x69,
	0x6f, 0x75, 0x73, 0x5f, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x56, 0x69,
	0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x69, 0x65, 0x77, 0x5f,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x76, 0x69, 0x65,
	0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x1a, 0x39, 0x0a, 0x07, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x12, 0x2e, 0x0a, 0x05, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x34, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x05, 0x61, 0x6c, 0x65, 0x72,
	0x74, 0x1a, 0x4b, 0x0a, 0x07, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x12, 0x21, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0d, 0xb2, 0xda, 0x21, 0x09,
	0x0a, 0x07, 0x0a, 0x05, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x76, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x3a, 0x0b,
	0x9a, 0xd9, 0x21, 0x07, 0x0a, 0x05, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x42, 0x0d, 0x0a, 0x0b, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x42, 0x72, 0xe8, 0xde, 0x21, 0x00,
	0x0a, 0x18, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f,
	0x72, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x34, 0x42, 0x10, 0x41, 0x6c, 0x65, 0x72,
	0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x00, 0x5a, 0x3e,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74,
	0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f,
	0x76, 0x34, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x3b, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	edgelq_monitoring_proto_v4_alert_change_proto_rawDescOnce sync.Once
	edgelq_monitoring_proto_v4_alert_change_proto_rawDescData = edgelq_monitoring_proto_v4_alert_change_proto_rawDesc
)

func edgelq_monitoring_proto_v4_alert_change_proto_rawDescGZIP() []byte {
	edgelq_monitoring_proto_v4_alert_change_proto_rawDescOnce.Do(func() {
		edgelq_monitoring_proto_v4_alert_change_proto_rawDescData = protoimpl.X.CompressGZIP(edgelq_monitoring_proto_v4_alert_change_proto_rawDescData)
	})
	return edgelq_monitoring_proto_v4_alert_change_proto_rawDescData
}

var edgelq_monitoring_proto_v4_alert_change_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var edgelq_monitoring_proto_v4_alert_change_proto_goTypes = []interface{}{
	(*AlertChange)(nil),          // 0: ntt.monitoring.v4.AlertChange
	(*AlertChange_Added)(nil),    // 1: ntt.monitoring.v4.AlertChange.Added
	(*AlertChange_Modified)(nil), // 2: ntt.monitoring.v4.AlertChange.Modified
	(*AlertChange_Current)(nil),  // 3: ntt.monitoring.v4.AlertChange.Current
	(*AlertChange_Removed)(nil),  // 4: ntt.monitoring.v4.AlertChange.Removed
	(*Alert)(nil),                // 5: ntt.monitoring.v4.Alert
	(*Alert_FieldMask)(nil),      // 6: ntt.monitoring.v4.Alert_FieldMask
}
var edgelq_monitoring_proto_v4_alert_change_proto_depIdxs = []int32{
	1, // 0: ntt.monitoring.v4.AlertChange.added:type_name -> ntt.monitoring.v4.AlertChange.Added
	2, // 1: ntt.monitoring.v4.AlertChange.modified:type_name -> ntt.monitoring.v4.AlertChange.Modified
	3, // 2: ntt.monitoring.v4.AlertChange.current:type_name -> ntt.monitoring.v4.AlertChange.Current
	4, // 3: ntt.monitoring.v4.AlertChange.removed:type_name -> ntt.monitoring.v4.AlertChange.Removed
	5, // 4: ntt.monitoring.v4.AlertChange.Added.alert:type_name -> ntt.monitoring.v4.Alert
	5, // 5: ntt.monitoring.v4.AlertChange.Modified.alert:type_name -> ntt.monitoring.v4.Alert
	6, // 6: ntt.monitoring.v4.AlertChange.Modified.field_mask:type_name -> ntt.monitoring.v4.Alert_FieldMask
	5, // 7: ntt.monitoring.v4.AlertChange.Current.alert:type_name -> ntt.monitoring.v4.Alert
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { edgelq_monitoring_proto_v4_alert_change_proto_init() }
func edgelq_monitoring_proto_v4_alert_change_proto_init() {
	if edgelq_monitoring_proto_v4_alert_change_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		edgelq_monitoring_proto_v4_alert_change_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlertChange); i {
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
		edgelq_monitoring_proto_v4_alert_change_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlertChange_Added); i {
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
		edgelq_monitoring_proto_v4_alert_change_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlertChange_Modified); i {
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
		edgelq_monitoring_proto_v4_alert_change_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlertChange_Current); i {
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
		edgelq_monitoring_proto_v4_alert_change_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlertChange_Removed); i {
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

	edgelq_monitoring_proto_v4_alert_change_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*AlertChange_Added_)(nil),
		(*AlertChange_Modified_)(nil),
		(*AlertChange_Current_)(nil),
		(*AlertChange_Removed_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: edgelq_monitoring_proto_v4_alert_change_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           edgelq_monitoring_proto_v4_alert_change_proto_goTypes,
		DependencyIndexes: edgelq_monitoring_proto_v4_alert_change_proto_depIdxs,
		MessageInfos:      edgelq_monitoring_proto_v4_alert_change_proto_msgTypes,
	}.Build()
	edgelq_monitoring_proto_v4_alert_change_proto = out.File
	edgelq_monitoring_proto_v4_alert_change_proto_rawDesc = nil
	edgelq_monitoring_proto_v4_alert_change_proto_goTypes = nil
	edgelq_monitoring_proto_v4_alert_change_proto_depIdxs = nil
}
