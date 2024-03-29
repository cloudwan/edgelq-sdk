// Code generated by protoc-gen-goten-go
// File: edgelq/ztp/proto/v1/edgelq_instance_custom.proto
// DO NOT EDIT!!!

package edgelq_instance_client

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
	edgelq_instance "github.com/cloudwan/edgelq-sdk/ztp/resources/v1/edgelq_instance"
	hardware "github.com/cloudwan/edgelq-sdk/ztp/resources/v1/hardware"
	project "github.com/cloudwan/edgelq-sdk/ztp/resources/v1/project"
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
	_ = &edgelq_instance.EdgelqInstance{}
	_ = &hardware.Hardware{}
	_ = &project.Project{}
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Request message for method [AssociateHardware][ntt.ztp.v1.AssociateHardware]
// only the fields serial_number and product_name are required for ztp function
// the other fields are present for debug/operations support
type AssociateHardwareRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Name of ntt.ztp.v1.EdgelqInstance
	Name                  *edgelq_instance.Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	SerialNumber          string                `protobuf:"bytes,2,opt,name=serial_number,json=serialNumber,proto3" json:"serial_number,omitempty" firestore:"serialNumber"`
	Manufacturer          string                `protobuf:"bytes,3,opt,name=manufacturer,proto3" json:"manufacturer,omitempty" firestore:"manufacturer"`
	ProductName           string                `protobuf:"bytes,4,opt,name=product_name,json=productName,proto3" json:"product_name,omitempty" firestore:"productName"`
	AssociatedProjectName string                `protobuf:"bytes,5,opt,name=associated_project_name,json=associatedProjectName,proto3" json:"associated_project_name,omitempty" firestore:"associatedProjectName"`
	// device could be associated to either provisioning policy or directly to
	// device resource only one will be set. This is not used for ztp itself, but
	// for debug purposes
	AssociatedProvisioningPolicyName string `protobuf:"bytes,6,opt,name=associated_provisioning_policy_name,json=associatedProvisioningPolicyName,proto3" json:"associated_provisioning_policy_name,omitempty" firestore:"associatedProvisioningPolicyName"`
	AssociatedDeviceName             string `protobuf:"bytes,7,opt,name=associated_device_name,json=associatedDeviceName,proto3" json:"associated_device_name,omitempty" firestore:"associatedDeviceName"`
	SimIccid                         string `protobuf:"bytes,8,opt,name=sim_iccid,json=simIccid,proto3" json:"sim_iccid,omitempty" firestore:"simIccid"`
	Imei                             string `protobuf:"bytes,9,opt,name=imei,proto3" json:"imei,omitempty" firestore:"imei"`
}

func (m *AssociateHardwareRequest) Reset() {
	*m = AssociateHardwareRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_ztp_proto_v1_edgelq_instance_custom_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *AssociateHardwareRequest) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*AssociateHardwareRequest) ProtoMessage() {}

func (m *AssociateHardwareRequest) ProtoReflect() preflect.Message {
	mi := &edgelq_ztp_proto_v1_edgelq_instance_custom_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*AssociateHardwareRequest) GotenMessage() {}

// Deprecated, Use AssociateHardwareRequest.ProtoReflect.Descriptor instead.
func (*AssociateHardwareRequest) Descriptor() ([]byte, []int) {
	return edgelq_ztp_proto_v1_edgelq_instance_custom_proto_rawDescGZIP(), []int{0}
}

func (m *AssociateHardwareRequest) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *AssociateHardwareRequest) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *AssociateHardwareRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *AssociateHardwareRequest) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *AssociateHardwareRequest) GetName() *edgelq_instance.Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *AssociateHardwareRequest) GetSerialNumber() string {
	if m != nil {
		return m.SerialNumber
	}
	return ""
}

func (m *AssociateHardwareRequest) GetManufacturer() string {
	if m != nil {
		return m.Manufacturer
	}
	return ""
}

func (m *AssociateHardwareRequest) GetProductName() string {
	if m != nil {
		return m.ProductName
	}
	return ""
}

func (m *AssociateHardwareRequest) GetAssociatedProjectName() string {
	if m != nil {
		return m.AssociatedProjectName
	}
	return ""
}

