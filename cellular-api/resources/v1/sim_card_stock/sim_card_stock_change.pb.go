// Code generated by protoc-gen-goten-go
// File: edgelq/cellular-api/proto/v1/sim_card_stock_change.proto
// DO NOT EDIT!!!

package sim_card_stock

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

// SimCardStockChange is used by Watch notifications Responses to describe
// change of single SimCardStock One of Added, Modified, Removed
type SimCardStockChange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// SimCardStock change
	//
	// Types that are valid to be assigned to ChangeType:
	//	*SimCardStockChange_Added_
	//	*SimCardStockChange_Modified_
	//	*SimCardStockChange_Current_
	//	*SimCardStockChange_Removed_
	ChangeType isSimCardStockChange_ChangeType `protobuf_oneof:"change_type"`
}

func (m *SimCardStockChange) Reset() {
	*m = SimCardStockChange{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_cellular_api_proto_v1_sim_card_stock_change_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *SimCardStockChange) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*SimCardStockChange) ProtoMessage() {}

func (m *SimCardStockChange) ProtoReflect() preflect.Message {
	mi := &edgelq_cellular_api_proto_v1_sim_card_stock_change_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*SimCardStockChange) GotenMessage() {}

// Deprecated, Use SimCardStockChange.ProtoReflect.Descriptor instead.
func (*SimCardStockChange) Descriptor() ([]byte, []int) {
	return edgelq_cellular_api_proto_v1_sim_card_stock_change_proto_rawDescGZIP(), []int{0}
}

func (m *SimCardStockChange) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *SimCardStockChange) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *SimCardStockChange) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *SimCardStockChange) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

type isSimCardStockChange_ChangeType interface {
	isSimCardStockChange_ChangeType()
}

type SimCardStockChange_Added_ struct {
	// Added is returned when watched document is added, either created or
	// enters Query view
	Added *SimCardStockChange_Added `protobuf:"bytes,1,opt,name=added,proto3,oneof"`
}
type SimCardStockChange_Modified_ struct {
	// Modified is returned when watched document is modified
	Modified *SimCardStockChange_Modified `protobuf:"bytes,2,opt,name=modified,proto3,oneof"`
}
type SimCardStockChange_Current_ struct {
	// Current is returned in stateless watch when document enters query view or
	// is modified within.
	Current *SimCardStockChange_Current `protobuf:"bytes,4,opt,name=current,proto3,oneof"`
}
type SimCardStockChange_Removed_ struct {
	// Removed is returned when SimCardStock is deleted or leaves Query view
	Removed *SimCardStockChange_Removed `protobuf:"bytes,3,opt,name=removed,proto3,oneof"`
}

func (*SimCardStockChange_Added_) isSimCardStockChange_ChangeType()    {}
func (*SimCardStockChange_Modified_) isSimCardStockChange_ChangeType() {}
func (*SimCardStockChange_Current_) isSimCardStockChange_ChangeType()  {}
func (*SimCardStockChange_Removed_) isSimCardStockChange_ChangeType()  {}
func (m *SimCardStockChange) GetChangeType() isSimCardStockChange_ChangeType {
	if m != nil {
		return m.ChangeType
	}
	return nil
}
func (m *SimCardStockChange) GetAdded() *SimCardStockChange_Added {
	if x, ok := m.GetChangeType().(*SimCardStockChange_Added_); ok {
		return x.Added
	}
	return nil
}
func (m *SimCardStockChange) GetModified() *SimCardStockChange_Modified {
	if x, ok := m.GetChangeType().(*SimCardStockChange_Modified_); ok {
		return x.Modified
	}
	return nil
}
func (m *SimCardStockChange) GetCurrent() *SimCardStockChange_Current {
	if x, ok := m.GetChangeType().(*SimCardStockChange_Current_); ok {
		return x.Current
	}
	return nil
}
func (m *SimCardStockChange) GetRemoved() *SimCardStockChange_Removed {
	if x, ok := m.GetChangeType().(*SimCardStockChange_Removed_); ok {
		return x.Removed
	}
	return nil
}
func (m *SimCardStockChange) SetChangeType(ofv isSimCardStockChange_ChangeType) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "isSimCardStockChange_ChangeType", "SimCardStockChange"))
	}
	m.ChangeType = ofv
}
func (m *SimCardStockChange) SetAdded(fv *SimCardStockChange_Added) {
	m.SetChangeType(&SimCardStockChange_Added_{Added: fv})
}
func (m *SimCardStockChange) SetModified(fv *SimCardStockChange_Modified) {
	m.SetChangeType(&SimCardStockChange_Modified_{Modified: fv})
}
func (m *SimCardStockChange) SetCurrent(fv *SimCardStockChange_Current) {
	m.SetChangeType(&SimCardStockChange_Current_{Current: fv})
}
func (m *SimCardStockChange) SetRemoved(fv *SimCardStockChange_Removed) {
	m.SetChangeType(&SimCardStockChange_Removed_{Removed: fv})
}

