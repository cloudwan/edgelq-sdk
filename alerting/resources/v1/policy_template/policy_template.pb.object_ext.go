// Code generated by protoc-gen-goten-object
// File: edgelq/alerting/proto/v1/policy_template.proto
// DO NOT EDIT!!!

package policy_template

import (
	"fmt"
	"sort"

	"google.golang.org/protobuf/proto"
	googlefieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"

	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
)

// proto imports
import (
	rcommon "github.com/cloudwan/edgelq-sdk/alerting/resources/v1/common"
	document "github.com/cloudwan/edgelq-sdk/alerting/resources/v1/document"
	iam_project "github.com/cloudwan/edgelq-sdk/iam/resources/v1/project"
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
	_ = &document.Document{}
	_ = &rcommon.LogCndSpec{}
	_ = &iam_project.Project{}
	_ = &meta.Meta{}
)

func (o *PolicyTemplate) GotenObjectExt() {}

func (o *PolicyTemplate) MakeFullFieldMask() *PolicyTemplate_FieldMask {
	return FullPolicyTemplate_FieldMask()
}

func (o *PolicyTemplate) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullPolicyTemplate_FieldMask()
}

func (o *PolicyTemplate) MakeDiffFieldMask(other *PolicyTemplate) *PolicyTemplate_FieldMask {
	if o == nil && other == nil {
		return &PolicyTemplate_FieldMask{}
	}
	if o == nil || other == nil {
		return FullPolicyTemplate_FieldMask()
	}

	res := &PolicyTemplate_FieldMask{}
	if o.GetName().String() != other.GetName().String() {
		res.Paths = append(res.Paths, &PolicyTemplate_FieldTerminalPath{selector: PolicyTemplate_FieldPathSelectorName})
	}
	{
		subMask := o.GetMetadata().MakeDiffFieldMask(other.GetMetadata())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &PolicyTemplate_FieldTerminalPath{selector: PolicyTemplate_FieldPathSelectorMetadata})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &PolicyTemplate_FieldSubPath{selector: PolicyTemplate_FieldPathSelectorMetadata, subPath: subpath})
			}
		}
	}
	if o.GetDisplayName() != other.GetDisplayName() {
		res.Paths = append(res.Paths, &PolicyTemplate_FieldTerminalPath{selector: PolicyTemplate_FieldPathSelectorDisplayName})
	}
	if o.GetDescription() != other.GetDescription() {
		res.Paths = append(res.Paths, &PolicyTemplate_FieldTerminalPath{selector: PolicyTemplate_FieldPathSelectorDescription})
	}

	if len(o.GetSupportingDocs()) == len(other.GetSupportingDocs()) {
		for i, lValue := range o.GetSupportingDocs() {
			rValue := other.GetSupportingDocs()[i]
			if lValue.String() != rValue.String() {
				res.Paths = append(res.Paths, &PolicyTemplate_FieldTerminalPath{selector: PolicyTemplate_FieldPathSelectorSupportingDocs})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &PolicyTemplate_FieldTerminalPath{selector: PolicyTemplate_FieldPathSelectorSupportingDocs})
	}
	{
		subMask := o.GetSpecTemplate().MakeDiffFieldMask(other.GetSpecTemplate())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &PolicyTemplate_FieldTerminalPath{selector: PolicyTemplate_FieldPathSelectorSpecTemplate})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &PolicyTemplate_FieldSubPath{selector: PolicyTemplate_FieldPathSelectorSpecTemplate, subPath: subpath})
			}
		}
	}
	return res
}

func (o *PolicyTemplate) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*PolicyTemplate))
}

func (o *PolicyTemplate) Clone() *PolicyTemplate {
	if o == nil {
		return nil
	}
	result := &PolicyTemplate{}
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
	result.SupportingDocs = make([]*document.Reference, len(o.SupportingDocs))
	for i, sourceValue := range o.SupportingDocs {
		if sourceValue == nil {
			result.SupportingDocs[i] = nil
		} else if data, err := sourceValue.ProtoString(); err != nil {
			panic(err)
		} else {
			result.SupportingDocs[i] = &document.Reference{}
			if err := result.SupportingDocs[i].ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	}
	result.SpecTemplate = o.SpecTemplate.Clone()
	return result
}

func (o *PolicyTemplate) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *PolicyTemplate) Merge(source *PolicyTemplate) {
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
	for _, sourceValue := range source.GetSupportingDocs() {
		exists := false
		for _, currentValue := range o.SupportingDocs {
			leftProtoStr, _ := currentValue.ProtoString()
			rightProtoStr, _ := sourceValue.ProtoString()
			if leftProtoStr == rightProtoStr {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement *document.Reference
			if sourceValue != nil {
				if data, err := sourceValue.ProtoString(); err != nil {
					panic(err)
				} else {
					newDstElement = &document.Reference{}
					if err := newDstElement.ParseProtoString(data); err != nil {
						panic(err)
					}
				}
			}
			o.SupportingDocs = append(o.SupportingDocs, newDstElement)
		}
	}

	if source.GetSpecTemplate() != nil {
		if o.SpecTemplate == nil {
			o.SpecTemplate = new(rcommon.PolicySpec)
		}
		o.SpecTemplate.Merge(source.GetSpecTemplate())
	}
}

func (o *PolicyTemplate) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*PolicyTemplate))
}
