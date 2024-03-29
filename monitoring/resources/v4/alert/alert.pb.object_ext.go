// Code generated by protoc-gen-goten-object
// File: edgelq/monitoring/proto/v4/alert.proto
// DO NOT EDIT!!!

package alert

import (
	"fmt"
	"sort"

	"google.golang.org/protobuf/proto"
	googlefieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"

	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
)

// proto imports
import (
	alerting_condition "github.com/cloudwan/edgelq-sdk/monitoring/resources/v4/alerting_condition"
	common "github.com/cloudwan/edgelq-sdk/monitoring/resources/v4/common"
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
	_ = &alerting_condition.AlertingCondition{}
	_ = &common.LabelDescriptor{}
	_ = &meta.Meta{}
)

func (o *Alert) GotenObjectExt() {}

func (o *Alert) MakeFullFieldMask() *Alert_FieldMask {
	return FullAlert_FieldMask()
}

func (o *Alert) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullAlert_FieldMask()
}

func (o *Alert) MakeDiffFieldMask(other *Alert) *Alert_FieldMask {
	if o == nil && other == nil {
		return &Alert_FieldMask{}
	}
	if o == nil || other == nil {
		return FullAlert_FieldMask()
	}

	res := &Alert_FieldMask{}
	if o.GetName().String() != other.GetName().String() {
		res.Paths = append(res.Paths, &Alert_FieldTerminalPath{selector: Alert_FieldPathSelectorName})
	}
	{
		subMask := o.GetMetadata().MakeDiffFieldMask(other.GetMetadata())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &Alert_FieldTerminalPath{selector: Alert_FieldPathSelectorMetadata})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &Alert_FieldSubPath{selector: Alert_FieldPathSelectorMetadata, subPath: subpath})
			}
		}
	}
	if o.GetDisplayName() != other.GetDisplayName() {
		res.Paths = append(res.Paths, &Alert_FieldTerminalPath{selector: Alert_FieldPathSelectorDisplayName})
	}
	{
		subMask := o.GetInfo().MakeDiffFieldMask(other.GetInfo())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &Alert_FieldTerminalPath{selector: Alert_FieldPathSelectorInfo})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &Alert_FieldSubPath{selector: Alert_FieldPathSelectorInfo, subPath: subpath})
			}
		}
	}
	{
		subMask := o.GetState().MakeDiffFieldMask(other.GetState())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &Alert_FieldTerminalPath{selector: Alert_FieldPathSelectorState})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &Alert_FieldSubPath{selector: Alert_FieldPathSelectorState, subPath: subpath})
			}
		}
	}
	return res
}

func (o *Alert) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*Alert))
}

func (o *Alert) Clone() *Alert {
	if o == nil {
		return nil
	}
	result := &Alert{}
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
	result.Info = o.Info.Clone()
	result.State = o.State.Clone()
	return result
}

func (o *Alert) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *Alert) Merge(source *Alert) {
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
	if source.GetInfo() != nil {
		if o.Info == nil {
			o.Info = new(Alert_Info)
		}
		o.Info.Merge(source.GetInfo())
	}
	if source.GetState() != nil {
		if o.State == nil {
			o.State = new(Alert_State)
		}
		o.State.Merge(source.GetState())
	}
}

func (o *Alert) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*Alert))
}

func (o *Alert_Info) GotenObjectExt() {}

func (o *Alert_Info) MakeFullFieldMask() *Alert_Info_FieldMask {
	return FullAlert_Info_FieldMask()
}

func (o *Alert_Info) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullAlert_Info_FieldMask()
}

func (o *Alert_Info) MakeDiffFieldMask(other *Alert_Info) *Alert_Info_FieldMask {
	if o == nil && other == nil {
		return &Alert_Info_FieldMask{}
	}
	if o == nil || other == nil {
		return FullAlert_Info_FieldMask()
	}

	res := &Alert_Info_FieldMask{}
	{
		subMask := o.GetTimeSerie().MakeDiffFieldMask(other.GetTimeSerie())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &AlertInfo_FieldTerminalPath{selector: AlertInfo_FieldPathSelectorTimeSerie})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &AlertInfo_FieldSubPath{selector: AlertInfo_FieldPathSelectorTimeSerie, subPath: subpath})
			}
		}
	}
	{
		subMask := o.GetObservedValues().MakeDiffFieldMask(other.GetObservedValues())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &AlertInfo_FieldTerminalPath{selector: AlertInfo_FieldPathSelectorObservedValues})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &AlertInfo_FieldSubPath{selector: AlertInfo_FieldPathSelectorObservedValues, subPath: subpath})
			}
		}
	}
	return res
}

