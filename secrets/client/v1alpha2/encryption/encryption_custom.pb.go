// Code generated by protoc-gen-goten-go
// File: edgelq/secrets/proto/v1alpha2/encryption_custom.proto
// DO NOT EDIT!!!

package encryption_client

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
import ()

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
var ()

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Request message for method [EncryptData][ntt.secrets.v1alpha2.EncryptData]
type EncryptDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	ProjectName   string `protobuf:"bytes,1,opt,name=project_name,json=projectName,proto3" json:"project_name,omitempty" firestore:"projectName"`
	Data          []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty" firestore:"data"`
}

func (m *EncryptDataRequest) Reset() {
	*m = EncryptDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_secrets_proto_v1alpha2_encryption_custom_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *EncryptDataRequest) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*EncryptDataRequest) ProtoMessage() {}

func (m *EncryptDataRequest) ProtoReflect() preflect.Message {
	mi := &edgelq_secrets_proto_v1alpha2_encryption_custom_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*EncryptDataRequest) GotenMessage() {}

// Deprecated, Use EncryptDataRequest.ProtoReflect.Descriptor instead.
func (*EncryptDataRequest) Descriptor() ([]byte, []int) {
	return edgelq_secrets_proto_v1alpha2_encryption_custom_proto_rawDescGZIP(), []int{0}
}

func (m *EncryptDataRequest) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *EncryptDataRequest) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *EncryptDataRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *EncryptDataRequest) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *EncryptDataRequest) GetProjectName() string {
	if m != nil {
		return m.ProjectName
	}
	return ""
}

func (m *EncryptDataRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *EncryptDataRequest) SetProjectName(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ProjectName", "EncryptDataRequest"))
	}
	m.ProjectName = fv
}

func (m *EncryptDataRequest) SetData(fv []byte) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Data", "EncryptDataRequest"))
	}
	m.Data = fv
}

// Response message for method [EncryptData][ntt.secrets.v1alpha2.EncryptData]
type EncryptDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Data          []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty" firestore:"data"`
}

func (m *EncryptDataResponse) Reset() {
	*m = EncryptDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_secrets_proto_v1alpha2_encryption_custom_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *EncryptDataResponse) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*EncryptDataResponse) ProtoMessage() {}

func (m *EncryptDataResponse) ProtoReflect() preflect.Message {
	mi := &edgelq_secrets_proto_v1alpha2_encryption_custom_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*EncryptDataResponse) GotenMessage() {}

// Deprecated, Use EncryptDataResponse.ProtoReflect.Descriptor instead.
func (*EncryptDataResponse) Descriptor() ([]byte, []int) {
	return edgelq_secrets_proto_v1alpha2_encryption_custom_proto_rawDescGZIP(), []int{1}
}

func (m *EncryptDataResponse) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *EncryptDataResponse) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *EncryptDataResponse) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *EncryptDataResponse) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *EncryptDataResponse) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *EncryptDataResponse) SetData(fv []byte) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Data", "EncryptDataResponse"))
	}
	m.Data = fv
}

// Request message for method [DecryptData][ntt.secrets.v1alpha2.DecryptData]
type DecryptDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	ProjectName   string `protobuf:"bytes,1,opt,name=project_name,json=projectName,proto3" json:"project_name,omitempty" firestore:"projectName"`
	Data          []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty" firestore:"data"`
}

func (m *DecryptDataRequest) Reset() {
	*m = DecryptDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_secrets_proto_v1alpha2_encryption_custom_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *DecryptDataRequest) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*DecryptDataRequest) ProtoMessage() {}

func (m *DecryptDataRequest) ProtoReflect() preflect.Message {
	mi := &edgelq_secrets_proto_v1alpha2_encryption_custom_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*DecryptDataRequest) GotenMessage() {}

// Deprecated, Use DecryptDataRequest.ProtoReflect.Descriptor instead.
func (*DecryptDataRequest) Descriptor() ([]byte, []int) {
	return edgelq_secrets_proto_v1alpha2_encryption_custom_proto_rawDescGZIP(), []int{2}
}

func (m *DecryptDataRequest) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *DecryptDataRequest) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *DecryptDataRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *DecryptDataRequest) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *DecryptDataRequest) GetProjectName() string {
	if m != nil {
		return m.ProjectName
	}
	return ""
}

func (m *DecryptDataRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *DecryptDataRequest) SetProjectName(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ProjectName", "DecryptDataRequest"))
	}
	m.ProjectName = fv
}

func (m *DecryptDataRequest) SetData(fv []byte) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Data", "DecryptDataRequest"))
	}
	m.Data = fv
}

// Response message for method [DecryptData][ntt.secrets.v1alpha2.DecryptData]
type DecryptDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Data          []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty" firestore:"data"`
}

func (m *DecryptDataResponse) Reset() {
	*m = DecryptDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_secrets_proto_v1alpha2_encryption_custom_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *DecryptDataResponse) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*DecryptDataResponse) ProtoMessage() {}

func (m *DecryptDataResponse) ProtoReflect() preflect.Message {
	mi := &edgelq_secrets_proto_v1alpha2_encryption_custom_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*DecryptDataResponse) GotenMessage() {}