// SimCardStock has been added to query view
type SimCardStockChange_Added struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	SimCardStock  *SimCardStock `protobuf:"bytes,1,opt,name=sim_card_stock,json=simCardStock,proto3" json:"sim_card_stock,omitempty"`
	// Integer describing index of added SimCardStock in resulting query view.
	ViewIndex int32 `protobuf:"varint,2,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty"`
}

func (m *SimCardStockChange_Added) Reset() {
	*m = SimCardStockChange_Added{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_cellular_api_proto_v1_sim_card_stock_change_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *SimCardStockChange_Added) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*SimCardStockChange_Added) ProtoMessage() {}

func (m *SimCardStockChange_Added) ProtoReflect() preflect.Message {
	mi := &edgelq_cellular_api_proto_v1_sim_card_stock_change_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*SimCardStockChange_Added) GotenMessage() {}

// Deprecated, Use SimCardStockChange_Added.ProtoReflect.Descriptor instead.
func (*SimCardStockChange_Added) Descriptor() ([]byte, []int) {
	return edgelq_cellular_api_proto_v1_sim_card_stock_change_proto_rawDescGZIP(), []int{0, 0}
}

func (m *SimCardStockChange_Added) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *SimCardStockChange_Added) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *SimCardStockChange_Added) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *SimCardStockChange_Added) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *SimCardStockChange_Added) GetSimCardStock() *SimCardStock {
	if m != nil {
		return m.SimCardStock
	}
	return nil
}

func (m *SimCardStockChange_Added) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *SimCardStockChange_Added) SetSimCardStock(fv *SimCardStock) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "SimCardStock", "SimCardStockChange_Added"))
	}
	m.SimCardStock = fv
}

func (m *SimCardStockChange_Added) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "SimCardStockChange_Added"))
	}
	m.ViewIndex = fv
}

// SimCardStock changed some of it's fields - contains either full document or
// masked change
type SimCardStockChange_Modified struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Name of modified SimCardStock
	Name *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty"`
	// New version of SimCardStock or masked difference, depending on
	// mask_changes instrumentation of issued [WatchSimCardStockRequest] or
	// [WatchSimCardStocksRequest]
	SimCardStock *SimCardStock `protobuf:"bytes,2,opt,name=sim_card_stock,json=simCardStock,proto3" json:"sim_card_stock,omitempty"`
	// Used when mask_changes is set, contains field paths of modified
	// properties.
	FieldMask *SimCardStock_FieldMask `protobuf:"bytes,3,opt,customtype=SimCardStock_FieldMask,name=field_mask,json=fieldMask,proto3" json:"field_mask,omitempty"`
	// Previous view index specifies previous position of modified SimCardStock.
	// When modification doesn't affect sorted order, value will remain
	// identical to [view_index].
	PreviousViewIndex int32 `protobuf:"varint,4,opt,name=previous_view_index,json=previousViewIndex,proto3" json:"previous_view_index,omitempty"`
	// Integer specifying SimCardStock new index in resulting query view.
	ViewIndex int32 `protobuf:"varint,5,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty"`
}

func (m *SimCardStockChange_Modified) Reset() {
	*m = SimCardStockChange_Modified{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_cellular_api_proto_v1_sim_card_stock_change_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *SimCardStockChange_Modified) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*SimCardStockChange_Modified) ProtoMessage() {}

