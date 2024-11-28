// Code generated by protoc-gen-goten-object
// File: edgelq/devices/proto/v1/provisioning_policy.proto
// DO NOT EDIT!!!

package provisioning_policy

import (
	"fmt"
	"sort"

	"google.golang.org/protobuf/proto"
	googlefieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"

	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
)

// proto imports
import (
	device "github.com/cloudwan/edgelq-sdk/devices/resources/v1/device"
	project "github.com/cloudwan/edgelq-sdk/devices/resources/v1/project"
	iam_condition "github.com/cloudwan/edgelq-sdk/iam/resources/v1/condition"
	iam_role "github.com/cloudwan/edgelq-sdk/iam/resources/v1/role"
	iam_service_account "github.com/cloudwan/edgelq-sdk/iam/resources/v1/service_account"
	meta "github.com/cloudwan/goten-sdk/types/meta"
	structpb "google.golang.org/protobuf/types/known/structpb"
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
	_ = &device.Device{}
	_ = &project.Project{}
	_ = &iam_condition.Condition{}
	_ = &iam_role.Role{}
	_ = &iam_service_account.ServiceAccount{}
	_ = &structpb.Struct{}
	_ = &meta.Meta{}
)

func (o *ProvisioningPolicy) GotenObjectExt() {}

func (o *ProvisioningPolicy) MakeFullFieldMask() *ProvisioningPolicy_FieldMask {
	return FullProvisioningPolicy_FieldMask()
}

func (o *ProvisioningPolicy) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullProvisioningPolicy_FieldMask()
}

func (o *ProvisioningPolicy) MakeDiffFieldMask(other *ProvisioningPolicy) *ProvisioningPolicy_FieldMask {
	if o == nil && other == nil {
		return &ProvisioningPolicy_FieldMask{}
	}
	if o == nil || other == nil {
		return FullProvisioningPolicy_FieldMask()
	}

	res := &ProvisioningPolicy_FieldMask{}
	if o.GetName().String() != other.GetName().String() {
		res.Paths = append(res.Paths, &ProvisioningPolicy_FieldTerminalPath{selector: ProvisioningPolicy_FieldPathSelectorName})
	}
	{
		subMask := o.GetMetadata().MakeDiffFieldMask(other.GetMetadata())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &ProvisioningPolicy_FieldTerminalPath{selector: ProvisioningPolicy_FieldPathSelectorMetadata})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &ProvisioningPolicy_FieldSubPath{selector: ProvisioningPolicy_FieldPathSelectorMetadata, subPath: subpath})
			}
		}
	}
	if o.GetDisplayName() != other.GetDisplayName() {
		res.Paths = append(res.Paths, &ProvisioningPolicy_FieldTerminalPath{selector: ProvisioningPolicy_FieldPathSelectorDisplayName})
	}
	if o.GetDescription() != other.GetDescription() {
		res.Paths = append(res.Paths, &ProvisioningPolicy_FieldTerminalPath{selector: ProvisioningPolicy_FieldPathSelectorDescription})
	}
	{
		subMask := o.GetSpec().MakeDiffFieldMask(other.GetSpec())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &ProvisioningPolicy_FieldTerminalPath{selector: ProvisioningPolicy_FieldPathSelectorSpec})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &ProvisioningPolicy_FieldSubPath{selector: ProvisioningPolicy_FieldPathSelectorSpec, subPath: subpath})
			}
		}
	}
	{
		subMask := o.GetStatus().MakeDiffFieldMask(other.GetStatus())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &ProvisioningPolicy_FieldTerminalPath{selector: ProvisioningPolicy_FieldPathSelectorStatus})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &ProvisioningPolicy_FieldSubPath{selector: ProvisioningPolicy_FieldPathSelectorStatus, subPath: subpath})
			}
		}
	}
	return res
}

func (o *ProvisioningPolicy) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*ProvisioningPolicy))
}

