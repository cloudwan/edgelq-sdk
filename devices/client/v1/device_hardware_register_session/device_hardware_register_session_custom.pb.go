// Code generated by protoc-gen-goten-go
// File: edgelq/devices/proto/v1/device_hardware_register_session_custom.proto
// DO NOT EDIT!!!

package device_hardware_register_session_client

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
	device_hardware "github.com/cloudwan/edgelq-sdk/devices/resources/v1/device_hardware"
	device_hardware_register_session "github.com/cloudwan/edgelq-sdk/devices/resources/v1/device_hardware_register_session"
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
	_ = &device_hardware.DeviceHardware{}
	_ = &device_hardware_register_session.DeviceHardwareRegisterSession{}
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Request message for method
// [RegisterHardware][ntt.devices.v1.RegisterHardware]
type RegisterHardwareRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// token from hardware_register_session
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty" firestore:"token"`
	// one of encrypted_serial_number or serial_number is required
	EncryptedSerialNumber string `protobuf:"bytes,2,opt,name=encrypted_serial_number,json=encryptedSerialNumber,proto3" json:"encrypted_serial_number,omitempty" firestore:"encryptedSerialNumber"`
	SerialNumber          string `protobuf:"bytes,3,opt,name=serial_number,json=serialNumber,proto3" json:"serial_number,omitempty" firestore:"serialNumber"`
	Manufacturer          string `protobuf:"bytes,4,opt,name=manufacturer,proto3" json:"manufacturer,omitempty" firestore:"manufacturer"`
	// Should be in the format "productname (sku)", required field
	ProductName string   `protobuf:"bytes,5,opt,name=product_name,json=productName,proto3" json:"product_name,omitempty" firestore:"productName"`
	MacAddress  []string `protobuf:"bytes,6,rep,name=mac_address,json=macAddress,proto3" json:"mac_address,omitempty" firestore:"macAddress"`
	SimIccid    string   `protobuf:"bytes,7,opt,name=sim_iccid,json=simIccid,proto3" json:"sim_iccid,omitempty" firestore:"simIccid"`
	Imei        string   `protobuf:"bytes,8,opt,name=imei,proto3" json:"imei,omitempty" firestore:"imei"`
}

func (m *RegisterHardwareRequest) Reset() {
	*m = RegisterHardwareRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_devices_proto_v1_device_hardware_register_session_custom_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *RegisterHardwareRequest) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*RegisterHardwareRequest) ProtoMessage() {}

func (m *RegisterHardwareRequest) ProtoReflect() preflect.Message {
	mi := &edgelq_devices_proto_v1_device_hardware_register_session_custom_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*RegisterHardwareRequest) GotenMessage() {}

// Deprecated, Use RegisterHardwareRequest.ProtoReflect.Descriptor instead.
func (*RegisterHardwareRequest) Descriptor() ([]byte, []int) {
	return edgelq_devices_proto_v1_device_hardware_register_session_custom_proto_rawDescGZIP(), []int{0}
}

func (m *RegisterHardwareRequest) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *RegisterHardwareRequest) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *RegisterHardwareRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *RegisterHardwareRequest) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *RegisterHardwareRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *RegisterHardwareRequest) GetEncryptedSerialNumber() string {
	if m != nil {
		return m.EncryptedSerialNumber
	}
	return ""
}

func (m *RegisterHardwareRequest) GetSerialNumber() string {
	if m != nil {
		return m.SerialNumber
	}
	return ""
}

func (m *RegisterHardwareRequest) GetManufacturer() string {
	if m != nil {
		return m.Manufacturer
	}
	return ""
}

func (m *RegisterHardwareRequest) GetProductName() string {
	if m != nil {
		return m.ProductName
	}
	return ""
}

func (m *RegisterHardwareRequest) GetMacAddress() []string {
	if m != nil {
		return m.MacAddress
	}
	return nil
}

func (m *RegisterHardwareRequest) GetSimIccid() string {
	if m != nil {
		return m.SimIccid
	}
	return ""
}

func (m *RegisterHardwareRequest) GetImei() string {
	if m != nil {
		return m.Imei
	}
	return ""
}

