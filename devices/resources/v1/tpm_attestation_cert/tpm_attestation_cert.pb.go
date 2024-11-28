// Code generated by protoc-gen-goten-go
// File: edgelq/devices/proto/v1/tpm_attestation_cert.proto
// DO NOT EDIT!!!

package tpm_attestation_cert

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
	_ = &project.Project{}
	_ = &meta.Meta{}
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// TpmAttestationCert Resource
type TpmAttestationCert struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Name of TpmAttestationCert
	// When creating a new instance, this field is optional and if not provided,
	// it will be generated automatically. Last ID segment must conform to the
	// following regex: [a-z][a-z0-9\\-]{0,28}[a-z0-9]
	Name *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// Metadata is an object with information like create, update and delete time
	// (for async deleted resources), has user labels/annotations, sharding
	// information, multi-region syncing information and may have non-schema
	// owners (useful for taking ownership of resources belonging to lower level
	// services by higher ones).
	Metadata     *meta.Meta `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty" firestore:"metadata"`
	DisplayName  string     `protobuf:"bytes,3,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty" firestore:"displayName"`
	Manufacturer string     `protobuf:"bytes,4,opt,name=manufacturer,proto3" json:"manufacturer,omitempty" firestore:"manufacturer"`
	// Should be in the format "productname (sku)"
	ProductName           string `protobuf:"bytes,5,opt,name=product_name,json=productName,proto3" json:"product_name,omitempty" firestore:"productName"`
	TpmManufacturerCaCert string `protobuf:"bytes,6,opt,name=tpm_manufacturer_ca_cert,json=tpmManufacturerCaCert,proto3" json:"tpm_manufacturer_ca_cert,omitempty" firestore:"tpmManufacturerCaCert"`
	IdevidIssuerCaCert    string `protobuf:"bytes,7,opt,name=idevid_issuer_ca_cert,json=idevidIssuerCaCert,proto3" json:"idevid_issuer_ca_cert,omitempty" firestore:"idevidIssuerCaCert"`
	LdevidIssuerCaCert    string `protobuf:"bytes,8,opt,name=ldevid_issuer_ca_cert,json=ldevidIssuerCaCert,proto3" json:"ldevid_issuer_ca_cert,omitempty" firestore:"ldevidIssuerCaCert"`
}

func (m *TpmAttestationCert) Reset() {
	*m = TpmAttestationCert{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_devices_proto_v1_tpm_attestation_cert_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *TpmAttestationCert) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*TpmAttestationCert) ProtoMessage() {}

func (m *TpmAttestationCert) ProtoReflect() preflect.Message {
	mi := &edgelq_devices_proto_v1_tpm_attestation_cert_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*TpmAttestationCert) GotenMessage() {}

// Deprecated, Use TpmAttestationCert.ProtoReflect.Descriptor instead.
func (*TpmAttestationCert) Descriptor() ([]byte, []int) {
	return edgelq_devices_proto_v1_tpm_attestation_cert_proto_rawDescGZIP(), []int{0}
}

func (m *TpmAttestationCert) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *TpmAttestationCert) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *TpmAttestationCert) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *TpmAttestationCert) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *TpmAttestationCert) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *TpmAttestationCert) GetMetadata() *meta.Meta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *TpmAttestationCert) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *TpmAttestationCert) GetManufacturer() string {
	if m != nil {
		return m.Manufacturer
	}
	return ""
}

func (m *TpmAttestationCert) GetProductName() string {
	if m != nil {
		return m.ProductName
	}
	return ""
}

func (m *TpmAttestationCert) GetTpmManufacturerCaCert() string {
	if m != nil {
		return m.TpmManufacturerCaCert
	}
	return ""
}

func (m *TpmAttestationCert) GetIdevidIssuerCaCert() string {
	if m != nil {
		return m.IdevidIssuerCaCert
	}
	return ""
}

func (m *TpmAttestationCert) GetLdevidIssuerCaCert() string {
	if m != nil {
		return m.LdevidIssuerCaCert
	}
	return ""
}

func (m *TpmAttestationCert) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "TpmAttestationCert"))
	}
	m.Name = fv
}

func (m *TpmAttestationCert) SetMetadata(fv *meta.Meta) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Metadata", "TpmAttestationCert"))
	}
	m.Metadata = fv
}

