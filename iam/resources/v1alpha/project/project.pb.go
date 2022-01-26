// Code generated by protoc-gen-goten-go
// File: edgelq/iam/proto/v1alpha/project.proto
// DO NOT EDIT!!!

package project

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
	organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/organization"
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
	_ = &organization.Organization{}
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Project Resource
type Project struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Name of Project
	Name *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// Title
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty" firestore:"title"`
	// Parent organization
	ParentOrganization *organization.Reference `protobuf:"bytes,3,opt,customtype=Reference,name=parent_organization,json=parentOrganization,proto3" json:"parent_organization,omitempty" firestore:"parentOrganization"`
	// Top parent
	RootOrganization *organization.Reference `protobuf:"bytes,4,opt,customtype=Reference,name=root_organization,json=rootOrganization,proto3" json:"root_organization,omitempty" firestore:"rootOrganization"`
	// Full ancestry path
	AncestryPath []*organization.Reference `protobuf:"bytes,5,rep,customtype=Reference,name=ancestry_path,json=ancestryPath,proto3" json:"ancestry_path,omitempty" firestore:"ancestryPath"`
	// Metadata
	Metadata *ntt_meta.Meta `protobuf:"bytes,6,opt,name=metadata,proto3" json:"metadata,omitempty" firestore:"metadata"`
}

func (m *Project) Reset() {
	*m = Project{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_iam_proto_v1alpha_project_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *Project) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*Project) ProtoMessage() {}

func (m *Project) ProtoReflect() preflect.Message {
	mi := &edgelq_iam_proto_v1alpha_project_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*Project) GotenMessage() {}

// Deprecated, Use Project.ProtoReflect.Descriptor instead.
func (*Project) Descriptor() ([]byte, []int) {
	return edgelq_iam_proto_v1alpha_project_proto_rawDescGZIP(), []int{0}
}

func (m *Project) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *Project) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *Project) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *Project) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *Project) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *Project) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Project) GetParentOrganization() *organization.Reference {
	if m != nil {
		return m.ParentOrganization
	}
	return nil
}

func (m *Project) GetRootOrganization() *organization.Reference {
	if m != nil {
		return m.RootOrganization
	}
	return nil
}

func (m *Project) GetAncestryPath() []*organization.Reference {
	if m != nil {
		return m.AncestryPath
	}
	return nil
}

func (m *Project) GetMetadata() *ntt_meta.Meta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Project) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "Project"))
	}
	m.Name = fv
}

func (m *Project) SetTitle(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Title", "Project"))
	}
	m.Title = fv
}

func (m *Project) SetParentOrganization(fv *organization.Reference) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ParentOrganization", "Project"))
	}
	m.ParentOrganization = fv
}

func (m *Project) SetRootOrganization(fv *organization.Reference) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "RootOrganization", "Project"))
	}
	m.RootOrganization = fv
}

func (m *Project) SetAncestryPath(fv []*organization.Reference) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "AncestryPath", "Project"))
	}
	m.AncestryPath = fv
}

func (m *Project) SetMetadata(fv *ntt_meta.Meta) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Metadata", "Project"))
	}
	m.Metadata = fv
}

var edgelq_iam_proto_v1alpha_project_proto preflect.FileDescriptor

