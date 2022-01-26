// Code generated by protoc-gen-goten-go
// File: edgelq/monitoring/proto/v3/phantom_time_serie_change.proto
// DO NOT EDIT!!!

package phantom_time_serie

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

// PhantomTimeSerieChange is used by Watch notifications Responses to describe
// change of single PhantomTimeSerie One of Added, Modified, Removed
type PhantomTimeSerieChange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// PhantomTimeSerie change
	//
	// Types that are valid to be assigned to ChangeType:
	//	*PhantomTimeSerieChange_Added_
	//	*PhantomTimeSerieChange_Modified_
	//	*PhantomTimeSerieChange_Current_
	//	*PhantomTimeSerieChange_Removed_
	ChangeType isPhantomTimeSerieChange_ChangeType `protobuf_oneof:"change_type"`
}

func (m *PhantomTimeSerieChange) Reset() {
	*m = PhantomTimeSerieChange{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_monitoring_proto_v3_phantom_time_serie_change_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *PhantomTimeSerieChange) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*PhantomTimeSerieChange) ProtoMessage() {}

func (m *PhantomTimeSerieChange) ProtoReflect() preflect.Message {
	mi := &edgelq_monitoring_proto_v3_phantom_time_serie_change_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*PhantomTimeSerieChange) GotenMessage() {}

// Deprecated, Use PhantomTimeSerieChange.ProtoReflect.Descriptor instead.
func (*PhantomTimeSerieChange) Descriptor() ([]byte, []int) {
	return edgelq_monitoring_proto_v3_phantom_time_serie_change_proto_rawDescGZIP(), []int{0}
}

func (m *PhantomTimeSerieChange) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *PhantomTimeSerieChange) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *PhantomTimeSerieChange) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *PhantomTimeSerieChange) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

type isPhantomTimeSerieChange_ChangeType interface {
	isPhantomTimeSerieChange_ChangeType()
}

type PhantomTimeSerieChange_Added_ struct {
	// Added is returned when watched document is added, either created or
	// enters Query view
	Added *PhantomTimeSerieChange_Added `protobuf:"bytes,1,opt,name=added,proto3,oneof" firestore:"added"`
}
type PhantomTimeSerieChange_Modified_ struct {
	// Modified is returned when watched document is modified
	Modified *PhantomTimeSerieChange_Modified `protobuf:"bytes,2,opt,name=modified,proto3,oneof" firestore:"modified"`
}
type PhantomTimeSerieChange_Current_ struct {
	// Current is returned in stateless watch when document enters query view or
	// is modified within.
	Current *PhantomTimeSerieChange_Current `protobuf:"bytes,4,opt,name=current,proto3,oneof" firestore:"current"`
}
type PhantomTimeSerieChange_Removed_ struct {
	// Removed is returned when PhantomTimeSerie is deleted or leaves Query view
	Removed *PhantomTimeSerieChange_Removed `protobuf:"bytes,3,opt,name=removed,proto3,oneof" firestore:"removed"`
}

