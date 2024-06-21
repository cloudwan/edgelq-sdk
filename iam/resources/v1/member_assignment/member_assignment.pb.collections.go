// Code generated by protoc-gen-goten-resource
// Resource: MemberAssignment
// DO NOT EDIT!!!

package member_assignment

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	iam_common "github.com/cloudwan/edgelq-sdk/iam/resources/v1/common"
	organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1/organization"
	role_binding "github.com/cloudwan/edgelq-sdk/iam/resources/v1/role_binding"
	meta_service "github.com/cloudwan/goten-sdk/meta-service/resources/v1/service"
	meta "github.com/cloudwan/goten-sdk/types/meta"
)

// ensure the imports are used
var (
	_ = new(gotenresource.ListQuery)
)

// make sure we're using proto imports
var (
	_ = &iam_common.PCR{}
	_ = &organization.Organization{}
	_ = &role_binding.RoleBinding{}
	_ = &meta_service.Service{}
	_ = &meta.Meta{}
)

type MemberAssignmentList []*MemberAssignment

func (l MemberAssignmentList) Append(item gotenresource.Resource) gotenresource.ResourceList {
	return append(l, item.(*MemberAssignment))
}

func (l MemberAssignmentList) AppendList(list gotenresource.ResourceList) gotenresource.ResourceList {
	return append(l, list.(MemberAssignmentList)...)
}

func (l MemberAssignmentList) Slice(first, second int) gotenresource.ResourceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l MemberAssignmentList) At(idx int) gotenresource.Resource {
	return l[idx]
}

func (l MemberAssignmentList) Set(idx int, res gotenresource.Resource) {
	l[idx] = res.(*MemberAssignment)
}

func (l MemberAssignmentList) Length() int {
	return len(l)
}

type MemberAssignmentChangeList []*MemberAssignmentChange

func (l MemberAssignmentChangeList) Append(item gotenresource.ResourceChange) gotenresource.ResourceChangeList {
	return append(l, item.(*MemberAssignmentChange))
}

func (l MemberAssignmentChangeList) AppendList(list gotenresource.ResourceChangeList) gotenresource.ResourceChangeList {
	return append(l, list.(MemberAssignmentChangeList)...)
}

func (l MemberAssignmentChangeList) Slice(first, second int) gotenresource.ResourceChangeList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l MemberAssignmentChangeList) At(idx int) gotenresource.ResourceChange {
	return l[idx]
}

func (l MemberAssignmentChangeList) Set(idx int, change gotenresource.ResourceChange) {
	l[idx] = change.(*MemberAssignmentChange)
}

func (l MemberAssignmentChangeList) Length() int {
	return len(l)
}

type MemberAssignmentNameList []*Name

func (l MemberAssignmentNameList) Append(name gotenresource.Name) gotenresource.NameList {
	return append(l, name.(*Name))
}

func (l MemberAssignmentNameList) AppendList(list gotenresource.NameList) gotenresource.NameList {
	return append(l, list.(MemberAssignmentNameList)...)
}

func (l MemberAssignmentNameList) Slice(first, second int) gotenresource.NameList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l MemberAssignmentNameList) At(idx int) gotenresource.Name {
	return l[idx]
}

func (l MemberAssignmentNameList) Set(idx int, name gotenresource.Name) {
	l[idx] = name.(*Name)
}

func (l MemberAssignmentNameList) Length() int {
	return len(l)
}

type MemberAssignmentReferenceList []*Reference

func (l MemberAssignmentReferenceList) Append(ref gotenresource.Reference) gotenresource.ReferenceList {
	return append(l, ref.(*Reference))
}

func (l MemberAssignmentReferenceList) AppendList(list gotenresource.ReferenceList) gotenresource.ReferenceList {
	return append(l, list.(MemberAssignmentReferenceList)...)
}

func (l MemberAssignmentReferenceList) Slice(first, second int) gotenresource.ReferenceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l MemberAssignmentReferenceList) At(idx int) gotenresource.Reference {
	return l[idx]
}

func (l MemberAssignmentReferenceList) Set(idx int, ref gotenresource.Reference) {
	l[idx] = ref.(*Reference)
}

func (l MemberAssignmentReferenceList) Length() int {
	return len(l)
}

type MemberAssignmentParentNameList []*ParentName

func (l MemberAssignmentParentNameList) Append(name gotenresource.Name) gotenresource.ParentNameList {
	return append(l, name.(*ParentName))
}

func (l MemberAssignmentParentNameList) AppendList(list gotenresource.ParentNameList) gotenresource.ParentNameList {
	return append(l, list.(MemberAssignmentParentNameList)...)
}

func (l MemberAssignmentParentNameList) Slice(first, second int) gotenresource.ParentNameList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l MemberAssignmentParentNameList) At(idx int) gotenresource.Name {
	return l[idx]
}

func (l MemberAssignmentParentNameList) Set(idx int, name gotenresource.Name) {
	l[idx] = name.(*ParentName)
}

func (l MemberAssignmentParentNameList) Length() int {
	return len(l)
}

type MemberAssignmentParentReferenceList []*ParentReference

func (l MemberAssignmentParentReferenceList) Append(ref gotenresource.Reference) gotenresource.ParentReferenceList {
	return append(l, ref.(*ParentReference))
}

func (l MemberAssignmentParentReferenceList) AppendList(list gotenresource.ParentReferenceList) gotenresource.ParentReferenceList {
	return append(l, list.(MemberAssignmentParentReferenceList)...)
}

func (l MemberAssignmentParentReferenceList) Slice(first, second int) gotenresource.ParentReferenceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l MemberAssignmentParentReferenceList) At(idx int) gotenresource.Reference {
	return l[idx]
}

func (l MemberAssignmentParentReferenceList) Set(idx int, ref gotenresource.Reference) {
	l[idx] = ref.(*ParentReference)
}

func (l MemberAssignmentParentReferenceList) Length() int {
	return len(l)
}

type MemberAssignmentMap map[Name]*MemberAssignment

func (m MemberAssignmentMap) Get(name gotenresource.Name) gotenresource.Resource {
	return m[*name.(*Name)]
}

func (m MemberAssignmentMap) Set(res gotenresource.Resource) {
	tRes := res.(*MemberAssignment)
	m[*tRes.Name] = tRes
}

func (m MemberAssignmentMap) Delete(name gotenresource.Name) {
	delete(m, *name.(*Name))
}

func (m MemberAssignmentMap) Length() int {
	return len(m)
}

func (m MemberAssignmentMap) ForEach(cb func(gotenresource.Name, gotenresource.Resource) bool) {
	for name, res := range m {
		if !cb(&name, res) {
			break
		}
	}
}

type MemberAssignmentChangeMap map[Name]*MemberAssignmentChange

func (m MemberAssignmentChangeMap) Get(name gotenresource.Name) gotenresource.ResourceChange {
	return m[*name.(*Name)]
}

func (m MemberAssignmentChangeMap) Set(change gotenresource.ResourceChange) {
	tChange := change.(*MemberAssignmentChange)
	m[*tChange.GetMemberAssignmentName()] = tChange
}

func (m MemberAssignmentChangeMap) Delete(name gotenresource.Name) {
	delete(m, *name.(*Name))
}

func (m MemberAssignmentChangeMap) Length() int {
	return len(m)
}

func (m MemberAssignmentChangeMap) ForEach(cb func(gotenresource.Name, gotenresource.ResourceChange) bool) {
	for name, res := range m {
		if !cb(&name, res) {
			break
		}
	}
}
