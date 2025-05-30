// Code generated by protoc-gen-goten-go
// File: edgelq/limits/proto/v1alpha2/limit_change.proto
// DO NOT EDIT!!!

package limit

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
	iam_project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/project"
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
	_ = &iam_project.Project{}
	_ = &fieldmaskpb.FieldMask{}
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// LimitChange is used by Watch notifications Responses to describe change of
// single Limit One of Added, Modified, Removed
type LimitChange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Limit change
	//
	// Types that are valid to be assigned to ChangeType:
	//	*LimitChange_Added_
	//	*LimitChange_Modified_
	//	*LimitChange_Current_
	//	*LimitChange_Removed_
	ChangeType isLimitChange_ChangeType `protobuf_oneof:"change_type"`
}

func (m *LimitChange) Reset() {
	*m = LimitChange{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_limits_proto_v1alpha2_limit_change_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *LimitChange) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*LimitChange) ProtoMessage() {}

func (m *LimitChange) ProtoReflect() preflect.Message {
	mi := &edgelq_limits_proto_v1alpha2_limit_change_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*LimitChange) GotenMessage() {}

// Deprecated, Use LimitChange.ProtoReflect.Descriptor instead.
func (*LimitChange) Descriptor() ([]byte, []int) {
	return edgelq_limits_proto_v1alpha2_limit_change_proto_rawDescGZIP(), []int{0}
}

func (m *LimitChange) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *LimitChange) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *LimitChange) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *LimitChange) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

type isLimitChange_ChangeType interface {
	isLimitChange_ChangeType()
}

type LimitChange_Added_ struct {
	// Added is returned when watched document is added, either created or
	// enters Query view
	Added *LimitChange_Added `protobuf:"bytes,1,opt,name=added,proto3,oneof"`
}
type LimitChange_Modified_ struct {
	// Modified is returned when watched document is modified
	Modified *LimitChange_Modified `protobuf:"bytes,2,opt,name=modified,proto3,oneof"`
}
type LimitChange_Current_ struct {
	// Current is returned in stateless watch when document enters query view or
	// is modified within.
	Current *LimitChange_Current `protobuf:"bytes,4,opt,name=current,proto3,oneof"`
}
type LimitChange_Removed_ struct {
	// Removed is returned when Limit is deleted or leaves Query view
	Removed *LimitChange_Removed `protobuf:"bytes,3,opt,name=removed,proto3,oneof"`
}

func (*LimitChange_Added_) isLimitChange_ChangeType()    {}
func (*LimitChange_Modified_) isLimitChange_ChangeType() {}
func (*LimitChange_Current_) isLimitChange_ChangeType()  {}
func (*LimitChange_Removed_) isLimitChange_ChangeType()  {}
func (m *LimitChange) GetChangeType() isLimitChange_ChangeType {
	if m != nil {
		return m.ChangeType
	}
	return nil
}
func (m *LimitChange) GetAdded() *LimitChange_Added {
	if x, ok := m.GetChangeType().(*LimitChange_Added_); ok {
		return x.Added
	}
	return nil
}
func (m *LimitChange) GetModified() *LimitChange_Modified {
	if x, ok := m.GetChangeType().(*LimitChange_Modified_); ok {
		return x.Modified
	}
	return nil
}
func (m *LimitChange) GetCurrent() *LimitChange_Current {
	if x, ok := m.GetChangeType().(*LimitChange_Current_); ok {
		return x.Current
	}
	return nil
}
func (m *LimitChange) GetRemoved() *LimitChange_Removed {
	if x, ok := m.GetChangeType().(*LimitChange_Removed_); ok {
		return x.Removed
	}
	return nil
}
func (m *LimitChange) SetChangeType(ofv isLimitChange_ChangeType) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "isLimitChange_ChangeType", "LimitChange"))
	}
	m.ChangeType = ofv
}
func (m *LimitChange) SetAdded(fv *LimitChange_Added) {
	m.SetChangeType(&LimitChange_Added_{Added: fv})
}
func (m *LimitChange) SetModified(fv *LimitChange_Modified) {
	m.SetChangeType(&LimitChange_Modified_{Modified: fv})
}
func (m *LimitChange) SetCurrent(fv *LimitChange_Current) {
	m.SetChangeType(&LimitChange_Current_{Current: fv})
}
func (m *LimitChange) SetRemoved(fv *LimitChange_Removed) {
	m.SetChangeType(&LimitChange_Removed_{Removed: fv})
}

// Limit has been added to query view
type LimitChange_Added struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Limit         *Limit `protobuf:"bytes,1,opt,name=limit,proto3" json:"limit,omitempty"`
	// Integer describing index of added Limit in resulting query view.
	ViewIndex int32 `protobuf:"varint,2,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty"`
}

