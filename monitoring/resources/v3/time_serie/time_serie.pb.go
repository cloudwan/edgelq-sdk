// Code generated by protoc-gen-goten-go
// File: edgelq/monitoring/proto/v3/time_serie.proto
// DO NOT EDIT!!!

package time_serie

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
	monitoring_common "github.com/cloudwan/edgelq-sdk/monitoring/common/v3"
	metric_descriptor "github.com/cloudwan/edgelq-sdk/monitoring/resources/v3/metric_descriptor"
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
	_ = &monitoring_common.LabelDescriptor{}
	_ = &metric_descriptor.MetricDescriptor{}
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// A single data point in a time series.
type Point struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// The time interval to which the data point applies.  For `GAUGE` metrics,
	// only the end time of the interval is used.  For `DELTA` metrics, the start
	// and end time should specify a non-zero interval, with subsequent points
	// specifying contiguous and non-overlapping intervals.  For `CUMULATIVE`
	// metrics, the start and end time should specify a non-zero interval, with
	// subsequent points specifying the same start time and increasing end times,
	// until an event resets the cumulative value to zero and sets a new start
	// time for the following points.
	Interval *monitoring_common.TimeInterval `protobuf:"bytes,1,opt,name=interval,proto3" json:"interval,omitempty" firestore:"interval"`
	// The value of the data point.
	Value *monitoring_common.TypedValue `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty" firestore:"value"`
	// Additional aggregation info
	// Used internally for batching rollup points
	Aggregation *monitoring_common.Aggregation `protobuf:"bytes,101,opt,name=aggregation,proto3" json:"aggregation,omitempty" firestore:"aggregation"`
}

func (m *Point) Reset() {
	*m = Point{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_monitoring_proto_v3_time_serie_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *Point) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*Point) ProtoMessage() {}

func (m *Point) ProtoReflect() preflect.Message {
	mi := &edgelq_monitoring_proto_v3_time_serie_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*Point) GotenMessage() {}

// Deprecated, Use Point.ProtoReflect.Descriptor instead.
func (*Point) Descriptor() ([]byte, []int) {
	return edgelq_monitoring_proto_v3_time_serie_proto_rawDescGZIP(), []int{0}
}

func (m *Point) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *Point) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *Point) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *Point) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *Point) GetInterval() *monitoring_common.TimeInterval {
	if m != nil {
		return m.Interval
	}
	return nil
}

func (m *Point) GetValue() *monitoring_common.TypedValue {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Point) GetAggregation() *monitoring_common.Aggregation {
	if m != nil {
		return m.Aggregation
	}
	return nil
}

func (m *Point) SetInterval(fv *monitoring_common.TimeInterval) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Interval", "Point"))
	}
	m.Interval = fv
}

func (m *Point) SetValue(fv *monitoring_common.TypedValue) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Value", "Point"))
	}
	m.Value = fv
}

func (m *Point) SetAggregation(fv *monitoring_common.Aggregation) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Aggregation", "Point"))
	}
	m.Aggregation = fv
}

