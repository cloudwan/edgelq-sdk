// Code generated by protoc-gen-goten-go
// File: edgelq/secrets/proto/v1/crypto_key_change.proto
// DO NOT EDIT!!!

package crypto_key

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
	project "github.com/cloudwan/edgelq-sdk/secrets/resources/v1/project"
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
	_ = &project.Project{}
	_ = &fieldmaskpb.FieldMask{}
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// CryptoKeyChange is used by Watch notifications Responses to describe change
// of single CryptoKey One of Added, Modified, Removed
type CryptoKeyChange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// CryptoKey change
	//
	// Types that are valid to be assigned to ChangeType:
	//	*CryptoKeyChange_Added_
	//	*CryptoKeyChange_Modified_
	//	*CryptoKeyChange_Current_
	//	*CryptoKeyChange_Removed_
	ChangeType isCryptoKeyChange_ChangeType `protobuf_oneof:"change_type"`
}

func (m *CryptoKeyChange) Reset() {
	*m = CryptoKeyChange{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_secrets_proto_v1_crypto_key_change_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *CryptoKeyChange) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*CryptoKeyChange) ProtoMessage() {}

func (m *CryptoKeyChange) ProtoReflect() preflect.Message {
	mi := &edgelq_secrets_proto_v1_crypto_key_change_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*CryptoKeyChange) GotenMessage() {}

// Deprecated, Use CryptoKeyChange.ProtoReflect.Descriptor instead.
func (*CryptoKeyChange) Descriptor() ([]byte, []int) {
	return edgelq_secrets_proto_v1_crypto_key_change_proto_rawDescGZIP(), []int{0}
}

func (m *CryptoKeyChange) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *CryptoKeyChange) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *CryptoKeyChange) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *CryptoKeyChange) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

type isCryptoKeyChange_ChangeType interface {
	isCryptoKeyChange_ChangeType()
}

type CryptoKeyChange_Added_ struct {
	// Added is returned when watched document is added, either created or
	// enters Query view
	Added *CryptoKeyChange_Added `protobuf:"bytes,1,opt,name=added,proto3,oneof" firestore:"added"`
}
type CryptoKeyChange_Modified_ struct {
	// Modified is returned when watched document is modified
	Modified *CryptoKeyChange_Modified `protobuf:"bytes,2,opt,name=modified,proto3,oneof" firestore:"modified"`
}
type CryptoKeyChange_Current_ struct {
	// Current is returned in stateless watch when document enters query view or
	// is modified within.
	Current *CryptoKeyChange_Current `protobuf:"bytes,4,opt,name=current,proto3,oneof" firestore:"current"`
}
type CryptoKeyChange_Removed_ struct {
	// Removed is returned when CryptoKey is deleted or leaves Query view
	Removed *CryptoKeyChange_Removed `protobuf:"bytes,3,opt,name=removed,proto3,oneof" firestore:"removed"`
}

func (*CryptoKeyChange_Added_) isCryptoKeyChange_ChangeType()    {}
func (*CryptoKeyChange_Modified_) isCryptoKeyChange_ChangeType() {}
func (*CryptoKeyChange_Current_) isCryptoKeyChange_ChangeType()  {}
func (*CryptoKeyChange_Removed_) isCryptoKeyChange_ChangeType()  {}
func (m *CryptoKeyChange) GetChangeType() isCryptoKeyChange_ChangeType {
	if m != nil {
		return m.ChangeType
	}
	return nil
}
func (m *CryptoKeyChange) GetAdded() *CryptoKeyChange_Added {
	if x, ok := m.GetChangeType().(*CryptoKeyChange_Added_); ok {
		return x.Added
	}
	return nil
}
func (m *CryptoKeyChange) GetModified() *CryptoKeyChange_Modified {
	if x, ok := m.GetChangeType().(*CryptoKeyChange_Modified_); ok {
		return x.Modified
	}
	return nil
}
func (m *CryptoKeyChange) GetCurrent() *CryptoKeyChange_Current {
	if x, ok := m.GetChangeType().(*CryptoKeyChange_Current_); ok {
		return x.Current
	}
	return nil
}
func (m *CryptoKeyChange) GetRemoved() *CryptoKeyChange_Removed {
	if x, ok := m.GetChangeType().(*CryptoKeyChange_Removed_); ok {
		return x.Removed
	}
	return nil
}
func (m *CryptoKeyChange) SetChangeType(ofv isCryptoKeyChange_ChangeType) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "isCryptoKeyChange_ChangeType", "CryptoKeyChange"))
	}
	m.ChangeType = ofv
}
func (m *CryptoKeyChange) SetAdded(fv *CryptoKeyChange_Added) {
	m.SetChangeType(&CryptoKeyChange_Added_{Added: fv})
}
func (m *CryptoKeyChange) SetModified(fv *CryptoKeyChange_Modified) {
	m.SetChangeType(&CryptoKeyChange_Modified_{Modified: fv})
}
func (m *CryptoKeyChange) SetCurrent(fv *CryptoKeyChange_Current) {
	m.SetChangeType(&CryptoKeyChange_Current_{Current: fv})
}
func (m *CryptoKeyChange) SetRemoved(fv *CryptoKeyChange_Removed) {
	m.SetChangeType(&CryptoKeyChange_Removed_{Removed: fv})
}

