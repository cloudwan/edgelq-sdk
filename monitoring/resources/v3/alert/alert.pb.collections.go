// Code generated by protoc-gen-goten-resource
// Resource: Alert
// DO NOT EDIT!!!

package alert

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
	monitoring_common "github.com/cloudwan/edgelq-sdk/monitoring/common/v3"
	alerting_condition "github.com/cloudwan/edgelq-sdk/monitoring/resources/v3/alerting_condition"
)

// ensure the imports are used
var (
	_ = gotenresource.ListQuery(nil)
)

// make sure we're using proto imports
var (
	_ = &ntt_meta.Meta{}
	_ = &alerting_condition.AlertingCondition{}
	_ = &monitoring_common.LabelDescriptor{}
)

type AlertList []*Alert

func (l AlertList) Append(item gotenresource.Resource) gotenresource.ResourceList {
	return append(l, item.(*Alert))
}

func (l AlertList) AppendList(list gotenresource.ResourceList) gotenresource.ResourceList {
	return append(l, list.(AlertList)...)
}

func (l AlertList) Slice(first, second int) gotenresource.ResourceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l AlertList) At(idx int) gotenresource.Resource {
	return l[idx]
}

func (l AlertList) Set(idx int, res gotenresource.Resource) {
	l[idx] = res.(*Alert)
}

func (l AlertList) Length() int {
	return len(l)
}

type AlertChangeList []*AlertChange

func (l AlertChangeList) Append(item gotenresource.ResourceChange) gotenresource.ResourceChangeList {
	return append(l, item.(*AlertChange))
}

func (l AlertChangeList) AppendList(list gotenresource.ResourceChangeList) gotenresource.ResourceChangeList {
	return append(l, list.(AlertChangeList)...)
}

func (l AlertChangeList) Slice(first, second int) gotenresource.ResourceChangeList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l AlertChangeList) At(idx int) gotenresource.ResourceChange {
	return l[idx]
}

func (l AlertChangeList) Set(idx int, change gotenresource.ResourceChange) {
	l[idx] = change.(*AlertChange)
}

func (l AlertChangeList) Length() int {
	return len(l)
}

type AlertNameList []*Name

func (l AlertNameList) Append(name gotenresource.Name) gotenresource.NameList {
	return append(l, name.(*Name))
}

func (l AlertNameList) AppendList(list gotenresource.NameList) gotenresource.NameList {
	return append(l, list.(AlertNameList)...)
}

func (l AlertNameList) Slice(first, second int) gotenresource.NameList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l AlertNameList) At(idx int) gotenresource.Name {
	return l[idx]
}

func (l AlertNameList) Set(idx int, name gotenresource.Name) {
	l[idx] = name.(*Name)
}

func (l AlertNameList) Length() int {
	return len(l)
}

type AlertReferenceList []*Reference

func (l AlertReferenceList) Append(ref gotenresource.Reference) gotenresource.ReferenceList {
	return append(l, ref.(*Reference))
}

func (l AlertReferenceList) AppendList(list gotenresource.ReferenceList) gotenresource.ReferenceList {
	return append(l, list.(AlertReferenceList)...)
}

func (l AlertReferenceList) Slice(first, second int) gotenresource.ReferenceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l AlertReferenceList) At(idx int) gotenresource.Reference {
	return l[idx]
}

func (l AlertReferenceList) Set(idx int, ref gotenresource.Reference) {
	l[idx] = ref.(*Reference)
}

func (l AlertReferenceList) Length() int {
	return len(l)
}

type AlertParentNameList []*ParentName

func (l AlertParentNameList) Append(name gotenresource.Name) gotenresource.ParentNameList {
	return append(l, name.(*ParentName))
}

func (l AlertParentNameList) AppendList(list gotenresource.ParentNameList) gotenresource.ParentNameList {
	return append(l, list.(AlertParentNameList)...)
}

func (l AlertParentNameList) Slice(first, second int) gotenresource.ParentNameList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l AlertParentNameList) At(idx int) gotenresource.Name {
	return l[idx]
}

func (l AlertParentNameList) Set(idx int, name gotenresource.Name) {
	l[idx] = name.(*ParentName)
}

func (l AlertParentNameList) Length() int {
	return len(l)
}

type AlertParentReferenceList []*ParentReference

func (l AlertParentReferenceList) Append(ref gotenresource.Reference) gotenresource.ParentReferenceList {
	return append(l, ref.(*ParentReference))
}

func (l AlertParentReferenceList) AppendList(list gotenresource.ParentReferenceList) gotenresource.ParentReferenceList {
	return append(l, list.(AlertParentReferenceList)...)
}

func (l AlertParentReferenceList) Slice(first, second int) gotenresource.ParentReferenceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l AlertParentReferenceList) At(idx int) gotenresource.Reference {
	return l[idx]
}

func (l AlertParentReferenceList) Set(idx int, ref gotenresource.Reference) {
	l[idx] = ref.(*ParentReference)
}

func (l AlertParentReferenceList) Length() int {
	return len(l)
}

type AlertMap map[Name]*Alert

func (m AlertMap) Get(name gotenresource.Name) gotenresource.Resource {
	return m[*name.(*Name)]
}

func (m AlertMap) Set(res gotenresource.Resource) {
	tRes := res.(*Alert)
	m[*tRes.Name] = tRes
}

func (m AlertMap) Delete(name gotenresource.Name) {
	delete(m, *name.(*Name))
}

func (m AlertMap) Length() int {
	return len(m)
}

func (m AlertMap) ForEach(cb func(gotenresource.Name, gotenresource.Resource) bool) {
	for name, res := range m {
		if !cb(&name, res) {
			break
		}
	}
}

type AlertChangeMap map[Name]*AlertChange

func (m AlertChangeMap) Get(name gotenresource.Name) gotenresource.ResourceChange {
	return m[*name.(*Name)]
}

func (m AlertChangeMap) Set(change gotenresource.ResourceChange) {
	tChange := change.(*AlertChange)
	m[*tChange.GetAlertName()] = tChange
}

func (m AlertChangeMap) Delete(name gotenresource.Name) {
	delete(m, *name.(*Name))
}

func (m AlertChangeMap) Length() int {
	return len(m)
}

func (m AlertChangeMap) ForEach(cb func(gotenresource.Name, gotenresource.ResourceChange) bool) {
	for name, res := range m {
		if !cb(&name, res) {
			break
		}
	}
}
