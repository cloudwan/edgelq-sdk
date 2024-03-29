// Code generated by protoc-gen-goten-object
// File: edgelq/applications/proto/v1alpha2/project.proto
// DO NOT EDIT!!!

package project

import (
	"fmt"
	"sort"

	"google.golang.org/protobuf/proto"
	googlefieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"

	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
)

// proto imports
import (
	meta "github.com/cloudwan/goten-sdk/types/meta"
	multi_region_policy "github.com/cloudwan/goten-sdk/types/multi_region_policy"
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
	_ = &meta.Meta{}
	_ = &multi_region_policy.MultiRegionPolicy{}
)

func (o *Project) GotenObjectExt() {}

func (o *Project) MakeFullFieldMask() *Project_FieldMask {
	return FullProject_FieldMask()
}

func (o *Project) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullProject_FieldMask()
}

func (o *Project) MakeDiffFieldMask(other *Project) *Project_FieldMask {
	if o == nil && other == nil {
		return &Project_FieldMask{}
	}
	if o == nil || other == nil {
		return FullProject_FieldMask()
	}

	res := &Project_FieldMask{}
	if o.GetName().String() != other.GetName().String() {
		res.Paths = append(res.Paths, &Project_FieldTerminalPath{selector: Project_FieldPathSelectorName})
	}
	{
		subMask := o.GetMetadata().MakeDiffFieldMask(other.GetMetadata())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &Project_FieldTerminalPath{selector: Project_FieldPathSelectorMetadata})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &Project_FieldSubPath{selector: Project_FieldPathSelectorMetadata, subPath: subpath})
			}
		}
	}
	{
		subMask := o.GetMultiRegionPolicy().MakeDiffFieldMask(other.GetMultiRegionPolicy())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &Project_FieldTerminalPath{selector: Project_FieldPathSelectorMultiRegionPolicy})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &Project_FieldSubPath{selector: Project_FieldPathSelectorMultiRegionPolicy, subPath: subpath})
			}
		}
	}
	return res
}

func (o *Project) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*Project))
}

func (o *Project) Clone() *Project {
	if o == nil {
		return nil
	}
	result := &Project{}
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
	result.MultiRegionPolicy = o.MultiRegionPolicy.Clone()
	return result
}

func (o *Project) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *Project) Merge(source *Project) {
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
	if source.GetMultiRegionPolicy() != nil {
		if o.MultiRegionPolicy == nil {
			o.MultiRegionPolicy = new(multi_region_policy.MultiRegionPolicy)
		}
		o.MultiRegionPolicy.Merge(source.GetMultiRegionPolicy())
	}
}

func (o *Project) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*Project))
}