func (m *SimCardStockChange_Modified) ProtoReflect() preflect.Message {
	mi := &edgelq_cellular_api_proto_v1_sim_card_stock_change_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*SimCardStockChange_Modified) GotenMessage() {}

// Deprecated, Use SimCardStockChange_Modified.ProtoReflect.Descriptor instead.
func (*SimCardStockChange_Modified) Descriptor() ([]byte, []int) {
	return edgelq_cellular_api_proto_v1_sim_card_stock_change_proto_rawDescGZIP(), []int{0, 1}
}

func (m *SimCardStockChange_Modified) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *SimCardStockChange_Modified) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *SimCardStockChange_Modified) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *SimCardStockChange_Modified) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *SimCardStockChange_Modified) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *SimCardStockChange_Modified) GetSimCardStock() *SimCardStock {
	if m != nil {
		return m.SimCardStock
	}
	return nil
}

func (m *SimCardStockChange_Modified) GetFieldMask() *SimCardStock_FieldMask {
	if m != nil {
		return m.FieldMask
	}
	return nil
}

func (m *SimCardStockChange_Modified) GetPreviousViewIndex() int32 {
	if m != nil {
		return m.PreviousViewIndex
	}
	return int32(0)
}

func (m *SimCardStockChange_Modified) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *SimCardStockChange_Modified) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "SimCardStockChange_Modified"))
	}
	m.Name = fv
}

func (m *SimCardStockChange_Modified) SetSimCardStock(fv *SimCardStock) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "SimCardStock", "SimCardStockChange_Modified"))
	}
	m.SimCardStock = fv
}

func (m *SimCardStockChange_Modified) SetFieldMask(fv *SimCardStock_FieldMask) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "FieldMask", "SimCardStockChange_Modified"))
	}
	m.FieldMask = fv
}

func (m *SimCardStockChange_Modified) SetPreviousViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "PreviousViewIndex", "SimCardStockChange_Modified"))
	}
	m.PreviousViewIndex = fv
}

func (m *SimCardStockChange_Modified) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "SimCardStockChange_Modified"))
	}
	m.ViewIndex = fv
}

// SimCardStock has been added or modified in a query view. Version used for
// stateless watching
type SimCardStockChange_Current struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	SimCardStock  *SimCardStock `protobuf:"bytes,1,opt,name=sim_card_stock,json=simCardStock,proto3" json:"sim_card_stock,omitempty"`
}

func (m *SimCardStockChange_Current) Reset() {
	*m = SimCardStockChange_Current{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_cellular_api_proto_v1_sim_card_stock_change_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *SimCardStockChange_Current) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*SimCardStockChange_Current) ProtoMessage() {}

func (m *SimCardStockChange_Current) ProtoReflect() preflect.Message {
	mi := &edgelq_cellular_api_proto_v1_sim_card_stock_change_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*SimCardStockChange_Current) GotenMessage() {}

// Deprecated, Use SimCardStockChange_Current.ProtoReflect.Descriptor instead.
func (*SimCardStockChange_Current) Descriptor() ([]byte, []int) {
	return edgelq_cellular_api_proto_v1_sim_card_stock_change_proto_rawDescGZIP(), []int{0, 2}
}

func (m *SimCardStockChange_Current) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *SimCardStockChange_Current) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *SimCardStockChange_Current) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *SimCardStockChange_Current) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *SimCardStockChange_Current) GetSimCardStock() *SimCardStock {
	if m != nil {
		return m.SimCardStock
	}
	return nil
}

func (m *SimCardStockChange_Current) SetSimCardStock(fv *SimCardStock) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "SimCardStock", "SimCardStockChange_Current"))
	}
	m.SimCardStock = fv
}

// Removed is returned when SimCardStock is deleted or leaves Query view
type SimCardStockChange_Removed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Name          *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty"`
	// Integer specifying removed SimCardStock index. Not populated in stateless
	// watch type.
	ViewIndex int32 `protobuf:"varint,2,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty"`
}

func (m *SimCardStockChange_Removed) Reset() {
	*m = SimCardStockChange_Removed{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_cellular_api_proto_v1_sim_card_stock_change_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *SimCardStockChange_Removed) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*SimCardStockChange_Removed) ProtoMessage() {}

