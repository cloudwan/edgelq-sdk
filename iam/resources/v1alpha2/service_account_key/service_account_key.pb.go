// Code generated by protoc-gen-goten-go
// File: edgelq/iam/proto/v1alpha2/service_account_key.proto
// DO NOT EDIT!!!

package service_account_key

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
	service_account "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/service_account"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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
	_ = &service_account.ServiceAccount{}
	_ = &timestamp.Timestamp{}
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ServiceAccountKey_Algorithm int32

const (
	ServiceAccountKey_KEY_ALGORITHM_UNSPECIFIED ServiceAccountKey_Algorithm = 0
	ServiceAccountKey_RSA_1024                  ServiceAccountKey_Algorithm = 1
	ServiceAccountKey_RSA_2048                  ServiceAccountKey_Algorithm = 2
	ServiceAccountKey_RSA_4096                  ServiceAccountKey_Algorithm = 3
	ServiceAccountKey_API_KEY                   ServiceAccountKey_Algorithm = 4
)

var (
	ServiceAccountKey_Algorithm_name = map[int32]string{
		0: "KEY_ALGORITHM_UNSPECIFIED",
		1: "RSA_1024",
		2: "RSA_2048",
		3: "RSA_4096",
		4: "API_KEY",
	}

	ServiceAccountKey_Algorithm_value = map[string]int32{
		"KEY_ALGORITHM_UNSPECIFIED": 0,
		"RSA_1024":                  1,
		"RSA_2048":                  2,
		"RSA_4096":                  3,
		"API_KEY":                   4,
	}
)

func (x ServiceAccountKey_Algorithm) Enum() *ServiceAccountKey_Algorithm {
	p := new(ServiceAccountKey_Algorithm)
	*p = x
	return p
}

func (x ServiceAccountKey_Algorithm) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), preflect.EnumNumber(x))
}

func (ServiceAccountKey_Algorithm) Descriptor() preflect.EnumDescriptor {
	return edgelq_iam_proto_v1alpha2_service_account_key_proto_enumTypes[0].Descriptor()
}

func (ServiceAccountKey_Algorithm) Type() preflect.EnumType {
	return &edgelq_iam_proto_v1alpha2_service_account_key_proto_enumTypes[0]
}

func (x ServiceAccountKey_Algorithm) Number() preflect.EnumNumber {
	return preflect.EnumNumber(x)
}

// Deprecated, Use ServiceAccountKey_Algorithm.ProtoReflect.Descriptor instead.
func (ServiceAccountKey_Algorithm) EnumDescriptor() ([]byte, []int) {
	return edgelq_iam_proto_v1alpha2_service_account_key_proto_rawDescGZIP(), []int{0, 0}
}

// Represents a key pair for the service account
type ServiceAccountKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Name of ServiceAccountKey
	Name *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// Display name
	DisplayName string `protobuf:"bytes,7,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty" firestore:"displayName"`
	// Public key contents
	PublicKeyData string `protobuf:"bytes,2,opt,name=public_key_data,json=publicKeyData,proto3" json:"public_key_data,omitempty" firestore:"publicKeyData"`
	// The private key of the pair. This field is only provided in
	// CreateServiceAccountKey responses. Private keys are NOT stored on the
	// server.
	PrivateKeyData string `protobuf:"bytes,3,opt,name=private_key_data,json=privateKeyData,proto3" json:"private_key_data,omitempty" firestore:"privateKeyData"`
	// Api key is set if algorithm is equal to API_KEY and provided in
	// CreateServiceAccountKey responses. They are not stored on the server,
	// so caller is obliged to remember its value. If lost, its is necessary to
	// create new key. Api key must be used in authorization header token when
	// making request:
	// - Authorization: "Bearer $API_KEY"
	// Example for curl:
	// $ curl -X GET -H "Authorization: Bearer $API_KEY" -s $URL
	ApiKey string `protobuf:"bytes,9,opt,name=api_key,json=apiKey,proto3" json:"api_key,omitempty" firestore:"apiKey"`
	// The algorithm used to generate the key.
	Algorithm ServiceAccountKey_Algorithm `protobuf:"varint,4,opt,name=algorithm,proto3,enum=ntt.iam.v1alpha2.ServiceAccountKey_Algorithm" json:"algorithm,omitempty" firestore:"algorithm"`
	// The key is not valid before this timestamp.
	ValidNotBefore *timestamp.Timestamp `protobuf:"bytes,5,opt,name=valid_not_before,json=validNotBefore,proto3" json:"valid_not_before,omitempty" firestore:"validNotBefore"`
	// The key is not valid after this timestamp.
	ValidNotAfter *timestamp.Timestamp `protobuf:"bytes,6,opt,name=valid_not_after,json=validNotAfter,proto3" json:"valid_not_after,omitempty" firestore:"validNotAfter"`
	// Metadata
	Metadata *ntt_meta.Meta `protobuf:"bytes,8,opt,name=metadata,proto3" json:"metadata,omitempty" firestore:"metadata"`
}