func (*PhantomTimeSerieChange_Added_) isPhantomTimeSerieChange_ChangeType()    {}
func (*PhantomTimeSerieChange_Modified_) isPhantomTimeSerieChange_ChangeType() {}
func (*PhantomTimeSerieChange_Current_) isPhantomTimeSerieChange_ChangeType()  {}
func (*PhantomTimeSerieChange_Removed_) isPhantomTimeSerieChange_ChangeType()  {}
func (m *PhantomTimeSerieChange) GetChangeType() isPhantomTimeSerieChange_ChangeType {
	if m != nil {
		return m.ChangeType
	}
	return nil
}
func (m *PhantomTimeSerieChange) GetAdded() *PhantomTimeSerieChange_Added {
	if x, ok := m.GetChangeType().(*PhantomTimeSerieChange_Added_); ok {
		return x.Added
	}
	return nil
}
func (m *PhantomTimeSerieChange) GetModified() *PhantomTimeSerieChange_Modified {
	if x, ok := m.GetChangeType().(*PhantomTimeSerieChange_Modified_); ok {
		return x.Modified
	}
	return nil
}
func (m *PhantomTimeSerieChange) GetCurrent() *PhantomTimeSerieChange_Current {
	if x, ok := m.GetChangeType().(*PhantomTimeSerieChange_Current_); ok {
		return x.Current
	}
	return nil
}
func (m *PhantomTimeSerieChange) GetRemoved() *PhantomTimeSerieChange_Removed {
	if x, ok := m.GetChangeType().(*PhantomTimeSerieChange_Removed_); ok {
		return x.Removed
	}
	return nil
}
func (m *PhantomTimeSerieChange) SetChangeType(ofv isPhantomTimeSerieChange_ChangeType) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "isPhantomTimeSerieChange_ChangeType", "PhantomTimeSerieChange"))
	}
	m.ChangeType = ofv
}
func (m *PhantomTimeSerieChange) SetAdded(fv *PhantomTimeSerieChange_Added) {
	m.SetChangeType(&PhantomTimeSerieChange_Added_{Added: fv})
}
func (m *PhantomTimeSerieChange) SetModified(fv *PhantomTimeSerieChange_Modified) {
	m.SetChangeType(&PhantomTimeSerieChange_Modified_{Modified: fv})
}
func (m *PhantomTimeSerieChange) SetCurrent(fv *PhantomTimeSerieChange_Current) {
	m.SetChangeType(&PhantomTimeSerieChange_Current_{Current: fv})
}
func (m *PhantomTimeSerieChange) SetRemoved(fv *PhantomTimeSerieChange_Removed) {
	m.SetChangeType(&PhantomTimeSerieChange_Removed_{Removed: fv})
}

// PhantomTimeSerie has been added to query view
type PhantomTimeSerieChange_Added struct {
	state            protoimpl.MessageState
	sizeCache        protoimpl.SizeCache
	unknownFields    protoimpl.UnknownFields
	PhantomTimeSerie *PhantomTimeSerie `protobuf:"bytes,1,opt,name=phantom_time_serie,json=phantomTimeSerie,proto3" json:"phantom_time_serie,omitempty" firestore:"phantomTimeSerie"`
	// Integer describing index of added PhantomTimeSerie in resulting query
	// view.
	ViewIndex int32 `protobuf:"varint,2,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty" firestore:"viewIndex"`
}

func (m *PhantomTimeSerieChange_Added) Reset() {
	*m = PhantomTimeSerieChange_Added{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_monitoring_proto_v3_phantom_time_serie_change_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *PhantomTimeSerieChange_Added) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*PhantomTimeSerieChange_Added) ProtoMessage() {}

func (m *PhantomTimeSerieChange_Added) ProtoReflect() preflect.Message {
	mi := &edgelq_monitoring_proto_v3_phantom_time_serie_change_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*PhantomTimeSerieChange_Added) GotenMessage() {}

// Deprecated, Use PhantomTimeSerieChange_Added.ProtoReflect.Descriptor instead.
func (*PhantomTimeSerieChange_Added) Descriptor() ([]byte, []int) {
	return edgelq_monitoring_proto_v3_phantom_time_serie_change_proto_rawDescGZIP(), []int{0, 0}
}

func (m *PhantomTimeSerieChange_Added) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *PhantomTimeSerieChange_Added) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *PhantomTimeSerieChange_Added) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *PhantomTimeSerieChange_Added) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *PhantomTimeSerieChange_Added) GetPhantomTimeSerie() *PhantomTimeSerie {
	if m != nil {
		return m.PhantomTimeSerie
	}
	return nil
}

func (m *PhantomTimeSerieChange_Added) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *PhantomTimeSerieChange_Added) SetPhantomTimeSerie(fv *PhantomTimeSerie) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "PhantomTimeSerie", "PhantomTimeSerieChange_Added"))
	}
	m.PhantomTimeSerie = fv
}

