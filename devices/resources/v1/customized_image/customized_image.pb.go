// Code generated by protoc-gen-goten-go
// File: edgelq/devices/proto/v1/customized_image.proto
// DO NOT EDIT!!!

package customized_image

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
	os_version "github.com/cloudwan/edgelq-sdk/devices/resources/v1/os_version"
	project "github.com/cloudwan/edgelq-sdk/devices/resources/v1/project"
	meta "github.com/cloudwan/goten-sdk/types/meta"
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
	_ = &os_version.OsVersion{}
	_ = &project.Project{}
	_ = &meta.Meta{}
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CustomizedImage_Status_State int32

const (
	// Initial state
	CustomizedImage_Status_STATE_UNSPECIFIED CustomizedImage_Status_State = 0
	// Backend is processing the customization request.
	CustomizedImage_Status_IN_PROGRESS CustomizedImage_Status_State = 1
	// There was an error processing the customization request. Check the
	// parameters and try again.
	CustomizedImage_Status_ERROR CustomizedImage_Status_State = 2
	// The customized image is ready to download.
	CustomizedImage_Status_READY CustomizedImage_Status_State = 3
)

var (
	CustomizedImage_Status_State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "IN_PROGRESS",
		2: "ERROR",
		3: "READY",
	}

	CustomizedImage_Status_State_value = map[string]int32{
		"STATE_UNSPECIFIED": 0,
		"IN_PROGRESS":       1,
		"ERROR":             2,
		"READY":             3,
	}
)

func (x CustomizedImage_Status_State) Enum() *CustomizedImage_Status_State {
	p := new(CustomizedImage_Status_State)
	*p = x
	return p
}

func (x CustomizedImage_Status_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), preflect.EnumNumber(x))
}

func (CustomizedImage_Status_State) Descriptor() preflect.EnumDescriptor {
	return edgelq_devices_proto_v1_customized_image_proto_enumTypes[0].Descriptor()
}

func (CustomizedImage_Status_State) Type() preflect.EnumType {
	return &edgelq_devices_proto_v1_customized_image_proto_enumTypes[0]
}

func (x CustomizedImage_Status_State) Number() preflect.EnumNumber {
	return preflect.EnumNumber(x)
}

// Deprecated, Use CustomizedImage_Status_State.ProtoReflect.Descriptor instead.
func (CustomizedImage_Status_State) EnumDescriptor() ([]byte, []int) {
	return edgelq_devices_proto_v1_customized_image_proto_rawDescGZIP(), []int{0, 1, 0}
}

// CustomizedImage Resource
type CustomizedImage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Name of CustomizedImage
	// When creating a new instance, this field is optional and if not provided,
	// it will be generated automatically. Last ID segment must conform to the
	// following regex: [a-z][a-z0-9\\-]{0,28}[a-z0-9]
	Name *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// Metadata is an object with information like create, update and delete time
	// (for async deleted resources), has user labels/annotations, sharding
	// information, multi-region syncing information and may have non-schema
	// owners (useful for taking ownership of resources belonging to lower level
	// services by higher ones).
	Metadata *meta.Meta              `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty" firestore:"metadata"`
	Spec     *CustomizedImage_Spec   `protobuf:"bytes,3,opt,name=spec,proto3" json:"spec,omitempty" firestore:"spec"`
	Status   *CustomizedImage_Status `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty" firestore:"status"`
}

func (m *CustomizedImage) Reset() {
	*m = CustomizedImage{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_devices_proto_v1_customized_image_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *CustomizedImage) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*CustomizedImage) ProtoMessage() {}

func (m *CustomizedImage) ProtoReflect() preflect.Message {
	mi := &edgelq_devices_proto_v1_customized_image_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*CustomizedImage) GotenMessage() {}

// Deprecated, Use CustomizedImage.ProtoReflect.Descriptor instead.
func (*CustomizedImage) Descriptor() ([]byte, []int) {
	return edgelq_devices_proto_v1_customized_image_proto_rawDescGZIP(), []int{0}
}

func (m *CustomizedImage) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *CustomizedImage) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *CustomizedImage) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *CustomizedImage) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *CustomizedImage) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *CustomizedImage) GetMetadata() *meta.Meta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *CustomizedImage) GetSpec() *CustomizedImage_Spec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *CustomizedImage) GetStatus() *CustomizedImage_Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *CustomizedImage) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "CustomizedImage"))
	}
	m.Name = fv
}

