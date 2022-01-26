// Code generated by protoc-gen-goten-resource
// Resource: Distribution
// DO NOT EDIT!!!

package distribution

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	pod "github.com/cloudwan/edgelq-sdk/applications/resources/v1alpha/pod"
	project "github.com/cloudwan/edgelq-sdk/applications/resources/v1alpha/project"
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
)

// ensure the imports are used
var (
	_ = gotenresource.ListQuery(nil)
)

// make sure we're using proto imports
var (
	_ = &pod.Pod{}
	_ = &project.Project{}
	_ = &ntt_meta.Meta{}
)

type DistributionList []*Distribution

func (l DistributionList) Append(item gotenresource.Resource) gotenresource.ResourceList {
	return append(l, item.(*Distribution))
}

func (l DistributionList) AppendList(list gotenresource.ResourceList) gotenresource.ResourceList {
	return append(l, list.(DistributionList)...)
}

func (l DistributionList) Slice(first, second int) gotenresource.ResourceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l DistributionList) At(idx int) gotenresource.Resource {
	return l[idx]
}

func (l DistributionList) Set(idx int, res gotenresource.Resource) {
	l[idx] = res.(*Distribution)
}

func (l DistributionList) Length() int {
	return len(l)
}

type DistributionChangeList []*DistributionChange

func (l DistributionChangeList) Append(item gotenresource.ResourceChange) gotenresource.ResourceChangeList {
	return append(l, item.(*DistributionChange))
}

func (l DistributionChangeList) AppendList(list gotenresource.ResourceChangeList) gotenresource.ResourceChangeList {
	return append(l, list.(DistributionChangeList)...)
}

func (l DistributionChangeList) Slice(first, second int) gotenresource.ResourceChangeList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l DistributionChangeList) At(idx int) gotenresource.ResourceChange {
	return l[idx]
}

func (l DistributionChangeList) Set(idx int, change gotenresource.ResourceChange) {
	l[idx] = change.(*DistributionChange)
}

func (l DistributionChangeList) Length() int {
	return len(l)
}

type DistributionNameList []*Name

func (l DistributionNameList) Append(name gotenresource.Name) gotenresource.NameList {
	return append(l, name.(*Name))
}

func (l DistributionNameList) AppendList(list gotenresource.NameList) gotenresource.NameList {
	return append(l, list.(DistributionNameList)...)
}

func (l DistributionNameList) Slice(first, second int) gotenresource.NameList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l DistributionNameList) At(idx int) gotenresource.Name {
	return l[idx]
}

func (l DistributionNameList) Set(idx int, name gotenresource.Name) {
	l[idx] = name.(*Name)
}

func (l DistributionNameList) Length() int {
	return len(l)
}

type DistributionReferenceList []*Reference

func (l DistributionReferenceList) Append(ref gotenresource.Reference) gotenresource.ReferenceList {
	return append(l, ref.(*Reference))
}

func (l DistributionReferenceList) AppendList(list gotenresource.ReferenceList) gotenresource.ReferenceList {
	return append(l, list.(DistributionReferenceList)...)
}

func (l DistributionReferenceList) Slice(first, second int) gotenresource.ReferenceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l DistributionReferenceList) At(idx int) gotenresource.Reference {
	return l[idx]
}

func (l DistributionReferenceList) Set(idx int, ref gotenresource.Reference) {
	l[idx] = ref.(*Reference)
}

func (l DistributionReferenceList) Length() int {
	return len(l)
}

type DistributionParentNameList []*ParentName

func (l DistributionParentNameList) Append(name gotenresource.Name) gotenresource.ParentNameList {
	return append(l, name.(*ParentName))
}

func (l DistributionParentNameList) AppendList(list gotenresource.ParentNameList) gotenresource.ParentNameList {
	return append(l, list.(DistributionParentNameList)...)
}

func (l DistributionParentNameList) Slice(first, second int) gotenresource.ParentNameList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l DistributionParentNameList) At(idx int) gotenresource.Name {
	return l[idx]
}

func (l DistributionParentNameList) Set(idx int, name gotenresource.Name) {
	l[idx] = name.(*ParentName)
}

func (l DistributionParentNameList) Length() int {
	return len(l)
}

type DistributionParentReferenceList []*ParentReference

func (l DistributionParentReferenceList) Append(ref gotenresource.Reference) gotenresource.ParentReferenceList {
	return append(l, ref.(*ParentReference))
}

func (l DistributionParentReferenceList) AppendList(list gotenresource.ParentReferenceList) gotenresource.ParentReferenceList {
	return append(l, list.(DistributionParentReferenceList)...)
}

func (l DistributionParentReferenceList) Slice(first, second int) gotenresource.ParentReferenceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l DistributionParentReferenceList) At(idx int) gotenresource.Reference {
	return l[idx]
}

func (l DistributionParentReferenceList) Set(idx int, ref gotenresource.Reference) {
	l[idx] = ref.(*ParentReference)
}

func (l DistributionParentReferenceList) Length() int {
	return len(l)
}

type DistributionMap map[Name]*Distribution

func (m DistributionMap) Get(name gotenresource.Name) gotenresource.Resource {
	return m[*name.(*Name)]
}

func (m DistributionMap) Set(res gotenresource.Resource) {
	tRes := res.(*Distribution)
	m[*tRes.Name] = tRes
}

func (m DistributionMap) Delete(name gotenresource.Name) {
	delete(m, *name.(*Name))
}

func (m DistributionMap) Length() int {
	return len(m)
}

func (m DistributionMap) ForEach(cb func(gotenresource.Name, gotenresource.Resource) bool) {
	for name, res := range m {
		if !cb(&name, res) {
			break
		}
	}
}

type DistributionChangeMap map[Name]*DistributionChange

func (m DistributionChangeMap) Get(name gotenresource.Name) gotenresource.ResourceChange {
	return m[*name.(*Name)]
}

func (m DistributionChangeMap) Set(change gotenresource.ResourceChange) {
	tChange := change.(*DistributionChange)
	m[*tChange.GetDistributionName()] = tChange
}

func (m DistributionChangeMap) Delete(name gotenresource.Name) {
	delete(m, *name.(*Name))
}

func (m DistributionChangeMap) Length() int {
	return len(m)
}

func (m DistributionChangeMap) ForEach(cb func(gotenresource.Name, gotenresource.ResourceChange) bool) {
	for name, res := range m {
		if !cb(&name, res) {
			break
		}
	}
}
