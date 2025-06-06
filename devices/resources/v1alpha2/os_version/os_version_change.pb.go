// Code generated by protoc-gen-goten-go
// File: edgelq/devices/proto/v1alpha2/os_version_change.proto
// DO NOT EDIT!!!

package os_version

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

// OsVersionChange is used by Watch notifications Responses to describe change
// of single OsVersion One of Added, Modified, Removed
type OsVersionChange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// OsVersion change
	//
	// Types that are valid to be assigned to ChangeType:
	//	*OsVersionChange_Added_
	//	*OsVersionChange_Modified_
	//	*OsVersionChange_Current_
	//	*OsVersionChange_Removed_
	ChangeType isOsVersionChange_ChangeType `protobuf_oneof:"change_type"`
}

func (m *OsVersionChange) Reset() {
	*m = OsVersionChange{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_devices_proto_v1alpha2_os_version_change_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *OsVersionChange) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*OsVersionChange) ProtoMessage() {}

func (m *OsVersionChange) ProtoReflect() preflect.Message {
	mi := &edgelq_devices_proto_v1alpha2_os_version_change_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*OsVersionChange) GotenMessage() {}

// Deprecated, Use OsVersionChange.ProtoReflect.Descriptor instead.
func (*OsVersionChange) Descriptor() ([]byte, []int) {
	return edgelq_devices_proto_v1alpha2_os_version_change_proto_rawDescGZIP(), []int{0}
}

func (m *OsVersionChange) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *OsVersionChange) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *OsVersionChange) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *OsVersionChange) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

type isOsVersionChange_ChangeType interface {
	isOsVersionChange_ChangeType()
}

type OsVersionChange_Added_ struct {
	// Added is returned when watched document is added, either created or
	// enters Query view
	Added *OsVersionChange_Added `protobuf:"bytes,1,opt,name=added,proto3,oneof"`
}
type OsVersionChange_Modified_ struct {
	// Modified is returned when watched document is modified
	Modified *OsVersionChange_Modified `protobuf:"bytes,2,opt,name=modified,proto3,oneof"`
}
type OsVersionChange_Current_ struct {
	// Current is returned in stateless watch when document enters query view or
	// is modified within.
	Current *OsVersionChange_Current `protobuf:"bytes,4,opt,name=current,proto3,oneof"`
}
type OsVersionChange_Removed_ struct {
	// Removed is returned when OsVersion is deleted or leaves Query view
	Removed *OsVersionChange_Removed `protobuf:"bytes,3,opt,name=removed,proto3,oneof"`
}

func (*OsVersionChange_Added_) isOsVersionChange_ChangeType()    {}
func (*OsVersionChange_Modified_) isOsVersionChange_ChangeType() {}
func (*OsVersionChange_Current_) isOsVersionChange_ChangeType()  {}
func (*OsVersionChange_Removed_) isOsVersionChange_ChangeType()  {}
func (m *OsVersionChange) GetChangeType() isOsVersionChange_ChangeType {
	if m != nil {
		return m.ChangeType
	}
	return nil
}
func (m *OsVersionChange) GetAdded() *OsVersionChange_Added {
	if x, ok := m.GetChangeType().(*OsVersionChange_Added_); ok {
		return x.Added
	}
	return nil
}
func (m *OsVersionChange) GetModified() *OsVersionChange_Modified {
	if x, ok := m.GetChangeType().(*OsVersionChange_Modified_); ok {
		return x.Modified
	}
	return nil
}
func (m *OsVersionChange) GetCurrent() *OsVersionChange_Current {
	if x, ok := m.GetChangeType().(*OsVersionChange_Current_); ok {
		return x.Current
	}
	return nil
}
func (m *OsVersionChange) GetRemoved() *OsVersionChange_Removed {
	if x, ok := m.GetChangeType().(*OsVersionChange_Removed_); ok {
		return x.Removed
	}
	return nil
}
func (m *OsVersionChange) SetChangeType(ofv isOsVersionChange_ChangeType) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "isOsVersionChange_ChangeType", "OsVersionChange"))
	}
	m.ChangeType = ofv
}
func (m *OsVersionChange) SetAdded(fv *OsVersionChange_Added) {
	m.SetChangeType(&OsVersionChange_Added_{Added: fv})
}
func (m *OsVersionChange) SetModified(fv *OsVersionChange_Modified) {
	m.SetChangeType(&OsVersionChange_Modified_{Modified: fv})
}
func (m *OsVersionChange) SetCurrent(fv *OsVersionChange_Current) {
	m.SetChangeType(&OsVersionChange_Current_{Current: fv})
}
func (m *OsVersionChange) SetRemoved(fv *OsVersionChange_Removed) {
	m.SetChangeType(&OsVersionChange_Removed_{Removed: fv})
}