func (m *PhantomTimeSerieChange_Added) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "PhantomTimeSerieChange_Added"))
	}
	m.ViewIndex = fv
}

// PhantomTimeSerie changed some of it's fields - contains either full
// document or masked change
type PhantomTimeSerieChange_Modified struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Name of modified PhantomTimeSerie
	Name *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// New version of PhantomTimeSerie or masked difference, depending on
	// mask_changes instrumentation of issued [WatchPhantomTimeSerieRequest] or
	// [WatchPhantomTimeSeriesRequest]
	PhantomTimeSerie *PhantomTimeSerie `protobuf:"bytes,2,opt,name=phantom_time_serie,json=phantomTimeSerie,proto3" json:"phantom_time_serie,omitempty" firestore:"phantomTimeSerie"`
	// Used when mask_changes is set, contains field paths of modified
	// properties.
	FieldMask *PhantomTimeSerie_FieldMask `protobuf:"bytes,3,opt,customtype=PhantomTimeSerie_FieldMask,name=field_mask,json=fieldMask,proto3" json:"field_mask,omitempty" firestore:"fieldMask"`
	// Previous view index specifies previous position of modified
	// PhantomTimeSerie. When modification doesn't affect sorted order, value
	// will remain identical to [view_index].
	PreviousViewIndex int32 `protobuf:"varint,4,opt,name=previous_view_index,json=previousViewIndex,proto3" json:"previous_view_index,omitempty" firestore:"previousViewIndex"`
	// Integer specifying PhantomTimeSerie new index in resulting query view.
	ViewIndex int32 `protobuf:"varint,5,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty" firestore:"viewIndex"`
}

func (m *PhantomTimeSerieChange_Modified) Reset() {
	*m = PhantomTimeSerieChange_Modified{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_monitoring_proto_v3_phantom_time_serie_change_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *PhantomTimeSerieChange_Modified) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*PhantomTimeSerieChange_Modified) ProtoMessage() {}

func (m *PhantomTimeSerieChange_Modified) ProtoReflect() preflect.Message {
	mi := &edgelq_monitoring_proto_v3_phantom_time_serie_change_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*PhantomTimeSerieChange_Modified) GotenMessage() {}

// Deprecated, Use PhantomTimeSerieChange_Modified.ProtoReflect.Descriptor instead.
func (*PhantomTimeSerieChange_Modified) Descriptor() ([]byte, []int) {
	return edgelq_monitoring_proto_v3_phantom_time_serie_change_proto_rawDescGZIP(), []int{0, 1}
}

func (m *PhantomTimeSerieChange_Modified) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *PhantomTimeSerieChange_Modified) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *PhantomTimeSerieChange_Modified) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *PhantomTimeSerieChange_Modified) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *PhantomTimeSerieChange_Modified) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *PhantomTimeSerieChange_Modified) GetPhantomTimeSerie() *PhantomTimeSerie {
	if m != nil {
		return m.PhantomTimeSerie
	}
	return nil
}

func (m *PhantomTimeSerieChange_Modified) GetFieldMask() *PhantomTimeSerie_FieldMask {
	if m != nil {
		return m.FieldMask
	}
	return nil
}

func (m *PhantomTimeSerieChange_Modified) GetPreviousViewIndex() int32 {
	if m != nil {
		return m.PreviousViewIndex
	}
	return int32(0)
}

func (m *PhantomTimeSerieChange_Modified) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *PhantomTimeSerieChange_Modified) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "PhantomTimeSerieChange_Modified"))
	}
	m.Name = fv
}

func (m *PhantomTimeSerieChange_Modified) SetPhantomTimeSerie(fv *PhantomTimeSerie) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "PhantomTimeSerie", "PhantomTimeSerieChange_Modified"))
	}
	m.PhantomTimeSerie = fv
}