func (m *SimCardStockChange_Removed) ProtoReflect() preflect.Message {
	mi := &edgelq_cellular_api_proto_v1_sim_card_stock_change_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*SimCardStockChange_Removed) GotenMessage() {}

// Deprecated, Use SimCardStockChange_Removed.ProtoReflect.Descriptor instead.
func (*SimCardStockChange_Removed) Descriptor() ([]byte, []int) {
	return edgelq_cellular_api_proto_v1_sim_card_stock_change_proto_rawDescGZIP(), []int{0, 3}
}

func (m *SimCardStockChange_Removed) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *SimCardStockChange_Removed) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *SimCardStockChange_Removed) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *SimCardStockChange_Removed) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *SimCardStockChange_Removed) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *SimCardStockChange_Removed) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *SimCardStockChange_Removed) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "SimCardStockChange_Removed"))
	}
	m.Name = fv
}

func (m *SimCardStockChange_Removed) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "SimCardStockChange_Removed"))
	}
	m.ViewIndex = fv
}

var edgelq_cellular_api_proto_v1_sim_card_stock_change_proto preflect.FileDescriptor

var edgelq_cellular_api_proto_v1_sim_card_stock_change_proto_rawDesc = []byte{
	0x0a, 0x38, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x63, 0x65, 0x6c, 0x6c, 0x75, 0x6c, 0x61,
	0x72, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x73,
	0x69, 0x6d, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x5f, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x6e, 0x74, 0x74, 0x2e,
	0x63, 0x65, 0x6c, 0x6c, 0x75, 0x6c, 0x61, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a,
	0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f,
	0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x74,
	0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x31, 0x65,
	0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x63, 0x65, 0x6c, 0x6c, 0x75, 0x6c, 0x61, 0x72, 0x2d, 0x61,
	0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x69, 0x6d, 0x5f,
	0x63, 0x61, 0x72, 0x64, 0x5f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xa1, 0x07, 0x0a, 0x12, 0x53, 0x69, 0x6d, 0x43, 0x61, 0x72, 0x64, 0x53, 0x74, 0x6f, 0x63,
	0x6b, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x45, 0x0a, 0x05, 0x61, 0x64, 0x64, 0x65, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x63, 0x65, 0x6c,
	0x6c, 0x75, 0x6c, 0x61, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x6d,
	0x43, 0x61, 0x72, 0x64, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e,
	0x41, 0x64, 0x64, 0x65, 0x64, 0x48, 0x00, 0x52, 0x05, 0x61, 0x64, 0x64, 0x65, 0x64, 0x12, 0x4e,
	0x0a, 0x08, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x30, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x63, 0x65, 0x6c, 0x6c, 0x75, 0x6c, 0x61, 0x72, 0x5f,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x6d, 0x43, 0x61, 0x72, 0x64, 0x53, 0x74,
	0x6f, 0x63, 0x6b, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69,
	0x65, 0x64, 0x48, 0x00, 0x52, 0x08, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x4b,
	0x0a, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2f, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x63, 0x65, 0x6c, 0x6c, 0x75, 0x6c, 0x61, 0x72, 0x5f, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x6d, 0x43, 0x61, 0x72, 0x64, 0x53, 0x74, 0x6f,
	0x63, 0x6b, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x48, 0x00, 0x52, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x4b, 0x0a, 0x07, 0x72,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x6e,
	0x74, 0x74, 0x2e, 0x63, 0x65, 0x6c, 0x6c, 0x75, 0x6c, 0x61, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x69, 0x6d, 0x43, 0x61, 0x72, 0x64, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x48, 0x00, 0x52,
	0x07, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x1a, 0x6f, 0x0a, 0x05, 0x41, 0x64, 0x64, 0x65,
	0x64, 0x12, 0x47, 0x0a, 0x0e, 0x73, 0x69, 0x6d, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x73, 0x74,
	0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6e, 0x74, 0x74, 0x2e,
	0x63, 0x65, 0x6c, 0x6c, 0x75, 0x6c, 0x61, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x69, 0x6d, 0x43, 0x61, 0x72, 0x64, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x0c, 0x73, 0x69,
	0x6d, 0x43, 0x61, 0x72, 0x64, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x69,
	0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x76, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x1a, 0x9d, 0x02, 0x0a, 0x08, 0x4d, 0x6f,
	0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x28, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x14, 0xb2, 0xda, 0x21, 0x10, 0x0a, 0x0e, 0x0a, 0x0c, 0x53, 0x69,
	0x6d, 0x43, 0x61, 0x72, 0x64, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x47, 0x0a, 0x0e, 0x73, 0x69, 0x6d, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x73, 0x74, 0x6f,
	0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x63,
	0x65, 0x6c, 0x6c, 0x75, 0x6c, 0x61, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x69, 0x6d, 0x43, 0x61, 0x72, 0x64, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x0c, 0x73, 0x69, 0x6d,
	0x43, 0x61, 0x72, 0x64, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x4f, 0x0a, 0x0a, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x42, 0x14, 0xb2, 0xda, 0x21, 0x10, 0x32,
	0x0e, 0x0a, 0x0c, 0x53, 0x69, 0x6d, 0x43, 0x61, 0x72, 0x64, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52,
	0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x2e, 0x0a, 0x13, 0x70, 0x72,
	0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x5f, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75,
	0x73, 0x56, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x69,
	0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x76, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x1a, 0x52, 0x0a, 0x07, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x12, 0x47, 0x0a, 0x0e, 0x73, 0x69, 0x6d, 0x5f, 0x63, 0x61, 0x72, 0x64,
	0x5f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6e,
	0x74, 0x74, 0x2e, 0x63, 0x65, 0x6c, 0x6c, 0x75, 0x6c, 0x61, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x69, 0x6d, 0x43, 0x61, 0x72, 0x64, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52,
	0x0c, 0x73, 0x69, 0x6d, 0x43, 0x61, 0x72, 0x64, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x1a, 0x52, 0x0a,
	0x07, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x12, 0x28, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x14, 0xb2, 0xda, 0x21, 0x10, 0x0a, 0x0e, 0x0a, 0x0c,
	0x53, 0x69, 0x6d, 0x43, 0x61, 0x72, 0x64, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x76, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x3a, 0x12, 0x9a, 0xd9, 0x21, 0x0e, 0x0a, 0x0c, 0x53, 0x69, 0x6d, 0x43, 0x61, 0x72, 0x64,
	0x53, 0x74, 0x6f, 0x63, 0x6b, 0x42, 0x0d, 0x0a, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x42, 0x8f, 0x01, 0xe8, 0xde, 0x21, 0x00, 0x0a, 0x1a, 0x63, 0x6f, 0x6d,
	0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x63, 0x65, 0x6c, 0x6c, 0x75, 0x6c, 0x61, 0x72, 0x5f, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x42, 0x17, 0x53, 0x69, 0x6d, 0x43, 0x61, 0x72, 0x64,
	0x53, 0x74, 0x6f, 0x63, 0x6b, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x00, 0x5a, 0x52, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x63,
	0x65, 0x6c, 0x6c, 0x75, 0x6c, 0x61, 0x72, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x69, 0x6d, 0x5f, 0x63, 0x61, 0x72,
	0x64, 0x5f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x3b, 0x73, 0x69, 0x6d, 0x5f, 0x63, 0x61, 0x72, 0x64,
	0x5f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	edgelq_cellular_api_proto_v1_sim_card_stock_change_proto_rawDescOnce sync.Once
	edgelq_cellular_api_proto_v1_sim_card_stock_change_proto_rawDescData = edgelq_cellular_api_proto_v1_sim_card_stock_change_proto_rawDesc
)

