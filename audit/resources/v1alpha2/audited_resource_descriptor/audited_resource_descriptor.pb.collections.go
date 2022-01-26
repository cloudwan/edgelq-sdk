// Code generated by protoc-gen-goten-resource
// Resource: AuditedResourceDescriptor
// DO NOT EDIT!!!

package audited_resource_descriptor

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	audit_common "github.com/cloudwan/edgelq-sdk/audit/common/v1alpha2"
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
)

// ensure the imports are used
var (
	_ = gotenresource.ListQuery(nil)
)

// make sure we're using proto imports
var (
	_ = &audit_common.Authentication{}
	_ = &ntt_meta.Meta{}
)

type AuditedResourceDescriptorList []*AuditedResourceDescriptor

func (l AuditedResourceDescriptorList) Append(item gotenresource.Resource) gotenresource.ResourceList {
	return append(l, item.(*AuditedResourceDescriptor))
}

func (l AuditedResourceDescriptorList) AppendList(list gotenresource.ResourceList) gotenresource.ResourceList {
	return append(l, list.(AuditedResourceDescriptorList)...)
}

func (l AuditedResourceDescriptorList) Slice(first, second int) gotenresource.ResourceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l AuditedResourceDescriptorList) At(idx int) gotenresource.Resource {
	return l[idx]
}

func (l AuditedResourceDescriptorList) Set(idx int, res gotenresource.Resource) {
	l[idx] = res.(*AuditedResourceDescriptor)
}

func (l AuditedResourceDescriptorList) Length() int {
	return len(l)
}

type AuditedResourceDescriptorChangeList []*AuditedResourceDescriptorChange

func (l AuditedResourceDescriptorChangeList) Append(item gotenresource.ResourceChange) gotenresource.ResourceChangeList {
	return append(l, item.(*AuditedResourceDescriptorChange))
}

func (l AuditedResourceDescriptorChangeList) AppendList(list gotenresource.ResourceChangeList) gotenresource.ResourceChangeList {
	return append(l, list.(AuditedResourceDescriptorChangeList)...)
}

func (l AuditedResourceDescriptorChangeList) Slice(first, second int) gotenresource.ResourceChangeList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l AuditedResourceDescriptorChangeList) At(idx int) gotenresource.ResourceChange {
	return l[idx]
}

func (l AuditedResourceDescriptorChangeList) Set(idx int, change gotenresource.ResourceChange) {
	l[idx] = change.(*AuditedResourceDescriptorChange)
}

func (l AuditedResourceDescriptorChangeList) Length() int {
	return len(l)
}

type AuditedResourceDescriptorNameList []*Name

func (l AuditedResourceDescriptorNameList) Append(name gotenresource.Name) gotenresource.NameList {
	return append(l, name.(*Name))
}

func (l AuditedResourceDescriptorNameList) AppendList(list gotenresource.NameList) gotenresource.NameList {
	return append(l, list.(AuditedResourceDescriptorNameList)...)
}

func (l AuditedResourceDescriptorNameList) Slice(first, second int) gotenresource.NameList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l AuditedResourceDescriptorNameList) At(idx int) gotenresource.Name {
	return l[idx]
}

func (l AuditedResourceDescriptorNameList) Set(idx int, name gotenresource.Name) {
	l[idx] = name.(*Name)
}

func (l AuditedResourceDescriptorNameList) Length() int {
	return len(l)
}

type AuditedResourceDescriptorReferenceList []*Reference

func (l AuditedResourceDescriptorReferenceList) Append(ref gotenresource.Reference) gotenresource.ReferenceList {
	return append(l, ref.(*Reference))
}

func (l AuditedResourceDescriptorReferenceList) AppendList(list gotenresource.ReferenceList) gotenresource.ReferenceList {
	return append(l, list.(AuditedResourceDescriptorReferenceList)...)
}

func (l AuditedResourceDescriptorReferenceList) Slice(first, second int) gotenresource.ReferenceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l AuditedResourceDescriptorReferenceList) At(idx int) gotenresource.Reference {
	return l[idx]
}

func (l AuditedResourceDescriptorReferenceList) Set(idx int, ref gotenresource.Reference) {
	l[idx] = ref.(*Reference)
}

func (l AuditedResourceDescriptorReferenceList) Length() int {
	return len(l)
}

type AuditedResourceDescriptorMap map[Name]*AuditedResourceDescriptor

func (m AuditedResourceDescriptorMap) Get(name gotenresource.Name) gotenresource.Resource {
	return m[*name.(*Name)]
}

func (m AuditedResourceDescriptorMap) Set(res gotenresource.Resource) {
	tRes := res.(*AuditedResourceDescriptor)
	m[*tRes.Name] = tRes
}

func (m AuditedResourceDescriptorMap) Delete(name gotenresource.Name) {
	delete(m, *name.(*Name))
}

func (m AuditedResourceDescriptorMap) Length() int {
	return len(m)
}

func (m AuditedResourceDescriptorMap) ForEach(cb func(gotenresource.Name, gotenresource.Resource) bool) {
	for name, res := range m {
		if !cb(&name, res) {
			break
		}
	}
}

type AuditedResourceDescriptorChangeMap map[Name]*AuditedResourceDescriptorChange

func (m AuditedResourceDescriptorChangeMap) Get(name gotenresource.Name) gotenresource.ResourceChange {
	return m[*name.(*Name)]
}

func (m AuditedResourceDescriptorChangeMap) Set(change gotenresource.ResourceChange) {
	tChange := change.(*AuditedResourceDescriptorChange)
	m[*tChange.GetAuditedResourceDescriptorName()] = tChange
}

func (m AuditedResourceDescriptorChangeMap) Delete(name gotenresource.Name) {
	delete(m, *name.(*Name))
}

func (m AuditedResourceDescriptorChangeMap) Length() int {
	return len(m)
}

func (m AuditedResourceDescriptorChangeMap) ForEach(cb func(gotenresource.Name, gotenresource.ResourceChange) bool) {
	for name, res := range m {
		if !cb(&name, res) {
			break
		}
	}
}
