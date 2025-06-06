// Code generated by protoc-gen-goten-object
// File: edgelq/monitoring/proto/v4/alerting_condition.proto
// DO NOT EDIT!!!

package alerting_condition

import (
	"fmt"
	"sort"

	"google.golang.org/protobuf/proto"
	googlefieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"

	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
)

// proto imports
import (
	alerting_policy "github.com/cloudwan/edgelq-sdk/monitoring/resources/v4/alerting_policy"
	common "github.com/cloudwan/edgelq-sdk/monitoring/resources/v4/common"
	time_serie "github.com/cloudwan/edgelq-sdk/monitoring/resources/v4/time_serie"
	meta "github.com/cloudwan/goten-sdk/types/meta"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
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
	_ = &alerting_policy.AlertingPolicy{}
	_ = &common.LabelDescriptor{}
	_ = &time_serie.Point{}
	_ = &durationpb.Duration{}
	_ = &meta.Meta{}
)

func (o *AlertingCondition) GotenObjectExt() {}

func (o *AlertingCondition) MakeFullFieldMask() *AlertingCondition_FieldMask {
	return FullAlertingCondition_FieldMask()
}

func (o *AlertingCondition) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullAlertingCondition_FieldMask()
}

func (o *AlertingCondition) MakeDiffFieldMask(other *AlertingCondition) *AlertingCondition_FieldMask {
	if o == nil && other == nil {
		return &AlertingCondition_FieldMask{}
	}
	if o == nil || other == nil {
		return FullAlertingCondition_FieldMask()
	}

	res := &AlertingCondition_FieldMask{}
	if o.GetName().String() != other.GetName().String() {
		res.Paths = append(res.Paths, &AlertingCondition_FieldTerminalPath{selector: AlertingCondition_FieldPathSelectorName})
	}
	{
		subMask := o.GetMetadata().MakeDiffFieldMask(other.GetMetadata())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &AlertingCondition_FieldTerminalPath{selector: AlertingCondition_FieldPathSelectorMetadata})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &AlertingCondition_FieldSubPath{selector: AlertingCondition_FieldPathSelectorMetadata, subPath: subpath})
			}
		}
	}
	if o.GetDisplayName() != other.GetDisplayName() {
		res.Paths = append(res.Paths, &AlertingCondition_FieldTerminalPath{selector: AlertingCondition_FieldPathSelectorDisplayName})
	}
	if o.GetDescription() != other.GetDescription() {
		res.Paths = append(res.Paths, &AlertingCondition_FieldTerminalPath{selector: AlertingCondition_FieldPathSelectorDescription})
	}
	{
		subMask := o.GetSpec().MakeDiffFieldMask(other.GetSpec())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &AlertingCondition_FieldTerminalPath{selector: AlertingCondition_FieldPathSelectorSpec})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &AlertingCondition_FieldSubPath{selector: AlertingCondition_FieldPathSelectorSpec, subPath: subpath})
			}
		}
	}
	{
		subMask := o.GetState().MakeDiffFieldMask(other.GetState())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &AlertingCondition_FieldTerminalPath{selector: AlertingCondition_FieldPathSelectorState})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &AlertingCondition_FieldSubPath{selector: AlertingCondition_FieldPathSelectorState, subPath: subpath})
			}
		}
	}
	return res
}

func (o *AlertingCondition) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*AlertingCondition))
}

func (o *AlertingCondition) Clone() *AlertingCondition {
	if o == nil {
		return nil
	}
	result := &AlertingCondition{}
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
	result.Description = o.Description
	result.Spec = o.Spec.Clone()
	result.State = o.State.Clone()
	return result
}

func (o *AlertingCondition) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *AlertingCondition) Merge(source *AlertingCondition) {
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
	o.Description = source.GetDescription()
	if source.GetSpec() != nil {
		if o.Spec == nil {
			o.Spec = new(AlertingCondition_Spec)
		}
		o.Spec.Merge(source.GetSpec())
	}
	if source.GetState() != nil {
		if o.State == nil {
			o.State = new(AlertingCondition_State)
		}
		o.State.Merge(source.GetState())
	}
}

func (o *AlertingCondition) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*AlertingCondition))
}