// A collection of data points that describes the time-varying values
// of a metric. A time series is identified by a combination of a
// fully-specified monitored resource and a fully-specified metric.
// This type is used for both listing and creating time series.
type TimeSerie struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// TimeSerie key identifies unique TimeSeries tuple:
	// <project, region, metric.type, metric.labels, resource.type,
	// resource.labels>
	//
	// Kind/ValueType are not present in key
	// Key is not to be decoded outside of service, but treated as opaque string
	//
	// Specific key is valid and understood only in single region only. If time
	// serie is created by merging from multiple regions, key must be ignore.
	Key []byte `protobuf:"bytes,101,opt,name=key,proto3" json:"key,omitempty" firestore:"key"`
	// Internal use - for bulk reporting of TimeSeries
	Project string `protobuf:"bytes,102,opt,name=project,proto3" json:"project,omitempty" firestore:"project"`
	// Region ID associated with time serie.
	Region string `protobuf:"bytes,103,opt,name=region,proto3" json:"region,omitempty" firestore:"region"`
	// The associated metric. A fully-specified metric used to identify the time
	// series.
	Metric *monitoring_common.Metric `protobuf:"bytes,1,opt,name=metric,proto3" json:"metric,omitempty" firestore:"metric"`
	// The associated monitored resource.  Custom metrics can use only certain
	// monitored resource types in their time series data.
	Resource *monitoring_common.MonitoredResource `protobuf:"bytes,2,opt,name=resource,proto3" json:"resource,omitempty" firestore:"resource"`
	// Output only. The associated monitored resource metadata. When reading a
	// a timeseries, this field will include metadata labels that are explicitly
	// named in the reduction. When creating a timeseries, this field is ignored.
	Metadata *monitoring_common.MonitoredResourceMetadata `protobuf:"bytes,7,opt,name=metadata,proto3" json:"metadata,omitempty" firestore:"metadata"`
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
	// The data points of this time series. When listing time series, points are
	// returned in reverse time order.
	//
	// When creating a time series, this field must contain exactly one point and
	// the point's type must be the same as the value type of the associated
	// metric. If the associated metric's descriptor must be auto-created, then
	// the value type of the descriptor is determined by the point's type, which
	// must be `BOOL`, `INT64`, `DOUBLE`, or `DISTRIBUTION`.
	Points []*Point `protobuf:"bytes,5,rep,name=points,proto3" json:"points,omitempty" firestore:"points"`
}

func (m *TimeSerie) Reset() {
	*m = TimeSerie{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_monitoring_proto_v3_time_serie_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *TimeSerie) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*TimeSerie) ProtoMessage() {}

func (m *TimeSerie) ProtoReflect() preflect.Message {
	mi := &edgelq_monitoring_proto_v3_time_serie_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*TimeSerie) GotenMessage() {}

// Deprecated, Use TimeSerie.ProtoReflect.Descriptor instead.
func (*TimeSerie) Descriptor() ([]byte, []int) {
	return edgelq_monitoring_proto_v3_time_serie_proto_rawDescGZIP(), []int{1}
}

func (m *TimeSerie) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *TimeSerie) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *TimeSerie) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *TimeSerie) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *TimeSerie) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *TimeSerie) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *TimeSerie) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *TimeSerie) GetMetric() *monitoring_common.Metric {
	if m != nil {
		return m.Metric
	}
	return nil
}

func (m *TimeSerie) GetResource() *monitoring_common.MonitoredResource {
	if m != nil {
		return m.Resource
	}
	return nil
}

func (m *TimeSerie) GetMetadata() *monitoring_common.MonitoredResourceMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *TimeSerie) GetMetricKind() metric_descriptor.MetricDescriptor_MetricKind {
	if m != nil {
		return m.MetricKind
	}
	return metric_descriptor.MetricDescriptor_METRIC_KIND_UNSPECIFIED
}

func (m *TimeSerie) GetValueType() metric_descriptor.MetricDescriptor_ValueType {
	if m != nil {
		return m.ValueType
	}
	return metric_descriptor.MetricDescriptor_VALUE_TYPE_UNSPECIFIED
}

func (m *TimeSerie) GetPoints() []*Point {
	if m != nil {
		return m.Points
	}
	return nil
}

func (m *TimeSerie) SetKey(fv []byte) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Key", "TimeSerie"))
	}
	m.Key = fv
}

func (m *TimeSerie) SetProject(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Project", "TimeSerie"))
	}
	m.Project = fv
}

func (m *TimeSerie) SetRegion(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Region", "TimeSerie"))
	}
	m.Region = fv
}

func (m *TimeSerie) SetMetric(fv *monitoring_common.Metric) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Metric", "TimeSerie"))
	}
	m.Metric = fv
}