func (o *Alert_Info) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*Alert_Info))
}

func (o *Alert_Info) Clone() *Alert_Info {
	if o == nil {
		return nil
	}
	result := &Alert_Info{}
	result.TimeSerie = o.TimeSerie.Clone()
	result.ObservedValues = o.ObservedValues.Clone()
	return result
}

func (o *Alert_Info) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *Alert_Info) Merge(source *Alert_Info) {
	if source.GetTimeSerie() != nil {
		if o.TimeSerie == nil {
			o.TimeSerie = new(Alert_Info_TimeSerie)
		}
		o.TimeSerie.Merge(source.GetTimeSerie())
	}
	if source.GetObservedValues() != nil {
		if o.ObservedValues == nil {
			o.ObservedValues = new(Alert_Info_ObservedValues)
		}
		o.ObservedValues.Merge(source.GetObservedValues())
	}
}

func (o *Alert_Info) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*Alert_Info))
}

func (o *Alert_State) GotenObjectExt() {}

func (o *Alert_State) MakeFullFieldMask() *Alert_State_FieldMask {
	return FullAlert_State_FieldMask()
}

func (o *Alert_State) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullAlert_State_FieldMask()
}

func (o *Alert_State) MakeDiffFieldMask(other *Alert_State) *Alert_State_FieldMask {
	if o == nil && other == nil {
		return &Alert_State_FieldMask{}
	}
	if o == nil || other == nil {
		return FullAlert_State_FieldMask()
	}

	res := &Alert_State_FieldMask{}
	if o.GetIsFiring() != other.GetIsFiring() {
		res.Paths = append(res.Paths, &AlertState_FieldTerminalPath{selector: AlertState_FieldPathSelectorIsFiring})
	}
	if o.GetIsAcknowledged() != other.GetIsAcknowledged() {
		res.Paths = append(res.Paths, &AlertState_FieldTerminalPath{selector: AlertState_FieldPathSelectorIsAcknowledged})
	}
	if o.GetIsSilenced() != other.GetIsSilenced() {
		res.Paths = append(res.Paths, &AlertState_FieldTerminalPath{selector: AlertState_FieldPathSelectorIsSilenced})
	}
	{
		subMask := o.GetLifetime().MakeDiffFieldMask(other.GetLifetime())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &AlertState_FieldTerminalPath{selector: AlertState_FieldPathSelectorLifetime})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &AlertState_FieldSubPath{selector: AlertState_FieldPathSelectorLifetime, subPath: subpath})
			}
		}
	}
	if o.GetNeedsNotification() != other.GetNeedsNotification() {
		res.Paths = append(res.Paths, &AlertState_FieldTerminalPath{selector: AlertState_FieldPathSelectorNeedsNotification})
	}
	if o.GetNotificationCreated() != other.GetNotificationCreated() {
		res.Paths = append(res.Paths, &AlertState_FieldTerminalPath{selector: AlertState_FieldPathSelectorNotificationCreated})
	}
	if o.GetLifecycleCompleted() != other.GetLifecycleCompleted() {
		res.Paths = append(res.Paths, &AlertState_FieldTerminalPath{selector: AlertState_FieldPathSelectorLifecycleCompleted})
	}
	return res
}

func (o *Alert_State) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*Alert_State))
}

func (o *Alert_State) Clone() *Alert_State {
	if o == nil {
		return nil
	}
	result := &Alert_State{}
	result.IsFiring = o.IsFiring
	result.IsAcknowledged = o.IsAcknowledged
	result.IsSilenced = o.IsSilenced
	result.Lifetime = o.Lifetime.Clone()
	result.NeedsNotification = o.NeedsNotification
	result.NotificationCreated = o.NotificationCreated
	result.LifecycleCompleted = o.LifecycleCompleted
	return result
}

