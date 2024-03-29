// Code generated by protoc-gen-goten-object
// File: edgelq/monitoring/proto/v3/notification.proto
// DO NOT EDIT!!!

package notification

import (
	"fmt"
	"sort"

	"google.golang.org/protobuf/proto"
	googlefieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"

	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
)

// proto imports
import (
	alert "github.com/cloudwan/edgelq-sdk/monitoring/resources/v3/alert"
	alerting_condition "github.com/cloudwan/edgelq-sdk/monitoring/resources/v3/alerting_condition"
	alerting_policy "github.com/cloudwan/edgelq-sdk/monitoring/resources/v3/alerting_policy"
	common "github.com/cloudwan/edgelq-sdk/monitoring/resources/v3/common"
	notification_channel "github.com/cloudwan/edgelq-sdk/monitoring/resources/v3/notification_channel"
	project "github.com/cloudwan/edgelq-sdk/monitoring/resources/v3/project"
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
	_ = &alert.Alert{}
	_ = &alerting_condition.AlertingCondition{}
	_ = &alerting_policy.AlertingPolicy{}
	_ = &common.LabelDescriptor{}
	_ = &notification_channel.NotificationChannel{}
	_ = &project.Project{}
	_ = &meta.Meta{}
)

func (o *Notification) GotenObjectExt() {}

func (o *Notification) MakeFullFieldMask() *Notification_FieldMask {
	return FullNotification_FieldMask()
}

func (o *Notification) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullNotification_FieldMask()
}

func (o *Notification) MakeDiffFieldMask(other *Notification) *Notification_FieldMask {
	if o == nil && other == nil {
		return &Notification_FieldMask{}
	}
	if o == nil || other == nil {
		return FullNotification_FieldMask()
	}

	res := &Notification_FieldMask{}
	if o.GetName().String() != other.GetName().String() {
		res.Paths = append(res.Paths, &Notification_FieldTerminalPath{selector: Notification_FieldPathSelectorName})
	}
	{
		subMask := o.GetMetadata().MakeDiffFieldMask(other.GetMetadata())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &Notification_FieldTerminalPath{selector: Notification_FieldPathSelectorMetadata})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &Notification_FieldSubPath{selector: Notification_FieldPathSelectorMetadata, subPath: subpath})
			}
		}
	}
	if o.GetAlertingPolicy().String() != other.GetAlertingPolicy().String() {
		res.Paths = append(res.Paths, &Notification_FieldTerminalPath{selector: Notification_FieldPathSelectorAlertingPolicy})
	}

	if len(o.GetAlerts()) == len(other.GetAlerts()) {
		for i, lValue := range o.GetAlerts() {
			rValue := other.GetAlerts()[i]
			if lValue.String() != rValue.String() {
				res.Paths = append(res.Paths, &Notification_FieldTerminalPath{selector: Notification_FieldPathSelectorAlerts})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &Notification_FieldTerminalPath{selector: Notification_FieldPathSelectorAlerts})
	}
	{
		subMask := o.GetState().MakeDiffFieldMask(other.GetState())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &Notification_FieldTerminalPath{selector: Notification_FieldPathSelectorState})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &Notification_FieldSubPath{selector: Notification_FieldPathSelectorState, subPath: subpath})
			}
		}
	}
	return res
}

func (o *Notification) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*Notification))
}

func (o *Notification) Clone() *Notification {
	if o == nil {
		return nil
	}
	result := &Notification{}
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
	if o.AlertingPolicy == nil {
		result.AlertingPolicy = nil
	} else if data, err := o.AlertingPolicy.ProtoString(); err != nil {
		panic(err)
	} else {
		result.AlertingPolicy = &alerting_policy.Name{}
		if err := result.AlertingPolicy.ParseProtoString(data); err != nil {
			panic(err)
		}
	}
	result.Alerts = make([]*alert.Name, len(o.Alerts))
	for i, sourceValue := range o.Alerts {
		if sourceValue == nil {
			result.Alerts[i] = nil
		} else if data, err := sourceValue.ProtoString(); err != nil {
			panic(err)
		} else {
			result.Alerts[i] = &alert.Name{}
			if err := result.Alerts[i].ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	}
	result.State = o.State.Clone()
	return result
}

func (o *Notification) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *Notification) Merge(source *Notification) {
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
	if source.GetAlertingPolicy() != nil {
		if data, err := source.GetAlertingPolicy().ProtoString(); err != nil {
			panic(err)
		} else {
			o.AlertingPolicy = &alerting_policy.Name{}
			if err := o.AlertingPolicy.ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	} else {
		o.AlertingPolicy = nil
	}
	for _, sourceValue := range source.GetAlerts() {
		exists := false
		for _, currentValue := range o.Alerts {
			leftProtoStr, _ := currentValue.ProtoString()
			rightProtoStr, _ := sourceValue.ProtoString()
			if leftProtoStr == rightProtoStr {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement *alert.Name
			if sourceValue != nil {
				if data, err := sourceValue.ProtoString(); err != nil {
					panic(err)
				} else {
					newDstElement = &alert.Name{}
					if err := newDstElement.ParseProtoString(data); err != nil {
						panic(err)
					}
				}
			}
			o.Alerts = append(o.Alerts, newDstElement)
		}
	}

	if source.GetState() != nil {
		if o.State == nil {
			o.State = new(Notification_State)
		}
		o.State.Merge(source.GetState())
	}
}

func (o *Notification) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*Notification))
}