// OsVersion has been added to query view
type OsVersionChange_Added struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	OsVersion     *OsVersion `protobuf:"bytes,1,opt,name=os_version,json=osVersion,proto3" json:"os_version,omitempty"`
	// Integer describing index of added OsVersion in resulting query view.
	ViewIndex int32 `protobuf:"varint,2,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty"`
}

func (m *OsVersionChange_Added) Reset() {
	*m = OsVersionChange_Added{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_devices_proto_v1alpha2_os_version_change_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *OsVersionChange_Added) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*OsVersionChange_Added) ProtoMessage() {}

func (m *OsVersionChange_Added) ProtoReflect() preflect.Message {
	mi := &edgelq_devices_proto_v1alpha2_os_version_change_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*OsVersionChange_Added) GotenMessage() {}

// Deprecated, Use OsVersionChange_Added.ProtoReflect.Descriptor instead.
func (*OsVersionChange_Added) Descriptor() ([]byte, []int) {
	return edgelq_devices_proto_v1alpha2_os_version_change_proto_rawDescGZIP(), []int{0, 0}
}

func (m *OsVersionChange_Added) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *OsVersionChange_Added) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *OsVersionChange_Added) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *OsVersionChange_Added) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *OsVersionChange_Added) GetOsVersion() *OsVersion {
	if m != nil {
		return m.OsVersion
	}
	return nil
}

func (m *OsVersionChange_Added) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *OsVersionChange_Added) SetOsVersion(fv *OsVersion) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "OsVersion", "OsVersionChange_Added"))
	}
	m.OsVersion = fv
}

func (m *OsVersionChange_Added) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "OsVersionChange_Added"))
	}
	m.ViewIndex = fv
}

// OsVersion changed some of it's fields - contains either full document or
// masked change
type OsVersionChange_Modified struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Name of modified OsVersion
	Name *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty"`
	// New version of OsVersion or masked difference, depending on mask_changes
	// instrumentation of issued [WatchOsVersionRequest] or
	// [WatchOsVersionsRequest]
	OsVersion *OsVersion `protobuf:"bytes,2,opt,name=os_version,json=osVersion,proto3" json:"os_version,omitempty"`
	// Used when mask_changes is set, contains field paths of modified
	// properties.
	FieldMask *OsVersion_FieldMask `protobuf:"bytes,3,opt,customtype=OsVersion_FieldMask,name=field_mask,json=fieldMask,proto3" json:"field_mask,omitempty"`
	// Previous view index specifies previous position of modified OsVersion.
	// When modification doesn't affect sorted order, value will remain
	// identical to [view_index].
	PreviousViewIndex int32 `protobuf:"varint,4,opt,name=previous_view_index,json=previousViewIndex,proto3" json:"previous_view_index,omitempty"`
	// Integer specifying OsVersion new index in resulting query view.
	ViewIndex int32 `protobuf:"varint,5,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty"`
}

func (m *OsVersionChange_Modified) Reset() {
	*m = OsVersionChange_Modified{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_devices_proto_v1alpha2_os_version_change_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *OsVersionChange_Modified) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*OsVersionChange_Modified) ProtoMessage() {}

func (m *OsVersionChange_Modified) ProtoReflect() preflect.Message {
	mi := &edgelq_devices_proto_v1alpha2_os_version_change_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*OsVersionChange_Modified) GotenMessage() {}

// Deprecated, Use OsVersionChange_Modified.ProtoReflect.Descriptor instead.
func (*OsVersionChange_Modified) Descriptor() ([]byte, []int) {
	return edgelq_devices_proto_v1alpha2_os_version_change_proto_rawDescGZIP(), []int{0, 1}
}