// CryptoKey has been added to query view
type CryptoKeyChange_Added struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	CryptoKey     *CryptoKey `protobuf:"bytes,1,opt,name=crypto_key,json=cryptoKey,proto3" json:"crypto_key,omitempty" firestore:"cryptoKey"`
	// Integer describing index of added CryptoKey in resulting query view.
	ViewIndex int32 `protobuf:"varint,2,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty" firestore:"viewIndex"`
}

func (m *CryptoKeyChange_Added) Reset() {
	*m = CryptoKeyChange_Added{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_secrets_proto_v1_crypto_key_change_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *CryptoKeyChange_Added) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*CryptoKeyChange_Added) ProtoMessage() {}

func (m *CryptoKeyChange_Added) ProtoReflect() preflect.Message {
	mi := &edgelq_secrets_proto_v1_crypto_key_change_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*CryptoKeyChange_Added) GotenMessage() {}

// Deprecated, Use CryptoKeyChange_Added.ProtoReflect.Descriptor instead.
func (*CryptoKeyChange_Added) Descriptor() ([]byte, []int) {
	return edgelq_secrets_proto_v1_crypto_key_change_proto_rawDescGZIP(), []int{0, 0}
}

func (m *CryptoKeyChange_Added) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *CryptoKeyChange_Added) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *CryptoKeyChange_Added) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *CryptoKeyChange_Added) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *CryptoKeyChange_Added) GetCryptoKey() *CryptoKey {
	if m != nil {
		return m.CryptoKey
	}
	return nil
}

func (m *CryptoKeyChange_Added) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *CryptoKeyChange_Added) SetCryptoKey(fv *CryptoKey) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "CryptoKey", "CryptoKeyChange_Added"))
	}
	m.CryptoKey = fv
}

func (m *CryptoKeyChange_Added) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "CryptoKeyChange_Added"))
	}
	m.ViewIndex = fv
}

// CryptoKey changed some of it's fields - contains either full document or
// masked change
type CryptoKeyChange_Modified struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Name of modified CryptoKey
	Name *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// New version of CryptoKey or masked difference, depending on mask_changes
	// instrumentation of issued [WatchCryptoKeyRequest] or
	// [WatchCryptoKeysRequest]
	CryptoKey *CryptoKey `protobuf:"bytes,2,opt,name=crypto_key,json=cryptoKey,proto3" json:"crypto_key,omitempty" firestore:"cryptoKey"`
	// Used when mask_changes is set, contains field paths of modified
	// properties.
	FieldMask *CryptoKey_FieldMask `protobuf:"bytes,3,opt,customtype=CryptoKey_FieldMask,name=field_mask,json=fieldMask,proto3" json:"field_mask,omitempty" firestore:"fieldMask"`
	// Previous view index specifies previous position of modified CryptoKey.
	// When modification doesn't affect sorted order, value will remain
	// identical to [view_index].
	PreviousViewIndex int32 `protobuf:"varint,4,opt,name=previous_view_index,json=previousViewIndex,proto3" json:"previous_view_index,omitempty" firestore:"previousViewIndex"`
	// Integer specifying CryptoKey new index in resulting query view.
	ViewIndex int32 `protobuf:"varint,5,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty" firestore:"viewIndex"`
}

func (m *CryptoKeyChange_Modified) Reset() {
	*m = CryptoKeyChange_Modified{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_secrets_proto_v1_crypto_key_change_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *CryptoKeyChange_Modified) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*CryptoKeyChange_Modified) ProtoMessage() {}

func (m *CryptoKeyChange_Modified) ProtoReflect() preflect.Message {
	mi := &edgelq_secrets_proto_v1_crypto_key_change_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*CryptoKeyChange_Modified) GotenMessage() {}