func (m *TpmAttestationCert) SetDisplayName(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "DisplayName", "TpmAttestationCert"))
	}
	m.DisplayName = fv
}

func (m *TpmAttestationCert) SetManufacturer(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Manufacturer", "TpmAttestationCert"))
	}
	m.Manufacturer = fv
}

func (m *TpmAttestationCert) SetProductName(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ProductName", "TpmAttestationCert"))
	}
	m.ProductName = fv
}

func (m *TpmAttestationCert) SetTpmManufacturerCaCert(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "TpmManufacturerCaCert", "TpmAttestationCert"))
	}
	m.TpmManufacturerCaCert = fv
}

func (m *TpmAttestationCert) SetIdevidIssuerCaCert(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "IdevidIssuerCaCert", "TpmAttestationCert"))
	}
	m.IdevidIssuerCaCert = fv
}

func (m *TpmAttestationCert) SetLdevidIssuerCaCert(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "LdevidIssuerCaCert", "TpmAttestationCert"))
	}
	m.LdevidIssuerCaCert = fv
}

var edgelq_devices_proto_v1_tpm_attestation_cert_proto preflect.FileDescriptor

var edgelq_devices_proto_v1_tpm_attestation_cert_proto_rawDesc = []byte{
	0x0a, 0x32, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x70, 0x6d, 0x5f, 0x61, 0x74,
	0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x6e, 0x74, 0x74, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x24, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f, 0x72, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xc2, 0x05, 0x0a, 0x12, 0x54, 0x70, 0x6d, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x43, 0x65, 0x72, 0x74, 0x12, 0x2e, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1a, 0xb2, 0xda, 0x21, 0x16, 0x0a, 0x14, 0x0a, 0x12, 0x54,
	0x70, 0x6d, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x65, 0x72,
	0x74, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67, 0x6f, 0x74, 0x65,
	0x6e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x2e, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61,
	0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0xca, 0xc6,
	0x27, 0x07, 0x2a, 0x05, 0x22, 0x03, 0x08, 0x80, 0x02, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x6d, 0x61, 0x6e, 0x75, 0x66, 0x61,
	0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6d, 0x61,
	0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x37, 0x0a,
	0x18, 0x74, 0x70, 0x6d, 0x5f, 0x6d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x72, 0x5f, 0x63, 0x61, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x15, 0x74, 0x70, 0x6d, 0x4d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72,
	0x43, 0x61, 0x43, 0x65, 0x72, 0x74, 0x12, 0x31, 0x0a, 0x15, 0x69, 0x64, 0x65, 0x76, 0x69, 0x64,
	0x5f, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x5f, 0x63, 0x61, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x69, 0x64, 0x65, 0x76, 0x69, 0x64, 0x49, 0x73, 0x73,
	0x75, 0x65, 0x72, 0x43, 0x61, 0x43, 0x65, 0x72, 0x74, 0x12, 0x31, 0x0a, 0x15, 0x6c, 0x64, 0x65,
	0x76, 0x69, 0x64, 0x5f, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x5f, 0x63, 0x61, 0x5f, 0x63, 0x65,
	0x72, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x6c, 0x64, 0x65, 0x76, 0x69, 0x64,
	0x49, 0x73, 0x73, 0x75, 0x65, 0x72, 0x43, 0x61, 0x43, 0x65, 0x72, 0x74, 0x3a, 0xb6, 0x02, 0xea,
	0x41, 0x92, 0x01, 0x0a, 0x25, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x65, 0x64, 0x67,
	0x65, 0x6c, 0x71, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x54, 0x70, 0x6d, 0x41, 0x74, 0x74, 0x65, 0x73,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x65, 0x72, 0x74, 0x12, 0x2a, 0x74, 0x70, 0x6d, 0x41,
	0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x65, 0x72, 0x74, 0x73, 0x2f,
	0x7b, 0x74, 0x70, 0x6d, 0x5f, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x63, 0x65, 0x72, 0x74, 0x7d, 0x12, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x74, 0x70, 0x6d, 0x41, 0x74,
	0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x7b,
	0x74, 0x70, 0x6d, 0x5f, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x63, 0x65, 0x72, 0x74, 0x7d, 0x92, 0xd9, 0x21, 0x75, 0x0a, 0x13, 0x74, 0x70, 0x6d, 0x41, 0x74,
	0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x65, 0x72, 0x74, 0x73, 0x12, 0x13,
	0x74, 0x70, 0x6d, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x65,
	0x72, 0x74, 0x73, 0x1a, 0x12, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x4e, 0x6f, 0x6e, 0x65, 0x1a, 0x07, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x38, 0x05, 0x42, 0x2a, 0x08, 0x02, 0x12, 0x06, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e,
	0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e,
	0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0xb2, 0xdf,
	0x21, 0x17, 0x12, 0x15, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x05, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0xda, 0x94, 0x23, 0x08, 0x12, 0x06, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0xd9, 0x02, 0xe8, 0xde, 0x21, 0x01, 0xd2, 0xff, 0xd0, 0x02,
	0x5e, 0x0a, 0x1a, 0x74, 0x70, 0x6d, 0x5f, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x40, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77,
	0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x70, 0x6d, 0x5f, 0x61,
	0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x65, 0x72, 0x74, 0xa2,
	0x80, 0xd1, 0x02, 0x60, 0x0a, 0x1b, 0x74, 0x70, 0x6d, 0x5f, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x12, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x74, 0x70, 0x6d, 0x5f, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x63, 0x65, 0x72, 0x74, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x42, 0x17, 0x54, 0x70, 0x6d,
	0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x65, 0x72, 0x74, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x59, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65,
	0x6c, 0x71, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x70, 0x6d, 0x5f, 0x61, 0x74, 0x74, 0x65,
	0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x3b, 0x74, 0x70, 0x6d,
	0x5f, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x65, 0x72,
	0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	edgelq_devices_proto_v1_tpm_attestation_cert_proto_rawDescOnce sync.Once
	edgelq_devices_proto_v1_tpm_attestation_cert_proto_rawDescData = edgelq_devices_proto_v1_tpm_attestation_cert_proto_rawDesc
)

