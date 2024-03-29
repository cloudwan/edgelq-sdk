// Code generated by protoc-gen-goten-go
// File: edgelq/logging/proto/v1/bucket_change.proto
// DO NOT EDIT!!!

package bucket

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
	iam_organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1/organization"
	iam_project "github.com/cloudwan/edgelq-sdk/iam/resources/v1/project"
	meta_service "github.com/cloudwan/goten-sdk/meta-service/resources/v1/service"
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
	_ = &iam_organization.Organization{}
	_ = &iam_project.Project{}
	_ = &fieldmaskpb.FieldMask{}
	_ = &meta_service.Service{}
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// BucketChange is used by Watch notifications Responses to describe change of
// single Bucket One of Added, Modified, Removed
type BucketChange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Bucket change
	//
	// Types that are valid to be assigned to ChangeType:
	//	*BucketChange_Added_
	//	*BucketChange_Modified_
	//	*BucketChange_Current_
	//	*BucketChange_Removed_
	ChangeType isBucketChange_ChangeType `protobuf_oneof:"change_type"`
}

func (m *BucketChange) Reset() {
	*m = BucketChange{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_logging_proto_v1_bucket_change_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *BucketChange) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*BucketChange) ProtoMessage() {}

func (m *BucketChange) ProtoReflect() preflect.Message {
	mi := &edgelq_logging_proto_v1_bucket_change_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*BucketChange) GotenMessage() {}

// Deprecated, Use BucketChange.ProtoReflect.Descriptor instead.
func (*BucketChange) Descriptor() ([]byte, []int) {
	return edgelq_logging_proto_v1_bucket_change_proto_rawDescGZIP(), []int{0}
}

func (m *BucketChange) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *BucketChange) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *BucketChange) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *BucketChange) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

type isBucketChange_ChangeType interface {
	isBucketChange_ChangeType()
}

type BucketChange_Added_ struct {
	// Added is returned when watched document is added, either created or
	// enters Query view
	Added *BucketChange_Added `protobuf:"bytes,1,opt,name=added,proto3,oneof" firestore:"added"`
}
type BucketChange_Modified_ struct {
	// Modified is returned when watched document is modified
	Modified *BucketChange_Modified `protobuf:"bytes,2,opt,name=modified,proto3,oneof" firestore:"modified"`
}
type BucketChange_Current_ struct {
	// Current is returned in stateless watch when document enters query view or
	// is modified within.
	Current *BucketChange_Current `protobuf:"bytes,4,opt,name=current,proto3,oneof" firestore:"current"`
}
type BucketChange_Removed_ struct {
	// Removed is returned when Bucket is deleted or leaves Query view
	Removed *BucketChange_Removed `protobuf:"bytes,3,opt,name=removed,proto3,oneof" firestore:"removed"`
}

func (*BucketChange_Added_) isBucketChange_ChangeType()    {}
func (*BucketChange_Modified_) isBucketChange_ChangeType() {}
func (*BucketChange_Current_) isBucketChange_ChangeType()  {}
func (*BucketChange_Removed_) isBucketChange_ChangeType()  {}
func (m *BucketChange) GetChangeType() isBucketChange_ChangeType {
	if m != nil {
		return m.ChangeType
	}
	return nil
}
func (m *BucketChange) GetAdded() *BucketChange_Added {
	if x, ok := m.GetChangeType().(*BucketChange_Added_); ok {
		return x.Added
	}
	return nil
}
func (m *BucketChange) GetModified() *BucketChange_Modified {
	if x, ok := m.GetChangeType().(*BucketChange_Modified_); ok {
		return x.Modified
	}
	return nil
}
func (m *BucketChange) GetCurrent() *BucketChange_Current {
	if x, ok := m.GetChangeType().(*BucketChange_Current_); ok {
		return x.Current
	}
	return nil
}
func (m *BucketChange) GetRemoved() *BucketChange_Removed {
	if x, ok := m.GetChangeType().(*BucketChange_Removed_); ok {
		return x.Removed
	}
	return nil
}
func (m *BucketChange) SetChangeType(ofv isBucketChange_ChangeType) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "isBucketChange_ChangeType", "BucketChange"))
	}
	m.ChangeType = ofv
}
func (m *BucketChange) SetAdded(fv *BucketChange_Added) {
	m.SetChangeType(&BucketChange_Added_{Added: fv})
}
func (m *BucketChange) SetModified(fv *BucketChange_Modified) {
	m.SetChangeType(&BucketChange_Modified_{Modified: fv})
}
func (m *BucketChange) SetCurrent(fv *BucketChange_Current) {
	m.SetChangeType(&BucketChange_Current_{Current: fv})
}
func (m *BucketChange) SetRemoved(fv *BucketChange_Removed) {
	m.SetChangeType(&BucketChange_Removed_{Removed: fv})
}

