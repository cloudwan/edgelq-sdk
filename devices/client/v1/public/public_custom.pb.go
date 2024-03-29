// Code generated by protoc-gen-goten-go
// File: edgelq/devices/proto/v1/public_custom.proto
// DO NOT EDIT!!!

package public_client

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
	device "github.com/cloudwan/edgelq-sdk/devices/resources/v1/device"
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
	_ = &device.Device{}
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
// [ListPublicDevices][ntt.devices.v1.ListPublicDevices]
type ListPublicDevicesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Parent name of ntt.devices.v1.Device
	Parent *device.ParentName `protobuf:"bytes,1,opt,customtype=ParentName,name=parent,proto3" json:"parent,omitempty" firestore:"parent"`
	// Requested page size. Server may return fewer Devices than requested.
	// If unspecified, server will pick an appropriate default.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty" firestore:"pageSize"`
	// A token identifying a page of results the server should return.
	// Typically, this is the value of
	// [ListDevicesResponse.next_page_token][ntt.devices.v1.ListDevicesResponse.next_page_token]
	PageToken *device.PagerCursor `protobuf:"bytes,3,opt,customtype=PagerCursor,name=page_token,json=pageToken,proto3" json:"page_token,omitempty" firestore:"pageToken"`
	// Order By -
	// https://cloud.google.com/apis/design/design_patterns#list_pagination list
	// of field path with order directive, either 'asc' or 'desc'. If direction is
	// not provided, 'asc' is assumed. e.g. "state.nested_field asc,
	// state.something.else desc, theme"
	OrderBy *device.OrderBy `protobuf:"bytes,4,opt,customtype=OrderBy,name=order_by,json=orderBy,proto3" json:"order_by,omitempty" firestore:"orderBy"`
	// Filter - filter results by field criteria. Simplified SQL-like syntax with
	// following operators:
	// <=, >=, =, !=, <, >, LIKE, CONTAINS (aliases CONTAIN, HAS, HAVE), IN, IS
	// [NOT] NULL | NaN . Combine conditions with OR | AND example: 'meta.labels
	// CONTAINS "severity:important" OR (state.last_error_time >
	// "2018-11-15T10:00:00Z" AND state.status = "ERROR")'
	Filter *device.Filter `protobuf:"bytes,5,opt,customtype=Filter,name=filter,proto3" json:"filter,omitempty" firestore:"filter"`
	// A list of extra fields to be requested on top of those defined by request
	// field View.
	FieldMask *device.Device_FieldMask `protobuf:"bytes,6,opt,customtype=Device_FieldMask,name=field_mask,json=fieldMask,proto3" json:"field_mask,omitempty" firestore:"fieldMask"`
	// View defines list of standard response fields present in response items.
	// Additional fields can be amended by request field field_mask
	View view.View `protobuf:"varint,7,opt,name=view,proto3,enum=goten.types.View" json:"view,omitempty" firestore:"view"`
}

func (m *ListPublicDevicesRequest) Reset() {
	*m = ListPublicDevicesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_devices_proto_v1_public_custom_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *ListPublicDevicesRequest) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*ListPublicDevicesRequest) ProtoMessage() {}

func (m *ListPublicDevicesRequest) ProtoReflect() preflect.Message {
	mi := &edgelq_devices_proto_v1_public_custom_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*ListPublicDevicesRequest) GotenMessage() {}

// Deprecated, Use ListPublicDevicesRequest.ProtoReflect.Descriptor instead.
func (*ListPublicDevicesRequest) Descriptor() ([]byte, []int) {
	return edgelq_devices_proto_v1_public_custom_proto_rawDescGZIP(), []int{0}
}

func (m *ListPublicDevicesRequest) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *ListPublicDevicesRequest) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *ListPublicDevicesRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *ListPublicDevicesRequest) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *ListPublicDevicesRequest) GetParent() *device.ParentName {
	if m != nil {
		return m.Parent
	}
	return nil
}

func (m *ListPublicDevicesRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return int32(0)
}

func (m *ListPublicDevicesRequest) GetPageToken() *device.PagerCursor {
	if m != nil {
		return m.PageToken
	}
	return nil
}

func (m *ListPublicDevicesRequest) GetOrderBy() *device.OrderBy {
	if m != nil {
		return m.OrderBy
	}
	return nil
}

func (m *ListPublicDevicesRequest) GetFilter() *device.Filter {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *ListPublicDevicesRequest) GetFieldMask() *device.Device_FieldMask {
	if m != nil {
		return m.FieldMask
	}
	return nil
}

func (m *ListPublicDevicesRequest) GetView() view.View {
	if m != nil {
		return m.View
	}
	return view.View_UNSPECIFIED
}

