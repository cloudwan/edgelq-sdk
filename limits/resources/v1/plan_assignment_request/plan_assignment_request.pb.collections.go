// Code generated by protoc-gen-goten-resource
// Resource: PlanAssignmentRequest
// DO NOT EDIT!!!

package plan_assignment_request

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	iam_organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1/organization"
	iam_project "github.com/cloudwan/edgelq-sdk/iam/resources/v1/project"
	accepted_plan "github.com/cloudwan/edgelq-sdk/limits/resources/v1/accepted_plan"
	common "github.com/cloudwan/edgelq-sdk/limits/resources/v1/common"
	plan "github.com/cloudwan/edgelq-sdk/limits/resources/v1/plan"
	plan_assignment "github.com/cloudwan/edgelq-sdk/limits/resources/v1/plan_assignment"
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
	_ = &accepted_plan.AcceptedPlan{}
	_ = &common.RegionalPlanAssignment{}
	_ = &plan.Plan{}
	_ = &plan_assignment.PlanAssignment{}
	_ = &meta_service.Service{}
	_ = &meta.Meta{}
)

type PlanAssignmentRequestList []*PlanAssignmentRequest

func (l PlanAssignmentRequestList) Append(item gotenresource.Resource) gotenresource.ResourceList {
	return append(l, item.(*PlanAssignmentRequest))
}

func (l PlanAssignmentRequestList) AppendList(list gotenresource.ResourceList) gotenresource.ResourceList {
	return append(l, list.(PlanAssignmentRequestList)...)
}

func (l PlanAssignmentRequestList) Slice(first, second int) gotenresource.ResourceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l PlanAssignmentRequestList) At(idx int) gotenresource.Resource {
	return l[idx]
}

func (l PlanAssignmentRequestList) Set(idx int, res gotenresource.Resource) {
	l[idx] = res.(*PlanAssignmentRequest)
}

func (l PlanAssignmentRequestList) Length() int {
	return len(l)
}

type PlanAssignmentRequestChangeList []*PlanAssignmentRequestChange

func (l PlanAssignmentRequestChangeList) Append(item gotenresource.ResourceChange) gotenresource.ResourceChangeList {
	return append(l, item.(*PlanAssignmentRequestChange))
}

func (l PlanAssignmentRequestChangeList) AppendList(list gotenresource.ResourceChangeList) gotenresource.ResourceChangeList {
	return append(l, list.(PlanAssignmentRequestChangeList)...)
}

func (l PlanAssignmentRequestChangeList) Slice(first, second int) gotenresource.ResourceChangeList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l PlanAssignmentRequestChangeList) At(idx int) gotenresource.ResourceChange {
	return l[idx]
}

func (l PlanAssignmentRequestChangeList) Set(idx int, change gotenresource.ResourceChange) {
	l[idx] = change.(*PlanAssignmentRequestChange)
}

func (l PlanAssignmentRequestChangeList) Length() int {
	return len(l)
}

type PlanAssignmentRequestNameList []*Name

func (l PlanAssignmentRequestNameList) Append(name gotenresource.Name) gotenresource.NameList {
	return append(l, name.(*Name))
}

func (l PlanAssignmentRequestNameList) AppendList(list gotenresource.NameList) gotenresource.NameList {
	return append(l, list.(PlanAssignmentRequestNameList)...)
}

func (l PlanAssignmentRequestNameList) Slice(first, second int) gotenresource.NameList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l PlanAssignmentRequestNameList) At(idx int) gotenresource.Name {
	return l[idx]
}

func (l PlanAssignmentRequestNameList) Set(idx int, name gotenresource.Name) {
	l[idx] = name.(*Name)
}

func (l PlanAssignmentRequestNameList) Length() int {
	return len(l)
}

type PlanAssignmentRequestReferenceList []*Reference

func (l PlanAssignmentRequestReferenceList) Append(ref gotenresource.Reference) gotenresource.ReferenceList {
	return append(l, ref.(*Reference))
}

func (l PlanAssignmentRequestReferenceList) AppendList(list gotenresource.ReferenceList) gotenresource.ReferenceList {
	return append(l, list.(PlanAssignmentRequestReferenceList)...)
}

