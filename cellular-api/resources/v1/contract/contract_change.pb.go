// Code generated by protoc-gen-goten-go
// File: edgelq/cellular-api/proto/v1/contract_change.proto
// DO NOT EDIT!!!

package contract

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

// ContractChange is used by Watch notifications Responses to describe change of
// single Contract One of Added, Modified, Removed
type ContractChange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Contract change
	//
	// Types that are valid to be assigned to ChangeType:
	//	*ContractChange_Added_
	//	*ContractChange_Modified_
	//	*ContractChange_Current_
	//	*ContractChange_Removed_
	ChangeType isContractChange_ChangeType `protobuf_oneof:"change_type"`
}

func (m *ContractChange) Reset() {
	*m = ContractChange{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_cellular_api_proto_v1_contract_change_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *ContractChange) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*ContractChange) ProtoMessage() {}

func (m *ContractChange) ProtoReflect() preflect.Message {
	mi := &edgelq_cellular_api_proto_v1_contract_change_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*ContractChange) GotenMessage() {}

// Deprecated, Use ContractChange.ProtoReflect.Descriptor instead.
func (*ContractChange) Descriptor() ([]byte, []int) {
	return edgelq_cellular_api_proto_v1_contract_change_proto_rawDescGZIP(), []int{0}
}

func (m *ContractChange) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *ContractChange) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *ContractChange) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *ContractChange) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

type isContractChange_ChangeType interface {
	isContractChange_ChangeType()
}

type ContractChange_Added_ struct {
	// Added is returned when watched document is added, either created or
	// enters Query view
	Added *ContractChange_Added `protobuf:"bytes,1,opt,name=added,proto3,oneof"`
}
type ContractChange_Modified_ struct {
	// Modified is returned when watched document is modified
	Modified *ContractChange_Modified `protobuf:"bytes,2,opt,name=modified,proto3,oneof"`
}
type ContractChange_Current_ struct {
	// Current is returned in stateless watch when document enters query view or
	// is modified within.
	Current *ContractChange_Current `protobuf:"bytes,4,opt,name=current,proto3,oneof"`
}
type ContractChange_Removed_ struct {
	// Removed is returned when Contract is deleted or leaves Query view
	Removed *ContractChange_Removed `protobuf:"bytes,3,opt,name=removed,proto3,oneof"`
}

func (*ContractChange_Added_) isContractChange_ChangeType()    {}
func (*ContractChange_Modified_) isContractChange_ChangeType() {}
func (*ContractChange_Current_) isContractChange_ChangeType()  {}
func (*ContractChange_Removed_) isContractChange_ChangeType()  {}
func (m *ContractChange) GetChangeType() isContractChange_ChangeType {
	if m != nil {
		return m.ChangeType
	}
	return nil
}
func (m *ContractChange) GetAdded() *ContractChange_Added {
	if x, ok := m.GetChangeType().(*ContractChange_Added_); ok {
		return x.Added
	}
	return nil
}
func (m *ContractChange) GetModified() *ContractChange_Modified {
	if x, ok := m.GetChangeType().(*ContractChange_Modified_); ok {
		return x.Modified
	}
	return nil
}
func (m *ContractChange) GetCurrent() *ContractChange_Current {
	if x, ok := m.GetChangeType().(*ContractChange_Current_); ok {
		return x.Current
	}
	return nil
}
func (m *ContractChange) GetRemoved() *ContractChange_Removed {
	if x, ok := m.GetChangeType().(*ContractChange_Removed_); ok {
		return x.Removed
	}
	return nil
}
func (m *ContractChange) SetChangeType(ofv isContractChange_ChangeType) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "isContractChange_ChangeType", "ContractChange"))
	}
	m.ChangeType = ofv
}
func (m *ContractChange) SetAdded(fv *ContractChange_Added) {
	m.SetChangeType(&ContractChange_Added_{Added: fv})
}
func (m *ContractChange) SetModified(fv *ContractChange_Modified) {
	m.SetChangeType(&ContractChange_Modified_{Modified: fv})
}
func (m *ContractChange) SetCurrent(fv *ContractChange_Current) {
	m.SetChangeType(&ContractChange_Current_{Current: fv})
}
func (m *ContractChange) SetRemoved(fv *ContractChange_Removed) {
	m.SetChangeType(&ContractChange_Removed_{Removed: fv})
}