func (o *Alert_State) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *Alert_State) Merge(source *Alert_State) {
	o.IsFiring = source.GetIsFiring()
	o.IsAcknowledged = source.GetIsAcknowledged()
	o.IsSilenced = source.GetIsSilenced()
	if source.GetLifetime() != nil {
		if o.Lifetime == nil {
			o.Lifetime = new(common.TimeRange)
		}
		o.Lifetime.Merge(source.GetLifetime())
	}
	o.NeedsNotification = source.GetNeedsNotification()
	o.NotificationCreated = source.GetNotificationCreated()
	o.LifecycleCompleted = source.GetLifecycleCompleted()
}

func (o *Alert_State) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*Alert_State))
}

func (o *Alert_Info_TimeSerie) GotenObjectExt() {}

func (o *Alert_Info_TimeSerie) MakeFullFieldMask() *Alert_Info_TimeSerie_FieldMask {
	return FullAlert_Info_TimeSerie_FieldMask()
}

func (o *Alert_Info_TimeSerie) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullAlert_Info_TimeSerie_FieldMask()
}

func (o *Alert_Info_TimeSerie) MakeDiffFieldMask(other *Alert_Info_TimeSerie) *Alert_Info_TimeSerie_FieldMask {
	if o == nil && other == nil {
		return &Alert_Info_TimeSerie_FieldMask{}
	}
	if o == nil || other == nil {
		return FullAlert_Info_TimeSerie_FieldMask()
	}

	res := &Alert_Info_TimeSerie_FieldMask{}
	if string(o.GetKey()) != string(other.GetKey()) {
		res.Paths = append(res.Paths, &AlertInfoTimeSerie_FieldTerminalPath{selector: AlertInfoTimeSerie_FieldPathSelectorKey})
	}
	{
		subMask := o.GetMetric().MakeDiffFieldMask(other.GetMetric())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &AlertInfoTimeSerie_FieldTerminalPath{selector: AlertInfoTimeSerie_FieldPathSelectorMetric})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &AlertInfoTimeSerie_FieldSubPath{selector: AlertInfoTimeSerie_FieldPathSelectorMetric, subPath: subpath})
			}
		}
	}
	{
		subMask := o.GetMonitoredResource().MakeDiffFieldMask(other.GetMonitoredResource())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &AlertInfoTimeSerie_FieldTerminalPath{selector: AlertInfoTimeSerie_FieldPathSelectorMonitoredResource})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &AlertInfoTimeSerie_FieldSubPath{selector: AlertInfoTimeSerie_FieldPathSelectorMonitoredResource, subPath: subpath})
			}
		}
	}
	return res
}

func (o *Alert_Info_TimeSerie) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*Alert_Info_TimeSerie))
}

func (o *Alert_Info_TimeSerie) Clone() *Alert_Info_TimeSerie {
	if o == nil {
		return nil
	}
	result := &Alert_Info_TimeSerie{}
	result.Key = make([]byte, len(o.Key))
	for i, bt := range o.Key {
		result.Key[i] = bt
	}
	result.Metric = o.Metric.Clone()
	result.MonitoredResource = o.MonitoredResource.Clone()
	return result
}

func (o *Alert_Info_TimeSerie) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *Alert_Info_TimeSerie) Merge(source *Alert_Info_TimeSerie) {
	o.Key = make([]byte, len(source.GetKey()))
	for i, bt := range source.GetKey() {
		o.Key[i] = bt
	}
	if source.GetMetric() != nil {
		if o.Metric == nil {
			o.Metric = new(common.Metric)
		}
		o.Metric.Merge(source.GetMetric())
	}
	if source.GetMonitoredResource() != nil {
		if o.MonitoredResource == nil {
			o.MonitoredResource = new(common.MonitoredResource)
		}
		o.MonitoredResource.Merge(source.GetMonitoredResource())
	}
}

func (o *Alert_Info_TimeSerie) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*Alert_Info_TimeSerie))
}

func (o *Alert_Info_ObservedValues) GotenObjectExt() {}

