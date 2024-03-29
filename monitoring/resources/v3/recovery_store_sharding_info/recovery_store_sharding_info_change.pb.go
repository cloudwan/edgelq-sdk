// Code generated by protoc-gen-goten-go
// File: edgelq/monitoring/proto/v3/recovery_store_sharding_info_change.proto
// DO NOT EDIT!!!

package recovery_store_sharding_info

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
	_ = &fieldmaskpb.FieldMask{}
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// RecoveryStoreShardingInfoChange is used by Watch notifications Responses to
// describe change of single RecoveryStoreShardingInfo One of Added, Modified,
// Removed
type RecoveryStoreShardingInfoChange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// RecoveryStoreShardingInfo change
	//
	// Types that are valid to be assigned to ChangeType:
	//	*RecoveryStoreShardingInfoChange_Added_
	//	*RecoveryStoreShardingInfoChange_Modified_
	//	*RecoveryStoreShardingInfoChange_Current_
	//	*RecoveryStoreShardingInfoChange_Removed_
	ChangeType isRecoveryStoreShardingInfoChange_ChangeType `protobuf_oneof:"change_type"`
}

func (m *RecoveryStoreShardingInfoChange) Reset() {
	*m = RecoveryStoreShardingInfoChange{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_monitoring_proto_v3_recovery_store_sharding_info_change_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *RecoveryStoreShardingInfoChange) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*RecoveryStoreShardingInfoChange) ProtoMessage() {}

func (m *RecoveryStoreShardingInfoChange) ProtoReflect() preflect.Message {
	mi := &edgelq_monitoring_proto_v3_recovery_store_sharding_info_change_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*RecoveryStoreShardingInfoChange) GotenMessage() {}

// Deprecated, Use RecoveryStoreShardingInfoChange.ProtoReflect.Descriptor instead.
func (*RecoveryStoreShardingInfoChange) Descriptor() ([]byte, []int) {
	return edgelq_monitoring_proto_v3_recovery_store_sharding_info_change_proto_rawDescGZIP(), []int{0}
}

func (m *RecoveryStoreShardingInfoChange) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *RecoveryStoreShardingInfoChange) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *RecoveryStoreShardingInfoChange) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *RecoveryStoreShardingInfoChange) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

type isRecoveryStoreShardingInfoChange_ChangeType interface {
	isRecoveryStoreShardingInfoChange_ChangeType()
}

type RecoveryStoreShardingInfoChange_Added_ struct {
	// Added is returned when watched document is added, either created or
	// enters Query view
	Added *RecoveryStoreShardingInfoChange_Added `protobuf:"bytes,1,opt,name=added,proto3,oneof" firestore:"added"`
}
type RecoveryStoreShardingInfoChange_Modified_ struct {
	// Modified is returned when watched document is modified
	Modified *RecoveryStoreShardingInfoChange_Modified `protobuf:"bytes,2,opt,name=modified,proto3,oneof" firestore:"modified"`
}
type RecoveryStoreShardingInfoChange_Current_ struct {
	// Current is returned in stateless watch when document enters query view or
	// is modified within.
	Current *RecoveryStoreShardingInfoChange_Current `protobuf:"bytes,4,opt,name=current,proto3,oneof" firestore:"current"`
}
type RecoveryStoreShardingInfoChange_Removed_ struct {
	// Removed is returned when RecoveryStoreShardingInfo is deleted or leaves
	// Query view
	Removed *RecoveryStoreShardingInfoChange_Removed `protobuf:"bytes,3,opt,name=removed,proto3,oneof" firestore:"removed"`
}

