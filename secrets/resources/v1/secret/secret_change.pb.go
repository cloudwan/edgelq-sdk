// Code generated by protoc-gen-goten-go
// File: edgelq/secrets/proto/v1/secret_change.proto
// DO NOT EDIT!!!

package secret

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

// SecretChange is used by Watch notifications Responses to describe change of
// single Secret One of Added, Modified, Removed
type SecretChange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Secret change
	//
	// Types that are valid to be assigned to ChangeType:
	//	*SecretChange_Added_
	//	*SecretChange_Modified_
	//	*SecretChange_Current_
	//	*SecretChange_Removed_
	ChangeType isSecretChange_ChangeType `protobuf_oneof:"change_type"`
}

func (m *SecretChange) Reset() {
	*m = SecretChange{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_secrets_proto_v1_secret_change_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *SecretChange) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*SecretChange) ProtoMessage() {}

func (m *SecretChange) ProtoReflect() preflect.Message {
	mi := &edgelq_secrets_proto_v1_secret_change_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*SecretChange) GotenMessage() {}

// Deprecated, Use SecretChange.ProtoReflect.Descriptor instead.
func (*SecretChange) Descriptor() ([]byte, []int) {
	return edgelq_secrets_proto_v1_secret_change_proto_rawDescGZIP(), []int{0}
}

func (m *SecretChange) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *SecretChange) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *SecretChange) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *SecretChange) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

type isSecretChange_ChangeType interface {
	isSecretChange_ChangeType()
}

type SecretChange_Added_ struct {
	// Added is returned when watched document is added, either created or
	// enters Query view
	Added *SecretChange_Added `protobuf:"bytes,1,opt,name=added,proto3,oneof" firestore:"added"`
}
type SecretChange_Modified_ struct {
	// Modified is returned when watched document is modified
	Modified *SecretChange_Modified `protobuf:"bytes,2,opt,name=modified,proto3,oneof" firestore:"modified"`
}
type SecretChange_Current_ struct {
	// Current is returned in stateless watch when document enters query view or
	// is modified within.
	Current *SecretChange_Current `protobuf:"bytes,4,opt,name=current,proto3,oneof" firestore:"current"`
}
type SecretChange_Removed_ struct {
	// Removed is returned when Secret is deleted or leaves Query view
	Removed *SecretChange_Removed `protobuf:"bytes,3,opt,name=removed,proto3,oneof" firestore:"removed"`
}

func (*SecretChange_Added_) isSecretChange_ChangeType()    {}
func (*SecretChange_Modified_) isSecretChange_ChangeType() {}
func (*SecretChange_Current_) isSecretChange_ChangeType()  {}
func (*SecretChange_Removed_) isSecretChange_ChangeType()  {}
func (m *SecretChange) GetChangeType() isSecretChange_ChangeType {
	if m != nil {
		return m.ChangeType
	}
	return nil
}
func (m *SecretChange) GetAdded() *SecretChange_Added {
	if x, ok := m.GetChangeType().(*SecretChange_Added_); ok {
		return x.Added
	}
	return nil
}
func (m *SecretChange) GetModified() *SecretChange_Modified {
	if x, ok := m.GetChangeType().(*SecretChange_Modified_); ok {
		return x.Modified
	}
	return nil
}
func (m *SecretChange) GetCurrent() *SecretChange_Current {
	if x, ok := m.GetChangeType().(*SecretChange_Current_); ok {
		return x.Current
	}
	return nil
}
func (m *SecretChange) GetRemoved() *SecretChange_Removed {
	if x, ok := m.GetChangeType().(*SecretChange_Removed_); ok {
		return x.Removed
	}
	return nil
}
func (m *SecretChange) SetChangeType(ofv isSecretChange_ChangeType) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "isSecretChange_ChangeType", "SecretChange"))
	}
	m.ChangeType = ofv
}
func (m *SecretChange) SetAdded(fv *SecretChange_Added) {
	m.SetChangeType(&SecretChange_Added_{Added: fv})
}
func (m *SecretChange) SetModified(fv *SecretChange_Modified) {
	m.SetChangeType(&SecretChange_Modified_{Modified: fv})
}
func (m *SecretChange) SetCurrent(fv *SecretChange_Current) {
	m.SetChangeType(&SecretChange_Current_{Current: fv})
}
func (m *SecretChange) SetRemoved(fv *SecretChange_Removed) {
	m.SetChangeType(&SecretChange_Removed_{Removed: fv})
}

// Secret has been added to query view
type SecretChange_Added struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Secret        *Secret `protobuf:"bytes,1,opt,name=secret,proto3" json:"secret,omitempty" firestore:"secret"`
	// Integer describing index of added Secret in resulting query view.
	ViewIndex int32 `protobuf:"varint,2,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty" firestore:"viewIndex"`
}