func (m *RegisterHardwareRequest) SetToken(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Token", "RegisterHardwareRequest"))
	}
	m.Token = fv
}

func (m *RegisterHardwareRequest) SetEncryptedSerialNumber(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "EncryptedSerialNumber", "RegisterHardwareRequest"))
	}
	m.EncryptedSerialNumber = fv
}

func (m *RegisterHardwareRequest) SetSerialNumber(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "SerialNumber", "RegisterHardwareRequest"))
	}
	m.SerialNumber = fv
}

func (m *RegisterHardwareRequest) SetManufacturer(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Manufacturer", "RegisterHardwareRequest"))
	}
	m.Manufacturer = fv
}

func (m *RegisterHardwareRequest) SetProductName(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ProductName", "RegisterHardwareRequest"))
	}
	m.ProductName = fv
}

func (m *RegisterHardwareRequest) SetMacAddress(fv []string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "MacAddress", "RegisterHardwareRequest"))
	}
	m.MacAddress = fv
}

func (m *RegisterHardwareRequest) SetSimIccid(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "SimIccid", "RegisterHardwareRequest"))
	}
	m.SimIccid = fv
}

func (m *RegisterHardwareRequest) SetImei(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Imei", "RegisterHardwareRequest"))
	}
	m.Imei = fv
}

// Response message for method
// [RegisterHardware][ntt.devices.v1.RegisterHardware]
type RegisterHardwareResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Hardware      *device_hardware.DeviceHardware `protobuf:"bytes,1,opt,name=hardware,proto3" json:"hardware,omitempty" firestore:"hardware"`
}

func (m *RegisterHardwareResponse) Reset() {
	*m = RegisterHardwareResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_devices_proto_v1_device_hardware_register_session_custom_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *RegisterHardwareResponse) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*RegisterHardwareResponse) ProtoMessage() {}

func (m *RegisterHardwareResponse) ProtoReflect() preflect.Message {
	mi := &edgelq_devices_proto_v1_device_hardware_register_session_custom_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*RegisterHardwareResponse) GotenMessage() {}

// Deprecated, Use RegisterHardwareResponse.ProtoReflect.Descriptor instead.
func (*RegisterHardwareResponse) Descriptor() ([]byte, []int) {
	return edgelq_devices_proto_v1_device_hardware_register_session_custom_proto_rawDescGZIP(), []int{1}
}

func (m *RegisterHardwareResponse) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *RegisterHardwareResponse) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *RegisterHardwareResponse) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *RegisterHardwareResponse) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *RegisterHardwareResponse) GetHardware() *device_hardware.DeviceHardware {
	if m != nil {
		return m.Hardware
	}
	return nil
}

func (m *RegisterHardwareResponse) SetHardware(fv *device_hardware.DeviceHardware) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Hardware", "RegisterHardwareResponse"))
	}
	m.Hardware = fv
}

// Request message for method [HardwareStatus][ntt.devices.v1.HardwareStatus]
type HardwareStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// token for hardware_register_session
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty" firestore:"token"`
	// one of encrypted_serial_number or serial_number or sim_iccid is mandatory
	EncryptedSerialNumber string `protobuf:"bytes,2,opt,name=encrypted_serial_number,json=encryptedSerialNumber,proto3" json:"encrypted_serial_number,omitempty" firestore:"encryptedSerialNumber"`
	SerialNumber          string `protobuf:"bytes,3,opt,name=serial_number,json=serialNumber,proto3" json:"serial_number,omitempty" firestore:"serialNumber"`
	// manufacturer is typically not required, only required in case of serial
	// number conflict with different vendors
	Manufacturer string `protobuf:"bytes,4,opt,name=manufacturer,proto3" json:"manufacturer,omitempty" firestore:"manufacturer"`
	// product_name is typically not required, only required in case of serial
	// number conflict with different productname/skus Should be in the format
	// "productname (sku)"
	ProductName string `protobuf:"bytes,5,opt,name=product_name,json=productName,proto3" json:"product_name,omitempty" firestore:"productName"`
}

func (m *HardwareStatusRequest) Reset() {
	*m = HardwareStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_devices_proto_v1_device_hardware_register_session_custom_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *HardwareStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*HardwareStatusRequest) ProtoMessage() {}