func (*RecoveryStoreShardingInfoChange_Added_) isRecoveryStoreShardingInfoChange_ChangeType()    {}
func (*RecoveryStoreShardingInfoChange_Modified_) isRecoveryStoreShardingInfoChange_ChangeType() {}
func (*RecoveryStoreShardingInfoChange_Current_) isRecoveryStoreShardingInfoChange_ChangeType()  {}
func (*RecoveryStoreShardingInfoChange_Removed_) isRecoveryStoreShardingInfoChange_ChangeType()  {}
func (m *RecoveryStoreShardingInfoChange) GetChangeType() isRecoveryStoreShardingInfoChange_ChangeType {
	if m != nil {
		return m.ChangeType
	}
	return nil
}
func (m *RecoveryStoreShardingInfoChange) GetAdded() *RecoveryStoreShardingInfoChange_Added {
	if x, ok := m.GetChangeType().(*RecoveryStoreShardingInfoChange_Added_); ok {
		return x.Added
	}
	return nil
}
func (m *RecoveryStoreShardingInfoChange) GetModified() *RecoveryStoreShardingInfoChange_Modified {
	if x, ok := m.GetChangeType().(*RecoveryStoreShardingInfoChange_Modified_); ok {
		return x.Modified
	}
	return nil
}
func (m *RecoveryStoreShardingInfoChange) GetCurrent() *RecoveryStoreShardingInfoChange_Current {
	if x, ok := m.GetChangeType().(*RecoveryStoreShardingInfoChange_Current_); ok {
		return x.Current
	}
	return nil
}
func (m *RecoveryStoreShardingInfoChange) GetRemoved() *RecoveryStoreShardingInfoChange_Removed {
	if x, ok := m.GetChangeType().(*RecoveryStoreShardingInfoChange_Removed_); ok {
		return x.Removed
	}
	return nil
}
func (m *RecoveryStoreShardingInfoChange) SetChangeType(ofv isRecoveryStoreShardingInfoChange_ChangeType) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "isRecoveryStoreShardingInfoChange_ChangeType", "RecoveryStoreShardingInfoChange"))
	}
	m.ChangeType = ofv
}
func (m *RecoveryStoreShardingInfoChange) SetAdded(fv *RecoveryStoreShardingInfoChange_Added) {
	m.SetChangeType(&RecoveryStoreShardingInfoChange_Added_{Added: fv})
}
func (m *RecoveryStoreShardingInfoChange) SetModified(fv *RecoveryStoreShardingInfoChange_Modified) {
	m.SetChangeType(&RecoveryStoreShardingInfoChange_Modified_{Modified: fv})
}
func (m *RecoveryStoreShardingInfoChange) SetCurrent(fv *RecoveryStoreShardingInfoChange_Current) {
	m.SetChangeType(&RecoveryStoreShardingInfoChange_Current_{Current: fv})
}
func (m *RecoveryStoreShardingInfoChange) SetRemoved(fv *RecoveryStoreShardingInfoChange_Removed) {
	m.SetChangeType(&RecoveryStoreShardingInfoChange_Removed_{Removed: fv})
}

// RecoveryStoreShardingInfo has been added to query view
type RecoveryStoreShardingInfoChange_Added struct {
	state                     protoimpl.MessageState
	sizeCache                 protoimpl.SizeCache
	unknownFields             protoimpl.UnknownFields
	RecoveryStoreShardingInfo *RecoveryStoreShardingInfo `protobuf:"bytes,1,opt,name=recovery_store_sharding_info,json=recoveryStoreShardingInfo,proto3" json:"recovery_store_sharding_info,omitempty" firestore:"recoveryStoreShardingInfo"`
	// Integer describing index of added RecoveryStoreShardingInfo in resulting
	// query view.
	ViewIndex int32 `protobuf:"varint,2,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty" firestore:"viewIndex"`
}

func (m *RecoveryStoreShardingInfoChange_Added) Reset() {
	*m = RecoveryStoreShardingInfoChange_Added{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_monitoring_proto_v3_recovery_store_sharding_info_change_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *RecoveryStoreShardingInfoChange_Added) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*RecoveryStoreShardingInfoChange_Added) ProtoMessage() {}

func (m *RecoveryStoreShardingInfoChange_Added) ProtoReflect() preflect.Message {
	mi := &edgelq_monitoring_proto_v3_recovery_store_sharding_info_change_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*RecoveryStoreShardingInfoChange_Added) GotenMessage() {}

// Deprecated, Use RecoveryStoreShardingInfoChange_Added.ProtoReflect.Descriptor instead.
func (*RecoveryStoreShardingInfoChange_Added) Descriptor() ([]byte, []int) {
	return edgelq_monitoring_proto_v3_recovery_store_sharding_info_change_proto_rawDescGZIP(), []int{0, 0}
}

