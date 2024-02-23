// Code generated by protoc-gen-goten-go
// File: edgelq/devices/proto/v1alpha2/os_version.proto
// DO NOT EDIT!!!

package os_version

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
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
	device_type "github.com/cloudwan/edgelq-sdk/devices/resources/v1alpha2/device_type"
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
	_ = &ntt_meta.Meta{}
	_ = &device_type.DeviceType{}
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type OsVersion_Channel int32

const (
	OsVersion_CHANNEL_UNSPECIFIED OsVersion_Channel = 0
	OsVersion_NORMAL              OsVersion_Channel = 1
	OsVersion_BETA                OsVersion_Channel = 2
)

var (
	OsVersion_Channel_name = map[int32]string{
		0: "CHANNEL_UNSPECIFIED",
		1: "NORMAL",
		2: "BETA",
	}

	OsVersion_Channel_value = map[string]int32{
		"CHANNEL_UNSPECIFIED": 0,
		"NORMAL":              1,
		"BETA":                2,
	}
)

func (x OsVersion_Channel) Enum() *OsVersion_Channel {
	p := new(OsVersion_Channel)
	*p = x
	return p
}

func (x OsVersion_Channel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), preflect.EnumNumber(x))
}

func (OsVersion_Channel) Descriptor() preflect.EnumDescriptor {
	return edgelq_devices_proto_v1alpha2_os_version_proto_enumTypes[0].Descriptor()
}

func (OsVersion_Channel) Type() preflect.EnumType {
	return &edgelq_devices_proto_v1alpha2_os_version_proto_enumTypes[0]
}

func (x OsVersion_Channel) Number() preflect.EnumNumber {
	return preflect.EnumNumber(x)
}

// Deprecated, Use OsVersion_Channel.ProtoReflect.Descriptor instead.
func (OsVersion_Channel) EnumDescriptor() ([]byte, []int) {
	return edgelq_devices_proto_v1alpha2_os_version_proto_rawDescGZIP(), []int{0, 0}
}

// OsVersion Resource
type OsVersion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Name of OsVersion
	// When creating a new instance, this field is optional and if not provided,
	// it will be generated automatically. Last ID segment must conform to the
	// following regex: [a-z][a-z0-9\\-]{0,28}[a-z0-9]
	Name     *Name          `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	Metadata *ntt_meta.Meta `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty" firestore:"metadata"`
	// The version name.
	Version string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty" firestore:"version"`
	// Supported device type.
	DeviceType *device_type.Reference `protobuf:"bytes,4,opt,customtype=Reference,name=device_type,json=deviceType,proto3" json:"device_type,omitempty" firestore:"deviceType"`
	// The minum previous OS version that is required to upgrade devices to the OS
	// version.
	MinimumPreviousVersion string            `protobuf:"bytes,5,opt,name=minimum_previous_version,json=minimumPreviousVersion,proto3" json:"minimum_previous_version,omitempty" firestore:"minimumPreviousVersion"`
	Channel                OsVersion_Channel `protobuf:"varint,6,opt,name=channel,proto3,enum=ntt.devices.v1alpha2.OsVersion_Channel" json:"channel,omitempty" firestore:"channel"`
}

func (m *OsVersion) Reset() {
	*m = OsVersion{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_devices_proto_v1alpha2_os_version_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *OsVersion) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*OsVersion) ProtoMessage() {}

func (m *OsVersion) ProtoReflect() preflect.Message {
	mi := &edgelq_devices_proto_v1alpha2_os_version_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*OsVersion) GotenMessage() {}

// Deprecated, Use OsVersion.ProtoReflect.Descriptor instead.
func (*OsVersion) Descriptor() ([]byte, []int) {
	return edgelq_devices_proto_v1alpha2_os_version_proto_rawDescGZIP(), []int{0}
}

func (m *OsVersion) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *OsVersion) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *OsVersion) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *OsVersion) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *OsVersion) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *OsVersion) GetMetadata() *ntt_meta.Meta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *OsVersion) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *OsVersion) GetDeviceType() *device_type.Reference {
	if m != nil {
		return m.DeviceType
	}
	return nil
}

func (m *OsVersion) GetMinimumPreviousVersion() string {
	if m != nil {
		return m.MinimumPreviousVersion
	}
	return ""
}

func (m *OsVersion) GetChannel() OsVersion_Channel {
	if m != nil {
		return m.Channel
	}
	return OsVersion_CHANNEL_UNSPECIFIED
}

func (m *OsVersion) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "OsVersion"))
	}
	m.Name = fv
}

func (m *OsVersion) SetMetadata(fv *ntt_meta.Meta) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Metadata", "OsVersion"))
	}
	m.Metadata = fv
}

func (m *OsVersion) SetVersion(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Version", "OsVersion"))
	}
	m.Version = fv
}

func (m *OsVersion) SetDeviceType(fv *device_type.Reference) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "DeviceType", "OsVersion"))
	}
	m.DeviceType = fv
}

func (m *OsVersion) SetMinimumPreviousVersion(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "MinimumPreviousVersion", "OsVersion"))
	}
	m.MinimumPreviousVersion = fv
}

func (m *OsVersion) SetChannel(fv OsVersion_Channel) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Channel", "OsVersion"))
	}
	m.Channel = fv
}

var edgelq_devices_proto_v1alpha2_os_version_proto preflect.FileDescriptor

