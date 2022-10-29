// Code generated by protoc-gen-goten-go
// File: edgelq/iam/proto/v1alpha2/user.proto
// DO NOT EDIT!!!

package user

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
	_ = &timestamp.Timestamp{}
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// User Resource
type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Name of User
	Name *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// Full Name
	FullName string `protobuf:"bytes,2,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty" firestore:"fullName"`
	// Metadata
	Metadata *ntt_meta.Meta `protobuf:"bytes,8,opt,name=metadata,proto3" json:"metadata,omitempty" firestore:"metadata"`
	// Email
	Email string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty" firestore:"email"`
	// Is email verified
	EmailVerified bool           `protobuf:"varint,4,opt,name=email_verified,json=emailVerified,proto3" json:"email_verified,omitempty" firestore:"emailVerified"`
	AuthInfo      *User_AuthInfo `protobuf:"bytes,5,opt,name=auth_info,json=authInfo,proto3" json:"auth_info,omitempty" firestore:"authInfo"`
	// User settings and preferences
	Settings      map[string]string    `protobuf:"bytes,7,rep,name=settings,proto3" json:"settings,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" firestore:"settings"`
	RefreshedTime *timestamp.Timestamp `protobuf:"bytes,6,opt,name=refreshed_time,json=refreshedTime,proto3" json:"refreshed_time,omitempty" firestore:"refreshedTime"`
}

func (m *User) Reset() {
	*m = User{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_iam_proto_v1alpha2_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *User) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*User) ProtoMessage() {}

func (m *User) ProtoReflect() preflect.Message {
	mi := &edgelq_iam_proto_v1alpha2_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*User) GotenMessage() {}

// Deprecated, Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return edgelq_iam_proto_v1alpha2_user_proto_rawDescGZIP(), []int{0}
}

func (m *User) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *User) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *User) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *User) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *User) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *User) GetFullName() string {
	if m != nil {
		return m.FullName
	}
	return ""
}

func (m *User) GetMetadata() *ntt_meta.Meta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetEmailVerified() bool {
	if m != nil {
		return m.EmailVerified
	}
	return false
}

func (m *User) GetAuthInfo() *User_AuthInfo {
	if m != nil {
		return m.AuthInfo
	}
	return nil
}

func (m *User) GetSettings() map[string]string {
	if m != nil {
		return m.Settings
	}
	return nil
}

func (m *User) GetRefreshedTime() *timestamp.Timestamp {
	if m != nil {
		return m.RefreshedTime
	}
	return nil
}

func (m *User) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "User"))
	}
	m.Name = fv
}

func (m *User) SetFullName(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "FullName", "User"))
	}
	m.FullName = fv
}

func (m *User) SetMetadata(fv *ntt_meta.Meta) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Metadata", "User"))
	}
	m.Metadata = fv
}

func (m *User) SetEmail(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Email", "User"))
	}
	m.Email = fv
}

func (m *User) SetEmailVerified(fv bool) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "EmailVerified", "User"))
	}
	m.EmailVerified = fv
}

func (m *User) SetAuthInfo(fv *User_AuthInfo) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "AuthInfo", "User"))
	}
	m.AuthInfo = fv
}

func (m *User) SetSettings(fv map[string]string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Settings", "User"))
	}
	m.Settings = fv
}

func (m *User) SetRefreshedTime(fv *timestamp.Timestamp) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "RefreshedTime", "User"))
	}
	m.RefreshedTime = fv
}

type User_AuthInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// auth provider
	Provider string `protobuf:"bytes,1,opt,name=provider,proto3" json:"provider,omitempty" firestore:"provider"`
	// auth provider id
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty" firestore:"id"`
}

func (m *User_AuthInfo) Reset() {
	*m = User_AuthInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_iam_proto_v1alpha2_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *User_AuthInfo) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*User_AuthInfo) ProtoMessage() {}

