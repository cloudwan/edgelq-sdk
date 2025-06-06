// Code generated by protoc-gen-goten-object
// File: edgelq/limits/proto/v1alpha2/plan_assignment_request.proto
// DO NOT EDIT!!!

package plan_assignment_request

import (
	"fmt"
	"sort"

	"google.golang.org/protobuf/proto"
	googlefieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"

	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
)

// proto imports
import (
	iam_organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/organization"
	iam_project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/project"
	common "github.com/cloudwan/edgelq-sdk/limits/resources/v1alpha2/common"
	plan "github.com/cloudwan/edgelq-sdk/limits/resources/v1alpha2/plan"
	plan_assignment "github.com/cloudwan/edgelq-sdk/limits/resources/v1alpha2/plan_assignment"
	meta_service "github.com/cloudwan/edgelq-sdk/meta/resources/v1alpha2/service"
	meta "github.com/cloudwan/goten-sdk/types/meta"
)

// ensure the imports are used
var (
	_ = new(fmt.Stringer)
	_ = new(sort.Interface)

	_ = new(proto.Message)
	_ = googlefieldmaskpb.FieldMask{}

	_ = new(gotenobject.FieldPath)
)

// make sure we're using proto imports
var (
	_ = &iam_organization.Organization{}
	_ = &iam_project.Project{}
	_ = &common.Allowance{}
	_ = &plan.Plan{}
	_ = &plan_assignment.PlanAssignment{}
	_ = &meta_service.Service{}
	_ = &meta.Meta{}
)

func (o *PlanAssignmentRequest) GotenObjectExt() {}

func (o *PlanAssignmentRequest) MakeFullFieldMask() *PlanAssignmentRequest_FieldMask {
	return FullPlanAssignmentRequest_FieldMask()
}

func (o *PlanAssignmentRequest) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullPlanAssignmentRequest_FieldMask()
}

func (o *PlanAssignmentRequest) MakeDiffFieldMask(other *PlanAssignmentRequest) *PlanAssignmentRequest_FieldMask {
	if o == nil && other == nil {
		return &PlanAssignmentRequest_FieldMask{}
	}
	if o == nil || other == nil {
		return FullPlanAssignmentRequest_FieldMask()
	}

	res := &PlanAssignmentRequest_FieldMask{}
	if o.GetName().String() != other.GetName().String() {
		res.Paths = append(res.Paths, &PlanAssignmentRequest_FieldTerminalPath{selector: PlanAssignmentRequest_FieldPathSelectorName})
	}
	{
		subMask := o.GetRequest().MakeDiffFieldMask(other.GetRequest())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &PlanAssignmentRequest_FieldTerminalPath{selector: PlanAssignmentRequest_FieldPathSelectorRequest})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &PlanAssignmentRequest_FieldSubPath{selector: PlanAssignmentRequest_FieldPathSelectorRequest, subPath: subpath})
			}
		}
	}
	if o.GetService().String() != other.GetService().String() {
		res.Paths = append(res.Paths, &PlanAssignmentRequest_FieldTerminalPath{selector: PlanAssignmentRequest_FieldPathSelectorService})
	}
	if o.GetApprover().String() != other.GetApprover().String() {
		res.Paths = append(res.Paths, &PlanAssignmentRequest_FieldTerminalPath{selector: PlanAssignmentRequest_FieldPathSelectorApprover})
	}
	{
		subMask := o.GetStatus().MakeDiffFieldMask(other.GetStatus())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &PlanAssignmentRequest_FieldTerminalPath{selector: PlanAssignmentRequest_FieldPathSelectorStatus})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &PlanAssignmentRequest_FieldSubPath{selector: PlanAssignmentRequest_FieldPathSelectorStatus, subPath: subpath})
			}
		}
	}
	{
		subMask := o.GetMetadata().MakeDiffFieldMask(other.GetMetadata())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &PlanAssignmentRequest_FieldTerminalPath{selector: PlanAssignmentRequest_FieldPathSelectorMetadata})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &PlanAssignmentRequest_FieldSubPath{selector: PlanAssignmentRequest_FieldPathSelectorMetadata, subPath: subpath})
			}
		}
	}
	return res
}

func (o *PlanAssignmentRequest) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*PlanAssignmentRequest))
}

