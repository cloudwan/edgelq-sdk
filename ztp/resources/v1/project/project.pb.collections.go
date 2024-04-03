// Code generated by protoc-gen-goten-resource
// Resource: Project
// DO NOT EDIT!!!

package project

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	meta "github.com/cloudwan/goten-sdk/types/meta"
	multi_region_policy "github.com/cloudwan/goten-sdk/types/multi_region_policy"
)

// ensure the imports are used
var (
	_ = new(gotenresource.ListQuery)
)

// make sure we're using proto imports
var (
	_ = &meta.Meta{}
	_ = &multi_region_policy.MultiRegionPolicy{}
)

type ProjectList []*Project

func (l ProjectList) Append(item gotenresource.Resource) gotenresource.ResourceList {
	return append(l, item.(*Project))
}

func (l ProjectList) AppendList(list gotenresource.ResourceList) gotenresource.ResourceList {
	return append(l, list.(ProjectList)...)
}

func (l ProjectList) Slice(first, second int) gotenresource.ResourceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l ProjectList) At(idx int) gotenresource.Resource {
	return l[idx]
}

func (l ProjectList) Set(idx int, res gotenresource.Resource) {
	l[idx] = res.(*Project)
}

func (l ProjectList) Length() int {
	return len(l)
}

type ProjectChangeList []*ProjectChange

func (l ProjectChangeList) Append(item gotenresource.ResourceChange) gotenresource.ResourceChangeList {
	return append(l, item.(*ProjectChange))
}

func (l ProjectChangeList) AppendList(list gotenresource.ResourceChangeList) gotenresource.ResourceChangeList {
	return append(l, list.(ProjectChangeList)...)
}

func (l ProjectChangeList) Slice(first, second int) gotenresource.ResourceChangeList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l ProjectChangeList) At(idx int) gotenresource.ResourceChange {
	return l[idx]
}

func (l ProjectChangeList) Set(idx int, change gotenresource.ResourceChange) {
	l[idx] = change.(*ProjectChange)
}

func (l ProjectChangeList) Length() int {
	return len(l)
}

type ProjectNameList []*Name

func (l ProjectNameList) Append(name gotenresource.Name) gotenresource.NameList {
	return append(l, name.(*Name))
}

func (l ProjectNameList) AppendList(list gotenresource.NameList) gotenresource.NameList {
	return append(l, list.(ProjectNameList)...)
}

func (l ProjectNameList) Slice(first, second int) gotenresource.NameList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l ProjectNameList) At(idx int) gotenresource.Name {
	return l[idx]
}

func (l ProjectNameList) Set(idx int, name gotenresource.Name) {
	l[idx] = name.(*Name)
}

func (l ProjectNameList) Length() int {
	return len(l)
}

type ProjectReferenceList []*Reference

func (l ProjectReferenceList) Append(ref gotenresource.Reference) gotenresource.ReferenceList {
	return append(l, ref.(*Reference))
}

func (l ProjectReferenceList) AppendList(list gotenresource.ReferenceList) gotenresource.ReferenceList {
	return append(l, list.(ProjectReferenceList)...)
}

func (l ProjectReferenceList) Slice(first, second int) gotenresource.ReferenceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l ProjectReferenceList) At(idx int) gotenresource.Reference {
	return l[idx]
}

func (l ProjectReferenceList) Set(idx int, ref gotenresource.Reference) {
	l[idx] = ref.(*Reference)
}

func (l ProjectReferenceList) Length() int {
	return len(l)
}

type ProjectMap map[Name]*Project

func (m ProjectMap) Get(name gotenresource.Name) gotenresource.Resource {
	return m[*name.(*Name)]
}

func (m ProjectMap) Set(res gotenresource.Resource) {
	tRes := res.(*Project)
	m[*tRes.Name] = tRes
}

func (m ProjectMap) Delete(name gotenresource.Name) {
	delete(m, *name.(*Name))
}

func (m ProjectMap) Length() int {
	return len(m)
}

func (m ProjectMap) ForEach(cb func(gotenresource.Name, gotenresource.Resource) bool) {
	for name, res := range m {
		if !cb(&name, res) {
			break
		}
	}
}

type ProjectChangeMap map[Name]*ProjectChange

func (m ProjectChangeMap) Get(name gotenresource.Name) gotenresource.ResourceChange {
	return m[*name.(*Name)]
}

func (m ProjectChangeMap) Set(change gotenresource.ResourceChange) {
	tChange := change.(*ProjectChange)
	m[*tChange.GetProjectName()] = tChange
}

func (m ProjectChangeMap) Delete(name gotenresource.Name) {
	delete(m, *name.(*Name))
}

func (m ProjectChangeMap) Length() int {
	return len(m)
}

func (m ProjectChangeMap) ForEach(cb func(gotenresource.Name, gotenresource.ResourceChange) bool) {
	for name, res := range m {
		if !cb(&name, res) {
			break
		}
	}
}