func (m *PhantomTimeSerieChange_Modified) SetFieldMask(fv *PhantomTimeSerie_FieldMask) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "FieldMask", "PhantomTimeSerieChange_Modified"))
	}
	m.FieldMask = fv
}

func (m *PhantomTimeSerieChange_Modified) SetPreviousViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "PreviousViewIndex", "PhantomTimeSerieChange_Modified"))
	}
	m.PreviousViewIndex = fv
}

func (m *PhantomTimeSerieChange_Modified) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "PhantomTimeSerieChange_Modified"))
	}
	m.ViewIndex = fv
}

// PhantomTimeSerie has been added or modified in a query view. Version used
// for stateless watching
type PhantomTimeSerieChange_Current struct {
	state            protoimpl.MessageState
	sizeCache        protoimpl.SizeCache
	unknownFields    protoimpl.UnknownFields
	PhantomTimeSerie *PhantomTimeSerie `protobuf:"bytes,1,opt,name=phantom_time_serie,json=phantomTimeSerie,proto3" json:"phantom_time_serie,omitempty" firestore:"phantomTimeSerie"`
}

func (m *PhantomTimeSerieChange_Current) Reset() {
	*m = PhantomTimeSerieChange_Current{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_monitoring_proto_v3_phantom_time_serie_change_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *PhantomTimeSerieChange_Current) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*PhantomTimeSerieChange_Current) ProtoMessage() {}

func (m *PhantomTimeSerieChange_Current) ProtoReflect() preflect.Message {
	mi := &edgelq_monitoring_proto_v3_phantom_time_serie_change_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*PhantomTimeSerieChange_Current) GotenMessage() {}

// Deprecated, Use PhantomTimeSerieChange_Current.ProtoReflect.Descriptor instead.
func (*PhantomTimeSerieChange_Current) Descriptor() ([]byte, []int) {
	return edgelq_monitoring_proto_v3_phantom_time_serie_change_proto_rawDescGZIP(), []int{0, 2}
}

func (m *PhantomTimeSerieChange_Current) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *PhantomTimeSerieChange_Current) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *PhantomTimeSerieChange_Current) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *PhantomTimeSerieChange_Current) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *PhantomTimeSerieChange_Current) GetPhantomTimeSerie() *PhantomTimeSerie {
	if m != nil {
		return m.PhantomTimeSerie
	}
	return nil
}

func (m *PhantomTimeSerieChange_Current) SetPhantomTimeSerie(fv *PhantomTimeSerie) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "PhantomTimeSerie", "PhantomTimeSerieChange_Current"))
	}
	m.PhantomTimeSerie = fv
}

// Removed is returned when PhantomTimeSerie is deleted or leaves Query view
type PhantomTimeSerieChange_Removed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Name          *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// Integer specifying removed PhantomTimeSerie index. Not populated in
	// stateless watch type.
	ViewIndex int32 `protobuf:"varint,2,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty" firestore:"viewIndex"`
}

func (m *PhantomTimeSerieChange_Removed) Reset() {
	*m = PhantomTimeSerieChange_Removed{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_monitoring_proto_v3_phantom_time_serie_change_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *PhantomTimeSerieChange_Removed) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*PhantomTimeSerieChange_Removed) ProtoMessage() {}

func (m *PhantomTimeSerieChange_Removed) ProtoReflect() preflect.Message {
	mi := &edgelq_monitoring_proto_v3_phantom_time_serie_change_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*PhantomTimeSerieChange_Removed) GotenMessage() {}

// Deprecated, Use PhantomTimeSerieChange_Removed.ProtoReflect.Descriptor instead.
func (*PhantomTimeSerieChange_Removed) Descriptor() ([]byte, []int) {
	return edgelq_monitoring_proto_v3_phantom_time_serie_change_proto_rawDescGZIP(), []int{0, 3}
}

