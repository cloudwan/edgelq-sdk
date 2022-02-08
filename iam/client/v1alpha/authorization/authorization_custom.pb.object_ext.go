// Code generated by protoc-gen-goten-object
// File: edgelq/iam/proto/v1alpha/authorization_custom.proto
// DO NOT EDIT!!!

package authorization_client

import (
	"fmt"
	"sort"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
)

// proto imports
import (
	condition "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/condition"
	permission "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/permission"
)

// ensure the imports are used
var (
	_ = fmt.Stringer(nil)
	_ = sort.Interface(nil)

	_ = proto.Message(nil)
	_ = fieldmaskpb.FieldMask{}

	_ = gotenobject.FieldPath(nil)
)

// make sure we're using proto imports
var (
	_ = &condition.Condition{}
	_ = &permission.Permission{}
)

func (o *Check) GotenObjectExt() {}

func (o *Check) MakeFullFieldMask() *Check_FieldMask {
	return FullCheck_FieldMask()
}

func (o *Check) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullCheck_FieldMask()
}

func (o *Check) MakeDiffFieldMask(other *Check) *Check_FieldMask {
	if o == nil && other == nil {
		return &Check_FieldMask{}
	}
	if o == nil || other == nil {
		return FullCheck_FieldMask()
	}

	res := &Check_FieldMask{}
	if o.GetObject() != other.GetObject() {
		res.Paths = append(res.Paths, &Check_FieldTerminalPath{selector: Check_FieldPathSelectorObject})
	}

	if len(o.GetPermissions()) == len(other.GetPermissions()) {
		for i, lValue := range o.GetPermissions() {
			rValue := other.GetPermissions()[i]
			if lValue.String() != rValue.String() {
				res.Paths = append(res.Paths, &Check_FieldTerminalPath{selector: Check_FieldPathSelectorPermissions})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &Check_FieldTerminalPath{selector: Check_FieldPathSelectorPermissions})
	}
	return res
}

func (o *Check) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*Check))
}

func (o *Check) Clone() *Check {
	if o == nil {
		return nil
	}
	result := &Check{}
	result.Object = o.Object
	result.Permissions = make([]*permission.Reference, len(o.Permissions))
	for i, sourceValue := range o.Permissions {
		if sourceValue == nil {
			result.Permissions[i] = nil
		} else if data, err := sourceValue.ProtoString(); err != nil {
			panic(err)
		} else {
			result.Permissions[i] = &permission.Reference{}
			if err := result.Permissions[i].ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	}
	return result
}

func (o *Check) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *Check) Merge(source *Check) {
	o.Object = source.GetObject()
	for _, sourceValue := range source.GetPermissions() {
		exists := false
		for _, currentValue := range o.Permissions {
			leftProtoStr, _ := currentValue.ProtoString()
			rightProtoStr, _ := sourceValue.ProtoString()
			if leftProtoStr == rightProtoStr {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement *permission.Reference
			if sourceValue != nil {
				if data, err := sourceValue.ProtoString(); err != nil {
					panic(err)
				} else {
					newDstElement = &permission.Reference{}
					if err := newDstElement.ParseProtoString(data); err != nil {
						panic(err)
					}
				}
			}
			o.Permissions = append(o.Permissions, newDstElement)
		}
	}

}

func (o *Check) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*Check))
}

func (o *ConditionalGrant) GotenObjectExt() {}

func (o *ConditionalGrant) MakeFullFieldMask() *ConditionalGrant_FieldMask {
	return FullConditionalGrant_FieldMask()
}

func (o *ConditionalGrant) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullConditionalGrant_FieldMask()
}

func (o *ConditionalGrant) MakeDiffFieldMask(other *ConditionalGrant) *ConditionalGrant_FieldMask {
	if o == nil && other == nil {
		return &ConditionalGrant_FieldMask{}
	}
	if o == nil || other == nil {
		return FullConditionalGrant_FieldMask()
	}

	res := &ConditionalGrant_FieldMask{}

	if len(o.GetPermissions()) == len(other.GetPermissions()) {
		for i, lValue := range o.GetPermissions() {
			rValue := other.GetPermissions()[i]
			if lValue.String() != rValue.String() {
				res.Paths = append(res.Paths, &ConditionalGrant_FieldTerminalPath{selector: ConditionalGrant_FieldPathSelectorPermissions})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &ConditionalGrant_FieldTerminalPath{selector: ConditionalGrant_FieldPathSelectorPermissions})
	}

	if len(o.GetConditionBindings()) == len(other.GetConditionBindings()) {
		for i, lValue := range o.GetConditionBindings() {
			rValue := other.GetConditionBindings()[i]
			if len(lValue.MakeDiffFieldMask(rValue).Paths) > 0 {
				res.Paths = append(res.Paths, &ConditionalGrant_FieldTerminalPath{selector: ConditionalGrant_FieldPathSelectorConditionBindings})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &ConditionalGrant_FieldTerminalPath{selector: ConditionalGrant_FieldPathSelectorConditionBindings})
	}
	return res
}

func (o *ConditionalGrant) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*ConditionalGrant))
}