// Contract has been added to query view
type ContractChange_Added struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Contract      *Contract `protobuf:"bytes,1,opt,name=contract,proto3" json:"contract,omitempty"`
	// Integer describing index of added Contract in resulting query view.
	ViewIndex int32 `protobuf:"varint,2,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty"`
}

func (m *ContractChange_Added) Reset() {
	*m = ContractChange_Added{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_cellular_api_proto_v1_contract_change_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *ContractChange_Added) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*ContractChange_Added) ProtoMessage() {}

func (m *ContractChange_Added) ProtoReflect() preflect.Message {
	mi := &edgelq_cellular_api_proto_v1_contract_change_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*ContractChange_Added) GotenMessage() {}

// Deprecated, Use ContractChange_Added.ProtoReflect.Descriptor instead.
func (*ContractChange_Added) Descriptor() ([]byte, []int) {
	return edgelq_cellular_api_proto_v1_contract_change_proto_rawDescGZIP(), []int{0, 0}
}

func (m *ContractChange_Added) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *ContractChange_Added) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *ContractChange_Added) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *ContractChange_Added) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *ContractChange_Added) GetContract() *Contract {
	if m != nil {
		return m.Contract
	}
	return nil
}

func (m *ContractChange_Added) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *ContractChange_Added) SetContract(fv *Contract) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Contract", "ContractChange_Added"))
	}
	m.Contract = fv
}

func (m *ContractChange_Added) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "ContractChange_Added"))
	}
	m.ViewIndex = fv
}

// Contract changed some of it's fields - contains either full document or
// masked change
type ContractChange_Modified struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Name of modified Contract
	Name *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty"`
	// New version of Contract or masked difference, depending on mask_changes
	// instrumentation of issued [WatchContractRequest] or
	// [WatchContractsRequest]
	Contract *Contract `protobuf:"bytes,2,opt,name=contract,proto3" json:"contract,omitempty"`
	// Used when mask_changes is set, contains field paths of modified
	// properties.
	FieldMask *Contract_FieldMask `protobuf:"bytes,3,opt,customtype=Contract_FieldMask,name=field_mask,json=fieldMask,proto3" json:"field_mask,omitempty"`
	// Previous view index specifies previous position of modified Contract.
	// When modification doesn't affect sorted order, value will remain
	// identical to [view_index].
	PreviousViewIndex int32 `protobuf:"varint,4,opt,name=previous_view_index,json=previousViewIndex,proto3" json:"previous_view_index,omitempty"`
	// Integer specifying Contract new index in resulting query view.
	ViewIndex int32 `protobuf:"varint,5,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty"`
}

func (m *ContractChange_Modified) Reset() {
	*m = ContractChange_Modified{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_cellular_api_proto_v1_contract_change_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *ContractChange_Modified) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*ContractChange_Modified) ProtoMessage() {}

func (m *ContractChange_Modified) ProtoReflect() preflect.Message {
	mi := &edgelq_cellular_api_proto_v1_contract_change_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*ContractChange_Modified) GotenMessage() {}

// Deprecated, Use ContractChange_Modified.ProtoReflect.Descriptor instead.
func (*ContractChange_Modified) Descriptor() ([]byte, []int) {
	return edgelq_cellular_api_proto_v1_contract_change_proto_rawDescGZIP(), []int{0, 1}
}

func (m *ContractChange_Modified) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *ContractChange_Modified) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *ContractChange_Modified) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *ContractChange_Modified) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *ContractChange_Modified) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *ContractChange_Modified) GetContract() *Contract {
	if m != nil {
		return m.Contract
	}
	return nil
}

