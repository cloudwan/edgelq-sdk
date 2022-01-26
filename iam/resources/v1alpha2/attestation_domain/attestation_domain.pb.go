// Code generated by protoc-gen-goten-go
// File: edgelq/iam/proto/v1alpha2/attestation_domain.proto
// DO NOT EDIT!!!

package attestation_domain

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
	iam_common "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/common"
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
	_ = &iam_common.Actor{}
	_ = &project.Project{}
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// AttestationDomain represents attestation parameters for a heterogenous fleet
// of devices.
type AttestationDomain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Meta
	Metadata *ntt_meta.Meta `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty" firestore:"metadata"`
	// Name of AttestationDomain
	Name *Name `protobuf:"bytes,2,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// Display name
	DisplayName string `protobuf:"bytes,3,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty" firestore:"displayName"`
	// If false (default), attestation verifier SHALL check CA certificate chain
	// up to the TPM chip manufacturers CA (defined in manufacturer_ca_issuers).
	// This step is necessary to prove that the attestation request comes from a
	// genuine TPM device, and not a TPM emulator.
	//
	// If true, attestation verifier SHALL skip EKcert verification -
	// !!INSECURE!! Without this check, any keys may come from a simulated TPM
	// on the device. An attacker may therefore forge any key they want.
	// Therefore, this option should only be changed for development purposes.
	InsecureSkipManufacturerEkcertVerification bool `protobuf:"varint,4,opt,name=insecure_skip_manufacturer_ekcert_verification,json=insecureSkipManufacturerEkcertVerification,proto3" json:"insecure_skip_manufacturer_ekcert_verification,omitempty" firestore:"insecureSkipManufacturerEkcertVerification"`
	// Attestees wanting to attest under this attestation domain SHALL fulfill
	// requirements of at least one of attestation policies defined in this list.
	// Policies are checked in the order they appear on this list.
	Policies []*AttestationDomain_Policy `protobuf:"bytes,5,rep,name=policies,proto3" json:"policies,omitempty" firestore:"policies"`
	// Attestation policy may require atestees pubkey to be present on
	// enrollment_list.
	EnrollmentList []*AttestationDomain_EnrolledKey `protobuf:"bytes,6,rep,name=enrollment_list,json=enrollmentList,proto3" json:"enrollment_list,omitempty" firestore:"enrollmentList"`
}

func (m *AttestationDomain) Reset() {
	*m = AttestationDomain{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_iam_proto_v1alpha2_attestation_domain_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *AttestationDomain) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*AttestationDomain) ProtoMessage() {}

func (m *AttestationDomain) ProtoReflect() preflect.Message {
	mi := &edgelq_iam_proto_v1alpha2_attestation_domain_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*AttestationDomain) GotenMessage() {}

// Deprecated, Use AttestationDomain.ProtoReflect.Descriptor instead.
func (*AttestationDomain) Descriptor() ([]byte, []int) {
	return edgelq_iam_proto_v1alpha2_attestation_domain_proto_rawDescGZIP(), []int{0}
}

func (m *AttestationDomain) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *AttestationDomain) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *AttestationDomain) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *AttestationDomain) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *AttestationDomain) GetMetadata() *ntt_meta.Meta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *AttestationDomain) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *AttestationDomain) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *AttestationDomain) GetInsecureSkipManufacturerEkcertVerification() bool {
	if m != nil {
		return m.InsecureSkipManufacturerEkcertVerification
	}
	return false
}

func (m *AttestationDomain) GetPolicies() []*AttestationDomain_Policy {
	if m != nil {
		return m.Policies
	}
	return nil
}

func (m *AttestationDomain) GetEnrollmentList() []*AttestationDomain_EnrolledKey {
	if m != nil {
		return m.EnrollmentList
	}
	return nil
}

func (m *AttestationDomain) SetMetadata(fv *ntt_meta.Meta) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Metadata", "AttestationDomain"))
	}
	m.Metadata = fv
}

func (m *AttestationDomain) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "AttestationDomain"))
	}
	m.Name = fv
}

func (m *AttestationDomain) SetDisplayName(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "DisplayName", "AttestationDomain"))
	}
	m.DisplayName = fv
}

func (m *AttestationDomain) SetInsecureSkipManufacturerEkcertVerification(fv bool) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "InsecureSkipManufacturerEkcertVerification", "AttestationDomain"))
	}
	m.InsecureSkipManufacturerEkcertVerification = fv
}

func (m *AttestationDomain) SetPolicies(fv []*AttestationDomain_Policy) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Policies", "AttestationDomain"))
	}
	m.Policies = fv
}

func (m *AttestationDomain) SetEnrollmentList(fv []*AttestationDomain_EnrolledKey) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "EnrollmentList", "AttestationDomain"))
	}
	m.EnrollmentList = fv
}

// Policy defines a singular attestation policy, that should match a
// homogenous class of device (i.e. hardware and firmware configuration) in
// the fleet.
type AttestationDomain_Policy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// List of acceptable manufacturer's root CAs certs (in PEM format).
	//
	// To obtain such a cert, run something like:
	// `tpm2_getekcertificate -o cert.bin
	// openssl x509 -inform der -in cert.bin -noout -text`
	// Then manually download certificate specified in Authority Information
	// Access (AIA). Repeat the above `openssl` command on the newly
	// downloadedfile (change the value of `-in` argument). Continue this until
	// you reach a certificate which does not have AIA. It's the root
	// certificate. Convert it to PEM like so: `openssl x509 -inform der -in
	// root.crt -out root.pem` Still, we need to encode newlines: `awk 'NF
	// {sub(/\r/, ""); printf "%s\\n",$0;}' root.pem` Then paste the output
	// here.
	//
	// Note: multiple PEM certificates may be included here by simple means of
	// concatenation.
	ManufacturerRootCaCertsPem string `protobuf:"bytes,1,opt,name=manufacturer_root_ca_certs_pem,json=manufacturerRootCaCertsPem,proto3" json:"manufacturer_root_ca_certs_pem,omitempty" firestore:"manufacturerRootCaCertsPem"`
	// If true, the atestee's pubkey SHALL be present on this
	// AttestationDomain's enrollment list, otherwise fail the attestation.
	RequireEnrollment bool `protobuf:"varint,2,opt,name=require_enrollment,json=requireEnrollment,proto3" json:"require_enrollment,omitempty" firestore:"requireEnrollment"`
	// If true, the verifier SHALL parse, replay and verify TPM event log
	// provided by the atestee, otherwise fail the attestation.
	// Note that requiring verification of event does not provide additional
	// security. See
	// https://github.com/google/go-attestation/blob/master/docs/event-log-disclosure.md#event-type-and-verification-footguns
	VerifyEventLog bool `protobuf:"varint,3,opt,name=verify_event_log,json=verifyEventLog,proto3" json:"verify_event_log,omitempty" firestore:"verifyEventLog"`
	// List of expected PCR values.
	// All PCRs on this list SHALL match exactly the PCRs provided by the
	// atestee, otherwise the attestation SHALL be failed. The verifier SHALL
	// perform sanity checks: PCR index/digest pairs are unique and the size of
	// the list is sane (TODO).
	// They can be read from TPM by running something like `tpm2_pcrread`.
	ExpectedPcrs []*iam_common.PCR `protobuf:"bytes,4,rep,name=expected_pcrs,json=expectedPcrs,proto3" json:"expected_pcrs,omitempty" firestore:"expectedPcrs"`
}

