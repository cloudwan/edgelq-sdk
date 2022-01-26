// Code generated by protoc-gen-goten-resource
// Resource: ProjectInvitation
// DO NOT EDIT!!!

package project_invitation

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	iam_common "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/common"
	project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/project"
)

// ensure the imports are used
var (
	_ = gotenresource.ListQuery(nil)
)

// make sure we're using proto imports
var (
	_ = &iam_common.Actor{}
	_ = &project.Project{}
)

type ProjectInvitationList []*ProjectInvitation

func (l ProjectInvitationList) Append(item gotenresource.Resource) gotenresource.ResourceList {
	return append(l, item.(*ProjectInvitation))
}

func (l ProjectInvitationList) AppendList(list gotenresource.ResourceList) gotenresource.ResourceList {
	return append(l, list.(ProjectInvitationList)...)
}

func (l ProjectInvitationList) Slice(first, second int) gotenresource.ResourceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l ProjectInvitationList) At(idx int) gotenresource.Resource {
	return l[idx]
}

func (l ProjectInvitationList) Set(idx int, res gotenresource.Resource) {
	l[idx] = res.(*ProjectInvitation)
}

func (l ProjectInvitationList) Length() int {
	return len(l)
}

type ProjectInvitationChangeList []*ProjectInvitationChange

func (l ProjectInvitationChangeList) Append(item gotenresource.ResourceChange) gotenresource.ResourceChangeList {
	return append(l, item.(*ProjectInvitationChange))
}

func (l ProjectInvitationChangeList) AppendList(list gotenresource.ResourceChangeList) gotenresource.ResourceChangeList {
	return append(l, list.(ProjectInvitationChangeList)...)
}

func (l ProjectInvitationChangeList) Slice(first, second int) gotenresource.ResourceChangeList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l ProjectInvitationChangeList) At(idx int) gotenresource.ResourceChange {
	return l[idx]
}

func (l ProjectInvitationChangeList) Set(idx int, change gotenresource.ResourceChange) {
	l[idx] = change.(*ProjectInvitationChange)
}

func (l ProjectInvitationChangeList) Length() int {
	return len(l)
}

type ProjectInvitationNameList []*Name

func (l ProjectInvitationNameList) Append(name gotenresource.Name) gotenresource.NameList {
	return append(l, name.(*Name))
}

func (l ProjectInvitationNameList) AppendList(list gotenresource.NameList) gotenresource.NameList {
	return append(l, list.(ProjectInvitationNameList)...)
}

func (l ProjectInvitationNameList) Slice(first, second int) gotenresource.NameList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l ProjectInvitationNameList) At(idx int) gotenresource.Name {
	return l[idx]
}

func (l ProjectInvitationNameList) Set(idx int, name gotenresource.Name) {
	l[idx] = name.(*Name)
}

func (l ProjectInvitationNameList) Length() int {
	return len(l)
}

type ProjectInvitationReferenceList []*Reference

func (l ProjectInvitationReferenceList) Append(ref gotenresource.Reference) gotenresource.ReferenceList {
	return append(l, ref.(*Reference))
}

func (l ProjectInvitationReferenceList) AppendList(list gotenresource.ReferenceList) gotenresource.ReferenceList {
	return append(l, list.(ProjectInvitationReferenceList)...)
}

func (l ProjectInvitationReferenceList) Slice(first, second int) gotenresource.ReferenceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l ProjectInvitationReferenceList) At(idx int) gotenresource.Reference {
	return l[idx]
}

func (l ProjectInvitationReferenceList) Set(idx int, ref gotenresource.Reference) {
	l[idx] = ref.(*Reference)
}

func (l ProjectInvitationReferenceList) Length() int {
	return len(l)
}

type ProjectInvitationParentNameList []*ParentName

func (l ProjectInvitationParentNameList) Append(name gotenresource.Name) gotenresource.ParentNameList {
	return append(l, name.(*ParentName))
}

func (l ProjectInvitationParentNameList) AppendList(list gotenresource.ParentNameList) gotenresource.ParentNameList {
	return append(l, list.(ProjectInvitationParentNameList)...)
}

func (l ProjectInvitationParentNameList) Slice(first, second int) gotenresource.ParentNameList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l ProjectInvitationParentNameList) At(idx int) gotenresource.Name {
	return l[idx]
}

func (l ProjectInvitationParentNameList) Set(idx int, name gotenresource.Name) {
	l[idx] = name.(*ParentName)
}

func (l ProjectInvitationParentNameList) Length() int {
	return len(l)
}

type ProjectInvitationParentReferenceList []*ParentReference

func (l ProjectInvitationParentReferenceList) Append(ref gotenresource.Reference) gotenresource.ParentReferenceList {
	return append(l, ref.(*ParentReference))
}

func (l ProjectInvitationParentReferenceList) AppendList(list gotenresource.ParentReferenceList) gotenresource.ParentReferenceList {
	return append(l, list.(ProjectInvitationParentReferenceList)...)
}

func (l ProjectInvitationParentReferenceList) Slice(first, second int) gotenresource.ParentReferenceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l ProjectInvitationParentReferenceList) At(idx int) gotenresource.Reference {
	return l[idx]
}

func (l ProjectInvitationParentReferenceList) Set(idx int, ref gotenresource.Reference) {
	l[idx] = ref.(*ParentReference)
}

func (l ProjectInvitationParentReferenceList) Length() int {
	return len(l)
}

type ProjectInvitationMap map[Name]*ProjectInvitation

func (m ProjectInvitationMap) Get(name gotenresource.Name) gotenresource.Resource {
	return m[*name.(*Name)]
}

func (m ProjectInvitationMap) Set(res gotenresource.Resource) {
	tRes := res.(*ProjectInvitation)
	m[*tRes.Name] = tRes
}

func (m ProjectInvitationMap) Delete(name gotenresource.Name) {
	delete(m, *name.(*Name))
}

func (m ProjectInvitationMap) Length() int {
	return len(m)
}

func (m ProjectInvitationMap) ForEach(cb func(gotenresource.Name, gotenresource.Resource) bool) {
	for name, res := range m {
		if !cb(&name, res) {
			break
		}
	}
}

type ProjectInvitationChangeMap map[Name]*ProjectInvitationChange

func (m ProjectInvitationChangeMap) Get(name gotenresource.Name) gotenresource.ResourceChange {
	return m[*name.(*Name)]
}

func (m ProjectInvitationChangeMap) Set(change gotenresource.ResourceChange) {
	tChange := change.(*ProjectInvitationChange)
	m[*tChange.GetProjectInvitationName()] = tChange
}

func (m ProjectInvitationChangeMap) Delete(name gotenresource.Name) {
	delete(m, *name.(*Name))
}

func (m ProjectInvitationChangeMap) Length() int {
	return len(m)
}

func (m ProjectInvitationChangeMap) ForEach(cb func(gotenresource.Name, gotenresource.ResourceChange) bool) {
	for name, res := range m {
		if !cb(&name, res) {
			break
		}
	}
}
