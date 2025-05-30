// Code generated by protoc-gen-goten-go
// File: edgelq/alerting/proto/v1/document.proto
// DO NOT EDIT!!!

package document

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
	iam_project "github.com/cloudwan/edgelq-sdk/iam/resources/v1/project"
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
	_ = &iam_project.Project{}
	_ = &meta.Meta{}
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Document can be attached to Policy, TsCondition or LogCondition
// resources. They should contain documentation that can be relevant when
// diagnosing & resolving alerts.
// They are especially important for AI Agent handling.
type Document struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Name of Document
	// When creating a new instance, this field is optional and if not provided,
	// it will be generated automatically. Last ID segment must conform to the
	// following regex: [a-z][a-z0-9\\-]{0,28}[a-z0-9]
	Name *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty"`
	// Metadata is an object with information like create, update and delete time
	// (for async deleted resources), has user labels/annotations, sharding
	// information, multi-region syncing information and may have non-schema
	// owners (useful for taking ownership of resources belonging to lower level
	// services by higher ones).
	Metadata *meta.Meta `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Document title
	Title string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	// Documentation content for operators
	Content string `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	// documentation mime type. Only `"text/markdown"` is supported.
	MimeType string `protobuf:"bytes,5,opt,name=mime_type,json=mimeType,proto3" json:"mime_type,omitempty"`
}

func (m *Document) Reset() {
	*m = Document{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_alerting_proto_v1_document_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *Document) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*Document) ProtoMessage() {}

func (m *Document) ProtoReflect() preflect.Message {
	mi := &edgelq_alerting_proto_v1_document_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*Document) GotenMessage() {}

// Deprecated, Use Document.ProtoReflect.Descriptor instead.
func (*Document) Descriptor() ([]byte, []int) {
	return edgelq_alerting_proto_v1_document_proto_rawDescGZIP(), []int{0}
}

func (m *Document) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *Document) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *Document) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *Document) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *Document) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *Document) GetMetadata() *meta.Meta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Document) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Document) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Document) GetMimeType() string {
	if m != nil {
		return m.MimeType
	}
	return ""
}

func (m *Document) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "Document"))
	}
	m.Name = fv
}

func (m *Document) SetMetadata(fv *meta.Meta) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Metadata", "Document"))
	}
	m.Metadata = fv
}

func (m *Document) SetTitle(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Title", "Document"))
	}
	m.Title = fv
}

func (m *Document) SetContent(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Content", "Document"))
	}
	m.Content = fv
}

func (m *Document) SetMimeType(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "MimeType", "Document"))
	}
	m.MimeType = fv
}

var edgelq_alerting_proto_v1_document_proto preflect.FileDescriptor