// Deprecated, Use CryptoKeyChange_Modified.ProtoReflect.Descriptor instead.
func (*CryptoKeyChange_Modified) Descriptor() ([]byte, []int) {
	return edgelq_secrets_proto_v1_crypto_key_change_proto_rawDescGZIP(), []int{0, 1}
}

func (m *CryptoKeyChange_Modified) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *CryptoKeyChange_Modified) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *CryptoKeyChange_Modified) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *CryptoKeyChange_Modified) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *CryptoKeyChange_Modified) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *CryptoKeyChange_Modified) GetCryptoKey() *CryptoKey {
	if m != nil {
		return m.CryptoKey
	}
	return nil
}

func (m *CryptoKeyChange_Modified) GetFieldMask() *CryptoKey_FieldMask {
	if m != nil {
		return m.FieldMask
	}
	return nil
}

func (m *CryptoKeyChange_Modified) GetPreviousViewIndex() int32 {
	if m != nil {
		return m.PreviousViewIndex
	}
	return int32(0)
}

func (m *CryptoKeyChange_Modified) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *CryptoKeyChange_Modified) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "CryptoKeyChange_Modified"))
	}
	m.Name = fv
}

func (m *CryptoKeyChange_Modified) SetCryptoKey(fv *CryptoKey) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "CryptoKey", "CryptoKeyChange_Modified"))
	}
	m.CryptoKey = fv
}

func (m *CryptoKeyChange_Modified) SetFieldMask(fv *CryptoKey_FieldMask) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "FieldMask", "CryptoKeyChange_Modified"))
	}
	m.FieldMask = fv
}

func (m *CryptoKeyChange_Modified) SetPreviousViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "PreviousViewIndex", "CryptoKeyChange_Modified"))
	}
	m.PreviousViewIndex = fv
}

func (m *CryptoKeyChange_Modified) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "CryptoKeyChange_Modified"))
	}
	m.ViewIndex = fv
}

// CryptoKey has been added or modified in a query view. Version used for
// stateless watching
type CryptoKeyChange_Current struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	CryptoKey     *CryptoKey `protobuf:"bytes,1,opt,name=crypto_key,json=cryptoKey,proto3" json:"crypto_key,omitempty" firestore:"cryptoKey"`
}

func (m *CryptoKeyChange_Current) Reset() {
	*m = CryptoKeyChange_Current{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_secrets_proto_v1_crypto_key_change_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *CryptoKeyChange_Current) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*CryptoKeyChange_Current) ProtoMessage() {}

func (m *CryptoKeyChange_Current) ProtoReflect() preflect.Message {
	mi := &edgelq_secrets_proto_v1_crypto_key_change_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*CryptoKeyChange_Current) GotenMessage() {}

// Deprecated, Use CryptoKeyChange_Current.ProtoReflect.Descriptor instead.
func (*CryptoKeyChange_Current) Descriptor() ([]byte, []int) {
	return edgelq_secrets_proto_v1_crypto_key_change_proto_rawDescGZIP(), []int{0, 2}
}

func (m *CryptoKeyChange_Current) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *CryptoKeyChange_Current) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *CryptoKeyChange_Current) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *CryptoKeyChange_Current) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *CryptoKeyChange_Current) GetCryptoKey() *CryptoKey {
	if m != nil {
		return m.CryptoKey
	}
	return nil
}

func (m *CryptoKeyChange_Current) SetCryptoKey(fv *CryptoKey) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "CryptoKey", "CryptoKeyChange_Current"))
	}
	m.CryptoKey = fv
}

// Removed is returned when CryptoKey is deleted or leaves Query view
type CryptoKeyChange_Removed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Name          *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// Integer specifying removed CryptoKey index. Not populated in stateless
	// watch type.
	ViewIndex int32 `protobuf:"varint,2,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty" firestore:"viewIndex"`
}

func (m *CryptoKeyChange_Removed) Reset() {
	*m = CryptoKeyChange_Removed{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_secrets_proto_v1_crypto_key_change_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *CryptoKeyChange_Removed) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*CryptoKeyChange_Removed) ProtoMessage() {}

func (m *CryptoKeyChange_Removed) ProtoReflect() preflect.Message {
	mi := &edgelq_secrets_proto_v1_crypto_key_change_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*CryptoKeyChange_Removed) GotenMessage() {}

// Deprecated, Use CryptoKeyChange_Removed.ProtoReflect.Descriptor instead.
func (*CryptoKeyChange_Removed) Descriptor() ([]byte, []int) {
	return edgelq_secrets_proto_v1_crypto_key_change_proto_rawDescGZIP(), []int{0, 3}
}