func (o *AlertingCondition_Spec) GotenObjectExt() {}

func (o *AlertingCondition_Spec) MakeFullFieldMask() *AlertingCondition_Spec_FieldMask {
	return FullAlertingCondition_Spec_FieldMask()
}

func (o *AlertingCondition_Spec) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullAlertingCondition_Spec_FieldMask()
}

func (o *AlertingCondition_Spec) MakeDiffFieldMask(other *AlertingCondition_Spec) *AlertingCondition_Spec_FieldMask {
	if o == nil && other == nil {
		return &AlertingCondition_Spec_FieldMask{}
	}
	if o == nil || other == nil {
		return FullAlertingCondition_Spec_FieldMask()
	}

	res := &AlertingCondition_Spec_FieldMask{}
	{
		subMask := o.GetTimeSeries().MakeDiffFieldMask(other.GetTimeSeries())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &AlertingConditionSpec_FieldTerminalPath{selector: AlertingConditionSpec_FieldPathSelectorTimeSeries})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &AlertingConditionSpec_FieldSubPath{selector: AlertingConditionSpec_FieldPathSelectorTimeSeries, subPath: subpath})
			}
		}
	}
	return res
}

func (o *AlertingCondition_Spec) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*AlertingCondition_Spec))
}

func (o *AlertingCondition_Spec) Clone() *AlertingCondition_Spec {
	if o == nil {
		return nil
	}
	result := &AlertingCondition_Spec{}
	result.TimeSeries = o.TimeSeries.Clone()
	return result
}

func (o *AlertingCondition_Spec) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *AlertingCondition_Spec) Merge(source *AlertingCondition_Spec) {
	if source.GetTimeSeries() != nil {
		if o.TimeSeries == nil {
			o.TimeSeries = new(AlertingCondition_Spec_TimeSeries)
		}
		o.TimeSeries.Merge(source.GetTimeSeries())
	}
}

func (o *AlertingCondition_Spec) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*AlertingCondition_Spec))
}

func (o *AlertingCondition_State) GotenObjectExt() {}

func (o *AlertingCondition_State) MakeFullFieldMask() *AlertingCondition_State_FieldMask {
	return FullAlertingCondition_State_FieldMask()
}

func (o *AlertingCondition_State) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullAlertingCondition_State_FieldMask()
}

func (o *AlertingCondition_State) MakeDiffFieldMask(other *AlertingCondition_State) *AlertingCondition_State_FieldMask {
	if o == nil && other == nil {
		return &AlertingCondition_State_FieldMask{}
	}
	if o == nil || other == nil {
		return FullAlertingCondition_State_FieldMask()
	}

	res := &AlertingCondition_State_FieldMask{}
	if o.GetFiringAlertsCount() != other.GetFiringAlertsCount() {
		res.Paths = append(res.Paths, &AlertingConditionState_FieldTerminalPath{selector: AlertingConditionState_FieldPathSelectorFiringAlertsCount})
	}
	return res
}

func (o *AlertingCondition_State) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*AlertingCondition_State))
}

func (o *AlertingCondition_State) Clone() *AlertingCondition_State {
	if o == nil {
		return nil
	}
	result := &AlertingCondition_State{}
	result.FiringAlertsCount = o.FiringAlertsCount
	return result
}

func (o *AlertingCondition_State) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *AlertingCondition_State) Merge(source *AlertingCondition_State) {
	o.FiringAlertsCount = source.GetFiringAlertsCount()
}

func (o *AlertingCondition_State) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*AlertingCondition_State))
}

func (o *AlertingCondition_Spec_TimeSeries) GotenObjectExt() {}

func (o *AlertingCondition_Spec_TimeSeries) MakeFullFieldMask() *AlertingCondition_Spec_TimeSeries_FieldMask {
	return FullAlertingCondition_Spec_TimeSeries_FieldMask()
}

func (o *AlertingCondition_Spec_TimeSeries) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullAlertingCondition_Spec_TimeSeries_FieldMask()
}

