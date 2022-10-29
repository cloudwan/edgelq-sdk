// Code generated by protoc-gen-goten-resource
// Resource: Resource
// DO NOT EDIT!!!

package resource

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
	service "github.com/cloudwan/edgelq-sdk/meta/resources/v1alpha2/service"
)

// ensure the imports are used
var (
	_ = new(gotenresource.ListQuery)
)

// make sure we're using proto imports
var (
	_ = &ntt_meta.Meta{}
	_ = &service.Service{}
)

type ResourceList []*Resource

func (l ResourceList) Append(item gotenresource.Resource) gotenresource.ResourceList {
	return append(l, item.(*Resource))
}

func (l ResourceList) AppendList(list gotenresource.ResourceList) gotenresource.ResourceList {
	return append(l, list.(ResourceList)...)
}

func (l ResourceList) Slice(first, second int) gotenresource.ResourceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l ResourceList) At(idx int) gotenresource.Resource {
	return l[idx]
}

func (l ResourceList) Set(idx int, res gotenresource.Resource) {
	l[idx] = res.(*Resource)
}

func (l ResourceList) Length() int {
	return len(l)
}

type ResourceChangeList []*ResourceChange

func (l ResourceChangeList) Append(item gotenresource.ResourceChange) gotenresource.ResourceChangeList {
	return append(l, item.(*ResourceChange))
}

func (l ResourceChangeList) AppendList(list gotenresource.ResourceChangeList) gotenresource.ResourceChangeList {
	return append(l, list.(ResourceChangeList)...)
}

func (l ResourceChangeList) Slice(first, second int) gotenresource.ResourceChangeList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l ResourceChangeList) At(idx int) gotenresource.ResourceChange {
	return l[idx]
}

func (l ResourceChangeList) Set(idx int, change gotenresource.ResourceChange) {
	l[idx] = change.(*ResourceChange)
}

func (l ResourceChangeList) Length() int {
	return len(l)
}

type ResourceNameList []*Name

func (l ResourceNameList) Append(name gotenresource.Name) gotenresource.NameList {
	return append(l, name.(*Name))
}

func (l ResourceNameList) AppendList(list gotenresource.NameList) gotenresource.NameList {
	return append(l, list.(ResourceNameList)...)
}

func (l ResourceNameList) Slice(first, second int) gotenresource.NameList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l ResourceNameList) At(idx int) gotenresource.Name {
	return l[idx]
}

func (l ResourceNameList) Set(idx int, name gotenresource.Name) {
	l[idx] = name.(*Name)
}

func (l ResourceNameList) Length() int {
	return len(l)
}

type ResourceReferenceList []*Reference

func (l ResourceReferenceList) Append(ref gotenresource.Reference) gotenresource.ReferenceList {
	return append(l, ref.(*Reference))
}

func (l ResourceReferenceList) AppendList(list gotenresource.ReferenceList) gotenresource.ReferenceList {
	return append(l, list.(ResourceReferenceList)...)
}

func (l ResourceReferenceList) Slice(first, second int) gotenresource.ReferenceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l ResourceReferenceList) At(idx int) gotenresource.Reference {
	return l[idx]
}

func (l ResourceReferenceList) Set(idx int, ref gotenresource.Reference) {
	l[idx] = ref.(*Reference)
}

func (l ResourceReferenceList) Length() int {
	return len(l)
}

type ResourceParentNameList []*ParentName

func (l ResourceParentNameList) Append(name gotenresource.Name) gotenresource.ParentNameList {
	return append(l, name.(*ParentName))
}

func (l ResourceParentNameList) AppendList(list gotenresource.ParentNameList) gotenresource.ParentNameList {
	return append(l, list.(ResourceParentNameList)...)
}

func (l ResourceParentNameList) Slice(first, second int) gotenresource.ParentNameList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l ResourceParentNameList) At(idx int) gotenresource.Name {
	return l[idx]
}

func (l ResourceParentNameList) Set(idx int, name gotenresource.Name) {
	l[idx] = name.(*ParentName)
}

func (l ResourceParentNameList) Length() int {
	return len(l)
}

type ResourceParentReferenceList []*ParentReference

func (l ResourceParentReferenceList) Append(ref gotenresource.Reference) gotenresource.ParentReferenceList {
	return append(l, ref.(*ParentReference))
}

func (l ResourceParentReferenceList) AppendList(list gotenresource.ParentReferenceList) gotenresource.ParentReferenceList {
	return append(l, list.(ResourceParentReferenceList)...)
}

func (l ResourceParentReferenceList) Slice(first, second int) gotenresource.ParentReferenceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l ResourceParentReferenceList) At(idx int) gotenresource.Reference {
	return l[idx]
}

func (l ResourceParentReferenceList) Set(idx int, ref gotenresource.Reference) {
	l[idx] = ref.(*ParentReference)
}

func (l ResourceParentReferenceList) Length() int {
	return len(l)
}

type ResourceMap map[Name]*Resource

func (m ResourceMap) Get(name gotenresource.Name) gotenresource.Resource {
	return m[*name.(*Name)]
}

func (m ResourceMap) Set(res gotenresource.Resource) {
	tRes := res.(*Resource)
	m[*tRes.Name] = tRes
}

func (m ResourceMap) Delete(name gotenresource.Name) {
	delete(m, *name.(*Name))
}

func (m ResourceMap) Length() int {
	return len(m)
}

func (m ResourceMap) ForEach(cb func(gotenresource.Name, gotenresource.Resource) bool) {
	for name, res := range m {
		if !cb(&name, res) {
			break
		}
	}
}

type ResourceChangeMap map[Name]*ResourceChange

func (m ResourceChangeMap) Get(name gotenresource.Name) gotenresource.ResourceChange {
	return m[*name.(*Name)]
}

func (m ResourceChangeMap) Set(change gotenresource.ResourceChange) {
	tChange := change.(*ResourceChange)
	m[*tChange.GetResourceName()] = tChange
}

func (m ResourceChangeMap) Delete(name gotenresource.Name) {
	delete(m, *name.(*Name))
}

func (m ResourceChangeMap) Length() int {
	return len(m)
}

func (m ResourceChangeMap) ForEach(cb func(gotenresource.Name, gotenresource.ResourceChange) bool) {
	for name, res := range m {
		if !cb(&name, res) {
			break
		}
	}
}
