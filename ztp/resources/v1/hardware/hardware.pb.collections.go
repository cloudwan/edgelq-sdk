// Code generated by protoc-gen-goten-resource
// Resource: Hardware
// DO NOT EDIT!!!

package hardware

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	edgelq_instance "github.com/cloudwan/edgelq-sdk/ztp/resources/v1/edgelq_instance"
	project "github.com/cloudwan/edgelq-sdk/ztp/resources/v1/project"
	meta "github.com/cloudwan/goten-sdk/types/meta"
)

// ensure the imports are used
var (
	_ = new(gotenresource.ListQuery)
)

// make sure we're using proto imports
var (
	_ = &edgelq_instance.EdgelqInstance{}
	_ = &project.Project{}
	_ = &meta.Meta{}
)

type HardwareList []*Hardware

func (l HardwareList) Append(item gotenresource.Resource) gotenresource.ResourceList {
	return append(l, item.(*Hardware))
}

func (l HardwareList) AppendList(list gotenresource.ResourceList) gotenresource.ResourceList {
	return append(l, list.(HardwareList)...)
}

func (l HardwareList) Slice(first, second int) gotenresource.ResourceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l HardwareList) At(idx int) gotenresource.Resource {
	return l[idx]
}

func (l HardwareList) Set(idx int, res gotenresource.Resource) {
	l[idx] = res.(*Hardware)
}

func (l HardwareList) Length() int {
	return len(l)
}

type HardwareChangeList []*HardwareChange

func (l HardwareChangeList) Append(item gotenresource.ResourceChange) gotenresource.ResourceChangeList {
	return append(l, item.(*HardwareChange))
}

func (l HardwareChangeList) AppendList(list gotenresource.ResourceChangeList) gotenresource.ResourceChangeList {
	return append(l, list.(HardwareChangeList)...)
}

func (l HardwareChangeList) Slice(first, second int) gotenresource.ResourceChangeList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l HardwareChangeList) At(idx int) gotenresource.ResourceChange {
	return l[idx]
}

func (l HardwareChangeList) Set(idx int, change gotenresource.ResourceChange) {
	l[idx] = change.(*HardwareChange)
}

func (l HardwareChangeList) Length() int {
	return len(l)
}

type HardwareNameList []*Name

func (l HardwareNameList) Append(name gotenresource.Name) gotenresource.NameList {
	return append(l, name.(*Name))
}

func (l HardwareNameList) AppendList(list gotenresource.NameList) gotenresource.NameList {
	return append(l, list.(HardwareNameList)...)
}

func (l HardwareNameList) Slice(first, second int) gotenresource.NameList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l HardwareNameList) At(idx int) gotenresource.Name {
	return l[idx]
}

func (l HardwareNameList) Set(idx int, name gotenresource.Name) {
	l[idx] = name.(*Name)
}

func (l HardwareNameList) Length() int {
	return len(l)
}

type HardwareReferenceList []*Reference

func (l HardwareReferenceList) Append(ref gotenresource.Reference) gotenresource.ReferenceList {
	return append(l, ref.(*Reference))
}

func (l HardwareReferenceList) AppendList(list gotenresource.ReferenceList) gotenresource.ReferenceList {
	return append(l, list.(HardwareReferenceList)...)
}

func (l HardwareReferenceList) Slice(first, second int) gotenresource.ReferenceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l HardwareReferenceList) At(idx int) gotenresource.Reference {
	return l[idx]
}

func (l HardwareReferenceList) Set(idx int, ref gotenresource.Reference) {
	l[idx] = ref.(*Reference)
}

func (l HardwareReferenceList) Length() int {
	return len(l)
}

type HardwareParentNameList []*ParentName

func (l HardwareParentNameList) Append(name gotenresource.Name) gotenresource.ParentNameList {
	return append(l, name.(*ParentName))
}

func (l HardwareParentNameList) AppendList(list gotenresource.ParentNameList) gotenresource.ParentNameList {
	return append(l, list.(HardwareParentNameList)...)
}

func (l HardwareParentNameList) Slice(first, second int) gotenresource.ParentNameList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l HardwareParentNameList) At(idx int) gotenresource.Name {
	return l[idx]
}

func (l HardwareParentNameList) Set(idx int, name gotenresource.Name) {
	l[idx] = name.(*ParentName)
}

func (l HardwareParentNameList) Length() int {
	return len(l)
}

type HardwareParentReferenceList []*ParentReference

func (l HardwareParentReferenceList) Append(ref gotenresource.Reference) gotenresource.ParentReferenceList {
	return append(l, ref.(*ParentReference))
}

func (l HardwareParentReferenceList) AppendList(list gotenresource.ParentReferenceList) gotenresource.ParentReferenceList {
	return append(l, list.(HardwareParentReferenceList)...)
}

func (l HardwareParentReferenceList) Slice(first, second int) gotenresource.ParentReferenceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l HardwareParentReferenceList) At(idx int) gotenresource.Reference {
	return l[idx]
}

func (l HardwareParentReferenceList) Set(idx int, ref gotenresource.Reference) {
	l[idx] = ref.(*ParentReference)
}

func (l HardwareParentReferenceList) Length() int {
	return len(l)
}

type HardwareMap map[Name]*Hardware

func (m HardwareMap) Get(name gotenresource.Name) gotenresource.Resource {
	return m[*name.(*Name)]
}

func (m HardwareMap) Set(res gotenresource.Resource) {
	tRes := res.(*Hardware)
	m[*tRes.Name] = tRes
}

func (m HardwareMap) Delete(name gotenresource.Name) {
	delete(m, *name.(*Name))
}

func (m HardwareMap) Length() int {
	return len(m)
}

func (m HardwareMap) ForEach(cb func(gotenresource.Name, gotenresource.Resource) bool) {
	for name, res := range m {
		if !cb(&name, res) {
			break
		}
	}
}

type HardwareChangeMap map[Name]*HardwareChange

func (m HardwareChangeMap) Get(name gotenresource.Name) gotenresource.ResourceChange {
	return m[*name.(*Name)]
}

func (m HardwareChangeMap) Set(change gotenresource.ResourceChange) {
	tChange := change.(*HardwareChange)
	m[*tChange.GetHardwareName()] = tChange
}

func (m HardwareChangeMap) Delete(name gotenresource.Name) {
	delete(m, *name.(*Name))
}

func (m HardwareChangeMap) Length() int {
	return len(m)
}

func (m HardwareChangeMap) ForEach(cb func(gotenresource.Name, gotenresource.ResourceChange) bool) {
	for name, res := range m {
		if !cb(&name, res) {
			break
		}
	}
}
