// Code generated by protoc-gen-goten-resource
// Resource: DeviceHardware
// DO NOT EDIT!!!

package device_hardware

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	cellular_api_sim_card "github.com/cloudwan/edgelq-sdk/cellular-api/resources/v1/sim_card"
	device "github.com/cloudwan/edgelq-sdk/devices/resources/v1/device"
	project "github.com/cloudwan/edgelq-sdk/devices/resources/v1/project"
	provisioning_policy "github.com/cloudwan/edgelq-sdk/devices/resources/v1/provisioning_policy"
	meta "github.com/cloudwan/goten-sdk/types/meta"
)

// ensure the imports are used
var (
	_ = new(gotenresource.ListQuery)
)

// make sure we're using proto imports
var (
	_ = &cellular_api_sim_card.SimCard{}
	_ = &device.Device{}
	_ = &project.Project{}
	_ = &provisioning_policy.ProvisioningPolicy{}
	_ = &meta.Meta{}
)

type DeviceHardwareList []*DeviceHardware

func (l DeviceHardwareList) Append(item gotenresource.Resource) gotenresource.ResourceList {
	return append(l, item.(*DeviceHardware))
}

func (l DeviceHardwareList) AppendList(list gotenresource.ResourceList) gotenresource.ResourceList {
	return append(l, list.(DeviceHardwareList)...)
}

func (l DeviceHardwareList) Slice(first, second int) gotenresource.ResourceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l DeviceHardwareList) At(idx int) gotenresource.Resource {
	return l[idx]
}

func (l DeviceHardwareList) Set(idx int, res gotenresource.Resource) {
	l[idx] = res.(*DeviceHardware)
}

func (l DeviceHardwareList) Length() int {
	return len(l)
}

type DeviceHardwareChangeList []*DeviceHardwareChange

func (l DeviceHardwareChangeList) Append(item gotenresource.ResourceChange) gotenresource.ResourceChangeList {
	return append(l, item.(*DeviceHardwareChange))
}

func (l DeviceHardwareChangeList) AppendList(list gotenresource.ResourceChangeList) gotenresource.ResourceChangeList {
	return append(l, list.(DeviceHardwareChangeList)...)
}

func (l DeviceHardwareChangeList) Slice(first, second int) gotenresource.ResourceChangeList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l DeviceHardwareChangeList) At(idx int) gotenresource.ResourceChange {
	return l[idx]
}

func (l DeviceHardwareChangeList) Set(idx int, change gotenresource.ResourceChange) {
	l[idx] = change.(*DeviceHardwareChange)
}

func (l DeviceHardwareChangeList) Length() int {
	return len(l)
}

type DeviceHardwareNameList []*Name

func (l DeviceHardwareNameList) Append(name gotenresource.Name) gotenresource.NameList {
	return append(l, name.(*Name))
}

func (l DeviceHardwareNameList) AppendList(list gotenresource.NameList) gotenresource.NameList {
	return append(l, list.(DeviceHardwareNameList)...)
}

func (l DeviceHardwareNameList) Slice(first, second int) gotenresource.NameList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l DeviceHardwareNameList) At(idx int) gotenresource.Name {
	return l[idx]
}

func (l DeviceHardwareNameList) Set(idx int, name gotenresource.Name) {
	l[idx] = name.(*Name)
}

func (l DeviceHardwareNameList) Length() int {
	return len(l)
}

type DeviceHardwareReferenceList []*Reference

func (l DeviceHardwareReferenceList) Append(ref gotenresource.Reference) gotenresource.ReferenceList {
	return append(l, ref.(*Reference))
}

func (l DeviceHardwareReferenceList) AppendList(list gotenresource.ReferenceList) gotenresource.ReferenceList {
	return append(l, list.(DeviceHardwareReferenceList)...)
}

func (l DeviceHardwareReferenceList) Slice(first, second int) gotenresource.ReferenceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l DeviceHardwareReferenceList) At(idx int) gotenresource.Reference {
	return l[idx]
}

func (l DeviceHardwareReferenceList) Set(idx int, ref gotenresource.Reference) {
	l[idx] = ref.(*Reference)
}

func (l DeviceHardwareReferenceList) Length() int {
	return len(l)
}

type DeviceHardwareParentNameList []*ParentName

func (l DeviceHardwareParentNameList) Append(name gotenresource.Name) gotenresource.ParentNameList {
	return append(l, name.(*ParentName))
}

func (l DeviceHardwareParentNameList) AppendList(list gotenresource.ParentNameList) gotenresource.ParentNameList {
	return append(l, list.(DeviceHardwareParentNameList)...)
}

func (l DeviceHardwareParentNameList) Slice(first, second int) gotenresource.ParentNameList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l DeviceHardwareParentNameList) At(idx int) gotenresource.Name {
	return l[idx]
}

func (l DeviceHardwareParentNameList) Set(idx int, name gotenresource.Name) {
	l[idx] = name.(*ParentName)
}

func (l DeviceHardwareParentNameList) Length() int {
	return len(l)
}

type DeviceHardwareParentReferenceList []*ParentReference

func (l DeviceHardwareParentReferenceList) Append(ref gotenresource.Reference) gotenresource.ParentReferenceList {
	return append(l, ref.(*ParentReference))
}

func (l DeviceHardwareParentReferenceList) AppendList(list gotenresource.ParentReferenceList) gotenresource.ParentReferenceList {
	return append(l, list.(DeviceHardwareParentReferenceList)...)
}

func (l DeviceHardwareParentReferenceList) Slice(first, second int) gotenresource.ParentReferenceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l DeviceHardwareParentReferenceList) At(idx int) gotenresource.Reference {
	return l[idx]
}

func (l DeviceHardwareParentReferenceList) Set(idx int, ref gotenresource.Reference) {
	l[idx] = ref.(*ParentReference)
}

func (l DeviceHardwareParentReferenceList) Length() int {
	return len(l)
}

type DeviceHardwareMap map[Name]*DeviceHardware

func (m DeviceHardwareMap) Get(name gotenresource.Name) gotenresource.Resource {
	return m[*name.(*Name)]
}

func (m DeviceHardwareMap) Set(res gotenresource.Resource) {
	tRes := res.(*DeviceHardware)
	m[*tRes.Name] = tRes
}

func (m DeviceHardwareMap) Delete(name gotenresource.Name) {
	delete(m, *name.(*Name))
}

func (m DeviceHardwareMap) Length() int {
	return len(m)
}

func (m DeviceHardwareMap) ForEach(cb func(gotenresource.Name, gotenresource.Resource) bool) {
	for name, res := range m {
		if !cb(&name, res) {
			break
		}
	}
}

type DeviceHardwareChangeMap map[Name]*DeviceHardwareChange

func (m DeviceHardwareChangeMap) Get(name gotenresource.Name) gotenresource.ResourceChange {
	return m[*name.(*Name)]
}

func (m DeviceHardwareChangeMap) Set(change gotenresource.ResourceChange) {
	tChange := change.(*DeviceHardwareChange)
	m[*tChange.GetDeviceHardwareName()] = tChange
}

func (m DeviceHardwareChangeMap) Delete(name gotenresource.Name) {
	delete(m, *name.(*Name))
}

func (m DeviceHardwareChangeMap) Length() int {
	return len(m)
}

func (m DeviceHardwareChangeMap) ForEach(cb func(gotenresource.Name, gotenresource.ResourceChange) bool) {
	for name, res := range m {
		if !cb(&name, res) {
			break
		}
	}
}