var edgelq_devices_proto_v1alpha2_os_version_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f,
	0x6f, 0x73, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x14, 0x6e, 0x74, 0x74, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x61, 0x75,
	0x64, 0x69, 0x74, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x34, 0x65, 0x64, 0x67,
	0x65, 0x6c, 0x71, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65,
	0x72, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6d,
	0x75, 0x6c, 0x74, 0x69, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1e, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x94, 0x04, 0x0a, 0x09, 0x4f, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x25, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x11, 0xb2,
	0xda, 0x21, 0x0d, 0x0a, 0x0b, 0x0a, 0x09, 0x4f, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2b, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x48, 0x0a,
	0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x27, 0xb2, 0xda, 0x21, 0x23, 0x12, 0x21, 0x0a, 0x1f, 0x6e, 0x74, 0x74, 0x2e,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32,
	0x2f, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x38, 0x0a, 0x18, 0x6d, 0x69, 0x6e, 0x69, 0x6d,
	0x75, 0x6d, 0x5f, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x5f, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x6d, 0x69, 0x6e, 0x69, 0x6d,
	0x75, 0x6d, 0x50, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x41, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x27, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x4f, 0x73, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x07, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x22, 0x38, 0x0a, 0x07, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12,
	0x17, 0x0a, 0x13, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x4f, 0x52, 0x4d,
	0x41, 0x4c, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x42, 0x45, 0x54, 0x41, 0x10, 0x02, 0x3a, 0x97,
	0x01, 0xea, 0x41, 0x48, 0x0a, 0x1c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x65, 0x64,
	0x67, 0x65, 0x6c, 0x71, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4f, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x28, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x6f, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x7b, 0x6f, 0x73, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x7d, 0x92, 0xd9, 0x21, 0x3c,
	0x0a, 0x0a, 0x6f, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x0a, 0x6f, 0x73,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x6a, 0x1c,
	0x6e, 0x74, 0x74, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2f, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0xda, 0x94, 0x23, 0x08,
	0x12, 0x06, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0xfb, 0x02, 0xe8, 0xde, 0x21, 0x01, 0xd2,
	0xff, 0xd0, 0x02, 0x50, 0x0a, 0x10, 0x6f, 0x73, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65,
	0x6c, 0x71, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x6f, 0x73, 0x5f, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x32, 0x42, 0x0e, 0x4f, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x4b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x6f, 0x73, 0x5f, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x3b, 0x6f, 0x73, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0xd2, 0x84, 0xd1, 0x02, 0x4a, 0x0a, 0x0d, 0x61, 0x75, 0x64, 0x69, 0x74, 0x68, 0x61, 0x6e, 0x64,
	0x6c, 0x65, 0x72, 0x73, 0x12, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71,
	0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x68, 0x61,
	0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0xa2,
	0x80, 0xd1, 0x02, 0x52, 0x0a, 0x11, 0x6f, 0x73, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67,
	0x65, 0x6c, 0x71, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x6f, 0x73, 0x5f, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	edgelq_devices_proto_v1alpha2_os_version_proto_rawDescOnce sync.Once
	edgelq_devices_proto_v1alpha2_os_version_proto_rawDescData = edgelq_devices_proto_v1alpha2_os_version_proto_rawDesc
)

func edgelq_devices_proto_v1alpha2_os_version_proto_rawDescGZIP() []byte {
	edgelq_devices_proto_v1alpha2_os_version_proto_rawDescOnce.Do(func() {
		edgelq_devices_proto_v1alpha2_os_version_proto_rawDescData = protoimpl.X.CompressGZIP(edgelq_devices_proto_v1alpha2_os_version_proto_rawDescData)
	})
	return edgelq_devices_proto_v1alpha2_os_version_proto_rawDescData
}

var edgelq_devices_proto_v1alpha2_os_version_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var edgelq_devices_proto_v1alpha2_os_version_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var edgelq_devices_proto_v1alpha2_os_version_proto_goTypes = []interface{}{
	(OsVersion_Channel)(0), // 0: ntt.devices.v1alpha2.OsVersion_Channel
	(*OsVersion)(nil),      // 1: ntt.devices.v1alpha2.OsVersion
	(*ntt_meta.Meta)(nil),  // 2: ntt.types.Meta
}
var edgelq_devices_proto_v1alpha2_os_version_proto_depIdxs = []int32{
	2, // 0: ntt.devices.v1alpha2.OsVersion.metadata:type_name -> ntt.types.Meta
	0, // 1: ntt.devices.v1alpha2.OsVersion.channel:type_name -> ntt.devices.v1alpha2.OsVersion_Channel
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { edgelq_devices_proto_v1alpha2_os_version_proto_init() }
func edgelq_devices_proto_v1alpha2_os_version_proto_init() {
	if edgelq_devices_proto_v1alpha2_os_version_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		edgelq_devices_proto_v1alpha2_os_version_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OsVersion); i {
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
			RawDescriptor: edgelq_devices_proto_v1alpha2_os_version_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           edgelq_devices_proto_v1alpha2_os_version_proto_goTypes,
		DependencyIndexes: edgelq_devices_proto_v1alpha2_os_version_proto_depIdxs,
		EnumInfos:         edgelq_devices_proto_v1alpha2_os_version_proto_enumTypes,
		MessageInfos:      edgelq_devices_proto_v1alpha2_os_version_proto_msgTypes,
	}.Build()
	edgelq_devices_proto_v1alpha2_os_version_proto = out.File
	edgelq_devices_proto_v1alpha2_os_version_proto_rawDesc = nil
	edgelq_devices_proto_v1alpha2_os_version_proto_goTypes = nil
	edgelq_devices_proto_v1alpha2_os_version_proto_depIdxs = nil
}
