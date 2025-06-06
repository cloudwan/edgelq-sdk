// Code generated by protoc-gen-goten-object
// File: edgelq/alerting/proto/v1/log_condition.proto
// DO NOT EDIT!!!

package log_condition

import (
	"fmt"
	"sort"

	"google.golang.org/protobuf/proto"
	googlefieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"

	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
)

// proto imports
import (
	document "github.com/cloudwan/edgelq-sdk/alerting/resources/v1/document"
	policy "github.com/cloudwan/edgelq-sdk/alerting/resources/v1/policy"
	logging_log "github.com/cloudwan/edgelq-sdk/logging/resources/v1/log"
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
	_ = &document.Document{}
	_ = &policy.Policy{}
	_ = &logging_log.Log{}
	_ = &durationpb.Duration{}
	_ = &meta.Meta{}
)

func (o *LogCondition) GotenObjectExt() {}

func (o *LogCondition) MakeFullFieldMask() *LogCondition_FieldMask {
	return FullLogCondition_FieldMask()
}

func (o *LogCondition) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullLogCondition_FieldMask()
}

func (o *LogCondition) MakeDiffFieldMask(other *LogCondition) *LogCondition_FieldMask {
	if o == nil && other == nil {
		return &LogCondition_FieldMask{}
	}
	if o == nil || other == nil {
		return FullLogCondition_FieldMask()
	}

	res := &LogCondition_FieldMask{}
	if o.GetName().String() != other.GetName().String() {
		res.Paths = append(res.Paths, &LogCondition_FieldTerminalPath{selector: LogCondition_FieldPathSelectorName})
	}
	{
		subMask := o.GetMetadata().MakeDiffFieldMask(other.GetMetadata())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &LogCondition_FieldTerminalPath{selector: LogCondition_FieldPathSelectorMetadata})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &LogCondition_FieldSubPath{selector: LogCondition_FieldPathSelectorMetadata, subPath: subpath})
			}
		}
	}
	if o.GetDisplayName() != other.GetDisplayName() {
		res.Paths = append(res.Paths, &LogCondition_FieldTerminalPath{selector: LogCondition_FieldPathSelectorDisplayName})
	}
	if o.GetDescription() != other.GetDescription() {
		res.Paths = append(res.Paths, &LogCondition_FieldTerminalPath{selector: LogCondition_FieldPathSelectorDescription})
	}

	if len(o.GetSupportingDocs()) == len(other.GetSupportingDocs()) {
		for i, lValue := range o.GetSupportingDocs() {
			rValue := other.GetSupportingDocs()[i]
			if lValue.String() != rValue.String() {
				res.Paths = append(res.Paths, &LogCondition_FieldTerminalPath{selector: LogCondition_FieldPathSelectorSupportingDocs})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &LogCondition_FieldTerminalPath{selector: LogCondition_FieldPathSelectorSupportingDocs})
	}
	{
		subMask := o.GetSpec().MakeDiffFieldMask(other.GetSpec())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &LogCondition_FieldTerminalPath{selector: LogCondition_FieldPathSelectorSpec})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &LogCondition_FieldSubPath{selector: LogCondition_FieldPathSelectorSpec, subPath: subpath})
			}
		}
	}
	{
		subMask := o.GetInternal().MakeDiffFieldMask(other.GetInternal())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &LogCondition_FieldTerminalPath{selector: LogCondition_FieldPathSelectorInternal})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &LogCondition_FieldSubPath{selector: LogCondition_FieldPathSelectorInternal, subPath: subpath})
			}
		}
	}
	return res
}

func (o *LogCondition) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*LogCondition))
}

