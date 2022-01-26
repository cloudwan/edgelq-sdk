// Code generated by protoc-gen-goten-go
// File: edgelq/iam/proto/v1alpha2/attestation_domain_change.proto
// DO NOT EDIT!!!

package attestation_domain

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
	project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/project"
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

// AttestationDomainChange is used by Watch notifications Responses to describe
// change of single AttestationDomain One of Added, Modified, Removed
type AttestationDomainChange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// AttestationDomain change
	//
	// Types that are valid to be assigned to ChangeType:
	//	*AttestationDomainChange_Added_
	//	*AttestationDomainChange_Modified_
	//	*AttestationDomainChange_Current_
	//	*AttestationDomainChange_Removed_
	ChangeType isAttestationDomainChange_ChangeType `protobuf_oneof:"change_type"`
}

func (m *AttestationDomainChange) Reset() {
	*m = AttestationDomainChange{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_iam_proto_v1alpha2_attestation_domain_change_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *AttestationDomainChange) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*AttestationDomainChange) ProtoMessage() {}

func (m *AttestationDomainChange) ProtoReflect() preflect.Message {
	mi := &edgelq_iam_proto_v1alpha2_attestation_domain_change_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*AttestationDomainChange) GotenMessage() {}

// Deprecated, Use AttestationDomainChange.ProtoReflect.Descriptor instead.
func (*AttestationDomainChange) Descriptor() ([]byte, []int) {
	return edgelq_iam_proto_v1alpha2_attestation_domain_change_proto_rawDescGZIP(), []int{0}
}

func (m *AttestationDomainChange) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *AttestationDomainChange) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *AttestationDomainChange) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *AttestationDomainChange) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

type isAttestationDomainChange_ChangeType interface {
	isAttestationDomainChange_ChangeType()
}

type AttestationDomainChange_Added_ struct {
	// Added is returned when watched document is added, either created or
	// enters Query view
	Added *AttestationDomainChange_Added `protobuf:"bytes,1,opt,name=added,proto3,oneof" firestore:"added"`
}
type AttestationDomainChange_Modified_ struct {
	// Modified is returned when watched document is modified
	Modified *AttestationDomainChange_Modified `protobuf:"bytes,2,opt,name=modified,proto3,oneof" firestore:"modified"`
}
type AttestationDomainChange_Current_ struct {
	// Current is returned in stateless watch when document enters query view or
	// is modified within.
	Current *AttestationDomainChange_Current `protobuf:"bytes,4,opt,name=current,proto3,oneof" firestore:"current"`
}
type AttestationDomainChange_Removed_ struct {
	// Removed is returned when AttestationDomain is deleted or leaves Query
	// view
	Removed *AttestationDomainChange_Removed `protobuf:"bytes,3,opt,name=removed,proto3,oneof" firestore:"removed"`
}

func (*AttestationDomainChange_Added_) isAttestationDomainChange_ChangeType()    {}
func (*AttestationDomainChange_Modified_) isAttestationDomainChange_ChangeType() {}
func (*AttestationDomainChange_Current_) isAttestationDomainChange_ChangeType()  {}
func (*AttestationDomainChange_Removed_) isAttestationDomainChange_ChangeType()  {}
func (m *AttestationDomainChange) GetChangeType() isAttestationDomainChange_ChangeType {
	if m != nil {
		return m.ChangeType
	}
	return nil
}
func (m *AttestationDomainChange) GetAdded() *AttestationDomainChange_Added {
	if x, ok := m.GetChangeType().(*AttestationDomainChange_Added_); ok {
		return x.Added
	}
	return nil
}
func (m *AttestationDomainChange) GetModified() *AttestationDomainChange_Modified {
	if x, ok := m.GetChangeType().(*AttestationDomainChange_Modified_); ok {
		return x.Modified
	}
	return nil
}
func (m *AttestationDomainChange) GetCurrent() *AttestationDomainChange_Current {
	if x, ok := m.GetChangeType().(*AttestationDomainChange_Current_); ok {
		return x.Current
	}
	return nil
}
func (m *AttestationDomainChange) GetRemoved() *AttestationDomainChange_Removed {
	if x, ok := m.GetChangeType().(*AttestationDomainChange_Removed_); ok {
		return x.Removed
	}
	return nil
}
func (m *AttestationDomainChange) SetChangeType(ofv isAttestationDomainChange_ChangeType) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "isAttestationDomainChange_ChangeType", "AttestationDomainChange"))
	}
	m.ChangeType = ofv
}
func (m *AttestationDomainChange) SetAdded(fv *AttestationDomainChange_Added) {
	m.SetChangeType(&AttestationDomainChange_Added_{Added: fv})
}
func (m *AttestationDomainChange) SetModified(fv *AttestationDomainChange_Modified) {
	m.SetChangeType(&AttestationDomainChange_Modified_{Modified: fv})
}
func (m *AttestationDomainChange) SetCurrent(fv *AttestationDomainChange_Current) {
	m.SetChangeType(&AttestationDomainChange_Current_{Current: fv})
}
func (m *AttestationDomainChange) SetRemoved(fv *AttestationDomainChange_Removed) {
	m.SetChangeType(&AttestationDomainChange_Removed_{Removed: fv})
}

