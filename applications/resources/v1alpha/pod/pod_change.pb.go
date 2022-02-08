// Code generated by protoc-gen-goten-go
// File: edgelq/applications/proto/v1alpha/pod_change.proto
// DO NOT EDIT!!!

package pod

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
	project "github.com/cloudwan/edgelq-sdk/applications/resources/v1alpha/project"
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

// PodChange is used by Watch notifications Responses to describe change of
// single Pod One of Added, Modified, Removed
type PodChange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Pod change
	//
	// Types that are valid to be assigned to ChangeType:
	//	*PodChange_Added_
	//	*PodChange_Modified_
	//	*PodChange_Current_
	//	*PodChange_Removed_
	ChangeType isPodChange_ChangeType `protobuf_oneof:"change_type"`
}

func (m *PodChange) Reset() {
	*m = PodChange{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_applications_proto_v1alpha_pod_change_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *PodChange) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*PodChange) ProtoMessage() {}

func (m *PodChange) ProtoReflect() preflect.Message {
	mi := &edgelq_applications_proto_v1alpha_pod_change_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*PodChange) GotenMessage() {}

// Deprecated, Use PodChange.ProtoReflect.Descriptor instead.
func (*PodChange) Descriptor() ([]byte, []int) {
	return edgelq_applications_proto_v1alpha_pod_change_proto_rawDescGZIP(), []int{0}
}

func (m *PodChange) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *PodChange) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *PodChange) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *PodChange) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

type isPodChange_ChangeType interface {
	isPodChange_ChangeType()
}

type PodChange_Added_ struct {
	// Added is returned when watched document is added, either created or
	// enters Query view
	Added *PodChange_Added `protobuf:"bytes,1,opt,name=added,proto3,oneof" firestore:"added"`
}
type PodChange_Modified_ struct {
	// Modified is returned when watched document is modified
	Modified *PodChange_Modified `protobuf:"bytes,2,opt,name=modified,proto3,oneof" firestore:"modified"`
}
type PodChange_Current_ struct {
	// Current is returned in stateless watch when document enters query view or
	// is modified within.
	Current *PodChange_Current `protobuf:"bytes,4,opt,name=current,proto3,oneof" firestore:"current"`
}
type PodChange_Removed_ struct {
	// Removed is returned when Pod is deleted or leaves Query view
	Removed *PodChange_Removed `protobuf:"bytes,3,opt,name=removed,proto3,oneof" firestore:"removed"`
}

func (*PodChange_Added_) isPodChange_ChangeType()    {}
func (*PodChange_Modified_) isPodChange_ChangeType() {}
func (*PodChange_Current_) isPodChange_ChangeType()  {}
func (*PodChange_Removed_) isPodChange_ChangeType()  {}
func (m *PodChange) GetChangeType() isPodChange_ChangeType {
	if m != nil {
		return m.ChangeType
	}
	return nil
}
func (m *PodChange) GetAdded() *PodChange_Added {
	if x, ok := m.GetChangeType().(*PodChange_Added_); ok {
		return x.Added
	}
	return nil
}
func (m *PodChange) GetModified() *PodChange_Modified {
	if x, ok := m.GetChangeType().(*PodChange_Modified_); ok {
		return x.Modified
	}
	return nil
}
func (m *PodChange) GetCurrent() *PodChange_Current {
	if x, ok := m.GetChangeType().(*PodChange_Current_); ok {
		return x.Current
	}
	return nil
}
func (m *PodChange) GetRemoved() *PodChange_Removed {
	if x, ok := m.GetChangeType().(*PodChange_Removed_); ok {
		return x.Removed
	}
	return nil
}
func (m *PodChange) SetChangeType(ofv isPodChange_ChangeType) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "isPodChange_ChangeType", "PodChange"))
	}
	m.ChangeType = ofv
}
func (m *PodChange) SetAdded(fv *PodChange_Added) {
	m.SetChangeType(&PodChange_Added_{Added: fv})
}
func (m *PodChange) SetModified(fv *PodChange_Modified) {
	m.SetChangeType(&PodChange_Modified_{Modified: fv})
}
func (m *PodChange) SetCurrent(fv *PodChange_Current) {
	m.SetChangeType(&PodChange_Current_{Current: fv})
}
func (m *PodChange) SetRemoved(fv *PodChange_Removed) {
	m.SetChangeType(&PodChange_Removed_{Removed: fv})
}

