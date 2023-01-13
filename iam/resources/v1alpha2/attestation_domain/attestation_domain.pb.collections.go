// Code generated by protoc-gen-goten-resource
// Resource: AttestationDomain
// DO NOT EDIT!!!

package attestation_domain

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
	iam_common "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/common"
	project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/project"
)

// ensure the imports are used
var (
	_ = new(gotenresource.ListQuery)
)

// make sure we're using proto imports
var (
	_ = &ntt_meta.Meta{}
	_ = &iam_common.PCR{}
	_ = &project.Project{}
)

type AttestationDomainList []*AttestationDomain

func (l AttestationDomainList) Append(item gotenresource.Resource) gotenresource.ResourceList {
	return append(l, item.(*AttestationDomain))
}

func (l AttestationDomainList) AppendList(list gotenresource.ResourceList) gotenresource.ResourceList {
	return append(l, list.(AttestationDomainList)...)
}

func (l AttestationDomainList) Slice(first, second int) gotenresource.ResourceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l AttestationDomainList) At(idx int) gotenresource.Resource {
	return l[idx]
}

func (l AttestationDomainList) Set(idx int, res gotenresource.Resource) {
	l[idx] = res.(*AttestationDomain)
}

func (l AttestationDomainList) Length() int {
	return len(l)
}

type AttestationDomainChangeList []*AttestationDomainChange

func (l AttestationDomainChangeList) Append(item gotenresource.ResourceChange) gotenresource.ResourceChangeList {
	return append(l, item.(*AttestationDomainChange))
}

func (l AttestationDomainChangeList) AppendList(list gotenresource.ResourceChangeList) gotenresource.ResourceChangeList {
	return append(l, list.(AttestationDomainChangeList)...)
}

func (l AttestationDomainChangeList) Slice(first, second int) gotenresource.ResourceChangeList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l AttestationDomainChangeList) At(idx int) gotenresource.ResourceChange {
	return l[idx]
}

func (l AttestationDomainChangeList) Set(idx int, change gotenresource.ResourceChange) {
	l[idx] = change.(*AttestationDomainChange)
}

func (l AttestationDomainChangeList) Length() int {
	return len(l)
}

type AttestationDomainNameList []*Name

func (l AttestationDomainNameList) Append(name gotenresource.Name) gotenresource.NameList {
	return append(l, name.(*Name))
}

func (l AttestationDomainNameList) AppendList(list gotenresource.NameList) gotenresource.NameList {
	return append(l, list.(AttestationDomainNameList)...)
}

func (l AttestationDomainNameList) Slice(first, second int) gotenresource.NameList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l AttestationDomainNameList) At(idx int) gotenresource.Name {
	return l[idx]
}

func (l AttestationDomainNameList) Set(idx int, name gotenresource.Name) {
	l[idx] = name.(*Name)
}

func (l AttestationDomainNameList) Length() int {
	return len(l)
}

type AttestationDomainReferenceList []*Reference

func (l AttestationDomainReferenceList) Append(ref gotenresource.Reference) gotenresource.ReferenceList {
	return append(l, ref.(*Reference))
}

func (l AttestationDomainReferenceList) AppendList(list gotenresource.ReferenceList) gotenresource.ReferenceList {
	return append(l, list.(AttestationDomainReferenceList)...)
}

func (l AttestationDomainReferenceList) Slice(first, second int) gotenresource.ReferenceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l AttestationDomainReferenceList) At(idx int) gotenresource.Reference {
	return l[idx]
}

func (l AttestationDomainReferenceList) Set(idx int, ref gotenresource.Reference) {
	l[idx] = ref.(*Reference)
}

func (l AttestationDomainReferenceList) Length() int {
	return len(l)
}

type AttestationDomainParentNameList []*ParentName

func (l AttestationDomainParentNameList) Append(name gotenresource.Name) gotenresource.ParentNameList {
	return append(l, name.(*ParentName))
}

func (l AttestationDomainParentNameList) AppendList(list gotenresource.ParentNameList) gotenresource.ParentNameList {
	return append(l, list.(AttestationDomainParentNameList)...)
}

func (l AttestationDomainParentNameList) Slice(first, second int) gotenresource.ParentNameList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l AttestationDomainParentNameList) At(idx int) gotenresource.Name {
	return l[idx]
}

func (l AttestationDomainParentNameList) Set(idx int, name gotenresource.Name) {
	l[idx] = name.(*ParentName)
}

func (l AttestationDomainParentNameList) Length() int {
	return len(l)
}

type AttestationDomainParentReferenceList []*ParentReference

func (l AttestationDomainParentReferenceList) Append(ref gotenresource.Reference) gotenresource.ParentReferenceList {
	return append(l, ref.(*ParentReference))
}

func (l AttestationDomainParentReferenceList) AppendList(list gotenresource.ParentReferenceList) gotenresource.ParentReferenceList {
	return append(l, list.(AttestationDomainParentReferenceList)...)
}

func (l AttestationDomainParentReferenceList) Slice(first, second int) gotenresource.ParentReferenceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l AttestationDomainParentReferenceList) At(idx int) gotenresource.Reference {
	return l[idx]
}

func (l AttestationDomainParentReferenceList) Set(idx int, ref gotenresource.Reference) {
	l[idx] = ref.(*ParentReference)
}

func (l AttestationDomainParentReferenceList) Length() int {
	return len(l)
}

type AttestationDomainMap map[Name]*AttestationDomain

func (m AttestationDomainMap) Get(name gotenresource.Name) gotenresource.Resource {
	return m[*name.(*Name)]
}

func (m AttestationDomainMap) Set(res gotenresource.Resource) {
	tRes := res.(*AttestationDomain)
	m[*tRes.Name] = tRes
}

func (m AttestationDomainMap) Delete(name gotenresource.Name) {
	delete(m, *name.(*Name))
}

func (m AttestationDomainMap) Length() int {
	return len(m)
}

func (m AttestationDomainMap) ForEach(cb func(gotenresource.Name, gotenresource.Resource) bool) {
	for name, res := range m {
		if !cb(&name, res) {
			break
		}
	}
}

type AttestationDomainChangeMap map[Name]*AttestationDomainChange

func (m AttestationDomainChangeMap) Get(name gotenresource.Name) gotenresource.ResourceChange {
	return m[*name.(*Name)]
}

func (m AttestationDomainChangeMap) Set(change gotenresource.ResourceChange) {
	tChange := change.(*AttestationDomainChange)
	m[*tChange.GetAttestationDomainName()] = tChange
}

func (m AttestationDomainChangeMap) Delete(name gotenresource.Name) {
	delete(m, *name.(*Name))
}

func (m AttestationDomainChangeMap) Length() int {
	return len(m)
}

func (m AttestationDomainChangeMap) ForEach(cb func(gotenresource.Name, gotenresource.ResourceChange) bool) {
	for name, res := range m {
		if !cb(&name, res) {
			break
		}
	}
}