func (m *TimeSerie) SetResource(fv *monitoring_common.MonitoredResource) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Resource", "TimeSerie"))
	}
	m.Resource = fv
}

func (m *TimeSerie) SetMetadata(fv *monitoring_common.MonitoredResourceMetadata) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Metadata", "TimeSerie"))
	}
	m.Metadata = fv
}

func (m *TimeSerie) SetMetricKind(fv metric_descriptor.MetricDescriptor_MetricKind) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "MetricKind", "TimeSerie"))
	}
	m.MetricKind = fv
}

func (m *TimeSerie) SetValueType(fv metric_descriptor.MetricDescriptor_ValueType) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ValueType", "TimeSerie"))
	}
	m.ValueType = fv
}

func (m *TimeSerie) SetPoints(fv []*Point) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Points", "TimeSerie"))
	}
	m.Points = fv
}

// Used for reporting rollups
type BulkTimeSeries struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	TimeSeries    []*TimeSerie `protobuf:"bytes,1,rep,name=time_series,json=timeSeries,proto3" json:"time_series,omitempty" firestore:"timeSeries"`
}

func (m *BulkTimeSeries) Reset() {
	*m = BulkTimeSeries{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_monitoring_proto_v3_time_serie_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *BulkTimeSeries) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*BulkTimeSeries) ProtoMessage() {}

func (m *BulkTimeSeries) ProtoReflect() preflect.Message {
	mi := &edgelq_monitoring_proto_v3_time_serie_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*BulkTimeSeries) GotenMessage() {}

// Deprecated, Use BulkTimeSeries.ProtoReflect.Descriptor instead.
func (*BulkTimeSeries) Descriptor() ([]byte, []int) {
	return edgelq_monitoring_proto_v3_time_serie_proto_rawDescGZIP(), []int{2}
}

func (m *BulkTimeSeries) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *BulkTimeSeries) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *BulkTimeSeries) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *BulkTimeSeries) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *BulkTimeSeries) GetTimeSeries() []*TimeSerie {
	if m != nil {
		return m.TimeSeries
	}
	return nil
}

func (m *BulkTimeSeries) SetTimeSeries(fv []*TimeSerie) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "TimeSeries", "BulkTimeSeries"))
	}
	m.TimeSeries = fv
}

var edgelq_monitoring_proto_v3_time_serie_proto preflect.FileDescriptor

