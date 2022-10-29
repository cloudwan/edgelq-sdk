// Code generated by protoc-gen-goten-object
// File: edgelq/iam/proto/v1alpha/role_binding.proto
// DO NOT EDIT!!!

package role_binding

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
	condition "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/condition"
	organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/organization"
	project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/project"
	role "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/role"
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
	_ = &condition.Condition{}
	_ = &organization.Organization{}
	_ = &project.Project{}
	_ = &role.Role{}
)

func (o *RoleBinding) GotenObjectExt() {}

func (o *RoleBinding) MakeFullFieldMask() *RoleBinding_FieldMask {
	return FullRoleBinding_FieldMask()
}

func (o *RoleBinding) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullRoleBinding_FieldMask()
}

func (o *RoleBinding) MakeDiffFieldMask(other *RoleBinding) *RoleBinding_FieldMask {
	if o == nil && other == nil {
		return &RoleBinding_FieldMask{}
	}
	if o == nil || other == nil {
		return FullRoleBinding_FieldMask()
	}

	res := &RoleBinding_FieldMask{}
	if o.GetName().String() != other.GetName().String() {
		res.Paths = append(res.Paths, &RoleBinding_FieldTerminalPath{selector: RoleBinding_FieldPathSelectorName})
	}
	if o.GetRole().String() != other.GetRole().String() {
		res.Paths = append(res.Paths, &RoleBinding_FieldTerminalPath{selector: RoleBinding_FieldPathSelectorRole})
	}
	if o.GetMember() != other.GetMember() {
		res.Paths = append(res.Paths, &RoleBinding_FieldTerminalPath{selector: RoleBinding_FieldPathSelectorMember})
	}
	{
		subMask := o.GetConditionBinding().MakeDiffFieldMask(other.GetConditionBinding())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &RoleBinding_FieldTerminalPath{selector: RoleBinding_FieldPathSelectorConditionBinding})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &RoleBinding_FieldSubPath{selector: RoleBinding_FieldPathSelectorConditionBinding, subPath: subpath})
			}
		}
	}

	if len(o.GetAncestryPath()) == len(other.GetAncestryPath()) {
		for i, lValue := range o.GetAncestryPath() {
			rValue := other.GetAncestryPath()[i]
			if len(lValue.MakeDiffFieldMask(rValue).Paths) > 0 {
				res.Paths = append(res.Paths, &RoleBinding_FieldTerminalPath{selector: RoleBinding_FieldPathSelectorAncestryPath})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &RoleBinding_FieldTerminalPath{selector: RoleBinding_FieldPathSelectorAncestryPath})
	}
	{
		subMask := o.GetMetadata().MakeDiffFieldMask(other.GetMetadata())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &RoleBinding_FieldTerminalPath{selector: RoleBinding_FieldPathSelectorMetadata})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &RoleBinding_FieldSubPath{selector: RoleBinding_FieldPathSelectorMetadata, subPath: subpath})
			}
		}
	}
	return res
}

func (o *RoleBinding) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*RoleBinding))
}

func (o *RoleBinding) Clone() *RoleBinding {
	if o == nil {
		return nil
	}
	result := &RoleBinding{}
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
	if o.Role == nil {
		result.Role = nil
	} else if data, err := o.Role.ProtoString(); err != nil {
		panic(err)
	} else {
		result.Role = &role.Reference{}
		if err := result.Role.ParseProtoString(data); err != nil {
			panic(err)
		}
	}
	result.Member = o.Member
	result.ConditionBinding = o.ConditionBinding.Clone()
	result.AncestryPath = make([]*RoleBinding_Parent, len(o.AncestryPath))
	for i, sourceValue := range o.AncestryPath {
		result.AncestryPath[i] = sourceValue.Clone()
	}
	result.Metadata = o.Metadata.Clone()
	return result
}

func (o *RoleBinding) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *RoleBinding) Merge(source *RoleBinding) {
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
	if source.GetRole() != nil {
		if data, err := source.GetRole().ProtoString(); err != nil {
			panic(err)
		} else {
			o.Role = &role.Reference{}
			if err := o.Role.ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	} else {
		o.Role = nil
	}
	o.Member = source.GetMember()
	if source.GetConditionBinding() != nil {
		if o.ConditionBinding == nil {
			o.ConditionBinding = new(condition.ConditionBinding)
		}
		o.ConditionBinding.Merge(source.GetConditionBinding())
	}
	for _, sourceValue := range source.GetAncestryPath() {
		exists := false
		for _, currentValue := range o.AncestryPath {
			if proto.Equal(sourceValue, currentValue) {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement *RoleBinding_Parent
			if sourceValue != nil {
				newDstElement = new(RoleBinding_Parent)
				newDstElement.Merge(sourceValue)
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
}

func (o *RoleBinding) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*RoleBinding))
}

func (o *RoleBinding_Parent) GotenObjectExt() {}

func (o *RoleBinding_Parent) MakeFullFieldMask() *RoleBinding_Parent_FieldMask {
	return FullRoleBinding_Parent_FieldMask()
}

func (o *RoleBinding_Parent) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullRoleBinding_Parent_FieldMask()
}

func (o *RoleBinding_Parent) MakeDiffFieldMask(other *RoleBinding_Parent) *RoleBinding_Parent_FieldMask {
	if o == nil && other == nil {
		return &RoleBinding_Parent_FieldMask{}
	}
	if o == nil || other == nil {
		return FullRoleBinding_Parent_FieldMask()
	}

	res := &RoleBinding_Parent_FieldMask{}
	if o.GetParent().String() != other.GetParent().String() {
		res.Paths = append(res.Paths, &RoleBindingParent_FieldTerminalPath{selector: RoleBindingParent_FieldPathSelectorParent})
	}
	if o.GetMember() != other.GetMember() {
		res.Paths = append(res.Paths, &RoleBindingParent_FieldTerminalPath{selector: RoleBindingParent_FieldPathSelectorMember})
	}
	return res
}

func (o *RoleBinding_Parent) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*RoleBinding_Parent))
}

func (o *RoleBinding_Parent) Clone() *RoleBinding_Parent {
	if o == nil {
		return nil
	}
	result := &RoleBinding_Parent{}
	if o.Parent == nil {
		result.Parent = nil
	} else if data, err := o.Parent.ProtoString(); err != nil {
		panic(err)
	} else {
		result.Parent = &Reference{}
		if err := result.Parent.ParseProtoString(data); err != nil {
			panic(err)
		}
	}
	result.Member = o.Member
	return result
}

func (o *RoleBinding_Parent) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *RoleBinding_Parent) Merge(source *RoleBinding_Parent) {
	if source.GetParent() != nil {
		if data, err := source.GetParent().ProtoString(); err != nil {
			panic(err)
		} else {
			o.Parent = &Reference{}
			if err := o.Parent.ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	} else {
		o.Parent = nil
	}
	o.Member = source.GetMember()
}

func (o *RoleBinding_Parent) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*RoleBinding_Parent))
}
