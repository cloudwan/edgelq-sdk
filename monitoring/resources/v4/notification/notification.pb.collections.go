// Code generated by protoc-gen-goten-resource
// Resource: Notification
// DO NOT EDIT!!!

package notification

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	alert "github.com/cloudwan/edgelq-sdk/monitoring/resources/v4/alert"
	alerting_condition "github.com/cloudwan/edgelq-sdk/monitoring/resources/v4/alerting_condition"
	alerting_policy "github.com/cloudwan/edgelq-sdk/monitoring/resources/v4/alerting_policy"
	common "github.com/cloudwan/edgelq-sdk/monitoring/resources/v4/common"
	notification_channel "github.com/cloudwan/edgelq-sdk/monitoring/resources/v4/notification_channel"
	project "github.com/cloudwan/edgelq-sdk/monitoring/resources/v4/project"
	meta "github.com/cloudwan/goten-sdk/types/meta"
)

// ensure the imports are used
var (
	_ = new(gotenresource.ListQuery)
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

type NotificationList []*Notification

func (l NotificationList) Append(item gotenresource.Resource) gotenresource.ResourceList {
	return append(l, item.(*Notification))
}

func (l NotificationList) AppendList(list gotenresource.ResourceList) gotenresource.ResourceList {
	return append(l, list.(NotificationList)...)
}

func (l NotificationList) Slice(first, second int) gotenresource.ResourceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l NotificationList) At(idx int) gotenresource.Resource {
	return l[idx]
}

func (l NotificationList) Set(idx int, res gotenresource.Resource) {
	l[idx] = res.(*Notification)
}

func (l NotificationList) Length() int {
	return len(l)
}

type NotificationChangeList []*NotificationChange

func (l NotificationChangeList) Append(item gotenresource.ResourceChange) gotenresource.ResourceChangeList {
	return append(l, item.(*NotificationChange))
}

func (l NotificationChangeList) AppendList(list gotenresource.ResourceChangeList) gotenresource.ResourceChangeList {
	return append(l, list.(NotificationChangeList)...)
}

func (l NotificationChangeList) Slice(first, second int) gotenresource.ResourceChangeList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l NotificationChangeList) At(idx int) gotenresource.ResourceChange {
	return l[idx]
}

func (l NotificationChangeList) Set(idx int, change gotenresource.ResourceChange) {
	l[idx] = change.(*NotificationChange)
}

func (l NotificationChangeList) Length() int {
	return len(l)
}

type NotificationNameList []*Name

func (l NotificationNameList) Append(name gotenresource.Name) gotenresource.NameList {
	return append(l, name.(*Name))
}

func (l NotificationNameList) AppendList(list gotenresource.NameList) gotenresource.NameList {
	return append(l, list.(NotificationNameList)...)
}

func (l NotificationNameList) Slice(first, second int) gotenresource.NameList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l NotificationNameList) At(idx int) gotenresource.Name {
	return l[idx]
}

func (l NotificationNameList) Set(idx int, name gotenresource.Name) {
	l[idx] = name.(*Name)
}

func (l NotificationNameList) Length() int {
	return len(l)
}

type NotificationReferenceList []*Reference

func (l NotificationReferenceList) Append(ref gotenresource.Reference) gotenresource.ReferenceList {
	return append(l, ref.(*Reference))
}

func (l NotificationReferenceList) AppendList(list gotenresource.ReferenceList) gotenresource.ReferenceList {
	return append(l, list.(NotificationReferenceList)...)
}

func (l NotificationReferenceList) Slice(first, second int) gotenresource.ReferenceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l NotificationReferenceList) At(idx int) gotenresource.Reference {
	return l[idx]
}

func (l NotificationReferenceList) Set(idx int, ref gotenresource.Reference) {
	l[idx] = ref.(*Reference)
}

func (l NotificationReferenceList) Length() int {
	return len(l)
}

type NotificationParentNameList []*ParentName

func (l NotificationParentNameList) Append(name gotenresource.Name) gotenresource.ParentNameList {
	return append(l, name.(*ParentName))
}

func (l NotificationParentNameList) AppendList(list gotenresource.ParentNameList) gotenresource.ParentNameList {
	return append(l, list.(NotificationParentNameList)...)
}

func (l NotificationParentNameList) Slice(first, second int) gotenresource.ParentNameList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l NotificationParentNameList) At(idx int) gotenresource.Name {
	return l[idx]
}

func (l NotificationParentNameList) Set(idx int, name gotenresource.Name) {
	l[idx] = name.(*ParentName)
}

func (l NotificationParentNameList) Length() int {
	return len(l)
}

type NotificationParentReferenceList []*ParentReference

func (l NotificationParentReferenceList) Append(ref gotenresource.Reference) gotenresource.ParentReferenceList {
	return append(l, ref.(*ParentReference))
}

func (l NotificationParentReferenceList) AppendList(list gotenresource.ParentReferenceList) gotenresource.ParentReferenceList {
	return append(l, list.(NotificationParentReferenceList)...)
}

func (l NotificationParentReferenceList) Slice(first, second int) gotenresource.ParentReferenceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l NotificationParentReferenceList) At(idx int) gotenresource.Reference {
	return l[idx]
}

func (l NotificationParentReferenceList) Set(idx int, ref gotenresource.Reference) {
	l[idx] = ref.(*ParentReference)
}

func (l NotificationParentReferenceList) Length() int {
	return len(l)
}

type NotificationMap map[Name]*Notification

func (m NotificationMap) Get(name gotenresource.Name) gotenresource.Resource {
	return m[*name.(*Name)]
}

func (m NotificationMap) Set(res gotenresource.Resource) {
	tRes := res.(*Notification)
	m[*tRes.Name] = tRes
}

func (m NotificationMap) Delete(name gotenresource.Name) {
	delete(m, *name.(*Name))
}

func (m NotificationMap) Length() int {
	return len(m)
}

func (m NotificationMap) ForEach(cb func(gotenresource.Name, gotenresource.Resource) bool) {
	for name, res := range m {
		if !cb(&name, res) {
			break
		}
	}
}

type NotificationChangeMap map[Name]*NotificationChange

func (m NotificationChangeMap) Get(name gotenresource.Name) gotenresource.ResourceChange {
	return m[*name.(*Name)]
}

func (m NotificationChangeMap) Set(change gotenresource.ResourceChange) {
	tChange := change.(*NotificationChange)
	m[*tChange.GetNotificationName()] = tChange
}

func (m NotificationChangeMap) Delete(name gotenresource.Name) {
	delete(m, *name.(*Name))
}

func (m NotificationChangeMap) Length() int {
	return len(m)
}

func (m NotificationChangeMap) ForEach(cb func(gotenresource.Name, gotenresource.ResourceChange) bool) {
	for name, res := range m {
		if !cb(&name, res) {
			break
		}
	}
}