func (o *AlertingCondition_Spec_TimeSeries) MakeDiffFieldMask(other *AlertingCondition_Spec_TimeSeries) *AlertingCondition_Spec_TimeSeries_FieldMask {
	if o == nil && other == nil {
		return &AlertingCondition_Spec_TimeSeries_FieldMask{}
	}
	if o == nil || other == nil {
		return FullAlertingCondition_Spec_TimeSeries_FieldMask()
	}

	res := &AlertingCondition_Spec_TimeSeries_FieldMask{}
	{
		subMask := o.GetQuery().MakeDiffFieldMask(other.GetQuery())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &AlertingConditionSpecTimeSeries_FieldTerminalPath{selector: AlertingConditionSpecTimeSeries_FieldPathSelectorQuery})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &AlertingConditionSpecTimeSeries_FieldSubPath{selector: AlertingConditionSpecTimeSeries_FieldPathSelectorQuery, subPath: subpath})
			}
		}
	}
	{
		subMask := o.GetThreshold().MakeDiffFieldMask(other.GetThreshold())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &AlertingConditionSpecTimeSeries_FieldTerminalPath{selector: AlertingConditionSpecTimeSeries_FieldPathSelectorThreshold})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &AlertingConditionSpecTimeSeries_FieldSubPath{selector: AlertingConditionSpecTimeSeries_FieldPathSelectorThreshold, subPath: subpath})
			}
		}
	}
	{
		subMask := o.GetCombineThreshold().MakeDiffFieldMask(other.GetCombineThreshold())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &AlertingConditionSpecTimeSeries_FieldTerminalPath{selector: AlertingConditionSpecTimeSeries_FieldPathSelectorCombineThreshold})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &AlertingConditionSpecTimeSeries_FieldSubPath{selector: AlertingConditionSpecTimeSeries_FieldPathSelectorCombineThreshold, subPath: subpath})
			}
		}
	}
	if !proto.Equal(o.GetDuration(), other.GetDuration()) {
		res.Paths = append(res.Paths, &AlertingConditionSpecTimeSeries_FieldTerminalPath{selector: AlertingConditionSpecTimeSeries_FieldPathSelectorDuration})
	}
	return res
}

func (o *AlertingCondition_Spec_TimeSeries) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*AlertingCondition_Spec_TimeSeries))
}

func (o *AlertingCondition_Spec_TimeSeries) Clone() *AlertingCondition_Spec_TimeSeries {
	if o == nil {
		return nil
	}
	result := &AlertingCondition_Spec_TimeSeries{}
	result.Query = o.Query.Clone()
	result.Threshold = o.Threshold.Clone()
	result.CombineThreshold = o.CombineThreshold.Clone()
	result.Duration = proto.Clone(o.Duration).(*durationpb.Duration)
	return result
}

func (o *AlertingCondition_Spec_TimeSeries) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *AlertingCondition_Spec_TimeSeries) Merge(source *AlertingCondition_Spec_TimeSeries) {
	if source.GetQuery() != nil {
		if o.Query == nil {
			o.Query = new(AlertingCondition_Spec_TimeSeries_Query)
		}
		o.Query.Merge(source.GetQuery())
	}
	if source.GetThreshold() != nil {
		if o.Threshold == nil {
			o.Threshold = new(AlertingCondition_Spec_TimeSeries_Threshold)
		}
		o.Threshold.Merge(source.GetThreshold())
	}
	if source.GetCombineThreshold() != nil {
		if o.CombineThreshold == nil {
			o.CombineThreshold = new(AlertingCondition_Spec_TimeSeries_CombineThreshold)
		}
		o.CombineThreshold.Merge(source.GetCombineThreshold())
	}
	if source.GetDuration() != nil {
		if o.Duration == nil {
			o.Duration = new(durationpb.Duration)
		}
		proto.Merge(o.Duration, source.GetDuration())
	}
}

func (o *AlertingCondition_Spec_TimeSeries) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*AlertingCondition_Spec_TimeSeries))
}

func (o *AlertingCondition_Spec_TimeSeries_Query) GotenObjectExt() {}

func (o *AlertingCondition_Spec_TimeSeries_Query) MakeFullFieldMask() *AlertingCondition_Spec_TimeSeries_Query_FieldMask {
	return FullAlertingCondition_Spec_TimeSeries_Query_FieldMask()
}