func (o *PlanAssignmentRequest) Clone() *PlanAssignmentRequest {
	if o == nil {
		return nil
	}
	result := &PlanAssignmentRequest{}
	if o.Name == nil {
		result.Name = nil
	} else if data, err := o.Name.ProtoString(); err != nil {
		panic(err)
	} else {
		result.Name = &Name{}
		if err := result.Name.ParseProtoString(data); err != nil {
			panic(err)
		}
	}
	result.Request = o.Request.Clone()
	if o.Service == nil {
		result.Service = nil
	} else if data, err := o.Service.ProtoString(); err != nil {
		panic(err)
	} else {
		result.Service = &meta_service.Reference{}
		if err := result.Service.ParseProtoString(data); err != nil {
			panic(err)
		}
	}
	if o.Approver == nil {
		result.Approver = nil
	} else if data, err := o.Approver.ProtoString(); err != nil {
		panic(err)
	} else {
		result.Approver = &iam_organization.Reference{}
		if err := result.Approver.ParseProtoString(data); err != nil {
			panic(err)
		}
	}
	result.Status = o.Status.Clone()
	result.Metadata = o.Metadata.Clone()
	return result
}

func (o *PlanAssignmentRequest) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *PlanAssignmentRequest) Merge(source *PlanAssignmentRequest) {
	if source.GetName() != nil {
		if data, err := source.GetName().ProtoString(); err != nil {
			panic(err)
		} else {
			o.Name = &Name{}
			if err := o.Name.ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	} else {
		o.Name = nil
	}
	if source.GetRequest() != nil {
		if o.Request == nil {
			o.Request = new(PlanAssignmentRequest_RequestType)
		}
		o.Request.Merge(source.GetRequest())
	}
	if source.GetService() != nil {
		if data, err := source.GetService().ProtoString(); err != nil {
			panic(err)
		} else {
			o.Service = &meta_service.Reference{}
			if err := o.Service.ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	} else {
		o.Service = nil
	}
	if source.GetApprover() != nil {
		if data, err := source.GetApprover().ProtoString(); err != nil {
			panic(err)
		} else {
			o.Approver = &iam_organization.Reference{}
			if err := o.Approver.ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	} else {
		o.Approver = nil
	}
	if source.GetStatus() != nil {
		if o.Status == nil {
			o.Status = new(PlanAssignmentRequest_Status)
		}
		o.Status.Merge(source.GetStatus())
	}
	if source.GetMetadata() != nil {
		if o.Metadata == nil {
			o.Metadata = new(meta.Meta)
		}
		o.Metadata.Merge(source.GetMetadata())
	}
}

func (o *PlanAssignmentRequest) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*PlanAssignmentRequest))
}

func (o *PlanAssignmentRequest_Status) GotenObjectExt() {}

func (o *PlanAssignmentRequest_Status) MakeFullFieldMask() *PlanAssignmentRequest_Status_FieldMask {
	return FullPlanAssignmentRequest_Status_FieldMask()
}

func (o *PlanAssignmentRequest_Status) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullPlanAssignmentRequest_Status_FieldMask()
}

func (o *PlanAssignmentRequest_Status) MakeDiffFieldMask(other *PlanAssignmentRequest_Status) *PlanAssignmentRequest_Status_FieldMask {
	if o == nil && other == nil {
		return &PlanAssignmentRequest_Status_FieldMask{}
	}
	if o == nil || other == nil {
		return FullPlanAssignmentRequest_Status_FieldMask()
	}

	res := &PlanAssignmentRequest_Status_FieldMask{}
	if o.GetConclusion() != other.GetConclusion() {
		res.Paths = append(res.Paths, &PlanAssignmentRequestStatus_FieldTerminalPath{selector: PlanAssignmentRequestStatus_FieldPathSelectorConclusion})
	}
	if o.GetReason() != other.GetReason() {
		res.Paths = append(res.Paths, &PlanAssignmentRequestStatus_FieldTerminalPath{selector: PlanAssignmentRequestStatus_FieldPathSelectorReason})
	}
	return res
}

func (o *PlanAssignmentRequest_Status) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*PlanAssignmentRequest_Status))
}

func (o *PlanAssignmentRequest_Status) Clone() *PlanAssignmentRequest_Status {
	if o == nil {
		return nil
	}
	result := &PlanAssignmentRequest_Status{}
	result.Conclusion = o.Conclusion
	result.Reason = o.Reason
	return result
}

func (o *PlanAssignmentRequest_Status) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *PlanAssignmentRequest_Status) Merge(source *PlanAssignmentRequest_Status) {
	o.Conclusion = source.GetConclusion()
	o.Reason = source.GetReason()
}

func (o *PlanAssignmentRequest_Status) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*PlanAssignmentRequest_Status))
}

func (o *PlanAssignmentRequest_RequestType) GotenObjectExt() {}