func (o *LogCondition) Clone() *LogCondition {
	if o == nil {
		return nil
	}
	result := &LogCondition{}
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
	result.SupportingDocs = make([]*document.Reference, len(o.SupportingDocs))
	for i, sourceValue := range o.SupportingDocs {
		if sourceValue == nil {
			result.SupportingDocs[i] = nil
		} else if data, err := sourceValue.ProtoString(); err != nil {
			panic(err)
		} else {
			result.SupportingDocs[i] = &document.Reference{}
			if err := result.SupportingDocs[i].ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	}
	result.Spec = o.Spec.Clone()
	result.Internal = o.Internal.Clone()
	return result
}

func (o *LogCondition) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *LogCondition) Merge(source *LogCondition) {
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
	for _, sourceValue := range source.GetSupportingDocs() {
		exists := false
		for _, currentValue := range o.SupportingDocs {
			leftProtoStr, _ := currentValue.ProtoString()
			rightProtoStr, _ := sourceValue.ProtoString()
			if leftProtoStr == rightProtoStr {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement *document.Reference
			if sourceValue != nil {
				if data, err := sourceValue.ProtoString(); err != nil {
					panic(err)
				} else {
					newDstElement = &document.Reference{}
					if err := newDstElement.ParseProtoString(data); err != nil {
						panic(err)
					}
				}
			}
			o.SupportingDocs = append(o.SupportingDocs, newDstElement)
		}
	}

	if source.GetSpec() != nil {
		if o.Spec == nil {
			o.Spec = new(LogCondition_Spec)
		}
		o.Spec.Merge(source.GetSpec())
	}
	if source.GetInternal() != nil {
		if o.Internal == nil {
			o.Internal = new(LogCondition_Internal)
		}
		o.Internal.Merge(source.GetInternal())
	}
}

func (o *LogCondition) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*LogCondition))
}

func (o *LogCondition_Spec) GotenObjectExt() {}

func (o *LogCondition_Spec) MakeFullFieldMask() *LogCondition_Spec_FieldMask {
	return FullLogCondition_Spec_FieldMask()
}

func (o *LogCondition_Spec) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullLogCondition_Spec_FieldMask()
}

func (o *LogCondition_Spec) MakeDiffFieldMask(other *LogCondition_Spec) *LogCondition_Spec_FieldMask {
	if o == nil && other == nil {
		return &LogCondition_Spec_FieldMask{}
	}
	if o == nil || other == nil {
		return FullLogCondition_Spec_FieldMask()
	}

	res := &LogCondition_Spec_FieldMask{}
	{
		subMask := o.GetQuery().MakeDiffFieldMask(other.GetQuery())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &LogConditionSpec_FieldTerminalPath{selector: LogConditionSpec_FieldPathSelectorQuery})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &LogConditionSpec_FieldSubPath{selector: LogConditionSpec_FieldPathSelectorQuery, subPath: subpath})
			}
		}
	}

	if len(o.GetGroupByLabels()) == len(other.GetGroupByLabels()) {
		for i, lValue := range o.GetGroupByLabels() {
			rValue := other.GetGroupByLabels()[i]
			if lValue != rValue {
				res.Paths = append(res.Paths, &LogConditionSpec_FieldTerminalPath{selector: LogConditionSpec_FieldPathSelectorGroupByLabels})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &LogConditionSpec_FieldTerminalPath{selector: LogConditionSpec_FieldPathSelectorGroupByLabels})
	}
	return res
}

func (o *LogCondition_Spec) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*LogCondition_Spec))
}

func (o *LogCondition_Spec) Clone() *LogCondition_Spec {
	if o == nil {
		return nil
	}
	result := &LogCondition_Spec{}
	result.Query = o.Query.Clone()
	result.GroupByLabels = make([]string, len(o.GroupByLabels))
	for i, sourceValue := range o.GroupByLabels {
		result.GroupByLabels[i] = sourceValue
	}
	return result
}

func (o *LogCondition_Spec) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *LogCondition_Spec) Merge(source *LogCondition_Spec) {
	if source.GetQuery() != nil {
		if o.Query == nil {
			o.Query = new(LogCondition_Spec_Query)
		}
		o.Query.Merge(source.GetQuery())
	}
	for _, sourceValue := range source.GetGroupByLabels() {
		exists := false
		for _, currentValue := range o.GroupByLabels {
			if currentValue == sourceValue {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement string
			newDstElement = sourceValue
			o.GroupByLabels = append(o.GroupByLabels, newDstElement)
		}
	}

}

func (o *LogCondition_Spec) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*LogCondition_Spec))
}

func (o *LogCondition_Internal) GotenObjectExt() {}

func (o *LogCondition_Internal) MakeFullFieldMask() *LogCondition_Internal_FieldMask {
	return FullLogCondition_Internal_FieldMask()
}

func (o *LogCondition_Internal) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullLogCondition_Internal_FieldMask()
}

func (o *LogCondition_Internal) MakeDiffFieldMask(other *LogCondition_Internal) *LogCondition_Internal_FieldMask {
	if o == nil && other == nil {
		return &LogCondition_Internal_FieldMask{}
	}
	if o == nil || other == nil {
		return FullLogCondition_Internal_FieldMask()
	}

	res := &LogCondition_Internal_FieldMask{}
	if o.GetAlertingLocation() != other.GetAlertingLocation() {
		res.Paths = append(res.Paths, &LogConditionInternal_FieldTerminalPath{selector: LogConditionInternal_FieldPathSelectorAlertingLocation})
	}
	return res
}

