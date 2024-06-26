// Code generated by protoc-gen-goten-object
// File: edgelq/monitoring/proto/v4/alerting_policy.proto
// DO NOT EDIT!!!

package alerting_policy

import (
	"fmt"
	"sort"

	"google.golang.org/protobuf/proto"
	googlefieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"

	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
)

// proto imports
import (
	notification_channel "github.com/cloudwan/edgelq-sdk/monitoring/resources/v4/notification_channel"
	project "github.com/cloudwan/edgelq-sdk/monitoring/resources/v4/project"
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
	_ = &notification_channel.NotificationChannel{}
	_ = &project.Project{}
	_ = &meta.Meta{}
)

func (o *AlertingPolicy) GotenObjectExt() {}

func (o *AlertingPolicy) MakeFullFieldMask() *AlertingPolicy_FieldMask {
	return FullAlertingPolicy_FieldMask()
}

func (o *AlertingPolicy) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullAlertingPolicy_FieldMask()
}

func (o *AlertingPolicy) MakeDiffFieldMask(other *AlertingPolicy) *AlertingPolicy_FieldMask {
	if o == nil && other == nil {
		return &AlertingPolicy_FieldMask{}
	}
	if o == nil || other == nil {
		return FullAlertingPolicy_FieldMask()
	}

	res := &AlertingPolicy_FieldMask{}
	if o.GetName().String() != other.GetName().String() {
		res.Paths = append(res.Paths, &AlertingPolicy_FieldTerminalPath{selector: AlertingPolicy_FieldPathSelectorName})
	}
	{
		subMask := o.GetMetadata().MakeDiffFieldMask(other.GetMetadata())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &AlertingPolicy_FieldTerminalPath{selector: AlertingPolicy_FieldPathSelectorMetadata})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &AlertingPolicy_FieldSubPath{selector: AlertingPolicy_FieldPathSelectorMetadata, subPath: subpath})
			}
		}
	}
	if o.GetDisplayName() != other.GetDisplayName() {
		res.Paths = append(res.Paths, &AlertingPolicy_FieldTerminalPath{selector: AlertingPolicy_FieldPathSelectorDisplayName})
	}
	if o.GetDescription() != other.GetDescription() {
		res.Paths = append(res.Paths, &AlertingPolicy_FieldTerminalPath{selector: AlertingPolicy_FieldPathSelectorDescription})
	}
	{
		subMask := o.GetDocumentation().MakeDiffFieldMask(other.GetDocumentation())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &AlertingPolicy_FieldTerminalPath{selector: AlertingPolicy_FieldPathSelectorDocumentation})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &AlertingPolicy_FieldSubPath{selector: AlertingPolicy_FieldPathSelectorDocumentation, subPath: subpath})
			}
		}
	}
	{
		subMask := o.GetSpec().MakeDiffFieldMask(other.GetSpec())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &AlertingPolicy_FieldTerminalPath{selector: AlertingPolicy_FieldPathSelectorSpec})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &AlertingPolicy_FieldSubPath{selector: AlertingPolicy_FieldPathSelectorSpec, subPath: subpath})
			}
		}
	}
	{
		subMask := o.GetState().MakeDiffFieldMask(other.GetState())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &AlertingPolicy_FieldTerminalPath{selector: AlertingPolicy_FieldPathSelectorState})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &AlertingPolicy_FieldSubPath{selector: AlertingPolicy_FieldPathSelectorState, subPath: subpath})
			}
		}
	}
	return res
}

func (o *AlertingPolicy) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*AlertingPolicy))
}

func (o *AlertingPolicy) Clone() *AlertingPolicy {
	if o == nil {
		return nil
	}
	result := &AlertingPolicy{}
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
	result.Documentation = o.Documentation.Clone()
	result.Spec = o.Spec.Clone()
	result.State = o.State.Clone()
	return result
}

func (o *AlertingPolicy) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *AlertingPolicy) Merge(source *AlertingPolicy) {
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
	if source.GetDocumentation() != nil {
		if o.Documentation == nil {
			o.Documentation = new(AlertingPolicy_Documentation)
		}
		o.Documentation.Merge(source.GetDocumentation())
	}
	if source.GetSpec() != nil {
		if o.Spec == nil {
			o.Spec = new(AlertingPolicy_Spec)
		}
		o.Spec.Merge(source.GetSpec())
	}
	if source.GetState() != nil {
		if o.State == nil {
			o.State = new(AlertingPolicy_State)
		}
		o.State.Merge(source.GetState())
	}
}

func (o *AlertingPolicy) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*AlertingPolicy))
}

func (o *AlertingPolicy_Documentation) GotenObjectExt() {}

func (o *AlertingPolicy_Documentation) MakeFullFieldMask() *AlertingPolicy_Documentation_FieldMask {
	return FullAlertingPolicy_Documentation_FieldMask()
}