func (o *PlanAssignmentRequest_RequestType) MakeFullFieldMask() *PlanAssignmentRequest_RequestType_FieldMask {
	return FullPlanAssignmentRequest_RequestType_FieldMask()
}

func (o *PlanAssignmentRequest_RequestType) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullPlanAssignmentRequest_RequestType_FieldMask()
}

func (o *PlanAssignmentRequest_RequestType) MakeDiffFieldMask(other *PlanAssignmentRequest_RequestType) *PlanAssignmentRequest_RequestType_FieldMask {
	if o == nil && other == nil {
		return &PlanAssignmentRequest_RequestType_FieldMask{}
	}
	if o == nil || other == nil {
		return FullPlanAssignmentRequest_RequestType_FieldMask()
	}

	res := &PlanAssignmentRequest_RequestType_FieldMask{}
	{
		_, leftSelected := o.Request.(*PlanAssignmentRequest_RequestType_Assign_)
		_, rightSelected := other.Request.(*PlanAssignmentRequest_RequestType_Assign_)
		if leftSelected == rightSelected {
			subMask := o.GetAssign().MakeDiffFieldMask(other.GetAssign())
			if subMask.IsFull() {
				res.Paths = append(res.Paths, &PlanAssignmentRequestRequestType_FieldTerminalPath{selector: PlanAssignmentRequestRequestType_FieldPathSelectorAssign})
			} else {
				for _, subpath := range subMask.Paths {
					res.Paths = append(res.Paths, &PlanAssignmentRequestRequestType_FieldSubPath{selector: PlanAssignmentRequestRequestType_FieldPathSelectorAssign, subPath: subpath})
				}
			}
		} else {
			res.Paths = append(res.Paths, &PlanAssignmentRequestRequestType_FieldTerminalPath{selector: PlanAssignmentRequestRequestType_FieldPathSelectorAssign})
		}
	}
	{
		_, leftSelected := o.Request.(*PlanAssignmentRequest_RequestType_Extend_)
		_, rightSelected := other.Request.(*PlanAssignmentRequest_RequestType_Extend_)
		if leftSelected == rightSelected {
			subMask := o.GetExtend().MakeDiffFieldMask(other.GetExtend())
			if subMask.IsFull() {
				res.Paths = append(res.Paths, &PlanAssignmentRequestRequestType_FieldTerminalPath{selector: PlanAssignmentRequestRequestType_FieldPathSelectorExtend})
			} else {
				for _, subpath := range subMask.Paths {
					res.Paths = append(res.Paths, &PlanAssignmentRequestRequestType_FieldSubPath{selector: PlanAssignmentRequestRequestType_FieldPathSelectorExtend, subPath: subpath})
				}
			}
		} else {
			res.Paths = append(res.Paths, &PlanAssignmentRequestRequestType_FieldTerminalPath{selector: PlanAssignmentRequestRequestType_FieldPathSelectorExtend})
		}
	}
	{
		_, leftSelected := o.Request.(*PlanAssignmentRequest_RequestType_Redistribute_)
		_, rightSelected := other.Request.(*PlanAssignmentRequest_RequestType_Redistribute_)
		if leftSelected == rightSelected {
			subMask := o.GetRedistribute().MakeDiffFieldMask(other.GetRedistribute())
			if subMask.IsFull() {
				res.Paths = append(res.Paths, &PlanAssignmentRequestRequestType_FieldTerminalPath{selector: PlanAssignmentRequestRequestType_FieldPathSelectorRedistribute})
			} else {
				for _, subpath := range subMask.Paths {
					res.Paths = append(res.Paths, &PlanAssignmentRequestRequestType_FieldSubPath{selector: PlanAssignmentRequestRequestType_FieldPathSelectorRedistribute, subPath: subpath})
				}
			}
		} else {
			res.Paths = append(res.Paths, &PlanAssignmentRequestRequestType_FieldTerminalPath{selector: PlanAssignmentRequestRequestType_FieldPathSelectorRedistribute})
		}
	}
	{
		_, leftSelected := o.Request.(*PlanAssignmentRequest_RequestType_Unassign_)
		_, rightSelected := other.Request.(*PlanAssignmentRequest_RequestType_Unassign_)
		if leftSelected == rightSelected {
			subMask := o.GetUnassign().MakeDiffFieldMask(other.GetUnassign())
			if subMask.IsFull() {
				res.Paths = append(res.Paths, &PlanAssignmentRequestRequestType_FieldTerminalPath{selector: PlanAssignmentRequestRequestType_FieldPathSelectorUnassign})
			} else {
				for _, subpath := range subMask.Paths {
					res.Paths = append(res.Paths, &PlanAssignmentRequestRequestType_FieldSubPath{selector: PlanAssignmentRequestRequestType_FieldPathSelectorUnassign, subPath: subpath})
				}
			}
		} else {
			res.Paths = append(res.Paths, &PlanAssignmentRequestRequestType_FieldTerminalPath{selector: PlanAssignmentRequestRequestType_FieldPathSelectorUnassign})
		}
	}
	return res
}