func (m *LimitChange_Added) Reset() {
	*m = LimitChange_Added{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_limits_proto_v1alpha2_limit_change_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *LimitChange_Added) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*LimitChange_Added) ProtoMessage() {}

func (m *LimitChange_Added) ProtoReflect() preflect.Message {
	mi := &edgelq_limits_proto_v1alpha2_limit_change_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*LimitChange_Added) GotenMessage() {}

// Deprecated, Use LimitChange_Added.ProtoReflect.Descriptor instead.
func (*LimitChange_Added) Descriptor() ([]byte, []int) {
	return edgelq_limits_proto_v1alpha2_limit_change_proto_rawDescGZIP(), []int{0, 0}
}

func (m *LimitChange_Added) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *LimitChange_Added) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *LimitChange_Added) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *LimitChange_Added) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *LimitChange_Added) GetLimit() *Limit {
	if m != nil {
		return m.Limit
	}
	return nil
}

func (m *LimitChange_Added) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *LimitChange_Added) SetLimit(fv *Limit) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Limit", "LimitChange_Added"))
	}
	m.Limit = fv
}

func (m *LimitChange_Added) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "LimitChange_Added"))
	}
	m.ViewIndex = fv
}

// Limit changed some of it's fields - contains either full document or masked
// change
type LimitChange_Modified struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Name of modified Limit
	Name *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty"`
	// New version of Limit or masked difference, depending on mask_changes
	// instrumentation of issued [WatchLimitRequest] or [WatchLimitsRequest]
	Limit *Limit `protobuf:"bytes,2,opt,name=limit,proto3" json:"limit,omitempty"`
	// Used when mask_changes is set, contains field paths of modified
	// properties.
	FieldMask *Limit_FieldMask `protobuf:"bytes,3,opt,customtype=Limit_FieldMask,name=field_mask,json=fieldMask,proto3" json:"field_mask,omitempty"`
	// Previous view index specifies previous position of modified Limit.
	// When modification doesn't affect sorted order, value will remain
	// identical to [view_index].
	PreviousViewIndex int32 `protobuf:"varint,4,opt,name=previous_view_index,json=previousViewIndex,proto3" json:"previous_view_index,omitempty"`
	// Integer specifying Limit new index in resulting query view.
	ViewIndex int32 `protobuf:"varint,5,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty"`
}

func (m *LimitChange_Modified) Reset() {
	*m = LimitChange_Modified{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_limits_proto_v1alpha2_limit_change_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *LimitChange_Modified) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*LimitChange_Modified) ProtoMessage() {}

func (m *LimitChange_Modified) ProtoReflect() preflect.Message {
	mi := &edgelq_limits_proto_v1alpha2_limit_change_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*LimitChange_Modified) GotenMessage() {}

// Deprecated, Use LimitChange_Modified.ProtoReflect.Descriptor instead.
func (*LimitChange_Modified) Descriptor() ([]byte, []int) {
	return edgelq_limits_proto_v1alpha2_limit_change_proto_rawDescGZIP(), []int{0, 1}
}

func (m *LimitChange_Modified) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *LimitChange_Modified) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *LimitChange_Modified) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *LimitChange_Modified) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *LimitChange_Modified) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *LimitChange_Modified) GetLimit() *Limit {
	if m != nil {
		return m.Limit
	}
	return nil
}

func (m *LimitChange_Modified) GetFieldMask() *Limit_FieldMask {
	if m != nil {
		return m.FieldMask
	}
	return nil
}

func (m *LimitChange_Modified) GetPreviousViewIndex() int32 {
	if m != nil {
		return m.PreviousViewIndex
	}
	return int32(0)
}

func (m *LimitChange_Modified) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *LimitChange_Modified) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "LimitChange_Modified"))
	}
	m.Name = fv
}

func (m *LimitChange_Modified) SetLimit(fv *Limit) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Limit", "LimitChange_Modified"))
	}
	m.Limit = fv
}

func (m *LimitChange_Modified) SetFieldMask(fv *Limit_FieldMask) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "FieldMask", "LimitChange_Modified"))
	}
	m.FieldMask = fv
}

func (m *LimitChange_Modified) SetPreviousViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "PreviousViewIndex", "LimitChange_Modified"))
	}
	m.PreviousViewIndex = fv
}

func (m *LimitChange_Modified) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "LimitChange_Modified"))
	}
	m.ViewIndex = fv
}

