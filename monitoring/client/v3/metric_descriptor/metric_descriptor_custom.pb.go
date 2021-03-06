// Code generated by protoc-gen-goten-go
// File: edgelq/monitoring/proto/v3/metric_descriptor_custom.proto
// DO NOT EDIT!!!

package metric_descriptor_client

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
	metric_descriptor "github.com/cloudwan/edgelq-sdk/monitoring/resources/v3/metric_descriptor"
	project "github.com/cloudwan/edgelq-sdk/monitoring/resources/v3/project"
	view "github.com/cloudwan/goten-sdk/runtime/api/view"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
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
	_ = &metric_descriptor.MetricDescriptor{}
	_ = &project.Project{}
	_ = &field_mask.FieldMask{}
	_ = view.View(0)
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Request message for method
// [ListMetricDescriptors][ntt.monitoring.v3.ListMetricDescriptors]
type ListMetricDescriptorsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// The project on which to execute the request. The format is
	// `"projects/{project_id_or_number}"`.
	// TODO: enable when multi-parents are supported
	Parent *metric_descriptor.ParentName `protobuf:"bytes,5,opt,customtype=ParentName,name=parent,proto3" json:"parent,omitempty" firestore:"parent"`
	// If this field is empty, all custom and
	// system-defined metric descriptors are returned.
	// Otherwise, the [filter](/monitoring/api/v3/filters)
	// specifies which metric descriptors are to be
	// returned. For example, the following filter matches all
	// [custom metrics](/monitoring/custom-metrics):
	//
	//     metric.type = starts_with("custom.googleapis.com/")
	Filter *metric_descriptor.Filter `protobuf:"bytes,2,opt,customtype=Filter,name=filter,proto3" json:"filter,omitempty" firestore:"filter"`
	// A positive number that is the maximum number of results to return.
	PageSize int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty" firestore:"pageSize"`
	// If this field is not empty then it must contain the `nextPageToken` value
	// returned by a previous call to this method.  Using this field causes the
	// method to return additional results from the previous method call.
	PageToken *metric_descriptor.PagerCursor `protobuf:"bytes,4,opt,customtype=PagerCursor,name=page_token,json=pageToken,proto3" json:"page_token,omitempty" firestore:"pageToken"`
	// Order By -
	// https://cloud.google.com/apis/design/design_patterns#list_pagination
	OrderBy *metric_descriptor.OrderBy `protobuf:"bytes,666,opt,customtype=OrderBy,name=order_by,json=orderBy,proto3" json:"order_by,omitempty" firestore:"orderBy"`
	// A list of extra fields to be obtained for each response item on top of
	// fields defined by request field view
	FieldMask *metric_descriptor.MetricDescriptor_FieldMask `protobuf:"bytes,6,opt,customtype=MetricDescriptor_FieldMask,name=field_mask,json=fieldMask,proto3" json:"field_mask,omitempty" firestore:"fieldMask"`
	// View defines list of standard response fields present in response items.
	// Additional fields can be amended by request field field_mask
	View view.View `protobuf:"varint,7,opt,name=view,proto3,enum=goten.view.View" json:"view,omitempty" firestore:"view"`
}

func (m *ListMetricDescriptorsRequest) Reset() {
	*m = ListMetricDescriptorsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_monitoring_proto_v3_metric_descriptor_custom_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *ListMetricDescriptorsRequest) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*ListMetricDescriptorsRequest) ProtoMessage() {}

func (m *ListMetricDescriptorsRequest) ProtoReflect() preflect.Message {
	mi := &edgelq_monitoring_proto_v3_metric_descriptor_custom_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*ListMetricDescriptorsRequest) GotenMessage() {}

// Deprecated, Use ListMetricDescriptorsRequest.ProtoReflect.Descriptor instead.
func (*ListMetricDescriptorsRequest) Descriptor() ([]byte, []int) {
	return edgelq_monitoring_proto_v3_metric_descriptor_custom_proto_rawDescGZIP(), []int{0}
}

func (m *ListMetricDescriptorsRequest) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *ListMetricDescriptorsRequest) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *ListMetricDescriptorsRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *ListMetricDescriptorsRequest) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *ListMetricDescriptorsRequest) GetParent() *metric_descriptor.ParentName {
	if m != nil {
		return m.Parent
	}
	return nil
}

func (m *ListMetricDescriptorsRequest) GetFilter() *metric_descriptor.Filter {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *ListMetricDescriptorsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return int32(0)
}