func (m *HardwareStatusRequest) ProtoReflect() preflect.Message {
	mi := &edgelq_devices_proto_v1_device_hardware_register_session_custom_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*HardwareStatusRequest) GotenMessage() {}

// Deprecated, Use HardwareStatusRequest.ProtoReflect.Descriptor instead.
func (*HardwareStatusRequest) Descriptor() ([]byte, []int) {
	return edgelq_devices_proto_v1_device_hardware_register_session_custom_proto_rawDescGZIP(), []int{2}
}

func (m *HardwareStatusRequest) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *HardwareStatusRequest) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *HardwareStatusRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *HardwareStatusRequest) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *HardwareStatusRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *HardwareStatusRequest) GetEncryptedSerialNumber() string {
	if m != nil {
		return m.EncryptedSerialNumber
	}
	return ""
}

func (m *HardwareStatusRequest) GetSerialNumber() string {
	if m != nil {
		return m.SerialNumber
	}
	return ""
}

func (m *HardwareStatusRequest) GetManufacturer() string {
	if m != nil {
		return m.Manufacturer
	}
	return ""
}

func (m *HardwareStatusRequest) GetProductName() string {
	if m != nil {
		return m.ProductName
	}
	return ""
}

func (m *HardwareStatusRequest) SetToken(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Token", "HardwareStatusRequest"))
	}
	m.Token = fv
}

func (m *HardwareStatusRequest) SetEncryptedSerialNumber(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "EncryptedSerialNumber", "HardwareStatusRequest"))
	}
	m.EncryptedSerialNumber = fv
}

func (m *HardwareStatusRequest) SetSerialNumber(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "SerialNumber", "HardwareStatusRequest"))
	}
	m.SerialNumber = fv
}

func (m *HardwareStatusRequest) SetManufacturer(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Manufacturer", "HardwareStatusRequest"))
	}
	m.Manufacturer = fv
}

func (m *HardwareStatusRequest) SetProductName(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ProductName", "HardwareStatusRequest"))
	}
	m.ProductName = fv
}

// Response message for method [HardwareStatus][ntt.devices.v1.HardwareStatus]
type HardwareStatusResponse struct {
	state           protoimpl.MessageState
	sizeCache       protoimpl.SizeCache
	unknownFields   protoimpl.UnknownFields
	DeviceHardwares []*device_hardware.DeviceHardware `protobuf:"bytes,1,rep,name=device_hardwares,json=deviceHardwares,proto3" json:"device_hardwares,omitempty" firestore:"deviceHardwares"`
}

func (m *HardwareStatusResponse) Reset() {
	*m = HardwareStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_devices_proto_v1_device_hardware_register_session_custom_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *HardwareStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*HardwareStatusResponse) ProtoMessage() {}

func (m *HardwareStatusResponse) ProtoReflect() preflect.Message {
	mi := &edgelq_devices_proto_v1_device_hardware_register_session_custom_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*HardwareStatusResponse) GotenMessage() {}

// Deprecated, Use HardwareStatusResponse.ProtoReflect.Descriptor instead.
func (*HardwareStatusResponse) Descriptor() ([]byte, []int) {
	return edgelq_devices_proto_v1_device_hardware_register_session_custom_proto_rawDescGZIP(), []int{3}
}

func (m *HardwareStatusResponse) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *HardwareStatusResponse) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *HardwareStatusResponse) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *HardwareStatusResponse) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *HardwareStatusResponse) GetDeviceHardwares() []*device_hardware.DeviceHardware {
	if m != nil {
		return m.DeviceHardwares
	}
	return nil
}

func (m *HardwareStatusResponse) SetDeviceHardwares(fv []*device_hardware.DeviceHardware) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "DeviceHardwares", "HardwareStatusResponse"))
	}
	m.DeviceHardwares = fv
}

// Request message for method
// [GetDeviceHardwareRegisterSessionFromToken][ntt.devices.v1.GetDeviceHardwareRegisterSessionFromToken]
type GetDeviceHardwareRegisterSessionFromTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// token from hardware_register_session
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty" firestore:"token"`
}