func (o *Notification_State) GotenObjectExt() {}

func (o *Notification_State) MakeFullFieldMask() *Notification_State_FieldMask {
	return FullNotification_State_FieldMask()
}

func (o *Notification_State) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullNotification_State_FieldMask()
}

func (o *Notification_State) MakeDiffFieldMask(other *Notification_State) *Notification_State_FieldMask {
	if o == nil && other == nil {
		return &Notification_State_FieldMask{}
	}
	if o == nil || other == nil {
		return FullNotification_State_FieldMask()
	}

	res := &Notification_State_FieldMask{}
	if o.GetIsResolved() != other.GetIsResolved() {
		res.Paths = append(res.Paths, &NotificationState_FieldTerminalPath{selector: NotificationState_FieldPathSelectorIsResolved})
	}

	if len(o.GetNotificationState()) == len(other.GetNotificationState()) {
		for i, lValue := range o.GetNotificationState() {
			rValue := other.GetNotificationState()[i]
			if len(lValue.MakeDiffFieldMask(rValue).Paths) > 0 {
				res.Paths = append(res.Paths, &NotificationState_FieldTerminalPath{selector: NotificationState_FieldPathSelectorNotificationState})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &NotificationState_FieldTerminalPath{selector: NotificationState_FieldPathSelectorNotificationState})
	}
	if o.GetIncidentNotifyAttemptsDone() != other.GetIncidentNotifyAttemptsDone() {
		res.Paths = append(res.Paths, &NotificationState_FieldTerminalPath{selector: NotificationState_FieldPathSelectorIncidentNotifyAttemptsDone})
	}
	if o.GetResolutionNotifyAttemptsDone() != other.GetResolutionNotifyAttemptsDone() {
		res.Paths = append(res.Paths, &NotificationState_FieldTerminalPath{selector: NotificationState_FieldPathSelectorResolutionNotifyAttemptsDone})
	}
	{
		subMask := o.GetAlertsLifetime().MakeDiffFieldMask(other.GetAlertsLifetime())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &NotificationState_FieldTerminalPath{selector: NotificationState_FieldPathSelectorAlertsLifetime})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &NotificationState_FieldSubPath{selector: NotificationState_FieldPathSelectorAlertsLifetime, subPath: subpath})
			}
		}
	}

	if len(o.GetResolutionNotificationState()) == len(other.GetResolutionNotificationState()) {
		for i, lValue := range o.GetResolutionNotificationState() {
			rValue := other.GetResolutionNotificationState()[i]
			if len(lValue.MakeDiffFieldMask(rValue).Paths) > 0 {
				res.Paths = append(res.Paths, &NotificationState_FieldTerminalPath{selector: NotificationState_FieldPathSelectorResolutionNotificationState})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &NotificationState_FieldTerminalPath{selector: NotificationState_FieldPathSelectorResolutionNotificationState})
	}
	if o.GetLifecycleCompleted() != other.GetLifecycleCompleted() {
		res.Paths = append(res.Paths, &NotificationState_FieldTerminalPath{selector: NotificationState_FieldPathSelectorLifecycleCompleted})
	}
	return res
}

func (o *Notification_State) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*Notification_State))
}