func (m *CryptoKeyChange_Removed) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *CryptoKeyChange_Removed) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *CryptoKeyChange_Removed) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *CryptoKeyChange_Removed) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *CryptoKeyChange_Removed) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *CryptoKeyChange_Removed) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *CryptoKeyChange_Removed) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "CryptoKeyChange_Removed"))
	}
	m.Name = fv
}

func (m *CryptoKeyChange_Removed) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "CryptoKeyChange_Removed"))
	}
	m.ViewIndex = fv
}

var edgelq_secrets_proto_v1_crypto_key_change_proto preflect.FileDescriptor

var edgelq_secrets_proto_v1_crypto_key_change_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f,
	0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0e, 0x6e, 0x74, 0x74, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2e, 0x76,
	0x31, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d,
	0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67,
	0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x28, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x5f,
	0x6b, 0x65, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc5, 0x06, 0x0a, 0x0f, 0x43, 0x72,
	0x79, 0x70, 0x74, 0x6f, 0x4b, 0x65, 0x79, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x3d, 0x0a,
	0x05, 0x61, 0x64, 0x64, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6e,
	0x74, 0x74, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x79, 0x70, 0x74, 0x6f, 0x4b, 0x65, 0x79, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x41, 0x64,
	0x64, 0x65, 0x64, 0x48, 0x00, 0x52, 0x05, 0x61, 0x64, 0x64, 0x65, 0x64, 0x12, 0x46, 0x0a, 0x08,
	0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28,
	0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x4b, 0x65, 0x79, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e,
	0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x48, 0x00, 0x52, 0x08, 0x6d, 0x6f, 0x64, 0x69,
	0x66, 0x69, 0x65, 0x64, 0x12, 0x43, 0x0a, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x73, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x4b, 0x65, 0x79,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x48, 0x00,
	0x52, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x43, 0x0a, 0x07, 0x72, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x6e, 0x74, 0x74,
	0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x79, 0x70,
	0x74, 0x6f, 0x4b, 0x65, 0x79, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x64, 0x48, 0x00, 0x52, 0x07, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x1a, 0x60,
	0x0a, 0x05, 0x41, 0x64, 0x64, 0x65, 0x64, 0x12, 0x38, 0x0a, 0x0a, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x6f, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6e, 0x74,
	0x74, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x79,
	0x70, 0x74, 0x6f, 0x4b, 0x65, 0x79, 0x52, 0x09, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x4b, 0x65,
	0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x76, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x1a, 0x88, 0x02, 0x0a, 0x08, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x25, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x11, 0xb2, 0xda, 0x21,
	0x0d, 0x0a, 0x0b, 0x0a, 0x09, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x4b, 0x65, 0x79, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x38, 0x0a, 0x0a, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x5f, 0x6b,
	0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x73,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f,
	0x4b, 0x65, 0x79, 0x52, 0x09, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x4b, 0x65, 0x79, 0x12, 0x4c,
	0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x42, 0x11,
	0xb2, 0xda, 0x21, 0x0d, 0x32, 0x0b, 0x0a, 0x09, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x4b, 0x65,
	0x79, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x2e, 0x0a, 0x13,
	0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x5f, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x70, 0x72, 0x65, 0x76, 0x69,
	0x6f, 0x75, 0x73, 0x56, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1d, 0x0a, 0x0a,
	0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x76, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x1a, 0x43, 0x0a, 0x07, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x38, 0x0a, 0x0a, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f,
	0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6e, 0x74, 0x74,
	0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x79, 0x70,
	0x74, 0x6f, 0x4b, 0x65, 0x79, 0x52, 0x09, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x4b, 0x65, 0x79,
	0x1a, 0x4f, 0x0a, 0x07, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x12, 0x25, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x11, 0xb2, 0xda, 0x21, 0x0d, 0x0a,
	0x0b, 0x0a, 0x09, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x4b, 0x65, 0x79, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x76, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x3a, 0x0f, 0x9a, 0xd9, 0x21, 0x0b, 0x0a, 0x09, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x4b,
	0x65, 0x79, 0x42, 0x0d, 0x0a, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x42, 0x7a, 0xe8, 0xde, 0x21, 0x00, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x74, 0x74,
	0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x42, 0x14,
	0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x4b, 0x65, 0x79, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x00, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65,
	0x6c, 0x71, 0x2f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x5f, 0x6b,
	0x65, 0x79, 0x3b, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x5f, 0x6b, 0x65, 0x79, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	edgelq_secrets_proto_v1_crypto_key_change_proto_rawDescOnce sync.Once
	edgelq_secrets_proto_v1_crypto_key_change_proto_rawDescData = edgelq_secrets_proto_v1_crypto_key_change_proto_rawDesc
)

