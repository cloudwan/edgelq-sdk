// Code generated by protoc-gen-goten-resource
// Resource: MetricDescriptor
// DO NOT EDIT!!!

package metric_descriptor

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	api "github.com/cloudwan/edgelq-sdk/common/api"
	common "github.com/cloudwan/edgelq-sdk/monitoring/resources/v4/common"
	monitored_resource_descriptor "github.com/cloudwan/edgelq-sdk/monitoring/resources/v4/monitored_resource_descriptor"
	project "github.com/cloudwan/edgelq-sdk/monitoring/resources/v4/project"
	meta "github.com/cloudwan/goten-sdk/types/meta"
)

// ensure the imports are used
var (
	_ = new(gotenresource.ListQuery)
)

// make sure we're using proto imports
var (
	_ = api.LaunchStage(0)
	_ = &common.LabelDescriptor{}
	_ = &monitored_resource_descriptor.MonitoredResourceDescriptor{}
	_ = &project.Project{}
	_ = &meta.Meta{}
)

type MetricDescriptorList []*MetricDescriptor

func (l MetricDescriptorList) Append(item gotenresource.Resource) gotenresource.ResourceList {
	return append(l, item.(*MetricDescriptor))
}

func (l MetricDescriptorList) AppendList(list gotenresource.ResourceList) gotenresource.ResourceList {
	return append(l, list.(MetricDescriptorList)...)
}

func (l MetricDescriptorList) Slice(first, second int) gotenresource.ResourceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l MetricDescriptorList) At(idx int) gotenresource.Resource {
	return l[idx]
}

func (l MetricDescriptorList) Set(idx int, res gotenresource.Resource) {
	l[idx] = res.(*MetricDescriptor)
}

func (l MetricDescriptorList) Length() int {
	return len(l)
}

type MetricDescriptorChangeList []*MetricDescriptorChange

func (l MetricDescriptorChangeList) Append(item gotenresource.ResourceChange) gotenresource.ResourceChangeList {
	return append(l, item.(*MetricDescriptorChange))
}

func (l MetricDescriptorChangeList) AppendList(list gotenresource.ResourceChangeList) gotenresource.ResourceChangeList {
	return append(l, list.(MetricDescriptorChangeList)...)
}

func (l MetricDescriptorChangeList) Slice(first, second int) gotenresource.ResourceChangeList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l MetricDescriptorChangeList) At(idx int) gotenresource.ResourceChange {
	return l[idx]
}

func (l MetricDescriptorChangeList) Set(idx int, change gotenresource.ResourceChange) {
	l[idx] = change.(*MetricDescriptorChange)
}

func (l MetricDescriptorChangeList) Length() int {
	return len(l)
}

type MetricDescriptorNameList []*Name

func (l MetricDescriptorNameList) Append(name gotenresource.Name) gotenresource.NameList {
	return append(l, name.(*Name))
}

func (l MetricDescriptorNameList) AppendList(list gotenresource.NameList) gotenresource.NameList {
	return append(l, list.(MetricDescriptorNameList)...)
}

func (l MetricDescriptorNameList) Slice(first, second int) gotenresource.NameList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l MetricDescriptorNameList) At(idx int) gotenresource.Name {
	return l[idx]
}

func (l MetricDescriptorNameList) Set(idx int, name gotenresource.Name) {
	l[idx] = name.(*Name)
}

func (l MetricDescriptorNameList) Length() int {
	return len(l)
}

type MetricDescriptorReferenceList []*Reference

func (l MetricDescriptorReferenceList) Append(ref gotenresource.Reference) gotenresource.ReferenceList {
	return append(l, ref.(*Reference))
}

func (l MetricDescriptorReferenceList) AppendList(list gotenresource.ReferenceList) gotenresource.ReferenceList {
	return append(l, list.(MetricDescriptorReferenceList)...)
}

func (l MetricDescriptorReferenceList) Slice(first, second int) gotenresource.ReferenceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l MetricDescriptorReferenceList) At(idx int) gotenresource.Reference {
	return l[idx]
}

func (l MetricDescriptorReferenceList) Set(idx int, ref gotenresource.Reference) {
	l[idx] = ref.(*Reference)
}

func (l MetricDescriptorReferenceList) Length() int {
	return len(l)
}

type MetricDescriptorParentNameList []*ParentName

func (l MetricDescriptorParentNameList) Append(name gotenresource.Name) gotenresource.ParentNameList {
	return append(l, name.(*ParentName))
}

func (l MetricDescriptorParentNameList) AppendList(list gotenresource.ParentNameList) gotenresource.ParentNameList {
	return append(l, list.(MetricDescriptorParentNameList)...)
}

func (l MetricDescriptorParentNameList) Slice(first, second int) gotenresource.ParentNameList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l MetricDescriptorParentNameList) At(idx int) gotenresource.Name {
	return l[idx]
}

func (l MetricDescriptorParentNameList) Set(idx int, name gotenresource.Name) {
	l[idx] = name.(*ParentName)
}

func (l MetricDescriptorParentNameList) Length() int {
	return len(l)
}

type MetricDescriptorParentReferenceList []*ParentReference

func (l MetricDescriptorParentReferenceList) Append(ref gotenresource.Reference) gotenresource.ParentReferenceList {
	return append(l, ref.(*ParentReference))
}

func (l MetricDescriptorParentReferenceList) AppendList(list gotenresource.ParentReferenceList) gotenresource.ParentReferenceList {
	return append(l, list.(MetricDescriptorParentReferenceList)...)
}

func (l MetricDescriptorParentReferenceList) Slice(first, second int) gotenresource.ParentReferenceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l MetricDescriptorParentReferenceList) At(idx int) gotenresource.Reference {
	return l[idx]
}

func (l MetricDescriptorParentReferenceList) Set(idx int, ref gotenresource.Reference) {
	l[idx] = ref.(*ParentReference)
}

func (l MetricDescriptorParentReferenceList) Length() int {
	return len(l)
}

type MetricDescriptorMap map[Name]*MetricDescriptor

func (m MetricDescriptorMap) Get(name gotenresource.Name) gotenresource.Resource {
	return m[*name.(*Name)]
}

func (m MetricDescriptorMap) Set(res gotenresource.Resource) {
	tRes := res.(*MetricDescriptor)
	m[*tRes.Name] = tRes
}

func (m MetricDescriptorMap) Delete(name gotenresource.Name) {
	delete(m, *name.(*Name))
}

func (m MetricDescriptorMap) Length() int {
	return len(m)
}

func (m MetricDescriptorMap) ForEach(cb func(gotenresource.Name, gotenresource.Resource) bool) {
	for name, res := range m {
		if !cb(&name, res) {
			break
		}
	}
}

type MetricDescriptorChangeMap map[Name]*MetricDescriptorChange

func (m MetricDescriptorChangeMap) Get(name gotenresource.Name) gotenresource.ResourceChange {
	return m[*name.(*Name)]
}

func (m MetricDescriptorChangeMap) Set(change gotenresource.ResourceChange) {
	tChange := change.(*MetricDescriptorChange)
	m[*tChange.GetMetricDescriptorName()] = tChange
}

func (m MetricDescriptorChangeMap) Delete(name gotenresource.Name) {
	delete(m, *name.(*Name))
}

func (m MetricDescriptorChangeMap) Length() int {
	return len(m)
}

func (m MetricDescriptorChangeMap) ForEach(cb func(gotenresource.Name, gotenresource.ResourceChange) bool) {
	for name, res := range m {
		if !cb(&name, res) {
			break
		}
	}
}