func (m *AttestationDomain_Policy) Reset() {
	*m = AttestationDomain_Policy{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_iam_proto_v1alpha2_attestation_domain_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *AttestationDomain_Policy) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*AttestationDomain_Policy) ProtoMessage() {}

func (m *AttestationDomain_Policy) ProtoReflect() preflect.Message {
	mi := &edgelq_iam_proto_v1alpha2_attestation_domain_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*AttestationDomain_Policy) GotenMessage() {}

// Deprecated, Use AttestationDomain_Policy.ProtoReflect.Descriptor instead.
func (*AttestationDomain_Policy) Descriptor() ([]byte, []int) {
	return edgelq_iam_proto_v1alpha2_attestation_domain_proto_rawDescGZIP(), []int{0, 0}
}

func (m *AttestationDomain_Policy) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *AttestationDomain_Policy) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *AttestationDomain_Policy) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *AttestationDomain_Policy) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *AttestationDomain_Policy) GetManufacturerRootCaCertsPem() string {
	if m != nil {
		return m.ManufacturerRootCaCertsPem
	}
	return ""
}

func (m *AttestationDomain_Policy) GetRequireEnrollment() bool {
	if m != nil {
		return m.RequireEnrollment
	}
	return false
}

func (m *AttestationDomain_Policy) GetVerifyEventLog() bool {
	if m != nil {
		return m.VerifyEventLog
	}
	return false
}