func (o *ConditionalGrant) Clone() *ConditionalGrant {
	if o == nil {
		return nil
	}
	result := &ConditionalGrant{}
	result.Permissions = make([]*permission.Reference, len(o.Permissions))
	for i, sourceValue := range o.Permissions {
		if sourceValue == nil {
			result.Permissions[i] = nil
		} else if data, err := sourceValue.ProtoString(); err != nil {
			panic(err)
		} else {
			result.Permissions[i] = &permission.Reference{}
			if err := result.Permissions[i].ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	}
	result.ConditionBindings = make([]*condition.ConditionBinding, len(o.ConditionBindings))
	for i, sourceValue := range o.ConditionBindings {
		result.ConditionBindings[i] = sourceValue.Clone()
	}
	return result
}

func (o *ConditionalGrant) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *ConditionalGrant) Merge(source *ConditionalGrant) {
	for _, sourceValue := range source.GetPermissions() {
		exists := false
		for _, currentValue := range o.Permissions {
			leftProtoStr, _ := currentValue.ProtoString()
			rightProtoStr, _ := sourceValue.ProtoString()
			if leftProtoStr == rightProtoStr {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement *permission.Reference
			if sourceValue != nil {
				if data, err := sourceValue.ProtoString(); err != nil {
					panic(err)
				} else {
					newDstElement = &permission.Reference{}
					if err := newDstElement.ParseProtoString(data); err != nil {
						panic(err)
					}
				}
			}
			o.Permissions = append(o.Permissions, newDstElement)
		}
	}

	for _, sourceValue := range source.GetConditionBindings() {
		exists := false
		for _, currentValue := range o.ConditionBindings {
			if proto.Equal(sourceValue, currentValue) {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement *condition.ConditionBinding
			if sourceValue != nil {
				newDstElement = new(condition.ConditionBinding)
				newDstElement.Merge(sourceValue)
			}
			o.ConditionBindings = append(o.ConditionBindings, newDstElement)
		}
	}

}

func (o *ConditionalGrant) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*ConditionalGrant))
}

func (o *CheckResult) GotenObjectExt() {}

func (o *CheckResult) MakeFullFieldMask() *CheckResult_FieldMask {
	return FullCheckResult_FieldMask()
}

func (o *CheckResult) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullCheckResult_FieldMask()
}

func (o *CheckResult) MakeDiffFieldMask(other *CheckResult) *CheckResult_FieldMask {
	if o == nil && other == nil {
		return &CheckResult_FieldMask{}
	}
	if o == nil || other == nil {
		return FullCheckResult_FieldMask()
	}

	res := &CheckResult_FieldMask{}
	if o.GetObject() != other.GetObject() {
		res.Paths = append(res.Paths, &CheckResult_FieldTerminalPath{selector: CheckResult_FieldPathSelectorObject})
	}

	if len(o.GetGrantedPermissions()) == len(other.GetGrantedPermissions()) {
		for i, lValue := range o.GetGrantedPermissions() {
			rValue := other.GetGrantedPermissions()[i]
			if lValue.String() != rValue.String() {
				res.Paths = append(res.Paths, &CheckResult_FieldTerminalPath{selector: CheckResult_FieldPathSelectorGrantedPermissions})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &CheckResult_FieldTerminalPath{selector: CheckResult_FieldPathSelectorGrantedPermissions})
	}

	if len(o.GetConditionallyGrantedPermissions()) == len(other.GetConditionallyGrantedPermissions()) {
		for i, lValue := range o.GetConditionallyGrantedPermissions() {
			rValue := other.GetConditionallyGrantedPermissions()[i]
			if len(lValue.MakeDiffFieldMask(rValue).Paths) > 0 {
				res.Paths = append(res.Paths, &CheckResult_FieldTerminalPath{selector: CheckResult_FieldPathSelectorConditionallyGrantedPermissions})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &CheckResult_FieldTerminalPath{selector: CheckResult_FieldPathSelectorConditionallyGrantedPermissions})
	}
	return res
}

func (o *CheckResult) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*CheckResult))
}