func (o *LogCondition_Internal) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*LogCondition_Internal))
}

func (o *LogCondition_Internal) Clone() *LogCondition_Internal {
	if o == nil {
		return nil
	}
	result := &LogCondition_Internal{}
	result.AlertingLocation = o.AlertingLocation
	return result
}

func (o *LogCondition_Internal) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *LogCondition_Internal) Merge(source *LogCondition_Internal) {
	o.AlertingLocation = source.GetAlertingLocation()
}

func (o *LogCondition_Internal) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*LogCondition_Internal))
}

func (o *LogCondition_Spec_Query) GotenObjectExt() {}

func (o *LogCondition_Spec_Query) MakeFullFieldMask() *LogCondition_Spec_Query_FieldMask {
	return FullLogCondition_Spec_Query_FieldMask()
}

func (o *LogCondition_Spec_Query) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullLogCondition_Spec_Query_FieldMask()
}

func (o *LogCondition_Spec_Query) MakeDiffFieldMask(other *LogCondition_Spec_Query) *LogCondition_Spec_Query_FieldMask {
	if o == nil && other == nil {
		return &LogCondition_Spec_Query_FieldMask{}
	}
	if o == nil || other == nil {
		return FullLogCondition_Spec_Query_FieldMask()
	}

	res := &LogCondition_Spec_Query_FieldMask{}
	if o.GetFilter().String() != other.GetFilter().String() {
		res.Paths = append(res.Paths, &LogConditionSpecQuery_FieldTerminalPath{selector: LogConditionSpecQuery_FieldPathSelectorFilter})
	}
	if !proto.Equal(o.GetTrigger(), other.GetTrigger()) {
		res.Paths = append(res.Paths, &LogConditionSpecQuery_FieldTerminalPath{selector: LogConditionSpecQuery_FieldPathSelectorTrigger})
	}
	if !proto.Equal(o.GetMinDuration(), other.GetMinDuration()) {
		res.Paths = append(res.Paths, &LogConditionSpecQuery_FieldTerminalPath{selector: LogConditionSpecQuery_FieldPathSelectorMinDuration})
	}
	return res
}

func (o *LogCondition_Spec_Query) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*LogCondition_Spec_Query))
}

func (o *LogCondition_Spec_Query) Clone() *LogCondition_Spec_Query {
	if o == nil {
		return nil
	}
	result := &LogCondition_Spec_Query{}
	if o.Filter == nil {
		result.Filter = nil
	} else if data, err := o.Filter.ProtoString(); err != nil {
		panic(err)
	} else {
		result.Filter = &logging_log.Filter{}
		if err := result.Filter.ParseProtoString(data); err != nil {
			panic(err)
		}
	}
	result.Trigger = proto.Clone(o.Trigger).(*LogCondition_Spec_Query_TriggerCnd)
	result.MinDuration = proto.Clone(o.MinDuration).(*durationpb.Duration)
	return result
}

func (o *LogCondition_Spec_Query) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *LogCondition_Spec_Query) Merge(source *LogCondition_Spec_Query) {
	if source.GetFilter() != nil {
		if data, err := source.GetFilter().ProtoString(); err != nil {
			panic(err)
		} else {
			o.Filter = &logging_log.Filter{}
			if err := o.Filter.ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	} else {
		o.Filter = nil
	}
	if source.GetTrigger() != nil {
		if o.Trigger == nil {
			o.Trigger = new(LogCondition_Spec_Query_TriggerCnd)
		}
		proto.Merge(o.Trigger, source.GetTrigger())
	}
	if source.GetMinDuration() != nil {
		if o.MinDuration == nil {
			o.MinDuration = new(durationpb.Duration)
		}
		proto.Merge(o.MinDuration, source.GetMinDuration())
	}
}

func (o *LogCondition_Spec_Query) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*LogCondition_Spec_Query))
}

func (o *LogCondition_Spec_Query_LabelTrigger) GotenObjectExt() {}

func (o *LogCondition_Spec_Query_LabelTrigger) MakeFullFieldMask() *LogCondition_Spec_Query_LabelTrigger_FieldMask {
	return FullLogCondition_Spec_Query_LabelTrigger_FieldMask()
}