func (m *User_AuthInfo) ProtoReflect() preflect.Message {
	mi := &edgelq_iam_proto_v1alpha2_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*User_AuthInfo) GotenMessage() {}

// Deprecated, Use User_AuthInfo.ProtoReflect.Descriptor instead.
func (*User_AuthInfo) Descriptor() ([]byte, []int) {
	return edgelq_iam_proto_v1alpha2_user_proto_rawDescGZIP(), []int{0, 0}
}

func (m *User_AuthInfo) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *User_AuthInfo) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *User_AuthInfo) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *User_AuthInfo) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *User_AuthInfo) GetProvider() string {
	if m != nil {
		return m.Provider
	}
	return ""
}

func (m *User_AuthInfo) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User_AuthInfo) SetProvider(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Provider", "User_AuthInfo"))
	}
	m.Provider = fv
}

func (m *User_AuthInfo) SetId(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Id", "User_AuthInfo"))
	}
	m.Id = fv
}

var edgelq_iam_proto_v1alpha2_user_proto preflect.FileDescriptor

var edgelq_iam_proto_v1alpha2_user_proto_rawDesc = []byte{
	0x0a, 0x24, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69,
	0x61, 0x6d, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x69,
	0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71,
	0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x34,
	0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x70,
	0x69, 0x6c, 0x65, 0x72, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa4, 0x07, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x20, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xb2, 0xda, 0x21,
	0x08, 0x0a, 0x06, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x23, 0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x06, 0xe2, 0xde, 0x21, 0x02, 0x08, 0x04, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2b, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x26, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x10, 0xe2, 0xde, 0x21, 0x02, 0x08, 0x04, 0xca, 0xc6, 0x27, 0x06, 0x2a, 0x04, 0x52, 0x02,
	0x58, 0x01, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x2d, 0x0a, 0x0e, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x08, 0x42, 0x06, 0xe2, 0xde, 0x21, 0x02, 0x08, 0x05, 0x52, 0x0d, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x3c, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68,
	0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6e, 0x74,
	0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x61, 0x75,
	0x74, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x40, 0x0a, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69,
	0x61, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x2e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08,
	0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x47, 0x0a, 0x0e, 0x72, 0x65, 0x66, 0x72,
	0x65, 0x73, 0x68, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x04, 0xf0, 0xd9,
	0x21, 0x01, 0x52, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x65, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x1a, 0x36, 0x0a, 0x08, 0x41, 0x75, 0x74, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x1a, 0x3b, 0x0a, 0x0d, 0x53, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x3a, 0x92, 0x03, 0xea, 0x41, 0x23, 0x0a, 0x13, 0x69, 0x61,
	0x6d, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x7d, 0x92,
	0xd9, 0x21, 0x78, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0x05, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x22, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x3a, 0x10, 0x5b, 0x5c, 0x77, 0x2e, 0x40, 0x7c, 0x5f,
	0x2d, 0x5d, 0x7b, 0x31, 0x2c, 0x31, 0x32, 0x38, 0x7d, 0x4a, 0x50, 0x08, 0x02, 0x12, 0x06, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x07, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x0b,
	0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x0e, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x1c, 0x0a,
	0x0e, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x12,
	0x0a, 0x4c, 0x61, 0x73, 0x74, 0x20, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0xaa, 0xd9, 0x21, 0x5a, 0x0a,
	0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x12, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x76,
	0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x1a, 0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x1a, 0x0e, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x65, 0x64, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x1a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x1a, 0x0e, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0xb2, 0xdf, 0x21, 0x0a, 0x0a, 0x08, 0x6c,
	0x69, 0x73, 0x74, 0x56, 0x69, 0x65, 0x77, 0xda, 0x94, 0x23, 0x09, 0x12, 0x07, 0x0a, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0xca, 0xa3, 0x22, 0x18, 0x0a, 0x0c, 0x72, 0x65, 0x61, 0x64, 0x53, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0xc2, 0x85, 0x2c, 0x57, 0x22, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x09, 0x66, 0x75, 0x6c, 0x6c,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x0e, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x22, 0x09, 0x61, 0x75,
	0x74, 0x68, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x42, 0x0e, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x42, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0xbe, 0x02, 0xe8, 0xde,
	0x21, 0x01, 0xd2, 0xff, 0xd0, 0x02, 0x40, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x12, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f,
	0x69, 0x61, 0x6d, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x32, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x74, 0x74,
	0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32,
	0x42, 0x09, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3b, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77,
	0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x3b, 0x75, 0x73, 0x65, 0x72, 0xd2, 0x84, 0xd1, 0x02, 0x46, 0x0a,
	0x0d, 0x61, 0x75, 0x64, 0x69, 0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x12, 0x35,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x61,
	0x75, 0x64, 0x69, 0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x32, 0xa2, 0x80, 0xd1, 0x02, 0x42, 0x0a, 0x0b, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67,
	0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	edgelq_iam_proto_v1alpha2_user_proto_rawDescOnce sync.Once
	edgelq_iam_proto_v1alpha2_user_proto_rawDescData = edgelq_iam_proto_v1alpha2_user_proto_rawDesc
)