func (m *CustomizedImage) SetMetadata(fv *meta.Meta) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Metadata", "CustomizedImage"))
	}
	m.Metadata = fv
}

func (m *CustomizedImage) SetSpec(fv *CustomizedImage_Spec) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Spec", "CustomizedImage"))
	}
	m.Spec = fv
}

func (m *CustomizedImage) SetStatus(fv *CustomizedImage_Status) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Status", "CustomizedImage"))
	}
	m.Status = fv
}

type CustomizedImage_Spec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Deprecated. Use os_version instead.
	// After a while, version and device_type are removed
	// and os_version becomes required
	Version              string                `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty" firestore:"version"`
	DeviceType           string                `protobuf:"bytes,2,opt,name=device_type,json=deviceType,proto3" json:"device_type,omitempty" firestore:"deviceType"`
	OsVersion            *os_version.Reference `protobuf:"bytes,14,opt,customtype=Reference,name=os_version,json=osVersion,proto3" json:"os_version,omitempty" firestore:"osVersion"`
	ProvisioningPolicy   string                `protobuf:"bytes,3,opt,name=provisioning_policy,json=provisioningPolicy,proto3" json:"provisioning_policy,omitempty" firestore:"provisioningPolicy"`
	InstallAiAccelerator bool                  `protobuf:"varint,4,opt,name=install_ai_accelerator,json=installAiAccelerator,proto3" json:"install_ai_accelerator,omitempty" firestore:"installAiAccelerator"`
	Password             string                `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty" firestore:"password"`
	Encryption           bool                  `protobuf:"varint,6,opt,name=encryption,proto3" json:"encryption,omitempty" firestore:"encryption"`
	EncryptionPassword   string                `protobuf:"bytes,7,opt,name=encryption_password,json=encryptionPassword,proto3" json:"encryption_password,omitempty" firestore:"encryptionPassword"`
	DiskMapping          string                `protobuf:"bytes,8,opt,name=disk_mapping,json=diskMapping,proto3" json:"disk_mapping,omitempty" firestore:"diskMapping"`
	NetworkAgent         string                `protobuf:"bytes,9,opt,name=network_agent,json=networkAgent,proto3" json:"network_agent,omitempty" firestore:"networkAgent"`
	Ntp                  string                `protobuf:"bytes,10,opt,name=ntp,proto3" json:"ntp,omitempty" firestore:"ntp"`
	HttpProxy            string                `protobuf:"bytes,11,opt,name=http_proxy,json=httpProxy,proto3" json:"http_proxy,omitempty" firestore:"httpProxy"`
	HttpsProxy           string                `protobuf:"bytes,12,opt,name=https_proxy,json=httpsProxy,proto3" json:"https_proxy,omitempty" firestore:"httpsProxy"`
	NoProxy              string                `protobuf:"bytes,13,opt,name=no_proxy,json=noProxy,proto3" json:"no_proxy,omitempty" firestore:"noProxy"`
}

func (m *CustomizedImage_Spec) Reset() {
	*m = CustomizedImage_Spec{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_devices_proto_v1_customized_image_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *CustomizedImage_Spec) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*CustomizedImage_Spec) ProtoMessage() {}

func (m *CustomizedImage_Spec) ProtoReflect() preflect.Message {
	mi := &edgelq_devices_proto_v1_customized_image_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*CustomizedImage_Spec) GotenMessage() {}

// Deprecated, Use CustomizedImage_Spec.ProtoReflect.Descriptor instead.
func (*CustomizedImage_Spec) Descriptor() ([]byte, []int) {
	return edgelq_devices_proto_v1_customized_image_proto_rawDescGZIP(), []int{0, 0}
}

func (m *CustomizedImage_Spec) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *CustomizedImage_Spec) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *CustomizedImage_Spec) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *CustomizedImage_Spec) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *CustomizedImage_Spec) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *CustomizedImage_Spec) GetDeviceType() string {
	if m != nil {
		return m.DeviceType
	}
	return ""
}

func (m *CustomizedImage_Spec) GetOsVersion() *os_version.Reference {
	if m != nil {
		return m.OsVersion
	}
	return nil
}

func (m *CustomizedImage_Spec) GetProvisioningPolicy() string {
	if m != nil {
		return m.ProvisioningPolicy
	}
	return ""
}

func (m *CustomizedImage_Spec) GetInstallAiAccelerator() bool {
	if m != nil {
		return m.InstallAiAccelerator
	}
	return false
}