func (o *Alert_Info_ObservedValues) MakeFullFieldMask() *Alert_Info_ObservedValues_FieldMask {
	return FullAlert_Info_ObservedValues_FieldMask()
}

func (o *Alert_Info_ObservedValues) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullAlert_Info_ObservedValues_FieldMask()
}

func (o *Alert_Info_ObservedValues) MakeDiffFieldMask(other *Alert_Info_ObservedValues) *Alert_Info_ObservedValues_FieldMask {
	if o == nil && other == nil {
		return &Alert_Info_ObservedValues_FieldMask{}
	}
	if o == nil || other == nil {
		return FullAlert_Info_ObservedValues_FieldMask()
	}

	res := &Alert_Info_ObservedValues_FieldMask{}
	if o.GetExampleValue() != other.GetExampleValue() {
		res.Paths = append(res.Paths, &AlertInfoObservedValues_FieldTerminalPath{selector: AlertInfoObservedValues_FieldPathSelectorExampleValue})
	}

	if len(o.GetPerMetric()) == len(other.GetPerMetric()) {
		for i, lValue := range o.GetPerMetric() {
			rValue := other.GetPerMetric()[i]
			if lValue != rValue {
				res.Paths = append(res.Paths, &AlertInfoObservedValues_FieldTerminalPath{selector: AlertInfoObservedValues_FieldPathSelectorPerMetric})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &AlertInfoObservedValues_FieldTerminalPath{selector: AlertInfoObservedValues_FieldPathSelectorPerMetric})
	}
	return res
}

func (o *Alert_Info_ObservedValues) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*Alert_Info_ObservedValues))
}

func (o *Alert_Info_ObservedValues) Clone() *Alert_Info_ObservedValues {
	if o == nil {
		return nil
	}
	result := &Alert_Info_ObservedValues{}
	result.ExampleValue = o.ExampleValue
	result.PerMetric = map[string]float64{}
	for key, sourceValue := range o.PerMetric {
		result.PerMetric[key] = sourceValue
	}
	return result
}

func (o *Alert_Info_ObservedValues) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *Alert_Info_ObservedValues) Merge(source *Alert_Info_ObservedValues) {
	o.ExampleValue = source.GetExampleValue()
	if source.GetPerMetric() != nil {
		if o.PerMetric == nil {
			o.PerMetric = make(map[string]float64, len(source.GetPerMetric()))
		}
		for key, sourceValue := range source.GetPerMetric() {
			o.PerMetric[key] = sourceValue
		}
	}
}

func (o *Alert_Info_ObservedValues) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*Alert_Info_ObservedValues))
}

func (o *Alert_State_Threshold) GotenObjectExt() {}

func (o *Alert_State_Threshold) MakeFullFieldMask() *Alert_State_Threshold_FieldMask {
	return FullAlert_State_Threshold_FieldMask()
}

func (o *Alert_State_Threshold) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullAlert_State_Threshold_FieldMask()
}

func (o *Alert_State_Threshold) MakeDiffFieldMask(other *Alert_State_Threshold) *Alert_State_Threshold_FieldMask {
	if o == nil && other == nil {
		return &Alert_State_Threshold_FieldMask{}
	}
	if o == nil || other == nil {
		return FullAlert_State_Threshold_FieldMask()
	}

	res := &Alert_State_Threshold_FieldMask{}
	if o.GetObservedValue() != other.GetObservedValue() {
		res.Paths = append(res.Paths, &AlertStateThreshold_FieldTerminalPath{selector: AlertStateThreshold_FieldPathSelectorObservedValue})
	}
	return res
}

func (o *Alert_State_Threshold) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*Alert_State_Threshold))
}

func (o *Alert_State_Threshold) Clone() *Alert_State_Threshold {
	if o == nil {
		return nil
	}
	result := &Alert_State_Threshold{}
	result.ObservedValue = o.ObservedValue
	return result
}

func (o *Alert_State_Threshold) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *Alert_State_Threshold) Merge(source *Alert_State_Threshold) {
	o.ObservedValue = source.GetObservedValue()
}

func (o *Alert_State_Threshold) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*Alert_State_Threshold))
}

