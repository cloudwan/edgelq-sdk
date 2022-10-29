// Code generated by protoc-gen-goten-resource
// Resource: Group
// DO NOT EDIT!!!

package group

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
	organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/organization"
	project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/project"
)

// ensure the imports are used
var (
	_ = new(gotenresource.ListQuery)
)

// make sure we're using proto imports
var (
	_ = &ntt_meta.Meta{}
	_ = &organization.Organization{}
	_ = &project.Project{}
)

type GroupList []*Group

func (l GroupList) Append(item gotenresource.Resource) gotenresource.ResourceList {
	return append(l, item.(*Group))
}

func (l GroupList) AppendList(list gotenresource.ResourceList) gotenresource.ResourceList {
	return append(l, list.(GroupList)...)
}

func (l GroupList) Slice(first, second int) gotenresource.ResourceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l GroupList) At(idx int) gotenresource.Resource {
	return l[idx]
}

func (l GroupList) Set(idx int, res gotenresource.Resource) {
	l[idx] = res.(*Group)
}

func (l GroupList) Length() int {
	return len(l)
}

type GroupChangeList []*GroupChange

func (l GroupChangeList) Append(item gotenresource.ResourceChange) gotenresource.ResourceChangeList {
	return append(l, item.(*GroupChange))
}

func (l GroupChangeList) AppendList(list gotenresource.ResourceChangeList) gotenresource.ResourceChangeList {
	return append(l, list.(GroupChangeList)...)
}

func (l GroupChangeList) Slice(first, second int) gotenresource.ResourceChangeList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l GroupChangeList) At(idx int) gotenresource.ResourceChange {
	return l[idx]
}

func (l GroupChangeList) Set(idx int, change gotenresource.ResourceChange) {
	l[idx] = change.(*GroupChange)
}

func (l GroupChangeList) Length() int {
	return len(l)
}

type GroupNameList []*Name

func (l GroupNameList) Append(name gotenresource.Name) gotenresource.NameList {
	return append(l, name.(*Name))
}

func (l GroupNameList) AppendList(list gotenresource.NameList) gotenresource.NameList {
	return append(l, list.(GroupNameList)...)
}

func (l GroupNameList) Slice(first, second int) gotenresource.NameList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l GroupNameList) At(idx int) gotenresource.Name {
	return l[idx]
}

func (l GroupNameList) Set(idx int, name gotenresource.Name) {
	l[idx] = name.(*Name)
}

func (l GroupNameList) Length() int {
	return len(l)
}

type GroupReferenceList []*Reference

func (l GroupReferenceList) Append(ref gotenresource.Reference) gotenresource.ReferenceList {
	return append(l, ref.(*Reference))
}

func (l GroupReferenceList) AppendList(list gotenresource.ReferenceList) gotenresource.ReferenceList {
	return append(l, list.(GroupReferenceList)...)
}

func (l GroupReferenceList) Slice(first, second int) gotenresource.ReferenceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l GroupReferenceList) At(idx int) gotenresource.Reference {
	return l[idx]
}

func (l GroupReferenceList) Set(idx int, ref gotenresource.Reference) {
	l[idx] = ref.(*Reference)
}

func (l GroupReferenceList) Length() int {
	return len(l)
}

type GroupParentNameList []*ParentName

func (l GroupParentNameList) Append(name gotenresource.Name) gotenresource.ParentNameList {
	return append(l, name.(*ParentName))
}

func (l GroupParentNameList) AppendList(list gotenresource.ParentNameList) gotenresource.ParentNameList {
	return append(l, list.(GroupParentNameList)...)
}

func (l GroupParentNameList) Slice(first, second int) gotenresource.ParentNameList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l GroupParentNameList) At(idx int) gotenresource.Name {
	return l[idx]
}

func (l GroupParentNameList) Set(idx int, name gotenresource.Name) {
	l[idx] = name.(*ParentName)
}

func (l GroupParentNameList) Length() int {
	return len(l)
}

type GroupParentReferenceList []*ParentReference

func (l GroupParentReferenceList) Append(ref gotenresource.Reference) gotenresource.ParentReferenceList {
	return append(l, ref.(*ParentReference))
}

func (l GroupParentReferenceList) AppendList(list gotenresource.ParentReferenceList) gotenresource.ParentReferenceList {
	return append(l, list.(GroupParentReferenceList)...)
}

func (l GroupParentReferenceList) Slice(first, second int) gotenresource.ParentReferenceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l GroupParentReferenceList) At(idx int) gotenresource.Reference {
	return l[idx]
}

func (l GroupParentReferenceList) Set(idx int, ref gotenresource.Reference) {
	l[idx] = ref.(*ParentReference)
}

func (l GroupParentReferenceList) Length() int {
	return len(l)
}

type GroupMap map[Name]*Group

func (m GroupMap) Get(name gotenresource.Name) gotenresource.Resource {
	return m[*name.(*Name)]
}

func (m GroupMap) Set(res gotenresource.Resource) {
	tRes := res.(*Group)
	m[*tRes.Name] = tRes
}

func (m GroupMap) Delete(name gotenresource.Name) {
	delete(m, *name.(*Name))
}

func (m GroupMap) Length() int {
	return len(m)
}

func (m GroupMap) ForEach(cb func(gotenresource.Name, gotenresource.Resource) bool) {
	for name, res := range m {
		if !cb(&name, res) {
			break
		}
	}
}

type GroupChangeMap map[Name]*GroupChange

func (m GroupChangeMap) Get(name gotenresource.Name) gotenresource.ResourceChange {
	return m[*name.(*Name)]
}

func (m GroupChangeMap) Set(change gotenresource.ResourceChange) {
	tChange := change.(*GroupChange)
	m[*tChange.GetGroupName()] = tChange
}

func (m GroupChangeMap) Delete(name gotenresource.Name) {
	delete(m, *name.(*Name))
}

func (m GroupChangeMap) Length() int {
	return len(m)
}

func (m GroupChangeMap) ForEach(cb func(gotenresource.Name, gotenresource.ResourceChange) bool) {
	for name, res := range m {
		if !cb(&name, res) {
			break
		}
	}
}
