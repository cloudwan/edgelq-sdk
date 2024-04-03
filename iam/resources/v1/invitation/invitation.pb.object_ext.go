// Code generated by protoc-gen-goten-object
// File: edgelq/iam/proto/v1/invitation.proto
// DO NOT EDIT!!!

package iam_invitation

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
	role "github.com/cloudwan/edgelq-sdk/iam/resources/v1/role"
	service_account "github.com/cloudwan/edgelq-sdk/iam/resources/v1/service_account"
	user "github.com/cloudwan/edgelq-sdk/iam/resources/v1/user"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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
	_ = &role.Role{}
	_ = &service_account.ServiceAccount{}
	_ = &user.User{}
	_ = &timestamppb.Timestamp{}
)

func (o *Actor) GotenObjectExt() {}

func (o *Actor) MakeFullFieldMask() *Actor_FieldMask {
	return FullActor_FieldMask()
}

func (o *Actor) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullActor_FieldMask()
}

func (o *Actor) MakeDiffFieldMask(other *Actor) *Actor_FieldMask {
	if o == nil && other == nil {
		return &Actor_FieldMask{}
	}
	if o == nil || other == nil {
		return FullActor_FieldMask()
	}

	res := &Actor_FieldMask{}
	if o.GetUser().String() != other.GetUser().String() {
		res.Paths = append(res.Paths, &Actor_FieldTerminalPath{selector: Actor_FieldPathSelectorUser})
	}
	if o.GetServiceAccount().String() != other.GetServiceAccount().String() {
		res.Paths = append(res.Paths, &Actor_FieldTerminalPath{selector: Actor_FieldPathSelectorServiceAccount})
	}
	return res
}

func (o *Actor) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*Actor))
}

func (o *Actor) Clone() *Actor {
	if o == nil {
		return nil
	}
	result := &Actor{}
	if o.User == nil {
		result.User = nil
	} else if data, err := o.User.ProtoString(); err != nil {
		panic(err)
	} else {
		result.User = &user.Reference{}
		if err := result.User.ParseProtoString(data); err != nil {
			panic(err)
		}
	}
	if o.ServiceAccount == nil {
		result.ServiceAccount = nil
	} else if data, err := o.ServiceAccount.ProtoString(); err != nil {
		panic(err)
	} else {
		result.ServiceAccount = &service_account.Reference{}
		if err := result.ServiceAccount.ParseProtoString(data); err != nil {
			panic(err)
		}
	}
	return result
}

func (o *Actor) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *Actor) Merge(source *Actor) {
	if source.GetUser() != nil {
		if data, err := source.GetUser().ProtoString(); err != nil {
			panic(err)
		} else {
			o.User = &user.Reference{}
			if err := o.User.ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	} else {
		o.User = nil
	}
	if source.GetServiceAccount() != nil {
		if data, err := source.GetServiceAccount().ProtoString(); err != nil {
			panic(err)
		} else {
			o.ServiceAccount = &service_account.Reference{}
			if err := o.ServiceAccount.ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	} else {
		o.ServiceAccount = nil
	}
}

func (o *Actor) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*Actor))
}

func (o *Invitation) GotenObjectExt() {}

func (o *Invitation) MakeFullFieldMask() *Invitation_FieldMask {
	return FullInvitation_FieldMask()
}

func (o *Invitation) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullInvitation_FieldMask()
}

