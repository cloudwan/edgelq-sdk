// Code generated by protoc-gen-goten-resource
// Resource: RoleBinding
// DO NOT EDIT!!!

package role_binding

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	condition "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/condition"
	organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/organization"
	project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/project"
	role "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/role"
	meta "github.com/cloudwan/goten-sdk/types/meta"
)

// ensure the imports are used
var (
	_ = new(gotenresource.ListQuery)
)

// make sure we're using proto imports
var (
	_ = &condition.Condition{}
	_ = &organization.Organization{}
	_ = &project.Project{}
	_ = &role.Role{}
	_ = &meta.Meta{}
)

type RoleBindingList []*RoleBinding

func (l RoleBindingList) Append(item gotenresource.Resource) gotenresource.ResourceList {
	return append(l, item.(*RoleBinding))
}

func (l RoleBindingList) AppendList(list gotenresource.ResourceList) gotenresource.ResourceList {
	return append(l, list.(RoleBindingList)...)
}

func (l RoleBindingList) Slice(first, second int) gotenresource.ResourceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l RoleBindingList) At(idx int) gotenresource.Resource {
	return l[idx]
}

func (l RoleBindingList) Set(idx int, res gotenresource.Resource) {
	l[idx] = res.(*RoleBinding)
}

func (l RoleBindingList) Length() int {
	return len(l)
}

type RoleBindingChangeList []*RoleBindingChange

func (l RoleBindingChangeList) Append(item gotenresource.ResourceChange) gotenresource.ResourceChangeList {
	return append(l, item.(*RoleBindingChange))
}

func (l RoleBindingChangeList) AppendList(list gotenresource.ResourceChangeList) gotenresource.ResourceChangeList {
	return append(l, list.(RoleBindingChangeList)...)
}

func (l RoleBindingChangeList) Slice(first, second int) gotenresource.ResourceChangeList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l RoleBindingChangeList) At(idx int) gotenresource.ResourceChange {
	return l[idx]
}

func (l RoleBindingChangeList) Set(idx int, change gotenresource.ResourceChange) {
	l[idx] = change.(*RoleBindingChange)
}

func (l RoleBindingChangeList) Length() int {
	return len(l)
}

type RoleBindingNameList []*Name

func (l RoleBindingNameList) Append(name gotenresource.Name) gotenresource.NameList {
	return append(l, name.(*Name))
}

func (l RoleBindingNameList) AppendList(list gotenresource.NameList) gotenresource.NameList {
	return append(l, list.(RoleBindingNameList)...)
}

func (l RoleBindingNameList) Slice(first, second int) gotenresource.NameList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l RoleBindingNameList) At(idx int) gotenresource.Name {
	return l[idx]
}

func (l RoleBindingNameList) Set(idx int, name gotenresource.Name) {
	l[idx] = name.(*Name)
}

func (l RoleBindingNameList) Length() int {
	return len(l)
}

type RoleBindingReferenceList []*Reference

func (l RoleBindingReferenceList) Append(ref gotenresource.Reference) gotenresource.ReferenceList {
	return append(l, ref.(*Reference))
}

func (l RoleBindingReferenceList) AppendList(list gotenresource.ReferenceList) gotenresource.ReferenceList {
	return append(l, list.(RoleBindingReferenceList)...)
}

func (l RoleBindingReferenceList) Slice(first, second int) gotenresource.ReferenceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l RoleBindingReferenceList) At(idx int) gotenresource.Reference {
	return l[idx]
}

func (l RoleBindingReferenceList) Set(idx int, ref gotenresource.Reference) {
	l[idx] = ref.(*Reference)
}

func (l RoleBindingReferenceList) Length() int {
	return len(l)
}

type RoleBindingParentNameList []*ParentName

func (l RoleBindingParentNameList) Append(name gotenresource.Name) gotenresource.ParentNameList {
	return append(l, name.(*ParentName))
}

func (l RoleBindingParentNameList) AppendList(list gotenresource.ParentNameList) gotenresource.ParentNameList {
	return append(l, list.(RoleBindingParentNameList)...)
}

func (l RoleBindingParentNameList) Slice(first, second int) gotenresource.ParentNameList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l RoleBindingParentNameList) At(idx int) gotenresource.Name {
	return l[idx]
}

func (l RoleBindingParentNameList) Set(idx int, name gotenresource.Name) {
	l[idx] = name.(*ParentName)
}

func (l RoleBindingParentNameList) Length() int {
	return len(l)
}

type RoleBindingParentReferenceList []*ParentReference

func (l RoleBindingParentReferenceList) Append(ref gotenresource.Reference) gotenresource.ParentReferenceList {
	return append(l, ref.(*ParentReference))
}

func (l RoleBindingParentReferenceList) AppendList(list gotenresource.ParentReferenceList) gotenresource.ParentReferenceList {
	return append(l, list.(RoleBindingParentReferenceList)...)
}

func (l RoleBindingParentReferenceList) Slice(first, second int) gotenresource.ParentReferenceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l RoleBindingParentReferenceList) At(idx int) gotenresource.Reference {
	return l[idx]
}

func (l RoleBindingParentReferenceList) Set(idx int, ref gotenresource.Reference) {
	l[idx] = ref.(*ParentReference)
}

func (l RoleBindingParentReferenceList) Length() int {
	return len(l)
}

type RoleBindingMap map[Name]*RoleBinding

func (m RoleBindingMap) Get(name gotenresource.Name) gotenresource.Resource {
	return m[*name.(*Name)]
}

func (m RoleBindingMap) Set(res gotenresource.Resource) {
	tRes := res.(*RoleBinding)
	m[*tRes.Name] = tRes
}

func (m RoleBindingMap) Delete(name gotenresource.Name) {
	delete(m, *name.(*Name))
}

func (m RoleBindingMap) Length() int {
	return len(m)
}

func (m RoleBindingMap) ForEach(cb func(gotenresource.Name, gotenresource.Resource) bool) {
	for name, res := range m {
		if !cb(&name, res) {
			break
		}
	}
}

type RoleBindingChangeMap map[Name]*RoleBindingChange

func (m RoleBindingChangeMap) Get(name gotenresource.Name) gotenresource.ResourceChange {
	return m[*name.(*Name)]
}

func (m RoleBindingChangeMap) Set(change gotenresource.ResourceChange) {
	tChange := change.(*RoleBindingChange)
	m[*tChange.GetRoleBindingName()] = tChange
}

func (m RoleBindingChangeMap) Delete(name gotenresource.Name) {
	delete(m, *name.(*Name))
}

func (m RoleBindingChangeMap) Length() int {
	return len(m)
}

func (m RoleBindingChangeMap) ForEach(cb func(gotenresource.Name, gotenresource.ResourceChange) bool) {
	for name, res := range m {
		if !cb(&name, res) {
			break
		}
	}
}