func (o *AlertingPolicy_Documentation) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullAlertingPolicy_Documentation_FieldMask()
}

func (o *AlertingPolicy_Documentation) MakeDiffFieldMask(other *AlertingPolicy_Documentation) *AlertingPolicy_Documentation_FieldMask {
	if o == nil && other == nil {
		return &AlertingPolicy_Documentation_FieldMask{}
	}
	if o == nil || other == nil {
		return FullAlertingPolicy_Documentation_FieldMask()
	}

	res := &AlertingPolicy_Documentation_FieldMask{}
	if o.GetContent() != other.GetContent() {
		res.Paths = append(res.Paths, &AlertingPolicyDocumentation_FieldTerminalPath{selector: AlertingPolicyDocumentation_FieldPathSelectorContent})
	}
	if o.GetMimeType() != other.GetMimeType() {
		res.Paths = append(res.Paths, &AlertingPolicyDocumentation_FieldTerminalPath{selector: AlertingPolicyDocumentation_FieldPathSelectorMimeType})
	}
	return res
}

func (o *AlertingPolicy_Documentation) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*AlertingPolicy_Documentation))
}

func (o *AlertingPolicy_Documentation) Clone() *AlertingPolicy_Documentation {
	if o == nil {
		return nil
	}
	result := &AlertingPolicy_Documentation{}
	result.Content = o.Content
	result.MimeType = o.MimeType
	return result
}

func (o *AlertingPolicy_Documentation) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *AlertingPolicy_Documentation) Merge(source *AlertingPolicy_Documentation) {
	o.Content = source.GetContent()
	o.MimeType = source.GetMimeType()
}

func (o *AlertingPolicy_Documentation) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*AlertingPolicy_Documentation))
}

func (o *AlertingPolicy_Spec) GotenObjectExt() {}

func (o *AlertingPolicy_Spec) MakeFullFieldMask() *AlertingPolicy_Spec_FieldMask {
	return FullAlertingPolicy_Spec_FieldMask()
}

func (o *AlertingPolicy_Spec) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullAlertingPolicy_Spec_FieldMask()
}

func (o *AlertingPolicy_Spec) MakeDiffFieldMask(other *AlertingPolicy_Spec) *AlertingPolicy_Spec_FieldMask {
	if o == nil && other == nil {
		return &AlertingPolicy_Spec_FieldMask{}
	}
	if o == nil || other == nil {
		return FullAlertingPolicy_Spec_FieldMask()
	}

	res := &AlertingPolicy_Spec_FieldMask{}
	if o.GetEnabled() != other.GetEnabled() {
		res.Paths = append(res.Paths, &AlertingPolicySpec_FieldTerminalPath{selector: AlertingPolicySpec_FieldPathSelectorEnabled})
	}
	if o.GetConditionCombiner() != other.GetConditionCombiner() {
		res.Paths = append(res.Paths, &AlertingPolicySpec_FieldTerminalPath{selector: AlertingPolicySpec_FieldPathSelectorConditionCombiner})
	}
	{
		subMask := o.GetNotification().MakeDiffFieldMask(other.GetNotification())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &AlertingPolicySpec_FieldTerminalPath{selector: AlertingPolicySpec_FieldPathSelectorNotification})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &AlertingPolicySpec_FieldSubPath{selector: AlertingPolicySpec_FieldPathSelectorNotification, subPath: subpath})
			}
		}
	}
	return res
}

func (o *AlertingPolicy_Spec) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*AlertingPolicy_Spec))
}

func (o *AlertingPolicy_Spec) Clone() *AlertingPolicy_Spec {
	if o == nil {
		return nil
	}
	result := &AlertingPolicy_Spec{}
	result.Enabled = o.Enabled
	result.ConditionCombiner = o.ConditionCombiner
	result.Notification = o.Notification.Clone()
	return result
}

func (o *AlertingPolicy_Spec) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *AlertingPolicy_Spec) Merge(source *AlertingPolicy_Spec) {
	o.Enabled = source.GetEnabled()
	o.ConditionCombiner = source.GetConditionCombiner()
	if source.GetNotification() != nil {
		if o.Notification == nil {
			o.Notification = new(AlertingPolicy_Spec_Notification)
		}
		o.Notification.Merge(source.GetNotification())
	}
}

func (o *AlertingPolicy_Spec) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*AlertingPolicy_Spec))
}

func (o *AlertingPolicy_State) GotenObjectExt() {}

func (o *AlertingPolicy_State) MakeFullFieldMask() *AlertingPolicy_State_FieldMask {
	return FullAlertingPolicy_State_FieldMask()
}

func (o *AlertingPolicy_State) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullAlertingPolicy_State_FieldMask()
}

