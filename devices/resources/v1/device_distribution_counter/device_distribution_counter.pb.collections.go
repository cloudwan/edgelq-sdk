// Code generated by protoc-gen-goten-resource
// Resource: DeviceDistributionCounter
// DO NOT EDIT!!!

package device_distribution_counter

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	project "github.com/cloudwan/edgelq-sdk/devices/resources/v1/project"
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

type DeviceDistributionCounterList []*DeviceDistributionCounter

func (l DeviceDistributionCounterList) Append(item gotenresource.Resource) gotenresource.ResourceList {
	return append(l, item.(*DeviceDistributionCounter))
}

func (l DeviceDistributionCounterList) AppendList(list gotenresource.ResourceList) gotenresource.ResourceList {
	return append(l, list.(DeviceDistributionCounterList)...)
}

func (l DeviceDistributionCounterList) Slice(first, second int) gotenresource.ResourceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l DeviceDistributionCounterList) At(idx int) gotenresource.Resource {
	return l[idx]
}

func (l DeviceDistributionCounterList) Set(idx int, res gotenresource.Resource) {
	l[idx] = res.(*DeviceDistributionCounter)
}

func (l DeviceDistributionCounterList) Length() int {
	return len(l)
}

type DeviceDistributionCounterChangeList []*DeviceDistributionCounterChange

func (l DeviceDistributionCounterChangeList) Append(item gotenresource.ResourceChange) gotenresource.ResourceChangeList {
	return append(l, item.(*DeviceDistributionCounterChange))
}

func (l DeviceDistributionCounterChangeList) AppendList(list gotenresource.ResourceChangeList) gotenresource.ResourceChangeList {
	return append(l, list.(DeviceDistributionCounterChangeList)...)
}

func (l DeviceDistributionCounterChangeList) Slice(first, second int) gotenresource.ResourceChangeList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l DeviceDistributionCounterChangeList) At(idx int) gotenresource.ResourceChange {
	return l[idx]
}

func (l DeviceDistributionCounterChangeList) Set(idx int, change gotenresource.ResourceChange) {
	l[idx] = change.(*DeviceDistributionCounterChange)
}

func (l DeviceDistributionCounterChangeList) Length() int {
	return len(l)
}

type DeviceDistributionCounterNameList []*Name

func (l DeviceDistributionCounterNameList) Append(name gotenresource.Name) gotenresource.NameList {
	return append(l, name.(*Name))
}

func (l DeviceDistributionCounterNameList) AppendList(list gotenresource.NameList) gotenresource.NameList {
	return append(l, list.(DeviceDistributionCounterNameList)...)
}

func (l DeviceDistributionCounterNameList) Slice(first, second int) gotenresource.NameList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l DeviceDistributionCounterNameList) At(idx int) gotenresource.Name {
	return l[idx]
}

func (l DeviceDistributionCounterNameList) Set(idx int, name gotenresource.Name) {
	l[idx] = name.(*Name)
}

func (l DeviceDistributionCounterNameList) Length() int {
	return len(l)
}

type DeviceDistributionCounterReferenceList []*Reference

func (l DeviceDistributionCounterReferenceList) Append(ref gotenresource.Reference) gotenresource.ReferenceList {
	return append(l, ref.(*Reference))
}

func (l DeviceDistributionCounterReferenceList) AppendList(list gotenresource.ReferenceList) gotenresource.ReferenceList {
	return append(l, list.(DeviceDistributionCounterReferenceList)...)
}

func (l DeviceDistributionCounterReferenceList) Slice(first, second int) gotenresource.ReferenceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l DeviceDistributionCounterReferenceList) At(idx int) gotenresource.Reference {
	return l[idx]
}

func (l DeviceDistributionCounterReferenceList) Set(idx int, ref gotenresource.Reference) {
	l[idx] = ref.(*Reference)
}

func (l DeviceDistributionCounterReferenceList) Length() int {
	return len(l)
}

type DeviceDistributionCounterParentNameList []*ParentName

func (l DeviceDistributionCounterParentNameList) Append(name gotenresource.Name) gotenresource.ParentNameList {
	return append(l, name.(*ParentName))
}

func (l DeviceDistributionCounterParentNameList) AppendList(list gotenresource.ParentNameList) gotenresource.ParentNameList {
	return append(l, list.(DeviceDistributionCounterParentNameList)...)
}

func (l DeviceDistributionCounterParentNameList) Slice(first, second int) gotenresource.ParentNameList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l DeviceDistributionCounterParentNameList) At(idx int) gotenresource.Name {
	return l[idx]
}

func (l DeviceDistributionCounterParentNameList) Set(idx int, name gotenresource.Name) {
	l[idx] = name.(*ParentName)
}

func (l DeviceDistributionCounterParentNameList) Length() int {
	return len(l)
}

type DeviceDistributionCounterParentReferenceList []*ParentReference

func (l DeviceDistributionCounterParentReferenceList) Append(ref gotenresource.Reference) gotenresource.ParentReferenceList {
	return append(l, ref.(*ParentReference))
}

func (l DeviceDistributionCounterParentReferenceList) AppendList(list gotenresource.ParentReferenceList) gotenresource.ParentReferenceList {
	return append(l, list.(DeviceDistributionCounterParentReferenceList)...)
}

func (l DeviceDistributionCounterParentReferenceList) Slice(first, second int) gotenresource.ParentReferenceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l DeviceDistributionCounterParentReferenceList) At(idx int) gotenresource.Reference {
	return l[idx]
}

func (l DeviceDistributionCounterParentReferenceList) Set(idx int, ref gotenresource.Reference) {
	l[idx] = ref.(*ParentReference)
}

func (l DeviceDistributionCounterParentReferenceList) Length() int {
	return len(l)
}

type DeviceDistributionCounterMap map[Name]*DeviceDistributionCounter

func (m DeviceDistributionCounterMap) Get(name gotenresource.Name) gotenresource.Resource {
	return m[*name.(*Name)]
}

func (m DeviceDistributionCounterMap) Set(res gotenresource.Resource) {
	tRes := res.(*DeviceDistributionCounter)
	m[*tRes.Name] = tRes
}

func (m DeviceDistributionCounterMap) Delete(name gotenresource.Name) {
	delete(m, *name.(*Name))
}

func (m DeviceDistributionCounterMap) Length() int {
	return len(m)
}

func (m DeviceDistributionCounterMap) ForEach(cb func(gotenresource.Name, gotenresource.Resource) bool) {
	for name, res := range m {
		if !cb(&name, res) {
			break
		}
	}
}

type DeviceDistributionCounterChangeMap map[Name]*DeviceDistributionCounterChange

func (m DeviceDistributionCounterChangeMap) Get(name gotenresource.Name) gotenresource.ResourceChange {
	return m[*name.(*Name)]
}

func (m DeviceDistributionCounterChangeMap) Set(change gotenresource.ResourceChange) {
	tChange := change.(*DeviceDistributionCounterChange)
	m[*tChange.GetDeviceDistributionCounterName()] = tChange
}

func (m DeviceDistributionCounterChangeMap) Delete(name gotenresource.Name) {
	delete(m, *name.(*Name))
}

func (m DeviceDistributionCounterChangeMap) Length() int {
	return len(m)
}

func (m DeviceDistributionCounterChangeMap) ForEach(cb func(gotenresource.Name, gotenresource.ResourceChange) bool) {
	for name, res := range m {
		if !cb(&name, res) {
			break
		}
	}
}
