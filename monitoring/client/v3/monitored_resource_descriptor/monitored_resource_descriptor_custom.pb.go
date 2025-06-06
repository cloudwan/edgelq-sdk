// Code generated by protoc-gen-goten-go
// File: edgelq/monitoring/proto/v3/monitored_resource_descriptor_custom.proto
// DO NOT EDIT!!!

package monitored_resource_descriptor_client

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
	monitored_resource_descriptor "github.com/cloudwan/edgelq-sdk/monitoring/resources/v3/monitored_resource_descriptor"
	view "github.com/cloudwan/goten-sdk/types/view"
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
	_ = &monitored_resource_descriptor.MonitoredResourceDescriptor{}
	_ = &fieldmaskpb.FieldMask{}
	_ = view.View(0)
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Request message for method
// [GetMonitoredResourceDescriptor][ntt.monitoring.v3.GetMonitoredResourceDescriptor]
type GetMonitoredResourceDescriptorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// The monitored resource descriptor to get.  The format is
	// `"monitoredResourceDescriptors/{resource_type}"`.
	// The `{resource_type}` is a predefined type, such as
	// `cloudsql_database`.
	Name *monitored_resource_descriptor.Name `protobuf:"bytes,3,opt,customtype=Name,name=name,proto3" json:"name,omitempty"`
	// A list of extra fields to be obtained for each response item on top of
	// fields defined by request field view
	FieldMask *monitored_resource_descriptor.MonitoredResourceDescriptor_FieldMask `protobuf:"bytes,6,opt,customtype=MonitoredResourceDescriptor_FieldMask,name=field_mask,json=fieldMask,proto3" json:"field_mask,omitempty"`
	// View defines list of standard response fields present in response items.
	// Additional fields can be amended by request field field_mask
	View view.View `protobuf:"varint,7,opt,name=view,proto3,enum=goten.types.View" json:"view,omitempty"`
}

func (m *GetMonitoredResourceDescriptorRequest) Reset() {
	*m = GetMonitoredResourceDescriptorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_monitoring_proto_v3_monitored_resource_descriptor_custom_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *GetMonitoredResourceDescriptorRequest) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*GetMonitoredResourceDescriptorRequest) ProtoMessage() {}

func (m *GetMonitoredResourceDescriptorRequest) ProtoReflect() preflect.Message {
	mi := &edgelq_monitoring_proto_v3_monitored_resource_descriptor_custom_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*GetMonitoredResourceDescriptorRequest) GotenMessage() {}

// Deprecated, Use GetMonitoredResourceDescriptorRequest.ProtoReflect.Descriptor instead.
func (*GetMonitoredResourceDescriptorRequest) Descriptor() ([]byte, []int) {
	return edgelq_monitoring_proto_v3_monitored_resource_descriptor_custom_proto_rawDescGZIP(), []int{0}
}

func (m *GetMonitoredResourceDescriptorRequest) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *GetMonitoredResourceDescriptorRequest) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *GetMonitoredResourceDescriptorRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *GetMonitoredResourceDescriptorRequest) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *GetMonitoredResourceDescriptorRequest) GetName() *monitored_resource_descriptor.Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *GetMonitoredResourceDescriptorRequest) GetFieldMask() *monitored_resource_descriptor.MonitoredResourceDescriptor_FieldMask {
	if m != nil {
		return m.FieldMask
	}
	return nil
}

func (m *GetMonitoredResourceDescriptorRequest) GetView() view.View {
	if m != nil {
		return m.View
	}
	return view.View_UNSPECIFIED
}

func (m *GetMonitoredResourceDescriptorRequest) SetName(fv *monitored_resource_descriptor.Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "GetMonitoredResourceDescriptorRequest"))
	}
	m.Name = fv
}

func (m *GetMonitoredResourceDescriptorRequest) SetFieldMask(fv *monitored_resource_descriptor.MonitoredResourceDescriptor_FieldMask) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "FieldMask", "GetMonitoredResourceDescriptorRequest"))
	}
	m.FieldMask = fv
}