func (m *SecretChange_Added) Reset() {
	*m = SecretChange_Added{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_secrets_proto_v1_secret_change_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *SecretChange_Added) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*SecretChange_Added) ProtoMessage() {}

func (m *SecretChange_Added) ProtoReflect() preflect.Message {
	mi := &edgelq_secrets_proto_v1_secret_change_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*SecretChange_Added) GotenMessage() {}

// Deprecated, Use SecretChange_Added.ProtoReflect.Descriptor instead.
func (*SecretChange_Added) Descriptor() ([]byte, []int) {
	return edgelq_secrets_proto_v1_secret_change_proto_rawDescGZIP(), []int{0, 0}
}

func (m *SecretChange_Added) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *SecretChange_Added) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *SecretChange_Added) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *SecretChange_Added) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *SecretChange_Added) GetSecret() *Secret {
	if m != nil {
		return m.Secret
	}
	return nil
}

func (m *SecretChange_Added) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *SecretChange_Added) SetSecret(fv *Secret) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Secret", "SecretChange_Added"))
	}
	m.Secret = fv
}

func (m *SecretChange_Added) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "SecretChange_Added"))
	}
	m.ViewIndex = fv
}

// Secret changed some of it's fields - contains either full document or
// masked change
type SecretChange_Modified struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Name of modified Secret
	Name *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// New version of Secret or masked difference, depending on mask_changes
	// instrumentation of issued [WatchSecretRequest] or [WatchSecretsRequest]
	Secret *Secret `protobuf:"bytes,2,opt,name=secret,proto3" json:"secret,omitempty" firestore:"secret"`
	// Used when mask_changes is set, contains field paths of modified
	// properties.
	FieldMask *Secret_FieldMask `protobuf:"bytes,3,opt,customtype=Secret_FieldMask,name=field_mask,json=fieldMask,proto3" json:"field_mask,omitempty" firestore:"fieldMask"`
	// Previous view index specifies previous position of modified Secret.
	// When modification doesn't affect sorted order, value will remain
	// identical to [view_index].
	PreviousViewIndex int32 `protobuf:"varint,4,opt,name=previous_view_index,json=previousViewIndex,proto3" json:"previous_view_index,omitempty" firestore:"previousViewIndex"`
	// Integer specifying Secret new index in resulting query view.
	ViewIndex int32 `protobuf:"varint,5,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty" firestore:"viewIndex"`
}

func (m *SecretChange_Modified) Reset() {
	*m = SecretChange_Modified{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_secrets_proto_v1_secret_change_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *SecretChange_Modified) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*SecretChange_Modified) ProtoMessage() {}

func (m *SecretChange_Modified) ProtoReflect() preflect.Message {
	mi := &edgelq_secrets_proto_v1_secret_change_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*SecretChange_Modified) GotenMessage() {}

// Deprecated, Use SecretChange_Modified.ProtoReflect.Descriptor instead.
func (*SecretChange_Modified) Descriptor() ([]byte, []int) {
	return edgelq_secrets_proto_v1_secret_change_proto_rawDescGZIP(), []int{0, 1}
}

func (m *SecretChange_Modified) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *SecretChange_Modified) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *SecretChange_Modified) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *SecretChange_Modified) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *SecretChange_Modified) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *SecretChange_Modified) GetSecret() *Secret {
	if m != nil {
		return m.Secret
	}
	return nil
}

func (m *SecretChange_Modified) GetFieldMask() *Secret_FieldMask {
	if m != nil {
		return m.FieldMask
	}
	return nil
}

func (m *SecretChange_Modified) GetPreviousViewIndex() int32 {
	if m != nil {
		return m.PreviousViewIndex
	}
	return int32(0)
}

func (m *SecretChange_Modified) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *SecretChange_Modified) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "SecretChange_Modified"))
	}
	m.Name = fv
}

func (m *SecretChange_Modified) SetSecret(fv *Secret) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Secret", "SecretChange_Modified"))
	}
	m.Secret = fv
}

func (m *SecretChange_Modified) SetFieldMask(fv *Secret_FieldMask) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "FieldMask", "SecretChange_Modified"))
	}
	m.FieldMask = fv
}

func (m *SecretChange_Modified) SetPreviousViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "PreviousViewIndex", "SecretChange_Modified"))
	}
	m.PreviousViewIndex = fv
}

func (m *SecretChange_Modified) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "SecretChange_Modified"))
	}
	m.ViewIndex = fv
}