func (o *CheckResult) Clone() *CheckResult {
	if o == nil {
		return nil
	}
	result := &CheckResult{}
	result.Object = o.Object
	result.GrantedPermissions = make([]*permission.Reference, len(o.GrantedPermissions))
	for i, sourceValue := range o.GrantedPermissions {
		if sourceValue == nil {
			result.GrantedPermissions[i] = nil
		} else if data, err := sourceValue.ProtoString(); err != nil {
			panic(err)
		} else {
			result.GrantedPermissions[i] = &permission.Reference{}
			if err := result.GrantedPermissions[i].ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	}
	result.ConditionallyGrantedPermissions = make([]*ConditionalGrant, len(o.ConditionallyGrantedPermissions))
	for i, sourceValue := range o.ConditionallyGrantedPermissions {
		result.ConditionallyGrantedPermissions[i] = sourceValue.Clone()
	}
	return result
}

func (o *CheckResult) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *CheckResult) Merge(source *CheckResult) {
	o.Object = source.GetObject()
	for _, sourceValue := range source.GetGrantedPermissions() {
		exists := false
		for _, currentValue := range o.GrantedPermissions {
			leftProtoStr, _ := currentValue.ProtoString()
			rightProtoStr, _ := sourceValue.ProtoString()
			if leftProtoStr == rightProtoStr {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement *permission.Reference
			if sourceValue != nil {
				if data, err := sourceValue.ProtoString(); err != nil {
					panic(err)
				} else {
					newDstElement = &permission.Reference{}
					if err := newDstElement.ParseProtoString(data); err != nil {
						panic(err)
					}
				}
			}
			o.GrantedPermissions = append(o.GrantedPermissions, newDstElement)
		}
	}

	for _, sourceValue := range source.GetConditionallyGrantedPermissions() {
		exists := false
		for _, currentValue := range o.ConditionallyGrantedPermissions {
			if proto.Equal(sourceValue, currentValue) {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement *ConditionalGrant
			if sourceValue != nil {
				newDstElement = new(ConditionalGrant)
				newDstElement.Merge(sourceValue)
			}
			o.ConditionallyGrantedPermissions = append(o.ConditionallyGrantedPermissions, newDstElement)
		}
	}

}

func (o *CheckResult) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*CheckResult))
}

func (o *CheckPermissionsRequest) GotenObjectExt() {}

func (o *CheckPermissionsRequest) MakeFullFieldMask() *CheckPermissionsRequest_FieldMask {
	return FullCheckPermissionsRequest_FieldMask()
}

func (o *CheckPermissionsRequest) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullCheckPermissionsRequest_FieldMask()
}

func (o *CheckPermissionsRequest) MakeDiffFieldMask(other *CheckPermissionsRequest) *CheckPermissionsRequest_FieldMask {
	if o == nil && other == nil {
		return &CheckPermissionsRequest_FieldMask{}
	}
	if o == nil || other == nil {
		return FullCheckPermissionsRequest_FieldMask()
	}

	res := &CheckPermissionsRequest_FieldMask{}
	if o.GetMember() != other.GetMember() {
		res.Paths = append(res.Paths, &CheckPermissionsRequest_FieldTerminalPath{selector: CheckPermissionsRequest_FieldPathSelectorMember})
	}

	if len(o.GetChecks()) == len(other.GetChecks()) {
		for i, lValue := range o.GetChecks() {
			rValue := other.GetChecks()[i]
			if len(lValue.MakeDiffFieldMask(rValue).Paths) > 0 {
				res.Paths = append(res.Paths, &CheckPermissionsRequest_FieldTerminalPath{selector: CheckPermissionsRequest_FieldPathSelectorChecks})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &CheckPermissionsRequest_FieldTerminalPath{selector: CheckPermissionsRequest_FieldPathSelectorChecks})
	}
	if o.GetSkipCache() != other.GetSkipCache() {
		res.Paths = append(res.Paths, &CheckPermissionsRequest_FieldTerminalPath{selector: CheckPermissionsRequest_FieldPathSelectorSkipCache})
	}
	return res
}

func (o *CheckPermissionsRequest) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*CheckPermissionsRequest))
}

