// Code generated by protoc-gen-goten-object
// File: edgelq/applications/proto/v1/distribution.proto
// DO NOT EDIT!!!

package distribution

import (
	"fmt"
	"sort"

	"google.golang.org/protobuf/proto"
	googlefieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"

	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
)

// proto imports
import (
	common "github.com/cloudwan/edgelq-sdk/applications/resources/v1/common"
	project "github.com/cloudwan/edgelq-sdk/applications/resources/v1/project"
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
	_ = &common.PodSpec{}
	_ = &project.Project{}
	_ = &meta.Meta{}
)

func (o *Distribution) GotenObjectExt() {}

func (o *Distribution) MakeFullFieldMask() *Distribution_FieldMask {
	return FullDistribution_FieldMask()
}

func (o *Distribution) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullDistribution_FieldMask()
}

func (o *Distribution) MakeDiffFieldMask(other *Distribution) *Distribution_FieldMask {
	if o == nil && other == nil {
		return &Distribution_FieldMask{}
	}
	if o == nil || other == nil {
		return FullDistribution_FieldMask()
	}

	res := &Distribution_FieldMask{}
	if o.GetName().String() != other.GetName().String() {
		res.Paths = append(res.Paths, &Distribution_FieldTerminalPath{selector: Distribution_FieldPathSelectorName})
	}
	{
		subMask := o.GetMetadata().MakeDiffFieldMask(other.GetMetadata())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &Distribution_FieldTerminalPath{selector: Distribution_FieldPathSelectorMetadata})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &Distribution_FieldSubPath{selector: Distribution_FieldPathSelectorMetadata, subPath: subpath})
			}
		}
	}
	if o.GetDisplayName() != other.GetDisplayName() {
		res.Paths = append(res.Paths, &Distribution_FieldTerminalPath{selector: Distribution_FieldPathSelectorDisplayName})
	}
	if o.GetDescription() != other.GetDescription() {
		res.Paths = append(res.Paths, &Distribution_FieldTerminalPath{selector: Distribution_FieldPathSelectorDescription})
	}
	{
		subMask := o.GetSpec().MakeDiffFieldMask(other.GetSpec())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &Distribution_FieldTerminalPath{selector: Distribution_FieldPathSelectorSpec})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &Distribution_FieldSubPath{selector: Distribution_FieldPathSelectorSpec, subPath: subpath})
			}
		}
	}
	{
		subMask := o.GetStatus().MakeDiffFieldMask(other.GetStatus())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &Distribution_FieldTerminalPath{selector: Distribution_FieldPathSelectorStatus})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &Distribution_FieldSubPath{selector: Distribution_FieldPathSelectorStatus, subPath: subpath})
			}
		}
	}
	return res
}

func (o *Distribution) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*Distribution))
}

func (o *Distribution) Clone() *Distribution {
	if o == nil {
		return nil
	}
	result := &Distribution{}
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

func (o *Distribution) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *Distribution) Merge(source *Distribution) {
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
			o.Spec = new(Distribution_Spec)
		}
		o.Spec.Merge(source.GetSpec())
	}
	if source.GetStatus() != nil {
		if o.Status == nil {
			o.Status = new(Distribution_Status)
		}
		o.Status.Merge(source.GetStatus())
	}
}

func (o *Distribution) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*Distribution))
}

func (o *Distribution_Spec) GotenObjectExt() {}

func (o *Distribution_Spec) MakeFullFieldMask() *Distribution_Spec_FieldMask {
	return FullDistribution_Spec_FieldMask()
}

func (o *Distribution_Spec) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullDistribution_Spec_FieldMask()
}

func (o *Distribution_Spec) MakeDiffFieldMask(other *Distribution_Spec) *Distribution_Spec_FieldMask {
	if o == nil && other == nil {
		return &Distribution_Spec_FieldMask{}
	}
	if o == nil || other == nil {
		return FullDistribution_Spec_FieldMask()
	}

	res := &Distribution_Spec_FieldMask{}
	{
		subMask := o.GetSelector().MakeDiffFieldMask(other.GetSelector())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &DistributionSpec_FieldTerminalPath{selector: DistributionSpec_FieldPathSelectorSelector})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &DistributionSpec_FieldSubPath{selector: DistributionSpec_FieldPathSelectorSelector, subPath: subpath})
			}
		}
	}
	{
		subMask := o.GetTemplate().MakeDiffFieldMask(other.GetTemplate())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &DistributionSpec_FieldTerminalPath{selector: DistributionSpec_FieldPathSelectorTemplate})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &DistributionSpec_FieldSubPath{selector: DistributionSpec_FieldPathSelectorTemplate, subPath: subpath})
			}
		}
	}
	if o.GetPodDisplayNameFormat() != other.GetPodDisplayNameFormat() {
		res.Paths = append(res.Paths, &DistributionSpec_FieldTerminalPath{selector: DistributionSpec_FieldPathSelectorPodDisplayNameFormat})
	}
	return res
}

func (o *Distribution_Spec) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*Distribution_Spec))
}