func (m *ServiceAccountKey) Reset() {
	*m = ServiceAccountKey{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_iam_proto_v1alpha2_service_account_key_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *ServiceAccountKey) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*ServiceAccountKey) ProtoMessage() {}

func (m *ServiceAccountKey) ProtoReflect() preflect.Message {
	mi := &edgelq_iam_proto_v1alpha2_service_account_key_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*ServiceAccountKey) GotenMessage() {}

// Deprecated, Use ServiceAccountKey.ProtoReflect.Descriptor instead.
func (*ServiceAccountKey) Descriptor() ([]byte, []int) {
	return edgelq_iam_proto_v1alpha2_service_account_key_proto_rawDescGZIP(), []int{0}
}

func (m *ServiceAccountKey) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *ServiceAccountKey) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *ServiceAccountKey) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *ServiceAccountKey) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *ServiceAccountKey) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *ServiceAccountKey) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *ServiceAccountKey) GetPublicKeyData() string {
	if m != nil {
		return m.PublicKeyData
	}
	return ""
}

func (m *ServiceAccountKey) GetPrivateKeyData() string {
	if m != nil {
		return m.PrivateKeyData
	}
	return ""
}

func (m *ServiceAccountKey) GetApiKey() string {
	if m != nil {
		return m.ApiKey
	}
	return ""
}

func (m *ServiceAccountKey) GetAlgorithm() ServiceAccountKey_Algorithm {
	if m != nil {
		return m.Algorithm
	}
	return ServiceAccountKey_KEY_ALGORITHM_UNSPECIFIED
}

func (m *ServiceAccountKey) GetValidNotBefore() *timestamp.Timestamp {
	if m != nil {
		return m.ValidNotBefore
	}
	return nil
}

func (m *ServiceAccountKey) GetValidNotAfter() *timestamp.Timestamp {
	if m != nil {
		return m.ValidNotAfter
	}
	return nil
}

func (m *ServiceAccountKey) GetMetadata() *ntt_meta.Meta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *ServiceAccountKey) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "ServiceAccountKey"))
	}
	m.Name = fv
}

func (m *ServiceAccountKey) SetDisplayName(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "DisplayName", "ServiceAccountKey"))
	}
	m.DisplayName = fv
}

func (m *ServiceAccountKey) SetPublicKeyData(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "PublicKeyData", "ServiceAccountKey"))
	}
	m.PublicKeyData = fv
}

func (m *ServiceAccountKey) SetPrivateKeyData(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "PrivateKeyData", "ServiceAccountKey"))
	}
	m.PrivateKeyData = fv
}

func (m *ServiceAccountKey) SetApiKey(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ApiKey", "ServiceAccountKey"))
	}
	m.ApiKey = fv
}

func (m *ServiceAccountKey) SetAlgorithm(fv ServiceAccountKey_Algorithm) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Algorithm", "ServiceAccountKey"))
	}
	m.Algorithm = fv
}

func (m *ServiceAccountKey) SetValidNotBefore(fv *timestamp.Timestamp) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ValidNotBefore", "ServiceAccountKey"))
	}
	m.ValidNotBefore = fv
}

func (m *ServiceAccountKey) SetValidNotAfter(fv *timestamp.Timestamp) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ValidNotAfter", "ServiceAccountKey"))
	}
	m.ValidNotAfter = fv
}

func (m *ServiceAccountKey) SetMetadata(fv *ntt_meta.Meta) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Metadata", "ServiceAccountKey"))
	}
	m.Metadata = fv
}

var edgelq_iam_proto_v1alpha2_service_account_key_proto preflect.FileDescriptor