// Pod has been added to query view
type PodChange_Added struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Pod           *Pod `protobuf:"bytes,1,opt,name=pod,proto3" json:"pod,omitempty" firestore:"pod"`
	// Integer describing index of added Pod in resulting query view.
	ViewIndex int32 `protobuf:"varint,2,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty" firestore:"viewIndex"`
}

func (m *PodChange_Added) Reset() {
	*m = PodChange_Added{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_applications_proto_v1alpha_pod_change_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *PodChange_Added) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*PodChange_Added) ProtoMessage() {}

func (m *PodChange_Added) ProtoReflect() preflect.Message {
	mi := &edgelq_applications_proto_v1alpha_pod_change_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*PodChange_Added) GotenMessage() {}

// Deprecated, Use PodChange_Added.ProtoReflect.Descriptor instead.
func (*PodChange_Added) Descriptor() ([]byte, []int) {
	return edgelq_applications_proto_v1alpha_pod_change_proto_rawDescGZIP(), []int{0, 0}
}

func (m *PodChange_Added) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *PodChange_Added) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *PodChange_Added) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *PodChange_Added) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *PodChange_Added) GetPod() *Pod {
	if m != nil {
		return m.Pod
	}
	return nil
}

func (m *PodChange_Added) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *PodChange_Added) SetPod(fv *Pod) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Pod", "PodChange_Added"))
	}
	m.Pod = fv
}

func (m *PodChange_Added) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "PodChange_Added"))
	}
	m.ViewIndex = fv
}

// Pod changed some of it's fields - contains either full document or masked
// change
type PodChange_Modified struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Name of modified Pod
	Name *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// New version of Pod or masked difference, depending on mask_changes
	// instrumentation of issued [WatchPodRequest] or [WatchPodsRequest]
	Pod *Pod `protobuf:"bytes,2,opt,name=pod,proto3" json:"pod,omitempty" firestore:"pod"`
	// Used when mask_changes is set, contains field paths of modified
	// properties.
	FieldMask *Pod_FieldMask `protobuf:"bytes,3,opt,customtype=Pod_FieldMask,name=field_mask,json=fieldMask,proto3" json:"field_mask,omitempty" firestore:"fieldMask"`
	// Previous view index specifies previous position of modified Pod.
	// When modification doesn't affect sorted order, value will remain
	// identical to [view_index].
	PreviousViewIndex int32 `protobuf:"varint,4,opt,name=previous_view_index,json=previousViewIndex,proto3" json:"previous_view_index,omitempty" firestore:"previousViewIndex"`
	// Integer specifying Pod new index in resulting query view.
	ViewIndex int32 `protobuf:"varint,5,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty" firestore:"viewIndex"`
}

func (m *PodChange_Modified) Reset() {
	*m = PodChange_Modified{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_applications_proto_v1alpha_pod_change_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *PodChange_Modified) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*PodChange_Modified) ProtoMessage() {}

func (m *PodChange_Modified) ProtoReflect() preflect.Message {
	mi := &edgelq_applications_proto_v1alpha_pod_change_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*PodChange_Modified) GotenMessage() {}

// Deprecated, Use PodChange_Modified.ProtoReflect.Descriptor instead.
func (*PodChange_Modified) Descriptor() ([]byte, []int) {
	return edgelq_applications_proto_v1alpha_pod_change_proto_rawDescGZIP(), []int{0, 1}
}

func (m *PodChange_Modified) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *PodChange_Modified) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *PodChange_Modified) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *PodChange_Modified) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *PodChange_Modified) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *PodChange_Modified) GetPod() *Pod {
	if m != nil {
		return m.Pod
	}
	return nil
}