func (o *Invitation) MakeDiffFieldMask(other *Invitation) *Invitation_FieldMask {
	if o == nil && other == nil {
		return &Invitation_FieldMask{}
	}
	if o == nil || other == nil {
		return FullInvitation_FieldMask()
	}

	res := &Invitation_FieldMask{}
	if o.GetInviteeEmail() != other.GetInviteeEmail() {
		res.Paths = append(res.Paths, &Invitation_FieldTerminalPath{selector: Invitation_FieldPathSelectorInviteeEmail})
	}
	{
		subMask := o.GetInviterActor().MakeDiffFieldMask(other.GetInviterActor())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &Invitation_FieldTerminalPath{selector: Invitation_FieldPathSelectorInviterActor})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &Invitation_FieldSubPath{selector: Invitation_FieldPathSelectorInviterActor, subPath: subpath})
			}
		}
	}
	if o.GetInviterFullName() != other.GetInviterFullName() {
		res.Paths = append(res.Paths, &Invitation_FieldTerminalPath{selector: Invitation_FieldPathSelectorInviterFullName})
	}
	if o.GetInviterEmail() != other.GetInviterEmail() {
		res.Paths = append(res.Paths, &Invitation_FieldTerminalPath{selector: Invitation_FieldPathSelectorInviterEmail})
	}
	if o.GetLanguageCode() != other.GetLanguageCode() {
		res.Paths = append(res.Paths, &Invitation_FieldTerminalPath{selector: Invitation_FieldPathSelectorLanguageCode})
	}

	if len(o.GetBindingRoles()) == len(other.GetBindingRoles()) {
		for i, lValue := range o.GetBindingRoles() {
			rValue := other.GetBindingRoles()[i]
			if len(lValue.MakeDiffFieldMask(rValue).Paths) > 0 {
				res.Paths = append(res.Paths, &Invitation_FieldTerminalPath{selector: Invitation_FieldPathSelectorBindingRoles})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &Invitation_FieldTerminalPath{selector: Invitation_FieldPathSelectorBindingRoles})
	}
	if !proto.Equal(o.GetExpirationDate(), other.GetExpirationDate()) {
		res.Paths = append(res.Paths, &Invitation_FieldTerminalPath{selector: Invitation_FieldPathSelectorExpirationDate})
	}

	if len(o.GetExtras()) == len(other.GetExtras()) {
		for i, lValue := range o.GetExtras() {
			rValue := other.GetExtras()[i]
			if lValue != rValue {
				res.Paths = append(res.Paths, &Invitation_FieldTerminalPath{selector: Invitation_FieldPathSelectorExtras})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &Invitation_FieldTerminalPath{selector: Invitation_FieldPathSelectorExtras})
	}
	if o.GetState() != other.GetState() {
		res.Paths = append(res.Paths, &Invitation_FieldTerminalPath{selector: Invitation_FieldPathSelectorState})
	}
	return res
}

func (o *Invitation) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*Invitation))
}

func (o *Invitation) Clone() *Invitation {
	if o == nil {
		return nil
	}
	result := &Invitation{}
	result.InviteeEmail = o.InviteeEmail
	result.InviterActor = o.InviterActor.Clone()
	result.InviterFullName = o.InviterFullName
	result.InviterEmail = o.InviterEmail
	result.LanguageCode = o.LanguageCode
	result.BindingRoles = make([]*Invitation_BindingRole, len(o.BindingRoles))
	for i, sourceValue := range o.BindingRoles {
		result.BindingRoles[i] = sourceValue.Clone()
	}
	result.ExpirationDate = proto.Clone(o.ExpirationDate).(*timestamppb.Timestamp)
	result.Extras = map[string]string{}
	for key, sourceValue := range o.Extras {
		result.Extras[key] = sourceValue
	}
	result.State = o.State
	return result
}

func (o *Invitation) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *Invitation) Merge(source *Invitation) {
	o.InviteeEmail = source.GetInviteeEmail()
	if source.GetInviterActor() != nil {
		if o.InviterActor == nil {
			o.InviterActor = new(Actor)
		}
		o.InviterActor.Merge(source.GetInviterActor())
	}
	o.InviterFullName = source.GetInviterFullName()
	o.InviterEmail = source.GetInviterEmail()
	o.LanguageCode = source.GetLanguageCode()
	for _, sourceValue := range source.GetBindingRoles() {
		exists := false
		for _, currentValue := range o.BindingRoles {
			if proto.Equal(sourceValue, currentValue) {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement *Invitation_BindingRole
			if sourceValue != nil {
				newDstElement = new(Invitation_BindingRole)
				newDstElement.Merge(sourceValue)
			}
			o.BindingRoles = append(o.BindingRoles, newDstElement)
		}
	}

	if source.GetExpirationDate() != nil {
		if o.ExpirationDate == nil {
			o.ExpirationDate = new(timestamppb.Timestamp)
		}
		proto.Merge(o.ExpirationDate, source.GetExpirationDate())
	}
	if source.GetExtras() != nil {
		if o.Extras == nil {
			o.Extras = make(map[string]string, len(source.GetExtras()))
		}
		for key, sourceValue := range source.GetExtras() {
			o.Extras[key] = sourceValue
		}
	}
	o.State = source.GetState()
}

func (o *Invitation) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*Invitation))
}