func (m *CustomizedImage_Spec) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *CustomizedImage_Spec) GetEncryption() bool {
	if m != nil {
		return m.Encryption
	}
	return false
}

func (m *CustomizedImage_Spec) GetEncryptionPassword() string {
	if m != nil {
		return m.EncryptionPassword
	}
	return ""
}

func (m *CustomizedImage_Spec) GetDiskMapping() string {
	if m != nil {
		return m.DiskMapping
	}
	return ""
}

func (m *CustomizedImage_Spec) GetNetworkAgent() string {
	if m != nil {
		return m.NetworkAgent
	}
	return ""
}

func (m *CustomizedImage_Spec) GetNtp() string {
	if m != nil {
		return m.Ntp
	}
	return ""
}

func (m *CustomizedImage_Spec) GetHttpProxy() string {
	if m != nil {
		return m.HttpProxy
	}
	return ""
}

func (m *CustomizedImage_Spec) GetHttpsProxy() string {
	if m != nil {
		return m.HttpsProxy
	}
	return ""
}

func (m *CustomizedImage_Spec) GetNoProxy() string {
	if m != nil {
		return m.NoProxy
	}
	return ""
}

func (m *CustomizedImage_Spec) SetVersion(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Version", "CustomizedImage_Spec"))
	}
	m.Version = fv
}

func (m *CustomizedImage_Spec) SetDeviceType(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "DeviceType", "CustomizedImage_Spec"))
	}
	m.DeviceType = fv
}

func (m *CustomizedImage_Spec) SetOsVersion(fv *os_version.Reference) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "OsVersion", "CustomizedImage_Spec"))
	}
	m.OsVersion = fv
}

func (m *CustomizedImage_Spec) SetProvisioningPolicy(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ProvisioningPolicy", "CustomizedImage_Spec"))
	}
	m.ProvisioningPolicy = fv
}

func (m *CustomizedImage_Spec) SetInstallAiAccelerator(fv bool) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "InstallAiAccelerator", "CustomizedImage_Spec"))
	}
	m.InstallAiAccelerator = fv
}

func (m *CustomizedImage_Spec) SetPassword(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Password", "CustomizedImage_Spec"))
	}
	m.Password = fv
}

func (m *CustomizedImage_Spec) SetEncryption(fv bool) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Encryption", "CustomizedImage_Spec"))
	}
	m.Encryption = fv
}

func (m *CustomizedImage_Spec) SetEncryptionPassword(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "EncryptionPassword", "CustomizedImage_Spec"))
	}
	m.EncryptionPassword = fv
}

func (m *CustomizedImage_Spec) SetDiskMapping(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "DiskMapping", "CustomizedImage_Spec"))
	}
	m.DiskMapping = fv
}

func (m *CustomizedImage_Spec) SetNetworkAgent(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "NetworkAgent", "CustomizedImage_Spec"))
	}
	m.NetworkAgent = fv
}

func (m *CustomizedImage_Spec) SetNtp(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Ntp", "CustomizedImage_Spec"))
	}
	m.Ntp = fv
}

func (m *CustomizedImage_Spec) SetHttpProxy(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "HttpProxy", "CustomizedImage_Spec"))
	}
	m.HttpProxy = fv
}

func (m *CustomizedImage_Spec) SetHttpsProxy(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "HttpsProxy", "CustomizedImage_Spec"))
	}
	m.HttpsProxy = fv
}

func (m *CustomizedImage_Spec) SetNoProxy(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "NoProxy", "CustomizedImage_Spec"))
	}
	m.NoProxy = fv
}

type CustomizedImage_Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// State of the image.
	State CustomizedImage_Status_State `protobuf:"varint,1,opt,name=state,proto3,enum=ntt.devices.v1.CustomizedImage_Status_State" json:"state,omitempty" firestore:"state"`
	// Error log when state is ERROR.
	Log string `protobuf:"bytes,2,opt,name=log,proto3" json:"log,omitempty" firestore:"log"`
	// TODO hide this from client
	File string `protobuf:"bytes,5,opt,name=file,proto3" json:"file,omitempty" firestore:"file"`
}

func (m *CustomizedImage_Status) Reset() {
	*m = CustomizedImage_Status{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_devices_proto_v1_customized_image_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *CustomizedImage_Status) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*CustomizedImage_Status) ProtoMessage() {}

