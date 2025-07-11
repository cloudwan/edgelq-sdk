// Code generated by protoc-gen-goten-resource
// Resource: LogCondition
// DO NOT EDIT!!!

package log_condition

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	rcommon "github.com/cloudwan/edgelq-sdk/alerting/resources/v1/common"
	document "github.com/cloudwan/edgelq-sdk/alerting/resources/v1/document"
	log_condition_template "github.com/cloudwan/edgelq-sdk/alerting/resources/v1/log_condition_template"
	policy "github.com/cloudwan/edgelq-sdk/alerting/resources/v1/policy"
	meta "github.com/cloudwan/goten-sdk/types/meta"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

// ensure the imports are used
var (
	_ = new(gotenresource.ListQuery)
)

// make sure we're using proto imports
var (
	_ = &document.Document{}
	_ = &log_condition_template.LogConditionTemplate{}
	_ = &policy.Policy{}
	_ = &rcommon.LogCndSpec{}
	_ = &fieldmaskpb.FieldMask{}
	_ = &meta.Meta{}
)

type LogConditionList []*LogCondition

func (l LogConditionList) Append(item gotenresource.Resource) gotenresource.ResourceList {
	return append(l, item.(*LogCondition))
}

func (l LogConditionList) AppendList(list gotenresource.ResourceList) gotenresource.ResourceList {
	return append(l, list.(LogConditionList)...)
}

func (l LogConditionList) Slice(first, second int) gotenresource.ResourceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l LogConditionList) At(idx int) gotenresource.Resource {
	return l[idx]
}

func (l LogConditionList) Set(idx int, res gotenresource.Resource) {
	l[idx] = res.(*LogCondition)
}

func (l LogConditionList) Length() int {
	return len(l)
}

type LogConditionChangeList []*LogConditionChange

func (l LogConditionChangeList) Append(item gotenresource.ResourceChange) gotenresource.ResourceChangeList {
	return append(l, item.(*LogConditionChange))
}

func (l LogConditionChangeList) AppendList(list gotenresource.ResourceChangeList) gotenresource.ResourceChangeList {
	return append(l, list.(LogConditionChangeList)...)
}

func (l LogConditionChangeList) Slice(first, second int) gotenresource.ResourceChangeList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l LogConditionChangeList) At(idx int) gotenresource.ResourceChange {
	return l[idx]
}

func (l LogConditionChangeList) Set(idx int, change gotenresource.ResourceChange) {
	l[idx] = change.(*LogConditionChange)
}

func (l LogConditionChangeList) Length() int {
	return len(l)
}

type LogConditionNameList []*Name

func (l LogConditionNameList) Append(name gotenresource.Name) gotenresource.NameList {
	return append(l, name.(*Name))
}

func (l LogConditionNameList) AppendList(list gotenresource.NameList) gotenresource.NameList {
	return append(l, list.(LogConditionNameList)...)
}

func (l LogConditionNameList) Slice(first, second int) gotenresource.NameList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l LogConditionNameList) At(idx int) gotenresource.Name {
	return l[idx]
}

func (l LogConditionNameList) Set(idx int, name gotenresource.Name) {
	l[idx] = name.(*Name)
}

func (l LogConditionNameList) Length() int {
	return len(l)
}

type LogConditionReferenceList []*Reference

func (l LogConditionReferenceList) Append(ref gotenresource.Reference) gotenresource.ReferenceList {
	return append(l, ref.(*Reference))
}

func (l LogConditionReferenceList) AppendList(list gotenresource.ReferenceList) gotenresource.ReferenceList {
	return append(l, list.(LogConditionReferenceList)...)
}

func (l LogConditionReferenceList) Slice(first, second int) gotenresource.ReferenceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l LogConditionReferenceList) At(idx int) gotenresource.Reference {
	return l[idx]
}

func (l LogConditionReferenceList) Set(idx int, ref gotenresource.Reference) {
	l[idx] = ref.(*Reference)
}

func (l LogConditionReferenceList) Length() int {
	return len(l)
}

type LogConditionParentNameList []*ParentName

func (l LogConditionParentNameList) Append(name gotenresource.Name) gotenresource.ParentNameList {
	return append(l, name.(*ParentName))
}

func (l LogConditionParentNameList) AppendList(list gotenresource.ParentNameList) gotenresource.ParentNameList {
	return append(l, list.(LogConditionParentNameList)...)
}

func (l LogConditionParentNameList) Slice(first, second int) gotenresource.ParentNameList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l LogConditionParentNameList) At(idx int) gotenresource.Name {
	return l[idx]
}

func (l LogConditionParentNameList) Set(idx int, name gotenresource.Name) {
	l[idx] = name.(*ParentName)
}

func (l LogConditionParentNameList) Length() int {
	return len(l)
}

type LogConditionParentReferenceList []*ParentReference

func (l LogConditionParentReferenceList) Append(ref gotenresource.Reference) gotenresource.ParentReferenceList {
	return append(l, ref.(*ParentReference))
}

func (l LogConditionParentReferenceList) AppendList(list gotenresource.ParentReferenceList) gotenresource.ParentReferenceList {
	return append(l, list.(LogConditionParentReferenceList)...)
}

func (l LogConditionParentReferenceList) Slice(first, second int) gotenresource.ParentReferenceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l LogConditionParentReferenceList) At(idx int) gotenresource.Reference {
	return l[idx]
}

func (l LogConditionParentReferenceList) Set(idx int, ref gotenresource.Reference) {
	l[idx] = ref.(*ParentReference)
}

func (l LogConditionParentReferenceList) Length() int {
	return len(l)
}

type LogConditionMap map[Name]*LogCondition

func (m LogConditionMap) Get(name gotenresource.Name) gotenresource.Resource {
	return m[*name.(*Name)]
}

func (m LogConditionMap) Set(res gotenresource.Resource) {
	tRes := res.(*LogCondition)
	m[*tRes.Name] = tRes
}

func (m LogConditionMap) Delete(name gotenresource.Name) {
	delete(m, *name.(*Name))
}

func (m LogConditionMap) Length() int {
	return len(m)
}

func (m LogConditionMap) ForEach(cb func(gotenresource.Name, gotenresource.Resource) bool) {
	for name, res := range m {
		if !cb(&name, res) {
			break
		}
	}
}

type LogConditionChangeMap map[Name]*LogConditionChange

func (m LogConditionChangeMap) Get(name gotenresource.Name) gotenresource.ResourceChange {
	return m[*name.(*Name)]
}

func (m LogConditionChangeMap) Set(change gotenresource.ResourceChange) {
	tChange := change.(*LogConditionChange)
	m[*tChange.GetLogConditionName()] = tChange
}

func (m LogConditionChangeMap) Delete(name gotenresource.Name) {
	delete(m, *name.(*Name))
}

func (m LogConditionChangeMap) Length() int {
	return len(m)
}

func (m LogConditionChangeMap) ForEach(cb func(gotenresource.Name, gotenresource.ResourceChange) bool) {
	for name, res := range m {
		if !cb(&name, res) {
			break
		}
	}
}
