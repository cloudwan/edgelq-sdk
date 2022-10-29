// Code generated by protoc-gen-goten-object
// File: edgelq/iam/proto/v1alpha2/organization.proto
// DO NOT EDIT!!!

package organization

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
	multi_region_policy "github.com/cloudwan/edgelq-sdk/common/types/multi_region_policy"
)

// ensure the imports are used
var (
	_ = new(fmt.Stringer)
	_ = new(sort.Interface)

	_ = new(proto.Message)
	_ = fieldmaskpb.FieldMask{}

	_ = new(gotenobject.FieldPath)
)

// make sure we're using proto imports
var (
	_ = &ntt_meta.Meta{}
	_ = &multi_region_policy.MultiRegionPolicy{}
)

func (o *Organization) GotenObjectExt() {}

func (o *Organization) MakeFullFieldMask() *Organization_FieldMask {
	return FullOrganization_FieldMask()
}

func (o *Organization) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullOrganization_FieldMask()
}

func (o *Organization) MakeDiffFieldMask(other *Organization) *Organization_FieldMask {
	if o == nil && other == nil {
		return &Organization_FieldMask{}
	}
	if o == nil || other == nil {
		return FullOrganization_FieldMask()
	}

	res := &Organization_FieldMask{}
	if o.GetName().String() != other.GetName().String() {
		res.Paths = append(res.Paths, &Organization_FieldTerminalPath{selector: Organization_FieldPathSelectorName})
	}
	if o.GetTitle() != other.GetTitle() {
		res.Paths = append(res.Paths, &Organization_FieldTerminalPath{selector: Organization_FieldPathSelectorTitle})
	}
	if o.GetParentOrganization().String() != other.GetParentOrganization().String() {
		res.Paths = append(res.Paths, &Organization_FieldTerminalPath{selector: Organization_FieldPathSelectorParentOrganization})
	}
	if o.GetRootOrganization().String() != other.GetRootOrganization().String() {
		res.Paths = append(res.Paths, &Organization_FieldTerminalPath{selector: Organization_FieldPathSelectorRootOrganization})
	}

	if len(o.GetAncestryPath()) == len(other.GetAncestryPath()) {
		for i, lValue := range o.GetAncestryPath() {
			rValue := other.GetAncestryPath()[i]
			if lValue.String() != rValue.String() {
				res.Paths = append(res.Paths, &Organization_FieldTerminalPath{selector: Organization_FieldPathSelectorAncestryPath})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &Organization_FieldTerminalPath{selector: Organization_FieldPathSelectorAncestryPath})
	}
	{
		subMask := o.GetMetadata().MakeDiffFieldMask(other.GetMetadata())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &Organization_FieldTerminalPath{selector: Organization_FieldPathSelectorMetadata})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &Organization_FieldSubPath{selector: Organization_FieldPathSelectorMetadata, subPath: subpath})
			}
		}
	}
	{
		subMask := o.GetMultiRegionPolicy().MakeDiffFieldMask(other.GetMultiRegionPolicy())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &Organization_FieldTerminalPath{selector: Organization_FieldPathSelectorMultiRegionPolicy})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &Organization_FieldSubPath{selector: Organization_FieldPathSelectorMultiRegionPolicy, subPath: subpath})
			}
		}
	}
	return res
}

func (o *Organization) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*Organization))
}

func (o *Organization) Clone() *Organization {
	if o == nil {
		return nil
	}
	result := &Organization{}
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
		result.ParentOrganization = &Reference{}
		if err := result.ParentOrganization.ParseProtoString(data); err != nil {
			panic(err)
		}
	}
	if o.RootOrganization == nil {
		result.RootOrganization = nil
	} else if data, err := o.RootOrganization.ProtoString(); err != nil {
		panic(err)
	} else {
		result.RootOrganization = &Reference{}
		if err := result.RootOrganization.ParseProtoString(data); err != nil {
			panic(err)
		}
	}
	result.AncestryPath = make([]*Reference, len(o.AncestryPath))
	for i, sourceValue := range o.AncestryPath {
		if sourceValue == nil {
			result.AncestryPath[i] = nil
		} else if data, err := sourceValue.ProtoString(); err != nil {
			panic(err)
		} else {
			result.AncestryPath[i] = &Reference{}
			if err := result.AncestryPath[i].ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	}
	result.Metadata = o.Metadata.Clone()
	result.MultiRegionPolicy = o.MultiRegionPolicy.Clone()
	return result
}

func (o *Organization) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *Organization) Merge(source *Organization) {
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
			o.ParentOrganization = &Reference{}
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
			o.RootOrganization = &Reference{}
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
			var newDstElement *Reference
			if sourceValue != nil {
				if data, err := sourceValue.ProtoString(); err != nil {
					panic(err)
				} else {
					newDstElement = &Reference{}
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
			o.MultiRegionPolicy = new(multi_region_policy.MultiRegionPolicy)
		}
		o.MultiRegionPolicy.Merge(source.GetMultiRegionPolicy())
	}
}

func (o *Organization) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*Organization))
}