func (o *AlertingCondition_Spec_TimeSeries_Query) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullAlertingCondition_Spec_TimeSeries_Query_FieldMask()
}

func (o *AlertingCondition_Spec_TimeSeries_Query) MakeDiffFieldMask(other *AlertingCondition_Spec_TimeSeries_Query) *AlertingCondition_Spec_TimeSeries_Query_FieldMask {
	if o == nil && other == nil {
		return &AlertingCondition_Spec_TimeSeries_Query_FieldMask{}
	}
	if o == nil || other == nil {
		return FullAlertingCondition_Spec_TimeSeries_Query_FieldMask()
	}

	res := &AlertingCondition_Spec_TimeSeries_Query_FieldMask{}
	if o.GetFilter().String() != other.GetFilter().String() {
		res.Paths = append(res.Paths, &AlertingConditionSpecTimeSeriesQuery_FieldTerminalPath{selector: AlertingConditionSpecTimeSeriesQuery_FieldPathSelectorFilter})
	}
	{
		subMask := o.GetSelector().MakeDiffFieldMask(other.GetSelector())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &AlertingConditionSpecTimeSeriesQuery_FieldTerminalPath{selector: AlertingConditionSpecTimeSeriesQuery_FieldPathSelectorSelector})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &AlertingConditionSpecTimeSeriesQuery_FieldSubPath{selector: AlertingConditionSpecTimeSeriesQuery_FieldPathSelectorSelector, subPath: subpath})
			}
		}
	}
	{
		subMask := o.GetAggregation().MakeDiffFieldMask(other.GetAggregation())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &AlertingConditionSpecTimeSeriesQuery_FieldTerminalPath{selector: AlertingConditionSpecTimeSeriesQuery_FieldPathSelectorAggregation})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &AlertingConditionSpecTimeSeriesQuery_FieldSubPath{selector: AlertingConditionSpecTimeSeriesQuery_FieldPathSelectorAggregation, subPath: subpath})
			}
		}
	}

	if len(o.GetPerMetricAggregations()) == len(other.GetPerMetricAggregations()) {
		for i, lValue := range o.GetPerMetricAggregations() {
			rValue := other.GetPerMetricAggregations()[i]
			if len(lValue.MakeDiffFieldMask(rValue).Paths) > 0 {
				res.Paths = append(res.Paths, &AlertingConditionSpecTimeSeriesQuery_FieldTerminalPath{selector: AlertingConditionSpecTimeSeriesQuery_FieldPathSelectorPerMetricAggregations})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &AlertingConditionSpecTimeSeriesQuery_FieldTerminalPath{selector: AlertingConditionSpecTimeSeriesQuery_FieldPathSelectorPerMetricAggregations})
	}
	return res
}

func (o *AlertingCondition_Spec_TimeSeries_Query) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*AlertingCondition_Spec_TimeSeries_Query))
}

func (o *AlertingCondition_Spec_TimeSeries_Query) Clone() *AlertingCondition_Spec_TimeSeries_Query {
	if o == nil {
		return nil
	}
	result := &AlertingCondition_Spec_TimeSeries_Query{}
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
	result.Selector = o.Selector.Clone()
	result.Aggregation = o.Aggregation.Clone()
	result.PerMetricAggregations = map[string]*common.Aggregation{}
	for key, sourceValue := range o.PerMetricAggregations {
		result.PerMetricAggregations[key] = sourceValue.Clone()
	}
	return result
}

func (o *AlertingCondition_Spec_TimeSeries_Query) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *AlertingCondition_Spec_TimeSeries_Query) Merge(source *AlertingCondition_Spec_TimeSeries_Query) {
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
	if source.GetSelector() != nil {
		if o.Selector == nil {
			o.Selector = new(common.TimeSeriesSelector)
		}
		o.Selector.Merge(source.GetSelector())
	}
	if source.GetAggregation() != nil {
		if o.Aggregation == nil {
			o.Aggregation = new(common.Aggregation)
		}
		o.Aggregation.Merge(source.GetAggregation())
	}
	if source.GetPerMetricAggregations() != nil {
		if o.PerMetricAggregations == nil {
			o.PerMetricAggregations = make(map[string]*common.Aggregation, len(source.GetPerMetricAggregations()))
		}
		for key, sourceValue := range source.GetPerMetricAggregations() {
			if sourceValue != nil {
				if o.PerMetricAggregations[key] == nil {
					o.PerMetricAggregations[key] = new(common.Aggregation)
				}
				o.PerMetricAggregations[key].Merge(sourceValue)
			}
		}
	}
}