func (o *LogCondition_Spec_Query_LabelTrigger) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullLogCondition_Spec_Query_LabelTrigger_FieldMask()
}

func (o *LogCondition_Spec_Query_LabelTrigger) MakeDiffFieldMask(other *LogCondition_Spec_Query_LabelTrigger) *LogCondition_Spec_Query_LabelTrigger_FieldMask {
	if o == nil && other == nil {
		return &LogCondition_Spec_Query_LabelTrigger_FieldMask{}
	}
	if o == nil || other == nil {
		return FullLogCondition_Spec_Query_LabelTrigger_FieldMask()
	}

	res := &LogCondition_Spec_Query_LabelTrigger_FieldMask{}
	if o.GetKey() != other.GetKey() {
		res.Paths = append(res.Paths, &LogConditionSpecQueryLabelTrigger_FieldTerminalPath{selector: LogConditionSpecQueryLabelTrigger_FieldPathSelectorKey})
	}

	if len(o.GetValues()) == len(other.GetValues()) {
		for i, lValue := range o.GetValues() {
			rValue := other.GetValues()[i]
			if lValue != rValue {
				res.Paths = append(res.Paths, &LogConditionSpecQueryLabelTrigger_FieldTerminalPath{selector: LogConditionSpecQueryLabelTrigger_FieldPathSelectorValues})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &LogConditionSpecQueryLabelTrigger_FieldTerminalPath{selector: LogConditionSpecQueryLabelTrigger_FieldPathSelectorValues})
	}
	return res
}

func (o *LogCondition_Spec_Query_LabelTrigger) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*LogCondition_Spec_Query_LabelTrigger))
}

func (o *LogCondition_Spec_Query_LabelTrigger) Clone() *LogCondition_Spec_Query_LabelTrigger {
	if o == nil {
		return nil
	}
	result := &LogCondition_Spec_Query_LabelTrigger{}
	result.Key = o.Key
	result.Values = make([]string, len(o.Values))
	for i, sourceValue := range o.Values {
		result.Values[i] = sourceValue
	}
	return result
}

func (o *LogCondition_Spec_Query_LabelTrigger) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *LogCondition_Spec_Query_LabelTrigger) Merge(source *LogCondition_Spec_Query_LabelTrigger) {
	o.Key = source.GetKey()
	for _, sourceValue := range source.GetValues() {
		exists := false
		for _, currentValue := range o.Values {
			if currentValue == sourceValue {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement string
			newDstElement = sourceValue
			o.Values = append(o.Values, newDstElement)
		}
	}

}

func (o *LogCondition_Spec_Query_LabelTrigger) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*LogCondition_Spec_Query_LabelTrigger))
}

func (o *LogCondition_Spec_Query_StringPayloadTrigger) GotenObjectExt() {}

func (o *LogCondition_Spec_Query_StringPayloadTrigger) MakeFullFieldMask() *LogCondition_Spec_Query_StringPayloadTrigger_FieldMask {
	return FullLogCondition_Spec_Query_StringPayloadTrigger_FieldMask()
}

func (o *LogCondition_Spec_Query_StringPayloadTrigger) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullLogCondition_Spec_Query_StringPayloadTrigger_FieldMask()
}

func (o *LogCondition_Spec_Query_StringPayloadTrigger) MakeDiffFieldMask(other *LogCondition_Spec_Query_StringPayloadTrigger) *LogCondition_Spec_Query_StringPayloadTrigger_FieldMask {
	if o == nil && other == nil {
		return &LogCondition_Spec_Query_StringPayloadTrigger_FieldMask{}
	}
	if o == nil || other == nil {
		return FullLogCondition_Spec_Query_StringPayloadTrigger_FieldMask()
	}

	res := &LogCondition_Spec_Query_StringPayloadTrigger_FieldMask{}
	if o.GetObjectSelector() != other.GetObjectSelector() {
		res.Paths = append(res.Paths, &LogConditionSpecQueryStringPayloadTrigger_FieldTerminalPath{selector: LogConditionSpecQueryStringPayloadTrigger_FieldPathSelectorObjectSelector})
	}
	if o.GetRegex() != other.GetRegex() {
		res.Paths = append(res.Paths, &LogConditionSpecQueryStringPayloadTrigger_FieldTerminalPath{selector: LogConditionSpecQueryStringPayloadTrigger_FieldPathSelectorRegex})
	}
	return res
}