// Limit has been added or modified in a query view. Version used for
// stateless watching
type LimitChange_Current struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Limit         *Limit `protobuf:"bytes,1,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (m *LimitChange_Current) Reset() {
	*m = LimitChange_Current{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_limits_proto_v1alpha2_limit_change_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *LimitChange_Current) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*LimitChange_Current) ProtoMessage() {}

func (m *LimitChange_Current) ProtoReflect() preflect.Message {
	mi := &edgelq_limits_proto_v1alpha2_limit_change_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*LimitChange_Current) GotenMessage() {}

// Deprecated, Use LimitChange_Current.ProtoReflect.Descriptor instead.
func (*LimitChange_Current) Descriptor() ([]byte, []int) {
	return edgelq_limits_proto_v1alpha2_limit_change_proto_rawDescGZIP(), []int{0, 2}
}

func (m *LimitChange_Current) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *LimitChange_Current) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *LimitChange_Current) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *LimitChange_Current) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *LimitChange_Current) GetLimit() *Limit {
	if m != nil {
		return m.Limit
	}
	return nil
}

func (m *LimitChange_Current) SetLimit(fv *Limit) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Limit", "LimitChange_Current"))
	}
	m.Limit = fv
}

// Removed is returned when Limit is deleted or leaves Query view
type LimitChange_Removed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Name          *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty"`
	// Integer specifying removed Limit index. Not populated in stateless watch
	// type.
	ViewIndex int32 `protobuf:"varint,2,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty"`
}

func (m *LimitChange_Removed) Reset() {
	*m = LimitChange_Removed{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_limits_proto_v1alpha2_limit_change_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *LimitChange_Removed) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*LimitChange_Removed) ProtoMessage() {}

func (m *LimitChange_Removed) ProtoReflect() preflect.Message {
	mi := &edgelq_limits_proto_v1alpha2_limit_change_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*LimitChange_Removed) GotenMessage() {}

// Deprecated, Use LimitChange_Removed.ProtoReflect.Descriptor instead.
func (*LimitChange_Removed) Descriptor() ([]byte, []int) {
	return edgelq_limits_proto_v1alpha2_limit_change_proto_rawDescGZIP(), []int{0, 3}
}

func (m *LimitChange_Removed) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *LimitChange_Removed) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *LimitChange_Removed) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *LimitChange_Removed) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *LimitChange_Removed) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *LimitChange_Removed) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *LimitChange_Removed) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "LimitChange_Removed"))
	}
	m.Name = fv
}

func (m *LimitChange_Removed) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "LimitChange_Removed"))
	}
	m.ViewIndex = fv
}

var edgelq_limits_proto_v1alpha2_limit_change_proto preflect.FileDescriptor

var edgelq_limits_proto_v1alpha2_limit_change_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x13, 0x6e, 0x74, 0x74, 0x2e, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x32, 0x2f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9d,
	0x06, 0x0a, 0x0b, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x3e,
	0x0a, 0x05, 0x61, 0x64, 0x64, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e,
	0x6e, 0x74, 0x74, 0x2e, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x32, 0x2e, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e,
	0x41, 0x64, 0x64, 0x65, 0x64, 0x48, 0x00, 0x52, 0x05, 0x61, 0x64, 0x64, 0x65, 0x64, 0x12, 0x47,
	0x0a, 0x08, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x29, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x48, 0x00, 0x52, 0x08, 0x6d,
	0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x44, 0x0a, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x48, 0x00, 0x52, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x44, 0x0a,
	0x07, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28,
	0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x32, 0x2e, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x48, 0x00, 0x52, 0x07, 0x72, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x64, 0x1a, 0x58, 0x0a, 0x05, 0x41, 0x64, 0x64, 0x65, 0x64, 0x12, 0x30, 0x0a, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6e, 0x74,
	0x74, 0x2e, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x32, 0x2e, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x76, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x1a, 0xf8, 0x01,
	0x0a, 0x08, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x21, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0d, 0xb2, 0xda, 0x21, 0x09, 0x0a, 0x07,
	0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x30, 0x0a,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6e,
	0x74, 0x74, 0x2e, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x32, 0x2e, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12,
	0x48, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x42,
	0x0d, 0xb2, 0xda, 0x21, 0x09, 0x32, 0x07, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x09,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x2e, 0x0a, 0x13, 0x70, 0x72, 0x65,
	0x76, 0x69, 0x6f, 0x75, 0x73, 0x5f, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73,
	0x56, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x69, 0x65,
	0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x76,
	0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x1a, 0x3b, 0x0a, 0x07, 0x43, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x12, 0x30, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x1a, 0x4b, 0x0a, 0x07, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64,
	0x12, 0x21, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0d,
	0xb2, 0xda, 0x21, 0x09, 0x0a, 0x07, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x76, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x3a, 0x0b, 0x9a, 0xd9, 0x21, 0x07, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x42,
	0x0d, 0x0a, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x42, 0x76,
	0xe8, 0xde, 0x21, 0x00, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x73, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32,
	0x42, 0x10, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x00, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71,
	0x2f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x3b, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	edgelq_limits_proto_v1alpha2_limit_change_proto_rawDescOnce sync.Once
	edgelq_limits_proto_v1alpha2_limit_change_proto_rawDescData = edgelq_limits_proto_v1alpha2_limit_change_proto_rawDesc
)