// Secret has been added or modified in a query view. Version used for
// stateless watching
type SecretChange_Current struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Secret        *Secret `protobuf:"bytes,1,opt,name=secret,proto3" json:"secret,omitempty" firestore:"secret"`
}

func (m *SecretChange_Current) Reset() {
	*m = SecretChange_Current{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_secrets_proto_v1_secret_change_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *SecretChange_Current) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*SecretChange_Current) ProtoMessage() {}

func (m *SecretChange_Current) ProtoReflect() preflect.Message {
	mi := &edgelq_secrets_proto_v1_secret_change_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*SecretChange_Current) GotenMessage() {}

// Deprecated, Use SecretChange_Current.ProtoReflect.Descriptor instead.
func (*SecretChange_Current) Descriptor() ([]byte, []int) {
	return edgelq_secrets_proto_v1_secret_change_proto_rawDescGZIP(), []int{0, 2}
}

func (m *SecretChange_Current) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *SecretChange_Current) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *SecretChange_Current) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *SecretChange_Current) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *SecretChange_Current) GetSecret() *Secret {
	if m != nil {
		return m.Secret
	}
	return nil
}

func (m *SecretChange_Current) SetSecret(fv *Secret) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Secret", "SecretChange_Current"))
	}
	m.Secret = fv
}

// Removed is returned when Secret is deleted or leaves Query view
type SecretChange_Removed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Name          *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// Integer specifying removed Secret index. Not populated in stateless watch
	// type.
	ViewIndex int32 `protobuf:"varint,2,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty" firestore:"viewIndex"`
}

func (m *SecretChange_Removed) Reset() {
	*m = SecretChange_Removed{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_secrets_proto_v1_secret_change_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *SecretChange_Removed) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*SecretChange_Removed) ProtoMessage() {}

func (m *SecretChange_Removed) ProtoReflect() preflect.Message {
	mi := &edgelq_secrets_proto_v1_secret_change_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*SecretChange_Removed) GotenMessage() {}

// Deprecated, Use SecretChange_Removed.ProtoReflect.Descriptor instead.
func (*SecretChange_Removed) Descriptor() ([]byte, []int) {
	return edgelq_secrets_proto_v1_secret_change_proto_rawDescGZIP(), []int{0, 3}
}

func (m *SecretChange_Removed) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *SecretChange_Removed) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *SecretChange_Removed) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *SecretChange_Removed) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *SecretChange_Removed) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *SecretChange_Removed) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *SecretChange_Removed) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "SecretChange_Removed"))
	}
	m.Name = fv
}

func (m *SecretChange_Removed) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "SecretChange_Removed"))
	}
	m.ViewIndex = fv
}

var edgelq_secrets_proto_v1_secret_change_proto preflect.FileDescriptor

var edgelq_secrets_proto_v1_secret_change_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x6e,
	0x74, 0x74, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x19, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65,
	0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f,
	0x74, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x65, 0x64, 0x67,
	0x65, 0x6c, 0x71, 0x2f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x8c, 0x06, 0x0a, 0x0c, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x12, 0x3a, 0x0a, 0x05, 0x61, 0x64, 0x64, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x22, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e,
	0x41, 0x64, 0x64, 0x65, 0x64, 0x48, 0x00, 0x52, 0x05, 0x61, 0x64, 0x64, 0x65, 0x64, 0x12, 0x43,
	0x0a, 0x08, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x25, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x4d,
	0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x48, 0x00, 0x52, 0x08, 0x6d, 0x6f, 0x64, 0x69, 0x66,
	0x69, 0x65, 0x64, 0x12, 0x40, 0x0a, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x07, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x40, 0x0a, 0x07, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x73, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x48, 0x00, 0x52, 0x07,
	0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x1a, 0x56, 0x0a, 0x05, 0x41, 0x64, 0x64, 0x65, 0x64,
	0x12, 0x2e, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x76, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x1a,
	0xf8, 0x01, 0x0a, 0x08, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x22, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0e, 0xb2, 0xda, 0x21, 0x0a,
	0x0a, 0x08, 0x0a, 0x06, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x2e, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x12, 0x49, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b,
	0x42, 0x0e, 0xb2, 0xda, 0x21, 0x0a, 0x32, 0x08, 0x0a, 0x06, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x2e, 0x0a, 0x13, 0x70,
	0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x5f, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f,
	0x75, 0x73, 0x56, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1d, 0x0a, 0x0a, 0x76,
	0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x76, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x1a, 0x39, 0x0a, 0x07, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x2e, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x73, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x06, 0x73,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x1a, 0x4c, 0x0a, 0x07, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64,
	0x12, 0x22, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0e,
	0xb2, 0xda, 0x21, 0x0a, 0x0a, 0x08, 0x0a, 0x06, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x76, 0x69, 0x65, 0x77, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x3a, 0x0c, 0x9a, 0xd9, 0x21, 0x08, 0x0a, 0x06, 0x53, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x42, 0x0d, 0x0a, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x42, 0x6f, 0xe8, 0xde, 0x21, 0x00, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x74, 0x74, 0x2e,
	0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x42, 0x11, 0x53,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x00, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x73,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x3b, 0x73, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	edgelq_secrets_proto_v1_secret_change_proto_rawDescOnce sync.Once
	edgelq_secrets_proto_v1_secret_change_proto_rawDescData = edgelq_secrets_proto_v1_secret_change_proto_rawDesc
)