func (l PlanAssignmentRequestReferenceList) Slice(first, second int) gotenresource.ReferenceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l PlanAssignmentRequestReferenceList) At(idx int) gotenresource.Reference {
	return l[idx]
}

func (l PlanAssignmentRequestReferenceList) Set(idx int, ref gotenresource.Reference) {
	l[idx] = ref.(*Reference)
}

func (l PlanAssignmentRequestReferenceList) Length() int {
	return len(l)
}

type PlanAssignmentRequestParentNameList []*ParentName

func (l PlanAssignmentRequestParentNameList) Append(name gotenresource.Name) gotenresource.ParentNameList {
	return append(l, name.(*ParentName))
}

func (l PlanAssignmentRequestParentNameList) AppendList(list gotenresource.ParentNameList) gotenresource.ParentNameList {
	return append(l, list.(PlanAssignmentRequestParentNameList)...)
}

func (l PlanAssignmentRequestParentNameList) Slice(first, second int) gotenresource.ParentNameList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l PlanAssignmentRequestParentNameList) At(idx int) gotenresource.Name {
	return l[idx]
}

func (l PlanAssignmentRequestParentNameList) Set(idx int, name gotenresource.Name) {
	l[idx] = name.(*ParentName)
}

func (l PlanAssignmentRequestParentNameList) Length() int {
	return len(l)
}

type PlanAssignmentRequestParentReferenceList []*ParentReference

func (l PlanAssignmentRequestParentReferenceList) Append(ref gotenresource.Reference) gotenresource.ParentReferenceList {
	return append(l, ref.(*ParentReference))
}

func (l PlanAssignmentRequestParentReferenceList) AppendList(list gotenresource.ParentReferenceList) gotenresource.ParentReferenceList {
	return append(l, list.(PlanAssignmentRequestParentReferenceList)...)
}

func (l PlanAssignmentRequestParentReferenceList) Slice(first, second int) gotenresource.ParentReferenceList {
	if first > 0 && second > 0 {
		return l[first:second]
	} else if first > 0 {
		return l[first:]
	} else if second > 0 {
		return l[:second]
	}
	return l[:]
}

func (l PlanAssignmentRequestParentReferenceList) At(idx int) gotenresource.Reference {
	return l[idx]
}

func (l PlanAssignmentRequestParentReferenceList) Set(idx int, ref gotenresource.Reference) {
	l[idx] = ref.(*ParentReference)
}

func (l PlanAssignmentRequestParentReferenceList) Length() int {
	return len(l)
}

type PlanAssignmentRequestMap map[Name]*PlanAssignmentRequest

func (m PlanAssignmentRequestMap) Get(name gotenresource.Name) gotenresource.Resource {
	return m[*name.(*Name)]
}

func (m PlanAssignmentRequestMap) Set(res gotenresource.Resource) {
	tRes := res.(*PlanAssignmentRequest)
	m[*tRes.Name] = tRes
}

func (m PlanAssignmentRequestMap) Delete(name gotenresource.Name) {
	delete(m, *name.(*Name))
}

func (m PlanAssignmentRequestMap) Length() int {
	return len(m)
}

func (m PlanAssignmentRequestMap) ForEach(cb func(gotenresource.Name, gotenresource.Resource) bool) {
	for name, res := range m {
		if !cb(&name, res) {
			break
		}
	}
}

type PlanAssignmentRequestChangeMap map[Name]*PlanAssignmentRequestChange

func (m PlanAssignmentRequestChangeMap) Get(name gotenresource.Name) gotenresource.ResourceChange {
	return m[*name.(*Name)]
}

func (m PlanAssignmentRequestChangeMap) Set(change gotenresource.ResourceChange) {
	tChange := change.(*PlanAssignmentRequestChange)
	m[*tChange.GetPlanAssignmentRequestName()] = tChange
}

func (m PlanAssignmentRequestChangeMap) Delete(name gotenresource.Name) {
	delete(m, *name.(*Name))
}

func (m PlanAssignmentRequestChangeMap) Length() int {
	return len(m)
}

func (m PlanAssignmentRequestChangeMap) ForEach(cb func(gotenresource.Name, gotenresource.ResourceChange) bool) {
	for name, res := range m {
		if !cb(&name, res) {
			break
		}
	}
}