func (o *ProvisioningPolicy) Clone() *ProvisioningPolicy {
	if o == nil {
		return nil
	}
	result := &ProvisioningPolicy{}
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
	result.Metadata = o.Metadata.Clone()
	result.DisplayName = o.DisplayName
	result.Description = o.Description
	result.Spec = o.Spec.Clone()
	result.Status = o.Status.Clone()
	return result
}

func (o *ProvisioningPolicy) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *ProvisioningPolicy) Merge(source *ProvisioningPolicy) {
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
	if source.GetMetadata() != nil {
		if o.Metadata == nil {
			o.Metadata = new(meta.Meta)
		}
		o.Metadata.Merge(source.GetMetadata())
	}
	o.DisplayName = source.GetDisplayName()
	o.Description = source.GetDescription()
	if source.GetSpec() != nil {
		if o.Spec == nil {
			o.Spec = new(ProvisioningPolicy_Spec)
		}
		o.Spec.Merge(source.GetSpec())
	}
	if source.GetStatus() != nil {
		if o.Status == nil {
			o.Status = new(ProvisioningPolicy_Status)
		}
		o.Status.Merge(source.GetStatus())
	}
}

func (o *ProvisioningPolicy) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*ProvisioningPolicy))
}

func (o *ProvisioningPolicy_Spec) GotenObjectExt() {}

func (o *ProvisioningPolicy_Spec) MakeFullFieldMask() *ProvisioningPolicy_Spec_FieldMask {
	return FullProvisioningPolicy_Spec_FieldMask()
}

func (o *ProvisioningPolicy_Spec) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullProvisioningPolicy_Spec_FieldMask()
}

func (o *ProvisioningPolicy_Spec) MakeDiffFieldMask(other *ProvisioningPolicy_Spec) *ProvisioningPolicy_Spec_FieldMask {
	if o == nil && other == nil {
		return &ProvisioningPolicy_Spec_FieldMask{}
	}
	if o == nil || other == nil {
		return FullProvisioningPolicy_Spec_FieldMask()
	}

	res := &ProvisioningPolicy_Spec_FieldMask{}
	if o.GetMode() != other.GetMode() {
		res.Paths = append(res.Paths, &ProvisioningPolicySpec_FieldTerminalPath{selector: ProvisioningPolicySpec_FieldPathSelectorMode})
	}
	if o.GetServiceAccount().String() != other.GetServiceAccount().String() {
		res.Paths = append(res.Paths, &ProvisioningPolicySpec_FieldTerminalPath{selector: ProvisioningPolicySpec_FieldPathSelectorServiceAccount})
	}
	if o.GetDeviceNameFormat() != other.GetDeviceNameFormat() {
		res.Paths = append(res.Paths, &ProvisioningPolicySpec_FieldTerminalPath{selector: ProvisioningPolicySpec_FieldPathSelectorDeviceNameFormat})
	}
	if o.GetDeviceDisplayNameFormat() != other.GetDeviceDisplayNameFormat() {
		res.Paths = append(res.Paths, &ProvisioningPolicySpec_FieldTerminalPath{selector: ProvisioningPolicySpec_FieldPathSelectorDeviceDisplayNameFormat})
	}

	if len(o.GetLabels()) == len(other.GetLabels()) {
		for i, lValue := range o.GetLabels() {
			rValue := other.GetLabels()[i]
			if lValue != rValue {
				res.Paths = append(res.Paths, &ProvisioningPolicySpec_FieldTerminalPath{selector: ProvisioningPolicySpec_FieldPathSelectorLabels})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &ProvisioningPolicySpec_FieldTerminalPath{selector: ProvisioningPolicySpec_FieldPathSelectorLabels})
	}
	{
		subMask := o.GetTemplate().MakeDiffFieldMask(other.GetTemplate())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &ProvisioningPolicySpec_FieldTerminalPath{selector: ProvisioningPolicySpec_FieldPathSelectorTemplate})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &ProvisioningPolicySpec_FieldSubPath{selector: ProvisioningPolicySpec_FieldPathSelectorTemplate, subPath: subpath})
			}
		}
	}

	if len(o.GetIdentityFieldPaths()) == len(other.GetIdentityFieldPaths()) {
		for i, lValue := range o.GetIdentityFieldPaths() {
			rValue := other.GetIdentityFieldPaths()[i]
			if lValue != rValue {
				res.Paths = append(res.Paths, &ProvisioningPolicySpec_FieldTerminalPath{selector: ProvisioningPolicySpec_FieldPathSelectorIdentityFieldPaths})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &ProvisioningPolicySpec_FieldTerminalPath{selector: ProvisioningPolicySpec_FieldPathSelectorIdentityFieldPaths})
	}
	if o.GetRole().String() != other.GetRole().String() {
		res.Paths = append(res.Paths, &ProvisioningPolicySpec_FieldTerminalPath{selector: ProvisioningPolicySpec_FieldPathSelectorRole})
	}

	if len(o.GetScopeParams()) == len(other.GetScopeParams()) {
		for i, lValue := range o.GetScopeParams() {
			rValue := other.GetScopeParams()[i]
			if len(lValue.MakeDiffFieldMask(rValue).Paths) > 0 {
				res.Paths = append(res.Paths, &ProvisioningPolicySpec_FieldTerminalPath{selector: ProvisioningPolicySpec_FieldPathSelectorScopeParams})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &ProvisioningPolicySpec_FieldTerminalPath{selector: ProvisioningPolicySpec_FieldPathSelectorScopeParams})
	}
	if o.GetCondition().String() != other.GetCondition().String() {
		res.Paths = append(res.Paths, &ProvisioningPolicySpec_FieldTerminalPath{selector: ProvisioningPolicySpec_FieldPathSelectorCondition})
	}
	if !proto.Equal(o.GetConditionParams(), other.GetConditionParams()) {
		res.Paths = append(res.Paths, &ProvisioningPolicySpec_FieldTerminalPath{selector: ProvisioningPolicySpec_FieldPathSelectorConditionParams})
	}
	return res
}

func (o *ProvisioningPolicy_Spec) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*ProvisioningPolicy_Spec))
}

