// Code generated by protoc-gen-goten-go
// File: edgelq/monitoring/proto/v4/time_series_collection_rule.proto
// DO NOT EDIT!!!

package time_series_collection_rule

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
	common "github.com/cloudwan/edgelq-sdk/monitoring/resources/v4/common"
	project "github.com/cloudwan/edgelq-sdk/monitoring/resources/v4/project"
	time_serie "github.com/cloudwan/edgelq-sdk/monitoring/resources/v4/time_serie"
	time_series_forwarder_sink "github.com/cloudwan/edgelq-sdk/monitoring/resources/v4/time_series_forwarder_sink"
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
	_ = &common.LabelDescriptor{}
	_ = &project.Project{}
	_ = &time_serie.Point{}
	_ = &time_series_forwarder_sink.TimeSeriesForwarderSink{}
	_ = &meta.Meta{}
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// TimeSeriesCollectionRule Resource is a persistent WatchTimeSeries
// session registered on the server side. It collects time series according
// to the specified filter/aggregation, and within a project where rule is.
// Sink resource can be from different project.
type TimeSeriesCollectionRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Name of TimeSeriesCollectionRule
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
	// Optional display name
	DisplayName string `protobuf:"bytes,7,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Time series filter to apply.
	Filter *time_serie.Filter `protobuf:"bytes,3,opt,customtype=Filter,name=filter,proto3" json:"filter,omitempty"`
	// Instructs how to transform individual time series (aligner) and combine
	// them together (reducer, group by fields).
	Aggregation *common.Aggregation `protobuf:"bytes,4,opt,name=aggregation,proto3" json:"aggregation,omitempty"`
	// Allocated persistent rule IDs for underlying watch.
	RuleIds []string `protobuf:"bytes,5,rep,name=rule_ids,json=ruleIds,proto3" json:"rule_ids,omitempty"`
	// Optional sink where data is automatically forwarder.
	// It can be nil, if intention for this collection rule is to aid in pulling
	// via time series watch feature (TODO: Not implemented, rule without sink has
	// no effect).
	Sink *time_series_forwarder_sink.Reference `protobuf:"bytes,6,opt,customtype=Reference,name=sink,proto3" json:"sink,omitempty"`
}

func (m *TimeSeriesCollectionRule) Reset() {
	*m = TimeSeriesCollectionRule{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_monitoring_proto_v4_time_series_collection_rule_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *TimeSeriesCollectionRule) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*TimeSeriesCollectionRule) ProtoMessage() {}

func (m *TimeSeriesCollectionRule) ProtoReflect() preflect.Message {
	mi := &edgelq_monitoring_proto_v4_time_series_collection_rule_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*TimeSeriesCollectionRule) GotenMessage() {}

// Deprecated, Use TimeSeriesCollectionRule.ProtoReflect.Descriptor instead.
func (*TimeSeriesCollectionRule) Descriptor() ([]byte, []int) {
	return edgelq_monitoring_proto_v4_time_series_collection_rule_proto_rawDescGZIP(), []int{0}
}

func (m *TimeSeriesCollectionRule) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *TimeSeriesCollectionRule) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *TimeSeriesCollectionRule) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *TimeSeriesCollectionRule) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *TimeSeriesCollectionRule) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *TimeSeriesCollectionRule) GetMetadata() *meta.Meta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *TimeSeriesCollectionRule) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *TimeSeriesCollectionRule) GetFilter() *time_serie.Filter {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *TimeSeriesCollectionRule) GetAggregation() *common.Aggregation {
	if m != nil {
		return m.Aggregation
	}
	return nil
}

func (m *TimeSeriesCollectionRule) GetRuleIds() []string {
	if m != nil {
		return m.RuleIds
	}
	return nil
}

func (m *TimeSeriesCollectionRule) GetSink() *time_series_forwarder_sink.Reference {
	if m != nil {
		return m.Sink
	}
	return nil
}

func (m *TimeSeriesCollectionRule) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "TimeSeriesCollectionRule"))
	}
	m.Name = fv
}

func (m *TimeSeriesCollectionRule) SetMetadata(fv *meta.Meta) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Metadata", "TimeSeriesCollectionRule"))
	}
	m.Metadata = fv
}

func (m *TimeSeriesCollectionRule) SetDisplayName(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "DisplayName", "TimeSeriesCollectionRule"))
	}
	m.DisplayName = fv
}

func (m *TimeSeriesCollectionRule) SetFilter(fv *time_serie.Filter) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Filter", "TimeSeriesCollectionRule"))
	}
	m.Filter = fv
}

func (m *TimeSeriesCollectionRule) SetAggregation(fv *common.Aggregation) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Aggregation", "TimeSeriesCollectionRule"))
	}
	m.Aggregation = fv
}

func (m *TimeSeriesCollectionRule) SetRuleIds(fv []string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "RuleIds", "TimeSeriesCollectionRule"))
	}
	m.RuleIds = fv
}

func (m *TimeSeriesCollectionRule) SetSink(fv *time_series_forwarder_sink.Reference) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Sink", "TimeSeriesCollectionRule"))
	}
	m.Sink = fv
}

var edgelq_monitoring_proto_v4_time_series_collection_rule_proto preflect.FileDescriptor

var edgelq_monitoring_proto_v4_time_series_collection_rule_proto_rawDesc = []byte{
	0x0a, 0x3c, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x69, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x34, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x5f, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x5f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11,
	0x6e, 0x74, 0x74, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76,
	0x34, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f,
	0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d,
	0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67,
	0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x24, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x65,
	0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x34, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcc, 0x04, 0x0a, 0x18, 0x54, 0x69, 0x6d, 0x65, 0x53,
	0x65, 0x72, 0x69, 0x65, 0x73, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x75, 0x6c, 0x65, 0x12, 0x34, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x20, 0xb2, 0xda, 0x21, 0x1c, 0x0a, 0x1a, 0x0a, 0x18, 0x54, 0x69, 0x6d, 0x65, 0x53,
	0x65, 0x72, 0x69, 0x65, 0x73, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x75, 0x6c, 0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x08, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67, 0x6f,
	0x74, 0x65, 0x6e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x2d, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70,
	0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a,
	0xca, 0xc6, 0x27, 0x06, 0x2a, 0x04, 0x22, 0x02, 0x08, 0x40, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70,
	0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x29, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x11, 0xb2, 0xda, 0x21, 0x0d, 0x1a, 0x0b, 0x0a,
	0x09, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x12, 0x40, 0x0a, 0x0b, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x6d, 0x6f,
	0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x34, 0x2e, 0x41, 0x67, 0x67, 0x72,
	0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x08, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x73,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x42, 0x04, 0xf0, 0xd9, 0x21, 0x01, 0x52, 0x07, 0x72, 0x75,
	0x6c, 0x65, 0x49, 0x64, 0x73, 0x12, 0x35, 0x0a, 0x04, 0x73, 0x69, 0x6e, 0x6b, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x21, 0xb2, 0xda, 0x21, 0x1d, 0x12, 0x1b, 0x0a, 0x17, 0x54, 0x69, 0x6d,
	0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72,
	0x53, 0x69, 0x6e, 0x6b, 0x10, 0x01, 0x52, 0x04, 0x73, 0x69, 0x6e, 0x6b, 0x3a, 0xd6, 0x01, 0xea,
	0x41, 0x7c, 0x0a, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x65,
	0x64, 0x67, 0x65, 0x6c, 0x71, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65,
	0x72, 0x69, 0x65, 0x73, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x75,
	0x6c, 0x65, 0x12, 0x4a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65,
	0x73, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x75, 0x6c, 0x65, 0x73,
	0x2f, 0x7b, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x5f, 0x63, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x7d, 0x92, 0xd9,
	0x21, 0x41, 0x0a, 0x19, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x43, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x19, 0x74,
	0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x1a, 0x07, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x38, 0x05, 0xda, 0x94, 0x23, 0x08, 0x12, 0x06, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0xe2,
	0xde, 0x21, 0x02, 0x08, 0x01, 0x42, 0x95, 0x03, 0xe8, 0xde, 0x21, 0x01, 0xd2, 0xff, 0xd0, 0x02,
	0x6f, 0x0a, 0x21, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x5f, 0x63,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x12, 0x4a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71,
	0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2f, 0x76, 0x34, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73,
	0x5f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x75, 0x6c, 0x65,
	0xa2, 0x80, 0xd1, 0x02, 0x71, 0x0a, 0x22, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x69,
	0x65, 0x73, 0x5f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x75,
	0x6c, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x4b, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65,
	0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67,
	0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x34, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x5f,
	0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x5f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x0a, 0x18, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x74, 0x74, 0x2e,
	0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x34,
	0x42, 0x1d, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x43, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x75, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x6a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6d, 0x6f,
	0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x2f, 0x76, 0x34, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x69, 0x65,
	0x73, 0x5f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x75, 0x6c,
	0x65, 0x3b, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x5f, 0x63, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	edgelq_monitoring_proto_v4_time_series_collection_rule_proto_rawDescOnce sync.Once
	edgelq_monitoring_proto_v4_time_series_collection_rule_proto_rawDescData = edgelq_monitoring_proto_v4_time_series_collection_rule_proto_rawDesc
)

func edgelq_monitoring_proto_v4_time_series_collection_rule_proto_rawDescGZIP() []byte {
	edgelq_monitoring_proto_v4_time_series_collection_rule_proto_rawDescOnce.Do(func() {
		edgelq_monitoring_proto_v4_time_series_collection_rule_proto_rawDescData = protoimpl.X.CompressGZIP(edgelq_monitoring_proto_v4_time_series_collection_rule_proto_rawDescData)
	})
	return edgelq_monitoring_proto_v4_time_series_collection_rule_proto_rawDescData
}

var edgelq_monitoring_proto_v4_time_series_collection_rule_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var edgelq_monitoring_proto_v4_time_series_collection_rule_proto_goTypes = []interface{}{
	(*TimeSeriesCollectionRule)(nil), // 0: ntt.monitoring.v4.TimeSeriesCollectionRule
	(*meta.Meta)(nil),                // 1: goten.types.Meta
	(*common.Aggregation)(nil),       // 2: ntt.monitoring.v4.Aggregation
}
var edgelq_monitoring_proto_v4_time_series_collection_rule_proto_depIdxs = []int32{
	1, // 0: ntt.monitoring.v4.TimeSeriesCollectionRule.metadata:type_name -> goten.types.Meta
	2, // 1: ntt.monitoring.v4.TimeSeriesCollectionRule.aggregation:type_name -> ntt.monitoring.v4.Aggregation
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { edgelq_monitoring_proto_v4_time_series_collection_rule_proto_init() }
func edgelq_monitoring_proto_v4_time_series_collection_rule_proto_init() {
	if edgelq_monitoring_proto_v4_time_series_collection_rule_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		edgelq_monitoring_proto_v4_time_series_collection_rule_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimeSeriesCollectionRule); i {
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
			RawDescriptor: edgelq_monitoring_proto_v4_time_series_collection_rule_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           edgelq_monitoring_proto_v4_time_series_collection_rule_proto_goTypes,
		DependencyIndexes: edgelq_monitoring_proto_v4_time_series_collection_rule_proto_depIdxs,
		MessageInfos:      edgelq_monitoring_proto_v4_time_series_collection_rule_proto_msgTypes,
	}.Build()
	edgelq_monitoring_proto_v4_time_series_collection_rule_proto = out.File
	edgelq_monitoring_proto_v4_time_series_collection_rule_proto_rawDesc = nil
	edgelq_monitoring_proto_v4_time_series_collection_rule_proto_goTypes = nil
	edgelq_monitoring_proto_v4_time_series_collection_rule_proto_depIdxs = nil
}
