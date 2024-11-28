// Code generated by protoc-gen-goten-go
// File: edgelq/monitoring/proto/v4/alert_custom.proto
// DO NOT EDIT!!!

package alert_client

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
	alert "github.com/cloudwan/edgelq-sdk/monitoring/resources/v4/alert"
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
	_ = &alert.Alert{}
	_ = &fieldmaskpb.FieldMask{}
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// A request message of the BulkCreateAlerts method.
type BulkCreateAlertsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Parent name of ntt.monitoring.v4.Alert
	Parent *alert.ParentName `protobuf:"bytes,1,opt,customtype=ParentName,name=parent,proto3" json:"parent,omitempty" firestore:"parent"`
	Alerts []*alert.Alert    `protobuf:"bytes,2,rep,name=alerts,proto3" json:"alerts,omitempty" firestore:"alerts"`
}

func (m *BulkCreateAlertsRequest) Reset() {
	*m = BulkCreateAlertsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_monitoring_proto_v4_alert_custom_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *BulkCreateAlertsRequest) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*BulkCreateAlertsRequest) ProtoMessage() {}

func (m *BulkCreateAlertsRequest) ProtoReflect() preflect.Message {
	mi := &edgelq_monitoring_proto_v4_alert_custom_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*BulkCreateAlertsRequest) GotenMessage() {}

// Deprecated, Use BulkCreateAlertsRequest.ProtoReflect.Descriptor instead.
func (*BulkCreateAlertsRequest) Descriptor() ([]byte, []int) {
	return edgelq_monitoring_proto_v4_alert_custom_proto_rawDescGZIP(), []int{0}
}

func (m *BulkCreateAlertsRequest) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *BulkCreateAlertsRequest) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *BulkCreateAlertsRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *BulkCreateAlertsRequest) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *BulkCreateAlertsRequest) GetParent() *alert.ParentName {
	if m != nil {
		return m.Parent
	}
	return nil
}

func (m *BulkCreateAlertsRequest) GetAlerts() []*alert.Alert {
	if m != nil {
		return m.Alerts
	}
	return nil
}

func (m *BulkCreateAlertsRequest) SetParent(fv *alert.ParentName) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Parent", "BulkCreateAlertsRequest"))
	}
	m.Parent = fv
}

func (m *BulkCreateAlertsRequest) SetAlerts(fv []*alert.Alert) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Alerts", "BulkCreateAlertsRequest"))
	}
	m.Alerts = fv
}

// A response message of the BulkCreateAlerts method.
type BulkCreateAlertsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (m *BulkCreateAlertsResponse) Reset() {
	*m = BulkCreateAlertsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_monitoring_proto_v4_alert_custom_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *BulkCreateAlertsResponse) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*BulkCreateAlertsResponse) ProtoMessage() {}

func (m *BulkCreateAlertsResponse) ProtoReflect() preflect.Message {
	mi := &edgelq_monitoring_proto_v4_alert_custom_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*BulkCreateAlertsResponse) GotenMessage() {}

// Deprecated, Use BulkCreateAlertsResponse.ProtoReflect.Descriptor instead.
func (*BulkCreateAlertsResponse) Descriptor() ([]byte, []int) {
	return edgelq_monitoring_proto_v4_alert_custom_proto_rawDescGZIP(), []int{1}
}

func (m *BulkCreateAlertsResponse) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *BulkCreateAlertsResponse) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *BulkCreateAlertsResponse) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *BulkCreateAlertsResponse) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

// A request message of the BulkUpdateAlerts method.
type BulkUpdateAlertsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Parent name of ntt.monitoring.v4.Alert
	Parent     *alert.ParentName      `protobuf:"bytes,1,opt,customtype=ParentName,name=parent,proto3" json:"parent,omitempty" firestore:"parent"`
	UpdateMask *alert.Alert_FieldMask `protobuf:"bytes,2,opt,customtype=Alert_FieldMask,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty" firestore:"updateMask"`
	Alerts     []*alert.Alert         `protobuf:"bytes,3,rep,name=alerts,proto3" json:"alerts,omitempty" firestore:"alerts"`
}

func (m *BulkUpdateAlertsRequest) Reset() {
	*m = BulkUpdateAlertsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_monitoring_proto_v4_alert_custom_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *BulkUpdateAlertsRequest) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*BulkUpdateAlertsRequest) ProtoMessage() {}