func (o *ProvisioningPolicy_Spec) Clone() *ProvisioningPolicy_Spec {
	if o == nil {
		return nil
	}
	result := &ProvisioningPolicy_Spec{}
	result.Mode = o.Mode
	if o.ServiceAccount == nil {
		result.ServiceAccount = nil
	} else if data, err := o.ServiceAccount.ProtoString(); err != nil {
		panic(err)
	} else {
		result.ServiceAccount = &iam_service_account.Reference{}
		if err := result.ServiceAccount.ParseProtoString(data); err != nil {
			panic(err)
		}
	}
	result.DeviceNameFormat = o.DeviceNameFormat
	result.DeviceDisplayNameFormat = o.DeviceDisplayNameFormat
	result.Labels = map[string]string{}
	for key, sourceValue := range o.Labels {
		result.Labels[key] = sourceValue
	}
	result.Template = o.Template.Clone()
	result.IdentityFieldPaths = make([]string, len(o.IdentityFieldPaths))
	for i, sourceValue := range o.IdentityFieldPaths {
		result.IdentityFieldPaths[i] = sourceValue
	}
	if o.Role == nil {
		result.Role = nil
	} else if data, err := o.Role.ProtoString(); err != nil {
		panic(err)
	} else {
		result.Role = &iam_role.Reference{}
		if err := result.Role.ParseProtoString(data); err != nil {
			panic(err)
		}
	}
	result.ScopeParams = make([]*iam_role.ScopeParam, len(o.ScopeParams))
	for i, sourceValue := range o.ScopeParams {
		result.ScopeParams[i] = sourceValue.Clone()
	}
	if o.Condition == nil {
		result.Condition = nil
	} else if data, err := o.Condition.ProtoString(); err != nil {
		panic(err)
	} else {
		result.Condition = &iam_condition.Reference{}
		if err := result.Condition.ParseProtoString(data); err != nil {
			panic(err)
		}
	}
	result.ConditionParams = proto.Clone(o.ConditionParams).(*structpb.Struct)
	return result
}