func (o *Distribution_Spec) Clone() *Distribution_Spec {
	if o == nil {
		return nil
	}
	result := &Distribution_Spec{}
	result.Selector = o.Selector.Clone()
	result.Template = o.Template.Clone()
	result.PodDisplayNameFormat = o.PodDisplayNameFormat
	return result
}

func (o *Distribution_Spec) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *Distribution_Spec) Merge(source *Distribution_Spec) {
	if source.GetSelector() != nil {
		if o.Selector == nil {
			o.Selector = new(LabelSelector)
		}
		o.Selector.Merge(source.GetSelector())
	}
	if source.GetTemplate() != nil {
		if o.Template == nil {
			o.Template = new(Distribution_Spec_Template)
		}
		o.Template.Merge(source.GetTemplate())
	}
	o.PodDisplayNameFormat = source.GetPodDisplayNameFormat()
}

func (o *Distribution_Spec) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*Distribution_Spec))
}

func (o *Distribution_Status) GotenObjectExt() {}

func (o *Distribution_Status) MakeFullFieldMask() *Distribution_Status_FieldMask {
	return FullDistribution_Status_FieldMask()
}

func (o *Distribution_Status) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullDistribution_Status_FieldMask()
}

func (o *Distribution_Status) MakeDiffFieldMask(other *Distribution_Status) *Distribution_Status_FieldMask {
	if o == nil && other == nil {
		return &Distribution_Status_FieldMask{}
	}
	if o == nil || other == nil {
		return FullDistribution_Status_FieldMask()
	}

	res := &Distribution_Status_FieldMask{}
	return res
}

func (o *Distribution_Status) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*Distribution_Status))
}

func (o *Distribution_Status) Clone() *Distribution_Status {
	if o == nil {
		return nil
	}
	result := &Distribution_Status{}
	return result
}

func (o *Distribution_Status) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *Distribution_Status) Merge(source *Distribution_Status) {
}

func (o *Distribution_Status) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*Distribution_Status))
}

func (o *Distribution_Spec_Template) GotenObjectExt() {}

func (o *Distribution_Spec_Template) MakeFullFieldMask() *Distribution_Spec_Template_FieldMask {
	return FullDistribution_Spec_Template_FieldMask()
}

func (o *Distribution_Spec_Template) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullDistribution_Spec_Template_FieldMask()
}

func (o *Distribution_Spec_Template) MakeDiffFieldMask(other *Distribution_Spec_Template) *Distribution_Spec_Template_FieldMask {
	if o == nil && other == nil {
		return &Distribution_Spec_Template_FieldMask{}
	}
	if o == nil || other == nil {
		return FullDistribution_Spec_Template_FieldMask()
	}

	res := &Distribution_Spec_Template_FieldMask{}
	{
		subMask := o.GetMetadata().MakeDiffFieldMask(other.GetMetadata())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &DistributionSpecTemplate_FieldTerminalPath{selector: DistributionSpecTemplate_FieldPathSelectorMetadata})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &DistributionSpecTemplate_FieldSubPath{selector: DistributionSpecTemplate_FieldPathSelectorMetadata, subPath: subpath})
			}
		}
	}
	{
		subMask := o.GetSpec().MakeDiffFieldMask(other.GetSpec())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &DistributionSpecTemplate_FieldTerminalPath{selector: DistributionSpecTemplate_FieldPathSelectorSpec})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &DistributionSpecTemplate_FieldSubPath{selector: DistributionSpecTemplate_FieldPathSelectorSpec, subPath: subpath})
			}
		}
	}
	return res
}

func (o *Distribution_Spec_Template) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*Distribution_Spec_Template))
}

func (o *Distribution_Spec_Template) Clone() *Distribution_Spec_Template {
	if o == nil {
		return nil
	}
	result := &Distribution_Spec_Template{}
	result.Metadata = o.Metadata.Clone()
	result.Spec = o.Spec.Clone()
	return result
}

func (o *Distribution_Spec_Template) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *Distribution_Spec_Template) Merge(source *Distribution_Spec_Template) {
	if source.GetMetadata() != nil {
		if o.Metadata == nil {
			o.Metadata = new(meta.Meta)
		}
		o.Metadata.Merge(source.GetMetadata())
	}
	if source.GetSpec() != nil {
		if o.Spec == nil {
			o.Spec = new(common.PodSpec)
		}
		o.Spec.Merge(source.GetSpec())
	}
}

func (o *Distribution_Spec_Template) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*Distribution_Spec_Template))
}

func (o *LabelSelector) GotenObjectExt() {}

func (o *LabelSelector) MakeFullFieldMask() *LabelSelector_FieldMask {
	return FullLabelSelector_FieldMask()
}

func (o *LabelSelector) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullLabelSelector_FieldMask()
}