func (m *ListPublicDevicesRequest) SetParent(fv *device.ParentName) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Parent", "ListPublicDevicesRequest"))
	}
	m.Parent = fv
}

func (m *ListPublicDevicesRequest) SetPageSize(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "PageSize", "ListPublicDevicesRequest"))
	}
	m.PageSize = fv
}

func (m *ListPublicDevicesRequest) SetPageToken(fv *device.PagerCursor) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "PageToken", "ListPublicDevicesRequest"))
	}
	m.PageToken = fv
}

func (m *ListPublicDevicesRequest) SetOrderBy(fv *device.OrderBy) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "OrderBy", "ListPublicDevicesRequest"))
	}
	m.OrderBy = fv
}

func (m *ListPublicDevicesRequest) SetFilter(fv *device.Filter) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Filter", "ListPublicDevicesRequest"))
	}
	m.Filter = fv
}

func (m *ListPublicDevicesRequest) SetFieldMask(fv *device.Device_FieldMask) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "FieldMask", "ListPublicDevicesRequest"))
	}
	m.FieldMask = fv
}

func (m *ListPublicDevicesRequest) SetView(fv view.View) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "View", "ListPublicDevicesRequest"))
	}
	m.View = fv
}

// Response message for method
// [ListPublicDevices][ntt.devices.v1.ListPublicDevices]
type ListPublicDevicesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// The list of Devices
	Devices []*device.Device `protobuf:"bytes,1,rep,name=devices,proto3" json:"devices,omitempty" firestore:"devices"`
	// A token to retrieve previous page of results. Pass this value in the
	// [ListDevicesRequest.page_token][ntt.devices.v1.ListDevicesRequest.page_token]
	PrevPageToken *device.PagerCursor `protobuf:"bytes,3,opt,customtype=PagerCursor,name=prev_page_token,json=prevPageToken,proto3" json:"prev_page_token,omitempty" firestore:"prevPageToken"`
	// A token to retrieve next page of results. Pass this value in the
	// [ListDevicesRequest.page_token][ntt.devices.v1.ListDevicesRequest.page_token]
	NextPageToken *device.PagerCursor `protobuf:"bytes,4,opt,customtype=PagerCursor,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty" firestore:"nextPageToken"`
}

func (m *ListPublicDevicesResponse) Reset() {
	*m = ListPublicDevicesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_devices_proto_v1_public_custom_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *ListPublicDevicesResponse) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*ListPublicDevicesResponse) ProtoMessage() {}

func (m *ListPublicDevicesResponse) ProtoReflect() preflect.Message {
	mi := &edgelq_devices_proto_v1_public_custom_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*ListPublicDevicesResponse) GotenMessage() {}

// Deprecated, Use ListPublicDevicesResponse.ProtoReflect.Descriptor instead.
func (*ListPublicDevicesResponse) Descriptor() ([]byte, []int) {
	return edgelq_devices_proto_v1_public_custom_proto_rawDescGZIP(), []int{1}
}

func (m *ListPublicDevicesResponse) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *ListPublicDevicesResponse) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *ListPublicDevicesResponse) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *ListPublicDevicesResponse) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *ListPublicDevicesResponse) GetDevices() []*device.Device {
	if m != nil {
		return m.Devices
	}
	return nil
}

func (m *ListPublicDevicesResponse) GetPrevPageToken() *device.PagerCursor {
	if m != nil {
		return m.PrevPageToken
	}
	return nil
}

func (m *ListPublicDevicesResponse) GetNextPageToken() *device.PagerCursor {
	if m != nil {
		return m.NextPageToken
	}
	return nil
}

func (m *ListPublicDevicesResponse) SetDevices(fv []*device.Device) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Devices", "ListPublicDevicesResponse"))
	}
	m.Devices = fv
}

func (m *ListPublicDevicesResponse) SetPrevPageToken(fv *device.PagerCursor) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "PrevPageToken", "ListPublicDevicesResponse"))
	}
	m.PrevPageToken = fv
}

func (m *ListPublicDevicesResponse) SetNextPageToken(fv *device.PagerCursor) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "NextPageToken", "ListPublicDevicesResponse"))
	}
	m.NextPageToken = fv
}

var edgelq_devices_proto_v1_public_custom_proto preflect.FileDescriptor

