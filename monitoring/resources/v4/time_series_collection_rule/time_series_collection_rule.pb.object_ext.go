// Code generated by protoc-gen-goten-object
// File: edgelq/monitoring/proto/v4/time_series_collection_rule.proto
// DO NOT EDIT!!!

package time_series_collection_rule

import (
	"fmt"
	"sort"

	"google.golang.org/protobuf/proto"
	googlefieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"

	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
)

// proto imports
import (
	common "github.com/cloudwan/edgelq-sdk/monitoring/resources/v4/common"
	project "github.com/cloudwan/edgelq-sdk/monitoring/resources/v4/project"
	time_serie "github.com/cloudwan/edgelq-sdk/monitoring/resources/v4/time_serie"
	time_series_forwarder_sink "github.com/cloudwan/edgelq-sdk/monitoring/resources/v4/time_series_forwarder_sink"
	meta "github.com/cloudwan/goten-sdk/types/meta"
)

// ensure the imports are used
var (
	_ = new(fmt.Stringer)
	_ = new(sort.Interface)

	_ = new(proto.Message)
	_ = googlefieldmaskpb.FieldMask{}

	_ = new(gotenobject.FieldPath)
)

// make sure we're using proto imports
var (
	_ = &common.LabelDescriptor{}
	_ = &project.Project{}
	_ = &time_serie.Point{}
	_ = &time_series_forwarder_sink.TimeSeriesForwarderSink{}
	_ = &meta.Meta{}
)

func (o *TimeSeriesCollectionRule) GotenObjectExt() {}

func (o *TimeSeriesCollectionRule) MakeFullFieldMask() *TimeSeriesCollectionRule_FieldMask {
	return FullTimeSeriesCollectionRule_FieldMask()
}

func (o *TimeSeriesCollectionRule) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullTimeSeriesCollectionRule_FieldMask()
}

func (o *TimeSeriesCollectionRule) MakeDiffFieldMask(other *TimeSeriesCollectionRule) *TimeSeriesCollectionRule_FieldMask {
	if o == nil && other == nil {
		return &TimeSeriesCollectionRule_FieldMask{}
	}
	if o == nil || other == nil {
		return FullTimeSeriesCollectionRule_FieldMask()
	}

	res := &TimeSeriesCollectionRule_FieldMask{}
	if o.GetName().String() != other.GetName().String() {
		res.Paths = append(res.Paths, &TimeSeriesCollectionRule_FieldTerminalPath{selector: TimeSeriesCollectionRule_FieldPathSelectorName})
	}
	{
		subMask := o.GetMetadata().MakeDiffFieldMask(other.GetMetadata())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &TimeSeriesCollectionRule_FieldTerminalPath{selector: TimeSeriesCollectionRule_FieldPathSelectorMetadata})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &TimeSeriesCollectionRule_FieldSubPath{selector: TimeSeriesCollectionRule_FieldPathSelectorMetadata, subPath: subpath})
			}
		}
	}
	if o.GetDisplayName() != other.GetDisplayName() {
		res.Paths = append(res.Paths, &TimeSeriesCollectionRule_FieldTerminalPath{selector: TimeSeriesCollectionRule_FieldPathSelectorDisplayName})
	}
	if o.GetFilter().String() != other.GetFilter().String() {
		res.Paths = append(res.Paths, &TimeSeriesCollectionRule_FieldTerminalPath{selector: TimeSeriesCollectionRule_FieldPathSelectorFilter})
	}
	{
		subMask := o.GetAggregation().MakeDiffFieldMask(other.GetAggregation())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &TimeSeriesCollectionRule_FieldTerminalPath{selector: TimeSeriesCollectionRule_FieldPathSelectorAggregation})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &TimeSeriesCollectionRule_FieldSubPath{selector: TimeSeriesCollectionRule_FieldPathSelectorAggregation, subPath: subpath})
			}
		}
	}

	if len(o.GetRuleIds()) == len(other.GetRuleIds()) {
		for i, lValue := range o.GetRuleIds() {
			rValue := other.GetRuleIds()[i]
			if lValue != rValue {
				res.Paths = append(res.Paths, &TimeSeriesCollectionRule_FieldTerminalPath{selector: TimeSeriesCollectionRule_FieldPathSelectorRuleIds})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &TimeSeriesCollectionRule_FieldTerminalPath{selector: TimeSeriesCollectionRule_FieldPathSelectorRuleIds})
	}
	if o.GetSink().String() != other.GetSink().String() {
		res.Paths = append(res.Paths, &TimeSeriesCollectionRule_FieldTerminalPath{selector: TimeSeriesCollectionRule_FieldPathSelectorSink})
	}
	return res
}