// Deprecated, Use DecryptDataResponse.ProtoReflect.Descriptor instead.
func (*DecryptDataResponse) Descriptor() ([]byte, []int) {
	return edgelq_secrets_proto_v1alpha2_encryption_custom_proto_rawDescGZIP(), []int{3}
}

func (m *DecryptDataResponse) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *DecryptDataResponse) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *DecryptDataResponse) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *DecryptDataResponse) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *DecryptDataResponse) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *DecryptDataResponse) SetData(fv []byte) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Data", "DecryptDataResponse"))
	}
	m.Data = fv
}

var edgelq_secrets_proto_v1alpha2_encryption_custom_proto preflect.FileDescriptor

var edgelq_secrets_proto_v1alpha2_encryption_custom_proto_rawDesc = []byte{
	0x0a, 0x35, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f,
	0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x6e, 0x74, 0x74, 0x2e, 0x73, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x1a, 0x20, 0x67,
	0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b,
	0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x32, 0x65, 0x64, 0x67,
	0x65, 0x6c, 0x71, 0x2f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x24, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x40, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6d, 0x65,
	0x74, 0x61, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2f,
	0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f,
	0x69, 0x61, 0x6d, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x69, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x57, 0x0a, 0x12, 0x45, 0x6e, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x3a, 0x0a, 0xc2, 0x85, 0x2c, 0x06, 0x1a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x22, 0x35, 0x0a, 0x13, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x3a, 0x0a, 0xc2,
	0x85, 0x2c, 0x06, 0x1a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x57, 0x0a, 0x12, 0x44, 0x65, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x3a, 0x0a, 0xc2, 0x85, 0x2c, 0x06, 0x1a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x22, 0x35, 0x0a, 0x13, 0x44, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x3a, 0x0a, 0xc2,
	0x85, 0x2c, 0x06, 0x1a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0xda, 0x01, 0xe8, 0xde, 0x21, 0x01,
	0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x73, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x42, 0x15, 0x45,
	0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x00, 0x5a, 0x4f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65,
	0x6c, 0x71, 0x2f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x65, 0x6e, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x3b, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0xd2, 0x84, 0xd1, 0x02, 0x4a, 0x0a, 0x0d, 0x61, 0x75,
	0x64, 0x69, 0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x12, 0x39, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e,
	0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2f,
	0x61, 0x75, 0x64, 0x69, 0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	edgelq_secrets_proto_v1alpha2_encryption_custom_proto_rawDescOnce sync.Once
	edgelq_secrets_proto_v1alpha2_encryption_custom_proto_rawDescData = edgelq_secrets_proto_v1alpha2_encryption_custom_proto_rawDesc
)

func edgelq_secrets_proto_v1alpha2_encryption_custom_proto_rawDescGZIP() []byte {
	edgelq_secrets_proto_v1alpha2_encryption_custom_proto_rawDescOnce.Do(func() {
		edgelq_secrets_proto_v1alpha2_encryption_custom_proto_rawDescData = protoimpl.X.CompressGZIP(edgelq_secrets_proto_v1alpha2_encryption_custom_proto_rawDescData)
	})
	return edgelq_secrets_proto_v1alpha2_encryption_custom_proto_rawDescData
}

var edgelq_secrets_proto_v1alpha2_encryption_custom_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var edgelq_secrets_proto_v1alpha2_encryption_custom_proto_goTypes = []interface{}{
	(*EncryptDataRequest)(nil),  // 0: ntt.secrets.v1alpha2.EncryptDataRequest
	(*EncryptDataResponse)(nil), // 1: ntt.secrets.v1alpha2.EncryptDataResponse
	(*DecryptDataRequest)(nil),  // 2: ntt.secrets.v1alpha2.DecryptDataRequest
	(*DecryptDataResponse)(nil), // 3: ntt.secrets.v1alpha2.DecryptDataResponse
}
var edgelq_secrets_proto_v1alpha2_encryption_custom_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { edgelq_secrets_proto_v1alpha2_encryption_custom_proto_init() }
func edgelq_secrets_proto_v1alpha2_encryption_custom_proto_init() {
	if edgelq_secrets_proto_v1alpha2_encryption_custom_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		edgelq_secrets_proto_v1alpha2_encryption_custom_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EncryptDataRequest); i {
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
		edgelq_secrets_proto_v1alpha2_encryption_custom_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EncryptDataResponse); i {
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
		edgelq_secrets_proto_v1alpha2_encryption_custom_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DecryptDataRequest); i {
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
		edgelq_secrets_proto_v1alpha2_encryption_custom_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DecryptDataResponse); i {
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
			RawDescriptor: edgelq_secrets_proto_v1alpha2_encryption_custom_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           edgelq_secrets_proto_v1alpha2_encryption_custom_proto_goTypes,
		DependencyIndexes: edgelq_secrets_proto_v1alpha2_encryption_custom_proto_depIdxs,
		MessageInfos:      edgelq_secrets_proto_v1alpha2_encryption_custom_proto_msgTypes,
	}.Build()
	edgelq_secrets_proto_v1alpha2_encryption_custom_proto = out.File
	edgelq_secrets_proto_v1alpha2_encryption_custom_proto_rawDesc = nil
	edgelq_secrets_proto_v1alpha2_encryption_custom_proto_goTypes = nil
	edgelq_secrets_proto_v1alpha2_encryption_custom_proto_depIdxs = nil
}