func (m *GetDeviceHardwareRegisterSessionFromTokenRequest) Reset() {
	*m = GetDeviceHardwareRegisterSessionFromTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_devices_proto_v1_device_hardware_register_session_custom_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *GetDeviceHardwareRegisterSessionFromTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*GetDeviceHardwareRegisterSessionFromTokenRequest) ProtoMessage() {}

func (m *GetDeviceHardwareRegisterSessionFromTokenRequest) ProtoReflect() preflect.Message {
	mi := &edgelq_devices_proto_v1_device_hardware_register_session_custom_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*GetDeviceHardwareRegisterSessionFromTokenRequest) GotenMessage() {}

// Deprecated, Use GetDeviceHardwareRegisterSessionFromTokenRequest.ProtoReflect.Descriptor instead.
func (*GetDeviceHardwareRegisterSessionFromTokenRequest) Descriptor() ([]byte, []int) {
	return edgelq_devices_proto_v1_device_hardware_register_session_custom_proto_rawDescGZIP(), []int{4}
}

func (m *GetDeviceHardwareRegisterSessionFromTokenRequest) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *GetDeviceHardwareRegisterSessionFromTokenRequest) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *GetDeviceHardwareRegisterSessionFromTokenRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *GetDeviceHardwareRegisterSessionFromTokenRequest) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *GetDeviceHardwareRegisterSessionFromTokenRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *GetDeviceHardwareRegisterSessionFromTokenRequest) SetToken(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Token", "GetDeviceHardwareRegisterSessionFromTokenRequest"))
	}
	m.Token = fv
}

// Response message for method
// [GetDeviceHardwareRegisterSessionFromToken][ntt.devices.v1.GetDeviceHardwareRegisterSessionFromToken]
type GetDeviceHardwareRegisterSessionFromTokenResponse struct {
	state                         protoimpl.MessageState
	sizeCache                     protoimpl.SizeCache
	unknownFields                 protoimpl.UnknownFields
	DeviceHardwareRegisterSession *device_hardware_register_session.DeviceHardwareRegisterSession `protobuf:"bytes,1,opt,name=device_hardware_register_session,json=deviceHardwareRegisterSession,proto3" json:"device_hardware_register_session,omitempty" firestore:"deviceHardwareRegisterSession"`
	// Used for showing project name in hardware registration dashboard
	ProjectDisplayName string `protobuf:"bytes,2,opt,name=project_display_name,json=projectDisplayName,proto3" json:"project_display_name,omitempty" firestore:"projectDisplayName"`
}

func (m *GetDeviceHardwareRegisterSessionFromTokenResponse) Reset() {
	*m = GetDeviceHardwareRegisterSessionFromTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_devices_proto_v1_device_hardware_register_session_custom_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *GetDeviceHardwareRegisterSessionFromTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*GetDeviceHardwareRegisterSessionFromTokenResponse) ProtoMessage() {}

func (m *GetDeviceHardwareRegisterSessionFromTokenResponse) ProtoReflect() preflect.Message {
	mi := &edgelq_devices_proto_v1_device_hardware_register_session_custom_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*GetDeviceHardwareRegisterSessionFromTokenResponse) GotenMessage() {}

// Deprecated, Use GetDeviceHardwareRegisterSessionFromTokenResponse.ProtoReflect.Descriptor instead.
func (*GetDeviceHardwareRegisterSessionFromTokenResponse) Descriptor() ([]byte, []int) {
	return edgelq_devices_proto_v1_device_hardware_register_session_custom_proto_rawDescGZIP(), []int{5}
}

func (m *GetDeviceHardwareRegisterSessionFromTokenResponse) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *GetDeviceHardwareRegisterSessionFromTokenResponse) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *GetDeviceHardwareRegisterSessionFromTokenResponse) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *GetDeviceHardwareRegisterSessionFromTokenResponse) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *GetDeviceHardwareRegisterSessionFromTokenResponse) GetDeviceHardwareRegisterSession() *device_hardware_register_session.DeviceHardwareRegisterSession {
	if m != nil {
		return m.DeviceHardwareRegisterSession
	}
	return nil
}

func (m *GetDeviceHardwareRegisterSessionFromTokenResponse) GetProjectDisplayName() string {
	if m != nil {
		return m.ProjectDisplayName
	}
	return ""
}