func (m *RecoveryStoreShardingInfoChange_Added) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *RecoveryStoreShardingInfoChange_Added) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *RecoveryStoreShardingInfoChange_Added) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *RecoveryStoreShardingInfoChange_Added) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *RecoveryStoreShardingInfoChange_Added) GetRecoveryStoreShardingInfo() *RecoveryStoreShardingInfo {
	if m != nil {
		return m.RecoveryStoreShardingInfo
	}
	return nil
}

func (m *RecoveryStoreShardingInfoChange_Added) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *RecoveryStoreShardingInfoChange_Added) SetRecoveryStoreShardingInfo(fv *RecoveryStoreShardingInfo) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "RecoveryStoreShardingInfo", "RecoveryStoreShardingInfoChange_Added"))
	}
	m.RecoveryStoreShardingInfo = fv
}

func (m *RecoveryStoreShardingInfoChange_Added) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "RecoveryStoreShardingInfoChange_Added"))
	}
	m.ViewIndex = fv
}

// RecoveryStoreShardingInfo changed some of it's fields - contains either
// full document or masked change
type RecoveryStoreShardingInfoChange_Modified struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Name of modified RecoveryStoreShardingInfo
	Name *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// New version of RecoveryStoreShardingInfo or masked difference, depending
	// on mask_changes instrumentation of issued
	// [WatchRecoveryStoreShardingInfoRequest] or
	// [WatchRecoveryStoreShardingInfosRequest]
	RecoveryStoreShardingInfo *RecoveryStoreShardingInfo `protobuf:"bytes,2,opt,name=recovery_store_sharding_info,json=recoveryStoreShardingInfo,proto3" json:"recovery_store_sharding_info,omitempty" firestore:"recoveryStoreShardingInfo"`
	// Used when mask_changes is set, contains field paths of modified
	// properties.
	FieldMask *RecoveryStoreShardingInfo_FieldMask `protobuf:"bytes,3,opt,customtype=RecoveryStoreShardingInfo_FieldMask,name=field_mask,json=fieldMask,proto3" json:"field_mask,omitempty" firestore:"fieldMask"`
	// Previous view index specifies previous position of modified
	// RecoveryStoreShardingInfo. When modification doesn't affect sorted order,
	// value will remain identical to [view_index].
	PreviousViewIndex int32 `protobuf:"varint,4,opt,name=previous_view_index,json=previousViewIndex,proto3" json:"previous_view_index,omitempty" firestore:"previousViewIndex"`
	// Integer specifying RecoveryStoreShardingInfo new index in resulting query
	// view.
	ViewIndex int32 `protobuf:"varint,5,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty" firestore:"viewIndex"`
}

func (m *RecoveryStoreShardingInfoChange_Modified) Reset() {
	*m = RecoveryStoreShardingInfoChange_Modified{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_monitoring_proto_v3_recovery_store_sharding_info_change_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *RecoveryStoreShardingInfoChange_Modified) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*RecoveryStoreShardingInfoChange_Modified) ProtoMessage() {}

func (m *RecoveryStoreShardingInfoChange_Modified) ProtoReflect() preflect.Message {
	mi := &edgelq_monitoring_proto_v3_recovery_store_sharding_info_change_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*RecoveryStoreShardingInfoChange_Modified) GotenMessage() {}

// Deprecated, Use RecoveryStoreShardingInfoChange_Modified.ProtoReflect.Descriptor instead.
func (*RecoveryStoreShardingInfoChange_Modified) Descriptor() ([]byte, []int) {
	return edgelq_monitoring_proto_v3_recovery_store_sharding_info_change_proto_rawDescGZIP(), []int{0, 1}
}

func (m *RecoveryStoreShardingInfoChange_Modified) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *RecoveryStoreShardingInfoChange_Modified) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *RecoveryStoreShardingInfoChange_Modified) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *RecoveryStoreShardingInfoChange_Modified) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *RecoveryStoreShardingInfoChange_Modified) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *RecoveryStoreShardingInfoChange_Modified) GetRecoveryStoreShardingInfo() *RecoveryStoreShardingInfo {
	if m != nil {
		return m.RecoveryStoreShardingInfo
	}
	return nil
}