func (o *Notification_State) Clone() *Notification_State {
	if o == nil {
		return nil
	}
	result := &Notification_State{}
	result.IsResolved = o.IsResolved
	result.NotificationState = make([]*Notification_State_NotificationState, len(o.NotificationState))
	for i, sourceValue := range o.NotificationState {
		result.NotificationState[i] = sourceValue.Clone()
	}
	result.IncidentNotifyAttemptsDone = o.IncidentNotifyAttemptsDone
	result.ResolutionNotifyAttemptsDone = o.ResolutionNotifyAttemptsDone
	result.AlertsLifetime = o.AlertsLifetime.Clone()
	result.ResolutionNotificationState = make([]*Notification_State_NotificationState, len(o.ResolutionNotificationState))
	for i, sourceValue := range o.ResolutionNotificationState {
		result.ResolutionNotificationState[i] = sourceValue.Clone()
	}
	result.LifecycleCompleted = o.LifecycleCompleted
	return result
}

func (o *Notification_State) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *Notification_State) Merge(source *Notification_State) {
	o.IsResolved = source.GetIsResolved()
	for _, sourceValue := range source.GetNotificationState() {
		exists := false
		for _, currentValue := range o.NotificationState {
			if proto.Equal(sourceValue, currentValue) {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement *Notification_State_NotificationState
			if sourceValue != nil {
				newDstElement = new(Notification_State_NotificationState)
				newDstElement.Merge(sourceValue)
			}
			o.NotificationState = append(o.NotificationState, newDstElement)
		}
	}

	o.IncidentNotifyAttemptsDone = source.GetIncidentNotifyAttemptsDone()
	o.ResolutionNotifyAttemptsDone = source.GetResolutionNotifyAttemptsDone()
	if source.GetAlertsLifetime() != nil {
		if o.AlertsLifetime == nil {
			o.AlertsLifetime = new(common.TimeRange)
		}
		o.AlertsLifetime.Merge(source.GetAlertsLifetime())
	}
	for _, sourceValue := range source.GetResolutionNotificationState() {
		exists := false
		for _, currentValue := range o.ResolutionNotificationState {
			if proto.Equal(sourceValue, currentValue) {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement *Notification_State_NotificationState
			if sourceValue != nil {
				newDstElement = new(Notification_State_NotificationState)
				newDstElement.Merge(sourceValue)
			}
			o.ResolutionNotificationState = append(o.ResolutionNotificationState, newDstElement)
		}
	}

	o.LifecycleCompleted = source.GetLifecycleCompleted()
}

func (o *Notification_State) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*Notification_State))
}

func (o *Notification_State_NotificationState) GotenObjectExt() {}

func (o *Notification_State_NotificationState) MakeFullFieldMask() *Notification_State_NotificationState_FieldMask {
	return FullNotification_State_NotificationState_FieldMask()
}

func (o *Notification_State_NotificationState) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullNotification_State_NotificationState_FieldMask()
}

func (o *Notification_State_NotificationState) MakeDiffFieldMask(other *Notification_State_NotificationState) *Notification_State_NotificationState_FieldMask {
	if o == nil && other == nil {
		return &Notification_State_NotificationState_FieldMask{}
	}
	if o == nil || other == nil {
		return FullNotification_State_NotificationState_FieldMask()
	}

	res := &Notification_State_NotificationState_FieldMask{}
	if o.GetNotificationChannel().String() != other.GetNotificationChannel().String() {
		res.Paths = append(res.Paths, &NotificationStateNotificationState_FieldTerminalPath{selector: NotificationStateNotificationState_FieldPathSelectorNotificationChannel})
	}
	if o.GetStatus() != other.GetStatus() {
		res.Paths = append(res.Paths, &NotificationStateNotificationState_FieldTerminalPath{selector: NotificationStateNotificationState_FieldPathSelectorStatus})
	}
	if o.GetError() != other.GetError() {
		res.Paths = append(res.Paths, &NotificationStateNotificationState_FieldTerminalPath{selector: NotificationStateNotificationState_FieldPathSelectorError})
	}
	{
		subMask := o.GetProviderData().MakeDiffFieldMask(other.GetProviderData())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &NotificationStateNotificationState_FieldTerminalPath{selector: NotificationStateNotificationState_FieldPathSelectorProviderData})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &NotificationStateNotificationState_FieldSubPath{selector: NotificationStateNotificationState_FieldPathSelectorProviderData, subPath: subpath})
			}
		}
	}
	if o.GetNotifyAttempts() != other.GetNotifyAttempts() {
		res.Paths = append(res.Paths, &NotificationStateNotificationState_FieldTerminalPath{selector: NotificationStateNotificationState_FieldPathSelectorNotifyAttempts})
	}
	return res
}

func (o *Notification_State_NotificationState) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*Notification_State_NotificationState))
}

