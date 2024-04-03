// Code generated by protoc-gen-goten-resource
// Resource: Bucket
// DO NOT EDIT!!!

package bucket

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	iam_organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1/organization"
	iam_project "github.com/cloudwan/edgelq-sdk/iam/resources/v1/project"
	log_descriptor "github.com/cloudwan/edgelq-sdk/logging/resources/v1/log_descriptor"
	meta_service "github.com/cloudwan/goten-sdk/meta-service/resources/v1/service"
	meta "github.com/cloudwan/goten-sdk/types/meta"
)

// ensure the imports are used
var (
	_ = new(gotenresource.ListQuery)
)

// make sure we're using proto imports
var (
	_ = &iam_organization.Organization{}
	_ = &iam_project.Project{}
	_ = &log_descriptor.LogDescriptor{}
	_ = &meta_service.Service{}
	_ = &meta.Meta{}
)

type BucketList []*Bucket

func (l BucketList) Append(item gotenresource.Resource) gotenresource.ResourceList {
	return append(l, item.(*Bucket))
}

func (l BucketList) AppendList(list gotenresource.ResourceList) gotenresource.ResourceList {
	return append(l, list.(BucketList)...)
}

func (l BucketList) Slice(first, second int) gotenresource.ResourceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l BucketList) At(idx int) gotenresource.Resource {
	return l[idx]
}

func (l BucketList) Set(idx int, res gotenresource.Resource) {
	l[idx] = res.(*Bucket)
}

func (l BucketList) Length() int {
	return len(l)
}

type BucketChangeList []*BucketChange

func (l BucketChangeList) Append(item gotenresource.ResourceChange) gotenresource.ResourceChangeList {
	return append(l, item.(*BucketChange))
}

func (l BucketChangeList) AppendList(list gotenresource.ResourceChangeList) gotenresource.ResourceChangeList {
	return append(l, list.(BucketChangeList)...)
}

func (l BucketChangeList) Slice(first, second int) gotenresource.ResourceChangeList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l BucketChangeList) At(idx int) gotenresource.ResourceChange {
	return l[idx]
}

func (l BucketChangeList) Set(idx int, change gotenresource.ResourceChange) {
	l[idx] = change.(*BucketChange)
}

func (l BucketChangeList) Length() int {
	return len(l)
}

type BucketNameList []*Name

func (l BucketNameList) Append(name gotenresource.Name) gotenresource.NameList {
	return append(l, name.(*Name))
}

func (l BucketNameList) AppendList(list gotenresource.NameList) gotenresource.NameList {
	return append(l, list.(BucketNameList)...)
}

func (l BucketNameList) Slice(first, second int) gotenresource.NameList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l BucketNameList) At(idx int) gotenresource.Name {
	return l[idx]
}

func (l BucketNameList) Set(idx int, name gotenresource.Name) {
	l[idx] = name.(*Name)
}

func (l BucketNameList) Length() int {
	return len(l)
}

type BucketReferenceList []*Reference

func (l BucketReferenceList) Append(ref gotenresource.Reference) gotenresource.ReferenceList {
	return append(l, ref.(*Reference))
}

func (l BucketReferenceList) AppendList(list gotenresource.ReferenceList) gotenresource.ReferenceList {
	return append(l, list.(BucketReferenceList)...)
}

func (l BucketReferenceList) Slice(first, second int) gotenresource.ReferenceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l BucketReferenceList) At(idx int) gotenresource.Reference {
	return l[idx]
}

func (l BucketReferenceList) Set(idx int, ref gotenresource.Reference) {
	l[idx] = ref.(*Reference)
}

func (l BucketReferenceList) Length() int {
	return len(l)
}

type BucketParentNameList []*ParentName

func (l BucketParentNameList) Append(name gotenresource.Name) gotenresource.ParentNameList {
	return append(l, name.(*ParentName))
}

func (l BucketParentNameList) AppendList(list gotenresource.ParentNameList) gotenresource.ParentNameList {
	return append(l, list.(BucketParentNameList)...)
}

func (l BucketParentNameList) Slice(first, second int) gotenresource.ParentNameList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l BucketParentNameList) At(idx int) gotenresource.Name {
	return l[idx]
}

func (l BucketParentNameList) Set(idx int, name gotenresource.Name) {
	l[idx] = name.(*ParentName)
}

func (l BucketParentNameList) Length() int {
	return len(l)
}

type BucketParentReferenceList []*ParentReference

func (l BucketParentReferenceList) Append(ref gotenresource.Reference) gotenresource.ParentReferenceList {
	return append(l, ref.(*ParentReference))
}

func (l BucketParentReferenceList) AppendList(list gotenresource.ParentReferenceList) gotenresource.ParentReferenceList {
	return append(l, list.(BucketParentReferenceList)...)
}

func (l BucketParentReferenceList) Slice(first, second int) gotenresource.ParentReferenceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l BucketParentReferenceList) At(idx int) gotenresource.Reference {
	return l[idx]
}

func (l BucketParentReferenceList) Set(idx int, ref gotenresource.Reference) {
	l[idx] = ref.(*ParentReference)
}

func (l BucketParentReferenceList) Length() int {
	return len(l)
}

type BucketMap map[Name]*Bucket

func (m BucketMap) Get(name gotenresource.Name) gotenresource.Resource {
	return m[*name.(*Name)]
}

func (m BucketMap) Set(res gotenresource.Resource) {
	tRes := res.(*Bucket)
	m[*tRes.Name] = tRes
}

func (m BucketMap) Delete(name gotenresource.Name) {
	delete(m, *name.(*Name))
}

func (m BucketMap) Length() int {
	return len(m)
}

func (m BucketMap) ForEach(cb func(gotenresource.Name, gotenresource.Resource) bool) {
	for name, res := range m {
		if !cb(&name, res) {
			break
		}
	}
}

type BucketChangeMap map[Name]*BucketChange

func (m BucketChangeMap) Get(name gotenresource.Name) gotenresource.ResourceChange {
	return m[*name.(*Name)]
}

func (m BucketChangeMap) Set(change gotenresource.ResourceChange) {
	tChange := change.(*BucketChange)
	m[*tChange.GetBucketName()] = tChange
}

func (m BucketChangeMap) Delete(name gotenresource.Name) {
	delete(m, *name.(*Name))
}

func (m BucketChangeMap) Length() int {
	return len(m)
}

func (m BucketChangeMap) ForEach(cb func(gotenresource.Name, gotenresource.ResourceChange) bool) {
	for name, res := range m {
		if !cb(&name, res) {
			break
		}
	}
}