func edgelq_secrets_proto_v1_crypto_key_change_proto_rawDescGZIP() []byte {
	edgelq_secrets_proto_v1_crypto_key_change_proto_rawDescOnce.Do(func() {
		edgelq_secrets_proto_v1_crypto_key_change_proto_rawDescData = protoimpl.X.CompressGZIP(edgelq_secrets_proto_v1_crypto_key_change_proto_rawDescData)
	})
	return edgelq_secrets_proto_v1_crypto_key_change_proto_rawDescData
}

var edgelq_secrets_proto_v1_crypto_key_change_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var edgelq_secrets_proto_v1_crypto_key_change_proto_goTypes = []interface{}{
	(*CryptoKeyChange)(nil),          // 0: ntt.secrets.v1.CryptoKeyChange
	(*CryptoKeyChange_Added)(nil),    // 1: ntt.secrets.v1.CryptoKeyChange.Added
	(*CryptoKeyChange_Modified)(nil), // 2: ntt.secrets.v1.CryptoKeyChange.Modified
	(*CryptoKeyChange_Current)(nil),  // 3: ntt.secrets.v1.CryptoKeyChange.Current
	(*CryptoKeyChange_Removed)(nil),  // 4: ntt.secrets.v1.CryptoKeyChange.Removed
	(*CryptoKey)(nil),                // 5: ntt.secrets.v1.CryptoKey
	(*CryptoKey_FieldMask)(nil),      // 6: ntt.secrets.v1.CryptoKey_FieldMask
}
var edgelq_secrets_proto_v1_crypto_key_change_proto_depIdxs = []int32{
	1, // 0: ntt.secrets.v1.CryptoKeyChange.added:type_name -> ntt.secrets.v1.CryptoKeyChange.Added
	2, // 1: ntt.secrets.v1.CryptoKeyChange.modified:type_name -> ntt.secrets.v1.CryptoKeyChange.Modified
	3, // 2: ntt.secrets.v1.CryptoKeyChange.current:type_name -> ntt.secrets.v1.CryptoKeyChange.Current
	4, // 3: ntt.secrets.v1.CryptoKeyChange.removed:type_name -> ntt.secrets.v1.CryptoKeyChange.Removed
	5, // 4: ntt.secrets.v1.CryptoKeyChange.Added.crypto_key:type_name -> ntt.secrets.v1.CryptoKey
	5, // 5: ntt.secrets.v1.CryptoKeyChange.Modified.crypto_key:type_name -> ntt.secrets.v1.CryptoKey
	6, // 6: ntt.secrets.v1.CryptoKeyChange.Modified.field_mask:type_name -> ntt.secrets.v1.CryptoKey_FieldMask
	5, // 7: ntt.secrets.v1.CryptoKeyChange.Current.crypto_key:type_name -> ntt.secrets.v1.CryptoKey
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { edgelq_secrets_proto_v1_crypto_key_change_proto_init() }
func edgelq_secrets_proto_v1_crypto_key_change_proto_init() {
	if edgelq_secrets_proto_v1_crypto_key_change_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		edgelq_secrets_proto_v1_crypto_key_change_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CryptoKeyChange); i {
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
		edgelq_secrets_proto_v1_crypto_key_change_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CryptoKeyChange_Added); i {
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
		edgelq_secrets_proto_v1_crypto_key_change_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CryptoKeyChange_Modified); i {
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
		edgelq_secrets_proto_v1_crypto_key_change_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CryptoKeyChange_Current); i {
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
		edgelq_secrets_proto_v1_crypto_key_change_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CryptoKeyChange_Removed); i {
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

	edgelq_secrets_proto_v1_crypto_key_change_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*CryptoKeyChange_Added_)(nil),
		(*CryptoKeyChange_Modified_)(nil),
		(*CryptoKeyChange_Current_)(nil),
		(*CryptoKeyChange_Removed_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: edgelq_secrets_proto_v1_crypto_key_change_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           edgelq_secrets_proto_v1_crypto_key_change_proto_goTypes,
		DependencyIndexes: edgelq_secrets_proto_v1_crypto_key_change_proto_depIdxs,
		MessageInfos:      edgelq_secrets_proto_v1_crypto_key_change_proto_msgTypes,
	}.Build()
	edgelq_secrets_proto_v1_crypto_key_change_proto = out.File
	edgelq_secrets_proto_v1_crypto_key_change_proto_rawDesc = nil
	edgelq_secrets_proto_v1_crypto_key_change_proto_goTypes = nil
	edgelq_secrets_proto_v1_crypto_key_change_proto_depIdxs = nil
}