var edgelq_monitoring_proto_v3_time_serie_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x69, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x33, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x5f, 0x73, 0x65, 0x72, 0x69, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x6e,
	0x74, 0x74, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33,
	0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x74,
	0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67,
	0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f,
	0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d,
	0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67,
	0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x65,
	0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x34, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6d, 0x65, 0x74, 0x61,
	0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x72, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x65, 0x64, 0x67, 0x65, 0x6c,
	0x71, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x76, 0x33, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x32, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74,
	0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x33, 0x2f, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbb, 0x01, 0x0a, 0x05, 0x50, 0x6f, 0x69, 0x6e, 0x74,
	0x12, 0x3b, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x76, 0x61, 0x6c, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x33, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6e,
	0x74, 0x74, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33,
	0x2e, 0x54, 0x79, 0x70, 0x65, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x6d, 0x6f,
	0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x41, 0x67, 0x67, 0x72,
	0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0xde, 0x04, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72,
	0x69, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18,
	0x66, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x67, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x31, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x6d, 0x6f, 0x6e,
	0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x12, 0x40, 0x0a, 0x08, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6e, 0x74,
	0x74, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e,
	0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x48, 0x0a, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e,
	0x6e, 0x74, 0x74, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76,
	0x33, 0x2e, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x4f, 0x0a, 0x0b, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f,
	0x6b, 0x69, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2e, 0x2e, 0x6e, 0x74, 0x74,
	0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x4d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x4b, 0x69, 0x6e, 0x64, 0x52, 0x0a, 0x6d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x4c, 0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2d, 0x2e, 0x6e, 0x74, 0x74,
	0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x4d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x30, 0x0a, 0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74,
	0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x06,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x3a, 0x7d, 0xea, 0x41, 0x4d, 0x0a, 0x1f, 0x6d, 0x6f, 0x6e,
	0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x12, 0x2a, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x2f, 0x7b, 0x74, 0x69, 0x6d,
	0x65, 0x5f, 0x73, 0x65, 0x72, 0x69, 0x65, 0x7d, 0x92, 0xd9, 0x21, 0x29, 0x0a, 0x0a, 0x74, 0x69,
	0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x12, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x65,
	0x72, 0x69, 0x65, 0x73, 0x1a, 0x07, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x28, 0x01, 0x40,
	0x01, 0x5a, 0x02, 0x08, 0x01, 0x22, 0x4f, 0x0a, 0x0e, 0x42, 0x75, 0x6c, 0x6b, 0x54, 0x69, 0x6d,
	0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x12, 0x3d, 0x0a, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x5f,
	0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6e,
	0x74, 0x74, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x52, 0x0a, 0x74, 0x69, 0x6d, 0x65,
	0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x42, 0xc6, 0x03, 0xe8, 0xde, 0x21, 0x01, 0xd2, 0xff, 0xd0,
	0x02, 0x4d, 0x0a, 0x10, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x69, 0x65, 0x5f, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x12, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71,
	0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2f, 0x76, 0x33, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x69, 0x65, 0x0a,
	0x18, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x69, 0x6e, 0x67, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x33, 0x42, 0x0e, 0x54, 0x69, 0x6d, 0x65, 0x53,
	0x65, 0x72, 0x69, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x48, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e,
	0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69,
	0x6e, 0x67, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x33, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x69, 0x65, 0x3b, 0x74, 0x69, 0x6d, 0x65, 0x5f,
	0x73, 0x65, 0x72, 0x69, 0x65, 0xd2, 0x84, 0xd1, 0x02, 0x47, 0x0a, 0x0d, 0x61, 0x75, 0x64, 0x69,
	0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x12, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65,
	0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67,
	0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x2f, 0x76,
	0x33, 0xa2, 0x80, 0xd1, 0x02, 0x4f, 0x0a, 0x11, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72,
	0x69, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65,
	0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67,
	0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x5f,
	0x73, 0x65, 0x72, 0x69, 0x65, 0xf2, 0x85, 0xd1, 0x02, 0x55, 0x0a, 0x14, 0x74, 0x69, 0x6d, 0x65,
	0x5f, 0x73, 0x65, 0x72, 0x69, 0x65, 0x5f, 0x64, 0x62, 0x5f, 0x73, 0x79, 0x6e, 0x63, 0x65, 0x72,
	0x12, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6d, 0x6f, 0x6e,
	0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x64, 0x62, 0x5f, 0x73, 0x79, 0x6e, 0x63, 0x65,
	0x72, 0x2f, 0x76, 0x33, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x69, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	edgelq_monitoring_proto_v3_time_serie_proto_rawDescOnce sync.Once
	edgelq_monitoring_proto_v3_time_serie_proto_rawDescData = edgelq_monitoring_proto_v3_time_serie_proto_rawDesc
)

func edgelq_monitoring_proto_v3_time_serie_proto_rawDescGZIP() []byte {
	edgelq_monitoring_proto_v3_time_serie_proto_rawDescOnce.Do(func() {
		edgelq_monitoring_proto_v3_time_serie_proto_rawDescData = protoimpl.X.CompressGZIP(edgelq_monitoring_proto_v3_time_serie_proto_rawDescData)
	})
	return edgelq_monitoring_proto_v3_time_serie_proto_rawDescData
}