// AttestationDomain has been added to query view
type AttestationDomainChange_Added struct {
	state             protoimpl.MessageState
	sizeCache         protoimpl.SizeCache
	unknownFields     protoimpl.UnknownFields
	AttestationDomain *AttestationDomain `protobuf:"bytes,1,opt,name=attestation_domain,json=attestationDomain,proto3" json:"attestation_domain,omitempty" firestore:"attestationDomain"`
	// Integer describing index of added AttestationDomain in resulting query
	// view.
	ViewIndex int32 `protobuf:"varint,2,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty" firestore:"viewIndex"`
}

func (m *AttestationDomainChange_Added) Reset() {
	*m = AttestationDomainChange_Added{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_iam_proto_v1alpha2_attestation_domain_change_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *AttestationDomainChange_Added) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*AttestationDomainChange_Added) ProtoMessage() {}

func (m *AttestationDomainChange_Added) ProtoReflect() preflect.Message {
	mi := &edgelq_iam_proto_v1alpha2_attestation_domain_change_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*AttestationDomainChange_Added) GotenMessage() {}

// Deprecated, Use AttestationDomainChange_Added.ProtoReflect.Descriptor instead.
func (*AttestationDomainChange_Added) Descriptor() ([]byte, []int) {
	return edgelq_iam_proto_v1alpha2_attestation_domain_change_proto_rawDescGZIP(), []int{0, 0}
}

func (m *AttestationDomainChange_Added) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *AttestationDomainChange_Added) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *AttestationDomainChange_Added) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *AttestationDomainChange_Added) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *AttestationDomainChange_Added) GetAttestationDomain() *AttestationDomain {
	if m != nil {
		return m.AttestationDomain
	}
	return nil
}

func (m *AttestationDomainChange_Added) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *AttestationDomainChange_Added) SetAttestationDomain(fv *AttestationDomain) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "AttestationDomain", "AttestationDomainChange_Added"))
	}
	m.AttestationDomain = fv
}

func (m *AttestationDomainChange_Added) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "AttestationDomainChange_Added"))
	}
	m.ViewIndex = fv
}

// AttestationDomain changed some of it's fields - contains either full
// document or masked change
type AttestationDomainChange_Modified struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Name of modified AttestationDomain
	Name *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// New version of AttestationDomain or masked difference, depending on
	// mask_changes instrumentation of issued [WatchAttestationDomainRequest] or
	// [WatchAttestationDomainsRequest]
	AttestationDomain *AttestationDomain `protobuf:"bytes,2,opt,name=attestation_domain,json=attestationDomain,proto3" json:"attestation_domain,omitempty" firestore:"attestationDomain"`
	// Used when mask_changes is set, contains field paths of modified
	// properties.
	FieldMask *AttestationDomain_FieldMask `protobuf:"bytes,3,opt,customtype=AttestationDomain_FieldMask,name=field_mask,json=fieldMask,proto3" json:"field_mask,omitempty" firestore:"fieldMask"`
	// Previous view index specifies previous position of modified
	// AttestationDomain. When modification doesn't affect sorted order, value
	// will remain identical to [view_index].
	PreviousViewIndex int32 `protobuf:"varint,4,opt,name=previous_view_index,json=previousViewIndex,proto3" json:"previous_view_index,omitempty" firestore:"previousViewIndex"`
	// Integer specifying AttestationDomain new index in resulting query view.
	ViewIndex int32 `protobuf:"varint,5,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty" firestore:"viewIndex"`
}

func (m *AttestationDomainChange_Modified) Reset() {
	*m = AttestationDomainChange_Modified{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_iam_proto_v1alpha2_attestation_domain_change_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *AttestationDomainChange_Modified) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*AttestationDomainChange_Modified) ProtoMessage() {}

func (m *AttestationDomainChange_Modified) ProtoReflect() preflect.Message {
	mi := &edgelq_iam_proto_v1alpha2_attestation_domain_change_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*AttestationDomainChange_Modified) GotenMessage() {}