func (m *AssociateHardwareRequest) GetAssociatedProvisioningPolicyName() string {
	if m != nil {
		return m.AssociatedProvisioningPolicyName
	}
	return ""
}

func (m *AssociateHardwareRequest) GetAssociatedDeviceName() string {
	if m != nil {
		return m.AssociatedDeviceName
	}
	return ""
}

func (m *AssociateHardwareRequest) GetSimIccid() string {
	if m != nil {
		return m.SimIccid
	}
	return ""
}

func (m *AssociateHardwareRequest) GetImei() string {
	if m != nil {
		return m.Imei
	}
	return ""
}

func (m *AssociateHardwareRequest) SetName(fv *edgelq_instance.Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "AssociateHardwareRequest"))
	}
	m.Name = fv
}

func (m *AssociateHardwareRequest) SetSerialNumber(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "SerialNumber", "AssociateHardwareRequest"))
	}
	m.SerialNumber = fv
}

func (m *AssociateHardwareRequest) SetManufacturer(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Manufacturer", "AssociateHardwareRequest"))
	}
	m.Manufacturer = fv
}

func (m *AssociateHardwareRequest) SetProductName(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ProductName", "AssociateHardwareRequest"))
	}
	m.ProductName = fv
}

func (m *AssociateHardwareRequest) SetAssociatedProjectName(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "AssociatedProjectName", "AssociateHardwareRequest"))
	}
	m.AssociatedProjectName = fv
}

func (m *AssociateHardwareRequest) SetAssociatedProvisioningPolicyName(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "AssociatedProvisioningPolicyName", "AssociateHardwareRequest"))
	}
	m.AssociatedProvisioningPolicyName = fv
}

func (m *AssociateHardwareRequest) SetAssociatedDeviceName(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "AssociatedDeviceName", "AssociateHardwareRequest"))
	}
	m.AssociatedDeviceName = fv
}

func (m *AssociateHardwareRequest) SetSimIccid(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "SimIccid", "AssociateHardwareRequest"))
	}
	m.SimIccid = fv
}

func (m *AssociateHardwareRequest) SetImei(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Imei", "AssociateHardwareRequest"))
	}
	m.Imei = fv
}

// Response message for method [AssociateHardware][ntt.ztp.v1.AssociateHardware]
type AssociateHardwareResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Hardware      *hardware.Hardware `protobuf:"bytes,1,opt,name=hardware,proto3" json:"hardware,omitempty" firestore:"hardware"`
}

func (m *AssociateHardwareResponse) Reset() {
	*m = AssociateHardwareResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_ztp_proto_v1_edgelq_instance_custom_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *AssociateHardwareResponse) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*AssociateHardwareResponse) ProtoMessage() {}

func (m *AssociateHardwareResponse) ProtoReflect() preflect.Message {
	mi := &edgelq_ztp_proto_v1_edgelq_instance_custom_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*AssociateHardwareResponse) GotenMessage() {}

// Deprecated, Use AssociateHardwareResponse.ProtoReflect.Descriptor instead.
func (*AssociateHardwareResponse) Descriptor() ([]byte, []int) {
	return edgelq_ztp_proto_v1_edgelq_instance_custom_proto_rawDescGZIP(), []int{1}
}

func (m *AssociateHardwareResponse) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *AssociateHardwareResponse) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *AssociateHardwareResponse) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *AssociateHardwareResponse) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *AssociateHardwareResponse) GetHardware() *hardware.Hardware {
	if m != nil {
		return m.Hardware
	}
	return nil
}

func (m *AssociateHardwareResponse) SetHardware(fv *hardware.Hardware) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Hardware", "AssociateHardwareResponse"))
	}
	m.Hardware = fv
}

// Request message for method
// [DissociateHardware][ntt.ztp.v1.DissociateHardware]
type DissociateHardwareRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Name of ntt.ztp.v1.EdgelqInstance
	Name         *edgelq_instance.Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	SerialNumber string                `protobuf:"bytes,2,opt,name=serial_number,json=serialNumber,proto3" json:"serial_number,omitempty" firestore:"serialNumber"`
	Manufacturer string                `protobuf:"bytes,3,opt,name=manufacturer,proto3" json:"manufacturer,omitempty" firestore:"manufacturer"`
	ProductName  string                `protobuf:"bytes,4,opt,name=product_name,json=productName,proto3" json:"product_name,omitempty" firestore:"productName"`
}