func (o *PlanAssignmentRequest_RequestType) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*PlanAssignmentRequest_RequestType))
}

func (o *PlanAssignmentRequest_RequestType) Clone() *PlanAssignmentRequest_RequestType {
	if o == nil {
		return nil
	}
	result := &PlanAssignmentRequest_RequestType{}
	if o, ok := o.Request.(*PlanAssignmentRequest_RequestType_Assign_); ok {
		result.Request = (*PlanAssignmentRequest_RequestType_Assign_)(nil)
		if o != nil {
			result.Request = &PlanAssignmentRequest_RequestType_Assign_{}
			result := result.Request.(*PlanAssignmentRequest_RequestType_Assign_)
			result.Assign = o.Assign.Clone()
		}
	}
	if o, ok := o.Request.(*PlanAssignmentRequest_RequestType_Extend_); ok {
		result.Request = (*PlanAssignmentRequest_RequestType_Extend_)(nil)
		if o != nil {
			result.Request = &PlanAssignmentRequest_RequestType_Extend_{}
			result := result.Request.(*PlanAssignmentRequest_RequestType_Extend_)
			result.Extend = o.Extend.Clone()
		}
	}
	if o, ok := o.Request.(*PlanAssignmentRequest_RequestType_Redistribute_); ok {
		result.Request = (*PlanAssignmentRequest_RequestType_Redistribute_)(nil)
		if o != nil {
			result.Request = &PlanAssignmentRequest_RequestType_Redistribute_{}
			result := result.Request.(*PlanAssignmentRequest_RequestType_Redistribute_)
			result.Redistribute = o.Redistribute.Clone()
		}
	}
	if o, ok := o.Request.(*PlanAssignmentRequest_RequestType_Unassign_); ok {
		result.Request = (*PlanAssignmentRequest_RequestType_Unassign_)(nil)
		if o != nil {
			result.Request = &PlanAssignmentRequest_RequestType_Unassign_{}
			result := result.Request.(*PlanAssignmentRequest_RequestType_Unassign_)
			result.Unassign = o.Unassign.Clone()
		}
	}
	return result
}

func (o *PlanAssignmentRequest_RequestType) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *PlanAssignmentRequest_RequestType) Merge(source *PlanAssignmentRequest_RequestType) {
	if source, ok := source.GetRequest().(*PlanAssignmentRequest_RequestType_Assign_); ok {
		if dstOneOf, ok := o.Request.(*PlanAssignmentRequest_RequestType_Assign_); !ok || dstOneOf == nil {
			o.Request = &PlanAssignmentRequest_RequestType_Assign_{}
		}
		if source != nil {
			o := o.Request.(*PlanAssignmentRequest_RequestType_Assign_)
			if source.Assign != nil {
				if o.Assign == nil {
					o.Assign = new(PlanAssignmentRequest_RequestType_Assign)
				}
				o.Assign.Merge(source.Assign)
			}
		}
	}
	if source, ok := source.GetRequest().(*PlanAssignmentRequest_RequestType_Extend_); ok {
		if dstOneOf, ok := o.Request.(*PlanAssignmentRequest_RequestType_Extend_); !ok || dstOneOf == nil {
			o.Request = &PlanAssignmentRequest_RequestType_Extend_{}
		}
		if source != nil {
			o := o.Request.(*PlanAssignmentRequest_RequestType_Extend_)
			if source.Extend != nil {
				if o.Extend == nil {
					o.Extend = new(PlanAssignmentRequest_RequestType_Extend)
				}
				o.Extend.Merge(source.Extend)
			}
		}
	}
	if source, ok := source.GetRequest().(*PlanAssignmentRequest_RequestType_Redistribute_); ok {
		if dstOneOf, ok := o.Request.(*PlanAssignmentRequest_RequestType_Redistribute_); !ok || dstOneOf == nil {
			o.Request = &PlanAssignmentRequest_RequestType_Redistribute_{}
		}
		if source != nil {
			o := o.Request.(*PlanAssignmentRequest_RequestType_Redistribute_)
			if source.Redistribute != nil {
				if o.Redistribute == nil {
					o.Redistribute = new(PlanAssignmentRequest_RequestType_Redistribute)
				}
				o.Redistribute.Merge(source.Redistribute)
			}
		}
	}
	if source, ok := source.GetRequest().(*PlanAssignmentRequest_RequestType_Unassign_); ok {
		if dstOneOf, ok := o.Request.(*PlanAssignmentRequest_RequestType_Unassign_); !ok || dstOneOf == nil {
			o.Request = &PlanAssignmentRequest_RequestType_Unassign_{}
		}
		if source != nil {
			o := o.Request.(*PlanAssignmentRequest_RequestType_Unassign_)
			if source.Unassign != nil {
				if o.Unassign == nil {
					o.Unassign = new(PlanAssignmentRequest_RequestType_Unassign)
				}
				o.Unassign.Merge(source.Unassign)
			}
		}
	}
}