func (o *Notification_State_NotificationState) Clone() *Notification_State_NotificationState {
	if o == nil {
		return nil
	}
	result := &Notification_State_NotificationState{}
	if o.NotificationChannel == nil {
		result.NotificationChannel = nil
	} else if data, err := o.NotificationChannel.ProtoString(); err != nil {
		panic(err)
	} else {
		result.NotificationChannel = &notification_channel.Name{}
		if err := result.NotificationChannel.ParseProtoString(data); err != nil {
			panic(err)
		}
	}
	result.Status = o.Status
	result.Error = o.Error
	result.ProviderData = o.ProviderData.Clone()
	result.NotifyAttempts = o.NotifyAttempts
	return result
}

func (o *Notification_State_NotificationState) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *Notification_State_NotificationState) Merge(source *Notification_State_NotificationState) {
	if source.GetNotificationChannel() != nil {
		if data, err := source.GetNotificationChannel().ProtoString(); err != nil {
			panic(err)
		} else {
			o.NotificationChannel = &notification_channel.Name{}
			if err := o.NotificationChannel.ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	} else {
		o.NotificationChannel = nil
	}
	o.Status = source.GetStatus()
	o.Error = source.GetError()
	if source.GetProviderData() != nil {
		if o.ProviderData == nil {
			o.ProviderData = new(Notification_State_NotificationState_ProviderData)
		}
		o.ProviderData.Merge(source.GetProviderData())
	}
	o.NotifyAttempts = source.GetNotifyAttempts()
}

func (o *Notification_State_NotificationState) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*Notification_State_NotificationState))
}

func (o *Notification_State_NotificationState_ProviderData) GotenObjectExt() {}

func (o *Notification_State_NotificationState_ProviderData) MakeFullFieldMask() *Notification_State_NotificationState_ProviderData_FieldMask {
	return FullNotification_State_NotificationState_ProviderData_FieldMask()
}

func (o *Notification_State_NotificationState_ProviderData) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullNotification_State_NotificationState_ProviderData_FieldMask()
}

func (o *Notification_State_NotificationState_ProviderData) MakeDiffFieldMask(other *Notification_State_NotificationState_ProviderData) *Notification_State_NotificationState_ProviderData_FieldMask {
	if o == nil && other == nil {
		return &Notification_State_NotificationState_ProviderData_FieldMask{}
	}
	if o == nil || other == nil {
		return FullNotification_State_NotificationState_ProviderData_FieldMask()
	}

	res := &Notification_State_NotificationState_ProviderData_FieldMask{}
	{
		subMask := o.GetSlack().MakeDiffFieldMask(other.GetSlack())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &NotificationStateNotificationStateProviderData_FieldTerminalPath{selector: NotificationStateNotificationStateProviderData_FieldPathSelectorSlack})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &NotificationStateNotificationStateProviderData_FieldSubPath{selector: NotificationStateNotificationStateProviderData_FieldPathSelectorSlack, subPath: subpath})
			}
		}
	}
	{
		subMask := o.GetPagerDuty().MakeDiffFieldMask(other.GetPagerDuty())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &NotificationStateNotificationStateProviderData_FieldTerminalPath{selector: NotificationStateNotificationStateProviderData_FieldPathSelectorPagerDuty})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &NotificationStateNotificationStateProviderData_FieldSubPath{selector: NotificationStateNotificationStateProviderData_FieldPathSelectorPagerDuty, subPath: subpath})
			}
		}
	}
	return res
}

func (o *Notification_State_NotificationState_ProviderData) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*Notification_State_NotificationState_ProviderData))
}

func (o *Notification_State_NotificationState_ProviderData) Clone() *Notification_State_NotificationState_ProviderData {
	if o == nil {
		return nil
	}
	result := &Notification_State_NotificationState_ProviderData{}
	result.Slack = o.Slack.Clone()
	result.PagerDuty = o.PagerDuty.Clone()
	return result
}

func (o *Notification_State_NotificationState_ProviderData) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *Notification_State_NotificationState_ProviderData) Merge(source *Notification_State_NotificationState_ProviderData) {
	if source.GetSlack() != nil {
		if o.Slack == nil {
			o.Slack = new(Notification_State_NotificationState_ProviderData_Slack)
		}
		o.Slack.Merge(source.GetSlack())
	}
	if source.GetPagerDuty() != nil {
		if o.PagerDuty == nil {
			o.PagerDuty = new(Notification_State_NotificationState_ProviderData_PagerDuty)
		}
		o.PagerDuty.Merge(source.GetPagerDuty())
	}
}

func (o *Notification_State_NotificationState_ProviderData) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*Notification_State_NotificationState_ProviderData))
}