func (o *Alert_State_CombineThreshold) GotenObjectExt() {}

func (o *Alert_State_CombineThreshold) MakeFullFieldMask() *Alert_State_CombineThreshold_FieldMask {
	return FullAlert_State_CombineThreshold_FieldMask()
}

func (o *Alert_State_CombineThreshold) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullAlert_State_CombineThreshold_FieldMask()
}

func (o *Alert_State_CombineThreshold) MakeDiffFieldMask(other *Alert_State_CombineThreshold) *Alert_State_CombineThreshold_FieldMask {
	if o == nil && other == nil {
		return &Alert_State_CombineThreshold_FieldMask{}
	}
	if o == nil || other == nil {
		return FullAlert_State_CombineThreshold_FieldMask()
	}

	res := &Alert_State_CombineThreshold_FieldMask{}
	return res
}

func (o *Alert_State_CombineThreshold) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*Alert_State_CombineThreshold))
}

func (o *Alert_State_CombineThreshold) Clone() *Alert_State_CombineThreshold {
	if o == nil {
		return nil
	}
	result := &Alert_State_CombineThreshold{}
	return result
}

func (o *Alert_State_CombineThreshold) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *Alert_State_CombineThreshold) Merge(source *Alert_State_CombineThreshold) {
}

func (o *Alert_State_CombineThreshold) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*Alert_State_CombineThreshold))
}

func (o *Alert_State_CombineThreshold_PerMetric) GotenObjectExt() {}

func (o *Alert_State_CombineThreshold_PerMetric) MakeFullFieldMask() *Alert_State_CombineThreshold_PerMetric_FieldMask {
	return FullAlert_State_CombineThreshold_PerMetric_FieldMask()
}

func (o *Alert_State_CombineThreshold_PerMetric) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullAlert_State_CombineThreshold_PerMetric_FieldMask()
}

func (o *Alert_State_CombineThreshold_PerMetric) MakeDiffFieldMask(other *Alert_State_CombineThreshold_PerMetric) *Alert_State_CombineThreshold_PerMetric_FieldMask {
	if o == nil && other == nil {
		return &Alert_State_CombineThreshold_PerMetric_FieldMask{}
	}
	if o == nil || other == nil {
		return FullAlert_State_CombineThreshold_PerMetric_FieldMask()
	}

	res := &Alert_State_CombineThreshold_PerMetric_FieldMask{}

	if len(o.GetObservedValues()) == len(other.GetObservedValues()) {
		for i, lValue := range o.GetObservedValues() {
			rValue := other.GetObservedValues()[i]
			if lValue != rValue {
				res.Paths = append(res.Paths, &AlertStateCombineThresholdPerMetric_FieldTerminalPath{selector: AlertStateCombineThresholdPerMetric_FieldPathSelectorObservedValues})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &AlertStateCombineThresholdPerMetric_FieldTerminalPath{selector: AlertStateCombineThresholdPerMetric_FieldPathSelectorObservedValues})
	}
	return res
}

func (o *Alert_State_CombineThreshold_PerMetric) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*Alert_State_CombineThreshold_PerMetric))
}

func (o *Alert_State_CombineThreshold_PerMetric) Clone() *Alert_State_CombineThreshold_PerMetric {
	if o == nil {
		return nil
	}
	result := &Alert_State_CombineThreshold_PerMetric{}
	result.ObservedValues = map[string]float64{}
	for key, sourceValue := range o.ObservedValues {
		result.ObservedValues[key] = sourceValue
	}
	return result
}

func (o *Alert_State_CombineThreshold_PerMetric) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *Alert_State_CombineThreshold_PerMetric) Merge(source *Alert_State_CombineThreshold_PerMetric) {
	if source.GetObservedValues() != nil {
		if o.ObservedValues == nil {
			o.ObservedValues = make(map[string]float64, len(source.GetObservedValues()))
		}
		for key, sourceValue := range source.GetObservedValues() {
			o.ObservedValues[key] = sourceValue
		}
	}
}

func (o *Alert_State_CombineThreshold_PerMetric) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*Alert_State_CombineThreshold_PerMetric))
}