func (m *ContractChange_Modified) GetFieldMask() *Contract_FieldMask {
	if m != nil {
		return m.FieldMask
	}
	return nil
}

func (m *ContractChange_Modified) GetPreviousViewIndex() int32 {
	if m != nil {
		return m.PreviousViewIndex
	}
	return int32(0)
}

func (m *ContractChange_Modified) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *ContractChange_Modified) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "ContractChange_Modified"))
	}
	m.Name = fv
}

func (m *ContractChange_Modified) SetContract(fv *Contract) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Contract", "ContractChange_Modified"))
	}
	m.Contract = fv
}

func (m *ContractChange_Modified) SetFieldMask(fv *Contract_FieldMask) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "FieldMask", "ContractChange_Modified"))
	}
	m.FieldMask = fv
}

func (m *ContractChange_Modified) SetPreviousViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "PreviousViewIndex", "ContractChange_Modified"))
	}
	m.PreviousViewIndex = fv
}

func (m *ContractChange_Modified) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "ContractChange_Modified"))
	}
	m.ViewIndex = fv
}

// Contract has been added or modified in a query view. Version used for
// stateless watching
type ContractChange_Current struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Contract      *Contract `protobuf:"bytes,1,opt,name=contract,proto3" json:"contract,omitempty"`
}

func (m *ContractChange_Current) Reset() {
	*m = ContractChange_Current{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_cellular_api_proto_v1_contract_change_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *ContractChange_Current) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*ContractChange_Current) ProtoMessage() {}

func (m *ContractChange_Current) ProtoReflect() preflect.Message {
	mi := &edgelq_cellular_api_proto_v1_contract_change_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*ContractChange_Current) GotenMessage() {}

// Deprecated, Use ContractChange_Current.ProtoReflect.Descriptor instead.
func (*ContractChange_Current) Descriptor() ([]byte, []int) {
	return edgelq_cellular_api_proto_v1_contract_change_proto_rawDescGZIP(), []int{0, 2}
}

func (m *ContractChange_Current) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *ContractChange_Current) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *ContractChange_Current) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *ContractChange_Current) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *ContractChange_Current) GetContract() *Contract {
	if m != nil {
		return m.Contract
	}
	return nil
}

func (m *ContractChange_Current) SetContract(fv *Contract) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Contract", "ContractChange_Current"))
	}
	m.Contract = fv
}

// Removed is returned when Contract is deleted or leaves Query view
type ContractChange_Removed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Name          *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty"`
	// Integer specifying removed Contract index. Not populated in stateless
	// watch type.
	ViewIndex int32 `protobuf:"varint,2,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty"`
}

func (m *ContractChange_Removed) Reset() {
	*m = ContractChange_Removed{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_cellular_api_proto_v1_contract_change_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *ContractChange_Removed) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*ContractChange_Removed) ProtoMessage() {}

func (m *ContractChange_Removed) ProtoReflect() preflect.Message {
	mi := &edgelq_cellular_api_proto_v1_contract_change_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*ContractChange_Removed) GotenMessage() {}

// Deprecated, Use ContractChange_Removed.ProtoReflect.Descriptor instead.
func (*ContractChange_Removed) Descriptor() ([]byte, []int) {
	return edgelq_cellular_api_proto_v1_contract_change_proto_rawDescGZIP(), []int{0, 3}
}

func (m *ContractChange_Removed) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *ContractChange_Removed) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *ContractChange_Removed) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *ContractChange_Removed) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *ContractChange_Removed) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *ContractChange_Removed) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *ContractChange_Removed) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "ContractChange_Removed"))
	}
	m.Name = fv
}

func (m *ContractChange_Removed) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "ContractChange_Removed"))
	}
	m.ViewIndex = fv
}

var edgelq_cellular_api_proto_v1_contract_change_proto preflect.FileDescriptor