func (m *GetMonitoredResourceDescriptorRequest) SetView(fv view.View) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "View", "GetMonitoredResourceDescriptorRequest"))
	}
	m.View = fv
}

// Request message for method
// [ListMonitoredResourceDescriptors][ntt.monitoring.v3.ListMonitoredResourceDescriptors]
type ListMonitoredResourceDescriptorsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// An optional [filter](/monitoring/api/v3/filters) describing
	// the descriptors to be returned.  The filter can reference
	// the descriptor's type and labels. For example, the
	// following filter returns only Google Compute Engine descriptors
	// that have an `id` label:
	//
	//     resource.type = starts_with("gce_") AND resource.label:id
	Filter *monitored_resource_descriptor.Filter `protobuf:"bytes,2,opt,customtype=Filter,name=filter,proto3" json:"filter,omitempty"`
	// A positive number that is the maximum number of results to return.
	PageSize int32                                  `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	OrderBy  *monitored_resource_descriptor.OrderBy `protobuf:"bytes,666,opt,customtype=OrderBy,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
	// If this field is not empty then it must contain the `nextPageToken` value
	// returned by a previous call to this method.  Using this field causes the
	// method to return additional results from the previous method call.
	PageToken *monitored_resource_descriptor.PagerCursor `protobuf:"bytes,4,opt,customtype=PagerCursor,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// A list of extra fields to be obtained for each response item on top of
	// fields defined by request field view
	FieldMask *monitored_resource_descriptor.MonitoredResourceDescriptor_FieldMask `protobuf:"bytes,6,opt,customtype=MonitoredResourceDescriptor_FieldMask,name=field_mask,json=fieldMask,proto3" json:"field_mask,omitempty"`
	// View defines list of standard response fields present in response items.
	// Additional fields can be amended by request field field_mask
	View view.View `protobuf:"varint,7,opt,name=view,proto3,enum=goten.types.View" json:"view,omitempty"`
	// Indicates if list response should contain total count and offset (fields
	// current_offset and total_results_count).
	IncludePagingInfo bool `protobuf:"varint,8,opt,name=include_paging_info,json=includePagingInfo,proto3" json:"include_paging_info,omitempty"`
}

func (m *ListMonitoredResourceDescriptorsRequest) Reset() {
	*m = ListMonitoredResourceDescriptorsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_monitoring_proto_v3_monitored_resource_descriptor_custom_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *ListMonitoredResourceDescriptorsRequest) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*ListMonitoredResourceDescriptorsRequest) ProtoMessage() {}

func (m *ListMonitoredResourceDescriptorsRequest) ProtoReflect() preflect.Message {
	mi := &edgelq_monitoring_proto_v3_monitored_resource_descriptor_custom_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*ListMonitoredResourceDescriptorsRequest) GotenMessage() {}

// Deprecated, Use ListMonitoredResourceDescriptorsRequest.ProtoReflect.Descriptor instead.
func (*ListMonitoredResourceDescriptorsRequest) Descriptor() ([]byte, []int) {
	return edgelq_monitoring_proto_v3_monitored_resource_descriptor_custom_proto_rawDescGZIP(), []int{1}
}

func (m *ListMonitoredResourceDescriptorsRequest) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *ListMonitoredResourceDescriptorsRequest) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *ListMonitoredResourceDescriptorsRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *ListMonitoredResourceDescriptorsRequest) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *ListMonitoredResourceDescriptorsRequest) GetFilter() *monitored_resource_descriptor.Filter {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *ListMonitoredResourceDescriptorsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return int32(0)
}

func (m *ListMonitoredResourceDescriptorsRequest) GetOrderBy() *monitored_resource_descriptor.OrderBy {
	if m != nil {
		return m.OrderBy
	}
	return nil
}

func (m *ListMonitoredResourceDescriptorsRequest) GetPageToken() *monitored_resource_descriptor.PagerCursor {
	if m != nil {
		return m.PageToken
	}
	return nil
}

func (m *ListMonitoredResourceDescriptorsRequest) GetFieldMask() *monitored_resource_descriptor.MonitoredResourceDescriptor_FieldMask {
	if m != nil {
		return m.FieldMask
	}
	return nil
}

func (m *ListMonitoredResourceDescriptorsRequest) GetView() view.View {
	if m != nil {
		return m.View
	}
	return view.View_UNSPECIFIED
}

func (m *ListMonitoredResourceDescriptorsRequest) GetIncludePagingInfo() bool {
	if m != nil {
		return m.IncludePagingInfo
	}
	return false
}

func (m *ListMonitoredResourceDescriptorsRequest) SetFilter(fv *monitored_resource_descriptor.Filter) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Filter", "ListMonitoredResourceDescriptorsRequest"))
	}
	m.Filter = fv
}

func (m *ListMonitoredResourceDescriptorsRequest) SetPageSize(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "PageSize", "ListMonitoredResourceDescriptorsRequest"))
	}
	m.PageSize = fv
}

func (m *ListMonitoredResourceDescriptorsRequest) SetOrderBy(fv *monitored_resource_descriptor.OrderBy) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "OrderBy", "ListMonitoredResourceDescriptorsRequest"))
	}
	m.OrderBy = fv
}

func (m *ListMonitoredResourceDescriptorsRequest) SetPageToken(fv *monitored_resource_descriptor.PagerCursor) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "PageToken", "ListMonitoredResourceDescriptorsRequest"))
	}
	m.PageToken = fv
}

func (m *ListMonitoredResourceDescriptorsRequest) SetFieldMask(fv *monitored_resource_descriptor.MonitoredResourceDescriptor_FieldMask) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "FieldMask", "ListMonitoredResourceDescriptorsRequest"))
	}
	m.FieldMask = fv
}

func (m *ListMonitoredResourceDescriptorsRequest) SetView(fv view.View) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "View", "ListMonitoredResourceDescriptorsRequest"))
	}
	m.View = fv
}

func (m *ListMonitoredResourceDescriptorsRequest) SetIncludePagingInfo(fv bool) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "IncludePagingInfo", "ListMonitoredResourceDescriptorsRequest"))
	}
	m.IncludePagingInfo = fv
}

// Response message for method
// [ListMonitoredResourceDescriptors][ntt.monitoring.v3.ListMonitoredResourceDescriptors]
type ListMonitoredResourceDescriptorsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// The monitored resource descriptors that are available to this project
	// and that match `filter`, if present.
	MonitoredResourceDescriptors []*monitored_resource_descriptor.MonitoredResourceDescriptor `protobuf:"bytes,1,rep,name=monitored_resource_descriptors,json=monitoredResourceDescriptors,proto3" json:"monitored_resource_descriptors,omitempty"`
	// If there are more results than have been returned, then this field is set
	// to a non-empty value.  To see the additional results,
	// use that value as `pageToken` in the next call to this method.
	NextPageToken *monitored_resource_descriptor.PagerCursor `protobuf:"bytes,2,opt,customtype=PagerCursor,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	PrevPageToken *monitored_resource_descriptor.PagerCursor `protobuf:"bytes,666,opt,customtype=PagerCursor,name=prev_page_token,json=prevPageToken,proto3" json:"prev_page_token,omitempty"`
	// Current offset from the first page (0 if no page tokens were given or
	// paging info was not requested). Page index can be computed from offset and
	// limit provided in a request.
	CurrentOffset int32 `protobuf:"varint,5,opt,name=current_offset,json=currentOffset,proto3" json:"current_offset,omitempty"`
	// Number of total MonitoresResourceDescriptors across all pages or 0, if
	// there are no items or paging info was not requested.
	TotalResultsCount int32 `protobuf:"varint,6,opt,name=total_results_count,json=totalResultsCount,proto3" json:"total_results_count,omitempty"`
}

