// Code generated by protoc-gen-goten-resource
// Resource: AlertingCondition
// DO NOT EDIT!!!

package alerting_condition

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	alerting_policy "github.com/cloudwan/edgelq-sdk/monitoring/resources/v3/alerting_policy"
	common "github.com/cloudwan/edgelq-sdk/monitoring/resources/v3/common"
	meta "github.com/cloudwan/goten-sdk/types/meta"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
)

// ensure the imports are used
var (
	_ = new(gotenresource.ListQuery)
)

// make sure we're using proto imports
var (
	_ = &alerting_policy.AlertingPolicy{}
	_ = &common.LabelDescriptor{}
	_ = &durationpb.Duration{}
	_ = &meta.Meta{}
)

type AlertingConditionList []*AlertingCondition

func (l AlertingConditionList) Append(item gotenresource.Resource) gotenresource.ResourceList {
	return append(l, item.(*AlertingCondition))
}

func (l AlertingConditionList) AppendList(list gotenresource.ResourceList) gotenresource.ResourceList {
	return append(l, list.(AlertingConditionList)...)
}

func (l AlertingConditionList) Slice(first, second int) gotenresource.ResourceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l AlertingConditionList) At(idx int) gotenresource.Resource {
	return l[idx]
}

func (l AlertingConditionList) Set(idx int, res gotenresource.Resource) {
	l[idx] = res.(*AlertingCondition)
}

func (l AlertingConditionList) Length() int {
	return len(l)
}

type AlertingConditionChangeList []*AlertingConditionChange

func (l AlertingConditionChangeList) Append(item gotenresource.ResourceChange) gotenresource.ResourceChangeList {
	return append(l, item.(*AlertingConditionChange))
}

func (l AlertingConditionChangeList) AppendList(list gotenresource.ResourceChangeList) gotenresource.ResourceChangeList {
	return append(l, list.(AlertingConditionChangeList)...)
}

func (l AlertingConditionChangeList) Slice(first, second int) gotenresource.ResourceChangeList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l AlertingConditionChangeList) At(idx int) gotenresource.ResourceChange {
	return l[idx]
}

func (l AlertingConditionChangeList) Set(idx int, change gotenresource.ResourceChange) {
	l[idx] = change.(*AlertingConditionChange)
}

func (l AlertingConditionChangeList) Length() int {
	return len(l)
}

type AlertingConditionNameList []*Name

func (l AlertingConditionNameList) Append(name gotenresource.Name) gotenresource.NameList {
	return append(l, name.(*Name))
}

func (l AlertingConditionNameList) AppendList(list gotenresource.NameList) gotenresource.NameList {
	return append(l, list.(AlertingConditionNameList)...)
}

func (l AlertingConditionNameList) Slice(first, second int) gotenresource.NameList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l AlertingConditionNameList) At(idx int) gotenresource.Name {
	return l[idx]
}

func (l AlertingConditionNameList) Set(idx int, name gotenresource.Name) {
	l[idx] = name.(*Name)
}

func (l AlertingConditionNameList) Length() int {
	return len(l)
}

type AlertingConditionReferenceList []*Reference

func (l AlertingConditionReferenceList) Append(ref gotenresource.Reference) gotenresource.ReferenceList {
	return append(l, ref.(*Reference))
}

func (l AlertingConditionReferenceList) AppendList(list gotenresource.ReferenceList) gotenresource.ReferenceList {
	return append(l, list.(AlertingConditionReferenceList)...)
}

func (l AlertingConditionReferenceList) Slice(first, second int) gotenresource.ReferenceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l AlertingConditionReferenceList) At(idx int) gotenresource.Reference {
	return l[idx]
}

func (l AlertingConditionReferenceList) Set(idx int, ref gotenresource.Reference) {
	l[idx] = ref.(*Reference)
}

func (l AlertingConditionReferenceList) Length() int {
	return len(l)
}

type AlertingConditionParentNameList []*ParentName

func (l AlertingConditionParentNameList) Append(name gotenresource.Name) gotenresource.ParentNameList {
	return append(l, name.(*ParentName))
}

func (l AlertingConditionParentNameList) AppendList(list gotenresource.ParentNameList) gotenresource.ParentNameList {
	return append(l, list.(AlertingConditionParentNameList)...)
}

func (l AlertingConditionParentNameList) Slice(first, second int) gotenresource.ParentNameList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l AlertingConditionParentNameList) At(idx int) gotenresource.Name {
	return l[idx]
}

func (l AlertingConditionParentNameList) Set(idx int, name gotenresource.Name) {
	l[idx] = name.(*ParentName)
}

func (l AlertingConditionParentNameList) Length() int {
	return len(l)
}

type AlertingConditionParentReferenceList []*ParentReference

func (l AlertingConditionParentReferenceList) Append(ref gotenresource.Reference) gotenresource.ParentReferenceList {
	return append(l, ref.(*ParentReference))
}

func (l AlertingConditionParentReferenceList) AppendList(list gotenresource.ParentReferenceList) gotenresource.ParentReferenceList {
	return append(l, list.(AlertingConditionParentReferenceList)...)
}

func (l AlertingConditionParentReferenceList) Slice(first, second int) gotenresource.ParentReferenceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l AlertingConditionParentReferenceList) At(idx int) gotenresource.Reference {
	return l[idx]
}

func (l AlertingConditionParentReferenceList) Set(idx int, ref gotenresource.Reference) {
	l[idx] = ref.(*ParentReference)
}

func (l AlertingConditionParentReferenceList) Length() int {
	return len(l)
}

type AlertingConditionMap map[Name]*AlertingCondition

func (m AlertingConditionMap) Get(name gotenresource.Name) gotenresource.Resource {
	return m[*name.(*Name)]
}

func (m AlertingConditionMap) Set(res gotenresource.Resource) {
	tRes := res.(*AlertingCondition)
	m[*tRes.Name] = tRes
}

func (m AlertingConditionMap) Delete(name gotenresource.Name) {
	delete(m, *name.(*Name))
}

func (m AlertingConditionMap) Length() int {
	return len(m)
}

func (m AlertingConditionMap) ForEach(cb func(gotenresource.Name, gotenresource.Resource) bool) {
	for name, res := range m {
		if !cb(&name, res) {
			break
		}
	}
}

type AlertingConditionChangeMap map[Name]*AlertingConditionChange

func (m AlertingConditionChangeMap) Get(name gotenresource.Name) gotenresource.ResourceChange {
	return m[*name.(*Name)]
}

func (m AlertingConditionChangeMap) Set(change gotenresource.ResourceChange) {
	tChange := change.(*AlertingConditionChange)
	m[*tChange.GetAlertingConditionName()] = tChange
}

func (m AlertingConditionChangeMap) Delete(name gotenresource.Name) {
	delete(m, *name.(*Name))
}

func (m AlertingConditionChangeMap) Length() int {
	return len(m)
}

func (m AlertingConditionChangeMap) ForEach(cb func(gotenresource.Name, gotenresource.ResourceChange) bool) {
	for name, res := range m {
		if !cb(&name, res) {
			break
		}
	}
}
