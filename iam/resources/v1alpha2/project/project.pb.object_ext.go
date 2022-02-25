// Code generated by protoc-gen-goten-object
// File: edgelq/iam/proto/v1alpha2/project.proto
// DO NOT EDIT!!!

package project

import (
	"fmt"
	"sort"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
)

// proto imports
import (
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
	organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/organization"
	policy "github.com/cloudwan/edgelq-sdk/meta/multi_region/proto/policy"
	meta_service "github.com/cloudwan/edgelq-sdk/meta/resources/v1alpha2/service"
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
	_ = &ntt_meta.Meta{}
	_ = &organization.Organization{}
	_ = &policy.Policy{}
	_ = &meta_service.Service{}
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
	if o.GetTitle() != other.GetTitle() {
		res.Paths = append(res.Paths, &Project_FieldTerminalPath{selector: Project_FieldPathSelectorTitle})
	}
	if o.GetParentOrganization().String() != other.GetParentOrganization().String() {
		res.Paths = append(res.Paths, &Project_FieldTerminalPath{selector: Project_FieldPathSelectorParentOrganization})
	}
	if o.GetRootOrganization().String() != other.GetRootOrganization().String() {
		res.Paths = append(res.Paths, &Project_FieldTerminalPath{selector: Project_FieldPathSelectorRootOrganization})
	}

	if len(o.GetAncestryPath()) == len(other.GetAncestryPath()) {
		for i, lValue := range o.GetAncestryPath() {
			rValue := other.GetAncestryPath()[i]
			if lValue.String() != rValue.String() {
				res.Paths = append(res.Paths, &Project_FieldTerminalPath{selector: Project_FieldPathSelectorAncestryPath})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &Project_FieldTerminalPath{selector: Project_FieldPathSelectorAncestryPath})
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

	if len(o.GetEnabledServices()) == len(other.GetEnabledServices()) {
		for i, lValue := range o.GetEnabledServices() {
			rValue := other.GetEnabledServices()[i]
			if lValue.String() != rValue.String() {
				res.Paths = append(res.Paths, &Project_FieldTerminalPath{selector: Project_FieldPathSelectorEnabledServices})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &Project_FieldTerminalPath{selector: Project_FieldPathSelectorEnabledServices})
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
	result.Title = o.Title
	if o.ParentOrganization == nil {
		result.ParentOrganization = nil
	} else if data, err := o.ParentOrganization.ProtoString(); err != nil {
		panic(err)
	} else {
		result.ParentOrganization = &organization.Reference{}
		if err := result.ParentOrganization.ParseProtoString(data); err != nil {
			panic(err)
		}
	}
	if o.RootOrganization == nil {
		result.RootOrganization = nil
	} else if data, err := o.RootOrganization.ProtoString(); err != nil {
		panic(err)
	} else {
		result.RootOrganization = &organization.Reference{}
		if err := result.RootOrganization.ParseProtoString(data); err != nil {
			panic(err)
		}
	}
	result.AncestryPath = make([]*organization.Reference, len(o.AncestryPath))
	for i, sourceValue := range o.AncestryPath {
		if sourceValue == nil {
			result.AncestryPath[i] = nil
		} else if data, err := sourceValue.ProtoString(); err != nil {
			panic(err)
		} else {
			result.AncestryPath[i] = &organization.Reference{}
			if err := result.AncestryPath[i].ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	}
	result.Metadata = o.Metadata.Clone()
	result.MultiRegionPolicy = o.MultiRegionPolicy.Clone()
	result.EnabledServices = make([]*meta_service.Reference, len(o.EnabledServices))
	for i, sourceValue := range o.EnabledServices {
		if sourceValue == nil {
			result.EnabledServices[i] = nil
		} else if data, err := sourceValue.ProtoString(); err != nil {
			panic(err)
		} else {
			result.EnabledServices[i] = &meta_service.Reference{}
			if err := result.EnabledServices[i].ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	}
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
	o.Title = source.GetTitle()
	if source.GetParentOrganization() != nil {
		if data, err := source.GetParentOrganization().ProtoString(); err != nil {
			panic(err)
		} else {
			o.ParentOrganization = &organization.Reference{}
			if err := o.ParentOrganization.ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	} else {
		o.ParentOrganization = nil
	}
	if source.GetRootOrganization() != nil {
		if data, err := source.GetRootOrganization().ProtoString(); err != nil {
			panic(err)
		} else {
			o.RootOrganization = &organization.Reference{}
			if err := o.RootOrganization.ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	} else {
		o.RootOrganization = nil
	}
	for _, sourceValue := range source.GetAncestryPath() {
		exists := false
		for _, currentValue := range o.AncestryPath {
			leftProtoStr, _ := currentValue.ProtoString()
			rightProtoStr, _ := sourceValue.ProtoString()
			if leftProtoStr == rightProtoStr {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement *organization.Reference
			if sourceValue != nil {
				if data, err := sourceValue.ProtoString(); err != nil {
					panic(err)
				} else {
					newDstElement = &organization.Reference{}
					if err := newDstElement.ParseProtoString(data); err != nil {
						panic(err)
					}
				}
			}
			o.AncestryPath = append(o.AncestryPath, newDstElement)
		}
	}

	if source.GetMetadata() != nil {
		if o.Metadata == nil {
			o.Metadata = new(ntt_meta.Meta)
		}
		o.Metadata.Merge(source.GetMetadata())
	}
	if source.GetMultiRegionPolicy() != nil {
		if o.MultiRegionPolicy == nil {
			o.MultiRegionPolicy = new(policy.Policy)
		}
		o.MultiRegionPolicy.Merge(source.GetMultiRegionPolicy())
	}
	for _, sourceValue := range source.GetEnabledServices() {
		exists := false
		for _, currentValue := range o.EnabledServices {
			leftProtoStr, _ := currentValue.ProtoString()
			rightProtoStr, _ := sourceValue.ProtoString()
			if leftProtoStr == rightProtoStr {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement *meta_service.Reference
			if sourceValue != nil {
				if data, err := sourceValue.ProtoString(); err != nil {
					panic(err)
				} else {
					newDstElement = &meta_service.Reference{}
					if err := newDstElement.ParseProtoString(data); err != nil {
						panic(err)
					}
				}
			}
			o.EnabledServices = append(o.EnabledServices, newDstElement)
		}
	}

}

func (o *Project) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*Project))
}