func (o *PlanAssignmentRequest_RequestType) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*PlanAssignmentRequest_RequestType))
}

func (o *PlanAssignmentRequest_RequestType_Assign) GotenObjectExt() {}

func (o *PlanAssignmentRequest_RequestType_Assign) MakeFullFieldMask() *PlanAssignmentRequest_RequestType_Assign_FieldMask {
	return FullPlanAssignmentRequest_RequestType_Assign_FieldMask()
}

func (o *PlanAssignmentRequest_RequestType_Assign) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullPlanAssignmentRequest_RequestType_Assign_FieldMask()
}

func (o *PlanAssignmentRequest_RequestType_Assign) MakeDiffFieldMask(other *PlanAssignmentRequest_RequestType_Assign) *PlanAssignmentRequest_RequestType_Assign_FieldMask {
	if o == nil && other == nil {
		return &PlanAssignmentRequest_RequestType_Assign_FieldMask{}
	}
	if o == nil || other == nil {
		return FullPlanAssignmentRequest_RequestType_Assign_FieldMask()
	}

	res := &PlanAssignmentRequest_RequestType_Assign_FieldMask{}
	if o.GetPlan().String() != other.GetPlan().String() {
		res.Paths = append(res.Paths, &PlanAssignmentRequestRequestTypeAssign_FieldTerminalPath{selector: PlanAssignmentRequestRequestTypeAssign_FieldPathSelectorPlan})
	}

	if len(o.GetExtensions()) == len(other.GetExtensions()) {
		for i, lValue := range o.GetExtensions() {
			rValue := other.GetExtensions()[i]
			if len(lValue.MakeDiffFieldMask(rValue).Paths) > 0 {
				res.Paths = append(res.Paths, &PlanAssignmentRequestRequestTypeAssign_FieldTerminalPath{selector: PlanAssignmentRequestRequestTypeAssign_FieldPathSelectorExtensions})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &PlanAssignmentRequestRequestTypeAssign_FieldTerminalPath{selector: PlanAssignmentRequestRequestTypeAssign_FieldPathSelectorExtensions})
	}

	if len(o.GetRegionalDistributions()) == len(other.GetRegionalDistributions()) {
		for i, lValue := range o.GetRegionalDistributions() {
			rValue := other.GetRegionalDistributions()[i]
			if len(lValue.MakeDiffFieldMask(rValue).Paths) > 0 {
				res.Paths = append(res.Paths, &PlanAssignmentRequestRequestTypeAssign_FieldTerminalPath{selector: PlanAssignmentRequestRequestTypeAssign_FieldPathSelectorRegionalDistributions})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &PlanAssignmentRequestRequestTypeAssign_FieldTerminalPath{selector: PlanAssignmentRequestRequestTypeAssign_FieldPathSelectorRegionalDistributions})
	}
	return res
}

func (o *PlanAssignmentRequest_RequestType_Assign) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*PlanAssignmentRequest_RequestType_Assign))
}

func (o *PlanAssignmentRequest_RequestType_Assign) Clone() *PlanAssignmentRequest_RequestType_Assign {
	if o == nil {
		return nil
	}
	result := &PlanAssignmentRequest_RequestType_Assign{}
	if o.Plan == nil {
		result.Plan = nil
	} else if data, err := o.Plan.ProtoString(); err != nil {
		panic(err)
	} else {
		result.Plan = &plan.Reference{}
		if err := result.Plan.ParseProtoString(data); err != nil {
			panic(err)
		}
	}
	result.Extensions = make([]*common.Allowance, len(o.Extensions))
	for i, sourceValue := range o.Extensions {
		result.Extensions[i] = sourceValue.Clone()
	}
	result.RegionalDistributions = make([]*common.RegionalDistribution, len(o.RegionalDistributions))
	for i, sourceValue := range o.RegionalDistributions {
		result.RegionalDistributions[i] = sourceValue.Clone()
	}
	return result
}