var edgelq_iam_proto_v1alpha2_service_account_key_proto_rawDesc = []byte{
	0x0a, 0x33, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x61, 0x75,
	0x64, 0x69, 0x74, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x34, 0x65, 0x64, 0x67,
	0x65, 0x6c, 0x71, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65,
	0x72, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6d,
	0x75, 0x6c, 0x74, 0x69, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1e, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xda, 0x08, 0x0a, 0x11, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x5d, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x49, 0xb2, 0xda, 0x21, 0x15, 0x0a, 0x13, 0x0a, 0x11, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4b, 0x65, 0x79,
	0xe2, 0xde, 0x21, 0x2c, 0x2a, 0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2a, 0x20,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61,
	0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69,
	0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x2e, 0x0a, 0x10, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xf0, 0xd9, 0x21,
	0x01, 0x52, 0x0e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x1d, 0x0a, 0x07, 0x61, 0x70, 0x69, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x04, 0xf0, 0xd9, 0x21, 0x01, 0x52, 0x06, 0x61, 0x70, 0x69, 0x4b, 0x65, 0x79,
	0x12, 0x4b, 0x0a, 0x09, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x2d, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4b, 0x65, 0x79, 0x2e, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74,
	0x68, 0x6d, 0x52, 0x09, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x12, 0x5a, 0x0a,
	0x10, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x6e, 0x6f, 0x74, 0x5f, 0x62, 0x65, 0x66, 0x6f, 0x72,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x42, 0x14, 0xf0, 0xd9, 0x21, 0x01, 0xe2, 0xde, 0x21, 0x0c, 0x08, 0x04, 0x2a,
	0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x52, 0x0e, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x4e, 0x6f, 0x74, 0x42, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x12, 0x58, 0x0a, 0x0f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x5f, 0x6e, 0x6f, 0x74, 0x5f, 0x61, 0x66, 0x74, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x14,
	0xf0, 0xd9, 0x21, 0x01, 0xe2, 0xde, 0x21, 0x0c, 0x08, 0x04, 0x2a, 0x08, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x52, 0x0d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x4e, 0x6f, 0x74, 0x41, 0x66,
	0x74, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x22, 0x61, 0x0a, 0x09, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x12, 0x1d, 0x0a,
	0x19, 0x4b, 0x45, 0x59, 0x5f, 0x41, 0x4c, 0x47, 0x4f, 0x52, 0x49, 0x54, 0x48, 0x4d, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08,
	0x52, 0x53, 0x41, 0x5f, 0x31, 0x30, 0x32, 0x34, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x53,
	0x41, 0x5f, 0x32, 0x30, 0x34, 0x38, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x53, 0x41, 0x5f,
	0x34, 0x30, 0x39, 0x36, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x41, 0x50, 0x49, 0x5f, 0x4b, 0x45,
	0x59, 0x10, 0x04, 0x3a, 0xb8, 0x03, 0xea, 0x41, 0x92, 0x01, 0x0a, 0x20, 0x69, 0x61, 0x6d, 0x2e,
	0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x6e, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d,
	0x2f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x7d, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x73, 0x2f, 0x7b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x7d, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x4b, 0x65, 0x79, 0x73, 0x2f, 0x7b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x7d, 0x92, 0xd9, 0x21, 0x90,
	0x01, 0x0a, 0x12, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x4b, 0x65, 0x79, 0x73, 0x12, 0x12, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4b, 0x65, 0x79, 0x73, 0x1a, 0x0e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x4a,
	0x4e, 0x08, 0x02, 0x12, 0x06, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x0c, 0x64,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0b, 0x0a, 0x09, 0x61,
	0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x12, 0x12, 0x0a, 0x10, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x5f, 0x6e, 0x6f, 0x74, 0x5f, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x12, 0x11, 0x0a, 0x0f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x6e, 0x6f, 0x74, 0x5f, 0x61, 0x66, 0x74, 0x65, 0x72, 0x50,
	0x02, 0xb2, 0xdf, 0x21, 0x0a, 0x0a, 0x08, 0x6c, 0x69, 0x73, 0x74, 0x56, 0x69, 0x65, 0x77, 0xda,
	0x94, 0x23, 0x08, 0x12, 0x06, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0xc2, 0x85, 0x2c, 0x6f, 0x1a,
	0x10, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x0f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65,
	0x79, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x22, 0x09, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68,
	0x6d, 0x22, 0x10, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x6e, 0x6f, 0x74, 0x5f, 0x62, 0x65, 0x66,
	0x6f, 0x72, 0x65, 0x22, 0x0f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x6e, 0x6f, 0x74, 0x5f, 0x61,
	0x66, 0x74, 0x65, 0x72, 0x42, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0xa5,
	0x03, 0xe8, 0xde, 0x21, 0x01, 0xd2, 0xff, 0xd0, 0x02, 0x5e, 0x0a, 0x19, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x5f,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c,
	0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x32, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x74,
	0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x32, 0x42, 0x16, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x4b, 0x65, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x59, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e,
	0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6b,
	0x65, 0x79, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0xd2, 0x84, 0xd1, 0x02, 0x46, 0x0a, 0x0d, 0x61, 0x75, 0x64,
	0x69, 0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x12, 0x35, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f,
	0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74,
	0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x32, 0xa2, 0x80, 0xd1, 0x02, 0x60, 0x0a, 0x1a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x12, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69,
	0x61, 0x6d, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x32, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	edgelq_iam_proto_v1alpha2_service_account_key_proto_rawDescOnce sync.Once
	edgelq_iam_proto_v1alpha2_service_account_key_proto_rawDescData = edgelq_iam_proto_v1alpha2_service_account_key_proto_rawDesc
)