func (o *ProvisioningPolicy_Spec) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *ProvisioningPolicy_Spec) Merge(source *ProvisioningPolicy_Spec) {
	o.Mode = source.GetMode()
	if source.GetServiceAccount() != nil {
		if data, err := source.GetServiceAccount().ProtoString(); err != nil {
			panic(err)
		} else {
			o.ServiceAccount = &iam_service_account.Reference{}
			if err := o.ServiceAccount.ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	} else {
		o.ServiceAccount = nil
	}
	o.DeviceNameFormat = source.GetDeviceNameFormat()
	o.DeviceDisplayNameFormat = source.GetDeviceDisplayNameFormat()
	if source.GetLabels() != nil {
		if o.Labels == nil {
			o.Labels = make(map[string]string, len(source.GetLabels()))
		}
		for key, sourceValue := range source.GetLabels() {
			o.Labels[key] = sourceValue
		}
	}
	if source.GetTemplate() != nil {
		if o.Template == nil {
			o.Template = new(ProvisioningPolicy_Spec_Template)
		}
		o.Template.Merge(source.GetTemplate())
	}
	for _, sourceValue := range source.GetIdentityFieldPaths() {
		exists := false
		for _, currentValue := range o.IdentityFieldPaths {
			if currentValue == sourceValue {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement string
			newDstElement = sourceValue
			o.IdentityFieldPaths = append(o.IdentityFieldPaths, newDstElement)
		}
	}

	if source.GetRole() != nil {
		if data, err := source.GetRole().ProtoString(); err != nil {
			panic(err)
		} else {
			o.Role = &iam_role.Reference{}
			if err := o.Role.ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	} else {
		o.Role = nil
	}
	for _, sourceValue := range source.GetScopeParams() {
		exists := false
		for _, currentValue := range o.ScopeParams {
			if proto.Equal(sourceValue, currentValue) {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement *iam_role.ScopeParam
			if sourceValue != nil {
				newDstElement = new(iam_role.ScopeParam)
				newDstElement.Merge(sourceValue)
			}
			o.ScopeParams = append(o.ScopeParams, newDstElement)
		}
	}

	if source.GetCondition() != nil {
		if data, err := source.GetCondition().ProtoString(); err != nil {
			panic(err)
		} else {
			o.Condition = &iam_condition.Reference{}
			if err := o.Condition.ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	} else {
		o.Condition = nil
	}
	if source.GetConditionParams() != nil {
		if o.ConditionParams == nil {
			o.ConditionParams = new(structpb.Struct)
		}
		proto.Merge(o.ConditionParams, source.GetConditionParams())
	}
}

func (o *ProvisioningPolicy_Spec) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*ProvisioningPolicy_Spec))
}

func (o *ProvisioningPolicy_Status) GotenObjectExt() {}

func (o *ProvisioningPolicy_Status) MakeFullFieldMask() *ProvisioningPolicy_Status_FieldMask {
	return FullProvisioningPolicy_Status_FieldMask()
}

func (o *ProvisioningPolicy_Status) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullProvisioningPolicy_Status_FieldMask()
}

func (o *ProvisioningPolicy_Status) MakeDiffFieldMask(other *ProvisioningPolicy_Status) *ProvisioningPolicy_Status_FieldMask {
	if o == nil && other == nil {
		return &ProvisioningPolicy_Status_FieldMask{}
	}
	if o == nil || other == nil {
		return FullProvisioningPolicy_Status_FieldMask()
	}

	res := &ProvisioningPolicy_Status_FieldMask{}
	return res
}

func (o *ProvisioningPolicy_Status) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*ProvisioningPolicy_Status))
}

func (o *ProvisioningPolicy_Status) Clone() *ProvisioningPolicy_Status {
	if o == nil {
		return nil
	}
	result := &ProvisioningPolicy_Status{}
	return result
}