func (o *LabelSelector) MakeDiffFieldMask(other *LabelSelector) *LabelSelector_FieldMask {
	if o == nil && other == nil {
		return &LabelSelector_FieldMask{}
	}
	if o == nil || other == nil {
		return FullLabelSelector_FieldMask()
	}

	res := &LabelSelector_FieldMask{}

	if len(o.GetMatchLabels()) == len(other.GetMatchLabels()) {
		for i, lValue := range o.GetMatchLabels() {
			rValue := other.GetMatchLabels()[i]
			if lValue != rValue {
				res.Paths = append(res.Paths, &LabelSelector_FieldTerminalPath{selector: LabelSelector_FieldPathSelectorMatchLabels})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &LabelSelector_FieldTerminalPath{selector: LabelSelector_FieldPathSelectorMatchLabels})
	}

	if len(o.GetMatchExpressions()) == len(other.GetMatchExpressions()) {
		for i, lValue := range o.GetMatchExpressions() {
			rValue := other.GetMatchExpressions()[i]
			if len(lValue.MakeDiffFieldMask(rValue).Paths) > 0 {
				res.Paths = append(res.Paths, &LabelSelector_FieldTerminalPath{selector: LabelSelector_FieldPathSelectorMatchExpressions})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &LabelSelector_FieldTerminalPath{selector: LabelSelector_FieldPathSelectorMatchExpressions})
	}
	return res
}

func (o *LabelSelector) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*LabelSelector))
}

func (o *LabelSelector) Clone() *LabelSelector {
	if o == nil {
		return nil
	}
	result := &LabelSelector{}
	result.MatchLabels = map[string]string{}
	for key, sourceValue := range o.MatchLabels {
		result.MatchLabels[key] = sourceValue
	}
	result.MatchExpressions = make([]*LabelSelectorRequirement, len(o.MatchExpressions))
	for i, sourceValue := range o.MatchExpressions {
		result.MatchExpressions[i] = sourceValue.Clone()
	}
	return result
}

func (o *LabelSelector) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *LabelSelector) Merge(source *LabelSelector) {
	if source.GetMatchLabels() != nil {
		if o.MatchLabels == nil {
			o.MatchLabels = make(map[string]string, len(source.GetMatchLabels()))
		}
		for key, sourceValue := range source.GetMatchLabels() {
			o.MatchLabels[key] = sourceValue
		}
	}
	for _, sourceValue := range source.GetMatchExpressions() {
		exists := false
		for _, currentValue := range o.MatchExpressions {
			if proto.Equal(sourceValue, currentValue) {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement *LabelSelectorRequirement
			if sourceValue != nil {
				newDstElement = new(LabelSelectorRequirement)
				newDstElement.Merge(sourceValue)
			}
			o.MatchExpressions = append(o.MatchExpressions, newDstElement)
		}
	}

}

func (o *LabelSelector) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*LabelSelector))
}

func (o *LabelSelectorRequirement) GotenObjectExt() {}

func (o *LabelSelectorRequirement) MakeFullFieldMask() *LabelSelectorRequirement_FieldMask {
	return FullLabelSelectorRequirement_FieldMask()
}

func (o *LabelSelectorRequirement) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullLabelSelectorRequirement_FieldMask()
}

func (o *LabelSelectorRequirement) MakeDiffFieldMask(other *LabelSelectorRequirement) *LabelSelectorRequirement_FieldMask {
	if o == nil && other == nil {
		return &LabelSelectorRequirement_FieldMask{}
	}
	if o == nil || other == nil {
		return FullLabelSelectorRequirement_FieldMask()
	}

	res := &LabelSelectorRequirement_FieldMask{}
	if o.GetKey() != other.GetKey() {
		res.Paths = append(res.Paths, &LabelSelectorRequirement_FieldTerminalPath{selector: LabelSelectorRequirement_FieldPathSelectorKey})
	}
	if o.GetOperator() != other.GetOperator() {
		res.Paths = append(res.Paths, &LabelSelectorRequirement_FieldTerminalPath{selector: LabelSelectorRequirement_FieldPathSelectorOperator})
	}

	if len(o.GetValues()) == len(other.GetValues()) {
		for i, lValue := range o.GetValues() {
			rValue := other.GetValues()[i]
			if lValue != rValue {
				res.Paths = append(res.Paths, &LabelSelectorRequirement_FieldTerminalPath{selector: LabelSelectorRequirement_FieldPathSelectorValues})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &LabelSelectorRequirement_FieldTerminalPath{selector: LabelSelectorRequirement_FieldPathSelectorValues})
	}
	return res
}

func (o *LabelSelectorRequirement) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*LabelSelectorRequirement))
}

func (o *LabelSelectorRequirement) Clone() *LabelSelectorRequirement {
	if o == nil {
		return nil
	}
	result := &LabelSelectorRequirement{}
	result.Key = o.Key
	result.Operator = o.Operator
	result.Values = make([]string, len(o.Values))
	for i, sourceValue := range o.Values {
		result.Values[i] = sourceValue
	}
	return result
}

func (o *LabelSelectorRequirement) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *LabelSelectorRequirement) Merge(source *LabelSelectorRequirement) {
	o.Key = source.GetKey()
	o.Operator = source.GetOperator()
	for _, sourceValue := range source.GetValues() {
		exists := false
		for _, currentValue := range o.Values {
			if currentValue == sourceValue {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement string
			newDstElement = sourceValue
			o.Values = append(o.Values, newDstElement)
		}
	}

}

func (o *LabelSelectorRequirement) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*LabelSelectorRequirement))
}