func (m *ListMonitoredResourceDescriptorsResponse) Reset() {
	*m = ListMonitoredResourceDescriptorsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_monitoring_proto_v3_monitored_resource_descriptor_custom_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *ListMonitoredResourceDescriptorsResponse) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*ListMonitoredResourceDescriptorsResponse) ProtoMessage() {}

func (m *ListMonitoredResourceDescriptorsResponse) ProtoReflect() preflect.Message {
	mi := &edgelq_monitoring_proto_v3_monitored_resource_descriptor_custom_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*ListMonitoredResourceDescriptorsResponse) GotenMessage() {}

// Deprecated, Use ListMonitoredResourceDescriptorsResponse.ProtoReflect.Descriptor instead.
func (*ListMonitoredResourceDescriptorsResponse) Descriptor() ([]byte, []int) {
	return edgelq_monitoring_proto_v3_monitored_resource_descriptor_custom_proto_rawDescGZIP(), []int{2}
}

func (m *ListMonitoredResourceDescriptorsResponse) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *ListMonitoredResourceDescriptorsResponse) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *ListMonitoredResourceDescriptorsResponse) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *ListMonitoredResourceDescriptorsResponse) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *ListMonitoredResourceDescriptorsResponse) GetMonitoredResourceDescriptors() []*monitored_resource_descriptor.MonitoredResourceDescriptor {
	if m != nil {
		return m.MonitoredResourceDescriptors
	}
	return nil
}