func edgelq_cellular_api_proto_v1_sim_card_stock_change_proto_rawDescGZIP() []byte {
	edgelq_cellular_api_proto_v1_sim_card_stock_change_proto_rawDescOnce.Do(func() {
		edgelq_cellular_api_proto_v1_sim_card_stock_change_proto_rawDescData = protoimpl.X.CompressGZIP(edgelq_cellular_api_proto_v1_sim_card_stock_change_proto_rawDescData)
	})
	return edgelq_cellular_api_proto_v1_sim_card_stock_change_proto_rawDescData
}

var edgelq_cellular_api_proto_v1_sim_card_stock_change_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var edgelq_cellular_api_proto_v1_sim_card_stock_change_proto_goTypes = []interface{}{
	(*SimCardStockChange)(nil),          // 0: ntt.cellular_api.v1.SimCardStockChange
	(*SimCardStockChange_Added)(nil),    // 1: ntt.cellular_api.v1.SimCardStockChange.Added
	(*SimCardStockChange_Modified)(nil), // 2: ntt.cellular_api.v1.SimCardStockChange.Modified
	(*SimCardStockChange_Current)(nil),  // 3: ntt.cellular_api.v1.SimCardStockChange.Current
	(*SimCardStockChange_Removed)(nil),  // 4: ntt.cellular_api.v1.SimCardStockChange.Removed
	(*SimCardStock)(nil),                // 5: ntt.cellular_api.v1.SimCardStock
	(*SimCardStock_FieldMask)(nil),      // 6: ntt.cellular_api.v1.SimCardStock_FieldMask
}
var edgelq_cellular_api_proto_v1_sim_card_stock_change_proto_depIdxs = []int32{
	1, // 0: ntt.cellular_api.v1.SimCardStockChange.added:type_name -> ntt.cellular_api.v1.SimCardStockChange.Added
	2, // 1: ntt.cellular_api.v1.SimCardStockChange.modified:type_name -> ntt.cellular_api.v1.SimCardStockChange.Modified
	3, // 2: ntt.cellular_api.v1.SimCardStockChange.current:type_name -> ntt.cellular_api.v1.SimCardStockChange.Current
	4, // 3: ntt.cellular_api.v1.SimCardStockChange.removed:type_name -> ntt.cellular_api.v1.SimCardStockChange.Removed
	5, // 4: ntt.cellular_api.v1.SimCardStockChange.Added.sim_card_stock:type_name -> ntt.cellular_api.v1.SimCardStock
	5, // 5: ntt.cellular_api.v1.SimCardStockChange.Modified.sim_card_stock:type_name -> ntt.cellular_api.v1.SimCardStock
	6, // 6: ntt.cellular_api.v1.SimCardStockChange.Modified.field_mask:type_name -> ntt.cellular_api.v1.SimCardStock_FieldMask
	5, // 7: ntt.cellular_api.v1.SimCardStockChange.Current.sim_card_stock:type_name -> ntt.cellular_api.v1.SimCardStock
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { edgelq_cellular_api_proto_v1_sim_card_stock_change_proto_init() }
func edgelq_cellular_api_proto_v1_sim_card_stock_change_proto_init() {
	if edgelq_cellular_api_proto_v1_sim_card_stock_change_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		edgelq_cellular_api_proto_v1_sim_card_stock_change_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimCardStockChange); i {
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
		edgelq_cellular_api_proto_v1_sim_card_stock_change_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimCardStockChange_Added); i {
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
		edgelq_cellular_api_proto_v1_sim_card_stock_change_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimCardStockChange_Modified); i {
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
		edgelq_cellular_api_proto_v1_sim_card_stock_change_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimCardStockChange_Current); i {
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
		edgelq_cellular_api_proto_v1_sim_card_stock_change_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimCardStockChange_Removed); i {
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

	edgelq_cellular_api_proto_v1_sim_card_stock_change_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*SimCardStockChange_Added_)(nil),
		(*SimCardStockChange_Modified_)(nil),
		(*SimCardStockChange_Current_)(nil),
		(*SimCardStockChange_Removed_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: edgelq_cellular_api_proto_v1_sim_card_stock_change_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           edgelq_cellular_api_proto_v1_sim_card_stock_change_proto_goTypes,
		DependencyIndexes: edgelq_cellular_api_proto_v1_sim_card_stock_change_proto_depIdxs,
		MessageInfos:      edgelq_cellular_api_proto_v1_sim_card_stock_change_proto_msgTypes,
	}.Build()
	edgelq_cellular_api_proto_v1_sim_card_stock_change_proto = out.File
	edgelq_cellular_api_proto_v1_sim_card_stock_change_proto_rawDesc = nil
	edgelq_cellular_api_proto_v1_sim_card_stock_change_proto_goTypes = nil
	edgelq_cellular_api_proto_v1_sim_card_stock_change_proto_depIdxs = nil
}
