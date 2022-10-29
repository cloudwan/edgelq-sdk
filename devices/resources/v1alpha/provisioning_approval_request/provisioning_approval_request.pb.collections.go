// Code generated by protoc-gen-goten-resource
// Resource: ProvisioningApprovalRequest
// DO NOT EDIT!!!

package provisioning_approval_request

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
	provisioning_policy "github.com/cloudwan/edgelq-sdk/devices/resources/v1alpha/provisioning_policy"
	iam_service_account "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/service_account"
)

// ensure the imports are used
var (
	_ = new(gotenresource.ListQuery)
)

// make sure we're using proto imports
var (
	_ = &ntt_meta.Meta{}
	_ = &provisioning_policy.ProvisioningPolicy{}
	_ = &iam_service_account.ServiceAccount{}
)

type ProvisioningApprovalRequestList []*ProvisioningApprovalRequest

func (l ProvisioningApprovalRequestList) Append(item gotenresource.Resource) gotenresource.ResourceList {
	return append(l, item.(*ProvisioningApprovalRequest))
}

func (l ProvisioningApprovalRequestList) AppendList(list gotenresource.ResourceList) gotenresource.ResourceList {
	return append(l, list.(ProvisioningApprovalRequestList)...)
}

func (l ProvisioningApprovalRequestList) Slice(first, second int) gotenresource.ResourceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l ProvisioningApprovalRequestList) At(idx int) gotenresource.Resource {
	return l[idx]
}

func (l ProvisioningApprovalRequestList) Set(idx int, res gotenresource.Resource) {
	l[idx] = res.(*ProvisioningApprovalRequest)
}

func (l ProvisioningApprovalRequestList) Length() int {
	return len(l)
}

type ProvisioningApprovalRequestChangeList []*ProvisioningApprovalRequestChange

func (l ProvisioningApprovalRequestChangeList) Append(item gotenresource.ResourceChange) gotenresource.ResourceChangeList {
	return append(l, item.(*ProvisioningApprovalRequestChange))
}

func (l ProvisioningApprovalRequestChangeList) AppendList(list gotenresource.ResourceChangeList) gotenresource.ResourceChangeList {
	return append(l, list.(ProvisioningApprovalRequestChangeList)...)
}

func (l ProvisioningApprovalRequestChangeList) Slice(first, second int) gotenresource.ResourceChangeList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l ProvisioningApprovalRequestChangeList) At(idx int) gotenresource.ResourceChange {
	return l[idx]
}

func (l ProvisioningApprovalRequestChangeList) Set(idx int, change gotenresource.ResourceChange) {
	l[idx] = change.(*ProvisioningApprovalRequestChange)
}

func (l ProvisioningApprovalRequestChangeList) Length() int {
	return len(l)
}

type ProvisioningApprovalRequestNameList []*Name

func (l ProvisioningApprovalRequestNameList) Append(name gotenresource.Name) gotenresource.NameList {
	return append(l, name.(*Name))
}

func (l ProvisioningApprovalRequestNameList) AppendList(list gotenresource.NameList) gotenresource.NameList {
	return append(l, list.(ProvisioningApprovalRequestNameList)...)
}

func (l ProvisioningApprovalRequestNameList) Slice(first, second int) gotenresource.NameList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l ProvisioningApprovalRequestNameList) At(idx int) gotenresource.Name {
	return l[idx]
}

func (l ProvisioningApprovalRequestNameList) Set(idx int, name gotenresource.Name) {
	l[idx] = name.(*Name)
}

func (l ProvisioningApprovalRequestNameList) Length() int {
	return len(l)
}

type ProvisioningApprovalRequestReferenceList []*Reference

func (l ProvisioningApprovalRequestReferenceList) Append(ref gotenresource.Reference) gotenresource.ReferenceList {
	return append(l, ref.(*Reference))
}

func (l ProvisioningApprovalRequestReferenceList) AppendList(list gotenresource.ReferenceList) gotenresource.ReferenceList {
	return append(l, list.(ProvisioningApprovalRequestReferenceList)...)
}