func (o *AlertingCondition_Spec_TimeSeries_Query) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*AlertingCondition_Spec_TimeSeries_Query))
}

func (o *AlertingCondition_Spec_TimeSeries_Threshold) GotenObjectExt() {}

func (o *AlertingCondition_Spec_TimeSeries_Threshold) MakeFullFieldMask() *AlertingCondition_Spec_TimeSeries_Threshold_FieldMask {
	return FullAlertingCondition_Spec_TimeSeries_Threshold_FieldMask()
}

func (o *AlertingCondition_Spec_TimeSeries_Threshold) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullAlertingCondition_Spec_TimeSeries_Threshold_FieldMask()
}

func (o *AlertingCondition_Spec_TimeSeries_Threshold) MakeDiffFieldMask(other *AlertingCondition_Spec_TimeSeries_Threshold) *AlertingCondition_Spec_TimeSeries_Threshold_FieldMask {
	if o == nil && other == nil {
		return &AlertingCondition_Spec_TimeSeries_Threshold_FieldMask{}
	}
	if o == nil || other == nil {
		return FullAlertingCondition_Spec_TimeSeries_Threshold_FieldMask()
	}

	res := &AlertingCondition_Spec_TimeSeries_Threshold_FieldMask{}
	if o.GetCompare() != other.GetCompare() {
		res.Paths = append(res.Paths, &AlertingConditionSpecTimeSeriesThreshold_FieldTerminalPath{selector: AlertingConditionSpecTimeSeriesThreshold_FieldPathSelectorCompare})
	}
	if o.GetValue() != other.GetValue() {
		res.Paths = append(res.Paths, &AlertingConditionSpecTimeSeriesThreshold_FieldTerminalPath{selector: AlertingConditionSpecTimeSeriesThreshold_FieldPathSelectorValue})
	}
	return res
}

func (o *AlertingCondition_Spec_TimeSeries_Threshold) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*AlertingCondition_Spec_TimeSeries_Threshold))
}

func (o *AlertingCondition_Spec_TimeSeries_Threshold) Clone() *AlertingCondition_Spec_TimeSeries_Threshold {
	if o == nil {
		return nil
	}
	result := &AlertingCondition_Spec_TimeSeries_Threshold{}
	result.Compare = o.Compare
	result.Value = o.Value
	return result
}

func (o *AlertingCondition_Spec_TimeSeries_Threshold) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *AlertingCondition_Spec_TimeSeries_Threshold) Merge(source *AlertingCondition_Spec_TimeSeries_Threshold) {
	o.Compare = source.GetCompare()
	o.Value = source.GetValue()
}

func (o *AlertingCondition_Spec_TimeSeries_Threshold) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*AlertingCondition_Spec_TimeSeries_Threshold))
}

func (o *AlertingCondition_Spec_TimeSeries_CombineThreshold) GotenObjectExt() {}

func (o *AlertingCondition_Spec_TimeSeries_CombineThreshold) MakeFullFieldMask() *AlertingCondition_Spec_TimeSeries_CombineThreshold_FieldMask {
	return FullAlertingCondition_Spec_TimeSeries_CombineThreshold_FieldMask()
}

func (o *AlertingCondition_Spec_TimeSeries_CombineThreshold) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullAlertingCondition_Spec_TimeSeries_CombineThreshold_FieldMask()
}

