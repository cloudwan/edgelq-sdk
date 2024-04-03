// Code generated by protoc-gen-goten-go
// File: edgelq/devices/proto/v1/provisioning_approval_request_custom.proto
// DO NOT EDIT!!!

package provisioning_approval_request_client

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
	device "github.com/cloudwan/edgelq-sdk/devices/resources/v1/device"
	provisioning_approval_request "github.com/cloudwan/edgelq-sdk/devices/resources/v1/provisioning_approval_request"
	provisioning_policy "github.com/cloudwan/edgelq-sdk/devices/resources/v1/provisioning_policy"
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
	_ = &device.Device{}
	_ = &provisioning_approval_request.ProvisioningApprovalRequest{}
	_ = &provisioning_policy.ProvisioningPolicy{}
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Request message for method
// [ProvisionDeviceForApprovedRequest][ntt.devices.v1.ProvisionDeviceForApprovedRequest]
type ProvisionDeviceForApprovedRequestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// name of ntt.devices.v1.ProvisioningApprovalRequest
	Name         *provisioning_approval_request.Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	DeviceStatus *device.Device_Status               `protobuf:"bytes,2,opt,name=device_status,json=deviceStatus,proto3" json:"device_status,omitempty" firestore:"deviceStatus"`
}

func (m *ProvisionDeviceForApprovedRequestRequest) Reset() {
	*m = ProvisionDeviceForApprovedRequestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_devices_proto_v1_provisioning_approval_request_custom_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *ProvisionDeviceForApprovedRequestRequest) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*ProvisionDeviceForApprovedRequestRequest) ProtoMessage() {}

func (m *ProvisionDeviceForApprovedRequestRequest) ProtoReflect() preflect.Message {
	mi := &edgelq_devices_proto_v1_provisioning_approval_request_custom_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*ProvisionDeviceForApprovedRequestRequest) GotenMessage() {}

// Deprecated, Use ProvisionDeviceForApprovedRequestRequest.ProtoReflect.Descriptor instead.
func (*ProvisionDeviceForApprovedRequestRequest) Descriptor() ([]byte, []int) {
	return edgelq_devices_proto_v1_provisioning_approval_request_custom_proto_rawDescGZIP(), []int{0}
}

func (m *ProvisionDeviceForApprovedRequestRequest) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *ProvisionDeviceForApprovedRequestRequest) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *ProvisionDeviceForApprovedRequestRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *ProvisionDeviceForApprovedRequestRequest) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *ProvisionDeviceForApprovedRequestRequest) GetName() *provisioning_approval_request.Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *ProvisionDeviceForApprovedRequestRequest) GetDeviceStatus() *device.Device_Status {
	if m != nil {
		return m.DeviceStatus
	}
	return nil
}

func (m *ProvisionDeviceForApprovedRequestRequest) SetName(fv *provisioning_approval_request.Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "ProvisionDeviceForApprovedRequestRequest"))
	}
	m.Name = fv
}

func (m *ProvisionDeviceForApprovedRequestRequest) SetDeviceStatus(fv *device.Device_Status) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "DeviceStatus", "ProvisionDeviceForApprovedRequestRequest"))
	}
	m.DeviceStatus = fv
}

// Response message for method
// [ProvisionDeviceForApprovedRequest][ntt.devices.v1.ProvisionDeviceForApprovedRequest]
type ProvisionDeviceForApprovedRequestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Device        *device.Device `protobuf:"bytes,1,opt,name=device,proto3" json:"device,omitempty" firestore:"device"`
}

func (m *ProvisionDeviceForApprovedRequestResponse) Reset() {
	*m = ProvisionDeviceForApprovedRequestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_devices_proto_v1_provisioning_approval_request_custom_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *ProvisionDeviceForApprovedRequestResponse) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*ProvisionDeviceForApprovedRequestResponse) ProtoMessage() {}

func (m *ProvisionDeviceForApprovedRequestResponse) ProtoReflect() preflect.Message {
	mi := &edgelq_devices_proto_v1_provisioning_approval_request_custom_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*ProvisionDeviceForApprovedRequestResponse) GotenMessage() {}

// Deprecated, Use ProvisionDeviceForApprovedRequestResponse.ProtoReflect.Descriptor instead.
func (*ProvisionDeviceForApprovedRequestResponse) Descriptor() ([]byte, []int) {
	return edgelq_devices_proto_v1_provisioning_approval_request_custom_proto_rawDescGZIP(), []int{1}
}

func (m *ProvisionDeviceForApprovedRequestResponse) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *ProvisionDeviceForApprovedRequestResponse) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *ProvisionDeviceForApprovedRequestResponse) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *ProvisionDeviceForApprovedRequestResponse) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *ProvisionDeviceForApprovedRequestResponse) GetDevice() *device.Device {
	if m != nil {
		return m.Device
	}
	return nil
}