func edgelq_iam_proto_v1alpha2_service_account_key_proto_rawDescGZIP() []byte {
	edgelq_iam_proto_v1alpha2_service_account_key_proto_rawDescOnce.Do(func() {
		edgelq_iam_proto_v1alpha2_service_account_key_proto_rawDescData = protoimpl.X.CompressGZIP(edgelq_iam_proto_v1alpha2_service_account_key_proto_rawDescData)
	})
	return edgelq_iam_proto_v1alpha2_service_account_key_proto_rawDescData
}

var edgelq_iam_proto_v1alpha2_service_account_key_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var edgelq_iam_proto_v1alpha2_service_account_key_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var edgelq_iam_proto_v1alpha2_service_account_key_proto_goTypes = []interface{}{
	(ServiceAccountKey_Algorithm)(0), // 0: ntt.iam.v1alpha2.ServiceAccountKey_Algorithm
	(*ServiceAccountKey)(nil),        // 1: ntt.iam.v1alpha2.ServiceAccountKey
	(*timestamp.Timestamp)(nil),      // 2: google.protobuf.Timestamp
	(*ntt_meta.Meta)(nil),            // 3: ntt.types.Meta
}
var edgelq_iam_proto_v1alpha2_service_account_key_proto_depIdxs = []int32{
	0, // 0: ntt.iam.v1alpha2.ServiceAccountKey.algorithm:type_name -> ntt.iam.v1alpha2.ServiceAccountKey_Algorithm
	2, // 1: ntt.iam.v1alpha2.ServiceAccountKey.valid_not_before:type_name -> google.protobuf.Timestamp
	2, // 2: ntt.iam.v1alpha2.ServiceAccountKey.valid_not_after:type_name -> google.protobuf.Timestamp
	3, // 3: ntt.iam.v1alpha2.ServiceAccountKey.metadata:type_name -> ntt.types.Meta
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { edgelq_iam_proto_v1alpha2_service_account_key_proto_init() }
func edgelq_iam_proto_v1alpha2_service_account_key_proto_init() {
	if edgelq_iam_proto_v1alpha2_service_account_key_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		edgelq_iam_proto_v1alpha2_service_account_key_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceAccountKey); i {
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
			RawDescriptor: edgelq_iam_proto_v1alpha2_service_account_key_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           edgelq_iam_proto_v1alpha2_service_account_key_proto_goTypes,
		DependencyIndexes: edgelq_iam_proto_v1alpha2_service_account_key_proto_depIdxs,
		EnumInfos:         edgelq_iam_proto_v1alpha2_service_account_key_proto_enumTypes,
		MessageInfos:      edgelq_iam_proto_v1alpha2_service_account_key_proto_msgTypes,
	}.Build()
	edgelq_iam_proto_v1alpha2_service_account_key_proto = out.File
	edgelq_iam_proto_v1alpha2_service_account_key_proto_rawDesc = nil
	edgelq_iam_proto_v1alpha2_service_account_key_proto_goTypes = nil
	edgelq_iam_proto_v1alpha2_service_account_key_proto_depIdxs = nil
}