func (m *PodChange_Modified) GetFieldMask() *Pod_FieldMask {
	if m != nil {
		return m.FieldMask
	}
	return nil
}

func (m *PodChange_Modified) GetPreviousViewIndex() int32 {
	if m != nil {
		return m.PreviousViewIndex
	}
	return int32(0)
}

func (m *PodChange_Modified) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *PodChange_Modified) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "PodChange_Modified"))
	}
	m.Name = fv
}

func (m *PodChange_Modified) SetPod(fv *Pod) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Pod", "PodChange_Modified"))
	}
	m.Pod = fv
}

func (m *PodChange_Modified) SetFieldMask(fv *Pod_FieldMask) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "FieldMask", "PodChange_Modified"))
	}
	m.FieldMask = fv
}

func (m *PodChange_Modified) SetPreviousViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "PreviousViewIndex", "PodChange_Modified"))
	}
	m.PreviousViewIndex = fv
}

func (m *PodChange_Modified) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "PodChange_Modified"))
	}
	m.ViewIndex = fv
}

// Pod has been added or modified in a query view. Version used for stateless
// watching
type PodChange_Current struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Pod           *Pod `protobuf:"bytes,1,opt,name=pod,proto3" json:"pod,omitempty" firestore:"pod"`
}

func (m *PodChange_Current) Reset() {
	*m = PodChange_Current{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_applications_proto_v1alpha_pod_change_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *PodChange_Current) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*PodChange_Current) ProtoMessage() {}

func (m *PodChange_Current) ProtoReflect() preflect.Message {
	mi := &edgelq_applications_proto_v1alpha_pod_change_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*PodChange_Current) GotenMessage() {}

// Deprecated, Use PodChange_Current.ProtoReflect.Descriptor instead.
func (*PodChange_Current) Descriptor() ([]byte, []int) {
	return edgelq_applications_proto_v1alpha_pod_change_proto_rawDescGZIP(), []int{0, 2}
}

func (m *PodChange_Current) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *PodChange_Current) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *PodChange_Current) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *PodChange_Current) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *PodChange_Current) GetPod() *Pod {
	if m != nil {
		return m.Pod
	}
	return nil
}

func (m *PodChange_Current) SetPod(fv *Pod) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Pod", "PodChange_Current"))
	}
	m.Pod = fv
}

// Removed is returned when Pod is deleted or leaves Query view
type PodChange_Removed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Name          *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// Integer specifying removed Pod index. Not populated in stateless watch
	// type.
	ViewIndex int32 `protobuf:"varint,2,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty" firestore:"viewIndex"`
}

func (m *PodChange_Removed) Reset() {
	*m = PodChange_Removed{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_applications_proto_v1alpha_pod_change_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *PodChange_Removed) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*PodChange_Removed) ProtoMessage() {}

func (m *PodChange_Removed) ProtoReflect() preflect.Message {
	mi := &edgelq_applications_proto_v1alpha_pod_change_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*PodChange_Removed) GotenMessage() {}

// Deprecated, Use PodChange_Removed.ProtoReflect.Descriptor instead.
func (*PodChange_Removed) Descriptor() ([]byte, []int) {
	return edgelq_applications_proto_v1alpha_pod_change_proto_rawDescGZIP(), []int{0, 3}
}

func (m *PodChange_Removed) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *PodChange_Removed) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *PodChange_Removed) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *PodChange_Removed) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *PodChange_Removed) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *PodChange_Removed) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *PodChange_Removed) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "PodChange_Removed"))
	}
	m.Name = fv
}

func (m *PodChange_Removed) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "PodChange_Removed"))
	}
	m.ViewIndex = fv
}

var edgelq_applications_proto_v1alpha_pod_change_proto preflect.FileDescriptor