func (m *ListMonitoredResourceDescriptorsResponse) GetNextPageToken() *monitored_resource_descriptor.PagerCursor {
	if m != nil {
		return m.NextPageToken
	}
	return nil
}

func (m *ListMonitoredResourceDescriptorsResponse) GetPrevPageToken() *monitored_resource_descriptor.PagerCursor {
	if m != nil {
		return m.PrevPageToken
	}
	return nil
}

func (m *ListMonitoredResourceDescriptorsResponse) GetCurrentOffset() int32 {
	if m != nil {
		return m.CurrentOffset
	}
	return int32(0)
}

func (m *ListMonitoredResourceDescriptorsResponse) GetTotalResultsCount() int32 {
	if m != nil {
		return m.TotalResultsCount
	}
	return int32(0)
}

func (m *ListMonitoredResourceDescriptorsResponse) SetMonitoredResourceDescriptors(fv []*monitored_resource_descriptor.MonitoredResourceDescriptor) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "MonitoredResourceDescriptors", "ListMonitoredResourceDescriptorsResponse"))
	}
	m.MonitoredResourceDescriptors = fv
}

func (m *ListMonitoredResourceDescriptorsResponse) SetNextPageToken(fv *monitored_resource_descriptor.PagerCursor) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "NextPageToken", "ListMonitoredResourceDescriptorsResponse"))
	}
	m.NextPageToken = fv
}

func (m *ListMonitoredResourceDescriptorsResponse) SetPrevPageToken(fv *monitored_resource_descriptor.PagerCursor) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "PrevPageToken", "ListMonitoredResourceDescriptorsResponse"))
	}
	m.PrevPageToken = fv
}

func (m *ListMonitoredResourceDescriptorsResponse) SetCurrentOffset(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "CurrentOffset", "ListMonitoredResourceDescriptorsResponse"))
	}
	m.CurrentOffset = fv
}

func (m *ListMonitoredResourceDescriptorsResponse) SetTotalResultsCount(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "TotalResultsCount", "ListMonitoredResourceDescriptorsResponse"))
	}
	m.TotalResultsCount = fv
}

var edgelq_monitoring_proto_v3_monitored_resource_descriptor_custom_proto preflect.FileDescriptor

