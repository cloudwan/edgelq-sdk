// Code generated by protoc-gen-goten-object
// File: edgelq/iam/proto/v1/role.proto
// DO NOT EDIT!!!

package role

import (
	"fmt"
	"sort"

	"google.golang.org/protobuf/proto"
	googlefieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"

	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
)

// proto imports
import (
	condition "github.com/cloudwan/edgelq-sdk/iam/resources/v1/condition"
	permission "github.com/cloudwan/edgelq-sdk/iam/resources/v1/permission"
	meta_service "github.com/cloudwan/goten-sdk/meta-service/resources/v1/service"
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
	_ = &condition.Condition{}
	_ = &permission.Permission{}
	_ = &meta_service.Service{}
	_ = &meta.Meta{}
)

func (o *Role) GotenObjectExt() {}

func (o *Role) MakeFullFieldMask() *Role_FieldMask {
	return FullRole_FieldMask()
}

func (o *Role) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullRole_FieldMask()
}

func (o *Role) MakeDiffFieldMask(other *Role) *Role_FieldMask {
	if o == nil && other == nil {
		return &Role_FieldMask{}
	}
	if o == nil || other == nil {
		return FullRole_FieldMask()
	}

	res := &Role_FieldMask{}
	if o.GetName().String() != other.GetName().String() {
		res.Paths = append(res.Paths, &Role_FieldTerminalPath{selector: Role_FieldPathSelectorName})
	}
	{
		subMask := o.GetMetadata().MakeDiffFieldMask(other.GetMetadata())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &Role_FieldTerminalPath{selector: Role_FieldPathSelectorMetadata})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &Role_FieldSubPath{selector: Role_FieldPathSelectorMetadata, subPath: subpath})
			}
		}
	}
	if o.GetDisplayName() != other.GetDisplayName() {
		res.Paths = append(res.Paths, &Role_FieldTerminalPath{selector: Role_FieldPathSelectorDisplayName})
	}

	if len(o.GetScopeParams()) == len(other.GetScopeParams()) {
		for i, lValue := range o.GetScopeParams() {
			rValue := other.GetScopeParams()[i]
			if len(lValue.MakeDiffFieldMask(rValue).Paths) > 0 {
				res.Paths = append(res.Paths, &Role_FieldTerminalPath{selector: Role_FieldPathSelectorScopeParams})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &Role_FieldTerminalPath{selector: Role_FieldPathSelectorScopeParams})
	}

	if len(o.GetGrants()) == len(other.GetGrants()) {
		for i, lValue := range o.GetGrants() {
			rValue := other.GetGrants()[i]
			if len(lValue.MakeDiffFieldMask(rValue).Paths) > 0 {
				res.Paths = append(res.Paths, &Role_FieldTerminalPath{selector: Role_FieldPathSelectorGrants})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &Role_FieldTerminalPath{selector: Role_FieldPathSelectorGrants})
	}

	if len(o.GetOwnedObjects()) == len(other.GetOwnedObjects()) {
		for i, lValue := range o.GetOwnedObjects() {
			rValue := other.GetOwnedObjects()[i]
			if lValue != rValue {
				res.Paths = append(res.Paths, &Role_FieldTerminalPath{selector: Role_FieldPathSelectorOwnedObjects})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &Role_FieldTerminalPath{selector: Role_FieldPathSelectorOwnedObjects})
	}

	if len(o.GetServices()) == len(other.GetServices()) {
		for i, lValue := range o.GetServices() {
			rValue := other.GetServices()[i]
			if lValue.String() != rValue.String() {
				res.Paths = append(res.Paths, &Role_FieldTerminalPath{selector: Role_FieldPathSelectorServices})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &Role_FieldTerminalPath{selector: Role_FieldPathSelectorServices})
	}
	if o.GetRbSpecGeneration() != other.GetRbSpecGeneration() {
		res.Paths = append(res.Paths, &Role_FieldTerminalPath{selector: Role_FieldPathSelectorRbSpecGeneration})
	}
	return res
}

func (o *Role) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*Role))
}