func edgelq_limits_proto_v1alpha2_limit_change_proto_rawDescGZIP() []byte {
	edgelq_limits_proto_v1alpha2_limit_change_proto_rawDescOnce.Do(func() {
		edgelq_limits_proto_v1alpha2_limit_change_proto_rawDescData = protoimpl.X.CompressGZIP(edgelq_limits_proto_v1alpha2_limit_change_proto_rawDescData)
	})
	return edgelq_limits_proto_v1alpha2_limit_change_proto_rawDescData
}

var edgelq_limits_proto_v1alpha2_limit_change_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var edgelq_limits_proto_v1alpha2_limit_change_proto_goTypes = []interface{}{
	(*LimitChange)(nil),          // 0: ntt.limits.v1alpha2.LimitChange
	(*LimitChange_Added)(nil),    // 1: ntt.limits.v1alpha2.LimitChange.Added
	(*LimitChange_Modified)(nil), // 2: ntt.limits.v1alpha2.LimitChange.Modified
	(*LimitChange_Current)(nil),  // 3: ntt.limits.v1alpha2.LimitChange.Current
	(*LimitChange_Removed)(nil),  // 4: ntt.limits.v1alpha2.LimitChange.Removed
	(*Limit)(nil),                // 5: ntt.limits.v1alpha2.Limit
	(*Limit_FieldMask)(nil),      // 6: ntt.limits.v1alpha2.Limit_FieldMask
}
var edgelq_limits_proto_v1alpha2_limit_change_proto_depIdxs = []int32{
	1, // 0: ntt.limits.v1alpha2.LimitChange.added:type_name -> ntt.limits.v1alpha2.LimitChange.Added
	2, // 1: ntt.limits.v1alpha2.LimitChange.modified:type_name -> ntt.limits.v1alpha2.LimitChange.Modified
	3, // 2: ntt.limits.v1alpha2.LimitChange.current:type_name -> ntt.limits.v1alpha2.LimitChange.Current
	4, // 3: ntt.limits.v1alpha2.LimitChange.removed:type_name -> ntt.limits.v1alpha2.LimitChange.Removed
	5, // 4: ntt.limits.v1alpha2.LimitChange.Added.limit:type_name -> ntt.limits.v1alpha2.Limit
	5, // 5: ntt.limits.v1alpha2.LimitChange.Modified.limit:type_name -> ntt.limits.v1alpha2.Limit
	6, // 6: ntt.limits.v1alpha2.LimitChange.Modified.field_mask:type_name -> ntt.limits.v1alpha2.Limit_FieldMask
	5, // 7: ntt.limits.v1alpha2.LimitChange.Current.limit:type_name -> ntt.limits.v1alpha2.Limit
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { edgelq_limits_proto_v1alpha2_limit_change_proto_init() }
func edgelq_limits_proto_v1alpha2_limit_change_proto_init() {
	if edgelq_limits_proto_v1alpha2_limit_change_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		edgelq_limits_proto_v1alpha2_limit_change_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LimitChange); i {
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
		edgelq_limits_proto_v1alpha2_limit_change_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LimitChange_Added); i {
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
		edgelq_limits_proto_v1alpha2_limit_change_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LimitChange_Modified); i {
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
		edgelq_limits_proto_v1alpha2_limit_change_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LimitChange_Current); i {
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
		edgelq_limits_proto_v1alpha2_limit_change_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LimitChange_Removed); i {
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

	edgelq_limits_proto_v1alpha2_limit_change_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*LimitChange_Added_)(nil),
		(*LimitChange_Modified_)(nil),
		(*LimitChange_Current_)(nil),
		(*LimitChange_Removed_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: edgelq_limits_proto_v1alpha2_limit_change_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           edgelq_limits_proto_v1alpha2_limit_change_proto_goTypes,
		DependencyIndexes: edgelq_limits_proto_v1alpha2_limit_change_proto_depIdxs,
		MessageInfos:      edgelq_limits_proto_v1alpha2_limit_change_proto_msgTypes,
	}.Build()
	edgelq_limits_proto_v1alpha2_limit_change_proto = out.File
	edgelq_limits_proto_v1alpha2_limit_change_proto_rawDesc = nil
	edgelq_limits_proto_v1alpha2_limit_change_proto_goTypes = nil
	edgelq_limits_proto_v1alpha2_limit_change_proto_depIdxs = nil
}
