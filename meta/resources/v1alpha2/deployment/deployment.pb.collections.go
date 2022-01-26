// Code generated by protoc-gen-goten-resource
// Resource: Deployment
// DO NOT EDIT!!!

package deployment

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
	region "github.com/cloudwan/edgelq-sdk/meta/resources/v1alpha2/region"
	service "github.com/cloudwan/edgelq-sdk/meta/resources/v1alpha2/service"
)

// ensure the imports are used
var (
	_ = gotenresource.ListQuery(nil)
)

// make sure we're using proto imports
var (
	_ = &ntt_meta.Meta{}
	_ = &region.Region{}
	_ = &service.Service{}
)

type DeploymentList []*Deployment

func (l DeploymentList) Append(item gotenresource.Resource) gotenresource.ResourceList {
	return append(l, item.(*Deployment))
}

func (l DeploymentList) AppendList(list gotenresource.ResourceList) gotenresource.ResourceList {
	return append(l, list.(DeploymentList)...)
}

func (l DeploymentList) Slice(first, second int) gotenresource.ResourceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l DeploymentList) At(idx int) gotenresource.Resource {
	return l[idx]
}

func (l DeploymentList) Set(idx int, res gotenresource.Resource) {
	l[idx] = res.(*Deployment)
}

func (l DeploymentList) Length() int {
	return len(l)
}

type DeploymentChangeList []*DeploymentChange

func (l DeploymentChangeList) Append(item gotenresource.ResourceChange) gotenresource.ResourceChangeList {
	return append(l, item.(*DeploymentChange))
}

func (l DeploymentChangeList) AppendList(list gotenresource.ResourceChangeList) gotenresource.ResourceChangeList {
	return append(l, list.(DeploymentChangeList)...)
}

func (l DeploymentChangeList) Slice(first, second int) gotenresource.ResourceChangeList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l DeploymentChangeList) At(idx int) gotenresource.ResourceChange {
	return l[idx]
}

func (l DeploymentChangeList) Set(idx int, change gotenresource.ResourceChange) {
	l[idx] = change.(*DeploymentChange)
}

func (l DeploymentChangeList) Length() int {
	return len(l)
}

type DeploymentNameList []*Name

func (l DeploymentNameList) Append(name gotenresource.Name) gotenresource.NameList {
	return append(l, name.(*Name))
}

func (l DeploymentNameList) AppendList(list gotenresource.NameList) gotenresource.NameList {
	return append(l, list.(DeploymentNameList)...)
}

func (l DeploymentNameList) Slice(first, second int) gotenresource.NameList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l DeploymentNameList) At(idx int) gotenresource.Name {
	return l[idx]
}

func (l DeploymentNameList) Set(idx int, name gotenresource.Name) {
	l[idx] = name.(*Name)
}

func (l DeploymentNameList) Length() int {
	return len(l)
}

type DeploymentReferenceList []*Reference

func (l DeploymentReferenceList) Append(ref gotenresource.Reference) gotenresource.ReferenceList {
	return append(l, ref.(*Reference))
}

func (l DeploymentReferenceList) AppendList(list gotenresource.ReferenceList) gotenresource.ReferenceList {
	return append(l, list.(DeploymentReferenceList)...)
}

func (l DeploymentReferenceList) Slice(first, second int) gotenresource.ReferenceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l DeploymentReferenceList) At(idx int) gotenresource.Reference {
	return l[idx]
}

func (l DeploymentReferenceList) Set(idx int, ref gotenresource.Reference) {
	l[idx] = ref.(*Reference)
}

func (l DeploymentReferenceList) Length() int {
	return len(l)
}

type DeploymentParentNameList []*ParentName

func (l DeploymentParentNameList) Append(name gotenresource.Name) gotenresource.ParentNameList {
	return append(l, name.(*ParentName))
}

func (l DeploymentParentNameList) AppendList(list gotenresource.ParentNameList) gotenresource.ParentNameList {
	return append(l, list.(DeploymentParentNameList)...)
}

func (l DeploymentParentNameList) Slice(first, second int) gotenresource.ParentNameList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l DeploymentParentNameList) At(idx int) gotenresource.Name {
	return l[idx]
}

func (l DeploymentParentNameList) Set(idx int, name gotenresource.Name) {
	l[idx] = name.(*ParentName)
}

func (l DeploymentParentNameList) Length() int {
	return len(l)
}

type DeploymentParentReferenceList []*ParentReference

func (l DeploymentParentReferenceList) Append(ref gotenresource.Reference) gotenresource.ParentReferenceList {
	return append(l, ref.(*ParentReference))
}

func (l DeploymentParentReferenceList) AppendList(list gotenresource.ParentReferenceList) gotenresource.ParentReferenceList {
	return append(l, list.(DeploymentParentReferenceList)...)
}

func (l DeploymentParentReferenceList) Slice(first, second int) gotenresource.ParentReferenceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l DeploymentParentReferenceList) At(idx int) gotenresource.Reference {
	return l[idx]
}

func (l DeploymentParentReferenceList) Set(idx int, ref gotenresource.Reference) {
	l[idx] = ref.(*ParentReference)
}

func (l DeploymentParentReferenceList) Length() int {
	return len(l)
}

type DeploymentMap map[Name]*Deployment

func (m DeploymentMap) Get(name gotenresource.Name) gotenresource.Resource {
	return m[*name.(*Name)]
}

func (m DeploymentMap) Set(res gotenresource.Resource) {
	tRes := res.(*Deployment)
	m[*tRes.Name] = tRes
}

func (m DeploymentMap) Delete(name gotenresource.Name) {
	delete(m, *name.(*Name))
}

func (m DeploymentMap) Length() int {
	return len(m)
}

func (m DeploymentMap) ForEach(cb func(gotenresource.Name, gotenresource.Resource) bool) {
	for name, res := range m {
		if !cb(&name, res) {
			break
		}
	}
}

type DeploymentChangeMap map[Name]*DeploymentChange

func (m DeploymentChangeMap) Get(name gotenresource.Name) gotenresource.ResourceChange {
	return m[*name.(*Name)]
}

func (m DeploymentChangeMap) Set(change gotenresource.ResourceChange) {
	tChange := change.(*DeploymentChange)
	m[*tChange.GetDeploymentName()] = tChange
}

func (m DeploymentChangeMap) Delete(name gotenresource.Name) {
	delete(m, *name.(*Name))
}

func (m DeploymentChangeMap) Length() int {
	return len(m)
}

func (m DeploymentChangeMap) ForEach(cb func(gotenresource.Name, gotenresource.ResourceChange) bool) {
	for name, res := range m {
		if !cb(&name, res) {
			break
		}
	}
}
