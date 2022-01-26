// Code generated by protoc-gen-goten-go
// File: edgelq/iam/proto/v1alpha2/service_account.proto
// DO NOT EDIT!!!

package service_account

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
	project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/project"
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
	_ = &project.Project{}
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ServiceAccount Resource
type ServiceAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Name of ServiceAccount
	Name *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// Display Name
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty" firestore:"displayName"`
	// Email
	Email string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty" firestore:"email"`
	// Additional meta information about service account
	Metadata *ntt_meta.Meta `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty" firestore:"metadata"`
}

func (m *ServiceAccount) Reset() {
	*m = ServiceAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_iam_proto_v1alpha2_service_account_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *ServiceAccount) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*ServiceAccount) ProtoMessage() {}

func (m *ServiceAccount) ProtoReflect() preflect.Message {
	mi := &edgelq_iam_proto_v1alpha2_service_account_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*ServiceAccount) GotenMessage() {}

// Deprecated, Use ServiceAccount.ProtoReflect.Descriptor instead.
func (*ServiceAccount) Descriptor() ([]byte, []int) {
	return edgelq_iam_proto_v1alpha2_service_account_proto_rawDescGZIP(), []int{0}
}

func (m *ServiceAccount) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *ServiceAccount) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *ServiceAccount) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *ServiceAccount) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *ServiceAccount) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *ServiceAccount) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *ServiceAccount) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *ServiceAccount) GetMetadata() *ntt_meta.Meta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *ServiceAccount) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "ServiceAccount"))
	}
	m.Name = fv
}

func (m *ServiceAccount) SetDisplayName(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "DisplayName", "ServiceAccount"))
	}
	m.DisplayName = fv
}

func (m *ServiceAccount) SetEmail(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Email", "ServiceAccount"))
	}
	m.Email = fv
}

func (m *ServiceAccount) SetMetadata(fv *ntt_meta.Meta) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Metadata", "ServiceAccount"))
	}
	m.Metadata = fv
}

var edgelq_iam_proto_v1alpha2_service_account_proto preflect.FileDescriptor

var edgelq_iam_proto_v1alpha2_service_account_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x10, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x32, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x24, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x75, 0x64, 0x69,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x40, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f,
	0x6d, 0x65, 0x74, 0x61, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x65, 0x64, 0x67, 0x65, 0x6c,
	0x71, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6d,
	0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xda, 0x03, 0x0a, 0x0e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x16, 0xb2, 0xda, 0x21, 0x12,
	0x0a, 0x10, 0x0a, 0x0e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x29, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70,
	0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06,
	0xe2, 0xde, 0x21, 0x02, 0x08, 0x04, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x07, 0xe0, 0x41, 0x03, 0xf0, 0xd9, 0x21, 0x01, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x2b, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x3a,
	0xa4, 0x02, 0xea, 0x41, 0x66, 0x0a, 0x1d, 0x69, 0x61, 0x6d, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x6c,
	0x71, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x45, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x7b, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x7d, 0x92, 0xd9, 0x21, 0x74, 0x0a,
	0x0f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73,
	0x12, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x73, 0x1a, 0x07, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x4a, 0x23, 0x08, 0x02, 0x12, 0x06, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x0c,
	0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x07, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x6a, 0x1c, 0x6e, 0x74, 0x74, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e,
	0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2f, 0x52, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0xb2, 0xdf, 0x21, 0x0a, 0x0a, 0x08, 0x6c, 0x69, 0x73, 0x74, 0x56, 0x69, 0x65,
	0x77, 0xda, 0x94, 0x23, 0x08, 0x12, 0x06, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0xc2, 0x85, 0x2c,
	0x25, 0x22, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x42, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0xed, 0x03, 0xe8, 0xde, 0x21, 0x01, 0xd2, 0xff, 0xd0,
	0x02, 0x56, 0x0a, 0x15, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65,
	0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x74,
	0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x32, 0x42, 0x13, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x51, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64,
	0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x3b, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0xd2, 0x84, 0xd1, 0x02, 0x46,
	0x0a, 0x0d, 0x61, 0x75, 0x64, 0x69, 0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x12,
	0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f,
	0x61, 0x75, 0x64, 0x69, 0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0xf2, 0x85, 0xd1, 0x02, 0x5e, 0x0a, 0x19, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x64, 0x62, 0x5f,
	0x73, 0x79, 0x6e, 0x63, 0x65, 0x72, 0x12, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65,
	0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x64, 0x62, 0x5f, 0x73, 0x79, 0x6e, 0x63, 0x65, 0x72,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0xa2, 0x80, 0xd1, 0x02, 0x58, 0x0a, 0x16,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65,
	0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	edgelq_iam_proto_v1alpha2_service_account_proto_rawDescOnce sync.Once
	edgelq_iam_proto_v1alpha2_service_account_proto_rawDescData = edgelq_iam_proto_v1alpha2_service_account_proto_rawDesc
)

func edgelq_iam_proto_v1alpha2_service_account_proto_rawDescGZIP() []byte {
	edgelq_iam_proto_v1alpha2_service_account_proto_rawDescOnce.Do(func() {
		edgelq_iam_proto_v1alpha2_service_account_proto_rawDescData = protoimpl.X.CompressGZIP(edgelq_iam_proto_v1alpha2_service_account_proto_rawDescData)
	})
	return edgelq_iam_proto_v1alpha2_service_account_proto_rawDescData
}

var edgelq_iam_proto_v1alpha2_service_account_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var edgelq_iam_proto_v1alpha2_service_account_proto_goTypes = []interface{}{
	(*ServiceAccount)(nil), // 0: ntt.iam.v1alpha2.ServiceAccount
	(*ntt_meta.Meta)(nil),  // 1: ntt.types.Meta
}
var edgelq_iam_proto_v1alpha2_service_account_proto_depIdxs = []int32{
	1, // 0: ntt.iam.v1alpha2.ServiceAccount.metadata:type_name -> ntt.types.Meta
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { edgelq_iam_proto_v1alpha2_service_account_proto_init() }
func edgelq_iam_proto_v1alpha2_service_account_proto_init() {
	if edgelq_iam_proto_v1alpha2_service_account_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		edgelq_iam_proto_v1alpha2_service_account_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceAccount); i {
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
			RawDescriptor: edgelq_iam_proto_v1alpha2_service_account_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           edgelq_iam_proto_v1alpha2_service_account_proto_goTypes,
		DependencyIndexes: edgelq_iam_proto_v1alpha2_service_account_proto_depIdxs,
		MessageInfos:      edgelq_iam_proto_v1alpha2_service_account_proto_msgTypes,
	}.Build()
	edgelq_iam_proto_v1alpha2_service_account_proto = out.File
	edgelq_iam_proto_v1alpha2_service_account_proto_rawDesc = nil
	edgelq_iam_proto_v1alpha2_service_account_proto_goTypes = nil
	edgelq_iam_proto_v1alpha2_service_account_proto_depIdxs = nil
}