func (o *AlertingPolicy_State) MakeDiffFieldMask(other *AlertingPolicy_State) *AlertingPolicy_State_FieldMask {
	if o == nil && other == nil {
		return &AlertingPolicy_State_FieldMask{}
	}
	if o == nil || other == nil {
		return FullAlertingPolicy_State_FieldMask()
	}

	res := &AlertingPolicy_State_FieldMask{}
	if o.GetActiveAlertsCount() != other.GetActiveAlertsCount() {
		res.Paths = append(res.Paths, &AlertingPolicyState_FieldTerminalPath{selector: AlertingPolicyState_FieldPathSelectorActiveAlertsCount})
	}
	return res
}

func (o *AlertingPolicy_State) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*AlertingPolicy_State))
}

func (o *AlertingPolicy_State) Clone() *AlertingPolicy_State {
	if o == nil {
		return nil
	}
	result := &AlertingPolicy_State{}
	result.ActiveAlertsCount = o.ActiveAlertsCount
	return result
}

func (o *AlertingPolicy_State) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *AlertingPolicy_State) Merge(source *AlertingPolicy_State) {
	o.ActiveAlertsCount = source.GetActiveAlertsCount()
}

func (o *AlertingPolicy_State) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*AlertingPolicy_State))
}

func (o *AlertingPolicy_Spec_Notification) GotenObjectExt() {}

func (o *AlertingPolicy_Spec_Notification) MakeFullFieldMask() *AlertingPolicy_Spec_Notification_FieldMask {
	return FullAlertingPolicy_Spec_Notification_FieldMask()
}

func (o *AlertingPolicy_Spec_Notification) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullAlertingPolicy_Spec_Notification_FieldMask()
}

func (o *AlertingPolicy_Spec_Notification) MakeDiffFieldMask(other *AlertingPolicy_Spec_Notification) *AlertingPolicy_Spec_Notification_FieldMask {
	if o == nil && other == nil {
		return &AlertingPolicy_Spec_Notification_FieldMask{}
	}
	if o == nil || other == nil {
		return FullAlertingPolicy_Spec_Notification_FieldMask()
	}

	res := &AlertingPolicy_Spec_Notification_FieldMask{}
	if o.GetEnabled() != other.GetEnabled() {
		res.Paths = append(res.Paths, &AlertingPolicySpecNotification_FieldTerminalPath{selector: AlertingPolicySpecNotification_FieldPathSelectorEnabled})
	}

	if len(o.GetChannels()) == len(other.GetChannels()) {
		for i, lValue := range o.GetChannels() {
			rValue := other.GetChannels()[i]
			if lValue.String() != rValue.String() {
				res.Paths = append(res.Paths, &AlertingPolicySpecNotification_FieldTerminalPath{selector: AlertingPolicySpecNotification_FieldPathSelectorChannels})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &AlertingPolicySpecNotification_FieldTerminalPath{selector: AlertingPolicySpecNotification_FieldPathSelectorChannels})
	}
	return res
}

func (o *AlertingPolicy_Spec_Notification) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*AlertingPolicy_Spec_Notification))
}

func (o *AlertingPolicy_Spec_Notification) Clone() *AlertingPolicy_Spec_Notification {
	if o == nil {
		return nil
	}
	result := &AlertingPolicy_Spec_Notification{}
	result.Enabled = o.Enabled
	result.Channels = make([]*notification_channel.Reference, len(o.Channels))
	for i, sourceValue := range o.Channels {
		if sourceValue == nil {
			result.Channels[i] = nil
		} else if data, err := sourceValue.ProtoString(); err != nil {
			panic(err)
		} else {
			result.Channels[i] = &notification_channel.Reference{}
			if err := result.Channels[i].ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	}
	return result
}

func (o *AlertingPolicy_Spec_Notification) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *AlertingPolicy_Spec_Notification) Merge(source *AlertingPolicy_Spec_Notification) {
	o.Enabled = source.GetEnabled()
	for _, sourceValue := range source.GetChannels() {
		exists := false
		for _, currentValue := range o.Channels {
			leftProtoStr, _ := currentValue.ProtoString()
			rightProtoStr, _ := sourceValue.ProtoString()
			if leftProtoStr == rightProtoStr {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement *notification_channel.Reference
			if sourceValue != nil {
				if data, err := sourceValue.ProtoString(); err != nil {
					panic(err)
				} else {
					newDstElement = &notification_channel.Reference{}
					if err := newDstElement.ParseProtoString(data); err != nil {
						panic(err)
					}
				}
			}
			o.Channels = append(o.Channels, newDstElement)
		}
	}

}

func (o *AlertingPolicy_Spec_Notification) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*AlertingPolicy_Spec_Notification))
}
