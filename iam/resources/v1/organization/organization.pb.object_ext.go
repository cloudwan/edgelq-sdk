// Code generated by protoc-gen-goten-object
// File: edgelq/iam/proto/v1/organization.proto
// DO NOT EDIT!!!

package organization

import (
	"fmt"
	"sort"

	"google.golang.org/protobuf/proto"
	googlefieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"

	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
)

// proto imports
import (
	iam_common "github.com/cloudwan/edgelq-sdk/iam/resources/v1/common"
	meta_service "github.com/cloudwan/goten-sdk/meta-service/resources/v1/service"
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
	_ = &iam_common.PCR{}
	_ = &meta_service.Service{}
	_ = &meta.Meta{}
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
	if o.GetTitle() != other.GetTitle() {
		res.Paths = append(res.Paths, &Organization_FieldTerminalPath{selector: Organization_FieldPathSelectorTitle})
	}
	if o.GetDescription() != other.GetDescription() {
		res.Paths = append(res.Paths, &Organization_FieldTerminalPath{selector: Organization_FieldPathSelectorDescription})
	}
	if o.GetParentOrganization().String() != other.GetParentOrganization().String() {
		res.Paths = append(res.Paths, &Organization_FieldTerminalPath{selector: Organization_FieldPathSelectorParentOrganization})
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

	if len(o.GetAllowedServices()) == len(other.GetAllowedServices()) {
		for i, lValue := range o.GetAllowedServices() {
			rValue := other.GetAllowedServices()[i]
			if lValue.String() != rValue.String() {
				res.Paths = append(res.Paths, &Organization_FieldTerminalPath{selector: Organization_FieldPathSelectorAllowedServices})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &Organization_FieldTerminalPath{selector: Organization_FieldPathSelectorAllowedServices})
	}
	if o.GetBusinessTier() != other.GetBusinessTier() {
		res.Paths = append(res.Paths, &Organization_FieldTerminalPath{selector: Organization_FieldPathSelectorBusinessTier})
	}

	if len(o.GetServiceTiers()) == len(other.GetServiceTiers()) {
		for i, lValue := range o.GetServiceTiers() {
			rValue := other.GetServiceTiers()[i]
			if len(lValue.MakeDiffFieldMask(rValue).Paths) > 0 {
				res.Paths = append(res.Paths, &Organization_FieldTerminalPath{selector: Organization_FieldPathSelectorServiceTiers})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &Organization_FieldTerminalPath{selector: Organization_FieldPathSelectorServiceTiers})
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

	if len(o.GetServiceErrors()) == len(other.GetServiceErrors()) {
		for i, lValue := range o.GetServiceErrors() {
			rValue := other.GetServiceErrors()[i]
			if len(lValue.MakeDiffFieldMask(rValue).Paths) > 0 {
				res.Paths = append(res.Paths, &Organization_FieldTerminalPath{selector: Organization_FieldPathSelectorServiceErrors})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &Organization_FieldTerminalPath{selector: Organization_FieldPathSelectorServiceErrors})
	}
	if o.GetServicesGeneration() != other.GetServicesGeneration() {
		res.Paths = append(res.Paths, &Organization_FieldTerminalPath{selector: Organization_FieldPathSelectorServicesGeneration})
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
	result.Metadata = o.Metadata.Clone()
	result.Title = o.Title
	result.Description = o.Description
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
	result.MultiRegionPolicy = o.MultiRegionPolicy.Clone()
	result.AllowedServices = make([]*meta_service.Reference, len(o.AllowedServices))
	for i, sourceValue := range o.AllowedServices {
		if sourceValue == nil {
			result.AllowedServices[i] = nil
		} else if data, err := sourceValue.ProtoString(); err != nil {
			panic(err)
		} else {
			result.AllowedServices[i] = &meta_service.Reference{}
			if err := result.AllowedServices[i].ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	}
	result.BusinessTier = o.BusinessTier
	result.ServiceTiers = make([]*iam_common.ServiceBusinessTier, len(o.ServiceTiers))
	for i, sourceValue := range o.ServiceTiers {
		result.ServiceTiers[i] = sourceValue.Clone()
	}
	if o.RootOrganization == nil {
		result.RootOrganization = nil
	} else if data, err := o.RootOrganization.ProtoString(); err != nil {
		panic(err)
	} else {
		result.RootOrganization = &Name{}
		if err := result.RootOrganization.ParseProtoString(data); err != nil {
			panic(err)
		}
	}
	result.AncestryPath = make([]*Name, len(o.AncestryPath))
	for i, sourceValue := range o.AncestryPath {
		if sourceValue == nil {
			result.AncestryPath[i] = nil
		} else if data, err := sourceValue.ProtoString(); err != nil {
			panic(err)
		} else {
			result.AncestryPath[i] = &Name{}
			if err := result.AncestryPath[i].ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	}
	result.ServiceErrors = map[string]*iam_common.ServiceErrors{}
	for key, sourceValue := range o.ServiceErrors {
		result.ServiceErrors[key] = sourceValue.Clone()
	}
	result.ServicesGeneration = o.ServicesGeneration
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
	if source.GetMetadata() != nil {
		if o.Metadata == nil {
			o.Metadata = new(meta.Meta)
		}
		o.Metadata.Merge(source.GetMetadata())
	}
	o.Title = source.GetTitle()
	o.Description = source.GetDescription()
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
	if source.GetMultiRegionPolicy() != nil {
		if o.MultiRegionPolicy == nil {
			o.MultiRegionPolicy = new(multi_region_policy.MultiRegionPolicy)
		}
		o.MultiRegionPolicy.Merge(source.GetMultiRegionPolicy())
	}
	for _, sourceValue := range source.GetAllowedServices() {
		exists := false
		for _, currentValue := range o.AllowedServices {
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
			o.AllowedServices = append(o.AllowedServices, newDstElement)
		}
	}

	o.BusinessTier = source.GetBusinessTier()
	for _, sourceValue := range source.GetServiceTiers() {
		exists := false
		for _, currentValue := range o.ServiceTiers {
			if proto.Equal(sourceValue, currentValue) {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement *iam_common.ServiceBusinessTier
			if sourceValue != nil {
				newDstElement = new(iam_common.ServiceBusinessTier)
				newDstElement.Merge(sourceValue)
			}
			o.ServiceTiers = append(o.ServiceTiers, newDstElement)
		}
	}

	if source.GetRootOrganization() != nil {
		if data, err := source.GetRootOrganization().ProtoString(); err != nil {
			panic(err)
		} else {
			o.RootOrganization = &Name{}
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
			var newDstElement *Name
			if sourceValue != nil {
				if data, err := sourceValue.ProtoString(); err != nil {
					panic(err)
				} else {
					newDstElement = &Name{}
					if err := newDstElement.ParseProtoString(data); err != nil {
						panic(err)
					}
				}
			}
			o.AncestryPath = append(o.AncestryPath, newDstElement)
		}
	}

	if source.GetServiceErrors() != nil {
		if o.ServiceErrors == nil {
			o.ServiceErrors = make(map[string]*iam_common.ServiceErrors, len(source.GetServiceErrors()))
		}
		for key, sourceValue := range source.GetServiceErrors() {
			if sourceValue != nil {
				if o.ServiceErrors[key] == nil {
					o.ServiceErrors[key] = new(iam_common.ServiceErrors)
				}
				o.ServiceErrors[key].Merge(sourceValue)
			}
		}
	}
	o.ServicesGeneration = source.GetServicesGeneration()
}

func (o *Organization) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*Organization))
}