func (m *RecoveryStoreShardingInfoChange_Modified) GetFieldMask() *RecoveryStoreShardingInfo_FieldMask {
	if m != nil {
		return m.FieldMask
	}
	return nil
}

func (m *RecoveryStoreShardingInfoChange_Modified) GetPreviousViewIndex() int32 {
	if m != nil {
		return m.PreviousViewIndex
	}
	return int32(0)
}

func (m *RecoveryStoreShardingInfoChange_Modified) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *RecoveryStoreShardingInfoChange_Modified) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "RecoveryStoreShardingInfoChange_Modified"))
	}
	m.Name = fv
}

func (m *RecoveryStoreShardingInfoChange_Modified) SetRecoveryStoreShardingInfo(fv *RecoveryStoreShardingInfo) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "RecoveryStoreShardingInfo", "RecoveryStoreShardingInfoChange_Modified"))
	}
	m.RecoveryStoreShardingInfo = fv
}

func (m *RecoveryStoreShardingInfoChange_Modified) SetFieldMask(fv *RecoveryStoreShardingInfo_FieldMask) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "FieldMask", "RecoveryStoreShardingInfoChange_Modified"))
	}
	m.FieldMask = fv
}

func (m *RecoveryStoreShardingInfoChange_Modified) SetPreviousViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "PreviousViewIndex", "RecoveryStoreShardingInfoChange_Modified"))
	}
	m.PreviousViewIndex = fv
}

func (m *RecoveryStoreShardingInfoChange_Modified) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "RecoveryStoreShardingInfoChange_Modified"))
	}
	m.ViewIndex = fv
}

// RecoveryStoreShardingInfo has been added or modified in a query view.
// Version used for stateless watching
type RecoveryStoreShardingInfoChange_Current struct {
	state                     protoimpl.MessageState
	sizeCache                 protoimpl.SizeCache
	unknownFields             protoimpl.UnknownFields
	RecoveryStoreShardingInfo *RecoveryStoreShardingInfo `protobuf:"bytes,1,opt,name=recovery_store_sharding_info,json=recoveryStoreShardingInfo,proto3" json:"recovery_store_sharding_info,omitempty" firestore:"recoveryStoreShardingInfo"`
}

func (m *RecoveryStoreShardingInfoChange_Current) Reset() {
	*m = RecoveryStoreShardingInfoChange_Current{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_monitoring_proto_v3_recovery_store_sharding_info_change_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *RecoveryStoreShardingInfoChange_Current) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*RecoveryStoreShardingInfoChange_Current) ProtoMessage() {}

func (m *RecoveryStoreShardingInfoChange_Current) ProtoReflect() preflect.Message {
	mi := &edgelq_monitoring_proto_v3_recovery_store_sharding_info_change_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*RecoveryStoreShardingInfoChange_Current) GotenMessage() {}

// Deprecated, Use RecoveryStoreShardingInfoChange_Current.ProtoReflect.Descriptor instead.
func (*RecoveryStoreShardingInfoChange_Current) Descriptor() ([]byte, []int) {
	return edgelq_monitoring_proto_v3_recovery_store_sharding_info_change_proto_rawDescGZIP(), []int{0, 2}
}

func (m *RecoveryStoreShardingInfoChange_Current) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *RecoveryStoreShardingInfoChange_Current) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *RecoveryStoreShardingInfoChange_Current) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *RecoveryStoreShardingInfoChange_Current) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *RecoveryStoreShardingInfoChange_Current) GetRecoveryStoreShardingInfo() *RecoveryStoreShardingInfo {
	if m != nil {
		return m.RecoveryStoreShardingInfo
	}
	return nil
}

func (m *RecoveryStoreShardingInfoChange_Current) SetRecoveryStoreShardingInfo(fv *RecoveryStoreShardingInfo) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "RecoveryStoreShardingInfo", "RecoveryStoreShardingInfoChange_Current"))
	}
	m.RecoveryStoreShardingInfo = fv
}