func (m *DissociateHardwareRequest) Reset() {
	*m = DissociateHardwareRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_ztp_proto_v1_edgelq_instance_custom_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *DissociateHardwareRequest) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*DissociateHardwareRequest) ProtoMessage() {}

func (m *DissociateHardwareRequest) ProtoReflect() preflect.Message {
	mi := &edgelq_ztp_proto_v1_edgelq_instance_custom_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*DissociateHardwareRequest) GotenMessage() {}

// Deprecated, Use DissociateHardwareRequest.ProtoReflect.Descriptor instead.
func (*DissociateHardwareRequest) Descriptor() ([]byte, []int) {
	return edgelq_ztp_proto_v1_edgelq_instance_custom_proto_rawDescGZIP(), []int{2}
}

func (m *DissociateHardwareRequest) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *DissociateHardwareRequest) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *DissociateHardwareRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *DissociateHardwareRequest) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *DissociateHardwareRequest) GetName() *edgelq_instance.Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *DissociateHardwareRequest) GetSerialNumber() string {
	if m != nil {
		return m.SerialNumber
	}
	return ""
}

func (m *DissociateHardwareRequest) GetManufacturer() string {
	if m != nil {
		return m.Manufacturer
	}
	return ""
}

func (m *DissociateHardwareRequest) GetProductName() string {
	if m != nil {
		return m.ProductName
	}
	return ""
}

func (m *DissociateHardwareRequest) SetName(fv *edgelq_instance.Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "DissociateHardwareRequest"))
	}
	m.Name = fv
}

func (m *DissociateHardwareRequest) SetSerialNumber(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "SerialNumber", "DissociateHardwareRequest"))
	}
	m.SerialNumber = fv
}

func (m *DissociateHardwareRequest) SetManufacturer(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Manufacturer", "DissociateHardwareRequest"))
	}
	m.Manufacturer = fv
}

func (m *DissociateHardwareRequest) SetProductName(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ProductName", "DissociateHardwareRequest"))
	}
	m.ProductName = fv
}

var edgelq_ztp_proto_v1_edgelq_instance_custom_proto preflect.FileDescriptor

var edgelq_ztp_proto_v1_edgelq_instance_custom_proto_rawDesc = []byte{
	0x0a, 0x30, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x7a, 0x74, 0x70, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x5f, 0x69, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x6e, 0x74, 0x74, 0x2e, 0x7a, 0x74, 0x70, 0x2e, 0x76, 0x31, 0x1a, 0x20,
	0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x24, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x7a,
	0x74, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x61, 0x72, 0x64,
	0x77, 0x61, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x65, 0x64, 0x67, 0x65,
	0x6c, 0x71, 0x2f, 0x7a, 0x74, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f,
	0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x30, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x7a, 0x74,
	0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c,
	0x71, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa0, 0x03, 0x0a, 0x18, 0x41, 0x73, 0x73, 0x6f,
	0x63, 0x69, 0x61, 0x74, 0x65, 0x48, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x16, 0xb2, 0xda, 0x21, 0x12, 0x0a, 0x10, 0x0a, 0x0e, 0x45, 0x64, 0x67, 0x65,
	0x6c, 0x71, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x23, 0x0a, 0x0d, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x0c, 0x6d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63,
	0x74, 0x75, 0x72, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6d, 0x61, 0x6e,
	0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x36, 0x0a, 0x17,
	0x61, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x61,
	0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x4d, 0x0a, 0x23, 0x61, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x5f,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x20, 0x61, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x65, 0x64, 0x50, 0x72, 0x6f,
	0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x34, 0x0a, 0x16, 0x61, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x14, 0x61, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x65, 0x64, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x69, 0x6d,
	0x5f, 0x69, 0x63, 0x63, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x69,
	0x6d, 0x49, 0x63, 0x63, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6d, 0x65, 0x69, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6d, 0x65, 0x69, 0x22, 0x4d, 0x0a, 0x19, 0x41, 0x73,
	0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x65, 0x48, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x08, 0x68, 0x61, 0x72, 0x64, 0x77,
	0x61, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6e, 0x74, 0x74, 0x2e,
	0x7a, 0x74, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x52,
	0x08, 0x68, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x22, 0xb3, 0x01, 0x0a, 0x19, 0x44, 0x69,
	0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x65, 0x48, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x16, 0xb2, 0xda, 0x21, 0x12, 0x0a, 0x10, 0x0a, 0x0e, 0x45,
	0x64, 0x67, 0x65, 0x6c, 0x71, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x65, 0x72, 0x69,
	0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x0c, 0x6d, 0x61, 0x6e, 0x75,
	0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x6d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x42,
	0x85, 0x01, 0xe8, 0xde, 0x21, 0x00, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x74, 0x74, 0x2e,
	0x7a, 0x74, 0x70, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x42, 0x19, 0x45, 0x64, 0x67, 0x65, 0x6c,
	0x71, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x00, 0x5a, 0x4f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65,
	0x6c, 0x71, 0x2f, 0x7a, 0x74, 0x70, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31,
	0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x3b, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	edgelq_ztp_proto_v1_edgelq_instance_custom_proto_rawDescOnce sync.Once
	edgelq_ztp_proto_v1_edgelq_instance_custom_proto_rawDescData = edgelq_ztp_proto_v1_edgelq_instance_custom_proto_rawDesc
)