// Bucket has been added to query view
type BucketChange_Added struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Bucket        *Bucket `protobuf:"bytes,1,opt,name=bucket,proto3" json:"bucket,omitempty" firestore:"bucket"`
	// Integer describing index of added Bucket in resulting query view.
	ViewIndex int32 `protobuf:"varint,2,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty" firestore:"viewIndex"`
}

func (m *BucketChange_Added) Reset() {
	*m = BucketChange_Added{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_logging_proto_v1_bucket_change_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *BucketChange_Added) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*BucketChange_Added) ProtoMessage() {}

func (m *BucketChange_Added) ProtoReflect() preflect.Message {
	mi := &edgelq_logging_proto_v1_bucket_change_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*BucketChange_Added) GotenMessage() {}

// Deprecated, Use BucketChange_Added.ProtoReflect.Descriptor instead.
func (*BucketChange_Added) Descriptor() ([]byte, []int) {
	return edgelq_logging_proto_v1_bucket_change_proto_rawDescGZIP(), []int{0, 0}
}

func (m *BucketChange_Added) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *BucketChange_Added) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *BucketChange_Added) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *BucketChange_Added) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *BucketChange_Added) GetBucket() *Bucket {
	if m != nil {
		return m.Bucket
	}
	return nil
}

func (m *BucketChange_Added) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *BucketChange_Added) SetBucket(fv *Bucket) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Bucket", "BucketChange_Added"))
	}
	m.Bucket = fv
}

func (m *BucketChange_Added) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "BucketChange_Added"))
	}
	m.ViewIndex = fv
}

// Bucket changed some of it's fields - contains either full document or
// masked change
type BucketChange_Modified struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Name of modified Bucket
	Name *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// New version of Bucket or masked difference, depending on mask_changes
	// instrumentation of issued [WatchBucketRequest] or [WatchBucketsRequest]
	Bucket *Bucket `protobuf:"bytes,2,opt,name=bucket,proto3" json:"bucket,omitempty" firestore:"bucket"`
	// Used when mask_changes is set, contains field paths of modified
	// properties.
	FieldMask *Bucket_FieldMask `protobuf:"bytes,3,opt,customtype=Bucket_FieldMask,name=field_mask,json=fieldMask,proto3" json:"field_mask,omitempty" firestore:"fieldMask"`
	// Previous view index specifies previous position of modified Bucket.
	// When modification doesn't affect sorted order, value will remain
	// identical to [view_index].
	PreviousViewIndex int32 `protobuf:"varint,4,opt,name=previous_view_index,json=previousViewIndex,proto3" json:"previous_view_index,omitempty" firestore:"previousViewIndex"`
	// Integer specifying Bucket new index in resulting query view.
	ViewIndex int32 `protobuf:"varint,5,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty" firestore:"viewIndex"`
}

func (m *BucketChange_Modified) Reset() {
	*m = BucketChange_Modified{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_logging_proto_v1_bucket_change_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *BucketChange_Modified) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*BucketChange_Modified) ProtoMessage() {}