func (m *OsVersionChange_Modified) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *OsVersionChange_Modified) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *OsVersionChange_Modified) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *OsVersionChange_Modified) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *OsVersionChange_Modified) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *OsVersionChange_Modified) GetOsVersion() *OsVersion {
	if m != nil {
		return m.OsVersion
	}
	return nil
}

func (m *OsVersionChange_Modified) GetFieldMask() *OsVersion_FieldMask {
	if m != nil {
		return m.FieldMask
	}
	return nil
}

func (m *OsVersionChange_Modified) GetPreviousViewIndex() int32 {
	if m != nil {
		return m.PreviousViewIndex
	}
	return int32(0)
}

func (m *OsVersionChange_Modified) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *OsVersionChange_Modified) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "OsVersionChange_Modified"))
	}
	m.Name = fv
}

func (m *OsVersionChange_Modified) SetOsVersion(fv *OsVersion) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "OsVersion", "OsVersionChange_Modified"))
	}
	m.OsVersion = fv
}

func (m *OsVersionChange_Modified) SetFieldMask(fv *OsVersion_FieldMask) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "FieldMask", "OsVersionChange_Modified"))
	}
	m.FieldMask = fv
}

func (m *OsVersionChange_Modified) SetPreviousViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "PreviousViewIndex", "OsVersionChange_Modified"))
	}
	m.PreviousViewIndex = fv
}

func (m *OsVersionChange_Modified) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "OsVersionChange_Modified"))
	}
	m.ViewIndex = fv
}

// OsVersion has been added or modified in a query view. Version used for
// stateless watching
type OsVersionChange_Current struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	OsVersion     *OsVersion `protobuf:"bytes,1,opt,name=os_version,json=osVersion,proto3" json:"os_version,omitempty"`
}

func (m *OsVersionChange_Current) Reset() {
	*m = OsVersionChange_Current{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_devices_proto_v1alpha2_os_version_change_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *OsVersionChange_Current) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*OsVersionChange_Current) ProtoMessage() {}

func (m *OsVersionChange_Current) ProtoReflect() preflect.Message {
	mi := &edgelq_devices_proto_v1alpha2_os_version_change_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*OsVersionChange_Current) GotenMessage() {}

// Deprecated, Use OsVersionChange_Current.ProtoReflect.Descriptor instead.
func (*OsVersionChange_Current) Descriptor() ([]byte, []int) {
	return edgelq_devices_proto_v1alpha2_os_version_change_proto_rawDescGZIP(), []int{0, 2}
}

func (m *OsVersionChange_Current) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *OsVersionChange_Current) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *OsVersionChange_Current) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *OsVersionChange_Current) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *OsVersionChange_Current) GetOsVersion() *OsVersion {
	if m != nil {
		return m.OsVersion
	}
	return nil
}

func (m *OsVersionChange_Current) SetOsVersion(fv *OsVersion) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "OsVersion", "OsVersionChange_Current"))
	}
	m.OsVersion = fv
}

// Removed is returned when OsVersion is deleted or leaves Query view
type OsVersionChange_Removed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Name          *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty"`
	// Integer specifying removed OsVersion index. Not populated in stateless
	// watch type.
	ViewIndex int32 `protobuf:"varint,2,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty"`
}

func (m *OsVersionChange_Removed) Reset() {
	*m = OsVersionChange_Removed{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_devices_proto_v1alpha2_os_version_change_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *OsVersionChange_Removed) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*OsVersionChange_Removed) ProtoMessage() {}

func (m *OsVersionChange_Removed) ProtoReflect() preflect.Message {
	mi := &edgelq_devices_proto_v1alpha2_os_version_change_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*OsVersionChange_Removed) GotenMessage() {}

// Deprecated, Use OsVersionChange_Removed.ProtoReflect.Descriptor instead.
func (*OsVersionChange_Removed) Descriptor() ([]byte, []int) {
	return edgelq_devices_proto_v1alpha2_os_version_change_proto_rawDescGZIP(), []int{0, 3}
}