func (o *Invitation_BindingRole) GotenObjectExt() {}

func (o *Invitation_BindingRole) MakeFullFieldMask() *Invitation_BindingRole_FieldMask {
	return FullInvitation_BindingRole_FieldMask()
}

func (o *Invitation_BindingRole) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullInvitation_BindingRole_FieldMask()
}

func (o *Invitation_BindingRole) MakeDiffFieldMask(other *Invitation_BindingRole) *Invitation_BindingRole_FieldMask {
	if o == nil && other == nil {
		return &Invitation_BindingRole_FieldMask{}
	}
	if o == nil || other == nil {
		return FullInvitation_BindingRole_FieldMask()
	}

	res := &Invitation_BindingRole_FieldMask{}
	if o.GetRole().String() != other.GetRole().String() {
		res.Paths = append(res.Paths, &InvitationBindingRole_FieldTerminalPath{selector: InvitationBindingRole_FieldPathSelectorRole})
	}

	if len(o.GetExecutableConditions()) == len(other.GetExecutableConditions()) {
		for i, lValue := range o.GetExecutableConditions() {
			rValue := other.GetExecutableConditions()[i]
			if len(lValue.MakeDiffFieldMask(rValue).Paths) > 0 {
				res.Paths = append(res.Paths, &InvitationBindingRole_FieldTerminalPath{selector: InvitationBindingRole_FieldPathSelectorExecutableConditions})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &InvitationBindingRole_FieldTerminalPath{selector: InvitationBindingRole_FieldPathSelectorExecutableConditions})
	}

	if len(o.GetScopeParams()) == len(other.GetScopeParams()) {
		for i, lValue := range o.GetScopeParams() {
			rValue := other.GetScopeParams()[i]
			if len(lValue.MakeDiffFieldMask(rValue).Paths) > 0 {
				res.Paths = append(res.Paths, &InvitationBindingRole_FieldTerminalPath{selector: InvitationBindingRole_FieldPathSelectorScopeParams})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &InvitationBindingRole_FieldTerminalPath{selector: InvitationBindingRole_FieldPathSelectorScopeParams})
	}
	return res
}

func (o *Invitation_BindingRole) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*Invitation_BindingRole))
}

func (o *Invitation_BindingRole) Clone() *Invitation_BindingRole {
	if o == nil {
		return nil
	}
	result := &Invitation_BindingRole{}
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
	result.ExecutableConditions = make([]*condition.ExecutableCondition, len(o.ExecutableConditions))
	for i, sourceValue := range o.ExecutableConditions {
		result.ExecutableConditions[i] = sourceValue.Clone()
	}
	result.ScopeParams = make([]*role.ScopeParam, len(o.ScopeParams))
	for i, sourceValue := range o.ScopeParams {
		result.ScopeParams[i] = sourceValue.Clone()
	}
	return result
}

func (o *Invitation_BindingRole) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *Invitation_BindingRole) Merge(source *Invitation_BindingRole) {
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

	for _, sourceValue := range source.GetScopeParams() {
		exists := false
		for _, currentValue := range o.ScopeParams {
			if proto.Equal(sourceValue, currentValue) {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement *role.ScopeParam
			if sourceValue != nil {
				newDstElement = new(role.ScopeParam)
				newDstElement.Merge(sourceValue)
			}
			o.ScopeParams = append(o.ScopeParams, newDstElement)
		}
	}

}

func (o *Invitation_BindingRole) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*Invitation_BindingRole))
}