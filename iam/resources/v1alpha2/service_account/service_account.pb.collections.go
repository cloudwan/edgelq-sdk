// Code generated by protoc-gen-goten-resource
// Resource: ServiceAccount
// DO NOT EDIT!!!

package service_account

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/project"
	meta "github.com/cloudwan/goten-sdk/types/meta"
)

// ensure the imports are used
var (
	_ = new(gotenresource.ListQuery)
)

// make sure we're using proto imports
var (
	_ = &project.Project{}
	_ = &meta.Meta{}
)

type ServiceAccountList []*ServiceAccount

func (l ServiceAccountList) Append(item gotenresource.Resource) gotenresource.ResourceList {
	return append(l, item.(*ServiceAccount))
}

func (l ServiceAccountList) AppendList(list gotenresource.ResourceList) gotenresource.ResourceList {
	return append(l, list.(ServiceAccountList)...)
}

func (l ServiceAccountList) Slice(first, second int) gotenresource.ResourceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l ServiceAccountList) At(idx int) gotenresource.Resource {
	return l[idx]
}

func (l ServiceAccountList) Set(idx int, res gotenresource.Resource) {
	l[idx] = res.(*ServiceAccount)
}

func (l ServiceAccountList) Length() int {
	return len(l)
}

type ServiceAccountChangeList []*ServiceAccountChange

func (l ServiceAccountChangeList) Append(item gotenresource.ResourceChange) gotenresource.ResourceChangeList {
	return append(l, item.(*ServiceAccountChange))
}

func (l ServiceAccountChangeList) AppendList(list gotenresource.ResourceChangeList) gotenresource.ResourceChangeList {
	return append(l, list.(ServiceAccountChangeList)...)
}

func (l ServiceAccountChangeList) Slice(first, second int) gotenresource.ResourceChangeList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l ServiceAccountChangeList) At(idx int) gotenresource.ResourceChange {
	return l[idx]
}

func (l ServiceAccountChangeList) Set(idx int, change gotenresource.ResourceChange) {
	l[idx] = change.(*ServiceAccountChange)
}

func (l ServiceAccountChangeList) Length() int {
	return len(l)
}

type ServiceAccountNameList []*Name

func (l ServiceAccountNameList) Append(name gotenresource.Name) gotenresource.NameList {
	return append(l, name.(*Name))
}

func (l ServiceAccountNameList) AppendList(list gotenresource.NameList) gotenresource.NameList {
	return append(l, list.(ServiceAccountNameList)...)
}

func (l ServiceAccountNameList) Slice(first, second int) gotenresource.NameList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l ServiceAccountNameList) At(idx int) gotenresource.Name {
	return l[idx]
}

func (l ServiceAccountNameList) Set(idx int, name gotenresource.Name) {
	l[idx] = name.(*Name)
}

func (l ServiceAccountNameList) Length() int {
	return len(l)
}

type ServiceAccountReferenceList []*Reference

func (l ServiceAccountReferenceList) Append(ref gotenresource.Reference) gotenresource.ReferenceList {
	return append(l, ref.(*Reference))
}

func (l ServiceAccountReferenceList) AppendList(list gotenresource.ReferenceList) gotenresource.ReferenceList {
	return append(l, list.(ServiceAccountReferenceList)...)
}

func (l ServiceAccountReferenceList) Slice(first, second int) gotenresource.ReferenceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l ServiceAccountReferenceList) At(idx int) gotenresource.Reference {
	return l[idx]
}

func (l ServiceAccountReferenceList) Set(idx int, ref gotenresource.Reference) {
	l[idx] = ref.(*Reference)
}

func (l ServiceAccountReferenceList) Length() int {
	return len(l)
}

type ServiceAccountParentNameList []*ParentName

func (l ServiceAccountParentNameList) Append(name gotenresource.Name) gotenresource.ParentNameList {
	return append(l, name.(*ParentName))
}

func (l ServiceAccountParentNameList) AppendList(list gotenresource.ParentNameList) gotenresource.ParentNameList {
	return append(l, list.(ServiceAccountParentNameList)...)
}

func (l ServiceAccountParentNameList) Slice(first, second int) gotenresource.ParentNameList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l ServiceAccountParentNameList) At(idx int) gotenresource.Name {
	return l[idx]
}

func (l ServiceAccountParentNameList) Set(idx int, name gotenresource.Name) {
	l[idx] = name.(*ParentName)
}

func (l ServiceAccountParentNameList) Length() int {
	return len(l)
}

type ServiceAccountParentReferenceList []*ParentReference

func (l ServiceAccountParentReferenceList) Append(ref gotenresource.Reference) gotenresource.ParentReferenceList {
	return append(l, ref.(*ParentReference))
}

func (l ServiceAccountParentReferenceList) AppendList(list gotenresource.ParentReferenceList) gotenresource.ParentReferenceList {
	return append(l, list.(ServiceAccountParentReferenceList)...)
}

func (l ServiceAccountParentReferenceList) Slice(first, second int) gotenresource.ParentReferenceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l ServiceAccountParentReferenceList) At(idx int) gotenresource.Reference {
	return l[idx]
}

func (l ServiceAccountParentReferenceList) Set(idx int, ref gotenresource.Reference) {
	l[idx] = ref.(*ParentReference)
}

func (l ServiceAccountParentReferenceList) Length() int {
	return len(l)
}

type ServiceAccountMap map[Name]*ServiceAccount

func (m ServiceAccountMap) Get(name gotenresource.Name) gotenresource.Resource {
	return m[*name.(*Name)]
}

func (m ServiceAccountMap) Set(res gotenresource.Resource) {
	tRes := res.(*ServiceAccount)
	m[*tRes.Name] = tRes
}

func (m ServiceAccountMap) Delete(name gotenresource.Name) {
	delete(m, *name.(*Name))
}

func (m ServiceAccountMap) Length() int {
	return len(m)
}

func (m ServiceAccountMap) ForEach(cb func(gotenresource.Name, gotenresource.Resource) bool) {
	for name, res := range m {
		if !cb(&name, res) {
			break
		}
	}
}

type ServiceAccountChangeMap map[Name]*ServiceAccountChange

func (m ServiceAccountChangeMap) Get(name gotenresource.Name) gotenresource.ResourceChange {
	return m[*name.(*Name)]
}

func (m ServiceAccountChangeMap) Set(change gotenresource.ResourceChange) {
	tChange := change.(*ServiceAccountChange)
	m[*tChange.GetServiceAccountName()] = tChange
}

func (m ServiceAccountChangeMap) Delete(name gotenresource.Name) {
	delete(m, *name.(*Name))
}

func (m ServiceAccountChangeMap) Length() int {
	return len(m)
}

func (m ServiceAccountChangeMap) ForEach(cb func(gotenresource.Name, gotenresource.ResourceChange) bool) {
	for name, res := range m {
		if !cb(&name, res) {
			break
		}
	}
}