func (m *AttestationDomain_Policy) GetExpectedPcrs() []*iam_common.PCR {
	if m != nil {
		return m.ExpectedPcrs
	}
	return nil
}

func (m *AttestationDomain_Policy) SetManufacturerRootCaCertsPem(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ManufacturerRootCaCertsPem", "AttestationDomain_Policy"))
	}
	m.ManufacturerRootCaCertsPem = fv
}

func (m *AttestationDomain_Policy) SetRequireEnrollment(fv bool) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "RequireEnrollment", "AttestationDomain_Policy"))
	}
	m.RequireEnrollment = fv
}

func (m *AttestationDomain_Policy) SetVerifyEventLog(fv bool) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "VerifyEventLog", "AttestationDomain_Policy"))
	}
	m.VerifyEventLog = fv
}

func (m *AttestationDomain_Policy) SetExpectedPcrs(fv []*iam_common.PCR) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ExpectedPcrs", "AttestationDomain_Policy"))
	}
	m.ExpectedPcrs = fv
}

// EnrolledKey defines an enrolled key.
type AttestationDomain_EnrolledKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// PEM encoded EK pubkey.
	// It can be read from TPM by running something like `tpm2_createek -G rsa
	// -u ek.pub -c key.ctx -f pem`.
	PubkeyPem string `protobuf:"bytes,1,opt,name=pubkey_pem,json=pubkeyPem,proto3" json:"pubkey_pem,omitempty" firestore:"pubkeyPem"`
	// User's comments for this entry
	Comment string `protobuf:"bytes,2,opt,name=comment,proto3" json:"comment,omitempty" firestore:"comment"`
}

func (m *AttestationDomain_EnrolledKey) Reset() {
	*m = AttestationDomain_EnrolledKey{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_iam_proto_v1alpha2_attestation_domain_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *AttestationDomain_EnrolledKey) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*AttestationDomain_EnrolledKey) ProtoMessage() {}

func (m *AttestationDomain_EnrolledKey) ProtoReflect() preflect.Message {
	mi := &edgelq_iam_proto_v1alpha2_attestation_domain_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*AttestationDomain_EnrolledKey) GotenMessage() {}

// Deprecated, Use AttestationDomain_EnrolledKey.ProtoReflect.Descriptor instead.
func (*AttestationDomain_EnrolledKey) Descriptor() ([]byte, []int) {
	return edgelq_iam_proto_v1alpha2_attestation_domain_proto_rawDescGZIP(), []int{0, 1}
}

func (m *AttestationDomain_EnrolledKey) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *AttestationDomain_EnrolledKey) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *AttestationDomain_EnrolledKey) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *AttestationDomain_EnrolledKey) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *AttestationDomain_EnrolledKey) GetPubkeyPem() string {
	if m != nil {
		return m.PubkeyPem
	}
	return ""
}

func (m *AttestationDomain_EnrolledKey) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *AttestationDomain_EnrolledKey) SetPubkeyPem(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "PubkeyPem", "AttestationDomain_EnrolledKey"))
	}
	m.PubkeyPem = fv
}

func (m *AttestationDomain_EnrolledKey) SetComment(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Comment", "AttestationDomain_EnrolledKey"))
	}
	m.Comment = fv
}

var edgelq_iam_proto_v1alpha2_attestation_domain_proto preflect.FileDescriptor