func (m *CustomizedImage_Status) ProtoReflect() preflect.Message {
	mi := &edgelq_devices_proto_v1_customized_image_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*CustomizedImage_Status) GotenMessage() {}

// Deprecated, Use CustomizedImage_Status.ProtoReflect.Descriptor instead.
func (*CustomizedImage_Status) Descriptor() ([]byte, []int) {
	return edgelq_devices_proto_v1_customized_image_proto_rawDescGZIP(), []int{0, 1}
}

func (m *CustomizedImage_Status) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *CustomizedImage_Status) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *CustomizedImage_Status) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *CustomizedImage_Status) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *CustomizedImage_Status) GetState() CustomizedImage_Status_State {
	if m != nil {
		return m.State
	}
	return CustomizedImage_Status_STATE_UNSPECIFIED
}

func (m *CustomizedImage_Status) GetLog() string {
	if m != nil {
		return m.Log
	}
	return ""
}

func (m *CustomizedImage_Status) GetFile() string {
	if m != nil {
		return m.File
	}
	return ""
}

func (m *CustomizedImage_Status) SetState(fv CustomizedImage_Status_State) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "State", "CustomizedImage_Status"))
	}
	m.State = fv
}

func (m *CustomizedImage_Status) SetLog(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Log", "CustomizedImage_Status"))
	}
	m.Log = fv
}

func (m *CustomizedImage_Status) SetFile(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "File", "CustomizedImage_Status"))
	}
	m.File = fv
}

var edgelq_devices_proto_v1_customized_image_proto preflect.FileDescriptor

var edgelq_devices_proto_v1_customized_image_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x69, 0x7a, 0x65, 0x64, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0e, 0x6e, 0x74, 0x74, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x74,
	0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67,
	0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f,
	0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d,
	0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24,
	0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x69, 0x61, 0x6d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x61,
	0x75, 0x64, 0x69, 0x74, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x67, 0x6f,
	0x74, 0x65, 0x6e, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd1, 0x09, 0x0a, 0x0f, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69,
	0x7a, 0x65, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x2b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x17, 0xb2, 0xda, 0x21, 0x13, 0x0a, 0x11, 0x0a, 0x0f,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x65, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x38, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x65, 0x64, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x2e, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12, 0x3e,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26,
	0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x65, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x1a, 0x86,
	0x04, 0x0a, 0x04, 0x53, 0x70, 0x65, 0x63, 0x12, 0x1c, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x02, 0x18, 0x01, 0x52, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x02, 0x18, 0x01, 0x52, 0x0a,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x32, 0x0a, 0x0a, 0x6f, 0x73,
	0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x42, 0x13,
	0xb2, 0xda, 0x21, 0x0f, 0x12, 0x0d, 0x0a, 0x09, 0x4f, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x10, 0x06, 0x52, 0x09, 0x6f, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2f,
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12,
	0x34, 0x0a, 0x16, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x5f, 0x61, 0x69, 0x5f, 0x61, 0x63,
	0x63, 0x65, 0x6c, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x14, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x41, 0x69, 0x41, 0x63, 0x63, 0x65, 0x6c, 0x65,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x2f, 0x0a, 0x13, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12,
	0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x6b, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69,
	0x6e, 0x67, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x6b, 0x4d, 0x61,
	0x70, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x23, 0x0a, 0x0d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x74,
	0x70, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6e, 0x74, 0x70, 0x12, 0x1d, 0x0a, 0x0a,
	0x68, 0x74, 0x74, 0x70, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x68, 0x74, 0x74, 0x70, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x68,
	0x74, 0x74, 0x70, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x68, 0x74, 0x74, 0x70, 0x73, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x12, 0x19, 0x0a, 0x08,
	0x6e, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6e, 0x6f, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x1a, 0xb9, 0x01, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x42, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x2c, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x65, 0x64, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6f, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6c, 0x6f, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x22, 0x45, 0x0a, 0x05,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b,
	0x49, 0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x47, 0x52, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x09, 0x0a,
	0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x52, 0x45, 0x41, 0x44,
	0x59, 0x10, 0x03, 0x3a, 0xa2, 0x02, 0xea, 0x41, 0x6d, 0x0a, 0x22, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x65, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x47, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x7d, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x7d, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x65, 0x64, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x65, 0x64, 0x5f,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x7d, 0x92, 0xd9, 0x21, 0x49, 0x0a, 0x10, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x69, 0x7a, 0x65, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x10, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x65, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x1a, 0x07,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x18, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x52, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x38, 0x05, 0xda, 0x94, 0x23, 0x08, 0x12, 0x06, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0xca,
	0xa3, 0x22, 0x33, 0x12, 0x31, 0x0a, 0x27, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x73, 0x2e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0xc2, 0x85, 0x2c, 0x1e, 0x22, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x04, 0x73, 0x70, 0x65, 0x63, 0x2a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0xbe, 0x02, 0xe8, 0xde, 0x21, 0x01, 0xd2,
	0xff, 0xd0, 0x02, 0x56, 0x0a, 0x16, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x65, 0x64,
	0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x3c, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61,
	0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x69, 0x7a, 0x65, 0x64, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0xa2, 0x80, 0xd1, 0x02, 0x58, 0x0a,
	0x17, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x65, 0x64, 0x5f, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64,
	0x67, 0x65, 0x6c, 0x71, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x65,
	0x64, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x74, 0x74,
	0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x42, 0x14,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x65, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x51, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65,
	0x6c, 0x71, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a,
	0x65, 0x64, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x3b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69,
	0x7a, 0x65, 0x64, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	edgelq_devices_proto_v1_customized_image_proto_rawDescOnce sync.Once
	edgelq_devices_proto_v1_customized_image_proto_rawDescData = edgelq_devices_proto_v1_customized_image_proto_rawDesc
)