// Removed is returned when RecoveryStoreShardingInfo is deleted or leaves
// Query view
type RecoveryStoreShardingInfoChange_Removed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Name          *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// Integer specifying removed RecoveryStoreShardingInfo index. Not populated
	// in stateless watch type.
	ViewIndex int32 `protobuf:"varint,2,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty" firestore:"viewIndex"`
}

func (m *RecoveryStoreShardingInfoChange_Removed) Reset() {
	*m = RecoveryStoreShardingInfoChange_Removed{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_monitoring_proto_v3_recovery_store_sharding_info_change_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *RecoveryStoreShardingInfoChange_Removed) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*RecoveryStoreShardingInfoChange_Removed) ProtoMessage() {}

func (m *RecoveryStoreShardingInfoChange_Removed) ProtoReflect() preflect.Message {
	mi := &edgelq_monitoring_proto_v3_recovery_store_sharding_info_change_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*RecoveryStoreShardingInfoChange_Removed) GotenMessage() {}

// Deprecated, Use RecoveryStoreShardingInfoChange_Removed.ProtoReflect.Descriptor instead.
func (*RecoveryStoreShardingInfoChange_Removed) Descriptor() ([]byte, []int) {
	return edgelq_monitoring_proto_v3_recovery_store_sharding_info_change_proto_rawDescGZIP(), []int{0, 3}
}

func (m *RecoveryStoreShardingInfoChange_Removed) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *RecoveryStoreShardingInfoChange_Removed) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *RecoveryStoreShardingInfoChange_Removed) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *RecoveryStoreShardingInfoChange_Removed) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *RecoveryStoreShardingInfoChange_Removed) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *RecoveryStoreShardingInfoChange_Removed) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *RecoveryStoreShardingInfoChange_Removed) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "RecoveryStoreShardingInfoChange_Removed"))
	}
	m.Name = fv
}

func (m *RecoveryStoreShardingInfoChange_Removed) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "RecoveryStoreShardingInfoChange_Removed"))
	}
	m.ViewIndex = fv
}

var edgelq_monitoring_proto_v3_recovery_store_sharding_info_change_proto preflect.FileDescriptor