func (m *GetDeviceHardwareRegisterSessionFromTokenResponse) SetDeviceHardwareRegisterSession(fv *device_hardware_register_session.DeviceHardwareRegisterSession) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "DeviceHardwareRegisterSession", "GetDeviceHardwareRegisterSessionFromTokenResponse"))
	}
	m.DeviceHardwareRegisterSession = fv
}

func (m *GetDeviceHardwareRegisterSessionFromTokenResponse) SetProjectDisplayName(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ProjectDisplayName", "GetDeviceHardwareRegisterSessionFromTokenResponse"))
	}
	m.ProjectDisplayName = fv
}

var edgelq_devices_proto_v1_device_hardware_register_session_custom_proto preflect.FileDescriptor

var edgelq_devices_proto_v1_device_hardware_register_session_custom_proto_rawDesc = []byte{
	0x0a, 0x45, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x68, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x6e, 0x74, 0x74, 0x2e, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x74,
	0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x67, 0x6f, 0x74, 0x65, 0x6e,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6d, 0x75, 0x6c,
	0x74, 0x69, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x2d, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f,
	0x68, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3e,
	0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x68,
	0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x45,
	0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x68,
	0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa5, 0x02, 0x0a, 0x17, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x48, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x36, 0x0a, 0x17, 0x65, 0x6e, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x65, 0x64, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x23, 0x0a, 0x0d, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x0c, 0x6d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74,
	0x75, 0x72, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6d, 0x61, 0x6e, 0x75,
	0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6d,
	0x61, 0x63, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0a, 0x6d, 0x61, 0x63, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1b, 0x0a, 0x09,
	0x73, 0x69, 0x6d, 0x5f, 0x69, 0x63, 0x63, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x73, 0x69, 0x6d, 0x49, 0x63, 0x63, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6d, 0x65,
	0x69, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6d, 0x65, 0x69, 0x22, 0x56, 0x0a,
	0x18, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x48, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x08, 0x68, 0x61, 0x72,
	0x64, 0x77, 0x61, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6e, 0x74,
	0x74, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x48, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x52, 0x08, 0x68, 0x61, 0x72,
	0x64, 0x77, 0x61, 0x72, 0x65, 0x22, 0xd1, 0x01, 0x0a, 0x15, 0x48, 0x61, 0x72, 0x64, 0x77, 0x61,
	0x72, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x36, 0x0a, 0x17, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x65, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65,
	0x64, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x23, 0x0a,
	0x0d, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x12, 0x22, 0x0a, 0x0c, 0x6d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72,
	0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6d, 0x61, 0x6e, 0x75, 0x66, 0x61,
	0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x63, 0x0a, 0x16, 0x48, 0x61, 0x72,
	0x64, 0x77, 0x61, 0x72, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x10, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x68, 0x61,
	0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x6e, 0x74, 0x74, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x48, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x52, 0x0f, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x48, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x73, 0x22, 0x48,
	0x0a, 0x30, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x48, 0x61, 0x72, 0x64, 0x77,
	0x61, 0x72, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x46, 0x72, 0x6f, 0x6d, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xdd, 0x01, 0x0a, 0x31, 0x47, 0x65, 0x74,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x48, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x46, 0x72, 0x6f,
	0x6d, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x76,
	0x0a, 0x20, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x68, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72,
	0x65, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x48, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x1d, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x48,
	0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x53,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x14, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x5f, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x44, 0x69, 0x73,
	0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0xbe, 0x01, 0xe8, 0xde, 0x21, 0x00, 0x0a,
	0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x42, 0x28, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x48, 0x61,
	0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x00, 0x5a, 0x75, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31,
	0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x68, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65,
	0x5f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x3b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x68, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72,
	0x65, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	edgelq_devices_proto_v1_device_hardware_register_session_custom_proto_rawDescOnce sync.Once
	edgelq_devices_proto_v1_device_hardware_register_session_custom_proto_rawDescData = edgelq_devices_proto_v1_device_hardware_register_session_custom_proto_rawDesc
)