func edgelq_devices_proto_v1_tpm_attestation_cert_proto_rawDescGZIP() []byte {
	edgelq_devices_proto_v1_tpm_attestation_cert_proto_rawDescOnce.Do(func() {
		edgelq_devices_proto_v1_tpm_attestation_cert_proto_rawDescData = protoimpl.X.CompressGZIP(edgelq_devices_proto_v1_tpm_attestation_cert_proto_rawDescData)
	})
	return edgelq_devices_proto_v1_tpm_attestation_cert_proto_rawDescData
}

var edgelq_devices_proto_v1_tpm_attestation_cert_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var edgelq_devices_proto_v1_tpm_attestation_cert_proto_goTypes = []interface{}{
	(*TpmAttestationCert)(nil), // 0: ntt.devices.v1.TpmAttestationCert
	(*meta.Meta)(nil),          // 1: goten.types.Meta
}
var edgelq_devices_proto_v1_tpm_attestation_cert_proto_depIdxs = []int32{
	1, // 0: ntt.devices.v1.TpmAttestationCert.metadata:type_name -> goten.types.Meta
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { edgelq_devices_proto_v1_tpm_attestation_cert_proto_init() }
func edgelq_devices_proto_v1_tpm_attestation_cert_proto_init() {
	if edgelq_devices_proto_v1_tpm_attestation_cert_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		edgelq_devices_proto_v1_tpm_attestation_cert_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TpmAttestationCert); i {
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
			RawDescriptor: edgelq_devices_proto_v1_tpm_attestation_cert_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           edgelq_devices_proto_v1_tpm_attestation_cert_proto_goTypes,
		DependencyIndexes: edgelq_devices_proto_v1_tpm_attestation_cert_proto_depIdxs,
		MessageInfos:      edgelq_devices_proto_v1_tpm_attestation_cert_proto_msgTypes,
	}.Build()
	edgelq_devices_proto_v1_tpm_attestation_cert_proto = out.File
	edgelq_devices_proto_v1_tpm_attestation_cert_proto_rawDesc = nil
	edgelq_devices_proto_v1_tpm_attestation_cert_proto_goTypes = nil
	edgelq_devices_proto_v1_tpm_attestation_cert_proto_depIdxs = nil
}