var edgelq_cellular_api_proto_v1_contract_change_proto_rawDesc = []byte{
	0x0a, 0x32, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x63, 0x65, 0x6c, 0x6c, 0x75, 0x6c, 0x61,
	0x72, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x6e, 0x74, 0x74, 0x2e, 0x63, 0x65, 0x6c, 0x6c, 0x75, 0x6c,
	0x61, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f,
	0x63, 0x65, 0x6c, 0x6c, 0x75, 0x6c, 0x61, 0x72, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd3, 0x06, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x41, 0x0a, 0x05, 0x61, 0x64, 0x64, 0x65, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x63, 0x65, 0x6c,
	0x6c, 0x75, 0x6c, 0x61, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x65,
	0x64, 0x48, 0x00, 0x52, 0x05, 0x61, 0x64, 0x64, 0x65, 0x64, 0x12, 0x4a, 0x0a, 0x08, 0x6d, 0x6f,
	0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x6e,
	0x74, 0x74, 0x2e, 0x63, 0x65, 0x6c, 0x6c, 0x75, 0x6c, 0x61, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x48, 0x00, 0x52, 0x08, 0x6d, 0x6f,
	0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x47, 0x0a, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x63, 0x65,
	0x6c, 0x6c, 0x75, 0x6c, 0x61, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12,
	0x47, 0x0a, 0x07, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2b, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x63, 0x65, 0x6c, 0x6c, 0x75, 0x6c, 0x61, 0x72, 0x5f,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x48, 0x00, 0x52,
	0x07, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x1a, 0x61, 0x0a, 0x05, 0x41, 0x64, 0x64, 0x65,
	0x64, 0x12, 0x39, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x63, 0x65, 0x6c, 0x6c, 0x75, 0x6c,
	0x61, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x76, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x1a, 0x87, 0x02, 0x0a, 0x08,
	0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x24, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x10, 0xb2, 0xda, 0x21, 0x0c, 0x0a, 0x0a, 0x0a, 0x08,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x39,
	0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x63, 0x65, 0x6c, 0x6c, 0x75, 0x6c, 0x61, 0x72, 0x5f,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52,
	0x08, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x4b, 0x0a, 0x0a, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x42, 0x10, 0xb2, 0xda, 0x21, 0x0c, 0x32,
	0x0a, 0x0a, 0x08, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x09, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x2e, 0x0a, 0x13, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f,
	0x75, 0x73, 0x5f, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x11, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x56, 0x69, 0x65,
	0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x76, 0x69, 0x65, 0x77,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x1a, 0x44, 0x0a, 0x07, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x12, 0x39, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x63, 0x65, 0x6c, 0x6c, 0x75, 0x6c, 0x61,
	0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x1a, 0x4e, 0x0a, 0x07, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x12, 0x24, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x10, 0xb2, 0xda, 0x21, 0x0c, 0x0a, 0x0a, 0x0a, 0x08, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x76, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x3a, 0x0e, 0x9a, 0xd9, 0x21,
	0x0a, 0x0a, 0x08, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x42, 0x0d, 0x0a, 0x0b, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x42, 0x7f, 0xe8, 0xde, 0x21, 0x00,
	0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x63, 0x65, 0x6c, 0x6c, 0x75, 0x6c,
	0x61, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x42, 0x13, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x00, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f,
	0x63, 0x65, 0x6c, 0x6c, 0x75, 0x6c, 0x61, 0x72, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x3b, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	edgelq_cellular_api_proto_v1_contract_change_proto_rawDescOnce sync.Once
	edgelq_cellular_api_proto_v1_contract_change_proto_rawDescData = edgelq_cellular_api_proto_v1_contract_change_proto_rawDesc
)

func edgelq_cellular_api_proto_v1_contract_change_proto_rawDescGZIP() []byte {
	edgelq_cellular_api_proto_v1_contract_change_proto_rawDescOnce.Do(func() {
		edgelq_cellular_api_proto_v1_contract_change_proto_rawDescData = protoimpl.X.CompressGZIP(edgelq_cellular_api_proto_v1_contract_change_proto_rawDescData)
	})
	return edgelq_cellular_api_proto_v1_contract_change_proto_rawDescData
}