func (m *ListMetricDescriptorsRequest) GetPageToken() *metric_descriptor.PagerCursor {
	if m != nil {
		return m.PageToken
	}
	return nil
}

func (m *ListMetricDescriptorsRequest) GetOrderBy() *metric_descriptor.OrderBy {
	if m != nil {
		return m.OrderBy
	}
	return nil
}

func (m *ListMetricDescriptorsRequest) GetFieldMask() *metric_descriptor.MetricDescriptor_FieldMask {
	if m != nil {
		return m.FieldMask
	}
	return nil
}

func (m *ListMetricDescriptorsRequest) GetView() view.View {
	if m != nil {
		return m.View
	}
	return view.View_UNSPECIFIED
}

func (m *ListMetricDescriptorsRequest) SetParent(fv *metric_descriptor.ParentName) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Parent", "ListMetricDescriptorsRequest"))
	}
	m.Parent = fv
}

func (m *ListMetricDescriptorsRequest) SetFilter(fv *metric_descriptor.Filter) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Filter", "ListMetricDescriptorsRequest"))
	}
	m.Filter = fv
}

func (m *ListMetricDescriptorsRequest) SetPageSize(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "PageSize", "ListMetricDescriptorsRequest"))
	}
	m.PageSize = fv
}

func (m *ListMetricDescriptorsRequest) SetPageToken(fv *metric_descriptor.PagerCursor) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "PageToken", "ListMetricDescriptorsRequest"))
	}
	m.PageToken = fv
}

func (m *ListMetricDescriptorsRequest) SetOrderBy(fv *metric_descriptor.OrderBy) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "OrderBy", "ListMetricDescriptorsRequest"))
	}
	m.OrderBy = fv
}

func (m *ListMetricDescriptorsRequest) SetFieldMask(fv *metric_descriptor.MetricDescriptor_FieldMask) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "FieldMask", "ListMetricDescriptorsRequest"))
	}
	m.FieldMask = fv
}

func (m *ListMetricDescriptorsRequest) SetView(fv view.View) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "View", "ListMetricDescriptorsRequest"))
	}
	m.View = fv
}

// Response message for method
// [ListMetricDescriptors][ntt.monitoring.v3.ListMetricDescriptors]
type ListMetricDescriptorsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// The metric descriptors that are available to the project
	// and that match the value of `filter`, if present.
	MetricDescriptors []*metric_descriptor.MetricDescriptor `protobuf:"bytes,1,rep,name=metric_descriptors,json=metricDescriptors,proto3" json:"metric_descriptors,omitempty" firestore:"metricDescriptors"`
	// If there are more results than have been returned, then this field is set
	// to a non-empty value.  To see the additional results,
	// use that value as `pageToken` in the next call to this method.
	NextPageToken *metric_descriptor.PagerCursor `protobuf:"bytes,2,opt,customtype=PagerCursor,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty" firestore:"nextPageToken"`
	PrevPageToken *metric_descriptor.PagerCursor `protobuf:"bytes,666,opt,customtype=PagerCursor,name=prev_page_token,json=prevPageToken,proto3" json:"prev_page_token,omitempty" firestore:"prevPageToken"`
}

func (m *ListMetricDescriptorsResponse) Reset() {
	*m = ListMetricDescriptorsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_monitoring_proto_v3_metric_descriptor_custom_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *ListMetricDescriptorsResponse) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*ListMetricDescriptorsResponse) ProtoMessage() {}

func (m *ListMetricDescriptorsResponse) ProtoReflect() preflect.Message {
	mi := &edgelq_monitoring_proto_v3_metric_descriptor_custom_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*ListMetricDescriptorsResponse) GotenMessage() {}

// Deprecated, Use ListMetricDescriptorsResponse.ProtoReflect.Descriptor instead.
func (*ListMetricDescriptorsResponse) Descriptor() ([]byte, []int) {
	return edgelq_monitoring_proto_v3_metric_descriptor_custom_proto_rawDescGZIP(), []int{1}
}

func (m *ListMetricDescriptorsResponse) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *ListMetricDescriptorsResponse) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *ListMetricDescriptorsResponse) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *ListMetricDescriptorsResponse) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *ListMetricDescriptorsResponse) GetMetricDescriptors() []*metric_descriptor.MetricDescriptor {
	if m != nil {
		return m.MetricDescriptors
	}
	return nil
}