func (m *BucketChange_Modified) ProtoReflect() preflect.Message {
	mi := &edgelq_logging_proto_v1_bucket_change_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*BucketChange_Modified) GotenMessage() {}

// Deprecated, Use BucketChange_Modified.ProtoReflect.Descriptor instead.
func (*BucketChange_Modified) Descriptor() ([]byte, []int) {
	return edgelq_logging_proto_v1_bucket_change_proto_rawDescGZIP(), []int{0, 1}
}

func (m *BucketChange_Modified) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *BucketChange_Modified) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *BucketChange_Modified) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *BucketChange_Modified) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *BucketChange_Modified) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *BucketChange_Modified) GetBucket() *Bucket {
	if m != nil {
		return m.Bucket
	}
	return nil
}

func (m *BucketChange_Modified) GetFieldMask() *Bucket_FieldMask {
	if m != nil {
		return m.FieldMask
	}
	return nil
}

func (m *BucketChange_Modified) GetPreviousViewIndex() int32 {
	if m != nil {
		return m.PreviousViewIndex
	}
	return int32(0)
}

func (m *BucketChange_Modified) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *BucketChange_Modified) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "BucketChange_Modified"))
	}
	m.Name = fv
}

func (m *BucketChange_Modified) SetBucket(fv *Bucket) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Bucket", "BucketChange_Modified"))
	}
	m.Bucket = fv
}

func (m *BucketChange_Modified) SetFieldMask(fv *Bucket_FieldMask) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "FieldMask", "BucketChange_Modified"))
	}
	m.FieldMask = fv
}

func (m *BucketChange_Modified) SetPreviousViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "PreviousViewIndex", "BucketChange_Modified"))
	}
	m.PreviousViewIndex = fv
}

func (m *BucketChange_Modified) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "BucketChange_Modified"))
	}
	m.ViewIndex = fv
}

// Bucket has been added or modified in a query view. Version used for
// stateless watching
type BucketChange_Current struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Bucket        *Bucket `protobuf:"bytes,1,opt,name=bucket,proto3" json:"bucket,omitempty" firestore:"bucket"`
}

func (m *BucketChange_Current) Reset() {
	*m = BucketChange_Current{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_logging_proto_v1_bucket_change_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *BucketChange_Current) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*BucketChange_Current) ProtoMessage() {}

func (m *BucketChange_Current) ProtoReflect() preflect.Message {
	mi := &edgelq_logging_proto_v1_bucket_change_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*BucketChange_Current) GotenMessage() {}

// Deprecated, Use BucketChange_Current.ProtoReflect.Descriptor instead.
func (*BucketChange_Current) Descriptor() ([]byte, []int) {
	return edgelq_logging_proto_v1_bucket_change_proto_rawDescGZIP(), []int{0, 2}
}

func (m *BucketChange_Current) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *BucketChange_Current) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *BucketChange_Current) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *BucketChange_Current) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *BucketChange_Current) GetBucket() *Bucket {
	if m != nil {
		return m.Bucket
	}
	return nil
}

func (m *BucketChange_Current) SetBucket(fv *Bucket) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Bucket", "BucketChange_Current"))
	}
	m.Bucket = fv
}

// Removed is returned when Bucket is deleted or leaves Query view
type BucketChange_Removed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Name          *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// Integer specifying removed Bucket index. Not populated in stateless watch
	// type.
	ViewIndex int32 `protobuf:"varint,2,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty" firestore:"viewIndex"`
}

func (m *BucketChange_Removed) Reset() {
	*m = BucketChange_Removed{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_logging_proto_v1_bucket_change_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *BucketChange_Removed) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*BucketChange_Removed) ProtoMessage() {}

func (m *BucketChange_Removed) ProtoReflect() preflect.Message {
	mi := &edgelq_logging_proto_v1_bucket_change_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*BucketChange_Removed) GotenMessage() {}