func edgelq_devices_proto_v1_device_hardware_register_session_custom_proto_rawDescGZIP() []byte {
	edgelq_devices_proto_v1_device_hardware_register_session_custom_proto_rawDescOnce.Do(func() {
		edgelq_devices_proto_v1_device_hardware_register_session_custom_proto_rawDescData = protoimpl.X.CompressGZIP(edgelq_devices_proto_v1_device_hardware_register_session_custom_proto_rawDescData)
	})
	return edgelq_devices_proto_v1_device_hardware_register_session_custom_proto_rawDescData
}

var edgelq_devices_proto_v1_device_hardware_register_session_custom_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var edgelq_devices_proto_v1_device_hardware_register_session_custom_proto_goTypes = []interface{}{
	(*RegisterHardwareRequest)(nil),                                        // 0: ntt.devices.v1.RegisterHardwareRequest
	(*RegisterHardwareResponse)(nil),                                       // 1: ntt.devices.v1.RegisterHardwareResponse
	(*HardwareStatusRequest)(nil),                                          // 2: ntt.devices.v1.HardwareStatusRequest
	(*HardwareStatusResponse)(nil),                                         // 3: ntt.devices.v1.HardwareStatusResponse
	(*GetDeviceHardwareRegisterSessionFromTokenRequest)(nil),               // 4: ntt.devices.v1.GetDeviceHardwareRegisterSessionFromTokenRequest
	(*GetDeviceHardwareRegisterSessionFromTokenResponse)(nil),              // 5: ntt.devices.v1.GetDeviceHardwareRegisterSessionFromTokenResponse
	(*device_hardware.DeviceHardware)(nil),                                 // 6: ntt.devices.v1.DeviceHardware
	(*device_hardware_register_session.DeviceHardwareRegisterSession)(nil), // 7: ntt.devices.v1.DeviceHardwareRegisterSession
}
var edgelq_devices_proto_v1_device_hardware_register_session_custom_proto_depIdxs = []int32{
	6, // 0: ntt.devices.v1.RegisterHardwareResponse.hardware:type_name -> ntt.devices.v1.DeviceHardware
	6, // 1: ntt.devices.v1.HardwareStatusResponse.device_hardwares:type_name -> ntt.devices.v1.DeviceHardware
	7, // 2: ntt.devices.v1.GetDeviceHardwareRegisterSessionFromTokenResponse.device_hardware_register_session:type_name -> ntt.devices.v1.DeviceHardwareRegisterSession
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { edgelq_devices_proto_v1_device_hardware_register_session_custom_proto_init() }
func edgelq_devices_proto_v1_device_hardware_register_session_custom_proto_init() {
	if edgelq_devices_proto_v1_device_hardware_register_session_custom_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		edgelq_devices_proto_v1_device_hardware_register_session_custom_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterHardwareRequest); i {
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
		edgelq_devices_proto_v1_device_hardware_register_session_custom_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterHardwareResponse); i {
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
		edgelq_devices_proto_v1_device_hardware_register_session_custom_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HardwareStatusRequest); i {
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
		edgelq_devices_proto_v1_device_hardware_register_session_custom_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HardwareStatusResponse); i {
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
		edgelq_devices_proto_v1_device_hardware_register_session_custom_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDeviceHardwareRegisterSessionFromTokenRequest); i {
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
		edgelq_devices_proto_v1_device_hardware_register_session_custom_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDeviceHardwareRegisterSessionFromTokenResponse); i {
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
			RawDescriptor: edgelq_devices_proto_v1_device_hardware_register_session_custom_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           edgelq_devices_proto_v1_device_hardware_register_session_custom_proto_goTypes,
		DependencyIndexes: edgelq_devices_proto_v1_device_hardware_register_session_custom_proto_depIdxs,
		MessageInfos:      edgelq_devices_proto_v1_device_hardware_register_session_custom_proto_msgTypes,
	}.Build()
	edgelq_devices_proto_v1_device_hardware_register_session_custom_proto = out.File
	edgelq_devices_proto_v1_device_hardware_register_session_custom_proto_rawDesc = nil
	edgelq_devices_proto_v1_device_hardware_register_session_custom_proto_goTypes = nil
	edgelq_devices_proto_v1_device_hardware_register_session_custom_proto_depIdxs = nil
}