func (m *OsVersionChange_Removed) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *OsVersionChange_Removed) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *OsVersionChange_Removed) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *OsVersionChange_Removed) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *OsVersionChange_Removed) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *OsVersionChange_Removed) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *OsVersionChange_Removed) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "OsVersionChange_Removed"))
	}
	m.Name = fv
}

func (m *OsVersionChange_Removed) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "OsVersionChange_Removed"))
	}
	m.ViewIndex = fv
}

var edgelq_devices_proto_v1alpha2_os_version_change_proto preflect.FileDescriptor

var edgelq_devices_proto_v1alpha2_os_version_change_proto_rawDesc = []byte{
	0x0a, 0x35, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f,
	0x6f, 0x73, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x6e, 0x74, 0x74, 0x2e, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x1a, 0x19, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65,
	0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f,
	0x74, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x65, 0x64, 0x67,
	0x65, 0x6c, 0x71, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x6f, 0x73, 0x5f, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xef, 0x06, 0x0a, 0x0f,
	0x4f, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12,
	0x43, 0x0a, 0x05, 0x61, 0x64, 0x64, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b,
	0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x4f, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x65, 0x64, 0x48, 0x00, 0x52, 0x05, 0x61,
	0x64, 0x64, 0x65, 0x64, 0x12, 0x4c, 0x0a, 0x08, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x4f, 0x73,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x4d, 0x6f,
	0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x48, 0x00, 0x52, 0x08, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69,
	0x65, 0x64, 0x12, 0x49, 0x0a, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x4f, 0x73, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x48, 0x00, 0x52, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x49, 0x0a,
	0x07, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d,
	0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x4f, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x48, 0x00, 0x52,
	0x07, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x1a, 0x66, 0x0a, 0x05, 0x41, 0x64, 0x64, 0x65,
	0x64, 0x12, 0x3e, 0x0a, 0x0a, 0x6f, 0x73, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x4f, 0x73, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x6f, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x76, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x1a, 0x8e, 0x02, 0x0a, 0x08, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x25, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x11, 0xb2, 0xda, 0x21,
	0x0d, 0x0a, 0x0b, 0x0a, 0x09, 0x4f, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3e, 0x0a, 0x0a, 0x6f, 0x73, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e,
	0x4f, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x6f, 0x73, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x4c, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61,
	0x73, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x4d, 0x61, 0x73, 0x6b, 0x42, 0x11, 0xb2, 0xda, 0x21, 0x0d, 0x32, 0x0b, 0x0a, 0x09, 0x4f, 0x73,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61,
	0x73, 0x6b, 0x12, 0x2e, 0x0a, 0x13, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x5f, 0x76,
	0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x11, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x56, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x76, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x1a, 0x49, 0x0a, 0x07, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x3e, 0x0a, 0x0a,
	0x6f, 0x73, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1f, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x4f, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x09, 0x6f, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x1a, 0x4f, 0x0a, 0x07,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x12, 0x25, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x11, 0xb2, 0xda, 0x21, 0x0d, 0x0a, 0x0b, 0x0a, 0x09, 0x4f,
	0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x76, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x3a, 0x0f, 0x9a,
	0xd9, 0x21, 0x0b, 0x0a, 0x09, 0x4f, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x0d,
	0x0a, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x42, 0x86, 0x01,
	0xe8, 0xde, 0x21, 0x00, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x32, 0x42, 0x14, 0x4f, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x00, 0x5a, 0x4b, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65,
	0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32,
	0x2f, 0x6f, 0x73, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x3b, 0x6f, 0x73, 0x5f, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	edgelq_devices_proto_v1alpha2_os_version_change_proto_rawDescOnce sync.Once
	edgelq_devices_proto_v1alpha2_os_version_change_proto_rawDescData = edgelq_devices_proto_v1alpha2_os_version_change_proto_rawDesc
)

func edgelq_devices_proto_v1alpha2_os_version_change_proto_rawDescGZIP() []byte {
	edgelq_devices_proto_v1alpha2_os_version_change_proto_rawDescOnce.Do(func() {
		edgelq_devices_proto_v1alpha2_os_version_change_proto_rawDescData = protoimpl.X.CompressGZIP(edgelq_devices_proto_v1alpha2_os_version_change_proto_rawDescData)
	})
	return edgelq_devices_proto_v1alpha2_os_version_change_proto_rawDescData
}

