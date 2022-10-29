// Code generated by protoc-gen-goten-resource
// Resource: LogDescriptor
// DO NOT EDIT!!!

package log_descriptor

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
	iam_organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/organization"
	iam_project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/project"
	logging_common "github.com/cloudwan/edgelq-sdk/logging/common/v1alpha2"
)

// ensure the imports are used
var (
	_ = new(gotenresource.ListQuery)
)

// make sure we're using proto imports
var (
	_ = &ntt_meta.Meta{}
	_ = &iam_organization.Organization{}
	_ = &iam_project.Project{}
	_ = &logging_common.LabelDescriptor{}
)

type LogDescriptorList []*LogDescriptor

func (l LogDescriptorList) Append(item gotenresource.Resource) gotenresource.ResourceList {
	return append(l, item.(*LogDescriptor))
}

func (l LogDescriptorList) AppendList(list gotenresource.ResourceList) gotenresource.ResourceList {
	return append(l, list.(LogDescriptorList)...)
}

func (l LogDescriptorList) Slice(first, second int) gotenresource.ResourceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l LogDescriptorList) At(idx int) gotenresource.Resource {
	return l[idx]
}

func (l LogDescriptorList) Set(idx int, res gotenresource.Resource) {
	l[idx] = res.(*LogDescriptor)
}

func (l LogDescriptorList) Length() int {
	return len(l)
}

type LogDescriptorChangeList []*LogDescriptorChange

func (l LogDescriptorChangeList) Append(item gotenresource.ResourceChange) gotenresource.ResourceChangeList {
	return append(l, item.(*LogDescriptorChange))
}

func (l LogDescriptorChangeList) AppendList(list gotenresource.ResourceChangeList) gotenresource.ResourceChangeList {
	return append(l, list.(LogDescriptorChangeList)...)
}

func (l LogDescriptorChangeList) Slice(first, second int) gotenresource.ResourceChangeList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l LogDescriptorChangeList) At(idx int) gotenresource.ResourceChange {
	return l[idx]
}

func (l LogDescriptorChangeList) Set(idx int, change gotenresource.ResourceChange) {
	l[idx] = change.(*LogDescriptorChange)
}

func (l LogDescriptorChangeList) Length() int {
	return len(l)
}

type LogDescriptorNameList []*Name

func (l LogDescriptorNameList) Append(name gotenresource.Name) gotenresource.NameList {
	return append(l, name.(*Name))
}

func (l LogDescriptorNameList) AppendList(list gotenresource.NameList) gotenresource.NameList {
	return append(l, list.(LogDescriptorNameList)...)
}

func (l LogDescriptorNameList) Slice(first, second int) gotenresource.NameList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l LogDescriptorNameList) At(idx int) gotenresource.Name {
	return l[idx]
}

func (l LogDescriptorNameList) Set(idx int, name gotenresource.Name) {
	l[idx] = name.(*Name)
}

func (l LogDescriptorNameList) Length() int {
	return len(l)
}

type LogDescriptorReferenceList []*Reference

func (l LogDescriptorReferenceList) Append(ref gotenresource.Reference) gotenresource.ReferenceList {
	return append(l, ref.(*Reference))
}

func (l LogDescriptorReferenceList) AppendList(list gotenresource.ReferenceList) gotenresource.ReferenceList {
	return append(l, list.(LogDescriptorReferenceList)...)
}

func (l LogDescriptorReferenceList) Slice(first, second int) gotenresource.ReferenceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l LogDescriptorReferenceList) At(idx int) gotenresource.Reference {
	return l[idx]
}

func (l LogDescriptorReferenceList) Set(idx int, ref gotenresource.Reference) {
	l[idx] = ref.(*Reference)
}

func (l LogDescriptorReferenceList) Length() int {
	return len(l)
}

type LogDescriptorParentNameList []*ParentName

func (l LogDescriptorParentNameList) Append(name gotenresource.Name) gotenresource.ParentNameList {
	return append(l, name.(*ParentName))
}

func (l LogDescriptorParentNameList) AppendList(list gotenresource.ParentNameList) gotenresource.ParentNameList {
	return append(l, list.(LogDescriptorParentNameList)...)
}

func (l LogDescriptorParentNameList) Slice(first, second int) gotenresource.ParentNameList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l LogDescriptorParentNameList) At(idx int) gotenresource.Name {
	return l[idx]
}

func (l LogDescriptorParentNameList) Set(idx int, name gotenresource.Name) {
	l[idx] = name.(*ParentName)
}

func (l LogDescriptorParentNameList) Length() int {
	return len(l)
}

type LogDescriptorParentReferenceList []*ParentReference

func (l LogDescriptorParentReferenceList) Append(ref gotenresource.Reference) gotenresource.ParentReferenceList {
	return append(l, ref.(*ParentReference))
}

func (l LogDescriptorParentReferenceList) AppendList(list gotenresource.ParentReferenceList) gotenresource.ParentReferenceList {
	return append(l, list.(LogDescriptorParentReferenceList)...)
}

func (l LogDescriptorParentReferenceList) Slice(first, second int) gotenresource.ParentReferenceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l LogDescriptorParentReferenceList) At(idx int) gotenresource.Reference {
	return l[idx]
}

func (l LogDescriptorParentReferenceList) Set(idx int, ref gotenresource.Reference) {
	l[idx] = ref.(*ParentReference)
}

func (l LogDescriptorParentReferenceList) Length() int {
	return len(l)
}

type LogDescriptorMap map[Name]*LogDescriptor

func (m LogDescriptorMap) Get(name gotenresource.Name) gotenresource.Resource {
	return m[*name.(*Name)]
}

func (m LogDescriptorMap) Set(res gotenresource.Resource) {
	tRes := res.(*LogDescriptor)
	m[*tRes.Name] = tRes
}

func (m LogDescriptorMap) Delete(name gotenresource.Name) {
	delete(m, *name.(*Name))
}

func (m LogDescriptorMap) Length() int {
	return len(m)
}

func (m LogDescriptorMap) ForEach(cb func(gotenresource.Name, gotenresource.Resource) bool) {
	for name, res := range m {
		if !cb(&name, res) {
			break
		}
	}
}

type LogDescriptorChangeMap map[Name]*LogDescriptorChange

func (m LogDescriptorChangeMap) Get(name gotenresource.Name) gotenresource.ResourceChange {
	return m[*name.(*Name)]
}

func (m LogDescriptorChangeMap) Set(change gotenresource.ResourceChange) {
	tChange := change.(*LogDescriptorChange)
	m[*tChange.GetLogDescriptorName()] = tChange
}

func (m LogDescriptorChangeMap) Delete(name gotenresource.Name) {
	delete(m, *name.(*Name))
}

func (m LogDescriptorChangeMap) Length() int {
	return len(m)
}

func (m LogDescriptorChangeMap) ForEach(cb func(gotenresource.Name, gotenresource.ResourceChange) bool) {
	for name, res := range m {
		if !cb(&name, res) {
			break
		}
	}
}