func (m *PhantomTimeSerieChange_Removed) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *PhantomTimeSerieChange_Removed) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *PhantomTimeSerieChange_Removed) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *PhantomTimeSerieChange_Removed) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *PhantomTimeSerieChange_Removed) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *PhantomTimeSerieChange_Removed) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *PhantomTimeSerieChange_Removed) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "PhantomTimeSerieChange_Removed"))
	}
	m.Name = fv
}

func (m *PhantomTimeSerieChange_Removed) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "PhantomTimeSerieChange_Removed"))
	}
	m.ViewIndex = fv
}

var edgelq_monitoring_proto_v3_phantom_time_serie_change_proto preflect.FileDescriptor

var edgelq_monitoring_proto_v3_phantom_time_serie_change_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x69, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x33, 0x2f, 0x70, 0x68, 0x61,
	0x6e, 0x74, 0x6f, 0x6d, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x69, 0x65, 0x5f,
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
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x33, 0x2f, 0x70, 0x68, 0x61, 0x6e, 0x74, 0x6f,
	0x6d, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x69, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xdb, 0x07, 0x0a, 0x16, 0x50, 0x68, 0x61, 0x6e, 0x74, 0x6f, 0x6d, 0x54, 0x69,
	0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x47, 0x0a,
	0x05, 0x61, 0x64, 0x64, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x6e,
	0x74, 0x74, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33,
	0x2e, 0x50, 0x68, 0x61, 0x6e, 0x74, 0x6f, 0x6d, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69,
	0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x65, 0x64, 0x48, 0x00, 0x52,
	0x05, 0x61, 0x64, 0x64, 0x65, 0x64, 0x12, 0x50, 0x0a, 0x08, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69,
	0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x6d,
	0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x50, 0x68, 0x61,
	0x6e, 0x74, 0x6f, 0x6d, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x48, 0x00, 0x52, 0x08,
	0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x4d, 0x0a, 0x07, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x6e, 0x74, 0x74, 0x2e,
	0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x50, 0x68,
	0x61, 0x6e, 0x74, 0x6f, 0x6d, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x07,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x4d, 0x0a, 0x07, 0x72, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x6d,
	0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x50, 0x68, 0x61,
	0x6e, 0x74, 0x6f, 0x6d, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x48, 0x00, 0x52, 0x07, 0x72,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x1a, 0x79, 0x0a, 0x05, 0x41, 0x64, 0x64, 0x65, 0x64, 0x12,
	0x51, 0x0a, 0x12, 0x70, 0x68, 0x61, 0x6e, 0x74, 0x6f, 0x6d, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f,
	0x73, 0x65, 0x72, 0x69, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6e, 0x74,
	0x74, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e,
	0x50, 0x68, 0x61, 0x6e, 0x74, 0x6f, 0x6d, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65,
	0x52, 0x10, 0x70, 0x68, 0x61, 0x6e, 0x74, 0x6f, 0x6d, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72,
	0x69, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x76, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x1a, 0xaf, 0x02, 0x0a, 0x08, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x2c,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xb2, 0xda,
	0x21, 0x14, 0x0a, 0x12, 0x0a, 0x10, 0x50, 0x68, 0x61, 0x6e, 0x74, 0x6f, 0x6d, 0x54, 0x69, 0x6d,
	0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x51, 0x0a, 0x12,
	0x70, 0x68, 0x61, 0x6e, 0x74, 0x6f, 0x6d, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72,
	0x69, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x6d,
	0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x50, 0x68, 0x61,
	0x6e, 0x74, 0x6f, 0x6d, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x52, 0x10, 0x70,
	0x68, 0x61, 0x6e, 0x74, 0x6f, 0x6d, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x12,
	0x53, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x42,
	0x18, 0xb2, 0xda, 0x21, 0x14, 0x32, 0x12, 0x0a, 0x10, 0x50, 0x68, 0x61, 0x6e, 0x74, 0x6f, 0x6d,
	0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x4d, 0x61, 0x73, 0x6b, 0x12, 0x2e, 0x0a, 0x13, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73,
	0x5f, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x11, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x56, 0x69, 0x65, 0x77, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x76, 0x69, 0x65, 0x77, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x1a, 0x5c, 0x0a, 0x07, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x51,
	0x0a, 0x12, 0x70, 0x68, 0x61, 0x6e, 0x74, 0x6f, 0x6d, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73,
	0x65, 0x72, 0x69, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6e, 0x74, 0x74,
	0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x50,
	0x68, 0x61, 0x6e, 0x74, 0x6f, 0x6d, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x52,
	0x10, 0x70, 0x68, 0x61, 0x6e, 0x74, 0x6f, 0x6d, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69,
	0x65, 0x1a, 0x56, 0x0a, 0x07, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x12, 0x2c, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xb2, 0xda, 0x21, 0x14,
	0x0a, 0x12, 0x0a, 0x10, 0x50, 0x68, 0x61, 0x6e, 0x74, 0x6f, 0x6d, 0x54, 0x69, 0x6d, 0x65, 0x53,
	0x65, 0x72, 0x69, 0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x69,
	0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x76, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x3a, 0x16, 0x9a, 0xd9, 0x21, 0x12, 0x0a,
	0x10, 0x50, 0x68, 0x61, 0x6e, 0x74, 0x6f, 0x6d, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69,
	0x65, 0x42, 0x0d, 0x0a, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x42, 0x97, 0x01, 0xe8, 0xde, 0x21, 0x00, 0x0a, 0x18, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x74, 0x74,
	0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x62, 0x2e, 0x76,
	0x33, 0x42, 0x1b, 0x50, 0x68, 0x61, 0x6e, 0x74, 0x6f, 0x6d, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65,
	0x72, 0x69, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x00,
	0x5a, 0x58, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6d, 0x6f, 0x6e,
	0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x2f, 0x76, 0x33, 0x2f, 0x70, 0x68, 0x61, 0x6e, 0x74, 0x6f, 0x6d, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x5f, 0x73, 0x65, 0x72, 0x69, 0x65, 0x3b, 0x70, 0x68, 0x61, 0x6e, 0x74, 0x6f, 0x6d, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x69, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	edgelq_monitoring_proto_v3_phantom_time_serie_change_proto_rawDescOnce sync.Once
	edgelq_monitoring_proto_v3_phantom_time_serie_change_proto_rawDescData = edgelq_monitoring_proto_v3_phantom_time_serie_change_proto_rawDesc
)