func (m *ProvisionDeviceForApprovedRequestResponse) SetDevice(fv *device.Device) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Device", "ProvisionDeviceForApprovedRequestResponse"))
	}
	m.Device = fv
}

var edgelq_devices_proto_v1_provisioning_approval_request_custom_proto preflect.FileDescriptor

var edgelq_devices_proto_v1_provisioning_approval_request_custom_proto_rawDesc = []byte{
	0x0a, 0x42, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x5f,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x6e, 0x74, 0x74, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f,
	0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x65, 0x64,
	0x67, 0x65, 0x6c, 0x71, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x3b, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61,
	0x6c, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x42, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69,
	0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x5f, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd9, 0x01, 0x0a, 0x28, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x46, 0x6f, 0x72,
	0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x23, 0xb2, 0xda, 0x21, 0x1f, 0x0a, 0x1d, 0x0a, 0x1b, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76,
	0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x42, 0x0a, 0x0d, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x3a, 0x30, 0xc2, 0x85, 0x2c, 0x2c, 0x0a, 0x2a, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x22, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x5f,
	0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x5b, 0x0a, 0x29, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69,
	0x6f, 0x6e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x46, 0x6f, 0x72, 0x41, 0x70, 0x70, 0x72, 0x6f,
	0x76, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x06, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x42, 0xb6, 0x01, 0xe8, 0xde, 0x21, 0x00, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x6e,
	0x74, 0x74, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31,
	0x42, 0x26, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x41, 0x70,
	0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x00, 0x5a, 0x6f, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f,
	0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69,
	0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x5f, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x3b, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x69, 0x6e, 0x67, 0x5f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	edgelq_devices_proto_v1_provisioning_approval_request_custom_proto_rawDescOnce sync.Once
	edgelq_devices_proto_v1_provisioning_approval_request_custom_proto_rawDescData = edgelq_devices_proto_v1_provisioning_approval_request_custom_proto_rawDesc
)

func edgelq_devices_proto_v1_provisioning_approval_request_custom_proto_rawDescGZIP() []byte {
	edgelq_devices_proto_v1_provisioning_approval_request_custom_proto_rawDescOnce.Do(func() {
		edgelq_devices_proto_v1_provisioning_approval_request_custom_proto_rawDescData = protoimpl.X.CompressGZIP(edgelq_devices_proto_v1_provisioning_approval_request_custom_proto_rawDescData)
	})
	return edgelq_devices_proto_v1_provisioning_approval_request_custom_proto_rawDescData
}

var edgelq_devices_proto_v1_provisioning_approval_request_custom_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var edgelq_devices_proto_v1_provisioning_approval_request_custom_proto_goTypes = []interface{}{
	(*ProvisionDeviceForApprovedRequestRequest)(nil),  // 0: ntt.devices.v1.ProvisionDeviceForApprovedRequestRequest
	(*ProvisionDeviceForApprovedRequestResponse)(nil), // 1: ntt.devices.v1.ProvisionDeviceForApprovedRequestResponse
	(*device.Device_Status)(nil),                      // 2: ntt.devices.v1.Device.Status
	(*device.Device)(nil),                             // 3: ntt.devices.v1.Device
}
var edgelq_devices_proto_v1_provisioning_approval_request_custom_proto_depIdxs = []int32{
	2, // 0: ntt.devices.v1.ProvisionDeviceForApprovedRequestRequest.device_status:type_name -> ntt.devices.v1.Device.Status
	3, // 1: ntt.devices.v1.ProvisionDeviceForApprovedRequestResponse.device:type_name -> ntt.devices.v1.Device
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { edgelq_devices_proto_v1_provisioning_approval_request_custom_proto_init() }
func edgelq_devices_proto_v1_provisioning_approval_request_custom_proto_init() {
	if edgelq_devices_proto_v1_provisioning_approval_request_custom_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		edgelq_devices_proto_v1_provisioning_approval_request_custom_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProvisionDeviceForApprovedRequestRequest); i {
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
		edgelq_devices_proto_v1_provisioning_approval_request_custom_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProvisionDeviceForApprovedRequestResponse); i {
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
			RawDescriptor: edgelq_devices_proto_v1_provisioning_approval_request_custom_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           edgelq_devices_proto_v1_provisioning_approval_request_custom_proto_goTypes,
		DependencyIndexes: edgelq_devices_proto_v1_provisioning_approval_request_custom_proto_depIdxs,
		MessageInfos:      edgelq_devices_proto_v1_provisioning_approval_request_custom_proto_msgTypes,
	}.Build()
	edgelq_devices_proto_v1_provisioning_approval_request_custom_proto = out.File
	edgelq_devices_proto_v1_provisioning_approval_request_custom_proto_rawDesc = nil
	edgelq_devices_proto_v1_provisioning_approval_request_custom_proto_goTypes = nil
	edgelq_devices_proto_v1_provisioning_approval_request_custom_proto_depIdxs = nil
}