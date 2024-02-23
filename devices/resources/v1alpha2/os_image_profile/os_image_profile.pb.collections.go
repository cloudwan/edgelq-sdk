// Code generated by protoc-gen-goten-resource
// Resource: OsImageProfile
// DO NOT EDIT!!!

package os_image_profile

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
	device_type "github.com/cloudwan/edgelq-sdk/devices/resources/v1alpha2/device_type"
	project "github.com/cloudwan/edgelq-sdk/devices/resources/v1alpha2/project"
)

// ensure the imports are used
var (
	_ = new(gotenresource.ListQuery)
)

// make sure we're using proto imports
var (
	_ = &ntt_meta.Meta{}
	_ = &device_type.DeviceType{}
	_ = &project.Project{}
)

type OsImageProfileList []*OsImageProfile

func (l OsImageProfileList) Append(item gotenresource.Resource) gotenresource.ResourceList {
	return append(l, item.(*OsImageProfile))
}

func (l OsImageProfileList) AppendList(list gotenresource.ResourceList) gotenresource.ResourceList {
	return append(l, list.(OsImageProfileList)...)
}

func (l OsImageProfileList) Slice(first, second int) gotenresource.ResourceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l OsImageProfileList) At(idx int) gotenresource.Resource {
	return l[idx]
}

func (l OsImageProfileList) Set(idx int, res gotenresource.Resource) {
	l[idx] = res.(*OsImageProfile)
}

func (l OsImageProfileList) Length() int {
	return len(l)
}

type OsImageProfileChangeList []*OsImageProfileChange

func (l OsImageProfileChangeList) Append(item gotenresource.ResourceChange) gotenresource.ResourceChangeList {
	return append(l, item.(*OsImageProfileChange))
}

func (l OsImageProfileChangeList) AppendList(list gotenresource.ResourceChangeList) gotenresource.ResourceChangeList {
	return append(l, list.(OsImageProfileChangeList)...)
}

func (l OsImageProfileChangeList) Slice(first, second int) gotenresource.ResourceChangeList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l OsImageProfileChangeList) At(idx int) gotenresource.ResourceChange {
	return l[idx]
}

func (l OsImageProfileChangeList) Set(idx int, change gotenresource.ResourceChange) {
	l[idx] = change.(*OsImageProfileChange)
}

func (l OsImageProfileChangeList) Length() int {
	return len(l)
}

type OsImageProfileNameList []*Name

func (l OsImageProfileNameList) Append(name gotenresource.Name) gotenresource.NameList {
	return append(l, name.(*Name))
}

func (l OsImageProfileNameList) AppendList(list gotenresource.NameList) gotenresource.NameList {
	return append(l, list.(OsImageProfileNameList)...)
}

func (l OsImageProfileNameList) Slice(first, second int) gotenresource.NameList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l OsImageProfileNameList) At(idx int) gotenresource.Name {
	return l[idx]
}

func (l OsImageProfileNameList) Set(idx int, name gotenresource.Name) {
	l[idx] = name.(*Name)
}

func (l OsImageProfileNameList) Length() int {
	return len(l)
}

type OsImageProfileReferenceList []*Reference

func (l OsImageProfileReferenceList) Append(ref gotenresource.Reference) gotenresource.ReferenceList {
	return append(l, ref.(*Reference))
}

func (l OsImageProfileReferenceList) AppendList(list gotenresource.ReferenceList) gotenresource.ReferenceList {
	return append(l, list.(OsImageProfileReferenceList)...)
}

func (l OsImageProfileReferenceList) Slice(first, second int) gotenresource.ReferenceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l OsImageProfileReferenceList) At(idx int) gotenresource.Reference {
	return l[idx]
}

func (l OsImageProfileReferenceList) Set(idx int, ref gotenresource.Reference) {
	l[idx] = ref.(*Reference)
}

func (l OsImageProfileReferenceList) Length() int {
	return len(l)
}

type OsImageProfileParentNameList []*ParentName

func (l OsImageProfileParentNameList) Append(name gotenresource.Name) gotenresource.ParentNameList {
	return append(l, name.(*ParentName))
}

func (l OsImageProfileParentNameList) AppendList(list gotenresource.ParentNameList) gotenresource.ParentNameList {
	return append(l, list.(OsImageProfileParentNameList)...)
}

func (l OsImageProfileParentNameList) Slice(first, second int) gotenresource.ParentNameList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l OsImageProfileParentNameList) At(idx int) gotenresource.Name {
	return l[idx]
}

func (l OsImageProfileParentNameList) Set(idx int, name gotenresource.Name) {
	l[idx] = name.(*ParentName)
}

func (l OsImageProfileParentNameList) Length() int {
	return len(l)
}

type OsImageProfileParentReferenceList []*ParentReference

func (l OsImageProfileParentReferenceList) Append(ref gotenresource.Reference) gotenresource.ParentReferenceList {
	return append(l, ref.(*ParentReference))
}

func (l OsImageProfileParentReferenceList) AppendList(list gotenresource.ParentReferenceList) gotenresource.ParentReferenceList {
	return append(l, list.(OsImageProfileParentReferenceList)...)
}

func (l OsImageProfileParentReferenceList) Slice(first, second int) gotenresource.ParentReferenceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l OsImageProfileParentReferenceList) At(idx int) gotenresource.Reference {
	return l[idx]
}

func (l OsImageProfileParentReferenceList) Set(idx int, ref gotenresource.Reference) {
	l[idx] = ref.(*ParentReference)
}

func (l OsImageProfileParentReferenceList) Length() int {
	return len(l)
}

type OsImageProfileMap map[Name]*OsImageProfile

func (m OsImageProfileMap) Get(name gotenresource.Name) gotenresource.Resource {
	return m[*name.(*Name)]
}

func (m OsImageProfileMap) Set(res gotenresource.Resource) {
	tRes := res.(*OsImageProfile)
	m[*tRes.Name] = tRes
}

func (m OsImageProfileMap) Delete(name gotenresource.Name) {
	delete(m, *name.(*Name))
}

func (m OsImageProfileMap) Length() int {
	return len(m)
}

func (m OsImageProfileMap) ForEach(cb func(gotenresource.Name, gotenresource.Resource) bool) {
	for name, res := range m {
		if !cb(&name, res) {
			break
		}
	}
}

type OsImageProfileChangeMap map[Name]*OsImageProfileChange

func (m OsImageProfileChangeMap) Get(name gotenresource.Name) gotenresource.ResourceChange {
	return m[*name.(*Name)]
}

func (m OsImageProfileChangeMap) Set(change gotenresource.ResourceChange) {
	tChange := change.(*OsImageProfileChange)
	m[*tChange.GetOsImageProfileName()] = tChange
}

func (m OsImageProfileChangeMap) Delete(name gotenresource.Name) {
	delete(m, *name.(*Name))
}

func (m OsImageProfileChangeMap) Length() int {
	return len(m)
}

func (m OsImageProfileChangeMap) ForEach(cb func(gotenresource.Name, gotenresource.ResourceChange) bool) {
	for name, res := range m {
		if !cb(&name, res) {
			break
		}
	}
}