func (m *BulkUpdateAlertsRequest) ProtoReflect() preflect.Message {
	mi := &edgelq_monitoring_proto_v4_alert_custom_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*BulkUpdateAlertsRequest) GotenMessage() {}

// Deprecated, Use BulkUpdateAlertsRequest.ProtoReflect.Descriptor instead.
func (*BulkUpdateAlertsRequest) Descriptor() ([]byte, []int) {
	return edgelq_monitoring_proto_v4_alert_custom_proto_rawDescGZIP(), []int{2}
}

func (m *BulkUpdateAlertsRequest) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *BulkUpdateAlertsRequest) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *BulkUpdateAlertsRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *BulkUpdateAlertsRequest) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *BulkUpdateAlertsRequest) GetParent() *alert.ParentName {
	if m != nil {
		return m.Parent
	}
	return nil
}

func (m *BulkUpdateAlertsRequest) GetUpdateMask() *alert.Alert_FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

func (m *BulkUpdateAlertsRequest) GetAlerts() []*alert.Alert {
	if m != nil {
		return m.Alerts
	}
	return nil
}

func (m *BulkUpdateAlertsRequest) SetParent(fv *alert.ParentName) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Parent", "BulkUpdateAlertsRequest"))
	}
	m.Parent = fv
}

func (m *BulkUpdateAlertsRequest) SetUpdateMask(fv *alert.Alert_FieldMask) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "UpdateMask", "BulkUpdateAlertsRequest"))
	}
	m.UpdateMask = fv
}

func (m *BulkUpdateAlertsRequest) SetAlerts(fv []*alert.Alert) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Alerts", "BulkUpdateAlertsRequest"))
	}
	m.Alerts = fv
}

// A response message of the BulkUpdateAlerts method.
type BulkUpdateAlertsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (m *BulkUpdateAlertsResponse) Reset() {
	*m = BulkUpdateAlertsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_monitoring_proto_v4_alert_custom_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *BulkUpdateAlertsResponse) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*BulkUpdateAlertsResponse) ProtoMessage() {}

func (m *BulkUpdateAlertsResponse) ProtoReflect() preflect.Message {
	mi := &edgelq_monitoring_proto_v4_alert_custom_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*BulkUpdateAlertsResponse) GotenMessage() {}

// Deprecated, Use BulkUpdateAlertsResponse.ProtoReflect.Descriptor instead.
func (*BulkUpdateAlertsResponse) Descriptor() ([]byte, []int) {
	return edgelq_monitoring_proto_v4_alert_custom_proto_rawDescGZIP(), []int{3}
}

func (m *BulkUpdateAlertsResponse) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *BulkUpdateAlertsResponse) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *BulkUpdateAlertsResponse) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *BulkUpdateAlertsResponse) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

var edgelq_monitoring_proto_v4_alert_custom_proto preflect.FileDescriptor

var edgelq_monitoring_proto_v4_alert_custom_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x69, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x34, 0x2f, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x11, 0x6e, 0x74, 0x74, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x34, 0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f, 0x72, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x65, 0x64,
	0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x34, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2d, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6d, 0x6f, 0x6e,
	0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x34,
	0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x82, 0x01, 0x0a, 0x17, 0x42, 0x75, 0x6c, 0x6b, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x25, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x0d, 0xb2, 0xda, 0x21, 0x09, 0x3a, 0x07, 0x0a, 0x05, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x06,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x40, 0x0a, 0x06, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x6d, 0x6f, 0x6e,
	0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x34, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74,
	0x42, 0x0e, 0xca, 0xc6, 0x27, 0x0a, 0x42, 0x08, 0x0a, 0x02, 0x08, 0x01, 0x12, 0x02, 0x08, 0x19,
	0x52, 0x06, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x22, 0x1a, 0x0a, 0x18, 0x42, 0x75, 0x6c, 0x6b,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0xce, 0x01, 0x0a, 0x17, 0x42, 0x75, 0x6c, 0x6b, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x25, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x0d, 0xb2, 0xda, 0x21, 0x09, 0x3a, 0x07, 0x0a, 0x05, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52,
	0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x4a, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x42, 0x0d, 0xb2, 0xda, 0x21, 0x09, 0x32, 0x07,
	0x0a, 0x05, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d,
	0x61, 0x73, 0x6b, 0x12, 0x40, 0x0a, 0x06, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f,
	0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x34, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x42, 0x0e, 0xca,
	0xc6, 0x27, 0x0a, 0x42, 0x08, 0x0a, 0x02, 0x08, 0x01, 0x12, 0x02, 0x08, 0x19, 0x52, 0x06, 0x61,
	0x6c, 0x65, 0x72, 0x74, 0x73, 0x22, 0x1a, 0x0a, 0x18, 0x42, 0x75, 0x6c, 0x6b, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x76, 0xe8, 0xde, 0x21, 0x00, 0x0a, 0x18, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x74, 0x74,
	0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x62, 0x2e, 0x76,
	0x34, 0x42, 0x10, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x00, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c,
	0x71, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2f, 0x76, 0x34, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x3b, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	edgelq_monitoring_proto_v4_alert_custom_proto_rawDescOnce sync.Once
	edgelq_monitoring_proto_v4_alert_custom_proto_rawDescData = edgelq_monitoring_proto_v4_alert_custom_proto_rawDesc
)