func edgelq_ztp_proto_v1_edgelq_instance_custom_proto_rawDescGZIP() []byte {
	edgelq_ztp_proto_v1_edgelq_instance_custom_proto_rawDescOnce.Do(func() {
		edgelq_ztp_proto_v1_edgelq_instance_custom_proto_rawDescData = protoimpl.X.CompressGZIP(edgelq_ztp_proto_v1_edgelq_instance_custom_proto_rawDescData)
	})
	return edgelq_ztp_proto_v1_edgelq_instance_custom_proto_rawDescData
}

var edgelq_ztp_proto_v1_edgelq_instance_custom_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var edgelq_ztp_proto_v1_edgelq_instance_custom_proto_goTypes = []interface{}{
	(*AssociateHardwareRequest)(nil),  // 0: ntt.ztp.v1.AssociateHardwareRequest
	(*AssociateHardwareResponse)(nil), // 1: ntt.ztp.v1.AssociateHardwareResponse
	(*DissociateHardwareRequest)(nil), // 2: ntt.ztp.v1.DissociateHardwareRequest
	(*hardware.Hardware)(nil),         // 3: ntt.ztp.v1.Hardware
}
var edgelq_ztp_proto_v1_edgelq_instance_custom_proto_depIdxs = []int32{
	3, // 0: ntt.ztp.v1.AssociateHardwareResponse.hardware:type_name -> ntt.ztp.v1.Hardware
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { edgelq_ztp_proto_v1_edgelq_instance_custom_proto_init() }
func edgelq_ztp_proto_v1_edgelq_instance_custom_proto_init() {
	if edgelq_ztp_proto_v1_edgelq_instance_custom_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		edgelq_ztp_proto_v1_edgelq_instance_custom_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssociateHardwareRequest); i {
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
		edgelq_ztp_proto_v1_edgelq_instance_custom_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssociateHardwareResponse); i {
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
		edgelq_ztp_proto_v1_edgelq_instance_custom_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DissociateHardwareRequest); i {
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
			RawDescriptor: edgelq_ztp_proto_v1_edgelq_instance_custom_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           edgelq_ztp_proto_v1_edgelq_instance_custom_proto_goTypes,
		DependencyIndexes: edgelq_ztp_proto_v1_edgelq_instance_custom_proto_depIdxs,
		MessageInfos:      edgelq_ztp_proto_v1_edgelq_instance_custom_proto_msgTypes,
	}.Build()
	edgelq_ztp_proto_v1_edgelq_instance_custom_proto = out.File
	edgelq_ztp_proto_v1_edgelq_instance_custom_proto_rawDesc = nil
	edgelq_ztp_proto_v1_edgelq_instance_custom_proto_goTypes = nil
	edgelq_ztp_proto_v1_edgelq_instance_custom_proto_depIdxs = nil
}