func (o *LogCondition_Spec_Query_StringPayloadTrigger) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*LogCondition_Spec_Query_StringPayloadTrigger))
}

func (o *LogCondition_Spec_Query_StringPayloadTrigger) Clone() *LogCondition_Spec_Query_StringPayloadTrigger {
	if o == nil {
		return nil
	}
	result := &LogCondition_Spec_Query_StringPayloadTrigger{}
	result.ObjectSelector = o.ObjectSelector
	result.Regex = o.Regex
	return result
}

func (o *LogCondition_Spec_Query_StringPayloadTrigger) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *LogCondition_Spec_Query_StringPayloadTrigger) Merge(source *LogCondition_Spec_Query_StringPayloadTrigger) {
	o.ObjectSelector = source.GetObjectSelector()
	o.Regex = source.GetRegex()
}

func (o *LogCondition_Spec_Query_StringPayloadTrigger) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*LogCondition_Spec_Query_StringPayloadTrigger))
}

func (o *LogCondition_Spec_Query_CompositeTrigger) GotenObjectExt() {}

func (o *LogCondition_Spec_Query_CompositeTrigger) MakeFullFieldMask() *LogCondition_Spec_Query_CompositeTrigger_FieldMask {
	return FullLogCondition_Spec_Query_CompositeTrigger_FieldMask()
}

func (o *LogCondition_Spec_Query_CompositeTrigger) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullLogCondition_Spec_Query_CompositeTrigger_FieldMask()
}

func (o *LogCondition_Spec_Query_CompositeTrigger) MakeDiffFieldMask(other *LogCondition_Spec_Query_CompositeTrigger) *LogCondition_Spec_Query_CompositeTrigger_FieldMask {
	if o == nil && other == nil {
		return &LogCondition_Spec_Query_CompositeTrigger_FieldMask{}
	}
	if o == nil || other == nil {
		return FullLogCondition_Spec_Query_CompositeTrigger_FieldMask()
	}

	res := &LogCondition_Spec_Query_CompositeTrigger_FieldMask{}

	if len(o.GetTriggers()) == len(other.GetTriggers()) {
		for i, lValue := range o.GetTriggers() {
			rValue := other.GetTriggers()[i]
			if !proto.Equal(lValue, rValue) {
				res.Paths = append(res.Paths, &LogConditionSpecQueryCompositeTrigger_FieldTerminalPath{selector: LogConditionSpecQueryCompositeTrigger_FieldPathSelectorTriggers})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &LogConditionSpecQueryCompositeTrigger_FieldTerminalPath{selector: LogConditionSpecQueryCompositeTrigger_FieldPathSelectorTriggers})
	}
	if o.GetOperator() != other.GetOperator() {
		res.Paths = append(res.Paths, &LogConditionSpecQueryCompositeTrigger_FieldTerminalPath{selector: LogConditionSpecQueryCompositeTrigger_FieldPathSelectorOperator})
	}
	return res
}

func (o *LogCondition_Spec_Query_CompositeTrigger) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*LogCondition_Spec_Query_CompositeTrigger))
}

func (o *LogCondition_Spec_Query_CompositeTrigger) Clone() *LogCondition_Spec_Query_CompositeTrigger {
	if o == nil {
		return nil
	}
	result := &LogCondition_Spec_Query_CompositeTrigger{}
	result.Triggers = make([]*LogCondition_Spec_Query_TriggerCnd, len(o.Triggers))
	for i, sourceValue := range o.Triggers {
		result.Triggers[i] = proto.Clone(sourceValue).(*LogCondition_Spec_Query_TriggerCnd)
	}
	result.Operator = o.Operator
	return result
}

func (o *LogCondition_Spec_Query_CompositeTrigger) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *LogCondition_Spec_Query_CompositeTrigger) Merge(source *LogCondition_Spec_Query_CompositeTrigger) {
	for _, sourceValue := range source.GetTriggers() {
		exists := false
		for _, currentValue := range o.Triggers {
			if proto.Equal(sourceValue, currentValue) {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement *LogCondition_Spec_Query_TriggerCnd
			if sourceValue != nil {
				newDstElement = new(LogCondition_Spec_Query_TriggerCnd)
				proto.Merge(newDstElement, sourceValue)
			}
			o.Triggers = append(o.Triggers, newDstElement)
		}
	}

	o.Operator = source.GetOperator()
}

func (o *LogCondition_Spec_Query_CompositeTrigger) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*LogCondition_Spec_Query_CompositeTrigger))
}