var edgelq_devices_proto_v1alpha2_os_version_change_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var edgelq_devices_proto_v1alpha2_os_version_change_proto_goTypes = []interface{}{
	(*OsVersionChange)(nil),          // 0: ntt.devices.v1alpha2.OsVersionChange
	(*OsVersionChange_Added)(nil),    // 1: ntt.devices.v1alpha2.OsVersionChange.Added
	(*OsVersionChange_Modified)(nil), // 2: ntt.devices.v1alpha2.OsVersionChange.Modified
	(*OsVersionChange_Current)(nil),  // 3: ntt.devices.v1alpha2.OsVersionChange.Current
	(*OsVersionChange_Removed)(nil),  // 4: ntt.devices.v1alpha2.OsVersionChange.Removed
	(*OsVersion)(nil),                // 5: ntt.devices.v1alpha2.OsVersion
	(*OsVersion_FieldMask)(nil),      // 6: ntt.devices.v1alpha2.OsVersion_FieldMask
}
var edgelq_devices_proto_v1alpha2_os_version_change_proto_depIdxs = []int32{
	1, // 0: ntt.devices.v1alpha2.OsVersionChange.added:type_name -> ntt.devices.v1alpha2.OsVersionChange.Added
	2, // 1: ntt.devices.v1alpha2.OsVersionChange.modified:type_name -> ntt.devices.v1alpha2.OsVersionChange.Modified
	3, // 2: ntt.devices.v1alpha2.OsVersionChange.current:type_name -> ntt.devices.v1alpha2.OsVersionChange.Current
	4, // 3: ntt.devices.v1alpha2.OsVersionChange.removed:type_name -> ntt.devices.v1alpha2.OsVersionChange.Removed
	5, // 4: ntt.devices.v1alpha2.OsVersionChange.Added.os_version:type_name -> ntt.devices.v1alpha2.OsVersion
	5, // 5: ntt.devices.v1alpha2.OsVersionChange.Modified.os_version:type_name -> ntt.devices.v1alpha2.OsVersion
	6, // 6: ntt.devices.v1alpha2.OsVersionChange.Modified.field_mask:type_name -> ntt.devices.v1alpha2.OsVersion_FieldMask
	5, // 7: ntt.devices.v1alpha2.OsVersionChange.Current.os_version:type_name -> ntt.devices.v1alpha2.OsVersion
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { edgelq_devices_proto_v1alpha2_os_version_change_proto_init() }
func edgelq_devices_proto_v1alpha2_os_version_change_proto_init() {
	if edgelq_devices_proto_v1alpha2_os_version_change_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		edgelq_devices_proto_v1alpha2_os_version_change_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OsVersionChange); i {
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
		edgelq_devices_proto_v1alpha2_os_version_change_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OsVersionChange_Added); i {
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
		edgelq_devices_proto_v1alpha2_os_version_change_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OsVersionChange_Modified); i {
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
		edgelq_devices_proto_v1alpha2_os_version_change_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OsVersionChange_Current); i {
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
		edgelq_devices_proto_v1alpha2_os_version_change_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OsVersionChange_Removed); i {
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

	edgelq_devices_proto_v1alpha2_os_version_change_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*OsVersionChange_Added_)(nil),
		(*OsVersionChange_Modified_)(nil),
		(*OsVersionChange_Current_)(nil),
		(*OsVersionChange_Removed_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: edgelq_devices_proto_v1alpha2_os_version_change_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           edgelq_devices_proto_v1alpha2_os_version_change_proto_goTypes,
		DependencyIndexes: edgelq_devices_proto_v1alpha2_os_version_change_proto_depIdxs,
		MessageInfos:      edgelq_devices_proto_v1alpha2_os_version_change_proto_msgTypes,
	}.Build()
	edgelq_devices_proto_v1alpha2_os_version_change_proto = out.File
	edgelq_devices_proto_v1alpha2_os_version_change_proto_rawDesc = nil
	edgelq_devices_proto_v1alpha2_os_version_change_proto_goTypes = nil
	edgelq_devices_proto_v1alpha2_os_version_change_proto_depIdxs = nil
}
