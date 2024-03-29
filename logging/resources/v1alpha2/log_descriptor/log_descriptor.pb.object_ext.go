// Code generated by protoc-gen-goten-object
// File: edgelq/logging/proto/v1alpha2/log_descriptor.proto
// DO NOT EDIT!!!

package log_descriptor

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
	common "github.com/cloudwan/edgelq-sdk/logging/resources/v1alpha2/common"
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
	_ = &common.LabelDescriptor{}
	_ = &meta.Meta{}
)

func (o *LogDescriptor) GotenObjectExt() {}

func (o *LogDescriptor) MakeFullFieldMask() *LogDescriptor_FieldMask {
	return FullLogDescriptor_FieldMask()
}

func (o *LogDescriptor) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullLogDescriptor_FieldMask()
}

func (o *LogDescriptor) MakeDiffFieldMask(other *LogDescriptor) *LogDescriptor_FieldMask {
	if o == nil && other == nil {
		return &LogDescriptor_FieldMask{}
	}
	if o == nil || other == nil {
		return FullLogDescriptor_FieldMask()
	}

	res := &LogDescriptor_FieldMask{}
	if o.GetName().String() != other.GetName().String() {
		res.Paths = append(res.Paths, &LogDescriptor_FieldTerminalPath{selector: LogDescriptor_FieldPathSelectorName})
	}
	if o.GetDisplayName() != other.GetDisplayName() {
		res.Paths = append(res.Paths, &LogDescriptor_FieldTerminalPath{selector: LogDescriptor_FieldPathSelectorDisplayName})
	}
	if o.GetDescription() != other.GetDescription() {
		res.Paths = append(res.Paths, &LogDescriptor_FieldTerminalPath{selector: LogDescriptor_FieldPathSelectorDescription})
	}

	if len(o.GetLabels()) == len(other.GetLabels()) {
		for i, lValue := range o.GetLabels() {
			rValue := other.GetLabels()[i]
			if len(lValue.MakeDiffFieldMask(rValue).Paths) > 0 {
				res.Paths = append(res.Paths, &LogDescriptor_FieldTerminalPath{selector: LogDescriptor_FieldPathSelectorLabels})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &LogDescriptor_FieldTerminalPath{selector: LogDescriptor_FieldPathSelectorLabels})
	}

	if len(o.GetPromotedLabelKeySets()) == len(other.GetPromotedLabelKeySets()) {
		for i, lValue := range o.GetPromotedLabelKeySets() {
			rValue := other.GetPromotedLabelKeySets()[i]
			if len(lValue.MakeDiffFieldMask(rValue).Paths) > 0 {
				res.Paths = append(res.Paths, &LogDescriptor_FieldTerminalPath{selector: LogDescriptor_FieldPathSelectorPromotedLabelKeySets})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &LogDescriptor_FieldTerminalPath{selector: LogDescriptor_FieldPathSelectorPromotedLabelKeySets})
	}
	{
		subMask := o.GetMetadata().MakeDiffFieldMask(other.GetMetadata())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &LogDescriptor_FieldTerminalPath{selector: LogDescriptor_FieldPathSelectorMetadata})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &LogDescriptor_FieldSubPath{selector: LogDescriptor_FieldPathSelectorMetadata, subPath: subpath})
			}
		}
	}
	return res
}

func (o *LogDescriptor) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*LogDescriptor))
}

func (o *LogDescriptor) Clone() *LogDescriptor {
	if o == nil {
		return nil
	}
	result := &LogDescriptor{}
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
	result.DisplayName = o.DisplayName
	result.Description = o.Description
	result.Labels = make([]*common.LabelDescriptor, len(o.Labels))
	for i, sourceValue := range o.Labels {
		result.Labels[i] = sourceValue.Clone()
	}
	result.PromotedLabelKeySets = make([]*common.LabelKeySet, len(o.PromotedLabelKeySets))
	for i, sourceValue := range o.PromotedLabelKeySets {
		result.PromotedLabelKeySets[i] = sourceValue.Clone()
	}
	result.Metadata = o.Metadata.Clone()
	return result
}

func (o *LogDescriptor) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *LogDescriptor) Merge(source *LogDescriptor) {
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
	o.DisplayName = source.GetDisplayName()
	o.Description = source.GetDescription()
	for _, sourceValue := range source.GetLabels() {
		exists := false
		for _, currentValue := range o.Labels {
			if proto.Equal(sourceValue, currentValue) {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement *common.LabelDescriptor
			if sourceValue != nil {
				newDstElement = new(common.LabelDescriptor)
				newDstElement.Merge(sourceValue)
			}
			o.Labels = append(o.Labels, newDstElement)
		}
	}

	for _, sourceValue := range source.GetPromotedLabelKeySets() {
		exists := false
		for _, currentValue := range o.PromotedLabelKeySets {
			if proto.Equal(sourceValue, currentValue) {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement *common.LabelKeySet
			if sourceValue != nil {
				newDstElement = new(common.LabelKeySet)
				newDstElement.Merge(sourceValue)
			}
			o.PromotedLabelKeySets = append(o.PromotedLabelKeySets, newDstElement)
		}
	}

	if source.GetMetadata() != nil {
		if o.Metadata == nil {
			o.Metadata = new(meta.Meta)
		}
		o.Metadata.Merge(source.GetMetadata())
	}
}

func (o *LogDescriptor) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*LogDescriptor))
}