func (o *AlertingCondition_Spec_TimeSeries_CombineThreshold) MakeDiffFieldMask(other *AlertingCondition_Spec_TimeSeries_CombineThreshold) *AlertingCondition_Spec_TimeSeries_CombineThreshold_FieldMask {
	if o == nil && other == nil {
		return &AlertingCondition_Spec_TimeSeries_CombineThreshold_FieldMask{}
	}
	if o == nil || other == nil {
		return FullAlertingCondition_Spec_TimeSeries_CombineThreshold_FieldMask()
	}

	res := &AlertingCondition_Spec_TimeSeries_CombineThreshold_FieldMask{}

	if len(o.GetPerMetric()) == len(other.GetPerMetric()) {
		for i, lValue := range o.GetPerMetric() {
			rValue := other.GetPerMetric()[i]
			if len(lValue.MakeDiffFieldMask(rValue).Paths) > 0 {
				res.Paths = append(res.Paths, &AlertingConditionSpecTimeSeriesCombineThreshold_FieldTerminalPath{selector: AlertingConditionSpecTimeSeriesCombineThreshold_FieldPathSelectorPerMetric})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &AlertingConditionSpecTimeSeriesCombineThreshold_FieldTerminalPath{selector: AlertingConditionSpecTimeSeriesCombineThreshold_FieldPathSelectorPerMetric})
	}
	if o.GetMainMetricType() != other.GetMainMetricType() {
		res.Paths = append(res.Paths, &AlertingConditionSpecTimeSeriesCombineThreshold_FieldTerminalPath{selector: AlertingConditionSpecTimeSeriesCombineThreshold_FieldPathSelectorMainMetricType})
	}

	if len(o.GetPerMetricTypeKv()) == len(other.GetPerMetricTypeKv()) {
		for i, lValue := range o.GetPerMetricTypeKv() {
			rValue := other.GetPerMetricTypeKv()[i]
			if string(lValue) != string(rValue) {
				res.Paths = append(res.Paths, &AlertingConditionSpecTimeSeriesCombineThreshold_FieldTerminalPath{selector: AlertingConditionSpecTimeSeriesCombineThreshold_FieldPathSelectorPerMetricTypeKv})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &AlertingConditionSpecTimeSeriesCombineThreshold_FieldTerminalPath{selector: AlertingConditionSpecTimeSeriesCombineThreshold_FieldPathSelectorPerMetricTypeKv})
	}
	return res
}

func (o *AlertingCondition_Spec_TimeSeries_CombineThreshold) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*AlertingCondition_Spec_TimeSeries_CombineThreshold))
}

func (o *AlertingCondition_Spec_TimeSeries_CombineThreshold) Clone() *AlertingCondition_Spec_TimeSeries_CombineThreshold {
	if o == nil {
		return nil
	}
	result := &AlertingCondition_Spec_TimeSeries_CombineThreshold{}
	result.PerMetric = map[string]*AlertingCondition_Spec_TimeSeries_Threshold{}
	for key, sourceValue := range o.PerMetric {
		result.PerMetric[key] = sourceValue.Clone()
	}
	result.MainMetricType = o.MainMetricType
	result.PerMetricTypeKv = map[string][]byte{}
	for key, sourceValue := range o.PerMetricTypeKv {
		result.PerMetricTypeKv[key] = sourceValue
	}
	return result
}

func (o *AlertingCondition_Spec_TimeSeries_CombineThreshold) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *AlertingCondition_Spec_TimeSeries_CombineThreshold) Merge(source *AlertingCondition_Spec_TimeSeries_CombineThreshold) {
	if source.GetPerMetric() != nil {
		if o.PerMetric == nil {
			o.PerMetric = make(map[string]*AlertingCondition_Spec_TimeSeries_Threshold, len(source.GetPerMetric()))
		}
		for key, sourceValue := range source.GetPerMetric() {
			if sourceValue != nil {
				if o.PerMetric[key] == nil {
					o.PerMetric[key] = new(AlertingCondition_Spec_TimeSeries_Threshold)
				}
				o.PerMetric[key].Merge(sourceValue)
			}
		}
	}
	o.MainMetricType = source.GetMainMetricType()
	if source.GetPerMetricTypeKv() != nil {
		if o.PerMetricTypeKv == nil {
			o.PerMetricTypeKv = make(map[string][]byte, len(source.GetPerMetricTypeKv()))
		}
		for key, sourceValue := range source.GetPerMetricTypeKv() {
			o.PerMetricTypeKv[key] = sourceValue
		}
	}
}

func (o *AlertingCondition_Spec_TimeSeries_CombineThreshold) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*AlertingCondition_Spec_TimeSeries_CombineThreshold))
}
