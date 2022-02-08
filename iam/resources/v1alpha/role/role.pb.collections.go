// Code generated by protoc-gen-goten-resource
// Resource: Role
// DO NOT EDIT!!!

package role

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	condition "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/condition"
	permission "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/permission"
)

// ensure the imports are used
var (
	_ = gotenresource.ListQuery(nil)
)

// make sure we're using proto imports
var (
	_ = &condition.Condition{}
	_ = &permission.Permission{}
)

type RoleList []*Role

func (l RoleList) Append(item gotenresource.Resource) gotenresource.ResourceList {
	return append(l, item.(*Role))
}

func (l RoleList) AppendList(list gotenresource.ResourceList) gotenresource.ResourceList {
	return append(l, list.(RoleList)...)
}

func (l RoleList) Slice(first, second int) gotenresource.ResourceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l RoleList) At(idx int) gotenresource.Resource {
	return l[idx]
}

func (l RoleList) Set(idx int, res gotenresource.Resource) {
	l[idx] = res.(*Role)
}

func (l RoleList) Length() int {
	return len(l)
}

type RoleChangeList []*RoleChange

func (l RoleChangeList) Append(item gotenresource.ResourceChange) gotenresource.ResourceChangeList {
	return append(l, item.(*RoleChange))
}

func (l RoleChangeList) AppendList(list gotenresource.ResourceChangeList) gotenresource.ResourceChangeList {
	return append(l, list.(RoleChangeList)...)
}

func (l RoleChangeList) Slice(first, second int) gotenresource.ResourceChangeList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l RoleChangeList) At(idx int) gotenresource.ResourceChange {
	return l[idx]
}

func (l RoleChangeList) Set(idx int, change gotenresource.ResourceChange) {
	l[idx] = change.(*RoleChange)
}

func (l RoleChangeList) Length() int {
	return len(l)
}

type RoleNameList []*Name

func (l RoleNameList) Append(name gotenresource.Name) gotenresource.NameList {
	return append(l, name.(*Name))
}

func (l RoleNameList) AppendList(list gotenresource.NameList) gotenresource.NameList {
	return append(l, list.(RoleNameList)...)
}

func (l RoleNameList) Slice(first, second int) gotenresource.NameList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l RoleNameList) At(idx int) gotenresource.Name {
	return l[idx]
}

func (l RoleNameList) Set(idx int, name gotenresource.Name) {
	l[idx] = name.(*Name)
}

func (l RoleNameList) Length() int {
	return len(l)
}

type RoleReferenceList []*Reference

func (l RoleReferenceList) Append(ref gotenresource.Reference) gotenresource.ReferenceList {
	return append(l, ref.(*Reference))
}

func (l RoleReferenceList) AppendList(list gotenresource.ReferenceList) gotenresource.ReferenceList {
	return append(l, list.(RoleReferenceList)...)
}

func (l RoleReferenceList) Slice(first, second int) gotenresource.ReferenceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l RoleReferenceList) At(idx int) gotenresource.Reference {
	return l[idx]
}

func (l RoleReferenceList) Set(idx int, ref gotenresource.Reference) {
	l[idx] = ref.(*Reference)
}

func (l RoleReferenceList) Length() int {
	return len(l)
}

type RoleMap map[Name]*Role

func (m RoleMap) Get(name gotenresource.Name) gotenresource.Resource {
	return m[*name.(*Name)]
}

func (m RoleMap) Set(res gotenresource.Resource) {
	tRes := res.(*Role)
	m[*tRes.Name] = tRes
}

func (m RoleMap) Delete(name gotenresource.Name) {
	delete(m, *name.(*Name))
}

func (m RoleMap) Length() int {
	return len(m)
}

func (m RoleMap) ForEach(cb func(gotenresource.Name, gotenresource.Resource) bool) {
	for name, res := range m {
		if !cb(&name, res) {
			break
		}
	}
}

type RoleChangeMap map[Name]*RoleChange

func (m RoleChangeMap) Get(name gotenresource.Name) gotenresource.ResourceChange {
	return m[*name.(*Name)]
}

func (m RoleChangeMap) Set(change gotenresource.ResourceChange) {
	tChange := change.(*RoleChange)
	m[*tChange.GetRoleName()] = tChange
}

func (m RoleChangeMap) Delete(name gotenresource.Name) {
	delete(m, *name.(*Name))
}

func (m RoleChangeMap) Length() int {
	return len(m)
}

func (m RoleChangeMap) ForEach(cb func(gotenresource.Name, gotenresource.ResourceChange) bool) {
	for name, res := range m {
		if !cb(&name, res) {
			break
		}
	}
}