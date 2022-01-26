// Code generated by protoc-gen-goten-go
// File: edgelq/devices/proto/v1alpha2/device_custom.proto
// DO NOT EDIT!!!

package device_client

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
	api "github.com/cloudwan/edgelq-sdk/common/api"
	device "github.com/cloudwan/edgelq-sdk/devices/resources/v1alpha2/device"
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
	_ = &api.Account{}
	_ = &device.Device{}
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Request message for method
// [ProvisionServiceAccountToDevice][ntt.devices.v1alpha2.ProvisionServiceAccountToDevice]
type ProvisionServiceAccountToDeviceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Name of Device to provision service account.
	// This method automatically generates a service account and a key for the
	// device and assign the key to .Spec.ServiceAccount
	Name *device.Reference `protobuf:"bytes,1,opt,customtype=Reference,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// external public key to use for service account for this device.
	// If not present, a keypair will be generated by the server.
	ExternalPubkey string `protobuf:"bytes,2,opt,name=external_pubkey,json=externalPubkey,proto3" json:"external_pubkey,omitempty" firestore:"externalPubkey"`
}

func (m *ProvisionServiceAccountToDeviceRequest) Reset() {
	*m = ProvisionServiceAccountToDeviceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_devices_proto_v1alpha2_device_custom_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *ProvisionServiceAccountToDeviceRequest) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*ProvisionServiceAccountToDeviceRequest) ProtoMessage() {}

func (m *ProvisionServiceAccountToDeviceRequest) ProtoReflect() preflect.Message {
	mi := &edgelq_devices_proto_v1alpha2_device_custom_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*ProvisionServiceAccountToDeviceRequest) GotenMessage() {}

// Deprecated, Use ProvisionServiceAccountToDeviceRequest.ProtoReflect.Descriptor instead.
func (*ProvisionServiceAccountToDeviceRequest) Descriptor() ([]byte, []int) {
	return edgelq_devices_proto_v1alpha2_device_custom_proto_rawDescGZIP(), []int{0}
}

func (m *ProvisionServiceAccountToDeviceRequest) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *ProvisionServiceAccountToDeviceRequest) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *ProvisionServiceAccountToDeviceRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *ProvisionServiceAccountToDeviceRequest) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *ProvisionServiceAccountToDeviceRequest) GetName() *device.Reference {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *ProvisionServiceAccountToDeviceRequest) GetExternalPubkey() string {
	if m != nil {
		return m.ExternalPubkey
	}
	return ""
}

func (m *ProvisionServiceAccountToDeviceRequest) SetName(fv *device.Reference) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "ProvisionServiceAccountToDeviceRequest"))
	}
	m.Name = fv
}

func (m *ProvisionServiceAccountToDeviceRequest) SetExternalPubkey(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ExternalPubkey", "ProvisionServiceAccountToDeviceRequest"))
	}
	m.ExternalPubkey = fv
}

// Response message for method
// [ProvisionServiceAccountToDevice][ntt.devices.v1alpha2.ProvisionServiceAccountToDevice]
type ProvisionServiceAccountToDeviceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// The service account generated by the endpoint. Save the private key
	// included.
	ServiceAccount *api.ServiceAccount `protobuf:"bytes,1,opt,name=service_account,json=serviceAccount,proto3" json:"service_account,omitempty" firestore:"serviceAccount"`
}

func (m *ProvisionServiceAccountToDeviceResponse) Reset() {
	*m = ProvisionServiceAccountToDeviceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_devices_proto_v1alpha2_device_custom_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *ProvisionServiceAccountToDeviceResponse) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*ProvisionServiceAccountToDeviceResponse) ProtoMessage() {}

func (m *ProvisionServiceAccountToDeviceResponse) ProtoReflect() preflect.Message {
	mi := &edgelq_devices_proto_v1alpha2_device_custom_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*ProvisionServiceAccountToDeviceResponse) GotenMessage() {}

// Deprecated, Use ProvisionServiceAccountToDeviceResponse.ProtoReflect.Descriptor instead.
func (*ProvisionServiceAccountToDeviceResponse) Descriptor() ([]byte, []int) {
	return edgelq_devices_proto_v1alpha2_device_custom_proto_rawDescGZIP(), []int{1}
}

func (m *ProvisionServiceAccountToDeviceResponse) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *ProvisionServiceAccountToDeviceResponse) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *ProvisionServiceAccountToDeviceResponse) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *ProvisionServiceAccountToDeviceResponse) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *ProvisionServiceAccountToDeviceResponse) GetServiceAccount() *api.ServiceAccount {
	if m != nil {
		return m.ServiceAccount
	}
	return nil
}

func (m *ProvisionServiceAccountToDeviceResponse) SetServiceAccount(fv *api.ServiceAccount) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ServiceAccount", "ProvisionServiceAccountToDeviceResponse"))
	}
	m.ServiceAccount = fv
}

// Request message for method
// [RemoveServiceAccountFromDevice][ntt.devices.v1alpha2.RemoveServiceAccountFromDevice]
type RemoveServiceAccountFromDeviceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Name of the Device to remove its service account
	// The service account and the key will be automatically deleted as well when
	// they have been generated automatically by `ProvisionServiceAccountToDevice`
	// action. Otherwise, the service account will be kept intact. The device will
	// lost the access to the resources in the project.
	Name *device.Reference `protobuf:"bytes,1,opt,customtype=Reference,name=name,proto3" json:"name,omitempty" firestore:"name"`
}

func (m *RemoveServiceAccountFromDeviceRequest) Reset() {
	*m = RemoveServiceAccountFromDeviceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_devices_proto_v1alpha2_device_custom_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *RemoveServiceAccountFromDeviceRequest) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*RemoveServiceAccountFromDeviceRequest) ProtoMessage() {}

func (m *RemoveServiceAccountFromDeviceRequest) ProtoReflect() preflect.Message {
	mi := &edgelq_devices_proto_v1alpha2_device_custom_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*RemoveServiceAccountFromDeviceRequest) GotenMessage() {}

// Deprecated, Use RemoveServiceAccountFromDeviceRequest.ProtoReflect.Descriptor instead.
func (*RemoveServiceAccountFromDeviceRequest) Descriptor() ([]byte, []int) {
	return edgelq_devices_proto_v1alpha2_device_custom_proto_rawDescGZIP(), []int{2}
}

func (m *RemoveServiceAccountFromDeviceRequest) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *RemoveServiceAccountFromDeviceRequest) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *RemoveServiceAccountFromDeviceRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *RemoveServiceAccountFromDeviceRequest) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *RemoveServiceAccountFromDeviceRequest) GetName() *device.Reference {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *RemoveServiceAccountFromDeviceRequest) SetName(fv *device.Reference) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "RemoveServiceAccountFromDeviceRequest"))
	}
	m.Name = fv
}

// Response message for method
// [RemoveServiceAccountFromDevice][ntt.devices.v1alpha2.RemoveServiceAccountFromDevice]
type RemoveServiceAccountFromDeviceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (m *RemoveServiceAccountFromDeviceResponse) Reset() {
	*m = RemoveServiceAccountFromDeviceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_devices_proto_v1alpha2_device_custom_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *RemoveServiceAccountFromDeviceResponse) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*RemoveServiceAccountFromDeviceResponse) ProtoMessage() {}

func (m *RemoveServiceAccountFromDeviceResponse) ProtoReflect() preflect.Message {
	mi := &edgelq_devices_proto_v1alpha2_device_custom_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*RemoveServiceAccountFromDeviceResponse) GotenMessage() {}

// Deprecated, Use RemoveServiceAccountFromDeviceResponse.ProtoReflect.Descriptor instead.
func (*RemoveServiceAccountFromDeviceResponse) Descriptor() ([]byte, []int) {
	return edgelq_devices_proto_v1alpha2_device_custom_proto_rawDescGZIP(), []int{3}
}

func (m *RemoveServiceAccountFromDeviceResponse) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *RemoveServiceAccountFromDeviceResponse) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *RemoveServiceAccountFromDeviceResponse) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *RemoveServiceAccountFromDeviceResponse) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

var edgelq_devices_proto_v1alpha2_device_custom_proto preflect.FileDescriptor

var edgelq_devices_proto_v1alpha2_device_custom_proto_rawDesc = []byte{
	0x0a, 0x31, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x14, 0x6e, 0x74, 0x74, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74,
	0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67,
	0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x65, 0x64, 0x67, 0x65,
	0x6c, 0x71, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x31, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x65, 0x64, 0x67, 0x65, 0x6c,
	0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x65, 0x64, 0x67,
	0x65, 0x6c, 0x71, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x40, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x6d,
	0x75, 0x6c, 0x74, 0x69, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x70,
	0x69, 0x6c, 0x65, 0x72, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61,
	0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8b, 0x01, 0x0a, 0x26, 0x50, 0x72, 0x6f,
	0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x54, 0x6f, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x12, 0xb2, 0xda, 0x21, 0x0a, 0x12, 0x08, 0x0a, 0x06, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0xba, 0x9d, 0x22, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x65,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x70, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x50, 0x75,
	0x62, 0x6b, 0x65, 0x79, 0x3a, 0x10, 0xc2, 0x85, 0x2c, 0x0c, 0x32, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x3a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x8e, 0x01, 0x0a, 0x27, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x54, 0x6f, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x40, 0x0a, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6e, 0x74,
	0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x52, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x3a, 0x21, 0xc2, 0x85, 0x2c, 0x1d, 0x1a, 0x1b, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x22, 0x61, 0x0a, 0x25, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x46,
	0x72, 0x6f, 0x6d, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x26, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x12,
	0xb2, 0xda, 0x21, 0x0a, 0x12, 0x08, 0x0a, 0x06, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0xba, 0x9d,
	0x22, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x3a, 0x10, 0xc2, 0x85, 0x2c, 0x0c, 0x32, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x3a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x28, 0x0a, 0x26, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0xce, 0x01, 0xe8, 0xde, 0x21, 0x01, 0x0a, 0x1b, 0x63, 0x6f, 0x6d,
	0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x62, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x42, 0x11, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x00, 0x5a, 0x47, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77,
	0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x32, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x3b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0xd2, 0x84, 0xd1, 0x02, 0x4a, 0x0a, 0x0d, 0x61, 0x75, 0x64,
	0x69, 0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x12, 0x39, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f,
	0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x61,
	0x75, 0x64, 0x69, 0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	edgelq_devices_proto_v1alpha2_device_custom_proto_rawDescOnce sync.Once
	edgelq_devices_proto_v1alpha2_device_custom_proto_rawDescData = edgelq_devices_proto_v1alpha2_device_custom_proto_rawDesc
)

func edgelq_devices_proto_v1alpha2_device_custom_proto_rawDescGZIP() []byte {
	edgelq_devices_proto_v1alpha2_device_custom_proto_rawDescOnce.Do(func() {
		edgelq_devices_proto_v1alpha2_device_custom_proto_rawDescData = protoimpl.X.CompressGZIP(edgelq_devices_proto_v1alpha2_device_custom_proto_rawDescData)
	})
	return edgelq_devices_proto_v1alpha2_device_custom_proto_rawDescData
}

var edgelq_devices_proto_v1alpha2_device_custom_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var edgelq_devices_proto_v1alpha2_device_custom_proto_goTypes = []interface{}{
	(*ProvisionServiceAccountToDeviceRequest)(nil),  // 0: ntt.devices.v1alpha2.ProvisionServiceAccountToDeviceRequest
	(*ProvisionServiceAccountToDeviceResponse)(nil), // 1: ntt.devices.v1alpha2.ProvisionServiceAccountToDeviceResponse
	(*RemoveServiceAccountFromDeviceRequest)(nil),   // 2: ntt.devices.v1alpha2.RemoveServiceAccountFromDeviceRequest
	(*RemoveServiceAccountFromDeviceResponse)(nil),  // 3: ntt.devices.v1alpha2.RemoveServiceAccountFromDeviceResponse
	(*api.ServiceAccount)(nil),                      // 4: ntt.api.ServiceAccount
}
var edgelq_devices_proto_v1alpha2_device_custom_proto_depIdxs = []int32{
	4, // 0: ntt.devices.v1alpha2.ProvisionServiceAccountToDeviceResponse.service_account:type_name -> ntt.api.ServiceAccount
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { edgelq_devices_proto_v1alpha2_device_custom_proto_init() }
func edgelq_devices_proto_v1alpha2_device_custom_proto_init() {
	if edgelq_devices_proto_v1alpha2_device_custom_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		edgelq_devices_proto_v1alpha2_device_custom_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProvisionServiceAccountToDeviceRequest); i {
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
		edgelq_devices_proto_v1alpha2_device_custom_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProvisionServiceAccountToDeviceResponse); i {
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
		edgelq_devices_proto_v1alpha2_device_custom_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveServiceAccountFromDeviceRequest); i {
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
		edgelq_devices_proto_v1alpha2_device_custom_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveServiceAccountFromDeviceResponse); i {
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
			RawDescriptor: edgelq_devices_proto_v1alpha2_device_custom_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           edgelq_devices_proto_v1alpha2_device_custom_proto_goTypes,
		DependencyIndexes: edgelq_devices_proto_v1alpha2_device_custom_proto_depIdxs,
		MessageInfos:      edgelq_devices_proto_v1alpha2_device_custom_proto_msgTypes,
	}.Build()
	edgelq_devices_proto_v1alpha2_device_custom_proto = out.File
	edgelq_devices_proto_v1alpha2_device_custom_proto_rawDesc = nil
	edgelq_devices_proto_v1alpha2_device_custom_proto_goTypes = nil
	edgelq_devices_proto_v1alpha2_device_custom_proto_depIdxs = nil
}