var edgelq_iam_proto_v1alpha_project_proto_rawDesc = []byte{
	0x0a, 0x26, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61,
	0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70,
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
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x65, 0x64, 0x67, 0x65,
	0x6c, 0x71, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x20, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1e, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xd1, 0x05, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x23,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0f, 0xb2, 0xda,
	0x21, 0x0b, 0x0a, 0x09, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x06, 0xe2, 0xde, 0x21, 0x02, 0x08, 0x04, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x5a, 0x0a, 0x13, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x6f, 0x72, 0x67, 0x61,
	0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x29,
	0xfa, 0x41, 0x0e, 0x0a, 0x0c, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0xb2, 0xda, 0x21, 0x10, 0x12, 0x0e, 0x0a, 0x0c, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0xba, 0x9d, 0x22, 0x00, 0x52, 0x12, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x58, 0x0a,
	0x11, 0x72, 0x6f, 0x6f, 0x74, 0x5f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2b, 0xfa, 0x41, 0x0e, 0x0a, 0x0c, 0x4f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0xb2, 0xda, 0x21, 0x12, 0x12,
	0x10, 0x0a, 0x0c, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x10,
	0x03, 0xba, 0x9d, 0x22, 0x00, 0x52, 0x10, 0x72, 0x6f, 0x6f, 0x74, 0x4f, 0x72, 0x67, 0x61, 0x6e,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x50, 0x0a, 0x0d, 0x61, 0x6e, 0x63, 0x65, 0x73,
	0x74, 0x72, 0x79, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x42, 0x2b,
	0xfa, 0x41, 0x0e, 0x0a, 0x0c, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0xb2, 0xda, 0x21, 0x12, 0x12, 0x10, 0x0a, 0x0c, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x03, 0xba, 0x9d, 0x22, 0x00, 0x52, 0x0c, 0x61, 0x6e, 0x63,
	0x65, 0x73, 0x74, 0x72, 0x79, 0x50, 0x61, 0x74, 0x68, 0x12, 0x2b, 0x0a, 0x08, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6e, 0x74,
	0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x3a, 0xcd, 0x02, 0xea, 0x41, 0x2c, 0x0a, 0x16, 0x69, 0x61,
	0x6d, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x12, 0x12, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x92, 0xd9, 0x21, 0x5b, 0x0a, 0x08, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x22, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x4a, 0x13, 0x08, 0x01, 0x12, 0x06, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x07, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x4a, 0x2a, 0x08, 0x02, 0x12,
	0x06, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x07, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x15, 0x0a, 0x13, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x6f, 0x72, 0x67, 0x61, 0x6e,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0xfa, 0xde, 0x21, 0x09, 0x0a, 0x07, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0xb2, 0xdf, 0x21, 0x0a, 0x0a, 0x08, 0x6c, 0x69, 0x73, 0x74, 0x56, 0x69,
	0x65, 0x77, 0xda, 0x94, 0x23, 0x08, 0x12, 0x06, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0xca, 0xa3,
	0x22, 0x42, 0x0a, 0x1e, 0x69, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e,
	0x72, 0x65, 0x61, 0x64, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x73, 0x12, 0x11, 0x72, 0x6f, 0x6f, 0x74, 0x5f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0d, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x74, 0x72, 0x79, 0x5f,
	0x70, 0x61, 0x74, 0x68, 0xc2, 0x85, 0x2c, 0x4e, 0x22, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x13, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x6f, 0x72,
	0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2a, 0x11, 0x72, 0x6f, 0x6f, 0x74,
	0x5f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2a, 0x0d, 0x61,
	0x6e, 0x63, 0x65, 0x73, 0x74, 0x72, 0x79, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x42, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0xa2, 0x03, 0xe8, 0xde, 0x21, 0x01, 0xd2, 0xff, 0xd0,
	0x02, 0x45, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x12, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61,
	0x6d, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x92, 0x8c, 0xd1, 0x02, 0x4f, 0x0a, 0x12, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67,
	0x12, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d,
	0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x0a, 0x16, 0x63, 0x6f, 0x6d,
	0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x42, 0x0c, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f,
	0x69, 0x61, 0x6d, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x3b, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0xd2, 0x84, 0xd1, 0x02, 0x45, 0x0a, 0x0d, 0x61, 0x75, 0x64, 0x69,
	0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x12, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65,
	0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x68,
	0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0xa2,
	0x80, 0xd1, 0x02, 0x47, 0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x12, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71,
	0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	edgelq_iam_proto_v1alpha_project_proto_rawDescOnce sync.Once
	edgelq_iam_proto_v1alpha_project_proto_rawDescData = edgelq_iam_proto_v1alpha_project_proto_rawDesc
)

func edgelq_iam_proto_v1alpha_project_proto_rawDescGZIP() []byte {
	edgelq_iam_proto_v1alpha_project_proto_rawDescOnce.Do(func() {
		edgelq_iam_proto_v1alpha_project_proto_rawDescData = protoimpl.X.CompressGZIP(edgelq_iam_proto_v1alpha_project_proto_rawDescData)
	})
	return edgelq_iam_proto_v1alpha_project_proto_rawDescData
}

var edgelq_iam_proto_v1alpha_project_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var edgelq_iam_proto_v1alpha_project_proto_goTypes = []interface{}{
	(*Project)(nil),       // 0: ntt.iam.v1alpha.Project
	(*ntt_meta.Meta)(nil), // 1: ntt.types.Meta
}
var edgelq_iam_proto_v1alpha_project_proto_depIdxs = []int32{
	1, // 0: ntt.iam.v1alpha.Project.metadata:type_name -> ntt.types.Meta
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { edgelq_iam_proto_v1alpha_project_proto_init() }
func edgelq_iam_proto_v1alpha_project_proto_init() {
	if edgelq_iam_proto_v1alpha_project_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		edgelq_iam_proto_v1alpha_project_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Project); i {
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
			RawDescriptor: edgelq_iam_proto_v1alpha_project_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           edgelq_iam_proto_v1alpha_project_proto_goTypes,
		DependencyIndexes: edgelq_iam_proto_v1alpha_project_proto_depIdxs,
		MessageInfos:      edgelq_iam_proto_v1alpha_project_proto_msgTypes,
	}.Build()
	edgelq_iam_proto_v1alpha_project_proto = out.File
	edgelq_iam_proto_v1alpha_project_proto_rawDesc = nil
	edgelq_iam_proto_v1alpha_project_proto_goTypes = nil
	edgelq_iam_proto_v1alpha_project_proto_depIdxs = nil
}