func (o *Role) Clone() *Role {
	if o == nil {
		return nil
	}
	result := &Role{}
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
	result.ScopeParams = make([]*Role_ScopeParamType, len(o.ScopeParams))
	for i, sourceValue := range o.ScopeParams {
		result.ScopeParams[i] = sourceValue.Clone()
	}
	result.Grants = make([]*Role_Grant, len(o.Grants))
	for i, sourceValue := range o.Grants {
		result.Grants[i] = sourceValue.Clone()
	}
	result.OwnedObjects = make([]string, len(o.OwnedObjects))
	for i, sourceValue := range o.OwnedObjects {
		result.OwnedObjects[i] = sourceValue
	}
	result.Services = make([]*meta_service.Reference, len(o.Services))
	for i, sourceValue := range o.Services {
		if sourceValue == nil {
			result.Services[i] = nil
		} else if data, err := sourceValue.ProtoString(); err != nil {
			panic(err)
		} else {
			result.Services[i] = &meta_service.Reference{}
			if err := result.Services[i].ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	}
	result.RbSpecGeneration = o.RbSpecGeneration
	return result
}

func (o *Role) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *Role) Merge(source *Role) {
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
	for _, sourceValue := range source.GetScopeParams() {
		exists := false
		for _, currentValue := range o.ScopeParams {
			if proto.Equal(sourceValue, currentValue) {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement *Role_ScopeParamType
			if sourceValue != nil {
				newDstElement = new(Role_ScopeParamType)
				newDstElement.Merge(sourceValue)
			}
			o.ScopeParams = append(o.ScopeParams, newDstElement)
		}
	}

	for _, sourceValue := range source.GetGrants() {
		exists := false
		for _, currentValue := range o.Grants {
			if proto.Equal(sourceValue, currentValue) {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement *Role_Grant
			if sourceValue != nil {
				newDstElement = new(Role_Grant)
				newDstElement.Merge(sourceValue)
			}
			o.Grants = append(o.Grants, newDstElement)
		}
	}

	for _, sourceValue := range source.GetOwnedObjects() {
		exists := false
		for _, currentValue := range o.OwnedObjects {
			if currentValue == sourceValue {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement string
			newDstElement = sourceValue
			o.OwnedObjects = append(o.OwnedObjects, newDstElement)
		}
	}

	for _, sourceValue := range source.GetServices() {
		exists := false
		for _, currentValue := range o.Services {
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
			o.Services = append(o.Services, newDstElement)
		}
	}

	o.RbSpecGeneration = source.GetRbSpecGeneration()
}

func (o *Role) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*Role))
}

func (o *Role_ScopeParamType) GotenObjectExt() {}

func (o *Role_ScopeParamType) MakeFullFieldMask() *Role_ScopeParamType_FieldMask {
	return FullRole_ScopeParamType_FieldMask()
}

func (o *Role_ScopeParamType) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullRole_ScopeParamType_FieldMask()
}

func (o *Role_ScopeParamType) MakeDiffFieldMask(other *Role_ScopeParamType) *Role_ScopeParamType_FieldMask {
	if o == nil && other == nil {
		return &Role_ScopeParamType_FieldMask{}
	}
	if o == nil || other == nil {
		return FullRole_ScopeParamType_FieldMask()
	}

	res := &Role_ScopeParamType_FieldMask{}
	if o.GetName() != other.GetName() {
		res.Paths = append(res.Paths, &RoleScopeParamType_FieldTerminalPath{selector: RoleScopeParamType_FieldPathSelectorName})
	}
	if o.GetType() != other.GetType() {
		res.Paths = append(res.Paths, &RoleScopeParamType_FieldTerminalPath{selector: RoleScopeParamType_FieldPathSelectorType})
	}
	return res
}