var edgelq_monitoring_proto_v3_monitored_resource_descriptor_custom_proto_rawDesc = []byte{
	0x0a, 0x45, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x69, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x33, 0x2f, 0x6d, 0x6f, 0x6e,
	0x69, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x6e, 0x74, 0x74, 0x2e, 0x6d, 0x6f, 0x6e,
	0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65,
	0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f,
	0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x65, 0x64, 0x67,
	0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x65,
	0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2f, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3e, 0x65,
	0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x33, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f,
	0x72, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe7, 0x01,
	0x0a, 0x25, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x23, 0xb2, 0xda, 0x21, 0x1f, 0x0a, 0x1d, 0x0a, 0x1b, 0x4d,
	0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x5e, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b,
	0x42, 0x23, 0xb2, 0xda, 0x21, 0x1f, 0x32, 0x1d, 0x0a, 0x1b, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f,
	0x72, 0x65, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b,
	0x12, 0x25, 0x0a, 0x04, 0x76, 0x69, 0x65, 0x77, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11,
	0x2e, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x56, 0x69, 0x65,
	0x77, 0x52, 0x04, 0x76, 0x69, 0x65, 0x77, 0x22, 0xbf, 0x03, 0x0a, 0x27, 0x4c, 0x69, 0x73, 0x74,
	0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x3b, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x23, 0xb2, 0xda, 0x21, 0x1f, 0x1a, 0x1d, 0x0a, 0x1b, 0x4d, 0x6f, 0x6e,
	0x69, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x3f, 0x0a,
	0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x18, 0x9a, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x23, 0xb2, 0xda, 0x21, 0x1f, 0x2a, 0x1d, 0x0a, 0x1b, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f,
	0x72, 0x65, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x12, 0x42,
	0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x23, 0xb2, 0xda, 0x21, 0x1f, 0x22, 0x1d, 0x0a, 0x1b, 0x4d, 0x6f, 0x6e, 0x69,
	0x74, 0x6f, 0x72, 0x65, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x5e, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61,
	0x73, 0x6b, 0x42, 0x23, 0xb2, 0xda, 0x21, 0x1f, 0x32, 0x1d, 0x0a, 0x1b, 0x4d, 0x6f, 0x6e, 0x69,
	0x74, 0x6f, 0x72, 0x65, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61,
	0x73, 0x6b, 0x12, 0x25, 0x0a, 0x04, 0x76, 0x69, 0x65, 0x77, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x11, 0x2e, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x56,
	0x69, 0x65, 0x77, 0x52, 0x04, 0x76, 0x69, 0x65, 0x77, 0x12, 0x2e, 0x0a, 0x13, 0x69, 0x6e, 0x63,
	0x6c, 0x75, 0x64, 0x65, 0x5f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x50,
	0x61, 0x67, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x92, 0x03, 0x0a, 0x28, 0x4c, 0x69,
	0x73, 0x74, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x74, 0x0a, 0x1e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f,
	0x72, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e,
	0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x33, 0x2e, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x1c,
	0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x4b, 0x0a, 0x0f,
	0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x23, 0xb2, 0xda, 0x21, 0x1f, 0x22, 0x1d, 0x0a, 0x1b, 0x4d,
	0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74,
	0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x4c, 0x0a, 0x0f, 0x70, 0x72, 0x65,
	0x76, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x9a, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x23, 0xb2, 0xda, 0x21, 0x1f, 0x22, 0x1d, 0x0a, 0x1b, 0x4d, 0x6f, 0x6e,
	0x69, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x0d, 0x70, 0x72, 0x65, 0x76, 0x50, 0x61,
	0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x2e,
	0x0a, 0x13, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0xb8,
	0x01, 0x0a, 0x18, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74,
	0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x33, 0x42, 0x26, 0x4d, 0x6f, 0x6e,
	0x69, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x00, 0x5a, 0x72, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c,
	0x71, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2f, 0x76, 0x33, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x65, 0x64,
	0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x6f, 0x72, 0x3b, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x5f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x6f, 0x72, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	edgelq_monitoring_proto_v3_monitored_resource_descriptor_custom_proto_rawDescOnce sync.Once
	edgelq_monitoring_proto_v3_monitored_resource_descriptor_custom_proto_rawDescData = edgelq_monitoring_proto_v3_monitored_resource_descriptor_custom_proto_rawDesc
)