var edgelq_devices_proto_v1_public_custom_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x6e,
	0x74, 0x74, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x20, 0x67,
	0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x24, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73,
	0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x24, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xdd, 0x02, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x26, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x0e, 0xb2, 0xda, 0x21, 0x0a, 0x3a, 0x08, 0x0a, 0x06, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52,
	0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x08, 0xca, 0xc6, 0x27, 0x04,
	0x12, 0x02, 0x2a, 0x00, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x2d,
	0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x0e, 0xb2, 0xda, 0x21, 0x0a, 0x22, 0x08, 0x0a, 0x06, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x29, 0x0a,
	0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x0e, 0xb2, 0xda, 0x21, 0x0a, 0x2a, 0x08, 0x0a, 0x06, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52,
	0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x12, 0x26, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0e, 0xb2, 0xda, 0x21, 0x0a, 0x1a, 0x08,
	0x0a, 0x06, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x12, 0x49, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b,
	0x42, 0x0e, 0xb2, 0xda, 0x21, 0x0a, 0x32, 0x08, 0x0a, 0x06, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x25, 0x0a, 0x04, 0x76,
	0x69, 0x65, 0x77, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x67, 0x6f, 0x74, 0x65,
	0x6e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x52, 0x04, 0x76, 0x69,
	0x65, 0x77, 0x22, 0xbd, 0x01, 0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x30, 0x0a, 0x07, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x07, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x12, 0x36, 0x0a, 0x0f, 0x70, 0x72, 0x65, 0x76, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0e, 0xb2, 0xda, 0x21,
	0x0a, 0x22, 0x08, 0x0a, 0x06, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x0d, 0x70, 0x72, 0x65,
	0x76, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x36, 0x0a, 0x0f, 0x6e, 0x65,
	0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x0e, 0xb2, 0xda, 0x21, 0x0a, 0x22, 0x08, 0x0a, 0x06, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x42, 0x73, 0xe8, 0xde, 0x21, 0x00, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x74,
	0x74, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x42,
	0x11, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x00, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71,
	0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f,
	0x76, 0x31, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x3b, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	edgelq_devices_proto_v1_public_custom_proto_rawDescOnce sync.Once
	edgelq_devices_proto_v1_public_custom_proto_rawDescData = edgelq_devices_proto_v1_public_custom_proto_rawDesc
)

func edgelq_devices_proto_v1_public_custom_proto_rawDescGZIP() []byte {
	edgelq_devices_proto_v1_public_custom_proto_rawDescOnce.Do(func() {
		edgelq_devices_proto_v1_public_custom_proto_rawDescData = protoimpl.X.CompressGZIP(edgelq_devices_proto_v1_public_custom_proto_rawDescData)
	})
	return edgelq_devices_proto_v1_public_custom_proto_rawDescData
}

var edgelq_devices_proto_v1_public_custom_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var edgelq_devices_proto_v1_public_custom_proto_goTypes = []interface{}{
	(*ListPublicDevicesRequest)(nil),  // 0: ntt.devices.v1.ListPublicDevicesRequest
	(*ListPublicDevicesResponse)(nil), // 1: ntt.devices.v1.ListPublicDevicesResponse
	(*device.Device_FieldMask)(nil),   // 2: ntt.devices.v1.Device_FieldMask
	(view.View)(0),                    // 3: goten.types.View
	(*device.Device)(nil),             // 4: ntt.devices.v1.Device
}
var edgelq_devices_proto_v1_public_custom_proto_depIdxs = []int32{
	2, // 0: ntt.devices.v1.ListPublicDevicesRequest.field_mask:type_name -> ntt.devices.v1.Device_FieldMask
	3, // 1: ntt.devices.v1.ListPublicDevicesRequest.view:type_name -> goten.types.View
	4, // 2: ntt.devices.v1.ListPublicDevicesResponse.devices:type_name -> ntt.devices.v1.Device
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { edgelq_devices_proto_v1_public_custom_proto_init() }
func edgelq_devices_proto_v1_public_custom_proto_init() {
	if edgelq_devices_proto_v1_public_custom_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		edgelq_devices_proto_v1_public_custom_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPublicDevicesRequest); i {
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
		edgelq_devices_proto_v1_public_custom_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPublicDevicesResponse); i {
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
			RawDescriptor: edgelq_devices_proto_v1_public_custom_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           edgelq_devices_proto_v1_public_custom_proto_goTypes,
		DependencyIndexes: edgelq_devices_proto_v1_public_custom_proto_depIdxs,
		MessageInfos:      edgelq_devices_proto_v1_public_custom_proto_msgTypes,
	}.Build()
	edgelq_devices_proto_v1_public_custom_proto = out.File
	edgelq_devices_proto_v1_public_custom_proto_rawDesc = nil
	edgelq_devices_proto_v1_public_custom_proto_goTypes = nil
	edgelq_devices_proto_v1_public_custom_proto_depIdxs = nil
}