func (o *PlanAssignmentRequest_RequestType_Assign) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *PlanAssignmentRequest_RequestType_Assign) Merge(source *PlanAssignmentRequest_RequestType_Assign) {
	if source.GetPlan() != nil {
		if data, err := source.GetPlan().ProtoString(); err != nil {
			panic(err)
		} else {
			o.Plan = &plan.Reference{}
			if err := o.Plan.ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	} else {
		o.Plan = nil
	}
	for _, sourceValue := range source.GetExtensions() {
		exists := false
		for _, currentValue := range o.Extensions {
			if proto.Equal(sourceValue, currentValue) {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement *common.Allowance
			if sourceValue != nil {
				newDstElement = new(common.Allowance)
				newDstElement.Merge(sourceValue)
			}
			o.Extensions = append(o.Extensions, newDstElement)
		}
	}

	for _, sourceValue := range source.GetRegionalDistributions() {
		exists := false
		for _, currentValue := range o.RegionalDistributions {
			if proto.Equal(sourceValue, currentValue) {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement *common.RegionalDistribution
			if sourceValue != nil {
				newDstElement = new(common.RegionalDistribution)
				newDstElement.Merge(sourceValue)
			}
			o.RegionalDistributions = append(o.RegionalDistributions, newDstElement)
		}
	}

}

func (o *PlanAssignmentRequest_RequestType_Assign) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*PlanAssignmentRequest_RequestType_Assign))
}

func (o *PlanAssignmentRequest_RequestType_Extend) GotenObjectExt() {}

func (o *PlanAssignmentRequest_RequestType_Extend) MakeFullFieldMask() *PlanAssignmentRequest_RequestType_Extend_FieldMask {
	return FullPlanAssignmentRequest_RequestType_Extend_FieldMask()
}

func (o *PlanAssignmentRequest_RequestType_Extend) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullPlanAssignmentRequest_RequestType_Extend_FieldMask()
}

func (o *PlanAssignmentRequest_RequestType_Extend) MakeDiffFieldMask(other *PlanAssignmentRequest_RequestType_Extend) *PlanAssignmentRequest_RequestType_Extend_FieldMask {
	if o == nil && other == nil {
		return &PlanAssignmentRequest_RequestType_Extend_FieldMask{}
	}
	if o == nil || other == nil {
		return FullPlanAssignmentRequest_RequestType_Extend_FieldMask()
	}

	res := &PlanAssignmentRequest_RequestType_Extend_FieldMask{}
	if o.GetAssignment().String() != other.GetAssignment().String() {
		res.Paths = append(res.Paths, &PlanAssignmentRequestRequestTypeExtend_FieldTerminalPath{selector: PlanAssignmentRequestRequestTypeExtend_FieldPathSelectorAssignment})
	}

	if len(o.GetAdditions()) == len(other.GetAdditions()) {
		for i, lValue := range o.GetAdditions() {
			rValue := other.GetAdditions()[i]
			if len(lValue.MakeDiffFieldMask(rValue).Paths) > 0 {
				res.Paths = append(res.Paths, &PlanAssignmentRequestRequestTypeExtend_FieldTerminalPath{selector: PlanAssignmentRequestRequestTypeExtend_FieldPathSelectorAdditions})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &PlanAssignmentRequestRequestTypeExtend_FieldTerminalPath{selector: PlanAssignmentRequestRequestTypeExtend_FieldPathSelectorAdditions})
	}

	if len(o.GetRegionalDistributions()) == len(other.GetRegionalDistributions()) {
		for i, lValue := range o.GetRegionalDistributions() {
			rValue := other.GetRegionalDistributions()[i]
			if len(lValue.MakeDiffFieldMask(rValue).Paths) > 0 {
				res.Paths = append(res.Paths, &PlanAssignmentRequestRequestTypeExtend_FieldTerminalPath{selector: PlanAssignmentRequestRequestTypeExtend_FieldPathSelectorRegionalDistributions})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &PlanAssignmentRequestRequestTypeExtend_FieldTerminalPath{selector: PlanAssignmentRequestRequestTypeExtend_FieldPathSelectorRegionalDistributions})
	}
	return res
}

func (o *PlanAssignmentRequest_RequestType_Extend) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*PlanAssignmentRequest_RequestType_Extend))
}