func edgelq_monitoring_proto_v3_phantom_time_serie_change_proto_rawDescGZIP() []byte {
	edgelq_monitoring_proto_v3_phantom_time_serie_change_proto_rawDescOnce.Do(func() {
		edgelq_monitoring_proto_v3_phantom_time_serie_change_proto_rawDescData = protoimpl.X.CompressGZIP(edgelq_monitoring_proto_v3_phantom_time_serie_change_proto_rawDescData)
	})
	return edgelq_monitoring_proto_v3_phantom_time_serie_change_proto_rawDescData
}

var edgelq_monitoring_proto_v3_phantom_time_serie_change_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var edgelq_monitoring_proto_v3_phantom_time_serie_change_proto_goTypes = []interface{}{
	(*PhantomTimeSerieChange)(nil),          // 0: ntt.monitoring.v3.PhantomTimeSerieChange
	(*PhantomTimeSerieChange_Added)(nil),    // 1: ntt.monitoring.v3.PhantomTimeSerieChange.Added
	(*PhantomTimeSerieChange_Modified)(nil), // 2: ntt.monitoring.v3.PhantomTimeSerieChange.Modified
	(*PhantomTimeSerieChange_Current)(nil),  // 3: ntt.monitoring.v3.PhantomTimeSerieChange.Current
	(*PhantomTimeSerieChange_Removed)(nil),  // 4: ntt.monitoring.v3.PhantomTimeSerieChange.Removed
	(*PhantomTimeSerie)(nil),                // 5: ntt.monitoring.v3.PhantomTimeSerie
	(*PhantomTimeSerie_FieldMask)(nil),      // 6: ntt.monitoring.v3.PhantomTimeSerie_FieldMask
}
var edgelq_monitoring_proto_v3_phantom_time_serie_change_proto_depIdxs = []int32{
	1, // 0: ntt.monitoring.v3.PhantomTimeSerieChange.added:type_name -> ntt.monitoring.v3.PhantomTimeSerieChange.Added
	2, // 1: ntt.monitoring.v3.PhantomTimeSerieChange.modified:type_name -> ntt.monitoring.v3.PhantomTimeSerieChange.Modified
	3, // 2: ntt.monitoring.v3.PhantomTimeSerieChange.current:type_name -> ntt.monitoring.v3.PhantomTimeSerieChange.Current
	4, // 3: ntt.monitoring.v3.PhantomTimeSerieChange.removed:type_name -> ntt.monitoring.v3.PhantomTimeSerieChange.Removed
	5, // 4: ntt.monitoring.v3.PhantomTimeSerieChange.Added.phantom_time_serie:type_name -> ntt.monitoring.v3.PhantomTimeSerie
	5, // 5: ntt.monitoring.v3.PhantomTimeSerieChange.Modified.phantom_time_serie:type_name -> ntt.monitoring.v3.PhantomTimeSerie
	6, // 6: ntt.monitoring.v3.PhantomTimeSerieChange.Modified.field_mask:type_name -> ntt.monitoring.v3.PhantomTimeSerie_FieldMask
	5, // 7: ntt.monitoring.v3.PhantomTimeSerieChange.Current.phantom_time_serie:type_name -> ntt.monitoring.v3.PhantomTimeSerie
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { edgelq_monitoring_proto_v3_phantom_time_serie_change_proto_init() }
func edgelq_monitoring_proto_v3_phantom_time_serie_change_proto_init() {
	if edgelq_monitoring_proto_v3_phantom_time_serie_change_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		edgelq_monitoring_proto_v3_phantom_time_serie_change_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PhantomTimeSerieChange); i {
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
		edgelq_monitoring_proto_v3_phantom_time_serie_change_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PhantomTimeSerieChange_Added); i {
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
		edgelq_monitoring_proto_v3_phantom_time_serie_change_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PhantomTimeSerieChange_Modified); i {
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
		edgelq_monitoring_proto_v3_phantom_time_serie_change_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PhantomTimeSerieChange_Current); i {
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
		edgelq_monitoring_proto_v3_phantom_time_serie_change_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PhantomTimeSerieChange_Removed); i {
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

	edgelq_monitoring_proto_v3_phantom_time_serie_change_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*PhantomTimeSerieChange_Added_)(nil),
		(*PhantomTimeSerieChange_Modified_)(nil),
		(*PhantomTimeSerieChange_Current_)(nil),
		(*PhantomTimeSerieChange_Removed_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: edgelq_monitoring_proto_v3_phantom_time_serie_change_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           edgelq_monitoring_proto_v3_phantom_time_serie_change_proto_goTypes,
		DependencyIndexes: edgelq_monitoring_proto_v3_phantom_time_serie_change_proto_depIdxs,
		MessageInfos:      edgelq_monitoring_proto_v3_phantom_time_serie_change_proto_msgTypes,
	}.Build()
	edgelq_monitoring_proto_v3_phantom_time_serie_change_proto = out.File
	edgelq_monitoring_proto_v3_phantom_time_serie_change_proto_rawDesc = nil
	edgelq_monitoring_proto_v3_phantom_time_serie_change_proto_goTypes = nil
	edgelq_monitoring_proto_v3_phantom_time_serie_change_proto_depIdxs = nil
}