func edgelq_secrets_proto_v1_secret_change_proto_rawDescGZIP() []byte {
	edgelq_secrets_proto_v1_secret_change_proto_rawDescOnce.Do(func() {
		edgelq_secrets_proto_v1_secret_change_proto_rawDescData = protoimpl.X.CompressGZIP(edgelq_secrets_proto_v1_secret_change_proto_rawDescData)
	})
	return edgelq_secrets_proto_v1_secret_change_proto_rawDescData
}

var edgelq_secrets_proto_v1_secret_change_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var edgelq_secrets_proto_v1_secret_change_proto_goTypes = []interface{}{
	(*SecretChange)(nil),          // 0: ntt.secrets.v1.SecretChange
	(*SecretChange_Added)(nil),    // 1: ntt.secrets.v1.SecretChange.Added
	(*SecretChange_Modified)(nil), // 2: ntt.secrets.v1.SecretChange.Modified
	(*SecretChange_Current)(nil),  // 3: ntt.secrets.v1.SecretChange.Current
	(*SecretChange_Removed)(nil),  // 4: ntt.secrets.v1.SecretChange.Removed
	(*Secret)(nil),                // 5: ntt.secrets.v1.Secret
	(*Secret_FieldMask)(nil),      // 6: ntt.secrets.v1.Secret_FieldMask
}
var edgelq_secrets_proto_v1_secret_change_proto_depIdxs = []int32{
	1, // 0: ntt.secrets.v1.SecretChange.added:type_name -> ntt.secrets.v1.SecretChange.Added
	2, // 1: ntt.secrets.v1.SecretChange.modified:type_name -> ntt.secrets.v1.SecretChange.Modified
	3, // 2: ntt.secrets.v1.SecretChange.current:type_name -> ntt.secrets.v1.SecretChange.Current
	4, // 3: ntt.secrets.v1.SecretChange.removed:type_name -> ntt.secrets.v1.SecretChange.Removed
	5, // 4: ntt.secrets.v1.SecretChange.Added.secret:type_name -> ntt.secrets.v1.Secret
	5, // 5: ntt.secrets.v1.SecretChange.Modified.secret:type_name -> ntt.secrets.v1.Secret
	6, // 6: ntt.secrets.v1.SecretChange.Modified.field_mask:type_name -> ntt.secrets.v1.Secret_FieldMask
	5, // 7: ntt.secrets.v1.SecretChange.Current.secret:type_name -> ntt.secrets.v1.Secret
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { edgelq_secrets_proto_v1_secret_change_proto_init() }
func edgelq_secrets_proto_v1_secret_change_proto_init() {
	if edgelq_secrets_proto_v1_secret_change_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		edgelq_secrets_proto_v1_secret_change_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecretChange); i {
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
		edgelq_secrets_proto_v1_secret_change_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecretChange_Added); i {
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
		edgelq_secrets_proto_v1_secret_change_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecretChange_Modified); i {
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
		edgelq_secrets_proto_v1_secret_change_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecretChange_Current); i {
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
		edgelq_secrets_proto_v1_secret_change_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecretChange_Removed); i {
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

	edgelq_secrets_proto_v1_secret_change_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*SecretChange_Added_)(nil),
		(*SecretChange_Modified_)(nil),
		(*SecretChange_Current_)(nil),
		(*SecretChange_Removed_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: edgelq_secrets_proto_v1_secret_change_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           edgelq_secrets_proto_v1_secret_change_proto_goTypes,
		DependencyIndexes: edgelq_secrets_proto_v1_secret_change_proto_depIdxs,
		MessageInfos:      edgelq_secrets_proto_v1_secret_change_proto_msgTypes,
	}.Build()
	edgelq_secrets_proto_v1_secret_change_proto = out.File
	edgelq_secrets_proto_v1_secret_change_proto_rawDesc = nil
	edgelq_secrets_proto_v1_secret_change_proto_goTypes = nil
	edgelq_secrets_proto_v1_secret_change_proto_depIdxs = nil
}