func (o *PlanAssignmentRequest_RequestType_Extend) Clone() *PlanAssignmentRequest_RequestType_Extend {
	if o == nil {
		return nil
	}
	result := &PlanAssignmentRequest_RequestType_Extend{}
	if o.Assignment == nil {
		result.Assignment = nil
	} else if data, err := o.Assignment.ProtoString(); err != nil {
		panic(err)
	} else {
		result.Assignment = &plan_assignment.Reference{}
		if err := result.Assignment.ParseProtoString(data); err != nil {
			panic(err)
		}
	}
	result.Additions = make([]*common.Allowance, len(o.Additions))
	for i, sourceValue := range o.Additions {
		result.Additions[i] = sourceValue.Clone()
	}
	result.RegionalDistributions = make([]*common.RegionalDistribution, len(o.RegionalDistributions))
	for i, sourceValue := range o.RegionalDistributions {
		result.RegionalDistributions[i] = sourceValue.Clone()
	}
	return result
}

func (o *PlanAssignmentRequest_RequestType_Extend) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *PlanAssignmentRequest_RequestType_Extend) Merge(source *PlanAssignmentRequest_RequestType_Extend) {
	if source.GetAssignment() != nil {
		if data, err := source.GetAssignment().ProtoString(); err != nil {
			panic(err)
		} else {
			o.Assignment = &plan_assignment.Reference{}
			if err := o.Assignment.ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	} else {
		o.Assignment = nil
	}
	for _, sourceValue := range source.GetAdditions() {
		exists := false
		for _, currentValue := range o.Additions {
			if proto.Equal(sourceValue, currentValue) {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement *common.Allowance
			if sourceValue != nil {
				newDstElement = new(common.Allowance)
				newDstElement.Merge(sourceValue)
			}
			o.Additions = append(o.Additions, newDstElement)
		}
	}

	for _, sourceValue := range source.GetRegionalDistributions() {
		exists := false
		for _, currentValue := range o.RegionalDistributions {
			if proto.Equal(sourceValue, currentValue) {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement *common.RegionalDistribution
			if sourceValue != nil {
				newDstElement = new(common.RegionalDistribution)
				newDstElement.Merge(sourceValue)
			}
			o.RegionalDistributions = append(o.RegionalDistributions, newDstElement)
		}
	}

}

func (o *PlanAssignmentRequest_RequestType_Extend) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*PlanAssignmentRequest_RequestType_Extend))
}

func (o *PlanAssignmentRequest_RequestType_Redistribute) GotenObjectExt() {}

func (o *PlanAssignmentRequest_RequestType_Redistribute) MakeFullFieldMask() *PlanAssignmentRequest_RequestType_Redistribute_FieldMask {
	return FullPlanAssignmentRequest_RequestType_Redistribute_FieldMask()
}

func (o *PlanAssignmentRequest_RequestType_Redistribute) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullPlanAssignmentRequest_RequestType_Redistribute_FieldMask()
}

func (o *PlanAssignmentRequest_RequestType_Redistribute) MakeDiffFieldMask(other *PlanAssignmentRequest_RequestType_Redistribute) *PlanAssignmentRequest_RequestType_Redistribute_FieldMask {
	if o == nil && other == nil {
		return &PlanAssignmentRequest_RequestType_Redistribute_FieldMask{}
	}
	if o == nil || other == nil {
		return FullPlanAssignmentRequest_RequestType_Redistribute_FieldMask()
	}

	res := &PlanAssignmentRequest_RequestType_Redistribute_FieldMask{}
	if o.GetAssignment().String() != other.GetAssignment().String() {
		res.Paths = append(res.Paths, &PlanAssignmentRequestRequestTypeRedistribute_FieldTerminalPath{selector: PlanAssignmentRequestRequestTypeRedistribute_FieldPathSelectorAssignment})
	}

	if len(o.GetRegionalDistributions()) == len(other.GetRegionalDistributions()) {
		for i, lValue := range o.GetRegionalDistributions() {
			rValue := other.GetRegionalDistributions()[i]
			if len(lValue.MakeDiffFieldMask(rValue).Paths) > 0 {
				res.Paths = append(res.Paths, &PlanAssignmentRequestRequestTypeRedistribute_FieldTerminalPath{selector: PlanAssignmentRequestRequestTypeRedistribute_FieldPathSelectorRegionalDistributions})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &PlanAssignmentRequestRequestTypeRedistribute_FieldTerminalPath{selector: PlanAssignmentRequestRequestTypeRedistribute_FieldPathSelectorRegionalDistributions})
	}
	return res
}

func (o *PlanAssignmentRequest_RequestType_Redistribute) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*PlanAssignmentRequest_RequestType_Redistribute))
}