func edgelq_monitoring_proto_v3_monitored_resource_descriptor_custom_proto_rawDescGZIP() []byte {
	edgelq_monitoring_proto_v3_monitored_resource_descriptor_custom_proto_rawDescOnce.Do(func() {
		edgelq_monitoring_proto_v3_monitored_resource_descriptor_custom_proto_rawDescData = protoimpl.X.CompressGZIP(edgelq_monitoring_proto_v3_monitored_resource_descriptor_custom_proto_rawDescData)
	})
	return edgelq_monitoring_proto_v3_monitored_resource_descriptor_custom_proto_rawDescData
}

var edgelq_monitoring_proto_v3_monitored_resource_descriptor_custom_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var edgelq_monitoring_proto_v3_monitored_resource_descriptor_custom_proto_goTypes = []interface{}{
	(*GetMonitoredResourceDescriptorRequest)(nil),                               // 0: ntt.monitoring.v3.GetMonitoredResourceDescriptorRequest
	(*ListMonitoredResourceDescriptorsRequest)(nil),                             // 1: ntt.monitoring.v3.ListMonitoredResourceDescriptorsRequest
	(*ListMonitoredResourceDescriptorsResponse)(nil),                            // 2: ntt.monitoring.v3.ListMonitoredResourceDescriptorsResponse
	(*monitored_resource_descriptor.MonitoredResourceDescriptor_FieldMask)(nil), // 3: ntt.monitoring.v3.MonitoredResourceDescriptor_FieldMask
	(view.View)(0), // 4: goten.types.View
	(*monitored_resource_descriptor.MonitoredResourceDescriptor)(nil), // 5: ntt.monitoring.v3.MonitoredResourceDescriptor
}
var edgelq_monitoring_proto_v3_monitored_resource_descriptor_custom_proto_depIdxs = []int32{
	3, // 0: ntt.monitoring.v3.GetMonitoredResourceDescriptorRequest.field_mask:type_name -> ntt.monitoring.v3.MonitoredResourceDescriptor_FieldMask
	4, // 1: ntt.monitoring.v3.GetMonitoredResourceDescriptorRequest.view:type_name -> goten.types.View
	3, // 2: ntt.monitoring.v3.ListMonitoredResourceDescriptorsRequest.field_mask:type_name -> ntt.monitoring.v3.MonitoredResourceDescriptor_FieldMask
	4, // 3: ntt.monitoring.v3.ListMonitoredResourceDescriptorsRequest.view:type_name -> goten.types.View
	5, // 4: ntt.monitoring.v3.ListMonitoredResourceDescriptorsResponse.monitored_resource_descriptors:type_name -> ntt.monitoring.v3.MonitoredResourceDescriptor
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { edgelq_monitoring_proto_v3_monitored_resource_descriptor_custom_proto_init() }
func edgelq_monitoring_proto_v3_monitored_resource_descriptor_custom_proto_init() {
	if edgelq_monitoring_proto_v3_monitored_resource_descriptor_custom_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		edgelq_monitoring_proto_v3_monitored_resource_descriptor_custom_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMonitoredResourceDescriptorRequest); i {
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
		edgelq_monitoring_proto_v3_monitored_resource_descriptor_custom_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMonitoredResourceDescriptorsRequest); i {
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
		edgelq_monitoring_proto_v3_monitored_resource_descriptor_custom_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMonitoredResourceDescriptorsResponse); i {
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
			RawDescriptor: edgelq_monitoring_proto_v3_monitored_resource_descriptor_custom_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           edgelq_monitoring_proto_v3_monitored_resource_descriptor_custom_proto_goTypes,
		DependencyIndexes: edgelq_monitoring_proto_v3_monitored_resource_descriptor_custom_proto_depIdxs,
		MessageInfos:      edgelq_monitoring_proto_v3_monitored_resource_descriptor_custom_proto_msgTypes,
	}.Build()
	edgelq_monitoring_proto_v3_monitored_resource_descriptor_custom_proto = out.File
	edgelq_monitoring_proto_v3_monitored_resource_descriptor_custom_proto_rawDesc = nil
	edgelq_monitoring_proto_v3_monitored_resource_descriptor_custom_proto_goTypes = nil
	edgelq_monitoring_proto_v3_monitored_resource_descriptor_custom_proto_depIdxs = nil
}