func (o *CheckPermissionsRequest) Clone() *CheckPermissionsRequest {
	if o == nil {
		return nil
	}
	result := &CheckPermissionsRequest{}
	result.Member = o.Member
	result.Checks = make([]*Check, len(o.Checks))
	for i, sourceValue := range o.Checks {
		result.Checks[i] = sourceValue.Clone()
	}
	result.SkipCache = o.SkipCache
	return result
}

func (o *CheckPermissionsRequest) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *CheckPermissionsRequest) Merge(source *CheckPermissionsRequest) {
	o.Member = source.GetMember()
	for _, sourceValue := range source.GetChecks() {
		exists := false
		for _, currentValue := range o.Checks {
			if proto.Equal(sourceValue, currentValue) {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement *Check
			if sourceValue != nil {
				newDstElement = new(Check)
				newDstElement.Merge(sourceValue)
			}
			o.Checks = append(o.Checks, newDstElement)
		}
	}

	o.SkipCache = source.GetSkipCache()
}

func (o *CheckPermissionsRequest) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*CheckPermissionsRequest))
}

func (o *CheckPermissionsResponse) GotenObjectExt() {}

func (o *CheckPermissionsResponse) MakeFullFieldMask() *CheckPermissionsResponse_FieldMask {
	return FullCheckPermissionsResponse_FieldMask()
}

func (o *CheckPermissionsResponse) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullCheckPermissionsResponse_FieldMask()
}

func (o *CheckPermissionsResponse) MakeDiffFieldMask(other *CheckPermissionsResponse) *CheckPermissionsResponse_FieldMask {
	if o == nil && other == nil {
		return &CheckPermissionsResponse_FieldMask{}
	}
	if o == nil || other == nil {
		return FullCheckPermissionsResponse_FieldMask()
	}

	res := &CheckPermissionsResponse_FieldMask{}

	if len(o.GetCheckResults()) == len(other.GetCheckResults()) {
		for i, lValue := range o.GetCheckResults() {
			rValue := other.GetCheckResults()[i]
			if len(lValue.MakeDiffFieldMask(rValue).Paths) > 0 {
				res.Paths = append(res.Paths, &CheckPermissionsResponse_FieldTerminalPath{selector: CheckPermissionsResponse_FieldPathSelectorCheckResults})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &CheckPermissionsResponse_FieldTerminalPath{selector: CheckPermissionsResponse_FieldPathSelectorCheckResults})
	}
	return res
}

func (o *CheckPermissionsResponse) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*CheckPermissionsResponse))
}

func (o *CheckPermissionsResponse) Clone() *CheckPermissionsResponse {
	if o == nil {
		return nil
	}
	result := &CheckPermissionsResponse{}
	result.CheckResults = make([]*CheckResult, len(o.CheckResults))
	for i, sourceValue := range o.CheckResults {
		result.CheckResults[i] = sourceValue.Clone()
	}
	return result
}

func (o *CheckPermissionsResponse) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *CheckPermissionsResponse) Merge(source *CheckPermissionsResponse) {
	for _, sourceValue := range source.GetCheckResults() {
		exists := false
		for _, currentValue := range o.CheckResults {
			if proto.Equal(sourceValue, currentValue) {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement *CheckResult
			if sourceValue != nil {
				newDstElement = new(CheckResult)
				newDstElement.Merge(sourceValue)
			}
			o.CheckResults = append(o.CheckResults, newDstElement)
		}
	}

}

func (o *CheckPermissionsResponse) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*CheckPermissionsResponse))
}

func (o *CheckMyPermissionsRequest) GotenObjectExt() {}

func (o *CheckMyPermissionsRequest) MakeFullFieldMask() *CheckMyPermissionsRequest_FieldMask {
	return FullCheckMyPermissionsRequest_FieldMask()
}

func (o *CheckMyPermissionsRequest) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullCheckMyPermissionsRequest_FieldMask()
}