var edgelq_iam_proto_v1alpha2_attestation_domain_proto_rawDesc = []byte{
	0x0a, 0x32, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x61, 0x74, 0x74, 0x65,
	0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31,
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
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x63, 0x6c, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1e, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x26, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71,
	0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x69, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x65, 0x64, 0x67, 0x65,
	0x6c, 0x71, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x40, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x6d, 0x75,
	0x6c, 0x74, 0x69, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x69,
	0x6c, 0x65, 0x72, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x83, 0x07, 0x0a, 0x11, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x2b, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6e, 0x74, 0x74,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x2d, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x19, 0xb2, 0xda, 0x21, 0x15, 0x0a, 0x13, 0x0a, 0x11, 0x41, 0x74, 0x74,
	0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x29, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xde, 0x21, 0x02,
	0x08, 0x04, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x62, 0x0a, 0x2e, 0x69, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x5f, 0x73, 0x6b, 0x69, 0x70,
	0x5f, 0x6d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x5f, 0x65, 0x6b,
	0x63, 0x65, 0x72, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x2a, 0x69, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x72,
	0x65, 0x53, 0x6b, 0x69, 0x70, 0x4d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x72, 0x45, 0x6b, 0x63, 0x65, 0x72, 0x74, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x46, 0x0a, 0x08, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x52, 0x08, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x12, 0x58, 0x0a, 0x0f, 0x65,
	0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c,
	0x65, 0x64, 0x4b, 0x65, 0x79, 0x52, 0x0e, 0x65, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0xe7, 0x01, 0x0a, 0x06, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x12, 0x48, 0x0a, 0x1e, 0x6d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72,
	0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x5f, 0x63, 0x61, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x73, 0x5f, 0x70,
	0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xd0, 0xd5, 0x22, 0x01, 0x52, 0x1a,
	0x6d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x52, 0x6f, 0x6f, 0x74,
	0x43, 0x61, 0x43, 0x65, 0x72, 0x74, 0x73, 0x50, 0x65, 0x6d, 0x12, 0x2d, 0x0a, 0x12, 0x72, 0x65,
	0x71, 0x75, 0x69, 0x72, 0x65, 0x5f, 0x65, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x45,
	0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x76, 0x65, 0x72,
	0x69, 0x66, 0x79, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0e, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x4c, 0x6f, 0x67, 0x12, 0x3a, 0x0a, 0x0d, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f,
	0x70, 0x63, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6e, 0x74, 0x74,
	0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x50, 0x43,
	0x52, 0x52, 0x0c, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x50, 0x63, 0x72, 0x73, 0x1a,
	0x4c, 0x0a, 0x0b, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x64, 0x4b, 0x65, 0x79, 0x12, 0x23,
	0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x5f, 0x70, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x04, 0xd0, 0xd5, 0x22, 0x01, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6b, 0x65, 0x79,
	0x50, 0x65, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x3a, 0xa8, 0x01,
	0xea, 0x41, 0x5e, 0x0a, 0x20, 0x69, 0x61, 0x6d, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x3a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f,
	0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73, 0x2f, 0x7b, 0x61, 0x74,
	0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x7d, 0x92, 0xd9, 0x21, 0x37, 0x0a, 0x12, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73, 0x12, 0x12, 0x61, 0x74, 0x74, 0x65, 0x73,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73, 0x1a, 0x07, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0xda, 0x94, 0x23, 0x08,
	0x12, 0x06, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x88, 0x04, 0xe8, 0xde, 0x21, 0x01, 0xd2,
	0xff, 0xd0, 0x02, 0x5c, 0x0a, 0x18, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x40,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x61, 0x74,
	0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x70, 0x62,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x42, 0x16, 0x41, 0x74, 0x74, 0x65, 0x73,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x57, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f,
	0x69, 0x61, 0x6d, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x3b, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0xd2, 0x84, 0xd1, 0x02,
	0x46, 0x0a, 0x0d, 0x61, 0x75, 0x64, 0x69, 0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73,
	0x12, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d,
	0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0xf2, 0x85, 0xd1, 0x02, 0x64, 0x0a, 0x1c, 0x61, 0x74,
	0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x5f, 0x64, 0x62, 0x5f, 0x73, 0x79, 0x6e, 0x63, 0x65, 0x72, 0x12, 0x44, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f,
	0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x64, 0x62, 0x5f, 0x73, 0x79,
	0x6e, 0x63, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x61, 0x74,
	0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0xa2, 0x80, 0xd1, 0x02, 0x5e, 0x0a, 0x19, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x12, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d,
	0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32,
	0x2f, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	edgelq_iam_proto_v1alpha2_attestation_domain_proto_rawDescOnce sync.Once
	edgelq_iam_proto_v1alpha2_attestation_domain_proto_rawDescData = edgelq_iam_proto_v1alpha2_attestation_domain_proto_rawDesc
)