func (o *PlanAssignmentRequest_RequestType_Redistribute) Clone() *PlanAssignmentRequest_RequestType_Redistribute {
	if o == nil {
		return nil
	}
	result := &PlanAssignmentRequest_RequestType_Redistribute{}
	if o.Assignment == nil {
		result.Assignment = nil
	} else if data, err := o.Assignment.ProtoString(); err != nil {
		panic(err)
	} else {
		result.Assignment = &plan_assignment.Reference{}
		if err := result.Assignment.ParseProtoString(data); err != nil {
			panic(err)
		}
	}
	result.RegionalDistributions = make([]*common.RegionalDistribution, len(o.RegionalDistributions))
	for i, sourceValue := range o.RegionalDistributions {
		result.RegionalDistributions[i] = sourceValue.Clone()
	}
	return result
}

func (o *PlanAssignmentRequest_RequestType_Redistribute) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *PlanAssignmentRequest_RequestType_Redistribute) Merge(source *PlanAssignmentRequest_RequestType_Redistribute) {
	if source.GetAssignment() != nil {
		if data, err := source.GetAssignment().ProtoString(); err != nil {
			panic(err)
		} else {
			o.Assignment = &plan_assignment.Reference{}
			if err := o.Assignment.ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	} else {
		o.Assignment = nil
	}
	for _, sourceValue := range source.GetRegionalDistributions() {
		exists := false
		for _, currentValue := range o.RegionalDistributions {
			if proto.Equal(sourceValue, currentValue) {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement *common.RegionalDistribution
			if sourceValue != nil {
				newDstElement = new(common.RegionalDistribution)
				newDstElement.Merge(sourceValue)
			}
			o.RegionalDistributions = append(o.RegionalDistributions, newDstElement)
		}
	}

}

func (o *PlanAssignmentRequest_RequestType_Redistribute) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*PlanAssignmentRequest_RequestType_Redistribute))
}

func (o *PlanAssignmentRequest_RequestType_Unassign) GotenObjectExt() {}

func (o *PlanAssignmentRequest_RequestType_Unassign) MakeFullFieldMask() *PlanAssignmentRequest_RequestType_Unassign_FieldMask {
	return FullPlanAssignmentRequest_RequestType_Unassign_FieldMask()
}

func (o *PlanAssignmentRequest_RequestType_Unassign) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullPlanAssignmentRequest_RequestType_Unassign_FieldMask()
}

func (o *PlanAssignmentRequest_RequestType_Unassign) MakeDiffFieldMask(other *PlanAssignmentRequest_RequestType_Unassign) *PlanAssignmentRequest_RequestType_Unassign_FieldMask {
	if o == nil && other == nil {
		return &PlanAssignmentRequest_RequestType_Unassign_FieldMask{}
	}
	if o == nil || other == nil {
		return FullPlanAssignmentRequest_RequestType_Unassign_FieldMask()
	}

	res := &PlanAssignmentRequest_RequestType_Unassign_FieldMask{}
	if o.GetAssignment().String() != other.GetAssignment().String() {
		res.Paths = append(res.Paths, &PlanAssignmentRequestRequestTypeUnassign_FieldTerminalPath{selector: PlanAssignmentRequestRequestTypeUnassign_FieldPathSelectorAssignment})
	}
	return res
}

func (o *PlanAssignmentRequest_RequestType_Unassign) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*PlanAssignmentRequest_RequestType_Unassign))
}

func (o *PlanAssignmentRequest_RequestType_Unassign) Clone() *PlanAssignmentRequest_RequestType_Unassign {
	if o == nil {
		return nil
	}
	result := &PlanAssignmentRequest_RequestType_Unassign{}
	if o.Assignment == nil {
		result.Assignment = nil
	} else if data, err := o.Assignment.ProtoString(); err != nil {
		panic(err)
	} else {
		result.Assignment = &plan_assignment.Reference{}
		if err := result.Assignment.ParseProtoString(data); err != nil {
			panic(err)
		}
	}
	return result
}

func (o *PlanAssignmentRequest_RequestType_Unassign) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *PlanAssignmentRequest_RequestType_Unassign) Merge(source *PlanAssignmentRequest_RequestType_Unassign) {
	if source.GetAssignment() != nil {
		if data, err := source.GetAssignment().ProtoString(); err != nil {
			panic(err)
		} else {
			o.Assignment = &plan_assignment.Reference{}
			if err := o.Assignment.ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	} else {
		o.Assignment = nil
	}
}

func (o *PlanAssignmentRequest_RequestType_Unassign) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*PlanAssignmentRequest_RequestType_Unassign))
}
