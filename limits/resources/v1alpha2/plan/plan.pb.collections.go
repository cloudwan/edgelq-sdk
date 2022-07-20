// Code generated by protoc-gen-goten-resource
// Resource: Plan
// DO NOT EDIT!!!

package plan

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
	common "github.com/cloudwan/edgelq-sdk/limits/resources/v1alpha2/common"
	meta_service "github.com/cloudwan/edgelq-sdk/meta/resources/v1alpha2/service"
)

// ensure the imports are used
var (
	_ = gotenresource.ListQuery(nil)
)

// make sure we're using proto imports
var (
	_ = &ntt_meta.Meta{}
	_ = &common.Allowance{}
	_ = &meta_service.Service{}
)

type PlanList []*Plan

func (l PlanList) Append(item gotenresource.Resource) gotenresource.ResourceList {
	return append(l, item.(*Plan))
}

func (l PlanList) AppendList(list gotenresource.ResourceList) gotenresource.ResourceList {
	return append(l, list.(PlanList)...)
}

func (l PlanList) Slice(first, second int) gotenresource.ResourceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l PlanList) At(idx int) gotenresource.Resource {
	return l[idx]
}

func (l PlanList) Set(idx int, res gotenresource.Resource) {
	l[idx] = res.(*Plan)
}

func (l PlanList) Length() int {
	return len(l)
}

type PlanChangeList []*PlanChange

func (l PlanChangeList) Append(item gotenresource.ResourceChange) gotenresource.ResourceChangeList {
	return append(l, item.(*PlanChange))
}

func (l PlanChangeList) AppendList(list gotenresource.ResourceChangeList) gotenresource.ResourceChangeList {
	return append(l, list.(PlanChangeList)...)
}

func (l PlanChangeList) Slice(first, second int) gotenresource.ResourceChangeList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l PlanChangeList) At(idx int) gotenresource.ResourceChange {
	return l[idx]
}

func (l PlanChangeList) Set(idx int, change gotenresource.ResourceChange) {
	l[idx] = change.(*PlanChange)
}

func (l PlanChangeList) Length() int {
	return len(l)
}

type PlanNameList []*Name

func (l PlanNameList) Append(name gotenresource.Name) gotenresource.NameList {
	return append(l, name.(*Name))
}

func (l PlanNameList) AppendList(list gotenresource.NameList) gotenresource.NameList {
	return append(l, list.(PlanNameList)...)
}

func (l PlanNameList) Slice(first, second int) gotenresource.NameList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l PlanNameList) At(idx int) gotenresource.Name {
	return l[idx]
}

func (l PlanNameList) Set(idx int, name gotenresource.Name) {
	l[idx] = name.(*Name)
}

func (l PlanNameList) Length() int {
	return len(l)
}

type PlanReferenceList []*Reference

func (l PlanReferenceList) Append(ref gotenresource.Reference) gotenresource.ReferenceList {
	return append(l, ref.(*Reference))
}

func (l PlanReferenceList) AppendList(list gotenresource.ReferenceList) gotenresource.ReferenceList {
	return append(l, list.(PlanReferenceList)...)
}

func (l PlanReferenceList) Slice(first, second int) gotenresource.ReferenceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l PlanReferenceList) At(idx int) gotenresource.Reference {
	return l[idx]
}

func (l PlanReferenceList) Set(idx int, ref gotenresource.Reference) {
	l[idx] = ref.(*Reference)
}

func (l PlanReferenceList) Length() int {
	return len(l)
}

type PlanMap map[Name]*Plan

func (m PlanMap) Get(name gotenresource.Name) gotenresource.Resource {
	return m[*name.(*Name)]
}

func (m PlanMap) Set(res gotenresource.Resource) {
	tRes := res.(*Plan)
	m[*tRes.Name] = tRes
}

func (m PlanMap) Delete(name gotenresource.Name) {
	delete(m, *name.(*Name))
}

func (m PlanMap) Length() int {
	return len(m)
}

func (m PlanMap) ForEach(cb func(gotenresource.Name, gotenresource.Resource) bool) {
	for name, res := range m {
		if !cb(&name, res) {
			break
		}
	}
}

type PlanChangeMap map[Name]*PlanChange

func (m PlanChangeMap) Get(name gotenresource.Name) gotenresource.ResourceChange {
	return m[*name.(*Name)]
}

func (m PlanChangeMap) Set(change gotenresource.ResourceChange) {
	tChange := change.(*PlanChange)
	m[*tChange.GetPlanName()] = tChange
}

func (m PlanChangeMap) Delete(name gotenresource.Name) {
	delete(m, *name.(*Name))
}

func (m PlanChangeMap) Length() int {
	return len(m)
}

func (m PlanChangeMap) ForEach(cb func(gotenresource.Name, gotenresource.ResourceChange) bool) {
	for name, res := range m {
		if !cb(&name, res) {
			break
		}
	}
}