var edgelq_monitoring_proto_v3_recovery_store_sharding_info_change_proto_rawDesc = []byte{
	0x0a, 0x44, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x69, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x33, 0x2f, 0x72, 0x65, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x73, 0x68, 0x61, 0x72,
	0x64, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x6e, 0x74, 0x74, 0x2e, 0x6d, 0x6f, 0x6e, 0x69,
	0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3d, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f,
	0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x76, 0x33, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x81, 0x09, 0x0a, 0x1f, 0x52, 0x65, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x79, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x53, 0x68, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67,
	0x49, 0x6e, 0x66, 0x6f, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x50, 0x0a, 0x05, 0x61, 0x64,
	0x64, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x6e, 0x74, 0x74, 0x2e,
	0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x52, 0x65,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x53, 0x68, 0x61, 0x72, 0x64,
	0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x41, 0x64,
	0x64, 0x65, 0x64, 0x48, 0x00, 0x52, 0x05, 0x61, 0x64, 0x64, 0x65, 0x64, 0x12, 0x59, 0x0a, 0x08,
	0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3b,
	0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x33, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x53, 0x68, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x48, 0x00, 0x52, 0x08, 0x6d,
	0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x56, 0x0a, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x6d,
	0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x52, 0x65, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x79, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x53, 0x68, 0x61, 0x72, 0x64, 0x69,
	0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12,
	0x56, 0x0a, 0x07, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x3a, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x33, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x53, 0x68, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x48, 0x00, 0x52, 0x07,
	0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x1a, 0x95, 0x01, 0x0a, 0x05, 0x41, 0x64, 0x64, 0x65,
	0x64, 0x12, 0x6d, 0x0a, 0x1c, 0x72, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x6e, 0x66,
	0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x6d, 0x6f,
	0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x52, 0x65, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x79, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x53, 0x68, 0x61, 0x72, 0x64, 0x69, 0x6e,
	0x67, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x19, 0x72, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x53, 0x68, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x76, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x1a,
	0xdd, 0x02, 0x0a, 0x08, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x35, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x21, 0xb2, 0xda, 0x21, 0x1d,
	0x0a, 0x1b, 0x0a, 0x19, 0x52, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x53, 0x74, 0x6f, 0x72,
	0x65, 0x53, 0x68, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x6d, 0x0a, 0x1c, 0x72, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x5f,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x6e, 0x74, 0x74, 0x2e,
	0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x52, 0x65,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x53, 0x68, 0x61, 0x72, 0x64,
	0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x19, 0x72, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x79, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x53, 0x68, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x5c, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61,
	0x73, 0x6b, 0x42, 0x21, 0xb2, 0xda, 0x21, 0x1d, 0x32, 0x1b, 0x0a, 0x19, 0x52, 0x65, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x79, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x53, 0x68, 0x61, 0x72, 0x64, 0x69, 0x6e,
	0x67, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b,
	0x12, 0x2e, 0x0a, 0x13, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x5f, 0x76, 0x69, 0x65,
	0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x70,
	0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x56, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x76, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x1a,
	0x78, 0x0a, 0x07, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x6d, 0x0a, 0x1c, 0x72, 0x65,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x73, 0x68, 0x61,
	0x72, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2c, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x33, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x53, 0x68, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x19,
	0x72, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x53, 0x68, 0x61,
	0x72, 0x64, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x5f, 0x0a, 0x07, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x64, 0x12, 0x35, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x21, 0xb2, 0xda, 0x21, 0x1d, 0x0a, 0x1b, 0x0a, 0x19, 0x52, 0x65, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x79, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x53, 0x68, 0x61, 0x72, 0x64, 0x69, 0x6e,
	0x67, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x76,
	0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x76, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x3a, 0x1f, 0x9a, 0xd9, 0x21, 0x1b,
	0x0a, 0x19, 0x52, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x53,
	0x68, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x0d, 0x0a, 0x0b, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x42, 0xb4, 0x01, 0xe8, 0xde, 0x21,
	0x00, 0x0a, 0x18, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74,
	0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x33, 0x42, 0x24, 0x52, 0x65, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x79, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x53, 0x68, 0x61, 0x72, 0x64, 0x69,
	0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x00, 0x5a, 0x6c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f,
	0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x3b, 0x72, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x6e, 0x66,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	edgelq_monitoring_proto_v3_recovery_store_sharding_info_change_proto_rawDescOnce sync.Once
	edgelq_monitoring_proto_v3_recovery_store_sharding_info_change_proto_rawDescData = edgelq_monitoring_proto_v3_recovery_store_sharding_info_change_proto_rawDesc
)

func edgelq_monitoring_proto_v3_recovery_store_sharding_info_change_proto_rawDescGZIP() []byte {
	edgelq_monitoring_proto_v3_recovery_store_sharding_info_change_proto_rawDescOnce.Do(func() {
		edgelq_monitoring_proto_v3_recovery_store_sharding_info_change_proto_rawDescData = protoimpl.X.CompressGZIP(edgelq_monitoring_proto_v3_recovery_store_sharding_info_change_proto_rawDescData)
	})
	return edgelq_monitoring_proto_v3_recovery_store_sharding_info_change_proto_rawDescData
}