func edgelq_devices_proto_v1_customized_image_proto_rawDescGZIP() []byte {
	edgelq_devices_proto_v1_customized_image_proto_rawDescOnce.Do(func() {
		edgelq_devices_proto_v1_customized_image_proto_rawDescData = protoimpl.X.CompressGZIP(edgelq_devices_proto_v1_customized_image_proto_rawDescData)
	})
	return edgelq_devices_proto_v1_customized_image_proto_rawDescData
}

var edgelq_devices_proto_v1_customized_image_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var edgelq_devices_proto_v1_customized_image_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var edgelq_devices_proto_v1_customized_image_proto_goTypes = []interface{}{
	(CustomizedImage_Status_State)(0), // 0: ntt.devices.v1.CustomizedImage_Status_State
	(*CustomizedImage)(nil),           // 1: ntt.devices.v1.CustomizedImage
	(*CustomizedImage_Spec)(nil),      // 2: ntt.devices.v1.CustomizedImage.Spec
	(*CustomizedImage_Status)(nil),    // 3: ntt.devices.v1.CustomizedImage.Status
	(*meta.Meta)(nil),                 // 4: goten.types.Meta
}
var edgelq_devices_proto_v1_customized_image_proto_depIdxs = []int32{
	4, // 0: ntt.devices.v1.CustomizedImage.metadata:type_name -> goten.types.Meta
	2, // 1: ntt.devices.v1.CustomizedImage.spec:type_name -> ntt.devices.v1.CustomizedImage.Spec
	3, // 2: ntt.devices.v1.CustomizedImage.status:type_name -> ntt.devices.v1.CustomizedImage.Status
	0, // 3: ntt.devices.v1.CustomizedImage.Status.state:type_name -> ntt.devices.v1.CustomizedImage_Status_State
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { edgelq_devices_proto_v1_customized_image_proto_init() }
func edgelq_devices_proto_v1_customized_image_proto_init() {
	if edgelq_devices_proto_v1_customized_image_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		edgelq_devices_proto_v1_customized_image_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomizedImage); i {
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
		edgelq_devices_proto_v1_customized_image_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomizedImage_Spec); i {
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
		edgelq_devices_proto_v1_customized_image_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomizedImage_Status); i {
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
			RawDescriptor: edgelq_devices_proto_v1_customized_image_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           edgelq_devices_proto_v1_customized_image_proto_goTypes,
		DependencyIndexes: edgelq_devices_proto_v1_customized_image_proto_depIdxs,
		EnumInfos:         edgelq_devices_proto_v1_customized_image_proto_enumTypes,
		MessageInfos:      edgelq_devices_proto_v1_customized_image_proto_msgTypes,
	}.Build()
	edgelq_devices_proto_v1_customized_image_proto = out.File
	edgelq_devices_proto_v1_customized_image_proto_rawDesc = nil
	edgelq_devices_proto_v1_customized_image_proto_goTypes = nil
	edgelq_devices_proto_v1_customized_image_proto_depIdxs = nil
}
