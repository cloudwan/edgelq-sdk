// Code generated by protoc-gen-goten-resource
// Resource: TimeSeriesForwarderSink
// DO NOT EDIT!!!

package time_series_forwarder_sink

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	project "github.com/cloudwan/edgelq-sdk/monitoring/resources/v4/project"
	secrets_secret "github.com/cloudwan/edgelq-sdk/secrets/resources/v1/secret"
	meta "github.com/cloudwan/goten-sdk/types/meta"
)

// ensure the imports are used
var (
	_ = new(gotenresource.ListQuery)
)

// make sure we're using proto imports
var (
	_ = &project.Project{}
	_ = &secrets_secret.Secret{}
	_ = &meta.Meta{}
)

type TimeSeriesForwarderSinkList []*TimeSeriesForwarderSink

func (l TimeSeriesForwarderSinkList) Append(item gotenresource.Resource) gotenresource.ResourceList {
	return append(l, item.(*TimeSeriesForwarderSink))
}

func (l TimeSeriesForwarderSinkList) AppendList(list gotenresource.ResourceList) gotenresource.ResourceList {
	return append(l, list.(TimeSeriesForwarderSinkList)...)
}

func (l TimeSeriesForwarderSinkList) Slice(first, second int) gotenresource.ResourceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l TimeSeriesForwarderSinkList) At(idx int) gotenresource.Resource {
	return l[idx]
}

func (l TimeSeriesForwarderSinkList) Set(idx int, res gotenresource.Resource) {
	l[idx] = res.(*TimeSeriesForwarderSink)
}

func (l TimeSeriesForwarderSinkList) Length() int {
	return len(l)
}

type TimeSeriesForwarderSinkChangeList []*TimeSeriesForwarderSinkChange

func (l TimeSeriesForwarderSinkChangeList) Append(item gotenresource.ResourceChange) gotenresource.ResourceChangeList {
	return append(l, item.(*TimeSeriesForwarderSinkChange))
}

func (l TimeSeriesForwarderSinkChangeList) AppendList(list gotenresource.ResourceChangeList) gotenresource.ResourceChangeList {
	return append(l, list.(TimeSeriesForwarderSinkChangeList)...)
}

func (l TimeSeriesForwarderSinkChangeList) Slice(first, second int) gotenresource.ResourceChangeList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l TimeSeriesForwarderSinkChangeList) At(idx int) gotenresource.ResourceChange {
	return l[idx]
}

func (l TimeSeriesForwarderSinkChangeList) Set(idx int, change gotenresource.ResourceChange) {
	l[idx] = change.(*TimeSeriesForwarderSinkChange)
}

func (l TimeSeriesForwarderSinkChangeList) Length() int {
	return len(l)
}

type TimeSeriesForwarderSinkNameList []*Name

func (l TimeSeriesForwarderSinkNameList) Append(name gotenresource.Name) gotenresource.NameList {
	return append(l, name.(*Name))
}

func (l TimeSeriesForwarderSinkNameList) AppendList(list gotenresource.NameList) gotenresource.NameList {
	return append(l, list.(TimeSeriesForwarderSinkNameList)...)
}

func (l TimeSeriesForwarderSinkNameList) Slice(first, second int) gotenresource.NameList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l TimeSeriesForwarderSinkNameList) At(idx int) gotenresource.Name {
	return l[idx]
}

func (l TimeSeriesForwarderSinkNameList) Set(idx int, name gotenresource.Name) {
	l[idx] = name.(*Name)
}

func (l TimeSeriesForwarderSinkNameList) Length() int {
	return len(l)
}

type TimeSeriesForwarderSinkReferenceList []*Reference

func (l TimeSeriesForwarderSinkReferenceList) Append(ref gotenresource.Reference) gotenresource.ReferenceList {
	return append(l, ref.(*Reference))
}

func (l TimeSeriesForwarderSinkReferenceList) AppendList(list gotenresource.ReferenceList) gotenresource.ReferenceList {
	return append(l, list.(TimeSeriesForwarderSinkReferenceList)...)
}

func (l TimeSeriesForwarderSinkReferenceList) Slice(first, second int) gotenresource.ReferenceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l TimeSeriesForwarderSinkReferenceList) At(idx int) gotenresource.Reference {
	return l[idx]
}