// Deprecated, Use AttestationDomainChange_Modified.ProtoReflect.Descriptor instead.
func (*AttestationDomainChange_Modified) Descriptor() ([]byte, []int) {
	return edgelq_iam_proto_v1alpha2_attestation_domain_change_proto_rawDescGZIP(), []int{0, 1}
}

func (m *AttestationDomainChange_Modified) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *AttestationDomainChange_Modified) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *AttestationDomainChange_Modified) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *AttestationDomainChange_Modified) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *AttestationDomainChange_Modified) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *AttestationDomainChange_Modified) GetAttestationDomain() *AttestationDomain {
	if m != nil {
		return m.AttestationDomain
	}
	return nil
}

func (m *AttestationDomainChange_Modified) GetFieldMask() *AttestationDomain_FieldMask {
	if m != nil {
		return m.FieldMask
	}
	return nil
}

func (m *AttestationDomainChange_Modified) GetPreviousViewIndex() int32 {
	if m != nil {
		return m.PreviousViewIndex
	}
	return int32(0)
}

func (m *AttestationDomainChange_Modified) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *AttestationDomainChange_Modified) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "AttestationDomainChange_Modified"))
	}
	m.Name = fv
}

func (m *AttestationDomainChange_Modified) SetAttestationDomain(fv *AttestationDomain) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "AttestationDomain", "AttestationDomainChange_Modified"))
	}
	m.AttestationDomain = fv
}

func (m *AttestationDomainChange_Modified) SetFieldMask(fv *AttestationDomain_FieldMask) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "FieldMask", "AttestationDomainChange_Modified"))
	}
	m.FieldMask = fv
}

func (m *AttestationDomainChange_Modified) SetPreviousViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "PreviousViewIndex", "AttestationDomainChange_Modified"))
	}
	m.PreviousViewIndex = fv
}

func (m *AttestationDomainChange_Modified) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "AttestationDomainChange_Modified"))
	}
	m.ViewIndex = fv
}

// AttestationDomain has been added or modified in a query view. Version used
// for stateless watching
type AttestationDomainChange_Current struct {
	state             protoimpl.MessageState
	sizeCache         protoimpl.SizeCache
	unknownFields     protoimpl.UnknownFields
	AttestationDomain *AttestationDomain `protobuf:"bytes,1,opt,name=attestation_domain,json=attestationDomain,proto3" json:"attestation_domain,omitempty" firestore:"attestationDomain"`
}

func (m *AttestationDomainChange_Current) Reset() {
	*m = AttestationDomainChange_Current{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_iam_proto_v1alpha2_attestation_domain_change_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *AttestationDomainChange_Current) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*AttestationDomainChange_Current) ProtoMessage() {}

func (m *AttestationDomainChange_Current) ProtoReflect() preflect.Message {
	mi := &edgelq_iam_proto_v1alpha2_attestation_domain_change_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*AttestationDomainChange_Current) GotenMessage() {}

// Deprecated, Use AttestationDomainChange_Current.ProtoReflect.Descriptor instead.
func (*AttestationDomainChange_Current) Descriptor() ([]byte, []int) {
	return edgelq_iam_proto_v1alpha2_attestation_domain_change_proto_rawDescGZIP(), []int{0, 2}
}

func (m *AttestationDomainChange_Current) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *AttestationDomainChange_Current) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *AttestationDomainChange_Current) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *AttestationDomainChange_Current) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *AttestationDomainChange_Current) GetAttestationDomain() *AttestationDomain {
	if m != nil {
		return m.AttestationDomain
	}
	return nil
}

func (m *AttestationDomainChange_Current) SetAttestationDomain(fv *AttestationDomain) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "AttestationDomain", "AttestationDomainChange_Current"))
	}
	m.AttestationDomain = fv
}

// Removed is returned when AttestationDomain is deleted or leaves Query view
type AttestationDomainChange_Removed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Name          *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// Integer specifying removed AttestationDomain index. Not populated in
	// stateless watch type.
	ViewIndex int32 `protobuf:"varint,2,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty" firestore:"viewIndex"`
}

func (m *AttestationDomainChange_Removed) Reset() {
	*m = AttestationDomainChange_Removed{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_iam_proto_v1alpha2_attestation_domain_change_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *AttestationDomainChange_Removed) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*AttestationDomainChange_Removed) ProtoMessage() {}

func (m *AttestationDomainChange_Removed) ProtoReflect() preflect.Message {
	mi := &edgelq_iam_proto_v1alpha2_attestation_domain_change_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*AttestationDomainChange_Removed) GotenMessage() {}