func (o *Notification_State_NotificationState_ProviderData_Slack) GotenObjectExt() {}

func (o *Notification_State_NotificationState_ProviderData_Slack) MakeFullFieldMask() *Notification_State_NotificationState_ProviderData_Slack_FieldMask {
	return FullNotification_State_NotificationState_ProviderData_Slack_FieldMask()
}

func (o *Notification_State_NotificationState_ProviderData_Slack) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullNotification_State_NotificationState_ProviderData_Slack_FieldMask()
}

func (o *Notification_State_NotificationState_ProviderData_Slack) MakeDiffFieldMask(other *Notification_State_NotificationState_ProviderData_Slack) *Notification_State_NotificationState_ProviderData_Slack_FieldMask {
	if o == nil && other == nil {
		return &Notification_State_NotificationState_ProviderData_Slack_FieldMask{}
	}
	if o == nil || other == nil {
		return FullNotification_State_NotificationState_ProviderData_Slack_FieldMask()
	}

	res := &Notification_State_NotificationState_ProviderData_Slack_FieldMask{}
	if o.GetTs() != other.GetTs() {
		res.Paths = append(res.Paths, &NotificationStateNotificationStateProviderDataSlack_FieldTerminalPath{selector: NotificationStateNotificationStateProviderDataSlack_FieldPathSelectorTs})
	}
	return res
}

func (o *Notification_State_NotificationState_ProviderData_Slack) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*Notification_State_NotificationState_ProviderData_Slack))
}

func (o *Notification_State_NotificationState_ProviderData_Slack) Clone() *Notification_State_NotificationState_ProviderData_Slack {
	if o == nil {
		return nil
	}
	result := &Notification_State_NotificationState_ProviderData_Slack{}
	result.Ts = o.Ts
	return result
}

func (o *Notification_State_NotificationState_ProviderData_Slack) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *Notification_State_NotificationState_ProviderData_Slack) Merge(source *Notification_State_NotificationState_ProviderData_Slack) {
	o.Ts = source.GetTs()
}

func (o *Notification_State_NotificationState_ProviderData_Slack) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*Notification_State_NotificationState_ProviderData_Slack))
}

func (o *Notification_State_NotificationState_ProviderData_PagerDuty) GotenObjectExt() {}

func (o *Notification_State_NotificationState_ProviderData_PagerDuty) MakeFullFieldMask() *Notification_State_NotificationState_ProviderData_PagerDuty_FieldMask {
	return FullNotification_State_NotificationState_ProviderData_PagerDuty_FieldMask()
}

func (o *Notification_State_NotificationState_ProviderData_PagerDuty) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullNotification_State_NotificationState_ProviderData_PagerDuty_FieldMask()
}

func (o *Notification_State_NotificationState_ProviderData_PagerDuty) MakeDiffFieldMask(other *Notification_State_NotificationState_ProviderData_PagerDuty) *Notification_State_NotificationState_ProviderData_PagerDuty_FieldMask {
	if o == nil && other == nil {
		return &Notification_State_NotificationState_ProviderData_PagerDuty_FieldMask{}
	}
	if o == nil || other == nil {
		return FullNotification_State_NotificationState_ProviderData_PagerDuty_FieldMask()
	}

	res := &Notification_State_NotificationState_ProviderData_PagerDuty_FieldMask{}
	if o.GetIncidentKey() != other.GetIncidentKey() {
		res.Paths = append(res.Paths, &NotificationStateNotificationStateProviderDataPagerDuty_FieldTerminalPath{selector: NotificationStateNotificationStateProviderDataPagerDuty_FieldPathSelectorIncidentKey})
	}
	return res
}

func (o *Notification_State_NotificationState_ProviderData_PagerDuty) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*Notification_State_NotificationState_ProviderData_PagerDuty))
}

func (o *Notification_State_NotificationState_ProviderData_PagerDuty) Clone() *Notification_State_NotificationState_ProviderData_PagerDuty {
	if o == nil {
		return nil
	}
	result := &Notification_State_NotificationState_ProviderData_PagerDuty{}
	result.IncidentKey = o.IncidentKey
	return result
}

func (o *Notification_State_NotificationState_ProviderData_PagerDuty) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *Notification_State_NotificationState_ProviderData_PagerDuty) Merge(source *Notification_State_NotificationState_ProviderData_PagerDuty) {
	o.IncidentKey = source.GetIncidentKey()
}

func (o *Notification_State_NotificationState_ProviderData_PagerDuty) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*Notification_State_NotificationState_ProviderData_PagerDuty))
}