func (o *Role_ScopeParamType) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*Role_ScopeParamType))
}

func (o *Role_ScopeParamType) Clone() *Role_ScopeParamType {
	if o == nil {
		return nil
	}
	result := &Role_ScopeParamType{}
	result.Name = o.Name
	result.Type = o.Type
	return result
}

func (o *Role_ScopeParamType) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *Role_ScopeParamType) Merge(source *Role_ScopeParamType) {
	o.Name = source.GetName()
	o.Type = source.GetType()
}

func (o *Role_ScopeParamType) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*Role_ScopeParamType))
}

func (o *Role_Grant) GotenObjectExt() {}

func (o *Role_Grant) MakeFullFieldMask() *Role_Grant_FieldMask {
	return FullRole_Grant_FieldMask()
}

func (o *Role_Grant) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullRole_Grant_FieldMask()
}

func (o *Role_Grant) MakeDiffFieldMask(other *Role_Grant) *Role_Grant_FieldMask {
	if o == nil && other == nil {
		return &Role_Grant_FieldMask{}
	}
	if o == nil || other == nil {
		return FullRole_Grant_FieldMask()
	}

	res := &Role_Grant_FieldMask{}
	if o.GetSubScope() != other.GetSubScope() {
		res.Paths = append(res.Paths, &RoleGrant_FieldTerminalPath{selector: RoleGrant_FieldPathSelectorSubScope})
	}

	if len(o.GetPermissions()) == len(other.GetPermissions()) {
		for i, lValue := range o.GetPermissions() {
			rValue := other.GetPermissions()[i]
			if lValue.String() != rValue.String() {
				res.Paths = append(res.Paths, &RoleGrant_FieldTerminalPath{selector: RoleGrant_FieldPathSelectorPermissions})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &RoleGrant_FieldTerminalPath{selector: RoleGrant_FieldPathSelectorPermissions})
	}

	if len(o.GetResourceFieldConditions()) == len(other.GetResourceFieldConditions()) {
		for i, lValue := range o.GetResourceFieldConditions() {
			rValue := other.GetResourceFieldConditions()[i]
			if len(lValue.MakeDiffFieldMask(rValue).Paths) > 0 {
				res.Paths = append(res.Paths, &RoleGrant_FieldTerminalPath{selector: RoleGrant_FieldPathSelectorResourceFieldConditions})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &RoleGrant_FieldTerminalPath{selector: RoleGrant_FieldPathSelectorResourceFieldConditions})
	}

	if len(o.GetRequestFieldConditions()) == len(other.GetRequestFieldConditions()) {
		for i, lValue := range o.GetRequestFieldConditions() {
			rValue := other.GetRequestFieldConditions()[i]
			if len(lValue.MakeDiffFieldMask(rValue).Paths) > 0 {
				res.Paths = append(res.Paths, &RoleGrant_FieldTerminalPath{selector: RoleGrant_FieldPathSelectorRequestFieldConditions})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &RoleGrant_FieldTerminalPath{selector: RoleGrant_FieldPathSelectorRequestFieldConditions})
	}

	if len(o.GetExecutableConditions()) == len(other.GetExecutableConditions()) {
		for i, lValue := range o.GetExecutableConditions() {
			rValue := other.GetExecutableConditions()[i]
			if len(lValue.MakeDiffFieldMask(rValue).Paths) > 0 {
				res.Paths = append(res.Paths, &RoleGrant_FieldTerminalPath{selector: RoleGrant_FieldPathSelectorExecutableConditions})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &RoleGrant_FieldTerminalPath{selector: RoleGrant_FieldPathSelectorExecutableConditions})
	}
	return res
}

func (o *Role_Grant) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*Role_Grant))
}

func (o *Role_Grant) Clone() *Role_Grant {
	if o == nil {
		return nil
	}
	result := &Role_Grant{}
	result.SubScope = o.SubScope
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
	result.ResourceFieldConditions = make([]*Role_Grant_FieldCondition, len(o.ResourceFieldConditions))
	for i, sourceValue := range o.ResourceFieldConditions {
		result.ResourceFieldConditions[i] = sourceValue.Clone()
	}
	result.RequestFieldConditions = make([]*Role_Grant_FieldCondition, len(o.RequestFieldConditions))
	for i, sourceValue := range o.RequestFieldConditions {
		result.RequestFieldConditions[i] = sourceValue.Clone()
	}
	result.ExecutableConditions = make([]*condition.ExecutableCondition, len(o.ExecutableConditions))
	for i, sourceValue := range o.ExecutableConditions {
		result.ExecutableConditions[i] = sourceValue.Clone()
	}
	return result
}

func (o *Role_Grant) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *Role_Grant) Merge(source *Role_Grant) {
	o.SubScope = source.GetSubScope()
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

	for _, sourceValue := range source.GetResourceFieldConditions() {
		exists := false
		for _, currentValue := range o.ResourceFieldConditions {
			if proto.Equal(sourceValue, currentValue) {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement *Role_Grant_FieldCondition
			if sourceValue != nil {
				newDstElement = new(Role_Grant_FieldCondition)
				newDstElement.Merge(sourceValue)
			}
			o.ResourceFieldConditions = append(o.ResourceFieldConditions, newDstElement)
		}
	}

	for _, sourceValue := range source.GetRequestFieldConditions() {
		exists := false
		for _, currentValue := range o.RequestFieldConditions {
			if proto.Equal(sourceValue, currentValue) {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement *Role_Grant_FieldCondition
			if sourceValue != nil {
				newDstElement = new(Role_Grant_FieldCondition)
				newDstElement.Merge(sourceValue)
			}
			o.RequestFieldConditions = append(o.RequestFieldConditions, newDstElement)
		}
	}

	for _, sourceValue := range source.GetExecutableConditions() {
		exists := false
		for _, currentValue := range o.ExecutableConditions {
			if proto.Equal(sourceValue, currentValue) {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement *condition.ExecutableCondition
			if sourceValue != nil {
				newDstElement = new(condition.ExecutableCondition)
				newDstElement.Merge(sourceValue)
			}
			o.ExecutableConditions = append(o.ExecutableConditions, newDstElement)
		}
	}

}

func (o *Role_Grant) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*Role_Grant))
}

