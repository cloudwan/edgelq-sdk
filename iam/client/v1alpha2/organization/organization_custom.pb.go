// Code generated by protoc-gen-goten-go
// File: edgelq/iam/proto/v1alpha2/organization_custom.proto
// DO NOT EDIT!!!

package organization_client

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
	organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/organization"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
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
	_ = &organization.Organization{}
	_ = &fieldmaskpb.FieldMask{}
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Request message for method
// [ListMyOrganizations][ntt.iam.v1alpha2.ListMyOrganizations]
type ListMyOrganizationsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Requested page size. Server may return fewer Organizations than requested.
	// If unspecified, server will pick an appropriate default.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// A token identifying a page of results the server should return.
	// Typically, this is the value of
	// [ListOrganizationsResponse.next_page_token][ntt.iam.v1alpha2.ListOrganizationsResponse.next_page_token]
	PageToken *organization.PagerCursor `protobuf:"bytes,3,opt,customtype=PagerCursor,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// Order By -
	// https://cloud.google.com/apis/design/design_patterns#list_pagination list
	// of field path with order directive, either 'asc' or 'desc'. If direction is
	// not provided, 'asc' is assumed. e.g. "state.nested_field asc,
	// state.something.else desc, theme"
	OrderBy *organization.OrderBy `protobuf:"bytes,4,opt,customtype=OrderBy,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
	// Filter - filter results by field criteria. Simplified SQL-like syntax with
	// following operators:
	// <=, >=, =, !=, <, >, LIKE, CONTAINS (aliases CONTAIN, HAS, HAVE), IN, IS
	// [NOT] NULL | NaN . Combine conditions with OR | AND example: 'meta.labels
	// CONTAINS "severity:important" OR (state.last_error_time >
	// "2018-11-15T10:00:00Z" AND state.status = "ERROR")'
	Filter *organization.Filter `protobuf:"bytes,5,opt,customtype=Filter,name=filter,proto3" json:"filter,omitempty"`
	// A FieldMask used to filter response fields. When present, only requested
	// fields will be present in each response item. Omitting field_mask results
	// will cause response items to contain all present fields.
	FieldMask *organization.Organization_FieldMask `protobuf:"bytes,6,opt,customtype=Organization_FieldMask,name=field_mask,json=fieldMask,proto3" json:"field_mask,omitempty"`
}

func (m *ListMyOrganizationsRequest) Reset() {
	*m = ListMyOrganizationsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_iam_proto_v1alpha2_organization_custom_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *ListMyOrganizationsRequest) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*ListMyOrganizationsRequest) ProtoMessage() {}

func (m *ListMyOrganizationsRequest) ProtoReflect() preflect.Message {
	mi := &edgelq_iam_proto_v1alpha2_organization_custom_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*ListMyOrganizationsRequest) GotenMessage() {}

// Deprecated, Use ListMyOrganizationsRequest.ProtoReflect.Descriptor instead.
func (*ListMyOrganizationsRequest) Descriptor() ([]byte, []int) {
	return edgelq_iam_proto_v1alpha2_organization_custom_proto_rawDescGZIP(), []int{0}
}

func (m *ListMyOrganizationsRequest) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *ListMyOrganizationsRequest) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *ListMyOrganizationsRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *ListMyOrganizationsRequest) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *ListMyOrganizationsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return int32(0)
}

func (m *ListMyOrganizationsRequest) GetPageToken() *organization.PagerCursor {
	if m != nil {
		return m.PageToken
	}
	return nil
}

func (m *ListMyOrganizationsRequest) GetOrderBy() *organization.OrderBy {
	if m != nil {
		return m.OrderBy
	}
	return nil
}

func (m *ListMyOrganizationsRequest) GetFilter() *organization.Filter {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *ListMyOrganizationsRequest) GetFieldMask() *organization.Organization_FieldMask {
	if m != nil {
		return m.FieldMask
	}
	return nil
}

func (m *ListMyOrganizationsRequest) SetPageSize(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "PageSize", "ListMyOrganizationsRequest"))
	}
	m.PageSize = fv
}

func (m *ListMyOrganizationsRequest) SetPageToken(fv *organization.PagerCursor) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "PageToken", "ListMyOrganizationsRequest"))
	}
	m.PageToken = fv
}

func (m *ListMyOrganizationsRequest) SetOrderBy(fv *organization.OrderBy) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "OrderBy", "ListMyOrganizationsRequest"))
	}
	m.OrderBy = fv
}