var edgelq_cellular_api_proto_v1_contract_change_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var edgelq_cellular_api_proto_v1_contract_change_proto_goTypes = []interface{}{
	(*ContractChange)(nil),          // 0: ntt.cellular_api.v1.ContractChange
	(*ContractChange_Added)(nil),    // 1: ntt.cellular_api.v1.ContractChange.Added
	(*ContractChange_Modified)(nil), // 2: ntt.cellular_api.v1.ContractChange.Modified
	(*ContractChange_Current)(nil),  // 3: ntt.cellular_api.v1.ContractChange.Current
	(*ContractChange_Removed)(nil),  // 4: ntt.cellular_api.v1.ContractChange.Removed
	(*Contract)(nil),                // 5: ntt.cellular_api.v1.Contract
	(*Contract_FieldMask)(nil),      // 6: ntt.cellular_api.v1.Contract_FieldMask
}
var edgelq_cellular_api_proto_v1_contract_change_proto_depIdxs = []int32{
	1, // 0: ntt.cellular_api.v1.ContractChange.added:type_name -> ntt.cellular_api.v1.ContractChange.Added
	2, // 1: ntt.cellular_api.v1.ContractChange.modified:type_name -> ntt.cellular_api.v1.ContractChange.Modified
	3, // 2: ntt.cellular_api.v1.ContractChange.current:type_name -> ntt.cellular_api.v1.ContractChange.Current
	4, // 3: ntt.cellular_api.v1.ContractChange.removed:type_name -> ntt.cellular_api.v1.ContractChange.Removed
	5, // 4: ntt.cellular_api.v1.ContractChange.Added.contract:type_name -> ntt.cellular_api.v1.Contract
	5, // 5: ntt.cellular_api.v1.ContractChange.Modified.contract:type_name -> ntt.cellular_api.v1.Contract
	6, // 6: ntt.cellular_api.v1.ContractChange.Modified.field_mask:type_name -> ntt.cellular_api.v1.Contract_FieldMask
	5, // 7: ntt.cellular_api.v1.ContractChange.Current.contract:type_name -> ntt.cellular_api.v1.Contract
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { edgelq_cellular_api_proto_v1_contract_change_proto_init() }
func edgelq_cellular_api_proto_v1_contract_change_proto_init() {
	if edgelq_cellular_api_proto_v1_contract_change_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		edgelq_cellular_api_proto_v1_contract_change_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContractChange); i {
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
		edgelq_cellular_api_proto_v1_contract_change_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContractChange_Added); i {
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
		edgelq_cellular_api_proto_v1_contract_change_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContractChange_Modified); i {
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
		edgelq_cellular_api_proto_v1_contract_change_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContractChange_Current); i {
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
		edgelq_cellular_api_proto_v1_contract_change_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContractChange_Removed); i {
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

	edgelq_cellular_api_proto_v1_contract_change_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ContractChange_Added_)(nil),
		(*ContractChange_Modified_)(nil),
		(*ContractChange_Current_)(nil),
		(*ContractChange_Removed_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: edgelq_cellular_api_proto_v1_contract_change_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           edgelq_cellular_api_proto_v1_contract_change_proto_goTypes,
		DependencyIndexes: edgelq_cellular_api_proto_v1_contract_change_proto_depIdxs,
		MessageInfos:      edgelq_cellular_api_proto_v1_contract_change_proto_msgTypes,
	}.Build()
	edgelq_cellular_api_proto_v1_contract_change_proto = out.File
	edgelq_cellular_api_proto_v1_contract_change_proto_rawDesc = nil
	edgelq_cellular_api_proto_v1_contract_change_proto_goTypes = nil
	edgelq_cellular_api_proto_v1_contract_change_proto_depIdxs = nil
}