func (o *Role_Grant_FieldCondition) GotenObjectExt() {}

func (o *Role_Grant_FieldCondition) MakeFullFieldMask() *Role_Grant_FieldCondition_FieldMask {
	return FullRole_Grant_FieldCondition_FieldMask()
}

func (o *Role_Grant_FieldCondition) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullRole_Grant_FieldCondition_FieldMask()
}

func (o *Role_Grant_FieldCondition) MakeDiffFieldMask(other *Role_Grant_FieldCondition) *Role_Grant_FieldCondition_FieldMask {
	if o == nil && other == nil {
		return &Role_Grant_FieldCondition_FieldMask{}
	}
	if o == nil || other == nil {
		return FullRole_Grant_FieldCondition_FieldMask()
	}

	res := &Role_Grant_FieldCondition_FieldMask{}
	if o.GetPath() != other.GetPath() {
		res.Paths = append(res.Paths, &RoleGrantFieldCondition_FieldTerminalPath{selector: RoleGrantFieldCondition_FieldPathSelectorPath})
	}
	if o.GetValue() != other.GetValue() {
		res.Paths = append(res.Paths, &RoleGrantFieldCondition_FieldTerminalPath{selector: RoleGrantFieldCondition_FieldPathSelectorValue})
	}
	return res
}

func (o *Role_Grant_FieldCondition) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*Role_Grant_FieldCondition))
}

func (o *Role_Grant_FieldCondition) Clone() *Role_Grant_FieldCondition {
	if o == nil {
		return nil
	}
	result := &Role_Grant_FieldCondition{}
	result.Path = o.Path
	result.Value = o.Value
	return result
}