func (m *ListMetricDescriptorsResponse) GetNextPageToken() *metric_descriptor.PagerCursor {
	if m != nil {
		return m.NextPageToken
	}
	return nil
}

func (m *ListMetricDescriptorsResponse) GetPrevPageToken() *metric_descriptor.PagerCursor {
	if m != nil {
		return m.PrevPageToken
	}
	return nil
}

func (m *ListMetricDescriptorsResponse) SetMetricDescriptors(fv []*metric_descriptor.MetricDescriptor) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "MetricDescriptors", "ListMetricDescriptorsResponse"))
	}
	m.MetricDescriptors = fv
}

func (m *ListMetricDescriptorsResponse) SetNextPageToken(fv *metric_descriptor.PagerCursor) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "NextPageToken", "ListMetricDescriptorsResponse"))
	}
	m.NextPageToken = fv
}

func (m *ListMetricDescriptorsResponse) SetPrevPageToken(fv *metric_descriptor.PagerCursor) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "PrevPageToken", "ListMetricDescriptorsResponse"))
	}
	m.PrevPageToken = fv
}

// Request message for method
// [GetMetricDescriptor][ntt.monitoring.v3.GetMetricDescriptor]
type GetMetricDescriptorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// The metric descriptor on which to execute the request. The format is
	// `"projects/{project_id_or_number}/metricDescriptors/{metric_id}"`.
	// An example value of `{metric_id}` is
	// `"compute.googleapis.com/instance/disk/read_bytes_count"`.
	Name *metric_descriptor.Reference `protobuf:"bytes,3,opt,customtype=Reference,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// A list of extra fields to be obtained for each response item on top of
	// fields defined by request field view
	FieldMask *metric_descriptor.MetricDescriptor_FieldMask `protobuf:"bytes,4,opt,customtype=MetricDescriptor_FieldMask,name=field_mask,json=fieldMask,proto3" json:"field_mask,omitempty" firestore:"fieldMask"`
	// View defines list of standard response fields present in response items.
	// Additional fields can be amended by request field field_mask
	View view.View `protobuf:"varint,5,opt,name=view,proto3,enum=goten.view.View" json:"view,omitempty" firestore:"view"`
}

func (m *GetMetricDescriptorRequest) Reset() {
	*m = GetMetricDescriptorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_monitoring_proto_v3_metric_descriptor_custom_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *GetMetricDescriptorRequest) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*GetMetricDescriptorRequest) ProtoMessage() {}

func (m *GetMetricDescriptorRequest) ProtoReflect() preflect.Message {
	mi := &edgelq_monitoring_proto_v3_metric_descriptor_custom_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*GetMetricDescriptorRequest) GotenMessage() {}

// Deprecated, Use GetMetricDescriptorRequest.ProtoReflect.Descriptor instead.
func (*GetMetricDescriptorRequest) Descriptor() ([]byte, []int) {
	return edgelq_monitoring_proto_v3_metric_descriptor_custom_proto_rawDescGZIP(), []int{2}
}

func (m *GetMetricDescriptorRequest) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *GetMetricDescriptorRequest) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *GetMetricDescriptorRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *GetMetricDescriptorRequest) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *GetMetricDescriptorRequest) GetName() *metric_descriptor.Reference {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *GetMetricDescriptorRequest) GetFieldMask() *metric_descriptor.MetricDescriptor_FieldMask {
	if m != nil {
		return m.FieldMask
	}
	return nil
}

func (m *GetMetricDescriptorRequest) GetView() view.View {
	if m != nil {
		return m.View
	}
	return view.View_UNSPECIFIED
}

func (m *GetMetricDescriptorRequest) SetName(fv *metric_descriptor.Reference) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "GetMetricDescriptorRequest"))
	}
	m.Name = fv
}

func (m *GetMetricDescriptorRequest) SetFieldMask(fv *metric_descriptor.MetricDescriptor_FieldMask) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "FieldMask", "GetMetricDescriptorRequest"))
	}
	m.FieldMask = fv
}

func (m *GetMetricDescriptorRequest) SetView(fv view.View) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "View", "GetMetricDescriptorRequest"))
	}
	m.View = fv
}

// Request message for method
// [CreateMetricDescriptor][ntt.monitoring.v3.CreateMetricDescriptor]
type CreateMetricDescriptorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// The project on which to execute the request. The format is
	// `"projects/{project_id_or_number}"`.
	// TODO: enable when multi-parents are supported
	Parent *metric_descriptor.ParentReference `protobuf:"bytes,3,opt,customtype=ParentReference,name=parent,proto3" json:"parent,omitempty" firestore:"parent"`
	// The new [custom metric](/monitoring/custom-metrics)
	// descriptor.
	MetricDescriptor *metric_descriptor.MetricDescriptor `protobuf:"bytes,2,opt,name=metric_descriptor,json=metricDescriptor,proto3" json:"metric_descriptor,omitempty" firestore:"metricDescriptor"`
}

func (m *CreateMetricDescriptorRequest) Reset() {
	*m = CreateMetricDescriptorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_monitoring_proto_v3_metric_descriptor_custom_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *CreateMetricDescriptorRequest) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*CreateMetricDescriptorRequest) ProtoMessage() {}

func (m *CreateMetricDescriptorRequest) ProtoReflect() preflect.Message {
	mi := &edgelq_monitoring_proto_v3_metric_descriptor_custom_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*CreateMetricDescriptorRequest) GotenMessage() {}

// Deprecated, Use CreateMetricDescriptorRequest.ProtoReflect.Descriptor instead.
func (*CreateMetricDescriptorRequest) Descriptor() ([]byte, []int) {
	return edgelq_monitoring_proto_v3_metric_descriptor_custom_proto_rawDescGZIP(), []int{3}
}

func (m *CreateMetricDescriptorRequest) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *CreateMetricDescriptorRequest) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *CreateMetricDescriptorRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *CreateMetricDescriptorRequest) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *CreateMetricDescriptorRequest) GetParent() *metric_descriptor.ParentReference {
	if m != nil {
		return m.Parent
	}
	return nil
}

func (m *CreateMetricDescriptorRequest) GetMetricDescriptor() *metric_descriptor.MetricDescriptor {
	if m != nil {
		return m.MetricDescriptor
	}
	return nil
}

func (m *CreateMetricDescriptorRequest) SetParent(fv *metric_descriptor.ParentReference) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Parent", "CreateMetricDescriptorRequest"))
	}
	m.Parent = fv
}

func (m *CreateMetricDescriptorRequest) SetMetricDescriptor(fv *metric_descriptor.MetricDescriptor) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "MetricDescriptor", "CreateMetricDescriptorRequest"))
	}
	m.MetricDescriptor = fv
}

// Request message for method
// [DeleteMetricDescriptor][ntt.monitoring.v3.DeleteMetricDescriptor]
type DeleteMetricDescriptorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// The metric descriptor on which to execute the request. The format is
	// `"projects/{project_id_or_number}/metricDescriptors/{metric_id}"`.
	// An example of `{metric_id}` is:
	// `"custom.googleapis.com/my_test_metric"`.
	Name *metric_descriptor.Reference `protobuf:"bytes,3,opt,customtype=Reference,name=name,proto3" json:"name,omitempty" firestore:"name"`
}

func (m *DeleteMetricDescriptorRequest) Reset() {
	*m = DeleteMetricDescriptorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_monitoring_proto_v3_metric_descriptor_custom_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *DeleteMetricDescriptorRequest) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*DeleteMetricDescriptorRequest) ProtoMessage() {}

func (m *DeleteMetricDescriptorRequest) ProtoReflect() preflect.Message {
	mi := &edgelq_monitoring_proto_v3_metric_descriptor_custom_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*DeleteMetricDescriptorRequest) GotenMessage() {}

// Deprecated, Use DeleteMetricDescriptorRequest.ProtoReflect.Descriptor instead.
func (*DeleteMetricDescriptorRequest) Descriptor() ([]byte, []int) {
	return edgelq_monitoring_proto_v3_metric_descriptor_custom_proto_rawDescGZIP(), []int{4}
}

func (m *DeleteMetricDescriptorRequest) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *DeleteMetricDescriptorRequest) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *DeleteMetricDescriptorRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *DeleteMetricDescriptorRequest) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *DeleteMetricDescriptorRequest) GetName() *metric_descriptor.Reference {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *DeleteMetricDescriptorRequest) SetName(fv *metric_descriptor.Reference) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "DeleteMetricDescriptorRequest"))
	}
	m.Name = fv
}

var edgelq_monitoring_proto_v3_metric_descriptor_custom_proto preflect.FileDescriptor

var edgelq_monitoring_proto_v3_metric_descriptor_custom_proto_rawDesc = []byte{
	0x0a, 0x39, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x69, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x33, 0x2f, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x5f, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x6e, 0x74, 0x74,
	0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x1a, 0x20,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x20, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x24, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x75, 0x64, 0x69,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x32, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f,
	0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x76, 0x33, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8d, 0x03, 0x0a, 0x1c,
	0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x06,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1c, 0xb2, 0xda,
	0x21, 0x14, 0x3a, 0x12, 0x0a, 0x10, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0xba, 0x9d, 0x22, 0x00, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x12, 0x30, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x18, 0xb2, 0xda, 0x21, 0x14, 0x1a, 0x12, 0x0a, 0x10, 0x4d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x06, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x37, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xb2, 0xda, 0x21, 0x14, 0x22, 0x12, 0x0a, 0x10, 0x4d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x34, 0x0a, 0x08, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x18, 0x9a, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xb2,
	0xda, 0x21, 0x14, 0x2a, 0x12, 0x0a, 0x10, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79,
	0x12, 0x53, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b,
	0x42, 0x18, 0xb2, 0xda, 0x21, 0x14, 0x32, 0x12, 0x0a, 0x10, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x24, 0x0a, 0x04, 0x76, 0x69, 0x65, 0x77, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x76, 0x69, 0x65, 0x77,
	0x2e, 0x56, 0x69, 0x65, 0x77, 0x52, 0x04, 0x76, 0x69, 0x65, 0x77, 0x22, 0xf8, 0x01, 0x0a, 0x1d,
	0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x52, 0x0a,
	0x12, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x6f, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6e, 0x74, 0x74, 0x2e,
	0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x4d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x11,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72,
	0x73, 0x12, 0x40, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xb2, 0xda, 0x21, 0x14,
	0x22, 0x12, 0x0a, 0x10, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x6f, 0x72, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x41, 0x0a, 0x0f, 0x70, 0x72, 0x65, 0x76, 0x5f, 0x70, 0x61, 0x67, 0x65,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x9a, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xb2,
	0xda, 0x21, 0x14, 0x22, 0x12, 0x0a, 0x10, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x0d, 0x70, 0x72, 0x65, 0x76, 0x50, 0x61, 0x67,
	0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xc5, 0x01, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x4d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x18, 0xb2, 0xda, 0x21, 0x14, 0x12, 0x12, 0x0a, 0x10, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x53, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73,
	0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d,
	0x61, 0x73, 0x6b, 0x42, 0x18, 0xb2, 0xda, 0x21, 0x14, 0x32, 0x12, 0x0a, 0x10, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x09, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x24, 0x0a, 0x04, 0x76, 0x69, 0x65, 0x77,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x76,
	0x69, 0x65, 0x77, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x52, 0x04, 0x76, 0x69, 0x65, 0x77, 0x22, 0xc9,
	0x01, 0x0a, 0x1d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x30, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x18, 0xb2, 0xda, 0x21, 0x14, 0x42, 0x12, 0x0a, 0x10, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x12, 0x50, 0x0a, 0x11, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e,
	0x6e, 0x74, 0x74, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76,
	0x33, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x6f, 0x72, 0x52, 0x10, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x6f, 0x72, 0x3a, 0x24, 0xc2, 0x85, 0x2c, 0x20, 0x32, 0x06, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x3a, 0x16, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x5f, 0x0a, 0x1d, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xb2, 0xda, 0x21, 0x14, 0x12,
	0x12, 0x0a, 0x10, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x6f, 0x72, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x3a, 0x10, 0xc2, 0x85, 0x2c, 0x0c, 0x32,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x3a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0xe5, 0x01, 0xe8, 0xde,
	0x21, 0x01, 0x0a, 0x18, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x6d, 0x6f, 0x6e, 0x69,
	0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x33, 0x42, 0x1b, 0x4d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x00, 0x5a, 0x5a, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e,
	0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69,
	0x6e, 0x67, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x33, 0x2f, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x3b, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72,
	0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0xd2, 0x84, 0xd1, 0x02, 0x47, 0x0a, 0x0d, 0x61, 0x75,
	0x64, 0x69, 0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x12, 0x36, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e,
	0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69,
	0x6e, 0x67, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73,
	0x2f, 0x76, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	edgelq_monitoring_proto_v3_metric_descriptor_custom_proto_rawDescOnce sync.Once
	edgelq_monitoring_proto_v3_metric_descriptor_custom_proto_rawDescData = edgelq_monitoring_proto_v3_metric_descriptor_custom_proto_rawDesc
)

func edgelq_monitoring_proto_v3_metric_descriptor_custom_proto_rawDescGZIP() []byte {
	edgelq_monitoring_proto_v3_metric_descriptor_custom_proto_rawDescOnce.Do(func() {
		edgelq_monitoring_proto_v3_metric_descriptor_custom_proto_rawDescData = protoimpl.X.CompressGZIP(edgelq_monitoring_proto_v3_metric_descriptor_custom_proto_rawDescData)
	})
	return edgelq_monitoring_proto_v3_metric_descriptor_custom_proto_rawDescData
}

var edgelq_monitoring_proto_v3_metric_descriptor_custom_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var edgelq_monitoring_proto_v3_metric_descriptor_custom_proto_goTypes = []interface{}{
	(*ListMetricDescriptorsRequest)(nil),                 // 0: ntt.monitoring.v3.ListMetricDescriptorsRequest
	(*ListMetricDescriptorsResponse)(nil),                // 1: ntt.monitoring.v3.ListMetricDescriptorsResponse
	(*GetMetricDescriptorRequest)(nil),                   // 2: ntt.monitoring.v3.GetMetricDescriptorRequest
	(*CreateMetricDescriptorRequest)(nil),                // 3: ntt.monitoring.v3.CreateMetricDescriptorRequest
	(*DeleteMetricDescriptorRequest)(nil),                // 4: ntt.monitoring.v3.DeleteMetricDescriptorRequest
	(*metric_descriptor.MetricDescriptor_FieldMask)(nil), // 5: ntt.monitoring.v3.MetricDescriptor_FieldMask
	(view.View)(0), // 6: goten.view.View
	(*metric_descriptor.MetricDescriptor)(nil), // 7: ntt.monitoring.v3.MetricDescriptor
}
var edgelq_monitoring_proto_v3_metric_descriptor_custom_proto_depIdxs = []int32{
	5, // 0: ntt.monitoring.v3.ListMetricDescriptorsRequest.field_mask:type_name -> ntt.monitoring.v3.MetricDescriptor_FieldMask
	6, // 1: ntt.monitoring.v3.ListMetricDescriptorsRequest.view:type_name -> goten.view.View
	7, // 2: ntt.monitoring.v3.ListMetricDescriptorsResponse.metric_descriptors:type_name -> ntt.monitoring.v3.MetricDescriptor
	5, // 3: ntt.monitoring.v3.GetMetricDescriptorRequest.field_mask:type_name -> ntt.monitoring.v3.MetricDescriptor_FieldMask
	6, // 4: ntt.monitoring.v3.GetMetricDescriptorRequest.view:type_name -> goten.view.View
	7, // 5: ntt.monitoring.v3.CreateMetricDescriptorRequest.metric_descriptor:type_name -> ntt.monitoring.v3.MetricDescriptor
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { edgelq_monitoring_proto_v3_metric_descriptor_custom_proto_init() }
func edgelq_monitoring_proto_v3_metric_descriptor_custom_proto_init() {
	if edgelq_monitoring_proto_v3_metric_descriptor_custom_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		edgelq_monitoring_proto_v3_metric_descriptor_custom_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMetricDescriptorsRequest); i {
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
		edgelq_monitoring_proto_v3_metric_descriptor_custom_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMetricDescriptorsResponse); i {
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
		edgelq_monitoring_proto_v3_metric_descriptor_custom_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMetricDescriptorRequest); i {
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
		edgelq_monitoring_proto_v3_metric_descriptor_custom_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMetricDescriptorRequest); i {
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
		edgelq_monitoring_proto_v3_metric_descriptor_custom_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteMetricDescriptorRequest); i {
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
			RawDescriptor: edgelq_monitoring_proto_v3_metric_descriptor_custom_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           edgelq_monitoring_proto_v3_metric_descriptor_custom_proto_goTypes,
		DependencyIndexes: edgelq_monitoring_proto_v3_metric_descriptor_custom_proto_depIdxs,
		MessageInfos:      edgelq_monitoring_proto_v3_metric_descriptor_custom_proto_msgTypes,
	}.Build()
	edgelq_monitoring_proto_v3_metric_descriptor_custom_proto = out.File
	edgelq_monitoring_proto_v3_metric_descriptor_custom_proto_rawDesc = nil
	edgelq_monitoring_proto_v3_metric_descriptor_custom_proto_goTypes = nil
	edgelq_monitoring_proto_v3_metric_descriptor_custom_proto_depIdxs = nil
}