// Deprecated, Use AttestationDomainChange_Removed.ProtoReflect.Descriptor instead.
func (*AttestationDomainChange_Removed) Descriptor() ([]byte, []int) {
	return edgelq_iam_proto_v1alpha2_attestation_domain_change_proto_rawDescGZIP(), []int{0, 3}
}

func (m *AttestationDomainChange_Removed) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *AttestationDomainChange_Removed) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *AttestationDomainChange_Removed) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *AttestationDomainChange_Removed) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *AttestationDomainChange_Removed) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *AttestationDomainChange_Removed) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *AttestationDomainChange_Removed) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "AttestationDomainChange_Removed"))
	}
	m.Name = fv
}

func (m *AttestationDomainChange_Removed) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "AttestationDomainChange_Removed"))
	}
	m.ViewIndex = fv
}

var edgelq_iam_proto_v1alpha2_attestation_domain_change_proto preflect.FileDescriptor

var edgelq_iam_proto_v1alpha2_attestation_domain_change_proto_rawDesc = []byte{
	0x0a, 0x39, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x61, 0x74, 0x74, 0x65,
	0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x6e, 0x74, 0x74,
	0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x1a, 0x19, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65,
	0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f,
	0x74, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x32, 0x65, 0x64, 0x67,
	0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xe3, 0x07, 0x0a, 0x17, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x47, 0x0a, 0x05, 0x61,
	0x64, 0x64, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x6e, 0x74, 0x74,
	0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x41, 0x74,
	0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x65, 0x64, 0x48, 0x00, 0x52, 0x05, 0x61,
	0x64, 0x64, 0x65, 0x64, 0x12, 0x50, 0x0a, 0x08, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x48, 0x00, 0x52, 0x08, 0x6d, 0x6f,
	0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x4d, 0x0a, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61,
	0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x41, 0x74, 0x74, 0x65, 0x73,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x07, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x4d, 0x0a, 0x07, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x48, 0x00, 0x52, 0x07, 0x72, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x64, 0x1a, 0x7a, 0x0a, 0x05, 0x41, 0x64, 0x64, 0x65, 0x64, 0x12, 0x52, 0x0a,
	0x12, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6e, 0x74, 0x74, 0x2e,
	0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x41, 0x74, 0x74,
	0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x11,
	0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x76, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x1a, 0xb2, 0x02, 0x0a, 0x08, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x2d, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x19, 0xb2, 0xda, 0x21,
	0x15, 0x0a, 0x13, 0x0a, 0x11, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x52, 0x0a, 0x12,
	0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69,
	0x61, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x41, 0x74, 0x74, 0x65,
	0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x11, 0x61,
	0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x12, 0x54, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b,
	0x42, 0x19, 0xb2, 0xda, 0x21, 0x15, 0x32, 0x13, 0x0a, 0x11, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x09, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x2e, 0x0a, 0x13, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f,
	0x75, 0x73, 0x5f, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x11, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x56, 0x69, 0x65,
	0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x76, 0x69, 0x65, 0x77,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x1a, 0x5d, 0x0a, 0x07, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x12, 0x52, 0x0a, 0x12, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6e,
	0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e,
	0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x52, 0x11, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x1a, 0x57, 0x0a, 0x07, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x12,
	0x2d, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x19, 0xb2,
	0xda, 0x21, 0x15, 0x0a, 0x13, 0x0a, 0x11, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x76, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x3a, 0x17, 0x9a,
	0xd9, 0x21, 0x13, 0x0a, 0x11, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x42, 0x0d, 0x0a, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x42, 0x96, 0x01, 0xe8, 0xde, 0x21, 0x00, 0x0a, 0x17, 0x63, 0x6f,
	0x6d, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x32, 0x42, 0x1c, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x00, 0x5a, 0x57, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c,
	0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x3b, 0x61, 0x74, 0x74, 0x65,
	0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	edgelq_iam_proto_v1alpha2_attestation_domain_change_proto_rawDescOnce sync.Once
	edgelq_iam_proto_v1alpha2_attestation_domain_change_proto_rawDescData = edgelq_iam_proto_v1alpha2_attestation_domain_change_proto_rawDesc
)

func edgelq_iam_proto_v1alpha2_attestation_domain_change_proto_rawDescGZIP() []byte {
	edgelq_iam_proto_v1alpha2_attestation_domain_change_proto_rawDescOnce.Do(func() {
		edgelq_iam_proto_v1alpha2_attestation_domain_change_proto_rawDescData = protoimpl.X.CompressGZIP(edgelq_iam_proto_v1alpha2_attestation_domain_change_proto_rawDescData)
	})
	return edgelq_iam_proto_v1alpha2_attestation_domain_change_proto_rawDescData
}