var edgelq_monitoring_proto_v3_time_serie_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var edgelq_monitoring_proto_v3_time_serie_proto_goTypes = []interface{}{
	(*Point)(nil),                                       // 0: ntt.monitoring.v3.Point
	(*TimeSerie)(nil),                                   // 1: ntt.monitoring.v3.TimeSerie
	(*BulkTimeSeries)(nil),                              // 2: ntt.monitoring.v3.BulkTimeSeries
	(*monitoring_common.TimeInterval)(nil),              // 3: ntt.monitoring.v3.TimeInterval
	(*monitoring_common.TypedValue)(nil),                // 4: ntt.monitoring.v3.TypedValue
	(*monitoring_common.Aggregation)(nil),               // 5: ntt.monitoring.v3.Aggregation
	(*monitoring_common.Metric)(nil),                    // 6: ntt.monitoring.v3.Metric
	(*monitoring_common.MonitoredResource)(nil),         // 7: ntt.monitoring.v3.MonitoredResource
	(*monitoring_common.MonitoredResourceMetadata)(nil), // 8: ntt.monitoring.v3.MonitoredResourceMetadata
	(metric_descriptor.MetricDescriptor_MetricKind)(0),  // 9: ntt.monitoring.v3.MetricDescriptor_MetricKind
	(metric_descriptor.MetricDescriptor_ValueType)(0),   // 10: ntt.monitoring.v3.MetricDescriptor_ValueType
}
var edgelq_monitoring_proto_v3_time_serie_proto_depIdxs = []int32{
	3,  // 0: ntt.monitoring.v3.Point.interval:type_name -> ntt.monitoring.v3.TimeInterval
	4,  // 1: ntt.monitoring.v3.Point.value:type_name -> ntt.monitoring.v3.TypedValue
	5,  // 2: ntt.monitoring.v3.Point.aggregation:type_name -> ntt.monitoring.v3.Aggregation
	6,  // 3: ntt.monitoring.v3.TimeSerie.metric:type_name -> ntt.monitoring.v3.Metric
	7,  // 4: ntt.monitoring.v3.TimeSerie.resource:type_name -> ntt.monitoring.v3.MonitoredResource
	8,  // 5: ntt.monitoring.v3.TimeSerie.metadata:type_name -> ntt.monitoring.v3.MonitoredResourceMetadata
	9,  // 6: ntt.monitoring.v3.TimeSerie.metric_kind:type_name -> ntt.monitoring.v3.MetricDescriptor_MetricKind
	10, // 7: ntt.monitoring.v3.TimeSerie.value_type:type_name -> ntt.monitoring.v3.MetricDescriptor_ValueType
	0,  // 8: ntt.monitoring.v3.TimeSerie.points:type_name -> ntt.monitoring.v3.Point
	1,  // 9: ntt.monitoring.v3.BulkTimeSeries.time_series:type_name -> ntt.monitoring.v3.TimeSerie
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { edgelq_monitoring_proto_v3_time_serie_proto_init() }
func edgelq_monitoring_proto_v3_time_serie_proto_init() {
	if edgelq_monitoring_proto_v3_time_serie_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		edgelq_monitoring_proto_v3_time_serie_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Point); i {
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
		edgelq_monitoring_proto_v3_time_serie_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimeSerie); i {
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
		edgelq_monitoring_proto_v3_time_serie_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BulkTimeSeries); i {
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
			RawDescriptor: edgelq_monitoring_proto_v3_time_serie_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           edgelq_monitoring_proto_v3_time_serie_proto_goTypes,
		DependencyIndexes: edgelq_monitoring_proto_v3_time_serie_proto_depIdxs,
		MessageInfos:      edgelq_monitoring_proto_v3_time_serie_proto_msgTypes,
	}.Build()
	edgelq_monitoring_proto_v3_time_serie_proto = out.File
	edgelq_monitoring_proto_v3_time_serie_proto_rawDesc = nil
	edgelq_monitoring_proto_v3_time_serie_proto_goTypes = nil
	edgelq_monitoring_proto_v3_time_serie_proto_depIdxs = nil
}