var edgelq_applications_proto_v1alpha_pod_change_proto_rawDesc = []byte{
	0x0a, 0x32, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2f, 0x70, 0x6f, 0x64, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x6e, 0x74, 0x74, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x19,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74,
	0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67,
	0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65,
	0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x67, 0x6f,
	0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x2b, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2f, 0x70, 0x6f, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xab, 0x06,
	0x0a, 0x09, 0x50, 0x6f, 0x64, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x41, 0x0a, 0x05, 0x61,
	0x64, 0x64, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x6e, 0x74, 0x74,
	0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x50, 0x6f, 0x64, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e,
	0x41, 0x64, 0x64, 0x65, 0x64, 0x48, 0x00, 0x52, 0x05, 0x61, 0x64, 0x64, 0x65, 0x64, 0x12, 0x4a,
	0x0a, 0x08, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2c, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x50, 0x6f, 0x64, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x48, 0x00,
	0x52, 0x08, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x47, 0x0a, 0x07, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x6e, 0x74,
	0x74, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x50, 0x6f, 0x64, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x07, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x12, 0x47, 0x0a, 0x07, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e,
	0x50, 0x6f, 0x64, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x64, 0x48, 0x00, 0x52, 0x07, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x1a, 0x57, 0x0a, 0x05,
	0x41, 0x64, 0x64, 0x65, 0x64, 0x12, 0x2f, 0x0a, 0x03, 0x70, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x50, 0x6f,
	0x64, 0x52, 0x03, 0x70, 0x6f, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x76, 0x69, 0x65, 0x77,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x1a, 0xf3, 0x01, 0x0a, 0x08, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69,
	0x65, 0x64, 0x12, 0x1f, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x0b, 0xb2, 0xda, 0x21, 0x07, 0x0a, 0x05, 0x0a, 0x03, 0x50, 0x6f, 0x64, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x2f, 0x0a, 0x03, 0x70, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x50, 0x6f, 0x64, 0x52,
	0x03, 0x70, 0x6f, 0x64, 0x12, 0x46, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61,
	0x73, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x4d, 0x61, 0x73, 0x6b, 0x42, 0x0b, 0xb2, 0xda, 0x21, 0x07, 0x32, 0x05, 0x0a, 0x03, 0x50, 0x6f,
	0x64, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x2e, 0x0a, 0x13,
	0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x5f, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x70, 0x72, 0x65, 0x76, 0x69,
	0x6f, 0x75, 0x73, 0x56, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1d, 0x0a, 0x0a,
	0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x76, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x1a, 0x3a, 0x0a, 0x07, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x2f, 0x0a, 0x03, 0x70, 0x6f, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x50,
	0x6f, 0x64, 0x52, 0x03, 0x70, 0x6f, 0x64, 0x1a, 0x49, 0x0a, 0x07, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x64, 0x12, 0x1f, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x0b, 0xb2, 0xda, 0x21, 0x07, 0x0a, 0x05, 0x0a, 0x03, 0x50, 0x6f, 0x64, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x76, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x3a, 0x18, 0x9a, 0xd9, 0x21, 0x05, 0x0a, 0x03, 0x50, 0x6f, 0x64, 0xfa, 0xde, 0x21,
	0x0b, 0x0a, 0x09, 0x50, 0x6f, 0x64, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x42, 0x0d, 0x0a, 0x0b,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x42, 0xcf, 0x01, 0xe8, 0xde,
	0x21, 0x00, 0x92, 0x8c, 0xd1, 0x02, 0x50, 0x0a, 0x0e, 0x70, 0x6f, 0x64, 0x5f, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67,
	0x65, 0x6c, 0x71, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2f, 0x70, 0x6f, 0x64, 0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x74, 0x74,
	0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x62,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x42, 0x0e, 0x50, 0x6f, 0x64, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x00, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f,
	0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x70, 0x6f, 0x64, 0x3b, 0x70, 0x6f, 0x64, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	edgelq_applications_proto_v1alpha_pod_change_proto_rawDescOnce sync.Once
	edgelq_applications_proto_v1alpha_pod_change_proto_rawDescData = edgelq_applications_proto_v1alpha_pod_change_proto_rawDesc
)

