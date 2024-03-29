// Code generated by protoc-gen-goten-go
// File: edgelq/secrets/proto/v1/secret.proto
// DO NOT EDIT!!!

package secret

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
	project "github.com/cloudwan/edgelq-sdk/secrets/resources/v1/project"
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
	_ = &project.Project{}
	_ = &meta.Meta{}
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Secret Resource
type Secret struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Name of Secret
	// When creating a new instance, this field is optional and if not provided,
	// it will be generated automatically. Last ID segment must conform to the
	// following regex: [a-z][a-z0-9\\-]{0,28}[a-z0-9]
	Name *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// Metadata is an object with information like create, update and delete time
	// (for async deleted resources), has user labels/annotations, sharding
	// information, multi-region syncing information and may have non-schema
	// owners (useful for taking ownership of resources belonging to lower level
	// services by higher ones).
	Metadata *meta.Meta `protobuf:"bytes,4,opt,name=metadata,proto3" json:"metadata,omitempty" firestore:"metadata"`
	// masked by read_checks. Database only.
	EncData []byte `protobuf:"bytes,2,opt,name=enc_data,json=encData,proto3" json:"enc_data,omitempty" firestore:"encData"`
	// Data to store as secret. Must be base64 encoded.
	Data map[string]string `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" firestore:"data"`
}

func (m *Secret) Reset() {
	*m = Secret{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_secrets_proto_v1_secret_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *Secret) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*Secret) ProtoMessage() {}

func (m *Secret) ProtoReflect() preflect.Message {
	mi := &edgelq_secrets_proto_v1_secret_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*Secret) GotenMessage() {}

// Deprecated, Use Secret.ProtoReflect.Descriptor instead.
func (*Secret) Descriptor() ([]byte, []int) {
	return edgelq_secrets_proto_v1_secret_proto_rawDescGZIP(), []int{0}
}

func (m *Secret) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *Secret) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *Secret) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *Secret) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *Secret) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *Secret) GetMetadata() *meta.Meta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Secret) GetEncData() []byte {
	if m != nil {
		return m.EncData
	}
	return nil
}

func (m *Secret) GetData() map[string]string {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Secret) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "Secret"))
	}
	m.Name = fv
}

func (m *Secret) SetMetadata(fv *meta.Meta) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Metadata", "Secret"))
	}
	m.Metadata = fv
}

func (m *Secret) SetEncData(fv []byte) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "EncData", "Secret"))
	}
	m.EncData = fv
}

func (m *Secret) SetData(fv map[string]string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Data", "Secret"))
	}
	m.Data = fv
}

var edgelq_secrets_proto_v1_secret_proto preflect.FileDescriptor

var edgelq_secrets_proto_v1_secret_proto_rawDesc = []byte{
	0x0a, 0x24, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x6e, 0x74, 0x74, 0x2e, 0x73, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f, 0x72, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x65, 0x64, 0x67, 0x65,
	0x6c, 0x71, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x20, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x16, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f,
	0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xeb, 0x03, 0x0a, 0x06, 0x53,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x22, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x0e, 0xb2, 0xda, 0x21, 0x0a, 0x0a, 0x08, 0x0a, 0x06, 0x53, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x08, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67, 0x6f,
	0x74, 0x65, 0x6e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x63, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x65, 0x6e, 0x63, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x34, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x37, 0x0a, 0x09, 0x44, 0x61, 0x74,
	0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x3a, 0x83, 0x02, 0xea, 0x41, 0x51, 0x0a, 0x19, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x73, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x12, 0x34, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x7b, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73,
	0x2f, 0x7b, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x7d, 0x92, 0xd9, 0x21, 0x37, 0x0a, 0x07, 0x73,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x12, 0x07, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x1a,
	0x07, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x18, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x52, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x38, 0x05, 0xda, 0x94, 0x23, 0x08, 0x12, 0x06, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0xca, 0xa3, 0x22, 0x40, 0x0a, 0x1f, 0x0a, 0x13, 0x6d, 0x61, 0x73, 0x6b, 0x5f, 0x65, 0x6e, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x12, 0x08, 0x65, 0x6e, 0x63,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x0a, 0x1d, 0x0a, 0x15, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73,
	0x2e, 0x73, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x04,
	0x64, 0x61, 0x74, 0x61, 0xc2, 0x85, 0x2c, 0x20, 0x1a, 0x08, 0x65, 0x6e, 0x63, 0x5f, 0x64, 0x61,
	0x74, 0x61, 0x1a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0xf9, 0x01, 0xe8, 0xde, 0x21, 0x01, 0xd2,
	0xff, 0xd0, 0x02, 0x42, 0x0a, 0x0c, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x12, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x73,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f,
	0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0xa2, 0x80, 0xd1, 0x02, 0x44, 0x0a, 0x0d, 0x73, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x33, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f,
	0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2f, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x0a,
	0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73,
	0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c,
	0x71, 0x2f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x3b, 0x73, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	edgelq_secrets_proto_v1_secret_proto_rawDescOnce sync.Once
	edgelq_secrets_proto_v1_secret_proto_rawDescData = edgelq_secrets_proto_v1_secret_proto_rawDesc
)

func edgelq_secrets_proto_v1_secret_proto_rawDescGZIP() []byte {
	edgelq_secrets_proto_v1_secret_proto_rawDescOnce.Do(func() {
		edgelq_secrets_proto_v1_secret_proto_rawDescData = protoimpl.X.CompressGZIP(edgelq_secrets_proto_v1_secret_proto_rawDescData)
	})
	return edgelq_secrets_proto_v1_secret_proto_rawDescData
}

var edgelq_secrets_proto_v1_secret_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var edgelq_secrets_proto_v1_secret_proto_goTypes = []interface{}{
	(*Secret)(nil),    // 0: ntt.secrets.v1.Secret
	nil,               // 1: ntt.secrets.v1.Secret.DataEntry
	(*meta.Meta)(nil), // 2: goten.types.Meta
}
var edgelq_secrets_proto_v1_secret_proto_depIdxs = []int32{
	2, // 0: ntt.secrets.v1.Secret.metadata:type_name -> goten.types.Meta
	1, // 1: ntt.secrets.v1.Secret.data:type_name -> ntt.secrets.v1.Secret.DataEntry
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { edgelq_secrets_proto_v1_secret_proto_init() }
func edgelq_secrets_proto_v1_secret_proto_init() {
	if edgelq_secrets_proto_v1_secret_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		edgelq_secrets_proto_v1_secret_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Secret); i {
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
			RawDescriptor: edgelq_secrets_proto_v1_secret_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           edgelq_secrets_proto_v1_secret_proto_goTypes,
		DependencyIndexes: edgelq_secrets_proto_v1_secret_proto_depIdxs,
		MessageInfos:      edgelq_secrets_proto_v1_secret_proto_msgTypes,
	}.Build()
	edgelq_secrets_proto_v1_secret_proto = out.File
	edgelq_secrets_proto_v1_secret_proto_rawDesc = nil
	edgelq_secrets_proto_v1_secret_proto_goTypes = nil
	edgelq_secrets_proto_v1_secret_proto_depIdxs = nil
}