func edgelq_iam_proto_v1alpha2_attestation_domain_proto_rawDescGZIP() []byte {
	edgelq_iam_proto_v1alpha2_attestation_domain_proto_rawDescOnce.Do(func() {
		edgelq_iam_proto_v1alpha2_attestation_domain_proto_rawDescData = protoimpl.X.CompressGZIP(edgelq_iam_proto_v1alpha2_attestation_domain_proto_rawDescData)
	})
	return edgelq_iam_proto_v1alpha2_attestation_domain_proto_rawDescData
}

var edgelq_iam_proto_v1alpha2_attestation_domain_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var edgelq_iam_proto_v1alpha2_attestation_domain_proto_goTypes = []interface{}{
	(*AttestationDomain)(nil),             // 0: ntt.iam.v1alpha2.AttestationDomain
	(*AttestationDomain_Policy)(nil),      // 1: ntt.iam.v1alpha2.AttestationDomain.Policy
	(*AttestationDomain_EnrolledKey)(nil), // 2: ntt.iam.v1alpha2.AttestationDomain.EnrolledKey
	(*ntt_meta.Meta)(nil),                 // 3: ntt.types.Meta
	(*iam_common.PCR)(nil),                // 4: ntt.iam.v1alpha2.PCR
}
var edgelq_iam_proto_v1alpha2_attestation_domain_proto_depIdxs = []int32{
	3, // 0: ntt.iam.v1alpha2.AttestationDomain.metadata:type_name -> ntt.types.Meta
	1, // 1: ntt.iam.v1alpha2.AttestationDomain.policies:type_name -> ntt.iam.v1alpha2.AttestationDomain.Policy
	2, // 2: ntt.iam.v1alpha2.AttestationDomain.enrollment_list:type_name -> ntt.iam.v1alpha2.AttestationDomain.EnrolledKey
	4, // 3: ntt.iam.v1alpha2.AttestationDomain.Policy.expected_pcrs:type_name -> ntt.iam.v1alpha2.PCR
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { edgelq_iam_proto_v1alpha2_attestation_domain_proto_init() }
func edgelq_iam_proto_v1alpha2_attestation_domain_proto_init() {
	if edgelq_iam_proto_v1alpha2_attestation_domain_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		edgelq_iam_proto_v1alpha2_attestation_domain_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttestationDomain); i {
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
		edgelq_iam_proto_v1alpha2_attestation_domain_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttestationDomain_Policy); i {
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
		edgelq_iam_proto_v1alpha2_attestation_domain_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttestationDomain_EnrolledKey); i {
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
			RawDescriptor: edgelq_iam_proto_v1alpha2_attestation_domain_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           edgelq_iam_proto_v1alpha2_attestation_domain_proto_goTypes,
		DependencyIndexes: edgelq_iam_proto_v1alpha2_attestation_domain_proto_depIdxs,
		MessageInfos:      edgelq_iam_proto_v1alpha2_attestation_domain_proto_msgTypes,
	}.Build()
	edgelq_iam_proto_v1alpha2_attestation_domain_proto = out.File
	edgelq_iam_proto_v1alpha2_attestation_domain_proto_rawDesc = nil
	edgelq_iam_proto_v1alpha2_attestation_domain_proto_goTypes = nil
	edgelq_iam_proto_v1alpha2_attestation_domain_proto_depIdxs = nil
}