var edgelq_iam_proto_v1alpha2_attestation_domain_change_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var edgelq_iam_proto_v1alpha2_attestation_domain_change_proto_goTypes = []interface{}{
	(*AttestationDomainChange)(nil),          // 0: ntt.iam.v1alpha2.AttestationDomainChange
	(*AttestationDomainChange_Added)(nil),    // 1: ntt.iam.v1alpha2.AttestationDomainChange.Added
	(*AttestationDomainChange_Modified)(nil), // 2: ntt.iam.v1alpha2.AttestationDomainChange.Modified
	(*AttestationDomainChange_Current)(nil),  // 3: ntt.iam.v1alpha2.AttestationDomainChange.Current
	(*AttestationDomainChange_Removed)(nil),  // 4: ntt.iam.v1alpha2.AttestationDomainChange.Removed
	(*AttestationDomain)(nil),                // 5: ntt.iam.v1alpha2.AttestationDomain
	(*AttestationDomain_FieldMask)(nil),      // 6: ntt.iam.v1alpha2.AttestationDomain_FieldMask
}
var edgelq_iam_proto_v1alpha2_attestation_domain_change_proto_depIdxs = []int32{
	1, // 0: ntt.iam.v1alpha2.AttestationDomainChange.added:type_name -> ntt.iam.v1alpha2.AttestationDomainChange.Added
	2, // 1: ntt.iam.v1alpha2.AttestationDomainChange.modified:type_name -> ntt.iam.v1alpha2.AttestationDomainChange.Modified
	3, // 2: ntt.iam.v1alpha2.AttestationDomainChange.current:type_name -> ntt.iam.v1alpha2.AttestationDomainChange.Current
	4, // 3: ntt.iam.v1alpha2.AttestationDomainChange.removed:type_name -> ntt.iam.v1alpha2.AttestationDomainChange.Removed
	5, // 4: ntt.iam.v1alpha2.AttestationDomainChange.Added.attestation_domain:type_name -> ntt.iam.v1alpha2.AttestationDomain
	5, // 5: ntt.iam.v1alpha2.AttestationDomainChange.Modified.attestation_domain:type_name -> ntt.iam.v1alpha2.AttestationDomain
	6, // 6: ntt.iam.v1alpha2.AttestationDomainChange.Modified.field_mask:type_name -> ntt.iam.v1alpha2.AttestationDomain_FieldMask
	5, // 7: ntt.iam.v1alpha2.AttestationDomainChange.Current.attestation_domain:type_name -> ntt.iam.v1alpha2.AttestationDomain
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { edgelq_iam_proto_v1alpha2_attestation_domain_change_proto_init() }
func edgelq_iam_proto_v1alpha2_attestation_domain_change_proto_init() {
	if edgelq_iam_proto_v1alpha2_attestation_domain_change_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		edgelq_iam_proto_v1alpha2_attestation_domain_change_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttestationDomainChange); i {
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
		edgelq_iam_proto_v1alpha2_attestation_domain_change_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttestationDomainChange_Added); i {
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
		edgelq_iam_proto_v1alpha2_attestation_domain_change_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttestationDomainChange_Modified); i {
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
		edgelq_iam_proto_v1alpha2_attestation_domain_change_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttestationDomainChange_Current); i {
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
		edgelq_iam_proto_v1alpha2_attestation_domain_change_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttestationDomainChange_Removed); i {
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

	edgelq_iam_proto_v1alpha2_attestation_domain_change_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*AttestationDomainChange_Added_)(nil),
		(*AttestationDomainChange_Modified_)(nil),
		(*AttestationDomainChange_Current_)(nil),
		(*AttestationDomainChange_Removed_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: edgelq_iam_proto_v1alpha2_attestation_domain_change_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           edgelq_iam_proto_v1alpha2_attestation_domain_change_proto_goTypes,
		DependencyIndexes: edgelq_iam_proto_v1alpha2_attestation_domain_change_proto_depIdxs,
		MessageInfos:      edgelq_iam_proto_v1alpha2_attestation_domain_change_proto_msgTypes,
	}.Build()
	edgelq_iam_proto_v1alpha2_attestation_domain_change_proto = out.File
	edgelq_iam_proto_v1alpha2_attestation_domain_change_proto_rawDesc = nil
	edgelq_iam_proto_v1alpha2_attestation_domain_change_proto_goTypes = nil
	edgelq_iam_proto_v1alpha2_attestation_domain_change_proto_depIdxs = nil
}