func edgelq_applications_proto_v1alpha_pod_change_proto_rawDescGZIP() []byte {
	edgelq_applications_proto_v1alpha_pod_change_proto_rawDescOnce.Do(func() {
		edgelq_applications_proto_v1alpha_pod_change_proto_rawDescData = protoimpl.X.CompressGZIP(edgelq_applications_proto_v1alpha_pod_change_proto_rawDescData)
	})
	return edgelq_applications_proto_v1alpha_pod_change_proto_rawDescData
}

var edgelq_applications_proto_v1alpha_pod_change_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var edgelq_applications_proto_v1alpha_pod_change_proto_goTypes = []interface{}{
	(*PodChange)(nil),          // 0: ntt.applications.v1alpha.PodChange
	(*PodChange_Added)(nil),    // 1: ntt.applications.v1alpha.PodChange.Added
	(*PodChange_Modified)(nil), // 2: ntt.applications.v1alpha.PodChange.Modified
	(*PodChange_Current)(nil),  // 3: ntt.applications.v1alpha.PodChange.Current
	(*PodChange_Removed)(nil),  // 4: ntt.applications.v1alpha.PodChange.Removed
	(*Pod)(nil),                // 5: ntt.applications.v1alpha.Pod
	(*Pod_FieldMask)(nil),      // 6: ntt.applications.v1alpha.Pod_FieldMask
}
var edgelq_applications_proto_v1alpha_pod_change_proto_depIdxs = []int32{
	1, // 0: ntt.applications.v1alpha.PodChange.added:type_name -> ntt.applications.v1alpha.PodChange.Added
	2, // 1: ntt.applications.v1alpha.PodChange.modified:type_name -> ntt.applications.v1alpha.PodChange.Modified
	3, // 2: ntt.applications.v1alpha.PodChange.current:type_name -> ntt.applications.v1alpha.PodChange.Current
	4, // 3: ntt.applications.v1alpha.PodChange.removed:type_name -> ntt.applications.v1alpha.PodChange.Removed
	5, // 4: ntt.applications.v1alpha.PodChange.Added.pod:type_name -> ntt.applications.v1alpha.Pod
	5, // 5: ntt.applications.v1alpha.PodChange.Modified.pod:type_name -> ntt.applications.v1alpha.Pod
	6, // 6: ntt.applications.v1alpha.PodChange.Modified.field_mask:type_name -> ntt.applications.v1alpha.Pod_FieldMask
	5, // 7: ntt.applications.v1alpha.PodChange.Current.pod:type_name -> ntt.applications.v1alpha.Pod
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { edgelq_applications_proto_v1alpha_pod_change_proto_init() }
func edgelq_applications_proto_v1alpha_pod_change_proto_init() {
	if edgelq_applications_proto_v1alpha_pod_change_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		edgelq_applications_proto_v1alpha_pod_change_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PodChange); i {
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
		edgelq_applications_proto_v1alpha_pod_change_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PodChange_Added); i {
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
		edgelq_applications_proto_v1alpha_pod_change_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PodChange_Modified); i {
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
		edgelq_applications_proto_v1alpha_pod_change_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PodChange_Current); i {
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
		edgelq_applications_proto_v1alpha_pod_change_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PodChange_Removed); i {
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

	edgelq_applications_proto_v1alpha_pod_change_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*PodChange_Added_)(nil),
		(*PodChange_Modified_)(nil),
		(*PodChange_Current_)(nil),
		(*PodChange_Removed_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: edgelq_applications_proto_v1alpha_pod_change_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           edgelq_applications_proto_v1alpha_pod_change_proto_goTypes,
		DependencyIndexes: edgelq_applications_proto_v1alpha_pod_change_proto_depIdxs,
		MessageInfos:      edgelq_applications_proto_v1alpha_pod_change_proto_msgTypes,
	}.Build()
	edgelq_applications_proto_v1alpha_pod_change_proto = out.File
	edgelq_applications_proto_v1alpha_pod_change_proto_rawDesc = nil
	edgelq_applications_proto_v1alpha_pod_change_proto_goTypes = nil
	edgelq_applications_proto_v1alpha_pod_change_proto_depIdxs = nil
}