func edgelq_iam_proto_v1alpha2_user_proto_rawDescGZIP() []byte {
	edgelq_iam_proto_v1alpha2_user_proto_rawDescOnce.Do(func() {
		edgelq_iam_proto_v1alpha2_user_proto_rawDescData = protoimpl.X.CompressGZIP(edgelq_iam_proto_v1alpha2_user_proto_rawDescData)
	})
	return edgelq_iam_proto_v1alpha2_user_proto_rawDescData
}

var edgelq_iam_proto_v1alpha2_user_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var edgelq_iam_proto_v1alpha2_user_proto_goTypes = []interface{}{
	(*User)(nil),                // 0: ntt.iam.v1alpha2.User
	(*User_AuthInfo)(nil),       // 1: ntt.iam.v1alpha2.User.AuthInfo
	nil,                         // 2: ntt.iam.v1alpha2.User.SettingsEntry
	(*ntt_meta.Meta)(nil),       // 3: ntt.types.Meta
	(*timestamp.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var edgelq_iam_proto_v1alpha2_user_proto_depIdxs = []int32{
	3, // 0: ntt.iam.v1alpha2.User.metadata:type_name -> ntt.types.Meta
	1, // 1: ntt.iam.v1alpha2.User.auth_info:type_name -> ntt.iam.v1alpha2.User.AuthInfo
	2, // 2: ntt.iam.v1alpha2.User.settings:type_name -> ntt.iam.v1alpha2.User.SettingsEntry
	4, // 3: ntt.iam.v1alpha2.User.refreshed_time:type_name -> google.protobuf.Timestamp
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { edgelq_iam_proto_v1alpha2_user_proto_init() }
func edgelq_iam_proto_v1alpha2_user_proto_init() {
	if edgelq_iam_proto_v1alpha2_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		edgelq_iam_proto_v1alpha2_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		edgelq_iam_proto_v1alpha2_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User_AuthInfo); i {
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
			RawDescriptor: edgelq_iam_proto_v1alpha2_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           edgelq_iam_proto_v1alpha2_user_proto_goTypes,
		DependencyIndexes: edgelq_iam_proto_v1alpha2_user_proto_depIdxs,
		MessageInfos:      edgelq_iam_proto_v1alpha2_user_proto_msgTypes,
	}.Build()
	edgelq_iam_proto_v1alpha2_user_proto = out.File
	edgelq_iam_proto_v1alpha2_user_proto_rawDesc = nil
	edgelq_iam_proto_v1alpha2_user_proto_goTypes = nil
	edgelq_iam_proto_v1alpha2_user_proto_depIdxs = nil
}