// Deprecated, Use BucketChange_Removed.ProtoReflect.Descriptor instead.
func (*BucketChange_Removed) Descriptor() ([]byte, []int) {
	return edgelq_logging_proto_v1_bucket_change_proto_rawDescGZIP(), []int{0, 3}
}

func (m *BucketChange_Removed) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *BucketChange_Removed) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *BucketChange_Removed) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *BucketChange_Removed) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *BucketChange_Removed) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *BucketChange_Removed) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *BucketChange_Removed) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "BucketChange_Removed"))
	}
	m.Name = fv
}

func (m *BucketChange_Removed) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "BucketChange_Removed"))
	}
	m.ViewIndex = fv
}

var edgelq_logging_proto_v1_bucket_change_proto preflect.FileDescriptor

var edgelq_logging_proto_v1_bucket_change_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x6e,
	0x74, 0x74, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x1a, 0x19, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65,
	0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f,
	0x74, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x65, 0x64, 0x67,
	0x65, 0x6c, 0x71, 0x2f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x8c, 0x06, 0x0a, 0x0c, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x12, 0x3a, 0x0a, 0x05, 0x61, 0x64, 0x64, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x22, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x31, 0x2e, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e,
	0x41, 0x64, 0x64, 0x65, 0x64, 0x48, 0x00, 0x52, 0x05, 0x61, 0x64, 0x64, 0x65, 0x64, 0x12, 0x43,
	0x0a, 0x08, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x25, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76,
	0x31, 0x2e, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x4d,
	0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x48, 0x00, 0x52, 0x08, 0x6d, 0x6f, 0x64, 0x69, 0x66,
	0x69, 0x65, 0x64, 0x12, 0x40, 0x0a, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x69,
	0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x07, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x40, 0x0a, 0x07, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x6c, 0x6f, 0x67,
	0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x48, 0x00, 0x52, 0x07,
	0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x1a, 0x56, 0x0a, 0x05, 0x41, 0x64, 0x64, 0x65, 0x64,
	0x12, 0x2e, 0x0a, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76,
	0x31, 0x2e, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x76, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x1a,
	0xf8, 0x01, 0x0a, 0x08, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x22, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0e, 0xb2, 0xda, 0x21, 0x0a,
	0x0a, 0x08, 0x0a, 0x06, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x2e, 0x0a, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76,
	0x31, 0x2e, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x12, 0x49, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b,
	0x42, 0x0e, 0xb2, 0xda, 0x21, 0x0a, 0x32, 0x08, 0x0a, 0x06, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x2e, 0x0a, 0x13, 0x70,
	0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x5f, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f,
	0x75, 0x73, 0x56, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1d, 0x0a, 0x0a, 0x76,
	0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x76, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x1a, 0x39, 0x0a, 0x07, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x2e, 0x0a, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x6c, 0x6f, 0x67, 0x67,
	0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x06, 0x62,
	0x75, 0x63, 0x6b, 0x65, 0x74, 0x1a, 0x4c, 0x0a, 0x07, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64,
	0x12, 0x22, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0e,
	0xb2, 0xda, 0x21, 0x0a, 0x0a, 0x08, 0x0a, 0x06, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x76, 0x69, 0x65, 0x77, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x3a, 0x0c, 0x9a, 0xd9, 0x21, 0x08, 0x0a, 0x06, 0x42, 0x75, 0x63, 0x6b, 0x65,
	0x74, 0x42, 0x0d, 0x0a, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x42, 0x6f, 0xe8, 0xde, 0x21, 0x00, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x74, 0x74, 0x2e,
	0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x42, 0x11, 0x42,
	0x75, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x00, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6c,
	0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x3b, 0x62, 0x75, 0x63, 0x6b, 0x65,
	0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	edgelq_logging_proto_v1_bucket_change_proto_rawDescOnce sync.Once
	edgelq_logging_proto_v1_bucket_change_proto_rawDescData = edgelq_logging_proto_v1_bucket_change_proto_rawDesc
)