func (l TimeSeriesForwarderSinkReferenceList) Set(idx int, ref gotenresource.Reference) {
	l[idx] = ref.(*Reference)
}

func (l TimeSeriesForwarderSinkReferenceList) Length() int {
	return len(l)
}

type TimeSeriesForwarderSinkParentNameList []*ParentName

func (l TimeSeriesForwarderSinkParentNameList) Append(name gotenresource.Name) gotenresource.ParentNameList {
	return append(l, name.(*ParentName))
}

func (l TimeSeriesForwarderSinkParentNameList) AppendList(list gotenresource.ParentNameList) gotenresource.ParentNameList {
	return append(l, list.(TimeSeriesForwarderSinkParentNameList)...)
}

func (l TimeSeriesForwarderSinkParentNameList) Slice(first, second int) gotenresource.ParentNameList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l TimeSeriesForwarderSinkParentNameList) At(idx int) gotenresource.Name {
	return l[idx]
}

func (l TimeSeriesForwarderSinkParentNameList) Set(idx int, name gotenresource.Name) {
	l[idx] = name.(*ParentName)
}

func (l TimeSeriesForwarderSinkParentNameList) Length() int {
	return len(l)
}

type TimeSeriesForwarderSinkParentReferenceList []*ParentReference

func (l TimeSeriesForwarderSinkParentReferenceList) Append(ref gotenresource.Reference) gotenresource.ParentReferenceList {
	return append(l, ref.(*ParentReference))
}

func (l TimeSeriesForwarderSinkParentReferenceList) AppendList(list gotenresource.ParentReferenceList) gotenresource.ParentReferenceList {
	return append(l, list.(TimeSeriesForwarderSinkParentReferenceList)...)
}

func (l TimeSeriesForwarderSinkParentReferenceList) Slice(first, second int) gotenresource.ParentReferenceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l TimeSeriesForwarderSinkParentReferenceList) At(idx int) gotenresource.Reference {
	return l[idx]
}

func (l TimeSeriesForwarderSinkParentReferenceList) Set(idx int, ref gotenresource.Reference) {
	l[idx] = ref.(*ParentReference)
}

func (l TimeSeriesForwarderSinkParentReferenceList) Length() int {
	return len(l)
}

type TimeSeriesForwarderSinkMap map[Name]*TimeSeriesForwarderSink

func (m TimeSeriesForwarderSinkMap) Get(name gotenresource.Name) gotenresource.Resource {
	return m[*name.(*Name)]
}

func (m TimeSeriesForwarderSinkMap) Set(res gotenresource.Resource) {
	tRes := res.(*TimeSeriesForwarderSink)
	m[*tRes.Name] = tRes
}

func (m TimeSeriesForwarderSinkMap) Delete(name gotenresource.Name) {
	delete(m, *name.(*Name))
}

func (m TimeSeriesForwarderSinkMap) Length() int {
	return len(m)
}

func (m TimeSeriesForwarderSinkMap) ForEach(cb func(gotenresource.Name, gotenresource.Resource) bool) {
	for name, res := range m {
		if !cb(&name, res) {
			break
		}
	}
}

type TimeSeriesForwarderSinkChangeMap map[Name]*TimeSeriesForwarderSinkChange

func (m TimeSeriesForwarderSinkChangeMap) Get(name gotenresource.Name) gotenresource.ResourceChange {
	return m[*name.(*Name)]
}

func (m TimeSeriesForwarderSinkChangeMap) Set(change gotenresource.ResourceChange) {
	tChange := change.(*TimeSeriesForwarderSinkChange)
	m[*tChange.GetTimeSeriesForwarderSinkName()] = tChange
}

func (m TimeSeriesForwarderSinkChangeMap) Delete(name gotenresource.Name) {
	delete(m, *name.(*Name))
}

func (m TimeSeriesForwarderSinkChangeMap) Length() int {
	return len(m)
}

func (m TimeSeriesForwarderSinkChangeMap) ForEach(cb func(gotenresource.Name, gotenresource.ResourceChange) bool) {
	for name, res := range m {
		if !cb(&name, res) {
			break
		}
	}
}
