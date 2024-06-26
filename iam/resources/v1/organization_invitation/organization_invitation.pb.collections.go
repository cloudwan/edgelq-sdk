// Code generated by protoc-gen-goten-resource
// Resource: OrganizationInvitation
// DO NOT EDIT!!!

package organization_invitation

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	iam_invitation "github.com/cloudwan/edgelq-sdk/iam/resources/v1/invitation"
	organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1/organization"
	meta "github.com/cloudwan/goten-sdk/types/meta"
)

// ensure the imports are used
var (
	_ = new(gotenresource.ListQuery)
)

// make sure we're using proto imports
var (
	_ = &iam_invitation.Actor{}
	_ = &organization.Organization{}
	_ = &meta.Meta{}
)

type OrganizationInvitationList []*OrganizationInvitation

func (l OrganizationInvitationList) Append(item gotenresource.Resource) gotenresource.ResourceList {
	return append(l, item.(*OrganizationInvitation))
}

func (l OrganizationInvitationList) AppendList(list gotenresource.ResourceList) gotenresource.ResourceList {
	return append(l, list.(OrganizationInvitationList)...)
}

func (l OrganizationInvitationList) Slice(first, second int) gotenresource.ResourceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l OrganizationInvitationList) At(idx int) gotenresource.Resource {
	return l[idx]
}

func (l OrganizationInvitationList) Set(idx int, res gotenresource.Resource) {
	l[idx] = res.(*OrganizationInvitation)
}

func (l OrganizationInvitationList) Length() int {
	return len(l)
}

type OrganizationInvitationChangeList []*OrganizationInvitationChange

func (l OrganizationInvitationChangeList) Append(item gotenresource.ResourceChange) gotenresource.ResourceChangeList {
	return append(l, item.(*OrganizationInvitationChange))
}

func (l OrganizationInvitationChangeList) AppendList(list gotenresource.ResourceChangeList) gotenresource.ResourceChangeList {
	return append(l, list.(OrganizationInvitationChangeList)...)
}

func (l OrganizationInvitationChangeList) Slice(first, second int) gotenresource.ResourceChangeList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l OrganizationInvitationChangeList) At(idx int) gotenresource.ResourceChange {
	return l[idx]
}

func (l OrganizationInvitationChangeList) Set(idx int, change gotenresource.ResourceChange) {
	l[idx] = change.(*OrganizationInvitationChange)
}

func (l OrganizationInvitationChangeList) Length() int {
	return len(l)
}

type OrganizationInvitationNameList []*Name

func (l OrganizationInvitationNameList) Append(name gotenresource.Name) gotenresource.NameList {
	return append(l, name.(*Name))
}

func (l OrganizationInvitationNameList) AppendList(list gotenresource.NameList) gotenresource.NameList {
	return append(l, list.(OrganizationInvitationNameList)...)
}

func (l OrganizationInvitationNameList) Slice(first, second int) gotenresource.NameList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l OrganizationInvitationNameList) At(idx int) gotenresource.Name {
	return l[idx]
}

func (l OrganizationInvitationNameList) Set(idx int, name gotenresource.Name) {
	l[idx] = name.(*Name)
}

func (l OrganizationInvitationNameList) Length() int {
	return len(l)
}

type OrganizationInvitationReferenceList []*Reference

func (l OrganizationInvitationReferenceList) Append(ref gotenresource.Reference) gotenresource.ReferenceList {
	return append(l, ref.(*Reference))
}

func (l OrganizationInvitationReferenceList) AppendList(list gotenresource.ReferenceList) gotenresource.ReferenceList {
	return append(l, list.(OrganizationInvitationReferenceList)...)
}

func (l OrganizationInvitationReferenceList) Slice(first, second int) gotenresource.ReferenceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l OrganizationInvitationReferenceList) At(idx int) gotenresource.Reference {
	return l[idx]
}

func (l OrganizationInvitationReferenceList) Set(idx int, ref gotenresource.Reference) {
	l[idx] = ref.(*Reference)
}

func (l OrganizationInvitationReferenceList) Length() int {
	return len(l)
}

type OrganizationInvitationParentNameList []*ParentName

func (l OrganizationInvitationParentNameList) Append(name gotenresource.Name) gotenresource.ParentNameList {
	return append(l, name.(*ParentName))
}

func (l OrganizationInvitationParentNameList) AppendList(list gotenresource.ParentNameList) gotenresource.ParentNameList {
	return append(l, list.(OrganizationInvitationParentNameList)...)
}

func (l OrganizationInvitationParentNameList) Slice(first, second int) gotenresource.ParentNameList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l OrganizationInvitationParentNameList) At(idx int) gotenresource.Name {
	return l[idx]
}

func (l OrganizationInvitationParentNameList) Set(idx int, name gotenresource.Name) {
	l[idx] = name.(*ParentName)
}

func (l OrganizationInvitationParentNameList) Length() int {
	return len(l)
}

type OrganizationInvitationParentReferenceList []*ParentReference

func (l OrganizationInvitationParentReferenceList) Append(ref gotenresource.Reference) gotenresource.ParentReferenceList {
	return append(l, ref.(*ParentReference))
}

func (l OrganizationInvitationParentReferenceList) AppendList(list gotenresource.ParentReferenceList) gotenresource.ParentReferenceList {
	return append(l, list.(OrganizationInvitationParentReferenceList)...)
}

func (l OrganizationInvitationParentReferenceList) Slice(first, second int) gotenresource.ParentReferenceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l OrganizationInvitationParentReferenceList) At(idx int) gotenresource.Reference {
	return l[idx]
}

func (l OrganizationInvitationParentReferenceList) Set(idx int, ref gotenresource.Reference) {
	l[idx] = ref.(*ParentReference)
}

func (l OrganizationInvitationParentReferenceList) Length() int {
	return len(l)
}

type OrganizationInvitationMap map[Name]*OrganizationInvitation

func (m OrganizationInvitationMap) Get(name gotenresource.Name) gotenresource.Resource {
	return m[*name.(*Name)]
}

func (m OrganizationInvitationMap) Set(res gotenresource.Resource) {
	tRes := res.(*OrganizationInvitation)
	m[*tRes.Name] = tRes
}

func (m OrganizationInvitationMap) Delete(name gotenresource.Name) {
	delete(m, *name.(*Name))
}

func (m OrganizationInvitationMap) Length() int {
	return len(m)
}

func (m OrganizationInvitationMap) ForEach(cb func(gotenresource.Name, gotenresource.Resource) bool) {
	for name, res := range m {
		if !cb(&name, res) {
			break
		}
	}
}

type OrganizationInvitationChangeMap map[Name]*OrganizationInvitationChange

func (m OrganizationInvitationChangeMap) Get(name gotenresource.Name) gotenresource.ResourceChange {
	return m[*name.(*Name)]
}

func (m OrganizationInvitationChangeMap) Set(change gotenresource.ResourceChange) {
	tChange := change.(*OrganizationInvitationChange)
	m[*tChange.GetOrganizationInvitationName()] = tChange
}

func (m OrganizationInvitationChangeMap) Delete(name gotenresource.Name) {
	delete(m, *name.(*Name))
}

func (m OrganizationInvitationChangeMap) Length() int {
	return len(m)
}

func (m OrganizationInvitationChangeMap) ForEach(cb func(gotenresource.Name, gotenresource.ResourceChange) bool) {
	for name, res := range m {
		if !cb(&name, res) {
			break
		}
	}
}