func edgelq_logging_proto_v1_bucket_change_proto_rawDescGZIP() []byte {
	edgelq_logging_proto_v1_bucket_change_proto_rawDescOnce.Do(func() {
		edgelq_logging_proto_v1_bucket_change_proto_rawDescData = protoimpl.X.CompressGZIP(edgelq_logging_proto_v1_bucket_change_proto_rawDescData)
	})
	return edgelq_logging_proto_v1_bucket_change_proto_rawDescData
}

var edgelq_logging_proto_v1_bucket_change_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var edgelq_logging_proto_v1_bucket_change_proto_goTypes = []interface{}{
	(*BucketChange)(nil),          // 0: ntt.logging.v1.BucketChange
	(*BucketChange_Added)(nil),    // 1: ntt.logging.v1.BucketChange.Added
	(*BucketChange_Modified)(nil), // 2: ntt.logging.v1.BucketChange.Modified
	(*BucketChange_Current)(nil),  // 3: ntt.logging.v1.BucketChange.Current
	(*BucketChange_Removed)(nil),  // 4: ntt.logging.v1.BucketChange.Removed
	(*Bucket)(nil),                // 5: ntt.logging.v1.Bucket
	(*Bucket_FieldMask)(nil),      // 6: ntt.logging.v1.Bucket_FieldMask
}
var edgelq_logging_proto_v1_bucket_change_proto_depIdxs = []int32{
	1, // 0: ntt.logging.v1.BucketChange.added:type_name -> ntt.logging.v1.BucketChange.Added
	2, // 1: ntt.logging.v1.BucketChange.modified:type_name -> ntt.logging.v1.BucketChange.Modified
	3, // 2: ntt.logging.v1.BucketChange.current:type_name -> ntt.logging.v1.BucketChange.Current
	4, // 3: ntt.logging.v1.BucketChange.removed:type_name -> ntt.logging.v1.BucketChange.Removed
	5, // 4: ntt.logging.v1.BucketChange.Added.bucket:type_name -> ntt.logging.v1.Bucket
	5, // 5: ntt.logging.v1.BucketChange.Modified.bucket:type_name -> ntt.logging.v1.Bucket
	6, // 6: ntt.logging.v1.BucketChange.Modified.field_mask:type_name -> ntt.logging.v1.Bucket_FieldMask
	5, // 7: ntt.logging.v1.BucketChange.Current.bucket:type_name -> ntt.logging.v1.Bucket
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { edgelq_logging_proto_v1_bucket_change_proto_init() }
func edgelq_logging_proto_v1_bucket_change_proto_init() {
	if edgelq_logging_proto_v1_bucket_change_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		edgelq_logging_proto_v1_bucket_change_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BucketChange); i {
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
		edgelq_logging_proto_v1_bucket_change_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BucketChange_Added); i {
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
		edgelq_logging_proto_v1_bucket_change_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BucketChange_Modified); i {
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
		edgelq_logging_proto_v1_bucket_change_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BucketChange_Current); i {
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
		edgelq_logging_proto_v1_bucket_change_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BucketChange_Removed); i {
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

	edgelq_logging_proto_v1_bucket_change_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*BucketChange_Added_)(nil),
		(*BucketChange_Modified_)(nil),
		(*BucketChange_Current_)(nil),
		(*BucketChange_Removed_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: edgelq_logging_proto_v1_bucket_change_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           edgelq_logging_proto_v1_bucket_change_proto_goTypes,
		DependencyIndexes: edgelq_logging_proto_v1_bucket_change_proto_depIdxs,
		MessageInfos:      edgelq_logging_proto_v1_bucket_change_proto_msgTypes,
	}.Build()
	edgelq_logging_proto_v1_bucket_change_proto = out.File
	edgelq_logging_proto_v1_bucket_change_proto_rawDesc = nil
	edgelq_logging_proto_v1_bucket_change_proto_goTypes = nil
	edgelq_logging_proto_v1_bucket_change_proto_depIdxs = nil
}