func (o *ProvisioningPolicy_Status) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *ProvisioningPolicy_Status) Merge(source *ProvisioningPolicy_Status) {
}

func (o *ProvisioningPolicy_Status) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*ProvisioningPolicy_Status))
}

func (o *ProvisioningPolicy_Spec_Template) GotenObjectExt() {}

func (o *ProvisioningPolicy_Spec_Template) MakeFullFieldMask() *ProvisioningPolicy_Spec_Template_FieldMask {
	return FullProvisioningPolicy_Spec_Template_FieldMask()
}

func (o *ProvisioningPolicy_Spec_Template) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullProvisioningPolicy_Spec_Template_FieldMask()
}

func (o *ProvisioningPolicy_Spec_Template) MakeDiffFieldMask(other *ProvisioningPolicy_Spec_Template) *ProvisioningPolicy_Spec_Template_FieldMask {
	if o == nil && other == nil {
		return &ProvisioningPolicy_Spec_Template_FieldMask{}
	}
	if o == nil || other == nil {
		return FullProvisioningPolicy_Spec_Template_FieldMask()
	}

	res := &ProvisioningPolicy_Spec_Template_FieldMask{}
	{
		subMask := o.GetMetadata().MakeDiffFieldMask(other.GetMetadata())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &ProvisioningPolicySpecTemplate_FieldTerminalPath{selector: ProvisioningPolicySpecTemplate_FieldPathSelectorMetadata})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &ProvisioningPolicySpecTemplate_FieldSubPath{selector: ProvisioningPolicySpecTemplate_FieldPathSelectorMetadata, subPath: subpath})
			}
		}
	}
	{
		subMask := o.GetSpec().MakeDiffFieldMask(other.GetSpec())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &ProvisioningPolicySpecTemplate_FieldTerminalPath{selector: ProvisioningPolicySpecTemplate_FieldPathSelectorSpec})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &ProvisioningPolicySpecTemplate_FieldSubPath{selector: ProvisioningPolicySpecTemplate_FieldPathSelectorSpec, subPath: subpath})
			}
		}
	}
	{
		subMask := o.GetPublicListingSpec().MakeDiffFieldMask(other.GetPublicListingSpec())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &ProvisioningPolicySpecTemplate_FieldTerminalPath{selector: ProvisioningPolicySpecTemplate_FieldPathSelectorPublicListingSpec})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &ProvisioningPolicySpecTemplate_FieldSubPath{selector: ProvisioningPolicySpecTemplate_FieldPathSelectorPublicListingSpec, subPath: subpath})
			}
		}
	}
	return res
}

func (o *ProvisioningPolicy_Spec_Template) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*ProvisioningPolicy_Spec_Template))
}

func (o *ProvisioningPolicy_Spec_Template) Clone() *ProvisioningPolicy_Spec_Template {
	if o == nil {
		return nil
	}
	result := &ProvisioningPolicy_Spec_Template{}
	result.Metadata = o.Metadata.Clone()
	result.Spec = o.Spec.Clone()
	result.PublicListingSpec = o.PublicListingSpec.Clone()
	return result
}

func (o *ProvisioningPolicy_Spec_Template) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *ProvisioningPolicy_Spec_Template) Merge(source *ProvisioningPolicy_Spec_Template) {
	if source.GetMetadata() != nil {
		if o.Metadata == nil {
			o.Metadata = new(meta.Meta)
		}
		o.Metadata.Merge(source.GetMetadata())
	}
	if source.GetSpec() != nil {
		if o.Spec == nil {
			o.Spec = new(device.Device_Spec)
		}
		o.Spec.Merge(source.GetSpec())
	}
	if source.GetPublicListingSpec() != nil {
		if o.PublicListingSpec == nil {
			o.PublicListingSpec = new(device.Device_PublicListingSpec)
		}
		o.PublicListingSpec.Merge(source.GetPublicListingSpec())
	}
}

func (o *ProvisioningPolicy_Spec_Template) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*ProvisioningPolicy_Spec_Template))
}