func (o *Role_Grant_FieldCondition) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *Role_Grant_FieldCondition) Merge(source *Role_Grant_FieldCondition) {
	o.Path = source.GetPath()
	o.Value = source.GetValue()
}

func (o *Role_Grant_FieldCondition) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*Role_Grant_FieldCondition))
}

func (o *ScopeParam) GotenObjectExt() {}

func (o *ScopeParam) MakeFullFieldMask() *ScopeParam_FieldMask {
	return FullScopeParam_FieldMask()
}

func (o *ScopeParam) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullScopeParam_FieldMask()
}

func (o *ScopeParam) MakeDiffFieldMask(other *ScopeParam) *ScopeParam_FieldMask {
	if o == nil && other == nil {
		return &ScopeParam_FieldMask{}
	}
	if o == nil || other == nil {
		return FullScopeParam_FieldMask()
	}

	res := &ScopeParam_FieldMask{}
	if o.GetName() != other.GetName() {
		res.Paths = append(res.Paths, &ScopeParam_FieldTerminalPath{selector: ScopeParam_FieldPathSelectorName})
	}
	{
		subMask := o.GetString_().MakeDiffFieldMask(other.GetString_())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &ScopeParam_FieldTerminalPath{selector: ScopeParam_FieldPathSelectorString})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &ScopeParam_FieldSubPath{selector: ScopeParam_FieldPathSelectorString, subPath: subpath})
			}
		}
	}
	{
		subMask := o.GetStrings().MakeDiffFieldMask(other.GetStrings())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &ScopeParam_FieldTerminalPath{selector: ScopeParam_FieldPathSelectorStrings})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &ScopeParam_FieldSubPath{selector: ScopeParam_FieldPathSelectorStrings, subPath: subpath})
			}
		}
	}
	return res
}

func (o *ScopeParam) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*ScopeParam))
}

func (o *ScopeParam) Clone() *ScopeParam {
	if o == nil {
		return nil
	}
	result := &ScopeParam{}
	result.Name = o.Name
	if o, ok := o.Value.(*ScopeParam_String_); ok {
		result.Value = (*ScopeParam_String_)(nil)
		if o != nil {
			result.Value = &ScopeParam_String_{}
			result := result.Value.(*ScopeParam_String_)
			result.String_ = o.String_.Clone()
		}
	}
	if o, ok := o.Value.(*ScopeParam_Strings); ok {
		result.Value = (*ScopeParam_Strings)(nil)
		if o != nil {
			result.Value = &ScopeParam_Strings{}
			result := result.Value.(*ScopeParam_Strings)
			result.Strings = o.Strings.Clone()
		}
	}
	return result
}

func (o *ScopeParam) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *ScopeParam) Merge(source *ScopeParam) {
	o.Name = source.GetName()
	if source, ok := source.GetValue().(*ScopeParam_String_); ok {
		if dstOneOf, ok := o.Value.(*ScopeParam_String_); !ok || dstOneOf == nil {
			o.Value = &ScopeParam_String_{}
		}
		if source != nil {
			o := o.Value.(*ScopeParam_String_)
			if source.String_ != nil {
				if o.String_ == nil {
					o.String_ = new(ScopeParam_StringValue)
				}
				o.String_.Merge(source.String_)
			}
		}
	}
	if source, ok := source.GetValue().(*ScopeParam_Strings); ok {
		if dstOneOf, ok := o.Value.(*ScopeParam_Strings); !ok || dstOneOf == nil {
			o.Value = &ScopeParam_Strings{}
		}
		if source != nil {
			o := o.Value.(*ScopeParam_Strings)
			if source.Strings != nil {
				if o.Strings == nil {
					o.Strings = new(ScopeParam_ArrayOfStringsValue)
				}
				o.Strings.Merge(source.Strings)
			}
		}
	}
}

