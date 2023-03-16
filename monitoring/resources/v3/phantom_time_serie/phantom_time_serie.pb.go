// Code generated by protoc-gen-goten-go
// File: edgelq/monitoring/proto/v3/phantom_time_serie.proto
// DO NOT EDIT!!!

package phantom_time_serie

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
	monitoring_common "github.com/cloudwan/edgelq-sdk/monitoring/common/v3"
	metric_descriptor "github.com/cloudwan/edgelq-sdk/monitoring/resources/v3/metric_descriptor"
	project "github.com/cloudwan/edgelq-sdk/monitoring/resources/v3/project"
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
	_ = &monitoring_common.LabelDescriptor{}
	_ = &metric_descriptor.MetricDescriptor{}
	_ = &project.Project{}
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// PhantomTimeSerie generates data in absence of real data
type PhantomTimeSerie struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Common resource Metadata
	Metadata *ntt_meta.Meta `protobuf:"bytes,11,opt,name=metadata,proto3" json:"metadata,omitempty" firestore:"metadata"`
	// Name of PhantomTimeSeries
	Name *Name `protobuf:"bytes,100,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// TimeSerie key identifies unique TimeSeries tuple:
	// <project, metric.type, metric.labels, resource.type, resource.labels>
	//
	// Kind/ValueType are not present in key
	// Key is not to be decoded outside of service, but treated as opaque string
	Key []byte `protobuf:"bytes,101,opt,name=key,proto3" json:"key,omitempty" firestore:"key"`
	// Internal use - for bulk reporting of TimeSeries
	Project string `protobuf:"bytes,102,opt,name=project,proto3" json:"project,omitempty" firestore:"project"`
	// The associated metric. A fully-specified metric used to identify the time
	// series.
	// This field cannot be updated, can be only set during creation.
	Metric *monitoring_common.Metric `protobuf:"bytes,1,opt,name=metric,proto3" json:"metric,omitempty" firestore:"metric"`
	// The associated monitored resource.  Custom metrics can use only certain
	// monitored resource types in their time series data.
	// This field cannot be updated, can be only set during creation.
	Resource *monitoring_common.MonitoredResource `protobuf:"bytes,2,opt,name=resource,proto3" json:"resource,omitempty" firestore:"resource"`
	// The metric kind of the time series. When listing time series, this metric
	// kind might be different from the metric kind of the associated metric if
	// this time series is an alignment or reduction of other time series.
	//
	// When creating a time series, this field is optional. If present, it must be
	// the same as the metric kind of the associated metric. If the associated
	// metric's descriptor must be auto-created, then this field specifies the
	// metric kind of the new descriptor and must be either `GAUGE` (the default)
	// or `CUMULATIVE`.
	MetricKind metric_descriptor.MetricDescriptor_MetricKind `protobuf:"varint,3,opt,name=metric_kind,json=metricKind,proto3,enum=ntt.monitoring.v3.MetricDescriptor_MetricKind" json:"metric_kind,omitempty" firestore:"metricKind"`
	// The value type of the time series. When listing time series, this value
	// type might be different from the value type of the associated metric if
	// this time series is an alignment or reduction of other time series.
	//
	// When creating a time series, this field is optional. If present, it must be
	// the same as the type of the data in the `points` field.
	ValueType metric_descriptor.MetricDescriptor_ValueType `protobuf:"varint,4,opt,name=value_type,json=valueType,proto3,enum=ntt.monitoring.v3.MetricDescriptor_ValueType" json:"value_type,omitempty" firestore:"valueType"`
	// Phantom value
	Value *monitoring_common.TypedValue `protobuf:"bytes,6,opt,name=value,proto3" json:"value,omitempty" firestore:"value"`
}

func (m *PhantomTimeSerie) Reset() {
	*m = PhantomTimeSerie{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_monitoring_proto_v3_phantom_time_serie_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *PhantomTimeSerie) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*PhantomTimeSerie) ProtoMessage() {}

func (m *PhantomTimeSerie) ProtoReflect() preflect.Message {
	mi := &edgelq_monitoring_proto_v3_phantom_time_serie_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*PhantomTimeSerie) GotenMessage() {}

// Deprecated, Use PhantomTimeSerie.ProtoReflect.Descriptor instead.
func (*PhantomTimeSerie) Descriptor() ([]byte, []int) {
	return edgelq_monitoring_proto_v3_phantom_time_serie_proto_rawDescGZIP(), []int{0}
}

func (m *PhantomTimeSerie) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *PhantomTimeSerie) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *PhantomTimeSerie) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *PhantomTimeSerie) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *PhantomTimeSerie) GetMetadata() *ntt_meta.Meta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *PhantomTimeSerie) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *PhantomTimeSerie) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *PhantomTimeSerie) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *PhantomTimeSerie) GetMetric() *monitoring_common.Metric {
	if m != nil {
		return m.Metric
	}
	return nil
}

func (m *PhantomTimeSerie) GetResource() *monitoring_common.MonitoredResource {
	if m != nil {
		return m.Resource
	}
	return nil
}

func (m *PhantomTimeSerie) GetMetricKind() metric_descriptor.MetricDescriptor_MetricKind {
	if m != nil {
		return m.MetricKind
	}
	return metric_descriptor.MetricDescriptor_METRIC_KIND_UNSPECIFIED
}

func (m *PhantomTimeSerie) GetValueType() metric_descriptor.MetricDescriptor_ValueType {
	if m != nil {
		return m.ValueType
	}
	return metric_descriptor.MetricDescriptor_VALUE_TYPE_UNSPECIFIED
}

func (m *PhantomTimeSerie) GetValue() *monitoring_common.TypedValue {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *PhantomTimeSerie) SetMetadata(fv *ntt_meta.Meta) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Metadata", "PhantomTimeSerie"))
	}
	m.Metadata = fv
}

func (m *PhantomTimeSerie) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "PhantomTimeSerie"))
	}
	m.Name = fv
}

func (m *PhantomTimeSerie) SetKey(fv []byte) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Key", "PhantomTimeSerie"))
	}
	m.Key = fv
}

func (m *PhantomTimeSerie) SetProject(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Project", "PhantomTimeSerie"))
	}
	m.Project = fv
}

func (m *PhantomTimeSerie) SetMetric(fv *monitoring_common.Metric) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Metric", "PhantomTimeSerie"))
	}
	m.Metric = fv
}

func (m *PhantomTimeSerie) SetResource(fv *monitoring_common.MonitoredResource) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Resource", "PhantomTimeSerie"))
	}
	m.Resource = fv
}

func (m *PhantomTimeSerie) SetMetricKind(fv metric_descriptor.MetricDescriptor_MetricKind) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "MetricKind", "PhantomTimeSerie"))
	}
	m.MetricKind = fv
}

func (m *PhantomTimeSerie) SetValueType(fv metric_descriptor.MetricDescriptor_ValueType) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ValueType", "PhantomTimeSerie"))
	}
	m.ValueType = fv
}

func (m *PhantomTimeSerie) SetValue(fv *monitoring_common.TypedValue) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Value", "PhantomTimeSerie"))
	}
	m.Value = fv
}

var edgelq_monitoring_proto_v3_phantom_time_serie_proto preflect.FileDescriptor

var edgelq_monitoring_proto_v3_phantom_time_serie_proto_rawDesc = []byte{
	0x0a, 0x33, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x69, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x33, 0x2f, 0x70, 0x68, 0x61,
	0x6e, 0x74, 0x6f, 0x6d, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x69, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x6e, 0x74, 0x74, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74,
	0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76,
	0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x74, 0x65, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f,
	0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x34, 0x65,
	0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x69,
	0x6c, 0x65, 0x72, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6d, 0x6f, 0x6e, 0x69,
	0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x33, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x32, 0x65, 0x64,
	0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x33, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xbe, 0x07, 0x0a, 0x10, 0x50, 0x68, 0x61, 0x6e, 0x74, 0x6f, 0x6d, 0x54, 0x69, 0x6d, 0x65,
	0x53, 0x65, 0x72, 0x69, 0x65, 0x12, 0x2b, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x2c, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x18, 0xb2, 0xda, 0x21, 0x14, 0x0a, 0x12, 0x0a, 0x10, 0x50, 0x68, 0x61, 0x6e, 0x74, 0x6f,
	0x6d, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x19, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x07, 0xe0,
	0x41, 0x03, 0xf0, 0xd9, 0x21, 0x01, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x21, 0x0a, 0x07, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x66, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xe0, 0x41,
	0x03, 0xf0, 0xd9, 0x21, 0x01, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x31,
	0x0a, 0x06, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x33, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x12, 0x40, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f,
	0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x65,
	0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x12, 0x58, 0x0a, 0x0b, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x6b, 0x69,
	0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2e, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x6d,
	0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x4d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x4b, 0x69, 0x6e, 0x64, 0x42, 0x07, 0xe0, 0x41, 0x03, 0xf0, 0xd9, 0x21,
	0x01, 0x52, 0x0a, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x55, 0x0a,
	0x0a, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x2d, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69,
	0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x42, 0x07, 0xe0, 0x41, 0x03, 0xf0, 0xd9, 0x21, 0x01, 0x52, 0x09, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x33, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f,
	0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x64, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0xb5, 0x03, 0xea, 0x41, 0x74, 0x0a,
	0x26, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x65, 0x64, 0x67, 0x65,
	0x6c, 0x71, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x68, 0x61, 0x6e, 0x74, 0x6f, 0x6d, 0x54, 0x69,
	0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x12, 0x4a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x72, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x70, 0x68, 0x61,
	0x6e, 0x74, 0x6f, 0x6d, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x2f, 0x7b,
	0x70, 0x68, 0x61, 0x6e, 0x74, 0x6f, 0x6d, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72,
	0x69, 0x65, 0x7d, 0x92, 0xd9, 0x21, 0xd3, 0x01, 0x0a, 0x11, 0x70, 0x68, 0x61, 0x6e, 0x74, 0x6f,
	0x6d, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x12, 0x11, 0x70, 0x68, 0x61,
	0x6e, 0x74, 0x6f, 0x6d, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x1a, 0x07,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x3a, 0x0e, 0x5b,
	0x5c, 0x77, 0x2b, 0x2f, 0x3d, 0x5d, 0x7b, 0x31, 0x2c, 0x32, 0x35, 0x36, 0x7d, 0x4a, 0x6c, 0x08,
	0x02, 0x12, 0x06, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x0b, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x20, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x12, 0x0d, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x20, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0d, 0x0a, 0x0b, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f,
	0x6b, 0x69, 0x6e, 0x64, 0x12, 0x0c, 0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x07, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x50, 0x05, 0x6a, 0x1c, 0x6e,
	0x74, 0x74, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f, 0x72, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x2f, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0xda, 0x94, 0x23, 0x08, 0x12,
	0x06, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0xc2, 0x85, 0x2c, 0x50, 0x22, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x22,
	0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x06, 0x6d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x22, 0x0b, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x6b, 0x69, 0x6e, 0x64, 0x22, 0x0a,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x42, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xe2, 0xde, 0x21, 0x02, 0x08,
	0x02, 0x42, 0xa3, 0x03, 0xe8, 0xde, 0x21, 0x01, 0xd2, 0xff, 0xd0, 0x02, 0x5d, 0x0a, 0x18, 0x70,
	0x68, 0x61, 0x6e, 0x74, 0x6f, 0x6d, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x69,
	0x65, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67,
	0x65, 0x6c, 0x71, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x70, 0x68, 0x61, 0x6e, 0x74, 0x6f, 0x6d, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x69, 0x65, 0x0a, 0x18, 0x63, 0x6f, 0x6d, 0x2e,
	0x6e, 0x74, 0x74, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x70,
	0x62, 0x2e, 0x76, 0x33, 0x42, 0x15, 0x50, 0x68, 0x61, 0x6e, 0x74, 0x6f, 0x6d, 0x54, 0x69, 0x6d,
	0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x58, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77,
	0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f,
	0x72, 0x69, 0x6e, 0x67, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76,
	0x33, 0x2f, 0x70, 0x68, 0x61, 0x6e, 0x74, 0x6f, 0x6d, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73,
	0x65, 0x72, 0x69, 0x65, 0x3b, 0x70, 0x68, 0x61, 0x6e, 0x74, 0x6f, 0x6d, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x5f, 0x73, 0x65, 0x72, 0x69, 0x65, 0xd2, 0x84, 0xd1, 0x02, 0x47, 0x0a, 0x0d, 0x61, 0x75,
	0x64, 0x69, 0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x12, 0x36, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e,
	0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69,
	0x6e, 0x67, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73,
	0x2f, 0x76, 0x33, 0xa2, 0x80, 0xd1, 0x02, 0x5f, 0x0a, 0x19, 0x70, 0x68, 0x61, 0x6e, 0x74, 0x6f,
	0x6d, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x69, 0x65, 0x5f, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x12, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f,
	0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x2f, 0x76, 0x33, 0x2f, 0x70, 0x68, 0x61, 0x6e, 0x74, 0x6f, 0x6d, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x5f, 0x73, 0x65, 0x72, 0x69, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	edgelq_monitoring_proto_v3_phantom_time_serie_proto_rawDescOnce sync.Once
	edgelq_monitoring_proto_v3_phantom_time_serie_proto_rawDescData = edgelq_monitoring_proto_v3_phantom_time_serie_proto_rawDesc
)

func edgelq_monitoring_proto_v3_phantom_time_serie_proto_rawDescGZIP() []byte {
	edgelq_monitoring_proto_v3_phantom_time_serie_proto_rawDescOnce.Do(func() {
		edgelq_monitoring_proto_v3_phantom_time_serie_proto_rawDescData = protoimpl.X.CompressGZIP(edgelq_monitoring_proto_v3_phantom_time_serie_proto_rawDescData)
	})
	return edgelq_monitoring_proto_v3_phantom_time_serie_proto_rawDescData
}

var edgelq_monitoring_proto_v3_phantom_time_serie_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var edgelq_monitoring_proto_v3_phantom_time_serie_proto_goTypes = []interface{}{
	(*PhantomTimeSerie)(nil),                           // 0: ntt.monitoring.v3.PhantomTimeSerie
	(*ntt_meta.Meta)(nil),                              // 1: ntt.types.Meta
	(*monitoring_common.Metric)(nil),                   // 2: ntt.monitoring.v3.Metric
	(*monitoring_common.MonitoredResource)(nil),        // 3: ntt.monitoring.v3.MonitoredResource
	(metric_descriptor.MetricDescriptor_MetricKind)(0), // 4: ntt.monitoring.v3.MetricDescriptor_MetricKind
	(metric_descriptor.MetricDescriptor_ValueType)(0),  // 5: ntt.monitoring.v3.MetricDescriptor_ValueType
	(*monitoring_common.TypedValue)(nil),               // 6: ntt.monitoring.v3.TypedValue
}
var edgelq_monitoring_proto_v3_phantom_time_serie_proto_depIdxs = []int32{
	1, // 0: ntt.monitoring.v3.PhantomTimeSerie.metadata:type_name -> ntt.types.Meta
	2, // 1: ntt.monitoring.v3.PhantomTimeSerie.metric:type_name -> ntt.monitoring.v3.Metric
	3, // 2: ntt.monitoring.v3.PhantomTimeSerie.resource:type_name -> ntt.monitoring.v3.MonitoredResource
	4, // 3: ntt.monitoring.v3.PhantomTimeSerie.metric_kind:type_name -> ntt.monitoring.v3.MetricDescriptor_MetricKind
	5, // 4: ntt.monitoring.v3.PhantomTimeSerie.value_type:type_name -> ntt.monitoring.v3.MetricDescriptor_ValueType
	6, // 5: ntt.monitoring.v3.PhantomTimeSerie.value:type_name -> ntt.monitoring.v3.TypedValue
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { edgelq_monitoring_proto_v3_phantom_time_serie_proto_init() }
func edgelq_monitoring_proto_v3_phantom_time_serie_proto_init() {
	if edgelq_monitoring_proto_v3_phantom_time_serie_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		edgelq_monitoring_proto_v3_phantom_time_serie_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PhantomTimeSerie); i {
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
			RawDescriptor: edgelq_monitoring_proto_v3_phantom_time_serie_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           edgelq_monitoring_proto_v3_phantom_time_serie_proto_goTypes,
		DependencyIndexes: edgelq_monitoring_proto_v3_phantom_time_serie_proto_depIdxs,
		MessageInfos:      edgelq_monitoring_proto_v3_phantom_time_serie_proto_msgTypes,
	}.Build()
	edgelq_monitoring_proto_v3_phantom_time_serie_proto = out.File
	edgelq_monitoring_proto_v3_phantom_time_serie_proto_rawDesc = nil
	edgelq_monitoring_proto_v3_phantom_time_serie_proto_goTypes = nil
	edgelq_monitoring_proto_v3_phantom_time_serie_proto_depIdxs = nil
}