func (o *CheckMyPermissionsRequest) MakeDiffFieldMask(other *CheckMyPermissionsRequest) *CheckMyPermissionsRequest_FieldMask {
	if o == nil && other == nil {
		return &CheckMyPermissionsRequest_FieldMask{}
	}
	if o == nil || other == nil {
		return FullCheckMyPermissionsRequest_FieldMask()
	}

	res := &CheckMyPermissionsRequest_FieldMask{}

	if len(o.GetChecks()) == len(other.GetChecks()) {
		for i, lValue := range o.GetChecks() {
			rValue := other.GetChecks()[i]
			if len(lValue.MakeDiffFieldMask(rValue).Paths) > 0 {
				res.Paths = append(res.Paths, &CheckMyPermissionsRequest_FieldTerminalPath{selector: CheckMyPermissionsRequest_FieldPathSelectorChecks})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &CheckMyPermissionsRequest_FieldTerminalPath{selector: CheckMyPermissionsRequest_FieldPathSelectorChecks})
	}
	if o.GetSkipCache() != other.GetSkipCache() {
		res.Paths = append(res.Paths, &CheckMyPermissionsRequest_FieldTerminalPath{selector: CheckMyPermissionsRequest_FieldPathSelectorSkipCache})
	}
	return res
}

func (o *CheckMyPermissionsRequest) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*CheckMyPermissionsRequest))
}

func (o *CheckMyPermissionsRequest) Clone() *CheckMyPermissionsRequest {
	if o == nil {
		return nil
	}
	result := &CheckMyPermissionsRequest{}
	result.Checks = make([]*Check, len(o.Checks))
	for i, sourceValue := range o.Checks {
		result.Checks[i] = sourceValue.Clone()
	}
	result.SkipCache = o.SkipCache
	return result
}

func (o *CheckMyPermissionsRequest) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *CheckMyPermissionsRequest) Merge(source *CheckMyPermissionsRequest) {
	for _, sourceValue := range source.GetChecks() {
		exists := false
		for _, currentValue := range o.Checks {
			if proto.Equal(sourceValue, currentValue) {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement *Check
			if sourceValue != nil {
				newDstElement = new(Check)
				newDstElement.Merge(sourceValue)
			}
			o.Checks = append(o.Checks, newDstElement)
		}
	}

	o.SkipCache = source.GetSkipCache()
}

func (o *CheckMyPermissionsRequest) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*CheckMyPermissionsRequest))
}

func (o *CheckMyPermissionsResponse) GotenObjectExt() {}

func (o *CheckMyPermissionsResponse) MakeFullFieldMask() *CheckMyPermissionsResponse_FieldMask {
	return FullCheckMyPermissionsResponse_FieldMask()
}

func (o *CheckMyPermissionsResponse) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullCheckMyPermissionsResponse_FieldMask()
}

func (o *CheckMyPermissionsResponse) MakeDiffFieldMask(other *CheckMyPermissionsResponse) *CheckMyPermissionsResponse_FieldMask {
	if o == nil && other == nil {
		return &CheckMyPermissionsResponse_FieldMask{}
	}
	if o == nil || other == nil {
		return FullCheckMyPermissionsResponse_FieldMask()
	}

	res := &CheckMyPermissionsResponse_FieldMask{}

	if len(o.GetCheckResults()) == len(other.GetCheckResults()) {
		for i, lValue := range o.GetCheckResults() {
			rValue := other.GetCheckResults()[i]
			if len(lValue.MakeDiffFieldMask(rValue).Paths) > 0 {
				res.Paths = append(res.Paths, &CheckMyPermissionsResponse_FieldTerminalPath{selector: CheckMyPermissionsResponse_FieldPathSelectorCheckResults})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &CheckMyPermissionsResponse_FieldTerminalPath{selector: CheckMyPermissionsResponse_FieldPathSelectorCheckResults})
	}
	return res
}

func (o *CheckMyPermissionsResponse) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*CheckMyPermissionsResponse))
}

func (o *CheckMyPermissionsResponse) Clone() *CheckMyPermissionsResponse {
	if o == nil {
		return nil
	}
	result := &CheckMyPermissionsResponse{}
	result.CheckResults = make([]*CheckResult, len(o.CheckResults))
	for i, sourceValue := range o.CheckResults {
		result.CheckResults[i] = sourceValue.Clone()
	}
	return result
}

func (o *CheckMyPermissionsResponse) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *CheckMyPermissionsResponse) Merge(source *CheckMyPermissionsResponse) {
	for _, sourceValue := range source.GetCheckResults() {
		exists := false
		for _, currentValue := range o.CheckResults {
			if proto.Equal(sourceValue, currentValue) {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement *CheckResult
			if sourceValue != nil {
				newDstElement = new(CheckResult)
				newDstElement.Merge(sourceValue)
			}
			o.CheckResults = append(o.CheckResults, newDstElement)
		}
	}

}

func (o *CheckMyPermissionsResponse) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*CheckMyPermissionsResponse))
}