var edgelq_alerting_proto_v1_document_proto_rawDesc = []byte{
	0x0a, 0x27, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e,
	0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x6f, 0x63, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x6e, 0x74, 0x74, 0x2e, 0x61,
	0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x74, 0x65, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6d, 0x75, 0x6c, 0x74,
	0x69, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24,
	0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe1, 0x04, 0x0a,
	0x08, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x10, 0xb2, 0xda, 0x21, 0x0c, 0x0a, 0x0a, 0x0a,
	0x08, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x2d, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x4d, 0x65, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x32,
	0x0a, 0x09, 0x6d, 0x69, 0x6d, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x15, 0xca, 0xc6, 0x27, 0x11, 0x2a, 0x0f, 0x5a, 0x0d, 0x74, 0x65, 0x78, 0x74, 0x2f,
	0x6d, 0x61, 0x72, 0x6b, 0x64, 0x6f, 0x77, 0x6e, 0x52, 0x08, 0x6d, 0x69, 0x6d, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x3a, 0x9b, 0x03, 0xea, 0x41, 0x47, 0x0a, 0x1c, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69,
	0x6e, 0x67, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x6f,
	0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x27, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x64, 0x6f, 0x63, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x7d, 0x92,
	0xd9, 0x21, 0xb8, 0x01, 0x0a, 0x09, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12,
	0x09, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0x16, 0x69, 0x61, 0x6d, 0x2e,
	0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x38, 0x05, 0x42, 0x38, 0x08, 0x02, 0x12, 0x0c, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x09, 0x6d, 0x69, 0x6d, 0x65, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x09, 0x4d, 0x69, 0x6d, 0x65, 0x20, 0x74, 0x79, 0x70, 0x65, 0x42, 0x4c,
	0x08, 0x03, 0x12, 0x0c, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x04, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x0e, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x12, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x07, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x09, 0x6d, 0x69, 0x6d, 0x65, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x09, 0x4d, 0x69, 0x6d, 0x65, 0x20, 0x74, 0x79, 0x70, 0x65, 0xaa, 0xd9, 0x21, 0x1d,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x1a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0xb2, 0xdf, 0x21,
	0x34, 0x0a, 0x32, 0x0a, 0x30, 0x0a, 0x07, 0x62, 0x79, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x1a, 0x1e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f,
	0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x2f, 0x2d, 0xda, 0x94, 0x23, 0x08, 0x12, 0x06, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0xc2, 0x85, 0x2c, 0x2b, 0x22, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x22, 0x09, 0x6d, 0x69, 0x6d, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x42, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x42, 0x8b, 0x02, 0xe8, 0xde, 0x21, 0x01, 0xd2, 0xff, 0xd0, 0x02, 0x47, 0x0a, 0x0e, 0x64, 0x6f,
	0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x35, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61,
	0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e,
	0x67, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x6f, 0x63, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0xa2, 0x80, 0xd1, 0x02, 0x49, 0x0a, 0x0f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65,
	0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74,
	0x69, 0x6e, 0x67, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x44, 0x6f, 0x63, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f,
	0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x6f, 0x63,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x3b, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	edgelq_alerting_proto_v1_document_proto_rawDescOnce sync.Once
	edgelq_alerting_proto_v1_document_proto_rawDescData = edgelq_alerting_proto_v1_document_proto_rawDesc
)

func edgelq_alerting_proto_v1_document_proto_rawDescGZIP() []byte {
	edgelq_alerting_proto_v1_document_proto_rawDescOnce.Do(func() {
		edgelq_alerting_proto_v1_document_proto_rawDescData = protoimpl.X.CompressGZIP(edgelq_alerting_proto_v1_document_proto_rawDescData)
	})
	return edgelq_alerting_proto_v1_document_proto_rawDescData
}

var edgelq_alerting_proto_v1_document_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var edgelq_alerting_proto_v1_document_proto_goTypes = []interface{}{
	(*Document)(nil),  // 0: ntt.alerting.v1.Document
	(*meta.Meta)(nil), // 1: goten.types.Meta
}
var edgelq_alerting_proto_v1_document_proto_depIdxs = []int32{
	1, // 0: ntt.alerting.v1.Document.metadata:type_name -> goten.types.Meta
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { edgelq_alerting_proto_v1_document_proto_init() }
func edgelq_alerting_proto_v1_document_proto_init() {
	if edgelq_alerting_proto_v1_document_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		edgelq_alerting_proto_v1_document_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Document); i {
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
			RawDescriptor: edgelq_alerting_proto_v1_document_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           edgelq_alerting_proto_v1_document_proto_goTypes,
		DependencyIndexes: edgelq_alerting_proto_v1_document_proto_depIdxs,
		MessageInfos:      edgelq_alerting_proto_v1_document_proto_msgTypes,
	}.Build()
	edgelq_alerting_proto_v1_document_proto = out.File
	edgelq_alerting_proto_v1_document_proto_rawDesc = nil
	edgelq_alerting_proto_v1_document_proto_goTypes = nil
	edgelq_alerting_proto_v1_document_proto_depIdxs = nil
}