func edgelq_monitoring_proto_v4_alert_custom_proto_rawDescGZIP() []byte {
	edgelq_monitoring_proto_v4_alert_custom_proto_rawDescOnce.Do(func() {
		edgelq_monitoring_proto_v4_alert_custom_proto_rawDescData = protoimpl.X.CompressGZIP(edgelq_monitoring_proto_v4_alert_custom_proto_rawDescData)
	})
	return edgelq_monitoring_proto_v4_alert_custom_proto_rawDescData
}

var edgelq_monitoring_proto_v4_alert_custom_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var edgelq_monitoring_proto_v4_alert_custom_proto_goTypes = []interface{}{
	(*BulkCreateAlertsRequest)(nil),  // 0: ntt.monitoring.v4.BulkCreateAlertsRequest
	(*BulkCreateAlertsResponse)(nil), // 1: ntt.monitoring.v4.BulkCreateAlertsResponse
	(*BulkUpdateAlertsRequest)(nil),  // 2: ntt.monitoring.v4.BulkUpdateAlertsRequest
	(*BulkUpdateAlertsResponse)(nil), // 3: ntt.monitoring.v4.BulkUpdateAlertsResponse
	(*alert.Alert)(nil),              // 4: ntt.monitoring.v4.Alert
	(*alert.Alert_FieldMask)(nil),    // 5: ntt.monitoring.v4.Alert_FieldMask
}
var edgelq_monitoring_proto_v4_alert_custom_proto_depIdxs = []int32{
	4, // 0: ntt.monitoring.v4.BulkCreateAlertsRequest.alerts:type_name -> ntt.monitoring.v4.Alert
	5, // 1: ntt.monitoring.v4.BulkUpdateAlertsRequest.update_mask:type_name -> ntt.monitoring.v4.Alert_FieldMask
	4, // 2: ntt.monitoring.v4.BulkUpdateAlertsRequest.alerts:type_name -> ntt.monitoring.v4.Alert
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { edgelq_monitoring_proto_v4_alert_custom_proto_init() }
func edgelq_monitoring_proto_v4_alert_custom_proto_init() {
	if edgelq_monitoring_proto_v4_alert_custom_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		edgelq_monitoring_proto_v4_alert_custom_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BulkCreateAlertsRequest); i {
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
		edgelq_monitoring_proto_v4_alert_custom_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BulkCreateAlertsResponse); i {
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
		edgelq_monitoring_proto_v4_alert_custom_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BulkUpdateAlertsRequest); i {
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
		edgelq_monitoring_proto_v4_alert_custom_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BulkUpdateAlertsResponse); i {
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

	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: edgelq_monitoring_proto_v4_alert_custom_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           edgelq_monitoring_proto_v4_alert_custom_proto_goTypes,
		DependencyIndexes: edgelq_monitoring_proto_v4_alert_custom_proto_depIdxs,
		MessageInfos:      edgelq_monitoring_proto_v4_alert_custom_proto_msgTypes,
	}.Build()
	edgelq_monitoring_proto_v4_alert_custom_proto = out.File
	edgelq_monitoring_proto_v4_alert_custom_proto_rawDesc = nil
	edgelq_monitoring_proto_v4_alert_custom_proto_goTypes = nil
	edgelq_monitoring_proto_v4_alert_custom_proto_depIdxs = nil
}