func (o *ScopeParam) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*ScopeParam))
}

func (o *ScopeParam_StringValue) GotenObjectExt() {}

func (o *ScopeParam_StringValue) MakeFullFieldMask() *ScopeParam_StringValue_FieldMask {
	return FullScopeParam_StringValue_FieldMask()
}

func (o *ScopeParam_StringValue) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullScopeParam_StringValue_FieldMask()
}

func (o *ScopeParam_StringValue) MakeDiffFieldMask(other *ScopeParam_StringValue) *ScopeParam_StringValue_FieldMask {
	if o == nil && other == nil {
		return &ScopeParam_StringValue_FieldMask{}
	}
	if o == nil || other == nil {
		return FullScopeParam_StringValue_FieldMask()
	}

	res := &ScopeParam_StringValue_FieldMask{}
	if o.GetValue() != other.GetValue() {
		res.Paths = append(res.Paths, &ScopeParamStringValue_FieldTerminalPath{selector: ScopeParamStringValue_FieldPathSelectorValue})
	}
	return res
}

func (o *ScopeParam_StringValue) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*ScopeParam_StringValue))
}

func (o *ScopeParam_StringValue) Clone() *ScopeParam_StringValue {
	if o == nil {
		return nil
	}
	result := &ScopeParam_StringValue{}
	result.Value = o.Value
	return result
}

func (o *ScopeParam_StringValue) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *ScopeParam_StringValue) Merge(source *ScopeParam_StringValue) {
	o.Value = source.GetValue()
}

func (o *ScopeParam_StringValue) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*ScopeParam_StringValue))
}

func (o *ScopeParam_ArrayOfStringsValue) GotenObjectExt() {}

func (o *ScopeParam_ArrayOfStringsValue) MakeFullFieldMask() *ScopeParam_ArrayOfStringsValue_FieldMask {
	return FullScopeParam_ArrayOfStringsValue_FieldMask()
}

func (o *ScopeParam_ArrayOfStringsValue) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullScopeParam_ArrayOfStringsValue_FieldMask()
}

func (o *ScopeParam_ArrayOfStringsValue) MakeDiffFieldMask(other *ScopeParam_ArrayOfStringsValue) *ScopeParam_ArrayOfStringsValue_FieldMask {
	if o == nil && other == nil {
		return &ScopeParam_ArrayOfStringsValue_FieldMask{}
	}
	if o == nil || other == nil {
		return FullScopeParam_ArrayOfStringsValue_FieldMask()
	}

	res := &ScopeParam_ArrayOfStringsValue_FieldMask{}

	if len(o.GetValues()) == len(other.GetValues()) {
		for i, lValue := range o.GetValues() {
			rValue := other.GetValues()[i]
			if lValue != rValue {
				res.Paths = append(res.Paths, &ScopeParamArrayOfStringsValue_FieldTerminalPath{selector: ScopeParamArrayOfStringsValue_FieldPathSelectorValues})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &ScopeParamArrayOfStringsValue_FieldTerminalPath{selector: ScopeParamArrayOfStringsValue_FieldPathSelectorValues})
	}
	return res
}

func (o *ScopeParam_ArrayOfStringsValue) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*ScopeParam_ArrayOfStringsValue))
}

func (o *ScopeParam_ArrayOfStringsValue) Clone() *ScopeParam_ArrayOfStringsValue {
	if o == nil {
		return nil
	}
	result := &ScopeParam_ArrayOfStringsValue{}
	result.Values = make([]string, len(o.Values))
	for i, sourceValue := range o.Values {
		result.Values[i] = sourceValue
	}
	return result
}

func (o *ScopeParam_ArrayOfStringsValue) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *ScopeParam_ArrayOfStringsValue) Merge(source *ScopeParam_ArrayOfStringsValue) {
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

func (o *ScopeParam_ArrayOfStringsValue) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*ScopeParam_ArrayOfStringsValue))
}