func (m *ListMyOrganizationsRequest) SetFilter(fv *organization.Filter) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Filter", "ListMyOrganizationsRequest"))
	}
	m.Filter = fv
}

func (m *ListMyOrganizationsRequest) SetFieldMask(fv *organization.Organization_FieldMask) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "FieldMask", "ListMyOrganizationsRequest"))
	}
	m.FieldMask = fv
}

// Response message for method
// [ListMyOrganizations][ntt.iam.v1alpha2.ListMyOrganizations]
type ListMyOrganizationsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// The list of Organizations
	Organizations []*organization.Organization `protobuf:"bytes,1,rep,name=organizations,proto3" json:"organizations,omitempty"`
	// A token to retrieve previous page of results. Pass this value in the
	// [ListOrganizationsRequest.page_token][ntt.iam.v1alpha2.ListOrganizationsRequest.page_token]
	PrevPageToken *organization.PagerCursor `protobuf:"bytes,3,opt,customtype=PagerCursor,name=prev_page_token,json=prevPageToken,proto3" json:"prev_page_token,omitempty"`
	// A token to retrieve next page of results. Pass this value in the
	// [ListOrganizationsRequest.page_token][ntt.iam.v1alpha2.ListOrganizationsRequest.page_token]
	NextPageToken *organization.PagerCursor `protobuf:"bytes,4,opt,customtype=PagerCursor,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (m *ListMyOrganizationsResponse) Reset() {
	*m = ListMyOrganizationsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_iam_proto_v1alpha2_organization_custom_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *ListMyOrganizationsResponse) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*ListMyOrganizationsResponse) ProtoMessage() {}

func (m *ListMyOrganizationsResponse) ProtoReflect() preflect.Message {
	mi := &edgelq_iam_proto_v1alpha2_organization_custom_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*ListMyOrganizationsResponse) GotenMessage() {}

// Deprecated, Use ListMyOrganizationsResponse.ProtoReflect.Descriptor instead.
func (*ListMyOrganizationsResponse) Descriptor() ([]byte, []int) {
	return edgelq_iam_proto_v1alpha2_organization_custom_proto_rawDescGZIP(), []int{1}
}

func (m *ListMyOrganizationsResponse) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *ListMyOrganizationsResponse) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *ListMyOrganizationsResponse) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *ListMyOrganizationsResponse) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *ListMyOrganizationsResponse) GetOrganizations() []*organization.Organization {
	if m != nil {
		return m.Organizations
	}
	return nil
}

func (m *ListMyOrganizationsResponse) GetPrevPageToken() *organization.PagerCursor {
	if m != nil {
		return m.PrevPageToken
	}
	return nil
}

func (m *ListMyOrganizationsResponse) GetNextPageToken() *organization.PagerCursor {
	if m != nil {
		return m.NextPageToken
	}
	return nil
}

func (m *ListMyOrganizationsResponse) SetOrganizations(fv []*organization.Organization) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Organizations", "ListMyOrganizationsResponse"))
	}
	m.Organizations = fv
}

func (m *ListMyOrganizationsResponse) SetPrevPageToken(fv *organization.PagerCursor) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "PrevPageToken", "ListMyOrganizationsResponse"))
	}
	m.PrevPageToken = fv
}

func (m *ListMyOrganizationsResponse) SetNextPageToken(fv *organization.PagerCursor) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "NextPageToken", "ListMyOrganizationsResponse"))
	}
	m.NextPageToken = fv
}

var edgelq_iam_proto_v1alpha2_organization_custom_proto preflect.FileDescriptor

var edgelq_iam_proto_v1alpha2_organization_custom_proto_rawDesc = []byte{
	0x0a, 0x33, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x6f, 0x72, 0x67, 0x61,
	0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x74,
	0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71,
	0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x69, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x65, 0x64, 0x67, 0x65,
	0x6c, 0x71, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x2c, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x6f, 0x72,
	0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x33, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x6f, 0x72, 0x67, 0x61,
	0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa8, 0x02, 0x0a, 0x1a, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x79,
	0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x08, 0xca, 0xc6, 0x27, 0x04, 0x12, 0x02, 0x2a,
	0x00, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x33, 0x0a, 0x0a, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x14, 0xb2, 0xda, 0x21, 0x10, 0x22, 0x0e, 0x0a, 0x0c, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x2f, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x14, 0xb2, 0xda, 0x21, 0x10, 0x2a, 0x0e, 0x0a, 0x0c, 0x4f, 0x72, 0x67, 0x61,
	0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42,
	0x79, 0x12, 0x2c, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x14, 0xb2, 0xda, 0x21, 0x10, 0x1a, 0x0e, 0x0a, 0x0c, 0x4f, 0x72, 0x67, 0x61, 0x6e,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12,
	0x4f, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x42,
	0x14, 0xb2, 0xda, 0x21, 0x10, 0x32, 0x0e, 0x0a, 0x0c, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b,
	0x22, 0xdf, 0x01, 0x0a, 0x1b, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x79, 0x4f, 0x72, 0x67, 0x61, 0x6e,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x44, 0x0a, 0x0d, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61,
	0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x4f, 0x72, 0x67, 0x61, 0x6e,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x3c, 0x0a, 0x0f, 0x70, 0x72, 0x65, 0x76, 0x5f, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x14, 0xb2, 0xda, 0x21, 0x10, 0x22, 0x0e, 0x0a, 0x0c, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x70, 0x72, 0x65, 0x76, 0x50, 0x61, 0x67, 0x65, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x3c, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x14, 0xb2,
	0xda, 0x21, 0x10, 0x22, 0x0e, 0x0a, 0x0c, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x42, 0x85, 0x01, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x69,
	0x61, 0x6d, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x42, 0x17,
	0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x00, 0x5a, 0x4f, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65,
	0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3b, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	edgelq_iam_proto_v1alpha2_organization_custom_proto_rawDescOnce sync.Once
	edgelq_iam_proto_v1alpha2_organization_custom_proto_rawDescData = edgelq_iam_proto_v1alpha2_organization_custom_proto_rawDesc
)

func edgelq_iam_proto_v1alpha2_organization_custom_proto_rawDescGZIP() []byte {
	edgelq_iam_proto_v1alpha2_organization_custom_proto_rawDescOnce.Do(func() {
		edgelq_iam_proto_v1alpha2_organization_custom_proto_rawDescData = protoimpl.X.CompressGZIP(edgelq_iam_proto_v1alpha2_organization_custom_proto_rawDescData)
	})
	return edgelq_iam_proto_v1alpha2_organization_custom_proto_rawDescData
}

var edgelq_iam_proto_v1alpha2_organization_custom_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var edgelq_iam_proto_v1alpha2_organization_custom_proto_goTypes = []interface{}{
	(*ListMyOrganizationsRequest)(nil),          // 0: ntt.iam.v1alpha2.ListMyOrganizationsRequest
	(*ListMyOrganizationsResponse)(nil),         // 1: ntt.iam.v1alpha2.ListMyOrganizationsResponse
	(*organization.Organization_FieldMask)(nil), // 2: ntt.iam.v1alpha2.Organization_FieldMask
	(*organization.Organization)(nil),           // 3: ntt.iam.v1alpha2.Organization
}
var edgelq_iam_proto_v1alpha2_organization_custom_proto_depIdxs = []int32{
	2, // 0: ntt.iam.v1alpha2.ListMyOrganizationsRequest.field_mask:type_name -> ntt.iam.v1alpha2.Organization_FieldMask
	3, // 1: ntt.iam.v1alpha2.ListMyOrganizationsResponse.organizations:type_name -> ntt.iam.v1alpha2.Organization
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { edgelq_iam_proto_v1alpha2_organization_custom_proto_init() }
func edgelq_iam_proto_v1alpha2_organization_custom_proto_init() {
	if edgelq_iam_proto_v1alpha2_organization_custom_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		edgelq_iam_proto_v1alpha2_organization_custom_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMyOrganizationsRequest); i {
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
		edgelq_iam_proto_v1alpha2_organization_custom_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMyOrganizationsResponse); i {
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
			RawDescriptor: edgelq_iam_proto_v1alpha2_organization_custom_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           edgelq_iam_proto_v1alpha2_organization_custom_proto_goTypes,
		DependencyIndexes: edgelq_iam_proto_v1alpha2_organization_custom_proto_depIdxs,
		MessageInfos:      edgelq_iam_proto_v1alpha2_organization_custom_proto_msgTypes,
	}.Build()
	edgelq_iam_proto_v1alpha2_organization_custom_proto = out.File
	edgelq_iam_proto_v1alpha2_organization_custom_proto_rawDesc = nil
	edgelq_iam_proto_v1alpha2_organization_custom_proto_goTypes = nil
	edgelq_iam_proto_v1alpha2_organization_custom_proto_depIdxs = nil
}