func (o *TimeSeriesCollectionRule) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*TimeSeriesCollectionRule))
}

func (o *TimeSeriesCollectionRule) Clone() *TimeSeriesCollectionRule {
	if o == nil {
		return nil
	}
	result := &TimeSeriesCollectionRule{}
	if o.Name == nil {
		result.Name = nil
	} else if data, err := o.Name.ProtoString(); err != nil {
		panic(err)
	} else {
		result.Name = &Name{}
		if err := result.Name.ParseProtoString(data); err != nil {
			panic(err)
		}
	}
	result.Metadata = o.Metadata.Clone()
	result.DisplayName = o.DisplayName
	if o.Filter == nil {
		result.Filter = nil
	} else if data, err := o.Filter.ProtoString(); err != nil {
		panic(err)
	} else {
		result.Filter = &time_serie.Filter{}
		if err := result.Filter.ParseProtoString(data); err != nil {
			panic(err)
		}
	}
	result.Aggregation = o.Aggregation.Clone()
	result.RuleIds = make([]string, len(o.RuleIds))
	for i, sourceValue := range o.RuleIds {
		result.RuleIds[i] = sourceValue
	}
	if o.Sink == nil {
		result.Sink = nil
	} else if data, err := o.Sink.ProtoString(); err != nil {
		panic(err)
	} else {
		result.Sink = &time_series_forwarder_sink.Reference{}
		if err := result.Sink.ParseProtoString(data); err != nil {
			panic(err)
		}
	}
	return result
}

func (o *TimeSeriesCollectionRule) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *TimeSeriesCollectionRule) Merge(source *TimeSeriesCollectionRule) {
	if source.GetName() != nil {
		if data, err := source.GetName().ProtoString(); err != nil {
			panic(err)
		} else {
			o.Name = &Name{}
			if err := o.Name.ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	} else {
		o.Name = nil
	}
	if source.GetMetadata() != nil {
		if o.Metadata == nil {
			o.Metadata = new(meta.Meta)
		}
		o.Metadata.Merge(source.GetMetadata())
	}
	o.DisplayName = source.GetDisplayName()
	if source.GetFilter() != nil {
		if data, err := source.GetFilter().ProtoString(); err != nil {
			panic(err)
		} else {
			o.Filter = &time_serie.Filter{}
			if err := o.Filter.ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	} else {
		o.Filter = nil
	}
	if source.GetAggregation() != nil {
		if o.Aggregation == nil {
			o.Aggregation = new(common.Aggregation)
		}
		o.Aggregation.Merge(source.GetAggregation())
	}
	for _, sourceValue := range source.GetRuleIds() {
		exists := false
		for _, currentValue := range o.RuleIds {
			if currentValue == sourceValue {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement string
			newDstElement = sourceValue
			o.RuleIds = append(o.RuleIds, newDstElement)
		}
	}

	if source.GetSink() != nil {
		if data, err := source.GetSink().ProtoString(); err != nil {
			panic(err)
		} else {
			o.Sink = &time_series_forwarder_sink.Reference{}
			if err := o.Sink.ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	} else {
		o.Sink = nil
	}
}

func (o *TimeSeriesCollectionRule) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*TimeSeriesCollectionRule))
}