var edgelq_monitoring_proto_v3_recovery_store_sharding_info_change_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var edgelq_monitoring_proto_v3_recovery_store_sharding_info_change_proto_goTypes = []interface{}{
	(*RecoveryStoreShardingInfoChange)(nil),          // 0: ntt.monitoring.v3.RecoveryStoreShardingInfoChange
	(*RecoveryStoreShardingInfoChange_Added)(nil),    // 1: ntt.monitoring.v3.RecoveryStoreShardingInfoChange.Added
	(*RecoveryStoreShardingInfoChange_Modified)(nil), // 2: ntt.monitoring.v3.RecoveryStoreShardingInfoChange.Modified
	(*RecoveryStoreShardingInfoChange_Current)(nil),  // 3: ntt.monitoring.v3.RecoveryStoreShardingInfoChange.Current
	(*RecoveryStoreShardingInfoChange_Removed)(nil),  // 4: ntt.monitoring.v3.RecoveryStoreShardingInfoChange.Removed
	(*RecoveryStoreShardingInfo)(nil),                // 5: ntt.monitoring.v3.RecoveryStoreShardingInfo
	(*RecoveryStoreShardingInfo_FieldMask)(nil),      // 6: ntt.monitoring.v3.RecoveryStoreShardingInfo_FieldMask
}
var edgelq_monitoring_proto_v3_recovery_store_sharding_info_change_proto_depIdxs = []int32{
	1, // 0: ntt.monitoring.v3.RecoveryStoreShardingInfoChange.added:type_name -> ntt.monitoring.v3.RecoveryStoreShardingInfoChange.Added
	2, // 1: ntt.monitoring.v3.RecoveryStoreShardingInfoChange.modified:type_name -> ntt.monitoring.v3.RecoveryStoreShardingInfoChange.Modified
	3, // 2: ntt.monitoring.v3.RecoveryStoreShardingInfoChange.current:type_name -> ntt.monitoring.v3.RecoveryStoreShardingInfoChange.Current
	4, // 3: ntt.monitoring.v3.RecoveryStoreShardingInfoChange.removed:type_name -> ntt.monitoring.v3.RecoveryStoreShardingInfoChange.Removed
	5, // 4: ntt.monitoring.v3.RecoveryStoreShardingInfoChange.Added.recovery_store_sharding_info:type_name -> ntt.monitoring.v3.RecoveryStoreShardingInfo
	5, // 5: ntt.monitoring.v3.RecoveryStoreShardingInfoChange.Modified.recovery_store_sharding_info:type_name -> ntt.monitoring.v3.RecoveryStoreShardingInfo
	6, // 6: ntt.monitoring.v3.RecoveryStoreShardingInfoChange.Modified.field_mask:type_name -> ntt.monitoring.v3.RecoveryStoreShardingInfo_FieldMask
	5, // 7: ntt.monitoring.v3.RecoveryStoreShardingInfoChange.Current.recovery_store_sharding_info:type_name -> ntt.monitoring.v3.RecoveryStoreShardingInfo
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { edgelq_monitoring_proto_v3_recovery_store_sharding_info_change_proto_init() }
func edgelq_monitoring_proto_v3_recovery_store_sharding_info_change_proto_init() {
	if edgelq_monitoring_proto_v3_recovery_store_sharding_info_change_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		edgelq_monitoring_proto_v3_recovery_store_sharding_info_change_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecoveryStoreShardingInfoChange); i {
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
		edgelq_monitoring_proto_v3_recovery_store_sharding_info_change_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecoveryStoreShardingInfoChange_Added); i {
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
		edgelq_monitoring_proto_v3_recovery_store_sharding_info_change_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecoveryStoreShardingInfoChange_Modified); i {
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
		edgelq_monitoring_proto_v3_recovery_store_sharding_info_change_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecoveryStoreShardingInfoChange_Current); i {
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
		edgelq_monitoring_proto_v3_recovery_store_sharding_info_change_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecoveryStoreShardingInfoChange_Removed); i {
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

	edgelq_monitoring_proto_v3_recovery_store_sharding_info_change_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*RecoveryStoreShardingInfoChange_Added_)(nil),
		(*RecoveryStoreShardingInfoChange_Modified_)(nil),
		(*RecoveryStoreShardingInfoChange_Current_)(nil),
		(*RecoveryStoreShardingInfoChange_Removed_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: edgelq_monitoring_proto_v3_recovery_store_sharding_info_change_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           edgelq_monitoring_proto_v3_recovery_store_sharding_info_change_proto_goTypes,
		DependencyIndexes: edgelq_monitoring_proto_v3_recovery_store_sharding_info_change_proto_depIdxs,
		MessageInfos:      edgelq_monitoring_proto_v3_recovery_store_sharding_info_change_proto_msgTypes,
	}.Build()
	edgelq_monitoring_proto_v3_recovery_store_sharding_info_change_proto = out.File
	edgelq_monitoring_proto_v3_recovery_store_sharding_info_change_proto_rawDesc = nil
	edgelq_monitoring_proto_v3_recovery_store_sharding_info_change_proto_goTypes = nil
	edgelq_monitoring_proto_v3_recovery_store_sharding_info_change_proto_depIdxs = nil
}