func (l ProvisioningApprovalRequestReferenceList) Slice(first, second int) gotenresource.ReferenceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l ProvisioningApprovalRequestReferenceList) At(idx int) gotenresource.Reference {
	return l[idx]
}

func (l ProvisioningApprovalRequestReferenceList) Set(idx int, ref gotenresource.Reference) {
	l[idx] = ref.(*Reference)
}

func (l ProvisioningApprovalRequestReferenceList) Length() int {
	return len(l)
}

type ProvisioningApprovalRequestParentNameList []*ParentName

func (l ProvisioningApprovalRequestParentNameList) Append(name gotenresource.Name) gotenresource.ParentNameList {
	return append(l, name.(*ParentName))
}

func (l ProvisioningApprovalRequestParentNameList) AppendList(list gotenresource.ParentNameList) gotenresource.ParentNameList {
	return append(l, list.(ProvisioningApprovalRequestParentNameList)...)
}

func (l ProvisioningApprovalRequestParentNameList) Slice(first, second int) gotenresource.ParentNameList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l ProvisioningApprovalRequestParentNameList) At(idx int) gotenresource.Name {
	return l[idx]
}

func (l ProvisioningApprovalRequestParentNameList) Set(idx int, name gotenresource.Name) {
	l[idx] = name.(*ParentName)
}

func (l ProvisioningApprovalRequestParentNameList) Length() int {
	return len(l)
}

type ProvisioningApprovalRequestParentReferenceList []*ParentReference

func (l ProvisioningApprovalRequestParentReferenceList) Append(ref gotenresource.Reference) gotenresource.ParentReferenceList {
	return append(l, ref.(*ParentReference))
}

func (l ProvisioningApprovalRequestParentReferenceList) AppendList(list gotenresource.ParentReferenceList) gotenresource.ParentReferenceList {
	return append(l, list.(ProvisioningApprovalRequestParentReferenceList)...)
}

func (l ProvisioningApprovalRequestParentReferenceList) Slice(first, second int) gotenresource.ParentReferenceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l ProvisioningApprovalRequestParentReferenceList) At(idx int) gotenresource.Reference {
	return l[idx]
}

func (l ProvisioningApprovalRequestParentReferenceList) Set(idx int, ref gotenresource.Reference) {
	l[idx] = ref.(*ParentReference)
}

func (l ProvisioningApprovalRequestParentReferenceList) Length() int {
	return len(l)
}

type ProvisioningApprovalRequestMap map[Name]*ProvisioningApprovalRequest

func (m ProvisioningApprovalRequestMap) Get(name gotenresource.Name) gotenresource.Resource {
	return m[*name.(*Name)]
}

func (m ProvisioningApprovalRequestMap) Set(res gotenresource.Resource) {
	tRes := res.(*ProvisioningApprovalRequest)
	m[*tRes.Name] = tRes
}

func (m ProvisioningApprovalRequestMap) Delete(name gotenresource.Name) {
	delete(m, *name.(*Name))
}

func (m ProvisioningApprovalRequestMap) Length() int {
	return len(m)
}

func (m ProvisioningApprovalRequestMap) ForEach(cb func(gotenresource.Name, gotenresource.Resource) bool) {
	for name, res := range m {
		if !cb(&name, res) {
			break
		}
	}
}

type ProvisioningApprovalRequestChangeMap map[Name]*ProvisioningApprovalRequestChange

func (m ProvisioningApprovalRequestChangeMap) Get(name gotenresource.Name) gotenresource.ResourceChange {
	return m[*name.(*Name)]
}

func (m ProvisioningApprovalRequestChangeMap) Set(change gotenresource.ResourceChange) {
	tChange := change.(*ProvisioningApprovalRequestChange)
	m[*tChange.GetProvisioningApprovalRequestName()] = tChange
}

func (m ProvisioningApprovalRequestChangeMap) Delete(name gotenresource.Name) {
	delete(m, *name.(*Name))
}

func (m ProvisioningApprovalRequestChangeMap) Length() int {
	return len(m)
}

func (m ProvisioningApprovalRequestChangeMap) ForEach(cb func(gotenresource.Name, gotenresource.ResourceChange) bool) {
	for name, res := range m {
		if !cb(&name, res) {
			break
		}
	}
}
